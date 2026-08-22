// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraphPooling4DOpDescriptor] class.
var (
	_MPSGraphPooling4DOpDescriptorClass     MPSGraphPooling4DOpDescriptorClass
	_MPSGraphPooling4DOpDescriptorClassOnce sync.Once
)

func getMPSGraphPooling4DOpDescriptorClass() MPSGraphPooling4DOpDescriptorClass {
	_MPSGraphPooling4DOpDescriptorClassOnce.Do(func() {
		_MPSGraphPooling4DOpDescriptorClass = MPSGraphPooling4DOpDescriptorClass{class: objc.GetClass("MPSGraphPooling4DOpDescriptor")}
	})
	return _MPSGraphPooling4DOpDescriptorClass
}

// GetMPSGraphPooling4DOpDescriptorClass returns the class object for MPSGraphPooling4DOpDescriptor.
func GetMPSGraphPooling4DOpDescriptorClass() MPSGraphPooling4DOpDescriptorClass {
	return getMPSGraphPooling4DOpDescriptorClass()
}

type MPSGraphPooling4DOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphPooling4DOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphPooling4DOpDescriptorClass) Alloc() MPSGraphPooling4DOpDescriptor {
	rv := objc.Send[MPSGraphPooling4DOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a 4D pooling operation.
//
// # Overview
//
// Use this descriptor with the following methods:
//
// - [MPSGraph.MaxPooling4DWithSourceTensorDescriptorName] -
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName] -
// [MPSGraph.MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName]
// -
// [MPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeDescriptorName]
// -
// [MPSGraph.MaxPooling4DGradientWithGradientTensorIndicesTensorOutputShapeTensorDescriptorName]
// - [MPSGraph.AvgPooling4DWithSourceTensorDescriptorName] -
// [MPSGraph.AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName]
// - [MPSGraph.L2NormPooling4DWithSourceTensorDescriptorName] -
// [MPSGraph.L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphPooling4DOpDescriptor.CeilMode]: Affects how MPSGraph computes the output size: if set to [YES] then output size is computed by rounding up instead of down when dividing input size by stride.
//   - [MPSGraphPooling4DOpDescriptor.SetCeilMode]
//   - [MPSGraphPooling4DOpDescriptor.DilationRates]: Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//   - [MPSGraphPooling4DOpDescriptor.SetDilationRates]
//   - [MPSGraphPooling4DOpDescriptor.IncludeZeroPadToAverage]: Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//   - [MPSGraphPooling4DOpDescriptor.SetIncludeZeroPadToAverage]
//   - [MPSGraphPooling4DOpDescriptor.KernelSizes]: Defines the pooling window size.
//   - [MPSGraphPooling4DOpDescriptor.SetKernelSizes]
//   - [MPSGraphPooling4DOpDescriptor.PaddingStyle]: Defines what kind of padding graph applies to the operation.
//   - [MPSGraphPooling4DOpDescriptor.SetPaddingStyle]
//   - [MPSGraphPooling4DOpDescriptor.PaddingValues]: Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//   - [MPSGraphPooling4DOpDescriptor.SetPaddingValues]
//   - [MPSGraphPooling4DOpDescriptor.ReturnIndicesDataType]: Defines the data type for returned indices.
//   - [MPSGraphPooling4DOpDescriptor.SetReturnIndicesDataType]
//   - [MPSGraphPooling4DOpDescriptor.ReturnIndicesMode]: Defines the mode for returned indices of maximum values within each pooling window.
//   - [MPSGraphPooling4DOpDescriptor.SetReturnIndicesMode]
//   - [MPSGraphPooling4DOpDescriptor.Strides]: Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//   - [MPSGraphPooling4DOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor
type MPSGraphPooling4DOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphPooling4DOpDescriptorFromID constructs a [MPSGraphPooling4DOpDescriptor] from an objc.ID.
//
// The class that defines the parameters for a 4D pooling operation.
func MPSGraphPooling4DOpDescriptorFromID(id objc.ID) MPSGraphPooling4DOpDescriptor {
	return MPSGraphPooling4DOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphPooling4DOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphPooling4DOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphPooling4DOpDescriptor.CeilMode]: Affects how MPSGraph computes the output size: if set to [YES] then output size is computed by rounding up instead of down when dividing input size by stride.
//   - [IMPSGraphPooling4DOpDescriptor.SetCeilMode]
//   - [IMPSGraphPooling4DOpDescriptor.DilationRates]: Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//   - [IMPSGraphPooling4DOpDescriptor.SetDilationRates]
//   - [IMPSGraphPooling4DOpDescriptor.IncludeZeroPadToAverage]: Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//   - [IMPSGraphPooling4DOpDescriptor.SetIncludeZeroPadToAverage]
//   - [IMPSGraphPooling4DOpDescriptor.KernelSizes]: Defines the pooling window size.
//   - [IMPSGraphPooling4DOpDescriptor.SetKernelSizes]
//   - [IMPSGraphPooling4DOpDescriptor.PaddingStyle]: Defines what kind of padding graph applies to the operation.
//   - [IMPSGraphPooling4DOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphPooling4DOpDescriptor.PaddingValues]: Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//   - [IMPSGraphPooling4DOpDescriptor.SetPaddingValues]
//   - [IMPSGraphPooling4DOpDescriptor.ReturnIndicesDataType]: Defines the data type for returned indices.
//   - [IMPSGraphPooling4DOpDescriptor.SetReturnIndicesDataType]
//   - [IMPSGraphPooling4DOpDescriptor.ReturnIndicesMode]: Defines the mode for returned indices of maximum values within each pooling window.
//   - [IMPSGraphPooling4DOpDescriptor.SetReturnIndicesMode]
//   - [IMPSGraphPooling4DOpDescriptor.Strides]: Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//   - [IMPSGraphPooling4DOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor
type IMPSGraphPooling4DOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// Affects how MPSGraph computes the output size: if set to [YES] then output size is computed by rounding up instead of down when dividing input size by stride.
	CeilMode() bool
	SetCeilMode(value bool)
	// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
	DilationRates() []foundation.NSNumber
	SetDilationRates(value []foundation.NSNumber)
	// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	// Defines the pooling window size.
	KernelSizes() []foundation.NSNumber
	SetKernelSizes(value []foundation.NSNumber)
	// Defines what kind of padding graph applies to the operation.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
	PaddingValues() []foundation.NSNumber
	SetPaddingValues(value []foundation.NSNumber)
	// Defines the data type for returned indices.
	ReturnIndicesDataType() uint32
	SetReturnIndicesDataType(value uint32)
	// Defines the mode for returned indices of maximum values within each pooling window.
	ReturnIndicesMode() MPSGraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value MPSGraphPoolingReturnIndicesMode)
	// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
	Strides() []foundation.NSNumber
	SetStrides(value []foundation.NSNumber)
}

