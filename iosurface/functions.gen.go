// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: IOSurface: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: IOSurface: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: IOSurface: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _iOSurfaceAlignProperty func(property corefoundation.CFStringRef, value uintptr) uintptr

// IOSurfaceAlignProperty returns the smallest aligned value greater than or equal to the specified value.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceAlignProperty(_:_:)
func IOSurfaceAlignProperty(property corefoundation.CFStringRef, value uintptr) uintptr {
	if _iOSurfaceAlignProperty == nil {
		panic("IOSurface: symbol IOSurfaceAlignProperty not loaded")
	}
	return _iOSurfaceAlignProperty(property, value)
}


var _iOSurfaceAllowsPixelSizeCasting func(buffer IOSurfaceRef) bool

// IOSurfaceAllowsPixelSizeCasting.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceAllowsPixelSizeCasting(_:)
func IOSurfaceAllowsPixelSizeCasting(buffer IOSurfaceRef) bool {
	if _iOSurfaceAllowsPixelSizeCasting == nil {
		panic("IOSurface: symbol IOSurfaceAllowsPixelSizeCasting not loaded")
	}
	return _iOSurfaceAllowsPixelSizeCasting(buffer)
}


var _iOSurfaceCopyAllValues func(buffer IOSurfaceRef) corefoundation.CFDictionaryRef

// IOSurfaceCopyAllValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyAllValues(_:)
func IOSurfaceCopyAllValues(buffer IOSurfaceRef) corefoundation.CFDictionaryRef {
	if _iOSurfaceCopyAllValues == nil {
		panic("IOSurface: symbol IOSurfaceCopyAllValues not loaded")
	}
	return _iOSurfaceCopyAllValues(buffer)
}


var _iOSurfaceCopyValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOSurfaceCopyValue retrieves a value from the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCopyValue(_:_:)
func IOSurfaceCopyValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOSurfaceCopyValue == nil {
		panic("IOSurface: symbol IOSurfaceCopyValue not loaded")
	}
	return _iOSurfaceCopyValue(buffer, key)
}


var _iOSurfaceCreate func(properties corefoundation.CFDictionaryRef) IOSurfaceRef

// IOSurfaceCreate creates a brand new IOSurface object
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreate(_:)
func IOSurfaceCreate(properties corefoundation.CFDictionaryRef) IOSurfaceRef {
	if _iOSurfaceCreate == nil {
		panic("IOSurface: symbol IOSurfaceCreate not loaded")
	}
	return _iOSurfaceCreate(properties)
}


var _iOSurfaceCreateMachPort func(buffer IOSurfaceRef) uint32

// IOSurfaceCreateMachPort returns a mach_port_t that holds a reference to the IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateMachPort(_:)
func IOSurfaceCreateMachPort(buffer IOSurfaceRef) uint32 {
	if _iOSurfaceCreateMachPort == nil {
		panic("IOSurface: symbol IOSurfaceCreateMachPort not loaded")
	}
	return _iOSurfaceCreateMachPort(buffer)
}


var _iOSurfaceCreateXPCObject func(aSurface IOSurfaceRef) unsafe.Pointer

// IOSurfaceCreateXPCObject returns an xpc_object_t that holds a reference to the IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateXPCObject(_:)
func IOSurfaceCreateXPCObject(aSurface IOSurfaceRef) unsafe.Pointer {
	if _iOSurfaceCreateXPCObject == nil {
		panic("IOSurface: symbol IOSurfaceCreateXPCObject not loaded")
	}
	return _iOSurfaceCreateXPCObject(aSurface)
}


var _iOSurfaceDecrementUseCount func(buffer IOSurfaceRef)

// IOSurfaceDecrementUseCount decrements the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceDecrementUseCount(_:)
func IOSurfaceDecrementUseCount(buffer IOSurfaceRef) {
	if _iOSurfaceDecrementUseCount == nil {
		panic("IOSurface: symbol IOSurfaceDecrementUseCount not loaded")
	}
	_iOSurfaceDecrementUseCount(buffer)
}


var _iOSurfaceGetAllocSize func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetAllocSize returns the total allocation size of the buffer including all planes.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetAllocSize(_:)
func IOSurfaceGetAllocSize(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetAllocSize == nil {
		panic("IOSurface: symbol IOSurfaceGetAllocSize not loaded")
	}
	return _iOSurfaceGetAllocSize(buffer)
}


