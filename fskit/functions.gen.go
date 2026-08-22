// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("FSKit: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("FSKit: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("FSKit: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("FSKit: register symbol %s: %v", name, r)
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

var _fs_errorForCocoaError func(errorCode int32) *foundation.NSError
var _fs_errorForCocoaErrorErr error

func tryFs_errorForCocoaError(errorCode int32) (*foundation.NSError, error) {
	if _fs_errorForCocoaError == nil {
		return nil, symbolCallError("fs_errorForCocoaError", "15.4", _fs_errorForCocoaErrorErr)
	}
	return _fs_errorForCocoaError(errorCode), nil
}

// Fs_errorForCocoaError creates an error object for the given Cocoa error code.
//
// See: https://developer.apple.com/documentation/FSKit/fs_errorForCocoaError(_:)
func Fs_errorForCocoaError(errorCode int32) *foundation.NSError {
	result, callErr := tryFs_errorForCocoaError(errorCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fs_errorForMachError func(errorCode int32) *foundation.NSError
var _fs_errorForMachErrorErr error

func tryFs_errorForMachError(errorCode int32) (*foundation.NSError, error) {
	if _fs_errorForMachError == nil {
		return nil, symbolCallError("fs_errorForMachError", "15.4", _fs_errorForMachErrorErr)
	}
	return _fs_errorForMachError(errorCode), nil
}

// Fs_errorForMachError creates an error object for the given Mach error code.
//
// See: https://developer.apple.com/documentation/FSKit/fs_errorForMachError(_:)
func Fs_errorForMachError(errorCode int32) *foundation.NSError {
	result, callErr := tryFs_errorForMachError(errorCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fs_errorForPOSIXError func(arg0 int32) *foundation.NSError
var _fs_errorForPOSIXErrorErr error

func tryFs_errorForPOSIXError(arg0 int32) (*foundation.NSError, error) {
	if _fs_errorForPOSIXError == nil {
		return nil, symbolCallError("fs_errorForPOSIXError", "15.4", _fs_errorForPOSIXErrorErr)
	}
	return _fs_errorForPOSIXError(arg0), nil
}

// Fs_errorForPOSIXError creates an error object for the given POSIX error code.
//
// See: https://developer.apple.com/documentation/FSKit/fs_errorForPOSIXError(_:)
func Fs_errorForPOSIXError(arg0 int32) *foundation.NSError {
	result, callErr := tryFs_errorForPOSIXError(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_fs_errorForCocoaError, &_fs_errorForCocoaErrorErr, frameworkHandle, "fs_errorForCocoaError", "15.4")
	registerFunc(&_fs_errorForMachError, &_fs_errorForMachErrorErr, frameworkHandle, "fs_errorForMachError", "15.4")
	registerFunc(&_fs_errorForPOSIXError, &_fs_errorForPOSIXErrorErr, frameworkHandle, "fs_errorForPOSIXError", "15.4")
}
