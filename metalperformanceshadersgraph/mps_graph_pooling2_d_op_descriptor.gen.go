// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphPooling2DOpDescriptor] class.
var (
	_MPSGraphPooling2DOpDescriptorClass     MPSGraphPooling2DOpDescriptorClass
	_MPSGraphPooling2DOpDescriptorClassOnce sync.Once
)

func getMPSGraphPooling2DOpDescriptorClass() MPSGraphPooling2DOpDescriptorClass {
	_MPSGraphPooling2DOpDescriptorClassOnce.Do(func() {
		_MPSGraphPooling2DOpDescriptorClass = MPSGraphPooling2DOpDescriptorClass{class: objc.GetClass("MPSGraphPooling2DOpDescriptor")}
	})
	return _MPSGraphPooling2DOpDescriptorClass
}

// GetMPSGraphPooling2DOpDescriptorClass returns the class object for MPSGraphPooling2DOpDescriptor.
func GetMPSGraphPooling2DOpDescriptorClass() MPSGraphPooling2DOpDescriptorClass {
	return getMPSGraphPooling2DOpDescriptorClass()
}

type MPSGraphPooling2DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphPooling2DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphPooling2DOpDescriptorClass) Alloc() MPSGraphPooling2DOpDescriptor {
	rv := objc.Send[MPSGraphPooling2DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a 2D pooling operation.
//
// # Overview
//
// Use this descriptor with the following methods:
//
// - [MPSGraph.MaxPooling2DWithSourceTensorDescriptorName] -
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName] -
// [MPSGraph.MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName]
// -
// [MPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]
// -
// [MPSGraph.MaxPooling2DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]
// - [MPSGraph.AvgPooling2DWithSourceTensorDescriptorName] -
// [MPSGraph.AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphPooling2DOpDescriptor.CeilMode]: Affects how the graph computes the output size.
//   - [MPSGraphPooling2DOpDescriptor.SetCeilMode]
//   - [MPSGraphPooling2DOpDescriptor.DataLayout]: Defines the data layout of the input data in the forward pass. See: [MPSGraphTensorNamedDataLayout](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout>).
//   - [MPSGraphPooling2DOpDescriptor.SetDataLayout]
//   - [MPSGraphPooling2DOpDescriptor.DilationRateInX]: Defines the dilation rate for the width dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetDilationRateInX]
//   - [MPSGraphPooling2DOpDescriptor.DilationRateInY]: Defines the dilation rate for the height dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetDilationRateInY]
//   - [MPSGraphPooling2DOpDescriptor.IncludeZeroPadToAverage]: Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//   - [MPSGraphPooling2DOpDescriptor.SetIncludeZeroPadToAverage]
//   - [MPSGraphPooling2DOpDescriptor.KernelHeight]: Defines the pooling window size for the height dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetKernelHeight]
//   - [MPSGraphPooling2DOpDescriptor.KernelWidth]: Defines the pooling window size for the width dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetKernelWidth]
//   - [MPSGraphPooling2DOpDescriptor.PaddingBottom]: Defines the explicit padding value for the height dimension to add after the data.
//   - [MPSGraphPooling2DOpDescriptor.SetPaddingBottom]
//   - [MPSGraphPooling2DOpDescriptor.PaddingLeft]: Defines the explicit padding value for the width dimension to add before the data.
//   - [MPSGraphPooling2DOpDescriptor.SetPaddingLeft]
//   - [MPSGraphPooling2DOpDescriptor.PaddingRight]: Defines the explicit padding value for the width dimension to add after the data.
//   - [MPSGraphPooling2DOpDescriptor.SetPaddingRight]
//   - [MPSGraphPooling2DOpDescriptor.PaddingStyle]: Defines what kind of padding graph applies to the operation.
//   - [MPSGraphPooling2DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphPooling2DOpDescriptor.PaddingTop]: Defines the explicit padding value for the height dimension to add before the data.
//   - [MPSGraphPooling2DOpDescriptor.SetPaddingTop]
//   - [MPSGraphPooling2DOpDescriptor.ReturnIndicesDataType]: Defines the data type for returned indices. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. Currently MPSGraph supports the following datatypes: [MPSDataTypeInt32]. Default value: [MPSDataTypeInt32].
//   - [MPSGraphPooling2DOpDescriptor.SetReturnIndicesDataType]
//   - [MPSGraphPooling2DOpDescriptor.ReturnIndicesMode]: Defines the mode for returned indices of maximum values within each pooling window. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. If `returnIndicesMode = MPSGraphPoolingReturnIndicesNone` then only the first result MPSGraph returns from [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) will be valid and using the second result will assert. Default value: [MPSGraphPoolingReturnIndicesNone].
//   - [MPSGraphPooling2DOpDescriptor.SetReturnIndicesMode]
//   - [MPSGraphPooling2DOpDescriptor.StrideInX]: Defines the stride for the width dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetStrideInX]
//   - [MPSGraphPooling2DOpDescriptor.StrideInY]: Defines the stride for the height dimension.
//   - [MPSGraphPooling2DOpDescriptor.SetStrideInY]
//
// # Instance Methods
//
//   - [MPSGraphPooling2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the explicit padding values and sets padding style to explicit.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor
type MPSGraphPooling2DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphPooling2DOpDescriptorFromID constructs a [MPSGraphPooling2DOpDescriptor] from an objc.ID.
//
// The class that defines the parameters for a 2D pooling operation.
func MPSGraphPooling2DOpDescriptorFromID(id objc.ID) MPSGraphPooling2DOpDescriptor {
	return MPSGraphPooling2DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphPooling2DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphPooling2DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphPooling2DOpDescriptor.CeilMode]: Affects how the graph computes the output size.
//   - [IMPSGraphPooling2DOpDescriptor.SetCeilMode]
//   - [IMPSGraphPooling2DOpDescriptor.DataLayout]: Defines the data layout of the input data in the forward pass. See: [MPSGraphTensorNamedDataLayout](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout>).
//   - [IMPSGraphPooling2DOpDescriptor.SetDataLayout]
//   - [IMPSGraphPooling2DOpDescriptor.DilationRateInX]: Defines the dilation rate for the width dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetDilationRateInX]
//   - [IMPSGraphPooling2DOpDescriptor.DilationRateInY]: Defines the dilation rate for the height dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetDilationRateInY]
//   - [IMPSGraphPooling2DOpDescriptor.IncludeZeroPadToAverage]: Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//   - [IMPSGraphPooling2DOpDescriptor.SetIncludeZeroPadToAverage]
//   - [IMPSGraphPooling2DOpDescriptor.KernelHeight]: Defines the pooling window size for the height dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetKernelHeight]
//   - [IMPSGraphPooling2DOpDescriptor.KernelWidth]: Defines the pooling window size for the width dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetKernelWidth]
//   - [IMPSGraphPooling2DOpDescriptor.PaddingBottom]: Defines the explicit padding value for the height dimension to add after the data.
//   - [IMPSGraphPooling2DOpDescriptor.SetPaddingBottom]
//   - [IMPSGraphPooling2DOpDescriptor.PaddingLeft]: Defines the explicit padding value for the width dimension to add before the data.
//   - [IMPSGraphPooling2DOpDescriptor.SetPaddingLeft]
//   - [IMPSGraphPooling2DOpDescriptor.PaddingRight]: Defines the explicit padding value for the width dimension to add after the data.
//   - [IMPSGraphPooling2DOpDescriptor.SetPaddingRight]
//   - [IMPSGraphPooling2DOpDescriptor.PaddingStyle]: Defines what kind of padding graph applies to the operation.
//   - [IMPSGraphPooling2DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphPooling2DOpDescriptor.PaddingTop]: Defines the explicit padding value for the height dimension to add before the data.
//   - [IMPSGraphPooling2DOpDescriptor.SetPaddingTop]
//   - [IMPSGraphPooling2DOpDescriptor.ReturnIndicesDataType]: Defines the data type for returned indices. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. Currently MPSGraph supports the following datatypes: [MPSDataTypeInt32]. Default value: [MPSDataTypeInt32].
//   - [IMPSGraphPooling2DOpDescriptor.SetReturnIndicesDataType]
//   - [IMPSGraphPooling2DOpDescriptor.ReturnIndicesMode]: Defines the mode for returned indices of maximum values within each pooling window. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. If `returnIndicesMode = MPSGraphPoolingReturnIndicesNone` then only the first result MPSGraph returns from [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) will be valid and using the second result will assert. Default value: [MPSGraphPoolingReturnIndicesNone].
//   - [IMPSGraphPooling2DOpDescriptor.SetReturnIndicesMode]
//   - [IMPSGraphPooling2DOpDescriptor.StrideInX]: Defines the stride for the width dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetStrideInX]
//   - [IMPSGraphPooling2DOpDescriptor.StrideInY]: Defines the stride for the height dimension.
//   - [IMPSGraphPooling2DOpDescriptor.SetStrideInY]
//
// # Instance Methods
//
//   - [IMPSGraphPooling2DOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the explicit padding values and sets padding style to explicit.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor
type IMPSGraphPooling2DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// Affects how the graph computes the output size.
	CeilMode() bool
	SetCeilMode(value bool)
	// Defines the data layout of the input data in the forward pass. See: [MPSGraphTensorNamedDataLayout](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout>).
	DataLayout() MPSGraphTensorNamedDataLayout
	SetDataLayout(value MPSGraphTensorNamedDataLayout)
	// Defines the dilation rate for the width dimension.
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	// Defines the dilation rate for the height dimension.
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	// Defines the pooling window size for the height dimension.
	KernelHeight() uint
	SetKernelHeight(value uint)
	// Defines the pooling window size for the width dimension.
	KernelWidth() uint
	SetKernelWidth(value uint)
	// Defines the explicit padding value for the height dimension to add after the data.
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	// Defines the explicit padding value for the width dimension to add before the data.
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	// Defines the explicit padding value for the width dimension to add after the data.
	PaddingRight() uint
	SetPaddingRight(value uint)
	// Defines what kind of padding graph applies to the operation.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// Defines the explicit padding value for the height dimension to add before the data.
	PaddingTop() uint
	SetPaddingTop(value uint)
	// Defines the data type for returned indices. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. Currently MPSGraph supports the following datatypes: [MPSDataTypeInt32]. Default value: [MPSDataTypeInt32].
	ReturnIndicesDataType() uint32
	SetReturnIndicesDataType(value uint32)
	// Defines the mode for returned indices of maximum values within each pooling window. Use this in conjunction with [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) API. If `returnIndicesMode = MPSGraphPoolingReturnIndicesNone` then only the first result MPSGraph returns from [maxPooling2DReturnIndices(_:descriptor:name:)](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)>) will be valid and using the second result will assert. Default value: [MPSGraphPoolingReturnIndicesNone].
	ReturnIndicesMode() MPSGraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value MPSGraphPoolingReturnIndicesMode)
	// Defines the stride for the width dimension.
	StrideInX() uint
	SetStrideInX(value uint)
	// Defines the stride for the height dimension.
	StrideInY() uint
	SetStrideInY(value uint)

	// Topic: Instance Methods

	// Sets the explicit padding values and sets padding style to explicit.
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
}

