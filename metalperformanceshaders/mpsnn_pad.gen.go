// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNPad] class.
var (
	_MPSNNPadClass     MPSNNPadClass
	_MPSNNPadClassOnce sync.Once
)

func getMPSNNPadClass() MPSNNPadClass {
	_MPSNNPadClassOnce.Do(func() {
		_MPSNNPadClass = MPSNNPadClass{class: objc.GetClass("MPSNNPad")}
	})
	return _MPSNNPadClass
}

// GetMPSNNPadClass returns the class object for MPSNNPad.
func GetMPSNNPadClass() MPSNNPadClass {
	return getMPSNNPadClass()
}

type MPSNNPadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNPadClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadClass) Alloc() MPSNNPad {
	rv := objc.Send[MPSNNPad](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfter]
//   - [MPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray]
//
// # Instance Properties
//
//   - [MPSNNPad.FillValue]
//   - [MPSNNPad.SetFillValue]
//   - [MPSNNPad.PaddingSizeAfter]
//   - [MPSNNPad.SetPaddingSizeAfter]
//   - [MPSNNPad.PaddingSizeBefore]
//   - [MPSNNPad.SetPaddingSizeBefore]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad
type MPSNNPad struct {
	MPSCNNKernel
}

// MPSNNPadFromID constructs a [MPSNNPad] from an objc.ID.
func MPSNNPadFromID(id objc.ID) MPSNNPad {
	return MPSNNPad{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNPad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNPad] class.
//
// # Initializers
//
//   - [IMPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfter]
//   - [IMPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray]
//
// # Instance Properties
//
//   - [IMPSNNPad.FillValue]
//   - [IMPSNNPad.SetFillValue]
//   - [IMPSNNPad.PaddingSizeAfter]
//   - [IMPSNNPad.SetPaddingSizeAfter]
//   - [IMPSNNPad.PaddingSizeBefore]
//   - [IMPSNNPad.SetPaddingSizeBefore]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad
type IMPSNNPad interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDevicePaddingSizeBeforePaddingSizeAfter(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate) MPSNNPad
	InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, fillValueArray foundation.NSData) MPSNNPad

	// Topic: Instance Properties

	FillValue() float32
	SetFillValue(value float32)
	PaddingSizeAfter() MPSImageCoordinate
	SetPaddingSizeAfter(value MPSImageCoordinate)
	PaddingSizeBefore() MPSImageCoordinate
	SetPaddingSizeBefore(value MPSImageCoordinate)
}

// Init initializes the instance.
func (p MPSNNPad) Init() MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPad) Autorelease() MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPad creates a new MPSNNPad instance.
func NewMPSNNPad() MPSNNPad {
	class := getMPSNNPadClass()
	rv := objc.Send[MPSNNPad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewPadWithCoder(aDecoder foundation.INSCoder) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNPadFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(coder:device:)
func NewPadWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNPadFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(device:)
func NewPadWithDevice(device metal.MTLDevice) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPadFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(device:paddingSizeBefore:paddingSizeAfter:)
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfter(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:"), device, paddingSizeBefore, paddingSizeAfter)
	return MPSNNPadFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(device:paddingSizeBefore:paddingSizeAfter:fillValueArray:)
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, fillValueArray foundation.NSData) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:"), device, paddingSizeBefore, paddingSizeAfter, fillValueArray)
	return MPSNNPadFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(device:paddingSizeBefore:paddingSizeAfter:)
func (p MPSNNPad) InitWithDevicePaddingSizeBeforePaddingSizeAfter(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:"), device, paddingSizeBefore, paddingSizeAfter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/init(device:paddingSizeBefore:paddingSizeAfter:fillValueArray:)
func (p MPSNNPad) InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device metal.MTLDevice, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, fillValueArray foundation.NSData) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:"), device, paddingSizeBefore, paddingSizeAfter, fillValueArray)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/fillValue
func (p MPSNNPad) FillValue() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("fillValue"))
	return rv
}
func (p MPSNNPad) SetFillValue(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setFillValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeAfter
func (p MPSNNPad) PaddingSizeAfter() MPSImageCoordinate {
	rv := objc.Send[MPSImageCoordinate](p.ID, objc.Sel("paddingSizeAfter"))
	return MPSImageCoordinate(rv)
}
func (p MPSNNPad) SetPaddingSizeAfter(value MPSImageCoordinate) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaddingSizeAfter:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPad/paddingSizeBefore
func (p MPSNNPad) PaddingSizeBefore() MPSImageCoordinate {
	rv := objc.Send[MPSImageCoordinate](p.ID, objc.Sel("paddingSizeBefore"))
	return MPSImageCoordinate(rv)
}
func (p MPSNNPad) SetPaddingSizeBefore(value MPSImageCoordinate) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaddingSizeBefore:"), value)
}
