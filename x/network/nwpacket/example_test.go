//go:build darwin

package nwpacket_test

import (
	"context"
	"errors"
	"fmt"
	"net"

	applenetwork "github.com/tmc/apple/network"
	"github.com/tmc/apple/x/network/nwpacket"
)

func ExampleConfig() {
	cfg := nwpacket.Config{
		InterfaceName:         "awdl0",
		LocalAddr:             &net.UDPAddr{IP: net.ParseIP("fe80::1"), Zone: "awdl0"},
		RequiredInterfaceType: applenetwork.NWInterfaceTypeWifi,
		SetRequiredInterface:  true,
		IncludePeerToPeer:     true,
		RequireInterface:      true,
		ReuseLocalAddress:     true,
	}
	fmt.Println(cfg.InterfaceName, cfg.IncludePeerToPeer)
	// Output: awdl0 true
}

func ExampleListenPacketContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := nwpacket.ListenPacketContext(ctx, nwpacket.Config{
		LocalAddr: &net.UDPAddr{IP: net.ParseIP("127.0.0.1")},
	})
	fmt.Println(errors.Is(err, context.Canceled))
	// Output: true
}
