package udpbatch

import (
	"fmt"
	"net"
	"syscall"
	"testing"
)

func TestAvailable(t *testing.T) {
	if !Available() {
		t.Fatalf("Available() = false on this system: %v", SelfTestError())
	}
}

func TestSendRecvRoundTrip(t *testing.T) {
	if !Available() {
		t.Skip("batched datapath unavailable:", SelfTestError())
	}

	// Unconnected receiver: Recv must report source addresses.
	rc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	// Connected sender, as Send requires.
	sc, err := net.DialUDP("udp4", nil, rc.LocalAddr().(*net.UDPAddr))
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Close()

	const count = 17
	payloads := make([][]byte, count)
	for i := range payloads {
		payloads[i] = []byte(fmt.Sprintf("datagram %02d payload", i))
	}

	sraw, err := sc.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	sent := 0
	for sent < count {
		n, err := Send(sraw, payloads[sent:])
		if err != nil {
			t.Fatalf("Send after %d: %v", sent, err)
		}
		if n == 0 {
			t.Fatal("Send accepted 0 datagrams")
		}
		sent += n
	}

	rraw, err := rc.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	msgs := make([]Message, count)
	for i := range msgs {
		msgs[i].Payload = make([]byte, 64)
	}
	got := 0
	for got < count {
		n, err := Recv(rraw, msgs[got:])
		if err != nil {
			t.Fatalf("Recv after %d: %v", got, err)
		}
		got += n
	}

	senderPort := sc.LocalAddr().(*net.UDPAddr).Port
	for i := 0; i < count; i++ {
		want := fmt.Sprintf("datagram %02d payload", i)
		if string(msgs[i].Payload[:msgs[i].N]) != want {
			t.Errorf("msg %d: got %q, want %q", i, msgs[i].Payload[:msgs[i].N], want)
		}
		if int(msgs[i].Addr.Port()) != senderPort {
			t.Errorf("msg %d: source port %d, want %d", i, msgs[i].Addr.Port(), senderPort)
		}
		if !msgs[i].Addr.Addr().IsLoopback() {
			t.Errorf("msg %d: source addr %v, want loopback", i, msgs[i].Addr.Addr())
		}
	}
}

func TestRecvTruncation(t *testing.T) {
	if !Available() {
		t.Skip("batched datapath unavailable:", SelfTestError())
	}
	rc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()
	sc, err := net.DialUDP("udp4", nil, rc.LocalAddr().(*net.UDPAddr))
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Close()

	if _, err := sc.Write(make([]byte, 100)); err != nil {
		t.Fatal(err)
	}

	rraw, err := rc.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	msgs := []Message{{Payload: make([]byte, 10)}}
	n, err := Recv(rraw, msgs)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("Recv = %d, want 1", n)
	}

	// The vendored header promises MSG_TRUNC in msg_flags on truncation.
	// Observed on macOS 26.x: the flag is NOT set and msg_datalen is clamped
	// to the buffer size, making truncation silent. This test documents the
	// divergence: it accepts either behavior but fails if the kernel starts
	// reporting something else entirely, so a behavior change is noticed.
	truncFlagged := msgs[0].Flags&syscall.MSG_TRUNC != 0
	datalenClamped := msgs[0].N == len(msgs[0].Payload)
	if !truncFlagged && !datalenClamped {
		t.Errorf("100-byte datagram into 10-byte buffer: N=%d flags=%#x — neither MSG_TRUNC nor clamped datalen", msgs[0].N, msgs[0].Flags)
	}
	if truncFlagged {
		t.Log("kernel reports MSG_TRUNC (header behavior)")
	} else {
		t.Log("kernel clamps msg_datalen silently, no MSG_TRUNC (macOS 26.x behavior)")
	}
}
