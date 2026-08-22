// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageDilate] class.
var (
	_MPSImageDilateClass     MPSImageDilateClass
	_MPSImageDilateClassOnce sync.Once
)

func getMPSImageDilateClass() MPSImageDilateClass {
	_MPSImageDilateClassOnce.Do(func() {
		_MPSImageDilateClass = MPSImageDilateClass{class: objc.GetClass("MPSImageDilate")}
	})
	return _MPSImageDilateClass
}

// GetMPSImageDilateClass returns the class object for MPSImageDilate.
func GetMPSImageDilateClass() MPSImageDilateClass {
	return getMPSImageDilateClass()
}

type MPSImageDilateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageDilateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageDilateClass) Alloc() MPSImageDilate {
	rv := objc.Send[MPSImageDilate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that finds the maximum pixel value in a rectangular region by
// applying a dilation function.
//
// # Overview
//
// An [MPSImageDilate] filter behaves like the [MPSImageAreaMax] filter,
// except Metal calculates the intensity at each position relative to a
// different value before determining which is the maximum pixel value,
// allowing for shaped, nonrectangular morphological probes.
//
// The code example below shows pseudocode for the calculation that returns
// each pixel value:
//
// A filter that contains all zeros is identical to an [MPSImageAreaMax]
// filter. Metal handles the center filter element as `0` to avoid causing a
// general darkening of the image, and it handles the
// [MPSUnaryImageKernel.EdgeMode] property as [MPSImageEdgeModeClamp] for this
// filter.
//
// # Methods
//
//   - [MPSImageDilate.InitWithDeviceKernelWidthKernelHeightValues]: Initializes the kernel with a specified width, height, and weight values.
//
// # Properties
//
//   - [MPSImageDilate.KernelHeight]: The height of the filter window. which must be an odd number.
//   - [MPSImageDilate.KernelWidth]: The width of the filter window which must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate
type MPSImageDilate struct {
	MPSUnaryImageKernel
}

// MPSImageDilateFromID constructs a [MPSImageDilate] from an objc.ID.
//
// A filter that finds the maximum pixel value in a rectangular region by
// applying a dilation function.
func MPSImageDilateFromID(id objc.ID) MPSImageDilate {
	return MPSImageDilate{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageDilate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageDilate] class.
//
// # Methods
//
//   - [IMPSImageDilate.InitWithDeviceKernelWidthKernelHeightValues]: Initializes the kernel with a specified width, height, and weight values.
//
// # Properties
//
//   - [IMPSImageDilate.KernelHeight]: The height of the filter window. which must be an odd number.
//   - [IMPSImageDilate.KernelWidth]: The width of the filter window which must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate
type IMPSImageDilate interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel with a specified width, height, and weight values.
	InitWithDeviceKernelWidthKernelHeightValues(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, values *float32) MPSImageDilate

	// Topic: Properties

	// The height of the filter window. which must be an odd number.
	KernelHeight() uint
	// The width of the filter window which must be an odd number.
	KernelWidth() uint
}

// Init initializes the instance.
func (i MPSImageDilate) Init() MPSImageDilate {
	rv := objc.Send[MPSImageDilate](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageDilate) Autorelease() MPSImageDilate {
	rv := objc.Send[MPSImageDilate](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageDilate creates a new MPSImageDilate instance.
func NewMPSImageDilate() MPSImageDilate {
	class := getMPSImageDilateClass()
	rv := objc.Send[MPSImageDilate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageDilateWithCoder(aDecoder foundation.INSCoder) MPSImageDilate {
	instance := getMPSImageDilateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageDilateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(coder:device:)
func NewImageDilateWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageDilate {
	instance := getMPSImageDilateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageDilateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageDilateWithDevice(device metal.MTLDevice) MPSImageDilate {
	instance := getMPSImageDilateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageDilateFromID(rv)
}

// Initializes the kernel with a specified width, height, and weight values.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// values: The set of values to use as the dilate probe. The values are copied into
// the filter. To avoid image lightening or darkening, the center value should
// be `0.0f`.
//
// # Return Value
//
// Returns an initialized kernel object with specific width, height, and
// weight values.
//
// # Discussion
//
// Each dilate shape probe defines a 3D surface of values. These are arranged
// in order left to right, then top to bottom in a 1D array.
// (`values[kernelWidth*y+x] = probe[y][x]`)
//
// Values should be generally be in the range `[0,1]` with the center pixel
// tending towards `0` and edges towards `1`. However, any numerical value is
// allowed. Calculations are subject to the usual floating-point rounding
// error.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(device:kernelWidth:kernelHeight:values:)
func NewImageDilateWithDeviceKernelWidthKernelHeightValues(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, values *float32) MPSImageDilate {
	instance := getMPSImageDilateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), device, kernelWidth, kernelHeight, values)
	return MPSImageDilateFromID(rv)
}

// Initializes the kernel with a specified width, height, and weight values.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// values: The set of values to use as the dilate probe. The values are copied into
// the filter. To avoid image lightening or darkening, the center value should
// be `0.0f`.
//
// # Return Value
//
// Returns an initialized kernel object with specific width, height, and
// weight values.
//
// # Discussion
//
// Each dilate shape probe defines a 3D surface of values. These are arranged
// in order left to right, then top to bottom in a 1D array.
// (`values[kernelWidth*y+x] = probe[y][x]`)
//
// Values should be generally be in the range `[0,1]` with the center pixel
// tending towards `0` and edges towards `1`. However, any numerical value is
// allowed. Calculations are subject to the usual floating-point rounding
// error.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(device:kernelWidth:kernelHeight:values:)
func (i MPSImageDilate) InitWithDeviceKernelWidthKernelHeightValues(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, values *float32) MPSImageDilate {
	rv := objc.Send[MPSImageDilate](i.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), device, kernelWidth, kernelHeight, values)
	return rv
}

// The height of the filter window. which must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/kernelHeight
func (i MPSImageDilate) KernelHeight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window which must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/kernelWidth
func (i MPSImageDilate) KernelWidth() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelWidth"))
	return rv
}
