// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4CommandBuffer] class.
var (
	_IOGPUMetal4CommandBufferClass     IOGPUMetal4CommandBufferClass
	_IOGPUMetal4CommandBufferClassOnce sync.Once
)

func getIOGPUMetal4CommandBufferClass() IOGPUMetal4CommandBufferClass {
	_IOGPUMetal4CommandBufferClassOnce.Do(func() {
		_IOGPUMetal4CommandBufferClass = IOGPUMetal4CommandBufferClass{class: objc.GetClass("IOGPUMetal4CommandBuffer")}
	})
	return _IOGPUMetal4CommandBufferClass
}

// GetIOGPUMetal4CommandBufferClass returns the class object for IOGPUMetal4CommandBuffer.
func GetIOGPUMetal4CommandBufferClass() IOGPUMetal4CommandBufferClass {
	return getIOGPUMetal4CommandBufferClass()
}

type IOGPUMetal4CommandBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4CommandBufferClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4CommandBufferClass) Alloc() IOGPUMetal4CommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetal4CommandBuffer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4CommandBuffer._reserveKernelCommandBufferSpace]
//   - [IOGPUMetal4CommandBuffer.AkPrivateResourceList]
//   - [IOGPUMetal4CommandBuffer.AkResourceList]
//   - [IOGPUMetal4CommandBuffer.AllocCommandBufferResourceAtIndex]
//   - [IOGPUMetal4CommandBuffer.AllocDebugBuffer]
//   - [IOGPUMetal4CommandBuffer.AllocateSidebandBuffer]
//   - [IOGPUMetal4CommandBuffer.BeginIOGPUCommandBufferWithAllocatorOptions]
//   - [IOGPUMetal4CommandBuffer.BeginSegment]
//   - [IOGPUMetal4CommandBuffer.CommandBufferStorage]
//   - [IOGPUMetal4CommandBuffer.CommitEncoder]
//   - [IOGPUMetal4CommandBuffer.CopyBufferMappingsFromBufferToBufferOperationsCount]
//   - [IOGPUMetal4CommandBuffer.EncodePostMappingWaitEventPostMappingValueTimeout]
//   - [IOGPUMetal4CommandBuffer.EncodeSignalEventValue]
//   - [IOGPUMetal4CommandBuffer.EncodeWaitForEventValue]
//   - [IOGPUMetal4CommandBuffer.EncodeWaitForEventValueTimeout]
//   - [IOGPUMetal4CommandBuffer.EndCurrentSegment]
//   - [IOGPUMetal4CommandBuffer.FillCommandBufferArgs]
//   - [IOGPUMetal4CommandBuffer.GetCurrentKernelCommandBufferPointerEnd]
//   - [IOGPUMetal4CommandBuffer.GetCurrentKernelCommandBufferStartCurrentEnd]
//   - [IOGPUMetal4CommandBuffer.GetDebugBufferPointerStartEnd]
//   - [IOGPUMetal4CommandBuffer.GetSegmentListPointerStartCurrentEnd]
//   - [IOGPUMetal4CommandBuffer.GrowDebugBuffer]
//   - [IOGPUMetal4CommandBuffer.GrowKernelCommandBuffer]
//   - [IOGPUMetal4CommandBuffer.GrowSegmentList]
//   - [IOGPUMetal4CommandBuffer.GrowSidebandBuffer]
//   - [IOGPUMetal4CommandBuffer.IoGPUResourceList]
//   - [IOGPUMetal4CommandBuffer.ProtectionOptions]
//   - [IOGPUMetal4CommandBuffer.ResetCommandBuffer]
//   - [IOGPUMetal4CommandBuffer.SetCurrentKernelCommandBufferPointer]
//   - [IOGPUMetal4CommandBuffer.SetCurrentSegmentListPointer]
//   - [IOGPUMetal4CommandBuffer.SetProtectionOptions]
//   - [IOGPUMetal4CommandBuffer.UpdateBufferMappingsHeapOperationsCount]
//   - [IOGPUMetal4CommandBuffer.UseInternalResidencySet]
//   - [IOGPUMetal4CommandBuffer.UseInternalResidencySetsCount]
//   - [IOGPUMetal4CommandBuffer.InitWithDevice]
type IOGPUMetal4CommandBuffer struct {
	metal.MTL4CommandBufferObject
}

