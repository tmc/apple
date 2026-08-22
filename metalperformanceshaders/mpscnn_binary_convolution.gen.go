// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBinaryConvolution] class.
var (
	_MPSCNNBinaryConvolutionClass     MPSCNNBinaryConvolutionClass
	_MPSCNNBinaryConvolutionClassOnce sync.Once
)

func getMPSCNNBinaryConvolutionClass() MPSCNNBinaryConvolutionClass {
	_MPSCNNBinaryConvolutionClassOnce.Do(func() {
		_MPSCNNBinaryConvolutionClass = MPSCNNBinaryConvolutionClass{class: objc.GetClass("MPSCNNBinaryConvolution")}
	})
	return _MPSCNNBinaryConvolutionClass
}

// GetMPSCNNBinaryConvolutionClass returns the class object for MPSCNNBinaryConvolution.
func GetMPSCNNBinaryConvolutionClass() MPSCNNBinaryConvolutionClass {
	return getMPSCNNBinaryConvolutionClass()
}

type MPSCNNBinaryConvolutionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBinaryConvolutionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryConvolutionClass) Alloc() MPSCNNBinaryConvolution {
	rv := objc.Send[MPSCNNBinaryConvolution](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A convolution kernel with binary weights and an input image using binary
// approximations.
//
// # Overview
//
// The [MPSCNNBinaryConvolution] optionally first binarizes the input image
// and then convolves the result with a set of binary-valued filters, each
// producing one feature map in the output image (which is a normal image).
//
// The output is computed as follows:
//
// [media-2903520]
//
// where the sum over dx,dy is over the spatial filter kernel window defined
// by [MPSCNNConvolutionDescriptor.KernelWidth] and
// [MPSCNNConvolutionDescriptor.KernelHeight], sum over f is over the input
// feature channel indices within group, B contains the binary weights,
// interpreted as `{-1, 1}` or `{0, 1}` and scale[c] is the `outputScaleTerms`
// array and bias is the `outputBiasTerms` array. Above i is the image index
// in batch the sum over input channels f runs through the group indices. The
// convolution operator ⊗ is defined by [MPSCNNBinaryConvolutionType] passed
// in at initialization time of the filter:
//
// [MPSCNNBinaryConvolutionTypeBinaryWeights]: The input image is not
// binarized at all and the convolution is computed interpreting the weights
// as `[0, 1] -> {-1, 1}` with the given scaling terms.
// [MPSCNNBinaryConvolutionTypeXNOR]: The convolution is computed by first
// binarizing the input image using the sign function `bin(x) = x < 0 ? -1 :
// 1` and the convolution multiplication is done with the XNOR-operator:
//
// `!(x ^ y) = delta_xy = { (x == y) ? 1 : 0 }`
//
// and scaled according to the optional scaling operations.
//
// Note that we output the values of the bitwise convolutions to interval
// `{-1, 1}`, which means that the output of the XNOR-operator is scaled
// implicitly as follows:
//
// `r = 2 * ( !(x ^ y) ) - 1 = { -1, 1 }`
//
// This means that for a dot-product of two 32-bit words the result is:
//
// `r = 2 * popcount(!(x ^ y) ) - 32 = 32 - 2 * popcount( x ^ y ) = { -32,
// -30, ..., 30, 32 }`
//
// [MPSCNNBinaryConvolutionTypeAND]: The convolution is computed by first
// binarizing the input image using the sign function `bin(x) = x < 0 ? -1 :
// 1` and the convolution multiplication is done with the AND-operator:
//
// `(x & y) = delta_xy * delta_x1 = { (x == y == 1) ? 1 : 0 }`
//
// and scaled according to the optional scaling operations.
//
// Note that we output the values of the AND-operation is assumed to lie in
// `{0, 1}` interval and hence no more implicit scaling takes place.
//
// This means that for a dot-product of two 32-bit words the result is:
//
// `r = popcount(x & y) = { 0, ..., 31, 32 }`
//
// The input data can be pre-offset and scaled by providing the
// `inputBiasTerms` and `inputScaleTerms` parameters for the initialization
// functions and this can be used for example to accomplish batch
// normalization of the data. The scaling of input values happens before
// possible beta-image computation.
//
// The parameter `beta` above is an optional image which is used to compute
// scaling factors for each spatial position and image index. For the XNOR-Net
// based networks this is computed as follows:
//
// [media-2903518]
//
// where (dx,dy) are summed over the convolution filter window.
//
// [media-2903519]
//
// where in is the original input image (in full precision) and Nc is the
// number of input channels in the input image. Parameter `beta` is not passed
// as input and to enable beta-scaling the user can provide
// [MPSCNNBinaryConvolutionFlagsUseBetaScaling] in the flags parameter in the
// initialization functions.
//
// Finally the normal activation neuron is applied and the result is written
// to the output image.
//
// # Initializers
//
//   - [MPSCNNBinaryConvolution.InitWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags]: Initializes a binary convolution kernel.
//   - [MPSCNNBinaryConvolution.InitWithDeviceConvolutionDataScaleValueTypeFlags]: Initializes a binary convolution kernel.
//
// # Instance Properties
//
//   - [MPSCNNBinaryConvolution.InputFeatureChannels]
//   - [MPSCNNBinaryConvolution.OutputFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution
//
// [MPSCNNBinaryConvolutionType]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionType
type MPSCNNBinaryConvolution struct {
	MPSCNNKernel
}

// MPSCNNBinaryConvolutionFromID constructs a [MPSCNNBinaryConvolution] from an objc.ID.
//
// A convolution kernel with binary weights and an input image using binary
// approximations.
func MPSCNNBinaryConvolutionFromID(id objc.ID) MPSCNNBinaryConvolution {
	return MPSCNNBinaryConvolution{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNBinaryConvolution adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBinaryConvolution] class.
//
// # Initializers
//
//   - [IMPSCNNBinaryConvolution.InitWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags]: Initializes a binary convolution kernel.
//   - [IMPSCNNBinaryConvolution.InitWithDeviceConvolutionDataScaleValueTypeFlags]: Initializes a binary convolution kernel.
//
// # Instance Properties
//
//   - [IMPSCNNBinaryConvolution.InputFeatureChannels]
//   - [IMPSCNNBinaryConvolution.OutputFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution
type IMPSCNNBinaryConvolution interface {
	IMPSCNNKernel

	// Topic: Initializers

	// Initializes a binary convolution kernel.
	InitWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution
	// Initializes a binary convolution kernel.
	InitWithDeviceConvolutionDataScaleValueTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution

	// Topic: Instance Properties

	InputFeatureChannels() uint
	OutputFeatureChannels() uint
}

// Init initializes the instance.
func (c MPSCNNBinaryConvolution) Init() MPSCNNBinaryConvolution {
	rv := objc.Send[MPSCNNBinaryConvolution](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryConvolution) Autorelease() MPSCNNBinaryConvolution {
	rv := objc.Send[MPSCNNBinaryConvolution](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryConvolution creates a new MPSCNNBinaryConvolution instance.
func NewMPSCNNBinaryConvolution() MPSCNNBinaryConvolution {
	class := getMPSCNNBinaryConvolutionClass()
	rv := objc.Send[MPSCNNBinaryConvolution](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBinaryConvolutionWithCoder(aDecoder foundation.INSCoder) MPSCNNBinaryConvolution {
	instance := getMPSCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBinaryConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/init(coder:device:)
func NewCNNBinaryConvolutionWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBinaryConvolution {
	instance := getMPSCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBinaryConvolutionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNBinaryConvolutionWithDevice(device metal.MTLDevice) MPSCNNBinaryConvolution {
	instance := getMPSCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBinaryConvolutionFromID(rv)
}

// Initializes a binary convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/init(device:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func NewCNNBinaryConvolutionWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution {
	instance := getMPSCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), device, convolutionData, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return MPSCNNBinaryConvolutionFromID(rv)
}

// Initializes a binary convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/init(device:convolutionData:scaleValue:type:flags:)
func NewCNNBinaryConvolutionWithDeviceConvolutionDataScaleValueTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution {
	instance := getMPSCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), device, convolutionData, scaleValue, type_, flags)
	return MPSCNNBinaryConvolutionFromID(rv)
}

// Initializes a binary convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/init(device:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:)
func (c MPSCNNBinaryConvolution) InitWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, outputBiasTerms *float32, outputScaleTerms *float32, inputBiasTerms *float32, inputScaleTerms *float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution {
	rv := objc.Send[MPSCNNBinaryConvolution](c.ID, objc.Sel("initWithDevice:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), device, convolutionData, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return rv
}

// Initializes a binary convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/init(device:convolutionData:scaleValue:type:flags:)
func (c MPSCNNBinaryConvolution) InitWithDeviceConvolutionDataScaleValueTypeFlags(device metal.MTLDevice, convolutionData MPSCNNConvolutionDataSource, scaleValue float32, type_ MPSCNNBinaryConvolutionType, flags MPSCNNBinaryConvolutionFlags) MPSCNNBinaryConvolution {
	rv := objc.Send[MPSCNNBinaryConvolution](c.ID, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), device, convolutionData, scaleValue, type_, flags)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/inputFeatureChannels
func (c MPSCNNBinaryConvolution) InputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution/outputFeatureChannels
func (c MPSCNNBinaryConvolution) OutputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("outputFeatureChannels"))
	return rv
}
