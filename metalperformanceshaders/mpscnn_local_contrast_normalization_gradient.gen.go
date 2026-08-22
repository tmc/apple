// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLocalContrastNormalizationGradient] class.
var (
	_MPSCNNLocalContrastNormalizationGradientClass     MPSCNNLocalContrastNormalizationGradientClass
	_MPSCNNLocalContrastNormalizationGradientClassOnce sync.Once
)

func getMPSCNNLocalContrastNormalizationGradientClass() MPSCNNLocalContrastNormalizationGradientClass {
	_MPSCNNLocalContrastNormalizationGradientClassOnce.Do(func() {
		_MPSCNNLocalContrastNormalizationGradientClass = MPSCNNLocalContrastNormalizationGradientClass{class: objc.GetClass("MPSCNNLocalContrastNormalizationGradient")}
	})
	return _MPSCNNLocalContrastNormalizationGradientClass
}

// GetMPSCNNLocalContrastNormalizationGradientClass returns the class object for MPSCNNLocalContrastNormalizationGradient.
func GetMPSCNNLocalContrastNormalizationGradientClass() MPSCNNLocalContrastNormalizationGradientClass {
	return getMPSCNNLocalContrastNormalizationGradientClass()
}

type MPSCNNLocalContrastNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLocalContrastNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLocalContrastNormalizationGradientClass) Alloc() MPSCNNLocalContrastNormalizationGradient {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient local-contrast normalization kernel.
//
// # Initializers
//
//   - [MPSCNNLocalContrastNormalizationGradient.InitWithDeviceKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [MPSCNNLocalContrastNormalizationGradient.Alpha]
//   - [MPSCNNLocalContrastNormalizationGradient.SetAlpha]
//   - [MPSCNNLocalContrastNormalizationGradient.Beta]
//   - [MPSCNNLocalContrastNormalizationGradient.SetBeta]
//   - [MPSCNNLocalContrastNormalizationGradient.Delta]
//   - [MPSCNNLocalContrastNormalizationGradient.SetDelta]
//   - [MPSCNNLocalContrastNormalizationGradient.P0]
//   - [MPSCNNLocalContrastNormalizationGradient.SetP0]
//   - [MPSCNNLocalContrastNormalizationGradient.Pm]
//   - [MPSCNNLocalContrastNormalizationGradient.SetPm]
//   - [MPSCNNLocalContrastNormalizationGradient.Ps]
//   - [MPSCNNLocalContrastNormalizationGradient.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient
type MPSCNNLocalContrastNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNLocalContrastNormalizationGradientFromID constructs a [MPSCNNLocalContrastNormalizationGradient] from an objc.ID.
//
// A gradient local-contrast normalization kernel.
func MPSCNNLocalContrastNormalizationGradientFromID(id objc.ID) MPSCNNLocalContrastNormalizationGradient {
	return MPSCNNLocalContrastNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNLocalContrastNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLocalContrastNormalizationGradient] class.
//
// # Initializers
//
//   - [IMPSCNNLocalContrastNormalizationGradient.InitWithDeviceKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [IMPSCNNLocalContrastNormalizationGradient.Alpha]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetAlpha]
//   - [IMPSCNNLocalContrastNormalizationGradient.Beta]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetBeta]
//   - [IMPSCNNLocalContrastNormalizationGradient.Delta]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetDelta]
//   - [IMPSCNNLocalContrastNormalizationGradient.P0]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetP0]
//   - [IMPSCNNLocalContrastNormalizationGradient.Pm]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetPm]
//   - [IMPSCNNLocalContrastNormalizationGradient.Ps]
//   - [IMPSCNNLocalContrastNormalizationGradient.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient
type IMPSCNNLocalContrastNormalizationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradient

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
	P0() float32
	SetP0(value float32)
	Pm() float32
	SetPm(value float32)
	Ps() float32
	SetPs(value float32)
}

// Init initializes the instance.
func (c MPSCNNLocalContrastNormalizationGradient) Init() MPSCNNLocalContrastNormalizationGradient {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLocalContrastNormalizationGradient) Autorelease() MPSCNNLocalContrastNormalizationGradient {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLocalContrastNormalizationGradient creates a new MPSCNNLocalContrastNormalizationGradient instance.
func NewMPSCNNLocalContrastNormalizationGradient() MPSCNNLocalContrastNormalizationGradient {
	class := getMPSCNNLocalContrastNormalizationGradientClass()
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNLocalContrastNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNLocalContrastNormalizationGradient {
	instance := getMPSCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNLocalContrastNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/init(coder:device:)
func NewCNNLocalContrastNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNLocalContrastNormalizationGradient {
	instance := getMPSCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNLocalContrastNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNLocalContrastNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNLocalContrastNormalizationGradient {
	instance := getMPSCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNLocalContrastNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNLocalContrastNormalizationGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradient {
	instance := getMPSCNNLocalContrastNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNLocalContrastNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNLocalContrastNormalizationGradient) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradient {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradient](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/alpha
func (c MPSCNNLocalContrastNormalizationGradient) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/beta
func (c MPSCNNLocalContrastNormalizationGradient) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/delta
func (c MPSCNNLocalContrastNormalizationGradient) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/p0
func (c MPSCNNLocalContrastNormalizationGradient) P0() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("p0"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetP0(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setP0:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/pm
func (c MPSCNNLocalContrastNormalizationGradient) Pm() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pm"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetPm(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPm:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradient/ps
func (c MPSCNNLocalContrastNormalizationGradient) Ps() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("ps"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradient) SetPs(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPs:"), value)
}
