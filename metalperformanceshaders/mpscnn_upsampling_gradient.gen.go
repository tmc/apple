// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingGradient] class.
var (
	_MPSCNNUpsamplingGradientClass     MPSCNNUpsamplingGradientClass
	_MPSCNNUpsamplingGradientClassOnce sync.Once
)

func getMPSCNNUpsamplingGradientClass() MPSCNNUpsamplingGradientClass {
	_MPSCNNUpsamplingGradientClassOnce.Do(func() {
		_MPSCNNUpsamplingGradientClass = MPSCNNUpsamplingGradientClass{class: objc.GetClass("MPSCNNUpsamplingGradient")}
	})
	return _MPSCNNUpsamplingGradientClass
}

// GetMPSCNNUpsamplingGradientClass returns the class object for MPSCNNUpsamplingGradient.
func GetMPSCNNUpsamplingGradientClass() MPSCNNUpsamplingGradientClass {
	return getMPSCNNUpsamplingGradientClass()
}

type MPSCNNUpsamplingGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingGradientClass) Alloc() MPSCNNUpsamplingGradient {
	rv := objc.Send[MPSCNNUpsamplingGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient filter that upsamples an existing Metal Performance Shaders
// image.
//
// # Instance Properties
//
//   - [MPSCNNUpsamplingGradient.ScaleFactorX]
//   - [MPSCNNUpsamplingGradient.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient
type MPSCNNUpsamplingGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNUpsamplingGradientFromID constructs a [MPSCNNUpsamplingGradient] from an objc.ID.
//
// A gradient filter that upsamples an existing Metal Performance Shaders
// image.
func MPSCNNUpsamplingGradientFromID(id objc.ID) MPSCNNUpsamplingGradient {
	return MPSCNNUpsamplingGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNUpsamplingGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingGradient] class.
//
// # Instance Properties
//
//   - [IMPSCNNUpsamplingGradient.ScaleFactorX]
//   - [IMPSCNNUpsamplingGradient.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient
type IMPSCNNUpsamplingGradient interface {
	IMPSCNNGradientKernel

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
}

// Init initializes the instance.
func (c MPSCNNUpsamplingGradient) Init() MPSCNNUpsamplingGradient {
	rv := objc.Send[MPSCNNUpsamplingGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingGradient) Autorelease() MPSCNNUpsamplingGradient {
	rv := objc.Send[MPSCNNUpsamplingGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingGradient creates a new MPSCNNUpsamplingGradient instance.
func NewMPSCNNUpsamplingGradient() MPSCNNUpsamplingGradient {
	class := getMPSCNNUpsamplingGradientClass()
	rv := objc.Send[MPSCNNUpsamplingGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsamplingGradient {
	instance := getMPSCNNUpsamplingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNUpsamplingGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsamplingGradient {
	instance := getMPSCNNUpsamplingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNUpsamplingGradientWithDevice(device metal.MTLDevice) MPSCNNUpsamplingGradient {
	instance := getMPSCNNUpsamplingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient/scaleFactorX
func (c MPSCNNUpsamplingGradient) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient/scaleFactorY
func (c MPSCNNUpsamplingGradient) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}
