// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNKernel] class.
var (
	_MPSCNNKernelClass     MPSCNNKernelClass
	_MPSCNNKernelClassOnce sync.Once
)

func getMPSCNNKernelClass() MPSCNNKernelClass {
	_MPSCNNKernelClassOnce.Do(func() {
		_MPSCNNKernelClass = MPSCNNKernelClass{class: objc.GetClass("MPSCNNKernel")}
	})
	return _MPSCNNKernelClass
}

// GetMPSCNNKernelClass returns the class object for MPSCNNKernel.
func GetMPSCNNKernelClass() MPSCNNKernelClass {
	return getMPSCNNKernelClass()
}

type MPSCNNKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNKernelClass) Alloc() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Base class for neural network layers.
//
// # Overview
//
// An [MPSCNNKernel] object consumes one [MPSImage] object and produces one
// [MPSImage] object.
//
// The region overwritten in the destination image is described by the
// [MPSCNNKernel.ClipRect] property. The top left corner of the region
// consumed (ignoring adjustments for filter size—for example, convolution
// filter size) is given by the [MPSCNNKernel.Offset] property. The size of
// the region consumed is a function of the size of the
// [MPSCNNKernel.ClipRect] property and any subsampling caused by pixel
// strides at work (for example,
// [MPSCNNKernel.StrideInPixelsX]/[MPSCNNKernel.StrideInPixelsY] in the
// [MPSCNNPooling] class). Wherever the [MPSCNNKernel.Offset] and
// [MPSCNNKernel.ClipRect] properties would cause an `{x,y}` pixel address not
// in the image to be read, the [MPSCNNKernel.EdgeMode] property is used to
// determine what value to read there.
//
// The `z` or `depth` component of the [MPSCNNKernel.Offset], [origin] and
// [size] properties indexes which images to use.
//
// - If the [MPSImage] object contains only a single image, then these values
// should be `offset.Z() = 0`, `clipRect.OriginXCUIElementTypeZ() = 0`, and
// `clipRect.SizeXCUIElementTypeDepth() = 1`. - If the [MPSImage] object
// contains multiple images, then the value of
// `clipRect.SizeXCUIElementTypeDepth()` determines the number of images to
// process. Both the source and destination [MPSImage] objects must have at
// least this many images. The value of `offset.Z()` refers to the starting
// image index of the source. Thus, the value of `offset.Z() +
// clipRect.SizeXCUIElementTypeDepth()` must be `<=source.NumberOfImages()`.
// Similarly, the value of `clipRect.OriginXCUIElementTypeZ()` determines the
// starting image index of the destination. Thus, the value of
// `clipRect.OriginXCUIElementTypeZ() + clipRect.SizeXCUIElementTypeDepth()`
// must be `<=destination.NumberOfImages()`.
//
// The [MPSCNNKernel.DestinationFeatureChannelOffset] property can be used to
// control where the kernel will start writing in terms of feature channel
// dimension. For example, if the destination has 64 channels and
// th[MPSCNNKernel.DestinationFeatureChannelOffset]e kernel outputs 32
// channels, channels 0-31 of the destination will be populated by the kernel
// (by default). But if you want the kernel to populate channels 32-63 of the
// destination, you can set the value of
// [MPSCNNKernel.DestinationFeatureChannelOffset] to 32. Suppose you have a
// source of dimensions `w x h x Ni`, where [N] is the number of channels,
// which goes through a convolution filter [C0] which produces the output `O0
// = w x h x N0` and [C]1 which produces the output `O1 = w x h x N1` followed
// by concatenation which produces `O = w x h x (N0 + N1)`. You can achieve
// this by creating an [MPSImage] object with dimensions `w x h x (N0 + N1)`
// and using this as the destination of both convolutions as follows:
//
// - `C0: destinationFeatureChannelOffset = 0`, this will output [N0] channels
// starting at channel `0` of destination thus populating `[0,N0-1]` channels.
// - `C1: destinationFeatureChannelOffset = N0`, this will output [N1]
// channels starting at channel [N0] of destination thus populating
// `[N0,N0+N1-1]` channels.
//
// # Instance Properties
//
//   - [MPSCNNKernel.Offset]: The position of the destination image’s clip rectangle origin, relative to the source image.
//   - [MPSCNNKernel.SetOffset]
//   - [MPSCNNKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//   - [MPSCNNKernel.SetClipRect]
//   - [MPSCNNKernel.DestinationFeatureChannelOffset]: The number of channels in the destination image to skip before writing output data.
//   - [MPSCNNKernel.SetDestinationFeatureChannelOffset]
//   - [MPSCNNKernel.EdgeMode]: The edge mode to use when texture reads stray off the edge of an image.
//   - [MPSCNNKernel.SetEdgeMode]
//   - [MPSCNNKernel.KernelHeight]
//   - [MPSCNNKernel.KernelWidth]
//   - [MPSCNNKernel.StrideInPixelsX]
//   - [MPSCNNKernel.StrideInPixelsY]
//   - [MPSCNNKernel.IsBackwards]
//   - [MPSCNNKernel.Padding]
//   - [MPSCNNKernel.SetPadding]
//   - [MPSCNNKernel.DestinationImageAllocator]
//   - [MPSCNNKernel.SetDestinationImageAllocator]
//   - [MPSCNNKernel.DilationRateX]
//   - [MPSCNNKernel.DilationRateY]
//   - [MPSCNNKernel.IsStateModified]
//   - [MPSCNNKernel.SourceFeatureChannelMaxCount]
//   - [MPSCNNKernel.SetSourceFeatureChannelMaxCount]
//   - [MPSCNNKernel.SourceFeatureChannelOffset]
//   - [MPSCNNKernel.SetSourceFeatureChannelOffset]
//
// # Instance Methods
//
//   - [MPSCNNKernel.EncodeToCommandBufferSourceImage]
//   - [MPSCNNKernel.EncodeToCommandBufferSourceImageDestinationImage]: Encodes a kernel into a command buffer.  The ensuing operation proceeds out-of-place.
//   - [MPSCNNKernel.AppendBatchBarrier]
//   - [MPSCNNKernel.BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage]
//   - [MPSCNNKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSCNNKernel.EncodeToCommandBufferSourceImageDestinationStateDestinationImage]
//   - [MPSCNNKernel.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary]
//   - [MPSCNNKernel.EncodeBatchToCommandBufferSourceImages]
//   - [MPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationImages]
//   - [MPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages]
//   - [MPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary]
//   - [MPSCNNKernel.EncodingStorageSizeForSourceImageSourceStatesDestinationImage]
//   - [MPSCNNKernel.IsResultStateReusedAcrossBatch]
//   - [MPSCNNKernel.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSCNNKernel.ResultStateBatchForSourceImageSourceStatesDestinationImage]
//   - [MPSCNNKernel.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSCNNKernel.TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel
//
// [origin]: https://developer.apple.com/documentation/Metal/MTLRegion/origin
// [size]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRegion/size
type MPSCNNKernel struct {
	MPSKernel
}

