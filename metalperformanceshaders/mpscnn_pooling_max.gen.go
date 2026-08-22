// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingMax] class.
var (
	_MPSCNNPoolingMaxClass     MPSCNNPoolingMaxClass
	_MPSCNNPoolingMaxClassOnce sync.Once
)

func getMPSCNNPoolingMaxClass() MPSCNNPoolingMaxClass {
	_MPSCNNPoolingMaxClassOnce.Do(func() {
		_MPSCNNPoolingMaxClass = MPSCNNPoolingMaxClass{class: objc.GetClass("MPSCNNPoolingMax")}
	})
	return _MPSCNNPoolingMaxClass
}

// GetMPSCNNPoolingMaxClass returns the class object for MPSCNNPoolingMax.
func GetMPSCNNPoolingMaxClass() MPSCNNPoolingMaxClass {
	return getMPSCNNPoolingMaxClass()
}

type MPSCNNPoolingMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingMaxClass) Alloc() MPSCNNPoolingMax {
	rv := objc.Send[MPSCNNPoolingMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A max pooling filter.
//
// # Overview
//
// For each pixel in an image, the filter returns the maximum value of the
// pixels in the filter region defined by `kernelWidth` x `kernelHeight`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMax
type MPSCNNPoolingMax struct {
	MPSCNNPooling
}

// MPSCNNPoolingMaxFromID constructs a [MPSCNNPoolingMax] from an objc.ID.
//
// A max pooling filter.
func MPSCNNPoolingMaxFromID(id objc.ID) MPSCNNPoolingMax {
	return MPSCNNPoolingMax{MPSCNNPooling: MPSCNNPoolingFromID(id)}
}

// NOTE: MPSCNNPoolingMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMax
type IMPSCNNPoolingMax interface {
	IMPSCNNPooling
}

// Init initializes the instance.
func (c MPSCNNPoolingMax) Init() MPSCNNPoolingMax {
	rv := objc.Send[MPSCNNPoolingMax](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingMax) Autorelease() MPSCNNPoolingMax {
	rv := objc.Send[MPSCNNPoolingMax](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingMax creates a new MPSCNNPoolingMax instance.
func NewMPSCNNPoolingMax() MPSCNNPoolingMax {
	class := getMPSCNNPoolingMaxClass()
	rv := objc.Send[MPSCNNPoolingMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingMaxWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingMax {
	instance := getMPSCNNPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingMaxFromID(rv)
}

// Initializes a max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMax/init(coder:device:)
func NewCNNPoolingMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingMax {
	instance := getMPSCNNPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNPoolingMaxWithDevice(device metal.MTLDevice) MPSCNNPoolingMax {
	instance := getMPSCNNPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingMaxFromID(rv)
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
func NewCNNPoolingMaxWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingMax {
	instance := getMPSCNNPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingMaxFromID(rv)
}

// Initializes a max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMax/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingMax {
	instance := getMPSCNNPoolingMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingMaxFromID(rv)
}
