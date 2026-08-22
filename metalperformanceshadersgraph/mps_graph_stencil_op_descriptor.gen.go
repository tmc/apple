// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphStencilOpDescriptor] class.
var (
	_MPSGraphStencilOpDescriptorClass     MPSGraphStencilOpDescriptorClass
	_MPSGraphStencilOpDescriptorClassOnce sync.Once
)

func getMPSGraphStencilOpDescriptorClass() MPSGraphStencilOpDescriptorClass {
	_MPSGraphStencilOpDescriptorClassOnce.Do(func() {
		_MPSGraphStencilOpDescriptorClass = MPSGraphStencilOpDescriptorClass{class: objc.GetClass("MPSGraphStencilOpDescriptor")}
	})
	return _MPSGraphStencilOpDescriptorClass
}

// GetMPSGraphStencilOpDescriptorClass returns the class object for MPSGraphStencilOpDescriptor.
func GetMPSGraphStencilOpDescriptorClass() MPSGraphStencilOpDescriptorClass {
	return getMPSGraphStencilOpDescriptorClass()
}

type MPSGraphStencilOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphStencilOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphStencilOpDescriptorClass) Alloc() MPSGraphStencilOpDescriptor {
	rv := objc.Send[MPSGraphStencilOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a stencil operation.
//
// # Overview
//
// Use this descriptor with the following [MPSGraph] method:
//
// - [MPSGraph.StencilWithSourceTensorWeightsTensorDescriptorName]
//
// # Instance Properties
//
//   - [MPSGraphStencilOpDescriptor.BoundaryMode]: The property that determines which values to use for padding the input tensor.
//   - [MPSGraphStencilOpDescriptor.SetBoundaryMode]
//   - [MPSGraphStencilOpDescriptor.DilationRates]: The property that defines dilation rates for spatial dimensions.
//   - [MPSGraphStencilOpDescriptor.SetDilationRates]
//   - [MPSGraphStencilOpDescriptor.ExplicitPadding]: The property that defines padding values for spatial dimensions.
//   - [MPSGraphStencilOpDescriptor.SetExplicitPadding]
//   - [MPSGraphStencilOpDescriptor.Offsets]: An array of length four that determines from which offset to start reading the input tensor.
//   - [MPSGraphStencilOpDescriptor.SetOffsets]
//   - [MPSGraphStencilOpDescriptor.PaddingConstant]: The padding value for `boundaryMode = MPSGraphPaddingModeConstant`.
//   - [MPSGraphStencilOpDescriptor.SetPaddingConstant]
//   - [MPSGraphStencilOpDescriptor.PaddingStyle]: The property that defines what kind of padding to apply to the stencil operation.
//   - [MPSGraphStencilOpDescriptor.SetPaddingStyle]
//   - [MPSGraphStencilOpDescriptor.ReductionMode]: The reduction mode to use within the stencil window.
//   - [MPSGraphStencilOpDescriptor.SetReductionMode]
//   - [MPSGraphStencilOpDescriptor.Strides]: The property that defines strides for spatial dimensions.
//   - [MPSGraphStencilOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor
type MPSGraphStencilOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphStencilOpDescriptorFromID constructs a [MPSGraphStencilOpDescriptor] from an objc.ID.
//
// The class that defines the parameters for a stencil operation.
func MPSGraphStencilOpDescriptorFromID(id objc.ID) MPSGraphStencilOpDescriptor {
	return MPSGraphStencilOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphStencilOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphStencilOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphStencilOpDescriptor.BoundaryMode]: The property that determines which values to use for padding the input tensor.
//   - [IMPSGraphStencilOpDescriptor.SetBoundaryMode]
//   - [IMPSGraphStencilOpDescriptor.DilationRates]: The property that defines dilation rates for spatial dimensions.
//   - [IMPSGraphStencilOpDescriptor.SetDilationRates]
//   - [IMPSGraphStencilOpDescriptor.ExplicitPadding]: The property that defines padding values for spatial dimensions.
//   - [IMPSGraphStencilOpDescriptor.SetExplicitPadding]
//   - [IMPSGraphStencilOpDescriptor.Offsets]: An array of length four that determines from which offset to start reading the input tensor.
//   - [IMPSGraphStencilOpDescriptor.SetOffsets]
//   - [IMPSGraphStencilOpDescriptor.PaddingConstant]: The padding value for `boundaryMode = MPSGraphPaddingModeConstant`.
//   - [IMPSGraphStencilOpDescriptor.SetPaddingConstant]
//   - [IMPSGraphStencilOpDescriptor.PaddingStyle]: The property that defines what kind of padding to apply to the stencil operation.
//   - [IMPSGraphStencilOpDescriptor.SetPaddingStyle]
//   - [IMPSGraphStencilOpDescriptor.ReductionMode]: The reduction mode to use within the stencil window.
//   - [IMPSGraphStencilOpDescriptor.SetReductionMode]
//   - [IMPSGraphStencilOpDescriptor.Strides]: The property that defines strides for spatial dimensions.
//   - [IMPSGraphStencilOpDescriptor.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor
type IMPSGraphStencilOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The property that determines which values to use for padding the input tensor.
	BoundaryMode() MPSGraphPaddingMode
	SetBoundaryMode(value MPSGraphPaddingMode)
	// The property that defines dilation rates for spatial dimensions.
	DilationRates() foundation.NSArray
	SetDilationRates(value foundation.NSArray)
	// The property that defines padding values for spatial dimensions.
	ExplicitPadding() foundation.NSArray
	SetExplicitPadding(value foundation.NSArray)
	// An array of length four that determines from which offset to start reading the input tensor.
	Offsets() foundation.NSArray
	SetOffsets(value foundation.NSArray)
	// The padding value for `boundaryMode = MPSGraphPaddingModeConstant`.
	PaddingConstant() float32
	SetPaddingConstant(value float32)
	// The property that defines what kind of padding to apply to the stencil operation.
	PaddingStyle() MPSGraphPaddingStyle
	SetPaddingStyle(value MPSGraphPaddingStyle)
	// The reduction mode to use within the stencil window.
	ReductionMode() MPSGraphReductionMode
	SetReductionMode(value MPSGraphReductionMode)
	// The property that defines strides for spatial dimensions.
	Strides() foundation.NSArray
	SetStrides(value foundation.NSArray)
}

// Init initializes the instance.
func (g MPSGraphStencilOpDescriptor) Init() MPSGraphStencilOpDescriptor {
	rv := objc.Send[MPSGraphStencilOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphStencilOpDescriptor) Autorelease() MPSGraphStencilOpDescriptor {
	rv := objc.Send[MPSGraphStencilOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphStencilOpDescriptor creates a new MPSGraphStencilOpDescriptor instance.
func NewMPSGraphStencilOpDescriptor() MPSGraphStencilOpDescriptor {
	class := getMPSGraphStencilOpDescriptorClass()
	rv := objc.Send[MPSGraphStencilOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a stencil operation descriptor with default values.
//
// explicitPadding: See `explicitPadding` property.
//
// # Return Value
//
// # A valid MPSGraphStencilOpDescriptor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(explicitPadding:)
func NewGraphStencilOpDescriptorWithExplicitPadding(explicitPadding foundation.NSArray) MPSGraphStencilOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithExplicitPadding:"), explicitPadding)
	return MPSGraphStencilOpDescriptorFromID(rv)
}

// Creates a stencil operation descriptor with default values.
//
// offsets: See `offsets` property.
//
// explicitPadding: See `explicitPadding` property.
//
// # Return Value
//
// # A valid MPSGraphStencilOpDescriptor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(offsets:explicitPadding:)
func NewGraphStencilOpDescriptorWithOffsetsExplicitPadding(offsets foundation.NSArray, explicitPadding foundation.NSArray) MPSGraphStencilOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithOffsets:explicitPadding:"), offsets, explicitPadding)
	return MPSGraphStencilOpDescriptorFromID(rv)
}

// Creates a stencil operation descriptor with default values.
//
// paddingStyle: See `paddingStyle` property.
//
// # Return Value
//
// # A valid MPSGraphStencilOpDescriptor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(paddingStyle:)
func NewGraphStencilOpDescriptorWithPaddingStyle(paddingStyle MPSGraphPaddingStyle) MPSGraphStencilOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return MPSGraphStencilOpDescriptorFromID(rv)
}

// Creates a stencil operation descriptor with given values.
//
// reductionMode: See `reductionMode` property.
//
// offsets: See `offsets` property.
//
// strides: See `strides` property.
//
// dilationRates: See `dilationRates` property.
//
// explicitPadding: See `explicitPadding` property.
//
// boundaryMode: See `boundaryMode` property.
//
// paddingStyle: See `paddingStyle` property.
//
// paddingConstant: See `paddingConstant` property.
//
// # Return Value
//
// # A valid MPSGraphStencilOpDescriptor object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/init(reductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:)
func NewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode MPSGraphReductionMode, offsets foundation.NSArray, strides foundation.NSArray, dilationRates foundation.NSArray, explicitPadding foundation.NSArray, boundaryMode MPSGraphPaddingMode, paddingStyle MPSGraphPaddingStyle, paddingConstant float32) MPSGraphStencilOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphStencilOpDescriptorClass().class), objc.Sel("descriptorWithReductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:"), reductionMode, offsets, strides, dilationRates, explicitPadding, boundaryMode, paddingStyle, paddingConstant)
	return MPSGraphStencilOpDescriptorFromID(rv)
}

