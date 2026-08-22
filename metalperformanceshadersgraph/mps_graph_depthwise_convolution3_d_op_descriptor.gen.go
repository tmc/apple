// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraphDepthwiseConvolution3DOpDescriptor] class.
var (
	_MPSGraphDepthwiseConvolution3DOpDescriptorClass     MPSGraphDepthwiseConvolution3DOpDescriptorClass
	_MPSGraphDepthwiseConvolution3DOpDescriptorClassOnce sync.Once
)

func getMPSGraphDepthwiseConvolution3DOpDescriptorClass() MPSGraphDepthwiseConvolution3DOpDescriptorClass {
	_MPSGraphDepthwiseConvolution3DOpDescriptorClassOnce.Do(func() {
		_MPSGraphDepthwiseConvolution3DOpDescriptorClass = MPSGraphDepthwiseConvolution3DOpDescriptorClass{class: objc.GetClass("MPSGraphDepthwiseConvolution3DOpDescriptor")}
	})
	return _MPSGraphDepthwiseConvolution3DOpDescriptorClass
}

// GetMPSGraphDepthwiseConvolution3DOpDescriptorClass returns the class object for MPSGraphDepthwiseConvolution3DOpDescriptor.
func GetMPSGraphDepthwiseConvolution3DOpDescriptorClass() MPSGraphDepthwiseConvolution3DOpDescriptorClass {
	return getMPSGraphDepthwiseConvolution3DOpDescriptorClass()
}

type MPSGraphDepthwiseConvolution3DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphDepthwiseConvolution3DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphDepthwiseConvolution3DOpDescriptorClass) Alloc() MPSGraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution3DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a 3D-depthwise convolution
// operation.
//
// # Overview
//
// A [MPSGraphDepthwiseConvolution3DOpDescriptor] defines constant parameters
// for 3D depthwise convolutions. Use this class with
// [MPSGraph.DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName],
// [MPSGraph.DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName]
// and
// [MPSGraph.DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]
// methods.
//
// # Instance Properties
//
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.ChannelDimensionIndex]: The axis that contains the channels in the input and the weights, within the 4D tile of the last dimensions.
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.SetChannelDimensionIndex]
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.DilationRates]: The dilation rates for spatial dimensions.
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.SetDilationRates]
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.PaddingStyle]: The padding style for the operation.
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.PaddingValues]: The padding values for spatial dimensions.
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.SetPaddingValues]
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.Strides]: The strides for spatial dimensions.
//   - [MPSGraphDepthwiseConvolution3DOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor
type MPSGraphDepthwiseConvolution3DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphDepthwiseConvolution3DOpDescriptorFromID constructs a [MPSGraphDepthwiseConvolution3DOpDescriptor] from an objc.ID.
//
// The class that defines the parameters for a 3D-depthwise convolution
// operation.
func MPSGraphDepthwiseConvolution3DOpDescriptorFromID(id objc.ID) MPSGraphDepthwiseConvolution3DOpDescriptor {
	return MPSGraphDepthwiseConvolution3DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphDepthwiseConvolution3DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphDepthwiseConvolution3DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.ChannelDimensionIndex]: The axis that contains the channels in the input and the weights, within the 4D tile of the last dimensions.
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.SetChannelDimensionIndex]
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.DilationRates]: The dilation rates for spatial dimensions.
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.SetDilationRates]
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.PaddingStyle]: The padding style for the operation.
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.PaddingValues]: The padding values for spatial dimensions.
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.SetPaddingValues]
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.Strides]: The strides for spatial dimensions.
//   - [IMPSGraphDepthwiseConvolution3DOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor
type IMPSGraphDepthwiseConvolution3DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The axis that contains the channels in the input and the weights, within the 4D tile of the last dimensions.
	ChannelDimensionIndex() int
	SetChannelDimensionIndex(value int)
	// The dilation rates for spatial dimensions.
	DilationRates() []foundation.NSNumber
	SetDilationRates(value []foundation.NSNumber)
	// The padding style for the operation.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// The padding values for spatial dimensions.
	PaddingValues() []foundation.NSNumber
	SetPaddingValues(value []foundation.NSNumber)
	// The strides for spatial dimensions.
	Strides() []foundation.NSNumber
	SetStrides(value []foundation.NSNumber)
}

