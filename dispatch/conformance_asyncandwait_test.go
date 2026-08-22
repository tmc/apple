//go:build darwin

package dispatch

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Regression test for the AsyncAndWait lost-wakeup deadlock.
//
// dispatch_async_and_wait_f parks the caller in a libdispatch
// ownership-handoff wait that loses its wakeup when Go's async-preemption
// signal (SIGURG) lands in the handoff window; with a pending Async item on
// the same serial queue this deadlocked within ~50k iterations
// (GODEBUG=asyncpreemptoff=1 was clean, plain C was clean, Sync/BarrierSync
// were clean). AsyncAndWait therefore waits on a Go channel instead of in
// the C primitive; this test pins that fix at the iteration count that
// reliably reproduced the hang.
func TestAsyncAndWaitBehindPendingAsync(t *testing.T) {
	if testing.Short() {
		t.Skip("50k-iteration regression guard; skipped in -short mode")
	}
	q := QueueCreate("com.appledocs.dispatch.conformance.aaw-pending")
	var completed atomic.Int64
	var wg sync.WaitGroup
	for range 50000 {
		wg.Add(1)
		q.Async(func() { wg.Done() })
		q.AsyncAndWait(func() { completed.Add(1) })
	}
	wg.Wait()
	if got := completed.Load(); got != 50000 {
		t.Fatalf("completed %d AsyncAndWait calls, want 50000", got)
	}
}
