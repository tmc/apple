// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingBilinear] class.
var (
	_MPSCNNUpsamplingBilinearClass     MPSCNNUpsamplingBilinearClass
	_MPSCNNUpsamplingBilinearClassOnce sync.Once
)

func getMPSCNNUpsamplingBilinearClass() MPSCNNUpsamplingBilinearClass {
	_MPSCNNUpsamplingBilinearClassOnce.Do(func() {
		_MPSCNNUpsamplingBilinearClass = MPSCNNUpsamplingBilinearClass{class: objc.GetClass("MPSCNNUpsamplingBilinear")}
	})
	return _MPSCNNUpsamplingBilinearClass
}

// GetMPSCNNUpsamplingBilinearClass returns the class object for MPSCNNUpsamplingBilinear.
func GetMPSCNNUpsamplingBilinearClass() MPSCNNUpsamplingBilinearClass {
	return getMPSCNNUpsamplingBilinearClass()
}

type MPSCNNUpsamplingBilinearClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingBilinearClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingBilinearClass) Alloc() MPSCNNUpsamplingBilinear {
	rv := objc.Send[MPSCNNUpsamplingBilinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A bilinear spatial upsampling filter.
//
// # Overview
//
// This filter can be used to resample an existing [MPSImage] using a
// different sampling frequency for the `x` and `y` dimensions with the
// purpose of enlarging the size of an image.
//
// The number of output feature channels remains the same as the number of
// input feature channels.
//
// The `scaleFactor` must be an integer value `>= 1`. The default value is
// `1`.
//
// # Initializers
//
//   - [MPSCNNUpsamplingBilinear.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]: Initializes a bilinear spatial upsampling filter.
//   - [MPSCNNUpsamplingBilinear.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear
type MPSCNNUpsamplingBilinear struct {
	MPSCNNUpsampling
}

// MPSCNNUpsamplingBilinearFromID constructs a [MPSCNNUpsamplingBilinear] from an objc.ID.
//
// A bilinear spatial upsampling filter.
func MPSCNNUpsamplingBilinearFromID(id objc.ID) MPSCNNUpsamplingBilinear {
	return MPSCNNUpsamplingBilinear{MPSCNNUpsampling: MPSCNNUpsamplingFromID(id)}
}

// NOTE: MPSCNNUpsamplingBilinear adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingBilinear] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingBilinear.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]: Initializes a bilinear spatial upsampling filter.
//   - [IMPSCNNUpsamplingBilinear.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear
type IMPSCNNUpsamplingBilinear interface {
	IMPSCNNUpsampling

	// Topic: Initializers

	// Initializes a bilinear spatial upsampling filter.
	InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinear
	InitWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinear
}

// Init initializes the instance.
func (c MPSCNNUpsamplingBilinear) Init() MPSCNNUpsamplingBilinear {
	rv := objc.Send[MPSCNNUpsamplingBilinear](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingBilinear) Autorelease() MPSCNNUpsamplingBilinear {
	rv := objc.Send[MPSCNNUpsamplingBilinear](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingBilinear creates a new MPSCNNUpsamplingBilinear instance.
func NewMPSCNNUpsamplingBilinear() MPSCNNUpsamplingBilinear {
	class := getMPSCNNUpsamplingBilinearClass()
	rv := objc.Send[MPSCNNUpsamplingBilinear](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingBilinearWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsamplingBilinear {
	instance := getMPSCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNUpsamplingBilinearWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsamplingBilinear {
	instance := getMPSCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNUpsamplingBilinearWithDevice(device metal.MTLDevice) MPSCNNUpsamplingBilinear {
	instance := getMPSCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingBilinearFromID(rv)
}

// Initializes a bilinear spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear/init(device:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinear {
	instance := getMPSCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear/init(device:integerScaleFactorX:integerScaleFactorY:alignCorners:)
func NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinear {
	instance := getMPSCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:alignCorners:"), device, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return MPSCNNUpsamplingBilinearFromID(rv)
}

// Initializes a bilinear spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear/init(device:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingBilinear) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinear {
	rv := objc.Send[MPSCNNUpsamplingBilinear](c.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear/init(device:integerScaleFactorX:integerScaleFactorY:alignCorners:)
func (c MPSCNNUpsamplingBilinear) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinear {
	rv := objc.Send[MPSCNNUpsamplingBilinear](c.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:alignCorners:"), device, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return rv
}
