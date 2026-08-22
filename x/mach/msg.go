package mach

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// Header is mach_msg_header_t with its field names restored.
//
// The generated kernel.Mach_msg_header_t is an opaque [6]uint32: Apple's
// documentation describes the typedef and the field names live on the
// anonymous struct, so the generator cannot recover them. The layout is ABI
// pinned to xnu, not to the SDK documentation, so owning it here as a
// hand-written struct is correct, with the size assertions below tying it
// to the generated type so a change on either side breaks the build.
type Header struct {
	Bits        uint32
	Size        uint32
	RemotePort  Port
	LocalPort   Port
	VoucherPort Port
	ID          int32
}

// Compile-time: Header and kernel.Mach_msg_header_t must be the same size.
var (
	_ [unsafe.Sizeof(Header{}) - unsafe.Sizeof(kernel.Mach_msg_header_t{})]byte
	_ [unsafe.Sizeof(kernel.Mach_msg_header_t{}) - unsafe.Sizeof(Header{})]byte
)

const (
	msghBitsComplex = 0x80000000

	machSendMsg     = 0x00000001
	machRcvMsg      = 0x00000002
	machSendTimeout = 0x00000010
	machRcvTimeout  = 0x00000100

	headerSize   = 24 // mach_msg_header_t
	bodySize     = 4  // mach_msg_body_t: descriptor count
	portDescSize = 12 // mach_msg_port_descriptor_t
)

// PortRight names a port right to carry in a message.
type PortRight struct {
	Port        Port
	Disposition Disposition
}

// Message is a received Mach message: the header, any port rights carried
// in descriptors, and the inline body bytes.
//
// Received rights are owned by the caller and must be balanced (Deallocate
// for send rights) like any other acquired right.
//
// Mach message sizes are word-aligned, so Body carries up to three bytes
// of zero padding beyond what was sent; protocols that need exact lengths
// carry them in the body or the ID.
type Message struct {
	Header Header
	Ports  []Port
	Body   []byte
}

// Send sends a message to dest. destDisp says what right the send consumes
// to address dest (CopySend keeps the caller's send right, MoveSend
// consumes it, MakeSend requires dest be a receive right). Attached port
// rights cross by their own dispositions. timeout zero blocks.
//
// The kernel both reads the buffer and requires it not move; the buffer is
// built pointer-free and pinned for the call.
func Send(dest Port, destDisp Disposition, id int32, rights []PortRight, body []byte, timeout time.Duration) error {
	size := headerSize + len(body)
	if len(rights) > 0 {
		size += bodySize + portDescSize*len(rights)
	}
	size = (size + 3) &^ 3
	buf := make([]byte, size)

	h := (*Header)(unsafe.Pointer(&buf[0]))
	h.Bits = uint32(destDisp)
	h.Size = uint32(size)
	h.RemotePort = dest
	h.ID = id
	off := headerSize
	if len(rights) > 0 {
		h.Bits |= msghBitsComplex
		*(*uint32)(unsafe.Pointer(&buf[off])) = uint32(len(rights))
		off += bodySize
		for _, r := range rights {
			*(*uint32)(unsafe.Pointer(&buf[off])) = uint32(r.Port)
			// word 3: pad2(0..15) | disposition(16..23) | type(24..31);
			// MACH_MSG_PORT_DESCRIPTOR is type 0.
			*(*uint32)(unsafe.Pointer(&buf[off+8])) = uint32(r.Disposition) << 16
			off += portDescSize
		}
	}
	copy(buf[off:], body)

	option := int32(machSendMsg)
	var to uint32
	if timeout > 0 {
		option |= machSendTimeout
		to = uint32(timeout.Milliseconds())
	}
	var pin runtime.Pinner
	pin.Pin(&buf[0])
	kr := kernel.Mach_msg(unsafe.Pointer(&buf[0]), option, kernel.Mach_msg_size_t(size), 0, 0, to, 0)
	pin.Unpin()
	if kr != 0 {
		return fmt.Errorf("mach: mach_msg send: 0x%x", uint32(kr))
	}
	return nil
}

// Receive blocks on p's receive right until a message arrives or timeout
// elapses (zero blocks indefinitely). Port rights carried in descriptors
// are extracted into Message.Ports and owned by the caller.
func Receive(p Port, timeout time.Duration) (*Message, error) {
	const bufSize = 4096
	buf := make([]byte, bufSize)

	option := int32(machRcvMsg)
	var to uint32
	if timeout > 0 {
		option |= machRcvTimeout
		to = uint32(timeout.Milliseconds())
	}
	var pin runtime.Pinner
	pin.Pin(&buf[0])
	kr := kernel.Mach_msg(unsafe.Pointer(&buf[0]), option, 0, bufSize, kernel.Mach_port_name_t(p), to, 0)
	pin.Unpin()
	if kr != 0 {
		return nil, fmt.Errorf("mach: mach_msg receive: 0x%x", uint32(kr))
	}

	m := &Message{Header: *(*Header)(unsafe.Pointer(&buf[0]))}
	off := headerSize
	if m.Header.Bits&msghBitsComplex != 0 {
		n := *(*uint32)(unsafe.Pointer(&buf[off]))
		off += bodySize
		for range n {
			word3 := *(*uint32)(unsafe.Pointer(&buf[off+8]))
			if typ := word3 >> 24; typ != 0 {
				return nil, fmt.Errorf("mach: receive: descriptor type %d unsupported (only port descriptors)", typ)
			}
			m.Ports = append(m.Ports, Port(*(*uint32)(unsafe.Pointer(&buf[off]))))
			off += portDescSize
		}
	}
	if int(m.Header.Size) > off {
		m.Body = append([]byte(nil), buf[off:m.Header.Size]...)
	}
	return m, nil
}
