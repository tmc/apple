// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Network: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Network: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Network: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _nw_advertise_descriptor_copy_txt_record_object func(advertise_descriptor Nw_advertise_descriptor_t) Nw_txt_record_t

// Nw_advertise_descriptor_copy_txt_record_object accesses the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_copy_txt_record_object(_:)
func Nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t) Nw_txt_record_t {
	if _nw_advertise_descriptor_copy_txt_record_object == nil {
		panic("Network: symbol nw_advertise_descriptor_copy_txt_record_object not loaded")
	}
	return _nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor)
}

var _nw_advertise_descriptor_create_application_service func(application_service_name string) Nw_advertise_descriptor_t

// Nw_advertise_descriptor_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_application_service(_:)
func Nw_advertise_descriptor_create_application_service(application_service_name string) Nw_advertise_descriptor_t {
	if _nw_advertise_descriptor_create_application_service == nil {
		panic("Network: symbol nw_advertise_descriptor_create_application_service not loaded")
	}
	return _nw_advertise_descriptor_create_application_service(application_service_name)
}

var _nw_advertise_descriptor_create_bonjour_service func(name string, type_ string, domain string) Nw_advertise_descriptor_t

// Nw_advertise_descriptor_create_bonjour_service initializes a Bonjour service to advertise.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_bonjour_service(_:_:_:)
func Nw_advertise_descriptor_create_bonjour_service(name string, type_ string, domain string) Nw_advertise_descriptor_t {
	if _nw_advertise_descriptor_create_bonjour_service == nil {
		panic("Network: symbol nw_advertise_descriptor_create_bonjour_service not loaded")
	}
	return _nw_advertise_descriptor_create_bonjour_service(name, type_, domain)
}

var _nw_advertise_descriptor_get_application_service_name func(advertise_descriptor Nw_advertise_descriptor_t) *byte

// Nw_advertise_descriptor_get_application_service_name.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_application_service_name(_:)
func Nw_advertise_descriptor_get_application_service_name(advertise_descriptor Nw_advertise_descriptor_t) *byte {
	if _nw_advertise_descriptor_get_application_service_name == nil {
		panic("Network: symbol nw_advertise_descriptor_get_application_service_name not loaded")
	}
	return _nw_advertise_descriptor_get_application_service_name(advertise_descriptor)
}

var _nw_advertise_descriptor_get_no_auto_rename func(advertise_descriptor Nw_advertise_descriptor_t) bool

// Nw_advertise_descriptor_get_no_auto_rename checks whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_no_auto_rename(_:)
func Nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t) bool {
	if _nw_advertise_descriptor_get_no_auto_rename == nil {
		panic("Network: symbol nw_advertise_descriptor_get_no_auto_rename not loaded")
	}
	return _nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor)
}

var _nw_advertise_descriptor_set_no_auto_rename func(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool)

// Nw_advertise_descriptor_set_no_auto_rename sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_no_auto_rename(_:_:)
func Nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool) {
	if _nw_advertise_descriptor_set_no_auto_rename == nil {
		panic("Network: symbol nw_advertise_descriptor_set_no_auto_rename not loaded")
	}
	_nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor, no_auto_rename)
}

var _nw_advertise_descriptor_set_txt_record func(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr)

// Nw_advertise_descriptor_set_txt_record sets the TXT record as a raw buffer to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record(_:_:_:)
func Nw_advertise_descriptor_set_txt_record(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr) {
	if _nw_advertise_descriptor_set_txt_record == nil {
		panic("Network: symbol nw_advertise_descriptor_set_txt_record not loaded")
	}
	_nw_advertise_descriptor_set_txt_record(advertise_descriptor, txt_record, txt_length)
}

var _nw_advertise_descriptor_set_txt_record_object func(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t)

// Nw_advertise_descriptor_set_txt_record_object sets the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record_object(_:_:)
func Nw_advertise_descriptor_set_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t) {
	if _nw_advertise_descriptor_set_txt_record_object == nil {
		panic("Network: symbol nw_advertise_descriptor_set_txt_record_object not loaded")
	}
	_nw_advertise_descriptor_set_txt_record_object(advertise_descriptor, txt_record)
}

var _nw_browse_descriptor_create_application_service func(application_service_name string) Nw_browse_descriptor_t

// Nw_browse_descriptor_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_application_service(_:)
func Nw_browse_descriptor_create_application_service(application_service_name string) Nw_browse_descriptor_t {
	if _nw_browse_descriptor_create_application_service == nil {
		panic("Network: symbol nw_browse_descriptor_create_application_service not loaded")
	}
	return _nw_browse_descriptor_create_application_service(application_service_name)
}

var _nw_browse_descriptor_create_bonjour_service func(type_ string, domain string) Nw_browse_descriptor_t

// Nw_browse_descriptor_create_bonjour_service initializes a service descriptor used to discover a Bonjour service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_bonjour_service(_:_:)
func Nw_browse_descriptor_create_bonjour_service(type_ string, domain string) Nw_browse_descriptor_t {
	if _nw_browse_descriptor_create_bonjour_service == nil {
		panic("Network: symbol nw_browse_descriptor_create_bonjour_service not loaded")
	}
	return _nw_browse_descriptor_create_bonjour_service(type_, domain)
}

var _nw_browse_descriptor_get_application_service_name func(descriptor Nw_browse_descriptor_t) *byte

// Nw_browse_descriptor_get_application_service_name.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_application_service_name(_:)
func Nw_browse_descriptor_get_application_service_name(descriptor Nw_browse_descriptor_t) *byte {
	if _nw_browse_descriptor_get_application_service_name == nil {
		panic("Network: symbol nw_browse_descriptor_get_application_service_name not loaded")
	}
	return _nw_browse_descriptor_get_application_service_name(descriptor)
}

var _nw_browse_descriptor_get_bonjour_service_domain func(descriptor Nw_browse_descriptor_t) *byte

// Nw_browse_descriptor_get_bonjour_service_domain accesses the Bonjour service domain set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_domain(_:)
func Nw_browse_descriptor_get_bonjour_service_domain(descriptor Nw_browse_descriptor_t) *byte {
	if _nw_browse_descriptor_get_bonjour_service_domain == nil {
		panic("Network: symbol nw_browse_descriptor_get_bonjour_service_domain not loaded")
	}
	return _nw_browse_descriptor_get_bonjour_service_domain(descriptor)
}

var _nw_browse_descriptor_get_bonjour_service_type func(descriptor Nw_browse_descriptor_t) *byte

// Nw_browse_descriptor_get_bonjour_service_type accesses the Bonjour service type set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_type(_:)
func Nw_browse_descriptor_get_bonjour_service_type(descriptor Nw_browse_descriptor_t) *byte {
	if _nw_browse_descriptor_get_bonjour_service_type == nil {
		panic("Network: symbol nw_browse_descriptor_get_bonjour_service_type not loaded")
	}
	return _nw_browse_descriptor_get_bonjour_service_type(descriptor)
}

var _nw_browse_descriptor_get_include_txt_record func(descriptor Nw_browse_descriptor_t) bool

// Nw_browse_descriptor_get_include_txt_record checks if the browse descriptor requires including associated TXT records with all results.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_include_txt_record(_:)
func Nw_browse_descriptor_get_include_txt_record(descriptor Nw_browse_descriptor_t) bool {
	if _nw_browse_descriptor_get_include_txt_record == nil {
		panic("Network: symbol nw_browse_descriptor_get_include_txt_record not loaded")
	}
	return _nw_browse_descriptor_get_include_txt_record(descriptor)
}

var _nw_browse_descriptor_set_include_txt_record func(descriptor Nw_browse_descriptor_t, include_txt_record bool)

// Nw_browse_descriptor_set_include_txt_record requires including associated TXT records with all results generated for this service descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_set_include_txt_record(_:_:)
func Nw_browse_descriptor_set_include_txt_record(descriptor Nw_browse_descriptor_t, include_txt_record bool) {
	if _nw_browse_descriptor_set_include_txt_record == nil {
		panic("Network: symbol nw_browse_descriptor_set_include_txt_record not loaded")
	}
	_nw_browse_descriptor_set_include_txt_record(descriptor, include_txt_record)
}

var _nw_browse_result_copy_endpoint func(result Nw_browse_result_t) Nw_endpoint_t

// Nw_browse_result_copy_endpoint the discovered service endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_endpoint(_:)
func Nw_browse_result_copy_endpoint(result Nw_browse_result_t) Nw_endpoint_t {
	if _nw_browse_result_copy_endpoint == nil {
		panic("Network: symbol nw_browse_result_copy_endpoint not loaded")
	}
	return _nw_browse_result_copy_endpoint(result)
}

var _nw_browse_result_copy_txt_record_object func(result Nw_browse_result_t) Nw_txt_record_t

// Nw_browse_result_copy_txt_record_object accesses the TXT record associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_txt_record_object(_:)
func Nw_browse_result_copy_txt_record_object(result Nw_browse_result_t) Nw_txt_record_t {
	if _nw_browse_result_copy_txt_record_object == nil {
		panic("Network: symbol nw_browse_result_copy_txt_record_object not loaded")
	}
	return _nw_browse_result_copy_txt_record_object(result)
}

var _nw_browse_result_enumerate_interfaces func(result Nw_browse_result_t, enumerator Nw_browse_result_enumerate_interface_t)

// Nw_browse_result_enumerate_interfaces enumerates the list of interfaces on which the service was discovered.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interfaces(_:_:)
func Nw_browse_result_enumerate_interfaces(result Nw_browse_result_t, enumerator Nw_browse_result_enumerate_interface_t) {
	if _nw_browse_result_enumerate_interfaces == nil {
		panic("Network: symbol nw_browse_result_enumerate_interfaces not loaded")
	}
	_nw_browse_result_enumerate_interfaces(result, enumerator)
}

var _nw_browse_result_get_changes func(old_result Nw_browse_result_t, new_result Nw_browse_result_t) Nw_browse_result_change_t

// Nw_browse_result_get_changes compares two discovered services and calculates changes between them.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_changes(_:_:)
func Nw_browse_result_get_changes(old_result Nw_browse_result_t, new_result Nw_browse_result_t) Nw_browse_result_change_t {
	if _nw_browse_result_get_changes == nil {
		panic("Network: symbol nw_browse_result_get_changes not loaded")
	}
	return _nw_browse_result_get_changes(old_result, new_result)
}

var _nw_browse_result_get_interfaces_count func(result Nw_browse_result_t) uintptr

// Nw_browse_result_get_interfaces_count accesses the number of interfaces associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_interfaces_count(_:)
func Nw_browse_result_get_interfaces_count(result Nw_browse_result_t) uintptr {
	if _nw_browse_result_get_interfaces_count == nil {
		panic("Network: symbol nw_browse_result_get_interfaces_count not loaded")
	}
	return _nw_browse_result_get_interfaces_count(result)
}

var _nw_browser_cancel func(browser Nw_browser_t)

// Nw_browser_cancel stops browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_cancel(_:)
func Nw_browser_cancel(browser Nw_browser_t) {
	if _nw_browser_cancel == nil {
		panic("Network: symbol nw_browser_cancel not loaded")
	}
	_nw_browser_cancel(browser)
}

var _nw_browser_copy_browse_descriptor func(browser Nw_browser_t) Nw_browse_descriptor_t

// Nw_browser_copy_browse_descriptor accesses the service descriptor with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_browse_descriptor(_:)
func Nw_browser_copy_browse_descriptor(browser Nw_browser_t) Nw_browse_descriptor_t {
	if _nw_browser_copy_browse_descriptor == nil {
		panic("Network: symbol nw_browser_copy_browse_descriptor not loaded")
	}
	return _nw_browser_copy_browse_descriptor(browser)
}

var _nw_browser_copy_parameters func(browser Nw_browser_t) Nw_parameters_t

// Nw_browser_copy_parameters accesses the parameters with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_parameters(_:)
func Nw_browser_copy_parameters(browser Nw_browser_t) Nw_parameters_t {
	if _nw_browser_copy_parameters == nil {
		panic("Network: symbol nw_browser_copy_parameters not loaded")
	}
	return _nw_browser_copy_parameters(browser)
}

var _nw_browser_create func(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) Nw_browser_t

// Nw_browser_create initializes a browser with a type of service to discover.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_create(_:_:)
func Nw_browser_create(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) Nw_browser_t {
	if _nw_browser_create == nil {
		panic("Network: symbol nw_browser_create not loaded")
	}
	return _nw_browser_create(descriptor, parameters)
}

var _nw_browser_set_browse_results_changed_handler func(browser Nw_browser_t, handler Nw_browser_browse_results_changed_handler_t)

// Nw_browser_set_browse_results_changed_handler sets the handler to receive updates about discovered services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_browse_results_changed_handler(_:_:)
func Nw_browser_set_browse_results_changed_handler(browser Nw_browser_t, handler Nw_browser_browse_results_changed_handler_t) {
	if _nw_browser_set_browse_results_changed_handler == nil {
		panic("Network: symbol nw_browser_set_browse_results_changed_handler not loaded")
	}
	_nw_browser_set_browse_results_changed_handler(browser, handler)
}

var _nw_browser_set_queue func(browser Nw_browser_t, queue uintptr)

// Nw_browser_set_queue sets the queue on which all browser events will be delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_queue(_:_:)
func Nw_browser_set_queue(browser Nw_browser_t, queue dispatch.Queue) {
	if _nw_browser_set_queue == nil {
		panic("Network: symbol nw_browser_set_queue not loaded")
	}
	_nw_browser_set_queue(browser, uintptr(queue.Handle()))
}

var _nw_browser_set_state_changed_handler func(browser Nw_browser_t, state_changed_handler Nw_browser_state_changed_handler_t)

// Nw_browser_set_state_changed_handler sets a handler to receive browser state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_state_changed_handler(_:_:)
func Nw_browser_set_state_changed_handler(browser Nw_browser_t, state_changed_handler Nw_browser_state_changed_handler_t) {
	if _nw_browser_set_state_changed_handler == nil {
		panic("Network: symbol nw_browser_set_state_changed_handler not loaded")
	}
	_nw_browser_set_state_changed_handler(browser, state_changed_handler)
}

var _nw_browser_start func(browser Nw_browser_t)

// Nw_browser_start starts browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_start(_:)
func Nw_browser_start(browser Nw_browser_t) {
	if _nw_browser_start == nil {
		panic("Network: symbol nw_browser_start not loaded")
	}
	_nw_browser_start(browser)
}

var _nw_connection_access_establishment_report func(connection Nw_connection_t, queue uintptr, access_block Nw_establishment_report_access_block_t)

// Nw_connection_access_establishment_report requests a copy of the connection’s establishment report once the connection is in the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_access_establishment_report(_:_:_:)
func Nw_connection_access_establishment_report(connection Nw_connection_t, queue dispatch.Queue, access_block Nw_establishment_report_access_block_t) {
	if _nw_connection_access_establishment_report == nil {
		panic("Network: symbol nw_connection_access_establishment_report not loaded")
	}
	_nw_connection_access_establishment_report(connection, uintptr(queue.Handle()), access_block)
}

var _nw_connection_batch func(connection Nw_connection_t, batch_block unsafe.Pointer)

// Nw_connection_batch defines a block in which calls to send and receive are processed as a batch to improve performance.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_batch(_:_:)
func Nw_connection_batch(connection Nw_connection_t, batch_block unsafe.Pointer) {
	if _nw_connection_batch == nil {
		panic("Network: symbol nw_connection_batch not loaded")
	}
	_nw_connection_batch(connection, batch_block)
}

var _nw_connection_cancel func(connection Nw_connection_t)

// Nw_connection_cancel cancels the connection and gracefully disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel(_:)
func Nw_connection_cancel(connection Nw_connection_t) {
	if _nw_connection_cancel == nil {
		panic("Network: symbol nw_connection_cancel not loaded")
	}
	_nw_connection_cancel(connection)
}

var _nw_connection_cancel_current_endpoint func(connection Nw_connection_t)

// Nw_connection_cancel_current_endpoint causes the current endpoint to be rejected, allowing the connection to try another resolved address.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel_current_endpoint(_:)
func Nw_connection_cancel_current_endpoint(connection Nw_connection_t) {
	if _nw_connection_cancel_current_endpoint == nil {
		panic("Network: symbol nw_connection_cancel_current_endpoint not loaded")
	}
	_nw_connection_cancel_current_endpoint(connection)
}

var _nw_connection_copy_current_path func(connection Nw_connection_t) Nw_path_t

// Nw_connection_copy_current_path accesses the network path the connection is using.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_current_path(_:)
func Nw_connection_copy_current_path(connection Nw_connection_t) Nw_path_t {
	if _nw_connection_copy_current_path == nil {
		panic("Network: symbol nw_connection_copy_current_path not loaded")
	}
	return _nw_connection_copy_current_path(connection)
}

var _nw_connection_copy_description func(connection Nw_connection_t) *byte

// Nw_connection_copy_description copies the description of the connection as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_description(_:)
func Nw_connection_copy_description(connection Nw_connection_t) *byte {
	if _nw_connection_copy_description == nil {
		panic("Network: symbol nw_connection_copy_description not loaded")
	}
	return _nw_connection_copy_description(connection)
}

var _nw_connection_copy_endpoint func(connection Nw_connection_t) Nw_endpoint_t

// Nw_connection_copy_endpoint accesses the endpoint with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_endpoint(_:)
func Nw_connection_copy_endpoint(connection Nw_connection_t) Nw_endpoint_t {
	if _nw_connection_copy_endpoint == nil {
		panic("Network: symbol nw_connection_copy_endpoint not loaded")
	}
	return _nw_connection_copy_endpoint(connection)
}

var _nw_connection_copy_parameters func(connection Nw_connection_t) Nw_parameters_t

// Nw_connection_copy_parameters accesses the parameters with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_parameters(_:)
func Nw_connection_copy_parameters(connection Nw_connection_t) Nw_parameters_t {
	if _nw_connection_copy_parameters == nil {
		panic("Network: symbol nw_connection_copy_parameters not loaded")
	}
	return _nw_connection_copy_parameters(connection)
}

var _nw_connection_copy_protocol_metadata func(connection Nw_connection_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t

// Nw_connection_copy_protocol_metadata retrieves the connection-wide metadata for a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_protocol_metadata(_:_:)
func Nw_connection_copy_protocol_metadata(connection Nw_connection_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	if _nw_connection_copy_protocol_metadata == nil {
		panic("Network: symbol nw_connection_copy_protocol_metadata not loaded")
	}
	return _nw_connection_copy_protocol_metadata(connection, definition)
}

var _nw_connection_create func(endpoint Nw_endpoint_t, parameters Nw_parameters_t) Nw_connection_t

// Nw_connection_create initializes a new connection to a remote endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create(_:_:)
func Nw_connection_create(endpoint Nw_endpoint_t, parameters Nw_parameters_t) Nw_connection_t {
	if _nw_connection_create == nil {
		panic("Network: symbol nw_connection_create not loaded")
	}
	return _nw_connection_create(endpoint, parameters)
}

var _nw_connection_create_new_data_transfer_report func(connection Nw_connection_t) Nw_data_transfer_report_t

// Nw_connection_create_new_data_transfer_report begins a new data transfer report, which can later be collected.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create_new_data_transfer_report(_:)
func Nw_connection_create_new_data_transfer_report(connection Nw_connection_t) Nw_data_transfer_report_t {
	if _nw_connection_create_new_data_transfer_report == nil {
		panic("Network: symbol nw_connection_create_new_data_transfer_report not loaded")
	}
	return _nw_connection_create_new_data_transfer_report(connection)
}

var _nw_connection_force_cancel func(connection Nw_connection_t)

// Nw_connection_force_cancel cancels the connection and immediately disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_force_cancel(_:)
func Nw_connection_force_cancel(connection Nw_connection_t) {
	if _nw_connection_force_cancel == nil {
		panic("Network: symbol nw_connection_force_cancel not loaded")
	}
	_nw_connection_force_cancel(connection)
}

var _nw_connection_get_maximum_datagram_size func(connection Nw_connection_t) uint32

// Nw_connection_get_maximum_datagram_size accesses the maximum size of a datagram message that can be sent on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_get_maximum_datagram_size(_:)
func Nw_connection_get_maximum_datagram_size(connection Nw_connection_t) uint32 {
	if _nw_connection_get_maximum_datagram_size == nil {
		panic("Network: symbol nw_connection_get_maximum_datagram_size not loaded")
	}
	return _nw_connection_get_maximum_datagram_size(connection)
}

var _nw_connection_group_cancel func(group Nw_connection_group_t)

// Nw_connection_group_cancel cancels the connection group object and leaves the network group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_cancel(_:)
func Nw_connection_group_cancel(group Nw_connection_group_t) {
	if _nw_connection_group_cancel == nil {
		panic("Network: symbol nw_connection_group_cancel not loaded")
	}
	_nw_connection_group_cancel(group)
}

var _nw_connection_group_copy_descriptor func(group Nw_connection_group_t) Nw_group_descriptor_t

// Nw_connection_group_copy_descriptor accesses the descriptor of the group you use to initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_descriptor(_:)
func Nw_connection_group_copy_descriptor(group Nw_connection_group_t) Nw_group_descriptor_t {
	if _nw_connection_group_copy_descriptor == nil {
		panic("Network: symbol nw_connection_group_copy_descriptor not loaded")
	}
	return _nw_connection_group_copy_descriptor(group)
}

var _nw_connection_group_copy_local_endpoint_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t

// Nw_connection_group_copy_local_endpoint_for_message accesses the local address and port you use to receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_local_endpoint_for_message(_:_:)
func Nw_connection_group_copy_local_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t {
	if _nw_connection_group_copy_local_endpoint_for_message == nil {
		panic("Network: symbol nw_connection_group_copy_local_endpoint_for_message not loaded")
	}
	return _nw_connection_group_copy_local_endpoint_for_message(group, context)
}

var _nw_connection_group_copy_parameters func(group Nw_connection_group_t) Nw_parameters_t

// Nw_connection_group_copy_parameters accesses the parameters with which you initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_parameters(_:)
func Nw_connection_group_copy_parameters(group Nw_connection_group_t) Nw_parameters_t {
	if _nw_connection_group_copy_parameters == nil {
		panic("Network: symbol nw_connection_group_copy_parameters not loaded")
	}
	return _nw_connection_group_copy_parameters(group)
}

var _nw_connection_group_copy_path_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_path_t

// Nw_connection_group_copy_path_for_message accesses the network path on which you receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_path_for_message(_:_:)
func Nw_connection_group_copy_path_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_path_t {
	if _nw_connection_group_copy_path_for_message == nil {
		panic("Network: symbol nw_connection_group_copy_path_for_message not loaded")
	}
	return _nw_connection_group_copy_path_for_message(group, context)
}

var _nw_connection_group_copy_protocol_metadata func(group Nw_connection_group_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t

// Nw_connection_group_copy_protocol_metadata.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata(_:_:)
func Nw_connection_group_copy_protocol_metadata(group Nw_connection_group_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	if _nw_connection_group_copy_protocol_metadata == nil {
		panic("Network: symbol nw_connection_group_copy_protocol_metadata not loaded")
	}
	return _nw_connection_group_copy_protocol_metadata(group, definition)
}

var _nw_connection_group_copy_protocol_metadata_for_message func(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t

// Nw_connection_group_copy_protocol_metadata_for_message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata_for_message(_:_:_:)
func Nw_connection_group_copy_protocol_metadata_for_message(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	if _nw_connection_group_copy_protocol_metadata_for_message == nil {
		panic("Network: symbol nw_connection_group_copy_protocol_metadata_for_message not loaded")
	}
	return _nw_connection_group_copy_protocol_metadata_for_message(group, context, definition)
}

var _nw_connection_group_copy_remote_endpoint_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t

// Nw_connection_group_copy_remote_endpoint_for_message accesses the endpoint that originates the message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_remote_endpoint_for_message(_:_:)
func Nw_connection_group_copy_remote_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t {
	if _nw_connection_group_copy_remote_endpoint_for_message == nil {
		panic("Network: symbol nw_connection_group_copy_remote_endpoint_for_message not loaded")
	}
	return _nw_connection_group_copy_remote_endpoint_for_message(group, context)
}

var _nw_connection_group_create func(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) Nw_connection_group_t

// Nw_connection_group_create initializes a new connection group with a group identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_create(_:_:)
func Nw_connection_group_create(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) Nw_connection_group_t {
	if _nw_connection_group_create == nil {
		panic("Network: symbol nw_connection_group_create not loaded")
	}
	return _nw_connection_group_create(group_descriptor, parameters)
}

var _nw_connection_group_extract_connection func(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) Nw_connection_t

// Nw_connection_group_extract_connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection(_:_:_:)
func Nw_connection_group_extract_connection(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) Nw_connection_t {
	if _nw_connection_group_extract_connection == nil {
		panic("Network: symbol nw_connection_group_extract_connection not loaded")
	}
	return _nw_connection_group_extract_connection(group, endpoint, protocol_options)
}

var _nw_connection_group_extract_connection_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_connection_t

// Nw_connection_group_extract_connection_for_message converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection_for_message(_:_:)
func Nw_connection_group_extract_connection_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_connection_t {
	if _nw_connection_group_extract_connection_for_message == nil {
		panic("Network: symbol nw_connection_group_extract_connection_for_message not loaded")
	}
	return _nw_connection_group_extract_connection_for_message(group, context)
}

var _nw_connection_group_reinsert_extracted_connection func(group Nw_connection_group_t, connection Nw_connection_t) bool

// Nw_connection_group_reinsert_extracted_connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reinsert_extracted_connection(_:_:)
func Nw_connection_group_reinsert_extracted_connection(group Nw_connection_group_t, connection Nw_connection_t) bool {
	if _nw_connection_group_reinsert_extracted_connection == nil {
		panic("Network: symbol nw_connection_group_reinsert_extracted_connection not loaded")
	}
	return _nw_connection_group_reinsert_extracted_connection(group, connection)
}

var _nw_connection_group_reply func(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content uintptr)

// Nw_connection_group_reply sends a reply to the specific endpoint that originates a group message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reply(_:_:_:_:)
func Nw_connection_group_reply(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content dispatch.Data) {
	if _nw_connection_group_reply == nil {
		panic("Network: symbol nw_connection_group_reply not loaded")
	}
	_nw_connection_group_reply(group, inbound_message, outbound_message, uintptr(content.Handle()))
}

var _nw_connection_group_send_message func(group Nw_connection_group_t, content uintptr, endpoint Nw_endpoint_t, context Nw_content_context_t, completion Nw_connection_group_send_completion_t)

// Nw_connection_group_send_message sends data to the entire group, or to a specific member of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_send_message(_:_:_:_:_:)
func Nw_connection_group_send_message(group Nw_connection_group_t, content dispatch.Data, endpoint Nw_endpoint_t, context Nw_content_context_t, completion Nw_connection_group_send_completion_t) {
	if _nw_connection_group_send_message == nil {
		panic("Network: symbol nw_connection_group_send_message not loaded")
	}
	_nw_connection_group_send_message(group, uintptr(content.Handle()), endpoint, context, completion)
}

var _nw_connection_group_set_new_connection_handler func(group Nw_connection_group_t, new_connection_handler Nw_connection_group_new_connection_handler_t)

// Nw_connection_group_set_new_connection_handler.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_new_connection_handler(_:_:)
func Nw_connection_group_set_new_connection_handler(group Nw_connection_group_t, new_connection_handler Nw_connection_group_new_connection_handler_t) {
	if _nw_connection_group_set_new_connection_handler == nil {
		panic("Network: symbol nw_connection_group_set_new_connection_handler not loaded")
	}
	_nw_connection_group_set_new_connection_handler(group, new_connection_handler)
}

var _nw_connection_group_set_queue func(group Nw_connection_group_t, queue uintptr)

// Nw_connection_group_set_queue sets the queue on which you handle connection group events.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_queue(_:_:)
func Nw_connection_group_set_queue(group Nw_connection_group_t, queue dispatch.Queue) {
	if _nw_connection_group_set_queue == nil {
		panic("Network: symbol nw_connection_group_set_queue not loaded")
	}
	_nw_connection_group_set_queue(group, uintptr(queue.Handle()))
}

var _nw_connection_group_set_receive_handler func(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler Nw_connection_group_receive_handler_t)

// Nw_connection_group_set_receive_handler sets a handler that receives inbound messages from members of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_receive_handler(_:_:_:_:)
func Nw_connection_group_set_receive_handler(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler Nw_connection_group_receive_handler_t) {
	if _nw_connection_group_set_receive_handler == nil {
		panic("Network: symbol nw_connection_group_set_receive_handler not loaded")
	}
	_nw_connection_group_set_receive_handler(group, maximum_message_size, reject_oversized_messages, receive_handler)
}

var _nw_connection_group_set_state_changed_handler func(group Nw_connection_group_t, state_changed_handler Nw_connection_group_state_changed_handler_t)

// Nw_connection_group_set_state_changed_handler sets a handler that receives connection group state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_state_changed_handler(_:_:)
func Nw_connection_group_set_state_changed_handler(group Nw_connection_group_t, state_changed_handler Nw_connection_group_state_changed_handler_t) {
	if _nw_connection_group_set_state_changed_handler == nil {
		panic("Network: symbol nw_connection_group_set_state_changed_handler not loaded")
	}
	_nw_connection_group_set_state_changed_handler(group, state_changed_handler)
}

var _nw_connection_group_start func(group Nw_connection_group_t)

// Nw_connection_group_start joins the group and registers to receive messages.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_start(_:)
func Nw_connection_group_start(group Nw_connection_group_t) {
	if _nw_connection_group_start == nil {
		panic("Network: symbol nw_connection_group_start not loaded")
	}
	_nw_connection_group_start(group)
}

var _nw_connection_receive func(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion Nw_connection_receive_completion_t)

// Nw_connection_receive schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive(_:_:_:_:)
func Nw_connection_receive(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion Nw_connection_receive_completion_t) {
	if _nw_connection_receive == nil {
		panic("Network: symbol nw_connection_receive not loaded")
	}
	_nw_connection_receive(connection, minimum_incomplete_length, maximum_length, completion)
}

var _nw_connection_receive_message func(connection Nw_connection_t, completion Nw_connection_receive_completion_t)

// Nw_connection_receive_message schedules a single receive completion handler for a complete message, as opposed to a range of bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive_message(_:_:)
func Nw_connection_receive_message(connection Nw_connection_t, completion Nw_connection_receive_completion_t) {
	if _nw_connection_receive_message == nil {
		panic("Network: symbol nw_connection_receive_message not loaded")
	}
	_nw_connection_receive_message(connection, completion)
}

var _nw_connection_restart func(connection Nw_connection_t)

// Nw_connection_restart restarts a connection that is in the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_restart(_:)
func Nw_connection_restart(connection Nw_connection_t) {
	if _nw_connection_restart == nil {
		panic("Network: symbol nw_connection_restart not loaded")
	}
	_nw_connection_restart(connection)
}

var _nw_connection_send func(connection Nw_connection_t, content uintptr, context Nw_content_context_t, is_complete bool, completion Nw_connection_send_completion_t)

// Nw_connection_send sends data on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_send(_:_:_:_:_:)
func Nw_connection_send(connection Nw_connection_t, content dispatch.Data, context Nw_content_context_t, is_complete bool, completion Nw_connection_send_completion_t) {
	if _nw_connection_send == nil {
		panic("Network: symbol nw_connection_send not loaded")
	}
	_nw_connection_send(connection, uintptr(content.Handle()), context, is_complete, completion)
}

var _nw_connection_set_better_path_available_handler func(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t)

// Nw_connection_set_better_path_available_handler sets a handler that receives updates when an alternative network path is preferred over the current path.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_better_path_available_handler(_:_:)
func Nw_connection_set_better_path_available_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) {
	if _nw_connection_set_better_path_available_handler == nil {
		panic("Network: symbol nw_connection_set_better_path_available_handler not loaded")
	}
	_nw_connection_set_better_path_available_handler(connection, handler)
}

var _nw_connection_set_path_changed_handler func(connection Nw_connection_t, handler Nw_connection_path_event_handler_t)

// Nw_connection_set_path_changed_handler sets a handler that receives network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_path_changed_handler(_:_:)
func Nw_connection_set_path_changed_handler(connection Nw_connection_t, handler Nw_connection_path_event_handler_t) {
	if _nw_connection_set_path_changed_handler == nil {
		panic("Network: symbol nw_connection_set_path_changed_handler not loaded")
	}
	_nw_connection_set_path_changed_handler(connection, handler)
}

var _nw_connection_set_queue func(connection Nw_connection_t, queue uintptr)

// Nw_connection_set_queue sets the queue on which all connection events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_queue(_:_:)
func Nw_connection_set_queue(connection Nw_connection_t, queue dispatch.Queue) {
	if _nw_connection_set_queue == nil {
		panic("Network: symbol nw_connection_set_queue not loaded")
	}
	_nw_connection_set_queue(connection, uintptr(queue.Handle()))
}

var _nw_connection_set_state_changed_handler func(connection Nw_connection_t, handler Nw_connection_state_changed_handler_t)

// Nw_connection_set_state_changed_handler sets a handler to receive connection state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_state_changed_handler(_:_:)
func Nw_connection_set_state_changed_handler(connection Nw_connection_t, handler Nw_connection_state_changed_handler_t) {
	if _nw_connection_set_state_changed_handler == nil {
		panic("Network: symbol nw_connection_set_state_changed_handler not loaded")
	}
	_nw_connection_set_state_changed_handler(connection, handler)
}

var _nw_connection_set_viability_changed_handler func(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t)

// Nw_connection_set_viability_changed_handler sets a handler that receives updates when data can be sent and received.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_viability_changed_handler(_:_:)
func Nw_connection_set_viability_changed_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) {
	if _nw_connection_set_viability_changed_handler == nil {
		panic("Network: symbol nw_connection_set_viability_changed_handler not loaded")
	}
	_nw_connection_set_viability_changed_handler(connection, handler)
}

var _nw_connection_start func(connection Nw_connection_t)

// Nw_connection_start starts establishing a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_start(_:)
func Nw_connection_start(connection Nw_connection_t) {
	if _nw_connection_start == nil {
		panic("Network: symbol nw_connection_start not loaded")
	}
	_nw_connection_start(connection)
}

var _nw_content_context_copy_antecedent func(context Nw_content_context_t) Nw_content_context_t

// Nw_content_context_copy_antecedent accesses the optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_antecedent(_:)
func Nw_content_context_copy_antecedent(context Nw_content_context_t) Nw_content_context_t {
	if _nw_content_context_copy_antecedent == nil {
		panic("Network: symbol nw_content_context_copy_antecedent not loaded")
	}
	return _nw_content_context_copy_antecedent(context)
}

var _nw_content_context_copy_protocol_metadata func(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) Nw_protocol_metadata_t

// Nw_content_context_copy_protocol_metadata retreives the metadata associated with a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_protocol_metadata(_:_:)
func Nw_content_context_copy_protocol_metadata(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) Nw_protocol_metadata_t {
	if _nw_content_context_copy_protocol_metadata == nil {
		panic("Network: symbol nw_content_context_copy_protocol_metadata not loaded")
	}
	return _nw_content_context_copy_protocol_metadata(context, protocol_)
}

var _nw_content_context_create func(context_identifier string) Nw_content_context_t

// Nw_content_context_create initializes a custom message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_create(_:)
func Nw_content_context_create(context_identifier string) Nw_content_context_t {
	if _nw_content_context_create == nil {
		panic("Network: symbol nw_content_context_create not loaded")
	}
	return _nw_content_context_create(context_identifier)
}

var _nw_content_context_foreach_protocol_metadata func(context Nw_content_context_t)

// Nw_content_context_foreach_protocol_metadata iterates through all protocol metadata associated with the message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_foreach_protocol_metadata(_:_:)
func Nw_content_context_foreach_protocol_metadata(context Nw_content_context_t) {
	if _nw_content_context_foreach_protocol_metadata == nil {
		panic("Network: symbol nw_content_context_foreach_protocol_metadata not loaded")
	}
	_nw_content_context_foreach_protocol_metadata(context)
}

var _nw_content_context_get_expiration_milliseconds func(context Nw_content_context_t) uint64

// Nw_content_context_get_expiration_milliseconds accesses the expiration set for this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_expiration_milliseconds(_:)
func Nw_content_context_get_expiration_milliseconds(context Nw_content_context_t) uint64 {
	if _nw_content_context_get_expiration_milliseconds == nil {
		panic("Network: symbol nw_content_context_get_expiration_milliseconds not loaded")
	}
	return _nw_content_context_get_expiration_milliseconds(context)
}

var _nw_content_context_get_identifier func(context Nw_content_context_t) *byte

// Nw_content_context_get_identifier accesses the identifier used to create this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_identifier(_:)
func Nw_content_context_get_identifier(context Nw_content_context_t) *byte {
	if _nw_content_context_get_identifier == nil {
		panic("Network: symbol nw_content_context_get_identifier not loaded")
	}
	return _nw_content_context_get_identifier(context)
}

var _nw_content_context_get_is_final func(context Nw_content_context_t) bool

// Nw_content_context_get_is_final checks whether this context represents the final message being received.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_is_final(_:)
func Nw_content_context_get_is_final(context Nw_content_context_t) bool {
	if _nw_content_context_get_is_final == nil {
		panic("Network: symbol nw_content_context_get_is_final not loaded")
	}
	return _nw_content_context_get_is_final(context)
}

var _nw_content_context_get_relative_priority func(context Nw_content_context_t) float64

// Nw_content_context_get_relative_priority accesses the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_relative_priority(_:)
func Nw_content_context_get_relative_priority(context Nw_content_context_t) float64 {
	if _nw_content_context_get_relative_priority == nil {
		panic("Network: symbol nw_content_context_get_relative_priority not loaded")
	}
	return _nw_content_context_get_relative_priority(context)
}

var _nw_content_context_set_antecedent func(context Nw_content_context_t, antecedent_context Nw_content_context_t)

// Nw_content_context_set_antecedent set an optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_antecedent(_:_:)
func Nw_content_context_set_antecedent(context Nw_content_context_t, antecedent_context Nw_content_context_t) {
	if _nw_content_context_set_antecedent == nil {
		panic("Network: symbol nw_content_context_set_antecedent not loaded")
	}
	_nw_content_context_set_antecedent(context, antecedent_context)
}

var _nw_content_context_set_expiration_milliseconds func(context Nw_content_context_t, expiration_milliseconds uint64)

// Nw_content_context_set_expiration_milliseconds sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_expiration_milliseconds(_:_:)
func Nw_content_context_set_expiration_milliseconds(context Nw_content_context_t, expiration_milliseconds uint64) {
	if _nw_content_context_set_expiration_milliseconds == nil {
		panic("Network: symbol nw_content_context_set_expiration_milliseconds not loaded")
	}
	_nw_content_context_set_expiration_milliseconds(context, expiration_milliseconds)
}

var _nw_content_context_set_is_final func(context Nw_content_context_t, is_final bool)

// Nw_content_context_set_is_final sets a Boolean indicating if this context represents the final message being sent.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_is_final(_:_:)
func Nw_content_context_set_is_final(context Nw_content_context_t, is_final bool) {
	if _nw_content_context_set_is_final == nil {
		panic("Network: symbol nw_content_context_set_is_final not loaded")
	}
	_nw_content_context_set_is_final(context, is_final)
}

var _nw_content_context_set_metadata_for_protocol func(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t)

// Nw_content_context_set_metadata_for_protocol sets protocol metadata to configure per-message or per-packet properties.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_metadata_for_protocol(_:_:)
func Nw_content_context_set_metadata_for_protocol(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t) {
	if _nw_content_context_set_metadata_for_protocol == nil {
		panic("Network: symbol nw_content_context_set_metadata_for_protocol not loaded")
	}
	_nw_content_context_set_metadata_for_protocol(context, protocol_metadata)
}

var _nw_content_context_set_relative_priority func(context Nw_content_context_t, relative_priority float64)

// Nw_content_context_set_relative_priority sets the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_relative_priority(_:_:)
func Nw_content_context_set_relative_priority(context Nw_content_context_t, relative_priority float64) {
	if _nw_content_context_set_relative_priority == nil {
		panic("Network: symbol nw_content_context_set_relative_priority not loaded")
	}
	_nw_content_context_set_relative_priority(context, relative_priority)
}

var _nw_data_transfer_report_collect func(report Nw_data_transfer_report_t, queue uintptr, collect_block Nw_data_transfer_report_collect_block_t)

// Nw_data_transfer_report_collect stops an outstanding data transfer report and calculates the results.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect(_:_:_:)
func Nw_data_transfer_report_collect(report Nw_data_transfer_report_t, queue dispatch.Queue, collect_block Nw_data_transfer_report_collect_block_t) {
	if _nw_data_transfer_report_collect == nil {
		panic("Network: symbol nw_data_transfer_report_collect not loaded")
	}
	_nw_data_transfer_report_collect(report, uintptr(queue.Handle()), collect_block)
}

var _nw_data_transfer_report_copy_path_interface func(report Nw_data_transfer_report_t, path_index uint32) Nw_interface_t

// Nw_data_transfer_report_copy_path_interface accesses the network interface the path used.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_copy_path_interface(_:_:)
func Nw_data_transfer_report_copy_path_interface(report Nw_data_transfer_report_t, path_index uint32) Nw_interface_t {
	if _nw_data_transfer_report_copy_path_interface == nil {
		panic("Network: symbol nw_data_transfer_report_copy_path_interface not loaded")
	}
	return _nw_data_transfer_report_copy_path_interface(report, path_index)
}

var _nw_data_transfer_report_get_duration_milliseconds func(report Nw_data_transfer_report_t) uint64

// Nw_data_transfer_report_get_duration_milliseconds checks the duration of the data transfer report, from when it was started to when it was collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_duration_milliseconds(_:)
func Nw_data_transfer_report_get_duration_milliseconds(report Nw_data_transfer_report_t) uint64 {
	if _nw_data_transfer_report_get_duration_milliseconds == nil {
		panic("Network: symbol nw_data_transfer_report_get_duration_milliseconds not loaded")
	}
	return _nw_data_transfer_report_get_duration_milliseconds(report)
}

var _nw_data_transfer_report_get_path_count func(report Nw_data_transfer_report_t) uint32

// Nw_data_transfer_report_get_path_count checks the number of valid paths in the report.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_count(_:)
func Nw_data_transfer_report_get_path_count(report Nw_data_transfer_report_t) uint32 {
	if _nw_data_transfer_report_get_path_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_path_count not loaded")
	}
	return _nw_data_transfer_report_get_path_count(report)
}

