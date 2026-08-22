// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageBilinearScale] class.
var (
	_MPSImageBilinearScaleClass     MPSImageBilinearScaleClass
	_MPSImageBilinearScaleClassOnce sync.Once
)

func getMPSImageBilinearScaleClass() MPSImageBilinearScaleClass {
	_MPSImageBilinearScaleClassOnce.Do(func() {
		_MPSImageBilinearScaleClass = MPSImageBilinearScaleClass{class: objc.GetClass("MPSImageBilinearScale")}
	})
	return _MPSImageBilinearScaleClass
}

// GetMPSImageBilinearScaleClass returns the class object for MPSImageBilinearScale.
func GetMPSImageBilinearScaleClass() MPSImageBilinearScaleClass {
	return getMPSImageBilinearScaleClass()
}

type MPSImageBilinearScaleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageBilinearScaleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageBilinearScaleClass) Alloc() MPSImageBilinearScale {
	rv := objc.Send[MPSImageBilinearScale](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that resizes and changes the aspect ratio of an image using
// Bilinear resampling.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBilinearScale
type MPSImageBilinearScale struct {
	MPSImageScale
}

// MPSImageBilinearScaleFromID constructs a [MPSImageBilinearScale] from an objc.ID.
//
// A filter that resizes and changes the aspect ratio of an image using
// Bilinear resampling.
func MPSImageBilinearScaleFromID(id objc.ID) MPSImageBilinearScale {
	return MPSImageBilinearScale{MPSImageScale: MPSImageScaleFromID(id)}
}

// NOTE: MPSImageBilinearScale adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageBilinearScale] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBilinearScale
type IMPSImageBilinearScale interface {
	IMPSImageScale
}

// Init initializes the instance.
func (i MPSImageBilinearScale) Init() MPSImageBilinearScale {
	rv := objc.Send[MPSImageBilinearScale](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageBilinearScale) Autorelease() MPSImageBilinearScale {
	rv := objc.Send[MPSImageBilinearScale](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageBilinearScale creates a new MPSImageBilinearScale instance.
func NewMPSImageBilinearScale() MPSImageBilinearScale {
	class := getMPSImageBilinearScaleClass()
	rv := objc.Send[MPSImageBilinearScale](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageBilinearScaleWithCoder(aDecoder foundation.INSCoder) MPSImageBilinearScale {
	instance := getMPSImageBilinearScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageBilinearScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBilinearScale/init(coder:device:)
func NewImageBilinearScaleWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageBilinearScale {
	instance := getMPSImageBilinearScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageBilinearScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBilinearScale/init(device:)
func NewImageBilinearScaleWithDevice(device metal.MTLDevice) MPSImageBilinearScale {
	instance := getMPSImageBilinearScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageBilinearScaleFromID(rv)
}
