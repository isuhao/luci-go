// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package tumble

import (
	"bytes"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/luci/gae/filter/txnBuf"
	ds "github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/datastore/serialize"
	"github.com/luci/gae/service/info"
	mc "github.com/luci/gae/service/memcache"
	"github.com/luci/luci-go/appengine/memlock"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/data/stringset"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/sync/parallel"

	"golang.org/x/net/context"
)

// expandedShardBounds returns the boundary of the expandedShard order that
// currently corresponds to this shard number. If Shard is < 0 or > NumShards
// (the currently configured number of shards), this will return a low > high.
// Otherwise low < high.
func expandedShardBounds(c context.Context, cfg *Config, shard uint64) (low, high int64) {
	if shard < 0 || uint64(shard) >= cfg.NumShards {
		logging.Warningf(c, "Invalid shard: %d", shard)
		// return inverted bounds
		return 0, -1
	}

	expandedShardsPerShard := int64(math.MaxUint64 / cfg.NumShards)
	low = math.MinInt64 + (int64(shard) * expandedShardsPerShard)
	if uint64(shard) == cfg.NumShards-1 {
		high = math.MaxInt64
	} else {
		high = low + expandedShardsPerShard
	}
	return
}

func processShardQuery(c context.Context, cfg *Config, shard uint64) *ds.Query {
	low, high := expandedShardBounds(c, cfg, shard)
	if low > high {
		return nil
	}

	return ds.NewQuery("tumble.Mutation").
		Gte("ExpandedShard", low).Lte("ExpandedShard", high).
		Project("TargetRoot").Distinct(true).
		Limit(cfg.ProcessMaxBatchSize)
}

// processShard is the tumble backend endpoint. This accepts a shard number which
// is expected to be < GlobalConfig.NumShards.
func processShard(c context.Context, cfg *Config, namespaces []string, timestamp time.Time, shard uint64) error {
	logging.Fields{
		"shard": shard,
	}.Infof(c, "Processing tumble shard.")

	q := processShardQuery(c, cfg, shard)

	if q == nil {
		logging.Warningf(c, "dead shard, quitting")
		return nil
	}

	// If there are no namesapces, there is nothing to process.
	if len(namespaces) == 0 {
		logging.Infof(c, "no namespaces, quitting")
		return nil
	}

	tasks := make([]*processTask, len(namespaces))
	for i, ns := range namespaces {
		tasks[i] = makeProcessTask(ns, timestamp, shard)
	}

	lockKey := fmt.Sprintf("%s.%d.lock", baseName, shard)
	clientID := fmt.Sprintf("%d_%d", timestamp.Unix(), shard)

	var err error
	for try := 0; try < 2; try++ {
		err = memlock.TryWithLock(c, lockKey, clientID, func(c context.Context) error {
			return parallel.FanOutIn(func(taskC chan<- func() error) {
				for _, task := range tasks {
					task := task

					taskC <- func() error {
						return task.process(c, cfg, q)
					}
				}
			})
		})
		if err != memlock.ErrFailedToLock {
			break
		}
		logging.Infof(c, "Couldn't obtain lock (try %d) (sleeping 2s)", try+1)
		if tr := clock.Sleep(c, time.Second*2); tr.Incomplete() {
			logging.Warningf(c, "sleep interrupted, context is done: %v", tr.Err)
			return tr.Err
		}
	}
	if err == memlock.ErrFailedToLock {
		logging.Infof(c, "Couldn't obtain lock (giving up): %s", err)
		err = nil
	}
	return err
}

// processTask is a stateful processing task. It is bound to a specific
// namespace.
type processTask struct {
	namespace string
	timestamp time.Time
	clientID  string
	lastKey   string
	banSets   map[string]stringset.Set
}

func makeProcessTask(namespace string, timestamp time.Time, shard uint64) *processTask {
	return &processTask{
		namespace: namespace,
		timestamp: timestamp,
		clientID:  fmt.Sprintf("%d_%d", timestamp.Unix(), shard),
		lastKey:   fmt.Sprintf("%s.%d.last", baseName, shard),
		banSets:   make(map[string]stringset.Set),
	}
}

