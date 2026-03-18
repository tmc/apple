// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureTriangleGeometryDescriptorClass     MTL4AccelerationStructureTriangleGeometryDescriptorClass
	_MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureTriangleGeometryDescriptorClass() MTL4AccelerationStructureTriangleGeometryDescriptorClass {
	_MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureTriangleGeometryDescriptorClass = MTL4AccelerationStructureTriangleGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureTriangleGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureTriangleGeometryDescriptorClass
}

// GetMTL4AccelerationStructureTriangleGeometryDescriptorClass returns the class object for MTL4AccelerationStructureTriangleGeometryDescriptor.
func GetMTL4AccelerationStructureTriangleGeometryDescriptorClass() MTL4AccelerationStructureTriangleGeometryDescriptorClass {
	return getMTL4AccelerationStructureTriangleGeometryDescriptorClass()
}

type MTL4AccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes triangle geometry suitable for ray tracing.
//
// # Overview
// 
// Use a [MTLResidencySet] to mark residency of all buffers this descriptor
// references when you build this acceleration structure.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.IndexBuffer]: Sets an optional index buffer containing references to vertices in the `vertexBuffer`.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetIndexBuffer]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.IndexType]: Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetIndexType]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBuffer]: Assigns an optional reference to a buffer containing a `float4x3` transformation matrix.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.TransformationMatrixLayout]: Configures the layout for the transformation matrix in the transformation matrix buffer.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.TriangleCount]: Declares the number of triangles in this geometry descriptor.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetTriangleCount]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.VertexBuffer]: Associates a vertex buffer containing triangle vertices.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexBuffer]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.VertexFormat]: Describes the format of the vertices in the vertex buffer.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexFormat]
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.VertexStride]: Sets the stride, in bytes, between vertices in the vertex buffer.
//   - [MTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor
type MTL4AccelerationStructureTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureTriangleGeometryDescriptorFromID constructs a [MTL4AccelerationStructureTriangleGeometryDescriptor] from an objc.ID.
//
// Describes triangle geometry suitable for ray tracing.
func MTL4AccelerationStructureTriangleGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureTriangleGeometryDescriptor {
	return MTL4AccelerationStructureTriangleGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTL4AccelerationStructureTriangleGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.IndexBuffer]: Sets an optional index buffer containing references to vertices in the `vertexBuffer`.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetIndexBuffer]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.IndexType]: Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetIndexType]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBuffer]: Assigns an optional reference to a buffer containing a `float4x3` transformation matrix.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.TransformationMatrixLayout]: Configures the layout for the transformation matrix in the transformation matrix buffer.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.TriangleCount]: Declares the number of triangles in this geometry descriptor.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetTriangleCount]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.VertexBuffer]: Associates a vertex buffer containing triangle vertices.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexBuffer]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.VertexFormat]: Describes the format of the vertices in the vertex buffer.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexFormat]
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.VertexStride]: Sets the stride, in bytes, between vertices in the vertex buffer.
//   - [IMTL4AccelerationStructureTriangleGeometryDescriptor.SetVertexStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor
type IMTL4AccelerationStructureTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// Sets an optional index buffer containing references to vertices in the `vertexBuffer`.
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	// Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// Assigns an optional reference to a buffer containing a `float4x3` transformation matrix.
	TransformationMatrixBuffer() MTL4BufferRange
	SetTransformationMatrixBuffer(value MTL4BufferRange)
	// Configures the layout for the transformation matrix in the transformation matrix buffer.
	TransformationMatrixLayout() MTLMatrixLayout
	SetTransformationMatrixLayout(value MTLMatrixLayout)
	// Declares the number of triangles in this geometry descriptor.
	TriangleCount() uint
	SetTriangleCount(value uint)
	// Associates a vertex buffer containing triangle vertices.
	VertexBuffer() MTL4BufferRange
	SetVertexBuffer(value MTL4BufferRange)
	// Describes the format of the vertices in the vertex buffer.
	VertexFormat() MTLAttributeFormat
	SetVertexFormat(value MTLAttributeFormat)
	// Sets the stride, in bytes, between vertices in the vertex buffer.
	VertexStride() uint
	SetVertexStride(value uint)
}

// Init initializes the instance.
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) Init() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureTriangleGeometryDescriptor creates a new MTL4AccelerationStructureTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureTriangleGeometryDescriptor() MTL4AccelerationStructureTriangleGeometryDescriptor {
	class := getMTL4AccelerationStructureTriangleGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets an optional index buffer containing references to vertices in the
// `vertexBuffer`.
//
// # Discussion
// 
// You can set this property to `0`, the default, to avoid specifying an index
// buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("indexBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexBuffer:"), value)
}

// Configures the size of the indices the `indexBuffer` contains, which is
// typically either 16 or 32-bits for each index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexType
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](m.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}

// Assigns an optional reference to a buffer containing a `float4x3`
// transformation matrix.
//
// # Discussion
// 
// When the buffer address is non-zero, Metal applies this transform to the
// vertex data positions when building the acceleration structure.
// 
// Building an acceleration structure with a descriptor that specifies this
// property doesn’t modify the contents of the input `vertexBuffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("transformationMatrixBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

// Configures the layout for the transformation matrix in the transformation
// matrix buffer.
//
// # Discussion
// 
// You can provide matrices in column-major or row-major form, and this
// property allows you to control how Metal interprets them.
// 
// Defaults to [MTLMatrixLayoutColumnMajor].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](m.ID, objc.Sel("transformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](m.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// Declares the number of triangles in this geometry descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/triangleCount
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("triangleCount"))
	return rv
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setTriangleCount:"), value)
}

// Associates a vertex buffer containing triangle vertices.
//
// # Discussion
// 
// You are responsible for ensuring that the format of all vertex positions
// match the [VertexFormat] property, and that the buffer address for the
// buffer range is not zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("vertexBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexBuffer:"), value)
}

// Describes the format of the vertices in the vertex buffer.
//
// # Discussion
// 
// This property controls the format of the position attribute of the vertices
// the [VertexBuffer] references.
// 
// The format defaults to [MTLAttributeFormatFloat3], corresponding to three
// packed floating point numbers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) VertexFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("vertexFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexFormat:"), value)
}

// Sets the stride, in bytes, between vertices in the vertex buffer.
//
// # Discussion
// 
// The stride you specify needs to be a multiple of the size of the vertex
// format you provide in the [VertexFormat] property. Similarly, you are
// responsible for ensuring this stride matches the vertex format data
// type’s alignment.
// 
// Defaults to `0`, which signals the stride matches the size of the
// [VertexFormat] data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexStride
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("vertexStride"))
	return rv
}
func (m MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexStride:"), value)
}