var _iOSurfaceGetBaseAddress func(buffer IOSurfaceRef) unsafe.Pointer

// IOSurfaceGetBaseAddress returns the address of the first byte of data in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddress(_:)
func IOSurfaceGetBaseAddress(buffer IOSurfaceRef) unsafe.Pointer {
	if _iOSurfaceGetBaseAddress == nil {
		panic("IOSurface: symbol IOSurfaceGetBaseAddress not loaded")
	}
	return _iOSurfaceGetBaseAddress(buffer)
}


var _iOSurfaceGetBaseAddressOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) unsafe.Pointer

// IOSurfaceGetBaseAddressOfPlane returns the address of the first byte of data in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBaseAddressOfPlane(_:_:)
func IOSurfaceGetBaseAddressOfPlane(buffer IOSurfaceRef, planeIndex uintptr) unsafe.Pointer {
	if _iOSurfaceGetBaseAddressOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetBaseAddressOfPlane not loaded")
	}
	return _iOSurfaceGetBaseAddressOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetBitDepthOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr

// IOSurfaceGetBitDepthOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitDepthOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitDepthOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr {
	if _iOSurfaceGetBitDepthOfComponentOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetBitDepthOfComponentOfPlane not loaded")
	}
	return _iOSurfaceGetBitDepthOfComponentOfPlane(buffer, planeIndex, componentIndex)
}


var _iOSurfaceGetBitOffsetOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr

// IOSurfaceGetBitOffsetOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBitOffsetOfComponentOfPlane(_:_:_:)
func IOSurfaceGetBitOffsetOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) uintptr {
	if _iOSurfaceGetBitOffsetOfComponentOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetBitOffsetOfComponentOfPlane not loaded")
	}
	return _iOSurfaceGetBitOffsetOfComponentOfPlane(buffer, planeIndex, componentIndex)
}


var _iOSurfaceGetBytesPerElement func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetBytesPerElement returns the length (in bytes) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElement(_:)
func IOSurfaceGetBytesPerElement(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetBytesPerElement == nil {
		panic("IOSurface: symbol IOSurfaceGetBytesPerElement not loaded")
	}
	return _iOSurfaceGetBytesPerElement(buffer)
}


var _iOSurfaceGetBytesPerElementOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetBytesPerElementOfPlane returns the size of each element (in bytes) in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerElementOfPlane(_:_:)
func IOSurfaceGetBytesPerElementOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetBytesPerElementOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetBytesPerElementOfPlane not loaded")
	}
	return _iOSurfaceGetBytesPerElementOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetBytesPerRow func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetBytesPerRow returns the length (in bytes) of each row in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRow(_:)
func IOSurfaceGetBytesPerRow(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetBytesPerRow == nil {
		panic("IOSurface: symbol IOSurfaceGetBytesPerRow not loaded")
	}
	return _iOSurfaceGetBytesPerRow(buffer)
}


var _iOSurfaceGetBytesPerRowOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetBytesPerRowOfPlane returns the size of each row (in bytes) in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetBytesPerRowOfPlane(_:_:)
func IOSurfaceGetBytesPerRowOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetBytesPerRowOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetBytesPerRowOfPlane not loaded")
	}
	return _iOSurfaceGetBytesPerRowOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetElementHeight func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetElementHeight returns the height (in pixels) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeight(_:)
func IOSurfaceGetElementHeight(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetElementHeight == nil {
		panic("IOSurface: symbol IOSurfaceGetElementHeight not loaded")
	}
	return _iOSurfaceGetElementHeight(buffer)
}


var _iOSurfaceGetElementHeightOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetElementHeightOfPlane returns the height (in pixels) of each element in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementHeightOfPlane(_:_:)
func IOSurfaceGetElementHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetElementHeightOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetElementHeightOfPlane not loaded")
	}
	return _iOSurfaceGetElementHeightOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetElementWidth func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetElementWidth returns the width (in pixels) of each element in a particular buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidth(_:)
