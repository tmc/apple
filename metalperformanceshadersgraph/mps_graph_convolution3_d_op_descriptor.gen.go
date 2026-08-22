// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphConvolution3DOpDescriptor] class.
var (
	_MPSGraphConvolution3DOpDescriptorClass     MPSGraphConvolution3DOpDescriptorClass
	_MPSGraphConvolution3DOpDescriptorClassOnce sync.Once
)

func getMPSGraphConvolution3DOpDescriptorClass() MPSGraphConvolution3DOpDescriptorClass {
	_MPSGraphConvolution3DOpDescriptorClassOnce.Do(func() {
		_MPSGraphConvolution3DOpDescriptorClass = MPSGraphConvolution3DOpDescriptorClass{class: objc.GetClass("MPSGraphConvolution3DOpDescriptor")}
	})
	return _MPSGraphConvolution3DOpDescriptorClass
}

// GetMPSGraphConvolution3DOpDescriptorClass returns the class object for MPSGraphConvolution3DOpDescriptor.
func GetMPSGraphConvolution3DOpDescriptorClass() MPSGraphConvolution3DOpDescriptorClass {
	return getMPSGraphConvolution3DOpDescriptorClass()
}

type MPSGraphConvolution3DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphConvolution3DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphConvolution3DOpDescriptorClass) Alloc() MPSGraphConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution3DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that describes the properties of a 3D-convolution operator.
//
// # Overview
//
// Use an instance of this class is to add a 3D-convolution operator with
// desired properties to the graph.
//
// # Instance Properties
//
//   - [MPSGraphConvolution3DOpDescriptor.DataLayout]: The named layout of data in the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetDataLayout]
//   - [MPSGraphConvolution3DOpDescriptor.DilationRateInX]: The amount by which weights tensor expands in the `x`-direction.
//   - [MPSGraphConvolution3DOpDescriptor.SetDilationRateInX]
//   - [MPSGraphConvolution3DOpDescriptor.DilationRateInY]: The amount by which weights tensor expands in the `y`-direction.
//   - [MPSGraphConvolution3DOpDescriptor.SetDilationRateInY]
//   - [MPSGraphConvolution3DOpDescriptor.DilationRateInZ]: The amount by which weights tensor expands in the `z`-direction.
//   - [MPSGraphConvolution3DOpDescriptor.SetDilationRateInZ]
//   - [MPSGraphConvolution3DOpDescriptor.Groups]: The number of partitions of the input and output channels.
//   - [MPSGraphConvolution3DOpDescriptor.SetGroups]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingBack]: The number of zeros added at the back of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingBack]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingBottom]: The number of zeros added at the bottom of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingBottom]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingFront]: The number of zeros added at the front of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingFront]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingLeft]: The number of zeros added on the left side of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingLeft]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingRight]: The number of zeros added on the right side of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingRight]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingStyle]: The type of padding that is applied to the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphConvolution3DOpDescriptor.PaddingTop]: The number of zeros added at the top of the source tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetPaddingTop]
//   - [MPSGraphConvolution3DOpDescriptor.StrideInX]: The scale that maps`x`-coordinate of destination to `x`-coordinate of source.
//   - [MPSGraphConvolution3DOpDescriptor.SetStrideInX]
//   - [MPSGraphConvolution3DOpDescriptor.StrideInY]: The scale that maps`y`-coordinate of destination to `y`-coordinate of source.
//   - [MPSGraphConvolution3DOpDescriptor.SetStrideInY]
//   - [MPSGraphConvolution3DOpDescriptor.StrideInZ]: The scale that maps`z`-coordinate of destination to `z`-coordinate of source.
//   - [MPSGraphConvolution3DOpDescriptor.SetStrideInZ]
//   - [MPSGraphConvolution3DOpDescriptor.WeightsLayout]: The named layout of data in the weights tensor.
//   - [MPSGraphConvolution3DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [MPSGraphConvolution3DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack]: Sets the left, right, top, bottom, front, and back padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor
type MPSGraphConvolution3DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphConvolution3DOpDescriptorFromID constructs a [MPSGraphConvolution3DOpDescriptor] from an objc.ID.
//
// A class that describes the properties of a 3D-convolution operator.
func MPSGraphConvolution3DOpDescriptorFromID(id objc.ID) MPSGraphConvolution3DOpDescriptor {
	return MPSGraphConvolution3DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphConvolution3DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphConvolution3DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphConvolution3DOpDescriptor.DataLayout]: The named layout of data in the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetDataLayout]