func (t *processTask) process(c context.Context, cfg *Config, q *ds.Query) error {
	if t.namespace != "" {
		c = logging.SetField(c, "namespace", t.namespace)
		c = info.MustNamespace(c, t.namespace)
	}

	// this last key allows buffered tasks to early exit if some other shard
	// processor has already processed past this task's target timestamp.
	lastItm, err := mc.GetKey(c, t.lastKey)
	if err != nil {
		if err != mc.ErrCacheMiss {
			logging.Warningf(c, "couldn't obtain last timestamp: %s", err)
		}
	} else {
		val := lastItm.Value()
		last, err := serialize.ReadTime(bytes.NewBuffer(val))
		if err != nil {
			logging.Warningf(c, "could not decode timestamp %v: %s", val, err)
		} else {
			last = last.Add(time.Duration(cfg.TemporalRoundFactor))
			if last.After(t.timestamp) {
				logging.Infof(c, "early exit, %s > %s", last, t.timestamp)
				return nil
			}
		}
	}
	err = nil

	for {
		processCounters := []*int64{}
		err := parallel.WorkPool(int(cfg.NumGoroutines), func(ch chan<- func() error) {
			err := ds.Run(c, q, func(pm ds.PropertyMap) error {
				root := pm.Slice("TargetRoot")[0].Value().(*ds.Key)
				encRoot := root.Encode()

				// TODO(riannucci): make banSets remove keys from the banSet which
				// weren't hit. Once they stop showing up, they'll never show up
				// again.

				bs := t.banSets[encRoot]
				if bs == nil {
					bs = stringset.New(0)
					t.banSets[encRoot] = bs
				}
				counter := new(int64)
				processCounters = append(processCounters, counter)

				ch <- func() error {
					switch err := processRoot(c, cfg, root, bs, counter); err {
					case nil:
						return nil

					case ds.ErrConcurrentTransaction:
						logging.Fields{
							logging.ErrorKey: err,
							"root":           root,
						}.Warningf(c, "Transient error encountered processing root.")
						return errors.WrapTransient(err)

					default:
						logging.Fields{
							logging.ErrorKey: err,
							"root":           root,
						}.Errorf(c, "Failed to process root.")
						return err
					}
				}

				if c.Err() != nil {
					logging.Warningf(c, "Lost lock! %s", c.Err())
					return ds.Stop
				}
				return nil
			})
			if err != nil {
				var qstr string
				if fq, err := q.Finalize(); err == nil {
					qstr = fq.String()
				} else {
					qstr = fmt.Sprintf("unable to finalize: %v", err)
				}

				logging.Fields{
					logging.ErrorKey: err,
					"query":          qstr,
				}.Errorf(c, "Failure to query.")
				ch <- func() error {
					return err
				}
			}
		})
		if err != nil {
			return err
		}
		numProcessed := int64(0)
		for _, n := range processCounters {
			numProcessed += *n
		}
		logging.Infof(c, "cumulatively processed %d items", numProcessed)
		if numProcessed == 0 {
			break
		}

		err = mc.Set(c, mc.NewItem(c, t.lastKey).SetValue(serialize.ToBytes(clock.Now(c).UTC())))
		if err != nil {
			logging.Warningf(c, "could not update last process memcache key %s: %s", t.lastKey, err)
		}

		if tr := clock.Sleep(c, time.Duration(cfg.DustSettleTimeout)); tr.Incomplete() {
			logging.Warningf(c, "sleep interrupted, context is done: %v", tr.Err)
			return tr.Err
		}
	}
	return nil
}

func getBatchByRoot(c context.Context, cfg *Config, root *ds.Key, banSet stringset.Set) ([]*realMutation, error) {
	q := ds.NewQuery("tumble.Mutation").Eq("TargetRoot", root)
	if cfg.DelayedMutations {
		q = q.Lte("ProcessAfter", clock.Now(c).UTC())
	}

	toFetch := make([]*realMutation, 0, cfg.ProcessMaxBatchSize)
	err := ds.Run(c, q, func(k *ds.Key) error {
		if !banSet.Has(k.Encode()) {
			toFetch = append(toFetch, &realMutation{
				ID:     k.StringID(),
				Parent: k.Parent(),
			})
		}
		if len(toFetch) < cap(toFetch) {
			return nil
		}
		return ds.Stop
	})
	return toFetch, err
}

