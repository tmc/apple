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

// The class instance for the [MPSRNNMatrixInferenceLayer] class.
var (
	_MPSRNNMatrixInferenceLayerClass     MPSRNNMatrixInferenceLayerClass
	_MPSRNNMatrixInferenceLayerClassOnce sync.Once
)

func getMPSRNNMatrixInferenceLayerClass() MPSRNNMatrixInferenceLayerClass {
	_MPSRNNMatrixInferenceLayerClassOnce.Do(func() {
		_MPSRNNMatrixInferenceLayerClass = MPSRNNMatrixInferenceLayerClass{class: objc.GetClass("MPSRNNMatrixInferenceLayer")}
	})
	return _MPSRNNMatrixInferenceLayerClass
}

// GetMPSRNNMatrixInferenceLayerClass returns the class object for MPSRNNMatrixInferenceLayer.
func GetMPSRNNMatrixInferenceLayerClass() MPSRNNMatrixInferenceLayerClass {
	return getMPSRNNMatrixInferenceLayerClass()
}

type MPSRNNMatrixInferenceLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNMatrixInferenceLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNMatrixInferenceLayerClass) Alloc() MPSRNNMatrixInferenceLayer {
	rv := objc.Send[MPSRNNMatrixInferenceLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A recurrent neural network layer for inference on Metal Performance Shaders
// matrices.
//
// # Overview
//
// The [MPSRNNMatrixInferenceLayer] specifies a recurrent neural network layer
// for inference on [MPSMatrix] objects. Two types of recurrent layers are
// supported:
//
// - [MPSRNNImageInferenceLayer]—Operates with convolutions on images. -
// [MPSRNNMatrixInferenceLayer]—Operates on matrices.
//
// You can use [MPSRNNImageInferenceLayer] to implement the latter by using 1
// x 1 images, but due to image size restrictions and performance,
// [MPSRNNMatrixInferenceLayer] is the better choice for linear recurrent
// layers.
//
// [MPSRNNMatrixInferenceLayer] is initialized using either of the following:
//
// - A single [MPSRNNDescriptor] instance, which further specifies the
// recurrent network layer. - An array of [MPSRNNDescriptor] instances, which
// specifies a stack of recurrent layers that can operate in parallel a subset
// of the inputs in a sequence of inputs and recurrent outputs.
//
// Stacks with bidirectionally traversing encode functions don’t support
// starting from a previous set of recurrent states. However, you can achieve
// this effect by defining two separate unidirectional stacks of layers,
// running the same input sequence on them separately (one forward and one
// backward), and ultimately combining the two result sequences.
//
// # Initializers
//
//   - [MPSRNNMatrixInferenceLayer.InitWithDeviceRnnDescriptor]
//   - [MPSRNNMatrixInferenceLayer.InitWithDeviceRnnDescriptors]
//
// # Instance Properties
//
//   - [MPSRNNMatrixInferenceLayer.BidirectionalCombineMode]
//   - [MPSRNNMatrixInferenceLayer.SetBidirectionalCombineMode]
//   - [MPSRNNMatrixInferenceLayer.InputFeatureChannels]
//   - [MPSRNNMatrixInferenceLayer.NumberOfLayers]
//   - [MPSRNNMatrixInferenceLayer.OutputFeatureChannels]
//   - [MPSRNNMatrixInferenceLayer.RecurrentOutputIsTemporary]
//   - [MPSRNNMatrixInferenceLayer.SetRecurrentOutputIsTemporary]
//   - [MPSRNNMatrixInferenceLayer.StoreAllIntermediateStates]
//   - [MPSRNNMatrixInferenceLayer.SetStoreAllIntermediateStates]
//
// # Instance Methods
//
//   - [MPSRNNMatrixInferenceLayer.EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices]
//   - [MPSRNNMatrixInferenceLayer.EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates]
//   - [MPSRNNMatrixInferenceLayer.EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer
type MPSRNNMatrixInferenceLayer struct {
	MPSKernel
}

// MPSRNNMatrixInferenceLayerFromID constructs a [MPSRNNMatrixInferenceLayer] from an objc.ID.
//
// A recurrent neural network layer for inference on Metal Performance Shaders
// matrices.
func MPSRNNMatrixInferenceLayerFromID(id objc.ID) MPSRNNMatrixInferenceLayer {
	return MPSRNNMatrixInferenceLayer{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSRNNMatrixInferenceLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNMatrixInferenceLayer] class.
//
// # Initializers
//
//   - [IMPSRNNMatrixInferenceLayer.InitWithDeviceRnnDescriptor]
//   - [IMPSRNNMatrixInferenceLayer.InitWithDeviceRnnDescriptors]
//
// # Instance Properties
//
//   - [IMPSRNNMatrixInferenceLayer.BidirectionalCombineMode]
//   - [IMPSRNNMatrixInferenceLayer.SetBidirectionalCombineMode]
//   - [IMPSRNNMatrixInferenceLayer.InputFeatureChannels]
//   - [IMPSRNNMatrixInferenceLayer.NumberOfLayers]
//   - [IMPSRNNMatrixInferenceLayer.OutputFeatureChannels]
//   - [IMPSRNNMatrixInferenceLayer.RecurrentOutputIsTemporary]
//   - [IMPSRNNMatrixInferenceLayer.SetRecurrentOutputIsTemporary]
//   - [IMPSRNNMatrixInferenceLayer.StoreAllIntermediateStates]
//   - [IMPSRNNMatrixInferenceLayer.SetStoreAllIntermediateStates]
//
// # Instance Methods
//
//   - [IMPSRNNMatrixInferenceLayer.EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices]
//   - [IMPSRNNMatrixInferenceLayer.EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates]
//   - [IMPSRNNMatrixInferenceLayer.EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer
type IMPSRNNMatrixInferenceLayer interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNMatrixInferenceLayer
	InitWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNMatrixInferenceLayer

	// Topic: Instance Properties

	BidirectionalCombineMode() MPSRNNBidirectionalCombineMode
	SetBidirectionalCombineMode(value MPSRNNBidirectionalCombineMode)
	InputFeatureChannels() uint
	NumberOfLayers() uint
	OutputFeatureChannels() uint
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)

	// Topic: Instance Methods

	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer metal.MTLCommandBuffer, sourceSequence []MPSMatrix, destinationForwardMatrices []MPSMatrix, destinationBackwardMatrices []MPSMatrix)
	EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray)
	EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, sourceOffsets *uint, destinationMatrices []MPSMatrix, destinationOffsets *uint, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray)
}

