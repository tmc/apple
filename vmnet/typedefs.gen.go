// Code generated from Apple documentation. DO NOT EDIT.

package vmnet

import (
	"github.com/tmc/apple/objectivec"
)

// Interface_ref is a virtual network interface.
//
// See: https://developer.apple.com/documentation/vmnet/interface_ref
type Interface_ref = uintptr

// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_completion_handler_t
type Vmnet_interface_completion_handler_t = func(objectivec.IObject)

// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_event_callback_t
type Vmnet_interface_event_callback_t = func(objectivec.IObject, objectivec.Object)

// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_ip_port_forwarding_rules_handler_t
type Vmnet_interface_get_ip_port_forwarding_rules_handler_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_port_forwarding_rules_handler_t
type Vmnet_interface_get_port_forwarding_rules_handler_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/vmnet/vmnet_mode_t
type Vmnet_mode_t = Operating_modes_t

// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_ref
type Vmnet_network_configuration_ref = uintptr

// See: https://developer.apple.com/documentation/vmnet/vmnet_network_ref
type Vmnet_network_ref = uintptr

// See: https://developer.apple.com/documentation/vmnet/vmnet_start_interface_completion_handler_t
type Vmnet_start_interface_completion_handler_t = func(objectivec.IObject, objectivec.Object)
