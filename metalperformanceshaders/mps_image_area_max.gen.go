// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageAreaMax] class.
var (
	_MPSImageAreaMaxClass     MPSImageAreaMaxClass
	_MPSImageAreaMaxClassOnce sync.Once
)

func getMPSImageAreaMaxClass() MPSImageAreaMaxClass {
	_MPSImageAreaMaxClassOnce.Do(func() {
		_MPSImageAreaMaxClass = MPSImageAreaMaxClass{class: objc.GetClass("MPSImageAreaMax")}
	})
	return _MPSImageAreaMaxClass
}

// GetMPSImageAreaMaxClass returns the class object for MPSImageAreaMax.
func GetMPSImageAreaMaxClass() MPSImageAreaMaxClass {
	return getMPSImageAreaMaxClass()
}

type MPSImageAreaMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageAreaMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageAreaMaxClass) Alloc() MPSImageAreaMax {
	rv := objc.Send[MPSImageAreaMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that finds the maximum pixel value in a rectangular region
// centered around each pixel in the source image.
//
// # Overview
//
// If there are multiple channels in the source image, each channel is
// processed independently. The [MPSUnaryImageKernel.EdgeMode] property value
// is assumed to always be [MPSImageEdgeModeClamp] for this filter.
//
// # Methods
//
//   - [MPSImageAreaMax.InitWithDeviceKernelWidthKernelHeight]: Initializes the kernel with a specified width and height.
//
// # Properties
//
//   - [MPSImageAreaMax.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [MPSImageAreaMax.KernelWidth]: The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax
type MPSImageAreaMax struct {
	MPSUnaryImageKernel
}

// MPSImageAreaMaxFromID constructs a [MPSImageAreaMax] from an objc.ID.
//
// A filter that finds the maximum pixel value in a rectangular region
// centered around each pixel in the source image.
func MPSImageAreaMaxFromID(id objc.ID) MPSImageAreaMax {
	return MPSImageAreaMax{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageAreaMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageAreaMax] class.
//
// # Methods
//
//   - [IMPSImageAreaMax.InitWithDeviceKernelWidthKernelHeight]: Initializes the kernel with a specified width and height.
//
// # Properties
//
//   - [IMPSImageAreaMax.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [IMPSImageAreaMax.KernelWidth]: The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax
type IMPSImageAreaMax interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes the kernel with a specified width and height.
	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageAreaMax

	// Topic: Properties

	// The height of the filter window. Must be an odd number.
	KernelHeight() uint
	// The width of the filter window. Must be an odd number.
	KernelWidth() uint
}

// Init initializes the instance.
func (i MPSImageAreaMax) Init() MPSImageAreaMax {
	rv := objc.Send[MPSImageAreaMax](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageAreaMax) Autorelease() MPSImageAreaMax {
	rv := objc.Send[MPSImageAreaMax](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageAreaMax creates a new MPSImageAreaMax instance.
func NewMPSImageAreaMax() MPSImageAreaMax {
	class := getMPSImageAreaMaxClass()
	rv := objc.Send[MPSImageAreaMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageAreaMaxWithCoder(aDecoder foundation.INSCoder) MPSImageAreaMax {
	instance := getMPSImageAreaMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageAreaMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/init(coder:device:)
func NewImageAreaMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageAreaMax {
	instance := getMPSImageAreaMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageAreaMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageAreaMaxWithDevice(device metal.MTLDevice) MPSImageAreaMax {
	instance := getMPSImageAreaMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageAreaMaxFromID(rv)
}

// Initializes the kernel with a specified width and height.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// # Return Value
//
// Returns an initialized kernel object with a specific width and height.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/init(device:kernelWidth:kernelHeight:)
func NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageAreaMax {
	instance := getMPSImageAreaMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSImageAreaMaxFromID(rv)
}

// Initializes the kernel with a specified width and height.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// # Return Value
//
// Returns an initialized kernel object with a specific width and height.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/init(device:kernelWidth:kernelHeight:)
func (i MPSImageAreaMax) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageAreaMax {
	rv := objc.Send[MPSImageAreaMax](i.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/kernelHeight
func (i MPSImageAreaMax) KernelHeight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/kernelWidth
func (i MPSImageAreaMax) KernelWidth() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelWidth"))
	return rv
}
