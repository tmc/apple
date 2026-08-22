// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingAverageGradient] class.
var (
	_MPSCNNPoolingAverageGradientClass     MPSCNNPoolingAverageGradientClass
	_MPSCNNPoolingAverageGradientClassOnce sync.Once
)

func getMPSCNNPoolingAverageGradientClass() MPSCNNPoolingAverageGradientClass {
	_MPSCNNPoolingAverageGradientClassOnce.Do(func() {
		_MPSCNNPoolingAverageGradientClass = MPSCNNPoolingAverageGradientClass{class: objc.GetClass("MPSCNNPoolingAverageGradient")}
	})
	return _MPSCNNPoolingAverageGradientClass
}

// GetMPSCNNPoolingAverageGradientClass returns the class object for MPSCNNPoolingAverageGradient.
func GetMPSCNNPoolingAverageGradientClass() MPSCNNPoolingAverageGradientClass {
	return getMPSCNNPoolingAverageGradientClass()
}

type MPSCNNPoolingAverageGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingAverageGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingAverageGradientClass) Alloc() MPSCNNPoolingAverageGradient {
	rv := objc.Send[MPSCNNPoolingAverageGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient average pooling filter.
//
// # Instance Properties
//
//   - [MPSCNNPoolingAverageGradient.ZeroPadSizeX]
//   - [MPSCNNPoolingAverageGradient.SetZeroPadSizeX]
//   - [MPSCNNPoolingAverageGradient.ZeroPadSizeY]
//   - [MPSCNNPoolingAverageGradient.SetZeroPadSizeY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient
type MPSCNNPoolingAverageGradient struct {
	MPSCNNPoolingGradient
}

// MPSCNNPoolingAverageGradientFromID constructs a [MPSCNNPoolingAverageGradient] from an objc.ID.
//
// A gradient average pooling filter.
func MPSCNNPoolingAverageGradientFromID(id objc.ID) MPSCNNPoolingAverageGradient {
	return MPSCNNPoolingAverageGradient{MPSCNNPoolingGradient: MPSCNNPoolingGradientFromID(id)}
}

// NOTE: MPSCNNPoolingAverageGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingAverageGradient] class.
//
// # Instance Properties
//
//   - [IMPSCNNPoolingAverageGradient.ZeroPadSizeX]
//   - [IMPSCNNPoolingAverageGradient.SetZeroPadSizeX]
//   - [IMPSCNNPoolingAverageGradient.ZeroPadSizeY]
//   - [IMPSCNNPoolingAverageGradient.SetZeroPadSizeY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient
type IMPSCNNPoolingAverageGradient interface {
	IMPSCNNPoolingGradient

	// Topic: Instance Properties

	ZeroPadSizeX() uint
	SetZeroPadSizeX(value uint)
	ZeroPadSizeY() uint
	SetZeroPadSizeY(value uint)
}

// Init initializes the instance.
func (c MPSCNNPoolingAverageGradient) Init() MPSCNNPoolingAverageGradient {
	rv := objc.Send[MPSCNNPoolingAverageGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingAverageGradient) Autorelease() MPSCNNPoolingAverageGradient {
	rv := objc.Send[MPSCNNPoolingAverageGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingAverageGradient creates a new MPSCNNPoolingAverageGradient instance.
func NewMPSCNNPoolingAverageGradient() MPSCNNPoolingAverageGradient {
	class := getMPSCNNPoolingAverageGradientClass()
	rv := objc.Send[MPSCNNPoolingAverageGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingAverageGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingAverageGradient {
	instance := getMPSCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingAverageGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient/init(coder:device:)
func NewCNNPoolingAverageGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingAverageGradient {
	instance := getMPSCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingAverageGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNPoolingAverageGradientWithDevice(device metal.MTLDevice) MPSCNNPoolingAverageGradient {
	instance := getMPSCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingAverageGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient/init(device:kernelWidth:kernelHeight:)
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingAverageGradient {
	instance := getMPSCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingAverageGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingAverageGradient {
	instance := getMPSCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingAverageGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient/zeroPadSizeX
func (c MPSCNNPoolingAverageGradient) ZeroPadSizeX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("zeroPadSizeX"))
	return rv
}
func (c MPSCNNPoolingAverageGradient) SetZeroPadSizeX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setZeroPadSizeX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient/zeroPadSizeY
func (c MPSCNNPoolingAverageGradient) ZeroPadSizeY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("zeroPadSizeY"))
	return rv
}
func (c MPSCNNPoolingAverageGradient) SetZeroPadSizeY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setZeroPadSizeY:"), value)
}
