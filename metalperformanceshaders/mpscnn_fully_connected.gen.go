// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNFullyConnected] class.
var (
	_MPSCNNFullyConnectedClass     MPSCNNFullyConnectedClass
	_MPSCNNFullyConnectedClassOnce sync.Once
)

func getMPSCNNFullyConnectedClass() MPSCNNFullyConnectedClass {
	_MPSCNNFullyConnectedClassOnce.Do(func() {
		_MPSCNNFullyConnectedClass = MPSCNNFullyConnectedClass{class: objc.GetClass("MPSCNNFullyConnected")}
	})
	return _MPSCNNFullyConnectedClass
}

// GetMPSCNNFullyConnectedClass returns the class object for MPSCNNFullyConnected.
func GetMPSCNNFullyConnectedClass() MPSCNNFullyConnectedClass {
	return getMPSCNNFullyConnectedClass()
}

type MPSCNNFullyConnectedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNFullyConnectedClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNFullyConnectedClass) Alloc() MPSCNNFullyConnected {
	rv := objc.Send[MPSCNNFullyConnected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A fully connected convolution layer, also known as an inner product layer.
//
// # Overview
//
// A fully connected layer in a Convolutional Neural Network (CNN) is one
// where every input channel is connected to every output channel. The kernel
// width is equal to the width of the source image, and the kernel height is
// equal to the height of the source image. The width and height of the output
// is `1 x 1`.
//
// A fully connected layer takes an [MPSImage] object with dimensions
// `source.Width() x source.Height() x Ni`, convolves it with
// `Weights[No][source.Width()][source.Height()][Ni]`,` `and produces a `1 x 1
// x No` output.
//
// Thus, the following conditions must be true:
//
// - `kernelWidth == source.Width()` - `kernelHeight == source.Height()` -
// `clipRect.SizeXCUIElementTypeWidth() == 1` -
// `clipRect.SizeXCUIElementTypeHeight() == 1`
//
// You can think of a fully connected layer as a matrix multiplication where
// the image is flattened into a vector of length
// `source.Width()*source.Height()*Ni`, and the weights are arranged in a
// matrix of dimension `No x (source.Width()*source.Height()*Ni)` to produce
// an output vector of length [No].
//
// The value of the `strideInPixelsX`,
// [MPSCNNConvolutionDescriptor.StrideInPixelsY], and
// [MPSCNNConvolution.Groups] properties must be `1`. The
// [MPSCNNKernel.Offset] property is not applicable and it is ignored. Because
// the clip rectangle is clamped to the destination image bounds, if the
// destination is `1 x 1`, you do not need to set the [MPSCNNKernel.ClipRect]
// property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected
type MPSCNNFullyConnected struct {
	MPSCNNConvolution
}

// MPSCNNFullyConnectedFromID constructs a [MPSCNNFullyConnected] from an objc.ID.
//
// A fully connected convolution layer, also known as an inner product layer.
func MPSCNNFullyConnectedFromID(id objc.ID) MPSCNNFullyConnected {
	return MPSCNNFullyConnected{MPSCNNConvolution: MPSCNNConvolutionFromID(id)}
}

// NOTE: MPSCNNFullyConnected adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNFullyConnected] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected
type IMPSCNNFullyConnected interface {
	IMPSCNNConvolution
}

// Init initializes the instance.
func (c MPSCNNFullyConnected) Init() MPSCNNFullyConnected {
	rv := objc.Send[MPSCNNFullyConnected](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNFullyConnected) Autorelease() MPSCNNFullyConnected {
	rv := objc.Send[MPSCNNFullyConnected](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNFullyConnected creates a new MPSCNNFullyConnected instance.
func NewMPSCNNFullyConnected() MPSCNNFullyConnected {
	class := getMPSCNNFullyConnectedClass()
	rv := objc.Send[MPSCNNFullyConnected](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNFullyConnectedWithCoder(aDecoder foundation.INSCoder) MPSCNNFullyConnected {
	instance := getMPSCNNFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(coder:device:)
func NewCNNFullyConnectedWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNFullyConnected {
	instance := getMPSCNNFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNFullyConnectedFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNFullyConnectedWithDevice(device metal.MTLDevice) MPSCNNFullyConnected {
	instance := getMPSCNNFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer.
//
// device: The device on which this kernel will run.
//
// convolutionDescriptor: A valid convolution descriptor.
//
// The values of the````MPSCNNConvolutionDescriptor/strideInPixelsX“,
// “MPSCNNConvolutionDescriptor/strideInPixelsY“, and
// “MPSCNNConvolutionDescriptor/groups“ properties of the descriptor must be
// set to`1` (i.e. their default values).
//
// kernelWeights: A pointer to a weights array.
//
// Each entry is a float value. The number of entries is equal to
// [MPSCNNConvolutionDescriptor.InputFeatureChannels] `*`
// [MPSCNNConvolutionDescriptor.OutputFeatureChannels] `*`
// [MPSCNNConvolutionDescriptor.KernelHeight] `*`
// [MPSCNNConvolutionDescriptor.KernelWidth].
//
// The layout of the filter weights is arranged so that it can be
// reinterpreted as a 4D tensor (array)
// `weight[outputFeatureChannels][kernelHeight][kernelWidth][inputChannels/groups].`
//
// Weights are converted to half float precision (`fp16`) internally for best
// performance.
//
// biasTerms: A pointer to bias terms to be applied to the convolution output.
//
// Each entry is a float value. The number of entries is the number of output
// feature maps.
//
// flags: Currently unused.
//
// This value must be [MPSCNNConvolutionFlagsNone].
//
// # Return Value
//
// A valid [MPSCNNFullyConnected] object or `nil`, if failure.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(device:convolutionDescriptor:kernelWeights:biasTerms:flags:)
func NewCNNFullyConnectedWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags(device metal.MTLDevice, convolutionDescriptor IMPSCNNConvolutionDescriptor, kernelWeights *float32, biasTerms *float32, flags MPSCNNConvolutionFlags) MPSCNNFullyConnected {
	instance := getMPSCNNFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionDescriptor:kernelWeights:biasTerms:flags:"), device, convolutionDescriptor, kernelWeights, biasTerms, flags)
	return MPSCNNFullyConnectedFromID(rv)
}

// Initializes a fully connected convolution layer.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(device:weights:)
func NewCNNFullyConnectedWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNFullyConnected {
	instance := getMPSCNNFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNFullyConnectedFromID(rv)
}
