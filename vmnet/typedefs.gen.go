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
type Vmnet_interface_completion_handler_t = func(Vmnet_return_t)

// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_event_callback_t
type Vmnet_interface_event_callback_t = func(Interface_event_t, objectivec.Object)

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
type Vmnet_start_interface_completion_handler_t = func(Vmnet_return_t, objectivec.Object)

// InterfaceRef is a Go-name alias for Interface_ref.
type InterfaceRef = Interface_ref

// VmnetInterfaceCompletionHandler is a Go-name alias for Vmnet_interface_completion_handler_t.
type VmnetInterfaceCompletionHandler = Vmnet_interface_completion_handler_t

// VmnetInterfaceEventCallback is a Go-name alias for Vmnet_interface_event_callback_t.
type VmnetInterfaceEventCallback = Vmnet_interface_event_callback_t

// VmnetInterfaceGetIPPortForwardingRulesHandler is a Go-name alias for Vmnet_interface_get_ip_port_forwarding_rules_handler_t.
type VmnetInterfaceGetIPPortForwardingRulesHandler = Vmnet_interface_get_ip_port_forwarding_rules_handler_t

// VmnetInterfaceGetPortForwardingRulesHandler is a Go-name alias for Vmnet_interface_get_port_forwarding_rules_handler_t.
type VmnetInterfaceGetPortForwardingRulesHandler = Vmnet_interface_get_port_forwarding_rules_handler_t

// VmnetMode is a Go-name alias for Vmnet_mode_t.
type VmnetMode = Vmnet_mode_t

// VmnetNetworkConfigurationRef is a Go-name alias for Vmnet_network_configuration_ref.
type VmnetNetworkConfigurationRef = Vmnet_network_configuration_ref

// VmnetNetworkRef is a Go-name alias for Vmnet_network_ref.
type VmnetNetworkRef = Vmnet_network_ref

// VmnetStartInterfaceCompletionHandler is a Go-name alias for Vmnet_start_interface_completion_handler_t.
type VmnetStartInterfaceCompletionHandler = Vmnet_start_interface_completion_handler_t
