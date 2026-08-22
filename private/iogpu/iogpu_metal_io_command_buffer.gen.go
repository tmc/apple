// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIOCommandBuffer] class.
var (
	_IOGPUMetalIOCommandBufferClass     IOGPUMetalIOCommandBufferClass
	_IOGPUMetalIOCommandBufferClassOnce sync.Once
)

func getIOGPUMetalIOCommandBufferClass() IOGPUMetalIOCommandBufferClass {
	_IOGPUMetalIOCommandBufferClassOnce.Do(func() {
		_IOGPUMetalIOCommandBufferClass = IOGPUMetalIOCommandBufferClass{class: objc.GetClass("IOGPUMetalIOCommandBuffer")}
	})
	return _IOGPUMetalIOCommandBufferClass
}

// GetIOGPUMetalIOCommandBufferClass returns the class object for IOGPUMetalIOCommandBuffer.
func GetIOGPUMetalIOCommandBufferClass() IOGPUMetalIOCommandBufferClass {
	return getIOGPUMetalIOCommandBufferClass()
}

type IOGPUMetalIOCommandBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIOCommandBufferClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIOCommandBufferClass) Alloc() IOGPUMetalIOCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIOCommandBuffer.Barrier]
//   - [IOGPUMetalIOCommandBuffer.BarrierComplete]
//   - [IOGPUMetalIOCommandBuffer.CompleteCommandCallbackBlocks]
//   - [IOGPUMetalIOCommandBuffer.CreateCommandList]
//   - [IOGPUMetalIOCommandBuffer.DidCompleteWithStatus]
//   - [IOGPUMetalIOCommandBuffer.EncodeSignalEventValue]
//   - [IOGPUMetalIOCommandBuffer.EncodeWaitForEventValue]
//   - [IOGPUMetalIOCommandBuffer.GetCommandBufferBytes]
//   - [IOGPUMetalIOCommandBuffer.GlobalTraceObjectID]
//   - [IOGPUMetalIOCommandBuffer.GrowKernelCommandBuffer]
//   - [IOGPUMetalIOCommandBuffer.HasFollowOnWork]
//   - [IOGPUMetalIOCommandBuffer.IsCommitted]
//   - [IOGPUMetalIOCommandBuffer.LoadBufferOffsetSizeHandleHandleOffset]
//   - [IOGPUMetalIOCommandBuffer.LoadTextureSliceLevelSizeBytesPerRowBytesPerImageDstOriginHandleHandleOffset]
//   - [IOGPUMetalIOCommandBuffer.ValidateNotificationCount]
//   - [IOGPUMetalIOCommandBuffer.InitWithQueue]
//   - [IOGPUMetalIOCommandBuffer.InitWithQueueResourceListRetained]
type IOGPUMetalIOCommandBuffer struct {
	metal.MTLIOCommandBufferObject
}

// IOGPUMetalIOCommandBufferFromID constructs a [IOGPUMetalIOCommandBuffer] from an objc.ID.
func IOGPUMetalIOCommandBufferFromID(id objc.ID) IOGPUMetalIOCommandBuffer {
	return IOGPUMetalIOCommandBuffer{MTLIOCommandBufferObject: metal.MTLIOCommandBufferObjectFromID(id)}
}

// Ensure IOGPUMetalIOCommandBuffer implements IIOGPUMetalIOCommandBuffer.
var _ IIOGPUMetalIOCommandBuffer = IOGPUMetalIOCommandBuffer{}

// An interface definition for the [IOGPUMetalIOCommandBuffer] class.
//
// # Methods
//
//   - [IIOGPUMetalIOCommandBuffer.Barrier]
//   - [IIOGPUMetalIOCommandBuffer.BarrierComplete]
//   - [IIOGPUMetalIOCommandBuffer.CompleteCommandCallbackBlocks]
//   - [IIOGPUMetalIOCommandBuffer.CreateCommandList]
//   - [IIOGPUMetalIOCommandBuffer.DidCompleteWithStatus]
//   - [IIOGPUMetalIOCommandBuffer.EncodeSignalEventValue]
//   - [IIOGPUMetalIOCommandBuffer.EncodeWaitForEventValue]
//   - [IIOGPUMetalIOCommandBuffer.GetCommandBufferBytes]
//   - [IIOGPUMetalIOCommandBuffer.GlobalTraceObjectID]
//   - [IIOGPUMetalIOCommandBuffer.GrowKernelCommandBuffer]
//   - [IIOGPUMetalIOCommandBuffer.HasFollowOnWork]
//   - [IIOGPUMetalIOCommandBuffer.IsCommitted]
//   - [IIOGPUMetalIOCommandBuffer.LoadBufferOffsetSizeHandleHandleOffset]
//   - [IIOGPUMetalIOCommandBuffer.LoadTextureSliceLevelSizeBytesPerRowBytesPerImageDstOriginHandleHandleOffset]
//   - [IIOGPUMetalIOCommandBuffer.ValidateNotificationCount]
//   - [IIOGPUMetalIOCommandBuffer.InitWithQueue]
//   - [IIOGPUMetalIOCommandBuffer.InitWithQueueResourceListRetained]
type IIOGPUMetalIOCommandBuffer interface {
	metal.MTLIOCommandBuffer

	// Topic: Methods

	Barrier()
	BarrierComplete(complete int64)
	CompleteCommandCallbackBlocks()
	CreateCommandList()
	DidCompleteWithStatus(status int64)
	EncodeSignalEventValue(event objectivec.IObject, value uint64)
	EncodeWaitForEventValue(event objectivec.IObject, value uint64)
	GetCommandBufferBytes(bytes uint32) unsafe.Pointer
	GlobalTraceObjectID() uint64
	GrowKernelCommandBuffer(buffer uint32)
	HasFollowOnWork() bool
	IsCommitted() bool
	LoadBufferOffsetSizeHandleHandleOffset(buffer objectivec.IObject, offset uint64, size uint64, handle objectivec.IObject, offset2 uint64)
	LoadTextureSliceLevelSizeBytesPerRowBytesPerImageDstOriginHandleHandleOffset(texture objectivec.IObject, slice uint64, level uint64, size unsafe.Pointer, row uint64, image uint64, origin unsafe.Pointer, handle objectivec.IObject, offset uint64)
	ValidateNotificationCount() bool
	InitWithQueue(queue objectivec.IObject) IOGPUMetalIOCommandBuffer
	InitWithQueueResourceListRetained(queue objectivec.IObject, list objectivec.IObject, retained bool) IOGPUMetalIOCommandBuffer
}