// MPSCNNKernelFromID constructs a [MPSCNNKernel] from an objc.ID.
//
// Base class for neural network layers.
func MPSCNNKernelFromID(id objc.ID) MPSCNNKernel {
	return MPSCNNKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSCNNKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNKernel] class.
//
// # Instance Properties
//
//   - [IMPSCNNKernel.Offset]: The position of the destination image’s clip rectangle origin, relative to the source image.
//   - [IMPSCNNKernel.SetOffset]
//   - [IMPSCNNKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//   - [IMPSCNNKernel.SetClipRect]
//   - [IMPSCNNKernel.DestinationFeatureChannelOffset]: The number of channels in the destination image to skip before writing output data.
//   - [IMPSCNNKernel.SetDestinationFeatureChannelOffset]
//   - [IMPSCNNKernel.EdgeMode]: The edge mode to use when texture reads stray off the edge of an image.
//   - [IMPSCNNKernel.SetEdgeMode]
//   - [IMPSCNNKernel.KernelHeight]
//   - [IMPSCNNKernel.KernelWidth]
//   - [IMPSCNNKernel.StrideInPixelsX]
//   - [IMPSCNNKernel.StrideInPixelsY]
//   - [IMPSCNNKernel.IsBackwards]
//   - [IMPSCNNKernel.Padding]
//   - [IMPSCNNKernel.SetPadding]
//   - [IMPSCNNKernel.DestinationImageAllocator]
//   - [IMPSCNNKernel.SetDestinationImageAllocator]
//   - [IMPSCNNKernel.DilationRateX]
//   - [IMPSCNNKernel.DilationRateY]
//   - [IMPSCNNKernel.IsStateModified]
//   - [IMPSCNNKernel.SourceFeatureChannelMaxCount]
//   - [IMPSCNNKernel.SetSourceFeatureChannelMaxCount]
//   - [IMPSCNNKernel.SourceFeatureChannelOffset]
//   - [IMPSCNNKernel.SetSourceFeatureChannelOffset]
//
// # Instance Methods
//
//   - [IMPSCNNKernel.EncodeToCommandBufferSourceImage]
//   - [IMPSCNNKernel.EncodeToCommandBufferSourceImageDestinationImage]: Encodes a kernel into a command buffer.  The ensuing operation proceeds out-of-place.
//   - [IMPSCNNKernel.AppendBatchBarrier]
//   - [IMPSCNNKernel.BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage]
//   - [IMPSCNNKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSCNNKernel.EncodeToCommandBufferSourceImageDestinationStateDestinationImage]
//   - [IMPSCNNKernel.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary]
//   - [IMPSCNNKernel.EncodeBatchToCommandBufferSourceImages]
//   - [IMPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationImages]
//   - [IMPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages]
//   - [IMPSCNNKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary]
//   - [IMPSCNNKernel.EncodingStorageSizeForSourceImageSourceStatesDestinationImage]
//   - [IMPSCNNKernel.IsResultStateReusedAcrossBatch]
//   - [IMPSCNNKernel.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSCNNKernel.ResultStateBatchForSourceImageSourceStatesDestinationImage]
//   - [IMPSCNNKernel.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSCNNKernel.TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel
type IMPSCNNKernel interface {
	IMPSKernel

	// Topic: Instance Properties

	// The position of the destination image’s clip rectangle origin, relative to the source image.
	Offset() MPSOffset
	SetOffset(value MPSOffset)
	// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
	ClipRect() metal.MTLRegion
	SetClipRect(value metal.MTLRegion)
	// The number of channels in the destination image to skip before writing output data.
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
	// The edge mode to use when texture reads stray off the edge of an image.
	EdgeMode() MPSImageEdgeMode
	SetEdgeMode(value MPSImageEdgeMode)
	KernelHeight() uint
	KernelWidth() uint
	StrideInPixelsX() uint
	StrideInPixelsY() uint
	IsBackwards() bool
	Padding() MPSNNPadding
	SetPadding(value MPSNNPadding)
	DestinationImageAllocator() MPSImageAllocator
	SetDestinationImageAllocator(value MPSImageAllocator)
	DilationRateX() uint
	DilationRateY() uint
	IsStateModified() bool
	SourceFeatureChannelMaxCount() uint
	SetSourceFeatureChannelMaxCount(value uint)
	SourceFeatureChannelOffset() uint
	SetSourceFeatureChannelOffset(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage) IMPSImage
	// Encodes a kernel into a command buffer.  The ensuing operation proceeds out-of-place.
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationImage IMPSImage)
	AppendBatchBarrier() bool
	BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) uint
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor
	EncodeToCommandBufferSourceImageDestinationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationState IMPSState, destinationImage IMPSImage)
	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, outState IMPSState, isTemporary bool) IMPSImage
	EncodeBatchToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationImages MPSImageBatch)
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationStates MPSStateBatch, destinationImages MPSImageBatch)
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, outStates MPSStateBatch, isTemporary bool) MPSImageBatch
	EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) uint
	IsResultStateReusedAcrossBatch() bool
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
}