// Init initializes the instance.
func (g MPSGraphPooling2DOpDescriptor) Init() MPSGraphPooling2DOpDescriptor {
	rv := objc.Send[MPSGraphPooling2DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphPooling2DOpDescriptor) Autorelease() MPSGraphPooling2DOpDescriptor {
	rv := objc.Send[MPSGraphPooling2DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphPooling2DOpDescriptor creates a new MPSGraphPooling2DOpDescriptor instance.
func NewMPSGraphPooling2DOpDescriptor() MPSGraphPooling2DOpDescriptor {
	class := getMPSGraphPooling2DOpDescriptorClass()
	rv := objc.Send[MPSGraphPooling2DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a 2D pooling descriptor with given values.
//
// kernelWidth: See `kernelWidth` property.
//
// kernelHeight: See `kernelHeight` property.
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
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:)
func NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout) MPSGraphPooling2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphPooling2DOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout)
	return MPSGraphPooling2DOpDescriptorFromID(rv)
}

// Creates a 2D pooling descriptor with given values.
//
// kernelWidth: See `kernelWidth` property.
//
// kernelHeight: See `kernelHeight“ property.
//
// strideInX: See `strideInX` property.
//
// strideInY: See `strideInY` property.
//
// paddingStyle: See `paddingStyle` property.
//
// dataLayout: See `dataLayout` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:)
func NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, paddingStyle MPSGraphPaddingStyle, dataLayout MPSGraphTensorNamedDataLayout) MPSGraphPooling2DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphPooling2DOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, paddingStyle, dataLayout)
	return MPSGraphPooling2DOpDescriptorFromID(rv)
}

