// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
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


var _vmnet_interface_add_ip_port_forwarding_rule func(interface_ uintptr, protocol_ uint8, external_port uint16, address_family uint8, internal_address unsafe.Pointer, internal_port uint16, handler uintptr) objectivec.IObject

// Vmnet_interface_add_ip_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_add_ip_port_forwarding_rule(_:_:_:_:_:_:_:)
func Vmnet_interface_add_ip_port_forwarding_rule(interface_ uintptr, protocol_ uint8, external_port uint16, address_family uint8, internal_address unsafe.Pointer, internal_port uint16, handler uintptr) objectivec.IObject {
	if _vmnet_interface_add_ip_port_forwarding_rule == nil {
		panic("vmnet: symbol vmnet_interface_add_ip_port_forwarding_rule not loaded")
	}
	return _vmnet_interface_add_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, internal_address, internal_port, handler)
}


var _vmnet_interface_add_port_forwarding_rule func(interface_ uintptr, protocol_ uint8, external_port uint16, internal_address uintptr, internal_port uint16, handler uintptr) objectivec.IObject

// Vmnet_interface_add_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_add_port_forwarding_rule(_:_:_:_:_:_:)
func Vmnet_interface_add_port_forwarding_rule(interface_ uintptr, protocol_ uint8, external_port uint16, internal_address uintptr, internal_port uint16, handler uintptr) objectivec.IObject {
	if _vmnet_interface_add_port_forwarding_rule == nil {
		panic("vmnet: symbol vmnet_interface_add_port_forwarding_rule not loaded")
	}
	return _vmnet_interface_add_port_forwarding_rule(interface_, protocol_, external_port, internal_address, internal_port, handler)
}


var _vmnet_interface_get_ip_port_forwarding_rules func(interface_ uintptr, address_family uintptr, handler uintptr) objectivec.IObject

// Vmnet_interface_get_ip_port_forwarding_rules.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_ip_port_forwarding_rules(_:_:_:)
func Vmnet_interface_get_ip_port_forwarding_rules(interface_ uintptr, address_family uintptr, handler uintptr) objectivec.IObject {
	if _vmnet_interface_get_ip_port_forwarding_rules == nil {
		panic("vmnet: symbol vmnet_interface_get_ip_port_forwarding_rules not loaded")
	}
	return _vmnet_interface_get_ip_port_forwarding_rules(interface_, address_family, handler)
}


var _vmnet_interface_get_port_forwarding_rules func(interface_ uintptr, handler uintptr) objectivec.IObject

// Vmnet_interface_get_port_forwarding_rules.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_port_forwarding_rules(_:_:)
func Vmnet_interface_get_port_forwarding_rules(interface_ uintptr, handler uintptr) objectivec.IObject {
	if _vmnet_interface_get_port_forwarding_rules == nil {
		panic("vmnet: symbol vmnet_interface_get_port_forwarding_rules not loaded")
	}
	return _vmnet_interface_get_port_forwarding_rules(interface_, handler)
}


var _vmnet_interface_remove_ip_port_forwarding_rule func(interface_ uintptr, protocol_ uint8, external_port uint16, address_family uint8, handler uintptr) objectivec.IObject

// Vmnet_interface_remove_ip_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_remove_ip_port_forwarding_rule(_:_:_:_:_:)
func Vmnet_interface_remove_ip_port_forwarding_rule(interface_ uintptr, protocol_ uint8, external_port uint16, address_family uint8, handler uintptr) objectivec.IObject {
	if _vmnet_interface_remove_ip_port_forwarding_rule == nil {
		panic("vmnet: symbol vmnet_interface_remove_ip_port_forwarding_rule not loaded")
	}
	return _vmnet_interface_remove_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, handler)
}


var _vmnet_interface_remove_port_forwarding_rule func(interface_ uintptr, protocol_ uint8, external_port uint16, handler uintptr) objectivec.IObject

// Vmnet_interface_remove_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_remove_port_forwarding_rule(_:_:_:_:)
func Vmnet_interface_remove_port_forwarding_rule(interface_ uintptr, protocol_ uint8, external_port uint16, handler uintptr) objectivec.IObject {
	if _vmnet_interface_remove_port_forwarding_rule == nil {
		panic("vmnet: symbol vmnet_interface_remove_port_forwarding_rule not loaded")
	}
	return _vmnet_interface_remove_port_forwarding_rule(interface_, protocol_, external_port, handler)
}


