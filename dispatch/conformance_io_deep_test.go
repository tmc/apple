//go:build darwin

package dispatch

// Deeper IO and source conformance tests: the IO channel surface under real
// transfers (large payloads, both directions, water marks, cancellation),
// the IO barrier contract, and the source types the basic suite never
// exercised (signal, suspend/resume gating).

import (
	"bytes"
	"crypto/sha256"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestIOBarrierRunsAfterWrites(t *testing.T) {
	// dispatch_io_barrier takes a block; the barrier closure must actually
	// execute after prior operations on the channel complete.
	f, err := os.CreateTemp(t.TempDir(), "io-barrier-*")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.deep.io-barrier")
	ch := IOCreate(IORandom, int(f.Fd()), q, func(error) {})
	defer ch.Close(0)

	payload := patternBytes(1<<16, 7)
	d := DataCreate(payload)
	defer d.Release()
	ch.Write(0, d, q, func(bool, Data, error) {})

	// The barrier orders against the channel's I/O operations (not against
	// handler delivery), so the observable postcondition is that the bytes
	// are in the file when the barrier body runs.
	fired := make(chan int64, 1)
	ch.Barrier(func() {
		st, err := f.Stat()
		if err != nil {
			fired <- -1
			return
		}
		fired <- st.Size()
	})
	select {
	case size := <-fired:
		if size != int64(len(payload)) {
			t.Fatalf("barrier observed file size %d, want %d (ran before prior write completed)", size, len(payload))
		}
	case <-time.After(10 * time.Second):
		t.Fatal("IO barrier closure never ran")
	}
}

func TestIOLargeTransferRoundTrip(t *testing.T) {
	// 4MB through a file: write via one channel, read via another, verify
	// byte-identity. Exercises multi-chunk delivery (handler called more
	// than once) and offset-addressed IORandom channels.
	const size = 4 << 20
	f, err := os.CreateTemp(t.TempDir(), "io-large-*")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.deep.io-large")
	payload := patternBytes(size, 11)
	wantSum := sha256.Sum256(payload)

	wch := IOCreate(IORandom, int(f.Fd()), q, func(error) {})
	d := DataCreate(payload)
	defer d.Release()
	wdone := make(chan error, 1)
	wch.Write(0, d, q, func(done bool, _ Data, err error) {
		if done {
			wdone <- err
		}
	})
	select {
	case err := <-wdone:
		if err != nil {
			t.Fatalf("write failed: %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("timeout writing 4MB")
	}
	wch.Close(0)

	rch := IOCreate(IORandom, int(f.Fd()), q, func(error) {})
	defer rch.Close(0)
	rch.SetHighWater(256 << 10) // force multiple delivery callbacks
	var got []byte
	var deliveries atomic.Int64
	rdone := make(chan error, 1)
	rch.Read(0, size, q, func(done bool, data Data, err error) {
		deliveries.Add(1)
		if data.Len() > 0 {
			got = append(got, data.Bytes()...)
		}
		if done {
			rdone <- err
		}
	})
	select {
	case err := <-rdone:
		if err != nil {
			t.Fatalf("read failed: %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("timeout reading 4MB")
	}
	if len(got) != size {
		t.Fatalf("read %d bytes, want %d", len(got), size)
	}
	if sha256.Sum256(got) != wantSum {
		t.Fatal("4MB IO round-trip corrupted data")
	}
	if deliveries.Load() < 2 {
		t.Fatalf("expected multiple delivery callbacks with 256KB high water, got %d", deliveries.Load())
	}
}

func TestIOStopCancelsPendingRead(t *testing.T) {
	// Start a read on a pipe, then leave it waiting for more data.
	// Close(IOStop) must cancel it and deliver done=true with ECANCELED.
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	q := QueueCreate("com.appledocs.dispatch.deep.io-stop")
	cleanup := make(chan error, 1)
	ch := IOCreate(IOStream, int(r.Fd()), q, func(err error) {
		cleanup <- err
	})
	ch.SetLowWater(1)
	started := make(chan struct{}, 1)
	done := make(chan error, 1)
	ch.Read(0, 1<<20, q, func(isDone bool, data Data, err error) {
		if data.Handle() != 0 && data.Len() > 0 {
			select {
			case started <- struct{}{}:
			default:
			}
		}
		if isDone {
			done <- err
		}
	})
	if _, err := w.Write([]byte{1}); err != nil {
		t.Fatal(err)
	}
	select {
	case <-started:
	case <-time.After(10 * time.Second):
		t.Fatal("read did not start")
	}
	ch.Close(IOStop)
	select {
	case err := <-done:
		if err != syscall.ECANCELED {
			t.Fatalf("cancelled read delivered err=%v, want ECANCELED", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("Close(IOStop) did not cancel the pending read")
	}
	select {
	case err := <-cleanup:
		if err != nil {
			t.Fatalf("IO cleanup: %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("timeout waiting for IO cleanup")
	}
	_dispatch_release(uintptr(ch.io))
	if err := r.Close(); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestIOPipeStreaming(t *testing.T) {
	// Stream 1MB through a pipe with concurrent writer and reader channels.
	const size = 1 << 20
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	q := QueueCreateConcurrent("com.appledocs.dispatch.deep.io-pipe")
	payload := patternBytes(size, 13)

	wch := IOCreate(IOStream, int(w.Fd()), q, func(error) {})
	rch := IOCreate(IOStream, int(r.Fd()), q, func(error) {})
	defer rch.Close(0)

	d := DataCreate(payload)
	defer d.Release()
	wch.Write(0, d, q, func(done bool, _ Data, err error) {
		if done {
			if err != nil {
				t.Errorf("pipe write failed: %v", err)
			}
			wch.Close(0)
			w.Close() // EOF for the reader
		}
	})

	var got []byte
	rdone := make(chan error, 1)
	rch.Read(0, size, q, func(done bool, data Data, err error) {
		if data.Len() > 0 {
			got = append(got, data.Bytes()...)
		}
		if done {
			rdone <- err
		}
	})
	select {
	case err := <-rdone:
		if err != nil {
			t.Fatalf("pipe read failed: %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("timeout streaming 1MB through pipe")
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("pipe streaming corrupted data (got %d bytes, want %d)", len(got), size)
	}
}

func TestSignalSourceDeliversCount(t *testing.T) {
	// SIGUSR2 routed through a dispatch signal source. Go's runtime must
	// not swallow it: signal sources use kqueue EVFILT_SIGNAL, which
	// observes delivery regardless of handler disposition — but the signal
	// must not kill the process, so ignore it at the OS level first.
	signal.Ignore(syscall.SIGUSR2)
	t.Cleanup(func() { signal.Reset(syscall.SIGUSR2) })

	q := QueueCreate("com.appledocs.dispatch.deep.signal")
	var count atomic.Int64
	src := NewSignalSource(syscall.SIGUSR2, q, func(n uintptr) {
		count.Add(int64(n))
	})
	defer src.Cancel()
	time.Sleep(50 * time.Millisecond) // let the source register

	const sends = 5
	for range sends {
		if err := syscall.Kill(os.Getpid(), syscall.SIGUSR2); err != nil {
			t.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	deadline := time.Now().Add(5 * time.Second)
	for count.Load() < sends && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if got := count.Load(); got < sends {
		t.Fatalf("signal source observed %d signals, want >= %d", got, sends)
	}
}

func TestSourceSuspendResumeGating(t *testing.T) {
	q := QueueCreate("com.appledocs.dispatch.deep.src-gate")
	var fired atomic.Int64
	src := NewDataAddSource(q, func(uintptr) {
		fired.Add(1)
	})
	defer src.Cancel()

	src.MergeData(1)
	deadline := time.Now().Add(5 * time.Second)
	for fired.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if fired.Load() == 0 {
		t.Fatal("active source never fired")
	}

	src.Suspend()
	before := fired.Load()
	src.MergeData(1)
	time.Sleep(100 * time.Millisecond)
	if fired.Load() != before {
		t.Fatal("suspended source fired")
	}
	src.Resume()
	deadline = time.Now().Add(5 * time.Second)
	for fired.Load() == before && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if fired.Load() == before {
		t.Fatal("resumed source never delivered the pending event")
	}
}

func TestAfterBurst(t *testing.T) {
	// 500 After timers with staggered deadlines, all through the shared
	// trampoline; every one must fire exactly once.
	q := QueueCreate("com.appledocs.dispatch.deep.after-burst")
	const n = 500
	var fired atomic.Int64
	for i := range n {
		After(TimeFromNow(int64(time.Duration(i%50)*time.Millisecond)), q, func() {
			fired.Add(1)
		})
	}
	deadline := time.Now().Add(15 * time.Second)
	for fired.Load() < n && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if got := fired.Load(); got != n {
		t.Fatalf("After burst fired %d/%d", got, n)
	}
}