func IOSurfaceGetElementWidth(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetElementWidth == nil {
		panic("IOSurface: symbol IOSurfaceGetElementWidth not loaded")
	}
	return _iOSurfaceGetElementWidth(buffer)
}


var _iOSurfaceGetElementWidthOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetElementWidthOfPlane returns the width (in pixels) of each element in the specified plane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetElementWidthOfPlane(_:_:)
func IOSurfaceGetElementWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetElementWidthOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetElementWidthOfPlane not loaded")
	}
	return _iOSurfaceGetElementWidthOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetHeight func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetHeight returns the height of the IOSurface buffer in pixels.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeight(_:)
func IOSurfaceGetHeight(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetHeight == nil {
		panic("IOSurface: symbol IOSurfaceGetHeight not loaded")
	}
	return _iOSurfaceGetHeight(buffer)
}


var _iOSurfaceGetHeightOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetHeightOfPlane returns the height of the specified plane (in pixels).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetHeightOfPlane(_:_:)
func IOSurfaceGetHeightOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetHeightOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetHeightOfPlane not loaded")
	}
	return _iOSurfaceGetHeightOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetID func(buffer IOSurfaceRef) uint32

// IOSurfaceGetID retrieves the unique IOSurfaceID value for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetID(_:)
func IOSurfaceGetID(buffer IOSurfaceRef) uint32 {
	if _iOSurfaceGetID == nil {
		panic("IOSurface: symbol IOSurfaceGetID not loaded")
	}
	return _iOSurfaceGetID(buffer)
}


var _iOSurfaceGetNameOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentName

// IOSurfaceGetNameOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNameOfComponentOfPlane(_:_:_:)
func IOSurfaceGetNameOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentName {
	if _iOSurfaceGetNameOfComponentOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetNameOfComponentOfPlane not loaded")
	}
	return _iOSurfaceGetNameOfComponentOfPlane(buffer, planeIndex, componentIndex)
}


var _iOSurfaceGetNumberOfComponentsOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetNumberOfComponentsOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetNumberOfComponentsOfPlane(_:_:)
func IOSurfaceGetNumberOfComponentsOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetNumberOfComponentsOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetNumberOfComponentsOfPlane not loaded")
	}
	return _iOSurfaceGetNumberOfComponentsOfPlane(buffer, planeIndex)
}


var _iOSurfaceGetPixelFormat func(buffer IOSurfaceRef) uint32

// IOSurfaceGetPixelFormat returns an unsigned integer that contains the traditional macOS buffer format.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPixelFormat(_:)
func IOSurfaceGetPixelFormat(buffer IOSurfaceRef) uint32 {
	if _iOSurfaceGetPixelFormat == nil {
		panic("IOSurface: symbol IOSurfaceGetPixelFormat not loaded")
	}
	return _iOSurfaceGetPixelFormat(buffer)
}


var _iOSurfaceGetPlaneCount func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetPlaneCount.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPlaneCount(_:)
func IOSurfaceGetPlaneCount(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetPlaneCount == nil {
		panic("IOSurface: symbol IOSurfaceGetPlaneCount not loaded")
	}
	return _iOSurfaceGetPlaneCount(buffer)
}


var _iOSurfaceGetPropertyAlignment func(property corefoundation.CFStringRef) uintptr

// IOSurfaceGetPropertyAlignment returns the alignment requirements for a property (if any).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyAlignment(_:)
func IOSurfaceGetPropertyAlignment(property corefoundation.CFStringRef) uintptr {
	if _iOSurfaceGetPropertyAlignment == nil {
		panic("IOSurface: symbol IOSurfaceGetPropertyAlignment not loaded")
	}
	return _iOSurfaceGetPropertyAlignment(property)
}


var _iOSurfaceGetPropertyMaximum func(property corefoundation.CFStringRef) uintptr

// IOSurfaceGetPropertyMaximum returns the maximum value for a given property that is guaranteed to be compatible with all of the current devices (GPUs, etc.) in the system.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetPropertyMaximum(_:)
func IOSurfaceGetPropertyMaximum(property corefoundation.CFStringRef) uintptr {
	if _iOSurfaceGetPropertyMaximum == nil {
		panic("IOSurface: symbol IOSurfaceGetPropertyMaximum not loaded")
	}
	return _iOSurfaceGetPropertyMaximum(property)
}


var _iOSurfaceGetRangeOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentRange

// IOSurfaceGetRangeOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetRangeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetRangeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentRange {
	if _iOSurfaceGetRangeOfComponentOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetRangeOfComponentOfPlane not loaded")
	}
	return _iOSurfaceGetRangeOfComponentOfPlane(buffer, planeIndex, componentIndex)
}


var _iOSurfaceGetSeed func(buffer IOSurfaceRef) uint32

// IOSurfaceGetSeed.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSeed(_:)
func IOSurfaceGetSeed(buffer IOSurfaceRef) uint32 {
	if _iOSurfaceGetSeed == nil {
		panic("IOSurface: symbol IOSurfaceGetSeed not loaded")
	}
	return _iOSurfaceGetSeed(buffer)
}


var _iOSurfaceGetSubsampling func(buffer IOSurfaceRef) IOSurfaceSubsampling

// IOSurfaceGetSubsampling.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetSubsampling(_:)
func IOSurfaceGetSubsampling(buffer IOSurfaceRef) IOSurfaceSubsampling {
	if _iOSurfaceGetSubsampling == nil {
		panic("IOSurface: symbol IOSurfaceGetSubsampling not loaded")
	}
	return _iOSurfaceGetSubsampling(buffer)
}


var _iOSurfaceGetTypeID func() uint

// IOSurfaceGetTypeID.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeID()
func IOSurfaceGetTypeID() uint {
	if _iOSurfaceGetTypeID == nil {
		panic("IOSurface: symbol IOSurfaceGetTypeID not loaded")
	}
	return _iOSurfaceGetTypeID()
}


var _iOSurfaceGetTypeOfComponentOfPlane func(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentType

// IOSurfaceGetTypeOfComponentOfPlane.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetTypeOfComponentOfPlane(_:_:_:)
func IOSurfaceGetTypeOfComponentOfPlane(buffer IOSurfaceRef, planeIndex uintptr, componentIndex uintptr) IOSurfaceComponentType {
	if _iOSurfaceGetTypeOfComponentOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetTypeOfComponentOfPlane not loaded")
	}
	return _iOSurfaceGetTypeOfComponentOfPlane(buffer, planeIndex, componentIndex)
}


var _iOSurfaceGetUseCount func(buffer IOSurfaceRef) int32

// IOSurfaceGetUseCount returns the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetUseCount(_:)
func IOSurfaceGetUseCount(buffer IOSurfaceRef) int32 {
	if _iOSurfaceGetUseCount == nil {
		panic("IOSurface: symbol IOSurfaceGetUseCount not loaded")
	}
	return _iOSurfaceGetUseCount(buffer)
}


var _iOSurfaceGetWidth func(buffer IOSurfaceRef) uintptr

// IOSurfaceGetWidth returns the width of the IOSurface buffer in pixels.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidth(_:)
func IOSurfaceGetWidth(buffer IOSurfaceRef) uintptr {
	if _iOSurfaceGetWidth == nil {
		panic("IOSurface: symbol IOSurfaceGetWidth not loaded")
	}
	return _iOSurfaceGetWidth(buffer)
}


var _iOSurfaceGetWidthOfPlane func(buffer IOSurfaceRef, planeIndex uintptr) uintptr

// IOSurfaceGetWidthOfPlane returns the width of the specified plane (in pixels).
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceGetWidthOfPlane(_:_:)
func IOSurfaceGetWidthOfPlane(buffer IOSurfaceRef, planeIndex uintptr) uintptr {
	if _iOSurfaceGetWidthOfPlane == nil {
		panic("IOSurface: symbol IOSurfaceGetWidthOfPlane not loaded")
	}
	return _iOSurfaceGetWidthOfPlane(buffer, planeIndex)
}


var _iOSurfaceIncrementUseCount func(buffer IOSurfaceRef)

// IOSurfaceIncrementUseCount increments the per-process usage count for an IOSurface.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceIncrementUseCount(_:)
func IOSurfaceIncrementUseCount(buffer IOSurfaceRef) {
	if _iOSurfaceIncrementUseCount == nil {
		panic("IOSurface: symbol IOSurfaceIncrementUseCount not loaded")
	}
	_iOSurfaceIncrementUseCount(buffer)
}


