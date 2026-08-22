// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingNearestGradient] class.
var (
	_MPSCNNUpsamplingNearestGradientClass     MPSCNNUpsamplingNearestGradientClass
	_MPSCNNUpsamplingNearestGradientClassOnce sync.Once
)

func getMPSCNNUpsamplingNearestGradientClass() MPSCNNUpsamplingNearestGradientClass {
	_MPSCNNUpsamplingNearestGradientClassOnce.Do(func() {
		_MPSCNNUpsamplingNearestGradientClass = MPSCNNUpsamplingNearestGradientClass{class: objc.GetClass("MPSCNNUpsamplingNearestGradient")}
	})
	return _MPSCNNUpsamplingNearestGradientClass
}

// GetMPSCNNUpsamplingNearestGradientClass returns the class object for MPSCNNUpsamplingNearestGradient.
func GetMPSCNNUpsamplingNearestGradientClass() MPSCNNUpsamplingNearestGradientClass {
	return getMPSCNNUpsamplingNearestGradientClass()
}

type MPSCNNUpsamplingNearestGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingNearestGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingNearestGradientClass) Alloc() MPSCNNUpsamplingNearestGradient {
	rv := objc.Send[MPSCNNUpsamplingNearestGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient upsampling filter that samples the pixel nearest to the source
// when upsampling to the destination pixel.
//
// # Initializers
//
//   - [MPSCNNUpsamplingNearestGradient.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient
type MPSCNNUpsamplingNearestGradient struct {
	MPSCNNUpsamplingGradient
}

// MPSCNNUpsamplingNearestGradientFromID constructs a [MPSCNNUpsamplingNearestGradient] from an objc.ID.
//
// A gradient upsampling filter that samples the pixel nearest to the source
// when upsampling to the destination pixel.
func MPSCNNUpsamplingNearestGradientFromID(id objc.ID) MPSCNNUpsamplingNearestGradient {
	return MPSCNNUpsamplingNearestGradient{MPSCNNUpsamplingGradient: MPSCNNUpsamplingGradientFromID(id)}
}

// NOTE: MPSCNNUpsamplingNearestGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingNearestGradient] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingNearestGradient.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient
type IMPSCNNUpsamplingNearestGradient interface {
	IMPSCNNUpsamplingGradient

	// Topic: Initializers

	InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestGradient
}

// Init initializes the instance.
func (c MPSCNNUpsamplingNearestGradient) Init() MPSCNNUpsamplingNearestGradient {
	rv := objc.Send[MPSCNNUpsamplingNearestGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingNearestGradient) Autorelease() MPSCNNUpsamplingNearestGradient {
	rv := objc.Send[MPSCNNUpsamplingNearestGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingNearestGradient creates a new MPSCNNUpsamplingNearestGradient instance.
func NewMPSCNNUpsamplingNearestGradient() MPSCNNUpsamplingNearestGradient {
	class := getMPSCNNUpsamplingNearestGradientClass()
	rv := objc.Send[MPSCNNUpsamplingNearestGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingNearestGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsamplingNearestGradient {
	instance := getMPSCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingNearestGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNUpsamplingNearestGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsamplingNearestGradient {
	instance := getMPSCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingNearestGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNUpsamplingNearestGradientWithDevice(device metal.MTLDevice) MPSCNNUpsamplingNearestGradient {
	instance := getMPSCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingNearestGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient/init(device:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingNearestGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestGradient {
	instance := getMPSCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingNearestGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient/init(device:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingNearestGradient) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestGradient {
	rv := objc.Send[MPSCNNUpsamplingNearestGradient](c.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return rv
}
