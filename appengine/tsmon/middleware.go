// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tsmon

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/info"
	gaeauth "github.com/luci/luci-go/appengine/gaeauth/client"
	"github.com/luci/luci-go/common/clock"
	gcps "github.com/luci/luci-go/common/gcloud/pubsub"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/tsmon"
	"github.com/luci/luci-go/common/tsmon/monitor"
	"github.com/luci/luci-go/common/tsmon/store"
	"github.com/luci/luci-go/common/tsmon/target"
	"github.com/luci/luci-go/server/middleware"
	"golang.org/x/net/context"
)

var (
	lastFlushed = struct {
		time.Time
		sync.Mutex
	}{}
)

// Middleware returns a middleware that must be inserted into the chain to
// enable tsmon metrics to be sent on App Engine.
func Middleware(h middleware.Handler) middleware.Handler {
	return func(c context.Context, rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if store.IsNilStore(tsmon.GetState(c).S) {
			if err := initialize(c); err != nil {
				logging.Errorf(c, "Failed to initialize tsmon: %s", err)
				// Don't fail the request.
			}
		}
		h(c, rw, r, p)
		flushIfNeeded(c)
	}
}

func initialize(c context.Context) error {
	var mon monitor.Monitor
	i := info.Get(c)
	if i.IsDevAppServer() {
		mon = monitor.NewDebugMonitor("")
	} else {
		// Create an HTTP client with the default appengine service account.
		auth, err := gaeauth.Authenticator(c, gcps.PublisherScopes, nil)
		if err != nil {
			return err
		}
		client, err := auth.Client()
		if err != nil {
			return err
		}

		mon, err = monitor.NewPubsubMonitor(c, client, gcps.NewTopic(pubsubProject, pubsubTopic))
		if err != nil {
			return err
		}
	}

	// Create the target.
	tar := &target.Task{
		DataCenter:  targetDataCenter,
		ServiceName: i.AppID(),
		JobName:     i.ModuleName(),
		HostName:    strings.SplitN(i.VersionID(), ".", 2)[0],
		TaskNum:     -1,
	}

	tsmon.Initialize(c, mon, store.NewInMemory(tar))
	return nil
}

func flushIfNeeded(c context.Context) {
	if !updateLastFlushed(c) {
		return
	}

	if err := updateInstanceEntityAndFlush(c); err != nil {
		logging.Errorf(c, "Failed to flush tsmon metrics: %s", err)
	}
}

func updateLastFlushed(c context.Context) bool {
	now := clock.Now(c)
	minuteAgo := now.Add(-time.Minute)

	lastFlushed.Lock()
	defer lastFlushed.Unlock()

	if lastFlushed.After(minuteAgo) {
		return false
	}
	lastFlushed.Time = now // Don't hammer the datastore if task_num is not yet assigned.
	return true
}

func updateInstanceEntityAndFlush(c context.Context) error {
	c = info.Get(c).MustNamespace(instanceNamespace)

	task, ok := tsmon.Store(c).DefaultTarget().(*target.Task)
	if !ok {
		// tsmon probably failed to initialize - just do nothing.
		return fmt.Errorf("default tsmon target is not a Task: %v", tsmon.Store(c).DefaultTarget())
	}

	logger := logging.Get(c)
	entity := getOrCreateInstanceEntity(c)
	now := clock.Now(c)

	if entity.TaskNum < 0 {
		if task.TaskNum >= 0 {
			// We used to have a task number but we don't any more (we were inactive
			// for too long), so clear our state.
			logging.Warningf(c, "Instance %s got purged from Datastore, but is still alive. "+
				"Clearing cumulative metrics", info.Get(c).InstanceID())
			tsmon.ResetCumulativeMetrics(c)
		}
		task.TaskNum = -1
		lastFlushed.Time = entity.LastUpdated

		// Start complaining if we haven't been given a task number after some time.
		shouldHaveTaskNumBy := entity.LastUpdated.Add(instanceExpectedToHaveTaskNum)
		if shouldHaveTaskNumBy.Before(now) {
			logger.Warningf("Instance %s is %s old with no task_num.",
				info.Get(c).InstanceID(), now.Sub(shouldHaveTaskNumBy).String())
		}
		return nil
	}

	task.TaskNum = int32(entity.TaskNum)
	tsmon.Store(c).SetDefaultTarget(task)

	// Update the instance entity and put it back in the datastore asynchronously.
	entity.LastUpdated = now
	putDone := make(chan struct{})
	go func() {
		defer close(putDone)
		if err := datastore.Get(c).Put(entity); err != nil {
			logger.Errorf("Failed to update instance entity: %s", err)
		}
	}()

	ret := tsmon.Flush(c)
	resetGlobalCallbackMetrics(c)

	<-putDone
	return ret
}