// Sets the explicit padding values and sets padding style to explicit.
//
// paddingLeft: See `paddingLeft` property.
//
// paddingRight: See `paddingRight` property.
//
// paddingTop: See `paddingTop` property.
//
// paddingBottom: See `paddingBottom` property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g MPSGraphPooling2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

// Affects how the graph computes the output size.
//
// # Discussion
//
// If set to [YES] then output size is computed by rounding up instead of down
// when dividing input size by stride. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/ceilMode
func (g MPSGraphPooling2DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("ceilMode"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setCeilMode:"), value)
}

// Defines the data layout of the input data in the forward pass. See:
// [MPSGraphTensorNamedDataLayout].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dataLayout
//
// [MPSGraphTensorNamedDataLayout]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorNamedDataLayout
func (g MPSGraphPooling2DOpDescriptor) DataLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("dataLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphPooling2DOpDescriptor) SetDataLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataLayout:"), value)
}

// Defines the dilation rate for the width dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInX
func (g MPSGraphPooling2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInX"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInX:"), value)
}

// Defines the dilation rate for the height dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInY
func (g MPSGraphPooling2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInY"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInY:"), value)
}

// Defines a mode for average pooling, where samples outside the input tensor
// count as zeroes in the average computation.
//
// # Discussion
//
// Otherwise the result is sum over samples divided by number of samples that
// didn’t come from padding. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/includeZeroPadToAverage
func (g MPSGraphPooling2DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

// Defines the pooling window size for the height dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelHeight
func (g MPSGraphPooling2DOpDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("kernelHeight"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetKernelHeight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setKernelHeight:"), value)
}

// Defines the pooling window size for the width dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelWidth
func (g MPSGraphPooling2DOpDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("kernelWidth"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetKernelWidth(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setKernelWidth:"), value)
}

// Defines the explicit padding value for the height dimension to add after
// the data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingBottom
func (g MPSGraphPooling2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBottom"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBottom:"), value)
}

// Defines the explicit padding value for the width dimension to add before
// the data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g MPSGraphPooling2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingLeft"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingLeft:"), value)
}

// Defines the explicit padding value for the width dimension to add after the
// data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingRight
func (g MPSGraphPooling2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingRight"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingRight:"), value)
}

