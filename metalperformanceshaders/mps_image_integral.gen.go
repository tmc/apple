// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageIntegral] class.
var (
	_MPSImageIntegralClass     MPSImageIntegralClass
	_MPSImageIntegralClassOnce sync.Once
)

func getMPSImageIntegralClass() MPSImageIntegralClass {
	_MPSImageIntegralClassOnce.Do(func() {
		_MPSImageIntegralClass = MPSImageIntegralClass{class: objc.GetClass("MPSImageIntegral")}
	})
	return _MPSImageIntegralClass
}

// GetMPSImageIntegralClass returns the class object for MPSImageIntegral.
func GetMPSImageIntegralClass() MPSImageIntegralClass {
	return getMPSImageIntegralClass()
}

type MPSImageIntegralClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageIntegralClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageIntegralClass) Alloc() MPSImageIntegral {
	rv := objc.Send[MPSImageIntegral](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that calculates the sum of pixels over a specified region in an
// image.
//
// # Overview
//
// The value at each position is the sum of all pixels in a source image
// rectangle, `sumRect.` The following listing shows the pseudocode used to
// calculate `sumRect`.
//
// Listing 1. Pseudocode for sumRect
//
// If the channels in the source image are normalized, half-float or floating
// values, the destination image is recommended to be a 32-bit floating-point
// image. If the channels in the source image are integer values, it is
// recommended that an appropriate 32-bit integer image destination format is
// used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegral
type MPSImageIntegral struct {
	MPSUnaryImageKernel
}

// MPSImageIntegralFromID constructs a [MPSImageIntegral] from an objc.ID.
//
// A filter that calculates the sum of pixels over a specified region in an
// image.
func MPSImageIntegralFromID(id objc.ID) MPSImageIntegral {
	return MPSImageIntegral{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageIntegral adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageIntegral] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegral
type IMPSImageIntegral interface {
	IMPSUnaryImageKernel
}

// Init initializes the instance.
func (i MPSImageIntegral) Init() MPSImageIntegral {
	rv := objc.Send[MPSImageIntegral](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageIntegral) Autorelease() MPSImageIntegral {
	rv := objc.Send[MPSImageIntegral](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageIntegral creates a new MPSImageIntegral instance.
func NewMPSImageIntegral() MPSImageIntegral {
	class := getMPSImageIntegralClass()
	rv := objc.Send[MPSImageIntegral](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageIntegralWithCoder(aDecoder foundation.INSCoder) MPSImageIntegral {
	instance := getMPSImageIntegralClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageIntegralFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageIntegralWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageIntegral {
	instance := getMPSImageIntegralClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageIntegralFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageIntegralWithDevice(device metal.MTLDevice) MPSImageIntegral {
	instance := getMPSImageIntegralClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageIntegralFromID(rv)
}
