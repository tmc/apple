//go:build darwin

package dispatch

// Preemption-pressure conformance tests.
//
// Go's async preemption delivers SIGURG to running threads; libdispatch's
// wait primitives park callers in kernel waits (ulock/turnstile) that must
// tolerate that interruption. dispatch_async_and_wait provably did not
// (see TestAsyncAndWaitBehindPendingAsync). These tests put every other
// C-side blocking entry point the package exposes under the same
// conditions that reproduced that deadlock: a wait entered behind pending
// work, repeated tens of thousands of times, with busy goroutines and a GC
// ticker keeping the preemption signal firing.
//
// A hang here means the primitive's wait path loses wakeups under signals
// and must be reimplemented with a Go-side wait like AsyncAndWait was.

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

const preemptionIters = 50000

// withPreemptionPressure runs fn while busy-loop goroutines and a GC ticker
// keep the runtime's SIGURG preemption signals firing.
func withPreemptionPressure(t *testing.T, fn func()) {
	t.Helper()
	if testing.Short() {
		t.Skip("50k-iteration preemption stress; skipped in -short mode")
	}
	stop := make(chan struct{})
	defer close(stop)
	// Busy loops draw SIGURG preemption at every GC and scheduler event.
	// Use half the Ps: enough to keep signals firing at the dispatch
	// threads, while leaving scheduler headroom so Go-side waits (channel
	// receives, callback goroutines) make progress instead of starving.
	busy := max(1, runtime.GOMAXPROCS(0)/2)
	var sink atomic.Int64
	for range busy {
		go func() {
			x := 0
			for {
				select {
				case <-stop:
					sink.Store(int64(x))
					return
				default:
					for i := range 1 << 16 {
						x += i
					}
				}
			}
		}()
	}
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				runtime.GC()
			}
		}
	}()
	fn()
}

func TestPreemptionSyncBehindPending(t *testing.T) {
	withPreemptionPressure(t, func() {
		q := QueueCreate("com.appledocs.dispatch.preempt.sync")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			q.Sync(func() { n.Add(1) })
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionBarrierSyncBehindPending(t *testing.T) {
	withPreemptionPressure(t, func() {
		q := QueueCreateConcurrent("com.appledocs.dispatch.preempt.barrier")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			q.BarrierSync(func() { n.Add(1) })
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionSyncWorkItemBehindPending(t *testing.T) {
	// Block-based dispatch_sync (not the _f variant the closure API uses).
	withPreemptionPressure(t, func() {
		q := QueueCreate("com.appledocs.dispatch.preempt.sync-item")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			item := WorkItemCreate(0, func() { n.Add(1) })
			q.SyncWorkItem(item)
			item.Release()
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionAsyncAndWaitWorkItemBehindPending(t *testing.T) {
	// Formerly block-based dispatch_async_and_wait — the same C primitive
	// whose _f variant deadlocked. Hung 3/3 here before being reimplemented
	// as dispatch_async + dispatch_block_wait; this test pins that fix.
	withPreemptionPressure(t, func() {
		q := QueueCreate("com.appledocs.dispatch.preempt.aaw-item")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			item := WorkItemCreate(0, func() { n.Add(1) })
			q.AsyncAndWaitWorkItem(item)
			item.Release()
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionWorkItemWaitBehindPending(t *testing.T) {
	// dispatch_block_wait on an item enqueued behind a pending async.
	withPreemptionPressure(t, func() {
		q := QueueCreate("com.appledocs.dispatch.preempt.item-wait")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			item := WorkItemCreate(0, func() { n.Add(1) })
			q.AsyncWorkItem(item)
			if !item.Wait(TimeForever) {
				t.Fatal("WorkItem.Wait(TimeForever) returned false")
			}
			item.Release()
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionGroupWait(t *testing.T) {
	withPreemptionPressure(t, func() {
		q := QueueCreateConcurrent("com.appledocs.dispatch.preempt.group")
		g := GroupCreate()
		var n atomic.Int64
		for range preemptionIters {
			g.Async(q, func() { n.Add(1) })
			g.Async(q, func() { n.Add(1) })
			if !g.Wait(TimeForever) {
				t.Fatal("Group.Wait(TimeForever) returned false")
			}
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

func TestPreemptionSemaphorePingPong(t *testing.T) {
	withPreemptionPressure(t, func() {
		ping := SemaphoreCreate(0)
		pong := SemaphoreCreate(0)
		q := QueueCreateConcurrent("com.appledocs.dispatch.preempt.sema")
		g := GroupCreate()
		var n atomic.Int64
		g.Async(q, func() {
			for range preemptionIters {
				ping.Signal()
				pong.Wait(TimeForever)
			}
		})
		g.Async(q, func() {
			for range preemptionIters {
				ping.Wait(TimeForever)
				n.Add(1)
				pong.Signal()
			}
		})
		if !g.Wait(TimeForever) {
			t.Fatal("semaphore ping-pong never completed")
		}
		if got := n.Load(); got != preemptionIters {
			t.Fatalf("completed %d rounds, want %d", got, preemptionIters)
		}
	})
}

func TestPreemptionApply(t *testing.T) {
	// dispatch_apply blocks the caller until all iterations finish.
	withPreemptionPressure(t, func() {
		q := QueueCreateConcurrent("com.appledocs.dispatch.preempt.apply")
		var n atomic.Int64
		const rounds = preemptionIters / 10
		for range rounds {
			Apply(10, q, func(int) { n.Add(1) })
		}
		if got := n.Load(); got != 10*rounds {
			t.Fatalf("ran %d iterations, want %d", got, 10*rounds)
		}
	})
}

func TestPreemptionAsyncAndWaitGoSide(t *testing.T) {
	// The fixed Go-side AsyncAndWait under the same pressure, for symmetry
	// with TestAsyncAndWaitBehindPendingAsync.
	withPreemptionPressure(t, func() {
		q := QueueCreate("com.appledocs.dispatch.preempt.aaw")
		var n atomic.Int64
		for range preemptionIters {
			q.Async(func() { n.Add(1) })
			q.AsyncAndWait(func() { n.Add(1) })
		}
		if got := n.Load(); got != 2*preemptionIters {
			t.Fatalf("ran %d closures, want %d", got, 2*preemptionIters)
		}
	})
}

// Keep the compiler honest about the timeout import if iterations change.
var _ = time.Second
