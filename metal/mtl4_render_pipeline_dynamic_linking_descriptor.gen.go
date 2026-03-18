// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4RenderPipelineDynamicLinkingDescriptor] class.
var (
	_MTL4RenderPipelineDynamicLinkingDescriptorClass     MTL4RenderPipelineDynamicLinkingDescriptorClass
	_MTL4RenderPipelineDynamicLinkingDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineDynamicLinkingDescriptorClass() MTL4RenderPipelineDynamicLinkingDescriptorClass {
	_MTL4RenderPipelineDynamicLinkingDescriptorClassOnce.Do(func() {
		_MTL4RenderPipelineDynamicLinkingDescriptorClass = MTL4RenderPipelineDynamicLinkingDescriptorClass{class: objc.GetClass("MTL4RenderPipelineDynamicLinkingDescriptor")}
	})
	return _MTL4RenderPipelineDynamicLinkingDescriptorClass
}

// GetMTL4RenderPipelineDynamicLinkingDescriptorClass returns the class object for MTL4RenderPipelineDynamicLinkingDescriptor.
func GetMTL4RenderPipelineDynamicLinkingDescriptorClass() MTL4RenderPipelineDynamicLinkingDescriptorClass {
	return getMTL4RenderPipelineDynamicLinkingDescriptorClass()
}

type MTL4RenderPipelineDynamicLinkingDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPipelineDynamicLinkingDescriptorClass) Alloc() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties that provide linking properties for render
// pipelines.
//
// # Instance Properties
//
//   - [MTL4RenderPipelineDynamicLinkingDescriptor.FragmentLinkingDescriptor]: Controls properties for linking the fragment stage of the render pipeline.
//   - [MTL4RenderPipelineDynamicLinkingDescriptor.MeshLinkingDescriptor]: Controls properties for linking the mesh stage of the render pipeline.
//   - [MTL4RenderPipelineDynamicLinkingDescriptor.ObjectLinkingDescriptor]: Controls properties for link the object stage of the render pipeline.
//   - [MTL4RenderPipelineDynamicLinkingDescriptor.TileLinkingDescriptor]: Controls properties for linking the tile stage of the render pipeline.
//   - [MTL4RenderPipelineDynamicLinkingDescriptor.VertexLinkingDescriptor]: Controls properties for linking the vertex stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor
type MTL4RenderPipelineDynamicLinkingDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineDynamicLinkingDescriptorFromID constructs a [MTL4RenderPipelineDynamicLinkingDescriptor] from an objc.ID.
//
// Groups together properties that provide linking properties for render
// pipelines.
func MTL4RenderPipelineDynamicLinkingDescriptorFromID(id objc.ID) MTL4RenderPipelineDynamicLinkingDescriptor {
	return MTL4RenderPipelineDynamicLinkingDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4RenderPipelineDynamicLinkingDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPipelineDynamicLinkingDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4RenderPipelineDynamicLinkingDescriptor.FragmentLinkingDescriptor]: Controls properties for linking the fragment stage of the render pipeline.
//   - [IMTL4RenderPipelineDynamicLinkingDescriptor.MeshLinkingDescriptor]: Controls properties for linking the mesh stage of the render pipeline.
//   - [IMTL4RenderPipelineDynamicLinkingDescriptor.ObjectLinkingDescriptor]: Controls properties for link the object stage of the render pipeline.
//   - [IMTL4RenderPipelineDynamicLinkingDescriptor.TileLinkingDescriptor]: Controls properties for linking the tile stage of the render pipeline.
//   - [IMTL4RenderPipelineDynamicLinkingDescriptor.VertexLinkingDescriptor]: Controls properties for linking the vertex stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor
type IMTL4RenderPipelineDynamicLinkingDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Controls properties for linking the fragment stage of the render pipeline.
	FragmentLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	// Controls properties for linking the mesh stage of the render pipeline.
	MeshLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	// Controls properties for link the object stage of the render pipeline.
	ObjectLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	// Controls properties for linking the tile stage of the render pipeline.
	TileLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	// Controls properties for linking the vertex stage of the render pipeline.
	VertexLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
}

// Init initializes the instance.
func (m MTL4RenderPipelineDynamicLinkingDescriptor) Init() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPipelineDynamicLinkingDescriptor) Autorelease() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineDynamicLinkingDescriptor creates a new MTL4RenderPipelineDynamicLinkingDescriptor instance.
func NewMTL4RenderPipelineDynamicLinkingDescriptor() MTL4RenderPipelineDynamicLinkingDescriptor {
	class := getMTL4RenderPipelineDynamicLinkingDescriptorClass()
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Controls properties for linking the fragment stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/fragmentLinkingDescriptor
func (m MTL4RenderPipelineDynamicLinkingDescriptor) FragmentLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentLinkingDescriptor"))
	return MTL4PipelineStageDynamicLinkingDescriptorFromID(objc.ID(rv))
}

// Controls properties for linking the mesh stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/meshLinkingDescriptor
func (m MTL4RenderPipelineDynamicLinkingDescriptor) MeshLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshLinkingDescriptor"))
	return MTL4PipelineStageDynamicLinkingDescriptorFromID(objc.ID(rv))
}

// Controls properties for link the object stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/objectLinkingDescriptor
func (m MTL4RenderPipelineDynamicLinkingDescriptor) ObjectLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectLinkingDescriptor"))
	return MTL4PipelineStageDynamicLinkingDescriptorFromID(objc.ID(rv))
}

// Controls properties for linking the tile stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/tileLinkingDescriptor
func (m MTL4RenderPipelineDynamicLinkingDescriptor) TileLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("tileLinkingDescriptor"))
	return MTL4PipelineStageDynamicLinkingDescriptorFromID(objc.ID(rv))
}

// Controls properties for linking the vertex stage of the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/vertexLinkingDescriptor
func (m MTL4RenderPipelineDynamicLinkingDescriptor) VertexLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vertexLinkingDescriptor"))
	return MTL4PipelineStageDynamicLinkingDescriptorFromID(objc.ID(rv))
}

