//go:build darwin

package dispatch

// Reentrancy conformance tests.
//
// Dispatch callbacks run on GCD-managed threads and re-enter Go through
// purego callbacks. These tests torture the Go->C->Go boundary: nested
// synchronous dispatch across queues, dispatch from inside callbacks,
// goroutine stack growth mid-callback, and callbacks that themselves spin
// up Go concurrency.

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNestedSyncAcrossQueues(t *testing.T) {
	// Sync onto q[i+1] from inside q[i]'s callback. Each level parks one
	// GCD worker thread; depth stays well below libdispatch's thread cap.
	const depth = 16
	queues := make([]Queue, depth)
	for i := range queues {
		queues[i] = QueueCreate("com.appledocs.dispatch.conformance.nest")
	}
	var reached int
	var recurse func(level int)
	recurse = func(level int) {
		if level == depth {
			return
		}
		queues[level].Sync(func() {
			queues[level].AssertCurrent()
			reached = level + 1
			recurse(level + 1)
		})
	}
	done := make(chan struct{})
	go func() {
		recurse(0)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(15 * time.Second):
		t.Fatalf("deadlock in nested Sync at depth %d/%d", reached, depth)
	}
	if reached != depth {
		t.Fatalf("reached depth %d, want %d", reached, depth)
	}
}

func TestAsyncChainReentry(t *testing.T) {
	// Each callback asynchronously dispatches the next link from inside the
	// previous callback — a long Go->C->Go->C chain through the shared
	// trampoline and workMap.
	const links = 2000
	q := QueueCreate("com.appledocs.dispatch.conformance.chain")
	done := make(chan int, 1)
	var step func(n int)
	step = func(n int) {
		if n == links {
			done <- n
			return
		}
		q.Async(func() { step(n + 1) })
	}
	step(0)
	select {
	case n := <-done:
		if n != links {
			t.Fatalf("chain terminated at %d, want %d", n, links)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("async chain never completed")
	}
}

func TestStackGrowthInsideCallback(t *testing.T) {
	// Force goroutine stack growth while executing inside a purego callback
	// on a GCD thread. If the bridge mishandled stack movement mid-callback,
	// this corrupts or crashes rather than computing the right answer.
	var deep func(n int) int
	deep = func(n int) int {
		var pad [256]byte
		pad[0] = byte(n)
		if n == 0 {
			return int(pad[0])
		}
		return deep(n-1) + int(pad[0])&1
	}
	q := QueueCreate("com.appledocs.dispatch.conformance.stackgrow")
	var got int
	q.Sync(func() {
		got = deep(20000)
	})
	if got == 0 && deep(20000) != 0 {
		t.Fatal("deep recursion inside callback returned wrong result")
	}
}

func TestGCDuringCallbacks(t *testing.T) {
	// Hammer the collector while callbacks are in flight. Closures capture
	// heap values; if the work-item plumbing dropped a reference, results
	// would be lost or wrong under GC pressure.
	const n = 2000
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.gc")
	var sum atomic.Int64
	var wg sync.WaitGroup
	wg.Add(n)

	stopGC := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopGC:
				return
			default:
				runtime.GC()
			}
		}
	}()

	for i := range n {
		v := make([]int64, 8) // heap value captured by the closure
		v[3] = int64(i)
		q.Async(func() {
			sum.Add(v[3])
			wg.Done()
		})
	}
	waitTimeout(t, &wg, 30*time.Second, "callbacks under GC pressure")
	close(stopGC)
	want := int64(n) * (n - 1) / 2
	if got := sum.Load(); got != want {
		t.Fatalf("sum = %d, want %d (lost or corrupted work under GC)", got, want)
	}
}

func TestGoroutinesSpawnedFromCallback(t *testing.T) {
	// A dispatch callback that spawns Go concurrency and blocks on it
	// exercises the runtime's scheduling of a C-callback goroutine.
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.spawn")
	var total atomic.Int64
	var outer sync.WaitGroup
	const callbacks, spawned = 50, 20
	outer.Add(callbacks)
	for range callbacks {
		q.Async(func() {
			defer outer.Done()
			var inner sync.WaitGroup
			inner.Add(spawned)
			for range spawned {
				go func() {
					total.Add(1)
					inner.Done()
				}()
			}
			inner.Wait()
		})
	}
	waitTimeout(t, &outer, 30*time.Second, "goroutine-spawning callbacks")
	if got := total.Load(); got != callbacks*spawned {
		t.Fatalf("spawned goroutines ran %d times, want %d", got, callbacks*spawned)
	}
}

func TestSyncFromApplyCallback(t *testing.T) {
	// Re-enter dispatch synchronously from inside a dispatch_apply block.
	const n = 64
	work := QueueCreateConcurrent("com.appledocs.dispatch.conformance.apply-outer")
	side := QueueCreate("com.appledocs.dispatch.conformance.apply-side")
	var sum atomic.Int64
	Apply(n, work, func(i int) {
		side.Sync(func() {
			sum.Add(int64(i))
		})
	})
	if want := int64(n) * (n - 1) / 2; sum.Load() != want {
		t.Fatalf("sum = %d, want %d", sum.Load(), want)
	}
}

func TestGroupReentrantAsync(t *testing.T) {
	// Group work that enqueues more group work from inside its callback,
	// racing the outstanding-count against Wait.
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.group-reent")
	g := GroupCreate()
	var count atomic.Int64
	const fanout, depth = 4, 5 // 4^0+...+4^5 = 1365 tasks
	var spawn func(level int)
	spawn = func(level int) {
		count.Add(1)
		if level == depth {
			return
		}
		for range fanout {
			// Enter before Async returns so Wait can't fire early.
			g.Enter()
			q.Async(func() {
				defer g.Leave()
				spawn(level + 1)
			})
		}
	}
	g.Enter()
	q.Async(func() {
		defer g.Leave()
		spawn(0)
	})
	if !g.Wait(TimeFromNow(int64(60 * time.Second))) {
		t.Fatal("timeout waiting for reentrant group work")
	}
	var want int64
	for l, per := 0, int64(1); l <= depth; l, per = l+1, per*fanout {
		want += per
	}
	if got := count.Load(); got != want {
		t.Fatalf("ran %d tasks, want %d", got, want)
	}
}

func TestTimerFiresRepeatedlyWhileDispatching(t *testing.T) {
	// A repeating timer source firing while the same process floods the
	// shared trampolines from other queues.
	q := QueueCreate("com.appledocs.dispatch.conformance.timer-mix")
	noise := QueueCreateConcurrent("com.appledocs.dispatch.conformance.timer-noise")
	var fires atomic.Int64
	src := NewTimerSource(2*time.Millisecond, time.Millisecond, q, func() {
		fires.Add(1)
	})
	defer src.Cancel()

	var wg sync.WaitGroup
	const n = 1000
	wg.Add(n)
	for range n {
		noise.Async(func() {
			wg.Done()
		})
	}
	waitTimeout(t, &wg, 15*time.Second, "noise dispatches")
	deadline := time.After(5 * time.Second)
	for fires.Load() < 5 {
		select {
		case <-deadline:
			t.Fatalf("timer fired only %d times under dispatch load", fires.Load())
		case <-time.After(5 * time.Millisecond):
		}
	}
}