var _vmnet_interface_set_event_callback func(interface_ uintptr, event_mask uintptr, queue uintptr, callback uintptr) objectivec.IObject

// Vmnet_interface_set_event_callback schedules a callback to be executed when events for the specified interface are received.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_set_event_callback(_:_:_:_:)
func Vmnet_interface_set_event_callback(interface_ uintptr, event_mask uintptr, queue dispatch.Queue, callback uintptr) objectivec.IObject {
	if _vmnet_interface_set_event_callback == nil {
		panic("vmnet: symbol vmnet_interface_set_event_callback not loaded")
	}
	return _vmnet_interface_set_event_callback(interface_, event_mask, uintptr(queue.Handle()), callback)
}


var _vmnet_interface_start_with_network func(network uintptr, interface_desc unsafe.Pointer, queue uintptr, start_block uintptr) objectivec.IObject

// Vmnet_interface_start_with_network.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_start_with_network(_:_:_:_:)
func Vmnet_interface_start_with_network(network uintptr, interface_desc unsafe.Pointer, queue dispatch.Queue, start_block uintptr) objectivec.IObject {
	if _vmnet_interface_start_with_network == nil {
		panic("vmnet: symbol vmnet_interface_start_with_network not loaded")
	}
	return _vmnet_interface_start_with_network(network, interface_desc, uintptr(queue.Handle()), start_block)
}


var _vmnet_ip_port_forwarding_rule_get_details func(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, address_family uint8, internal_address unsafe.Pointer, internal_port *uint16) objectivec.IObject

// Vmnet_ip_port_forwarding_rule_get_details.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_ip_port_forwarding_rule_get_details(_:_:_:_:_:_:)
func Vmnet_ip_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, address_family uint8, internal_address unsafe.Pointer, internal_port *uint16) objectivec.IObject {
	if _vmnet_ip_port_forwarding_rule_get_details == nil {
		panic("vmnet: symbol vmnet_ip_port_forwarding_rule_get_details not loaded")
	}
	return _vmnet_ip_port_forwarding_rule_get_details(rule, protocol_, external_port, address_family, internal_address, internal_port)
}



var _vmnet_network_configuration_add_port_forwarding_rule func(config uintptr, protocol_ uint8, address_family uint8, internal_port uint16, external_port uint16, internal_address unsafe.Pointer) objectivec.IObject

// Vmnet_network_configuration_add_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_add_port_forwarding_rule(_:_:_:_:_:_:)
func Vmnet_network_configuration_add_port_forwarding_rule(config uintptr, protocol_ uint8, address_family uint8, internal_port uint16, external_port uint16, internal_address unsafe.Pointer) objectivec.IObject {
	if _vmnet_network_configuration_add_port_forwarding_rule == nil {
		panic("vmnet: symbol vmnet_network_configuration_add_port_forwarding_rule not loaded")
	}
	return _vmnet_network_configuration_add_port_forwarding_rule(config, protocol_, address_family, internal_port, external_port, internal_address)
}


var _vmnet_network_configuration_create func(mode uintptr, status uintptr) objectivec.IObject

// Vmnet_network_configuration_create.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_create(_:_:)
func Vmnet_network_configuration_create(mode uintptr, status uintptr) objectivec.IObject {
	if _vmnet_network_configuration_create == nil {
		panic("vmnet: symbol vmnet_network_configuration_create not loaded")
	}
	return _vmnet_network_configuration_create(mode, status)
}


var _vmnet_network_configuration_disable_dhcp func(config uintptr) 

// Vmnet_network_configuration_disable_dhcp.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dhcp(_:)
func Vmnet_network_configuration_disable_dhcp(config uintptr)  {
	if _vmnet_network_configuration_disable_dhcp == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_dhcp not loaded")
	}
	_vmnet_network_configuration_disable_dhcp(config)
}


var _vmnet_network_configuration_disable_dns_proxy func(config uintptr) 

// Vmnet_network_configuration_disable_dns_proxy.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dns_proxy(_:)
func Vmnet_network_configuration_disable_dns_proxy(config uintptr)  {
	if _vmnet_network_configuration_disable_dns_proxy == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_dns_proxy not loaded")
	}
	_vmnet_network_configuration_disable_dns_proxy(config)
}


var _vmnet_network_configuration_disable_nat44 func(config uintptr) 