var _iOSurfaceIsInUse func(buffer IOSurfaceRef) bool

// IOSurfaceIsInUse returns true of an IOSurface is in use by any process in the system, otherwise false.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceIsInUse(_:)
func IOSurfaceIsInUse(buffer IOSurfaceRef) bool {
	if _iOSurfaceIsInUse == nil {
		panic("IOSurface: symbol IOSurfaceIsInUse not loaded")
	}
	return _iOSurfaceIsInUse(buffer)
}


var _iOSurfaceLock func(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32

// IOSurfaceLock “Lock” an IOSurface for reading or writing.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLock(_:_:_:)
func IOSurfaceLock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32 {
	if _iOSurfaceLock == nil {
		panic("IOSurface: symbol IOSurfaceLock not loaded")
	}
	return _iOSurfaceLock(buffer, options, seed)
}


var _iOSurfaceLookup func(csid uint32) IOSurfaceRef

// IOSurfaceLookup performs an atomic lookup and retain of an IOSurface by its IOSurfaceID.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookup(_:)
func IOSurfaceLookup(csid uint32) IOSurfaceRef {
	if _iOSurfaceLookup == nil {
		panic("IOSurface: symbol IOSurfaceLookup not loaded")
	}
	return _iOSurfaceLookup(csid)
}


var _iOSurfaceLookupFromMachPort func(port uint32) IOSurfaceRef

// IOSurfaceLookupFromMachPort recreates an IOSurfaceRef from a mach port.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromMachPort(_:)
func IOSurfaceLookupFromMachPort(port uint32) IOSurfaceRef {
	if _iOSurfaceLookupFromMachPort == nil {
		panic("IOSurface: symbol IOSurfaceLookupFromMachPort not loaded")
	}
	return _iOSurfaceLookupFromMachPort(port)
}


var _iOSurfaceLookupFromXPCObject func(xobj unsafe.Pointer) IOSurfaceRef

// IOSurfaceLookupFromXPCObject.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLookupFromXPCObject(_:)
func IOSurfaceLookupFromXPCObject(xobj unsafe.Pointer) IOSurfaceRef {
	if _iOSurfaceLookupFromXPCObject == nil {
		panic("IOSurface: symbol IOSurfaceLookupFromXPCObject not loaded")
	}
	return _iOSurfaceLookupFromXPCObject(xobj)
}


var _iOSurfaceRemoveAllValues func(buffer IOSurfaceRef)

// IOSurfaceRemoveAllValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveAllValues(_:)
func IOSurfaceRemoveAllValues(buffer IOSurfaceRef) {
	if _iOSurfaceRemoveAllValues == nil {
		panic("IOSurface: symbol IOSurfaceRemoveAllValues not loaded")
	}
	_iOSurfaceRemoveAllValues(buffer)
}


var _iOSurfaceRemoveValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef)

// IOSurfaceRemoveValue deletes a value in the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceRemoveValue(_:_:)
func IOSurfaceRemoveValue(buffer IOSurfaceRef, key corefoundation.CFStringRef) {
	if _iOSurfaceRemoveValue == nil {
		panic("IOSurface: symbol IOSurfaceRemoveValue not loaded")
	}
	_iOSurfaceRemoveValue(buffer, key)
}


var _iOSurfaceSetOwnershipIdentity func(buffer IOSurfaceRef, task_id_token uintptr, newLedgerTag int, newLedgerOptions uint32) int32

// IOSurfaceSetOwnershipIdentity.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetOwnershipIdentity(_:_:_:_:)
func IOSurfaceSetOwnershipIdentity(buffer IOSurfaceRef, task_id_token uintptr, newLedgerTag int, newLedgerOptions uint32) int32 {
	if _iOSurfaceSetOwnershipIdentity == nil {
		panic("IOSurface: symbol IOSurfaceSetOwnershipIdentity not loaded")
	}
	return _iOSurfaceSetOwnershipIdentity(buffer, task_id_token, newLedgerTag, newLedgerOptions)
}


var _iOSurfaceSetPurgeable func(buffer IOSurfaceRef, newState uint32, oldState *uint32) int32

