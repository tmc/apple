// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSPolygonAccelerationStructure] class.
var (
	_MPSPolygonAccelerationStructureClass     MPSPolygonAccelerationStructureClass
	_MPSPolygonAccelerationStructureClassOnce sync.Once
)

func getMPSPolygonAccelerationStructureClass() MPSPolygonAccelerationStructureClass {
	_MPSPolygonAccelerationStructureClassOnce.Do(func() {
		_MPSPolygonAccelerationStructureClass = MPSPolygonAccelerationStructureClass{class: objc.GetClass("MPSPolygonAccelerationStructure")}
	})
	return _MPSPolygonAccelerationStructureClass
}

// GetMPSPolygonAccelerationStructureClass returns the class object for MPSPolygonAccelerationStructure.
func GetMPSPolygonAccelerationStructureClass() MPSPolygonAccelerationStructureClass {
	return getMPSPolygonAccelerationStructureClass()
}

type MPSPolygonAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSPolygonAccelerationStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSPolygonAccelerationStructureClass) Alloc() MPSPolygonAccelerationStructure {
	rv := objc.Send[MPSPolygonAccelerationStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSPolygonAccelerationStructure.IndexBuffer]
//   - [MPSPolygonAccelerationStructure.SetIndexBuffer]
//   - [MPSPolygonAccelerationStructure.IndexBufferOffset]
//   - [MPSPolygonAccelerationStructure.SetIndexBufferOffset]
//   - [MPSPolygonAccelerationStructure.IndexType]
//   - [MPSPolygonAccelerationStructure.SetIndexType]
//   - [MPSPolygonAccelerationStructure.MaskBuffer]
//   - [MPSPolygonAccelerationStructure.SetMaskBuffer]
//   - [MPSPolygonAccelerationStructure.MaskBufferOffset]
//   - [MPSPolygonAccelerationStructure.SetMaskBufferOffset]
//   - [MPSPolygonAccelerationStructure.PolygonBuffers]
//   - [MPSPolygonAccelerationStructure.SetPolygonBuffers]
//   - [MPSPolygonAccelerationStructure.PolygonCount]
//   - [MPSPolygonAccelerationStructure.SetPolygonCount]
//   - [MPSPolygonAccelerationStructure.PolygonType]
//   - [MPSPolygonAccelerationStructure.SetPolygonType]
//   - [MPSPolygonAccelerationStructure.VertexBuffer]
//   - [MPSPolygonAccelerationStructure.SetVertexBuffer]
//   - [MPSPolygonAccelerationStructure.VertexBufferOffset]
//   - [MPSPolygonAccelerationStructure.SetVertexBufferOffset]
//   - [MPSPolygonAccelerationStructure.VertexStride]
//   - [MPSPolygonAccelerationStructure.SetVertexStride]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure
type MPSPolygonAccelerationStructure struct {
	MPSAccelerationStructure
}

// MPSPolygonAccelerationStructureFromID constructs a [MPSPolygonAccelerationStructure] from an objc.ID.
func MPSPolygonAccelerationStructureFromID(id objc.ID) MPSPolygonAccelerationStructure {
	return MPSPolygonAccelerationStructure{MPSAccelerationStructure: MPSAccelerationStructureFromID(id)}
}

// NOTE: MPSPolygonAccelerationStructure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSPolygonAccelerationStructure] class.
//
// # Instance Properties
//
//   - [IMPSPolygonAccelerationStructure.IndexBuffer]
//   - [IMPSPolygonAccelerationStructure.SetIndexBuffer]
//   - [IMPSPolygonAccelerationStructure.IndexBufferOffset]
//   - [IMPSPolygonAccelerationStructure.SetIndexBufferOffset]
//   - [IMPSPolygonAccelerationStructure.IndexType]
//   - [IMPSPolygonAccelerationStructure.SetIndexType]
//   - [IMPSPolygonAccelerationStructure.MaskBuffer]
//   - [IMPSPolygonAccelerationStructure.SetMaskBuffer]
//   - [IMPSPolygonAccelerationStructure.MaskBufferOffset]
//   - [IMPSPolygonAccelerationStructure.SetMaskBufferOffset]
//   - [IMPSPolygonAccelerationStructure.PolygonBuffers]
//   - [IMPSPolygonAccelerationStructure.SetPolygonBuffers]
//   - [IMPSPolygonAccelerationStructure.PolygonCount]
//   - [IMPSPolygonAccelerationStructure.SetPolygonCount]
//   - [IMPSPolygonAccelerationStructure.PolygonType]
//   - [IMPSPolygonAccelerationStructure.SetPolygonType]
//   - [IMPSPolygonAccelerationStructure.VertexBuffer]
//   - [IMPSPolygonAccelerationStructure.SetVertexBuffer]
//   - [IMPSPolygonAccelerationStructure.VertexBufferOffset]
//   - [IMPSPolygonAccelerationStructure.SetVertexBufferOffset]
//   - [IMPSPolygonAccelerationStructure.VertexStride]
//   - [IMPSPolygonAccelerationStructure.SetVertexStride]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure
type IMPSPolygonAccelerationStructure interface {
	IMPSAccelerationStructure

	// Topic: Instance Properties

	IndexBuffer() metal.MTLBuffer
	SetIndexBuffer(value metal.MTLBuffer)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexType() MPSDataType
	SetIndexType(value MPSDataType)
	MaskBuffer() metal.MTLBuffer
	SetMaskBuffer(value metal.MTLBuffer)
	MaskBufferOffset() uint
	SetMaskBufferOffset(value uint)
	PolygonBuffers() []MPSPolygonBuffer
	SetPolygonBuffers(value []MPSPolygonBuffer)
	PolygonCount() uint
	SetPolygonCount(value uint)
	PolygonType() MPSPolygonType
	SetPolygonType(value MPSPolygonType)
	VertexBuffer() metal.MTLBuffer
	SetVertexBuffer(value metal.MTLBuffer)
	VertexBufferOffset() uint
	SetVertexBufferOffset(value uint)
	VertexStride() uint
	SetVertexStride(value uint)
}

// Init initializes the instance.
func (p MPSPolygonAccelerationStructure) Init() MPSPolygonAccelerationStructure {
	rv := objc.Send[MPSPolygonAccelerationStructure](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSPolygonAccelerationStructure) Autorelease() MPSPolygonAccelerationStructure {
	rv := objc.Send[MPSPolygonAccelerationStructure](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSPolygonAccelerationStructure creates a new MPSPolygonAccelerationStructure instance.
func NewMPSPolygonAccelerationStructure() MPSPolygonAccelerationStructure {
	class := getMPSPolygonAccelerationStructureClass()
	rv := objc.Send[MPSPolygonAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewPolygonAccelerationStructureWithCoder(aDecoder foundation.INSCoder) MPSPolygonAccelerationStructure {
	instance := getMPSPolygonAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSPolygonAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:device:)
func NewPolygonAccelerationStructureWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSPolygonAccelerationStructure {
	instance := getMPSPolygonAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSPolygonAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(coder:group:)
func NewPolygonAccelerationStructureWithCoderGroup(aDecoder foundation.INSCoder, group IMPSAccelerationStructureGroup) MPSPolygonAccelerationStructure {
	instance := getMPSPolygonAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	return MPSPolygonAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(device:)
func NewPolygonAccelerationStructureWithDevice(device metal.MTLDevice) MPSPolygonAccelerationStructure {
	instance := getMPSPolygonAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSPolygonAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewPolygonAccelerationStructureWithGroup(group IMPSAccelerationStructureGroup) MPSPolygonAccelerationStructure {
	instance := getMPSPolygonAccelerationStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroup:"), group)
	return MPSPolygonAccelerationStructureFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/indexBuffer
func (p MPSPolygonAccelerationStructure) IndexBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("indexBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonAccelerationStructure) SetIndexBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/indexBufferOffset
func (p MPSPolygonAccelerationStructure) IndexBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (p MPSPolygonAccelerationStructure) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/indexType
func (p MPSPolygonAccelerationStructure) IndexType() MPSDataType {
	rv := objc.Send[MPSDataType](p.ID, objc.Sel("indexType"))
	return MPSDataType(rv)
}
func (p MPSPolygonAccelerationStructure) SetIndexType(value MPSDataType) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/maskBuffer
func (p MPSPolygonAccelerationStructure) MaskBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("maskBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonAccelerationStructure) SetMaskBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaskBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/maskBufferOffset
func (p MPSPolygonAccelerationStructure) MaskBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("maskBufferOffset"))
	return rv
}
func (p MPSPolygonAccelerationStructure) SetMaskBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaskBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonBuffers
func (p MPSPolygonAccelerationStructure) PolygonBuffers() []MPSPolygonBuffer {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("polygonBuffers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSPolygonBuffer {
		return MPSPolygonBufferFromID(id)
	})
}
func (p MPSPolygonAccelerationStructure) SetPolygonBuffers(value []MPSPolygonBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setPolygonBuffers:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonCount
func (p MPSPolygonAccelerationStructure) PolygonCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("polygonCount"))
	return rv
}
func (p MPSPolygonAccelerationStructure) SetPolygonCount(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setPolygonCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonType
func (p MPSPolygonAccelerationStructure) PolygonType() MPSPolygonType {
	rv := objc.Send[MPSPolygonType](p.ID, objc.Sel("polygonType"))
	return MPSPolygonType(rv)
}
func (p MPSPolygonAccelerationStructure) SetPolygonType(value MPSPolygonType) {
	objc.Send[struct{}](p.ID, objc.Sel("setPolygonType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/vertexBuffer
func (p MPSPolygonAccelerationStructure) VertexBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("vertexBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonAccelerationStructure) SetVertexBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setVertexBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/vertexBufferOffset
func (p MPSPolygonAccelerationStructure) VertexBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("vertexBufferOffset"))
	return rv
}
func (p MPSPolygonAccelerationStructure) SetVertexBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setVertexBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/vertexStride
func (p MPSPolygonAccelerationStructure) VertexStride() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("vertexStride"))
	return rv
}
func (p MPSPolygonAccelerationStructure) SetVertexStride(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setVertexStride:"), value)
}
