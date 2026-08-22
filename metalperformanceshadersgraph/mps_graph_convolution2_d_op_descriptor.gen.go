// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphConvolution2DOpDescriptor] class.
var (
	_MPSGraphConvolution2DOpDescriptorClass     MPSGraphConvolution2DOpDescriptorClass
	_MPSGraphConvolution2DOpDescriptorClassOnce sync.Once
)

func getMPSGraphConvolution2DOpDescriptorClass() MPSGraphConvolution2DOpDescriptorClass {
	_MPSGraphConvolution2DOpDescriptorClassOnce.Do(func() {
		_MPSGraphConvolution2DOpDescriptorClass = MPSGraphConvolution2DOpDescriptorClass{class: objc.GetClass("MPSGraphConvolution2DOpDescriptor")}
	})
	return _MPSGraphConvolution2DOpDescriptorClass
}

// GetMPSGraphConvolution2DOpDescriptorClass returns the class object for MPSGraphConvolution2DOpDescriptor.
func GetMPSGraphConvolution2DOpDescriptorClass() MPSGraphConvolution2DOpDescriptorClass {
	return getMPSGraphConvolution2DOpDescriptorClass()
}

type MPSGraphConvolution2DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphConvolution2DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphConvolution2DOpDescriptorClass) Alloc() MPSGraphConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution2DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that describes the properties of a 2D-convolution operator.
//
// # Overview
//
// Use an instance of this class is to add a 2D-convolution operator with the
// desired properties to the graph.
//
// # Instance Properties
//
//   - [MPSGraphConvolution2DOpDescriptor.DataLayout]: The named layout of data in the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetDataLayout]
//   - [MPSGraphConvolution2DOpDescriptor.DilationRateInX]: The amount by which the weights tensor expands in the `x`-direction.
//   - [MPSGraphConvolution2DOpDescriptor.SetDilationRateInX]
//   - [MPSGraphConvolution2DOpDescriptor.DilationRateInY]: The amount by which the weights tensor expands in the `y`-direction.
//   - [MPSGraphConvolution2DOpDescriptor.SetDilationRateInY]
//   - [MPSGraphConvolution2DOpDescriptor.Groups]: The number of partitions of the input and output channels.
//   - [MPSGraphConvolution2DOpDescriptor.SetGroups]
//   - [MPSGraphConvolution2DOpDescriptor.PaddingBottom]: The number of zeros added at the bottom of the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetPaddingBottom]
//   - [MPSGraphConvolution2DOpDescriptor.PaddingLeft]: The number of zeros added on the left side of the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetPaddingLeft]
//   - [MPSGraphConvolution2DOpDescriptor.PaddingRight]: The number of zeros added on the right side of the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetPaddingRight]
//   - [MPSGraphConvolution2DOpDescriptor.PaddingStyle]: The type of padding applied to the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphConvolution2DOpDescriptor.PaddingTop]: The number of zeros added at the top of the source tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetPaddingTop]
//   - [MPSGraphConvolution2DOpDescriptor.StrideInX]: The scale that maps `x`-coordinate of the destination to `x`-coordinate of the source.
//   - [MPSGraphConvolution2DOpDescriptor.SetStrideInX]
//   - [MPSGraphConvolution2DOpDescriptor.StrideInY]: The scale that maps `y`-coordinate of the destination to `y`-coordinate of the source.
//   - [MPSGraphConvolution2DOpDescriptor.SetStrideInY]
//   - [MPSGraphConvolution2DOpDescriptor.WeightsLayout]: The named layout of data in the weights tensor.
//   - [MPSGraphConvolution2DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [MPSGraphConvolution2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the left, right, top, and bottom padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor
type MPSGraphConvolution2DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphConvolution2DOpDescriptorFromID constructs a [MPSGraphConvolution2DOpDescriptor] from an objc.ID.
//
// A class that describes the properties of a 2D-convolution operator.
func MPSGraphConvolution2DOpDescriptorFromID(id objc.ID) MPSGraphConvolution2DOpDescriptor {
	return MPSGraphConvolution2DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphConvolution2DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphConvolution2DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphConvolution2DOpDescriptor.DataLayout]: The named layout of data in the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetDataLayout]