// IOGPUMetal4CommandBufferFromID constructs a [IOGPUMetal4CommandBuffer] from an objc.ID.
func IOGPUMetal4CommandBufferFromID(id objc.ID) IOGPUMetal4CommandBuffer {
	return IOGPUMetal4CommandBuffer{MTL4CommandBufferObject: metal.MTL4CommandBufferObjectFromID(id)}
}

// Ensure IOGPUMetal4CommandBuffer implements IIOGPUMetal4CommandBuffer.
var _ IIOGPUMetal4CommandBuffer = IOGPUMetal4CommandBuffer{}

// An interface definition for the [IOGPUMetal4CommandBuffer] class.
//
// # Methods
//
//   - [IIOGPUMetal4CommandBuffer._reserveKernelCommandBufferSpace]
//   - [IIOGPUMetal4CommandBuffer.AkPrivateResourceList]
//   - [IIOGPUMetal4CommandBuffer.AkResourceList]
//   - [IIOGPUMetal4CommandBuffer.AllocCommandBufferResourceAtIndex]
//   - [IIOGPUMetal4CommandBuffer.AllocDebugBuffer]
//   - [IIOGPUMetal4CommandBuffer.AllocateSidebandBuffer]
//   - [IIOGPUMetal4CommandBuffer.BeginIOGPUCommandBufferWithAllocatorOptions]
//   - [IIOGPUMetal4CommandBuffer.BeginSegment]
//   - [IIOGPUMetal4CommandBuffer.CommandBufferStorage]
//   - [IIOGPUMetal4CommandBuffer.CommitEncoder]
//   - [IIOGPUMetal4CommandBuffer.CopyBufferMappingsFromBufferToBufferOperationsCount]
//   - [IIOGPUMetal4CommandBuffer.EncodePostMappingWaitEventPostMappingValueTimeout]
//   - [IIOGPUMetal4CommandBuffer.EncodeSignalEventValue]
//   - [IIOGPUMetal4CommandBuffer.EncodeWaitForEventValue]
//   - [IIOGPUMetal4CommandBuffer.EncodeWaitForEventValueTimeout]
//   - [IIOGPUMetal4CommandBuffer.EndCurrentSegment]
//   - [IIOGPUMetal4CommandBuffer.FillCommandBufferArgs]
//   - [IIOGPUMetal4CommandBuffer.GetCurrentKernelCommandBufferPointerEnd]
//   - [IIOGPUMetal4CommandBuffer.GetCurrentKernelCommandBufferStartCurrentEnd]
//   - [IIOGPUMetal4CommandBuffer.GetDebugBufferPointerStartEnd]
//   - [IIOGPUMetal4CommandBuffer.GetSegmentListPointerStartCurrentEnd]
//   - [IIOGPUMetal4CommandBuffer.GrowDebugBuffer]
//   - [IIOGPUMetal4CommandBuffer.GrowKernelCommandBuffer]
//   - [IIOGPUMetal4CommandBuffer.GrowSegmentList]
//   - [IIOGPUMetal4CommandBuffer.GrowSidebandBuffer]
//   - [IIOGPUMetal4CommandBuffer.IoGPUResourceList]
//   - [IIOGPUMetal4CommandBuffer.ProtectionOptions]
//   - [IIOGPUMetal4CommandBuffer.ResetCommandBuffer]
//   - [IIOGPUMetal4CommandBuffer.SetCurrentKernelCommandBufferPointer]
//   - [IIOGPUMetal4CommandBuffer.SetCurrentSegmentListPointer]
//   - [IIOGPUMetal4CommandBuffer.SetProtectionOptions]
//   - [IIOGPUMetal4CommandBuffer.UpdateBufferMappingsHeapOperationsCount]
//   - [IIOGPUMetal4CommandBuffer.UseInternalResidencySet]
//   - [IIOGPUMetal4CommandBuffer.UseInternalResidencySetsCount]
//   - [IIOGPUMetal4CommandBuffer.InitWithDevice]
type IIOGPUMetal4CommandBuffer interface {
	metal.MTL4CommandBuffer

	// Topic: Methods

	_reserveKernelCommandBufferSpace(space uint64) unsafe.Pointer
	AkPrivateResourceList() objectivec.IObject
	AkResourceList() objectivec.IObject
	AllocCommandBufferResourceAtIndex(index uint32)
	AllocDebugBuffer()
	AllocateSidebandBuffer(buffer uint32)
	BeginIOGPUCommandBufferWithAllocatorOptions(allocator objectivec.IObject, options objectivec.IObject)
	BeginSegment(segment unsafe.Pointer)
	CommandBufferStorage() *IOGPUMetalCommandBufferStorage
	CommitEncoder()
	CopyBufferMappingsFromBufferToBufferOperationsCount(buffer objectivec.IObject, buffer2 objectivec.IObject, operations unsafe.Pointer, count uint64)
	EncodePostMappingWaitEventPostMappingValueTimeout(event objectivec.IObject, value uint64, timeout uint32)
	EncodeSignalEventValue(event objectivec.IObject, value uint64)
	EncodeWaitForEventValue(event objectivec.IObject, value uint64)
	EncodeWaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint32)
	EndCurrentSegment()
	FillCommandBufferArgs(args *IOGPUCommandQueueCommandBufferArgs)
	GetCurrentKernelCommandBufferPointerEnd(pointer unsafe.Pointer, end unsafe.Pointer)
	GetCurrentKernelCommandBufferStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer)
	GetDebugBufferPointerStartEnd(start unsafe.Pointer, end unsafe.Pointer)
	GetSegmentListPointerStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer)
	GrowDebugBuffer(buffer uint32)
	GrowKernelCommandBuffer(buffer uint64)
	GrowSegmentList()
	GrowSidebandBuffer(buffer uint32)
	IoGPUResourceList() *IOGPUResourceList
	ProtectionOptions() uint64
	ResetCommandBuffer()
	SetCurrentKernelCommandBufferPointer(pointer unsafe.Pointer)
	SetCurrentSegmentListPointer(pointer unsafe.Pointer)
	SetProtectionOptions(options uint64)
	UpdateBufferMappingsHeapOperationsCount(mappings objectivec.IObject, heap objectivec.IObject, operations unsafe.Pointer, count uint64)
	UseInternalResidencySet(set objectivec.IObject)
	UseInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandBuffer
}

