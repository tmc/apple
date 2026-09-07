//go:build darwin

package irecovery

import (
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func descriptor(text string) []byte {
	b := make([]byte, 2+2*len(text))
	b[0] = byte(len(b))
	b[1] = 3
	for i := range text {
		binary.LittleEndian.PutUint16(b[2+2*i:], uint16(text[i]))
	}
	return b
}
func infoConnection(serial, nonce string, calls *[]uint16) *Conn {
	l := &library{control: func(_ uintptr, kind, request uint8, value, index uint16, p *byte, size uint16, timeout uint32) int32 {
		if kind != 0x80 || request != 6 || timeout == 0 {
			panic("unexpected descriptor request")
		}
		*calls = append(*calls, value)
		var b []byte
		switch value {
		case 0x0300:
			if index != 0 {
				panic("language index")
			}
			b = []byte{4, 3, 9, 4}
		case 0x0303:
			if index != 0x0409 {
				panic("serial language")
			}
			b = descriptor(serial)
		case 0x0301:
			if index != 0x0409 {
				panic("nonce language")
			}
			b = descriptor(nonce)
		default:
			panic("unexpected descriptor index")
		}
		return int32(copy(unsafe.Slice(p, int(size)), b))
	}}
	return &Conn{library: l, handle: 1, info: Device{ECID: 1<<63 | 17, CPID: 0xfe01, Mode: "dfu", serialIndex: 3}, dfuBlocks: 2}
}
func TestReadInfo(t *testing.T) {
	var calls []uint16
	conn := infoConnection("ECID:8000000000000011 CPID:FE01 BDID:00 CPFM:03 IBFL:04", "NONC:010203 SNON:aabb", &calls)
	got, err := conn.ReadInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got.Device.ECID != 1<<63|17 || *got.BoardID != 0 || *got.CPFM != 3 || *got.IBFL != 4 || !reflect.DeepEqual(got.APNonce, []byte{1, 2, 3}) || !reflect.DeepEqual(got.SEPNonce, []byte{0xaa, 0xbb}) {
		t.Fatalf("got %#v", got)
	}
	if !reflect.DeepEqual(calls, []uint16{0x0300, 0x0303, 0x0301}) || conn.dfuBlocks != 2 {
		t.Fatalf("calls %x; blocks %d", calls, conn.dfuBlocks)
	}
	got.APNonce[0] = 99
	*got.BoardID = 99
	again, err := conn.ReadInfo(context.Background())
	if err != nil || again.APNonce[0] != 1 || *again.BoardID != 0 {
		t.Fatal("aliased observation")
	}
}
func TestReadInfoReject(t *testing.T) {
	for _, tt := range []struct{ name, serial, nonce string }{
		{"ECID changed", "ECID:12 CPID:FE01", "NONC:01"},
		{"chip changed", "ECID:8000000000000011 CPID:12", "NONC:01"},
		{"bad board", "ECID:8000000000000011 CPID:FE01 BDID:zz", "NONC:01"},
		{"duplicate flags", "ECID:8000000000000011 CPID:FE01 CPFM:0 CPFM:0", "NONC:01"},
		{"wide flags", "ECID:8000000000000011 CPID:FE01 IBFL:100000000", "NONC:01"},
		{"duplicate nonce", "ECID:8000000000000011 CPID:FE01", "NONC:01 NONC:02"},
		{"empty nonce", "ECID:8000000000000011 CPID:FE01", "NONC:"},
		{"odd nonce", "ECID:8000000000000011 CPID:FE01", "NONC:123"},
		{"bad nonce", "ECID:8000000000000011 CPID:FE01", "SNON:zz"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var calls []uint16
			c := infoConnection(tt.serial, tt.nonce, &calls)
			if _, err := c.ReadInfo(context.Background()); err == nil {
				t.Fatal("accepted malformed observation")
			}
		})
	}
	var calls []uint16
	c := infoConnection("ECID:8000000000000011 CPID:FE01", "other:value", &calls)
	info, err := c.ReadInfo(context.Background())
	if err != nil || info.BoardID != nil || info.CPFM != nil || info.IBFL != nil || info.APNonce != nil || info.SEPNonce != nil {
		t.Fatalf("missing fields: %#v, %v", info, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls = nil
	if _, err := c.ReadInfo(ctx); err != context.Canceled || len(calls) != 0 {
		t.Fatal("ignored cancellation")
	}
}
func TestStringDescriptorReject(t *testing.T) {
	for _, b := range [][]byte{nil, {2, 3}, {4, 2, 9, 4}, {5, 3, 9, 4, 0}, {6, 3, 9, 4}} {
		c := &Conn{library: &library{control: func(_ uintptr, _ uint8, _ uint8, _ uint16, _ uint16, p *byte, size uint16, _ uint32) int32 {
			return int32(copy(unsafe.Slice(p, int(size)), b))
		}}, info: Device{serialIndex: 3}}
		if _, err := c.ReadInfo(context.Background()); err == nil {
			t.Fatalf("accepted %x", b)
		}
	}
}
func ExampleConn_ReadInfo() {
	var c Conn
	_, err := c.ReadInfo(context.Background())
	fmt.Println(err)
	// Output: recovery connection is closed
}
