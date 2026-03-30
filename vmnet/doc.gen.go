// Code generated from Apple documentation for vmnet. DO NOT EDIT.

// Package vmnet provides Go bindings for the vmnet framework.
//
// Connect with network interfaces to read and write packets on guest operating systems.
//
// The vmnet framework is an API for virtual machines to read and write packets.
//
// # Essentials
//
//   - com.apple.vm.networking: A Boolean that indicates whether the app manages virtual network interfaces without escalating privileges to the root user.
//
// # Starting and Stopping Interfaces
//
//   - vmnet_start_interface(_:_:_:): Starts host or shared mode on an interface with a specified configuration.
//   - vmnet_interface_set_event_callback(_:_:_:_:): Schedules a callback to be executed when events for the specified interface are received.
//   - vmnet_stop_interface(_:_:_:): Stops the interface.
//
// # Reading and Writing Packets
//
//   - vmnet_read(_:_:_:): Attempts to read a specified number of packets from an interface.
//   - vmnet_write(_:_:_:): Attempts to write specified packets to an interface.
//
// # Data Types
//
//   - vmnet_return_t: Values returned by functions in the vmnet Framework.
//   - vmpktdesc: Describes a packet.
//   - interface_ref: A virtual network interface.
//   - interface_event_t: Interface event types.
//   - operating_modes_t: The operating modes for an interface.
//
// # Constants
//
//   - interface_desc XPC Dictionary Keys: XPC dictionary keys supported by the  parameter passed to the  function to describe the parameters of the network interface.
//   - interface_param XPC Dictionary Keys: XPC dictionary keys used by the  argument returned by the completion handler of the  function that describes the parameters that should be used to configure the network interface.
//   - event XPC Dictionary: XPC dictionary keys used by the  value returned to the client in the  callback specified by the  function that provides information about the callback event.
//
// # Variables
//
//   - vmnet_enable_virtio_header_key
//   - vmnet_read_max_packets_key
//   - vmnet_write_max_packets_key
//
// # Functions
//
//   - vmnet_interface_start_with_network(_:_:_:_:)
//   - vmnet_network_configuration_add_dhcp_reservation(_:_:_:)
//   - vmnet_network_configuration_add_port_forwarding_rule(_:_:_:_:_:_:)
//   - vmnet_network_configuration_create(_:_:)
//   - vmnet_network_configuration_disable_dhcp(_:)
//   - vmnet_network_configuration_disable_dns_proxy(_:)
//   - vmnet_network_configuration_disable_nat44(_:)
//   - vmnet_network_configuration_disable_nat66(_:)
//   - vmnet_network_configuration_disable_router_advertisement(_:)
//   - vmnet_network_configuration_set_external_interface(_:_:)
//   - vmnet_network_configuration_set_ipv4_subnet(_:_:_:)
//   - vmnet_network_configuration_set_ipv6_prefix(_:_:_:)
//   - vmnet_network_configuration_set_mtu(_:_:)
//   - vmnet_network_copy_serialization(_:_:)
//   - vmnet_network_create(_:_:)
//   - vmnet_network_create_with_serialization(_:_:)
//   - vmnet_network_get_ipv4_subnet(_:_:_:)
//   - vmnet_network_get_ipv6_prefix(_:_:_:)
//
// # Type Aliases
//
//   - vmnet_mode_t
//   - vmnet_network_configuration_ref
//   - vmnet_network_ref
//
// [vmnet Documentation]: https://developer.apple.com/documentation/vmnet
package vmnet

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the vmnet library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/vmnet.framework/vmnet",
	"/usr/lib/libvmnet.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: vmnet: failed to load framework from any known path\n")
}