// The property that determines which values to use for padding the input
// tensor.
//
// # Discussion
//
// Default value: [MPSGraphPaddingModeZero].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/boundaryMode
func (g MPSGraphStencilOpDescriptor) BoundaryMode() MPSGraphPaddingMode {
	rv := objc.Send[MPSGraphPaddingMode](g.ID, objc.Sel("boundaryMode"))
	return MPSGraphPaddingMode(rv)
}
func (g MPSGraphStencilOpDescriptor) SetBoundaryMode(value MPSGraphPaddingMode) {
	objc.Send[struct{}](g.ID, objc.Sel("setBoundaryMode:"), value)
}

// The property that defines dilation rates for spatial dimensions.
//
// # Discussion
//
// Must be four numbers, one for each spatial dimension, fastest running index
// last. Default value: `@[ @1, @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/dilationRates
func (g MPSGraphStencilOpDescriptor) DilationRates() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dilationRates"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MPSGraphStencilOpDescriptor) SetDilationRates(value foundation.NSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setDilationRates:"), value)
}

// The property that defines padding values for spatial dimensions.
//
// # Discussion
//
// Must be eight numbers, two for each spatial dimension. For example
// `paddingValues[0]` defines the explicit padding amount before the first
// spatial dimension (slowest running index of spatial dimensions),
// `paddingValues[1]` defines the padding amount after the first spatial
// dimension etc. Used only when `paddingStyle =
// MPSGraphPaddingStyleExplicit`. Default value: `@[ @0, @0, @0, @0, @0, @0,
// @0, @0 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/explicitPadding
func (g MPSGraphStencilOpDescriptor) ExplicitPadding() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("explicitPadding"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MPSGraphStencilOpDescriptor) SetExplicitPadding(value foundation.NSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setExplicitPadding:"), value)
}

