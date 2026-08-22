// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("ParavirtualizedGraphics: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ParavirtualizedGraphics: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("ParavirtualizedGraphics: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("ParavirtualizedGraphics: register symbol %s: %v", name, r)
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

var _pGCopyOptionROMURL func() *foundation.NSURL
var _pGCopyOptionROMURLErr error

func tryPGCopyOptionROMURL() (*foundation.NSURL, error) {
	if _pGCopyOptionROMURL == nil {
		return nil, symbolCallError("PGCopyOptionROMURL", "11.0", _pGCopyOptionROMURLErr)
	}
	return _pGCopyOptionROMURL(), nil
}

// PGCopyOptionROMURL copies the URL of the ROM image to use on the guest graphics device.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGCopyOptionROMURL()
func PGCopyOptionROMURL() *foundation.NSURL {
	result, callErr := tryPGCopyOptionROMURL()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pGCreateDeviceWithDescriptor func(descriptor *PGDeviceDescriptor) unsafe.Pointer
var _pGCreateDeviceWithDescriptorErr error

func tryPGCreateDeviceWithDescriptor(descriptor *PGDeviceDescriptor) (PGDeviceObject, error) {
	if _pGCreateDeviceWithDescriptor == nil {
		return *new(PGDeviceObject), symbolCallError("PGCreateDeviceWithDescriptor", "15.2", _pGCreateDeviceWithDescriptorErr)
	}
	rv := _pGCreateDeviceWithDescriptor(descriptor)
	return PGDeviceObjectFromID(objc.IDFrom(rv)), nil
}

// PGCreateDeviceWithDescriptor.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGCreateDeviceWithDescriptor(_:)
func PGCreateDeviceWithDescriptor(descriptor *PGDeviceDescriptor) PGDeviceObject {
	result, callErr := tryPGCreateDeviceWithDescriptor(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pGMaxDisplayPortCount func() uint32
var _pGMaxDisplayPortCountErr error

func tryPGMaxDisplayPortCount() (uint32, error) {
	if _pGMaxDisplayPortCount == nil {
		return 0, symbolCallError("PGMaxDisplayPortCount", "13.0", _pGMaxDisplayPortCountErr)
	}
	return _pGMaxDisplayPortCount(), nil
}

// PGMaxDisplayPortCount.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGMaxDisplayPortCount()
func PGMaxDisplayPortCount() uint32 {
	result, callErr := tryPGMaxDisplayPortCount()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pGNewDeviceWithDescriptor func(descriptor *PGDeviceDescriptor) unsafe.Pointer
var _pGNewDeviceWithDescriptorErr error

func tryPGNewDeviceWithDescriptor(descriptor *PGDeviceDescriptor) (PGDeviceObject, error) {
	if _pGNewDeviceWithDescriptor == nil {
		return *new(PGDeviceObject), symbolCallError("PGNewDeviceWithDescriptor", "11.0", _pGNewDeviceWithDescriptorErr)
	}
	rv := _pGNewDeviceWithDescriptor(descriptor)
	return PGDeviceObjectFromID(objc.IDFrom(rv)), nil
}

// PGNewDeviceWithDescriptor creates a new paravirtualized graphics device.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGNewDeviceWithDescriptor(_:)
func PGNewDeviceWithDescriptor(descriptor *PGDeviceDescriptor) PGDeviceObject {
	result, callErr := tryPGNewDeviceWithDescriptor(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_pGCopyOptionROMURL, &_pGCopyOptionROMURLErr, frameworkHandle, "PGCopyOptionROMURL", "11.0")
	registerFunc(&_pGCreateDeviceWithDescriptor, &_pGCreateDeviceWithDescriptorErr, frameworkHandle, "PGCreateDeviceWithDescriptor", "15.2")
	registerFunc(&_pGMaxDisplayPortCount, &_pGMaxDisplayPortCountErr, frameworkHandle, "PGMaxDisplayPortCount", "13.0")
	registerFunc(&_pGNewDeviceWithDescriptor, &_pGNewDeviceWithDescriptorErr, frameworkHandle, "PGNewDeviceWithDescriptor", "11.0")
}
