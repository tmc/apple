// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphImToColOpDescriptor] class.
var (
	_MPSGraphImToColOpDescriptorClass     MPSGraphImToColOpDescriptorClass
	_MPSGraphImToColOpDescriptorClassOnce sync.Once
)

func getMPSGraphImToColOpDescriptorClass() MPSGraphImToColOpDescriptorClass {
	_MPSGraphImToColOpDescriptorClassOnce.Do(func() {
		_MPSGraphImToColOpDescriptorClass = MPSGraphImToColOpDescriptorClass{class: objc.GetClass("MPSGraphImToColOpDescriptor")}
	})
	return _MPSGraphImToColOpDescriptorClass
}

// GetMPSGraphImToColOpDescriptorClass returns the class object for MPSGraphImToColOpDescriptor.
func GetMPSGraphImToColOpDescriptorClass() MPSGraphImToColOpDescriptorClass {
	return getMPSGraphImToColOpDescriptorClass()
}

type MPSGraphImToColOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphImToColOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphImToColOpDescriptorClass) Alloc() MPSGraphImToColOpDescriptor {
	rv := objc.Send[MPSGraphImToColOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for an image to column or column to
// image operation.
//
// # Overview
//
// Use this descriptor with the following [MPSGraph] methods:
//
// - [MPSGraph.ImToColWithSourceTensorDescriptorName] -
// [MPSGraph.ColToImWithSourceTensorOutputShapeDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphImToColOpDescriptor.DataLayout]: The property that defines the layout of source or output  tensor. e.g. `batch x channels x width x height` for [NCHW] layout
//   - [MPSGraphImToColOpDescriptor.SetDataLayout]
//   - [MPSGraphImToColOpDescriptor.DilationRateInX]: The property that defines the dilation in width dimension.
//   - [MPSGraphImToColOpDescriptor.SetDilationRateInX]
//   - [MPSGraphImToColOpDescriptor.DilationRateInY]: The property that defines the dilation in height dimension.
//   - [MPSGraphImToColOpDescriptor.SetDilationRateInY]
//   - [MPSGraphImToColOpDescriptor.KernelHeight]: The property that defines the kernel size  in height dimension.
//   - [MPSGraphImToColOpDescriptor.SetKernelHeight]
//   - [MPSGraphImToColOpDescriptor.KernelWidth]: The property that defines the kernel size in width dimension.
//   - [MPSGraphImToColOpDescriptor.SetKernelWidth]
//   - [MPSGraphImToColOpDescriptor.PaddingBottom]: The property that defines the padding in height dimension at the bottom.
//   - [MPSGraphImToColOpDescriptor.SetPaddingBottom]
//   - [MPSGraphImToColOpDescriptor.PaddingLeft]: The property that defines the padding in width dimension on the left side.
//   - [MPSGraphImToColOpDescriptor.SetPaddingLeft]
//   - [MPSGraphImToColOpDescriptor.PaddingRight]: The property that defines the padding in width dimension on the right side.
//   - [MPSGraphImToColOpDescriptor.SetPaddingRight]
//   - [MPSGraphImToColOpDescriptor.PaddingTop]: The property that defines the padding in height dimension at the top.
//   - [MPSGraphImToColOpDescriptor.SetPaddingTop]
//   - [MPSGraphImToColOpDescriptor.StrideInX]: The property that defines the stride in width dimension.
//   - [MPSGraphImToColOpDescriptor.SetStrideInX]
//   - [MPSGraphImToColOpDescriptor.StrideInY]: The property that defines the stride in height dimension.
//   - [MPSGraphImToColOpDescriptor.SetStrideInY]
//
// # Instance Methods
//
//   - [MPSGraphImToColOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the descriptor’s padding to the given values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor
type MPSGraphImToColOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphImToColOpDescriptorFromID constructs a [MPSGraphImToColOpDescriptor] from an objc.ID.
//
// The class that defines the parameters for an image to column or column to
// image operation.
func MPSGraphImToColOpDescriptorFromID(id objc.ID) MPSGraphImToColOpDescriptor {
	return MPSGraphImToColOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphImToColOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphImToColOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphImToColOpDescriptor.DataLayout]: The property that defines the layout of source or output  tensor. e.g. `batch x channels x width x height` for [NCHW] layout
//   - [IMPSGraphImToColOpDescriptor.SetDataLayout]
//   - [IMPSGraphImToColOpDescriptor.DilationRateInX]: The property that defines the dilation in width dimension.
//   - [IMPSGraphImToColOpDescriptor.SetDilationRateInX]
//   - [IMPSGraphImToColOpDescriptor.DilationRateInY]: The property that defines the dilation in height dimension.
//   - [IMPSGraphImToColOpDescriptor.SetDilationRateInY]
//   - [IMPSGraphImToColOpDescriptor.KernelHeight]: The property that defines the kernel size  in height dimension.
//   - [IMPSGraphImToColOpDescriptor.SetKernelHeight]
//   - [IMPSGraphImToColOpDescriptor.KernelWidth]: The property that defines the kernel size in width dimension.
//   - [IMPSGraphImToColOpDescriptor.SetKernelWidth]
//   - [IMPSGraphImToColOpDescriptor.PaddingBottom]: The property that defines the padding in height dimension at the bottom.
//   - [IMPSGraphImToColOpDescriptor.SetPaddingBottom]
//   - [IMPSGraphImToColOpDescriptor.PaddingLeft]: The property that defines the padding in width dimension on the left side.
//   - [IMPSGraphImToColOpDescriptor.SetPaddingLeft]
//   - [IMPSGraphImToColOpDescriptor.PaddingRight]: The property that defines the padding in width dimension on the right side.
//   - [IMPSGraphImToColOpDescriptor.SetPaddingRight]
//   - [IMPSGraphImToColOpDescriptor.PaddingTop]: The property that defines the padding in height dimension at the top.
//   - [IMPSGraphImToColOpDescriptor.SetPaddingTop]
//   - [IMPSGraphImToColOpDescriptor.StrideInX]: The property that defines the stride in width dimension.
//   - [IMPSGraphImToColOpDescriptor.SetStrideInX]
//   - [IMPSGraphImToColOpDescriptor.StrideInY]: The property that defines the stride in height dimension.
//   - [IMPSGraphImToColOpDescriptor.SetStrideInY]
//
// # Instance Methods
//
//   - [IMPSGraphImToColOpDescriptor.SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom]: Sets the descriptor’s padding to the given values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor
type IMPSGraphImToColOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The property that defines the layout of source or output  tensor. e.g. `batch x channels x width x height` for [NCHW] layout
	DataLayout() MPSGraphTensorNamedDataLayout
	SetDataLayout(value MPSGraphTensorNamedDataLayout)
	// The property that defines the dilation in width dimension.
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	// The property that defines the dilation in height dimension.
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	// The property that defines the kernel size  in height dimension.
	KernelHeight() uint
	SetKernelHeight(value uint)
	// The property that defines the kernel size in width dimension.
	KernelWidth() uint
	SetKernelWidth(value uint)
	// The property that defines the padding in height dimension at the bottom.
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	// The property that defines the padding in width dimension on the left side.
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	// The property that defines the padding in width dimension on the right side.
	PaddingRight() uint
	SetPaddingRight(value uint)
	// The property that defines the padding in height dimension at the top.
	PaddingTop() uint
	SetPaddingTop(value uint)
	// The property that defines the stride in width dimension.
	StrideInX() uint
	SetStrideInX(value uint)
	// The property that defines the stride in height dimension.
	StrideInY() uint
	SetStrideInY(value uint)

	// Topic: Instance Methods

	// Sets the descriptor’s padding to the given values.
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
}

// Init initializes the instance.
func (g MPSGraphImToColOpDescriptor) Init() MPSGraphImToColOpDescriptor {
	rv := objc.Send[MPSGraphImToColOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphImToColOpDescriptor) Autorelease() MPSGraphImToColOpDescriptor {
	rv := objc.Send[MPSGraphImToColOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphImToColOpDescriptor creates a new MPSGraphImToColOpDescriptor instance.
func NewMPSGraphImToColOpDescriptor() MPSGraphImToColOpDescriptor {
	class := getMPSGraphImToColOpDescriptorClass()
	rv := objc.Send[MPSGraphImToColOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates column to image descriptor with given values for parameters.
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
// dataLayout: See `dataLayout` property.
//
// # Return Value
//
// A valid MPSGraphImToColOpDescriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, dataLayout MPSGraphTensorNamedDataLayout) MPSGraphImToColOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, dataLayout)
	return MPSGraphImToColOpDescriptorFromID(rv)
}

// Creates an image to column descriptor with given values for parameters.
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
// dataLayout: See `dataLayout` property.
//
// # Return Value
//
// A valid MPSGraphImToColOpDescriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, dataLayout MPSGraphTensorNamedDataLayout) MPSGraphImToColOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, dataLayout)
	return MPSGraphImToColOpDescriptorFromID(rv)
}

// Sets the descriptor’s padding to the given values.
//
// paddingLeft: See `paddingLeft` property.
//
// paddingRight: See `paddingRight` property.
//
// paddingTop: See `paddingTop` property.
//
// paddingBottom: See `paddingBottom` property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g MPSGraphImToColOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

// The property that defines the layout of source or output tensor. e.g.
// `batch x channels x width x height` for [NCHW] layout
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dataLayout
func (g MPSGraphImToColOpDescriptor) DataLayout() MPSGraphTensorNamedDataLayout {
	rv := objc.Send[MPSGraphTensorNamedDataLayout](g.ID, objc.Sel("dataLayout"))
	return MPSGraphTensorNamedDataLayout(rv)
}
func (g MPSGraphImToColOpDescriptor) SetDataLayout(value MPSGraphTensorNamedDataLayout) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataLayout:"), value)
}

// The property that defines the dilation in width dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g MPSGraphImToColOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInX"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInX:"), value)
}

// The property that defines the dilation in height dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInY
func (g MPSGraphImToColOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("dilationRateInY"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRateInY:"), value)
}

// The property that defines the kernel size in height dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelHeight
func (g MPSGraphImToColOpDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("kernelHeight"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetKernelHeight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setKernelHeight:"), value)
}

// The property that defines the kernel size in width dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelWidth
func (g MPSGraphImToColOpDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("kernelWidth"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetKernelWidth(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setKernelWidth:"), value)
}

// The property that defines the padding in height dimension at the bottom.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingBottom
func (g MPSGraphImToColOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingBottom"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingBottom:"), value)
}

// The property that defines the padding in width dimension on the left side.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingLeft
func (g MPSGraphImToColOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingLeft"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingLeft:"), value)
}

// The property that defines the padding in width dimension on the right side.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingRight
func (g MPSGraphImToColOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingRight"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingRight:"), value)
}

// The property that defines the padding in height dimension at the top.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingTop
func (g MPSGraphImToColOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("paddingTop"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingTop:"), value)
}

// The property that defines the stride in width dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInX
func (g MPSGraphImToColOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInX"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetStrideInX(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInX:"), value)
}

// The property that defines the stride in height dimension.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g MPSGraphImToColOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("strideInY"))
	return rv
}
func (g MPSGraphImToColOpDescriptor) SetStrideInY(value uint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrideInY:"), value)
}
