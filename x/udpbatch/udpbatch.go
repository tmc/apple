// Package udpbatch sends and receives multiple UDP datagrams per system call
// on Darwin, using the kernel's sendmsg_x and recvmsg_x entry points.
//
// Linux programs batch UDP with sendmmsg and recvmmsg. Darwin has the same
// capability but exposes it only through undocumented symbols declared in
// xnu's bsd/sys/socket_private.h, so Go UDP stacks on macOS issue one system
// call per datagram. This package binds those symbols.
//
// The symbols are private. [Available] reports whether they resolved and
// passed a self-test on this system, and every call returns [ErrUnavailable]
// if they did not. Callers must keep their per-datagram path and fall back to
// it; this package is an optimization, never a requirement.
//
// # Connections, not descriptors
//
// Every entry point takes a [syscall.RawConn] rather than a file descriptor,
// and performs its work inside the conn's Read or Write method. A bare int fd
// races Go's network poller and invites use-after-close: the descriptor may
// be closed and reused by an unrelated goroutine between the caller obtaining
// it and the kernel acting on it.
//
// Going through [syscall.RawConn] also fixes blocking. A *net.UDPConn's
// descriptor is always non-blocking, so a loop that called recvmsg_x on the
// raw fd would busy-spin on EAGAIN rather than sleep. RawConn.Read parks the
// goroutine on the poller and retries when the socket is readable.
//
// # Sending is restricted to connected sockets
//
// socket_private.h states that msg_name, msg_namelen, msg_control and
// msg_controllen must be zero on input to sendmsg_x. [Send] therefore takes
// payloads with no destinations and no control messages, and is usable only
// on a connected socket.
//
// This restriction is deliberately stricter than what some implementations do
// in practice. Whether the kernel honors the documented prohibition has not
// been established empirically, and the fast paths that ignore it are not
// known to run in production anywhere. Until that is measured, this package
// promises only what the header promises: if per-datagram addressing proves
// to work, an addressed variant can be added compatibly, whereas withdrawing
// one could not.
//
// Receiving carries no such restriction. recvmsg_x does not forbid msg_name
// on output, so [Recv] reports each datagram's source and suits an
// unconnected socket serving many peers.
//
// # Partial transfers are normal
//
// Like sendmmsg, the kernel may accept fewer datagrams than were offered.
// [Send] reports how many it took; the caller retries the remainder. A short
// count is ordinary operation, not an error.
//
// # Version behavior
//
// Known quirks by release, from implementations that shipped this surface
// and from this package's own tests:
//
//   - macOS 10.15: recvmsg_x does not write back msg_controllen (reported by
//     noq's implementation). This package does not expose control messages,
//     so that quirk is absorbed here.
//   - macOS 26.x (measured by this package's tests): truncation is silent.
//     The header promises MSG_TRUNC in msg_flags when a datagram exceeds the
//     buffer; the kernel instead clamps msg_datalen to the buffer size and
//     sets no flag. A Message whose N equals len(Payload) may have been
//     truncated — size Payload above the largest expected datagram, because
//     nothing distinguishes an exact fit from a loss.
//
// These are recorded because they demonstrate the class: private syscalls
// change behavior between releases without notice. The [Available] self-test
// exchanges real datagrams with the running kernel precisely because a
// layout check against a pinned header cannot detect that the header and the
// kernel disagree.
//
// # Distribution
//
// Resolving private symbols with dlsym is exactly what Mac App Store static
// analysis flags. Do not link this package into an App Store build. Like the
// repository's private/ bindings, this surface is unstable, may change or
// disappear between OS releases, and carries no compatibility guarantees.
package udpbatch

import (
	"errors"
	"fmt"
	"net/netip"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/ebitengine/purego"
)

// ErrUnavailable is returned by every call in this package when the batched
// datapath is not usable on the running system.
var ErrUnavailable = errors.New("udpbatch: sendmsg_x/recvmsg_x unavailable")

// BatchSize is the number of datagrams to offer per call that implementations
// have converged on. It is a starting point, not a tuned value: the point at
// which the benefit flattens has not been measured here, and the best size
// depends on payload size and whether the socket is on loopback or a real
// interface.
const BatchSize = 32

