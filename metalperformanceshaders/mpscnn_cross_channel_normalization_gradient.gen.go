// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNCrossChannelNormalizationGradient] class.
var (
	_MPSCNNCrossChannelNormalizationGradientClass     MPSCNNCrossChannelNormalizationGradientClass
	_MPSCNNCrossChannelNormalizationGradientClassOnce sync.Once
)

func getMPSCNNCrossChannelNormalizationGradientClass() MPSCNNCrossChannelNormalizationGradientClass {
	_MPSCNNCrossChannelNormalizationGradientClassOnce.Do(func() {
		_MPSCNNCrossChannelNormalizationGradientClass = MPSCNNCrossChannelNormalizationGradientClass{class: objc.GetClass("MPSCNNCrossChannelNormalizationGradient")}
	})
	return _MPSCNNCrossChannelNormalizationGradientClass
}

// GetMPSCNNCrossChannelNormalizationGradientClass returns the class object for MPSCNNCrossChannelNormalizationGradient.
func GetMPSCNNCrossChannelNormalizationGradientClass() MPSCNNCrossChannelNormalizationGradientClass {
	return getMPSCNNCrossChannelNormalizationGradientClass()
}

type MPSCNNCrossChannelNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNCrossChannelNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNCrossChannelNormalizationGradientClass) Alloc() MPSCNNCrossChannelNormalizationGradient {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient normalization kernel applied across feature channels.
//
// # Initializers
//
//   - [MPSCNNCrossChannelNormalizationGradient.InitWithDeviceKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNCrossChannelNormalizationGradient.Alpha]
//   - [MPSCNNCrossChannelNormalizationGradient.SetAlpha]
//   - [MPSCNNCrossChannelNormalizationGradient.Beta]
//   - [MPSCNNCrossChannelNormalizationGradient.SetBeta]
//   - [MPSCNNCrossChannelNormalizationGradient.Delta]
//   - [MPSCNNCrossChannelNormalizationGradient.SetDelta]
//   - [MPSCNNCrossChannelNormalizationGradient.KernelSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient
type MPSCNNCrossChannelNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNCrossChannelNormalizationGradientFromID constructs a [MPSCNNCrossChannelNormalizationGradient] from an objc.ID.
//
// A gradient normalization kernel applied across feature channels.
func MPSCNNCrossChannelNormalizationGradientFromID(id objc.ID) MPSCNNCrossChannelNormalizationGradient {
	return MPSCNNCrossChannelNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNCrossChannelNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNCrossChannelNormalizationGradient] class.
//
// # Initializers
//
//   - [IMPSCNNCrossChannelNormalizationGradient.InitWithDeviceKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNCrossChannelNormalizationGradient.Alpha]
//   - [IMPSCNNCrossChannelNormalizationGradient.SetAlpha]
//   - [IMPSCNNCrossChannelNormalizationGradient.Beta]
//   - [IMPSCNNCrossChannelNormalizationGradient.SetBeta]
//   - [IMPSCNNCrossChannelNormalizationGradient.Delta]
//   - [IMPSCNNCrossChannelNormalizationGradient.SetDelta]
//   - [IMPSCNNCrossChannelNormalizationGradient.KernelSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient
type IMPSCNNCrossChannelNormalizationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalizationGradient

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
	KernelSize() uint
}

// Init initializes the instance.
func (c MPSCNNCrossChannelNormalizationGradient) Init() MPSCNNCrossChannelNormalizationGradient {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNCrossChannelNormalizationGradient) Autorelease() MPSCNNCrossChannelNormalizationGradient {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNCrossChannelNormalizationGradient creates a new MPSCNNCrossChannelNormalizationGradient instance.
func NewMPSCNNCrossChannelNormalizationGradient() MPSCNNCrossChannelNormalizationGradient {
	class := getMPSCNNCrossChannelNormalizationGradientClass()
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNCrossChannelNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNCrossChannelNormalizationGradient {
	instance := getMPSCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNCrossChannelNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/init(coder:device:)
func NewCNNCrossChannelNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNCrossChannelNormalizationGradient {
	instance := getMPSCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNCrossChannelNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNCrossChannelNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNCrossChannelNormalizationGradient {
	instance := getMPSCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNCrossChannelNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/init(device:kernelSize:)
func NewCNNCrossChannelNormalizationGradientWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalizationGradient {
	instance := getMPSCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	return MPSCNNCrossChannelNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/init(device:kernelSize:)
func (c MPSCNNCrossChannelNormalizationGradient) InitWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalizationGradient {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradient](c.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/alpha
func (c MPSCNNCrossChannelNormalizationGradient) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNCrossChannelNormalizationGradient) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/beta
func (c MPSCNNCrossChannelNormalizationGradient) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNCrossChannelNormalizationGradient) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/delta
func (c MPSCNNCrossChannelNormalizationGradient) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNCrossChannelNormalizationGradient) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient/kernelSize
func (c MPSCNNCrossChannelNormalizationGradient) KernelSize() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelSize"))
	return rv
}