// Vmnet_network_configuration_disable_nat44.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat44(_:)
func Vmnet_network_configuration_disable_nat44(config uintptr)  {
	if _vmnet_network_configuration_disable_nat44 == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_nat44 not loaded")
	}
	_vmnet_network_configuration_disable_nat44(config)
}


var _vmnet_network_configuration_disable_nat66 func(config uintptr) 

// Vmnet_network_configuration_disable_nat66.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat66(_:)
func Vmnet_network_configuration_disable_nat66(config uintptr)  {
	if _vmnet_network_configuration_disable_nat66 == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_nat66 not loaded")
	}
	_vmnet_network_configuration_disable_nat66(config)
}


var _vmnet_network_configuration_disable_router_advertisement func(config uintptr) 

// Vmnet_network_configuration_disable_router_advertisement.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_router_advertisement(_:)
func Vmnet_network_configuration_disable_router_advertisement(config uintptr)  {
	if _vmnet_network_configuration_disable_router_advertisement == nil {
		panic("vmnet: symbol vmnet_network_configuration_disable_router_advertisement not loaded")
	}
	_vmnet_network_configuration_disable_router_advertisement(config)
}


var _vmnet_network_configuration_set_external_interface func(config uintptr, interface_name *byte) objectivec.IObject

// Vmnet_network_configuration_set_external_interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_external_interface(_:_:)
func Vmnet_network_configuration_set_external_interface(config uintptr, interface_name *byte) objectivec.IObject {
	if _vmnet_network_configuration_set_external_interface == nil {
		panic("vmnet: symbol vmnet_network_configuration_set_external_interface not loaded")
	}
	return _vmnet_network_configuration_set_external_interface(config, interface_name)
}


var _vmnet_network_configuration_set_ipv4_subnet func(config uintptr, subnet_addr uintptr, subnet_mask uintptr) objectivec.IObject

// Vmnet_network_configuration_set_ipv4_subnet.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_ipv4_subnet(_:_:_:)
func Vmnet_network_configuration_set_ipv4_subnet(config uintptr, subnet_addr uintptr, subnet_mask uintptr) objectivec.IObject {
	if _vmnet_network_configuration_set_ipv4_subnet == nil {
		panic("vmnet: symbol vmnet_network_configuration_set_ipv4_subnet not loaded")
	}
	return _vmnet_network_configuration_set_ipv4_subnet(config, subnet_addr, subnet_mask)
}


var _vmnet_network_configuration_set_ipv6_prefix func(config uintptr, prefix uintptr, len_ uint8) objectivec.IObject

// Vmnet_network_configuration_set_ipv6_prefix.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_ipv6_prefix(_:_:_:)
func Vmnet_network_configuration_set_ipv6_prefix(config uintptr, prefix uintptr, len_ uint8) objectivec.IObject {
	if _vmnet_network_configuration_set_ipv6_prefix == nil {
		panic("vmnet: symbol vmnet_network_configuration_set_ipv6_prefix not loaded")
	}
	return _vmnet_network_configuration_set_ipv6_prefix(config, prefix, len_)
}


var _vmnet_network_configuration_set_mtu func(config uintptr, mtu uint32) objectivec.IObject

// Vmnet_network_configuration_set_mtu.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_mtu(_:_:)
func Vmnet_network_configuration_set_mtu(config uintptr, mtu uint32) objectivec.IObject {
	if _vmnet_network_configuration_set_mtu == nil {
		panic("vmnet: symbol vmnet_network_configuration_set_mtu not loaded")
	}
	return _vmnet_network_configuration_set_mtu(config, mtu)
}


var _vmnet_network_copy_serialization func(network uintptr, status uintptr) unsafe.Pointer

// Vmnet_network_copy_serialization.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_copy_serialization(_:_:)
func Vmnet_network_copy_serialization(network uintptr, status uintptr) unsafe.Pointer {
	if _vmnet_network_copy_serialization == nil {
		panic("vmnet: symbol vmnet_network_copy_serialization not loaded")
	}
	return _vmnet_network_copy_serialization(network, status)
}


var _vmnet_network_create func(configuration uintptr, status uintptr) objectivec.IObject

// Vmnet_network_create.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_create(_:_:)
func Vmnet_network_create(configuration uintptr, status uintptr) objectivec.IObject {
	if _vmnet_network_create == nil {
		panic("vmnet: symbol vmnet_network_create not loaded")
	}
	return _vmnet_network_create(configuration, status)
}


