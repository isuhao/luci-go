// Copyright 2016 The LUCI Authors.
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

package gitiles

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/pubsub/v1"

	"go.chromium.org/luci/common/api/gitiles"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gitilespb "go.chromium.org/luci/common/proto/gitiles"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/config/validation"
	"go.chromium.org/luci/server/auth"

	api "go.chromium.org/luci/scheduler/api/scheduler/v1"
	"go.chromium.org/luci/scheduler/appengine/internal"
	"go.chromium.org/luci/scheduler/appengine/messages"
	"go.chromium.org/luci/scheduler/appengine/task"
)

// defaultMaxTriggersPerInvocation limits number of triggers emitted per one
// invocation.
const defaultMaxTriggersPerInvocation = 100

// TaskManager implements task.Manager interface for tasks defined with
// GitilesTask proto message.
type TaskManager struct {
	mockGitilesClient        gitilespb.GitilesClient // Used for testing only.
	maxTriggersPerInvocation int                     // Avoid choking on DS or runtime limits.
}

// Name is part of Manager interface.
func (m TaskManager) Name() string {
	return "gitiles"
}

// ProtoMessageType is part of Manager interface.
func (m TaskManager) ProtoMessageType() proto.Message {
	return (*messages.GitilesTask)(nil)
}

// Traits is part of Manager interface.
func (m TaskManager) Traits() task.Traits {
	return task.Traits{
		Multistage: false, // we don't use task.StatusRunning state
	}
}

// ValidateProtoMessage is part of Manager interface.
func (m TaskManager) ValidateProtoMessage(c *validation.Context, msg proto.Message) {
	cfg, ok := msg.(*messages.GitilesTask)
	if !ok {
		c.Errorf("wrong type %T, expecting *messages.GitilesTask", msg)
		return
	}

	// Validate 'repo' field.
	c.Enter("repo")
	if cfg.Repo == "" {
		c.Errorf("field 'repository' is required")
	} else {
		u, err := url.Parse(cfg.Repo)
		if err != nil {
			c.Errorf("invalid URL %q: %s", cfg.Repo, err)
		} else if !u.IsAbs() {
			c.Errorf("not an absolute url: %q", cfg.Repo)
		}
	}
	c.Exit()

	c.Enter("refs")
	for _, ref := range cfg.Refs {
		if !strings.HasPrefix(ref, "refs/") {
			c.Errorf("ref must start with 'refs/' not %q", ref)
		}
		cnt := strings.Count(ref, "*")
		if cnt > 1 || (cnt == 1 && !strings.HasSuffix(ref, "/*")) {
			c.Errorf("only trailing (e.g. refs/blah/*) globs are supported, not %q", ref)
		}
	}
	c.Exit()
}

// LaunchTask is part of Manager interface.
func (m TaskManager) LaunchTask(c context.Context, ctl task.Controller) error {
	cfg := ctl.Task().(*messages.GitilesTask)
	ctl.DebugLog("Repo: %s, Refs: %s", cfg.Repo, cfg.Refs)
	repoURL, err := url.Parse(cfg.Repo)
	if err != nil {
		return err
	}

	refs, err := m.fetchRefsState(c, ctl, cfg, repoURL)
	if err != nil {
		ctl.DebugLog("Error fetching state of the world: %s", err)
		return err
	}

	refs.pruneKnown(ctl)
	leftToProcess := m.emitTriggersRefAtATime(c, ctl, cfg.Repo, refs)

	if refs.changed == 0 {
		ctl.DebugLog("No changes detected")
		ctl.State().Status = task.StatusSucceeded
		return nil
	}
	ctl.DebugLog("%d changed refs processed, %d refs not yet examined", refs.changed, leftToProcess)
	// Force save to ensure triggers are actually emitted.
	if err := ctl.Save(c); err != nil {
		// At this point, triggers have not been sent, so bail now and don't save
		// the refs' heads newest values.
		return err
	}
	if err := saveState(c, ctl.JobID(), repoURL, refs.known); err != nil {
		return err
	}
	ctl.DebugLog("Saved %d known refs", len(refs.known))
	ctl.State().Status = task.StatusSucceeded
	return nil
}

// AbortTask is part of Manager interface.
func (m TaskManager) AbortTask(c context.Context, ctl task.Controller) error {
	return nil
}

