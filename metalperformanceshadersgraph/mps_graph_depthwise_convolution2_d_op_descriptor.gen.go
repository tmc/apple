// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphDepthwiseConvolution2DOpDescriptor] class.
var (
	_MPSGraphDepthwiseConvolution2DOpDescriptorClass     MPSGraphDepthwiseConvolution2DOpDescriptorClass
	_MPSGraphDepthwiseConvolution2DOpDescriptorClassOnce sync.Once
)

func getMPSGraphDepthwiseConvolution2DOpDescriptorClass() MPSGraphDepthwiseConvolution2DOpDescriptorClass {
	_MPSGraphDepthwiseConvolution2DOpDescriptorClassOnce.Do(func() {
		_MPSGraphDepthwiseConvolution2DOpDescriptorClass = MPSGraphDepthwiseConvolution2DOpDescriptorClass{class: objc.GetClass("MPSGraphDepthwiseConvolution2DOpDescriptor")}
	})
	return _MPSGraphDepthwiseConvolution2DOpDescriptorClass
}

// GetMPSGraphDepthwiseConvolution2DOpDescriptorClass returns the class object for MPSGraphDepthwiseConvolution2DOpDescriptor.
func GetMPSGraphDepthwiseConvolution2DOpDescriptorClass() MPSGraphDepthwiseConvolution2DOpDescriptorClass {
	return getMPSGraphDepthwiseConvolution2DOpDescriptorClass()
}

type MPSGraphDepthwiseConvolution2DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphDepthwiseConvolution2DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphDepthwiseConvolution2DOpDescriptorClass) Alloc() MPSGraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution2DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that defines the parameters for a 2D-depthwise convolution
// operation.
//
// # Overview
//
// An [MPSGraphDepthwiseConvolution2DOpDescriptor] defines constant parameters
// for 2D-depthwise convolutions. Use this class with
// [MPSGraph.DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName],
// [MPSGraph.DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName],
// and
// [MPSGraph.DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName]
// methods.
//
// # Instance Properties
//
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.DataLayout]: The data layout of the input data in the forward pass.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetDataLayout]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.DilationRateInX]: The dilation rate for the x dimension.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetDilationRateInX]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.DilationRateInY]: The dilation rate for the y dimension.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetDilationRateInY]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.PaddingBottom]: The explicit padding value for the y dimension operation adds after the data.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingBottom]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.PaddingLeft]: The explicit padding value for the x dimension the operation adds before the data.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingLeft]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.PaddingRight]: The explicit padding value for the x dimension operation adds after the data.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingRight]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.PaddingStyle]: The padding style for the operation.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.PaddingTop]: The explicit padding value for the y dimension operation adds before the data.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingTop]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.StrideInX]: The stride for the x dimension.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetStrideInX]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.StrideInY]: The stride for the y dimension.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetStrideInY]
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.WeightsLayout]: The data layout of the weights.
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [MPSGraphDepthwiseConvolution2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the explicit padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor
type MPSGraphDepthwiseConvolution2DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphDepthwiseConvolution2DOpDescriptorFromID constructs a [MPSGraphDepthwiseConvolution2DOpDescriptor] from an objc.ID.
//
// A class that defines the parameters for a 2D-depthwise convolution
// operation.
func MPSGraphDepthwiseConvolution2DOpDescriptorFromID(id objc.ID) MPSGraphDepthwiseConvolution2DOpDescriptor {
	return MPSGraphDepthwiseConvolution2DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphDepthwiseConvolution2DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphDepthwiseConvolution2DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.DataLayout]: The data layout of the input data in the forward pass.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetDataLayout]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.DilationRateInX]: The dilation rate for the x dimension.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetDilationRateInX]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.DilationRateInY]: The dilation rate for the y dimension.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetDilationRateInY]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.PaddingBottom]: The explicit padding value for the y dimension operation adds after the data.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingBottom]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.PaddingLeft]: The explicit padding value for the x dimension the operation adds before the data.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingLeft]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.PaddingRight]: The explicit padding value for the x dimension operation adds after the data.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingRight]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.PaddingStyle]: The padding style for the operation.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.PaddingTop]: The explicit padding value for the y dimension operation adds before the data.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetPaddingTop]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.StrideInX]: The stride for the x dimension.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetStrideInX]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.StrideInY]: The stride for the y dimension.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetStrideInY]
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.WeightsLayout]: The data layout of the weights.
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [IMPSGraphDepthwiseConvolution2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the explicit padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor
type IMPSGraphDepthwiseConvolution2DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The data layout of the input data in the forward pass.
	DataLayout() MPSGraphTensorNamedDataLayout
	SetDataLayout(value MPSGraphTensorNamedDataLayout)
	// The dilation rate for the x dimension.
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	// The dilation rate for the y dimension.
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	// The explicit padding value for the y dimension operation adds after the data.
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	// The explicit padding value for the x dimension the operation adds before the data.
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	// The explicit padding value for the x dimension operation adds after the data.
	PaddingRight() uint
	SetPaddingRight(value uint)
	// The padding style for the operation.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// The explicit padding value for the y dimension operation adds before the data.
	PaddingTop() uint
	SetPaddingTop(value uint)
	// The stride for the x dimension.
	StrideInX() uint
	SetStrideInX(value uint)
	// The stride for the y dimension.
	StrideInY() uint
	SetStrideInY(value uint)
	// The data layout of the weights.
	WeightsLayout() MPSGraphTensorNamedDataLayout
	SetWeightsLayout(value MPSGraphTensorNamedDataLayout)

	// Topic: Instance Methods

	// Sets the explicit padding values.
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
}

