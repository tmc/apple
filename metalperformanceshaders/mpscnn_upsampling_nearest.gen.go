// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingNearest] class.
var (
	_MPSCNNUpsamplingNearestClass     MPSCNNUpsamplingNearestClass
	_MPSCNNUpsamplingNearestClassOnce sync.Once
)

func getMPSCNNUpsamplingNearestClass() MPSCNNUpsamplingNearestClass {
	_MPSCNNUpsamplingNearestClassOnce.Do(func() {
		_MPSCNNUpsamplingNearestClass = MPSCNNUpsamplingNearestClass{class: objc.GetClass("MPSCNNUpsamplingNearest")}
	})
	return _MPSCNNUpsamplingNearestClass
}

// GetMPSCNNUpsamplingNearestClass returns the class object for MPSCNNUpsamplingNearest.
func GetMPSCNNUpsamplingNearestClass() MPSCNNUpsamplingNearestClass {
	return getMPSCNNUpsamplingNearestClass()
}

type MPSCNNUpsamplingNearestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingNearestClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingNearestClass) Alloc() MPSCNNUpsamplingNearest {
	rv := objc.Send[MPSCNNUpsamplingNearest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A nearest spatial upsampling filter.
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
//   - [MPSCNNUpsamplingNearest.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]: Initializes a nearest spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearest
type MPSCNNUpsamplingNearest struct {
	MPSCNNUpsampling
}

// MPSCNNUpsamplingNearestFromID constructs a [MPSCNNUpsamplingNearest] from an objc.ID.
//
// A nearest spatial upsampling filter.
func MPSCNNUpsamplingNearestFromID(id objc.ID) MPSCNNUpsamplingNearest {
	return MPSCNNUpsamplingNearest{MPSCNNUpsampling: MPSCNNUpsamplingFromID(id)}
}

// NOTE: MPSCNNUpsamplingNearest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingNearest] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingNearest.InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY]: Initializes a nearest spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearest
type IMPSCNNUpsamplingNearest interface {
	IMPSCNNUpsampling

	// Topic: Initializers

	// Initializes a nearest spatial upsampling filter.
	InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearest
}

// Init initializes the instance.
func (c MPSCNNUpsamplingNearest) Init() MPSCNNUpsamplingNearest {
	rv := objc.Send[MPSCNNUpsamplingNearest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingNearest) Autorelease() MPSCNNUpsamplingNearest {
	rv := objc.Send[MPSCNNUpsamplingNearest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingNearest creates a new MPSCNNUpsamplingNearest instance.
func NewMPSCNNUpsamplingNearest() MPSCNNUpsamplingNearest {
	class := getMPSCNNUpsamplingNearestClass()
	rv := objc.Send[MPSCNNUpsamplingNearest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingNearestWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsamplingNearest {
	instance := getMPSCNNUpsamplingNearestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingNearestFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNUpsamplingNearestWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsamplingNearest {
	instance := getMPSCNNUpsamplingNearestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingNearestFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNUpsamplingNearestWithDevice(device metal.MTLDevice) MPSCNNUpsamplingNearest {
	instance := getMPSCNNUpsamplingNearestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingNearestFromID(rv)
}

// Initializes a nearest spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearest/init(device:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingNearestWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearest {
	instance := getMPSCNNUpsamplingNearestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingNearestFromID(rv)
}

// Initializes a nearest spatial upsampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearest/init(device:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingNearest) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.MTLDevice, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearest {
	rv := objc.Send[MPSCNNUpsamplingNearest](c.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	return rv
}
