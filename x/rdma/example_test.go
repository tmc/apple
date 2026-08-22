//go:build darwin

package rdma_test

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/rdma"
	xrdma "github.com/tmc/apple/x/rdma"
)

func ExampleSelectRouteGID() {
	gid0 := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	gid1 := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	route, ok := xrdma.SelectRouteGID([]xrdma.RouteGID{
		{Index: 0, GID: gid0},
		{Index: 1, GID: gid1},
	}, -1, xrdma.LinkLayerThunderbolt)

	if ok {
		fmt.Println("selected route index:", route.Index)
	}
	// Output:
	// selected route index: 1
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

	if ok {
		fmt.Println(reasons[0])
	}
	// Output:
	// read-only preflight passed; safe_to_attempt_rtr is necessary, not sufficient
}

func ExampleRequireRTRAttemptAllowed() {
	err := xrdma.RequireRTRAttemptAllowed(false)

	if errors.Is(err, xrdma.ErrRTRUnsafe) {
		fmt.Println(err)
	}
	// Output:
	// rdma-pingpong drives QP INIT->RTR, which can wedge Apple Thunderbolt RDMA ports; run rdmainfo preflight, run rdma-probe, and read the README first, then pass -allow-rtr for one bounded attempt
}

func ExampleRTRAttr() {
	remoteGID := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0xde, 0xad, 0, 0, 0, 0, 0, 1}

	attr, mask, err := xrdma.RTRAttr(
		xrdma.LocalQP{PortNum: 1, GIDIndex: 0, ActiveMTU: 5},
		xrdma.RemoteQP{LID: 2, QPN: 42, PSN: 7, GID: remoteGID, UseGlobal: true, ActiveMTU: 5},
		xrdma.RTRPolicy{},
	)

	fmt.Println(err)
	fmt.Printf("%#x %d %d %d\n", mask, attr.PathMTU, attr.AHAttr.DLID, attr.AHAttr.GRH.HopLimit)
	// Output:
	// <nil>
	// 0x101181 5 2 1
}
