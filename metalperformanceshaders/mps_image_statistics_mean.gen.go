// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageStatisticsMean] class.
var (
	_MPSImageStatisticsMeanClass     MPSImageStatisticsMeanClass
	_MPSImageStatisticsMeanClassOnce sync.Once
)

func getMPSImageStatisticsMeanClass() MPSImageStatisticsMeanClass {
	_MPSImageStatisticsMeanClassOnce.Do(func() {
		_MPSImageStatisticsMeanClass = MPSImageStatisticsMeanClass{class: objc.GetClass("MPSImageStatisticsMean")}
	})
	return _MPSImageStatisticsMeanClass
}

// GetMPSImageStatisticsMeanClass returns the class object for MPSImageStatisticsMean.
func GetMPSImageStatisticsMeanClass() MPSImageStatisticsMeanClass {
	return getMPSImageStatisticsMeanClass()
}

type MPSImageStatisticsMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageStatisticsMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageStatisticsMeanClass) Alloc() MPSImageStatisticsMean {
	rv := objc.Send[MPSImageStatisticsMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that computes the mean for a given region of an image.
//
// # Instance Properties
//
//   - [MPSImageStatisticsMean.ClipRectSource]
//   - [MPSImageStatisticsMean.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean
type MPSImageStatisticsMean struct {
	MPSUnaryImageKernel
}

// MPSImageStatisticsMeanFromID constructs a [MPSImageStatisticsMean] from an objc.ID.
//
// A kernel that computes the mean for a given region of an image.
func MPSImageStatisticsMeanFromID(id objc.ID) MPSImageStatisticsMean {
	return MPSImageStatisticsMean{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageStatisticsMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageStatisticsMean] class.
//
// # Instance Properties
//
//   - [IMPSImageStatisticsMean.ClipRectSource]
//   - [IMPSImageStatisticsMean.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean
type IMPSImageStatisticsMean interface {
	IMPSUnaryImageKernel

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (i MPSImageStatisticsMean) Init() MPSImageStatisticsMean {
	rv := objc.Send[MPSImageStatisticsMean](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageStatisticsMean) Autorelease() MPSImageStatisticsMean {
	rv := objc.Send[MPSImageStatisticsMean](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageStatisticsMean creates a new MPSImageStatisticsMean instance.
func NewMPSImageStatisticsMean() MPSImageStatisticsMean {
	class := getMPSImageStatisticsMeanClass()
	rv := objc.Send[MPSImageStatisticsMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageStatisticsMeanWithCoder(aDecoder foundation.INSCoder) MPSImageStatisticsMean {
	instance := getMPSImageStatisticsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageStatisticsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean/init(coder:device:)
func NewImageStatisticsMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageStatisticsMean {
	instance := getMPSImageStatisticsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageStatisticsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean/init(device:)
func NewImageStatisticsMeanWithDevice(device metal.MTLDevice) MPSImageStatisticsMean {
	instance := getMPSImageStatisticsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageStatisticsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean/clipRectSource
func (i MPSImageStatisticsMean) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageStatisticsMean) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}
