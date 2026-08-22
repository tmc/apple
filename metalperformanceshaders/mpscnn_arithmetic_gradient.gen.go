// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNArithmeticGradient] class.
var (
	_MPSCNNArithmeticGradientClass     MPSCNNArithmeticGradientClass
	_MPSCNNArithmeticGradientClassOnce sync.Once
)

func getMPSCNNArithmeticGradientClass() MPSCNNArithmeticGradientClass {
	_MPSCNNArithmeticGradientClassOnce.Do(func() {
		_MPSCNNArithmeticGradientClass = MPSCNNArithmeticGradientClass{class: objc.GetClass("MPSCNNArithmeticGradient")}
	})
	return _MPSCNNArithmeticGradientClass
}

// GetMPSCNNArithmeticGradientClass returns the class object for MPSCNNArithmeticGradient.
func GetMPSCNNArithmeticGradientClass() MPSCNNArithmeticGradientClass {
	return getMPSCNNArithmeticGradientClass()
}

type MPSCNNArithmeticGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNArithmeticGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNArithmeticGradientClass) Alloc() MPSCNNArithmeticGradient {
	rv := objc.Send[MPSCNNArithmeticGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for gradient arithmetic operators.
//
// # Instance Properties
//
//   - [MPSCNNArithmeticGradient.Bias]
//   - [MPSCNNArithmeticGradient.SetBias]
//   - [MPSCNNArithmeticGradient.IsSecondarySourceFilter]
//   - [MPSCNNArithmeticGradient.MaximumValue]
//   - [MPSCNNArithmeticGradient.SetMaximumValue]
//   - [MPSCNNArithmeticGradient.MinimumValue]
//   - [MPSCNNArithmeticGradient.SetMinimumValue]
//   - [MPSCNNArithmeticGradient.PrimaryScale]
//   - [MPSCNNArithmeticGradient.SetPrimaryScale]
//   - [MPSCNNArithmeticGradient.SecondaryScale]
//   - [MPSCNNArithmeticGradient.SetSecondaryScale]
//   - [MPSCNNArithmeticGradient.SecondaryStrideInFeatureChannels]
//   - [MPSCNNArithmeticGradient.SetSecondaryStrideInFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient
type MPSCNNArithmeticGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNArithmeticGradientFromID constructs a [MPSCNNArithmeticGradient] from an objc.ID.
//
// The base class for gradient arithmetic operators.
func MPSCNNArithmeticGradientFromID(id objc.ID) MPSCNNArithmeticGradient {
	return MPSCNNArithmeticGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNArithmeticGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNArithmeticGradient] class.
//
// # Instance Properties
//
//   - [IMPSCNNArithmeticGradient.Bias]
//   - [IMPSCNNArithmeticGradient.SetBias]
//   - [IMPSCNNArithmeticGradient.IsSecondarySourceFilter]
//   - [IMPSCNNArithmeticGradient.MaximumValue]
//   - [IMPSCNNArithmeticGradient.SetMaximumValue]
//   - [IMPSCNNArithmeticGradient.MinimumValue]
//   - [IMPSCNNArithmeticGradient.SetMinimumValue]
//   - [IMPSCNNArithmeticGradient.PrimaryScale]
//   - [IMPSCNNArithmeticGradient.SetPrimaryScale]
//   - [IMPSCNNArithmeticGradient.SecondaryScale]
//   - [IMPSCNNArithmeticGradient.SetSecondaryScale]
//   - [IMPSCNNArithmeticGradient.SecondaryStrideInFeatureChannels]
//   - [IMPSCNNArithmeticGradient.SetSecondaryStrideInFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient
type IMPSCNNArithmeticGradient interface {
	IMPSCNNGradientKernel

	// Topic: Instance Properties

	Bias() float32
	SetBias(value float32)
	IsSecondarySourceFilter() bool
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
}

// Init initializes the instance.
func (c MPSCNNArithmeticGradient) Init() MPSCNNArithmeticGradient {
	rv := objc.Send[MPSCNNArithmeticGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNArithmeticGradient) Autorelease() MPSCNNArithmeticGradient {
	rv := objc.Send[MPSCNNArithmeticGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNArithmeticGradient creates a new MPSCNNArithmeticGradient instance.
func NewMPSCNNArithmeticGradient() MPSCNNArithmeticGradient {
	class := getMPSCNNArithmeticGradientClass()
	rv := objc.Send[MPSCNNArithmeticGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNArithmeticGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNArithmeticGradient {
	instance := getMPSCNNArithmeticGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNArithmeticGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNArithmeticGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNArithmeticGradient {
	instance := getMPSCNNArithmeticGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNArithmeticGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNArithmeticGradientWithDevice(device metal.MTLDevice) MPSCNNArithmeticGradient {
	instance := getMPSCNNArithmeticGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNArithmeticGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/bias
func (c MPSCNNArithmeticGradient) Bias() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("bias"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetBias(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/isSecondarySourceFilter
func (c MPSCNNArithmeticGradient) IsSecondarySourceFilter() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSecondarySourceFilter"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/maximumValue
func (c MPSCNNArithmeticGradient) MaximumValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("maximumValue"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetMaximumValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/minimumValue
func (c MPSCNNArithmeticGradient) MinimumValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("minimumValue"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetMinimumValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/primaryScale
func (c MPSCNNArithmeticGradient) PrimaryScale() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("primaryScale"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetPrimaryScale(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/secondaryScale
func (c MPSCNNArithmeticGradient) SecondaryScale() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("secondaryScale"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetSecondaryScale(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient/secondaryStrideInFeatureChannels
func (c MPSCNNArithmeticGradient) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (c MPSCNNArithmeticGradient) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}
