// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageLaplacian] class.
var (
	_MPSImageLaplacianClass     MPSImageLaplacianClass
	_MPSImageLaplacianClassOnce sync.Once
)

func getMPSImageLaplacianClass() MPSImageLaplacianClass {
	_MPSImageLaplacianClassOnce.Do(func() {
		_MPSImageLaplacianClass = MPSImageLaplacianClass{class: objc.GetClass("MPSImageLaplacian")}
	})
	return _MPSImageLaplacianClass
}

// GetMPSImageLaplacianClass returns the class object for MPSImageLaplacian.
func GetMPSImageLaplacianClass() MPSImageLaplacianClass {
	return getMPSImageLaplacianClass()
}

type MPSImageLaplacianClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageLaplacianClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageLaplacianClass) Alloc() MPSImageLaplacian {
	rv := objc.Send[MPSImageLaplacian](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An optimized Laplacian filter, provided for ease of use.
//
// # Overview
//
// This filter uses an optimized convolution filter with a 3x3 kernel with the
// following weights:
//
// [media-2556916]
//
// # Properties
//
//   - [MPSImageLaplacian.Bias]: The value added to a convolved pixel before it is converted back to its intended storage format.
//   - [MPSImageLaplacian.SetBias]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian
type MPSImageLaplacian struct {
	MPSUnaryImageKernel
}

// MPSImageLaplacianFromID constructs a [MPSImageLaplacian] from an objc.ID.
//
// An optimized Laplacian filter, provided for ease of use.
func MPSImageLaplacianFromID(id objc.ID) MPSImageLaplacian {
	return MPSImageLaplacian{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageLaplacian adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageLaplacian] class.
//
// # Properties
//
//   - [IMPSImageLaplacian.Bias]: The value added to a convolved pixel before it is converted back to its intended storage format.
//   - [IMPSImageLaplacian.SetBias]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian
type IMPSImageLaplacian interface {
	IMPSUnaryImageKernel

	// Topic: Properties

	// The value added to a convolved pixel before it is converted back to its intended storage format.
	Bias() float32
	SetBias(value float32)
}

// Init initializes the instance.
func (i MPSImageLaplacian) Init() MPSImageLaplacian {
	rv := objc.Send[MPSImageLaplacian](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageLaplacian) Autorelease() MPSImageLaplacian {
	rv := objc.Send[MPSImageLaplacian](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageLaplacian creates a new MPSImageLaplacian instance.
func NewMPSImageLaplacian() MPSImageLaplacian {
	class := getMPSImageLaplacianClass()
	rv := objc.Send[MPSImageLaplacian](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageLaplacianWithCoder(aDecoder foundation.INSCoder) MPSImageLaplacian {
	instance := getMPSImageLaplacianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageLaplacianFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageLaplacianWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageLaplacian {
	instance := getMPSImageLaplacianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageLaplacianFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageLaplacianWithDevice(device metal.MTLDevice) MPSImageLaplacian {
	instance := getMPSImageLaplacianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageLaplacianFromID(rv)
}

// The value added to a convolved pixel before it is converted back to its
// intended storage format.
//
// # Discussion
//
// This value can be used to convert negative values into a representable
// range for a unsigned pixel format. For example, many edge detection filters
// produce results in the range `[-k,k]`. By scaling the filter weights by
// `0.5/k` and adding `0.5`, the results will be in range `[0,1]` suitable for
// use with unsigned normalized pixel formats.
//
// This value can also be used in combination with renormalization of the
// filter weights to do video ranging as part of the convolution effect. It
// can also just be used to increase the brightness of the image.
//
// The default value is `0.0f`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian/bias
func (i MPSImageLaplacian) Bias() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("bias"))
	return rv
}
func (i MPSImageLaplacian) SetBias(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setBias:"), value)
}
