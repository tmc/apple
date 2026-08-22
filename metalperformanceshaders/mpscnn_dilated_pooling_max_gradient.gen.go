// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDilatedPoolingMaxGradient] class.
var (
	_MPSCNNDilatedPoolingMaxGradientClass     MPSCNNDilatedPoolingMaxGradientClass
	_MPSCNNDilatedPoolingMaxGradientClassOnce sync.Once
)

func getMPSCNNDilatedPoolingMaxGradientClass() MPSCNNDilatedPoolingMaxGradientClass {
	_MPSCNNDilatedPoolingMaxGradientClassOnce.Do(func() {
		_MPSCNNDilatedPoolingMaxGradientClass = MPSCNNDilatedPoolingMaxGradientClass{class: objc.GetClass("MPSCNNDilatedPoolingMaxGradient")}
	})
	return _MPSCNNDilatedPoolingMaxGradientClass
}

// GetMPSCNNDilatedPoolingMaxGradientClass returns the class object for MPSCNNDilatedPoolingMaxGradient.
func GetMPSCNNDilatedPoolingMaxGradientClass() MPSCNNDilatedPoolingMaxGradientClass {
	return getMPSCNNDilatedPoolingMaxGradientClass()
}

type MPSCNNDilatedPoolingMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDilatedPoolingMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDilatedPoolingMaxGradientClass) Alloc() MPSCNNDilatedPoolingMaxGradient {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient dilated max pooling filter.
//
// # Overview
//
// A gradient max pooling filter but the pixels selected in each
// “application” of the max pooling operation are exactly the same pixels
// that would be selected with dilated convolution
//
// # Initializers
//
//   - [MPSCNNDilatedPoolingMaxGradient.InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient
type MPSCNNDilatedPoolingMaxGradient struct {
	MPSCNNPoolingGradient
}

// MPSCNNDilatedPoolingMaxGradientFromID constructs a [MPSCNNDilatedPoolingMaxGradient] from an objc.ID.
//
// A gradient dilated max pooling filter.
func MPSCNNDilatedPoolingMaxGradientFromID(id objc.ID) MPSCNNDilatedPoolingMaxGradient {
	return MPSCNNDilatedPoolingMaxGradient{MPSCNNPoolingGradient: MPSCNNPoolingGradientFromID(id)}
}

// NOTE: MPSCNNDilatedPoolingMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDilatedPoolingMaxGradient] class.
//
// # Initializers
//
//   - [IMPSCNNDilatedPoolingMaxGradient.InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient
type IMPSCNNDilatedPoolingMaxGradient interface {
	IMPSCNNPoolingGradient

	// Topic: Initializers

	InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMaxGradient
}

// Init initializes the instance.
func (c MPSCNNDilatedPoolingMaxGradient) Init() MPSCNNDilatedPoolingMaxGradient {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDilatedPoolingMaxGradient) Autorelease() MPSCNNDilatedPoolingMaxGradient {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDilatedPoolingMaxGradient creates a new MPSCNNDilatedPoolingMaxGradient instance.
func NewMPSCNNDilatedPoolingMaxGradient() MPSCNNDilatedPoolingMaxGradient {
	class := getMPSCNNDilatedPoolingMaxGradientClass()
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNDilatedPoolingMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient/init(coder:device:)
func NewCNNDilatedPoolingMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNDilatedPoolingMaxGradientWithDevice(device metal.MTLDevice) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient/init(device:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:)
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMaxGradient {
	instance := getMPSCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNDilatedPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient/init(device:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:)
func (c MPSCNNDilatedPoolingMaxGradient) InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNDilatedPoolingMaxGradient {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradient](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return rv
}
