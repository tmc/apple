// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalBlitCommandEncoder] class.
var (
	_IOGPUMetalBlitCommandEncoderClass     IOGPUMetalBlitCommandEncoderClass
	_IOGPUMetalBlitCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalBlitCommandEncoderClass() IOGPUMetalBlitCommandEncoderClass {
	_IOGPUMetalBlitCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalBlitCommandEncoderClass = IOGPUMetalBlitCommandEncoderClass{class: objc.GetClass("IOGPUMetalBlitCommandEncoder")}
	})
	return _IOGPUMetalBlitCommandEncoderClass
}

// GetIOGPUMetalBlitCommandEncoderClass returns the class object for IOGPUMetalBlitCommandEncoder.
func GetIOGPUMetalBlitCommandEncoderClass() IOGPUMetalBlitCommandEncoderClass {
	return getIOGPUMetalBlitCommandEncoderClass()
}

type IOGPUMetalBlitCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalBlitCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalBlitCommandEncoderClass) Alloc() IOGPUMetalBlitCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalBlitCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalBlitCommandEncoder.CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount]
//   - [IOGPUMetalBlitCommandEncoder.CopyFromTextureToTexture]
//   - [IOGPUMetalBlitCommandEncoder.CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex]
//   - [IOGPUMetalBlitCommandEncoder.FillBufferRangePattern4]
//   - [IOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionBytesLength]
//   - [IOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionColor]
//   - [IOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionColorPixelFormat]
//   - [IOGPUMetalBlitCommandEncoder.GetType]
//   - [IOGPUMetalBlitCommandEncoder.OptimizeContentsForCPUAccess]
//   - [IOGPUMetalBlitCommandEncoder.OptimizeContentsForCPUAccessSliceLevel]
//   - [IOGPUMetalBlitCommandEncoder.OptimizeContentsForGPUAccess]
//   - [IOGPUMetalBlitCommandEncoder.OptimizeContentsForGPUAccessSliceLevel]
//   - [IOGPUMetalBlitCommandEncoder.OptimizeIndirectCommandBufferWithRange]
//   - [IOGPUMetalBlitCommandEncoder.ResetCommandsInBufferWithRange]
//   - [IOGPUMetalBlitCommandEncoder.ResolveCountersInRangeDestinationBufferDestinationOffset]
//   - [IOGPUMetalBlitCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IOGPUMetalBlitCommandEncoder.SynchronizeResource]
//   - [IOGPUMetalBlitCommandEncoder.SynchronizeTextureSliceLevel]
//   - [IOGPUMetalBlitCommandEncoder.UpdateFence]
//   - [IOGPUMetalBlitCommandEncoder.WaitForFence]
type IOGPUMetalBlitCommandEncoder struct {
	IOGPUMetalCommandEncoder
}

// IOGPUMetalBlitCommandEncoderFromID constructs a [IOGPUMetalBlitCommandEncoder] from an objc.ID.
func IOGPUMetalBlitCommandEncoderFromID(id objc.ID) IOGPUMetalBlitCommandEncoder {
	return IOGPUMetalBlitCommandEncoder{IOGPUMetalCommandEncoder: IOGPUMetalCommandEncoderFromID(id)}
}

// Ensure IOGPUMetalBlitCommandEncoder implements IIOGPUMetalBlitCommandEncoder.
var _ IIOGPUMetalBlitCommandEncoder = IOGPUMetalBlitCommandEncoder{}