// Init initializes the instance.
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) Init() MPSGraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution3DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) Autorelease() MPSGraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution3DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphDepthwiseConvolution3DOpDescriptor creates a new MPSGraphDepthwiseConvolution3DOpDescriptor instance.
func NewMPSGraphDepthwiseConvolution3DOpDescriptor() MPSGraphDepthwiseConvolution3DOpDescriptor {
	class := getMPSGraphDepthwiseConvolution3DOpDescriptorClass()
	rv := objc.Send[MPSGraphDepthwiseConvolution3DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a 3D depthwise convolution descriptor with default values.
//
// paddingStyle: See `paddingStyle` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(paddingStyle:)
func NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle(paddingStyle MPSGraphPaddingStyle) MPSGraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphDepthwiseConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return MPSGraphDepthwiseConvolution3DOpDescriptorFromID(rv)
}

// Creates a 3D depthwise convolution descriptor with given values.
//
// strides: See `strides` property.
//
// dilationRates: See `dilationRates` property.
//
// paddingValues: See `paddingValues` property.
//
// paddingStyle: See `paddingStyle` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(strides:dilationRates:paddingValues:paddingStyle:)
func NewGraphDepthwiseConvolution3DOpDescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides []foundation.NSNumber, dilationRates []foundation.NSNumber, paddingValues []foundation.NSNumber, paddingStyle MPSGraphPaddingStyle) MPSGraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphDepthwiseConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrides:dilationRates:paddingValues:paddingStyle:"), objectivec.IObjectSliceToNSArray(strides), objectivec.IObjectSliceToNSArray(dilationRates), objectivec.IObjectSliceToNSArray(paddingValues), paddingStyle)
	return MPSGraphDepthwiseConvolution3DOpDescriptorFromID(rv)
}

// The axis that contains the channels in the input and the weights, within
// the 4D tile of the last dimensions.
//
// # Discussion
//
// For example the value of `-1` corresponds to [NDHWC], [NHWC] layouts. This
// allows the placement of the channel index anywhere within the last 4
// dimensions of the tensor. In case your weights are in a different layout
// you can bring them to the same layout as inputs using transposes or
// permutations. Default value: `-4`, corresponds to [NCDHW] and [CDHW]
// layouts.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/channelDimensionIndex
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) ChannelDimensionIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("channelDimensionIndex"))
	return rv
}
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) SetChannelDimensionIndex(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setChannelDimensionIndex:"), value)
}

// The dilation rates for spatial dimensions.
//
// # Discussion
//
// Must be three numbers, one for each spatial dimension, fastest running
// index last. Default value: `@[ @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/dilationRates
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) DilationRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("dilationRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) SetDilationRates(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRates:"), objectivec.IObjectSliceToNSArray(value))
}

// The padding style for the operation.
//
// # Discussion
//
// Default value: [MPSGraphPaddingStyleExplicit].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingStyle
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// The padding values for spatial dimensions.
//
// # Discussion
//
// Must be six numbers, two for each spatial dimension. For example
// `paddingValues[0]` defines the explicit padding amount before the first
// spatial dimension (slowest running index of spatial dimensions),
// `paddingValues[1]` defines the padding amount after the first spatial
// dimension etc. Use only with `paddingStyle = MPSGraphPaddingStyleExplicit`.
// Default value: `@[ @0, @0, @0, @0, @0, @0 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingValues
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) PaddingValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("paddingValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) SetPaddingValues(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingValues:"), objectivec.IObjectSliceToNSArray(value))
}

// The strides for spatial dimensions.
//
// # Discussion
//
// Must be three numbers, one for each spatial dimension, fastest running
// index last. Default value: `@[ @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/strides
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) Strides() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("strides"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphDepthwiseConvolution3DOpDescriptor) SetStrides(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrides:"), objectivec.IObjectSliceToNSArray(value))
}
