// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNMultiaryKernel] class.
var (
	_MPSCNNMultiaryKernelClass     MPSCNNMultiaryKernelClass
	_MPSCNNMultiaryKernelClassOnce sync.Once
)

func getMPSCNNMultiaryKernelClass() MPSCNNMultiaryKernelClass {
	_MPSCNNMultiaryKernelClassOnce.Do(func() {
		_MPSCNNMultiaryKernelClass = MPSCNNMultiaryKernelClass{class: objc.GetClass("MPSCNNMultiaryKernel")}
	})
	return _MPSCNNMultiaryKernelClass
}

// GetMPSCNNMultiaryKernelClass returns the class object for MPSCNNMultiaryKernel.
func GetMPSCNNMultiaryKernelClass() MPSCNNMultiaryKernelClass {
	return getMPSCNNMultiaryKernelClass()
}

type MPSCNNMultiaryKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNMultiaryKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNMultiaryKernelClass) Alloc() MPSCNNMultiaryKernel {
	rv := objc.Send[MPSCNNMultiaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNMultiaryKernel.InitWithDeviceSourceCount]
//
// # Instance Properties
//
//   - [MPSCNNMultiaryKernel.ClipRect]
//   - [MPSCNNMultiaryKernel.SetClipRect]
//   - [MPSCNNMultiaryKernel.DestinationFeatureChannelOffset]
//   - [MPSCNNMultiaryKernel.SetDestinationFeatureChannelOffset]
//   - [MPSCNNMultiaryKernel.DestinationImageAllocator]
//   - [MPSCNNMultiaryKernel.SetDestinationImageAllocator]
//   - [MPSCNNMultiaryKernel.IsBackwards]
//   - [MPSCNNMultiaryKernel.IsStateModified]
//   - [MPSCNNMultiaryKernel.Padding]
//   - [MPSCNNMultiaryKernel.SetPadding]
//   - [MPSCNNMultiaryKernel.SourceCount]
//
// # Instance Methods
//
//   - [MPSCNNMultiaryKernel.AppendBatchBarrier]
//   - [MPSCNNMultiaryKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSCNNMultiaryKernel.DilationRateXatIndex]
//   - [MPSCNNMultiaryKernel.DilationRateYatIndex]
//   - [MPSCNNMultiaryKernel.EdgeModeAtIndex]
//   - [MPSCNNMultiaryKernel.EncodeToCommandBufferSourceImages]
//   - [MPSCNNMultiaryKernel.EncodeToCommandBufferSourceImagesDestinationImage]
//   - [MPSCNNMultiaryKernel.EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary]
//   - [MPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImages]
//   - [MPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImagesDestinationImages]
//   - [MPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary]
//   - [MPSCNNMultiaryKernel.IsResultStateReusedAcrossBatch]
//   - [MPSCNNMultiaryKernel.KernelHeightAtIndex]
//   - [MPSCNNMultiaryKernel.KernelWidthAtIndex]
//   - [MPSCNNMultiaryKernel.OffsetAtIndex]
//   - [MPSCNNMultiaryKernel.ResultStateForSourceImagesSourceStatesDestinationImage]
//   - [MPSCNNMultiaryKernel.ResultStateBatchForSourceImagesSourceStatesDestinationImage]
//   - [MPSCNNMultiaryKernel.SetDilationRateXAtIndex]
//   - [MPSCNNMultiaryKernel.SetDilationRateYAtIndex]
//   - [MPSCNNMultiaryKernel.SetEdgeModeAtIndex]
//   - [MPSCNNMultiaryKernel.SetKernelHeightAtIndex]
//   - [MPSCNNMultiaryKernel.SetKernelWidthAtIndex]
//   - [MPSCNNMultiaryKernel.SetOffsetAtIndex]
//   - [MPSCNNMultiaryKernel.SetSourceFeatureChannelMaxCountAtIndex]
//   - [MPSCNNMultiaryKernel.SetSourceFeatureChannelOffsetAtIndex]
//   - [MPSCNNMultiaryKernel.SetStrideInPixelsXAtIndex]
//   - [MPSCNNMultiaryKernel.SetStrideInPixelsYAtIndex]
//   - [MPSCNNMultiaryKernel.SourceFeatureChannelMaxCountAtIndex]
//   - [MPSCNNMultiaryKernel.SourceFeatureChannelOffsetAtIndex]
//   - [MPSCNNMultiaryKernel.StrideInPixelsXatIndex]
//   - [MPSCNNMultiaryKernel.StrideInPixelsYatIndex]
//   - [MPSCNNMultiaryKernel.TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [MPSCNNMultiaryKernel.TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel
type MPSCNNMultiaryKernel struct {
	MPSKernel
}

// MPSCNNMultiaryKernelFromID constructs a [MPSCNNMultiaryKernel] from an objc.ID.
func MPSCNNMultiaryKernelFromID(id objc.ID) MPSCNNMultiaryKernel {
	return MPSCNNMultiaryKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSCNNMultiaryKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNMultiaryKernel] class.
//
// # Initializers
//
//   - [IMPSCNNMultiaryKernel.InitWithDeviceSourceCount]
//
// # Instance Properties
//
//   - [IMPSCNNMultiaryKernel.ClipRect]
//   - [IMPSCNNMultiaryKernel.SetClipRect]
//   - [IMPSCNNMultiaryKernel.DestinationFeatureChannelOffset]
//   - [IMPSCNNMultiaryKernel.SetDestinationFeatureChannelOffset]
//   - [IMPSCNNMultiaryKernel.DestinationImageAllocator]
//   - [IMPSCNNMultiaryKernel.SetDestinationImageAllocator]
//   - [IMPSCNNMultiaryKernel.IsBackwards]
//   - [IMPSCNNMultiaryKernel.IsStateModified]
//   - [IMPSCNNMultiaryKernel.Padding]
//   - [IMPSCNNMultiaryKernel.SetPadding]
//   - [IMPSCNNMultiaryKernel.SourceCount]
//
// # Instance Methods
//
//   - [IMPSCNNMultiaryKernel.AppendBatchBarrier]
//   - [IMPSCNNMultiaryKernel.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSCNNMultiaryKernel.DilationRateXatIndex]
//   - [IMPSCNNMultiaryKernel.DilationRateYatIndex]
//   - [IMPSCNNMultiaryKernel.EdgeModeAtIndex]
//   - [IMPSCNNMultiaryKernel.EncodeToCommandBufferSourceImages]
//   - [IMPSCNNMultiaryKernel.EncodeToCommandBufferSourceImagesDestinationImage]
//   - [IMPSCNNMultiaryKernel.EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary]
//   - [IMPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImages]
//   - [IMPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImagesDestinationImages]
//   - [IMPSCNNMultiaryKernel.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary]
//   - [IMPSCNNMultiaryKernel.IsResultStateReusedAcrossBatch]
//   - [IMPSCNNMultiaryKernel.KernelHeightAtIndex]
//   - [IMPSCNNMultiaryKernel.KernelWidthAtIndex]
//   - [IMPSCNNMultiaryKernel.OffsetAtIndex]
//   - [IMPSCNNMultiaryKernel.ResultStateForSourceImagesSourceStatesDestinationImage]
//   - [IMPSCNNMultiaryKernel.ResultStateBatchForSourceImagesSourceStatesDestinationImage]
//   - [IMPSCNNMultiaryKernel.SetDilationRateXAtIndex]
//   - [IMPSCNNMultiaryKernel.SetDilationRateYAtIndex]
//   - [IMPSCNNMultiaryKernel.SetEdgeModeAtIndex]
//   - [IMPSCNNMultiaryKernel.SetKernelHeightAtIndex]
//   - [IMPSCNNMultiaryKernel.SetKernelWidthAtIndex]
//   - [IMPSCNNMultiaryKernel.SetOffsetAtIndex]
//   - [IMPSCNNMultiaryKernel.SetSourceFeatureChannelMaxCountAtIndex]
//   - [IMPSCNNMultiaryKernel.SetSourceFeatureChannelOffsetAtIndex]
//   - [IMPSCNNMultiaryKernel.SetStrideInPixelsXAtIndex]
//   - [IMPSCNNMultiaryKernel.SetStrideInPixelsYAtIndex]
//   - [IMPSCNNMultiaryKernel.SourceFeatureChannelMaxCountAtIndex]
//   - [IMPSCNNMultiaryKernel.SourceFeatureChannelOffsetAtIndex]
//   - [IMPSCNNMultiaryKernel.StrideInPixelsXatIndex]
//   - [IMPSCNNMultiaryKernel.StrideInPixelsYatIndex]
//   - [IMPSCNNMultiaryKernel.TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [IMPSCNNMultiaryKernel.TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel
type IMPSCNNMultiaryKernel interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceSourceCount(device metal.MTLDevice, sourceCount uint) MPSCNNMultiaryKernel

	// Topic: Instance Properties

	ClipRect() metal.MTLRegion
	SetClipRect(value metal.MTLRegion)
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
	DestinationImageAllocator() MPSImageAllocator
	SetDestinationImageAllocator(value MPSImageAllocator)
	IsBackwards() bool
	IsStateModified() bool
	Padding() MPSNNPadding
	SetPadding(value MPSNNPadding)
	SourceCount() uint

	// Topic: Instance Methods

	AppendBatchBarrier() bool
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor
	DilationRateXatIndex(index uint) uint
	DilationRateYatIndex(index uint) uint
	EdgeModeAtIndex(index uint) MPSImageEdgeMode
	EncodeToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage) IMPSImage
	EncodeToCommandBufferSourceImagesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, destinationImage IMPSImage)
	EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, outState IMPSState, isTemporary bool) IMPSImage
	EncodeBatchToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImageBatches []foundation.NSArray) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, destinationImages MPSImageBatch)
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImageBatches []foundation.NSArray, outState MPSStateBatch, isTemporary bool) MPSImageBatch
	IsResultStateReusedAcrossBatch() bool
	KernelHeightAtIndex(index uint) uint
	KernelWidthAtIndex(index uint) uint
	OffsetAtIndex(index uint) MPSOffset
	ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages []MPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages []foundation.NSArray, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
	SetDilationRateXAtIndex(dilationRate uint, index uint)
	SetDilationRateYAtIndex(dilationRate uint, index uint)
	SetEdgeModeAtIndex(edgeMode MPSImageEdgeMode, index uint)
	SetKernelHeightAtIndex(height uint, index uint)
	SetKernelWidthAtIndex(width uint, index uint)
	SetOffsetAtIndex(offset MPSOffset, index uint)
	SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint)
	SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint)
	SetStrideInPixelsXAtIndex(stride uint, index uint)
	SetStrideInPixelsYAtIndex(stride uint, index uint)
	SourceFeatureChannelMaxCountAtIndex(index uint) uint
	SourceFeatureChannelOffsetAtIndex(index uint) uint
	StrideInPixelsXatIndex(index uint) uint
	StrideInPixelsYatIndex(index uint) uint
	TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage []MPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState
	TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage []foundation.NSArray, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch
}

