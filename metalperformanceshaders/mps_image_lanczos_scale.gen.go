// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageLanczosScale] class.
var (
	_MPSImageLanczosScaleClass     MPSImageLanczosScaleClass
	_MPSImageLanczosScaleClassOnce sync.Once
)

func getMPSImageLanczosScaleClass() MPSImageLanczosScaleClass {
	_MPSImageLanczosScaleClassOnce.Do(func() {
		_MPSImageLanczosScaleClass = MPSImageLanczosScaleClass{class: objc.GetClass("MPSImageLanczosScale")}
	})
	return _MPSImageLanczosScaleClass
}

// GetMPSImageLanczosScaleClass returns the class object for MPSImageLanczosScale.
func GetMPSImageLanczosScaleClass() MPSImageLanczosScaleClass {
	return getMPSImageLanczosScaleClass()
}

type MPSImageLanczosScaleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageLanczosScaleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageLanczosScaleClass) Alloc() MPSImageLanczosScale {
	rv := objc.Send[MPSImageLanczosScale](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that resizes and changes the aspect ratio of an image using
// Lanczos resampling.
//
// # Overview
//
// You can use this filter to enlarge or reduce the size of an image, or to
// change the aspect ratio of an image. The filter uses a Lanczos resampling
// algorithm, that typically produces better quality for photographs, but is
// slower than linear sampling that uses GPU texture units. Lanczos
// downsampling does not require a low pass filter to be applied before it is
// used. Because the resampling function has negative lobes, Lanczos can
// result in ringing artifacts near sharp edges, making it less suitable for
// vector art.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLanczosScale
type MPSImageLanczosScale struct {
	MPSImageScale
}

// MPSImageLanczosScaleFromID constructs a [MPSImageLanczosScale] from an objc.ID.
//
// A filter that resizes and changes the aspect ratio of an image using
// Lanczos resampling.
func MPSImageLanczosScaleFromID(id objc.ID) MPSImageLanczosScale {
	return MPSImageLanczosScale{MPSImageScale: MPSImageScaleFromID(id)}
}

// NOTE: MPSImageLanczosScale adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageLanczosScale] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLanczosScale
type IMPSImageLanczosScale interface {
	IMPSImageScale
}

// Init initializes the instance.
func (i MPSImageLanczosScale) Init() MPSImageLanczosScale {
	rv := objc.Send[MPSImageLanczosScale](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageLanczosScale) Autorelease() MPSImageLanczosScale {
	rv := objc.Send[MPSImageLanczosScale](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageLanczosScale creates a new MPSImageLanczosScale instance.
func NewMPSImageLanczosScale() MPSImageLanczosScale {
	class := getMPSImageLanczosScaleClass()
	rv := objc.Send[MPSImageLanczosScale](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageLanczosScaleWithCoder(aDecoder foundation.INSCoder) MPSImageLanczosScale {
	instance := getMPSImageLanczosScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageLanczosScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLanczosScale/init(coder:device:)
func NewImageLanczosScaleWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageLanczosScale {
	instance := getMPSImageLanczosScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageLanczosScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLanczosScale/init(device:)
func NewImageLanczosScaleWithDevice(device metal.MTLDevice) MPSImageLanczosScale {
	instance := getMPSImageLanczosScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageLanczosScaleFromID(rv)
}