// A Message is one received datagram.
//
// Payload is supplied by the caller and is filled by [Recv]; N reports how
// many bytes were written into it. A datagram larger than the payload is
// truncated, and Flags reports the truncation via syscall.MSG_TRUNC.
type Message struct {
	// Payload is the caller-provided buffer to receive into.
	Payload []byte

	// N is the number of bytes written into Payload.
	N int

	// Addr is the datagram's source address.
	Addr netip.AddrPort

	// Flags holds the kernel's per-message flags, including truncation.
	Flags int
}

// msghdrX corresponds to struct msghdr_x.
//
// Layout proven against internal/xnu/socket_private_excerpt.h (xnu revision
// 8d741a5de7ff4191bf97d57b9f54c2f6d4a15585) by TestMsghdrXLayout, which
// compiles the vendored header and compares offsets field by field.
type msghdrX struct {
	Name       unsafe.Pointer // msg_name
	Namelen    uint32         // msg_namelen
	_          [4]byte
	Iov        *syscall.Iovec // msg_iov
	Iovlen     int32          // msg_iovlen
	_          [4]byte
	Control    unsafe.Pointer // msg_control
	Controllen uint32         // msg_controllen
	Flags      int32          // msg_flags
	Datalen    uint64         // msg_datalen (size_t)
}

var (
	loadOnce  sync.Once
	available bool

	recvmsgX func(s int32, msgp *msghdrX, cnt uint32, flags int32) int64
	sendmsgX func(s int32, msgp *msghdrX, cnt uint32, flags int32) int64
	errnoLoc func() *int32
)

// Available reports whether the batched datapath is usable on this system.
//
// On first use it resolves the underlying symbols and runs a self-test: it
// exchanges known datagrams over a loopback socket pair and verifies that the
// kernel wrote back what was expected, including per-datagram lengths and the
// sender's address. The self-test is the point. The committed layout record
// checks this package's struct against a pinned copy of Apple's header, but
// only an exchange with the running kernel can detect that the header and the
// kernel disagree, and only that check fails safe on a system this package
// was never tested against.
//
// A false result is permanent for the life of the process and means every
// call in this package returns [ErrUnavailable].
func Available() bool {
	loadOnce.Do(func() { available = load() && selfTest() == nil })
	return available
}

// SelfTestError runs the availability check and, if it failed, reports why.
// It exists for diagnostics; callers should gate on [Available].
func SelfTestError() error {
	if Available() {
		return nil
	}
	if !load() {
		return fmt.Errorf("%w: symbols did not resolve", ErrUnavailable)
	}
	return fmt.Errorf("%w: self-test: %v", ErrUnavailable, selfTest())
}

func load() bool {
	if recvmsgX != nil && sendmsgX != nil && errnoLoc != nil {
		return true
	}
	r, err := purego.Dlsym(purego.RTLD_DEFAULT, "recvmsg_x")
	if err != nil || r == 0 {
		return false
	}
	s, err := purego.Dlsym(purego.RTLD_DEFAULT, "sendmsg_x")
	if err != nil || s == 0 {
		return false
	}
	e, err := purego.Dlsym(purego.RTLD_DEFAULT, "__error")
	if err != nil || e == 0 {
		return false
	}
	purego.RegisterFunc(&recvmsgX, r)
	purego.RegisterFunc(&sendmsgX, s)
	purego.RegisterFunc(&errnoLoc, e)
	return true
}

// call invokes fn on a locked OS thread so that the errno read that follows a
// failure observes the same thread's errno slot.
func call(fn func(s int32, msgp *msghdrX, cnt uint32, flags int32) int64, fd uintptr, hdrs []msghdrX) (int, syscall.Errno) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	n := fn(int32(fd), &hdrs[0], uint32(len(hdrs)), 0)
	if n < 0 {
		return 0, syscall.Errno(*errnoLoc())
	}
	return int(n), 0
}

