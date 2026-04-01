// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
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
		return fmt.Sprintf("Network: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Network: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Network: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Network: register symbol %s: %v", name, r)
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

type networkAsyncBlockKey struct {
	owner  objc.ID
	setter string
}

var (
	networkAsyncBlockMu sync.Mutex
	networkAsyncBlocks  = make(map[networkAsyncBlockKey]objc.Block)

	_nw_parameters_configure_protocol_default_configurationSymbol uintptr
	_nw_parameters_configure_protocol_default_configurationErr    error
	_nw_parameters_configure_protocol_disableSymbol               uintptr
	_nw_parameters_configure_protocol_disableErr                  error
)

func retainNetworkAsyncBlock(owner objc.ID, setter string, block objc.Block) {
	if owner == 0 || block == 0 {
		return
	}
	key := networkAsyncBlockKey{owner: owner, setter: setter}
	var old objc.Block
	networkAsyncBlockMu.Lock()
	old = networkAsyncBlocks[key]
	networkAsyncBlocks[key] = block
	networkAsyncBlockMu.Unlock()
	if old != 0 {
		old.Release()
	}
}

func clearNetworkAsyncBlock(owner objc.ID, setter string) {
	if owner == 0 {
		return
	}
	key := networkAsyncBlockKey{owner: owner, setter: setter}
	var old objc.Block
	networkAsyncBlockMu.Lock()
	old = networkAsyncBlocks[key]
	delete(networkAsyncBlocks, key)
	networkAsyncBlockMu.Unlock()
	if old != 0 {
		old.Release()
	}
}

func networkProtocolBlockValue(sym uintptr) unsafe.Pointer {
	if sym == 0 {
		return nil
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(sym))
}

var _nw_advertise_descriptor_copy_txt_record_object func(advertise_descriptor Nw_advertise_descriptor_t) Nw_txt_record_t
var _nw_advertise_descriptor_copy_txt_record_objectErr error

func tryNw_advertise_descriptor_copy_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t) (Nw_txt_record_t, error) {
	if _nw_advertise_descriptor_copy_txt_record_object == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_advertise_descriptor_copy_txt_record_object", "10.15", _nw_advertise_descriptor_copy_txt_record_objectErr)
	}
	return _nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor), nil
}

// Nw_advertise_descriptor_copy_txt_record_object accesses the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_copy_txt_record_object(_:)
func Nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t) Nw_txt_record_t {
	result, callErr := tryNw_advertise_descriptor_copy_txt_record_object(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_advertise_descriptor_create_application_service func(application_service_name string) Nw_advertise_descriptor_t
var _nw_advertise_descriptor_create_application_serviceErr error

func tryNw_advertise_descriptor_create_application_service(application_service_name string) (Nw_advertise_descriptor_t, error) {
	if _nw_advertise_descriptor_create_application_service == nil {
		return *new(Nw_advertise_descriptor_t), symbolCallError("nw_advertise_descriptor_create_application_service", "13.0", _nw_advertise_descriptor_create_application_serviceErr)
	}
	return _nw_advertise_descriptor_create_application_service(application_service_name), nil
}

// Nw_advertise_descriptor_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_application_service(_:)
func Nw_advertise_descriptor_create_application_service(application_service_name string) Nw_advertise_descriptor_t {
	result, callErr := tryNw_advertise_descriptor_create_application_service(application_service_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_advertise_descriptor_create_bonjour_service func(name string, type_ string, domain string) Nw_advertise_descriptor_t
var _nw_advertise_descriptor_create_bonjour_serviceErr error

func tryNw_advertise_descriptor_create_bonjour_service(name string, type_ string, domain string) (Nw_advertise_descriptor_t, error) {
	if _nw_advertise_descriptor_create_bonjour_service == nil {
		return *new(Nw_advertise_descriptor_t), symbolCallError("nw_advertise_descriptor_create_bonjour_service", "10.14", _nw_advertise_descriptor_create_bonjour_serviceErr)
	}
	return _nw_advertise_descriptor_create_bonjour_service(name, type_, domain), nil
}

// Nw_advertise_descriptor_create_bonjour_service initializes a Bonjour service to advertise.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_bonjour_service(_:_:_:)
func Nw_advertise_descriptor_create_bonjour_service(name string, type_ string, domain string) Nw_advertise_descriptor_t {
	result, callErr := tryNw_advertise_descriptor_create_bonjour_service(name, type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_advertise_descriptor_get_application_service_name func(advertise_descriptor Nw_advertise_descriptor_t) *byte
var _nw_advertise_descriptor_get_application_service_nameErr error

func tryNw_advertise_descriptor_get_application_service_name(advertise_descriptor Nw_advertise_descriptor_t) (*byte, error) {
	if _nw_advertise_descriptor_get_application_service_name == nil {
		return nil, symbolCallError("nw_advertise_descriptor_get_application_service_name", "13.0", _nw_advertise_descriptor_get_application_service_nameErr)
	}
	return _nw_advertise_descriptor_get_application_service_name(advertise_descriptor), nil
}

// Nw_advertise_descriptor_get_application_service_name.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_application_service_name(_:)
func Nw_advertise_descriptor_get_application_service_name(advertise_descriptor Nw_advertise_descriptor_t) *byte {
	result, callErr := tryNw_advertise_descriptor_get_application_service_name(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_advertise_descriptor_get_no_auto_rename func(advertise_descriptor Nw_advertise_descriptor_t) bool
var _nw_advertise_descriptor_get_no_auto_renameErr error

func tryNw_advertise_descriptor_get_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t) (bool, error) {
	if _nw_advertise_descriptor_get_no_auto_rename == nil {
		return false, symbolCallError("nw_advertise_descriptor_get_no_auto_rename", "10.14", _nw_advertise_descriptor_get_no_auto_renameErr)
	}
	return _nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor), nil
}

// Nw_advertise_descriptor_get_no_auto_rename checks whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_no_auto_rename(_:)
func Nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t) bool {
	result, callErr := tryNw_advertise_descriptor_get_no_auto_rename(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_advertise_descriptor_set_no_auto_rename func(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool)
var _nw_advertise_descriptor_set_no_auto_renameErr error

func tryNw_advertise_descriptor_set_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool) error {
	if _nw_advertise_descriptor_set_no_auto_rename == nil {
		return symbolCallError("nw_advertise_descriptor_set_no_auto_rename", "10.14", _nw_advertise_descriptor_set_no_auto_renameErr)
	}
	_nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor, no_auto_rename)
	return nil
}

// Nw_advertise_descriptor_set_no_auto_rename sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_no_auto_rename(_:_:)
func Nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool) {
	if callErr := tryNw_advertise_descriptor_set_no_auto_rename(advertise_descriptor, no_auto_rename); callErr != nil {
		panic(callErr)
	}
}

var _nw_advertise_descriptor_set_txt_record func(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr)
var _nw_advertise_descriptor_set_txt_recordErr error

func tryNw_advertise_descriptor_set_txt_record(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr) error {
	if _nw_advertise_descriptor_set_txt_record == nil {
		return symbolCallError("nw_advertise_descriptor_set_txt_record", "10.14", _nw_advertise_descriptor_set_txt_recordErr)
	}
	_nw_advertise_descriptor_set_txt_record(advertise_descriptor, txt_record, txt_length)
	return nil
}

// Nw_advertise_descriptor_set_txt_record sets the TXT record as a raw buffer to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record(_:_:_:)
func Nw_advertise_descriptor_set_txt_record(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr) {
	if callErr := tryNw_advertise_descriptor_set_txt_record(advertise_descriptor, txt_record, txt_length); callErr != nil {
		panic(callErr)
	}
}

var _nw_advertise_descriptor_set_txt_record_object func(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t)
var _nw_advertise_descriptor_set_txt_record_objectErr error

func tryNw_advertise_descriptor_set_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t) error {
	if _nw_advertise_descriptor_set_txt_record_object == nil {
		return symbolCallError("nw_advertise_descriptor_set_txt_record_object", "10.15", _nw_advertise_descriptor_set_txt_record_objectErr)
	}
	_nw_advertise_descriptor_set_txt_record_object(advertise_descriptor, txt_record)
	return nil
}

// Nw_advertise_descriptor_set_txt_record_object sets the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record_object(_:_:)
func Nw_advertise_descriptor_set_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t) {
	if callErr := tryNw_advertise_descriptor_set_txt_record_object(advertise_descriptor, txt_record); callErr != nil {
		panic(callErr)
	}
}

var _nw_browse_descriptor_create_application_service func(application_service_name string) Nw_browse_descriptor_t
var _nw_browse_descriptor_create_application_serviceErr error

func tryNw_browse_descriptor_create_application_service(application_service_name string) (Nw_browse_descriptor_t, error) {
	if _nw_browse_descriptor_create_application_service == nil {
		return *new(Nw_browse_descriptor_t), symbolCallError("nw_browse_descriptor_create_application_service", "13.0", _nw_browse_descriptor_create_application_serviceErr)
	}
	return _nw_browse_descriptor_create_application_service(application_service_name), nil
}

// Nw_browse_descriptor_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_application_service(_:)
func Nw_browse_descriptor_create_application_service(application_service_name string) Nw_browse_descriptor_t {
	result, callErr := tryNw_browse_descriptor_create_application_service(application_service_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_create_bonjour_service func(type_ string, domain string) Nw_browse_descriptor_t
var _nw_browse_descriptor_create_bonjour_serviceErr error

func tryNw_browse_descriptor_create_bonjour_service(type_ string, domain string) (Nw_browse_descriptor_t, error) {
	if _nw_browse_descriptor_create_bonjour_service == nil {
		return *new(Nw_browse_descriptor_t), symbolCallError("nw_browse_descriptor_create_bonjour_service", "10.15", _nw_browse_descriptor_create_bonjour_serviceErr)
	}
	return _nw_browse_descriptor_create_bonjour_service(type_, domain), nil
}

// Nw_browse_descriptor_create_bonjour_service initializes a service descriptor used to discover a Bonjour service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_bonjour_service(_:_:)
func Nw_browse_descriptor_create_bonjour_service(type_ string, domain string) Nw_browse_descriptor_t {
	result, callErr := tryNw_browse_descriptor_create_bonjour_service(type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_get_application_service_name func(descriptor Nw_browse_descriptor_t) *byte
var _nw_browse_descriptor_get_application_service_nameErr error

func tryNw_browse_descriptor_get_application_service_name(descriptor Nw_browse_descriptor_t) (*byte, error) {
	if _nw_browse_descriptor_get_application_service_name == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_application_service_name", "13.0", _nw_browse_descriptor_get_application_service_nameErr)
	}
	return _nw_browse_descriptor_get_application_service_name(descriptor), nil
}

// Nw_browse_descriptor_get_application_service_name.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_application_service_name(_:)
func Nw_browse_descriptor_get_application_service_name(descriptor Nw_browse_descriptor_t) *byte {
	result, callErr := tryNw_browse_descriptor_get_application_service_name(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_get_bonjour_service_domain func(descriptor Nw_browse_descriptor_t) *byte
var _nw_browse_descriptor_get_bonjour_service_domainErr error

func tryNw_browse_descriptor_get_bonjour_service_domain(descriptor Nw_browse_descriptor_t) (*byte, error) {
	if _nw_browse_descriptor_get_bonjour_service_domain == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_bonjour_service_domain", "10.15", _nw_browse_descriptor_get_bonjour_service_domainErr)
	}
	return _nw_browse_descriptor_get_bonjour_service_domain(descriptor), nil
}

// Nw_browse_descriptor_get_bonjour_service_domain accesses the Bonjour service domain set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_domain(_:)
func Nw_browse_descriptor_get_bonjour_service_domain(descriptor Nw_browse_descriptor_t) *byte {
	result, callErr := tryNw_browse_descriptor_get_bonjour_service_domain(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_get_bonjour_service_type func(descriptor Nw_browse_descriptor_t) *byte
var _nw_browse_descriptor_get_bonjour_service_typeErr error

func tryNw_browse_descriptor_get_bonjour_service_type(descriptor Nw_browse_descriptor_t) (*byte, error) {
	if _nw_browse_descriptor_get_bonjour_service_type == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_bonjour_service_type", "10.15", _nw_browse_descriptor_get_bonjour_service_typeErr)
	}
	return _nw_browse_descriptor_get_bonjour_service_type(descriptor), nil
}

// Nw_browse_descriptor_get_bonjour_service_type accesses the Bonjour service type set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_type(_:)
func Nw_browse_descriptor_get_bonjour_service_type(descriptor Nw_browse_descriptor_t) *byte {
	result, callErr := tryNw_browse_descriptor_get_bonjour_service_type(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_get_include_txt_record func(descriptor Nw_browse_descriptor_t) bool
var _nw_browse_descriptor_get_include_txt_recordErr error

func tryNw_browse_descriptor_get_include_txt_record(descriptor Nw_browse_descriptor_t) (bool, error) {
	if _nw_browse_descriptor_get_include_txt_record == nil {
		return false, symbolCallError("nw_browse_descriptor_get_include_txt_record", "10.15", _nw_browse_descriptor_get_include_txt_recordErr)
	}
	return _nw_browse_descriptor_get_include_txt_record(descriptor), nil
}

// Nw_browse_descriptor_get_include_txt_record checks if the browse descriptor requires including associated TXT records with all results.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_include_txt_record(_:)
func Nw_browse_descriptor_get_include_txt_record(descriptor Nw_browse_descriptor_t) bool {
	result, callErr := tryNw_browse_descriptor_get_include_txt_record(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_descriptor_set_include_txt_record func(descriptor Nw_browse_descriptor_t, include_txt_record bool)
var _nw_browse_descriptor_set_include_txt_recordErr error

func tryNw_browse_descriptor_set_include_txt_record(descriptor Nw_browse_descriptor_t, include_txt_record bool) error {
	if _nw_browse_descriptor_set_include_txt_record == nil {
		return symbolCallError("nw_browse_descriptor_set_include_txt_record", "10.15", _nw_browse_descriptor_set_include_txt_recordErr)
	}
	_nw_browse_descriptor_set_include_txt_record(descriptor, include_txt_record)
	return nil
}

// Nw_browse_descriptor_set_include_txt_record requires including associated TXT records with all results generated for this service descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_set_include_txt_record(_:_:)
func Nw_browse_descriptor_set_include_txt_record(descriptor Nw_browse_descriptor_t, include_txt_record bool) {
	if callErr := tryNw_browse_descriptor_set_include_txt_record(descriptor, include_txt_record); callErr != nil {
		panic(callErr)
	}
}

var _nw_browse_result_copy_endpoint func(result Nw_browse_result_t) Nw_endpoint_t
var _nw_browse_result_copy_endpointErr error

func tryNw_browse_result_copy_endpoint(result Nw_browse_result_t) (Nw_endpoint_t, error) {
	if _nw_browse_result_copy_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_browse_result_copy_endpoint", "10.15", _nw_browse_result_copy_endpointErr)
	}
	return _nw_browse_result_copy_endpoint(result), nil
}

// Nw_browse_result_copy_endpoint the discovered service endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_endpoint(_:)
func Nw_browse_result_copy_endpoint(result Nw_browse_result_t) Nw_endpoint_t {
	result0, callErr := tryNw_browse_result_copy_endpoint(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nw_browse_result_copy_txt_record_object func(result Nw_browse_result_t) Nw_txt_record_t
var _nw_browse_result_copy_txt_record_objectErr error

func tryNw_browse_result_copy_txt_record_object(result Nw_browse_result_t) (Nw_txt_record_t, error) {
	if _nw_browse_result_copy_txt_record_object == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_browse_result_copy_txt_record_object", "10.15", _nw_browse_result_copy_txt_record_objectErr)
	}
	return _nw_browse_result_copy_txt_record_object(result), nil
}

// Nw_browse_result_copy_txt_record_object accesses the TXT record associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_txt_record_object(_:)
func Nw_browse_result_copy_txt_record_object(result Nw_browse_result_t) Nw_txt_record_t {
	result0, callErr := tryNw_browse_result_copy_txt_record_object(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nw_browse_result_enumerate_interfaces func(result Nw_browse_result_t, enumerator unsafe.Pointer)
var _nw_browse_result_enumerate_interfacesErr error

func tryNw_browse_result_enumerate_interfaces(result Nw_browse_result_t, enumerator Nw_browse_result_enumerate_interface_t) error {
	if _nw_browse_result_enumerate_interfaces == nil {
		return symbolCallError("nw_browse_result_enumerate_interfaces", "10.15", _nw_browse_result_enumerate_interfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return enumerator(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_browse_result_enumerate_interfaces(result, _block0)
	return nil
}

// Nw_browse_result_enumerate_interfaces enumerates the list of interfaces on which the service was discovered.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interfaces(_:_:)
func Nw_browse_result_enumerate_interfaces(result Nw_browse_result_t, enumerator Nw_browse_result_enumerate_interface_t) {
	if callErr := tryNw_browse_result_enumerate_interfaces(result, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nw_browse_result_get_changes func(old_result Nw_browse_result_t, new_result Nw_browse_result_t) Nw_browse_result_change_t
var _nw_browse_result_get_changesErr error

func tryNw_browse_result_get_changes(old_result Nw_browse_result_t, new_result Nw_browse_result_t) (Nw_browse_result_change_t, error) {
	if _nw_browse_result_get_changes == nil {
		return *new(Nw_browse_result_change_t), symbolCallError("nw_browse_result_get_changes", "10.15", _nw_browse_result_get_changesErr)
	}
	return _nw_browse_result_get_changes(old_result, new_result), nil
}

// Nw_browse_result_get_changes compares two discovered services and calculates changes between them.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_changes(_:_:)
func Nw_browse_result_get_changes(old_result Nw_browse_result_t, new_result Nw_browse_result_t) Nw_browse_result_change_t {
	result, callErr := tryNw_browse_result_get_changes(old_result, new_result)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browse_result_get_interfaces_count func(result Nw_browse_result_t) uintptr
var _nw_browse_result_get_interfaces_countErr error

func tryNw_browse_result_get_interfaces_count(result Nw_browse_result_t) (uintptr, error) {
	if _nw_browse_result_get_interfaces_count == nil {
		return 0, symbolCallError("nw_browse_result_get_interfaces_count", "10.15", _nw_browse_result_get_interfaces_countErr)
	}
	return _nw_browse_result_get_interfaces_count(result), nil
}

// Nw_browse_result_get_interfaces_count accesses the number of interfaces associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_interfaces_count(_:)
func Nw_browse_result_get_interfaces_count(result Nw_browse_result_t) uintptr {
	result0, callErr := tryNw_browse_result_get_interfaces_count(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nw_browser_cancel func(browser Nw_browser_t)
var _nw_browser_cancelErr error

func tryNw_browser_cancel(browser Nw_browser_t) error {
	if _nw_browser_cancel == nil {
		return symbolCallError("nw_browser_cancel", "10.15", _nw_browser_cancelErr)
	}
	_nw_browser_cancel(browser)
	return nil
}

// Nw_browser_cancel stops browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_cancel(_:)
func Nw_browser_cancel(browser Nw_browser_t) {
	if callErr := tryNw_browser_cancel(browser); callErr != nil {
		panic(callErr)
	}
}

var _nw_browser_copy_browse_descriptor func(browser Nw_browser_t) Nw_browse_descriptor_t
var _nw_browser_copy_browse_descriptorErr error

func tryNw_browser_copy_browse_descriptor(browser Nw_browser_t) (Nw_browse_descriptor_t, error) {
	if _nw_browser_copy_browse_descriptor == nil {
		return *new(Nw_browse_descriptor_t), symbolCallError("nw_browser_copy_browse_descriptor", "10.15", _nw_browser_copy_browse_descriptorErr)
	}
	return _nw_browser_copy_browse_descriptor(browser), nil
}

// Nw_browser_copy_browse_descriptor accesses the service descriptor with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_browse_descriptor(_:)
func Nw_browser_copy_browse_descriptor(browser Nw_browser_t) Nw_browse_descriptor_t {
	result, callErr := tryNw_browser_copy_browse_descriptor(browser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browser_copy_parameters func(browser Nw_browser_t) Nw_parameters_t
var _nw_browser_copy_parametersErr error

func tryNw_browser_copy_parameters(browser Nw_browser_t) (Nw_parameters_t, error) {
	if _nw_browser_copy_parameters == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_browser_copy_parameters", "10.15", _nw_browser_copy_parametersErr)
	}
	return _nw_browser_copy_parameters(browser), nil
}

// Nw_browser_copy_parameters accesses the parameters with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_parameters(_:)
func Nw_browser_copy_parameters(browser Nw_browser_t) Nw_parameters_t {
	result, callErr := tryNw_browser_copy_parameters(browser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browser_create func(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) Nw_browser_t
var _nw_browser_createErr error

func tryNw_browser_create(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) (Nw_browser_t, error) {
	if _nw_browser_create == nil {
		return *new(Nw_browser_t), symbolCallError("nw_browser_create", "10.15", _nw_browser_createErr)
	}
	return _nw_browser_create(descriptor, parameters), nil
}

// Nw_browser_create initializes a browser with a type of service to discover.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_create(_:_:)
func Nw_browser_create(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) Nw_browser_t {
	result, callErr := tryNw_browser_create(descriptor, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_browser_set_browse_results_changed_handler func(browser Nw_browser_t, handler unsafe.Pointer)
var _nw_browser_set_browse_results_changed_handlerErr error

func tryNw_browser_set_browse_results_changed_handler(browser Nw_browser_t, handler Nw_browser_browse_results_changed_handler_t) error {
	if _nw_browser_set_browse_results_changed_handler == nil {
		return symbolCallError("nw_browser_set_browse_results_changed_handler", "10.15", _nw_browser_set_browse_results_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 bool) {
		handler(arg0, arg1, arg2)
	})
	retainNetworkAsyncBlock(browser.ID, "nw_browser_set_browse_results_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_browser_set_browse_results_changed_handler(browser, _block0)
	return nil
}

// Nw_browser_set_browse_results_changed_handler sets the handler to receive updates about discovered services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_browse_results_changed_handler(_:_:)
func Nw_browser_set_browse_results_changed_handler(browser Nw_browser_t, handler Nw_browser_browse_results_changed_handler_t) {
	if callErr := tryNw_browser_set_browse_results_changed_handler(browser, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_browser_set_queue func(browser Nw_browser_t, queue uintptr)
var _nw_browser_set_queueErr error

func tryNw_browser_set_queue(browser Nw_browser_t, queue dispatch.Queue) error {
	if _nw_browser_set_queue == nil {
		return symbolCallError("nw_browser_set_queue", "10.15", _nw_browser_set_queueErr)
	}
	_nw_browser_set_queue(browser, uintptr(queue.Handle()))
	return nil
}

// Nw_browser_set_queue sets the queue on which all browser events will be delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_queue(_:_:)
func Nw_browser_set_queue(browser Nw_browser_t, queue dispatch.Queue) {
	if callErr := tryNw_browser_set_queue(browser, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_browser_set_state_changed_handler func(browser Nw_browser_t, state_changed_handler unsafe.Pointer)
var _nw_browser_set_state_changed_handlerErr error

func tryNw_browser_set_state_changed_handler(browser Nw_browser_t, state_changed_handler Nw_browser_state_changed_handler_t) error {
	if _nw_browser_set_state_changed_handler == nil {
		return symbolCallError("nw_browser_set_state_changed_handler", "10.15", _nw_browser_set_state_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwBrowserState, arg1 objectivec.Object) { state_changed_handler(arg0, arg1) })
	retainNetworkAsyncBlock(browser.ID, "nw_browser_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_browser_set_state_changed_handler(browser, _block0)
	return nil
}

// Nw_browser_set_state_changed_handler sets a handler to receive browser state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_state_changed_handler(_:_:)
func Nw_browser_set_state_changed_handler(browser Nw_browser_t, state_changed_handler Nw_browser_state_changed_handler_t) {
	if callErr := tryNw_browser_set_state_changed_handler(browser, state_changed_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_browser_start func(browser Nw_browser_t)
var _nw_browser_startErr error

func tryNw_browser_start(browser Nw_browser_t) error {
	if _nw_browser_start == nil {
		return symbolCallError("nw_browser_start", "10.15", _nw_browser_startErr)
	}
	_nw_browser_start(browser)
	return nil
}

// Nw_browser_start starts browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_start(_:)
func Nw_browser_start(browser Nw_browser_t) {
	if callErr := tryNw_browser_start(browser); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_access_establishment_report func(connection Nw_connection_t, queue uintptr, access_block unsafe.Pointer)
var _nw_connection_access_establishment_reportErr error

func tryNw_connection_access_establishment_report(connection Nw_connection_t, queue dispatch.Queue, access_block Nw_establishment_report_access_block_t) error {
	if _nw_connection_access_establishment_report == nil {
		return symbolCallError("nw_connection_access_establishment_report", "10.15", _nw_connection_access_establishment_reportErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { access_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_access_establishment_report(connection, uintptr(queue.Handle()), _block0)
	return nil
}

// Nw_connection_access_establishment_report requests a copy of the connection’s establishment report once the connection is in the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_access_establishment_report(_:_:_:)
func Nw_connection_access_establishment_report(connection Nw_connection_t, queue dispatch.Queue, access_block Nw_establishment_report_access_block_t) {
	if callErr := tryNw_connection_access_establishment_report(connection, queue, access_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_batch func(connection Nw_connection_t, batch_block unsafe.Pointer)
var _nw_connection_batchErr error

func tryNw_connection_batch(connection Nw_connection_t, batch_block unsafe.Pointer) error {
	if _nw_connection_batch == nil {
		return symbolCallError("nw_connection_batch", "10.14", _nw_connection_batchErr)
	}
	_nw_connection_batch(connection, batch_block)
	return nil
}

// Nw_connection_batch defines a block in which calls to send and receive are processed as a batch to improve performance.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_batch(_:_:)
func Nw_connection_batch(connection Nw_connection_t, batch_block unsafe.Pointer) {
	if callErr := tryNw_connection_batch(connection, batch_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_cancel func(connection Nw_connection_t)
var _nw_connection_cancelErr error

func tryNw_connection_cancel(connection Nw_connection_t) error {
	if _nw_connection_cancel == nil {
		return symbolCallError("nw_connection_cancel", "10.14", _nw_connection_cancelErr)
	}
	_nw_connection_cancel(connection)
	return nil
}

// Nw_connection_cancel cancels the connection and gracefully disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel(_:)
func Nw_connection_cancel(connection Nw_connection_t) {
	if callErr := tryNw_connection_cancel(connection); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_cancel_current_endpoint func(connection Nw_connection_t)
var _nw_connection_cancel_current_endpointErr error

func tryNw_connection_cancel_current_endpoint(connection Nw_connection_t) error {
	if _nw_connection_cancel_current_endpoint == nil {
		return symbolCallError("nw_connection_cancel_current_endpoint", "10.14", _nw_connection_cancel_current_endpointErr)
	}
	_nw_connection_cancel_current_endpoint(connection)
	return nil
}

// Nw_connection_cancel_current_endpoint causes the current endpoint to be rejected, allowing the connection to try another resolved address.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel_current_endpoint(_:)
func Nw_connection_cancel_current_endpoint(connection Nw_connection_t) {
	if callErr := tryNw_connection_cancel_current_endpoint(connection); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_copy_current_path func(connection Nw_connection_t) Nw_path_t
var _nw_connection_copy_current_pathErr error

func tryNw_connection_copy_current_path(connection Nw_connection_t) (Nw_path_t, error) {
	if _nw_connection_copy_current_path == nil {
		return *new(Nw_path_t), symbolCallError("nw_connection_copy_current_path", "10.14", _nw_connection_copy_current_pathErr)
	}
	return _nw_connection_copy_current_path(connection), nil
}

// Nw_connection_copy_current_path accesses the network path the connection is using.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_current_path(_:)
func Nw_connection_copy_current_path(connection Nw_connection_t) Nw_path_t {
	result, callErr := tryNw_connection_copy_current_path(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_copy_description func(connection Nw_connection_t) *byte
var _nw_connection_copy_descriptionErr error

func tryNw_connection_copy_description(connection Nw_connection_t) (*byte, error) {
	if _nw_connection_copy_description == nil {
		return nil, symbolCallError("nw_connection_copy_description", "10.14", _nw_connection_copy_descriptionErr)
	}
	return _nw_connection_copy_description(connection), nil
}

// Nw_connection_copy_description copies the description of the connection as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_description(_:)
func Nw_connection_copy_description(connection Nw_connection_t) *byte {
	result, callErr := tryNw_connection_copy_description(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_copy_endpoint func(connection Nw_connection_t) Nw_endpoint_t
var _nw_connection_copy_endpointErr error

func tryNw_connection_copy_endpoint(connection Nw_connection_t) (Nw_endpoint_t, error) {
	if _nw_connection_copy_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_connection_copy_endpoint", "10.14", _nw_connection_copy_endpointErr)
	}
	return _nw_connection_copy_endpoint(connection), nil
}

// Nw_connection_copy_endpoint accesses the endpoint with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_endpoint(_:)
func Nw_connection_copy_endpoint(connection Nw_connection_t) Nw_endpoint_t {
	result, callErr := tryNw_connection_copy_endpoint(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_copy_parameters func(connection Nw_connection_t) Nw_parameters_t
var _nw_connection_copy_parametersErr error

func tryNw_connection_copy_parameters(connection Nw_connection_t) (Nw_parameters_t, error) {
	if _nw_connection_copy_parameters == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_connection_copy_parameters", "10.14", _nw_connection_copy_parametersErr)
	}
	return _nw_connection_copy_parameters(connection), nil
}

// Nw_connection_copy_parameters accesses the parameters with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_parameters(_:)
func Nw_connection_copy_parameters(connection Nw_connection_t) Nw_parameters_t {
	result, callErr := tryNw_connection_copy_parameters(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_copy_protocol_metadata func(connection Nw_connection_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t
var _nw_connection_copy_protocol_metadataErr error

func tryNw_connection_copy_protocol_metadata(connection Nw_connection_t, definition Nw_protocol_definition_t) (Nw_protocol_metadata_t, error) {
	if _nw_connection_copy_protocol_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_connection_copy_protocol_metadata", "10.14", _nw_connection_copy_protocol_metadataErr)
	}
	return _nw_connection_copy_protocol_metadata(connection, definition), nil
}

// Nw_connection_copy_protocol_metadata retrieves the connection-wide metadata for a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_protocol_metadata(_:_:)
func Nw_connection_copy_protocol_metadata(connection Nw_connection_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	result, callErr := tryNw_connection_copy_protocol_metadata(connection, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_create func(endpoint Nw_endpoint_t, parameters Nw_parameters_t) Nw_connection_t
var _nw_connection_createErr error

func tryNw_connection_create(endpoint Nw_endpoint_t, parameters Nw_parameters_t) (Nw_connection_t, error) {
	if _nw_connection_create == nil {
		return *new(Nw_connection_t), symbolCallError("nw_connection_create", "10.14", _nw_connection_createErr)
	}
	return _nw_connection_create(endpoint, parameters), nil
}

// Nw_connection_create initializes a new connection to a remote endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create(_:_:)
func Nw_connection_create(endpoint Nw_endpoint_t, parameters Nw_parameters_t) Nw_connection_t {
	result, callErr := tryNw_connection_create(endpoint, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_create_new_data_transfer_report func(connection Nw_connection_t) Nw_data_transfer_report_t
var _nw_connection_create_new_data_transfer_reportErr error

func tryNw_connection_create_new_data_transfer_report(connection Nw_connection_t) (Nw_data_transfer_report_t, error) {
	if _nw_connection_create_new_data_transfer_report == nil {
		return *new(Nw_data_transfer_report_t), symbolCallError("nw_connection_create_new_data_transfer_report", "10.15", _nw_connection_create_new_data_transfer_reportErr)
	}
	return _nw_connection_create_new_data_transfer_report(connection), nil
}

// Nw_connection_create_new_data_transfer_report begins a new data transfer report, which can later be collected.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create_new_data_transfer_report(_:)
func Nw_connection_create_new_data_transfer_report(connection Nw_connection_t) Nw_data_transfer_report_t {
	result, callErr := tryNw_connection_create_new_data_transfer_report(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_force_cancel func(connection Nw_connection_t)
var _nw_connection_force_cancelErr error

func tryNw_connection_force_cancel(connection Nw_connection_t) error {
	if _nw_connection_force_cancel == nil {
		return symbolCallError("nw_connection_force_cancel", "10.14", _nw_connection_force_cancelErr)
	}
	_nw_connection_force_cancel(connection)
	return nil
}

// Nw_connection_force_cancel cancels the connection and immediately disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_force_cancel(_:)
func Nw_connection_force_cancel(connection Nw_connection_t) {
	if callErr := tryNw_connection_force_cancel(connection); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_get_maximum_datagram_size func(connection Nw_connection_t) uint32
var _nw_connection_get_maximum_datagram_sizeErr error

func tryNw_connection_get_maximum_datagram_size(connection Nw_connection_t) (uint32, error) {
	if _nw_connection_get_maximum_datagram_size == nil {
		return 0, symbolCallError("nw_connection_get_maximum_datagram_size", "10.14", _nw_connection_get_maximum_datagram_sizeErr)
	}
	return _nw_connection_get_maximum_datagram_size(connection), nil
}

// Nw_connection_get_maximum_datagram_size accesses the maximum size of a datagram message that can be sent on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_get_maximum_datagram_size(_:)
func Nw_connection_get_maximum_datagram_size(connection Nw_connection_t) uint32 {
	result, callErr := tryNw_connection_get_maximum_datagram_size(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_cancel func(group Nw_connection_group_t)
var _nw_connection_group_cancelErr error

func tryNw_connection_group_cancel(group Nw_connection_group_t) error {
	if _nw_connection_group_cancel == nil {
		return symbolCallError("nw_connection_group_cancel", "11.0", _nw_connection_group_cancelErr)
	}
	_nw_connection_group_cancel(group)
	return nil
}

// Nw_connection_group_cancel cancels the connection group object and leaves the network group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_cancel(_:)
func Nw_connection_group_cancel(group Nw_connection_group_t) {
	if callErr := tryNw_connection_group_cancel(group); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_copy_descriptor func(group Nw_connection_group_t) Nw_group_descriptor_t
var _nw_connection_group_copy_descriptorErr error

func tryNw_connection_group_copy_descriptor(group Nw_connection_group_t) (Nw_group_descriptor_t, error) {
	if _nw_connection_group_copy_descriptor == nil {
		return *new(Nw_group_descriptor_t), symbolCallError("nw_connection_group_copy_descriptor", "11.0", _nw_connection_group_copy_descriptorErr)
	}
	return _nw_connection_group_copy_descriptor(group), nil
}

// Nw_connection_group_copy_descriptor accesses the descriptor of the group you use to initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_descriptor(_:)
func Nw_connection_group_copy_descriptor(group Nw_connection_group_t) Nw_group_descriptor_t {
	result, callErr := tryNw_connection_group_copy_descriptor(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_local_endpoint_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t
var _nw_connection_group_copy_local_endpoint_for_messageErr error

func tryNw_connection_group_copy_local_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) (Nw_endpoint_t, error) {
	if _nw_connection_group_copy_local_endpoint_for_message == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_connection_group_copy_local_endpoint_for_message", "11.0", _nw_connection_group_copy_local_endpoint_for_messageErr)
	}
	return _nw_connection_group_copy_local_endpoint_for_message(group, context), nil
}

// Nw_connection_group_copy_local_endpoint_for_message accesses the local address and port you use to receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_local_endpoint_for_message(_:_:)
func Nw_connection_group_copy_local_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t {
	result, callErr := tryNw_connection_group_copy_local_endpoint_for_message(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_parameters func(group Nw_connection_group_t) Nw_parameters_t
var _nw_connection_group_copy_parametersErr error

func tryNw_connection_group_copy_parameters(group Nw_connection_group_t) (Nw_parameters_t, error) {
	if _nw_connection_group_copy_parameters == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_connection_group_copy_parameters", "11.0", _nw_connection_group_copy_parametersErr)
	}
	return _nw_connection_group_copy_parameters(group), nil
}

// Nw_connection_group_copy_parameters accesses the parameters with which you initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_parameters(_:)
func Nw_connection_group_copy_parameters(group Nw_connection_group_t) Nw_parameters_t {
	result, callErr := tryNw_connection_group_copy_parameters(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_path_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_path_t
var _nw_connection_group_copy_path_for_messageErr error

func tryNw_connection_group_copy_path_for_message(group Nw_connection_group_t, context Nw_content_context_t) (Nw_path_t, error) {
	if _nw_connection_group_copy_path_for_message == nil {
		return *new(Nw_path_t), symbolCallError("nw_connection_group_copy_path_for_message", "11.0", _nw_connection_group_copy_path_for_messageErr)
	}
	return _nw_connection_group_copy_path_for_message(group, context), nil
}

// Nw_connection_group_copy_path_for_message accesses the network path on which you receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_path_for_message(_:_:)
func Nw_connection_group_copy_path_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_path_t {
	result, callErr := tryNw_connection_group_copy_path_for_message(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_protocol_metadata func(group Nw_connection_group_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t
var _nw_connection_group_copy_protocol_metadataErr error

func tryNw_connection_group_copy_protocol_metadata(group Nw_connection_group_t, definition Nw_protocol_definition_t) (Nw_protocol_metadata_t, error) {
	if _nw_connection_group_copy_protocol_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_connection_group_copy_protocol_metadata", "12.0", _nw_connection_group_copy_protocol_metadataErr)
	}
	return _nw_connection_group_copy_protocol_metadata(group, definition), nil
}

// Nw_connection_group_copy_protocol_metadata.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata(_:_:)
func Nw_connection_group_copy_protocol_metadata(group Nw_connection_group_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	result, callErr := tryNw_connection_group_copy_protocol_metadata(group, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_protocol_metadata_for_message func(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t
var _nw_connection_group_copy_protocol_metadata_for_messageErr error

func tryNw_connection_group_copy_protocol_metadata_for_message(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) (Nw_protocol_metadata_t, error) {
	if _nw_connection_group_copy_protocol_metadata_for_message == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_connection_group_copy_protocol_metadata_for_message", "12.0", _nw_connection_group_copy_protocol_metadata_for_messageErr)
	}
	return _nw_connection_group_copy_protocol_metadata_for_message(group, context, definition), nil
}

// Nw_connection_group_copy_protocol_metadata_for_message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata_for_message(_:_:_:)
func Nw_connection_group_copy_protocol_metadata_for_message(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	result, callErr := tryNw_connection_group_copy_protocol_metadata_for_message(group, context, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_copy_remote_endpoint_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t
var _nw_connection_group_copy_remote_endpoint_for_messageErr error

func tryNw_connection_group_copy_remote_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) (Nw_endpoint_t, error) {
	if _nw_connection_group_copy_remote_endpoint_for_message == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_connection_group_copy_remote_endpoint_for_message", "11.0", _nw_connection_group_copy_remote_endpoint_for_messageErr)
	}
	return _nw_connection_group_copy_remote_endpoint_for_message(group, context), nil
}

// Nw_connection_group_copy_remote_endpoint_for_message accesses the endpoint that originates the message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_remote_endpoint_for_message(_:_:)
func Nw_connection_group_copy_remote_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t {
	result, callErr := tryNw_connection_group_copy_remote_endpoint_for_message(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_create func(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) Nw_connection_group_t
var _nw_connection_group_createErr error

func tryNw_connection_group_create(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) (Nw_connection_group_t, error) {
	if _nw_connection_group_create == nil {
		return *new(Nw_connection_group_t), symbolCallError("nw_connection_group_create", "11.0", _nw_connection_group_createErr)
	}
	return _nw_connection_group_create(group_descriptor, parameters), nil
}

// Nw_connection_group_create initializes a new connection group with a group identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_create(_:_:)
func Nw_connection_group_create(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) Nw_connection_group_t {
	result, callErr := tryNw_connection_group_create(group_descriptor, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_extract_connection func(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) Nw_connection_t
var _nw_connection_group_extract_connectionErr error

func tryNw_connection_group_extract_connection(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) (Nw_connection_t, error) {
	if _nw_connection_group_extract_connection == nil {
		return *new(Nw_connection_t), symbolCallError("nw_connection_group_extract_connection", "12.0", _nw_connection_group_extract_connectionErr)
	}
	return _nw_connection_group_extract_connection(group, endpoint, protocol_options), nil
}

// Nw_connection_group_extract_connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection(_:_:_:)
func Nw_connection_group_extract_connection(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) Nw_connection_t {
	result, callErr := tryNw_connection_group_extract_connection(group, endpoint, protocol_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_extract_connection_for_message func(group Nw_connection_group_t, context Nw_content_context_t) Nw_connection_t
var _nw_connection_group_extract_connection_for_messageErr error

func tryNw_connection_group_extract_connection_for_message(group Nw_connection_group_t, context Nw_content_context_t) (Nw_connection_t, error) {
	if _nw_connection_group_extract_connection_for_message == nil {
		return *new(Nw_connection_t), symbolCallError("nw_connection_group_extract_connection_for_message", "11.0", _nw_connection_group_extract_connection_for_messageErr)
	}
	return _nw_connection_group_extract_connection_for_message(group, context), nil
}

// Nw_connection_group_extract_connection_for_message converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection_for_message(_:_:)
func Nw_connection_group_extract_connection_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_connection_t {
	result, callErr := tryNw_connection_group_extract_connection_for_message(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_reinsert_extracted_connection func(group Nw_connection_group_t, connection Nw_connection_t) bool
var _nw_connection_group_reinsert_extracted_connectionErr error

func tryNw_connection_group_reinsert_extracted_connection(group Nw_connection_group_t, connection Nw_connection_t) (bool, error) {
	if _nw_connection_group_reinsert_extracted_connection == nil {
		return false, symbolCallError("nw_connection_group_reinsert_extracted_connection", "12.0", _nw_connection_group_reinsert_extracted_connectionErr)
	}
	return _nw_connection_group_reinsert_extracted_connection(group, connection), nil
}

// Nw_connection_group_reinsert_extracted_connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reinsert_extracted_connection(_:_:)
func Nw_connection_group_reinsert_extracted_connection(group Nw_connection_group_t, connection Nw_connection_t) bool {
	result, callErr := tryNw_connection_group_reinsert_extracted_connection(group, connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_connection_group_reply func(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content uintptr)
var _nw_connection_group_replyErr error

func tryNw_connection_group_reply(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content dispatch.Data) error {
	if _nw_connection_group_reply == nil {
		return symbolCallError("nw_connection_group_reply", "11.0", _nw_connection_group_replyErr)
	}
	_nw_connection_group_reply(group, inbound_message, outbound_message, uintptr(content.Handle()))
	return nil
}

// Nw_connection_group_reply sends a reply to the specific endpoint that originates a group message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reply(_:_:_:_:)
func Nw_connection_group_reply(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content dispatch.Data) {
	if callErr := tryNw_connection_group_reply(group, inbound_message, outbound_message, content); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_send_message func(group Nw_connection_group_t, content uintptr, endpoint Nw_endpoint_t, context Nw_content_context_t, completion unsafe.Pointer)
var _nw_connection_group_send_messageErr error

func tryNw_connection_group_send_message(group Nw_connection_group_t, content dispatch.Data, endpoint Nw_endpoint_t, context Nw_content_context_t, completion Nw_connection_group_send_completion_t) error {
	if _nw_connection_group_send_message == nil {
		return symbolCallError("nw_connection_group_send_message", "11.0", _nw_connection_group_send_messageErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { completion(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_group_send_message(group, uintptr(content.Handle()), endpoint, context, _block0)
	return nil
}

// Nw_connection_group_send_message sends data to the entire group, or to a specific member of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_send_message(_:_:_:_:_:)
func Nw_connection_group_send_message(group Nw_connection_group_t, content dispatch.Data, endpoint Nw_endpoint_t, context Nw_content_context_t, completion Nw_connection_group_send_completion_t) {
	if callErr := tryNw_connection_group_send_message(group, content, endpoint, context, completion); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_set_new_connection_handler func(group Nw_connection_group_t, new_connection_handler unsafe.Pointer)
var _nw_connection_group_set_new_connection_handlerErr error

func tryNw_connection_group_set_new_connection_handler(group Nw_connection_group_t, new_connection_handler Nw_connection_group_new_connection_handler_t) error {
	if _nw_connection_group_set_new_connection_handler == nil {
		return symbolCallError("nw_connection_group_set_new_connection_handler", "12.0", _nw_connection_group_set_new_connection_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { new_connection_handler(arg0) })
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_new_connection_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_group_set_new_connection_handler(group, _block0)
	return nil
}

// Nw_connection_group_set_new_connection_handler.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_new_connection_handler(_:_:)
func Nw_connection_group_set_new_connection_handler(group Nw_connection_group_t, new_connection_handler Nw_connection_group_new_connection_handler_t) {
	if callErr := tryNw_connection_group_set_new_connection_handler(group, new_connection_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_set_queue func(group Nw_connection_group_t, queue uintptr)
var _nw_connection_group_set_queueErr error

func tryNw_connection_group_set_queue(group Nw_connection_group_t, queue dispatch.Queue) error {
	if _nw_connection_group_set_queue == nil {
		return symbolCallError("nw_connection_group_set_queue", "11.0", _nw_connection_group_set_queueErr)
	}
	_nw_connection_group_set_queue(group, uintptr(queue.Handle()))
	return nil
}

// Nw_connection_group_set_queue sets the queue on which you handle connection group events.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_queue(_:_:)
func Nw_connection_group_set_queue(group Nw_connection_group_t, queue dispatch.Queue) {
	if callErr := tryNw_connection_group_set_queue(group, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_set_receive_handler func(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler unsafe.Pointer)
var _nw_connection_group_set_receive_handlerErr error

func tryNw_connection_group_set_receive_handler(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler Nw_connection_group_receive_handler_t) error {
	if _nw_connection_group_set_receive_handler == nil {
		return symbolCallError("nw_connection_group_set_receive_handler", "11.0", _nw_connection_group_set_receive_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 bool) {
		receive_handler(arg0, arg1, arg2)
	})
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_receive_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_group_set_receive_handler(group, maximum_message_size, reject_oversized_messages, _block0)
	return nil
}

// Nw_connection_group_set_receive_handler sets a handler that receives inbound messages from members of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_receive_handler(_:_:_:_:)
func Nw_connection_group_set_receive_handler(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler Nw_connection_group_receive_handler_t) {
	if callErr := tryNw_connection_group_set_receive_handler(group, maximum_message_size, reject_oversized_messages, receive_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_set_state_changed_handler func(group Nw_connection_group_t, state_changed_handler unsafe.Pointer)
var _nw_connection_group_set_state_changed_handlerErr error

func tryNw_connection_group_set_state_changed_handler(group Nw_connection_group_t, state_changed_handler Nw_connection_group_state_changed_handler_t) error {
	if _nw_connection_group_set_state_changed_handler == nil {
		return symbolCallError("nw_connection_group_set_state_changed_handler", "11.0", _nw_connection_group_set_state_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwConnectionGroupState, arg1 objectivec.Object) {
		state_changed_handler(arg0, arg1)
	})
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_group_set_state_changed_handler(group, _block0)
	return nil
}

// Nw_connection_group_set_state_changed_handler sets a handler that receives connection group state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_state_changed_handler(_:_:)
func Nw_connection_group_set_state_changed_handler(group Nw_connection_group_t, state_changed_handler Nw_connection_group_state_changed_handler_t) {
	if callErr := tryNw_connection_group_set_state_changed_handler(group, state_changed_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_group_start func(group Nw_connection_group_t)
var _nw_connection_group_startErr error

func tryNw_connection_group_start(group Nw_connection_group_t) error {
	if _nw_connection_group_start == nil {
		return symbolCallError("nw_connection_group_start", "11.0", _nw_connection_group_startErr)
	}
	_nw_connection_group_start(group)
	return nil
}

// Nw_connection_group_start joins the group and registers to receive messages.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_start(_:)
func Nw_connection_group_start(group Nw_connection_group_t) {
	if callErr := tryNw_connection_group_start(group); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_receive func(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion unsafe.Pointer)
var _nw_connection_receiveErr error

func tryNw_connection_receive(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion Nw_connection_receive_completion_t) error {
	if _nw_connection_receive == nil {
		return symbolCallError("nw_connection_receive", "10.14", _nw_connection_receiveErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 bool, arg3 objectivec.Object) {
		completion(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_receive(connection, minimum_incomplete_length, maximum_length, _block0)
	return nil
}

// Nw_connection_receive schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive(_:_:_:_:)
func Nw_connection_receive(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion Nw_connection_receive_completion_t) {
	if callErr := tryNw_connection_receive(connection, minimum_incomplete_length, maximum_length, completion); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_receive_message func(connection Nw_connection_t, completion unsafe.Pointer)
var _nw_connection_receive_messageErr error

func tryNw_connection_receive_message(connection Nw_connection_t, completion Nw_connection_receive_completion_t) error {
	if _nw_connection_receive_message == nil {
		return symbolCallError("nw_connection_receive_message", "10.14", _nw_connection_receive_messageErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 bool, arg3 objectivec.Object) {
		completion(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_receive_message(connection, _block0)
	return nil
}

// Nw_connection_receive_message schedules a single receive completion handler for a complete message, as opposed to a range of bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive_message(_:_:)
func Nw_connection_receive_message(connection Nw_connection_t, completion Nw_connection_receive_completion_t) {
	if callErr := tryNw_connection_receive_message(connection, completion); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_restart func(connection Nw_connection_t)
var _nw_connection_restartErr error

func tryNw_connection_restart(connection Nw_connection_t) error {
	if _nw_connection_restart == nil {
		return symbolCallError("nw_connection_restart", "10.14", _nw_connection_restartErr)
	}
	_nw_connection_restart(connection)
	return nil
}

// Nw_connection_restart restarts a connection that is in the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_restart(_:)
func Nw_connection_restart(connection Nw_connection_t) {
	if callErr := tryNw_connection_restart(connection); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_send func(connection Nw_connection_t, content uintptr, context Nw_content_context_t, is_complete bool, completion unsafe.Pointer)
var _nw_connection_sendErr error

func tryNw_connection_send(connection Nw_connection_t, content dispatch.Data, context Nw_content_context_t, is_complete bool, completion Nw_connection_send_completion_t) error {
	if _nw_connection_send == nil {
		return symbolCallError("nw_connection_send", "10.14", _nw_connection_sendErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { completion(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_send(connection, uintptr(content.Handle()), context, is_complete, _block0)
	return nil
}

// Nw_connection_send sends data on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_send(_:_:_:_:_:)
func Nw_connection_send(connection Nw_connection_t, content dispatch.Data, context Nw_content_context_t, is_complete bool, completion Nw_connection_send_completion_t) {
	if callErr := tryNw_connection_send(connection, content, context, is_complete, completion); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_set_better_path_available_handler func(connection Nw_connection_t, handler unsafe.Pointer)
var _nw_connection_set_better_path_available_handlerErr error

func tryNw_connection_set_better_path_available_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) error {
	if _nw_connection_set_better_path_available_handler == nil {
		return symbolCallError("nw_connection_set_better_path_available_handler", "10.14", _nw_connection_set_better_path_available_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 bool) { handler(arg0) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_better_path_available_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_set_better_path_available_handler(connection, _block0)
	return nil
}

// Nw_connection_set_better_path_available_handler sets a handler that receives updates when an alternative network path is preferred over the current path.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_better_path_available_handler(_:_:)
func Nw_connection_set_better_path_available_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) {
	if callErr := tryNw_connection_set_better_path_available_handler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_set_path_changed_handler func(connection Nw_connection_t, handler unsafe.Pointer)
var _nw_connection_set_path_changed_handlerErr error

func tryNw_connection_set_path_changed_handler(connection Nw_connection_t, handler Nw_connection_path_event_handler_t) error {
	if _nw_connection_set_path_changed_handler == nil {
		return symbolCallError("nw_connection_set_path_changed_handler", "10.14", _nw_connection_set_path_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { handler(arg0) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_path_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_set_path_changed_handler(connection, _block0)
	return nil
}

// Nw_connection_set_path_changed_handler sets a handler that receives network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_path_changed_handler(_:_:)
func Nw_connection_set_path_changed_handler(connection Nw_connection_t, handler Nw_connection_path_event_handler_t) {
	if callErr := tryNw_connection_set_path_changed_handler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_set_queue func(connection Nw_connection_t, queue uintptr)
var _nw_connection_set_queueErr error

func tryNw_connection_set_queue(connection Nw_connection_t, queue dispatch.Queue) error {
	if _nw_connection_set_queue == nil {
		return symbolCallError("nw_connection_set_queue", "10.14", _nw_connection_set_queueErr)
	}
	_nw_connection_set_queue(connection, uintptr(queue.Handle()))
	return nil
}

// Nw_connection_set_queue sets the queue on which all connection events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_queue(_:_:)
func Nw_connection_set_queue(connection Nw_connection_t, queue dispatch.Queue) {
	if callErr := tryNw_connection_set_queue(connection, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_set_state_changed_handler func(connection Nw_connection_t, handler unsafe.Pointer)
var _nw_connection_set_state_changed_handlerErr error

func tryNw_connection_set_state_changed_handler(connection Nw_connection_t, handler Nw_connection_state_changed_handler_t) error {
	if _nw_connection_set_state_changed_handler == nil {
		return symbolCallError("nw_connection_set_state_changed_handler", "10.14", _nw_connection_set_state_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwConnectionState, arg1 objectivec.Object) { handler(arg0, arg1) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_set_state_changed_handler(connection, _block0)
	return nil
}

// Nw_connection_set_state_changed_handler sets a handler to receive connection state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_state_changed_handler(_:_:)
func Nw_connection_set_state_changed_handler(connection Nw_connection_t, handler Nw_connection_state_changed_handler_t) {
	if callErr := tryNw_connection_set_state_changed_handler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_set_viability_changed_handler func(connection Nw_connection_t, handler unsafe.Pointer)
var _nw_connection_set_viability_changed_handlerErr error

func tryNw_connection_set_viability_changed_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) error {
	if _nw_connection_set_viability_changed_handler == nil {
		return symbolCallError("nw_connection_set_viability_changed_handler", "10.14", _nw_connection_set_viability_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 bool) { handler(arg0) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_viability_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_connection_set_viability_changed_handler(connection, _block0)
	return nil
}

// Nw_connection_set_viability_changed_handler sets a handler that receives updates when data can be sent and received.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_viability_changed_handler(_:_:)
func Nw_connection_set_viability_changed_handler(connection Nw_connection_t, handler Nw_connection_boolean_event_handler_t) {
	if callErr := tryNw_connection_set_viability_changed_handler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_connection_start func(connection Nw_connection_t)
var _nw_connection_startErr error

func tryNw_connection_start(connection Nw_connection_t) error {
	if _nw_connection_start == nil {
		return symbolCallError("nw_connection_start", "10.14", _nw_connection_startErr)
	}
	_nw_connection_start(connection)
	return nil
}

// Nw_connection_start starts establishing a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_start(_:)
func Nw_connection_start(connection Nw_connection_t) {
	if callErr := tryNw_connection_start(connection); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_copy_antecedent func(context Nw_content_context_t) Nw_content_context_t
var _nw_content_context_copy_antecedentErr error

func tryNw_content_context_copy_antecedent(context Nw_content_context_t) (Nw_content_context_t, error) {
	if _nw_content_context_copy_antecedent == nil {
		return *new(Nw_content_context_t), symbolCallError("nw_content_context_copy_antecedent", "10.14", _nw_content_context_copy_antecedentErr)
	}
	return _nw_content_context_copy_antecedent(context), nil
}

// Nw_content_context_copy_antecedent accesses the optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_antecedent(_:)
func Nw_content_context_copy_antecedent(context Nw_content_context_t) Nw_content_context_t {
	result, callErr := tryNw_content_context_copy_antecedent(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_copy_protocol_metadata func(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) Nw_protocol_metadata_t
var _nw_content_context_copy_protocol_metadataErr error

func tryNw_content_context_copy_protocol_metadata(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) (Nw_protocol_metadata_t, error) {
	if _nw_content_context_copy_protocol_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_content_context_copy_protocol_metadata", "10.14", _nw_content_context_copy_protocol_metadataErr)
	}
	return _nw_content_context_copy_protocol_metadata(context, protocol_), nil
}

// Nw_content_context_copy_protocol_metadata retreives the metadata associated with a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_protocol_metadata(_:_:)
func Nw_content_context_copy_protocol_metadata(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) Nw_protocol_metadata_t {
	result, callErr := tryNw_content_context_copy_protocol_metadata(context, protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_create func(context_identifier string) Nw_content_context_t
var _nw_content_context_createErr error

func tryNw_content_context_create(context_identifier string) (Nw_content_context_t, error) {
	if _nw_content_context_create == nil {
		return *new(Nw_content_context_t), symbolCallError("nw_content_context_create", "10.14", _nw_content_context_createErr)
	}
	return _nw_content_context_create(context_identifier), nil
}

// Nw_content_context_create initializes a custom message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_create(_:)
func Nw_content_context_create(context_identifier string) Nw_content_context_t {
	result, callErr := tryNw_content_context_create(context_identifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_foreach_protocol_metadata func(context Nw_content_context_t)
var _nw_content_context_foreach_protocol_metadataErr error

func tryNw_content_context_foreach_protocol_metadata(context Nw_content_context_t) error {
	if _nw_content_context_foreach_protocol_metadata == nil {
		return symbolCallError("nw_content_context_foreach_protocol_metadata", "10.14", _nw_content_context_foreach_protocol_metadataErr)
	}
	_nw_content_context_foreach_protocol_metadata(context)
	return nil
}

// Nw_content_context_foreach_protocol_metadata iterates through all protocol metadata associated with the message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_foreach_protocol_metadata(_:_:)
func Nw_content_context_foreach_protocol_metadata(context Nw_content_context_t) {
	if callErr := tryNw_content_context_foreach_protocol_metadata(context); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_get_expiration_milliseconds func(context Nw_content_context_t) uint64
var _nw_content_context_get_expiration_millisecondsErr error

func tryNw_content_context_get_expiration_milliseconds(context Nw_content_context_t) (uint64, error) {
	if _nw_content_context_get_expiration_milliseconds == nil {
		return 0, symbolCallError("nw_content_context_get_expiration_milliseconds", "10.14", _nw_content_context_get_expiration_millisecondsErr)
	}
	return _nw_content_context_get_expiration_milliseconds(context), nil
}

// Nw_content_context_get_expiration_milliseconds accesses the expiration set for this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_expiration_milliseconds(_:)
func Nw_content_context_get_expiration_milliseconds(context Nw_content_context_t) uint64 {
	result, callErr := tryNw_content_context_get_expiration_milliseconds(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_get_identifier func(context Nw_content_context_t) *byte
var _nw_content_context_get_identifierErr error

func tryNw_content_context_get_identifier(context Nw_content_context_t) (*byte, error) {
	if _nw_content_context_get_identifier == nil {
		return nil, symbolCallError("nw_content_context_get_identifier", "10.14", _nw_content_context_get_identifierErr)
	}
	return _nw_content_context_get_identifier(context), nil
}

// Nw_content_context_get_identifier accesses the identifier used to create this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_identifier(_:)
func Nw_content_context_get_identifier(context Nw_content_context_t) *byte {
	result, callErr := tryNw_content_context_get_identifier(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_get_is_final func(context Nw_content_context_t) bool
var _nw_content_context_get_is_finalErr error

func tryNw_content_context_get_is_final(context Nw_content_context_t) (bool, error) {
	if _nw_content_context_get_is_final == nil {
		return false, symbolCallError("nw_content_context_get_is_final", "10.14", _nw_content_context_get_is_finalErr)
	}
	return _nw_content_context_get_is_final(context), nil
}

// Nw_content_context_get_is_final checks whether this context represents the final message being received.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_is_final(_:)
func Nw_content_context_get_is_final(context Nw_content_context_t) bool {
	result, callErr := tryNw_content_context_get_is_final(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_get_relative_priority func(context Nw_content_context_t) float64
var _nw_content_context_get_relative_priorityErr error

func tryNw_content_context_get_relative_priority(context Nw_content_context_t) (float64, error) {
	if _nw_content_context_get_relative_priority == nil {
		return 0.0, symbolCallError("nw_content_context_get_relative_priority", "10.14", _nw_content_context_get_relative_priorityErr)
	}
	return _nw_content_context_get_relative_priority(context), nil
}

// Nw_content_context_get_relative_priority accesses the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_relative_priority(_:)
func Nw_content_context_get_relative_priority(context Nw_content_context_t) float64 {
	result, callErr := tryNw_content_context_get_relative_priority(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_content_context_set_antecedent func(context Nw_content_context_t, antecedent_context Nw_content_context_t)
var _nw_content_context_set_antecedentErr error

func tryNw_content_context_set_antecedent(context Nw_content_context_t, antecedent_context Nw_content_context_t) error {
	if _nw_content_context_set_antecedent == nil {
		return symbolCallError("nw_content_context_set_antecedent", "10.14", _nw_content_context_set_antecedentErr)
	}
	_nw_content_context_set_antecedent(context, antecedent_context)
	return nil
}

// Nw_content_context_set_antecedent set an optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_antecedent(_:_:)
func Nw_content_context_set_antecedent(context Nw_content_context_t, antecedent_context Nw_content_context_t) {
	if callErr := tryNw_content_context_set_antecedent(context, antecedent_context); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_set_expiration_milliseconds func(context Nw_content_context_t, expiration_milliseconds uint64)
var _nw_content_context_set_expiration_millisecondsErr error

func tryNw_content_context_set_expiration_milliseconds(context Nw_content_context_t, expiration_milliseconds uint64) error {
	if _nw_content_context_set_expiration_milliseconds == nil {
		return symbolCallError("nw_content_context_set_expiration_milliseconds", "10.14", _nw_content_context_set_expiration_millisecondsErr)
	}
	_nw_content_context_set_expiration_milliseconds(context, expiration_milliseconds)
	return nil
}

// Nw_content_context_set_expiration_milliseconds sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_expiration_milliseconds(_:_:)
func Nw_content_context_set_expiration_milliseconds(context Nw_content_context_t, expiration_milliseconds uint64) {
	if callErr := tryNw_content_context_set_expiration_milliseconds(context, expiration_milliseconds); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_set_is_final func(context Nw_content_context_t, is_final bool)
var _nw_content_context_set_is_finalErr error

func tryNw_content_context_set_is_final(context Nw_content_context_t, is_final bool) error {
	if _nw_content_context_set_is_final == nil {
		return symbolCallError("nw_content_context_set_is_final", "10.14", _nw_content_context_set_is_finalErr)
	}
	_nw_content_context_set_is_final(context, is_final)
	return nil
}

// Nw_content_context_set_is_final sets a Boolean indicating if this context represents the final message being sent.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_is_final(_:_:)
func Nw_content_context_set_is_final(context Nw_content_context_t, is_final bool) {
	if callErr := tryNw_content_context_set_is_final(context, is_final); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_set_metadata_for_protocol func(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t)
var _nw_content_context_set_metadata_for_protocolErr error

func tryNw_content_context_set_metadata_for_protocol(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t) error {
	if _nw_content_context_set_metadata_for_protocol == nil {
		return symbolCallError("nw_content_context_set_metadata_for_protocol", "10.14", _nw_content_context_set_metadata_for_protocolErr)
	}
	_nw_content_context_set_metadata_for_protocol(context, protocol_metadata)
	return nil
}

// Nw_content_context_set_metadata_for_protocol sets protocol metadata to configure per-message or per-packet properties.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_metadata_for_protocol(_:_:)
func Nw_content_context_set_metadata_for_protocol(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t) {
	if callErr := tryNw_content_context_set_metadata_for_protocol(context, protocol_metadata); callErr != nil {
		panic(callErr)
	}
}

var _nw_content_context_set_relative_priority func(context Nw_content_context_t, relative_priority float64)
var _nw_content_context_set_relative_priorityErr error

func tryNw_content_context_set_relative_priority(context Nw_content_context_t, relative_priority float64) error {
	if _nw_content_context_set_relative_priority == nil {
		return symbolCallError("nw_content_context_set_relative_priority", "10.14", _nw_content_context_set_relative_priorityErr)
	}
	_nw_content_context_set_relative_priority(context, relative_priority)
	return nil
}

// Nw_content_context_set_relative_priority sets the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_relative_priority(_:_:)
func Nw_content_context_set_relative_priority(context Nw_content_context_t, relative_priority float64) {
	if callErr := tryNw_content_context_set_relative_priority(context, relative_priority); callErr != nil {
		panic(callErr)
	}
}

var _nw_data_transfer_report_collect func(report Nw_data_transfer_report_t, queue uintptr, collect_block unsafe.Pointer)
var _nw_data_transfer_report_collectErr error

func tryNw_data_transfer_report_collect(report Nw_data_transfer_report_t, queue dispatch.Queue, collect_block Nw_data_transfer_report_collect_block_t) error {
	if _nw_data_transfer_report_collect == nil {
		return symbolCallError("nw_data_transfer_report_collect", "10.15", _nw_data_transfer_report_collectErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { collect_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_data_transfer_report_collect(report, uintptr(queue.Handle()), _block0)
	return nil
}

// Nw_data_transfer_report_collect stops an outstanding data transfer report and calculates the results.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect(_:_:_:)
func Nw_data_transfer_report_collect(report Nw_data_transfer_report_t, queue dispatch.Queue, collect_block Nw_data_transfer_report_collect_block_t) {
	if callErr := tryNw_data_transfer_report_collect(report, queue, collect_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_data_transfer_report_copy_path_interface func(report Nw_data_transfer_report_t, path_index uint32) Nw_interface_t
var _nw_data_transfer_report_copy_path_interfaceErr error

func tryNw_data_transfer_report_copy_path_interface(report Nw_data_transfer_report_t, path_index uint32) (Nw_interface_t, error) {
	if _nw_data_transfer_report_copy_path_interface == nil {
		return *new(Nw_interface_t), symbolCallError("nw_data_transfer_report_copy_path_interface", "10.15", _nw_data_transfer_report_copy_path_interfaceErr)
	}
	return _nw_data_transfer_report_copy_path_interface(report, path_index), nil
}

// Nw_data_transfer_report_copy_path_interface accesses the network interface the path used.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_copy_path_interface(_:_:)
func Nw_data_transfer_report_copy_path_interface(report Nw_data_transfer_report_t, path_index uint32) Nw_interface_t {
	result, callErr := tryNw_data_transfer_report_copy_path_interface(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_duration_milliseconds func(report Nw_data_transfer_report_t) uint64
var _nw_data_transfer_report_get_duration_millisecondsErr error

func tryNw_data_transfer_report_get_duration_milliseconds(report Nw_data_transfer_report_t) (uint64, error) {
	if _nw_data_transfer_report_get_duration_milliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_duration_milliseconds", "10.15", _nw_data_transfer_report_get_duration_millisecondsErr)
	}
	return _nw_data_transfer_report_get_duration_milliseconds(report), nil
}

// Nw_data_transfer_report_get_duration_milliseconds checks the duration of the data transfer report, from when it was started to when it was collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_duration_milliseconds(_:)
func Nw_data_transfer_report_get_duration_milliseconds(report Nw_data_transfer_report_t) uint64 {
	result, callErr := tryNw_data_transfer_report_get_duration_milliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_path_count func(report Nw_data_transfer_report_t) uint32
var _nw_data_transfer_report_get_path_countErr error

func tryNw_data_transfer_report_get_path_count(report Nw_data_transfer_report_t) (uint32, error) {
	if _nw_data_transfer_report_get_path_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_path_count", "10.15", _nw_data_transfer_report_get_path_countErr)
	}
	return _nw_data_transfer_report_get_path_count(report), nil
}

// Nw_data_transfer_report_get_path_count checks the number of valid paths in the report.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_count(_:)
func Nw_data_transfer_report_get_path_count(report Nw_data_transfer_report_t) uint32 {
	result, callErr := tryNw_data_transfer_report_get_path_count(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_path_radio_type func(report Nw_data_transfer_report_t, path_index uint32) NwInterfaceRadioType
var _nw_data_transfer_report_get_path_radio_typeErr error

func tryNw_data_transfer_report_get_path_radio_type(report Nw_data_transfer_report_t, path_index uint32) (NwInterfaceRadioType, error) {
	if _nw_data_transfer_report_get_path_radio_type == nil {
		return *new(NwInterfaceRadioType), symbolCallError("nw_data_transfer_report_get_path_radio_type", "12.0", _nw_data_transfer_report_get_path_radio_typeErr)
	}
	return _nw_data_transfer_report_get_path_radio_type(report, path_index), nil
}

// Nw_data_transfer_report_get_path_radio_type.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_radio_type(_:_:)
func Nw_data_transfer_report_get_path_radio_type(report Nw_data_transfer_report_t, path_index uint32) NwInterfaceRadioType {
	result, callErr := tryNw_data_transfer_report_get_path_radio_type(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_received_application_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_received_application_byte_countErr error

func tryNw_data_transfer_report_get_received_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_received_application_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_application_byte_count", "10.15", _nw_data_transfer_report_get_received_application_byte_countErr)
	}
	return _nw_data_transfer_report_get_received_application_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_received_application_byte_count accesses the number of bytes the connection delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_application_byte_count(_:_:)
func Nw_data_transfer_report_get_received_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_received_application_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_received_ip_packet_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_received_ip_packet_countErr error

func tryNw_data_transfer_report_get_received_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_received_ip_packet_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_ip_packet_count", "10.15", _nw_data_transfer_report_get_received_ip_packet_countErr)
	}
	return _nw_data_transfer_report_get_received_ip_packet_count(report, path_index), nil
}

// Nw_data_transfer_report_get_received_ip_packet_count accesses the number of IP packets the connection received.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_ip_packet_count(_:_:)
func Nw_data_transfer_report_get_received_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_received_ip_packet_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_received_transport_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_received_transport_byte_countErr error

func tryNw_data_transfer_report_get_received_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_received_transport_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_byte_count", "10.15", _nw_data_transfer_report_get_received_transport_byte_countErr)
	}
	return _nw_data_transfer_report_get_received_transport_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_received_transport_byte_count accesses the number of bytes the transport protocol delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_received_transport_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_received_transport_duplicate_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_received_transport_duplicate_byte_countErr error

func tryNw_data_transfer_report_get_received_transport_duplicate_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_received_transport_duplicate_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_duplicate_byte_count", "10.15", _nw_data_transfer_report_get_received_transport_duplicate_byte_countErr)
	}
	return _nw_data_transfer_report_get_received_transport_duplicate_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_received_transport_duplicate_byte_count accesses the number of duplicated bytes the transport protocol detected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_duplicate_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_duplicate_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_received_transport_duplicate_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_received_transport_out_of_order_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_received_transport_out_of_order_byte_countErr error

func tryNw_data_transfer_report_get_received_transport_out_of_order_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_received_transport_out_of_order_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_out_of_order_byte_count", "10.15", _nw_data_transfer_report_get_received_transport_out_of_order_byte_countErr)
	}
	return _nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_received_transport_out_of_order_byte_count accesses the number of bytes the transport protocol received out of order.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_out_of_order_byte_count(_:_:)
func Nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_received_transport_out_of_order_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_sent_application_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_sent_application_byte_countErr error

func tryNw_data_transfer_report_get_sent_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_sent_application_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_application_byte_count", "10.15", _nw_data_transfer_report_get_sent_application_byte_countErr)
	}
	return _nw_data_transfer_report_get_sent_application_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_sent_application_byte_count accesses the number of bytes sent on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_application_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_sent_application_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_sent_ip_packet_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_sent_ip_packet_countErr error

func tryNw_data_transfer_report_get_sent_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_sent_ip_packet_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_ip_packet_count", "10.15", _nw_data_transfer_report_get_sent_ip_packet_countErr)
	}
	return _nw_data_transfer_report_get_sent_ip_packet_count(report, path_index), nil
}

// Nw_data_transfer_report_get_sent_ip_packet_count accesses the number of IP packets the connection sent.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_ip_packet_count(_:_:)
func Nw_data_transfer_report_get_sent_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_sent_ip_packet_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_sent_transport_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_sent_transport_byte_countErr error

func tryNw_data_transfer_report_get_sent_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_sent_transport_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_transport_byte_count", "10.15", _nw_data_transfer_report_get_sent_transport_byte_countErr)
	}
	return _nw_data_transfer_report_get_sent_transport_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_sent_transport_byte_count accesses the number of bytes sent into the transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_sent_transport_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_sent_transport_retransmitted_byte_countErr error

func tryNw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_transport_retransmitted_byte_count", "10.15", _nw_data_transfer_report_get_sent_transport_retransmitted_byte_countErr)
	}
	return _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report, path_index), nil
}

// Nw_data_transfer_report_get_sent_transport_retransmitted_byte_count accesses the number of bytes the transport protocol retransmitted.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(_:_:)
func Nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_state func(report Nw_data_transfer_report_t) NwDataTransferReportState
var _nw_data_transfer_report_get_stateErr error

func tryNw_data_transfer_report_get_state(report Nw_data_transfer_report_t) (NwDataTransferReportState, error) {
	if _nw_data_transfer_report_get_state == nil {
		return *new(NwDataTransferReportState), symbolCallError("nw_data_transfer_report_get_state", "10.15", _nw_data_transfer_report_get_stateErr)
	}
	return _nw_data_transfer_report_get_state(report), nil
}

// Nw_data_transfer_report_get_state checks whether a data transfer report is collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_state(_:)
func Nw_data_transfer_report_get_state(report Nw_data_transfer_report_t) NwDataTransferReportState {
	result, callErr := tryNw_data_transfer_report_get_state(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_transport_minimum_rtt_millisecondsErr error

func tryNw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_minimum_rtt_milliseconds", "10.15", _nw_data_transfer_report_get_transport_minimum_rtt_millisecondsErr)
	}
	return _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report, path_index), nil
}

// Nw_data_transfer_report_get_transport_minimum_rtt_milliseconds accesses the minimum round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(_:_:)
func Nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_transport_rtt_variance func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_transport_rtt_varianceErr error

func tryNw_data_transfer_report_get_transport_rtt_variance(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_transport_rtt_variance == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_rtt_variance", "10.15", _nw_data_transfer_report_get_transport_rtt_varianceErr)
	}
	return _nw_data_transfer_report_get_transport_rtt_variance(report, path_index), nil
}

// Nw_data_transfer_report_get_transport_rtt_variance accesses the variance of the round-trip time the transport protocol measured.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_rtt_variance(_:_:)
func Nw_data_transfer_report_get_transport_rtt_variance(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_transport_rtt_variance(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds func(report Nw_data_transfer_report_t, path_index uint32) uint64
var _nw_data_transfer_report_get_transport_smoothed_rtt_millisecondsErr error

func tryNw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) (uint64, error) {
	if _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds", "10.15", _nw_data_transfer_report_get_transport_smoothed_rtt_millisecondsErr)
	}
	return _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report, path_index), nil
}

// Nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds accesses the smoothed round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(_:_:)
func Nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	result, callErr := tryNw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_copy_address_string func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_copy_address_stringErr error

func tryNw_endpoint_copy_address_string(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_copy_address_string == nil {
		return nil, symbolCallError("nw_endpoint_copy_address_string", "10.14", _nw_endpoint_copy_address_stringErr)
	}
	return _nw_endpoint_copy_address_string(endpoint), nil
}

// Nw_endpoint_copy_address_string copies the address of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_address_string(_:)
func Nw_endpoint_copy_address_string(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_copy_address_string(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_copy_port_string func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_copy_port_stringErr error

func tryNw_endpoint_copy_port_string(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_copy_port_string == nil {
		return nil, symbolCallError("nw_endpoint_copy_port_string", "10.14", _nw_endpoint_copy_port_stringErr)
	}
	return _nw_endpoint_copy_port_string(endpoint), nil
}

// Nw_endpoint_copy_port_string copies the port of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_port_string(_:)
func Nw_endpoint_copy_port_string(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_copy_port_string(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_copy_txt_record func(endpoint Nw_endpoint_t) Nw_txt_record_t
var _nw_endpoint_copy_txt_recordErr error

func tryNw_endpoint_copy_txt_record(endpoint Nw_endpoint_t) (Nw_txt_record_t, error) {
	if _nw_endpoint_copy_txt_record == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_endpoint_copy_txt_record", "13.0", _nw_endpoint_copy_txt_recordErr)
	}
	return _nw_endpoint_copy_txt_record(endpoint), nil
}

// Nw_endpoint_copy_txt_record.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_txt_record(_:)
func Nw_endpoint_copy_txt_record(endpoint Nw_endpoint_t) Nw_txt_record_t {
	result, callErr := tryNw_endpoint_copy_txt_record(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_create_address func(address uintptr) Nw_endpoint_t
var _nw_endpoint_create_addressErr error

func tryNw_endpoint_create_address(address uintptr) (Nw_endpoint_t, error) {
	if _nw_endpoint_create_address == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_endpoint_create_address", "10.14", _nw_endpoint_create_addressErr)
	}
	return _nw_endpoint_create_address(address), nil
}

// Nw_endpoint_create_address creates a network endpoint with an address structure.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_address(_:)
func Nw_endpoint_create_address(address uintptr) Nw_endpoint_t {
	result, callErr := tryNw_endpoint_create_address(address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_create_bonjour_service func(name string, type_ string, domain string) Nw_endpoint_t
var _nw_endpoint_create_bonjour_serviceErr error

func tryNw_endpoint_create_bonjour_service(name string, type_ string, domain string) (Nw_endpoint_t, error) {
	if _nw_endpoint_create_bonjour_service == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_endpoint_create_bonjour_service", "10.14", _nw_endpoint_create_bonjour_serviceErr)
	}
	return _nw_endpoint_create_bonjour_service(name, type_, domain), nil
}

// Nw_endpoint_create_bonjour_service creates a network endpoint with a Bonjour service name, type, and domain.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_bonjour_service(_:_:_:)
func Nw_endpoint_create_bonjour_service(name string, type_ string, domain string) Nw_endpoint_t {
	result, callErr := tryNw_endpoint_create_bonjour_service(name, type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_create_host func(hostname string, port string) Nw_endpoint_t
var _nw_endpoint_create_hostErr error

func tryNw_endpoint_create_host(hostname string, port string) (Nw_endpoint_t, error) {
	if _nw_endpoint_create_host == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_endpoint_create_host", "10.14", _nw_endpoint_create_hostErr)
	}
	return _nw_endpoint_create_host(hostname, port), nil
}

// Nw_endpoint_create_host creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_host(_:_:)
func Nw_endpoint_create_host(hostname string, port string) Nw_endpoint_t {
	result, callErr := tryNw_endpoint_create_host(hostname, port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_create_url func(url string) Nw_endpoint_t
var _nw_endpoint_create_urlErr error

func tryNw_endpoint_create_url(url string) (Nw_endpoint_t, error) {
	if _nw_endpoint_create_url == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_endpoint_create_url", "10.15", _nw_endpoint_create_urlErr)
	}
	return _nw_endpoint_create_url(url), nil
}

// Nw_endpoint_create_url creates a network endpoint with a URL string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_url(_:)
func Nw_endpoint_create_url(url string) Nw_endpoint_t {
	result, callErr := tryNw_endpoint_create_url(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_address func(endpoint Nw_endpoint_t) objectivec.IObject
var _nw_endpoint_get_addressErr error

func tryNw_endpoint_get_address(endpoint Nw_endpoint_t) (objectivec.IObject, error) {
	if _nw_endpoint_get_address == nil {
		return nil, symbolCallError("nw_endpoint_get_address", "10.14", _nw_endpoint_get_addressErr)
	}
	return _nw_endpoint_get_address(endpoint), nil
}

// Nw_endpoint_get_address accesses the address structure stored in an address endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_address(_:)
func Nw_endpoint_get_address(endpoint Nw_endpoint_t) objectivec.IObject {
	result, callErr := tryNw_endpoint_get_address(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_bonjour_service_domain func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_get_bonjour_service_domainErr error

func tryNw_endpoint_get_bonjour_service_domain(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_get_bonjour_service_domain == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_domain", "10.14", _nw_endpoint_get_bonjour_service_domainErr)
	}
	return _nw_endpoint_get_bonjour_service_domain(endpoint), nil
}

// Nw_endpoint_get_bonjour_service_domain accesses the Bonjour service domain stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_domain(_:)
func Nw_endpoint_get_bonjour_service_domain(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_get_bonjour_service_domain(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_bonjour_service_name func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_get_bonjour_service_nameErr error

func tryNw_endpoint_get_bonjour_service_name(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_get_bonjour_service_name == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_name", "10.14", _nw_endpoint_get_bonjour_service_nameErr)
	}
	return _nw_endpoint_get_bonjour_service_name(endpoint), nil
}

// Nw_endpoint_get_bonjour_service_name accesses the Bonjour service name stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_name(_:)
func Nw_endpoint_get_bonjour_service_name(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_get_bonjour_service_name(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_bonjour_service_type func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_get_bonjour_service_typeErr error

func tryNw_endpoint_get_bonjour_service_type(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_get_bonjour_service_type == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_type", "10.14", _nw_endpoint_get_bonjour_service_typeErr)
	}
	return _nw_endpoint_get_bonjour_service_type(endpoint), nil
}

// Nw_endpoint_get_bonjour_service_type accesses the Bonjour service type stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_type(_:)
func Nw_endpoint_get_bonjour_service_type(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_get_bonjour_service_type(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_hostname func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_get_hostnameErr error

func tryNw_endpoint_get_hostname(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_get_hostname == nil {
		return nil, symbolCallError("nw_endpoint_get_hostname", "10.14", _nw_endpoint_get_hostnameErr)
	}
	return _nw_endpoint_get_hostname(endpoint), nil
}

// Nw_endpoint_get_hostname accesses the hostname stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_hostname(_:)
func Nw_endpoint_get_hostname(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_get_hostname(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_port func(endpoint Nw_endpoint_t) uint16
var _nw_endpoint_get_portErr error

func tryNw_endpoint_get_port(endpoint Nw_endpoint_t) (uint16, error) {
	if _nw_endpoint_get_port == nil {
		return 0, symbolCallError("nw_endpoint_get_port", "10.14", _nw_endpoint_get_portErr)
	}
	return _nw_endpoint_get_port(endpoint), nil
}

// Nw_endpoint_get_port accesses the port stored in an endpoint, in host-byte order.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_port(_:)
func Nw_endpoint_get_port(endpoint Nw_endpoint_t) uint16 {
	result, callErr := tryNw_endpoint_get_port(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_signature func(endpoint Nw_endpoint_t, out_signature_length *uintptr) *uint8
var _nw_endpoint_get_signatureErr error

func tryNw_endpoint_get_signature(endpoint Nw_endpoint_t, out_signature_length *uintptr) (*uint8, error) {
	if _nw_endpoint_get_signature == nil {
		return nil, symbolCallError("nw_endpoint_get_signature", "13.0", _nw_endpoint_get_signatureErr)
	}
	return _nw_endpoint_get_signature(endpoint, out_signature_length), nil
}

// Nw_endpoint_get_signature.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_signature(_:_:)
func Nw_endpoint_get_signature(endpoint Nw_endpoint_t, out_signature_length *uintptr) *uint8 {
	result, callErr := tryNw_endpoint_get_signature(endpoint, out_signature_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_type func(endpoint Nw_endpoint_t) NwEndpointType
var _nw_endpoint_get_typeErr error

func tryNw_endpoint_get_type(endpoint Nw_endpoint_t) (NwEndpointType, error) {
	if _nw_endpoint_get_type == nil {
		return *new(NwEndpointType), symbolCallError("nw_endpoint_get_type", "10.14", _nw_endpoint_get_typeErr)
	}
	return _nw_endpoint_get_type(endpoint), nil
}

// Nw_endpoint_get_type accesses the type of a endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_type(_:)
func Nw_endpoint_get_type(endpoint Nw_endpoint_t) NwEndpointType {
	result, callErr := tryNw_endpoint_get_type(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_endpoint_get_url func(endpoint Nw_endpoint_t) *byte
var _nw_endpoint_get_urlErr error

func tryNw_endpoint_get_url(endpoint Nw_endpoint_t) (*byte, error) {
	if _nw_endpoint_get_url == nil {
		return nil, symbolCallError("nw_endpoint_get_url", "10.15", _nw_endpoint_get_urlErr)
	}
	return _nw_endpoint_get_url(endpoint), nil
}

// Nw_endpoint_get_url accesses the URL string stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_url(_:)
func Nw_endpoint_get_url(endpoint Nw_endpoint_t) *byte {
	result, callErr := tryNw_endpoint_get_url(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_error_copy_cf_error func(err Nw_error_t) corefoundation.CFErrorRef
var _nw_error_copy_cf_errorErr error

func tryNw_error_copy_cf_error(err Nw_error_t) (corefoundation.CFErrorRef, error) {
	if _nw_error_copy_cf_error == nil {
		return 0, symbolCallError("nw_error_copy_cf_error", "10.14", _nw_error_copy_cf_errorErr)
	}
	return _nw_error_copy_cf_error(err), nil
}

// Nw_error_copy_cf_error returns a copy of a network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_copy_cf_error(_:)
func Nw_error_copy_cf_error(err Nw_error_t) corefoundation.CFErrorRef {
	result, callErr := tryNw_error_copy_cf_error(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_error_get_error_code func(err Nw_error_t) int
var _nw_error_get_error_codeErr error

func tryNw_error_get_error_code(err Nw_error_t) (int, error) {
	if _nw_error_get_error_code == nil {
		return 0, symbolCallError("nw_error_get_error_code", "10.14", _nw_error_get_error_codeErr)
	}
	return _nw_error_get_error_code(err), nil
}

// Nw_error_get_error_code accesses the specific code of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_code(_:)
func Nw_error_get_error_code(err Nw_error_t) int {
	result, callErr := tryNw_error_get_error_code(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_error_get_error_domain func(err Nw_error_t) NwErrorDomain
var _nw_error_get_error_domainErr error

func tryNw_error_get_error_domain(err Nw_error_t) (NwErrorDomain, error) {
	if _nw_error_get_error_domain == nil {
		return *new(NwErrorDomain), symbolCallError("nw_error_get_error_domain", "10.14", _nw_error_get_error_domainErr)
	}
	return _nw_error_get_error_domain(err), nil
}

// Nw_error_get_error_domain accesses the domain of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_domain(_:)
func Nw_error_get_error_domain(err Nw_error_t) NwErrorDomain {
	result, callErr := tryNw_error_get_error_domain(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_copy_proxy_endpoint func(report Nw_establishment_report_t) Nw_endpoint_t
var _nw_establishment_report_copy_proxy_endpointErr error

func tryNw_establishment_report_copy_proxy_endpoint(report Nw_establishment_report_t) (Nw_endpoint_t, error) {
	if _nw_establishment_report_copy_proxy_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_establishment_report_copy_proxy_endpoint", "10.15", _nw_establishment_report_copy_proxy_endpointErr)
	}
	return _nw_establishment_report_copy_proxy_endpoint(report), nil
}

// Nw_establishment_report_copy_proxy_endpoint accesses the endpoint of the proxy the connection used.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_copy_proxy_endpoint(_:)
func Nw_establishment_report_copy_proxy_endpoint(report Nw_establishment_report_t) Nw_endpoint_t {
	result, callErr := tryNw_establishment_report_copy_proxy_endpoint(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_enumerate_protocols func(report Nw_establishment_report_t, enumerate_block unsafe.Pointer)
var _nw_establishment_report_enumerate_protocolsErr error

func tryNw_establishment_report_enumerate_protocols(report Nw_establishment_report_t, enumerate_block Nw_report_protocol_enumerator_t) error {
	if _nw_establishment_report_enumerate_protocols == nil {
		return symbolCallError("nw_establishment_report_enumerate_protocols", "10.15", _nw_establishment_report_enumerate_protocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 uint64, arg2 uint64) bool {
		return enumerate_block(arg0, arg1, arg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_establishment_report_enumerate_protocols(report, _block0)
	return nil
}

// Nw_establishment_report_enumerate_protocols iterates a list of protocol handshakes in order from first completed to last completed.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_protocols(_:_:)
func Nw_establishment_report_enumerate_protocols(report Nw_establishment_report_t, enumerate_block Nw_report_protocol_enumerator_t) {
	if callErr := tryNw_establishment_report_enumerate_protocols(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_establishment_report_enumerate_resolution_reports func(report Nw_establishment_report_t, enumerate_block unsafe.Pointer)
var _nw_establishment_report_enumerate_resolution_reportsErr error

func tryNw_establishment_report_enumerate_resolution_reports(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_report_enumerator_t) error {
	if _nw_establishment_report_enumerate_resolution_reports == nil {
		return symbolCallError("nw_establishment_report_enumerate_resolution_reports", "11.0", _nw_establishment_report_enumerate_resolution_reportsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return enumerate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_establishment_report_enumerate_resolution_reports(report, _block0)
	return nil
}

// Nw_establishment_report_enumerate_resolution_reports.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolution_reports(_:_:)
func Nw_establishment_report_enumerate_resolution_reports(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_report_enumerator_t) {
	if callErr := tryNw_establishment_report_enumerate_resolution_reports(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_establishment_report_enumerate_resolutions func(report Nw_establishment_report_t, enumerate_block unsafe.Pointer)
var _nw_establishment_report_enumerate_resolutionsErr error

func tryNw_establishment_report_enumerate_resolutions(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_enumerator_t) error {
	if _nw_establishment_report_enumerate_resolutions == nil {
		return symbolCallError("nw_establishment_report_enumerate_resolutions", "10.15", _nw_establishment_report_enumerate_resolutionsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwReportResolutionSource, arg1 uint64, arg2 uint32, arg3 objectivec.Object, arg4 objectivec.Object) bool {
		return enumerate_block(arg0, arg1, arg2, arg3, arg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_establishment_report_enumerate_resolutions(report, _block0)
	return nil
}

// Nw_establishment_report_enumerate_resolutions iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolutions(_:_:)
func Nw_establishment_report_enumerate_resolutions(report Nw_establishment_report_t, enumerate_block Nw_report_resolution_enumerator_t) {
	if callErr := tryNw_establishment_report_enumerate_resolutions(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_establishment_report_get_attempt_started_after_milliseconds func(report Nw_establishment_report_t) uint64
var _nw_establishment_report_get_attempt_started_after_millisecondsErr error

func tryNw_establishment_report_get_attempt_started_after_milliseconds(report Nw_establishment_report_t) (uint64, error) {
	if _nw_establishment_report_get_attempt_started_after_milliseconds == nil {
		return 0, symbolCallError("nw_establishment_report_get_attempt_started_after_milliseconds", "10.15", _nw_establishment_report_get_attempt_started_after_millisecondsErr)
	}
	return _nw_establishment_report_get_attempt_started_after_milliseconds(report), nil
}

// Nw_establishment_report_get_attempt_started_after_milliseconds accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_attempt_started_after_milliseconds(_:)
func Nw_establishment_report_get_attempt_started_after_milliseconds(report Nw_establishment_report_t) uint64 {
	result, callErr := tryNw_establishment_report_get_attempt_started_after_milliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_get_duration_milliseconds func(report Nw_establishment_report_t) uint64
var _nw_establishment_report_get_duration_millisecondsErr error

func tryNw_establishment_report_get_duration_milliseconds(report Nw_establishment_report_t) (uint64, error) {
	if _nw_establishment_report_get_duration_milliseconds == nil {
		return 0, symbolCallError("nw_establishment_report_get_duration_milliseconds", "10.15", _nw_establishment_report_get_duration_millisecondsErr)
	}
	return _nw_establishment_report_get_duration_milliseconds(report), nil
}

// Nw_establishment_report_get_duration_milliseconds checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_duration_milliseconds(_:)
func Nw_establishment_report_get_duration_milliseconds(report Nw_establishment_report_t) uint64 {
	result, callErr := tryNw_establishment_report_get_duration_milliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_get_previous_attempt_count func(report Nw_establishment_report_t) uint32
var _nw_establishment_report_get_previous_attempt_countErr error

func tryNw_establishment_report_get_previous_attempt_count(report Nw_establishment_report_t) (uint32, error) {
	if _nw_establishment_report_get_previous_attempt_count == nil {
		return 0, symbolCallError("nw_establishment_report_get_previous_attempt_count", "10.15", _nw_establishment_report_get_previous_attempt_countErr)
	}
	return _nw_establishment_report_get_previous_attempt_count(report), nil
}

// Nw_establishment_report_get_previous_attempt_count checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_previous_attempt_count(_:)
func Nw_establishment_report_get_previous_attempt_count(report Nw_establishment_report_t) uint32 {
	result, callErr := tryNw_establishment_report_get_previous_attempt_count(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_get_proxy_configured func(report Nw_establishment_report_t) bool
var _nw_establishment_report_get_proxy_configuredErr error

func tryNw_establishment_report_get_proxy_configured(report Nw_establishment_report_t) (bool, error) {
	if _nw_establishment_report_get_proxy_configured == nil {
		return false, symbolCallError("nw_establishment_report_get_proxy_configured", "10.15", _nw_establishment_report_get_proxy_configuredErr)
	}
	return _nw_establishment_report_get_proxy_configured(report), nil
}

// Nw_establishment_report_get_proxy_configured checks whether a proxy was configured on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_proxy_configured(_:)
func Nw_establishment_report_get_proxy_configured(report Nw_establishment_report_t) bool {
	result, callErr := tryNw_establishment_report_get_proxy_configured(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_establishment_report_get_used_proxy func(report Nw_establishment_report_t) bool
var _nw_establishment_report_get_used_proxyErr error

func tryNw_establishment_report_get_used_proxy(report Nw_establishment_report_t) (bool, error) {
	if _nw_establishment_report_get_used_proxy == nil {
		return false, symbolCallError("nw_establishment_report_get_used_proxy", "10.15", _nw_establishment_report_get_used_proxyErr)
	}
	return _nw_establishment_report_get_used_proxy(report), nil
}

// Nw_establishment_report_get_used_proxy checks whether the connection used a proxy.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_used_proxy(_:)
func Nw_establishment_report_get_used_proxy(report Nw_establishment_report_t) bool {
	result, callErr := tryNw_establishment_report_get_used_proxy(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ethernet_channel_cancel func(ethernet_channel Nw_ethernet_channel_t)
var _nw_ethernet_channel_cancelErr error

func tryNw_ethernet_channel_cancel(ethernet_channel Nw_ethernet_channel_t) error {
	if _nw_ethernet_channel_cancel == nil {
		return symbolCallError("nw_ethernet_channel_cancel", "10.15", _nw_ethernet_channel_cancelErr)
	}
	_nw_ethernet_channel_cancel(ethernet_channel)
	return nil
}

// Nw_ethernet_channel_cancel unregisters the channel from the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_cancel(_:)
func Nw_ethernet_channel_cancel(ethernet_channel Nw_ethernet_channel_t) {
	if callErr := tryNw_ethernet_channel_cancel(ethernet_channel); callErr != nil {
		panic(callErr)
	}
}

var _nw_ethernet_channel_create func(ether_type uint16, interface_ Nw_interface_t) Nw_ethernet_channel_t
var _nw_ethernet_channel_createErr error

func tryNw_ethernet_channel_create(ether_type uint16, interface_ Nw_interface_t) (Nw_ethernet_channel_t, error) {
	if _nw_ethernet_channel_create == nil {
		return *new(Nw_ethernet_channel_t), symbolCallError("nw_ethernet_channel_create", "10.15", _nw_ethernet_channel_createErr)
	}
	return _nw_ethernet_channel_create(ether_type, interface_), nil
}

// Nw_ethernet_channel_create initializes an Ethernet channel on a specific interface with a custom Ethernet type.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create(_:_:)
func Nw_ethernet_channel_create(ether_type uint16, interface_ Nw_interface_t) Nw_ethernet_channel_t {
	result, callErr := tryNw_ethernet_channel_create(ether_type, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ethernet_channel_create_with_parameters func(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) Nw_ethernet_channel_t
var _nw_ethernet_channel_create_with_parametersErr error

func tryNw_ethernet_channel_create_with_parameters(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) (Nw_ethernet_channel_t, error) {
	if _nw_ethernet_channel_create_with_parameters == nil {
		return *new(Nw_ethernet_channel_t), symbolCallError("nw_ethernet_channel_create_with_parameters", "13.0", _nw_ethernet_channel_create_with_parametersErr)
	}
	return _nw_ethernet_channel_create_with_parameters(ether_type, interface_, parameters), nil
}

// Nw_ethernet_channel_create_with_parameters.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create_with_parameters(_:_:_:)
func Nw_ethernet_channel_create_with_parameters(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) Nw_ethernet_channel_t {
	result, callErr := tryNw_ethernet_channel_create_with_parameters(ether_type, interface_, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ethernet_channel_get_maximum_payload_size func(ethernet_channel Nw_ethernet_channel_t) uint32
var _nw_ethernet_channel_get_maximum_payload_sizeErr error

func tryNw_ethernet_channel_get_maximum_payload_size(ethernet_channel Nw_ethernet_channel_t) (uint32, error) {
	if _nw_ethernet_channel_get_maximum_payload_size == nil {
		return 0, symbolCallError("nw_ethernet_channel_get_maximum_payload_size", "13.0", _nw_ethernet_channel_get_maximum_payload_sizeErr)
	}
	return _nw_ethernet_channel_get_maximum_payload_size(ethernet_channel), nil
}

// Nw_ethernet_channel_get_maximum_payload_size.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_get_maximum_payload_size(_:)
func Nw_ethernet_channel_get_maximum_payload_size(ethernet_channel Nw_ethernet_channel_t) uint32 {
	result, callErr := tryNw_ethernet_channel_get_maximum_payload_size(ethernet_channel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ethernet_channel_send func(ethernet_channel Nw_ethernet_channel_t, content uintptr, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion unsafe.Pointer)
var _nw_ethernet_channel_sendErr error

func tryNw_ethernet_channel_send(ethernet_channel Nw_ethernet_channel_t, content dispatch.Data, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion Nw_ethernet_channel_send_completion_t) error {
	if _nw_ethernet_channel_send == nil {
		return symbolCallError("nw_ethernet_channel_send", "10.15", _nw_ethernet_channel_sendErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { completion(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_ethernet_channel_send(ethernet_channel, uintptr(content.Handle()), vlan_tag, remote_address, _block0)
	return nil
}

// Nw_ethernet_channel_send sends a single Ethernet frame over a channel to a specific Ethernet address.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send(_:_:_:_:_:)
func Nw_ethernet_channel_send(ethernet_channel Nw_ethernet_channel_t, content dispatch.Data, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion Nw_ethernet_channel_send_completion_t) {
	if callErr := tryNw_ethernet_channel_send(ethernet_channel, content, vlan_tag, remote_address, completion); callErr != nil {
		panic(callErr)
	}
}

var _nw_ethernet_channel_set_queue func(ethernet_channel Nw_ethernet_channel_t, queue uintptr)
var _nw_ethernet_channel_set_queueErr error

func tryNw_ethernet_channel_set_queue(ethernet_channel Nw_ethernet_channel_t, queue dispatch.Queue) error {
	if _nw_ethernet_channel_set_queue == nil {
		return symbolCallError("nw_ethernet_channel_set_queue", "10.15", _nw_ethernet_channel_set_queueErr)
	}
	_nw_ethernet_channel_set_queue(ethernet_channel, uintptr(queue.Handle()))
	return nil
}

// Nw_ethernet_channel_set_queue sets the queue on which all channel events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_queue(_:_:)
func Nw_ethernet_channel_set_queue(ethernet_channel Nw_ethernet_channel_t, queue dispatch.Queue) {
	if callErr := tryNw_ethernet_channel_set_queue(ethernet_channel, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_ethernet_channel_set_receive_handler func(ethernet_channel Nw_ethernet_channel_t, handler unsafe.Pointer)
var _nw_ethernet_channel_set_receive_handlerErr error

func tryNw_ethernet_channel_set_receive_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_receive_handler_t) error {
	if _nw_ethernet_channel_set_receive_handler == nil {
		return symbolCallError("nw_ethernet_channel_set_receive_handler", "10.15", _nw_ethernet_channel_set_receive_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 uint32, arg2 *uint8, arg3 *uint8) {
		handler(arg0, arg1, arg2, arg3)
	})
	retainNetworkAsyncBlock(ethernet_channel.ID, "nw_ethernet_channel_set_receive_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_ethernet_channel_set_receive_handler(ethernet_channel, _block0)
	return nil
}

// Nw_ethernet_channel_set_receive_handler sets a handler to receive inbound Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_receive_handler(_:_:)
func Nw_ethernet_channel_set_receive_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_receive_handler_t) {
	if callErr := tryNw_ethernet_channel_set_receive_handler(ethernet_channel, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_ethernet_channel_set_state_changed_handler func(ethernet_channel Nw_ethernet_channel_t, handler unsafe.Pointer)
var _nw_ethernet_channel_set_state_changed_handlerErr error

func tryNw_ethernet_channel_set_state_changed_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_state_changed_handler_t) error {
	if _nw_ethernet_channel_set_state_changed_handler == nil {
		return symbolCallError("nw_ethernet_channel_set_state_changed_handler", "10.15", _nw_ethernet_channel_set_state_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwEthernetChannelState, arg1 objectivec.Object) { handler(arg0, arg1) })
	retainNetworkAsyncBlock(ethernet_channel.ID, "nw_ethernet_channel_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_ethernet_channel_set_state_changed_handler(ethernet_channel, _block0)
	return nil
}

// Nw_ethernet_channel_set_state_changed_handler sets a handler to receive channel state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_state_changed_handler(_:_:)
func Nw_ethernet_channel_set_state_changed_handler(ethernet_channel Nw_ethernet_channel_t, handler Nw_ethernet_channel_state_changed_handler_t) {
	if callErr := tryNw_ethernet_channel_set_state_changed_handler(ethernet_channel, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_ethernet_channel_start func(ethernet_channel Nw_ethernet_channel_t)
var _nw_ethernet_channel_startErr error

func tryNw_ethernet_channel_start(ethernet_channel Nw_ethernet_channel_t) error {
	if _nw_ethernet_channel_start == nil {
		return symbolCallError("nw_ethernet_channel_start", "10.15", _nw_ethernet_channel_startErr)
	}
	_nw_ethernet_channel_start(ethernet_channel)
	return nil
}

// Nw_ethernet_channel_start starts the process of registering the channel.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_start(_:)
func Nw_ethernet_channel_start(ethernet_channel Nw_ethernet_channel_t) {
	if callErr := tryNw_ethernet_channel_start(ethernet_channel); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_async func(framer Nw_framer_t, async_block unsafe.Pointer)
var _nw_framer_asyncErr error

func tryNw_framer_async(framer Nw_framer_t, async_block Nw_framer_block_t) error {
	if _nw_framer_async == nil {
		return symbolCallError("nw_framer_async", "10.15", _nw_framer_asyncErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { async_block() })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_async(framer, _block0)
	return nil
}

// Nw_framer_async requests that a block be executed on the connection’s internal scheduling context.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_async(_:_:)
func Nw_framer_async(framer Nw_framer_t, async_block Nw_framer_block_t) {
	if callErr := tryNw_framer_async(framer, async_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_copy_local_endpoint func(framer Nw_framer_t) Nw_endpoint_t
var _nw_framer_copy_local_endpointErr error

func tryNw_framer_copy_local_endpoint(framer Nw_framer_t) (Nw_endpoint_t, error) {
	if _nw_framer_copy_local_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_framer_copy_local_endpoint", "10.15", _nw_framer_copy_local_endpointErr)
	}
	return _nw_framer_copy_local_endpoint(framer), nil
}

// Nw_framer_copy_local_endpoint accesses the local endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_local_endpoint(_:)
func Nw_framer_copy_local_endpoint(framer Nw_framer_t) Nw_endpoint_t {
	result, callErr := tryNw_framer_copy_local_endpoint(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_copy_options func(framer Nw_framer_t) Nw_protocol_options_t
var _nw_framer_copy_optionsErr error

func tryNw_framer_copy_options(framer Nw_framer_t) (Nw_protocol_options_t, error) {
	if _nw_framer_copy_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_framer_copy_options", "12.3", _nw_framer_copy_optionsErr)
	}
	return _nw_framer_copy_options(framer), nil
}

// Nw_framer_copy_options.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_options(_:)
func Nw_framer_copy_options(framer Nw_framer_t) Nw_protocol_options_t {
	result, callErr := tryNw_framer_copy_options(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_copy_parameters func(framer Nw_framer_t) Nw_parameters_t
var _nw_framer_copy_parametersErr error

func tryNw_framer_copy_parameters(framer Nw_framer_t) (Nw_parameters_t, error) {
	if _nw_framer_copy_parameters == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_framer_copy_parameters", "10.15", _nw_framer_copy_parametersErr)
	}
	return _nw_framer_copy_parameters(framer), nil
}

// Nw_framer_copy_parameters accesses the parameters of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_parameters(_:)
func Nw_framer_copy_parameters(framer Nw_framer_t) Nw_parameters_t {
	result, callErr := tryNw_framer_copy_parameters(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_copy_remote_endpoint func(framer Nw_framer_t) Nw_endpoint_t
var _nw_framer_copy_remote_endpointErr error

func tryNw_framer_copy_remote_endpoint(framer Nw_framer_t) (Nw_endpoint_t, error) {
	if _nw_framer_copy_remote_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_framer_copy_remote_endpoint", "10.15", _nw_framer_copy_remote_endpointErr)
	}
	return _nw_framer_copy_remote_endpoint(framer), nil
}

// Nw_framer_copy_remote_endpoint accesses the remote endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_remote_endpoint(_:)
func Nw_framer_copy_remote_endpoint(framer Nw_framer_t) Nw_endpoint_t {
	result, callErr := tryNw_framer_copy_remote_endpoint(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_create_definition func(identifier string, flags uint32, start_handler unsafe.Pointer) Nw_protocol_definition_t
var _nw_framer_create_definitionErr error

func tryNw_framer_create_definition(identifier string, flags uint32, start_handler Nw_framer_start_handler_t) (Nw_protocol_definition_t, error) {
	if _nw_framer_create_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_framer_create_definition", "10.15", _nw_framer_create_definitionErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) NwFramerStartResult { return start_handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_framer_create_definition(identifier, flags, _block0), nil
}

// Nw_framer_create_definition initializes a new protocol definition based on your protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_definition(_:_:_:)
func Nw_framer_create_definition(identifier string, flags uint32, start_handler Nw_framer_start_handler_t) Nw_protocol_definition_t {
	result, callErr := tryNw_framer_create_definition(identifier, flags, start_handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_create_options func(framer_definition Nw_protocol_definition_t) Nw_protocol_options_t
var _nw_framer_create_optionsErr error

func tryNw_framer_create_options(framer_definition Nw_protocol_definition_t) (Nw_protocol_options_t, error) {
	if _nw_framer_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_framer_create_options", "10.15", _nw_framer_create_optionsErr)
	}
	return _nw_framer_create_options(framer_definition), nil
}

// Nw_framer_create_options initializes a set of protocol options with a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_options(_:)
func Nw_framer_create_options(framer_definition Nw_protocol_definition_t) Nw_protocol_options_t {
	result, callErr := tryNw_framer_create_options(framer_definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_deliver_input func(framer Nw_framer_t, input_buffer *uint8, input_length uintptr, message Nw_framer_message_t, is_complete bool)
var _nw_framer_deliver_inputErr error

func tryNw_framer_deliver_input(framer Nw_framer_t, input_buffer *uint8, input_length uintptr, message Nw_framer_message_t, is_complete bool) error {
	if _nw_framer_deliver_input == nil {
		return symbolCallError("nw_framer_deliver_input", "10.15", _nw_framer_deliver_inputErr)
	}
	_nw_framer_deliver_input(framer, input_buffer, input_length, message, is_complete)
	return nil
}

// Nw_framer_deliver_input delivers an inbound message containing arbitrary data from your protocol to the application.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input(_:_:_:_:_:)
func Nw_framer_deliver_input(framer Nw_framer_t, input_buffer *uint8, input_length uintptr, message Nw_framer_message_t, is_complete bool) {
	if callErr := tryNw_framer_deliver_input(framer, input_buffer, input_length, message, is_complete); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_deliver_input_no_copy func(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) bool
var _nw_framer_deliver_input_no_copyErr error

func tryNw_framer_deliver_input_no_copy(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) (bool, error) {
	if _nw_framer_deliver_input_no_copy == nil {
		return false, symbolCallError("nw_framer_deliver_input_no_copy", "10.15", _nw_framer_deliver_input_no_copyErr)
	}
	return _nw_framer_deliver_input_no_copy(framer, input_length, message, is_complete), nil
}

// Nw_framer_deliver_input_no_copy delivers an inbound message containing a specific number of next received bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input_no_copy(_:_:_:_:)
func Nw_framer_deliver_input_no_copy(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) bool {
	result, callErr := tryNw_framer_deliver_input_no_copy(framer, input_length, message, is_complete)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_mark_failed_with_error func(framer Nw_framer_t, error_code int)
var _nw_framer_mark_failed_with_errorErr error

func tryNw_framer_mark_failed_with_error(framer Nw_framer_t, error_code int) error {
	if _nw_framer_mark_failed_with_error == nil {
		return symbolCallError("nw_framer_mark_failed_with_error", "10.15", _nw_framer_mark_failed_with_errorErr)
	}
	_nw_framer_mark_failed_with_error(framer, error_code)
	return nil
}

// Nw_framer_mark_failed_with_error indicates to a connection that your protocol has encountered an error, or has gracefully closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_failed_with_error(_:_:)
func Nw_framer_mark_failed_with_error(framer Nw_framer_t, error_code int) {
	if callErr := tryNw_framer_mark_failed_with_error(framer, error_code); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_mark_ready func(framer Nw_framer_t)
var _nw_framer_mark_readyErr error

func tryNw_framer_mark_ready(framer Nw_framer_t) error {
	if _nw_framer_mark_ready == nil {
		return symbolCallError("nw_framer_mark_ready", "10.15", _nw_framer_mark_readyErr)
	}
	_nw_framer_mark_ready(framer)
	return nil
}

// Nw_framer_mark_ready indicates to a connection that your protocol’s handshake is complete.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_ready(_:)
func Nw_framer_mark_ready(framer Nw_framer_t) {
	if callErr := tryNw_framer_mark_ready(framer); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_message_access_value func(message Nw_framer_message_t, key string, access_value bool) bool
var _nw_framer_message_access_valueErr error

func tryNw_framer_message_access_value(message Nw_framer_message_t, key string, access_value bool) (bool, error) {
	if _nw_framer_message_access_value == nil {
		return false, symbolCallError("nw_framer_message_access_value", "10.15", _nw_framer_message_access_valueErr)
	}
	return _nw_framer_message_access_value(message, key, access_value), nil
}

// Nw_framer_message_access_value accesses a custom value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_access_value(_:_:_:)
func Nw_framer_message_access_value(message Nw_framer_message_t, key string, access_value bool) bool {
	result, callErr := tryNw_framer_message_access_value(message, key, access_value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_message_copy_object_value func(message Nw_framer_message_t, key string) objectivec.Object
var _nw_framer_message_copy_object_valueErr error

func tryNw_framer_message_copy_object_value(message Nw_framer_message_t, key string) (objectivec.Object, error) {
	if _nw_framer_message_copy_object_value == nil {
		return objectivec.Object{}, symbolCallError("nw_framer_message_copy_object_value", "10.15", _nw_framer_message_copy_object_valueErr)
	}
	return _nw_framer_message_copy_object_value(message, key), nil
}

// Nw_framer_message_copy_object_value accesses an NSObject value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_copy_object_value(_:_:)
func Nw_framer_message_copy_object_value(message Nw_framer_message_t, key string) objectivec.Object {
	result, callErr := tryNw_framer_message_copy_object_value(message, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_message_create func(framer Nw_framer_t) Nw_framer_message_t
var _nw_framer_message_createErr error

func tryNw_framer_message_create(framer Nw_framer_t) (Nw_framer_message_t, error) {
	if _nw_framer_message_create == nil {
		return *new(Nw_framer_message_t), symbolCallError("nw_framer_message_create", "10.15", _nw_framer_message_createErr)
	}
	return _nw_framer_message_create(framer), nil
}

// Nw_framer_message_create initializes an empty message from within a framer implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_create(_:)
func Nw_framer_message_create(framer Nw_framer_t) Nw_framer_message_t {
	result, callErr := tryNw_framer_message_create(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_message_set_object_value func(message Nw_framer_message_t, key string, value objectivec.Object)
var _nw_framer_message_set_object_valueErr error

func tryNw_framer_message_set_object_value(message Nw_framer_message_t, key string, value objectivec.Object) error {
	if _nw_framer_message_set_object_value == nil {
		return symbolCallError("nw_framer_message_set_object_value", "10.15", _nw_framer_message_set_object_valueErr)
	}
	_nw_framer_message_set_object_value(message, key, value)
	return nil
}

// Nw_framer_message_set_object_value sets an NSObject value to be stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_object_value(_:_:_:)
func Nw_framer_message_set_object_value(message Nw_framer_message_t, key string, value objectivec.Object) {
	if callErr := tryNw_framer_message_set_object_value(message, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_message_set_value func(message Nw_framer_message_t, key string, value unsafe.Pointer, dispose_value unsafe.Pointer)
var _nw_framer_message_set_valueErr error

func tryNw_framer_message_set_value(message Nw_framer_message_t, key string, value unsafe.Pointer, dispose_value Nw_framer_message_dispose_value_t) error {
	if _nw_framer_message_set_value == nil {
		return symbolCallError("nw_framer_message_set_value", "10.15", _nw_framer_message_set_valueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer) { dispose_value(arg0) })
	retainNetworkAsyncBlock(message.ID, "nw_framer_message_set_value:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_message_set_value(message, key, value, _block0)
	return nil
}

// Nw_framer_message_set_value sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_value(_:_:_:_:)
func Nw_framer_message_set_value(message Nw_framer_message_t, key string, value unsafe.Pointer, dispose_value Nw_framer_message_dispose_value_t) {
	if callErr := tryNw_framer_message_set_value(message, key, value, dispose_value); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_options_copy_object_value func(options Nw_protocol_options_t, key string) objectivec.Object
var _nw_framer_options_copy_object_valueErr error

func tryNw_framer_options_copy_object_value(options Nw_protocol_options_t, key string) (objectivec.Object, error) {
	if _nw_framer_options_copy_object_value == nil {
		return objectivec.Object{}, symbolCallError("nw_framer_options_copy_object_value", "12.3", _nw_framer_options_copy_object_valueErr)
	}
	return _nw_framer_options_copy_object_value(options, key), nil
}

// Nw_framer_options_copy_object_value.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_copy_object_value(_:_:)
func Nw_framer_options_copy_object_value(options Nw_protocol_options_t, key string) objectivec.Object {
	result, callErr := tryNw_framer_options_copy_object_value(options, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_options_set_object_value func(options Nw_protocol_options_t, key string, value objectivec.Object)
var _nw_framer_options_set_object_valueErr error

func tryNw_framer_options_set_object_value(options Nw_protocol_options_t, key string, value objectivec.Object) error {
	if _nw_framer_options_set_object_value == nil {
		return symbolCallError("nw_framer_options_set_object_value", "12.3", _nw_framer_options_set_object_valueErr)
	}
	_nw_framer_options_set_object_value(options, key, value)
	return nil
}

// Nw_framer_options_set_object_value.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_set_object_value(_:_:_:)
func Nw_framer_options_set_object_value(options Nw_protocol_options_t, key string, value objectivec.Object) {
	if callErr := tryNw_framer_options_set_object_value(options, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_parse_input func(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse unsafe.Pointer) bool
var _nw_framer_parse_inputErr error

func tryNw_framer_parse_input(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) (bool, error) {
	if _nw_framer_parse_input == nil {
		return false, symbolCallError("nw_framer_parse_input", "10.15", _nw_framer_parse_inputErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 *uint8, arg1 uint32, arg2 bool) uint64 { return parse(arg0, arg1, arg2) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_framer_parse_input(framer, minimum_incomplete_length, maximum_length, temp_buffer, _block0), nil
}

// Nw_framer_parse_input examines the content of input data while inside your input handler block.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_input(_:_:_:_:_:)
func Nw_framer_parse_input(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool {
	result, callErr := tryNw_framer_parse_input(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_parse_output func(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse unsafe.Pointer) bool
var _nw_framer_parse_outputErr error

func tryNw_framer_parse_output(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) (bool, error) {
	if _nw_framer_parse_output == nil {
		return false, symbolCallError("nw_framer_parse_output", "10.15", _nw_framer_parse_outputErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 *uint8, arg1 uint32, arg2 bool) uint64 { return parse(arg0, arg1, arg2) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_framer_parse_output(framer, minimum_incomplete_length, maximum_length, temp_buffer, _block0), nil
}

// Nw_framer_parse_output examines the content of output data while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_output(_:_:_:_:_:)
func Nw_framer_parse_output(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *uint8, parse Nw_framer_parse_completion_t) bool {
	result, callErr := tryNw_framer_parse_output(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_pass_through_input func(framer Nw_framer_t)
var _nw_framer_pass_through_inputErr error

func tryNw_framer_pass_through_input(framer Nw_framer_t) error {
	if _nw_framer_pass_through_input == nil {
		return symbolCallError("nw_framer_pass_through_input", "10.15", _nw_framer_pass_through_inputErr)
	}
	_nw_framer_pass_through_input(framer)
	return nil
}

// Nw_framer_pass_through_input indicates that your protocol no longer needs to handle input data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_input(_:)
func Nw_framer_pass_through_input(framer Nw_framer_t) {
	if callErr := tryNw_framer_pass_through_input(framer); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_pass_through_output func(framer Nw_framer_t)
var _nw_framer_pass_through_outputErr error

func tryNw_framer_pass_through_output(framer Nw_framer_t) error {
	if _nw_framer_pass_through_output == nil {
		return symbolCallError("nw_framer_pass_through_output", "10.15", _nw_framer_pass_through_outputErr)
	}
	_nw_framer_pass_through_output(framer)
	return nil
}

// Nw_framer_pass_through_output indicates that your protocol no longer needs to handle output data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_output(_:)
func Nw_framer_pass_through_output(framer Nw_framer_t) {
	if callErr := tryNw_framer_pass_through_output(framer); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_prepend_application_protocol func(framer Nw_framer_t, protocol_options Nw_protocol_options_t) bool
var _nw_framer_prepend_application_protocolErr error

func tryNw_framer_prepend_application_protocol(framer Nw_framer_t, protocol_options Nw_protocol_options_t) (bool, error) {
	if _nw_framer_prepend_application_protocol == nil {
		return false, symbolCallError("nw_framer_prepend_application_protocol", "10.15", _nw_framer_prepend_application_protocolErr)
	}
	return _nw_framer_prepend_application_protocol(framer, protocol_options), nil
}

// Nw_framer_prepend_application_protocol dynamically adds another protocol that will run above your protocol after your protocol calls nw_framer_mark_ready(_:).
//
// See: https://developer.apple.com/documentation/Network/nw_framer_prepend_application_protocol(_:_:)
func Nw_framer_prepend_application_protocol(framer Nw_framer_t, protocol_options Nw_protocol_options_t) bool {
	result, callErr := tryNw_framer_prepend_application_protocol(framer, protocol_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_protocol_create_message func(definition Nw_protocol_definition_t) Nw_framer_message_t
var _nw_framer_protocol_create_messageErr error

func tryNw_framer_protocol_create_message(definition Nw_protocol_definition_t) (Nw_framer_message_t, error) {
	if _nw_framer_protocol_create_message == nil {
		return *new(Nw_framer_message_t), symbolCallError("nw_framer_protocol_create_message", "10.15", _nw_framer_protocol_create_messageErr)
	}
	return _nw_framer_protocol_create_message(definition), nil
}

// Nw_framer_protocol_create_message initializes an empty message for a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_protocol_create_message(_:)
func Nw_framer_protocol_create_message(definition Nw_protocol_definition_t) Nw_framer_message_t {
	result, callErr := tryNw_framer_protocol_create_message(definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_framer_schedule_wakeup func(framer Nw_framer_t, milliseconds uint64)
var _nw_framer_schedule_wakeupErr error

func tryNw_framer_schedule_wakeup(framer Nw_framer_t, milliseconds uint64) error {
	if _nw_framer_schedule_wakeup == nil {
		return symbolCallError("nw_framer_schedule_wakeup", "10.15", _nw_framer_schedule_wakeupErr)
	}
	_nw_framer_schedule_wakeup(framer, milliseconds)
	return nil
}

// Nw_framer_schedule_wakeup requests that the nw_framer_wakeup_handler_t be called on your protocol at a specific time in the future.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_schedule_wakeup(_:_:)
func Nw_framer_schedule_wakeup(framer Nw_framer_t, milliseconds uint64) {
	if callErr := tryNw_framer_schedule_wakeup(framer, milliseconds); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_set_cleanup_handler func(framer Nw_framer_t, cleanup_handler unsafe.Pointer)
var _nw_framer_set_cleanup_handlerErr error

func tryNw_framer_set_cleanup_handler(framer Nw_framer_t, cleanup_handler Nw_framer_cleanup_handler_t) error {
	if _nw_framer_set_cleanup_handler == nil {
		return symbolCallError("nw_framer_set_cleanup_handler", "10.15", _nw_framer_set_cleanup_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { cleanup_handler(arg0) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_cleanup_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_set_cleanup_handler(framer, _block0)
	return nil
}

// Nw_framer_set_cleanup_handler sets a block to handle the final cleanup of allocations made by your protocol instance.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_cleanup_handler(_:_:)
func Nw_framer_set_cleanup_handler(framer Nw_framer_t, cleanup_handler Nw_framer_cleanup_handler_t) {
	if callErr := tryNw_framer_set_cleanup_handler(framer, cleanup_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_set_input_handler func(framer Nw_framer_t, input_handler unsafe.Pointer)
var _nw_framer_set_input_handlerErr error

func tryNw_framer_set_input_handler(framer Nw_framer_t, input_handler Nw_framer_input_handler_t) error {
	if _nw_framer_set_input_handler == nil {
		return symbolCallError("nw_framer_set_input_handler", "10.15", _nw_framer_set_input_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) uint64 { return input_handler(arg0) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_input_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_set_input_handler(framer, _block0)
	return nil
}

// Nw_framer_set_input_handler sets a block to handle new inbound data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_input_handler(_:_:)
func Nw_framer_set_input_handler(framer Nw_framer_t, input_handler Nw_framer_input_handler_t) {
	if callErr := tryNw_framer_set_input_handler(framer, input_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_set_output_handler func(framer Nw_framer_t, output_handler unsafe.Pointer)
var _nw_framer_set_output_handlerErr error

func tryNw_framer_set_output_handler(framer Nw_framer_t, output_handler Nw_framer_output_handler_t) error {
	if _nw_framer_set_output_handler == nil {
		return symbolCallError("nw_framer_set_output_handler", "10.15", _nw_framer_set_output_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 uint32, arg3 bool) {
		output_handler(arg0, arg1, arg2, arg3)
	})
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_output_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_set_output_handler(framer, _block0)
	return nil
}

// Nw_framer_set_output_handler sets a block to handle new outbound messages.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_output_handler(_:_:)
func Nw_framer_set_output_handler(framer Nw_framer_t, output_handler Nw_framer_output_handler_t) {
	if callErr := tryNw_framer_set_output_handler(framer, output_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_set_stop_handler func(framer Nw_framer_t, stop_handler unsafe.Pointer)
var _nw_framer_set_stop_handlerErr error

func tryNw_framer_set_stop_handler(framer Nw_framer_t, stop_handler Nw_framer_stop_handler_t) error {
	if _nw_framer_set_stop_handler == nil {
		return symbolCallError("nw_framer_set_stop_handler", "10.15", _nw_framer_set_stop_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return stop_handler(arg0) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_stop_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_set_stop_handler(framer, _block0)
	return nil
}

// Nw_framer_set_stop_handler sets a block to handle when the connection is being closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_stop_handler(_:_:)
func Nw_framer_set_stop_handler(framer Nw_framer_t, stop_handler Nw_framer_stop_handler_t) {
	if callErr := tryNw_framer_set_stop_handler(framer, stop_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_set_wakeup_handler func(framer Nw_framer_t, wakeup_handler unsafe.Pointer)
var _nw_framer_set_wakeup_handlerErr error

func tryNw_framer_set_wakeup_handler(framer Nw_framer_t, wakeup_handler Nw_framer_wakeup_handler_t) error {
	if _nw_framer_set_wakeup_handler == nil {
		return symbolCallError("nw_framer_set_wakeup_handler", "10.15", _nw_framer_set_wakeup_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { wakeup_handler(arg0) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_wakeup_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_framer_set_wakeup_handler(framer, _block0)
	return nil
}

// Nw_framer_set_wakeup_handler sets a handler to receive scheduled wakeup events.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_wakeup_handler(_:_:)
func Nw_framer_set_wakeup_handler(framer Nw_framer_t, wakeup_handler Nw_framer_wakeup_handler_t) {
	if callErr := tryNw_framer_set_wakeup_handler(framer, wakeup_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_write_output func(framer Nw_framer_t, output_buffer *uint8, output_length uintptr)
var _nw_framer_write_outputErr error

func tryNw_framer_write_output(framer Nw_framer_t, output_buffer *uint8, output_length uintptr) error {
	if _nw_framer_write_output == nil {
		return symbolCallError("nw_framer_write_output", "10.15", _nw_framer_write_outputErr)
	}
	_nw_framer_write_output(framer, output_buffer, output_length)
	return nil
}

// Nw_framer_write_output sends arbitrary output data in a buffer from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output(_:_:_:)
func Nw_framer_write_output(framer Nw_framer_t, output_buffer *uint8, output_length uintptr) {
	if callErr := tryNw_framer_write_output(framer, output_buffer, output_length); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_write_output_data func(framer Nw_framer_t, output_data uintptr)
var _nw_framer_write_output_dataErr error

func tryNw_framer_write_output_data(framer Nw_framer_t, output_data dispatch.Data) error {
	if _nw_framer_write_output_data == nil {
		return symbolCallError("nw_framer_write_output_data", "10.15", _nw_framer_write_output_dataErr)
	}
	_nw_framer_write_output_data(framer, uintptr(output_data.Handle()))
	return nil
}

// Nw_framer_write_output_data sends arbitrary output data from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_data(_:_:)
func Nw_framer_write_output_data(framer Nw_framer_t, output_data dispatch.Data) {
	if callErr := tryNw_framer_write_output_data(framer, output_data); callErr != nil {
		panic(callErr)
	}
}

var _nw_framer_write_output_no_copy func(framer Nw_framer_t, output_length uintptr) bool
var _nw_framer_write_output_no_copyErr error

func tryNw_framer_write_output_no_copy(framer Nw_framer_t, output_length uintptr) (bool, error) {
	if _nw_framer_write_output_no_copy == nil {
		return false, symbolCallError("nw_framer_write_output_no_copy", "10.15", _nw_framer_write_output_no_copyErr)
	}
	return _nw_framer_write_output_no_copy(framer, output_length), nil
}

// Nw_framer_write_output_no_copy sends a specific number of bytes from a message while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_no_copy(_:_:)
func Nw_framer_write_output_no_copy(framer Nw_framer_t, output_length uintptr) bool {
	result, callErr := tryNw_framer_write_output_no_copy(framer, output_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_group_descriptor_add_endpoint func(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) bool
var _nw_group_descriptor_add_endpointErr error

func tryNw_group_descriptor_add_endpoint(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) (bool, error) {
	if _nw_group_descriptor_add_endpoint == nil {
		return false, symbolCallError("nw_group_descriptor_add_endpoint", "11.0", _nw_group_descriptor_add_endpointErr)
	}
	return _nw_group_descriptor_add_endpoint(descriptor, endpoint), nil
}

// Nw_group_descriptor_add_endpoint adds a multicast address endpoint you specify to define an extra IP multicast group to join.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_add_endpoint(_:_:)
func Nw_group_descriptor_add_endpoint(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) bool {
	result, callErr := tryNw_group_descriptor_add_endpoint(descriptor, endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_group_descriptor_create_multicast func(multicast_group Nw_endpoint_t) Nw_group_descriptor_t
var _nw_group_descriptor_create_multicastErr error

func tryNw_group_descriptor_create_multicast(multicast_group Nw_endpoint_t) (Nw_group_descriptor_t, error) {
	if _nw_group_descriptor_create_multicast == nil {
		return *new(Nw_group_descriptor_t), symbolCallError("nw_group_descriptor_create_multicast", "11.0", _nw_group_descriptor_create_multicastErr)
	}
	return _nw_group_descriptor_create_multicast(multicast_group), nil
}

// Nw_group_descriptor_create_multicast creates group descriptor you use to join an IP multicast group on a local network.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multicast(_:)
func Nw_group_descriptor_create_multicast(multicast_group Nw_endpoint_t) Nw_group_descriptor_t {
	result, callErr := tryNw_group_descriptor_create_multicast(multicast_group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_group_descriptor_create_multiplex func(remote_endpoint Nw_endpoint_t) Nw_group_descriptor_t
var _nw_group_descriptor_create_multiplexErr error

func tryNw_group_descriptor_create_multiplex(remote_endpoint Nw_endpoint_t) (Nw_group_descriptor_t, error) {
	if _nw_group_descriptor_create_multiplex == nil {
		return *new(Nw_group_descriptor_t), symbolCallError("nw_group_descriptor_create_multiplex", "12.0", _nw_group_descriptor_create_multiplexErr)
	}
	return _nw_group_descriptor_create_multiplex(remote_endpoint), nil
}

// Nw_group_descriptor_create_multiplex.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multiplex(_:)
func Nw_group_descriptor_create_multiplex(remote_endpoint Nw_endpoint_t) Nw_group_descriptor_t {
	result, callErr := tryNw_group_descriptor_create_multiplex(remote_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_group_descriptor_enumerate_endpoints func(descriptor Nw_group_descriptor_t, enumerate_block unsafe.Pointer)
var _nw_group_descriptor_enumerate_endpointsErr error

func tryNw_group_descriptor_enumerate_endpoints(descriptor Nw_group_descriptor_t, enumerate_block Nw_group_descriptor_enumerate_endpoints_block_t) error {
	if _nw_group_descriptor_enumerate_endpoints == nil {
		return symbolCallError("nw_group_descriptor_enumerate_endpoints", "11.0", _nw_group_descriptor_enumerate_endpointsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return enumerate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_group_descriptor_enumerate_endpoints(descriptor, _block0)
	return nil
}

// Nw_group_descriptor_enumerate_endpoints sets a handler to list all endpoints added to the group descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints(_:_:)
func Nw_group_descriptor_enumerate_endpoints(descriptor Nw_group_descriptor_t, enumerate_block Nw_group_descriptor_enumerate_endpoints_block_t) {
	if callErr := tryNw_group_descriptor_enumerate_endpoints(descriptor, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_interface_get_index func(interface_ Nw_interface_t) uint32
var _nw_interface_get_indexErr error

func tryNw_interface_get_index(interface_ Nw_interface_t) (uint32, error) {
	if _nw_interface_get_index == nil {
		return 0, symbolCallError("nw_interface_get_index", "10.14", _nw_interface_get_indexErr)
	}
	return _nw_interface_get_index(interface_), nil
}

// Nw_interface_get_index accesses the system interface index associated with the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_index(_:)
func Nw_interface_get_index(interface_ Nw_interface_t) uint32 {
	result, callErr := tryNw_interface_get_index(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_interface_get_name func(interface_ Nw_interface_t) *byte
var _nw_interface_get_nameErr error

func tryNw_interface_get_name(interface_ Nw_interface_t) (*byte, error) {
	if _nw_interface_get_name == nil {
		return nil, symbolCallError("nw_interface_get_name", "10.14", _nw_interface_get_nameErr)
	}
	return _nw_interface_get_name(interface_), nil
}

// Nw_interface_get_name accesses the name of the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_name(_:)
func Nw_interface_get_name(interface_ Nw_interface_t) *byte {
	result, callErr := tryNw_interface_get_name(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_interface_get_type func(interface_ Nw_interface_t) NwInterfaceType
var _nw_interface_get_typeErr error

func tryNw_interface_get_type(interface_ Nw_interface_t) (NwInterfaceType, error) {
	if _nw_interface_get_type == nil {
		return *new(NwInterfaceType), symbolCallError("nw_interface_get_type", "10.14", _nw_interface_get_typeErr)
	}
	return _nw_interface_get_type(interface_), nil
}

// Nw_interface_get_type accesses the type of the interface, such as Wi-Fi or Loopback.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_type(_:)
func Nw_interface_get_type(interface_ Nw_interface_t) NwInterfaceType {
	result, callErr := tryNw_interface_get_type(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ip_create_metadata func() Nw_protocol_metadata_t
var _nw_ip_create_metadataErr error

func tryNw_ip_create_metadata() (Nw_protocol_metadata_t, error) {
	if _nw_ip_create_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_ip_create_metadata", "10.14", _nw_ip_create_metadataErr)
	}
	return _nw_ip_create_metadata(), nil
}

// Nw_ip_create_metadata initializes an IP packet configuration with default settings.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_create_metadata()
func Nw_ip_create_metadata() Nw_protocol_metadata_t {
	result, callErr := tryNw_ip_create_metadata()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ip_metadata_get_ecn_flag func(metadata Nw_protocol_metadata_t) NwIpEcnFlag
var _nw_ip_metadata_get_ecn_flagErr error

func tryNw_ip_metadata_get_ecn_flag(metadata Nw_protocol_metadata_t) (NwIpEcnFlag, error) {
	if _nw_ip_metadata_get_ecn_flag == nil {
		return *new(NwIpEcnFlag), symbolCallError("nw_ip_metadata_get_ecn_flag", "10.14", _nw_ip_metadata_get_ecn_flagErr)
	}
	return _nw_ip_metadata_get_ecn_flag(metadata), nil
}

// Nw_ip_metadata_get_ecn_flag checks the Explicit Congestion Notification flag value received on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_ecn_flag(_:)
func Nw_ip_metadata_get_ecn_flag(metadata Nw_protocol_metadata_t) NwIpEcnFlag {
	result, callErr := tryNw_ip_metadata_get_ecn_flag(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ip_metadata_get_receive_time func(metadata Nw_protocol_metadata_t) uint64
var _nw_ip_metadata_get_receive_timeErr error

func tryNw_ip_metadata_get_receive_time(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_ip_metadata_get_receive_time == nil {
		return 0, symbolCallError("nw_ip_metadata_get_receive_time", "10.14", _nw_ip_metadata_get_receive_timeErr)
	}
	return _nw_ip_metadata_get_receive_time(metadata), nil
}

// Nw_ip_metadata_get_receive_time access the time at which a packet was received, in nanoseconds, based on `CLOCK_MONOTONIC_RAW`.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_receive_time(_:)
func Nw_ip_metadata_get_receive_time(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_ip_metadata_get_receive_time(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ip_metadata_get_service_class func(metadata Nw_protocol_metadata_t) NwServiceClass
var _nw_ip_metadata_get_service_classErr error

func tryNw_ip_metadata_get_service_class(metadata Nw_protocol_metadata_t) (NwServiceClass, error) {
	if _nw_ip_metadata_get_service_class == nil {
		return *new(NwServiceClass), symbolCallError("nw_ip_metadata_get_service_class", "10.14", _nw_ip_metadata_get_service_classErr)
	}
	return _nw_ip_metadata_get_service_class(metadata), nil
}

// Nw_ip_metadata_get_service_class accesses a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_service_class(_:)
func Nw_ip_metadata_get_service_class(metadata Nw_protocol_metadata_t) NwServiceClass {
	result, callErr := tryNw_ip_metadata_get_service_class(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ip_metadata_set_ecn_flag func(metadata Nw_protocol_metadata_t, ecn_flag NwIpEcnFlag)
var _nw_ip_metadata_set_ecn_flagErr error

func tryNw_ip_metadata_set_ecn_flag(metadata Nw_protocol_metadata_t, ecn_flag NwIpEcnFlag) error {
	if _nw_ip_metadata_set_ecn_flag == nil {
		return symbolCallError("nw_ip_metadata_set_ecn_flag", "10.14", _nw_ip_metadata_set_ecn_flagErr)
	}
	_nw_ip_metadata_set_ecn_flag(metadata, ecn_flag)
	return nil
}

// Nw_ip_metadata_set_ecn_flag sets a specific Explicit Congestion Notification flag value to set on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_ecn_flag(_:_:)
func Nw_ip_metadata_set_ecn_flag(metadata Nw_protocol_metadata_t, ecn_flag NwIpEcnFlag) {
	if callErr := tryNw_ip_metadata_set_ecn_flag(metadata, ecn_flag); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_metadata_set_service_class func(metadata Nw_protocol_metadata_t, service_class NwServiceClass)
var _nw_ip_metadata_set_service_classErr error

func tryNw_ip_metadata_set_service_class(metadata Nw_protocol_metadata_t, service_class NwServiceClass) error {
	if _nw_ip_metadata_set_service_class == nil {
		return symbolCallError("nw_ip_metadata_set_service_class", "10.14", _nw_ip_metadata_set_service_classErr)
	}
	_nw_ip_metadata_set_service_class(metadata, service_class)
	return nil
}

// Nw_ip_metadata_set_service_class sets a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_service_class(_:_:)
func Nw_ip_metadata_set_service_class(metadata Nw_protocol_metadata_t, service_class NwServiceClass) {
	if callErr := tryNw_ip_metadata_set_service_class(metadata, service_class); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_calculate_receive_time func(options Nw_protocol_options_t, calculate_receive_time bool)
var _nw_ip_options_set_calculate_receive_timeErr error

func tryNw_ip_options_set_calculate_receive_time(options Nw_protocol_options_t, calculate_receive_time bool) error {
	if _nw_ip_options_set_calculate_receive_time == nil {
		return symbolCallError("nw_ip_options_set_calculate_receive_time", "10.14", _nw_ip_options_set_calculate_receive_timeErr)
	}
	_nw_ip_options_set_calculate_receive_time(options, calculate_receive_time)
	return nil
}

// Nw_ip_options_set_calculate_receive_time configures a connection to deliver receive timestamps for IP packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_calculate_receive_time(_:_:)
func Nw_ip_options_set_calculate_receive_time(options Nw_protocol_options_t, calculate_receive_time bool) {
	if callErr := tryNw_ip_options_set_calculate_receive_time(options, calculate_receive_time); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_disable_fragmentation func(options Nw_protocol_options_t, disable_fragmentation bool)
var _nw_ip_options_set_disable_fragmentationErr error

func tryNw_ip_options_set_disable_fragmentation(options Nw_protocol_options_t, disable_fragmentation bool) error {
	if _nw_ip_options_set_disable_fragmentation == nil {
		return symbolCallError("nw_ip_options_set_disable_fragmentation", "10.14", _nw_ip_options_set_disable_fragmentationErr)
	}
	_nw_ip_options_set_disable_fragmentation(options, disable_fragmentation)
	return nil
}

// Nw_ip_options_set_disable_fragmentation configures a connection to disable fragmentation on outbound packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_fragmentation(_:_:)
func Nw_ip_options_set_disable_fragmentation(options Nw_protocol_options_t, disable_fragmentation bool) {
	if callErr := tryNw_ip_options_set_disable_fragmentation(options, disable_fragmentation); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_disable_multicast_loopback func(options Nw_protocol_options_t, disable_multicast_loopback bool)
var _nw_ip_options_set_disable_multicast_loopbackErr error

func tryNw_ip_options_set_disable_multicast_loopback(options Nw_protocol_options_t, disable_multicast_loopback bool) error {
	if _nw_ip_options_set_disable_multicast_loopback == nil {
		return symbolCallError("nw_ip_options_set_disable_multicast_loopback", "11.0", _nw_ip_options_set_disable_multicast_loopbackErr)
	}
	_nw_ip_options_set_disable_multicast_loopback(options, disable_multicast_loopback)
	return nil
}

// Nw_ip_options_set_disable_multicast_loopback.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_multicast_loopback(_:_:)
func Nw_ip_options_set_disable_multicast_loopback(options Nw_protocol_options_t, disable_multicast_loopback bool) {
	if callErr := tryNw_ip_options_set_disable_multicast_loopback(options, disable_multicast_loopback); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_hop_limit func(options Nw_protocol_options_t, hop_limit uint8)
var _nw_ip_options_set_hop_limitErr error

func tryNw_ip_options_set_hop_limit(options Nw_protocol_options_t, hop_limit uint8) error {
	if _nw_ip_options_set_hop_limit == nil {
		return symbolCallError("nw_ip_options_set_hop_limit", "10.14", _nw_ip_options_set_hop_limitErr)
	}
	_nw_ip_options_set_hop_limit(options, hop_limit)
	return nil
}

// Nw_ip_options_set_hop_limit configures the default hop limit for packets generated by a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_hop_limit(_:_:)
func Nw_ip_options_set_hop_limit(options Nw_protocol_options_t, hop_limit uint8) {
	if callErr := tryNw_ip_options_set_hop_limit(options, hop_limit); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_local_address_preference func(options Nw_protocol_options_t, preference NwIpLocalAddressPreference)
var _nw_ip_options_set_local_address_preferenceErr error

func tryNw_ip_options_set_local_address_preference(options Nw_protocol_options_t, preference NwIpLocalAddressPreference) error {
	if _nw_ip_options_set_local_address_preference == nil {
		return symbolCallError("nw_ip_options_set_local_address_preference", "10.15", _nw_ip_options_set_local_address_preferenceErr)
	}
	_nw_ip_options_set_local_address_preference(options, preference)
	return nil
}

// Nw_ip_options_set_local_address_preference configures a connection to prefer certain types of local addresses, such as temporary or stable.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_local_address_preference(_:_:)
func Nw_ip_options_set_local_address_preference(options Nw_protocol_options_t, preference NwIpLocalAddressPreference) {
	if callErr := tryNw_ip_options_set_local_address_preference(options, preference); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_use_minimum_mtu func(options Nw_protocol_options_t, use_minimum_mtu bool)
var _nw_ip_options_set_use_minimum_mtuErr error

func tryNw_ip_options_set_use_minimum_mtu(options Nw_protocol_options_t, use_minimum_mtu bool) error {
	if _nw_ip_options_set_use_minimum_mtu == nil {
		return symbolCallError("nw_ip_options_set_use_minimum_mtu", "10.14", _nw_ip_options_set_use_minimum_mtuErr)
	}
	_nw_ip_options_set_use_minimum_mtu(options, use_minimum_mtu)
	return nil
}

// Nw_ip_options_set_use_minimum_mtu configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_use_minimum_mtu(_:_:)
func Nw_ip_options_set_use_minimum_mtu(options Nw_protocol_options_t, use_minimum_mtu bool) {
	if callErr := tryNw_ip_options_set_use_minimum_mtu(options, use_minimum_mtu); callErr != nil {
		panic(callErr)
	}
}

var _nw_ip_options_set_version func(options Nw_protocol_options_t, version NwIpVersion)
var _nw_ip_options_set_versionErr error

func tryNw_ip_options_set_version(options Nw_protocol_options_t, version NwIpVersion) error {
	if _nw_ip_options_set_version == nil {
		return symbolCallError("nw_ip_options_set_version", "10.14", _nw_ip_options_set_versionErr)
	}
	_nw_ip_options_set_version(options, version)
	return nil
}

// Nw_ip_options_set_version sets a required IP version to disable all other versions for a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_version(_:_:)
func Nw_ip_options_set_version(options Nw_protocol_options_t, version NwIpVersion) {
	if callErr := tryNw_ip_options_set_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_cancel func(listener Nw_listener_t)
var _nw_listener_cancelErr error

func tryNw_listener_cancel(listener Nw_listener_t) error {
	if _nw_listener_cancel == nil {
		return symbolCallError("nw_listener_cancel", "10.14", _nw_listener_cancelErr)
	}
	_nw_listener_cancel(listener)
	return nil
}

// Nw_listener_cancel stops listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_cancel(_:)
func Nw_listener_cancel(listener Nw_listener_t) {
	if callErr := tryNw_listener_cancel(listener); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_create func(parameters Nw_parameters_t) Nw_listener_t
var _nw_listener_createErr error

func tryNw_listener_create(parameters Nw_parameters_t) (Nw_listener_t, error) {
	if _nw_listener_create == nil {
		return *new(Nw_listener_t), symbolCallError("nw_listener_create", "10.14", _nw_listener_createErr)
	}
	return _nw_listener_create(parameters), nil
}

// Nw_listener_create initializes a network listener, which will select a random port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create(_:)
func Nw_listener_create(parameters Nw_parameters_t) Nw_listener_t {
	result, callErr := tryNw_listener_create(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_create_with_connection func(connection Nw_connection_t, parameters Nw_parameters_t) Nw_listener_t
var _nw_listener_create_with_connectionErr error

func tryNw_listener_create_with_connection(connection Nw_connection_t, parameters Nw_parameters_t) (Nw_listener_t, error) {
	if _nw_listener_create_with_connection == nil {
		return *new(Nw_listener_t), symbolCallError("nw_listener_create_with_connection", "10.14", _nw_listener_create_with_connectionErr)
	}
	return _nw_listener_create_with_connection(connection, parameters), nil
}

// Nw_listener_create_with_connection initializes a network listener to receive new streams on a multiplexed connection.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_connection(_:_:)
func Nw_listener_create_with_connection(connection Nw_connection_t, parameters Nw_parameters_t) Nw_listener_t {
	result, callErr := tryNw_listener_create_with_connection(connection, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_create_with_launchd_key func(parameters Nw_parameters_t, launchd_key string) Nw_listener_t
var _nw_listener_create_with_launchd_keyErr error

func tryNw_listener_create_with_launchd_key(parameters Nw_parameters_t, launchd_key string) (Nw_listener_t, error) {
	if _nw_listener_create_with_launchd_key == nil {
		return *new(Nw_listener_t), symbolCallError("nw_listener_create_with_launchd_key", "10.14", _nw_listener_create_with_launchd_keyErr)
	}
	return _nw_listener_create_with_launchd_key(parameters, launchd_key), nil
}

// Nw_listener_create_with_launchd_key.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_launchd_key(_:_:)
func Nw_listener_create_with_launchd_key(parameters Nw_parameters_t, launchd_key string) Nw_listener_t {
	result, callErr := tryNw_listener_create_with_launchd_key(parameters, launchd_key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_create_with_port func(port string, parameters Nw_parameters_t) Nw_listener_t
var _nw_listener_create_with_portErr error

func tryNw_listener_create_with_port(port string, parameters Nw_parameters_t) (Nw_listener_t, error) {
	if _nw_listener_create_with_port == nil {
		return *new(Nw_listener_t), symbolCallError("nw_listener_create_with_port", "10.14", _nw_listener_create_with_portErr)
	}
	return _nw_listener_create_with_port(port, parameters), nil
}

// Nw_listener_create_with_port initializes a network listener with a specified local port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_port(_:_:)
func Nw_listener_create_with_port(port string, parameters Nw_parameters_t) Nw_listener_t {
	result, callErr := tryNw_listener_create_with_port(port, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_get_new_connection_limit func(listener Nw_listener_t) uint32
var _nw_listener_get_new_connection_limitErr error

func tryNw_listener_get_new_connection_limit(listener Nw_listener_t) (uint32, error) {
	if _nw_listener_get_new_connection_limit == nil {
		return 0, symbolCallError("nw_listener_get_new_connection_limit", "10.15", _nw_listener_get_new_connection_limitErr)
	}
	return _nw_listener_get_new_connection_limit(listener), nil
}

// Nw_listener_get_new_connection_limit checks the remaining number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_new_connection_limit(_:)
func Nw_listener_get_new_connection_limit(listener Nw_listener_t) uint32 {
	result, callErr := tryNw_listener_get_new_connection_limit(listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_get_port func(listener Nw_listener_t) uint16
var _nw_listener_get_portErr error

func tryNw_listener_get_port(listener Nw_listener_t) (uint16, error) {
	if _nw_listener_get_port == nil {
		return 0, symbolCallError("nw_listener_get_port", "10.14", _nw_listener_get_portErr)
	}
	return _nw_listener_get_port(listener), nil
}

// Nw_listener_get_port the port on which the listener can accept connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_port(_:)
func Nw_listener_get_port(listener Nw_listener_t) uint16 {
	result, callErr := tryNw_listener_get_port(listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_listener_set_advertise_descriptor func(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t)
var _nw_listener_set_advertise_descriptorErr error

func tryNw_listener_set_advertise_descriptor(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t) error {
	if _nw_listener_set_advertise_descriptor == nil {
		return symbolCallError("nw_listener_set_advertise_descriptor", "10.14", _nw_listener_set_advertise_descriptorErr)
	}
	_nw_listener_set_advertise_descriptor(listener, advertise_descriptor)
	return nil
}

// Nw_listener_set_advertise_descriptor sets a Bonjour service that advertises the listener on the local network.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertise_descriptor(_:_:)
func Nw_listener_set_advertise_descriptor(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t) {
	if callErr := tryNw_listener_set_advertise_descriptor(listener, advertise_descriptor); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_advertised_endpoint_changed_handler func(listener Nw_listener_t, handler unsafe.Pointer)
var _nw_listener_set_advertised_endpoint_changed_handlerErr error

func tryNw_listener_set_advertised_endpoint_changed_handler(listener Nw_listener_t, handler Nw_listener_advertised_endpoint_changed_handler_t) error {
	if _nw_listener_set_advertised_endpoint_changed_handler == nil {
		return symbolCallError("nw_listener_set_advertised_endpoint_changed_handler", "10.14", _nw_listener_set_advertised_endpoint_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 bool) { handler(arg0, arg1) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_advertised_endpoint_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_listener_set_advertised_endpoint_changed_handler(listener, _block0)
	return nil
}

// Nw_listener_set_advertised_endpoint_changed_handler sets a handler that receives updates for the service endpoint being advertised.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertised_endpoint_changed_handler(_:_:)
func Nw_listener_set_advertised_endpoint_changed_handler(listener Nw_listener_t, handler Nw_listener_advertised_endpoint_changed_handler_t) {
	if callErr := tryNw_listener_set_advertised_endpoint_changed_handler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_new_connection_group_handler func(listener Nw_listener_t, handler unsafe.Pointer)
var _nw_listener_set_new_connection_group_handlerErr error

func tryNw_listener_set_new_connection_group_handler(listener Nw_listener_t, handler Nw_listener_new_connection_group_handler_t) error {
	if _nw_listener_set_new_connection_group_handler == nil {
		return symbolCallError("nw_listener_set_new_connection_group_handler", "12.0", _nw_listener_set_new_connection_group_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { handler(arg0) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_new_connection_group_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_listener_set_new_connection_group_handler(listener, _block0)
	return nil
}

// Nw_listener_set_new_connection_group_handler.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_group_handler(_:_:)
func Nw_listener_set_new_connection_group_handler(listener Nw_listener_t, handler Nw_listener_new_connection_group_handler_t) {
	if callErr := tryNw_listener_set_new_connection_group_handler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_new_connection_handler func(listener Nw_listener_t, handler unsafe.Pointer)
var _nw_listener_set_new_connection_handlerErr error

func tryNw_listener_set_new_connection_handler(listener Nw_listener_t, handler Nw_listener_new_connection_handler_t) error {
	if _nw_listener_set_new_connection_handler == nil {
		return symbolCallError("nw_listener_set_new_connection_handler", "10.14", _nw_listener_set_new_connection_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { handler(arg0) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_new_connection_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_listener_set_new_connection_handler(listener, _block0)
	return nil
}

// Nw_listener_set_new_connection_handler sets a handler that receives inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_handler(_:_:)
func Nw_listener_set_new_connection_handler(listener Nw_listener_t, handler Nw_listener_new_connection_handler_t) {
	if callErr := tryNw_listener_set_new_connection_handler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_new_connection_limit func(listener Nw_listener_t, new_connection_limit uint32)
var _nw_listener_set_new_connection_limitErr error

func tryNw_listener_set_new_connection_limit(listener Nw_listener_t, new_connection_limit uint32) error {
	if _nw_listener_set_new_connection_limit == nil {
		return symbolCallError("nw_listener_set_new_connection_limit", "10.15", _nw_listener_set_new_connection_limitErr)
	}
	_nw_listener_set_new_connection_limit(listener, new_connection_limit)
	return nil
}

// Nw_listener_set_new_connection_limit resets the number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_limit(_:_:)
func Nw_listener_set_new_connection_limit(listener Nw_listener_t, new_connection_limit uint32) {
	if callErr := tryNw_listener_set_new_connection_limit(listener, new_connection_limit); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_queue func(listener Nw_listener_t, queue uintptr)
var _nw_listener_set_queueErr error

func tryNw_listener_set_queue(listener Nw_listener_t, queue dispatch.Queue) error {
	if _nw_listener_set_queue == nil {
		return symbolCallError("nw_listener_set_queue", "10.14", _nw_listener_set_queueErr)
	}
	_nw_listener_set_queue(listener, uintptr(queue.Handle()))
	return nil
}

// Nw_listener_set_queue sets the queue on which all listener events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_queue(_:_:)
func Nw_listener_set_queue(listener Nw_listener_t, queue dispatch.Queue) {
	if callErr := tryNw_listener_set_queue(listener, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_set_state_changed_handler func(listener Nw_listener_t, handler unsafe.Pointer)
var _nw_listener_set_state_changed_handlerErr error

func tryNw_listener_set_state_changed_handler(listener Nw_listener_t, handler Nw_listener_state_changed_handler_t) error {
	if _nw_listener_set_state_changed_handler == nil {
		return symbolCallError("nw_listener_set_state_changed_handler", "10.14", _nw_listener_set_state_changed_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwListenerState, arg1 objectivec.Object) { handler(arg0, arg1) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_listener_set_state_changed_handler(listener, _block0)
	return nil
}

// Nw_listener_set_state_changed_handler sets a handler to receive listener state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_state_changed_handler(_:_:)
func Nw_listener_set_state_changed_handler(listener Nw_listener_t, handler Nw_listener_state_changed_handler_t) {
	if callErr := tryNw_listener_set_state_changed_handler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_listener_start func(listener Nw_listener_t)
var _nw_listener_startErr error

func tryNw_listener_start(listener Nw_listener_t) error {
	if _nw_listener_start == nil {
		return symbolCallError("nw_listener_start", "10.14", _nw_listener_startErr)
	}
	_nw_listener_start(listener)
	return nil
}

// Nw_listener_start registers for listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_start(_:)
func Nw_listener_start(listener Nw_listener_t) {
	if callErr := tryNw_listener_start(listener); callErr != nil {
		panic(callErr)
	}
}

var _nw_multicast_group_descriptor_get_disable_unicast_traffic func(multicast_descriptor Nw_group_descriptor_t) bool
var _nw_multicast_group_descriptor_get_disable_unicast_trafficErr error

func tryNw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t) (bool, error) {
	if _nw_multicast_group_descriptor_get_disable_unicast_traffic == nil {
		return false, symbolCallError("nw_multicast_group_descriptor_get_disable_unicast_traffic", "11.0", _nw_multicast_group_descriptor_get_disable_unicast_trafficErr)
	}
	return _nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor), nil
}

// Nw_multicast_group_descriptor_get_disable_unicast_traffic checks a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_get_disable_unicast_traffic(_:)
func Nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t) bool {
	result, callErr := tryNw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_multicast_group_descriptor_set_disable_unicast_traffic func(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool)
var _nw_multicast_group_descriptor_set_disable_unicast_trafficErr error

func tryNw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool) error {
	if _nw_multicast_group_descriptor_set_disable_unicast_traffic == nil {
		return symbolCallError("nw_multicast_group_descriptor_set_disable_unicast_traffic", "11.0", _nw_multicast_group_descriptor_set_disable_unicast_trafficErr)
	}
	_nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor, disable_unicast_traffic)
	return nil
}

// Nw_multicast_group_descriptor_set_disable_unicast_traffic sets a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_disable_unicast_traffic(_:_:)
func Nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool) {
	if callErr := tryNw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor, disable_unicast_traffic); callErr != nil {
		panic(callErr)
	}
}

var _nw_multicast_group_descriptor_set_specific_source func(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t)
var _nw_multicast_group_descriptor_set_specific_sourceErr error

func tryNw_multicast_group_descriptor_set_specific_source(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t) error {
	if _nw_multicast_group_descriptor_set_specific_source == nil {
		return symbolCallError("nw_multicast_group_descriptor_set_specific_source", "11.0", _nw_multicast_group_descriptor_set_specific_sourceErr)
	}
	_nw_multicast_group_descriptor_set_specific_source(multicast_descriptor, source)
	return nil
}

// Nw_multicast_group_descriptor_set_specific_source sets an optional address endpoint used to filter received multicast packets.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_specific_source(_:_:)
func Nw_multicast_group_descriptor_set_specific_source(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t) {
	if callErr := tryNw_multicast_group_descriptor_set_specific_source(multicast_descriptor, source); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_clear_prohibited_interface_types func(parameters Nw_parameters_t)
var _nw_parameters_clear_prohibited_interface_typesErr error

func tryNw_parameters_clear_prohibited_interface_types(parameters Nw_parameters_t) error {
	if _nw_parameters_clear_prohibited_interface_types == nil {
		return symbolCallError("nw_parameters_clear_prohibited_interface_types", "10.14", _nw_parameters_clear_prohibited_interface_typesErr)
	}
	_nw_parameters_clear_prohibited_interface_types(parameters)
	return nil
}

// Nw_parameters_clear_prohibited_interface_types removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interface_types(_:)
func Nw_parameters_clear_prohibited_interface_types(parameters Nw_parameters_t) {
	if callErr := tryNw_parameters_clear_prohibited_interface_types(parameters); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_clear_prohibited_interfaces func(parameters Nw_parameters_t)
var _nw_parameters_clear_prohibited_interfacesErr error

func tryNw_parameters_clear_prohibited_interfaces(parameters Nw_parameters_t) error {
	if _nw_parameters_clear_prohibited_interfaces == nil {
		return symbolCallError("nw_parameters_clear_prohibited_interfaces", "10.14", _nw_parameters_clear_prohibited_interfacesErr)
	}
	_nw_parameters_clear_prohibited_interfaces(parameters)
	return nil
}

// Nw_parameters_clear_prohibited_interfaces removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interfaces(_:)
func Nw_parameters_clear_prohibited_interfaces(parameters Nw_parameters_t) {
	if callErr := tryNw_parameters_clear_prohibited_interfaces(parameters); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_copy func(parameters Nw_parameters_t) Nw_parameters_t
var _nw_parameters_copyErr error

func tryNw_parameters_copy(parameters Nw_parameters_t) (Nw_parameters_t, error) {
	if _nw_parameters_copy == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_copy", "10.14", _nw_parameters_copyErr)
	}
	return _nw_parameters_copy(parameters), nil
}

// Nw_parameters_copy peforms a deep copy of a parameters object.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy(_:)
func Nw_parameters_copy(parameters Nw_parameters_t) Nw_parameters_t {
	result, callErr := tryNw_parameters_copy(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_copy_default_protocol_stack func(parameters Nw_parameters_t) Nw_protocol_stack_t
var _nw_parameters_copy_default_protocol_stackErr error

func tryNw_parameters_copy_default_protocol_stack(parameters Nw_parameters_t) (Nw_protocol_stack_t, error) {
	if _nw_parameters_copy_default_protocol_stack == nil {
		return *new(Nw_protocol_stack_t), symbolCallError("nw_parameters_copy_default_protocol_stack", "10.14", _nw_parameters_copy_default_protocol_stackErr)
	}
	return _nw_parameters_copy_default_protocol_stack(parameters), nil
}

// Nw_parameters_copy_default_protocol_stack accesses the protocol stack used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_default_protocol_stack(_:)
func Nw_parameters_copy_default_protocol_stack(parameters Nw_parameters_t) Nw_protocol_stack_t {
	result, callErr := tryNw_parameters_copy_default_protocol_stack(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_copy_local_endpoint func(parameters Nw_parameters_t) Nw_endpoint_t
var _nw_parameters_copy_local_endpointErr error

func tryNw_parameters_copy_local_endpoint(parameters Nw_parameters_t) (Nw_endpoint_t, error) {
	if _nw_parameters_copy_local_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_parameters_copy_local_endpoint", "10.14", _nw_parameters_copy_local_endpointErr)
	}
	return _nw_parameters_copy_local_endpoint(parameters), nil
}

// Nw_parameters_copy_local_endpoint accesses the local IP address and port used for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_local_endpoint(_:)
func Nw_parameters_copy_local_endpoint(parameters Nw_parameters_t) Nw_endpoint_t {
	result, callErr := tryNw_parameters_copy_local_endpoint(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_copy_required_interface func(parameters Nw_parameters_t) Nw_interface_t
var _nw_parameters_copy_required_interfaceErr error

func tryNw_parameters_copy_required_interface(parameters Nw_parameters_t) (Nw_interface_t, error) {
	if _nw_parameters_copy_required_interface == nil {
		return *new(Nw_interface_t), symbolCallError("nw_parameters_copy_required_interface", "10.14", _nw_parameters_copy_required_interfaceErr)
	}
	return _nw_parameters_copy_required_interface(parameters), nil
}

// Nw_parameters_copy_required_interface accesses the interface required on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_required_interface(_:)
func Nw_parameters_copy_required_interface(parameters Nw_parameters_t) Nw_interface_t {
	result, callErr := tryNw_parameters_copy_required_interface(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create func() Nw_parameters_t
var _nw_parameters_createErr error

func tryNw_parameters_create() (Nw_parameters_t, error) {
	if _nw_parameters_create == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create", "10.14", _nw_parameters_createErr)
	}
	return _nw_parameters_create(), nil
}

// Nw_parameters_create initializes parameters for connections, listeners, and browsers with no protocols specified.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create()
func Nw_parameters_create() Nw_parameters_t {
	result, callErr := tryNw_parameters_create()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create_application_service func() Nw_parameters_t
var _nw_parameters_create_application_serviceErr error

func tryNw_parameters_create_application_service() (Nw_parameters_t, error) {
	if _nw_parameters_create_application_service == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create_application_service", "13.0", _nw_parameters_create_application_serviceErr)
	}
	return _nw_parameters_create_application_service(), nil
}

// Nw_parameters_create_application_service.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_application_service()
func Nw_parameters_create_application_service() Nw_parameters_t {
	result, callErr := tryNw_parameters_create_application_service()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create_custom_ip func(custom_ip_protocol_number uint8, configure_ip unsafe.Pointer) Nw_parameters_t
var _nw_parameters_create_custom_ipErr error

func tryNw_parameters_create_custom_ip(custom_ip_protocol_number uint8, configure_ip Nw_parameters_configure_protocol_block_t) (Nw_parameters_t, error) {
	if _nw_parameters_create_custom_ip == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create_custom_ip", "10.15", _nw_parameters_create_custom_ipErr)
	}
	var _block0 unsafe.Pointer
	if configure_ip == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(Nw_parameters_t), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_ip(arg0) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	return _nw_parameters_create_custom_ip(custom_ip_protocol_number, _block0), nil
}

// Nw_parameters_create_custom_ip initializes parameters for connections and listeners using a custom IP protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_custom_ip(_:_:)
func Nw_parameters_create_custom_ip(custom_ip_protocol_number uint8, configure_ip Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	result, callErr := tryNw_parameters_create_custom_ip(custom_ip_protocol_number, configure_ip)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create_quic func(configure_quic unsafe.Pointer) Nw_parameters_t
var _nw_parameters_create_quicErr error

func tryNw_parameters_create_quic(configure_quic Nw_parameters_configure_protocol_block_t) (Nw_parameters_t, error) {
	if _nw_parameters_create_quic == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create_quic", "12.0", _nw_parameters_create_quicErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_quic(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_parameters_create_quic(_block0), nil
}

// Nw_parameters_create_quic initializes parameters for QUIC connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_quic(_:)
func Nw_parameters_create_quic(configure_quic Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	result, callErr := tryNw_parameters_create_quic(configure_quic)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create_secure_tcp func(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) Nw_parameters_t
var _nw_parameters_create_secure_tcpErr error

func tryNw_parameters_create_secure_tcp(configure_tls Nw_parameters_configure_protocol_block_t, configure_tcp Nw_parameters_configure_protocol_block_t) (Nw_parameters_t, error) {
	if _nw_parameters_create_secure_tcp == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create_secure_tcp", "10.14", _nw_parameters_create_secure_tcpErr)
	}
	var _block0 unsafe.Pointer
	if configure_tls == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(Nw_parameters_t), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_tls(arg0) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	var _block1 unsafe.Pointer
	if configure_tcp == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(Nw_parameters_t), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_tcp(arg0) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return _nw_parameters_create_secure_tcp(_block0, _block1), nil
}

// Nw_parameters_create_secure_tcp initializes parameters for TLS or TCP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_tcp(_:_:)
func Nw_parameters_create_secure_tcp(configure_tls Nw_parameters_configure_protocol_block_t, configure_tcp Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	result, callErr := tryNw_parameters_create_secure_tcp(configure_tls, configure_tcp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_create_secure_udp func(configure_dtls unsafe.Pointer, configure_udp unsafe.Pointer) Nw_parameters_t
var _nw_parameters_create_secure_udpErr error

func tryNw_parameters_create_secure_udp(configure_dtls Nw_parameters_configure_protocol_block_t, configure_udp Nw_parameters_configure_protocol_block_t) (Nw_parameters_t, error) {
	if _nw_parameters_create_secure_udp == nil {
		return *new(Nw_parameters_t), symbolCallError("nw_parameters_create_secure_udp", "10.14", _nw_parameters_create_secure_udpErr)
	}
	var _block0 unsafe.Pointer
	if configure_dtls == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(Nw_parameters_t), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_dtls(arg0) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	var _block1 unsafe.Pointer
	if configure_udp == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(Nw_parameters_t), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { configure_udp(arg0) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return _nw_parameters_create_secure_udp(_block0, _block1), nil
}

// Nw_parameters_create_secure_udp initializes parameters for DTLS or UDP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_udp(_:_:)
func Nw_parameters_create_secure_udp(configure_dtls Nw_parameters_configure_protocol_block_t, configure_udp Nw_parameters_configure_protocol_block_t) Nw_parameters_t {
	result, callErr := tryNw_parameters_create_secure_udp(configure_dtls, configure_udp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_allow_ultra_constrained func(parameters Nw_parameters_t) bool
var _nw_parameters_get_allow_ultra_constrainedErr error

func tryNw_parameters_get_allow_ultra_constrained(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_allow_ultra_constrained == nil {
		return false, symbolCallError("nw_parameters_get_allow_ultra_constrained", "26.0", _nw_parameters_get_allow_ultra_constrainedErr)
	}
	return _nw_parameters_get_allow_ultra_constrained(parameters), nil
}

// Nw_parameters_get_allow_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_allow_ultra_constrained(_:)
func Nw_parameters_get_allow_ultra_constrained(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_allow_ultra_constrained(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_attribution func(parameters Nw_parameters_t) Nw_parameters_attribution_t
var _nw_parameters_get_attributionErr error

func tryNw_parameters_get_attribution(parameters Nw_parameters_t) (Nw_parameters_attribution_t, error) {
	if _nw_parameters_get_attribution == nil {
		return *new(Nw_parameters_attribution_t), symbolCallError("nw_parameters_get_attribution", "12.0", _nw_parameters_get_attributionErr)
	}
	return _nw_parameters_get_attribution(parameters), nil
}

// Nw_parameters_get_attribution gets a flag that indicates whether the network request originates from the developer or the user.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_attribution(_:)
func Nw_parameters_get_attribution(parameters Nw_parameters_t) Nw_parameters_attribution_t {
	result, callErr := tryNw_parameters_get_attribution(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_expired_dns_behavior func(parameters Nw_parameters_t) NwParametersExpiredDnsBehavior
var _nw_parameters_get_expired_dns_behaviorErr error

func tryNw_parameters_get_expired_dns_behavior(parameters Nw_parameters_t) (NwParametersExpiredDnsBehavior, error) {
	if _nw_parameters_get_expired_dns_behavior == nil {
		return *new(NwParametersExpiredDnsBehavior), symbolCallError("nw_parameters_get_expired_dns_behavior", "10.14", _nw_parameters_get_expired_dns_behaviorErr)
	}
	return _nw_parameters_get_expired_dns_behavior(parameters), nil
}

// Nw_parameters_get_expired_dns_behavior checks the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_expired_dns_behavior(_:)
func Nw_parameters_get_expired_dns_behavior(parameters Nw_parameters_t) NwParametersExpiredDnsBehavior {
	result, callErr := tryNw_parameters_get_expired_dns_behavior(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_fast_open_enabled func(parameters Nw_parameters_t) bool
var _nw_parameters_get_fast_open_enabledErr error

func tryNw_parameters_get_fast_open_enabled(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_fast_open_enabled == nil {
		return false, symbolCallError("nw_parameters_get_fast_open_enabled", "10.14", _nw_parameters_get_fast_open_enabledErr)
	}
	return _nw_parameters_get_fast_open_enabled(parameters), nil
}

// Nw_parameters_get_fast_open_enabled checks if sending application data with protocol handshakes is enabled.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_fast_open_enabled(_:)
func Nw_parameters_get_fast_open_enabled(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_fast_open_enabled(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_include_peer_to_peer func(parameters Nw_parameters_t) bool
var _nw_parameters_get_include_peer_to_peerErr error

func tryNw_parameters_get_include_peer_to_peer(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_include_peer_to_peer == nil {
		return false, symbolCallError("nw_parameters_get_include_peer_to_peer", "10.14", _nw_parameters_get_include_peer_to_peerErr)
	}
	return _nw_parameters_get_include_peer_to_peer(parameters), nil
}

// Nw_parameters_get_include_peer_to_peer checks whether a connection is allowed to use peer-to-peer link technologies.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_include_peer_to_peer(_:)
func Nw_parameters_get_include_peer_to_peer(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_include_peer_to_peer(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_local_only func(parameters Nw_parameters_t) bool
var _nw_parameters_get_local_onlyErr error

func tryNw_parameters_get_local_only(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_local_only == nil {
		return false, symbolCallError("nw_parameters_get_local_only", "10.14", _nw_parameters_get_local_onlyErr)
	}
	return _nw_parameters_get_local_only(parameters), nil
}

// Nw_parameters_get_local_only checks if a listener is restricted to accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_local_only(_:)
func Nw_parameters_get_local_only(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_local_only(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_multipath_service func(parameters Nw_parameters_t) NwMultipathService
var _nw_parameters_get_multipath_serviceErr error

func tryNw_parameters_get_multipath_service(parameters Nw_parameters_t) (NwMultipathService, error) {
	if _nw_parameters_get_multipath_service == nil {
		return *new(NwMultipathService), symbolCallError("nw_parameters_get_multipath_service", "10.14", _nw_parameters_get_multipath_serviceErr)
	}
	return _nw_parameters_get_multipath_service(parameters), nil
}

// Nw_parameters_get_multipath_service checks if multipath is enabled on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_multipath_service(_:)
func Nw_parameters_get_multipath_service(parameters Nw_parameters_t) NwMultipathService {
	result, callErr := tryNw_parameters_get_multipath_service(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_prefer_no_proxy func(parameters Nw_parameters_t) bool
var _nw_parameters_get_prefer_no_proxyErr error

func tryNw_parameters_get_prefer_no_proxy(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_prefer_no_proxy == nil {
		return false, symbolCallError("nw_parameters_get_prefer_no_proxy", "10.14", _nw_parameters_get_prefer_no_proxyErr)
	}
	return _nw_parameters_get_prefer_no_proxy(parameters), nil
}

// Nw_parameters_get_prefer_no_proxy checks if proxies are ignored by default.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prefer_no_proxy(_:)
func Nw_parameters_get_prefer_no_proxy(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_prefer_no_proxy(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_prohibit_constrained func(parameters Nw_parameters_t) bool
var _nw_parameters_get_prohibit_constrainedErr error

func tryNw_parameters_get_prohibit_constrained(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_prohibit_constrained == nil {
		return false, symbolCallError("nw_parameters_get_prohibit_constrained", "10.15", _nw_parameters_get_prohibit_constrainedErr)
	}
	return _nw_parameters_get_prohibit_constrained(parameters), nil
}

// Nw_parameters_get_prohibit_constrained checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_constrained(_:)
func Nw_parameters_get_prohibit_constrained(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_prohibit_constrained(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_prohibit_expensive func(parameters Nw_parameters_t) bool
var _nw_parameters_get_prohibit_expensiveErr error

func tryNw_parameters_get_prohibit_expensive(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_prohibit_expensive == nil {
		return false, symbolCallError("nw_parameters_get_prohibit_expensive", "10.14", _nw_parameters_get_prohibit_expensiveErr)
	}
	return _nw_parameters_get_prohibit_expensive(parameters), nil
}

// Nw_parameters_get_prohibit_expensive checks if connections, listeners, and browsers are prevented from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_expensive(_:)
func Nw_parameters_get_prohibit_expensive(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_prohibit_expensive(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_required_interface_type func(parameters Nw_parameters_t) NwInterfaceType
var _nw_parameters_get_required_interface_typeErr error

func tryNw_parameters_get_required_interface_type(parameters Nw_parameters_t) (NwInterfaceType, error) {
	if _nw_parameters_get_required_interface_type == nil {
		return *new(NwInterfaceType), symbolCallError("nw_parameters_get_required_interface_type", "10.14", _nw_parameters_get_required_interface_typeErr)
	}
	return _nw_parameters_get_required_interface_type(parameters), nil
}

// Nw_parameters_get_required_interface_type accesses the interface type required on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_required_interface_type(_:)
func Nw_parameters_get_required_interface_type(parameters Nw_parameters_t) NwInterfaceType {
	result, callErr := tryNw_parameters_get_required_interface_type(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_reuse_local_address func(parameters Nw_parameters_t) bool
var _nw_parameters_get_reuse_local_addressErr error

func tryNw_parameters_get_reuse_local_address(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_get_reuse_local_address == nil {
		return false, symbolCallError("nw_parameters_get_reuse_local_address", "10.14", _nw_parameters_get_reuse_local_addressErr)
	}
	return _nw_parameters_get_reuse_local_address(parameters), nil
}

// Nw_parameters_get_reuse_local_address checks whether a connection allows reusing local addresses and ports.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_reuse_local_address(_:)
func Nw_parameters_get_reuse_local_address(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_get_reuse_local_address(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_get_service_class func(parameters Nw_parameters_t) NwServiceClass
var _nw_parameters_get_service_classErr error

func tryNw_parameters_get_service_class(parameters Nw_parameters_t) (NwServiceClass, error) {
	if _nw_parameters_get_service_class == nil {
		return *new(NwServiceClass), symbolCallError("nw_parameters_get_service_class", "10.14", _nw_parameters_get_service_classErr)
	}
	return _nw_parameters_get_service_class(parameters), nil
}

// Nw_parameters_get_service_class checks the level of service quality used for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_service_class(_:)
func Nw_parameters_get_service_class(parameters Nw_parameters_t) NwServiceClass {
	result, callErr := tryNw_parameters_get_service_class(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_iterate_prohibited_interface_types func(parameters Nw_parameters_t, iterate_block unsafe.Pointer)
var _nw_parameters_iterate_prohibited_interface_typesErr error

func tryNw_parameters_iterate_prohibited_interface_types(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interface_types_block_t) error {
	if _nw_parameters_iterate_prohibited_interface_types == nil {
		return symbolCallError("nw_parameters_iterate_prohibited_interface_types", "10.14", _nw_parameters_iterate_prohibited_interface_typesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 NwInterfaceType) bool { return iterate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_parameters_iterate_prohibited_interface_types(parameters, _block0)
	return nil
}

// Nw_parameters_iterate_prohibited_interface_types examines the list of prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interface_types(_:_:)
func Nw_parameters_iterate_prohibited_interface_types(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interface_types_block_t) {
	if callErr := tryNw_parameters_iterate_prohibited_interface_types(parameters, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_iterate_prohibited_interfaces func(parameters Nw_parameters_t, iterate_block unsafe.Pointer)
var _nw_parameters_iterate_prohibited_interfacesErr error

func tryNw_parameters_iterate_prohibited_interfaces(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interfaces_block_t) error {
	if _nw_parameters_iterate_prohibited_interfaces == nil {
		return symbolCallError("nw_parameters_iterate_prohibited_interfaces", "10.14", _nw_parameters_iterate_prohibited_interfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return iterate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_parameters_iterate_prohibited_interfaces(parameters, _block0)
	return nil
}

// Nw_parameters_iterate_prohibited_interfaces examines the list of prohibited interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interfaces(_:_:)
func Nw_parameters_iterate_prohibited_interfaces(parameters Nw_parameters_t, iterate_block Nw_parameters_iterate_interfaces_block_t) {
	if callErr := tryNw_parameters_iterate_prohibited_interfaces(parameters, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_prohibit_interface func(parameters Nw_parameters_t, interface_ Nw_interface_t)
var _nw_parameters_prohibit_interfaceErr error

func tryNw_parameters_prohibit_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) error {
	if _nw_parameters_prohibit_interface == nil {
		return symbolCallError("nw_parameters_prohibit_interface", "10.14", _nw_parameters_prohibit_interfaceErr)
	}
	_nw_parameters_prohibit_interface(parameters, interface_)
	return nil
}

// Nw_parameters_prohibit_interface prevents connections and listeners from using a specific interface.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface(_:_:)
func Nw_parameters_prohibit_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	if callErr := tryNw_parameters_prohibit_interface(parameters, interface_); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_prohibit_interface_type func(parameters Nw_parameters_t, interface_type NwInterfaceType)
var _nw_parameters_prohibit_interface_typeErr error

func tryNw_parameters_prohibit_interface_type(parameters Nw_parameters_t, interface_type NwInterfaceType) error {
	if _nw_parameters_prohibit_interface_type == nil {
		return symbolCallError("nw_parameters_prohibit_interface_type", "10.14", _nw_parameters_prohibit_interface_typeErr)
	}
	_nw_parameters_prohibit_interface_type(parameters, interface_type)
	return nil
}

// Nw_parameters_prohibit_interface_type prevents connections, listeners, and browsers from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface_type(_:_:)
func Nw_parameters_prohibit_interface_type(parameters Nw_parameters_t, interface_type NwInterfaceType) {
	if callErr := tryNw_parameters_prohibit_interface_type(parameters, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_require_interface func(parameters Nw_parameters_t, interface_ Nw_interface_t)
var _nw_parameters_require_interfaceErr error

func tryNw_parameters_require_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) error {
	if _nw_parameters_require_interface == nil {
		return symbolCallError("nw_parameters_require_interface", "10.14", _nw_parameters_require_interfaceErr)
	}
	_nw_parameters_require_interface(parameters, interface_)
	return nil
}

// Nw_parameters_require_interface sets a specific interface to require on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_require_interface(_:_:)
func Nw_parameters_require_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	if callErr := tryNw_parameters_require_interface(parameters, interface_); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_requires_dnssec_validation func(parameters Nw_parameters_t) bool
var _nw_parameters_requires_dnssec_validationErr error

func tryNw_parameters_requires_dnssec_validation(parameters Nw_parameters_t) (bool, error) {
	if _nw_parameters_requires_dnssec_validation == nil {
		return false, symbolCallError("nw_parameters_requires_dnssec_validation", "13.0", _nw_parameters_requires_dnssec_validationErr)
	}
	return _nw_parameters_requires_dnssec_validation(parameters), nil
}

// Nw_parameters_requires_dnssec_validation checks whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_requires_dnssec_validation(_:)
func Nw_parameters_requires_dnssec_validation(parameters Nw_parameters_t) bool {
	result, callErr := tryNw_parameters_requires_dnssec_validation(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_parameters_set_allow_ultra_constrained func(parameters Nw_parameters_t, allow_ultra_constrained bool)
var _nw_parameters_set_allow_ultra_constrainedErr error

func tryNw_parameters_set_allow_ultra_constrained(parameters Nw_parameters_t, allow_ultra_constrained bool) error {
	if _nw_parameters_set_allow_ultra_constrained == nil {
		return symbolCallError("nw_parameters_set_allow_ultra_constrained", "26.0", _nw_parameters_set_allow_ultra_constrainedErr)
	}
	_nw_parameters_set_allow_ultra_constrained(parameters, allow_ultra_constrained)
	return nil
}

// Nw_parameters_set_allow_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_allow_ultra_constrained(_:_:)
func Nw_parameters_set_allow_ultra_constrained(parameters Nw_parameters_t, allow_ultra_constrained bool) {
	if callErr := tryNw_parameters_set_allow_ultra_constrained(parameters, allow_ultra_constrained); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_attribution func(parameters Nw_parameters_t, attribution Nw_parameters_attribution_t)
var _nw_parameters_set_attributionErr error

func tryNw_parameters_set_attribution(parameters Nw_parameters_t, attribution Nw_parameters_attribution_t) error {
	if _nw_parameters_set_attribution == nil {
		return symbolCallError("nw_parameters_set_attribution", "12.0", _nw_parameters_set_attributionErr)
	}
	_nw_parameters_set_attribution(parameters, attribution)
	return nil
}

// Nw_parameters_set_attribution sets a flag that indicates whether the network request originates from the developer or the user.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_attribution(_:_:)
func Nw_parameters_set_attribution(parameters Nw_parameters_t, attribution Nw_parameters_attribution_t) {
	if callErr := tryNw_parameters_set_attribution(parameters, attribution); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_expired_dns_behavior func(parameters Nw_parameters_t, expired_dns_behavior NwParametersExpiredDnsBehavior)
var _nw_parameters_set_expired_dns_behaviorErr error

func tryNw_parameters_set_expired_dns_behavior(parameters Nw_parameters_t, expired_dns_behavior NwParametersExpiredDnsBehavior) error {
	if _nw_parameters_set_expired_dns_behavior == nil {
		return symbolCallError("nw_parameters_set_expired_dns_behavior", "10.14", _nw_parameters_set_expired_dns_behaviorErr)
	}
	_nw_parameters_set_expired_dns_behavior(parameters, expired_dns_behavior)
	return nil
}

// Nw_parameters_set_expired_dns_behavior sets the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_expired_dns_behavior(_:_:)
func Nw_parameters_set_expired_dns_behavior(parameters Nw_parameters_t, expired_dns_behavior NwParametersExpiredDnsBehavior) {
	if callErr := tryNw_parameters_set_expired_dns_behavior(parameters, expired_dns_behavior); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_fast_open_enabled func(parameters Nw_parameters_t, fast_open_enabled bool)
var _nw_parameters_set_fast_open_enabledErr error

func tryNw_parameters_set_fast_open_enabled(parameters Nw_parameters_t, fast_open_enabled bool) error {
	if _nw_parameters_set_fast_open_enabled == nil {
		return symbolCallError("nw_parameters_set_fast_open_enabled", "10.14", _nw_parameters_set_fast_open_enabledErr)
	}
	_nw_parameters_set_fast_open_enabled(parameters, fast_open_enabled)
	return nil
}

// Nw_parameters_set_fast_open_enabled enables sending application data with protocol handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_fast_open_enabled(_:_:)
func Nw_parameters_set_fast_open_enabled(parameters Nw_parameters_t, fast_open_enabled bool) {
	if callErr := tryNw_parameters_set_fast_open_enabled(parameters, fast_open_enabled); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_include_peer_to_peer func(parameters Nw_parameters_t, include_peer_to_peer bool)
var _nw_parameters_set_include_peer_to_peerErr error

func tryNw_parameters_set_include_peer_to_peer(parameters Nw_parameters_t, include_peer_to_peer bool) error {
	if _nw_parameters_set_include_peer_to_peer == nil {
		return symbolCallError("nw_parameters_set_include_peer_to_peer", "10.14", _nw_parameters_set_include_peer_to_peerErr)
	}
	_nw_parameters_set_include_peer_to_peer(parameters, include_peer_to_peer)
	return nil
}

// Nw_parameters_set_include_peer_to_peer enables peer-to-peer link technologies for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_include_peer_to_peer(_:_:)
func Nw_parameters_set_include_peer_to_peer(parameters Nw_parameters_t, include_peer_to_peer bool) {
	if callErr := tryNw_parameters_set_include_peer_to_peer(parameters, include_peer_to_peer); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_local_endpoint func(parameters Nw_parameters_t, local_endpoint Nw_endpoint_t)
var _nw_parameters_set_local_endpointErr error

func tryNw_parameters_set_local_endpoint(parameters Nw_parameters_t, local_endpoint Nw_endpoint_t) error {
	if _nw_parameters_set_local_endpoint == nil {
		return symbolCallError("nw_parameters_set_local_endpoint", "10.14", _nw_parameters_set_local_endpointErr)
	}
	_nw_parameters_set_local_endpoint(parameters, local_endpoint)
	return nil
}

// Nw_parameters_set_local_endpoint sets a specific local IP address and port to use for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_endpoint(_:_:)
func Nw_parameters_set_local_endpoint(parameters Nw_parameters_t, local_endpoint Nw_endpoint_t) {
	if callErr := tryNw_parameters_set_local_endpoint(parameters, local_endpoint); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_local_only func(parameters Nw_parameters_t, local_only bool)
var _nw_parameters_set_local_onlyErr error

func tryNw_parameters_set_local_only(parameters Nw_parameters_t, local_only bool) error {
	if _nw_parameters_set_local_only == nil {
		return symbolCallError("nw_parameters_set_local_only", "10.14", _nw_parameters_set_local_onlyErr)
	}
	_nw_parameters_set_local_only(parameters, local_only)
	return nil
}

// Nw_parameters_set_local_only restricts listeners to only accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_only(_:_:)
func Nw_parameters_set_local_only(parameters Nw_parameters_t, local_only bool) {
	if callErr := tryNw_parameters_set_local_only(parameters, local_only); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_multipath_service func(parameters Nw_parameters_t, multipath_service NwMultipathService)
var _nw_parameters_set_multipath_serviceErr error

func tryNw_parameters_set_multipath_service(parameters Nw_parameters_t, multipath_service NwMultipathService) error {
	if _nw_parameters_set_multipath_service == nil {
		return symbolCallError("nw_parameters_set_multipath_service", "10.14", _nw_parameters_set_multipath_serviceErr)
	}
	_nw_parameters_set_multipath_service(parameters, multipath_service)
	return nil
}

// Nw_parameters_set_multipath_service enables multipath protocols to allow connections to use multiple interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_multipath_service(_:_:)
func Nw_parameters_set_multipath_service(parameters Nw_parameters_t, multipath_service NwMultipathService) {
	if callErr := tryNw_parameters_set_multipath_service(parameters, multipath_service); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_prefer_no_proxy func(parameters Nw_parameters_t, prefer_no_proxy bool)
var _nw_parameters_set_prefer_no_proxyErr error

func tryNw_parameters_set_prefer_no_proxy(parameters Nw_parameters_t, prefer_no_proxy bool) error {
	if _nw_parameters_set_prefer_no_proxy == nil {
		return symbolCallError("nw_parameters_set_prefer_no_proxy", "10.14", _nw_parameters_set_prefer_no_proxyErr)
	}
	_nw_parameters_set_prefer_no_proxy(parameters, prefer_no_proxy)
	return nil
}

// Nw_parameters_set_prefer_no_proxy sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prefer_no_proxy(_:_:)
func Nw_parameters_set_prefer_no_proxy(parameters Nw_parameters_t, prefer_no_proxy bool) {
	if callErr := tryNw_parameters_set_prefer_no_proxy(parameters, prefer_no_proxy); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_privacy_context func(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t)
var _nw_parameters_set_privacy_contextErr error

func tryNw_parameters_set_privacy_context(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t) error {
	if _nw_parameters_set_privacy_context == nil {
		return symbolCallError("nw_parameters_set_privacy_context", "11.0", _nw_parameters_set_privacy_contextErr)
	}
	_nw_parameters_set_privacy_context(parameters, privacy_context)
	return nil
}

// Nw_parameters_set_privacy_context associates a privacy context with any connections or listeners that use the parameters.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_privacy_context(_:_:)
func Nw_parameters_set_privacy_context(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t) {
	if callErr := tryNw_parameters_set_privacy_context(parameters, privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_prohibit_constrained func(parameters Nw_parameters_t, prohibit_constrained bool)
var _nw_parameters_set_prohibit_constrainedErr error

func tryNw_parameters_set_prohibit_constrained(parameters Nw_parameters_t, prohibit_constrained bool) error {
	if _nw_parameters_set_prohibit_constrained == nil {
		return symbolCallError("nw_parameters_set_prohibit_constrained", "10.15", _nw_parameters_set_prohibit_constrainedErr)
	}
	_nw_parameters_set_prohibit_constrained(parameters, prohibit_constrained)
	return nil
}

// Nw_parameters_set_prohibit_constrained prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_constrained(_:_:)
func Nw_parameters_set_prohibit_constrained(parameters Nw_parameters_t, prohibit_constrained bool) {
	if callErr := tryNw_parameters_set_prohibit_constrained(parameters, prohibit_constrained); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_prohibit_expensive func(parameters Nw_parameters_t, prohibit_expensive bool)
var _nw_parameters_set_prohibit_expensiveErr error

func tryNw_parameters_set_prohibit_expensive(parameters Nw_parameters_t, prohibit_expensive bool) error {
	if _nw_parameters_set_prohibit_expensive == nil {
		return symbolCallError("nw_parameters_set_prohibit_expensive", "10.14", _nw_parameters_set_prohibit_expensiveErr)
	}
	_nw_parameters_set_prohibit_expensive(parameters, prohibit_expensive)
	return nil
}

// Nw_parameters_set_prohibit_expensive prevents connections, listeners, and browsers from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_expensive(_:_:)
func Nw_parameters_set_prohibit_expensive(parameters Nw_parameters_t, prohibit_expensive bool) {
	if callErr := tryNw_parameters_set_prohibit_expensive(parameters, prohibit_expensive); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_required_interface_type func(parameters Nw_parameters_t, interface_type NwInterfaceType)
var _nw_parameters_set_required_interface_typeErr error

func tryNw_parameters_set_required_interface_type(parameters Nw_parameters_t, interface_type NwInterfaceType) error {
	if _nw_parameters_set_required_interface_type == nil {
		return symbolCallError("nw_parameters_set_required_interface_type", "10.14", _nw_parameters_set_required_interface_typeErr)
	}
	_nw_parameters_set_required_interface_type(parameters, interface_type)
	return nil
}

// Nw_parameters_set_required_interface_type sets an interface type to require on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_required_interface_type(_:_:)
func Nw_parameters_set_required_interface_type(parameters Nw_parameters_t, interface_type NwInterfaceType) {
	if callErr := tryNw_parameters_set_required_interface_type(parameters, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_requires_dnssec_validation func(parameters Nw_parameters_t, requires_dnssec_validation bool)
var _nw_parameters_set_requires_dnssec_validationErr error

func tryNw_parameters_set_requires_dnssec_validation(parameters Nw_parameters_t, requires_dnssec_validation bool) error {
	if _nw_parameters_set_requires_dnssec_validation == nil {
		return symbolCallError("nw_parameters_set_requires_dnssec_validation", "13.0", _nw_parameters_set_requires_dnssec_validationErr)
	}
	_nw_parameters_set_requires_dnssec_validation(parameters, requires_dnssec_validation)
	return nil
}

// Nw_parameters_set_requires_dnssec_validation determines whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_requires_dnssec_validation(_:_:)
func Nw_parameters_set_requires_dnssec_validation(parameters Nw_parameters_t, requires_dnssec_validation bool) {
	if callErr := tryNw_parameters_set_requires_dnssec_validation(parameters, requires_dnssec_validation); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_reuse_local_address func(parameters Nw_parameters_t, reuse_local_address bool)
var _nw_parameters_set_reuse_local_addressErr error

func tryNw_parameters_set_reuse_local_address(parameters Nw_parameters_t, reuse_local_address bool) error {
	if _nw_parameters_set_reuse_local_address == nil {
		return symbolCallError("nw_parameters_set_reuse_local_address", "10.14", _nw_parameters_set_reuse_local_addressErr)
	}
	_nw_parameters_set_reuse_local_address(parameters, reuse_local_address)
	return nil
}

// Nw_parameters_set_reuse_local_address allows reusing local addresses and ports across connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_reuse_local_address(_:_:)
func Nw_parameters_set_reuse_local_address(parameters Nw_parameters_t, reuse_local_address bool) {
	if callErr := tryNw_parameters_set_reuse_local_address(parameters, reuse_local_address); callErr != nil {
		panic(callErr)
	}
}

var _nw_parameters_set_service_class func(parameters Nw_parameters_t, service_class NwServiceClass)
var _nw_parameters_set_service_classErr error

func tryNw_parameters_set_service_class(parameters Nw_parameters_t, service_class NwServiceClass) error {
	if _nw_parameters_set_service_class == nil {
		return symbolCallError("nw_parameters_set_service_class", "10.14", _nw_parameters_set_service_classErr)
	}
	_nw_parameters_set_service_class(parameters, service_class)
	return nil
}

// Nw_parameters_set_service_class sets a level of service quality to use for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_service_class(_:_:)
func Nw_parameters_set_service_class(parameters Nw_parameters_t, service_class NwServiceClass) {
	if callErr := tryNw_parameters_set_service_class(parameters, service_class); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_copy_effective_local_endpoint func(path Nw_path_t) Nw_endpoint_t
var _nw_path_copy_effective_local_endpointErr error

func tryNw_path_copy_effective_local_endpoint(path Nw_path_t) (Nw_endpoint_t, error) {
	if _nw_path_copy_effective_local_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_path_copy_effective_local_endpoint", "10.14", _nw_path_copy_effective_local_endpointErr)
	}
	return _nw_path_copy_effective_local_endpoint(path), nil
}

// Nw_path_copy_effective_local_endpoint accesses the local endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_local_endpoint(_:)
func Nw_path_copy_effective_local_endpoint(path Nw_path_t) Nw_endpoint_t {
	result, callErr := tryNw_path_copy_effective_local_endpoint(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_copy_effective_remote_endpoint func(path Nw_path_t) Nw_endpoint_t
var _nw_path_copy_effective_remote_endpointErr error

func tryNw_path_copy_effective_remote_endpoint(path Nw_path_t) (Nw_endpoint_t, error) {
	if _nw_path_copy_effective_remote_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_path_copy_effective_remote_endpoint", "10.14", _nw_path_copy_effective_remote_endpointErr)
	}
	return _nw_path_copy_effective_remote_endpoint(path), nil
}

// Nw_path_copy_effective_remote_endpoint accesses the remote endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_remote_endpoint(_:)
func Nw_path_copy_effective_remote_endpoint(path Nw_path_t) Nw_endpoint_t {
	result, callErr := tryNw_path_copy_effective_remote_endpoint(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_enumerate_gateways func(path Nw_path_t, enumerate_block unsafe.Pointer)
var _nw_path_enumerate_gatewaysErr error

func tryNw_path_enumerate_gateways(path Nw_path_t, enumerate_block Nw_path_enumerate_gateways_block_t) error {
	if _nw_path_enumerate_gateways == nil {
		return symbolCallError("nw_path_enumerate_gateways", "10.15", _nw_path_enumerate_gatewaysErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return enumerate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_path_enumerate_gateways(path, _block0)
	return nil
}

// Nw_path_enumerate_gateways enumerates the list of gateways configured on the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways(_:_:)
func Nw_path_enumerate_gateways(path Nw_path_t, enumerate_block Nw_path_enumerate_gateways_block_t) {
	if callErr := tryNw_path_enumerate_gateways(path, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_enumerate_interfaces func(path Nw_path_t, enumerate_block unsafe.Pointer)
var _nw_path_enumerate_interfacesErr error

func tryNw_path_enumerate_interfaces(path Nw_path_t, enumerate_block Nw_path_enumerate_interfaces_block_t) error {
	if _nw_path_enumerate_interfaces == nil {
		return symbolCallError("nw_path_enumerate_interfaces", "10.14", _nw_path_enumerate_interfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) bool { return enumerate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_path_enumerate_interfaces(path, _block0)
	return nil
}

// Nw_path_enumerate_interfaces enumerates the list of all interfaces available to the path, in order of preference.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces(_:_:)
func Nw_path_enumerate_interfaces(path Nw_path_t, enumerate_block Nw_path_enumerate_interfaces_block_t) {
	if callErr := tryNw_path_enumerate_interfaces(path, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_get_link_quality func(path Nw_path_t) NwLinkQuality
var _nw_path_get_link_qualityErr error

func tryNw_path_get_link_quality(path Nw_path_t) (NwLinkQuality, error) {
	if _nw_path_get_link_quality == nil {
		return *new(NwLinkQuality), symbolCallError("nw_path_get_link_quality", "26.0", _nw_path_get_link_qualityErr)
	}
	return _nw_path_get_link_quality(path), nil
}

// Nw_path_get_link_quality.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_link_quality(_:)
func Nw_path_get_link_quality(path Nw_path_t) NwLinkQuality {
	result, callErr := tryNw_path_get_link_quality(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_get_status func(path Nw_path_t) NwPathStatus
var _nw_path_get_statusErr error

func tryNw_path_get_status(path Nw_path_t) (NwPathStatus, error) {
	if _nw_path_get_status == nil {
		return *new(NwPathStatus), symbolCallError("nw_path_get_status", "10.14", _nw_path_get_statusErr)
	}
	return _nw_path_get_status(path), nil
}

// Nw_path_get_status checks whether a path can be used by connections.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_status(_:)
func Nw_path_get_status(path Nw_path_t) NwPathStatus {
	result, callErr := tryNw_path_get_status(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_get_unsatisfied_reason func(path Nw_path_t) NwPathUnsatisfiedReason
var _nw_path_get_unsatisfied_reasonErr error

func tryNw_path_get_unsatisfied_reason(path Nw_path_t) (NwPathUnsatisfiedReason, error) {
	if _nw_path_get_unsatisfied_reason == nil {
		return *new(NwPathUnsatisfiedReason), symbolCallError("nw_path_get_unsatisfied_reason", "11.0", _nw_path_get_unsatisfied_reasonErr)
	}
	return _nw_path_get_unsatisfied_reason(path), nil
}

// Nw_path_get_unsatisfied_reason.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_unsatisfied_reason(_:)
func Nw_path_get_unsatisfied_reason(path Nw_path_t) NwPathUnsatisfiedReason {
	result, callErr := tryNw_path_get_unsatisfied_reason(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_has_dns func(path Nw_path_t) bool
var _nw_path_has_dnsErr error

func tryNw_path_has_dns(path Nw_path_t) (bool, error) {
	if _nw_path_has_dns == nil {
		return false, symbolCallError("nw_path_has_dns", "10.14", _nw_path_has_dnsErr)
	}
	return _nw_path_has_dns(path), nil
}

// Nw_path_has_dns checks whether the path has a DNS server configured.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_dns(_:)
func Nw_path_has_dns(path Nw_path_t) bool {
	result, callErr := tryNw_path_has_dns(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_has_ipv4 func(path Nw_path_t) bool
var _nw_path_has_ipv4Err error

func tryNw_path_has_ipv4(path Nw_path_t) (bool, error) {
	if _nw_path_has_ipv4 == nil {
		return false, symbolCallError("nw_path_has_ipv4", "10.14", _nw_path_has_ipv4Err)
	}
	return _nw_path_has_ipv4(path), nil
}

// Nw_path_has_ipv4 checks whether the path can route IPv4 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv4(_:)
func Nw_path_has_ipv4(path Nw_path_t) bool {
	result, callErr := tryNw_path_has_ipv4(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_has_ipv6 func(path Nw_path_t) bool
var _nw_path_has_ipv6Err error

func tryNw_path_has_ipv6(path Nw_path_t) (bool, error) {
	if _nw_path_has_ipv6 == nil {
		return false, symbolCallError("nw_path_has_ipv6", "10.14", _nw_path_has_ipv6Err)
	}
	return _nw_path_has_ipv6(path), nil
}

// Nw_path_has_ipv6 checks whether the path can route IPv6 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv6(_:)
func Nw_path_has_ipv6(path Nw_path_t) bool {
	result, callErr := tryNw_path_has_ipv6(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_is_constrained func(path Nw_path_t) bool
var _nw_path_is_constrainedErr error

func tryNw_path_is_constrained(path Nw_path_t) (bool, error) {
	if _nw_path_is_constrained == nil {
		return false, symbolCallError("nw_path_is_constrained", "10.15", _nw_path_is_constrainedErr)
	}
	return _nw_path_is_constrained(path), nil
}

// Nw_path_is_constrained checks whether the path uses an interface in Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_constrained(_:)
func Nw_path_is_constrained(path Nw_path_t) bool {
	result, callErr := tryNw_path_is_constrained(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_is_equal func(path Nw_path_t, other_path Nw_path_t) bool
var _nw_path_is_equalErr error

func tryNw_path_is_equal(path Nw_path_t, other_path Nw_path_t) (bool, error) {
	if _nw_path_is_equal == nil {
		return false, symbolCallError("nw_path_is_equal", "10.14", _nw_path_is_equalErr)
	}
	return _nw_path_is_equal(path, other_path), nil
}

// Nw_path_is_equal compares if two paths are identical.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_equal(_:_:)
func Nw_path_is_equal(path Nw_path_t, other_path Nw_path_t) bool {
	result, callErr := tryNw_path_is_equal(path, other_path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_is_expensive func(path Nw_path_t) bool
var _nw_path_is_expensiveErr error

func tryNw_path_is_expensive(path Nw_path_t) (bool, error) {
	if _nw_path_is_expensive == nil {
		return false, symbolCallError("nw_path_is_expensive", "10.14", _nw_path_is_expensiveErr)
	}
	return _nw_path_is_expensive(path), nil
}

// Nw_path_is_expensive checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_expensive(_:)
func Nw_path_is_expensive(path Nw_path_t) bool {
	result, callErr := tryNw_path_is_expensive(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_is_ultra_constrained func(path Nw_path_t) bool
var _nw_path_is_ultra_constrainedErr error

func tryNw_path_is_ultra_constrained(path Nw_path_t) (bool, error) {
	if _nw_path_is_ultra_constrained == nil {
		return false, symbolCallError("nw_path_is_ultra_constrained", "26.0", _nw_path_is_ultra_constrainedErr)
	}
	return _nw_path_is_ultra_constrained(path), nil
}

// Nw_path_is_ultra_constrained.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_ultra_constrained(_:)
func Nw_path_is_ultra_constrained(path Nw_path_t) bool {
	result, callErr := tryNw_path_is_ultra_constrained(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_monitor_cancel func(monitor Nw_path_monitor_t)
var _nw_path_monitor_cancelErr error

func tryNw_path_monitor_cancel(monitor Nw_path_monitor_t) error {
	if _nw_path_monitor_cancel == nil {
		return symbolCallError("nw_path_monitor_cancel", "10.14", _nw_path_monitor_cancelErr)
	}
	_nw_path_monitor_cancel(monitor)
	return nil
}

// Nw_path_monitor_cancel stops receiving network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel(_:)
func Nw_path_monitor_cancel(monitor Nw_path_monitor_t) {
	if callErr := tryNw_path_monitor_cancel(monitor); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_monitor_create func() Nw_path_monitor_t
var _nw_path_monitor_createErr error

func tryNw_path_monitor_create() (Nw_path_monitor_t, error) {
	if _nw_path_monitor_create == nil {
		return *new(Nw_path_monitor_t), symbolCallError("nw_path_monitor_create", "10.14", _nw_path_monitor_createErr)
	}
	return _nw_path_monitor_create(), nil
}

// Nw_path_monitor_create initializes a path monitor to observe all available interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create()
func Nw_path_monitor_create() Nw_path_monitor_t {
	result, callErr := tryNw_path_monitor_create()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_monitor_create_for_ethernet_channel func() Nw_path_monitor_t
var _nw_path_monitor_create_for_ethernet_channelErr error

func tryNw_path_monitor_create_for_ethernet_channel() (Nw_path_monitor_t, error) {
	if _nw_path_monitor_create_for_ethernet_channel == nil {
		return *new(Nw_path_monitor_t), symbolCallError("nw_path_monitor_create_for_ethernet_channel", "13.0", _nw_path_monitor_create_for_ethernet_channelErr)
	}
	return _nw_path_monitor_create_for_ethernet_channel(), nil
}

// Nw_path_monitor_create_for_ethernet_channel.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_for_ethernet_channel()
func Nw_path_monitor_create_for_ethernet_channel() Nw_path_monitor_t {
	result, callErr := tryNw_path_monitor_create_for_ethernet_channel()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_monitor_create_with_type func(required_interface_type NwInterfaceType) Nw_path_monitor_t
var _nw_path_monitor_create_with_typeErr error

func tryNw_path_monitor_create_with_type(required_interface_type NwInterfaceType) (Nw_path_monitor_t, error) {
	if _nw_path_monitor_create_with_type == nil {
		return *new(Nw_path_monitor_t), symbolCallError("nw_path_monitor_create_with_type", "10.14", _nw_path_monitor_create_with_typeErr)
	}
	return _nw_path_monitor_create_with_type(required_interface_type), nil
}

// Nw_path_monitor_create_with_type initializes a path monitor to observe a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_with_type(_:)
func Nw_path_monitor_create_with_type(required_interface_type NwInterfaceType) Nw_path_monitor_t {
	result, callErr := tryNw_path_monitor_create_with_type(required_interface_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_path_monitor_prohibit_interface_type func(monitor Nw_path_monitor_t, interface_type NwInterfaceType)
var _nw_path_monitor_prohibit_interface_typeErr error

func tryNw_path_monitor_prohibit_interface_type(monitor Nw_path_monitor_t, interface_type NwInterfaceType) error {
	if _nw_path_monitor_prohibit_interface_type == nil {
		return symbolCallError("nw_path_monitor_prohibit_interface_type", "11.0", _nw_path_monitor_prohibit_interface_typeErr)
	}
	_nw_path_monitor_prohibit_interface_type(monitor, interface_type)
	return nil
}

// Nw_path_monitor_prohibit_interface_type prohibit a path monitor from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_prohibit_interface_type(_:_:)
func Nw_path_monitor_prohibit_interface_type(monitor Nw_path_monitor_t, interface_type NwInterfaceType) {
	if callErr := tryNw_path_monitor_prohibit_interface_type(monitor, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_monitor_set_cancel_handler func(monitor Nw_path_monitor_t, cancel_handler unsafe.Pointer)
var _nw_path_monitor_set_cancel_handlerErr error

func tryNw_path_monitor_set_cancel_handler(monitor Nw_path_monitor_t, cancel_handler Nw_path_monitor_cancel_handler_t) error {
	if _nw_path_monitor_set_cancel_handler == nil {
		return symbolCallError("nw_path_monitor_set_cancel_handler", "10.14", _nw_path_monitor_set_cancel_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { cancel_handler() })
	retainNetworkAsyncBlock(monitor.ID, "nw_path_monitor_set_cancel_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_path_monitor_set_cancel_handler(monitor, _block0)
	return nil
}

// Nw_path_monitor_set_cancel_handler sets a handler to determine when a monitor is fully cancelled and will no longer deliver events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_cancel_handler(_:_:)
func Nw_path_monitor_set_cancel_handler(monitor Nw_path_monitor_t, cancel_handler Nw_path_monitor_cancel_handler_t) {
	if callErr := tryNw_path_monitor_set_cancel_handler(monitor, cancel_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_monitor_set_queue func(monitor Nw_path_monitor_t, queue uintptr)
var _nw_path_monitor_set_queueErr error

func tryNw_path_monitor_set_queue(monitor Nw_path_monitor_t, queue dispatch.Queue) error {
	if _nw_path_monitor_set_queue == nil {
		return symbolCallError("nw_path_monitor_set_queue", "10.14", _nw_path_monitor_set_queueErr)
	}
	_nw_path_monitor_set_queue(monitor, uintptr(queue.Handle()))
	return nil
}

// Nw_path_monitor_set_queue sets a queue on which to deliver path events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_queue(_:_:)
func Nw_path_monitor_set_queue(monitor Nw_path_monitor_t, queue dispatch.Queue) {
	if callErr := tryNw_path_monitor_set_queue(monitor, queue); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_monitor_set_update_handler func(monitor Nw_path_monitor_t, update_handler unsafe.Pointer)
var _nw_path_monitor_set_update_handlerErr error

func tryNw_path_monitor_set_update_handler(monitor Nw_path_monitor_t, update_handler Nw_path_monitor_update_handler_t) error {
	if _nw_path_monitor_set_update_handler == nil {
		return symbolCallError("nw_path_monitor_set_update_handler", "10.14", _nw_path_monitor_set_update_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { update_handler(arg0) })
	retainNetworkAsyncBlock(monitor.ID, "nw_path_monitor_set_update_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_path_monitor_set_update_handler(monitor, _block0)
	return nil
}

// Nw_path_monitor_set_update_handler sets a handler to receive network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_update_handler(_:_:)
func Nw_path_monitor_set_update_handler(monitor Nw_path_monitor_t, update_handler Nw_path_monitor_update_handler_t) {
	if callErr := tryNw_path_monitor_set_update_handler(monitor, update_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_monitor_start func(monitor Nw_path_monitor_t)
var _nw_path_monitor_startErr error

func tryNw_path_monitor_start(monitor Nw_path_monitor_t) error {
	if _nw_path_monitor_start == nil {
		return symbolCallError("nw_path_monitor_start", "10.14", _nw_path_monitor_startErr)
	}
	_nw_path_monitor_start(monitor)
	return nil
}

// Nw_path_monitor_start starts monitoring path changes.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_start(_:)
func Nw_path_monitor_start(monitor Nw_path_monitor_t) {
	if callErr := tryNw_path_monitor_start(monitor); callErr != nil {
		panic(callErr)
	}
}

var _nw_path_uses_interface_type func(path Nw_path_t, interface_type NwInterfaceType) bool
var _nw_path_uses_interface_typeErr error

func tryNw_path_uses_interface_type(path Nw_path_t, interface_type NwInterfaceType) (bool, error) {
	if _nw_path_uses_interface_type == nil {
		return false, symbolCallError("nw_path_uses_interface_type", "10.14", _nw_path_uses_interface_typeErr)
	}
	return _nw_path_uses_interface_type(path, interface_type), nil
}

// Nw_path_uses_interface_type checks if connections using the path may send traffic over a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_uses_interface_type(_:_:)
func Nw_path_uses_interface_type(path Nw_path_t, interface_type NwInterfaceType) bool {
	result, callErr := tryNw_path_uses_interface_type(path, interface_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_privacy_context_add_proxy func(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t)
var _nw_privacy_context_add_proxyErr error

func tryNw_privacy_context_add_proxy(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t) error {
	if _nw_privacy_context_add_proxy == nil {
		return symbolCallError("nw_privacy_context_add_proxy", "14.0", _nw_privacy_context_add_proxyErr)
	}
	_nw_privacy_context_add_proxy(privacy_context, proxy_config)
	return nil
}

// Nw_privacy_context_add_proxy applies a proxy configuration to all connections associated with this context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_add_proxy(_:_:)
func Nw_privacy_context_add_proxy(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t) {
	if callErr := tryNw_privacy_context_add_proxy(privacy_context, proxy_config); callErr != nil {
		panic(callErr)
	}
}

var _nw_privacy_context_clear_proxies func(privacy_context Nw_privacy_context_t)
var _nw_privacy_context_clear_proxiesErr error

func tryNw_privacy_context_clear_proxies(privacy_context Nw_privacy_context_t) error {
	if _nw_privacy_context_clear_proxies == nil {
		return symbolCallError("nw_privacy_context_clear_proxies", "14.0", _nw_privacy_context_clear_proxiesErr)
	}
	_nw_privacy_context_clear_proxies(privacy_context)
	return nil
}

// Nw_privacy_context_clear_proxies clears out any proxies added using nw_privacy_context_add_proxy(_:_:)
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_clear_proxies(_:)
func Nw_privacy_context_clear_proxies(privacy_context Nw_privacy_context_t) {
	if callErr := tryNw_privacy_context_clear_proxies(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nw_privacy_context_create func(description string) Nw_privacy_context_t
var _nw_privacy_context_createErr error

func tryNw_privacy_context_create(description string) (Nw_privacy_context_t, error) {
	if _nw_privacy_context_create == nil {
		return *new(Nw_privacy_context_t), symbolCallError("nw_privacy_context_create", "11.0", _nw_privacy_context_createErr)
	}
	return _nw_privacy_context_create(description), nil
}

// Nw_privacy_context_create initializes a privacy context with a description string.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_create(_:)
func Nw_privacy_context_create(description string) Nw_privacy_context_t {
	result, callErr := tryNw_privacy_context_create(description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_privacy_context_disable_logging func(privacy_context Nw_privacy_context_t)
var _nw_privacy_context_disable_loggingErr error

func tryNw_privacy_context_disable_logging(privacy_context Nw_privacy_context_t) error {
	if _nw_privacy_context_disable_logging == nil {
		return symbolCallError("nw_privacy_context_disable_logging", "11.0", _nw_privacy_context_disable_loggingErr)
	}
	_nw_privacy_context_disable_logging(privacy_context)
	return nil
}

// Nw_privacy_context_disable_logging disables system logging of connection activity.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_disable_logging(_:)
func Nw_privacy_context_disable_logging(privacy_context Nw_privacy_context_t) {
	if callErr := tryNw_privacy_context_disable_logging(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nw_privacy_context_flush_cache func(privacy_context Nw_privacy_context_t)
var _nw_privacy_context_flush_cacheErr error

func tryNw_privacy_context_flush_cache(privacy_context Nw_privacy_context_t) error {
	if _nw_privacy_context_flush_cache == nil {
		return symbolCallError("nw_privacy_context_flush_cache", "11.0", _nw_privacy_context_flush_cacheErr)
	}
	_nw_privacy_context_flush_cache(privacy_context)
	return nil
}

// Nw_privacy_context_flush_cache flushes all cached data, such as TLS session state, created by connections associated with the privacy context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_flush_cache(_:)
func Nw_privacy_context_flush_cache(privacy_context Nw_privacy_context_t) {
	if callErr := tryNw_privacy_context_flush_cache(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nw_privacy_context_require_encrypted_name_resolution func(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t)
var _nw_privacy_context_require_encrypted_name_resolutionErr error

func tryNw_privacy_context_require_encrypted_name_resolution(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t) error {
	if _nw_privacy_context_require_encrypted_name_resolution == nil {
		return symbolCallError("nw_privacy_context_require_encrypted_name_resolution", "11.0", _nw_privacy_context_require_encrypted_name_resolutionErr)
	}
	_nw_privacy_context_require_encrypted_name_resolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config)
	return nil
}

// Nw_privacy_context_require_encrypted_name_resolution requires that any DNS name resolution for connections associated with this context use encrypted transports, such as TLS or HTTPS.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_require_encrypted_name_resolution(_:_:_:)
func Nw_privacy_context_require_encrypted_name_resolution(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t) {
	if callErr := tryNw_privacy_context_require_encrypted_name_resolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config); callErr != nil {
		panic(callErr)
	}
}

var _nw_protocol_copy_ip_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_ip_definitionErr error

func tryNw_protocol_copy_ip_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_ip_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_ip_definition", "10.14", _nw_protocol_copy_ip_definitionErr)
	}
	return _nw_protocol_copy_ip_definition(), nil
}

// Nw_protocol_copy_ip_definition accesses the system definition of the Internet Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ip_definition()
func Nw_protocol_copy_ip_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_ip_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_copy_quic_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_quic_definitionErr error

func tryNw_protocol_copy_quic_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_quic_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_quic_definition", "12.0", _nw_protocol_copy_quic_definitionErr)
	}
	return _nw_protocol_copy_quic_definition(), nil
}

// Nw_protocol_copy_quic_definition accesses the system definition of the QUIC transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_quic_definition()
func Nw_protocol_copy_quic_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_quic_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_copy_tcp_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_tcp_definitionErr error

func tryNw_protocol_copy_tcp_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_tcp_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_tcp_definition", "10.14", _nw_protocol_copy_tcp_definitionErr)
	}
	return _nw_protocol_copy_tcp_definition(), nil
}

// Nw_protocol_copy_tcp_definition accesses the system definition of the Transport Control Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tcp_definition()
func Nw_protocol_copy_tcp_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_tcp_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_copy_tls_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_tls_definitionErr error

func tryNw_protocol_copy_tls_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_tls_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_tls_definition", "10.14", _nw_protocol_copy_tls_definitionErr)
	}
	return _nw_protocol_copy_tls_definition(), nil
}

// Nw_protocol_copy_tls_definition accesses the system definition of the Transport Layer Security protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tls_definition()
func Nw_protocol_copy_tls_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_tls_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_copy_udp_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_udp_definitionErr error

func tryNw_protocol_copy_udp_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_udp_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_udp_definition", "10.14", _nw_protocol_copy_udp_definitionErr)
	}
	return _nw_protocol_copy_udp_definition(), nil
}

// Nw_protocol_copy_udp_definition accesses the system definition of the User Datagram Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_udp_definition()
func Nw_protocol_copy_udp_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_udp_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_copy_ws_definition func() Nw_protocol_definition_t
var _nw_protocol_copy_ws_definitionErr error

func tryNw_protocol_copy_ws_definition() (Nw_protocol_definition_t, error) {
	if _nw_protocol_copy_ws_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_copy_ws_definition", "10.15", _nw_protocol_copy_ws_definitionErr)
	}
	return _nw_protocol_copy_ws_definition(), nil
}

// Nw_protocol_copy_ws_definition accesses the system definition of the WebSocket protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ws_definition()
func Nw_protocol_copy_ws_definition() Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_copy_ws_definition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_definition_is_equal func(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) bool
var _nw_protocol_definition_is_equalErr error

func tryNw_protocol_definition_is_equal(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) (bool, error) {
	if _nw_protocol_definition_is_equal == nil {
		return false, symbolCallError("nw_protocol_definition_is_equal", "10.14", _nw_protocol_definition_is_equalErr)
	}
	return _nw_protocol_definition_is_equal(definition1, definition2), nil
}

// Nw_protocol_definition_is_equal compares two protocol definitions, and returns true if they represent the same protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_definition_is_equal(_:_:)
func Nw_protocol_definition_is_equal(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) bool {
	result, callErr := tryNw_protocol_definition_is_equal(definition1, definition2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_copy_definition func(metadata Nw_protocol_metadata_t) Nw_protocol_definition_t
var _nw_protocol_metadata_copy_definitionErr error

func tryNw_protocol_metadata_copy_definition(metadata Nw_protocol_metadata_t) (Nw_protocol_definition_t, error) {
	if _nw_protocol_metadata_copy_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_metadata_copy_definition", "10.14", _nw_protocol_metadata_copy_definitionErr)
	}
	return _nw_protocol_metadata_copy_definition(metadata), nil
}

// Nw_protocol_metadata_copy_definition accesses the protocol definition associated with the metadata object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_copy_definition(_:)
func Nw_protocol_metadata_copy_definition(metadata Nw_protocol_metadata_t) Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_metadata_copy_definition(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_framer_message func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_framer_messageErr error

func tryNw_protocol_metadata_is_framer_message(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_framer_message == nil {
		return false, symbolCallError("nw_protocol_metadata_is_framer_message", "10.15", _nw_protocol_metadata_is_framer_messageErr)
	}
	return _nw_protocol_metadata_is_framer_message(metadata), nil
}

// Nw_protocol_metadata_is_framer_message checks if a metadata object represents a custom framer protocol message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_framer_message(_:)
func Nw_protocol_metadata_is_framer_message(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_framer_message(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_ip func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_ipErr error

func tryNw_protocol_metadata_is_ip(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_ip == nil {
		return false, symbolCallError("nw_protocol_metadata_is_ip", "10.14", _nw_protocol_metadata_is_ipErr)
	}
	return _nw_protocol_metadata_is_ip(metadata), nil
}

// Nw_protocol_metadata_is_ip checks whether a metadata object represents an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ip(_:)
func Nw_protocol_metadata_is_ip(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_ip(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_quic func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_quicErr error

func tryNw_protocol_metadata_is_quic(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_quic == nil {
		return false, symbolCallError("nw_protocol_metadata_is_quic", "12.0", _nw_protocol_metadata_is_quicErr)
	}
	return _nw_protocol_metadata_is_quic(metadata), nil
}

// Nw_protocol_metadata_is_quic checks whether a metadata object contains QUIC connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_quic(_:)
func Nw_protocol_metadata_is_quic(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_quic(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_tcp func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_tcpErr error

func tryNw_protocol_metadata_is_tcp(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_tcp == nil {
		return false, symbolCallError("nw_protocol_metadata_is_tcp", "10.14", _nw_protocol_metadata_is_tcpErr)
	}
	return _nw_protocol_metadata_is_tcp(metadata), nil
}

// Nw_protocol_metadata_is_tcp checks whether a metadata object contains TCP connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tcp(_:)
func Nw_protocol_metadata_is_tcp(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_tcp(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_tls func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_tlsErr error

func tryNw_protocol_metadata_is_tls(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_tls == nil {
		return false, symbolCallError("nw_protocol_metadata_is_tls", "10.14", _nw_protocol_metadata_is_tlsErr)
	}
	return _nw_protocol_metadata_is_tls(metadata), nil
}

// Nw_protocol_metadata_is_tls checks whether a metadata object contains TLS connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tls(_:)
func Nw_protocol_metadata_is_tls(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_tls(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_udp func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_udpErr error

func tryNw_protocol_metadata_is_udp(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_udp == nil {
		return false, symbolCallError("nw_protocol_metadata_is_udp", "10.14", _nw_protocol_metadata_is_udpErr)
	}
	return _nw_protocol_metadata_is_udp(metadata), nil
}

// Nw_protocol_metadata_is_udp checks whether a metadata object represents a UDP datagram.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_udp(_:)
func Nw_protocol_metadata_is_udp(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_udp(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_metadata_is_ws func(metadata Nw_protocol_metadata_t) bool
var _nw_protocol_metadata_is_wsErr error

func tryNw_protocol_metadata_is_ws(metadata Nw_protocol_metadata_t) (bool, error) {
	if _nw_protocol_metadata_is_ws == nil {
		return false, symbolCallError("nw_protocol_metadata_is_ws", "10.15", _nw_protocol_metadata_is_wsErr)
	}
	return _nw_protocol_metadata_is_ws(metadata), nil
}

// Nw_protocol_metadata_is_ws checks whether a metadata object represents a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ws(_:)
func Nw_protocol_metadata_is_ws(metadata Nw_protocol_metadata_t) bool {
	result, callErr := tryNw_protocol_metadata_is_ws(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_options_copy_definition func(options Nw_protocol_options_t) Nw_protocol_definition_t
var _nw_protocol_options_copy_definitionErr error

func tryNw_protocol_options_copy_definition(options Nw_protocol_options_t) (Nw_protocol_definition_t, error) {
	if _nw_protocol_options_copy_definition == nil {
		return *new(Nw_protocol_definition_t), symbolCallError("nw_protocol_options_copy_definition", "10.14", _nw_protocol_options_copy_definitionErr)
	}
	return _nw_protocol_options_copy_definition(options), nil
}

// Nw_protocol_options_copy_definition accesses the protocol definition associated with the options object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_copy_definition(_:)
func Nw_protocol_options_copy_definition(options Nw_protocol_options_t) Nw_protocol_definition_t {
	result, callErr := tryNw_protocol_options_copy_definition(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_options_is_quic func(options Nw_protocol_options_t) bool
var _nw_protocol_options_is_quicErr error

func tryNw_protocol_options_is_quic(options Nw_protocol_options_t) (bool, error) {
	if _nw_protocol_options_is_quic == nil {
		return false, symbolCallError("nw_protocol_options_is_quic", "12.0", _nw_protocol_options_is_quicErr)
	}
	return _nw_protocol_options_is_quic(options), nil
}

// Nw_protocol_options_is_quic checks whether an options object uses the QUIC protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_is_quic(_:)
func Nw_protocol_options_is_quic(options Nw_protocol_options_t) bool {
	result, callErr := tryNw_protocol_options_is_quic(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_stack_clear_application_protocols func(stack Nw_protocol_stack_t)
var _nw_protocol_stack_clear_application_protocolsErr error

func tryNw_protocol_stack_clear_application_protocols(stack Nw_protocol_stack_t) error {
	if _nw_protocol_stack_clear_application_protocols == nil {
		return symbolCallError("nw_protocol_stack_clear_application_protocols", "10.14", _nw_protocol_stack_clear_application_protocolsErr)
	}
	_nw_protocol_stack_clear_application_protocols(stack)
	return nil
}

// Nw_protocol_stack_clear_application_protocols removes all application protocols from the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_clear_application_protocols(_:)
func Nw_protocol_stack_clear_application_protocols(stack Nw_protocol_stack_t) {
	if callErr := tryNw_protocol_stack_clear_application_protocols(stack); callErr != nil {
		panic(callErr)
	}
}

var _nw_protocol_stack_copy_internet_protocol func(stack Nw_protocol_stack_t) Nw_protocol_options_t
var _nw_protocol_stack_copy_internet_protocolErr error

func tryNw_protocol_stack_copy_internet_protocol(stack Nw_protocol_stack_t) (Nw_protocol_options_t, error) {
	if _nw_protocol_stack_copy_internet_protocol == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_protocol_stack_copy_internet_protocol", "10.14", _nw_protocol_stack_copy_internet_protocolErr)
	}
	return _nw_protocol_stack_copy_internet_protocol(stack), nil
}

// Nw_protocol_stack_copy_internet_protocol accesses the protocol stack’s Internet Protocol options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_internet_protocol(_:)
func Nw_protocol_stack_copy_internet_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	result, callErr := tryNw_protocol_stack_copy_internet_protocol(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_stack_copy_transport_protocol func(stack Nw_protocol_stack_t) Nw_protocol_options_t
var _nw_protocol_stack_copy_transport_protocolErr error

func tryNw_protocol_stack_copy_transport_protocol(stack Nw_protocol_stack_t) (Nw_protocol_options_t, error) {
	if _nw_protocol_stack_copy_transport_protocol == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_protocol_stack_copy_transport_protocol", "10.14", _nw_protocol_stack_copy_transport_protocolErr)
	}
	return _nw_protocol_stack_copy_transport_protocol(stack), nil
}

// Nw_protocol_stack_copy_transport_protocol accesses the options for the protocol stack’s transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_transport_protocol(_:)
func Nw_protocol_stack_copy_transport_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	result, callErr := tryNw_protocol_stack_copy_transport_protocol(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_protocol_stack_iterate_application_protocols func(stack Nw_protocol_stack_t, iterate_block unsafe.Pointer)
var _nw_protocol_stack_iterate_application_protocolsErr error

func tryNw_protocol_stack_iterate_application_protocols(stack Nw_protocol_stack_t, iterate_block Nw_protocol_stack_iterate_protocols_block_t) error {
	if _nw_protocol_stack_iterate_application_protocols == nil {
		return symbolCallError("nw_protocol_stack_iterate_application_protocols", "10.14", _nw_protocol_stack_iterate_application_protocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { iterate_block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_protocol_stack_iterate_application_protocols(stack, _block0)
	return nil
}

// Nw_protocol_stack_iterate_application_protocols iterates through the array of application protocol options that will be used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_application_protocols(_:_:)
func Nw_protocol_stack_iterate_application_protocols(stack Nw_protocol_stack_t, iterate_block Nw_protocol_stack_iterate_protocols_block_t) {
	if callErr := tryNw_protocol_stack_iterate_application_protocols(stack, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nw_protocol_stack_prepend_application_protocol func(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t)
var _nw_protocol_stack_prepend_application_protocolErr error

func tryNw_protocol_stack_prepend_application_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) error {
	if _nw_protocol_stack_prepend_application_protocol == nil {
		return symbolCallError("nw_protocol_stack_prepend_application_protocol", "10.14", _nw_protocol_stack_prepend_application_protocolErr)
	}
	_nw_protocol_stack_prepend_application_protocol(stack, protocol_)
	return nil
}

// Nw_protocol_stack_prepend_application_protocol adds a protocol onto the top of the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_prepend_application_protocol(_:_:)
func Nw_protocol_stack_prepend_application_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	if callErr := tryNw_protocol_stack_prepend_application_protocol(stack, protocol_); callErr != nil {
		panic(callErr)
	}
}

var _nw_protocol_stack_set_transport_protocol func(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t)
var _nw_protocol_stack_set_transport_protocolErr error

func tryNw_protocol_stack_set_transport_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) error {
	if _nw_protocol_stack_set_transport_protocol == nil {
		return symbolCallError("nw_protocol_stack_set_transport_protocol", "10.14", _nw_protocol_stack_set_transport_protocolErr)
	}
	_nw_protocol_stack_set_transport_protocol(stack, protocol_)
	return nil
}

// Nw_protocol_stack_set_transport_protocol replaces the protocol stack’s transport protocol with a new set of options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_set_transport_protocol(_:_:)
func Nw_protocol_stack_set_transport_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	if callErr := tryNw_protocol_stack_set_transport_protocol(stack, protocol_); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_add_excluded_domain func(config Nw_proxy_config_t, excluded_domain string)
var _nw_proxy_config_add_excluded_domainErr error

func tryNw_proxy_config_add_excluded_domain(config Nw_proxy_config_t, excluded_domain string) error {
	if _nw_proxy_config_add_excluded_domain == nil {
		return symbolCallError("nw_proxy_config_add_excluded_domain", "14.0", _nw_proxy_config_add_excluded_domainErr)
	}
	_nw_proxy_config_add_excluded_domain(config, excluded_domain)
	return nil
}

// Nw_proxy_config_add_excluded_domain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_excluded_domain(_:_:)
func Nw_proxy_config_add_excluded_domain(config Nw_proxy_config_t, excluded_domain string) {
	if callErr := tryNw_proxy_config_add_excluded_domain(config, excluded_domain); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_add_match_domain func(config Nw_proxy_config_t, match_domain string)
var _nw_proxy_config_add_match_domainErr error

func tryNw_proxy_config_add_match_domain(config Nw_proxy_config_t, match_domain string) error {
	if _nw_proxy_config_add_match_domain == nil {
		return symbolCallError("nw_proxy_config_add_match_domain", "14.0", _nw_proxy_config_add_match_domainErr)
	}
	_nw_proxy_config_add_match_domain(config, match_domain)
	return nil
}

// Nw_proxy_config_add_match_domain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_match_domain(_:_:)
func Nw_proxy_config_add_match_domain(config Nw_proxy_config_t, match_domain string) {
	if callErr := tryNw_proxy_config_add_match_domain(config, match_domain); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_clear_excluded_domains func(config Nw_proxy_config_t)
var _nw_proxy_config_clear_excluded_domainsErr error

func tryNw_proxy_config_clear_excluded_domains(config Nw_proxy_config_t) error {
	if _nw_proxy_config_clear_excluded_domains == nil {
		return symbolCallError("nw_proxy_config_clear_excluded_domains", "14.0", _nw_proxy_config_clear_excluded_domainsErr)
	}
	_nw_proxy_config_clear_excluded_domains(config)
	return nil
}

// Nw_proxy_config_clear_excluded_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_excluded_domains(_:)
func Nw_proxy_config_clear_excluded_domains(config Nw_proxy_config_t) {
	if callErr := tryNw_proxy_config_clear_excluded_domains(config); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_clear_match_domains func(config Nw_proxy_config_t)
var _nw_proxy_config_clear_match_domainsErr error

func tryNw_proxy_config_clear_match_domains(config Nw_proxy_config_t) error {
	if _nw_proxy_config_clear_match_domains == nil {
		return symbolCallError("nw_proxy_config_clear_match_domains", "14.0", _nw_proxy_config_clear_match_domainsErr)
	}
	_nw_proxy_config_clear_match_domains(config)
	return nil
}

// Nw_proxy_config_clear_match_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_match_domains(_:)
func Nw_proxy_config_clear_match_domains(config Nw_proxy_config_t) {
	if callErr := tryNw_proxy_config_clear_match_domains(config); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_create_http_connect func(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) Nw_proxy_config_t
var _nw_proxy_config_create_http_connectErr error

func tryNw_proxy_config_create_http_connect(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) (Nw_proxy_config_t, error) {
	if _nw_proxy_config_create_http_connect == nil {
		return *new(Nw_proxy_config_t), symbolCallError("nw_proxy_config_create_http_connect", "14.0", _nw_proxy_config_create_http_connectErr)
	}
	return _nw_proxy_config_create_http_connect(proxy_endpoint, proxy_tls_options), nil
}

// Nw_proxy_config_create_http_connect initializes a legacy HTTP CONNECT configuration for a proxy server accessible using HTTP/1.1.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_http_connect(_:_:)
func Nw_proxy_config_create_http_connect(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) Nw_proxy_config_t {
	result, callErr := tryNw_proxy_config_create_http_connect(proxy_endpoint, proxy_tls_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_proxy_config_create_oblivious_http func(relay Nw_relay_hop_t, relay_resource_path string, gateway_key_config *uint8, gateway_key_config_length uintptr) Nw_proxy_config_t
var _nw_proxy_config_create_oblivious_httpErr error

func tryNw_proxy_config_create_oblivious_http(relay Nw_relay_hop_t, relay_resource_path string, gateway_key_config *uint8, gateway_key_config_length uintptr) (Nw_proxy_config_t, error) {
	if _nw_proxy_config_create_oblivious_http == nil {
		return *new(Nw_proxy_config_t), symbolCallError("nw_proxy_config_create_oblivious_http", "14.0", _nw_proxy_config_create_oblivious_httpErr)
	}
	return _nw_proxy_config_create_oblivious_http(relay, relay_resource_path, gateway_key_config, gateway_key_config_length), nil
}

// Nw_proxy_config_create_oblivious_http initializes an Oblivious HTTP proxy configuration using a relay and a gateway.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_oblivious_http(_:_:_:_:)
func Nw_proxy_config_create_oblivious_http(relay Nw_relay_hop_t, relay_resource_path string, gateway_key_config *uint8, gateway_key_config_length uintptr) Nw_proxy_config_t {
	result, callErr := tryNw_proxy_config_create_oblivious_http(relay, relay_resource_path, gateway_key_config, gateway_key_config_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_proxy_config_create_relay func(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) Nw_proxy_config_t
var _nw_proxy_config_create_relayErr error

func tryNw_proxy_config_create_relay(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) (Nw_proxy_config_t, error) {
	if _nw_proxy_config_create_relay == nil {
		return *new(Nw_proxy_config_t), symbolCallError("nw_proxy_config_create_relay", "14.0", _nw_proxy_config_create_relayErr)
	}
	return _nw_proxy_config_create_relay(first_hop, second_hop), nil
}

// Nw_proxy_config_create_relay initializes a proxy configuration with one or two relay hops.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_relay(_:_:)
func Nw_proxy_config_create_relay(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) Nw_proxy_config_t {
	result, callErr := tryNw_proxy_config_create_relay(first_hop, second_hop)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_proxy_config_create_socksv5 func(proxy_endpoint Nw_endpoint_t) Nw_proxy_config_t
var _nw_proxy_config_create_socksv5Err error

func tryNw_proxy_config_create_socksv5(proxy_endpoint Nw_endpoint_t) (Nw_proxy_config_t, error) {
	if _nw_proxy_config_create_socksv5 == nil {
		return *new(Nw_proxy_config_t), symbolCallError("nw_proxy_config_create_socksv5", "14.0", _nw_proxy_config_create_socksv5Err)
	}
	return _nw_proxy_config_create_socksv5(proxy_endpoint), nil
}

// Nw_proxy_config_create_socksv5 initializes a SOCKSv5 proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_socksv5(_:)
func Nw_proxy_config_create_socksv5(proxy_endpoint Nw_endpoint_t) Nw_proxy_config_t {
	result, callErr := tryNw_proxy_config_create_socksv5(proxy_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_proxy_config_enumerate_excluded_domains func(config Nw_proxy_config_t, enumerator unsafe.Pointer)
var _nw_proxy_config_enumerate_excluded_domainsErr error

func tryNw_proxy_config_enumerate_excluded_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) error {
	if _nw_proxy_config_enumerate_excluded_domains == nil {
		return symbolCallError("nw_proxy_config_enumerate_excluded_domains", "14.0", _nw_proxy_config_enumerate_excluded_domainsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string) { enumerator(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_proxy_config_enumerate_excluded_domains(config, _block0)
	return nil
}

// Nw_proxy_config_enumerate_excluded_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_excluded_domains(_:_:)
func Nw_proxy_config_enumerate_excluded_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) {
	if callErr := tryNw_proxy_config_enumerate_excluded_domains(config, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_enumerate_match_domains func(config Nw_proxy_config_t, enumerator unsafe.Pointer)
var _nw_proxy_config_enumerate_match_domainsErr error

func tryNw_proxy_config_enumerate_match_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) error {
	if _nw_proxy_config_enumerate_match_domains == nil {
		return symbolCallError("nw_proxy_config_enumerate_match_domains", "14.0", _nw_proxy_config_enumerate_match_domainsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string) { enumerator(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nw_proxy_config_enumerate_match_domains(config, _block0)
	return nil
}

// Nw_proxy_config_enumerate_match_domains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_match_domains(_:_:)
func Nw_proxy_config_enumerate_match_domains(config Nw_proxy_config_t, enumerator Nw_proxy_domain_enumerator_t) {
	if callErr := tryNw_proxy_config_enumerate_match_domains(config, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_get_failover_allowed func(proxy_config Nw_proxy_config_t) bool
var _nw_proxy_config_get_failover_allowedErr error

func tryNw_proxy_config_get_failover_allowed(proxy_config Nw_proxy_config_t) (bool, error) {
	if _nw_proxy_config_get_failover_allowed == nil {
		return false, symbolCallError("nw_proxy_config_get_failover_allowed", "14.0", _nw_proxy_config_get_failover_allowedErr)
	}
	return _nw_proxy_config_get_failover_allowed(proxy_config), nil
}

// Nw_proxy_config_get_failover_allowed checks if a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_get_failover_allowed(_:)
func Nw_proxy_config_get_failover_allowed(proxy_config Nw_proxy_config_t) bool {
	result, callErr := tryNw_proxy_config_get_failover_allowed(proxy_config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_proxy_config_set_failover_allowed func(proxy_config Nw_proxy_config_t, failover_allowed bool)
var _nw_proxy_config_set_failover_allowedErr error

func tryNw_proxy_config_set_failover_allowed(proxy_config Nw_proxy_config_t, failover_allowed bool) error {
	if _nw_proxy_config_set_failover_allowed == nil {
		return symbolCallError("nw_proxy_config_set_failover_allowed", "14.0", _nw_proxy_config_set_failover_allowedErr)
	}
	_nw_proxy_config_set_failover_allowed(proxy_config, failover_allowed)
	return nil
}

// Nw_proxy_config_set_failover_allowed configures whether or not a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_failover_allowed(_:_:)
func Nw_proxy_config_set_failover_allowed(proxy_config Nw_proxy_config_t, failover_allowed bool) {
	if callErr := tryNw_proxy_config_set_failover_allowed(proxy_config, failover_allowed); callErr != nil {
		panic(callErr)
	}
}

var _nw_proxy_config_set_username_and_password func(proxy_config Nw_proxy_config_t, username string, password string)
var _nw_proxy_config_set_username_and_passwordErr error

func tryNw_proxy_config_set_username_and_password(proxy_config Nw_proxy_config_t, username string, password string) error {
	if _nw_proxy_config_set_username_and_password == nil {
		return symbolCallError("nw_proxy_config_set_username_and_password", "14.0", _nw_proxy_config_set_username_and_passwordErr)
	}
	_nw_proxy_config_set_username_and_password(proxy_config, username, password)
	return nil
}

// Nw_proxy_config_set_username_and_password sets a username and password to use as authentication for a proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_username_and_password(_:_:_:)
func Nw_proxy_config_set_username_and_password(proxy_config Nw_proxy_config_t, username string, password string) {
	if callErr := tryNw_proxy_config_set_username_and_password(proxy_config, username, password); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_add_tls_application_protocol func(options Nw_protocol_options_t, application_protocol string)
var _nw_quic_add_tls_application_protocolErr error

func tryNw_quic_add_tls_application_protocol(options Nw_protocol_options_t, application_protocol string) error {
	if _nw_quic_add_tls_application_protocol == nil {
		return symbolCallError("nw_quic_add_tls_application_protocol", "12.0", _nw_quic_add_tls_application_protocolErr)
	}
	_nw_quic_add_tls_application_protocol(options, application_protocol)
	return nil
}

// Nw_quic_add_tls_application_protocol adds a supported Application-Layer Protocol Negotiation value.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_add_tls_application_protocol(_:_:)
func Nw_quic_add_tls_application_protocol(options Nw_protocol_options_t, application_protocol string) {
	if callErr := tryNw_quic_add_tls_application_protocol(options, application_protocol); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_copy_sec_protocol_metadata func(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t
var _nw_quic_copy_sec_protocol_metadataErr error

func tryNw_quic_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) (security.Sec_protocol_metadata_t, error) {
	if _nw_quic_copy_sec_protocol_metadata == nil {
		return *new(security.Sec_protocol_metadata_t), symbolCallError("nw_quic_copy_sec_protocol_metadata", "12.0", _nw_quic_copy_sec_protocol_metadataErr)
	}
	return _nw_quic_copy_sec_protocol_metadata(metadata), nil
}

// Nw_quic_copy_sec_protocol_metadata accesses the result of the QUIC handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_metadata(_:)
func Nw_quic_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t {
	result, callErr := tryNw_quic_copy_sec_protocol_metadata(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_copy_sec_protocol_options func(options Nw_protocol_options_t) security.Sec_protocol_options_t
var _nw_quic_copy_sec_protocol_optionsErr error

func tryNw_quic_copy_sec_protocol_options(options Nw_protocol_options_t) (security.Sec_protocol_options_t, error) {
	if _nw_quic_copy_sec_protocol_options == nil {
		return *new(security.Sec_protocol_options_t), symbolCallError("nw_quic_copy_sec_protocol_options", "12.0", _nw_quic_copy_sec_protocol_optionsErr)
	}
	return _nw_quic_copy_sec_protocol_options(options), nil
}

// Nw_quic_copy_sec_protocol_options accesses the handshake security options QUIC will use.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_options(_:)
func Nw_quic_copy_sec_protocol_options(options Nw_protocol_options_t) security.Sec_protocol_options_t {
	result, callErr := tryNw_quic_copy_sec_protocol_options(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_create_options func() Nw_protocol_options_t
var _nw_quic_create_optionsErr error

func tryNw_quic_create_options() (Nw_protocol_options_t, error) {
	if _nw_quic_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_quic_create_options", "12.0", _nw_quic_create_optionsErr)
	}
	return _nw_quic_create_options(), nil
}

// Nw_quic_create_options initializes a default set of QUIC connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_create_options()
func Nw_quic_create_options() Nw_protocol_options_t {
	result, callErr := tryNw_quic_create_options()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_application_error func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_application_errorErr error

func tryNw_quic_get_application_error(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_application_error == nil {
		return 0, symbolCallError("nw_quic_get_application_error", "12.0", _nw_quic_get_application_errorErr)
	}
	return _nw_quic_get_application_error(metadata), nil
}

// Nw_quic_get_application_error accesses the QUIC application error code received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error(_:)
func Nw_quic_get_application_error(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_application_error(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_application_error_reason func(metadata Nw_protocol_metadata_t) *byte
var _nw_quic_get_application_error_reasonErr error

func tryNw_quic_get_application_error_reason(metadata Nw_protocol_metadata_t) (*byte, error) {
	if _nw_quic_get_application_error_reason == nil {
		return nil, symbolCallError("nw_quic_get_application_error_reason", "12.0", _nw_quic_get_application_error_reasonErr)
	}
	return _nw_quic_get_application_error_reason(metadata), nil
}

// Nw_quic_get_application_error_reason accesses the QUIC application error reason received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error_reason(_:)
func Nw_quic_get_application_error_reason(metadata Nw_protocol_metadata_t) *byte {
	result, callErr := tryNw_quic_get_application_error_reason(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_idle_timeout func(options Nw_protocol_options_t) uint32
var _nw_quic_get_idle_timeoutErr error

func tryNw_quic_get_idle_timeout(options Nw_protocol_options_t) (uint32, error) {
	if _nw_quic_get_idle_timeout == nil {
		return 0, symbolCallError("nw_quic_get_idle_timeout", "12.0", _nw_quic_get_idle_timeoutErr)
	}
	return _nw_quic_get_idle_timeout(options), nil
}

// Nw_quic_get_idle_timeout accesses the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_idle_timeout(_:)
func Nw_quic_get_idle_timeout(options Nw_protocol_options_t) uint32 {
	result, callErr := tryNw_quic_get_idle_timeout(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_data func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_dataErr error

func tryNw_quic_get_initial_max_data(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_data == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_data", "12.0", _nw_quic_get_initial_max_dataErr)
	}
	return _nw_quic_get_initial_max_data(options), nil
}

// Nw_quic_get_initial_max_data accesses a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_data(_:)
func Nw_quic_get_initial_max_data(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_data(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_stream_data_bidirectional_local func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_stream_data_bidirectional_localErr error

func tryNw_quic_get_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_stream_data_bidirectional_local == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_bidirectional_local", "12.0", _nw_quic_get_initial_max_stream_data_bidirectional_localErr)
	}
	return _nw_quic_get_initial_max_stream_data_bidirectional_local(options), nil
}

// Nw_quic_get_initial_max_stream_data_bidirectional_local accesses a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_local(_:)
func Nw_quic_get_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_stream_data_bidirectional_local(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_stream_data_bidirectional_remote func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_stream_data_bidirectional_remoteErr error

func tryNw_quic_get_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_stream_data_bidirectional_remote == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_bidirectional_remote", "12.0", _nw_quic_get_initial_max_stream_data_bidirectional_remoteErr)
	}
	return _nw_quic_get_initial_max_stream_data_bidirectional_remote(options), nil
}

// Nw_quic_get_initial_max_stream_data_bidirectional_remote accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_remote(_:)
func Nw_quic_get_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_stream_data_bidirectional_remote(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_stream_data_unidirectional func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_stream_data_unidirectionalErr error

func tryNw_quic_get_initial_max_stream_data_unidirectional(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_stream_data_unidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_unidirectional", "12.0", _nw_quic_get_initial_max_stream_data_unidirectionalErr)
	}
	return _nw_quic_get_initial_max_stream_data_unidirectional(options), nil
}

// Nw_quic_get_initial_max_stream_data_unidirectional accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_unidirectional(_:)
func Nw_quic_get_initial_max_stream_data_unidirectional(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_stream_data_unidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_streams_bidirectional func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_streams_bidirectionalErr error

func tryNw_quic_get_initial_max_streams_bidirectional(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_streams_bidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_streams_bidirectional", "12.0", _nw_quic_get_initial_max_streams_bidirectionalErr)
	}
	return _nw_quic_get_initial_max_streams_bidirectional(options), nil
}

// Nw_quic_get_initial_max_streams_bidirectional accesses a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_bidirectional(_:)
func Nw_quic_get_initial_max_streams_bidirectional(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_streams_bidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_initial_max_streams_unidirectional func(options Nw_protocol_options_t) uint64
var _nw_quic_get_initial_max_streams_unidirectionalErr error

func tryNw_quic_get_initial_max_streams_unidirectional(options Nw_protocol_options_t) (uint64, error) {
	if _nw_quic_get_initial_max_streams_unidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_streams_unidirectional", "12.0", _nw_quic_get_initial_max_streams_unidirectionalErr)
	}
	return _nw_quic_get_initial_max_streams_unidirectional(options), nil
}

// Nw_quic_get_initial_max_streams_unidirectional accesses a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_unidirectional(_:)
func Nw_quic_get_initial_max_streams_unidirectional(options Nw_protocol_options_t) uint64 {
	result, callErr := tryNw_quic_get_initial_max_streams_unidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_keepalive_interval func(metadata Nw_protocol_metadata_t) uint16
var _nw_quic_get_keepalive_intervalErr error

func tryNw_quic_get_keepalive_interval(metadata Nw_protocol_metadata_t) (uint16, error) {
	if _nw_quic_get_keepalive_interval == nil {
		return 0, symbolCallError("nw_quic_get_keepalive_interval", "12.0", _nw_quic_get_keepalive_intervalErr)
	}
	return _nw_quic_get_keepalive_interval(metadata), nil
}

// Nw_quic_get_keepalive_interval accesses the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_keepalive_interval(_:)
func Nw_quic_get_keepalive_interval(metadata Nw_protocol_metadata_t) uint16 {
	result, callErr := tryNw_quic_get_keepalive_interval(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_local_max_streams_bidirectional func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_local_max_streams_bidirectionalErr error

func tryNw_quic_get_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_local_max_streams_bidirectional == nil {
		return 0, symbolCallError("nw_quic_get_local_max_streams_bidirectional", "12.0", _nw_quic_get_local_max_streams_bidirectionalErr)
	}
	return _nw_quic_get_local_max_streams_bidirectional(metadata), nil
}

// Nw_quic_get_local_max_streams_bidirectional accesses the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_bidirectional(_:)
func Nw_quic_get_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_local_max_streams_bidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_local_max_streams_unidirectional func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_local_max_streams_unidirectionalErr error

func tryNw_quic_get_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_local_max_streams_unidirectional == nil {
		return 0, symbolCallError("nw_quic_get_local_max_streams_unidirectional", "12.0", _nw_quic_get_local_max_streams_unidirectionalErr)
	}
	return _nw_quic_get_local_max_streams_unidirectional(metadata), nil
}

// Nw_quic_get_local_max_streams_unidirectional accesses the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_unidirectional(_:)
func Nw_quic_get_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_local_max_streams_unidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_max_datagram_frame_size func(options Nw_protocol_options_t) uint16
var _nw_quic_get_max_datagram_frame_sizeErr error

func tryNw_quic_get_max_datagram_frame_size(options Nw_protocol_options_t) (uint16, error) {
	if _nw_quic_get_max_datagram_frame_size == nil {
		return 0, symbolCallError("nw_quic_get_max_datagram_frame_size", "13.0", _nw_quic_get_max_datagram_frame_sizeErr)
	}
	return _nw_quic_get_max_datagram_frame_size(options), nil
}

// Nw_quic_get_max_datagram_frame_size accesses a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_datagram_frame_size(_:)
func Nw_quic_get_max_datagram_frame_size(options Nw_protocol_options_t) uint16 {
	result, callErr := tryNw_quic_get_max_datagram_frame_size(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_max_udp_payload_size func(options Nw_protocol_options_t) uint16
var _nw_quic_get_max_udp_payload_sizeErr error

func tryNw_quic_get_max_udp_payload_size(options Nw_protocol_options_t) (uint16, error) {
	if _nw_quic_get_max_udp_payload_size == nil {
		return 0, symbolCallError("nw_quic_get_max_udp_payload_size", "12.0", _nw_quic_get_max_udp_payload_sizeErr)
	}
	return _nw_quic_get_max_udp_payload_size(options), nil
}

// Nw_quic_get_max_udp_payload_size accesses the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_udp_payload_size(_:)
func Nw_quic_get_max_udp_payload_size(options Nw_protocol_options_t) uint16 {
	result, callErr := tryNw_quic_get_max_udp_payload_size(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_remote_idle_timeout func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_remote_idle_timeoutErr error

func tryNw_quic_get_remote_idle_timeout(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_remote_idle_timeout == nil {
		return 0, symbolCallError("nw_quic_get_remote_idle_timeout", "12.0", _nw_quic_get_remote_idle_timeoutErr)
	}
	return _nw_quic_get_remote_idle_timeout(metadata), nil
}

// Nw_quic_get_remote_idle_timeout accesses the idle timeout value from the peer’s transport parameters, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_idle_timeout(_:)
func Nw_quic_get_remote_idle_timeout(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_remote_idle_timeout(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_remote_max_streams_bidirectional func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_remote_max_streams_bidirectionalErr error

func tryNw_quic_get_remote_max_streams_bidirectional(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_remote_max_streams_bidirectional == nil {
		return 0, symbolCallError("nw_quic_get_remote_max_streams_bidirectional", "12.0", _nw_quic_get_remote_max_streams_bidirectionalErr)
	}
	return _nw_quic_get_remote_max_streams_bidirectional(metadata), nil
}

// Nw_quic_get_remote_max_streams_bidirectional accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_bidirectional(_:)
func Nw_quic_get_remote_max_streams_bidirectional(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_remote_max_streams_bidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_remote_max_streams_unidirectional func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_remote_max_streams_unidirectionalErr error

func tryNw_quic_get_remote_max_streams_unidirectional(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_remote_max_streams_unidirectional == nil {
		return 0, symbolCallError("nw_quic_get_remote_max_streams_unidirectional", "12.0", _nw_quic_get_remote_max_streams_unidirectionalErr)
	}
	return _nw_quic_get_remote_max_streams_unidirectional(metadata), nil
}

// Nw_quic_get_remote_max_streams_unidirectional accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_unidirectional(_:)
func Nw_quic_get_remote_max_streams_unidirectional(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_remote_max_streams_unidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_application_error func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_stream_application_errorErr error

func tryNw_quic_get_stream_application_error(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_stream_application_error == nil {
		return 0, symbolCallError("nw_quic_get_stream_application_error", "12.0", _nw_quic_get_stream_application_errorErr)
	}
	return _nw_quic_get_stream_application_error(metadata), nil
}

// Nw_quic_get_stream_application_error accesses the QUIC application error code received from the peer for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_application_error(_:)
func Nw_quic_get_stream_application_error(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_stream_application_error(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_id func(metadata Nw_protocol_metadata_t) uint64
var _nw_quic_get_stream_idErr error

func tryNw_quic_get_stream_id(metadata Nw_protocol_metadata_t) (uint64, error) {
	if _nw_quic_get_stream_id == nil {
		return 0, symbolCallError("nw_quic_get_stream_id", "12.0", _nw_quic_get_stream_idErr)
	}
	return _nw_quic_get_stream_id(metadata), nil
}

// Nw_quic_get_stream_id accesses the QUIC stream identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_id(_:)
func Nw_quic_get_stream_id(metadata Nw_protocol_metadata_t) uint64 {
	result, callErr := tryNw_quic_get_stream_id(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_is_datagram func(options Nw_protocol_options_t) bool
var _nw_quic_get_stream_is_datagramErr error

func tryNw_quic_get_stream_is_datagram(options Nw_protocol_options_t) (bool, error) {
	if _nw_quic_get_stream_is_datagram == nil {
		return false, symbolCallError("nw_quic_get_stream_is_datagram", "13.0", _nw_quic_get_stream_is_datagramErr)
	}
	return _nw_quic_get_stream_is_datagram(options), nil
}

// Nw_quic_get_stream_is_datagram checks if a QUIC stream is a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_datagram(_:)
func Nw_quic_get_stream_is_datagram(options Nw_protocol_options_t) bool {
	result, callErr := tryNw_quic_get_stream_is_datagram(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_is_unidirectional func(options Nw_protocol_options_t) bool
var _nw_quic_get_stream_is_unidirectionalErr error

func tryNw_quic_get_stream_is_unidirectional(options Nw_protocol_options_t) (bool, error) {
	if _nw_quic_get_stream_is_unidirectional == nil {
		return false, symbolCallError("nw_quic_get_stream_is_unidirectional", "12.0", _nw_quic_get_stream_is_unidirectionalErr)
	}
	return _nw_quic_get_stream_is_unidirectional(options), nil
}

// Nw_quic_get_stream_is_unidirectional checks if a QUIC stream is unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_unidirectional(_:)
func Nw_quic_get_stream_is_unidirectional(options Nw_protocol_options_t) bool {
	result, callErr := tryNw_quic_get_stream_is_unidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_type func(stream_metadata Nw_protocol_metadata_t) uint8
var _nw_quic_get_stream_typeErr error

func tryNw_quic_get_stream_type(stream_metadata Nw_protocol_metadata_t) (uint8, error) {
	if _nw_quic_get_stream_type == nil {
		return 0, symbolCallError("nw_quic_get_stream_type", "12.0", _nw_quic_get_stream_typeErr)
	}
	return _nw_quic_get_stream_type(stream_metadata), nil
}

// Nw_quic_get_stream_type accesses the stream type of the QUIC stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_type(_:)
func Nw_quic_get_stream_type(stream_metadata Nw_protocol_metadata_t) uint8 {
	result, callErr := tryNw_quic_get_stream_type(stream_metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_get_stream_usable_datagram_frame_size func(metadata Nw_protocol_metadata_t) uint16
var _nw_quic_get_stream_usable_datagram_frame_sizeErr error

func tryNw_quic_get_stream_usable_datagram_frame_size(metadata Nw_protocol_metadata_t) (uint16, error) {
	if _nw_quic_get_stream_usable_datagram_frame_size == nil {
		return 0, symbolCallError("nw_quic_get_stream_usable_datagram_frame_size", "13.0", _nw_quic_get_stream_usable_datagram_frame_sizeErr)
	}
	return _nw_quic_get_stream_usable_datagram_frame_size(metadata), nil
}

// Nw_quic_get_stream_usable_datagram_frame_size accesses the maximum usable size of a datagram frame on a QUIC datagram flow.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_usable_datagram_frame_size(_:)
func Nw_quic_get_stream_usable_datagram_frame_size(metadata Nw_protocol_metadata_t) uint16 {
	result, callErr := tryNw_quic_get_stream_usable_datagram_frame_size(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_quic_set_application_error func(metadata Nw_protocol_metadata_t, application_error uint64, reason string)
var _nw_quic_set_application_errorErr error

func tryNw_quic_set_application_error(metadata Nw_protocol_metadata_t, application_error uint64, reason string) error {
	if _nw_quic_set_application_error == nil {
		return symbolCallError("nw_quic_set_application_error", "12.0", _nw_quic_set_application_errorErr)
	}
	_nw_quic_set_application_error(metadata, application_error, reason)
	return nil
}

// Nw_quic_set_application_error sets the QUIC application error code to send for the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_application_error(_:_:_:)
func Nw_quic_set_application_error(metadata Nw_protocol_metadata_t, application_error uint64, reason string) {
	if callErr := tryNw_quic_set_application_error(metadata, application_error, reason); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_idle_timeout func(options Nw_protocol_options_t, idle_timeout uint32)
var _nw_quic_set_idle_timeoutErr error

func tryNw_quic_set_idle_timeout(options Nw_protocol_options_t, idle_timeout uint32) error {
	if _nw_quic_set_idle_timeout == nil {
		return symbolCallError("nw_quic_set_idle_timeout", "12.0", _nw_quic_set_idle_timeoutErr)
	}
	_nw_quic_set_idle_timeout(options, idle_timeout)
	return nil
}

// Nw_quic_set_idle_timeout sets the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_idle_timeout(_:_:)
func Nw_quic_set_idle_timeout(options Nw_protocol_options_t, idle_timeout uint32) {
	if callErr := tryNw_quic_set_idle_timeout(options, idle_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_data func(options Nw_protocol_options_t, initial_max_data uint64)
var _nw_quic_set_initial_max_dataErr error

func tryNw_quic_set_initial_max_data(options Nw_protocol_options_t, initial_max_data uint64) error {
	if _nw_quic_set_initial_max_data == nil {
		return symbolCallError("nw_quic_set_initial_max_data", "12.0", _nw_quic_set_initial_max_dataErr)
	}
	_nw_quic_set_initial_max_data(options, initial_max_data)
	return nil
}

// Nw_quic_set_initial_max_data sets a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_data(_:_:)
func Nw_quic_set_initial_max_data(options Nw_protocol_options_t, initial_max_data uint64) {
	if callErr := tryNw_quic_set_initial_max_data(options, initial_max_data); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_stream_data_bidirectional_local func(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_local uint64)
var _nw_quic_set_initial_max_stream_data_bidirectional_localErr error

func tryNw_quic_set_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_local uint64) error {
	if _nw_quic_set_initial_max_stream_data_bidirectional_local == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_bidirectional_local", "12.0", _nw_quic_set_initial_max_stream_data_bidirectional_localErr)
	}
	_nw_quic_set_initial_max_stream_data_bidirectional_local(options, initial_max_stream_data_bidirectional_local)
	return nil
}

// Nw_quic_set_initial_max_stream_data_bidirectional_local sets a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_local(_:_:)
func Nw_quic_set_initial_max_stream_data_bidirectional_local(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_local uint64) {
	if callErr := tryNw_quic_set_initial_max_stream_data_bidirectional_local(options, initial_max_stream_data_bidirectional_local); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_stream_data_bidirectional_remote func(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64)
var _nw_quic_set_initial_max_stream_data_bidirectional_remoteErr error

func tryNw_quic_set_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64) error {
	if _nw_quic_set_initial_max_stream_data_bidirectional_remote == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_bidirectional_remote", "12.0", _nw_quic_set_initial_max_stream_data_bidirectional_remoteErr)
	}
	_nw_quic_set_initial_max_stream_data_bidirectional_remote(options, initial_max_stream_data_bidirectional_remote)
	return nil
}

// Nw_quic_set_initial_max_stream_data_bidirectional_remote sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_remote(_:_:)
func Nw_quic_set_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64) {
	if callErr := tryNw_quic_set_initial_max_stream_data_bidirectional_remote(options, initial_max_stream_data_bidirectional_remote); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_stream_data_unidirectional func(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64)
var _nw_quic_set_initial_max_stream_data_unidirectionalErr error

func tryNw_quic_set_initial_max_stream_data_unidirectional(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64) error {
	if _nw_quic_set_initial_max_stream_data_unidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_unidirectional", "12.0", _nw_quic_set_initial_max_stream_data_unidirectionalErr)
	}
	_nw_quic_set_initial_max_stream_data_unidirectional(options, initial_max_stream_data_unidirectional)
	return nil
}

// Nw_quic_set_initial_max_stream_data_unidirectional sets a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_unidirectional(_:_:)
func Nw_quic_set_initial_max_stream_data_unidirectional(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64) {
	if callErr := tryNw_quic_set_initial_max_stream_data_unidirectional(options, initial_max_stream_data_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_streams_bidirectional func(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64)
var _nw_quic_set_initial_max_streams_bidirectionalErr error

func tryNw_quic_set_initial_max_streams_bidirectional(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64) error {
	if _nw_quic_set_initial_max_streams_bidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_streams_bidirectional", "12.0", _nw_quic_set_initial_max_streams_bidirectionalErr)
	}
	_nw_quic_set_initial_max_streams_bidirectional(options, initial_max_streams_bidirectional)
	return nil
}

// Nw_quic_set_initial_max_streams_bidirectional sets a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_bidirectional(_:_:)
func Nw_quic_set_initial_max_streams_bidirectional(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64) {
	if callErr := tryNw_quic_set_initial_max_streams_bidirectional(options, initial_max_streams_bidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_initial_max_streams_unidirectional func(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64)
var _nw_quic_set_initial_max_streams_unidirectionalErr error

func tryNw_quic_set_initial_max_streams_unidirectional(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64) error {
	if _nw_quic_set_initial_max_streams_unidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_streams_unidirectional", "12.0", _nw_quic_set_initial_max_streams_unidirectionalErr)
	}
	_nw_quic_set_initial_max_streams_unidirectional(options, initial_max_streams_unidirectional)
	return nil
}

// Nw_quic_set_initial_max_streams_unidirectional sets a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_unidirectional(_:_:)
func Nw_quic_set_initial_max_streams_unidirectional(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64) {
	if callErr := tryNw_quic_set_initial_max_streams_unidirectional(options, initial_max_streams_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_keepalive_interval func(metadata Nw_protocol_metadata_t, keepalive_interval uint16)
var _nw_quic_set_keepalive_intervalErr error

func tryNw_quic_set_keepalive_interval(metadata Nw_protocol_metadata_t, keepalive_interval uint16) error {
	if _nw_quic_set_keepalive_interval == nil {
		return symbolCallError("nw_quic_set_keepalive_interval", "12.0", _nw_quic_set_keepalive_intervalErr)
	}
	_nw_quic_set_keepalive_interval(metadata, keepalive_interval)
	return nil
}

// Nw_quic_set_keepalive_interval sets the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_keepalive_interval(_:_:)
func Nw_quic_set_keepalive_interval(metadata Nw_protocol_metadata_t, keepalive_interval uint16) {
	if callErr := tryNw_quic_set_keepalive_interval(metadata, keepalive_interval); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_local_max_streams_bidirectional func(metadata Nw_protocol_metadata_t, max_streams_bidirectional uint64)
var _nw_quic_set_local_max_streams_bidirectionalErr error

func tryNw_quic_set_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t, max_streams_bidirectional uint64) error {
	if _nw_quic_set_local_max_streams_bidirectional == nil {
		return symbolCallError("nw_quic_set_local_max_streams_bidirectional", "12.0", _nw_quic_set_local_max_streams_bidirectionalErr)
	}
	_nw_quic_set_local_max_streams_bidirectional(metadata, max_streams_bidirectional)
	return nil
}

// Nw_quic_set_local_max_streams_bidirectional sets the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_bidirectional(_:_:)
func Nw_quic_set_local_max_streams_bidirectional(metadata Nw_protocol_metadata_t, max_streams_bidirectional uint64) {
	if callErr := tryNw_quic_set_local_max_streams_bidirectional(metadata, max_streams_bidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_local_max_streams_unidirectional func(metadata Nw_protocol_metadata_t, max_streams_unidirectional uint64)
var _nw_quic_set_local_max_streams_unidirectionalErr error

func tryNw_quic_set_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t, max_streams_unidirectional uint64) error {
	if _nw_quic_set_local_max_streams_unidirectional == nil {
		return symbolCallError("nw_quic_set_local_max_streams_unidirectional", "12.0", _nw_quic_set_local_max_streams_unidirectionalErr)
	}
	_nw_quic_set_local_max_streams_unidirectional(metadata, max_streams_unidirectional)
	return nil
}

// Nw_quic_set_local_max_streams_unidirectional sets the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_unidirectional(_:_:)
func Nw_quic_set_local_max_streams_unidirectional(metadata Nw_protocol_metadata_t, max_streams_unidirectional uint64) {
	if callErr := tryNw_quic_set_local_max_streams_unidirectional(metadata, max_streams_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_max_datagram_frame_size func(options Nw_protocol_options_t, max_datagram_frame_size uint16)
var _nw_quic_set_max_datagram_frame_sizeErr error

func tryNw_quic_set_max_datagram_frame_size(options Nw_protocol_options_t, max_datagram_frame_size uint16) error {
	if _nw_quic_set_max_datagram_frame_size == nil {
		return symbolCallError("nw_quic_set_max_datagram_frame_size", "13.0", _nw_quic_set_max_datagram_frame_sizeErr)
	}
	_nw_quic_set_max_datagram_frame_size(options, max_datagram_frame_size)
	return nil
}

// Nw_quic_set_max_datagram_frame_size sets a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_datagram_frame_size(_:_:)
func Nw_quic_set_max_datagram_frame_size(options Nw_protocol_options_t, max_datagram_frame_size uint16) {
	if callErr := tryNw_quic_set_max_datagram_frame_size(options, max_datagram_frame_size); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_max_udp_payload_size func(options Nw_protocol_options_t, max_udp_payload_size uint16)
var _nw_quic_set_max_udp_payload_sizeErr error

func tryNw_quic_set_max_udp_payload_size(options Nw_protocol_options_t, max_udp_payload_size uint16) error {
	if _nw_quic_set_max_udp_payload_size == nil {
		return symbolCallError("nw_quic_set_max_udp_payload_size", "12.0", _nw_quic_set_max_udp_payload_sizeErr)
	}
	_nw_quic_set_max_udp_payload_size(options, max_udp_payload_size)
	return nil
}

// Nw_quic_set_max_udp_payload_size sets the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_udp_payload_size(_:_:)
func Nw_quic_set_max_udp_payload_size(options Nw_protocol_options_t, max_udp_payload_size uint16) {
	if callErr := tryNw_quic_set_max_udp_payload_size(options, max_udp_payload_size); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_stream_application_error func(metadata Nw_protocol_metadata_t, application_error uint64)
var _nw_quic_set_stream_application_errorErr error

func tryNw_quic_set_stream_application_error(metadata Nw_protocol_metadata_t, application_error uint64) error {
	if _nw_quic_set_stream_application_error == nil {
		return symbolCallError("nw_quic_set_stream_application_error", "12.0", _nw_quic_set_stream_application_errorErr)
	}
	_nw_quic_set_stream_application_error(metadata, application_error)
	return nil
}

// Nw_quic_set_stream_application_error sets the QUIC application error code to send for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_application_error(_:_:)
func Nw_quic_set_stream_application_error(metadata Nw_protocol_metadata_t, application_error uint64) {
	if callErr := tryNw_quic_set_stream_application_error(metadata, application_error); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_stream_is_datagram func(options Nw_protocol_options_t, is_datagram bool)
var _nw_quic_set_stream_is_datagramErr error

func tryNw_quic_set_stream_is_datagram(options Nw_protocol_options_t, is_datagram bool) error {
	if _nw_quic_set_stream_is_datagram == nil {
		return symbolCallError("nw_quic_set_stream_is_datagram", "13.0", _nw_quic_set_stream_is_datagramErr)
	}
	_nw_quic_set_stream_is_datagram(options, is_datagram)
	return nil
}

// Nw_quic_set_stream_is_datagram configures a QUIC stream as a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_datagram(_:_:)
func Nw_quic_set_stream_is_datagram(options Nw_protocol_options_t, is_datagram bool) {
	if callErr := tryNw_quic_set_stream_is_datagram(options, is_datagram); callErr != nil {
		panic(callErr)
	}
}

var _nw_quic_set_stream_is_unidirectional func(options Nw_protocol_options_t, is_unidirectional bool)
var _nw_quic_set_stream_is_unidirectionalErr error

func tryNw_quic_set_stream_is_unidirectional(options Nw_protocol_options_t, is_unidirectional bool) error {
	if _nw_quic_set_stream_is_unidirectional == nil {
		return symbolCallError("nw_quic_set_stream_is_unidirectional", "12.0", _nw_quic_set_stream_is_unidirectionalErr)
	}
	_nw_quic_set_stream_is_unidirectional(options, is_unidirectional)
	return nil
}

// Nw_quic_set_stream_is_unidirectional configures a QUIC stream as unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_unidirectional(_:_:)
func Nw_quic_set_stream_is_unidirectional(options Nw_protocol_options_t, is_unidirectional bool) {
	if callErr := tryNw_quic_set_stream_is_unidirectional(options, is_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nw_relay_hop_add_additional_http_header_field func(relay_hop Nw_relay_hop_t, field_name string, field_value string)
var _nw_relay_hop_add_additional_http_header_fieldErr error

func tryNw_relay_hop_add_additional_http_header_field(relay_hop Nw_relay_hop_t, field_name string, field_value string) error {
	if _nw_relay_hop_add_additional_http_header_field == nil {
		return symbolCallError("nw_relay_hop_add_additional_http_header_field", "14.0", _nw_relay_hop_add_additional_http_header_fieldErr)
	}
	_nw_relay_hop_add_additional_http_header_field(relay_hop, field_name, field_value)
	return nil
}

// Nw_relay_hop_add_additional_http_header_field adds an HTTP header name and value pair to send as part of [CONNECT] requests to the relay.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_add_additional_http_header_field(_:_:_:)
func Nw_relay_hop_add_additional_http_header_field(relay_hop Nw_relay_hop_t, field_name string, field_value string) {
	if callErr := tryNw_relay_hop_add_additional_http_header_field(relay_hop, field_name, field_value); callErr != nil {
		panic(callErr)
	}
}

var _nw_relay_hop_create func(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) Nw_relay_hop_t
var _nw_relay_hop_createErr error

func tryNw_relay_hop_create(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) (Nw_relay_hop_t, error) {
	if _nw_relay_hop_create == nil {
		return *new(Nw_relay_hop_t), symbolCallError("nw_relay_hop_create", "14.0", _nw_relay_hop_createErr)
	}
	return _nw_relay_hop_create(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options), nil
}

// Nw_relay_hop_create creates a configuration for a secure relay accessible using HTTP/3, with an optional HTTP/2 fallback.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_create(_:_:_:)
func Nw_relay_hop_create(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) Nw_relay_hop_t {
	result, callErr := tryNw_relay_hop_create(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_release func(obj unsafe.Pointer)
var _nw_releaseErr error

func tryNw_release(obj unsafe.Pointer) error {
	if _nw_release == nil {
		return symbolCallError("nw_release", "10.14", _nw_releaseErr)
	}
	_nw_release(obj)
	return nil
}

// Nw_release releases a reference count on a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_release
func Nw_release(obj unsafe.Pointer) {
	if callErr := tryNw_release(obj); callErr != nil {
		panic(callErr)
	}
}

var _nw_resolution_report_copy_preferred_endpoint func(resolution_report Nw_resolution_report_t) Nw_endpoint_t
var _nw_resolution_report_copy_preferred_endpointErr error

func tryNw_resolution_report_copy_preferred_endpoint(resolution_report Nw_resolution_report_t) (Nw_endpoint_t, error) {
	if _nw_resolution_report_copy_preferred_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_resolution_report_copy_preferred_endpoint", "11.0", _nw_resolution_report_copy_preferred_endpointErr)
	}
	return _nw_resolution_report_copy_preferred_endpoint(resolution_report), nil
}

// Nw_resolution_report_copy_preferred_endpoint accesses the resolved endpoint that the connection used for its first connection attempt.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_preferred_endpoint(_:)
func Nw_resolution_report_copy_preferred_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	result, callErr := tryNw_resolution_report_copy_preferred_endpoint(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolution_report_copy_successful_endpoint func(resolution_report Nw_resolution_report_t) Nw_endpoint_t
var _nw_resolution_report_copy_successful_endpointErr error

func tryNw_resolution_report_copy_successful_endpoint(resolution_report Nw_resolution_report_t) (Nw_endpoint_t, error) {
	if _nw_resolution_report_copy_successful_endpoint == nil {
		return *new(Nw_endpoint_t), symbolCallError("nw_resolution_report_copy_successful_endpoint", "11.0", _nw_resolution_report_copy_successful_endpointErr)
	}
	return _nw_resolution_report_copy_successful_endpoint(resolution_report), nil
}

// Nw_resolution_report_copy_successful_endpoint accesses the resolved endpoint that led to the established connection.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_successful_endpoint(_:)
func Nw_resolution_report_copy_successful_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	result, callErr := tryNw_resolution_report_copy_successful_endpoint(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolution_report_get_endpoint_count func(resolution_report Nw_resolution_report_t) uint32
var _nw_resolution_report_get_endpoint_countErr error

func tryNw_resolution_report_get_endpoint_count(resolution_report Nw_resolution_report_t) (uint32, error) {
	if _nw_resolution_report_get_endpoint_count == nil {
		return 0, symbolCallError("nw_resolution_report_get_endpoint_count", "11.0", _nw_resolution_report_get_endpoint_countErr)
	}
	return _nw_resolution_report_get_endpoint_count(resolution_report), nil
}

// Nw_resolution_report_get_endpoint_count accesses the number of endpoints resolved in this step.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_endpoint_count(_:)
func Nw_resolution_report_get_endpoint_count(resolution_report Nw_resolution_report_t) uint32 {
	result, callErr := tryNw_resolution_report_get_endpoint_count(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolution_report_get_milliseconds func(resolution_report Nw_resolution_report_t) uint64
var _nw_resolution_report_get_millisecondsErr error

func tryNw_resolution_report_get_milliseconds(resolution_report Nw_resolution_report_t) (uint64, error) {
	if _nw_resolution_report_get_milliseconds == nil {
		return 0, symbolCallError("nw_resolution_report_get_milliseconds", "11.0", _nw_resolution_report_get_millisecondsErr)
	}
	return _nw_resolution_report_get_milliseconds(resolution_report), nil
}

// Nw_resolution_report_get_milliseconds accesses the duration of this resolution step, from when the query was issued to when the response was complete.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_milliseconds(_:)
func Nw_resolution_report_get_milliseconds(resolution_report Nw_resolution_report_t) uint64 {
	result, callErr := tryNw_resolution_report_get_milliseconds(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolution_report_get_protocol func(resolution_report Nw_resolution_report_t) NwReportResolutionProtocol
var _nw_resolution_report_get_protocolErr error

func tryNw_resolution_report_get_protocol(resolution_report Nw_resolution_report_t) (NwReportResolutionProtocol, error) {
	if _nw_resolution_report_get_protocol == nil {
		return *new(NwReportResolutionProtocol), symbolCallError("nw_resolution_report_get_protocol", "11.0", _nw_resolution_report_get_protocolErr)
	}
	return _nw_resolution_report_get_protocol(resolution_report), nil
}

// Nw_resolution_report_get_protocol accesses the transport protocol your connection used for DNS resolution.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_protocol(_:)
func Nw_resolution_report_get_protocol(resolution_report Nw_resolution_report_t) NwReportResolutionProtocol {
	result, callErr := tryNw_resolution_report_get_protocol(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolution_report_get_source func(resolution_report Nw_resolution_report_t) NwReportResolutionSource
var _nw_resolution_report_get_sourceErr error

func tryNw_resolution_report_get_source(resolution_report Nw_resolution_report_t) (NwReportResolutionSource, error) {
	if _nw_resolution_report_get_source == nil {
		return *new(NwReportResolutionSource), symbolCallError("nw_resolution_report_get_source", "11.0", _nw_resolution_report_get_sourceErr)
	}
	return _nw_resolution_report_get_source(resolution_report), nil
}

// Nw_resolution_report_get_source accesses the source of the DNS response.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_source(_:)
func Nw_resolution_report_get_source(resolution_report Nw_resolution_report_t) NwReportResolutionSource {
	result, callErr := tryNw_resolution_report_get_source(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolver_config_add_server_address func(config Nw_resolver_config_t, server_address Nw_endpoint_t)
var _nw_resolver_config_add_server_addressErr error

func tryNw_resolver_config_add_server_address(config Nw_resolver_config_t, server_address Nw_endpoint_t) error {
	if _nw_resolver_config_add_server_address == nil {
		return symbolCallError("nw_resolver_config_add_server_address", "11.0", _nw_resolver_config_add_server_addressErr)
	}
	_nw_resolver_config_add_server_address(config, server_address)
	return nil
}

// Nw_resolver_config_add_server_address provides a well-known DNS server address to use instead of looking up the address dynamically.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_add_server_address(_:_:)
func Nw_resolver_config_add_server_address(config Nw_resolver_config_t, server_address Nw_endpoint_t) {
	if callErr := tryNw_resolver_config_add_server_address(config, server_address); callErr != nil {
		panic(callErr)
	}
}

var _nw_resolver_config_create_https func(url_endpoint Nw_endpoint_t) Nw_resolver_config_t
var _nw_resolver_config_create_httpsErr error

func tryNw_resolver_config_create_https(url_endpoint Nw_endpoint_t) (Nw_resolver_config_t, error) {
	if _nw_resolver_config_create_https == nil {
		return *new(Nw_resolver_config_t), symbolCallError("nw_resolver_config_create_https", "11.0", _nw_resolver_config_create_httpsErr)
	}
	return _nw_resolver_config_create_https(url_endpoint), nil
}

// Nw_resolver_config_create_https initializes a DNS-over-HTTPS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_https(_:)
func Nw_resolver_config_create_https(url_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	result, callErr := tryNw_resolver_config_create_https(url_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_resolver_config_create_tls func(server_endpoint Nw_endpoint_t) Nw_resolver_config_t
var _nw_resolver_config_create_tlsErr error

func tryNw_resolver_config_create_tls(server_endpoint Nw_endpoint_t) (Nw_resolver_config_t, error) {
	if _nw_resolver_config_create_tls == nil {
		return *new(Nw_resolver_config_t), symbolCallError("nw_resolver_config_create_tls", "11.0", _nw_resolver_config_create_tlsErr)
	}
	return _nw_resolver_config_create_tls(server_endpoint), nil
}

// Nw_resolver_config_create_tls initializes a DNS-over-TLS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_tls(_:)
func Nw_resolver_config_create_tls(server_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	result, callErr := tryNw_resolver_config_create_tls(server_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_retain func(obj unsafe.Pointer) unsafe.Pointer
var _nw_retainErr error

func tryNw_retain(obj unsafe.Pointer) (unsafe.Pointer, error) {
	if _nw_retain == nil {
		return nil, symbolCallError("nw_retain", "10.14", _nw_retainErr)
	}
	return _nw_retain(obj), nil
}

// Nw_retain adds a reference count to a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_retain
func Nw_retain(obj unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNw_retain(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tcp_create_options func() Nw_protocol_options_t
var _nw_tcp_create_optionsErr error

func tryNw_tcp_create_options() (Nw_protocol_options_t, error) {
	if _nw_tcp_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_tcp_create_options", "10.14", _nw_tcp_create_optionsErr)
	}
	return _nw_tcp_create_options(), nil
}

// Nw_tcp_create_options initializes a default set of TCP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_create_options()
func Nw_tcp_create_options() Nw_protocol_options_t {
	result, callErr := tryNw_tcp_create_options()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tcp_get_available_receive_buffer func(metadata Nw_protocol_metadata_t) uint32
var _nw_tcp_get_available_receive_bufferErr error

func tryNw_tcp_get_available_receive_buffer(metadata Nw_protocol_metadata_t) (uint32, error) {
	if _nw_tcp_get_available_receive_buffer == nil {
		return 0, symbolCallError("nw_tcp_get_available_receive_buffer", "10.14", _nw_tcp_get_available_receive_bufferErr)
	}
	return _nw_tcp_get_available_receive_buffer(metadata), nil
}

// Nw_tcp_get_available_receive_buffer accesses the number of available bytes in the TCP receive buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_receive_buffer(_:)
func Nw_tcp_get_available_receive_buffer(metadata Nw_protocol_metadata_t) uint32 {
	result, callErr := tryNw_tcp_get_available_receive_buffer(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tcp_get_available_send_buffer func(metadata Nw_protocol_metadata_t) uint32
var _nw_tcp_get_available_send_bufferErr error

func tryNw_tcp_get_available_send_buffer(metadata Nw_protocol_metadata_t) (uint32, error) {
	if _nw_tcp_get_available_send_buffer == nil {
		return 0, symbolCallError("nw_tcp_get_available_send_buffer", "10.14", _nw_tcp_get_available_send_bufferErr)
	}
	return _nw_tcp_get_available_send_buffer(metadata), nil
}

// Nw_tcp_get_available_send_buffer accesses the number of available bytes in the TCP send buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_send_buffer(_:)
func Nw_tcp_get_available_send_buffer(metadata Nw_protocol_metadata_t) uint32 {
	result, callErr := tryNw_tcp_get_available_send_buffer(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tcp_options_set_connection_timeout func(options Nw_protocol_options_t, connection_timeout uint32)
var _nw_tcp_options_set_connection_timeoutErr error

func tryNw_tcp_options_set_connection_timeout(options Nw_protocol_options_t, connection_timeout uint32) error {
	if _nw_tcp_options_set_connection_timeout == nil {
		return symbolCallError("nw_tcp_options_set_connection_timeout", "10.14", _nw_tcp_options_set_connection_timeoutErr)
	}
	_nw_tcp_options_set_connection_timeout(options, connection_timeout)
	return nil
}

// Nw_tcp_options_set_connection_timeout sets the number of seconds that TCP waits before timing out its handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_connection_timeout(_:_:)
func Nw_tcp_options_set_connection_timeout(options Nw_protocol_options_t, connection_timeout uint32) {
	if callErr := tryNw_tcp_options_set_connection_timeout(options, connection_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_disable_ack_stretching func(options Nw_protocol_options_t, disable_ack_stretching bool)
var _nw_tcp_options_set_disable_ack_stretchingErr error

func tryNw_tcp_options_set_disable_ack_stretching(options Nw_protocol_options_t, disable_ack_stretching bool) error {
	if _nw_tcp_options_set_disable_ack_stretching == nil {
		return symbolCallError("nw_tcp_options_set_disable_ack_stretching", "10.14", _nw_tcp_options_set_disable_ack_stretchingErr)
	}
	_nw_tcp_options_set_disable_ack_stretching(options, disable_ack_stretching)
	return nil
}

// Nw_tcp_options_set_disable_ack_stretching disables TCP acknowledgment stretching.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ack_stretching(_:_:)
func Nw_tcp_options_set_disable_ack_stretching(options Nw_protocol_options_t, disable_ack_stretching bool) {
	if callErr := tryNw_tcp_options_set_disable_ack_stretching(options, disable_ack_stretching); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_disable_ecn func(options Nw_protocol_options_t, disable_ecn bool)
var _nw_tcp_options_set_disable_ecnErr error

func tryNw_tcp_options_set_disable_ecn(options Nw_protocol_options_t, disable_ecn bool) error {
	if _nw_tcp_options_set_disable_ecn == nil {
		return symbolCallError("nw_tcp_options_set_disable_ecn", "10.14", _nw_tcp_options_set_disable_ecnErr)
	}
	_nw_tcp_options_set_disable_ecn(options, disable_ecn)
	return nil
}

// Nw_tcp_options_set_disable_ecn disables negotiation of Explicit Congestion Notification markings.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ecn(_:_:)
func Nw_tcp_options_set_disable_ecn(options Nw_protocol_options_t, disable_ecn bool) {
	if callErr := tryNw_tcp_options_set_disable_ecn(options, disable_ecn); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_enable_fast_open func(options Nw_protocol_options_t, enable_fast_open bool)
var _nw_tcp_options_set_enable_fast_openErr error

func tryNw_tcp_options_set_enable_fast_open(options Nw_protocol_options_t, enable_fast_open bool) error {
	if _nw_tcp_options_set_enable_fast_open == nil {
		return symbolCallError("nw_tcp_options_set_enable_fast_open", "10.14", _nw_tcp_options_set_enable_fast_openErr)
	}
	_nw_tcp_options_set_enable_fast_open(options, enable_fast_open)
	return nil
}

// Nw_tcp_options_set_enable_fast_open enables TCP Fast Open on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_fast_open(_:_:)
func Nw_tcp_options_set_enable_fast_open(options Nw_protocol_options_t, enable_fast_open bool) {
	if callErr := tryNw_tcp_options_set_enable_fast_open(options, enable_fast_open); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_enable_keepalive func(options Nw_protocol_options_t, enable_keepalive bool)
var _nw_tcp_options_set_enable_keepaliveErr error

func tryNw_tcp_options_set_enable_keepalive(options Nw_protocol_options_t, enable_keepalive bool) error {
	if _nw_tcp_options_set_enable_keepalive == nil {
		return symbolCallError("nw_tcp_options_set_enable_keepalive", "10.14", _nw_tcp_options_set_enable_keepaliveErr)
	}
	_nw_tcp_options_set_enable_keepalive(options, enable_keepalive)
	return nil
}

// Nw_tcp_options_set_enable_keepalive enables TCP keepalives.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_keepalive(_:_:)
func Nw_tcp_options_set_enable_keepalive(options Nw_protocol_options_t, enable_keepalive bool) {
	if callErr := tryNw_tcp_options_set_enable_keepalive(options, enable_keepalive); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_keepalive_count func(options Nw_protocol_options_t, keepalive_count uint32)
var _nw_tcp_options_set_keepalive_countErr error

func tryNw_tcp_options_set_keepalive_count(options Nw_protocol_options_t, keepalive_count uint32) error {
	if _nw_tcp_options_set_keepalive_count == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_count", "10.14", _nw_tcp_options_set_keepalive_countErr)
	}
	_nw_tcp_options_set_keepalive_count(options, keepalive_count)
	return nil
}

// Nw_tcp_options_set_keepalive_count sets the number of keepalive probes that TCP sends before terminating the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_count(_:_:)
func Nw_tcp_options_set_keepalive_count(options Nw_protocol_options_t, keepalive_count uint32) {
	if callErr := tryNw_tcp_options_set_keepalive_count(options, keepalive_count); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_keepalive_idle_time func(options Nw_protocol_options_t, keepalive_idle_time uint32)
var _nw_tcp_options_set_keepalive_idle_timeErr error

func tryNw_tcp_options_set_keepalive_idle_time(options Nw_protocol_options_t, keepalive_idle_time uint32) error {
	if _nw_tcp_options_set_keepalive_idle_time == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_idle_time", "10.14", _nw_tcp_options_set_keepalive_idle_timeErr)
	}
	_nw_tcp_options_set_keepalive_idle_time(options, keepalive_idle_time)
	return nil
}

// Nw_tcp_options_set_keepalive_idle_time sets the number of seconds of idleness that TCP waits before sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_idle_time(_:_:)
func Nw_tcp_options_set_keepalive_idle_time(options Nw_protocol_options_t, keepalive_idle_time uint32) {
	if callErr := tryNw_tcp_options_set_keepalive_idle_time(options, keepalive_idle_time); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_keepalive_interval func(options Nw_protocol_options_t, keepalive_interval uint32)
var _nw_tcp_options_set_keepalive_intervalErr error

func tryNw_tcp_options_set_keepalive_interval(options Nw_protocol_options_t, keepalive_interval uint32) error {
	if _nw_tcp_options_set_keepalive_interval == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_interval", "10.14", _nw_tcp_options_set_keepalive_intervalErr)
	}
	_nw_tcp_options_set_keepalive_interval(options, keepalive_interval)
	return nil
}

// Nw_tcp_options_set_keepalive_interval sets the number of seconds that TCP waits between sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_interval(_:_:)
func Nw_tcp_options_set_keepalive_interval(options Nw_protocol_options_t, keepalive_interval uint32) {
	if callErr := tryNw_tcp_options_set_keepalive_interval(options, keepalive_interval); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_maximum_segment_size func(options Nw_protocol_options_t, maximum_segment_size uint32)
var _nw_tcp_options_set_maximum_segment_sizeErr error

func tryNw_tcp_options_set_maximum_segment_size(options Nw_protocol_options_t, maximum_segment_size uint32) error {
	if _nw_tcp_options_set_maximum_segment_size == nil {
		return symbolCallError("nw_tcp_options_set_maximum_segment_size", "10.14", _nw_tcp_options_set_maximum_segment_sizeErr)
	}
	_nw_tcp_options_set_maximum_segment_size(options, maximum_segment_size)
	return nil
}

// Nw_tcp_options_set_maximum_segment_size sets TCP’s maximum segment size in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_maximum_segment_size(_:_:)
func Nw_tcp_options_set_maximum_segment_size(options Nw_protocol_options_t, maximum_segment_size uint32) {
	if callErr := tryNw_tcp_options_set_maximum_segment_size(options, maximum_segment_size); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_multipath_force_version func(options Nw_protocol_options_t, multipath_force_version NwMultipathVersion)
var _nw_tcp_options_set_multipath_force_versionErr error

func tryNw_tcp_options_set_multipath_force_version(options Nw_protocol_options_t, multipath_force_version NwMultipathVersion) error {
	if _nw_tcp_options_set_multipath_force_version == nil {
		return symbolCallError("nw_tcp_options_set_multipath_force_version", "12.0", _nw_tcp_options_set_multipath_force_versionErr)
	}
	_nw_tcp_options_set_multipath_force_version(options, multipath_force_version)
	return nil
}

// Nw_tcp_options_set_multipath_force_version.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_multipath_force_version(_:_:)
func Nw_tcp_options_set_multipath_force_version(options Nw_protocol_options_t, multipath_force_version NwMultipathVersion) {
	if callErr := tryNw_tcp_options_set_multipath_force_version(options, multipath_force_version); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_no_delay func(options Nw_protocol_options_t, no_delay bool)
var _nw_tcp_options_set_no_delayErr error

func tryNw_tcp_options_set_no_delay(options Nw_protocol_options_t, no_delay bool) error {
	if _nw_tcp_options_set_no_delay == nil {
		return symbolCallError("nw_tcp_options_set_no_delay", "10.14", _nw_tcp_options_set_no_delayErr)
	}
	_nw_tcp_options_set_no_delay(options, no_delay)
	return nil
}

// Nw_tcp_options_set_no_delay disables Nagle’s algorithm for TCP.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_delay(_:_:)
func Nw_tcp_options_set_no_delay(options Nw_protocol_options_t, no_delay bool) {
	if callErr := tryNw_tcp_options_set_no_delay(options, no_delay); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_no_options func(options Nw_protocol_options_t, no_options bool)
var _nw_tcp_options_set_no_optionsErr error

func tryNw_tcp_options_set_no_options(options Nw_protocol_options_t, no_options bool) error {
	if _nw_tcp_options_set_no_options == nil {
		return symbolCallError("nw_tcp_options_set_no_options", "10.14", _nw_tcp_options_set_no_optionsErr)
	}
	_nw_tcp_options_set_no_options(options, no_options)
	return nil
}

// Nw_tcp_options_set_no_options sets TCP into no-options mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_options(_:_:)
func Nw_tcp_options_set_no_options(options Nw_protocol_options_t, no_options bool) {
	if callErr := tryNw_tcp_options_set_no_options(options, no_options); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_no_push func(options Nw_protocol_options_t, no_push bool)
var _nw_tcp_options_set_no_pushErr error

func tryNw_tcp_options_set_no_push(options Nw_protocol_options_t, no_push bool) error {
	if _nw_tcp_options_set_no_push == nil {
		return symbolCallError("nw_tcp_options_set_no_push", "10.14", _nw_tcp_options_set_no_pushErr)
	}
	_nw_tcp_options_set_no_push(options, no_push)
	return nil
}

// Nw_tcp_options_set_no_push sets TCP into no-push mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_push(_:_:)
func Nw_tcp_options_set_no_push(options Nw_protocol_options_t, no_push bool) {
	if callErr := tryNw_tcp_options_set_no_push(options, no_push); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_persist_timeout func(options Nw_protocol_options_t, persist_timeout uint32)
var _nw_tcp_options_set_persist_timeoutErr error

func tryNw_tcp_options_set_persist_timeout(options Nw_protocol_options_t, persist_timeout uint32) error {
	if _nw_tcp_options_set_persist_timeout == nil {
		return symbolCallError("nw_tcp_options_set_persist_timeout", "10.14", _nw_tcp_options_set_persist_timeoutErr)
	}
	_nw_tcp_options_set_persist_timeout(options, persist_timeout)
	return nil
}

// Nw_tcp_options_set_persist_timeout sets the TCP persist timeout in seconds, as defined by RFC 6429.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_persist_timeout(_:_:)
func Nw_tcp_options_set_persist_timeout(options Nw_protocol_options_t, persist_timeout uint32) {
	if callErr := tryNw_tcp_options_set_persist_timeout(options, persist_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_retransmit_connection_drop_time func(options Nw_protocol_options_t, retransmit_connection_drop_time uint32)
var _nw_tcp_options_set_retransmit_connection_drop_timeErr error

func tryNw_tcp_options_set_retransmit_connection_drop_time(options Nw_protocol_options_t, retransmit_connection_drop_time uint32) error {
	if _nw_tcp_options_set_retransmit_connection_drop_time == nil {
		return symbolCallError("nw_tcp_options_set_retransmit_connection_drop_time", "10.14", _nw_tcp_options_set_retransmit_connection_drop_timeErr)
	}
	_nw_tcp_options_set_retransmit_connection_drop_time(options, retransmit_connection_drop_time)
	return nil
}

// Nw_tcp_options_set_retransmit_connection_drop_time sets the number of seconds that TCP waits between retransmission attempts.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_connection_drop_time(_:_:)
func Nw_tcp_options_set_retransmit_connection_drop_time(options Nw_protocol_options_t, retransmit_connection_drop_time uint32) {
	if callErr := tryNw_tcp_options_set_retransmit_connection_drop_time(options, retransmit_connection_drop_time); callErr != nil {
		panic(callErr)
	}
}

var _nw_tcp_options_set_retransmit_fin_drop func(options Nw_protocol_options_t, retransmit_fin_drop bool)
var _nw_tcp_options_set_retransmit_fin_dropErr error

func tryNw_tcp_options_set_retransmit_fin_drop(options Nw_protocol_options_t, retransmit_fin_drop bool) error {
	if _nw_tcp_options_set_retransmit_fin_drop == nil {
		return symbolCallError("nw_tcp_options_set_retransmit_fin_drop", "10.14", _nw_tcp_options_set_retransmit_fin_dropErr)
	}
	_nw_tcp_options_set_retransmit_fin_drop(options, retransmit_fin_drop)
	return nil
}

// Nw_tcp_options_set_retransmit_fin_drop causes TCP to drop its connection after not receiving an ACK after a FIN.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_fin_drop(_:_:)
func Nw_tcp_options_set_retransmit_fin_drop(options Nw_protocol_options_t, retransmit_fin_drop bool) {
	if callErr := tryNw_tcp_options_set_retransmit_fin_drop(options, retransmit_fin_drop); callErr != nil {
		panic(callErr)
	}
}

var _nw_tls_copy_sec_protocol_metadata func(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t
var _nw_tls_copy_sec_protocol_metadataErr error

func tryNw_tls_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) (security.Sec_protocol_metadata_t, error) {
	if _nw_tls_copy_sec_protocol_metadata == nil {
		return *new(security.Sec_protocol_metadata_t), symbolCallError("nw_tls_copy_sec_protocol_metadata", "10.14", _nw_tls_copy_sec_protocol_metadataErr)
	}
	return _nw_tls_copy_sec_protocol_metadata(metadata), nil
}

// Nw_tls_copy_sec_protocol_metadata accesses the result of the TLS handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_metadata(_:)
func Nw_tls_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) security.Sec_protocol_metadata_t {
	result, callErr := tryNw_tls_copy_sec_protocol_metadata(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tls_copy_sec_protocol_options func(options Nw_protocol_options_t) security.Sec_protocol_options_t
var _nw_tls_copy_sec_protocol_optionsErr error

func tryNw_tls_copy_sec_protocol_options(options Nw_protocol_options_t) (security.Sec_protocol_options_t, error) {
	if _nw_tls_copy_sec_protocol_options == nil {
		return *new(security.Sec_protocol_options_t), symbolCallError("nw_tls_copy_sec_protocol_options", "10.14", _nw_tls_copy_sec_protocol_optionsErr)
	}
	return _nw_tls_copy_sec_protocol_options(options), nil
}

// Nw_tls_copy_sec_protocol_options accesses the handshake security options TLS will use.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_options(_:)
func Nw_tls_copy_sec_protocol_options(options Nw_protocol_options_t) security.Sec_protocol_options_t {
	result, callErr := tryNw_tls_copy_sec_protocol_options(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_tls_create_options func() Nw_protocol_options_t
var _nw_tls_create_optionsErr error

func tryNw_tls_create_options() (Nw_protocol_options_t, error) {
	if _nw_tls_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_tls_create_options", "10.14", _nw_tls_create_optionsErr)
	}
	return _nw_tls_create_options(), nil
}

// Nw_tls_create_options initializes a default set of TLS connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_create_options()
func Nw_tls_create_options() Nw_protocol_options_t {
	result, callErr := tryNw_tls_create_options()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_access_bytes func(txt_record Nw_txt_record_t, access_bytes unsafe.Pointer) bool
var _nw_txt_record_access_bytesErr error

func tryNw_txt_record_access_bytes(txt_record Nw_txt_record_t, access_bytes Nw_txt_record_access_bytes_t) (bool, error) {
	if _nw_txt_record_access_bytes == nil {
		return false, symbolCallError("nw_txt_record_access_bytes", "10.15", _nw_txt_record_access_bytesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 *uint8, arg1 uint32) bool { return access_bytes(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_txt_record_access_bytes(txt_record, _block0), nil
}

// Nw_txt_record_access_bytes accesses the raw bytes contained within a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes(_:_:)
func Nw_txt_record_access_bytes(txt_record Nw_txt_record_t, access_bytes Nw_txt_record_access_bytes_t) bool {
	result, callErr := tryNw_txt_record_access_bytes(txt_record, access_bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_access_key func(txt_record Nw_txt_record_t, key string, access_value unsafe.Pointer) bool
var _nw_txt_record_access_keyErr error

func tryNw_txt_record_access_key(txt_record Nw_txt_record_t, key string, access_value Nw_txt_record_access_key_t) (bool, error) {
	if _nw_txt_record_access_key == nil {
		return false, symbolCallError("nw_txt_record_access_key", "10.15", _nw_txt_record_access_keyErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string, arg1 NwTxtRecordFindKey, arg2 *uint8, arg3 uint32) bool {
		return access_value(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_txt_record_access_key(txt_record, key, _block0), nil
}

// Nw_txt_record_access_key accesses the value for a specific key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_key(_:_:_:)
func Nw_txt_record_access_key(txt_record Nw_txt_record_t, key string, access_value Nw_txt_record_access_key_t) bool {
	result, callErr := tryNw_txt_record_access_key(txt_record, key, access_value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_apply func(txt_record Nw_txt_record_t, applier unsafe.Pointer) bool
var _nw_txt_record_applyErr error

func tryNw_txt_record_apply(txt_record Nw_txt_record_t, applier Nw_txt_record_applier_t) (bool, error) {
	if _nw_txt_record_apply == nil {
		return false, symbolCallError("nw_txt_record_apply", "10.15", _nw_txt_record_applyErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string, arg1 NwTxtRecordFindKey, arg2 *uint8, arg3 uint32) bool {
		return applier(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_txt_record_apply(txt_record, _block0), nil
}

// Nw_txt_record_apply iterates through all keys in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_apply(_:_:)
func Nw_txt_record_apply(txt_record Nw_txt_record_t, applier Nw_txt_record_applier_t) bool {
	result, callErr := tryNw_txt_record_apply(txt_record, applier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_copy func(txt_record Nw_txt_record_t) Nw_txt_record_t
var _nw_txt_record_copyErr error

func tryNw_txt_record_copy(txt_record Nw_txt_record_t) (Nw_txt_record_t, error) {
	if _nw_txt_record_copy == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_txt_record_copy", "10.15", _nw_txt_record_copyErr)
	}
	return _nw_txt_record_copy(txt_record), nil
}

// Nw_txt_record_copy performs a deep copy of a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_copy(_:)
func Nw_txt_record_copy(txt_record Nw_txt_record_t) Nw_txt_record_t {
	result, callErr := tryNw_txt_record_copy(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_create_dictionary func() Nw_txt_record_t
var _nw_txt_record_create_dictionaryErr error

func tryNw_txt_record_create_dictionary() (Nw_txt_record_t, error) {
	if _nw_txt_record_create_dictionary == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_txt_record_create_dictionary", "10.15", _nw_txt_record_create_dictionaryErr)
	}
	return _nw_txt_record_create_dictionary(), nil
}

// Nw_txt_record_create_dictionary initializes a TXT record as a dictionary of strings.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_dictionary()
func Nw_txt_record_create_dictionary() Nw_txt_record_t {
	result, callErr := tryNw_txt_record_create_dictionary()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_create_with_bytes func(txt_bytes *uint8, txt_len uintptr) Nw_txt_record_t
var _nw_txt_record_create_with_bytesErr error

func tryNw_txt_record_create_with_bytes(txt_bytes *uint8, txt_len uintptr) (Nw_txt_record_t, error) {
	if _nw_txt_record_create_with_bytes == nil {
		return *new(Nw_txt_record_t), symbolCallError("nw_txt_record_create_with_bytes", "10.15", _nw_txt_record_create_with_bytesErr)
	}
	return _nw_txt_record_create_with_bytes(txt_bytes, txt_len), nil
}

// Nw_txt_record_create_with_bytes initializes a TXT record with raw bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_with_bytes(_:_:)
func Nw_txt_record_create_with_bytes(txt_bytes *uint8, txt_len uintptr) Nw_txt_record_t {
	result, callErr := tryNw_txt_record_create_with_bytes(txt_bytes, txt_len)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_find_key func(txt_record Nw_txt_record_t, key string) NwTxtRecordFindKey
var _nw_txt_record_find_keyErr error

func tryNw_txt_record_find_key(txt_record Nw_txt_record_t, key string) (NwTxtRecordFindKey, error) {
	if _nw_txt_record_find_key == nil {
		return *new(NwTxtRecordFindKey), symbolCallError("nw_txt_record_find_key", "10.15", _nw_txt_record_find_keyErr)
	}
	return _nw_txt_record_find_key(txt_record, key), nil
}

// Nw_txt_record_find_key checks the status of value associated with a key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_find_key(_:_:)
func Nw_txt_record_find_key(txt_record Nw_txt_record_t, key string) NwTxtRecordFindKey {
	result, callErr := tryNw_txt_record_find_key(txt_record, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_get_key_count func(txt_record Nw_txt_record_t) uintptr
var _nw_txt_record_get_key_countErr error

func tryNw_txt_record_get_key_count(txt_record Nw_txt_record_t) (uintptr, error) {
	if _nw_txt_record_get_key_count == nil {
		return 0, symbolCallError("nw_txt_record_get_key_count", "10.15", _nw_txt_record_get_key_countErr)
	}
	return _nw_txt_record_get_key_count(txt_record), nil
}

// Nw_txt_record_get_key_count accesses the number of keys stored in the TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_get_key_count(_:)
func Nw_txt_record_get_key_count(txt_record Nw_txt_record_t) uintptr {
	result, callErr := tryNw_txt_record_get_key_count(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_is_dictionary func(txt_record Nw_txt_record_t) bool
var _nw_txt_record_is_dictionaryErr error

func tryNw_txt_record_is_dictionary(txt_record Nw_txt_record_t) (bool, error) {
	if _nw_txt_record_is_dictionary == nil {
		return false, symbolCallError("nw_txt_record_is_dictionary", "10.15", _nw_txt_record_is_dictionaryErr)
	}
	return _nw_txt_record_is_dictionary(txt_record), nil
}

// Nw_txt_record_is_dictionary checks whether a TXT record conforms to a dictionary format.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_dictionary(_:)
func Nw_txt_record_is_dictionary(txt_record Nw_txt_record_t) bool {
	result, callErr := tryNw_txt_record_is_dictionary(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_is_equal func(left Nw_txt_record_t, right Nw_txt_record_t) bool
var _nw_txt_record_is_equalErr error

func tryNw_txt_record_is_equal(left Nw_txt_record_t, right Nw_txt_record_t) (bool, error) {
	if _nw_txt_record_is_equal == nil {
		return false, symbolCallError("nw_txt_record_is_equal", "10.15", _nw_txt_record_is_equalErr)
	}
	return _nw_txt_record_is_equal(left, right), nil
}

// Nw_txt_record_is_equal checks whether two TXT records are equivalent.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_equal(_:_:)
func Nw_txt_record_is_equal(left Nw_txt_record_t, right Nw_txt_record_t) bool {
	result, callErr := tryNw_txt_record_is_equal(left, right)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_remove_key func(txt_record Nw_txt_record_t, key string) bool
var _nw_txt_record_remove_keyErr error

func tryNw_txt_record_remove_key(txt_record Nw_txt_record_t, key string) (bool, error) {
	if _nw_txt_record_remove_key == nil {
		return false, symbolCallError("nw_txt_record_remove_key", "10.15", _nw_txt_record_remove_keyErr)
	}
	return _nw_txt_record_remove_key(txt_record, key), nil
}

// Nw_txt_record_remove_key removes a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_remove_key(_:_:)
func Nw_txt_record_remove_key(txt_record Nw_txt_record_t, key string) bool {
	result, callErr := tryNw_txt_record_remove_key(txt_record, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_txt_record_set_key func(txt_record Nw_txt_record_t, key string, value *uint8, value_len uintptr) bool
var _nw_txt_record_set_keyErr error

func tryNw_txt_record_set_key(txt_record Nw_txt_record_t, key string, value *uint8, value_len uintptr) (bool, error) {
	if _nw_txt_record_set_key == nil {
		return false, symbolCallError("nw_txt_record_set_key", "10.15", _nw_txt_record_set_keyErr)
	}
	return _nw_txt_record_set_key(txt_record, key, value, value_len), nil
}

// Nw_txt_record_set_key sets a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_set_key(_:_:_:_:)
func Nw_txt_record_set_key(txt_record Nw_txt_record_t, key string, value *uint8, value_len uintptr) bool {
	result, callErr := tryNw_txt_record_set_key(txt_record, key, value, value_len)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_udp_create_metadata func() Nw_protocol_metadata_t
var _nw_udp_create_metadataErr error

func tryNw_udp_create_metadata() (Nw_protocol_metadata_t, error) {
	if _nw_udp_create_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_udp_create_metadata", "10.14", _nw_udp_create_metadataErr)
	}
	return _nw_udp_create_metadata(), nil
}

// Nw_udp_create_metadata initializes a default UDP message.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_metadata()
func Nw_udp_create_metadata() Nw_protocol_metadata_t {
	result, callErr := tryNw_udp_create_metadata()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_udp_create_options func() Nw_protocol_options_t
var _nw_udp_create_optionsErr error

func tryNw_udp_create_options() (Nw_protocol_options_t, error) {
	if _nw_udp_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_udp_create_options", "10.14", _nw_udp_create_optionsErr)
	}
	return _nw_udp_create_options(), nil
}

// Nw_udp_create_options initializes a default set of UDP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_options()
func Nw_udp_create_options() Nw_protocol_options_t {
	result, callErr := tryNw_udp_create_options()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_udp_options_set_prefer_no_checksum func(options Nw_protocol_options_t, prefer_no_checksum bool)
var _nw_udp_options_set_prefer_no_checksumErr error

func tryNw_udp_options_set_prefer_no_checksum(options Nw_protocol_options_t, prefer_no_checksum bool) error {
	if _nw_udp_options_set_prefer_no_checksum == nil {
		return symbolCallError("nw_udp_options_set_prefer_no_checksum", "10.14", _nw_udp_options_set_prefer_no_checksumErr)
	}
	_nw_udp_options_set_prefer_no_checksum(options, prefer_no_checksum)
	return nil
}

// Nw_udp_options_set_prefer_no_checksum configures the connection to not send UDP checksums.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_options_set_prefer_no_checksum(_:_:)
func Nw_udp_options_set_prefer_no_checksum(options Nw_protocol_options_t, prefer_no_checksum bool) {
	if callErr := tryNw_udp_options_set_prefer_no_checksum(options, prefer_no_checksum); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_create_metadata func(opcode NwWsOpcode) Nw_protocol_metadata_t
var _nw_ws_create_metadataErr error

func tryNw_ws_create_metadata(opcode NwWsOpcode) (Nw_protocol_metadata_t, error) {
	if _nw_ws_create_metadata == nil {
		return *new(Nw_protocol_metadata_t), symbolCallError("nw_ws_create_metadata", "10.15", _nw_ws_create_metadataErr)
	}
	return _nw_ws_create_metadata(opcode), nil
}

// Nw_ws_create_metadata initializes a WebSocket message with a specific type code.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_metadata(_:)
func Nw_ws_create_metadata(opcode NwWsOpcode) Nw_protocol_metadata_t {
	result, callErr := tryNw_ws_create_metadata(opcode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_create_options func(version NwWsVersion) Nw_protocol_options_t
var _nw_ws_create_optionsErr error

func tryNw_ws_create_options(version NwWsVersion) (Nw_protocol_options_t, error) {
	if _nw_ws_create_options == nil {
		return *new(Nw_protocol_options_t), symbolCallError("nw_ws_create_options", "10.15", _nw_ws_create_optionsErr)
	}
	return _nw_ws_create_options(version), nil
}

// Nw_ws_create_options initializes a default set of WebSocket connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_options(_:)
func Nw_ws_create_options(version NwWsVersion) Nw_protocol_options_t {
	result, callErr := tryNw_ws_create_options(version)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_metadata_copy_server_response func(metadata Nw_protocol_metadata_t) Nw_ws_response_t
var _nw_ws_metadata_copy_server_responseErr error

func tryNw_ws_metadata_copy_server_response(metadata Nw_protocol_metadata_t) (Nw_ws_response_t, error) {
	if _nw_ws_metadata_copy_server_response == nil {
		return *new(Nw_ws_response_t), symbolCallError("nw_ws_metadata_copy_server_response", "10.15", _nw_ws_metadata_copy_server_responseErr)
	}
	return _nw_ws_metadata_copy_server_response(metadata), nil
}

// Nw_ws_metadata_copy_server_response accesses the WebSocket server’s response sent during the handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_copy_server_response(_:)
func Nw_ws_metadata_copy_server_response(metadata Nw_protocol_metadata_t) Nw_ws_response_t {
	result, callErr := tryNw_ws_metadata_copy_server_response(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_metadata_get_close_code func(metadata Nw_protocol_metadata_t) NwWsCloseCode
var _nw_ws_metadata_get_close_codeErr error

func tryNw_ws_metadata_get_close_code(metadata Nw_protocol_metadata_t) (NwWsCloseCode, error) {
	if _nw_ws_metadata_get_close_code == nil {
		return *new(NwWsCloseCode), symbolCallError("nw_ws_metadata_get_close_code", "10.15", _nw_ws_metadata_get_close_codeErr)
	}
	return _nw_ws_metadata_get_close_code(metadata), nil
}

// Nw_ws_metadata_get_close_code accesses the close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_close_code(_:)
func Nw_ws_metadata_get_close_code(metadata Nw_protocol_metadata_t) NwWsCloseCode {
	result, callErr := tryNw_ws_metadata_get_close_code(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_metadata_get_opcode func(metadata Nw_protocol_metadata_t) NwWsOpcode
var _nw_ws_metadata_get_opcodeErr error

func tryNw_ws_metadata_get_opcode(metadata Nw_protocol_metadata_t) (NwWsOpcode, error) {
	if _nw_ws_metadata_get_opcode == nil {
		return *new(NwWsOpcode), symbolCallError("nw_ws_metadata_get_opcode", "10.15", _nw_ws_metadata_get_opcodeErr)
	}
	return _nw_ws_metadata_get_opcode(metadata), nil
}

// Nw_ws_metadata_get_opcode checks the type code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_opcode(_:)
func Nw_ws_metadata_get_opcode(metadata Nw_protocol_metadata_t) NwWsOpcode {
	result, callErr := tryNw_ws_metadata_get_opcode(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_metadata_set_close_code func(metadata Nw_protocol_metadata_t, close_code NwWsCloseCode)
var _nw_ws_metadata_set_close_codeErr error

func tryNw_ws_metadata_set_close_code(metadata Nw_protocol_metadata_t, close_code NwWsCloseCode) error {
	if _nw_ws_metadata_set_close_code == nil {
		return symbolCallError("nw_ws_metadata_set_close_code", "10.15", _nw_ws_metadata_set_close_codeErr)
	}
	_nw_ws_metadata_set_close_code(metadata, close_code)
	return nil
}

// Nw_ws_metadata_set_close_code sets a close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_close_code(_:_:)
func Nw_ws_metadata_set_close_code(metadata Nw_protocol_metadata_t, close_code NwWsCloseCode) {
	if callErr := tryNw_ws_metadata_set_close_code(metadata, close_code); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_metadata_set_pong_handler func(metadata Nw_protocol_metadata_t, client_queue uintptr, pong_handler unsafe.Pointer)
var _nw_ws_metadata_set_pong_handlerErr error

func tryNw_ws_metadata_set_pong_handler(metadata Nw_protocol_metadata_t, client_queue dispatch.Queue, pong_handler Nw_ws_pong_handler_t) error {
	if _nw_ws_metadata_set_pong_handler == nil {
		return symbolCallError("nw_ws_metadata_set_pong_handler", "10.15", _nw_ws_metadata_set_pong_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) { pong_handler(arg0) })
	retainNetworkAsyncBlock(metadata.ID, "nw_ws_metadata_set_pong_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_ws_metadata_set_pong_handler(metadata, uintptr(client_queue.Handle()), _block0)
	return nil
}

// Nw_ws_metadata_set_pong_handler sets a handler on a Ping message to be invoked when the corresponding Pong message is received.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_pong_handler(_:_:_:)
func Nw_ws_metadata_set_pong_handler(metadata Nw_protocol_metadata_t, client_queue dispatch.Queue, pong_handler Nw_ws_pong_handler_t) {
	if callErr := tryNw_ws_metadata_set_pong_handler(metadata, client_queue, pong_handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_add_additional_header func(options Nw_protocol_options_t, name string, value string)
var _nw_ws_options_add_additional_headerErr error

func tryNw_ws_options_add_additional_header(options Nw_protocol_options_t, name string, value string) error {
	if _nw_ws_options_add_additional_header == nil {
		return symbolCallError("nw_ws_options_add_additional_header", "10.15", _nw_ws_options_add_additional_headerErr)
	}
	_nw_ws_options_add_additional_header(options, name, value)
	return nil
}

// Nw_ws_options_add_additional_header adds additional HTTP header fields to be sent by the client during the WebSocket handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_additional_header(_:_:_:)
func Nw_ws_options_add_additional_header(options Nw_protocol_options_t, name string, value string) {
	if callErr := tryNw_ws_options_add_additional_header(options, name, value); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_add_subprotocol func(options Nw_protocol_options_t, subprotocol string)
var _nw_ws_options_add_subprotocolErr error

func tryNw_ws_options_add_subprotocol(options Nw_protocol_options_t, subprotocol string) error {
	if _nw_ws_options_add_subprotocol == nil {
		return symbolCallError("nw_ws_options_add_subprotocol", "10.15", _nw_ws_options_add_subprotocolErr)
	}
	_nw_ws_options_add_subprotocol(options, subprotocol)
	return nil
}

// Nw_ws_options_add_subprotocol adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_subprotocol(_:_:)
func Nw_ws_options_add_subprotocol(options Nw_protocol_options_t, subprotocol string) {
	if callErr := tryNw_ws_options_add_subprotocol(options, subprotocol); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_set_auto_reply_ping func(options Nw_protocol_options_t, auto_reply_ping bool)
var _nw_ws_options_set_auto_reply_pingErr error

func tryNw_ws_options_set_auto_reply_ping(options Nw_protocol_options_t, auto_reply_ping bool) error {
	if _nw_ws_options_set_auto_reply_ping == nil {
		return symbolCallError("nw_ws_options_set_auto_reply_ping", "10.15", _nw_ws_options_set_auto_reply_pingErr)
	}
	_nw_ws_options_set_auto_reply_ping(options, auto_reply_ping)
	return nil
}

// Nw_ws_options_set_auto_reply_ping configures the connection to automatically reply to Ping messages instead of delivering them to you.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_auto_reply_ping(_:_:)
func Nw_ws_options_set_auto_reply_ping(options Nw_protocol_options_t, auto_reply_ping bool) {
	if callErr := tryNw_ws_options_set_auto_reply_ping(options, auto_reply_ping); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_set_client_request_handler func(options Nw_protocol_options_t, client_queue uintptr, handler unsafe.Pointer)
var _nw_ws_options_set_client_request_handlerErr error

func tryNw_ws_options_set_client_request_handler(options Nw_protocol_options_t, client_queue dispatch.Queue, handler Nw_ws_client_request_handler_t) error {
	if _nw_ws_options_set_client_request_handler == nil {
		return symbolCallError("nw_ws_options_set_client_request_handler", "10.15", _nw_ws_options_set_client_request_handlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object) objectivec.Object { return handler(arg0) })
	retainNetworkAsyncBlock(options.ID, "nw_ws_options_set_client_request_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nw_ws_options_set_client_request_handler(options, uintptr(client_queue.Handle()), _block0)
	return nil
}

// Nw_ws_options_set_client_request_handler sets a handler to react to as a server to inbound WebSocket client handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_client_request_handler(_:_:_:)
func Nw_ws_options_set_client_request_handler(options Nw_protocol_options_t, client_queue dispatch.Queue, handler Nw_ws_client_request_handler_t) {
	if callErr := tryNw_ws_options_set_client_request_handler(options, client_queue, handler); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_set_maximum_message_size func(options Nw_protocol_options_t, maximum_message_size uintptr)
var _nw_ws_options_set_maximum_message_sizeErr error

func tryNw_ws_options_set_maximum_message_size(options Nw_protocol_options_t, maximum_message_size uintptr) error {
	if _nw_ws_options_set_maximum_message_size == nil {
		return symbolCallError("nw_ws_options_set_maximum_message_size", "10.15", _nw_ws_options_set_maximum_message_sizeErr)
	}
	_nw_ws_options_set_maximum_message_size(options, maximum_message_size)
	return nil
}

// Nw_ws_options_set_maximum_message_size sets the maximum allowed message size, in bytes, to be received by the WebSocket connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_maximum_message_size(_:_:)
func Nw_ws_options_set_maximum_message_size(options Nw_protocol_options_t, maximum_message_size uintptr) {
	if callErr := tryNw_ws_options_set_maximum_message_size(options, maximum_message_size); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_options_set_skip_handshake func(options Nw_protocol_options_t, skip_handshake bool)
var _nw_ws_options_set_skip_handshakeErr error

func tryNw_ws_options_set_skip_handshake(options Nw_protocol_options_t, skip_handshake bool) error {
	if _nw_ws_options_set_skip_handshake == nil {
		return symbolCallError("nw_ws_options_set_skip_handshake", "10.15", _nw_ws_options_set_skip_handshakeErr)
	}
	_nw_ws_options_set_skip_handshake(options, skip_handshake)
	return nil
}

// Nw_ws_options_set_skip_handshake specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_skip_handshake(_:_:)
func Nw_ws_options_set_skip_handshake(options Nw_protocol_options_t, skip_handshake bool) {
	if callErr := tryNw_ws_options_set_skip_handshake(options, skip_handshake); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_request_enumerate_additional_headers func(request Nw_ws_request_t, enumerator unsafe.Pointer) bool
var _nw_ws_request_enumerate_additional_headersErr error

func tryNw_ws_request_enumerate_additional_headers(request Nw_ws_request_t, enumerator Nw_ws_additional_header_enumerator_t) (bool, error) {
	if _nw_ws_request_enumerate_additional_headers == nil {
		return false, symbolCallError("nw_ws_request_enumerate_additional_headers", "10.15", _nw_ws_request_enumerate_additional_headersErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string, arg1 string) bool { return enumerator(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_ws_request_enumerate_additional_headers(request, _block0), nil
}

// Nw_ws_request_enumerate_additional_headers enumerates additional HTTP headers in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_additional_headers(_:_:)
func Nw_ws_request_enumerate_additional_headers(request Nw_ws_request_t, enumerator Nw_ws_additional_header_enumerator_t) bool {
	result, callErr := tryNw_ws_request_enumerate_additional_headers(request, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_request_enumerate_subprotocols func(request Nw_ws_request_t, enumerator unsafe.Pointer) bool
var _nw_ws_request_enumerate_subprotocolsErr error

func tryNw_ws_request_enumerate_subprotocols(request Nw_ws_request_t, enumerator Nw_ws_subprotocol_enumerator_t) (bool, error) {
	if _nw_ws_request_enumerate_subprotocols == nil {
		return false, symbolCallError("nw_ws_request_enumerate_subprotocols", "10.15", _nw_ws_request_enumerate_subprotocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string) bool { return enumerator(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_ws_request_enumerate_subprotocols(request, _block0), nil
}

// Nw_ws_request_enumerate_subprotocols enumerates the supported subprotocols in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_subprotocols(_:_:)
func Nw_ws_request_enumerate_subprotocols(request Nw_ws_request_t, enumerator Nw_ws_subprotocol_enumerator_t) bool {
	result, callErr := tryNw_ws_request_enumerate_subprotocols(request, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_response_add_additional_header func(response Nw_ws_response_t, name string, value string)
var _nw_ws_response_add_additional_headerErr error

func tryNw_ws_response_add_additional_header(response Nw_ws_response_t, name string, value string) error {
	if _nw_ws_response_add_additional_header == nil {
		return symbolCallError("nw_ws_response_add_additional_header", "10.15", _nw_ws_response_add_additional_headerErr)
	}
	_nw_ws_response_add_additional_header(response, name, value)
	return nil
}

// Nw_ws_response_add_additional_header adds an additional HTTP header to a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_add_additional_header(_:_:_:)
func Nw_ws_response_add_additional_header(response Nw_ws_response_t, name string, value string) {
	if callErr := tryNw_ws_response_add_additional_header(response, name, value); callErr != nil {
		panic(callErr)
	}
}

var _nw_ws_response_create func(status NwWsResponseStatus, selected_subprotocol string) Nw_ws_response_t
var _nw_ws_response_createErr error

func tryNw_ws_response_create(status NwWsResponseStatus, selected_subprotocol string) (Nw_ws_response_t, error) {
	if _nw_ws_response_create == nil {
		return *new(Nw_ws_response_t), symbolCallError("nw_ws_response_create", "10.15", _nw_ws_response_createErr)
	}
	return _nw_ws_response_create(status, selected_subprotocol), nil
}

// Nw_ws_response_create initializes a WebSocket server response with a status and selected subprotocol.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_create(_:_:)
func Nw_ws_response_create(status NwWsResponseStatus, selected_subprotocol string) Nw_ws_response_t {
	result, callErr := tryNw_ws_response_create(status, selected_subprotocol)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_response_enumerate_additional_headers func(response Nw_ws_response_t, enumerator unsafe.Pointer) bool
var _nw_ws_response_enumerate_additional_headersErr error

func tryNw_ws_response_enumerate_additional_headers(response Nw_ws_response_t, enumerator Nw_ws_additional_header_enumerator_t) (bool, error) {
	if _nw_ws_response_enumerate_additional_headers == nil {
		return false, symbolCallError("nw_ws_response_enumerate_additional_headers", "10.15", _nw_ws_response_enumerate_additional_headersErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 string, arg1 string) bool { return enumerator(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nw_ws_response_enumerate_additional_headers(response, _block0), nil
}

// Nw_ws_response_enumerate_additional_headers enumerates the additional HTTP headers in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_enumerate_additional_headers(_:_:)
func Nw_ws_response_enumerate_additional_headers(response Nw_ws_response_t, enumerator Nw_ws_additional_header_enumerator_t) bool {
	result, callErr := tryNw_ws_response_enumerate_additional_headers(response, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_response_get_selected_subprotocol func(response Nw_ws_response_t) *byte
var _nw_ws_response_get_selected_subprotocolErr error

func tryNw_ws_response_get_selected_subprotocol(response Nw_ws_response_t) (*byte, error) {
	if _nw_ws_response_get_selected_subprotocol == nil {
		return nil, symbolCallError("nw_ws_response_get_selected_subprotocol", "10.15", _nw_ws_response_get_selected_subprotocolErr)
	}
	return _nw_ws_response_get_selected_subprotocol(response), nil
}

// Nw_ws_response_get_selected_subprotocol accesses the selected subprotocol in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_selected_subprotocol(_:)
func Nw_ws_response_get_selected_subprotocol(response Nw_ws_response_t) *byte {
	result, callErr := tryNw_ws_response_get_selected_subprotocol(response)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nw_ws_response_get_status func(response Nw_ws_response_t) NwWsResponseStatus
var _nw_ws_response_get_statusErr error

func tryNw_ws_response_get_status(response Nw_ws_response_t) (NwWsResponseStatus, error) {
	if _nw_ws_response_get_status == nil {
		return *new(NwWsResponseStatus), symbolCallError("nw_ws_response_get_status", "10.15", _nw_ws_response_get_statusErr)
	}
	return _nw_ws_response_get_status(response), nil
}

// Nw_ws_response_get_status accesses the status of a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_status(_:)
func Nw_ws_response_get_status(response Nw_ws_response_t) NwWsResponseStatus {
	result, callErr := tryNw_ws_response_get_status(response)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerSymbol(&_nw_parameters_configure_protocol_default_configurationSymbol, &_nw_parameters_configure_protocol_default_configurationErr, frameworkHandle, "_nw_parameters_configure_protocol_default_configuration", "10.14")
	registerSymbol(&_nw_parameters_configure_protocol_disableSymbol, &_nw_parameters_configure_protocol_disableErr, frameworkHandle, "_nw_parameters_configure_protocol_disable", "10.14")
	registerFunc(&_nw_advertise_descriptor_copy_txt_record_object, &_nw_advertise_descriptor_copy_txt_record_objectErr, frameworkHandle, "nw_advertise_descriptor_copy_txt_record_object", "10.15")
	registerFunc(&_nw_advertise_descriptor_create_application_service, &_nw_advertise_descriptor_create_application_serviceErr, frameworkHandle, "nw_advertise_descriptor_create_application_service", "13.0")
	registerFunc(&_nw_advertise_descriptor_create_bonjour_service, &_nw_advertise_descriptor_create_bonjour_serviceErr, frameworkHandle, "nw_advertise_descriptor_create_bonjour_service", "10.14")
	registerFunc(&_nw_advertise_descriptor_get_application_service_name, &_nw_advertise_descriptor_get_application_service_nameErr, frameworkHandle, "nw_advertise_descriptor_get_application_service_name", "13.0")
	registerFunc(&_nw_advertise_descriptor_get_no_auto_rename, &_nw_advertise_descriptor_get_no_auto_renameErr, frameworkHandle, "nw_advertise_descriptor_get_no_auto_rename", "10.14")
	registerFunc(&_nw_advertise_descriptor_set_no_auto_rename, &_nw_advertise_descriptor_set_no_auto_renameErr, frameworkHandle, "nw_advertise_descriptor_set_no_auto_rename", "10.14")
	registerFunc(&_nw_advertise_descriptor_set_txt_record, &_nw_advertise_descriptor_set_txt_recordErr, frameworkHandle, "nw_advertise_descriptor_set_txt_record", "10.14")
	registerFunc(&_nw_advertise_descriptor_set_txt_record_object, &_nw_advertise_descriptor_set_txt_record_objectErr, frameworkHandle, "nw_advertise_descriptor_set_txt_record_object", "10.15")
	registerFunc(&_nw_browse_descriptor_create_application_service, &_nw_browse_descriptor_create_application_serviceErr, frameworkHandle, "nw_browse_descriptor_create_application_service", "13.0")
	registerFunc(&_nw_browse_descriptor_create_bonjour_service, &_nw_browse_descriptor_create_bonjour_serviceErr, frameworkHandle, "nw_browse_descriptor_create_bonjour_service", "10.15")
	registerFunc(&_nw_browse_descriptor_get_application_service_name, &_nw_browse_descriptor_get_application_service_nameErr, frameworkHandle, "nw_browse_descriptor_get_application_service_name", "13.0")
	registerFunc(&_nw_browse_descriptor_get_bonjour_service_domain, &_nw_browse_descriptor_get_bonjour_service_domainErr, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_domain", "10.15")
	registerFunc(&_nw_browse_descriptor_get_bonjour_service_type, &_nw_browse_descriptor_get_bonjour_service_typeErr, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_type", "10.15")
	registerFunc(&_nw_browse_descriptor_get_include_txt_record, &_nw_browse_descriptor_get_include_txt_recordErr, frameworkHandle, "nw_browse_descriptor_get_include_txt_record", "10.15")
	registerFunc(&_nw_browse_descriptor_set_include_txt_record, &_nw_browse_descriptor_set_include_txt_recordErr, frameworkHandle, "nw_browse_descriptor_set_include_txt_record", "10.15")
	registerFunc(&_nw_browse_result_copy_endpoint, &_nw_browse_result_copy_endpointErr, frameworkHandle, "nw_browse_result_copy_endpoint", "10.15")
	registerFunc(&_nw_browse_result_copy_txt_record_object, &_nw_browse_result_copy_txt_record_objectErr, frameworkHandle, "nw_browse_result_copy_txt_record_object", "10.15")
	registerFunc(&_nw_browse_result_enumerate_interfaces, &_nw_browse_result_enumerate_interfacesErr, frameworkHandle, "nw_browse_result_enumerate_interfaces", "10.15")
	registerFunc(&_nw_browse_result_get_changes, &_nw_browse_result_get_changesErr, frameworkHandle, "nw_browse_result_get_changes", "10.15")
	registerFunc(&_nw_browse_result_get_interfaces_count, &_nw_browse_result_get_interfaces_countErr, frameworkHandle, "nw_browse_result_get_interfaces_count", "10.15")
	registerFunc(&_nw_browser_cancel, &_nw_browser_cancelErr, frameworkHandle, "nw_browser_cancel", "10.15")
	registerFunc(&_nw_browser_copy_browse_descriptor, &_nw_browser_copy_browse_descriptorErr, frameworkHandle, "nw_browser_copy_browse_descriptor", "10.15")
	registerFunc(&_nw_browser_copy_parameters, &_nw_browser_copy_parametersErr, frameworkHandle, "nw_browser_copy_parameters", "10.15")
	registerFunc(&_nw_browser_create, &_nw_browser_createErr, frameworkHandle, "nw_browser_create", "10.15")
	registerFunc(&_nw_browser_set_browse_results_changed_handler, &_nw_browser_set_browse_results_changed_handlerErr, frameworkHandle, "nw_browser_set_browse_results_changed_handler", "10.15")
	registerFunc(&_nw_browser_set_queue, &_nw_browser_set_queueErr, frameworkHandle, "nw_browser_set_queue", "10.15")
	registerFunc(&_nw_browser_set_state_changed_handler, &_nw_browser_set_state_changed_handlerErr, frameworkHandle, "nw_browser_set_state_changed_handler", "10.15")
	registerFunc(&_nw_browser_start, &_nw_browser_startErr, frameworkHandle, "nw_browser_start", "10.15")
	registerFunc(&_nw_connection_access_establishment_report, &_nw_connection_access_establishment_reportErr, frameworkHandle, "nw_connection_access_establishment_report", "10.15")
	registerFunc(&_nw_connection_batch, &_nw_connection_batchErr, frameworkHandle, "nw_connection_batch", "10.14")
	registerFunc(&_nw_connection_cancel, &_nw_connection_cancelErr, frameworkHandle, "nw_connection_cancel", "10.14")
	registerFunc(&_nw_connection_cancel_current_endpoint, &_nw_connection_cancel_current_endpointErr, frameworkHandle, "nw_connection_cancel_current_endpoint", "10.14")
	registerFunc(&_nw_connection_copy_current_path, &_nw_connection_copy_current_pathErr, frameworkHandle, "nw_connection_copy_current_path", "10.14")
	registerFunc(&_nw_connection_copy_description, &_nw_connection_copy_descriptionErr, frameworkHandle, "nw_connection_copy_description", "10.14")
	registerFunc(&_nw_connection_copy_endpoint, &_nw_connection_copy_endpointErr, frameworkHandle, "nw_connection_copy_endpoint", "10.14")
	registerFunc(&_nw_connection_copy_parameters, &_nw_connection_copy_parametersErr, frameworkHandle, "nw_connection_copy_parameters", "10.14")
	registerFunc(&_nw_connection_copy_protocol_metadata, &_nw_connection_copy_protocol_metadataErr, frameworkHandle, "nw_connection_copy_protocol_metadata", "10.14")
	registerFunc(&_nw_connection_create, &_nw_connection_createErr, frameworkHandle, "nw_connection_create", "10.14")
	registerFunc(&_nw_connection_create_new_data_transfer_report, &_nw_connection_create_new_data_transfer_reportErr, frameworkHandle, "nw_connection_create_new_data_transfer_report", "10.15")
	registerFunc(&_nw_connection_force_cancel, &_nw_connection_force_cancelErr, frameworkHandle, "nw_connection_force_cancel", "10.14")
	registerFunc(&_nw_connection_get_maximum_datagram_size, &_nw_connection_get_maximum_datagram_sizeErr, frameworkHandle, "nw_connection_get_maximum_datagram_size", "10.14")
	registerFunc(&_nw_connection_group_cancel, &_nw_connection_group_cancelErr, frameworkHandle, "nw_connection_group_cancel", "11.0")
	registerFunc(&_nw_connection_group_copy_descriptor, &_nw_connection_group_copy_descriptorErr, frameworkHandle, "nw_connection_group_copy_descriptor", "11.0")
	registerFunc(&_nw_connection_group_copy_local_endpoint_for_message, &_nw_connection_group_copy_local_endpoint_for_messageErr, frameworkHandle, "nw_connection_group_copy_local_endpoint_for_message", "11.0")
	registerFunc(&_nw_connection_group_copy_parameters, &_nw_connection_group_copy_parametersErr, frameworkHandle, "nw_connection_group_copy_parameters", "11.0")
	registerFunc(&_nw_connection_group_copy_path_for_message, &_nw_connection_group_copy_path_for_messageErr, frameworkHandle, "nw_connection_group_copy_path_for_message", "11.0")
	registerFunc(&_nw_connection_group_copy_protocol_metadata, &_nw_connection_group_copy_protocol_metadataErr, frameworkHandle, "nw_connection_group_copy_protocol_metadata", "12.0")
	registerFunc(&_nw_connection_group_copy_protocol_metadata_for_message, &_nw_connection_group_copy_protocol_metadata_for_messageErr, frameworkHandle, "nw_connection_group_copy_protocol_metadata_for_message", "12.0")
	registerFunc(&_nw_connection_group_copy_remote_endpoint_for_message, &_nw_connection_group_copy_remote_endpoint_for_messageErr, frameworkHandle, "nw_connection_group_copy_remote_endpoint_for_message", "11.0")
	registerFunc(&_nw_connection_group_create, &_nw_connection_group_createErr, frameworkHandle, "nw_connection_group_create", "11.0")
	registerFunc(&_nw_connection_group_extract_connection, &_nw_connection_group_extract_connectionErr, frameworkHandle, "nw_connection_group_extract_connection", "12.0")
	registerFunc(&_nw_connection_group_extract_connection_for_message, &_nw_connection_group_extract_connection_for_messageErr, frameworkHandle, "nw_connection_group_extract_connection_for_message", "11.0")
	registerFunc(&_nw_connection_group_reinsert_extracted_connection, &_nw_connection_group_reinsert_extracted_connectionErr, frameworkHandle, "nw_connection_group_reinsert_extracted_connection", "12.0")
	registerFunc(&_nw_connection_group_reply, &_nw_connection_group_replyErr, frameworkHandle, "nw_connection_group_reply", "11.0")
	registerFunc(&_nw_connection_group_send_message, &_nw_connection_group_send_messageErr, frameworkHandle, "nw_connection_group_send_message", "11.0")
	registerFunc(&_nw_connection_group_set_new_connection_handler, &_nw_connection_group_set_new_connection_handlerErr, frameworkHandle, "nw_connection_group_set_new_connection_handler", "12.0")
	registerFunc(&_nw_connection_group_set_queue, &_nw_connection_group_set_queueErr, frameworkHandle, "nw_connection_group_set_queue", "11.0")
	registerFunc(&_nw_connection_group_set_receive_handler, &_nw_connection_group_set_receive_handlerErr, frameworkHandle, "nw_connection_group_set_receive_handler", "11.0")
	registerFunc(&_nw_connection_group_set_state_changed_handler, &_nw_connection_group_set_state_changed_handlerErr, frameworkHandle, "nw_connection_group_set_state_changed_handler", "11.0")
	registerFunc(&_nw_connection_group_start, &_nw_connection_group_startErr, frameworkHandle, "nw_connection_group_start", "11.0")
	registerFunc(&_nw_connection_receive, &_nw_connection_receiveErr, frameworkHandle, "nw_connection_receive", "10.14")
	registerFunc(&_nw_connection_receive_message, &_nw_connection_receive_messageErr, frameworkHandle, "nw_connection_receive_message", "10.14")
	registerFunc(&_nw_connection_restart, &_nw_connection_restartErr, frameworkHandle, "nw_connection_restart", "10.14")
	registerFunc(&_nw_connection_send, &_nw_connection_sendErr, frameworkHandle, "nw_connection_send", "10.14")
	registerFunc(&_nw_connection_set_better_path_available_handler, &_nw_connection_set_better_path_available_handlerErr, frameworkHandle, "nw_connection_set_better_path_available_handler", "10.14")
	registerFunc(&_nw_connection_set_path_changed_handler, &_nw_connection_set_path_changed_handlerErr, frameworkHandle, "nw_connection_set_path_changed_handler", "10.14")
	registerFunc(&_nw_connection_set_queue, &_nw_connection_set_queueErr, frameworkHandle, "nw_connection_set_queue", "10.14")
	registerFunc(&_nw_connection_set_state_changed_handler, &_nw_connection_set_state_changed_handlerErr, frameworkHandle, "nw_connection_set_state_changed_handler", "10.14")
	registerFunc(&_nw_connection_set_viability_changed_handler, &_nw_connection_set_viability_changed_handlerErr, frameworkHandle, "nw_connection_set_viability_changed_handler", "10.14")
	registerFunc(&_nw_connection_start, &_nw_connection_startErr, frameworkHandle, "nw_connection_start", "10.14")
	registerFunc(&_nw_content_context_copy_antecedent, &_nw_content_context_copy_antecedentErr, frameworkHandle, "nw_content_context_copy_antecedent", "10.14")
	registerFunc(&_nw_content_context_copy_protocol_metadata, &_nw_content_context_copy_protocol_metadataErr, frameworkHandle, "nw_content_context_copy_protocol_metadata", "10.14")
	registerFunc(&_nw_content_context_create, &_nw_content_context_createErr, frameworkHandle, "nw_content_context_create", "10.14")
	registerFunc(&_nw_content_context_foreach_protocol_metadata, &_nw_content_context_foreach_protocol_metadataErr, frameworkHandle, "nw_content_context_foreach_protocol_metadata", "10.14")
	registerFunc(&_nw_content_context_get_expiration_milliseconds, &_nw_content_context_get_expiration_millisecondsErr, frameworkHandle, "nw_content_context_get_expiration_milliseconds", "10.14")
	registerFunc(&_nw_content_context_get_identifier, &_nw_content_context_get_identifierErr, frameworkHandle, "nw_content_context_get_identifier", "10.14")
	registerFunc(&_nw_content_context_get_is_final, &_nw_content_context_get_is_finalErr, frameworkHandle, "nw_content_context_get_is_final", "10.14")
	registerFunc(&_nw_content_context_get_relative_priority, &_nw_content_context_get_relative_priorityErr, frameworkHandle, "nw_content_context_get_relative_priority", "10.14")
	registerFunc(&_nw_content_context_set_antecedent, &_nw_content_context_set_antecedentErr, frameworkHandle, "nw_content_context_set_antecedent", "10.14")
	registerFunc(&_nw_content_context_set_expiration_milliseconds, &_nw_content_context_set_expiration_millisecondsErr, frameworkHandle, "nw_content_context_set_expiration_milliseconds", "10.14")
	registerFunc(&_nw_content_context_set_is_final, &_nw_content_context_set_is_finalErr, frameworkHandle, "nw_content_context_set_is_final", "10.14")
	registerFunc(&_nw_content_context_set_metadata_for_protocol, &_nw_content_context_set_metadata_for_protocolErr, frameworkHandle, "nw_content_context_set_metadata_for_protocol", "10.14")
	registerFunc(&_nw_content_context_set_relative_priority, &_nw_content_context_set_relative_priorityErr, frameworkHandle, "nw_content_context_set_relative_priority", "10.14")
	registerFunc(&_nw_data_transfer_report_collect, &_nw_data_transfer_report_collectErr, frameworkHandle, "nw_data_transfer_report_collect", "10.15")
	registerFunc(&_nw_data_transfer_report_copy_path_interface, &_nw_data_transfer_report_copy_path_interfaceErr, frameworkHandle, "nw_data_transfer_report_copy_path_interface", "10.15")
	registerFunc(&_nw_data_transfer_report_get_duration_milliseconds, &_nw_data_transfer_report_get_duration_millisecondsErr, frameworkHandle, "nw_data_transfer_report_get_duration_milliseconds", "10.15")
	registerFunc(&_nw_data_transfer_report_get_path_count, &_nw_data_transfer_report_get_path_countErr, frameworkHandle, "nw_data_transfer_report_get_path_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_path_radio_type, &_nw_data_transfer_report_get_path_radio_typeErr, frameworkHandle, "nw_data_transfer_report_get_path_radio_type", "12.0")
	registerFunc(&_nw_data_transfer_report_get_received_application_byte_count, &_nw_data_transfer_report_get_received_application_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_received_application_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_received_ip_packet_count, &_nw_data_transfer_report_get_received_ip_packet_countErr, frameworkHandle, "nw_data_transfer_report_get_received_ip_packet_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_received_transport_byte_count, &_nw_data_transfer_report_get_received_transport_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_received_transport_duplicate_byte_count, &_nw_data_transfer_report_get_received_transport_duplicate_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_duplicate_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_received_transport_out_of_order_byte_count, &_nw_data_transfer_report_get_received_transport_out_of_order_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_out_of_order_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_sent_application_byte_count, &_nw_data_transfer_report_get_sent_application_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_sent_application_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_sent_ip_packet_count, &_nw_data_transfer_report_get_sent_ip_packet_countErr, frameworkHandle, "nw_data_transfer_report_get_sent_ip_packet_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_sent_transport_byte_count, &_nw_data_transfer_report_get_sent_transport_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_sent_transport_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count, &_nw_data_transfer_report_get_sent_transport_retransmitted_byte_countErr, frameworkHandle, "nw_data_transfer_report_get_sent_transport_retransmitted_byte_count", "10.15")
	registerFunc(&_nw_data_transfer_report_get_state, &_nw_data_transfer_report_get_stateErr, frameworkHandle, "nw_data_transfer_report_get_state", "10.15")
	registerFunc(&_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds, &_nw_data_transfer_report_get_transport_minimum_rtt_millisecondsErr, frameworkHandle, "nw_data_transfer_report_get_transport_minimum_rtt_milliseconds", "10.15")
	registerFunc(&_nw_data_transfer_report_get_transport_rtt_variance, &_nw_data_transfer_report_get_transport_rtt_varianceErr, frameworkHandle, "nw_data_transfer_report_get_transport_rtt_variance", "10.15")
	registerFunc(&_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds, &_nw_data_transfer_report_get_transport_smoothed_rtt_millisecondsErr, frameworkHandle, "nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds", "10.15")
	registerFunc(&_nw_endpoint_copy_address_string, &_nw_endpoint_copy_address_stringErr, frameworkHandle, "nw_endpoint_copy_address_string", "10.14")
	registerFunc(&_nw_endpoint_copy_port_string, &_nw_endpoint_copy_port_stringErr, frameworkHandle, "nw_endpoint_copy_port_string", "10.14")
	registerFunc(&_nw_endpoint_copy_txt_record, &_nw_endpoint_copy_txt_recordErr, frameworkHandle, "nw_endpoint_copy_txt_record", "13.0")
	registerFunc(&_nw_endpoint_create_address, &_nw_endpoint_create_addressErr, frameworkHandle, "nw_endpoint_create_address", "10.14")
	registerFunc(&_nw_endpoint_create_bonjour_service, &_nw_endpoint_create_bonjour_serviceErr, frameworkHandle, "nw_endpoint_create_bonjour_service", "10.14")
	registerFunc(&_nw_endpoint_create_host, &_nw_endpoint_create_hostErr, frameworkHandle, "nw_endpoint_create_host", "10.14")
	registerFunc(&_nw_endpoint_create_url, &_nw_endpoint_create_urlErr, frameworkHandle, "nw_endpoint_create_url", "10.15")
	registerFunc(&_nw_endpoint_get_address, &_nw_endpoint_get_addressErr, frameworkHandle, "nw_endpoint_get_address", "10.14")
	registerFunc(&_nw_endpoint_get_bonjour_service_domain, &_nw_endpoint_get_bonjour_service_domainErr, frameworkHandle, "nw_endpoint_get_bonjour_service_domain", "10.14")
	registerFunc(&_nw_endpoint_get_bonjour_service_name, &_nw_endpoint_get_bonjour_service_nameErr, frameworkHandle, "nw_endpoint_get_bonjour_service_name", "10.14")
	registerFunc(&_nw_endpoint_get_bonjour_service_type, &_nw_endpoint_get_bonjour_service_typeErr, frameworkHandle, "nw_endpoint_get_bonjour_service_type", "10.14")
	registerFunc(&_nw_endpoint_get_hostname, &_nw_endpoint_get_hostnameErr, frameworkHandle, "nw_endpoint_get_hostname", "10.14")
	registerFunc(&_nw_endpoint_get_port, &_nw_endpoint_get_portErr, frameworkHandle, "nw_endpoint_get_port", "10.14")
	registerFunc(&_nw_endpoint_get_signature, &_nw_endpoint_get_signatureErr, frameworkHandle, "nw_endpoint_get_signature", "13.0")
	registerFunc(&_nw_endpoint_get_type, &_nw_endpoint_get_typeErr, frameworkHandle, "nw_endpoint_get_type", "10.14")
	registerFunc(&_nw_endpoint_get_url, &_nw_endpoint_get_urlErr, frameworkHandle, "nw_endpoint_get_url", "10.15")
	registerFunc(&_nw_error_copy_cf_error, &_nw_error_copy_cf_errorErr, frameworkHandle, "nw_error_copy_cf_error", "10.14")
	registerFunc(&_nw_error_get_error_code, &_nw_error_get_error_codeErr, frameworkHandle, "nw_error_get_error_code", "10.14")
	registerFunc(&_nw_error_get_error_domain, &_nw_error_get_error_domainErr, frameworkHandle, "nw_error_get_error_domain", "10.14")
	registerFunc(&_nw_establishment_report_copy_proxy_endpoint, &_nw_establishment_report_copy_proxy_endpointErr, frameworkHandle, "nw_establishment_report_copy_proxy_endpoint", "10.15")
	registerFunc(&_nw_establishment_report_enumerate_protocols, &_nw_establishment_report_enumerate_protocolsErr, frameworkHandle, "nw_establishment_report_enumerate_protocols", "10.15")
	registerFunc(&_nw_establishment_report_enumerate_resolution_reports, &_nw_establishment_report_enumerate_resolution_reportsErr, frameworkHandle, "nw_establishment_report_enumerate_resolution_reports", "11.0")
	registerFunc(&_nw_establishment_report_enumerate_resolutions, &_nw_establishment_report_enumerate_resolutionsErr, frameworkHandle, "nw_establishment_report_enumerate_resolutions", "10.15")
	registerFunc(&_nw_establishment_report_get_attempt_started_after_milliseconds, &_nw_establishment_report_get_attempt_started_after_millisecondsErr, frameworkHandle, "nw_establishment_report_get_attempt_started_after_milliseconds", "10.15")
	registerFunc(&_nw_establishment_report_get_duration_milliseconds, &_nw_establishment_report_get_duration_millisecondsErr, frameworkHandle, "nw_establishment_report_get_duration_milliseconds", "10.15")
	registerFunc(&_nw_establishment_report_get_previous_attempt_count, &_nw_establishment_report_get_previous_attempt_countErr, frameworkHandle, "nw_establishment_report_get_previous_attempt_count", "10.15")
	registerFunc(&_nw_establishment_report_get_proxy_configured, &_nw_establishment_report_get_proxy_configuredErr, frameworkHandle, "nw_establishment_report_get_proxy_configured", "10.15")
	registerFunc(&_nw_establishment_report_get_used_proxy, &_nw_establishment_report_get_used_proxyErr, frameworkHandle, "nw_establishment_report_get_used_proxy", "10.15")
	registerFunc(&_nw_ethernet_channel_cancel, &_nw_ethernet_channel_cancelErr, frameworkHandle, "nw_ethernet_channel_cancel", "10.15")
	registerFunc(&_nw_ethernet_channel_create, &_nw_ethernet_channel_createErr, frameworkHandle, "nw_ethernet_channel_create", "10.15")
	registerFunc(&_nw_ethernet_channel_create_with_parameters, &_nw_ethernet_channel_create_with_parametersErr, frameworkHandle, "nw_ethernet_channel_create_with_parameters", "13.0")
	registerFunc(&_nw_ethernet_channel_get_maximum_payload_size, &_nw_ethernet_channel_get_maximum_payload_sizeErr, frameworkHandle, "nw_ethernet_channel_get_maximum_payload_size", "13.0")
	registerFunc(&_nw_ethernet_channel_send, &_nw_ethernet_channel_sendErr, frameworkHandle, "nw_ethernet_channel_send", "10.15")
	registerFunc(&_nw_ethernet_channel_set_queue, &_nw_ethernet_channel_set_queueErr, frameworkHandle, "nw_ethernet_channel_set_queue", "10.15")
	registerFunc(&_nw_ethernet_channel_set_receive_handler, &_nw_ethernet_channel_set_receive_handlerErr, frameworkHandle, "nw_ethernet_channel_set_receive_handler", "10.15")
	registerFunc(&_nw_ethernet_channel_set_state_changed_handler, &_nw_ethernet_channel_set_state_changed_handlerErr, frameworkHandle, "nw_ethernet_channel_set_state_changed_handler", "10.15")
	registerFunc(&_nw_ethernet_channel_start, &_nw_ethernet_channel_startErr, frameworkHandle, "nw_ethernet_channel_start", "10.15")
	registerFunc(&_nw_framer_async, &_nw_framer_asyncErr, frameworkHandle, "nw_framer_async", "10.15")
	registerFunc(&_nw_framer_copy_local_endpoint, &_nw_framer_copy_local_endpointErr, frameworkHandle, "nw_framer_copy_local_endpoint", "10.15")
	registerFunc(&_nw_framer_copy_options, &_nw_framer_copy_optionsErr, frameworkHandle, "nw_framer_copy_options", "12.3")
	registerFunc(&_nw_framer_copy_parameters, &_nw_framer_copy_parametersErr, frameworkHandle, "nw_framer_copy_parameters", "10.15")
	registerFunc(&_nw_framer_copy_remote_endpoint, &_nw_framer_copy_remote_endpointErr, frameworkHandle, "nw_framer_copy_remote_endpoint", "10.15")
	registerFunc(&_nw_framer_create_definition, &_nw_framer_create_definitionErr, frameworkHandle, "nw_framer_create_definition", "10.15")
	registerFunc(&_nw_framer_create_options, &_nw_framer_create_optionsErr, frameworkHandle, "nw_framer_create_options", "10.15")
	registerFunc(&_nw_framer_deliver_input, &_nw_framer_deliver_inputErr, frameworkHandle, "nw_framer_deliver_input", "10.15")
	registerFunc(&_nw_framer_deliver_input_no_copy, &_nw_framer_deliver_input_no_copyErr, frameworkHandle, "nw_framer_deliver_input_no_copy", "10.15")
	registerFunc(&_nw_framer_mark_failed_with_error, &_nw_framer_mark_failed_with_errorErr, frameworkHandle, "nw_framer_mark_failed_with_error", "10.15")
	registerFunc(&_nw_framer_mark_ready, &_nw_framer_mark_readyErr, frameworkHandle, "nw_framer_mark_ready", "10.15")
	registerFunc(&_nw_framer_message_access_value, &_nw_framer_message_access_valueErr, frameworkHandle, "nw_framer_message_access_value", "10.15")
	registerFunc(&_nw_framer_message_copy_object_value, &_nw_framer_message_copy_object_valueErr, frameworkHandle, "nw_framer_message_copy_object_value", "10.15")
	registerFunc(&_nw_framer_message_create, &_nw_framer_message_createErr, frameworkHandle, "nw_framer_message_create", "10.15")
	registerFunc(&_nw_framer_message_set_object_value, &_nw_framer_message_set_object_valueErr, frameworkHandle, "nw_framer_message_set_object_value", "10.15")
	registerFunc(&_nw_framer_message_set_value, &_nw_framer_message_set_valueErr, frameworkHandle, "nw_framer_message_set_value", "10.15")
	registerFunc(&_nw_framer_options_copy_object_value, &_nw_framer_options_copy_object_valueErr, frameworkHandle, "nw_framer_options_copy_object_value", "12.3")
	registerFunc(&_nw_framer_options_set_object_value, &_nw_framer_options_set_object_valueErr, frameworkHandle, "nw_framer_options_set_object_value", "12.3")
	registerFunc(&_nw_framer_parse_input, &_nw_framer_parse_inputErr, frameworkHandle, "nw_framer_parse_input", "10.15")
	registerFunc(&_nw_framer_parse_output, &_nw_framer_parse_outputErr, frameworkHandle, "nw_framer_parse_output", "10.15")
	registerFunc(&_nw_framer_pass_through_input, &_nw_framer_pass_through_inputErr, frameworkHandle, "nw_framer_pass_through_input", "10.15")
	registerFunc(&_nw_framer_pass_through_output, &_nw_framer_pass_through_outputErr, frameworkHandle, "nw_framer_pass_through_output", "10.15")
	registerFunc(&_nw_framer_prepend_application_protocol, &_nw_framer_prepend_application_protocolErr, frameworkHandle, "nw_framer_prepend_application_protocol", "10.15")
	registerFunc(&_nw_framer_protocol_create_message, &_nw_framer_protocol_create_messageErr, frameworkHandle, "nw_framer_protocol_create_message", "10.15")
	registerFunc(&_nw_framer_schedule_wakeup, &_nw_framer_schedule_wakeupErr, frameworkHandle, "nw_framer_schedule_wakeup", "10.15")
	registerFunc(&_nw_framer_set_cleanup_handler, &_nw_framer_set_cleanup_handlerErr, frameworkHandle, "nw_framer_set_cleanup_handler", "10.15")
	registerFunc(&_nw_framer_set_input_handler, &_nw_framer_set_input_handlerErr, frameworkHandle, "nw_framer_set_input_handler", "10.15")
	registerFunc(&_nw_framer_set_output_handler, &_nw_framer_set_output_handlerErr, frameworkHandle, "nw_framer_set_output_handler", "10.15")
	registerFunc(&_nw_framer_set_stop_handler, &_nw_framer_set_stop_handlerErr, frameworkHandle, "nw_framer_set_stop_handler", "10.15")
	registerFunc(&_nw_framer_set_wakeup_handler, &_nw_framer_set_wakeup_handlerErr, frameworkHandle, "nw_framer_set_wakeup_handler", "10.15")
	registerFunc(&_nw_framer_write_output, &_nw_framer_write_outputErr, frameworkHandle, "nw_framer_write_output", "10.15")
	registerFunc(&_nw_framer_write_output_data, &_nw_framer_write_output_dataErr, frameworkHandle, "nw_framer_write_output_data", "10.15")
	registerFunc(&_nw_framer_write_output_no_copy, &_nw_framer_write_output_no_copyErr, frameworkHandle, "nw_framer_write_output_no_copy", "10.15")
	registerFunc(&_nw_group_descriptor_add_endpoint, &_nw_group_descriptor_add_endpointErr, frameworkHandle, "nw_group_descriptor_add_endpoint", "11.0")
	registerFunc(&_nw_group_descriptor_create_multicast, &_nw_group_descriptor_create_multicastErr, frameworkHandle, "nw_group_descriptor_create_multicast", "11.0")
	registerFunc(&_nw_group_descriptor_create_multiplex, &_nw_group_descriptor_create_multiplexErr, frameworkHandle, "nw_group_descriptor_create_multiplex", "12.0")
	registerFunc(&_nw_group_descriptor_enumerate_endpoints, &_nw_group_descriptor_enumerate_endpointsErr, frameworkHandle, "nw_group_descriptor_enumerate_endpoints", "11.0")
	registerFunc(&_nw_interface_get_index, &_nw_interface_get_indexErr, frameworkHandle, "nw_interface_get_index", "10.14")
	registerFunc(&_nw_interface_get_name, &_nw_interface_get_nameErr, frameworkHandle, "nw_interface_get_name", "10.14")
	registerFunc(&_nw_interface_get_type, &_nw_interface_get_typeErr, frameworkHandle, "nw_interface_get_type", "10.14")
	registerFunc(&_nw_ip_create_metadata, &_nw_ip_create_metadataErr, frameworkHandle, "nw_ip_create_metadata", "10.14")
	registerFunc(&_nw_ip_metadata_get_ecn_flag, &_nw_ip_metadata_get_ecn_flagErr, frameworkHandle, "nw_ip_metadata_get_ecn_flag", "10.14")
	registerFunc(&_nw_ip_metadata_get_receive_time, &_nw_ip_metadata_get_receive_timeErr, frameworkHandle, "nw_ip_metadata_get_receive_time", "10.14")
	registerFunc(&_nw_ip_metadata_get_service_class, &_nw_ip_metadata_get_service_classErr, frameworkHandle, "nw_ip_metadata_get_service_class", "10.14")
	registerFunc(&_nw_ip_metadata_set_ecn_flag, &_nw_ip_metadata_set_ecn_flagErr, frameworkHandle, "nw_ip_metadata_set_ecn_flag", "10.14")
	registerFunc(&_nw_ip_metadata_set_service_class, &_nw_ip_metadata_set_service_classErr, frameworkHandle, "nw_ip_metadata_set_service_class", "10.14")
	registerFunc(&_nw_ip_options_set_calculate_receive_time, &_nw_ip_options_set_calculate_receive_timeErr, frameworkHandle, "nw_ip_options_set_calculate_receive_time", "10.14")
	registerFunc(&_nw_ip_options_set_disable_fragmentation, &_nw_ip_options_set_disable_fragmentationErr, frameworkHandle, "nw_ip_options_set_disable_fragmentation", "10.14")
	registerFunc(&_nw_ip_options_set_disable_multicast_loopback, &_nw_ip_options_set_disable_multicast_loopbackErr, frameworkHandle, "nw_ip_options_set_disable_multicast_loopback", "11.0")
	registerFunc(&_nw_ip_options_set_hop_limit, &_nw_ip_options_set_hop_limitErr, frameworkHandle, "nw_ip_options_set_hop_limit", "10.14")
	registerFunc(&_nw_ip_options_set_local_address_preference, &_nw_ip_options_set_local_address_preferenceErr, frameworkHandle, "nw_ip_options_set_local_address_preference", "10.15")
	registerFunc(&_nw_ip_options_set_use_minimum_mtu, &_nw_ip_options_set_use_minimum_mtuErr, frameworkHandle, "nw_ip_options_set_use_minimum_mtu", "10.14")
	registerFunc(&_nw_ip_options_set_version, &_nw_ip_options_set_versionErr, frameworkHandle, "nw_ip_options_set_version", "10.14")
	registerFunc(&_nw_listener_cancel, &_nw_listener_cancelErr, frameworkHandle, "nw_listener_cancel", "10.14")
	registerFunc(&_nw_listener_create, &_nw_listener_createErr, frameworkHandle, "nw_listener_create", "10.14")
	registerFunc(&_nw_listener_create_with_connection, &_nw_listener_create_with_connectionErr, frameworkHandle, "nw_listener_create_with_connection", "10.14")
	registerFunc(&_nw_listener_create_with_launchd_key, &_nw_listener_create_with_launchd_keyErr, frameworkHandle, "nw_listener_create_with_launchd_key", "10.14")
	registerFunc(&_nw_listener_create_with_port, &_nw_listener_create_with_portErr, frameworkHandle, "nw_listener_create_with_port", "10.14")
	registerFunc(&_nw_listener_get_new_connection_limit, &_nw_listener_get_new_connection_limitErr, frameworkHandle, "nw_listener_get_new_connection_limit", "10.15")
	registerFunc(&_nw_listener_get_port, &_nw_listener_get_portErr, frameworkHandle, "nw_listener_get_port", "10.14")
	registerFunc(&_nw_listener_set_advertise_descriptor, &_nw_listener_set_advertise_descriptorErr, frameworkHandle, "nw_listener_set_advertise_descriptor", "10.14")
	registerFunc(&_nw_listener_set_advertised_endpoint_changed_handler, &_nw_listener_set_advertised_endpoint_changed_handlerErr, frameworkHandle, "nw_listener_set_advertised_endpoint_changed_handler", "10.14")
	registerFunc(&_nw_listener_set_new_connection_group_handler, &_nw_listener_set_new_connection_group_handlerErr, frameworkHandle, "nw_listener_set_new_connection_group_handler", "12.0")
	registerFunc(&_nw_listener_set_new_connection_handler, &_nw_listener_set_new_connection_handlerErr, frameworkHandle, "nw_listener_set_new_connection_handler", "10.14")
	registerFunc(&_nw_listener_set_new_connection_limit, &_nw_listener_set_new_connection_limitErr, frameworkHandle, "nw_listener_set_new_connection_limit", "10.15")
	registerFunc(&_nw_listener_set_queue, &_nw_listener_set_queueErr, frameworkHandle, "nw_listener_set_queue", "10.14")
	registerFunc(&_nw_listener_set_state_changed_handler, &_nw_listener_set_state_changed_handlerErr, frameworkHandle, "nw_listener_set_state_changed_handler", "10.14")
	registerFunc(&_nw_listener_start, &_nw_listener_startErr, frameworkHandle, "nw_listener_start", "10.14")
	registerFunc(&_nw_multicast_group_descriptor_get_disable_unicast_traffic, &_nw_multicast_group_descriptor_get_disable_unicast_trafficErr, frameworkHandle, "nw_multicast_group_descriptor_get_disable_unicast_traffic", "11.0")
	registerFunc(&_nw_multicast_group_descriptor_set_disable_unicast_traffic, &_nw_multicast_group_descriptor_set_disable_unicast_trafficErr, frameworkHandle, "nw_multicast_group_descriptor_set_disable_unicast_traffic", "11.0")
	registerFunc(&_nw_multicast_group_descriptor_set_specific_source, &_nw_multicast_group_descriptor_set_specific_sourceErr, frameworkHandle, "nw_multicast_group_descriptor_set_specific_source", "11.0")
	registerFunc(&_nw_parameters_clear_prohibited_interface_types, &_nw_parameters_clear_prohibited_interface_typesErr, frameworkHandle, "nw_parameters_clear_prohibited_interface_types", "10.14")
	registerFunc(&_nw_parameters_clear_prohibited_interfaces, &_nw_parameters_clear_prohibited_interfacesErr, frameworkHandle, "nw_parameters_clear_prohibited_interfaces", "10.14")
	registerFunc(&_nw_parameters_copy, &_nw_parameters_copyErr, frameworkHandle, "nw_parameters_copy", "10.14")
	registerFunc(&_nw_parameters_copy_default_protocol_stack, &_nw_parameters_copy_default_protocol_stackErr, frameworkHandle, "nw_parameters_copy_default_protocol_stack", "10.14")
	registerFunc(&_nw_parameters_copy_local_endpoint, &_nw_parameters_copy_local_endpointErr, frameworkHandle, "nw_parameters_copy_local_endpoint", "10.14")
	registerFunc(&_nw_parameters_copy_required_interface, &_nw_parameters_copy_required_interfaceErr, frameworkHandle, "nw_parameters_copy_required_interface", "10.14")
	registerFunc(&_nw_parameters_create, &_nw_parameters_createErr, frameworkHandle, "nw_parameters_create", "10.14")
	registerFunc(&_nw_parameters_create_application_service, &_nw_parameters_create_application_serviceErr, frameworkHandle, "nw_parameters_create_application_service", "13.0")
	registerFunc(&_nw_parameters_create_custom_ip, &_nw_parameters_create_custom_ipErr, frameworkHandle, "nw_parameters_create_custom_ip", "10.15")
	registerFunc(&_nw_parameters_create_quic, &_nw_parameters_create_quicErr, frameworkHandle, "nw_parameters_create_quic", "12.0")
	registerFunc(&_nw_parameters_create_secure_tcp, &_nw_parameters_create_secure_tcpErr, frameworkHandle, "nw_parameters_create_secure_tcp", "10.14")
	registerFunc(&_nw_parameters_create_secure_udp, &_nw_parameters_create_secure_udpErr, frameworkHandle, "nw_parameters_create_secure_udp", "10.14")
	registerFunc(&_nw_parameters_get_allow_ultra_constrained, &_nw_parameters_get_allow_ultra_constrainedErr, frameworkHandle, "nw_parameters_get_allow_ultra_constrained", "26.0")
	registerFunc(&_nw_parameters_get_attribution, &_nw_parameters_get_attributionErr, frameworkHandle, "nw_parameters_get_attribution", "12.0")
	registerFunc(&_nw_parameters_get_expired_dns_behavior, &_nw_parameters_get_expired_dns_behaviorErr, frameworkHandle, "nw_parameters_get_expired_dns_behavior", "10.14")
	registerFunc(&_nw_parameters_get_fast_open_enabled, &_nw_parameters_get_fast_open_enabledErr, frameworkHandle, "nw_parameters_get_fast_open_enabled", "10.14")
	registerFunc(&_nw_parameters_get_include_peer_to_peer, &_nw_parameters_get_include_peer_to_peerErr, frameworkHandle, "nw_parameters_get_include_peer_to_peer", "10.14")
	registerFunc(&_nw_parameters_get_local_only, &_nw_parameters_get_local_onlyErr, frameworkHandle, "nw_parameters_get_local_only", "10.14")
	registerFunc(&_nw_parameters_get_multipath_service, &_nw_parameters_get_multipath_serviceErr, frameworkHandle, "nw_parameters_get_multipath_service", "10.14")
	registerFunc(&_nw_parameters_get_prefer_no_proxy, &_nw_parameters_get_prefer_no_proxyErr, frameworkHandle, "nw_parameters_get_prefer_no_proxy", "10.14")
	registerFunc(&_nw_parameters_get_prohibit_constrained, &_nw_parameters_get_prohibit_constrainedErr, frameworkHandle, "nw_parameters_get_prohibit_constrained", "10.15")
	registerFunc(&_nw_parameters_get_prohibit_expensive, &_nw_parameters_get_prohibit_expensiveErr, frameworkHandle, "nw_parameters_get_prohibit_expensive", "10.14")
	registerFunc(&_nw_parameters_get_required_interface_type, &_nw_parameters_get_required_interface_typeErr, frameworkHandle, "nw_parameters_get_required_interface_type", "10.14")
	registerFunc(&_nw_parameters_get_reuse_local_address, &_nw_parameters_get_reuse_local_addressErr, frameworkHandle, "nw_parameters_get_reuse_local_address", "10.14")
	registerFunc(&_nw_parameters_get_service_class, &_nw_parameters_get_service_classErr, frameworkHandle, "nw_parameters_get_service_class", "10.14")
	registerFunc(&_nw_parameters_iterate_prohibited_interface_types, &_nw_parameters_iterate_prohibited_interface_typesErr, frameworkHandle, "nw_parameters_iterate_prohibited_interface_types", "10.14")
	registerFunc(&_nw_parameters_iterate_prohibited_interfaces, &_nw_parameters_iterate_prohibited_interfacesErr, frameworkHandle, "nw_parameters_iterate_prohibited_interfaces", "10.14")
	registerFunc(&_nw_parameters_prohibit_interface, &_nw_parameters_prohibit_interfaceErr, frameworkHandle, "nw_parameters_prohibit_interface", "10.14")
	registerFunc(&_nw_parameters_prohibit_interface_type, &_nw_parameters_prohibit_interface_typeErr, frameworkHandle, "nw_parameters_prohibit_interface_type", "10.14")
	registerFunc(&_nw_parameters_require_interface, &_nw_parameters_require_interfaceErr, frameworkHandle, "nw_parameters_require_interface", "10.14")
	registerFunc(&_nw_parameters_requires_dnssec_validation, &_nw_parameters_requires_dnssec_validationErr, frameworkHandle, "nw_parameters_requires_dnssec_validation", "13.0")
	registerFunc(&_nw_parameters_set_allow_ultra_constrained, &_nw_parameters_set_allow_ultra_constrainedErr, frameworkHandle, "nw_parameters_set_allow_ultra_constrained", "26.0")
	registerFunc(&_nw_parameters_set_attribution, &_nw_parameters_set_attributionErr, frameworkHandle, "nw_parameters_set_attribution", "12.0")
	registerFunc(&_nw_parameters_set_expired_dns_behavior, &_nw_parameters_set_expired_dns_behaviorErr, frameworkHandle, "nw_parameters_set_expired_dns_behavior", "10.14")
	registerFunc(&_nw_parameters_set_fast_open_enabled, &_nw_parameters_set_fast_open_enabledErr, frameworkHandle, "nw_parameters_set_fast_open_enabled", "10.14")
	registerFunc(&_nw_parameters_set_include_peer_to_peer, &_nw_parameters_set_include_peer_to_peerErr, frameworkHandle, "nw_parameters_set_include_peer_to_peer", "10.14")
	registerFunc(&_nw_parameters_set_local_endpoint, &_nw_parameters_set_local_endpointErr, frameworkHandle, "nw_parameters_set_local_endpoint", "10.14")
	registerFunc(&_nw_parameters_set_local_only, &_nw_parameters_set_local_onlyErr, frameworkHandle, "nw_parameters_set_local_only", "10.14")
	registerFunc(&_nw_parameters_set_multipath_service, &_nw_parameters_set_multipath_serviceErr, frameworkHandle, "nw_parameters_set_multipath_service", "10.14")
	registerFunc(&_nw_parameters_set_prefer_no_proxy, &_nw_parameters_set_prefer_no_proxyErr, frameworkHandle, "nw_parameters_set_prefer_no_proxy", "10.14")
	registerFunc(&_nw_parameters_set_privacy_context, &_nw_parameters_set_privacy_contextErr, frameworkHandle, "nw_parameters_set_privacy_context", "11.0")
	registerFunc(&_nw_parameters_set_prohibit_constrained, &_nw_parameters_set_prohibit_constrainedErr, frameworkHandle, "nw_parameters_set_prohibit_constrained", "10.15")
	registerFunc(&_nw_parameters_set_prohibit_expensive, &_nw_parameters_set_prohibit_expensiveErr, frameworkHandle, "nw_parameters_set_prohibit_expensive", "10.14")
	registerFunc(&_nw_parameters_set_required_interface_type, &_nw_parameters_set_required_interface_typeErr, frameworkHandle, "nw_parameters_set_required_interface_type", "10.14")
	registerFunc(&_nw_parameters_set_requires_dnssec_validation, &_nw_parameters_set_requires_dnssec_validationErr, frameworkHandle, "nw_parameters_set_requires_dnssec_validation", "13.0")
	registerFunc(&_nw_parameters_set_reuse_local_address, &_nw_parameters_set_reuse_local_addressErr, frameworkHandle, "nw_parameters_set_reuse_local_address", "10.14")
	registerFunc(&_nw_parameters_set_service_class, &_nw_parameters_set_service_classErr, frameworkHandle, "nw_parameters_set_service_class", "10.14")
	registerFunc(&_nw_path_copy_effective_local_endpoint, &_nw_path_copy_effective_local_endpointErr, frameworkHandle, "nw_path_copy_effective_local_endpoint", "10.14")
	registerFunc(&_nw_path_copy_effective_remote_endpoint, &_nw_path_copy_effective_remote_endpointErr, frameworkHandle, "nw_path_copy_effective_remote_endpoint", "10.14")
	registerFunc(&_nw_path_enumerate_gateways, &_nw_path_enumerate_gatewaysErr, frameworkHandle, "nw_path_enumerate_gateways", "10.15")
	registerFunc(&_nw_path_enumerate_interfaces, &_nw_path_enumerate_interfacesErr, frameworkHandle, "nw_path_enumerate_interfaces", "10.14")
	registerFunc(&_nw_path_get_link_quality, &_nw_path_get_link_qualityErr, frameworkHandle, "nw_path_get_link_quality", "26.0")
	registerFunc(&_nw_path_get_status, &_nw_path_get_statusErr, frameworkHandle, "nw_path_get_status", "10.14")
	registerFunc(&_nw_path_get_unsatisfied_reason, &_nw_path_get_unsatisfied_reasonErr, frameworkHandle, "nw_path_get_unsatisfied_reason", "11.0")
	registerFunc(&_nw_path_has_dns, &_nw_path_has_dnsErr, frameworkHandle, "nw_path_has_dns", "10.14")
	registerFunc(&_nw_path_has_ipv4, &_nw_path_has_ipv4Err, frameworkHandle, "nw_path_has_ipv4", "10.14")
	registerFunc(&_nw_path_has_ipv6, &_nw_path_has_ipv6Err, frameworkHandle, "nw_path_has_ipv6", "10.14")
	registerFunc(&_nw_path_is_constrained, &_nw_path_is_constrainedErr, frameworkHandle, "nw_path_is_constrained", "10.15")
	registerFunc(&_nw_path_is_equal, &_nw_path_is_equalErr, frameworkHandle, "nw_path_is_equal", "10.14")
	registerFunc(&_nw_path_is_expensive, &_nw_path_is_expensiveErr, frameworkHandle, "nw_path_is_expensive", "10.14")
	registerFunc(&_nw_path_is_ultra_constrained, &_nw_path_is_ultra_constrainedErr, frameworkHandle, "nw_path_is_ultra_constrained", "26.0")
	registerFunc(&_nw_path_monitor_cancel, &_nw_path_monitor_cancelErr, frameworkHandle, "nw_path_monitor_cancel", "10.14")
	registerFunc(&_nw_path_monitor_create, &_nw_path_monitor_createErr, frameworkHandle, "nw_path_monitor_create", "10.14")
	registerFunc(&_nw_path_monitor_create_for_ethernet_channel, &_nw_path_monitor_create_for_ethernet_channelErr, frameworkHandle, "nw_path_monitor_create_for_ethernet_channel", "13.0")
	registerFunc(&_nw_path_monitor_create_with_type, &_nw_path_monitor_create_with_typeErr, frameworkHandle, "nw_path_monitor_create_with_type", "10.14")
	registerFunc(&_nw_path_monitor_prohibit_interface_type, &_nw_path_monitor_prohibit_interface_typeErr, frameworkHandle, "nw_path_monitor_prohibit_interface_type", "11.0")
	registerFunc(&_nw_path_monitor_set_cancel_handler, &_nw_path_monitor_set_cancel_handlerErr, frameworkHandle, "nw_path_monitor_set_cancel_handler", "10.14")
	registerFunc(&_nw_path_monitor_set_queue, &_nw_path_monitor_set_queueErr, frameworkHandle, "nw_path_monitor_set_queue", "10.14")
	registerFunc(&_nw_path_monitor_set_update_handler, &_nw_path_monitor_set_update_handlerErr, frameworkHandle, "nw_path_monitor_set_update_handler", "10.14")
	registerFunc(&_nw_path_monitor_start, &_nw_path_monitor_startErr, frameworkHandle, "nw_path_monitor_start", "10.14")
	registerFunc(&_nw_path_uses_interface_type, &_nw_path_uses_interface_typeErr, frameworkHandle, "nw_path_uses_interface_type", "10.14")
	registerFunc(&_nw_privacy_context_add_proxy, &_nw_privacy_context_add_proxyErr, frameworkHandle, "nw_privacy_context_add_proxy", "14.0")
	registerFunc(&_nw_privacy_context_clear_proxies, &_nw_privacy_context_clear_proxiesErr, frameworkHandle, "nw_privacy_context_clear_proxies", "14.0")
	registerFunc(&_nw_privacy_context_create, &_nw_privacy_context_createErr, frameworkHandle, "nw_privacy_context_create", "11.0")
	registerFunc(&_nw_privacy_context_disable_logging, &_nw_privacy_context_disable_loggingErr, frameworkHandle, "nw_privacy_context_disable_logging", "11.0")
	registerFunc(&_nw_privacy_context_flush_cache, &_nw_privacy_context_flush_cacheErr, frameworkHandle, "nw_privacy_context_flush_cache", "11.0")
	registerFunc(&_nw_privacy_context_require_encrypted_name_resolution, &_nw_privacy_context_require_encrypted_name_resolutionErr, frameworkHandle, "nw_privacy_context_require_encrypted_name_resolution", "11.0")
	registerFunc(&_nw_protocol_copy_ip_definition, &_nw_protocol_copy_ip_definitionErr, frameworkHandle, "nw_protocol_copy_ip_definition", "10.14")
	registerFunc(&_nw_protocol_copy_quic_definition, &_nw_protocol_copy_quic_definitionErr, frameworkHandle, "nw_protocol_copy_quic_definition", "12.0")
	registerFunc(&_nw_protocol_copy_tcp_definition, &_nw_protocol_copy_tcp_definitionErr, frameworkHandle, "nw_protocol_copy_tcp_definition", "10.14")
	registerFunc(&_nw_protocol_copy_tls_definition, &_nw_protocol_copy_tls_definitionErr, frameworkHandle, "nw_protocol_copy_tls_definition", "10.14")
	registerFunc(&_nw_protocol_copy_udp_definition, &_nw_protocol_copy_udp_definitionErr, frameworkHandle, "nw_protocol_copy_udp_definition", "10.14")
	registerFunc(&_nw_protocol_copy_ws_definition, &_nw_protocol_copy_ws_definitionErr, frameworkHandle, "nw_protocol_copy_ws_definition", "10.15")
	registerFunc(&_nw_protocol_definition_is_equal, &_nw_protocol_definition_is_equalErr, frameworkHandle, "nw_protocol_definition_is_equal", "10.14")
	registerFunc(&_nw_protocol_metadata_copy_definition, &_nw_protocol_metadata_copy_definitionErr, frameworkHandle, "nw_protocol_metadata_copy_definition", "10.14")
	registerFunc(&_nw_protocol_metadata_is_framer_message, &_nw_protocol_metadata_is_framer_messageErr, frameworkHandle, "nw_protocol_metadata_is_framer_message", "10.15")
	registerFunc(&_nw_protocol_metadata_is_ip, &_nw_protocol_metadata_is_ipErr, frameworkHandle, "nw_protocol_metadata_is_ip", "10.14")
	registerFunc(&_nw_protocol_metadata_is_quic, &_nw_protocol_metadata_is_quicErr, frameworkHandle, "nw_protocol_metadata_is_quic", "12.0")
	registerFunc(&_nw_protocol_metadata_is_tcp, &_nw_protocol_metadata_is_tcpErr, frameworkHandle, "nw_protocol_metadata_is_tcp", "10.14")
	registerFunc(&_nw_protocol_metadata_is_tls, &_nw_protocol_metadata_is_tlsErr, frameworkHandle, "nw_protocol_metadata_is_tls", "10.14")
	registerFunc(&_nw_protocol_metadata_is_udp, &_nw_protocol_metadata_is_udpErr, frameworkHandle, "nw_protocol_metadata_is_udp", "10.14")
	registerFunc(&_nw_protocol_metadata_is_ws, &_nw_protocol_metadata_is_wsErr, frameworkHandle, "nw_protocol_metadata_is_ws", "10.15")
	registerFunc(&_nw_protocol_options_copy_definition, &_nw_protocol_options_copy_definitionErr, frameworkHandle, "nw_protocol_options_copy_definition", "10.14")
	registerFunc(&_nw_protocol_options_is_quic, &_nw_protocol_options_is_quicErr, frameworkHandle, "nw_protocol_options_is_quic", "12.0")
	registerFunc(&_nw_protocol_stack_clear_application_protocols, &_nw_protocol_stack_clear_application_protocolsErr, frameworkHandle, "nw_protocol_stack_clear_application_protocols", "10.14")
	registerFunc(&_nw_protocol_stack_copy_internet_protocol, &_nw_protocol_stack_copy_internet_protocolErr, frameworkHandle, "nw_protocol_stack_copy_internet_protocol", "10.14")
	registerFunc(&_nw_protocol_stack_copy_transport_protocol, &_nw_protocol_stack_copy_transport_protocolErr, frameworkHandle, "nw_protocol_stack_copy_transport_protocol", "10.14")
	registerFunc(&_nw_protocol_stack_iterate_application_protocols, &_nw_protocol_stack_iterate_application_protocolsErr, frameworkHandle, "nw_protocol_stack_iterate_application_protocols", "10.14")
	registerFunc(&_nw_protocol_stack_prepend_application_protocol, &_nw_protocol_stack_prepend_application_protocolErr, frameworkHandle, "nw_protocol_stack_prepend_application_protocol", "10.14")
	registerFunc(&_nw_protocol_stack_set_transport_protocol, &_nw_protocol_stack_set_transport_protocolErr, frameworkHandle, "nw_protocol_stack_set_transport_protocol", "10.14")
	registerFunc(&_nw_proxy_config_add_excluded_domain, &_nw_proxy_config_add_excluded_domainErr, frameworkHandle, "nw_proxy_config_add_excluded_domain", "14.0")
	registerFunc(&_nw_proxy_config_add_match_domain, &_nw_proxy_config_add_match_domainErr, frameworkHandle, "nw_proxy_config_add_match_domain", "14.0")
	registerFunc(&_nw_proxy_config_clear_excluded_domains, &_nw_proxy_config_clear_excluded_domainsErr, frameworkHandle, "nw_proxy_config_clear_excluded_domains", "14.0")
	registerFunc(&_nw_proxy_config_clear_match_domains, &_nw_proxy_config_clear_match_domainsErr, frameworkHandle, "nw_proxy_config_clear_match_domains", "14.0")
	registerFunc(&_nw_proxy_config_create_http_connect, &_nw_proxy_config_create_http_connectErr, frameworkHandle, "nw_proxy_config_create_http_connect", "14.0")
	registerFunc(&_nw_proxy_config_create_oblivious_http, &_nw_proxy_config_create_oblivious_httpErr, frameworkHandle, "nw_proxy_config_create_oblivious_http", "14.0")
	registerFunc(&_nw_proxy_config_create_relay, &_nw_proxy_config_create_relayErr, frameworkHandle, "nw_proxy_config_create_relay", "14.0")
	registerFunc(&_nw_proxy_config_create_socksv5, &_nw_proxy_config_create_socksv5Err, frameworkHandle, "nw_proxy_config_create_socksv5", "14.0")
	registerFunc(&_nw_proxy_config_enumerate_excluded_domains, &_nw_proxy_config_enumerate_excluded_domainsErr, frameworkHandle, "nw_proxy_config_enumerate_excluded_domains", "14.0")
	registerFunc(&_nw_proxy_config_enumerate_match_domains, &_nw_proxy_config_enumerate_match_domainsErr, frameworkHandle, "nw_proxy_config_enumerate_match_domains", "14.0")
	registerFunc(&_nw_proxy_config_get_failover_allowed, &_nw_proxy_config_get_failover_allowedErr, frameworkHandle, "nw_proxy_config_get_failover_allowed", "14.0")
	registerFunc(&_nw_proxy_config_set_failover_allowed, &_nw_proxy_config_set_failover_allowedErr, frameworkHandle, "nw_proxy_config_set_failover_allowed", "14.0")
	registerFunc(&_nw_proxy_config_set_username_and_password, &_nw_proxy_config_set_username_and_passwordErr, frameworkHandle, "nw_proxy_config_set_username_and_password", "14.0")
	registerFunc(&_nw_quic_add_tls_application_protocol, &_nw_quic_add_tls_application_protocolErr, frameworkHandle, "nw_quic_add_tls_application_protocol", "12.0")
	registerFunc(&_nw_quic_copy_sec_protocol_metadata, &_nw_quic_copy_sec_protocol_metadataErr, frameworkHandle, "nw_quic_copy_sec_protocol_metadata", "12.0")
	registerFunc(&_nw_quic_copy_sec_protocol_options, &_nw_quic_copy_sec_protocol_optionsErr, frameworkHandle, "nw_quic_copy_sec_protocol_options", "12.0")
	registerFunc(&_nw_quic_create_options, &_nw_quic_create_optionsErr, frameworkHandle, "nw_quic_create_options", "12.0")
	registerFunc(&_nw_quic_get_application_error, &_nw_quic_get_application_errorErr, frameworkHandle, "nw_quic_get_application_error", "12.0")
	registerFunc(&_nw_quic_get_application_error_reason, &_nw_quic_get_application_error_reasonErr, frameworkHandle, "nw_quic_get_application_error_reason", "12.0")
	registerFunc(&_nw_quic_get_idle_timeout, &_nw_quic_get_idle_timeoutErr, frameworkHandle, "nw_quic_get_idle_timeout", "12.0")
	registerFunc(&_nw_quic_get_initial_max_data, &_nw_quic_get_initial_max_dataErr, frameworkHandle, "nw_quic_get_initial_max_data", "12.0")
	registerFunc(&_nw_quic_get_initial_max_stream_data_bidirectional_local, &_nw_quic_get_initial_max_stream_data_bidirectional_localErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_local", "12.0")
	registerFunc(&_nw_quic_get_initial_max_stream_data_bidirectional_remote, &_nw_quic_get_initial_max_stream_data_bidirectional_remoteErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_remote", "12.0")
	registerFunc(&_nw_quic_get_initial_max_stream_data_unidirectional, &_nw_quic_get_initial_max_stream_data_unidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_unidirectional", "12.0")
	registerFunc(&_nw_quic_get_initial_max_streams_bidirectional, &_nw_quic_get_initial_max_streams_bidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_streams_bidirectional", "12.0")
	registerFunc(&_nw_quic_get_initial_max_streams_unidirectional, &_nw_quic_get_initial_max_streams_unidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_streams_unidirectional", "12.0")
	registerFunc(&_nw_quic_get_keepalive_interval, &_nw_quic_get_keepalive_intervalErr, frameworkHandle, "nw_quic_get_keepalive_interval", "12.0")
	registerFunc(&_nw_quic_get_local_max_streams_bidirectional, &_nw_quic_get_local_max_streams_bidirectionalErr, frameworkHandle, "nw_quic_get_local_max_streams_bidirectional", "12.0")
	registerFunc(&_nw_quic_get_local_max_streams_unidirectional, &_nw_quic_get_local_max_streams_unidirectionalErr, frameworkHandle, "nw_quic_get_local_max_streams_unidirectional", "12.0")
	registerFunc(&_nw_quic_get_max_datagram_frame_size, &_nw_quic_get_max_datagram_frame_sizeErr, frameworkHandle, "nw_quic_get_max_datagram_frame_size", "13.0")
	registerFunc(&_nw_quic_get_max_udp_payload_size, &_nw_quic_get_max_udp_payload_sizeErr, frameworkHandle, "nw_quic_get_max_udp_payload_size", "12.0")
	registerFunc(&_nw_quic_get_remote_idle_timeout, &_nw_quic_get_remote_idle_timeoutErr, frameworkHandle, "nw_quic_get_remote_idle_timeout", "12.0")
	registerFunc(&_nw_quic_get_remote_max_streams_bidirectional, &_nw_quic_get_remote_max_streams_bidirectionalErr, frameworkHandle, "nw_quic_get_remote_max_streams_bidirectional", "12.0")
	registerFunc(&_nw_quic_get_remote_max_streams_unidirectional, &_nw_quic_get_remote_max_streams_unidirectionalErr, frameworkHandle, "nw_quic_get_remote_max_streams_unidirectional", "12.0")
	registerFunc(&_nw_quic_get_stream_application_error, &_nw_quic_get_stream_application_errorErr, frameworkHandle, "nw_quic_get_stream_application_error", "12.0")
	registerFunc(&_nw_quic_get_stream_id, &_nw_quic_get_stream_idErr, frameworkHandle, "nw_quic_get_stream_id", "12.0")
	registerFunc(&_nw_quic_get_stream_is_datagram, &_nw_quic_get_stream_is_datagramErr, frameworkHandle, "nw_quic_get_stream_is_datagram", "13.0")
	registerFunc(&_nw_quic_get_stream_is_unidirectional, &_nw_quic_get_stream_is_unidirectionalErr, frameworkHandle, "nw_quic_get_stream_is_unidirectional", "12.0")
	registerFunc(&_nw_quic_get_stream_type, &_nw_quic_get_stream_typeErr, frameworkHandle, "nw_quic_get_stream_type", "12.0")
	registerFunc(&_nw_quic_get_stream_usable_datagram_frame_size, &_nw_quic_get_stream_usable_datagram_frame_sizeErr, frameworkHandle, "nw_quic_get_stream_usable_datagram_frame_size", "13.0")
	registerFunc(&_nw_quic_set_application_error, &_nw_quic_set_application_errorErr, frameworkHandle, "nw_quic_set_application_error", "12.0")
	registerFunc(&_nw_quic_set_idle_timeout, &_nw_quic_set_idle_timeoutErr, frameworkHandle, "nw_quic_set_idle_timeout", "12.0")
	registerFunc(&_nw_quic_set_initial_max_data, &_nw_quic_set_initial_max_dataErr, frameworkHandle, "nw_quic_set_initial_max_data", "12.0")
	registerFunc(&_nw_quic_set_initial_max_stream_data_bidirectional_local, &_nw_quic_set_initial_max_stream_data_bidirectional_localErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_local", "12.0")
	registerFunc(&_nw_quic_set_initial_max_stream_data_bidirectional_remote, &_nw_quic_set_initial_max_stream_data_bidirectional_remoteErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_remote", "12.0")
	registerFunc(&_nw_quic_set_initial_max_stream_data_unidirectional, &_nw_quic_set_initial_max_stream_data_unidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_unidirectional", "12.0")
	registerFunc(&_nw_quic_set_initial_max_streams_bidirectional, &_nw_quic_set_initial_max_streams_bidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_streams_bidirectional", "12.0")
	registerFunc(&_nw_quic_set_initial_max_streams_unidirectional, &_nw_quic_set_initial_max_streams_unidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_streams_unidirectional", "12.0")
	registerFunc(&_nw_quic_set_keepalive_interval, &_nw_quic_set_keepalive_intervalErr, frameworkHandle, "nw_quic_set_keepalive_interval", "12.0")
	registerFunc(&_nw_quic_set_local_max_streams_bidirectional, &_nw_quic_set_local_max_streams_bidirectionalErr, frameworkHandle, "nw_quic_set_local_max_streams_bidirectional", "12.0")
	registerFunc(&_nw_quic_set_local_max_streams_unidirectional, &_nw_quic_set_local_max_streams_unidirectionalErr, frameworkHandle, "nw_quic_set_local_max_streams_unidirectional", "12.0")
	registerFunc(&_nw_quic_set_max_datagram_frame_size, &_nw_quic_set_max_datagram_frame_sizeErr, frameworkHandle, "nw_quic_set_max_datagram_frame_size", "13.0")
	registerFunc(&_nw_quic_set_max_udp_payload_size, &_nw_quic_set_max_udp_payload_sizeErr, frameworkHandle, "nw_quic_set_max_udp_payload_size", "12.0")
	registerFunc(&_nw_quic_set_stream_application_error, &_nw_quic_set_stream_application_errorErr, frameworkHandle, "nw_quic_set_stream_application_error", "12.0")
	registerFunc(&_nw_quic_set_stream_is_datagram, &_nw_quic_set_stream_is_datagramErr, frameworkHandle, "nw_quic_set_stream_is_datagram", "13.0")
	registerFunc(&_nw_quic_set_stream_is_unidirectional, &_nw_quic_set_stream_is_unidirectionalErr, frameworkHandle, "nw_quic_set_stream_is_unidirectional", "12.0")
	registerFunc(&_nw_relay_hop_add_additional_http_header_field, &_nw_relay_hop_add_additional_http_header_fieldErr, frameworkHandle, "nw_relay_hop_add_additional_http_header_field", "14.0")
	registerFunc(&_nw_relay_hop_create, &_nw_relay_hop_createErr, frameworkHandle, "nw_relay_hop_create", "14.0")
	registerFunc(&_nw_release, &_nw_releaseErr, frameworkHandle, "nw_release", "10.14")
	registerFunc(&_nw_resolution_report_copy_preferred_endpoint, &_nw_resolution_report_copy_preferred_endpointErr, frameworkHandle, "nw_resolution_report_copy_preferred_endpoint", "11.0")
	registerFunc(&_nw_resolution_report_copy_successful_endpoint, &_nw_resolution_report_copy_successful_endpointErr, frameworkHandle, "nw_resolution_report_copy_successful_endpoint", "11.0")
	registerFunc(&_nw_resolution_report_get_endpoint_count, &_nw_resolution_report_get_endpoint_countErr, frameworkHandle, "nw_resolution_report_get_endpoint_count", "11.0")
	registerFunc(&_nw_resolution_report_get_milliseconds, &_nw_resolution_report_get_millisecondsErr, frameworkHandle, "nw_resolution_report_get_milliseconds", "11.0")
	registerFunc(&_nw_resolution_report_get_protocol, &_nw_resolution_report_get_protocolErr, frameworkHandle, "nw_resolution_report_get_protocol", "11.0")
	registerFunc(&_nw_resolution_report_get_source, &_nw_resolution_report_get_sourceErr, frameworkHandle, "nw_resolution_report_get_source", "11.0")
	registerFunc(&_nw_resolver_config_add_server_address, &_nw_resolver_config_add_server_addressErr, frameworkHandle, "nw_resolver_config_add_server_address", "11.0")
	registerFunc(&_nw_resolver_config_create_https, &_nw_resolver_config_create_httpsErr, frameworkHandle, "nw_resolver_config_create_https", "11.0")
	registerFunc(&_nw_resolver_config_create_tls, &_nw_resolver_config_create_tlsErr, frameworkHandle, "nw_resolver_config_create_tls", "11.0")
	registerFunc(&_nw_retain, &_nw_retainErr, frameworkHandle, "nw_retain", "10.14")
	registerFunc(&_nw_tcp_create_options, &_nw_tcp_create_optionsErr, frameworkHandle, "nw_tcp_create_options", "10.14")
	registerFunc(&_nw_tcp_get_available_receive_buffer, &_nw_tcp_get_available_receive_bufferErr, frameworkHandle, "nw_tcp_get_available_receive_buffer", "10.14")
	registerFunc(&_nw_tcp_get_available_send_buffer, &_nw_tcp_get_available_send_bufferErr, frameworkHandle, "nw_tcp_get_available_send_buffer", "10.14")
	registerFunc(&_nw_tcp_options_set_connection_timeout, &_nw_tcp_options_set_connection_timeoutErr, frameworkHandle, "nw_tcp_options_set_connection_timeout", "10.14")
	registerFunc(&_nw_tcp_options_set_disable_ack_stretching, &_nw_tcp_options_set_disable_ack_stretchingErr, frameworkHandle, "nw_tcp_options_set_disable_ack_stretching", "10.14")
	registerFunc(&_nw_tcp_options_set_disable_ecn, &_nw_tcp_options_set_disable_ecnErr, frameworkHandle, "nw_tcp_options_set_disable_ecn", "10.14")
	registerFunc(&_nw_tcp_options_set_enable_fast_open, &_nw_tcp_options_set_enable_fast_openErr, frameworkHandle, "nw_tcp_options_set_enable_fast_open", "10.14")
	registerFunc(&_nw_tcp_options_set_enable_keepalive, &_nw_tcp_options_set_enable_keepaliveErr, frameworkHandle, "nw_tcp_options_set_enable_keepalive", "10.14")
	registerFunc(&_nw_tcp_options_set_keepalive_count, &_nw_tcp_options_set_keepalive_countErr, frameworkHandle, "nw_tcp_options_set_keepalive_count", "10.14")
	registerFunc(&_nw_tcp_options_set_keepalive_idle_time, &_nw_tcp_options_set_keepalive_idle_timeErr, frameworkHandle, "nw_tcp_options_set_keepalive_idle_time", "10.14")
	registerFunc(&_nw_tcp_options_set_keepalive_interval, &_nw_tcp_options_set_keepalive_intervalErr, frameworkHandle, "nw_tcp_options_set_keepalive_interval", "10.14")
	registerFunc(&_nw_tcp_options_set_maximum_segment_size, &_nw_tcp_options_set_maximum_segment_sizeErr, frameworkHandle, "nw_tcp_options_set_maximum_segment_size", "10.14")
	registerFunc(&_nw_tcp_options_set_multipath_force_version, &_nw_tcp_options_set_multipath_force_versionErr, frameworkHandle, "nw_tcp_options_set_multipath_force_version", "12.0")
	registerFunc(&_nw_tcp_options_set_no_delay, &_nw_tcp_options_set_no_delayErr, frameworkHandle, "nw_tcp_options_set_no_delay", "10.14")
	registerFunc(&_nw_tcp_options_set_no_options, &_nw_tcp_options_set_no_optionsErr, frameworkHandle, "nw_tcp_options_set_no_options", "10.14")
	registerFunc(&_nw_tcp_options_set_no_push, &_nw_tcp_options_set_no_pushErr, frameworkHandle, "nw_tcp_options_set_no_push", "10.14")
	registerFunc(&_nw_tcp_options_set_persist_timeout, &_nw_tcp_options_set_persist_timeoutErr, frameworkHandle, "nw_tcp_options_set_persist_timeout", "10.14")
	registerFunc(&_nw_tcp_options_set_retransmit_connection_drop_time, &_nw_tcp_options_set_retransmit_connection_drop_timeErr, frameworkHandle, "nw_tcp_options_set_retransmit_connection_drop_time", "10.14")
	registerFunc(&_nw_tcp_options_set_retransmit_fin_drop, &_nw_tcp_options_set_retransmit_fin_dropErr, frameworkHandle, "nw_tcp_options_set_retransmit_fin_drop", "10.14")
	registerFunc(&_nw_tls_copy_sec_protocol_metadata, &_nw_tls_copy_sec_protocol_metadataErr, frameworkHandle, "nw_tls_copy_sec_protocol_metadata", "10.14")
	registerFunc(&_nw_tls_copy_sec_protocol_options, &_nw_tls_copy_sec_protocol_optionsErr, frameworkHandle, "nw_tls_copy_sec_protocol_options", "10.14")
	registerFunc(&_nw_tls_create_options, &_nw_tls_create_optionsErr, frameworkHandle, "nw_tls_create_options", "10.14")
	registerFunc(&_nw_txt_record_access_bytes, &_nw_txt_record_access_bytesErr, frameworkHandle, "nw_txt_record_access_bytes", "10.15")
	registerFunc(&_nw_txt_record_access_key, &_nw_txt_record_access_keyErr, frameworkHandle, "nw_txt_record_access_key", "10.15")
	registerFunc(&_nw_txt_record_apply, &_nw_txt_record_applyErr, frameworkHandle, "nw_txt_record_apply", "10.15")
	registerFunc(&_nw_txt_record_copy, &_nw_txt_record_copyErr, frameworkHandle, "nw_txt_record_copy", "10.15")
	registerFunc(&_nw_txt_record_create_dictionary, &_nw_txt_record_create_dictionaryErr, frameworkHandle, "nw_txt_record_create_dictionary", "10.15")
	registerFunc(&_nw_txt_record_create_with_bytes, &_nw_txt_record_create_with_bytesErr, frameworkHandle, "nw_txt_record_create_with_bytes", "10.15")
	registerFunc(&_nw_txt_record_find_key, &_nw_txt_record_find_keyErr, frameworkHandle, "nw_txt_record_find_key", "10.15")
	registerFunc(&_nw_txt_record_get_key_count, &_nw_txt_record_get_key_countErr, frameworkHandle, "nw_txt_record_get_key_count", "10.15")
	registerFunc(&_nw_txt_record_is_dictionary, &_nw_txt_record_is_dictionaryErr, frameworkHandle, "nw_txt_record_is_dictionary", "10.15")
	registerFunc(&_nw_txt_record_is_equal, &_nw_txt_record_is_equalErr, frameworkHandle, "nw_txt_record_is_equal", "10.15")
	registerFunc(&_nw_txt_record_remove_key, &_nw_txt_record_remove_keyErr, frameworkHandle, "nw_txt_record_remove_key", "10.15")
	registerFunc(&_nw_txt_record_set_key, &_nw_txt_record_set_keyErr, frameworkHandle, "nw_txt_record_set_key", "10.15")
	registerFunc(&_nw_udp_create_metadata, &_nw_udp_create_metadataErr, frameworkHandle, "nw_udp_create_metadata", "10.14")
	registerFunc(&_nw_udp_create_options, &_nw_udp_create_optionsErr, frameworkHandle, "nw_udp_create_options", "10.14")
	registerFunc(&_nw_udp_options_set_prefer_no_checksum, &_nw_udp_options_set_prefer_no_checksumErr, frameworkHandle, "nw_udp_options_set_prefer_no_checksum", "10.14")
	registerFunc(&_nw_ws_create_metadata, &_nw_ws_create_metadataErr, frameworkHandle, "nw_ws_create_metadata", "10.15")
	registerFunc(&_nw_ws_create_options, &_nw_ws_create_optionsErr, frameworkHandle, "nw_ws_create_options", "10.15")
	registerFunc(&_nw_ws_metadata_copy_server_response, &_nw_ws_metadata_copy_server_responseErr, frameworkHandle, "nw_ws_metadata_copy_server_response", "10.15")
	registerFunc(&_nw_ws_metadata_get_close_code, &_nw_ws_metadata_get_close_codeErr, frameworkHandle, "nw_ws_metadata_get_close_code", "10.15")
	registerFunc(&_nw_ws_metadata_get_opcode, &_nw_ws_metadata_get_opcodeErr, frameworkHandle, "nw_ws_metadata_get_opcode", "10.15")
	registerFunc(&_nw_ws_metadata_set_close_code, &_nw_ws_metadata_set_close_codeErr, frameworkHandle, "nw_ws_metadata_set_close_code", "10.15")
	registerFunc(&_nw_ws_metadata_set_pong_handler, &_nw_ws_metadata_set_pong_handlerErr, frameworkHandle, "nw_ws_metadata_set_pong_handler", "10.15")
	registerFunc(&_nw_ws_options_add_additional_header, &_nw_ws_options_add_additional_headerErr, frameworkHandle, "nw_ws_options_add_additional_header", "10.15")
	registerFunc(&_nw_ws_options_add_subprotocol, &_nw_ws_options_add_subprotocolErr, frameworkHandle, "nw_ws_options_add_subprotocol", "10.15")
	registerFunc(&_nw_ws_options_set_auto_reply_ping, &_nw_ws_options_set_auto_reply_pingErr, frameworkHandle, "nw_ws_options_set_auto_reply_ping", "10.15")
	registerFunc(&_nw_ws_options_set_client_request_handler, &_nw_ws_options_set_client_request_handlerErr, frameworkHandle, "nw_ws_options_set_client_request_handler", "10.15")
	registerFunc(&_nw_ws_options_set_maximum_message_size, &_nw_ws_options_set_maximum_message_sizeErr, frameworkHandle, "nw_ws_options_set_maximum_message_size", "10.15")
	registerFunc(&_nw_ws_options_set_skip_handshake, &_nw_ws_options_set_skip_handshakeErr, frameworkHandle, "nw_ws_options_set_skip_handshake", "10.15")
	registerFunc(&_nw_ws_request_enumerate_additional_headers, &_nw_ws_request_enumerate_additional_headersErr, frameworkHandle, "nw_ws_request_enumerate_additional_headers", "10.15")
	registerFunc(&_nw_ws_request_enumerate_subprotocols, &_nw_ws_request_enumerate_subprotocolsErr, frameworkHandle, "nw_ws_request_enumerate_subprotocols", "10.15")
	registerFunc(&_nw_ws_response_add_additional_header, &_nw_ws_response_add_additional_headerErr, frameworkHandle, "nw_ws_response_add_additional_header", "10.15")
	registerFunc(&_nw_ws_response_create, &_nw_ws_response_createErr, frameworkHandle, "nw_ws_response_create", "10.15")
	registerFunc(&_nw_ws_response_enumerate_additional_headers, &_nw_ws_response_enumerate_additional_headersErr, frameworkHandle, "nw_ws_response_enumerate_additional_headers", "10.15")
	registerFunc(&_nw_ws_response_get_selected_subprotocol, &_nw_ws_response_get_selected_subprotocolErr, frameworkHandle, "nw_ws_response_get_selected_subprotocol", "10.15")
	registerFunc(&_nw_ws_response_get_status, &_nw_ws_response_get_statusErr, frameworkHandle, "nw_ws_response_get_status", "10.15")
}
