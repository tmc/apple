//go:build darwin

package irecovery

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func fakeUpload(mode string, events *[]string) *Conn {
	l := &library{}
	l.control = func(_ uintptr, kind, request uint8, value, index uint16, p *byte, n uint16, _ uint32) int32 {
		data := unsafe.Slice(p, int(n))
		switch {
		case kind == 0xa1 && request == 5:
			*events = append(*events, "state")
			data[0] = 2
		case kind == 0xa1 && request == 3:
			*events = append(*events, "status")
			copy(data, []byte{0, 0, 0, 0, 5, 0})
		case kind == 0x21 && request == 1:
			*events = append(*events, fmt.Sprintf("download %d %d", value, n))
		case kind == 0x41 && request == 0:
			*events = append(*events, "begin")
		default:
			panic("unexpected test transfer")
		}
		return int32(n)
	}
	l.bulk = func(_ uintptr, endpoint uint8, p *byte, n int32, actual *int32, _ uint32) int32 {
		*events = append(*events, fmt.Sprintf("bulk %d %d", endpoint, n))
		*actual = n
		return 0
	}
	return &Conn{library: l, info: Device{Mode: mode}}
}

func TestUpload_ProtocolFlow(t *testing.T) {
	for _, tt := range []struct {
		name, mode string
		size       int
		want       []string
	}{
		{"dfu small", "dfu", 1, []string{"state", "download 0 17", "status"}},
		{"dfu fits footer", "dfu", 2032, []string{"state", "download 0 2048", "status"}},
		{"dfu split footer", "dfu", 2033, []string{"state", "download 0 2033", "download 0 16", "status"}},
		{"dfu full block", "dfu", 2048, []string{"state", "download 0 2048", "download 0 16", "status"}},
		{"dfu two blocks", "dfu", 2049, []string{"state", "download 0 2048", "status", "download 1 17", "status"}},
		{"recovery small", "recovery", 1, []string{"begin", "bulk 4 1"}},
		{"recovery zlp", "recovery", 512, []string{"begin", "bulk 4 512", "bulk 4 0"}},
		{"recovery split", "recovery", 32769, []string{"begin", "bulk 4 32768", "bulk 4 1"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var events []string
			c := fakeUpload(tt.mode, &events)
			if err := c.Upload(context.Background(), make([]byte, tt.size)); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(events, tt.want) {
				t.Fatalf("got %v want %v", events, tt.want)
			}
		})
	}
}

func TestDFUFooterFixtures(t *testing.T) {
	// Values independently calculated with the pinned libirecovery CRC convention.
	for _, tt := range []struct {
		data []byte
		want string
	}{
		{[]byte("a"), "ffffffffac0500015546441065e31834"},
		{[]byte("123456789"), "ffffffffac05000155464410a4dac214"},
		{make([]byte, 2048), "ffffffffac050001554644100a8dae0e"},
	} {
		got := dfuFooter(tt.data)
		if hex.EncodeToString(got[:]) != tt.want {
			t.Fatalf("got %x want %s", got, tt.want)
		}
	}
}

func TestUploadPollTiming(t *testing.T) {
	var events []string
	c := fakeUpload("dfu", &events)
	original := c.library.control
	var first time.Time
	polls := 0
	c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
		if k != 0xa1 || r != 3 {
			return original(h, k, r, v, i, p, n, ms)
		}
		polls++
		data := unsafe.Slice(p, int(n))
		if polls == 1 {
			first = time.Now()
			copy(data, []byte{0, 25, 0, 0, 4, 0})
		} else {
			if time.Since(first) < 25*time.Millisecond {
				t.Error("polled before device timeout")
			}
			copy(data, []byte{0, 0, 0, 0, 5, 0})
		}
		return int32(n)
	}
	if err := c.Upload(context.Background(), []byte{1}); err != nil {
		t.Fatal(err)
	}
	if polls != 2 {
		t.Fatal(polls)
	}
}

