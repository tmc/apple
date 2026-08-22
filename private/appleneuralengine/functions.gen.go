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

var _aNEGetValidateNetworkSupportedVersion func() uint32
var _aNEGetValidateNetworkSupportedVersionErr error

func tryANEGetValidateNetworkSupportedVersion() (uint32, error) {
	if _aNEGetValidateNetworkSupportedVersion == nil {
		return 0, symbolCallError("ANEGetValidateNetworkSupportedVersion", "", _aNEGetValidateNetworkSupportedVersionErr)
	}
	return _aNEGetValidateNetworkSupportedVersion(), nil
}

// ANEGetValidateNetworkSupportedVersion.
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

// ANEValidateNetworkCreate.
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

// ANEValidateNetworkCreateVMHost.
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