// IOSurfaceSetPurgeable.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetPurgeable(_:_:_:)
func IOSurfaceSetPurgeable(buffer IOSurfaceRef, newState uint32, oldState *uint32) int32 {
	if _iOSurfaceSetPurgeable == nil {
		panic("IOSurface: symbol IOSurfaceSetPurgeable not loaded")
	}
	return _iOSurfaceSetPurgeable(buffer, newState, oldState)
}


var _iOSurfaceSetValue func(buffer IOSurfaceRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef)

// IOSurfaceSetValue sets a value in the dictionary associated with the buffer.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValue(_:_:_:)
func IOSurfaceSetValue(buffer IOSurfaceRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) {
	if _iOSurfaceSetValue == nil {
		panic("IOSurface: symbol IOSurfaceSetValue not loaded")
	}
	_iOSurfaceSetValue(buffer, key, value)
}


var _iOSurfaceSetValues func(buffer IOSurfaceRef, keysAndValues corefoundation.CFDictionaryRef)

// IOSurfaceSetValues.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSetValues(_:_:)
func IOSurfaceSetValues(buffer IOSurfaceRef, keysAndValues corefoundation.CFDictionaryRef) {
	if _iOSurfaceSetValues == nil {
		panic("IOSurface: symbol IOSurfaceSetValues not loaded")
	}
	_iOSurfaceSetValues(buffer, keysAndValues)
}


