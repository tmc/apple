// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
var (
	_MTL4AccelerationStructureMotionCurveGeometryDescriptorClass     MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
	_MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass() MTL4AccelerationStructureMotionCurveGeometryDescriptorClass {
	_MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureMotionCurveGeometryDescriptorClass = MTL4AccelerationStructureMotionCurveGeometryDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureMotionCurveGeometryDescriptor")}
	})
	return _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
}

// GetMTL4AccelerationStructureMotionCurveGeometryDescriptorClass returns the class object for MTL4AccelerationStructureMotionCurveGeometryDescriptor.
func GetMTL4AccelerationStructureMotionCurveGeometryDescriptorClass() MTL4AccelerationStructureMotionCurveGeometryDescriptorClass {
	return getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass()
}

type MTL4AccelerationStructureMotionCurveGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Describes motion curve geometry, suitable for motion ray tracing.
//
// # Overview
// 
// Use a [MTLResidencySet] to mark residency of all buffers this descriptor
// references when you build this acceleration structure.
//
// # Instance Properties
//
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointBuffers]: Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointBuffers]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointCount]: Specifies the number of control points in the buffers the control point buffers reference.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointCount]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointFormat]: Declares the format of the control points in the buffers that the control point buffers reference.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointFormat]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointStride]: Sets the stride, in bytes, between control points in the control point buffer.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointStride]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveBasis]: Sets the curve basis function, determining how Metal interpolates the control points.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveBasis]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveEndCaps]: Configures the type of curve end caps.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveEndCaps]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveType]: Controls the curve type.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveType]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to control points in the control point buffers.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetIndexBuffer]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.IndexType]: Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetIndexType]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusBuffers]: Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusBuffers]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusFormat]: Sets the format of the radii in the radius buffer.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusFormat]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusStride]: Sets the stride, in bytes, between radii in the radius buffer.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusStride]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SegmentControlPointCount]: Controls the number of control points per curve segment.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SegmentCount]: Declares the number of curve segments.
//   - [MTL4AccelerationStructureMotionCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor
type MTL4AccelerationStructureMotionCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionCurveGeometryDescriptorFromID constructs a [MTL4AccelerationStructureMotionCurveGeometryDescriptor] from an objc.ID.
//
// Describes motion curve geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionCurveGeometryDescriptorFromID(id objc.ID) MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	return MTL4AccelerationStructureMotionCurveGeometryDescriptor{MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTL4AccelerationStructureMotionCurveGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointBuffers]: Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointBuffers]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointCount]: Specifies the number of control points in the buffers the control point buffers reference.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointCount]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointFormat]: Declares the format of the control points in the buffers that the control point buffers reference.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointFormat]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.ControlPointStride]: Sets the stride, in bytes, between control points in the control point buffer.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetControlPointStride]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveBasis]: Sets the curve basis function, determining how Metal interpolates the control points.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveBasis]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveEndCaps]: Configures the type of curve end caps.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveEndCaps]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.CurveType]: Controls the curve type.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetCurveType]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.IndexBuffer]: Assigns an optional index buffer containing references to control points in the control point buffers.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetIndexBuffer]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.IndexType]: Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetIndexType]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusBuffers]: Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusBuffers]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusFormat]: Sets the format of the radii in the radius buffer.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusFormat]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.RadiusStride]: Sets the stride, in bytes, between radii in the radius buffer.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetRadiusStride]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SegmentControlPointCount]: Controls the number of control points per curve segment.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SegmentCount]: Declares the number of curve segments.
//   - [IMTL4AccelerationStructureMotionCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor
type IMTL4AccelerationStructureMotionCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	// Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
	ControlPointBuffers() MTL4BufferRange
	SetControlPointBuffers(value MTL4BufferRange)
	// Specifies the number of control points in the buffers the control point buffers reference.
	ControlPointCount() uint
	SetControlPointCount(value uint)
	// Declares the format of the control points in the buffers that the control point buffers reference.
	ControlPointFormat() MTLAttributeFormat
	SetControlPointFormat(value MTLAttributeFormat)
	// Sets the stride, in bytes, between control points in the control point buffer.
	ControlPointStride() uint
	SetControlPointStride(value uint)
	// Sets the curve basis function, determining how Metal interpolates the control points.
	CurveBasis() MTLCurveBasis
	SetCurveBasis(value MTLCurveBasis)
	// Configures the type of curve end caps.
	CurveEndCaps() MTLCurveEndCaps
	SetCurveEndCaps(value MTLCurveEndCaps)
	// Controls the curve type.
	CurveType() MTLCurveType
	SetCurveType(value MTLCurveType)
	// Assigns an optional index buffer containing references to control points in the control point buffers.
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	// Configures the size of the indices the `indexBuffer` contains, which is typically either 16 or 32-bits for each index.
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	// Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
	RadiusBuffers() MTL4BufferRange
	SetRadiusBuffers(value MTL4BufferRange)
	// Sets the format of the radii in the radius buffer.
	RadiusFormat() MTLAttributeFormat
	SetRadiusFormat(value MTLAttributeFormat)
	// Sets the stride, in bytes, between radii in the radius buffer.
	RadiusStride() uint
	SetRadiusStride(value uint)
	// Controls the number of control points per curve segment.
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	// Declares the number of curve segments.
	SegmentCount() uint
	SetSegmentCount(value uint)
}





