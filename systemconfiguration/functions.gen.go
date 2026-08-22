// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
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
		return fmt.Sprintf("SystemConfiguration: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("SystemConfiguration: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("SystemConfiguration: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("SystemConfiguration: register symbol %s: %v", name, r)
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

var _cNCopySupportedInterfaces func() corefoundation.CFArrayRef
var _cNCopySupportedInterfacesErr error

func tryCNCopySupportedInterfaces() (corefoundation.CFArrayRef, error) {
	if _cNCopySupportedInterfaces == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CNCopySupportedInterfaces", "10.8", _cNCopySupportedInterfacesErr)
	}
	return _cNCopySupportedInterfaces(), nil
}

// CNCopySupportedInterfaces returns the names of all network interfaces Captive Network Support is monitoring.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/CNCopySupportedInterfaces
func CNCopySupportedInterfaces() corefoundation.CFArrayRef {
	result, callErr := tryCNCopySupportedInterfaces()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cNMarkPortalOffline func(interfaceName corefoundation.CFStringRef) bool
var _cNMarkPortalOfflineErr error

func tryCNMarkPortalOffline(interfaceName corefoundation.CFStringRef) (bool, error) {
	if _cNMarkPortalOffline == nil {
		return false, symbolCallError("CNMarkPortalOffline", "10.8", _cNMarkPortalOfflineErr)
	}
	return _cNMarkPortalOffline(interfaceName), nil
}

// CNMarkPortalOffline informs Captive Network Support that the device is not authenticated on a captive network.
//
// Deprecated: Deprecated.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOffline
func CNMarkPortalOffline(interfaceName corefoundation.CFStringRef) bool {
	result, callErr := tryCNMarkPortalOffline(interfaceName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cNMarkPortalOnline func(interfaceName corefoundation.CFStringRef) bool
var _cNMarkPortalOnlineErr error

func tryCNMarkPortalOnline(interfaceName corefoundation.CFStringRef) (bool, error) {
	if _cNMarkPortalOnline == nil {
		return false, symbolCallError("CNMarkPortalOnline", "10.8", _cNMarkPortalOnlineErr)
	}
	return _cNMarkPortalOnline(interfaceName), nil
}

// CNMarkPortalOnline informs Captive Network Support that the application has successfully authenticated the device to a captive network.
//
// Deprecated: Deprecated.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOnline
func CNMarkPortalOnline(interfaceName corefoundation.CFStringRef) bool {
	result, callErr := tryCNMarkPortalOnline(interfaceName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cNSetSupportedSSIDs func(ssidArray corefoundation.CFArrayRef) bool
var _cNSetSupportedSSIDsErr error

func tryCNSetSupportedSSIDs(ssidArray corefoundation.CFArrayRef) (bool, error) {
	if _cNSetSupportedSSIDs == nil {
		return false, symbolCallError("CNSetSupportedSSIDs", "10.8", _cNSetSupportedSSIDsErr)
	}
	return _cNSetSupportedSSIDs(ssidArray), nil
}

// CNSetSupportedSSIDs specifies an updated list of captive network SSIDs that the application performs authentication on.
//
// Deprecated: Deprecated.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/CNSetSupportedSSIDs
func CNSetSupportedSSIDs(ssidArray corefoundation.CFArrayRef) bool {
	result, callErr := tryCNSetSupportedSSIDs(ssidArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dHCPClientPreferencesCopyApplicationOptions func(applicationID corefoundation.CFStringRef, count *int) *uint8
var _dHCPClientPreferencesCopyApplicationOptionsErr error

func tryDHCPClientPreferencesCopyApplicationOptions(applicationID corefoundation.CFStringRef, count *int) (*uint8, error) {
	if _dHCPClientPreferencesCopyApplicationOptions == nil {
		return nil, symbolCallError("DHCPClientPreferencesCopyApplicationOptions", "10.1", _dHCPClientPreferencesCopyApplicationOptionsErr)
	}
	return _dHCPClientPreferencesCopyApplicationOptions(applicationID, count), nil
}

// DHCPClientPreferencesCopyApplicationOptions returns the list of options for the specified application ID.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesCopyApplicationOptions
func DHCPClientPreferencesCopyApplicationOptions(applicationID corefoundation.CFStringRef, count *int) *uint8 {
	result, callErr := tryDHCPClientPreferencesCopyApplicationOptions(applicationID, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dHCPClientPreferencesSetApplicationOptions func(applicationID corefoundation.CFStringRef, options *byte, count int) bool
var _dHCPClientPreferencesSetApplicationOptionsErr error

func tryDHCPClientPreferencesSetApplicationOptions(applicationID corefoundation.CFStringRef, options []byte, count int) (bool, error) {
	if _dHCPClientPreferencesSetApplicationOptions == nil {
		return false, symbolCallError("DHCPClientPreferencesSetApplicationOptions", "10.1", _dHCPClientPreferencesSetApplicationOptionsErr)
	}
	return _dHCPClientPreferencesSetApplicationOptions(applicationID, unsafe.SliceData(options), count), nil
}

// DHCPClientPreferencesSetApplicationOptions updates the DHCP client preferences to include the specified list of options for the specified application ID.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesSetApplicationOptions
func DHCPClientPreferencesSetApplicationOptions(applicationID corefoundation.CFStringRef, options []byte, count int) bool {
	result, callErr := tryDHCPClientPreferencesSetApplicationOptions(applicationID, options, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dHCPInfoGetLeaseExpirationTime func(info corefoundation.CFDictionaryRef) corefoundation.CFDateRef
var _dHCPInfoGetLeaseExpirationTimeErr error

func tryDHCPInfoGetLeaseExpirationTime(info corefoundation.CFDictionaryRef) (corefoundation.CFDateRef, error) {
	if _dHCPInfoGetLeaseExpirationTime == nil {
		return *new(corefoundation.CFDateRef), symbolCallError("DHCPInfoGetLeaseExpirationTime", "10.8", _dHCPInfoGetLeaseExpirationTimeErr)
	}
	return _dHCPInfoGetLeaseExpirationTime(info), nil
}

// DHCPInfoGetLeaseExpirationTime returns the lease expiration time data.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseExpirationTime
func DHCPInfoGetLeaseExpirationTime(info corefoundation.CFDictionaryRef) corefoundation.CFDateRef {
	result, callErr := tryDHCPInfoGetLeaseExpirationTime(info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dHCPInfoGetLeaseStartTime func(info corefoundation.CFDictionaryRef) corefoundation.CFDateRef
var _dHCPInfoGetLeaseStartTimeErr error

func tryDHCPInfoGetLeaseStartTime(info corefoundation.CFDictionaryRef) (corefoundation.CFDateRef, error) {
	if _dHCPInfoGetLeaseStartTime == nil {
		return *new(corefoundation.CFDateRef), symbolCallError("DHCPInfoGetLeaseStartTime", "10.1", _dHCPInfoGetLeaseStartTimeErr)
	}
	return _dHCPInfoGetLeaseStartTime(info), nil
}

// DHCPInfoGetLeaseStartTime returns the lease start time data.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseStartTime
func DHCPInfoGetLeaseStartTime(info corefoundation.CFDictionaryRef) corefoundation.CFDateRef {
	result, callErr := tryDHCPInfoGetLeaseStartTime(info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dHCPInfoGetOptionData func(info corefoundation.CFDictionaryRef, code uint8) corefoundation.CFDataRef
var _dHCPInfoGetOptionDataErr error

func tryDHCPInfoGetOptionData(info corefoundation.CFDictionaryRef, code uint8) (corefoundation.CFDataRef, error) {
	if _dHCPInfoGetOptionData == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("DHCPInfoGetOptionData", "10.1", _dHCPInfoGetOptionDataErr)
	}
	return _dHCPInfoGetOptionData(info, code), nil
}

// DHCPInfoGetOptionData returns DHCP option data, if present.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetOptionData
func DHCPInfoGetOptionData(info corefoundation.CFDictionaryRef, code uint8) corefoundation.CFDataRef {
	result, callErr := tryDHCPInfoGetOptionData(info, code)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceCopyAll func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCBondInterfaceCopyAllErr error

func trySCBondInterfaceCopyAll(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCBondInterfaceCopyAll == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCBondInterfaceCopyAll", "10.5", _sCBondInterfaceCopyAllErr)
	}
	return _sCBondInterfaceCopyAll(prefs), nil
}

// SCBondInterfaceCopyAll returns all Ethernet bond interfaces on the system.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAll(_:)
func SCBondInterfaceCopyAll(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCBondInterfaceCopyAll(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceCopyAvailableMemberInterfaces func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCBondInterfaceCopyAvailableMemberInterfacesErr error

func trySCBondInterfaceCopyAvailableMemberInterfaces(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCBondInterfaceCopyAvailableMemberInterfaces == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCBondInterfaceCopyAvailableMemberInterfaces", "10.5", _sCBondInterfaceCopyAvailableMemberInterfacesErr)
	}
	return _sCBondInterfaceCopyAvailableMemberInterfaces(prefs), nil
}

// SCBondInterfaceCopyAvailableMemberInterfaces returns all network capable devices on the system that can be added to an Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAvailableMemberInterfaces(_:)
func SCBondInterfaceCopyAvailableMemberInterfaces(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCBondInterfaceCopyAvailableMemberInterfaces(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceCopyStatus func(bond SCBondInterfaceRef) SCBondStatusRef
var _sCBondInterfaceCopyStatusErr error

func trySCBondInterfaceCopyStatus(bond SCBondInterfaceRef) (SCBondStatusRef, error) {
	if _sCBondInterfaceCopyStatus == nil {
		return *new(SCBondStatusRef), symbolCallError("SCBondInterfaceCopyStatus", "10.5", _sCBondInterfaceCopyStatusErr)
	}
	return _sCBondInterfaceCopyStatus(bond), nil
}

// SCBondInterfaceCopyStatus returns the status of the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyStatus(_:)
func SCBondInterfaceCopyStatus(bond SCBondInterfaceRef) SCBondStatusRef {
	result, callErr := trySCBondInterfaceCopyStatus(bond)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceCreate func(prefs SCPreferencesRef) SCBondInterfaceRef
var _sCBondInterfaceCreateErr error

func trySCBondInterfaceCreate(prefs SCPreferencesRef) (SCBondInterfaceRef, error) {
	if _sCBondInterfaceCreate == nil {
		return *new(SCBondInterfaceRef), symbolCallError("SCBondInterfaceCreate", "10.5", _sCBondInterfaceCreateErr)
	}
	return _sCBondInterfaceCreate(prefs), nil
}

// SCBondInterfaceCreate creates a new Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCreate(_:)
func SCBondInterfaceCreate(prefs SCPreferencesRef) SCBondInterfaceRef {
	result, callErr := trySCBondInterfaceCreate(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceGetMemberInterfaces func(bond SCBondInterfaceRef) corefoundation.CFArrayRef
var _sCBondInterfaceGetMemberInterfacesErr error

func trySCBondInterfaceGetMemberInterfaces(bond SCBondInterfaceRef) (corefoundation.CFArrayRef, error) {
	if _sCBondInterfaceGetMemberInterfaces == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCBondInterfaceGetMemberInterfaces", "10.5", _sCBondInterfaceGetMemberInterfacesErr)
	}
	return _sCBondInterfaceGetMemberInterfaces(bond), nil
}

// SCBondInterfaceGetMemberInterfaces returns the member interfaces for the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetMemberInterfaces(_:)
func SCBondInterfaceGetMemberInterfaces(bond SCBondInterfaceRef) corefoundation.CFArrayRef {
	result, callErr := trySCBondInterfaceGetMemberInterfaces(bond)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceGetOptions func(bond SCBondInterfaceRef) corefoundation.CFDictionaryRef
var _sCBondInterfaceGetOptionsErr error

func trySCBondInterfaceGetOptions(bond SCBondInterfaceRef) (corefoundation.CFDictionaryRef, error) {
	if _sCBondInterfaceGetOptions == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCBondInterfaceGetOptions", "10.5", _sCBondInterfaceGetOptionsErr)
	}
	return _sCBondInterfaceGetOptions(bond), nil
}

// SCBondInterfaceGetOptions returns the configuration settings associated with the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetOptions(_:)
func SCBondInterfaceGetOptions(bond SCBondInterfaceRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCBondInterfaceGetOptions(bond)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceRemove func(bond SCBondInterfaceRef) bool
var _sCBondInterfaceRemoveErr error

func trySCBondInterfaceRemove(bond SCBondInterfaceRef) (bool, error) {
	if _sCBondInterfaceRemove == nil {
		return false, symbolCallError("SCBondInterfaceRemove", "10.5", _sCBondInterfaceRemoveErr)
	}
	return _sCBondInterfaceRemove(bond), nil
}

// SCBondInterfaceRemove removes the Ethernet bond interface from the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceRemove(_:)
func SCBondInterfaceRemove(bond SCBondInterfaceRef) bool {
	result, callErr := trySCBondInterfaceRemove(bond)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceSetLocalizedDisplayName func(bond SCBondInterfaceRef, newName corefoundation.CFStringRef) bool
var _sCBondInterfaceSetLocalizedDisplayNameErr error

func trySCBondInterfaceSetLocalizedDisplayName(bond SCBondInterfaceRef, newName corefoundation.CFStringRef) (bool, error) {
	if _sCBondInterfaceSetLocalizedDisplayName == nil {
		return false, symbolCallError("SCBondInterfaceSetLocalizedDisplayName", "10.5", _sCBondInterfaceSetLocalizedDisplayNameErr)
	}
	return _sCBondInterfaceSetLocalizedDisplayName(bond, newName), nil
}

// SCBondInterfaceSetLocalizedDisplayName sets the localized display name for the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetLocalizedDisplayName(_:_:)
func SCBondInterfaceSetLocalizedDisplayName(bond SCBondInterfaceRef, newName corefoundation.CFStringRef) bool {
	result, callErr := trySCBondInterfaceSetLocalizedDisplayName(bond, newName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceSetMemberInterfaces func(bond SCBondInterfaceRef, members corefoundation.CFArrayRef) bool
var _sCBondInterfaceSetMemberInterfacesErr error

func trySCBondInterfaceSetMemberInterfaces(bond SCBondInterfaceRef, members corefoundation.CFArrayRef) (bool, error) {
	if _sCBondInterfaceSetMemberInterfaces == nil {
		return false, symbolCallError("SCBondInterfaceSetMemberInterfaces", "10.5", _sCBondInterfaceSetMemberInterfacesErr)
	}
	return _sCBondInterfaceSetMemberInterfaces(bond, members), nil
}

// SCBondInterfaceSetMemberInterfaces sets the member interfaces for the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetMemberInterfaces(_:_:)
func SCBondInterfaceSetMemberInterfaces(bond SCBondInterfaceRef, members corefoundation.CFArrayRef) bool {
	result, callErr := trySCBondInterfaceSetMemberInterfaces(bond, members)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondInterfaceSetOptions func(bond SCBondInterfaceRef, newOptions corefoundation.CFDictionaryRef) bool
var _sCBondInterfaceSetOptionsErr error

func trySCBondInterfaceSetOptions(bond SCBondInterfaceRef, newOptions corefoundation.CFDictionaryRef) (bool, error) {
	if _sCBondInterfaceSetOptions == nil {
		return false, symbolCallError("SCBondInterfaceSetOptions", "10.5", _sCBondInterfaceSetOptionsErr)
	}
	return _sCBondInterfaceSetOptions(bond, newOptions), nil
}

// SCBondInterfaceSetOptions sets the configuration settings for the specified Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetOptions(_:_:)
func SCBondInterfaceSetOptions(bond SCBondInterfaceRef, newOptions corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCBondInterfaceSetOptions(bond, newOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondStatusGetInterfaceStatus func(bondStatus SCBondStatusRef, interface_ SCNetworkInterfaceRef) corefoundation.CFDictionaryRef
var _sCBondStatusGetInterfaceStatusErr error

func trySCBondStatusGetInterfaceStatus(bondStatus SCBondStatusRef, interface_ SCNetworkInterfaceRef) (corefoundation.CFDictionaryRef, error) {
	if _sCBondStatusGetInterfaceStatus == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCBondStatusGetInterfaceStatus", "10.5", _sCBondStatusGetInterfaceStatusErr)
	}
	return _sCBondStatusGetInterfaceStatus(bondStatus, interface_), nil
}

// SCBondStatusGetInterfaceStatus returns the status of the specified member interface of an Ethernet bond or the status of the bond as a whole.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetInterfaceStatus(_:_:)
func SCBondStatusGetInterfaceStatus(bondStatus SCBondStatusRef, interface_ SCNetworkInterfaceRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCBondStatusGetInterfaceStatus(bondStatus, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondStatusGetMemberInterfaces func(bondStatus SCBondStatusRef) corefoundation.CFArrayRef
var _sCBondStatusGetMemberInterfacesErr error

func trySCBondStatusGetMemberInterfaces(bondStatus SCBondStatusRef) (corefoundation.CFArrayRef, error) {
	if _sCBondStatusGetMemberInterfaces == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCBondStatusGetMemberInterfaces", "10.5", _sCBondStatusGetMemberInterfacesErr)
	}
	return _sCBondStatusGetMemberInterfaces(bondStatus), nil
}

// SCBondStatusGetMemberInterfaces returns the member interfaces that are represented with the Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetMemberInterfaces(_:)
func SCBondStatusGetMemberInterfaces(bondStatus SCBondStatusRef) corefoundation.CFArrayRef {
	result, callErr := trySCBondStatusGetMemberInterfaces(bondStatus)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCBondStatusGetTypeID func() uint
var _sCBondStatusGetTypeIDErr error

func trySCBondStatusGetTypeID() (uint, error) {
	if _sCBondStatusGetTypeID == nil {
		return 0, symbolCallError("SCBondStatusGetTypeID", "10.5", _sCBondStatusGetTypeIDErr)
	}
	return _sCBondStatusGetTypeID(), nil
}

// SCBondStatusGetTypeID returns the type identifier of all [SCBondStatusRef] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetTypeID()
func SCBondStatusGetTypeID() uint {
	result, callErr := trySCBondStatusGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCCopyLastError func() corefoundation.CFErrorRef
var _sCCopyLastErrorErr error

func trySCCopyLastError() (corefoundation.CFErrorRef, error) {
	if _sCCopyLastError == nil {
		return *new(corefoundation.CFErrorRef), symbolCallError("SCCopyLastError", "10.5", _sCCopyLastErrorErr)
	}
	return _sCCopyLastError(), nil
}

// SCCopyLastError returns an error or status code associated with the most recent function call.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCCopyLastError()
func SCCopyLastError() corefoundation.CFErrorRef {
	result, callErr := trySCCopyLastError()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreAddTemporaryValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool
var _sCDynamicStoreAddTemporaryValueErr error

func trySCDynamicStoreAddTemporaryValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) (bool, error) {
	if _sCDynamicStoreAddTemporaryValue == nil {
		return false, symbolCallError("SCDynamicStoreAddTemporaryValue", "10.1", _sCDynamicStoreAddTemporaryValueErr)
	}
	return _sCDynamicStoreAddTemporaryValue(store, key, value), nil
}

// SCDynamicStoreAddTemporaryValue temporarily adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddTemporaryValue(_:_:_:)
func SCDynamicStoreAddTemporaryValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool {
	result, callErr := trySCDynamicStoreAddTemporaryValue(store, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreAddValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool
var _sCDynamicStoreAddValueErr error

func trySCDynamicStoreAddValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) (bool, error) {
	if _sCDynamicStoreAddValue == nil {
		return false, symbolCallError("SCDynamicStoreAddValue", "10.1", _sCDynamicStoreAddValueErr)
	}
	return _sCDynamicStoreAddValue(store, key, value), nil
}

// SCDynamicStoreAddValue adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddValue(_:_:_:)
func SCDynamicStoreAddValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool {
	result, callErr := trySCDynamicStoreAddValue(store, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyComputerName func(store SCDynamicStoreRef, nameEncoding *uint32) corefoundation.CFStringRef
var _sCDynamicStoreCopyComputerNameErr error

func trySCDynamicStoreCopyComputerName(store SCDynamicStoreRef, nameEncoding *uint32) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreCopyComputerName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreCopyComputerName", "10.1", _sCDynamicStoreCopyComputerNameErr)
	}
	return _sCDynamicStoreCopyComputerName(store, nameEncoding), nil
}

// SCDynamicStoreCopyComputerName returns the current computer name.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyComputerName(_:_:)
func SCDynamicStoreCopyComputerName(store SCDynamicStoreRef, nameEncoding *uint32) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreCopyComputerName(store, nameEncoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyConsoleUser func(store SCDynamicStoreRef, uid *uint32, gid *uint32) corefoundation.CFStringRef
var _sCDynamicStoreCopyConsoleUserErr error

func trySCDynamicStoreCopyConsoleUser(store SCDynamicStoreRef, uid *uint32, gid *uint32) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreCopyConsoleUser == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreCopyConsoleUser", "10.1", _sCDynamicStoreCopyConsoleUserErr)
	}
	return _sCDynamicStoreCopyConsoleUser(store, uid, gid), nil
}

// SCDynamicStoreCopyConsoleUser returns information about the user currently logged into the system.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyConsoleUser(_:_:_:)
func SCDynamicStoreCopyConsoleUser(store SCDynamicStoreRef, uid *uint32, gid *uint32) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreCopyConsoleUser(store, uid, gid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyDHCPInfo func(store SCDynamicStoreRef, serviceID corefoundation.CFStringRef) corefoundation.CFDictionaryRef
var _sCDynamicStoreCopyDHCPInfoErr error

func trySCDynamicStoreCopyDHCPInfo(store SCDynamicStoreRef, serviceID corefoundation.CFStringRef) (corefoundation.CFDictionaryRef, error) {
	if _sCDynamicStoreCopyDHCPInfo == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCDynamicStoreCopyDHCPInfo", "10.1", _sCDynamicStoreCopyDHCPInfoErr)
	}
	return _sCDynamicStoreCopyDHCPInfo(store, serviceID), nil
}

// SCDynamicStoreCopyDHCPInfo returns the DHCP information for the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyDHCPInfo
func SCDynamicStoreCopyDHCPInfo(store SCDynamicStoreRef, serviceID corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCDynamicStoreCopyDHCPInfo(store, serviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyKeyList func(store SCDynamicStoreRef, pattern corefoundation.CFStringRef) corefoundation.CFArrayRef
var _sCDynamicStoreCopyKeyListErr error

func trySCDynamicStoreCopyKeyList(store SCDynamicStoreRef, pattern corefoundation.CFStringRef) (corefoundation.CFArrayRef, error) {
	if _sCDynamicStoreCopyKeyList == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCDynamicStoreCopyKeyList", "10.1", _sCDynamicStoreCopyKeyListErr)
	}
	return _sCDynamicStoreCopyKeyList(store, pattern), nil
}

// SCDynamicStoreCopyKeyList returns the keys that represent the current dynamic store entries that match the specified pattern.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyKeyList(_:_:)
func SCDynamicStoreCopyKeyList(store SCDynamicStoreRef, pattern corefoundation.CFStringRef) corefoundation.CFArrayRef {
	result, callErr := trySCDynamicStoreCopyKeyList(store, pattern)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyLocalHostName func(store SCDynamicStoreRef) corefoundation.CFStringRef
var _sCDynamicStoreCopyLocalHostNameErr error

func trySCDynamicStoreCopyLocalHostName(store SCDynamicStoreRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreCopyLocalHostName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreCopyLocalHostName", "10.1", _sCDynamicStoreCopyLocalHostNameErr)
	}
	return _sCDynamicStoreCopyLocalHostName(store), nil
}

// SCDynamicStoreCopyLocalHostName returns the current local host name.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocalHostName(_:)
func SCDynamicStoreCopyLocalHostName(store SCDynamicStoreRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreCopyLocalHostName(store)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyLocation func(store SCDynamicStoreRef) corefoundation.CFStringRef
var _sCDynamicStoreCopyLocationErr error

func trySCDynamicStoreCopyLocation(store SCDynamicStoreRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreCopyLocation == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreCopyLocation", "10.1", _sCDynamicStoreCopyLocationErr)
	}
	return _sCDynamicStoreCopyLocation(store), nil
}

// SCDynamicStoreCopyLocation returns the current location identifier.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocation(_:)
func SCDynamicStoreCopyLocation(store SCDynamicStoreRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreCopyLocation(store)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyMultiple func(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) corefoundation.CFDictionaryRef
var _sCDynamicStoreCopyMultipleErr error

func trySCDynamicStoreCopyMultiple(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) (corefoundation.CFDictionaryRef, error) {
	if _sCDynamicStoreCopyMultiple == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCDynamicStoreCopyMultiple", "10.1", _sCDynamicStoreCopyMultipleErr)
	}
	return _sCDynamicStoreCopyMultiple(store, keys, patterns), nil
}

// SCDynamicStoreCopyMultiple returns the key-value pairs that match the specified keys and key patterns.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyMultiple(_:_:_:)
func SCDynamicStoreCopyMultiple(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCDynamicStoreCopyMultiple(store, keys, patterns)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyNotifiedKeys func(store SCDynamicStoreRef) corefoundation.CFArrayRef
var _sCDynamicStoreCopyNotifiedKeysErr error

func trySCDynamicStoreCopyNotifiedKeys(store SCDynamicStoreRef) (corefoundation.CFArrayRef, error) {
	if _sCDynamicStoreCopyNotifiedKeys == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCDynamicStoreCopyNotifiedKeys", "10.1", _sCDynamicStoreCopyNotifiedKeysErr)
	}
	return _sCDynamicStoreCopyNotifiedKeys(store), nil
}

// SCDynamicStoreCopyNotifiedKeys returns the keys that have changed since the last call to this function.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyNotifiedKeys(_:)
func SCDynamicStoreCopyNotifiedKeys(store SCDynamicStoreRef) corefoundation.CFArrayRef {
	result, callErr := trySCDynamicStoreCopyNotifiedKeys(store)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyProxies func(store SCDynamicStoreRef) corefoundation.CFDictionaryRef
var _sCDynamicStoreCopyProxiesErr error

func trySCDynamicStoreCopyProxies(store SCDynamicStoreRef) (corefoundation.CFDictionaryRef, error) {
	if _sCDynamicStoreCopyProxies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCDynamicStoreCopyProxies", "10.1", _sCDynamicStoreCopyProxiesErr)
	}
	return _sCDynamicStoreCopyProxies(store), nil
}

// SCDynamicStoreCopyProxies returns the key-value pairs that represent the current internet proxy settings.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyProxies(_:)
func SCDynamicStoreCopyProxies(store SCDynamicStoreRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCDynamicStoreCopyProxies(store)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCopyValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef) corefoundation.CFPropertyListRef
var _sCDynamicStoreCopyValueErr error

func trySCDynamicStoreCopyValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) (corefoundation.CFPropertyListRef, error) {
	if _sCDynamicStoreCopyValue == nil {
		return *new(corefoundation.CFPropertyListRef), symbolCallError("SCDynamicStoreCopyValue", "10.1", _sCDynamicStoreCopyValueErr)
	}
	return _sCDynamicStoreCopyValue(store, key), nil
}

// SCDynamicStoreCopyValue returns the value associated with the specified key.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyValue(_:_:)
func SCDynamicStoreCopyValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) corefoundation.CFPropertyListRef {
	result, callErr := trySCDynamicStoreCopyValue(store, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCreate func(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) SCDynamicStoreRef
var _sCDynamicStoreCreateErr error

func trySCDynamicStoreCreate(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) (SCDynamicStoreRef, error) {
	if _sCDynamicStoreCreate == nil {
		return *new(SCDynamicStoreRef), symbolCallError("SCDynamicStoreCreate", "10.1", _sCDynamicStoreCreateErr)
	}
	return _sCDynamicStoreCreate(allocator, name, callout, context), nil
}

// SCDynamicStoreCreate creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreate(_:_:_:_:)
func SCDynamicStoreCreate(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) SCDynamicStoreRef {
	result, callErr := trySCDynamicStoreCreate(allocator, name, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCreateRunLoopSource func(allocator corefoundation.CFAllocatorRef, store SCDynamicStoreRef, order int) corefoundation.CFRunLoopSourceRef
var _sCDynamicStoreCreateRunLoopSourceErr error

func trySCDynamicStoreCreateRunLoopSource(allocator corefoundation.CFAllocatorRef, store SCDynamicStoreRef, order int) (corefoundation.CFRunLoopSourceRef, error) {
	if _sCDynamicStoreCreateRunLoopSource == nil {
		return *new(corefoundation.CFRunLoopSourceRef), symbolCallError("SCDynamicStoreCreateRunLoopSource", "10.1", _sCDynamicStoreCreateRunLoopSourceErr)
	}
	return _sCDynamicStoreCreateRunLoopSource(allocator, store, order), nil
}

// SCDynamicStoreCreateRunLoopSource creates a run loop source object that can be added to the application’s run loop.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateRunLoopSource(_:_:_:)
func SCDynamicStoreCreateRunLoopSource(allocator corefoundation.CFAllocatorRef, store SCDynamicStoreRef, order int) corefoundation.CFRunLoopSourceRef {
	result, callErr := trySCDynamicStoreCreateRunLoopSource(allocator, store, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreCreateWithOptions func(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, storeOptions corefoundation.CFDictionaryRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) SCDynamicStoreRef
var _sCDynamicStoreCreateWithOptionsErr error

func trySCDynamicStoreCreateWithOptions(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, storeOptions corefoundation.CFDictionaryRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) (SCDynamicStoreRef, error) {
	if _sCDynamicStoreCreateWithOptions == nil {
		return *new(SCDynamicStoreRef), symbolCallError("SCDynamicStoreCreateWithOptions", "10.4", _sCDynamicStoreCreateWithOptionsErr)
	}
	return _sCDynamicStoreCreateWithOptions(allocator, name, storeOptions, callout, context), nil
}

// SCDynamicStoreCreateWithOptions creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateWithOptions(_:_:_:_:_:)
func SCDynamicStoreCreateWithOptions(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, storeOptions corefoundation.CFDictionaryRef, callout SCDynamicStoreCallBack, context *SCDynamicStoreContext) SCDynamicStoreRef {
	result, callErr := trySCDynamicStoreCreateWithOptions(allocator, name, storeOptions, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreGetTypeID func() uint
var _sCDynamicStoreGetTypeIDErr error

func trySCDynamicStoreGetTypeID() (uint, error) {
	if _sCDynamicStoreGetTypeID == nil {
		return 0, symbolCallError("SCDynamicStoreGetTypeID", "10.1", _sCDynamicStoreGetTypeIDErr)
	}
	return _sCDynamicStoreGetTypeID(), nil
}

// SCDynamicStoreGetTypeID returns the type identifier of all [SCDynamicStore] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreGetTypeID()
func SCDynamicStoreGetTypeID() uint {
	result, callErr := trySCDynamicStoreGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreate func(allocator corefoundation.CFAllocatorRef, fmt corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateErr error

func trySCDynamicStoreKeyCreate(allocator corefoundation.CFAllocatorRef, fmt corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreate == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreate", "10.1", _sCDynamicStoreKeyCreateErr)
	}
	return _sCDynamicStoreKeyCreate(allocator, fmt), nil
}

// SCDynamicStoreKeyCreate creates a dynamic store key using the specified format.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreate
func SCDynamicStoreKeyCreate(allocator corefoundation.CFAllocatorRef, fmt corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreate(allocator, fmt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateComputerName func(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateComputerNameErr error

func trySCDynamicStoreKeyCreateComputerName(allocator corefoundation.CFAllocatorRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateComputerName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateComputerName", "10.1", _sCDynamicStoreKeyCreateComputerNameErr)
	}
	return _sCDynamicStoreKeyCreateComputerName(allocator), nil
}

// SCDynamicStoreKeyCreateComputerName creates a key that can be used to receive notifications when the current computer name changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateComputerName(_:)
func SCDynamicStoreKeyCreateComputerName(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateComputerName(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateConsoleUser func(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateConsoleUserErr error

func trySCDynamicStoreKeyCreateConsoleUser(allocator corefoundation.CFAllocatorRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateConsoleUser == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateConsoleUser", "10.1", _sCDynamicStoreKeyCreateConsoleUserErr)
	}
	return _sCDynamicStoreKeyCreateConsoleUser(allocator), nil
}

// SCDynamicStoreKeyCreateConsoleUser creates a key that can be used to receive notifications when the current console user changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateConsoleUser(_:)
func SCDynamicStoreKeyCreateConsoleUser(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateConsoleUser(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateHostNames func(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateHostNamesErr error

func trySCDynamicStoreKeyCreateHostNames(allocator corefoundation.CFAllocatorRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateHostNames == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateHostNames", "10.2", _sCDynamicStoreKeyCreateHostNamesErr)
	}
	return _sCDynamicStoreKeyCreateHostNames(allocator), nil
}

// SCDynamicStoreKeyCreateHostNames creates a key that can be used to receive notifications when the [HostNames] entity changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateHostNames(_:)
func SCDynamicStoreKeyCreateHostNames(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateHostNames(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateLocation func(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateLocationErr error

func trySCDynamicStoreKeyCreateLocation(allocator corefoundation.CFAllocatorRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateLocation == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateLocation", "10.2", _sCDynamicStoreKeyCreateLocationErr)
	}
	return _sCDynamicStoreKeyCreateLocation(allocator), nil
}

// SCDynamicStoreKeyCreateLocation creates a key that can be used to receive notifications when the location identifier changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateLocation(_:)
func SCDynamicStoreKeyCreateLocation(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateLocation(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateNetworkGlobalEntity func(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateNetworkGlobalEntityErr error

func trySCDynamicStoreKeyCreateNetworkGlobalEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, entity corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateNetworkGlobalEntity == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateNetworkGlobalEntity", "10.1", _sCDynamicStoreKeyCreateNetworkGlobalEntityErr)
	}
	return _sCDynamicStoreKeyCreateNetworkGlobalEntity(allocator, domain, entity), nil
}

// SCDynamicStoreKeyCreateNetworkGlobalEntity creates a dynamic store key that can be used to access a specific global (as opposed to a per-service or per-interface) network configuration entity.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkGlobalEntity(_:_:_:)
func SCDynamicStoreKeyCreateNetworkGlobalEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateNetworkGlobalEntity(allocator, domain, entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateNetworkInterface func(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateNetworkInterfaceErr error

func trySCDynamicStoreKeyCreateNetworkInterface(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateNetworkInterface == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateNetworkInterface", "10.1", _sCDynamicStoreKeyCreateNetworkInterfaceErr)
	}
	return _sCDynamicStoreKeyCreateNetworkInterface(allocator, domain), nil
}

// SCDynamicStoreKeyCreateNetworkInterface creates a dynamic store key that can be used to access the network interface configuration information in the dynamic store.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterface(_:_:)
func SCDynamicStoreKeyCreateNetworkInterface(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateNetworkInterface(allocator, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateNetworkInterfaceEntity func(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, ifname corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateNetworkInterfaceEntityErr error

func trySCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, ifname corefoundation.CFStringRef, entity corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateNetworkInterfaceEntity == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateNetworkInterfaceEntity", "10.1", _sCDynamicStoreKeyCreateNetworkInterfaceEntityErr)
	}
	return _sCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator, domain, ifname, entity), nil
}

// SCDynamicStoreKeyCreateNetworkInterfaceEntity creates a dynamic store key that can be used to access the per-interface network configuration information in the dynamic store.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterfaceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, ifname corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator, domain, ifname, entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateNetworkServiceEntity func(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, serviceID corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateNetworkServiceEntityErr error

func trySCDynamicStoreKeyCreateNetworkServiceEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, serviceID corefoundation.CFStringRef, entity corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateNetworkServiceEntity == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateNetworkServiceEntity", "10.1", _sCDynamicStoreKeyCreateNetworkServiceEntityErr)
	}
	return _sCDynamicStoreKeyCreateNetworkServiceEntity(allocator, domain, serviceID, entity), nil
}

// SCDynamicStoreKeyCreateNetworkServiceEntity creates a dynamic store key that can be used to access the per-service network configuration information.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkServiceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkServiceEntity(allocator corefoundation.CFAllocatorRef, domain corefoundation.CFStringRef, serviceID corefoundation.CFStringRef, entity corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateNetworkServiceEntity(allocator, domain, serviceID, entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreKeyCreateProxies func(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef
var _sCDynamicStoreKeyCreateProxiesErr error

func trySCDynamicStoreKeyCreateProxies(allocator corefoundation.CFAllocatorRef) (corefoundation.CFStringRef, error) {
	if _sCDynamicStoreKeyCreateProxies == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCDynamicStoreKeyCreateProxies", "10.1", _sCDynamicStoreKeyCreateProxiesErr)
	}
	return _sCDynamicStoreKeyCreateProxies(allocator), nil
}

// SCDynamicStoreKeyCreateProxies creates a key that can be used to receive notifications when the current network proxy settings are changed.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateProxies(_:)
func SCDynamicStoreKeyCreateProxies(allocator corefoundation.CFAllocatorRef) corefoundation.CFStringRef {
	result, callErr := trySCDynamicStoreKeyCreateProxies(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreNotifyValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef) bool
var _sCDynamicStoreNotifyValueErr error

func trySCDynamicStoreNotifyValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) (bool, error) {
	if _sCDynamicStoreNotifyValue == nil {
		return false, symbolCallError("SCDynamicStoreNotifyValue", "10.1", _sCDynamicStoreNotifyValueErr)
	}
	return _sCDynamicStoreNotifyValue(store, key), nil
}

// SCDynamicStoreNotifyValue causes a notification to be delivered for the specified key in the dynamic store.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreNotifyValue(_:_:)
func SCDynamicStoreNotifyValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) bool {
	result, callErr := trySCDynamicStoreNotifyValue(store, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreRemoveValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef) bool
var _sCDynamicStoreRemoveValueErr error

func trySCDynamicStoreRemoveValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) (bool, error) {
	if _sCDynamicStoreRemoveValue == nil {
		return false, symbolCallError("SCDynamicStoreRemoveValue", "10.1", _sCDynamicStoreRemoveValueErr)
	}
	return _sCDynamicStoreRemoveValue(store, key), nil
}

// SCDynamicStoreRemoveValue removes the value of the specified key from the dynamic store.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreRemoveValue(_:_:)
func SCDynamicStoreRemoveValue(store SCDynamicStoreRef, key corefoundation.CFStringRef) bool {
	result, callErr := trySCDynamicStoreRemoveValue(store, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreSetDispatchQueue func(store SCDynamicStoreRef, queue uintptr) bool
var _sCDynamicStoreSetDispatchQueueErr error

func trySCDynamicStoreSetDispatchQueue(store SCDynamicStoreRef, queue dispatch.Queue) (bool, error) {
	if _sCDynamicStoreSetDispatchQueue == nil {
		return false, symbolCallError("SCDynamicStoreSetDispatchQueue", "10.6", _sCDynamicStoreSetDispatchQueueErr)
	}
	return _sCDynamicStoreSetDispatchQueue(store, uintptr(queue.Handle())), nil
}

// SCDynamicStoreSetDispatchQueue initiates notifications for the notification keys, using the specified dispatch queue for the callback.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetDispatchQueue(_:_:)
func SCDynamicStoreSetDispatchQueue(store SCDynamicStoreRef, queue dispatch.Queue) bool {
	result, callErr := trySCDynamicStoreSetDispatchQueue(store, queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreSetMultiple func(store SCDynamicStoreRef, keysToSet corefoundation.CFDictionaryRef, keysToRemove corefoundation.CFArrayRef, keysToNotify corefoundation.CFArrayRef) bool
var _sCDynamicStoreSetMultipleErr error

func trySCDynamicStoreSetMultiple(store SCDynamicStoreRef, keysToSet corefoundation.CFDictionaryRef, keysToRemove corefoundation.CFArrayRef, keysToNotify corefoundation.CFArrayRef) (bool, error) {
	if _sCDynamicStoreSetMultiple == nil {
		return false, symbolCallError("SCDynamicStoreSetMultiple", "10.1", _sCDynamicStoreSetMultipleErr)
	}
	return _sCDynamicStoreSetMultiple(store, keysToSet, keysToRemove, keysToNotify), nil
}

// SCDynamicStoreSetMultiple updates multiple values in the dynamic store.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetMultiple(_:_:_:_:)
func SCDynamicStoreSetMultiple(store SCDynamicStoreRef, keysToSet corefoundation.CFDictionaryRef, keysToRemove corefoundation.CFArrayRef, keysToNotify corefoundation.CFArrayRef) bool {
	result, callErr := trySCDynamicStoreSetMultiple(store, keysToSet, keysToRemove, keysToNotify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreSetNotificationKeys func(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) bool
var _sCDynamicStoreSetNotificationKeysErr error

func trySCDynamicStoreSetNotificationKeys(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) (bool, error) {
	if _sCDynamicStoreSetNotificationKeys == nil {
		return false, symbolCallError("SCDynamicStoreSetNotificationKeys", "10.1", _sCDynamicStoreSetNotificationKeysErr)
	}
	return _sCDynamicStoreSetNotificationKeys(store, keys, patterns), nil
}

// SCDynamicStoreSetNotificationKeys specifies a set of keys and key patterns that should be monitored for changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetNotificationKeys(_:_:_:)
func SCDynamicStoreSetNotificationKeys(store SCDynamicStoreRef, keys corefoundation.CFArrayRef, patterns corefoundation.CFArrayRef) bool {
	result, callErr := trySCDynamicStoreSetNotificationKeys(store, keys, patterns)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCDynamicStoreSetValue func(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool
var _sCDynamicStoreSetValueErr error

func trySCDynamicStoreSetValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) (bool, error) {
	if _sCDynamicStoreSetValue == nil {
		return false, symbolCallError("SCDynamicStoreSetValue", "10.1", _sCDynamicStoreSetValueErr)
	}
	return _sCDynamicStoreSetValue(store, key, value), nil
}

// SCDynamicStoreSetValue adds or replaces a value in the dynamic store for the specified key.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetValue(_:_:_:)
func SCDynamicStoreSetValue(store SCDynamicStoreRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool {
	result, callErr := trySCDynamicStoreSetValue(store, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCError func() int32
var _sCErrorErr error

func trySCError() (int32, error) {
	if _sCError == nil {
		return 0, symbolCallError("SCError", "10.1", _sCErrorErr)
	}
	return _sCError(), nil
}

// SCError returns an error or status code associated with the most recent function call.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCError()
func SCError() int32 {
	result, callErr := trySCError()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCErrorString func(status int32) *byte
var _sCErrorStringErr error

func trySCErrorString(status int32) (*byte, error) {
	if _sCErrorString == nil {
		return nil, symbolCallError("SCErrorString", "10.1", _sCErrorStringErr)
	}
	return _sCErrorString(status), nil
}

// SCErrorString returns a string describing the specified status code or error code.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCErrorString(_:)
func SCErrorString(status int32) *byte {
	result, callErr := trySCErrorString(status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCopyExtendedStatus func(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef
var _sCNetworkConnectionCopyExtendedStatusErr error

func trySCNetworkConnectionCopyExtendedStatus(connection SCNetworkConnectionRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkConnectionCopyExtendedStatus == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkConnectionCopyExtendedStatus", "10.3", _sCNetworkConnectionCopyExtendedStatusErr)
	}
	return _sCNetworkConnectionCopyExtendedStatus(connection), nil
}

// SCNetworkConnectionCopyExtendedStatus returns the extended status of the connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyExtendedStatus(_:)
func SCNetworkConnectionCopyExtendedStatus(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkConnectionCopyExtendedStatus(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCopyServiceID func(connection SCNetworkConnectionRef) corefoundation.CFStringRef
var _sCNetworkConnectionCopyServiceIDErr error

func trySCNetworkConnectionCopyServiceID(connection SCNetworkConnectionRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkConnectionCopyServiceID == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkConnectionCopyServiceID", "10.3", _sCNetworkConnectionCopyServiceIDErr)
	}
	return _sCNetworkConnectionCopyServiceID(connection), nil
}

// SCNetworkConnectionCopyServiceID returns the service ID associated with the specified network connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyServiceID(_:)
func SCNetworkConnectionCopyServiceID(connection SCNetworkConnectionRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkConnectionCopyServiceID(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCopyStatistics func(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef
var _sCNetworkConnectionCopyStatisticsErr error

func trySCNetworkConnectionCopyStatistics(connection SCNetworkConnectionRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkConnectionCopyStatistics == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkConnectionCopyStatistics", "10.3", _sCNetworkConnectionCopyStatisticsErr)
	}
	return _sCNetworkConnectionCopyStatistics(connection), nil
}

// SCNetworkConnectionCopyStatistics returns the statistics of the specified connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyStatistics(_:)
func SCNetworkConnectionCopyStatistics(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkConnectionCopyStatistics(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCopyUserOptions func(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef
var _sCNetworkConnectionCopyUserOptionsErr error

func trySCNetworkConnectionCopyUserOptions(connection SCNetworkConnectionRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkConnectionCopyUserOptions == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkConnectionCopyUserOptions", "10.3", _sCNetworkConnectionCopyUserOptionsErr)
	}
	return _sCNetworkConnectionCopyUserOptions(connection), nil
}

// SCNetworkConnectionCopyUserOptions gets the user options used to start the specified connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserOptions(_:)
func SCNetworkConnectionCopyUserOptions(connection SCNetworkConnectionRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkConnectionCopyUserOptions(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCopyUserPreferences func(selectionOptions corefoundation.CFDictionaryRef, serviceID *corefoundation.CFStringRef, userOptions *corefoundation.CFDictionaryRef) bool
var _sCNetworkConnectionCopyUserPreferencesErr error

func trySCNetworkConnectionCopyUserPreferences(selectionOptions corefoundation.CFDictionaryRef, serviceID *corefoundation.CFStringRef, userOptions *corefoundation.CFDictionaryRef) (bool, error) {
	if _sCNetworkConnectionCopyUserPreferences == nil {
		return false, symbolCallError("SCNetworkConnectionCopyUserPreferences", "10.3", _sCNetworkConnectionCopyUserPreferencesErr)
	}
	return _sCNetworkConnectionCopyUserPreferences(selectionOptions, serviceID, userOptions), nil
}

// SCNetworkConnectionCopyUserPreferences provides the default service ID and a dictionary of user options for the specified connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserPreferences(_:_:_:)
func SCNetworkConnectionCopyUserPreferences(selectionOptions corefoundation.CFDictionaryRef, serviceID *corefoundation.CFStringRef, userOptions *corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCNetworkConnectionCopyUserPreferences(selectionOptions, serviceID, userOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionCreateWithServiceID func(allocator corefoundation.CFAllocatorRef, serviceID corefoundation.CFStringRef, callout SCNetworkConnectionCallBack, context *SCNetworkConnectionContext) SCNetworkConnectionRef
var _sCNetworkConnectionCreateWithServiceIDErr error

func trySCNetworkConnectionCreateWithServiceID(allocator corefoundation.CFAllocatorRef, serviceID corefoundation.CFStringRef, callout SCNetworkConnectionCallBack, context *SCNetworkConnectionContext) (SCNetworkConnectionRef, error) {
	if _sCNetworkConnectionCreateWithServiceID == nil {
		return *new(SCNetworkConnectionRef), symbolCallError("SCNetworkConnectionCreateWithServiceID", "10.3", _sCNetworkConnectionCreateWithServiceIDErr)
	}
	return _sCNetworkConnectionCreateWithServiceID(allocator, serviceID, callout, context), nil
}

// SCNetworkConnectionCreateWithServiceID creates a new connection reference to use for getting the status or for connecting or disconnecting the associated service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCreateWithServiceID(_:_:_:_:)
func SCNetworkConnectionCreateWithServiceID(allocator corefoundation.CFAllocatorRef, serviceID corefoundation.CFStringRef, callout SCNetworkConnectionCallBack, context *SCNetworkConnectionContext) SCNetworkConnectionRef {
	result, callErr := trySCNetworkConnectionCreateWithServiceID(allocator, serviceID, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionGetStatus func(connection SCNetworkConnectionRef) SCNetworkConnectionStatus
var _sCNetworkConnectionGetStatusErr error

func trySCNetworkConnectionGetStatus(connection SCNetworkConnectionRef) (SCNetworkConnectionStatus, error) {
	if _sCNetworkConnectionGetStatus == nil {
		return *new(SCNetworkConnectionStatus), symbolCallError("SCNetworkConnectionGetStatus", "10.3", _sCNetworkConnectionGetStatusErr)
	}
	return _sCNetworkConnectionGetStatus(connection), nil
}

// SCNetworkConnectionGetStatus returns the status of the specified network connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetStatus(_:)
func SCNetworkConnectionGetStatus(connection SCNetworkConnectionRef) SCNetworkConnectionStatus {
	result, callErr := trySCNetworkConnectionGetStatus(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionGetTypeID func() uint
var _sCNetworkConnectionGetTypeIDErr error

func trySCNetworkConnectionGetTypeID() (uint, error) {
	if _sCNetworkConnectionGetTypeID == nil {
		return 0, symbolCallError("SCNetworkConnectionGetTypeID", "10.3", _sCNetworkConnectionGetTypeIDErr)
	}
	return _sCNetworkConnectionGetTypeID(), nil
}

// SCNetworkConnectionGetTypeID returns the type identifier of all [SCNetworkConnection] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetTypeID()
func SCNetworkConnectionGetTypeID() uint {
	result, callErr := trySCNetworkConnectionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionScheduleWithRunLoop func(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCNetworkConnectionScheduleWithRunLoopErr error

func trySCNetworkConnectionScheduleWithRunLoop(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkConnectionScheduleWithRunLoop == nil {
		return false, symbolCallError("SCNetworkConnectionScheduleWithRunLoop", "10.3", _sCNetworkConnectionScheduleWithRunLoopErr)
	}
	return _sCNetworkConnectionScheduleWithRunLoop(connection, runLoop, runLoopMode), nil
}

// SCNetworkConnectionScheduleWithRunLoop schedules the specified connection with the specified run loop.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionScheduleWithRunLoop(_:_:_:)
func SCNetworkConnectionScheduleWithRunLoop(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkConnectionScheduleWithRunLoop(connection, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionSetDispatchQueue func(connection SCNetworkConnectionRef, queue uintptr) bool
var _sCNetworkConnectionSetDispatchQueueErr error

func trySCNetworkConnectionSetDispatchQueue(connection SCNetworkConnectionRef, queue dispatch.Queue) (bool, error) {
	if _sCNetworkConnectionSetDispatchQueue == nil {
		return false, symbolCallError("SCNetworkConnectionSetDispatchQueue", "10.6", _sCNetworkConnectionSetDispatchQueueErr)
	}
	return _sCNetworkConnectionSetDispatchQueue(connection, uintptr(queue.Handle())), nil
}

// SCNetworkConnectionSetDispatchQueue specifies a dispatch queue to use for the connection’s callback function and enables notifications.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionSetDispatchQueue(_:_:)
func SCNetworkConnectionSetDispatchQueue(connection SCNetworkConnectionRef, queue dispatch.Queue) bool {
	result, callErr := trySCNetworkConnectionSetDispatchQueue(connection, queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionStart func(connection SCNetworkConnectionRef, userOptions corefoundation.CFDictionaryRef, linger bool) bool
var _sCNetworkConnectionStartErr error

func trySCNetworkConnectionStart(connection SCNetworkConnectionRef, userOptions corefoundation.CFDictionaryRef, linger bool) (bool, error) {
	if _sCNetworkConnectionStart == nil {
		return false, symbolCallError("SCNetworkConnectionStart", "10.3", _sCNetworkConnectionStartErr)
	}
	return _sCNetworkConnectionStart(connection, userOptions, linger), nil
}

// SCNetworkConnectionStart starts the connection process for the specified network connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStart(_:_:_:)
func SCNetworkConnectionStart(connection SCNetworkConnectionRef, userOptions corefoundation.CFDictionaryRef, linger bool) bool {
	result, callErr := trySCNetworkConnectionStart(connection, userOptions, linger)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionStop func(connection SCNetworkConnectionRef, forceDisconnect bool) bool
var _sCNetworkConnectionStopErr error

func trySCNetworkConnectionStop(connection SCNetworkConnectionRef, forceDisconnect bool) (bool, error) {
	if _sCNetworkConnectionStop == nil {
		return false, symbolCallError("SCNetworkConnectionStop", "10.3", _sCNetworkConnectionStopErr)
	}
	return _sCNetworkConnectionStop(connection, forceDisconnect), nil
}

// SCNetworkConnectionStop stops the connection process for the specified network connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStop(_:_:)
func SCNetworkConnectionStop(connection SCNetworkConnectionRef, forceDisconnect bool) bool {
	result, callErr := trySCNetworkConnectionStop(connection, forceDisconnect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkConnectionUnscheduleFromRunLoop func(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCNetworkConnectionUnscheduleFromRunLoopErr error

func trySCNetworkConnectionUnscheduleFromRunLoop(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkConnectionUnscheduleFromRunLoop == nil {
		return false, symbolCallError("SCNetworkConnectionUnscheduleFromRunLoop", "10.3", _sCNetworkConnectionUnscheduleFromRunLoopErr)
	}
	return _sCNetworkConnectionUnscheduleFromRunLoop(connection, runLoop, runLoopMode), nil
}

// SCNetworkConnectionUnscheduleFromRunLoop unschedules the specified connection from the specified run loop.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionUnscheduleFromRunLoop(_:_:_:)
func SCNetworkConnectionUnscheduleFromRunLoop(connection SCNetworkConnectionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkConnectionUnscheduleFromRunLoop(connection, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCopyAll func() corefoundation.CFArrayRef
var _sCNetworkInterfaceCopyAllErr error

func trySCNetworkInterfaceCopyAll() (corefoundation.CFArrayRef, error) {
	if _sCNetworkInterfaceCopyAll == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkInterfaceCopyAll", "10.4", _sCNetworkInterfaceCopyAllErr)
	}
	return _sCNetworkInterfaceCopyAll(), nil
}

// SCNetworkInterfaceCopyAll returns all network-capable interfaces on the system.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyAll()
func SCNetworkInterfaceCopyAll() corefoundation.CFArrayRef {
	result, callErr := trySCNetworkInterfaceCopyAll()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCopyMTU func(interface_ SCNetworkInterfaceRef, mtu_cur *int32, mtu_min *int32, mtu_max *int32) bool
var _sCNetworkInterfaceCopyMTUErr error

func trySCNetworkInterfaceCopyMTU(interface_ SCNetworkInterfaceRef, mtu_cur *int32, mtu_min *int32, mtu_max *int32) (bool, error) {
	if _sCNetworkInterfaceCopyMTU == nil {
		return false, symbolCallError("SCNetworkInterfaceCopyMTU", "10.5", _sCNetworkInterfaceCopyMTUErr)
	}
	return _sCNetworkInterfaceCopyMTU(interface_, mtu_cur, mtu_min, mtu_max), nil
}

// SCNetworkInterfaceCopyMTU returns the current MTU setting and the range of allowable values for the specified network interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMTU(_:_:_:_:)
func SCNetworkInterfaceCopyMTU(interface_ SCNetworkInterfaceRef, mtu_cur *int32, mtu_min *int32, mtu_max *int32) bool {
	result, callErr := trySCNetworkInterfaceCopyMTU(interface_, mtu_cur, mtu_min, mtu_max)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCopyMediaOptions func(interface_ SCNetworkInterfaceRef, current *corefoundation.CFDictionaryRef, active *corefoundation.CFDictionaryRef, available *corefoundation.CFArrayRef, filter bool) bool
var _sCNetworkInterfaceCopyMediaOptionsErr error

func trySCNetworkInterfaceCopyMediaOptions(interface_ SCNetworkInterfaceRef, current *corefoundation.CFDictionaryRef, active *corefoundation.CFDictionaryRef, available *corefoundation.CFArrayRef, filter bool) (bool, error) {
	if _sCNetworkInterfaceCopyMediaOptions == nil {
		return false, symbolCallError("SCNetworkInterfaceCopyMediaOptions", "10.5", _sCNetworkInterfaceCopyMediaOptionsErr)
	}
	return _sCNetworkInterfaceCopyMediaOptions(interface_, current, active, available, filter), nil
}

// SCNetworkInterfaceCopyMediaOptions returns information media options for the specified network interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaOptions(_:_:_:_:_:)
func SCNetworkInterfaceCopyMediaOptions(interface_ SCNetworkInterfaceRef, current *corefoundation.CFDictionaryRef, active *corefoundation.CFDictionaryRef, available *corefoundation.CFArrayRef, filter bool) bool {
	result, callErr := trySCNetworkInterfaceCopyMediaOptions(interface_, current, active, available, filter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCopyMediaSubTypeOptions func(available corefoundation.CFArrayRef, subType corefoundation.CFStringRef) corefoundation.CFArrayRef
var _sCNetworkInterfaceCopyMediaSubTypeOptionsErr error

func trySCNetworkInterfaceCopyMediaSubTypeOptions(available corefoundation.CFArrayRef, subType corefoundation.CFStringRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkInterfaceCopyMediaSubTypeOptions == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkInterfaceCopyMediaSubTypeOptions", "10.5", _sCNetworkInterfaceCopyMediaSubTypeOptionsErr)
	}
	return _sCNetworkInterfaceCopyMediaSubTypeOptions(available, subType), nil
}

// SCNetworkInterfaceCopyMediaSubTypeOptions returns a list of available media options for the specified interface configuration options and subtype.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypeOptions(_:_:)
func SCNetworkInterfaceCopyMediaSubTypeOptions(available corefoundation.CFArrayRef, subType corefoundation.CFStringRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkInterfaceCopyMediaSubTypeOptions(available, subType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCopyMediaSubTypes func(available corefoundation.CFArrayRef) corefoundation.CFArrayRef
var _sCNetworkInterfaceCopyMediaSubTypesErr error

func trySCNetworkInterfaceCopyMediaSubTypes(available corefoundation.CFArrayRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkInterfaceCopyMediaSubTypes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkInterfaceCopyMediaSubTypes", "10.5", _sCNetworkInterfaceCopyMediaSubTypesErr)
	}
	return _sCNetworkInterfaceCopyMediaSubTypes(available), nil
}

// SCNetworkInterfaceCopyMediaSubTypes returns a list of available media subtypes for the specified interface configuration options.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypes(_:)
func SCNetworkInterfaceCopyMediaSubTypes(available corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkInterfaceCopyMediaSubTypes(available)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceCreateWithInterface func(interface_ SCNetworkInterfaceRef, interfaceType corefoundation.CFStringRef) SCNetworkInterfaceRef
var _sCNetworkInterfaceCreateWithInterfaceErr error

func trySCNetworkInterfaceCreateWithInterface(interface_ SCNetworkInterfaceRef, interfaceType corefoundation.CFStringRef) (SCNetworkInterfaceRef, error) {
	if _sCNetworkInterfaceCreateWithInterface == nil {
		return *new(SCNetworkInterfaceRef), symbolCallError("SCNetworkInterfaceCreateWithInterface", "10.4", _sCNetworkInterfaceCreateWithInterfaceErr)
	}
	return _sCNetworkInterfaceCreateWithInterface(interface_, interfaceType), nil
}

// SCNetworkInterfaceCreateWithInterface creates a new network interface layered on top of the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCreateWithInterface(_:_:)
func SCNetworkInterfaceCreateWithInterface(interface_ SCNetworkInterfaceRef, interfaceType corefoundation.CFStringRef) SCNetworkInterfaceRef {
	result, callErr := trySCNetworkInterfaceCreateWithInterface(interface_, interfaceType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceForceConfigurationRefresh func(interface_ SCNetworkInterfaceRef) bool
var _sCNetworkInterfaceForceConfigurationRefreshErr error

func trySCNetworkInterfaceForceConfigurationRefresh(interface_ SCNetworkInterfaceRef) (bool, error) {
	if _sCNetworkInterfaceForceConfigurationRefresh == nil {
		return false, symbolCallError("SCNetworkInterfaceForceConfigurationRefresh", "10.5", _sCNetworkInterfaceForceConfigurationRefreshErr)
	}
	return _sCNetworkInterfaceForceConfigurationRefresh(interface_), nil
}

// SCNetworkInterfaceForceConfigurationRefresh sends a notification to interested network configuration agents to immediately retry their configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceForceConfigurationRefresh(_:)
func SCNetworkInterfaceForceConfigurationRefresh(interface_ SCNetworkInterfaceRef) bool {
	result, callErr := trySCNetworkInterfaceForceConfigurationRefresh(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetBSDName func(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef
var _sCNetworkInterfaceGetBSDNameErr error

func trySCNetworkInterfaceGetBSDName(interface_ SCNetworkInterfaceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkInterfaceGetBSDName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkInterfaceGetBSDName", "10.4", _sCNetworkInterfaceGetBSDNameErr)
	}
	return _sCNetworkInterfaceGetBSDName(interface_), nil
}

// SCNetworkInterfaceGetBSDName returns the BSD interface or device name for the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetBSDName(_:)
func SCNetworkInterfaceGetBSDName(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkInterfaceGetBSDName(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetConfiguration func(interface_ SCNetworkInterfaceRef) corefoundation.CFDictionaryRef
var _sCNetworkInterfaceGetConfigurationErr error

func trySCNetworkInterfaceGetConfiguration(interface_ SCNetworkInterfaceRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkInterfaceGetConfiguration == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkInterfaceGetConfiguration", "10.4", _sCNetworkInterfaceGetConfigurationErr)
	}
	return _sCNetworkInterfaceGetConfiguration(interface_), nil
}

// SCNetworkInterfaceGetConfiguration returns the configuration settings associated with the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetConfiguration(_:)
func SCNetworkInterfaceGetConfiguration(interface_ SCNetworkInterfaceRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkInterfaceGetConfiguration(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetExtendedConfiguration func(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef) corefoundation.CFDictionaryRef
var _sCNetworkInterfaceGetExtendedConfigurationErr error

func trySCNetworkInterfaceGetExtendedConfiguration(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkInterfaceGetExtendedConfiguration == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkInterfaceGetExtendedConfiguration", "10.5", _sCNetworkInterfaceGetExtendedConfigurationErr)
	}
	return _sCNetworkInterfaceGetExtendedConfiguration(interface_, extendedType), nil
}

// SCNetworkInterfaceGetExtendedConfiguration returns the extended configuration settings associated with the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetExtendedConfiguration(_:_:)
func SCNetworkInterfaceGetExtendedConfiguration(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkInterfaceGetExtendedConfiguration(interface_, extendedType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetHardwareAddressString func(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef
var _sCNetworkInterfaceGetHardwareAddressStringErr error

func trySCNetworkInterfaceGetHardwareAddressString(interface_ SCNetworkInterfaceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkInterfaceGetHardwareAddressString == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkInterfaceGetHardwareAddressString", "10.4", _sCNetworkInterfaceGetHardwareAddressStringErr)
	}
	return _sCNetworkInterfaceGetHardwareAddressString(interface_), nil
}

// SCNetworkInterfaceGetHardwareAddressString returns a displayable link layer address for the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetHardwareAddressString(_:)
func SCNetworkInterfaceGetHardwareAddressString(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkInterfaceGetHardwareAddressString(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetInterface func(interface_ SCNetworkInterfaceRef) SCNetworkInterfaceRef
var _sCNetworkInterfaceGetInterfaceErr error

func trySCNetworkInterfaceGetInterface(interface_ SCNetworkInterfaceRef) (SCNetworkInterfaceRef, error) {
	if _sCNetworkInterfaceGetInterface == nil {
		return *new(SCNetworkInterfaceRef), symbolCallError("SCNetworkInterfaceGetInterface", "10.4", _sCNetworkInterfaceGetInterfaceErr)
	}
	return _sCNetworkInterfaceGetInterface(interface_), nil
}

// SCNetworkInterfaceGetInterface returns the underlying interface, for layered network interfaces.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterface(_:)
func SCNetworkInterfaceGetInterface(interface_ SCNetworkInterfaceRef) SCNetworkInterfaceRef {
	result, callErr := trySCNetworkInterfaceGetInterface(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetInterfaceType func(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef
var _sCNetworkInterfaceGetInterfaceTypeErr error

func trySCNetworkInterfaceGetInterfaceType(interface_ SCNetworkInterfaceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkInterfaceGetInterfaceType == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkInterfaceGetInterfaceType", "10.4", _sCNetworkInterfaceGetInterfaceTypeErr)
	}
	return _sCNetworkInterfaceGetInterfaceType(interface_), nil
}

// SCNetworkInterfaceGetInterfaceType returns the network interface type of the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterfaceType(_:)
func SCNetworkInterfaceGetInterfaceType(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkInterfaceGetInterfaceType(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetLocalizedDisplayName func(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef
var _sCNetworkInterfaceGetLocalizedDisplayNameErr error

func trySCNetworkInterfaceGetLocalizedDisplayName(interface_ SCNetworkInterfaceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkInterfaceGetLocalizedDisplayName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkInterfaceGetLocalizedDisplayName", "10.4", _sCNetworkInterfaceGetLocalizedDisplayNameErr)
	}
	return _sCNetworkInterfaceGetLocalizedDisplayName(interface_), nil
}

// SCNetworkInterfaceGetLocalizedDisplayName returns the localized display name, such as “Ethernet” or “FireWire”, for the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetLocalizedDisplayName(_:)
func SCNetworkInterfaceGetLocalizedDisplayName(interface_ SCNetworkInterfaceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkInterfaceGetLocalizedDisplayName(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetSupportedInterfaceTypes func(interface_ SCNetworkInterfaceRef) corefoundation.CFArrayRef
var _sCNetworkInterfaceGetSupportedInterfaceTypesErr error

func trySCNetworkInterfaceGetSupportedInterfaceTypes(interface_ SCNetworkInterfaceRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkInterfaceGetSupportedInterfaceTypes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkInterfaceGetSupportedInterfaceTypes", "10.4", _sCNetworkInterfaceGetSupportedInterfaceTypesErr)
	}
	return _sCNetworkInterfaceGetSupportedInterfaceTypes(interface_), nil
}

// SCNetworkInterfaceGetSupportedInterfaceTypes identifies all of the network interface types, such as PPP, that can be layered on top of the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedInterfaceTypes(_:)
func SCNetworkInterfaceGetSupportedInterfaceTypes(interface_ SCNetworkInterfaceRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkInterfaceGetSupportedInterfaceTypes(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetSupportedProtocolTypes func(interface_ SCNetworkInterfaceRef) corefoundation.CFArrayRef
var _sCNetworkInterfaceGetSupportedProtocolTypesErr error

func trySCNetworkInterfaceGetSupportedProtocolTypes(interface_ SCNetworkInterfaceRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkInterfaceGetSupportedProtocolTypes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkInterfaceGetSupportedProtocolTypes", "10.4", _sCNetworkInterfaceGetSupportedProtocolTypesErr)
	}
	return _sCNetworkInterfaceGetSupportedProtocolTypes(interface_), nil
}

// SCNetworkInterfaceGetSupportedProtocolTypes identifies all of the network protocol types, such as IPv4 and IPv6, that can be layered on top of the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedProtocolTypes(_:)
func SCNetworkInterfaceGetSupportedProtocolTypes(interface_ SCNetworkInterfaceRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkInterfaceGetSupportedProtocolTypes(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceGetTypeID func() uint
var _sCNetworkInterfaceGetTypeIDErr error

func trySCNetworkInterfaceGetTypeID() (uint, error) {
	if _sCNetworkInterfaceGetTypeID == nil {
		return 0, symbolCallError("SCNetworkInterfaceGetTypeID", "10.4", _sCNetworkInterfaceGetTypeIDErr)
	}
	return _sCNetworkInterfaceGetTypeID(), nil
}

// SCNetworkInterfaceGetTypeID returns the type identifier of all [SCNetworkInterface] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetTypeID()
func SCNetworkInterfaceGetTypeID() uint {
	result, callErr := trySCNetworkInterfaceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceSetConfiguration func(interface_ SCNetworkInterfaceRef, config corefoundation.CFDictionaryRef) bool
var _sCNetworkInterfaceSetConfigurationErr error

func trySCNetworkInterfaceSetConfiguration(interface_ SCNetworkInterfaceRef, config corefoundation.CFDictionaryRef) (bool, error) {
	if _sCNetworkInterfaceSetConfiguration == nil {
		return false, symbolCallError("SCNetworkInterfaceSetConfiguration", "10.4", _sCNetworkInterfaceSetConfigurationErr)
	}
	return _sCNetworkInterfaceSetConfiguration(interface_, config), nil
}

// SCNetworkInterfaceSetConfiguration stores the configuration settings for the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetConfiguration(_:_:)
func SCNetworkInterfaceSetConfiguration(interface_ SCNetworkInterfaceRef, config corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCNetworkInterfaceSetConfiguration(interface_, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceSetExtendedConfiguration func(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef, config corefoundation.CFDictionaryRef) bool
var _sCNetworkInterfaceSetExtendedConfigurationErr error

func trySCNetworkInterfaceSetExtendedConfiguration(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef, config corefoundation.CFDictionaryRef) (bool, error) {
	if _sCNetworkInterfaceSetExtendedConfiguration == nil {
		return false, symbolCallError("SCNetworkInterfaceSetExtendedConfiguration", "10.5", _sCNetworkInterfaceSetExtendedConfigurationErr)
	}
	return _sCNetworkInterfaceSetExtendedConfiguration(interface_, extendedType, config), nil
}

// SCNetworkInterfaceSetExtendedConfiguration stores the extended configuration settings for the specified interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetExtendedConfiguration(_:_:_:)
func SCNetworkInterfaceSetExtendedConfiguration(interface_ SCNetworkInterfaceRef, extendedType corefoundation.CFStringRef, config corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCNetworkInterfaceSetExtendedConfiguration(interface_, extendedType, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceSetMTU func(interface_ SCNetworkInterfaceRef, mtu int32) bool
var _sCNetworkInterfaceSetMTUErr error

func trySCNetworkInterfaceSetMTU(interface_ SCNetworkInterfaceRef, mtu int32) (bool, error) {
	if _sCNetworkInterfaceSetMTU == nil {
		return false, symbolCallError("SCNetworkInterfaceSetMTU", "10.5", _sCNetworkInterfaceSetMTUErr)
	}
	return _sCNetworkInterfaceSetMTU(interface_, mtu), nil
}

// SCNetworkInterfaceSetMTU sets the requested MTU setting for the specified network interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMTU(_:_:)
func SCNetworkInterfaceSetMTU(interface_ SCNetworkInterfaceRef, mtu int32) bool {
	result, callErr := trySCNetworkInterfaceSetMTU(interface_, mtu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkInterfaceSetMediaOptions func(interface_ SCNetworkInterfaceRef, subtype corefoundation.CFStringRef, options corefoundation.CFArrayRef) bool
var _sCNetworkInterfaceSetMediaOptionsErr error

func trySCNetworkInterfaceSetMediaOptions(interface_ SCNetworkInterfaceRef, subtype corefoundation.CFStringRef, options corefoundation.CFArrayRef) (bool, error) {
	if _sCNetworkInterfaceSetMediaOptions == nil {
		return false, symbolCallError("SCNetworkInterfaceSetMediaOptions", "10.5", _sCNetworkInterfaceSetMediaOptionsErr)
	}
	return _sCNetworkInterfaceSetMediaOptions(interface_, subtype, options), nil
}

// SCNetworkInterfaceSetMediaOptions sets the requested media subtype and options for the specified network interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMediaOptions(_:_:_:)
func SCNetworkInterfaceSetMediaOptions(interface_ SCNetworkInterfaceRef, subtype corefoundation.CFStringRef, options corefoundation.CFArrayRef) bool {
	result, callErr := trySCNetworkInterfaceSetMediaOptions(interface_, subtype, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolGetConfiguration func(protocol_ SCNetworkProtocolRef) corefoundation.CFDictionaryRef
var _sCNetworkProtocolGetConfigurationErr error

func trySCNetworkProtocolGetConfiguration(protocol_ SCNetworkProtocolRef) (corefoundation.CFDictionaryRef, error) {
	if _sCNetworkProtocolGetConfiguration == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCNetworkProtocolGetConfiguration", "10.4", _sCNetworkProtocolGetConfigurationErr)
	}
	return _sCNetworkProtocolGetConfiguration(protocol_), nil
}

// SCNetworkProtocolGetConfiguration returns the configuration settings associated with the specified protocol.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetConfiguration(_:)
func SCNetworkProtocolGetConfiguration(protocol_ SCNetworkProtocolRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCNetworkProtocolGetConfiguration(protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolGetEnabled func(protocol_ SCNetworkProtocolRef) bool
var _sCNetworkProtocolGetEnabledErr error

func trySCNetworkProtocolGetEnabled(protocol_ SCNetworkProtocolRef) (bool, error) {
	if _sCNetworkProtocolGetEnabled == nil {
		return false, symbolCallError("SCNetworkProtocolGetEnabled", "10.4", _sCNetworkProtocolGetEnabledErr)
	}
	return _sCNetworkProtocolGetEnabled(protocol_), nil
}

// SCNetworkProtocolGetEnabled returns a Boolean value indicating whether the specified protocol is enabled.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetEnabled(_:)
func SCNetworkProtocolGetEnabled(protocol_ SCNetworkProtocolRef) bool {
	result, callErr := trySCNetworkProtocolGetEnabled(protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolGetProtocolType func(protocol_ SCNetworkProtocolRef) corefoundation.CFStringRef
var _sCNetworkProtocolGetProtocolTypeErr error

func trySCNetworkProtocolGetProtocolType(protocol_ SCNetworkProtocolRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkProtocolGetProtocolType == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkProtocolGetProtocolType", "10.4", _sCNetworkProtocolGetProtocolTypeErr)
	}
	return _sCNetworkProtocolGetProtocolType(protocol_), nil
}

// SCNetworkProtocolGetProtocolType returns the type of the specified network protocol.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetProtocolType(_:)
func SCNetworkProtocolGetProtocolType(protocol_ SCNetworkProtocolRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkProtocolGetProtocolType(protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolGetTypeID func() uint
var _sCNetworkProtocolGetTypeIDErr error

func trySCNetworkProtocolGetTypeID() (uint, error) {
	if _sCNetworkProtocolGetTypeID == nil {
		return 0, symbolCallError("SCNetworkProtocolGetTypeID", "10.4", _sCNetworkProtocolGetTypeIDErr)
	}
	return _sCNetworkProtocolGetTypeID(), nil
}

// SCNetworkProtocolGetTypeID returns the type identifier of all [SCNetworkProtocol] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetTypeID()
func SCNetworkProtocolGetTypeID() uint {
	result, callErr := trySCNetworkProtocolGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolSetConfiguration func(protocol_ SCNetworkProtocolRef, config corefoundation.CFDictionaryRef) bool
var _sCNetworkProtocolSetConfigurationErr error

func trySCNetworkProtocolSetConfiguration(protocol_ SCNetworkProtocolRef, config corefoundation.CFDictionaryRef) (bool, error) {
	if _sCNetworkProtocolSetConfiguration == nil {
		return false, symbolCallError("SCNetworkProtocolSetConfiguration", "10.4", _sCNetworkProtocolSetConfigurationErr)
	}
	return _sCNetworkProtocolSetConfiguration(protocol_, config), nil
}

// SCNetworkProtocolSetConfiguration stores the configuration settings for the specified network protocol.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetConfiguration(_:_:)
func SCNetworkProtocolSetConfiguration(protocol_ SCNetworkProtocolRef, config corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCNetworkProtocolSetConfiguration(protocol_, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkProtocolSetEnabled func(protocol_ SCNetworkProtocolRef, enabled bool) bool
var _sCNetworkProtocolSetEnabledErr error

func trySCNetworkProtocolSetEnabled(protocol_ SCNetworkProtocolRef, enabled bool) (bool, error) {
	if _sCNetworkProtocolSetEnabled == nil {
		return false, symbolCallError("SCNetworkProtocolSetEnabled", "10.4", _sCNetworkProtocolSetEnabledErr)
	}
	return _sCNetworkProtocolSetEnabled(protocol_, enabled), nil
}

// SCNetworkProtocolSetEnabled enables or disables the specified protocol.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetEnabled(_:_:)
func SCNetworkProtocolSetEnabled(protocol_ SCNetworkProtocolRef, enabled bool) bool {
	result, callErr := trySCNetworkProtocolSetEnabled(protocol_, enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityCreateWithAddress func(allocator corefoundation.CFAllocatorRef, address unsafe.Pointer) SCNetworkReachabilityRef
var _sCNetworkReachabilityCreateWithAddressErr error

func trySCNetworkReachabilityCreateWithAddress(allocator corefoundation.CFAllocatorRef, address unsafe.Pointer) (SCNetworkReachabilityRef, error) {
	if _sCNetworkReachabilityCreateWithAddress == nil {
		return *new(SCNetworkReachabilityRef), symbolCallError("SCNetworkReachabilityCreateWithAddress", "10.3", _sCNetworkReachabilityCreateWithAddressErr)
	}
	return _sCNetworkReachabilityCreateWithAddress(allocator, address), nil
}

// SCNetworkReachabilityCreateWithAddress creates a reachability reference to the specified network address.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddress(_:_:)
func SCNetworkReachabilityCreateWithAddress(allocator corefoundation.CFAllocatorRef, address unsafe.Pointer) SCNetworkReachabilityRef {
	result, callErr := trySCNetworkReachabilityCreateWithAddress(allocator, address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityCreateWithAddressPair func(allocator corefoundation.CFAllocatorRef, localAddress unsafe.Pointer, remoteAddress unsafe.Pointer) SCNetworkReachabilityRef
var _sCNetworkReachabilityCreateWithAddressPairErr error

func trySCNetworkReachabilityCreateWithAddressPair(allocator corefoundation.CFAllocatorRef, localAddress unsafe.Pointer, remoteAddress unsafe.Pointer) (SCNetworkReachabilityRef, error) {
	if _sCNetworkReachabilityCreateWithAddressPair == nil {
		return *new(SCNetworkReachabilityRef), symbolCallError("SCNetworkReachabilityCreateWithAddressPair", "10.3", _sCNetworkReachabilityCreateWithAddressPairErr)
	}
	return _sCNetworkReachabilityCreateWithAddressPair(allocator, localAddress, remoteAddress), nil
}

// SCNetworkReachabilityCreateWithAddressPair creates a reachability reference to the specified network address.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddressPair(_:_:_:)
func SCNetworkReachabilityCreateWithAddressPair(allocator corefoundation.CFAllocatorRef, localAddress unsafe.Pointer, remoteAddress unsafe.Pointer) SCNetworkReachabilityRef {
	result, callErr := trySCNetworkReachabilityCreateWithAddressPair(allocator, localAddress, remoteAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityCreateWithName func(allocator corefoundation.CFAllocatorRef, nodename string) SCNetworkReachabilityRef
var _sCNetworkReachabilityCreateWithNameErr error

func trySCNetworkReachabilityCreateWithName(allocator corefoundation.CFAllocatorRef, nodename string) (SCNetworkReachabilityRef, error) {
	if _sCNetworkReachabilityCreateWithName == nil {
		return *new(SCNetworkReachabilityRef), symbolCallError("SCNetworkReachabilityCreateWithName", "10.3", _sCNetworkReachabilityCreateWithNameErr)
	}
	return _sCNetworkReachabilityCreateWithName(allocator, nodename), nil
}

// SCNetworkReachabilityCreateWithName creates a reachability reference to the specified network host or node name.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithName(_:_:)
func SCNetworkReachabilityCreateWithName(allocator corefoundation.CFAllocatorRef, nodename string) SCNetworkReachabilityRef {
	result, callErr := trySCNetworkReachabilityCreateWithName(allocator, nodename)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityGetFlags func(target SCNetworkReachabilityRef, flags *SCNetworkReachabilityFlags) bool
var _sCNetworkReachabilityGetFlagsErr error

func trySCNetworkReachabilityGetFlags(target SCNetworkReachabilityRef, flags *SCNetworkReachabilityFlags) (bool, error) {
	if _sCNetworkReachabilityGetFlags == nil {
		return false, symbolCallError("SCNetworkReachabilityGetFlags", "10.3", _sCNetworkReachabilityGetFlagsErr)
	}
	return _sCNetworkReachabilityGetFlags(target, flags), nil
}

// SCNetworkReachabilityGetFlags determines if the specified network target is reachable using the current network configuration.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetFlags(_:_:)
func SCNetworkReachabilityGetFlags(target SCNetworkReachabilityRef, flags *SCNetworkReachabilityFlags) bool {
	result, callErr := trySCNetworkReachabilityGetFlags(target, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityGetTypeID func() uint
var _sCNetworkReachabilityGetTypeIDErr error

func trySCNetworkReachabilityGetTypeID() (uint, error) {
	if _sCNetworkReachabilityGetTypeID == nil {
		return 0, symbolCallError("SCNetworkReachabilityGetTypeID", "10.3", _sCNetworkReachabilityGetTypeIDErr)
	}
	return _sCNetworkReachabilityGetTypeID(), nil
}

// SCNetworkReachabilityGetTypeID returns the type identifier of all [SCNetworkReachability] instances.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetTypeID()
func SCNetworkReachabilityGetTypeID() uint {
	result, callErr := trySCNetworkReachabilityGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityScheduleWithRunLoop func(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCNetworkReachabilityScheduleWithRunLoopErr error

func trySCNetworkReachabilityScheduleWithRunLoop(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkReachabilityScheduleWithRunLoop == nil {
		return false, symbolCallError("SCNetworkReachabilityScheduleWithRunLoop", "10.3", _sCNetworkReachabilityScheduleWithRunLoopErr)
	}
	return _sCNetworkReachabilityScheduleWithRunLoop(target, runLoop, runLoopMode), nil
}

// SCNetworkReachabilityScheduleWithRunLoop schedules the specified network target with the specified run loop and mode.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityScheduleWithRunLoop(_:_:_:)
func SCNetworkReachabilityScheduleWithRunLoop(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkReachabilityScheduleWithRunLoop(target, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilitySetCallback func(target SCNetworkReachabilityRef, callout SCNetworkReachabilityCallBack, context *SCNetworkReachabilityContext) bool
var _sCNetworkReachabilitySetCallbackErr error

func trySCNetworkReachabilitySetCallback(target SCNetworkReachabilityRef, callout SCNetworkReachabilityCallBack, context *SCNetworkReachabilityContext) (bool, error) {
	if _sCNetworkReachabilitySetCallback == nil {
		return false, symbolCallError("SCNetworkReachabilitySetCallback", "10.3", _sCNetworkReachabilitySetCallbackErr)
	}
	return _sCNetworkReachabilitySetCallback(target, callout, context), nil
}

// SCNetworkReachabilitySetCallback assigns a client to the specified target, which receives callbacks when the reachability of the target changes.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetCallback(_:_:_:)
func SCNetworkReachabilitySetCallback(target SCNetworkReachabilityRef, callout SCNetworkReachabilityCallBack, context *SCNetworkReachabilityContext) bool {
	result, callErr := trySCNetworkReachabilitySetCallback(target, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilitySetDispatchQueue func(target SCNetworkReachabilityRef, queue uintptr) bool
var _sCNetworkReachabilitySetDispatchQueueErr error

func trySCNetworkReachabilitySetDispatchQueue(target SCNetworkReachabilityRef, queue dispatch.Queue) (bool, error) {
	if _sCNetworkReachabilitySetDispatchQueue == nil {
		return false, symbolCallError("SCNetworkReachabilitySetDispatchQueue", "10.6", _sCNetworkReachabilitySetDispatchQueueErr)
	}
	return _sCNetworkReachabilitySetDispatchQueue(target, uintptr(queue.Handle())), nil
}

// SCNetworkReachabilitySetDispatchQueue schedules callbacks for the specified target on the specified dispatch queue.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetDispatchQueue(_:_:)
func SCNetworkReachabilitySetDispatchQueue(target SCNetworkReachabilityRef, queue dispatch.Queue) bool {
	result, callErr := trySCNetworkReachabilitySetDispatchQueue(target, queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkReachabilityUnscheduleFromRunLoop func(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCNetworkReachabilityUnscheduleFromRunLoopErr error

func trySCNetworkReachabilityUnscheduleFromRunLoop(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkReachabilityUnscheduleFromRunLoop == nil {
		return false, symbolCallError("SCNetworkReachabilityUnscheduleFromRunLoop", "10.3", _sCNetworkReachabilityUnscheduleFromRunLoopErr)
	}
	return _sCNetworkReachabilityUnscheduleFromRunLoop(target, runLoop, runLoopMode), nil
}

// SCNetworkReachabilityUnscheduleFromRunLoop unschedules the specified target from the specified run loop and mode.
//
// Deprecated: Deprecated since macOS 14.4.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityUnscheduleFromRunLoop(_:_:_:)
func SCNetworkReachabilityUnscheduleFromRunLoop(target SCNetworkReachabilityRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkReachabilityUnscheduleFromRunLoop(target, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceAddProtocolType func(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) bool
var _sCNetworkServiceAddProtocolTypeErr error

func trySCNetworkServiceAddProtocolType(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkServiceAddProtocolType == nil {
		return false, symbolCallError("SCNetworkServiceAddProtocolType", "10.4", _sCNetworkServiceAddProtocolTypeErr)
	}
	return _sCNetworkServiceAddProtocolType(service, protocolType), nil
}

// SCNetworkServiceAddProtocolType adds the network protocol of the specified type to the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceAddProtocolType(_:_:)
func SCNetworkServiceAddProtocolType(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkServiceAddProtocolType(service, protocolType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceCopy func(prefs SCPreferencesRef, serviceID corefoundation.CFStringRef) SCNetworkServiceRef
var _sCNetworkServiceCopyErr error

func trySCNetworkServiceCopy(prefs SCPreferencesRef, serviceID corefoundation.CFStringRef) (SCNetworkServiceRef, error) {
	if _sCNetworkServiceCopy == nil {
		return *new(SCNetworkServiceRef), symbolCallError("SCNetworkServiceCopy", "10.4", _sCNetworkServiceCopyErr)
	}
	return _sCNetworkServiceCopy(prefs, serviceID), nil
}

// SCNetworkServiceCopy returns the network service with the specified identifier.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopy(_:_:)
func SCNetworkServiceCopy(prefs SCPreferencesRef, serviceID corefoundation.CFStringRef) SCNetworkServiceRef {
	result, callErr := trySCNetworkServiceCopy(prefs, serviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceCopyAll func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCNetworkServiceCopyAllErr error

func trySCNetworkServiceCopyAll(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkServiceCopyAll == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkServiceCopyAll", "10.4", _sCNetworkServiceCopyAllErr)
	}
	return _sCNetworkServiceCopyAll(prefs), nil
}

// SCNetworkServiceCopyAll returns all available network services for the specified preferences.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyAll(_:)
func SCNetworkServiceCopyAll(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkServiceCopyAll(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceCopyProtocol func(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) SCNetworkProtocolRef
var _sCNetworkServiceCopyProtocolErr error

func trySCNetworkServiceCopyProtocol(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) (SCNetworkProtocolRef, error) {
	if _sCNetworkServiceCopyProtocol == nil {
		return *new(SCNetworkProtocolRef), symbolCallError("SCNetworkServiceCopyProtocol", "10.4", _sCNetworkServiceCopyProtocolErr)
	}
	return _sCNetworkServiceCopyProtocol(service, protocolType), nil
}

// SCNetworkServiceCopyProtocol returns the network protocol of the specified type for the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocol(_:_:)
func SCNetworkServiceCopyProtocol(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) SCNetworkProtocolRef {
	result, callErr := trySCNetworkServiceCopyProtocol(service, protocolType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceCopyProtocols func(service SCNetworkServiceRef) corefoundation.CFArrayRef
var _sCNetworkServiceCopyProtocolsErr error

func trySCNetworkServiceCopyProtocols(service SCNetworkServiceRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkServiceCopyProtocols == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkServiceCopyProtocols", "10.4", _sCNetworkServiceCopyProtocolsErr)
	}
	return _sCNetworkServiceCopyProtocols(service), nil
}

// SCNetworkServiceCopyProtocols returns all network protocols associated with the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocols(_:)
func SCNetworkServiceCopyProtocols(service SCNetworkServiceRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkServiceCopyProtocols(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceCreate func(prefs SCPreferencesRef, interface_ SCNetworkInterfaceRef) SCNetworkServiceRef
var _sCNetworkServiceCreateErr error

func trySCNetworkServiceCreate(prefs SCPreferencesRef, interface_ SCNetworkInterfaceRef) (SCNetworkServiceRef, error) {
	if _sCNetworkServiceCreate == nil {
		return *new(SCNetworkServiceRef), symbolCallError("SCNetworkServiceCreate", "10.4", _sCNetworkServiceCreateErr)
	}
	return _sCNetworkServiceCreate(prefs, interface_), nil
}

// SCNetworkServiceCreate creates a new network service for the specified interface in the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCreate(_:_:)
func SCNetworkServiceCreate(prefs SCPreferencesRef, interface_ SCNetworkInterfaceRef) SCNetworkServiceRef {
	result, callErr := trySCNetworkServiceCreate(prefs, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceEstablishDefaultConfiguration func(service SCNetworkServiceRef) bool
var _sCNetworkServiceEstablishDefaultConfigurationErr error

func trySCNetworkServiceEstablishDefaultConfiguration(service SCNetworkServiceRef) (bool, error) {
	if _sCNetworkServiceEstablishDefaultConfiguration == nil {
		return false, symbolCallError("SCNetworkServiceEstablishDefaultConfiguration", "10.5", _sCNetworkServiceEstablishDefaultConfigurationErr)
	}
	return _sCNetworkServiceEstablishDefaultConfiguration(service), nil
}

// SCNetworkServiceEstablishDefaultConfiguration establishes the default configuration for the specified network service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceEstablishDefaultConfiguration(_:)
func SCNetworkServiceEstablishDefaultConfiguration(service SCNetworkServiceRef) bool {
	result, callErr := trySCNetworkServiceEstablishDefaultConfiguration(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceGetEnabled func(service SCNetworkServiceRef) bool
var _sCNetworkServiceGetEnabledErr error

func trySCNetworkServiceGetEnabled(service SCNetworkServiceRef) (bool, error) {
	if _sCNetworkServiceGetEnabled == nil {
		return false, symbolCallError("SCNetworkServiceGetEnabled", "10.4", _sCNetworkServiceGetEnabledErr)
	}
	return _sCNetworkServiceGetEnabled(service), nil
}

// SCNetworkServiceGetEnabled returns a Boolean value indicating whether the specified service is enabled.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetEnabled(_:)
func SCNetworkServiceGetEnabled(service SCNetworkServiceRef) bool {
	result, callErr := trySCNetworkServiceGetEnabled(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceGetInterface func(service SCNetworkServiceRef) SCNetworkInterfaceRef
var _sCNetworkServiceGetInterfaceErr error

func trySCNetworkServiceGetInterface(service SCNetworkServiceRef) (SCNetworkInterfaceRef, error) {
	if _sCNetworkServiceGetInterface == nil {
		return *new(SCNetworkInterfaceRef), symbolCallError("SCNetworkServiceGetInterface", "10.4", _sCNetworkServiceGetInterfaceErr)
	}
	return _sCNetworkServiceGetInterface(service), nil
}

// SCNetworkServiceGetInterface returns the network interface associated with the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetInterface(_:)
func SCNetworkServiceGetInterface(service SCNetworkServiceRef) SCNetworkInterfaceRef {
	result, callErr := trySCNetworkServiceGetInterface(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceGetName func(service SCNetworkServiceRef) corefoundation.CFStringRef
var _sCNetworkServiceGetNameErr error

func trySCNetworkServiceGetName(service SCNetworkServiceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkServiceGetName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkServiceGetName", "10.4", _sCNetworkServiceGetNameErr)
	}
	return _sCNetworkServiceGetName(service), nil
}

// SCNetworkServiceGetName returns the user-specified name associated with the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetName(_:)
func SCNetworkServiceGetName(service SCNetworkServiceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkServiceGetName(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceGetServiceID func(service SCNetworkServiceRef) corefoundation.CFStringRef
var _sCNetworkServiceGetServiceIDErr error

func trySCNetworkServiceGetServiceID(service SCNetworkServiceRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkServiceGetServiceID == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkServiceGetServiceID", "10.4", _sCNetworkServiceGetServiceIDErr)
	}
	return _sCNetworkServiceGetServiceID(service), nil
}

// SCNetworkServiceGetServiceID returns the identifier for the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetServiceID(_:)
func SCNetworkServiceGetServiceID(service SCNetworkServiceRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkServiceGetServiceID(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceGetTypeID func() uint
var _sCNetworkServiceGetTypeIDErr error

func trySCNetworkServiceGetTypeID() (uint, error) {
	if _sCNetworkServiceGetTypeID == nil {
		return 0, symbolCallError("SCNetworkServiceGetTypeID", "10.4", _sCNetworkServiceGetTypeIDErr)
	}
	return _sCNetworkServiceGetTypeID(), nil
}

// SCNetworkServiceGetTypeID returns the type identifier of all [SCNetworkService] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetTypeID()
func SCNetworkServiceGetTypeID() uint {
	result, callErr := trySCNetworkServiceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceRemove func(service SCNetworkServiceRef) bool
var _sCNetworkServiceRemoveErr error

func trySCNetworkServiceRemove(service SCNetworkServiceRef) (bool, error) {
	if _sCNetworkServiceRemove == nil {
		return false, symbolCallError("SCNetworkServiceRemove", "10.4", _sCNetworkServiceRemoveErr)
	}
	return _sCNetworkServiceRemove(service), nil
}

// SCNetworkServiceRemove removes the specified network service from the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemove(_:)
func SCNetworkServiceRemove(service SCNetworkServiceRef) bool {
	result, callErr := trySCNetworkServiceRemove(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceRemoveProtocolType func(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) bool
var _sCNetworkServiceRemoveProtocolTypeErr error

func trySCNetworkServiceRemoveProtocolType(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkServiceRemoveProtocolType == nil {
		return false, symbolCallError("SCNetworkServiceRemoveProtocolType", "10.4", _sCNetworkServiceRemoveProtocolTypeErr)
	}
	return _sCNetworkServiceRemoveProtocolType(service, protocolType), nil
}

// SCNetworkServiceRemoveProtocolType removes the network protocol of the specified type from the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemoveProtocolType(_:_:)
func SCNetworkServiceRemoveProtocolType(service SCNetworkServiceRef, protocolType corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkServiceRemoveProtocolType(service, protocolType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceSetEnabled func(service SCNetworkServiceRef, enabled bool) bool
var _sCNetworkServiceSetEnabledErr error

func trySCNetworkServiceSetEnabled(service SCNetworkServiceRef, enabled bool) (bool, error) {
	if _sCNetworkServiceSetEnabled == nil {
		return false, symbolCallError("SCNetworkServiceSetEnabled", "10.4", _sCNetworkServiceSetEnabledErr)
	}
	return _sCNetworkServiceSetEnabled(service, enabled), nil
}

// SCNetworkServiceSetEnabled enables or disables the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetEnabled(_:_:)
func SCNetworkServiceSetEnabled(service SCNetworkServiceRef, enabled bool) bool {
	result, callErr := trySCNetworkServiceSetEnabled(service, enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkServiceSetName func(service SCNetworkServiceRef, name corefoundation.CFStringRef) bool
var _sCNetworkServiceSetNameErr error

func trySCNetworkServiceSetName(service SCNetworkServiceRef, name corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkServiceSetName == nil {
		return false, symbolCallError("SCNetworkServiceSetName", "10.4", _sCNetworkServiceSetNameErr)
	}
	return _sCNetworkServiceSetName(service, name), nil
}

// SCNetworkServiceSetName stores the user-specified name for the specified service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetName(_:_:)
func SCNetworkServiceSetName(service SCNetworkServiceRef, name corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkServiceSetName(service, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetAddService func(set SCNetworkSetRef, service SCNetworkServiceRef) bool
var _sCNetworkSetAddServiceErr error

func trySCNetworkSetAddService(set SCNetworkSetRef, service SCNetworkServiceRef) (bool, error) {
	if _sCNetworkSetAddService == nil {
		return false, symbolCallError("SCNetworkSetAddService", "10.4", _sCNetworkSetAddServiceErr)
	}
	return _sCNetworkSetAddService(set, service), nil
}

// SCNetworkSetAddService adds the specified network service to the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetAddService(_:_:)
func SCNetworkSetAddService(set SCNetworkSetRef, service SCNetworkServiceRef) bool {
	result, callErr := trySCNetworkSetAddService(set, service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetContainsInterface func(set SCNetworkSetRef, interface_ SCNetworkInterfaceRef) bool
var _sCNetworkSetContainsInterfaceErr error

func trySCNetworkSetContainsInterface(set SCNetworkSetRef, interface_ SCNetworkInterfaceRef) (bool, error) {
	if _sCNetworkSetContainsInterface == nil {
		return false, symbolCallError("SCNetworkSetContainsInterface", "10.5", _sCNetworkSetContainsInterfaceErr)
	}
	return _sCNetworkSetContainsInterface(set, interface_), nil
}

// SCNetworkSetContainsInterface returns a Boolean value indicating whether the specified interface is represented by at least one network service in the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetContainsInterface(_:_:)
func SCNetworkSetContainsInterface(set SCNetworkSetRef, interface_ SCNetworkInterfaceRef) bool {
	result, callErr := trySCNetworkSetContainsInterface(set, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetCopy func(prefs SCPreferencesRef, setID corefoundation.CFStringRef) SCNetworkSetRef
var _sCNetworkSetCopyErr error

func trySCNetworkSetCopy(prefs SCPreferencesRef, setID corefoundation.CFStringRef) (SCNetworkSetRef, error) {
	if _sCNetworkSetCopy == nil {
		return *new(SCNetworkSetRef), symbolCallError("SCNetworkSetCopy", "10.4", _sCNetworkSetCopyErr)
	}
	return _sCNetworkSetCopy(prefs, setID), nil
}

// SCNetworkSetCopy returns the set with the specified identifier.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopy(_:_:)
func SCNetworkSetCopy(prefs SCPreferencesRef, setID corefoundation.CFStringRef) SCNetworkSetRef {
	result, callErr := trySCNetworkSetCopy(prefs, setID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetCopyAll func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCNetworkSetCopyAllErr error

func trySCNetworkSetCopyAll(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkSetCopyAll == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkSetCopyAll", "10.4", _sCNetworkSetCopyAllErr)
	}
	return _sCNetworkSetCopyAll(prefs), nil
}

// SCNetworkSetCopyAll returns all available sets for the specified preferences session.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyAll(_:)
func SCNetworkSetCopyAll(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkSetCopyAll(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetCopyCurrent func(prefs SCPreferencesRef) SCNetworkSetRef
var _sCNetworkSetCopyCurrentErr error

func trySCNetworkSetCopyCurrent(prefs SCPreferencesRef) (SCNetworkSetRef, error) {
	if _sCNetworkSetCopyCurrent == nil {
		return *new(SCNetworkSetRef), symbolCallError("SCNetworkSetCopyCurrent", "10.4", _sCNetworkSetCopyCurrentErr)
	}
	return _sCNetworkSetCopyCurrent(prefs), nil
}

// SCNetworkSetCopyCurrent returns the current set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyCurrent(_:)
func SCNetworkSetCopyCurrent(prefs SCPreferencesRef) SCNetworkSetRef {
	result, callErr := trySCNetworkSetCopyCurrent(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetCopyServices func(set SCNetworkSetRef) corefoundation.CFArrayRef
var _sCNetworkSetCopyServicesErr error

func trySCNetworkSetCopyServices(set SCNetworkSetRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkSetCopyServices == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkSetCopyServices", "10.4", _sCNetworkSetCopyServicesErr)
	}
	return _sCNetworkSetCopyServices(set), nil
}

// SCNetworkSetCopyServices returns all network services associated with the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyServices(_:)
func SCNetworkSetCopyServices(set SCNetworkSetRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkSetCopyServices(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetCreate func(prefs SCPreferencesRef) SCNetworkSetRef
var _sCNetworkSetCreateErr error

func trySCNetworkSetCreate(prefs SCPreferencesRef) (SCNetworkSetRef, error) {
	if _sCNetworkSetCreate == nil {
		return *new(SCNetworkSetRef), symbolCallError("SCNetworkSetCreate", "10.4", _sCNetworkSetCreateErr)
	}
	return _sCNetworkSetCreate(prefs), nil
}

// SCNetworkSetCreate creates a new set in the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCreate(_:)
func SCNetworkSetCreate(prefs SCPreferencesRef) SCNetworkSetRef {
	result, callErr := trySCNetworkSetCreate(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetGetName func(set SCNetworkSetRef) corefoundation.CFStringRef
var _sCNetworkSetGetNameErr error

func trySCNetworkSetGetName(set SCNetworkSetRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkSetGetName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkSetGetName", "10.4", _sCNetworkSetGetNameErr)
	}
	return _sCNetworkSetGetName(set), nil
}

// SCNetworkSetGetName returns the user-specified name associated with the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetName(_:)
func SCNetworkSetGetName(set SCNetworkSetRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkSetGetName(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetGetServiceOrder func(set SCNetworkSetRef) corefoundation.CFArrayRef
var _sCNetworkSetGetServiceOrderErr error

func trySCNetworkSetGetServiceOrder(set SCNetworkSetRef) (corefoundation.CFArrayRef, error) {
	if _sCNetworkSetGetServiceOrder == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCNetworkSetGetServiceOrder", "10.4", _sCNetworkSetGetServiceOrderErr)
	}
	return _sCNetworkSetGetServiceOrder(set), nil
}

// SCNetworkSetGetServiceOrder returns the user-specified ordering of network services within the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetServiceOrder(_:)
func SCNetworkSetGetServiceOrder(set SCNetworkSetRef) corefoundation.CFArrayRef {
	result, callErr := trySCNetworkSetGetServiceOrder(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetGetSetID func(set SCNetworkSetRef) corefoundation.CFStringRef
var _sCNetworkSetGetSetIDErr error

func trySCNetworkSetGetSetID(set SCNetworkSetRef) (corefoundation.CFStringRef, error) {
	if _sCNetworkSetGetSetID == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCNetworkSetGetSetID", "10.4", _sCNetworkSetGetSetIDErr)
	}
	return _sCNetworkSetGetSetID(set), nil
}

// SCNetworkSetGetSetID returns the identifier for the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetSetID(_:)
func SCNetworkSetGetSetID(set SCNetworkSetRef) corefoundation.CFStringRef {
	result, callErr := trySCNetworkSetGetSetID(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetGetTypeID func() uint
var _sCNetworkSetGetTypeIDErr error

func trySCNetworkSetGetTypeID() (uint, error) {
	if _sCNetworkSetGetTypeID == nil {
		return 0, symbolCallError("SCNetworkSetGetTypeID", "10.4", _sCNetworkSetGetTypeIDErr)
	}
	return _sCNetworkSetGetTypeID(), nil
}

// SCNetworkSetGetTypeID returns the type identifier of all [SCNetworkSet] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetTypeID()
func SCNetworkSetGetTypeID() uint {
	result, callErr := trySCNetworkSetGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetRemove func(set SCNetworkSetRef) bool
var _sCNetworkSetRemoveErr error

func trySCNetworkSetRemove(set SCNetworkSetRef) (bool, error) {
	if _sCNetworkSetRemove == nil {
		return false, symbolCallError("SCNetworkSetRemove", "10.4", _sCNetworkSetRemoveErr)
	}
	return _sCNetworkSetRemove(set), nil
}

// SCNetworkSetRemove removes the specified set from the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemove(_:)
func SCNetworkSetRemove(set SCNetworkSetRef) bool {
	result, callErr := trySCNetworkSetRemove(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetRemoveService func(set SCNetworkSetRef, service SCNetworkServiceRef) bool
var _sCNetworkSetRemoveServiceErr error

func trySCNetworkSetRemoveService(set SCNetworkSetRef, service SCNetworkServiceRef) (bool, error) {
	if _sCNetworkSetRemoveService == nil {
		return false, symbolCallError("SCNetworkSetRemoveService", "10.4", _sCNetworkSetRemoveServiceErr)
	}
	return _sCNetworkSetRemoveService(set, service), nil
}

// SCNetworkSetRemoveService removes the specified network service from the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemoveService(_:_:)
func SCNetworkSetRemoveService(set SCNetworkSetRef, service SCNetworkServiceRef) bool {
	result, callErr := trySCNetworkSetRemoveService(set, service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetSetCurrent func(set SCNetworkSetRef) bool
var _sCNetworkSetSetCurrentErr error

func trySCNetworkSetSetCurrent(set SCNetworkSetRef) (bool, error) {
	if _sCNetworkSetSetCurrent == nil {
		return false, symbolCallError("SCNetworkSetSetCurrent", "10.4", _sCNetworkSetSetCurrentErr)
	}
	return _sCNetworkSetSetCurrent(set), nil
}

// SCNetworkSetSetCurrent specifies the set that should be the current set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetCurrent(_:)
func SCNetworkSetSetCurrent(set SCNetworkSetRef) bool {
	result, callErr := trySCNetworkSetSetCurrent(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetSetName func(set SCNetworkSetRef, name corefoundation.CFStringRef) bool
var _sCNetworkSetSetNameErr error

func trySCNetworkSetSetName(set SCNetworkSetRef, name corefoundation.CFStringRef) (bool, error) {
	if _sCNetworkSetSetName == nil {
		return false, symbolCallError("SCNetworkSetSetName", "10.4", _sCNetworkSetSetNameErr)
	}
	return _sCNetworkSetSetName(set, name), nil
}

// SCNetworkSetSetName stores the user-specified name for the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetName(_:_:)
func SCNetworkSetSetName(set SCNetworkSetRef, name corefoundation.CFStringRef) bool {
	result, callErr := trySCNetworkSetSetName(set, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCNetworkSetSetServiceOrder func(set SCNetworkSetRef, newOrder corefoundation.CFArrayRef) bool
var _sCNetworkSetSetServiceOrderErr error

func trySCNetworkSetSetServiceOrder(set SCNetworkSetRef, newOrder corefoundation.CFArrayRef) (bool, error) {
	if _sCNetworkSetSetServiceOrder == nil {
		return false, symbolCallError("SCNetworkSetSetServiceOrder", "10.4", _sCNetworkSetSetServiceOrderErr)
	}
	return _sCNetworkSetSetServiceOrder(set, newOrder), nil
}

// SCNetworkSetSetServiceOrder stores the user-specified ordering of network services for the specified set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetServiceOrder(_:_:)
func SCNetworkSetSetServiceOrder(set SCNetworkSetRef, newOrder corefoundation.CFArrayRef) bool {
	result, callErr := trySCNetworkSetSetServiceOrder(set, newOrder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesAddValue func(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool
var _sCPreferencesAddValueErr error

func trySCPreferencesAddValue(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) (bool, error) {
	if _sCPreferencesAddValue == nil {
		return false, symbolCallError("SCPreferencesAddValue", "10.1", _sCPreferencesAddValueErr)
	}
	return _sCPreferencesAddValue(prefs, key, value), nil
}

// SCPreferencesAddValue associates the specified value with the specified preference key.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesAddValue(_:_:_:)
func SCPreferencesAddValue(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool {
	result, callErr := trySCPreferencesAddValue(prefs, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesApplyChanges func(prefs SCPreferencesRef) bool
var _sCPreferencesApplyChangesErr error

func trySCPreferencesApplyChanges(prefs SCPreferencesRef) (bool, error) {
	if _sCPreferencesApplyChanges == nil {
		return false, symbolCallError("SCPreferencesApplyChanges", "10.1", _sCPreferencesApplyChangesErr)
	}
	return _sCPreferencesApplyChanges(prefs), nil
}

// SCPreferencesApplyChanges requests that the currently stored configuration preferences be applied to the active configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesApplyChanges(_:)
func SCPreferencesApplyChanges(prefs SCPreferencesRef) bool {
	result, callErr := trySCPreferencesApplyChanges(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesCommitChanges func(prefs SCPreferencesRef) bool
var _sCPreferencesCommitChangesErr error

func trySCPreferencesCommitChanges(prefs SCPreferencesRef) (bool, error) {
	if _sCPreferencesCommitChanges == nil {
		return false, symbolCallError("SCPreferencesCommitChanges", "10.1", _sCPreferencesCommitChangesErr)
	}
	return _sCPreferencesCommitChanges(prefs), nil
}

// SCPreferencesCommitChanges commits changes made to the configuration preferences to persistent storage.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCommitChanges(_:)
func SCPreferencesCommitChanges(prefs SCPreferencesRef) bool {
	result, callErr := trySCPreferencesCommitChanges(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesCopyKeyList func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCPreferencesCopyKeyListErr error

func trySCPreferencesCopyKeyList(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCPreferencesCopyKeyList == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCPreferencesCopyKeyList", "10.1", _sCPreferencesCopyKeyListErr)
	}
	return _sCPreferencesCopyKeyList(prefs), nil
}

// SCPreferencesCopyKeyList returns the currently defined preference keys.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCopyKeyList(_:)
func SCPreferencesCopyKeyList(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCPreferencesCopyKeyList(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesCreate func(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef) SCPreferencesRef
var _sCPreferencesCreateErr error

func trySCPreferencesCreate(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef) (SCPreferencesRef, error) {
	if _sCPreferencesCreate == nil {
		return *new(SCPreferencesRef), symbolCallError("SCPreferencesCreate", "10.1", _sCPreferencesCreateErr)
	}
	return _sCPreferencesCreate(allocator, name, prefsID), nil
}

// SCPreferencesCreate initiates access to the per-system set of configuration preferences.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreate(_:_:_:)
func SCPreferencesCreate(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef) SCPreferencesRef {
	result, callErr := trySCPreferencesCreate(allocator, name, prefsID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesCreateWithAuthorization func(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef, authorization security.AuthorizationRef) SCPreferencesRef
var _sCPreferencesCreateWithAuthorizationErr error

func trySCPreferencesCreateWithAuthorization(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef, authorization security.AuthorizationRef) (SCPreferencesRef, error) {
	if _sCPreferencesCreateWithAuthorization == nil {
		return *new(SCPreferencesRef), symbolCallError("SCPreferencesCreateWithAuthorization", "10.5", _sCPreferencesCreateWithAuthorizationErr)
	}
	return _sCPreferencesCreateWithAuthorization(allocator, name, prefsID, authorization), nil
}

// SCPreferencesCreateWithAuthorization initiates access to the per-system set of configuration preferences with the specified authorization.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreateWithAuthorization(_:_:_:_:)
func SCPreferencesCreateWithAuthorization(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, prefsID corefoundation.CFStringRef, authorization security.AuthorizationRef) SCPreferencesRef {
	result, callErr := trySCPreferencesCreateWithAuthorization(allocator, name, prefsID, authorization)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesGetSignature func(prefs SCPreferencesRef) corefoundation.CFDataRef
var _sCPreferencesGetSignatureErr error

func trySCPreferencesGetSignature(prefs SCPreferencesRef) (corefoundation.CFDataRef, error) {
	if _sCPreferencesGetSignature == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("SCPreferencesGetSignature", "10.1", _sCPreferencesGetSignatureErr)
	}
	return _sCPreferencesGetSignature(prefs), nil
}

// SCPreferencesGetSignature returns a value that can be used to determine if the saved configuration preferences have changed.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetSignature(_:)
func SCPreferencesGetSignature(prefs SCPreferencesRef) corefoundation.CFDataRef {
	result, callErr := trySCPreferencesGetSignature(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesGetTypeID func() uint
var _sCPreferencesGetTypeIDErr error

func trySCPreferencesGetTypeID() (uint, error) {
	if _sCPreferencesGetTypeID == nil {
		return 0, symbolCallError("SCPreferencesGetTypeID", "10.1", _sCPreferencesGetTypeIDErr)
	}
	return _sCPreferencesGetTypeID(), nil
}

// SCPreferencesGetTypeID returns the type identifier of all [SCPreferences] instances.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetTypeID()
func SCPreferencesGetTypeID() uint {
	result, callErr := trySCPreferencesGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesGetValue func(prefs SCPreferencesRef, key corefoundation.CFStringRef) corefoundation.CFPropertyListRef
var _sCPreferencesGetValueErr error

func trySCPreferencesGetValue(prefs SCPreferencesRef, key corefoundation.CFStringRef) (corefoundation.CFPropertyListRef, error) {
	if _sCPreferencesGetValue == nil {
		return *new(corefoundation.CFPropertyListRef), symbolCallError("SCPreferencesGetValue", "10.1", _sCPreferencesGetValueErr)
	}
	return _sCPreferencesGetValue(prefs, key), nil
}

// SCPreferencesGetValue retrieves the value associated with the specified preference key.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetValue(_:_:)
func SCPreferencesGetValue(prefs SCPreferencesRef, key corefoundation.CFStringRef) corefoundation.CFPropertyListRef {
	result, callErr := trySCPreferencesGetValue(prefs, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesLock func(prefs SCPreferencesRef, wait bool) bool
var _sCPreferencesLockErr error

func trySCPreferencesLock(prefs SCPreferencesRef, wait bool) (bool, error) {
	if _sCPreferencesLock == nil {
		return false, symbolCallError("SCPreferencesLock", "10.1", _sCPreferencesLockErr)
	}
	return _sCPreferencesLock(prefs, wait), nil
}

// SCPreferencesLock locks access to the configuration preferences.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesLock(_:_:)
func SCPreferencesLock(prefs SCPreferencesRef, wait bool) bool {
	result, callErr := trySCPreferencesLock(prefs, wait)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathCreateUniqueChild func(prefs SCPreferencesRef, prefix corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCPreferencesPathCreateUniqueChildErr error

func trySCPreferencesPathCreateUniqueChild(prefs SCPreferencesRef, prefix corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCPreferencesPathCreateUniqueChild == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCPreferencesPathCreateUniqueChild", "10.1", _sCPreferencesPathCreateUniqueChildErr)
	}
	return _sCPreferencesPathCreateUniqueChild(prefs, prefix), nil
}

// SCPreferencesPathCreateUniqueChild creates a new path component rooted at the specified path in the dictionary hierarchy.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathCreateUniqueChild(_:_:)
func SCPreferencesPathCreateUniqueChild(prefs SCPreferencesRef, prefix corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCPreferencesPathCreateUniqueChild(prefs, prefix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathGetLink func(prefs SCPreferencesRef, path corefoundation.CFStringRef) corefoundation.CFStringRef
var _sCPreferencesPathGetLinkErr error

func trySCPreferencesPathGetLink(prefs SCPreferencesRef, path corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _sCPreferencesPathGetLink == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("SCPreferencesPathGetLink", "10.1", _sCPreferencesPathGetLinkErr)
	}
	return _sCPreferencesPathGetLink(prefs, path), nil
}

// SCPreferencesPathGetLink returns the link associated with the specified path.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetLink(_:_:)
func SCPreferencesPathGetLink(prefs SCPreferencesRef, path corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := trySCPreferencesPathGetLink(prefs, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathGetValue func(prefs SCPreferencesRef, path corefoundation.CFStringRef) corefoundation.CFDictionaryRef
var _sCPreferencesPathGetValueErr error

func trySCPreferencesPathGetValue(prefs SCPreferencesRef, path corefoundation.CFStringRef) (corefoundation.CFDictionaryRef, error) {
	if _sCPreferencesPathGetValue == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCPreferencesPathGetValue", "10.1", _sCPreferencesPathGetValueErr)
	}
	return _sCPreferencesPathGetValue(prefs, path), nil
}

// SCPreferencesPathGetValue returns the dictionary associated with the specified path.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetValue(_:_:)
func SCPreferencesPathGetValue(prefs SCPreferencesRef, path corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCPreferencesPathGetValue(prefs, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathRemoveValue func(prefs SCPreferencesRef, path corefoundation.CFStringRef) bool
var _sCPreferencesPathRemoveValueErr error

func trySCPreferencesPathRemoveValue(prefs SCPreferencesRef, path corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesPathRemoveValue == nil {
		return false, symbolCallError("SCPreferencesPathRemoveValue", "10.1", _sCPreferencesPathRemoveValueErr)
	}
	return _sCPreferencesPathRemoveValue(prefs, path), nil
}

// SCPreferencesPathRemoveValue removes the data associated with the specified path.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathRemoveValue(_:_:)
func SCPreferencesPathRemoveValue(prefs SCPreferencesRef, path corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesPathRemoveValue(prefs, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathSetLink func(prefs SCPreferencesRef, path corefoundation.CFStringRef, link corefoundation.CFStringRef) bool
var _sCPreferencesPathSetLinkErr error

func trySCPreferencesPathSetLink(prefs SCPreferencesRef, path corefoundation.CFStringRef, link corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesPathSetLink == nil {
		return false, symbolCallError("SCPreferencesPathSetLink", "10.1", _sCPreferencesPathSetLinkErr)
	}
	return _sCPreferencesPathSetLink(prefs, path, link), nil
}

// SCPreferencesPathSetLink associates a link to a second dictionary at the specified path.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetLink(_:_:_:)
func SCPreferencesPathSetLink(prefs SCPreferencesRef, path corefoundation.CFStringRef, link corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesPathSetLink(prefs, path, link)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesPathSetValue func(prefs SCPreferencesRef, path corefoundation.CFStringRef, value corefoundation.CFDictionaryRef) bool
var _sCPreferencesPathSetValueErr error

func trySCPreferencesPathSetValue(prefs SCPreferencesRef, path corefoundation.CFStringRef, value corefoundation.CFDictionaryRef) (bool, error) {
	if _sCPreferencesPathSetValue == nil {
		return false, symbolCallError("SCPreferencesPathSetValue", "10.1", _sCPreferencesPathSetValueErr)
	}
	return _sCPreferencesPathSetValue(prefs, path, value), nil
}

// SCPreferencesPathSetValue associates the specified dictionary with the specified path.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetValue(_:_:_:)
func SCPreferencesPathSetValue(prefs SCPreferencesRef, path corefoundation.CFStringRef, value corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCPreferencesPathSetValue(prefs, path, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesRemoveValue func(prefs SCPreferencesRef, key corefoundation.CFStringRef) bool
var _sCPreferencesRemoveValueErr error

func trySCPreferencesRemoveValue(prefs SCPreferencesRef, key corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesRemoveValue == nil {
		return false, symbolCallError("SCPreferencesRemoveValue", "10.1", _sCPreferencesRemoveValueErr)
	}
	return _sCPreferencesRemoveValue(prefs, key), nil
}

// SCPreferencesRemoveValue removes the data associated with the specified preference key.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesRemoveValue(_:_:)
func SCPreferencesRemoveValue(prefs SCPreferencesRef, key corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesRemoveValue(prefs, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesScheduleWithRunLoop func(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCPreferencesScheduleWithRunLoopErr error

func trySCPreferencesScheduleWithRunLoop(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesScheduleWithRunLoop == nil {
		return false, symbolCallError("SCPreferencesScheduleWithRunLoop", "10.4", _sCPreferencesScheduleWithRunLoopErr)
	}
	return _sCPreferencesScheduleWithRunLoop(prefs, runLoop, runLoopMode), nil
}

// SCPreferencesScheduleWithRunLoop schedules commit and apply notifications for the specified preferences session using the specified run loop and mode.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesScheduleWithRunLoop(_:_:_:)
func SCPreferencesScheduleWithRunLoop(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesScheduleWithRunLoop(prefs, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSetCallback func(prefs SCPreferencesRef, callout SCPreferencesCallBack, context *SCPreferencesContext) bool
var _sCPreferencesSetCallbackErr error

func trySCPreferencesSetCallback(prefs SCPreferencesRef, callout SCPreferencesCallBack, context *SCPreferencesContext) (bool, error) {
	if _sCPreferencesSetCallback == nil {
		return false, symbolCallError("SCPreferencesSetCallback", "10.4", _sCPreferencesSetCallbackErr)
	}
	return _sCPreferencesSetCallback(prefs, callout, context), nil
}

// SCPreferencesSetCallback assigns the specified callback to the specified preferences session.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetCallback(_:_:_:)
func SCPreferencesSetCallback(prefs SCPreferencesRef, callout SCPreferencesCallBack, context *SCPreferencesContext) bool {
	result, callErr := trySCPreferencesSetCallback(prefs, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSetComputerName func(prefs SCPreferencesRef, name corefoundation.CFStringRef, nameEncoding uint32) bool
var _sCPreferencesSetComputerNameErr error

func trySCPreferencesSetComputerName(prefs SCPreferencesRef, name corefoundation.CFStringRef, nameEncoding uint32) (bool, error) {
	if _sCPreferencesSetComputerName == nil {
		return false, symbolCallError("SCPreferencesSetComputerName", "10.1", _sCPreferencesSetComputerNameErr)
	}
	return _sCPreferencesSetComputerName(prefs, name, nameEncoding), nil
}

// SCPreferencesSetComputerName sets the computer name preference to the specified name.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetComputerName(_:_:_:)
func SCPreferencesSetComputerName(prefs SCPreferencesRef, name corefoundation.CFStringRef, nameEncoding uint32) bool {
	result, callErr := trySCPreferencesSetComputerName(prefs, name, nameEncoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSetDispatchQueue func(prefs SCPreferencesRef, queue uintptr) bool
var _sCPreferencesSetDispatchQueueErr error

func trySCPreferencesSetDispatchQueue(prefs SCPreferencesRef, queue dispatch.Queue) (bool, error) {
	if _sCPreferencesSetDispatchQueue == nil {
		return false, symbolCallError("SCPreferencesSetDispatchQueue", "10.6", _sCPreferencesSetDispatchQueueErr)
	}
	return _sCPreferencesSetDispatchQueue(prefs, uintptr(queue.Handle())), nil
}

// SCPreferencesSetDispatchQueue schedules commit and apply notifications for the specified preferences session using the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetDispatchQueue(_:_:)
func SCPreferencesSetDispatchQueue(prefs SCPreferencesRef, queue dispatch.Queue) bool {
	result, callErr := trySCPreferencesSetDispatchQueue(prefs, queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSetLocalHostName func(prefs SCPreferencesRef, name corefoundation.CFStringRef) bool
var _sCPreferencesSetLocalHostNameErr error

func trySCPreferencesSetLocalHostName(prefs SCPreferencesRef, name corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesSetLocalHostName == nil {
		return false, symbolCallError("SCPreferencesSetLocalHostName", "10.2", _sCPreferencesSetLocalHostNameErr)
	}
	return _sCPreferencesSetLocalHostName(prefs, name), nil
}

// SCPreferencesSetLocalHostName sets the local host name to the specified name.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetLocalHostName(_:_:)
func SCPreferencesSetLocalHostName(prefs SCPreferencesRef, name corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesSetLocalHostName(prefs, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSetValue func(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool
var _sCPreferencesSetValueErr error

func trySCPreferencesSetValue(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) (bool, error) {
	if _sCPreferencesSetValue == nil {
		return false, symbolCallError("SCPreferencesSetValue", "10.1", _sCPreferencesSetValueErr)
	}
	return _sCPreferencesSetValue(prefs, key, value), nil
}

// SCPreferencesSetValue updates the data associated with the specified preference key with the specified value.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetValue(_:_:_:)
func SCPreferencesSetValue(prefs SCPreferencesRef, key corefoundation.CFStringRef, value corefoundation.CFPropertyListRef) bool {
	result, callErr := trySCPreferencesSetValue(prefs, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesSynchronize func(prefs SCPreferencesRef)
var _sCPreferencesSynchronizeErr error

func trySCPreferencesSynchronize(prefs SCPreferencesRef) error {
	if _sCPreferencesSynchronize == nil {
		return symbolCallError("SCPreferencesSynchronize", "10.4", _sCPreferencesSynchronizeErr)
	}
	_sCPreferencesSynchronize(prefs)
	return nil
}

// SCPreferencesSynchronize synchronizes accessed preferences with committed changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSynchronize(_:)
func SCPreferencesSynchronize(prefs SCPreferencesRef) {
	if callErr := trySCPreferencesSynchronize(prefs); callErr != nil {
		panic(callErr)
	}
}

var _sCPreferencesUnlock func(prefs SCPreferencesRef) bool
var _sCPreferencesUnlockErr error

func trySCPreferencesUnlock(prefs SCPreferencesRef) (bool, error) {
	if _sCPreferencesUnlock == nil {
		return false, symbolCallError("SCPreferencesUnlock", "10.1", _sCPreferencesUnlockErr)
	}
	return _sCPreferencesUnlock(prefs), nil
}

// SCPreferencesUnlock releases exclusive access to the configuration preferences.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnlock(_:)
func SCPreferencesUnlock(prefs SCPreferencesRef) bool {
	result, callErr := trySCPreferencesUnlock(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCPreferencesUnscheduleFromRunLoop func(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool
var _sCPreferencesUnscheduleFromRunLoopErr error

func trySCPreferencesUnscheduleFromRunLoop(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) (bool, error) {
	if _sCPreferencesUnscheduleFromRunLoop == nil {
		return false, symbolCallError("SCPreferencesUnscheduleFromRunLoop", "10.4", _sCPreferencesUnscheduleFromRunLoopErr)
	}
	return _sCPreferencesUnscheduleFromRunLoop(prefs, runLoop, runLoopMode), nil
}

// SCPreferencesUnscheduleFromRunLoop unschedules commit and apply notifications for the specified preferences session from the specified run loop and mode.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnscheduleFromRunLoop(_:_:_:)
func SCPreferencesUnscheduleFromRunLoop(prefs SCPreferencesRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	result, callErr := trySCPreferencesUnscheduleFromRunLoop(prefs, runLoop, runLoopMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceCopyAll func(prefs SCPreferencesRef) corefoundation.CFArrayRef
var _sCVLANInterfaceCopyAllErr error

func trySCVLANInterfaceCopyAll(prefs SCPreferencesRef) (corefoundation.CFArrayRef, error) {
	if _sCVLANInterfaceCopyAll == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCVLANInterfaceCopyAll", "10.5", _sCVLANInterfaceCopyAllErr)
	}
	return _sCVLANInterfaceCopyAll(prefs), nil
}

// SCVLANInterfaceCopyAll returns all virtual LAN (VLAN) interfaces on the system.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAll(_:)
func SCVLANInterfaceCopyAll(prefs SCPreferencesRef) corefoundation.CFArrayRef {
	result, callErr := trySCVLANInterfaceCopyAll(prefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceCopyAvailablePhysicalInterfaces func() corefoundation.CFArrayRef
var _sCVLANInterfaceCopyAvailablePhysicalInterfacesErr error

func trySCVLANInterfaceCopyAvailablePhysicalInterfaces() (corefoundation.CFArrayRef, error) {
	if _sCVLANInterfaceCopyAvailablePhysicalInterfaces == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("SCVLANInterfaceCopyAvailablePhysicalInterfaces", "10.5", _sCVLANInterfaceCopyAvailablePhysicalInterfacesErr)
	}
	return _sCVLANInterfaceCopyAvailablePhysicalInterfaces(), nil
}

// SCVLANInterfaceCopyAvailablePhysicalInterfaces returns the network capable devices on the system that can be associated with a virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAvailablePhysicalInterfaces()
func SCVLANInterfaceCopyAvailablePhysicalInterfaces() corefoundation.CFArrayRef {
	result, callErr := trySCVLANInterfaceCopyAvailablePhysicalInterfaces()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceCreate func(prefs SCPreferencesRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) SCVLANInterfaceRef
var _sCVLANInterfaceCreateErr error

func trySCVLANInterfaceCreate(prefs SCPreferencesRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) (SCVLANInterfaceRef, error) {
	if _sCVLANInterfaceCreate == nil {
		return *new(SCVLANInterfaceRef), symbolCallError("SCVLANInterfaceCreate", "10.5", _sCVLANInterfaceCreateErr)
	}
	return _sCVLANInterfaceCreate(prefs, physical, tag), nil
}

// SCVLANInterfaceCreate creates a new virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCreate(_:_:_:)
func SCVLANInterfaceCreate(prefs SCPreferencesRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) SCVLANInterfaceRef {
	result, callErr := trySCVLANInterfaceCreate(prefs, physical, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceGetOptions func(vlan SCVLANInterfaceRef) corefoundation.CFDictionaryRef
var _sCVLANInterfaceGetOptionsErr error

func trySCVLANInterfaceGetOptions(vlan SCVLANInterfaceRef) (corefoundation.CFDictionaryRef, error) {
	if _sCVLANInterfaceGetOptions == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("SCVLANInterfaceGetOptions", "10.5", _sCVLANInterfaceGetOptionsErr)
	}
	return _sCVLANInterfaceGetOptions(vlan), nil
}

// SCVLANInterfaceGetOptions returns the configuration settings associated with the virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetOptions(_:)
func SCVLANInterfaceGetOptions(vlan SCVLANInterfaceRef) corefoundation.CFDictionaryRef {
	result, callErr := trySCVLANInterfaceGetOptions(vlan)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceGetPhysicalInterface func(vlan SCVLANInterfaceRef) SCNetworkInterfaceRef
var _sCVLANInterfaceGetPhysicalInterfaceErr error

func trySCVLANInterfaceGetPhysicalInterface(vlan SCVLANInterfaceRef) (SCNetworkInterfaceRef, error) {
	if _sCVLANInterfaceGetPhysicalInterface == nil {
		return *new(SCNetworkInterfaceRef), symbolCallError("SCVLANInterfaceGetPhysicalInterface", "10.5", _sCVLANInterfaceGetPhysicalInterfaceErr)
	}
	return _sCVLANInterfaceGetPhysicalInterface(vlan), nil
}

// SCVLANInterfaceGetPhysicalInterface returns the physical interface for the specified virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetPhysicalInterface(_:)
func SCVLANInterfaceGetPhysicalInterface(vlan SCVLANInterfaceRef) SCNetworkInterfaceRef {
	result, callErr := trySCVLANInterfaceGetPhysicalInterface(vlan)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceGetTag func(vlan SCVLANInterfaceRef) corefoundation.CFNumberRef
var _sCVLANInterfaceGetTagErr error

func trySCVLANInterfaceGetTag(vlan SCVLANInterfaceRef) (corefoundation.CFNumberRef, error) {
	if _sCVLANInterfaceGetTag == nil {
		return *new(corefoundation.CFNumberRef), symbolCallError("SCVLANInterfaceGetTag", "10.5", _sCVLANInterfaceGetTagErr)
	}
	return _sCVLANInterfaceGetTag(vlan), nil
}

// SCVLANInterfaceGetTag returns the tag for the specified virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetTag(_:)
func SCVLANInterfaceGetTag(vlan SCVLANInterfaceRef) corefoundation.CFNumberRef {
	result, callErr := trySCVLANInterfaceGetTag(vlan)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceRemove func(vlan SCVLANInterfaceRef) bool
var _sCVLANInterfaceRemoveErr error

func trySCVLANInterfaceRemove(vlan SCVLANInterfaceRef) (bool, error) {
	if _sCVLANInterfaceRemove == nil {
		return false, symbolCallError("SCVLANInterfaceRemove", "10.5", _sCVLANInterfaceRemoveErr)
	}
	return _sCVLANInterfaceRemove(vlan), nil
}

// SCVLANInterfaceRemove removes the virtual LAN (VLAN) interface from the configuration.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceRemove(_:)
func SCVLANInterfaceRemove(vlan SCVLANInterfaceRef) bool {
	result, callErr := trySCVLANInterfaceRemove(vlan)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceSetLocalizedDisplayName func(vlan SCVLANInterfaceRef, newName corefoundation.CFStringRef) bool
var _sCVLANInterfaceSetLocalizedDisplayNameErr error

func trySCVLANInterfaceSetLocalizedDisplayName(vlan SCVLANInterfaceRef, newName corefoundation.CFStringRef) (bool, error) {
	if _sCVLANInterfaceSetLocalizedDisplayName == nil {
		return false, symbolCallError("SCVLANInterfaceSetLocalizedDisplayName", "10.5", _sCVLANInterfaceSetLocalizedDisplayNameErr)
	}
	return _sCVLANInterfaceSetLocalizedDisplayName(vlan, newName), nil
}

// SCVLANInterfaceSetLocalizedDisplayName sets the localized display name for the specified virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetLocalizedDisplayName(_:_:)
func SCVLANInterfaceSetLocalizedDisplayName(vlan SCVLANInterfaceRef, newName corefoundation.CFStringRef) bool {
	result, callErr := trySCVLANInterfaceSetLocalizedDisplayName(vlan, newName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceSetOptions func(vlan SCVLANInterfaceRef, newOptions corefoundation.CFDictionaryRef) bool
var _sCVLANInterfaceSetOptionsErr error

func trySCVLANInterfaceSetOptions(vlan SCVLANInterfaceRef, newOptions corefoundation.CFDictionaryRef) (bool, error) {
	if _sCVLANInterfaceSetOptions == nil {
		return false, symbolCallError("SCVLANInterfaceSetOptions", "10.5", _sCVLANInterfaceSetOptionsErr)
	}
	return _sCVLANInterfaceSetOptions(vlan, newOptions), nil
}

// SCVLANInterfaceSetOptions sets the specified configuration settings for the specified virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetOptions(_:_:)
func SCVLANInterfaceSetOptions(vlan SCVLANInterfaceRef, newOptions corefoundation.CFDictionaryRef) bool {
	result, callErr := trySCVLANInterfaceSetOptions(vlan, newOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sCVLANInterfaceSetPhysicalInterfaceAndTag func(vlan SCVLANInterfaceRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) bool
var _sCVLANInterfaceSetPhysicalInterfaceAndTagErr error

func trySCVLANInterfaceSetPhysicalInterfaceAndTag(vlan SCVLANInterfaceRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) (bool, error) {
	if _sCVLANInterfaceSetPhysicalInterfaceAndTag == nil {
		return false, symbolCallError("SCVLANInterfaceSetPhysicalInterfaceAndTag", "10.5", _sCVLANInterfaceSetPhysicalInterfaceAndTagErr)
	}
	return _sCVLANInterfaceSetPhysicalInterfaceAndTag(vlan, physical, tag), nil
}

// SCVLANInterfaceSetPhysicalInterfaceAndTag updates the specified virtual LAN (VLAN) interface with the specified information.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetPhysicalInterfaceAndTag(_:_:_:)
func SCVLANInterfaceSetPhysicalInterfaceAndTag(vlan SCVLANInterfaceRef, physical SCNetworkInterfaceRef, tag corefoundation.CFNumberRef) bool {
	result, callErr := trySCVLANInterfaceSetPhysicalInterfaceAndTag(vlan, physical, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cNCopySupportedInterfaces, &_cNCopySupportedInterfacesErr, frameworkHandle, "CNCopySupportedInterfaces", "10.8")
	registerFunc(&_cNMarkPortalOffline, &_cNMarkPortalOfflineErr, frameworkHandle, "CNMarkPortalOffline", "10.8")
	registerFunc(&_cNMarkPortalOnline, &_cNMarkPortalOnlineErr, frameworkHandle, "CNMarkPortalOnline", "10.8")
	registerFunc(&_cNSetSupportedSSIDs, &_cNSetSupportedSSIDsErr, frameworkHandle, "CNSetSupportedSSIDs", "10.8")
	registerFunc(&_dHCPClientPreferencesCopyApplicationOptions, &_dHCPClientPreferencesCopyApplicationOptionsErr, frameworkHandle, "DHCPClientPreferencesCopyApplicationOptions", "10.1")
	registerFunc(&_dHCPClientPreferencesSetApplicationOptions, &_dHCPClientPreferencesSetApplicationOptionsErr, frameworkHandle, "DHCPClientPreferencesSetApplicationOptions", "10.1")
	registerFunc(&_dHCPInfoGetLeaseExpirationTime, &_dHCPInfoGetLeaseExpirationTimeErr, frameworkHandle, "DHCPInfoGetLeaseExpirationTime", "10.8")
	registerFunc(&_dHCPInfoGetLeaseStartTime, &_dHCPInfoGetLeaseStartTimeErr, frameworkHandle, "DHCPInfoGetLeaseStartTime", "10.1")
	registerFunc(&_dHCPInfoGetOptionData, &_dHCPInfoGetOptionDataErr, frameworkHandle, "DHCPInfoGetOptionData", "10.1")
	registerFunc(&_sCBondInterfaceCopyAll, &_sCBondInterfaceCopyAllErr, frameworkHandle, "SCBondInterfaceCopyAll", "10.5")
	registerFunc(&_sCBondInterfaceCopyAvailableMemberInterfaces, &_sCBondInterfaceCopyAvailableMemberInterfacesErr, frameworkHandle, "SCBondInterfaceCopyAvailableMemberInterfaces", "10.5")
	registerFunc(&_sCBondInterfaceCopyStatus, &_sCBondInterfaceCopyStatusErr, frameworkHandle, "SCBondInterfaceCopyStatus", "10.5")
	registerFunc(&_sCBondInterfaceCreate, &_sCBondInterfaceCreateErr, frameworkHandle, "SCBondInterfaceCreate", "10.5")
	registerFunc(&_sCBondInterfaceGetMemberInterfaces, &_sCBondInterfaceGetMemberInterfacesErr, frameworkHandle, "SCBondInterfaceGetMemberInterfaces", "10.5")
	registerFunc(&_sCBondInterfaceGetOptions, &_sCBondInterfaceGetOptionsErr, frameworkHandle, "SCBondInterfaceGetOptions", "10.5")
	registerFunc(&_sCBondInterfaceRemove, &_sCBondInterfaceRemoveErr, frameworkHandle, "SCBondInterfaceRemove", "10.5")
	registerFunc(&_sCBondInterfaceSetLocalizedDisplayName, &_sCBondInterfaceSetLocalizedDisplayNameErr, frameworkHandle, "SCBondInterfaceSetLocalizedDisplayName", "10.5")
	registerFunc(&_sCBondInterfaceSetMemberInterfaces, &_sCBondInterfaceSetMemberInterfacesErr, frameworkHandle, "SCBondInterfaceSetMemberInterfaces", "10.5")
	registerFunc(&_sCBondInterfaceSetOptions, &_sCBondInterfaceSetOptionsErr, frameworkHandle, "SCBondInterfaceSetOptions", "10.5")
	registerFunc(&_sCBondStatusGetInterfaceStatus, &_sCBondStatusGetInterfaceStatusErr, frameworkHandle, "SCBondStatusGetInterfaceStatus", "10.5")
	registerFunc(&_sCBondStatusGetMemberInterfaces, &_sCBondStatusGetMemberInterfacesErr, frameworkHandle, "SCBondStatusGetMemberInterfaces", "10.5")
	registerFunc(&_sCBondStatusGetTypeID, &_sCBondStatusGetTypeIDErr, frameworkHandle, "SCBondStatusGetTypeID", "10.5")
	registerFunc(&_sCCopyLastError, &_sCCopyLastErrorErr, frameworkHandle, "SCCopyLastError", "10.5")
	registerFunc(&_sCDynamicStoreAddTemporaryValue, &_sCDynamicStoreAddTemporaryValueErr, frameworkHandle, "SCDynamicStoreAddTemporaryValue", "10.1")
	registerFunc(&_sCDynamicStoreAddValue, &_sCDynamicStoreAddValueErr, frameworkHandle, "SCDynamicStoreAddValue", "10.1")
	registerFunc(&_sCDynamicStoreCopyComputerName, &_sCDynamicStoreCopyComputerNameErr, frameworkHandle, "SCDynamicStoreCopyComputerName", "10.1")
	registerFunc(&_sCDynamicStoreCopyConsoleUser, &_sCDynamicStoreCopyConsoleUserErr, frameworkHandle, "SCDynamicStoreCopyConsoleUser", "10.1")
	registerFunc(&_sCDynamicStoreCopyDHCPInfo, &_sCDynamicStoreCopyDHCPInfoErr, frameworkHandle, "SCDynamicStoreCopyDHCPInfo", "10.1")
	registerFunc(&_sCDynamicStoreCopyKeyList, &_sCDynamicStoreCopyKeyListErr, frameworkHandle, "SCDynamicStoreCopyKeyList", "10.1")
	registerFunc(&_sCDynamicStoreCopyLocalHostName, &_sCDynamicStoreCopyLocalHostNameErr, frameworkHandle, "SCDynamicStoreCopyLocalHostName", "10.1")
	registerFunc(&_sCDynamicStoreCopyLocation, &_sCDynamicStoreCopyLocationErr, frameworkHandle, "SCDynamicStoreCopyLocation", "10.1")
	registerFunc(&_sCDynamicStoreCopyMultiple, &_sCDynamicStoreCopyMultipleErr, frameworkHandle, "SCDynamicStoreCopyMultiple", "10.1")
	registerFunc(&_sCDynamicStoreCopyNotifiedKeys, &_sCDynamicStoreCopyNotifiedKeysErr, frameworkHandle, "SCDynamicStoreCopyNotifiedKeys", "10.1")
	registerFunc(&_sCDynamicStoreCopyProxies, &_sCDynamicStoreCopyProxiesErr, frameworkHandle, "SCDynamicStoreCopyProxies", "10.1")
	registerFunc(&_sCDynamicStoreCopyValue, &_sCDynamicStoreCopyValueErr, frameworkHandle, "SCDynamicStoreCopyValue", "10.1")
	registerFunc(&_sCDynamicStoreCreate, &_sCDynamicStoreCreateErr, frameworkHandle, "SCDynamicStoreCreate", "10.1")
	registerFunc(&_sCDynamicStoreCreateRunLoopSource, &_sCDynamicStoreCreateRunLoopSourceErr, frameworkHandle, "SCDynamicStoreCreateRunLoopSource", "10.1")
	registerFunc(&_sCDynamicStoreCreateWithOptions, &_sCDynamicStoreCreateWithOptionsErr, frameworkHandle, "SCDynamicStoreCreateWithOptions", "10.4")
	registerFunc(&_sCDynamicStoreGetTypeID, &_sCDynamicStoreGetTypeIDErr, frameworkHandle, "SCDynamicStoreGetTypeID", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreate, &_sCDynamicStoreKeyCreateErr, frameworkHandle, "SCDynamicStoreKeyCreate", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateComputerName, &_sCDynamicStoreKeyCreateComputerNameErr, frameworkHandle, "SCDynamicStoreKeyCreateComputerName", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateConsoleUser, &_sCDynamicStoreKeyCreateConsoleUserErr, frameworkHandle, "SCDynamicStoreKeyCreateConsoleUser", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateHostNames, &_sCDynamicStoreKeyCreateHostNamesErr, frameworkHandle, "SCDynamicStoreKeyCreateHostNames", "10.2")
	registerFunc(&_sCDynamicStoreKeyCreateLocation, &_sCDynamicStoreKeyCreateLocationErr, frameworkHandle, "SCDynamicStoreKeyCreateLocation", "10.2")
	registerFunc(&_sCDynamicStoreKeyCreateNetworkGlobalEntity, &_sCDynamicStoreKeyCreateNetworkGlobalEntityErr, frameworkHandle, "SCDynamicStoreKeyCreateNetworkGlobalEntity", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateNetworkInterface, &_sCDynamicStoreKeyCreateNetworkInterfaceErr, frameworkHandle, "SCDynamicStoreKeyCreateNetworkInterface", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateNetworkInterfaceEntity, &_sCDynamicStoreKeyCreateNetworkInterfaceEntityErr, frameworkHandle, "SCDynamicStoreKeyCreateNetworkInterfaceEntity", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateNetworkServiceEntity, &_sCDynamicStoreKeyCreateNetworkServiceEntityErr, frameworkHandle, "SCDynamicStoreKeyCreateNetworkServiceEntity", "10.1")
	registerFunc(&_sCDynamicStoreKeyCreateProxies, &_sCDynamicStoreKeyCreateProxiesErr, frameworkHandle, "SCDynamicStoreKeyCreateProxies", "10.1")
	registerFunc(&_sCDynamicStoreNotifyValue, &_sCDynamicStoreNotifyValueErr, frameworkHandle, "SCDynamicStoreNotifyValue", "10.1")
	registerFunc(&_sCDynamicStoreRemoveValue, &_sCDynamicStoreRemoveValueErr, frameworkHandle, "SCDynamicStoreRemoveValue", "10.1")
	registerFunc(&_sCDynamicStoreSetDispatchQueue, &_sCDynamicStoreSetDispatchQueueErr, frameworkHandle, "SCDynamicStoreSetDispatchQueue", "10.6")
	registerFunc(&_sCDynamicStoreSetMultiple, &_sCDynamicStoreSetMultipleErr, frameworkHandle, "SCDynamicStoreSetMultiple", "10.1")
	registerFunc(&_sCDynamicStoreSetNotificationKeys, &_sCDynamicStoreSetNotificationKeysErr, frameworkHandle, "SCDynamicStoreSetNotificationKeys", "10.1")
	registerFunc(&_sCDynamicStoreSetValue, &_sCDynamicStoreSetValueErr, frameworkHandle, "SCDynamicStoreSetValue", "10.1")
	registerFunc(&_sCError, &_sCErrorErr, frameworkHandle, "SCError", "10.1")
	registerFunc(&_sCErrorString, &_sCErrorStringErr, frameworkHandle, "SCErrorString", "10.1")
	registerFunc(&_sCNetworkConnectionCopyExtendedStatus, &_sCNetworkConnectionCopyExtendedStatusErr, frameworkHandle, "SCNetworkConnectionCopyExtendedStatus", "10.3")
	registerFunc(&_sCNetworkConnectionCopyServiceID, &_sCNetworkConnectionCopyServiceIDErr, frameworkHandle, "SCNetworkConnectionCopyServiceID", "10.3")
	registerFunc(&_sCNetworkConnectionCopyStatistics, &_sCNetworkConnectionCopyStatisticsErr, frameworkHandle, "SCNetworkConnectionCopyStatistics", "10.3")
	registerFunc(&_sCNetworkConnectionCopyUserOptions, &_sCNetworkConnectionCopyUserOptionsErr, frameworkHandle, "SCNetworkConnectionCopyUserOptions", "10.3")
	registerFunc(&_sCNetworkConnectionCopyUserPreferences, &_sCNetworkConnectionCopyUserPreferencesErr, frameworkHandle, "SCNetworkConnectionCopyUserPreferences", "10.3")
	registerFunc(&_sCNetworkConnectionCreateWithServiceID, &_sCNetworkConnectionCreateWithServiceIDErr, frameworkHandle, "SCNetworkConnectionCreateWithServiceID", "10.3")
	registerFunc(&_sCNetworkConnectionGetStatus, &_sCNetworkConnectionGetStatusErr, frameworkHandle, "SCNetworkConnectionGetStatus", "10.3")
	registerFunc(&_sCNetworkConnectionGetTypeID, &_sCNetworkConnectionGetTypeIDErr, frameworkHandle, "SCNetworkConnectionGetTypeID", "10.3")
	registerFunc(&_sCNetworkConnectionScheduleWithRunLoop, &_sCNetworkConnectionScheduleWithRunLoopErr, frameworkHandle, "SCNetworkConnectionScheduleWithRunLoop", "10.3")
	registerFunc(&_sCNetworkConnectionSetDispatchQueue, &_sCNetworkConnectionSetDispatchQueueErr, frameworkHandle, "SCNetworkConnectionSetDispatchQueue", "10.6")
	registerFunc(&_sCNetworkConnectionStart, &_sCNetworkConnectionStartErr, frameworkHandle, "SCNetworkConnectionStart", "10.3")
	registerFunc(&_sCNetworkConnectionStop, &_sCNetworkConnectionStopErr, frameworkHandle, "SCNetworkConnectionStop", "10.3")
	registerFunc(&_sCNetworkConnectionUnscheduleFromRunLoop, &_sCNetworkConnectionUnscheduleFromRunLoopErr, frameworkHandle, "SCNetworkConnectionUnscheduleFromRunLoop", "10.3")
	registerFunc(&_sCNetworkInterfaceCopyAll, &_sCNetworkInterfaceCopyAllErr, frameworkHandle, "SCNetworkInterfaceCopyAll", "10.4")
	registerFunc(&_sCNetworkInterfaceCopyMTU, &_sCNetworkInterfaceCopyMTUErr, frameworkHandle, "SCNetworkInterfaceCopyMTU", "10.5")
	registerFunc(&_sCNetworkInterfaceCopyMediaOptions, &_sCNetworkInterfaceCopyMediaOptionsErr, frameworkHandle, "SCNetworkInterfaceCopyMediaOptions", "10.5")
	registerFunc(&_sCNetworkInterfaceCopyMediaSubTypeOptions, &_sCNetworkInterfaceCopyMediaSubTypeOptionsErr, frameworkHandle, "SCNetworkInterfaceCopyMediaSubTypeOptions", "10.5")
	registerFunc(&_sCNetworkInterfaceCopyMediaSubTypes, &_sCNetworkInterfaceCopyMediaSubTypesErr, frameworkHandle, "SCNetworkInterfaceCopyMediaSubTypes", "10.5")
	registerFunc(&_sCNetworkInterfaceCreateWithInterface, &_sCNetworkInterfaceCreateWithInterfaceErr, frameworkHandle, "SCNetworkInterfaceCreateWithInterface", "10.4")
	registerFunc(&_sCNetworkInterfaceForceConfigurationRefresh, &_sCNetworkInterfaceForceConfigurationRefreshErr, frameworkHandle, "SCNetworkInterfaceForceConfigurationRefresh", "10.5")
	registerFunc(&_sCNetworkInterfaceGetBSDName, &_sCNetworkInterfaceGetBSDNameErr, frameworkHandle, "SCNetworkInterfaceGetBSDName", "10.4")
	registerFunc(&_sCNetworkInterfaceGetConfiguration, &_sCNetworkInterfaceGetConfigurationErr, frameworkHandle, "SCNetworkInterfaceGetConfiguration", "10.4")
	registerFunc(&_sCNetworkInterfaceGetExtendedConfiguration, &_sCNetworkInterfaceGetExtendedConfigurationErr, frameworkHandle, "SCNetworkInterfaceGetExtendedConfiguration", "10.5")
	registerFunc(&_sCNetworkInterfaceGetHardwareAddressString, &_sCNetworkInterfaceGetHardwareAddressStringErr, frameworkHandle, "SCNetworkInterfaceGetHardwareAddressString", "10.4")
	registerFunc(&_sCNetworkInterfaceGetInterface, &_sCNetworkInterfaceGetInterfaceErr, frameworkHandle, "SCNetworkInterfaceGetInterface", "10.4")
	registerFunc(&_sCNetworkInterfaceGetInterfaceType, &_sCNetworkInterfaceGetInterfaceTypeErr, frameworkHandle, "SCNetworkInterfaceGetInterfaceType", "10.4")
	registerFunc(&_sCNetworkInterfaceGetLocalizedDisplayName, &_sCNetworkInterfaceGetLocalizedDisplayNameErr, frameworkHandle, "SCNetworkInterfaceGetLocalizedDisplayName", "10.4")
	registerFunc(&_sCNetworkInterfaceGetSupportedInterfaceTypes, &_sCNetworkInterfaceGetSupportedInterfaceTypesErr, frameworkHandle, "SCNetworkInterfaceGetSupportedInterfaceTypes", "10.4")
	registerFunc(&_sCNetworkInterfaceGetSupportedProtocolTypes, &_sCNetworkInterfaceGetSupportedProtocolTypesErr, frameworkHandle, "SCNetworkInterfaceGetSupportedProtocolTypes", "10.4")
	registerFunc(&_sCNetworkInterfaceGetTypeID, &_sCNetworkInterfaceGetTypeIDErr, frameworkHandle, "SCNetworkInterfaceGetTypeID", "10.4")
	registerFunc(&_sCNetworkInterfaceSetConfiguration, &_sCNetworkInterfaceSetConfigurationErr, frameworkHandle, "SCNetworkInterfaceSetConfiguration", "10.4")
	registerFunc(&_sCNetworkInterfaceSetExtendedConfiguration, &_sCNetworkInterfaceSetExtendedConfigurationErr, frameworkHandle, "SCNetworkInterfaceSetExtendedConfiguration", "10.5")
	registerFunc(&_sCNetworkInterfaceSetMTU, &_sCNetworkInterfaceSetMTUErr, frameworkHandle, "SCNetworkInterfaceSetMTU", "10.5")
	registerFunc(&_sCNetworkInterfaceSetMediaOptions, &_sCNetworkInterfaceSetMediaOptionsErr, frameworkHandle, "SCNetworkInterfaceSetMediaOptions", "10.5")
	registerFunc(&_sCNetworkProtocolGetConfiguration, &_sCNetworkProtocolGetConfigurationErr, frameworkHandle, "SCNetworkProtocolGetConfiguration", "10.4")
	registerFunc(&_sCNetworkProtocolGetEnabled, &_sCNetworkProtocolGetEnabledErr, frameworkHandle, "SCNetworkProtocolGetEnabled", "10.4")
	registerFunc(&_sCNetworkProtocolGetProtocolType, &_sCNetworkProtocolGetProtocolTypeErr, frameworkHandle, "SCNetworkProtocolGetProtocolType", "10.4")
	registerFunc(&_sCNetworkProtocolGetTypeID, &_sCNetworkProtocolGetTypeIDErr, frameworkHandle, "SCNetworkProtocolGetTypeID", "10.4")
	registerFunc(&_sCNetworkProtocolSetConfiguration, &_sCNetworkProtocolSetConfigurationErr, frameworkHandle, "SCNetworkProtocolSetConfiguration", "10.4")
	registerFunc(&_sCNetworkProtocolSetEnabled, &_sCNetworkProtocolSetEnabledErr, frameworkHandle, "SCNetworkProtocolSetEnabled", "10.4")
	registerFunc(&_sCNetworkReachabilityCreateWithAddress, &_sCNetworkReachabilityCreateWithAddressErr, frameworkHandle, "SCNetworkReachabilityCreateWithAddress", "10.3")
	registerFunc(&_sCNetworkReachabilityCreateWithAddressPair, &_sCNetworkReachabilityCreateWithAddressPairErr, frameworkHandle, "SCNetworkReachabilityCreateWithAddressPair", "10.3")
	registerFunc(&_sCNetworkReachabilityCreateWithName, &_sCNetworkReachabilityCreateWithNameErr, frameworkHandle, "SCNetworkReachabilityCreateWithName", "10.3")
	registerFunc(&_sCNetworkReachabilityGetFlags, &_sCNetworkReachabilityGetFlagsErr, frameworkHandle, "SCNetworkReachabilityGetFlags", "10.3")
	registerFunc(&_sCNetworkReachabilityGetTypeID, &_sCNetworkReachabilityGetTypeIDErr, frameworkHandle, "SCNetworkReachabilityGetTypeID", "10.3")
	registerFunc(&_sCNetworkReachabilityScheduleWithRunLoop, &_sCNetworkReachabilityScheduleWithRunLoopErr, frameworkHandle, "SCNetworkReachabilityScheduleWithRunLoop", "10.3")
	registerFunc(&_sCNetworkReachabilitySetCallback, &_sCNetworkReachabilitySetCallbackErr, frameworkHandle, "SCNetworkReachabilitySetCallback", "10.3")
	registerFunc(&_sCNetworkReachabilitySetDispatchQueue, &_sCNetworkReachabilitySetDispatchQueueErr, frameworkHandle, "SCNetworkReachabilitySetDispatchQueue", "10.6")
	registerFunc(&_sCNetworkReachabilityUnscheduleFromRunLoop, &_sCNetworkReachabilityUnscheduleFromRunLoopErr, frameworkHandle, "SCNetworkReachabilityUnscheduleFromRunLoop", "10.3")
	registerFunc(&_sCNetworkServiceAddProtocolType, &_sCNetworkServiceAddProtocolTypeErr, frameworkHandle, "SCNetworkServiceAddProtocolType", "10.4")
	registerFunc(&_sCNetworkServiceCopy, &_sCNetworkServiceCopyErr, frameworkHandle, "SCNetworkServiceCopy", "10.4")
	registerFunc(&_sCNetworkServiceCopyAll, &_sCNetworkServiceCopyAllErr, frameworkHandle, "SCNetworkServiceCopyAll", "10.4")
	registerFunc(&_sCNetworkServiceCopyProtocol, &_sCNetworkServiceCopyProtocolErr, frameworkHandle, "SCNetworkServiceCopyProtocol", "10.4")
	registerFunc(&_sCNetworkServiceCopyProtocols, &_sCNetworkServiceCopyProtocolsErr, frameworkHandle, "SCNetworkServiceCopyProtocols", "10.4")
	registerFunc(&_sCNetworkServiceCreate, &_sCNetworkServiceCreateErr, frameworkHandle, "SCNetworkServiceCreate", "10.4")
	registerFunc(&_sCNetworkServiceEstablishDefaultConfiguration, &_sCNetworkServiceEstablishDefaultConfigurationErr, frameworkHandle, "SCNetworkServiceEstablishDefaultConfiguration", "10.5")
	registerFunc(&_sCNetworkServiceGetEnabled, &_sCNetworkServiceGetEnabledErr, frameworkHandle, "SCNetworkServiceGetEnabled", "10.4")
	registerFunc(&_sCNetworkServiceGetInterface, &_sCNetworkServiceGetInterfaceErr, frameworkHandle, "SCNetworkServiceGetInterface", "10.4")
	registerFunc(&_sCNetworkServiceGetName, &_sCNetworkServiceGetNameErr, frameworkHandle, "SCNetworkServiceGetName", "10.4")
	registerFunc(&_sCNetworkServiceGetServiceID, &_sCNetworkServiceGetServiceIDErr, frameworkHandle, "SCNetworkServiceGetServiceID", "10.4")
	registerFunc(&_sCNetworkServiceGetTypeID, &_sCNetworkServiceGetTypeIDErr, frameworkHandle, "SCNetworkServiceGetTypeID", "10.4")
	registerFunc(&_sCNetworkServiceRemove, &_sCNetworkServiceRemoveErr, frameworkHandle, "SCNetworkServiceRemove", "10.4")
	registerFunc(&_sCNetworkServiceRemoveProtocolType, &_sCNetworkServiceRemoveProtocolTypeErr, frameworkHandle, "SCNetworkServiceRemoveProtocolType", "10.4")
	registerFunc(&_sCNetworkServiceSetEnabled, &_sCNetworkServiceSetEnabledErr, frameworkHandle, "SCNetworkServiceSetEnabled", "10.4")
	registerFunc(&_sCNetworkServiceSetName, &_sCNetworkServiceSetNameErr, frameworkHandle, "SCNetworkServiceSetName", "10.4")
	registerFunc(&_sCNetworkSetAddService, &_sCNetworkSetAddServiceErr, frameworkHandle, "SCNetworkSetAddService", "10.4")
	registerFunc(&_sCNetworkSetContainsInterface, &_sCNetworkSetContainsInterfaceErr, frameworkHandle, "SCNetworkSetContainsInterface", "10.5")
	registerFunc(&_sCNetworkSetCopy, &_sCNetworkSetCopyErr, frameworkHandle, "SCNetworkSetCopy", "10.4")
	registerFunc(&_sCNetworkSetCopyAll, &_sCNetworkSetCopyAllErr, frameworkHandle, "SCNetworkSetCopyAll", "10.4")
	registerFunc(&_sCNetworkSetCopyCurrent, &_sCNetworkSetCopyCurrentErr, frameworkHandle, "SCNetworkSetCopyCurrent", "10.4")
	registerFunc(&_sCNetworkSetCopyServices, &_sCNetworkSetCopyServicesErr, frameworkHandle, "SCNetworkSetCopyServices", "10.4")
	registerFunc(&_sCNetworkSetCreate, &_sCNetworkSetCreateErr, frameworkHandle, "SCNetworkSetCreate", "10.4")
	registerFunc(&_sCNetworkSetGetName, &_sCNetworkSetGetNameErr, frameworkHandle, "SCNetworkSetGetName", "10.4")
	registerFunc(&_sCNetworkSetGetServiceOrder, &_sCNetworkSetGetServiceOrderErr, frameworkHandle, "SCNetworkSetGetServiceOrder", "10.4")
	registerFunc(&_sCNetworkSetGetSetID, &_sCNetworkSetGetSetIDErr, frameworkHandle, "SCNetworkSetGetSetID", "10.4")
	registerFunc(&_sCNetworkSetGetTypeID, &_sCNetworkSetGetTypeIDErr, frameworkHandle, "SCNetworkSetGetTypeID", "10.4")
	registerFunc(&_sCNetworkSetRemove, &_sCNetworkSetRemoveErr, frameworkHandle, "SCNetworkSetRemove", "10.4")
	registerFunc(&_sCNetworkSetRemoveService, &_sCNetworkSetRemoveServiceErr, frameworkHandle, "SCNetworkSetRemoveService", "10.4")
	registerFunc(&_sCNetworkSetSetCurrent, &_sCNetworkSetSetCurrentErr, frameworkHandle, "SCNetworkSetSetCurrent", "10.4")
	registerFunc(&_sCNetworkSetSetName, &_sCNetworkSetSetNameErr, frameworkHandle, "SCNetworkSetSetName", "10.4")
	registerFunc(&_sCNetworkSetSetServiceOrder, &_sCNetworkSetSetServiceOrderErr, frameworkHandle, "SCNetworkSetSetServiceOrder", "10.4")
	registerFunc(&_sCPreferencesAddValue, &_sCPreferencesAddValueErr, frameworkHandle, "SCPreferencesAddValue", "10.1")
	registerFunc(&_sCPreferencesApplyChanges, &_sCPreferencesApplyChangesErr, frameworkHandle, "SCPreferencesApplyChanges", "10.1")
	registerFunc(&_sCPreferencesCommitChanges, &_sCPreferencesCommitChangesErr, frameworkHandle, "SCPreferencesCommitChanges", "10.1")
	registerFunc(&_sCPreferencesCopyKeyList, &_sCPreferencesCopyKeyListErr, frameworkHandle, "SCPreferencesCopyKeyList", "10.1")
	registerFunc(&_sCPreferencesCreate, &_sCPreferencesCreateErr, frameworkHandle, "SCPreferencesCreate", "10.1")
	registerFunc(&_sCPreferencesCreateWithAuthorization, &_sCPreferencesCreateWithAuthorizationErr, frameworkHandle, "SCPreferencesCreateWithAuthorization", "10.5")
	registerFunc(&_sCPreferencesGetSignature, &_sCPreferencesGetSignatureErr, frameworkHandle, "SCPreferencesGetSignature", "10.1")
	registerFunc(&_sCPreferencesGetTypeID, &_sCPreferencesGetTypeIDErr, frameworkHandle, "SCPreferencesGetTypeID", "10.1")
	registerFunc(&_sCPreferencesGetValue, &_sCPreferencesGetValueErr, frameworkHandle, "SCPreferencesGetValue", "10.1")
	registerFunc(&_sCPreferencesLock, &_sCPreferencesLockErr, frameworkHandle, "SCPreferencesLock", "10.1")
	registerFunc(&_sCPreferencesPathCreateUniqueChild, &_sCPreferencesPathCreateUniqueChildErr, frameworkHandle, "SCPreferencesPathCreateUniqueChild", "10.1")
	registerFunc(&_sCPreferencesPathGetLink, &_sCPreferencesPathGetLinkErr, frameworkHandle, "SCPreferencesPathGetLink", "10.1")
	registerFunc(&_sCPreferencesPathGetValue, &_sCPreferencesPathGetValueErr, frameworkHandle, "SCPreferencesPathGetValue", "10.1")
	registerFunc(&_sCPreferencesPathRemoveValue, &_sCPreferencesPathRemoveValueErr, frameworkHandle, "SCPreferencesPathRemoveValue", "10.1")
	registerFunc(&_sCPreferencesPathSetLink, &_sCPreferencesPathSetLinkErr, frameworkHandle, "SCPreferencesPathSetLink", "10.1")
	registerFunc(&_sCPreferencesPathSetValue, &_sCPreferencesPathSetValueErr, frameworkHandle, "SCPreferencesPathSetValue", "10.1")
	registerFunc(&_sCPreferencesRemoveValue, &_sCPreferencesRemoveValueErr, frameworkHandle, "SCPreferencesRemoveValue", "10.1")
	registerFunc(&_sCPreferencesScheduleWithRunLoop, &_sCPreferencesScheduleWithRunLoopErr, frameworkHandle, "SCPreferencesScheduleWithRunLoop", "10.4")
	registerFunc(&_sCPreferencesSetCallback, &_sCPreferencesSetCallbackErr, frameworkHandle, "SCPreferencesSetCallback", "10.4")
	registerFunc(&_sCPreferencesSetComputerName, &_sCPreferencesSetComputerNameErr, frameworkHandle, "SCPreferencesSetComputerName", "10.1")
	registerFunc(&_sCPreferencesSetDispatchQueue, &_sCPreferencesSetDispatchQueueErr, frameworkHandle, "SCPreferencesSetDispatchQueue", "10.6")
	registerFunc(&_sCPreferencesSetLocalHostName, &_sCPreferencesSetLocalHostNameErr, frameworkHandle, "SCPreferencesSetLocalHostName", "10.2")
	registerFunc(&_sCPreferencesSetValue, &_sCPreferencesSetValueErr, frameworkHandle, "SCPreferencesSetValue", "10.1")
	registerFunc(&_sCPreferencesSynchronize, &_sCPreferencesSynchronizeErr, frameworkHandle, "SCPreferencesSynchronize", "10.4")
	registerFunc(&_sCPreferencesUnlock, &_sCPreferencesUnlockErr, frameworkHandle, "SCPreferencesUnlock", "10.1")
	registerFunc(&_sCPreferencesUnscheduleFromRunLoop, &_sCPreferencesUnscheduleFromRunLoopErr, frameworkHandle, "SCPreferencesUnscheduleFromRunLoop", "10.4")
	registerFunc(&_sCVLANInterfaceCopyAll, &_sCVLANInterfaceCopyAllErr, frameworkHandle, "SCVLANInterfaceCopyAll", "10.5")
	registerFunc(&_sCVLANInterfaceCopyAvailablePhysicalInterfaces, &_sCVLANInterfaceCopyAvailablePhysicalInterfacesErr, frameworkHandle, "SCVLANInterfaceCopyAvailablePhysicalInterfaces", "10.5")
	registerFunc(&_sCVLANInterfaceCreate, &_sCVLANInterfaceCreateErr, frameworkHandle, "SCVLANInterfaceCreate", "10.5")
	registerFunc(&_sCVLANInterfaceGetOptions, &_sCVLANInterfaceGetOptionsErr, frameworkHandle, "SCVLANInterfaceGetOptions", "10.5")
	registerFunc(&_sCVLANInterfaceGetPhysicalInterface, &_sCVLANInterfaceGetPhysicalInterfaceErr, frameworkHandle, "SCVLANInterfaceGetPhysicalInterface", "10.5")
	registerFunc(&_sCVLANInterfaceGetTag, &_sCVLANInterfaceGetTagErr, frameworkHandle, "SCVLANInterfaceGetTag", "10.5")
	registerFunc(&_sCVLANInterfaceRemove, &_sCVLANInterfaceRemoveErr, frameworkHandle, "SCVLANInterfaceRemove", "10.5")
	registerFunc(&_sCVLANInterfaceSetLocalizedDisplayName, &_sCVLANInterfaceSetLocalizedDisplayNameErr, frameworkHandle, "SCVLANInterfaceSetLocalizedDisplayName", "10.5")
	registerFunc(&_sCVLANInterfaceSetOptions, &_sCVLANInterfaceSetOptionsErr, frameworkHandle, "SCVLANInterfaceSetOptions", "10.5")
	registerFunc(&_sCVLANInterfaceSetPhysicalInterfaceAndTag, &_sCVLANInterfaceSetPhysicalInterfaceAndTagErr, frameworkHandle, "SCVLANInterfaceSetPhysicalInterfaceAndTag", "10.5")
}
