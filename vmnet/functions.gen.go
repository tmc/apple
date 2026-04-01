// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("vmnet: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("vmnet: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("vmnet: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("vmnet: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _vmnet_copy_shared_interface_list func() unsafe.Pointer
var _vmnet_copy_shared_interface_listErr error

func tryVmnet_copy_shared_interface_list() (unsafe.Pointer, error) {
	if _vmnet_copy_shared_interface_list == nil {
		return nil, symbolCallError("vmnet_copy_shared_interface_list", "10.15", _vmnet_copy_shared_interface_listErr)
	}
	return _vmnet_copy_shared_interface_list(), nil
}

// Vmnet_copy_shared_interface_list.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_copy_shared_interface_list()
func Vmnet_copy_shared_interface_list() unsafe.Pointer {
	result, callErr := tryVmnet_copy_shared_interface_list()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_add_ip_port_forwarding_rule func(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, internal_address unsafe.Pointer, internal_port uint16, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_add_ip_port_forwarding_ruleErr error

func tryVmnet_interface_add_ip_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, internal_address unsafe.Pointer, internal_port uint16, handler Vmnet_interface_completion_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_add_ip_port_forwarding_rule == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_add_ip_port_forwarding_rule", "11.0", _vmnet_interface_add_ip_port_forwarding_ruleErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_add_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, internal_address, internal_port, _block0), nil
}

// Vmnet_interface_add_ip_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_add_ip_port_forwarding_rule(_:_:_:_:_:_:_:)
func Vmnet_interface_add_ip_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, internal_address unsafe.Pointer, internal_port uint16, handler Vmnet_interface_completion_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_add_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, internal_address, internal_port, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_add_port_forwarding_rule func(interface_ Interface_ref, protocol_ uint8, external_port uint16, internal_address uintptr, internal_port uint16, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_add_port_forwarding_ruleErr error

func tryVmnet_interface_add_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, internal_address uintptr, internal_port uint16, handler Vmnet_interface_completion_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_add_port_forwarding_rule == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_add_port_forwarding_rule", "10.15", _vmnet_interface_add_port_forwarding_ruleErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_add_port_forwarding_rule(interface_, protocol_, external_port, internal_address, internal_port, _block0), nil
}

// Vmnet_interface_add_port_forwarding_rule.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_add_port_forwarding_rule(_:_:_:_:_:_:)
func Vmnet_interface_add_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, internal_address uintptr, internal_port uint16, handler Vmnet_interface_completion_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_add_port_forwarding_rule(interface_, protocol_, external_port, internal_address, internal_port, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_get_ip_port_forwarding_rules func(interface_ Interface_ref, address_family kernel.U_int8_t, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_get_ip_port_forwarding_rulesErr error

func tryVmnet_interface_get_ip_port_forwarding_rules(interface_ Interface_ref, address_family kernel.U_int8_t, handler Vmnet_interface_get_ip_port_forwarding_rules_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_get_ip_port_forwarding_rules == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_get_ip_port_forwarding_rules", "11.0", _vmnet_interface_get_ip_port_forwarding_rulesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_get_ip_port_forwarding_rules(interface_, address_family, _block0), nil
}

// Vmnet_interface_get_ip_port_forwarding_rules.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_ip_port_forwarding_rules(_:_:_:)
func Vmnet_interface_get_ip_port_forwarding_rules(interface_ Interface_ref, address_family kernel.U_int8_t, handler Vmnet_interface_get_ip_port_forwarding_rules_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_get_ip_port_forwarding_rules(interface_, address_family, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_get_port_forwarding_rules func(interface_ Interface_ref, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_get_port_forwarding_rulesErr error

func tryVmnet_interface_get_port_forwarding_rules(interface_ Interface_ref, handler Vmnet_interface_get_port_forwarding_rules_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_get_port_forwarding_rules == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_get_port_forwarding_rules", "10.15", _vmnet_interface_get_port_forwarding_rulesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_get_port_forwarding_rules(interface_, _block0), nil
}

// Vmnet_interface_get_port_forwarding_rules.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_get_port_forwarding_rules(_:_:)
func Vmnet_interface_get_port_forwarding_rules(interface_ Interface_ref, handler Vmnet_interface_get_port_forwarding_rules_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_get_port_forwarding_rules(interface_, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_remove_ip_port_forwarding_rule func(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_remove_ip_port_forwarding_ruleErr error

func tryVmnet_interface_remove_ip_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, handler Vmnet_interface_completion_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_remove_ip_port_forwarding_rule == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_remove_ip_port_forwarding_rule", "11.0", _vmnet_interface_remove_ip_port_forwarding_ruleErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_remove_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, _block0), nil
}

// Vmnet_interface_remove_ip_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_remove_ip_port_forwarding_rule(_:_:_:_:_:)
func Vmnet_interface_remove_ip_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, address_family uint8, handler Vmnet_interface_completion_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_remove_ip_port_forwarding_rule(interface_, protocol_, external_port, address_family, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_remove_port_forwarding_rule func(interface_ Interface_ref, protocol_ uint8, external_port uint16, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_remove_port_forwarding_ruleErr error

func tryVmnet_interface_remove_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, handler Vmnet_interface_completion_handler_t) (Vmnet_return_t, error) {
	if _vmnet_interface_remove_port_forwarding_rule == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_remove_port_forwarding_rule", "10.15", _vmnet_interface_remove_port_forwarding_ruleErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_remove_port_forwarding_rule(interface_, protocol_, external_port, _block0), nil
}

// Vmnet_interface_remove_port_forwarding_rule.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_remove_port_forwarding_rule(_:_:_:_:)
func Vmnet_interface_remove_port_forwarding_rule(interface_ Interface_ref, protocol_ uint8, external_port uint16, handler Vmnet_interface_completion_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_remove_port_forwarding_rule(interface_, protocol_, external_port, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_set_event_callback func(interface_ Interface_ref, event_mask Interface_event_t, queue uintptr, callback unsafe.Pointer) Vmnet_return_t
var _vmnet_interface_set_event_callbackErr error

func tryVmnet_interface_set_event_callback(interface_ Interface_ref, event_mask Interface_event_t, queue dispatch.Queue, callback Vmnet_interface_event_callback_t) (Vmnet_return_t, error) {
	if _vmnet_interface_set_event_callback == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_interface_set_event_callback", "10.10", _vmnet_interface_set_event_callbackErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Interface_event_t, arg1 objectivec.Object) { callback(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_set_event_callback(interface_, event_mask, uintptr(queue.Handle()), _block0), nil
}

// Vmnet_interface_set_event_callback schedules a callback to be executed when events for the specified interface are received.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_set_event_callback(_:_:_:_:)
func Vmnet_interface_set_event_callback(interface_ Interface_ref, event_mask Interface_event_t, queue dispatch.Queue, callback Vmnet_interface_event_callback_t) Vmnet_return_t {
	result, callErr := tryVmnet_interface_set_event_callback(interface_, event_mask, queue, callback)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_interface_start_with_network func(network Vmnet_network_ref, interface_desc unsafe.Pointer, queue uintptr, start_block unsafe.Pointer) Interface_ref
var _vmnet_interface_start_with_networkErr error

func tryVmnet_interface_start_with_network(network Vmnet_network_ref, interface_desc unsafe.Pointer, queue dispatch.Queue, start_block Vmnet_start_interface_completion_handler_t) (Interface_ref, error) {
	if _vmnet_interface_start_with_network == nil {
		return *new(Interface_ref), symbolCallError("vmnet_interface_start_with_network", "26.0", _vmnet_interface_start_with_networkErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t, arg1 objectivec.Object) { start_block(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_interface_start_with_network(network, interface_desc, uintptr(queue.Handle()), _block0), nil
}

// Vmnet_interface_start_with_network.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_interface_start_with_network(_:_:_:_:)
func Vmnet_interface_start_with_network(network Vmnet_network_ref, interface_desc unsafe.Pointer, queue dispatch.Queue, start_block Vmnet_start_interface_completion_handler_t) Interface_ref {
	result, callErr := tryVmnet_interface_start_with_network(network, interface_desc, queue, start_block)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_ip_port_forwarding_rule_get_details func(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, address_family uint8, internal_address unsafe.Pointer, internal_port *uint16) Vmnet_return_t
var _vmnet_ip_port_forwarding_rule_get_detailsErr error

func tryVmnet_ip_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, address_family uint8, internal_address unsafe.Pointer, internal_port *uint16) (Vmnet_return_t, error) {
	if _vmnet_ip_port_forwarding_rule_get_details == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_ip_port_forwarding_rule_get_details", "11.0", _vmnet_ip_port_forwarding_rule_get_detailsErr)
	}
	return _vmnet_ip_port_forwarding_rule_get_details(rule, protocol_, external_port, address_family, internal_address, internal_port), nil
}

// Vmnet_ip_port_forwarding_rule_get_details.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_ip_port_forwarding_rule_get_details(_:_:_:_:_:_:)
func Vmnet_ip_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, address_family uint8, internal_address unsafe.Pointer, internal_port *uint16) Vmnet_return_t {
	result, callErr := tryVmnet_ip_port_forwarding_rule_get_details(rule, protocol_, external_port, address_family, internal_address, internal_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_add_port_forwarding_rule func(config Vmnet_network_configuration_ref, protocol_ uint8, address_family uint8, internal_port uint16, external_port uint16, internal_address unsafe.Pointer) Vmnet_return_t
var _vmnet_network_configuration_add_port_forwarding_ruleErr error

func tryVmnet_network_configuration_add_port_forwarding_rule(config Vmnet_network_configuration_ref, protocol_ uint8, address_family uint8, internal_port uint16, external_port uint16, internal_address unsafe.Pointer) (Vmnet_return_t, error) {
	if _vmnet_network_configuration_add_port_forwarding_rule == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_network_configuration_add_port_forwarding_rule", "26.0", _vmnet_network_configuration_add_port_forwarding_ruleErr)
	}
	return _vmnet_network_configuration_add_port_forwarding_rule(config, protocol_, address_family, internal_port, external_port, internal_address), nil
}

// Vmnet_network_configuration_add_port_forwarding_rule.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_add_port_forwarding_rule(_:_:_:_:_:_:)
func Vmnet_network_configuration_add_port_forwarding_rule(config Vmnet_network_configuration_ref, protocol_ uint8, address_family uint8, internal_port uint16, external_port uint16, internal_address unsafe.Pointer) Vmnet_return_t {
	result, callErr := tryVmnet_network_configuration_add_port_forwarding_rule(config, protocol_, address_family, internal_port, external_port, internal_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_create func(mode Operating_modes_t, status *Vmnet_return_t) Vmnet_network_configuration_ref
var _vmnet_network_configuration_createErr error

func tryVmnet_network_configuration_create(mode Operating_modes_t, status *Vmnet_return_t) (Vmnet_network_configuration_ref, error) {
	if _vmnet_network_configuration_create == nil {
		return *new(Vmnet_network_configuration_ref), symbolCallError("vmnet_network_configuration_create", "26.0", _vmnet_network_configuration_createErr)
	}
	return _vmnet_network_configuration_create(mode, status), nil
}

// Vmnet_network_configuration_create.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_create(_:_:)
func Vmnet_network_configuration_create(mode Operating_modes_t, status *Vmnet_return_t) Vmnet_network_configuration_ref {
	result, callErr := tryVmnet_network_configuration_create(mode, status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_disable_dhcp func(config Vmnet_network_configuration_ref)
var _vmnet_network_configuration_disable_dhcpErr error

func tryVmnet_network_configuration_disable_dhcp(config Vmnet_network_configuration_ref) error {
	if _vmnet_network_configuration_disable_dhcp == nil {
		return symbolCallError("vmnet_network_configuration_disable_dhcp", "26.0", _vmnet_network_configuration_disable_dhcpErr)
	}
	_vmnet_network_configuration_disable_dhcp(config)
	return nil
}

// Vmnet_network_configuration_disable_dhcp.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dhcp(_:)
func Vmnet_network_configuration_disable_dhcp(config Vmnet_network_configuration_ref) {
	if callErr := tryVmnet_network_configuration_disable_dhcp(config); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_configuration_disable_dns_proxy func(config Vmnet_network_configuration_ref)
var _vmnet_network_configuration_disable_dns_proxyErr error

func tryVmnet_network_configuration_disable_dns_proxy(config Vmnet_network_configuration_ref) error {
	if _vmnet_network_configuration_disable_dns_proxy == nil {
		return symbolCallError("vmnet_network_configuration_disable_dns_proxy", "26.0", _vmnet_network_configuration_disable_dns_proxyErr)
	}
	_vmnet_network_configuration_disable_dns_proxy(config)
	return nil
}

// Vmnet_network_configuration_disable_dns_proxy.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_dns_proxy(_:)
func Vmnet_network_configuration_disable_dns_proxy(config Vmnet_network_configuration_ref) {
	if callErr := tryVmnet_network_configuration_disable_dns_proxy(config); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_configuration_disable_nat44 func(config Vmnet_network_configuration_ref)
var _vmnet_network_configuration_disable_nat44Err error

func tryVmnet_network_configuration_disable_nat44(config Vmnet_network_configuration_ref) error {
	if _vmnet_network_configuration_disable_nat44 == nil {
		return symbolCallError("vmnet_network_configuration_disable_nat44", "26.0", _vmnet_network_configuration_disable_nat44Err)
	}
	_vmnet_network_configuration_disable_nat44(config)
	return nil
}

// Vmnet_network_configuration_disable_nat44.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat44(_:)
func Vmnet_network_configuration_disable_nat44(config Vmnet_network_configuration_ref) {
	if callErr := tryVmnet_network_configuration_disable_nat44(config); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_configuration_disable_nat66 func(config Vmnet_network_configuration_ref)
var _vmnet_network_configuration_disable_nat66Err error

func tryVmnet_network_configuration_disable_nat66(config Vmnet_network_configuration_ref) error {
	if _vmnet_network_configuration_disable_nat66 == nil {
		return symbolCallError("vmnet_network_configuration_disable_nat66", "26.0", _vmnet_network_configuration_disable_nat66Err)
	}
	_vmnet_network_configuration_disable_nat66(config)
	return nil
}

// Vmnet_network_configuration_disable_nat66.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_nat66(_:)
func Vmnet_network_configuration_disable_nat66(config Vmnet_network_configuration_ref) {
	if callErr := tryVmnet_network_configuration_disable_nat66(config); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_configuration_disable_router_advertisement func(config Vmnet_network_configuration_ref)
var _vmnet_network_configuration_disable_router_advertisementErr error

func tryVmnet_network_configuration_disable_router_advertisement(config Vmnet_network_configuration_ref) error {
	if _vmnet_network_configuration_disable_router_advertisement == nil {
		return symbolCallError("vmnet_network_configuration_disable_router_advertisement", "26.0", _vmnet_network_configuration_disable_router_advertisementErr)
	}
	_vmnet_network_configuration_disable_router_advertisement(config)
	return nil
}

// Vmnet_network_configuration_disable_router_advertisement.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_disable_router_advertisement(_:)
func Vmnet_network_configuration_disable_router_advertisement(config Vmnet_network_configuration_ref) {
	if callErr := tryVmnet_network_configuration_disable_router_advertisement(config); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_configuration_set_external_interface func(config Vmnet_network_configuration_ref, interface_name string) Vmnet_return_t
var _vmnet_network_configuration_set_external_interfaceErr error

func tryVmnet_network_configuration_set_external_interface(config Vmnet_network_configuration_ref, interface_name string) (Vmnet_return_t, error) {
	if _vmnet_network_configuration_set_external_interface == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_network_configuration_set_external_interface", "26.0", _vmnet_network_configuration_set_external_interfaceErr)
	}
	return _vmnet_network_configuration_set_external_interface(config, interface_name), nil
}

// Vmnet_network_configuration_set_external_interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_external_interface(_:_:)
func Vmnet_network_configuration_set_external_interface(config Vmnet_network_configuration_ref, interface_name string) Vmnet_return_t {
	result, callErr := tryVmnet_network_configuration_set_external_interface(config, interface_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_set_ipv4_subnet func(config Vmnet_network_configuration_ref, subnet_addr uintptr, subnet_mask uintptr) Vmnet_return_t
var _vmnet_network_configuration_set_ipv4_subnetErr error

func tryVmnet_network_configuration_set_ipv4_subnet(config Vmnet_network_configuration_ref, subnet_addr uintptr, subnet_mask uintptr) (Vmnet_return_t, error) {
	if _vmnet_network_configuration_set_ipv4_subnet == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_network_configuration_set_ipv4_subnet", "26.0", _vmnet_network_configuration_set_ipv4_subnetErr)
	}
	return _vmnet_network_configuration_set_ipv4_subnet(config, subnet_addr, subnet_mask), nil
}

// Vmnet_network_configuration_set_ipv4_subnet.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_ipv4_subnet(_:_:_:)
func Vmnet_network_configuration_set_ipv4_subnet(config Vmnet_network_configuration_ref, subnet_addr uintptr, subnet_mask uintptr) Vmnet_return_t {
	result, callErr := tryVmnet_network_configuration_set_ipv4_subnet(config, subnet_addr, subnet_mask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_set_ipv6_prefix func(config Vmnet_network_configuration_ref, prefix uintptr, len_ uint8) Vmnet_return_t
var _vmnet_network_configuration_set_ipv6_prefixErr error

func tryVmnet_network_configuration_set_ipv6_prefix(config Vmnet_network_configuration_ref, prefix uintptr, len_ uint8) (Vmnet_return_t, error) {
	if _vmnet_network_configuration_set_ipv6_prefix == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_network_configuration_set_ipv6_prefix", "26.0", _vmnet_network_configuration_set_ipv6_prefixErr)
	}
	return _vmnet_network_configuration_set_ipv6_prefix(config, prefix, len_), nil
}

// Vmnet_network_configuration_set_ipv6_prefix.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_ipv6_prefix(_:_:_:)
func Vmnet_network_configuration_set_ipv6_prefix(config Vmnet_network_configuration_ref, prefix uintptr, len_ uint8) Vmnet_return_t {
	result, callErr := tryVmnet_network_configuration_set_ipv6_prefix(config, prefix, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_configuration_set_mtu func(config Vmnet_network_configuration_ref, mtu uint32) Vmnet_return_t
var _vmnet_network_configuration_set_mtuErr error

func tryVmnet_network_configuration_set_mtu(config Vmnet_network_configuration_ref, mtu uint32) (Vmnet_return_t, error) {
	if _vmnet_network_configuration_set_mtu == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_network_configuration_set_mtu", "26.0", _vmnet_network_configuration_set_mtuErr)
	}
	return _vmnet_network_configuration_set_mtu(config, mtu), nil
}

// Vmnet_network_configuration_set_mtu.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_configuration_set_mtu(_:_:)
func Vmnet_network_configuration_set_mtu(config Vmnet_network_configuration_ref, mtu uint32) Vmnet_return_t {
	result, callErr := tryVmnet_network_configuration_set_mtu(config, mtu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_copy_serialization func(network Vmnet_network_ref, status *Vmnet_return_t) unsafe.Pointer
var _vmnet_network_copy_serializationErr error

func tryVmnet_network_copy_serialization(network Vmnet_network_ref, status *Vmnet_return_t) (unsafe.Pointer, error) {
	if _vmnet_network_copy_serialization == nil {
		return nil, symbolCallError("vmnet_network_copy_serialization", "26.0", _vmnet_network_copy_serializationErr)
	}
	return _vmnet_network_copy_serialization(network, status), nil
}

// Vmnet_network_copy_serialization.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_copy_serialization(_:_:)
func Vmnet_network_copy_serialization(network Vmnet_network_ref, status *Vmnet_return_t) unsafe.Pointer {
	result, callErr := tryVmnet_network_copy_serialization(network, status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_create func(configuration Vmnet_network_configuration_ref, status *Vmnet_return_t) Vmnet_network_ref
var _vmnet_network_createErr error

func tryVmnet_network_create(configuration Vmnet_network_configuration_ref, status *Vmnet_return_t) (Vmnet_network_ref, error) {
	if _vmnet_network_create == nil {
		return *new(Vmnet_network_ref), symbolCallError("vmnet_network_create", "26.0", _vmnet_network_createErr)
	}
	return _vmnet_network_create(configuration, status), nil
}

// Vmnet_network_create.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_create(_:_:)
func Vmnet_network_create(configuration Vmnet_network_configuration_ref, status *Vmnet_return_t) Vmnet_network_ref {
	result, callErr := tryVmnet_network_create(configuration, status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_create_with_serialization func(network unsafe.Pointer, status *Vmnet_return_t) Vmnet_network_ref
var _vmnet_network_create_with_serializationErr error

func tryVmnet_network_create_with_serialization(network unsafe.Pointer, status *Vmnet_return_t) (Vmnet_network_ref, error) {
	if _vmnet_network_create_with_serialization == nil {
		return *new(Vmnet_network_ref), symbolCallError("vmnet_network_create_with_serialization", "26.0", _vmnet_network_create_with_serializationErr)
	}
	return _vmnet_network_create_with_serialization(network, status), nil
}

// Vmnet_network_create_with_serialization.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_create_with_serialization(_:_:)
func Vmnet_network_create_with_serialization(network unsafe.Pointer, status *Vmnet_return_t) Vmnet_network_ref {
	result, callErr := tryVmnet_network_create_with_serialization(network, status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_network_get_ipv4_subnet func(network Vmnet_network_ref, subnet uintptr, mask uintptr)
var _vmnet_network_get_ipv4_subnetErr error

func tryVmnet_network_get_ipv4_subnet(network Vmnet_network_ref, subnet uintptr, mask uintptr) error {
	if _vmnet_network_get_ipv4_subnet == nil {
		return symbolCallError("vmnet_network_get_ipv4_subnet", "26.0", _vmnet_network_get_ipv4_subnetErr)
	}
	_vmnet_network_get_ipv4_subnet(network, subnet, mask)
	return nil
}

// Vmnet_network_get_ipv4_subnet.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv4_subnet(_:_:_:)
func Vmnet_network_get_ipv4_subnet(network Vmnet_network_ref, subnet uintptr, mask uintptr) {
	if callErr := tryVmnet_network_get_ipv4_subnet(network, subnet, mask); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_network_get_ipv6_prefix func(network Vmnet_network_ref, prefix uintptr, prefix_len *uint8)
var _vmnet_network_get_ipv6_prefixErr error

func tryVmnet_network_get_ipv6_prefix(network Vmnet_network_ref, prefix uintptr, prefix_len *uint8) error {
	if _vmnet_network_get_ipv6_prefix == nil {
		return symbolCallError("vmnet_network_get_ipv6_prefix", "26.0", _vmnet_network_get_ipv6_prefixErr)
	}
	_vmnet_network_get_ipv6_prefix(network, prefix, prefix_len)
	return nil
}

// Vmnet_network_get_ipv6_prefix.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_network_get_ipv6_prefix(_:_:_:)
func Vmnet_network_get_ipv6_prefix(network Vmnet_network_ref, prefix uintptr, prefix_len *uint8) {
	if callErr := tryVmnet_network_get_ipv6_prefix(network, prefix, prefix_len); callErr != nil {
		panic(callErr)
	}
}

var _vmnet_port_forwarding_rule_get_details func(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, internal_address uintptr, internal_port *uint16) Vmnet_return_t
var _vmnet_port_forwarding_rule_get_detailsErr error

func tryVmnet_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, internal_address uintptr, internal_port *uint16) (Vmnet_return_t, error) {
	if _vmnet_port_forwarding_rule_get_details == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_port_forwarding_rule_get_details", "10.15", _vmnet_port_forwarding_rule_get_detailsErr)
	}
	return _vmnet_port_forwarding_rule_get_details(rule, protocol_, external_port, internal_address, internal_port), nil
}

// Vmnet_port_forwarding_rule_get_details.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_port_forwarding_rule_get_details(_:_:_:_:_:)
func Vmnet_port_forwarding_rule_get_details(rule unsafe.Pointer, protocol_ *uint8, external_port *uint16, internal_address uintptr, internal_port *uint16) Vmnet_return_t {
	result, callErr := tryVmnet_port_forwarding_rule_get_details(rule, protocol_, external_port, internal_address, internal_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_read func(interface_ Interface_ref, packets *Vmpktdesc, pktcnt *int) Vmnet_return_t
var _vmnet_readErr error

func tryVmnet_read(interface_ Interface_ref, packets *Vmpktdesc, pktcnt []int) (Vmnet_return_t, error) {
	if _vmnet_read == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_read", "10.10", _vmnet_readErr)
	}
	return _vmnet_read(interface_, packets, unsafe.SliceData(pktcnt)), nil
}

// Vmnet_read attempts to read a specified number of packets from an interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_read(_:_:_:)
func Vmnet_read(interface_ Interface_ref, packets *Vmpktdesc, pktcnt []int) Vmnet_return_t {
	result, callErr := tryVmnet_read(interface_, packets, pktcnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_start_interface func(interface_desc unsafe.Pointer, queue uintptr, handler unsafe.Pointer) Interface_ref
var _vmnet_start_interfaceErr error

func tryVmnet_start_interface(interface_desc unsafe.Pointer, queue dispatch.Queue, handler Vmnet_start_interface_completion_handler_t) (Interface_ref, error) {
	if _vmnet_start_interface == nil {
		return *new(Interface_ref), symbolCallError("vmnet_start_interface", "10.10", _vmnet_start_interfaceErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t, arg1 objectivec.Object) { handler(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_start_interface(interface_desc, uintptr(queue.Handle()), _block0), nil
}

// Vmnet_start_interface starts host or shared mode on an interface with a specified configuration.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_start_interface(_:_:_:)
func Vmnet_start_interface(interface_desc unsafe.Pointer, queue dispatch.Queue, handler Vmnet_start_interface_completion_handler_t) Interface_ref {
	result, callErr := tryVmnet_start_interface(interface_desc, queue, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_stop_interface func(interface_ Interface_ref, queue uintptr, handler unsafe.Pointer) Vmnet_return_t
var _vmnet_stop_interfaceErr error

func tryVmnet_stop_interface(interface_ Interface_ref, queue dispatch.Queue, handler Vmnet_interface_completion_handler_t) (Vmnet_return_t, error) {
	if _vmnet_stop_interface == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_stop_interface", "10.10", _vmnet_stop_interfaceErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 Vmnet_return_t) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vmnet_stop_interface(interface_, uintptr(queue.Handle()), _block0), nil
}

// Vmnet_stop_interface stops the interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_stop_interface(_:_:_:)
func Vmnet_stop_interface(interface_ Interface_ref, queue dispatch.Queue, handler Vmnet_interface_completion_handler_t) Vmnet_return_t {
	result, callErr := tryVmnet_stop_interface(interface_, queue, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vmnet_write func(interface_ Interface_ref, packets *Vmpktdesc, pktcnt *int) Vmnet_return_t
var _vmnet_writeErr error

func tryVmnet_write(interface_ Interface_ref, packets *Vmpktdesc, pktcnt []int) (Vmnet_return_t, error) {
	if _vmnet_write == nil {
		return *new(Vmnet_return_t), symbolCallError("vmnet_write", "10.10", _vmnet_writeErr)
	}
	return _vmnet_write(interface_, packets, unsafe.SliceData(pktcnt)), nil
}

// Vmnet_write attempts to write specified packets to an interface.
//
// See: https://developer.apple.com/documentation/vmnet/vmnet_write(_:_:_:)
func Vmnet_write(interface_ Interface_ref, packets *Vmpktdesc, pktcnt []int) Vmnet_return_t {
	result, callErr := tryVmnet_write(interface_, packets, pktcnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_vmnet_copy_shared_interface_list, &_vmnet_copy_shared_interface_listErr, frameworkHandle, "vmnet_copy_shared_interface_list", "10.15")
	registerFunc(&_vmnet_interface_add_ip_port_forwarding_rule, &_vmnet_interface_add_ip_port_forwarding_ruleErr, frameworkHandle, "vmnet_interface_add_ip_port_forwarding_rule", "11.0")
	registerFunc(&_vmnet_interface_add_port_forwarding_rule, &_vmnet_interface_add_port_forwarding_ruleErr, frameworkHandle, "vmnet_interface_add_port_forwarding_rule", "10.15")
	registerFunc(&_vmnet_interface_get_ip_port_forwarding_rules, &_vmnet_interface_get_ip_port_forwarding_rulesErr, frameworkHandle, "vmnet_interface_get_ip_port_forwarding_rules", "11.0")
	registerFunc(&_vmnet_interface_get_port_forwarding_rules, &_vmnet_interface_get_port_forwarding_rulesErr, frameworkHandle, "vmnet_interface_get_port_forwarding_rules", "10.15")
	registerFunc(&_vmnet_interface_remove_ip_port_forwarding_rule, &_vmnet_interface_remove_ip_port_forwarding_ruleErr, frameworkHandle, "vmnet_interface_remove_ip_port_forwarding_rule", "11.0")
	registerFunc(&_vmnet_interface_remove_port_forwarding_rule, &_vmnet_interface_remove_port_forwarding_ruleErr, frameworkHandle, "vmnet_interface_remove_port_forwarding_rule", "10.15")
	registerFunc(&_vmnet_interface_set_event_callback, &_vmnet_interface_set_event_callbackErr, frameworkHandle, "vmnet_interface_set_event_callback", "10.10")
	registerFunc(&_vmnet_interface_start_with_network, &_vmnet_interface_start_with_networkErr, frameworkHandle, "vmnet_interface_start_with_network", "26.0")
	registerFunc(&_vmnet_ip_port_forwarding_rule_get_details, &_vmnet_ip_port_forwarding_rule_get_detailsErr, frameworkHandle, "vmnet_ip_port_forwarding_rule_get_details", "11.0")
	registerFunc(&_vmnet_network_configuration_add_port_forwarding_rule, &_vmnet_network_configuration_add_port_forwarding_ruleErr, frameworkHandle, "vmnet_network_configuration_add_port_forwarding_rule", "26.0")
	registerFunc(&_vmnet_network_configuration_create, &_vmnet_network_configuration_createErr, frameworkHandle, "vmnet_network_configuration_create", "26.0")
	registerFunc(&_vmnet_network_configuration_disable_dhcp, &_vmnet_network_configuration_disable_dhcpErr, frameworkHandle, "vmnet_network_configuration_disable_dhcp", "26.0")
	registerFunc(&_vmnet_network_configuration_disable_dns_proxy, &_vmnet_network_configuration_disable_dns_proxyErr, frameworkHandle, "vmnet_network_configuration_disable_dns_proxy", "26.0")
	registerFunc(&_vmnet_network_configuration_disable_nat44, &_vmnet_network_configuration_disable_nat44Err, frameworkHandle, "vmnet_network_configuration_disable_nat44", "26.0")
	registerFunc(&_vmnet_network_configuration_disable_nat66, &_vmnet_network_configuration_disable_nat66Err, frameworkHandle, "vmnet_network_configuration_disable_nat66", "26.0")
	registerFunc(&_vmnet_network_configuration_disable_router_advertisement, &_vmnet_network_configuration_disable_router_advertisementErr, frameworkHandle, "vmnet_network_configuration_disable_router_advertisement", "26.0")
	registerFunc(&_vmnet_network_configuration_set_external_interface, &_vmnet_network_configuration_set_external_interfaceErr, frameworkHandle, "vmnet_network_configuration_set_external_interface", "26.0")
	registerFunc(&_vmnet_network_configuration_set_ipv4_subnet, &_vmnet_network_configuration_set_ipv4_subnetErr, frameworkHandle, "vmnet_network_configuration_set_ipv4_subnet", "26.0")
	registerFunc(&_vmnet_network_configuration_set_ipv6_prefix, &_vmnet_network_configuration_set_ipv6_prefixErr, frameworkHandle, "vmnet_network_configuration_set_ipv6_prefix", "26.0")
	registerFunc(&_vmnet_network_configuration_set_mtu, &_vmnet_network_configuration_set_mtuErr, frameworkHandle, "vmnet_network_configuration_set_mtu", "26.0")
	registerFunc(&_vmnet_network_copy_serialization, &_vmnet_network_copy_serializationErr, frameworkHandle, "vmnet_network_copy_serialization", "26.0")
	registerFunc(&_vmnet_network_create, &_vmnet_network_createErr, frameworkHandle, "vmnet_network_create", "26.0")
	registerFunc(&_vmnet_network_create_with_serialization, &_vmnet_network_create_with_serializationErr, frameworkHandle, "vmnet_network_create_with_serialization", "26.0")
	registerFunc(&_vmnet_network_get_ipv4_subnet, &_vmnet_network_get_ipv4_subnetErr, frameworkHandle, "vmnet_network_get_ipv4_subnet", "26.0")
	registerFunc(&_vmnet_network_get_ipv6_prefix, &_vmnet_network_get_ipv6_prefixErr, frameworkHandle, "vmnet_network_get_ipv6_prefix", "26.0")
	registerFunc(&_vmnet_port_forwarding_rule_get_details, &_vmnet_port_forwarding_rule_get_detailsErr, frameworkHandle, "vmnet_port_forwarding_rule_get_details", "10.15")
	registerFunc(&_vmnet_read, &_vmnet_readErr, frameworkHandle, "vmnet_read", "10.10")
	registerFunc(&_vmnet_start_interface, &_vmnet_start_interfaceErr, frameworkHandle, "vmnet_start_interface", "10.10")
	registerFunc(&_vmnet_stop_interface, &_vmnet_stop_interfaceErr, frameworkHandle, "vmnet_stop_interface", "10.10")
	registerFunc(&_vmnet_write, &_vmnet_writeErr, frameworkHandle, "vmnet_write", "10.10")
}
