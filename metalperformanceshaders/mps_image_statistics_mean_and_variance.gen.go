// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageStatisticsMeanAndVariance] class.
var (
	_MPSImageStatisticsMeanAndVarianceClass     MPSImageStatisticsMeanAndVarianceClass
	_MPSImageStatisticsMeanAndVarianceClassOnce sync.Once
)

func getMPSImageStatisticsMeanAndVarianceClass() MPSImageStatisticsMeanAndVarianceClass {
	_MPSImageStatisticsMeanAndVarianceClassOnce.Do(func() {
		_MPSImageStatisticsMeanAndVarianceClass = MPSImageStatisticsMeanAndVarianceClass{class: objc.GetClass("MPSImageStatisticsMeanAndVariance")}
	})
	return _MPSImageStatisticsMeanAndVarianceClass
}

// GetMPSImageStatisticsMeanAndVarianceClass returns the class object for MPSImageStatisticsMeanAndVariance.
func GetMPSImageStatisticsMeanAndVarianceClass() MPSImageStatisticsMeanAndVarianceClass {
	return getMPSImageStatisticsMeanAndVarianceClass()
}

type MPSImageStatisticsMeanAndVarianceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageStatisticsMeanAndVarianceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageStatisticsMeanAndVarianceClass) Alloc() MPSImageStatisticsMeanAndVariance {
	rv := objc.Send[MPSImageStatisticsMeanAndVariance](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that computes the mean and variance for a given region of an
// image.
//
// # Overview
//
// The mean and variance values are written to the destination image at the
// following pixel locations:
//
// - Mean value is written at pixel location `(0, 0)` - Variance value is
// written at pixel location `(1, 0)`
//
// # Instance Properties
//
//   - [MPSImageStatisticsMeanAndVariance.ClipRectSource]
//   - [MPSImageStatisticsMeanAndVariance.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance
type MPSImageStatisticsMeanAndVariance struct {
	MPSUnaryImageKernel
}

// MPSImageStatisticsMeanAndVarianceFromID constructs a [MPSImageStatisticsMeanAndVariance] from an objc.ID.
//
// A kernel that computes the mean and variance for a given region of an
// image.
func MPSImageStatisticsMeanAndVarianceFromID(id objc.ID) MPSImageStatisticsMeanAndVariance {
	return MPSImageStatisticsMeanAndVariance{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageStatisticsMeanAndVariance adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageStatisticsMeanAndVariance] class.
//
// # Instance Properties
//
//   - [IMPSImageStatisticsMeanAndVariance.ClipRectSource]
//   - [IMPSImageStatisticsMeanAndVariance.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance
type IMPSImageStatisticsMeanAndVariance interface {
	IMPSUnaryImageKernel

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (i MPSImageStatisticsMeanAndVariance) Init() MPSImageStatisticsMeanAndVariance {
	rv := objc.Send[MPSImageStatisticsMeanAndVariance](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageStatisticsMeanAndVariance) Autorelease() MPSImageStatisticsMeanAndVariance {
	rv := objc.Send[MPSImageStatisticsMeanAndVariance](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageStatisticsMeanAndVariance creates a new MPSImageStatisticsMeanAndVariance instance.
func NewMPSImageStatisticsMeanAndVariance() MPSImageStatisticsMeanAndVariance {
	class := getMPSImageStatisticsMeanAndVarianceClass()
	rv := objc.Send[MPSImageStatisticsMeanAndVariance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageStatisticsMeanAndVarianceWithCoder(aDecoder foundation.INSCoder) MPSImageStatisticsMeanAndVariance {
	instance := getMPSImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageStatisticsMeanAndVarianceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance/init(coder:device:)
func NewImageStatisticsMeanAndVarianceWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageStatisticsMeanAndVariance {
	instance := getMPSImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageStatisticsMeanAndVarianceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance/init(device:)
func NewImageStatisticsMeanAndVarianceWithDevice(device metal.MTLDevice) MPSImageStatisticsMeanAndVariance {
	instance := getMPSImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageStatisticsMeanAndVarianceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance/clipRectSource
func (i MPSImageStatisticsMeanAndVariance) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageStatisticsMeanAndVariance) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}