// Init initializes the instance.
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) Init() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionCurveGeometryDescriptor creates a new MTL4AccelerationStructureMotionCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionCurveGeometryDescriptor() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	class := getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Assigns a reference to a buffer where each entry contains a reference to a
// buffer of control points.
//
// # Discussion
// 
// This property references a buffer that conceptually represents an array
// with one entry for each keyframe in the motion animation. Each one of these
// entries consists of a [MTL4BufferRange] that, in turn, references a buffer
// containing the control points corresponding to the keyframe.
// 
// You are responsible for ensuring the buffer address is not zero for the
// top-level buffer, as well as for all the vertex buffers it references.
//
// [MTL4BufferRange]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("controlPointBuffers"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointBuffers:"), value)
}



// Specifies the number of control points in the buffers the control point
// buffers reference.
//
// # Discussion
// 
// All keyframes have the same number of control points.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("controlPointCount"))
	return rv
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointCount:"), value)
}



// Declares the format of the control points in the buffers that the control
// point buffers reference.
//
// # Discussion
// 
// All keyframes share the same control point format. Defaults to
// [MTLAttributeFormatFloat3], representing 3 floating point values tightly
// packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("controlPointFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointFormat:"), value)
}



// Sets the stride, in bytes, between control points in the control point
// buffer.
//
// # Discussion
// 
// All keyframes share the same control point stride.
// 
// You are responsible for ensuring this stride is a multiple of the control
// point format’s element size, and at a minimum exactly the control point
// format’s size.
// 
// This property defaults to `0`, indicating that the control points are
// tightly-packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("controlPointStride"))
	return rv
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setControlPointStride:"), value)
}



// Sets the curve basis function, determining how Metal interpolates the
// control points.
//
// # Discussion
// 
// Defaults to [MTLCurveBasisBSpline]. All keyframes share the same curve
// basis function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() MTLCurveBasis {
	rv := objc.Send[MTLCurveBasis](m.ID, objc.Sel("curveBasis"))
	return MTLCurveBasis(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value MTLCurveBasis) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveBasis:"), value)
}



// Configures the type of curve end caps.
//
// # Discussion
// 
// Defaults to [MTLCurveEndCapsNone]. All keyframes share the same end cap
// type.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() MTLCurveEndCaps {
	rv := objc.Send[MTLCurveEndCaps](m.ID, objc.Sel("curveEndCaps"))
	return MTLCurveEndCaps(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value MTLCurveEndCaps) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveEndCaps:"), value)
}



// Controls the curve type.
//
// # Discussion
// 
// Defaults to [MTLCurveTypeRound]. All keyframes share the same curve type.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveType
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveType() MTLCurveType {
	rv := objc.Send[MTLCurveType](m.ID, objc.Sel("curveType"))
	return MTLCurveType(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value MTLCurveType) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurveType:"), value)
}



// Assigns an optional index buffer containing references to control points in
// the control point buffers.
//
// # Discussion
// 
// All keyframes share the same index buffer, with each index representing the
// first control point of a curve segment.
// 
// You are responsible for ensuring the buffer address of the range is not
// zero.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("indexBuffer"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexBuffer:"), value)
}



// Configures the size of the indices the `indexBuffer` contains, which is
// typically either 16 or 32-bits for each index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexType
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](m.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}



// Assigns a reference to a buffer containing, in turn, references to curve
// radii buffers.
//
// # Discussion
// 
// This property references a buffer that conceptually represents an array
// with one entry for each keyframe in the motion animation. Each one of these
// entries consists of a [MTL4BufferRange] that, in turn, references a buffer
// containing the radii corresponding to the keyframe.
// 
// Metal interpolates curve radii according to the basis function you specify
// via [CurveBasis].
// 
// You are responsible for ensuring the type of each radius matches the type
// property [RadiusFormat] specifies, that each radius is at least zero, and
// that the buffer address of the top-level buffer, as well as of buffer it
// references, is not zero.
//
// [MTL4BufferRange]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() MTL4BufferRange {
	rv := objc.Send[MTL4BufferRange](m.ID, objc.Sel("radiusBuffers"))
	return MTL4BufferRange(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value MTL4BufferRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusBuffers:"), value)
}



// Sets the format of the radii in the radius buffer.
//
// # Discussion
// 
// Defaults to [MTLAttributeFormatFloat]. All keyframes share the same radius
// format.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](m.ID, objc.Sel("radiusFormat"))
	return MTLAttributeFormat(rv)
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusFormat:"), value)
}



// Sets the stride, in bytes, between radii in the radius buffer.
//
// # Discussion
// 
// You are responsible for ensuring this property is set to a multiple of the
// size corresponding to the [RadiusFormat]. All keyframes share the same
// radius stride.
// 
// This property defaults to `0` bytes, indicating that the radii are tightly
// packed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("radiusStride"))
	return rv
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRadiusStride:"), value)
}



// Controls the number of control points per curve segment.
//
// # Discussion
// 
// Valid values for this property are `2`, `3`, or `4`. All keyframes have the
// same number of control points per curve segment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("segmentControlPointCount"))
	return rv
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSegmentControlPointCount:"), value)
}



// Declares the number of curve segments.
//
// # Discussion
// 
// All keyframes have the same number of curve segments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("segmentCount"))
	return rv
}
func (m MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSegmentCount:"), value)
}
