// Init initializes the instance.
func (c MPSCNNMultiaryKernel) Init() MPSCNNMultiaryKernel {
	rv := objc.Send[MPSCNNMultiaryKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNMultiaryKernel) Autorelease() MPSCNNMultiaryKernel {
	rv := objc.Send[MPSCNNMultiaryKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNMultiaryKernel creates a new MPSCNNMultiaryKernel instance.
func NewMPSCNNMultiaryKernel() MPSCNNMultiaryKernel {
	class := getMPSCNNMultiaryKernelClass()
	rv := objc.Send[MPSCNNMultiaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNMultiaryKernelWithCoder(aDecoder foundation.INSCoder) MPSCNNMultiaryKernel {
	instance := getMPSCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/init(coder:device:)
func NewCNNMultiaryKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNMultiaryKernel {
	instance := getMPSCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNMultiaryKernelFromID(rv)
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
func NewCNNMultiaryKernelWithDevice(device metal.MTLDevice) MPSCNNMultiaryKernel {
	instance := getMPSCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/init(device:sourceCount:)
func NewCNNMultiaryKernelWithDeviceSourceCount(device metal.MTLDevice, sourceCount uint) MPSCNNMultiaryKernel {
	instance := getMPSCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, sourceCount)
	return MPSCNNMultiaryKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/init(device:sourceCount:)
func (c MPSCNNMultiaryKernel) InitWithDeviceSourceCount(device metal.MTLDevice, sourceCount uint) MPSCNNMultiaryKernel {
	rv := objc.Send[MPSCNNMultiaryKernel](c.ID, objc.Sel("initWithDevice:sourceCount:"), device, sourceCount)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/appendBatchBarrier()
func (c MPSCNNMultiaryKernel) AppendBatchBarrier() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("appendBatchBarrier"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/destinationImageDescriptor(sourceImages:sourceStates:)
func (c MPSCNNMultiaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []MPSImage, sourceStates []MPSState) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates))
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateXatIndex(_:)
func (c MPSCNNMultiaryKernel) DilationRateXatIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateXatIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateYatIndex(_:)
func (c MPSCNNMultiaryKernel) DilationRateYatIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateYatIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/edgeMode(at:)
func (c MPSCNNMultiaryKernel) EdgeModeAtIndex(index uint) MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](c.ID, objc.Sel("edgeModeAtIndex:"), index)
	return MPSImageEdgeMode(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encode(commandBuffer:sourceImages:)
func (c MPSCNNMultiaryKernel) EncodeToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImages:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages))
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encode(commandBuffer:sourceImages:destinationImage:)
func (c MPSCNNMultiaryKernel) EncodeToCommandBufferSourceImagesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImages:destinationImage:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encode(commandBuffer:sourceImages:destinationState:destinationStateIsTemporary:)
func (c MPSCNNMultiaryKernel) EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, outState IMPSState, isTemporary bool) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImages:destinationState:destinationStateIsTemporary:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), outState, isTemporary)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encodeBatch(commandBuffer:sourceImages:)
func (c MPSCNNMultiaryKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImageBatches []foundation.NSArray) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImageBatches))
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encodeBatch(commandBuffer:sourceImages:destinationImages:)
func (c MPSCNNMultiaryKernel) EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationImages:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/encodeBatch(commandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:)
func (c MPSCNNMultiaryKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImageBatches []foundation.NSArray, outState MPSStateBatch, isTemporary bool) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImageBatches), outState, isTemporary)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/isResultStateReusedAcrossBatch()
func (c MPSCNNMultiaryKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/kernelHeight(at:)
func (c MPSCNNMultiaryKernel) KernelHeightAtIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeightAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/kernelWidth(at:)
func (c MPSCNNMultiaryKernel) KernelWidthAtIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidthAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/offset(at:)
func (c MPSCNNMultiaryKernel) OffsetAtIndex(index uint) MPSOffset {
	rv := objc.Send[MPSOffset](c.ID, objc.Sel("offsetAtIndex:"), index)
	return MPSOffset(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/resultState(sourceImages:sourceStates:destinationImage:)
func (c MPSCNNMultiaryKernel) ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages []MPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateForSourceImages:sourceStates:destinationImage:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/resultStateBatch(sourceImages:sourceStates:destinationImage:)
func (c MPSCNNMultiaryKernel) ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages []foundation.NSArray, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("resultStateBatchForSourceImages:sourceStates:destinationImage:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setDilationRateX(_:at:)
func (c MPSCNNMultiaryKernel) SetDilationRateXAtIndex(dilationRate uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDilationRateX:atIndex:"), dilationRate, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setDilationRateY(_:at:)
func (c MPSCNNMultiaryKernel) SetDilationRateYAtIndex(dilationRate uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDilationRateY:atIndex:"), dilationRate, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setEdgeMode(_:at:)
func (c MPSCNNMultiaryKernel) SetEdgeModeAtIndex(edgeMode MPSImageEdgeMode, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setEdgeMode:atIndex:"), edgeMode, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setKernelHeight(_:at:)
func (c MPSCNNMultiaryKernel) SetKernelHeightAtIndex(height uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setKernelHeight:atIndex:"), height, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setKernelWidth(_:at:)
func (c MPSCNNMultiaryKernel) SetKernelWidthAtIndex(width uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setKernelWidth:atIndex:"), width, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setOffset(_:at:)
func (c MPSCNNMultiaryKernel) SetOffsetAtIndex(offset MPSOffset, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setOffset:atIndex:"), offset, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setSourceFeatureChannelMaxCount(_:at:)
func (c MPSCNNMultiaryKernel) SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setSourceFeatureChannelMaxCount:atIndex:"), count, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setSourceFeatureChannelOffset(_:at:)
func (c MPSCNNMultiaryKernel) SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setSourceFeatureChannelOffset:atIndex:"), offset, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setStrideInPixelsX(_:at:)
func (c MPSCNNMultiaryKernel) SetStrideInPixelsXAtIndex(stride uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setStrideInPixelsX:atIndex:"), stride, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setStrideInPixelsY(_:at:)
func (c MPSCNNMultiaryKernel) SetStrideInPixelsYAtIndex(stride uint, index uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setStrideInPixelsY:atIndex:"), stride, index)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/sourceFeatureChannelMaxCount(at:)
func (c MPSCNNMultiaryKernel) SourceFeatureChannelMaxCountAtIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceFeatureChannelMaxCountAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/sourceFeatureChannelOffset(at:)
func (c MPSCNNMultiaryKernel) SourceFeatureChannelOffsetAtIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceFeatureChannelOffsetAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/stride(inPixelsXatIndex:)
func (c MPSCNNMultiaryKernel) StrideInPixelsXatIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsXatIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/stride(inPixelsYatIndex:)
func (c MPSCNNMultiaryKernel) StrideInPixelsYatIndex(index uint) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsYatIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/temporaryResultState(commandBuffer:sourceImages:sourceStates:destinationImage:)
func (c MPSCNNMultiaryKernel) TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage []MPSImage, sourceStates []MPSState, destinationImage IMPSImage) IMPSState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImage), objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/temporaryResultStateBatch(commandBuffer:sourceImages:sourceStates:destinationImage:)
func (c MPSCNNMultiaryKernel) TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage []foundation.NSArray, sourceStates []foundation.NSArray, destinationImage MPSImageBatch) MPSStateBatch {
	rv := objc.Send[MPSStateBatch](c.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImage), objectivec.IObjectSliceToNSArray(sourceStates), destinationImage)
	return MPSStateBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/clipRect
func (c MPSCNNMultiaryKernel) ClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](c.ID, objc.Sel("clipRect"))
	return metal.MTLRegion(rv)
}
func (c MPSCNNMultiaryKernel) SetClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](c.ID, objc.Sel("setClipRect:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/destinationFeatureChannelOffset
func (c MPSCNNMultiaryKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}
func (c MPSCNNMultiaryKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/destinationImageAllocator
func (c MPSCNNMultiaryKernel) DestinationImageAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}
func (c MPSCNNMultiaryKernel) SetDestinationImageAllocator(value MPSImageAllocator) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/isBackwards
func (c MPSCNNMultiaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBackwards"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/isStateModified
func (c MPSCNNMultiaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStateModified"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/padding
func (c MPSCNNMultiaryKernel) Padding() MPSNNPadding {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("padding"))
	return MPSNNPaddingObjectFromID(rv)
}
func (c MPSCNNMultiaryKernel) SetPadding(value MPSNNPadding) {
	objc.Send[struct{}](c.ID, objc.Sel("setPadding:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/sourceCount
func (c MPSCNNMultiaryKernel) SourceCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceCount"))
	return rv
}