// Init initializes the instance.
func (c MPSCNNKernel) Init() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNKernel) Autorelease() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNKernel creates a new MPSCNNKernel instance.
func NewMPSCNNKernel() MPSCNNKernel {
	class := getMPSCNNKernelClass()
	rv := objc.Send[MPSCNNKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNKernelWithCoder(aDecoder foundation.INSCoder) MPSCNNKernel {
	instance := getMPSCNNKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNKernel {
	instance := getMPSCNNKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNKernelWithDevice(device metal.MTLDevice) MPSCNNKernel {
	instance := getMPSCNNKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encode(commandBuffer:sourceImage:)
func (c MPSCNNKernel) EncodeToCommandBufferSourceImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:"), commandBuffer, sourceImage)
	return MPSImageFromID(rv)
}

// Encodes a kernel into a command buffer. The ensuing operation proceeds
// out-of-place.
//
// commandBuffer: A valid command buffer to receive the encoded filter.
//
// sourceImage: A valid source image.
//
// destinationImage: A valid destination image to be overwritten by the results.
//
// # Discussion
//
// The `destinationImage` object may not alias the `sourceImage` object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encode(commandBuffer:sourceImage:destinationImage:)
func (c MPSCNNKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/appendBatchBarrier()
func (c MPSCNNKernel) AppendBatchBarrier() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("appendBatchBarrier"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/batchEncodingStorageSize(sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("batchEncodingStorageSizeForSourceImage:sourceStates:destinationImage:"), sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/destinationImageDescriptor(sourceImages:sourceStates:)
func (c MPSCNNKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates))
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encode(commandBuffer:sourceImage:destinationState:destinationImage:)
func (c MPSCNNKernel) EncodeToCommandBufferSourceImageDestinationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationState IMPSState, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationImage:"), commandBuffer, sourceImage, destinationState, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encode(commandBuffer:sourceImage:destinationState:destinationStateIsTemporary:)
func (c MPSCNNKernel) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, outState IMPSState, isTemporary bool) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImage, outState, isTemporary)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encodeBatch(commandBuffer:sourceImages:)
func (c MPSCNNKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), commandBuffer, sourceImages)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encodeBatch(commandBuffer:sourceImages:destinationImages:)
func (c MPSCNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationImages:"), commandBuffer, sourceImages, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encodeBatch(commandBuffer:sourceImages:destinationStates:destinationImages:)
func (c MPSCNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, destinationStates MPSStateBatch, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationImages:"), commandBuffer, sourceImages, destinationStates, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encodeBatch(commandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:)
func (c MPSCNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, outStates MPSStateBatch, isTemporary bool) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, outStates, isTemporary)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/encodingStorageSize(sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("encodingStorageSizeForSourceImage:sourceStates:destinationImage:"), sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/isResultStateReusedAcrossBatch()
func (c MPSCNNKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/resultState(sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/resultStateBatch(sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/temporaryResultState(commandBuffer:sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/temporaryResultStateBatch(commandBuffer:sourceImage:sourceStates:destinationImage:)
func (c MPSCNNKernel) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// The position of the destination image’s clip rectangle origin, relative
// to the source image.
//
// # Discussion
//
// The offset is defined as the position of `clipRect.Origin()` in source
// image coordinates. The default value is `{0,0,0}`, indicating that the top
// left corners of the clip rectangle and the source image align.
//
// The value of `offset.Z()` is the index of the starting source image in
// batch processing mode.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/offset
func (c MPSCNNKernel) Offset() MPSOffset {
	rv := objc.Send[MPSOffset](c.ID, objc.Sel("offset"))
	return MPSOffset(rv)
}
func (c MPSCNNKernel) SetOffset(value MPSOffset) {
	objc.Send[struct{}](c.ID, objc.Sel("setOffset:"), value)
}

// An optional clip rectangle to use when writing data. Only the pixels in the
// clip rectangle will be overwritten.
//
// # Discussion
//
// A region that indicates which part of the destination image to overwrite.
// If the clip rectangle does not lie completely within the destination image,
// the intersection between the clip rectangle and the destination image
// bounds is used instead.
//
// The default value is [MPSRectNoClip], indicating that the entire
// destination image will be overwritten.
//
// The value of `clipRect.OriginXCUIElementTypeZ()` is the index of the
// starting destination image in batch processing mode. The value of
// `clipRect.SizeXCUIElementTypeDepth()` is the number of images to process in
// batch processing mode.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/clipRect
//
// [MPSRectNoClip]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRectNoClip
func (c MPSCNNKernel) ClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](c.ID, objc.Sel("clipRect"))
	return metal.MTLRegion(rv)
}
func (c MPSCNNKernel) SetClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](c.ID, objc.Sel("setClipRect:"), value)
}

// The number of channels in the destination image to skip before writing
// output data.
//
// # Discussion
//
// This is the starting offset in the destination image in the feature channel
// dimension at which destination output data is written. This allows you to
// pass a subset of all the channels in an image as the output of a kernel.
//
// For example, suppose a destination image has 24 channels and a kernel
// outputs 8 channels. If we want channels 8 to 15 of this destination image
// to be used for the output, we can set the value of the
// [MPSCNNKernel.DestinationFeatureChannelOffset] property to 8.
//
// Note that this offset applies independently to each image when the
// [MPSImage] object is a container for multiple images and the [MPSCNNKernel]
// object is processing multiple images (i.e.,
// `clipRect.SizeXCUIElementTypeDepth() > 1`).
//
// The default value is `0`. Any other value specified must be a multiple of
// `4`. If the kernel outputs [N] channels, the destination image must have at
// least `destinationFeatureChannelOffset + N` channels. Using a destination
// image with an insufficient number of feature channels results in an error.
//
// For example, if a convolution filter outputs 32 channels, and the
// destination image has 64 channels, then it is an error to set
// `destinationFeatureChannelOffset > 32`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/destinationFeatureChannelOffset
func (c MPSCNNKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}
func (c MPSCNNKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

// The edge mode to use when texture reads stray off the edge of an image.
//
// # Discussion
//
// Most [MPSKernel] objects can read off the edge of the source image. This
// can happen because of a negative offset property, because the value of
// `offset + clipRect.Size()` is larger than the source image or because the
// filter looks at neighboring pixels, such as a convolution filter.
//
// The default value is [MPSImageEdgeModeZero].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/edgeMode
func (c MPSCNNKernel) EdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](c.ID, objc.Sel("edgeMode"))
	return MPSImageEdgeMode(rv)
}
func (c MPSCNNKernel) SetEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setEdgeMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/kernelHeight
func (c MPSCNNKernel) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/kernelWidth
func (c MPSCNNKernel) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/strideInPixelsX
func (c MPSCNNKernel) StrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/strideInPixelsY
func (c MPSCNNKernel) StrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/isBackwards
func (c MPSCNNKernel) IsBackwards() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBackwards"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/padding
func (c MPSCNNKernel) Padding() MPSNNPadding {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("padding"))
	return MPSNNPaddingObjectFromID(rv)
}
func (c MPSCNNKernel) SetPadding(value MPSNNPadding) {
	objc.Send[struct{}](c.ID, objc.Sel("setPadding:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/destinationImageAllocator
func (c MPSCNNKernel) DestinationImageAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}
func (c MPSCNNKernel) SetDestinationImageAllocator(value MPSImageAllocator) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/dilationRateX
func (c MPSCNNKernel) DilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/dilationRateY
func (c MPSCNNKernel) DilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/isStateModified
func (c MPSCNNKernel) IsStateModified() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStateModified"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/sourceFeatureChannelMaxCount
func (c MPSCNNKernel) SourceFeatureChannelMaxCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceFeatureChannelMaxCount"))
	return rv
}
func (c MPSCNNKernel) SetSourceFeatureChannelMaxCount(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSourceFeatureChannelMaxCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/sourceFeatureChannelOffset
func (c MPSCNNKernel) SourceFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceFeatureChannelOffset"))
	return rv
}
func (c MPSCNNKernel) SetSourceFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSourceFeatureChannelOffset:"), value)
}
