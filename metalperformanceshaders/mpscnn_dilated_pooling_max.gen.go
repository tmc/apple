// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDilatedPoolingMax] class.
var (
	_MPSCNNDilatedPoolingMaxClass     MPSCNNDilatedPoolingMaxClass
	_MPSCNNDilatedPoolingMaxClassOnce sync.Once
)

func getMPSCNNDilatedPoolingMaxClass() MPSCNNDilatedPoolingMaxClass {
	_MPSCNNDilatedPoolingMaxClassOnce.Do(func() {
		_MPSCNNDilatedPoolingMaxClass = MPSCNNDilatedPoolingMaxClass{class: objc.GetClass("MPSCNNDilatedPoolingMax")}
	})
	return _MPSCNNDilatedPoolingMaxClass
}

// GetMPSCNNDilatedPoolingMaxClass returns the class object for MPSCNNDilatedPoolingMax.
func GetMPSCNNDilatedPoolingMaxClass() MPSCNNDilatedPoolingMaxClass {
	return getMPSCNNDilatedPoolingMaxClass()
}

type MPSCNNDilatedPoolingMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDilatedPoolingMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDilatedPoolingMaxClass) Alloc() MPSCNNDilatedPoolingMax {
	rv := objc.Send[MPSCNNDilatedPoolingMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A dilated max pooling filter.
//
// # Overview
//
// For each pixel, returns the maximum value of pixels in the `kernelWidth *
// kernelHeight` filter region by step size `dilationFactorX` `*`
// `dilationFactorY`.
//
// # Initializers
//
//   - [MPSCNNDilatedPoolingMax.InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY]: Initializes a dilated max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax
type MPSCNNDilatedPoolingMax struct {
	MPSCNNPooling
}

// MPSCNNDilatedPoolingMaxFromID constructs a [MPSCNNDilatedPoolingMax] from an objc.ID.
//
// A dilated max pooling filter.
func MPSCNNDilatedPoolingMaxFromID(id objc.ID) MPSCNNDilatedPoolingMax {
	return MPSCNNDilatedPoolingMax{MPSCNNPooling: MPSCNNPoolingFromID(id)}
}

// NOTE: MPSCNNDilatedPoolingMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDilatedPoolingMax] class.
//
// # Initializers
//
//   - [IMPSCNNDilatedPoolingMax.InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY]: Initializes a dilated max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax
type IMPSCNNDilatedPoolingMax interface {
	IMPSCNNPooling

	// Topic: Initializers

	// Initializes a dilated max pooling filter.
	InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMax
}

// Init initializes the instance.
func (c MPSCNNDilatedPoolingMax) Init() MPSCNNDilatedPoolingMax {
	rv := objc.Send[MPSCNNDilatedPoolingMax](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDilatedPoolingMax) Autorelease() MPSCNNDilatedPoolingMax {
	rv := objc.Send[MPSCNNDilatedPoolingMax](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDilatedPoolingMax creates a new MPSCNNDilatedPoolingMax instance.
func NewMPSCNNDilatedPoolingMax() MPSCNNDilatedPoolingMax {
	class := getMPSCNNDilatedPoolingMaxClass()
	rv := objc.Send[MPSCNNDilatedPoolingMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNDilatedPoolingMaxWithCoder(aDecoder foundation.INSCoder) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// Initializes a dilated max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax/init(coder:device:)
func NewCNNDilatedPoolingMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNDilatedPoolingMaxWithDevice(device metal.MTLDevice) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:)
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// Initializes a dilated max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax/init(device:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:)
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// strideInPixelsX: The output stride (downsampling factor) in the x dimension.
//
// strideInPixelsY: The output stride (downsampling factor) in the y dimension.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMax {
	instance := getMPSCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNDilatedPoolingMaxFromID(rv)
}

// Initializes a dilated max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax/init(device:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:)
func (c MPSCNNDilatedPoolingMax) InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMax {
	rv := objc.Send[MPSCNNDilatedPoolingMax](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return rv
}