var _nw_data_transfer_report_get_path_radio_type func(report Nw_data_transfer_report_t, path_index uint32) unsafe.Pointer

// Nw_data_transfer_report_get_path_radio_type.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_radio_type(_:_:)
func Nw_data_transfer_report_get_path_radio_type(report Nw_data_transfer_report_t, path_index uint32) unsafe.Pointer {
	if _nw_data_transfer_report_get_path_radio_type == nil {
		panic("Network: symbol nw_data_transfer_report_get_path_radio_type not loaded")
	}
	return _nw_data_transfer_report_get_path_radio_type(report, path_index)
}

var _nw_data_transfer_report_get_received_application_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_received_application_byte_count accesses the number of bytes the connection delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_application_byte_count(_:_:)
func Nw_data_transfer_report_get_received_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_received_application_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_received_application_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_received_application_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_received_ip_packet_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_received_ip_packet_count accesses the number of IP packets the connection received.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_ip_packet_count(_:_:)
func Nw_data_transfer_report_get_received_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_received_ip_packet_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_received_ip_packet_count not loaded")
	}
	return _nw_data_transfer_report_get_received_ip_packet_count(report, path_index)
}

var _nw_data_transfer_report_get_received_transport_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_received_transport_byte_count accesses the number of bytes the transport protocol delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_received_transport_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_received_transport_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_received_transport_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_received_transport_duplicate_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_received_transport_duplicate_byte_count accesses the number of duplicated bytes the transport protocol detected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_duplicate_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_duplicate_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_received_transport_duplicate_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_received_transport_duplicate_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_received_transport_duplicate_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_received_transport_out_of_order_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_received_transport_out_of_order_byte_count accesses the number of bytes the transport protocol received out of order.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_out_of_order_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_received_transport_out_of_order_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_received_transport_out_of_order_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_sent_application_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_sent_application_byte_count accesses the number of bytes sent on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_application_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_sent_application_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_sent_application_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_sent_application_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_sent_ip_packet_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_sent_ip_packet_count accesses the number of IP packets the connection sent.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_ip_packet_count(_:_:)
func Nw_data_transfer_report_get_sent_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_sent_ip_packet_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_sent_ip_packet_count not loaded")
	}
	return _nw_data_transfer_report_get_sent_ip_packet_count(report, path_index)
}

var _nw_data_transfer_report_get_sent_transport_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_sent_transport_byte_count accesses the number of bytes sent into the transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_sent_transport_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_sent_transport_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_sent_transport_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_sent_transport_retransmitted_byte_count accesses the number of bytes the transport protocol retransmitted.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count == nil {
		panic("Network: symbol nw_data_transfer_report_get_sent_transport_retransmitted_byte_count not loaded")
	}
	return _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report, path_index)
}

var _nw_data_transfer_report_get_state func(report Nw_data_transfer_report_t) unsafe.Pointer

// Nw_data_transfer_report_get_state checks whether a data transfer report is collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_state(_:)
func Nw_data_transfer_report_get_state(report Nw_data_transfer_report_t) unsafe.Pointer {
	if _nw_data_transfer_report_get_state == nil {
		panic("Network: symbol nw_data_transfer_report_get_state not loaded")
	}
	return _nw_data_transfer_report_get_state(report)
}

var _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_transport_minimum_rtt_milliseconds accesses the minimum round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(_:_:)
func Nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds == nil {
		panic("Network: symbol nw_data_transfer_report_get_transport_minimum_rtt_milliseconds not loaded")
	}
	return _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report, path_index)
}

var _nw_data_transfer_report_get_transport_rtt_variance func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_transport_rtt_variance accesses the variance of the round-trip time the transport protocol measured.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_rtt_variance(_:_:)
func Nw_data_transfer_report_get_transport_rtt_variance(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_transport_rtt_variance == nil {
		panic("Network: symbol nw_data_transfer_report_get_transport_rtt_variance not loaded")
	}
	return _nw_data_transfer_report_get_transport_rtt_variance(report, path_index)
}

var _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds func(report Nw_data_transfer_report_t, path_index uint32) uint64

// Nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds accesses the smoothed round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(_:_:)
func Nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	if _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds == nil {
		panic("Network: symbol nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds not loaded")
	}
	return _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report, path_index)
}

var _nw_endpoint_copy_address_string func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_copy_address_string copies the address of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_address_string(_:)
func Nw_endpoint_copy_address_string(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_copy_address_string == nil {
		panic("Network: symbol nw_endpoint_copy_address_string not loaded")
	}
	return _nw_endpoint_copy_address_string(endpoint)
}

var _nw_endpoint_copy_port_string func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_copy_port_string copies the port of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_port_string(_:)
func Nw_endpoint_copy_port_string(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_copy_port_string == nil {
		panic("Network: symbol nw_endpoint_copy_port_string not loaded")
	}
	return _nw_endpoint_copy_port_string(endpoint)
}

var _nw_endpoint_copy_txt_record func(endpoint Nw_endpoint_t) Nw_txt_record_t

// Nw_endpoint_copy_txt_record.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_txt_record(_:)
func Nw_endpoint_copy_txt_record(endpoint Nw_endpoint_t) Nw_txt_record_t {
	if _nw_endpoint_copy_txt_record == nil {
		panic("Network: symbol nw_endpoint_copy_txt_record not loaded")
	}
	return _nw_endpoint_copy_txt_record(endpoint)
}

var _nw_endpoint_create_address func(address uintptr) Nw_endpoint_t

// Nw_endpoint_create_address creates a network endpoint with an address structure.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_address(_:)
func Nw_endpoint_create_address(address uintptr) Nw_endpoint_t {
	if _nw_endpoint_create_address == nil {
		panic("Network: symbol nw_endpoint_create_address not loaded")
	}
	return _nw_endpoint_create_address(address)
}

var _nw_endpoint_create_bonjour_service func(name string, type_ string, domain string) Nw_endpoint_t

// Nw_endpoint_create_bonjour_service creates a network endpoint with a Bonjour service name, type, and domain.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_bonjour_service(_:_:_:)
func Nw_endpoint_create_bonjour_service(name string, type_ string, domain string) Nw_endpoint_t {
	if _nw_endpoint_create_bonjour_service == nil {
		panic("Network: symbol nw_endpoint_create_bonjour_service not loaded")
	}
	return _nw_endpoint_create_bonjour_service(name, type_, domain)
}

var _nw_endpoint_create_host func(hostname string, port string) Nw_endpoint_t

// Nw_endpoint_create_host creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_host(_:_:)
func Nw_endpoint_create_host(hostname string, port string) Nw_endpoint_t {
	if _nw_endpoint_create_host == nil {
		panic("Network: symbol nw_endpoint_create_host not loaded")
	}
	return _nw_endpoint_create_host(hostname, port)
}

var _nw_endpoint_create_url func(url string) Nw_endpoint_t

// Nw_endpoint_create_url creates a network endpoint with a URL string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_url(_:)
func Nw_endpoint_create_url(url string) Nw_endpoint_t {
	if _nw_endpoint_create_url == nil {
		panic("Network: symbol nw_endpoint_create_url not loaded")
	}
	return _nw_endpoint_create_url(url)
}

var _nw_endpoint_get_address func(endpoint Nw_endpoint_t) objectivec.IObject

// Nw_endpoint_get_address accesses the address structure stored in an address endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_address(_:)
func Nw_endpoint_get_address(endpoint Nw_endpoint_t) objectivec.IObject {
	if _nw_endpoint_get_address == nil {
		panic("Network: symbol nw_endpoint_get_address not loaded")
	}
	return _nw_endpoint_get_address(endpoint)
}

var _nw_endpoint_get_bonjour_service_domain func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_get_bonjour_service_domain accesses the Bonjour service domain stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_domain(_:)
func Nw_endpoint_get_bonjour_service_domain(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_get_bonjour_service_domain == nil {
		panic("Network: symbol nw_endpoint_get_bonjour_service_domain not loaded")
	}
	return _nw_endpoint_get_bonjour_service_domain(endpoint)
}

var _nw_endpoint_get_bonjour_service_name func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_get_bonjour_service_name accesses the Bonjour service name stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_name(_:)
func Nw_endpoint_get_bonjour_service_name(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_get_bonjour_service_name == nil {
		panic("Network: symbol nw_endpoint_get_bonjour_service_name not loaded")
	}
	return _nw_endpoint_get_bonjour_service_name(endpoint)
}

var _nw_endpoint_get_bonjour_service_type func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_get_bonjour_service_type accesses the Bonjour service type stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_type(_:)
func Nw_endpoint_get_bonjour_service_type(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_get_bonjour_service_type == nil {
		panic("Network: symbol nw_endpoint_get_bonjour_service_type not loaded")
	}
	return _nw_endpoint_get_bonjour_service_type(endpoint)
}

var _nw_endpoint_get_hostname func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_get_hostname accesses the hostname stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_hostname(_:)
func Nw_endpoint_get_hostname(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_get_hostname == nil {
		panic("Network: symbol nw_endpoint_get_hostname not loaded")
	}
	return _nw_endpoint_get_hostname(endpoint)
}

var _nw_endpoint_get_port func(endpoint Nw_endpoint_t) uint16

// Nw_endpoint_get_port accesses the port stored in an endpoint, in host-byte order.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_port(_:)
func Nw_endpoint_get_port(endpoint Nw_endpoint_t) uint16 {
	if _nw_endpoint_get_port == nil {
		panic("Network: symbol nw_endpoint_get_port not loaded")
	}
	return _nw_endpoint_get_port(endpoint)
}

var _nw_endpoint_get_signature func(endpoint Nw_endpoint_t, out_signature_length *uintptr) *uint8

// Nw_endpoint_get_signature.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_signature(_:_:)
func Nw_endpoint_get_signature(endpoint Nw_endpoint_t, out_signature_length *uintptr) *uint8 {
	if _nw_endpoint_get_signature == nil {
		panic("Network: symbol nw_endpoint_get_signature not loaded")
	}
	return _nw_endpoint_get_signature(endpoint, out_signature_length)
}

var _nw_endpoint_get_type func(endpoint Nw_endpoint_t) unsafe.Pointer

// Nw_endpoint_get_type accesses the type of a endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_type(_:)
func Nw_endpoint_get_type(endpoint Nw_endpoint_t) unsafe.Pointer {
	if _nw_endpoint_get_type == nil {
		panic("Network: symbol nw_endpoint_get_type not loaded")
	}
	return _nw_endpoint_get_type(endpoint)
}

var _nw_endpoint_get_url func(endpoint Nw_endpoint_t) *byte

// Nw_endpoint_get_url accesses the URL string stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_url(_:)
func Nw_endpoint_get_url(endpoint Nw_endpoint_t) *byte {
	if _nw_endpoint_get_url == nil {
		panic("Network: symbol nw_endpoint_get_url not loaded")
	}
	return _nw_endpoint_get_url(endpoint)
}

var _nw_error_copy_cf_error func(err Nw_error_t) corefoundation.CFErrorRef

// Nw_error_copy_cf_error returns a copy of a network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_copy_cf_error(_:)
func Nw_error_copy_cf_error(err Nw_error_t) corefoundation.CFErrorRef {
	if _nw_error_copy_cf_error == nil {
		panic("Network: symbol nw_error_copy_cf_error not loaded")
	}
	return _nw_error_copy_cf_error(err)
}

var _nw_error_get_error_code func(err Nw_error_t) int

// Nw_error_get_error_code accesses the specific code of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_code(_:)
func Nw_error_get_error_code(err Nw_error_t) int {
	if _nw_error_get_error_code == nil {
		panic("Network: symbol nw_error_get_error_code not loaded")
	}
	return _nw_error_get_error_code(err)
}

var _nw_error_get_error_domain func(err Nw_error_t) unsafe.Pointer

// Nw_error_get_error_domain accesses the domain of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_domain(_:)
func Nw_error_get_error_domain(err Nw_error_t) unsafe.Pointer {
	if _nw_error_get_error_domain == nil {
		panic("Network: symbol nw_error_get_error_domain not loaded")
	}
	return _nw_error_get_error_domain(err)
}

var _nw_establishment_report_copy_proxy_endpoint func(report Nw_establishment_report_t) Nw_endpoint_t

// Nw_establishment_report_copy_proxy_endpoint accesses the endpoint of the proxy the connection used.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_copy_proxy_endpoint(_:)
func Nw_establishment_report_copy_proxy_endpoint(report Nw_establishment_report_t) Nw_endpoint_t {
	if _nw_establishment_report_copy_proxy_endpoint == nil {
		panic("Network: symbol nw_establishment_report_copy_proxy_endpoint not loaded")
	}
	return _nw_establishment_report_copy_proxy_endpoint(report)
}

var _nw_establishment_report_enumerate_protocols func(report Nw_establishment_report_t, enumerate_block Nw_report_protocol_enumerator_t)

// Nw_establishment_report_enumerate_protocols iterates a list of protocol handshakes in order from first completed to last completed.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_protocols(_:_:)
func Nw_establishment_report_enumerate_protocols(report Nw_establishment_report_t, enumerate_block Nw_report_protocol_enumerator_t) {
	if _nw_establishment_report_enumerate_protocols == nil {
		panic("Network: symbol nw_establishment_report_enumerate_protocols not loaded")
	}
	_nw_establishment_report_enumerate_protocols(report, enumerate_block)
}

var _nw_establishment_report_enumerate_resolution_reports func(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_report_enumerator_t)

// Nw_establishment_report_enumerate_resolution_reports.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolution_reports(_:_:)
func Nw_establishment_report_enumerate_resolution_reports(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_report_enumerator_t) {
	if _nw_establishment_report_enumerate_resolution_reports == nil {
		panic("Network: symbol nw_establishment_report_enumerate_resolution_reports not loaded")
	}
	_nw_establishment_report_enumerate_resolution_reports(report, enumerate_block)
}

var _nw_establishment_report_enumerate_resolutions func(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_enumerator_t)

// Nw_establishment_report_enumerate_resolutions iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolutions(_:_:)
func Nw_establishment_report_enumerate_resolutions(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_enumerator_t) {
	if _nw_establishment_report_enumerate_resolutions == nil {
		panic("Network: symbol nw_establishment_report_enumerate_resolutions not loaded")
	}
	_nw_establishment_report_enumerate_resolutions(report, enumerate_block)
}

var _nw_establishment_report_get_attempt_started_after_milliseconds func(report Nw_establishment_report_t) uint64

// Nw_establishment_report_get_attempt_started_after_milliseconds accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_attempt_started_after_milliseconds(_:)
func Nw_establishment_report_get_attempt_started_after_milliseconds(report Nw_establishment_report_t) uint64 {
	if _nw_establishment_report_get_attempt_started_after_milliseconds == nil {
		panic("Network: symbol nw_establishment_report_get_attempt_started_after_milliseconds not loaded")
	}
	return _nw_establishment_report_get_attempt_started_after_milliseconds(report)
}

var _nw_establishment_report_get_duration_milliseconds func(report Nw_establishment_report_t) uint64

// Nw_establishment_report_get_duration_milliseconds checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_duration_milliseconds(_:)
func Nw_establishment_report_get_duration_milliseconds(report Nw_establishment_report_t) uint64 {
	if _nw_establishment_report_get_duration_milliseconds == nil {
		panic("Network: symbol nw_establishment_report_get_duration_milliseconds not loaded")
	}
	return _nw_establishment_report_get_duration_milliseconds(report)
}

var _nw_establishment_report_get_previous_attempt_count func(report Nw_establishment_report_t) uint32

// Nw_establishment_report_get_previous_attempt_count checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_previous_attempt_count(_:)
func Nw_establishment_report_get_previous_attempt_count(report Nw_establishment_report_t) uint32 {
	if _nw_establishment_report_get_previous_attempt_count == nil {
		panic("Network: symbol nw_establishment_report_get_previous_attempt_count not loaded")
	}
	return _nw_establishment_report_get_previous_attempt_count(report)
}

var _nw_establishment_report_get_proxy_configured func(report Nw_establishment_report_t) bool

// Nw_establishment_report_get_proxy_configured checks whether a proxy was configured on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_proxy_configured(_:)
func Nw_establishment_report_get_proxy_configured(report Nw_establishment_report_t) bool {
	if _nw_establishment_report_get_proxy_configured == nil {
		panic("Network: symbol nw_establishment_report_get_proxy_configured not loaded")
	}
	return _nw_establishment_report_get_proxy_configured(report)
}

var _nw_establishment_report_get_used_proxy func(report Nw_establishment_report_t) bool

// Nw_establishment_report_get_used_proxy checks whether the connection used a proxy.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_used_proxy(_:)
func Nw_establishment_report_get_used_proxy(report Nw_establishment_report_t) bool {
	if _nw_establishment_report_get_used_proxy == nil {
		panic("Network: symbol nw_establishment_report_get_used_proxy not loaded")
	}
	return _nw_establishment_report_get_used_proxy(report)
}

var _nw_ethernet_channel_cancel func(ethernet_channel Nw_ethernet_channel_t)

// Nw_ethernet_channel_cancel unregisters the channel from the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_cancel(_:)
func Nw_ethernet_channel_cancel(ethernet_channel Nw_ethernet_channel_t) {
	if _nw_ethernet_channel_cancel == nil {
		panic("Network: symbol nw_ethernet_channel_cancel not loaded")
	}
	_nw_ethernet_channel_cancel(ethernet_channel)
}

var _nw_ethernet_channel_create func(ether_type uint16, interface_ Nw_interface_t) Nw_ethernet_channel_t

// Nw_ethernet_channel_create initializes an Ethernet channel on a specific interface with a custom Ethernet type.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create(_:_:)
func Nw_ethernet_channel_create(ether_type uint16, interface_ Nw_interface_t) Nw_ethernet_channel_t {
	if _nw_ethernet_channel_create == nil {
		panic("Network: symbol nw_ethernet_channel_create not loaded")
	}
	return _nw_ethernet_channel_create(ether_type, interface_)
}

var _nw_ethernet_channel_create_with_parameters func(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) Nw_ethernet_channel_t

// Nw_ethernet_channel_create_with_parameters.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create_with_parameters(_:_:_:)
func Nw_ethernet_channel_create_with_parameters(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) Nw_ethernet_channel_t {
	if _nw_ethernet_channel_create_with_parameters == nil {
		panic("Network: symbol nw_ethernet_channel_create_with_parameters not loaded")
	}
	return _nw_ethernet_channel_create_with_parameters(ether_type, interface_, parameters)
}

var _nw_ethernet_channel_get_maximum_payload_size func(ethernet_channel Nw_ethernet_channel_t) uint32

// Nw_ethernet_channel_get_maximum_payload_size.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_get_maximum_payload_size(_:)
func Nw_ethernet_channel_get_maximum_payload_size(ethernet_channel Nw_ethernet_channel_t) uint32 {
	if _nw_ethernet_channel_get_maximum_payload_size == nil {
		panic("Network: symbol nw_ethernet_channel_get_maximum_payload_size not loaded")
	}
	return _nw_ethernet_channel_get_maximum_payload_size(ethernet_channel)
}

var _nw_ethernet_channel_send func(ethernet_channel Nw_ethernet_channel_t, content uintptr, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion Nw_ethernet_channel_send_completion_t)

// Nw_ethernet_channel_send sends a single Ethernet frame over a channel to a specific Ethernet address.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send(_:_:_:_:_:)
func Nw_ethernet_channel_send(ethernet_channel Nw_ethernet_channel_t, content dispatch.Data, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion Nw_ethernet_channel_send_completion_t) {
	if _nw_ethernet_channel_send == nil {
		panic("Network: symbol nw_ethernet_channel_send not loaded")
	}
	_nw_ethernet_channel_send(ethernet_channel, uintptr(content.Handle()), vlan_tag, remote_address, completion)
}

var _nw_ethernet_channel_set_queue func(ethernet_channel Nw_ethernet_channel_t, queue uintptr)

// Nw_ethernet_channel_set_queue sets the queue on which all channel events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_queue(_:_:)
func Nw_ethernet_channel_set_queue(ethernet_channel Nw_ethernet_channel_t, queue dispatch.Queue) {
	if _nw_ethernet_channel_set_queue == nil {
		panic("Network: symbol nw_ethernet_channel_set_queue not loaded")
	}
	_nw_ethernet_channel_set_queue(ethernet_channel, uintptr(queue.Handle()))
}

var _nw_ethernet_channel_set_receive_handler func(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_receive_handler_t)

// Nw_ethernet_channel_set_receive_handler sets a handler to receive inbound Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_receive_handler(_:_:)
func Nw_ethernet_channel_set_receive_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_receive_handler_t) {
	if _nw_ethernet_channel_set_receive_handler == nil {
		panic("Network: symbol nw_ethernet_channel_set_receive_handler not loaded")
	}
	_nw_ethernet_channel_set_receive_handler(ethernet_channel, handler)
}

var _nw_ethernet_channel_set_state_changed_handler func(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_state_changed_handler_t)

// Nw_ethernet_channel_set_state_changed_handler sets a handler to receive channel state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_state_changed_handler(_:_:)
func Nw_ethernet_channel_set_state_changed_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_state_changed_handler_t) {
	if _nw_ethernet_channel_set_state_changed_handler == nil {
		panic("Network: symbol nw_ethernet_channel_set_state_changed_handler not loaded")
	}
	_nw_ethernet_channel_set_state_changed_handler(ethernet_channel, handler)
}

var _nw_ethernet_channel_start func(ethernet_channel Nw_ethernet_channel_t)

// Nw_ethernet_channel_start starts the process of registering the channel.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_start(_:)
func Nw_ethernet_channel_start(ethernet_channel Nw_ethernet_channel_t) {
	if _nw_ethernet_channel_start == nil {
		panic("Network: symbol nw_ethernet_channel_start not loaded")
	}
	_nw_ethernet_channel_start(ethernet_channel)
}

var _nw_framer_async func(framer Nw_framer_t, async_block Nw_framer_block_t)

// Nw_framer_async requests that a block be executed on the connection’s internal scheduling context.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_async(_:_:)
func Nw_framer_async(framer Nw_framer_t, async_block Nw_framer_block_t) {
	if _nw_framer_async == nil {
		panic("Network: symbol nw_framer_async not loaded")
	}
	_nw_framer_async(framer, async_block)
}

var _nw_framer_copy_local_endpoint func(framer Nw_framer_t) Nw_endpoint_t

// Nw_framer_copy_local_endpoint accesses the local endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_local_endpoint(_:)
func Nw_framer_copy_local_endpoint(framer Nw_framer_t) Nw_endpoint_t {
	if _nw_framer_copy_local_endpoint == nil {
		panic("Network: symbol nw_framer_copy_local_endpoint not loaded")
	}
	return _nw_framer_copy_local_endpoint(framer)
}

var _nw_framer_copy_options func(framer Nw_framer_t) Nw_protocol_options_t

// Nw_framer_copy_options.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_options(_:)
func Nw_framer_copy_options(framer Nw_framer_t) Nw_protocol_options_t {
	if _nw_framer_copy_options == nil {
		panic("Network: symbol nw_framer_copy_options not loaded")
	}
	return _nw_framer_copy_options(framer)
}

var _nw_framer_copy_parameters func(framer Nw_framer_t) Nw_parameters_t