var _iOSurfaceUnlock func(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32

// IOSurfaceUnlock “Unlock” an IOSurface for reading or writing.
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceUnlock(_:_:_:)
func IOSurfaceUnlock(buffer IOSurfaceRef, options IOSurfaceLockOptions, seed *uint32) int32 {
	if _iOSurfaceUnlock == nil {
		panic("IOSurface: symbol IOSurfaceUnlock not loaded")
	}
	return _iOSurfaceUnlock(buffer, options, seed)
}



func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_iOSurfaceAlignProperty, frameworkHandle, "IOSurfaceAlignProperty")
		registerFunc(&_iOSurfaceAllowsPixelSizeCasting, frameworkHandle, "IOSurfaceAllowsPixelSizeCasting")
		registerFunc(&_iOSurfaceCopyAllValues, frameworkHandle, "IOSurfaceCopyAllValues")
		registerFunc(&_iOSurfaceCopyValue, frameworkHandle, "IOSurfaceCopyValue")
		registerFunc(&_iOSurfaceCreate, frameworkHandle, "IOSurfaceCreate")
		registerFunc(&_iOSurfaceCreateMachPort, frameworkHandle, "IOSurfaceCreateMachPort")
		registerFunc(&_iOSurfaceCreateXPCObject, frameworkHandle, "IOSurfaceCreateXPCObject")
		registerFunc(&_iOSurfaceDecrementUseCount, frameworkHandle, "IOSurfaceDecrementUseCount")
		registerFunc(&_iOSurfaceGetAllocSize, frameworkHandle, "IOSurfaceGetAllocSize")
		registerFunc(&_iOSurfaceGetBaseAddress, frameworkHandle, "IOSurfaceGetBaseAddress")
		registerFunc(&_iOSurfaceGetBaseAddressOfPlane, frameworkHandle, "IOSurfaceGetBaseAddressOfPlane")
		registerFunc(&_iOSurfaceGetBitDepthOfComponentOfPlane, frameworkHandle, "IOSurfaceGetBitDepthOfComponentOfPlane")
		registerFunc(&_iOSurfaceGetBitOffsetOfComponentOfPlane, frameworkHandle, "IOSurfaceGetBitOffsetOfComponentOfPlane")
		registerFunc(&_iOSurfaceGetBytesPerElement, frameworkHandle, "IOSurfaceGetBytesPerElement")
		registerFunc(&_iOSurfaceGetBytesPerElementOfPlane, frameworkHandle, "IOSurfaceGetBytesPerElementOfPlane")
		registerFunc(&_iOSurfaceGetBytesPerRow, frameworkHandle, "IOSurfaceGetBytesPerRow")
		registerFunc(&_iOSurfaceGetBytesPerRowOfPlane, frameworkHandle, "IOSurfaceGetBytesPerRowOfPlane")
		registerFunc(&_iOSurfaceGetElementHeight, frameworkHandle, "IOSurfaceGetElementHeight")
		registerFunc(&_iOSurfaceGetElementHeightOfPlane, frameworkHandle, "IOSurfaceGetElementHeightOfPlane")
		registerFunc(&_iOSurfaceGetElementWidth, frameworkHandle, "IOSurfaceGetElementWidth")
		registerFunc(&_iOSurfaceGetElementWidthOfPlane, frameworkHandle, "IOSurfaceGetElementWidthOfPlane")
		registerFunc(&_iOSurfaceGetHeight, frameworkHandle, "IOSurfaceGetHeight")
		registerFunc(&_iOSurfaceGetHeightOfPlane, frameworkHandle, "IOSurfaceGetHeightOfPlane")
		registerFunc(&_iOSurfaceGetID, frameworkHandle, "IOSurfaceGetID")
		registerFunc(&_iOSurfaceGetNameOfComponentOfPlane, frameworkHandle, "IOSurfaceGetNameOfComponentOfPlane")
		registerFunc(&_iOSurfaceGetNumberOfComponentsOfPlane, frameworkHandle, "IOSurfaceGetNumberOfComponentsOfPlane")
		registerFunc(&_iOSurfaceGetPixelFormat, frameworkHandle, "IOSurfaceGetPixelFormat")
		registerFunc(&_iOSurfaceGetPlaneCount, frameworkHandle, "IOSurfaceGetPlaneCount")
		registerFunc(&_iOSurfaceGetPropertyAlignment, frameworkHandle, "IOSurfaceGetPropertyAlignment")
		registerFunc(&_iOSurfaceGetPropertyMaximum, frameworkHandle, "IOSurfaceGetPropertyMaximum")
		registerFunc(&_iOSurfaceGetRangeOfComponentOfPlane, frameworkHandle, "IOSurfaceGetRangeOfComponentOfPlane")
		registerFunc(&_iOSurfaceGetSeed, frameworkHandle, "IOSurfaceGetSeed")
		registerFunc(&_iOSurfaceGetSubsampling, frameworkHandle, "IOSurfaceGetSubsampling")
		registerFunc(&_iOSurfaceGetTypeID, frameworkHandle, "IOSurfaceGetTypeID")
		registerFunc(&_iOSurfaceGetTypeOfComponentOfPlane, frameworkHandle, "IOSurfaceGetTypeOfComponentOfPlane")
		registerFunc(&_iOSurfaceGetUseCount, frameworkHandle, "IOSurfaceGetUseCount")
		registerFunc(&_iOSurfaceGetWidth, frameworkHandle, "IOSurfaceGetWidth")
		registerFunc(&_iOSurfaceGetWidthOfPlane, frameworkHandle, "IOSurfaceGetWidthOfPlane")
		registerFunc(&_iOSurfaceIncrementUseCount, frameworkHandle, "IOSurfaceIncrementUseCount")
		registerFunc(&_iOSurfaceIsInUse, frameworkHandle, "IOSurfaceIsInUse")
		registerFunc(&_iOSurfaceLock, frameworkHandle, "IOSurfaceLock")
		registerFunc(&_iOSurfaceLookup, frameworkHandle, "IOSurfaceLookup")
		registerFunc(&_iOSurfaceLookupFromMachPort, frameworkHandle, "IOSurfaceLookupFromMachPort")
		registerFunc(&_iOSurfaceLookupFromXPCObject, frameworkHandle, "IOSurfaceLookupFromXPCObject")
		registerFunc(&_iOSurfaceRemoveAllValues, frameworkHandle, "IOSurfaceRemoveAllValues")
		registerFunc(&_iOSurfaceRemoveValue, frameworkHandle, "IOSurfaceRemoveValue")
		registerFunc(&_iOSurfaceSetOwnershipIdentity, frameworkHandle, "IOSurfaceSetOwnershipIdentity")
		registerFunc(&_iOSurfaceSetPurgeable, frameworkHandle, "IOSurfaceSetPurgeable")
		registerFunc(&_iOSurfaceSetValue, frameworkHandle, "IOSurfaceSetValue")
		registerFunc(&_iOSurfaceSetValues, frameworkHandle, "IOSurfaceSetValues")
		registerFunc(&_iOSurfaceUnlock, frameworkHandle, "IOSurfaceUnlock")
	}



