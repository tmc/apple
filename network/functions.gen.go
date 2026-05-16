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

var networkCreateSecureTCPForPlain func(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) NWParameters
var networkCreateSecureTCPForPlainErr error

func tryNWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if networkCreateSecureTCPForPlain == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_secure_tcp", "10.14", networkCreateSecureTCPForPlainErr)
	}
	if _nw_parameters_configure_protocol_disableSymbol == 0 {
		return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_disable", "10.14", _nw_parameters_configure_protocol_disableErr)
	}
	var _block1 unsafe.Pointer
	if configureTCP == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configureTCP(objectivec.ObjectFromID(blockArg0)) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return networkCreateSecureTCPForPlain(networkProtocolBlockValue(_nw_parameters_configure_protocol_disableSymbol), _block1), nil
}

// TryNWParametersCreatePlainTCP initializes parameters for cleartext TCP.
func TryNWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) (NWParameters, error) {
	return tryNWParametersCreatePlainTCP(configureTCP)
}

// NWParametersCreatePlainTCP initializes parameters for cleartext TCP.
func NWParametersCreatePlainTCP(configureTCP NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreatePlainTCP(configureTCP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorCopyTXTRecordObject func(advertise_descriptor NWAdvertiseDescriptor) NWTXTRecord
var _nWAdvertiseDescriptorCopyTXTRecordObjectErr error

func tryNWAdvertiseDescriptorCopyTXTRecordObject(advertise_descriptor NWAdvertiseDescriptor) (NWTXTRecord, error) {
	if _nWAdvertiseDescriptorCopyTXTRecordObject == nil {
		return *new(NWTXTRecord), symbolCallError("nw_advertise_descriptor_copy_txt_record_object", "10.15", _nWAdvertiseDescriptorCopyTXTRecordObjectErr)
	}
	return _nWAdvertiseDescriptorCopyTXTRecordObject(advertise_descriptor), nil
}

// NWAdvertiseDescriptorCopyTXTRecordObject accesses the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_copy_txt_record_object(_:)
func NWAdvertiseDescriptorCopyTXTRecordObject(advertise_descriptor NWAdvertiseDescriptor) NWTXTRecord {
	result, callErr := tryNWAdvertiseDescriptorCopyTXTRecordObject(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorCreateApplicationService func(application_service_name string) NWAdvertiseDescriptor
var _nWAdvertiseDescriptorCreateApplicationServiceErr error

func tryNWAdvertiseDescriptorCreateApplicationService(application_service_name string) (NWAdvertiseDescriptor, error) {
	if _nWAdvertiseDescriptorCreateApplicationService == nil {
		return *new(NWAdvertiseDescriptor), symbolCallError("nw_advertise_descriptor_create_application_service", "13.0", _nWAdvertiseDescriptorCreateApplicationServiceErr)
	}
	return _nWAdvertiseDescriptorCreateApplicationService(application_service_name), nil
}

// NWAdvertiseDescriptorCreateApplicationService.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_application_service(_:)
func NWAdvertiseDescriptorCreateApplicationService(application_service_name string) NWAdvertiseDescriptor {
	result, callErr := tryNWAdvertiseDescriptorCreateApplicationService(application_service_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorCreateBonjourService func(name string, type_ string, domain string) NWAdvertiseDescriptor
var _nWAdvertiseDescriptorCreateBonjourServiceErr error

func tryNWAdvertiseDescriptorCreateBonjourService(name string, type_ string, domain string) (NWAdvertiseDescriptor, error) {
	if _nWAdvertiseDescriptorCreateBonjourService == nil {
		return *new(NWAdvertiseDescriptor), symbolCallError("nw_advertise_descriptor_create_bonjour_service", "10.14", _nWAdvertiseDescriptorCreateBonjourServiceErr)
	}
	return _nWAdvertiseDescriptorCreateBonjourService(name, type_, domain), nil
}

// NWAdvertiseDescriptorCreateBonjourService initializes a Bonjour service to advertise.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_bonjour_service(_:_:_:)
func NWAdvertiseDescriptorCreateBonjourService(name string, type_ string, domain string) NWAdvertiseDescriptor {
	result, callErr := tryNWAdvertiseDescriptorCreateBonjourService(name, type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorGetApplicationServiceName func(advertise_descriptor NWAdvertiseDescriptor) *byte
var _nWAdvertiseDescriptorGetApplicationServiceNameErr error

func tryNWAdvertiseDescriptorGetApplicationServiceName(advertise_descriptor NWAdvertiseDescriptor) (*byte, error) {
	if _nWAdvertiseDescriptorGetApplicationServiceName == nil {
		return nil, symbolCallError("nw_advertise_descriptor_get_application_service_name", "13.0", _nWAdvertiseDescriptorGetApplicationServiceNameErr)
	}
	return _nWAdvertiseDescriptorGetApplicationServiceName(advertise_descriptor), nil
}

// NWAdvertiseDescriptorGetApplicationServiceName.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_application_service_name(_:)
func NWAdvertiseDescriptorGetApplicationServiceName(advertise_descriptor NWAdvertiseDescriptor) *byte {
	result, callErr := tryNWAdvertiseDescriptorGetApplicationServiceName(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorGetNoAutoRename func(advertise_descriptor NWAdvertiseDescriptor) bool
var _nWAdvertiseDescriptorGetNoAutoRenameErr error

func tryNWAdvertiseDescriptorGetNoAutoRename(advertise_descriptor NWAdvertiseDescriptor) (bool, error) {
	if _nWAdvertiseDescriptorGetNoAutoRename == nil {
		return false, symbolCallError("nw_advertise_descriptor_get_no_auto_rename", "10.14", _nWAdvertiseDescriptorGetNoAutoRenameErr)
	}
	return _nWAdvertiseDescriptorGetNoAutoRename(advertise_descriptor), nil
}

// NWAdvertiseDescriptorGetNoAutoRename checks whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_no_auto_rename(_:)
func NWAdvertiseDescriptorGetNoAutoRename(advertise_descriptor NWAdvertiseDescriptor) bool {
	result, callErr := tryNWAdvertiseDescriptorGetNoAutoRename(advertise_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWAdvertiseDescriptorSetNoAutoRename func(advertise_descriptor NWAdvertiseDescriptor, no_auto_rename bool)
var _nWAdvertiseDescriptorSetNoAutoRenameErr error

func tryNWAdvertiseDescriptorSetNoAutoRename(advertise_descriptor NWAdvertiseDescriptor, no_auto_rename bool) error {
	if _nWAdvertiseDescriptorSetNoAutoRename == nil {
		return symbolCallError("nw_advertise_descriptor_set_no_auto_rename", "10.14", _nWAdvertiseDescriptorSetNoAutoRenameErr)
	}
	_nWAdvertiseDescriptorSetNoAutoRename(advertise_descriptor, no_auto_rename)
	return nil
}

// NWAdvertiseDescriptorSetNoAutoRename sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_no_auto_rename(_:_:)
func NWAdvertiseDescriptorSetNoAutoRename(advertise_descriptor NWAdvertiseDescriptor, no_auto_rename bool) {
	if callErr := tryNWAdvertiseDescriptorSetNoAutoRename(advertise_descriptor, no_auto_rename); callErr != nil {
		panic(callErr)
	}
}

var _nWAdvertiseDescriptorSetTXTRecord func(advertise_descriptor NWAdvertiseDescriptor, txt_record unsafe.Pointer, txt_length uintptr)
var _nWAdvertiseDescriptorSetTXTRecordErr error

func tryNWAdvertiseDescriptorSetTXTRecord(advertise_descriptor NWAdvertiseDescriptor, txt_record unsafe.Pointer, txt_length uintptr) error {
	if _nWAdvertiseDescriptorSetTXTRecord == nil {
		return symbolCallError("nw_advertise_descriptor_set_txt_record", "10.14", _nWAdvertiseDescriptorSetTXTRecordErr)
	}
	_nWAdvertiseDescriptorSetTXTRecord(advertise_descriptor, txt_record, txt_length)
	return nil
}

// NWAdvertiseDescriptorSetTXTRecord sets the TXT record as a raw buffer to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record(_:_:_:)
func NWAdvertiseDescriptorSetTXTRecord(advertise_descriptor NWAdvertiseDescriptor, txt_record unsafe.Pointer, txt_length uintptr) {
	if callErr := tryNWAdvertiseDescriptorSetTXTRecord(advertise_descriptor, txt_record, txt_length); callErr != nil {
		panic(callErr)
	}
}

var _nWAdvertiseDescriptorSetTXTRecordObject func(advertise_descriptor NWAdvertiseDescriptor, txt_record NWTXTRecord)
var _nWAdvertiseDescriptorSetTXTRecordObjectErr error

func tryNWAdvertiseDescriptorSetTXTRecordObject(advertise_descriptor NWAdvertiseDescriptor, txt_record NWTXTRecord) error {
	if _nWAdvertiseDescriptorSetTXTRecordObject == nil {
		return symbolCallError("nw_advertise_descriptor_set_txt_record_object", "10.15", _nWAdvertiseDescriptorSetTXTRecordObjectErr)
	}
	_nWAdvertiseDescriptorSetTXTRecordObject(advertise_descriptor, txt_record)
	return nil
}

// NWAdvertiseDescriptorSetTXTRecordObject sets the TXT record to advertise with the service.
//
// See: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record_object(_:_:)
func NWAdvertiseDescriptorSetTXTRecordObject(advertise_descriptor NWAdvertiseDescriptor, txt_record NWTXTRecord) {
	if callErr := tryNWAdvertiseDescriptorSetTXTRecordObject(advertise_descriptor, txt_record); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowseDescriptorCreateApplicationService func(application_service_name string) NWBrowseDescriptor
var _nWBrowseDescriptorCreateApplicationServiceErr error

func tryNWBrowseDescriptorCreateApplicationService(application_service_name string) (NWBrowseDescriptor, error) {
	if _nWBrowseDescriptorCreateApplicationService == nil {
		return *new(NWBrowseDescriptor), symbolCallError("nw_browse_descriptor_create_application_service", "13.0", _nWBrowseDescriptorCreateApplicationServiceErr)
	}
	return _nWBrowseDescriptorCreateApplicationService(application_service_name), nil
}

// NWBrowseDescriptorCreateApplicationService.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_application_service(_:)
func NWBrowseDescriptorCreateApplicationService(application_service_name string) NWBrowseDescriptor {
	result, callErr := tryNWBrowseDescriptorCreateApplicationService(application_service_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorCreateBonjourService func(type_ string, domain string) NWBrowseDescriptor
var _nWBrowseDescriptorCreateBonjourServiceErr error

func tryNWBrowseDescriptorCreateBonjourService(type_ string, domain string) (NWBrowseDescriptor, error) {
	if _nWBrowseDescriptorCreateBonjourService == nil {
		return *new(NWBrowseDescriptor), symbolCallError("nw_browse_descriptor_create_bonjour_service", "10.15", _nWBrowseDescriptorCreateBonjourServiceErr)
	}
	return _nWBrowseDescriptorCreateBonjourService(type_, domain), nil
}

// NWBrowseDescriptorCreateBonjourService initializes a service descriptor used to discover a Bonjour service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_bonjour_service(_:_:)
func NWBrowseDescriptorCreateBonjourService(type_ string, domain string) NWBrowseDescriptor {
	result, callErr := tryNWBrowseDescriptorCreateBonjourService(type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorGetApplicationServiceName func(descriptor NWBrowseDescriptor) *byte
var _nWBrowseDescriptorGetApplicationServiceNameErr error

func tryNWBrowseDescriptorGetApplicationServiceName(descriptor NWBrowseDescriptor) (*byte, error) {
	if _nWBrowseDescriptorGetApplicationServiceName == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_application_service_name", "13.0", _nWBrowseDescriptorGetApplicationServiceNameErr)
	}
	return _nWBrowseDescriptorGetApplicationServiceName(descriptor), nil
}

// NWBrowseDescriptorGetApplicationServiceName.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_application_service_name(_:)
func NWBrowseDescriptorGetApplicationServiceName(descriptor NWBrowseDescriptor) *byte {
	result, callErr := tryNWBrowseDescriptorGetApplicationServiceName(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorGetBonjourServiceDomain func(descriptor NWBrowseDescriptor) *byte
var _nWBrowseDescriptorGetBonjourServiceDomainErr error

func tryNWBrowseDescriptorGetBonjourServiceDomain(descriptor NWBrowseDescriptor) (*byte, error) {
	if _nWBrowseDescriptorGetBonjourServiceDomain == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_bonjour_service_domain", "10.15", _nWBrowseDescriptorGetBonjourServiceDomainErr)
	}
	return _nWBrowseDescriptorGetBonjourServiceDomain(descriptor), nil
}

// NWBrowseDescriptorGetBonjourServiceDomain accesses the Bonjour service domain set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_domain(_:)
func NWBrowseDescriptorGetBonjourServiceDomain(descriptor NWBrowseDescriptor) *byte {
	result, callErr := tryNWBrowseDescriptorGetBonjourServiceDomain(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorGetBonjourServiceType func(descriptor NWBrowseDescriptor) *byte
var _nWBrowseDescriptorGetBonjourServiceTypeErr error

func tryNWBrowseDescriptorGetBonjourServiceType(descriptor NWBrowseDescriptor) (*byte, error) {
	if _nWBrowseDescriptorGetBonjourServiceType == nil {
		return nil, symbolCallError("nw_browse_descriptor_get_bonjour_service_type", "10.15", _nWBrowseDescriptorGetBonjourServiceTypeErr)
	}
	return _nWBrowseDescriptorGetBonjourServiceType(descriptor), nil
}

// NWBrowseDescriptorGetBonjourServiceType accesses the Bonjour service type set on a browse descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_type(_:)
func NWBrowseDescriptorGetBonjourServiceType(descriptor NWBrowseDescriptor) *byte {
	result, callErr := tryNWBrowseDescriptorGetBonjourServiceType(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorGetIncludeTXTRecord func(descriptor NWBrowseDescriptor) bool
var _nWBrowseDescriptorGetIncludeTXTRecordErr error

func tryNWBrowseDescriptorGetIncludeTXTRecord(descriptor NWBrowseDescriptor) (bool, error) {
	if _nWBrowseDescriptorGetIncludeTXTRecord == nil {
		return false, symbolCallError("nw_browse_descriptor_get_include_txt_record", "10.15", _nWBrowseDescriptorGetIncludeTXTRecordErr)
	}
	return _nWBrowseDescriptorGetIncludeTXTRecord(descriptor), nil
}

// NWBrowseDescriptorGetIncludeTXTRecord checks if the browse descriptor requires including associated TXT records with all results.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_include_txt_record(_:)
func NWBrowseDescriptorGetIncludeTXTRecord(descriptor NWBrowseDescriptor) bool {
	result, callErr := tryNWBrowseDescriptorGetIncludeTXTRecord(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseDescriptorSetIncludeTXTRecord func(descriptor NWBrowseDescriptor, include_txt_record bool)
var _nWBrowseDescriptorSetIncludeTXTRecordErr error

func tryNWBrowseDescriptorSetIncludeTXTRecord(descriptor NWBrowseDescriptor, include_txt_record bool) error {
	if _nWBrowseDescriptorSetIncludeTXTRecord == nil {
		return symbolCallError("nw_browse_descriptor_set_include_txt_record", "10.15", _nWBrowseDescriptorSetIncludeTXTRecordErr)
	}
	_nWBrowseDescriptorSetIncludeTXTRecord(descriptor, include_txt_record)
	return nil
}

// NWBrowseDescriptorSetIncludeTXTRecord requires including associated TXT records with all results generated for this service descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_descriptor_set_include_txt_record(_:_:)
func NWBrowseDescriptorSetIncludeTXTRecord(descriptor NWBrowseDescriptor, include_txt_record bool) {
	if callErr := tryNWBrowseDescriptorSetIncludeTXTRecord(descriptor, include_txt_record); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowseResultCopyEndpoint func(result NWBrowseResult) NWEndpoint
var _nWBrowseResultCopyEndpointErr error

func tryNWBrowseResultCopyEndpoint(result NWBrowseResult) (NWEndpoint, error) {
	if _nWBrowseResultCopyEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_browse_result_copy_endpoint", "10.15", _nWBrowseResultCopyEndpointErr)
	}
	return _nWBrowseResultCopyEndpoint(result), nil
}

// NWBrowseResultCopyEndpoint the discovered service endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_endpoint(_:)
func NWBrowseResultCopyEndpoint(result NWBrowseResult) NWEndpoint {
	result0, callErr := tryNWBrowseResultCopyEndpoint(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nWBrowseResultCopyTXTRecordObject func(result NWBrowseResult) NWTXTRecord
var _nWBrowseResultCopyTXTRecordObjectErr error

func tryNWBrowseResultCopyTXTRecordObject(result NWBrowseResult) (NWTXTRecord, error) {
	if _nWBrowseResultCopyTXTRecordObject == nil {
		return *new(NWTXTRecord), symbolCallError("nw_browse_result_copy_txt_record_object", "10.15", _nWBrowseResultCopyTXTRecordObjectErr)
	}
	return _nWBrowseResultCopyTXTRecordObject(result), nil
}

// NWBrowseResultCopyTXTRecordObject accesses the TXT record associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_copy_txt_record_object(_:)
func NWBrowseResultCopyTXTRecordObject(result NWBrowseResult) NWTXTRecord {
	result0, callErr := tryNWBrowseResultCopyTXTRecordObject(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nWBrowseResultEnumerateInterfaces func(result NWBrowseResult, enumerator unsafe.Pointer)
var _nWBrowseResultEnumerateInterfacesErr error

func tryNWBrowseResultEnumerateInterfaces(result NWBrowseResult, enumerator NWBrowseResultEnumerateInterface) error {
	if _nWBrowseResultEnumerateInterfaces == nil {
		return symbolCallError("nw_browse_result_enumerate_interfaces", "10.15", _nWBrowseResultEnumerateInterfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return enumerator(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWBrowseResultEnumerateInterfaces(result, _block0)
	return nil
}

// NWBrowseResultEnumerateInterfaces enumerates the list of interfaces on which the service was discovered.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interfaces(_:_:)
func NWBrowseResultEnumerateInterfaces(result NWBrowseResult, enumerator NWBrowseResultEnumerateInterface) {
	if callErr := tryNWBrowseResultEnumerateInterfaces(result, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowseResultGetChanges func(old_result NWBrowseResult, new_result NWBrowseResult) NWBrowseResultChange
var _nWBrowseResultGetChangesErr error

func tryNWBrowseResultGetChanges(old_result NWBrowseResult, new_result NWBrowseResult) (NWBrowseResultChange, error) {
	if _nWBrowseResultGetChanges == nil {
		return *new(NWBrowseResultChange), symbolCallError("nw_browse_result_get_changes", "10.15", _nWBrowseResultGetChangesErr)
	}
	return _nWBrowseResultGetChanges(old_result, new_result), nil
}

// NWBrowseResultGetChanges compares two discovered services and calculates changes between them.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_changes(_:_:)
func NWBrowseResultGetChanges(old_result NWBrowseResult, new_result NWBrowseResult) NWBrowseResultChange {
	result, callErr := tryNWBrowseResultGetChanges(old_result, new_result)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowseResultGetInterfacesCount func(result NWBrowseResult) uintptr
var _nWBrowseResultGetInterfacesCountErr error

func tryNWBrowseResultGetInterfacesCount(result NWBrowseResult) (uintptr, error) {
	if _nWBrowseResultGetInterfacesCount == nil {
		return 0, symbolCallError("nw_browse_result_get_interfaces_count", "10.15", _nWBrowseResultGetInterfacesCountErr)
	}
	return _nWBrowseResultGetInterfacesCount(result), nil
}

// NWBrowseResultGetInterfacesCount accesses the number of interfaces associated with a discovered service.
//
// See: https://developer.apple.com/documentation/Network/nw_browse_result_get_interfaces_count(_:)
func NWBrowseResultGetInterfacesCount(result NWBrowseResult) uintptr {
	result0, callErr := tryNWBrowseResultGetInterfacesCount(result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nWBrowserCancel func(browser NWBrowser)
var _nWBrowserCancelErr error

func tryNWBrowserCancel(browser NWBrowser) error {
	if _nWBrowserCancel == nil {
		return symbolCallError("nw_browser_cancel", "10.15", _nWBrowserCancelErr)
	}
	_nWBrowserCancel(browser)
	return nil
}

// NWBrowserCancel stops browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_cancel(_:)
func NWBrowserCancel(browser NWBrowser) {
	if callErr := tryNWBrowserCancel(browser); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowserCopyBrowseDescriptor func(browser NWBrowser) NWBrowseDescriptor
var _nWBrowserCopyBrowseDescriptorErr error

func tryNWBrowserCopyBrowseDescriptor(browser NWBrowser) (NWBrowseDescriptor, error) {
	if _nWBrowserCopyBrowseDescriptor == nil {
		return *new(NWBrowseDescriptor), symbolCallError("nw_browser_copy_browse_descriptor", "10.15", _nWBrowserCopyBrowseDescriptorErr)
	}
	return _nWBrowserCopyBrowseDescriptor(browser), nil
}

// NWBrowserCopyBrowseDescriptor accesses the service descriptor with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_browse_descriptor(_:)
func NWBrowserCopyBrowseDescriptor(browser NWBrowser) NWBrowseDescriptor {
	result, callErr := tryNWBrowserCopyBrowseDescriptor(browser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowserCopyParameters func(browser NWBrowser) NWParameters
var _nWBrowserCopyParametersErr error

func tryNWBrowserCopyParameters(browser NWBrowser) (NWParameters, error) {
	if _nWBrowserCopyParameters == nil {
		return *new(NWParameters), symbolCallError("nw_browser_copy_parameters", "10.15", _nWBrowserCopyParametersErr)
	}
	return _nWBrowserCopyParameters(browser), nil
}

// NWBrowserCopyParameters accesses the parameters with which the browser was created.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_copy_parameters(_:)
func NWBrowserCopyParameters(browser NWBrowser) NWParameters {
	result, callErr := tryNWBrowserCopyParameters(browser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowserCreate func(descriptor NWBrowseDescriptor, parameters NWParameters) NWBrowser
var _nWBrowserCreateErr error

func tryNWBrowserCreate(descriptor NWBrowseDescriptor, parameters NWParameters) (NWBrowser, error) {
	if _nWBrowserCreate == nil {
		return *new(NWBrowser), symbolCallError("nw_browser_create", "10.15", _nWBrowserCreateErr)
	}
	return _nWBrowserCreate(descriptor, parameters), nil
}

// NWBrowserCreate initializes a browser with a type of service to discover.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_create(_:_:)
func NWBrowserCreate(descriptor NWBrowseDescriptor, parameters NWParameters) NWBrowser {
	result, callErr := tryNWBrowserCreate(descriptor, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWBrowserSetBrowseResultsChangedHandler func(browser NWBrowser, handler unsafe.Pointer)
var _nWBrowserSetBrowseResultsChangedHandlerErr error

func tryNWBrowserSetBrowseResultsChangedHandler(browser NWBrowser, handler NWBrowserBrowseResultsChangedHandler) error {
	if _nWBrowserSetBrowseResultsChangedHandler == nil {
		return symbolCallError("nw_browser_set_browse_results_changed_handler", "10.15", _nWBrowserSetBrowseResultsChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 objc.ID, blockArg2 bool) {
		handler(objectivec.ObjectFromID(blockArg0), objectivec.ObjectFromID(blockArg1), blockArg2)
	})
	retainNetworkAsyncBlock(browser.ID, "nw_browser_set_browse_results_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWBrowserSetBrowseResultsChangedHandler(browser, _block0)
	return nil
}

// NWBrowserSetBrowseResultsChangedHandler sets the handler to receive updates about discovered services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_browse_results_changed_handler(_:_:)
func NWBrowserSetBrowseResultsChangedHandler(browser NWBrowser, handler NWBrowserBrowseResultsChangedHandler) {
	if callErr := tryNWBrowserSetBrowseResultsChangedHandler(browser, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowserSetQueue func(browser NWBrowser, queue uintptr)
var _nWBrowserSetQueueErr error

func tryNWBrowserSetQueue(browser NWBrowser, queue dispatch.Queue) error {
	if _nWBrowserSetQueue == nil {
		return symbolCallError("nw_browser_set_queue", "10.15", _nWBrowserSetQueueErr)
	}
	_nWBrowserSetQueue(browser, uintptr(queue.Handle()))
	return nil
}

// NWBrowserSetQueue sets the queue on which all browser events will be delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_queue(_:_:)
func NWBrowserSetQueue(browser NWBrowser, queue dispatch.Queue) {
	if callErr := tryNWBrowserSetQueue(browser, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowserSetStateChangedHandler func(browser NWBrowser, state_changed_handler unsafe.Pointer)
var _nWBrowserSetStateChangedHandlerErr error

func tryNWBrowserSetStateChangedHandler(browser NWBrowser, state_changed_handler NWBrowserStateChangedHandler) error {
	if _nWBrowserSetStateChangedHandler == nil {
		return symbolCallError("nw_browser_set_state_changed_handler", "10.15", _nWBrowserSetStateChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWBrowserState, blockArg1 objc.ID) {
		state_changed_handler(blockArg0, NWError{Object: objectivec.ObjectFromID(blockArg1)})
	})
	retainNetworkAsyncBlock(browser.ID, "nw_browser_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWBrowserSetStateChangedHandler(browser, _block0)
	return nil
}

// NWBrowserSetStateChangedHandler sets a handler to receive browser state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_set_state_changed_handler(_:_:)
func NWBrowserSetStateChangedHandler(browser NWBrowser, state_changed_handler NWBrowserStateChangedHandler) {
	if callErr := tryNWBrowserSetStateChangedHandler(browser, state_changed_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWBrowserStart func(browser NWBrowser)
var _nWBrowserStartErr error

func tryNWBrowserStart(browser NWBrowser) error {
	if _nWBrowserStart == nil {
		return symbolCallError("nw_browser_start", "10.15", _nWBrowserStartErr)
	}
	_nWBrowserStart(browser)
	return nil
}

// NWBrowserStart starts browsing for services.
//
// See: https://developer.apple.com/documentation/Network/nw_browser_start(_:)
func NWBrowserStart(browser NWBrowser) {
	if callErr := tryNWBrowserStart(browser); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionAccessEstablishmentReport func(connection NWConnection, queue uintptr, access_block unsafe.Pointer)
var _nWConnectionAccessEstablishmentReportErr error

func tryNWConnectionAccessEstablishmentReport(connection NWConnection, queue dispatch.Queue, access_block NWEstablishmentReportAccessBlock) error {
	if _nWConnectionAccessEstablishmentReport == nil {
		return symbolCallError("nw_connection_access_establishment_report", "10.15", _nWConnectionAccessEstablishmentReportErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { access_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionAccessEstablishmentReport(connection, uintptr(queue.Handle()), _block0)
	return nil
}

// NWConnectionAccessEstablishmentReport requests a copy of the connection’s establishment report once the connection is in the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_access_establishment_report(_:_:_:)
func NWConnectionAccessEstablishmentReport(connection NWConnection, queue dispatch.Queue, access_block NWEstablishmentReportAccessBlock) {
	if callErr := tryNWConnectionAccessEstablishmentReport(connection, queue, access_block); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionBatch func(connection NWConnection, batch_block unsafe.Pointer)
var _nWConnectionBatchErr error

func tryNWConnectionBatch(connection NWConnection, batch_block unsafe.Pointer) error {
	if _nWConnectionBatch == nil {
		return symbolCallError("nw_connection_batch", "10.14", _nWConnectionBatchErr)
	}
	_nWConnectionBatch(connection, batch_block)
	return nil
}

// NWConnectionBatch defines a block in which calls to send and receive are processed as a batch to improve performance.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_batch(_:_:)
func NWConnectionBatch(connection NWConnection, batch_block unsafe.Pointer) {
	if callErr := tryNWConnectionBatch(connection, batch_block); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionCancel func(connection NWConnection)
var _nWConnectionCancelErr error

func tryNWConnectionCancel(connection NWConnection) error {
	if _nWConnectionCancel == nil {
		return symbolCallError("nw_connection_cancel", "10.14", _nWConnectionCancelErr)
	}
	_nWConnectionCancel(connection)
	return nil
}

// NWConnectionCancel cancels the connection and gracefully disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel(_:)
func NWConnectionCancel(connection NWConnection) {
	if callErr := tryNWConnectionCancel(connection); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionCancelCurrentEndpoint func(connection NWConnection)
var _nWConnectionCancelCurrentEndpointErr error

func tryNWConnectionCancelCurrentEndpoint(connection NWConnection) error {
	if _nWConnectionCancelCurrentEndpoint == nil {
		return symbolCallError("nw_connection_cancel_current_endpoint", "10.14", _nWConnectionCancelCurrentEndpointErr)
	}
	_nWConnectionCancelCurrentEndpoint(connection)
	return nil
}

// NWConnectionCancelCurrentEndpoint causes the current endpoint to be rejected, allowing the connection to try another resolved address.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_cancel_current_endpoint(_:)
func NWConnectionCancelCurrentEndpoint(connection NWConnection) {
	if callErr := tryNWConnectionCancelCurrentEndpoint(connection); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionCopyCurrentPath func(connection NWConnection) NWPath
var _nWConnectionCopyCurrentPathErr error

func tryNWConnectionCopyCurrentPath(connection NWConnection) (NWPath, error) {
	if _nWConnectionCopyCurrentPath == nil {
		return NWPath{}, symbolCallError("nw_connection_copy_current_path", "10.14", _nWConnectionCopyCurrentPathErr)
	}
	return _nWConnectionCopyCurrentPath(connection), nil
}

// NWConnectionCopyCurrentPath accesses the network path the connection is using.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_current_path(_:)
func NWConnectionCopyCurrentPath(connection NWConnection) NWPath {
	result, callErr := tryNWConnectionCopyCurrentPath(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCopyDescription func(connection NWConnection) *byte
var _nWConnectionCopyDescriptionErr error

func tryNWConnectionCopyDescription(connection NWConnection) (*byte, error) {
	if _nWConnectionCopyDescription == nil {
		return nil, symbolCallError("nw_connection_copy_description", "10.14", _nWConnectionCopyDescriptionErr)
	}
	return _nWConnectionCopyDescription(connection), nil
}

// NWConnectionCopyDescription copies the description of the connection as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_description(_:)
func NWConnectionCopyDescription(connection NWConnection) *byte {
	result, callErr := tryNWConnectionCopyDescription(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCopyEndpoint func(connection NWConnection) NWEndpoint
var _nWConnectionCopyEndpointErr error

func tryNWConnectionCopyEndpoint(connection NWConnection) (NWEndpoint, error) {
	if _nWConnectionCopyEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_connection_copy_endpoint", "10.14", _nWConnectionCopyEndpointErr)
	}
	return _nWConnectionCopyEndpoint(connection), nil
}

// NWConnectionCopyEndpoint accesses the endpoint with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_endpoint(_:)
func NWConnectionCopyEndpoint(connection NWConnection) NWEndpoint {
	result, callErr := tryNWConnectionCopyEndpoint(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCopyParameters func(connection NWConnection) NWParameters
var _nWConnectionCopyParametersErr error

func tryNWConnectionCopyParameters(connection NWConnection) (NWParameters, error) {
	if _nWConnectionCopyParameters == nil {
		return *new(NWParameters), symbolCallError("nw_connection_copy_parameters", "10.14", _nWConnectionCopyParametersErr)
	}
	return _nWConnectionCopyParameters(connection), nil
}

// NWConnectionCopyParameters accesses the parameters with which the connection was created.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_parameters(_:)
func NWConnectionCopyParameters(connection NWConnection) NWParameters {
	result, callErr := tryNWConnectionCopyParameters(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCopyProtocolMetadata func(connection NWConnection, definition NWProtocolDefinition) NWProtocolMetadata
var _nWConnectionCopyProtocolMetadataErr error

func tryNWConnectionCopyProtocolMetadata(connection NWConnection, definition NWProtocolDefinition) (NWProtocolMetadata, error) {
	if _nWConnectionCopyProtocolMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_connection_copy_protocol_metadata", "10.14", _nWConnectionCopyProtocolMetadataErr)
	}
	return _nWConnectionCopyProtocolMetadata(connection, definition), nil
}

// NWConnectionCopyProtocolMetadata retrieves the connection-wide metadata for a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_copy_protocol_metadata(_:_:)
func NWConnectionCopyProtocolMetadata(connection NWConnection, definition NWProtocolDefinition) NWProtocolMetadata {
	result, callErr := tryNWConnectionCopyProtocolMetadata(connection, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCreate func(endpoint NWEndpoint, parameters NWParameters) NWConnection
var _nWConnectionCreateErr error

func tryNWConnectionCreate(endpoint NWEndpoint, parameters NWParameters) (NWConnection, error) {
	if _nWConnectionCreate == nil {
		return *new(NWConnection), symbolCallError("nw_connection_create", "10.14", _nWConnectionCreateErr)
	}
	return _nWConnectionCreate(endpoint, parameters), nil
}

// NWConnectionCreate initializes a new connection to a remote endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create(_:_:)
func NWConnectionCreate(endpoint NWEndpoint, parameters NWParameters) NWConnection {
	result, callErr := tryNWConnectionCreate(endpoint, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionCreateNewDataTransferReport func(connection NWConnection) NWDataTransferReport
var _nWConnectionCreateNewDataTransferReportErr error

func tryNWConnectionCreateNewDataTransferReport(connection NWConnection) (NWDataTransferReport, error) {
	if _nWConnectionCreateNewDataTransferReport == nil {
		return *new(NWDataTransferReport), symbolCallError("nw_connection_create_new_data_transfer_report", "10.15", _nWConnectionCreateNewDataTransferReportErr)
	}
	return _nWConnectionCreateNewDataTransferReport(connection), nil
}

// NWConnectionCreateNewDataTransferReport begins a new data transfer report, which can later be collected.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_create_new_data_transfer_report(_:)
func NWConnectionCreateNewDataTransferReport(connection NWConnection) NWDataTransferReport {
	result, callErr := tryNWConnectionCreateNewDataTransferReport(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionForceCancel func(connection NWConnection)
var _nWConnectionForceCancelErr error

func tryNWConnectionForceCancel(connection NWConnection) error {
	if _nWConnectionForceCancel == nil {
		return symbolCallError("nw_connection_force_cancel", "10.14", _nWConnectionForceCancelErr)
	}
	_nWConnectionForceCancel(connection)
	return nil
}

// NWConnectionForceCancel cancels the connection and immediately disconnects any established network protocols.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_force_cancel(_:)
func NWConnectionForceCancel(connection NWConnection) {
	if callErr := tryNWConnectionForceCancel(connection); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGetMaximumDatagramSize func(connection NWConnection) uint32
var _nWConnectionGetMaximumDatagramSizeErr error

func tryNWConnectionGetMaximumDatagramSize(connection NWConnection) (uint32, error) {
	if _nWConnectionGetMaximumDatagramSize == nil {
		return 0, symbolCallError("nw_connection_get_maximum_datagram_size", "10.14", _nWConnectionGetMaximumDatagramSizeErr)
	}
	return _nWConnectionGetMaximumDatagramSize(connection), nil
}

// NWConnectionGetMaximumDatagramSize accesses the maximum size of a datagram message that can be sent on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_get_maximum_datagram_size(_:)
func NWConnectionGetMaximumDatagramSize(connection NWConnection) uint32 {
	result, callErr := tryNWConnectionGetMaximumDatagramSize(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCancel func(group NWConnectionGroup)
var _nWConnectionGroupCancelErr error

func tryNWConnectionGroupCancel(group NWConnectionGroup) error {
	if _nWConnectionGroupCancel == nil {
		return symbolCallError("nw_connection_group_cancel", "11.0", _nWConnectionGroupCancelErr)
	}
	_nWConnectionGroupCancel(group)
	return nil
}

// NWConnectionGroupCancel cancels the connection group object and leaves the network group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_cancel(_:)
func NWConnectionGroupCancel(group NWConnectionGroup) {
	if callErr := tryNWConnectionGroupCancel(group); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupCopyDescriptor func(group NWConnectionGroup) NWGroupDescriptor
var _nWConnectionGroupCopyDescriptorErr error

func tryNWConnectionGroupCopyDescriptor(group NWConnectionGroup) (NWGroupDescriptor, error) {
	if _nWConnectionGroupCopyDescriptor == nil {
		return *new(NWGroupDescriptor), symbolCallError("nw_connection_group_copy_descriptor", "11.0", _nWConnectionGroupCopyDescriptorErr)
	}
	return _nWConnectionGroupCopyDescriptor(group), nil
}

// NWConnectionGroupCopyDescriptor accesses the descriptor of the group you use to initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_descriptor(_:)
func NWConnectionGroupCopyDescriptor(group NWConnectionGroup) NWGroupDescriptor {
	result, callErr := tryNWConnectionGroupCopyDescriptor(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyLocalEndpointForMessage func(group NWConnectionGroup, context NWContentContext) NWEndpoint
var _nWConnectionGroupCopyLocalEndpointForMessageErr error

func tryNWConnectionGroupCopyLocalEndpointForMessage(group NWConnectionGroup, context NWContentContext) (NWEndpoint, error) {
	if _nWConnectionGroupCopyLocalEndpointForMessage == nil {
		return NWEndpoint{}, symbolCallError("nw_connection_group_copy_local_endpoint_for_message", "11.0", _nWConnectionGroupCopyLocalEndpointForMessageErr)
	}
	return _nWConnectionGroupCopyLocalEndpointForMessage(group, context), nil
}

// NWConnectionGroupCopyLocalEndpointForMessage accesses the local address and port you use to receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_local_endpoint_for_message(_:_:)
func NWConnectionGroupCopyLocalEndpointForMessage(group NWConnectionGroup, context NWContentContext) NWEndpoint {
	result, callErr := tryNWConnectionGroupCopyLocalEndpointForMessage(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyParameters func(group NWConnectionGroup) NWParameters
var _nWConnectionGroupCopyParametersErr error

func tryNWConnectionGroupCopyParameters(group NWConnectionGroup) (NWParameters, error) {
	if _nWConnectionGroupCopyParameters == nil {
		return *new(NWParameters), symbolCallError("nw_connection_group_copy_parameters", "11.0", _nWConnectionGroupCopyParametersErr)
	}
	return _nWConnectionGroupCopyParameters(group), nil
}

// NWConnectionGroupCopyParameters accesses the parameters with which you initialize the connection group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_parameters(_:)
func NWConnectionGroupCopyParameters(group NWConnectionGroup) NWParameters {
	result, callErr := tryNWConnectionGroupCopyParameters(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyPathForMessage func(group NWConnectionGroup, context NWContentContext) NWPath
var _nWConnectionGroupCopyPathForMessageErr error

func tryNWConnectionGroupCopyPathForMessage(group NWConnectionGroup, context NWContentContext) (NWPath, error) {
	if _nWConnectionGroupCopyPathForMessage == nil {
		return NWPath{}, symbolCallError("nw_connection_group_copy_path_for_message", "11.0", _nWConnectionGroupCopyPathForMessageErr)
	}
	return _nWConnectionGroupCopyPathForMessage(group, context), nil
}

// NWConnectionGroupCopyPathForMessage accesses the network path on which you receive the message.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_path_for_message(_:_:)
func NWConnectionGroupCopyPathForMessage(group NWConnectionGroup, context NWContentContext) NWPath {
	result, callErr := tryNWConnectionGroupCopyPathForMessage(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyProtocolMetadata func(group NWConnectionGroup, definition NWProtocolDefinition) NWProtocolMetadata
var _nWConnectionGroupCopyProtocolMetadataErr error

func tryNWConnectionGroupCopyProtocolMetadata(group NWConnectionGroup, definition NWProtocolDefinition) (NWProtocolMetadata, error) {
	if _nWConnectionGroupCopyProtocolMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_connection_group_copy_protocol_metadata", "12.0", _nWConnectionGroupCopyProtocolMetadataErr)
	}
	return _nWConnectionGroupCopyProtocolMetadata(group, definition), nil
}

// NWConnectionGroupCopyProtocolMetadata.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata(_:_:)
func NWConnectionGroupCopyProtocolMetadata(group NWConnectionGroup, definition NWProtocolDefinition) NWProtocolMetadata {
	result, callErr := tryNWConnectionGroupCopyProtocolMetadata(group, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyProtocolMetadataForMessage func(group NWConnectionGroup, context NWContentContext, definition NWProtocolDefinition) NWProtocolMetadata
var _nWConnectionGroupCopyProtocolMetadataForMessageErr error

func tryNWConnectionGroupCopyProtocolMetadataForMessage(group NWConnectionGroup, context NWContentContext, definition NWProtocolDefinition) (NWProtocolMetadata, error) {
	if _nWConnectionGroupCopyProtocolMetadataForMessage == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_connection_group_copy_protocol_metadata_for_message", "12.0", _nWConnectionGroupCopyProtocolMetadataForMessageErr)
	}
	return _nWConnectionGroupCopyProtocolMetadataForMessage(group, context, definition), nil
}

// NWConnectionGroupCopyProtocolMetadataForMessage.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata_for_message(_:_:_:)
func NWConnectionGroupCopyProtocolMetadataForMessage(group NWConnectionGroup, context NWContentContext, definition NWProtocolDefinition) NWProtocolMetadata {
	result, callErr := tryNWConnectionGroupCopyProtocolMetadataForMessage(group, context, definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCopyRemoteEndpointForMessage func(group NWConnectionGroup, context NWContentContext) NWEndpoint
var _nWConnectionGroupCopyRemoteEndpointForMessageErr error

func tryNWConnectionGroupCopyRemoteEndpointForMessage(group NWConnectionGroup, context NWContentContext) (NWEndpoint, error) {
	if _nWConnectionGroupCopyRemoteEndpointForMessage == nil {
		return NWEndpoint{}, symbolCallError("nw_connection_group_copy_remote_endpoint_for_message", "11.0", _nWConnectionGroupCopyRemoteEndpointForMessageErr)
	}
	return _nWConnectionGroupCopyRemoteEndpointForMessage(group, context), nil
}

// NWConnectionGroupCopyRemoteEndpointForMessage accesses the endpoint that originates the message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_copy_remote_endpoint_for_message(_:_:)
func NWConnectionGroupCopyRemoteEndpointForMessage(group NWConnectionGroup, context NWContentContext) NWEndpoint {
	result, callErr := tryNWConnectionGroupCopyRemoteEndpointForMessage(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupCreate func(group_descriptor NWGroupDescriptor, parameters NWParameters) NWConnectionGroup
var _nWConnectionGroupCreateErr error

func tryNWConnectionGroupCreate(group_descriptor NWGroupDescriptor, parameters NWParameters) (NWConnectionGroup, error) {
	if _nWConnectionGroupCreate == nil {
		return *new(NWConnectionGroup), symbolCallError("nw_connection_group_create", "11.0", _nWConnectionGroupCreateErr)
	}
	return _nWConnectionGroupCreate(group_descriptor, parameters), nil
}

// NWConnectionGroupCreate initializes a new connection group with a group identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_create(_:_:)
func NWConnectionGroupCreate(group_descriptor NWGroupDescriptor, parameters NWParameters) NWConnectionGroup {
	result, callErr := tryNWConnectionGroupCreate(group_descriptor, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupExtractConnection func(group NWConnectionGroup, endpoint NWEndpoint, protocol_options NWProtocolOptions) NWConnection
var _nWConnectionGroupExtractConnectionErr error

func tryNWConnectionGroupExtractConnection(group NWConnectionGroup, endpoint NWEndpoint, protocol_options NWProtocolOptions) (NWConnection, error) {
	if _nWConnectionGroupExtractConnection == nil {
		return *new(NWConnection), symbolCallError("nw_connection_group_extract_connection", "12.0", _nWConnectionGroupExtractConnectionErr)
	}
	return _nWConnectionGroupExtractConnection(group, endpoint, protocol_options), nil
}

// NWConnectionGroupExtractConnection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection(_:_:_:)
func NWConnectionGroupExtractConnection(group NWConnectionGroup, endpoint NWEndpoint, protocol_options NWProtocolOptions) NWConnection {
	result, callErr := tryNWConnectionGroupExtractConnection(group, endpoint, protocol_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupExtractConnectionForMessage func(group NWConnectionGroup, context NWContentContext) NWConnection
var _nWConnectionGroupExtractConnectionForMessageErr error

func tryNWConnectionGroupExtractConnectionForMessage(group NWConnectionGroup, context NWContentContext) (NWConnection, error) {
	if _nWConnectionGroupExtractConnectionForMessage == nil {
		return *new(NWConnection), symbolCallError("nw_connection_group_extract_connection_for_message", "11.0", _nWConnectionGroupExtractConnectionForMessageErr)
	}
	return _nWConnectionGroupExtractConnectionForMessage(group, context), nil
}

// NWConnectionGroupExtractConnectionForMessage converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection_for_message(_:_:)
func NWConnectionGroupExtractConnectionForMessage(group NWConnectionGroup, context NWContentContext) NWConnection {
	result, callErr := tryNWConnectionGroupExtractConnectionForMessage(group, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupReinsertExtractedConnection func(group NWConnectionGroup, connection NWConnection) bool
var _nWConnectionGroupReinsertExtractedConnectionErr error

func tryNWConnectionGroupReinsertExtractedConnection(group NWConnectionGroup, connection NWConnection) (bool, error) {
	if _nWConnectionGroupReinsertExtractedConnection == nil {
		return false, symbolCallError("nw_connection_group_reinsert_extracted_connection", "12.0", _nWConnectionGroupReinsertExtractedConnectionErr)
	}
	return _nWConnectionGroupReinsertExtractedConnection(group, connection), nil
}

// NWConnectionGroupReinsertExtractedConnection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reinsert_extracted_connection(_:_:)
func NWConnectionGroupReinsertExtractedConnection(group NWConnectionGroup, connection NWConnection) bool {
	result, callErr := tryNWConnectionGroupReinsertExtractedConnection(group, connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWConnectionGroupReply func(group NWConnectionGroup, inbound_message NWContentContext, outbound_message NWContentContext, content uintptr)
var _nWConnectionGroupReplyErr error

func tryNWConnectionGroupReply(group NWConnectionGroup, inbound_message NWContentContext, outbound_message NWContentContext, content dispatch.Data) error {
	if _nWConnectionGroupReply == nil {
		return symbolCallError("nw_connection_group_reply", "11.0", _nWConnectionGroupReplyErr)
	}
	_nWConnectionGroupReply(group, inbound_message, outbound_message, uintptr(content.Handle()))
	return nil
}

// NWConnectionGroupReply sends a reply to the specific endpoint that originates a group message you receive.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_reply(_:_:_:_:)
func NWConnectionGroupReply(group NWConnectionGroup, inbound_message NWContentContext, outbound_message NWContentContext, content dispatch.Data) {
	if callErr := tryNWConnectionGroupReply(group, inbound_message, outbound_message, content); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupSendMessage func(group NWConnectionGroup, content uintptr, endpoint NWEndpoint, context NWContentContext, completion unsafe.Pointer)
var _nWConnectionGroupSendMessageErr error

func tryNWConnectionGroupSendMessage(group NWConnectionGroup, content dispatch.Data, endpoint NWEndpoint, context NWContentContext, completion NWConnectionGroupSendCompletion) error {
	if _nWConnectionGroupSendMessage == nil {
		return symbolCallError("nw_connection_group_send_message", "11.0", _nWConnectionGroupSendMessageErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { completion(NWError{Object: objectivec.ObjectFromID(blockArg0)}) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionGroupSendMessage(group, uintptr(content.Handle()), endpoint, context, _block0)
	return nil
}

// NWConnectionGroupSendMessage sends data to the entire group, or to a specific member of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_send_message(_:_:_:_:_:)
func NWConnectionGroupSendMessage(group NWConnectionGroup, content dispatch.Data, endpoint NWEndpoint, context NWContentContext, completion NWConnectionGroupSendCompletion) {
	if callErr := tryNWConnectionGroupSendMessage(group, content, endpoint, context, completion); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupSetNewConnectionHandler func(group NWConnectionGroup, new_connection_handler unsafe.Pointer)
var _nWConnectionGroupSetNewConnectionHandlerErr error

func tryNWConnectionGroupSetNewConnectionHandler(group NWConnectionGroup, new_connection_handler NWConnectionGroupNewConnectionHandler) error {
	if _nWConnectionGroupSetNewConnectionHandler == nil {
		return symbolCallError("nw_connection_group_set_new_connection_handler", "12.0", _nWConnectionGroupSetNewConnectionHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { new_connection_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_new_connection_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionGroupSetNewConnectionHandler(group, _block0)
	return nil
}

// NWConnectionGroupSetNewConnectionHandler.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_new_connection_handler(_:_:)
func NWConnectionGroupSetNewConnectionHandler(group NWConnectionGroup, new_connection_handler NWConnectionGroupNewConnectionHandler) {
	if callErr := tryNWConnectionGroupSetNewConnectionHandler(group, new_connection_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupSetQueue func(group NWConnectionGroup, queue uintptr)
var _nWConnectionGroupSetQueueErr error

func tryNWConnectionGroupSetQueue(group NWConnectionGroup, queue dispatch.Queue) error {
	if _nWConnectionGroupSetQueue == nil {
		return symbolCallError("nw_connection_group_set_queue", "11.0", _nWConnectionGroupSetQueueErr)
	}
	_nWConnectionGroupSetQueue(group, uintptr(queue.Handle()))
	return nil
}

// NWConnectionGroupSetQueue sets the queue on which you handle connection group events.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_queue(_:_:)
func NWConnectionGroupSetQueue(group NWConnectionGroup, queue dispatch.Queue) {
	if callErr := tryNWConnectionGroupSetQueue(group, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupSetReceiveHandler func(group NWConnectionGroup, maximum_message_size uint32, reject_oversized_messages bool, receive_handler unsafe.Pointer)
var _nWConnectionGroupSetReceiveHandlerErr error

func tryNWConnectionGroupSetReceiveHandler(group NWConnectionGroup, maximum_message_size uint32, reject_oversized_messages bool, receive_handler NWConnectionGroupReceiveHandler) error {
	if _nWConnectionGroupSetReceiveHandler == nil {
		return symbolCallError("nw_connection_group_set_receive_handler", "11.0", _nWConnectionGroupSetReceiveHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 objc.ID, blockArg2 bool) {
		receive_handler(objectivec.ObjectFromID(blockArg0), objectivec.ObjectFromID(blockArg1), blockArg2)
	})
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_receive_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionGroupSetReceiveHandler(group, maximum_message_size, reject_oversized_messages, _block0)
	return nil
}

// NWConnectionGroupSetReceiveHandler sets a handler that receives inbound messages from members of the group.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_receive_handler(_:_:_:_:)
func NWConnectionGroupSetReceiveHandler(group NWConnectionGroup, maximum_message_size uint32, reject_oversized_messages bool, receive_handler NWConnectionGroupReceiveHandler) {
	if callErr := tryNWConnectionGroupSetReceiveHandler(group, maximum_message_size, reject_oversized_messages, receive_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupSetStateChangedHandler func(group NWConnectionGroup, state_changed_handler unsafe.Pointer)
var _nWConnectionGroupSetStateChangedHandlerErr error

func tryNWConnectionGroupSetStateChangedHandler(group NWConnectionGroup, state_changed_handler NWConnectionGroupStateChangedHandler) error {
	if _nWConnectionGroupSetStateChangedHandler == nil {
		return symbolCallError("nw_connection_group_set_state_changed_handler", "11.0", _nWConnectionGroupSetStateChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWConnectionGroupState, blockArg1 objc.ID) {
		state_changed_handler(blockArg0, NWError{Object: objectivec.ObjectFromID(blockArg1)})
	})
	retainNetworkAsyncBlock(group.ID, "nw_connection_group_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionGroupSetStateChangedHandler(group, _block0)
	return nil
}

// NWConnectionGroupSetStateChangedHandler sets a handler that receives connection group state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_set_state_changed_handler(_:_:)
func NWConnectionGroupSetStateChangedHandler(group NWConnectionGroup, state_changed_handler NWConnectionGroupStateChangedHandler) {
	if callErr := tryNWConnectionGroupSetStateChangedHandler(group, state_changed_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionGroupStart func(group NWConnectionGroup)
var _nWConnectionGroupStartErr error

func tryNWConnectionGroupStart(group NWConnectionGroup) error {
	if _nWConnectionGroupStart == nil {
		return symbolCallError("nw_connection_group_start", "11.0", _nWConnectionGroupStartErr)
	}
	_nWConnectionGroupStart(group)
	return nil
}

// NWConnectionGroupStart joins the group and registers to receive messages.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_group_start(_:)
func NWConnectionGroupStart(group NWConnectionGroup) {
	if callErr := tryNWConnectionGroupStart(group); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionReceive func(connection NWConnection, minimum_incomplete_length uint32, maximum_length uint32, completion unsafe.Pointer)
var _nWConnectionReceiveErr error

func tryNWConnectionReceive(connection NWConnection, minimum_incomplete_length uint32, maximum_length uint32, completion NWConnectionReceiveCompletion) error {
	if _nWConnectionReceive == nil {
		return symbolCallError("nw_connection_receive", "10.14", _nWConnectionReceiveErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 objc.ID, blockArg2 bool, blockArg3 objc.ID) {
		completion(objectivec.ObjectFromID(blockArg0), objectivec.ObjectFromID(blockArg1), blockArg2, NWError{Object: objectivec.ObjectFromID(blockArg3)})
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionReceive(connection, minimum_incomplete_length, maximum_length, _block0)
	return nil
}

// NWConnectionReceive schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive(_:_:_:_:)
func NWConnectionReceive(connection NWConnection, minimum_incomplete_length uint32, maximum_length uint32, completion NWConnectionReceiveCompletion) {
	if callErr := tryNWConnectionReceive(connection, minimum_incomplete_length, maximum_length, completion); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionReceiveMessage func(connection NWConnection, completion unsafe.Pointer)
var _nWConnectionReceiveMessageErr error

func tryNWConnectionReceiveMessage(connection NWConnection, completion NWConnectionReceiveCompletion) error {
	if _nWConnectionReceiveMessage == nil {
		return symbolCallError("nw_connection_receive_message", "10.14", _nWConnectionReceiveMessageErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 objc.ID, blockArg2 bool, blockArg3 objc.ID) {
		completion(objectivec.ObjectFromID(blockArg0), objectivec.ObjectFromID(blockArg1), blockArg2, NWError{Object: objectivec.ObjectFromID(blockArg3)})
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionReceiveMessage(connection, _block0)
	return nil
}

// NWConnectionReceiveMessage schedules a single receive completion handler for a complete message, as opposed to a range of bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_receive_message(_:_:)
func NWConnectionReceiveMessage(connection NWConnection, completion NWConnectionReceiveCompletion) {
	if callErr := tryNWConnectionReceiveMessage(connection, completion); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionRestart func(connection NWConnection)
var _nWConnectionRestartErr error

func tryNWConnectionRestart(connection NWConnection) error {
	if _nWConnectionRestart == nil {
		return symbolCallError("nw_connection_restart", "10.14", _nWConnectionRestartErr)
	}
	_nWConnectionRestart(connection)
	return nil
}

// NWConnectionRestart restarts a connection that is in the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_restart(_:)
func NWConnectionRestart(connection NWConnection) {
	if callErr := tryNWConnectionRestart(connection); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSend func(connection NWConnection, content uintptr, context NWContentContext, is_complete bool, completion unsafe.Pointer)
var _nWConnectionSendErr error

func tryNWConnectionSend(connection NWConnection, content dispatch.Data, context NWContentContext, is_complete bool, completion NWConnectionSendCompletion) error {
	if _nWConnectionSend == nil {
		return symbolCallError("nw_connection_send", "10.14", _nWConnectionSendErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { completion(NWError{Object: objectivec.ObjectFromID(blockArg0)}) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionSend(connection, uintptr(content.Handle()), context, is_complete, _block0)
	return nil
}

// NWConnectionSend sends data on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_send(_:_:_:_:_:)
func NWConnectionSend(connection NWConnection, content dispatch.Data, context NWContentContext, is_complete bool, completion NWConnectionSendCompletion) {
	if callErr := tryNWConnectionSend(connection, content, context, is_complete, completion); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSetBetterPathAvailableHandler func(connection NWConnection, handler unsafe.Pointer)
var _nWConnectionSetBetterPathAvailableHandlerErr error

func tryNWConnectionSetBetterPathAvailableHandler(connection NWConnection, handler NWConnectionBooleanEventHandler) error {
	if _nWConnectionSetBetterPathAvailableHandler == nil {
		return symbolCallError("nw_connection_set_better_path_available_handler", "10.14", _nWConnectionSetBetterPathAvailableHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 bool) { handler(blockArg0) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_better_path_available_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionSetBetterPathAvailableHandler(connection, _block0)
	return nil
}

// NWConnectionSetBetterPathAvailableHandler sets a handler that receives updates when an alternative network path is preferred over the current path.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_better_path_available_handler(_:_:)
func NWConnectionSetBetterPathAvailableHandler(connection NWConnection, handler NWConnectionBooleanEventHandler) {
	if callErr := tryNWConnectionSetBetterPathAvailableHandler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSetPathChangedHandler func(connection NWConnection, handler unsafe.Pointer)
var _nWConnectionSetPathChangedHandlerErr error

func tryNWConnectionSetPathChangedHandler(connection NWConnection, handler NWConnectionPathEventHandler) error {
	if _nWConnectionSetPathChangedHandler == nil {
		return symbolCallError("nw_connection_set_path_changed_handler", "10.14", _nWConnectionSetPathChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_path_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionSetPathChangedHandler(connection, _block0)
	return nil
}

// NWConnectionSetPathChangedHandler sets a handler that receives network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_path_changed_handler(_:_:)
func NWConnectionSetPathChangedHandler(connection NWConnection, handler NWConnectionPathEventHandler) {
	if callErr := tryNWConnectionSetPathChangedHandler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSetQueue func(connection NWConnection, queue uintptr)
var _nWConnectionSetQueueErr error

func tryNWConnectionSetQueue(connection NWConnection, queue dispatch.Queue) error {
	if _nWConnectionSetQueue == nil {
		return symbolCallError("nw_connection_set_queue", "10.14", _nWConnectionSetQueueErr)
	}
	_nWConnectionSetQueue(connection, uintptr(queue.Handle()))
	return nil
}

// NWConnectionSetQueue sets the queue on which all connection events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_queue(_:_:)
func NWConnectionSetQueue(connection NWConnection, queue dispatch.Queue) {
	if callErr := tryNWConnectionSetQueue(connection, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSetStateChangedHandler func(connection NWConnection, handler unsafe.Pointer)
var _nWConnectionSetStateChangedHandlerErr error

func tryNWConnectionSetStateChangedHandler(connection NWConnection, handler NWConnectionStateChangedHandler) error {
	if _nWConnectionSetStateChangedHandler == nil {
		return symbolCallError("nw_connection_set_state_changed_handler", "10.14", _nWConnectionSetStateChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWConnectionState, blockArg1 objc.ID) {
		handler(blockArg0, NWError{Object: objectivec.ObjectFromID(blockArg1)})
	})
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionSetStateChangedHandler(connection, _block0)
	return nil
}

// NWConnectionSetStateChangedHandler sets a handler to receive connection state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_state_changed_handler(_:_:)
func NWConnectionSetStateChangedHandler(connection NWConnection, handler NWConnectionStateChangedHandler) {
	if callErr := tryNWConnectionSetStateChangedHandler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionSetViabilityChangedHandler func(connection NWConnection, handler unsafe.Pointer)
var _nWConnectionSetViabilityChangedHandlerErr error

func tryNWConnectionSetViabilityChangedHandler(connection NWConnection, handler NWConnectionBooleanEventHandler) error {
	if _nWConnectionSetViabilityChangedHandler == nil {
		return symbolCallError("nw_connection_set_viability_changed_handler", "10.14", _nWConnectionSetViabilityChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 bool) { handler(blockArg0) })
	retainNetworkAsyncBlock(connection.ID, "nw_connection_set_viability_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWConnectionSetViabilityChangedHandler(connection, _block0)
	return nil
}

// NWConnectionSetViabilityChangedHandler sets a handler that receives updates when data can be sent and received.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_set_viability_changed_handler(_:_:)
func NWConnectionSetViabilityChangedHandler(connection NWConnection, handler NWConnectionBooleanEventHandler) {
	if callErr := tryNWConnectionSetViabilityChangedHandler(connection, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWConnectionStart func(connection NWConnection)
var _nWConnectionStartErr error

func tryNWConnectionStart(connection NWConnection) error {
	if _nWConnectionStart == nil {
		return symbolCallError("nw_connection_start", "10.14", _nWConnectionStartErr)
	}
	_nWConnectionStart(connection)
	return nil
}

// NWConnectionStart starts establishing a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_connection_start(_:)
func NWConnectionStart(connection NWConnection) {
	if callErr := tryNWConnectionStart(connection); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextCopyAntecedent func(context NWContentContext) NWContentContext
var _nWContentContextCopyAntecedentErr error

func tryNWContentContextCopyAntecedent(context NWContentContext) (NWContentContext, error) {
	if _nWContentContextCopyAntecedent == nil {
		return *new(NWContentContext), symbolCallError("nw_content_context_copy_antecedent", "10.14", _nWContentContextCopyAntecedentErr)
	}
	return _nWContentContextCopyAntecedent(context), nil
}

// NWContentContextCopyAntecedent accesses the optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_antecedent(_:)
func NWContentContextCopyAntecedent(context NWContentContext) NWContentContext {
	result, callErr := tryNWContentContextCopyAntecedent(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextCopyProtocolMetadata func(context NWContentContext, protocol_ NWProtocolDefinition) NWProtocolMetadata
var _nWContentContextCopyProtocolMetadataErr error

func tryNWContentContextCopyProtocolMetadata(context NWContentContext, protocol_ NWProtocolDefinition) (NWProtocolMetadata, error) {
	if _nWContentContextCopyProtocolMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_content_context_copy_protocol_metadata", "10.14", _nWContentContextCopyProtocolMetadataErr)
	}
	return _nWContentContextCopyProtocolMetadata(context, protocol_), nil
}

// NWContentContextCopyProtocolMetadata retreives the metadata associated with a specific protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_copy_protocol_metadata(_:_:)
func NWContentContextCopyProtocolMetadata(context NWContentContext, protocol_ NWProtocolDefinition) NWProtocolMetadata {
	result, callErr := tryNWContentContextCopyProtocolMetadata(context, protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextCreate func(context_identifier string) NWContentContext
var _nWContentContextCreateErr error

func tryNWContentContextCreate(context_identifier string) (NWContentContext, error) {
	if _nWContentContextCreate == nil {
		return *new(NWContentContext), symbolCallError("nw_content_context_create", "10.14", _nWContentContextCreateErr)
	}
	return _nWContentContextCreate(context_identifier), nil
}

// NWContentContextCreate initializes a custom message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_create(_:)
func NWContentContextCreate(context_identifier string) NWContentContext {
	result, callErr := tryNWContentContextCreate(context_identifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextForeachProtocolMetadata func(context NWContentContext)
var _nWContentContextForeachProtocolMetadataErr error

func tryNWContentContextForeachProtocolMetadata(context NWContentContext) error {
	if _nWContentContextForeachProtocolMetadata == nil {
		return symbolCallError("nw_content_context_foreach_protocol_metadata", "10.14", _nWContentContextForeachProtocolMetadataErr)
	}
	_nWContentContextForeachProtocolMetadata(context)
	return nil
}

// NWContentContextForeachProtocolMetadata iterates through all protocol metadata associated with the message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_foreach_protocol_metadata(_:_:)
func NWContentContextForeachProtocolMetadata(context NWContentContext) {
	if callErr := tryNWContentContextForeachProtocolMetadata(context); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextGetExpirationMilliseconds func(context NWContentContext) uint64
var _nWContentContextGetExpirationMillisecondsErr error

func tryNWContentContextGetExpirationMilliseconds(context NWContentContext) (uint64, error) {
	if _nWContentContextGetExpirationMilliseconds == nil {
		return 0, symbolCallError("nw_content_context_get_expiration_milliseconds", "10.14", _nWContentContextGetExpirationMillisecondsErr)
	}
	return _nWContentContextGetExpirationMilliseconds(context), nil
}

// NWContentContextGetExpirationMilliseconds accesses the expiration set for this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_expiration_milliseconds(_:)
func NWContentContextGetExpirationMilliseconds(context NWContentContext) uint64 {
	result, callErr := tryNWContentContextGetExpirationMilliseconds(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextGetIdentifier func(context NWContentContext) *byte
var _nWContentContextGetIdentifierErr error

func tryNWContentContextGetIdentifier(context NWContentContext) (*byte, error) {
	if _nWContentContextGetIdentifier == nil {
		return nil, symbolCallError("nw_content_context_get_identifier", "10.14", _nWContentContextGetIdentifierErr)
	}
	return _nWContentContextGetIdentifier(context), nil
}

// NWContentContextGetIdentifier accesses the identifier used to create this message context.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_identifier(_:)
func NWContentContextGetIdentifier(context NWContentContext) *byte {
	result, callErr := tryNWContentContextGetIdentifier(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextGetIsFinal func(context NWContentContext) bool
var _nWContentContextGetIsFinalErr error

func tryNWContentContextGetIsFinal(context NWContentContext) (bool, error) {
	if _nWContentContextGetIsFinal == nil {
		return false, symbolCallError("nw_content_context_get_is_final", "10.14", _nWContentContextGetIsFinalErr)
	}
	return _nWContentContextGetIsFinal(context), nil
}

// NWContentContextGetIsFinal checks whether this context represents the final message being received.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_is_final(_:)
func NWContentContextGetIsFinal(context NWContentContext) bool {
	result, callErr := tryNWContentContextGetIsFinal(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextGetRelativePriority func(context NWContentContext) float64
var _nWContentContextGetRelativePriorityErr error

func tryNWContentContextGetRelativePriority(context NWContentContext) (float64, error) {
	if _nWContentContextGetRelativePriority == nil {
		return 0.0, symbolCallError("nw_content_context_get_relative_priority", "10.14", _nWContentContextGetRelativePriorityErr)
	}
	return _nWContentContextGetRelativePriority(context), nil
}

// NWContentContextGetRelativePriority accesses the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_get_relative_priority(_:)
func NWContentContextGetRelativePriority(context NWContentContext) float64 {
	result, callErr := tryNWContentContextGetRelativePriority(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWContentContextSetAntecedent func(context NWContentContext, antecedent_context NWContentContext)
var _nWContentContextSetAntecedentErr error

func tryNWContentContextSetAntecedent(context NWContentContext, antecedent_context NWContentContext) error {
	if _nWContentContextSetAntecedent == nil {
		return symbolCallError("nw_content_context_set_antecedent", "10.14", _nWContentContextSetAntecedentErr)
	}
	_nWContentContextSetAntecedent(context, antecedent_context)
	return nil
}

// NWContentContextSetAntecedent set an optional message context that must be sent before the context you are sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_antecedent(_:_:)
func NWContentContextSetAntecedent(context NWContentContext, antecedent_context NWContentContext) {
	if callErr := tryNWContentContextSetAntecedent(context, antecedent_context); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextSetExpirationMilliseconds func(context NWContentContext, expiration_milliseconds uint64)
var _nWContentContextSetExpirationMillisecondsErr error

func tryNWContentContextSetExpirationMilliseconds(context NWContentContext, expiration_milliseconds uint64) error {
	if _nWContentContextSetExpirationMilliseconds == nil {
		return symbolCallError("nw_content_context_set_expiration_milliseconds", "10.14", _nWContentContextSetExpirationMillisecondsErr)
	}
	_nWContentContextSetExpirationMilliseconds(context, expiration_milliseconds)
	return nil
}

// NWContentContextSetExpirationMilliseconds sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_expiration_milliseconds(_:_:)
func NWContentContextSetExpirationMilliseconds(context NWContentContext, expiration_milliseconds uint64) {
	if callErr := tryNWContentContextSetExpirationMilliseconds(context, expiration_milliseconds); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextSetIsFinal func(context NWContentContext, is_final bool)
var _nWContentContextSetIsFinalErr error

func tryNWContentContextSetIsFinal(context NWContentContext, is_final bool) error {
	if _nWContentContextSetIsFinal == nil {
		return symbolCallError("nw_content_context_set_is_final", "10.14", _nWContentContextSetIsFinalErr)
	}
	_nWContentContextSetIsFinal(context, is_final)
	return nil
}

// NWContentContextSetIsFinal sets a Boolean indicating if this context represents the final message being sent.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_is_final(_:_:)
func NWContentContextSetIsFinal(context NWContentContext, is_final bool) {
	if callErr := tryNWContentContextSetIsFinal(context, is_final); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextSetMetadataForProtocol func(context NWContentContext, protocol_metadata NWProtocolMetadata)
var _nWContentContextSetMetadataForProtocolErr error

func tryNWContentContextSetMetadataForProtocol(context NWContentContext, protocol_metadata NWProtocolMetadata) error {
	if _nWContentContextSetMetadataForProtocol == nil {
		return symbolCallError("nw_content_context_set_metadata_for_protocol", "10.14", _nWContentContextSetMetadataForProtocolErr)
	}
	_nWContentContextSetMetadataForProtocol(context, protocol_metadata)
	return nil
}

// NWContentContextSetMetadataForProtocol sets protocol metadata to configure per-message or per-packet properties.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_metadata_for_protocol(_:_:)
func NWContentContextSetMetadataForProtocol(context NWContentContext, protocol_metadata NWProtocolMetadata) {
	if callErr := tryNWContentContextSetMetadataForProtocol(context, protocol_metadata); callErr != nil {
		panic(callErr)
	}
}

var _nWContentContextSetRelativePriority func(context NWContentContext, relative_priority float64)
var _nWContentContextSetRelativePriorityErr error

func tryNWContentContextSetRelativePriority(context NWContentContext, relative_priority float64) error {
	if _nWContentContextSetRelativePriority == nil {
		return symbolCallError("nw_content_context_set_relative_priority", "10.14", _nWContentContextSetRelativePriorityErr)
	}
	_nWContentContextSetRelativePriority(context, relative_priority)
	return nil
}

// NWContentContextSetRelativePriority sets the relative value of priority used to reorder contexts when sending.
//
// See: https://developer.apple.com/documentation/Network/nw_content_context_set_relative_priority(_:_:)
func NWContentContextSetRelativePriority(context NWContentContext, relative_priority float64) {
	if callErr := tryNWContentContextSetRelativePriority(context, relative_priority); callErr != nil {
		panic(callErr)
	}
}

var _nWDataTransferReportCollect func(report NWDataTransferReport, queue uintptr, collect_block unsafe.Pointer)
var _nWDataTransferReportCollectErr error

func tryNWDataTransferReportCollect(report NWDataTransferReport, queue dispatch.Queue, collect_block NWDataTransferReportCollectBlock) error {
	if _nWDataTransferReportCollect == nil {
		return symbolCallError("nw_data_transfer_report_collect", "10.15", _nWDataTransferReportCollectErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { collect_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWDataTransferReportCollect(report, uintptr(queue.Handle()), _block0)
	return nil
}

// NWDataTransferReportCollect stops an outstanding data transfer report and calculates the results.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect(_:_:_:)
func NWDataTransferReportCollect(report NWDataTransferReport, queue dispatch.Queue, collect_block NWDataTransferReportCollectBlock) {
	if callErr := tryNWDataTransferReportCollect(report, queue, collect_block); callErr != nil {
		panic(callErr)
	}
}

var _nWDataTransferReportCopyPathInterface func(report NWDataTransferReport, path_index uint32) NWInterface
var _nWDataTransferReportCopyPathInterfaceErr error

func tryNWDataTransferReportCopyPathInterface(report NWDataTransferReport, path_index uint32) (NWInterface, error) {
	if _nWDataTransferReportCopyPathInterface == nil {
		return *new(NWInterface), symbolCallError("nw_data_transfer_report_copy_path_interface", "10.15", _nWDataTransferReportCopyPathInterfaceErr)
	}
	return _nWDataTransferReportCopyPathInterface(report, path_index), nil
}

// NWDataTransferReportCopyPathInterface accesses the network interface the path used.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_copy_path_interface(_:_:)
func NWDataTransferReportCopyPathInterface(report NWDataTransferReport, path_index uint32) NWInterface {
	result, callErr := tryNWDataTransferReportCopyPathInterface(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetDurationMilliseconds func(report NWDataTransferReport) uint64
var _nWDataTransferReportGetDurationMillisecondsErr error

func tryNWDataTransferReportGetDurationMilliseconds(report NWDataTransferReport) (uint64, error) {
	if _nWDataTransferReportGetDurationMilliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_duration_milliseconds", "10.15", _nWDataTransferReportGetDurationMillisecondsErr)
	}
	return _nWDataTransferReportGetDurationMilliseconds(report), nil
}

// NWDataTransferReportGetDurationMilliseconds checks the duration of the data transfer report, from when it was started to when it was collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_duration_milliseconds(_:)
func NWDataTransferReportGetDurationMilliseconds(report NWDataTransferReport) uint64 {
	result, callErr := tryNWDataTransferReportGetDurationMilliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetPathCount func(report NWDataTransferReport) uint32
var _nWDataTransferReportGetPathCountErr error

func tryNWDataTransferReportGetPathCount(report NWDataTransferReport) (uint32, error) {
	if _nWDataTransferReportGetPathCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_path_count", "10.15", _nWDataTransferReportGetPathCountErr)
	}
	return _nWDataTransferReportGetPathCount(report), nil
}

// NWDataTransferReportGetPathCount checks the number of valid paths in the report.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_count(_:)
func NWDataTransferReportGetPathCount(report NWDataTransferReport) uint32 {
	result, callErr := tryNWDataTransferReportGetPathCount(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetPathRadioType func(report NWDataTransferReport, path_index uint32) NWInterfaceRadioType
var _nWDataTransferReportGetPathRadioTypeErr error

func tryNWDataTransferReportGetPathRadioType(report NWDataTransferReport, path_index uint32) (NWInterfaceRadioType, error) {
	if _nWDataTransferReportGetPathRadioType == nil {
		return *new(NWInterfaceRadioType), symbolCallError("nw_data_transfer_report_get_path_radio_type", "12.0", _nWDataTransferReportGetPathRadioTypeErr)
	}
	return _nWDataTransferReportGetPathRadioType(report, path_index), nil
}

// NWDataTransferReportGetPathRadioType.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_radio_type(_:_:)
func NWDataTransferReportGetPathRadioType(report NWDataTransferReport, path_index uint32) NWInterfaceRadioType {
	result, callErr := tryNWDataTransferReportGetPathRadioType(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetReceivedApplicationByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetReceivedApplicationByteCountErr error

func tryNWDataTransferReportGetReceivedApplicationByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetReceivedApplicationByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_application_byte_count", "10.15", _nWDataTransferReportGetReceivedApplicationByteCountErr)
	}
	return _nWDataTransferReportGetReceivedApplicationByteCount(report, path_index), nil
}

// NWDataTransferReportGetReceivedApplicationByteCount accesses the number of bytes the connection delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_application_byte_count(_:_:)
func NWDataTransferReportGetReceivedApplicationByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetReceivedApplicationByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetReceivedIPPacketCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetReceivedIPPacketCountErr error

func tryNWDataTransferReportGetReceivedIPPacketCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetReceivedIPPacketCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_ip_packet_count", "10.15", _nWDataTransferReportGetReceivedIPPacketCountErr)
	}
	return _nWDataTransferReportGetReceivedIPPacketCount(report, path_index), nil
}

// NWDataTransferReportGetReceivedIPPacketCount accesses the number of IP packets the connection received.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_ip_packet_count(_:_:)
func NWDataTransferReportGetReceivedIPPacketCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetReceivedIPPacketCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetReceivedTransportByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetReceivedTransportByteCountErr error

func tryNWDataTransferReportGetReceivedTransportByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetReceivedTransportByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_byte_count", "10.15", _nWDataTransferReportGetReceivedTransportByteCountErr)
	}
	return _nWDataTransferReportGetReceivedTransportByteCount(report, path_index), nil
}

// NWDataTransferReportGetReceivedTransportByteCount accesses the number of bytes the transport protocol delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_byte_count(_:_:)
func NWDataTransferReportGetReceivedTransportByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetReceivedTransportByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetReceivedTransportDuplicateByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetReceivedTransportDuplicateByteCountErr error

func tryNWDataTransferReportGetReceivedTransportDuplicateByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetReceivedTransportDuplicateByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_duplicate_byte_count", "10.15", _nWDataTransferReportGetReceivedTransportDuplicateByteCountErr)
	}
	return _nWDataTransferReportGetReceivedTransportDuplicateByteCount(report, path_index), nil
}

// NWDataTransferReportGetReceivedTransportDuplicateByteCount accesses the number of duplicated bytes the transport protocol detected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_duplicate_byte_count(_:_:)
func NWDataTransferReportGetReceivedTransportDuplicateByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetReceivedTransportDuplicateByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetReceivedTransportOutOfOrderByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetReceivedTransportOutOfOrderByteCountErr error

func tryNWDataTransferReportGetReceivedTransportOutOfOrderByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetReceivedTransportOutOfOrderByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_received_transport_out_of_order_byte_count", "10.15", _nWDataTransferReportGetReceivedTransportOutOfOrderByteCountErr)
	}
	return _nWDataTransferReportGetReceivedTransportOutOfOrderByteCount(report, path_index), nil
}

// NWDataTransferReportGetReceivedTransportOutOfOrderByteCount accesses the number of bytes the transport protocol received out of order.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_out_of_order_byte_count(_:_:)
func NWDataTransferReportGetReceivedTransportOutOfOrderByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetReceivedTransportOutOfOrderByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetSentApplicationByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetSentApplicationByteCountErr error

func tryNWDataTransferReportGetSentApplicationByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetSentApplicationByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_application_byte_count", "10.15", _nWDataTransferReportGetSentApplicationByteCountErr)
	}
	return _nWDataTransferReportGetSentApplicationByteCount(report, path_index), nil
}

// NWDataTransferReportGetSentApplicationByteCount accesses the number of bytes sent on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_application_byte_count(_:_:)
func NWDataTransferReportGetSentApplicationByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetSentApplicationByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetSentIPPacketCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetSentIPPacketCountErr error

func tryNWDataTransferReportGetSentIPPacketCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetSentIPPacketCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_ip_packet_count", "10.15", _nWDataTransferReportGetSentIPPacketCountErr)
	}
	return _nWDataTransferReportGetSentIPPacketCount(report, path_index), nil
}

// NWDataTransferReportGetSentIPPacketCount accesses the number of IP packets the connection sent.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_ip_packet_count(_:_:)
func NWDataTransferReportGetSentIPPacketCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetSentIPPacketCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetSentTransportByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetSentTransportByteCountErr error

func tryNWDataTransferReportGetSentTransportByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetSentTransportByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_transport_byte_count", "10.15", _nWDataTransferReportGetSentTransportByteCountErr)
	}
	return _nWDataTransferReportGetSentTransportByteCount(report, path_index), nil
}

// NWDataTransferReportGetSentTransportByteCount accesses the number of bytes sent into the transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_byte_count(_:_:)
func NWDataTransferReportGetSentTransportByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetSentTransportByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetSentTransportRetransmittedByteCount func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetSentTransportRetransmittedByteCountErr error

func tryNWDataTransferReportGetSentTransportRetransmittedByteCount(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetSentTransportRetransmittedByteCount == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_sent_transport_retransmitted_byte_count", "10.15", _nWDataTransferReportGetSentTransportRetransmittedByteCountErr)
	}
	return _nWDataTransferReportGetSentTransportRetransmittedByteCount(report, path_index), nil
}

// NWDataTransferReportGetSentTransportRetransmittedByteCount accesses the number of bytes the transport protocol retransmitted.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(_:_:)
func NWDataTransferReportGetSentTransportRetransmittedByteCount(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetSentTransportRetransmittedByteCount(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetState func(report NWDataTransferReport) NWDataTransferReportState
var _nWDataTransferReportGetStateErr error

func tryNWDataTransferReportGetState(report NWDataTransferReport) (NWDataTransferReportState, error) {
	if _nWDataTransferReportGetState == nil {
		return *new(NWDataTransferReportState), symbolCallError("nw_data_transfer_report_get_state", "10.15", _nWDataTransferReportGetStateErr)
	}
	return _nWDataTransferReportGetState(report), nil
}

// NWDataTransferReportGetState checks whether a data transfer report is collected.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_state(_:)
func NWDataTransferReportGetState(report NWDataTransferReport) NWDataTransferReportState {
	result, callErr := tryNWDataTransferReportGetState(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetTransportMinimumRttMilliseconds func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetTransportMinimumRttMillisecondsErr error

func tryNWDataTransferReportGetTransportMinimumRttMilliseconds(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetTransportMinimumRttMilliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_minimum_rtt_milliseconds", "10.15", _nWDataTransferReportGetTransportMinimumRttMillisecondsErr)
	}
	return _nWDataTransferReportGetTransportMinimumRttMilliseconds(report, path_index), nil
}

// NWDataTransferReportGetTransportMinimumRttMilliseconds accesses the minimum round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(_:_:)
func NWDataTransferReportGetTransportMinimumRttMilliseconds(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetTransportMinimumRttMilliseconds(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetTransportRttVariance func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetTransportRttVarianceErr error

func tryNWDataTransferReportGetTransportRttVariance(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetTransportRttVariance == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_rtt_variance", "10.15", _nWDataTransferReportGetTransportRttVarianceErr)
	}
	return _nWDataTransferReportGetTransportRttVariance(report, path_index), nil
}

// NWDataTransferReportGetTransportRttVariance accesses the variance of the round-trip time the transport protocol measured.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_rtt_variance(_:_:)
func NWDataTransferReportGetTransportRttVariance(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetTransportRttVariance(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWDataTransferReportGetTransportSmoothedRttMilliseconds func(report NWDataTransferReport, path_index uint32) uint64
var _nWDataTransferReportGetTransportSmoothedRttMillisecondsErr error

func tryNWDataTransferReportGetTransportSmoothedRttMilliseconds(report NWDataTransferReport, path_index uint32) (uint64, error) {
	if _nWDataTransferReportGetTransportSmoothedRttMilliseconds == nil {
		return 0, symbolCallError("nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds", "10.15", _nWDataTransferReportGetTransportSmoothedRttMillisecondsErr)
	}
	return _nWDataTransferReportGetTransportSmoothedRttMilliseconds(report, path_index), nil
}

// NWDataTransferReportGetTransportSmoothedRttMilliseconds accesses the smoothed round-trip time the transport protocol measured, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(_:_:)
func NWDataTransferReportGetTransportSmoothedRttMilliseconds(report NWDataTransferReport, path_index uint32) uint64 {
	result, callErr := tryNWDataTransferReportGetTransportSmoothedRttMilliseconds(report, path_index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCopyAddressString func(endpoint NWEndpoint) *byte
var _nWEndpointCopyAddressStringErr error

func tryNWEndpointCopyAddressString(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointCopyAddressString == nil {
		return nil, symbolCallError("nw_endpoint_copy_address_string", "10.14", _nWEndpointCopyAddressStringErr)
	}
	return _nWEndpointCopyAddressString(endpoint), nil
}

// NWEndpointCopyAddressString copies the address of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_address_string(_:)
func NWEndpointCopyAddressString(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointCopyAddressString(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCopyPortString func(endpoint NWEndpoint) *byte
var _nWEndpointCopyPortStringErr error

func tryNWEndpointCopyPortString(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointCopyPortString == nil {
		return nil, symbolCallError("nw_endpoint_copy_port_string", "10.14", _nWEndpointCopyPortStringErr)
	}
	return _nWEndpointCopyPortString(endpoint), nil
}

// NWEndpointCopyPortString copies the port of an endpoint as a string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_port_string(_:)
func NWEndpointCopyPortString(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointCopyPortString(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCopyTXTRecord func(endpoint NWEndpoint) NWTXTRecord
var _nWEndpointCopyTXTRecordErr error

func tryNWEndpointCopyTXTRecord(endpoint NWEndpoint) (NWTXTRecord, error) {
	if _nWEndpointCopyTXTRecord == nil {
		return *new(NWTXTRecord), symbolCallError("nw_endpoint_copy_txt_record", "13.0", _nWEndpointCopyTXTRecordErr)
	}
	return _nWEndpointCopyTXTRecord(endpoint), nil
}

// NWEndpointCopyTXTRecord.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_copy_txt_record(_:)
func NWEndpointCopyTXTRecord(endpoint NWEndpoint) NWTXTRecord {
	result, callErr := tryNWEndpointCopyTXTRecord(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCreateAddress func(address unsafe.Pointer) NWEndpoint
var _nWEndpointCreateAddressErr error

func tryNWEndpointCreateAddress(address unsafe.Pointer) (NWEndpoint, error) {
	if _nWEndpointCreateAddress == nil {
		return NWEndpoint{}, symbolCallError("nw_endpoint_create_address", "10.14", _nWEndpointCreateAddressErr)
	}
	return _nWEndpointCreateAddress(address), nil
}

// NWEndpointCreateAddress creates a network endpoint with an address structure.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_address(_:)
func NWEndpointCreateAddress(address unsafe.Pointer) NWEndpoint {
	result, callErr := tryNWEndpointCreateAddress(address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCreateBonjourService func(name string, type_ string, domain string) NWEndpoint
var _nWEndpointCreateBonjourServiceErr error

func tryNWEndpointCreateBonjourService(name string, type_ string, domain string) (NWEndpoint, error) {
	if _nWEndpointCreateBonjourService == nil {
		return NWEndpoint{}, symbolCallError("nw_endpoint_create_bonjour_service", "10.14", _nWEndpointCreateBonjourServiceErr)
	}
	return _nWEndpointCreateBonjourService(name, type_, domain), nil
}

// NWEndpointCreateBonjourService creates a network endpoint with a Bonjour service name, type, and domain.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_bonjour_service(_:_:_:)
func NWEndpointCreateBonjourService(name string, type_ string, domain string) NWEndpoint {
	result, callErr := tryNWEndpointCreateBonjourService(name, type_, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCreateHost func(hostname string, port string) NWEndpoint
var _nWEndpointCreateHostErr error

func tryNWEndpointCreateHost(hostname string, port string) (NWEndpoint, error) {
	if _nWEndpointCreateHost == nil {
		return NWEndpoint{}, symbolCallError("nw_endpoint_create_host", "10.14", _nWEndpointCreateHostErr)
	}
	return _nWEndpointCreateHost(hostname, port), nil
}

// NWEndpointCreateHost creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_host(_:_:)
func NWEndpointCreateHost(hostname string, port string) NWEndpoint {
	result, callErr := tryNWEndpointCreateHost(hostname, port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointCreateURL func(url string) NWEndpoint
var _nWEndpointCreateURLErr error

func tryNWEndpointCreateURL(url string) (NWEndpoint, error) {
	if _nWEndpointCreateURL == nil {
		return NWEndpoint{}, symbolCallError("nw_endpoint_create_url", "10.15", _nWEndpointCreateURLErr)
	}
	return _nWEndpointCreateURL(url), nil
}

// NWEndpointCreateURL creates a network endpoint with a URL string.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_create_url(_:)
func NWEndpointCreateURL(url string) NWEndpoint {
	result, callErr := tryNWEndpointCreateURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetAddress func(endpoint NWEndpoint) unsafe.Pointer
var _nWEndpointGetAddressErr error

func tryNWEndpointGetAddress(endpoint NWEndpoint) (unsafe.Pointer, error) {
	if _nWEndpointGetAddress == nil {
		return nil, symbolCallError("nw_endpoint_get_address", "10.14", _nWEndpointGetAddressErr)
	}
	return _nWEndpointGetAddress(endpoint), nil
}

// NWEndpointGetAddress accesses the address structure stored in an address endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_address(_:)
func NWEndpointGetAddress(endpoint NWEndpoint) unsafe.Pointer {
	result, callErr := tryNWEndpointGetAddress(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetBonjourServiceDomain func(endpoint NWEndpoint) *byte
var _nWEndpointGetBonjourServiceDomainErr error

func tryNWEndpointGetBonjourServiceDomain(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointGetBonjourServiceDomain == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_domain", "10.14", _nWEndpointGetBonjourServiceDomainErr)
	}
	return _nWEndpointGetBonjourServiceDomain(endpoint), nil
}

// NWEndpointGetBonjourServiceDomain accesses the Bonjour service domain stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_domain(_:)
func NWEndpointGetBonjourServiceDomain(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointGetBonjourServiceDomain(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetBonjourServiceName func(endpoint NWEndpoint) *byte
var _nWEndpointGetBonjourServiceNameErr error

func tryNWEndpointGetBonjourServiceName(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointGetBonjourServiceName == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_name", "10.14", _nWEndpointGetBonjourServiceNameErr)
	}
	return _nWEndpointGetBonjourServiceName(endpoint), nil
}

// NWEndpointGetBonjourServiceName accesses the Bonjour service name stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_name(_:)
func NWEndpointGetBonjourServiceName(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointGetBonjourServiceName(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetBonjourServiceType func(endpoint NWEndpoint) *byte
var _nWEndpointGetBonjourServiceTypeErr error

func tryNWEndpointGetBonjourServiceType(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointGetBonjourServiceType == nil {
		return nil, symbolCallError("nw_endpoint_get_bonjour_service_type", "10.14", _nWEndpointGetBonjourServiceTypeErr)
	}
	return _nWEndpointGetBonjourServiceType(endpoint), nil
}

// NWEndpointGetBonjourServiceType accesses the Bonjour service type stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_type(_:)
func NWEndpointGetBonjourServiceType(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointGetBonjourServiceType(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetHostname func(endpoint NWEndpoint) *byte
var _nWEndpointGetHostnameErr error

func tryNWEndpointGetHostname(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointGetHostname == nil {
		return nil, symbolCallError("nw_endpoint_get_hostname", "10.14", _nWEndpointGetHostnameErr)
	}
	return _nWEndpointGetHostname(endpoint), nil
}

// NWEndpointGetHostname accesses the hostname stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_hostname(_:)
func NWEndpointGetHostname(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointGetHostname(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetPort func(endpoint NWEndpoint) uint16
var _nWEndpointGetPortErr error

func tryNWEndpointGetPort(endpoint NWEndpoint) (uint16, error) {
	if _nWEndpointGetPort == nil {
		return 0, symbolCallError("nw_endpoint_get_port", "10.14", _nWEndpointGetPortErr)
	}
	return _nWEndpointGetPort(endpoint), nil
}

// NWEndpointGetPort accesses the port stored in an endpoint, in host-byte order.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_port(_:)
func NWEndpointGetPort(endpoint NWEndpoint) uint16 {
	result, callErr := tryNWEndpointGetPort(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetSignature func(endpoint NWEndpoint, out_signature_length *uintptr) *uint8
var _nWEndpointGetSignatureErr error

func tryNWEndpointGetSignature(endpoint NWEndpoint, out_signature_length *uintptr) (*uint8, error) {
	if _nWEndpointGetSignature == nil {
		return nil, symbolCallError("nw_endpoint_get_signature", "13.0", _nWEndpointGetSignatureErr)
	}
	return _nWEndpointGetSignature(endpoint, out_signature_length), nil
}

// NWEndpointGetSignature.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_signature(_:_:)
func NWEndpointGetSignature(endpoint NWEndpoint, out_signature_length *uintptr) *uint8 {
	result, callErr := tryNWEndpointGetSignature(endpoint, out_signature_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetType func(endpoint NWEndpoint) NWEndpointType
var _nWEndpointGetTypeErr error

func tryNWEndpointGetType(endpoint NWEndpoint) (NWEndpointType, error) {
	if _nWEndpointGetType == nil {
		return *new(NWEndpointType), symbolCallError("nw_endpoint_get_type", "10.14", _nWEndpointGetTypeErr)
	}
	return _nWEndpointGetType(endpoint), nil
}

// NWEndpointGetType accesses the type of a endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_type(_:)
func NWEndpointGetType(endpoint NWEndpoint) NWEndpointType {
	result, callErr := tryNWEndpointGetType(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEndpointGetURL func(endpoint NWEndpoint) *byte
var _nWEndpointGetURLErr error

func tryNWEndpointGetURL(endpoint NWEndpoint) (*byte, error) {
	if _nWEndpointGetURL == nil {
		return nil, symbolCallError("nw_endpoint_get_url", "10.15", _nWEndpointGetURLErr)
	}
	return _nWEndpointGetURL(endpoint), nil
}

// NWEndpointGetURL accesses the URL string stored in an endpoint.
//
// See: https://developer.apple.com/documentation/Network/nw_endpoint_get_url(_:)
func NWEndpointGetURL(endpoint NWEndpoint) *byte {
	result, callErr := tryNWEndpointGetURL(endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWErrorCopyCfError func(err NWError) corefoundation.CFErrorRef
var _nWErrorCopyCfErrorErr error

func tryNWErrorCopyCfError(err NWError) (corefoundation.CFErrorRef, error) {
	if _nWErrorCopyCfError == nil {
		return 0, symbolCallError("nw_error_copy_cf_error", "10.14", _nWErrorCopyCfErrorErr)
	}
	return _nWErrorCopyCfError(err), nil
}

// NWErrorCopyCfError returns a copy of a network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_copy_cf_error(_:)
func NWErrorCopyCfError(err NWError) corefoundation.CFErrorRef {
	result, callErr := tryNWErrorCopyCfError(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWErrorGetErrorCode func(err NWError) int
var _nWErrorGetErrorCodeErr error

func tryNWErrorGetErrorCode(err NWError) (int, error) {
	if _nWErrorGetErrorCode == nil {
		return 0, symbolCallError("nw_error_get_error_code", "10.14", _nWErrorGetErrorCodeErr)
	}
	return _nWErrorGetErrorCode(err), nil
}

// NWErrorGetErrorCode accesses the specific code of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_code(_:)
func NWErrorGetErrorCode(err NWError) int {
	result, callErr := tryNWErrorGetErrorCode(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWErrorGetErrorDomain func(err NWError) NWErrorDomain
var _nWErrorGetErrorDomainErr error

func tryNWErrorGetErrorDomain(err NWError) (NWErrorDomain, error) {
	if _nWErrorGetErrorDomain == nil {
		return *new(NWErrorDomain), symbolCallError("nw_error_get_error_domain", "10.14", _nWErrorGetErrorDomainErr)
	}
	return _nWErrorGetErrorDomain(err), nil
}

// NWErrorGetErrorDomain accesses the domain of the network error.
//
// See: https://developer.apple.com/documentation/Network/nw_error_get_error_domain(_:)
func NWErrorGetErrorDomain(err NWError) NWErrorDomain {
	result, callErr := tryNWErrorGetErrorDomain(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportCopyProxyEndpoint func(report NWEstablishmentReport) NWEndpoint
var _nWEstablishmentReportCopyProxyEndpointErr error

func tryNWEstablishmentReportCopyProxyEndpoint(report NWEstablishmentReport) (NWEndpoint, error) {
	if _nWEstablishmentReportCopyProxyEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_establishment_report_copy_proxy_endpoint", "10.15", _nWEstablishmentReportCopyProxyEndpointErr)
	}
	return _nWEstablishmentReportCopyProxyEndpoint(report), nil
}

// NWEstablishmentReportCopyProxyEndpoint accesses the endpoint of the proxy the connection used.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_copy_proxy_endpoint(_:)
func NWEstablishmentReportCopyProxyEndpoint(report NWEstablishmentReport) NWEndpoint {
	result, callErr := tryNWEstablishmentReportCopyProxyEndpoint(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportEnumerateProtocols func(report NWEstablishmentReport, enumerate_block unsafe.Pointer)
var _nWEstablishmentReportEnumerateProtocolsErr error

func tryNWEstablishmentReportEnumerateProtocols(report NWEstablishmentReport, enumerate_block NWReportProtocolEnumerator) error {
	if _nWEstablishmentReportEnumerateProtocols == nil {
		return symbolCallError("nw_establishment_report_enumerate_protocols", "10.15", _nWEstablishmentReportEnumerateProtocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 uint64, blockArg2 uint64) bool {
		return enumerate_block(objectivec.ObjectFromID(blockArg0), blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWEstablishmentReportEnumerateProtocols(report, _block0)
	return nil
}

// NWEstablishmentReportEnumerateProtocols iterates a list of protocol handshakes in order from first completed to last completed.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_protocols(_:_:)
func NWEstablishmentReportEnumerateProtocols(report NWEstablishmentReport, enumerate_block NWReportProtocolEnumerator) {
	if callErr := tryNWEstablishmentReportEnumerateProtocols(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWEstablishmentReportEnumerateResolutionReports func(report NWEstablishmentReport, enumerate_block unsafe.Pointer)
var _nWEstablishmentReportEnumerateResolutionReportsErr error

func tryNWEstablishmentReportEnumerateResolutionReports(report NWEstablishmentReport, enumerate_block NWReportResolutionReportEnumerator) error {
	if _nWEstablishmentReportEnumerateResolutionReports == nil {
		return symbolCallError("nw_establishment_report_enumerate_resolution_reports", "11.0", _nWEstablishmentReportEnumerateResolutionReportsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return enumerate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWEstablishmentReportEnumerateResolutionReports(report, _block0)
	return nil
}

// NWEstablishmentReportEnumerateResolutionReports.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolution_reports(_:_:)
func NWEstablishmentReportEnumerateResolutionReports(report NWEstablishmentReport, enumerate_block NWReportResolutionReportEnumerator) {
	if callErr := tryNWEstablishmentReportEnumerateResolutionReports(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWEstablishmentReportEnumerateResolutions func(report NWEstablishmentReport, enumerate_block unsafe.Pointer)
var _nWEstablishmentReportEnumerateResolutionsErr error

func tryNWEstablishmentReportEnumerateResolutions(report NWEstablishmentReport, enumerate_block NWReportResolutionEnumerator) error {
	if _nWEstablishmentReportEnumerateResolutions == nil {
		return symbolCallError("nw_establishment_report_enumerate_resolutions", "10.15", _nWEstablishmentReportEnumerateResolutionsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWReportResolutionSource, blockArg1 uint64, blockArg2 uint32, blockArg3 objc.ID, blockArg4 objc.ID) bool {
		return enumerate_block(blockArg0, blockArg1, blockArg2, objectivec.ObjectFromID(blockArg3), objectivec.ObjectFromID(blockArg4))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWEstablishmentReportEnumerateResolutions(report, _block0)
	return nil
}

// NWEstablishmentReportEnumerateResolutions iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolutions(_:_:)
func NWEstablishmentReportEnumerateResolutions(report NWEstablishmentReport, enumerate_block NWReportResolutionEnumerator) {
	if callErr := tryNWEstablishmentReportEnumerateResolutions(report, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWEstablishmentReportGetAttemptStartedAfterMilliseconds func(report NWEstablishmentReport) uint64
var _nWEstablishmentReportGetAttemptStartedAfterMillisecondsErr error

func tryNWEstablishmentReportGetAttemptStartedAfterMilliseconds(report NWEstablishmentReport) (uint64, error) {
	if _nWEstablishmentReportGetAttemptStartedAfterMilliseconds == nil {
		return 0, symbolCallError("nw_establishment_report_get_attempt_started_after_milliseconds", "10.15", _nWEstablishmentReportGetAttemptStartedAfterMillisecondsErr)
	}
	return _nWEstablishmentReportGetAttemptStartedAfterMilliseconds(report), nil
}

// NWEstablishmentReportGetAttemptStartedAfterMilliseconds accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_attempt_started_after_milliseconds(_:)
func NWEstablishmentReportGetAttemptStartedAfterMilliseconds(report NWEstablishmentReport) uint64 {
	result, callErr := tryNWEstablishmentReportGetAttemptStartedAfterMilliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportGetDurationMilliseconds func(report NWEstablishmentReport) uint64
var _nWEstablishmentReportGetDurationMillisecondsErr error

func tryNWEstablishmentReportGetDurationMilliseconds(report NWEstablishmentReport) (uint64, error) {
	if _nWEstablishmentReportGetDurationMilliseconds == nil {
		return 0, symbolCallError("nw_establishment_report_get_duration_milliseconds", "10.15", _nWEstablishmentReportGetDurationMillisecondsErr)
	}
	return _nWEstablishmentReportGetDurationMilliseconds(report), nil
}

// NWEstablishmentReportGetDurationMilliseconds checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_duration_milliseconds(_:)
func NWEstablishmentReportGetDurationMilliseconds(report NWEstablishmentReport) uint64 {
	result, callErr := tryNWEstablishmentReportGetDurationMilliseconds(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportGetPreviousAttemptCount func(report NWEstablishmentReport) uint32
var _nWEstablishmentReportGetPreviousAttemptCountErr error

func tryNWEstablishmentReportGetPreviousAttemptCount(report NWEstablishmentReport) (uint32, error) {
	if _nWEstablishmentReportGetPreviousAttemptCount == nil {
		return 0, symbolCallError("nw_establishment_report_get_previous_attempt_count", "10.15", _nWEstablishmentReportGetPreviousAttemptCountErr)
	}
	return _nWEstablishmentReportGetPreviousAttemptCount(report), nil
}

// NWEstablishmentReportGetPreviousAttemptCount checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_previous_attempt_count(_:)
func NWEstablishmentReportGetPreviousAttemptCount(report NWEstablishmentReport) uint32 {
	result, callErr := tryNWEstablishmentReportGetPreviousAttemptCount(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportGetProxyConfigured func(report NWEstablishmentReport) bool
var _nWEstablishmentReportGetProxyConfiguredErr error

func tryNWEstablishmentReportGetProxyConfigured(report NWEstablishmentReport) (bool, error) {
	if _nWEstablishmentReportGetProxyConfigured == nil {
		return false, symbolCallError("nw_establishment_report_get_proxy_configured", "10.15", _nWEstablishmentReportGetProxyConfiguredErr)
	}
	return _nWEstablishmentReportGetProxyConfigured(report), nil
}

// NWEstablishmentReportGetProxyConfigured checks whether a proxy was configured on the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_proxy_configured(_:)
func NWEstablishmentReportGetProxyConfigured(report NWEstablishmentReport) bool {
	result, callErr := tryNWEstablishmentReportGetProxyConfigured(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEstablishmentReportGetUsedProxy func(report NWEstablishmentReport) bool
var _nWEstablishmentReportGetUsedProxyErr error

func tryNWEstablishmentReportGetUsedProxy(report NWEstablishmentReport) (bool, error) {
	if _nWEstablishmentReportGetUsedProxy == nil {
		return false, symbolCallError("nw_establishment_report_get_used_proxy", "10.15", _nWEstablishmentReportGetUsedProxyErr)
	}
	return _nWEstablishmentReportGetUsedProxy(report), nil
}

// NWEstablishmentReportGetUsedProxy checks whether the connection used a proxy.
//
// See: https://developer.apple.com/documentation/Network/nw_establishment_report_get_used_proxy(_:)
func NWEstablishmentReportGetUsedProxy(report NWEstablishmentReport) bool {
	result, callErr := tryNWEstablishmentReportGetUsedProxy(report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEthernetChannelCancel func(ethernet_channel NWEthernetChannel)
var _nWEthernetChannelCancelErr error

func tryNWEthernetChannelCancel(ethernet_channel NWEthernetChannel) error {
	if _nWEthernetChannelCancel == nil {
		return symbolCallError("nw_ethernet_channel_cancel", "10.15", _nWEthernetChannelCancelErr)
	}
	_nWEthernetChannelCancel(ethernet_channel)
	return nil
}

// NWEthernetChannelCancel unregisters the channel from the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_cancel(_:)
func NWEthernetChannelCancel(ethernet_channel NWEthernetChannel) {
	if callErr := tryNWEthernetChannelCancel(ethernet_channel); callErr != nil {
		panic(callErr)
	}
}

var _nWEthernetChannelCreate func(ether_type uint16, interface_ NWInterface) NWEthernetChannel
var _nWEthernetChannelCreateErr error

func tryNWEthernetChannelCreate(ether_type uint16, interface_ NWInterface) (NWEthernetChannel, error) {
	if _nWEthernetChannelCreate == nil {
		return *new(NWEthernetChannel), symbolCallError("nw_ethernet_channel_create", "10.15", _nWEthernetChannelCreateErr)
	}
	return _nWEthernetChannelCreate(ether_type, interface_), nil
}

// NWEthernetChannelCreate initializes an Ethernet channel on a specific interface with a custom Ethernet type.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create(_:_:)
func NWEthernetChannelCreate(ether_type uint16, interface_ NWInterface) NWEthernetChannel {
	result, callErr := tryNWEthernetChannelCreate(ether_type, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEthernetChannelCreateWithParameters func(ether_type uint16, interface_ NWInterface, parameters NWParameters) NWEthernetChannel
var _nWEthernetChannelCreateWithParametersErr error

func tryNWEthernetChannelCreateWithParameters(ether_type uint16, interface_ NWInterface, parameters NWParameters) (NWEthernetChannel, error) {
	if _nWEthernetChannelCreateWithParameters == nil {
		return *new(NWEthernetChannel), symbolCallError("nw_ethernet_channel_create_with_parameters", "13.0", _nWEthernetChannelCreateWithParametersErr)
	}
	return _nWEthernetChannelCreateWithParameters(ether_type, interface_, parameters), nil
}

// NWEthernetChannelCreateWithParameters.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create_with_parameters(_:_:_:)
func NWEthernetChannelCreateWithParameters(ether_type uint16, interface_ NWInterface, parameters NWParameters) NWEthernetChannel {
	result, callErr := tryNWEthernetChannelCreateWithParameters(ether_type, interface_, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEthernetChannelGetMaximumPayloadSize func(ethernet_channel NWEthernetChannel) uint32
var _nWEthernetChannelGetMaximumPayloadSizeErr error

func tryNWEthernetChannelGetMaximumPayloadSize(ethernet_channel NWEthernetChannel) (uint32, error) {
	if _nWEthernetChannelGetMaximumPayloadSize == nil {
		return 0, symbolCallError("nw_ethernet_channel_get_maximum_payload_size", "13.0", _nWEthernetChannelGetMaximumPayloadSizeErr)
	}
	return _nWEthernetChannelGetMaximumPayloadSize(ethernet_channel), nil
}

// NWEthernetChannelGetMaximumPayloadSize.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_get_maximum_payload_size(_:)
func NWEthernetChannelGetMaximumPayloadSize(ethernet_channel NWEthernetChannel) uint32 {
	result, callErr := tryNWEthernetChannelGetMaximumPayloadSize(ethernet_channel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWEthernetChannelSend func(ethernet_channel NWEthernetChannel, content uintptr, vlan_tag uint16, remote_address NWEthernetAddress, completion unsafe.Pointer)
var _nWEthernetChannelSendErr error

func tryNWEthernetChannelSend(ethernet_channel NWEthernetChannel, content dispatch.Data, vlan_tag uint16, remote_address NWEthernetAddress, completion NWEthernetChannelSendCompletion) error {
	if _nWEthernetChannelSend == nil {
		return symbolCallError("nw_ethernet_channel_send", "10.15", _nWEthernetChannelSendErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { completion(NWError{Object: objectivec.ObjectFromID(blockArg0)}) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWEthernetChannelSend(ethernet_channel, uintptr(content.Handle()), vlan_tag, remote_address, _block0)
	return nil
}

// NWEthernetChannelSend sends a single Ethernet frame over a channel to a specific Ethernet address.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send(_:_:_:_:_:)
func NWEthernetChannelSend(ethernet_channel NWEthernetChannel, content dispatch.Data, vlan_tag uint16, remote_address NWEthernetAddress, completion NWEthernetChannelSendCompletion) {
	if callErr := tryNWEthernetChannelSend(ethernet_channel, content, vlan_tag, remote_address, completion); callErr != nil {
		panic(callErr)
	}
}

var _nWEthernetChannelSetQueue func(ethernet_channel NWEthernetChannel, queue uintptr)
var _nWEthernetChannelSetQueueErr error

func tryNWEthernetChannelSetQueue(ethernet_channel NWEthernetChannel, queue dispatch.Queue) error {
	if _nWEthernetChannelSetQueue == nil {
		return symbolCallError("nw_ethernet_channel_set_queue", "10.15", _nWEthernetChannelSetQueueErr)
	}
	_nWEthernetChannelSetQueue(ethernet_channel, uintptr(queue.Handle()))
	return nil
}

// NWEthernetChannelSetQueue sets the queue on which all channel events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_queue(_:_:)
func NWEthernetChannelSetQueue(ethernet_channel NWEthernetChannel, queue dispatch.Queue) {
	if callErr := tryNWEthernetChannelSetQueue(ethernet_channel, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWEthernetChannelSetReceiveHandler func(ethernet_channel NWEthernetChannel, handler unsafe.Pointer)
var _nWEthernetChannelSetReceiveHandlerErr error

func tryNWEthernetChannelSetReceiveHandler(ethernet_channel NWEthernetChannel, handler NWEthernetChannelReceiveHandler) error {
	if _nWEthernetChannelSetReceiveHandler == nil {
		return symbolCallError("nw_ethernet_channel_set_receive_handler", "10.15", _nWEthernetChannelSetReceiveHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 uint32, blockArg2 *uint8, blockArg3 *uint8) {
		handler(objectivec.ObjectFromID(blockArg0), blockArg1, blockArg2, blockArg3)
	})
	retainNetworkAsyncBlock(ethernet_channel.ID, "nw_ethernet_channel_set_receive_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWEthernetChannelSetReceiveHandler(ethernet_channel, _block0)
	return nil
}

// NWEthernetChannelSetReceiveHandler sets a handler to receive inbound Ethernet frames.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_receive_handler(_:_:)
func NWEthernetChannelSetReceiveHandler(ethernet_channel NWEthernetChannel, handler NWEthernetChannelReceiveHandler) {
	if callErr := tryNWEthernetChannelSetReceiveHandler(ethernet_channel, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWEthernetChannelSetStateChangedHandler func(ethernet_channel NWEthernetChannel, handler unsafe.Pointer)
var _nWEthernetChannelSetStateChangedHandlerErr error

func tryNWEthernetChannelSetStateChangedHandler(ethernet_channel NWEthernetChannel, handler NWEthernetChannelStateChangedHandler) error {
	if _nWEthernetChannelSetStateChangedHandler == nil {
		return symbolCallError("nw_ethernet_channel_set_state_changed_handler", "10.15", _nWEthernetChannelSetStateChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWEthernetChannelState, blockArg1 objc.ID) {
		handler(blockArg0, NWError{Object: objectivec.ObjectFromID(blockArg1)})
	})
	retainNetworkAsyncBlock(ethernet_channel.ID, "nw_ethernet_channel_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWEthernetChannelSetStateChangedHandler(ethernet_channel, _block0)
	return nil
}

// NWEthernetChannelSetStateChangedHandler sets a handler to receive channel state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_state_changed_handler(_:_:)
func NWEthernetChannelSetStateChangedHandler(ethernet_channel NWEthernetChannel, handler NWEthernetChannelStateChangedHandler) {
	if callErr := tryNWEthernetChannelSetStateChangedHandler(ethernet_channel, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWEthernetChannelStart func(ethernet_channel NWEthernetChannel)
var _nWEthernetChannelStartErr error

func tryNWEthernetChannelStart(ethernet_channel NWEthernetChannel) error {
	if _nWEthernetChannelStart == nil {
		return symbolCallError("nw_ethernet_channel_start", "10.15", _nWEthernetChannelStartErr)
	}
	_nWEthernetChannelStart(ethernet_channel)
	return nil
}

// NWEthernetChannelStart starts the process of registering the channel.
//
// See: https://developer.apple.com/documentation/Network/nw_ethernet_channel_start(_:)
func NWEthernetChannelStart(ethernet_channel NWEthernetChannel) {
	if callErr := tryNWEthernetChannelStart(ethernet_channel); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerAsync func(framer NWFramer, async_block unsafe.Pointer)
var _nWFramerAsyncErr error

func tryNWFramerAsync(framer NWFramer, async_block NWFramerBlock) error {
	if _nWFramerAsync == nil {
		return symbolCallError("nw_framer_async", "10.15", _nWFramerAsyncErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { async_block() })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerAsync(framer, _block0)
	return nil
}

// NWFramerAsync requests that a block be executed on the connection’s internal scheduling context.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_async(_:_:)
func NWFramerAsync(framer NWFramer, async_block NWFramerBlock) {
	if callErr := tryNWFramerAsync(framer, async_block); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerCopyLocalEndpoint func(framer NWFramer) NWEndpoint
var _nWFramerCopyLocalEndpointErr error

func tryNWFramerCopyLocalEndpoint(framer NWFramer) (NWEndpoint, error) {
	if _nWFramerCopyLocalEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_framer_copy_local_endpoint", "10.15", _nWFramerCopyLocalEndpointErr)
	}
	return _nWFramerCopyLocalEndpoint(framer), nil
}

// NWFramerCopyLocalEndpoint accesses the local endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_local_endpoint(_:)
func NWFramerCopyLocalEndpoint(framer NWFramer) NWEndpoint {
	result, callErr := tryNWFramerCopyLocalEndpoint(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerCopyOptions func(framer NWFramer) NWProtocolOptions
var _nWFramerCopyOptionsErr error

func tryNWFramerCopyOptions(framer NWFramer) (NWProtocolOptions, error) {
	if _nWFramerCopyOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_framer_copy_options", "12.3", _nWFramerCopyOptionsErr)
	}
	return _nWFramerCopyOptions(framer), nil
}

// NWFramerCopyOptions.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_options(_:)
func NWFramerCopyOptions(framer NWFramer) NWProtocolOptions {
	result, callErr := tryNWFramerCopyOptions(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerCopyParameters func(framer NWFramer) NWParameters
var _nWFramerCopyParametersErr error

func tryNWFramerCopyParameters(framer NWFramer) (NWParameters, error) {
	if _nWFramerCopyParameters == nil {
		return *new(NWParameters), symbolCallError("nw_framer_copy_parameters", "10.15", _nWFramerCopyParametersErr)
	}
	return _nWFramerCopyParameters(framer), nil
}

// NWFramerCopyParameters accesses the parameters of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_parameters(_:)
func NWFramerCopyParameters(framer NWFramer) NWParameters {
	result, callErr := tryNWFramerCopyParameters(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerCopyRemoteEndpoint func(framer NWFramer) NWEndpoint
var _nWFramerCopyRemoteEndpointErr error

func tryNWFramerCopyRemoteEndpoint(framer NWFramer) (NWEndpoint, error) {
	if _nWFramerCopyRemoteEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_framer_copy_remote_endpoint", "10.15", _nWFramerCopyRemoteEndpointErr)
	}
	return _nWFramerCopyRemoteEndpoint(framer), nil
}

// NWFramerCopyRemoteEndpoint accesses the remote endpoint of the connection in which your protocol is running.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_copy_remote_endpoint(_:)
func NWFramerCopyRemoteEndpoint(framer NWFramer) NWEndpoint {
	result, callErr := tryNWFramerCopyRemoteEndpoint(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerCreateDefinition func(identifier string, flags uint32, start_handler unsafe.Pointer) NWProtocolDefinition
var _nWFramerCreateDefinitionErr error

func tryNWFramerCreateDefinition(identifier string, flags uint32, start_handler NWFramerStartHandler) (NWProtocolDefinition, error) {
	if _nWFramerCreateDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_framer_create_definition", "10.15", _nWFramerCreateDefinitionErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) NWFramerStartResult {
		return start_handler(objectivec.ObjectFromID(blockArg0))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWFramerCreateDefinition(identifier, flags, _block0), nil
}

// NWFramerCreateDefinition initializes a new protocol definition based on your protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_definition(_:_:_:)
func NWFramerCreateDefinition(identifier string, flags uint32, start_handler NWFramerStartHandler) NWProtocolDefinition {
	result, callErr := tryNWFramerCreateDefinition(identifier, flags, start_handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerCreateOptions func(framer_definition NWProtocolDefinition) NWProtocolOptions
var _nWFramerCreateOptionsErr error

func tryNWFramerCreateOptions(framer_definition NWProtocolDefinition) (NWProtocolOptions, error) {
	if _nWFramerCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_framer_create_options", "10.15", _nWFramerCreateOptionsErr)
	}
	return _nWFramerCreateOptions(framer_definition), nil
}

// NWFramerCreateOptions initializes a set of protocol options with a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_create_options(_:)
func NWFramerCreateOptions(framer_definition NWProtocolDefinition) NWProtocolOptions {
	result, callErr := tryNWFramerCreateOptions(framer_definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerDeliverInput func(framer NWFramer, input_buffer *byte, input_length uintptr, message NWFramerMessage, is_complete bool)
var _nWFramerDeliverInputErr error

func tryNWFramerDeliverInput(framer NWFramer, input_buffer []byte, input_length uintptr, message NWFramerMessage, is_complete bool) error {
	if _nWFramerDeliverInput == nil {
		return symbolCallError("nw_framer_deliver_input", "10.15", _nWFramerDeliverInputErr)
	}
	_nWFramerDeliverInput(framer, unsafe.SliceData(input_buffer), input_length, message, is_complete)
	return nil
}

// NWFramerDeliverInput delivers an inbound message containing arbitrary data from your protocol to the application.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input(_:_:_:_:_:)
func NWFramerDeliverInput(framer NWFramer, input_buffer []byte, input_length uintptr, message NWFramerMessage, is_complete bool) {
	if callErr := tryNWFramerDeliverInput(framer, input_buffer, input_length, message, is_complete); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerDeliverInputNoCopy func(framer NWFramer, input_length uintptr, message NWFramerMessage, is_complete bool) bool
var _nWFramerDeliverInputNoCopyErr error

func tryNWFramerDeliverInputNoCopy(framer NWFramer, input_length uintptr, message NWFramerMessage, is_complete bool) (bool, error) {
	if _nWFramerDeliverInputNoCopy == nil {
		return false, symbolCallError("nw_framer_deliver_input_no_copy", "10.15", _nWFramerDeliverInputNoCopyErr)
	}
	return _nWFramerDeliverInputNoCopy(framer, input_length, message, is_complete), nil
}

// NWFramerDeliverInputNoCopy delivers an inbound message containing a specific number of next received bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_deliver_input_no_copy(_:_:_:_:)
func NWFramerDeliverInputNoCopy(framer NWFramer, input_length uintptr, message NWFramerMessage, is_complete bool) bool {
	result, callErr := tryNWFramerDeliverInputNoCopy(framer, input_length, message, is_complete)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerMarkFailedWithError func(framer NWFramer, error_code int)
var _nWFramerMarkFailedWithErrorErr error

func tryNWFramerMarkFailedWithError(framer NWFramer, error_code int) error {
	if _nWFramerMarkFailedWithError == nil {
		return symbolCallError("nw_framer_mark_failed_with_error", "10.15", _nWFramerMarkFailedWithErrorErr)
	}
	_nWFramerMarkFailedWithError(framer, error_code)
	return nil
}

// NWFramerMarkFailedWithError indicates to a connection that your protocol has encountered an error, or has gracefully closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_failed_with_error(_:_:)
func NWFramerMarkFailedWithError(framer NWFramer, error_code int) {
	if callErr := tryNWFramerMarkFailedWithError(framer, error_code); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerMarkReady func(framer NWFramer)
var _nWFramerMarkReadyErr error

func tryNWFramerMarkReady(framer NWFramer) error {
	if _nWFramerMarkReady == nil {
		return symbolCallError("nw_framer_mark_ready", "10.15", _nWFramerMarkReadyErr)
	}
	_nWFramerMarkReady(framer)
	return nil
}

// NWFramerMarkReady indicates to a connection that your protocol’s handshake is complete.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_mark_ready(_:)
func NWFramerMarkReady(framer NWFramer) {
	if callErr := tryNWFramerMarkReady(framer); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerMessageAccessValue func(message NWFramerMessage, key string, access_value bool) bool
var _nWFramerMessageAccessValueErr error

func tryNWFramerMessageAccessValue(message NWFramerMessage, key string, access_value bool) (bool, error) {
	if _nWFramerMessageAccessValue == nil {
		return false, symbolCallError("nw_framer_message_access_value", "10.15", _nWFramerMessageAccessValueErr)
	}
	return _nWFramerMessageAccessValue(message, key, access_value), nil
}

// NWFramerMessageAccessValue accesses a custom value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_access_value(_:_:_:)
func NWFramerMessageAccessValue(message NWFramerMessage, key string, access_value bool) bool {
	result, callErr := tryNWFramerMessageAccessValue(message, key, access_value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerMessageCopyObjectValue func(message NWFramerMessage, key string) objectivec.Object
var _nWFramerMessageCopyObjectValueErr error

func tryNWFramerMessageCopyObjectValue(message NWFramerMessage, key string) (objectivec.Object, error) {
	if _nWFramerMessageCopyObjectValue == nil {
		return objectivec.Object{}, symbolCallError("nw_framer_message_copy_object_value", "10.15", _nWFramerMessageCopyObjectValueErr)
	}
	return _nWFramerMessageCopyObjectValue(message, key), nil
}

// NWFramerMessageCopyObjectValue accesses an NSObject value stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_copy_object_value(_:_:)
func NWFramerMessageCopyObjectValue(message NWFramerMessage, key string) objectivec.Object {
	result, callErr := tryNWFramerMessageCopyObjectValue(message, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerMessageCreate func(framer NWFramer) NWFramerMessage
var _nWFramerMessageCreateErr error

func tryNWFramerMessageCreate(framer NWFramer) (NWFramerMessage, error) {
	if _nWFramerMessageCreate == nil {
		return *new(NWFramerMessage), symbolCallError("nw_framer_message_create", "10.15", _nWFramerMessageCreateErr)
	}
	return _nWFramerMessageCreate(framer), nil
}

// NWFramerMessageCreate initializes an empty message from within a framer implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_create(_:)
func NWFramerMessageCreate(framer NWFramer) NWFramerMessage {
	result, callErr := tryNWFramerMessageCreate(framer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerMessageSetObjectValue func(message NWFramerMessage, key string, value objectivec.Object)
var _nWFramerMessageSetObjectValueErr error

func tryNWFramerMessageSetObjectValue(message NWFramerMessage, key string, value objectivec.Object) error {
	if _nWFramerMessageSetObjectValue == nil {
		return symbolCallError("nw_framer_message_set_object_value", "10.15", _nWFramerMessageSetObjectValueErr)
	}
	_nWFramerMessageSetObjectValue(message, key, value)
	return nil
}

// NWFramerMessageSetObjectValue sets an NSObject value to be stored in a framer message.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_object_value(_:_:_:)
func NWFramerMessageSetObjectValue(message NWFramerMessage, key string, value objectivec.Object) {
	if callErr := tryNWFramerMessageSetObjectValue(message, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerMessageSetValue func(message NWFramerMessage, key string, value unsafe.Pointer, dispose_value unsafe.Pointer)
var _nWFramerMessageSetValueErr error

func tryNWFramerMessageSetValue(message NWFramerMessage, key string, value unsafe.Pointer, dispose_value NWFramerMessageDisposeValue) error {
	if _nWFramerMessageSetValue == nil {
		return symbolCallError("nw_framer_message_set_value", "10.15", _nWFramerMessageSetValueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer) { dispose_value(blockArg0) })
	retainNetworkAsyncBlock(message.ID, "nw_framer_message_set_value:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerMessageSetValue(message, key, value, _block0)
	return nil
}

// NWFramerMessageSetValue sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_message_set_value(_:_:_:_:)
func NWFramerMessageSetValue(message NWFramerMessage, key string, value unsafe.Pointer, dispose_value NWFramerMessageDisposeValue) {
	if callErr := tryNWFramerMessageSetValue(message, key, value, dispose_value); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerOptionsCopyObjectValue func(options NWProtocolOptions, key string) objectivec.Object
var _nWFramerOptionsCopyObjectValueErr error

func tryNWFramerOptionsCopyObjectValue(options NWProtocolOptions, key string) (objectivec.Object, error) {
	if _nWFramerOptionsCopyObjectValue == nil {
		return objectivec.Object{}, symbolCallError("nw_framer_options_copy_object_value", "12.3", _nWFramerOptionsCopyObjectValueErr)
	}
	return _nWFramerOptionsCopyObjectValue(options, key), nil
}

// NWFramerOptionsCopyObjectValue.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_copy_object_value(_:_:)
func NWFramerOptionsCopyObjectValue(options NWProtocolOptions, key string) objectivec.Object {
	result, callErr := tryNWFramerOptionsCopyObjectValue(options, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerOptionsSetObjectValue func(options NWProtocolOptions, key string, value objectivec.Object)
var _nWFramerOptionsSetObjectValueErr error

func tryNWFramerOptionsSetObjectValue(options NWProtocolOptions, key string, value objectivec.Object) error {
	if _nWFramerOptionsSetObjectValue == nil {
		return symbolCallError("nw_framer_options_set_object_value", "12.3", _nWFramerOptionsSetObjectValueErr)
	}
	_nWFramerOptionsSetObjectValue(options, key, value)
	return nil
}

// NWFramerOptionsSetObjectValue.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_options_set_object_value(_:_:_:)
func NWFramerOptionsSetObjectValue(options NWProtocolOptions, key string, value objectivec.Object) {
	if callErr := tryNWFramerOptionsSetObjectValue(options, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerParseInput func(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *byte, parse unsafe.Pointer) bool
var _nWFramerParseInputErr error

func tryNWFramerParseInput(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer []byte, parse NWFramerParseCompletion) (bool, error) {
	if _nWFramerParseInput == nil {
		return false, symbolCallError("nw_framer_parse_input", "10.15", _nWFramerParseInputErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *uint8, blockArg1 uint32, blockArg2 bool) uint64 {
		return parse(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWFramerParseInput(framer, minimum_incomplete_length, maximum_length, unsafe.SliceData(temp_buffer), _block0), nil
}

// NWFramerParseInput examines the content of input data while inside your input handler block.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_input(_:_:_:_:_:)
func NWFramerParseInput(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer []byte, parse NWFramerParseCompletion) bool {
	result, callErr := tryNWFramerParseInput(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerParseOutput func(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer *byte, parse unsafe.Pointer) bool
var _nWFramerParseOutputErr error

func tryNWFramerParseOutput(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer []byte, parse NWFramerParseCompletion) (bool, error) {
	if _nWFramerParseOutput == nil {
		return false, symbolCallError("nw_framer_parse_output", "10.15", _nWFramerParseOutputErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *uint8, blockArg1 uint32, blockArg2 bool) uint64 {
		return parse(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWFramerParseOutput(framer, minimum_incomplete_length, maximum_length, unsafe.SliceData(temp_buffer), _block0), nil
}

// NWFramerParseOutput examines the content of output data while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_parse_output(_:_:_:_:_:)
func NWFramerParseOutput(framer NWFramer, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer []byte, parse NWFramerParseCompletion) bool {
	result, callErr := tryNWFramerParseOutput(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerPassThroughInput func(framer NWFramer)
var _nWFramerPassThroughInputErr error

func tryNWFramerPassThroughInput(framer NWFramer) error {
	if _nWFramerPassThroughInput == nil {
		return symbolCallError("nw_framer_pass_through_input", "10.15", _nWFramerPassThroughInputErr)
	}
	_nWFramerPassThroughInput(framer)
	return nil
}

// NWFramerPassThroughInput indicates that your protocol no longer needs to handle input data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_input(_:)
func NWFramerPassThroughInput(framer NWFramer) {
	if callErr := tryNWFramerPassThroughInput(framer); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerPassThroughOutput func(framer NWFramer)
var _nWFramerPassThroughOutputErr error

func tryNWFramerPassThroughOutput(framer NWFramer) error {
	if _nWFramerPassThroughOutput == nil {
		return symbolCallError("nw_framer_pass_through_output", "10.15", _nWFramerPassThroughOutputErr)
	}
	_nWFramerPassThroughOutput(framer)
	return nil
}

// NWFramerPassThroughOutput indicates that your protocol no longer needs to handle output data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_pass_through_output(_:)
func NWFramerPassThroughOutput(framer NWFramer) {
	if callErr := tryNWFramerPassThroughOutput(framer); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerPrependApplicationProtocol func(framer NWFramer, protocol_options NWProtocolOptions) bool
var _nWFramerPrependApplicationProtocolErr error

func tryNWFramerPrependApplicationProtocol(framer NWFramer, protocol_options NWProtocolOptions) (bool, error) {
	if _nWFramerPrependApplicationProtocol == nil {
		return false, symbolCallError("nw_framer_prepend_application_protocol", "10.15", _nWFramerPrependApplicationProtocolErr)
	}
	return _nWFramerPrependApplicationProtocol(framer, protocol_options), nil
}

// NWFramerPrependApplicationProtocol dynamically adds another protocol that will run above your protocol after your protocol calls nw_framer_mark_ready(_:).
//
// See: https://developer.apple.com/documentation/Network/nw_framer_prepend_application_protocol(_:_:)
func NWFramerPrependApplicationProtocol(framer NWFramer, protocol_options NWProtocolOptions) bool {
	result, callErr := tryNWFramerPrependApplicationProtocol(framer, protocol_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerProtocolCreateMessage func(definition NWProtocolDefinition) NWFramerMessage
var _nWFramerProtocolCreateMessageErr error

func tryNWFramerProtocolCreateMessage(definition NWProtocolDefinition) (NWFramerMessage, error) {
	if _nWFramerProtocolCreateMessage == nil {
		return *new(NWFramerMessage), symbolCallError("nw_framer_protocol_create_message", "10.15", _nWFramerProtocolCreateMessageErr)
	}
	return _nWFramerProtocolCreateMessage(definition), nil
}

// NWFramerProtocolCreateMessage initializes an empty message for a custom framer definition.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_protocol_create_message(_:)
func NWFramerProtocolCreateMessage(definition NWProtocolDefinition) NWFramerMessage {
	result, callErr := tryNWFramerProtocolCreateMessage(definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWFramerScheduleWakeup func(framer NWFramer, milliseconds uint64)
var _nWFramerScheduleWakeupErr error

func tryNWFramerScheduleWakeup(framer NWFramer, milliseconds uint64) error {
	if _nWFramerScheduleWakeup == nil {
		return symbolCallError("nw_framer_schedule_wakeup", "10.15", _nWFramerScheduleWakeupErr)
	}
	_nWFramerScheduleWakeup(framer, milliseconds)
	return nil
}

// NWFramerScheduleWakeup requests that the nw_framer_wakeup_handler_t be called on your protocol at a specific time in the future.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_schedule_wakeup(_:_:)
func NWFramerScheduleWakeup(framer NWFramer, milliseconds uint64) {
	if callErr := tryNWFramerScheduleWakeup(framer, milliseconds); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerSetCleanupHandler func(framer NWFramer, cleanup_handler unsafe.Pointer)
var _nWFramerSetCleanupHandlerErr error

func tryNWFramerSetCleanupHandler(framer NWFramer, cleanup_handler NWFramerCleanupHandler) error {
	if _nWFramerSetCleanupHandler == nil {
		return symbolCallError("nw_framer_set_cleanup_handler", "10.15", _nWFramerSetCleanupHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { cleanup_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_cleanup_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerSetCleanupHandler(framer, _block0)
	return nil
}

// NWFramerSetCleanupHandler sets a block to handle the final cleanup of allocations made by your protocol instance.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_cleanup_handler(_:_:)
func NWFramerSetCleanupHandler(framer NWFramer, cleanup_handler NWFramerCleanupHandler) {
	if callErr := tryNWFramerSetCleanupHandler(framer, cleanup_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerSetInputHandler func(framer NWFramer, input_handler unsafe.Pointer)
var _nWFramerSetInputHandlerErr error

func tryNWFramerSetInputHandler(framer NWFramer, input_handler NWFramerInputHandler) error {
	if _nWFramerSetInputHandler == nil {
		return symbolCallError("nw_framer_set_input_handler", "10.15", _nWFramerSetInputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) uint64 { return input_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_input_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerSetInputHandler(framer, _block0)
	return nil
}

// NWFramerSetInputHandler sets a block to handle new inbound data.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_input_handler(_:_:)
func NWFramerSetInputHandler(framer NWFramer, input_handler NWFramerInputHandler) {
	if callErr := tryNWFramerSetInputHandler(framer, input_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerSetOutputHandler func(framer NWFramer, output_handler unsafe.Pointer)
var _nWFramerSetOutputHandlerErr error

func tryNWFramerSetOutputHandler(framer NWFramer, output_handler NWFramerOutputHandler) error {
	if _nWFramerSetOutputHandler == nil {
		return symbolCallError("nw_framer_set_output_handler", "10.15", _nWFramerSetOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 objc.ID, blockArg2 uint32, blockArg3 bool) {
		output_handler(objectivec.ObjectFromID(blockArg0), objectivec.ObjectFromID(blockArg1), blockArg2, blockArg3)
	})
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_output_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerSetOutputHandler(framer, _block0)
	return nil
}

// NWFramerSetOutputHandler sets a block to handle new outbound messages.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_output_handler(_:_:)
func NWFramerSetOutputHandler(framer NWFramer, output_handler NWFramerOutputHandler) {
	if callErr := tryNWFramerSetOutputHandler(framer, output_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerSetStopHandler func(framer NWFramer, stop_handler unsafe.Pointer)
var _nWFramerSetStopHandlerErr error

func tryNWFramerSetStopHandler(framer NWFramer, stop_handler NWFramerStopHandler) error {
	if _nWFramerSetStopHandler == nil {
		return symbolCallError("nw_framer_set_stop_handler", "10.15", _nWFramerSetStopHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return stop_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_stop_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerSetStopHandler(framer, _block0)
	return nil
}

// NWFramerSetStopHandler sets a block to handle when the connection is being closed.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_stop_handler(_:_:)
func NWFramerSetStopHandler(framer NWFramer, stop_handler NWFramerStopHandler) {
	if callErr := tryNWFramerSetStopHandler(framer, stop_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerSetWakeupHandler func(framer NWFramer, wakeup_handler unsafe.Pointer)
var _nWFramerSetWakeupHandlerErr error

func tryNWFramerSetWakeupHandler(framer NWFramer, wakeup_handler NWFramerWakeupHandler) error {
	if _nWFramerSetWakeupHandler == nil {
		return symbolCallError("nw_framer_set_wakeup_handler", "10.15", _nWFramerSetWakeupHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { wakeup_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(framer.ID, "nw_framer_set_wakeup_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWFramerSetWakeupHandler(framer, _block0)
	return nil
}

// NWFramerSetWakeupHandler sets a handler to receive scheduled wakeup events.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_set_wakeup_handler(_:_:)
func NWFramerSetWakeupHandler(framer NWFramer, wakeup_handler NWFramerWakeupHandler) {
	if callErr := tryNWFramerSetWakeupHandler(framer, wakeup_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerWriteOutput func(framer NWFramer, output_buffer *byte, output_length uintptr)
var _nWFramerWriteOutputErr error

func tryNWFramerWriteOutput(framer NWFramer, output_buffer []byte, output_length uintptr) error {
	if _nWFramerWriteOutput == nil {
		return symbolCallError("nw_framer_write_output", "10.15", _nWFramerWriteOutputErr)
	}
	_nWFramerWriteOutput(framer, unsafe.SliceData(output_buffer), output_length)
	return nil
}

// NWFramerWriteOutput sends arbitrary output data in a buffer from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output(_:_:_:)
func NWFramerWriteOutput(framer NWFramer, output_buffer []byte, output_length uintptr) {
	if callErr := tryNWFramerWriteOutput(framer, output_buffer, output_length); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerWriteOutputData func(framer NWFramer, output_data uintptr)
var _nWFramerWriteOutputDataErr error

func tryNWFramerWriteOutputData(framer NWFramer, output_data dispatch.Data) error {
	if _nWFramerWriteOutputData == nil {
		return symbolCallError("nw_framer_write_output_data", "10.15", _nWFramerWriteOutputDataErr)
	}
	_nWFramerWriteOutputData(framer, uintptr(output_data.Handle()))
	return nil
}

// NWFramerWriteOutputData sends arbitrary output data from your protocol to the next protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_data(_:_:)
func NWFramerWriteOutputData(framer NWFramer, output_data dispatch.Data) {
	if callErr := tryNWFramerWriteOutputData(framer, output_data); callErr != nil {
		panic(callErr)
	}
}

var _nWFramerWriteOutputNoCopy func(framer NWFramer, output_length uintptr) bool
var _nWFramerWriteOutputNoCopyErr error

func tryNWFramerWriteOutputNoCopy(framer NWFramer, output_length uintptr) (bool, error) {
	if _nWFramerWriteOutputNoCopy == nil {
		return false, symbolCallError("nw_framer_write_output_no_copy", "10.15", _nWFramerWriteOutputNoCopyErr)
	}
	return _nWFramerWriteOutputNoCopy(framer, output_length), nil
}

// NWFramerWriteOutputNoCopy sends a specific number of bytes from a message while inside your output handler.
//
// See: https://developer.apple.com/documentation/Network/nw_framer_write_output_no_copy(_:_:)
func NWFramerWriteOutputNoCopy(framer NWFramer, output_length uintptr) bool {
	result, callErr := tryNWFramerWriteOutputNoCopy(framer, output_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWGroupDescriptorAddEndpoint func(descriptor NWGroupDescriptor, endpoint NWEndpoint) bool
var _nWGroupDescriptorAddEndpointErr error

func tryNWGroupDescriptorAddEndpoint(descriptor NWGroupDescriptor, endpoint NWEndpoint) (bool, error) {
	if _nWGroupDescriptorAddEndpoint == nil {
		return false, symbolCallError("nw_group_descriptor_add_endpoint", "11.0", _nWGroupDescriptorAddEndpointErr)
	}
	return _nWGroupDescriptorAddEndpoint(descriptor, endpoint), nil
}

// NWGroupDescriptorAddEndpoint adds a multicast address endpoint you specify to define an extra IP multicast group to join.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_add_endpoint(_:_:)
func NWGroupDescriptorAddEndpoint(descriptor NWGroupDescriptor, endpoint NWEndpoint) bool {
	result, callErr := tryNWGroupDescriptorAddEndpoint(descriptor, endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWGroupDescriptorCreateMulticast func(multicast_group NWEndpoint) NWGroupDescriptor
var _nWGroupDescriptorCreateMulticastErr error

func tryNWGroupDescriptorCreateMulticast(multicast_group NWEndpoint) (NWGroupDescriptor, error) {
	if _nWGroupDescriptorCreateMulticast == nil {
		return *new(NWGroupDescriptor), symbolCallError("nw_group_descriptor_create_multicast", "11.0", _nWGroupDescriptorCreateMulticastErr)
	}
	return _nWGroupDescriptorCreateMulticast(multicast_group), nil
}

// NWGroupDescriptorCreateMulticast creates group descriptor you use to join an IP multicast group on a local network.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multicast(_:)
func NWGroupDescriptorCreateMulticast(multicast_group NWEndpoint) NWGroupDescriptor {
	result, callErr := tryNWGroupDescriptorCreateMulticast(multicast_group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWGroupDescriptorCreateMultiplex func(remote_endpoint NWEndpoint) NWGroupDescriptor
var _nWGroupDescriptorCreateMultiplexErr error

func tryNWGroupDescriptorCreateMultiplex(remote_endpoint NWEndpoint) (NWGroupDescriptor, error) {
	if _nWGroupDescriptorCreateMultiplex == nil {
		return *new(NWGroupDescriptor), symbolCallError("nw_group_descriptor_create_multiplex", "12.0", _nWGroupDescriptorCreateMultiplexErr)
	}
	return _nWGroupDescriptorCreateMultiplex(remote_endpoint), nil
}

// NWGroupDescriptorCreateMultiplex.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multiplex(_:)
func NWGroupDescriptorCreateMultiplex(remote_endpoint NWEndpoint) NWGroupDescriptor {
	result, callErr := tryNWGroupDescriptorCreateMultiplex(remote_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWGroupDescriptorEnumerateEndpoints func(descriptor NWGroupDescriptor, enumerate_block unsafe.Pointer)
var _nWGroupDescriptorEnumerateEndpointsErr error

func tryNWGroupDescriptorEnumerateEndpoints(descriptor NWGroupDescriptor, enumerate_block NWGroupDescriptorEnumerateEndpointsBlock) error {
	if _nWGroupDescriptorEnumerateEndpoints == nil {
		return symbolCallError("nw_group_descriptor_enumerate_endpoints", "11.0", _nWGroupDescriptorEnumerateEndpointsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return enumerate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWGroupDescriptorEnumerateEndpoints(descriptor, _block0)
	return nil
}

// NWGroupDescriptorEnumerateEndpoints sets a handler to list all endpoints added to the group descriptor.
//
// See: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints(_:_:)
func NWGroupDescriptorEnumerateEndpoints(descriptor NWGroupDescriptor, enumerate_block NWGroupDescriptorEnumerateEndpointsBlock) {
	if callErr := tryNWGroupDescriptorEnumerateEndpoints(descriptor, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWInterfaceGetIndex func(interface_ NWInterface) uint32
var _nWInterfaceGetIndexErr error

func tryNWInterfaceGetIndex(interface_ NWInterface) (uint32, error) {
	if _nWInterfaceGetIndex == nil {
		return 0, symbolCallError("nw_interface_get_index", "10.14", _nWInterfaceGetIndexErr)
	}
	return _nWInterfaceGetIndex(interface_), nil
}

// NWInterfaceGetIndex accesses the system interface index associated with the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_index(_:)
func NWInterfaceGetIndex(interface_ NWInterface) uint32 {
	result, callErr := tryNWInterfaceGetIndex(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWInterfaceGetName func(interface_ NWInterface) *byte
var _nWInterfaceGetNameErr error

func tryNWInterfaceGetName(interface_ NWInterface) (*byte, error) {
	if _nWInterfaceGetName == nil {
		return nil, symbolCallError("nw_interface_get_name", "10.14", _nWInterfaceGetNameErr)
	}
	return _nWInterfaceGetName(interface_), nil
}

// NWInterfaceGetName accesses the name of the interface.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_name(_:)
func NWInterfaceGetName(interface_ NWInterface) *byte {
	result, callErr := tryNWInterfaceGetName(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWInterfaceGetType func(interface_ NWInterface) NWInterfaceType
var _nWInterfaceGetTypeErr error

func tryNWInterfaceGetType(interface_ NWInterface) (NWInterfaceType, error) {
	if _nWInterfaceGetType == nil {
		return *new(NWInterfaceType), symbolCallError("nw_interface_get_type", "10.14", _nWInterfaceGetTypeErr)
	}
	return _nWInterfaceGetType(interface_), nil
}

// NWInterfaceGetType accesses the type of the interface, such as Wi-Fi or Loopback.
//
// See: https://developer.apple.com/documentation/Network/nw_interface_get_type(_:)
func NWInterfaceGetType(interface_ NWInterface) NWInterfaceType {
	result, callErr := tryNWInterfaceGetType(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWIPCreateMetadata func() NWProtocolMetadata
var _nWIPCreateMetadataErr error

func tryNWIPCreateMetadata() (NWProtocolMetadata, error) {
	if _nWIPCreateMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_ip_create_metadata", "10.14", _nWIPCreateMetadataErr)
	}
	return _nWIPCreateMetadata(), nil
}

// NWIPCreateMetadata initializes an IP packet configuration with default settings.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_create_metadata()
func NWIPCreateMetadata() NWProtocolMetadata {
	result, callErr := tryNWIPCreateMetadata()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWIPMetadataGetEcnFlag func(metadata NWProtocolMetadata) NWIPEcnFlag
var _nWIPMetadataGetEcnFlagErr error

func tryNWIPMetadataGetEcnFlag(metadata NWProtocolMetadata) (NWIPEcnFlag, error) {
	if _nWIPMetadataGetEcnFlag == nil {
		return *new(NWIPEcnFlag), symbolCallError("nw_ip_metadata_get_ecn_flag", "10.14", _nWIPMetadataGetEcnFlagErr)
	}
	return _nWIPMetadataGetEcnFlag(metadata), nil
}

// NWIPMetadataGetEcnFlag checks the Explicit Congestion Notification flag value received on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_ecn_flag(_:)
func NWIPMetadataGetEcnFlag(metadata NWProtocolMetadata) NWIPEcnFlag {
	result, callErr := tryNWIPMetadataGetEcnFlag(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWIPMetadataGetReceiveTime func(metadata NWProtocolMetadata) uint64
var _nWIPMetadataGetReceiveTimeErr error

func tryNWIPMetadataGetReceiveTime(metadata NWProtocolMetadata) (uint64, error) {
	if _nWIPMetadataGetReceiveTime == nil {
		return 0, symbolCallError("nw_ip_metadata_get_receive_time", "10.14", _nWIPMetadataGetReceiveTimeErr)
	}
	return _nWIPMetadataGetReceiveTime(metadata), nil
}

// NWIPMetadataGetReceiveTime access the time at which a packet was received, in nanoseconds, based on `CLOCK_MONOTONIC_RAW`.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_receive_time(_:)
func NWIPMetadataGetReceiveTime(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWIPMetadataGetReceiveTime(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWIPMetadataGetServiceClass func(metadata NWProtocolMetadata) NWServiceClass
var _nWIPMetadataGetServiceClassErr error

func tryNWIPMetadataGetServiceClass(metadata NWProtocolMetadata) (NWServiceClass, error) {
	if _nWIPMetadataGetServiceClass == nil {
		return *new(NWServiceClass), symbolCallError("nw_ip_metadata_get_service_class", "10.14", _nWIPMetadataGetServiceClassErr)
	}
	return _nWIPMetadataGetServiceClass(metadata), nil
}

// NWIPMetadataGetServiceClass accesses a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_service_class(_:)
func NWIPMetadataGetServiceClass(metadata NWProtocolMetadata) NWServiceClass {
	result, callErr := tryNWIPMetadataGetServiceClass(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWIPMetadataSetEcnFlag func(metadata NWProtocolMetadata, ecn_flag NWIPEcnFlag)
var _nWIPMetadataSetEcnFlagErr error

func tryNWIPMetadataSetEcnFlag(metadata NWProtocolMetadata, ecn_flag NWIPEcnFlag) error {
	if _nWIPMetadataSetEcnFlag == nil {
		return symbolCallError("nw_ip_metadata_set_ecn_flag", "10.14", _nWIPMetadataSetEcnFlagErr)
	}
	_nWIPMetadataSetEcnFlag(metadata, ecn_flag)
	return nil
}

// NWIPMetadataSetEcnFlag sets a specific Explicit Congestion Notification flag value to set on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_ecn_flag(_:_:)
func NWIPMetadataSetEcnFlag(metadata NWProtocolMetadata, ecn_flag NWIPEcnFlag) {
	if callErr := tryNWIPMetadataSetEcnFlag(metadata, ecn_flag); callErr != nil {
		panic(callErr)
	}
}

var _nWIPMetadataSetServiceClass func(metadata NWProtocolMetadata, service_class NWServiceClass)
var _nWIPMetadataSetServiceClassErr error

func tryNWIPMetadataSetServiceClass(metadata NWProtocolMetadata, service_class NWServiceClass) error {
	if _nWIPMetadataSetServiceClass == nil {
		return symbolCallError("nw_ip_metadata_set_service_class", "10.14", _nWIPMetadataSetServiceClassErr)
	}
	_nWIPMetadataSetServiceClass(metadata, service_class)
	return nil
}

// NWIPMetadataSetServiceClass sets a specific service class to mark on an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_service_class(_:_:)
func NWIPMetadataSetServiceClass(metadata NWProtocolMetadata, service_class NWServiceClass) {
	if callErr := tryNWIPMetadataSetServiceClass(metadata, service_class); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetCalculateReceiveTime func(options NWProtocolOptions, calculate_receive_time bool)
var _nWIPOptionsSetCalculateReceiveTimeErr error

func tryNWIPOptionsSetCalculateReceiveTime(options NWProtocolOptions, calculate_receive_time bool) error {
	if _nWIPOptionsSetCalculateReceiveTime == nil {
		return symbolCallError("nw_ip_options_set_calculate_receive_time", "10.14", _nWIPOptionsSetCalculateReceiveTimeErr)
	}
	_nWIPOptionsSetCalculateReceiveTime(options, calculate_receive_time)
	return nil
}

// NWIPOptionsSetCalculateReceiveTime configures a connection to deliver receive timestamps for IP packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_calculate_receive_time(_:_:)
func NWIPOptionsSetCalculateReceiveTime(options NWProtocolOptions, calculate_receive_time bool) {
	if callErr := tryNWIPOptionsSetCalculateReceiveTime(options, calculate_receive_time); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetDisableFragmentation func(options NWProtocolOptions, disable_fragmentation bool)
var _nWIPOptionsSetDisableFragmentationErr error

func tryNWIPOptionsSetDisableFragmentation(options NWProtocolOptions, disable_fragmentation bool) error {
	if _nWIPOptionsSetDisableFragmentation == nil {
		return symbolCallError("nw_ip_options_set_disable_fragmentation", "10.14", _nWIPOptionsSetDisableFragmentationErr)
	}
	_nWIPOptionsSetDisableFragmentation(options, disable_fragmentation)
	return nil
}

// NWIPOptionsSetDisableFragmentation configures a connection to disable fragmentation on outbound packets.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_fragmentation(_:_:)
func NWIPOptionsSetDisableFragmentation(options NWProtocolOptions, disable_fragmentation bool) {
	if callErr := tryNWIPOptionsSetDisableFragmentation(options, disable_fragmentation); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetDisableMulticastLoopback func(options NWProtocolOptions, disable_multicast_loopback bool)
var _nWIPOptionsSetDisableMulticastLoopbackErr error

func tryNWIPOptionsSetDisableMulticastLoopback(options NWProtocolOptions, disable_multicast_loopback bool) error {
	if _nWIPOptionsSetDisableMulticastLoopback == nil {
		return symbolCallError("nw_ip_options_set_disable_multicast_loopback", "11.0", _nWIPOptionsSetDisableMulticastLoopbackErr)
	}
	_nWIPOptionsSetDisableMulticastLoopback(options, disable_multicast_loopback)
	return nil
}

// NWIPOptionsSetDisableMulticastLoopback.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_multicast_loopback(_:_:)
func NWIPOptionsSetDisableMulticastLoopback(options NWProtocolOptions, disable_multicast_loopback bool) {
	if callErr := tryNWIPOptionsSetDisableMulticastLoopback(options, disable_multicast_loopback); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetHopLimit func(options NWProtocolOptions, hop_limit uint8)
var _nWIPOptionsSetHopLimitErr error

func tryNWIPOptionsSetHopLimit(options NWProtocolOptions, hop_limit uint8) error {
	if _nWIPOptionsSetHopLimit == nil {
		return symbolCallError("nw_ip_options_set_hop_limit", "10.14", _nWIPOptionsSetHopLimitErr)
	}
	_nWIPOptionsSetHopLimit(options, hop_limit)
	return nil
}

// NWIPOptionsSetHopLimit configures the default hop limit for packets generated by a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_hop_limit(_:_:)
func NWIPOptionsSetHopLimit(options NWProtocolOptions, hop_limit uint8) {
	if callErr := tryNWIPOptionsSetHopLimit(options, hop_limit); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetLocalAddressPreference func(options NWProtocolOptions, preference NWIPLocalAddressPreference)
var _nWIPOptionsSetLocalAddressPreferenceErr error

func tryNWIPOptionsSetLocalAddressPreference(options NWProtocolOptions, preference NWIPLocalAddressPreference) error {
	if _nWIPOptionsSetLocalAddressPreference == nil {
		return symbolCallError("nw_ip_options_set_local_address_preference", "10.15", _nWIPOptionsSetLocalAddressPreferenceErr)
	}
	_nWIPOptionsSetLocalAddressPreference(options, preference)
	return nil
}

// NWIPOptionsSetLocalAddressPreference configures a connection to prefer certain types of local addresses, such as temporary or stable.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_local_address_preference(_:_:)
func NWIPOptionsSetLocalAddressPreference(options NWProtocolOptions, preference NWIPLocalAddressPreference) {
	if callErr := tryNWIPOptionsSetLocalAddressPreference(options, preference); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetUseMinimumMtu func(options NWProtocolOptions, use_minimum_mtu bool)
var _nWIPOptionsSetUseMinimumMtuErr error

func tryNWIPOptionsSetUseMinimumMtu(options NWProtocolOptions, use_minimum_mtu bool) error {
	if _nWIPOptionsSetUseMinimumMtu == nil {
		return symbolCallError("nw_ip_options_set_use_minimum_mtu", "10.14", _nWIPOptionsSetUseMinimumMtuErr)
	}
	_nWIPOptionsSetUseMinimumMtu(options, use_minimum_mtu)
	return nil
}

// NWIPOptionsSetUseMinimumMtu configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_use_minimum_mtu(_:_:)
func NWIPOptionsSetUseMinimumMtu(options NWProtocolOptions, use_minimum_mtu bool) {
	if callErr := tryNWIPOptionsSetUseMinimumMtu(options, use_minimum_mtu); callErr != nil {
		panic(callErr)
	}
}

var _nWIPOptionsSetVersion func(options NWProtocolOptions, version NWIPVersion)
var _nWIPOptionsSetVersionErr error

func tryNWIPOptionsSetVersion(options NWProtocolOptions, version NWIPVersion) error {
	if _nWIPOptionsSetVersion == nil {
		return symbolCallError("nw_ip_options_set_version", "10.14", _nWIPOptionsSetVersionErr)
	}
	_nWIPOptionsSetVersion(options, version)
	return nil
}

// NWIPOptionsSetVersion sets a required IP version to disable all other versions for a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ip_options_set_version(_:_:)
func NWIPOptionsSetVersion(options NWProtocolOptions, version NWIPVersion) {
	if callErr := tryNWIPOptionsSetVersion(options, version); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerCancel func(listener NWListener)
var _nWListenerCancelErr error

func tryNWListenerCancel(listener NWListener) error {
	if _nWListenerCancel == nil {
		return symbolCallError("nw_listener_cancel", "10.14", _nWListenerCancelErr)
	}
	_nWListenerCancel(listener)
	return nil
}

// NWListenerCancel stops listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_cancel(_:)
func NWListenerCancel(listener NWListener) {
	if callErr := tryNWListenerCancel(listener); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerCreate func(parameters NWParameters) NWListener
var _nWListenerCreateErr error

func tryNWListenerCreate(parameters NWParameters) (NWListener, error) {
	if _nWListenerCreate == nil {
		return *new(NWListener), symbolCallError("nw_listener_create", "10.14", _nWListenerCreateErr)
	}
	return _nWListenerCreate(parameters), nil
}

// NWListenerCreate initializes a network listener, which will select a random port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create(_:)
func NWListenerCreate(parameters NWParameters) NWListener {
	result, callErr := tryNWListenerCreate(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerCreateWithConnection func(connection NWConnection, parameters NWParameters) NWListener
var _nWListenerCreateWithConnectionErr error

func tryNWListenerCreateWithConnection(connection NWConnection, parameters NWParameters) (NWListener, error) {
	if _nWListenerCreateWithConnection == nil {
		return *new(NWListener), symbolCallError("nw_listener_create_with_connection", "10.14", _nWListenerCreateWithConnectionErr)
	}
	return _nWListenerCreateWithConnection(connection, parameters), nil
}

// NWListenerCreateWithConnection initializes a network listener to receive new streams on a multiplexed connection.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_connection(_:_:)
func NWListenerCreateWithConnection(connection NWConnection, parameters NWParameters) NWListener {
	result, callErr := tryNWListenerCreateWithConnection(connection, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerCreateWithLaunchdKey func(parameters NWParameters, launchd_key string) NWListener
var _nWListenerCreateWithLaunchdKeyErr error

func tryNWListenerCreateWithLaunchdKey(parameters NWParameters, launchd_key string) (NWListener, error) {
	if _nWListenerCreateWithLaunchdKey == nil {
		return *new(NWListener), symbolCallError("nw_listener_create_with_launchd_key", "10.14", _nWListenerCreateWithLaunchdKeyErr)
	}
	return _nWListenerCreateWithLaunchdKey(parameters, launchd_key), nil
}

// NWListenerCreateWithLaunchdKey.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_launchd_key(_:_:)
func NWListenerCreateWithLaunchdKey(parameters NWParameters, launchd_key string) NWListener {
	result, callErr := tryNWListenerCreateWithLaunchdKey(parameters, launchd_key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerCreateWithPort func(port string, parameters NWParameters) NWListener
var _nWListenerCreateWithPortErr error

func tryNWListenerCreateWithPort(port string, parameters NWParameters) (NWListener, error) {
	if _nWListenerCreateWithPort == nil {
		return *new(NWListener), symbolCallError("nw_listener_create_with_port", "10.14", _nWListenerCreateWithPortErr)
	}
	return _nWListenerCreateWithPort(port, parameters), nil
}

// NWListenerCreateWithPort initializes a network listener with a specified local port.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_create_with_port(_:_:)
func NWListenerCreateWithPort(port string, parameters NWParameters) NWListener {
	result, callErr := tryNWListenerCreateWithPort(port, parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerGetNewConnectionLimit func(listener NWListener) uint32
var _nWListenerGetNewConnectionLimitErr error

func tryNWListenerGetNewConnectionLimit(listener NWListener) (uint32, error) {
	if _nWListenerGetNewConnectionLimit == nil {
		return 0, symbolCallError("nw_listener_get_new_connection_limit", "10.15", _nWListenerGetNewConnectionLimitErr)
	}
	return _nWListenerGetNewConnectionLimit(listener), nil
}

// NWListenerGetNewConnectionLimit checks the remaining number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_new_connection_limit(_:)
func NWListenerGetNewConnectionLimit(listener NWListener) uint32 {
	result, callErr := tryNWListenerGetNewConnectionLimit(listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerGetPort func(listener NWListener) uint16
var _nWListenerGetPortErr error

func tryNWListenerGetPort(listener NWListener) (uint16, error) {
	if _nWListenerGetPort == nil {
		return 0, symbolCallError("nw_listener_get_port", "10.14", _nWListenerGetPortErr)
	}
	return _nWListenerGetPort(listener), nil
}

// NWListenerGetPort the port on which the listener can accept connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_get_port(_:)
func NWListenerGetPort(listener NWListener) uint16 {
	result, callErr := tryNWListenerGetPort(listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWListenerSetAdvertiseDescriptor func(listener NWListener, advertise_descriptor NWAdvertiseDescriptor)
var _nWListenerSetAdvertiseDescriptorErr error

func tryNWListenerSetAdvertiseDescriptor(listener NWListener, advertise_descriptor NWAdvertiseDescriptor) error {
	if _nWListenerSetAdvertiseDescriptor == nil {
		return symbolCallError("nw_listener_set_advertise_descriptor", "10.14", _nWListenerSetAdvertiseDescriptorErr)
	}
	_nWListenerSetAdvertiseDescriptor(listener, advertise_descriptor)
	return nil
}

// NWListenerSetAdvertiseDescriptor sets a Bonjour service that advertises the listener on the local network.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertise_descriptor(_:_:)
func NWListenerSetAdvertiseDescriptor(listener NWListener, advertise_descriptor NWAdvertiseDescriptor) {
	if callErr := tryNWListenerSetAdvertiseDescriptor(listener, advertise_descriptor); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetAdvertisedEndpointChangedHandler func(listener NWListener, handler unsafe.Pointer)
var _nWListenerSetAdvertisedEndpointChangedHandlerErr error

func tryNWListenerSetAdvertisedEndpointChangedHandler(listener NWListener, handler NWListenerAdvertisedEndpointChangedHandler) error {
	if _nWListenerSetAdvertisedEndpointChangedHandler == nil {
		return symbolCallError("nw_listener_set_advertised_endpoint_changed_handler", "10.14", _nWListenerSetAdvertisedEndpointChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID, blockArg1 bool) {
		handler(objectivec.ObjectFromID(blockArg0), blockArg1)
	})
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_advertised_endpoint_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWListenerSetAdvertisedEndpointChangedHandler(listener, _block0)
	return nil
}

// NWListenerSetAdvertisedEndpointChangedHandler sets a handler that receives updates for the service endpoint being advertised.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_advertised_endpoint_changed_handler(_:_:)
func NWListenerSetAdvertisedEndpointChangedHandler(listener NWListener, handler NWListenerAdvertisedEndpointChangedHandler) {
	if callErr := tryNWListenerSetAdvertisedEndpointChangedHandler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetNewConnectionGroupHandler func(listener NWListener, handler unsafe.Pointer)
var _nWListenerSetNewConnectionGroupHandlerErr error

func tryNWListenerSetNewConnectionGroupHandler(listener NWListener, handler NWListenerNewConnectionGroupHandler) error {
	if _nWListenerSetNewConnectionGroupHandler == nil {
		return symbolCallError("nw_listener_set_new_connection_group_handler", "12.0", _nWListenerSetNewConnectionGroupHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_new_connection_group_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWListenerSetNewConnectionGroupHandler(listener, _block0)
	return nil
}

// NWListenerSetNewConnectionGroupHandler.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_group_handler(_:_:)
func NWListenerSetNewConnectionGroupHandler(listener NWListener, handler NWListenerNewConnectionGroupHandler) {
	if callErr := tryNWListenerSetNewConnectionGroupHandler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetNewConnectionHandler func(listener NWListener, handler unsafe.Pointer)
var _nWListenerSetNewConnectionHandlerErr error

func tryNWListenerSetNewConnectionHandler(listener NWListener, handler NWListenerNewConnectionHandler) error {
	if _nWListenerSetNewConnectionHandler == nil {
		return symbolCallError("nw_listener_set_new_connection_handler", "10.14", _nWListenerSetNewConnectionHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_new_connection_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWListenerSetNewConnectionHandler(listener, _block0)
	return nil
}

// NWListenerSetNewConnectionHandler sets a handler that receives inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_handler(_:_:)
func NWListenerSetNewConnectionHandler(listener NWListener, handler NWListenerNewConnectionHandler) {
	if callErr := tryNWListenerSetNewConnectionHandler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetNewConnectionLimit func(listener NWListener, new_connection_limit uint32)
var _nWListenerSetNewConnectionLimitErr error

func tryNWListenerSetNewConnectionLimit(listener NWListener, new_connection_limit uint32) error {
	if _nWListenerSetNewConnectionLimit == nil {
		return symbolCallError("nw_listener_set_new_connection_limit", "10.15", _nWListenerSetNewConnectionLimitErr)
	}
	_nWListenerSetNewConnectionLimit(listener, new_connection_limit)
	return nil
}

// NWListenerSetNewConnectionLimit resets the number of inbound connections to deliver before rejecting connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_limit(_:_:)
func NWListenerSetNewConnectionLimit(listener NWListener, new_connection_limit uint32) {
	if callErr := tryNWListenerSetNewConnectionLimit(listener, new_connection_limit); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetQueue func(listener NWListener, queue uintptr)
var _nWListenerSetQueueErr error

func tryNWListenerSetQueue(listener NWListener, queue dispatch.Queue) error {
	if _nWListenerSetQueue == nil {
		return symbolCallError("nw_listener_set_queue", "10.14", _nWListenerSetQueueErr)
	}
	_nWListenerSetQueue(listener, uintptr(queue.Handle()))
	return nil
}

// NWListenerSetQueue sets the queue on which all listener events are delivered.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_queue(_:_:)
func NWListenerSetQueue(listener NWListener, queue dispatch.Queue) {
	if callErr := tryNWListenerSetQueue(listener, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerSetStateChangedHandler func(listener NWListener, handler unsafe.Pointer)
var _nWListenerSetStateChangedHandlerErr error

func tryNWListenerSetStateChangedHandler(listener NWListener, handler NWListenerStateChangedHandler) error {
	if _nWListenerSetStateChangedHandler == nil {
		return symbolCallError("nw_listener_set_state_changed_handler", "10.14", _nWListenerSetStateChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWListenerState, blockArg1 objc.ID) {
		handler(blockArg0, NWError{Object: objectivec.ObjectFromID(blockArg1)})
	})
	retainNetworkAsyncBlock(listener.ID, "nw_listener_set_state_changed_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWListenerSetStateChangedHandler(listener, _block0)
	return nil
}

// NWListenerSetStateChangedHandler sets a handler to receive listener state updates.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_set_state_changed_handler(_:_:)
func NWListenerSetStateChangedHandler(listener NWListener, handler NWListenerStateChangedHandler) {
	if callErr := tryNWListenerSetStateChangedHandler(listener, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWListenerStart func(listener NWListener)
var _nWListenerStartErr error

func tryNWListenerStart(listener NWListener) error {
	if _nWListenerStart == nil {
		return symbolCallError("nw_listener_start", "10.14", _nWListenerStartErr)
	}
	_nWListenerStart(listener)
	return nil
}

// NWListenerStart registers for listening for inbound connections.
//
// See: https://developer.apple.com/documentation/Network/nw_listener_start(_:)
func NWListenerStart(listener NWListener) {
	if callErr := tryNWListenerStart(listener); callErr != nil {
		panic(callErr)
	}
}

var _nWMulticastGroupDescriptorGetDisableUnicastTraffic func(multicast_descriptor NWGroupDescriptor) bool
var _nWMulticastGroupDescriptorGetDisableUnicastTrafficErr error

func tryNWMulticastGroupDescriptorGetDisableUnicastTraffic(multicast_descriptor NWGroupDescriptor) (bool, error) {
	if _nWMulticastGroupDescriptorGetDisableUnicastTraffic == nil {
		return false, symbolCallError("nw_multicast_group_descriptor_get_disable_unicast_traffic", "11.0", _nWMulticastGroupDescriptorGetDisableUnicastTrafficErr)
	}
	return _nWMulticastGroupDescriptorGetDisableUnicastTraffic(multicast_descriptor), nil
}

// NWMulticastGroupDescriptorGetDisableUnicastTraffic checks a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_get_disable_unicast_traffic(_:)
func NWMulticastGroupDescriptorGetDisableUnicastTraffic(multicast_descriptor NWGroupDescriptor) bool {
	result, callErr := tryNWMulticastGroupDescriptorGetDisableUnicastTraffic(multicast_descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWMulticastGroupDescriptorSetDisableUnicastTraffic func(multicast_descriptor NWGroupDescriptor, disable_unicast_traffic bool)
var _nWMulticastGroupDescriptorSetDisableUnicastTrafficErr error

func tryNWMulticastGroupDescriptorSetDisableUnicastTraffic(multicast_descriptor NWGroupDescriptor, disable_unicast_traffic bool) error {
	if _nWMulticastGroupDescriptorSetDisableUnicastTraffic == nil {
		return symbolCallError("nw_multicast_group_descriptor_set_disable_unicast_traffic", "11.0", _nWMulticastGroupDescriptorSetDisableUnicastTrafficErr)
	}
	_nWMulticastGroupDescriptorSetDisableUnicastTraffic(multicast_descriptor, disable_unicast_traffic)
	return nil
}

// NWMulticastGroupDescriptorSetDisableUnicastTraffic sets a Boolean that indicates whether a connection group should reject unicast traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_disable_unicast_traffic(_:_:)
func NWMulticastGroupDescriptorSetDisableUnicastTraffic(multicast_descriptor NWGroupDescriptor, disable_unicast_traffic bool) {
	if callErr := tryNWMulticastGroupDescriptorSetDisableUnicastTraffic(multicast_descriptor, disable_unicast_traffic); callErr != nil {
		panic(callErr)
	}
}

var _nWMulticastGroupDescriptorSetSpecificSource func(multicast_descriptor NWGroupDescriptor, source NWEndpoint)
var _nWMulticastGroupDescriptorSetSpecificSourceErr error

func tryNWMulticastGroupDescriptorSetSpecificSource(multicast_descriptor NWGroupDescriptor, source NWEndpoint) error {
	if _nWMulticastGroupDescriptorSetSpecificSource == nil {
		return symbolCallError("nw_multicast_group_descriptor_set_specific_source", "11.0", _nWMulticastGroupDescriptorSetSpecificSourceErr)
	}
	_nWMulticastGroupDescriptorSetSpecificSource(multicast_descriptor, source)
	return nil
}

// NWMulticastGroupDescriptorSetSpecificSource sets an optional address endpoint used to filter received multicast packets.
//
// See: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_specific_source(_:_:)
func NWMulticastGroupDescriptorSetSpecificSource(multicast_descriptor NWGroupDescriptor, source NWEndpoint) {
	if callErr := tryNWMulticastGroupDescriptorSetSpecificSource(multicast_descriptor, source); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersClearProhibitedInterfaceTypes func(parameters NWParameters)
var _nWParametersClearProhibitedInterfaceTypesErr error

func tryNWParametersClearProhibitedInterfaceTypes(parameters NWParameters) error {
	if _nWParametersClearProhibitedInterfaceTypes == nil {
		return symbolCallError("nw_parameters_clear_prohibited_interface_types", "10.14", _nWParametersClearProhibitedInterfaceTypesErr)
	}
	_nWParametersClearProhibitedInterfaceTypes(parameters)
	return nil
}

// NWParametersClearProhibitedInterfaceTypes removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interface_types(_:)
func NWParametersClearProhibitedInterfaceTypes(parameters NWParameters) {
	if callErr := tryNWParametersClearProhibitedInterfaceTypes(parameters); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersClearProhibitedInterfaces func(parameters NWParameters)
var _nWParametersClearProhibitedInterfacesErr error

func tryNWParametersClearProhibitedInterfaces(parameters NWParameters) error {
	if _nWParametersClearProhibitedInterfaces == nil {
		return symbolCallError("nw_parameters_clear_prohibited_interfaces", "10.14", _nWParametersClearProhibitedInterfacesErr)
	}
	_nWParametersClearProhibitedInterfaces(parameters)
	return nil
}

// NWParametersClearProhibitedInterfaces removes all prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interfaces(_:)
func NWParametersClearProhibitedInterfaces(parameters NWParameters) {
	if callErr := tryNWParametersClearProhibitedInterfaces(parameters); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersCopy func(parameters NWParameters) NWParameters
var _nWParametersCopyErr error

func tryNWParametersCopy(parameters NWParameters) (NWParameters, error) {
	if _nWParametersCopy == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_copy", "10.14", _nWParametersCopyErr)
	}
	return _nWParametersCopy(parameters), nil
}

// NWParametersCopy peforms a deep copy of a parameters object.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy(_:)
func NWParametersCopy(parameters NWParameters) NWParameters {
	result, callErr := tryNWParametersCopy(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCopyDefaultProtocolStack func(parameters NWParameters) NWProtocolStack
var _nWParametersCopyDefaultProtocolStackErr error

func tryNWParametersCopyDefaultProtocolStack(parameters NWParameters) (NWProtocolStack, error) {
	if _nWParametersCopyDefaultProtocolStack == nil {
		return *new(NWProtocolStack), symbolCallError("nw_parameters_copy_default_protocol_stack", "10.14", _nWParametersCopyDefaultProtocolStackErr)
	}
	return _nWParametersCopyDefaultProtocolStack(parameters), nil
}

// NWParametersCopyDefaultProtocolStack accesses the protocol stack used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_default_protocol_stack(_:)
func NWParametersCopyDefaultProtocolStack(parameters NWParameters) NWProtocolStack {
	result, callErr := tryNWParametersCopyDefaultProtocolStack(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCopyLocalEndpoint func(parameters NWParameters) NWEndpoint
var _nWParametersCopyLocalEndpointErr error

func tryNWParametersCopyLocalEndpoint(parameters NWParameters) (NWEndpoint, error) {
	if _nWParametersCopyLocalEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_parameters_copy_local_endpoint", "10.14", _nWParametersCopyLocalEndpointErr)
	}
	return _nWParametersCopyLocalEndpoint(parameters), nil
}

// NWParametersCopyLocalEndpoint accesses the local IP address and port used for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_local_endpoint(_:)
func NWParametersCopyLocalEndpoint(parameters NWParameters) NWEndpoint {
	result, callErr := tryNWParametersCopyLocalEndpoint(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCopyRequiredInterface func(parameters NWParameters) NWInterface
var _nWParametersCopyRequiredInterfaceErr error

func tryNWParametersCopyRequiredInterface(parameters NWParameters) (NWInterface, error) {
	if _nWParametersCopyRequiredInterface == nil {
		return *new(NWInterface), symbolCallError("nw_parameters_copy_required_interface", "10.14", _nWParametersCopyRequiredInterfaceErr)
	}
	return _nWParametersCopyRequiredInterface(parameters), nil
}

// NWParametersCopyRequiredInterface accesses the interface required on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_copy_required_interface(_:)
func NWParametersCopyRequiredInterface(parameters NWParameters) NWInterface {
	result, callErr := tryNWParametersCopyRequiredInterface(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreate func() NWParameters
var _nWParametersCreateErr error

func tryNWParametersCreate() (NWParameters, error) {
	if _nWParametersCreate == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create", "10.14", _nWParametersCreateErr)
	}
	return _nWParametersCreate(), nil
}

// NWParametersCreate initializes parameters for connections, listeners, and browsers with no protocols specified.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create()
func NWParametersCreate() NWParameters {
	result, callErr := tryNWParametersCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreateApplicationService func() NWParameters
var _nWParametersCreateApplicationServiceErr error

func tryNWParametersCreateApplicationService() (NWParameters, error) {
	if _nWParametersCreateApplicationService == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_application_service", "13.0", _nWParametersCreateApplicationServiceErr)
	}
	return _nWParametersCreateApplicationService(), nil
}

// NWParametersCreateApplicationService.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_application_service()
func NWParametersCreateApplicationService() NWParameters {
	result, callErr := tryNWParametersCreateApplicationService()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreateCustomIP func(custom_ip_protocol_number uint8, configure_ip unsafe.Pointer) NWParameters
var _nWParametersCreateCustomIPErr error

func tryNWParametersCreateCustomIP(custom_ip_protocol_number uint8, configure_ip NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if _nWParametersCreateCustomIP == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_custom_ip", "10.15", _nWParametersCreateCustomIPErr)
	}
	var _block0 unsafe.Pointer
	if configure_ip == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_ip(objectivec.ObjectFromID(blockArg0)) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	return _nWParametersCreateCustomIP(custom_ip_protocol_number, _block0), nil
}

// NWParametersCreateCustomIP initializes parameters for connections and listeners using a custom IP protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_custom_ip(_:_:)
func NWParametersCreateCustomIP(custom_ip_protocol_number uint8, configure_ip NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreateCustomIP(custom_ip_protocol_number, configure_ip)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreateQuic func(configure_quic unsafe.Pointer) NWParameters
var _nWParametersCreateQuicErr error

func tryNWParametersCreateQuic(configure_quic NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if _nWParametersCreateQuic == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_quic", "12.0", _nWParametersCreateQuicErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_quic(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWParametersCreateQuic(_block0), nil
}

// NWParametersCreateQuic initializes parameters for QUIC connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_quic(_:)
func NWParametersCreateQuic(configure_quic NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreateQuic(configure_quic)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreateSecureTCP func(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) NWParameters
var _nWParametersCreateSecureTCPErr error

func tryNWParametersCreateSecureTCP(configure_tls NWParametersConfigureProtocolBlock, configure_tcp NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if _nWParametersCreateSecureTCP == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_secure_tcp", "10.14", _nWParametersCreateSecureTCPErr)
	}
	var _block0 unsafe.Pointer
	if configure_tls == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_tls(objectivec.ObjectFromID(blockArg0)) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	var _block1 unsafe.Pointer
	if configure_tcp == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_tcp(objectivec.ObjectFromID(blockArg0)) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return _nWParametersCreateSecureTCP(_block0, _block1), nil
}

// NWParametersCreateSecureTCP initializes parameters for TLS or TCP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_tcp(_:_:)
func NWParametersCreateSecureTCP(configure_tls NWParametersConfigureProtocolBlock, configure_tcp NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreateSecureTCP(configure_tls, configure_tcp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersCreateSecureUDP func(configure_dtls unsafe.Pointer, configure_udp unsafe.Pointer) NWParameters
var _nWParametersCreateSecureUDPErr error

func tryNWParametersCreateSecureUDP(configure_dtls NWParametersConfigureProtocolBlock, configure_udp NWParametersConfigureProtocolBlock) (NWParameters, error) {
	if _nWParametersCreateSecureUDP == nil {
		return *new(NWParameters), symbolCallError("nw_parameters_create_secure_udp", "10.14", _nWParametersCreateSecureUDPErr)
	}
	var _block0 unsafe.Pointer
	if configure_dtls == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block0 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_dtls(objectivec.ObjectFromID(blockArg0)) })
		defer _block0Value.Release()
		_block0 = unsafe.Pointer(_block0Value)
	}
	var _block1 unsafe.Pointer
	if configure_udp == nil {
		if _nw_parameters_configure_protocol_default_configurationSymbol == 0 {
			return *new(NWParameters), symbolCallError("_nw_parameters_configure_protocol_default_configuration", "10.14", _nw_parameters_configure_protocol_default_configurationErr)
		}
		_block1 = networkProtocolBlockValue(_nw_parameters_configure_protocol_default_configurationSymbol)
	} else {
		_block1Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { configure_udp(objectivec.ObjectFromID(blockArg0)) })
		defer _block1Value.Release()
		_block1 = unsafe.Pointer(_block1Value)
	}
	return _nWParametersCreateSecureUDP(_block0, _block1), nil
}

// NWParametersCreateSecureUDP initializes parameters for DTLS or UDP connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_udp(_:_:)
func NWParametersCreateSecureUDP(configure_dtls NWParametersConfigureProtocolBlock, configure_udp NWParametersConfigureProtocolBlock) NWParameters {
	result, callErr := tryNWParametersCreateSecureUDP(configure_dtls, configure_udp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetAllowUltraConstrained func(parameters NWParameters) bool
var _nWParametersGetAllowUltraConstrainedErr error

func tryNWParametersGetAllowUltraConstrained(parameters NWParameters) (bool, error) {
	if _nWParametersGetAllowUltraConstrained == nil {
		return false, symbolCallError("nw_parameters_get_allow_ultra_constrained", "26.0", _nWParametersGetAllowUltraConstrainedErr)
	}
	return _nWParametersGetAllowUltraConstrained(parameters), nil
}

// NWParametersGetAllowUltraConstrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_allow_ultra_constrained(_:)
func NWParametersGetAllowUltraConstrained(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetAllowUltraConstrained(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetAttribution func(parameters NWParameters) NWParametersAttribution
var _nWParametersGetAttributionErr error

func tryNWParametersGetAttribution(parameters NWParameters) (NWParametersAttribution, error) {
	if _nWParametersGetAttribution == nil {
		return *new(NWParametersAttribution), symbolCallError("nw_parameters_get_attribution", "12.0", _nWParametersGetAttributionErr)
	}
	return _nWParametersGetAttribution(parameters), nil
}

// NWParametersGetAttribution gets a flag that indicates whether the network request originates from the developer or the user.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_attribution(_:)
func NWParametersGetAttribution(parameters NWParameters) NWParametersAttribution {
	result, callErr := tryNWParametersGetAttribution(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetExpiredDnsBehavior func(parameters NWParameters) NWParametersExpiredDnsBehavior
var _nWParametersGetExpiredDnsBehaviorErr error

func tryNWParametersGetExpiredDnsBehavior(parameters NWParameters) (NWParametersExpiredDnsBehavior, error) {
	if _nWParametersGetExpiredDnsBehavior == nil {
		return *new(NWParametersExpiredDnsBehavior), symbolCallError("nw_parameters_get_expired_dns_behavior", "10.14", _nWParametersGetExpiredDnsBehaviorErr)
	}
	return _nWParametersGetExpiredDnsBehavior(parameters), nil
}

// NWParametersGetExpiredDnsBehavior checks the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_expired_dns_behavior(_:)
func NWParametersGetExpiredDnsBehavior(parameters NWParameters) NWParametersExpiredDnsBehavior {
	result, callErr := tryNWParametersGetExpiredDnsBehavior(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetFastOpenEnabled func(parameters NWParameters) bool
var _nWParametersGetFastOpenEnabledErr error

func tryNWParametersGetFastOpenEnabled(parameters NWParameters) (bool, error) {
	if _nWParametersGetFastOpenEnabled == nil {
		return false, symbolCallError("nw_parameters_get_fast_open_enabled", "10.14", _nWParametersGetFastOpenEnabledErr)
	}
	return _nWParametersGetFastOpenEnabled(parameters), nil
}

// NWParametersGetFastOpenEnabled checks if sending application data with protocol handshakes is enabled.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_fast_open_enabled(_:)
func NWParametersGetFastOpenEnabled(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetFastOpenEnabled(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetIncludePeerToPeer func(parameters NWParameters) bool
var _nWParametersGetIncludePeerToPeerErr error

func tryNWParametersGetIncludePeerToPeer(parameters NWParameters) (bool, error) {
	if _nWParametersGetIncludePeerToPeer == nil {
		return false, symbolCallError("nw_parameters_get_include_peer_to_peer", "10.14", _nWParametersGetIncludePeerToPeerErr)
	}
	return _nWParametersGetIncludePeerToPeer(parameters), nil
}

// NWParametersGetIncludePeerToPeer checks whether a connection is allowed to use peer-to-peer link technologies.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_include_peer_to_peer(_:)
func NWParametersGetIncludePeerToPeer(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetIncludePeerToPeer(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetLocalOnly func(parameters NWParameters) bool
var _nWParametersGetLocalOnlyErr error

func tryNWParametersGetLocalOnly(parameters NWParameters) (bool, error) {
	if _nWParametersGetLocalOnly == nil {
		return false, symbolCallError("nw_parameters_get_local_only", "10.14", _nWParametersGetLocalOnlyErr)
	}
	return _nWParametersGetLocalOnly(parameters), nil
}

// NWParametersGetLocalOnly checks if a listener is restricted to accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_local_only(_:)
func NWParametersGetLocalOnly(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetLocalOnly(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetMultipathService func(parameters NWParameters) NWMultipathService
var _nWParametersGetMultipathServiceErr error

func tryNWParametersGetMultipathService(parameters NWParameters) (NWMultipathService, error) {
	if _nWParametersGetMultipathService == nil {
		return *new(NWMultipathService), symbolCallError("nw_parameters_get_multipath_service", "10.14", _nWParametersGetMultipathServiceErr)
	}
	return _nWParametersGetMultipathService(parameters), nil
}

// NWParametersGetMultipathService checks if multipath is enabled on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_multipath_service(_:)
func NWParametersGetMultipathService(parameters NWParameters) NWMultipathService {
	result, callErr := tryNWParametersGetMultipathService(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetPreferNoProxy func(parameters NWParameters) bool
var _nWParametersGetPreferNoProxyErr error

func tryNWParametersGetPreferNoProxy(parameters NWParameters) (bool, error) {
	if _nWParametersGetPreferNoProxy == nil {
		return false, symbolCallError("nw_parameters_get_prefer_no_proxy", "10.14", _nWParametersGetPreferNoProxyErr)
	}
	return _nWParametersGetPreferNoProxy(parameters), nil
}

// NWParametersGetPreferNoProxy checks if proxies are ignored by default.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prefer_no_proxy(_:)
func NWParametersGetPreferNoProxy(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetPreferNoProxy(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetProhibitConstrained func(parameters NWParameters) bool
var _nWParametersGetProhibitConstrainedErr error

func tryNWParametersGetProhibitConstrained(parameters NWParameters) (bool, error) {
	if _nWParametersGetProhibitConstrained == nil {
		return false, symbolCallError("nw_parameters_get_prohibit_constrained", "10.15", _nWParametersGetProhibitConstrainedErr)
	}
	return _nWParametersGetProhibitConstrained(parameters), nil
}

// NWParametersGetProhibitConstrained checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_constrained(_:)
func NWParametersGetProhibitConstrained(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetProhibitConstrained(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetProhibitExpensive func(parameters NWParameters) bool
var _nWParametersGetProhibitExpensiveErr error

func tryNWParametersGetProhibitExpensive(parameters NWParameters) (bool, error) {
	if _nWParametersGetProhibitExpensive == nil {
		return false, symbolCallError("nw_parameters_get_prohibit_expensive", "10.14", _nWParametersGetProhibitExpensiveErr)
	}
	return _nWParametersGetProhibitExpensive(parameters), nil
}

// NWParametersGetProhibitExpensive checks if connections, listeners, and browsers are prevented from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_expensive(_:)
func NWParametersGetProhibitExpensive(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetProhibitExpensive(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetRequiredInterfaceType func(parameters NWParameters) NWInterfaceType
var _nWParametersGetRequiredInterfaceTypeErr error

func tryNWParametersGetRequiredInterfaceType(parameters NWParameters) (NWInterfaceType, error) {
	if _nWParametersGetRequiredInterfaceType == nil {
		return *new(NWInterfaceType), symbolCallError("nw_parameters_get_required_interface_type", "10.14", _nWParametersGetRequiredInterfaceTypeErr)
	}
	return _nWParametersGetRequiredInterfaceType(parameters), nil
}

// NWParametersGetRequiredInterfaceType accesses the interface type required on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_required_interface_type(_:)
func NWParametersGetRequiredInterfaceType(parameters NWParameters) NWInterfaceType {
	result, callErr := tryNWParametersGetRequiredInterfaceType(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetReuseLocalAddress func(parameters NWParameters) bool
var _nWParametersGetReuseLocalAddressErr error

func tryNWParametersGetReuseLocalAddress(parameters NWParameters) (bool, error) {
	if _nWParametersGetReuseLocalAddress == nil {
		return false, symbolCallError("nw_parameters_get_reuse_local_address", "10.14", _nWParametersGetReuseLocalAddressErr)
	}
	return _nWParametersGetReuseLocalAddress(parameters), nil
}

// NWParametersGetReuseLocalAddress checks whether a connection allows reusing local addresses and ports.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_reuse_local_address(_:)
func NWParametersGetReuseLocalAddress(parameters NWParameters) bool {
	result, callErr := tryNWParametersGetReuseLocalAddress(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersGetServiceClass func(parameters NWParameters) NWServiceClass
var _nWParametersGetServiceClassErr error

func tryNWParametersGetServiceClass(parameters NWParameters) (NWServiceClass, error) {
	if _nWParametersGetServiceClass == nil {
		return *new(NWServiceClass), symbolCallError("nw_parameters_get_service_class", "10.14", _nWParametersGetServiceClassErr)
	}
	return _nWParametersGetServiceClass(parameters), nil
}

// NWParametersGetServiceClass checks the level of service quality used for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_get_service_class(_:)
func NWParametersGetServiceClass(parameters NWParameters) NWServiceClass {
	result, callErr := tryNWParametersGetServiceClass(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersIterateProhibitedInterfaceTypes func(parameters NWParameters, iterate_block unsafe.Pointer)
var _nWParametersIterateProhibitedInterfaceTypesErr error

func tryNWParametersIterateProhibitedInterfaceTypes(parameters NWParameters, iterate_block NWParametersIterateInterfaceTypesBlock) error {
	if _nWParametersIterateProhibitedInterfaceTypes == nil {
		return symbolCallError("nw_parameters_iterate_prohibited_interface_types", "10.14", _nWParametersIterateProhibitedInterfaceTypesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 NWInterfaceType) bool { return iterate_block(blockArg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWParametersIterateProhibitedInterfaceTypes(parameters, _block0)
	return nil
}

// NWParametersIterateProhibitedInterfaceTypes examines the list of prohibited interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interface_types(_:_:)
func NWParametersIterateProhibitedInterfaceTypes(parameters NWParameters, iterate_block NWParametersIterateInterfaceTypesBlock) {
	if callErr := tryNWParametersIterateProhibitedInterfaceTypes(parameters, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersIterateProhibitedInterfaces func(parameters NWParameters, iterate_block unsafe.Pointer)
var _nWParametersIterateProhibitedInterfacesErr error

func tryNWParametersIterateProhibitedInterfaces(parameters NWParameters, iterate_block NWParametersIterateInterfacesBlock) error {
	if _nWParametersIterateProhibitedInterfaces == nil {
		return symbolCallError("nw_parameters_iterate_prohibited_interfaces", "10.14", _nWParametersIterateProhibitedInterfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return iterate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWParametersIterateProhibitedInterfaces(parameters, _block0)
	return nil
}

// NWParametersIterateProhibitedInterfaces examines the list of prohibited interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interfaces(_:_:)
func NWParametersIterateProhibitedInterfaces(parameters NWParameters, iterate_block NWParametersIterateInterfacesBlock) {
	if callErr := tryNWParametersIterateProhibitedInterfaces(parameters, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersProhibitInterface func(parameters NWParameters, interface_ NWInterface)
var _nWParametersProhibitInterfaceErr error

func tryNWParametersProhibitInterface(parameters NWParameters, interface_ NWInterface) error {
	if _nWParametersProhibitInterface == nil {
		return symbolCallError("nw_parameters_prohibit_interface", "10.14", _nWParametersProhibitInterfaceErr)
	}
	_nWParametersProhibitInterface(parameters, interface_)
	return nil
}

// NWParametersProhibitInterface prevents connections and listeners from using a specific interface.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface(_:_:)
func NWParametersProhibitInterface(parameters NWParameters, interface_ NWInterface) {
	if callErr := tryNWParametersProhibitInterface(parameters, interface_); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersProhibitInterfaceType func(parameters NWParameters, interface_type NWInterfaceType)
var _nWParametersProhibitInterfaceTypeErr error

func tryNWParametersProhibitInterfaceType(parameters NWParameters, interface_type NWInterfaceType) error {
	if _nWParametersProhibitInterfaceType == nil {
		return symbolCallError("nw_parameters_prohibit_interface_type", "10.14", _nWParametersProhibitInterfaceTypeErr)
	}
	_nWParametersProhibitInterfaceType(parameters, interface_type)
	return nil
}

// NWParametersProhibitInterfaceType prevents connections, listeners, and browsers from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface_type(_:_:)
func NWParametersProhibitInterfaceType(parameters NWParameters, interface_type NWInterfaceType) {
	if callErr := tryNWParametersProhibitInterfaceType(parameters, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersRequireInterface func(parameters NWParameters, interface_ NWInterface)
var _nWParametersRequireInterfaceErr error

func tryNWParametersRequireInterface(parameters NWParameters, interface_ NWInterface) error {
	if _nWParametersRequireInterface == nil {
		return symbolCallError("nw_parameters_require_interface", "10.14", _nWParametersRequireInterfaceErr)
	}
	_nWParametersRequireInterface(parameters, interface_)
	return nil
}

// NWParametersRequireInterface sets a specific interface to require on connections, listeners, and browsers.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_require_interface(_:_:)
func NWParametersRequireInterface(parameters NWParameters, interface_ NWInterface) {
	if callErr := tryNWParametersRequireInterface(parameters, interface_); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersRequiresDnssecValidation func(parameters NWParameters) bool
var _nWParametersRequiresDnssecValidationErr error

func tryNWParametersRequiresDnssecValidation(parameters NWParameters) (bool, error) {
	if _nWParametersRequiresDnssecValidation == nil {
		return false, symbolCallError("nw_parameters_requires_dnssec_validation", "13.0", _nWParametersRequiresDnssecValidationErr)
	}
	return _nWParametersRequiresDnssecValidation(parameters), nil
}

// NWParametersRequiresDnssecValidation checks whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_requires_dnssec_validation(_:)
func NWParametersRequiresDnssecValidation(parameters NWParameters) bool {
	result, callErr := tryNWParametersRequiresDnssecValidation(parameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWParametersSetAllowUltraConstrained func(parameters NWParameters, allow_ultra_constrained bool)
var _nWParametersSetAllowUltraConstrainedErr error

func tryNWParametersSetAllowUltraConstrained(parameters NWParameters, allow_ultra_constrained bool) error {
	if _nWParametersSetAllowUltraConstrained == nil {
		return symbolCallError("nw_parameters_set_allow_ultra_constrained", "26.0", _nWParametersSetAllowUltraConstrainedErr)
	}
	_nWParametersSetAllowUltraConstrained(parameters, allow_ultra_constrained)
	return nil
}

// NWParametersSetAllowUltraConstrained.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_allow_ultra_constrained(_:_:)
func NWParametersSetAllowUltraConstrained(parameters NWParameters, allow_ultra_constrained bool) {
	if callErr := tryNWParametersSetAllowUltraConstrained(parameters, allow_ultra_constrained); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetAttribution func(parameters NWParameters, attribution NWParametersAttribution)
var _nWParametersSetAttributionErr error

func tryNWParametersSetAttribution(parameters NWParameters, attribution NWParametersAttribution) error {
	if _nWParametersSetAttribution == nil {
		return symbolCallError("nw_parameters_set_attribution", "12.0", _nWParametersSetAttributionErr)
	}
	_nWParametersSetAttribution(parameters, attribution)
	return nil
}

// NWParametersSetAttribution sets a flag that indicates whether the network request originates from the developer or the user.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_attribution(_:_:)
func NWParametersSetAttribution(parameters NWParameters, attribution NWParametersAttribution) {
	if callErr := tryNWParametersSetAttribution(parameters, attribution); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetExpiredDnsBehavior func(parameters NWParameters, expired_dns_behavior NWParametersExpiredDnsBehavior)
var _nWParametersSetExpiredDnsBehaviorErr error

func tryNWParametersSetExpiredDnsBehavior(parameters NWParameters, expired_dns_behavior NWParametersExpiredDnsBehavior) error {
	if _nWParametersSetExpiredDnsBehavior == nil {
		return symbolCallError("nw_parameters_set_expired_dns_behavior", "10.14", _nWParametersSetExpiredDnsBehaviorErr)
	}
	_nWParametersSetExpiredDnsBehavior(parameters, expired_dns_behavior)
	return nil
}

// NWParametersSetExpiredDnsBehavior sets the behavior for how expired DNS answers should be used.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_expired_dns_behavior(_:_:)
func NWParametersSetExpiredDnsBehavior(parameters NWParameters, expired_dns_behavior NWParametersExpiredDnsBehavior) {
	if callErr := tryNWParametersSetExpiredDnsBehavior(parameters, expired_dns_behavior); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetFastOpenEnabled func(parameters NWParameters, fast_open_enabled bool)
var _nWParametersSetFastOpenEnabledErr error

func tryNWParametersSetFastOpenEnabled(parameters NWParameters, fast_open_enabled bool) error {
	if _nWParametersSetFastOpenEnabled == nil {
		return symbolCallError("nw_parameters_set_fast_open_enabled", "10.14", _nWParametersSetFastOpenEnabledErr)
	}
	_nWParametersSetFastOpenEnabled(parameters, fast_open_enabled)
	return nil
}

// NWParametersSetFastOpenEnabled enables sending application data with protocol handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_fast_open_enabled(_:_:)
func NWParametersSetFastOpenEnabled(parameters NWParameters, fast_open_enabled bool) {
	if callErr := tryNWParametersSetFastOpenEnabled(parameters, fast_open_enabled); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetIncludePeerToPeer func(parameters NWParameters, include_peer_to_peer bool)
var _nWParametersSetIncludePeerToPeerErr error

func tryNWParametersSetIncludePeerToPeer(parameters NWParameters, include_peer_to_peer bool) error {
	if _nWParametersSetIncludePeerToPeer == nil {
		return symbolCallError("nw_parameters_set_include_peer_to_peer", "10.14", _nWParametersSetIncludePeerToPeerErr)
	}
	_nWParametersSetIncludePeerToPeer(parameters, include_peer_to_peer)
	return nil
}

// NWParametersSetIncludePeerToPeer enables peer-to-peer link technologies for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_include_peer_to_peer(_:_:)
func NWParametersSetIncludePeerToPeer(parameters NWParameters, include_peer_to_peer bool) {
	if callErr := tryNWParametersSetIncludePeerToPeer(parameters, include_peer_to_peer); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetLocalEndpoint func(parameters NWParameters, local_endpoint NWEndpoint)
var _nWParametersSetLocalEndpointErr error

func tryNWParametersSetLocalEndpoint(parameters NWParameters, local_endpoint NWEndpoint) error {
	if _nWParametersSetLocalEndpoint == nil {
		return symbolCallError("nw_parameters_set_local_endpoint", "10.14", _nWParametersSetLocalEndpointErr)
	}
	_nWParametersSetLocalEndpoint(parameters, local_endpoint)
	return nil
}

// NWParametersSetLocalEndpoint sets a specific local IP address and port to use for connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_endpoint(_:_:)
func NWParametersSetLocalEndpoint(parameters NWParameters, local_endpoint NWEndpoint) {
	if callErr := tryNWParametersSetLocalEndpoint(parameters, local_endpoint); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetLocalOnly func(parameters NWParameters, local_only bool)
var _nWParametersSetLocalOnlyErr error

func tryNWParametersSetLocalOnly(parameters NWParameters, local_only bool) error {
	if _nWParametersSetLocalOnly == nil {
		return symbolCallError("nw_parameters_set_local_only", "10.14", _nWParametersSetLocalOnlyErr)
	}
	_nWParametersSetLocalOnly(parameters, local_only)
	return nil
}

// NWParametersSetLocalOnly restricts listeners to only accepting connections from the local link.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_local_only(_:_:)
func NWParametersSetLocalOnly(parameters NWParameters, local_only bool) {
	if callErr := tryNWParametersSetLocalOnly(parameters, local_only); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetMultipathService func(parameters NWParameters, multipath_service NWMultipathService)
var _nWParametersSetMultipathServiceErr error

func tryNWParametersSetMultipathService(parameters NWParameters, multipath_service NWMultipathService) error {
	if _nWParametersSetMultipathService == nil {
		return symbolCallError("nw_parameters_set_multipath_service", "10.14", _nWParametersSetMultipathServiceErr)
	}
	_nWParametersSetMultipathService(parameters, multipath_service)
	return nil
}

// NWParametersSetMultipathService enables multipath protocols to allow connections to use multiple interfaces.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_multipath_service(_:_:)
func NWParametersSetMultipathService(parameters NWParameters, multipath_service NWMultipathService) {
	if callErr := tryNWParametersSetMultipathService(parameters, multipath_service); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetPreferNoProxy func(parameters NWParameters, prefer_no_proxy bool)
var _nWParametersSetPreferNoProxyErr error

func tryNWParametersSetPreferNoProxy(parameters NWParameters, prefer_no_proxy bool) error {
	if _nWParametersSetPreferNoProxy == nil {
		return symbolCallError("nw_parameters_set_prefer_no_proxy", "10.14", _nWParametersSetPreferNoProxyErr)
	}
	_nWParametersSetPreferNoProxy(parameters, prefer_no_proxy)
	return nil
}

// NWParametersSetPreferNoProxy sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prefer_no_proxy(_:_:)
func NWParametersSetPreferNoProxy(parameters NWParameters, prefer_no_proxy bool) {
	if callErr := tryNWParametersSetPreferNoProxy(parameters, prefer_no_proxy); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetPrivacyContext func(parameters NWParameters, privacy_context NWPrivacyContext)
var _nWParametersSetPrivacyContextErr error

func tryNWParametersSetPrivacyContext(parameters NWParameters, privacy_context NWPrivacyContext) error {
	if _nWParametersSetPrivacyContext == nil {
		return symbolCallError("nw_parameters_set_privacy_context", "11.0", _nWParametersSetPrivacyContextErr)
	}
	_nWParametersSetPrivacyContext(parameters, privacy_context)
	return nil
}

// NWParametersSetPrivacyContext associates a privacy context with any connections or listeners that use the parameters.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_privacy_context(_:_:)
func NWParametersSetPrivacyContext(parameters NWParameters, privacy_context NWPrivacyContext) {
	if callErr := tryNWParametersSetPrivacyContext(parameters, privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetProhibitConstrained func(parameters NWParameters, prohibit_constrained bool)
var _nWParametersSetProhibitConstrainedErr error

func tryNWParametersSetProhibitConstrained(parameters NWParameters, prohibit_constrained bool) error {
	if _nWParametersSetProhibitConstrained == nil {
		return symbolCallError("nw_parameters_set_prohibit_constrained", "10.15", _nWParametersSetProhibitConstrainedErr)
	}
	_nWParametersSetProhibitConstrained(parameters, prohibit_constrained)
	return nil
}

// NWParametersSetProhibitConstrained prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_constrained(_:_:)
func NWParametersSetProhibitConstrained(parameters NWParameters, prohibit_constrained bool) {
	if callErr := tryNWParametersSetProhibitConstrained(parameters, prohibit_constrained); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetProhibitExpensive func(parameters NWParameters, prohibit_expensive bool)
var _nWParametersSetProhibitExpensiveErr error

func tryNWParametersSetProhibitExpensive(parameters NWParameters, prohibit_expensive bool) error {
	if _nWParametersSetProhibitExpensive == nil {
		return symbolCallError("nw_parameters_set_prohibit_expensive", "10.14", _nWParametersSetProhibitExpensiveErr)
	}
	_nWParametersSetProhibitExpensive(parameters, prohibit_expensive)
	return nil
}

// NWParametersSetProhibitExpensive prevents connections, listeners, and browsers from using network paths marked as expensive.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_expensive(_:_:)
func NWParametersSetProhibitExpensive(parameters NWParameters, prohibit_expensive bool) {
	if callErr := tryNWParametersSetProhibitExpensive(parameters, prohibit_expensive); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetRequiredInterfaceType func(parameters NWParameters, interface_type NWInterfaceType)
var _nWParametersSetRequiredInterfaceTypeErr error

func tryNWParametersSetRequiredInterfaceType(parameters NWParameters, interface_type NWInterfaceType) error {
	if _nWParametersSetRequiredInterfaceType == nil {
		return symbolCallError("nw_parameters_set_required_interface_type", "10.14", _nWParametersSetRequiredInterfaceTypeErr)
	}
	_nWParametersSetRequiredInterfaceType(parameters, interface_type)
	return nil
}

// NWParametersSetRequiredInterfaceType sets an interface type to require on connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_required_interface_type(_:_:)
func NWParametersSetRequiredInterfaceType(parameters NWParameters, interface_type NWInterfaceType) {
	if callErr := tryNWParametersSetRequiredInterfaceType(parameters, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetRequiresDnssecValidation func(parameters NWParameters, requires_dnssec_validation bool)
var _nWParametersSetRequiresDnssecValidationErr error

func tryNWParametersSetRequiresDnssecValidation(parameters NWParameters, requires_dnssec_validation bool) error {
	if _nWParametersSetRequiresDnssecValidation == nil {
		return symbolCallError("nw_parameters_set_requires_dnssec_validation", "13.0", _nWParametersSetRequiresDnssecValidationErr)
	}
	_nWParametersSetRequiresDnssecValidation(parameters, requires_dnssec_validation)
	return nil
}

// NWParametersSetRequiresDnssecValidation determines whether a connection requires DNSSEC validation when resolving endpoints.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_requires_dnssec_validation(_:_:)
func NWParametersSetRequiresDnssecValidation(parameters NWParameters, requires_dnssec_validation bool) {
	if callErr := tryNWParametersSetRequiresDnssecValidation(parameters, requires_dnssec_validation); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetReuseLocalAddress func(parameters NWParameters, reuse_local_address bool)
var _nWParametersSetReuseLocalAddressErr error

func tryNWParametersSetReuseLocalAddress(parameters NWParameters, reuse_local_address bool) error {
	if _nWParametersSetReuseLocalAddress == nil {
		return symbolCallError("nw_parameters_set_reuse_local_address", "10.14", _nWParametersSetReuseLocalAddressErr)
	}
	_nWParametersSetReuseLocalAddress(parameters, reuse_local_address)
	return nil
}

// NWParametersSetReuseLocalAddress allows reusing local addresses and ports across connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_reuse_local_address(_:_:)
func NWParametersSetReuseLocalAddress(parameters NWParameters, reuse_local_address bool) {
	if callErr := tryNWParametersSetReuseLocalAddress(parameters, reuse_local_address); callErr != nil {
		panic(callErr)
	}
}

var _nWParametersSetServiceClass func(parameters NWParameters, service_class NWServiceClass)
var _nWParametersSetServiceClassErr error

func tryNWParametersSetServiceClass(parameters NWParameters, service_class NWServiceClass) error {
	if _nWParametersSetServiceClass == nil {
		return symbolCallError("nw_parameters_set_service_class", "10.14", _nWParametersSetServiceClassErr)
	}
	_nWParametersSetServiceClass(parameters, service_class)
	return nil
}

// NWParametersSetServiceClass sets a level of service quality to use for connections.
//
// See: https://developer.apple.com/documentation/Network/nw_parameters_set_service_class(_:_:)
func NWParametersSetServiceClass(parameters NWParameters, service_class NWServiceClass) {
	if callErr := tryNWParametersSetServiceClass(parameters, service_class); callErr != nil {
		panic(callErr)
	}
}

var _nWPathCopyEffectiveLocalEndpoint func(path NWPath) NWEndpoint
var _nWPathCopyEffectiveLocalEndpointErr error

func tryNWPathCopyEffectiveLocalEndpoint(path NWPath) (NWEndpoint, error) {
	if _nWPathCopyEffectiveLocalEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_path_copy_effective_local_endpoint", "10.14", _nWPathCopyEffectiveLocalEndpointErr)
	}
	return _nWPathCopyEffectiveLocalEndpoint(path), nil
}

// NWPathCopyEffectiveLocalEndpoint accesses the local endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_local_endpoint(_:)
func NWPathCopyEffectiveLocalEndpoint(path NWPath) NWEndpoint {
	result, callErr := tryNWPathCopyEffectiveLocalEndpoint(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathCopyEffectiveRemoteEndpoint func(path NWPath) NWEndpoint
var _nWPathCopyEffectiveRemoteEndpointErr error

func tryNWPathCopyEffectiveRemoteEndpoint(path NWPath) (NWEndpoint, error) {
	if _nWPathCopyEffectiveRemoteEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_path_copy_effective_remote_endpoint", "10.14", _nWPathCopyEffectiveRemoteEndpointErr)
	}
	return _nWPathCopyEffectiveRemoteEndpoint(path), nil
}

// NWPathCopyEffectiveRemoteEndpoint accesses the remote endpoint in use by a connection’s network path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_copy_effective_remote_endpoint(_:)
func NWPathCopyEffectiveRemoteEndpoint(path NWPath) NWEndpoint {
	result, callErr := tryNWPathCopyEffectiveRemoteEndpoint(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathEnumerateGateways func(path NWPath, enumerate_block unsafe.Pointer)
var _nWPathEnumerateGatewaysErr error

func tryNWPathEnumerateGateways(path NWPath, enumerate_block NWPathEnumerateGatewaysBlock) error {
	if _nWPathEnumerateGateways == nil {
		return symbolCallError("nw_path_enumerate_gateways", "10.15", _nWPathEnumerateGatewaysErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return enumerate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWPathEnumerateGateways(path, _block0)
	return nil
}

// NWPathEnumerateGateways enumerates the list of gateways configured on the interfaces available to a path.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways(_:_:)
func NWPathEnumerateGateways(path NWPath, enumerate_block NWPathEnumerateGatewaysBlock) {
	if callErr := tryNWPathEnumerateGateways(path, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWPathEnumerateInterfaces func(path NWPath, enumerate_block unsafe.Pointer)
var _nWPathEnumerateInterfacesErr error

func tryNWPathEnumerateInterfaces(path NWPath, enumerate_block NWPathEnumerateInterfacesBlock) error {
	if _nWPathEnumerateInterfaces == nil {
		return symbolCallError("nw_path_enumerate_interfaces", "10.14", _nWPathEnumerateInterfacesErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) bool { return enumerate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWPathEnumerateInterfaces(path, _block0)
	return nil
}

// NWPathEnumerateInterfaces enumerates the list of all interfaces available to the path, in order of preference.
//
// See: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces(_:_:)
func NWPathEnumerateInterfaces(path NWPath, enumerate_block NWPathEnumerateInterfacesBlock) {
	if callErr := tryNWPathEnumerateInterfaces(path, enumerate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWPathGetLinkQuality func(path NWPath) NWLinkQuality
var _nWPathGetLinkQualityErr error

func tryNWPathGetLinkQuality(path NWPath) (NWLinkQuality, error) {
	if _nWPathGetLinkQuality == nil {
		return *new(NWLinkQuality), symbolCallError("nw_path_get_link_quality", "26.0", _nWPathGetLinkQualityErr)
	}
	return _nWPathGetLinkQuality(path), nil
}

// NWPathGetLinkQuality.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_link_quality(_:)
func NWPathGetLinkQuality(path NWPath) NWLinkQuality {
	result, callErr := tryNWPathGetLinkQuality(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathGetStatus func(path NWPath) NWPathStatus
var _nWPathGetStatusErr error

func tryNWPathGetStatus(path NWPath) (NWPathStatus, error) {
	if _nWPathGetStatus == nil {
		return *new(NWPathStatus), symbolCallError("nw_path_get_status", "10.14", _nWPathGetStatusErr)
	}
	return _nWPathGetStatus(path), nil
}

// NWPathGetStatus checks whether a path can be used by connections.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_status(_:)
func NWPathGetStatus(path NWPath) NWPathStatus {
	result, callErr := tryNWPathGetStatus(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathGetUnsatisfiedReason func(path NWPath) NWPathUnsatisfiedReason
var _nWPathGetUnsatisfiedReasonErr error

func tryNWPathGetUnsatisfiedReason(path NWPath) (NWPathUnsatisfiedReason, error) {
	if _nWPathGetUnsatisfiedReason == nil {
		return *new(NWPathUnsatisfiedReason), symbolCallError("nw_path_get_unsatisfied_reason", "11.0", _nWPathGetUnsatisfiedReasonErr)
	}
	return _nWPathGetUnsatisfiedReason(path), nil
}

// NWPathGetUnsatisfiedReason.
//
// See: https://developer.apple.com/documentation/Network/nw_path_get_unsatisfied_reason(_:)
func NWPathGetUnsatisfiedReason(path NWPath) NWPathUnsatisfiedReason {
	result, callErr := tryNWPathGetUnsatisfiedReason(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathHasDns func(path NWPath) bool
var _nWPathHasDnsErr error

func tryNWPathHasDns(path NWPath) (bool, error) {
	if _nWPathHasDns == nil {
		return false, symbolCallError("nw_path_has_dns", "10.14", _nWPathHasDnsErr)
	}
	return _nWPathHasDns(path), nil
}

// NWPathHasDns checks whether the path has a DNS server configured.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_dns(_:)
func NWPathHasDns(path NWPath) bool {
	result, callErr := tryNWPathHasDns(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathHasIpv4 func(path NWPath) bool
var _nWPathHasIpv4Err error

func tryNWPathHasIpv4(path NWPath) (bool, error) {
	if _nWPathHasIpv4 == nil {
		return false, symbolCallError("nw_path_has_ipv4", "10.14", _nWPathHasIpv4Err)
	}
	return _nWPathHasIpv4(path), nil
}

// NWPathHasIpv4 checks whether the path can route IPv4 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv4(_:)
func NWPathHasIpv4(path NWPath) bool {
	result, callErr := tryNWPathHasIpv4(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathHasIpv6 func(path NWPath) bool
var _nWPathHasIpv6Err error

func tryNWPathHasIpv6(path NWPath) (bool, error) {
	if _nWPathHasIpv6 == nil {
		return false, symbolCallError("nw_path_has_ipv6", "10.14", _nWPathHasIpv6Err)
	}
	return _nWPathHasIpv6(path), nil
}

// NWPathHasIpv6 checks whether the path can route IPv6 traffic.
//
// See: https://developer.apple.com/documentation/Network/nw_path_has_ipv6(_:)
func NWPathHasIpv6(path NWPath) bool {
	result, callErr := tryNWPathHasIpv6(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathIsConstrained func(path NWPath) bool
var _nWPathIsConstrainedErr error

func tryNWPathIsConstrained(path NWPath) (bool, error) {
	if _nWPathIsConstrained == nil {
		return false, symbolCallError("nw_path_is_constrained", "10.15", _nWPathIsConstrainedErr)
	}
	return _nWPathIsConstrained(path), nil
}

// NWPathIsConstrained checks whether the path uses an interface in Low Data Mode.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_constrained(_:)
func NWPathIsConstrained(path NWPath) bool {
	result, callErr := tryNWPathIsConstrained(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathIsEqual func(path NWPath, other_path NWPath) bool
var _nWPathIsEqualErr error

func tryNWPathIsEqual(path NWPath, other_path NWPath) (bool, error) {
	if _nWPathIsEqual == nil {
		return false, symbolCallError("nw_path_is_equal", "10.14", _nWPathIsEqualErr)
	}
	return _nWPathIsEqual(path, other_path), nil
}

// NWPathIsEqual compares if two paths are identical.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_equal(_:_:)
func NWPathIsEqual(path NWPath, other_path NWPath) bool {
	result, callErr := tryNWPathIsEqual(path, other_path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathIsExpensive func(path NWPath) bool
var _nWPathIsExpensiveErr error

func tryNWPathIsExpensive(path NWPath) (bool, error) {
	if _nWPathIsExpensive == nil {
		return false, symbolCallError("nw_path_is_expensive", "10.14", _nWPathIsExpensiveErr)
	}
	return _nWPathIsExpensive(path), nil
}

// NWPathIsExpensive checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_expensive(_:)
func NWPathIsExpensive(path NWPath) bool {
	result, callErr := tryNWPathIsExpensive(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathIsUltraConstrained func(path NWPath) bool
var _nWPathIsUltraConstrainedErr error

func tryNWPathIsUltraConstrained(path NWPath) (bool, error) {
	if _nWPathIsUltraConstrained == nil {
		return false, symbolCallError("nw_path_is_ultra_constrained", "26.0", _nWPathIsUltraConstrainedErr)
	}
	return _nWPathIsUltraConstrained(path), nil
}

// NWPathIsUltraConstrained.
//
// See: https://developer.apple.com/documentation/Network/nw_path_is_ultra_constrained(_:)
func NWPathIsUltraConstrained(path NWPath) bool {
	result, callErr := tryNWPathIsUltraConstrained(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathMonitorCancel func(monitor NWPathMonitor)
var _nWPathMonitorCancelErr error

func tryNWPathMonitorCancel(monitor NWPathMonitor) error {
	if _nWPathMonitorCancel == nil {
		return symbolCallError("nw_path_monitor_cancel", "10.14", _nWPathMonitorCancelErr)
	}
	_nWPathMonitorCancel(monitor)
	return nil
}

// NWPathMonitorCancel stops receiving network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel(_:)
func NWPathMonitorCancel(monitor NWPathMonitor) {
	if callErr := tryNWPathMonitorCancel(monitor); callErr != nil {
		panic(callErr)
	}
}

var _nWPathMonitorCreate func() NWPathMonitor
var _nWPathMonitorCreateErr error

func tryNWPathMonitorCreate() (NWPathMonitor, error) {
	if _nWPathMonitorCreate == nil {
		return *new(NWPathMonitor), symbolCallError("nw_path_monitor_create", "10.14", _nWPathMonitorCreateErr)
	}
	return _nWPathMonitorCreate(), nil
}

// NWPathMonitorCreate initializes a path monitor to observe all available interface types.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create()
func NWPathMonitorCreate() NWPathMonitor {
	result, callErr := tryNWPathMonitorCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathMonitorCreateForEthernetChannel func() NWPathMonitor
var _nWPathMonitorCreateForEthernetChannelErr error

func tryNWPathMonitorCreateForEthernetChannel() (NWPathMonitor, error) {
	if _nWPathMonitorCreateForEthernetChannel == nil {
		return *new(NWPathMonitor), symbolCallError("nw_path_monitor_create_for_ethernet_channel", "13.0", _nWPathMonitorCreateForEthernetChannelErr)
	}
	return _nWPathMonitorCreateForEthernetChannel(), nil
}

// NWPathMonitorCreateForEthernetChannel.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_for_ethernet_channel()
func NWPathMonitorCreateForEthernetChannel() NWPathMonitor {
	result, callErr := tryNWPathMonitorCreateForEthernetChannel()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathMonitorCreateWithType func(required_interface_type NWInterfaceType) NWPathMonitor
var _nWPathMonitorCreateWithTypeErr error

func tryNWPathMonitorCreateWithType(required_interface_type NWInterfaceType) (NWPathMonitor, error) {
	if _nWPathMonitorCreateWithType == nil {
		return *new(NWPathMonitor), symbolCallError("nw_path_monitor_create_with_type", "10.14", _nWPathMonitorCreateWithTypeErr)
	}
	return _nWPathMonitorCreateWithType(required_interface_type), nil
}

// NWPathMonitorCreateWithType initializes a path monitor to observe a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_create_with_type(_:)
func NWPathMonitorCreateWithType(required_interface_type NWInterfaceType) NWPathMonitor {
	result, callErr := tryNWPathMonitorCreateWithType(required_interface_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPathMonitorProhibitInterfaceType func(monitor NWPathMonitor, interface_type NWInterfaceType)
var _nWPathMonitorProhibitInterfaceTypeErr error

func tryNWPathMonitorProhibitInterfaceType(monitor NWPathMonitor, interface_type NWInterfaceType) error {
	if _nWPathMonitorProhibitInterfaceType == nil {
		return symbolCallError("nw_path_monitor_prohibit_interface_type", "11.0", _nWPathMonitorProhibitInterfaceTypeErr)
	}
	_nWPathMonitorProhibitInterfaceType(monitor, interface_type)
	return nil
}

// NWPathMonitorProhibitInterfaceType prohibit a path monitor from using a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_prohibit_interface_type(_:_:)
func NWPathMonitorProhibitInterfaceType(monitor NWPathMonitor, interface_type NWInterfaceType) {
	if callErr := tryNWPathMonitorProhibitInterfaceType(monitor, interface_type); callErr != nil {
		panic(callErr)
	}
}

var _nWPathMonitorSetCancelHandler func(monitor NWPathMonitor, cancel_handler unsafe.Pointer)
var _nWPathMonitorSetCancelHandlerErr error

func tryNWPathMonitorSetCancelHandler(monitor NWPathMonitor, cancel_handler NWPathMonitorCancelHandler) error {
	if _nWPathMonitorSetCancelHandler == nil {
		return symbolCallError("nw_path_monitor_set_cancel_handler", "10.14", _nWPathMonitorSetCancelHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { cancel_handler() })
	retainNetworkAsyncBlock(monitor.ID, "nw_path_monitor_set_cancel_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWPathMonitorSetCancelHandler(monitor, _block0)
	return nil
}

// NWPathMonitorSetCancelHandler sets a handler to determine when a monitor is fully cancelled and will no longer deliver events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_cancel_handler(_:_:)
func NWPathMonitorSetCancelHandler(monitor NWPathMonitor, cancel_handler NWPathMonitorCancelHandler) {
	if callErr := tryNWPathMonitorSetCancelHandler(monitor, cancel_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWPathMonitorSetQueue func(monitor NWPathMonitor, queue uintptr)
var _nWPathMonitorSetQueueErr error

func tryNWPathMonitorSetQueue(monitor NWPathMonitor, queue dispatch.Queue) error {
	if _nWPathMonitorSetQueue == nil {
		return symbolCallError("nw_path_monitor_set_queue", "10.14", _nWPathMonitorSetQueueErr)
	}
	_nWPathMonitorSetQueue(monitor, uintptr(queue.Handle()))
	return nil
}

// NWPathMonitorSetQueue sets a queue on which to deliver path events.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_queue(_:_:)
func NWPathMonitorSetQueue(monitor NWPathMonitor, queue dispatch.Queue) {
	if callErr := tryNWPathMonitorSetQueue(monitor, queue); callErr != nil {
		panic(callErr)
	}
}

var _nWPathMonitorSetUpdateHandler func(monitor NWPathMonitor, update_handler unsafe.Pointer)
var _nWPathMonitorSetUpdateHandlerErr error

func tryNWPathMonitorSetUpdateHandler(monitor NWPathMonitor, update_handler NWPathMonitorUpdateHandler) error {
	if _nWPathMonitorSetUpdateHandler == nil {
		return symbolCallError("nw_path_monitor_set_update_handler", "10.14", _nWPathMonitorSetUpdateHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { update_handler(objectivec.ObjectFromID(blockArg0)) })
	retainNetworkAsyncBlock(monitor.ID, "nw_path_monitor_set_update_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWPathMonitorSetUpdateHandler(monitor, _block0)
	return nil
}

// NWPathMonitorSetUpdateHandler sets a handler to receive network path updates.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_set_update_handler(_:_:)
func NWPathMonitorSetUpdateHandler(monitor NWPathMonitor, update_handler NWPathMonitorUpdateHandler) {
	if callErr := tryNWPathMonitorSetUpdateHandler(monitor, update_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWPathMonitorStart func(monitor NWPathMonitor)
var _nWPathMonitorStartErr error

func tryNWPathMonitorStart(monitor NWPathMonitor) error {
	if _nWPathMonitorStart == nil {
		return symbolCallError("nw_path_monitor_start", "10.14", _nWPathMonitorStartErr)
	}
	_nWPathMonitorStart(monitor)
	return nil
}

// NWPathMonitorStart starts monitoring path changes.
//
// See: https://developer.apple.com/documentation/Network/nw_path_monitor_start(_:)
func NWPathMonitorStart(monitor NWPathMonitor) {
	if callErr := tryNWPathMonitorStart(monitor); callErr != nil {
		panic(callErr)
	}
}

var _nWPathUsesInterfaceType func(path NWPath, interface_type NWInterfaceType) bool
var _nWPathUsesInterfaceTypeErr error

func tryNWPathUsesInterfaceType(path NWPath, interface_type NWInterfaceType) (bool, error) {
	if _nWPathUsesInterfaceType == nil {
		return false, symbolCallError("nw_path_uses_interface_type", "10.14", _nWPathUsesInterfaceTypeErr)
	}
	return _nWPathUsesInterfaceType(path, interface_type), nil
}

// NWPathUsesInterfaceType checks if connections using the path may send traffic over a specific interface type.
//
// See: https://developer.apple.com/documentation/Network/nw_path_uses_interface_type(_:_:)
func NWPathUsesInterfaceType(path NWPath, interface_type NWInterfaceType) bool {
	result, callErr := tryNWPathUsesInterfaceType(path, interface_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPrivacyContextAddProxy func(privacy_context NWPrivacyContext, proxy_config NWProxyConfig)
var _nWPrivacyContextAddProxyErr error

func tryNWPrivacyContextAddProxy(privacy_context NWPrivacyContext, proxy_config NWProxyConfig) error {
	if _nWPrivacyContextAddProxy == nil {
		return symbolCallError("nw_privacy_context_add_proxy", "14.0", _nWPrivacyContextAddProxyErr)
	}
	_nWPrivacyContextAddProxy(privacy_context, proxy_config)
	return nil
}

// NWPrivacyContextAddProxy applies a proxy configuration to all connections associated with this context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_add_proxy(_:_:)
func NWPrivacyContextAddProxy(privacy_context NWPrivacyContext, proxy_config NWProxyConfig) {
	if callErr := tryNWPrivacyContextAddProxy(privacy_context, proxy_config); callErr != nil {
		panic(callErr)
	}
}

var _nWPrivacyContextClearProxies func(privacy_context NWPrivacyContext)
var _nWPrivacyContextClearProxiesErr error

func tryNWPrivacyContextClearProxies(privacy_context NWPrivacyContext) error {
	if _nWPrivacyContextClearProxies == nil {
		return symbolCallError("nw_privacy_context_clear_proxies", "14.0", _nWPrivacyContextClearProxiesErr)
	}
	_nWPrivacyContextClearProxies(privacy_context)
	return nil
}

// NWPrivacyContextClearProxies clears out any proxies added using nw_privacy_context_add_proxy(_:_:)
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_clear_proxies(_:)
func NWPrivacyContextClearProxies(privacy_context NWPrivacyContext) {
	if callErr := tryNWPrivacyContextClearProxies(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nWPrivacyContextCreate func(description string) NWPrivacyContext
var _nWPrivacyContextCreateErr error

func tryNWPrivacyContextCreate(description string) (NWPrivacyContext, error) {
	if _nWPrivacyContextCreate == nil {
		return *new(NWPrivacyContext), symbolCallError("nw_privacy_context_create", "11.0", _nWPrivacyContextCreateErr)
	}
	return _nWPrivacyContextCreate(description), nil
}

// NWPrivacyContextCreate initializes a privacy context with a description string.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_create(_:)
func NWPrivacyContextCreate(description string) NWPrivacyContext {
	result, callErr := tryNWPrivacyContextCreate(description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWPrivacyContextDisableLogging func(privacy_context NWPrivacyContext)
var _nWPrivacyContextDisableLoggingErr error

func tryNWPrivacyContextDisableLogging(privacy_context NWPrivacyContext) error {
	if _nWPrivacyContextDisableLogging == nil {
		return symbolCallError("nw_privacy_context_disable_logging", "11.0", _nWPrivacyContextDisableLoggingErr)
	}
	_nWPrivacyContextDisableLogging(privacy_context)
	return nil
}

// NWPrivacyContextDisableLogging disables system logging of connection activity.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_disable_logging(_:)
func NWPrivacyContextDisableLogging(privacy_context NWPrivacyContext) {
	if callErr := tryNWPrivacyContextDisableLogging(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nWPrivacyContextFlushCache func(privacy_context NWPrivacyContext)
var _nWPrivacyContextFlushCacheErr error

func tryNWPrivacyContextFlushCache(privacy_context NWPrivacyContext) error {
	if _nWPrivacyContextFlushCache == nil {
		return symbolCallError("nw_privacy_context_flush_cache", "11.0", _nWPrivacyContextFlushCacheErr)
	}
	_nWPrivacyContextFlushCache(privacy_context)
	return nil
}

// NWPrivacyContextFlushCache flushes all cached data, such as TLS session state, created by connections associated with the privacy context.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_flush_cache(_:)
func NWPrivacyContextFlushCache(privacy_context NWPrivacyContext) {
	if callErr := tryNWPrivacyContextFlushCache(privacy_context); callErr != nil {
		panic(callErr)
	}
}

var _nWPrivacyContextRequireEncryptedNameResolution func(privacy_context NWPrivacyContext, require_encrypted_name_resolution bool, fallback_resolver_config NWResolverConfig)
var _nWPrivacyContextRequireEncryptedNameResolutionErr error

func tryNWPrivacyContextRequireEncryptedNameResolution(privacy_context NWPrivacyContext, require_encrypted_name_resolution bool, fallback_resolver_config NWResolverConfig) error {
	if _nWPrivacyContextRequireEncryptedNameResolution == nil {
		return symbolCallError("nw_privacy_context_require_encrypted_name_resolution", "11.0", _nWPrivacyContextRequireEncryptedNameResolutionErr)
	}
	_nWPrivacyContextRequireEncryptedNameResolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config)
	return nil
}

// NWPrivacyContextRequireEncryptedNameResolution requires that any DNS name resolution for connections associated with this context use encrypted transports, such as TLS or HTTPS.
//
// See: https://developer.apple.com/documentation/Network/nw_privacy_context_require_encrypted_name_resolution(_:_:_:)
func NWPrivacyContextRequireEncryptedNameResolution(privacy_context NWPrivacyContext, require_encrypted_name_resolution bool, fallback_resolver_config NWResolverConfig) {
	if callErr := tryNWPrivacyContextRequireEncryptedNameResolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config); callErr != nil {
		panic(callErr)
	}
}

var _nWProtocolCopyIPDefinition func() NWProtocolDefinition
var _nWProtocolCopyIPDefinitionErr error

func tryNWProtocolCopyIPDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyIPDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_ip_definition", "10.14", _nWProtocolCopyIPDefinitionErr)
	}
	return _nWProtocolCopyIPDefinition(), nil
}

// NWProtocolCopyIPDefinition accesses the system definition of the Internet Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ip_definition()
func NWProtocolCopyIPDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyIPDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolCopyQuicDefinition func() NWProtocolDefinition
var _nWProtocolCopyQuicDefinitionErr error

func tryNWProtocolCopyQuicDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyQuicDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_quic_definition", "12.0", _nWProtocolCopyQuicDefinitionErr)
	}
	return _nWProtocolCopyQuicDefinition(), nil
}

// NWProtocolCopyQuicDefinition accesses the system definition of the QUIC transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_quic_definition()
func NWProtocolCopyQuicDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyQuicDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolCopyTCPDefinition func() NWProtocolDefinition
var _nWProtocolCopyTCPDefinitionErr error

func tryNWProtocolCopyTCPDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyTCPDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_tcp_definition", "10.14", _nWProtocolCopyTCPDefinitionErr)
	}
	return _nWProtocolCopyTCPDefinition(), nil
}

// NWProtocolCopyTCPDefinition accesses the system definition of the Transport Control Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tcp_definition()
func NWProtocolCopyTCPDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyTCPDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolCopyTLSDefinition func() NWProtocolDefinition
var _nWProtocolCopyTLSDefinitionErr error

func tryNWProtocolCopyTLSDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyTLSDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_tls_definition", "10.14", _nWProtocolCopyTLSDefinitionErr)
	}
	return _nWProtocolCopyTLSDefinition(), nil
}

// NWProtocolCopyTLSDefinition accesses the system definition of the Transport Layer Security protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_tls_definition()
func NWProtocolCopyTLSDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyTLSDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolCopyUDPDefinition func() NWProtocolDefinition
var _nWProtocolCopyUDPDefinitionErr error

func tryNWProtocolCopyUDPDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyUDPDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_udp_definition", "10.14", _nWProtocolCopyUDPDefinitionErr)
	}
	return _nWProtocolCopyUDPDefinition(), nil
}

// NWProtocolCopyUDPDefinition accesses the system definition of the User Datagram Protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_udp_definition()
func NWProtocolCopyUDPDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyUDPDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolCopyWsDefinition func() NWProtocolDefinition
var _nWProtocolCopyWsDefinitionErr error

func tryNWProtocolCopyWsDefinition() (NWProtocolDefinition, error) {
	if _nWProtocolCopyWsDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_copy_ws_definition", "10.15", _nWProtocolCopyWsDefinitionErr)
	}
	return _nWProtocolCopyWsDefinition(), nil
}

// NWProtocolCopyWsDefinition accesses the system definition of the WebSocket protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_copy_ws_definition()
func NWProtocolCopyWsDefinition() NWProtocolDefinition {
	result, callErr := tryNWProtocolCopyWsDefinition()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolDefinitionIsEqual func(definition1 NWProtocolDefinition, definition2 NWProtocolDefinition) bool
var _nWProtocolDefinitionIsEqualErr error

func tryNWProtocolDefinitionIsEqual(definition1 NWProtocolDefinition, definition2 NWProtocolDefinition) (bool, error) {
	if _nWProtocolDefinitionIsEqual == nil {
		return false, symbolCallError("nw_protocol_definition_is_equal", "10.14", _nWProtocolDefinitionIsEqualErr)
	}
	return _nWProtocolDefinitionIsEqual(definition1, definition2), nil
}

// NWProtocolDefinitionIsEqual compares two protocol definitions, and returns true if they represent the same protocol implementation.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_definition_is_equal(_:_:)
func NWProtocolDefinitionIsEqual(definition1 NWProtocolDefinition, definition2 NWProtocolDefinition) bool {
	result, callErr := tryNWProtocolDefinitionIsEqual(definition1, definition2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataCopyDefinition func(metadata NWProtocolMetadata) NWProtocolDefinition
var _nWProtocolMetadataCopyDefinitionErr error

func tryNWProtocolMetadataCopyDefinition(metadata NWProtocolMetadata) (NWProtocolDefinition, error) {
	if _nWProtocolMetadataCopyDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_metadata_copy_definition", "10.14", _nWProtocolMetadataCopyDefinitionErr)
	}
	return _nWProtocolMetadataCopyDefinition(metadata), nil
}

// NWProtocolMetadataCopyDefinition accesses the protocol definition associated with the metadata object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_copy_definition(_:)
func NWProtocolMetadataCopyDefinition(metadata NWProtocolMetadata) NWProtocolDefinition {
	result, callErr := tryNWProtocolMetadataCopyDefinition(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsFramerMessage func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsFramerMessageErr error

func tryNWProtocolMetadataIsFramerMessage(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsFramerMessage == nil {
		return false, symbolCallError("nw_protocol_metadata_is_framer_message", "10.15", _nWProtocolMetadataIsFramerMessageErr)
	}
	return _nWProtocolMetadataIsFramerMessage(metadata), nil
}

// NWProtocolMetadataIsFramerMessage checks if a metadata object represents a custom framer protocol message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_framer_message(_:)
func NWProtocolMetadataIsFramerMessage(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsFramerMessage(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsIP func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsIPErr error

func tryNWProtocolMetadataIsIP(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsIP == nil {
		return false, symbolCallError("nw_protocol_metadata_is_ip", "10.14", _nWProtocolMetadataIsIPErr)
	}
	return _nWProtocolMetadataIsIP(metadata), nil
}

// NWProtocolMetadataIsIP checks whether a metadata object represents an IP packet.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ip(_:)
func NWProtocolMetadataIsIP(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsIP(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsQuic func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsQuicErr error

func tryNWProtocolMetadataIsQuic(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsQuic == nil {
		return false, symbolCallError("nw_protocol_metadata_is_quic", "12.0", _nWProtocolMetadataIsQuicErr)
	}
	return _nWProtocolMetadataIsQuic(metadata), nil
}

// NWProtocolMetadataIsQuic checks whether a metadata object contains QUIC connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_quic(_:)
func NWProtocolMetadataIsQuic(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsQuic(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsTCP func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsTCPErr error

func tryNWProtocolMetadataIsTCP(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsTCP == nil {
		return false, symbolCallError("nw_protocol_metadata_is_tcp", "10.14", _nWProtocolMetadataIsTCPErr)
	}
	return _nWProtocolMetadataIsTCP(metadata), nil
}

// NWProtocolMetadataIsTCP checks whether a metadata object contains TCP connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tcp(_:)
func NWProtocolMetadataIsTCP(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsTCP(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsTLS func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsTLSErr error

func tryNWProtocolMetadataIsTLS(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsTLS == nil {
		return false, symbolCallError("nw_protocol_metadata_is_tls", "10.14", _nWProtocolMetadataIsTLSErr)
	}
	return _nWProtocolMetadataIsTLS(metadata), nil
}

// NWProtocolMetadataIsTLS checks whether a metadata object contains TLS connection state.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tls(_:)
func NWProtocolMetadataIsTLS(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsTLS(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsUDP func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsUDPErr error

func tryNWProtocolMetadataIsUDP(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsUDP == nil {
		return false, symbolCallError("nw_protocol_metadata_is_udp", "10.14", _nWProtocolMetadataIsUDPErr)
	}
	return _nWProtocolMetadataIsUDP(metadata), nil
}

// NWProtocolMetadataIsUDP checks whether a metadata object represents a UDP datagram.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_udp(_:)
func NWProtocolMetadataIsUDP(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsUDP(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolMetadataIsWs func(metadata NWProtocolMetadata) bool
var _nWProtocolMetadataIsWsErr error

func tryNWProtocolMetadataIsWs(metadata NWProtocolMetadata) (bool, error) {
	if _nWProtocolMetadataIsWs == nil {
		return false, symbolCallError("nw_protocol_metadata_is_ws", "10.15", _nWProtocolMetadataIsWsErr)
	}
	return _nWProtocolMetadataIsWs(metadata), nil
}

// NWProtocolMetadataIsWs checks whether a metadata object represents a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ws(_:)
func NWProtocolMetadataIsWs(metadata NWProtocolMetadata) bool {
	result, callErr := tryNWProtocolMetadataIsWs(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolOptionsCopyDefinition func(options NWProtocolOptions) NWProtocolDefinition
var _nWProtocolOptionsCopyDefinitionErr error

func tryNWProtocolOptionsCopyDefinition(options NWProtocolOptions) (NWProtocolDefinition, error) {
	if _nWProtocolOptionsCopyDefinition == nil {
		return *new(NWProtocolDefinition), symbolCallError("nw_protocol_options_copy_definition", "10.14", _nWProtocolOptionsCopyDefinitionErr)
	}
	return _nWProtocolOptionsCopyDefinition(options), nil
}

// NWProtocolOptionsCopyDefinition accesses the protocol definition associated with the options object.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_copy_definition(_:)
func NWProtocolOptionsCopyDefinition(options NWProtocolOptions) NWProtocolDefinition {
	result, callErr := tryNWProtocolOptionsCopyDefinition(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolOptionsIsQuic func(options NWProtocolOptions) bool
var _nWProtocolOptionsIsQuicErr error

func tryNWProtocolOptionsIsQuic(options NWProtocolOptions) (bool, error) {
	if _nWProtocolOptionsIsQuic == nil {
		return false, symbolCallError("nw_protocol_options_is_quic", "12.0", _nWProtocolOptionsIsQuicErr)
	}
	return _nWProtocolOptionsIsQuic(options), nil
}

// NWProtocolOptionsIsQuic checks whether an options object uses the QUIC protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_options_is_quic(_:)
func NWProtocolOptionsIsQuic(options NWProtocolOptions) bool {
	result, callErr := tryNWProtocolOptionsIsQuic(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolStackClearApplicationProtocols func(stack NWProtocolStack)
var _nWProtocolStackClearApplicationProtocolsErr error

func tryNWProtocolStackClearApplicationProtocols(stack NWProtocolStack) error {
	if _nWProtocolStackClearApplicationProtocols == nil {
		return symbolCallError("nw_protocol_stack_clear_application_protocols", "10.14", _nWProtocolStackClearApplicationProtocolsErr)
	}
	_nWProtocolStackClearApplicationProtocols(stack)
	return nil
}

// NWProtocolStackClearApplicationProtocols removes all application protocols from the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_clear_application_protocols(_:)
func NWProtocolStackClearApplicationProtocols(stack NWProtocolStack) {
	if callErr := tryNWProtocolStackClearApplicationProtocols(stack); callErr != nil {
		panic(callErr)
	}
}

var _nWProtocolStackCopyInternetProtocol func(stack NWProtocolStack) NWProtocolOptions
var _nWProtocolStackCopyInternetProtocolErr error

func tryNWProtocolStackCopyInternetProtocol(stack NWProtocolStack) (NWProtocolOptions, error) {
	if _nWProtocolStackCopyInternetProtocol == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_protocol_stack_copy_internet_protocol", "10.14", _nWProtocolStackCopyInternetProtocolErr)
	}
	return _nWProtocolStackCopyInternetProtocol(stack), nil
}

// NWProtocolStackCopyInternetProtocol accesses the protocol stack’s Internet Protocol options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_internet_protocol(_:)
func NWProtocolStackCopyInternetProtocol(stack NWProtocolStack) NWProtocolOptions {
	result, callErr := tryNWProtocolStackCopyInternetProtocol(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolStackCopyTransportProtocol func(stack NWProtocolStack) NWProtocolOptions
var _nWProtocolStackCopyTransportProtocolErr error

func tryNWProtocolStackCopyTransportProtocol(stack NWProtocolStack) (NWProtocolOptions, error) {
	if _nWProtocolStackCopyTransportProtocol == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_protocol_stack_copy_transport_protocol", "10.14", _nWProtocolStackCopyTransportProtocolErr)
	}
	return _nWProtocolStackCopyTransportProtocol(stack), nil
}

// NWProtocolStackCopyTransportProtocol accesses the options for the protocol stack’s transport protocol.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_transport_protocol(_:)
func NWProtocolStackCopyTransportProtocol(stack NWProtocolStack) NWProtocolOptions {
	result, callErr := tryNWProtocolStackCopyTransportProtocol(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProtocolStackIterateApplicationProtocols func(stack NWProtocolStack, iterate_block unsafe.Pointer)
var _nWProtocolStackIterateApplicationProtocolsErr error

func tryNWProtocolStackIterateApplicationProtocols(stack NWProtocolStack, iterate_block NWProtocolStackIterateProtocolsBlock) error {
	if _nWProtocolStackIterateApplicationProtocols == nil {
		return symbolCallError("nw_protocol_stack_iterate_application_protocols", "10.14", _nWProtocolStackIterateApplicationProtocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) { iterate_block(objectivec.ObjectFromID(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWProtocolStackIterateApplicationProtocols(stack, _block0)
	return nil
}

// NWProtocolStackIterateApplicationProtocols iterates through the array of application protocol options that will be used by connections and listeners.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_application_protocols(_:_:)
func NWProtocolStackIterateApplicationProtocols(stack NWProtocolStack, iterate_block NWProtocolStackIterateProtocolsBlock) {
	if callErr := tryNWProtocolStackIterateApplicationProtocols(stack, iterate_block); callErr != nil {
		panic(callErr)
	}
}

var _nWProtocolStackPrependApplicationProtocol func(stack NWProtocolStack, protocol_ NWProtocolOptions)
var _nWProtocolStackPrependApplicationProtocolErr error

func tryNWProtocolStackPrependApplicationProtocol(stack NWProtocolStack, protocol_ NWProtocolOptions) error {
	if _nWProtocolStackPrependApplicationProtocol == nil {
		return symbolCallError("nw_protocol_stack_prepend_application_protocol", "10.14", _nWProtocolStackPrependApplicationProtocolErr)
	}
	_nWProtocolStackPrependApplicationProtocol(stack, protocol_)
	return nil
}

// NWProtocolStackPrependApplicationProtocol adds a protocol onto the top of the protocol stack.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_prepend_application_protocol(_:_:)
func NWProtocolStackPrependApplicationProtocol(stack NWProtocolStack, protocol_ NWProtocolOptions) {
	if callErr := tryNWProtocolStackPrependApplicationProtocol(stack, protocol_); callErr != nil {
		panic(callErr)
	}
}

var _nWProtocolStackSetTransportProtocol func(stack NWProtocolStack, protocol_ NWProtocolOptions)
var _nWProtocolStackSetTransportProtocolErr error

func tryNWProtocolStackSetTransportProtocol(stack NWProtocolStack, protocol_ NWProtocolOptions) error {
	if _nWProtocolStackSetTransportProtocol == nil {
		return symbolCallError("nw_protocol_stack_set_transport_protocol", "10.14", _nWProtocolStackSetTransportProtocolErr)
	}
	_nWProtocolStackSetTransportProtocol(stack, protocol_)
	return nil
}

// NWProtocolStackSetTransportProtocol replaces the protocol stack’s transport protocol with a new set of options.
//
// See: https://developer.apple.com/documentation/Network/nw_protocol_stack_set_transport_protocol(_:_:)
func NWProtocolStackSetTransportProtocol(stack NWProtocolStack, protocol_ NWProtocolOptions) {
	if callErr := tryNWProtocolStackSetTransportProtocol(stack, protocol_); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigAddExcludedDomain func(config NWProxyConfig, excluded_domain string)
var _nWProxyConfigAddExcludedDomainErr error

func tryNWProxyConfigAddExcludedDomain(config NWProxyConfig, excluded_domain string) error {
	if _nWProxyConfigAddExcludedDomain == nil {
		return symbolCallError("nw_proxy_config_add_excluded_domain", "14.0", _nWProxyConfigAddExcludedDomainErr)
	}
	_nWProxyConfigAddExcludedDomain(config, excluded_domain)
	return nil
}

// NWProxyConfigAddExcludedDomain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_excluded_domain(_:_:)
func NWProxyConfigAddExcludedDomain(config NWProxyConfig, excluded_domain string) {
	if callErr := tryNWProxyConfigAddExcludedDomain(config, excluded_domain); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigAddMatchDomain func(config NWProxyConfig, match_domain string)
var _nWProxyConfigAddMatchDomainErr error

func tryNWProxyConfigAddMatchDomain(config NWProxyConfig, match_domain string) error {
	if _nWProxyConfigAddMatchDomain == nil {
		return symbolCallError("nw_proxy_config_add_match_domain", "14.0", _nWProxyConfigAddMatchDomainErr)
	}
	_nWProxyConfigAddMatchDomain(config, match_domain)
	return nil
}

// NWProxyConfigAddMatchDomain.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_add_match_domain(_:_:)
func NWProxyConfigAddMatchDomain(config NWProxyConfig, match_domain string) {
	if callErr := tryNWProxyConfigAddMatchDomain(config, match_domain); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigClearExcludedDomains func(config NWProxyConfig)
var _nWProxyConfigClearExcludedDomainsErr error

func tryNWProxyConfigClearExcludedDomains(config NWProxyConfig) error {
	if _nWProxyConfigClearExcludedDomains == nil {
		return symbolCallError("nw_proxy_config_clear_excluded_domains", "14.0", _nWProxyConfigClearExcludedDomainsErr)
	}
	_nWProxyConfigClearExcludedDomains(config)
	return nil
}

// NWProxyConfigClearExcludedDomains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_excluded_domains(_:)
func NWProxyConfigClearExcludedDomains(config NWProxyConfig) {
	if callErr := tryNWProxyConfigClearExcludedDomains(config); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigClearMatchDomains func(config NWProxyConfig)
var _nWProxyConfigClearMatchDomainsErr error

func tryNWProxyConfigClearMatchDomains(config NWProxyConfig) error {
	if _nWProxyConfigClearMatchDomains == nil {
		return symbolCallError("nw_proxy_config_clear_match_domains", "14.0", _nWProxyConfigClearMatchDomainsErr)
	}
	_nWProxyConfigClearMatchDomains(config)
	return nil
}

// NWProxyConfigClearMatchDomains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_match_domains(_:)
func NWProxyConfigClearMatchDomains(config NWProxyConfig) {
	if callErr := tryNWProxyConfigClearMatchDomains(config); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigCreateHttpConnect func(proxy_endpoint NWEndpoint, proxy_tls_options NWProtocolOptions) NWProxyConfig
var _nWProxyConfigCreateHttpConnectErr error

func tryNWProxyConfigCreateHttpConnect(proxy_endpoint NWEndpoint, proxy_tls_options NWProtocolOptions) (NWProxyConfig, error) {
	if _nWProxyConfigCreateHttpConnect == nil {
		return *new(NWProxyConfig), symbolCallError("nw_proxy_config_create_http_connect", "14.0", _nWProxyConfigCreateHttpConnectErr)
	}
	return _nWProxyConfigCreateHttpConnect(proxy_endpoint, proxy_tls_options), nil
}

// NWProxyConfigCreateHttpConnect initializes a legacy HTTP CONNECT configuration for a proxy server accessible using HTTP/1.1.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_http_connect(_:_:)
func NWProxyConfigCreateHttpConnect(proxy_endpoint NWEndpoint, proxy_tls_options NWProtocolOptions) NWProxyConfig {
	result, callErr := tryNWProxyConfigCreateHttpConnect(proxy_endpoint, proxy_tls_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProxyConfigCreateObliviousHttp func(relay NWRelayHop, relay_resource_path string, gateway_key_config *byte, gateway_key_config_length uintptr) NWProxyConfig
var _nWProxyConfigCreateObliviousHttpErr error

func tryNWProxyConfigCreateObliviousHttp(relay NWRelayHop, relay_resource_path string, gateway_key_config []byte, gateway_key_config_length uintptr) (NWProxyConfig, error) {
	if _nWProxyConfigCreateObliviousHttp == nil {
		return *new(NWProxyConfig), symbolCallError("nw_proxy_config_create_oblivious_http", "14.0", _nWProxyConfigCreateObliviousHttpErr)
	}
	return _nWProxyConfigCreateObliviousHttp(relay, relay_resource_path, unsafe.SliceData(gateway_key_config), gateway_key_config_length), nil
}

// NWProxyConfigCreateObliviousHttp initializes an Oblivious HTTP proxy configuration using a relay and a gateway.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_oblivious_http(_:_:_:_:)
func NWProxyConfigCreateObliviousHttp(relay NWRelayHop, relay_resource_path string, gateway_key_config []byte, gateway_key_config_length uintptr) NWProxyConfig {
	result, callErr := tryNWProxyConfigCreateObliviousHttp(relay, relay_resource_path, gateway_key_config, gateway_key_config_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProxyConfigCreateRelay func(first_hop NWRelayHop, second_hop NWRelayHop) NWProxyConfig
var _nWProxyConfigCreateRelayErr error

func tryNWProxyConfigCreateRelay(first_hop NWRelayHop, second_hop NWRelayHop) (NWProxyConfig, error) {
	if _nWProxyConfigCreateRelay == nil {
		return *new(NWProxyConfig), symbolCallError("nw_proxy_config_create_relay", "14.0", _nWProxyConfigCreateRelayErr)
	}
	return _nWProxyConfigCreateRelay(first_hop, second_hop), nil
}

// NWProxyConfigCreateRelay initializes a proxy configuration with one or two relay hops.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_relay(_:_:)
func NWProxyConfigCreateRelay(first_hop NWRelayHop, second_hop NWRelayHop) NWProxyConfig {
	result, callErr := tryNWProxyConfigCreateRelay(first_hop, second_hop)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProxyConfigCreateSocksv5 func(proxy_endpoint NWEndpoint) NWProxyConfig
var _nWProxyConfigCreateSocksv5Err error

func tryNWProxyConfigCreateSocksv5(proxy_endpoint NWEndpoint) (NWProxyConfig, error) {
	if _nWProxyConfigCreateSocksv5 == nil {
		return *new(NWProxyConfig), symbolCallError("nw_proxy_config_create_socksv5", "14.0", _nWProxyConfigCreateSocksv5Err)
	}
	return _nWProxyConfigCreateSocksv5(proxy_endpoint), nil
}

// NWProxyConfigCreateSocksv5 initializes a SOCKSv5 proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_create_socksv5(_:)
func NWProxyConfigCreateSocksv5(proxy_endpoint NWEndpoint) NWProxyConfig {
	result, callErr := tryNWProxyConfigCreateSocksv5(proxy_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProxyConfigEnumerateExcludedDomains func(config NWProxyConfig, enumerator unsafe.Pointer)
var _nWProxyConfigEnumerateExcludedDomainsErr error

func tryNWProxyConfigEnumerateExcludedDomains(config NWProxyConfig, enumerator NWProxyDomainEnumerator) error {
	if _nWProxyConfigEnumerateExcludedDomains == nil {
		return symbolCallError("nw_proxy_config_enumerate_excluded_domains", "14.0", _nWProxyConfigEnumerateExcludedDomainsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte) { enumerator(objc.GoString(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWProxyConfigEnumerateExcludedDomains(config, _block0)
	return nil
}

// NWProxyConfigEnumerateExcludedDomains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_excluded_domains(_:_:)
func NWProxyConfigEnumerateExcludedDomains(config NWProxyConfig, enumerator NWProxyDomainEnumerator) {
	if callErr := tryNWProxyConfigEnumerateExcludedDomains(config, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigEnumerateMatchDomains func(config NWProxyConfig, enumerator unsafe.Pointer)
var _nWProxyConfigEnumerateMatchDomainsErr error

func tryNWProxyConfigEnumerateMatchDomains(config NWProxyConfig, enumerator NWProxyDomainEnumerator) error {
	if _nWProxyConfigEnumerateMatchDomains == nil {
		return symbolCallError("nw_proxy_config_enumerate_match_domains", "14.0", _nWProxyConfigEnumerateMatchDomainsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte) { enumerator(objc.GoString(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_nWProxyConfigEnumerateMatchDomains(config, _block0)
	return nil
}

// NWProxyConfigEnumerateMatchDomains.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_match_domains(_:_:)
func NWProxyConfigEnumerateMatchDomains(config NWProxyConfig, enumerator NWProxyDomainEnumerator) {
	if callErr := tryNWProxyConfigEnumerateMatchDomains(config, enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigGetFailoverAllowed func(proxy_config NWProxyConfig) bool
var _nWProxyConfigGetFailoverAllowedErr error

func tryNWProxyConfigGetFailoverAllowed(proxy_config NWProxyConfig) (bool, error) {
	if _nWProxyConfigGetFailoverAllowed == nil {
		return false, symbolCallError("nw_proxy_config_get_failover_allowed", "14.0", _nWProxyConfigGetFailoverAllowedErr)
	}
	return _nWProxyConfigGetFailoverAllowed(proxy_config), nil
}

// NWProxyConfigGetFailoverAllowed checks if a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_get_failover_allowed(_:)
func NWProxyConfigGetFailoverAllowed(proxy_config NWProxyConfig) bool {
	result, callErr := tryNWProxyConfigGetFailoverAllowed(proxy_config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWProxyConfigSetFailoverAllowed func(proxy_config NWProxyConfig, failover_allowed bool)
var _nWProxyConfigSetFailoverAllowedErr error

func tryNWProxyConfigSetFailoverAllowed(proxy_config NWProxyConfig, failover_allowed bool) error {
	if _nWProxyConfigSetFailoverAllowed == nil {
		return symbolCallError("nw_proxy_config_set_failover_allowed", "14.0", _nWProxyConfigSetFailoverAllowedErr)
	}
	_nWProxyConfigSetFailoverAllowed(proxy_config, failover_allowed)
	return nil
}

// NWProxyConfigSetFailoverAllowed configures whether or not a proxy configuration allows failover to non-proxied connections.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_failover_allowed(_:_:)
func NWProxyConfigSetFailoverAllowed(proxy_config NWProxyConfig, failover_allowed bool) {
	if callErr := tryNWProxyConfigSetFailoverAllowed(proxy_config, failover_allowed); callErr != nil {
		panic(callErr)
	}
}

var _nWProxyConfigSetUsernameAndPassword func(proxy_config NWProxyConfig, username string, password string)
var _nWProxyConfigSetUsernameAndPasswordErr error

func tryNWProxyConfigSetUsernameAndPassword(proxy_config NWProxyConfig, username string, password string) error {
	if _nWProxyConfigSetUsernameAndPassword == nil {
		return symbolCallError("nw_proxy_config_set_username_and_password", "14.0", _nWProxyConfigSetUsernameAndPasswordErr)
	}
	_nWProxyConfigSetUsernameAndPassword(proxy_config, username, password)
	return nil
}

// NWProxyConfigSetUsernameAndPassword sets a username and password to use as authentication for a proxy configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_proxy_config_set_username_and_password(_:_:_:)
func NWProxyConfigSetUsernameAndPassword(proxy_config NWProxyConfig, username string, password string) {
	if callErr := tryNWProxyConfigSetUsernameAndPassword(proxy_config, username, password); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicAddTLSApplicationProtocol func(options NWProtocolOptions, application_protocol string)
var _nWQuicAddTLSApplicationProtocolErr error

func tryNWQuicAddTLSApplicationProtocol(options NWProtocolOptions, application_protocol string) error {
	if _nWQuicAddTLSApplicationProtocol == nil {
		return symbolCallError("nw_quic_add_tls_application_protocol", "12.0", _nWQuicAddTLSApplicationProtocolErr)
	}
	_nWQuicAddTLSApplicationProtocol(options, application_protocol)
	return nil
}

// NWQuicAddTLSApplicationProtocol adds a supported Application-Layer Protocol Negotiation value.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_add_tls_application_protocol(_:_:)
func NWQuicAddTLSApplicationProtocol(options NWProtocolOptions, application_protocol string) {
	if callErr := tryNWQuicAddTLSApplicationProtocol(options, application_protocol); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicCopySecProtocolMetadata func(metadata NWProtocolMetadata) security.Sec_protocol_metadata_t
var _nWQuicCopySecProtocolMetadataErr error

func tryNWQuicCopySecProtocolMetadata(metadata NWProtocolMetadata) (security.Sec_protocol_metadata_t, error) {
	if _nWQuicCopySecProtocolMetadata == nil {
		return *new(security.Sec_protocol_metadata_t), symbolCallError("nw_quic_copy_sec_protocol_metadata", "12.0", _nWQuicCopySecProtocolMetadataErr)
	}
	return _nWQuicCopySecProtocolMetadata(metadata), nil
}

// NWQuicCopySecProtocolMetadata accesses the result of the QUIC handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_metadata(_:)
func NWQuicCopySecProtocolMetadata(metadata NWProtocolMetadata) security.Sec_protocol_metadata_t {
	result, callErr := tryNWQuicCopySecProtocolMetadata(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicCopySecProtocolOptions func(options NWProtocolOptions) security.Sec_protocol_options_t
var _nWQuicCopySecProtocolOptionsErr error

func tryNWQuicCopySecProtocolOptions(options NWProtocolOptions) (security.Sec_protocol_options_t, error) {
	if _nWQuicCopySecProtocolOptions == nil {
		return *new(security.Sec_protocol_options_t), symbolCallError("nw_quic_copy_sec_protocol_options", "12.0", _nWQuicCopySecProtocolOptionsErr)
	}
	return _nWQuicCopySecProtocolOptions(options), nil
}

// NWQuicCopySecProtocolOptions accesses the handshake security options QUIC will use.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_options(_:)
func NWQuicCopySecProtocolOptions(options NWProtocolOptions) security.Sec_protocol_options_t {
	result, callErr := tryNWQuicCopySecProtocolOptions(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicCreateOptions func() NWProtocolOptions
var _nWQuicCreateOptionsErr error

func tryNWQuicCreateOptions() (NWProtocolOptions, error) {
	if _nWQuicCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_quic_create_options", "12.0", _nWQuicCreateOptionsErr)
	}
	return _nWQuicCreateOptions(), nil
}

// NWQuicCreateOptions initializes a default set of QUIC connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_create_options()
func NWQuicCreateOptions() NWProtocolOptions {
	result, callErr := tryNWQuicCreateOptions()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetApplicationError func(metadata NWProtocolMetadata) uint64
var _nWQuicGetApplicationErrorErr error

func tryNWQuicGetApplicationError(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetApplicationError == nil {
		return 0, symbolCallError("nw_quic_get_application_error", "12.0", _nWQuicGetApplicationErrorErr)
	}
	return _nWQuicGetApplicationError(metadata), nil
}

// NWQuicGetApplicationError accesses the QUIC application error code received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error(_:)
func NWQuicGetApplicationError(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetApplicationError(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetApplicationErrorReason func(metadata NWProtocolMetadata) *byte
var _nWQuicGetApplicationErrorReasonErr error

func tryNWQuicGetApplicationErrorReason(metadata NWProtocolMetadata) (*byte, error) {
	if _nWQuicGetApplicationErrorReason == nil {
		return nil, symbolCallError("nw_quic_get_application_error_reason", "12.0", _nWQuicGetApplicationErrorReasonErr)
	}
	return _nWQuicGetApplicationErrorReason(metadata), nil
}

// NWQuicGetApplicationErrorReason accesses the QUIC application error reason received from the peer.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_application_error_reason(_:)
func NWQuicGetApplicationErrorReason(metadata NWProtocolMetadata) *byte {
	result, callErr := tryNWQuicGetApplicationErrorReason(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetIdleTimeout func(options NWProtocolOptions) uint32
var _nWQuicGetIdleTimeoutErr error

func tryNWQuicGetIdleTimeout(options NWProtocolOptions) (uint32, error) {
	if _nWQuicGetIdleTimeout == nil {
		return 0, symbolCallError("nw_quic_get_idle_timeout", "12.0", _nWQuicGetIdleTimeoutErr)
	}
	return _nWQuicGetIdleTimeout(options), nil
}

// NWQuicGetIdleTimeout accesses the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_idle_timeout(_:)
func NWQuicGetIdleTimeout(options NWProtocolOptions) uint32 {
	result, callErr := tryNWQuicGetIdleTimeout(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxData func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxDataErr error

func tryNWQuicGetInitialMaxData(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxData == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_data", "12.0", _nWQuicGetInitialMaxDataErr)
	}
	return _nWQuicGetInitialMaxData(options), nil
}

// NWQuicGetInitialMaxData accesses a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_data(_:)
func NWQuicGetInitialMaxData(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxData(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxStreamDataBidirectionalLocal func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxStreamDataBidirectionalLocalErr error

func tryNWQuicGetInitialMaxStreamDataBidirectionalLocal(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxStreamDataBidirectionalLocal == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_bidirectional_local", "12.0", _nWQuicGetInitialMaxStreamDataBidirectionalLocalErr)
	}
	return _nWQuicGetInitialMaxStreamDataBidirectionalLocal(options), nil
}

// NWQuicGetInitialMaxStreamDataBidirectionalLocal accesses a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_local(_:)
func NWQuicGetInitialMaxStreamDataBidirectionalLocal(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxStreamDataBidirectionalLocal(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxStreamDataBidirectionalRemote func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxStreamDataBidirectionalRemoteErr error

func tryNWQuicGetInitialMaxStreamDataBidirectionalRemote(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxStreamDataBidirectionalRemote == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_bidirectional_remote", "12.0", _nWQuicGetInitialMaxStreamDataBidirectionalRemoteErr)
	}
	return _nWQuicGetInitialMaxStreamDataBidirectionalRemote(options), nil
}

// NWQuicGetInitialMaxStreamDataBidirectionalRemote accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_remote(_:)
func NWQuicGetInitialMaxStreamDataBidirectionalRemote(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxStreamDataBidirectionalRemote(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxStreamDataUnidirectional func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxStreamDataUnidirectionalErr error

func tryNWQuicGetInitialMaxStreamDataUnidirectional(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxStreamDataUnidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_stream_data_unidirectional", "12.0", _nWQuicGetInitialMaxStreamDataUnidirectionalErr)
	}
	return _nWQuicGetInitialMaxStreamDataUnidirectional(options), nil
}

// NWQuicGetInitialMaxStreamDataUnidirectional accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_unidirectional(_:)
func NWQuicGetInitialMaxStreamDataUnidirectional(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxStreamDataUnidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxStreamsBidirectional func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxStreamsBidirectionalErr error

func tryNWQuicGetInitialMaxStreamsBidirectional(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxStreamsBidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_streams_bidirectional", "12.0", _nWQuicGetInitialMaxStreamsBidirectionalErr)
	}
	return _nWQuicGetInitialMaxStreamsBidirectional(options), nil
}

// NWQuicGetInitialMaxStreamsBidirectional accesses a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_bidirectional(_:)
func NWQuicGetInitialMaxStreamsBidirectional(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxStreamsBidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetInitialMaxStreamsUnidirectional func(options NWProtocolOptions) uint64
var _nWQuicGetInitialMaxStreamsUnidirectionalErr error

func tryNWQuicGetInitialMaxStreamsUnidirectional(options NWProtocolOptions) (uint64, error) {
	if _nWQuicGetInitialMaxStreamsUnidirectional == nil {
		return 0, symbolCallError("nw_quic_get_initial_max_streams_unidirectional", "12.0", _nWQuicGetInitialMaxStreamsUnidirectionalErr)
	}
	return _nWQuicGetInitialMaxStreamsUnidirectional(options), nil
}

// NWQuicGetInitialMaxStreamsUnidirectional accesses a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_unidirectional(_:)
func NWQuicGetInitialMaxStreamsUnidirectional(options NWProtocolOptions) uint64 {
	result, callErr := tryNWQuicGetInitialMaxStreamsUnidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetKeepaliveInterval func(metadata NWProtocolMetadata) uint16
var _nWQuicGetKeepaliveIntervalErr error

func tryNWQuicGetKeepaliveInterval(metadata NWProtocolMetadata) (uint16, error) {
	if _nWQuicGetKeepaliveInterval == nil {
		return 0, symbolCallError("nw_quic_get_keepalive_interval", "12.0", _nWQuicGetKeepaliveIntervalErr)
	}
	return _nWQuicGetKeepaliveInterval(metadata), nil
}

// NWQuicGetKeepaliveInterval accesses the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_keepalive_interval(_:)
func NWQuicGetKeepaliveInterval(metadata NWProtocolMetadata) uint16 {
	result, callErr := tryNWQuicGetKeepaliveInterval(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetLocalMaxStreamsBidirectional func(metadata NWProtocolMetadata) uint64
var _nWQuicGetLocalMaxStreamsBidirectionalErr error

func tryNWQuicGetLocalMaxStreamsBidirectional(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetLocalMaxStreamsBidirectional == nil {
		return 0, symbolCallError("nw_quic_get_local_max_streams_bidirectional", "12.0", _nWQuicGetLocalMaxStreamsBidirectionalErr)
	}
	return _nWQuicGetLocalMaxStreamsBidirectional(metadata), nil
}

// NWQuicGetLocalMaxStreamsBidirectional accesses the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_bidirectional(_:)
func NWQuicGetLocalMaxStreamsBidirectional(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetLocalMaxStreamsBidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetLocalMaxStreamsUnidirectional func(metadata NWProtocolMetadata) uint64
var _nWQuicGetLocalMaxStreamsUnidirectionalErr error

func tryNWQuicGetLocalMaxStreamsUnidirectional(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetLocalMaxStreamsUnidirectional == nil {
		return 0, symbolCallError("nw_quic_get_local_max_streams_unidirectional", "12.0", _nWQuicGetLocalMaxStreamsUnidirectionalErr)
	}
	return _nWQuicGetLocalMaxStreamsUnidirectional(metadata), nil
}

// NWQuicGetLocalMaxStreamsUnidirectional accesses the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_unidirectional(_:)
func NWQuicGetLocalMaxStreamsUnidirectional(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetLocalMaxStreamsUnidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetMaxDatagramFrameSize func(options NWProtocolOptions) uint16
var _nWQuicGetMaxDatagramFrameSizeErr error

func tryNWQuicGetMaxDatagramFrameSize(options NWProtocolOptions) (uint16, error) {
	if _nWQuicGetMaxDatagramFrameSize == nil {
		return 0, symbolCallError("nw_quic_get_max_datagram_frame_size", "13.0", _nWQuicGetMaxDatagramFrameSizeErr)
	}
	return _nWQuicGetMaxDatagramFrameSize(options), nil
}

// NWQuicGetMaxDatagramFrameSize accesses a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_datagram_frame_size(_:)
func NWQuicGetMaxDatagramFrameSize(options NWProtocolOptions) uint16 {
	result, callErr := tryNWQuicGetMaxDatagramFrameSize(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetMaxUDPPayloadSize func(options NWProtocolOptions) uint16
var _nWQuicGetMaxUDPPayloadSizeErr error

func tryNWQuicGetMaxUDPPayloadSize(options NWProtocolOptions) (uint16, error) {
	if _nWQuicGetMaxUDPPayloadSize == nil {
		return 0, symbolCallError("nw_quic_get_max_udp_payload_size", "12.0", _nWQuicGetMaxUDPPayloadSizeErr)
	}
	return _nWQuicGetMaxUDPPayloadSize(options), nil
}

// NWQuicGetMaxUDPPayloadSize accesses the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_max_udp_payload_size(_:)
func NWQuicGetMaxUDPPayloadSize(options NWProtocolOptions) uint16 {
	result, callErr := tryNWQuicGetMaxUDPPayloadSize(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetRemoteIdleTimeout func(metadata NWProtocolMetadata) uint64
var _nWQuicGetRemoteIdleTimeoutErr error

func tryNWQuicGetRemoteIdleTimeout(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetRemoteIdleTimeout == nil {
		return 0, symbolCallError("nw_quic_get_remote_idle_timeout", "12.0", _nWQuicGetRemoteIdleTimeoutErr)
	}
	return _nWQuicGetRemoteIdleTimeout(metadata), nil
}

// NWQuicGetRemoteIdleTimeout accesses the idle timeout value from the peer’s transport parameters, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_idle_timeout(_:)
func NWQuicGetRemoteIdleTimeout(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetRemoteIdleTimeout(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetRemoteMaxStreamsBidirectional func(metadata NWProtocolMetadata) uint64
var _nWQuicGetRemoteMaxStreamsBidirectionalErr error

func tryNWQuicGetRemoteMaxStreamsBidirectional(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetRemoteMaxStreamsBidirectional == nil {
		return 0, symbolCallError("nw_quic_get_remote_max_streams_bidirectional", "12.0", _nWQuicGetRemoteMaxStreamsBidirectionalErr)
	}
	return _nWQuicGetRemoteMaxStreamsBidirectional(metadata), nil
}

// NWQuicGetRemoteMaxStreamsBidirectional accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_bidirectional(_:)
func NWQuicGetRemoteMaxStreamsBidirectional(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetRemoteMaxStreamsBidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetRemoteMaxStreamsUnidirectional func(metadata NWProtocolMetadata) uint64
var _nWQuicGetRemoteMaxStreamsUnidirectionalErr error

func tryNWQuicGetRemoteMaxStreamsUnidirectional(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetRemoteMaxStreamsUnidirectional == nil {
		return 0, symbolCallError("nw_quic_get_remote_max_streams_unidirectional", "12.0", _nWQuicGetRemoteMaxStreamsUnidirectionalErr)
	}
	return _nWQuicGetRemoteMaxStreamsUnidirectional(metadata), nil
}

// NWQuicGetRemoteMaxStreamsUnidirectional accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_unidirectional(_:)
func NWQuicGetRemoteMaxStreamsUnidirectional(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetRemoteMaxStreamsUnidirectional(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamApplicationError func(metadata NWProtocolMetadata) uint64
var _nWQuicGetStreamApplicationErrorErr error

func tryNWQuicGetStreamApplicationError(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetStreamApplicationError == nil {
		return 0, symbolCallError("nw_quic_get_stream_application_error", "12.0", _nWQuicGetStreamApplicationErrorErr)
	}
	return _nWQuicGetStreamApplicationError(metadata), nil
}

// NWQuicGetStreamApplicationError accesses the QUIC application error code received from the peer for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_application_error(_:)
func NWQuicGetStreamApplicationError(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetStreamApplicationError(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamID func(metadata NWProtocolMetadata) uint64
var _nWQuicGetStreamIDErr error

func tryNWQuicGetStreamID(metadata NWProtocolMetadata) (uint64, error) {
	if _nWQuicGetStreamID == nil {
		return 0, symbolCallError("nw_quic_get_stream_id", "12.0", _nWQuicGetStreamIDErr)
	}
	return _nWQuicGetStreamID(metadata), nil
}

// NWQuicGetStreamID accesses the QUIC stream identifier.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_id(_:)
func NWQuicGetStreamID(metadata NWProtocolMetadata) uint64 {
	result, callErr := tryNWQuicGetStreamID(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamIsDatagram func(options NWProtocolOptions) bool
var _nWQuicGetStreamIsDatagramErr error

func tryNWQuicGetStreamIsDatagram(options NWProtocolOptions) (bool, error) {
	if _nWQuicGetStreamIsDatagram == nil {
		return false, symbolCallError("nw_quic_get_stream_is_datagram", "13.0", _nWQuicGetStreamIsDatagramErr)
	}
	return _nWQuicGetStreamIsDatagram(options), nil
}

// NWQuicGetStreamIsDatagram checks if a QUIC stream is a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_datagram(_:)
func NWQuicGetStreamIsDatagram(options NWProtocolOptions) bool {
	result, callErr := tryNWQuicGetStreamIsDatagram(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamIsUnidirectional func(options NWProtocolOptions) bool
var _nWQuicGetStreamIsUnidirectionalErr error

func tryNWQuicGetStreamIsUnidirectional(options NWProtocolOptions) (bool, error) {
	if _nWQuicGetStreamIsUnidirectional == nil {
		return false, symbolCallError("nw_quic_get_stream_is_unidirectional", "12.0", _nWQuicGetStreamIsUnidirectionalErr)
	}
	return _nWQuicGetStreamIsUnidirectional(options), nil
}

// NWQuicGetStreamIsUnidirectional checks if a QUIC stream is unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_unidirectional(_:)
func NWQuicGetStreamIsUnidirectional(options NWProtocolOptions) bool {
	result, callErr := tryNWQuicGetStreamIsUnidirectional(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamType func(stream_metadata NWProtocolMetadata) uint8
var _nWQuicGetStreamTypeErr error

func tryNWQuicGetStreamType(stream_metadata NWProtocolMetadata) (uint8, error) {
	if _nWQuicGetStreamType == nil {
		return 0, symbolCallError("nw_quic_get_stream_type", "12.0", _nWQuicGetStreamTypeErr)
	}
	return _nWQuicGetStreamType(stream_metadata), nil
}

// NWQuicGetStreamType accesses the stream type of the QUIC stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_type(_:)
func NWQuicGetStreamType(stream_metadata NWProtocolMetadata) uint8 {
	result, callErr := tryNWQuicGetStreamType(stream_metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicGetStreamUsableDatagramFrameSize func(metadata NWProtocolMetadata) uint16
var _nWQuicGetStreamUsableDatagramFrameSizeErr error

func tryNWQuicGetStreamUsableDatagramFrameSize(metadata NWProtocolMetadata) (uint16, error) {
	if _nWQuicGetStreamUsableDatagramFrameSize == nil {
		return 0, symbolCallError("nw_quic_get_stream_usable_datagram_frame_size", "13.0", _nWQuicGetStreamUsableDatagramFrameSizeErr)
	}
	return _nWQuicGetStreamUsableDatagramFrameSize(metadata), nil
}

// NWQuicGetStreamUsableDatagramFrameSize accesses the maximum usable size of a datagram frame on a QUIC datagram flow.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_get_stream_usable_datagram_frame_size(_:)
func NWQuicGetStreamUsableDatagramFrameSize(metadata NWProtocolMetadata) uint16 {
	result, callErr := tryNWQuicGetStreamUsableDatagramFrameSize(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWQuicSetApplicationError func(metadata NWProtocolMetadata, application_error uint64, reason string)
var _nWQuicSetApplicationErrorErr error

func tryNWQuicSetApplicationError(metadata NWProtocolMetadata, application_error uint64, reason string) error {
	if _nWQuicSetApplicationError == nil {
		return symbolCallError("nw_quic_set_application_error", "12.0", _nWQuicSetApplicationErrorErr)
	}
	_nWQuicSetApplicationError(metadata, application_error, reason)
	return nil
}

// NWQuicSetApplicationError sets the QUIC application error code to send for the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_application_error(_:_:_:)
func NWQuicSetApplicationError(metadata NWProtocolMetadata, application_error uint64, reason string) {
	if callErr := tryNWQuicSetApplicationError(metadata, application_error, reason); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetIdleTimeout func(options NWProtocolOptions, idle_timeout uint32)
var _nWQuicSetIdleTimeoutErr error

func tryNWQuicSetIdleTimeout(options NWProtocolOptions, idle_timeout uint32) error {
	if _nWQuicSetIdleTimeout == nil {
		return symbolCallError("nw_quic_set_idle_timeout", "12.0", _nWQuicSetIdleTimeoutErr)
	}
	_nWQuicSetIdleTimeout(options, idle_timeout)
	return nil
}

// NWQuicSetIdleTimeout sets the idle timeout for the QUIC connection, in milliseconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_idle_timeout(_:_:)
func NWQuicSetIdleTimeout(options NWProtocolOptions, idle_timeout uint32) {
	if callErr := tryNWQuicSetIdleTimeout(options, idle_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxData func(options NWProtocolOptions, initial_max_data uint64)
var _nWQuicSetInitialMaxDataErr error

func tryNWQuicSetInitialMaxData(options NWProtocolOptions, initial_max_data uint64) error {
	if _nWQuicSetInitialMaxData == nil {
		return symbolCallError("nw_quic_set_initial_max_data", "12.0", _nWQuicSetInitialMaxDataErr)
	}
	_nWQuicSetInitialMaxData(options, initial_max_data)
	return nil
}

// NWQuicSetInitialMaxData sets a QUIC connection’s initial maximum data transport parameter.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_data(_:_:)
func NWQuicSetInitialMaxData(options NWProtocolOptions, initial_max_data uint64) {
	if callErr := tryNWQuicSetInitialMaxData(options, initial_max_data); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxStreamDataBidirectionalLocal func(options NWProtocolOptions, initial_max_stream_data_bidirectional_local uint64)
var _nWQuicSetInitialMaxStreamDataBidirectionalLocalErr error

func tryNWQuicSetInitialMaxStreamDataBidirectionalLocal(options NWProtocolOptions, initial_max_stream_data_bidirectional_local uint64) error {
	if _nWQuicSetInitialMaxStreamDataBidirectionalLocal == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_bidirectional_local", "12.0", _nWQuicSetInitialMaxStreamDataBidirectionalLocalErr)
	}
	_nWQuicSetInitialMaxStreamDataBidirectionalLocal(options, initial_max_stream_data_bidirectional_local)
	return nil
}

// NWQuicSetInitialMaxStreamDataBidirectionalLocal sets a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_local(_:_:)
func NWQuicSetInitialMaxStreamDataBidirectionalLocal(options NWProtocolOptions, initial_max_stream_data_bidirectional_local uint64) {
	if callErr := tryNWQuicSetInitialMaxStreamDataBidirectionalLocal(options, initial_max_stream_data_bidirectional_local); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxStreamDataBidirectionalRemote func(options NWProtocolOptions, initial_max_stream_data_bidirectional_remote uint64)
var _nWQuicSetInitialMaxStreamDataBidirectionalRemoteErr error

func tryNWQuicSetInitialMaxStreamDataBidirectionalRemote(options NWProtocolOptions, initial_max_stream_data_bidirectional_remote uint64) error {
	if _nWQuicSetInitialMaxStreamDataBidirectionalRemote == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_bidirectional_remote", "12.0", _nWQuicSetInitialMaxStreamDataBidirectionalRemoteErr)
	}
	_nWQuicSetInitialMaxStreamDataBidirectionalRemote(options, initial_max_stream_data_bidirectional_remote)
	return nil
}

// NWQuicSetInitialMaxStreamDataBidirectionalRemote sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_remote(_:_:)
func NWQuicSetInitialMaxStreamDataBidirectionalRemote(options NWProtocolOptions, initial_max_stream_data_bidirectional_remote uint64) {
	if callErr := tryNWQuicSetInitialMaxStreamDataBidirectionalRemote(options, initial_max_stream_data_bidirectional_remote); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxStreamDataUnidirectional func(options NWProtocolOptions, initial_max_stream_data_unidirectional uint64)
var _nWQuicSetInitialMaxStreamDataUnidirectionalErr error

func tryNWQuicSetInitialMaxStreamDataUnidirectional(options NWProtocolOptions, initial_max_stream_data_unidirectional uint64) error {
	if _nWQuicSetInitialMaxStreamDataUnidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_stream_data_unidirectional", "12.0", _nWQuicSetInitialMaxStreamDataUnidirectionalErr)
	}
	_nWQuicSetInitialMaxStreamDataUnidirectional(options, initial_max_stream_data_unidirectional)
	return nil
}

// NWQuicSetInitialMaxStreamDataUnidirectional sets a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_unidirectional(_:_:)
func NWQuicSetInitialMaxStreamDataUnidirectional(options NWProtocolOptions, initial_max_stream_data_unidirectional uint64) {
	if callErr := tryNWQuicSetInitialMaxStreamDataUnidirectional(options, initial_max_stream_data_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxStreamsBidirectional func(options NWProtocolOptions, initial_max_streams_bidirectional uint64)
var _nWQuicSetInitialMaxStreamsBidirectionalErr error

func tryNWQuicSetInitialMaxStreamsBidirectional(options NWProtocolOptions, initial_max_streams_bidirectional uint64) error {
	if _nWQuicSetInitialMaxStreamsBidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_streams_bidirectional", "12.0", _nWQuicSetInitialMaxStreamsBidirectionalErr)
	}
	_nWQuicSetInitialMaxStreamsBidirectional(options, initial_max_streams_bidirectional)
	return nil
}

// NWQuicSetInitialMaxStreamsBidirectional sets a QUIC connection’s initial maximum number of bidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_bidirectional(_:_:)
func NWQuicSetInitialMaxStreamsBidirectional(options NWProtocolOptions, initial_max_streams_bidirectional uint64) {
	if callErr := tryNWQuicSetInitialMaxStreamsBidirectional(options, initial_max_streams_bidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetInitialMaxStreamsUnidirectional func(options NWProtocolOptions, initial_max_streams_unidirectional uint64)
var _nWQuicSetInitialMaxStreamsUnidirectionalErr error

func tryNWQuicSetInitialMaxStreamsUnidirectional(options NWProtocolOptions, initial_max_streams_unidirectional uint64) error {
	if _nWQuicSetInitialMaxStreamsUnidirectional == nil {
		return symbolCallError("nw_quic_set_initial_max_streams_unidirectional", "12.0", _nWQuicSetInitialMaxStreamsUnidirectionalErr)
	}
	_nWQuicSetInitialMaxStreamsUnidirectional(options, initial_max_streams_unidirectional)
	return nil
}

// NWQuicSetInitialMaxStreamsUnidirectional sets a QUIC connection’s initial maximum number of unidirectional streams.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_unidirectional(_:_:)
func NWQuicSetInitialMaxStreamsUnidirectional(options NWProtocolOptions, initial_max_streams_unidirectional uint64) {
	if callErr := tryNWQuicSetInitialMaxStreamsUnidirectional(options, initial_max_streams_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetKeepaliveInterval func(metadata NWProtocolMetadata, keepalive_interval uint16)
var _nWQuicSetKeepaliveIntervalErr error

func tryNWQuicSetKeepaliveInterval(metadata NWProtocolMetadata, keepalive_interval uint16) error {
	if _nWQuicSetKeepaliveInterval == nil {
		return symbolCallError("nw_quic_set_keepalive_interval", "12.0", _nWQuicSetKeepaliveIntervalErr)
	}
	_nWQuicSetKeepaliveInterval(metadata, keepalive_interval)
	return nil
}

// NWQuicSetKeepaliveInterval sets the keepalive interval for the QUIC connection, in seconds.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_keepalive_interval(_:_:)
func NWQuicSetKeepaliveInterval(metadata NWProtocolMetadata, keepalive_interval uint16) {
	if callErr := tryNWQuicSetKeepaliveInterval(metadata, keepalive_interval); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetLocalMaxStreamsBidirectional func(metadata NWProtocolMetadata, max_streams_bidirectional uint64)
var _nWQuicSetLocalMaxStreamsBidirectionalErr error

func tryNWQuicSetLocalMaxStreamsBidirectional(metadata NWProtocolMetadata, max_streams_bidirectional uint64) error {
	if _nWQuicSetLocalMaxStreamsBidirectional == nil {
		return symbolCallError("nw_quic_set_local_max_streams_bidirectional", "12.0", _nWQuicSetLocalMaxStreamsBidirectionalErr)
	}
	_nWQuicSetLocalMaxStreamsBidirectional(metadata, max_streams_bidirectional)
	return nil
}

// NWQuicSetLocalMaxStreamsBidirectional sets the maximum number of bidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_bidirectional(_:_:)
func NWQuicSetLocalMaxStreamsBidirectional(metadata NWProtocolMetadata, max_streams_bidirectional uint64) {
	if callErr := tryNWQuicSetLocalMaxStreamsBidirectional(metadata, max_streams_bidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetLocalMaxStreamsUnidirectional func(metadata NWProtocolMetadata, max_streams_unidirectional uint64)
var _nWQuicSetLocalMaxStreamsUnidirectionalErr error

func tryNWQuicSetLocalMaxStreamsUnidirectional(metadata NWProtocolMetadata, max_streams_unidirectional uint64) error {
	if _nWQuicSetLocalMaxStreamsUnidirectional == nil {
		return symbolCallError("nw_quic_set_local_max_streams_unidirectional", "12.0", _nWQuicSetLocalMaxStreamsUnidirectionalErr)
	}
	_nWQuicSetLocalMaxStreamsUnidirectional(metadata, max_streams_unidirectional)
	return nil
}

// NWQuicSetLocalMaxStreamsUnidirectional sets the maximum number of unidirectional streams that the peer can create on a QUIC connection.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_unidirectional(_:_:)
func NWQuicSetLocalMaxStreamsUnidirectional(metadata NWProtocolMetadata, max_streams_unidirectional uint64) {
	if callErr := tryNWQuicSetLocalMaxStreamsUnidirectional(metadata, max_streams_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetMaxDatagramFrameSize func(options NWProtocolOptions, max_datagram_frame_size uint16)
var _nWQuicSetMaxDatagramFrameSizeErr error

func tryNWQuicSetMaxDatagramFrameSize(options NWProtocolOptions, max_datagram_frame_size uint16) error {
	if _nWQuicSetMaxDatagramFrameSize == nil {
		return symbolCallError("nw_quic_set_max_datagram_frame_size", "13.0", _nWQuicSetMaxDatagramFrameSizeErr)
	}
	_nWQuicSetMaxDatagramFrameSize(options, max_datagram_frame_size)
	return nil
}

// NWQuicSetMaxDatagramFrameSize sets a QUIC connection’s maximum DATAGRAM frame size.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_datagram_frame_size(_:_:)
func NWQuicSetMaxDatagramFrameSize(options NWProtocolOptions, max_datagram_frame_size uint16) {
	if callErr := tryNWQuicSetMaxDatagramFrameSize(options, max_datagram_frame_size); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetMaxUDPPayloadSize func(options NWProtocolOptions, max_udp_payload_size uint16)
var _nWQuicSetMaxUDPPayloadSizeErr error

func tryNWQuicSetMaxUDPPayloadSize(options NWProtocolOptions, max_udp_payload_size uint16) error {
	if _nWQuicSetMaxUDPPayloadSize == nil {
		return symbolCallError("nw_quic_set_max_udp_payload_size", "12.0", _nWQuicSetMaxUDPPayloadSizeErr)
	}
	_nWQuicSetMaxUDPPayloadSize(options, max_udp_payload_size)
	return nil
}

// NWQuicSetMaxUDPPayloadSize sets the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_max_udp_payload_size(_:_:)
func NWQuicSetMaxUDPPayloadSize(options NWProtocolOptions, max_udp_payload_size uint16) {
	if callErr := tryNWQuicSetMaxUDPPayloadSize(options, max_udp_payload_size); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetStreamApplicationError func(metadata NWProtocolMetadata, application_error uint64)
var _nWQuicSetStreamApplicationErrorErr error

func tryNWQuicSetStreamApplicationError(metadata NWProtocolMetadata, application_error uint64) error {
	if _nWQuicSetStreamApplicationError == nil {
		return symbolCallError("nw_quic_set_stream_application_error", "12.0", _nWQuicSetStreamApplicationErrorErr)
	}
	_nWQuicSetStreamApplicationError(metadata, application_error)
	return nil
}

// NWQuicSetStreamApplicationError sets the QUIC application error code to send for the stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_application_error(_:_:)
func NWQuicSetStreamApplicationError(metadata NWProtocolMetadata, application_error uint64) {
	if callErr := tryNWQuicSetStreamApplicationError(metadata, application_error); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetStreamIsDatagram func(options NWProtocolOptions, is_datagram bool)
var _nWQuicSetStreamIsDatagramErr error

func tryNWQuicSetStreamIsDatagram(options NWProtocolOptions, is_datagram bool) error {
	if _nWQuicSetStreamIsDatagram == nil {
		return symbolCallError("nw_quic_set_stream_is_datagram", "13.0", _nWQuicSetStreamIsDatagramErr)
	}
	_nWQuicSetStreamIsDatagram(options, is_datagram)
	return nil
}

// NWQuicSetStreamIsDatagram configures a QUIC stream as a datagram flow, instead of a byte stream.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_datagram(_:_:)
func NWQuicSetStreamIsDatagram(options NWProtocolOptions, is_datagram bool) {
	if callErr := tryNWQuicSetStreamIsDatagram(options, is_datagram); callErr != nil {
		panic(callErr)
	}
}

var _nWQuicSetStreamIsUnidirectional func(options NWProtocolOptions, is_unidirectional bool)
var _nWQuicSetStreamIsUnidirectionalErr error

func tryNWQuicSetStreamIsUnidirectional(options NWProtocolOptions, is_unidirectional bool) error {
	if _nWQuicSetStreamIsUnidirectional == nil {
		return symbolCallError("nw_quic_set_stream_is_unidirectional", "12.0", _nWQuicSetStreamIsUnidirectionalErr)
	}
	_nWQuicSetStreamIsUnidirectional(options, is_unidirectional)
	return nil
}

// NWQuicSetStreamIsUnidirectional configures a QUIC stream as unidirectional, instead of bidirectional.
//
// See: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_unidirectional(_:_:)
func NWQuicSetStreamIsUnidirectional(options NWProtocolOptions, is_unidirectional bool) {
	if callErr := tryNWQuicSetStreamIsUnidirectional(options, is_unidirectional); callErr != nil {
		panic(callErr)
	}
}

var _nWRelayHopAddAdditionalHttpHeaderField func(relay_hop NWRelayHop, field_name string, field_value string)
var _nWRelayHopAddAdditionalHttpHeaderFieldErr error

func tryNWRelayHopAddAdditionalHttpHeaderField(relay_hop NWRelayHop, field_name string, field_value string) error {
	if _nWRelayHopAddAdditionalHttpHeaderField == nil {
		return symbolCallError("nw_relay_hop_add_additional_http_header_field", "14.0", _nWRelayHopAddAdditionalHttpHeaderFieldErr)
	}
	_nWRelayHopAddAdditionalHttpHeaderField(relay_hop, field_name, field_value)
	return nil
}

// NWRelayHopAddAdditionalHttpHeaderField adds an HTTP header name and value pair to send as part of [CONNECT] requests to the relay.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_add_additional_http_header_field(_:_:_:)
func NWRelayHopAddAdditionalHttpHeaderField(relay_hop NWRelayHop, field_name string, field_value string) {
	if callErr := tryNWRelayHopAddAdditionalHttpHeaderField(relay_hop, field_name, field_value); callErr != nil {
		panic(callErr)
	}
}

var _nWRelayHopCreate func(http3_relay_endpoint NWEndpoint, http2_relay_endpoint NWEndpoint, relay_tls_options NWProtocolOptions) NWRelayHop
var _nWRelayHopCreateErr error

func tryNWRelayHopCreate(http3_relay_endpoint NWEndpoint, http2_relay_endpoint NWEndpoint, relay_tls_options NWProtocolOptions) (NWRelayHop, error) {
	if _nWRelayHopCreate == nil {
		return *new(NWRelayHop), symbolCallError("nw_relay_hop_create", "14.0", _nWRelayHopCreateErr)
	}
	return _nWRelayHopCreate(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options), nil
}

// NWRelayHopCreate creates a configuration for a secure relay accessible using HTTP/3, with an optional HTTP/2 fallback.
//
// See: https://developer.apple.com/documentation/Network/nw_relay_hop_create(_:_:_:)
func NWRelayHopCreate(http3_relay_endpoint NWEndpoint, http2_relay_endpoint NWEndpoint, relay_tls_options NWProtocolOptions) NWRelayHop {
	result, callErr := tryNWRelayHopCreate(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWRelease func(obj unsafe.Pointer)
var _nWReleaseErr error

func tryNWRelease(obj unsafe.Pointer) error {
	if _nWRelease == nil {
		return symbolCallError("nw_release", "10.14", _nWReleaseErr)
	}
	_nWRelease(obj)
	return nil
}

// NWRelease releases a reference count on a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_release
func NWRelease(obj unsafe.Pointer) {
	if callErr := tryNWRelease(obj); callErr != nil {
		panic(callErr)
	}
}

var _nWResolutionReportCopyPreferredEndpoint func(resolution_report NWResolutionReport) NWEndpoint
var _nWResolutionReportCopyPreferredEndpointErr error

func tryNWResolutionReportCopyPreferredEndpoint(resolution_report NWResolutionReport) (NWEndpoint, error) {
	if _nWResolutionReportCopyPreferredEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_resolution_report_copy_preferred_endpoint", "11.0", _nWResolutionReportCopyPreferredEndpointErr)
	}
	return _nWResolutionReportCopyPreferredEndpoint(resolution_report), nil
}

// NWResolutionReportCopyPreferredEndpoint accesses the resolved endpoint that the connection used for its first connection attempt.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_preferred_endpoint(_:)
func NWResolutionReportCopyPreferredEndpoint(resolution_report NWResolutionReport) NWEndpoint {
	result, callErr := tryNWResolutionReportCopyPreferredEndpoint(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolutionReportCopySuccessfulEndpoint func(resolution_report NWResolutionReport) NWEndpoint
var _nWResolutionReportCopySuccessfulEndpointErr error

func tryNWResolutionReportCopySuccessfulEndpoint(resolution_report NWResolutionReport) (NWEndpoint, error) {
	if _nWResolutionReportCopySuccessfulEndpoint == nil {
		return NWEndpoint{}, symbolCallError("nw_resolution_report_copy_successful_endpoint", "11.0", _nWResolutionReportCopySuccessfulEndpointErr)
	}
	return _nWResolutionReportCopySuccessfulEndpoint(resolution_report), nil
}

// NWResolutionReportCopySuccessfulEndpoint accesses the resolved endpoint that led to the established connection.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_successful_endpoint(_:)
func NWResolutionReportCopySuccessfulEndpoint(resolution_report NWResolutionReport) NWEndpoint {
	result, callErr := tryNWResolutionReportCopySuccessfulEndpoint(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolutionReportGetEndpointCount func(resolution_report NWResolutionReport) uint32
var _nWResolutionReportGetEndpointCountErr error

func tryNWResolutionReportGetEndpointCount(resolution_report NWResolutionReport) (uint32, error) {
	if _nWResolutionReportGetEndpointCount == nil {
		return 0, symbolCallError("nw_resolution_report_get_endpoint_count", "11.0", _nWResolutionReportGetEndpointCountErr)
	}
	return _nWResolutionReportGetEndpointCount(resolution_report), nil
}

// NWResolutionReportGetEndpointCount accesses the number of endpoints resolved in this step.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_endpoint_count(_:)
func NWResolutionReportGetEndpointCount(resolution_report NWResolutionReport) uint32 {
	result, callErr := tryNWResolutionReportGetEndpointCount(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolutionReportGetMilliseconds func(resolution_report NWResolutionReport) uint64
var _nWResolutionReportGetMillisecondsErr error

func tryNWResolutionReportGetMilliseconds(resolution_report NWResolutionReport) (uint64, error) {
	if _nWResolutionReportGetMilliseconds == nil {
		return 0, symbolCallError("nw_resolution_report_get_milliseconds", "11.0", _nWResolutionReportGetMillisecondsErr)
	}
	return _nWResolutionReportGetMilliseconds(resolution_report), nil
}

// NWResolutionReportGetMilliseconds accesses the duration of this resolution step, from when the query was issued to when the response was complete.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_milliseconds(_:)
func NWResolutionReportGetMilliseconds(resolution_report NWResolutionReport) uint64 {
	result, callErr := tryNWResolutionReportGetMilliseconds(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolutionReportGetProtocol func(resolution_report NWResolutionReport) NWReportResolutionProtocol
var _nWResolutionReportGetProtocolErr error

func tryNWResolutionReportGetProtocol(resolution_report NWResolutionReport) (NWReportResolutionProtocol, error) {
	if _nWResolutionReportGetProtocol == nil {
		return *new(NWReportResolutionProtocol), symbolCallError("nw_resolution_report_get_protocol", "11.0", _nWResolutionReportGetProtocolErr)
	}
	return _nWResolutionReportGetProtocol(resolution_report), nil
}

// NWResolutionReportGetProtocol accesses the transport protocol your connection used for DNS resolution.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_protocol(_:)
func NWResolutionReportGetProtocol(resolution_report NWResolutionReport) NWReportResolutionProtocol {
	result, callErr := tryNWResolutionReportGetProtocol(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolutionReportGetSource func(resolution_report NWResolutionReport) NWReportResolutionSource
var _nWResolutionReportGetSourceErr error

func tryNWResolutionReportGetSource(resolution_report NWResolutionReport) (NWReportResolutionSource, error) {
	if _nWResolutionReportGetSource == nil {
		return *new(NWReportResolutionSource), symbolCallError("nw_resolution_report_get_source", "11.0", _nWResolutionReportGetSourceErr)
	}
	return _nWResolutionReportGetSource(resolution_report), nil
}

// NWResolutionReportGetSource accesses the source of the DNS response.
//
// See: https://developer.apple.com/documentation/Network/nw_resolution_report_get_source(_:)
func NWResolutionReportGetSource(resolution_report NWResolutionReport) NWReportResolutionSource {
	result, callErr := tryNWResolutionReportGetSource(resolution_report)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolverConfigAddServerAddress func(config NWResolverConfig, server_address NWEndpoint)
var _nWResolverConfigAddServerAddressErr error

func tryNWResolverConfigAddServerAddress(config NWResolverConfig, server_address NWEndpoint) error {
	if _nWResolverConfigAddServerAddress == nil {
		return symbolCallError("nw_resolver_config_add_server_address", "11.0", _nWResolverConfigAddServerAddressErr)
	}
	_nWResolverConfigAddServerAddress(config, server_address)
	return nil
}

// NWResolverConfigAddServerAddress provides a well-known DNS server address to use instead of looking up the address dynamically.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_add_server_address(_:_:)
func NWResolverConfigAddServerAddress(config NWResolverConfig, server_address NWEndpoint) {
	if callErr := tryNWResolverConfigAddServerAddress(config, server_address); callErr != nil {
		panic(callErr)
	}
}

var _nWResolverConfigCreateHttps func(url_endpoint NWEndpoint) NWResolverConfig
var _nWResolverConfigCreateHttpsErr error

func tryNWResolverConfigCreateHttps(url_endpoint NWEndpoint) (NWResolverConfig, error) {
	if _nWResolverConfigCreateHttps == nil {
		return *new(NWResolverConfig), symbolCallError("nw_resolver_config_create_https", "11.0", _nWResolverConfigCreateHttpsErr)
	}
	return _nWResolverConfigCreateHttps(url_endpoint), nil
}

// NWResolverConfigCreateHttps initializes a DNS-over-HTTPS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_https(_:)
func NWResolverConfigCreateHttps(url_endpoint NWEndpoint) NWResolverConfig {
	result, callErr := tryNWResolverConfigCreateHttps(url_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWResolverConfigCreateTLS func(server_endpoint NWEndpoint) NWResolverConfig
var _nWResolverConfigCreateTLSErr error

func tryNWResolverConfigCreateTLS(server_endpoint NWEndpoint) (NWResolverConfig, error) {
	if _nWResolverConfigCreateTLS == nil {
		return *new(NWResolverConfig), symbolCallError("nw_resolver_config_create_tls", "11.0", _nWResolverConfigCreateTLSErr)
	}
	return _nWResolverConfigCreateTLS(server_endpoint), nil
}

// NWResolverConfigCreateTLS initializes a DNS-over-TLS resolver configuration.
//
// See: https://developer.apple.com/documentation/Network/nw_resolver_config_create_tls(_:)
func NWResolverConfigCreateTLS(server_endpoint NWEndpoint) NWResolverConfig {
	result, callErr := tryNWResolverConfigCreateTLS(server_endpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWRetain func(obj unsafe.Pointer) unsafe.Pointer
var _nWRetainErr error

func tryNWRetain(obj unsafe.Pointer) (unsafe.Pointer, error) {
	if _nWRetain == nil {
		return nil, symbolCallError("nw_retain", "10.14", _nWRetainErr)
	}
	return _nWRetain(obj), nil
}

// NWRetain adds a reference count to a Network.framework object.
//
// See: https://developer.apple.com/documentation/Network/nw_retain
func NWRetain(obj unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNWRetain(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTCPCreateOptions func() NWProtocolOptions
var _nWTCPCreateOptionsErr error

func tryNWTCPCreateOptions() (NWProtocolOptions, error) {
	if _nWTCPCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_tcp_create_options", "10.14", _nWTCPCreateOptionsErr)
	}
	return _nWTCPCreateOptions(), nil
}

// NWTCPCreateOptions initializes a default set of TCP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_create_options()
func NWTCPCreateOptions() NWProtocolOptions {
	result, callErr := tryNWTCPCreateOptions()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTCPGetAvailableReceiveBuffer func(metadata NWProtocolMetadata) uint32
var _nWTCPGetAvailableReceiveBufferErr error

func tryNWTCPGetAvailableReceiveBuffer(metadata NWProtocolMetadata) (uint32, error) {
	if _nWTCPGetAvailableReceiveBuffer == nil {
		return 0, symbolCallError("nw_tcp_get_available_receive_buffer", "10.14", _nWTCPGetAvailableReceiveBufferErr)
	}
	return _nWTCPGetAvailableReceiveBuffer(metadata), nil
}

// NWTCPGetAvailableReceiveBuffer accesses the number of available bytes in the TCP receive buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_receive_buffer(_:)
func NWTCPGetAvailableReceiveBuffer(metadata NWProtocolMetadata) uint32 {
	result, callErr := tryNWTCPGetAvailableReceiveBuffer(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTCPGetAvailableSendBuffer func(metadata NWProtocolMetadata) uint32
var _nWTCPGetAvailableSendBufferErr error

func tryNWTCPGetAvailableSendBuffer(metadata NWProtocolMetadata) (uint32, error) {
	if _nWTCPGetAvailableSendBuffer == nil {
		return 0, symbolCallError("nw_tcp_get_available_send_buffer", "10.14", _nWTCPGetAvailableSendBufferErr)
	}
	return _nWTCPGetAvailableSendBuffer(metadata), nil
}

// NWTCPGetAvailableSendBuffer accesses the number of available bytes in the TCP send buffer.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_get_available_send_buffer(_:)
func NWTCPGetAvailableSendBuffer(metadata NWProtocolMetadata) uint32 {
	result, callErr := tryNWTCPGetAvailableSendBuffer(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTCPOptionsSetConnectionTimeout func(options NWProtocolOptions, connection_timeout uint32)
var _nWTCPOptionsSetConnectionTimeoutErr error

func tryNWTCPOptionsSetConnectionTimeout(options NWProtocolOptions, connection_timeout uint32) error {
	if _nWTCPOptionsSetConnectionTimeout == nil {
		return symbolCallError("nw_tcp_options_set_connection_timeout", "10.14", _nWTCPOptionsSetConnectionTimeoutErr)
	}
	_nWTCPOptionsSetConnectionTimeout(options, connection_timeout)
	return nil
}

// NWTCPOptionsSetConnectionTimeout sets the number of seconds that TCP waits before timing out its handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_connection_timeout(_:_:)
func NWTCPOptionsSetConnectionTimeout(options NWProtocolOptions, connection_timeout uint32) {
	if callErr := tryNWTCPOptionsSetConnectionTimeout(options, connection_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetDisableAckStretching func(options NWProtocolOptions, disable_ack_stretching bool)
var _nWTCPOptionsSetDisableAckStretchingErr error

func tryNWTCPOptionsSetDisableAckStretching(options NWProtocolOptions, disable_ack_stretching bool) error {
	if _nWTCPOptionsSetDisableAckStretching == nil {
		return symbolCallError("nw_tcp_options_set_disable_ack_stretching", "10.14", _nWTCPOptionsSetDisableAckStretchingErr)
	}
	_nWTCPOptionsSetDisableAckStretching(options, disable_ack_stretching)
	return nil
}

// NWTCPOptionsSetDisableAckStretching disables TCP acknowledgment stretching.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ack_stretching(_:_:)
func NWTCPOptionsSetDisableAckStretching(options NWProtocolOptions, disable_ack_stretching bool) {
	if callErr := tryNWTCPOptionsSetDisableAckStretching(options, disable_ack_stretching); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetDisableEcn func(options NWProtocolOptions, disable_ecn bool)
var _nWTCPOptionsSetDisableEcnErr error

func tryNWTCPOptionsSetDisableEcn(options NWProtocolOptions, disable_ecn bool) error {
	if _nWTCPOptionsSetDisableEcn == nil {
		return symbolCallError("nw_tcp_options_set_disable_ecn", "10.14", _nWTCPOptionsSetDisableEcnErr)
	}
	_nWTCPOptionsSetDisableEcn(options, disable_ecn)
	return nil
}

// NWTCPOptionsSetDisableEcn disables negotiation of Explicit Congestion Notification markings.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ecn(_:_:)
func NWTCPOptionsSetDisableEcn(options NWProtocolOptions, disable_ecn bool) {
	if callErr := tryNWTCPOptionsSetDisableEcn(options, disable_ecn); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetEnableFastOpen func(options NWProtocolOptions, enable_fast_open bool)
var _nWTCPOptionsSetEnableFastOpenErr error

func tryNWTCPOptionsSetEnableFastOpen(options NWProtocolOptions, enable_fast_open bool) error {
	if _nWTCPOptionsSetEnableFastOpen == nil {
		return symbolCallError("nw_tcp_options_set_enable_fast_open", "10.14", _nWTCPOptionsSetEnableFastOpenErr)
	}
	_nWTCPOptionsSetEnableFastOpen(options, enable_fast_open)
	return nil
}

// NWTCPOptionsSetEnableFastOpen enables TCP Fast Open on a connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_fast_open(_:_:)
func NWTCPOptionsSetEnableFastOpen(options NWProtocolOptions, enable_fast_open bool) {
	if callErr := tryNWTCPOptionsSetEnableFastOpen(options, enable_fast_open); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetEnableKeepalive func(options NWProtocolOptions, enable_keepalive bool)
var _nWTCPOptionsSetEnableKeepaliveErr error

func tryNWTCPOptionsSetEnableKeepalive(options NWProtocolOptions, enable_keepalive bool) error {
	if _nWTCPOptionsSetEnableKeepalive == nil {
		return symbolCallError("nw_tcp_options_set_enable_keepalive", "10.14", _nWTCPOptionsSetEnableKeepaliveErr)
	}
	_nWTCPOptionsSetEnableKeepalive(options, enable_keepalive)
	return nil
}

// NWTCPOptionsSetEnableKeepalive enables TCP keepalives.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_keepalive(_:_:)
func NWTCPOptionsSetEnableKeepalive(options NWProtocolOptions, enable_keepalive bool) {
	if callErr := tryNWTCPOptionsSetEnableKeepalive(options, enable_keepalive); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetKeepaliveCount func(options NWProtocolOptions, keepalive_count uint32)
var _nWTCPOptionsSetKeepaliveCountErr error

func tryNWTCPOptionsSetKeepaliveCount(options NWProtocolOptions, keepalive_count uint32) error {
	if _nWTCPOptionsSetKeepaliveCount == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_count", "10.14", _nWTCPOptionsSetKeepaliveCountErr)
	}
	_nWTCPOptionsSetKeepaliveCount(options, keepalive_count)
	return nil
}

// NWTCPOptionsSetKeepaliveCount sets the number of keepalive probes that TCP sends before terminating the connection.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_count(_:_:)
func NWTCPOptionsSetKeepaliveCount(options NWProtocolOptions, keepalive_count uint32) {
	if callErr := tryNWTCPOptionsSetKeepaliveCount(options, keepalive_count); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetKeepaliveIdleTime func(options NWProtocolOptions, keepalive_idle_time uint32)
var _nWTCPOptionsSetKeepaliveIdleTimeErr error

func tryNWTCPOptionsSetKeepaliveIdleTime(options NWProtocolOptions, keepalive_idle_time uint32) error {
	if _nWTCPOptionsSetKeepaliveIdleTime == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_idle_time", "10.14", _nWTCPOptionsSetKeepaliveIdleTimeErr)
	}
	_nWTCPOptionsSetKeepaliveIdleTime(options, keepalive_idle_time)
	return nil
}

// NWTCPOptionsSetKeepaliveIdleTime sets the number of seconds of idleness that TCP waits before sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_idle_time(_:_:)
func NWTCPOptionsSetKeepaliveIdleTime(options NWProtocolOptions, keepalive_idle_time uint32) {
	if callErr := tryNWTCPOptionsSetKeepaliveIdleTime(options, keepalive_idle_time); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetKeepaliveInterval func(options NWProtocolOptions, keepalive_interval uint32)
var _nWTCPOptionsSetKeepaliveIntervalErr error

func tryNWTCPOptionsSetKeepaliveInterval(options NWProtocolOptions, keepalive_interval uint32) error {
	if _nWTCPOptionsSetKeepaliveInterval == nil {
		return symbolCallError("nw_tcp_options_set_keepalive_interval", "10.14", _nWTCPOptionsSetKeepaliveIntervalErr)
	}
	_nWTCPOptionsSetKeepaliveInterval(options, keepalive_interval)
	return nil
}

// NWTCPOptionsSetKeepaliveInterval sets the number of seconds that TCP waits between sending keepalive probes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_interval(_:_:)
func NWTCPOptionsSetKeepaliveInterval(options NWProtocolOptions, keepalive_interval uint32) {
	if callErr := tryNWTCPOptionsSetKeepaliveInterval(options, keepalive_interval); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetMaximumSegmentSize func(options NWProtocolOptions, maximum_segment_size uint32)
var _nWTCPOptionsSetMaximumSegmentSizeErr error

func tryNWTCPOptionsSetMaximumSegmentSize(options NWProtocolOptions, maximum_segment_size uint32) error {
	if _nWTCPOptionsSetMaximumSegmentSize == nil {
		return symbolCallError("nw_tcp_options_set_maximum_segment_size", "10.14", _nWTCPOptionsSetMaximumSegmentSizeErr)
	}
	_nWTCPOptionsSetMaximumSegmentSize(options, maximum_segment_size)
	return nil
}

// NWTCPOptionsSetMaximumSegmentSize sets TCP’s maximum segment size in bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_maximum_segment_size(_:_:)
func NWTCPOptionsSetMaximumSegmentSize(options NWProtocolOptions, maximum_segment_size uint32) {
	if callErr := tryNWTCPOptionsSetMaximumSegmentSize(options, maximum_segment_size); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetMultipathForceVersion func(options NWProtocolOptions, multipath_force_version NWMultipathVersion)
var _nWTCPOptionsSetMultipathForceVersionErr error

func tryNWTCPOptionsSetMultipathForceVersion(options NWProtocolOptions, multipath_force_version NWMultipathVersion) error {
	if _nWTCPOptionsSetMultipathForceVersion == nil {
		return symbolCallError("nw_tcp_options_set_multipath_force_version", "12.0", _nWTCPOptionsSetMultipathForceVersionErr)
	}
	_nWTCPOptionsSetMultipathForceVersion(options, multipath_force_version)
	return nil
}

// NWTCPOptionsSetMultipathForceVersion.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_multipath_force_version(_:_:)
func NWTCPOptionsSetMultipathForceVersion(options NWProtocolOptions, multipath_force_version NWMultipathVersion) {
	if callErr := tryNWTCPOptionsSetMultipathForceVersion(options, multipath_force_version); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetNoDelay func(options NWProtocolOptions, no_delay bool)
var _nWTCPOptionsSetNoDelayErr error

func tryNWTCPOptionsSetNoDelay(options NWProtocolOptions, no_delay bool) error {
	if _nWTCPOptionsSetNoDelay == nil {
		return symbolCallError("nw_tcp_options_set_no_delay", "10.14", _nWTCPOptionsSetNoDelayErr)
	}
	_nWTCPOptionsSetNoDelay(options, no_delay)
	return nil
}

// NWTCPOptionsSetNoDelay disables Nagle’s algorithm for TCP.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_delay(_:_:)
func NWTCPOptionsSetNoDelay(options NWProtocolOptions, no_delay bool) {
	if callErr := tryNWTCPOptionsSetNoDelay(options, no_delay); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetNoOptions func(options NWProtocolOptions, no_options bool)
var _nWTCPOptionsSetNoOptionsErr error

func tryNWTCPOptionsSetNoOptions(options NWProtocolOptions, no_options bool) error {
	if _nWTCPOptionsSetNoOptions == nil {
		return symbolCallError("nw_tcp_options_set_no_options", "10.14", _nWTCPOptionsSetNoOptionsErr)
	}
	_nWTCPOptionsSetNoOptions(options, no_options)
	return nil
}

// NWTCPOptionsSetNoOptions sets TCP into no-options mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_options(_:_:)
func NWTCPOptionsSetNoOptions(options NWProtocolOptions, no_options bool) {
	if callErr := tryNWTCPOptionsSetNoOptions(options, no_options); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetNoPush func(options NWProtocolOptions, no_push bool)
var _nWTCPOptionsSetNoPushErr error

func tryNWTCPOptionsSetNoPush(options NWProtocolOptions, no_push bool) error {
	if _nWTCPOptionsSetNoPush == nil {
		return symbolCallError("nw_tcp_options_set_no_push", "10.14", _nWTCPOptionsSetNoPushErr)
	}
	_nWTCPOptionsSetNoPush(options, no_push)
	return nil
}

// NWTCPOptionsSetNoPush sets TCP into no-push mode.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_push(_:_:)
func NWTCPOptionsSetNoPush(options NWProtocolOptions, no_push bool) {
	if callErr := tryNWTCPOptionsSetNoPush(options, no_push); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetPersistTimeout func(options NWProtocolOptions, persist_timeout uint32)
var _nWTCPOptionsSetPersistTimeoutErr error

func tryNWTCPOptionsSetPersistTimeout(options NWProtocolOptions, persist_timeout uint32) error {
	if _nWTCPOptionsSetPersistTimeout == nil {
		return symbolCallError("nw_tcp_options_set_persist_timeout", "10.14", _nWTCPOptionsSetPersistTimeoutErr)
	}
	_nWTCPOptionsSetPersistTimeout(options, persist_timeout)
	return nil
}

// NWTCPOptionsSetPersistTimeout sets the TCP persist timeout in seconds, as defined by RFC 6429.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_persist_timeout(_:_:)
func NWTCPOptionsSetPersistTimeout(options NWProtocolOptions, persist_timeout uint32) {
	if callErr := tryNWTCPOptionsSetPersistTimeout(options, persist_timeout); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetRetransmitConnectionDropTime func(options NWProtocolOptions, retransmit_connection_drop_time uint32)
var _nWTCPOptionsSetRetransmitConnectionDropTimeErr error

func tryNWTCPOptionsSetRetransmitConnectionDropTime(options NWProtocolOptions, retransmit_connection_drop_time uint32) error {
	if _nWTCPOptionsSetRetransmitConnectionDropTime == nil {
		return symbolCallError("nw_tcp_options_set_retransmit_connection_drop_time", "10.14", _nWTCPOptionsSetRetransmitConnectionDropTimeErr)
	}
	_nWTCPOptionsSetRetransmitConnectionDropTime(options, retransmit_connection_drop_time)
	return nil
}

// NWTCPOptionsSetRetransmitConnectionDropTime sets the number of seconds that TCP waits between retransmission attempts.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_connection_drop_time(_:_:)
func NWTCPOptionsSetRetransmitConnectionDropTime(options NWProtocolOptions, retransmit_connection_drop_time uint32) {
	if callErr := tryNWTCPOptionsSetRetransmitConnectionDropTime(options, retransmit_connection_drop_time); callErr != nil {
		panic(callErr)
	}
}

var _nWTCPOptionsSetRetransmitFinDrop func(options NWProtocolOptions, retransmit_fin_drop bool)
var _nWTCPOptionsSetRetransmitFinDropErr error

func tryNWTCPOptionsSetRetransmitFinDrop(options NWProtocolOptions, retransmit_fin_drop bool) error {
	if _nWTCPOptionsSetRetransmitFinDrop == nil {
		return symbolCallError("nw_tcp_options_set_retransmit_fin_drop", "10.14", _nWTCPOptionsSetRetransmitFinDropErr)
	}
	_nWTCPOptionsSetRetransmitFinDrop(options, retransmit_fin_drop)
	return nil
}

// NWTCPOptionsSetRetransmitFinDrop causes TCP to drop its connection after not receiving an ACK after a FIN.
//
// See: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_fin_drop(_:_:)
func NWTCPOptionsSetRetransmitFinDrop(options NWProtocolOptions, retransmit_fin_drop bool) {
	if callErr := tryNWTCPOptionsSetRetransmitFinDrop(options, retransmit_fin_drop); callErr != nil {
		panic(callErr)
	}
}

var _nWTLSCopySecProtocolMetadata func(metadata NWProtocolMetadata) security.Sec_protocol_metadata_t
var _nWTLSCopySecProtocolMetadataErr error

func tryNWTLSCopySecProtocolMetadata(metadata NWProtocolMetadata) (security.Sec_protocol_metadata_t, error) {
	if _nWTLSCopySecProtocolMetadata == nil {
		return *new(security.Sec_protocol_metadata_t), symbolCallError("nw_tls_copy_sec_protocol_metadata", "10.14", _nWTLSCopySecProtocolMetadataErr)
	}
	return _nWTLSCopySecProtocolMetadata(metadata), nil
}

// NWTLSCopySecProtocolMetadata accesses the result of the TLS handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_metadata(_:)
func NWTLSCopySecProtocolMetadata(metadata NWProtocolMetadata) security.Sec_protocol_metadata_t {
	result, callErr := tryNWTLSCopySecProtocolMetadata(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTLSCopySecProtocolOptions func(options NWProtocolOptions) security.Sec_protocol_options_t
var _nWTLSCopySecProtocolOptionsErr error

func tryNWTLSCopySecProtocolOptions(options NWProtocolOptions) (security.Sec_protocol_options_t, error) {
	if _nWTLSCopySecProtocolOptions == nil {
		return *new(security.Sec_protocol_options_t), symbolCallError("nw_tls_copy_sec_protocol_options", "10.14", _nWTLSCopySecProtocolOptionsErr)
	}
	return _nWTLSCopySecProtocolOptions(options), nil
}

// NWTLSCopySecProtocolOptions accesses the handshake security options TLS will use.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_options(_:)
func NWTLSCopySecProtocolOptions(options NWProtocolOptions) security.Sec_protocol_options_t {
	result, callErr := tryNWTLSCopySecProtocolOptions(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTLSCreateOptions func() NWProtocolOptions
var _nWTLSCreateOptionsErr error

func tryNWTLSCreateOptions() (NWProtocolOptions, error) {
	if _nWTLSCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_tls_create_options", "10.14", _nWTLSCreateOptionsErr)
	}
	return _nWTLSCreateOptions(), nil
}

// NWTLSCreateOptions initializes a default set of TLS connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_tls_create_options()
func NWTLSCreateOptions() NWProtocolOptions {
	result, callErr := tryNWTLSCreateOptions()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordAccessBytesFunc func(txt_record NWTXTRecord, access_bytes unsafe.Pointer) bool
var _nWTXTRecordAccessBytesFuncErr error

func tryNWTXTRecordAccessBytesFunc(txt_record NWTXTRecord, access_bytes NWTXTRecordAccessBytes) (bool, error) {
	if _nWTXTRecordAccessBytesFunc == nil {
		return false, symbolCallError("nw_txt_record_access_bytes", "10.15", _nWTXTRecordAccessBytesFuncErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *uint8, blockArg1 uint32) bool { return access_bytes(blockArg0, blockArg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWTXTRecordAccessBytesFunc(txt_record, _block0), nil
}

// NWTXTRecordAccessBytesFunc accesses the raw bytes contained within a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes(_:_:)
func NWTXTRecordAccessBytesFunc(txt_record NWTXTRecord, access_bytes NWTXTRecordAccessBytes) bool {
	result, callErr := tryNWTXTRecordAccessBytesFunc(txt_record, access_bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordAccessKeyFunc func(txt_record NWTXTRecord, key string, access_value unsafe.Pointer) bool
var _nWTXTRecordAccessKeyFuncErr error

func tryNWTXTRecordAccessKeyFunc(txt_record NWTXTRecord, key string, access_value NWTXTRecordAccessKey) (bool, error) {
	if _nWTXTRecordAccessKeyFunc == nil {
		return false, symbolCallError("nw_txt_record_access_key", "10.15", _nWTXTRecordAccessKeyFuncErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte, blockArg1 NWTXTRecordFindKey, blockArg2 *uint8, blockArg3 uint32) bool {
		return access_value(objc.GoString(blockArg0), blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWTXTRecordAccessKeyFunc(txt_record, key, _block0), nil
}

// NWTXTRecordAccessKeyFunc accesses the value for a specific key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_access_key(_:_:_:)
func NWTXTRecordAccessKeyFunc(txt_record NWTXTRecord, key string, access_value NWTXTRecordAccessKey) bool {
	result, callErr := tryNWTXTRecordAccessKeyFunc(txt_record, key, access_value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordApply func(txt_record NWTXTRecord, applier unsafe.Pointer) bool
var _nWTXTRecordApplyErr error

func tryNWTXTRecordApply(txt_record NWTXTRecord, applier NWTXTRecordApplier) (bool, error) {
	if _nWTXTRecordApply == nil {
		return false, symbolCallError("nw_txt_record_apply", "10.15", _nWTXTRecordApplyErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte, blockArg1 NWTXTRecordFindKey, blockArg2 *uint8, blockArg3 uint32) bool {
		return applier(objc.GoString(blockArg0), blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWTXTRecordApply(txt_record, _block0), nil
}

// NWTXTRecordApply iterates through all keys in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_apply(_:_:)
func NWTXTRecordApply(txt_record NWTXTRecord, applier NWTXTRecordApplier) bool {
	result, callErr := tryNWTXTRecordApply(txt_record, applier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordCopy func(txt_record NWTXTRecord) NWTXTRecord
var _nWTXTRecordCopyErr error

func tryNWTXTRecordCopy(txt_record NWTXTRecord) (NWTXTRecord, error) {
	if _nWTXTRecordCopy == nil {
		return *new(NWTXTRecord), symbolCallError("nw_txt_record_copy", "10.15", _nWTXTRecordCopyErr)
	}
	return _nWTXTRecordCopy(txt_record), nil
}

// NWTXTRecordCopy performs a deep copy of a TXT record.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_copy(_:)
func NWTXTRecordCopy(txt_record NWTXTRecord) NWTXTRecord {
	result, callErr := tryNWTXTRecordCopy(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordCreateDictionary func() NWTXTRecord
var _nWTXTRecordCreateDictionaryErr error

func tryNWTXTRecordCreateDictionary() (NWTXTRecord, error) {
	if _nWTXTRecordCreateDictionary == nil {
		return *new(NWTXTRecord), symbolCallError("nw_txt_record_create_dictionary", "10.15", _nWTXTRecordCreateDictionaryErr)
	}
	return _nWTXTRecordCreateDictionary(), nil
}

// NWTXTRecordCreateDictionary initializes a TXT record as a dictionary of strings.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_dictionary()
func NWTXTRecordCreateDictionary() NWTXTRecord {
	result, callErr := tryNWTXTRecordCreateDictionary()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordCreateWithBytes func(txt_bytes *byte, txt_len uintptr) NWTXTRecord
var _nWTXTRecordCreateWithBytesErr error

func tryNWTXTRecordCreateWithBytes(txt_bytes []byte, txt_len uintptr) (NWTXTRecord, error) {
	if _nWTXTRecordCreateWithBytes == nil {
		return *new(NWTXTRecord), symbolCallError("nw_txt_record_create_with_bytes", "10.15", _nWTXTRecordCreateWithBytesErr)
	}
	return _nWTXTRecordCreateWithBytes(unsafe.SliceData(txt_bytes), txt_len), nil
}

// NWTXTRecordCreateWithBytes initializes a TXT record with raw bytes.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_create_with_bytes(_:_:)
func NWTXTRecordCreateWithBytes(txt_bytes []byte, txt_len uintptr) NWTXTRecord {
	result, callErr := tryNWTXTRecordCreateWithBytes(txt_bytes, txt_len)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordFindKeyFunc func(txt_record NWTXTRecord, key string) NWTXTRecordFindKey
var _nWTXTRecordFindKeyFuncErr error

func tryNWTXTRecordFindKeyFunc(txt_record NWTXTRecord, key string) (NWTXTRecordFindKey, error) {
	if _nWTXTRecordFindKeyFunc == nil {
		return *new(NWTXTRecordFindKey), symbolCallError("nw_txt_record_find_key", "10.15", _nWTXTRecordFindKeyFuncErr)
	}
	return _nWTXTRecordFindKeyFunc(txt_record, key), nil
}

// NWTXTRecordFindKeyFunc checks the status of value associated with a key in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_find_key(_:_:)
func NWTXTRecordFindKeyFunc(txt_record NWTXTRecord, key string) NWTXTRecordFindKey {
	result, callErr := tryNWTXTRecordFindKeyFunc(txt_record, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordGetKeyCount func(txt_record NWTXTRecord) uintptr
var _nWTXTRecordGetKeyCountErr error

func tryNWTXTRecordGetKeyCount(txt_record NWTXTRecord) (uintptr, error) {
	if _nWTXTRecordGetKeyCount == nil {
		return 0, symbolCallError("nw_txt_record_get_key_count", "10.15", _nWTXTRecordGetKeyCountErr)
	}
	return _nWTXTRecordGetKeyCount(txt_record), nil
}

// NWTXTRecordGetKeyCount accesses the number of keys stored in the TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_get_key_count(_:)
func NWTXTRecordGetKeyCount(txt_record NWTXTRecord) uintptr {
	result, callErr := tryNWTXTRecordGetKeyCount(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordIsDictionary func(txt_record NWTXTRecord) bool
var _nWTXTRecordIsDictionaryErr error

func tryNWTXTRecordIsDictionary(txt_record NWTXTRecord) (bool, error) {
	if _nWTXTRecordIsDictionary == nil {
		return false, symbolCallError("nw_txt_record_is_dictionary", "10.15", _nWTXTRecordIsDictionaryErr)
	}
	return _nWTXTRecordIsDictionary(txt_record), nil
}

// NWTXTRecordIsDictionary checks whether a TXT record conforms to a dictionary format.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_dictionary(_:)
func NWTXTRecordIsDictionary(txt_record NWTXTRecord) bool {
	result, callErr := tryNWTXTRecordIsDictionary(txt_record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordIsEqual func(left NWTXTRecord, right NWTXTRecord) bool
var _nWTXTRecordIsEqualErr error

func tryNWTXTRecordIsEqual(left NWTXTRecord, right NWTXTRecord) (bool, error) {
	if _nWTXTRecordIsEqual == nil {
		return false, symbolCallError("nw_txt_record_is_equal", "10.15", _nWTXTRecordIsEqualErr)
	}
	return _nWTXTRecordIsEqual(left, right), nil
}

// NWTXTRecordIsEqual checks whether two TXT records are equivalent.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_is_equal(_:_:)
func NWTXTRecordIsEqual(left NWTXTRecord, right NWTXTRecord) bool {
	result, callErr := tryNWTXTRecordIsEqual(left, right)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordRemoveKey func(txt_record NWTXTRecord, key string) bool
var _nWTXTRecordRemoveKeyErr error

func tryNWTXTRecordRemoveKey(txt_record NWTXTRecord, key string) (bool, error) {
	if _nWTXTRecordRemoveKey == nil {
		return false, symbolCallError("nw_txt_record_remove_key", "10.15", _nWTXTRecordRemoveKeyErr)
	}
	return _nWTXTRecordRemoveKey(txt_record, key), nil
}

// NWTXTRecordRemoveKey removes a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_remove_key(_:_:)
func NWTXTRecordRemoveKey(txt_record NWTXTRecord, key string) bool {
	result, callErr := tryNWTXTRecordRemoveKey(txt_record, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWTXTRecordSetKey func(txt_record NWTXTRecord, key string, value *byte, value_len uintptr) bool
var _nWTXTRecordSetKeyErr error

func tryNWTXTRecordSetKey(txt_record NWTXTRecord, key string, value []byte, value_len uintptr) (bool, error) {
	if _nWTXTRecordSetKey == nil {
		return false, symbolCallError("nw_txt_record_set_key", "10.15", _nWTXTRecordSetKeyErr)
	}
	return _nWTXTRecordSetKey(txt_record, key, unsafe.SliceData(value), value_len), nil
}

// NWTXTRecordSetKey sets a data value in a TXT record dictionary.
//
// See: https://developer.apple.com/documentation/Network/nw_txt_record_set_key(_:_:_:_:)
func NWTXTRecordSetKey(txt_record NWTXTRecord, key string, value []byte, value_len uintptr) bool {
	result, callErr := tryNWTXTRecordSetKey(txt_record, key, value, value_len)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWUDPCreateMetadata func() NWProtocolMetadata
var _nWUDPCreateMetadataErr error

func tryNWUDPCreateMetadata() (NWProtocolMetadata, error) {
	if _nWUDPCreateMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_udp_create_metadata", "10.14", _nWUDPCreateMetadataErr)
	}
	return _nWUDPCreateMetadata(), nil
}

// NWUDPCreateMetadata initializes a default UDP message.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_metadata()
func NWUDPCreateMetadata() NWProtocolMetadata {
	result, callErr := tryNWUDPCreateMetadata()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWUDPCreateOptions func() NWProtocolOptions
var _nWUDPCreateOptionsErr error

func tryNWUDPCreateOptions() (NWProtocolOptions, error) {
	if _nWUDPCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_udp_create_options", "10.14", _nWUDPCreateOptionsErr)
	}
	return _nWUDPCreateOptions(), nil
}

// NWUDPCreateOptions initializes a default set of UDP connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_create_options()
func NWUDPCreateOptions() NWProtocolOptions {
	result, callErr := tryNWUDPCreateOptions()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWUDPOptionsSetPreferNoChecksum func(options NWProtocolOptions, prefer_no_checksum bool)
var _nWUDPOptionsSetPreferNoChecksumErr error

func tryNWUDPOptionsSetPreferNoChecksum(options NWProtocolOptions, prefer_no_checksum bool) error {
	if _nWUDPOptionsSetPreferNoChecksum == nil {
		return symbolCallError("nw_udp_options_set_prefer_no_checksum", "10.14", _nWUDPOptionsSetPreferNoChecksumErr)
	}
	_nWUDPOptionsSetPreferNoChecksum(options, prefer_no_checksum)
	return nil
}

// NWUDPOptionsSetPreferNoChecksum configures the connection to not send UDP checksums.
//
// See: https://developer.apple.com/documentation/Network/nw_udp_options_set_prefer_no_checksum(_:_:)
func NWUDPOptionsSetPreferNoChecksum(options NWProtocolOptions, prefer_no_checksum bool) {
	if callErr := tryNWUDPOptionsSetPreferNoChecksum(options, prefer_no_checksum); callErr != nil {
		panic(callErr)
	}
}

var _nWWsCreateMetadata func(opcode NWWsOpcode) NWProtocolMetadata
var _nWWsCreateMetadataErr error

func tryNWWsCreateMetadata(opcode NWWsOpcode) (NWProtocolMetadata, error) {
	if _nWWsCreateMetadata == nil {
		return *new(NWProtocolMetadata), symbolCallError("nw_ws_create_metadata", "10.15", _nWWsCreateMetadataErr)
	}
	return _nWWsCreateMetadata(opcode), nil
}

// NWWsCreateMetadata initializes a WebSocket message with a specific type code.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_metadata(_:)
func NWWsCreateMetadata(opcode NWWsOpcode) NWProtocolMetadata {
	result, callErr := tryNWWsCreateMetadata(opcode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsCreateOptions func(version NWWsVersion) NWProtocolOptions
var _nWWsCreateOptionsErr error

func tryNWWsCreateOptions(version NWWsVersion) (NWProtocolOptions, error) {
	if _nWWsCreateOptions == nil {
		return *new(NWProtocolOptions), symbolCallError("nw_ws_create_options", "10.15", _nWWsCreateOptionsErr)
	}
	return _nWWsCreateOptions(version), nil
}

// NWWsCreateOptions initializes a default set of WebSocket connection options.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_create_options(_:)
func NWWsCreateOptions(version NWWsVersion) NWProtocolOptions {
	result, callErr := tryNWWsCreateOptions(version)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsMetadataCopyServerResponse func(metadata NWProtocolMetadata) NWWsResponse
var _nWWsMetadataCopyServerResponseErr error

func tryNWWsMetadataCopyServerResponse(metadata NWProtocolMetadata) (NWWsResponse, error) {
	if _nWWsMetadataCopyServerResponse == nil {
		return *new(NWWsResponse), symbolCallError("nw_ws_metadata_copy_server_response", "10.15", _nWWsMetadataCopyServerResponseErr)
	}
	return _nWWsMetadataCopyServerResponse(metadata), nil
}

// NWWsMetadataCopyServerResponse accesses the WebSocket server’s response sent during the handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_copy_server_response(_:)
func NWWsMetadataCopyServerResponse(metadata NWProtocolMetadata) NWWsResponse {
	result, callErr := tryNWWsMetadataCopyServerResponse(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsMetadataGetCloseCode func(metadata NWProtocolMetadata) NWWsCloseCode
var _nWWsMetadataGetCloseCodeErr error

func tryNWWsMetadataGetCloseCode(metadata NWProtocolMetadata) (NWWsCloseCode, error) {
	if _nWWsMetadataGetCloseCode == nil {
		return *new(NWWsCloseCode), symbolCallError("nw_ws_metadata_get_close_code", "10.15", _nWWsMetadataGetCloseCodeErr)
	}
	return _nWWsMetadataGetCloseCode(metadata), nil
}

// NWWsMetadataGetCloseCode accesses the close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_close_code(_:)
func NWWsMetadataGetCloseCode(metadata NWProtocolMetadata) NWWsCloseCode {
	result, callErr := tryNWWsMetadataGetCloseCode(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsMetadataGetOpcode func(metadata NWProtocolMetadata) NWWsOpcode
var _nWWsMetadataGetOpcodeErr error

func tryNWWsMetadataGetOpcode(metadata NWProtocolMetadata) (NWWsOpcode, error) {
	if _nWWsMetadataGetOpcode == nil {
		return *new(NWWsOpcode), symbolCallError("nw_ws_metadata_get_opcode", "10.15", _nWWsMetadataGetOpcodeErr)
	}
	return _nWWsMetadataGetOpcode(metadata), nil
}

// NWWsMetadataGetOpcode checks the type code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_opcode(_:)
func NWWsMetadataGetOpcode(metadata NWProtocolMetadata) NWWsOpcode {
	result, callErr := tryNWWsMetadataGetOpcode(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsMetadataSetCloseCode func(metadata NWProtocolMetadata, close_code NWWsCloseCode)
var _nWWsMetadataSetCloseCodeErr error

func tryNWWsMetadataSetCloseCode(metadata NWProtocolMetadata, close_code NWWsCloseCode) error {
	if _nWWsMetadataSetCloseCode == nil {
		return symbolCallError("nw_ws_metadata_set_close_code", "10.15", _nWWsMetadataSetCloseCodeErr)
	}
	_nWWsMetadataSetCloseCode(metadata, close_code)
	return nil
}

// NWWsMetadataSetCloseCode sets a close code on a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_close_code(_:_:)
func NWWsMetadataSetCloseCode(metadata NWProtocolMetadata, close_code NWWsCloseCode) {
	if callErr := tryNWWsMetadataSetCloseCode(metadata, close_code); callErr != nil {
		panic(callErr)
	}
}

var _nWWsMetadataSetPongHandler func(metadata NWProtocolMetadata, client_queue uintptr, pong_handler unsafe.Pointer)
var _nWWsMetadataSetPongHandlerErr error

func tryNWWsMetadataSetPongHandler(metadata NWProtocolMetadata, client_queue dispatch.Queue, pong_handler NWWsPongHandler) error {
	if _nWWsMetadataSetPongHandler == nil {
		return symbolCallError("nw_ws_metadata_set_pong_handler", "10.15", _nWWsMetadataSetPongHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) {
		pong_handler(NWError{Object: objectivec.ObjectFromID(blockArg0)})
	})
	retainNetworkAsyncBlock(metadata.ID, "nw_ws_metadata_set_pong_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWWsMetadataSetPongHandler(metadata, uintptr(client_queue.Handle()), _block0)
	return nil
}

// NWWsMetadataSetPongHandler sets a handler on a Ping message to be invoked when the corresponding Pong message is received.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_pong_handler(_:_:_:)
func NWWsMetadataSetPongHandler(metadata NWProtocolMetadata, client_queue dispatch.Queue, pong_handler NWWsPongHandler) {
	if callErr := tryNWWsMetadataSetPongHandler(metadata, client_queue, pong_handler); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsAddAdditionalHeader func(options NWProtocolOptions, name string, value string)
var _nWWsOptionsAddAdditionalHeaderErr error

func tryNWWsOptionsAddAdditionalHeader(options NWProtocolOptions, name string, value string) error {
	if _nWWsOptionsAddAdditionalHeader == nil {
		return symbolCallError("nw_ws_options_add_additional_header", "10.15", _nWWsOptionsAddAdditionalHeaderErr)
	}
	_nWWsOptionsAddAdditionalHeader(options, name, value)
	return nil
}

// NWWsOptionsAddAdditionalHeader adds additional HTTP header fields to be sent by the client during the WebSocket handshake.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_additional_header(_:_:_:)
func NWWsOptionsAddAdditionalHeader(options NWProtocolOptions, name string, value string) {
	if callErr := tryNWWsOptionsAddAdditionalHeader(options, name, value); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsAddSubprotocol func(options NWProtocolOptions, subprotocol string)
var _nWWsOptionsAddSubprotocolErr error

func tryNWWsOptionsAddSubprotocol(options NWProtocolOptions, subprotocol string) error {
	if _nWWsOptionsAddSubprotocol == nil {
		return symbolCallError("nw_ws_options_add_subprotocol", "10.15", _nWWsOptionsAddSubprotocolErr)
	}
	_nWWsOptionsAddSubprotocol(options, subprotocol)
	return nil
}

// NWWsOptionsAddSubprotocol adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_add_subprotocol(_:_:)
func NWWsOptionsAddSubprotocol(options NWProtocolOptions, subprotocol string) {
	if callErr := tryNWWsOptionsAddSubprotocol(options, subprotocol); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsSetAutoReplyPing func(options NWProtocolOptions, auto_reply_ping bool)
var _nWWsOptionsSetAutoReplyPingErr error

func tryNWWsOptionsSetAutoReplyPing(options NWProtocolOptions, auto_reply_ping bool) error {
	if _nWWsOptionsSetAutoReplyPing == nil {
		return symbolCallError("nw_ws_options_set_auto_reply_ping", "10.15", _nWWsOptionsSetAutoReplyPingErr)
	}
	_nWWsOptionsSetAutoReplyPing(options, auto_reply_ping)
	return nil
}

// NWWsOptionsSetAutoReplyPing configures the connection to automatically reply to Ping messages instead of delivering them to you.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_auto_reply_ping(_:_:)
func NWWsOptionsSetAutoReplyPing(options NWProtocolOptions, auto_reply_ping bool) {
	if callErr := tryNWWsOptionsSetAutoReplyPing(options, auto_reply_ping); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsSetClientRequestHandler func(options NWProtocolOptions, client_queue uintptr, handler unsafe.Pointer)
var _nWWsOptionsSetClientRequestHandlerErr error

func tryNWWsOptionsSetClientRequestHandler(options NWProtocolOptions, client_queue dispatch.Queue, handler NWWsClientRequestHandler) error {
	if _nWWsOptionsSetClientRequestHandler == nil {
		return symbolCallError("nw_ws_options_set_client_request_handler", "10.15", _nWWsOptionsSetClientRequestHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objc.ID) objc.ID { return handler(objectivec.ObjectFromID(blockArg0)).ID })
	retainNetworkAsyncBlock(options.ID, "nw_ws_options_set_client_request_handler:0", _block0Value)
	_block0 := unsafe.Pointer(_block0Value)
	_nWWsOptionsSetClientRequestHandler(options, uintptr(client_queue.Handle()), _block0)
	return nil
}

// NWWsOptionsSetClientRequestHandler sets a handler to react to as a server to inbound WebSocket client handshakes.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_client_request_handler(_:_:_:)
func NWWsOptionsSetClientRequestHandler(options NWProtocolOptions, client_queue dispatch.Queue, handler NWWsClientRequestHandler) {
	if callErr := tryNWWsOptionsSetClientRequestHandler(options, client_queue, handler); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsSetMaximumMessageSize func(options NWProtocolOptions, maximum_message_size uintptr)
var _nWWsOptionsSetMaximumMessageSizeErr error

func tryNWWsOptionsSetMaximumMessageSize(options NWProtocolOptions, maximum_message_size uintptr) error {
	if _nWWsOptionsSetMaximumMessageSize == nil {
		return symbolCallError("nw_ws_options_set_maximum_message_size", "10.15", _nWWsOptionsSetMaximumMessageSizeErr)
	}
	_nWWsOptionsSetMaximumMessageSize(options, maximum_message_size)
	return nil
}

// NWWsOptionsSetMaximumMessageSize sets the maximum allowed message size, in bytes, to be received by the WebSocket connection.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_maximum_message_size(_:_:)
func NWWsOptionsSetMaximumMessageSize(options NWProtocolOptions, maximum_message_size uintptr) {
	if callErr := tryNWWsOptionsSetMaximumMessageSize(options, maximum_message_size); callErr != nil {
		panic(callErr)
	}
}

var _nWWsOptionsSetSkipHandshake func(options NWProtocolOptions, skip_handshake bool)
var _nWWsOptionsSetSkipHandshakeErr error

func tryNWWsOptionsSetSkipHandshake(options NWProtocolOptions, skip_handshake bool) error {
	if _nWWsOptionsSetSkipHandshake == nil {
		return symbolCallError("nw_ws_options_set_skip_handshake", "10.15", _nWWsOptionsSetSkipHandshakeErr)
	}
	_nWWsOptionsSetSkipHandshake(options, skip_handshake)
	return nil
}

// NWWsOptionsSetSkipHandshake specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_options_set_skip_handshake(_:_:)
func NWWsOptionsSetSkipHandshake(options NWProtocolOptions, skip_handshake bool) {
	if callErr := tryNWWsOptionsSetSkipHandshake(options, skip_handshake); callErr != nil {
		panic(callErr)
	}
}

var _nWWsRequestEnumerateAdditionalHeaders func(request NWWsRequest, enumerator unsafe.Pointer) bool
var _nWWsRequestEnumerateAdditionalHeadersErr error

func tryNWWsRequestEnumerateAdditionalHeaders(request NWWsRequest, enumerator NWWsAdditionalHeaderEnumerator) (bool, error) {
	if _nWWsRequestEnumerateAdditionalHeaders == nil {
		return false, symbolCallError("nw_ws_request_enumerate_additional_headers", "10.15", _nWWsRequestEnumerateAdditionalHeadersErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte, blockArg1 *byte) bool {
		return enumerator(objc.GoString(blockArg0), objc.GoString(blockArg1))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWWsRequestEnumerateAdditionalHeaders(request, _block0), nil
}

// NWWsRequestEnumerateAdditionalHeaders enumerates additional HTTP headers in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_additional_headers(_:_:)
func NWWsRequestEnumerateAdditionalHeaders(request NWWsRequest, enumerator NWWsAdditionalHeaderEnumerator) bool {
	result, callErr := tryNWWsRequestEnumerateAdditionalHeaders(request, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsRequestEnumerateSubprotocols func(request NWWsRequest, enumerator unsafe.Pointer) bool
var _nWWsRequestEnumerateSubprotocolsErr error

func tryNWWsRequestEnumerateSubprotocols(request NWWsRequest, enumerator NWWsSubprotocolEnumerator) (bool, error) {
	if _nWWsRequestEnumerateSubprotocols == nil {
		return false, symbolCallError("nw_ws_request_enumerate_subprotocols", "10.15", _nWWsRequestEnumerateSubprotocolsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte) bool { return enumerator(objc.GoString(blockArg0)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWWsRequestEnumerateSubprotocols(request, _block0), nil
}

// NWWsRequestEnumerateSubprotocols enumerates the supported subprotocols in a WebSocket message.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_subprotocols(_:_:)
func NWWsRequestEnumerateSubprotocols(request NWWsRequest, enumerator NWWsSubprotocolEnumerator) bool {
	result, callErr := tryNWWsRequestEnumerateSubprotocols(request, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsResponseAddAdditionalHeader func(response NWWsResponse, name string, value string)
var _nWWsResponseAddAdditionalHeaderErr error

func tryNWWsResponseAddAdditionalHeader(response NWWsResponse, name string, value string) error {
	if _nWWsResponseAddAdditionalHeader == nil {
		return symbolCallError("nw_ws_response_add_additional_header", "10.15", _nWWsResponseAddAdditionalHeaderErr)
	}
	_nWWsResponseAddAdditionalHeader(response, name, value)
	return nil
}

// NWWsResponseAddAdditionalHeader adds an additional HTTP header to a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_add_additional_header(_:_:_:)
func NWWsResponseAddAdditionalHeader(response NWWsResponse, name string, value string) {
	if callErr := tryNWWsResponseAddAdditionalHeader(response, name, value); callErr != nil {
		panic(callErr)
	}
}

var _nWWsResponseCreate func(status NWWsResponseStatus, selected_subprotocol string) NWWsResponse
var _nWWsResponseCreateErr error

func tryNWWsResponseCreate(status NWWsResponseStatus, selected_subprotocol string) (NWWsResponse, error) {
	if _nWWsResponseCreate == nil {
		return *new(NWWsResponse), symbolCallError("nw_ws_response_create", "10.15", _nWWsResponseCreateErr)
	}
	return _nWWsResponseCreate(status, selected_subprotocol), nil
}

// NWWsResponseCreate initializes a WebSocket server response with a status and selected subprotocol.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_create(_:_:)
func NWWsResponseCreate(status NWWsResponseStatus, selected_subprotocol string) NWWsResponse {
	result, callErr := tryNWWsResponseCreate(status, selected_subprotocol)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsResponseEnumerateAdditionalHeaders func(response NWWsResponse, enumerator unsafe.Pointer) bool
var _nWWsResponseEnumerateAdditionalHeadersErr error

func tryNWWsResponseEnumerateAdditionalHeaders(response NWWsResponse, enumerator NWWsAdditionalHeaderEnumerator) (bool, error) {
	if _nWWsResponseEnumerateAdditionalHeaders == nil {
		return false, symbolCallError("nw_ws_response_enumerate_additional_headers", "10.15", _nWWsResponseEnumerateAdditionalHeadersErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *byte, blockArg1 *byte) bool {
		return enumerator(objc.GoString(blockArg0), objc.GoString(blockArg1))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _nWWsResponseEnumerateAdditionalHeaders(response, _block0), nil
}

// NWWsResponseEnumerateAdditionalHeaders enumerates the additional HTTP headers in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_enumerate_additional_headers(_:_:)
func NWWsResponseEnumerateAdditionalHeaders(response NWWsResponse, enumerator NWWsAdditionalHeaderEnumerator) bool {
	result, callErr := tryNWWsResponseEnumerateAdditionalHeaders(response, enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsResponseGetSelectedSubprotocol func(response NWWsResponse) *byte
var _nWWsResponseGetSelectedSubprotocolErr error

func tryNWWsResponseGetSelectedSubprotocol(response NWWsResponse) (*byte, error) {
	if _nWWsResponseGetSelectedSubprotocol == nil {
		return nil, symbolCallError("nw_ws_response_get_selected_subprotocol", "10.15", _nWWsResponseGetSelectedSubprotocolErr)
	}
	return _nWWsResponseGetSelectedSubprotocol(response), nil
}

// NWWsResponseGetSelectedSubprotocol accesses the selected subprotocol in a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_selected_subprotocol(_:)
func NWWsResponseGetSelectedSubprotocol(response NWWsResponse) *byte {
	result, callErr := tryNWWsResponseGetSelectedSubprotocol(response)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nWWsResponseGetStatus func(response NWWsResponse) NWWsResponseStatus
var _nWWsResponseGetStatusErr error

func tryNWWsResponseGetStatus(response NWWsResponse) (NWWsResponseStatus, error) {
	if _nWWsResponseGetStatus == nil {
		return *new(NWWsResponseStatus), symbolCallError("nw_ws_response_get_status", "10.15", _nWWsResponseGetStatusErr)
	}
	return _nWWsResponseGetStatus(response), nil
}

// NWWsResponseGetStatus accesses the status of a WebSocket server response.
//
// See: https://developer.apple.com/documentation/Network/nw_ws_response_get_status(_:)
func NWWsResponseGetStatus(response NWWsResponse) NWWsResponseStatus {
	result, callErr := tryNWWsResponseGetStatus(response)
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
	registerFunc(&networkCreateSecureTCPForPlain, &networkCreateSecureTCPForPlainErr, frameworkHandle, "nw_parameters_create_secure_tcp", "10.14")
	registerFunc(&_nWAdvertiseDescriptorCopyTXTRecordObject, &_nWAdvertiseDescriptorCopyTXTRecordObjectErr, frameworkHandle, "nw_advertise_descriptor_copy_txt_record_object", "10.15")
	registerFunc(&_nWAdvertiseDescriptorCreateApplicationService, &_nWAdvertiseDescriptorCreateApplicationServiceErr, frameworkHandle, "nw_advertise_descriptor_create_application_service", "13.0")
	registerFunc(&_nWAdvertiseDescriptorCreateBonjourService, &_nWAdvertiseDescriptorCreateBonjourServiceErr, frameworkHandle, "nw_advertise_descriptor_create_bonjour_service", "10.14")
	registerFunc(&_nWAdvertiseDescriptorGetApplicationServiceName, &_nWAdvertiseDescriptorGetApplicationServiceNameErr, frameworkHandle, "nw_advertise_descriptor_get_application_service_name", "13.0")
	registerFunc(&_nWAdvertiseDescriptorGetNoAutoRename, &_nWAdvertiseDescriptorGetNoAutoRenameErr, frameworkHandle, "nw_advertise_descriptor_get_no_auto_rename", "10.14")
	registerFunc(&_nWAdvertiseDescriptorSetNoAutoRename, &_nWAdvertiseDescriptorSetNoAutoRenameErr, frameworkHandle, "nw_advertise_descriptor_set_no_auto_rename", "10.14")
	registerFunc(&_nWAdvertiseDescriptorSetTXTRecord, &_nWAdvertiseDescriptorSetTXTRecordErr, frameworkHandle, "nw_advertise_descriptor_set_txt_record", "10.14")
	registerFunc(&_nWAdvertiseDescriptorSetTXTRecordObject, &_nWAdvertiseDescriptorSetTXTRecordObjectErr, frameworkHandle, "nw_advertise_descriptor_set_txt_record_object", "10.15")
	registerFunc(&_nWBrowseDescriptorCreateApplicationService, &_nWBrowseDescriptorCreateApplicationServiceErr, frameworkHandle, "nw_browse_descriptor_create_application_service", "13.0")
	registerFunc(&_nWBrowseDescriptorCreateBonjourService, &_nWBrowseDescriptorCreateBonjourServiceErr, frameworkHandle, "nw_browse_descriptor_create_bonjour_service", "10.15")
	registerFunc(&_nWBrowseDescriptorGetApplicationServiceName, &_nWBrowseDescriptorGetApplicationServiceNameErr, frameworkHandle, "nw_browse_descriptor_get_application_service_name", "13.0")
	registerFunc(&_nWBrowseDescriptorGetBonjourServiceDomain, &_nWBrowseDescriptorGetBonjourServiceDomainErr, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_domain", "10.15")
	registerFunc(&_nWBrowseDescriptorGetBonjourServiceType, &_nWBrowseDescriptorGetBonjourServiceTypeErr, frameworkHandle, "nw_browse_descriptor_get_bonjour_service_type", "10.15")
	registerFunc(&_nWBrowseDescriptorGetIncludeTXTRecord, &_nWBrowseDescriptorGetIncludeTXTRecordErr, frameworkHandle, "nw_browse_descriptor_get_include_txt_record", "10.15")
	registerFunc(&_nWBrowseDescriptorSetIncludeTXTRecord, &_nWBrowseDescriptorSetIncludeTXTRecordErr, frameworkHandle, "nw_browse_descriptor_set_include_txt_record", "10.15")
	registerFunc(&_nWBrowseResultCopyEndpoint, &_nWBrowseResultCopyEndpointErr, frameworkHandle, "nw_browse_result_copy_endpoint", "10.15")
	registerFunc(&_nWBrowseResultCopyTXTRecordObject, &_nWBrowseResultCopyTXTRecordObjectErr, frameworkHandle, "nw_browse_result_copy_txt_record_object", "10.15")
	registerFunc(&_nWBrowseResultEnumerateInterfaces, &_nWBrowseResultEnumerateInterfacesErr, frameworkHandle, "nw_browse_result_enumerate_interfaces", "10.15")
	registerFunc(&_nWBrowseResultGetChanges, &_nWBrowseResultGetChangesErr, frameworkHandle, "nw_browse_result_get_changes", "10.15")
	registerFunc(&_nWBrowseResultGetInterfacesCount, &_nWBrowseResultGetInterfacesCountErr, frameworkHandle, "nw_browse_result_get_interfaces_count", "10.15")
	registerFunc(&_nWBrowserCancel, &_nWBrowserCancelErr, frameworkHandle, "nw_browser_cancel", "10.15")
	registerFunc(&_nWBrowserCopyBrowseDescriptor, &_nWBrowserCopyBrowseDescriptorErr, frameworkHandle, "nw_browser_copy_browse_descriptor", "10.15")
	registerFunc(&_nWBrowserCopyParameters, &_nWBrowserCopyParametersErr, frameworkHandle, "nw_browser_copy_parameters", "10.15")
	registerFunc(&_nWBrowserCreate, &_nWBrowserCreateErr, frameworkHandle, "nw_browser_create", "10.15")
	registerFunc(&_nWBrowserSetBrowseResultsChangedHandler, &_nWBrowserSetBrowseResultsChangedHandlerErr, frameworkHandle, "nw_browser_set_browse_results_changed_handler", "10.15")
	registerFunc(&_nWBrowserSetQueue, &_nWBrowserSetQueueErr, frameworkHandle, "nw_browser_set_queue", "10.15")
	registerFunc(&_nWBrowserSetStateChangedHandler, &_nWBrowserSetStateChangedHandlerErr, frameworkHandle, "nw_browser_set_state_changed_handler", "10.15")
	registerFunc(&_nWBrowserStart, &_nWBrowserStartErr, frameworkHandle, "nw_browser_start", "10.15")
	registerFunc(&_nWConnectionAccessEstablishmentReport, &_nWConnectionAccessEstablishmentReportErr, frameworkHandle, "nw_connection_access_establishment_report", "10.15")
	registerFunc(&_nWConnectionBatch, &_nWConnectionBatchErr, frameworkHandle, "nw_connection_batch", "10.14")
	registerFunc(&_nWConnectionCancel, &_nWConnectionCancelErr, frameworkHandle, "nw_connection_cancel", "10.14")
	registerFunc(&_nWConnectionCancelCurrentEndpoint, &_nWConnectionCancelCurrentEndpointErr, frameworkHandle, "nw_connection_cancel_current_endpoint", "10.14")
	registerFunc(&_nWConnectionCopyCurrentPath, &_nWConnectionCopyCurrentPathErr, frameworkHandle, "nw_connection_copy_current_path", "10.14")
	registerFunc(&_nWConnectionCopyDescription, &_nWConnectionCopyDescriptionErr, frameworkHandle, "nw_connection_copy_description", "10.14")
	registerFunc(&_nWConnectionCopyEndpoint, &_nWConnectionCopyEndpointErr, frameworkHandle, "nw_connection_copy_endpoint", "10.14")
	registerFunc(&_nWConnectionCopyParameters, &_nWConnectionCopyParametersErr, frameworkHandle, "nw_connection_copy_parameters", "10.14")
	registerFunc(&_nWConnectionCopyProtocolMetadata, &_nWConnectionCopyProtocolMetadataErr, frameworkHandle, "nw_connection_copy_protocol_metadata", "10.14")
	registerFunc(&_nWConnectionCreate, &_nWConnectionCreateErr, frameworkHandle, "nw_connection_create", "10.14")
	registerFunc(&_nWConnectionCreateNewDataTransferReport, &_nWConnectionCreateNewDataTransferReportErr, frameworkHandle, "nw_connection_create_new_data_transfer_report", "10.15")
	registerFunc(&_nWConnectionForceCancel, &_nWConnectionForceCancelErr, frameworkHandle, "nw_connection_force_cancel", "10.14")
	registerFunc(&_nWConnectionGetMaximumDatagramSize, &_nWConnectionGetMaximumDatagramSizeErr, frameworkHandle, "nw_connection_get_maximum_datagram_size", "10.14")
	registerFunc(&_nWConnectionGroupCancel, &_nWConnectionGroupCancelErr, frameworkHandle, "nw_connection_group_cancel", "11.0")
	registerFunc(&_nWConnectionGroupCopyDescriptor, &_nWConnectionGroupCopyDescriptorErr, frameworkHandle, "nw_connection_group_copy_descriptor", "11.0")
	registerFunc(&_nWConnectionGroupCopyLocalEndpointForMessage, &_nWConnectionGroupCopyLocalEndpointForMessageErr, frameworkHandle, "nw_connection_group_copy_local_endpoint_for_message", "11.0")
	registerFunc(&_nWConnectionGroupCopyParameters, &_nWConnectionGroupCopyParametersErr, frameworkHandle, "nw_connection_group_copy_parameters", "11.0")
	registerFunc(&_nWConnectionGroupCopyPathForMessage, &_nWConnectionGroupCopyPathForMessageErr, frameworkHandle, "nw_connection_group_copy_path_for_message", "11.0")
	registerFunc(&_nWConnectionGroupCopyProtocolMetadata, &_nWConnectionGroupCopyProtocolMetadataErr, frameworkHandle, "nw_connection_group_copy_protocol_metadata", "12.0")
	registerFunc(&_nWConnectionGroupCopyProtocolMetadataForMessage, &_nWConnectionGroupCopyProtocolMetadataForMessageErr, frameworkHandle, "nw_connection_group_copy_protocol_metadata_for_message", "12.0")
	registerFunc(&_nWConnectionGroupCopyRemoteEndpointForMessage, &_nWConnectionGroupCopyRemoteEndpointForMessageErr, frameworkHandle, "nw_connection_group_copy_remote_endpoint_for_message", "11.0")
	registerFunc(&_nWConnectionGroupCreate, &_nWConnectionGroupCreateErr, frameworkHandle, "nw_connection_group_create", "11.0")
	registerFunc(&_nWConnectionGroupExtractConnection, &_nWConnectionGroupExtractConnectionErr, frameworkHandle, "nw_connection_group_extract_connection", "12.0")
	registerFunc(&_nWConnectionGroupExtractConnectionForMessage, &_nWConnectionGroupExtractConnectionForMessageErr, frameworkHandle, "nw_connection_group_extract_connection_for_message", "11.0")
	registerFunc(&_nWConnectionGroupReinsertExtractedConnection, &_nWConnectionGroupReinsertExtractedConnectionErr, frameworkHandle, "nw_connection_group_reinsert_extracted_connection", "12.0")
	registerFunc(&_nWConnectionGroupReply, &_nWConnectionGroupReplyErr, frameworkHandle, "nw_connection_group_reply", "11.0")
	registerFunc(&_nWConnectionGroupSendMessage, &_nWConnectionGroupSendMessageErr, frameworkHandle, "nw_connection_group_send_message", "11.0")
	registerFunc(&_nWConnectionGroupSetNewConnectionHandler, &_nWConnectionGroupSetNewConnectionHandlerErr, frameworkHandle, "nw_connection_group_set_new_connection_handler", "12.0")
	registerFunc(&_nWConnectionGroupSetQueue, &_nWConnectionGroupSetQueueErr, frameworkHandle, "nw_connection_group_set_queue", "11.0")
	registerFunc(&_nWConnectionGroupSetReceiveHandler, &_nWConnectionGroupSetReceiveHandlerErr, frameworkHandle, "nw_connection_group_set_receive_handler", "11.0")
	registerFunc(&_nWConnectionGroupSetStateChangedHandler, &_nWConnectionGroupSetStateChangedHandlerErr, frameworkHandle, "nw_connection_group_set_state_changed_handler", "11.0")
	registerFunc(&_nWConnectionGroupStart, &_nWConnectionGroupStartErr, frameworkHandle, "nw_connection_group_start", "11.0")
	registerFunc(&_nWConnectionReceive, &_nWConnectionReceiveErr, frameworkHandle, "nw_connection_receive", "10.14")
	registerFunc(&_nWConnectionReceiveMessage, &_nWConnectionReceiveMessageErr, frameworkHandle, "nw_connection_receive_message", "10.14")
	registerFunc(&_nWConnectionRestart, &_nWConnectionRestartErr, frameworkHandle, "nw_connection_restart", "10.14")
	registerFunc(&_nWConnectionSend, &_nWConnectionSendErr, frameworkHandle, "nw_connection_send", "10.14")
	registerFunc(&_nWConnectionSetBetterPathAvailableHandler, &_nWConnectionSetBetterPathAvailableHandlerErr, frameworkHandle, "nw_connection_set_better_path_available_handler", "10.14")
	registerFunc(&_nWConnectionSetPathChangedHandler, &_nWConnectionSetPathChangedHandlerErr, frameworkHandle, "nw_connection_set_path_changed_handler", "10.14")
	registerFunc(&_nWConnectionSetQueue, &_nWConnectionSetQueueErr, frameworkHandle, "nw_connection_set_queue", "10.14")
	registerFunc(&_nWConnectionSetStateChangedHandler, &_nWConnectionSetStateChangedHandlerErr, frameworkHandle, "nw_connection_set_state_changed_handler", "10.14")
	registerFunc(&_nWConnectionSetViabilityChangedHandler, &_nWConnectionSetViabilityChangedHandlerErr, frameworkHandle, "nw_connection_set_viability_changed_handler", "10.14")
	registerFunc(&_nWConnectionStart, &_nWConnectionStartErr, frameworkHandle, "nw_connection_start", "10.14")
	registerFunc(&_nWContentContextCopyAntecedent, &_nWContentContextCopyAntecedentErr, frameworkHandle, "nw_content_context_copy_antecedent", "10.14")
	registerFunc(&_nWContentContextCopyProtocolMetadata, &_nWContentContextCopyProtocolMetadataErr, frameworkHandle, "nw_content_context_copy_protocol_metadata", "10.14")
	registerFunc(&_nWContentContextCreate, &_nWContentContextCreateErr, frameworkHandle, "nw_content_context_create", "10.14")
	registerFunc(&_nWContentContextForeachProtocolMetadata, &_nWContentContextForeachProtocolMetadataErr, frameworkHandle, "nw_content_context_foreach_protocol_metadata", "10.14")
	registerFunc(&_nWContentContextGetExpirationMilliseconds, &_nWContentContextGetExpirationMillisecondsErr, frameworkHandle, "nw_content_context_get_expiration_milliseconds", "10.14")
	registerFunc(&_nWContentContextGetIdentifier, &_nWContentContextGetIdentifierErr, frameworkHandle, "nw_content_context_get_identifier", "10.14")
	registerFunc(&_nWContentContextGetIsFinal, &_nWContentContextGetIsFinalErr, frameworkHandle, "nw_content_context_get_is_final", "10.14")
	registerFunc(&_nWContentContextGetRelativePriority, &_nWContentContextGetRelativePriorityErr, frameworkHandle, "nw_content_context_get_relative_priority", "10.14")
	registerFunc(&_nWContentContextSetAntecedent, &_nWContentContextSetAntecedentErr, frameworkHandle, "nw_content_context_set_antecedent", "10.14")
	registerFunc(&_nWContentContextSetExpirationMilliseconds, &_nWContentContextSetExpirationMillisecondsErr, frameworkHandle, "nw_content_context_set_expiration_milliseconds", "10.14")
	registerFunc(&_nWContentContextSetIsFinal, &_nWContentContextSetIsFinalErr, frameworkHandle, "nw_content_context_set_is_final", "10.14")
	registerFunc(&_nWContentContextSetMetadataForProtocol, &_nWContentContextSetMetadataForProtocolErr, frameworkHandle, "nw_content_context_set_metadata_for_protocol", "10.14")
	registerFunc(&_nWContentContextSetRelativePriority, &_nWContentContextSetRelativePriorityErr, frameworkHandle, "nw_content_context_set_relative_priority", "10.14")
	registerFunc(&_nWDataTransferReportCollect, &_nWDataTransferReportCollectErr, frameworkHandle, "nw_data_transfer_report_collect", "10.15")
	registerFunc(&_nWDataTransferReportCopyPathInterface, &_nWDataTransferReportCopyPathInterfaceErr, frameworkHandle, "nw_data_transfer_report_copy_path_interface", "10.15")
	registerFunc(&_nWDataTransferReportGetDurationMilliseconds, &_nWDataTransferReportGetDurationMillisecondsErr, frameworkHandle, "nw_data_transfer_report_get_duration_milliseconds", "10.15")
	registerFunc(&_nWDataTransferReportGetPathCount, &_nWDataTransferReportGetPathCountErr, frameworkHandle, "nw_data_transfer_report_get_path_count", "10.15")
	registerFunc(&_nWDataTransferReportGetPathRadioType, &_nWDataTransferReportGetPathRadioTypeErr, frameworkHandle, "nw_data_transfer_report_get_path_radio_type", "12.0")
	registerFunc(&_nWDataTransferReportGetReceivedApplicationByteCount, &_nWDataTransferReportGetReceivedApplicationByteCountErr, frameworkHandle, "nw_data_transfer_report_get_received_application_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetReceivedIPPacketCount, &_nWDataTransferReportGetReceivedIPPacketCountErr, frameworkHandle, "nw_data_transfer_report_get_received_ip_packet_count", "10.15")
	registerFunc(&_nWDataTransferReportGetReceivedTransportByteCount, &_nWDataTransferReportGetReceivedTransportByteCountErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetReceivedTransportDuplicateByteCount, &_nWDataTransferReportGetReceivedTransportDuplicateByteCountErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_duplicate_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetReceivedTransportOutOfOrderByteCount, &_nWDataTransferReportGetReceivedTransportOutOfOrderByteCountErr, frameworkHandle, "nw_data_transfer_report_get_received_transport_out_of_order_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetSentApplicationByteCount, &_nWDataTransferReportGetSentApplicationByteCountErr, frameworkHandle, "nw_data_transfer_report_get_sent_application_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetSentIPPacketCount, &_nWDataTransferReportGetSentIPPacketCountErr, frameworkHandle, "nw_data_transfer_report_get_sent_ip_packet_count", "10.15")
	registerFunc(&_nWDataTransferReportGetSentTransportByteCount, &_nWDataTransferReportGetSentTransportByteCountErr, frameworkHandle, "nw_data_transfer_report_get_sent_transport_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetSentTransportRetransmittedByteCount, &_nWDataTransferReportGetSentTransportRetransmittedByteCountErr, frameworkHandle, "nw_data_transfer_report_get_sent_transport_retransmitted_byte_count", "10.15")
	registerFunc(&_nWDataTransferReportGetState, &_nWDataTransferReportGetStateErr, frameworkHandle, "nw_data_transfer_report_get_state", "10.15")
	registerFunc(&_nWDataTransferReportGetTransportMinimumRttMilliseconds, &_nWDataTransferReportGetTransportMinimumRttMillisecondsErr, frameworkHandle, "nw_data_transfer_report_get_transport_minimum_rtt_milliseconds", "10.15")
	registerFunc(&_nWDataTransferReportGetTransportRttVariance, &_nWDataTransferReportGetTransportRttVarianceErr, frameworkHandle, "nw_data_transfer_report_get_transport_rtt_variance", "10.15")
	registerFunc(&_nWDataTransferReportGetTransportSmoothedRttMilliseconds, &_nWDataTransferReportGetTransportSmoothedRttMillisecondsErr, frameworkHandle, "nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds", "10.15")
	registerFunc(&_nWEndpointCopyAddressString, &_nWEndpointCopyAddressStringErr, frameworkHandle, "nw_endpoint_copy_address_string", "10.14")
	registerFunc(&_nWEndpointCopyPortString, &_nWEndpointCopyPortStringErr, frameworkHandle, "nw_endpoint_copy_port_string", "10.14")
	registerFunc(&_nWEndpointCopyTXTRecord, &_nWEndpointCopyTXTRecordErr, frameworkHandle, "nw_endpoint_copy_txt_record", "13.0")
	registerFunc(&_nWEndpointCreateAddress, &_nWEndpointCreateAddressErr, frameworkHandle, "nw_endpoint_create_address", "10.14")
	registerFunc(&_nWEndpointCreateBonjourService, &_nWEndpointCreateBonjourServiceErr, frameworkHandle, "nw_endpoint_create_bonjour_service", "10.14")
	registerFunc(&_nWEndpointCreateHost, &_nWEndpointCreateHostErr, frameworkHandle, "nw_endpoint_create_host", "10.14")
	registerFunc(&_nWEndpointCreateURL, &_nWEndpointCreateURLErr, frameworkHandle, "nw_endpoint_create_url", "10.15")
	registerFunc(&_nWEndpointGetAddress, &_nWEndpointGetAddressErr, frameworkHandle, "nw_endpoint_get_address", "10.14")
	registerFunc(&_nWEndpointGetBonjourServiceDomain, &_nWEndpointGetBonjourServiceDomainErr, frameworkHandle, "nw_endpoint_get_bonjour_service_domain", "10.14")
	registerFunc(&_nWEndpointGetBonjourServiceName, &_nWEndpointGetBonjourServiceNameErr, frameworkHandle, "nw_endpoint_get_bonjour_service_name", "10.14")
	registerFunc(&_nWEndpointGetBonjourServiceType, &_nWEndpointGetBonjourServiceTypeErr, frameworkHandle, "nw_endpoint_get_bonjour_service_type", "10.14")
	registerFunc(&_nWEndpointGetHostname, &_nWEndpointGetHostnameErr, frameworkHandle, "nw_endpoint_get_hostname", "10.14")
	registerFunc(&_nWEndpointGetPort, &_nWEndpointGetPortErr, frameworkHandle, "nw_endpoint_get_port", "10.14")
	registerFunc(&_nWEndpointGetSignature, &_nWEndpointGetSignatureErr, frameworkHandle, "nw_endpoint_get_signature", "13.0")
	registerFunc(&_nWEndpointGetType, &_nWEndpointGetTypeErr, frameworkHandle, "nw_endpoint_get_type", "10.14")
	registerFunc(&_nWEndpointGetURL, &_nWEndpointGetURLErr, frameworkHandle, "nw_endpoint_get_url", "10.15")
	registerFunc(&_nWErrorCopyCfError, &_nWErrorCopyCfErrorErr, frameworkHandle, "nw_error_copy_cf_error", "10.14")
	registerFunc(&_nWErrorGetErrorCode, &_nWErrorGetErrorCodeErr, frameworkHandle, "nw_error_get_error_code", "10.14")
	registerFunc(&_nWErrorGetErrorDomain, &_nWErrorGetErrorDomainErr, frameworkHandle, "nw_error_get_error_domain", "10.14")
	registerFunc(&_nWEstablishmentReportCopyProxyEndpoint, &_nWEstablishmentReportCopyProxyEndpointErr, frameworkHandle, "nw_establishment_report_copy_proxy_endpoint", "10.15")
	registerFunc(&_nWEstablishmentReportEnumerateProtocols, &_nWEstablishmentReportEnumerateProtocolsErr, frameworkHandle, "nw_establishment_report_enumerate_protocols", "10.15")
	registerFunc(&_nWEstablishmentReportEnumerateResolutionReports, &_nWEstablishmentReportEnumerateResolutionReportsErr, frameworkHandle, "nw_establishment_report_enumerate_resolution_reports", "11.0")
	registerFunc(&_nWEstablishmentReportEnumerateResolutions, &_nWEstablishmentReportEnumerateResolutionsErr, frameworkHandle, "nw_establishment_report_enumerate_resolutions", "10.15")
	registerFunc(&_nWEstablishmentReportGetAttemptStartedAfterMilliseconds, &_nWEstablishmentReportGetAttemptStartedAfterMillisecondsErr, frameworkHandle, "nw_establishment_report_get_attempt_started_after_milliseconds", "10.15")
	registerFunc(&_nWEstablishmentReportGetDurationMilliseconds, &_nWEstablishmentReportGetDurationMillisecondsErr, frameworkHandle, "nw_establishment_report_get_duration_milliseconds", "10.15")
	registerFunc(&_nWEstablishmentReportGetPreviousAttemptCount, &_nWEstablishmentReportGetPreviousAttemptCountErr, frameworkHandle, "nw_establishment_report_get_previous_attempt_count", "10.15")
	registerFunc(&_nWEstablishmentReportGetProxyConfigured, &_nWEstablishmentReportGetProxyConfiguredErr, frameworkHandle, "nw_establishment_report_get_proxy_configured", "10.15")
	registerFunc(&_nWEstablishmentReportGetUsedProxy, &_nWEstablishmentReportGetUsedProxyErr, frameworkHandle, "nw_establishment_report_get_used_proxy", "10.15")
	registerFunc(&_nWEthernetChannelCancel, &_nWEthernetChannelCancelErr, frameworkHandle, "nw_ethernet_channel_cancel", "10.15")
	registerFunc(&_nWEthernetChannelCreate, &_nWEthernetChannelCreateErr, frameworkHandle, "nw_ethernet_channel_create", "10.15")
	registerFunc(&_nWEthernetChannelCreateWithParameters, &_nWEthernetChannelCreateWithParametersErr, frameworkHandle, "nw_ethernet_channel_create_with_parameters", "13.0")
	registerFunc(&_nWEthernetChannelGetMaximumPayloadSize, &_nWEthernetChannelGetMaximumPayloadSizeErr, frameworkHandle, "nw_ethernet_channel_get_maximum_payload_size", "13.0")
	registerFunc(&_nWEthernetChannelSend, &_nWEthernetChannelSendErr, frameworkHandle, "nw_ethernet_channel_send", "10.15")
	registerFunc(&_nWEthernetChannelSetQueue, &_nWEthernetChannelSetQueueErr, frameworkHandle, "nw_ethernet_channel_set_queue", "10.15")
	registerFunc(&_nWEthernetChannelSetReceiveHandler, &_nWEthernetChannelSetReceiveHandlerErr, frameworkHandle, "nw_ethernet_channel_set_receive_handler", "10.15")
	registerFunc(&_nWEthernetChannelSetStateChangedHandler, &_nWEthernetChannelSetStateChangedHandlerErr, frameworkHandle, "nw_ethernet_channel_set_state_changed_handler", "10.15")
	registerFunc(&_nWEthernetChannelStart, &_nWEthernetChannelStartErr, frameworkHandle, "nw_ethernet_channel_start", "10.15")
	registerFunc(&_nWFramerAsync, &_nWFramerAsyncErr, frameworkHandle, "nw_framer_async", "10.15")
	registerFunc(&_nWFramerCopyLocalEndpoint, &_nWFramerCopyLocalEndpointErr, frameworkHandle, "nw_framer_copy_local_endpoint", "10.15")
	registerFunc(&_nWFramerCopyOptions, &_nWFramerCopyOptionsErr, frameworkHandle, "nw_framer_copy_options", "12.3")
	registerFunc(&_nWFramerCopyParameters, &_nWFramerCopyParametersErr, frameworkHandle, "nw_framer_copy_parameters", "10.15")
	registerFunc(&_nWFramerCopyRemoteEndpoint, &_nWFramerCopyRemoteEndpointErr, frameworkHandle, "nw_framer_copy_remote_endpoint", "10.15")
	registerFunc(&_nWFramerCreateDefinition, &_nWFramerCreateDefinitionErr, frameworkHandle, "nw_framer_create_definition", "10.15")
	registerFunc(&_nWFramerCreateOptions, &_nWFramerCreateOptionsErr, frameworkHandle, "nw_framer_create_options", "10.15")
	registerFunc(&_nWFramerDeliverInput, &_nWFramerDeliverInputErr, frameworkHandle, "nw_framer_deliver_input", "10.15")
	registerFunc(&_nWFramerDeliverInputNoCopy, &_nWFramerDeliverInputNoCopyErr, frameworkHandle, "nw_framer_deliver_input_no_copy", "10.15")
	registerFunc(&_nWFramerMarkFailedWithError, &_nWFramerMarkFailedWithErrorErr, frameworkHandle, "nw_framer_mark_failed_with_error", "10.15")
	registerFunc(&_nWFramerMarkReady, &_nWFramerMarkReadyErr, frameworkHandle, "nw_framer_mark_ready", "10.15")
	registerFunc(&_nWFramerMessageAccessValue, &_nWFramerMessageAccessValueErr, frameworkHandle, "nw_framer_message_access_value", "10.15")
	registerFunc(&_nWFramerMessageCopyObjectValue, &_nWFramerMessageCopyObjectValueErr, frameworkHandle, "nw_framer_message_copy_object_value", "10.15")
	registerFunc(&_nWFramerMessageCreate, &_nWFramerMessageCreateErr, frameworkHandle, "nw_framer_message_create", "10.15")
	registerFunc(&_nWFramerMessageSetObjectValue, &_nWFramerMessageSetObjectValueErr, frameworkHandle, "nw_framer_message_set_object_value", "10.15")
	registerFunc(&_nWFramerMessageSetValue, &_nWFramerMessageSetValueErr, frameworkHandle, "nw_framer_message_set_value", "10.15")
	registerFunc(&_nWFramerOptionsCopyObjectValue, &_nWFramerOptionsCopyObjectValueErr, frameworkHandle, "nw_framer_options_copy_object_value", "12.3")
	registerFunc(&_nWFramerOptionsSetObjectValue, &_nWFramerOptionsSetObjectValueErr, frameworkHandle, "nw_framer_options_set_object_value", "12.3")
	registerFunc(&_nWFramerParseInput, &_nWFramerParseInputErr, frameworkHandle, "nw_framer_parse_input", "10.15")
	registerFunc(&_nWFramerParseOutput, &_nWFramerParseOutputErr, frameworkHandle, "nw_framer_parse_output", "10.15")
	registerFunc(&_nWFramerPassThroughInput, &_nWFramerPassThroughInputErr, frameworkHandle, "nw_framer_pass_through_input", "10.15")
	registerFunc(&_nWFramerPassThroughOutput, &_nWFramerPassThroughOutputErr, frameworkHandle, "nw_framer_pass_through_output", "10.15")
	registerFunc(&_nWFramerPrependApplicationProtocol, &_nWFramerPrependApplicationProtocolErr, frameworkHandle, "nw_framer_prepend_application_protocol", "10.15")
	registerFunc(&_nWFramerProtocolCreateMessage, &_nWFramerProtocolCreateMessageErr, frameworkHandle, "nw_framer_protocol_create_message", "10.15")
	registerFunc(&_nWFramerScheduleWakeup, &_nWFramerScheduleWakeupErr, frameworkHandle, "nw_framer_schedule_wakeup", "10.15")
	registerFunc(&_nWFramerSetCleanupHandler, &_nWFramerSetCleanupHandlerErr, frameworkHandle, "nw_framer_set_cleanup_handler", "10.15")
	registerFunc(&_nWFramerSetInputHandler, &_nWFramerSetInputHandlerErr, frameworkHandle, "nw_framer_set_input_handler", "10.15")
	registerFunc(&_nWFramerSetOutputHandler, &_nWFramerSetOutputHandlerErr, frameworkHandle, "nw_framer_set_output_handler", "10.15")
	registerFunc(&_nWFramerSetStopHandler, &_nWFramerSetStopHandlerErr, frameworkHandle, "nw_framer_set_stop_handler", "10.15")
	registerFunc(&_nWFramerSetWakeupHandler, &_nWFramerSetWakeupHandlerErr, frameworkHandle, "nw_framer_set_wakeup_handler", "10.15")
	registerFunc(&_nWFramerWriteOutput, &_nWFramerWriteOutputErr, frameworkHandle, "nw_framer_write_output", "10.15")
	registerFunc(&_nWFramerWriteOutputData, &_nWFramerWriteOutputDataErr, frameworkHandle, "nw_framer_write_output_data", "10.15")
	registerFunc(&_nWFramerWriteOutputNoCopy, &_nWFramerWriteOutputNoCopyErr, frameworkHandle, "nw_framer_write_output_no_copy", "10.15")
	registerFunc(&_nWGroupDescriptorAddEndpoint, &_nWGroupDescriptorAddEndpointErr, frameworkHandle, "nw_group_descriptor_add_endpoint", "11.0")
	registerFunc(&_nWGroupDescriptorCreateMulticast, &_nWGroupDescriptorCreateMulticastErr, frameworkHandle, "nw_group_descriptor_create_multicast", "11.0")
	registerFunc(&_nWGroupDescriptorCreateMultiplex, &_nWGroupDescriptorCreateMultiplexErr, frameworkHandle, "nw_group_descriptor_create_multiplex", "12.0")
	registerFunc(&_nWGroupDescriptorEnumerateEndpoints, &_nWGroupDescriptorEnumerateEndpointsErr, frameworkHandle, "nw_group_descriptor_enumerate_endpoints", "11.0")
	registerFunc(&_nWInterfaceGetIndex, &_nWInterfaceGetIndexErr, frameworkHandle, "nw_interface_get_index", "10.14")
	registerFunc(&_nWInterfaceGetName, &_nWInterfaceGetNameErr, frameworkHandle, "nw_interface_get_name", "10.14")
	registerFunc(&_nWInterfaceGetType, &_nWInterfaceGetTypeErr, frameworkHandle, "nw_interface_get_type", "10.14")
	registerFunc(&_nWIPCreateMetadata, &_nWIPCreateMetadataErr, frameworkHandle, "nw_ip_create_metadata", "10.14")
	registerFunc(&_nWIPMetadataGetEcnFlag, &_nWIPMetadataGetEcnFlagErr, frameworkHandle, "nw_ip_metadata_get_ecn_flag", "10.14")
	registerFunc(&_nWIPMetadataGetReceiveTime, &_nWIPMetadataGetReceiveTimeErr, frameworkHandle, "nw_ip_metadata_get_receive_time", "10.14")
	registerFunc(&_nWIPMetadataGetServiceClass, &_nWIPMetadataGetServiceClassErr, frameworkHandle, "nw_ip_metadata_get_service_class", "10.14")
	registerFunc(&_nWIPMetadataSetEcnFlag, &_nWIPMetadataSetEcnFlagErr, frameworkHandle, "nw_ip_metadata_set_ecn_flag", "10.14")
	registerFunc(&_nWIPMetadataSetServiceClass, &_nWIPMetadataSetServiceClassErr, frameworkHandle, "nw_ip_metadata_set_service_class", "10.14")
	registerFunc(&_nWIPOptionsSetCalculateReceiveTime, &_nWIPOptionsSetCalculateReceiveTimeErr, frameworkHandle, "nw_ip_options_set_calculate_receive_time", "10.14")
	registerFunc(&_nWIPOptionsSetDisableFragmentation, &_nWIPOptionsSetDisableFragmentationErr, frameworkHandle, "nw_ip_options_set_disable_fragmentation", "10.14")
	registerFunc(&_nWIPOptionsSetDisableMulticastLoopback, &_nWIPOptionsSetDisableMulticastLoopbackErr, frameworkHandle, "nw_ip_options_set_disable_multicast_loopback", "11.0")
	registerFunc(&_nWIPOptionsSetHopLimit, &_nWIPOptionsSetHopLimitErr, frameworkHandle, "nw_ip_options_set_hop_limit", "10.14")
	registerFunc(&_nWIPOptionsSetLocalAddressPreference, &_nWIPOptionsSetLocalAddressPreferenceErr, frameworkHandle, "nw_ip_options_set_local_address_preference", "10.15")
	registerFunc(&_nWIPOptionsSetUseMinimumMtu, &_nWIPOptionsSetUseMinimumMtuErr, frameworkHandle, "nw_ip_options_set_use_minimum_mtu", "10.14")
	registerFunc(&_nWIPOptionsSetVersion, &_nWIPOptionsSetVersionErr, frameworkHandle, "nw_ip_options_set_version", "10.14")
	registerFunc(&_nWListenerCancel, &_nWListenerCancelErr, frameworkHandle, "nw_listener_cancel", "10.14")
	registerFunc(&_nWListenerCreate, &_nWListenerCreateErr, frameworkHandle, "nw_listener_create", "10.14")
	registerFunc(&_nWListenerCreateWithConnection, &_nWListenerCreateWithConnectionErr, frameworkHandle, "nw_listener_create_with_connection", "10.14")
	registerFunc(&_nWListenerCreateWithLaunchdKey, &_nWListenerCreateWithLaunchdKeyErr, frameworkHandle, "nw_listener_create_with_launchd_key", "10.14")
	registerFunc(&_nWListenerCreateWithPort, &_nWListenerCreateWithPortErr, frameworkHandle, "nw_listener_create_with_port", "10.14")
	registerFunc(&_nWListenerGetNewConnectionLimit, &_nWListenerGetNewConnectionLimitErr, frameworkHandle, "nw_listener_get_new_connection_limit", "10.15")
	registerFunc(&_nWListenerGetPort, &_nWListenerGetPortErr, frameworkHandle, "nw_listener_get_port", "10.14")
	registerFunc(&_nWListenerSetAdvertiseDescriptor, &_nWListenerSetAdvertiseDescriptorErr, frameworkHandle, "nw_listener_set_advertise_descriptor", "10.14")
	registerFunc(&_nWListenerSetAdvertisedEndpointChangedHandler, &_nWListenerSetAdvertisedEndpointChangedHandlerErr, frameworkHandle, "nw_listener_set_advertised_endpoint_changed_handler", "10.14")
	registerFunc(&_nWListenerSetNewConnectionGroupHandler, &_nWListenerSetNewConnectionGroupHandlerErr, frameworkHandle, "nw_listener_set_new_connection_group_handler", "12.0")
	registerFunc(&_nWListenerSetNewConnectionHandler, &_nWListenerSetNewConnectionHandlerErr, frameworkHandle, "nw_listener_set_new_connection_handler", "10.14")
	registerFunc(&_nWListenerSetNewConnectionLimit, &_nWListenerSetNewConnectionLimitErr, frameworkHandle, "nw_listener_set_new_connection_limit", "10.15")
	registerFunc(&_nWListenerSetQueue, &_nWListenerSetQueueErr, frameworkHandle, "nw_listener_set_queue", "10.14")
	registerFunc(&_nWListenerSetStateChangedHandler, &_nWListenerSetStateChangedHandlerErr, frameworkHandle, "nw_listener_set_state_changed_handler", "10.14")
	registerFunc(&_nWListenerStart, &_nWListenerStartErr, frameworkHandle, "nw_listener_start", "10.14")
	registerFunc(&_nWMulticastGroupDescriptorGetDisableUnicastTraffic, &_nWMulticastGroupDescriptorGetDisableUnicastTrafficErr, frameworkHandle, "nw_multicast_group_descriptor_get_disable_unicast_traffic", "11.0")
	registerFunc(&_nWMulticastGroupDescriptorSetDisableUnicastTraffic, &_nWMulticastGroupDescriptorSetDisableUnicastTrafficErr, frameworkHandle, "nw_multicast_group_descriptor_set_disable_unicast_traffic", "11.0")
	registerFunc(&_nWMulticastGroupDescriptorSetSpecificSource, &_nWMulticastGroupDescriptorSetSpecificSourceErr, frameworkHandle, "nw_multicast_group_descriptor_set_specific_source", "11.0")
	registerFunc(&_nWParametersClearProhibitedInterfaceTypes, &_nWParametersClearProhibitedInterfaceTypesErr, frameworkHandle, "nw_parameters_clear_prohibited_interface_types", "10.14")
	registerFunc(&_nWParametersClearProhibitedInterfaces, &_nWParametersClearProhibitedInterfacesErr, frameworkHandle, "nw_parameters_clear_prohibited_interfaces", "10.14")
	registerFunc(&_nWParametersCopy, &_nWParametersCopyErr, frameworkHandle, "nw_parameters_copy", "10.14")
	registerFunc(&_nWParametersCopyDefaultProtocolStack, &_nWParametersCopyDefaultProtocolStackErr, frameworkHandle, "nw_parameters_copy_default_protocol_stack", "10.14")
	registerFunc(&_nWParametersCopyLocalEndpoint, &_nWParametersCopyLocalEndpointErr, frameworkHandle, "nw_parameters_copy_local_endpoint", "10.14")
	registerFunc(&_nWParametersCopyRequiredInterface, &_nWParametersCopyRequiredInterfaceErr, frameworkHandle, "nw_parameters_copy_required_interface", "10.14")
	registerFunc(&_nWParametersCreate, &_nWParametersCreateErr, frameworkHandle, "nw_parameters_create", "10.14")
	registerFunc(&_nWParametersCreateApplicationService, &_nWParametersCreateApplicationServiceErr, frameworkHandle, "nw_parameters_create_application_service", "13.0")
	registerFunc(&_nWParametersCreateCustomIP, &_nWParametersCreateCustomIPErr, frameworkHandle, "nw_parameters_create_custom_ip", "10.15")
	registerFunc(&_nWParametersCreateQuic, &_nWParametersCreateQuicErr, frameworkHandle, "nw_parameters_create_quic", "12.0")
	registerFunc(&_nWParametersCreateSecureTCP, &_nWParametersCreateSecureTCPErr, frameworkHandle, "nw_parameters_create_secure_tcp", "10.14")
	registerFunc(&_nWParametersCreateSecureUDP, &_nWParametersCreateSecureUDPErr, frameworkHandle, "nw_parameters_create_secure_udp", "10.14")
	registerFunc(&_nWParametersGetAllowUltraConstrained, &_nWParametersGetAllowUltraConstrainedErr, frameworkHandle, "nw_parameters_get_allow_ultra_constrained", "26.0")
	registerFunc(&_nWParametersGetAttribution, &_nWParametersGetAttributionErr, frameworkHandle, "nw_parameters_get_attribution", "12.0")
	registerFunc(&_nWParametersGetExpiredDnsBehavior, &_nWParametersGetExpiredDnsBehaviorErr, frameworkHandle, "nw_parameters_get_expired_dns_behavior", "10.14")
	registerFunc(&_nWParametersGetFastOpenEnabled, &_nWParametersGetFastOpenEnabledErr, frameworkHandle, "nw_parameters_get_fast_open_enabled", "10.14")
	registerFunc(&_nWParametersGetIncludePeerToPeer, &_nWParametersGetIncludePeerToPeerErr, frameworkHandle, "nw_parameters_get_include_peer_to_peer", "10.14")
	registerFunc(&_nWParametersGetLocalOnly, &_nWParametersGetLocalOnlyErr, frameworkHandle, "nw_parameters_get_local_only", "10.14")
	registerFunc(&_nWParametersGetMultipathService, &_nWParametersGetMultipathServiceErr, frameworkHandle, "nw_parameters_get_multipath_service", "10.14")
	registerFunc(&_nWParametersGetPreferNoProxy, &_nWParametersGetPreferNoProxyErr, frameworkHandle, "nw_parameters_get_prefer_no_proxy", "10.14")
	registerFunc(&_nWParametersGetProhibitConstrained, &_nWParametersGetProhibitConstrainedErr, frameworkHandle, "nw_parameters_get_prohibit_constrained", "10.15")
	registerFunc(&_nWParametersGetProhibitExpensive, &_nWParametersGetProhibitExpensiveErr, frameworkHandle, "nw_parameters_get_prohibit_expensive", "10.14")
	registerFunc(&_nWParametersGetRequiredInterfaceType, &_nWParametersGetRequiredInterfaceTypeErr, frameworkHandle, "nw_parameters_get_required_interface_type", "10.14")
	registerFunc(&_nWParametersGetReuseLocalAddress, &_nWParametersGetReuseLocalAddressErr, frameworkHandle, "nw_parameters_get_reuse_local_address", "10.14")
	registerFunc(&_nWParametersGetServiceClass, &_nWParametersGetServiceClassErr, frameworkHandle, "nw_parameters_get_service_class", "10.14")
	registerFunc(&_nWParametersIterateProhibitedInterfaceTypes, &_nWParametersIterateProhibitedInterfaceTypesErr, frameworkHandle, "nw_parameters_iterate_prohibited_interface_types", "10.14")
	registerFunc(&_nWParametersIterateProhibitedInterfaces, &_nWParametersIterateProhibitedInterfacesErr, frameworkHandle, "nw_parameters_iterate_prohibited_interfaces", "10.14")
	registerFunc(&_nWParametersProhibitInterface, &_nWParametersProhibitInterfaceErr, frameworkHandle, "nw_parameters_prohibit_interface", "10.14")
	registerFunc(&_nWParametersProhibitInterfaceType, &_nWParametersProhibitInterfaceTypeErr, frameworkHandle, "nw_parameters_prohibit_interface_type", "10.14")
	registerFunc(&_nWParametersRequireInterface, &_nWParametersRequireInterfaceErr, frameworkHandle, "nw_parameters_require_interface", "10.14")
	registerFunc(&_nWParametersRequiresDnssecValidation, &_nWParametersRequiresDnssecValidationErr, frameworkHandle, "nw_parameters_requires_dnssec_validation", "13.0")
	registerFunc(&_nWParametersSetAllowUltraConstrained, &_nWParametersSetAllowUltraConstrainedErr, frameworkHandle, "nw_parameters_set_allow_ultra_constrained", "26.0")
	registerFunc(&_nWParametersSetAttribution, &_nWParametersSetAttributionErr, frameworkHandle, "nw_parameters_set_attribution", "12.0")
	registerFunc(&_nWParametersSetExpiredDnsBehavior, &_nWParametersSetExpiredDnsBehaviorErr, frameworkHandle, "nw_parameters_set_expired_dns_behavior", "10.14")
	registerFunc(&_nWParametersSetFastOpenEnabled, &_nWParametersSetFastOpenEnabledErr, frameworkHandle, "nw_parameters_set_fast_open_enabled", "10.14")
	registerFunc(&_nWParametersSetIncludePeerToPeer, &_nWParametersSetIncludePeerToPeerErr, frameworkHandle, "nw_parameters_set_include_peer_to_peer", "10.14")
	registerFunc(&_nWParametersSetLocalEndpoint, &_nWParametersSetLocalEndpointErr, frameworkHandle, "nw_parameters_set_local_endpoint", "10.14")
	registerFunc(&_nWParametersSetLocalOnly, &_nWParametersSetLocalOnlyErr, frameworkHandle, "nw_parameters_set_local_only", "10.14")
	registerFunc(&_nWParametersSetMultipathService, &_nWParametersSetMultipathServiceErr, frameworkHandle, "nw_parameters_set_multipath_service", "10.14")
	registerFunc(&_nWParametersSetPreferNoProxy, &_nWParametersSetPreferNoProxyErr, frameworkHandle, "nw_parameters_set_prefer_no_proxy", "10.14")
	registerFunc(&_nWParametersSetPrivacyContext, &_nWParametersSetPrivacyContextErr, frameworkHandle, "nw_parameters_set_privacy_context", "11.0")
	registerFunc(&_nWParametersSetProhibitConstrained, &_nWParametersSetProhibitConstrainedErr, frameworkHandle, "nw_parameters_set_prohibit_constrained", "10.15")
	registerFunc(&_nWParametersSetProhibitExpensive, &_nWParametersSetProhibitExpensiveErr, frameworkHandle, "nw_parameters_set_prohibit_expensive", "10.14")
	registerFunc(&_nWParametersSetRequiredInterfaceType, &_nWParametersSetRequiredInterfaceTypeErr, frameworkHandle, "nw_parameters_set_required_interface_type", "10.14")
	registerFunc(&_nWParametersSetRequiresDnssecValidation, &_nWParametersSetRequiresDnssecValidationErr, frameworkHandle, "nw_parameters_set_requires_dnssec_validation", "13.0")
	registerFunc(&_nWParametersSetReuseLocalAddress, &_nWParametersSetReuseLocalAddressErr, frameworkHandle, "nw_parameters_set_reuse_local_address", "10.14")
	registerFunc(&_nWParametersSetServiceClass, &_nWParametersSetServiceClassErr, frameworkHandle, "nw_parameters_set_service_class", "10.14")
	registerFunc(&_nWPathCopyEffectiveLocalEndpoint, &_nWPathCopyEffectiveLocalEndpointErr, frameworkHandle, "nw_path_copy_effective_local_endpoint", "10.14")
	registerFunc(&_nWPathCopyEffectiveRemoteEndpoint, &_nWPathCopyEffectiveRemoteEndpointErr, frameworkHandle, "nw_path_copy_effective_remote_endpoint", "10.14")
	registerFunc(&_nWPathEnumerateGateways, &_nWPathEnumerateGatewaysErr, frameworkHandle, "nw_path_enumerate_gateways", "10.15")
	registerFunc(&_nWPathEnumerateInterfaces, &_nWPathEnumerateInterfacesErr, frameworkHandle, "nw_path_enumerate_interfaces", "10.14")
	registerFunc(&_nWPathGetLinkQuality, &_nWPathGetLinkQualityErr, frameworkHandle, "nw_path_get_link_quality", "26.0")
	registerFunc(&_nWPathGetStatus, &_nWPathGetStatusErr, frameworkHandle, "nw_path_get_status", "10.14")
	registerFunc(&_nWPathGetUnsatisfiedReason, &_nWPathGetUnsatisfiedReasonErr, frameworkHandle, "nw_path_get_unsatisfied_reason", "11.0")
	registerFunc(&_nWPathHasDns, &_nWPathHasDnsErr, frameworkHandle, "nw_path_has_dns", "10.14")
	registerFunc(&_nWPathHasIpv4, &_nWPathHasIpv4Err, frameworkHandle, "nw_path_has_ipv4", "10.14")
	registerFunc(&_nWPathHasIpv6, &_nWPathHasIpv6Err, frameworkHandle, "nw_path_has_ipv6", "10.14")
	registerFunc(&_nWPathIsConstrained, &_nWPathIsConstrainedErr, frameworkHandle, "nw_path_is_constrained", "10.15")
	registerFunc(&_nWPathIsEqual, &_nWPathIsEqualErr, frameworkHandle, "nw_path_is_equal", "10.14")
	registerFunc(&_nWPathIsExpensive, &_nWPathIsExpensiveErr, frameworkHandle, "nw_path_is_expensive", "10.14")
	registerFunc(&_nWPathIsUltraConstrained, &_nWPathIsUltraConstrainedErr, frameworkHandle, "nw_path_is_ultra_constrained", "26.0")
	registerFunc(&_nWPathMonitorCancel, &_nWPathMonitorCancelErr, frameworkHandle, "nw_path_monitor_cancel", "10.14")
	registerFunc(&_nWPathMonitorCreate, &_nWPathMonitorCreateErr, frameworkHandle, "nw_path_monitor_create", "10.14")
	registerFunc(&_nWPathMonitorCreateForEthernetChannel, &_nWPathMonitorCreateForEthernetChannelErr, frameworkHandle, "nw_path_monitor_create_for_ethernet_channel", "13.0")
	registerFunc(&_nWPathMonitorCreateWithType, &_nWPathMonitorCreateWithTypeErr, frameworkHandle, "nw_path_monitor_create_with_type", "10.14")
	registerFunc(&_nWPathMonitorProhibitInterfaceType, &_nWPathMonitorProhibitInterfaceTypeErr, frameworkHandle, "nw_path_monitor_prohibit_interface_type", "11.0")
	registerFunc(&_nWPathMonitorSetCancelHandler, &_nWPathMonitorSetCancelHandlerErr, frameworkHandle, "nw_path_monitor_set_cancel_handler", "10.14")
	registerFunc(&_nWPathMonitorSetQueue, &_nWPathMonitorSetQueueErr, frameworkHandle, "nw_path_monitor_set_queue", "10.14")
	registerFunc(&_nWPathMonitorSetUpdateHandler, &_nWPathMonitorSetUpdateHandlerErr, frameworkHandle, "nw_path_monitor_set_update_handler", "10.14")
	registerFunc(&_nWPathMonitorStart, &_nWPathMonitorStartErr, frameworkHandle, "nw_path_monitor_start", "10.14")
	registerFunc(&_nWPathUsesInterfaceType, &_nWPathUsesInterfaceTypeErr, frameworkHandle, "nw_path_uses_interface_type", "10.14")
	registerFunc(&_nWPrivacyContextAddProxy, &_nWPrivacyContextAddProxyErr, frameworkHandle, "nw_privacy_context_add_proxy", "14.0")
	registerFunc(&_nWPrivacyContextClearProxies, &_nWPrivacyContextClearProxiesErr, frameworkHandle, "nw_privacy_context_clear_proxies", "14.0")
	registerFunc(&_nWPrivacyContextCreate, &_nWPrivacyContextCreateErr, frameworkHandle, "nw_privacy_context_create", "11.0")
	registerFunc(&_nWPrivacyContextDisableLogging, &_nWPrivacyContextDisableLoggingErr, frameworkHandle, "nw_privacy_context_disable_logging", "11.0")
	registerFunc(&_nWPrivacyContextFlushCache, &_nWPrivacyContextFlushCacheErr, frameworkHandle, "nw_privacy_context_flush_cache", "11.0")
	registerFunc(&_nWPrivacyContextRequireEncryptedNameResolution, &_nWPrivacyContextRequireEncryptedNameResolutionErr, frameworkHandle, "nw_privacy_context_require_encrypted_name_resolution", "11.0")
	registerFunc(&_nWProtocolCopyIPDefinition, &_nWProtocolCopyIPDefinitionErr, frameworkHandle, "nw_protocol_copy_ip_definition", "10.14")
	registerFunc(&_nWProtocolCopyQuicDefinition, &_nWProtocolCopyQuicDefinitionErr, frameworkHandle, "nw_protocol_copy_quic_definition", "12.0")
	registerFunc(&_nWProtocolCopyTCPDefinition, &_nWProtocolCopyTCPDefinitionErr, frameworkHandle, "nw_protocol_copy_tcp_definition", "10.14")
	registerFunc(&_nWProtocolCopyTLSDefinition, &_nWProtocolCopyTLSDefinitionErr, frameworkHandle, "nw_protocol_copy_tls_definition", "10.14")
	registerFunc(&_nWProtocolCopyUDPDefinition, &_nWProtocolCopyUDPDefinitionErr, frameworkHandle, "nw_protocol_copy_udp_definition", "10.14")
	registerFunc(&_nWProtocolCopyWsDefinition, &_nWProtocolCopyWsDefinitionErr, frameworkHandle, "nw_protocol_copy_ws_definition", "10.15")
	registerFunc(&_nWProtocolDefinitionIsEqual, &_nWProtocolDefinitionIsEqualErr, frameworkHandle, "nw_protocol_definition_is_equal", "10.14")
	registerFunc(&_nWProtocolMetadataCopyDefinition, &_nWProtocolMetadataCopyDefinitionErr, frameworkHandle, "nw_protocol_metadata_copy_definition", "10.14")
	registerFunc(&_nWProtocolMetadataIsFramerMessage, &_nWProtocolMetadataIsFramerMessageErr, frameworkHandle, "nw_protocol_metadata_is_framer_message", "10.15")
	registerFunc(&_nWProtocolMetadataIsIP, &_nWProtocolMetadataIsIPErr, frameworkHandle, "nw_protocol_metadata_is_ip", "10.14")
	registerFunc(&_nWProtocolMetadataIsQuic, &_nWProtocolMetadataIsQuicErr, frameworkHandle, "nw_protocol_metadata_is_quic", "12.0")
	registerFunc(&_nWProtocolMetadataIsTCP, &_nWProtocolMetadataIsTCPErr, frameworkHandle, "nw_protocol_metadata_is_tcp", "10.14")
	registerFunc(&_nWProtocolMetadataIsTLS, &_nWProtocolMetadataIsTLSErr, frameworkHandle, "nw_protocol_metadata_is_tls", "10.14")
	registerFunc(&_nWProtocolMetadataIsUDP, &_nWProtocolMetadataIsUDPErr, frameworkHandle, "nw_protocol_metadata_is_udp", "10.14")
	registerFunc(&_nWProtocolMetadataIsWs, &_nWProtocolMetadataIsWsErr, frameworkHandle, "nw_protocol_metadata_is_ws", "10.15")
	registerFunc(&_nWProtocolOptionsCopyDefinition, &_nWProtocolOptionsCopyDefinitionErr, frameworkHandle, "nw_protocol_options_copy_definition", "10.14")
	registerFunc(&_nWProtocolOptionsIsQuic, &_nWProtocolOptionsIsQuicErr, frameworkHandle, "nw_protocol_options_is_quic", "12.0")
	registerFunc(&_nWProtocolStackClearApplicationProtocols, &_nWProtocolStackClearApplicationProtocolsErr, frameworkHandle, "nw_protocol_stack_clear_application_protocols", "10.14")
	registerFunc(&_nWProtocolStackCopyInternetProtocol, &_nWProtocolStackCopyInternetProtocolErr, frameworkHandle, "nw_protocol_stack_copy_internet_protocol", "10.14")
	registerFunc(&_nWProtocolStackCopyTransportProtocol, &_nWProtocolStackCopyTransportProtocolErr, frameworkHandle, "nw_protocol_stack_copy_transport_protocol", "10.14")
	registerFunc(&_nWProtocolStackIterateApplicationProtocols, &_nWProtocolStackIterateApplicationProtocolsErr, frameworkHandle, "nw_protocol_stack_iterate_application_protocols", "10.14")
	registerFunc(&_nWProtocolStackPrependApplicationProtocol, &_nWProtocolStackPrependApplicationProtocolErr, frameworkHandle, "nw_protocol_stack_prepend_application_protocol", "10.14")
	registerFunc(&_nWProtocolStackSetTransportProtocol, &_nWProtocolStackSetTransportProtocolErr, frameworkHandle, "nw_protocol_stack_set_transport_protocol", "10.14")
	registerFunc(&_nWProxyConfigAddExcludedDomain, &_nWProxyConfigAddExcludedDomainErr, frameworkHandle, "nw_proxy_config_add_excluded_domain", "14.0")
	registerFunc(&_nWProxyConfigAddMatchDomain, &_nWProxyConfigAddMatchDomainErr, frameworkHandle, "nw_proxy_config_add_match_domain", "14.0")
	registerFunc(&_nWProxyConfigClearExcludedDomains, &_nWProxyConfigClearExcludedDomainsErr, frameworkHandle, "nw_proxy_config_clear_excluded_domains", "14.0")
	registerFunc(&_nWProxyConfigClearMatchDomains, &_nWProxyConfigClearMatchDomainsErr, frameworkHandle, "nw_proxy_config_clear_match_domains", "14.0")
	registerFunc(&_nWProxyConfigCreateHttpConnect, &_nWProxyConfigCreateHttpConnectErr, frameworkHandle, "nw_proxy_config_create_http_connect", "14.0")
	registerFunc(&_nWProxyConfigCreateObliviousHttp, &_nWProxyConfigCreateObliviousHttpErr, frameworkHandle, "nw_proxy_config_create_oblivious_http", "14.0")
	registerFunc(&_nWProxyConfigCreateRelay, &_nWProxyConfigCreateRelayErr, frameworkHandle, "nw_proxy_config_create_relay", "14.0")
	registerFunc(&_nWProxyConfigCreateSocksv5, &_nWProxyConfigCreateSocksv5Err, frameworkHandle, "nw_proxy_config_create_socksv5", "14.0")
	registerFunc(&_nWProxyConfigEnumerateExcludedDomains, &_nWProxyConfigEnumerateExcludedDomainsErr, frameworkHandle, "nw_proxy_config_enumerate_excluded_domains", "14.0")
	registerFunc(&_nWProxyConfigEnumerateMatchDomains, &_nWProxyConfigEnumerateMatchDomainsErr, frameworkHandle, "nw_proxy_config_enumerate_match_domains", "14.0")
	registerFunc(&_nWProxyConfigGetFailoverAllowed, &_nWProxyConfigGetFailoverAllowedErr, frameworkHandle, "nw_proxy_config_get_failover_allowed", "14.0")
	registerFunc(&_nWProxyConfigSetFailoverAllowed, &_nWProxyConfigSetFailoverAllowedErr, frameworkHandle, "nw_proxy_config_set_failover_allowed", "14.0")
	registerFunc(&_nWProxyConfigSetUsernameAndPassword, &_nWProxyConfigSetUsernameAndPasswordErr, frameworkHandle, "nw_proxy_config_set_username_and_password", "14.0")
	registerFunc(&_nWQuicAddTLSApplicationProtocol, &_nWQuicAddTLSApplicationProtocolErr, frameworkHandle, "nw_quic_add_tls_application_protocol", "12.0")
	registerFunc(&_nWQuicCopySecProtocolMetadata, &_nWQuicCopySecProtocolMetadataErr, frameworkHandle, "nw_quic_copy_sec_protocol_metadata", "12.0")
	registerFunc(&_nWQuicCopySecProtocolOptions, &_nWQuicCopySecProtocolOptionsErr, frameworkHandle, "nw_quic_copy_sec_protocol_options", "12.0")
	registerFunc(&_nWQuicCreateOptions, &_nWQuicCreateOptionsErr, frameworkHandle, "nw_quic_create_options", "12.0")
	registerFunc(&_nWQuicGetApplicationError, &_nWQuicGetApplicationErrorErr, frameworkHandle, "nw_quic_get_application_error", "12.0")
	registerFunc(&_nWQuicGetApplicationErrorReason, &_nWQuicGetApplicationErrorReasonErr, frameworkHandle, "nw_quic_get_application_error_reason", "12.0")
	registerFunc(&_nWQuicGetIdleTimeout, &_nWQuicGetIdleTimeoutErr, frameworkHandle, "nw_quic_get_idle_timeout", "12.0")
	registerFunc(&_nWQuicGetInitialMaxData, &_nWQuicGetInitialMaxDataErr, frameworkHandle, "nw_quic_get_initial_max_data", "12.0")
	registerFunc(&_nWQuicGetInitialMaxStreamDataBidirectionalLocal, &_nWQuicGetInitialMaxStreamDataBidirectionalLocalErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_local", "12.0")
	registerFunc(&_nWQuicGetInitialMaxStreamDataBidirectionalRemote, &_nWQuicGetInitialMaxStreamDataBidirectionalRemoteErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_bidirectional_remote", "12.0")
	registerFunc(&_nWQuicGetInitialMaxStreamDataUnidirectional, &_nWQuicGetInitialMaxStreamDataUnidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_stream_data_unidirectional", "12.0")
	registerFunc(&_nWQuicGetInitialMaxStreamsBidirectional, &_nWQuicGetInitialMaxStreamsBidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_streams_bidirectional", "12.0")
	registerFunc(&_nWQuicGetInitialMaxStreamsUnidirectional, &_nWQuicGetInitialMaxStreamsUnidirectionalErr, frameworkHandle, "nw_quic_get_initial_max_streams_unidirectional", "12.0")
	registerFunc(&_nWQuicGetKeepaliveInterval, &_nWQuicGetKeepaliveIntervalErr, frameworkHandle, "nw_quic_get_keepalive_interval", "12.0")
	registerFunc(&_nWQuicGetLocalMaxStreamsBidirectional, &_nWQuicGetLocalMaxStreamsBidirectionalErr, frameworkHandle, "nw_quic_get_local_max_streams_bidirectional", "12.0")
	registerFunc(&_nWQuicGetLocalMaxStreamsUnidirectional, &_nWQuicGetLocalMaxStreamsUnidirectionalErr, frameworkHandle, "nw_quic_get_local_max_streams_unidirectional", "12.0")
	registerFunc(&_nWQuicGetMaxDatagramFrameSize, &_nWQuicGetMaxDatagramFrameSizeErr, frameworkHandle, "nw_quic_get_max_datagram_frame_size", "13.0")
	registerFunc(&_nWQuicGetMaxUDPPayloadSize, &_nWQuicGetMaxUDPPayloadSizeErr, frameworkHandle, "nw_quic_get_max_udp_payload_size", "12.0")
	registerFunc(&_nWQuicGetRemoteIdleTimeout, &_nWQuicGetRemoteIdleTimeoutErr, frameworkHandle, "nw_quic_get_remote_idle_timeout", "12.0")
	registerFunc(&_nWQuicGetRemoteMaxStreamsBidirectional, &_nWQuicGetRemoteMaxStreamsBidirectionalErr, frameworkHandle, "nw_quic_get_remote_max_streams_bidirectional", "12.0")
	registerFunc(&_nWQuicGetRemoteMaxStreamsUnidirectional, &_nWQuicGetRemoteMaxStreamsUnidirectionalErr, frameworkHandle, "nw_quic_get_remote_max_streams_unidirectional", "12.0")
	registerFunc(&_nWQuicGetStreamApplicationError, &_nWQuicGetStreamApplicationErrorErr, frameworkHandle, "nw_quic_get_stream_application_error", "12.0")
	registerFunc(&_nWQuicGetStreamID, &_nWQuicGetStreamIDErr, frameworkHandle, "nw_quic_get_stream_id", "12.0")
	registerFunc(&_nWQuicGetStreamIsDatagram, &_nWQuicGetStreamIsDatagramErr, frameworkHandle, "nw_quic_get_stream_is_datagram", "13.0")
	registerFunc(&_nWQuicGetStreamIsUnidirectional, &_nWQuicGetStreamIsUnidirectionalErr, frameworkHandle, "nw_quic_get_stream_is_unidirectional", "12.0")
	registerFunc(&_nWQuicGetStreamType, &_nWQuicGetStreamTypeErr, frameworkHandle, "nw_quic_get_stream_type", "12.0")
	registerFunc(&_nWQuicGetStreamUsableDatagramFrameSize, &_nWQuicGetStreamUsableDatagramFrameSizeErr, frameworkHandle, "nw_quic_get_stream_usable_datagram_frame_size", "13.0")
	registerFunc(&_nWQuicSetApplicationError, &_nWQuicSetApplicationErrorErr, frameworkHandle, "nw_quic_set_application_error", "12.0")
	registerFunc(&_nWQuicSetIdleTimeout, &_nWQuicSetIdleTimeoutErr, frameworkHandle, "nw_quic_set_idle_timeout", "12.0")
	registerFunc(&_nWQuicSetInitialMaxData, &_nWQuicSetInitialMaxDataErr, frameworkHandle, "nw_quic_set_initial_max_data", "12.0")
	registerFunc(&_nWQuicSetInitialMaxStreamDataBidirectionalLocal, &_nWQuicSetInitialMaxStreamDataBidirectionalLocalErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_local", "12.0")
	registerFunc(&_nWQuicSetInitialMaxStreamDataBidirectionalRemote, &_nWQuicSetInitialMaxStreamDataBidirectionalRemoteErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_bidirectional_remote", "12.0")
	registerFunc(&_nWQuicSetInitialMaxStreamDataUnidirectional, &_nWQuicSetInitialMaxStreamDataUnidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_stream_data_unidirectional", "12.0")
	registerFunc(&_nWQuicSetInitialMaxStreamsBidirectional, &_nWQuicSetInitialMaxStreamsBidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_streams_bidirectional", "12.0")
	registerFunc(&_nWQuicSetInitialMaxStreamsUnidirectional, &_nWQuicSetInitialMaxStreamsUnidirectionalErr, frameworkHandle, "nw_quic_set_initial_max_streams_unidirectional", "12.0")
	registerFunc(&_nWQuicSetKeepaliveInterval, &_nWQuicSetKeepaliveIntervalErr, frameworkHandle, "nw_quic_set_keepalive_interval", "12.0")
	registerFunc(&_nWQuicSetLocalMaxStreamsBidirectional, &_nWQuicSetLocalMaxStreamsBidirectionalErr, frameworkHandle, "nw_quic_set_local_max_streams_bidirectional", "12.0")
	registerFunc(&_nWQuicSetLocalMaxStreamsUnidirectional, &_nWQuicSetLocalMaxStreamsUnidirectionalErr, frameworkHandle, "nw_quic_set_local_max_streams_unidirectional", "12.0")
	registerFunc(&_nWQuicSetMaxDatagramFrameSize, &_nWQuicSetMaxDatagramFrameSizeErr, frameworkHandle, "nw_quic_set_max_datagram_frame_size", "13.0")
	registerFunc(&_nWQuicSetMaxUDPPayloadSize, &_nWQuicSetMaxUDPPayloadSizeErr, frameworkHandle, "nw_quic_set_max_udp_payload_size", "12.0")
	registerFunc(&_nWQuicSetStreamApplicationError, &_nWQuicSetStreamApplicationErrorErr, frameworkHandle, "nw_quic_set_stream_application_error", "12.0")
	registerFunc(&_nWQuicSetStreamIsDatagram, &_nWQuicSetStreamIsDatagramErr, frameworkHandle, "nw_quic_set_stream_is_datagram", "13.0")
	registerFunc(&_nWQuicSetStreamIsUnidirectional, &_nWQuicSetStreamIsUnidirectionalErr, frameworkHandle, "nw_quic_set_stream_is_unidirectional", "12.0")
	registerFunc(&_nWRelayHopAddAdditionalHttpHeaderField, &_nWRelayHopAddAdditionalHttpHeaderFieldErr, frameworkHandle, "nw_relay_hop_add_additional_http_header_field", "14.0")
	registerFunc(&_nWRelayHopCreate, &_nWRelayHopCreateErr, frameworkHandle, "nw_relay_hop_create", "14.0")
	registerFunc(&_nWRelease, &_nWReleaseErr, frameworkHandle, "nw_release", "10.14")
	registerFunc(&_nWResolutionReportCopyPreferredEndpoint, &_nWResolutionReportCopyPreferredEndpointErr, frameworkHandle, "nw_resolution_report_copy_preferred_endpoint", "11.0")
	registerFunc(&_nWResolutionReportCopySuccessfulEndpoint, &_nWResolutionReportCopySuccessfulEndpointErr, frameworkHandle, "nw_resolution_report_copy_successful_endpoint", "11.0")
	registerFunc(&_nWResolutionReportGetEndpointCount, &_nWResolutionReportGetEndpointCountErr, frameworkHandle, "nw_resolution_report_get_endpoint_count", "11.0")
	registerFunc(&_nWResolutionReportGetMilliseconds, &_nWResolutionReportGetMillisecondsErr, frameworkHandle, "nw_resolution_report_get_milliseconds", "11.0")
	registerFunc(&_nWResolutionReportGetProtocol, &_nWResolutionReportGetProtocolErr, frameworkHandle, "nw_resolution_report_get_protocol", "11.0")
	registerFunc(&_nWResolutionReportGetSource, &_nWResolutionReportGetSourceErr, frameworkHandle, "nw_resolution_report_get_source", "11.0")
	registerFunc(&_nWResolverConfigAddServerAddress, &_nWResolverConfigAddServerAddressErr, frameworkHandle, "nw_resolver_config_add_server_address", "11.0")
	registerFunc(&_nWResolverConfigCreateHttps, &_nWResolverConfigCreateHttpsErr, frameworkHandle, "nw_resolver_config_create_https", "11.0")
	registerFunc(&_nWResolverConfigCreateTLS, &_nWResolverConfigCreateTLSErr, frameworkHandle, "nw_resolver_config_create_tls", "11.0")
	registerFunc(&_nWRetain, &_nWRetainErr, frameworkHandle, "nw_retain", "10.14")
	registerFunc(&_nWTCPCreateOptions, &_nWTCPCreateOptionsErr, frameworkHandle, "nw_tcp_create_options", "10.14")
	registerFunc(&_nWTCPGetAvailableReceiveBuffer, &_nWTCPGetAvailableReceiveBufferErr, frameworkHandle, "nw_tcp_get_available_receive_buffer", "10.14")
	registerFunc(&_nWTCPGetAvailableSendBuffer, &_nWTCPGetAvailableSendBufferErr, frameworkHandle, "nw_tcp_get_available_send_buffer", "10.14")
	registerFunc(&_nWTCPOptionsSetConnectionTimeout, &_nWTCPOptionsSetConnectionTimeoutErr, frameworkHandle, "nw_tcp_options_set_connection_timeout", "10.14")
	registerFunc(&_nWTCPOptionsSetDisableAckStretching, &_nWTCPOptionsSetDisableAckStretchingErr, frameworkHandle, "nw_tcp_options_set_disable_ack_stretching", "10.14")
	registerFunc(&_nWTCPOptionsSetDisableEcn, &_nWTCPOptionsSetDisableEcnErr, frameworkHandle, "nw_tcp_options_set_disable_ecn", "10.14")
	registerFunc(&_nWTCPOptionsSetEnableFastOpen, &_nWTCPOptionsSetEnableFastOpenErr, frameworkHandle, "nw_tcp_options_set_enable_fast_open", "10.14")
	registerFunc(&_nWTCPOptionsSetEnableKeepalive, &_nWTCPOptionsSetEnableKeepaliveErr, frameworkHandle, "nw_tcp_options_set_enable_keepalive", "10.14")
	registerFunc(&_nWTCPOptionsSetKeepaliveCount, &_nWTCPOptionsSetKeepaliveCountErr, frameworkHandle, "nw_tcp_options_set_keepalive_count", "10.14")
	registerFunc(&_nWTCPOptionsSetKeepaliveIdleTime, &_nWTCPOptionsSetKeepaliveIdleTimeErr, frameworkHandle, "nw_tcp_options_set_keepalive_idle_time", "10.14")
	registerFunc(&_nWTCPOptionsSetKeepaliveInterval, &_nWTCPOptionsSetKeepaliveIntervalErr, frameworkHandle, "nw_tcp_options_set_keepalive_interval", "10.14")
	registerFunc(&_nWTCPOptionsSetMaximumSegmentSize, &_nWTCPOptionsSetMaximumSegmentSizeErr, frameworkHandle, "nw_tcp_options_set_maximum_segment_size", "10.14")
	registerFunc(&_nWTCPOptionsSetMultipathForceVersion, &_nWTCPOptionsSetMultipathForceVersionErr, frameworkHandle, "nw_tcp_options_set_multipath_force_version", "12.0")
	registerFunc(&_nWTCPOptionsSetNoDelay, &_nWTCPOptionsSetNoDelayErr, frameworkHandle, "nw_tcp_options_set_no_delay", "10.14")
	registerFunc(&_nWTCPOptionsSetNoOptions, &_nWTCPOptionsSetNoOptionsErr, frameworkHandle, "nw_tcp_options_set_no_options", "10.14")
	registerFunc(&_nWTCPOptionsSetNoPush, &_nWTCPOptionsSetNoPushErr, frameworkHandle, "nw_tcp_options_set_no_push", "10.14")
	registerFunc(&_nWTCPOptionsSetPersistTimeout, &_nWTCPOptionsSetPersistTimeoutErr, frameworkHandle, "nw_tcp_options_set_persist_timeout", "10.14")
	registerFunc(&_nWTCPOptionsSetRetransmitConnectionDropTime, &_nWTCPOptionsSetRetransmitConnectionDropTimeErr, frameworkHandle, "nw_tcp_options_set_retransmit_connection_drop_time", "10.14")
	registerFunc(&_nWTCPOptionsSetRetransmitFinDrop, &_nWTCPOptionsSetRetransmitFinDropErr, frameworkHandle, "nw_tcp_options_set_retransmit_fin_drop", "10.14")
	registerFunc(&_nWTLSCopySecProtocolMetadata, &_nWTLSCopySecProtocolMetadataErr, frameworkHandle, "nw_tls_copy_sec_protocol_metadata", "10.14")
	registerFunc(&_nWTLSCopySecProtocolOptions, &_nWTLSCopySecProtocolOptionsErr, frameworkHandle, "nw_tls_copy_sec_protocol_options", "10.14")
	registerFunc(&_nWTLSCreateOptions, &_nWTLSCreateOptionsErr, frameworkHandle, "nw_tls_create_options", "10.14")
	registerFunc(&_nWTXTRecordAccessBytesFunc, &_nWTXTRecordAccessBytesFuncErr, frameworkHandle, "nw_txt_record_access_bytes", "10.15")
	registerFunc(&_nWTXTRecordAccessKeyFunc, &_nWTXTRecordAccessKeyFuncErr, frameworkHandle, "nw_txt_record_access_key", "10.15")
	registerFunc(&_nWTXTRecordApply, &_nWTXTRecordApplyErr, frameworkHandle, "nw_txt_record_apply", "10.15")
	registerFunc(&_nWTXTRecordCopy, &_nWTXTRecordCopyErr, frameworkHandle, "nw_txt_record_copy", "10.15")
	registerFunc(&_nWTXTRecordCreateDictionary, &_nWTXTRecordCreateDictionaryErr, frameworkHandle, "nw_txt_record_create_dictionary", "10.15")
	registerFunc(&_nWTXTRecordCreateWithBytes, &_nWTXTRecordCreateWithBytesErr, frameworkHandle, "nw_txt_record_create_with_bytes", "10.15")
	registerFunc(&_nWTXTRecordFindKeyFunc, &_nWTXTRecordFindKeyFuncErr, frameworkHandle, "nw_txt_record_find_key", "10.15")
	registerFunc(&_nWTXTRecordGetKeyCount, &_nWTXTRecordGetKeyCountErr, frameworkHandle, "nw_txt_record_get_key_count", "10.15")
	registerFunc(&_nWTXTRecordIsDictionary, &_nWTXTRecordIsDictionaryErr, frameworkHandle, "nw_txt_record_is_dictionary", "10.15")
	registerFunc(&_nWTXTRecordIsEqual, &_nWTXTRecordIsEqualErr, frameworkHandle, "nw_txt_record_is_equal", "10.15")
	registerFunc(&_nWTXTRecordRemoveKey, &_nWTXTRecordRemoveKeyErr, frameworkHandle, "nw_txt_record_remove_key", "10.15")
	registerFunc(&_nWTXTRecordSetKey, &_nWTXTRecordSetKeyErr, frameworkHandle, "nw_txt_record_set_key", "10.15")
	registerFunc(&_nWUDPCreateMetadata, &_nWUDPCreateMetadataErr, frameworkHandle, "nw_udp_create_metadata", "10.14")
	registerFunc(&_nWUDPCreateOptions, &_nWUDPCreateOptionsErr, frameworkHandle, "nw_udp_create_options", "10.14")
	registerFunc(&_nWUDPOptionsSetPreferNoChecksum, &_nWUDPOptionsSetPreferNoChecksumErr, frameworkHandle, "nw_udp_options_set_prefer_no_checksum", "10.14")
	registerFunc(&_nWWsCreateMetadata, &_nWWsCreateMetadataErr, frameworkHandle, "nw_ws_create_metadata", "10.15")
	registerFunc(&_nWWsCreateOptions, &_nWWsCreateOptionsErr, frameworkHandle, "nw_ws_create_options", "10.15")
	registerFunc(&_nWWsMetadataCopyServerResponse, &_nWWsMetadataCopyServerResponseErr, frameworkHandle, "nw_ws_metadata_copy_server_response", "10.15")
	registerFunc(&_nWWsMetadataGetCloseCode, &_nWWsMetadataGetCloseCodeErr, frameworkHandle, "nw_ws_metadata_get_close_code", "10.15")
	registerFunc(&_nWWsMetadataGetOpcode, &_nWWsMetadataGetOpcodeErr, frameworkHandle, "nw_ws_metadata_get_opcode", "10.15")
	registerFunc(&_nWWsMetadataSetCloseCode, &_nWWsMetadataSetCloseCodeErr, frameworkHandle, "nw_ws_metadata_set_close_code", "10.15")
	registerFunc(&_nWWsMetadataSetPongHandler, &_nWWsMetadataSetPongHandlerErr, frameworkHandle, "nw_ws_metadata_set_pong_handler", "10.15")
	registerFunc(&_nWWsOptionsAddAdditionalHeader, &_nWWsOptionsAddAdditionalHeaderErr, frameworkHandle, "nw_ws_options_add_additional_header", "10.15")
	registerFunc(&_nWWsOptionsAddSubprotocol, &_nWWsOptionsAddSubprotocolErr, frameworkHandle, "nw_ws_options_add_subprotocol", "10.15")
	registerFunc(&_nWWsOptionsSetAutoReplyPing, &_nWWsOptionsSetAutoReplyPingErr, frameworkHandle, "nw_ws_options_set_auto_reply_ping", "10.15")
	registerFunc(&_nWWsOptionsSetClientRequestHandler, &_nWWsOptionsSetClientRequestHandlerErr, frameworkHandle, "nw_ws_options_set_client_request_handler", "10.15")
	registerFunc(&_nWWsOptionsSetMaximumMessageSize, &_nWWsOptionsSetMaximumMessageSizeErr, frameworkHandle, "nw_ws_options_set_maximum_message_size", "10.15")
	registerFunc(&_nWWsOptionsSetSkipHandshake, &_nWWsOptionsSetSkipHandshakeErr, frameworkHandle, "nw_ws_options_set_skip_handshake", "10.15")
	registerFunc(&_nWWsRequestEnumerateAdditionalHeaders, &_nWWsRequestEnumerateAdditionalHeadersErr, frameworkHandle, "nw_ws_request_enumerate_additional_headers", "10.15")
	registerFunc(&_nWWsRequestEnumerateSubprotocols, &_nWWsRequestEnumerateSubprotocolsErr, frameworkHandle, "nw_ws_request_enumerate_subprotocols", "10.15")
	registerFunc(&_nWWsResponseAddAdditionalHeader, &_nWWsResponseAddAdditionalHeaderErr, frameworkHandle, "nw_ws_response_add_additional_header", "10.15")
	registerFunc(&_nWWsResponseCreate, &_nWWsResponseCreateErr, frameworkHandle, "nw_ws_response_create", "10.15")
	registerFunc(&_nWWsResponseEnumerateAdditionalHeaders, &_nWWsResponseEnumerateAdditionalHeadersErr, frameworkHandle, "nw_ws_response_enumerate_additional_headers", "10.15")
	registerFunc(&_nWWsResponseGetSelectedSubprotocol, &_nWWsResponseGetSelectedSubprotocolErr, frameworkHandle, "nw_ws_response_get_selected_subprotocol", "10.15")
	registerFunc(&_nWWsResponseGetStatus, &_nWWsResponseGetStatusErr, frameworkHandle, "nw_ws_response_get_status", "10.15")
}
