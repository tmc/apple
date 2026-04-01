//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const dispatchOperationStressCount = 4096

func runManyAsyncOperations(tb testing.TB, count int) {
	tb.Helper()

	q := QueueCreate("com.appledocs.dispatch.test.async-stress")
	var completed atomic.Int64
	var wg sync.WaitGroup
	wg.Add(count)
	for range count {
		q.Async(func() {
			completed.Add(1)
			wg.Done()
		})
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		tb.Fatalf("timeout waiting for %d dispatch operations", count)
	}
	if got := completed.Load(); got != int64(count) {
		tb.Fatalf("completed operations = %d, want %d", got, count)
	}
}

func TestQueueCreateInactive(t *testing.T) {
	q := QueueCreateWithAttr("com.appledocs.dispatch.test.inactive", SerialQueueAttr().InitiallyInactive())
	var ran atomic.Bool
	q.Async(func() {
		ran.Store(true)
	})
	time.Sleep(50 * time.Millisecond)
	if ran.Load() {
		t.Fatal("inactive queue ran work before activation")
	}
	q.Activate()
	time.Sleep(100 * time.Millisecond)
	if !ran.Load() {
		t.Fatal("inactive queue did not run work after activation")
	}
}

func TestQueueQOSClass(t *testing.T) {
	attr, err := SerialQueueAttr().WithQOS(QOSUtility, 0)
	if err != nil {
		t.Fatal(err)
	}
	q := QueueCreateWithAttr("com.appledocs.dispatch.test.qos", attr)
	qos, relativePriority := q.QOSClass()
	if qos != QOSUtility || relativePriority != 0 {
		t.Fatalf("QOSClass() = (%d, %d), want (%d, 0)", qos, relativePriority, QOSUtility)
	}
}

func TestQueueAsyncAndWait(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.test.async-and-wait")
	var value atomic.Int64
	q.AsyncAndWait(func() {
		value.Store(42)
	})
	if value.Load() != 42 {
		t.Fatalf("AsyncAndWait value = %d, want 42", value.Load())
	}
}

func TestApplyAuto(t *testing.T) {
	var sum atomic.Int64
	ApplyAuto(10, func(iteration int) {
		sum.Add(int64(iteration))
	})
	if sum.Load() != 45 {
		t.Fatalf("ApplyAuto sum = %d, want 45", sum.Load())
	}
}

func TestQueueAsyncMoreThanPuregoCallbackLimit(t *testing.T) {
	runManyAsyncOperations(t, dispatchOperationStressCount)
}

func TestWorkItemWaitAndNotify(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.test.work-item")
	var ran atomic.Bool
	item := WorkItemCreate(0, func() {
		ran.Store(true)
	})
	defer item.Release()

	done := make(chan struct{}, 1)
	item.Notify(q, func() {
		select {
		case done <- struct{}{}:
		default:
		}
	})
	q.AsyncWorkItem(item)
	if !item.Wait(TimeFromNow(int64(5 * time.Second))) {
		t.Fatal("timeout waiting for work item")
	}
	if !ran.Load() {
		t.Fatal("work item did not run")
	}
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for work item notification")
	}
}

func TestWorkItemCancel(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.test.work-item-cancel")
	var ran atomic.Bool
	item := WorkItemCreate(0, func() {
		ran.Store(true)
	})
	defer item.Release()

	item.Cancel()
	if !item.IsCancelled() {
		t.Fatal("work item should be cancelled")
	}
	q.AsyncWorkItem(item)
	if !item.Wait(TimeFromNow(int64(5 * time.Second))) {
		t.Fatal("timeout waiting for cancelled work item")
	}
	if ran.Load() {
		t.Fatal("cancelled work item still ran")
	}
}

func BenchmarkQueueAsyncOverCallbackLimit(b *testing.B) {
	for b.Loop() {
		runManyAsyncOperations(b, dispatchOperationStressCount)
	}
}