var _vmnet_network_create_with_serialization func(network unsafe.Pointer, status uintptr) objectivec.IObject

// Vmnet_network_create_with_serialization.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_create_with_serialization(_:_:)
func Vmnet_network_create_with_serialization(network unsafe.Pointer, status uintptr) objectivec.IObject {
	if _vmnet_network_create_with_serialization == nil {
		panic("vmnet: symbol vmnet_network_create_with_serialization not loaded")
	}
	return _vmnet_network_create_with_serialization(network, status)
}


var _vmnet_network_get_ipv4_subnet func(network uintptr, subnet uintptr, mask uintptr) 

// Vmnet_network_get_ipv4_subnet.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv4_subnet(_:_:_:)
func Vmnet_network_get_ipv4_subnet(network uintptr, subnet uintptr, mask uintptr)  {
	if _vmnet_network_get_ipv4_subnet == nil {
		panic("vmnet: symbol vmnet_network_get_ipv4_subnet not loaded")
	}
	_vmnet_network_get_ipv4_subnet(network, subnet, mask)
}


var _vmnet_network_get_ipv6_prefix func(network uintptr, prefix uintptr, prefix_len *uint8) 

// Vmnet_network_get_ipv6_prefix.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv6_prefix(_:_:_:)
func Vmnet_network_get_ipv6_prefix(network uintptr, prefix uintptr, prefix_len *uint8)  {
	if _vmnet_network_get_ipv6_prefix == nil {
		panic("vmnet: symbol vmnet_network_get_ipv6_prefix not loaded")
	}
	_vmnet_network_get_ipv6_prefix(network, prefix, prefix_len)
}


var _vmnet_port_forwarding_rule_get_details func(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, internal_address uintptr, internal_port *uint16) objectivec.IObject

// Vmnet_port_forwarding_rule_get_details.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_port_forwarding_rule_get_details(_:_:_:_:_:)
func Vmnet_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, internal_address uintptr, internal_port *uint16) objectivec.IObject {
	if _vmnet_port_forwarding_rule_get_details == nil {
		panic("vmnet: symbol vmnet_port_forwarding_rule_get_details not loaded")
	}
	return _vmnet_port_forwarding_rule_get_details(rule, protocol_, external_port, internal_address, internal_port)
}


var _vmnet_read func(interface_ uintptr, packets *Vmpktdesc, pktcnt *int) objectivec.IObject

// Vmnet_read attempts to read a specified number of packets from an interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_read(_:_:_:)
func Vmnet_read(interface_ uintptr, packets *Vmpktdesc, pktcnt []int) objectivec.IObject {
	if _vmnet_read == nil {
		panic("vmnet: symbol vmnet_read not loaded")
	}
	return _vmnet_read(interface_, packets, unsafe.SliceData(pktcnt))
}


var _vmnet_start_interface func(interface_desc unsafe.Pointer, queue uintptr, handler uintptr) objectivec.IObject

// Vmnet_start_interface starts host or shared mode on an interface with a specified configuration.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_start_interface(_:_:_:)
func Vmnet_start_interface(interface_desc unsafe.Pointer, queue dispatch.Queue, handler uintptr) objectivec.IObject {
	if _vmnet_start_interface == nil {
		panic("vmnet: symbol vmnet_start_interface not loaded")
	}
	return _vmnet_start_interface(interface_desc, uintptr(queue.Handle()), handler)
}


var _vmnet_stop_interface func(interface_ uintptr, queue uintptr, handler uintptr) objectivec.IObject

// Vmnet_stop_interface stops the interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_stop_interface(_:_:_:)
func Vmnet_stop_interface(interface_ uintptr, queue dispatch.Queue, handler uintptr) objectivec.IObject {
	if _vmnet_stop_interface == nil {
		panic("vmnet: symbol vmnet_stop_interface not loaded")
	}
	return _vmnet_stop_interface(interface_, uintptr(queue.Handle()), handler)
}


var _vmnet_write func(interface_ uintptr, packets *Vmpktdesc, pktcnt *int) objectivec.IObject

// Vmnet_write attempts to write specified packets to an interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_write(_:_:_:)
func Vmnet_write(interface_ uintptr, packets *Vmpktdesc, pktcnt []int) objectivec.IObject {
	if _vmnet_write == nil {
		panic("vmnet: symbol vmnet_write not loaded")
	}
	return _vmnet_write(interface_, packets, unsafe.SliceData(pktcnt))
}



