// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageMedian] class.
var (
	_MPSImageMedianClass     MPSImageMedianClass
	_MPSImageMedianClassOnce sync.Once
)

func getMPSImageMedianClass() MPSImageMedianClass {
	_MPSImageMedianClassOnce.Do(func() {
		_MPSImageMedianClass = MPSImageMedianClass{class: objc.GetClass("MPSImageMedian")}
	})
	return _MPSImageMedianClass
}

// GetMPSImageMedianClass returns the class object for MPSImageMedian.
func GetMPSImageMedianClass() MPSImageMedianClass {
	return getMPSImageMedianClass()
}

type MPSImageMedianClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageMedianClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageMedianClass) Alloc() MPSImageMedian {
	rv := objc.Send[MPSImageMedian](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that applies a median filter in a square region centered around
// each pixel in the source image.
//
// # Overview
//
// An [MPSImageMedian] filter finds the median color value for each channel
// within a `kernelDiameter * kernelDiameter` window surrounding the pixel of
// interest. It is a common means of noise reduction and also as a smoothing
// filter with edge preserving qualities.
//
// # Methods
//
//   - [MPSImageMedian.InitWithDeviceKernelDiameter]: Initializes a filter for a particular kernel size and device.
//
// # Properties
//
//   - [MPSImageMedian.KernelDiameter]: The diameter, in pixels, of the filter window.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian
type MPSImageMedian struct {
	MPSUnaryImageKernel
}

// MPSImageMedianFromID constructs a [MPSImageMedian] from an objc.ID.
//
// A filter that applies a median filter in a square region centered around
// each pixel in the source image.
func MPSImageMedianFromID(id objc.ID) MPSImageMedian {
	return MPSImageMedian{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageMedian adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageMedian] class.
//
// # Methods
//
//   - [IMPSImageMedian.InitWithDeviceKernelDiameter]: Initializes a filter for a particular kernel size and device.
//
// # Properties
//
//   - [IMPSImageMedian.KernelDiameter]: The diameter, in pixels, of the filter window.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian
type IMPSImageMedian interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a filter for a particular kernel size and device.
	InitWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageMedian

	// Topic: Properties

	// The diameter, in pixels, of the filter window.
	KernelDiameter() uint
}

// Init initializes the instance.
func (i MPSImageMedian) Init() MPSImageMedian {
	rv := objc.Send[MPSImageMedian](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageMedian) Autorelease() MPSImageMedian {
	rv := objc.Send[MPSImageMedian](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageMedian creates a new MPSImageMedian instance.
func NewMPSImageMedian() MPSImageMedian {
	class := getMPSImageMedianClass()
	rv := objc.Send[MPSImageMedian](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageMedianWithCoder(aDecoder foundation.INSCoder) MPSImageMedian {
	instance := getMPSImageMedianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageMedianFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/init(coder:device:)
func NewImageMedianWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageMedian {
	instance := getMPSImageMedianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageMedianFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageMedianWithDevice(device metal.MTLDevice) MPSImageMedian {
	instance := getMPSImageMedianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageMedianFromID(rv)
}

// Initializes a filter for a particular kernel size and device.
//
// device: The Metal device the filter will run on.
//
// kernelDiameter: The diameter of the median filter, in pixels. Must be an odd number.
//
// # Return Value
//
// An initialized median filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/init(device:kernelDiameter:)
func NewImageMedianWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageMedian {
	instance := getMPSImageMedianClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	return MPSImageMedianFromID(rv)
}

// Initializes a filter for a particular kernel size and device.
//
// device: The Metal device the filter will run on.
//
// kernelDiameter: The diameter of the median filter, in pixels. Must be an odd number.
//
// # Return Value
//
// An initialized median filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/init(device:kernelDiameter:)
func (i MPSImageMedian) InitWithDeviceKernelDiameter(device metal.MTLDevice, kernelDiameter uint) MPSImageMedian {
	rv := objc.Send[MPSImageMedian](i.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	return rv
}

// Queries the maximum diameter, in pixels, of the filter window supported by
// the median filter.
//
// # Return Value
//
// Returns the maximum diameter, in pixels, of the filter window supported by
// the median filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/maxKernelDiameter()
func (_MPSImageMedianClass MPSImageMedianClass) MaxKernelDiameter() uint {
	rv := objc.Send[uint](objc.ID(_MPSImageMedianClass.class), objc.Sel("maxKernelDiameter"))
	return rv
}

// Queries the minimum diameter, in pixels, of the filter window supported by
// the median filter.
//
// # Return Value
//
// Returns the minimum diameter, in pixels, of the filter window supported by
// the median filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/minKernelDiameter()
func (_MPSImageMedianClass MPSImageMedianClass) MinKernelDiameter() uint {
	rv := objc.Send[uint](objc.ID(_MPSImageMedianClass.class), objc.Sel("minKernelDiameter"))
	return rv
}

// The diameter, in pixels, of the filter window.
//
// # Discussion
//
// The median filter is applied to a `kernelDiameter * kernelDiameter` window
// of pixels centered on the corresponding source pixel for each destination
// pixel. The kernel diameter must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/kernelDiameter
func (i MPSImageMedian) KernelDiameter() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelDiameter"))
	return rv
}
