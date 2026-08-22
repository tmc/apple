// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingL2Norm] class.
var (
	_MPSCNNPoolingL2NormClass     MPSCNNPoolingL2NormClass
	_MPSCNNPoolingL2NormClassOnce sync.Once
)

func getMPSCNNPoolingL2NormClass() MPSCNNPoolingL2NormClass {
	_MPSCNNPoolingL2NormClassOnce.Do(func() {
		_MPSCNNPoolingL2NormClass = MPSCNNPoolingL2NormClass{class: objc.GetClass("MPSCNNPoolingL2Norm")}
	})
	return _MPSCNNPoolingL2NormClass
}

// GetMPSCNNPoolingL2NormClass returns the class object for MPSCNNPoolingL2Norm.
func GetMPSCNNPoolingL2NormClass() MPSCNNPoolingL2NormClass {
	return getMPSCNNPoolingL2NormClass()
}

type MPSCNNPoolingL2NormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingL2NormClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingL2NormClass) Alloc() MPSCNNPoolingL2Norm {
	rv := objc.Send[MPSCNNPoolingL2Norm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An L2-norm pooling filter.
//
// # Overview
//
// For each pixel, returns L2-Norm of pixels in the `kernelWidth *
// kernelHeight` filter region:
//
// [media-2903549]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2Norm
type MPSCNNPoolingL2Norm struct {
	MPSCNNPooling
}

// MPSCNNPoolingL2NormFromID constructs a [MPSCNNPoolingL2Norm] from an objc.ID.
//
// An L2-norm pooling filter.
func MPSCNNPoolingL2NormFromID(id objc.ID) MPSCNNPoolingL2Norm {
	return MPSCNNPoolingL2Norm{MPSCNNPooling: MPSCNNPoolingFromID(id)}
}

// NOTE: MPSCNNPoolingL2Norm adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingL2Norm] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2Norm
type IMPSCNNPoolingL2Norm interface {
	IMPSCNNPooling
}

// Init initializes the instance.
func (c MPSCNNPoolingL2Norm) Init() MPSCNNPoolingL2Norm {
	rv := objc.Send[MPSCNNPoolingL2Norm](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingL2Norm) Autorelease() MPSCNNPoolingL2Norm {
	rv := objc.Send[MPSCNNPoolingL2Norm](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingL2Norm creates a new MPSCNNPoolingL2Norm instance.
func NewMPSCNNPoolingL2Norm() MPSCNNPoolingL2Norm {
	class := getMPSCNNPoolingL2NormClass()
	rv := objc.Send[MPSCNNPoolingL2Norm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingL2NormWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingL2Norm {
	instance := getMPSCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingL2NormFromID(rv)
}

// Initializes an L2-norm pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2Norm/init(coder:device:)
func NewCNNPoolingL2NormWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingL2Norm {
	instance := getMPSCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingL2NormFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNPoolingL2NormWithDevice(device metal.MTLDevice) MPSCNNPoolingL2Norm {
	instance := getMPSCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingL2NormFromID(rv)
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
func NewCNNPoolingL2NormWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingL2Norm {
	instance := getMPSCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingL2NormFromID(rv)
}

// Initializes an L2-norm pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2Norm/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingL2NormWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingL2Norm {
	instance := getMPSCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingL2NormFromID(rv)
}
