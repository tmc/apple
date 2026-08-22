// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingL2NormGradient] class.
var (
	_MPSCNNPoolingL2NormGradientClass     MPSCNNPoolingL2NormGradientClass
	_MPSCNNPoolingL2NormGradientClassOnce sync.Once
)

func getMPSCNNPoolingL2NormGradientClass() MPSCNNPoolingL2NormGradientClass {
	_MPSCNNPoolingL2NormGradientClassOnce.Do(func() {
		_MPSCNNPoolingL2NormGradientClass = MPSCNNPoolingL2NormGradientClass{class: objc.GetClass("MPSCNNPoolingL2NormGradient")}
	})
	return _MPSCNNPoolingL2NormGradientClass
}

// GetMPSCNNPoolingL2NormGradientClass returns the class object for MPSCNNPoolingL2NormGradient.
func GetMPSCNNPoolingL2NormGradientClass() MPSCNNPoolingL2NormGradientClass {
	return getMPSCNNPoolingL2NormGradientClass()
}

type MPSCNNPoolingL2NormGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingL2NormGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingL2NormGradientClass) Alloc() MPSCNNPoolingL2NormGradient {
	rv := objc.Send[MPSCNNPoolingL2NormGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient L2-norm pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradient
type MPSCNNPoolingL2NormGradient struct {
	MPSCNNPoolingGradient
}

// MPSCNNPoolingL2NormGradientFromID constructs a [MPSCNNPoolingL2NormGradient] from an objc.ID.
//
// A gradient L2-norm pooling filter.
func MPSCNNPoolingL2NormGradientFromID(id objc.ID) MPSCNNPoolingL2NormGradient {
	return MPSCNNPoolingL2NormGradient{MPSCNNPoolingGradient: MPSCNNPoolingGradientFromID(id)}
}

// NOTE: MPSCNNPoolingL2NormGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingL2NormGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradient
type IMPSCNNPoolingL2NormGradient interface {
	IMPSCNNPoolingGradient
}

// Init initializes the instance.
func (c MPSCNNPoolingL2NormGradient) Init() MPSCNNPoolingL2NormGradient {
	rv := objc.Send[MPSCNNPoolingL2NormGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingL2NormGradient) Autorelease() MPSCNNPoolingL2NormGradient {
	rv := objc.Send[MPSCNNPoolingL2NormGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingL2NormGradient creates a new MPSCNNPoolingL2NormGradient instance.
func NewMPSCNNPoolingL2NormGradient() MPSCNNPoolingL2NormGradient {
	class := getMPSCNNPoolingL2NormGradientClass()
	rv := objc.Send[MPSCNNPoolingL2NormGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingL2NormGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingL2NormGradient {
	instance := getMPSCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingL2NormGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradient/init(coder:device:)
func NewCNNPoolingL2NormGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingL2NormGradient {
	instance := getMPSCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingL2NormGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNPoolingL2NormGradientWithDevice(device metal.MTLDevice) MPSCNNPoolingL2NormGradient {
	instance := getMPSCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingL2NormGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNPoolingL2NormGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingL2NormGradient {
	instance := getMPSCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingL2NormGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingL2NormGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingL2NormGradient {
	instance := getMPSCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingL2NormGradientFromID(rv)
}
