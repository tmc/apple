package rdma_test

import (
	"errors"
	"fmt"

	xrdma "github.com/tmc/apple/x/rdma"
)

func ExampleSelectRouteGID() {
	gid0 := xrdma.GID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	gid1 := xrdma.GID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	route, ok := xrdma.SelectRouteGID([]xrdma.RouteGID{
		{Index: 0, GID: gid0},
		{Index: 1, GID: gid1},
	}, -1, xrdma.LinkLayerThunderbolt)

	fmt.Println(ok, route.Index)
	// Output: true 1
}

func ExampleDerivePreflightSafety() {
	routeIndex := 1
	ok, reasons := xrdma.DerivePreflightSafety(xrdma.PreflightReport{
		Devices: []xrdma.PreflightDevice{{
			Name:          "rdma_en1",
			State:         xrdma.PortActive,
			LinkLayer:     xrdma.LinkLayerThunderbolt,
			RouteGIDIndex: &routeIndex,
		}},
		IORegistry: map[string]int{
			xrdma.IORegistryPeerInterface:  1,
			xrdma.IORegistryXDomainService: 1,
		},
		RecentLog: []string{"AppleThunderboltRDMA context allocation/query/free"},
	})

	fmt.Println(ok)
	fmt.Println(reasons[0])
	// Output:
	// true
	// read-only preflight passed; safe_to_attempt_rtr is necessary, not sufficient
}

func ExampleRequireRTRAttemptAllowed() {
	err := xrdma.RequireRTRAttemptAllowed(false)

	fmt.Println(errors.Is(err, xrdma.ErrRTRUnsafe))
	fmt.Println(err)
	// Output:
	// true
	// rdma-pingpong drives QP INIT->RTR, which can wedge Apple Thunderbolt RDMA ports; run rdmainfo preflight, run rdma-probe, and read the README first, then pass -allow-rtr for one bounded attempt
}
