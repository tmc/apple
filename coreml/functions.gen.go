// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
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
		return fmt.Sprintf("CoreML: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreML: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreML: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreML: register symbol %s: %v", name, r)
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

var _mLAllComputeDevices func() []unsafe.Pointer
var _mLAllComputeDevicesErr error

func tryMLAllComputeDevices() ([]MLComputeDeviceProtocolObject, error) {
	if _mLAllComputeDevices == nil {
		return nil, symbolCallError("MLAllComputeDevices", "14.0", _mLAllComputeDevicesErr)
	}
	ptrs := _mLAllComputeDevices()
	result := make([]MLComputeDeviceProtocolObject, len(ptrs))
	for i, p := range ptrs {
		result[i] = MLComputeDeviceProtocolObjectFromID(objc.IDFrom(p))
	}
	return result, nil
}

// MLAllComputeDevices returns an array that contains all of the compute devices that are accessible.
//
// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDevices
func MLAllComputeDevices() []MLComputeDeviceProtocolObject {
	result, callErr := tryMLAllComputeDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mLAllComputeDevices, &_mLAllComputeDevicesErr, frameworkHandle, "MLAllComputeDevices", "14.0")
}
