// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"fmt"

	"github.com/ebitengine/purego"
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
		return fmt.Sprintf("appleneuralengine: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("appleneuralengine: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("appleneuralengine: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("appleneuralengine: register symbol %s: %v", name, r)
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

// SymbolAddress returns the address of name in appleneuralengine, whether or not
// this package generated a binding for it.
//
// What is generated is bounded by what is documented, and for a private
// framework that is whatever a manifest happened to enumerate. The dylib
// usually exports far more. Without this, a symbol nobody wrote down is
// unreachable from a package that has already loaded the image holding it, and
// the generated surface becomes a ceiling instead of a floor.
//
// The lookup is scoped to this framework's handle, not RTLD_DEFAULT, so a
// symbol some other loaded image exports is not reported as this one's.
func SymbolAddress(name string) (uintptr, error) {
	if frameworkHandle == 0 {
		return 0, fmt.Errorf("appleneuralengine: symbol %s unavailable because the framework could not be loaded", name)
	}
	sym, err := purego.Dlsym(frameworkHandle, name)
	if err != nil || sym == 0 {
		return 0, missingSymbolError(name, "", err)
	}
	return sym, nil
}

// BindFunc binds the appleneuralengine symbol name into fptr, which must be a
// pointer to a func variable.
//
// The caller supplies the signature, and nothing checks it. A dylib records no
// argument count or types for a C symbol, so a wrong signature here is not a
// type error: it is the wrong number of machine words moved on a live stack,
// and the failure surfaces somewhere else entirely. Prefer a generated binding,
// whose signature carries recorded evidence, and reach for this only for a
// symbol that has none.
// purego.RegisterFunc panics on a signature it cannot lower; that is recovered
// and returned, because an escape hatch that takes down the process on a
// mistyped experiment is not one anybody can experiment with.
func BindFunc(fptr any, name string) (err error) {
	sym, err := SymbolAddress(name)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("appleneuralengine: bind symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	return nil
}

var _aNEGetValidateNetworkSupportedVersion func() uint32
var _aNEGetValidateNetworkSupportedVersionErr error

func tryANEGetValidateNetworkSupportedVersion() (uint32, error) {
	if _aNEGetValidateNetworkSupportedVersion == nil {
		return 0, symbolCallError("ANEGetValidateNetworkSupportedVersion", "", _aNEGetValidateNetworkSupportedVersionErr)
	}
	return _aNEGetValidateNetworkSupportedVersion(), nil
}

// ANEGetValidateNetworkSupportedVersion signature evidence: disassembly, macOS 26.6 arm64e: reads no argument register before its first call, consistent with taking no parameters.
func ANEGetValidateNetworkSupportedVersion() (uint32, error) {
	return tryANEGetValidateNetworkSupportedVersion()
}

var _aNEValidateNetworkCreate func(arg0 objectivec.Object, arg1 objectivec.Object) int32
var _aNEValidateNetworkCreateErr error

func tryANEValidateNetworkCreate(arg0 objectivec.Object, arg1 objectivec.Object) (int32, error) {
	if _aNEValidateNetworkCreate == nil {
		return 0, symbolCallError("ANEValidateNetworkCreate", "", _aNEValidateNetworkCreateErr)
	}
	return _aNEValidateNetworkCreate(arg0, arg1), nil
}

// ANEValidateNetworkCreate signature evidence: disassembly, macOS 26.6 arm64e, binary extracted from the dyld shared cache with ipsw: arity from both the prologue and the call site in ANEValidateNetworkCreateVMHost; both parameters shown to be Objective-C objects by objc_release and by objc_msgSend receiver use.
func ANEValidateNetworkCreate(arg0 objectivec.Object, arg1 objectivec.Object) (int32, error) {
	return tryANEValidateNetworkCreate(arg0, arg1)
}

var _aNEValidateNetworkCreateVMHost func(arg0 objectivec.Object, arg1 objectivec.Object, arg2 objectivec.Object, arg3 objectivec.Object, arg4 uintptr) int32
var _aNEValidateNetworkCreateVMHostErr error

func tryANEValidateNetworkCreateVMHost(arg0 objectivec.Object, arg1 objectivec.Object, arg2 objectivec.Object, arg3 objectivec.Object, arg4 uintptr) (int32, error) {
	if _aNEValidateNetworkCreateVMHost == nil {
		return 0, symbolCallError("ANEValidateNetworkCreateVMHost", "", _aNEValidateNetworkCreateVMHostErr)
	}
	return _aNEValidateNetworkCreateVMHost(arg0, arg1, arg2, arg3, arg4), nil
}

// ANEValidateNetworkCreateVMHost signature evidence: disassembly, macOS 26.6 arm64e, binary extracted from the dyld shared cache with ipsw: arity from the prologue; four parameters classified by objc_msgSend receiver use and objc_retain; arg4 shown to be a pointer by an authenticated load, kind unmeasured.
func ANEValidateNetworkCreateVMHost(arg0 objectivec.Object, arg1 objectivec.Object, arg2 objectivec.Object, arg3 objectivec.Object, arg4 uintptr) (int32, error) {
	return tryANEValidateNetworkCreateVMHost(arg0, arg1, arg2, arg3, arg4)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_aNEGetValidateNetworkSupportedVersion, &_aNEGetValidateNetworkSupportedVersionErr, frameworkHandle, "ANEGetValidateNetworkSupportedVersion", "")
	registerFunc(&_aNEValidateNetworkCreate, &_aNEValidateNetworkCreateErr, frameworkHandle, "ANEValidateNetworkCreate", "")
	registerFunc(&_aNEValidateNetworkCreateVMHost, &_aNEValidateNetworkCreateVMHostErr, frameworkHandle, "ANEValidateNetworkCreateVMHost", "")
}
