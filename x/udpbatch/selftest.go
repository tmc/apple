package udpbatch

import (
	"bytes"
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

// selfTest exchanges two known datagrams from a connected sender to an
// unconnected receiver over loopback, using sendmsg_x on the way out and
// recvmsg_x on the way in, and verifies the kernel wrote back the expected
// byte counts, contents, and source address. It exercises exactly the
// contract [Send] and [Recv] rely on: zeroed name fields on send, populated
// msg_name and msg_datalen on receive.
func selfTest() error {
	recv, recvAddr, err := boundUDPSocket()
	if err != nil {
		return err
	}
	defer syscall.Close(recv)

	send, _, err := boundUDPSocket()
	if err != nil {
		return err
	}
	defer syscall.Close(send)
	if err := syscall.Connect(send, recvAddr); err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	// Give the receive side a deadline so a kernel that accepts the send but
	// never delivers cannot hang the check.
	tv := syscall.Timeval{Sec: 2}
	if err := syscall.SetsockoptTimeval(recv, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv); err != nil {
		return fmt.Errorf("setsockopt SO_RCVTIMEO: %w", err)
	}

	want := [][]byte{
		[]byte("udpbatch self-test datagram zero"),
		[]byte("udpbatch self-test datagram one, longer"),
	}

	// Send both datagrams in one sendmsg_x call.
	hdrs := make([]msghdrX, len(want))
	iovs := make([]syscall.Iovec, len(want))
	var pin runtime.Pinner
	defer pin.Unpin()
	for i, p := range want {
		pin.Pin(&p[0])
		iovs[i] = syscall.Iovec{Base: &p[0], Len: uint64(len(p))}
		hdrs[i] = msghdrX{Iov: &iovs[i], Iovlen: 1}
	}
	pin.Pin(&iovs[0])
	n, errno := call(sendmsgX, uintptr(send), hdrs)
	if errno != 0 {
		return fmt.Errorf("sendmsg_x: %w", errno)
	}
	if n != len(want) {
		return fmt.Errorf("sendmsg_x accepted %d of %d datagrams", n, len(want))
	}

	// Receive them in one recvmsg_x call and verify every output field this
	// package's API depends on.
	bufs := make([][]byte, len(want))
	rhdrs := make([]msghdrX, len(want))
	riovs := make([]syscall.Iovec, len(want))
	names := make([]syscall.RawSockaddrInet6, len(want))
	for i := range bufs {
		bufs[i] = make([]byte, 128)
		pin.Pin(&bufs[i][0])
		riovs[i] = syscall.Iovec{Base: &bufs[i][0], Len: uint64(len(bufs[i]))}
		rhdrs[i] = msghdrX{
			Name:    unsafe.Pointer(&names[i]),
			Namelen: uint32(unsafe.Sizeof(names[i])),
			Iov:     &riovs[i],
			Iovlen:  1,
		}
	}
	pin.Pin(&riovs[0])
	pin.Pin(&names[0])
	got, errno := call(recvmsgX, uintptr(recv), rhdrs)
	if errno != 0 {
		return fmt.Errorf("recvmsg_x: %w", errno)
	}
	if got != len(want) {
		return fmt.Errorf("recvmsg_x returned %d of %d datagrams", got, len(want))
	}
	for i, p := range want {
		if int(rhdrs[i].Datalen) != len(p) {
			return fmt.Errorf("datagram %d: msg_datalen = %d, want %d", i, rhdrs[i].Datalen, len(p))
		}
		if !bytes.Equal(bufs[i][:rhdrs[i].Datalen], p) {
			return fmt.Errorf("datagram %d: payload mismatch", i)
		}
		if rhdrs[i].Flags&syscall.MSG_TRUNC != 0 {
			return fmt.Errorf("datagram %d: unexpected MSG_TRUNC", i)
		}
		if ap := sockaddrToAddrPort(&names[i]); !ap.IsValid() || ap.Port() == 0 {
			return fmt.Errorf("datagram %d: msg_name not populated (got %v)", i, ap)
		}
	}
	return nil
}

// boundUDPSocket creates a blocking UDP socket bound to an ephemeral loopback
// port and reports its bound address.
func boundUDPSocket() (int, syscall.Sockaddr, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return 0, nil, fmt.Errorf("socket: %w", err)
	}
	sa := &syscall.SockaddrInet4{Addr: [4]byte{127, 0, 0, 1}}
	if err := syscall.Bind(fd, sa); err != nil {
		syscall.Close(fd)
		return 0, nil, fmt.Errorf("bind: %w", err)
	}
	bound, err := syscall.Getsockname(fd)
	if err != nil {
		syscall.Close(fd)
		return 0, nil, fmt.Errorf("getsockname: %w", err)
	}
	return fd, bound, nil
}
