// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingBilinearGradient] class.
var (
	_MPSCNNUpsamplingBilinearGradientClass     MPSCNNUpsamplingBilinearGradientClass
	_MPSCNNUpsamplingBilinearGradientClassOnce sync.Once
)

func getMPSCNNUpsamplingBilinearGradientClass() MPSCNNUpsamplingBilinearGradientClass {
	_MPSCNNUpsamplingBilinearGradientClassOnce.Do(func() {
		_MPSCNNUpsamplingBilinearGradientClass = MPSCNNUpsamplingBilinearGradientClass{class: objc.GetClass("MPSCNNUpsamplingBilinearGradient")}
	})
	return _MPSCNNUpsamplingBilinearGradientClass
}

// GetMPSCNNUpsamplingBilinearGradientClass returns the class object for MPSCNNUpsamplingBilinearGradient.
func GetMPSCNNUpsamplingBilinearGradientClass() MPSCNNUpsamplingBilinearGradientClass {
	return getMPSCNNUpsamplingBilinearGradientClass()
}

type MPSCNNUpsamplingBilinearGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingBilinearGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingBilinearGradientClass) Alloc() MPSCNNUpsamplingBilinearGradient {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient bilinear spatial upsampling filter.
//
// # Initializers
//
//   - [MPSCNNUpsamplingBilinearGradient.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient
type MPSCNNUpsamplingBilinearGradient struct {
	MPSCNNUpsamplingGradient
}

// MPSCNNUpsamplingBilinearGradientFromID constructs a [MPSCNNUpsamplingBilinearGradient] from an objc.ID.
//
// A gradient bilinear spatial upsampling filter.
func MPSCNNUpsamplingBilinearGradientFromID(id objc.ID) MPSCNNUpsamplingBilinearGradient {
	return MPSCNNUpsamplingBilinearGradient{MPSCNNUpsamplingGradient: MPSCNNUpsamplingGradientFromID(id)}
}

// NOTE: MPSCNNUpsamplingBilinearGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingBilinearGradient] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingBilinearGradient.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient
type IMPSCNNUpsamplingBilinearGradient interface {
	IMPSCNNUpsamplingGradient

	// Topic: Initializers

	InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearGradient
}

// Init initializes the instance.
func (c MPSCNNUpsamplingBilinearGradient) Init() MPSCNNUpsamplingBilinearGradient {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingBilinearGradient) Autorelease() MPSCNNUpsamplingBilinearGradient {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingBilinearGradient creates a new MPSCNNUpsamplingBilinearGradient instance.
func NewMPSCNNUpsamplingBilinearGradient() MPSCNNUpsamplingBilinearGradient {
	class := getMPSCNNUpsamplingBilinearGradientClass()
	rv := objc.Send[MPSCNNUpsamplingBilinearGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingBilinearGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsamplingBilinearGradient {
	instance := getMPSCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingBilinearGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNUpsamplingBilinearGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsamplingBilinearGradient {
	instance := getMPSCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingBilinearGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNUpsamplingBilinearGradientWithDevice(device metal.MTLDevice) MPSCNNUpsamplingBilinearGradient {
	instance := getMPSCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingBilinearGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient/init(device:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingBilinearGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearGradient {
	instance := getMPSCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingBilinearGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient/init(device:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingBilinearGradient) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearGradient {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradient](c.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return rv
}
