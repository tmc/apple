// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageThresholdTruncate] class.
var (
	_MPSImageThresholdTruncateClass     MPSImageThresholdTruncateClass
	_MPSImageThresholdTruncateClassOnce sync.Once
)

func getMPSImageThresholdTruncateClass() MPSImageThresholdTruncateClass {
	_MPSImageThresholdTruncateClassOnce.Do(func() {
		_MPSImageThresholdTruncateClass = MPSImageThresholdTruncateClass{class: objc.GetClass("MPSImageThresholdTruncate")}
	})
	return _MPSImageThresholdTruncateClass
}

// GetMPSImageThresholdTruncateClass returns the class object for MPSImageThresholdTruncate.
func GetMPSImageThresholdTruncateClass() MPSImageThresholdTruncateClass {
	return getMPSImageThresholdTruncateClass()
}

type MPSImageThresholdTruncateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageThresholdTruncateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageThresholdTruncateClass) Alloc() MPSImageThresholdTruncate {
	rv := objc.Send[MPSImageThresholdTruncate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that clamps the return value to an upper specified value.
//
// # Overview
//
// An [MPSImageThresholdTruncate] filter converts a single channel image to a
// binary image. If the input image is not a single channel image, the
// function first converts the input image into a single channel luminance
// image using the linear gray color transform, and then it applies the
// threshold.
//
// The following listing shows the threshold truncate function.
//
// Listing 1. Threshold truncate function
//
// # Methods
//
//   - [MPSImageThresholdTruncate.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [MPSImageThresholdTruncate.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [MPSImageThresholdTruncate.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate
type MPSImageThresholdTruncate struct {
	MPSUnaryImageKernel
}

// MPSImageThresholdTruncateFromID constructs a [MPSImageThresholdTruncate] from an objc.ID.
//
// A filter that clamps the return value to an upper specified value.
func MPSImageThresholdTruncateFromID(id objc.ID) MPSImageThresholdTruncate {
	return MPSImageThresholdTruncate{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageThresholdTruncate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageThresholdTruncate] class.
//
// # Methods
//
//   - [IMPSImageThresholdTruncate.InitWithDeviceThresholdValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [IMPSImageThresholdTruncate.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [IMPSImageThresholdTruncate.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate
type IMPSImageThresholdTruncate interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel.
	InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdTruncate

	// Topic: Properties

	// The threshold value used to initialize the threshold filter.
	ThresholdValue() float32
	// The color transform used to initialize the threshold filter.
	Transform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageThresholdTruncate) Init() MPSImageThresholdTruncate {
	rv := objc.Send[MPSImageThresholdTruncate](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageThresholdTruncate) Autorelease() MPSImageThresholdTruncate {
	rv := objc.Send[MPSImageThresholdTruncate](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageThresholdTruncate creates a new MPSImageThresholdTruncate instance.
func NewMPSImageThresholdTruncate() MPSImageThresholdTruncate {
	class := getMPSImageThresholdTruncateClass()
	rv := objc.Send[MPSImageThresholdTruncate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageThresholdTruncateWithCoder(aDecoder foundation.INSCoder) MPSImageThresholdTruncate {
	instance := getMPSImageThresholdTruncateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageThresholdTruncateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate/init(coder:device:)
func NewImageThresholdTruncateWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageThresholdTruncate {
	instance := getMPSImageThresholdTruncateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageThresholdTruncateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageThresholdTruncateWithDevice(device metal.MTLDevice) MPSImageThresholdTruncate {
	instance := getMPSImageThresholdTruncateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageThresholdTruncateFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate/init(device:thresholdValue:linearGrayColorTransform:)
func NewImageThresholdTruncateWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdTruncate {
	instance := getMPSImageThresholdTruncateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return MPSImageThresholdTruncateFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate/init(device:thresholdValue:linearGrayColorTransform:)
func (i MPSImageThresholdTruncate) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, transform *float32) MPSImageThresholdTruncate {
	rv := objc.Send[MPSImageThresholdTruncate](i.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	return rv
}

// The threshold value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate/thresholdValue
func (i MPSImageThresholdTruncate) ThresholdValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("thresholdValue"))
	return rv
}

// The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate/transform
func (i MPSImageThresholdTruncate) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("transform"))
	return rv
}