// Init initializes the instance.
func (g MPSGraphPooling4DOpDescriptor) Init() MPSGraphPooling4DOpDescriptor {
	rv := objc.Send[MPSGraphPooling4DOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphPooling4DOpDescriptor) Autorelease() MPSGraphPooling4DOpDescriptor {
	rv := objc.Send[MPSGraphPooling4DOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphPooling4DOpDescriptor creates a new MPSGraphPooling4DOpDescriptor instance.
func NewMPSGraphPooling4DOpDescriptor() MPSGraphPooling4DOpDescriptor {
	class := getMPSGraphPooling4DOpDescriptorClass()
	rv := objc.Send[MPSGraphPooling4DOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a 4D pooling descriptor with default values.
//
// kernelSizes: See `kernelSizes` property.
//
// paddingStyle: See `paddingStyle` property.
//
// # Return Value
//
// The descriptor on autoreleasepool.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:paddingStyle:)
func NewGraphPooling4DOpDescriptorWithKernelSizesPaddingStyle(kernelSizes []foundation.NSNumber, paddingStyle MPSGraphPaddingStyle) MPSGraphPooling4DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphPooling4DOpDescriptorClass().class), objc.Sel("descriptorWithKernelSizes:paddingStyle:"), objectivec.IObjectSliceToNSArray(kernelSizes), paddingStyle)
	return MPSGraphPooling4DOpDescriptorFromID(rv)
}

// Creates a 4D pooling descriptor with given values.
//
// kernelSizes: See `kernelSizes` property.
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
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:strides:dilationRates:paddingValues:paddingStyle:)
func NewGraphPooling4DOpDescriptorWithKernelSizesStridesDilationRatesPaddingValuesPaddingStyle(kernelSizes []foundation.NSNumber, strides []foundation.NSNumber, dilationRates []foundation.NSNumber, paddingValues []foundation.NSNumber, paddingStyle MPSGraphPaddingStyle) MPSGraphPooling4DOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphPooling4DOpDescriptorClass().class), objc.Sel("descriptorWithKernelSizes:strides:dilationRates:paddingValues:paddingStyle:"), objectivec.IObjectSliceToNSArray(kernelSizes), objectivec.IObjectSliceToNSArray(strides), objectivec.IObjectSliceToNSArray(dilationRates), objectivec.IObjectSliceToNSArray(paddingValues), paddingStyle)
	return MPSGraphPooling4DOpDescriptorFromID(rv)
}

