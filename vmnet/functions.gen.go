// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: vmnet: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _vmnet_copy_shared_interface_list func() unsafe.Pointer

// Vmnet_copy_shared_interface_list.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_copy_shared_interface_list()
func Vmnet_copy_shared_interface_list() unsafe.Pointer {
	if _vmnet_copy_shared_interface_list == nil {
		panic("vmnet: symbol vmnet_copy_shared_interface_list not loaded")
	}
	return _vmnet_copy_shared_interface_list()
}

var _vmnet_interface_start_with_network func(network Vmnet_network_ref, interface_desc unsafe.Pointer, queue uintptr, start_block Vmnet_start_interface_completion_handler_t) Interface_ref

// Vmnet_interface_start_with_network.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_start_with_network(_:_:_:_:)
func Vmnet_interface_start_with_network(network Vmnet_network_ref, interface_desc unsafe.Pointer, queue dispatch.Queue, start_block Vmnet_start_interface_completion_handler_t) Interface_ref {
	if _vmnet_interface_start_with_network == nil {
		panic("vmnet: symbol vmnet_interface_start_with_network not loaded")
	}
	return _vmnet_interface_start_with_network(network, interface_desc, uintptr(queue.Handle()), start_block)
}

var _vmnet_network_configuration_disable_dhcp func(config Vmnet_network_configuration_ref)

// Vmnet_network_configuration_disable_dhcp.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dhcp(_:)
func Vmnet_network_configuration_disable_dhcp(config Vmnet_network_configuration_ref) {
	if _vmnet_network_configuration_disable_dhcp == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_dhcp not loaded")
	}
	_vmnet_network_configuration_disable_dhcp(config)
}

var _vmnet_network_configuration_disable_dns_proxy func(config Vmnet_network_configuration_ref)

// Vmnet_network_configuration_disable_dns_proxy.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dns_proxy(_:)
func Vmnet_network_configuration_disable_dns_proxy(config Vmnet_network_configuration_ref) {
	if _vmnet_network_configuration_disable_dns_proxy == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_dns_proxy not loaded")
	}
	_vmnet_network_configuration_disable_dns_proxy(config)
}

var _vmnet_network_configuration_disable_nat44 func(config Vmnet_network_configuration_ref)

// Vmnet_network_configuration_disable_nat44.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat44(_:)
func Vmnet_network_configuration_disable_nat44(config Vmnet_network_configuration_ref) {
	if _vmnet_network_configuration_disable_nat44 == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_nat44 not loaded")
	}
	_vmnet_network_configuration_disable_nat44(config)
}

var _vmnet_network_configuration_disable_nat66 func(config Vmnet_network_configuration_ref)

// Vmnet_network_configuration_disable_nat66.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat66(_:)
func Vmnet_network_configuration_disable_nat66(config Vmnet_network_configuration_ref) {
	if _vmnet_network_configuration_disable_nat66 == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_nat66 not loaded")
	}
	_vmnet_network_configuration_disable_nat66(config)
}

var _vmnet_network_configuration_disable_router_advertisement func(config Vmnet_network_configuration_ref)

// Vmnet_network_configuration_disable_router_advertisement.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_router_advertisement(_:)
func Vmnet_network_configuration_disable_router_advertisement(config Vmnet_network_configuration_ref) {
	if _vmnet_network_configuration_disable_router_advertisement == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_router_advertisement not loaded")
	}
	_vmnet_network_configuration_disable_router_advertisement(config)
}

var _vmnet_network_get_ipv4_subnet func(network Vmnet_network_ref, subnet uintptr, mask uintptr)

// Vmnet_network_get_ipv4_subnet.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv4_subnet(_:_:_:)
func Vmnet_network_get_ipv4_subnet(network Vmnet_network_ref, subnet uintptr, mask uintptr) {
	if _vmnet_network_get_ipv4_subnet == nil {
		panic("vmnet: symbol vmnet_network_get_ipv4_subnet not loaded")
	}
	_vmnet_network_get_ipv4_subnet(network, subnet, mask)
}

var _vmnet_network_get_ipv6_prefix func(network Vmnet_network_ref, prefix uintptr, prefix_len *uint8)

// Vmnet_network_get_ipv6_prefix.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv6_prefix(_:_:_:)
func Vmnet_network_get_ipv6_prefix(network Vmnet_network_ref, prefix uintptr, prefix_len *uint8) {
	if _vmnet_network_get_ipv6_prefix == nil {
		panic("vmnet: symbol vmnet_network_get_ipv6_prefix not loaded")
	}
	_vmnet_network_get_ipv6_prefix(network, prefix, prefix_len)
}

var _vmnet_start_interface func(interface_desc unsafe.Pointer, queue uintptr, handler Vmnet_start_interface_completion_handler_t) Interface_ref

// Vmnet_start_interface starts host or shared mode on an interface with a specified configuration.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_start_interface(_:_:_:)
func Vmnet_start_interface(interface_desc unsafe.Pointer, queue dispatch.Queue, handler Vmnet_start_interface_completion_handler_t) Interface_ref {
	if _vmnet_start_interface == nil {
		panic("vmnet: symbol vmnet_start_interface not loaded")
	}
	return _vmnet_start_interface(interface_desc, uintptr(queue.Handle()), handler)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_vmnet_copy_shared_interface_list, frameworkHandle, "vmnet_copy_shared_interface_list")
	registerFunc(&_vmnet_interface_start_with_network, frameworkHandle, "vmnet_interface_start_with_network")
	registerFunc(&_vmnet_network_configuration_disable_dhcp, frameworkHandle, "vmnet_network_configuration_disable_dhcp")
	registerFunc(&_vmnet_network_configuration_disable_dns_proxy, frameworkHandle, "vmnet_network_configuration_disable_dns_proxy")
	registerFunc(&_vmnet_network_configuration_disable_nat44, frameworkHandle, "vmnet_network_configuration_disable_nat44")
	registerFunc(&_vmnet_network_configuration_disable_nat66, frameworkHandle, "vmnet_network_configuration_disable_nat66")
	registerFunc(&_vmnet_network_configuration_disable_router_advertisement, frameworkHandle, "vmnet_network_configuration_disable_router_advertisement")
	registerFunc(&_vmnet_network_get_ipv4_subnet, frameworkHandle, "vmnet_network_get_ipv4_subnet")
	registerFunc(&_vmnet_network_get_ipv6_prefix, frameworkHandle, "vmnet_network_get_ipv6_prefix")
	registerFunc(&_vmnet_start_interface, frameworkHandle, "vmnet_start_interface")
}
