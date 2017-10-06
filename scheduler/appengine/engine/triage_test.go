// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"

	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/scheduler/appengine/internal"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTriageOp(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		tb := triageTestBed{}

		Convey("noop triage", func() {
			before := &Job{
				JobID:             "job",
				ActiveInvocations: []int64{1, 2, 3},
			}
			after, err := tb.runTestTriage(c, before)
			So(err, ShouldBeNil)
			So(after, ShouldResemble, before)
		})

		Convey("pops finished invocations", func() {
			before := &Job{
				JobID:             "job",
				ActiveInvocations: []int64{1, 2, 3, 4, 5},
			}
			recentlyFinishedSet(c, before.JobID).Add(c, []int64{1, 2, 4})

			after, err := tb.runTestTriage(c, before)
			So(err, ShouldBeNil)
			So(after.ActiveInvocations, ShouldResemble, []int64{3, 5})
		})

		Convey("processes triggers", func() {
			triggers := []*internal.Trigger{
				{
					Id:      "t0",
					Created: google.NewTimestamp(epoch.Add(20 * time.Second)),
				},
				{
					Id:      "t1",
					Created: google.NewTimestamp(epoch.Add(20 * time.Second)),
				},
				{
					Id:      "t2",
					Created: google.NewTimestamp(epoch.Add(10 * time.Second)),
				},
				{
					Id:      "a_t3", // to make its ID be before t0 and t1
					Created: google.NewTimestamp(epoch.Add(20 * time.Second)),
				},
			}

			before := &Job{JobID: "job"}
			pendingTriggersSet(c, before.JobID).Add(c, triggers)

			// Cycle 1. Pops triggers, converts them into a bunch of invocations.
			// Triggers are in sorted order!
			after, err := tb.runTestTriage(c, before)
			So(err, ShouldBeNil)
			So(after.ActiveInvocations, ShouldResemble, []int64{1, 2, 3, 4})
			So(tb.requests, ShouldResemble, []InvocationRequest{
				{IncomingTriggers: []*internal.Trigger{triggers[2]}},
				{IncomingTriggers: []*internal.Trigger{triggers[3]}},
				{IncomingTriggers: []*internal.Trigger{triggers[0]}},
				{IncomingTriggers: []*internal.Trigger{triggers[1]}},
			})

			tb.requests = nil

			// Cycle 2. Nothing new.
			after, err = tb.runTestTriage(c, after)
			So(err, ShouldBeNil)
			So(after.ActiveInvocations, ShouldResemble, []int64{1, 2, 3, 4})
			So(tb.requests, ShouldBeNil)
		})
	})
}

type triageTestBed struct {
	// Inputs.
	triggeringPolicy func(c context.Context, job *Job, triggers []*internal.Trigger) ([]InvocationRequest, error)

	// Outputs.
	nextInvID int64
	requests  []InvocationRequest
}

func (t *triageTestBed) runTestTriage(c context.Context, before *Job) (after *Job, err error) {
	if err := datastore.Put(c, before); err != nil {
		return nil, err
	}

	policy := t.triggeringPolicy
	if policy == nil {
		policy = t.defaultTriggeringPolicy
	}

	op := triageOp{
		jobID:            before.JobID,
		triggeringPolicy: policy,
		enqueueInvocations: func(c context.Context, job *Job, req []InvocationRequest) error {
			t.requests = append(t.requests, req...)
			for range req {
				t.nextInvID++
				job.ActiveInvocations = append(job.ActiveInvocations, t.nextInvID)
			}
			return nil
		},
	}
	if err := op.prepare(c); err != nil {
		return nil, err
	}

	err = runTxn(c, func(c context.Context) error {
		job := Job{JobID: op.jobID}
		if err := datastore.Get(c, &job); err != nil {
			return err
		}
		if err := op.transaction(c, &job); err != nil {
			return err
		}
		return datastore.Put(c, &job)
	})
	if err != nil {
		return nil, err
	}

	op.finalize(c)

	after = &Job{JobID: before.JobID}
	if err := datastore.Get(c, after); err != nil {
		return nil, err
	}
	return after, nil
}

func (t *triageTestBed) defaultTriggeringPolicy(c context.Context, job *Job, triggers []*internal.Trigger) (out []InvocationRequest, err error) {
	for _, t := range triggers {
		out = append(out, InvocationRequest{
			IncomingTriggers: []*internal.Trigger{t},
		})
	}
	return
}