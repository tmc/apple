// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNMatrixTrainingLayer] class.
var (
	_MPSRNNMatrixTrainingLayerClass     MPSRNNMatrixTrainingLayerClass
	_MPSRNNMatrixTrainingLayerClassOnce sync.Once
)

func getMPSRNNMatrixTrainingLayerClass() MPSRNNMatrixTrainingLayerClass {
	_MPSRNNMatrixTrainingLayerClassOnce.Do(func() {
		_MPSRNNMatrixTrainingLayerClass = MPSRNNMatrixTrainingLayerClass{class: objc.GetClass("MPSRNNMatrixTrainingLayer")}
	})
	return _MPSRNNMatrixTrainingLayerClass
}

// GetMPSRNNMatrixTrainingLayerClass returns the class object for MPSRNNMatrixTrainingLayer.
func GetMPSRNNMatrixTrainingLayerClass() MPSRNNMatrixTrainingLayerClass {
	return getMPSRNNMatrixTrainingLayerClass()
}

type MPSRNNMatrixTrainingLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNMatrixTrainingLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNMatrixTrainingLayerClass) Alloc() MPSRNNMatrixTrainingLayer {
	rv := objc.Send[MPSRNNMatrixTrainingLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A layer for training recurrent neural networks on Metal Performance Shaders
// matrices.
//
// # Initializers
//
//   - [MPSRNNMatrixTrainingLayer.InitWithDeviceRnnDescriptorTrainableWeights]
//
// # Instance Properties
//
//   - [MPSRNNMatrixTrainingLayer.AccumulateWeightGradients]
//   - [MPSRNNMatrixTrainingLayer.SetAccumulateWeightGradients]
//   - [MPSRNNMatrixTrainingLayer.InputFeatureChannels]
//   - [MPSRNNMatrixTrainingLayer.OutputFeatureChannels]
//   - [MPSRNNMatrixTrainingLayer.RecurrentOutputIsTemporary]
//   - [MPSRNNMatrixTrainingLayer.SetRecurrentOutputIsTemporary]
//   - [MPSRNNMatrixTrainingLayer.StoreAllIntermediateStates]
//   - [MPSRNNMatrixTrainingLayer.SetStoreAllIntermediateStates]
//   - [MPSRNNMatrixTrainingLayer.TrainingStateIsTemporary]
//   - [MPSRNNMatrixTrainingLayer.SetTrainingStateIsTemporary]
//
// # Instance Methods
//
//   - [MPSRNNMatrixTrainingLayer.CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer]
//   - [MPSRNNMatrixTrainingLayer.CreateWeightGradientMatricesDataType]
//   - [MPSRNNMatrixTrainingLayer.CreateWeightMatrices]
//   - [MPSRNNMatrixTrainingLayer.EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset]
//   - [MPSRNNMatrixTrainingLayer.EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights]
//   - [MPSRNNMatrixTrainingLayer.EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights]
//   - [MPSRNNMatrixTrainingLayer.EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights]
//   - [MPSRNNMatrixTrainingLayer.EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer
type MPSRNNMatrixTrainingLayer struct {
	MPSKernel
}

// MPSRNNMatrixTrainingLayerFromID constructs a [MPSRNNMatrixTrainingLayer] from an objc.ID.
//
// A layer for training recurrent neural networks on Metal Performance Shaders
// matrices.
func MPSRNNMatrixTrainingLayerFromID(id objc.ID) MPSRNNMatrixTrainingLayer {
	return MPSRNNMatrixTrainingLayer{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSRNNMatrixTrainingLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNMatrixTrainingLayer] class.
//
// # Initializers
//
//   - [IMPSRNNMatrixTrainingLayer.InitWithDeviceRnnDescriptorTrainableWeights]
//
// # Instance Properties
//
//   - [IMPSRNNMatrixTrainingLayer.AccumulateWeightGradients]
//   - [IMPSRNNMatrixTrainingLayer.SetAccumulateWeightGradients]
//   - [IMPSRNNMatrixTrainingLayer.InputFeatureChannels]
//   - [IMPSRNNMatrixTrainingLayer.OutputFeatureChannels]
//   - [IMPSRNNMatrixTrainingLayer.RecurrentOutputIsTemporary]
//   - [IMPSRNNMatrixTrainingLayer.SetRecurrentOutputIsTemporary]
//   - [IMPSRNNMatrixTrainingLayer.StoreAllIntermediateStates]
//   - [IMPSRNNMatrixTrainingLayer.SetStoreAllIntermediateStates]
//   - [IMPSRNNMatrixTrainingLayer.TrainingStateIsTemporary]
//   - [IMPSRNNMatrixTrainingLayer.SetTrainingStateIsTemporary]
//
// # Instance Methods
//
//   - [IMPSRNNMatrixTrainingLayer.CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer]
//   - [IMPSRNNMatrixTrainingLayer.CreateWeightGradientMatricesDataType]
//   - [IMPSRNNMatrixTrainingLayer.CreateWeightMatrices]
//   - [IMPSRNNMatrixTrainingLayer.EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset]
//   - [IMPSRNNMatrixTrainingLayer.EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights]
//   - [IMPSRNNMatrixTrainingLayer.EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights]
//   - [IMPSRNNMatrixTrainingLayer.EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights]
//   - [IMPSRNNMatrixTrainingLayer.EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer
type IMPSRNNMatrixTrainingLayer interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceRnnDescriptorTrainableWeights(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor, trainableWeights foundation.INSArray) MPSRNNMatrixTrainingLayer

	// Topic: Instance Properties

	AccumulateWeightGradients() bool
	SetAccumulateWeightGradients(value bool)
	InputFeatureChannels() uint
	OutputFeatureChannels() uint
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
	TrainingStateIsTemporary() bool
	SetTrainingStateIsTemporary(value bool)

	// Topic: Instance Methods

	CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut foundation.INSArray, dataType MPSDataType, commandBuffer metal.MTLCommandBuffer)
	CreateWeightGradientMatricesDataType(matricesOut foundation.INSArray, dataType MPSDataType)
	CreateWeightMatrices(matricesOut foundation.INSArray)
	EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer metal.MTLCommandBuffer, weights []MPSMatrix, matrixId MPSRNNMatrixId, matrix IMPSMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.MTLOrigin)
	EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, trainingStates foundation.INSArray, weights []MPSMatrix)
	EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, sourceOffsets *uint, destinationMatrices []MPSMatrix, destinationOffsets *uint, trainingStates foundation.INSArray, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray, weights []MPSMatrix)
	EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer metal.MTLCommandBuffer, forwardSources []MPSMatrix, forwardSourceOffsets *uint, sourceGradients []MPSMatrix, sourceGradientOffsets *uint, destinationGradients []MPSMatrix, destinationOffsets *uint, weightGradients []MPSMatrix, trainingStates []MPSRNNMatrixTrainingState, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray, weights []MPSMatrix)
	EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer metal.MTLCommandBuffer, forwardSources []MPSMatrix, sourceGradients []MPSMatrix, destinationGradients []MPSMatrix, weightGradients []MPSMatrix, trainingStates []MPSRNNMatrixTrainingState, weights []MPSMatrix)
}

