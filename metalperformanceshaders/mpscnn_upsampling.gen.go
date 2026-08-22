// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsampling] class.
var (
	_MPSCNNUpsamplingClass     MPSCNNUpsamplingClass
	_MPSCNNUpsamplingClassOnce sync.Once
)

func getMPSCNNUpsamplingClass() MPSCNNUpsamplingClass {
	_MPSCNNUpsamplingClassOnce.Do(func() {
		_MPSCNNUpsamplingClass = MPSCNNUpsamplingClass{class: objc.GetClass("MPSCNNUpsampling")}
	})
	return _MPSCNNUpsamplingClass
}

// GetMPSCNNUpsamplingClass returns the class object for MPSCNNUpsampling.
func GetMPSCNNUpsamplingClass() MPSCNNUpsamplingClass {
	return getMPSCNNUpsamplingClass()
}

type MPSCNNUpsamplingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingClass) Alloc() MPSCNNUpsampling {
	rv := objc.Send[MPSCNNUpsampling](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that resamples an existing MPS image.
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
// Nearest and bilinear variants are supported.
//
// # Instance Properties
//
//   - [MPSCNNUpsampling.ScaleFactorX]
//   - [MPSCNNUpsampling.ScaleFactorY]
//   - [MPSCNNUpsampling.AlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling
type MPSCNNUpsampling struct {
	MPSCNNKernel
}

// MPSCNNUpsamplingFromID constructs a [MPSCNNUpsampling] from an objc.ID.
//
// A filter that resamples an existing MPS image.
func MPSCNNUpsamplingFromID(id objc.ID) MPSCNNUpsampling {
	return MPSCNNUpsampling{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNUpsampling adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsampling] class.
//
// # Instance Properties
//
//   - [IMPSCNNUpsampling.ScaleFactorX]
//   - [IMPSCNNUpsampling.ScaleFactorY]
//   - [IMPSCNNUpsampling.AlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling
type IMPSCNNUpsampling interface {
	IMPSCNNKernel

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
	AlignCorners() bool
}

// Init initializes the instance.
func (c MPSCNNUpsampling) Init() MPSCNNUpsampling {
	rv := objc.Send[MPSCNNUpsampling](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsampling) Autorelease() MPSCNNUpsampling {
	rv := objc.Send[MPSCNNUpsampling](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsampling creates a new MPSCNNUpsampling instance.
func NewMPSCNNUpsampling() MPSCNNUpsampling {
	class := getMPSCNNUpsamplingClass()
	rv := objc.Send[MPSCNNUpsampling](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNUpsamplingWithCoder(aDecoder foundation.INSCoder) MPSCNNUpsampling {
	instance := getMPSCNNUpsamplingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNUpsamplingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNUpsamplingWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNUpsampling {
	instance := getMPSCNNUpsamplingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNUpsamplingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNUpsamplingWithDevice(device metal.MTLDevice) MPSCNNUpsampling {
	instance := getMPSCNNUpsamplingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNUpsamplingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling/scaleFactorX
func (c MPSCNNUpsampling) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling/scaleFactorY
func (c MPSCNNUpsampling) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling/alignCorners
func (c MPSCNNUpsampling) AlignCorners() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("alignCorners"))
	return rv
}
