// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexDescriptor] class.
var (
	_MTLVertexDescriptorClass     MTLVertexDescriptorClass
	_MTLVertexDescriptorClassOnce sync.Once
)

func getMTLVertexDescriptorClass() MTLVertexDescriptorClass {
	_MTLVertexDescriptorClassOnce.Do(func() {
		_MTLVertexDescriptorClass = MTLVertexDescriptorClass{class: objc.GetClass("MTLVertexDescriptor")}
	})
	return _MTLVertexDescriptorClass
}

// GetMTLVertexDescriptorClass returns the class object for MTLVertexDescriptor.
func GetMTLVertexDescriptorClass() MTLVertexDescriptorClass {
	return getMTLVertexDescriptorClass()
}

type MTLVertexDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLVertexDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexDescriptorClass) Alloc() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance that describes how to organize and map data to a vertex
// function.
//
// # Overview
// 
// An [MTLVertexDescriptor] instance is used to configure how vertex data
// stored in memory is mapped to attributes in a vertex shader.
// 
// A pipeline state is the state of the graphics rendering pipeline, including
// shaders, blending, multisampling, and visibility testing. For every
// pipeline state, there can be only one [MTLVertexDescriptor] instance. When
// you configure an [MTLRenderPipelineDescriptor] instance to create this
// pipeline state, you use an [MTLVertexDescriptor] instance to establish the
// vertex layout for the function associated with the pipeline. Create and
// configure an [MTLVertexDescriptor] instance, then use this instance to set
// the [VertexDescriptor] property of the [MTLRenderPipelineDescriptor]
// instance.
//
// # Setting default values
//
//   - [MTLVertexDescriptor.Reset]: Resets the default state for the vertex descriptor.
//
// # Accessing the vertex buffer layouts and vertex attributes
//
//   - [MTLVertexDescriptor.Attributes]: An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
//   - [MTLVertexDescriptor.Layouts]: An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor
type MTLVertexDescriptor struct {
	objectivec.Object
}

// MTLVertexDescriptorFromID constructs a [MTLVertexDescriptor] from an objc.ID.
//
// An instance that describes how to organize and map data to a vertex
// function.
func MTLVertexDescriptorFromID(id objc.ID) MTLVertexDescriptor {
	return MTLVertexDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLVertexDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVertexDescriptor] class.
//
// # Setting default values
//
//   - [IMTLVertexDescriptor.Reset]: Resets the default state for the vertex descriptor.
//
// # Accessing the vertex buffer layouts and vertex attributes
//
//   - [IMTLVertexDescriptor.Attributes]: An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
//   - [IMTLVertexDescriptor.Layouts]: An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor
type IMTLVertexDescriptor interface {
	objectivec.IObject

	// Topic: Setting default values

	// Resets the default state for the vertex descriptor.
	Reset()

	// Topic: Accessing the vertex buffer layouts and vertex attributes

	// An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
	Attributes() IMTLVertexAttributeDescriptorArray
	// An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
	Layouts() IMTLVertexBufferLayoutDescriptorArray

	MTLBufferLayoutStrideDynamic() int
	// The organization of vertex data in an attribute’s argument table.
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)
}

// Init initializes the instance.
func (v MTLVertexDescriptor) Init() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexDescriptor) Autorelease() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexDescriptor creates a new MTLVertexDescriptor instance.
func NewMTLVertexDescriptor() MTLVertexDescriptor {
	class := getMTLVertexDescriptorClass()
	rv := objc.Send[MTLVertexDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets the default state for the vertex descriptor.
//
// # Discussion
// 
// After reset, each element of the [Attributes] array has a default vertex
// attribute descriptor, and each element of the [Layouts] array has a default
// vertex buffer layout descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/reset()
func (v MTLVertexDescriptor) Reset() {
	objc.Send[objc.ID](v.ID, objc.Sel("reset"))
}

// Creates and returns a new vertex descriptor.
//
// # Return Value
// 
// A default object with allocated arrays in the [Attributes] and [Layouts]
// properties.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/vertexDescriptor
func (_MTLVertexDescriptorClass MTLVertexDescriptorClass) VertexDescriptor() MTLVertexDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLVertexDescriptorClass.class), objc.Sel("vertexDescriptor"))
	return MTLVertexDescriptorFromID(rv)
}

// An array of state data that describes how vertex attribute data is stored
// in memory and is mapped to arguments for a vertex shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/attributes
func (v MTLVertexDescriptor) Attributes() IMTLVertexAttributeDescriptorArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("attributes"))
	return MTLVertexAttributeDescriptorArrayFromID(objc.ID(rv))
}
// An array of state data that describes how data are fetched by a vertex
// shader function when rendering primitives.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/layouts
func (v MTLVertexDescriptor) Layouts() IMTLVertexBufferLayoutDescriptorArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layouts"))
	return MTLVertexBufferLayoutDescriptorArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v MTLVertexDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}
// The organization of vertex data in an attribute’s argument table.
//
// See: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/vertexdescriptor
func (v MTLVertexDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("vertexDescriptor"))
	return MTLVertexDescriptorFromID(objc.ID(rv))
}
func (v MTLVertexDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[struct{}](v.ID, objc.Sel("setVertexDescriptor:"), value)
}