// Init initializes the instance.
func (r MPSRNNMatrixTrainingLayer) Init() MPSRNNMatrixTrainingLayer {
	rv := objc.Send[MPSRNNMatrixTrainingLayer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNMatrixTrainingLayer) Autorelease() MPSRNNMatrixTrainingLayer {
	rv := objc.Send[MPSRNNMatrixTrainingLayer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNMatrixTrainingLayer creates a new MPSRNNMatrixTrainingLayer instance.
func NewMPSRNNMatrixTrainingLayer() MPSRNNMatrixTrainingLayer {
	class := getMPSRNNMatrixTrainingLayerClass()
	rv := objc.Send[MPSRNNMatrixTrainingLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewRNNMatrixTrainingLayerWithCoder(aDecoder foundation.INSCoder) MPSRNNMatrixTrainingLayer {
	instance := getMPSRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSRNNMatrixTrainingLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/init(coder:device:)
func NewRNNMatrixTrainingLayerWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSRNNMatrixTrainingLayer {
	instance := getMPSRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSRNNMatrixTrainingLayerFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewRNNMatrixTrainingLayerWithDevice(device metal.MTLDevice) MPSRNNMatrixTrainingLayer {
	instance := getMPSRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSRNNMatrixTrainingLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/init(device:rnnDescriptor:trainableWeights:)
func NewRNNMatrixTrainingLayerWithDeviceRnnDescriptorTrainableWeights(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor, trainableWeights foundation.INSArray) MPSRNNMatrixTrainingLayer {
	instance := getMPSRNNMatrixTrainingLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:trainableWeights:"), device, rnnDescriptor, trainableWeights)
	return MPSRNNMatrixTrainingLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/init(device:rnnDescriptor:trainableWeights:)
func (r MPSRNNMatrixTrainingLayer) InitWithDeviceRnnDescriptorTrainableWeights(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor, trainableWeights foundation.INSArray) MPSRNNMatrixTrainingLayer {
	rv := objc.Send[MPSRNNMatrixTrainingLayer](r.ID, objc.Sel("initWithDevice:rnnDescriptor:trainableWeights:"), device, rnnDescriptor, trainableWeights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createTemporaryWeightGradientMatrices(_:dataType:commandBuffer:)
func (r MPSRNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut foundation.INSArray, dataType MPSDataType, commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](r.ID, objc.Sel("createTemporaryWeightGradientMatrices:dataType:commandBuffer:"), matricesOut, dataType, commandBuffer)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createWeightGradientMatrices(_:dataType:)
func (r MPSRNNMatrixTrainingLayer) CreateWeightGradientMatricesDataType(matricesOut foundation.INSArray, dataType MPSDataType) {
	objc.Send[objc.ID](r.ID, objc.Sel("createWeightGradientMatrices:dataType:"), matricesOut, dataType)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/createWeightMatrices(_:)
func (r MPSRNNMatrixTrainingLayer) CreateWeightMatrices(matricesOut foundation.INSArray) {
	objc.Send[objc.ID](r.ID, objc.Sel("createWeightMatrices:"), matricesOut)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/encodeCopyWeights(commandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:)
func (r MPSRNNMatrixTrainingLayer) EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer metal.MTLCommandBuffer, weights []MPSMatrix, matrixId MPSRNNMatrixId, matrix IMPSMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.MTLOrigin) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeCopyWeightsToCommandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:"), commandBuffer, objectivec.IObjectSliceToNSArray(weights), matrixId, matrix, copyFromWeightsToMatrix, matrixOffset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/encodeForwardSequence(commandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:)
func (r MPSRNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, trainingStates foundation.INSArray, weights []MPSMatrix) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceMatrices), objectivec.IObjectSliceToNSArray(destinationMatrices), trainingStates, objectivec.IObjectSliceToNSArray(weights))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/encodeForwardSequence(commandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:trainingStates:recurrentInputState:recurrentOutputStates:weights:)
func (r MPSRNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, sourceOffsets *uint, destinationMatrices []MPSMatrix, destinationOffsets *uint, trainingStates foundation.INSArray, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray, weights []MPSMatrix) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceMatrices), unsafe.Pointer(sourceOffsets), objectivec.IObjectSliceToNSArray(destinationMatrices), unsafe.Pointer(destinationOffsets), trainingStates, recurrentInputState, recurrentOutputStates, objectivec.IObjectSliceToNSArray(weights))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/encodeGradientSequence(commandBuffer:forwardSources:forwardSourceOffsets:sourceGradients:sourceOffsets:destinationGradients:destinationOffsets:weightGradients:trainingStates:recurrentInputState:recurrentOutputStates:weights:)
func (r MPSRNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesForwardSourceOffsetsSourceGradientsSourceGradientOffsetsDestinationGradientsDestinationOffsetsWeightGradientsTrainingStatesRecurrentInputStateRecurrentOutputStatesWeights(commandBuffer metal.MTLCommandBuffer, forwardSources []MPSMatrix, forwardSourceOffsets *uint, sourceGradients []MPSMatrix, sourceGradientOffsets *uint, destinationGradients []MPSMatrix, destinationOffsets *uint, weightGradients []MPSMatrix, trainingStates []MPSRNNMatrixTrainingState, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray, weights []MPSMatrix) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:forwardSourceOffsets:sourceGradients:sourceGradientOffsets:destinationGradients:destinationOffsets:weightGradients:trainingStates:recurrentInputState:recurrentOutputStates:weights:"), commandBuffer, objectivec.IObjectSliceToNSArray(forwardSources), unsafe.Pointer(forwardSourceOffsets), objectivec.IObjectSliceToNSArray(sourceGradients), unsafe.Pointer(sourceGradientOffsets), objectivec.IObjectSliceToNSArray(destinationGradients), unsafe.Pointer(destinationOffsets), objectivec.IObjectSliceToNSArray(weightGradients), objectivec.IObjectSliceToNSArray(trainingStates), recurrentInputState, recurrentOutputStates, objectivec.IObjectSliceToNSArray(weights))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/encodeGradientSequence(commandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:)
func (r MPSRNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer metal.MTLCommandBuffer, forwardSources []MPSMatrix, sourceGradients []MPSMatrix, destinationGradients []MPSMatrix, weightGradients []MPSMatrix, trainingStates []MPSRNNMatrixTrainingState, weights []MPSMatrix) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:"), commandBuffer, objectivec.IObjectSliceToNSArray(forwardSources), objectivec.IObjectSliceToNSArray(sourceGradients), objectivec.IObjectSliceToNSArray(destinationGradients), objectivec.IObjectSliceToNSArray(weightGradients), objectivec.IObjectSliceToNSArray(trainingStates), objectivec.IObjectSliceToNSArray(weights))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/accumulateWeightGradients
func (r MPSRNNMatrixTrainingLayer) AccumulateWeightGradients() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("accumulateWeightGradients"))
	return rv
}
func (r MPSRNNMatrixTrainingLayer) SetAccumulateWeightGradients(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setAccumulateWeightGradients:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/inputFeatureChannels
func (r MPSRNNMatrixTrainingLayer) InputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/outputFeatureChannels
func (r MPSRNNMatrixTrainingLayer) OutputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("outputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/recurrentOutputIsTemporary
func (r MPSRNNMatrixTrainingLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}
func (r MPSRNNMatrixTrainingLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/storeAllIntermediateStates
func (r MPSRNNMatrixTrainingLayer) StoreAllIntermediateStates() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}
func (r MPSRNNMatrixTrainingLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer/trainingStateIsTemporary
func (r MPSRNNMatrixTrainingLayer) TrainingStateIsTemporary() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("trainingStateIsTemporary"))
	return rv
}
func (r MPSRNNMatrixTrainingLayer) SetTrainingStateIsTemporary(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setTrainingStateIsTemporary:"), value)
}
