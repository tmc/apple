// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLAccelerationStructureTriangleGeometryDescriptor] class.
var (
	_MTLAccelerationStructureTriangleGeometryDescriptorClass     MTLAccelerationStructureTriangleGeometryDescriptorClass
	_MTLAccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureTriangleGeometryDescriptorClass() MTLAccelerationStructureTriangleGeometryDescriptorClass {
	_MTLAccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureTriangleGeometryDescriptorClass = MTLAccelerationStructureTriangleGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureTriangleGeometryDescriptor")}
	})
	return _MTLAccelerationStructureTriangleGeometryDescriptorClass
}

// GetMTLAccelerationStructureTriangleGeometryDescriptorClass returns the class object for MTLAccelerationStructureTriangleGeometryDescriptor.
func GetMTLAccelerationStructureTriangleGeometryDescriptorClass() MTLAccelerationStructureTriangleGeometryDescriptorClass {
	return getMTLAccelerationStructureTriangleGeometryDescriptorClass()
}

type MTLAccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureTriangleGeometryDescriptorClass) Alloc() MTLAccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A description of a list of triangle primitives to turn into an acceleration
// structure.
//
// # Configuring the number of triangles
//
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.TriangleCount]: The number of triangles in the buffers.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetTriangleCount]
//
// # Configuring index data
//
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.IndexType]: The data type of indices in the index buffer.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetIndexType]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.IndexBuffer]: A buffer that contains indices for the vertices that compose the triangle list.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetIndexBuffer]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the first index in the buffer.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetIndexBufferOffset]
//
// # Configuring vertex data
//
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.VertexFormat]: The format of each vertex position in the vertex buffer property.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetVertexFormat]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.VertexBuffer]: A buffer that contains vertex data.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetVertexBuffer]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.VertexBufferOffset]: The offset, in bytes, for the first vertex in the vertex buffer.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetVertexBufferOffset]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.VertexStride]: The stride, in bytes, between vertices in the vertex buffer.
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetVertexStride]
//
// # Configuring transformation data
//
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixLayout]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBuffer]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBufferOffset]
//   - [MTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBufferOffset]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor
type MTLAccelerationStructureTriangleGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureTriangleGeometryDescriptorFromID constructs a [MTLAccelerationStructureTriangleGeometryDescriptor] from an objc.ID.
//
// A description of a list of triangle primitives to turn into an acceleration
// structure.
func MTLAccelerationStructureTriangleGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureTriangleGeometryDescriptor {
	return MTLAccelerationStructureTriangleGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTLAccelerationStructureTriangleGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLAccelerationStructureTriangleGeometryDescriptor] class.
//
// # Configuring the number of triangles
//
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.TriangleCount]: The number of triangles in the buffers.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetTriangleCount]
//
// # Configuring index data
//
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.IndexType]: The data type of indices in the index buffer.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetIndexType]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.IndexBuffer]: A buffer that contains indices for the vertices that compose the triangle list.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetIndexBuffer]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the first index in the buffer.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetIndexBufferOffset]
//
// # Configuring vertex data
//
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.VertexFormat]: The format of each vertex position in the vertex buffer property.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetVertexFormat]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.VertexBuffer]: A buffer that contains vertex data.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetVertexBuffer]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.VertexBufferOffset]: The offset, in bytes, for the first vertex in the vertex buffer.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetVertexBufferOffset]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.VertexStride]: The stride, in bytes, between vertices in the vertex buffer.
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetVertexStride]
//
// # Configuring transformation data
//
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixLayout]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBuffer]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.TransformationMatrixBufferOffset]
//   - [IMTLAccelerationStructureTriangleGeometryDescriptor.SetTransformationMatrixBufferOffset]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor
type IMTLAccelerationStructureTriangleGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Configuring the number of triangles

	// The number of triangles in the buffers.
	TriangleCount() uint
	SetTriangleCount(value uint)

	// Topic: Configuring index data

	// The data type of indices in the index buffer.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// A buffer that contains indices for the vertices that compose the triangle list.
	IndexBuffer() MTLBuffer
	SetIndexBuffer(value MTLBuffer)
	// The offset, in bytes, to the first index in the buffer.
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)

	// Topic: Configuring vertex data

	// The format of each vertex position in the vertex buffer property.
	VertexFormat() MTLAttributeFormat
	SetVertexFormat(value MTLAttributeFormat)
	// A buffer that contains vertex data.
	VertexBuffer() MTLBuffer
	SetVertexBuffer(value MTLBuffer)
	// The offset, in bytes, for the first vertex in the vertex buffer.
	VertexBufferOffset() uint
	SetVertexBufferOffset(value uint)
	// The stride, in bytes, between vertices in the vertex buffer.
	VertexStride() uint
	SetVertexStride(value uint)

	// Topic: Configuring transformation data

	TransformationMatrixLayout() MTLMatrixLayout
	SetTransformationMatrixLayout(value MTLMatrixLayout)
	TransformationMatrixBuffer() MTLBuffer
	SetTransformationMatrixBuffer(value MTLBuffer)
	TransformationMatrixBufferOffset() uint
	SetTransformationMatrixBufferOffset(value uint)
}





