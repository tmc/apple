package udpbatch

import (
	"runtime"
	"syscall"
	"testing"
	"unsafe"
)

// TestSendmsgXWithMsgName probes the open question in design/udp-datapath.md:
// socket_private.h says msg_name must be zero on input to sendmsg_x, but noq
// sets it anyway (dormant in production). This test asks the running kernel.
// It asserts nothing about which answer is correct — it records the answer.
func TestSendmsgXWithMsgName(t *testing.T) {
	if !Available() {
		t.Skip("unavailable")
	}
	recv, recvAddr, err := boundUDPSocket()
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(recv)
	send, _, err := boundUDPSocket() // NOT connected
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(send)

	dst := recvAddr.(*syscall.SockaddrInet4)
	var raw syscall.RawSockaddrInet4
	raw.Len = uint8(unsafe.Sizeof(raw))
	raw.Family = syscall.AF_INET
	raw.Addr = dst.Addr
	raw.Port = uint16(dst.Port>>8) | uint16(dst.Port&0xff)<<8

	payload := []byte("addressed datagram")
	var pin runtime.Pinner
	defer pin.Unpin()
	pin.Pin(&payload[0])
	pin.Pin(&raw)
	iov := syscall.Iovec{Base: &payload[0], Len: uint64(len(payload))}
	pin.Pin(&iov)
	hdrs := []msghdrX{{
		Name:    unsafe.Pointer(&raw),
		Namelen: uint32(unsafe.Sizeof(raw)),
		Iov:     &iov,
		Iovlen:  1,
	}}
	n, errno := call(sendmsgX, uintptr(send), hdrs)
	if errno != 0 {
		t.Logf("PROBE ANSWER: kernel REJECTS msg_name on sendmsg_x: %v (header prohibition is enforced)", errno)
		return
	}
	t.Logf("sendmsg_x accepted %d addressed datagram(s); verifying delivery", n)

	tv := syscall.Timeval{Sec: 2}
	syscall.SetsockoptTimeval(recv, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)
	buf := make([]byte, 64)
	rn, _, err := syscall.Recvfrom(recv, buf, 0)
	if err != nil {
		t.Logf("PROBE ANSWER: kernel ACCEPTED msg_name but did not deliver (recv: %v) — silent drop, worst case", err)
		return
	}
	t.Logf("PROBE ANSWER: kernel ACCEPTS AND DELIVERS msg_name on sendmsg_x (%d bytes: %q) — header comment is stale; addressed variant is viable", rn, buf[:rn])
}