// An array of length four that determines from which offset to start reading
// the input tensor.
//
// # Discussion
//
// Only used when `paddingStyle` is [MPSGraphPaddingStyleExplicitOffset]. For
// example zero offset means that the first stencil window will align its
// top-left corner (in 4 dimensions) to the top-left corner of the input
// tensor. Default value: `@[ @0, @0, @0, @0 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/offsets
func (g MPSGraphStencilOpDescriptor) Offsets() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("offsets"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MPSGraphStencilOpDescriptor) SetOffsets(value foundation.NSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setOffsets:"), value)
}

// The padding value for `boundaryMode = MPSGraphPaddingModeConstant`.
//
// # Discussion
//
// Default value: 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingConstant
func (g MPSGraphStencilOpDescriptor) PaddingConstant() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("paddingConstant"))
	return rv
}
func (g MPSGraphStencilOpDescriptor) SetPaddingConstant(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingConstant:"), value)
}

// The property that defines what kind of padding to apply to the stencil
// operation.
//
// # Discussion
//
// Default value: [MPSGraphPaddingStyleExplicit].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/paddingStyle
func (g MPSGraphStencilOpDescriptor) PaddingStyle() MPSGraphPaddingStyle {
	rv := objc.Send[MPSGraphPaddingStyle](g.ID, objc.Sel("paddingStyle"))
	return MPSGraphPaddingStyle(rv)
}
func (g MPSGraphStencilOpDescriptor) SetPaddingStyle(value MPSGraphPaddingStyle) {
	objc.Send[struct{}](g.ID, objc.Sel("setPaddingStyle:"), value)
}

// The reduction mode to use within the stencil window.
//
// # Discussion
//
// Default value: [MPSGraphReductionModeSum].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/reductionMode
func (g MPSGraphStencilOpDescriptor) ReductionMode() MPSGraphReductionMode {
	rv := objc.Send[MPSGraphReductionMode](g.ID, objc.Sel("reductionMode"))
	return MPSGraphReductionMode(rv)
}
func (g MPSGraphStencilOpDescriptor) SetReductionMode(value MPSGraphReductionMode) {
	objc.Send[struct{}](g.ID, objc.Sel("setReductionMode:"), value)
}

// The property that defines strides for spatial dimensions.
//
// # Discussion
//
// Must be four numbers, one for each spatial dimension, fastest running index
// last. Default value: `@[ @1, @1, @1, @1 ]`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphStencilOpDescriptor/strides
func (g MPSGraphStencilOpDescriptor) Strides() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MPSGraphStencilOpDescriptor) SetStrides(value foundation.NSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setStrides:"), value)
}