//   - [IMPSGraphConvolution3DOpDescriptor.DilationRateInX]: The amount by which weights tensor expands in the `x`-direction.
//   - [IMPSGraphConvolution3DOpDescriptor.SetDilationRateInX]
//   - [IMPSGraphConvolution3DOpDescriptor.DilationRateInY]: The amount by which weights tensor expands in the `y`-direction.
//   - [IMPSGraphConvolution3DOpDescriptor.SetDilationRateInY]
//   - [IMPSGraphConvolution3DOpDescriptor.DilationRateInZ]: The amount by which weights tensor expands in the `z`-direction.
//   - [IMPSGraphConvolution3DOpDescriptor.SetDilationRateInZ]
//   - [IMPSGraphConvolution3DOpDescriptor.Groups]: The number of partitions of the input and output channels.
//   - [IMPSGraphConvolution3DOpDescriptor.SetGroups]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingBack]: The number of zeros added at the back of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingBack]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingBottom]: The number of zeros added at the bottom of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingBottom]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingFront]: The number of zeros added at the front of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingFront]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingLeft]: The number of zeros added on the left side of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingLeft]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingRight]: The number of zeros added on the right side of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingRight]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingStyle]: The type of padding that is applied to the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphConvolution3DOpDescriptor.PaddingTop]: The number of zeros added at the top of the source tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetPaddingTop]
//   - [IMPSGraphConvolution3DOpDescriptor.StrideInX]: The scale that maps`x`-coordinate of destination to `x`-coordinate of source.
//   - [IMPSGraphConvolution3DOpDescriptor.SetStrideInX]
//   - [IMPSGraphConvolution3DOpDescriptor.StrideInY]: The scale that maps`y`-coordinate of destination to `y`-coordinate of source.
//   - [IMPSGraphConvolution3DOpDescriptor.SetStrideInY]
//   - [IMPSGraphConvolution3DOpDescriptor.StrideInZ]: The scale that maps`z`-coordinate of destination to `z`-coordinate of source.
//   - [IMPSGraphConvolution3DOpDescriptor.SetStrideInZ]
//   - [IMPSGraphConvolution3DOpDescriptor.WeightsLayout]: The named layout of data in the weights tensor.
//   - [IMPSGraphConvolution3DOpDescriptor.SetWeightsLayout]
//
// # Instance Methods
//
//   - [IMPSGraphConvolution3DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack]: Sets the left, right, top, bottom, front, and back padding values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor
type IMPSGraphConvolution3DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The named layout of data in the source tensor.
	DataLayout() MPSGraphTensorNamedDataLayout
	SetDataLayout(value MPSGraphTensorNamedDataLayout)
	// The amount by which weights tensor expands in the `x`-direction.
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	// The amount by which weights tensor expands in the `y`-direction.
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	// The amount by which weights tensor expands in the `z`-direction.
	DilationRateInZ() uint
	SetDilationRateInZ(value uint)
	// The number of partitions of the input and output channels.
	Groups() uint
	SetGroups(value uint)
	// The number of zeros added at the back of the source tensor.
	PaddingBack() uint
	SetPaddingBack(value uint)
	// The number of zeros added at the bottom of the source tensor.
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	// The number of zeros added at the front of the source tensor.
	PaddingFront() uint
	SetPaddingFront(value uint)
	// The number of zeros added on the left side of the source tensor.
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	// The number of zeros added on the right side of the source tensor.
	PaddingRight() uint
	SetPaddingRight(value uint)
	// The type of padding that is applied to the source tensor.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// The number of zeros added at the top of the source tensor.
	PaddingTop() uint
	SetPaddingTop(value uint)
	// The scale that maps`x`-coordinate of destination to `x`-coordinate of source.
	StrideInX() uint
	SetStrideInX(value uint)
	// The scale that maps`y`-coordinate of destination to `y`-coordinate of source.
	StrideInY() uint
	SetStrideInY(value uint)
	// The scale that maps`z`-coordinate of destination to `z`-coordinate of source.
	StrideInZ() uint
	SetStrideInZ(value uint)
	// The named layout of data in the weights tensor.
	WeightsLayout() MPSGraphTensorNamedDataLayout
	SetWeightsLayout(value MPSGraphTensorNamedDataLayout)

	// Topic: Instance Methods

	// Sets the left, right, top, bottom, front, and back padding values.
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint)
}

