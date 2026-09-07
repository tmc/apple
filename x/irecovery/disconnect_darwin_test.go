//go:build darwin

package irecovery

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestWaitDisconnected(t *testing.T) {
	for _, tt := range []struct {
		name string
		code int32
		bad  bool
	}{{"removed", -4, false}, {"timeout", -7, false}, {"pipe", -9, false}, {"access", -3, true}, {"short", 2, true}} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			c := &Conn{library: &library{control: func(_ uintptr, kind, request uint8, value, index uint16, p *byte, size uint16, _ uint32) int32 {
				calls++
				if kind != 0x80 || request != 6 || value != 0x0100 || index != 0 || size != 18 {
					t.Fatal("incorrect presence request")
				}
				if calls == 1 {
					b := unsafe.Slice(p, 18)
					b[0] = 18
					b[1] = 1
					return 18
				}
				if calls == 2 {
					return tt.code
				}
				return -4
			}}}
			err := c.WaitDisconnected(context.Background())
			wantCalls := 2
			if tt.code == -7 || tt.code == -9 {
				wantCalls = 3
			}
			if (err != nil) != tt.bad || calls != wantCalls {
				t.Fatalf("got %v, calls %d", err, calls)
			}
			if c.library == nil {
				t.Fatal("closed caller's connection")
			}
		})
	}
}
func ExampleConn_WaitDisconnected() {
	var c Conn
	fmt.Println(c.WaitDisconnected(context.Background()))
	// Output: recovery connection is closed
}

func TestWaitDisconnectedTransientDeadline(t *testing.T) {
	for _, code := range []int32{-7, -9} {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		c := &Conn{library: &library{control: func(uintptr, uint8, uint8, uint16, uint16, *byte, uint16, uint32) int32 { return code }}}
		err := c.WaitDisconnected(ctx)
		cancel()
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("code %d treated as removal: %v", code, err)
		}
	}
}