// Nw_framer_copy_parameters accesses the parameters of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_parameters(_:)
func Nw_framer_copy_parameters(framer Nw_framer_t) Nw_parameters_t {
	if _nw_framer_copy_parameters == nil {
		panic("Network: symbol nw_framer_copy_parameters not loaded")
	}
	return _nw_framer_copy_parameters(framer)
}

var _nw_framer_copy_remote_endpoint func(framer Nw_framer_t) Nw_endpoint_t

// Nw_framer_copy_remote_endpoint accesses the remote endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_remote_endpoint(_:)
func Nw_framer_copy_remote_endpoint(framer Nw_framer_t) Nw_endpoint_t {
	if _nw_framer_copy_remote_endpoint == nil {
		panic("Network: symbol nw_framer_copy_remote_endpoint not loaded")
	}
	return _nw_framer_copy_remote_endpoint(framer)
}

var _nw_framer_create_definition func(identifier string, flags uint32, start_handler Nw_framer_start_handler_t) Nw_protocol_definition_t

// Nw_framer_create_definition initializes a new protocol definition based on your protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_definition(_:_:_:)
func Nw_framer_create_definition(identifier string, flags uint32, start_handler Nw_framer_start_handler_t) Nw_protocol_definition_t {
	if _nw_framer_create_definition == nil {
		panic("Network: symbol nw_framer_create_definition not loaded")
	}
	return _nw_framer_create_definition(identifier, flags, start_handler)
}

var _nw_framer_create_options func(framer_definition Nw_protocol_definition_t) Nw_protocol_options_t

// Nw_framer_create_options initializes a set of protocol options with a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_options(_:)
func Nw_framer_create_options(framer_definition Nw_protocol_definition_t) Nw_protocol_options_t {
	if _nw_framer_create_options == nil {
		panic("Network: symbol nw_framer_create_options not loaded")
	}
	return _nw_framer_create_options(framer_definition)
}

var _nw_framer_deliver_input func(framer Nw_framer_t, input_buffer *uint8, input_length uintptr, message Nw_framer_message_t, is_complete bool)

// Nw_framer_deliver_input delivers an inbound message containing arbitrary data from your protocol to the application.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input(_:_:_:_:_:)
func Nw_framer_deliver_input(framer Nw_framer_t, input_buffer *uint8, input_length uintptr, message Nw_framer_message_t, is_complete bool) {
	if _nw_framer_deliver_input == nil {
		panic("Network: symbol nw_framer_deliver_input not loaded")
	}
	_nw_framer_deliver_input(framer, input_buffer, input_length, message, is_complete)
}

var _nw_framer_deliver_input_no_copy func(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) bool

// Nw_framer_deliver_input_no_copy delivers an inbound message containing a specific number of next received bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input_no_copy(_:_:_:_:)
func Nw_framer_deliver_input_no_copy(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) bool {
	if _nw_framer_deliver_input_no_copy == nil {
		panic("Network: symbol nw_framer_deliver_input_no_copy not loaded")
	}
	return _nw_framer_deliver_input_no_copy(framer, input_length, message, is_complete)
}

var _nw_framer_mark_failed_with_error func(framer Nw_framer_t, error_code int)

// Nw_framer_mark_failed_with_error indicates to a connection that your protocol has encountered an error, or has gracefully closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_failed_with_error(_:_:)
func Nw_framer_mark_failed_with_error(framer Nw_framer_t, error_code int) {
	if _nw_framer_mark_failed_with_error == nil {
		panic("Network: symbol nw_framer_mark_failed_with_error not loaded")
	}
	_nw_framer_mark_failed_with_error(framer, error_code)
}

var _nw_framer_mark_ready func(framer Nw_framer_t)

// Nw_framer_mark_ready indicates to a connection that your protocol’s handshake is complete.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_ready(_:)
func Nw_framer_mark_ready(framer Nw_framer_t) {
	if _nw_framer_mark_ready == nil {
		panic("Network: symbol nw_framer_mark_ready not loaded")
	}
	_nw_framer_mark_ready(framer)
}

var _nw_framer_message_access_value func(message Nw_framer_message_t, key string, access_value bool) bool

// Nw_framer_message_access_value accesses a custom value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_access_value(_:_:_:)
func Nw_framer_message_access_value(message Nw_framer_message_t, key string, access_value bool) bool {
	if _nw_framer_message_access_value == nil {
		panic("Network: symbol nw_framer_message_access_value not loaded")
	}
	return _nw_framer_message_access_value(message, key, access_value)
}

var _nw_framer_message_copy_object_value func(message Nw_framer_message_t, key string) objectivec.Object

// Nw_framer_message_copy_object_value accesses an NSObject value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_copy_object_value(_:_:)
func Nw_framer_message_copy_object_value(message Nw_framer_message_t, key string) objectivec.Object {
	if _nw_framer_message_copy_object_value == nil {
		panic("Network: symbol nw_framer_message_copy_object_value not loaded")
	}
	return _nw_framer_message_copy_object_value(message, key)
}

var _nw_framer_message_create func(framer Nw_framer_t) Nw_framer_message_t

// Nw_framer_message_create initializes an empty message from within a framer implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_create(_:)
func Nw_framer_message_create(framer Nw_framer_t) Nw_framer_message_t {
	if _nw_framer_message_create == nil {
		panic("Network: symbol nw_framer_message_create not loaded")
	}
	return _nw_framer_message_create(framer)
}

var _nw_framer_message_set_object_value func(message Nw_framer_message_t, key string, value objectivec.Object)

// Nw_framer_message_set_object_value sets an NSObject value to be stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_object_value(_:_:_:)
func Nw_framer_message_set_object_value(message Nw_framer_message_t, key string, value objectivec.Object) {
	if _nw_framer_message_set_object_value == nil {
		panic("Network: symbol nw_framer_message_set_object_value not loaded")
	}
	_nw_framer_message_set_object_value(message, key, value)
}

var _nw_framer_message_set_value func(message Nw_framer_message_t, key string, value unsafe.Pointer, dispose_value Nw_framer_message_dispose_value_t)

// Nw_framer_message_set_value sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_value(_:_:_:_:)
func Nw_framer_message_set_value(message Nw_framer_message_t, key string, value unsafe.Pointer, dispose_value Nw_framer_message_dispose_value_t) {
	if _nw_framer_message_set_value == nil {
		panic("Network: symbol nw_framer_message_set_value not loaded")
	}
	_nw_framer_message_set_value(message, key, value, dispose_value)
}

var _nw_framer_options_copy_object_value func(options Nw_protocol_options_t, key string) objectivec.Object

// Nw_framer_options_copy_object_value.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_copy_object_value(_:_:)
func Nw_framer_options_copy_object_value(options Nw_protocol_options_t, key string) objectivec.Object {
	if _nw_framer_options_copy_object_value == nil {
		panic("Network: symbol nw_framer_options_copy_object_value not loaded")
	}
	return _nw_framer_options_copy_object_value(options, key)
}

var _nw_framer_options_set_object_value func(options Nw_protocol_options_t, key string, value objectivec.Object)

// Nw_framer_options_set_object_value.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_set_object_value(_:_:_:)
func Nw_framer_options_set_object_value(options Nw_protocol_options_t, key string, value objectivec.Object) {
	if _nw_framer_options_set_object_value == nil {
		panic("Network: symbol nw_framer_options_set_object_value not loaded")
	}
	_nw_framer_options_set_object_value(options, key, value)
}

var _nw_framer_parse_input func(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool

// Nw_framer_parse_input examines the content of input data while inside your input handler block.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_input(_:_:_:_:_:)
func Nw_framer_parse_input(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool {
	if _nw_framer_parse_input == nil {
		panic("Network: symbol nw_framer_parse_input not loaded")
	}
	return _nw_framer_parse_input(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
}

var _nw_framer_parse_output func(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool

// Nw_framer_parse_output examines the content of output data while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_output(_:_:_:_:_:)
func Nw_framer_parse_output(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool {
	if _nw_framer_parse_output == nil {
		panic("Network: symbol nw_framer_parse_output not loaded")
	}
	return _nw_framer_parse_output(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
}

var _nw_framer_pass_through_input func(framer Nw_framer_t)

// Nw_framer_pass_through_input indicates that your protocol no longer needs to handle input data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_input(_:)
func Nw_framer_pass_through_input(framer Nw_framer_t) {
	if _nw_framer_pass_through_input == nil {
		panic("Network: symbol nw_framer_pass_through_input not loaded")
	}
	_nw_framer_pass_through_input(framer)
}

var _nw_framer_pass_through_output func(framer Nw_framer_t)

// Nw_framer_pass_through_output indicates that your protocol no longer needs to handle output data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_output(_:)
func Nw_framer_pass_through_output(framer Nw_framer_t) {
	if _nw_framer_pass_through_output == nil {
		panic("Network: symbol nw_framer_pass_through_output not loaded")
	}
	_nw_framer_pass_through_output(framer)
}

var _nw_framer_prepend_application_protocol func(framer Nw_framer_t, protocol_options Nw_protocol_options_t) bool

// Nw_framer_prepend_application_protocol dynamically adds another protocol that will run above your protocol after your protocol calls nw_framer_mark_ready(_:).
//
// See: https://developer.apple.com/documentation/Network/nw_framer_prepend_application_protocol(_:_:)
func Nw_framer_prepend_application_protocol(framer Nw_framer_t, protocol_options Nw_protocol_options_t) bool {
	if _nw_framer_prepend_application_protocol == nil {
		panic("Network: symbol nw_framer_prepend_application_protocol not loaded")
	}
	return _nw_framer_prepend_application_protocol(framer, protocol_options)
}

var _nw_framer_protocol_create_message func(definition Nw_protocol_definition_t) Nw_framer_message_t

// Nw_framer_protocol_create_message initializes an empty message for a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_protocol_create_message(_:)
func Nw_framer_protocol_create_message(definition Nw_protocol_definition_t) Nw_framer_message_t {
	if _nw_framer_protocol_create_message == nil {
		panic("Network: symbol nw_framer_protocol_create_message not loaded")
	}
	return _nw_framer_protocol_create_message(definition)
}

var _nw_framer_schedule_wakeup func(framer Nw_framer_t, milliseconds uint64)

// Nw_framer_schedule_wakeup requests that the nw_framer_wakeup_handler_t be called on your protocol at a specific time in the future.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_schedule_wakeup(_:_:)
func Nw_framer_schedule_wakeup(framer Nw_framer_t, milliseconds uint64) {
	if _nw_framer_schedule_wakeup == nil {
		panic("Network: symbol nw_framer_schedule_wakeup not loaded")
	}
	_nw_framer_schedule_wakeup(framer, milliseconds)
}

var _nw_framer_set_cleanup_handler func(framer Nw_framer_t, cleanup_handler Nw_framer_cleanup_handler_t)

// Nw_framer_set_cleanup_handler sets a block to handle the final cleanup of allocations made by your protocol instance.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_cleanup_handler(_:_:)
func Nw_framer_set_cleanup_handler(framer Nw_framer_t, cleanup_handler Nw_framer_cleanup_handler_t) {
	if _nw_framer_set_cleanup_handler == nil {
		panic("Network: symbol nw_framer_set_cleanup_handler not loaded")
	}
	_nw_framer_set_cleanup_handler(framer, cleanup_handler)
}

var _nw_framer_set_input_handler func(framer Nw_framer_t, input_handler Nw_framer_input_handler_t)

// Nw_framer_set_input_handler sets a block to handle new inbound data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_input_handler(_:_:)
func Nw_framer_set_input_handler(framer Nw_framer_t, input_handler Nw_framer_input_handler_t) {
	if _nw_framer_set_input_handler == nil {
		panic("Network: symbol nw_framer_set_input_handler not loaded")
	}
	_nw_framer_set_input_handler(framer, input_handler)
}

var _nw_framer_set_output_handler func(framer Nw_framer_t, output_handler Nw_framer_output_handler_t)

// Nw_framer_set_output_handler sets a block to handle new outbound messages.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_output_handler(_:_:)
func Nw_framer_set_output_handler(framer Nw_framer_t, output_handler Nw_framer_output_handler_t) {
	if _nw_framer_set_output_handler == nil {
		panic("Network: symbol nw_framer_set_output_handler not loaded")
	}
	_nw_framer_set_output_handler(framer, output_handler)
}

var _nw_framer_set_stop_handler func(framer Nw_framer_t, stop_handler Nw_framer_stop_handler_t)

// Nw_framer_set_stop_handler sets a block to handle when the connection is being closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_stop_handler(_:_:)
func Nw_framer_set_stop_handler(framer Nw_framer_t, stop_handler Nw_framer_stop_handler_t) {
	if _nw_framer_set_stop_handler == nil {
		panic("Network: symbol nw_framer_set_stop_handler not loaded")
	}
	_nw_framer_set_stop_handler(framer, stop_handler)
}

var _nw_framer_set_wakeup_handler func(framer Nw_framer_t, wakeup_handler Nw_framer_wakeup_handler_t)

// Nw_framer_set_wakeup_handler sets a handler to receive scheduled wakeup events.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_wakeup_handler(_:_:)
func Nw_framer_set_wakeup_handler(framer Nw_framer_t, wakeup_handler Nw_framer_wakeup_handler_t) {
	if _nw_framer_set_wakeup_handler == nil {
		panic("Network: symbol nw_framer_set_wakeup_handler not loaded")
	}
	_nw_framer_set_wakeup_handler(framer, wakeup_handler)
}

var _nw_framer_write_output func(framer Nw_framer_t, output_buffer *uint8, output_length uintptr)

// Nw_framer_write_output sends arbitrary output data in a buffer from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output(_:_:_:)
func Nw_framer_write_output(framer Nw_framer_t, output_buffer *uint8, output_length uintptr) {
	if _nw_framer_write_output == nil {
		panic("Network: symbol nw_framer_write_output not loaded")
	}
	_nw_framer_write_output(framer, output_buffer, output_length)
}

var _nw_framer_write_output_data func(framer Nw_framer_t, output_data uintptr)

// Nw_framer_write_output_data sends arbitrary output data from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_data(_:_:)
func Nw_framer_write_output_data(framer Nw_framer_t, output_data dispatch.Data) {
	if _nw_framer_write_output_data == nil {
		panic("Network: symbol nw_framer_write_output_data not loaded")
	}
	_nw_framer_write_output_data(framer, uintptr(output_data.Handle()))
}

var _nw_framer_write_output_no_copy func(framer Nw_framer_t, output_length uintptr) bool

// Nw_framer_write_output_no_copy sends a specific number of bytes from a message while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_no_copy(_:_:)
func Nw_framer_write_output_no_copy(framer Nw_framer_t, output_length uintptr) bool {
	if _nw_framer_write_output_no_copy == nil {
		panic("Network: symbol nw_framer_write_output_no_copy not loaded")
	}
	return _nw_framer_write_output_no_copy(framer, output_length)
}

var _nw_group_descriptor_add_endpoint func(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) bool

// Nw_group_descriptor_add_endpoint adds a multicast address endpoint you specify to define an extra IP multicast group to join.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_add_endpoint(_:_:)
func Nw_group_descriptor_add_endpoint(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) bool {
	if _nw_group_descriptor_add_endpoint == nil {
		panic("Network: symbol nw_group_descriptor_add_endpoint not loaded")
	}
	return _nw_group_descriptor_add_endpoint(descriptor, endpoint)
}

var _nw_group_descriptor_create_multicast func(multicast_group Nw_endpoint_t) Nw_group_descriptor_t

// Nw_group_descriptor_create_multicast creates group descriptor you use to join an IP multicast group on a local network.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multicast(_:)
func Nw_group_descriptor_create_multicast(multicast_group Nw_endpoint_t) Nw_group_descriptor_t {
	if _nw_group_descriptor_create_multicast == nil {
		panic("Network: symbol nw_group_descriptor_create_multicast not loaded")
	}
	return _nw_group_descriptor_create_multicast(multicast_group)
}

var _nw_group_descriptor_create_multiplex func(remote_endpoint Nw_endpoint_t) Nw_group_descriptor_t

// Nw_group_descriptor_create_multiplex.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multiplex(_:)
func Nw_group_descriptor_create_multiplex(remote_endpoint Nw_endpoint_t) Nw_group_descriptor_t {
	if _nw_group_descriptor_create_multiplex == nil {
		panic("Network: symbol nw_group_descriptor_create_multiplex not loaded")
	}
	return _nw_group_descriptor_create_multiplex(remote_endpoint)
}

var _nw_group_descriptor_enumerate_endpoints func(descriptor Nw_group_descriptor_t, enumerate_block Nw_group_descriptor_enumerate_endpoints_block_t)

// Nw_group_descriptor_enumerate_endpoints sets a handler to list all endpoints added to the group descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints(_:_:)
func Nw_group_descriptor_enumerate_endpoints(descriptor Nw_group_descriptor_t, enumerate_block Nw_group_descriptor_enumerate_endpoints_block_t) {
	if _nw_group_descriptor_enumerate_endpoints == nil {
		panic("Network: symbol nw_group_descriptor_enumerate_endpoints not loaded")
	}
	_nw_group_descriptor_enumerate_endpoints(descriptor, enumerate_block)
}

var _nw_interface_get_index func(interface_ Nw_interface_t) uint32

// Nw_interface_get_index accesses the system interface index associated with the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_index(_:)
func Nw_interface_get_index(interface_ Nw_interface_t) uint32 {
	if _nw_interface_get_index == nil {
		panic("Network: symbol nw_interface_get_index not loaded")
	}
	return _nw_interface_get_index(interface_)
}

var _nw_interface_get_name func(interface_ Nw_interface_t) *byte

// Nw_interface_get_name accesses the name of the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_name(_:)
func Nw_interface_get_name(interface_ Nw_interface_t) *byte {
	if _nw_interface_get_name == nil {
		panic("Network: symbol nw_interface_get_name not loaded")
	}
	return _nw_interface_get_name(interface_)
}

var _nw_interface_get_type func(interface_ Nw_interface_t) unsafe.Pointer

// Nw_interface_get_type accesses the type of the interface, such as Wi-Fi or Loopback.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_type(_:)
func Nw_interface_get_type(interface_ Nw_interface_t) unsafe.Pointer {
	if _nw_interface_get_type == nil {
		panic("Network: symbol nw_interface_get_type not loaded")
	}
	return _nw_interface_get_type(interface_)
}

var _nw_ip_create_metadata func() Nw_protocol_metadata_t

// Nw_ip_create_metadata initializes an IP packet configuration with default settings.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_create_metadata()
func Nw_ip_create_metadata() Nw_protocol_metadata_t {
	if _nw_ip_create_metadata == nil {
		panic("Network: symbol nw_ip_create_metadata not loaded")
	}
	return _nw_ip_create_metadata()
}

var _nw_ip_metadata_get_ecn_flag func(metadata Nw_protocol_metadata_t) unsafe.Pointer

// Nw_ip_metadata_get_ecn_flag checks the Explicit Congestion Notification flag value received on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_ecn_flag(_:)
func Nw_ip_metadata_get_ecn_flag(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	if _nw_ip_metadata_get_ecn_flag == nil {
		panic("Network: symbol nw_ip_metadata_get_ecn_flag not loaded")
	}
	return _nw_ip_metadata_get_ecn_flag(metadata)
}

var _nw_ip_metadata_get_receive_time func(metadata Nw_protocol_metadata_t) uint64

// Nw_ip_metadata_get_receive_time access the time at which a packet was received, in nanoseconds, based on `CLOCK_MONOTONIC_RAW`.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_receive_time(_:)
func Nw_ip_metadata_get_receive_time(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_ip_metadata_get_receive_time == nil {
		panic("Network: symbol nw_ip_metadata_get_receive_time not loaded")
	}
	return _nw_ip_metadata_get_receive_time(metadata)
}

var _nw_ip_metadata_get_service_class func(metadata Nw_protocol_metadata_t) unsafe.Pointer

// Nw_ip_metadata_get_service_class accesses a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_service_class(_:)
func Nw_ip_metadata_get_service_class(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	if _nw_ip_metadata_get_service_class == nil {
		panic("Network: symbol nw_ip_metadata_get_service_class not loaded")
	}
	return _nw_ip_metadata_get_service_class(metadata)
}

var _nw_ip_metadata_set_ecn_flag func(metadata Nw_protocol_metadata_t, ecn_flag unsafe.Pointer)

// Nw_ip_metadata_set_ecn_flag sets a specific Explicit Congestion Notification flag value to set on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_ecn_flag(_:_:)
func Nw_ip_metadata_set_ecn_flag(metadata Nw_protocol_metadata_t, ecn_flag unsafe.Pointer) {
	if _nw_ip_metadata_set_ecn_flag == nil {
		panic("Network: symbol nw_ip_metadata_set_ecn_flag not loaded")
	}
	_nw_ip_metadata_set_ecn_flag(metadata, ecn_flag)
}

var _nw_ip_metadata_set_service_class func(metadata Nw_protocol_metadata_t, service_class unsafe.Pointer)

// Nw_ip_metadata_set_service_class sets a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_service_class(_:_:)
func Nw_ip_metadata_set_service_class(metadata Nw_protocol_metadata_t, service_class unsafe.Pointer) {
	if _nw_ip_metadata_set_service_class == nil {
		panic("Network: symbol nw_ip_metadata_set_service_class not loaded")
	}
	_nw_ip_metadata_set_service_class(metadata, service_class)
}

var _nw_ip_options_set_calculate_receive_time func(options Nw_protocol_options_t, calculate_receive_time bool)

// Nw_ip_options_set_calculate_receive_time configures a connection to deliver receive timestamps for IP packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_calculate_receive_time(_:_:)
func Nw_ip_options_set_calculate_receive_time(options Nw_protocol_options_t, calculate_receive_time bool) {
	if _nw_ip_options_set_calculate_receive_time == nil {
		panic("Network: symbol nw_ip_options_set_calculate_receive_time not loaded")
	}
	_nw_ip_options_set_calculate_receive_time(options, calculate_receive_time)
}

var _nw_ip_options_set_disable_fragmentation func(options Nw_protocol_options_t, disable_fragmentation bool)

// Nw_ip_options_set_disable_fragmentation configures a connection to disable fragmentation on outbound packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_fragmentation(_:_:)
func Nw_ip_options_set_disable_fragmentation(options Nw_protocol_options_t, disable_fragmentation bool) {
	if _nw_ip_options_set_disable_fragmentation == nil {
		panic("Network: symbol nw_ip_options_set_disable_fragmentation not loaded")
	}
	_nw_ip_options_set_disable_fragmentation(options, disable_fragmentation)
}

var _nw_ip_options_set_disable_multicast_loopback func(options Nw_protocol_options_t, disable_multicast_loopback bool)

// Nw_ip_options_set_disable_multicast_loopback.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_multicast_loopback(_:_:)
func Nw_ip_options_set_disable_multicast_loopback(options Nw_protocol_options_t, disable_multicast_loopback bool) {
	if _nw_ip_options_set_disable_multicast_loopback == nil {
		panic("Network: symbol nw_ip_options_set_disable_multicast_loopback not loaded")
	}
	_nw_ip_options_set_disable_multicast_loopback(options, disable_multicast_loopback)
}

var _nw_ip_options_set_hop_limit func(options Nw_protocol_options_t, hop_limit uint8)

// Nw_ip_options_set_hop_limit configures the default hop limit for packets generated by a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_hop_limit(_:_:)
func Nw_ip_options_set_hop_limit(options Nw_protocol_options_t, hop_limit uint8) {
	if _nw_ip_options_set_hop_limit == nil {
		panic("Network: symbol nw_ip_options_set_hop_limit not loaded")
	}
	_nw_ip_options_set_hop_limit(options, hop_limit)
}

var _nw_ip_options_set_local_address_preference func(options Nw_protocol_options_t, preference unsafe.Pointer)

// Nw_ip_options_set_local_address_preference configures a connection to prefer certain types of local addresses, such as temporary or stable.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_local_address_preference(_:_:)
func Nw_ip_options_set_local_address_preference(options Nw_protocol_options_t, preference unsafe.Pointer) {
	if _nw_ip_options_set_local_address_preference == nil {
		panic("Network: symbol nw_ip_options_set_local_address_preference not loaded")
	}
	_nw_ip_options_set_local_address_preference(options, preference)
}

var _nw_ip_options_set_use_minimum_mtu func(options Nw_protocol_options_t, use_minimum_mtu bool)

// Nw_ip_options_set_use_minimum_mtu configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_use_minimum_mtu(_:_:)
func Nw_ip_options_set_use_minimum_mtu(options Nw_protocol_options_t, use_minimum_mtu bool) {
	if _nw_ip_options_set_use_minimum_mtu == nil {
		panic("Network: symbol nw_ip_options_set_use_minimum_mtu not loaded")
	}
	_nw_ip_options_set_use_minimum_mtu(options, use_minimum_mtu)
}

var _nw_ip_options_set_version func(options Nw_protocol_options_t, version unsafe.Pointer)

// Nw_ip_options_set_version sets a required IP version to disable all other versions for a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_version(_:_:)
func Nw_ip_options_set_version(options Nw_protocol_options_t, version unsafe.Pointer) {
	if _nw_ip_options_set_version == nil {
		panic("Network: symbol nw_ip_options_set_version not loaded")
	}
	_nw_ip_options_set_version(options, version)
}

var _nw_listener_cancel func(listener Nw_listener_t)

// Nw_listener_cancel stops listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_cancel(_:)
func Nw_listener_cancel(listener Nw_listener_t) {
	if _nw_listener_cancel == nil {
		panic("Network: symbol nw_listener_cancel not loaded")
	}
	_nw_listener_cancel(listener)
}

var _nw_listener_create func(parameters Nw_parameters_t) Nw_listener_t

// Nw_listener_create initializes a network listener, which will select a random port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create(_:)
func Nw_listener_create(parameters Nw_parameters_t) Nw_listener_t {
	if _nw_listener_create == nil {
		panic("Network: symbol nw_listener_create not loaded")
	}
	return _nw_listener_create(parameters)
}

var _nw_listener_create_with_connection func(connection Nw_connection_t, parameters Nw_parameters_t) Nw_listener_t

// Nw_listener_create_with_connection initializes a network listener to receive new streams on a multiplexed connection.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_connection(_:_:)
func Nw_listener_create_with_connection(connection Nw_connection_t, parameters Nw_parameters_t) Nw_listener_t {
	if _nw_listener_create_with_connection == nil {
		panic("Network: symbol nw_listener_create_with_connection not loaded")
	}
	return _nw_listener_create_with_connection(connection, parameters)
}

var _nw_listener_create_with_launchd_key func(parameters Nw_parameters_t, launchd_key string) Nw_listener_t

// Nw_listener_create_with_launchd_key.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_launchd_key(_:_:)
func Nw_listener_create_with_launchd_key(parameters Nw_parameters_t, launchd_key string) Nw_listener_t {
	if _nw_listener_create_with_launchd_key == nil {
		panic("Network: symbol nw_listener_create_with_launchd_key not loaded")
	}
	return _nw_listener_create_with_launchd_key(parameters, launchd_key)
}

var _nw_listener_create_with_port func(port string, parameters Nw_parameters_t) Nw_listener_t

// Nw_listener_create_with_port initializes a network listener with a specified local port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_port(_:_:)
func Nw_listener_create_with_port(port string, parameters Nw_parameters_t) Nw_listener_t {
	if _nw_listener_create_with_port == nil {
		panic("Network: symbol nw_listener_create_with_port not loaded")
	}
	return _nw_listener_create_with_port(port, parameters)
}

var _nw_listener_get_new_connection_limit func(listener Nw_listener_t) uint32

// Nw_listener_get_new_connection_limit checks the remaining number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_new_connection_limit(_:)
func Nw_listener_get_new_connection_limit(listener Nw_listener_t) uint32 {
	if _nw_listener_get_new_connection_limit == nil {
		panic("Network: symbol nw_listener_get_new_connection_limit not loaded")
	}
	return _nw_listener_get_new_connection_limit(listener)
}

var _nw_listener_get_port func(listener Nw_listener_t) uint16

// Nw_listener_get_port the port on which the listener can accept connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_port(_:)
func Nw_listener_get_port(listener Nw_listener_t) uint16 {
	if _nw_listener_get_port == nil {
		panic("Network: symbol nw_listener_get_port not loaded")
	}
	return _nw_listener_get_port(listener)
}

var _nw_listener_set_advertise_descriptor func(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t)

// Nw_listener_set_advertise_descriptor sets a Bonjour service that advertises the listener on the local network.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertise_descriptor(_:_:)
func Nw_listener_set_advertise_descriptor(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t) {
	if _nw_listener_set_advertise_descriptor == nil {
		panic("Network: symbol nw_listener_set_advertise_descriptor not loaded")
	}
	_nw_listener_set_advertise_descriptor(listener, advertise_descriptor)
}

var _nw_listener_set_advertised_endpoint_changed_handler func(listener Nw_listener_t, handler Nw_listener_advertised_endpoint_changed_handler_t)

// Nw_listener_set_advertised_endpoint_changed_handler sets a handler that receives updates for the service endpoint being advertised.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertised_endpoint_changed_handler(_:_:)
func Nw_listener_set_advertised_endpoint_changed_handler(listener Nw_listener_t, handler Nw_listener_advertised_endpoint_changed_handler_t) {
	if _nw_listener_set_advertised_endpoint_changed_handler == nil {
		panic("Network: symbol nw_listener_set_advertised_endpoint_changed_handler not loaded")
	}
	_nw_listener_set_advertised_endpoint_changed_handler(listener, handler)
}

var _nw_listener_set_new_connection_group_handler func(listener Nw_listener_t, handler Nw_listener_new_connection_group_handler_t)

// Nw_listener_set_new_connection_group_handler.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_group_handler(_:_:)
func Nw_listener_set_new_connection_group_handler(listener Nw_listener_t, handler Nw_listener_new_connection_group_handler_t) {
	if _nw_listener_set_new_connection_group_handler == nil {
		panic("Network: symbol nw_listener_set_new_connection_group_handler not loaded")
	}
	_nw_listener_set_new_connection_group_handler(listener, handler)
}

var _nw_listener_set_new_connection_handler func(listener Nw_listener_t, handler Nw_listener_new_connection_handler_t)

// Nw_listener_set_new_connection_handler sets a handler that receives inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_handler(_:_:)
func Nw_listener_set_new_connection_handler(listener Nw_listener_t, handler Nw_listener_new_connection_handler_t) {
	if _nw_listener_set_new_connection_handler == nil {
		panic("Network: symbol nw_listener_set_new_connection_handler not loaded")
	}
	_nw_listener_set_new_connection_handler(listener, handler)
}

var _nw_listener_set_new_connection_limit func(listener Nw_listener_t, new_connection_limit uint32)

// Nw_listener_set_new_connection_limit resets the number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_limit(_:_:)
func Nw_listener_set_new_connection_limit(listener Nw_listener_t, new_connection_limit uint32) {
	if _nw_listener_set_new_connection_limit == nil {
		panic("Network: symbol nw_listener_set_new_connection_limit not loaded")
	}
	_nw_listener_set_new_connection_limit(listener, new_connection_limit)
}

var _nw_listener_set_queue func(listener Nw_listener_t, queue uintptr)

// Nw_listener_set_queue sets the queue on which all listener events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_queue(_:_:)
func Nw_listener_set_queue(listener Nw_listener_t, queue dispatch.Queue) {
	if _nw_listener_set_queue == nil {
		panic("Network: symbol nw_listener_set_queue not loaded")
	}
	_nw_listener_set_queue(listener, uintptr(queue.Handle()))
}

var _nw_listener_set_state_changed_handler func(listener Nw_listener_t, handler Nw_listener_state_changed_handler_t)

// Nw_listener_set_state_changed_handler sets a handler to receive listener state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_state_changed_handler(_:_:)
func Nw_listener_set_state_changed_handler(listener Nw_listener_t, handler Nw_listener_state_changed_handler_t) {
	if _nw_listener_set_state_changed_handler == nil {
		panic("Network: symbol nw_listener_set_state_changed_handler not loaded")
	}
	_nw_listener_set_state_changed_handler(listener, handler)
}

var _nw_listener_start func(listener Nw_listener_t)

// Nw_listener_start registers for listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_start(_:)
func Nw_listener_start(listener Nw_listener_t) {
	if _nw_listener_start == nil {
		panic("Network: symbol nw_listener_start not loaded")
	}
	_nw_listener_start(listener)
}

var _nw_multicast_group_descriptor_get_disable_unicast_traffic func(multicast_descriptor Nw_group_descriptor_t) bool

// Nw_multicast_group_descriptor_get_disable_unicast_traffic checks a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_get_disable_unicast_traffic(_:)
func Nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t) bool {
	if _nw_multicast_group_descriptor_get_disable_unicast_traffic == nil {
		panic("Network: symbol nw_multicast_group_descriptor_get_disable_unicast_traffic not loaded")
	}
	return _nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor)
}

var _nw_multicast_group_descriptor_set_disable_unicast_traffic func(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool)

// Nw_multicast_group_descriptor_set_disable_unicast_traffic sets a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_disable_unicast_traffic(_:_:)
func Nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool) {
	if _nw_multicast_group_descriptor_set_disable_unicast_traffic == nil {
		panic("Network: symbol nw_multicast_group_descriptor_set_disable_unicast_traffic not loaded")
	}
	_nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor, disable_unicast_traffic)
}

var _nw_multicast_group_descriptor_set_specific_source func(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t)

// Nw_multicast_group_descriptor_set_specific_source sets an optional address endpoint used to filter received multicast packets.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_specific_source(_:_:)
func Nw_multicast_group_descriptor_set_specific_source(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t) {
	if _nw_multicast_group_descriptor_set_specific_source == nil {
		panic("Network: symbol nw_multicast_group_descriptor_set_specific_source not loaded")
	}
	_nw_multicast_group_descriptor_set_specific_source(multicast_descriptor, source)
}

var _nw_parameters_clear_prohibited_interface_types func(parameters Nw_parameters_t)

// Nw_parameters_clear_prohibited_interface_types removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interface_types(_:)
func Nw_parameters_clear_prohibited_interface_types(parameters Nw_parameters_t) {
	if _nw_parameters_clear_prohibited_interface_types == nil {
		panic("Network: symbol nw_parameters_clear_prohibited_interface_types not loaded")
	}
	_nw_parameters_clear_prohibited_interface_types(parameters)
}

var _nw_parameters_clear_prohibited_interfaces func(parameters Nw_parameters_t)

// Nw_parameters_clear_prohibited_interfaces removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interfaces(_:)
func Nw_parameters_clear_prohibited_interfaces(parameters Nw_parameters_t) {
	if _nw_parameters_clear_prohibited_interfaces == nil {
		panic("Network: symbol nw_parameters_clear_prohibited_interfaces not loaded")
	}
	_nw_parameters_clear_prohibited_interfaces(parameters)
}

var _nw_parameters_copy func(parameters Nw_parameters_t) Nw_parameters_t

// Nw_parameters_copy peforms a deep copy of a parameters object.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy(_:)
func Nw_parameters_copy(parameters Nw_parameters_t) Nw_parameters_t {
	if _nw_parameters_copy == nil {
		panic("Network: symbol nw_parameters_copy not loaded")
	}
	return _nw_parameters_copy(parameters)
}

var _nw_parameters_copy_default_protocol_stack func(parameters Nw_parameters_t) Nw_protocol_stack_t

// Nw_parameters_copy_default_protocol_stack accesses the protocol stack used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_default_protocol_stack(_:)
func Nw_parameters_copy_default_protocol_stack(parameters Nw_parameters_t) Nw_protocol_stack_t {
	if _nw_parameters_copy_default_protocol_stack == nil {
		panic("Network: symbol nw_parameters_copy_default_protocol_stack not loaded")
	}
	return _nw_parameters_copy_default_protocol_stack(parameters)
}

var _nw_parameters_copy_local_endpoint func(parameters Nw_parameters_t) Nw_endpoint_t

// Nw_parameters_copy_local_endpoint accesses the local IP address and port used for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_local_endpoint(_:)
func Nw_parameters_copy_local_endpoint(parameters Nw_parameters_t) Nw_endpoint_t {
	if _nw_parameters_copy_local_endpoint == nil {
		panic("Network: symbol nw_parameters_copy_local_endpoint not loaded")
	}
	return _nw_parameters_copy_local_endpoint(parameters)
}

var _nw_parameters_copy_required_interface func(parameters Nw_parameters_t) Nw_interface_t

// Nw_parameters_copy_required_interface accesses the interface required on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_required_interface(_:)
func Nw_parameters_copy_required_interface(parameters Nw_parameters_t) Nw_interface_t {
	if _nw_parameters_copy_required_interface == nil {
		panic("Network: symbol nw_parameters_copy_required_interface not loaded")
	}
	return _nw_parameters_copy_required_interface(parameters)
}

var _nw_parameters_create func() Nw_parameters_t

// Nw_parameters_create initializes parameters for connections, listeners, and browsers with no protocols specified.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create()
func Nw_parameters_create() Nw_parameters_t {
	if _nw_parameters_create == nil {
		panic("Network: symbol nw_parameters_create not loaded")
	}
	return _nw_parameters_create()
}

var _nw_parameters_create_application_service func() Nw_parameters_t

// Nw_parameters_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_application_service()
func Nw_parameters_create_application_service() Nw_parameters_t {
	if _nw_parameters_create_application_service == nil {
		panic("Network: symbol nw_parameters_create_application_service not loaded")
	}
	return _nw_parameters_create_application_service()
}

var _nw_parameters_create_custom_ip func(custom_ip_protocol_number uint8, configure_ip Nw_parameters_configure_protocol_block_t) Nw_parameters_t

// Nw_parameters_create_custom_ip initializes parameters for connections and listeners using a custom IP protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_custom_ip(_:_:)
func Nw_parameters_create_custom_ip(custom_ip_protocol_number uint8, configure_ip Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	if _nw_parameters_create_custom_ip == nil {
		panic("Network: symbol nw_parameters_create_custom_ip not loaded")
	}
	return _nw_parameters_create_custom_ip(custom_ip_protocol_number, configure_ip)
}

var _nw_parameters_create_quic func(configure_quic Nw_parameters_configure_protocol_block_t) Nw_parameters_t

// Nw_parameters_create_quic initializes parameters for QUIC connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_quic(_:)
func Nw_parameters_create_quic(configure_quic Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	if _nw_parameters_create_quic == nil {
		panic("Network: symbol nw_parameters_create_quic not loaded")
	}
	return _nw_parameters_create_quic(configure_quic)
}

var _nw_parameters_create_secure_tcp func(configure_tls Nw_parameters_configure_protocol_block_t, configure_tcp Nw_parameters_configure_protocol_block_t) Nw_parameters_t

// Nw_parameters_create_secure_tcp initializes parameters for TLS or TCP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_tcp(_:_:)
func Nw_parameters_create_secure_tcp(configure_tls Nw_parameters_configure_protocol_block_t, configure_tcp Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	if _nw_parameters_create_secure_tcp == nil {
		panic("Network: symbol nw_parameters_create_secure_tcp not loaded")
	}
	return _nw_parameters_create_secure_tcp(configure_tls, configure_tcp)
}

var _nw_parameters_create_secure_udp func(configure_dtls Nw_parameters_configure_protocol_block_t, configure_udp Nw_parameters_configure_protocol_block_t) Nw_parameters_t

// Nw_parameters_create_secure_udp initializes parameters for DTLS or UDP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_udp(_:_:)
func Nw_parameters_create_secure_udp(configure_dtls Nw_parameters_configure_protocol_block_t, configure_udp Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	if _nw_parameters_create_secure_udp == nil {
		panic("Network: symbol nw_parameters_create_secure_udp not loaded")
	}
	return _nw_parameters_create_secure_udp(configure_dtls, configure_udp)
}

var _nw_parameters_get_allow_ultra_constrained func(parameters Nw_parameters_t) bool

// Nw_parameters_get_allow_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_allow_ultra_constrained(_:)
func Nw_parameters_get_allow_ultra_constrained(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_allow_ultra_constrained == nil {
		panic("Network: symbol nw_parameters_get_allow_ultra_constrained not loaded")
	}
	return _nw_parameters_get_allow_ultra_constrained(parameters)
}

var _nw_parameters_get_expired_dns_behavior func(parameters Nw_parameters_t) unsafe.Pointer

// Nw_parameters_get_expired_dns_behavior checks the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_expired_dns_behavior(_:)
func Nw_parameters_get_expired_dns_behavior(parameters Nw_parameters_t) unsafe.Pointer {
	if _nw_parameters_get_expired_dns_behavior == nil {
		panic("Network: symbol nw_parameters_get_expired_dns_behavior not loaded")
	}
	return _nw_parameters_get_expired_dns_behavior(parameters)
}

var _nw_parameters_get_fast_open_enabled func(parameters Nw_parameters_t) bool

// Nw_parameters_get_fast_open_enabled checks if sending application data with protocol handshakes is enabled.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_fast_open_enabled(_:)
func Nw_parameters_get_fast_open_enabled(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_fast_open_enabled == nil {
		panic("Network: symbol nw_parameters_get_fast_open_enabled not loaded")
	}
	return _nw_parameters_get_fast_open_enabled(parameters)
}

var _nw_parameters_get_include_peer_to_peer func(parameters Nw_parameters_t) bool

// Nw_parameters_get_include_peer_to_peer checks whether a connection is allowed to use peer-to-peer link technologies.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_include_peer_to_peer(_:)
func Nw_parameters_get_include_peer_to_peer(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_include_peer_to_peer == nil {
		panic("Network: symbol nw_parameters_get_include_peer_to_peer not loaded")
	}
	return _nw_parameters_get_include_peer_to_peer(parameters)
}

var _nw_parameters_get_local_only func(parameters Nw_parameters_t) bool

// Nw_parameters_get_local_only checks if a listener is restricted to accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_local_only(_:)
func Nw_parameters_get_local_only(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_local_only == nil {
		panic("Network: symbol nw_parameters_get_local_only not loaded")
	}
	return _nw_parameters_get_local_only(parameters)
}

var _nw_parameters_get_multipath_service func(parameters Nw_parameters_t) unsafe.Pointer

// Nw_parameters_get_multipath_service checks if multipath is enabled on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_multipath_service(_:)
func Nw_parameters_get_multipath_service(parameters Nw_parameters_t) unsafe.Pointer {
	if _nw_parameters_get_multipath_service == nil {
		panic("Network: symbol nw_parameters_get_multipath_service not loaded")
	}
	return _nw_parameters_get_multipath_service(parameters)
}

var _nw_parameters_get_prefer_no_proxy func(parameters Nw_parameters_t) bool

// Nw_parameters_get_prefer_no_proxy checks if proxies are ignored by default.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prefer_no_proxy(_:)
func Nw_parameters_get_prefer_no_proxy(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_prefer_no_proxy == nil {
		panic("Network: symbol nw_parameters_get_prefer_no_proxy not loaded")
	}
	return _nw_parameters_get_prefer_no_proxy(parameters)
}

var _nw_parameters_get_prohibit_constrained func(parameters Nw_parameters_t) bool

// Nw_parameters_get_prohibit_constrained checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_constrained(_:)
func Nw_parameters_get_prohibit_constrained(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_prohibit_constrained == nil {
		panic("Network: symbol nw_parameters_get_prohibit_constrained not loaded")
	}
	return _nw_parameters_get_prohibit_constrained(parameters)
}

var _nw_parameters_get_prohibit_expensive func(parameters Nw_parameters_t) bool

// Nw_parameters_get_prohibit_expensive checks if connections, listeners, and browsers are prevented from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_expensive(_:)
func Nw_parameters_get_prohibit_expensive(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_prohibit_expensive == nil {
		panic("Network: symbol nw_parameters_get_prohibit_expensive not loaded")
	}
	return _nw_parameters_get_prohibit_expensive(parameters)
}

var _nw_parameters_get_required_interface_type func(parameters Nw_parameters_t) unsafe.Pointer

// Nw_parameters_get_required_interface_type accesses the interface type required on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_required_interface_type(_:)
func Nw_parameters_get_required_interface_type(parameters Nw_parameters_t) unsafe.Pointer {
	if _nw_parameters_get_required_interface_type == nil {
		panic("Network: symbol nw_parameters_get_required_interface_type not loaded")
	}
	return _nw_parameters_get_required_interface_type(parameters)
}

var _nw_parameters_get_reuse_local_address func(parameters Nw_parameters_t) bool

// Nw_parameters_get_reuse_local_address checks whether a connection allows reusing local addresses and ports.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_reuse_local_address(_:)
func Nw_parameters_get_reuse_local_address(parameters Nw_parameters_t) bool {
	if _nw_parameters_get_reuse_local_address == nil {
		panic("Network: symbol nw_parameters_get_reuse_local_address not loaded")
	}
	return _nw_parameters_get_reuse_local_address(parameters)
}

var _nw_parameters_get_service_class func(parameters Nw_parameters_t) unsafe.Pointer

// Nw_parameters_get_service_class checks the level of service quality used for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_service_class(_:)
func Nw_parameters_get_service_class(parameters Nw_parameters_t) unsafe.Pointer {
	if _nw_parameters_get_service_class == nil {
		panic("Network: symbol nw_parameters_get_service_class not loaded")
	}
	return _nw_parameters_get_service_class(parameters)
}

var _nw_parameters_iterate_prohibited_interface_types func(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interface_types_block_t)

// Nw_parameters_iterate_prohibited_interface_types examines the list of prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interface_types(_:_:)
func Nw_parameters_iterate_prohibited_interface_types(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interface_types_block_t) {
	if _nw_parameters_iterate_prohibited_interface_types == nil {
		panic("Network: symbol nw_parameters_iterate_prohibited_interface_types not loaded")
	}
	_nw_parameters_iterate_prohibited_interface_types(parameters, iterate_block)
}

var _nw_parameters_iterate_prohibited_interfaces func(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interfaces_block_t)

// Nw_parameters_iterate_prohibited_interfaces examines the list of prohibited interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interfaces(_:_:)
func Nw_parameters_iterate_prohibited_interfaces(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interfaces_block_t) {
	if _nw_parameters_iterate_prohibited_interfaces == nil {
		panic("Network: symbol nw_parameters_iterate_prohibited_interfaces not loaded")
	}
	_nw_parameters_iterate_prohibited_interfaces(parameters, iterate_block)
}

var _nw_parameters_prohibit_interface func(parameters Nw_parameters_t, interface_ Nw_interface_t)

// Nw_parameters_prohibit_interface prevents connections and listeners from using a specific interface.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface(_:_:)
func Nw_parameters_prohibit_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	if _nw_parameters_prohibit_interface == nil {
		panic("Network: symbol nw_parameters_prohibit_interface not loaded")
	}
	_nw_parameters_prohibit_interface(parameters, interface_)
}

var _nw_parameters_prohibit_interface_type func(parameters Nw_parameters_t, interface_type unsafe.Pointer)

// Nw_parameters_prohibit_interface_type prevents connections, listeners, and browsers from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface_type(_:_:)
func Nw_parameters_prohibit_interface_type(parameters Nw_parameters_t, interface_type unsafe.Pointer) {
	if _nw_parameters_prohibit_interface_type == nil {
		panic("Network: symbol nw_parameters_prohibit_interface_type not loaded")
	}
	_nw_parameters_prohibit_interface_type(parameters, interface_type)
}

var _nw_parameters_require_interface func(parameters Nw_parameters_t, interface_ Nw_interface_t)

// Nw_parameters_require_interface sets a specific interface to require on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_require_interface(_:_:)
func Nw_parameters_require_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	if _nw_parameters_require_interface == nil {
		panic("Network: symbol nw_parameters_require_interface not loaded")
	}
	_nw_parameters_require_interface(parameters, interface_)
}