// Init initializes the instance.
func (g MPSGraphConvolution3DOpDescriptor) Init() MPSGraphConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution3DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphConvolution3DOpDescriptor) Autorelease() MPSGraphConvolution3DOpDescriptor {
	rv := objc.Send[MPSGraphConvolution3DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphConvolution3DOpDescriptor creates a new MPSGraphConvolution3DOpDescriptor instance.
func NewMPSGraphConvolution3DOpDescriptor() MPSGraphConvolution3DOpDescriptor {
	class := getMPSGraphConvolution3DOpDescriptorClass()
	rv := objc.Send[MPSGraphConvolution3DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a convolution descriptor with given values for parameters.
//
// strideInX: See [MPSGraphConvolution3DOpDescriptor.StrideInX] property.
//
// strideInY: See [MPSGraphConvolution3DOpDescriptor.StrideInY] property.
//
// strideInZ: See [MPSGraphConvolution3DOpDescriptor.StrideInZ] property.
//
// dilationRateInX: See [MPSGraphConvolution3DOpDescriptor.DilationRateInX] property.
//
// dilationRateInY: See [MPSGraphConvolution3DOpDescriptor.DilationRateInY] property.
//
// dilationRateInZ: See [MPSGraphConvolution3DOpDescriptor.DilationRateInZ] property.
//
// groups: See [MPSGraphConvolution3DOpDescriptor.Groups] property.
//
// paddingLeft: See [MPSGraphConvolution3DOpDescriptor.PaddingLeft] property.
//
// paddingRight: See [MPSGraphConvolution3DOpDescriptor.PaddingRight] property.
//
// paddingTop: See [MPSGraphConvolution3DOpDescriptor.PaddingTop] property.
//
// paddingBottom: See [MPSGraphConvolution3DOpDescriptor.PaddingBottom] property.
//
// paddingFront: See [MPSGraphConvolution3DOpDescriptor.PaddingFront] property.
//
// paddingBack: See [MPSGraphConvolution3DOpDescriptor.PaddingBack] property.
//
// paddingStyle: See [MPSGraphConvolution3DOpDescriptor.PaddingStyle] property.
//
// dataLayout: See [MPSGraphConvolution3DOpDescriptor.DataLayout] property.
//
// weightsLayout: See [MPSGraphConvolution3DOpDescriptor.WeightsLayout] property.
//
// # Return Value
//
// The [MPSGraphConvolution3DOpDescriptor] on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphConvolution3DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack, paddingStyle, dataLayout, weightsLayout)
	return MPSGraphConvolution3DOpDescriptorFromID(rv)
}

// Creates a convolution descriptor with given values for parameters.
//
// strideInX: See [MPSGraphConvolution3DOpDescriptor.StrideInX] property.
//
// strideInY: See [MPSGraphConvolution3DOpDescriptor.StrideInY] property.
//
// strideInZ: See [MPSGraphConvolution3DOpDescriptor.StrideInZ] property.
//
// dilationRateInX: See [MPSGraphConvolution3DOpDescriptor.DilationRateInX] property.
//
// dilationRateInY: See [MPSGraphConvolution3DOpDescriptor.DilationRateInY] property.
//
// dilationRateInZ: See [MPSGraphConvolution3DOpDescriptor.DilationRateInZ] property.
//
// groups: See [MPSGraphConvolution3DOpDescriptor.Groups] property.
//
// paddingStyle: See [MPSGraphConvolution3DOpDescriptor.PaddingStyle] property.
//
// dataLayout: See [MPSGraphConvolution3DOpDescriptor.DataLayout] property.
//
// weightsLayout: See [MPSGraphConvolution3DOpDescriptor.WeightsLayout] property.
//
// # Return Value
//
// The [MPSGraphConvolution3DOpDescriptor] on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout, weightsLayout MPSGraphTensorNamedDataLayout) MPSGraphConvolution3DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingStyle, dataLayout, weightsLayout)
	return MPSGraphConvolution3DOpDescriptorFromID(rv)
}

// Sets the left, right, top, bottom, front, and back padding values.
//
// paddingLeft: See [MPSGraphConvolution3DOpDescriptor.PaddingLeft] property.
//
// paddingRight: See [MPSGraphConvolution3DOpDescriptor.PaddingRight] property.
//
// paddingTop: See [MPSGraphConvolution3DOpDescriptor.PaddingTop] property.
//
// paddingBottom: See [MPSGraphConvolution3DOpDescriptor.PaddingBottom] property.
//
// paddingFront: See [MPSGraphConvolution3DOpDescriptor.PaddingFront] property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:)
func (g MPSGraphConvolution3DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint) {
	objc.Send[objc.ID](g.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:"), paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack)
}

// The named layout of data in the source tensor.
//
// # Discussion
//
// It defines the order of named dimensions (Batch, Channel, Depth, Height,
// Width). The convolution operation uses this to interpret data in the source
// tensor. For example, if `dataLayout` is
// [MPSGraphTensorNamedDataLayoutNCDHW], frameork interprets data in source
// tensor as `batch x channels x depth x height x width` with `width` as
// fastest moving dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dataLayout
func (g MPSGraphConvolution3DOpDescriptor) DataLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("dataLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphConvolution3DOpDescriptor) SetDataLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataLayout:"), value)
}

// The amount by which weights tensor expands in the `x`-direction.
//
// # Discussion
//
// The weights tensor is dilated by inserting `dilationRateInX-1` zeros
// between consecutive values in `x`-dimension. Dilated weights tensor width
// is `(dilationRateInX-1)*kernelWidth+1`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInX
func (g MPSGraphConvolution3DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInX"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInX:"), value)
}

