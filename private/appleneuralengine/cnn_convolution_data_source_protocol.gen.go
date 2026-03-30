// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSCNNConvolutionDataSource protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource
type MPSCNNConvolutionDataSource interface {
	objectivec.IObject

	// BiasTerms protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/biasTerms
	BiasTerms() unsafe.Pointer

	// DataType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/dataType
	DataType() uint32

	// KernelWeightsDataType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/kernelWeightsDataType
	KernelWeightsDataType() uint32

	// Load protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/load
	Load() bool

	// LookupTableForUInt8Kernel protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/lookupTableForUInt8Kernel
	LookupTableForUInt8Kernel() unsafe.Pointer

	// Purge protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/purge
	Purge()

	// RangesForUInt8Kernel protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/rangesForUInt8Kernel
	RangesForUInt8Kernel() []objectivec.IObject

	// Weights protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weights
	Weights() unsafe.Pointer

	// WeightsLayout protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weightsLayout
	WeightsLayout() uint32

	// WeightsQuantizationType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weightsQuantizationType
	WeightsQuantizationType() uint32
}

// MPSCNNConvolutionDataSourceObject wraps an existing Objective-C object that conforms to the MPSCNNConvolutionDataSource protocol.
type MPSCNNConvolutionDataSourceObject struct {
	objectivec.Object
}

func (o MPSCNNConvolutionDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSCNNConvolutionDataSourceObjectFromID constructs a [MPSCNNConvolutionDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSCNNConvolutionDataSourceObjectFromID(id objc.ID) MPSCNNConvolutionDataSourceObject {
	return MPSCNNConvolutionDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/biasTerms
func (o MPSCNNConvolutionDataSourceObject) BiasTerms() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("biasTerms"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/copyWithZone:device:
func (o MPSCNNConvolutionDataSourceObject) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/dataType
func (o MPSCNNConvolutionDataSourceObject) DataType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/descriptor
func (o MPSCNNConvolutionDataSourceObject) Descriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/kernelWeightsDataType
func (o MPSCNNConvolutionDataSourceObject) KernelWeightsDataType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("kernelWeightsDataType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/label
func (o MPSCNNConvolutionDataSourceObject) Label() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/load
func (o MPSCNNConvolutionDataSourceObject) Load() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("load"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/lookupTableForUInt8Kernel
func (o MPSCNNConvolutionDataSourceObject) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/purge
func (o MPSCNNConvolutionDataSourceObject) Purge() {
	objc.Send[struct{}](o.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/rangesForUInt8Kernel
func (o MPSCNNConvolutionDataSourceObject) RangesForUInt8Kernel() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("rangesForUInt8Kernel"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/updateWithCommandBuffer:gradientState:sourceState:
func (o MPSCNNConvolutionDataSourceObject) UpdateWithCommandBufferGradientStateSourceState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateWithCommandBuffer:gradientState:sourceState:"), buffer, state, state2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/updateWithGradientState:sourceState:
func (o MPSCNNConvolutionDataSourceObject) UpdateWithGradientStateSourceState(state objectivec.IObject, state2 objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateWithGradientState:sourceState:"), state, state2)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weights
func (o MPSCNNConvolutionDataSourceObject) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("weights"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weightsLayout
func (o MPSCNNConvolutionDataSourceObject) WeightsLayout() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("weightsLayout"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/MPSCNNConvolutionDataSource/weightsQuantizationType
func (o MPSCNNConvolutionDataSourceObject) WeightsQuantizationType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("weightsQuantizationType"))
	return rv
}
