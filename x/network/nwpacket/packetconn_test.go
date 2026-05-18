//go:build darwin

package nwpacket

import (
	"context"
	"errors"
	"net"
	"testing"
)

func TestEndpointHost(t *testing.T) {
	tests := []struct {
		name string
		addr *net.UDPAddr
		want string
	}{
		{
			name: "ipv4",
			addr: &net.UDPAddr{IP: net.ParseIP("10.0.0.1"), Port: 1234},
			want: "10.0.0.1",
		},
		{
			name: "link local zone",
			addr: &net.UDPAddr{IP: net.ParseIP("fe80::1"), Zone: "awdl0", Port: 1234},
			want: "fe80::1%awdl0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nwEndpointHost(tt.addr); got != tt.want {
				t.Fatalf("nwEndpointHost(%v) = %q, want %q", tt.addr, got, tt.want)
			}
		})
	}
}

func TestNetAddrToUDP(t *testing.T) {
	addr := &net.UDPAddr{IP: net.ParseIP("fe80::1"), Port: 1234, Zone: "awdl0"}
	got, err := netAddrToUDP(addr)
	if err != nil {
		t.Fatal(err)
	}
	if got == addr {
		t.Fatal("netAddrToUDP returned original UDPAddr")
	}
	if got.String() != addr.String() {
		t.Fatalf("netAddrToUDP = %s, want %s", got, addr)
	}
	got.IP[0] ^= 0xff
	if got.IP.Equal(addr.IP) {
		t.Fatal("netAddrToUDP did not copy IP storage")
	}
}

func TestTimeoutError(t *testing.T) {
	err := nwTimeoutError{op: "read"}
	if !err.Timeout() {
		t.Fatal("Timeout returned false")
	}
	if !err.Temporary() {
		t.Fatal("Temporary returned false")
	}
	var netErr net.Error
	if !errors.As(err, &netErr) {
		t.Fatal("nwTimeoutError does not satisfy net.Error")
	}
}

func TestListenPacketRequiresLocalAddr(t *testing.T) {
	if _, err := ListenPacket(Config{}); err == nil {
		t.Fatal("ListenPacket succeeded without LocalAddr")
	}
}

func TestListenPacketContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := ListenPacketContext(ctx, Config{
		LocalAddr: &net.UDPAddr{IP: net.ParseIP("127.0.0.1")},
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("ListenPacketContext err = %v, want %v", err, context.Canceled)
	}
}
