// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageThresholdToZeroInverse] class.
var (
	_MPSImageThresholdToZeroInverseClass     MPSImageThresholdToZeroInverseClass
	_MPSImageThresholdToZeroInverseClassOnce sync.Once
)

func getMPSImageThresholdToZeroInverseClass() MPSImageThresholdToZeroInverseClass {
	_MPSImageThresholdToZeroInverseClassOnce.Do(func() {
		_MPSImageThresholdToZeroInverseClass = MPSImageThresholdToZeroInverseClass{class: objc.GetClass("MPSImageThresholdToZeroInverse")}
	})
	return _MPSImageThresholdToZeroInverseClass
}

// GetMPSImageThresholdToZeroInverseClass returns the class object for MPSImageThresholdToZeroInverse.
func GetMPSImageThresholdToZeroInverseClass() MPSImageThresholdToZeroInverseClass {
	return getMPSImageThresholdToZeroInverseClass()
}

type MPSImageThresholdToZeroInverseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageThresholdToZeroInverseClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageThresholdToZeroInverseClass) Alloc() MPSImageThresholdToZeroInverse {
	rv := objc.Send[MPSImageThresholdToZeroInverse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns 0 for each pixel with a value greater than a
// specified threshold or the original value otherwise.
//
// # Overview
//
// An [MPSImageThresholdToZeroInverse] filter converts a single channel image
// to a binary image. If the input image is not a single channel image, the
// function first converts the input image into a single channel luminance
// image using the linear gray color transform, and then it applies the
// threshold.
//
// The following listing shows the threshold to zero inverse function.
//
// Listing 1. Threshold to zero inverse function
//
// # New Methods
//
//   - [MPSImageThresholdToZeroInverse.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [MPSImageThresholdToZeroInverse.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [MPSImageThresholdToZeroInverse.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse
type MPSImageThresholdToZeroInverse struct {
	MPSUnaryImageKernel
}

// MPSImageThresholdToZeroInverseFromID constructs a [MPSImageThresholdToZeroInverse] from an objc.ID.
//
// A filter that returns 0 for each pixel with a value greater than a
// specified threshold or the original value otherwise.
func MPSImageThresholdToZeroInverseFromID(id objc.ID) MPSImageThresholdToZeroInverse {
	return MPSImageThresholdToZeroInverse{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageThresholdToZeroInverse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageThresholdToZeroInverse] class.
//
// # New Methods
//
//   - [IMPSImageThresholdToZeroInverse.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [IMPSImageThresholdToZeroInverse.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [IMPSImageThresholdToZeroInverse.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse
type IMPSImageThresholdToZeroInverse interface {
	IMPSUnaryImageKernel

	// Topic: New Methods

	// Initializes the kernel.
	InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZeroInverse

	// Topic: Properties

	// The threshold value used to initialize the threshold filter.
	ThresholdValue() float32
	// The color transform used to initialize the threshold filter.
	Transform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageThresholdToZeroInverse) Init() MPSImageThresholdToZeroInverse {
	rv := objc.Send[MPSImageThresholdToZeroInverse](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageThresholdToZeroInverse) Autorelease() MPSImageThresholdToZeroInverse {
	rv := objc.Send[MPSImageThresholdToZeroInverse](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageThresholdToZeroInverse creates a new MPSImageThresholdToZeroInverse instance.
func NewMPSImageThresholdToZeroInverse() MPSImageThresholdToZeroInverse {
	class := getMPSImageThresholdToZeroInverseClass()
	rv := objc.Send[MPSImageThresholdToZeroInverse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageThresholdToZeroInverseWithCoder(aDecoder foundation.INSCoder) MPSImageThresholdToZeroInverse {
	instance := getMPSImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageThresholdToZeroInverseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse/init(coder:device:)
func NewImageThresholdToZeroInverseWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageThresholdToZeroInverse {
	instance := getMPSImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageThresholdToZeroInverseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageThresholdToZeroInverseWithDevice(device metal.MTLDevice) MPSImageThresholdToZeroInverse {
	instance := getMPSImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageThresholdToZeroInverseFromID(rv)
}

// Initializes the kernel.
//
// device: The Metal device the filter will run on.
//
// thresholdValue: The threshold value to use.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// defaults to the BT.601/JPEG standard: `{0.299f, 0.587f, 0.114f}`
//
// # Return Value
//
// An initialized kernel object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse/init(device:thresholdValue:linearGrayColorTransform:)
func NewImageThresholdToZeroInverseWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZeroInverse {
	instance := getMPSImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return MPSImageThresholdToZeroInverseFromID(rv)
}

// Initializes the kernel.
//
// device: The Metal device the filter will run on.
//
// thresholdValue: The threshold value to use.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// defaults to the BT.601/JPEG standard: `{0.299f, 0.587f, 0.114f}`
//
// # Return Value
//
// An initialized kernel object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse/init(device:thresholdValue:linearGrayColorTransform:)
func (i MPSImageThresholdToZeroInverse) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZeroInverse {
	rv := objc.Send[MPSImageThresholdToZeroInverse](i.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return rv
}

// The threshold value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse/thresholdValue
func (i MPSImageThresholdToZeroInverse) ThresholdValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("thresholdValue"))
	return rv
}

// The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse/transform
func (i MPSImageThresholdToZeroInverse) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("transform"))
	return rv
}
