// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageThresholdToZero] class.
var (
	_MPSImageThresholdToZeroClass     MPSImageThresholdToZeroClass
	_MPSImageThresholdToZeroClassOnce sync.Once
)

func getMPSImageThresholdToZeroClass() MPSImageThresholdToZeroClass {
	_MPSImageThresholdToZeroClassOnce.Do(func() {
		_MPSImageThresholdToZeroClass = MPSImageThresholdToZeroClass{class: objc.GetClass("MPSImageThresholdToZero")}
	})
	return _MPSImageThresholdToZeroClass
}

// GetMPSImageThresholdToZeroClass returns the class object for MPSImageThresholdToZero.
func GetMPSImageThresholdToZeroClass() MPSImageThresholdToZeroClass {
	return getMPSImageThresholdToZeroClass()
}

type MPSImageThresholdToZeroClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageThresholdToZeroClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageThresholdToZeroClass) Alloc() MPSImageThresholdToZero {
	rv := objc.Send[MPSImageThresholdToZero](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the original value for each pixel with a value
// greater than a specified threshold or 0 otherwise.
//
// # Overview
//
// An [MPSImageThresholdToZero] filter converts a single channel image to a
// binary image. If the input image is not a single channel image, the
// function first converts the input image into a single channel luminance
// image using the linear gray color transform, and then it applies the
// threshold.
//
// The following listing shows the threshold to zero function.
//
// Listing 1. Threshold to zero function
//
// # Methods
//
//   - [MPSImageThresholdToZero.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [MPSImageThresholdToZero.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [MPSImageThresholdToZero.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero
type MPSImageThresholdToZero struct {
	MPSUnaryImageKernel
}

// MPSImageThresholdToZeroFromID constructs a [MPSImageThresholdToZero] from an objc.ID.
//
// A filter that returns the original value for each pixel with a value
// greater than a specified threshold or 0 otherwise.
func MPSImageThresholdToZeroFromID(id objc.ID) MPSImageThresholdToZero {
	return MPSImageThresholdToZero{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageThresholdToZero adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageThresholdToZero] class.
//
// # Methods
//
//   - [IMPSImageThresholdToZero.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [IMPSImageThresholdToZero.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [IMPSImageThresholdToZero.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero
type IMPSImageThresholdToZero interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel.
	InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZero

	// Topic: Properties

	// The threshold value used to initialize the threshold filter.
	ThresholdValue() float32
	// The color transform used to initialize the threshold filter.
	Transform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageThresholdToZero) Init() MPSImageThresholdToZero {
	rv := objc.Send[MPSImageThresholdToZero](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageThresholdToZero) Autorelease() MPSImageThresholdToZero {
	rv := objc.Send[MPSImageThresholdToZero](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageThresholdToZero creates a new MPSImageThresholdToZero instance.
func NewMPSImageThresholdToZero() MPSImageThresholdToZero {
	class := getMPSImageThresholdToZeroClass()
	rv := objc.Send[MPSImageThresholdToZero](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageThresholdToZeroWithCoder(aDecoder foundation.INSCoder) MPSImageThresholdToZero {
	instance := getMPSImageThresholdToZeroClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageThresholdToZeroFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/init(coder:device:)
func NewImageThresholdToZeroWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageThresholdToZero {
	instance := getMPSImageThresholdToZeroClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageThresholdToZeroFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageThresholdToZeroWithDevice(device metal.MTLDevice) MPSImageThresholdToZero {
	instance := getMPSImageThresholdToZeroClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageThresholdToZeroFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/init(device:thresholdValue:linearGrayColorTransform:)
func NewImageThresholdToZeroWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZero {
	instance := getMPSImageThresholdToZeroClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return MPSImageThresholdToZeroFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/init(device:thresholdValue:linearGrayColorTransform:)
func (i MPSImageThresholdToZero) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdToZero {
	rv := objc.Send[MPSImageThresholdToZero](i.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return rv
}

// The threshold value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/thresholdValue
func (i MPSImageThresholdToZero) ThresholdValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("thresholdValue"))
	return rv
}

// The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/transform
func (i MPSImageThresholdToZero) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("transform"))
	return rv
}
