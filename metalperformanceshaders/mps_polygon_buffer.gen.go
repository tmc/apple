// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSPolygonBuffer] class.
var (
	_MPSPolygonBufferClass     MPSPolygonBufferClass
	_MPSPolygonBufferClassOnce sync.Once
)

func getMPSPolygonBufferClass() MPSPolygonBufferClass {
	_MPSPolygonBufferClassOnce.Do(func() {
		_MPSPolygonBufferClass = MPSPolygonBufferClass{class: objc.GetClass("MPSPolygonBuffer")}
	})
	return _MPSPolygonBufferClass
}

// GetMPSPolygonBufferClass returns the class object for MPSPolygonBuffer.
func GetMPSPolygonBufferClass() MPSPolygonBufferClass {
	return getMPSPolygonBufferClass()
}

type MPSPolygonBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSPolygonBufferClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSPolygonBufferClass) Alloc() MPSPolygonBuffer {
	rv := objc.Send[MPSPolygonBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSPolygonBuffer.IndexBuffer]
//   - [MPSPolygonBuffer.SetIndexBuffer]
//   - [MPSPolygonBuffer.IndexBufferOffset]
//   - [MPSPolygonBuffer.SetIndexBufferOffset]
//   - [MPSPolygonBuffer.MaskBuffer]
//   - [MPSPolygonBuffer.SetMaskBuffer]
//   - [MPSPolygonBuffer.MaskBufferOffset]
//   - [MPSPolygonBuffer.SetMaskBufferOffset]
//   - [MPSPolygonBuffer.PolygonCount]
//   - [MPSPolygonBuffer.SetPolygonCount]
//   - [MPSPolygonBuffer.VertexBuffer]
//   - [MPSPolygonBuffer.SetVertexBuffer]
//   - [MPSPolygonBuffer.VertexBufferOffset]
//   - [MPSPolygonBuffer.SetVertexBufferOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer
type MPSPolygonBuffer struct {
	objectivec.Object
}

// MPSPolygonBufferFromID constructs a [MPSPolygonBuffer] from an objc.ID.
func MPSPolygonBufferFromID(id objc.ID) MPSPolygonBuffer {
	return MPSPolygonBuffer{objectivec.Object{ID: id}}
}

// NOTE: MPSPolygonBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSPolygonBuffer] class.
//
// # Instance Properties
//
//   - [IMPSPolygonBuffer.IndexBuffer]
//   - [IMPSPolygonBuffer.SetIndexBuffer]
//   - [IMPSPolygonBuffer.IndexBufferOffset]
//   - [IMPSPolygonBuffer.SetIndexBufferOffset]
//   - [IMPSPolygonBuffer.MaskBuffer]
//   - [IMPSPolygonBuffer.SetMaskBuffer]
//   - [IMPSPolygonBuffer.MaskBufferOffset]
//   - [IMPSPolygonBuffer.SetMaskBufferOffset]
//   - [IMPSPolygonBuffer.PolygonCount]
//   - [IMPSPolygonBuffer.SetPolygonCount]
//   - [IMPSPolygonBuffer.VertexBuffer]
//   - [IMPSPolygonBuffer.SetVertexBuffer]
//   - [IMPSPolygonBuffer.VertexBufferOffset]
//   - [IMPSPolygonBuffer.SetVertexBufferOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer
type IMPSPolygonBuffer interface {
	objectivec.IObject

	// Topic: Instance Properties

	IndexBuffer() metal.MTLBuffer
	SetIndexBuffer(value metal.MTLBuffer)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	MaskBuffer() metal.MTLBuffer
	SetMaskBuffer(value metal.MTLBuffer)
	MaskBufferOffset() uint
	SetMaskBufferOffset(value uint)
	PolygonCount() uint
	SetPolygonCount(value uint)
	VertexBuffer() metal.MTLBuffer
	SetVertexBuffer(value metal.MTLBuffer)
	VertexBufferOffset() uint
	SetVertexBufferOffset(value uint)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p MPSPolygonBuffer) Init() MPSPolygonBuffer {
	rv := objc.Send[MPSPolygonBuffer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSPolygonBuffer) Autorelease() MPSPolygonBuffer {
	rv := objc.Send[MPSPolygonBuffer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSPolygonBuffer creates a new MPSPolygonBuffer instance.
func NewMPSPolygonBuffer() MPSPolygonBuffer {
	class := getMPSPolygonBufferClass()
	rv := objc.Send[MPSPolygonBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/init(coder:)
func NewPolygonBufferWithCoder(aDecoder foundation.INSCoder) MPSPolygonBuffer {
	instance := getMPSPolygonBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSPolygonBufferFromID(rv)
}

func (p MPSPolygonBuffer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/indexBuffer
func (p MPSPolygonBuffer) IndexBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("indexBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonBuffer) SetIndexBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/indexBufferOffset
func (p MPSPolygonBuffer) IndexBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("indexBufferOffset"))
	return rv
}
func (p MPSPolygonBuffer) SetIndexBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/maskBuffer
func (p MPSPolygonBuffer) MaskBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("maskBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonBuffer) SetMaskBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaskBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/maskBufferOffset
func (p MPSPolygonBuffer) MaskBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("maskBufferOffset"))
	return rv
}
func (p MPSPolygonBuffer) SetMaskBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaskBufferOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/polygonCount
func (p MPSPolygonBuffer) PolygonCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("polygonCount"))
	return rv
}
func (p MPSPolygonBuffer) SetPolygonCount(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setPolygonCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/vertexBuffer
func (p MPSPolygonBuffer) VertexBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("vertexBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}
func (p MPSPolygonBuffer) SetVertexBuffer(value metal.MTLBuffer) {
	objc.Send[struct{}](p.ID, objc.Sel("setVertexBuffer:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/vertexBufferOffset
func (p MPSPolygonBuffer) VertexBufferOffset() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("vertexBufferOffset"))
	return rv
}
func (p MPSPolygonBuffer) SetVertexBufferOffset(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setVertexBufferOffset:"), value)
}