// Init initializes the instance.
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) Init() MPSGraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution2DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) Autorelease() MPSGraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphDepthwiseConvolution2DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphDepthwiseConvolution2DOpDescriptor creates a new MPSGraphDepthwiseConvolution2DOpDescriptor instance.
func NewMPSGraphDepthwiseConvolution2DOpDescriptor() MPSGraphDepthwiseConvolution2DOpDescriptor {
	class := getMPSGraphDepthwiseConvolution2DOpDescriptorClass()
	rv := objc.Send[MPSGraphDepthwiseConvolution2DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a 2D-depthwise convolution descriptor with given properties and
// default values.
//
// dataLayout: See `dataLayout` property.
//
// weightsLayout: See `weightsLayout` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(dataLayout:weightsLayout:)
func NewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout(dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphDepthwiseConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithDataLayout:weightsLayout:"), dataLayout, weightsLayout)
	return MPSGraphDepthwiseConvolution2DOpDescriptorFromID(rv)
}

// Creates a 2D-depthwise convolution descriptor with given values.
//
// strideInX: See `strideInX` property.
//
// strideInY: See `strideInY` property.
//
// dilationRateInX: See `dilationRateInX` property.
//
// dilationRateInY: See `dilationRateInY` property.
//
// paddingLeft: See `paddingLeft` property.
//
// paddingRight: See `paddingRight` property.
//
// paddingTop: See `paddingTop` property.
//
// paddingBottom: See `paddingBottom` property.
//
// paddingStyle: See `paddingStyle` property.
//
// dataLayout: See `dataLayout` property.
//
// weightsLayout: See `weightsLayout` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphDepthwiseConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return MPSGraphDepthwiseConvolution2DOpDescriptorFromID(rv)
}

// Sets the explicit padding values.
//
// paddingLeft: See `paddingLeft` property.
//
// paddingRight: See `paddingRight` property.
//
// paddingTop: See `paddingTop` property.
//
// paddingBottom: See `paddingBottom` property.
//
// # Discussion
//
// Note: this method also sets `paddingStyle` to
// [MPSGraphPaddingStyleExplicit] (see [MPSGraphPaddingStyle]).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
//
// [MPSGraphPaddingStyle]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPaddingStyle
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

// The data layout of the input data in the forward pass.
//
// # Discussion
//
// See: [MPSGraphTensorNamedDataLayout].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dataLayout
//
// [MPSGraphTensorNamedDataLayout]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) DataLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("dataLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetDataLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataLayout:"), value)
}

// The dilation rate for the x dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInX
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInX"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInX:"), value)
}

// The dilation rate for the y dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInY
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInY"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInY:"), value)
}

// The explicit padding value for the y dimension operation adds after the
// data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingBottom
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBottom"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBottom:"), value)
}

// The explicit padding value for the x dimension the operation adds before
// the data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingLeft
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingLeft"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingLeft:"), value)
}

// The explicit padding value for the x dimension operation adds after the
// data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingRight
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingRight"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingRight:"), value)
}

// The padding style for the operation.
//
// # Discussion
//
// Default value is [MPSGraphPaddingStyleExplicit].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingStyle
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// The explicit padding value for the y dimension operation adds before the
// data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingTop
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingTop"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingTop:"), value)
}

// The stride for the x dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInX
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInX"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInX:"), value)
}

// The stride for the y dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInY
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInY"))
	return rv
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInY:"), value)
}

// The data layout of the weights.
//
// # Discussion
//
// NOTE: ‘O’ index is channel multiplier index. See:
// [MPSGraphTensorNamedDataLayout].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/weightsLayout
//
// [MPSGraphTensorNamedDataLayout]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) WeightsLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("weightsLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphDepthwiseConvolution2DOpDescriptor) SetWeightsLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setWeightsLayout:"), value)
}
