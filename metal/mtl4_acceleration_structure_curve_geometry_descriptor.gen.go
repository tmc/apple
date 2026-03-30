// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureCurveGeometryDescriptorClass     MTL4AccelerationStructureCurveGeometryDescriptorClass
	_MTL4AccelerationStructureCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureCurveGeometryDescriptorClass() MTL4AccelerationStructureCurveGeometryDescriptorClass {
	_MTL4AccelerationStructureCurveGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureCurveGeometryDescriptorClass = MTL4AccelerationStructureCurveGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureCurveGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureCurveGeometryDescriptorClass
}

// GetMTL4AccelerationStructureCurveGeometryDescriptorClass returns the class object for MTL4AccelerationStructureCurveGeometryDescriptor.
func GetMTL4AccelerationStructureCurveGeometryDescriptorClass() MTL4AccelerationStructureCurveGeometryDescriptorClass {
	return getMTL4AccelerationStructureCurveGeometryDescriptorClass()
}

type MTL4AccelerationStructureCurveGeometryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4AccelerationStructureCurveGeometryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes curve geometry suitable for ray tracing.
//
// # Overview
//
// Use a [MTLResidencySet] to mark residency of all buffers this descriptor
// references when you build this acceleration structure.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.ControlPointBuffer]: References a buffer containing curve control points.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointBuffer]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.ControlPointCount]: Declares the number of control points in the control point buffer.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointCount]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.ControlPointFormat]: Declares the format of the control points the control point buffer references.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointFormat]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.ControlPointStride]: Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointStride]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.CurveBasis]: Controls the curve basis function, determining how Metal interpolates the control points.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetCurveBasis]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.CurveEndCaps]: Sets the type of curve end caps.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetCurveEndCaps]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.CurveType]: Controls the curve type.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetCurveType]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to control points in the control point buffer.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetIndexBuffer]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.IndexType]: Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetIndexType]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.RadiusBuffer]: Assigns a reference to a buffer containing the curve radius for each control point.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusBuffer]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.RadiusFormat]: Declares the format of the radii in the radius buffer.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusFormat]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.RadiusStride]: Configures the stride, in bytes, between radii in the radius buffer.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusStride]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SegmentControlPointCount]: Declares the number of control points per curve segment.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SegmentCount]: Declares the number of curve segments.
//   - [MTL4AccelerationStructureCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor
type MTL4AccelerationStructureCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureCurveGeometryDescriptorFromID constructs a [MTL4AccelerationStructureCurveGeometryDescriptor] from an objc.ID.
//
// Describes curve geometry suitable for ray tracing.
func MTL4AccelerationStructureCurveGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureCurveGeometryDescriptor {
	return MTL4AccelerationStructureCurveGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}

// NOTE: MTL4AccelerationStructureCurveGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.ControlPointBuffer]: References a buffer containing curve control points.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointBuffer]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.ControlPointCount]: Declares the number of control points in the control point buffer.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointCount]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.ControlPointFormat]: Declares the format of the control points the control point buffer references.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointFormat]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.ControlPointStride]: Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetControlPointStride]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.CurveBasis]: Controls the curve basis function, determining how Metal interpolates the control points.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetCurveBasis]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.CurveEndCaps]: Sets the type of curve end caps.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetCurveEndCaps]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.CurveType]: Controls the curve type.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetCurveType]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to control points in the control point buffer.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetIndexBuffer]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.IndexType]: Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetIndexType]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.RadiusBuffer]: Assigns a reference to a buffer containing the curve radius for each control point.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusBuffer]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.RadiusFormat]: Declares the format of the radii in the radius buffer.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusFormat]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.RadiusStride]: Configures the stride, in bytes, between radii in the radius buffer.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetRadiusStride]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SegmentControlPointCount]: Declares the number of control points per curve segment.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SegmentCount]: Declares the number of curve segments.
//   - [IMTL4AccelerationStructureCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor
type IMTL4AccelerationStructureCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// References a buffer containing curve control points.
	ControlPointBuffer() MTL4BufferRange
	SetControlPointBuffer(value MTL4BufferRange)
	// Declares the number of control points in the control point buffer.
	ControlPointCount() uint
	SetControlPointCount(value uint)
	// Declares the format of the control points the control point buffer references.
	ControlPointFormat() MTLAttributeFormat
	SetControlPointFormat(value MTLAttributeFormat)
	// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
	ControlPointStride() uint
	SetControlPointStride(value uint)
	// Controls the curve basis function, determining how Metal interpolates the control points.
	CurveBasis() MTLCurveBasis
	SetCurveBasis(value MTLCurveBasis)
	// Sets the type of curve end caps.
	CurveEndCaps() MTLCurveEndCaps
	SetCurveEndCaps(value MTLCurveEndCaps)
	// Controls the curve type.
	CurveType() MTLCurveType
	SetCurveType(value MTLCurveType)
	// Assigns an optional index buffer containing references to control points in the control point buffer.
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	// Specifies the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// Assigns a reference to a buffer containing the curve radius for each control point.
	RadiusBuffer() MTL4BufferRange
	SetRadiusBuffer(value MTL4BufferRange)
	// Declares the format of the radii in the radius buffer.
	RadiusFormat() MTLAttributeFormat
	SetRadiusFormat(value MTLAttributeFormat)
	// Configures the stride, in bytes, between radii in the radius buffer.
	RadiusStride() uint
	SetRadiusStride(value uint)
	// Declares the number of control points per curve segment.
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	// Declares the number of curve segments.
	SegmentCount() uint
	SetSegmentCount(value uint)
}

