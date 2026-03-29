// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4PipelineStageDynamicLinkingDescriptor] class.
var (
	_MTL4PipelineStageDynamicLinkingDescriptorClass     MTL4PipelineStageDynamicLinkingDescriptorClass
	_MTL4PipelineStageDynamicLinkingDescriptorClassOnce sync.Once
)

func getMTL4PipelineStageDynamicLinkingDescriptorClass() MTL4PipelineStageDynamicLinkingDescriptorClass {
	_MTL4PipelineStageDynamicLinkingDescriptorClassOnce.Do(func() {
		_MTL4PipelineStageDynamicLinkingDescriptorClass = MTL4PipelineStageDynamicLinkingDescriptorClass{class: objc.GetClass("MTL4PipelineStageDynamicLinkingDescriptor")}
	})
	return _MTL4PipelineStageDynamicLinkingDescriptorClass
}

// GetMTL4PipelineStageDynamicLinkingDescriptorClass returns the class object for MTL4PipelineStageDynamicLinkingDescriptor.
func GetMTL4PipelineStageDynamicLinkingDescriptorClass() MTL4PipelineStageDynamicLinkingDescriptorClass {
	return getMTL4PipelineStageDynamicLinkingDescriptorClass()
}

type MTL4PipelineStageDynamicLinkingDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4PipelineStageDynamicLinkingDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4PipelineStageDynamicLinkingDescriptorClass) Alloc() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties to drive the dynamic linking process of a
// pipeline stage.
//
// # Instance Properties
//
//   - [MTL4PipelineStageDynamicLinkingDescriptor.BinaryLinkedFunctions]: Provides the array of binary functions to link.
//   - [MTL4PipelineStageDynamicLinkingDescriptor.SetBinaryLinkedFunctions]
//   - [MTL4PipelineStageDynamicLinkingDescriptor.MaxCallStackDepth]: Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//   - [MTL4PipelineStageDynamicLinkingDescriptor.SetMaxCallStackDepth]
//   - [MTL4PipelineStageDynamicLinkingDescriptor.PreloadedLibraries]: Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
//   - [MTL4PipelineStageDynamicLinkingDescriptor.SetPreloadedLibraries]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor
type MTL4PipelineStageDynamicLinkingDescriptor struct {
	objectivec.Object
}

// MTL4PipelineStageDynamicLinkingDescriptorFromID constructs a [MTL4PipelineStageDynamicLinkingDescriptor] from an objc.ID.
//
// Groups together properties to drive the dynamic linking process of a
// pipeline stage.
func MTL4PipelineStageDynamicLinkingDescriptorFromID(id objc.ID) MTL4PipelineStageDynamicLinkingDescriptor {
	return MTL4PipelineStageDynamicLinkingDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4PipelineStageDynamicLinkingDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4PipelineStageDynamicLinkingDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.BinaryLinkedFunctions]: Provides the array of binary functions to link.
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.SetBinaryLinkedFunctions]
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.MaxCallStackDepth]: Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.SetMaxCallStackDepth]
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.PreloadedLibraries]: Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
//   - [IMTL4PipelineStageDynamicLinkingDescriptor.SetPreloadedLibraries]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor
type IMTL4PipelineStageDynamicLinkingDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides the array of binary functions to link.
	BinaryLinkedFunctions() []objectivec.IObject
	SetBinaryLinkedFunctions(value []objectivec.IObject)
	// Limits the maximum depth of the call stack for indirect function calls in the pipeline stage function.
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	// Provides an array of dynamic libraries the compiler loads when it builds the pipeline.
	PreloadedLibraries() []objectivec.IObject
	SetPreloadedLibraries(value []objectivec.IObject)
}

// Init initializes the instance.
func (m MTL4PipelineStageDynamicLinkingDescriptor) Init() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4PipelineStageDynamicLinkingDescriptor) Autorelease() MTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineStageDynamicLinkingDescriptor creates a new MTL4PipelineStageDynamicLinkingDescriptor instance.
func NewMTL4PipelineStageDynamicLinkingDescriptor() MTL4PipelineStageDynamicLinkingDescriptor {
	class := getMTL4PipelineStageDynamicLinkingDescriptorClass()
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides the array of binary functions to link.
//
// # Discussion
// 
// Binary functions are shader functions that you compile from Metal IR to
// machine code ahead of time using instances of [MTL4Compiler].
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/binaryLinkedFunctions
func (m MTL4PipelineStageDynamicLinkingDescriptor) BinaryLinkedFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("binaryLinkedFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4PipelineStageDynamicLinkingDescriptor) SetBinaryLinkedFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setBinaryLinkedFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Limits the maximum depth of the call stack for indirect function calls in
// the pipeline stage function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/maxCallStackDepth
func (m MTL4PipelineStageDynamicLinkingDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxCallStackDepth"))
	return rv
}
func (m MTL4PipelineStageDynamicLinkingDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxCallStackDepth:"), value)
}
// Provides an array of dynamic libraries the compiler loads when it builds
// the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineStageDynamicLinkingDescriptor/preloadedLibraries
func (m MTL4PipelineStageDynamicLinkingDescriptor) PreloadedLibraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("preloadedLibraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4PipelineStageDynamicLinkingDescriptor) SetPreloadedLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreloadedLibraries:"), objectivec.IObjectSliceToNSArray(value))
}

