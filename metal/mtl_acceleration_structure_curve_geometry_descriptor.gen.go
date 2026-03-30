// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLAccelerationStructureCurveGeometryDescriptor] class.
var (
	_MTLAccelerationStructureCurveGeometryDescriptorClass     MTLAccelerationStructureCurveGeometryDescriptorClass
	_MTLAccelerationStructureCurveGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureCurveGeometryDescriptorClass() MTLAccelerationStructureCurveGeometryDescriptorClass {
	_MTLAccelerationStructureCurveGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureCurveGeometryDescriptorClass = MTLAccelerationStructureCurveGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureCurveGeometryDescriptor")}
	})
	return _MTLAccelerationStructureCurveGeometryDescriptorClass
}

// GetMTLAccelerationStructureCurveGeometryDescriptorClass returns the class object for MTLAccelerationStructureCurveGeometryDescriptor.
func GetMTLAccelerationStructureCurveGeometryDescriptorClass() MTLAccelerationStructureCurveGeometryDescriptorClass {
	return getMTLAccelerationStructureCurveGeometryDescriptorClass()
}

type MTLAccelerationStructureCurveGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructureCurveGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureCurveGeometryDescriptorClass) Alloc() MTLAccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A descriptor you configure with curve geometry for building acceleration
// structures.
//
// # Instance Properties
//
//   - [MTLAccelerationStructureCurveGeometryDescriptor.ControlPointBuffer]: A buffer that contains curve control points.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetControlPointBuffer]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.ControlPointBufferOffset]: The offset, in bytes, to the control point data in the buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetControlPointBufferOffset]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.ControlPointCount]: The number of control points in the control point buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetControlPointCount]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.ControlPointFormat]: The format of the control points in the buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetControlPointFormat]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.ControlPointStride]: The stride, in bytes, between control points in the buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetControlPointStride]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.CurveBasis]: The basis function for the curve geometry.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetCurveBasis]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.CurveEndCaps]: An end-cap type for the curves in the geometry.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetCurveEndCaps]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.CurveType]: A curve type for curves in the geometry.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetCurveType]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.IndexBuffer]: A buffer that contains references to control points in the control point buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetIndexBuffer]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the index data in the buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetIndexBufferOffset]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.IndexType]: The size of each index in the index buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetIndexType]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.RadiusBuffer]: A buffer that contains the curve radius for each control point.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetRadiusBuffer]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.RadiusBufferOffset]: The offset, in bytes, to the radius data in the buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetRadiusBufferOffset]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.RadiusFormat]: The format of each radius in the radius buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetRadiusFormat]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.RadiusStride]: The stride, in bytes, between the radius elements in the radius buffer.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetRadiusStride]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SegmentControlPointCount]: The number of control points in each curve segment.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SegmentCount]: The number of curve segments in each curve of the geometry.
//   - [MTLAccelerationStructureCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor
type MTLAccelerationStructureCurveGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureCurveGeometryDescriptorFromID constructs a [MTLAccelerationStructureCurveGeometryDescriptor] from an objc.ID.
//
// A descriptor you configure with curve geometry for building acceleration
// structures.
func MTLAccelerationStructureCurveGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureCurveGeometryDescriptor {
	return MTLAccelerationStructureCurveGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}

// NOTE: MTLAccelerationStructureCurveGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureCurveGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.ControlPointBuffer]: A buffer that contains curve control points.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetControlPointBuffer]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.ControlPointBufferOffset]: The offset, in bytes, to the control point data in the buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetControlPointBufferOffset]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.ControlPointCount]: The number of control points in the control point buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetControlPointCount]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.ControlPointFormat]: The format of the control points in the buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetControlPointFormat]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.ControlPointStride]: The stride, in bytes, between control points in the buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetControlPointStride]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.CurveBasis]: The basis function for the curve geometry.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetCurveBasis]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.CurveEndCaps]: An end-cap type for the curves in the geometry.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetCurveEndCaps]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.CurveType]: A curve type for curves in the geometry.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetCurveType]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.IndexBuffer]: A buffer that contains references to control points in the control point buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetIndexBuffer]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.IndexBufferOffset]: The offset, in bytes, to the index data in the buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetIndexBufferOffset]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.IndexType]: The size of each index in the index buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetIndexType]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.RadiusBuffer]: A buffer that contains the curve radius for each control point.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetRadiusBuffer]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.RadiusBufferOffset]: The offset, in bytes, to the radius data in the buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetRadiusBufferOffset]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.RadiusFormat]: The format of each radius in the radius buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetRadiusFormat]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.RadiusStride]: The stride, in bytes, between the radius elements in the radius buffer.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetRadiusStride]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SegmentControlPointCount]: The number of control points in each curve segment.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SegmentCount]: The number of curve segments in each curve of the geometry.
//   - [IMTLAccelerationStructureCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor
type IMTLAccelerationStructureCurveGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// A buffer that contains curve control points.
	ControlPointBuffer() MTLBuffer
	SetControlPointBuffer(value MTLBuffer)
	// The offset, in bytes, to the control point data in the buffer.
	ControlPointBufferOffset() uint
	SetControlPointBufferOffset(value uint)
	// The number of control points in the control point buffer.
	ControlPointCount() uint
	SetControlPointCount(value uint)
	// The format of the control points in the buffer.
	ControlPointFormat() MTLAttributeFormat
	SetControlPointFormat(value MTLAttributeFormat)
	// The stride, in bytes, between control points in the buffer.
	ControlPointStride() uint
	SetControlPointStride(value uint)
	// The basis function for the curve geometry.
	CurveBasis() MTLCurveBasis
	SetCurveBasis(value MTLCurveBasis)
	// An end-cap type for the curves in the geometry.
	CurveEndCaps() MTLCurveEndCaps
	SetCurveEndCaps(value MTLCurveEndCaps)
	// A curve type for curves in the geometry.
	CurveType() MTLCurveType
	SetCurveType(value MTLCurveType)
	// A buffer that contains references to control points in the control point buffer.
	IndexBuffer() MTLBuffer
	SetIndexBuffer(value MTLBuffer)
	// The offset, in bytes, to the index data in the buffer.
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	// The size of each index in the index buffer.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// A buffer that contains the curve radius for each control point.
	RadiusBuffer() MTLBuffer
	SetRadiusBuffer(value MTLBuffer)
	// The offset, in bytes, to the radius data in the buffer.
	RadiusBufferOffset() uint
	SetRadiusBufferOffset(value uint)
	// The format of each radius in the radius buffer.
	RadiusFormat() MTLAttributeFormat
	SetRadiusFormat(value MTLAttributeFormat)
	// The stride, in bytes, between the radius elements in the radius buffer.
	RadiusStride() uint
	SetRadiusStride(value uint)
	// The number of control points in each curve segment.
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	// The number of curve segments in each curve of the geometry.
	SegmentCount() uint
	SetSegmentCount(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructureCurveGeometryDescriptor) Init() MTLAccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureCurveGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureCurveGeometryDescriptor) Autorelease() MTLAccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureCurveGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureCurveGeometryDescriptor creates a new MTLAccelerationStructureCurveGeometryDescriptor instance.
func NewMTLAccelerationStructureCurveGeometryDescriptor() MTLAccelerationStructureCurveGeometryDescriptor {
	class := getMTLAccelerationStructureCurveGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureCurveGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A buffer that contains curve control points.
//
// # Discussion
//
// You provide control points in the format that matches the
// [ControlPointFormat] property. This property needs to have a non-nil value
// when you build an acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (a MTLAccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("controlPointBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointBuffer:"), value)
}

// The offset, in bytes, to the control point data in the buffer.
//
// # Discussion
//
// The offset needs to be a multiple of the format element size you configure
// with the [ControlPointFormat] property. You also need to align the offset
// to the platform’s buffer alignment requirement.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBufferOffset
func (a MTLAccelerationStructureCurveGeometryDescriptor) ControlPointBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("controlPointBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointBufferOffset:"), value)
}

// The number of control points in the control point buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointCount
func (a MTLAccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("controlPointCount"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointCount:"), value)
}

// The format of the control points in the buffer.
//
// # Discussion
//
// The default value is [MTLAttributeFormat.float3].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointFormat
//
// [MTLAttributeFormat.float3]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float3
func (a MTLAccelerationStructureCurveGeometryDescriptor) ControlPointFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("controlPointFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointFormat:"), value)
}

