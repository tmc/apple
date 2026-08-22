//go:build darwin

package dispatch

// Ordering and semantics conformance tests.
//
// These tests prove the semantic contracts of the queue wrappers rather than
// just "did not crash": serial FIFO order, barrier exclusivity on concurrent
// queues, suspend/resume gating, target-queue serialization, and After timing.

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// waitTimeout runs wait-for-wg with a deadline so a semantic failure surfaces
// as a test failure instead of a hung test binary.
func waitTimeout(t *testing.T, wg *sync.WaitGroup, d time.Duration, what string) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(d):
		t.Fatalf("timeout waiting for %s", what)
	}
}

func TestSerialQueueFIFOOrder(t *testing.T) {
	const n = 1000
	q := QueueCreate("com.appledocs.dispatch.conformance.fifo")
	order := make([]int, 0, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := range n {
		q.Async(func() {
			order = append(order, i) // serialized by the queue
			wg.Done()
		})
	}
	waitTimeout(t, &wg, 10*time.Second, "serial FIFO work")
	if len(order) != n {
		t.Fatalf("ran %d items, want %d", len(order), n)
	}
	for i, got := range order {
		if got != i {
			t.Fatalf("order[%d] = %d, want %d (serial queue is not FIFO)", i, got, i)
		}
	}
}

func TestConcurrentQueueBarrierExclusivity(t *testing.T) {
	const readers, barriers = 200, 20
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.barrier")

	var active atomic.Int64   // readers currently inside their body
	var inBarrier atomic.Bool // a barrier body is running
	var violations atomic.Int64
	var wg sync.WaitGroup

	reader := func() {
		defer wg.Done()
		if inBarrier.Load() {
			violations.Add(1)
		}
		active.Add(1)
		time.Sleep(100 * time.Microsecond)
		active.Add(-1)
	}
	barrier := func() {
		defer wg.Done()
		inBarrier.Store(true)
		if active.Load() != 0 {
			violations.Add(1)
		}
		time.Sleep(100 * time.Microsecond)
		if active.Load() != 0 {
			violations.Add(1)
		}
		inBarrier.Store(false)
	}

	wg.Add(readers + barriers)
	for i := range readers {
		q.Async(reader)
		if i%(readers/barriers) == 0 && i/(readers/barriers) < barriers {
			q.BarrierAsync(barrier)
		}
	}
	waitTimeout(t, &wg, 15*time.Second, "barrier/reader mix")
	if v := violations.Load(); v != 0 {
		t.Fatalf("%d barrier exclusivity violations", v)
	}
}

func TestConcurrentQueueActuallyConcurrent(t *testing.T) {
	// Prove the concurrent attr took effect: two blocking tasks must overlap.
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.overlap")
	var wg sync.WaitGroup
	rendezvous := make(chan struct{})
	overlapped := atomic.Bool{}
	wg.Add(2)
	for range 2 {
		q.Async(func() {
			defer wg.Done()
			select {
			case rendezvous <- struct{}{}:
				// partner picked it up: both were running at once
			case <-rendezvous:
				overlapped.Store(true)
			case <-time.After(3 * time.Second):
			}
		})
	}
	waitTimeout(t, &wg, 10*time.Second, "overlap probe")
	if !overlapped.Load() {
		t.Fatal("two tasks on a concurrent queue never ran simultaneously")
	}
}

func TestQueueSuspendResumeGating(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.conformance.suspend")
	q.Suspend()
	var ran atomic.Bool
	done := make(chan struct{})
	q.Async(func() {
		ran.Store(true)
		close(done)
	})
	time.Sleep(100 * time.Millisecond)
	if ran.Load() {
		t.Fatal("suspended queue executed work")
	}
	q.Resume()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("resumed queue never executed work")
	}
}

func TestSetTargetQueueSerializes(t *testing.T) {
	// Two serial queues targeting the same serial queue must never run
	// their work concurrently.
	target := QueueCreate("com.appledocs.dispatch.conformance.target")
	q1 := QueueCreateWithTarget("com.appledocs.dispatch.conformance.t1", SerialQueueAttr(), target)
	q2 := QueueCreateWithTarget("com.appledocs.dispatch.conformance.t2", SerialQueueAttr(), target)

	var active, violations atomic.Int64
	var wg sync.WaitGroup
	work := func() {
		defer wg.Done()
		if active.Add(1) > 1 {
			violations.Add(1)
		}
		time.Sleep(50 * time.Microsecond)
		active.Add(-1)
	}
	const per = 200
	wg.Add(2 * per)
	for range per {
		q1.Async(work)
		q2.Async(work)
	}
	waitTimeout(t, &wg, 15*time.Second, "target-queue work")
	if v := violations.Load(); v != 0 {
		t.Fatalf("%d concurrency violations through shared target queue", v)
	}
}