var _nw_parameters_requires_dnssec_validation func(parameters Nw_parameters_t) bool

// Nw_parameters_requires_dnssec_validation checks whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_requires_dnssec_validation(_:)
func Nw_parameters_requires_dnssec_validation(parameters Nw_parameters_t) bool {
	if _nw_parameters_requires_dnssec_validation == nil {
		panic("Network: symbol nw_parameters_requires_dnssec_validation not loaded")
	}
	return _nw_parameters_requires_dnssec_validation(parameters)
}

var _nw_parameters_set_allow_ultra_constrained func(parameters Nw_parameters_t, allow_ultra_constrained bool)

// Nw_parameters_set_allow_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_allow_ultra_constrained(_:_:)
func Nw_parameters_set_allow_ultra_constrained(parameters Nw_parameters_t, allow_ultra_constrained bool) {
	if _nw_parameters_set_allow_ultra_constrained == nil {
		panic("Network: symbol nw_parameters_set_allow_ultra_constrained not loaded")
	}
	_nw_parameters_set_allow_ultra_constrained(parameters, allow_ultra_constrained)
}

var _nw_parameters_set_expired_dns_behavior func(parameters Nw_parameters_t, expired_dns_behavior unsafe.Pointer)

// Nw_parameters_set_expired_dns_behavior sets the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_expired_dns_behavior(_:_:)
func Nw_parameters_set_expired_dns_behavior(parameters Nw_parameters_t, expired_dns_behavior unsafe.Pointer) {
	if _nw_parameters_set_expired_dns_behavior == nil {
		panic("Network: symbol nw_parameters_set_expired_dns_behavior not loaded")
	}
	_nw_parameters_set_expired_dns_behavior(parameters, expired_dns_behavior)
}

var _nw_parameters_set_fast_open_enabled func(parameters Nw_parameters_t, fast_open_enabled bool)

// Nw_parameters_set_fast_open_enabled enables sending application data with protocol handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_fast_open_enabled(_:_:)
func Nw_parameters_set_fast_open_enabled(parameters Nw_parameters_t, fast_open_enabled bool) {
	if _nw_parameters_set_fast_open_enabled == nil {
		panic("Network: symbol nw_parameters_set_fast_open_enabled not loaded")
	}
	_nw_parameters_set_fast_open_enabled(parameters, fast_open_enabled)
}

var _nw_parameters_set_include_peer_to_peer func(parameters Nw_parameters_t, include_peer_to_peer bool)

// Nw_parameters_set_include_peer_to_peer enables peer-to-peer link technologies for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_include_peer_to_peer(_:_:)
func Nw_parameters_set_include_peer_to_peer(parameters Nw_parameters_t, include_peer_to_peer bool) {
	if _nw_parameters_set_include_peer_to_peer == nil {
		panic("Network: symbol nw_parameters_set_include_peer_to_peer not loaded")
	}
	_nw_parameters_set_include_peer_to_peer(parameters, include_peer_to_peer)
}

var _nw_parameters_set_local_endpoint func(parameters Nw_parameters_t, local_endpoint Nw_endpoint_t)

// Nw_parameters_set_local_endpoint sets a specific local IP address and port to use for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_endpoint(_:_:)
func Nw_parameters_set_local_endpoint(parameters Nw_parameters_t, local_endpoint Nw_endpoint_t) {
	if _nw_parameters_set_local_endpoint == nil {
		panic("Network: symbol nw_parameters_set_local_endpoint not loaded")
	}
	_nw_parameters_set_local_endpoint(parameters, local_endpoint)
}

var _nw_parameters_set_local_only func(parameters Nw_parameters_t, local_only bool)

// Nw_parameters_set_local_only restricts listeners to only accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_only(_:_:)
func Nw_parameters_set_local_only(parameters Nw_parameters_t, local_only bool) {
	if _nw_parameters_set_local_only == nil {
		panic("Network: symbol nw_parameters_set_local_only not loaded")
	}
	_nw_parameters_set_local_only(parameters, local_only)
}

var _nw_parameters_set_multipath_service func(parameters Nw_parameters_t, multipath_service unsafe.Pointer)

// Nw_parameters_set_multipath_service enables multipath protocols to allow connections to use multiple interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_multipath_service(_:_:)
func Nw_parameters_set_multipath_service(parameters Nw_parameters_t, multipath_service unsafe.Pointer) {
	if _nw_parameters_set_multipath_service == nil {
		panic("Network: symbol nw_parameters_set_multipath_service not loaded")
	}
	_nw_parameters_set_multipath_service(parameters, multipath_service)
}

var _nw_parameters_set_prefer_no_proxy func(parameters Nw_parameters_t, prefer_no_proxy bool)

// Nw_parameters_set_prefer_no_proxy sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prefer_no_proxy(_:_:)
func Nw_parameters_set_prefer_no_proxy(parameters Nw_parameters_t, prefer_no_proxy bool) {
	if _nw_parameters_set_prefer_no_proxy == nil {
		panic("Network: symbol nw_parameters_set_prefer_no_proxy not loaded")
	}
	_nw_parameters_set_prefer_no_proxy(parameters, prefer_no_proxy)
}

var _nw_parameters_set_privacy_context func(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t)

// Nw_parameters_set_privacy_context associates a privacy context with any connections or listeners that use the parameters.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_privacy_context(_:_:)
func Nw_parameters_set_privacy_context(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t) {
	if _nw_parameters_set_privacy_context == nil {
		panic("Network: symbol nw_parameters_set_privacy_context not loaded")
	}
	_nw_parameters_set_privacy_context(parameters, privacy_context)
}

var _nw_parameters_set_prohibit_constrained func(parameters Nw_parameters_t, prohibit_constrained bool)

// Nw_parameters_set_prohibit_constrained prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_constrained(_:_:)
func Nw_parameters_set_prohibit_constrained(parameters Nw_parameters_t, prohibit_constrained bool) {
	if _nw_parameters_set_prohibit_constrained == nil {
		panic("Network: symbol nw_parameters_set_prohibit_constrained not loaded")
	}
	_nw_parameters_set_prohibit_constrained(parameters, prohibit_constrained)
}

var _nw_parameters_set_prohibit_expensive func(parameters Nw_parameters_t, prohibit_expensive bool)

// Nw_parameters_set_prohibit_expensive prevents connections, listeners, and browsers from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_expensive(_:_:)
func Nw_parameters_set_prohibit_expensive(parameters Nw_parameters_t, prohibit_expensive bool) {
	if _nw_parameters_set_prohibit_expensive == nil {
		panic("Network: symbol nw_parameters_set_prohibit_expensive not loaded")
	}
	_nw_parameters_set_prohibit_expensive(parameters, prohibit_expensive)
}

var _nw_parameters_set_required_interface_type func(parameters Nw_parameters_t, interface_type unsafe.Pointer)

// Nw_parameters_set_required_interface_type sets an interface type to require on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_required_interface_type(_:_:)
func Nw_parameters_set_required_interface_type(parameters Nw_parameters_t, interface_type unsafe.Pointer) {
	if _nw_parameters_set_required_interface_type == nil {
		panic("Network: symbol nw_parameters_set_required_interface_type not loaded")
	}
	_nw_parameters_set_required_interface_type(parameters, interface_type)
}

var _nw_parameters_set_requires_dnssec_validation func(parameters Nw_parameters_t, requires_dnssec_validation bool)

// Nw_parameters_set_requires_dnssec_validation determines whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_requires_dnssec_validation(_:_:)
func Nw_parameters_set_requires_dnssec_validation(parameters Nw_parameters_t, requires_dnssec_validation bool) {
	if _nw_parameters_set_requires_dnssec_validation == nil {
		panic("Network: symbol nw_parameters_set_requires_dnssec_validation not loaded")
	}
	_nw_parameters_set_requires_dnssec_validation(parameters, requires_dnssec_validation)
}

var _nw_parameters_set_reuse_local_address func(parameters Nw_parameters_t, reuse_local_address bool)

// Nw_parameters_set_reuse_local_address allows reusing local addresses and ports across connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_reuse_local_address(_:_:)
func Nw_parameters_set_reuse_local_address(parameters Nw_parameters_t, reuse_local_address bool) {
	if _nw_parameters_set_reuse_local_address == nil {
		panic("Network: symbol nw_parameters_set_reuse_local_address not loaded")
	}
	_nw_parameters_set_reuse_local_address(parameters, reuse_local_address)
}

var _nw_parameters_set_service_class func(parameters Nw_parameters_t, service_class unsafe.Pointer)

// Nw_parameters_set_service_class sets a level of service quality to use for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_service_class(_:_:)
func Nw_parameters_set_service_class(parameters Nw_parameters_t, service_class unsafe.Pointer) {
	if _nw_parameters_set_service_class == nil {
		panic("Network: symbol nw_parameters_set_service_class not loaded")
	}
	_nw_parameters_set_service_class(parameters, service_class)
}

var _nw_path_copy_effective_local_endpoint func(path Nw_path_t) Nw_endpoint_t

// Nw_path_copy_effective_local_endpoint accesses the local endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_local_endpoint(_:)
func Nw_path_copy_effective_local_endpoint(path Nw_path_t) Nw_endpoint_t {
	if _nw_path_copy_effective_local_endpoint == nil {
		panic("Network: symbol nw_path_copy_effective_local_endpoint not loaded")
	}
	return _nw_path_copy_effective_local_endpoint(path)
}

var _nw_path_copy_effective_remote_endpoint func(path Nw_path_t) Nw_endpoint_t

// Nw_path_copy_effective_remote_endpoint accesses the remote endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_remote_endpoint(_:)
func Nw_path_copy_effective_remote_endpoint(path Nw_path_t) Nw_endpoint_t {
	if _nw_path_copy_effective_remote_endpoint == nil {
		panic("Network: symbol nw_path_copy_effective_remote_endpoint not loaded")
	}
	return _nw_path_copy_effective_remote_endpoint(path)
}

var _nw_path_enumerate_gateways func(path Nw_path_t, enumerate_block Nw_path_enumerate_gateways_block_t)

// Nw_path_enumerate_gateways enumerates the list of gateways configured on the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways(_:_:)
func Nw_path_enumerate_gateways(path Nw_path_t, enumerate_block Nw_path_enumerate_gateways_block_t) {
	if _nw_path_enumerate_gateways == nil {
		panic("Network: symbol nw_path_enumerate_gateways not loaded")
	}
	_nw_path_enumerate_gateways(path, enumerate_block)
}

var _nw_path_enumerate_interfaces func(path Nw_path_t, enumerate_block Nw_path_enumerate_interfaces_block_t)

// Nw_path_enumerate_interfaces enumerates the list of all interfaces available to the path, in order of preference.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces(_:_:)
func Nw_path_enumerate_interfaces(path Nw_path_t, enumerate_block Nw_path_enumerate_interfaces_block_t) {
	if _nw_path_enumerate_interfaces == nil {
		panic("Network: symbol nw_path_enumerate_interfaces not loaded")
	}
	_nw_path_enumerate_interfaces(path, enumerate_block)
}

var _nw_path_get_link_quality func(path Nw_path_t) unsafe.Pointer

// Nw_path_get_link_quality.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_link_quality(_:)
func Nw_path_get_link_quality(path Nw_path_t) unsafe.Pointer {
	if _nw_path_get_link_quality == nil {
		panic("Network: symbol nw_path_get_link_quality not loaded")
	}
	return _nw_path_get_link_quality(path)
}

var _nw_path_get_status func(path Nw_path_t) unsafe.Pointer

// Nw_path_get_status checks whether a path can be used by connections.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_status(_:)
func Nw_path_get_status(path Nw_path_t) unsafe.Pointer {
	if _nw_path_get_status == nil {
		panic("Network: symbol nw_path_get_status not loaded")
	}
	return _nw_path_get_status(path)
}

var _nw_path_get_unsatisfied_reason func(path Nw_path_t) unsafe.Pointer

// Nw_path_get_unsatisfied_reason.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_unsatisfied_reason(_:)
func Nw_path_get_unsatisfied_reason(path Nw_path_t) unsafe.Pointer {
	if _nw_path_get_unsatisfied_reason == nil {
		panic("Network: symbol nw_path_get_unsatisfied_reason not loaded")
	}
	return _nw_path_get_unsatisfied_reason(path)
}

var _nw_path_has_dns func(path Nw_path_t) bool

// Nw_path_has_dns checks whether the path has a DNS server configured.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_dns(_:)
func Nw_path_has_dns(path Nw_path_t) bool {
	if _nw_path_has_dns == nil {
		panic("Network: symbol nw_path_has_dns not loaded")
	}
	return _nw_path_has_dns(path)
}

var _nw_path_has_ipv4 func(path Nw_path_t) bool

// Nw_path_has_ipv4 checks whether the path can route IPv4 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv4(_:)
func Nw_path_has_ipv4(path Nw_path_t) bool {
	if _nw_path_has_ipv4 == nil {
		panic("Network: symbol nw_path_has_ipv4 not loaded")
	}
	return _nw_path_has_ipv4(path)
}

var _nw_path_has_ipv6 func(path Nw_path_t) bool

// Nw_path_has_ipv6 checks whether the path can route IPv6 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv6(_:)
func Nw_path_has_ipv6(path Nw_path_t) bool {
	if _nw_path_has_ipv6 == nil {
		panic("Network: symbol nw_path_has_ipv6 not loaded")
	}
	return _nw_path_has_ipv6(path)
}

var _nw_path_is_constrained func(path Nw_path_t) bool

// Nw_path_is_constrained checks whether the path uses an interface in Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_constrained(_:)
func Nw_path_is_constrained(path Nw_path_t) bool {
	if _nw_path_is_constrained == nil {
		panic("Network: symbol nw_path_is_constrained not loaded")
	}
	return _nw_path_is_constrained(path)
}

var _nw_path_is_equal func(path Nw_path_t, other_path Nw_path_t) bool

// Nw_path_is_equal compares if two paths are identical.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_equal(_:_:)
func Nw_path_is_equal(path Nw_path_t, other_path Nw_path_t) bool {
	if _nw_path_is_equal == nil {
		panic("Network: symbol nw_path_is_equal not loaded")
	}
	return _nw_path_is_equal(path, other_path)
}

var _nw_path_is_expensive func(path Nw_path_t) bool

// Nw_path_is_expensive checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_expensive(_:)
func Nw_path_is_expensive(path Nw_path_t) bool {
	if _nw_path_is_expensive == nil {
		panic("Network: symbol nw_path_is_expensive not loaded")
	}
	return _nw_path_is_expensive(path)
}

var _nw_path_is_ultra_constrained func(path Nw_path_t) bool

// Nw_path_is_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_ultra_constrained(_:)
func Nw_path_is_ultra_constrained(path Nw_path_t) bool {
	if _nw_path_is_ultra_constrained == nil {
		panic("Network: symbol nw_path_is_ultra_constrained not loaded")
	}
	return _nw_path_is_ultra_constrained(path)
}

var _nw_path_monitor_cancel func(monitor Nw_path_monitor_t)

// Nw_path_monitor_cancel stops receiving network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel(_:)
func Nw_path_monitor_cancel(monitor Nw_path_monitor_t) {
	if _nw_path_monitor_cancel == nil {
		panic("Network: symbol nw_path_monitor_cancel not loaded")
	}
	_nw_path_monitor_cancel(monitor)
}

var _nw_path_monitor_create func() Nw_path_monitor_t

// Nw_path_monitor_create initializes a path monitor to observe all available interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create()
func Nw_path_monitor_create() Nw_path_monitor_t {
	if _nw_path_monitor_create == nil {
		panic("Network: symbol nw_path_monitor_create not loaded")
	}
	return _nw_path_monitor_create()
}

var _nw_path_monitor_create_for_ethernet_channel func() Nw_path_monitor_t

// Nw_path_monitor_create_for_ethernet_channel.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_for_ethernet_channel()
func Nw_path_monitor_create_for_ethernet_channel() Nw_path_monitor_t {
	if _nw_path_monitor_create_for_ethernet_channel == nil {
		panic("Network: symbol nw_path_monitor_create_for_ethernet_channel not loaded")
	}
	return _nw_path_monitor_create_for_ethernet_channel()
}

var _nw_path_monitor_create_with_type func(required_interface_type unsafe.Pointer) Nw_path_monitor_t

// Nw_path_monitor_create_with_type initializes a path monitor to observe a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_with_type(_:)
func Nw_path_monitor_create_with_type(required_interface_type unsafe.Pointer) Nw_path_monitor_t {
	if _nw_path_monitor_create_with_type == nil {
		panic("Network: symbol nw_path_monitor_create_with_type not loaded")
	}
	return _nw_path_monitor_create_with_type(required_interface_type)
}

var _nw_path_monitor_prohibit_interface_type func(monitor Nw_path_monitor_t, interface_type unsafe.Pointer)

// Nw_path_monitor_prohibit_interface_type prohibit a path monitor from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_prohibit_interface_type(_:_:)
func Nw_path_monitor_prohibit_interface_type(monitor Nw_path_monitor_t, interface_type unsafe.Pointer) {
	if _nw_path_monitor_prohibit_interface_type == nil {
		panic("Network: symbol nw_path_monitor_prohibit_interface_type not loaded")
	}
	_nw_path_monitor_prohibit_interface_type(monitor, interface_type)
}

var _nw_path_monitor_set_cancel_handler func(monitor Nw_path_monitor_t, cancel_handler Nw_path_monitor_cancel_handler_t)

// Nw_path_monitor_set_cancel_handler sets a handler to determine when a monitor is fully cancelled and will no longer deliver events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_cancel_handler(_:_:)
func Nw_path_monitor_set_cancel_handler(monitor Nw_path_monitor_t, cancel_handler Nw_path_monitor_cancel_handler_t) {
	if _nw_path_monitor_set_cancel_handler == nil {
		panic("Network: symbol nw_path_monitor_set_cancel_handler not loaded")
	}
	_nw_path_monitor_set_cancel_handler(monitor, cancel_handler)
}

var _nw_path_monitor_set_queue func(monitor Nw_path_monitor_t, queue uintptr)

