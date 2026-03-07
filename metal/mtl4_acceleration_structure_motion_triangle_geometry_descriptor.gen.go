// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass     MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
	_MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass() MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass {
	_MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass = MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
}

// GetMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass returns the class object for MTL4AccelerationStructureMotionTriangleGeometryDescriptor.
func GetMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass() MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass {
	return getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass()
}

type MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Describes motion triangle geometry, suitable for motion ray tracing.
//
// # Overview
// 
// Use a [MTLResidencySet] to mark residency of all buffers this descriptor
// references when you build this acceleration structure.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the vertex buffers property.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBuffer]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.IndexType]: Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetIndexType]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBuffer]: Assings an optional reference to a buffer containing a `float4x3` transformation matrix.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixLayout]: Configures the layout for the transformation matrix in the transformation matrix buffer.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.TriangleCount]: Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTriangleCount]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexBuffers]: Assigns a buffer where each entry contains a reference to a vertex buffer.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexBuffers]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexFormat]: Defines the format of the vertices in the vertex buffers.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexFormat]
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexStride]: Sets the stride, in bytes, between vertices in all the vertex buffer.
//   - [MTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor
type MTL4AccelerationStructureMotionTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionTriangleGeometryDescriptorFromID constructs a [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] from an objc.ID.
//
// Describes motion triangle geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionTriangleGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	return MTL4AccelerationStructureMotionTriangleGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTL4AccelerationStructureMotionTriangleGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the vertex buffers property.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBuffer]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.IndexType]: Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetIndexType]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBuffer]: Assings an optional reference to a buffer containing a `float4x3` transformation matrix.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixLayout]: Configures the layout for the transformation matrix in the transformation matrix buffer.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.TriangleCount]: Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetTriangleCount]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexBuffers]: Assigns a buffer where each entry contains a reference to a vertex buffer.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexBuffers]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexFormat]: Defines the format of the vertices in the vertex buffers.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexFormat]
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.VertexStride]: Sets the stride, in bytes, between vertices in all the vertex buffer.
//   - [IMTL4AccelerationStructureMotionTriangleGeometryDescriptor.SetVertexStride]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor
type IMTL4AccelerationStructureMotionTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the vertex buffers property.
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	// Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// Assings an optional reference to a buffer containing a `float4x3` transformation matrix.
	TransformationMatrixBuffer() MTL4BufferRange
	SetTransformationMatrixBuffer(value MTL4BufferRange)
	// Configures the layout for the transformation matrix in the transformation matrix buffer.
	TransformationMatrixLayout() MTLMatrixLayout
	SetTransformationMatrixLayout(value MTLMatrixLayout)
	// Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
	TriangleCount() uint
	SetTriangleCount(value uint)
	// Assigns a buffer where each entry contains a reference to a vertex buffer.
	VertexBuffers() MTL4BufferRange
	SetVertexBuffers(value MTL4BufferRange)
	// Defines the format of the vertices in the vertex buffers.
	VertexFormat() MTLAttributeFormat
	SetVertexFormat(value MTLAttributeFormat)
	// Sets the stride, in bytes, between vertices in all the vertex buffer.
	VertexStride() uint
	SetVertexStride(value uint)
}





// Init initializes the instance.
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Init() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor creates a new MTL4AccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	class := getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Assigns an optional index buffer containing references to vertices in the
// vertex buffers you reference through the vertex buffers property.
//
// # Discussion
// 
// You can set this property to `0`, the default, to avoid specifying an index
// buffer. All keyframes share the same index buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("indexBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexBuffer:"), value)
}



// Specifies the size of the indices the `indexBuffer` contains, which is
// typically either 16 or 32-bits for each index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](m.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}



// Assings an optional reference to a buffer containing a `float4x3`
// transformation matrix.
//
// # Discussion
// 
// When the buffer address is non-zero, Metal applies this transform to the
// vertex data positions when building the acceleration structure. All
// keyframes share the same transformation matrix.
// 
// Building an acceleration structure with a descriptor that specifies this
// property doesn’t modify the contents of the input `vertexBuffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("transformationMatrixBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value MTL4BufferRange) {
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
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](m.ID, objc.Sel("transformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](m.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}



// Declares the number of triangles in the vertex buffers that the buffer in
// the vertex buffers property references.
//
// # Discussion
// 
// All keyframes share the same triangle count.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("triangleCount"))
	return rv
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setTriangleCount:"), value)
}



// Assigns a buffer where each entry contains a reference to a vertex buffer.
//
// # Discussion
// 
// This property references a buffer that conceptually represents an array
// with one entry for each keyframe in the motion animation. Each one of these
// entries consists of a [MTL4BufferRange] that, in turn, references a vertex
// buffer containing the vertex data for the keyframe.
// 
// You are responsible for ensuring the buffer address is not zero for the
// top-level buffer, as well as for all the vertex buffers it references.
//
// [MTL4BufferRange]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("vertexBuffers"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexBuffers:"), value)
}



// Defines the format of the vertices in the vertex buffers.
//
// # Discussion
// 
// All keyframes share the same vertex format. Defaults to
// [MTLAttributeFormatFloat3], corresponding to three packed floating point
// numbers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("vertexFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexFormat:"), value)
}



// Sets the stride, in bytes, between vertices in all the vertex buffer.
//
// # Discussion
// 
// All keyframes share the same vertex stride. This stride needs to be a
// multiple of the size of the vertex format you provide in the [VertexFormat]
// property.
// 
// Similarly, you are responsible for ensuring this stride matches the vertex
// format data type’s alignment.
// 
// Defaults to `0`, which signals the stride matches the size of the
// [VertexFormat] data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("vertexStride"))
	return rv
}
func (m MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexStride:"), value)
}
