//   - [IMPSGraphConvolution2DOpDescriptor.DilationRateInX]: The amount by which the weights tensor expands in the `x`-direction.
//   - [IMPSGraphConvolution2DOpDescriptor.SetDilationRateInX]
//   - [IMPSGraphConvolution2DOpDescriptor.DilationRateInY]: The amount by which the weights tensor expands in the `y`-direction.
//   - [IMPSGraphConvolution2DOpDescriptor.SetDilationRateInY]
//   - [IMPSGraphConvolution2DOpDescriptor.Groups]: The number of partitions of the input and output channels.
//   - [IMPSGraphConvolution2DOpDescriptor.SetGroups]
//   - [IMPSGraphConvolution2DOpDescriptor.PaddingBottom]: The number of zeros added at the bottom of the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetPaddingBottom]
//   - [IMPSGraphConvolution2DOpDescriptor.PaddingLeft]: The number of zeros added on the left side of the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetPaddingLeft]
//   - [IMPSGraphConvolution2DOpDescriptor.PaddingRight]: The number of zeros added on the right side of the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetPaddingRight]
//   - [IMPSGraphConvolution2DOpDescriptor.PaddingStyle]: The type of padding applied to the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphConvolution2DOpDescriptor.PaddingTop]: The number of zeros added at the top of the source tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetPaddingTop]
//   - [IMPSGraphConvolution2DOpDescriptor.StrideInX]: The scale that maps `x`-coordinate of the destination to `x`-coordinate of the source.
//   - [IMPSGraphConvolution2DOpDescriptor.SetStrideInX]
//   - [IMPSGraphConvolution2DOpDescriptor.StrideInY]: The scale that maps `y`-coordinate of the destination to `y`-coordinate of the source.
//   - [IMPSGraphConvolution2DOpDescriptor.SetStrideInY]
//   - [IMPSGraphConvolution2DOpDescriptor.WeightsLayout]: The named layout of data in the weights tensor.
//   - [IMPSGraphConvolution2DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [IMPSGraphConvolution2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the left, right, top, and bottom padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor
type IMPSGraphConvolution2DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The named layout of data in the source tensor.
	DataLayout() MPSGraphTensorNamedDataLayout
	SetDataLayout(value MPSGraphTensorNamedDataLayout)
	// The amount by which the weights tensor expands in the `x`-direction.
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	// The amount by which the weights tensor expands in the `y`-direction.
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	// The number of partitions of the input and output channels.
	Groups() uint
	SetGroups(value uint)
	// The number of zeros added at the bottom of the source tensor.
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	// The number of zeros added on the left side of the source tensor.
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	// The number of zeros added on the right side of the source tensor.
	PaddingRight() uint
	SetPaddingRight(value uint)
	// The type of padding applied to the source tensor.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// The number of zeros added at the top of the source tensor.
	PaddingTop() uint
	SetPaddingTop(value uint)
	// The scale that maps `x`-coordinate of the destination to `x`-coordinate of the source.
	StrideInX() uint
	SetStrideInX(value uint)
	// The scale that maps `y`-coordinate of the destination to `y`-coordinate of the source.
	StrideInY() uint
	SetStrideInY(value uint)
	// The named layout of data in the weights tensor.
	WeightsLayout() MPSGraphTensorNamedDataLayout
	SetWeightsLayout(value MPSGraphTensorNamedDataLayout)

	// Topic: Instance Methods

	// Sets the left, right, top, and bottom padding values.
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
}

// Init initializes the instance.
func (g MPSGraphConvolution2DOpDescriptor) Init() MPSGraphConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution2DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphConvolution2DOpDescriptor) Autorelease() MPSGraphConvolution2DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution2DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphConvolution2DOpDescriptor creates a new MPSGraphConvolution2DOpDescriptor instance.
func NewMPSGraphConvolution2DOpDescriptor() MPSGraphConvolution2DOpDescriptor {
	class := getMPSGraphConvolution2DOpDescriptorClass()
	rv := objc.Send[MPSGraphConvolution2DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a convolution descriptor with given values for parameters.
//
// strideInX: See [MPSGraphConvolution2DOpDescriptor.StrideInX] property.
//
// strideInY: See [MPSGraphConvolution2DOpDescriptor.StrideInY] property.
//
// dilationRateInX: See [MPSGraphConvolution2DOpDescriptor.DilationRateInX] property.
//
// dilationRateInY: See [MPSGraphConvolution2DOpDescriptor.DilationRateInY] property.
//
// groups: See [MPSGraphConvolution2DOpDescriptor.Groups] property.
//
// paddingLeft: See [MPSGraphConvolution2DOpDescriptor.PaddingLeft] property.
//
// paddingRight: See [MPSGraphConvolution2DOpDescriptor.PaddingRight] property.
//
// paddingTop: See [MPSGraphConvolution2DOpDescriptor.PaddingTop] property.
//
// paddingBottom: See [MPSGraphConvolution2DOpDescriptor.PaddingBottom] property.
//
// paddingStyle: See [MPSGraphConvolution2DOpDescriptor.PaddingStyle] property.
//
// dataLayout: See [MPSGraphConvolution2DOpDescriptor.DataLayout] property.
//
// weightsLayout: See [MPSGraphConvolution2DOpDescriptor.WeightsLayout] property.
//
// # Return Value
//
// The [MPSGraphConvolution2DOpDescriptor] on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphConvolution2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return MPSGraphConvolution2DOpDescriptorFromID(rv)
}

// Creates a convolution descriptor with given values for parameters.
//
// strideInX: See [MPSGraphConvolution2DOpDescriptor.StrideInX] property.
//
// strideInY: See [MPSGraphConvolution2DOpDescriptor.StrideInY] property.
//
// dilationRateInX: See [MPSGraphConvolution2DOpDescriptor.DilationRateInX] property.
//
// dilationRateInY: See [MPSGraphConvolution2DOpDescriptor.DilationRateInY] property.
//
// groups: See [MPSGraphConvolution2DOpDescriptor.Groups] property.
//
// paddingStyle: See [MPSGraphConvolution2DOpDescriptor.PaddingStyle] property.
//
// dataLayout: See [MPSGraphConvolution2DOpDescriptor.DataLayout] property.
//
// weightsLayout: See [MPSGraphConvolution2DOpDescriptor.WeightsLayout] property.
//
// # Return Value
//
// The [MPSGraphConvolution2DOpDescriptor] on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphConvolution2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingStyle, dataLayout, weightsLayout)
	return MPSGraphConvolution2DOpDescriptorFromID(rv)
}

