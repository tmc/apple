//go:build darwin

package irecovery

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"slices"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func finalizer(events *[]string) *Conn {
	l := &library{}
	l.control = func(_ uintptr, k, r uint8, v, _ uint16, p *byte, n uint16, _ uint32) int32 {
		if k == 0x21 && r == 1 {
			*events = append(*events, fmt.Sprintf("notify %d %d", v, n))
			return int32(n)
		}
		*events = append(*events, "status")
		copy(unsafe.Slice(p, int(n)), []byte{0, 0, 0, 0, 8, 0})
		return int32(n)
	}
	l.reset = func(uintptr) int32 { *events = append(*events, "reset"); return 0 }
	l.release = func(uintptr, int32) int32 { *events = append(*events, "release"); return 0 }
	l.close = func(uintptr) { *events = append(*events, "close") }
	l.exit = func(uintptr) { *events = append(*events, "exit") }
	return &Conn{library: l, handle: 42, info: Device{Mode: "dfu", ECID: 123}, dfuBlocks: 2}
}

func TestFinalizeDFUSequence(t *testing.T) {
	var events []string
	c := finalizer(&events)
	if err := c.FinalizeDFU(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := c.FinalizeDFU(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
	want := []string{"notify 2 0", "status", "reset", "release", "close", "exit"}
	if !slices.Equal(events, want) {
		t.Fatal(events)
	}
	if _, err := c.Control(context.Background(), 0, 0, 0, 0, nil); err == nil {
		t.Fatal("retired handle reused")
	}
}

func TestFinalizeCancellationRetainsHandle(t *testing.T) {
	var events []string
	c := finalizer(&events)
	entered, finish := make(chan struct{}), make(chan struct{})
	c.library.reset = func(uintptr) int32 { close(entered); <-finish; return 0 }
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() { result <- c.FinalizeDFU(ctx) }()
	<-entered
	cancel()
	if err := <-result; !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	if !slices.Equal(events, []string{"notify 2 0", "status"}) {
		t.Fatalf("freed during reset: %v", events)
	}
	if _, err := c.Control(context.Background(), 0, 0, 0, 0, nil); err == nil {
		t.Fatal("retired handle reused")
	}
	close(finish)
	if err := c.FinalizeDFU(context.Background()); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(events, []string{"notify 2 0", "status", "release", "close", "exit"}) {
		t.Fatal(events)
	}
}

func TestFinalizeFailures(t *testing.T) {
	for _, tt := range []struct {
		name      string
		reset     int32
		status    byte
		wantError bool
	}{
		{"success", 0, 0, false}, {"rediscover", -5, 0, false}, {"io error", -1, 0, true}, {"no device", -4, 0, true}, {"status error", 0, 7, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var events []string
			c := finalizer(&events)
			reset := false
			c.library.reset = func(uintptr) int32 { reset = true; return tt.reset }
			if tt.status != 0 {
				original := c.library.control
				c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
					result := original(h, k, r, v, i, p, n, ms)
					if k == 0xa1 {
						*p = tt.status
					}
					return result
				}
			}
			err := c.FinalizeDFU(context.Background())
			if (err != nil) != tt.wantError {
				t.Fatal(err)
			}
			if tt.status != 0 && reset {
				t.Fatal("reset after device status error")
			}
			if events[len(events)-1] != "exit" {
				t.Fatal("missing cleanup", events)
			}
		})
	}
}

func TestFinalizePollTiming(t *testing.T) {
	var events []string
	c := finalizer(&events)
	original := c.library.control
	polls := 0
	var first time.Time
	c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
		if k != 0xa1 {
			return original(h, k, r, v, i, p, n, ms)
		}
		polls++
		if polls == 1 {
			first = time.Now()
			copy(unsafe.Slice(p, int(n)), []byte{0, 20, 0, 0, 7, 0})
		} else {
			if time.Since(first) < 20*time.Millisecond {
				t.Error("ignored poll timeout")
			}
			copy(unsafe.Slice(p, int(n)), []byte{0, 0, 0, 0, 8, 0})
		}
		return int32(n)
	}
	if err := c.FinalizeDFU(context.Background()); err != nil {
		t.Fatal(err)
	}
	if polls != 2 {
		t.Fatal(polls)
	}
}

func TestWaitOpenRetriesOnlyAbsence(t *testing.T) {
	attempts := 0
	want := &Conn{}
	got, err := waitOpen(context.Background(), func(context.Context) (*Conn, error) {
		attempts++
		if attempts == 1 {
			return nil, ErrNotFound
		}
		return want, nil
	})
	if err != nil || got != want || attempts != 2 {
		t.Fatal(got, err, attempts)
	}
	denied := errors.New("permission denied")
	if _, err := waitOpen(context.Background(), func(context.Context) (*Conn, error) { return nil, denied }); !errors.Is(err, denied) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := WaitOpen(ctx, "unused", 123, "recovery"); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}

func TestSelectDeviceRejectsAmbiguityAndWrongMode(t *testing.T) {
	for _, tt := range []struct {
		name      string
		devices   int
		mode      string
		wantClaim bool
	}{
		{"one", 1, "dfu", true}, {"duplicate", 2, "dfu", false}, {"wrong mode", 1, "recovery", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// Use native memory because libusb returns an address, not a Go pointer.
			memory, err := syscall.Mmap(-1, 0, 16, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if err := syscall.Munmap(memory); err != nil {
					t.Error(err)
				}
			})
			binary.LittleEndian.PutUint64(memory, 1)
			binary.LittleEndian.PutUint64(memory[8:], 2)
			closed := 0
			claimed := 0
			l := &library{}
			l.list = func(_ uintptr, p *uintptr) int64 { *p = uintptr(unsafe.Pointer(&memory[0])); return int64(tt.devices) }
			l.freeList = func(uintptr, int32) {}
			l.descriptor = func(_ uintptr, d *[18]byte) int32 {
				d[8] = 0xac
				d[9] = 5
				d[10] = 0x27
				d[11] = 0x12
				d[16] = 1
				return 0
			}
			l.open = func(d uintptr, h *uintptr) int32 { *h = d; return 0 }
			l.serial = func(_ uintptr, _ uint8, p *byte, n int32) int32 {
				return int32(copy(unsafe.Slice(p, int(n)), "ECID:7b CPID:fe01"))
			}
			l.close = func(uintptr) { closed++ }
			l.claim = func(uintptr, int32) int32 { claimed++; return 0 }
			conn, err := l.selectDevice(context.Background(), 123, tt.mode)
			if (err == nil) != tt.wantClaim {
				t.Fatal(conn, err)
			}
			if tt.wantClaim {
				if claimed != 1 || closed != 0 {
					t.Fatal(claimed, closed)
				}
				l.close(conn.handle)
			} else if claimed != 0 || closed != tt.devices {
				t.Fatal(claimed, closed)
			}
		})
	}
}

func ExampleConn_FinalizeDFU() {
	var c Conn
	fmt.Println(c.FinalizeDFU(context.Background()))
	// Output: recovery connection is closed
}
func ExampleWaitOpen() {
	_, err := WaitOpen(context.Background(), "libusb.dylib", 0, "recovery")
	fmt.Println(err)
	// Output: nonzero ECID is required
}
