// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPipelineFunctionsDescriptor] class.
var (
	_MTLRenderPipelineFunctionsDescriptorClass     MTLRenderPipelineFunctionsDescriptorClass
	_MTLRenderPipelineFunctionsDescriptorClassOnce sync.Once
)

func getMTLRenderPipelineFunctionsDescriptorClass() MTLRenderPipelineFunctionsDescriptorClass {
	_MTLRenderPipelineFunctionsDescriptorClassOnce.Do(func() {
		_MTLRenderPipelineFunctionsDescriptorClass = MTLRenderPipelineFunctionsDescriptorClass{class: objc.GetClass("MTLRenderPipelineFunctionsDescriptor")}
	})
	return _MTLRenderPipelineFunctionsDescriptorClass
}

// GetMTLRenderPipelineFunctionsDescriptorClass returns the class object for MTLRenderPipelineFunctionsDescriptor.
func GetMTLRenderPipelineFunctionsDescriptorClass() MTLRenderPipelineFunctionsDescriptorClass {
	return getMTLRenderPipelineFunctionsDescriptorClass()
}

type MTLRenderPipelineFunctionsDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLRenderPipelineFunctionsDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPipelineFunctionsDescriptorClass) Alloc() MTLRenderPipelineFunctionsDescriptor {
	rv := objc.Send[MTLRenderPipelineFunctionsDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A collection of functions for updating a render pipeline.
//
// # Overview
// 
// When you create a render pipeline that takes visible functions as
// parameters, you need to specify all possible functions that the render
// pipeline can call. If you already have a pipeline, you can create a new
// render pipeline with the same configuration but additional callable
// functions. To create the new pipeline state, configure an
// [MTLRenderPipelineFunctionsDescriptor] instance with the additional
// callable functions to add, and then call the pipeline state’s
// [NewRenderPipelineStateWithAdditionalBinaryFunctionsError] method, passing
// the descriptor.
//
// # Configuring the descriptor’s functions
//
//   - [MTLRenderPipelineFunctionsDescriptor.VertexAdditionalBinaryFunctions]: The vertex functions to add to the render pipeline.
//   - [MTLRenderPipelineFunctionsDescriptor.SetVertexAdditionalBinaryFunctions]
//   - [MTLRenderPipelineFunctionsDescriptor.FragmentAdditionalBinaryFunctions]: The fragment functions to add to the render pipeline.
//   - [MTLRenderPipelineFunctionsDescriptor.SetFragmentAdditionalBinaryFunctions]
//   - [MTLRenderPipelineFunctionsDescriptor.TileAdditionalBinaryFunctions]: The tile functions to add to the render pipeline.
//   - [MTLRenderPipelineFunctionsDescriptor.SetTileAdditionalBinaryFunctions]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor
type MTLRenderPipelineFunctionsDescriptor struct {
	objectivec.Object
}

// MTLRenderPipelineFunctionsDescriptorFromID constructs a [MTLRenderPipelineFunctionsDescriptor] from an objc.ID.
//
// A collection of functions for updating a render pipeline.
func MTLRenderPipelineFunctionsDescriptorFromID(id objc.ID) MTLRenderPipelineFunctionsDescriptor {
	return MTLRenderPipelineFunctionsDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLRenderPipelineFunctionsDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRenderPipelineFunctionsDescriptor] class.
//
// # Configuring the descriptor’s functions
//
//   - [IMTLRenderPipelineFunctionsDescriptor.VertexAdditionalBinaryFunctions]: The vertex functions to add to the render pipeline.
//   - [IMTLRenderPipelineFunctionsDescriptor.SetVertexAdditionalBinaryFunctions]
//   - [IMTLRenderPipelineFunctionsDescriptor.FragmentAdditionalBinaryFunctions]: The fragment functions to add to the render pipeline.
//   - [IMTLRenderPipelineFunctionsDescriptor.SetFragmentAdditionalBinaryFunctions]
//   - [IMTLRenderPipelineFunctionsDescriptor.TileAdditionalBinaryFunctions]: The tile functions to add to the render pipeline.
//   - [IMTLRenderPipelineFunctionsDescriptor.SetTileAdditionalBinaryFunctions]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor
type IMTLRenderPipelineFunctionsDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the descriptor’s functions

	// The vertex functions to add to the render pipeline.
	VertexAdditionalBinaryFunctions() []objectivec.IObject
	SetVertexAdditionalBinaryFunctions(value []objectivec.IObject)
	// The fragment functions to add to the render pipeline.
	FragmentAdditionalBinaryFunctions() []objectivec.IObject
	SetFragmentAdditionalBinaryFunctions(value []objectivec.IObject)
	// The tile functions to add to the render pipeline.
	TileAdditionalBinaryFunctions() []objectivec.IObject
	SetTileAdditionalBinaryFunctions(value []objectivec.IObject)
}

// Init initializes the instance.
func (r MTLRenderPipelineFunctionsDescriptor) Init() MTLRenderPipelineFunctionsDescriptor {
	rv := objc.Send[MTLRenderPipelineFunctionsDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPipelineFunctionsDescriptor) Autorelease() MTLRenderPipelineFunctionsDescriptor {
	rv := objc.Send[MTLRenderPipelineFunctionsDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPipelineFunctionsDescriptor creates a new MTLRenderPipelineFunctionsDescriptor instance.
func NewMTLRenderPipelineFunctionsDescriptor() MTLRenderPipelineFunctionsDescriptor {
	class := getMTLRenderPipelineFunctionsDescriptorClass()
	rv := objc.Send[MTLRenderPipelineFunctionsDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The vertex functions to add to the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (r MTLRenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertexAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// The fragment functions to add to the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (r MTLRenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setFragmentAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// The tile functions to add to the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineFunctionsDescriptor/tileAdditionalBinaryFunctions
func (r MTLRenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setTileAdditionalBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}

