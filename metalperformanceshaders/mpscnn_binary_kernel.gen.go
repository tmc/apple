// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNBinaryKernel] class.
var (
	_MPSCNNBinaryKernelClass     MPSCNNBinaryKernelClass
	_MPSCNNBinaryKernelClassOnce sync.Once
)

func getMPSCNNBinaryKernelClass() MPSCNNBinaryKernelClass {
	_MPSCNNBinaryKernelClassOnce.Do(func() {
		_MPSCNNBinaryKernelClass = MPSCNNBinaryKernelClass{class: objc.GetClass("MPSCNNBinaryKernel")}
	})
	return _MPSCNNBinaryKernelClass
}

// GetMPSCNNBinaryKernelClass returns the class object for MPSCNNBinaryKernel.
func GetMPSCNNBinaryKernelClass() MPSCNNBinaryKernelClass {
	return getMPSCNNBinaryKernelClass()
}

type MPSCNNBinaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBinaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryKernelClass) Alloc() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A convolution neural network kernel.
//
// # Instance Properties
//
//   - [MPSCNNBinaryKernel.ClipRect]
//   - [MPSCNNBinaryKernel.SetClipRect]
//   - [MPSCNNBinaryKernel.DestinationFeatureChannelOffset]
//   - [MPSCNNBinaryKernel.SetDestinationFeatureChannelOffset]
//   - [MPSCNNBinaryKernel.DestinationImageAllocator]
//   - [MPSCNNBinaryKernel.SetDestinationImageAllocator]
//   - [MPSCNNBinaryKernel.IsBackwards]
//   - [MPSCNNBinaryKernel.Padding]
//   - [MPSCNNBinaryKernel.SetPadding]
//   - [MPSCNNBinaryKernel.PrimaryEdgeMode]
//   - [MPSCNNBinaryKernel.SetPrimaryEdgeMode]
//   - [MPSCNNBinaryKernel.PrimaryOffset]
//   - [MPSCNNBinaryKernel.SetPrimaryOffset]
//   - [MPSCNNBinaryKernel.PrimaryStrideInPixelsX]
//   - [MPSCNNBinaryKernel.SetPrimaryStrideInPixelsX]
//   - [MPSCNNBinaryKernel.PrimaryStrideInPixelsY]
//   - [MPSCNNBinaryKernel.SetPrimaryStrideInPixelsY]
//   - [MPSCNNBinaryKernel.SecondaryEdgeMode]
//   - [MPSCNNBinaryKernel.SetSecondaryEdgeMode]
//   - [MPSCNNBinaryKernel.SecondaryOffset]
//   - [MPSCNNBinaryKernel.SetSecondaryOffset]
//   - [MPSCNNBinaryKernel.SecondaryStrideInPixelsX]
//   - [MPSCNNBinaryKernel.SetSecondaryStrideInPixelsX]
//   - [MPSCNNBinaryKernel.SecondaryStrideInPixelsY]
//   - [MPSCNNBinaryKernel.SetSecondaryStrideInPixelsY]
//   - [MPSCNNBinaryKernel.IsStateModified]
//   - [MPSCNNBinaryKernel.PrimaryDilationRateX]
//   - [MPSCNNBinaryKernel.PrimaryDilationRateY]
//   - [MPSCNNBinaryKernel.PrimaryKernelHeight]
//   - [MPSCNNBinaryKernel.PrimaryKernelWidth]
//   - [MPSCNNBinaryKernel.PrimarySourceFeatureChannelMaxCount]
//   - [MPSCNNBinaryKernel.SetPrimarySourceFeatureChannelMaxCount]
//   - [MPSCNNBinaryKernel.PrimarySourceFeatureChannelOffset]
//   - [MPSCNNBinaryKernel.SetPrimarySourceFeatureChannelOffset]
//   - [MPSCNNBinaryKernel.SecondaryDilationRateX]
//   - [MPSCNNBinaryKernel.SecondaryDilationRateY]
//   - [MPSCNNBinaryKernel.SecondaryKernelHeight]
//   - [MPSCNNBinaryKernel.SecondaryKernelWidth]
//   - [MPSCNNBinaryKernel.SecondarySourceFeatureChannelMaxCount]
//   - [MPSCNNBinaryKernel.SetSecondarySourceFeatureChannelMaxCount]
//   - [MPSCNNBinaryKernel.SecondarySourceFeatureChannelOffset]
//   - [MPSCNNBinaryKernel.SetSecondarySourceFeatureChannelOffset]
//
// # Instance Methods
//
//   - [MPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImage]
//   - [MPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage]
//   - [MPSCNNBinaryKernel.AppendBatchBarrier]
//   - [MPSCNNBinaryKernel.BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [MPSCNNBinaryKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary]
//   - [MPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImages]
//   - [MPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages]
//   - [MPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary]
//   - [MPSCNNBinaryKernel.EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [MPSCNNBinaryKernel.IsResultStateReusedAcrossBatch]
//   - [MPSCNNBinaryKernel.ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [MPSCNNBinaryKernel.ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [MPSCNNBinaryKernel.TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [MPSCNNBinaryKernel.TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel
type MPSCNNBinaryKernel struct {
	MPSKernel
}