// Defines what kind of padding graph applies to the operation.
//
// # Discussion
//
// Default value: [MPSGraphPaddingStyleExplicit].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingStyle
func (g MPSGraphPooling2DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphPooling2DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// Defines the explicit padding value for the height dimension to add before
// the data.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingTop
func (g MPSGraphPooling2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingTop"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingTop:"), value)
}

// Defines the data type for returned indices. Use this in conjunction with
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName] API.
// Currently MPSGraph supports the following datatypes: [MPSDataTypeInt32].
// Default value: [MPSDataTypeInt32].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesDataType
func (g MPSGraphPooling2DOpDescriptor) ReturnIndicesDataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("returnIndicesDataType"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetReturnIndicesDataType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setReturnIndicesDataType:"), value)
}

// Defines the mode for returned indices of maximum values within each pooling
// window. Use this in conjunction with
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName] API. If
// `returnIndicesMode = MPSGraphPoolingReturnIndicesNone` then only the first
// result MPSGraph returns from
// [MPSGraph.MaxPooling2DReturnIndicesWithSourceTensorDescriptorName] will be
// valid and using the second result will assert. Default value:
// [MPSGraphPoolingReturnIndicesNone].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesMode
func (g MPSGraphPooling2DOpDescriptor) ReturnIndicesMode() MPSGraphPoolingReturnIndicesMode {
	rv := objc.Send[MPSGraphPoolingReturnIndicesMode](g.ID, objc.Sel("returnIndicesMode"))
	return MPSGraphPoolingReturnIndicesMode(rv)
}
func (g MPSGraphPooling2DOpDescriptor) SetReturnIndicesMode(value MPSGraphPoolingReturnIndicesMode) {
	objc.Send[struct{}](g.ID, objc.Sel("setReturnIndicesMode:"), value)
}

// Defines the stride for the width dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInX
func (g MPSGraphPooling2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInX"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInX:"), value)
}

// Defines the stride for the height dimension.
//
// # Discussion
//
// Default value: 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInY
func (g MPSGraphPooling2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInY"))
	return rv
}
func (g MPSGraphPooling2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInY:"), value)
}
