// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
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
		return fmt.Sprintf("MetalKit: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("MetalKit: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("MetalKit: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("MetalKit: register symbol %s: %v", name, r)
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

var _mTKMetalVertexDescriptorFromModelIO func(modelIODescriptor uintptr) *metal.MTLVertexDescriptor
var _mTKMetalVertexDescriptorFromModelIOErr error

func tryMTKMetalVertexDescriptorFromModelIO(modelIODescriptor uintptr) (*metal.MTLVertexDescriptor, error) {
	if _mTKMetalVertexDescriptorFromModelIO == nil {
		return nil, symbolCallError("MTKMetalVertexDescriptorFromModelIO", "10.11", _mTKMetalVertexDescriptorFromModelIOErr)
	}
	return _mTKMetalVertexDescriptorFromModelIO(modelIODescriptor), nil
}

// MTKMetalVertexDescriptorFromModelIO returns a partially converted Metal vertex descriptor.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIO(_:)
func MTKMetalVertexDescriptorFromModelIO(modelIODescriptor uintptr) *metal.MTLVertexDescriptor {
	result, callErr := tryMTKMetalVertexDescriptorFromModelIO(modelIODescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTKMetalVertexDescriptorFromModelIOWithError func(modelIODescriptor uintptr, err *foundation.NSError) *metal.MTLVertexDescriptor
var _mTKMetalVertexDescriptorFromModelIOWithErrorErr error

func tryMTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor uintptr, err *foundation.NSError) (*metal.MTLVertexDescriptor, error) {
	if _mTKMetalVertexDescriptorFromModelIOWithError == nil {
		return nil, symbolCallError("MTKMetalVertexDescriptorFromModelIOWithError", "10.12", _mTKMetalVertexDescriptorFromModelIOWithErrorErr)
	}
	return _mTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor, err), nil
}

// MTKMetalVertexDescriptorFromModelIOWithError returns a partially converted Metal vertex descriptor, reporting any error that occurs.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexDescriptorFromModelIOWithError
func MTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor uintptr, err *foundation.NSError) *metal.MTLVertexDescriptor {
	result, callErr := tryMTKMetalVertexDescriptorFromModelIOWithError(modelIODescriptor, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTKMetalVertexFormatFromModelIO func(vertexFormat uint) metal.MTLVertexFormat
var _mTKMetalVertexFormatFromModelIOErr error

func tryMTKMetalVertexFormatFromModelIO(vertexFormat uint) (metal.MTLVertexFormat, error) {
	if _mTKMetalVertexFormatFromModelIO == nil {
		return *new(metal.MTLVertexFormat), symbolCallError("MTKMetalVertexFormatFromModelIO", "10.11", _mTKMetalVertexFormatFromModelIOErr)
	}
	return _mTKMetalVertexFormatFromModelIO(vertexFormat), nil
}

// MTKMetalVertexFormatFromModelIO returns a converted Metal vertex format.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMetalVertexFormatFromModelIO(_:)
func MTKMetalVertexFormatFromModelIO(vertexFormat uint) metal.MTLVertexFormat {
	result, callErr := tryMTKMetalVertexFormatFromModelIO(vertexFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTKModelIOVertexDescriptorFromMetal func(metalDescriptor *metal.MTLVertexDescriptor) *objc.ID
var _mTKModelIOVertexDescriptorFromMetalErr error

func tryMTKModelIOVertexDescriptorFromMetal(metalDescriptor *metal.MTLVertexDescriptor) (*objc.ID, error) {
	if _mTKModelIOVertexDescriptorFromMetal == nil {
		return nil, symbolCallError("MTKModelIOVertexDescriptorFromMetal", "10.11", _mTKModelIOVertexDescriptorFromMetalErr)
	}
	return _mTKModelIOVertexDescriptorFromMetal(metalDescriptor), nil
}

// MTKModelIOVertexDescriptorFromMetal returns a partially converted Model I/O vertex descriptor.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetal(_:)
func MTKModelIOVertexDescriptorFromMetal(metalDescriptor *metal.MTLVertexDescriptor) *objc.ID {
	result, callErr := tryMTKModelIOVertexDescriptorFromMetal(metalDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTKModelIOVertexDescriptorFromMetalWithError func(metalDescriptor *metal.MTLVertexDescriptor, err *foundation.NSError) *objc.ID
var _mTKModelIOVertexDescriptorFromMetalWithErrorErr error

func tryMTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor *metal.MTLVertexDescriptor, err *foundation.NSError) (*objc.ID, error) {
	if _mTKModelIOVertexDescriptorFromMetalWithError == nil {
		return nil, symbolCallError("MTKModelIOVertexDescriptorFromMetalWithError", "10.12", _mTKModelIOVertexDescriptorFromMetalWithErrorErr)
	}
	return _mTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor, err), nil
}

// MTKModelIOVertexDescriptorFromMetalWithError returns a partially converted Model I/O vertex descriptor, reporting any error that occurs.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexDescriptorFromMetalWithError
func MTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor *metal.MTLVertexDescriptor, err *foundation.NSError) *objc.ID {
	result, callErr := tryMTKModelIOVertexDescriptorFromMetalWithError(metalDescriptor, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTKModelIOVertexFormatFromMetal func(vertexFormat metal.MTLVertexFormat) uint
var _mTKModelIOVertexFormatFromMetalErr error

func tryMTKModelIOVertexFormatFromMetal(vertexFormat metal.MTLVertexFormat) (uint, error) {
	if _mTKModelIOVertexFormatFromMetal == nil {
		return 0, symbolCallError("MTKModelIOVertexFormatFromMetal", "10.11", _mTKModelIOVertexFormatFromMetalErr)
	}
	return _mTKModelIOVertexFormatFromMetal(vertexFormat), nil
}

// MTKModelIOVertexFormatFromMetal returns a converted Model I/O vertex format.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelIOVertexFormatFromMetal(_:)
func MTKModelIOVertexFormatFromMetal(vertexFormat metal.MTLVertexFormat) uint {
	result, callErr := tryMTKModelIOVertexFormatFromMetal(vertexFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mTKMetalVertexDescriptorFromModelIO, &_mTKMetalVertexDescriptorFromModelIOErr, frameworkHandle, "MTKMetalVertexDescriptorFromModelIO", "10.11")
	registerFunc(&_mTKMetalVertexDescriptorFromModelIOWithError, &_mTKMetalVertexDescriptorFromModelIOWithErrorErr, frameworkHandle, "MTKMetalVertexDescriptorFromModelIOWithError", "10.12")
	registerFunc(&_mTKMetalVertexFormatFromModelIO, &_mTKMetalVertexFormatFromModelIOErr, frameworkHandle, "MTKMetalVertexFormatFromModelIO", "10.11")
	registerFunc(&_mTKModelIOVertexDescriptorFromMetal, &_mTKModelIOVertexDescriptorFromMetalErr, frameworkHandle, "MTKModelIOVertexDescriptorFromMetal", "10.11")
	registerFunc(&_mTKModelIOVertexDescriptorFromMetalWithError, &_mTKModelIOVertexDescriptorFromMetalWithErrorErr, frameworkHandle, "MTKModelIOVertexDescriptorFromMetalWithError", "10.12")
	registerFunc(&_mTKModelIOVertexFormatFromMetal, &_mTKModelIOVertexFormatFromMetalErr, frameworkHandle, "MTKModelIOVertexFormatFromMetal", "10.11")
}
