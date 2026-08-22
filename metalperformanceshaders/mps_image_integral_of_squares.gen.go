// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageIntegralOfSquares] class.
var (
	_MPSImageIntegralOfSquaresClass     MPSImageIntegralOfSquaresClass
	_MPSImageIntegralOfSquaresClassOnce sync.Once
)

func getMPSImageIntegralOfSquaresClass() MPSImageIntegralOfSquaresClass {
	_MPSImageIntegralOfSquaresClassOnce.Do(func() {
		_MPSImageIntegralOfSquaresClass = MPSImageIntegralOfSquaresClass{class: objc.GetClass("MPSImageIntegralOfSquares")}
	})
	return _MPSImageIntegralOfSquaresClass
}

// GetMPSImageIntegralOfSquaresClass returns the class object for MPSImageIntegralOfSquares.
func GetMPSImageIntegralOfSquaresClass() MPSImageIntegralOfSquaresClass {
	return getMPSImageIntegralOfSquaresClass()
}

type MPSImageIntegralOfSquaresClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageIntegralOfSquaresClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageIntegralOfSquaresClass) Alloc() MPSImageIntegralOfSquares {
	rv := objc.Send[MPSImageIntegralOfSquares](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that calculates the sum of squared pixels over a specified region
// in an image.
//
// # Overview
//
// The value at each position is the sum of all squared pixels in a source
// image rectangle, `sumRect.` The following listing shows the pseudocode used
// to calculate `sumRect`.
//
// Listing 1. Pseudocode for sumRect
//
// If the channels in the source image are normalized, half-float or floating
// values, the destination image is recommended to be a 32-bit floating-point
// image. If the channels in the source image are integer values, it is
// recommended that an appropriate 32-bit integer image destination format is
// used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegralOfSquares
type MPSImageIntegralOfSquares struct {
	MPSUnaryImageKernel
}

// MPSImageIntegralOfSquaresFromID constructs a [MPSImageIntegralOfSquares] from an objc.ID.
//
// A filter that calculates the sum of squared pixels over a specified region
// in an image.
func MPSImageIntegralOfSquaresFromID(id objc.ID) MPSImageIntegralOfSquares {
	return MPSImageIntegralOfSquares{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageIntegralOfSquares adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageIntegralOfSquares] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegralOfSquares
type IMPSImageIntegralOfSquares interface {
	IMPSUnaryImageKernel
}

// Init initializes the instance.
func (i MPSImageIntegralOfSquares) Init() MPSImageIntegralOfSquares {
	rv := objc.Send[MPSImageIntegralOfSquares](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageIntegralOfSquares) Autorelease() MPSImageIntegralOfSquares {
	rv := objc.Send[MPSImageIntegralOfSquares](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageIntegralOfSquares creates a new MPSImageIntegralOfSquares instance.
func NewMPSImageIntegralOfSquares() MPSImageIntegralOfSquares {
	class := getMPSImageIntegralOfSquaresClass()
	rv := objc.Send[MPSImageIntegralOfSquares](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageIntegralOfSquaresWithCoder(aDecoder foundation.INSCoder) MPSImageIntegralOfSquares {
	instance := getMPSImageIntegralOfSquaresClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageIntegralOfSquaresFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageIntegralOfSquaresWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageIntegralOfSquares {
	instance := getMPSImageIntegralOfSquaresClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageIntegralOfSquaresFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageIntegralOfSquaresWithDevice(device metal.MTLDevice) MPSImageIntegralOfSquares {
	instance := getMPSImageIntegralOfSquaresClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageIntegralOfSquaresFromID(rv)
}