// An interface definition for the [IOGPUMetalBlitCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalBlitCommandEncoder.CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount]
//   - [IIOGPUMetalBlitCommandEncoder.CopyFromTextureToTexture]
//   - [IIOGPUMetalBlitCommandEncoder.CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex]
//   - [IIOGPUMetalBlitCommandEncoder.FillBufferRangePattern4]
//   - [IIOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionBytesLength]
//   - [IIOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionColor]
//   - [IIOGPUMetalBlitCommandEncoder.FillTextureLevelSliceRegionColorPixelFormat]
//   - [IIOGPUMetalBlitCommandEncoder.GetType]
//   - [IIOGPUMetalBlitCommandEncoder.OptimizeContentsForCPUAccess]
//   - [IIOGPUMetalBlitCommandEncoder.OptimizeContentsForCPUAccessSliceLevel]
//   - [IIOGPUMetalBlitCommandEncoder.OptimizeContentsForGPUAccess]
//   - [IIOGPUMetalBlitCommandEncoder.OptimizeContentsForGPUAccessSliceLevel]
//   - [IIOGPUMetalBlitCommandEncoder.OptimizeIndirectCommandBufferWithRange]
//   - [IIOGPUMetalBlitCommandEncoder.ResetCommandsInBufferWithRange]
//   - [IIOGPUMetalBlitCommandEncoder.ResolveCountersInRangeDestinationBufferDestinationOffset]
//   - [IIOGPUMetalBlitCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IIOGPUMetalBlitCommandEncoder.SynchronizeResource]
//   - [IIOGPUMetalBlitCommandEncoder.SynchronizeTextureSliceLevel]
//   - [IIOGPUMetalBlitCommandEncoder.UpdateFence]
//   - [IIOGPUMetalBlitCommandEncoder.WaitForFence]
type IIOGPUMetalBlitCommandEncoder interface {
	IIOGPUMetalCommandEncoder

	// Topic: Methods

	CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(texture objectivec.IObject, slice uint64, level uint64, texture2 objectivec.IObject, slice2 uint64, level2 uint64, count uint64, count2 uint64)
	CopyFromTextureToTexture(texture objectivec.IObject, texture2 objectivec.IObject)
	CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(buffer objectivec.IObject, range_ foundation.NSRange, destination objectivec.IObject, index uint64)
	FillBufferRangePattern4(buffer objectivec.IObject, range_ foundation.NSRange, pattern4 uint32)
	FillTextureLevelSliceRegionBytesLength(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, bytes []byte)
	FillTextureLevelSliceRegionColor(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, color unsafe.Pointer)
	FillTextureLevelSliceRegionColorPixelFormat(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, color unsafe.Pointer, format uint64)
	GetType() uint64
	OptimizeContentsForCPUAccess(cPUAccess objectivec.IObject)
	OptimizeContentsForCPUAccessSliceLevel(cPUAccess objectivec.IObject, slice uint64, level uint64)
	OptimizeContentsForGPUAccess(gPUAccess objectivec.IObject)
	OptimizeContentsForGPUAccessSliceLevel(gPUAccess objectivec.IObject, slice uint64, level uint64)
	OptimizeIndirectCommandBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange)
	ResetCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange)
	ResolveCountersInRangeDestinationBufferDestinationOffset(counters objectivec.IObject, range_ foundation.NSRange, buffer objectivec.IObject, offset uint64)
	SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool)
	SynchronizeResource(resource objectivec.IObject)
	SynchronizeTextureSliceLevel(texture objectivec.IObject, slice uint64, level uint64)
	UpdateFence(fence objectivec.IObject)
	WaitForFence(fence objectivec.IObject)
}

