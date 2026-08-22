// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSCNNConvolutionDataSource protocol.
type MPSCNNConvolutionDataSource interface {
	objectivec.IObject

	// BiasTerms protocol.
	BiasTerms() unsafe.Pointer

	// CopyWithZoneDevice protocol.
	CopyWithZoneDevice(zone unsafe.Pointer, device objectivec.IObject) objectivec.IObject

	// DataType protocol.
	DataType() uint32

	// Descriptor protocol.
	Descriptor() objectivec.IObject

	// KernelWeightsDataType protocol.
	KernelWeightsDataType() uint32

	// Label protocol.
	Label() objectivec.IObject

	// Load protocol.
	Load() bool

	// LookupTableForUInt8Kernel protocol.
	LookupTableForUInt8Kernel() unsafe.Pointer

	// Purge protocol.
	Purge()

	// RangesForUInt8Kernel protocol.
	RangesForUInt8Kernel() unsafe.Pointer

	// UpdateWithCommandBufferGradientStateSourceState protocol.
	UpdateWithCommandBufferGradientStateSourceState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject) objectivec.IObject

	// UpdateWithGradientStateSourceState protocol.
	UpdateWithGradientStateSourceState(state objectivec.IObject, state2 objectivec.IObject) bool

	// Weights protocol.
	Weights() unsafe.Pointer

	// WeightsLayout protocol.
	WeightsLayout() uint32

	// WeightsQuantizationType protocol.
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

func (o MPSCNNConvolutionDataSourceObject) BiasTerms() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("biasTerms"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}
func (o MPSCNNConvolutionDataSourceObject) DataType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("dataType"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) Descriptor() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("descriptor"))
	return objectivec.Object{ID: rv}
}
func (o MPSCNNConvolutionDataSourceObject) KernelWeightsDataType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("kernelWeightsDataType"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) Label() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}
func (o MPSCNNConvolutionDataSourceObject) Load() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("load"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) Purge() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("purge"))
}
func (o MPSCNNConvolutionDataSourceObject) RangesForUInt8Kernel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("rangesForUInt8Kernel"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) UpdateWithCommandBufferGradientStateSourceState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("updateWithCommandBuffer:gradientState:sourceState:"), buffer, state, state2)
	return objectivec.Object{ID: rv}
}
func (o MPSCNNConvolutionDataSourceObject) UpdateWithGradientStateSourceState(state objectivec.IObject, state2 objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("updateWithGradientState:sourceState:"), state, state2)
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) Weights() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("weights"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) WeightsLayout() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("weightsLayout"))
	return rv
}
func (o MPSCNNConvolutionDataSourceObject) WeightsQuantizationType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("weightsQuantizationType"))
	return rv
}
