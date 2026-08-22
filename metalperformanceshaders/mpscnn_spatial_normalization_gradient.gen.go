// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSpatialNormalizationGradient] class.
var (
	_MPSCNNSpatialNormalizationGradientClass     MPSCNNSpatialNormalizationGradientClass
	_MPSCNNSpatialNormalizationGradientClassOnce sync.Once
)

func getMPSCNNSpatialNormalizationGradientClass() MPSCNNSpatialNormalizationGradientClass {
	_MPSCNNSpatialNormalizationGradientClassOnce.Do(func() {
		_MPSCNNSpatialNormalizationGradientClass = MPSCNNSpatialNormalizationGradientClass{class: objc.GetClass("MPSCNNSpatialNormalizationGradient")}
	})
	return _MPSCNNSpatialNormalizationGradientClass
}

// GetMPSCNNSpatialNormalizationGradientClass returns the class object for MPSCNNSpatialNormalizationGradient.
func GetMPSCNNSpatialNormalizationGradientClass() MPSCNNSpatialNormalizationGradientClass {
	return getMPSCNNSpatialNormalizationGradientClass()
}

type MPSCNNSpatialNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSpatialNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSpatialNormalizationGradientClass) Alloc() MPSCNNSpatialNormalizationGradient {
	rv := objc.Send[MPSCNNSpatialNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient spatial normalization kernel.
//
// # Initializers
//
//   - [MPSCNNSpatialNormalizationGradient.InitWithDeviceKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [MPSCNNSpatialNormalizationGradient.Alpha]
//   - [MPSCNNSpatialNormalizationGradient.SetAlpha]
//   - [MPSCNNSpatialNormalizationGradient.Beta]
//   - [MPSCNNSpatialNormalizationGradient.SetBeta]
//   - [MPSCNNSpatialNormalizationGradient.Delta]
//   - [MPSCNNSpatialNormalizationGradient.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient
type MPSCNNSpatialNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNSpatialNormalizationGradientFromID constructs a [MPSCNNSpatialNormalizationGradient] from an objc.ID.
//
// A gradient spatial normalization kernel.
func MPSCNNSpatialNormalizationGradientFromID(id objc.ID) MPSCNNSpatialNormalizationGradient {
	return MPSCNNSpatialNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNSpatialNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSpatialNormalizationGradient] class.
//
// # Initializers
//
//   - [IMPSCNNSpatialNormalizationGradient.InitWithDeviceKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [IMPSCNNSpatialNormalizationGradient.Alpha]
//   - [IMPSCNNSpatialNormalizationGradient.SetAlpha]
//   - [IMPSCNNSpatialNormalizationGradient.Beta]
//   - [IMPSCNNSpatialNormalizationGradient.SetBeta]
//   - [IMPSCNNSpatialNormalizationGradient.Delta]
//   - [IMPSCNNSpatialNormalizationGradient.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient
type IMPSCNNSpatialNormalizationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalizationGradient

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
}

// Init initializes the instance.
func (c MPSCNNSpatialNormalizationGradient) Init() MPSCNNSpatialNormalizationGradient {
	rv := objc.Send[MPSCNNSpatialNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSpatialNormalizationGradient) Autorelease() MPSCNNSpatialNormalizationGradient {
	rv := objc.Send[MPSCNNSpatialNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSpatialNormalizationGradient creates a new MPSCNNSpatialNormalizationGradient instance.
func NewMPSCNNSpatialNormalizationGradient() MPSCNNSpatialNormalizationGradient {
	class := getMPSCNNSpatialNormalizationGradientClass()
	rv := objc.Send[MPSCNNSpatialNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSpatialNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNSpatialNormalizationGradient {
	instance := getMPSCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSpatialNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/init(coder:device:)
func NewCNNSpatialNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSpatialNormalizationGradient {
	instance := getMPSCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSpatialNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNSpatialNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNSpatialNormalizationGradient {
	instance := getMPSCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSpatialNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNSpatialNormalizationGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalizationGradient {
	instance := getMPSCNNSpatialNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNSpatialNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNSpatialNormalizationGradient) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalizationGradient {
	rv := objc.Send[MPSCNNSpatialNormalizationGradient](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/alpha
func (c MPSCNNSpatialNormalizationGradient) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradient) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/beta
func (c MPSCNNSpatialNormalizationGradient) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradient) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradient/delta
func (c MPSCNNSpatialNormalizationGradient) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradient) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}
