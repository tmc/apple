// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNImageInferenceLayer] class.
var (
	_MPSRNNImageInferenceLayerClass     MPSRNNImageInferenceLayerClass
	_MPSRNNImageInferenceLayerClassOnce sync.Once
)

func getMPSRNNImageInferenceLayerClass() MPSRNNImageInferenceLayerClass {
	_MPSRNNImageInferenceLayerClassOnce.Do(func() {
		_MPSRNNImageInferenceLayerClass = MPSRNNImageInferenceLayerClass{class: objc.GetClass("MPSRNNImageInferenceLayer")}
	})
	return _MPSRNNImageInferenceLayerClass
}

// GetMPSRNNImageInferenceLayerClass returns the class object for MPSRNNImageInferenceLayer.
func GetMPSRNNImageInferenceLayerClass() MPSRNNImageInferenceLayerClass {
	return getMPSRNNImageInferenceLayerClass()
}

type MPSRNNImageInferenceLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNImageInferenceLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNImageInferenceLayerClass) Alloc() MPSRNNImageInferenceLayer {
	rv := objc.Send[MPSRNNImageInferenceLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A recurrent neural network layer for inference on Metal Performance Shaders
// images.
//
// # Overview
//
// The [MPSRNNImageInferenceLayer] specifies a recurrent neural network layer
// for inference on [MPSImage] objects. Two types of recurrent layers are
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
// [MPSRNNImageInferenceLayer] is initialized using either of the following:
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
//   - [MPSRNNImageInferenceLayer.InitWithDeviceRnnDescriptor]
//   - [MPSRNNImageInferenceLayer.InitWithDeviceRnnDescriptors]
//
// # Instance Properties
//
//   - [MPSRNNImageInferenceLayer.BidirectionalCombineMode]
//   - [MPSRNNImageInferenceLayer.SetBidirectionalCombineMode]
//   - [MPSRNNImageInferenceLayer.NumberOfLayers]
//   - [MPSRNNImageInferenceLayer.RecurrentOutputIsTemporary]
//   - [MPSRNNImageInferenceLayer.SetRecurrentOutputIsTemporary]
//   - [MPSRNNImageInferenceLayer.StoreAllIntermediateStates]
//   - [MPSRNNImageInferenceLayer.SetStoreAllIntermediateStates]
//   - [MPSRNNImageInferenceLayer.InputFeatureChannels]
//   - [MPSRNNImageInferenceLayer.OutputFeatureChannels]
//
// # Instance Methods
//
//   - [MPSRNNImageInferenceLayer.EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages]
//   - [MPSRNNImageInferenceLayer.EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer
type MPSRNNImageInferenceLayer struct {
	MPSCNNKernel
}

// MPSRNNImageInferenceLayerFromID constructs a [MPSRNNImageInferenceLayer] from an objc.ID.
//
// A recurrent neural network layer for inference on Metal Performance Shaders
// images.
func MPSRNNImageInferenceLayerFromID(id objc.ID) MPSRNNImageInferenceLayer {
	return MPSRNNImageInferenceLayer{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSRNNImageInferenceLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNImageInferenceLayer] class.
//
// # Initializers
//
//   - [IMPSRNNImageInferenceLayer.InitWithDeviceRnnDescriptor]
//   - [IMPSRNNImageInferenceLayer.InitWithDeviceRnnDescriptors]
//
// # Instance Properties
//
//   - [IMPSRNNImageInferenceLayer.BidirectionalCombineMode]
//   - [IMPSRNNImageInferenceLayer.SetBidirectionalCombineMode]
//   - [IMPSRNNImageInferenceLayer.NumberOfLayers]
//   - [IMPSRNNImageInferenceLayer.RecurrentOutputIsTemporary]
//   - [IMPSRNNImageInferenceLayer.SetRecurrentOutputIsTemporary]
//   - [IMPSRNNImageInferenceLayer.StoreAllIntermediateStates]
//   - [IMPSRNNImageInferenceLayer.SetStoreAllIntermediateStates]
//   - [IMPSRNNImageInferenceLayer.InputFeatureChannels]
//   - [IMPSRNNImageInferenceLayer.OutputFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSRNNImageInferenceLayer.EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages]
//   - [IMPSRNNImageInferenceLayer.EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer
type IMPSRNNImageInferenceLayer interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNImageInferenceLayer
	InitWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNImageInferenceLayer

	// Topic: Instance Properties

	BidirectionalCombineMode() MPSRNNBidirectionalCombineMode
	SetBidirectionalCombineMode(value MPSRNNBidirectionalCombineMode)
	NumberOfLayers() uint
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
	InputFeatureChannels() uint
	OutputFeatureChannels() uint

	// Topic: Instance Methods

	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer metal.MTLCommandBuffer, sourceSequence []MPSImage, destinationForwardImages []MPSImage, destinationBackwardImages []MPSImage)
	EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, destinationImages []MPSImage, recurrentInputState IMPSRNNRecurrentImageState, recurrentOutputStates foundation.INSArray)
}

