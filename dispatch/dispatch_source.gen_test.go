//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"
)

func TestTimerSource(t *testing.T) {
	var count atomic.Int64
	q := QueueCreate("com.appledocs.dispatch.test.timer")
	s := NewTimerSource(10*time.Millisecond, 0, q, func() {
		count.Add(1)
	})
	defer s.Cancel()
	time.Sleep(100 * time.Millisecond)
	got := count.Load()
	if got < 3 {
		t.Fatalf("timer fired %d times, want >= 3", got)
	}
}

func TestTimerCancel(t *testing.T) {
	var count atomic.Int64
	q := QueueCreate("com.appledocs.dispatch.test.timer-cancel")
	s := NewTimerSource(10*time.Millisecond, 0, q, func() {
		count.Add(1)
	})
	time.Sleep(50 * time.Millisecond)
	s.Cancel()
	snapshot := count.Load()
	time.Sleep(50 * time.Millisecond)
	if count.Load() != snapshot {
		t.Fatal("timer continued firing after cancel")
	}
}

func TestDataAddSource(t *testing.T) {
	var total atomic.Uint64
	q := QueueCreate("com.appledocs.dispatch.test.data-add")
	s := SourceCreate(SourceTypeDataAdd, 0, 0, q)
	s.SetEventHandler(func() {
		total.Add(uint64(s.GetData()))
	})
	s.Activate()
	s.MergeData(10)
	s.MergeData(20)
	s.MergeData(30)
	time.Sleep(100 * time.Millisecond)
	got := total.Load()
	if got < 30 {
		t.Fatalf("data add total = %d, want >= 30", got)
	}
	s.Cancel()
}

func TestSourceIsCancelled(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.test.cancelled")
	s := SourceCreate(SourceTypeDataAdd, 0, 0, q)
	s.SetEventHandler(func() {})
	s.Activate()
	if s.IsCancelled() {
		t.Fatal("source should not be cancelled yet")
	}
	s.Cancel()
	time.Sleep(50 * time.Millisecond)
	if !s.IsCancelled() {
		t.Fatal("source should be cancelled")
	}
}

func TestSourceCancelCleanup(t *testing.T) {
	var cleaned atomic.Bool
	q := QueueCreate("com.appledocs.dispatch.test.cancel-cleanup")
	s := SourceCreate(SourceTypeDataAdd, 0, 0, q)
	s.SetEventHandler(func() {})
	s.SetCancelHandler(func() {
		cleaned.Store(true)
	})
	s.Activate()
	s.Cancel()
	time.Sleep(100 * time.Millisecond)
	if !cleaned.Load() {
		t.Fatal("cancel handler was not called")
	}
}

func TestSourceGetHandleAndMask(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "dispatch-source-*")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.handle-mask")
	mask := VnodeWrite | VnodeDelete
	s := SourceCreate(SourceTypeVnode, uintptr(f.Fd()), uintptr(mask), q)
	if got := s.GetHandle(); got != uintptr(f.Fd()) {
		t.Fatalf("GetHandle() = %d, want %d", got, f.Fd())
	}
	if got := s.GetMask(); got != uintptr(mask) {
		t.Fatalf("GetMask() = %#x, want %#x", got, mask)
	}
	s.Cancel()
}

func TestDataReplaceSource(t *testing.T) {
	if SourceTypeDataReplace == 0 {
		t.Skip("dispatch data replace sources unavailable")
	}

	var got atomic.Uint64
	done := make(chan struct{}, 1)
	q := QueueCreate("com.appledocs.dispatch.test.data-replace")
	s := SourceCreate(SourceTypeDataReplace, 0, 0, q)
	s.SetEventHandler(func() {
		got.Store(uint64(s.GetData()))
		select {
		case done <- struct{}{}:
		default:
		}
	})
	s.MergeData(10)
	s.MergeData(20)
	s.Activate()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for data replace source")
	}
	if got.Load() != 20 {
		t.Fatalf("GetData() = %d, want 20", got.Load())
	}
	s.Cancel()
}

func TestReadSource(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	q := QueueCreate("com.appledocs.dispatch.test.read-source")
	done := make(chan uintptr, 1)
	var s Source
	s = NewReadSource(int(r.Fd()), q, func(bytesAvailable uintptr) {
		select {
		case done <- bytesAvailable:
		default:
		}
		s.Cancel()
	})
	if _, err := w.Write([]byte("hello")); err != nil {
		t.Fatal(err)
	}
	select {
	case bytesAvailable := <-done:
		if bytesAvailable == 0 {
			t.Fatal("read source reported zero bytes available")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for read source")
	}
}

func TestProcessSource(t *testing.T) {
	cmd := exec.Command("sleep", "30")
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	q := QueueCreate("com.appledocs.dispatch.test.process-source")
	done := make(chan ProcFlags, 1)
	var s Source
	s = NewProcessSource(cmd.Process.Pid, ProcExit, q, func(events ProcFlags) {
		select {
		case done <- events:
		default:
		}
		s.Cancel()
	})
	if err := cmd.Process.Kill(); err != nil {
		t.Fatal(err)
	}
	_ = cmd.Wait()
	select {
	case events := <-done:
		if events&ProcExit == 0 {
			t.Fatalf("process source events = %#x, want %#x", events, ProcExit)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for process source")
	}
}

func TestVnodeSource(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dispatch-vnode.txt")
	if err := os.WriteFile(path, []byte("before"), 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.vnode-source")
	done := make(chan VnodeFlags, 1)
	var s Source
	s = NewVnodeSource(int(f.Fd()), VnodeWrite, q, func(events VnodeFlags) {
		select {
		case done <- events:
		default:
		}
		s.Cancel()
	})
	if err := os.WriteFile(path, []byte("after"), 0644); err != nil {
		t.Fatal(err)
	}
	select {
	case events := <-done:
		if events&VnodeWrite == 0 {
			t.Fatalf("vnode source events = %#x, want %#x", events, VnodeWrite)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for vnode source")
	}
}
