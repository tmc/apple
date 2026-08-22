// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("CoreWLAN: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreWLAN: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreWLAN: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreWLAN: register symbol %s: %v", name, r)
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

var _cWKeychainCopyEAPIdentity func(ssidData corefoundation.CFDataRef, identity *security.SecIdentityRef) int32
var _cWKeychainCopyEAPIdentityErr error

func tryCWKeychainCopyEAPIdentity(ssidData corefoundation.CFDataRef, identity *security.SecIdentityRef) (int32, error) {
	if _cWKeychainCopyEAPIdentity == nil {
		return 0, symbolCallError("CWKeychainCopyEAPIdentity", "10.7", _cWKeychainCopyEAPIdentityErr)
	}
	return _cWKeychainCopyEAPIdentity(ssidData, identity), nil
}

// CWKeychainCopyEAPIdentity finds and returns the identity stored for corresponding network with the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPIdentity
func CWKeychainCopyEAPIdentity(ssidData corefoundation.CFDataRef, identity *security.SecIdentityRef) int32 {
	result, callErr := tryCWKeychainCopyEAPIdentity(ssidData, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainCopyEAPIdentityList func(list *corefoundation.CFArrayRef) int32
var _cWKeychainCopyEAPIdentityListErr error

func tryCWKeychainCopyEAPIdentityList(list *corefoundation.CFArrayRef) (int32, error) {
	if _cWKeychainCopyEAPIdentityList == nil {
		return 0, symbolCallError("CWKeychainCopyEAPIdentityList", "10.7", _cWKeychainCopyEAPIdentityListErr)
	}
	return _cWKeychainCopyEAPIdentityList(list), nil
}

// CWKeychainCopyEAPIdentityList finds and returns the available identities stored in the keychain.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPIdentityList(_:)
func CWKeychainCopyEAPIdentityList(list *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCWKeychainCopyEAPIdentityList(list)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainCopyEAPUsernameAndPassword func(ssidData corefoundation.CFDataRef, username *corefoundation.CFStringRef, password *corefoundation.CFStringRef) int32
var _cWKeychainCopyEAPUsernameAndPasswordErr error

func tryCWKeychainCopyEAPUsernameAndPassword(ssidData corefoundation.CFDataRef, username *corefoundation.CFStringRef, password *corefoundation.CFStringRef) (int32, error) {
	if _cWKeychainCopyEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainCopyEAPUsernameAndPassword", "10.7", _cWKeychainCopyEAPUsernameAndPasswordErr)
	}
	return _cWKeychainCopyEAPUsernameAndPassword(ssidData, username, password), nil
}

// CWKeychainCopyEAPUsernameAndPassword finds and returns the username and password stored for corresponding network with the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPUsernameAndPassword
func CWKeychainCopyEAPUsernameAndPassword(ssidData corefoundation.CFDataRef, username *corefoundation.CFStringRef, password *corefoundation.CFStringRef) int32 {
	result, callErr := tryCWKeychainCopyEAPUsernameAndPassword(ssidData, username, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainCopyPassword func(ssidData corefoundation.CFDataRef, password *corefoundation.CFStringRef) int32
var _cWKeychainCopyPasswordErr error

func tryCWKeychainCopyPassword(ssidData corefoundation.CFDataRef, password *corefoundation.CFStringRef) (int32, error) {
	if _cWKeychainCopyPassword == nil {
		return 0, symbolCallError("CWKeychainCopyPassword", "10.7", _cWKeychainCopyPasswordErr)
	}
	return _cWKeychainCopyPassword(ssidData, password), nil
}

// CWKeychainCopyPassword finds and returns the keychain password stored for the corresponding network with the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyPassword
func CWKeychainCopyPassword(ssidData corefoundation.CFDataRef, password *corefoundation.CFStringRef) int32 {
	result, callErr := tryCWKeychainCopyPassword(ssidData, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainCopyWiFiEAPIdentity func(domain CWKeychainDomain, ssid *foundation.NSData, identity *security.SecIdentityRef) int32
var _cWKeychainCopyWiFiEAPIdentityErr error

func tryCWKeychainCopyWiFiEAPIdentity(domain CWKeychainDomain, ssid *foundation.NSData, identity *security.SecIdentityRef) (int32, error) {
	if _cWKeychainCopyWiFiEAPIdentity == nil {
		return 0, symbolCallError("CWKeychainCopyWiFiEAPIdentity", "10.9", _cWKeychainCopyWiFiEAPIdentityErr)
	}
	return _cWKeychainCopyWiFiEAPIdentity(domain, ssid, identity), nil
}

// CWKeychainCopyWiFiEAPIdentity finds and returns the identity stored for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyWiFiEAPIdentity(_:_:_:)
func CWKeychainCopyWiFiEAPIdentity(domain CWKeychainDomain, ssid *foundation.NSData, identity *security.SecIdentityRef) int32 {
	result, callErr := tryCWKeychainCopyWiFiEAPIdentity(domain, ssid, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainDeleteEAPUsernameAndPassword func(ssidData corefoundation.CFDataRef) int32
var _cWKeychainDeleteEAPUsernameAndPasswordErr error

func tryCWKeychainDeleteEAPUsernameAndPassword(ssidData corefoundation.CFDataRef) (int32, error) {
	if _cWKeychainDeleteEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainDeleteEAPUsernameAndPassword", "10.7", _cWKeychainDeleteEAPUsernameAndPasswordErr)
	}
	return _cWKeychainDeleteEAPUsernameAndPassword(ssidData), nil
}

// CWKeychainDeleteEAPUsernameAndPassword deletes the keychain item containing the 802.1X username and password for the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteEAPUsernameAndPassword
func CWKeychainDeleteEAPUsernameAndPassword(ssidData corefoundation.CFDataRef) int32 {
	result, callErr := tryCWKeychainDeleteEAPUsernameAndPassword(ssidData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainDeletePassword func(ssidData corefoundation.CFDataRef) int32
var _cWKeychainDeletePasswordErr error

func tryCWKeychainDeletePassword(ssidData corefoundation.CFDataRef) (int32, error) {
	if _cWKeychainDeletePassword == nil {
		return 0, symbolCallError("CWKeychainDeletePassword", "10.7", _cWKeychainDeletePasswordErr)
	}
	return _cWKeychainDeletePassword(ssidData), nil
}

// CWKeychainDeletePassword deletes the network password for the specified SSID from the default keychain.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeletePassword
func CWKeychainDeletePassword(ssidData corefoundation.CFDataRef) int32 {
	result, callErr := tryCWKeychainDeletePassword(ssidData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainDeleteWiFiEAPUsernameAndPassword func(domain CWKeychainDomain, ssid *foundation.NSData) int32
var _cWKeychainDeleteWiFiEAPUsernameAndPasswordErr error

func tryCWKeychainDeleteWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData) (int32, error) {
	if _cWKeychainDeleteWiFiEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainDeleteWiFiEAPUsernameAndPassword", "10.9", _cWKeychainDeleteWiFiEAPUsernameAndPasswordErr)
	}
	return _cWKeychainDeleteWiFiEAPUsernameAndPassword(domain, ssid), nil
}

// CWKeychainDeleteWiFiEAPUsernameAndPassword deletes the 802.1X username and password for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteWiFiEAPUsernameAndPassword(_:_:)
func CWKeychainDeleteWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData) int32 {
	result, callErr := tryCWKeychainDeleteWiFiEAPUsernameAndPassword(domain, ssid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainDeleteWiFiPassword func(domain CWKeychainDomain, ssid *foundation.NSData) int32
var _cWKeychainDeleteWiFiPasswordErr error

func tryCWKeychainDeleteWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData) (int32, error) {
	if _cWKeychainDeleteWiFiPassword == nil {
		return 0, symbolCallError("CWKeychainDeleteWiFiPassword", "10.9", _cWKeychainDeleteWiFiPasswordErr)
	}
	return _cWKeychainDeleteWiFiPassword(domain, ssid), nil
}

// CWKeychainDeleteWiFiPassword deletes the password for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteWiFiPassword(_:_:)
func CWKeychainDeleteWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData) int32 {
	result, callErr := tryCWKeychainDeleteWiFiPassword(domain, ssid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainFindWiFiEAPUsernameAndPassword func(domain CWKeychainDomain, ssid *foundation.NSData, username foundation.NSString, password foundation.NSString) int32
var _cWKeychainFindWiFiEAPUsernameAndPasswordErr error

func tryCWKeychainFindWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData, username foundation.NSString, password foundation.NSString) (int32, error) {
	if _cWKeychainFindWiFiEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainFindWiFiEAPUsernameAndPassword", "10.9", _cWKeychainFindWiFiEAPUsernameAndPasswordErr)
	}
	return _cWKeychainFindWiFiEAPUsernameAndPassword(domain, ssid, username, password), nil
}

// CWKeychainFindWiFiEAPUsernameAndPassword finds and returns the 802.1X username and password stored for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainFindWiFiEAPUsernameAndPassword(_:_:_:_:)
func CWKeychainFindWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData, username foundation.NSString, password foundation.NSString) int32 {
	result, callErr := tryCWKeychainFindWiFiEAPUsernameAndPassword(domain, ssid, username, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainFindWiFiPassword func(domain CWKeychainDomain, ssid *foundation.NSData, password foundation.NSString) int32
var _cWKeychainFindWiFiPasswordErr error

func tryCWKeychainFindWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData, password foundation.NSString) (int32, error) {
	if _cWKeychainFindWiFiPassword == nil {
		return 0, symbolCallError("CWKeychainFindWiFiPassword", "10.9", _cWKeychainFindWiFiPasswordErr)
	}
	return _cWKeychainFindWiFiPassword(domain, ssid, password), nil
}

// CWKeychainFindWiFiPassword finds and returns, by reference, the password for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainFindWiFiPassword(_:_:_:)
func CWKeychainFindWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData, password foundation.NSString) int32 {
	result, callErr := tryCWKeychainFindWiFiPassword(domain, ssid, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetEAPIdentity func(ssidData corefoundation.CFDataRef, identity security.SecIdentityRef) int32
var _cWKeychainSetEAPIdentityErr error

func tryCWKeychainSetEAPIdentity(ssidData corefoundation.CFDataRef, identity security.SecIdentityRef) (int32, error) {
	if _cWKeychainSetEAPIdentity == nil {
		return 0, symbolCallError("CWKeychainSetEAPIdentity", "10.7", _cWKeychainSetEAPIdentityErr)
	}
	return _cWKeychainSetEAPIdentity(ssidData, identity), nil
}

// CWKeychainSetEAPIdentity associates an exisiting identity item to the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetEAPIdentity
func CWKeychainSetEAPIdentity(ssidData corefoundation.CFDataRef, identity security.SecIdentityRef) int32 {
	result, callErr := tryCWKeychainSetEAPIdentity(ssidData, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetEAPUsernameAndPassword func(ssidData corefoundation.CFDataRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef) int32
var _cWKeychainSetEAPUsernameAndPasswordErr error

func tryCWKeychainSetEAPUsernameAndPassword(ssidData corefoundation.CFDataRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef) (int32, error) {
	if _cWKeychainSetEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainSetEAPUsernameAndPassword", "10.7", _cWKeychainSetEAPUsernameAndPasswordErr)
	}
	return _cWKeychainSetEAPUsernameAndPassword(ssidData, username, password), nil
}

// CWKeychainSetEAPUsernameAndPassword sets the keychain item containing the 802.1X username and password for the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetEAPUsernameAndPassword
func CWKeychainSetEAPUsernameAndPassword(ssidData corefoundation.CFDataRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef) int32 {
	result, callErr := tryCWKeychainSetEAPUsernameAndPassword(ssidData, username, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetPassword func(ssidData corefoundation.CFDataRef, password corefoundation.CFStringRef) int32
var _cWKeychainSetPasswordErr error

func tryCWKeychainSetPassword(ssidData corefoundation.CFDataRef, password corefoundation.CFStringRef) (int32, error) {
	if _cWKeychainSetPassword == nil {
		return 0, symbolCallError("CWKeychainSetPassword", "10.7", _cWKeychainSetPasswordErr)
	}
	return _cWKeychainSetPassword(ssidData, password), nil
}

// CWKeychainSetPassword sets the network keychain password for the specified SSID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetPassword
func CWKeychainSetPassword(ssidData corefoundation.CFDataRef, password corefoundation.CFStringRef) int32 {
	result, callErr := tryCWKeychainSetPassword(ssidData, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetWiFiEAPIdentity func(domain CWKeychainDomain, ssid *foundation.NSData, identity security.SecIdentityRef) int32
var _cWKeychainSetWiFiEAPIdentityErr error

func tryCWKeychainSetWiFiEAPIdentity(domain CWKeychainDomain, ssid *foundation.NSData, identity security.SecIdentityRef) (int32, error) {
	if _cWKeychainSetWiFiEAPIdentity == nil {
		return 0, symbolCallError("CWKeychainSetWiFiEAPIdentity", "10.9", _cWKeychainSetWiFiEAPIdentityErr)
	}
	return _cWKeychainSetWiFiEAPIdentity(domain, ssid, identity), nil
}

// CWKeychainSetWiFiEAPIdentity associates an identity to the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiEAPIdentity(_:_:_:)
func CWKeychainSetWiFiEAPIdentity(domain CWKeychainDomain, ssid *foundation.NSData, identity security.SecIdentityRef) int32 {
	result, callErr := tryCWKeychainSetWiFiEAPIdentity(domain, ssid, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetWiFiEAPUsernameAndPassword func(domain CWKeychainDomain, ssid *foundation.NSData, username *foundation.NSString, password *foundation.NSString) int32
var _cWKeychainSetWiFiEAPUsernameAndPasswordErr error

func tryCWKeychainSetWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData, username *foundation.NSString, password *foundation.NSString) (int32, error) {
	if _cWKeychainSetWiFiEAPUsernameAndPassword == nil {
		return 0, symbolCallError("CWKeychainSetWiFiEAPUsernameAndPassword", "10.9", _cWKeychainSetWiFiEAPUsernameAndPasswordErr)
	}
	return _cWKeychainSetWiFiEAPUsernameAndPassword(domain, ssid, username, password), nil
}

// CWKeychainSetWiFiEAPUsernameAndPassword sets the 802.1X username and password for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiEAPUsernameAndPassword(_:_:_:_:)
func CWKeychainSetWiFiEAPUsernameAndPassword(domain CWKeychainDomain, ssid *foundation.NSData, username *foundation.NSString, password *foundation.NSString) int32 {
	result, callErr := tryCWKeychainSetWiFiEAPUsernameAndPassword(domain, ssid, username, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWKeychainSetWiFiPassword func(domain CWKeychainDomain, ssid *foundation.NSData, password *foundation.NSString) int32
var _cWKeychainSetWiFiPasswordErr error

func tryCWKeychainSetWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData, password *foundation.NSString) (int32, error) {
	if _cWKeychainSetWiFiPassword == nil {
		return 0, symbolCallError("CWKeychainSetWiFiPassword", "10.9", _cWKeychainSetWiFiPasswordErr)
	}
	return _cWKeychainSetWiFiPassword(domain, ssid, password), nil
}

// CWKeychainSetWiFiPassword sets the Wi-Fi network keychain password for the SSID and keychain domain you specify.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiPassword(_:_:_:)
func CWKeychainSetWiFiPassword(domain CWKeychainDomain, ssid *foundation.NSData, password *foundation.NSString) int32 {
	result, callErr := tryCWKeychainSetWiFiPassword(domain, ssid, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cWMergeNetworks func(networks uintptr) uintptr
var _cWMergeNetworksErr error

func tryCWMergeNetworks(networks uintptr) (uintptr, error) {
	if _cWMergeNetworks == nil {
		return 0, symbolCallError("CWMergeNetworks", "10.7", _cWMergeNetworksErr)
	}
	return _cWMergeNetworks(networks), nil
}

// CWMergeNetworks merges the specified set of CWNetwork objects.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWMergeNetworks(_:)
func CWMergeNetworks(networks uintptr) uintptr {
	result, callErr := tryCWMergeNetworks(networks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cWKeychainCopyEAPIdentity, &_cWKeychainCopyEAPIdentityErr, frameworkHandle, "CWKeychainCopyEAPIdentity", "10.7")
	registerFunc(&_cWKeychainCopyEAPIdentityList, &_cWKeychainCopyEAPIdentityListErr, frameworkHandle, "CWKeychainCopyEAPIdentityList", "10.7")
	registerFunc(&_cWKeychainCopyEAPUsernameAndPassword, &_cWKeychainCopyEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainCopyEAPUsernameAndPassword", "10.7")
	registerFunc(&_cWKeychainCopyPassword, &_cWKeychainCopyPasswordErr, frameworkHandle, "CWKeychainCopyPassword", "10.7")
	registerFunc(&_cWKeychainCopyWiFiEAPIdentity, &_cWKeychainCopyWiFiEAPIdentityErr, frameworkHandle, "CWKeychainCopyWiFiEAPIdentity", "10.9")
	registerFunc(&_cWKeychainDeleteEAPUsernameAndPassword, &_cWKeychainDeleteEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainDeleteEAPUsernameAndPassword", "10.7")
	registerFunc(&_cWKeychainDeletePassword, &_cWKeychainDeletePasswordErr, frameworkHandle, "CWKeychainDeletePassword", "10.7")
	registerFunc(&_cWKeychainDeleteWiFiEAPUsernameAndPassword, &_cWKeychainDeleteWiFiEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainDeleteWiFiEAPUsernameAndPassword", "10.9")
	registerFunc(&_cWKeychainDeleteWiFiPassword, &_cWKeychainDeleteWiFiPasswordErr, frameworkHandle, "CWKeychainDeleteWiFiPassword", "10.9")
	registerFunc(&_cWKeychainFindWiFiEAPUsernameAndPassword, &_cWKeychainFindWiFiEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainFindWiFiEAPUsernameAndPassword", "10.9")
	registerFunc(&_cWKeychainFindWiFiPassword, &_cWKeychainFindWiFiPasswordErr, frameworkHandle, "CWKeychainFindWiFiPassword", "10.9")
	registerFunc(&_cWKeychainSetEAPIdentity, &_cWKeychainSetEAPIdentityErr, frameworkHandle, "CWKeychainSetEAPIdentity", "10.7")
	registerFunc(&_cWKeychainSetEAPUsernameAndPassword, &_cWKeychainSetEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainSetEAPUsernameAndPassword", "10.7")
	registerFunc(&_cWKeychainSetPassword, &_cWKeychainSetPasswordErr, frameworkHandle, "CWKeychainSetPassword", "10.7")
	registerFunc(&_cWKeychainSetWiFiEAPIdentity, &_cWKeychainSetWiFiEAPIdentityErr, frameworkHandle, "CWKeychainSetWiFiEAPIdentity", "10.9")
	registerFunc(&_cWKeychainSetWiFiEAPUsernameAndPassword, &_cWKeychainSetWiFiEAPUsernameAndPasswordErr, frameworkHandle, "CWKeychainSetWiFiEAPUsernameAndPassword", "10.9")
	registerFunc(&_cWKeychainSetWiFiPassword, &_cWKeychainSetWiFiPasswordErr, frameworkHandle, "CWKeychainSetWiFiPassword", "10.9")
	registerFunc(&_cWMergeNetworks, &_cWMergeNetworksErr, frameworkHandle, "CWMergeNetworks", "10.7")
}