// Nw_path_monitor_set_queue sets a queue on which to deliver path events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_queue(_:_:)
func Nw_path_monitor_set_queue(monitor Nw_path_monitor_t, queue dispatch.Queue) {
	if _nw_path_monitor_set_queue == nil {
		panic("Network: symbol nw_path_monitor_set_queue not loaded")
	}
	_nw_path_monitor_set_queue(monitor, uintptr(queue.Handle()))
}

var _nw_path_monitor_set_update_handler func(monitor Nw_path_monitor_t, update_handler Nw_path_monitor_update_handler_t)

// Nw_path_monitor_set_update_handler sets a handler to receive network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_update_handler(_:_:)
func Nw_path_monitor_set_update_handler(monitor Nw_path_monitor_t, update_handler Nw_path_monitor_update_handler_t) {
	if _nw_path_monitor_set_update_handler == nil {
		panic("Network: symbol nw_path_monitor_set_update_handler not loaded")
	}
	_nw_path_monitor_set_update_handler(monitor, update_handler)
}

var _nw_path_monitor_start func(monitor Nw_path_monitor_t)

// Nw_path_monitor_start starts monitoring path changes.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_start(_:)
func Nw_path_monitor_start(monitor Nw_path_monitor_t) {
	if _nw_path_monitor_start == nil {
		panic("Network: symbol nw_path_monitor_start not loaded")
	}
	_nw_path_monitor_start(monitor)
}

var _nw_path_uses_interface_type func(path Nw_path_t, interface_type unsafe.Pointer) bool

// Nw_path_uses_interface_type checks if connections using the path may send traffic over a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_uses_interface_type(_:_:)
func Nw_path_uses_interface_type(path Nw_path_t, interface_type unsafe.Pointer) bool {
	if _nw_path_uses_interface_type == nil {
		panic("Network: symbol nw_path_uses_interface_type not loaded")
	}
	return _nw_path_uses_interface_type(path, interface_type)
}

var _nw_privacy_context_add_proxy func(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t)

// Nw_privacy_context_add_proxy applies a proxy configuration to all connections associated with this context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_add_proxy(_:_:)
func Nw_privacy_context_add_proxy(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t) {
	if _nw_privacy_context_add_proxy == nil {
		panic("Network: symbol nw_privacy_context_add_proxy not loaded")
	}
	_nw_privacy_context_add_proxy(privacy_context, proxy_config)
}

var _nw_privacy_context_clear_proxies func(privacy_context Nw_privacy_context_t)

// Nw_privacy_context_clear_proxies clears out any proxies added using nw_privacy_context_add_proxy(_:_:)
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_clear_proxies(_:)
func Nw_privacy_context_clear_proxies(privacy_context Nw_privacy_context_t) {
	if _nw_privacy_context_clear_proxies == nil {
		panic("Network: symbol nw_privacy_context_clear_proxies not loaded")
	}
	_nw_privacy_context_clear_proxies(privacy_context)
}

var _nw_privacy_context_create func(description string) Nw_privacy_context_t

// Nw_privacy_context_create initializes a privacy context with a description string.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_create(_:)
func Nw_privacy_context_create(description string) Nw_privacy_context_t {
	if _nw_privacy_context_create == nil {
		panic("Network: symbol nw_privacy_context_create not loaded")
	}
	return _nw_privacy_context_create(description)
}

var _nw_privacy_context_disable_logging func(privacy_context Nw_privacy_context_t)

// Nw_privacy_context_disable_logging disables system logging of connection activity.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_disable_logging(_:)
func Nw_privacy_context_disable_logging(privacy_context Nw_privacy_context_t) {
	if _nw_privacy_context_disable_logging == nil {
		panic("Network: symbol nw_privacy_context_disable_logging not loaded")
	}
	_nw_privacy_context_disable_logging(privacy_context)
}

var _nw_privacy_context_flush_cache func(privacy_context Nw_privacy_context_t)

// Nw_privacy_context_flush_cache flushes all cached data, such as TLS session state, created by connections associated with the privacy context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_flush_cache(_:)
func Nw_privacy_context_flush_cache(privacy_context Nw_privacy_context_t) {
	if _nw_privacy_context_flush_cache == nil {
		panic("Network: symbol nw_privacy_context_flush_cache not loaded")
	}
	_nw_privacy_context_flush_cache(privacy_context)
}

var _nw_privacy_context_require_encrypted_name_resolution func(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t)

// Nw_privacy_context_require_encrypted_name_resolution requires that any DNS name resolution for connections associated with this context use encrypted transports, such as TLS or HTTPS.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_require_encrypted_name_resolution(_:_:_:)
func Nw_privacy_context_require_encrypted_name_resolution(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t) {
	if _nw_privacy_context_require_encrypted_name_resolution == nil {
		panic("Network: symbol nw_privacy_context_require_encrypted_name_resolution not loaded")
	}
	_nw_privacy_context_require_encrypted_name_resolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config)
}

var _nw_protocol_copy_ip_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_ip_definition accesses the system definition of the Internet Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ip_definition()
func Nw_protocol_copy_ip_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_ip_definition == nil {
		panic("Network: symbol nw_protocol_copy_ip_definition not loaded")
	}
	return _nw_protocol_copy_ip_definition()
}

var _nw_protocol_copy_quic_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_quic_definition accesses the system definition of the QUIC transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_quic_definition()
func Nw_protocol_copy_quic_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_quic_definition == nil {
		panic("Network: symbol nw_protocol_copy_quic_definition not loaded")
	}
	return _nw_protocol_copy_quic_definition()
}

var _nw_protocol_copy_tcp_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_tcp_definition accesses the system definition of the Transport Control Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tcp_definition()
func Nw_protocol_copy_tcp_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_tcp_definition == nil {
		panic("Network: symbol nw_protocol_copy_tcp_definition not loaded")
	}
	return _nw_protocol_copy_tcp_definition()
}

var _nw_protocol_copy_tls_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_tls_definition accesses the system definition of the Transport Layer Security protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tls_definition()
func Nw_protocol_copy_tls_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_tls_definition == nil {
		panic("Network: symbol nw_protocol_copy_tls_definition not loaded")
	}
	return _nw_protocol_copy_tls_definition()
}

var _nw_protocol_copy_udp_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_udp_definition accesses the system definition of the User Datagram Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_udp_definition()
func Nw_protocol_copy_udp_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_udp_definition == nil {
		panic("Network: symbol nw_protocol_copy_udp_definition not loaded")
	}
	return _nw_protocol_copy_udp_definition()
}

var _nw_protocol_copy_ws_definition func() Nw_protocol_definition_t

// Nw_protocol_copy_ws_definition accesses the system definition of the WebSocket protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ws_definition()
func Nw_protocol_copy_ws_definition() Nw_protocol_definition_t {
	if _nw_protocol_copy_ws_definition == nil {
		panic("Network: symbol nw_protocol_copy_ws_definition not loaded")
	}
	return _nw_protocol_copy_ws_definition()
}

var _nw_protocol_definition_is_equal func(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) bool

// Nw_protocol_definition_is_equal compares two protocol definitions, and returns true if they represent the same protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_definition_is_equal(_:_:)
func Nw_protocol_definition_is_equal(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) bool {
	if _nw_protocol_definition_is_equal == nil {
		panic("Network: symbol nw_protocol_definition_is_equal not loaded")
	}
	return _nw_protocol_definition_is_equal(definition1, definition2)
}

var _nw_protocol_metadata_copy_definition func(metadata Nw_protocol_metadata_t) Nw_protocol_definition_t

// Nw_protocol_metadata_copy_definition accesses the protocol definition associated with the metadata object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_copy_definition(_:)
func Nw_protocol_metadata_copy_definition(metadata Nw_protocol_metadata_t) Nw_protocol_definition_t {
	if _nw_protocol_metadata_copy_definition == nil {
		panic("Network: symbol nw_protocol_metadata_copy_definition not loaded")
	}
	return _nw_protocol_metadata_copy_definition(metadata)
}

var _nw_protocol_metadata_is_framer_message func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_framer_message checks if a metadata object represents a custom framer protocol message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_framer_message(_:)
func Nw_protocol_metadata_is_framer_message(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_framer_message == nil {
		panic("Network: symbol nw_protocol_metadata_is_framer_message not loaded")
	}
	return _nw_protocol_metadata_is_framer_message(metadata)
}

var _nw_protocol_metadata_is_ip func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_ip checks whether a metadata object represents an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ip(_:)
func Nw_protocol_metadata_is_ip(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_ip == nil {
		panic("Network: symbol nw_protocol_metadata_is_ip not loaded")
	}
	return _nw_protocol_metadata_is_ip(metadata)
}

var _nw_protocol_metadata_is_quic func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_quic checks whether a metadata object contains QUIC connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_quic(_:)
func Nw_protocol_metadata_is_quic(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_quic == nil {
		panic("Network: symbol nw_protocol_metadata_is_quic not loaded")
	}
	return _nw_protocol_metadata_is_quic(metadata)
}

var _nw_protocol_metadata_is_tcp func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_tcp checks whether a metadata object contains TCP connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tcp(_:)
func Nw_protocol_metadata_is_tcp(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_tcp == nil {
		panic("Network: symbol nw_protocol_metadata_is_tcp not loaded")
	}
	return _nw_protocol_metadata_is_tcp(metadata)
}

var _nw_protocol_metadata_is_tls func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_tls checks whether a metadata object contains TLS connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tls(_:)
func Nw_protocol_metadata_is_tls(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_tls == nil {
		panic("Network: symbol nw_protocol_metadata_is_tls not loaded")
	}
	return _nw_protocol_metadata_is_tls(metadata)
}

var _nw_protocol_metadata_is_udp func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_udp checks whether a metadata object represents a UDP datagram.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_udp(_:)
func Nw_protocol_metadata_is_udp(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_udp == nil {
		panic("Network: symbol nw_protocol_metadata_is_udp not loaded")
	}
	return _nw_protocol_metadata_is_udp(metadata)
}

var _nw_protocol_metadata_is_ws func(metadata Nw_protocol_metadata_t) bool

// Nw_protocol_metadata_is_ws checks whether a metadata object represents a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ws(_:)
func Nw_protocol_metadata_is_ws(metadata Nw_protocol_metadata_t) bool {
	if _nw_protocol_metadata_is_ws == nil {
		panic("Network: symbol nw_protocol_metadata_is_ws not loaded")
	}
	return _nw_protocol_metadata_is_ws(metadata)
}

var _nw_protocol_options_copy_definition func(options Nw_protocol_options_t) Nw_protocol_definition_t

// Nw_protocol_options_copy_definition accesses the protocol definition associated with the options object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_copy_definition(_:)
func Nw_protocol_options_copy_definition(options Nw_protocol_options_t) Nw_protocol_definition_t {
	if _nw_protocol_options_copy_definition == nil {
		panic("Network: symbol nw_protocol_options_copy_definition not loaded")
	}
	return _nw_protocol_options_copy_definition(options)
}

var _nw_protocol_options_is_quic func(options Nw_protocol_options_t) bool

// Nw_protocol_options_is_quic checks whether an options object uses the QUIC protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_is_quic(_:)
func Nw_protocol_options_is_quic(options Nw_protocol_options_t) bool {
	if _nw_protocol_options_is_quic == nil {
		panic("Network: symbol nw_protocol_options_is_quic not loaded")
	}
	return _nw_protocol_options_is_quic(options)
}

var _nw_protocol_stack_clear_application_protocols func(stack Nw_protocol_stack_t)

// Nw_protocol_stack_clear_application_protocols removes all application protocols from the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_clear_application_protocols(_:)
func Nw_protocol_stack_clear_application_protocols(stack Nw_protocol_stack_t) {
	if _nw_protocol_stack_clear_application_protocols == nil {
		panic("Network: symbol nw_protocol_stack_clear_application_protocols not loaded")
	}
	_nw_protocol_stack_clear_application_protocols(stack)
}

var _nw_protocol_stack_copy_internet_protocol func(stack Nw_protocol_stack_t) Nw_protocol_options_t

// Nw_protocol_stack_copy_internet_protocol accesses the protocol stack’s Internet Protocol options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_internet_protocol(_:)
func Nw_protocol_stack_copy_internet_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	if _nw_protocol_stack_copy_internet_protocol == nil {
		panic("Network: symbol nw_protocol_stack_copy_internet_protocol not loaded")
	}
	return _nw_protocol_stack_copy_internet_protocol(stack)
}

var _nw_protocol_stack_copy_transport_protocol func(stack Nw_protocol_stack_t) Nw_protocol_options_t

// Nw_protocol_stack_copy_transport_protocol accesses the options for the protocol stack’s transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_transport_protocol(_:)
func Nw_protocol_stack_copy_transport_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	if _nw_protocol_stack_copy_transport_protocol == nil {
		panic("Network: symbol nw_protocol_stack_copy_transport_protocol not loaded")
	}
	return _nw_protocol_stack_copy_transport_protocol(stack)
}

var _nw_protocol_stack_iterate_application_protocols func(stack Nw_protocol_stack_t, iterate_block Nw_protocol_stack_iterate_protocols_block_t)

// Nw_protocol_stack_iterate_application_protocols iterates through the array of application protocol options that will be used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_application_protocols(_:_:)
func Nw_protocol_stack_iterate_application_protocols(stack Nw_protocol_stack_t, iterate_block Nw_protocol_stack_iterate_protocols_block_t) {
	if _nw_protocol_stack_iterate_application_protocols == nil {
		panic("Network: symbol nw_protocol_stack_iterate_application_protocols not loaded")
	}
	_nw_protocol_stack_iterate_application_protocols(stack, iterate_block)
}

var _nw_protocol_stack_prepend_application_protocol func(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t)

// Nw_protocol_stack_prepend_application_protocol adds a protocol onto the top of the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_prepend_application_protocol(_:_:)
func Nw_protocol_stack_prepend_application_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	if _nw_protocol_stack_prepend_application_protocol == nil {
		panic("Network: symbol nw_protocol_stack_prepend_application_protocol not loaded")
	}
	_nw_protocol_stack_prepend_application_protocol(stack, protocol_)
}

var _nw_protocol_stack_set_transport_protocol func(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t)

// Nw_protocol_stack_set_transport_protocol replaces the protocol stack’s transport protocol with a new set of options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_set_transport_protocol(_:_:)
func Nw_protocol_stack_set_transport_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	if _nw_protocol_stack_set_transport_protocol == nil {
		panic("Network: symbol nw_protocol_stack_set_transport_protocol not loaded")
	}
	_nw_protocol_stack_set_transport_protocol(stack, protocol_)
}

var _nw_proxy_config_add_excluded_domain func(config Nw_proxy_config_t, excluded_domain string)

// Nw_proxy_config_add_excluded_domain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_excluded_domain(_:_:)
func Nw_proxy_config_add_excluded_domain(config Nw_proxy_config_t, excluded_domain string) {
	if _nw_proxy_config_add_excluded_domain == nil {
		panic("Network: symbol nw_proxy_config_add_excluded_domain not loaded")
	}
	_nw_proxy_config_add_excluded_domain(config, excluded_domain)
}

var _nw_proxy_config_add_match_domain func(config Nw_proxy_config_t, match_domain string)

// Nw_proxy_config_add_match_domain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_match_domain(_:_:)
func Nw_proxy_config_add_match_domain(config Nw_proxy_config_t, match_domain string) {
	if _nw_proxy_config_add_match_domain == nil {
		panic("Network: symbol nw_proxy_config_add_match_domain not loaded")
	}
	_nw_proxy_config_add_match_domain(config, match_domain)
}

var _nw_proxy_config_clear_excluded_domains func(config Nw_proxy_config_t)

// Nw_proxy_config_clear_excluded_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_excluded_domains(_:)
func Nw_proxy_config_clear_excluded_domains(config Nw_proxy_config_t) {
	if _nw_proxy_config_clear_excluded_domains == nil {
		panic("Network: symbol nw_proxy_config_clear_excluded_domains not loaded")
	}
	_nw_proxy_config_clear_excluded_domains(config)
}

var _nw_proxy_config_clear_match_domains func(config Nw_proxy_config_t)

// Nw_proxy_config_clear_match_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_match_domains(_:)
func Nw_proxy_config_clear_match_domains(config Nw_proxy_config_t) {
	if _nw_proxy_config_clear_match_domains == nil {
		panic("Network: symbol nw_proxy_config_clear_match_domains not loaded")
	}
	_nw_proxy_config_clear_match_domains(config)
}

var _nw_proxy_config_create_http_connect func(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) Nw_proxy_config_t

// Nw_proxy_config_create_http_connect initializes a legacy HTTP CONNECT configuration for a proxy server accessible using HTTP/1.1.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_http_connect(_:_:)
func Nw_proxy_config_create_http_connect(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) Nw_proxy_config_t {
	if _nw_proxy_config_create_http_connect == nil {
		panic("Network: symbol nw_proxy_config_create_http_connect not loaded")
	}
	return _nw_proxy_config_create_http_connect(proxy_endpoint, proxy_tls_options)
}

var _nw_proxy_config_create_oblivious_http func(relay Nw_relay_hop_t, relay_resource_path string, gateway_key_config *uint8, gateway_key_config_length uintptr) Nw_proxy_config_t

// Nw_proxy_config_create_oblivious_http initializes an Oblivious HTTP proxy configuration using a relay and a gateway.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_oblivious_http(_:_:_:_:)
func Nw_proxy_config_create_oblivious_http(relay Nw_relay_hop_t, relay_resource_path string, gateway_key_config *uint8, gateway_key_config_length uintptr) Nw_proxy_config_t {
	if _nw_proxy_config_create_oblivious_http == nil {
		panic("Network: symbol nw_proxy_config_create_oblivious_http not loaded")
	}
	return _nw_proxy_config_create_oblivious_http(relay, relay_resource_path, gateway_key_config, gateway_key_config_length)
}

var _nw_proxy_config_create_relay func(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) Nw_proxy_config_t

// Nw_proxy_config_create_relay initializes a proxy configuration with one or two relay hops.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_relay(_:_:)
func Nw_proxy_config_create_relay(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) Nw_proxy_config_t {
	if _nw_proxy_config_create_relay == nil {
		panic("Network: symbol nw_proxy_config_create_relay not loaded")
	}
	return _nw_proxy_config_create_relay(first_hop, second_hop)
}

var _nw_proxy_config_create_socksv5 func(proxy_endpoint Nw_endpoint_t) Nw_proxy_config_t

// Nw_proxy_config_create_socksv5 initializes a SOCKSv5 proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_socksv5(_:)
func Nw_proxy_config_create_socksv5(proxy_endpoint Nw_endpoint_t) Nw_proxy_config_t {
	if _nw_proxy_config_create_socksv5 == nil {
		panic("Network: symbol nw_proxy_config_create_socksv5 not loaded")
	}
	return _nw_proxy_config_create_socksv5(proxy_endpoint)
}

var _nw_proxy_config_enumerate_excluded_domains func(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t)

// Nw_proxy_config_enumerate_excluded_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_excluded_domains(_:_:)
func Nw_proxy_config_enumerate_excluded_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) {
	if _nw_proxy_config_enumerate_excluded_domains == nil {
		panic("Network: symbol nw_proxy_config_enumerate_excluded_domains not loaded")
	}
	_nw_proxy_config_enumerate_excluded_domains(config, enumerator)
}

var _nw_proxy_config_enumerate_match_domains func(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t)

// Nw_proxy_config_enumerate_match_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_match_domains(_:_:)
func Nw_proxy_config_enumerate_match_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) {
	if _nw_proxy_config_enumerate_match_domains == nil {
		panic("Network: symbol nw_proxy_config_enumerate_match_domains not loaded")
	}
	_nw_proxy_config_enumerate_match_domains(config, enumerator)
}

var _nw_proxy_config_get_failover_allowed func(proxy_config Nw_proxy_config_t) bool

// Nw_proxy_config_get_failover_allowed checks if a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_get_failover_allowed(_:)
func Nw_proxy_config_get_failover_allowed(proxy_config Nw_proxy_config_t) bool {
	if _nw_proxy_config_get_failover_allowed == nil {
		panic("Network: symbol nw_proxy_config_get_failover_allowed not loaded")
	}
	return _nw_proxy_config_get_failover_allowed(proxy_config)
}

var _nw_proxy_config_set_failover_allowed func(proxy_config Nw_proxy_config_t, failover_allowed bool)

// Nw_proxy_config_set_failover_allowed configures whether or not a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_failover_allowed(_:_:)
func Nw_proxy_config_set_failover_allowed(proxy_config Nw_proxy_config_t, failover_allowed bool) {
	if _nw_proxy_config_set_failover_allowed == nil {
		panic("Network: symbol nw_proxy_config_set_failover_allowed not loaded")
	}
	_nw_proxy_config_set_failover_allowed(proxy_config, failover_allowed)
}

var _nw_proxy_config_set_username_and_password func(proxy_config Nw_proxy_config_t, username string, password string)

// Nw_proxy_config_set_username_and_password sets a username and password to use as authentication for a proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_username_and_password(_:_:_:)
func Nw_proxy_config_set_username_and_password(proxy_config Nw_proxy_config_t, username string, password string) {
	if _nw_proxy_config_set_username_and_password == nil {
		panic("Network: symbol nw_proxy_config_set_username_and_password not loaded")
	}
	_nw_proxy_config_set_username_and_password(proxy_config, username, password)
}

var _nw_quic_add_tls_application_protocol func(options Nw_protocol_options_t, application_protocol string)

// Nw_quic_add_tls_application_protocol adds a supported Application-Layer Protocol Negotiation value.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_add_tls_application_protocol(_:_:)
func Nw_quic_add_tls_application_protocol(options Nw_protocol_options_t, application_protocol string) {
	if _nw_quic_add_tls_application_protocol == nil {
		panic("Network: symbol nw_quic_add_tls_application_protocol not loaded")
	}
	_nw_quic_add_tls_application_protocol(options, application_protocol)
}

var _nw_quic_copy_sec_protocol_metadata func(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t

// Nw_quic_copy_sec_protocol_metadata accesses the result of the QUIC handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_metadata(_:)
func Nw_quic_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t {
	if _nw_quic_copy_sec_protocol_metadata == nil {
		panic("Network: symbol nw_quic_copy_sec_protocol_metadata not loaded")
	}
	return _nw_quic_copy_sec_protocol_metadata(metadata)
}

var _nw_quic_copy_sec_protocol_options func(options Nw_protocol_options_t) security.Sec_protocol_options_t

// Nw_quic_copy_sec_protocol_options accesses the handshake security options QUIC will use.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_options(_:)
func Nw_quic_copy_sec_protocol_options(options Nw_protocol_options_t) security.Sec_protocol_options_t {
	if _nw_quic_copy_sec_protocol_options == nil {
		panic("Network: symbol nw_quic_copy_sec_protocol_options not loaded")
	}
	return _nw_quic_copy_sec_protocol_options(options)
}

var _nw_quic_create_options func() Nw_protocol_options_t

// Nw_quic_create_options initializes a default set of QUIC connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_create_options()
func Nw_quic_create_options() Nw_protocol_options_t {
	if _nw_quic_create_options == nil {
		panic("Network: symbol nw_quic_create_options not loaded")
	}
	return _nw_quic_create_options()
}

var _nw_quic_get_application_error func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_application_error accesses the QUIC application error code received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error(_:)
func Nw_quic_get_application_error(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_application_error == nil {
		panic("Network: symbol nw_quic_get_application_error not loaded")
	}
	return _nw_quic_get_application_error(metadata)
}

var _nw_quic_get_application_error_reason func(metadata Nw_protocol_metadata_t) *byte

// Nw_quic_get_application_error_reason accesses the QUIC application error reason received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error_reason(_:)
func Nw_quic_get_application_error_reason(metadata Nw_protocol_metadata_t) *byte {
	if _nw_quic_get_application_error_reason == nil {
		panic("Network: symbol nw_quic_get_application_error_reason not loaded")
	}
	return _nw_quic_get_application_error_reason(metadata)
}

var _nw_quic_get_idle_timeout func(options Nw_protocol_options_t) uint32

// Nw_quic_get_idle_timeout accesses the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_idle_timeout(_:)
func Nw_quic_get_idle_timeout(options Nw_protocol_options_t) uint32 {
	if _nw_quic_get_idle_timeout == nil {
		panic("Network: symbol nw_quic_get_idle_timeout not loaded")
	}
	return _nw_quic_get_idle_timeout(options)
}

var _nw_quic_get_initial_max_data func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_data accesses a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_data(_:)
func Nw_quic_get_initial_max_data(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_data == nil {
		panic("Network: symbol nw_quic_get_initial_max_data not loaded")
	}
	return _nw_quic_get_initial_max_data(options)
}

var _nw_quic_get_initial_max_stream_data_bidirectional_local func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_stream_data_bidirectional_local accesses a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_local(_:)
func Nw_quic_get_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_stream_data_bidirectional_local == nil {
		panic("Network: symbol nw_quic_get_initial_max_stream_data_bidirectional_local not loaded")
	}
	return _nw_quic_get_initial_max_stream_data_bidirectional_local(options)
}

var _nw_quic_get_initial_max_stream_data_bidirectional_remote func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_stream_data_bidirectional_remote accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_remote(_:)
func Nw_quic_get_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_stream_data_bidirectional_remote == nil {
		panic("Network: symbol nw_quic_get_initial_max_stream_data_bidirectional_remote not loaded")
	}
	return _nw_quic_get_initial_max_stream_data_bidirectional_remote(options)
}

var _nw_quic_get_initial_max_stream_data_unidirectional func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_stream_data_unidirectional accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_unidirectional(_:)
func Nw_quic_get_initial_max_stream_data_unidirectional(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_stream_data_unidirectional == nil {
		panic("Network: symbol nw_quic_get_initial_max_stream_data_unidirectional not loaded")
	}
	return _nw_quic_get_initial_max_stream_data_unidirectional(options)
}

var _nw_quic_get_initial_max_streams_bidirectional func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_streams_bidirectional accesses a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_bidirectional(_:)
func Nw_quic_get_initial_max_streams_bidirectional(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_streams_bidirectional == nil {
		panic("Network: symbol nw_quic_get_initial_max_streams_bidirectional not loaded")
	}
	return _nw_quic_get_initial_max_streams_bidirectional(options)
}

var _nw_quic_get_initial_max_streams_unidirectional func(options Nw_protocol_options_t) uint64

// Nw_quic_get_initial_max_streams_unidirectional accesses a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_unidirectional(_:)
func Nw_quic_get_initial_max_streams_unidirectional(options Nw_protocol_options_t) uint64 {
	if _nw_quic_get_initial_max_streams_unidirectional == nil {
		panic("Network: symbol nw_quic_get_initial_max_streams_unidirectional not loaded")
	}
	return _nw_quic_get_initial_max_streams_unidirectional(options)
}

var _nw_quic_get_keepalive_interval func(metadata Nw_protocol_metadata_t) uint16

// Nw_quic_get_keepalive_interval accesses the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_keepalive_interval(_:)
func Nw_quic_get_keepalive_interval(metadata Nw_protocol_metadata_t) uint16 {
	if _nw_quic_get_keepalive_interval == nil {
		panic("Network: symbol nw_quic_get_keepalive_interval not loaded")
	}
	return _nw_quic_get_keepalive_interval(metadata)
}

var _nw_quic_get_local_max_streams_bidirectional func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_local_max_streams_bidirectional accesses the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_bidirectional(_:)
func Nw_quic_get_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_local_max_streams_bidirectional == nil {
		panic("Network: symbol nw_quic_get_local_max_streams_bidirectional not loaded")
	}
	return _nw_quic_get_local_max_streams_bidirectional(metadata)
}

var _nw_quic_get_local_max_streams_unidirectional func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_local_max_streams_unidirectional accesses the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_unidirectional(_:)
func Nw_quic_get_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_local_max_streams_unidirectional == nil {
		panic("Network: symbol nw_quic_get_local_max_streams_unidirectional not loaded")
	}
	return _nw_quic_get_local_max_streams_unidirectional(metadata)
}

var _nw_quic_get_max_datagram_frame_size func(options Nw_protocol_options_t) uint16

// Nw_quic_get_max_datagram_frame_size accesses a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_datagram_frame_size(_:)
func Nw_quic_get_max_datagram_frame_size(options Nw_protocol_options_t) uint16 {
	if _nw_quic_get_max_datagram_frame_size == nil {
		panic("Network: symbol nw_quic_get_max_datagram_frame_size not loaded")
	}
	return _nw_quic_get_max_datagram_frame_size(options)
}

var _nw_quic_get_max_udp_payload_size func(options Nw_protocol_options_t) uint16

// Nw_quic_get_max_udp_payload_size accesses the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_udp_payload_size(_:)
func Nw_quic_get_max_udp_payload_size(options Nw_protocol_options_t) uint16 {
	if _nw_quic_get_max_udp_payload_size == nil {
		panic("Network: symbol nw_quic_get_max_udp_payload_size not loaded")
	}
	return _nw_quic_get_max_udp_payload_size(options)
}

var _nw_quic_get_remote_idle_timeout func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_remote_idle_timeout accesses the idle timeout value from the peer’s transport parameters, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_idle_timeout(_:)
func Nw_quic_get_remote_idle_timeout(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_remote_idle_timeout == nil {
		panic("Network: symbol nw_quic_get_remote_idle_timeout not loaded")
	}
	return _nw_quic_get_remote_idle_timeout(metadata)
}

var _nw_quic_get_remote_max_streams_bidirectional func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_remote_max_streams_bidirectional accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_bidirectional(_:)
func Nw_quic_get_remote_max_streams_bidirectional(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_remote_max_streams_bidirectional == nil {
		panic("Network: symbol nw_quic_get_remote_max_streams_bidirectional not loaded")
	}
	return _nw_quic_get_remote_max_streams_bidirectional(metadata)
}

var _nw_quic_get_remote_max_streams_unidirectional func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_remote_max_streams_unidirectional accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_unidirectional(_:)
func Nw_quic_get_remote_max_streams_unidirectional(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_remote_max_streams_unidirectional == nil {
		panic("Network: symbol nw_quic_get_remote_max_streams_unidirectional not loaded")
	}
	return _nw_quic_get_remote_max_streams_unidirectional(metadata)
}

var _nw_quic_get_stream_application_error func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_stream_application_error accesses the QUIC application error code received from the peer for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_application_error(_:)
func Nw_quic_get_stream_application_error(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_stream_application_error == nil {
		panic("Network: symbol nw_quic_get_stream_application_error not loaded")
	}
	return _nw_quic_get_stream_application_error(metadata)
}

var _nw_quic_get_stream_id func(metadata Nw_protocol_metadata_t) uint64

// Nw_quic_get_stream_id accesses the QUIC stream identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_id(_:)
func Nw_quic_get_stream_id(metadata Nw_protocol_metadata_t) uint64 {
	if _nw_quic_get_stream_id == nil {
		panic("Network: symbol nw_quic_get_stream_id not loaded")
	}
	return _nw_quic_get_stream_id(metadata)
}

var _nw_quic_get_stream_is_datagram func(options Nw_protocol_options_t) bool

// Nw_quic_get_stream_is_datagram checks if a QUIC stream is a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_datagram(_:)
func Nw_quic_get_stream_is_datagram(options Nw_protocol_options_t) bool {
	if _nw_quic_get_stream_is_datagram == nil {
		panic("Network: symbol nw_quic_get_stream_is_datagram not loaded")
	}
	return _nw_quic_get_stream_is_datagram(options)
}

var _nw_quic_get_stream_is_unidirectional func(options Nw_protocol_options_t) bool

// Nw_quic_get_stream_is_unidirectional checks if a QUIC stream is unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_unidirectional(_:)
func Nw_quic_get_stream_is_unidirectional(options Nw_protocol_options_t) bool {
	if _nw_quic_get_stream_is_unidirectional == nil {
		panic("Network: symbol nw_quic_get_stream_is_unidirectional not loaded")
	}
	return _nw_quic_get_stream_is_unidirectional(options)
}

var _nw_quic_get_stream_type func(stream_metadata Nw_protocol_metadata_t) uint8

// Nw_quic_get_stream_type accesses the stream type of the QUIC stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_type(_:)
func Nw_quic_get_stream_type(stream_metadata Nw_protocol_metadata_t) uint8 {
	if _nw_quic_get_stream_type == nil {
		panic("Network: symbol nw_quic_get_stream_type not loaded")
	}
	return _nw_quic_get_stream_type(stream_metadata)
}

var _nw_quic_get_stream_usable_datagram_frame_size func(metadata Nw_protocol_metadata_t) uint16

// Nw_quic_get_stream_usable_datagram_frame_size accesses the maximum usable size of a datagram frame on a QUIC datagram flow.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_usable_datagram_frame_size(_:)
func Nw_quic_get_stream_usable_datagram_frame_size(metadata Nw_protocol_metadata_t) uint16 {
	if _nw_quic_get_stream_usable_datagram_frame_size == nil {
		panic("Network: symbol nw_quic_get_stream_usable_datagram_frame_size not loaded")
	}
	return _nw_quic_get_stream_usable_datagram_frame_size(metadata)
}