// Sets the left, right, top, and bottom padding values.
//
// paddingLeft: See [MPSGraphConvolution2DOpDescriptor.PaddingLeft] property.
//
// paddingRight: See [MPSGraphConvolution2DOpDescriptor.PaddingRight] property.
//
// paddingTop: See [MPSGraphConvolution2DOpDescriptor.PaddingTop] property.
//
// paddingBottom: See [MPSGraphConvolution2DOpDescriptor.PaddingBottom] property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g MPSGraphConvolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

// The named layout of data in the source tensor.
//
// # Discussion
//
// It defines the order of named dimensions (Batch, Channel, Height, Width).
// The convolution operation uses this to interpret data in the source tensor.
// For example, if `dataLayout` is [MPSGraphTensorNamedDataLayoutNCHW],
// frameork interprets data in source tensor as `batch x channels x height x
// width` with `width` as fastest moving dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dataLayout
func (g MPSGraphConvolution2DOpDescriptor) DataLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("dataLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphConvolution2DOpDescriptor) SetDataLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataLayout:"), value)
}

// The amount by which the weights tensor expands in the `x`-direction.
//
// # Discussion
//
// The weights tensor is dilated by inserting `dilationRateInX-1` zeros
// between consecutive values in `x`-dimension. Dilated weights tensor width
// is `(dilationRateInX-1)*kernelWidth+1`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInX
func (g MPSGraphConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInX"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInX:"), value)
}

// The amount by which the weights tensor expands in the `y`-direction.
//
// # Discussion
//
// The weights tensor is dilated by inserting `dilationRateInY-1` zeros
// between consecutive values in `y`-dimension. Dilated weights tensor width
// is `(dilationRateInY-1)*kernelHeight+1`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInY
func (g MPSGraphConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInY"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInY:"), value)
}

// The number of partitions of the input and output channels.
//
// # Discussion
//
// The convolution operation divides input and output channels in `groups`
// partitions. input channels in a group or partition are only connected to
// output channels in corresponding group. Number of weights the convolution
// needs is `outputFeatureChannels x inputFeatureChannels/groups x kernelWidth
// x kernelHeight`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/groups
func (g MPSGraphConvolution2DOpDescriptor) Groups() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("groups"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetGroups(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setGroups:"), value)
}

// The number of zeros added at the bottom of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingBottom
func (g MPSGraphConvolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBottom"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBottom:"), value)
}

// The number of zeros added on the left side of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingLeft
func (g MPSGraphConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingLeft"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingLeft:"), value)
}

// The number of zeros added on the right side of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingRight
func (g MPSGraphConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingRight"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingRight:"), value)
}

// The type of padding applied to the source tensor.
//
// # Discussion
//
// If paddingStyle is [MPSGraphPaddingStyleExplicit], `paddingLeft`,
// `laddingRight`, `paddingTop`, and `paddingBottom` must to be specified. For
// all other padding styles, framework compute these values so you dont need
// to provide these values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingStyle
func (g MPSGraphConvolution2DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphConvolution2DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// The number of zeros added at the top of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingTop
func (g MPSGraphConvolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingTop"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingTop:"), value)
}

// The scale that maps `x`-coordinate of the destination to `x`-coordinate of
// the source.
//
// # Discussion
//
// Source `x`-coordinate, `sx` is computed from destination `x`-coordinate,
// `dx` as `sx = strideInX*dx`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInX
func (g MPSGraphConvolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInX"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInX:"), value)
}

// The scale that maps `y`-coordinate of the destination to `y`-coordinate of
// the source.
//
// # Discussion
//
// Source `y`-coordinate, `sy` is computed from destination `y`-coordinate,
// `dy` as `sy = strideInY*dy`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInY
func (g MPSGraphConvolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInY"))
	return rv
}
func (g MPSGraphConvolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInY:"), value)
}

// The named layout of data in the weights tensor.
//
// # Discussion
//
// It defines the order of named dimensions (Output channels, Input channels,
// Kernel height, Kernel width). The convolution operation uses this to
// interpret data in the weights tensor. For example, if `weightsLayout` is
// [MPSGraphTensorNamedDataLayoutOIHW], frameork interprets data in weights
// tensor as `outputChannels x inputChannels x kernelHeight x kernelWidth`
// with `kernelWidth` as fastest moving dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/weightsLayout
func (g MPSGraphConvolution2DOpDescriptor) WeightsLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("weightsLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphConvolution2DOpDescriptor) SetWeightsLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setWeightsLayout:"), value)
}
