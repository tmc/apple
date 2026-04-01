// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("IOSurface: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("IOSurface: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("IOSurface: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("IOSurface: register symbol %s: %v", name, r)
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

var _iOSurfaceAlignProperty func(property corefoundation.CFStringRef, value uintptr) uintptr
var _iOSurfaceAlignPropertyErr error

func tryIOSurfaceAlignProperty(property corefoundation.CFStringRef, value uintptr) (uintptr, error) {
	if _iOSurfaceAlignProperty == nil {
		return 0, symbolCallError("IOSurfaceAlignProperty", "10.6", _iOSurfaceAlignPropertyErr)
	}
	return _iOSurfaceAlignProperty(property, value), nil
}

// IOSurfaceAlignProperty returns the smallest aligned value greater than or equal to the specified value.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceAlignProperty(_:_:)
func IOSurfaceAlignProperty(property corefoundation.CFStringRef, value uintptr) uintptr {
	result, callErr := tryIOSurfaceAlignProperty(property, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceAllowsPixelSizeCasting func(buffer IOSurfaceRef) bool
var _iOSurfaceAllowsPixelSizeCastingErr error

func tryIOSurfaceAllowsPixelSizeCasting(buffer IOSurfaceRef) (bool, error) {
	if _iOSurfaceAllowsPixelSizeCasting == nil {
		return false, symbolCallError("IOSurfaceAllowsPixelSizeCasting", "10.12", _iOSurfaceAllowsPixelSizeCastingErr)
	}
	return _iOSurfaceAllowsPixelSizeCasting(buffer), nil
}

// IOSurfaceAllowsPixelSizeCasting.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceAllowsPixelSizeCasting(_:)
func IOSurfaceAllowsPixelSizeCasting(buffer IOSurfaceRef) bool {
	result, callErr := tryIOSurfaceAllowsPixelSizeCasting(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceCopyAllValues func(buffer IOSurfaceRef) corefoundation.CFDictionaryRef
var _iOSurfaceCopyAllValuesErr error

func tryIOSurfaceCopyAllValues(buffer IOSurfaceRef) (corefoundation.CFDictionaryRef, error) {
	if _iOSurfaceCopyAllValues == nil {
		return 0, symbolCallError("IOSurfaceCopyAllValues", "10.6", _iOSurfaceCopyAllValuesErr)
	}
	return _iOSurfaceCopyAllValues(buffer), nil
}

// IOSurfaceCopyAllValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyAllValues(_:)
func IOSurfaceCopyAllValues(buffer IOSurfaceRef) corefoundation.CFDictionaryRef {
	result, callErr := tryIOSurfaceCopyAllValues(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceCopyValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOSurfaceCopyValueErr error

func tryIOSurfaceCopyValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOSurfaceCopyValue == nil {
		return 0, symbolCallError("IOSurfaceCopyValue", "10.6", _iOSurfaceCopyValueErr)
	}
	return _iOSurfaceCopyValue(buffer, key), nil
}

// IOSurfaceCopyValue retrieves a value from the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyValue(_:_:)
func IOSurfaceCopyValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOSurfaceCopyValue(buffer, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceCreate func(properties corefoundation.CFDictionaryRef) IOSurfaceRef
var _iOSurfaceCreateErr error

func tryIOSurfaceCreate(properties corefoundation.CFDictionaryRef) (IOSurfaceRef, error) {
	if _iOSurfaceCreate == nil {
		return 0, symbolCallError("IOSurfaceCreate", "10.6", _iOSurfaceCreateErr)
	}
	return _iOSurfaceCreate(properties), nil
}

// IOSurfaceCreate creates a brand new IOSurface object
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreate(_:)
func IOSurfaceCreate(properties corefoundation.CFDictionaryRef) IOSurfaceRef {
	result, callErr := tryIOSurfaceCreate(properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceCreateMachPort func(buffer IOSurfaceRef) uint32
var _iOSurfaceCreateMachPortErr error

func tryIOSurfaceCreateMachPort(buffer IOSurfaceRef) (uint32, error) {
	if _iOSurfaceCreateMachPort == nil {
		return 0, symbolCallError("IOSurfaceCreateMachPort", "10.6", _iOSurfaceCreateMachPortErr)
	}
	return _iOSurfaceCreateMachPort(buffer), nil
}

// IOSurfaceCreateMachPort returns a mach_port_t that holds a reference to the IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateMachPort(_:)
func IOSurfaceCreateMachPort(buffer IOSurfaceRef) uint32 {
	result, callErr := tryIOSurfaceCreateMachPort(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceCreateXPCObject func(aSurface IOSurfaceRef) unsafe.Pointer
var _iOSurfaceCreateXPCObjectErr error

func tryIOSurfaceCreateXPCObject(aSurface IOSurfaceRef) (unsafe.Pointer, error) {
	if _iOSurfaceCreateXPCObject == nil {
		return nil, symbolCallError("IOSurfaceCreateXPCObject", "10.7", _iOSurfaceCreateXPCObjectErr)
	}
	return _iOSurfaceCreateXPCObject(aSurface), nil
}

// IOSurfaceCreateXPCObject returns an xpc_object_t that holds a reference to the IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateXPCObject(_:)
func IOSurfaceCreateXPCObject(aSurface IOSurfaceRef) unsafe.Pointer {
	result, callErr := tryIOSurfaceCreateXPCObject(aSurface)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceDecrementUseCount func(buffer IOSurfaceRef)
var _iOSurfaceDecrementUseCountErr error

func tryIOSurfaceDecrementUseCount(buffer IOSurfaceRef) error {
	if _iOSurfaceDecrementUseCount == nil {
		return symbolCallError("IOSurfaceDecrementUseCount", "10.6", _iOSurfaceDecrementUseCountErr)
	}
	_iOSurfaceDecrementUseCount(buffer)
	return nil
}

// IOSurfaceDecrementUseCount decrements the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceDecrementUseCount(_:)
func IOSurfaceDecrementUseCount(buffer IOSurfaceRef) {
	if callErr := tryIOSurfaceDecrementUseCount(buffer); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceGetAllocSize func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetAllocSizeErr error

func tryIOSurfaceGetAllocSize(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetAllocSize == nil {
		return 0, symbolCallError("IOSurfaceGetAllocSize", "10.6", _iOSurfaceGetAllocSizeErr)
	}
	return _iOSurfaceGetAllocSize(buffer), nil
}

// IOSurfaceGetAllocSize returns the total allocation size of the buffer including all planes.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetAllocSize(_:)
func IOSurfaceGetAllocSize(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetAllocSize(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBaseAddress func(buffer IOSurfaceRef) unsafe.Pointer
var _iOSurfaceGetBaseAddressErr error

func tryIOSurfaceGetBaseAddress(buffer IOSurfaceRef) (unsafe.Pointer, error) {
	if _iOSurfaceGetBaseAddress == nil {
		return nil, symbolCallError("IOSurfaceGetBaseAddress", "10.6", _iOSurfaceGetBaseAddressErr)
	}
	return _iOSurfaceGetBaseAddress(buffer), nil
}

// IOSurfaceGetBaseAddress returns the address of the first byte of data in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddress(_:)
func IOSurfaceGetBaseAddress(buffer IOSurfaceRef) unsafe.Pointer {
	result, callErr := tryIOSurfaceGetBaseAddress(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBaseAddressOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) unsafe.Pointer
var _iOSurfaceGetBaseAddressOfPlaneErr error

func tryIOSurfaceGetBaseAddressOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (unsafe.Pointer, error) {
	if _iOSurfaceGetBaseAddressOfPlane == nil {
		return nil, symbolCallError("IOSurfaceGetBaseAddressOfPlane", "10.6", _iOSurfaceGetBaseAddressOfPlaneErr)
	}
	return _iOSurfaceGetBaseAddressOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetBaseAddressOfPlane returns the address of the first byte of data in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddressOfPlane(_:_:)
func IOSurfaceGetBaseAddressOfPlane(buffer IOSurfaceRef, planeIndex uintptr) unsafe.Pointer {
	result, callErr := tryIOSurfaceGetBaseAddressOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBitDepthOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr
var _iOSurfaceGetBitDepthOfComponentOfPlaneErr error

func tryIOSurfaceGetBitDepthOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetBitDepthOfComponentOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetBitDepthOfComponentOfPlane", "10.13", _iOSurfaceGetBitDepthOfComponentOfPlaneErr)
	}
	return _iOSurfaceGetBitDepthOfComponentOfPlane(buffer, planeIndex, componentIndex), nil
}

// IOSurfaceGetBitDepthOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitDepthOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitDepthOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetBitDepthOfComponentOfPlane(buffer, planeIndex, componentIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBitOffsetOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr
var _iOSurfaceGetBitOffsetOfComponentOfPlaneErr error

func tryIOSurfaceGetBitOffsetOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetBitOffsetOfComponentOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetBitOffsetOfComponentOfPlane", "10.13", _iOSurfaceGetBitOffsetOfComponentOfPlaneErr)
	}
	return _iOSurfaceGetBitOffsetOfComponentOfPlane(buffer, planeIndex, componentIndex), nil
}

// IOSurfaceGetBitOffsetOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitOffsetOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitOffsetOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetBitOffsetOfComponentOfPlane(buffer, planeIndex, componentIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBytesPerElement func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetBytesPerElementErr error

func tryIOSurfaceGetBytesPerElement(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetBytesPerElement == nil {
		return 0, symbolCallError("IOSurfaceGetBytesPerElement", "10.6", _iOSurfaceGetBytesPerElementErr)
	}
	return _iOSurfaceGetBytesPerElement(buffer), nil
}

// IOSurfaceGetBytesPerElement returns the length (in bytes) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElement(_:)
func IOSurfaceGetBytesPerElement(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetBytesPerElement(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBytesPerElementOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetBytesPerElementOfPlaneErr error

func tryIOSurfaceGetBytesPerElementOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetBytesPerElementOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetBytesPerElementOfPlane", "10.6", _iOSurfaceGetBytesPerElementOfPlaneErr)
	}
	return _iOSurfaceGetBytesPerElementOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetBytesPerElementOfPlane returns the size of each element (in bytes) in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElementOfPlane(_:_:)
func IOSurfaceGetBytesPerElementOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetBytesPerElementOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBytesPerRow func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetBytesPerRowErr error

func tryIOSurfaceGetBytesPerRow(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetBytesPerRow == nil {
		return 0, symbolCallError("IOSurfaceGetBytesPerRow", "10.6", _iOSurfaceGetBytesPerRowErr)
	}
	return _iOSurfaceGetBytesPerRow(buffer), nil
}

// IOSurfaceGetBytesPerRow returns the length (in bytes) of each row in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRow(_:)
func IOSurfaceGetBytesPerRow(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetBytesPerRow(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetBytesPerRowOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetBytesPerRowOfPlaneErr error

func tryIOSurfaceGetBytesPerRowOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetBytesPerRowOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetBytesPerRowOfPlane", "10.6", _iOSurfaceGetBytesPerRowOfPlaneErr)
	}
	return _iOSurfaceGetBytesPerRowOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetBytesPerRowOfPlane returns the size of each row (in bytes) in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRowOfPlane(_:_:)
func IOSurfaceGetBytesPerRowOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetBytesPerRowOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetElementHeight func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetElementHeightErr error

func tryIOSurfaceGetElementHeight(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetElementHeight == nil {
		return 0, symbolCallError("IOSurfaceGetElementHeight", "10.6", _iOSurfaceGetElementHeightErr)
	}
	return _iOSurfaceGetElementHeight(buffer), nil
}

// IOSurfaceGetElementHeight returns the height (in pixels) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeight(_:)
func IOSurfaceGetElementHeight(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetElementHeight(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetElementHeightOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetElementHeightOfPlaneErr error

func tryIOSurfaceGetElementHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetElementHeightOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetElementHeightOfPlane", "10.6", _iOSurfaceGetElementHeightOfPlaneErr)
	}
	return _iOSurfaceGetElementHeightOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetElementHeightOfPlane returns the height (in pixels) of each element in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeightOfPlane(_:_:)
func IOSurfaceGetElementHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetElementHeightOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetElementWidth func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetElementWidthErr error

func tryIOSurfaceGetElementWidth(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetElementWidth == nil {
		return 0, symbolCallError("IOSurfaceGetElementWidth", "10.6", _iOSurfaceGetElementWidthErr)
	}
	return _iOSurfaceGetElementWidth(buffer), nil
}

// IOSurfaceGetElementWidth returns the width (in pixels) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidth(_:)
func IOSurfaceGetElementWidth(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetElementWidth(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetElementWidthOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetElementWidthOfPlaneErr error

func tryIOSurfaceGetElementWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetElementWidthOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetElementWidthOfPlane", "10.6", _iOSurfaceGetElementWidthOfPlaneErr)
	}
	return _iOSurfaceGetElementWidthOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetElementWidthOfPlane returns the width (in pixels) of each element in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidthOfPlane(_:_:)
func IOSurfaceGetElementWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetElementWidthOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetHeight func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetHeightErr error

func tryIOSurfaceGetHeight(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetHeight == nil {
		return 0, symbolCallError("IOSurfaceGetHeight", "10.6", _iOSurfaceGetHeightErr)
	}
	return _iOSurfaceGetHeight(buffer), nil
}

// IOSurfaceGetHeight returns the height of the IOSurface buffer in pixels.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeight(_:)
func IOSurfaceGetHeight(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetHeight(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetHeightOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetHeightOfPlaneErr error

func tryIOSurfaceGetHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetHeightOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetHeightOfPlane", "10.6", _iOSurfaceGetHeightOfPlaneErr)
	}
	return _iOSurfaceGetHeightOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetHeightOfPlane returns the height of the specified plane (in pixels).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeightOfPlane(_:_:)
func IOSurfaceGetHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetHeightOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetID func(buffer IOSurfaceRef) uint32
var _iOSurfaceGetIDErr error

func tryIOSurfaceGetID(buffer IOSurfaceRef) (uint32, error) {
	if _iOSurfaceGetID == nil {
		return 0, symbolCallError("IOSurfaceGetID", "10.6", _iOSurfaceGetIDErr)
	}
	return _iOSurfaceGetID(buffer), nil
}

// IOSurfaceGetID retrieves the unique IOSurfaceID value for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetID(_:)
func IOSurfaceGetID(buffer IOSurfaceRef) uint32 {
	result, callErr := tryIOSurfaceGetID(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetNameOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentName
var _iOSurfaceGetNameOfComponentOfPlaneErr error

func tryIOSurfaceGetNameOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) (IOSurfaceComponentName, error) {
	if _iOSurfaceGetNameOfComponentOfPlane == nil {
		return *new(IOSurfaceComponentName), symbolCallError("IOSurfaceGetNameOfComponentOfPlane", "10.13", _iOSurfaceGetNameOfComponentOfPlaneErr)
	}
	return _iOSurfaceGetNameOfComponentOfPlane(buffer, planeIndex, componentIndex), nil
}

// IOSurfaceGetNameOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNameOfComponentOfPlane(_:_:_:)
func IOSurfaceGetNameOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentName {
	result, callErr := tryIOSurfaceGetNameOfComponentOfPlane(buffer, planeIndex, componentIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetNumberOfComponentsOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetNumberOfComponentsOfPlaneErr error

func tryIOSurfaceGetNumberOfComponentsOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetNumberOfComponentsOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetNumberOfComponentsOfPlane", "10.13", _iOSurfaceGetNumberOfComponentsOfPlaneErr)
	}
	return _iOSurfaceGetNumberOfComponentsOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetNumberOfComponentsOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNumberOfComponentsOfPlane(_:_:)
func IOSurfaceGetNumberOfComponentsOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetNumberOfComponentsOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetPixelFormat func(buffer IOSurfaceRef) uint32
var _iOSurfaceGetPixelFormatErr error

func tryIOSurfaceGetPixelFormat(buffer IOSurfaceRef) (uint32, error) {
	if _iOSurfaceGetPixelFormat == nil {
		return 0, symbolCallError("IOSurfaceGetPixelFormat", "10.6", _iOSurfaceGetPixelFormatErr)
	}
	return _iOSurfaceGetPixelFormat(buffer), nil
}

// IOSurfaceGetPixelFormat returns an unsigned integer that contains the traditional macOS buffer format.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPixelFormat(_:)
func IOSurfaceGetPixelFormat(buffer IOSurfaceRef) uint32 {
	result, callErr := tryIOSurfaceGetPixelFormat(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetPlaneCount func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetPlaneCountErr error

func tryIOSurfaceGetPlaneCount(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetPlaneCount == nil {
		return 0, symbolCallError("IOSurfaceGetPlaneCount", "10.6", _iOSurfaceGetPlaneCountErr)
	}
	return _iOSurfaceGetPlaneCount(buffer), nil
}

// IOSurfaceGetPlaneCount.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPlaneCount(_:)
func IOSurfaceGetPlaneCount(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetPlaneCount(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetPropertyAlignment func(property corefoundation.CFStringRef) uintptr
var _iOSurfaceGetPropertyAlignmentErr error

func tryIOSurfaceGetPropertyAlignment(property corefoundation.CFStringRef) (uintptr, error) {
	if _iOSurfaceGetPropertyAlignment == nil {
		return 0, symbolCallError("IOSurfaceGetPropertyAlignment", "10.6", _iOSurfaceGetPropertyAlignmentErr)
	}
	return _iOSurfaceGetPropertyAlignment(property), nil
}

// IOSurfaceGetPropertyAlignment returns the alignment requirements for a property (if any).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyAlignment(_:)
func IOSurfaceGetPropertyAlignment(property corefoundation.CFStringRef) uintptr {
	result, callErr := tryIOSurfaceGetPropertyAlignment(property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetPropertyMaximum func(property corefoundation.CFStringRef) uintptr
var _iOSurfaceGetPropertyMaximumErr error

func tryIOSurfaceGetPropertyMaximum(property corefoundation.CFStringRef) (uintptr, error) {
	if _iOSurfaceGetPropertyMaximum == nil {
		return 0, symbolCallError("IOSurfaceGetPropertyMaximum", "10.6", _iOSurfaceGetPropertyMaximumErr)
	}
	return _iOSurfaceGetPropertyMaximum(property), nil
}

// IOSurfaceGetPropertyMaximum returns the maximum value for a given property that is guaranteed to be compatible with all of the current devices (GPUs, etc.) in the system.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyMaximum(_:)
func IOSurfaceGetPropertyMaximum(property corefoundation.CFStringRef) uintptr {
	result, callErr := tryIOSurfaceGetPropertyMaximum(property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetRangeOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentRange
var _iOSurfaceGetRangeOfComponentOfPlaneErr error

func tryIOSurfaceGetRangeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) (IOSurfaceComponentRange, error) {
	if _iOSurfaceGetRangeOfComponentOfPlane == nil {
		return *new(IOSurfaceComponentRange), symbolCallError("IOSurfaceGetRangeOfComponentOfPlane", "10.13", _iOSurfaceGetRangeOfComponentOfPlaneErr)
	}
	return _iOSurfaceGetRangeOfComponentOfPlane(buffer, planeIndex, componentIndex), nil
}

// IOSurfaceGetRangeOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetRangeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetRangeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentRange {
	result, callErr := tryIOSurfaceGetRangeOfComponentOfPlane(buffer, planeIndex, componentIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetSeed func(buffer IOSurfaceRef) uint32
var _iOSurfaceGetSeedErr error

func tryIOSurfaceGetSeed(buffer IOSurfaceRef) (uint32, error) {
	if _iOSurfaceGetSeed == nil {
		return 0, symbolCallError("IOSurfaceGetSeed", "10.6", _iOSurfaceGetSeedErr)
	}
	return _iOSurfaceGetSeed(buffer), nil
}

// IOSurfaceGetSeed.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSeed(_:)
func IOSurfaceGetSeed(buffer IOSurfaceRef) uint32 {
	result, callErr := tryIOSurfaceGetSeed(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetSubsampling func(buffer IOSurfaceRef) IOSurfaceSubsampling
var _iOSurfaceGetSubsamplingErr error

func tryIOSurfaceGetSubsampling(buffer IOSurfaceRef) (IOSurfaceSubsampling, error) {
	if _iOSurfaceGetSubsampling == nil {
		return *new(IOSurfaceSubsampling), symbolCallError("IOSurfaceGetSubsampling", "10.13", _iOSurfaceGetSubsamplingErr)
	}
	return _iOSurfaceGetSubsampling(buffer), nil
}

// IOSurfaceGetSubsampling.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSubsampling(_:)
func IOSurfaceGetSubsampling(buffer IOSurfaceRef) IOSurfaceSubsampling {
	result, callErr := tryIOSurfaceGetSubsampling(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetTypeID func() uint
var _iOSurfaceGetTypeIDErr error

func tryIOSurfaceGetTypeID() (uint, error) {
	if _iOSurfaceGetTypeID == nil {
		return 0, symbolCallError("IOSurfaceGetTypeID", "10.6", _iOSurfaceGetTypeIDErr)
	}
	return _iOSurfaceGetTypeID(), nil
}

// IOSurfaceGetTypeID.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeID()
func IOSurfaceGetTypeID() uint {
	result, callErr := tryIOSurfaceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetTypeOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentType
var _iOSurfaceGetTypeOfComponentOfPlaneErr error

func tryIOSurfaceGetTypeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) (IOSurfaceComponentType, error) {
	if _iOSurfaceGetTypeOfComponentOfPlane == nil {
		return *new(IOSurfaceComponentType), symbolCallError("IOSurfaceGetTypeOfComponentOfPlane", "10.13", _iOSurfaceGetTypeOfComponentOfPlaneErr)
	}
	return _iOSurfaceGetTypeOfComponentOfPlane(buffer, planeIndex, componentIndex), nil
}

// IOSurfaceGetTypeOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetTypeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentType {
	result, callErr := tryIOSurfaceGetTypeOfComponentOfPlane(buffer, planeIndex, componentIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetUseCount func(buffer IOSurfaceRef) int32
var _iOSurfaceGetUseCountErr error

func tryIOSurfaceGetUseCount(buffer IOSurfaceRef) (int32, error) {
	if _iOSurfaceGetUseCount == nil {
		return 0, symbolCallError("IOSurfaceGetUseCount", "10.6", _iOSurfaceGetUseCountErr)
	}
	return _iOSurfaceGetUseCount(buffer), nil
}

// IOSurfaceGetUseCount returns the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetUseCount(_:)
func IOSurfaceGetUseCount(buffer IOSurfaceRef) int32 {
	result, callErr := tryIOSurfaceGetUseCount(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetWidth func(buffer IOSurfaceRef) uintptr
var _iOSurfaceGetWidthErr error

func tryIOSurfaceGetWidth(buffer IOSurfaceRef) (uintptr, error) {
	if _iOSurfaceGetWidth == nil {
		return 0, symbolCallError("IOSurfaceGetWidth", "10.6", _iOSurfaceGetWidthErr)
	}
	return _iOSurfaceGetWidth(buffer), nil
}

// IOSurfaceGetWidth returns the width of the IOSurface buffer in pixels.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidth(_:)
func IOSurfaceGetWidth(buffer IOSurfaceRef) uintptr {
	result, callErr := tryIOSurfaceGetWidth(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceGetWidthOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr
var _iOSurfaceGetWidthOfPlaneErr error

func tryIOSurfaceGetWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) (uintptr, error) {
	if _iOSurfaceGetWidthOfPlane == nil {
		return 0, symbolCallError("IOSurfaceGetWidthOfPlane", "10.6", _iOSurfaceGetWidthOfPlaneErr)
	}
	return _iOSurfaceGetWidthOfPlane(buffer, planeIndex), nil
}

// IOSurfaceGetWidthOfPlane returns the width of the specified plane (in pixels).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidthOfPlane(_:_:)
func IOSurfaceGetWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	result, callErr := tryIOSurfaceGetWidthOfPlane(buffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceIncrementUseCount func(buffer IOSurfaceRef)
var _iOSurfaceIncrementUseCountErr error

func tryIOSurfaceIncrementUseCount(buffer IOSurfaceRef) error {
	if _iOSurfaceIncrementUseCount == nil {
		return symbolCallError("IOSurfaceIncrementUseCount", "10.6", _iOSurfaceIncrementUseCountErr)
	}
	_iOSurfaceIncrementUseCount(buffer)
	return nil
}

// IOSurfaceIncrementUseCount increments the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceIncrementUseCount(_:)
func IOSurfaceIncrementUseCount(buffer IOSurfaceRef) {
	if callErr := tryIOSurfaceIncrementUseCount(buffer); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceIsInUse func(buffer IOSurfaceRef) bool
var _iOSurfaceIsInUseErr error

func tryIOSurfaceIsInUse(buffer IOSurfaceRef) (bool, error) {
	if _iOSurfaceIsInUse == nil {
		return false, symbolCallError("IOSurfaceIsInUse", "10.6", _iOSurfaceIsInUseErr)
	}
	return _iOSurfaceIsInUse(buffer), nil
}

// IOSurfaceIsInUse returns true of an IOSurface is in use by any process in the system, otherwise false.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceIsInUse(_:)
func IOSurfaceIsInUse(buffer IOSurfaceRef) bool {
	result, callErr := tryIOSurfaceIsInUse(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceLock func(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32
var _iOSurfaceLockErr error

func tryIOSurfaceLock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) (int32, error) {
	if _iOSurfaceLock == nil {
		return 0, symbolCallError("IOSurfaceLock", "10.6", _iOSurfaceLockErr)
	}
	return _iOSurfaceLock(buffer, options, seed), nil
}

// IOSurfaceLock “Lock” an IOSurface for reading or writing.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLock(_:_:_:)
func IOSurfaceLock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32 {
	result, callErr := tryIOSurfaceLock(buffer, options, seed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceLookup func(csid uint32) IOSurfaceRef
var _iOSurfaceLookupErr error

func tryIOSurfaceLookup(csid uint32) (IOSurfaceRef, error) {
	if _iOSurfaceLookup == nil {
		return 0, symbolCallError("IOSurfaceLookup", "10.6", _iOSurfaceLookupErr)
	}
	return _iOSurfaceLookup(csid), nil
}

// IOSurfaceLookup performs an atomic lookup and retain of an IOSurface by its IOSurfaceID.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookup(_:)
func IOSurfaceLookup(csid uint32) IOSurfaceRef {
	result, callErr := tryIOSurfaceLookup(csid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceLookupFromMachPort func(port uint32) IOSurfaceRef
var _iOSurfaceLookupFromMachPortErr error

func tryIOSurfaceLookupFromMachPort(port uint32) (IOSurfaceRef, error) {
	if _iOSurfaceLookupFromMachPort == nil {
		return 0, symbolCallError("IOSurfaceLookupFromMachPort", "10.6", _iOSurfaceLookupFromMachPortErr)
	}
	return _iOSurfaceLookupFromMachPort(port), nil
}

// IOSurfaceLookupFromMachPort recreates an IOSurfaceRef from a mach port.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromMachPort(_:)
func IOSurfaceLookupFromMachPort(port uint32) IOSurfaceRef {
	result, callErr := tryIOSurfaceLookupFromMachPort(port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceLookupFromXPCObject func(xobj unsafe.Pointer) IOSurfaceRef
var _iOSurfaceLookupFromXPCObjectErr error

func tryIOSurfaceLookupFromXPCObject(xobj unsafe.Pointer) (IOSurfaceRef, error) {
	if _iOSurfaceLookupFromXPCObject == nil {
		return 0, symbolCallError("IOSurfaceLookupFromXPCObject", "10.7", _iOSurfaceLookupFromXPCObjectErr)
	}
	return _iOSurfaceLookupFromXPCObject(xobj), nil
}

// IOSurfaceLookupFromXPCObject.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromXPCObject(_:)
func IOSurfaceLookupFromXPCObject(xobj unsafe.Pointer) IOSurfaceRef {
	result, callErr := tryIOSurfaceLookupFromXPCObject(xobj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceRemoveAllValues func(buffer IOSurfaceRef)
var _iOSurfaceRemoveAllValuesErr error

func tryIOSurfaceRemoveAllValues(buffer IOSurfaceRef) error {
	if _iOSurfaceRemoveAllValues == nil {
		return symbolCallError("IOSurfaceRemoveAllValues", "10.6", _iOSurfaceRemoveAllValuesErr)
	}
	_iOSurfaceRemoveAllValues(buffer)
	return nil
}

// IOSurfaceRemoveAllValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveAllValues(_:)
func IOSurfaceRemoveAllValues(buffer IOSurfaceRef) {
	if callErr := tryIOSurfaceRemoveAllValues(buffer); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceRemoveValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef)
var _iOSurfaceRemoveValueErr error

func tryIOSurfaceRemoveValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) error {
	if _iOSurfaceRemoveValue == nil {
		return symbolCallError("IOSurfaceRemoveValue", "10.6", _iOSurfaceRemoveValueErr)
	}
	_iOSurfaceRemoveValue(buffer, key)
	return nil
}

// IOSurfaceRemoveValue deletes a value in the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveValue(_:_:)
func IOSurfaceRemoveValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) {
	if callErr := tryIOSurfaceRemoveValue(buffer, key); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceSetOwnershipIdentity func(buffer IOSurfaceRef, task_id_token kernel.Task_id_token_t, newLedgerTag int, newLedgerOptions uint32) int32
var _iOSurfaceSetOwnershipIdentityErr error

func tryIOSurfaceSetOwnershipIdentity(buffer IOSurfaceRef, task_id_token kernel.Task_id_token_t, newLedgerTag int, newLedgerOptions uint32) (int32, error) {
	if _iOSurfaceSetOwnershipIdentity == nil {
		return 0, symbolCallError("IOSurfaceSetOwnershipIdentity", "14.4", _iOSurfaceSetOwnershipIdentityErr)
	}
	return _iOSurfaceSetOwnershipIdentity(buffer, task_id_token, newLedgerTag, newLedgerOptions), nil
}

// IOSurfaceSetOwnershipIdentity.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetOwnershipIdentity(_:_:_:_:)
func IOSurfaceSetOwnershipIdentity(buffer IOSurfaceRef, task_id_token kernel.Task_id_token_t, newLedgerTag int, newLedgerOptions uint32) int32 {
	result, callErr := tryIOSurfaceSetOwnershipIdentity(buffer, task_id_token, newLedgerTag, newLedgerOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceSetPurgeable func(buffer IOSurfaceRef, newState uint32, oldState *uint32) int32
var _iOSurfaceSetPurgeableErr error

func tryIOSurfaceSetPurgeable(buffer IOSurfaceRef, newState uint32, oldState *uint32) (int32, error) {
	if _iOSurfaceSetPurgeable == nil {
		return 0, symbolCallError("IOSurfaceSetPurgeable", "10.12", _iOSurfaceSetPurgeableErr)
	}
	return _iOSurfaceSetPurgeable(buffer, newState, oldState), nil
}

// IOSurfaceSetPurgeable.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetPurgeable(_:_:_:)
func IOSurfaceSetPurgeable(buffer IOSurfaceRef, newState uint32, oldState *uint32) int32 {
	result, callErr := tryIOSurfaceSetPurgeable(buffer, newState, oldState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOSurfaceSetValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef)
var _iOSurfaceSetValueErr error

func tryIOSurfaceSetValue(buffer IOSurfaceRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) error {
	if _iOSurfaceSetValue == nil {
		return symbolCallError("IOSurfaceSetValue", "10.6", _iOSurfaceSetValueErr)
	}
	_iOSurfaceSetValue(buffer, key, value)
	return nil
}

// IOSurfaceSetValue sets a value in the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValue(_:_:_:)
func IOSurfaceSetValue(buffer IOSurfaceRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) {
	if callErr := tryIOSurfaceSetValue(buffer, key, value); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceSetValues func(buffer IOSurfaceRef, keysAndValues corefoundation.CFDictionaryRef)
var _iOSurfaceSetValuesErr error

func tryIOSurfaceSetValues(buffer IOSurfaceRef, keysAndValues corefoundation.CFDictionaryRef) error {
	if _iOSurfaceSetValues == nil {
		return symbolCallError("IOSurfaceSetValues", "10.6", _iOSurfaceSetValuesErr)
	}
	_iOSurfaceSetValues(buffer, keysAndValues)
	return nil
}

// IOSurfaceSetValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValues(_:_:)
func IOSurfaceSetValues(buffer IOSurfaceRef, keysAndValues corefoundation.CFDictionaryRef) {
	if callErr := tryIOSurfaceSetValues(buffer, keysAndValues); callErr != nil {
		panic(callErr)
	}
}

var _iOSurfaceUnlock func(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32
var _iOSurfaceUnlockErr error

func tryIOSurfaceUnlock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) (int32, error) {
	if _iOSurfaceUnlock == nil {
		return 0, symbolCallError("IOSurfaceUnlock", "10.6", _iOSurfaceUnlockErr)
	}
	return _iOSurfaceUnlock(buffer, options, seed), nil
}

// IOSurfaceUnlock “Unlock” an IOSurface for reading or writing.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceUnlock(_:_:_:)
func IOSurfaceUnlock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32 {
	result, callErr := tryIOSurfaceUnlock(buffer, options, seed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_iOSurfaceAlignProperty, &_iOSurfaceAlignPropertyErr, frameworkHandle, "IOSurfaceAlignProperty", "10.6")
	registerFunc(&_iOSurfaceAllowsPixelSizeCasting, &_iOSurfaceAllowsPixelSizeCastingErr, frameworkHandle, "IOSurfaceAllowsPixelSizeCasting", "10.12")
	registerFunc(&_iOSurfaceCopyAllValues, &_iOSurfaceCopyAllValuesErr, frameworkHandle, "IOSurfaceCopyAllValues", "10.6")
	registerFunc(&_iOSurfaceCopyValue, &_iOSurfaceCopyValueErr, frameworkHandle, "IOSurfaceCopyValue", "10.6")
	registerFunc(&_iOSurfaceCreate, &_iOSurfaceCreateErr, frameworkHandle, "IOSurfaceCreate", "10.6")
	registerFunc(&_iOSurfaceCreateMachPort, &_iOSurfaceCreateMachPortErr, frameworkHandle, "IOSurfaceCreateMachPort", "10.6")
	registerFunc(&_iOSurfaceCreateXPCObject, &_iOSurfaceCreateXPCObjectErr, frameworkHandle, "IOSurfaceCreateXPCObject", "10.7")
	registerFunc(&_iOSurfaceDecrementUseCount, &_iOSurfaceDecrementUseCountErr, frameworkHandle, "IOSurfaceDecrementUseCount", "10.6")
	registerFunc(&_iOSurfaceGetAllocSize, &_iOSurfaceGetAllocSizeErr, frameworkHandle, "IOSurfaceGetAllocSize", "10.6")
	registerFunc(&_iOSurfaceGetBaseAddress, &_iOSurfaceGetBaseAddressErr, frameworkHandle, "IOSurfaceGetBaseAddress", "10.6")
	registerFunc(&_iOSurfaceGetBaseAddressOfPlane, &_iOSurfaceGetBaseAddressOfPlaneErr, frameworkHandle, "IOSurfaceGetBaseAddressOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetBitDepthOfComponentOfPlane, &_iOSurfaceGetBitDepthOfComponentOfPlaneErr, frameworkHandle, "IOSurfaceGetBitDepthOfComponentOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetBitOffsetOfComponentOfPlane, &_iOSurfaceGetBitOffsetOfComponentOfPlaneErr, frameworkHandle, "IOSurfaceGetBitOffsetOfComponentOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetBytesPerElement, &_iOSurfaceGetBytesPerElementErr, frameworkHandle, "IOSurfaceGetBytesPerElement", "10.6")
	registerFunc(&_iOSurfaceGetBytesPerElementOfPlane, &_iOSurfaceGetBytesPerElementOfPlaneErr, frameworkHandle, "IOSurfaceGetBytesPerElementOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetBytesPerRow, &_iOSurfaceGetBytesPerRowErr, frameworkHandle, "IOSurfaceGetBytesPerRow", "10.6")
	registerFunc(&_iOSurfaceGetBytesPerRowOfPlane, &_iOSurfaceGetBytesPerRowOfPlaneErr, frameworkHandle, "IOSurfaceGetBytesPerRowOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetElementHeight, &_iOSurfaceGetElementHeightErr, frameworkHandle, "IOSurfaceGetElementHeight", "10.6")
	registerFunc(&_iOSurfaceGetElementHeightOfPlane, &_iOSurfaceGetElementHeightOfPlaneErr, frameworkHandle, "IOSurfaceGetElementHeightOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetElementWidth, &_iOSurfaceGetElementWidthErr, frameworkHandle, "IOSurfaceGetElementWidth", "10.6")
	registerFunc(&_iOSurfaceGetElementWidthOfPlane, &_iOSurfaceGetElementWidthOfPlaneErr, frameworkHandle, "IOSurfaceGetElementWidthOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetHeight, &_iOSurfaceGetHeightErr, frameworkHandle, "IOSurfaceGetHeight", "10.6")
	registerFunc(&_iOSurfaceGetHeightOfPlane, &_iOSurfaceGetHeightOfPlaneErr, frameworkHandle, "IOSurfaceGetHeightOfPlane", "10.6")
	registerFunc(&_iOSurfaceGetID, &_iOSurfaceGetIDErr, frameworkHandle, "IOSurfaceGetID", "10.6")
	registerFunc(&_iOSurfaceGetNameOfComponentOfPlane, &_iOSurfaceGetNameOfComponentOfPlaneErr, frameworkHandle, "IOSurfaceGetNameOfComponentOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetNumberOfComponentsOfPlane, &_iOSurfaceGetNumberOfComponentsOfPlaneErr, frameworkHandle, "IOSurfaceGetNumberOfComponentsOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetPixelFormat, &_iOSurfaceGetPixelFormatErr, frameworkHandle, "IOSurfaceGetPixelFormat", "10.6")
	registerFunc(&_iOSurfaceGetPlaneCount, &_iOSurfaceGetPlaneCountErr, frameworkHandle, "IOSurfaceGetPlaneCount", "10.6")
	registerFunc(&_iOSurfaceGetPropertyAlignment, &_iOSurfaceGetPropertyAlignmentErr, frameworkHandle, "IOSurfaceGetPropertyAlignment", "10.6")
	registerFunc(&_iOSurfaceGetPropertyMaximum, &_iOSurfaceGetPropertyMaximumErr, frameworkHandle, "IOSurfaceGetPropertyMaximum", "10.6")
	registerFunc(&_iOSurfaceGetRangeOfComponentOfPlane, &_iOSurfaceGetRangeOfComponentOfPlaneErr, frameworkHandle, "IOSurfaceGetRangeOfComponentOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetSeed, &_iOSurfaceGetSeedErr, frameworkHandle, "IOSurfaceGetSeed", "10.6")
	registerFunc(&_iOSurfaceGetSubsampling, &_iOSurfaceGetSubsamplingErr, frameworkHandle, "IOSurfaceGetSubsampling", "10.13")
	registerFunc(&_iOSurfaceGetTypeID, &_iOSurfaceGetTypeIDErr, frameworkHandle, "IOSurfaceGetTypeID", "10.6")
	registerFunc(&_iOSurfaceGetTypeOfComponentOfPlane, &_iOSurfaceGetTypeOfComponentOfPlaneErr, frameworkHandle, "IOSurfaceGetTypeOfComponentOfPlane", "10.13")
	registerFunc(&_iOSurfaceGetUseCount, &_iOSurfaceGetUseCountErr, frameworkHandle, "IOSurfaceGetUseCount", "10.6")
	registerFunc(&_iOSurfaceGetWidth, &_iOSurfaceGetWidthErr, frameworkHandle, "IOSurfaceGetWidth", "10.6")
	registerFunc(&_iOSurfaceGetWidthOfPlane, &_iOSurfaceGetWidthOfPlaneErr, frameworkHandle, "IOSurfaceGetWidthOfPlane", "10.6")
	registerFunc(&_iOSurfaceIncrementUseCount, &_iOSurfaceIncrementUseCountErr, frameworkHandle, "IOSurfaceIncrementUseCount", "10.6")
	registerFunc(&_iOSurfaceIsInUse, &_iOSurfaceIsInUseErr, frameworkHandle, "IOSurfaceIsInUse", "10.6")
	registerFunc(&_iOSurfaceLock, &_iOSurfaceLockErr, frameworkHandle, "IOSurfaceLock", "10.6")
	registerFunc(&_iOSurfaceLookup, &_iOSurfaceLookupErr, frameworkHandle, "IOSurfaceLookup", "10.6")
	registerFunc(&_iOSurfaceLookupFromMachPort, &_iOSurfaceLookupFromMachPortErr, frameworkHandle, "IOSurfaceLookupFromMachPort", "10.6")
	registerFunc(&_iOSurfaceLookupFromXPCObject, &_iOSurfaceLookupFromXPCObjectErr, frameworkHandle, "IOSurfaceLookupFromXPCObject", "10.7")
	registerFunc(&_iOSurfaceRemoveAllValues, &_iOSurfaceRemoveAllValuesErr, frameworkHandle, "IOSurfaceRemoveAllValues", "10.6")
	registerFunc(&_iOSurfaceRemoveValue, &_iOSurfaceRemoveValueErr, frameworkHandle, "IOSurfaceRemoveValue", "10.6")
	registerFunc(&_iOSurfaceSetOwnershipIdentity, &_iOSurfaceSetOwnershipIdentityErr, frameworkHandle, "IOSurfaceSetOwnershipIdentity", "14.4")
	registerFunc(&_iOSurfaceSetPurgeable, &_iOSurfaceSetPurgeableErr, frameworkHandle, "IOSurfaceSetPurgeable", "10.12")
	registerFunc(&_iOSurfaceSetValue, &_iOSurfaceSetValueErr, frameworkHandle, "IOSurfaceSetValue", "10.6")
	registerFunc(&_iOSurfaceSetValues, &_iOSurfaceSetValuesErr, frameworkHandle, "IOSurfaceSetValues", "10.6")
	registerFunc(&_iOSurfaceUnlock, &_iOSurfaceUnlockErr, frameworkHandle, "IOSurfaceUnlock", "10.6")
}
