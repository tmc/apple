// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol that provides convolution filter weights and bias terms.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource
type MPSCNNConvolutionDataSource interface {
	objectivec.IObject
	foundation.NSCopying

	// BiasTerms protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/biasTerms()
	BiasTerms() unsafe.Pointer

	// DataType protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/dataType()
	DataType() MPSDataType

	// Descriptor protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/descriptor()
	Descriptor() IMPSCNNConvolutionDescriptor

	// Label protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/label()
	Label() string

	// Load protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/load()
	Load() bool

	// Purge protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/purge()
	Purge()

	// Weights protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/weights()
	Weights() unsafe.Pointer
}

// MPSCNNConvolutionDataSourceObject wraps an existing Objective-C object that conforms to the MPSCNNConvolutionDataSource protocol.
type MPSCNNConvolutionDataSourceObject struct {
	foundation.NSCopyingObject
}

func (o MPSCNNConvolutionDataSourceObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MPSCNNConvolutionDataSourceObjectFromID constructs a [MPSCNNConvolutionDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSCNNConvolutionDataSourceObjectFromID(id objc.ID) MPSCNNConvolutionDataSourceObject {
	return MPSCNNConvolutionDataSourceObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/biasTerms()
func (o MPSCNNConvolutionDataSourceObject) BiasTerms() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("biasTerms"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/dataType()
func (o MPSCNNConvolutionDataSourceObject) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](o.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/descriptor()
func (o MPSCNNConvolutionDataSourceObject) Descriptor() IMPSCNNConvolutionDescriptor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("descriptor"))
	return MPSCNNConvolutionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/label()
func (o MPSCNNConvolutionDataSourceObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/load()
func (o MPSCNNConvolutionDataSourceObject) Load() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("load"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/purge()
func (o MPSCNNConvolutionDataSourceObject) Purge() {
	objc.Send[struct{}](o.ID, objc.Sel("purge"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/weights()
func (o MPSCNNConvolutionDataSourceObject) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("weights"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/lookupTableForUInt8Kernel()
func (o MPSCNNConvolutionDataSourceObject) LookupTableForUInt8Kernel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/rangesForUInt8Kernel()
func (o MPSCNNConvolutionDataSourceObject) RangesForUInt8Kernel() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("rangesForUInt8Kernel"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/copy(with:device:)
func (o MPSCNNConvolutionDataSourceObject) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/kernelWeightsDataType()
func (o MPSCNNConvolutionDataSourceObject) KernelWeightsDataType() MPSDataType {
	rv := objc.Send[MPSDataType](o.ID, objc.Sel("kernelWeightsDataType"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/update(with:gradientState:sourceState:)
func (o MPSCNNConvolutionDataSourceObject) UpdateWithCommandBufferGradientStateSourceState(commandBuffer metal.MTLCommandBuffer, gradientState IMPSCNNConvolutionGradientState, sourceState IMPSCNNConvolutionWeightsAndBiasesState) IMPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("updateWithCommandBuffer:gradientState:sourceState:"), commandBuffer, gradientState, sourceState)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/update(with:sourceState:)
func (o MPSCNNConvolutionDataSourceObject) UpdateWithGradientStateSourceState(gradientState IMPSCNNConvolutionGradientState, sourceState IMPSCNNConvolutionWeightsAndBiasesState) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("updateWithGradientState:sourceState:"), gradientState, sourceState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/weightsLayout()
func (o MPSCNNConvolutionDataSourceObject) WeightsLayout() MPSCNNConvolutionWeightsLayout {
	rv := objc.Send[MPSCNNConvolutionWeightsLayout](o.ID, objc.Sel("weightsLayout"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource/weightsQuantizationType()
func (o MPSCNNConvolutionDataSourceObject) WeightsQuantizationType() MPSCNNWeightsQuantizationType {
	rv := objc.Send[MPSCNNWeightsQuantizationType](o.ID, objc.Sel("weightsQuantizationType"))
	return rv
}
