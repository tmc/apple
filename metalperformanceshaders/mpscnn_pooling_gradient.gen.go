// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingGradient] class.
var (
	_MPSCNNPoolingGradientClass     MPSCNNPoolingGradientClass
	_MPSCNNPoolingGradientClassOnce sync.Once
)

func getMPSCNNPoolingGradientClass() MPSCNNPoolingGradientClass {
	_MPSCNNPoolingGradientClassOnce.Do(func() {
		_MPSCNNPoolingGradientClass = MPSCNNPoolingGradientClass{class: objc.GetClass("MPSCNNPoolingGradient")}
	})
	return _MPSCNNPoolingGradientClass
}

// GetMPSCNNPoolingGradientClass returns the class object for MPSCNNPoolingGradient.
func GetMPSCNNPoolingGradientClass() MPSCNNPoolingGradientClass {
	return getMPSCNNPoolingGradientClass()
}

type MPSCNNPoolingGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingGradientClass) Alloc() MPSCNNPoolingGradient {
	rv := objc.Send[MPSCNNPoolingGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient pooling kernel.
//
// # Initializers
//
//   - [MPSCNNPoolingGradient.InitWithDeviceKernelWidthKernelHeight]
//   - [MPSCNNPoolingGradient.InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]
//
// # Instance Properties
//
//   - [MPSCNNPoolingGradient.SourceSize]
//   - [MPSCNNPoolingGradient.SetSourceSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient
type MPSCNNPoolingGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNPoolingGradientFromID constructs a [MPSCNNPoolingGradient] from an objc.ID.
//
// A gradient pooling kernel.
func MPSCNNPoolingGradientFromID(id objc.ID) MPSCNNPoolingGradient {
	return MPSCNNPoolingGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNPoolingGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingGradient] class.
//
// # Initializers
//
//   - [IMPSCNNPoolingGradient.InitWithDeviceKernelWidthKernelHeight]
//   - [IMPSCNNPoolingGradient.InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]
//
// # Instance Properties
//
//   - [IMPSCNNPoolingGradient.SourceSize]
//   - [IMPSCNNPoolingGradient.SetSourceSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient
type IMPSCNNPoolingGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingGradient
	InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingGradient

	// Topic: Instance Properties

	SourceSize() metal.MTLSize
	SetSourceSize(value metal.MTLSize)
}

// Init initializes the instance.
func (c MPSCNNPoolingGradient) Init() MPSCNNPoolingGradient {
	rv := objc.Send[MPSCNNPoolingGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingGradient) Autorelease() MPSCNNPoolingGradient {
	rv := objc.Send[MPSCNNPoolingGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingGradient creates a new MPSCNNPoolingGradient instance.
func NewMPSCNNPoolingGradient() MPSCNNPoolingGradient {
	class := getMPSCNNPoolingGradientClass()
	rv := objc.Send[MPSCNNPoolingGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingGradient {
	instance := getMPSCNNPoolingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(coder:device:)
func NewCNNPoolingGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingGradient {
	instance := getMPSCNNPoolingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNPoolingGradientWithDevice(device metal.MTLDevice) MPSCNNPoolingGradient {
	instance := getMPSCNNPoolingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNPoolingGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingGradient {
	instance := getMPSCNNPoolingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingGradient {
	instance := getMPSCNNPoolingGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNPoolingGradient) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingGradient {
	rv := objc.Send[MPSCNNPoolingGradient](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func (c MPSCNNPoolingGradient) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingGradient {
	rv := objc.Send[MPSCNNPoolingGradient](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/sourceSize
func (c MPSCNNPoolingGradient) SourceSize() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("sourceSize"))
	return metal.MTLSize(rv)
}
func (c MPSCNNPoolingGradient) SetSourceSize(value metal.MTLSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setSourceSize:"), value)
}