var _nw_quic_set_application_error func(metadata Nw_protocol_metadata_t, application_error uint64, reason string)

// Nw_quic_set_application_error sets the QUIC application error code to send for the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_application_error(_:_:_:)
func Nw_quic_set_application_error(metadata Nw_protocol_metadata_t, application_error uint64, reason string) {
	if _nw_quic_set_application_error == nil {
		panic("Network: symbol nw_quic_set_application_error not loaded")
	}
	_nw_quic_set_application_error(metadata, application_error, reason)
}

var _nw_quic_set_idle_timeout func(options Nw_protocol_options_t, idle_timeout uint32)

// Nw_quic_set_idle_timeout sets the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_idle_timeout(_:_:)
func Nw_quic_set_idle_timeout(options Nw_protocol_options_t, idle_timeout uint32) {
	if _nw_quic_set_idle_timeout == nil {
		panic("Network: symbol nw_quic_set_idle_timeout not loaded")
	}
	_nw_quic_set_idle_timeout(options, idle_timeout)
}

var _nw_quic_set_initial_max_data func(options Nw_protocol_options_t, initial_max_data uint64)

// Nw_quic_set_initial_max_data sets a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_data(_:_:)
func Nw_quic_set_initial_max_data(options Nw_protocol_options_t, initial_max_data uint64) {
	if _nw_quic_set_initial_max_data == nil {
		panic("Network: symbol nw_quic_set_initial_max_data not loaded")
	}
	_nw_quic_set_initial_max_data(options, initial_max_data)
}

var _nw_quic_set_initial_max_stream_data_bidirectional_local func(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_local uint64)

// Nw_quic_set_initial_max_stream_data_bidirectional_local sets a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_local(_:_:)
func Nw_quic_set_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_local uint64) {
	if _nw_quic_set_initial_max_stream_data_bidirectional_local == nil {
		panic("Network: symbol nw_quic_set_initial_max_stream_data_bidirectional_local not loaded")
	}
	_nw_quic_set_initial_max_stream_data_bidirectional_local(options, initial_max_stream_data_bidirectional_local)
}

var _nw_quic_set_initial_max_stream_data_bidirectional_remote func(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64)

// Nw_quic_set_initial_max_stream_data_bidirectional_remote sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_remote(_:_:)
func Nw_quic_set_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64) {
	if _nw_quic_set_initial_max_stream_data_bidirectional_remote == nil {
		panic("Network: symbol nw_quic_set_initial_max_stream_data_bidirectional_remote not loaded")
	}
	_nw_quic_set_initial_max_stream_data_bidirectional_remote(options, initial_max_stream_data_bidirectional_remote)
}

var _nw_quic_set_initial_max_stream_data_unidirectional func(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64)

// Nw_quic_set_initial_max_stream_data_unidirectional sets a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_unidirectional(_:_:)
func Nw_quic_set_initial_max_stream_data_unidirectional(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64) {
	if _nw_quic_set_initial_max_stream_data_unidirectional == nil {
		panic("Network: symbol nw_quic_set_initial_max_stream_data_unidirectional not loaded")
	}
	_nw_quic_set_initial_max_stream_data_unidirectional(options, initial_max_stream_data_unidirectional)
}

var _nw_quic_set_initial_max_streams_bidirectional func(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64)

// Nw_quic_set_initial_max_streams_bidirectional sets a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_bidirectional(_:_:)
func Nw_quic_set_initial_max_streams_bidirectional(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64) {
	if _nw_quic_set_initial_max_streams_bidirectional == nil {
		panic("Network: symbol nw_quic_set_initial_max_streams_bidirectional not loaded")
	}
	_nw_quic_set_initial_max_streams_bidirectional(options, initial_max_streams_bidirectional)
}

var _nw_quic_set_initial_max_streams_unidirectional func(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64)

// Nw_quic_set_initial_max_streams_unidirectional sets a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_unidirectional(_:_:)
func Nw_quic_set_initial_max_streams_unidirectional(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64) {
	if _nw_quic_set_initial_max_streams_unidirectional == nil {
		panic("Network: symbol nw_quic_set_initial_max_streams_unidirectional not loaded")
	}
	_nw_quic_set_initial_max_streams_unidirectional(options, initial_max_streams_unidirectional)
}

var _nw_quic_set_keepalive_interval func(metadata Nw_protocol_metadata_t, keepalive_interval uint16)

// Nw_quic_set_keepalive_interval sets the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_keepalive_interval(_:_:)
func Nw_quic_set_keepalive_interval(metadata Nw_protocol_metadata_t, keepalive_interval uint16) {
	if _nw_quic_set_keepalive_interval == nil {
		panic("Network: symbol nw_quic_set_keepalive_interval not loaded")
	}
	_nw_quic_set_keepalive_interval(metadata, keepalive_interval)
}

var _nw_quic_set_local_max_streams_bidirectional func(metadata Nw_protocol_metadata_t, max_streams_bidirectional uint64)

// Nw_quic_set_local_max_streams_bidirectional sets the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_bidirectional(_:_:)
func Nw_quic_set_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t, max_streams_bidirectional uint64) {
	if _nw_quic_set_local_max_streams_bidirectional == nil {
		panic("Network: symbol nw_quic_set_local_max_streams_bidirectional not loaded")
	}
	_nw_quic_set_local_max_streams_bidirectional(metadata, max_streams_bidirectional)
}

var _nw_quic_set_local_max_streams_unidirectional func(metadata Nw_protocol_metadata_t, max_streams_unidirectional uint64)

// Nw_quic_set_local_max_streams_unidirectional sets the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_unidirectional(_:_:)
func Nw_quic_set_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t, max_streams_unidirectional uint64) {
	if _nw_quic_set_local_max_streams_unidirectional == nil {
		panic("Network: symbol nw_quic_set_local_max_streams_unidirectional not loaded")
	}
	_nw_quic_set_local_max_streams_unidirectional(metadata, max_streams_unidirectional)
}

var _nw_quic_set_max_datagram_frame_size func(options Nw_protocol_options_t, max_datagram_frame_size uint16)

// Nw_quic_set_max_datagram_frame_size sets a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_datagram_frame_size(_:_:)
func Nw_quic_set_max_datagram_frame_size(options Nw_protocol_options_t, max_datagram_frame_size uint16) {
	if _nw_quic_set_max_datagram_frame_size == nil {
		panic("Network: symbol nw_quic_set_max_datagram_frame_size not loaded")
	}
	_nw_quic_set_max_datagram_frame_size(options, max_datagram_frame_size)
}

var _nw_quic_set_max_udp_payload_size func(options Nw_protocol_options_t, max_udp_payload_size uint16)

// Nw_quic_set_max_udp_payload_size sets the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_udp_payload_size(_:_:)
func Nw_quic_set_max_udp_payload_size(options Nw_protocol_options_t, max_udp_payload_size uint16) {
	if _nw_quic_set_max_udp_payload_size == nil {
		panic("Network: symbol nw_quic_set_max_udp_payload_size not loaded")
	}
	_nw_quic_set_max_udp_payload_size(options, max_udp_payload_size)
}

var _nw_quic_set_stream_application_error func(metadata Nw_protocol_metadata_t, application_error uint64)

// Nw_quic_set_stream_application_error sets the QUIC application error code to send for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_application_error(_:_:)
func Nw_quic_set_stream_application_error(metadata Nw_protocol_metadata_t, application_error uint64) {
	if _nw_quic_set_stream_application_error == nil {
		panic("Network: symbol nw_quic_set_stream_application_error not loaded")
	}
	_nw_quic_set_stream_application_error(metadata, application_error)
}

var _nw_quic_set_stream_is_datagram func(options Nw_protocol_options_t, is_datagram bool)

// Nw_quic_set_stream_is_datagram configures a QUIC stream as a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_datagram(_:_:)
func Nw_quic_set_stream_is_datagram(options Nw_protocol_options_t, is_datagram bool) {
	if _nw_quic_set_stream_is_datagram == nil {
		panic("Network: symbol nw_quic_set_stream_is_datagram not loaded")
	}
	_nw_quic_set_stream_is_datagram(options, is_datagram)
}

var _nw_quic_set_stream_is_unidirectional func(options Nw_protocol_options_t, is_unidirectional bool)

// Nw_quic_set_stream_is_unidirectional configures a QUIC stream as unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_unidirectional(_:_:)
func Nw_quic_set_stream_is_unidirectional(options Nw_protocol_options_t, is_unidirectional bool) {
	if _nw_quic_set_stream_is_unidirectional == nil {
		panic("Network: symbol nw_quic_set_stream_is_unidirectional not loaded")
	}
	_nw_quic_set_stream_is_unidirectional(options, is_unidirectional)
}

var _nw_relay_hop_add_additional_http_header_field func(relay_hop Nw_relay_hop_t, field_name string, field_value string)

// Nw_relay_hop_add_additional_http_header_field adds an HTTP header name and value pair to send as part of [CONNECT] requests to the relay.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_add_additional_http_header_field(_:_:_:)
func Nw_relay_hop_add_additional_http_header_field(relay_hop Nw_relay_hop_t, field_name string, field_value string) {
	if _nw_relay_hop_add_additional_http_header_field == nil {
		panic("Network: symbol nw_relay_hop_add_additional_http_header_field not loaded")
	}
	_nw_relay_hop_add_additional_http_header_field(relay_hop, field_name, field_value)
}

var _nw_relay_hop_create func(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) Nw_relay_hop_t

// Nw_relay_hop_create creates a configuration for a secure relay accessible using HTTP/3, with an optional HTTP/2 fallback.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_create(_:_:_:)
func Nw_relay_hop_create(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) Nw_relay_hop_t {
	if _nw_relay_hop_create == nil {
		panic("Network: symbol nw_relay_hop_create not loaded")
	}
	return _nw_relay_hop_create(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options)
}

var _nw_release func(obj unsafe.Pointer)

// Nw_release releases a reference count on a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_release
func Nw_release(obj unsafe.Pointer) {
	if _nw_release == nil {
		panic("Network: symbol nw_release not loaded")
	}
	_nw_release(obj)
}

var _nw_resolution_report_copy_preferred_endpoint func(resolution_report Nw_resolution_report_t) Nw_endpoint_t

// Nw_resolution_report_copy_preferred_endpoint accesses the resolved endpoint that the connection used for its first connection attempt.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_preferred_endpoint(_:)
func Nw_resolution_report_copy_preferred_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	if _nw_resolution_report_copy_preferred_endpoint == nil {
		panic("Network: symbol nw_resolution_report_copy_preferred_endpoint not loaded")
	}
	return _nw_resolution_report_copy_preferred_endpoint(resolution_report)
}

var _nw_resolution_report_copy_successful_endpoint func(resolution_report Nw_resolution_report_t) Nw_endpoint_t

// Nw_resolution_report_copy_successful_endpoint accesses the resolved endpoint that led to the established connection.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_successful_endpoint(_:)
func Nw_resolution_report_copy_successful_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	if _nw_resolution_report_copy_successful_endpoint == nil {
		panic("Network: symbol nw_resolution_report_copy_successful_endpoint not loaded")
	}
	return _nw_resolution_report_copy_successful_endpoint(resolution_report)
}

var _nw_resolution_report_get_endpoint_count func(resolution_report Nw_resolution_report_t) uint32

// Nw_resolution_report_get_endpoint_count accesses the number of endpoints resolved in this step.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_endpoint_count(_:)
func Nw_resolution_report_get_endpoint_count(resolution_report Nw_resolution_report_t) uint32 {
	if _nw_resolution_report_get_endpoint_count == nil {
		panic("Network: symbol nw_resolution_report_get_endpoint_count not loaded")
	}
	return _nw_resolution_report_get_endpoint_count(resolution_report)
}

var _nw_resolution_report_get_milliseconds func(resolution_report Nw_resolution_report_t) uint64

// Nw_resolution_report_get_milliseconds accesses the duration of this resolution step, from when the query was issued to when the response was complete.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_milliseconds(_:)
func Nw_resolution_report_get_milliseconds(resolution_report Nw_resolution_report_t) uint64 {
	if _nw_resolution_report_get_milliseconds == nil {
		panic("Network: symbol nw_resolution_report_get_milliseconds not loaded")
	}
	return _nw_resolution_report_get_milliseconds(resolution_report)
}

var _nw_resolution_report_get_protocol func(resolution_report Nw_resolution_report_t) unsafe.Pointer

// Nw_resolution_report_get_protocol accesses the transport protocol your connection used for DNS resolution.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_protocol(_:)
func Nw_resolution_report_get_protocol(resolution_report Nw_resolution_report_t) unsafe.Pointer {
	if _nw_resolution_report_get_protocol == nil {
		panic("Network: symbol nw_resolution_report_get_protocol not loaded")
	}
	return _nw_resolution_report_get_protocol(resolution_report)
}

var _nw_resolution_report_get_source func(resolution_report Nw_resolution_report_t) unsafe.Pointer

// Nw_resolution_report_get_source accesses the source of the DNS response.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_source(_:)
func Nw_resolution_report_get_source(resolution_report Nw_resolution_report_t) unsafe.Pointer {
	if _nw_resolution_report_get_source == nil {
		panic("Network: symbol nw_resolution_report_get_source not loaded")
	}
	return _nw_resolution_report_get_source(resolution_report)
}

var _nw_resolver_config_add_server_address func(config Nw_resolver_config_t, server_address Nw_endpoint_t)

// Nw_resolver_config_add_server_address provides a well-known DNS server address to use instead of looking up the address dynamically.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_add_server_address(_:_:)
func Nw_resolver_config_add_server_address(config Nw_resolver_config_t, server_address Nw_endpoint_t) {
	if _nw_resolver_config_add_server_address == nil {
		panic("Network: symbol nw_resolver_config_add_server_address not loaded")
	}
	_nw_resolver_config_add_server_address(config, server_address)
}

var _nw_resolver_config_create_https func(url_endpoint Nw_endpoint_t) Nw_resolver_config_t

// Nw_resolver_config_create_https initializes a DNS-over-HTTPS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_https(_:)
func Nw_resolver_config_create_https(url_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	if _nw_resolver_config_create_https == nil {
		panic("Network: symbol nw_resolver_config_create_https not loaded")
	}
	return _nw_resolver_config_create_https(url_endpoint)
}

var _nw_resolver_config_create_tls func(server_endpoint Nw_endpoint_t) Nw_resolver_config_t

// Nw_resolver_config_create_tls initializes a DNS-over-TLS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_tls(_:)
func Nw_resolver_config_create_tls(server_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	if _nw_resolver_config_create_tls == nil {
		panic("Network: symbol nw_resolver_config_create_tls not loaded")
	}
	return _nw_resolver_config_create_tls(server_endpoint)
}

var _nw_retain func(obj unsafe.Pointer) unsafe.Pointer

// Nw_retain adds a reference count to a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_retain
func Nw_retain(obj unsafe.Pointer) unsafe.Pointer {
	if _nw_retain == nil {
		panic("Network: symbol nw_retain not loaded")
	}
	return _nw_retain(obj)
}

var _nw_tcp_create_options func() Nw_protocol_options_t

// Nw_tcp_create_options initializes a default set of TCP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_create_options()
func Nw_tcp_create_options() Nw_protocol_options_t {
	if _nw_tcp_create_options == nil {
		panic("Network: symbol nw_tcp_create_options not loaded")
	}
	return _nw_tcp_create_options()
}

var _nw_tcp_get_available_receive_buffer func(metadata Nw_protocol_metadata_t) uint32

// Nw_tcp_get_available_receive_buffer accesses the number of available bytes in the TCP receive buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_receive_buffer(_:)
func Nw_tcp_get_available_receive_buffer(metadata Nw_protocol_metadata_t) uint32 {
	if _nw_tcp_get_available_receive_buffer == nil {
		panic("Network: symbol nw_tcp_get_available_receive_buffer not loaded")
	}
	return _nw_tcp_get_available_receive_buffer(metadata)
}

var _nw_tcp_get_available_send_buffer func(metadata Nw_protocol_metadata_t) uint32

// Nw_tcp_get_available_send_buffer accesses the number of available bytes in the TCP send buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_send_buffer(_:)
func Nw_tcp_get_available_send_buffer(metadata Nw_protocol_metadata_t) uint32 {
	if _nw_tcp_get_available_send_buffer == nil {
		panic("Network: symbol nw_tcp_get_available_send_buffer not loaded")
	}
	return _nw_tcp_get_available_send_buffer(metadata)
}

var _nw_tcp_options_set_connection_timeout func(options Nw_protocol_options_t, connection_timeout uint32)

// Nw_tcp_options_set_connection_timeout sets the number of seconds that TCP waits before timing out its handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_connection_timeout(_:_:)
func Nw_tcp_options_set_connection_timeout(options Nw_protocol_options_t, connection_timeout uint32) {
	if _nw_tcp_options_set_connection_timeout == nil {
		panic("Network: symbol nw_tcp_options_set_connection_timeout not loaded")
	}
	_nw_tcp_options_set_connection_timeout(options, connection_timeout)
}

var _nw_tcp_options_set_disable_ack_stretching func(options Nw_protocol_options_t, disable_ack_stretching bool)

// Nw_tcp_options_set_disable_ack_stretching disables TCP acknowledgment stretching.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ack_stretching(_:_:)
func Nw_tcp_options_set_disable_ack_stretching(options Nw_protocol_options_t, disable_ack_stretching bool) {
	if _nw_tcp_options_set_disable_ack_stretching == nil {
		panic("Network: symbol nw_tcp_options_set_disable_ack_stretching not loaded")
	}
	_nw_tcp_options_set_disable_ack_stretching(options, disable_ack_stretching)
}

var _nw_tcp_options_set_disable_ecn func(options Nw_protocol_options_t, disable_ecn bool)

// Nw_tcp_options_set_disable_ecn disables negotiation of Explicit Congestion Notification markings.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ecn(_:_:)
func Nw_tcp_options_set_disable_ecn(options Nw_protocol_options_t, disable_ecn bool) {
	if _nw_tcp_options_set_disable_ecn == nil {
		panic("Network: symbol nw_tcp_options_set_disable_ecn not loaded")
	}
	_nw_tcp_options_set_disable_ecn(options, disable_ecn)
}

var _nw_tcp_options_set_enable_fast_open func(options Nw_protocol_options_t, enable_fast_open bool)

// Nw_tcp_options_set_enable_fast_open enables TCP Fast Open on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_fast_open(_:_:)
func Nw_tcp_options_set_enable_fast_open(options Nw_protocol_options_t, enable_fast_open bool) {
	if _nw_tcp_options_set_enable_fast_open == nil {
		panic("Network: symbol nw_tcp_options_set_enable_fast_open not loaded")
	}
	_nw_tcp_options_set_enable_fast_open(options, enable_fast_open)
}

var _nw_tcp_options_set_enable_keepalive func(options Nw_protocol_options_t, enable_keepalive bool)

// Nw_tcp_options_set_enable_keepalive enables TCP keepalives.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_keepalive(_:_:)
func Nw_tcp_options_set_enable_keepalive(options Nw_protocol_options_t, enable_keepalive bool) {
	if _nw_tcp_options_set_enable_keepalive == nil {
		panic("Network: symbol nw_tcp_options_set_enable_keepalive not loaded")
	}
	_nw_tcp_options_set_enable_keepalive(options, enable_keepalive)
}

var _nw_tcp_options_set_keepalive_count func(options Nw_protocol_options_t, keepalive_count uint32)

// Nw_tcp_options_set_keepalive_count sets the number of keepalive probes that TCP sends before terminating the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_count(_:_:)
func Nw_tcp_options_set_keepalive_count(options Nw_protocol_options_t, keepalive_count uint32) {
	if _nw_tcp_options_set_keepalive_count == nil {
		panic("Network: symbol nw_tcp_options_set_keepalive_count not loaded")
	}
	_nw_tcp_options_set_keepalive_count(options, keepalive_count)
}

var _nw_tcp_options_set_keepalive_idle_time func(options Nw_protocol_options_t, keepalive_idle_time uint32)

// Nw_tcp_options_set_keepalive_idle_time sets the number of seconds of idleness that TCP waits before sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_idle_time(_:_:)
func Nw_tcp_options_set_keepalive_idle_time(options Nw_protocol_options_t, keepalive_idle_time uint32) {
	if _nw_tcp_options_set_keepalive_idle_time == nil {
		panic("Network: symbol nw_tcp_options_set_keepalive_idle_time not loaded")
	}
	_nw_tcp_options_set_keepalive_idle_time(options, keepalive_idle_time)
}

var _nw_tcp_options_set_keepalive_interval func(options Nw_protocol_options_t, keepalive_interval uint32)

// Nw_tcp_options_set_keepalive_interval sets the number of seconds that TCP waits between sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_interval(_:_:)
func Nw_tcp_options_set_keepalive_interval(options Nw_protocol_options_t, keepalive_interval uint32) {
	if _nw_tcp_options_set_keepalive_interval == nil {
		panic("Network: symbol nw_tcp_options_set_keepalive_interval not loaded")
	}
	_nw_tcp_options_set_keepalive_interval(options, keepalive_interval)
}

var _nw_tcp_options_set_maximum_segment_size func(options Nw_protocol_options_t, maximum_segment_size uint32)

// Nw_tcp_options_set_maximum_segment_size sets TCP’s maximum segment size in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_maximum_segment_size(_:_:)
func Nw_tcp_options_set_maximum_segment_size(options Nw_protocol_options_t, maximum_segment_size uint32) {
	if _nw_tcp_options_set_maximum_segment_size == nil {
		panic("Network: symbol nw_tcp_options_set_maximum_segment_size not loaded")
	}
	_nw_tcp_options_set_maximum_segment_size(options, maximum_segment_size)
}

var _nw_tcp_options_set_multipath_force_version func(options Nw_protocol_options_t, multipath_force_version unsafe.Pointer)

// Nw_tcp_options_set_multipath_force_version.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_multipath_force_version(_:_:)
func Nw_tcp_options_set_multipath_force_version(options Nw_protocol_options_t, multipath_force_version unsafe.Pointer) {
	if _nw_tcp_options_set_multipath_force_version == nil {
		panic("Network: symbol nw_tcp_options_set_multipath_force_version not loaded")
	}
	_nw_tcp_options_set_multipath_force_version(options, multipath_force_version)
}

var _nw_tcp_options_set_no_delay func(options Nw_protocol_options_t, no_delay bool)

// Nw_tcp_options_set_no_delay disables Nagle’s algorithm for TCP.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_delay(_:_:)
func Nw_tcp_options_set_no_delay(options Nw_protocol_options_t, no_delay bool) {
	if _nw_tcp_options_set_no_delay == nil {
		panic("Network: symbol nw_tcp_options_set_no_delay not loaded")
	}
	_nw_tcp_options_set_no_delay(options, no_delay)
}

var _nw_tcp_options_set_no_options func(options Nw_protocol_options_t, no_options bool)

// Nw_tcp_options_set_no_options sets TCP into no-options mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_options(_:_:)
func Nw_tcp_options_set_no_options(options Nw_protocol_options_t, no_options bool) {
	if _nw_tcp_options_set_no_options == nil {
		panic("Network: symbol nw_tcp_options_set_no_options not loaded")
	}
	_nw_tcp_options_set_no_options(options, no_options)
}

var _nw_tcp_options_set_no_push func(options Nw_protocol_options_t, no_push bool)

// Nw_tcp_options_set_no_push sets TCP into no-push mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_push(_:_:)
func Nw_tcp_options_set_no_push(options Nw_protocol_options_t, no_push bool) {
	if _nw_tcp_options_set_no_push == nil {
		panic("Network: symbol nw_tcp_options_set_no_push not loaded")
	}
	_nw_tcp_options_set_no_push(options, no_push)
}

var _nw_tcp_options_set_persist_timeout func(options Nw_protocol_options_t, persist_timeout uint32)

// Nw_tcp_options_set_persist_timeout sets the TCP persist timeout in seconds, as defined by RFC 6429.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_persist_timeout(_:_:)
func Nw_tcp_options_set_persist_timeout(options Nw_protocol_options_t, persist_timeout uint32) {
	if _nw_tcp_options_set_persist_timeout == nil {
		panic("Network: symbol nw_tcp_options_set_persist_timeout not loaded")
	}
	_nw_tcp_options_set_persist_timeout(options, persist_timeout)
}

var _nw_tcp_options_set_retransmit_connection_drop_time func(options Nw_protocol_options_t, retransmit_connection_drop_time uint32)

// Nw_tcp_options_set_retransmit_connection_drop_time sets the number of seconds that TCP waits between retransmission attempts.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_connection_drop_time(_:_:)
func Nw_tcp_options_set_retransmit_connection_drop_time(options Nw_protocol_options_t, retransmit_connection_drop_time uint32) {
	if _nw_tcp_options_set_retransmit_connection_drop_time == nil {
		panic("Network: symbol nw_tcp_options_set_retransmit_connection_drop_time not loaded")
	}
	_nw_tcp_options_set_retransmit_connection_drop_time(options, retransmit_connection_drop_time)
}

var _nw_tcp_options_set_retransmit_fin_drop func(options Nw_protocol_options_t, retransmit_fin_drop bool)

// Nw_tcp_options_set_retransmit_fin_drop causes TCP to drop its connection after not receiving an ACK after a FIN.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_fin_drop(_:_:)
func Nw_tcp_options_set_retransmit_fin_drop(options Nw_protocol_options_t, retransmit_fin_drop bool) {
	if _nw_tcp_options_set_retransmit_fin_drop == nil {
		panic("Network: symbol nw_tcp_options_set_retransmit_fin_drop not loaded")
	}
	_nw_tcp_options_set_retransmit_fin_drop(options, retransmit_fin_drop)
}

