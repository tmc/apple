//go:build darwin

package irecovery

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"unsafe"
)

func TestSendCommand(t *testing.T) {
	for _, tt := range []struct {
		name, mode, command string
		request             uint8
		count               int
		bad                 bool
	}{
		{"normal", "recovery", "setenv auto-boot false", 0, -1, false},
		{"go", "recovery", "go", 1, -1, false},
		{"DFU", "dfu", "go", 1, -1, true},
		{"short", "recovery", "go", 1, 1, true},
		{"empty", "recovery", "", 0, -1, true},
		{"long", "recovery", strings.Repeat("x", 256), 0, -1, true},
		{"NUL", "recovery", "go\x00reboot", 0, -1, true},
		{"newline", "recovery", "go\nreboot", 0, -1, true},
		{"request", "recovery", "go", 2, -1, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			c := &Conn{info: Device{Mode: tt.mode}, library: &library{control: func(_ uintptr, kind, request uint8, value, index uint16, p *byte, size uint16, _ uint32) int32 {
				calls++
				if kind != 0x40 || request != tt.request || value != 0 || index != 0 || string(unsafe.Slice(p, int(size))) != tt.command+"\x00" {
					t.Fatal("incorrect command framing")
				}
				if tt.count >= 0 {
					return int32(tt.count)
				}
				return int32(size)
			}}}
			err := c.SendCommand(context.Background(), tt.command, tt.request)
			if (err != nil) != tt.bad {
				t.Fatalf("got %v", err)
			}
			if calls > 1 {
				t.Fatal("retried command")
			}
			if tt.name == "short" && !errors.Is(err, io.ErrShortWrite) {
				t.Fatal(err)
			}
		})
	}
}
func TestGetenv(t *testing.T) {
	for _, tt := range []struct {
		name, response, want string
		bad                  bool
	}{
		{"plain", "value", "value", false}, {"padded", "value\x00\x00", "value", false}, {"empty", "", "", false},
		{"embedded data", "x\x00y", "", true}, {"full", strings.Repeat("x", 255), "", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			c := &Conn{info: Device{Mode: "recovery"}, library: &library{control: func(_ uintptr, kind, request uint8, value, index uint16, p *byte, size uint16, _ uint32) int32 {
				calls++
				if calls == 1 {
					if kind != 0x40 || string(unsafe.Slice(p, int(size))) != "getenv boot-args\x00" {
						t.Fatal("bad getenv command")
					}
					return int32(size)
				}
				if kind != 0xc0 || request != 0 || value != 0 || index != 0 || size != 255 {
					t.Fatal("bad getenv response request")
				}
				return int32(copy(unsafe.Slice(p, int(size)), tt.response))
			}}}
			got, err := c.Getenv(context.Background(), "boot-args")
			if (err != nil) != tt.bad || got != tt.want || calls != 2 {
				t.Fatalf("got %q %v, calls %d", got, err, calls)
			}
		})
	}
}
func ExampleConn_SendCommand() {
	var c Conn
	fmt.Println(c.SendCommand(context.Background(), "go", 1))
	// Output: recovery connection is closed
}
func ExampleConn_Getenv() {
	var c Conn
	_, err := c.Getenv(context.Background(), "boot-args")
	fmt.Println(err)
	// Output: recovery connection is closed
}