// Init initializes the instance.
func (r MPSRNNImageInferenceLayer) Init() MPSRNNImageInferenceLayer {
	rv := objc.Send[MPSRNNImageInferenceLayer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNImageInferenceLayer) Autorelease() MPSRNNImageInferenceLayer {
	rv := objc.Send[MPSRNNImageInferenceLayer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNImageInferenceLayer creates a new MPSRNNImageInferenceLayer instance.
func NewMPSRNNImageInferenceLayer() MPSRNNImageInferenceLayer {
	class := getMPSRNNImageInferenceLayerClass()
	rv := objc.Send[MPSRNNImageInferenceLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewRNNImageInferenceLayerWithCoder(aDecoder foundation.INSCoder) MPSRNNImageInferenceLayer {
	instance := getMPSRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSRNNImageInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/init(coder:device:)
func NewRNNImageInferenceLayerWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSRNNImageInferenceLayer {
	instance := getMPSRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSRNNImageInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewRNNImageInferenceLayerWithDevice(device metal.MTLDevice) MPSRNNImageInferenceLayer {
	instance := getMPSRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSRNNImageInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/init(device:rnnDescriptor:)
func NewRNNImageInferenceLayerWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNImageInferenceLayer {
	instance := getMPSRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	return MPSRNNImageInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/init(device:rnnDescriptors:)
func NewRNNImageInferenceLayerWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNImageInferenceLayer {
	instance := getMPSRNNImageInferenceLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, objectivec.IObjectSliceToNSArray(rnnDescriptors))
	return MPSRNNImageInferenceLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/init(device:rnnDescriptor:)
func (r MPSRNNImageInferenceLayer) InitWithDeviceRnnDescriptor(device metal.MTLDevice, rnnDescriptor IMPSRNNDescriptor) MPSRNNImageInferenceLayer {
	rv := objc.Send[MPSRNNImageInferenceLayer](r.ID, objc.Sel("initWithDevice:rnnDescriptor:"), device, rnnDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/init(device:rnnDescriptors:)
func (r MPSRNNImageInferenceLayer) InitWithDeviceRnnDescriptors(device metal.MTLDevice, rnnDescriptors []MPSRNNDescriptor) MPSRNNImageInferenceLayer {
	rv := objc.Send[MPSRNNImageInferenceLayer](r.ID, objc.Sel("initWithDevice:rnnDescriptors:"), device, objectivec.IObjectSliceToNSArray(rnnDescriptors))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/encodeBidirectionalSequence(commandBuffer:sourceSequence:destinationForwardImages:destinationBackwardImages:)
func (r MPSRNNImageInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer metal.MTLCommandBuffer, sourceSequence []MPSImage, destinationForwardImages []MPSImage, destinationBackwardImages []MPSImage) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardImages:destinationBackwardImages:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceSequence), objectivec.IObjectSliceToNSArray(destinationForwardImages), objectivec.IObjectSliceToNSArray(destinationBackwardImages))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/encodeSequence(commandBuffer:sourceImages:destinationImages:recurrentInputState:recurrentOutputStates:)
func (r MPSRNNImageInferenceLayer) EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, destinationImages []MPSImage, recurrentInputState IMPSRNNRecurrentImageState, recurrentOutputStates foundation.INSArray) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeSequenceToCommandBuffer:sourceImages:destinationImages:recurrentInputState:recurrentOutputStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(destinationImages), recurrentInputState, recurrentOutputStates)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/bidirectionalCombineMode
func (r MPSRNNImageInferenceLayer) BidirectionalCombineMode() MPSRNNBidirectionalCombineMode {
	rv := objc.Send[MPSRNNBidirectionalCombineMode](r.ID, objc.Sel("bidirectionalCombineMode"))
	return MPSRNNBidirectionalCombineMode(rv)
}
func (r MPSRNNImageInferenceLayer) SetBidirectionalCombineMode(value MPSRNNBidirectionalCombineMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setBidirectionalCombineMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/numberOfLayers
func (r MPSRNNImageInferenceLayer) NumberOfLayers() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("numberOfLayers"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/recurrentOutputIsTemporary
func (r MPSRNNImageInferenceLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}
func (r MPSRNNImageInferenceLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/storeAllIntermediateStates
func (r MPSRNNImageInferenceLayer) StoreAllIntermediateStates() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}
func (r MPSRNNImageInferenceLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/inputFeatureChannels
func (r MPSRNNImageInferenceLayer) InputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNImageInferenceLayer/outputFeatureChannels
func (r MPSRNNImageInferenceLayer) OutputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("outputFeatureChannels"))
	return rv
}