// Init initializes the instance.
func (r MPSRNNMatrixInferenceLayer) Init() MPSRNNMatrixInferenceLayer {
	rv := objc.Send[MPSRNNMatrixInferenceLayer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNMatrixInferenceLayer) Autorelease() MPSRNNMatrixInferenceLayer {
	rv := objc.Send[MPSRNNMatrixInferenceLayer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNMatrixInferenceLayer creates a new MPSRNNMatrixInferenceLayer instance.
func NewMPSRNNMatrixInferenceLayer() MPSRNNMatrixInferenceLayer {
	class := getMPSRNNMatrixInferenceLayerClass()
	rv := objc.Send[MPSRNNMatrixInferenceLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewRNNMatrixInferenceLayerWithCoder(aDecoder foundation.INSCoder) MPSRNNMatrixInferenceLayer {
	instance := getMPSRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSRNNMatrixInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/init(coder:device:)
func NewRNNMatrixInferenceLayerWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSRNNMatrixInferenceLayer {
	instance := getMPSRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSRNNMatrixInferenceLayerFromID(rv)
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
func NewRNNMatrixInferenceLayerWithDevice(device metal.MTLDevice) MPSRNNMatrixInferenceLayer {
	instance := getMPSRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSRNNMatrixInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/init(device:rnnDescriptor:)
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNMatrixInferenceLayer {
	instance := getMPSRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	return MPSRNNMatrixInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/init(device:rnnDescriptors:)
func NewRNNMatrixInferenceLayerWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNMatrixInferenceLayer {
	instance := getMPSRNNMatrixInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, objectivec.IObjectSliceToNSArray(rnnDescriptors))
	return MPSRNNMatrixInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/init(device:rnnDescriptor:)
func (r MPSRNNMatrixInferenceLayer) InitWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNMatrixInferenceLayer {
	rv := objc.Send[MPSRNNMatrixInferenceLayer](r.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/init(device:rnnDescriptors:)
func (r MPSRNNMatrixInferenceLayer) InitWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNMatrixInferenceLayer {
	rv := objc.Send[MPSRNNMatrixInferenceLayer](r.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, objectivec.IObjectSliceToNSArray(rnnDescriptors))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/encodeBidirectionalSequence(commandBuffer:sourceSequence:destinationForwardMatrices:destinationBackwardMatrices:)
func (r MPSRNNMatrixInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardMatricesDestinationBackwardMatrices(commandBuffer metal.MTLCommandBuffer, sourceSequence []MPSMatrix, destinationForwardMatrices []MPSMatrix, destinationBackwardMatrices []MPSMatrix) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardMatrices:destinationBackwardMatrices:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceSequence), objectivec.IObjectSliceToNSArray(destinationForwardMatrices), objectivec.IObjectSliceToNSArray(destinationBackwardMatrices))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/encodeSequence(commandBuffer:sourceMatrices:destinationMatrices:recurrentInputState:recurrentOutputStates:)
func (r MPSRNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesDestinationMatricesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:destinationMatrices:recurrentInputState:recurrentOutputStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceMatrices), objectivec.IObjectSliceToNSArray(destinationMatrices), recurrentInputState, recurrentOutputStates)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/encodeSequence(commandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:recurrentInputState:recurrentOutputStates:)
func (r MPSRNNMatrixInferenceLayer) EncodeSequenceToCommandBufferSourceMatricesSourceOffsetsDestinationMatricesDestinationOffsetsRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceMatrices []MPSMatrix, sourceOffsets *uint, destinationMatrices []MPSMatrix, destinationOffsets *uint, recurrentInputState IMPSRNNRecurrentMatrixState, recurrentOutputStates foundation.INSArray) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceMatrices:sourceOffsets:destinationMatrices:destinationOffsets:recurrentInputState:recurrentOutputStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceMatrices), unsafe.Pointer(sourceOffsets), objectivec.IObjectSliceToNSArray(destinationMatrices), unsafe.Pointer(destinationOffsets), recurrentInputState, recurrentOutputStates)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/bidirectionalCombineMode
func (r MPSRNNMatrixInferenceLayer) BidirectionalCombineMode() MPSRNNBidirectionalCombineMode {
	rv := objc.Send[MPSRNNBidirectionalCombineMode](r.ID, objc.Sel("bidirectionalCombineMode"))
	return MPSRNNBidirectionalCombineMode(rv)
}
func (r MPSRNNMatrixInferenceLayer) SetBidirectionalCombineMode(value MPSRNNBidirectionalCombineMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setBidirectionalCombineMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/inputFeatureChannels
func (r MPSRNNMatrixInferenceLayer) InputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/numberOfLayers
func (r MPSRNNMatrixInferenceLayer) NumberOfLayers() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("numberOfLayers"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/outputFeatureChannels
func (r MPSRNNMatrixInferenceLayer) OutputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("outputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/recurrentOutputIsTemporary
func (r MPSRNNMatrixInferenceLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}
func (r MPSRNNMatrixInferenceLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixInferenceLayer/storeAllIntermediateStates
func (r MPSRNNMatrixInferenceLayer) StoreAllIntermediateStates() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}
func (r MPSRNNMatrixInferenceLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}
