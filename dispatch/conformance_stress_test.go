//go:build darwin

package dispatch

// Stress and bookkeeping conformance tests.
//
// The package funnels every dispatched closure through shared purego
// trampolines and package-level maps (workMap, persistentWorkMap, ...).
// These tests hammer that plumbing from many goroutines at once and then
// verify the bookkeeping drained — a leak in workMap is a real bug even
// when every closure ran.

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// pendingWork returns the number of undelivered closures in workMap.
func pendingWork() int {
	workMap.mu.Lock()
	defer workMap.mu.Unlock()
	return len(workMap.items)
}

func TestWorkMapDrainsAfterMixedOps(t *testing.T) {
	before := pendingWork()
	serial := QueueCreate("com.appledocs.dispatch.conformance.drain-serial")
	conc := QueueCreateConcurrent("com.appledocs.dispatch.conformance.drain-conc")
	g := GroupCreate()

	const n = 5000
	var ran atomic.Int64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := range n {
		work := func() {
			ran.Add(1)
			wg.Done()
		}
		switch i % 5 {
		case 0:
			serial.Async(work)
		case 1:
			conc.Async(work)
		case 2:
			conc.BarrierAsync(work)
		case 3:
			g.Async(conc, work)
		case 4:
			serial.AsyncAndWait(work)
		}
	}
	waitTimeout(t, &wg, 30*time.Second, "mixed dispatch operations")
	g.Wait(TimeForever)
	if got := ran.Load(); got != n {
		t.Fatalf("ran %d closures, want %d", got, n)
	}
	// Delivery decrements happen inside the trampoline before fn runs, so
	// by the time every fn has finished the map must be back to baseline.
	deadline := time.Now().Add(5 * time.Second)
	for pendingWork() > before && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if after := pendingWork(); after > before {
		t.Fatalf("workMap leaked %d entries (before=%d after=%d)", after-before, before, after)
	}
}

func TestManyGoroutinesManyQueues(t *testing.T) {
	// Concurrent submitters × concurrent queues × mixed sync/async ops.
	// This is the closest analogue to mlxpurego's reentrant_concurrent
	// suite: every operation crosses the Go->C boundary under contention.
	const submitters = 16
	const opsPer = 500
	queues := make([]Queue, 8)
	for i := range queues {
		if i%2 == 0 {
			queues[i] = QueueCreate("com.appledocs.dispatch.conformance.mq-serial")
		} else {
			queues[i] = QueueCreateConcurrent("com.appledocs.dispatch.conformance.mq-conc")
		}
	}
	var sum atomic.Int64
	var done atomic.Int64
	var wg sync.WaitGroup
	wg.Add(submitters)
	for s := range submitters {
		go func() {
			defer wg.Done()
			rng := rand.New(rand.NewSource(int64(s)))
			var local sync.WaitGroup
			for op := range opsPer {
				q := queues[rng.Intn(len(queues))]
				v := int64(s*opsPer + op)
				switch rng.Intn(3) {
				case 0:
					local.Add(1)
					q.Async(func() {
						sum.Add(v)
						done.Add(1)
						local.Done()
					})
				case 1:
					q.Sync(func() {
						sum.Add(v)
						done.Add(1)
					})
				default:
					q.AsyncAndWait(func() {
						sum.Add(v)
						done.Add(1)
					})
				}
			}
			local.Wait()
		}()
	}
	waitTimeout(t, &wg, 60*time.Second, "concurrent submitters")
	const total = submitters * opsPer
	if got := done.Load(); got != total {
		t.Fatalf("completed %d ops, want %d", got, total)
	}
	want := int64(total) * (total - 1) / 2
	if got := sum.Load(); got != want {
		t.Fatalf("sum = %d, want %d (an op ran twice, never, or with a torn closure)", got, want)
	}
}

func TestSemaphorePingPongLatency(t *testing.T) {
	// Two GCD threads ping-ponging through a pair of semaphores; a missed
	// wakeup deadlocks, which the timeout converts into a failure.
	const rounds = 10000
	ping := SemaphoreCreate(0)
	pong := SemaphoreCreate(0)
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.pingpong")
	g := GroupCreate()
	var count atomic.Int64
	g.Async(q, func() {
		for range rounds {
			ping.Signal()
			pong.Wait(TimeForever)
		}
	})
	g.Async(q, func() {
		for range rounds {
			ping.Wait(TimeForever)
			count.Add(1)
			pong.Signal()
		}
	})
	if !g.Wait(TimeFromNow(int64(60 * time.Second))) {
		t.Fatalf("ping-pong deadlocked after %d rounds", count.Load())
	}
	if count.Load() != rounds {
		t.Fatalf("completed %d rounds, want %d", count.Load(), rounds)
	}
}