func TestUploadErrors(t *testing.T) {
	for _, tt := range []struct {
		name, mode    string
		kind, request uint8
		short         bool
		status        byte
	}{
		{"short state", "dfu", 0xa1, 5, true, 0},
		{"short download", "dfu", 0x21, 1, true, 0},
		{"short status", "dfu", 0xa1, 3, true, 0},
		{"device error", "dfu", 0xa1, 3, false, 7},
		{"short bulk", "recovery", 0, 0, true, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var events []string
			c := fakeUpload(tt.mode, &events)
			original := c.library.control
			c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
				result := original(h, k, r, v, i, p, n, ms)
				if k == tt.kind && r == tt.request {
					if tt.short {
						return result - 1
					}
					unsafe.Slice(p, int(n))[0] = tt.status
				}
				return result
			}
			if tt.kind == 0 {
				c.library.bulk = func(_ uintptr, _ uint8, _ *byte, n int32, actual *int32, _ uint32) int32 { *actual = n - 1; return 0 }
			}
			err := c.Upload(context.Background(), []byte{1})
			if err == nil {
				t.Fatal("accepted transfer error")
			}
			if tt.status != 0 {
				var status *StatusError
				if !errors.As(err, &status) || status.Status != tt.status {
					t.Fatal(err)
				}
			} else if !errors.Is(err, io.ErrShortWrite) && !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatal(err)
			}
		})
	}
}

func TestUploadCancellationDuringPoll(t *testing.T) {
	var events []string
	c := fakeUpload("dfu", &events)
	original := c.library.control
	polls := 0
	c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
		if k == 0xa1 && r == 3 {
			polls++
			copy(unsafe.Slice(p, int(n)), []byte{0, 0, 1, 1, 4, 0})
			return int32(n)
		}
		return original(h, k, r, v, i, p, n, ms)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := c.Upload(ctx, []byte{1}); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	if polls != 1 {
		t.Fatalf("ignored 24-bit poll delay: %d polls", polls)
	}
}

func TestUploadLockCancellation(t *testing.T) {
	var c Conn
	c.mu.Lock()
	defer c.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := c.Upload(ctx, []byte{1}); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
}

func ExampleConn_Upload() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var c Conn
	fmt.Println(c.Upload(ctx, []byte{1}))
	// Output: context canceled
}
func ExampleStatusError_Error() {
	fmt.Println((&StatusError{Status: 7, State: 10}).Error())
	// Output: dfu status 7 in state 10
}

func TestUploadPayloadAndFooter(t *testing.T) {
	var events []string
	c := fakeUpload("dfu", &events)
	original := c.library.control
	var packet []byte
	c.library.control = func(h uintptr, k, r uint8, v, i uint16, p *byte, n uint16, ms uint32) int32 {
		if k == 0x21 && r == 1 {
			packet = append([]byte(nil), unsafe.Slice(p, int(n))...)
		}
		return original(h, k, r, v, i, p, n, ms)
	}
	data := []byte("123456789")
	if err := c.Upload(context.Background(), data); err != nil {
		t.Fatal(err)
	}
	if string(data) != "123456789" {
		t.Fatal("mutated caller's image")
	}
	if hex.EncodeToString(packet) != "313233343536373839ffffffffac05000155464410a4dac214" {
		t.Fatalf("unexpected transmitted bytes: %x", packet)
	}
}

func TestUploadInitialStateRecovery(t *testing.T) {
	for _, tt := range []struct{ state, request byte }{{10, 4}, {3, 6}} {
		t.Run(fmt.Sprint(tt.state), func(t *testing.T) {
			calls := 0
			c := &Conn{info: Device{Mode: "dfu"}, library: &library{}}
			c.library.control = func(_ uintptr, k, r uint8, _ uint16, _ uint16, p *byte, n uint16, _ uint32) int32 {
				calls++
				if calls == 1 {
					if k != 0xa1 || r != 5 {
						t.Error("expected GETSTATE")
					}
					*p = tt.state
					return 1
				}
				if calls != 2 || k != 0x21 || r != tt.request || n != 0 {
					t.Errorf("unexpected recovery request %x %d", k, r)
				}
				return 0
			}
			if err := c.Upload(context.Background(), []byte{1}); err == nil {
				t.Fatal("non-idle upload succeeded")
			}
			if calls != 2 {
				t.Fatal(calls)
			}
		})
	}
}

func TestUploadLimitIsDFUOnly(t *testing.T) {
	data := make([]byte, 0xffff*0x800+1)
	for _, mode := range []string{"dfu", "recovery"} {
		t.Run(mode, func(t *testing.T) {
			called := false
			c := &Conn{info: Device{Mode: mode}, library: &library{}}
			c.library.control = func(_ uintptr, _ uint8, _ uint8, _ uint16, _ uint16, _ *byte, _ uint16, _ uint32) int32 {
				called = true
				return -7
			}
			if err := c.Upload(context.Background(), data); err == nil {
				t.Fatal("expected rejected DFU size or injected recovery failure")
			}
			if called != (mode == "recovery") {
				t.Fatalf("mode %s reached USB: %v", mode, called)
			}
		})
	}
}