func TestAfterFiresNotBeforeDeadline(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.conformance.after")
	const delay = 100 * time.Millisecond
	start := time.Now()
	done := make(chan time.Duration, 1)
	After(TimeFromNow(int64(delay)), q, func() {
		done <- time.Since(start)
	})
	select {
	case elapsed := <-done:
		// Allow small clock-domain skew between mach_absolute_time and
		// Go's monotonic clock, but a grossly early fire is a real bug.
		if elapsed < delay-10*time.Millisecond {
			t.Fatalf("After fired at %v, before the %v deadline", elapsed, delay)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("After never fired")
	}
}

func TestSyncSeesCurrentQueue(t *testing.T) {
	// AssertCurrent inside Sync must not abort; AssertNotCurrent from
	// outside must not abort. (dispatch_assert_queue crashes the process on
	// violation, so simply surviving these calls is the assertion.)
	q := QueueCreate("com.appledocs.dispatch.conformance.assert")
	q.AssertNotCurrent()
	q.Sync(func() {
		q.AssertCurrent()
	})
}

func TestQueueLabelRoundTrip(t *testing.T) {
	labels := []string{
		"com.appledocs.dispatch.conformance.label",
		"unicode-λ-队列-🚀",
		"",
	}
	for _, label := range labels {
		q := QueueCreate(label)
		if got := q.Label(); got != label {
			t.Errorf("Label() = %q, want %q", got, label)
		}
	}
}

func TestApplyCoversAllIterationsExactlyOnce(t *testing.T) {
	const n = 512
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.apply")
	counts := make([]atomic.Int32, n)
	Apply(n, q, func(i int) {
		counts[i].Add(1)
	})
	for i := range counts {
		if c := counts[i].Load(); c != 1 {
			t.Fatalf("iteration %d ran %d times, want 1", i, c)
		}
	}
}

func TestGroupNotifyAfterAllWork(t *testing.T) {
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.group-notify")
	g := GroupCreate()
	const n = 100
	var completed atomic.Int64
	for range n {
		g.Async(q, func() {
			time.Sleep(time.Millisecond)
			completed.Add(1)
		})
	}
	notified := make(chan int64, 1)
	g.Notify(q, func() {
		notified <- completed.Load()
	})
	select {
	case seen := <-notified:
		if seen != n {
			t.Fatalf("Notify fired with %d/%d tasks complete", seen, n)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("group Notify never fired")
	}
}

func TestGroupWaitTimeoutExpires(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.conformance.group-timeout")
	g := GroupCreate()
	release := make(chan struct{})
	g.Async(q, func() { <-release })
	if g.Wait(TimeFromNow(int64(50 * time.Millisecond))) {
		t.Fatal("Wait returned true while work was still blocked")
	}
	close(release)
	if !g.Wait(TimeForever) {
		t.Fatal("Wait(TimeForever) returned false after work completed")
	}
}

func TestSemaphoreBoundsConcurrency(t *testing.T) {
	const capacity, workers = 4, 64
	sem := SemaphoreCreate(capacity)
	q := QueueCreateConcurrent("com.appledocs.dispatch.conformance.sema")
	var inside, violations atomic.Int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		q.Async(func() {
			defer wg.Done()
			if !sem.Wait(TimeForever) {
				violations.Add(1)
				return
			}
			if inside.Add(1) > capacity {
				violations.Add(1)
			}
			time.Sleep(time.Millisecond)
			inside.Add(-1)
			sem.Signal()
		})
	}
	waitTimeout(t, &wg, 30*time.Second, "semaphore workers")
	if v := violations.Load(); v != 0 {
		t.Fatalf("%d semaphore capacity violations", v)
	}
}

func TestSemaphoreWaitTimeout(t *testing.T) {
	sem := SemaphoreCreate(0)
	start := time.Now()
	if sem.Wait(TimeFromNow(int64(50 * time.Millisecond))) {
		t.Fatal("Wait on empty semaphore returned acquired")
	}
	if elapsed := time.Since(start); elapsed < 40*time.Millisecond {
		t.Fatalf("Wait returned after %v, before the 50ms timeout", elapsed)
	}
	sem.Signal()
	if !sem.Wait(TimeNow) {
		t.Fatal("Wait failed to acquire a signaled semaphore")
	}
}
