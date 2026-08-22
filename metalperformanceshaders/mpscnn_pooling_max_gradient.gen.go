// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingMaxGradient] class.
var (
	_MPSCNNPoolingMaxGradientClass     MPSCNNPoolingMaxGradientClass
	_MPSCNNPoolingMaxGradientClassOnce sync.Once
)

func getMPSCNNPoolingMaxGradientClass() MPSCNNPoolingMaxGradientClass {
	_MPSCNNPoolingMaxGradientClassOnce.Do(func() {
		_MPSCNNPoolingMaxGradientClass = MPSCNNPoolingMaxGradientClass{class: objc.GetClass("MPSCNNPoolingMaxGradient")}
	})
	return _MPSCNNPoolingMaxGradientClass
}

// GetMPSCNNPoolingMaxGradientClass returns the class object for MPSCNNPoolingMaxGradient.
func GetMPSCNNPoolingMaxGradientClass() MPSCNNPoolingMaxGradientClass {
	return getMPSCNNPoolingMaxGradientClass()
}

type MPSCNNPoolingMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingMaxGradientClass) Alloc() MPSCNNPoolingMaxGradient {
	rv := objc.Send[MPSCNNPoolingMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient
type MPSCNNPoolingMaxGradient struct {
	MPSCNNPoolingGradient
}

// MPSCNNPoolingMaxGradientFromID constructs a [MPSCNNPoolingMaxGradient] from an objc.ID.
//
// A gradient max pooling filter.
func MPSCNNPoolingMaxGradientFromID(id objc.ID) MPSCNNPoolingMaxGradient {
	return MPSCNNPoolingMaxGradient{MPSCNNPoolingGradient: MPSCNNPoolingGradientFromID(id)}
}

// NOTE: MPSCNNPoolingMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingMaxGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient
type IMPSCNNPoolingMaxGradient interface {
	IMPSCNNPoolingGradient
}

// Init initializes the instance.
func (c MPSCNNPoolingMaxGradient) Init() MPSCNNPoolingMaxGradient {
	rv := objc.Send[MPSCNNPoolingMaxGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingMaxGradient) Autorelease() MPSCNNPoolingMaxGradient {
	rv := objc.Send[MPSCNNPoolingMaxGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingMaxGradient creates a new MPSCNNPoolingMaxGradient instance.
func NewMPSCNNPoolingMaxGradient() MPSCNNPoolingMaxGradient {
	class := getMPSCNNPoolingMaxGradientClass()
	rv := objc.Send[MPSCNNPoolingMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingMaxGradient {
	instance := getMPSCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient/init(coder:device:)
func NewCNNPoolingMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingMaxGradient {
	instance := getMPSCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNPoolingMaxGradientWithDevice(device metal.MTLDevice) MPSCNNPoolingMaxGradient {
	instance := getMPSCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingMaxGradient {
	instance := getMPSCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingMaxGradient {
	instance := getMPSCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingMaxGradientFromID(rv)
}
