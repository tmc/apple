//go:build darwin

package network_test

import (
	"fmt"

	"github.com/tmc/apple/network"
)

func ExampleNWEndpointType() {
	hostType := network.NWEndpointTypeHost
	fmt.Println(hostType)

	addrType := network.NWEndpointTypeAddress
	fmt.Println(addrType)

	// Output:
	// NWEndpointTypeHost
	// NWEndpointTypeAddress
}

func ExampleNWInterfaceType() {
	wifi := network.NWInterfaceTypeWifi
	fmt.Println(wifi)

	cellular := network.NWInterfaceTypeCellular
	fmt.Println(cellular)

	// Output:
	// NWInterfaceTypeWifi
	// NWInterfaceTypeCellular
}

func ExampleNWEndpointCreateHost() {
	endpoint := network.NWEndpointCreateHost("localhost", "8080")
	endpointType := network.NWEndpointGetType(endpoint)
	port := network.NWEndpointGetPort(endpoint)

	fmt.Println(endpointType)
	fmt.Println(port)

	// Output:
	// NWEndpointTypeHost
	// 8080
}
