// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageThresholdBinary] class.
var (
	_MPSImageThresholdBinaryClass     MPSImageThresholdBinaryClass
	_MPSImageThresholdBinaryClassOnce sync.Once
)

func getMPSImageThresholdBinaryClass() MPSImageThresholdBinaryClass {
	_MPSImageThresholdBinaryClassOnce.Do(func() {
		_MPSImageThresholdBinaryClass = MPSImageThresholdBinaryClass{class: objc.GetClass("MPSImageThresholdBinary")}
	})
	return _MPSImageThresholdBinaryClass
}

// GetMPSImageThresholdBinaryClass returns the class object for MPSImageThresholdBinary.
func GetMPSImageThresholdBinaryClass() MPSImageThresholdBinaryClass {
	return getMPSImageThresholdBinaryClass()
}

type MPSImageThresholdBinaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageThresholdBinaryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageThresholdBinaryClass) Alloc() MPSImageThresholdBinary {
	rv := objc.Send[MPSImageThresholdBinary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns a specified value for each pixel with a value greater
// than a specified threshold or 0 otherwise.
//
// # Overview
//
// An [MPSImageThresholdBinary] filter converts a single channel image to a
// binary image. If the input image is not a single channel image, the
// function first converts the input image into a single channel luminance
// image using the linear gray color transform, and then it applies the
// threshold.
//
// The following listing shows the threshold binary function.
//
// Listing 1. Threshold binary function
//
// # Methods
//
//   - [MPSImageThresholdBinary.InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [MPSImageThresholdBinary.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [MPSImageThresholdBinary.MaximumValue]: The maximum value used to initialize the threshold filter.
//   - [MPSImageThresholdBinary.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary
type MPSImageThresholdBinary struct {
	MPSUnaryImageKernel
}

// MPSImageThresholdBinaryFromID constructs a [MPSImageThresholdBinary] from an objc.ID.
//
// A filter that returns a specified value for each pixel with a value greater
// than a specified threshold or 0 otherwise.
func MPSImageThresholdBinaryFromID(id objc.ID) MPSImageThresholdBinary {
	return MPSImageThresholdBinary{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageThresholdBinary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageThresholdBinary] class.
//
// # Methods
//
//   - [IMPSImageThresholdBinary.InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform]: Initializes the kernel.
//
// # Properties
//
//   - [IMPSImageThresholdBinary.ThresholdValue]: The threshold value used to initialize the threshold filter.
//   - [IMPSImageThresholdBinary.MaximumValue]: The maximum value used to initialize the threshold filter.
//   - [IMPSImageThresholdBinary.Transform]: The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary
type IMPSImageThresholdBinary interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel.
	InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinary

	// Topic: Properties

	// The threshold value used to initialize the threshold filter.
	ThresholdValue() float32
	// The maximum value used to initialize the threshold filter.
	MaximumValue() float32
	// The color transform used to initialize the threshold filter.
	Transform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageThresholdBinary) Init() MPSImageThresholdBinary {
	rv := objc.Send[MPSImageThresholdBinary](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageThresholdBinary) Autorelease() MPSImageThresholdBinary {
	rv := objc.Send[MPSImageThresholdBinary](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageThresholdBinary creates a new MPSImageThresholdBinary instance.
func NewMPSImageThresholdBinary() MPSImageThresholdBinary {
	class := getMPSImageThresholdBinaryClass()
	rv := objc.Send[MPSImageThresholdBinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageThresholdBinaryWithCoder(aDecoder foundation.INSCoder) MPSImageThresholdBinary {
	instance := getMPSImageThresholdBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageThresholdBinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/init(coder:device:)
func NewImageThresholdBinaryWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageThresholdBinary {
	instance := getMPSImageThresholdBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageThresholdBinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageThresholdBinaryWithDevice(device metal.MTLDevice) MPSImageThresholdBinary {
	instance := getMPSImageThresholdBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageThresholdBinaryFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/init(device:thresholdValue:maximumValue:linearGrayColorTransform:)
func NewImageThresholdBinaryWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinary {
	instance := getMPSImageThresholdBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	return MPSImageThresholdBinaryFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/init(device:thresholdValue:maximumValue:linearGrayColorTransform:)
func (i MPSImageThresholdBinary) InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.MTLDevice, thresholdValue float32, maximumValue float32, transform *float32) MPSImageThresholdBinary {
	rv := objc.Send[MPSImageThresholdBinary](i.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	return rv
}

// The threshold value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/thresholdValue
func (i MPSImageThresholdBinary) ThresholdValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("thresholdValue"))
	return rv
}

// The maximum value used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/maximumValue
func (i MPSImageThresholdBinary) MaximumValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("maximumValue"))
	return rv
}

// The color transform used to initialize the threshold filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary/transform
func (i MPSImageThresholdBinary) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("transform"))
	return rv
}