// Init initializes the instance.
func (i IOGPUMetalBlitCommandEncoder) Init() IOGPUMetalBlitCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalBlitCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalBlitCommandEncoder) Autorelease() IOGPUMetalBlitCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalBlitCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalBlitCommandEncoder creates a new IOGPUMetalBlitCommandEncoder instance.
func NewIOGPUMetalBlitCommandEncoder() IOGPUMetalBlitCommandEncoder {
	class := getIOGPUMetalBlitCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalBlitCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalBlitCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalBlitCommandEncoder {
	instance := getIOGPUMetalBlitCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalBlitCommandEncoderFromID(rv)
}

func (i IOGPUMetalBlitCommandEncoder) CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(texture objectivec.IObject, slice uint64, level uint64, texture2 objectivec.IObject, slice2 uint64, level2 uint64, count uint64, count2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:toTexture:destinationSlice:destinationLevel:sliceCount:levelCount:"), texture, slice, level, texture2, slice2, level2, count, count2)
}
func (i IOGPUMetalBlitCommandEncoder) CopyFromTextureToTexture(texture objectivec.IObject, texture2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyFromTexture:toTexture:"), texture, texture2)
}
func (i IOGPUMetalBlitCommandEncoder) CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(buffer objectivec.IObject, range_ foundation.NSRange, destination objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:"), buffer, range_, destination, index)
}
func (i IOGPUMetalBlitCommandEncoder) FillBufferRangePattern4(buffer objectivec.IObject, range_ foundation.NSRange, pattern4 uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillBuffer:range:pattern4:"), buffer, range_, pattern4)
}
func (i IOGPUMetalBlitCommandEncoder) FillTextureLevelSliceRegionBytesLength(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, bytes []byte) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillTexture:level:slice:region:bytes:length:"), texture, level, slice, region, objc.BytesPointer(bytes), uint(len(bytes)))
}
func (i IOGPUMetalBlitCommandEncoder) FillTextureLevelSliceRegionColor(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, color unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillTexture:level:slice:region:color:"), texture, level, slice, region, color)
}
func (i IOGPUMetalBlitCommandEncoder) FillTextureLevelSliceRegionColorPixelFormat(texture objectivec.IObject, level uint64, slice uint64, region unsafe.Pointer, color unsafe.Pointer, format uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillTexture:level:slice:region:color:pixelFormat:"), texture, level, slice, region, color, format)
}
func (i IOGPUMetalBlitCommandEncoder) GetType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetalBlitCommandEncoder) OptimizeContentsForCPUAccess(cPUAccess objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("optimizeContentsForCPUAccess:"), cPUAccess)
}
func (i IOGPUMetalBlitCommandEncoder) OptimizeContentsForCPUAccessSliceLevel(cPUAccess objectivec.IObject, slice uint64, level uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("optimizeContentsForCPUAccess:slice:level:"), cPUAccess, slice, level)
}
func (i IOGPUMetalBlitCommandEncoder) OptimizeContentsForGPUAccess(gPUAccess objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("optimizeContentsForGPUAccess:"), gPUAccess)
}
func (i IOGPUMetalBlitCommandEncoder) OptimizeContentsForGPUAccessSliceLevel(gPUAccess objectivec.IObject, slice uint64, level uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("optimizeContentsForGPUAccess:slice:level:"), gPUAccess, slice, level)
}
func (i IOGPUMetalBlitCommandEncoder) OptimizeIndirectCommandBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("optimizeIndirectCommandBuffer:withRange:"), buffer, range_)
}
func (i IOGPUMetalBlitCommandEncoder) ResetCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("resetCommandsInBuffer:withRange:"), buffer, range_)
}
func (i IOGPUMetalBlitCommandEncoder) ResolveCountersInRangeDestinationBufferDestinationOffset(counters objectivec.IObject, range_ foundation.NSRange, buffer objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("resolveCounters:inRange:destinationBuffer:destinationOffset:"), counters, range_, buffer, offset)
}
func (i IOGPUMetalBlitCommandEncoder) SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), buffer, index, barrier)
}
func (i IOGPUMetalBlitCommandEncoder) SynchronizeResource(resource objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("synchronizeResource:"), resource)
}
func (i IOGPUMetalBlitCommandEncoder) SynchronizeTextureSliceLevel(texture objectivec.IObject, slice uint64, level uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("synchronizeTexture:slice:level:"), texture, slice, level)
}
func (i IOGPUMetalBlitCommandEncoder) UpdateFence(fence objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateFence:"), fence)
}
func (i IOGPUMetalBlitCommandEncoder) WaitForFence(fence objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("waitForFence:"), fence)
}