// Init initializes the instance.
func (i IOGPUMetalIOCommandBuffer) Init() IOGPUMetalIOCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIOCommandBuffer) Autorelease() IOGPUMetalIOCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIOCommandBuffer creates a new IOGPUMetalIOCommandBuffer instance.
func NewIOGPUMetalIOCommandBuffer() IOGPUMetalIOCommandBuffer {
	class := getIOGPUMetalIOCommandBufferClass()
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIOCommandBufferWithQueue(queue objectivec.IObject) IOGPUMetalIOCommandBuffer {
	instance := getIOGPUMetalIOCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithQueue:"), queue)
	return IOGPUMetalIOCommandBufferFromID(rv)
}

func NewGPUMetalIOCommandBufferWithQueueResourceListRetained(queue objectivec.IObject, list objectivec.IObject, retained bool) IOGPUMetalIOCommandBuffer {
	instance := getIOGPUMetalIOCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithQueue:resourceList:retained:"), queue, list, retained)
	return IOGPUMetalIOCommandBufferFromID(rv)
}

func (i IOGPUMetalIOCommandBuffer) Barrier() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("barrier"))
}
func (i IOGPUMetalIOCommandBuffer) BarrierComplete(complete int64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("barrierComplete:"), complete)
}
func (i IOGPUMetalIOCommandBuffer) CompleteCommandCallbackBlocks() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("completeCommandCallbackBlocks"))
}
func (i IOGPUMetalIOCommandBuffer) CreateCommandList() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("createCommandList"))
}
func (i IOGPUMetalIOCommandBuffer) DidCompleteWithStatus(status int64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("didCompleteWithStatus:"), status)
}
func (i IOGPUMetalIOCommandBuffer) EncodeSignalEventValue(event objectivec.IObject, value uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeSignalEvent:value:"), event, value)
}
func (i IOGPUMetalIOCommandBuffer) EncodeWaitForEventValue(event objectivec.IObject, value uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeWaitForEvent:value:"), event, value)
}
func (i IOGPUMetalIOCommandBuffer) GetCommandBufferBytes(bytes uint32) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getCommandBufferBytes:"), bytes)
	return rv
}
func (i IOGPUMetalIOCommandBuffer) GlobalTraceObjectID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalTraceObjectID"))
	return rv
}
func (i IOGPUMetalIOCommandBuffer) GrowKernelCommandBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growKernelCommandBuffer:"), buffer)
}
func (i IOGPUMetalIOCommandBuffer) HasFollowOnWork() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("hasFollowOnWork"))
	return rv
}
func (i IOGPUMetalIOCommandBuffer) IsCommitted() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isCommitted"))
	return rv
}
func (i IOGPUMetalIOCommandBuffer) LoadBufferOffsetSizeHandleHandleOffset(buffer objectivec.IObject, offset uint64, size uint64, handle objectivec.IObject, offset2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("loadBuffer:offset:size:handle:handleOffset:"), buffer, offset, size, handle, offset2)
}
func (i IOGPUMetalIOCommandBuffer) LoadTextureSliceLevelSizeBytesPerRowBytesPerImageDstOriginHandleHandleOffset(texture objectivec.IObject, slice uint64, level uint64, size unsafe.Pointer, row uint64, image uint64, origin unsafe.Pointer, handle objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("loadTexture:slice:level:size:bytesPerRow:bytesPerImage:dstOrigin:handle:handleOffset:"), texture, slice, level, size, row, image, origin, handle, offset)
}
func (i IOGPUMetalIOCommandBuffer) ValidateNotificationCount() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("validateNotificationCount"))
	return rv
}
func (i IOGPUMetalIOCommandBuffer) InitWithQueue(queue objectivec.IObject) IOGPUMetalIOCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](i.ID, objc.Sel("initWithQueue:"), queue)
	return rv
}
func (i IOGPUMetalIOCommandBuffer) InitWithQueueResourceListRetained(queue objectivec.IObject, list objectivec.IObject, retained bool) IOGPUMetalIOCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandBuffer](i.ID, objc.Sel("initWithQueue:resourceList:retained:"), queue, list, retained)
	return rv
}