// HandleNotification is part of Manager interface.
func (m TaskManager) HandleNotification(c context.Context, ctl task.Controller, msg *pubsub.PubsubMessage) error {
	return errors.New("not implemented")
}

// HandleTimer is part of Manager interface.
func (m TaskManager) HandleTimer(c context.Context, ctl task.Controller, name string, payload []byte) error {
	return errors.New("not implemented")
}

func (m TaskManager) fetchRefsState(c context.Context, ctl task.Controller, cfg *messages.GitilesTask, repoURL *url.URL) (*refsState, error) {
	refs := &refsState{}
	refs.watched.init(cfg.GetRefs())
	return refs, parallel.FanOutIn(func(work chan<- func() error) {
		work <- func() (loadErr error) {
			refs.known, loadErr = loadState(c, ctl.JobID(), repoURL)
			return
		}
		work <- func() (refsErr error) {
			// Merge getRefsTips here.
			refs.current, refsErr = m.getRefsTips(c, ctl, cfg.Repo, refs.watched)
			return
		}
	})
}

// getRefsTips returns tip for each ref being watched.
func (m TaskManager) getRefsTips(c context.Context, ctl task.Controller, repoURL string, watched watchedRefs) (map[string]string, error) {
	host, project, err := gitiles.ParseRepoURL(repoURL)
	if err != nil {
		return nil, errors.Annotate(err, "invalid repo URL %q", repoURL).Err()
	}

	g, err := m.getGitilesClient(c, ctl, host)
	if err != nil {
		return nil, err
	}

	// Query gitiles for each namespace in parallel.
	var wg sync.WaitGroup
	var lock sync.Mutex
	errs := []error{}
	allTips := map[string]string{}
	// Group all refs by their namespace to reduce # of RPCs.
	for _, wrs := range watched.namespaces {
		wg.Add(1)
		go func(wrs *watchedRefNamespace) {
			defer wg.Done()
			res, err := g.Refs(c, &gitilespb.RefsRequest{
				Project:  project,
				RefsPath: wrs.namespace,
			})
			lock.Lock()
			defer lock.Unlock()
			if err != nil {
				ctl.DebugLog("failed to fetch %q namespace tips for %q: %q", wrs.namespace, err)
				errs = append(errs, err)
				return
			}
			for ref, tip := range res.Revisions {
				if watched.hasRef(ref) {
					allTips[ref] = tip
				}
			}
		}(wrs)
	}
	wg.Wait()
	if len(errs) > 0 {
		return nil, errors.NewMultiError(errs...)
	}
	return allTips, nil
}

// emitTriggersRefAtATime processes refs one a time and emits triggers if ref
// changed. Limits number of triggers emitted and so may stop early.
//
// Returns how many refs were not examined.
func (m TaskManager) emitTriggersRefAtATime(c context.Context, ctl task.Controller, repo string, refs *refsState) (left int) {
	maxTriggersPerInvocation := m.maxTriggersPerInvocation
	if maxTriggersPerInvocation == 0 {
		maxTriggersPerInvocation = defaultMaxTriggersPerInvocation
	}
	emittedTriggers := 0
	// Note, that refs.current contain only watched refs (see getRefsTips).
	// For determinism, sort refs by name.
	sortedRefs := make([]string, 0, len(refs.current))
	for ref := range refs.current {
		sortedRefs = append(sortedRefs, ref)
	}
	sort.Strings(sortedRefs)
	for i, ref := range sortedRefs {
		// TODO(tandrii): enforce limit on emitted triggers here.
		emittedTriggers += m.emitTriggersForRef(c, ctl, repo, ref, refs)
		// Safeguard against too many changes such as the first run after
		// config change to watch many more refs than before.
		if emittedTriggers >= maxTriggersPerInvocation {
			ctl.DebugLog("Emitted %d triggers, postponing the rest", emittedTriggers)
			left = len(sortedRefs) - i - 1
			return left
		}
	}
	left = 0
	return
}

