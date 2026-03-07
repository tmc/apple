// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: MetalKit: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: MetalKit: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: MetalKit: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _mTKMetalVertexDescriptorFromModelIO func(modelIODescriptor uintptr) *metal.MTLVertexDescriptor

// MTKMetalVertexDescriptorFromModelIO returns a partially converted Metal vertex descriptor.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIO(_:)
func MTKMetalVertexDescriptorFromModelIO(modelIODescriptor uintptr) *metal.MTLVertexDescriptor {
	if _mTKMetalVertexDescriptorFromModelIO == nil {
		panic("MetalKit: symbol MTKMetalVertexDescriptorFromModelIO not loaded")
	}
	return _mTKMetalVertexDescriptorFromModelIO(modelIODescriptor)
}


var _mTKMetalVertexDescriptorFromModelIOWithError func(modelIODescriptor uintptr, err *foundation.NSError) *metal.MTLVertexDescriptor

// MTKMetalVertexDescriptorFromModelIOWithError returns a partially converted Metal vertex descriptor, reporting any error that occurs.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIOWithError
func MTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor uintptr, err *foundation.NSError) *metal.MTLVertexDescriptor {
	if _mTKMetalVertexDescriptorFromModelIOWithError == nil {
		panic("MetalKit: symbol MTKMetalVertexDescriptorFromModelIOWithError not loaded")
	}
	return _mTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor, err)
}



var _mTKModelIOVertexDescriptorFromMetal func(metalDescriptor *metal.MTLVertexDescriptor) *objc.ID

// MTKModelIOVertexDescriptorFromMetal returns a partially converted Model I/O vertex descriptor.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetal(_:)
func MTKModelIOVertexDescriptorFromMetal(metalDescriptor *metal.MTLVertexDescriptor) *objc.ID {
	if _mTKModelIOVertexDescriptorFromMetal == nil {
		panic("MetalKit: symbol MTKModelIOVertexDescriptorFromMetal not loaded")
	}
	return _mTKModelIOVertexDescriptorFromMetal(metalDescriptor)
}


var _mTKModelIOVertexDescriptorFromMetalWithError func(metalDescriptor *metal.MTLVertexDescriptor, err *foundation.NSError) *objc.ID

// MTKModelIOVertexDescriptorFromMetalWithError returns a partially converted Model I/O vertex descriptor, reporting any error that occurs.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetalWithError
func MTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor *metal.MTLVertexDescriptor, err *foundation.NSError) *objc.ID {
	if _mTKModelIOVertexDescriptorFromMetalWithError == nil {
		panic("MetalKit: symbol MTKModelIOVertexDescriptorFromMetalWithError not loaded")
	}
	return _mTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor, err)
}




func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_mTKMetalVertexDescriptorFromModelIO, frameworkHandle, "MTKMetalVertexDescriptorFromModelIO")
		registerFunc(&_mTKMetalVertexDescriptorFromModelIOWithError, frameworkHandle, "MTKMetalVertexDescriptorFromModelIOWithError")
		registerFunc(&_mTKModelIOVertexDescriptorFromMetal, frameworkHandle, "MTKModelIOVertexDescriptorFromMetal")
		registerFunc(&_mTKModelIOVertexDescriptorFromMetalWithError, frameworkHandle, "MTKModelIOVertexDescriptorFromMetalWithError")
	}