// MPSCNNBinaryKernelFromID constructs a [MPSCNNBinaryKernel] from an objc.ID.
//
// A convolution neural network kernel.
func MPSCNNBinaryKernelFromID(id objc.ID) MPSCNNBinaryKernel {
	return MPSCNNBinaryKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSCNNBinaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBinaryKernel] class.
//
// # Instance Properties
//
//   - [IMPSCNNBinaryKernel.ClipRect]
//   - [IMPSCNNBinaryKernel.SetClipRect]
//   - [IMPSCNNBinaryKernel.DestinationFeatureChannelOffset]
//   - [IMPSCNNBinaryKernel.SetDestinationFeatureChannelOffset]
//   - [IMPSCNNBinaryKernel.DestinationImageAllocator]
//   - [IMPSCNNBinaryKernel.SetDestinationImageAllocator]
//   - [IMPSCNNBinaryKernel.IsBackwards]
//   - [IMPSCNNBinaryKernel.Padding]
//   - [IMPSCNNBinaryKernel.SetPadding]
//   - [IMPSCNNBinaryKernel.PrimaryEdgeMode]
//   - [IMPSCNNBinaryKernel.SetPrimaryEdgeMode]
//   - [IMPSCNNBinaryKernel.PrimaryOffset]
//   - [IMPSCNNBinaryKernel.SetPrimaryOffset]
//   - [IMPSCNNBinaryKernel.PrimaryStrideInPixelsX]
//   - [IMPSCNNBinaryKernel.SetPrimaryStrideInPixelsX]
//   - [IMPSCNNBinaryKernel.PrimaryStrideInPixelsY]
//   - [IMPSCNNBinaryKernel.SetPrimaryStrideInPixelsY]
//   - [IMPSCNNBinaryKernel.SecondaryEdgeMode]
//   - [IMPSCNNBinaryKernel.SetSecondaryEdgeMode]
//   - [IMPSCNNBinaryKernel.SecondaryOffset]
//   - [IMPSCNNBinaryKernel.SetSecondaryOffset]
//   - [IMPSCNNBinaryKernel.SecondaryStrideInPixelsX]
//   - [IMPSCNNBinaryKernel.SetSecondaryStrideInPixelsX]
//   - [IMPSCNNBinaryKernel.SecondaryStrideInPixelsY]
//   - [IMPSCNNBinaryKernel.SetSecondaryStrideInPixelsY]
//   - [IMPSCNNBinaryKernel.IsStateModified]
//   - [IMPSCNNBinaryKernel.PrimaryDilationRateX]
//   - [IMPSCNNBinaryKernel.PrimaryDilationRateY]
//   - [IMPSCNNBinaryKernel.PrimaryKernelHeight]
//   - [IMPSCNNBinaryKernel.PrimaryKernelWidth]
//   - [IMPSCNNBinaryKernel.PrimarySourceFeatureChannelMaxCount]
//   - [IMPSCNNBinaryKernel.SetPrimarySourceFeatureChannelMaxCount]
//   - [IMPSCNNBinaryKernel.PrimarySourceFeatureChannelOffset]
//   - [IMPSCNNBinaryKernel.SetPrimarySourceFeatureChannelOffset]
//   - [IMPSCNNBinaryKernel.SecondaryDilationRateX]
//   - [IMPSCNNBinaryKernel.SecondaryDilationRateY]
//   - [IMPSCNNBinaryKernel.SecondaryKernelHeight]
//   - [IMPSCNNBinaryKernel.SecondaryKernelWidth]
//   - [IMPSCNNBinaryKernel.SecondarySourceFeatureChannelMaxCount]
//   - [IMPSCNNBinaryKernel.SetSecondarySourceFeatureChannelMaxCount]
//   - [IMPSCNNBinaryKernel.SecondarySourceFeatureChannelOffset]
//   - [IMPSCNNBinaryKernel.SetSecondarySourceFeatureChannelOffset]
//
// # Instance Methods
//
//   - [IMPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImage]
//   - [IMPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage]
//   - [IMPSCNNBinaryKernel.AppendBatchBarrier]
//   - [IMPSCNNBinaryKernel.BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [IMPSCNNBinaryKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSCNNBinaryKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary]
//   - [IMPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImages]
//   - [IMPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages]
//   - [IMPSCNNBinaryKernel.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary]
//   - [IMPSCNNBinaryKernel.EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [IMPSCNNBinaryKernel.IsResultStateReusedAcrossBatch]
//   - [IMPSCNNBinaryKernel.ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [IMPSCNNBinaryKernel.ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [IMPSCNNBinaryKernel.TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage]
//   - [IMPSCNNBinaryKernel.TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel
type IMPSCNNBinaryKernel interface {
	IMPSKernel

	// Topic: Instance Properties

	ClipRect() metal.MTLRegion
	SetClipRect(value metal.MTLRegion)
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
	DestinationImageAllocator() MPSImageAllocator
	SetDestinationImageAllocator(value MPSImageAllocator)
	IsBackwards() bool
	Padding() MPSNNPadding
	SetPadding(value MPSNNPadding)
	PrimaryEdgeMode() MPSImageEdgeMode
	SetPrimaryEdgeMode(value MPSImageEdgeMode)
	PrimaryOffset() MPSOffset
	SetPrimaryOffset(value MPSOffset)
	PrimaryStrideInPixelsX() uint
	SetPrimaryStrideInPixelsX(value uint)
	PrimaryStrideInPixelsY() uint
	SetPrimaryStrideInPixelsY(value uint)
	SecondaryEdgeMode() MPSImageEdgeMode
	SetSecondaryEdgeMode(value MPSImageEdgeMode)
	SecondaryOffset() MPSOffset
	SetSecondaryOffset(value MPSOffset)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
	IsStateModified() bool
	PrimaryDilationRateX() uint
	PrimaryDilationRateY() uint
	PrimaryKernelHeight() uint
	PrimaryKernelWidth() uint
	PrimarySourceFeatureChannelMaxCount() uint
	SetPrimarySourceFeatureChannelMaxCount(value uint)
	PrimarySourceFeatureChannelOffset() uint
	SetPrimarySourceFeatureChannelOffset(value uint)
	SecondaryDilationRateX() uint
	SecondaryDilationRateY() uint
	SecondaryKernelHeight() uint
	SecondaryKernelWidth() uint
	SecondarySourceFeatureChannelMaxCount() uint
	SetSecondarySourceFeatureChannelMaxCount(value uint)
	SecondarySourceFeatureChannelOffset() uint
	SetSecondarySourceFeatureChannelOffset(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferPrimaryImageSecondaryImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage) IMPSImage
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationImage IMPSImage)
	AppendBatchBarrier() bool
	BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) uint
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, outState IMPSState, isTemporary bool) IMPSImage
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImages(commandBuffer metal.MTLCommandBuffer, primaryImage MPSImageBatch, secondaryImage MPSImageBatch) MPSImageBatch
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, destinationImages MPSImageBatch)
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, outState MPSStateBatch, isTemporary bool) MPSImageBatch
	EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) uint
	IsResultStateReusedAcrossBatch() bool
	ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
	TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
}

