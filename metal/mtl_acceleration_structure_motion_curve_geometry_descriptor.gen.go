// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructureMotionCurveGeometryDescriptor] class.
var (
	_MTLAccelerationStructureMotionCurveGeometryDescriptorClass     MTLAccelerationStructureMotionCurveGeometryDescriptorClass
	_MTLAccelerationStructureMotionCurveGeometryDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureMotionCurveGeometryDescriptorClass() MTLAccelerationStructureMotionCurveGeometryDescriptorClass {
	_MTLAccelerationStructureMotionCurveGeometryDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureMotionCurveGeometryDescriptorClass = MTLAccelerationStructureMotionCurveGeometryDescriptorClass{class: objc.GetClass("MTLAccelerationStructureMotionCurveGeometryDescriptor")}
	})
	return _MTLAccelerationStructureMotionCurveGeometryDescriptorClass
}

// GetMTLAccelerationStructureMotionCurveGeometryDescriptorClass returns the class object for MTLAccelerationStructureMotionCurveGeometryDescriptor.
func GetMTLAccelerationStructureMotionCurveGeometryDescriptorClass() MTLAccelerationStructureMotionCurveGeometryDescriptorClass {
	return getMTLAccelerationStructureMotionCurveGeometryDescriptorClass()
}

type MTLAccelerationStructureMotionCurveGeometryDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() MTLAccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointBuffers]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointBuffers]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointCount]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointCount]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointFormat]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointFormat]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointStride]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointStride]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.CurveBasis]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveBasis]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.CurveEndCaps]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveEndCaps]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.CurveType]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveType]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.IndexBuffer]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexBuffer]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.IndexBufferOffset]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexBufferOffset]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.IndexType]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexType]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusBuffers]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusBuffers]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusFormat]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusFormat]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusStride]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusStride]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SegmentControlPointCount]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SegmentCount]
//   - [MTLAccelerationStructureMotionCurveGeometryDescriptor.SetSegmentCount]
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor
type MTLAccelerationStructureMotionCurveGeometryDescriptor struct {
	MTLAccelerationStructureGeometryDescriptor
}

// MTLAccelerationStructureMotionCurveGeometryDescriptorFromID constructs a [MTLAccelerationStructureMotionCurveGeometryDescriptor] from an objc.ID.
func MTLAccelerationStructureMotionCurveGeometryDescriptorFromID(id objc.ID) MTLAccelerationStructureMotionCurveGeometryDescriptor {
	return MTLAccelerationStructureMotionCurveGeometryDescriptor{MTLAccelerationStructureGeometryDescriptor: MTLAccelerationStructureGeometryDescriptorFromID(id)}
}
// NOTE: MTLAccelerationStructureMotionCurveGeometryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureMotionCurveGeometryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointBuffers]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointBuffers]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointCount]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointCount]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointFormat]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointFormat]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.ControlPointStride]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetControlPointStride]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.CurveBasis]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveBasis]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.CurveEndCaps]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveEndCaps]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.CurveType]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetCurveType]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.IndexBuffer]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexBuffer]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.IndexBufferOffset]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexBufferOffset]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.IndexType]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetIndexType]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusBuffers]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusBuffers]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusFormat]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusFormat]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.RadiusStride]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetRadiusStride]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SegmentControlPointCount]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetSegmentControlPointCount]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SegmentCount]
//   - [IMTLAccelerationStructureMotionCurveGeometryDescriptor.SetSegmentCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor
type IMTLAccelerationStructureMotionCurveGeometryDescriptor interface {
	IMTLAccelerationStructureGeometryDescriptor

	// Topic: Instance Properties

	ControlPointBuffers() []MTLMotionKeyframeData
	SetControlPointBuffers(value []MTLMotionKeyframeData)
	ControlPointCount() uint
	SetControlPointCount(value uint)
	ControlPointFormat() MTLAttributeFormat
	SetControlPointFormat(value MTLAttributeFormat)
	ControlPointStride() uint
	SetControlPointStride(value uint)
	CurveBasis() MTLCurveBasis
	SetCurveBasis(value MTLCurveBasis)
	CurveEndCaps() MTLCurveEndCaps
	SetCurveEndCaps(value MTLCurveEndCaps)
	CurveType() MTLCurveType
	SetCurveType(value MTLCurveType)
	IndexBuffer() MTLBuffer
	SetIndexBuffer(value MTLBuffer)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexType() MTLIndexType
	SetIndexType(value MTLIndexType)
	RadiusBuffers() []MTLMotionKeyframeData
	SetRadiusBuffers(value []MTLMotionKeyframeData)
	RadiusFormat() MTLAttributeFormat
	SetRadiusFormat(value MTLAttributeFormat)
	RadiusStride() uint
	SetRadiusStride(value uint)
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	SegmentCount() uint
	SetSegmentCount(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) Init() MTLAccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionCurveGeometryDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) Autorelease() MTLAccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureMotionCurveGeometryDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureMotionCurveGeometryDescriptor creates a new MTLAccelerationStructureMotionCurveGeometryDescriptor instance.
func NewMTLAccelerationStructureMotionCurveGeometryDescriptor() MTLAccelerationStructureMotionCurveGeometryDescriptor {
	class := getMTLAccelerationStructureMotionCurveGeometryDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureMotionCurveGeometryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/descriptor
func (_MTLAccelerationStructureMotionCurveGeometryDescriptorClass MTLAccelerationStructureMotionCurveGeometryDescriptorClass) Descriptor() MTLAccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLAccelerationStructureMotionCurveGeometryDescriptorClass.class), objc.Sel("descriptor"))
	return MTLAccelerationStructureMotionCurveGeometryDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() []MTLMotionKeyframeData {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("controlPointBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLMotionKeyframeData {
		return MTLMotionKeyframeDataFromID(id)
	})
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value []MTLMotionKeyframeData) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointBuffers:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("controlPointCount"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointCount:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("controlPointFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointFormat:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("controlPointStride"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setControlPointStride:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() MTLCurveBasis {
	rv := objc.Send[MTLCurveBasis](a.ID, objc.Sel("curveBasis"))
	return MTLCurveBasis(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value MTLCurveBasis) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveBasis:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() MTLCurveEndCaps {
	rv := objc.Send[MTLCurveEndCaps](a.ID, objc.Sel("curveEndCaps"))
	return MTLCurveEndCaps(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value MTLCurveEndCaps) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveEndCaps:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveType
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveType() MTLCurveType {
	rv := objc.Send[MTLCurveType](a.ID, objc.Sel("curveType"))
	return MTLCurveType(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value MTLCurveType) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurveType:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("indexBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value MTLBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBuffer:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBufferOffset
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexBufferOffset:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexType
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexType() MTLIndexType {
	rv := objc.Send[MTLIndexType](a.ID, objc.Sel("indexType"))
	return MTLIndexType(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value MTLIndexType) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndexType:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() []MTLMotionKeyframeData {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("radiusBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLMotionKeyframeData {
		return MTLMotionKeyframeDataFromID(id)
	})
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value []MTLMotionKeyframeData) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusBuffers:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() MTLAttributeFormat {
	rv := objc.Send[MTLAttributeFormat](a.ID, objc.Sel("radiusFormat"))
	return MTLAttributeFormat(rv)
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value MTLAttributeFormat) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusFormat:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("radiusStride"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setRadiusStride:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("segmentControlPointCount"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSegmentControlPointCount:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("segmentCount"))
	return rv
}
func (a MTLAccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSegmentCount:"), value)
}

