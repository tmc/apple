// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLComputePipelineReflection] class.
var (
	_MTLComputePipelineReflectionClass     MTLComputePipelineReflectionClass
	_MTLComputePipelineReflectionClassOnce sync.Once
)

func getMTLComputePipelineReflectionClass() MTLComputePipelineReflectionClass {
	_MTLComputePipelineReflectionClassOnce.Do(func() {
		_MTLComputePipelineReflectionClass = MTLComputePipelineReflectionClass{class: objc.GetClass("MTLComputePipelineReflection")}
	})
	return _MTLComputePipelineReflectionClass
}

// GetMTLComputePipelineReflectionClass returns the class object for MTLComputePipelineReflection.
func GetMTLComputePipelineReflectionClass() MTLComputePipelineReflectionClass {
	return getMTLComputePipelineReflectionClass()
}

type MTLComputePipelineReflectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLComputePipelineReflectionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLComputePipelineReflectionClass) Alloc() MTLComputePipelineReflection {
	rv := objc.Send[MTLComputePipelineReflection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Information about the arguments of a compute function.
//
// # Overview
// 
// An [MTLComputePipelineReflection] object provides access to the arguments
// of the compute function used in an [MTLComputePipelineState] object. An
// [MTLComputePipelineReflection] object can be created along with an
// [MTLComputePipelineState] object. Don’t create an
// [MTLComputePipelineReflection] object directly. Instead, call either the
// [NewComputePipelineStateWithFunctionOptionsReflectionError] or
// [NewComputePipelineStateWithFunctionOptionsCompletionHandler] method of
// [MTLDevice] to create both an [MTLComputePipelineState] object and an
// [MTLComputePipelineReflection] object.
// 
// [MTLComputePipelineReflection] objects can use a significant amount of
// memory; release any strong references to them after you finish creating
// pipeline objects.
//
// # Instance Properties
//
//   - [MTLComputePipelineReflection.Bindings]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection
type MTLComputePipelineReflection struct {
	objectivec.Object
}

// MTLComputePipelineReflectionFromID constructs a [MTLComputePipelineReflection] from an objc.ID.
//
// Information about the arguments of a compute function.
func MTLComputePipelineReflectionFromID(id objc.ID) MTLComputePipelineReflection {
	return MTLComputePipelineReflection{objectivec.Object{ID: id}}
}
// NOTE: MTLComputePipelineReflection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLComputePipelineReflection] class.
//
// # Instance Properties
//
//   - [IMTLComputePipelineReflection.Bindings]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection
type IMTLComputePipelineReflection interface {
	objectivec.IObject

	// Topic: Instance Properties

	Bindings() []objectivec.IObject
}

// Init initializes the instance.
func (c MTLComputePipelineReflection) Init() MTLComputePipelineReflection {
	rv := objc.Send[MTLComputePipelineReflection](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLComputePipelineReflection) Autorelease() MTLComputePipelineReflection {
	rv := objc.Send[MTLComputePipelineReflection](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLComputePipelineReflection creates a new MTLComputePipelineReflection instance.
func NewMTLComputePipelineReflection() MTLComputePipelineReflection {
	class := getMTLComputePipelineReflectionClass()
	rv := objc.Send[MTLComputePipelineReflection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection/bindings
func (c MTLComputePipelineReflection) Bindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("bindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

