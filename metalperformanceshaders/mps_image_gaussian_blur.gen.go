// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageGaussianBlur] class.
var (
	_MPSImageGaussianBlurClass     MPSImageGaussianBlurClass
	_MPSImageGaussianBlurClassOnce sync.Once
)

func getMPSImageGaussianBlurClass() MPSImageGaussianBlurClass {
	_MPSImageGaussianBlurClassOnce.Do(func() {
		_MPSImageGaussianBlurClass = MPSImageGaussianBlurClass{class: objc.GetClass("MPSImageGaussianBlur")}
	})
	return _MPSImageGaussianBlurClass
}

// GetMPSImageGaussianBlurClass returns the class object for MPSImageGaussianBlur.
func GetMPSImageGaussianBlurClass() MPSImageGaussianBlurClass {
	return getMPSImageGaussianBlurClass()
}

type MPSImageGaussianBlurClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageGaussianBlurClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageGaussianBlurClass) Alloc() MPSImageGaussianBlur {
	rv := objc.Send[MPSImageGaussianBlur](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a Gaussian blur of a given sigma in
// both the x and y directions.
//
// # Overview
//
// # Methods
//
//   - [MPSImageGaussianBlur.InitWithDeviceSigma]: Initializes a Gaussian blur filter.
//
// # Properties
//
//   - [MPSImageGaussianBlur.Sigma]: The sigma value with which the filter was created.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur
type MPSImageGaussianBlur struct {
	MPSUnaryImageKernel
}

// MPSImageGaussianBlurFromID constructs a [MPSImageGaussianBlur] from an objc.ID.
//
// A filter that convolves an image with a Gaussian blur of a given sigma in
// both the x and y directions.
func MPSImageGaussianBlurFromID(id objc.ID) MPSImageGaussianBlur {
	return MPSImageGaussianBlur{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageGaussianBlur adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageGaussianBlur] class.
//
// # Methods
//
//   - [IMPSImageGaussianBlur.InitWithDeviceSigma]: Initializes a Gaussian blur filter.
//
// # Properties
//
//   - [IMPSImageGaussianBlur.Sigma]: The sigma value with which the filter was created.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur
type IMPSImageGaussianBlur interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a Gaussian blur filter.
	InitWithDeviceSigma(device metal.MTLDevice, sigma float32) MPSImageGaussianBlur

	// Topic: Properties

	// The sigma value with which the filter was created.
	Sigma() float32
}

// Init initializes the instance.
func (i MPSImageGaussianBlur) Init() MPSImageGaussianBlur {
	rv := objc.Send[MPSImageGaussianBlur](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageGaussianBlur) Autorelease() MPSImageGaussianBlur {
	rv := objc.Send[MPSImageGaussianBlur](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageGaussianBlur creates a new MPSImageGaussianBlur instance.
func NewMPSImageGaussianBlur() MPSImageGaussianBlur {
	class := getMPSImageGaussianBlurClass()
	rv := objc.Send[MPSImageGaussianBlur](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageGaussianBlurWithCoder(aDecoder foundation.INSCoder) MPSImageGaussianBlur {
	instance := getMPSImageGaussianBlurClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageGaussianBlurFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur/init(coder:device:)
func NewImageGaussianBlurWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageGaussianBlur {
	instance := getMPSImageGaussianBlurClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageGaussianBlurFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageGaussianBlurWithDevice(device metal.MTLDevice) MPSImageGaussianBlur {
	instance := getMPSImageGaussianBlurClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageGaussianBlurFromID(rv)
}

// Initializes a Gaussian blur filter.
//
// device: The Metal device the filter will run on.
//
// sigma: The standard deviation of the gaussian blur filter.
//
// Gaussian weight `w`, centered at `0`, at integer grid `i`, is given as:
//
// `w(i) = 1/sqrt(2*pi*sigma) * exp(-i^2/2*sigma^2)`
//
// If we take cut off at 1% of `w(0)` (max weight) beyond which weights are
// considered `0`, we have `ceil(sqrt(-log(0.01)*2)*sigma) ~ ceil(3.7*sigma)`
// as the rough estimate of the filter width.
//
// # Return Value
//
// An initialized Gaussian blur filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur/init(device:sigma:)
func NewImageGaussianBlurWithDeviceSigma(device metal.MTLDevice, sigma float32) MPSImageGaussianBlur {
	instance := getMPSImageGaussianBlurClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sigma:"), device, sigma)
	return MPSImageGaussianBlurFromID(rv)
}

// Initializes a Gaussian blur filter.
//
// device: The Metal device the filter will run on.
//
// sigma: The standard deviation of the gaussian blur filter.
//
// Gaussian weight `w`, centered at `0`, at integer grid `i`, is given as:
//
// `w(i) = 1/sqrt(2*pi*sigma) * exp(-i^2/2*sigma^2)`
//
// If we take cut off at 1% of `w(0)` (max weight) beyond which weights are
// considered `0`, we have `ceil(sqrt(-log(0.01)*2)*sigma) ~ ceil(3.7*sigma)`
// as the rough estimate of the filter width.
//
// # Return Value
//
// An initialized Gaussian blur filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur/init(device:sigma:)
func (i MPSImageGaussianBlur) InitWithDeviceSigma(device metal.MTLDevice, sigma float32) MPSImageGaussianBlur {
	rv := objc.Send[MPSImageGaussianBlur](i.ID, objc.Sel("initWithDevice:sigma:"), device, sigma)
	return rv
}

// The sigma value with which the filter was created.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur/sigma
func (i MPSImageGaussianBlur) Sigma() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("sigma"))
	return rv
}