// Affects how MPSGraph computes the output size: if set to [YES] then output
// size is computed by rounding up instead of down when dividing input size by
// stride.
//
// # Discussion
//
// Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/ceilMode
func (g MPSGraphPooling4DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("ceilMode"))
	return rv
}
func (g MPSGraphPooling4DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setCeilMode:"), value)
}

// Defines dilation rates for spatial dimensions. Must be four numbers, one
// for each spatial dimension, fastest running index last.
//
// # Discussion
//
// Default value: `@[ @1, @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/dilationRates
func (g MPSGraphPooling4DOpDescriptor) DilationRates() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("dilationRates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphPooling4DOpDescriptor) SetDilationRates(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRates:"), objectivec.IObjectSliceToNSArray(value))
}

// Defines a mode for average pooling, where samples outside the input tensor
// count as zeroes in the average computation.
//
// # Discussion
//
// Otherwise the result is sum over samples divided by number of samples that
// didn’t come from padding. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/includeZeroPadToAverage
func (g MPSGraphPooling4DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}
func (g MPSGraphPooling4DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

// Defines the pooling window size.
//
// # Discussion
//
// Must be four numbers, one for each spatial dimension, fastest running index
// last.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/kernelSizes
func (g MPSGraphPooling4DOpDescriptor) KernelSizes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("kernelSizes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphPooling4DOpDescriptor) SetKernelSizes(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setKernelSizes:"), objectivec.IObjectSliceToNSArray(value))
}

// Defines what kind of padding graph applies to the operation.
//
// # Discussion
//
// Default value: [MPSGraphPaddingStyleExplicit].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingStyle
func (g MPSGraphPooling4DOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphPooling4DOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// Defines padding values for spatial dimensions which must be eight numbers,
// two for each spatial dimension.
//
// # Discussion
//
// For example `paddingValues[0]` defines the explicit padding amount before
// the first spatial dimension (slowest running index of spatial dimensions),
// `paddingValues[1]` defines the padding amount after the first spatial
// dimension etc. Used only when `paddingStyle =
// MPSGraphPaddingStyleExplicit`. Default value: `@[ @0, @0, @0, @0, @0, @0,
// @0, @0 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingValues
func (g MPSGraphPooling4DOpDescriptor) PaddingValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("paddingValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphPooling4DOpDescriptor) SetPaddingValues(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingValues:"), objectivec.IObjectSliceToNSArray(value))
}

// Defines the data type for returned indices.
//
// # Discussion
//
// Use this in conjunction with
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName] API.
// Currently MPSGraph supports the following datatypes: [MPSDataTypeInt32].
// Default value: [MPSDataTypeInt32].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesDataType
func (g MPSGraphPooling4DOpDescriptor) ReturnIndicesDataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("returnIndicesDataType"))
	return rv
}
func (g MPSGraphPooling4DOpDescriptor) SetReturnIndicesDataType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setReturnIndicesDataType:"), value)
}

// Defines the mode for returned indices of maximum values within each pooling
// window.
//
// # Discussion
//
// Use this in conjunction with
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName] API. If
// `returnIndicesMode = MPSGraphPoolingReturnIndicesNone` then only the first
// result MPSGraph returns from
// [MPSGraph.MaxPooling4DReturnIndicesWithSourceTensorDescriptorName] will be
// valid and using the second result will assert. Default value:
// [MPSGraphPoolingReturnIndicesNone].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesMode
func (g MPSGraphPooling4DOpDescriptor) ReturnIndicesMode() MPSGraphPoolingReturnIndicesMode {
	rv := objc.Send[MPSGraphPoolingReturnIndicesMode](g.ID, objc.Sel("returnIndicesMode"))
	return MPSGraphPoolingReturnIndicesMode(rv)
}
func (g MPSGraphPooling4DOpDescriptor) SetReturnIndicesMode(value MPSGraphPoolingReturnIndicesMode) {
	objc.Send[struct{}](g.ID, objc.Sel("setReturnIndicesMode:"), value)
}

// Defines strides for spatial dimensions. Must be four numbers, one for each
// spatial dimension, fastest running index last.
//
// # Discussion
//
// Default value: `@[ @1, @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/strides
func (g MPSGraphPooling4DOpDescriptor) Strides() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("strides"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g MPSGraphPooling4DOpDescriptor) SetStrides(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrides:"), objectivec.IObjectSliceToNSArray(value))
}