// Init initializes the instance.
func (m MTL4AccelerationStructureCurveGeometryDescriptor) Init() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureCurveGeometryDescriptor creates a new MTL4AccelerationStructureCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureCurveGeometryDescriptor() MTL4AccelerationStructureCurveGeometryDescriptor {
	class := getMTL4AccelerationStructureCurveGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// References a buffer containing curve control points.
//
// # Discussion
//
// Control points are interpolated according to the basis function you specify
// in [CurveBasis].
//
// You are responsible for ensuring each control is in a format matching the
// control point format [ControlPointFormat] specifies, as well as ensuring
// that the buffer address of the range is not zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (m MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("controlPointBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointBuffer:"), value)
}

// Declares the number of control points in the control point buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointCount
func (m MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("controlPointCount"))
	return rv
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointCount:"), value)
}

// Declares the format of the control points the control point buffer
// references.
//
// # Discussion
//
// Defaults to [MTLAttributeFormatFloat3], representing 3 floating point
// values tightly packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (m MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("controlPointFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointFormat:"), value)
}

// Sets the stride, in bytes, between control points in the control point
// buffer the control point buffer references.
//
// # Discussion
//
// You are responsible for ensuring this stride is a multiple of the control
// point format’s element size, and at a minimum exactly the control point
// format’s size.
//
// This property defaults to `0`, indicating that the control points are
// tightly-packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointStride
func (m MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("controlPointStride"))
	return rv
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointStride:"), value)
}

// Controls the curve basis function, determining how Metal interpolates the
// control points.
//
// # Discussion
//
// Defaults to [MTLCurveBasisBSpline].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveBasis
func (m MTL4AccelerationStructureCurveGeometryDescriptor) CurveBasis() MTLCurveBasis {
	rv := objc.Send[MTLCurveBasis](m.ID, objc.Sel("curveBasis"))
	return MTLCurveBasis(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value MTLCurveBasis) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveBasis:"), value)
}

// Sets the type of curve end caps.
//
// # Discussion
//
// Defaults to [MTLCurveEndCapsNone].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (m MTL4AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() MTLCurveEndCaps {
	rv := objc.Send[MTLCurveEndCaps](m.ID, objc.Sel("curveEndCaps"))
	return MTLCurveEndCaps(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value MTLCurveEndCaps) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveEndCaps:"), value)
}

// Controls the curve type.
//
// # Discussion
//
// Defaults to [MTLCurveTypeRound].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveType
func (m MTL4AccelerationStructureCurveGeometryDescriptor) CurveType() MTLCurveType {
	rv := objc.Send[MTLCurveType](m.ID, objc.Sel("curveType"))
	return MTLCurveType(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveType(value MTLCurveType) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveType:"), value)
}

// Assigns an optional index buffer containing references to control points in
// the control point buffer.
//
// # Discussion
//
// Each index represents the first control point of a curve segment. You are
// responsible for ensuring the buffer address of the range is not zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexBuffer
func (m MTL4AccelerationStructureCurveGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("indexBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexBuffer:"), value)
}

// Specifies the size of the indices the `indexBuffer` contains, which is
// typically either 16 or 32-bits for each index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexType
func (m MTL4AccelerationStructureCurveGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](m.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}

// Assigns a reference to a buffer containing the curve radius for each
// control point.
//
// # Discussion
//
// Metal interpolates curve radii according to the basis function you specify
// via [CurveBasis].
//
// You are responsible for ensuring the type of each radius matches the type
// property [RadiusFormat] specifies, that each radius is at least zero, and
// that the buffer address of the range is not zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (m MTL4AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("radiusBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusBuffer:"), value)
}

// Declares the format of the radii in the radius buffer.
//
// # Discussion
//
// Defaults to [MTLAttributeFormatFloat].
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusFormat
func (m MTL4AccelerationStructureCurveGeometryDescriptor) RadiusFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("radiusFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusFormat:"), value)
}

// Configures the stride, in bytes, between radii in the radius buffer.
//
// # Discussion
//
// You are responsible for ensuring this property is set to a multiple of the
// size corresponding to the [RadiusFormat].
//
// This property defaults to `0` bytes, indicating that the radii are tightly
// packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusStride
func (m MTL4AccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("radiusStride"))
	return rv
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusStride:"), value)
}

// Declares the number of control points per curve segment.
//
// # Discussion
//
// Valid values for this property are `2`, `3`, or `4`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("segmentControlPointCount"))
	return rv
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

// Declares the number of curve segments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentCount
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("segmentCount"))
	return rv
}
func (m MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSegmentCount:"), value)
}