// Init initializes the instance.
func (a MTLAccelerationStructureTriangleGeometryDescriptor) Init() MTLAccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureTriangleGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureTriangleGeometryDescriptor) Autorelease() MTLAccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureTriangleGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureTriangleGeometryDescriptor creates a new MTLAccelerationStructureTriangleGeometryDescriptor instance.
func NewMTLAccelerationStructureTriangleGeometryDescriptor() MTLAccelerationStructureTriangleGeometryDescriptor {
	class := getMTLAccelerationStructureTriangleGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureTriangleGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Creates a new triangle descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/descriptor
func (_MTLAccelerationStructureTriangleGeometryDescriptorClass MTLAccelerationStructureTriangleGeometryDescriptorClass) Descriptor() MTLAccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLAccelerationStructureTriangleGeometryDescriptorClass.class), objc.Sel("descriptor"))
	return MTLAccelerationStructureTriangleGeometryDescriptorFromID(rv)
}








// The number of triangles in the buffers.
//
// # Discussion
// 
// If the triangle descriptor contains an index buffer, then the index buffer
// needs to provide indices for this many triangles. If the triangle
// descriptor doesn’t provide an index buffer, then the vertex buffer
// provides 3 vertices for each triangle.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/triangleCount
func (a MTLAccelerationStructureTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("triangleCount"))
	return rv
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setTriangleCount:"), value)
}



// The data type of indices in the index buffer.
//
// # Discussion
// 
// The index type needs to be [IndexTypeUInt16] or [IndexTypeUInt32]. The
// default is [IndexTypeUInt32].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexType
func (a MTLAccelerationStructureTriangleGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](a.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexType:"), value)
}



// A buffer that contains indices for the vertices that compose the triangle
// list.
//
// # Discussion
// 
// This property can be `nil`, in which case the vertex data defines the
// triangle list implicitly. You need to store indices in a packed data
// format.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (a MTLAccelerationStructureTriangleGeometryDescriptor) IndexBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBuffer:"), value)
}



// The offset, in bytes, to the first index in the buffer.
//
// # Discussion
// 
// Specify an offset that is a multiple of the index data type size and a
// multiple of the platform’s buffer offset alignment.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBufferOffset
func (a MTLAccelerationStructureTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBufferOffset:"), value)
}



// The format of each vertex position in the vertex buffer property.
//
// # Discussion
// 
// Set this property to a value that represents the pixel format of the data
// you assign to the [VertexBuffer] property. The property’s default is
// [AttributeFormatFloat3].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (a MTLAccelerationStructureTriangleGeometryDescriptor) VertexFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("vertexFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexFormat:"), value)
}



// A buffer that contains vertex data.
//
// # Discussion
// 
// The [VertexFormat] property defines the format of each vertex position in
// the buffer. You need to set a vertex buffer before creating the
// acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (a MTLAccelerationStructureTriangleGeometryDescriptor) VertexBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("vertexBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexBuffer:"), value)
}



// The offset, in bytes, for the first vertex in the vertex buffer.
//
// # Discussion
// 
// The vertex needs to be a multiple of the vertex stride and be a multiple of
// 4 bytes. The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBufferOffset
func (a MTLAccelerationStructureTriangleGeometryDescriptor) VertexBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("vertexBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexBufferOffset:"), value)
}



// The stride, in bytes, between vertices in the vertex buffer.
//
// # Discussion
// 
// The stride needs to be at least 12 bytes and needs to be a multiple of 4
// bytes. The default value is 12 bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexStride
func (a MTLAccelerationStructureTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("vertexStride"))
	return rv
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexStride:"), value)
}



// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (a MTLAccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](a.ID, objc.Sel("transformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}



// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (a MTLAccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transformationMatrixBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}



// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a MTLAccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}
























