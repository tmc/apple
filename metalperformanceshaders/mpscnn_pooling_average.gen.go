// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingAverage] class.
var (
	_MPSCNNPoolingAverageClass     MPSCNNPoolingAverageClass
	_MPSCNNPoolingAverageClassOnce sync.Once
)

func getMPSCNNPoolingAverageClass() MPSCNNPoolingAverageClass {
	_MPSCNNPoolingAverageClassOnce.Do(func() {
		_MPSCNNPoolingAverageClass = MPSCNNPoolingAverageClass{class: objc.GetClass("MPSCNNPoolingAverage")}
	})
	return _MPSCNNPoolingAverageClass
}

// GetMPSCNNPoolingAverageClass returns the class object for MPSCNNPoolingAverage.
func GetMPSCNNPoolingAverageClass() MPSCNNPoolingAverageClass {
	return getMPSCNNPoolingAverageClass()
}

type MPSCNNPoolingAverageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingAverageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingAverageClass) Alloc() MPSCNNPoolingAverage {
	rv := objc.Send[MPSCNNPoolingAverage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An average pooling filter.
//
// # Overview
//
// For each pixel in an image, the filter returns the average value of the
// pixels in the filter region defined by `kernelWidth` `x` `kernelHeight`.
//
// When the value of the [MPSCNNKernel.EdgeMode] property is set to
// [MPSImageEdgeModeClamp], the filtering window is shrunk to remain within
// the source image borders. For pixels close to the image borders, the
// filtering window will be smaller in order to fit inside the source image
// and less values will be used to compute the average value. In case the
// filtering window is entirely outside the source image border, the output
// value will be `0`.
//
// # Instance Properties
//
//   - [MPSCNNPoolingAverage.ZeroPadSizeX]
//   - [MPSCNNPoolingAverage.SetZeroPadSizeX]
//   - [MPSCNNPoolingAverage.ZeroPadSizeY]
//   - [MPSCNNPoolingAverage.SetZeroPadSizeY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage
type MPSCNNPoolingAverage struct {
	MPSCNNPooling
}

// MPSCNNPoolingAverageFromID constructs a [MPSCNNPoolingAverage] from an objc.ID.
//
// An average pooling filter.
func MPSCNNPoolingAverageFromID(id objc.ID) MPSCNNPoolingAverage {
	return MPSCNNPoolingAverage{MPSCNNPooling: MPSCNNPoolingFromID(id)}
}

// NOTE: MPSCNNPoolingAverage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingAverage] class.
//
// # Instance Properties
//
//   - [IMPSCNNPoolingAverage.ZeroPadSizeX]
//   - [IMPSCNNPoolingAverage.SetZeroPadSizeX]
//   - [IMPSCNNPoolingAverage.ZeroPadSizeY]
//   - [IMPSCNNPoolingAverage.SetZeroPadSizeY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage
type IMPSCNNPoolingAverage interface {
	IMPSCNNPooling

	// Topic: Instance Properties

	ZeroPadSizeX() uint
	SetZeroPadSizeX(value uint)
	ZeroPadSizeY() uint
	SetZeroPadSizeY(value uint)
}

// Init initializes the instance.
func (c MPSCNNPoolingAverage) Init() MPSCNNPoolingAverage {
	rv := objc.Send[MPSCNNPoolingAverage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingAverage) Autorelease() MPSCNNPoolingAverage {
	rv := objc.Send[MPSCNNPoolingAverage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingAverage creates a new MPSCNNPoolingAverage instance.
func NewMPSCNNPoolingAverage() MPSCNNPoolingAverage {
	class := getMPSCNNPoolingAverageClass()
	rv := objc.Send[MPSCNNPoolingAverage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingAverageWithCoder(aDecoder foundation.INSCoder) MPSCNNPoolingAverage {
	instance := getMPSCNNPoolingAverageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingAverageFromID(rv)
}

// Initializes an average pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage/init(coder:device:)
func NewCNNPoolingAverageWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPoolingAverage {
	instance := getMPSCNNPoolingAverageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingAverageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNPoolingAverageWithDevice(device metal.MTLDevice) MPSCNNPoolingAverage {
	instance := getMPSCNNPoolingAverageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingAverageFromID(rv)
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
func NewCNNPoolingAverageWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPoolingAverage {
	instance := getMPSCNNPoolingAverageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingAverageFromID(rv)
}

// Initializes an average pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingAverageWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingAverage {
	instance := getMPSCNNPoolingAverageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingAverageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage/zeroPadSizeX
func (c MPSCNNPoolingAverage) ZeroPadSizeX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("zeroPadSizeX"))
	return rv
}
func (c MPSCNNPoolingAverage) SetZeroPadSizeX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setZeroPadSizeX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage/zeroPadSizeY
func (c MPSCNNPoolingAverage) ZeroPadSizeY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("zeroPadSizeY"))
	return rv
}
func (c MPSCNNPoolingAverage) SetZeroPadSizeY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setZeroPadSizeY:"), value)
}