// Recv receives up to len(msgs) datagrams in one system call and reports how
// many were received, filling the N, Addr and Flags fields of those entries.
//
// It parks the calling goroutine on the network poller until the socket is
// readable, so it neither busy-spins nor blocks an OS thread.
//
// Every payload buffer and the internal header arrays are pinned with a
// [runtime.Pinner] for the duration of the call: the kernel is handed
// pointers into Go memory, which the garbage collector must not move
// underneath it.
func Recv(conn syscall.RawConn, msgs []Message) (n int, err error) {
	if !Available() {
		return 0, ErrUnavailable
	}
	if len(msgs) == 0 {
		return 0, nil
	}

	hdrs := make([]msghdrX, len(msgs))
	iovs := make([]syscall.Iovec, len(msgs))
	names := make([]syscall.RawSockaddrInet6, len(msgs))

	var pin runtime.Pinner
	defer pin.Unpin()
	pin.Pin(&iovs[0])
	pin.Pin(&names[0])
	for i := range msgs {
		if len(msgs[i].Payload) == 0 {
			return 0, fmt.Errorf("udpbatch: message %d has empty payload buffer", i)
		}
		pin.Pin(&msgs[i].Payload[0])
		iovs[i] = syscall.Iovec{Base: &msgs[i].Payload[0], Len: uint64(len(msgs[i].Payload))}
		hdrs[i] = msghdrX{
			Name:    unsafe.Pointer(&names[i]),
			Namelen: uint32(unsafe.Sizeof(names[i])),
			Iov:     &iovs[i],
			Iovlen:  1,
		}
	}

	var errno syscall.Errno
	rcErr := conn.Read(func(fd uintptr) bool {
		n, errno = call(recvmsgX, fd, hdrs)
		return errno != syscall.EAGAIN
	})
	if rcErr != nil {
		return 0, rcErr
	}
	if errno != 0 {
		return 0, fmt.Errorf("udpbatch: recvmsg_x: %w", errno)
	}

	for i := 0; i < n; i++ {
		msgs[i].N = int(hdrs[i].Datalen)
		msgs[i].Flags = int(hdrs[i].Flags)
		msgs[i].Addr = sockaddrToAddrPort(&names[i])
	}
	return n, nil
}

// Send transmits payloads on a connected socket in one system call and
// reports how many datagrams the kernel accepted, which may be fewer than
// len(payloads).
//
// The caller retries the unaccepted tail. Destinations and control messages
// are not supported; see the package documentation for why.
//
// Payloads are pinned for the duration of the call, as described on [Recv].
func Send(conn syscall.RawConn, payloads [][]byte) (n int, err error) {
	if !Available() {
		return 0, ErrUnavailable
	}
	if len(payloads) == 0 {
		return 0, nil
	}

	hdrs := make([]msghdrX, len(payloads))
	iovs := make([]syscall.Iovec, len(payloads))

	var pin runtime.Pinner
	defer pin.Unpin()
	pin.Pin(&iovs[0])
	for i, p := range payloads {
		if len(p) == 0 {
			return 0, fmt.Errorf("udpbatch: payload %d is empty", i)
		}
		pin.Pin(&p[0])
		iovs[i] = syscall.Iovec{Base: &p[0], Len: uint64(len(p))}
		// All other fields zero, as the header requires on input.
		hdrs[i] = msghdrX{Iov: &iovs[i], Iovlen: 1}
	}

	var errno syscall.Errno
	wcErr := conn.Write(func(fd uintptr) bool {
		n, errno = call(sendmsgX, fd, hdrs)
		return errno != syscall.EAGAIN
	})
	if wcErr != nil {
		return 0, wcErr
	}
	if errno != 0 {
		return 0, fmt.Errorf("udpbatch: sendmsg_x: %w", errno)
	}
	return n, nil
}

func sockaddrToAddrPort(sa *syscall.RawSockaddrInet6) netip.AddrPort {
	switch sa.Family {
	case syscall.AF_INET:
		sa4 := (*syscall.RawSockaddrInet4)(unsafe.Pointer(sa))
		port := uint16(sa4.Port>>8) | uint16(sa4.Port&0xff)<<8
		return netip.AddrPortFrom(netip.AddrFrom4(sa4.Addr), port)
	case syscall.AF_INET6:
		port := uint16(sa.Port>>8) | uint16(sa.Port&0xff)<<8
		addr := netip.AddrFrom16(sa.Addr)
		if sa.Scope_id != 0 {
			addr = addr.WithZone(fmt.Sprint(sa.Scope_id))
		}
		return netip.AddrPortFrom(addr, port)
	}
	return netip.AddrPort{}
}