var _nw_tls_copy_sec_protocol_metadata func(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t

// Nw_tls_copy_sec_protocol_metadata accesses the result of the TLS handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_metadata(_:)
func Nw_tls_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t {
	if _nw_tls_copy_sec_protocol_metadata == nil {
		panic("Network: symbol nw_tls_copy_sec_protocol_metadata not loaded")
	}
	return _nw_tls_copy_sec_protocol_metadata(metadata)
}

var _nw_tls_copy_sec_protocol_options func(options Nw_protocol_options_t) security.Sec_protocol_options_t

// Nw_tls_copy_sec_protocol_options accesses the handshake security options TLS will use.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_options(_:)
func Nw_tls_copy_sec_protocol_options(options Nw_protocol_options_t) security.Sec_protocol_options_t {
	if _nw_tls_copy_sec_protocol_options == nil {
		panic("Network: symbol nw_tls_copy_sec_protocol_options not loaded")
	}
	return _nw_tls_copy_sec_protocol_options(options)
}

var _nw_tls_create_options func() Nw_protocol_options_t

// Nw_tls_create_options initializes a default set of TLS connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_create_options()
func Nw_tls_create_options() Nw_protocol_options_t {
	if _nw_tls_create_options == nil {
		panic("Network: symbol nw_tls_create_options not loaded")
	}
	return _nw_tls_create_options()
}

var _nw_txt_record_access_bytes func(txt_record Nw_txt_record_t, access_bytes Nw_txt_record_access_bytes_t) bool

// Nw_txt_record_access_bytes accesses the raw bytes contained within a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes(_:_:)
func Nw_txt_record_access_bytes(txt_record Nw_txt_record_t, access_bytes Nw_txt_record_access_bytes_t) bool {
	if _nw_txt_record_access_bytes == nil {
		panic("Network: symbol nw_txt_record_access_bytes not loaded")
	}
	return _nw_txt_record_access_bytes(txt_record, access_bytes)
}

var _nw_txt_record_access_key func(txt_record Nw_txt_record_t, key string, access_value Nw_txt_record_access_key_t) bool

// Nw_txt_record_access_key accesses the value for a specific key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_key(_:_:_:)
func Nw_txt_record_access_key(txt_record Nw_txt_record_t, key string, access_value Nw_txt_record_access_key_t) bool {
	if _nw_txt_record_access_key == nil {
		panic("Network: symbol nw_txt_record_access_key not loaded")
	}
	return _nw_txt_record_access_key(txt_record, key, access_value)
}

var _nw_txt_record_apply func(txt_record Nw_txt_record_t, applier Nw_txt_record_applier_t) bool

// Nw_txt_record_apply iterates through all keys in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_apply(_:_:)
func Nw_txt_record_apply(txt_record Nw_txt_record_t, applier Nw_txt_record_applier_t) bool {
	if _nw_txt_record_apply == nil {
		panic("Network: symbol nw_txt_record_apply not loaded")
	}
	return _nw_txt_record_apply(txt_record, applier)
}

var _nw_txt_record_copy func(txt_record Nw_txt_record_t) Nw_txt_record_t

// Nw_txt_record_copy performs a deep copy of a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_copy(_:)
func Nw_txt_record_copy(txt_record Nw_txt_record_t) Nw_txt_record_t {
	if _nw_txt_record_copy == nil {
		panic("Network: symbol nw_txt_record_copy not loaded")
	}
	return _nw_txt_record_copy(txt_record)
}

var _nw_txt_record_create_dictionary func() Nw_txt_record_t

// Nw_txt_record_create_dictionary initializes a TXT record as a dictionary of strings.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_dictionary()
func Nw_txt_record_create_dictionary() Nw_txt_record_t {
	if _nw_txt_record_create_dictionary == nil {
		panic("Network: symbol nw_txt_record_create_dictionary not loaded")
	}
	return _nw_txt_record_create_dictionary()
}

var _nw_txt_record_create_with_bytes func(txt_bytes *uint8, txt_len uintptr) Nw_txt_record_t

// Nw_txt_record_create_with_bytes initializes a TXT record with raw bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_with_bytes(_:_:)
func Nw_txt_record_create_with_bytes(txt_bytes *uint8, txt_len uintptr) Nw_txt_record_t {
	if _nw_txt_record_create_with_bytes == nil {
		panic("Network: symbol nw_txt_record_create_with_bytes not loaded")
	}
	return _nw_txt_record_create_with_bytes(txt_bytes, txt_len)
}

var _nw_txt_record_find_key func(txt_record Nw_txt_record_t, key string) unsafe.Pointer

// Nw_txt_record_find_key checks the status of value associated with a key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_find_key(_:_:)
func Nw_txt_record_find_key(txt_record Nw_txt_record_t, key string) unsafe.Pointer {
	if _nw_txt_record_find_key == nil {
		panic("Network: symbol nw_txt_record_find_key not loaded")
	}
	return _nw_txt_record_find_key(txt_record, key)
}

var _nw_txt_record_get_key_count func(txt_record Nw_txt_record_t) uintptr

// Nw_txt_record_get_key_count accesses the number of keys stored in the TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_get_key_count(_:)
func Nw_txt_record_get_key_count(txt_record Nw_txt_record_t) uintptr {
	if _nw_txt_record_get_key_count == nil {
		panic("Network: symbol nw_txt_record_get_key_count not loaded")
	}
	return _nw_txt_record_get_key_count(txt_record)
}

var _nw_txt_record_is_dictionary func(txt_record Nw_txt_record_t) bool

// Nw_txt_record_is_dictionary checks whether a TXT record conforms to a dictionary format.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_dictionary(_:)
func Nw_txt_record_is_dictionary(txt_record Nw_txt_record_t) bool {
	if _nw_txt_record_is_dictionary == nil {
		panic("Network: symbol nw_txt_record_is_dictionary not loaded")
	}
	return _nw_txt_record_is_dictionary(txt_record)
}

var _nw_txt_record_is_equal func(left Nw_txt_record_t, right Nw_txt_record_t) bool

// Nw_txt_record_is_equal checks whether two TXT records are equivalent.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_equal(_:_:)
func Nw_txt_record_is_equal(left Nw_txt_record_t, right Nw_txt_record_t) bool {
	if _nw_txt_record_is_equal == nil {
		panic("Network: symbol nw_txt_record_is_equal not loaded")
	}
	return _nw_txt_record_is_equal(left, right)
}

var _nw_txt_record_remove_key func(txt_record Nw_txt_record_t, key string) bool

// Nw_txt_record_remove_key removes a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_remove_key(_:_:)
func Nw_txt_record_remove_key(txt_record Nw_txt_record_t, key string) bool {
	if _nw_txt_record_remove_key == nil {
		panic("Network: symbol nw_txt_record_remove_key not loaded")
	}
	return _nw_txt_record_remove_key(txt_record, key)
}

var _nw_txt_record_set_key func(txt_record Nw_txt_record_t, key string, value *uint8, value_len uintptr) bool

// Nw_txt_record_set_key sets a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_set_key(_:_:_:_:)
func Nw_txt_record_set_key(txt_record Nw_txt_record_t, key string, value *uint8, value_len uintptr) bool {
	if _nw_txt_record_set_key == nil {
		panic("Network: symbol nw_txt_record_set_key not loaded")
	}
	return _nw_txt_record_set_key(txt_record, key, value, value_len)
}

var _nw_udp_create_metadata func() Nw_protocol_metadata_t

// Nw_udp_create_metadata initializes a default UDP message.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_metadata()
func Nw_udp_create_metadata() Nw_protocol_metadata_t {
	if _nw_udp_create_metadata == nil {
		panic("Network: symbol nw_udp_create_metadata not loaded")
	}
	return _nw_udp_create_metadata()
}

var _nw_udp_create_options func() Nw_protocol_options_t

// Nw_udp_create_options initializes a default set of UDP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_options()
func Nw_udp_create_options() Nw_protocol_options_t {
	if _nw_udp_create_options == nil {
		panic("Network: symbol nw_udp_create_options not loaded")
	}
	return _nw_udp_create_options()
}

var _nw_udp_options_set_prefer_no_checksum func(options Nw_protocol_options_t, prefer_no_checksum bool)

// Nw_udp_options_set_prefer_no_checksum configures the connection to not send UDP checksums.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_options_set_prefer_no_checksum(_:_:)
func Nw_udp_options_set_prefer_no_checksum(options Nw_protocol_options_t, prefer_no_checksum bool) {
	if _nw_udp_options_set_prefer_no_checksum == nil {
		panic("Network: symbol nw_udp_options_set_prefer_no_checksum not loaded")
	}
	_nw_udp_options_set_prefer_no_checksum(options, prefer_no_checksum)
}

var _nw_ws_create_metadata func(opcode unsafe.Pointer) Nw_protocol_metadata_t

// Nw_ws_create_metadata initializes a WebSocket message with a specific type code.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_metadata(_:)
func Nw_ws_create_metadata(opcode unsafe.Pointer) Nw_protocol_metadata_t {
	if _nw_ws_create_metadata == nil {
		panic("Network: symbol nw_ws_create_metadata not loaded")
	}
	return _nw_ws_create_metadata(opcode)
}

var _nw_ws_create_options func(version unsafe.Pointer) Nw_protocol_options_t

// Nw_ws_create_options initializes a default set of WebSocket connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_options(_:)
func Nw_ws_create_options(version unsafe.Pointer) Nw_protocol_options_t {
	if _nw_ws_create_options == nil {
		panic("Network: symbol nw_ws_create_options not loaded")
	}
	return _nw_ws_create_options(version)
}

var _nw_ws_metadata_copy_server_response func(metadata Nw_protocol_metadata_t) Nw_ws_response_t

// Nw_ws_metadata_copy_server_response accesses the WebSocket server’s response sent during the handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_copy_server_response(_:)
func Nw_ws_metadata_copy_server_response(metadata Nw_protocol_metadata_t) Nw_ws_response_t {
	if _nw_ws_metadata_copy_server_response == nil {
		panic("Network: symbol nw_ws_metadata_copy_server_response not loaded")
	}
	return _nw_ws_metadata_copy_server_response(metadata)
}

var _nw_ws_metadata_get_close_code func(metadata Nw_protocol_metadata_t) unsafe.Pointer

// Nw_ws_metadata_get_close_code accesses the close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_close_code(_:)
func Nw_ws_metadata_get_close_code(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	if _nw_ws_metadata_get_close_code == nil {
		panic("Network: symbol nw_ws_metadata_get_close_code not loaded")
	}
	return _nw_ws_metadata_get_close_code(metadata)
}

var _nw_ws_metadata_get_opcode func(metadata Nw_protocol_metadata_t) unsafe.Pointer

// Nw_ws_metadata_get_opcode checks the type code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_opcode(_:)
func Nw_ws_metadata_get_opcode(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	if _nw_ws_metadata_get_opcode == nil {
		panic("Network: symbol nw_ws_metadata_get_opcode not loaded")
	}
	return _nw_ws_metadata_get_opcode(metadata)
}

var _nw_ws_metadata_set_close_code func(metadata Nw_protocol_metadata_t, close_code unsafe.Pointer)

// Nw_ws_metadata_set_close_code sets a close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_close_code(_:_:)
func Nw_ws_metadata_set_close_code(metadata Nw_protocol_metadata_t, close_code unsafe.Pointer) {
	if _nw_ws_metadata_set_close_code == nil {
		panic("Network: symbol nw_ws_metadata_set_close_code not loaded")
	}
	_nw_ws_metadata_set_close_code(metadata, close_code)
}

var _nw_ws_metadata_set_pong_handler func(metadata Nw_protocol_metadata_t, client_queue uintptr, pong_handler Nw_ws_pong_handler_t)

// Nw_ws_metadata_set_pong_handler sets a handler on a Ping message to be invoked when the corresponding Pong message is received.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_pong_handler(_:_:_:)
func Nw_ws_metadata_set_pong_handler(metadata Nw_protocol_metadata_t, client_queue dispatch.Queue, pong_handler Nw_ws_pong_handler_t) {
	if _nw_ws_metadata_set_pong_handler == nil {
		panic("Network: symbol nw_ws_metadata_set_pong_handler not loaded")
	}
	_nw_ws_metadata_set_pong_handler(metadata, uintptr(client_queue.Handle()), pong_handler)
}

var _nw_ws_options_add_additional_header func(options Nw_protocol_options_t, name string, value string)

// Nw_ws_options_add_additional_header adds additional HTTP header fields to be sent by the client during the WebSocket handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_additional_header(_:_:_:)
func Nw_ws_options_add_additional_header(options Nw_protocol_options_t, name string, value string) {
	if _nw_ws_options_add_additional_header == nil {
		panic("Network: symbol nw_ws_options_add_additional_header not loaded")
	}
	_nw_ws_options_add_additional_header(options, name, value)
}

var _nw_ws_options_add_subprotocol func(options Nw_protocol_options_t, subprotocol string)

// Nw_ws_options_add_subprotocol adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_subprotocol(_:_:)
func Nw_ws_options_add_subprotocol(options Nw_protocol_options_t, subprotocol string) {
	if _nw_ws_options_add_subprotocol == nil {
		panic("Network: symbol nw_ws_options_add_subprotocol not loaded")
	}
	_nw_ws_options_add_subprotocol(options, subprotocol)
}

var _nw_ws_options_set_auto_reply_ping func(options Nw_protocol_options_t, auto_reply_ping bool)

// Nw_ws_options_set_auto_reply_ping configures the connection to automatically reply to Ping messages instead of delivering them to you.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_auto_reply_ping(_:_:)
func Nw_ws_options_set_auto_reply_ping(options Nw_protocol_options_t, auto_reply_ping bool) {
	if _nw_ws_options_set_auto_reply_ping == nil {
		panic("Network: symbol nw_ws_options_set_auto_reply_ping not loaded")
	}
	_nw_ws_options_set_auto_reply_ping(options, auto_reply_ping)
}

var _nw_ws_options_set_client_request_handler func(options Nw_protocol_options_t, client_queue uintptr, handler Nw_ws_client_request_handler_t)

// Nw_ws_options_set_client_request_handler sets a handler to react to as a server to inbound WebSocket client handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_client_request_handler(_:_:_:)
func Nw_ws_options_set_client_request_handler(options Nw_protocol_options_t, client_queue dispatch.Queue, handler Nw_ws_client_request_handler_t) {
	if _nw_ws_options_set_client_request_handler == nil {
		panic("Network: symbol nw_ws_options_set_client_request_handler not loaded")
	}
	_nw_ws_options_set_client_request_handler(options, uintptr(client_queue.Handle()), handler)
}

var _nw_ws_options_set_maximum_message_size func(options Nw_protocol_options_t, maximum_message_size uintptr)

// Nw_ws_options_set_maximum_message_size sets the maximum allowed message size, in bytes, to be received by the WebSocket connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_maximum_message_size(_:_:)
func Nw_ws_options_set_maximum_message_size(options Nw_protocol_options_t, maximum_message_size uintptr) {
	if _nw_ws_options_set_maximum_message_size == nil {
		panic("Network: symbol nw_ws_options_set_maximum_message_size not loaded")
	}
	_nw_ws_options_set_maximum_message_size(options, maximum_message_size)
}

var _nw_ws_options_set_skip_handshake func(options Nw_protocol_options_t, skip_handshake bool)

// Nw_ws_options_set_skip_handshake specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_skip_handshake(_:_:)
func Nw_ws_options_set_skip_handshake(options Nw_protocol_options_t, skip_handshake bool) {
	if _nw_ws_options_set_skip_handshake == nil {
		panic("Network: symbol nw_ws_options_set_skip_handshake not loaded")
	}
	_nw_ws_options_set_skip_handshake(options, skip_handshake)
}

var _nw_ws_request_enumerate_additional_headers func(request Nw_ws_request_t, enumerator Nw_ws_additional_header_enumerator_t) bool

// Nw_ws_request_enumerate_additional_headers enumerates additional HTTP headers in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_additional_headers(_:_:)
func Nw_ws_request_enumerate_additional_headers(request Nw_ws_request_t, enumerator Nw_ws_additional_header_enumerator_t) bool {
	if _nw_ws_request_enumerate_additional_headers == nil {
		panic("Network: symbol nw_ws_request_enumerate_additional_headers not loaded")
	}
	return _nw_ws_request_enumerate_additional_headers(request, enumerator)
}

var _nw_ws_request_enumerate_subprotocols func(request Nw_ws_request_t, enumerator Nw_ws_subprotocol_enumerator_t) bool

// Nw_ws_request_enumerate_subprotocols enumerates the supported subprotocols in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_subprotocols(_:_:)
func Nw_ws_request_enumerate_subprotocols(request Nw_ws_request_t, enumerator Nw_ws_subprotocol_enumerator_t) bool {
	if _nw_ws_request_enumerate_subprotocols == nil {
		panic("Network: symbol nw_ws_request_enumerate_subprotocols not loaded")
	}
	return _nw_ws_request_enumerate_subprotocols(request, enumerator)
}

var _nw_ws_response_add_additional_header func(response Nw_ws_response_t, name string, value string)

// Nw_ws_response_add_additional_header adds an additional HTTP header to a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_add_additional_header(_:_:_:)
func Nw_ws_response_add_additional_header(response Nw_ws_response_t, name string, value string) {
	if _nw_ws_response_add_additional_header == nil {
		panic("Network: symbol nw_ws_response_add_additional_header not loaded")
	}
	_nw_ws_response_add_additional_header(response, name, value)
}

var _nw_ws_response_create func(status unsafe.Pointer, selected_subprotocol string) Nw_ws_response_t

// Nw_ws_response_create initializes a WebSocket server response with a status and selected subprotocol.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_create(_:_:)
func Nw_ws_response_create(status unsafe.Pointer, selected_subprotocol string) Nw_ws_response_t {
	if _nw_ws_response_create == nil {
		panic("Network: symbol nw_ws_response_create not loaded")
	}
	return _nw_ws_response_create(status, selected_subprotocol)
}

var _nw_ws_response_enumerate_additional_headers func(response Nw_ws_response_t, enumerator Nw_ws_additional_header_enumerator_t) bool

// Nw_ws_response_enumerate_additional_headers enumerates the additional HTTP headers in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_enumerate_additional_headers(_:_:)
func Nw_ws_response_enumerate_additional_headers(response Nw_ws_response_t, enumerator Nw_ws_additional_header_enumerator_t) bool {
	if _nw_ws_response_enumerate_additional_headers == nil {
		panic("Network: symbol nw_ws_response_enumerate_additional_headers not loaded")
	}
	return _nw_ws_response_enumerate_additional_headers(response, enumerator)
}

var _nw_ws_response_get_selected_subprotocol func(response Nw_ws_response_t) *byte

// Nw_ws_response_get_selected_subprotocol accesses the selected subprotocol in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_selected_subprotocol(_:)
func Nw_ws_response_get_selected_subprotocol(response Nw_ws_response_t) *byte {
	if _nw_ws_response_get_selected_subprotocol == nil {
		panic("Network: symbol nw_ws_response_get_selected_subprotocol not loaded")
	}
	return _nw_ws_response_get_selected_subprotocol(response)
}

var _nw_ws_response_get_status func(response Nw_ws_response_t) unsafe.Pointer

// Nw_ws_response_get_status accesses the status of a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_status(_:)
func Nw_ws_response_get_status(response Nw_ws_response_t) unsafe.Pointer {
	if _nw_ws_response_get_status == nil {
		panic("Network: symbol nw_ws_response_get_status not loaded")
	}
	return _nw_ws_response_get_status(response)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nw_advertise_descriptor_copy_txt_record_object, frameworkHandle, "nw_advertise_descriptor_copy_txt_record_object")
	registerFunc(&_nw_advertise_descriptor_create_application_service, frameworkHandle, "nw_advertise_descriptor_create_application_service")
	registerFunc(&_nw_advertise_descriptor_create_bonjour_service, frameworkHandle, "nw_advertise_descriptor_create_bonjour_service")
	registerFunc(&_nw_advertise_descriptor_get_application_service_name, frameworkHandle, "nw_advertise_descriptor_get_application_service_name")
	registerFunc(&_nw_advertise_descriptor_get_no_auto_rename, frameworkHandle, "nw_advertise_descriptor_get_no_auto_rename")
	registerFunc(&_nw_advertise_descriptor_set_no_auto_rename, frameworkHandle, "nw_advertise_descriptor_set_no_auto_rename")
	registerFunc(&_nw_advertise_descriptor_set_txt_record, frameworkHandle, "nw_advertise_descriptor_set_txt_record")
	registerFunc(&_nw_advertise_descriptor_set_txt_record_object, frameworkHandle, "nw_advertise_descriptor_set_txt_record_object")
	registerFunc(&_nw_browse_descriptor_create_application_service, frameworkHandle, "nw_browse_descriptor_create_application_service")
	registerFunc(&_nw_browse_descriptor_create_bonjour_service, frameworkHandle, "nw_browse_descriptor_create_bonjour_service")
	registerFunc(&_nw_browse_descriptor_get_application_service_name, frameworkHandle, "nw_browse_descriptor_get_application_service_name")
	registerFunc(&_nw_browse_descriptor_get_bonjour_service_domain, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_domain")
	registerFunc(&_nw_browse_descriptor_get_bonjour_service_type, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_type")
	registerFunc(&_nw_browse_descriptor_get_include_txt_record, frameworkHandle, "nw_browse_descriptor_get_include_txt_record")
	registerFunc(&_nw_browse_descriptor_set_include_txt_record, frameworkHandle, "nw_browse_descriptor_set_include_txt_record")
	registerFunc(&_nw_browse_result_copy_endpoint, frameworkHandle, "nw_browse_result_copy_endpoint")
	registerFunc(&_nw_browse_result_copy_txt_record_object, frameworkHandle, "nw_browse_result_copy_txt_record_object")
	registerFunc(&_nw_browse_result_enumerate_interfaces, frameworkHandle, "nw_browse_result_enumerate_interfaces")
	registerFunc(&_nw_browse_result_get_changes, frameworkHandle, "nw_browse_result_get_changes")
	registerFunc(&_nw_browse_result_get_interfaces_count, frameworkHandle, "nw_browse_result_get_interfaces_count")
	registerFunc(&_nw_browser_cancel, frameworkHandle, "nw_browser_cancel")
	registerFunc(&_nw_browser_copy_browse_descriptor, frameworkHandle, "nw_browser_copy_browse_descriptor")
	registerFunc(&_nw_browser_copy_parameters, frameworkHandle, "nw_browser_copy_parameters")
	registerFunc(&_nw_browser_create, frameworkHandle, "nw_browser_create")
	registerFunc(&_nw_browser_set_browse_results_changed_handler, frameworkHandle, "nw_browser_set_browse_results_changed_handler")
	registerFunc(&_nw_browser_set_queue, frameworkHandle, "nw_browser_set_queue")
	registerFunc(&_nw_browser_set_state_changed_handler, frameworkHandle, "nw_browser_set_state_changed_handler")
	registerFunc(&_nw_browser_start, frameworkHandle, "nw_browser_start")
	registerFunc(&_nw_connection_access_establishment_report, frameworkHandle, "nw_connection_access_establishment_report")
	registerFunc(&_nw_connection_batch, frameworkHandle, "nw_connection_batch")
	registerFunc(&_nw_connection_cancel, frameworkHandle, "nw_connection_cancel")
	registerFunc(&_nw_connection_cancel_current_endpoint, frameworkHandle, "nw_connection_cancel_current_endpoint")
	registerFunc(&_nw_connection_copy_current_path, frameworkHandle, "nw_connection_copy_current_path")
	registerFunc(&_nw_connection_copy_description, frameworkHandle, "nw_connection_copy_description")
	registerFunc(&_nw_connection_copy_endpoint, frameworkHandle, "nw_connection_copy_endpoint")
	registerFunc(&_nw_connection_copy_parameters, frameworkHandle, "nw_connection_copy_parameters")
	registerFunc(&_nw_connection_copy_protocol_metadata, frameworkHandle, "nw_connection_copy_protocol_metadata")
	registerFunc(&_nw_connection_create, frameworkHandle, "nw_connection_create")
	registerFunc(&_nw_connection_create_new_data_transfer_report, frameworkHandle, "nw_connection_create_new_data_transfer_report")
	registerFunc(&_nw_connection_force_cancel, frameworkHandle, "nw_connection_force_cancel")
	registerFunc(&_nw_connection_get_maximum_datagram_size, frameworkHandle, "nw_connection_get_maximum_datagram_size")
	registerFunc(&_nw_connection_group_cancel, frameworkHandle, "nw_connection_group_cancel")
	registerFunc(&_nw_connection_group_copy_descriptor, frameworkHandle, "nw_connection_group_copy_descriptor")
	registerFunc(&_nw_connection_group_copy_local_endpoint_for_message, frameworkHandle, "nw_connection_group_copy_local_endpoint_for_message")
	registerFunc(&_nw_connection_group_copy_parameters, frameworkHandle, "nw_connection_group_copy_parameters")
	registerFunc(&_nw_connection_group_copy_path_for_message, frameworkHandle, "nw_connection_group_copy_path_for_message")
	registerFunc(&_nw_connection_group_copy_protocol_metadata, frameworkHandle, "nw_connection_group_copy_protocol_metadata")
	registerFunc(&_nw_connection_group_copy_protocol_metadata_for_message, frameworkHandle, "nw_connection_group_copy_protocol_metadata_for_message")
	registerFunc(&_nw_connection_group_copy_remote_endpoint_for_message, frameworkHandle, "nw_connection_group_copy_remote_endpoint_for_message")
	registerFunc(&_nw_connection_group_create, frameworkHandle, "nw_connection_group_create")
	registerFunc(&_nw_connection_group_extract_connection, frameworkHandle, "nw_connection_group_extract_connection")
	registerFunc(&_nw_connection_group_extract_connection_for_message, frameworkHandle, "nw_connection_group_extract_connection_for_message")
	registerFunc(&_nw_connection_group_reinsert_extracted_connection, frameworkHandle, "nw_connection_group_reinsert_extracted_connection")
	registerFunc(&_nw_connection_group_reply, frameworkHandle, "nw_connection_group_reply")
	registerFunc(&_nw_connection_group_send_message, frameworkHandle, "nw_connection_group_send_message")
	registerFunc(&_nw_connection_group_set_new_connection_handler, frameworkHandle, "nw_connection_group_set_new_connection_handler")
	registerFunc(&_nw_connection_group_set_queue, frameworkHandle, "nw_connection_group_set_queue")
	registerFunc(&_nw_connection_group_set_receive_handler, frameworkHandle, "nw_connection_group_set_receive_handler")
	registerFunc(&_nw_connection_group_set_state_changed_handler, frameworkHandle, "nw_connection_group_set_state_changed_handler")
	registerFunc(&_nw_connection_group_start, frameworkHandle, "nw_connection_group_start")
	registerFunc(&_nw_connection_receive, frameworkHandle, "nw_connection_receive")
	registerFunc(&_nw_connection_receive_message, frameworkHandle, "nw_connection_receive_message")
	registerFunc(&_nw_connection_restart, frameworkHandle, "nw_connection_restart")
	registerFunc(&_nw_connection_send, frameworkHandle, "nw_connection_send")
	registerFunc(&_nw_connection_set_better_path_available_handler, frameworkHandle, "nw_connection_set_better_path_available_handler")
	registerFunc(&_nw_connection_set_path_changed_handler, frameworkHandle, "nw_connection_set_path_changed_handler")
	registerFunc(&_nw_connection_set_queue, frameworkHandle, "nw_connection_set_queue")
	registerFunc(&_nw_connection_set_state_changed_handler, frameworkHandle, "nw_connection_set_state_changed_handler")
	registerFunc(&_nw_connection_set_viability_changed_handler, frameworkHandle, "nw_connection_set_viability_changed_handler")
	registerFunc(&_nw_connection_start, frameworkHandle, "nw_connection_start")
	registerFunc(&_nw_content_context_copy_antecedent, frameworkHandle, "nw_content_context_copy_antecedent")
	registerFunc(&_nw_content_context_copy_protocol_metadata, frameworkHandle, "nw_content_context_copy_protocol_metadata")
	registerFunc(&_nw_content_context_create, frameworkHandle, "nw_content_context_create")
	registerFunc(&_nw_content_context_foreach_protocol_metadata, frameworkHandle, "nw_content_context_foreach_protocol_metadata")
	registerFunc(&_nw_content_context_get_expiration_milliseconds, frameworkHandle, "nw_content_context_get_expiration_milliseconds")
	registerFunc(&_nw_content_context_get_identifier, frameworkHandle, "nw_content_context_get_identifier")
	registerFunc(&_nw_content_context_get_is_final, frameworkHandle, "nw_content_context_get_is_final")
	registerFunc(&_nw_content_context_get_relative_priority, frameworkHandle, "nw_content_context_get_relative_priority")
	registerFunc(&_nw_content_context_set_antecedent, frameworkHandle, "nw_content_context_set_antecedent")
	registerFunc(&_nw_content_context_set_expiration_milliseconds, frameworkHandle, "nw_content_context_set_expiration_milliseconds")
	registerFunc(&_nw_content_context_set_is_final, frameworkHandle, "nw_content_context_set_is_final")
	registerFunc(&_nw_content_context_set_metadata_for_protocol, frameworkHandle, "nw_content_context_set_metadata_for_protocol")
	registerFunc(&_nw_content_context_set_relative_priority, frameworkHandle, "nw_content_context_set_relative_priority")
	registerFunc(&_nw_data_transfer_report_collect, frameworkHandle, "nw_data_transfer_report_collect")
	registerFunc(&_nw_data_transfer_report_copy_path_interface, frameworkHandle, "nw_data_transfer_report_copy_path_interface")
	registerFunc(&_nw_data_transfer_report_get_duration_milliseconds, frameworkHandle, "nw_data_transfer_report_get_duration_milliseconds")
	registerFunc(&_nw_data_transfer_report_get_path_count, frameworkHandle, "nw_data_transfer_report_get_path_count")
	registerFunc(&_nw_data_transfer_report_get_path_radio_type, frameworkHandle, "nw_data_transfer_report_get_path_radio_type")
	registerFunc(&_nw_data_transfer_report_get_received_application_byte_count, frameworkHandle, "nw_data_transfer_report_get_received_application_byte_count")
	registerFunc(&_nw_data_transfer_report_get_received_ip_packet_count, frameworkHandle, "nw_data_transfer_report_get_received_ip_packet_count")
	registerFunc(&_nw_data_transfer_report_get_received_transport_byte_count, frameworkHandle, "nw_data_transfer_report_get_received_transport_byte_count")
	registerFunc(&_nw_data_transfer_report_get_received_transport_duplicate_byte_count, frameworkHandle, "nw_data_transfer_report_get_received_transport_duplicate_byte_count")
	registerFunc(&_nw_data_transfer_report_get_received_transport_out_of_order_byte_count, frameworkHandle, "nw_data_transfer_report_get_received_transport_out_of_order_byte_count")
	registerFunc(&_nw_data_transfer_report_get_sent_application_byte_count, frameworkHandle, "nw_data_transfer_report_get_sent_application_byte_count")
	registerFunc(&_nw_data_transfer_report_get_sent_ip_packet_count, frameworkHandle, "nw_data_transfer_report_get_sent_ip_packet_count")
	registerFunc(&_nw_data_transfer_report_get_sent_transport_byte_count, frameworkHandle, "nw_data_transfer_report_get_sent_transport_byte_count")
	registerFunc(&_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count, frameworkHandle, "nw_data_transfer_report_get_sent_transport_retransmitted_byte_count")
	registerFunc(&_nw_data_transfer_report_get_state, frameworkHandle, "nw_data_transfer_report_get_state")
	registerFunc(&_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds, frameworkHandle, "nw_data_transfer_report_get_transport_minimum_rtt_milliseconds")
	registerFunc(&_nw_data_transfer_report_get_transport_rtt_variance, frameworkHandle, "nw_data_transfer_report_get_transport_rtt_variance")
	registerFunc(&_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds, frameworkHandle, "nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds")
	registerFunc(&_nw_endpoint_copy_address_string, frameworkHandle, "nw_endpoint_copy_address_string")
	registerFunc(&_nw_endpoint_copy_port_string, frameworkHandle, "nw_endpoint_copy_port_string")
	registerFunc(&_nw_endpoint_copy_txt_record, frameworkHandle, "nw_endpoint_copy_txt_record")
	registerFunc(&_nw_endpoint_create_address, frameworkHandle, "nw_endpoint_create_address")
	registerFunc(&_nw_endpoint_create_bonjour_service, frameworkHandle, "nw_endpoint_create_bonjour_service")
	registerFunc(&_nw_endpoint_create_host, frameworkHandle, "nw_endpoint_create_host")
	registerFunc(&_nw_endpoint_create_url, frameworkHandle, "nw_endpoint_create_url")
	registerFunc(&_nw_endpoint_get_address, frameworkHandle, "nw_endpoint_get_address")
	registerFunc(&_nw_endpoint_get_bonjour_service_domain, frameworkHandle, "nw_endpoint_get_bonjour_service_domain")
	registerFunc(&_nw_endpoint_get_bonjour_service_name, frameworkHandle, "nw_endpoint_get_bonjour_service_name")
	registerFunc(&_nw_endpoint_get_bonjour_service_type, frameworkHandle, "nw_endpoint_get_bonjour_service_type")
	registerFunc(&_nw_endpoint_get_hostname, frameworkHandle, "nw_endpoint_get_hostname")
	registerFunc(&_nw_endpoint_get_port, frameworkHandle, "nw_endpoint_get_port")
	registerFunc(&_nw_endpoint_get_signature, frameworkHandle, "nw_endpoint_get_signature")
	registerFunc(&_nw_endpoint_get_type, frameworkHandle, "nw_endpoint_get_type")
	registerFunc(&_nw_endpoint_get_url, frameworkHandle, "nw_endpoint_get_url")
	registerFunc(&_nw_error_copy_cf_error, frameworkHandle, "nw_error_copy_cf_error")
	registerFunc(&_nw_error_get_error_code, frameworkHandle, "nw_error_get_error_code")
	registerFunc(&_nw_error_get_error_domain, frameworkHandle, "nw_error_get_error_domain")
	registerFunc(&_nw_establishment_report_copy_proxy_endpoint, frameworkHandle, "nw_establishment_report_copy_proxy_endpoint")
	registerFunc(&_nw_establishment_report_enumerate_protocols, frameworkHandle, "nw_establishment_report_enumerate_protocols")
	registerFunc(&_nw_establishment_report_enumerate_resolution_reports, frameworkHandle, "nw_establishment_report_enumerate_resolution_reports")
	registerFunc(&_nw_establishment_report_enumerate_resolutions, frameworkHandle, "nw_establishment_report_enumerate_resolutions")
	registerFunc(&_nw_establishment_report_get_attempt_started_after_milliseconds, frameworkHandle, "nw_establishment_report_get_attempt_started_after_milliseconds")
	registerFunc(&_nw_establishment_report_get_duration_milliseconds, frameworkHandle, "nw_establishment_report_get_duration_milliseconds")
	registerFunc(&_nw_establishment_report_get_previous_attempt_count, frameworkHandle, "nw_establishment_report_get_previous_attempt_count")
	registerFunc(&_nw_establishment_report_get_proxy_configured, frameworkHandle, "nw_establishment_report_get_proxy_configured")
	registerFunc(&_nw_establishment_report_get_used_proxy, frameworkHandle, "nw_establishment_report_get_used_proxy")
	registerFunc(&_nw_ethernet_channel_cancel, frameworkHandle, "nw_ethernet_channel_cancel")
	registerFunc(&_nw_ethernet_channel_create, frameworkHandle, "nw_ethernet_channel_create")
	registerFunc(&_nw_ethernet_channel_create_with_parameters, frameworkHandle, "nw_ethernet_channel_create_with_parameters")
	registerFunc(&_nw_ethernet_channel_get_maximum_payload_size, frameworkHandle, "nw_ethernet_channel_get_maximum_payload_size")
	registerFunc(&_nw_ethernet_channel_send, frameworkHandle, "nw_ethernet_channel_send")
	registerFunc(&_nw_ethernet_channel_set_queue, frameworkHandle, "nw_ethernet_channel_set_queue")
	registerFunc(&_nw_ethernet_channel_set_receive_handler, frameworkHandle, "nw_ethernet_channel_set_receive_handler")
	registerFunc(&_nw_ethernet_channel_set_state_changed_handler, frameworkHandle, "nw_ethernet_channel_set_state_changed_handler")
	registerFunc(&_nw_ethernet_channel_start, frameworkHandle, "nw_ethernet_channel_start")
	registerFunc(&_nw_framer_async, frameworkHandle, "nw_framer_async")
	registerFunc(&_nw_framer_copy_local_endpoint, frameworkHandle, "nw_framer_copy_local_endpoint")
	registerFunc(&_nw_framer_copy_options, frameworkHandle, "nw_framer_copy_options")
	registerFunc(&_nw_framer_copy_parameters, frameworkHandle, "nw_framer_copy_parameters")
	registerFunc(&_nw_framer_copy_remote_endpoint, frameworkHandle, "nw_framer_copy_remote_endpoint")
	registerFunc(&_nw_framer_create_definition, frameworkHandle, "nw_framer_create_definition")
	registerFunc(&_nw_framer_create_options, frameworkHandle, "nw_framer_create_options")
	registerFunc(&_nw_framer_deliver_input, frameworkHandle, "nw_framer_deliver_input")
	registerFunc(&_nw_framer_deliver_input_no_copy, frameworkHandle, "nw_framer_deliver_input_no_copy")
	registerFunc(&_nw_framer_mark_failed_with_error, frameworkHandle, "nw_framer_mark_failed_with_error")
	registerFunc(&_nw_framer_mark_ready, frameworkHandle, "nw_framer_mark_ready")
	registerFunc(&_nw_framer_message_access_value, frameworkHandle, "nw_framer_message_access_value")
	registerFunc(&_nw_framer_message_copy_object_value, frameworkHandle, "nw_framer_message_copy_object_value")
	registerFunc(&_nw_framer_message_create, frameworkHandle, "nw_framer_message_create")
	registerFunc(&_nw_framer_message_set_object_value, frameworkHandle, "nw_framer_message_set_object_value")
	registerFunc(&_nw_framer_message_set_value, frameworkHandle, "nw_framer_message_set_value")
	registerFunc(&_nw_framer_options_copy_object_value, frameworkHandle, "nw_framer_options_copy_object_value")
	registerFunc(&_nw_framer_options_set_object_value, frameworkHandle, "nw_framer_options_set_object_value")
	registerFunc(&_nw_framer_parse_input, frameworkHandle, "nw_framer_parse_input")
	registerFunc(&_nw_framer_parse_output, frameworkHandle, "nw_framer_parse_output")
	registerFunc(&_nw_framer_pass_through_input, frameworkHandle, "nw_framer_pass_through_input")
	registerFunc(&_nw_framer_pass_through_output, frameworkHandle, "nw_framer_pass_through_output")
	registerFunc(&_nw_framer_prepend_application_protocol, frameworkHandle, "nw_framer_prepend_application_protocol")
	registerFunc(&_nw_framer_protocol_create_message, frameworkHandle, "nw_framer_protocol_create_message")
	registerFunc(&_nw_framer_schedule_wakeup, frameworkHandle, "nw_framer_schedule_wakeup")
	registerFunc(&_nw_framer_set_cleanup_handler, frameworkHandle, "nw_framer_set_cleanup_handler")
	registerFunc(&_nw_framer_set_input_handler, frameworkHandle, "nw_framer_set_input_handler")
	registerFunc(&_nw_framer_set_output_handler, frameworkHandle, "nw_framer_set_output_handler")
	registerFunc(&_nw_framer_set_stop_handler, frameworkHandle, "nw_framer_set_stop_handler")
	registerFunc(&_nw_framer_set_wakeup_handler, frameworkHandle, "nw_framer_set_wakeup_handler")
	registerFunc(&_nw_framer_write_output, frameworkHandle, "nw_framer_write_output")
	registerFunc(&_nw_framer_write_output_data, frameworkHandle, "nw_framer_write_output_data")
	registerFunc(&_nw_framer_write_output_no_copy, frameworkHandle, "nw_framer_write_output_no_copy")
	registerFunc(&_nw_group_descriptor_add_endpoint, frameworkHandle, "nw_group_descriptor_add_endpoint")
	registerFunc(&_nw_group_descriptor_create_multicast, frameworkHandle, "nw_group_descriptor_create_multicast")
	registerFunc(&_nw_group_descriptor_create_multiplex, frameworkHandle, "nw_group_descriptor_create_multiplex")
	registerFunc(&_nw_group_descriptor_enumerate_endpoints, frameworkHandle, "nw_group_descriptor_enumerate_endpoints")
	registerFunc(&_nw_interface_get_index, frameworkHandle, "nw_interface_get_index")
	registerFunc(&_nw_interface_get_name, frameworkHandle, "nw_interface_get_name")
	registerFunc(&_nw_interface_get_type, frameworkHandle, "nw_interface_get_type")
	registerFunc(&_nw_ip_create_metadata, frameworkHandle, "nw_ip_create_metadata")
	registerFunc(&_nw_ip_metadata_get_ecn_flag, frameworkHandle, "nw_ip_metadata_get_ecn_flag")
	registerFunc(&_nw_ip_metadata_get_receive_time, frameworkHandle, "nw_ip_metadata_get_receive_time")
	registerFunc(&_nw_ip_metadata_get_service_class, frameworkHandle, "nw_ip_metadata_get_service_class")
	registerFunc(&_nw_ip_metadata_set_ecn_flag, frameworkHandle, "nw_ip_metadata_set_ecn_flag")
	registerFunc(&_nw_ip_metadata_set_service_class, frameworkHandle, "nw_ip_metadata_set_service_class")
	registerFunc(&_nw_ip_options_set_calculate_receive_time, frameworkHandle, "nw_ip_options_set_calculate_receive_time")
	registerFunc(&_nw_ip_options_set_disable_fragmentation, frameworkHandle, "nw_ip_options_set_disable_fragmentation")
	registerFunc(&_nw_ip_options_set_disable_multicast_loopback, frameworkHandle, "nw_ip_options_set_disable_multicast_loopback")
	registerFunc(&_nw_ip_options_set_hop_limit, frameworkHandle, "nw_ip_options_set_hop_limit")
	registerFunc(&_nw_ip_options_set_local_address_preference, frameworkHandle, "nw_ip_options_set_local_address_preference")
	registerFunc(&_nw_ip_options_set_use_minimum_mtu, frameworkHandle, "nw_ip_options_set_use_minimum_mtu")
	registerFunc(&_nw_ip_options_set_version, frameworkHandle, "nw_ip_options_set_version")
	registerFunc(&_nw_listener_cancel, frameworkHandle, "nw_listener_cancel")
	registerFunc(&_nw_listener_create, frameworkHandle, "nw_listener_create")
	registerFunc(&_nw_listener_create_with_connection, frameworkHandle, "nw_listener_create_with_connection")
	registerFunc(&_nw_listener_create_with_launchd_key, frameworkHandle, "nw_listener_create_with_launchd_key")
	registerFunc(&_nw_listener_create_with_port, frameworkHandle, "nw_listener_create_with_port")
	registerFunc(&_nw_listener_get_new_connection_limit, frameworkHandle, "nw_listener_get_new_connection_limit")
	registerFunc(&_nw_listener_get_port, frameworkHandle, "nw_listener_get_port")
	registerFunc(&_nw_listener_set_advertise_descriptor, frameworkHandle, "nw_listener_set_advertise_descriptor")
	registerFunc(&_nw_listener_set_advertised_endpoint_changed_handler, frameworkHandle, "nw_listener_set_advertised_endpoint_changed_handler")
	registerFunc(&_nw_listener_set_new_connection_group_handler, frameworkHandle, "nw_listener_set_new_connection_group_handler")
	registerFunc(&_nw_listener_set_new_connection_handler, frameworkHandle, "nw_listener_set_new_connection_handler")
	registerFunc(&_nw_listener_set_new_connection_limit, frameworkHandle, "nw_listener_set_new_connection_limit")
	registerFunc(&_nw_listener_set_queue, frameworkHandle, "nw_listener_set_queue")
	registerFunc(&_nw_listener_set_state_changed_handler, frameworkHandle, "nw_listener_set_state_changed_handler")
	registerFunc(&_nw_listener_start, frameworkHandle, "nw_listener_start")
	registerFunc(&_nw_multicast_group_descriptor_get_disable_unicast_traffic, frameworkHandle, "nw_multicast_group_descriptor_get_disable_unicast_traffic")
	registerFunc(&_nw_multicast_group_descriptor_set_disable_unicast_traffic, frameworkHandle, "nw_multicast_group_descriptor_set_disable_unicast_traffic")
	registerFunc(&_nw_multicast_group_descriptor_set_specific_source, frameworkHandle, "nw_multicast_group_descriptor_set_specific_source")
	registerFunc(&_nw_parameters_clear_prohibited_interface_types, frameworkHandle, "nw_parameters_clear_prohibited_interface_types")
	registerFunc(&_nw_parameters_clear_prohibited_interfaces, frameworkHandle, "nw_parameters_clear_prohibited_interfaces")
	registerFunc(&_nw_parameters_copy, frameworkHandle, "nw_parameters_copy")
	registerFunc(&_nw_parameters_copy_default_protocol_stack, frameworkHandle, "nw_parameters_copy_default_protocol_stack")
	registerFunc(&_nw_parameters_copy_local_endpoint, frameworkHandle, "nw_parameters_copy_local_endpoint")
	registerFunc(&_nw_parameters_copy_required_interface, frameworkHandle, "nw_parameters_copy_required_interface")
	registerFunc(&_nw_parameters_create, frameworkHandle, "nw_parameters_create")
	registerFunc(&_nw_parameters_create_application_service, frameworkHandle, "nw_parameters_create_application_service")
	registerFunc(&_nw_parameters_create_custom_ip, frameworkHandle, "nw_parameters_create_custom_ip")
	registerFunc(&_nw_parameters_create_quic, frameworkHandle, "nw_parameters_create_quic")
	registerFunc(&_nw_parameters_create_secure_tcp, frameworkHandle, "nw_parameters_create_secure_tcp")
	registerFunc(&_nw_parameters_create_secure_udp, frameworkHandle, "nw_parameters_create_secure_udp")
	registerFunc(&_nw_parameters_get_allow_ultra_constrained, frameworkHandle, "nw_parameters_get_allow_ultra_constrained")
	registerFunc(&_nw_parameters_get_expired_dns_behavior, frameworkHandle, "nw_parameters_get_expired_dns_behavior")
	registerFunc(&_nw_parameters_get_fast_open_enabled, frameworkHandle, "nw_parameters_get_fast_open_enabled")
	registerFunc(&_nw_parameters_get_include_peer_to_peer, frameworkHandle, "nw_parameters_get_include_peer_to_peer")
	registerFunc(&_nw_parameters_get_local_only, frameworkHandle, "nw_parameters_get_local_only")
	registerFunc(&_nw_parameters_get_multipath_service, frameworkHandle, "nw_parameters_get_multipath_service")
	registerFunc(&_nw_parameters_get_prefer_no_proxy, frameworkHandle, "nw_parameters_get_prefer_no_proxy")
	registerFunc(&_nw_parameters_get_prohibit_constrained, frameworkHandle, "nw_parameters_get_prohibit_constrained")
	registerFunc(&_nw_parameters_get_prohibit_expensive, frameworkHandle, "nw_parameters_get_prohibit_expensive")
	registerFunc(&_nw_parameters_get_required_interface_type, frameworkHandle, "nw_parameters_get_required_interface_type")
	registerFunc(&_nw_parameters_get_reuse_local_address, frameworkHandle, "nw_parameters_get_reuse_local_address")
	registerFunc(&_nw_parameters_get_service_class, frameworkHandle, "nw_parameters_get_service_class")
	registerFunc(&_nw_parameters_iterate_prohibited_interface_types, frameworkHandle, "nw_parameters_iterate_prohibited_interface_types")
	registerFunc(&_nw_parameters_iterate_prohibited_interfaces, frameworkHandle, "nw_parameters_iterate_prohibited_interfaces")
	registerFunc(&_nw_parameters_prohibit_interface, frameworkHandle, "nw_parameters_prohibit_interface")
	registerFunc(&_nw_parameters_prohibit_interface_type, frameworkHandle, "nw_parameters_prohibit_interface_type")
	registerFunc(&_nw_parameters_require_interface, frameworkHandle, "nw_parameters_require_interface")
	registerFunc(&_nw_parameters_requires_dnssec_validation, frameworkHandle, "nw_parameters_requires_dnssec_validation")
	registerFunc(&_nw_parameters_set_allow_ultra_constrained, frameworkHandle, "nw_parameters_set_allow_ultra_constrained")
	registerFunc(&_nw_parameters_set_expired_dns_behavior, frameworkHandle, "nw_parameters_set_expired_dns_behavior")
	registerFunc(&_nw_parameters_set_fast_open_enabled, frameworkHandle, "nw_parameters_set_fast_open_enabled")
	registerFunc(&_nw_parameters_set_include_peer_to_peer, frameworkHandle, "nw_parameters_set_include_peer_to_peer")
	registerFunc(&_nw_parameters_set_local_endpoint, frameworkHandle, "nw_parameters_set_local_endpoint")
	registerFunc(&_nw_parameters_set_local_only, frameworkHandle, "nw_parameters_set_local_only")
	registerFunc(&_nw_parameters_set_multipath_service, frameworkHandle, "nw_parameters_set_multipath_service")
	registerFunc(&_nw_parameters_set_prefer_no_proxy, frameworkHandle, "nw_parameters_set_prefer_no_proxy")
	registerFunc(&_nw_parameters_set_privacy_context, frameworkHandle, "nw_parameters_set_privacy_context")
	registerFunc(&_nw_parameters_set_prohibit_constrained, frameworkHandle, "nw_parameters_set_prohibit_constrained")
	registerFunc(&_nw_parameters_set_prohibit_expensive, frameworkHandle, "nw_parameters_set_prohibit_expensive")
	registerFunc(&_nw_parameters_set_required_interface_type, frameworkHandle, "nw_parameters_set_required_interface_type")
	registerFunc(&_nw_parameters_set_requires_dnssec_validation, frameworkHandle, "nw_parameters_set_requires_dnssec_validation")
	registerFunc(&_nw_parameters_set_reuse_local_address, frameworkHandle, "nw_parameters_set_reuse_local_address")
	registerFunc(&_nw_parameters_set_service_class, frameworkHandle, "nw_parameters_set_service_class")
	registerFunc(&_nw_path_copy_effective_local_endpoint, frameworkHandle, "nw_path_copy_effective_local_endpoint")
	registerFunc(&_nw_path_copy_effective_remote_endpoint, frameworkHandle, "nw_path_copy_effective_remote_endpoint")
	registerFunc(&_nw_path_enumerate_gateways, frameworkHandle, "nw_path_enumerate_gateways")
	registerFunc(&_nw_path_enumerate_interfaces, frameworkHandle, "nw_path_enumerate_interfaces")
	registerFunc(&_nw_path_get_link_quality, frameworkHandle, "nw_path_get_link_quality")
	registerFunc(&_nw_path_get_status, frameworkHandle, "nw_path_get_status")
	registerFunc(&_nw_path_get_unsatisfied_reason, frameworkHandle, "nw_path_get_unsatisfied_reason")
	registerFunc(&_nw_path_has_dns, frameworkHandle, "nw_path_has_dns")
	registerFunc(&_nw_path_has_ipv4, frameworkHandle, "nw_path_has_ipv4")
	registerFunc(&_nw_path_has_ipv6, frameworkHandle, "nw_path_has_ipv6")
	registerFunc(&_nw_path_is_constrained, frameworkHandle, "nw_path_is_constrained")
	registerFunc(&_nw_path_is_equal, frameworkHandle, "nw_path_is_equal")
	registerFunc(&_nw_path_is_expensive, frameworkHandle, "nw_path_is_expensive")
	registerFunc(&_nw_path_is_ultra_constrained, frameworkHandle, "nw_path_is_ultra_constrained")
	registerFunc(&_nw_path_monitor_cancel, frameworkHandle, "nw_path_monitor_cancel")
	registerFunc(&_nw_path_monitor_create, frameworkHandle, "nw_path_monitor_create")
	registerFunc(&_nw_path_monitor_create_for_ethernet_channel, frameworkHandle, "nw_path_monitor_create_for_ethernet_channel")
	registerFunc(&_nw_path_monitor_create_with_type, frameworkHandle, "nw_path_monitor_create_with_type")
	registerFunc(&_nw_path_monitor_prohibit_interface_type, frameworkHandle, "nw_path_monitor_prohibit_interface_type")
	registerFunc(&_nw_path_monitor_set_cancel_handler, frameworkHandle, "nw_path_monitor_set_cancel_handler")
	registerFunc(&_nw_path_monitor_set_queue, frameworkHandle, "nw_path_monitor_set_queue")
	registerFunc(&_nw_path_monitor_set_update_handler, frameworkHandle, "nw_path_monitor_set_update_handler")
	registerFunc(&_nw_path_monitor_start, frameworkHandle, "nw_path_monitor_start")
	registerFunc(&_nw_path_uses_interface_type, frameworkHandle, "nw_path_uses_interface_type")
	registerFunc(&_nw_privacy_context_add_proxy, frameworkHandle, "nw_privacy_context_add_proxy")
	registerFunc(&_nw_privacy_context_clear_proxies, frameworkHandle, "nw_privacy_context_clear_proxies")
	registerFunc(&_nw_privacy_context_create, frameworkHandle, "nw_privacy_context_create")
	registerFunc(&_nw_privacy_context_disable_logging, frameworkHandle, "nw_privacy_context_disable_logging")
	registerFunc(&_nw_privacy_context_flush_cache, frameworkHandle, "nw_privacy_context_flush_cache")
	registerFunc(&_nw_privacy_context_require_encrypted_name_resolution, frameworkHandle, "nw_privacy_context_require_encrypted_name_resolution")
	registerFunc(&_nw_protocol_copy_ip_definition, frameworkHandle, "nw_protocol_copy_ip_definition")
	registerFunc(&_nw_protocol_copy_quic_definition, frameworkHandle, "nw_protocol_copy_quic_definition")
	registerFunc(&_nw_protocol_copy_tcp_definition, frameworkHandle, "nw_protocol_copy_tcp_definition")
	registerFunc(&_nw_protocol_copy_tls_definition, frameworkHandle, "nw_protocol_copy_tls_definition")
	registerFunc(&_nw_protocol_copy_udp_definition, frameworkHandle, "nw_protocol_copy_udp_definition")
	registerFunc(&_nw_protocol_copy_ws_definition, frameworkHandle, "nw_protocol_copy_ws_definition")
	registerFunc(&_nw_protocol_definition_is_equal, frameworkHandle, "nw_protocol_definition_is_equal")
	registerFunc(&_nw_protocol_metadata_copy_definition, frameworkHandle, "nw_protocol_metadata_copy_definition")
	registerFunc(&_nw_protocol_metadata_is_framer_message, frameworkHandle, "nw_protocol_metadata_is_framer_message")
	registerFunc(&_nw_protocol_metadata_is_ip, frameworkHandle, "nw_protocol_metadata_is_ip")
	registerFunc(&_nw_protocol_metadata_is_quic, frameworkHandle, "nw_protocol_metadata_is_quic")
	registerFunc(&_nw_protocol_metadata_is_tcp, frameworkHandle, "nw_protocol_metadata_is_tcp")
	registerFunc(&_nw_protocol_metadata_is_tls, frameworkHandle, "nw_protocol_metadata_is_tls")
	registerFunc(&_nw_protocol_metadata_is_udp, frameworkHandle, "nw_protocol_metadata_is_udp")
	registerFunc(&_nw_protocol_metadata_is_ws, frameworkHandle, "nw_protocol_metadata_is_ws")
	registerFunc(&_nw_protocol_options_copy_definition, frameworkHandle, "nw_protocol_options_copy_definition")
	registerFunc(&_nw_protocol_options_is_quic, frameworkHandle, "nw_protocol_options_is_quic")
	registerFunc(&_nw_protocol_stack_clear_application_protocols, frameworkHandle, "nw_protocol_stack_clear_application_protocols")
	registerFunc(&_nw_protocol_stack_copy_internet_protocol, frameworkHandle, "nw_protocol_stack_copy_internet_protocol")
	registerFunc(&_nw_protocol_stack_copy_transport_protocol, frameworkHandle, "nw_protocol_stack_copy_transport_protocol")
	registerFunc(&_nw_protocol_stack_iterate_application_protocols, frameworkHandle, "nw_protocol_stack_iterate_application_protocols")
	registerFunc(&_nw_protocol_stack_prepend_application_protocol, frameworkHandle, "nw_protocol_stack_prepend_application_protocol")
	registerFunc(&_nw_protocol_stack_set_transport_protocol, frameworkHandle, "nw_protocol_stack_set_transport_protocol")
	registerFunc(&_nw_proxy_config_add_excluded_domain, frameworkHandle, "nw_proxy_config_add_excluded_domain")
	registerFunc(&_nw_proxy_config_add_match_domain, frameworkHandle, "nw_proxy_config_add_match_domain")
	registerFunc(&_nw_proxy_config_clear_excluded_domains, frameworkHandle, "nw_proxy_config_clear_excluded_domains")
	registerFunc(&_nw_proxy_config_clear_match_domains, frameworkHandle, "nw_proxy_config_clear_match_domains")
	registerFunc(&_nw_proxy_config_create_http_connect, frameworkHandle, "nw_proxy_config_create_http_connect")
	registerFunc(&_nw_proxy_config_create_oblivious_http, frameworkHandle, "nw_proxy_config_create_oblivious_http")
	registerFunc(&_nw_proxy_config_create_relay, frameworkHandle, "nw_proxy_config_create_relay")
	registerFunc(&_nw_proxy_config_create_socksv5, frameworkHandle, "nw_proxy_config_create_socksv5")
	registerFunc(&_nw_proxy_config_enumerate_excluded_domains, frameworkHandle, "nw_proxy_config_enumerate_excluded_domains")
	registerFunc(&_nw_proxy_config_enumerate_match_domains, frameworkHandle, "nw_proxy_config_enumerate_match_domains")
	registerFunc(&_nw_proxy_config_get_failover_allowed, frameworkHandle, "nw_proxy_config_get_failover_allowed")
	registerFunc(&_nw_proxy_config_set_failover_allowed, frameworkHandle, "nw_proxy_config_set_failover_allowed")
	registerFunc(&_nw_proxy_config_set_username_and_password, frameworkHandle, "nw_proxy_config_set_username_and_password")
	registerFunc(&_nw_quic_add_tls_application_protocol, frameworkHandle, "nw_quic_add_tls_application_protocol")
	registerFunc(&_nw_quic_copy_sec_protocol_metadata, frameworkHandle, "nw_quic_copy_sec_protocol_metadata")
	registerFunc(&_nw_quic_copy_sec_protocol_options, frameworkHandle, "nw_quic_copy_sec_protocol_options")
	registerFunc(&_nw_quic_create_options, frameworkHandle, "nw_quic_create_options")
	registerFunc(&_nw_quic_get_application_error, frameworkHandle, "nw_quic_get_application_error")
	registerFunc(&_nw_quic_get_application_error_reason, frameworkHandle, "nw_quic_get_application_error_reason")
	registerFunc(&_nw_quic_get_idle_timeout, frameworkHandle, "nw_quic_get_idle_timeout")
	registerFunc(&_nw_quic_get_initial_max_data, frameworkHandle, "nw_quic_get_initial_max_data")
	registerFunc(&_nw_quic_get_initial_max_stream_data_bidirectional_local, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_local")
	registerFunc(&_nw_quic_get_initial_max_stream_data_bidirectional_remote, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_remote")
	registerFunc(&_nw_quic_get_initial_max_stream_data_unidirectional, frameworkHandle, "nw_quic_get_initial_max_stream_data_unidirectional")
	registerFunc(&_nw_quic_get_initial_max_streams_bidirectional, frameworkHandle, "nw_quic_get_initial_max_streams_bidirectional")
	registerFunc(&_nw_quic_get_initial_max_streams_unidirectional, frameworkHandle, "nw_quic_get_initial_max_streams_unidirectional")
	registerFunc(&_nw_quic_get_keepalive_interval, frameworkHandle, "nw_quic_get_keepalive_interval")
	registerFunc(&_nw_quic_get_local_max_streams_bidirectional, frameworkHandle, "nw_quic_get_local_max_streams_bidirectional")
	registerFunc(&_nw_quic_get_local_max_streams_unidirectional, frameworkHandle, "nw_quic_get_local_max_streams_unidirectional")
	registerFunc(&_nw_quic_get_max_datagram_frame_size, frameworkHandle, "nw_quic_get_max_datagram_frame_size")
	registerFunc(&_nw_quic_get_max_udp_payload_size, frameworkHandle, "nw_quic_get_max_udp_payload_size")
	registerFunc(&_nw_quic_get_remote_idle_timeout, frameworkHandle, "nw_quic_get_remote_idle_timeout")
	registerFunc(&_nw_quic_get_remote_max_streams_bidirectional, frameworkHandle, "nw_quic_get_remote_max_streams_bidirectional")
	registerFunc(&_nw_quic_get_remote_max_streams_unidirectional, frameworkHandle, "nw_quic_get_remote_max_streams_unidirectional")
	registerFunc(&_nw_quic_get_stream_application_error, frameworkHandle, "nw_quic_get_stream_application_error")
	registerFunc(&_nw_quic_get_stream_id, frameworkHandle, "nw_quic_get_stream_id")
	registerFunc(&_nw_quic_get_stream_is_datagram, frameworkHandle, "nw_quic_get_stream_is_datagram")
	registerFunc(&_nw_quic_get_stream_is_unidirectional, frameworkHandle, "nw_quic_get_stream_is_unidirectional")
	registerFunc(&_nw_quic_get_stream_type, frameworkHandle, "nw_quic_get_stream_type")
	registerFunc(&_nw_quic_get_stream_usable_datagram_frame_size, frameworkHandle, "nw_quic_get_stream_usable_datagram_frame_size")
	registerFunc(&_nw_quic_set_application_error, frameworkHandle, "nw_quic_set_application_error")
	registerFunc(&_nw_quic_set_idle_timeout, frameworkHandle, "nw_quic_set_idle_timeout")
	registerFunc(&_nw_quic_set_initial_max_data, frameworkHandle, "nw_quic_set_initial_max_data")
	registerFunc(&_nw_quic_set_initial_max_stream_data_bidirectional_local, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_local")
	registerFunc(&_nw_quic_set_initial_max_stream_data_bidirectional_remote, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_remote")
	registerFunc(&_nw_quic_set_initial_max_stream_data_unidirectional, frameworkHandle, "nw_quic_set_initial_max_stream_data_unidirectional")
	registerFunc(&_nw_quic_set_initial_max_streams_bidirectional, frameworkHandle, "nw_quic_set_initial_max_streams_bidirectional")
	registerFunc(&_nw_quic_set_initial_max_streams_unidirectional, frameworkHandle, "nw_quic_set_initial_max_streams_unidirectional")
	registerFunc(&_nw_quic_set_keepalive_interval, frameworkHandle, "nw_quic_set_keepalive_interval")
	registerFunc(&_nw_quic_set_local_max_streams_bidirectional, frameworkHandle, "nw_quic_set_local_max_streams_bidirectional")
	registerFunc(&_nw_quic_set_local_max_streams_unidirectional, frameworkHandle, "nw_quic_set_local_max_streams_unidirectional")
	registerFunc(&_nw_quic_set_max_datagram_frame_size, frameworkHandle, "nw_quic_set_max_datagram_frame_size")
	registerFunc(&_nw_quic_set_max_udp_payload_size, frameworkHandle, "nw_quic_set_max_udp_payload_size")
	registerFunc(&_nw_quic_set_stream_application_error, frameworkHandle, "nw_quic_set_stream_application_error")
	registerFunc(&_nw_quic_set_stream_is_datagram, frameworkHandle, "nw_quic_set_stream_is_datagram")
	registerFunc(&_nw_quic_set_stream_is_unidirectional, frameworkHandle, "nw_quic_set_stream_is_unidirectional")
	registerFunc(&_nw_relay_hop_add_additional_http_header_field, frameworkHandle, "nw_relay_hop_add_additional_http_header_field")
	registerFunc(&_nw_relay_hop_create, frameworkHandle, "nw_relay_hop_create")
	registerFunc(&_nw_release, frameworkHandle, "nw_release")
	registerFunc(&_nw_resolution_report_copy_preferred_endpoint, frameworkHandle, "nw_resolution_report_copy_preferred_endpoint")
	registerFunc(&_nw_resolution_report_copy_successful_endpoint, frameworkHandle, "nw_resolution_report_copy_successful_endpoint")
	registerFunc(&_nw_resolution_report_get_endpoint_count, frameworkHandle, "nw_resolution_report_get_endpoint_count")
	registerFunc(&_nw_resolution_report_get_milliseconds, frameworkHandle, "nw_resolution_report_get_milliseconds")
	registerFunc(&_nw_resolution_report_get_protocol, frameworkHandle, "nw_resolution_report_get_protocol")
	registerFunc(&_nw_resolution_report_get_source, frameworkHandle, "nw_resolution_report_get_source")
	registerFunc(&_nw_resolver_config_add_server_address, frameworkHandle, "nw_resolver_config_add_server_address")
	registerFunc(&_nw_resolver_config_create_https, frameworkHandle, "nw_resolver_config_create_https")
	registerFunc(&_nw_resolver_config_create_tls, frameworkHandle, "nw_resolver_config_create_tls")
	registerFunc(&_nw_retain, frameworkHandle, "nw_retain")
	registerFunc(&_nw_tcp_create_options, frameworkHandle, "nw_tcp_create_options")
	registerFunc(&_nw_tcp_get_available_receive_buffer, frameworkHandle, "nw_tcp_get_available_receive_buffer")
	registerFunc(&_nw_tcp_get_available_send_buffer, frameworkHandle, "nw_tcp_get_available_send_buffer")
	registerFunc(&_nw_tcp_options_set_connection_timeout, frameworkHandle, "nw_tcp_options_set_connection_timeout")
	registerFunc(&_nw_tcp_options_set_disable_ack_stretching, frameworkHandle, "nw_tcp_options_set_disable_ack_stretching")
	registerFunc(&_nw_tcp_options_set_disable_ecn, frameworkHandle, "nw_tcp_options_set_disable_ecn")
	registerFunc(&_nw_tcp_options_set_enable_fast_open, frameworkHandle, "nw_tcp_options_set_enable_fast_open")
	registerFunc(&_nw_tcp_options_set_enable_keepalive, frameworkHandle, "nw_tcp_options_set_enable_keepalive")
	registerFunc(&_nw_tcp_options_set_keepalive_count, frameworkHandle, "nw_tcp_options_set_keepalive_count")
	registerFunc(&_nw_tcp_options_set_keepalive_idle_time, frameworkHandle, "nw_tcp_options_set_keepalive_idle_time")
	registerFunc(&_nw_tcp_options_set_keepalive_interval, frameworkHandle, "nw_tcp_options_set_keepalive_interval")
	registerFunc(&_nw_tcp_options_set_maximum_segment_size, frameworkHandle, "nw_tcp_options_set_maximum_segment_size")
	registerFunc(&_nw_tcp_options_set_multipath_force_version, frameworkHandle, "nw_tcp_options_set_multipath_force_version")
	registerFunc(&_nw_tcp_options_set_no_delay, frameworkHandle, "nw_tcp_options_set_no_delay")
	registerFunc(&_nw_tcp_options_set_no_options, frameworkHandle, "nw_tcp_options_set_no_options")
	registerFunc(&_nw_tcp_options_set_no_push, frameworkHandle, "nw_tcp_options_set_no_push")
	registerFunc(&_nw_tcp_options_set_persist_timeout, frameworkHandle, "nw_tcp_options_set_persist_timeout")
	registerFunc(&_nw_tcp_options_set_retransmit_connection_drop_time, frameworkHandle, "nw_tcp_options_set_retransmit_connection_drop_time")
	registerFunc(&_nw_tcp_options_set_retransmit_fin_drop, frameworkHandle, "nw_tcp_options_set_retransmit_fin_drop")
	registerFunc(&_nw_tls_copy_sec_protocol_metadata, frameworkHandle, "nw_tls_copy_sec_protocol_metadata")
	registerFunc(&_nw_tls_copy_sec_protocol_options, frameworkHandle, "nw_tls_copy_sec_protocol_options")
	registerFunc(&_nw_tls_create_options, frameworkHandle, "nw_tls_create_options")
	registerFunc(&_nw_txt_record_access_bytes, frameworkHandle, "nw_txt_record_access_bytes")
	registerFunc(&_nw_txt_record_access_key, frameworkHandle, "nw_txt_record_access_key")
	registerFunc(&_nw_txt_record_apply, frameworkHandle, "nw_txt_record_apply")
	registerFunc(&_nw_txt_record_copy, frameworkHandle, "nw_txt_record_copy")
	registerFunc(&_nw_txt_record_create_dictionary, frameworkHandle, "nw_txt_record_create_dictionary")
	registerFunc(&_nw_txt_record_create_with_bytes, frameworkHandle, "nw_txt_record_create_with_bytes")
	registerFunc(&_nw_txt_record_find_key, frameworkHandle, "nw_txt_record_find_key")
	registerFunc(&_nw_txt_record_get_key_count, frameworkHandle, "nw_txt_record_get_key_count")
	registerFunc(&_nw_txt_record_is_dictionary, frameworkHandle, "nw_txt_record_is_dictionary")
	registerFunc(&_nw_txt_record_is_equal, frameworkHandle, "nw_txt_record_is_equal")
	registerFunc(&_nw_txt_record_remove_key, frameworkHandle, "nw_txt_record_remove_key")
	registerFunc(&_nw_txt_record_set_key, frameworkHandle, "nw_txt_record_set_key")
	registerFunc(&_nw_udp_create_metadata, frameworkHandle, "nw_udp_create_metadata")
	registerFunc(&_nw_udp_create_options, frameworkHandle, "nw_udp_create_options")
	registerFunc(&_nw_udp_options_set_prefer_no_checksum, frameworkHandle, "nw_udp_options_set_prefer_no_checksum")
	registerFunc(&_nw_ws_create_metadata, frameworkHandle, "nw_ws_create_metadata")
	registerFunc(&_nw_ws_create_options, frameworkHandle, "nw_ws_create_options")
	registerFunc(&_nw_ws_metadata_copy_server_response, frameworkHandle, "nw_ws_metadata_copy_server_response")
	registerFunc(&_nw_ws_metadata_get_close_code, frameworkHandle, "nw_ws_metadata_get_close_code")
	registerFunc(&_nw_ws_metadata_get_opcode, frameworkHandle, "nw_ws_metadata_get_opcode")
	registerFunc(&_nw_ws_metadata_set_close_code, frameworkHandle, "nw_ws_metadata_set_close_code")
	registerFunc(&_nw_ws_metadata_set_pong_handler, frameworkHandle, "nw_ws_metadata_set_pong_handler")
	registerFunc(&_nw_ws_options_add_additional_header, frameworkHandle, "nw_ws_options_add_additional_header")
	registerFunc(&_nw_ws_options_add_subprotocol, frameworkHandle, "nw_ws_options_add_subprotocol")
	registerFunc(&_nw_ws_options_set_auto_reply_ping, frameworkHandle, "nw_ws_options_set_auto_reply_ping")
	registerFunc(&_nw_ws_options_set_client_request_handler, frameworkHandle, "nw_ws_options_set_client_request_handler")
	registerFunc(&_nw_ws_options_set_maximum_message_size, frameworkHandle, "nw_ws_options_set_maximum_message_size")
	registerFunc(&_nw_ws_options_set_skip_handshake, frameworkHandle, "nw_ws_options_set_skip_handshake")
	registerFunc(&_nw_ws_request_enumerate_additional_headers, frameworkHandle, "nw_ws_request_enumerate_additional_headers")
	registerFunc(&_nw_ws_request_enumerate_subprotocols, frameworkHandle, "nw_ws_request_enumerate_subprotocols")
	registerFunc(&_nw_ws_response_add_additional_header, frameworkHandle, "nw_ws_response_add_additional_header")
	registerFunc(&_nw_ws_response_create, frameworkHandle, "nw_ws_response_create")
	registerFunc(&_nw_ws_response_enumerate_additional_headers, frameworkHandle, "nw_ws_response_enumerate_additional_headers")
	registerFunc(&_nw_ws_response_get_selected_subprotocol, frameworkHandle, "nw_ws_response_get_selected_subprotocol")
	registerFunc(&_nw_ws_response_get_status, frameworkHandle, "nw_ws_response_get_status")
}