// emitTriggersForRef emits triggers for 1 ref if there are any changes and
// returns number of triggers emitted.
func (m TaskManager) emitTriggersForRef(c context.Context, ctl task.Controller, repo, ref string, refs *refsState) int {
	newHead := refs.current[ref]
	switch oldHead, existed := refs.known[ref]; {
	case !existed:
		ctl.DebugLog("Ref %s is new: %s", ref, newHead)
	case oldHead != newHead:
		ctl.DebugLog("Ref %s updated: %s => %s", ref, oldHead, newHead)
	default:
		return 0 // no change
	}
	refs.known[ref] = newHead
	refs.changed++
	// TODO(tandrii): actually look at commits between current and previously
	// known tips of each ref.
	// In current (v1) engine, all triggers emitted around the same time will
	// result in just 1 invocation of each triggered job. Therefore,
	// passing just HEAD's revision is good enough.
	// For the same reason, only 1 of the refs will actually be processed if
	// several refs changed at the same time.
	ctl.EmitTrigger(c, &internal.Trigger{
		Id:    fmt.Sprintf("%s/+/%s@%s", repo, ref, newHead),
		Title: newHead,
		Url:   fmt.Sprintf("%s/+/%s", repo, newHead),
		Payload: &internal.Trigger_Gitiles{
			Gitiles: &api.GitilesTrigger{Repo: repo, Ref: ref, Revision: newHead},
		},
	})
	return 1
}

func (m TaskManager) getGitilesClient(c context.Context, ctl task.Controller, host string) (gitilespb.GitilesClient, error) {
	if m.mockGitilesClient != nil {
		// Used for testing only.
		logging.Infof(c, "using mockGitilesClient")
		return m.mockGitilesClient, nil
	}

	httpClient, err := ctl.GetClient(c, time.Minute, auth.WithScopes(gitiles.OAuthScope))
	if err != nil {
		return nil, err
	}

	return gitiles.NewRESTClient(httpClient, host, true)
}

type refsState struct {
	watched watchedRefs
	known   map[string]string // HEADs we saw before
	current map[string]string // HEADs available now
	changed int
}

func (s *refsState) pruneKnown(ctl task.Controller) {
	for ref := range s.known {
		switch {
		case !s.watched.hasRef(ref):
			ctl.DebugLog("Ref %s is no longer watched", ref)
			delete(s.known, ref)
			s.changed++
		case s.current[ref] == "":
			ctl.DebugLog("Ref %s deleted", ref)
			delete(s.known, ref)
			s.changed++
		}
	}
}

func (s *refsState) sortedCurrentRefNames() []string {
	sortedRefs := make([]string, 0, len(s.current))
	for ref := range s.current {
		sortedRefs = append(sortedRefs, ref)
	}
	sort.Strings(sortedRefs)
	return sortedRefs
}

type watchedRefNamespace struct {
	namespace    string // no trailing "/".
	allChildren  bool   // if true, someChildren is ignored.
	someChildren stringset.Set
}

func (w watchedRefNamespace) hasSuffix(suffix string) bool {
	switch {
	case suffix == "*":
		panic(fmt.Errorf("watchedRefNamespace membership should only be checked for refs, not ref glob %s", suffix))
	case w.allChildren:
		return true
	case w.someChildren == nil:
		return false
	default:
		return w.someChildren.Has(suffix)
	}
}

func (w *watchedRefNamespace) addSuffix(suffix string) {
	switch {
	case w.allChildren:
		return
	case suffix == "*":
		w.allChildren = true
		w.someChildren = nil
		return
	case w.someChildren == nil:
		w.someChildren = stringset.New(1)
	}
	w.someChildren.Add(suffix)
}

type watchedRefs struct {
	namespaces map[string]*watchedRefNamespace
}

func (w *watchedRefs) init(refsConfig []string) {
	w.namespaces = map[string]*watchedRefNamespace{}
	for _, ref := range refsConfig {
		ns, suffix := splitRef(ref)
		if _, exists := w.namespaces[ns]; !exists {
			w.namespaces[ns] = &watchedRefNamespace{namespace: ns}
		}
		w.namespaces[ns].addSuffix(suffix)
	}
}

func (w *watchedRefs) hasRef(ref string) bool {
	ns, suffix := splitRef(ref)
	if wrn, exists := w.namespaces[ns]; exists {
		return wrn.hasSuffix(suffix)
	}
	return false
}

func splitRef(s string) (string, string) {
	if i := strings.LastIndex(s, "/"); i <= 0 {
		return s, ""
	} else {
		return s[:i], s[i+1:]
	}
}