// Init initializes the instance.
func (c MPSCNNBinaryKernel) Init() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryKernel) Autorelease() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryKernel creates a new MPSCNNBinaryKernel instance.
func NewMPSCNNBinaryKernel() MPSCNNBinaryKernel {
	class := getMPSCNNBinaryKernelClass()
	rv := objc.Send[MPSCNNBinaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBinaryKernelWithCoder(aDecoder foundation.INSCoder) MPSCNNBinaryKernel {
	instance := getMPSCNNBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNBinaryKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBinaryKernel {
	instance := getMPSCNNBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(device:)
func NewCNNBinaryKernelWithDevice(device metal.MTLDevice) MPSCNNBinaryKernel {
	instance := getMPSCNNBinaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBinaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encode(commandBuffer:primaryImage:secondaryImage:)
func (c MPSCNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:"), commandBuffer, primaryImage, secondaryImage)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encode(commandBuffer:primaryImage:secondaryImage:destinationImage:)
func (c MPSCNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/appendBatchBarrier()
func (c MPSCNNBinaryKernel) AppendBatchBarrier() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("appendBatchBarrier"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/batchEncodingStorageSize(primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("batchEncodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/destinationImageDescriptor(forSourceImages:sourceStates:)
func (c MPSCNNBinaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates))
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encode(commandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:)
func (c MPSCNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, outState IMPSState, isTemporary bool) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:"), commandBuffer, primaryImage, secondaryImage, outState, isTemporary)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encodeBatch(commandBuffer:primaryImages:secondaryImages:)
func (c MPSCNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImages(commandBuffer metal.MTLCommandBuffer, primaryImage MPSImageBatch, secondaryImage MPSImageBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:"), commandBuffer, primaryImage, secondaryImage)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encodeBatch(commandBuffer:primaryImages:secondaryImages:destinationImages:)
func (c MPSCNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationImages:"), commandBuffer, primaryImages, secondaryImages, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encodeBatch(commandBuffer:primaryImages:secondaryImages:destinationStates:destinationStateIsTemporary:)
func (c MPSCNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, primaryImages MPSImageBatch, secondaryImages MPSImageBatch, outState MPSStateBatch, isTemporary bool) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, primaryImages, secondaryImages, outState, isTemporary)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/encodingStorageSize(primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("encodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/isResultStateReusedAcrossBatch()
func (c MPSCNNBinaryKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/resultState(primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/resultStateBatch(primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("resultStateBatchForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/temporaryResultState(commandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/temporaryResultStateBatch(commandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:)
func (c MPSCNNBinaryKernel) TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage MPSImageBatch, secondaryImage MPSImageBatch, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/clipRect
func (c MPSCNNBinaryKernel) ClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](c.ID, objc.Sel("clipRect"))
	return metal.MTLRegion(rv)
}
func (c MPSCNNBinaryKernel) SetClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](c.ID, objc.Sel("setClipRect:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/destinationFeatureChannelOffset
func (c MPSCNNBinaryKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}
func (c MPSCNNBinaryKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/destinationImageAllocator
func (c MPSCNNBinaryKernel) DestinationImageAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}
func (c MPSCNNBinaryKernel) SetDestinationImageAllocator(value MPSImageAllocator) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/isBackwards
func (c MPSCNNBinaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBackwards"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/padding
func (c MPSCNNBinaryKernel) Padding() MPSNNPadding {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("padding"))
	return MPSNNPaddingObjectFromID(rv)
}
func (c MPSCNNBinaryKernel) SetPadding(value MPSNNPadding) {
	objc.Send[struct{}](c.ID, objc.Sel("setPadding:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryEdgeMode
func (c MPSCNNBinaryKernel) PrimaryEdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](c.ID, objc.Sel("primaryEdgeMode"))
	return MPSImageEdgeMode(rv)
}
func (c MPSCNNBinaryKernel) SetPrimaryEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryOffset
func (c MPSCNNBinaryKernel) PrimaryOffset() MPSOffset {
	rv := objc.Send[MPSOffset](c.ID, objc.Sel("primaryOffset"))
	return MPSOffset(rv)
}
func (c MPSCNNBinaryKernel) SetPrimaryOffset(value MPSOffset) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryStrideInPixelsX
func (c MPSCNNBinaryKernel) PrimaryStrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}
func (c MPSCNNBinaryKernel) SetPrimaryStrideInPixelsX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryStrideInPixelsY
func (c MPSCNNBinaryKernel) PrimaryStrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}
func (c MPSCNNBinaryKernel) SetPrimaryStrideInPixelsY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryEdgeMode
func (c MPSCNNBinaryKernel) SecondaryEdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](c.ID, objc.Sel("secondaryEdgeMode"))
	return MPSImageEdgeMode(rv)
}
func (c MPSCNNBinaryKernel) SetSecondaryEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryOffset
func (c MPSCNNBinaryKernel) SecondaryOffset() MPSOffset {
	rv := objc.Send[MPSOffset](c.ID, objc.Sel("secondaryOffset"))
	return MPSOffset(rv)
}
func (c MPSCNNBinaryKernel) SetSecondaryOffset(value MPSOffset) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryStrideInPixelsX
func (c MPSCNNBinaryKernel) SecondaryStrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}
func (c MPSCNNBinaryKernel) SetSecondaryStrideInPixelsX(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryStrideInPixelsY
func (c MPSCNNBinaryKernel) SecondaryStrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}
func (c MPSCNNBinaryKernel) SetSecondaryStrideInPixelsY(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/isStateModified
func (c MPSCNNBinaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStateModified"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryDilationRateX
func (c MPSCNNBinaryKernel) PrimaryDilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryDilationRateX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryDilationRateY
func (c MPSCNNBinaryKernel) PrimaryDilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryDilationRateY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryKernelHeight
func (c MPSCNNBinaryKernel) PrimaryKernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryKernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primaryKernelWidth
func (c MPSCNNBinaryKernel) PrimaryKernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primaryKernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primarySourceFeatureChannelMaxCount
func (c MPSCNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}
func (c MPSCNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/primarySourceFeatureChannelOffset
func (c MPSCNNBinaryKernel) PrimarySourceFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}
func (c MPSCNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryDilationRateX
func (c MPSCNNBinaryKernel) SecondaryDilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryDilationRateX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryDilationRateY
func (c MPSCNNBinaryKernel) SecondaryDilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryDilationRateY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryKernelHeight
func (c MPSCNNBinaryKernel) SecondaryKernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryKernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondaryKernelWidth
func (c MPSCNNBinaryKernel) SecondaryKernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondaryKernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondarySourceFeatureChannelMaxCount
func (c MPSCNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}
func (c MPSCNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/secondarySourceFeatureChannelOffset
func (c MPSCNNBinaryKernel) SecondarySourceFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}
func (c MPSCNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}
