// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageThresholdBinaryInverse] class.
var (
	_MPSImageThresholdBinaryInverseClass     MPSImageThresholdBinaryInverseClass
	_MPSImageThresholdBinaryInverseClassOnce sync.Once
)

func getMPSImageThresholdBinaryInverseClass() MPSImageThresholdBinaryInverseClass {
	_MPSImageThresholdBinaryInverseClassOnce.Do(func() {
		_MPSImageThresholdBinaryInverseClass = MPSImageThresholdBinaryInverseClass{class: objc.GetClass("MPSImageThresholdBinaryInverse")}
	})
	return _MPSImageThresholdBinaryInverseClass
}

// GetMPSImageThresholdBinaryInverseClass returns the class object for MPSImageThresholdBinaryInverse.
func GetMPSImageThresholdBinaryInverseClass() MPSImageThresholdBinaryInverseClass {
	return getMPSImageThresholdBinaryInverseClass()
}

type MPSImageThresholdBinaryInverseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageThresholdBinaryInverseClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageThresholdBinaryInverseClass) Alloc() MPSImageThresholdBinaryInverse {
	rv := objc.Send[MPSImageThresholdBinaryInverse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns 0 for each pixel with a value greater than a
// specified threshold or a specified value otherwise.
//
// # Overview
//
// An [MPSImageThresholdBinaryInverse] function converts a single channel
// image to a binary image. If the input image is not a single channel image,
// the function first converts the input image into a single channel luminance
// image using the linear gray color transform, and then it applies the
// threshold. The following listing shows the threshold binary inverse
// function.
//
// Listing 1. Threshold binary inverse function
//
// # Methods
//
//   - [MPSImageThresholdBinaryInverse.InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [MPSImageThresholdBinaryInverse.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [MPSImageThresholdBinaryInverse.MaximumValue]: The maximum value used to initialize the threshold filter.
//   - [MPSImageThresholdBinaryInverse.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse
type MPSImageThresholdBinaryInverse struct {
	MPSUnaryImageKernel
}

// MPSImageThresholdBinaryInverseFromID constructs a [MPSImageThresholdBinaryInverse] from an objc.ID.
//
// A filter that returns 0 for each pixel with a value greater than a
// specified threshold or a specified value otherwise.
func MPSImageThresholdBinaryInverseFromID(id objc.ID) MPSImageThresholdBinaryInverse {
	return MPSImageThresholdBinaryInverse{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageThresholdBinaryInverse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageThresholdBinaryInverse] class.
//
// # Methods
//
//   - [IMPSImageThresholdBinaryInverse.InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [IMPSImageThresholdBinaryInverse.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [IMPSImageThresholdBinaryInverse.MaximumValue]: The maximum value used to initialize the threshold filter.
//   - [IMPSImageThresholdBinaryInverse.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse
type IMPSImageThresholdBinaryInverse interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel.
	InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinaryInverse

	// Topic: Properties

	// The threshold value used to initialize the threshold filter.
	ThresholdValue() float32
	// The maximum value used to initialize the threshold filter.
	MaximumValue() float32
	// The color transform used to initialize the threshold filter.
	Transform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageThresholdBinaryInverse) Init() MPSImageThresholdBinaryInverse {
	rv := objc.Send[MPSImageThresholdBinaryInverse](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageThresholdBinaryInverse) Autorelease() MPSImageThresholdBinaryInverse {
	rv := objc.Send[MPSImageThresholdBinaryInverse](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageThresholdBinaryInverse creates a new MPSImageThresholdBinaryInverse instance.
func NewMPSImageThresholdBinaryInverse() MPSImageThresholdBinaryInverse {
	class := getMPSImageThresholdBinaryInverseClass()
	rv := objc.Send[MPSImageThresholdBinaryInverse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageThresholdBinaryInverseWithCoder(aDecoder foundation.INSCoder) MPSImageThresholdBinaryInverse {
	instance := getMPSImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageThresholdBinaryInverseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/init(coder:device:)
func NewImageThresholdBinaryInverseWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageThresholdBinaryInverse {
	instance := getMPSImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageThresholdBinaryInverseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageThresholdBinaryInverseWithDevice(device metal.MTLDevice) MPSImageThresholdBinaryInverse {
	instance := getMPSImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageThresholdBinaryInverseFromID(rv)
}

// Initializes the kernel.
//
// device: The Metal device the filter will run on.
//
// thresholdValue: The threshold value to use.
//
// maximumValue: The maximum value to use.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// defaults to the BT.601/JPEG standard: `{0.299f, 0.587f, 0.114f}`
//
// # Return Value
//
// An initialized kernel object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/init(device:thresholdValue:maximumValue:linearGrayColorTransform:)
func NewImageThresholdBinaryInverseWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinaryInverse {
	instance := getMPSImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	return MPSImageThresholdBinaryInverseFromID(rv)
}

// Initializes the kernel.
//
// device: The Metal device the filter will run on.
//
// thresholdValue: The threshold value to use.
//
// maximumValue: The maximum value to use.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// defaults to the BT.601/JPEG standard: `{0.299f, 0.587f, 0.114f}`
//
// # Return Value
//
// An initialized kernel object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/init(device:thresholdValue:maximumValue:linearGrayColorTransform:)
func (i MPSImageThresholdBinaryInverse) InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinaryInverse {
	rv := objc.Send[MPSImageThresholdBinaryInverse](i.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	return rv
}

// The threshold value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/thresholdValue
func (i MPSImageThresholdBinaryInverse) ThresholdValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("thresholdValue"))
	return rv
}

// The maximum value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/maximumValue
func (i MPSImageThresholdBinaryInverse) MaximumValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("maximumValue"))
	return rv
}

// The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/transform
func (i MPSImageThresholdBinaryInverse) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("transform"))
	return rv
}
