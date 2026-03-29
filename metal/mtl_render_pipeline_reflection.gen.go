// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPipelineReflection] class.
var (
	_MTLRenderPipelineReflectionClass     MTLRenderPipelineReflectionClass
	_MTLRenderPipelineReflectionClassOnce sync.Once
)

func getMTLRenderPipelineReflectionClass() MTLRenderPipelineReflectionClass {
	_MTLRenderPipelineReflectionClassOnce.Do(func() {
		_MTLRenderPipelineReflectionClass = MTLRenderPipelineReflectionClass{class: objc.GetClass("MTLRenderPipelineReflection")}
	})
	return _MTLRenderPipelineReflectionClass
}

// GetMTLRenderPipelineReflectionClass returns the class object for MTLRenderPipelineReflection.
func GetMTLRenderPipelineReflectionClass() MTLRenderPipelineReflectionClass {
	return getMTLRenderPipelineReflectionClass()
}

type MTLRenderPipelineReflectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLRenderPipelineReflectionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPipelineReflectionClass) Alloc() MTLRenderPipelineReflection {
	rv := objc.Send[MTLRenderPipelineReflection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Information about the arguments of a graphics function.
//
// # Overview
// 
// The [MTLRenderPipelineReflection] class is an interface that represents the
// parameters for the shaders in a render pipeline state (see
// [MTLRenderPipelineState]). Each pipeline state can include object, mesh,
// vertex, fragment, and tile shaders.
// 
// You create a reflection instance at the same time as the pipeline state
// that it represents by calling the appropriate [MTLDevice] method. For
// example, the [NewRenderPipelineStateWithDescriptorOptionsReflectionError]
// and [NewRenderPipelineStateWithDescriptorOptionsCompletionHandler] methods
// create the pipeline state and the reflection instances at the same time.
// 
// For more information, see [Pipeline state creation].
//
// [Pipeline state creation]: https://developer.apple.com/documentation/Metal/pipeline-state-creation
//
// # Inspecting a shader’s parameter
//
//   - [MTLRenderPipelineReflection.FragmentBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s fragment shader.
//   - [MTLRenderPipelineReflection.MeshBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s mesh shader.
//   - [MTLRenderPipelineReflection.ObjectBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s object shader.
//   - [MTLRenderPipelineReflection.TileBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s tile shader.
//   - [MTLRenderPipelineReflection.VertexBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s vertex shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection
type MTLRenderPipelineReflection struct {
	objectivec.Object
}

// MTLRenderPipelineReflectionFromID constructs a [MTLRenderPipelineReflection] from an objc.ID.
//
// Information about the arguments of a graphics function.
func MTLRenderPipelineReflectionFromID(id objc.ID) MTLRenderPipelineReflection {
	return MTLRenderPipelineReflection{objectivec.Object{ID: id}}
}
// NOTE: MTLRenderPipelineReflection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRenderPipelineReflection] class.
//
// # Inspecting a shader’s parameter
//
//   - [IMTLRenderPipelineReflection.FragmentBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s fragment shader.
//   - [IMTLRenderPipelineReflection.MeshBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s mesh shader.
//   - [IMTLRenderPipelineReflection.ObjectBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s object shader.
//   - [IMTLRenderPipelineReflection.TileBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s tile shader.
//   - [IMTLRenderPipelineReflection.VertexBindings]: An array of binding instances, each of which represents a parameter of the pipeline state’s vertex shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection
type IMTLRenderPipelineReflection interface {
	objectivec.IObject

	// Topic: Inspecting a shader’s parameter

	// An array of binding instances, each of which represents a parameter of the pipeline state’s fragment shader.
	FragmentBindings() []objectivec.IObject
	// An array of binding instances, each of which represents a parameter of the pipeline state’s mesh shader.
	MeshBindings() []objectivec.IObject
	// An array of binding instances, each of which represents a parameter of the pipeline state’s object shader.
	ObjectBindings() []objectivec.IObject
	// An array of binding instances, each of which represents a parameter of the pipeline state’s tile shader.
	TileBindings() []objectivec.IObject
	// An array of binding instances, each of which represents a parameter of the pipeline state’s vertex shader.
	VertexBindings() []objectivec.IObject
}

// Init initializes the instance.
func (r MTLRenderPipelineReflection) Init() MTLRenderPipelineReflection {
	rv := objc.Send[MTLRenderPipelineReflection](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPipelineReflection) Autorelease() MTLRenderPipelineReflection {
	rv := objc.Send[MTLRenderPipelineReflection](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPipelineReflection creates a new MTLRenderPipelineReflection instance.
func NewMTLRenderPipelineReflection() MTLRenderPipelineReflection {
	class := getMTLRenderPipelineReflectionClass()
	rv := objc.Send[MTLRenderPipelineReflection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of binding instances, each of which represents a parameter of the
// pipeline state’s fragment shader.
//
// # Discussion
// 
// The [MTLBinding] elements in the array are in the same order as the
// fragment shader’s declaration signature.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/fragmentBindings
func (r MTLRenderPipelineReflection) FragmentBindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("fragmentBindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// An array of binding instances, each of which represents a parameter of the
// pipeline state’s mesh shader.
//
// # Discussion
// 
// The [MTLBinding] elements in the array are in the same order as the mesh
// shader’s declaration signature.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/meshBindings
func (r MTLRenderPipelineReflection) MeshBindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("meshBindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// An array of binding instances, each of which represents a parameter of the
// pipeline state’s object shader.
//
// # Discussion
// 
// The [MTLBinding] elements in the array are in the same order as the object
// shader’s declaration signature.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/objectBindings
func (r MTLRenderPipelineReflection) ObjectBindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("objectBindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// An array of binding instances, each of which represents a parameter of the
// pipeline state’s tile shader.
//
// # Discussion
// 
// The [MTLBinding] elements in the array are in the same order as the tile
// shader’s declaration signature.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/tileBindings
func (r MTLRenderPipelineReflection) TileBindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("tileBindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// An array of binding instances, each of which represents a parameter of the
// pipeline state’s vertex shader.
//
// # Discussion
// 
// The [MTLBinding] elements in the array are in the same order as the vertex
// shader’s declaration signature.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineReflection/vertexBindings
func (r MTLRenderPipelineReflection) VertexBindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("vertexBindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

