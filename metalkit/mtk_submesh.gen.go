// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKSubmesh] class.
var (
	_MTKSubmeshClass     MTKSubmeshClass
	_MTKSubmeshClassOnce sync.Once
)

func getMTKSubmeshClass() MTKSubmeshClass {
	_MTKSubmeshClassOnce.Do(func() {
		_MTKSubmeshClass = MTKSubmeshClass{class: objc.GetClass("MTKSubmesh")}
	})
	return _MTKSubmeshClass
}

// GetMTKSubmeshClass returns the class object for MTKSubmesh.
func GetMTKSubmeshClass() MTKSubmeshClass {
	return getMTKSubmeshClass()
}

type MTKSubmeshClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTKSubmeshClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKSubmeshClass) Alloc() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A container for the index data of a Model I/O submesh, suitable for use in
// a Metal app.
//
// # Overview
//
// The [MTKSubmesh] class provides a container for a segment of mesh data that
// can be rendered in a single draw call. A submesh can only be initialized as
// part of a [MTKMesh] object. Each submesh contains an index buffer with
// which the parent’s mesh data can be rendered. Actual submesh vertex data
// resides in the submesh’s parent mesh. For more information on Model I/O
// submeshes, see [MDLSubmesh].
//
// # Parent Mesh
//
//   - [MTKSubmesh.Mesh]: The parent mesh containing the vertex data of this submesh.
//
// # Properties used to Draw Indexed Primitives
//
//   - [MTKSubmesh.IndexBuffer]: The index buffer used to render the submesh object.
//   - [MTKSubmesh.IndexCount]: The number of indices in the index buffer.
//   - [MTKSubmesh.IndexType]: The type of index data in the index buffer.
//   - [MTKSubmesh.PrimitiveType]: The primitive type with which to draw the submesh object.
//
// # Identifying Properties
//
//   - [MTKSubmesh.Name]: The name of the submesh.
//   - [MTKSubmesh.SetName]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh
//
// [MDLSubmesh]: https://developer.apple.com/documentation/ModelIO/MDLSubmesh
type MTKSubmesh struct {
	objectivec.Object
}

// MTKSubmeshFromID constructs a [MTKSubmesh] from an objc.ID.
//
// A container for the index data of a Model I/O submesh, suitable for use in
// a Metal app.
func MTKSubmeshFromID(id objc.ID) MTKSubmesh {
	return MTKSubmesh{objectivec.Object{ID: id}}
}

// NOTE: MTKSubmesh adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKSubmesh] class.
//
// # Parent Mesh
//
//   - [IMTKSubmesh.Mesh]: The parent mesh containing the vertex data of this submesh.
//
// # Properties used to Draw Indexed Primitives
//
//   - [IMTKSubmesh.IndexBuffer]: The index buffer used to render the submesh object.
//   - [IMTKSubmesh.IndexCount]: The number of indices in the index buffer.
//   - [IMTKSubmesh.IndexType]: The type of index data in the index buffer.
//   - [IMTKSubmesh.PrimitiveType]: The primitive type with which to draw the submesh object.
//
// # Identifying Properties
//
//   - [IMTKSubmesh.Name]: The name of the submesh.
//   - [IMTKSubmesh.SetName]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh
type IMTKSubmesh interface {
	objectivec.IObject

	// Topic: Parent Mesh

	// The parent mesh containing the vertex data of this submesh.
	Mesh() IMTKMesh

	// Topic: Properties used to Draw Indexed Primitives

	// The index buffer used to render the submesh object.
	IndexBuffer() IMTKMeshBuffer
	// The number of indices in the index buffer.
	IndexCount() uint
	// The type of index data in the index buffer.
	IndexType() unsafe.Pointer
	// The primitive type with which to draw the submesh object.
	PrimitiveType() unsafe.Pointer

	// Topic: Identifying Properties

	// The name of the submesh.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (s MTKSubmesh) Init() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTKSubmesh) Autorelease() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKSubmesh creates a new MTKSubmesh instance.
func NewMTKSubmesh() MTKSubmesh {
	class := getMTKSubmeshClass()
	rv := objc.Send[MTKSubmesh](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The parent mesh containing the vertex data of this submesh.
//
// # Discussion
//
// The buffer of this parent mesh should be set in the encoder before a call
// to
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]
// is made.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/mesh
//
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (s MTKSubmesh) Mesh() IMTKMesh {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("mesh"))
	return MTKMeshFromID(objc.ID(rv))
}

// The index buffer used to render the submesh object.
//
// # Discussion
//
// Use this object for the `indexBuffer` parameter in a call to
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)].
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexBuffer
//
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (s MTKSubmesh) IndexBuffer() IMTKMeshBuffer {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("indexBuffer"))
	return MTKMeshBufferFromID(objc.ID(rv))
}

// The number of indices in the index buffer.
//
// # Discussion
//
// Use this value for the `indexCount` parameter in a call to
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)].
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexCount
//
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (s MTKSubmesh) IndexCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("indexCount"))
	return rv
}

// The type of index data in the index buffer.
//
// # Discussion
//
// Use this value for the `indexType` parameter in a call to
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)].
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexType
//
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (s MTKSubmesh) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("indexType"))
	return rv
}

// The primitive type with which to draw the submesh object.
//
// # Discussion
//
// Use this value for the `primitiveType` parameter in a call to
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)].
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/primitiveType
//
// [drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (s MTKSubmesh) PrimitiveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("primitiveType"))
	return rv
}

// The name of the submesh.
//
// # Discussion
//
// Your application can use this value to identify the mesh in a scene.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/name
func (s MTKSubmesh) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (s MTKSubmesh) SetName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setName:"), objc.String(value))
}