// The stride, in bytes, between control points in the buffer.
//
// # Discussion
//
// The stride needs to be a multiple of the format element size you configure
// with the [ControlPointFormat] property, and at least the format’s size.
// The default value is `0`, which indicates that the control point elements
// in the buffer have zero bytes of padding between them.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointStride
func (a MTLAccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("controlPointStride"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointStride:"), value)
}

// The basis function for the curve geometry.
//
// # Discussion
//
// The default value is [MTLCurveBasis.bSpline].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveBasis
//
// [MTLCurveBasis.bSpline]: https://developer.apple.com/documentation/Metal/MTLCurveBasis/bSpline
func (a MTLAccelerationStructureCurveGeometryDescriptor) CurveBasis() MTLCurveBasis {
	rv := objc.Send[MTLCurveBasis](a.ID, objc.Sel("curveBasis"))
	return MTLCurveBasis(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value MTLCurveBasis) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveBasis:"), value)
}

// An end-cap type for the curves in the geometry.
//
// # Discussion
//
// The default value is [MTLCurveEndCapsNone].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (a MTLAccelerationStructureCurveGeometryDescriptor) CurveEndCaps() MTLCurveEndCaps {
	rv := objc.Send[MTLCurveEndCaps](a.ID, objc.Sel("curveEndCaps"))
	return MTLCurveEndCaps(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value MTLCurveEndCaps) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveEndCaps:"), value)
}

// A curve type for curves in the geometry.
//
// # Discussion
//
// The default value is [MTLCurveTypeRound].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveType
func (a MTLAccelerationStructureCurveGeometryDescriptor) CurveType() MTLCurveType {
	rv := objc.Send[MTLCurveType](a.ID, objc.Sel("curveType"))
	return MTLCurveType(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetCurveType(value MTLCurveType) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveType:"), value)
}

// A buffer that contains references to control points in the control point
// buffer.
//
// # Discussion
//
// This property needs to have a non-nil value when you build an acceleration
// structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBuffer
func (a MTLAccelerationStructureCurveGeometryDescriptor) IndexBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBuffer:"), value)
}

// The offset, in bytes, to the index data in the buffer.
//
// # Discussion
//
// The offset needs to be a multiple of the index data type you configure with
// the [IndexType] property. You also need to align the offset to both the
// index type’s size and the platform’s buffer alignment requirement.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBufferOffset
func (a MTLAccelerationStructureCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// The size of each index in the index buffer.
//
// # Discussion
//
// Set this property to a value that reflects the size of the indices in the
// [IndexBuffer] property, such as [MTLIndexTypeUInt16] or
// [MTLIndexTypeUInt32].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexType
func (a MTLAccelerationStructureCurveGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](a.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexType:"), value)
}

// A buffer that contains the curve radius for each control point.
//
// # Discussion
//
// The buffer contains values that are greater than or equal to `0.0` in the
// format you configure with the [RadiusFormat] property. This property needs
// to have a non-nil value when you build an acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (a MTLAccelerationStructureCurveGeometryDescriptor) RadiusBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("radiusBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusBuffer:"), value)
}

// The offset, in bytes, to the radius data in the buffer.
//
// # Discussion
//
// The offset needs to be a multiple of the radius format you configure with
// the [RadiusFormat] property. You also need to align the offset to both the
// radius format’s size and the platform’s buffer alignment requirement.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBufferOffset
func (a MTLAccelerationStructureCurveGeometryDescriptor) RadiusBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("radiusBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusBufferOffset:"), value)
}

// The format of each radius in the radius buffer.
//
// # Discussion
//
// The property’s default value is [MTLAttributeFormat.float].
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusFormat
//
// [MTLAttributeFormat.float]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float
func (a MTLAccelerationStructureCurveGeometryDescriptor) RadiusFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("radiusFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusFormat:"), value)
}

// The stride, in bytes, between the radius elements in the radius buffer.
//
// # Discussion
//
// The stride needs to be a multiple of the radius format size you configure
// with the [RadiusFormat] property. The default value is `0`, which indicates
// that the radius elements in the buffer have zero bytes of padding between
// them.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusStride
func (a MTLAccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("radiusStride"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusStride:"), value)
}

// The number of control points in each curve segment.
//
// # Discussion
//
// This value can be `2`, `3`, or `4`.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (a MTLAccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("segmentControlPointCount"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

// The number of curve segments in each curve of the geometry.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentCount
func (a MTLAccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("segmentCount"))
	return rv
}
func (a MTLAccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSegmentCount:"), value)
}