func loadFilteredMutations(c context.Context, rms []*realMutation) ([]*ds.Key, []Mutation, error) {
	mutKeys := make([]*ds.Key, 0, len(rms))
	muts := make([]Mutation, 0, len(rms))
	err := ds.Get(c, rms)
	me, ok := err.(errors.MultiError)
	if !ok && err != nil {
		return nil, nil, err
	}

	for i, rm := range rms {
		err = nil
		if me != nil {
			err = me[i]
		}
		if err == nil {
			if rm.Version != getAppVersion(c) {
				logging.Fields{
					"mut_version": rm.Version,
					"cur_version": getAppVersion(c),
				}.Warningf(c, "loading mutation with different code version")
			}
			m, err := rm.GetMutation()
			if err != nil {
				logging.Errorf(c, "couldn't load mutation: %s", err)
				continue
			}
			muts = append(muts, m)
			mutKeys = append(mutKeys, ds.KeyForObj(c, rm))
		} else if err != ds.ErrNoSuchEntity {
			return nil, nil, me
		}
	}

	return mutKeys, muts, nil
}

type overrideRoot struct {
	Mutation

	root *ds.Key
}

func (o overrideRoot) Root(context.Context) *ds.Key {
	return o.root
}

func processRoot(c context.Context, cfg *Config, root *ds.Key, banSet stringset.Set, counter *int64) error {
	l := logging.Get(c)

	toFetch, err := getBatchByRoot(c, cfg, root, banSet)
	if err != nil || len(toFetch) == 0 {
		return err
	}

	mutKeys, muts, err := loadFilteredMutations(c, toFetch)
	if err != nil {
		return err
	}

	if c.Err() != nil {
		l.Warningf("Lost lock during processRoot")
		return nil
	}

	allShards := map[taskShard]struct{}{}

	toDel := make([]*ds.Key, 0, len(muts))
	numMuts := uint64(0)
	deletedMuts := uint64(0)
	processedMuts := uint64(0)
	err = ds.RunInTransaction(txnBuf.FilterRDS(c), func(c context.Context) error {
		toDel = toDel[:0]
		numMuts = 0
		deletedMuts = 0
		processedMuts = 0

		iterMuts := muts
		iterMutKeys := mutKeys

		for i := 0; i < len(iterMuts); i++ {
			m := iterMuts[i]

			logging.Fields{"m": m}.Infof(c, "running RollForward")
			shards, newMuts, newMutKeys, err := enterTransactionMutation(c, cfg, overrideRoot{m, root}, uint64(i))
			if err != nil {
				l.Errorf("Executing decoded gob(%T) failed: %q: %+v", m, err, m)
				continue
			}
			processedMuts++
			for j, nm := range newMuts {
				if nm.Root(c).HasAncestor(root) {
					runNow := !cfg.DelayedMutations
					if !runNow {
						dm, isDelayedMutation := nm.(DelayedMutation)
						runNow = !isDelayedMutation || clock.Now(c).UTC().After(dm.ProcessAfter())
					}
					if runNow {
						iterMuts = append(iterMuts, nm)
						iterMutKeys = append(iterMutKeys, newMutKeys[j])
					}
				}
			}

			key := iterMutKeys[i]
			if key.HasAncestor(root) {
				// try to delete it as part of the same transaction.
				if err := ds.Delete(c, key); err == nil {
					deletedMuts++
				} else {
					toDel = append(toDel, key)
				}
			} else {
				toDel = append(toDel, key)
			}

			numMuts += uint64(len(newMuts))
			for shard := range shards {
				allShards[shard] = struct{}{}
			}
		}

		return nil
	}, nil)
	if err != nil {
		l.Errorf("failed running transaction: %s", err)
		return err
	}
	numMuts -= deletedMuts

	fireTasks(c, cfg, allShards)
	l.Infof("successfully processed %d mutations (%d tail-call), adding %d more", processedMuts, deletedMuts, numMuts)

	if len(toDel) > 0 {
		atomic.StoreInt64(counter, int64(len(toDel)))

		for _, k := range toDel {
			banSet.Add(k.Encode())
		}
		if err := ds.Delete(c, toDel); err != nil {
			l.Warningf("error deleting finished mutations: %s", err)
		}
	}

	return nil
}