func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_vmnet_copy_shared_interface_list, frameworkHandle, "vmnet_copy_shared_interface_list")
		registerFunc(&_vmnet_interface_add_ip_port_forwarding_rule, frameworkHandle, "vmnet_interface_add_ip_port_forwarding_rule")
		registerFunc(&_vmnet_interface_add_port_forwarding_rule, frameworkHandle, "vmnet_interface_add_port_forwarding_rule")
		registerFunc(&_vmnet_interface_get_ip_port_forwarding_rules, frameworkHandle, "vmnet_interface_get_ip_port_forwarding_rules")
		registerFunc(&_vmnet_interface_get_port_forwarding_rules, frameworkHandle, "vmnet_interface_get_port_forwarding_rules")
		registerFunc(&_vmnet_interface_remove_ip_port_forwarding_rule, frameworkHandle, "vmnet_interface_remove_ip_port_forwarding_rule")
		registerFunc(&_vmnet_interface_remove_port_forwarding_rule, frameworkHandle, "vmnet_interface_remove_port_forwarding_rule")
		registerFunc(&_vmnet_interface_set_event_callback, frameworkHandle, "vmnet_interface_set_event_callback")
		registerFunc(&_vmnet_interface_start_with_network, frameworkHandle, "vmnet_interface_start_with_network")
		registerFunc(&_vmnet_ip_port_forwarding_rule_get_details, frameworkHandle, "vmnet_ip_port_forwarding_rule_get_details")
		registerFunc(&_vmnet_network_configuration_add_port_forwarding_rule, frameworkHandle, "vmnet_network_configuration_add_port_forwarding_rule")
		registerFunc(&_vmnet_network_configuration_create, frameworkHandle, "vmnet_network_configuration_create")
		registerFunc(&_vmnet_network_configuration_disable_dhcp, frameworkHandle, "vmnet_network_configuration_disable_dhcp")
		registerFunc(&_vmnet_network_configuration_disable_dns_proxy, frameworkHandle, "vmnet_network_configuration_disable_dns_proxy")
		registerFunc(&_vmnet_network_configuration_disable_nat44, frameworkHandle, "vmnet_network_configuration_disable_nat44")
		registerFunc(&_vmnet_network_configuration_disable_nat66, frameworkHandle, "vmnet_network_configuration_disable_nat66")
		registerFunc(&_vmnet_network_configuration_disable_router_advertisement, frameworkHandle, "vmnet_network_configuration_disable_router_advertisement")
		registerFunc(&_vmnet_network_configuration_set_external_interface, frameworkHandle, "vmnet_network_configuration_set_external_interface")
		registerFunc(&_vmnet_network_configuration_set_ipv4_subnet, frameworkHandle, "vmnet_network_configuration_set_ipv4_subnet")
		registerFunc(&_vmnet_network_configuration_set_ipv6_prefix, frameworkHandle, "vmnet_network_configuration_set_ipv6_prefix")
		registerFunc(&_vmnet_network_configuration_set_mtu, frameworkHandle, "vmnet_network_configuration_set_mtu")
		registerFunc(&_vmnet_network_copy_serialization, frameworkHandle, "vmnet_network_copy_serialization")
		registerFunc(&_vmnet_network_create, frameworkHandle, "vmnet_network_create")
		registerFunc(&_vmnet_network_create_with_serialization, frameworkHandle, "vmnet_network_create_with_serialization")
		registerFunc(&_vmnet_network_get_ipv4_subnet, frameworkHandle, "vmnet_network_get_ipv4_subnet")
		registerFunc(&_vmnet_network_get_ipv6_prefix, frameworkHandle, "vmnet_network_get_ipv6_prefix")
		registerFunc(&_vmnet_port_forwarding_rule_get_details, frameworkHandle, "vmnet_port_forwarding_rule_get_details")
		registerFunc(&_vmnet_read, frameworkHandle, "vmnet_read")
		registerFunc(&_vmnet_start_interface, frameworkHandle, "vmnet_start_interface")
		registerFunc(&_vmnet_stop_interface, frameworkHandle, "vmnet_stop_interface")
		registerFunc(&_vmnet_write, frameworkHandle, "vmnet_write")
	}