// Init initializes the instance.
func (i IOGPUMetal4CommandBuffer) Init() IOGPUMetal4CommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetal4CommandBuffer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4CommandBuffer) Autorelease() IOGPUMetal4CommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetal4CommandBuffer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4CommandBuffer creates a new IOGPUMetal4CommandBuffer instance.
func NewIOGPUMetal4CommandBuffer() IOGPUMetal4CommandBuffer {
	class := getIOGPUMetal4CommandBufferClass()
	rv := objc.SendIfResponds[IOGPUMetal4CommandBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4CommandBufferWithDevice(device objectivec.IObject) IOGPUMetal4CommandBuffer {
	instance := getIOGPUMetal4CommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetal4CommandBufferFromID(rv)
}

func (i IOGPUMetal4CommandBuffer) _reserveKernelCommandBufferSpace(space uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("_reserveKernelCommandBufferSpace:"), space)
	return rv
}

// ReserveKernelCommandBufferSpace is an exported wrapper for the private method _reserveKernelCommandBufferSpace.
func (i IOGPUMetal4CommandBuffer) ReserveKernelCommandBufferSpace(space uint64) (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_reserveKernelCommandBufferSpace:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reserveKernelCommandBufferSpace:"}
		return nil, err
	}
	return i._reserveKernelCommandBufferSpace(space), nil
}