// The amount by which weights tensor expands in the `y`-direction.
//
// # Discussion
//
// The weights tensor is dilated by inserting `dilationRateInY-1` zeros
// between consecutive values in `y`-dimension. Dilated weights tensor width
// is `(dilationRateInY-1)*kernelHeight+1`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInY
func (g MPSGraphConvolution3DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInY"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInY:"), value)
}

// The amount by which weights tensor expands in the `z`-direction.
//
// # Discussion
//
// The weights tensor is dilated by inserting `dilationRateInZ-1` zeros
// between consecutive values in `z`-dimension. Dilated weights tensor depth
// is `(dilationRateInZ-1)*kernelDepth+1`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInZ
func (g MPSGraphConvolution3DOpDescriptor) DilationRateInZ() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInZ"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetDilationRateInZ(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInZ:"), value)
}

// The number of partitions of the input and output channels.
//
// # Discussion
//
// The convolution operation divides input and output channels in `groups`
// partitions. input channels in a group or partition are only connected to
// output channels in corresponding group. Number of weights the convolution
// needs is `outputFeatureChannels x inputFeatureChannels/groups x kernelDepth
// x kernelWidth x kernelHeight`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/groups
func (g MPSGraphConvolution3DOpDescriptor) Groups() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("groups"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetGroups(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setGroups:"), value)
}

// The number of zeros added at the back of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBack
func (g MPSGraphConvolution3DOpDescriptor) PaddingBack() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBack"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingBack(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBack:"), value)
}

// The number of zeros added at the bottom of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBottom
func (g MPSGraphConvolution3DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBottom"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBottom:"), value)
}

// The number of zeros added at the front of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingFront
func (g MPSGraphConvolution3DOpDescriptor) PaddingFront() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingFront"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingFront(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingFront:"), value)
}

// The number of zeros added on the left side of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingLeft
func (g MPSGraphConvolution3DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingLeft"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingLeft:"), value)
}

// The number of zeros added on the right side of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingRight
func (g MPSGraphConvolution3DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingRight"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingRight:"), value)
}

// The type of padding that is applied to the source tensor.
//
// # Discussion
//
// If paddingStyle is [MPSGraphPaddingStyleExplicit], `paddingLeft`,
// `laddingRight`, `paddingTop`, `paddingBottom`, `paddingFront` and
// `paddingBack` must to be specified. For all other padding styles, framework
// compute these values so you dont need to provide these values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingStyle
func (g MPSGraphConvolution3DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// The number of zeros added at the top of the source tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingTop
func (g MPSGraphConvolution3DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingTop"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingTop:"), value)
}

// The scale that maps`x`-coordinate of destination to `x`-coordinate of
// source.
//
// # Discussion
//
// Source `x`-coordinate, `sx` is computed from destination `x`-coordinate,
// `dx` as `sx = strideInX*dx`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInX
func (g MPSGraphConvolution3DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInX"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInX:"), value)
}

// The scale that maps`y`-coordinate of destination to `y`-coordinate of
// source.
//
// # Discussion
//
// Source `y`-coordinate, `sy` is computed from destination `y`-coordinate,
// `dy` as `sy = strideInY*dy`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInY
func (g MPSGraphConvolution3DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInY"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInY:"), value)
}

// The scale that maps`z`-coordinate of destination to `z`-coordinate of
// source.
//
// # Discussion
//
// Source `z`-coordinate, `sz` is computed from destination `z`-coordinate,
// `dz` as `sz = strideInZ*dz`. Default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInZ
func (g MPSGraphConvolution3DOpDescriptor) StrideInZ() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInZ"))
	return rv
}
func (g MPSGraphConvolution3DOpDescriptor) SetStrideInZ(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInZ:"), value)
}

// The named layout of data in the weights tensor.
//
// # Discussion
//
// It defines the order of named dimensions (Output channels, Input channels,
// Kernel depth, Kernel height, Kernel width). The convolution operation uses
// this to interpret data in the weights tensor. For example, if
// `weightsLayout` is [MPSGraphTensorNamedDataLayoutOIDHW], frameork
// interprets data in weights tensor as `outputChannels x inputChannels x
// kernelDepth x kernelHeight x kernelWidth` with `kernelWidth` as fastest
// moving dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/weightsLayout
func (g MPSGraphConvolution3DOpDescriptor) WeightsLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("weightsLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphConvolution3DOpDescriptor) SetWeightsLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setWeightsLayout:"), value)
}
