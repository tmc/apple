// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	_MTLAccelerationStructureMotionTriangleGeometryDescriptorClass     MTLAccelerationStructureMotionTriangleGeometryDescriptorClass
	_MTLAccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureMotionTriangleGeometryDescriptorClass() MTLAccelerationStructureMotionTriangleGeometryDescriptorClass {
	_MTLAccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureMotionTriangleGeometryDescriptorClass = MTLAccelerationStructureMotionTriangleGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return _MTLAccelerationStructureMotionTriangleGeometryDescriptorClass
}

// GetMTLAccelerationStructureMotionTriangleGeometryDescriptorClass returns the class object for MTLAccelerationStructureMotionTriangleGeometryDescriptor.
func GetMTLAccelerationStructureMotionTriangleGeometryDescriptorClass() MTLAccelerationStructureMotionTriangleGeometryDescriptorClass {
	return getMTLAccelerationStructureMotionTriangleGeometryDescriptorClass()
}

type MTLAccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a list of triangle primitives, as motion keyframe data, to
// turn into an acceleration structure.
//
// # Specifying the number of triangles
//
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.TriangleCount]: The number of triangles in the buffers.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTriangleCount]
//
// # Specifying index data
//
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexBuffer]: A buffer that contains indices for the vertices that compose the triangle list.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBuffer]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexType]: The data type of indices in the index buffer.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexType]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the first index in the buffer.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBufferOffset]
//
// # Specifying vertex data
//
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexBuffers]: An array of motion keyframes, each containing triangle data.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexBuffers]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexStride]: The stride, in bytes, between vertices in each vertex buffer.
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexStride]
//
// # Instance Properties
//
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBuffer]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBufferOffset]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBufferOffset]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixLayout]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexFormat]
//   - [MTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor
type MTLAccelerationStructureMotionTriangleGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureMotionTriangleGeometryDescriptorFromID constructs a [MTLAccelerationStructureMotionTriangleGeometryDescriptor] from an objc.ID.
//
// A description of a list of triangle primitives, as motion keyframe data, to
// turn into an acceleration structure.
func MTLAccelerationStructureMotionTriangleGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	return MTLAccelerationStructureMotionTriangleGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTLAccelerationStructureMotionTriangleGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureMotionTriangleGeometryDescriptor] class.
//
// # Specifying the number of triangles
//
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.TriangleCount]: The number of triangles in the buffers.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTriangleCount]
//
// # Specifying index data
//
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexBuffer]: A buffer that contains indices for the vertices that compose the triangle list.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBuffer]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexType]: The data type of indices in the index buffer.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexType]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the first index in the buffer.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetIndexBufferOffset]
//
// # Specifying vertex data
//
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexBuffers]: An array of motion keyframes, each containing triangle data.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexBuffers]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexStride]: The stride, in bytes, between vertices in each vertex buffer.
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexStride]
//
// # Instance Properties
//
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBuffer]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBuffer]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixBufferOffset]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixBufferOffset]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.TransformationMatrixLayout]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetTransformationMatrixLayout]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.VertexFormat]
//   - [IMTLAccelerationStructureMotionTriangleGeometryDescriptor.SetVertexFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor
type IMTLAccelerationStructureMotionTriangleGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Specifying the number of triangles

	// The number of triangles in the buffers.
	TriangleCount() uint
	SetTriangleCount(value uint)

	// Topic: Specifying index data

	// A buffer that contains indices for the vertices that compose the triangle list.
	IndexBuffer() MTLBuffer
	SetIndexBuffer(value MTLBuffer)
	// The data type of indices in the index buffer.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// The offset, in bytes, to the first index in the buffer.
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)

	// Topic: Specifying vertex data

	// An array of motion keyframes, each containing triangle data.
	VertexBuffers() []MTLMotionKeyframeData
	SetVertexBuffers(value []MTLMotionKeyframeData)
	// The stride, in bytes, between vertices in each vertex buffer.
	VertexStride() uint
	SetVertexStride(value uint)

	// Topic: Instance Properties

	TransformationMatrixBuffer() MTLBuffer
	SetTransformationMatrixBuffer(value MTLBuffer)
	TransformationMatrixBufferOffset() uint
	SetTransformationMatrixBufferOffset(value uint)
	TransformationMatrixLayout() MTLMatrixLayout
	SetTransformationMatrixLayout(value MTLMatrixLayout)
	VertexFormat() MTLAttributeFormat
	SetVertexFormat(value MTLAttributeFormat)
}

// Init initializes the instance.
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) Init() MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionTriangleGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionTriangleGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureMotionTriangleGeometryDescriptor creates a new MTLAccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewMTLAccelerationStructureMotionTriangleGeometryDescriptor() MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	class := getMTLAccelerationStructureMotionTriangleGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new triangle descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/descriptor
func (_MTLAccelerationStructureMotionTriangleGeometryDescriptorClass MTLAccelerationStructureMotionTriangleGeometryDescriptorClass) Descriptor() MTLAccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLAccelerationStructureMotionTriangleGeometryDescriptorClass.class), objc.Sel("descriptor"))
	return MTLAccelerationStructureMotionTriangleGeometryDescriptorFromID(rv)
}

// The number of triangles in the buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("triangleCount"))
	return rv
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setTriangleCount:"), value)
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
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBuffer:"), value)
}

// The data type of indices in the index buffer.
//
// # Discussion
// 
// The index type needs to be [IndexTypeUInt16] or [IndexTypeUInt32]. The
// default is [IndexTypeUInt32].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](a.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexType:"), value)
}

// The offset, in bytes, to the first index in the buffer.
//
// # Discussion
// 
// Specify an offset that is a multiple of the index data type size and a
// multiple of the platform’s buffer offset alignment.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBufferOffset
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// An array of motion keyframes, each containing triangle data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() []MTLMotionKeyframeData {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("vertexBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLMotionKeyframeData {
		return MTLMotionKeyframeDataFromID(id)
	})
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value []MTLMotionKeyframeData) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexBuffers:"), objectivec.IObjectSliceToNSArray(value))
}

// The stride, in bytes, between vertices in each vertex buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("vertexStride"))
	return rv
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexStride:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transformationMatrixBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() MTLMatrixLayout {
	rv := objc.Send[MTLMatrixLayout](a.ID, objc.Sel("transformationMatrixLayout"))
	return MTLMatrixLayout(rv)
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MTLMatrixLayout) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("vertexFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setVertexFormat:"), value)
}