// CanReserveKernelCommandBufferSpace reports whether the receiver responds to the private selector _reserveKernelCommandBufferSpace:.
func (i IOGPUMetal4CommandBuffer) CanReserveKernelCommandBufferSpace() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_reserveKernelCommandBufferSpace:"))
}
func (i IOGPUMetal4CommandBuffer) AkPrivateResourceList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akPrivateResourceList"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetal4CommandBuffer) AkResourceList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akResourceList"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetal4CommandBuffer) AllocCommandBufferResourceAtIndex(index uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocCommandBufferResourceAtIndex:"), index)
}
func (i IOGPUMetal4CommandBuffer) AllocDebugBuffer() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocDebugBuffer"))
}
func (i IOGPUMetal4CommandBuffer) AllocateSidebandBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocateSidebandBuffer:"), buffer)
}
func (i IOGPUMetal4CommandBuffer) BeginIOGPUCommandBufferWithAllocatorOptions(allocator objectivec.IObject, options objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("beginIOGPUCommandBufferWithAllocator:options:"), allocator, options)
}
func (i IOGPUMetal4CommandBuffer) BeginSegment(segment unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("beginSegment:"), segment)
}
func (i IOGPUMetal4CommandBuffer) CommandBufferStorage() *IOGPUMetalCommandBufferStorage {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("commandBufferStorage"))
	return (*IOGPUMetalCommandBufferStorage)(rv)
}
func (i IOGPUMetal4CommandBuffer) CommitEncoder() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitEncoder"))
}
func (i IOGPUMetal4CommandBuffer) CopyBufferMappingsFromBufferToBufferOperationsCount(buffer objectivec.IObject, buffer2 objectivec.IObject, operations unsafe.Pointer, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyBufferMappingsFromBuffer:toBuffer:operations:count:"), buffer, buffer2, objc.CArray(operations), count)
}
func (i IOGPUMetal4CommandBuffer) EncodePostMappingWaitEventPostMappingValueTimeout(event objectivec.IObject, value uint64, timeout uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodePostMappingWaitEvent:postMappingValue:timeout:"), event, value, timeout)
}
func (i IOGPUMetal4CommandBuffer) EncodeSignalEventValue(event objectivec.IObject, value uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeSignalEvent:value:"), event, value)
}
func (i IOGPUMetal4CommandBuffer) EncodeWaitForEventValue(event objectivec.IObject, value uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeWaitForEvent:value:"), event, value)
}
func (i IOGPUMetal4CommandBuffer) EncodeWaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeWaitForEvent:value:timeout:"), event, value, timeout)
}
func (i IOGPUMetal4CommandBuffer) EndCurrentSegment() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("endCurrentSegment"))
}
func (i IOGPUMetal4CommandBuffer) FillCommandBufferArgs(args *IOGPUCommandQueueCommandBufferArgs) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillCommandBufferArgs:"), unsafe.Pointer(args))
}
func (i IOGPUMetal4CommandBuffer) GetCurrentKernelCommandBufferPointerEnd(pointer unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getCurrentKernelCommandBufferPointer:end:"), pointer, end)
}
func (i IOGPUMetal4CommandBuffer) GetCurrentKernelCommandBufferStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getCurrentKernelCommandBufferStart:current:end:"), start, current, end)
}
func (i IOGPUMetal4CommandBuffer) GetDebugBufferPointerStartEnd(start unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getDebugBufferPointerStart:end:"), start, end)
}
func (i IOGPUMetal4CommandBuffer) GetSegmentListPointerStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getSegmentListPointerStart:current:end:"), start, current, end)
}
func (i IOGPUMetal4CommandBuffer) GrowDebugBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growDebugBuffer:"), buffer)
}
func (i IOGPUMetal4CommandBuffer) GrowKernelCommandBuffer(buffer uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growKernelCommandBuffer:"), buffer)
}
func (i IOGPUMetal4CommandBuffer) GrowSegmentList() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growSegmentList"))
}
func (i IOGPUMetal4CommandBuffer) GrowSidebandBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growSidebandBuffer:"), buffer)
}
func (i IOGPUMetal4CommandBuffer) IoGPUResourceList() *IOGPUResourceList {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("ioGPUResourceList"))
	return (*IOGPUResourceList)(rv)
}
func (i IOGPUMetal4CommandBuffer) ProtectionOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("protectionOptions"))
	return rv
}
func (i IOGPUMetal4CommandBuffer) ResetCommandBuffer() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("resetCommandBuffer"))
}
func (i IOGPUMetal4CommandBuffer) SetCurrentKernelCommandBufferPointer(pointer unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentKernelCommandBufferPointer:"), pointer)
}
func (i IOGPUMetal4CommandBuffer) SetCurrentSegmentListPointer(pointer unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentSegmentListPointer:"), pointer)
}
func (i IOGPUMetal4CommandBuffer) SetProtectionOptions(options uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setProtectionOptions:"), options)
}
func (i IOGPUMetal4CommandBuffer) UpdateBufferMappingsHeapOperationsCount(mappings objectivec.IObject, heap objectivec.IObject, operations unsafe.Pointer, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateBufferMappings:heap:operations:count:"), mappings, heap, objc.CArray(operations), count)
}
func (i IOGPUMetal4CommandBuffer) UseInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useInternalResidencySet:"), set)
}
func (i IOGPUMetal4CommandBuffer) UseInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetal4CommandBuffer) InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetal4CommandBuffer](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
