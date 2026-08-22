//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
	"time"
)

func ioWorkMapCount() int {
	count := 0
	ioWorkMap.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}

func waitForIOWorkMapCount(t *testing.T, want int) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if got := ioWorkMapCount(); got == want {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("ioWorkMap entries = %d, want %d", ioWorkMapCount(), want)
}

func waitForIOCleanup(t *testing.T, cleanup <-chan error) {
	t.Helper()
	select {
	case err := <-cleanup:
		if err != nil {
			t.Fatalf("IO cleanup: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for IO cleanup")
	}
}

func TestIOStopDrainsWorkMap(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	before := ioWorkMapCount()
	q := QueueCreate("com.appledocs.dispatch.test.io-stop-drain")
	cleanup := make(chan error, 1)
	ch := IOCreate(IOStream, int(r.Fd()), q, func(err error) {
		cleanup <- err
	})
	done := make(chan struct{}, 1)
	ch.Read(0, 1, q, func(final bool, _ Data, _ error) {
		if final {
			done <- struct{}{}
		}
	})
	if got := ioWorkMapCount(); got != before+1 {
		t.Fatalf("ioWorkMap entries after Read = %d, want %d", got, before+1)
	}
	ch.Close(IOStop)
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for final IO callback after IOStop")
	}
	waitForIOWorkMapCount(t, before)
	waitForIOCleanup(t, cleanup)
	ch.Release()
}

func TestIODescriptorOwnershipHandoff(t *testing.T) {
	const iterations = 1000
	q := QueueCreate("com.appledocs.dispatch.test.io-descriptor-ownership")
	for i := 0; i < iterations; i++ {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		cleanup := make(chan error, 1)
		ch := IOCreate(IOStream, int(r.Fd()), q, func(err error) {
			cleanup <- err
		})
		ch.SetLowWater(1)
		started := make(chan struct{}, 1)
		done := make(chan error, 1)
		ch.Read(0, 1<<20, q, func(final bool, data Data, err error) {
			if data.Handle() != 0 && data.Len() > 0 {
				select {
				case started <- struct{}{}:
				default:
				}
			}
			if final {
				done <- err
			}
		})
		if _, err := w.Write([]byte{1}); err != nil {
			t.Fatal(err)
		}
		select {
		case <-started:
		case <-time.After(5 * time.Second):
			t.Fatalf("iteration %d: timeout waiting for read", i)
		}
		ch.Close(IOStop)
		select {
		case err := <-done:
			if err != syscall.ECANCELED {
				t.Fatalf("iteration %d: final callback error = %v, want ECANCELED", i, err)
			}
		case <-time.After(5 * time.Second):
			t.Fatalf("iteration %d: timeout waiting for final IO callback", i)
		}
		waitForIOCleanup(t, cleanup)
		ch.Release()
		if err := r.Close(); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestIOStreamRead(t *testing.T) {
	content := []byte("hello dispatch IO stream read")
	path := filepath.Join(t.TempDir(), "stream.txt")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.io-stream")
	var result []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	ch := IOCreate(IOStream, int(f.Fd()), q, func(err error) {})
	ch.Read(0, uint64(len(content)+1), q, func(done bool, data Data, err error) {
		if data.Len() > 0 {
			result = append(result, data.Bytes()...)
		}
		if err != nil {
			resultErr = err
		}
		if done {
			wg.Done()
		}
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for IO read")
	}
	if resultErr != nil {
		t.Fatalf("IO read error: %v", resultErr)
	}
	if string(result) != string(content) {
		t.Fatalf("IO read = %q, want %q", result, content)
	}
}

func TestReadFD(t *testing.T) {
	content := []byte("ReadFD convenience function test")
	path := filepath.Join(t.TempDir(), "readfd.txt")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.readfd")
	var result []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	ReadFD(int(f.Fd()), len(content)+1, q, func(data []byte, err error) {
		result = data
		resultErr = err
		wg.Done()
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for ReadFD")
	}
	if resultErr != nil {
		t.Fatalf("ReadFD error: %v", resultErr)
	}
	if string(result) != string(content) {
		t.Fatalf("ReadFD = %q, want %q", result, content)
	}
}

func TestReadBytes(t *testing.T) {
	content := []byte("dispatch read bytes test")
	path := filepath.Join(t.TempDir(), "read-bytes.txt")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.read-bytes")
	var result []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	ReadBytes(int(f.Fd()), len(content)+1, q, func(data []byte, err error) {
		result = data
		resultErr = err
		wg.Done()
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for ReadBytes")
	}
	if resultErr != nil {
		t.Fatalf("ReadBytes error: %v", resultErr)
	}
	if string(result) != string(content) {
		t.Fatalf("ReadBytes = %q, want %q", result, content)
	}
}

func TestIOCreateWithIO(t *testing.T) {
	content := []byte("hello dispatch IO create with io")
	path := filepath.Join(t.TempDir(), "stream-child.txt")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.io-create-with-io")
	base := IOCreate(IOStream, int(f.Fd()), q, func(err error) {})
	ch := IOCreateWithIO(IOStream, base, q, func(err error) {})
	var result []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	ch.Read(0, uint64(len(content)+1), q, func(done bool, data Data, err error) {
		if data.Len() > 0 {
			result = append(result, data.Bytes()...)
		}
		if err != nil {
			resultErr = err
		}
		if done {
			wg.Done()
		}
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for IOCreateWithIO read")
	}
	if resultErr != nil {
		t.Fatalf("IOCreateWithIO read error: %v", resultErr)
	}
	if string(result) != string(content) {
		t.Fatalf("IOCreateWithIO read = %q, want %q", result, content)
	}
}

func TestIOSetInterval(t *testing.T) {
	content := []byte("hello dispatch IO interval")
	path := filepath.Join(t.TempDir(), "interval.txt")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.io-interval")
	ch := IOCreate(IOStream, int(f.Fd()), q, func(err error) {})
	ch.SetInterval(10*time.Millisecond, IOStrictInterval)
	var result []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	ch.Read(0, uint64(len(content)+1), q, func(done bool, data Data, err error) {
		if data.Len() > 0 {
			result = append(result, data.Bytes()...)
		}
		if err != nil {
			resultErr = err
		}
		if done {
			wg.Done()
		}
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for IO interval read")
	}
	if resultErr != nil {
		t.Fatalf("IO interval read error: %v", resultErr)
	}
	if string(result) != string(content) {
		t.Fatalf("IO interval read = %q, want %q", result, content)
	}
}

func TestWriteBytes(t *testing.T) {
	content := []byte("dispatch write bytes test")
	path := filepath.Join(t.TempDir(), "write-bytes.txt")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	q := QueueCreate("com.appledocs.dispatch.test.write-bytes")
	var remaining []byte
	var resultErr error
	var wg sync.WaitGroup
	wg.Add(1)
	WriteBytes(int(f.Fd()), content, q, func(data []byte, err error) {
		remaining = data
		resultErr = err
		wg.Done()
	})
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for WriteBytes")
	}
	if resultErr != nil {
		t.Fatalf("WriteBytes error: %v", resultErr)
	}
	if len(remaining) != 0 {
		t.Fatalf("WriteBytes remaining = %q, want empty", remaining)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(content) {
		t.Fatalf("file contents = %q, want %q", got, content)
	}
}