func TestDataSourceCoalescingUnderLoad(t *testing.T) {
	// DATA_ADD sources coalesce merged values; the sum delivered across all
	// fires must equal the sum merged, regardless of how libdispatch batches.
	q := QueueCreate("com.appledocs.dispatch.conformance.coalesce")
	var delivered atomic.Int64
	src := NewDataAddSource(q, func(value uintptr) {
		delivered.Add(int64(value))
	})
	defer src.Cancel()

	const merges = 10000
	var want int64
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := range workers {
		go func() {
			defer wg.Done()
			for i := range merges / workers {
				src.MergeData(uintptr(w + i + 1))
			}
		}()
		for i := range merges / workers {
			want += int64(w + i + 1)
		}
	}
	waitTimeout(t, &wg, 30*time.Second, "data-add mergers")
	deadline := time.Now().Add(10 * time.Second)
	for delivered.Load() != want && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if got := delivered.Load(); got != want {
		t.Fatalf("delivered sum %d, want %d (coalesced events lost)", got, want)
	}
}

func TestManyTimersConcurrently(t *testing.T) {
	// 50 repeating timers on 50 queues, all flowing through the single
	// persistent trampoline; each must fire independently.
	const timers = 50
	fires := make([]atomic.Int64, timers)
	sources := make([]Source, timers)
	for i := range timers {
		q := QueueCreate("com.appledocs.dispatch.conformance.many-timers")
		sources[i] = NewTimerSource(5*time.Millisecond, time.Millisecond, q, func() {
			fires[i].Add(1)
		})
	}
	defer func() {
		for _, s := range sources {
			s.Cancel()
		}
	}()
	deadline := time.Now().Add(10 * time.Second)
	for {
		all := true
		for i := range fires {
			if fires[i].Load() < 3 {
				all = false
				break
			}
		}
		if all {
			return
		}
		if time.Now().After(deadline) {
			stuck := 0
			for i := range fires {
				if fires[i].Load() < 3 {
					stuck++
				}
			}
			t.Fatalf("%d/%d timers failed to fire 3 times", stuck, timers)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestSourceCancelHandlerRunsExactlyOnce(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.conformance.cancel-once")
	var cancels atomic.Int64
	src := SourceCreate(SourceTypeDataAdd, 0, 0, q)
	src.SetEventHandler(func() {})
	src.SetCancelHandler(func() {
		cancels.Add(1)
	})
	src.Activate()
	src.Cancel()
	src.Cancel() // second cancel must be a no-op
	deadline := time.Now().Add(5 * time.Second)
	for cancels.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond) // window for a (buggy) second delivery
	if got := cancels.Load(); got != 1 {
		t.Fatalf("cancel handler ran %d times, want 1", got)
	}
}

func TestWorkItemBarrierFlagOnConcurrentQueue(t *testing.T) {
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.item-barrier")
	var active, violations atomic.Int64
	var wg sync.WaitGroup
	const readers = 100
	wg.Add(readers + 1)
	for range readers / 2 {
		q.Async(func() {
			defer wg.Done()
			active.Add(1)
			time.Sleep(100 * time.Microsecond)
			active.Add(-1)
		})
	}
	item := WorkItemCreate(WorkItemBarrier, func() {
		defer wg.Done()
		if active.Load() != 0 {
			violations.Add(1)
		}
	})
	defer item.Release()
	q.AsyncWorkItem(item)
	for range readers / 2 {
		q.Async(func() {
			defer wg.Done()
			active.Add(1)
			time.Sleep(100 * time.Microsecond)
			active.Add(-1)
		})
	}
	waitTimeout(t, &wg, 15*time.Second, "barrier work item mix")
	if violations.Load() != 0 {
		t.Fatal("WorkItemBarrier ran concurrently with readers")
	}
}

func BenchmarkQueueSync(b *testing.B) {
	q := QueueCreate("com.appledocs.dispatch.bench.sync")
	f := func() {}
	b.ReportAllocs()
	for b.Loop() {
		q.Sync(f)
	}
}

func BenchmarkQueueAsyncThroughput(b *testing.B) {
	q := QueueCreate("com.appledocs.dispatch.bench.async")
	var wg sync.WaitGroup
	b.ReportAllocs()
	for b.Loop() {
		wg.Add(1)
		q.Async(func() { wg.Done() })
	}
	wg.Wait()
}

func BenchmarkSemaphoreSignalWait(b *testing.B) {
	sem := SemaphoreCreate(0)
	b.ReportAllocs()
	for b.Loop() {
		sem.Signal()
		sem.Wait(TimeForever)
	}
}
