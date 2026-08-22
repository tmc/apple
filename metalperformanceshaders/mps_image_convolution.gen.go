// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageConvolution] class.
var (
	_MPSImageConvolutionClass     MPSImageConvolutionClass
	_MPSImageConvolutionClassOnce sync.Once
)

func getMPSImageConvolutionClass() MPSImageConvolutionClass {
	_MPSImageConvolutionClassOnce.Do(func() {
		_MPSImageConvolutionClass = MPSImageConvolutionClass{class: objc.GetClass("MPSImageConvolution")}
	})
	return _MPSImageConvolutionClass
}

// GetMPSImageConvolutionClass returns the class object for MPSImageConvolution.
func GetMPSImageConvolutionClass() MPSImageConvolutionClass {
	return getMPSImageConvolutionClass()
}

type MPSImageConvolutionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageConvolutionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageConvolutionClass) Alloc() MPSImageConvolution {
	rv := objc.Send[MPSImageConvolution](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a given kernel of odd width and
// height.
//
// # Overview
//
// Filter width and height can be either 3, 5, 7 or 9. If there are multiple
// channels in the source image, each channel is processed independently.
//
// A separable convolution filter may perform better when done in two passes.
// . A convolution filter is separable if the ratio of filter values between
// all rows is constant over the whole row. For example, this edge detection
// filter:
//
// [media-2556905]
//
// Can instead be separated into the product of two vectors, like so:
//
// [media-2556912]
//
// And consequently can be done as two, one-dimensional convolution passes
// back to back on the same image. In this way, the number of multiplies
// (ignoring the fact that we could skip zeros here) is reduced from `3*3=9`
// to `3+3=6`. There are similar savings for addition. For large filters, the
// savings can be profound.
//
// # Methods
//
//   - [MPSImageConvolution.InitWithDeviceKernelWidthKernelHeightWeights]: Initializes a convolution filter.
//
// # Properties
//
//   - [MPSImageConvolution.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [MPSImageConvolution.KernelWidth]: The width of the filter window. Must be an odd number.
//   - [MPSImageConvolution.Bias]: The value added to a convolved pixel before it is converted back to its intended storage format.
//   - [MPSImageConvolution.SetBias]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution
type MPSImageConvolution struct {
	MPSUnaryImageKernel
}

// MPSImageConvolutionFromID constructs a [MPSImageConvolution] from an objc.ID.
//
// A filter that convolves an image with a given kernel of odd width and
// height.
func MPSImageConvolutionFromID(id objc.ID) MPSImageConvolution {
	return MPSImageConvolution{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageConvolution adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageConvolution] class.
//
// # Methods
//
//   - [IMPSImageConvolution.InitWithDeviceKernelWidthKernelHeightWeights]: Initializes a convolution filter.
//
// # Properties
//
//   - [IMPSImageConvolution.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [IMPSImageConvolution.KernelWidth]: The width of the filter window. Must be an odd number.
//   - [IMPSImageConvolution.Bias]: The value added to a convolved pixel before it is converted back to its intended storage format.
//   - [IMPSImageConvolution.SetBias]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution
type IMPSImageConvolution interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a convolution filter.
	InitWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImageConvolution

	// Topic: Properties

	// The height of the filter window. Must be an odd number.
	KernelHeight() uint
	// The width of the filter window. Must be an odd number.
	KernelWidth() uint
	// The value added to a convolved pixel before it is converted back to its intended storage format.
	Bias() float32
	SetBias(value float32)
}

// Init initializes the instance.
func (i MPSImageConvolution) Init() MPSImageConvolution {
	rv := objc.Send[MPSImageConvolution](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageConvolution) Autorelease() MPSImageConvolution {
	rv := objc.Send[MPSImageConvolution](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageConvolution creates a new MPSImageConvolution instance.
func NewMPSImageConvolution() MPSImageConvolution {
	class := getMPSImageConvolutionClass()
	rv := objc.Send[MPSImageConvolution](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageConvolutionWithCoder(aDecoder foundation.INSCoder) MPSImageConvolution {
	instance := getMPSImageConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/init(coder:device:)
func NewImageConvolutionWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageConvolution {
	instance := getMPSImageConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageConvolutionWithDevice(device metal.MTLDevice) MPSImageConvolution {
	instance := getMPSImageConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageConvolutionFromID(rv)
}

// Initializes a convolution filter.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// kernelWeights: A pointer to an array of `kernelWidth * kernelHeight` values to be used as
// the kernel. These values should be in row-major order.
//
// # Return Value
//
// An initialized convolution filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/init(device:kernelWidth:kernelHeight:weights:)
func NewImageConvolutionWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImageConvolution {
	instance := getMPSImageConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return MPSImageConvolutionFromID(rv)
}

// Initializes a convolution filter.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// kernelWeights: A pointer to an array of `kernelWidth * kernelHeight` values to be used as
// the kernel. These values should be in row-major order.
//
// # Return Value
//
// An initialized convolution filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/init(device:kernelWidth:kernelHeight:weights:)
func (i MPSImageConvolution) InitWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImageConvolution {
	rv := objc.Send[MPSImageConvolution](i.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return rv
}

// The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/kernelHeight
func (i MPSImageConvolution) KernelHeight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/kernelWidth
func (i MPSImageConvolution) KernelWidth() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelWidth"))
	return rv
}

// The value added to a convolved pixel before it is converted back to its
// intended storage format.
//
// # Discussion
//
// This value can be used to convert negative values into a representable
// range for an unsigned pixel format. For example, many edge detection
// filters produce results in the range `[-k,k]`. By scaling the filter
// weights by `0.5/k` and adding `0.5`, the results will be in the range
// `[0,1]` suitable for use with unsigned normalized formats.
//
// This value can be used in combination with renormalization of the filter
// weights to do video ranging as part of the convolution effect. It can also
// just be used to increase the brightness of the image.
//
// The default value is `0.0f`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/bias
func (i MPSImageConvolution) Bias() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("bias"))
	return rv
}
func (i MPSImageConvolution) SetBias(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setBias:"), value)
}
