// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageStatisticsMinAndMax] class.
var (
	_MPSImageStatisticsMinAndMaxClass     MPSImageStatisticsMinAndMaxClass
	_MPSImageStatisticsMinAndMaxClassOnce sync.Once
)

func getMPSImageStatisticsMinAndMaxClass() MPSImageStatisticsMinAndMaxClass {
	_MPSImageStatisticsMinAndMaxClassOnce.Do(func() {
		_MPSImageStatisticsMinAndMaxClass = MPSImageStatisticsMinAndMaxClass{class: objc.GetClass("MPSImageStatisticsMinAndMax")}
	})
	return _MPSImageStatisticsMinAndMaxClass
}

// GetMPSImageStatisticsMinAndMaxClass returns the class object for MPSImageStatisticsMinAndMax.
func GetMPSImageStatisticsMinAndMaxClass() MPSImageStatisticsMinAndMaxClass {
	return getMPSImageStatisticsMinAndMaxClass()
}

type MPSImageStatisticsMinAndMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageStatisticsMinAndMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageStatisticsMinAndMaxClass) Alloc() MPSImageStatisticsMinAndMax {
	rv := objc.Send[MPSImageStatisticsMinAndMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that computes the minimum and maximum pixel values for a given
// region of an image.
//
// # Overview
//
// The minimum and maximum values are written to the destination image at the
// following pixel locations:
//
// - Minimum value is written at pixel location `(0, 0)` - Maximum value is
// written at pixel location `(1, 0)`
//
// # Instance Properties
//
//   - [MPSImageStatisticsMinAndMax.ClipRectSource]
//   - [MPSImageStatisticsMinAndMax.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax
type MPSImageStatisticsMinAndMax struct {
	MPSUnaryImageKernel
}

// MPSImageStatisticsMinAndMaxFromID constructs a [MPSImageStatisticsMinAndMax] from an objc.ID.
//
// A kernel that computes the minimum and maximum pixel values for a given
// region of an image.
func MPSImageStatisticsMinAndMaxFromID(id objc.ID) MPSImageStatisticsMinAndMax {
	return MPSImageStatisticsMinAndMax{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageStatisticsMinAndMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageStatisticsMinAndMax] class.
//
// # Instance Properties
//
//   - [IMPSImageStatisticsMinAndMax.ClipRectSource]
//   - [IMPSImageStatisticsMinAndMax.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax
type IMPSImageStatisticsMinAndMax interface {
	IMPSUnaryImageKernel

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (i MPSImageStatisticsMinAndMax) Init() MPSImageStatisticsMinAndMax {
	rv := objc.Send[MPSImageStatisticsMinAndMax](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageStatisticsMinAndMax) Autorelease() MPSImageStatisticsMinAndMax {
	rv := objc.Send[MPSImageStatisticsMinAndMax](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageStatisticsMinAndMax creates a new MPSImageStatisticsMinAndMax instance.
func NewMPSImageStatisticsMinAndMax() MPSImageStatisticsMinAndMax {
	class := getMPSImageStatisticsMinAndMaxClass()
	rv := objc.Send[MPSImageStatisticsMinAndMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageStatisticsMinAndMaxWithCoder(aDecoder foundation.INSCoder) MPSImageStatisticsMinAndMax {
	instance := getMPSImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageStatisticsMinAndMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/init(coder:device:)
func NewImageStatisticsMinAndMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageStatisticsMinAndMax {
	instance := getMPSImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageStatisticsMinAndMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/init(device:)
func NewImageStatisticsMinAndMaxWithDevice(device metal.MTLDevice) MPSImageStatisticsMinAndMax {
	instance := getMPSImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageStatisticsMinAndMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/clipRectSource
func (i MPSImageStatisticsMinAndMax) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageStatisticsMinAndMax) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}
