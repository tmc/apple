// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalCommandBuffer] class.
var (
	_IOGPUMetalCommandBufferClass     IOGPUMetalCommandBufferClass
	_IOGPUMetalCommandBufferClassOnce sync.Once
)

func getIOGPUMetalCommandBufferClass() IOGPUMetalCommandBufferClass {
	_IOGPUMetalCommandBufferClassOnce.Do(func() {
		_IOGPUMetalCommandBufferClass = IOGPUMetalCommandBufferClass{class: objc.GetClass("IOGPUMetalCommandBuffer")}
	})
	return _IOGPUMetalCommandBufferClass
}

// GetIOGPUMetalCommandBufferClass returns the class object for IOGPUMetalCommandBuffer.
func GetIOGPUMetalCommandBufferClass() IOGPUMetalCommandBufferClass {
	return getIOGPUMetalCommandBufferClass()
}

type IOGPUMetalCommandBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalCommandBufferClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalCommandBufferClass) Alloc() IOGPUMetalCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalCommandBuffer._debugBytesLengthOutput_type]
//   - [IOGPUMetalCommandBuffer._encodePurgedResources]
//   - [IOGPUMetalCommandBuffer._reserveKernelCommandBufferSpace]
//   - [IOGPUMetalCommandBuffer.AddPurgedHeap]
//   - [IOGPUMetalCommandBuffer.AddPurgedResource]
//   - [IOGPUMetalCommandBuffer.AkPrivateResourceList]
//   - [IOGPUMetalCommandBuffer.AkResourceList]
//   - [IOGPUMetalCommandBuffer.AllocCommandBufferResourceAtIndex]
//   - [IOGPUMetalCommandBuffer.AllocDebugBuffer]
//   - [IOGPUMetalCommandBuffer.AllocateSidebandBuffer]
//   - [IOGPUMetalCommandBuffer.BeginSegment]
//   - [IOGPUMetalCommandBuffer.CommandBufferResourceInfo]
//   - [IOGPUMetalCommandBuffer.CommandBufferStorage]
//   - [IOGPUMetalCommandBuffer.CommitAndReset]
//   - [IOGPUMetalCommandBuffer.CommitEncoder]
//   - [IOGPUMetalCommandBuffer.DidCompleteWithStartTimeEndTimeError]
//   - [IOGPUMetalCommandBuffer.DoCorruptCBSPI]
//   - [IOGPUMetalCommandBuffer.EncodeConditionalAbortEvent]
//   - [IOGPUMetalCommandBuffer.EncodeSignalEventValueAgentMask]
//   - [IOGPUMetalCommandBuffer.EncodeSignalEventScheduledValue]
//   - [IOGPUMetalCommandBuffer.EncodeSubmitSleepMS]
//   - [IOGPUMetalCommandBuffer.EncodeWaitForEventValueTimeout]
//   - [IOGPUMetalCommandBuffer.EndCurrentSegment]
//   - [IOGPUMetalCommandBuffer.FillCommandBufferArgsCommandQueue]
//   - [IOGPUMetalCommandBuffer.GetCurrentKernelCommandBufferPointerEnd]
//   - [IOGPUMetalCommandBuffer.GetCurrentKernelCommandBufferStartCurrentEnd]
//   - [IOGPUMetalCommandBuffer.GetDebugBufferPointerStartEnd]
//   - [IOGPUMetalCommandBuffer.GetSegmentListHeader]
//   - [IOGPUMetalCommandBuffer.GetSegmentListLockedPeerIndex]
//   - [IOGPUMetalCommandBuffer.GetSegmentListPointerStartCurrentEnd]
//   - [IOGPUMetalCommandBuffer.GpuFaultAddress]
//   - [IOGPUMetalCommandBuffer.GrowDebugBuffer]
//   - [IOGPUMetalCommandBuffer.GrowKernelCommandBuffer]
//   - [IOGPUMetalCommandBuffer.GrowSegmentList]
//   - [IOGPUMetalCommandBuffer.GrowSidebandBuffer]
//   - [IOGPUMetalCommandBuffer.IoGPUResourceList]
//   - [IOGPUMetalCommandBuffer.KernelCommandCollectTimeStamp]
//   - [IOGPUMetalCommandBuffer.KprintfBytesLength]
//   - [IOGPUMetalCommandBuffer.ProtectionOptions]
//   - [IOGPUMetalCommandBuffer.SetCurrentCommandEncoder]
//   - [IOGPUMetalCommandBuffer.SetCurrentKernelCommandBufferPointer]
//   - [IOGPUMetalCommandBuffer.SetCurrentSegmentListPointer]
//   - [IOGPUMetalCommandBuffer.SetProtectionOptions]
//   - [IOGPUMetalCommandBuffer.SetResponsibleTaskIDsCount]
//   - [IOGPUMetalCommandBuffer.SetSegmentListLockedPeerIndex]
//   - [IOGPUMetalCommandBuffer.UseInternalResidencySet]
//   - [IOGPUMetalCommandBuffer.UseInternalResidencySetsCount]
//   - [IOGPUMetalCommandBuffer.Validate]
//   - [IOGPUMetalCommandBuffer.InitWithQueueRetainedReferences]
//   - [IOGPUMetalCommandBuffer.InitWithQueueRetainedReferencesSynchronousDebugMode]
type IOGPUMetalCommandBuffer struct {
	metal.MTLCommandBufferObject
}

// IOGPUMetalCommandBufferFromID constructs a [IOGPUMetalCommandBuffer] from an objc.ID.
func IOGPUMetalCommandBufferFromID(id objc.ID) IOGPUMetalCommandBuffer {
	return IOGPUMetalCommandBuffer{MTLCommandBufferObject: metal.MTLCommandBufferObjectFromID(id)}
}

// Ensure IOGPUMetalCommandBuffer implements IIOGPUMetalCommandBuffer.
var _ IIOGPUMetalCommandBuffer = IOGPUMetalCommandBuffer{}

// An interface definition for the [IOGPUMetalCommandBuffer] class.
//
// # Methods
//
//   - [IIOGPUMetalCommandBuffer._debugBytesLengthOutput_type]
//   - [IIOGPUMetalCommandBuffer._encodePurgedResources]
//   - [IIOGPUMetalCommandBuffer._reserveKernelCommandBufferSpace]
//   - [IIOGPUMetalCommandBuffer.AddPurgedHeap]
//   - [IIOGPUMetalCommandBuffer.AddPurgedResource]
//   - [IIOGPUMetalCommandBuffer.AkPrivateResourceList]
//   - [IIOGPUMetalCommandBuffer.AkResourceList]
//   - [IIOGPUMetalCommandBuffer.AllocCommandBufferResourceAtIndex]
//   - [IIOGPUMetalCommandBuffer.AllocDebugBuffer]
//   - [IIOGPUMetalCommandBuffer.AllocateSidebandBuffer]
//   - [IIOGPUMetalCommandBuffer.BeginSegment]
//   - [IIOGPUMetalCommandBuffer.CommandBufferResourceInfo]
//   - [IIOGPUMetalCommandBuffer.CommandBufferStorage]
//   - [IIOGPUMetalCommandBuffer.CommitAndReset]
//   - [IIOGPUMetalCommandBuffer.CommitEncoder]
//   - [IIOGPUMetalCommandBuffer.DidCompleteWithStartTimeEndTimeError]
//   - [IIOGPUMetalCommandBuffer.DoCorruptCBSPI]
//   - [IIOGPUMetalCommandBuffer.EncodeConditionalAbortEvent]
//   - [IIOGPUMetalCommandBuffer.EncodeSignalEventValueAgentMask]
//   - [IIOGPUMetalCommandBuffer.EncodeSignalEventScheduledValue]
//   - [IIOGPUMetalCommandBuffer.EncodeSubmitSleepMS]
//   - [IIOGPUMetalCommandBuffer.EncodeWaitForEventValueTimeout]
//   - [IIOGPUMetalCommandBuffer.EndCurrentSegment]
//   - [IIOGPUMetalCommandBuffer.FillCommandBufferArgsCommandQueue]
//   - [IIOGPUMetalCommandBuffer.GetCurrentKernelCommandBufferPointerEnd]
//   - [IIOGPUMetalCommandBuffer.GetCurrentKernelCommandBufferStartCurrentEnd]
//   - [IIOGPUMetalCommandBuffer.GetDebugBufferPointerStartEnd]
//   - [IIOGPUMetalCommandBuffer.GetSegmentListHeader]
//   - [IIOGPUMetalCommandBuffer.GetSegmentListLockedPeerIndex]
//   - [IIOGPUMetalCommandBuffer.GetSegmentListPointerStartCurrentEnd]
//   - [IIOGPUMetalCommandBuffer.GpuFaultAddress]
//   - [IIOGPUMetalCommandBuffer.GrowDebugBuffer]
//   - [IIOGPUMetalCommandBuffer.GrowKernelCommandBuffer]
//   - [IIOGPUMetalCommandBuffer.GrowSegmentList]
//   - [IIOGPUMetalCommandBuffer.GrowSidebandBuffer]
//   - [IIOGPUMetalCommandBuffer.IoGPUResourceList]
//   - [IIOGPUMetalCommandBuffer.KernelCommandCollectTimeStamp]
//   - [IIOGPUMetalCommandBuffer.KprintfBytesLength]
//   - [IIOGPUMetalCommandBuffer.ProtectionOptions]
//   - [IIOGPUMetalCommandBuffer.SetCurrentCommandEncoder]
//   - [IIOGPUMetalCommandBuffer.SetCurrentKernelCommandBufferPointer]
//   - [IIOGPUMetalCommandBuffer.SetCurrentSegmentListPointer]
//   - [IIOGPUMetalCommandBuffer.SetProtectionOptions]
//   - [IIOGPUMetalCommandBuffer.SetResponsibleTaskIDsCount]
//   - [IIOGPUMetalCommandBuffer.SetSegmentListLockedPeerIndex]
//   - [IIOGPUMetalCommandBuffer.UseInternalResidencySet]
//   - [IIOGPUMetalCommandBuffer.UseInternalResidencySetsCount]
//   - [IIOGPUMetalCommandBuffer.Validate]
//   - [IIOGPUMetalCommandBuffer.InitWithQueueRetainedReferences]
//   - [IIOGPUMetalCommandBuffer.InitWithQueueRetainedReferencesSynchronousDebugMode]
type IIOGPUMetalCommandBuffer interface {
	metal.MTLCommandBuffer

	// Topic: Methods

	_debugBytesLengthOutput_type(bytes string, length uint64, output_type uint32)
	_encodePurgedResources()
	_reserveKernelCommandBufferSpace(space uint64) unsafe.Pointer
	AddPurgedHeap(heap objectivec.IObject)
	AddPurgedResource(resource objectivec.IObject)
	AkPrivateResourceList() objectivec.IObject
	AkResourceList() objectivec.IObject
	AllocCommandBufferResourceAtIndex(index uint32)
	AllocDebugBuffer()
	AllocateSidebandBuffer(buffer uint32)
	BeginSegment(segment unsafe.Pointer)
	CommandBufferResourceInfo() *IOGPUMetalCommandBufferResourceInfo
	CommandBufferStorage() *IOGPUMetalCommandBufferStorage
	CommitAndReset()
	CommitEncoder()
	DidCompleteWithStartTimeEndTimeError(time uint64, time2 uint64, error_ objectivec.IObject)
	DoCorruptCBSPI(cbspi int)
	EncodeConditionalAbortEvent(event objectivec.IObject)
	EncodeSignalEventValueAgentMask(event objectivec.IObject, value uint64, mask uint64)
	EncodeSignalEventScheduledValue(scheduled objectivec.IObject, value uint64)
	EncodeSubmitSleepMS(ms uint32)
	EncodeWaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint32)
	EndCurrentSegment()
	FillCommandBufferArgsCommandQueue(args *IOGPUCommandQueueCommandBufferArgs, queue objectivec.IObject)
	GetCurrentKernelCommandBufferPointerEnd(pointer unsafe.Pointer, end unsafe.Pointer)
	GetCurrentKernelCommandBufferStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer)
	GetDebugBufferPointerStartEnd(start unsafe.Pointer, end unsafe.Pointer)
	GetSegmentListHeader() *IOGPUSegmentListHeader
	GetSegmentListLockedPeerIndex(index *uint32) bool
	GetSegmentListPointerStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer)
	GpuFaultAddress() uint64
	GrowDebugBuffer(buffer uint32)
	GrowKernelCommandBuffer(buffer uint64)
	GrowSegmentList()
	GrowSidebandBuffer(buffer uint32)
	IoGPUResourceList() *IOGPUResourceList
	KernelCommandCollectTimeStamp()
	KprintfBytesLength(bytes string, length uint64)
	ProtectionOptions() uint64
	SetCurrentCommandEncoder(encoder objectivec.IObject)
	SetCurrentKernelCommandBufferPointer(pointer unsafe.Pointer)
	SetCurrentSegmentListPointer(pointer unsafe.Pointer)
	SetProtectionOptions(options uint64)
	SetResponsibleTaskIDsCount(iDs []uint32, count uint32)
	SetSegmentListLockedPeerIndex(index uint32)
	UseInternalResidencySet(set objectivec.IObject)
	UseInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	Validate()
	InitWithQueueRetainedReferences(queue objectivec.IObject, references bool) IOGPUMetalCommandBuffer
	InitWithQueueRetainedReferencesSynchronousDebugMode(queue objectivec.IObject, references bool, mode bool) IOGPUMetalCommandBuffer
}

// Init initializes the instance.
func (i IOGPUMetalCommandBuffer) Init() IOGPUMetalCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalCommandBuffer) Autorelease() IOGPUMetalCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalCommandBuffer creates a new IOGPUMetalCommandBuffer instance.
func NewIOGPUMetalCommandBuffer() IOGPUMetalCommandBuffer {
	class := getIOGPUMetalCommandBufferClass()
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalCommandBufferWithQueueRetainedReferences(queue objectivec.IObject, references bool) IOGPUMetalCommandBuffer {
	instance := getIOGPUMetalCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithQueue:retainedReferences:"), queue, references)
	return IOGPUMetalCommandBufferFromID(rv)
}

func NewGPUMetalCommandBufferWithQueueRetainedReferencesSynchronousDebugMode(queue objectivec.IObject, references bool, mode bool) IOGPUMetalCommandBuffer {
	instance := getIOGPUMetalCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithQueue:retainedReferences:synchronousDebugMode:"), queue, references, mode)
	return IOGPUMetalCommandBufferFromID(rv)
}

func (i IOGPUMetalCommandBuffer) _debugBytesLengthOutput_type(bytes string, length uint64, output_type uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_debugBytes:length:output_type:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length, output_type)
}

// DebugBytesLengthOutput_type is an exported wrapper for the private method _debugBytesLengthOutput_type.
func (i IOGPUMetalCommandBuffer) DebugBytesLengthOutput_type(bytes string, length uint64, output_type uint32) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_debugBytes:length:output_type:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_debugBytes:length:output_type:"}
		return err
	}
	i._debugBytesLengthOutput_type(bytes, length, output_type)
	return nil
}

// CanDebugBytesLengthOutput_type reports whether the receiver responds to the private selector _debugBytes:length:output_type:.
func (i IOGPUMetalCommandBuffer) CanDebugBytesLengthOutput_type() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_debugBytes:length:output_type:"))
}
func (i IOGPUMetalCommandBuffer) _encodePurgedResources() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_encodePurgedResources"))
}

// EncodePurgedResources is an exported wrapper for the private method _encodePurgedResources.
func (i IOGPUMetalCommandBuffer) EncodePurgedResources() error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodePurgedResources")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodePurgedResources"}
		return err
	}
	i._encodePurgedResources()
	return nil
}

// CanEncodePurgedResources reports whether the receiver responds to the private selector _encodePurgedResources.
func (i IOGPUMetalCommandBuffer) CanEncodePurgedResources() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodePurgedResources"))
}
func (i IOGPUMetalCommandBuffer) _reserveKernelCommandBufferSpace(space uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("_reserveKernelCommandBufferSpace:"), space)
	return rv
}

// ReserveKernelCommandBufferSpace is an exported wrapper for the private method _reserveKernelCommandBufferSpace.
func (i IOGPUMetalCommandBuffer) ReserveKernelCommandBufferSpace(space uint64) (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_reserveKernelCommandBufferSpace:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reserveKernelCommandBufferSpace:"}
		return nil, err
	}
	return i._reserveKernelCommandBufferSpace(space), nil
}

// CanReserveKernelCommandBufferSpace reports whether the receiver responds to the private selector _reserveKernelCommandBufferSpace:.
func (i IOGPUMetalCommandBuffer) CanReserveKernelCommandBufferSpace() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_reserveKernelCommandBufferSpace:"))
}
func (i IOGPUMetalCommandBuffer) AddPurgedHeap(heap objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addPurgedHeap:"), heap)
}
func (i IOGPUMetalCommandBuffer) AddPurgedResource(resource objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addPurgedResource:"), resource)
}
func (i IOGPUMetalCommandBuffer) AkPrivateResourceList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akPrivateResourceList"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalCommandBuffer) AkResourceList() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akResourceList"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalCommandBuffer) AllocCommandBufferResourceAtIndex(index uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocCommandBufferResourceAtIndex:"), index)
}
func (i IOGPUMetalCommandBuffer) AllocDebugBuffer() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocDebugBuffer"))
}
func (i IOGPUMetalCommandBuffer) AllocateSidebandBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocateSidebandBuffer:"), buffer)
}
func (i IOGPUMetalCommandBuffer) BeginSegment(segment unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("beginSegment:"), segment)
}
func (i IOGPUMetalCommandBuffer) CommitAndReset() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitAndReset"))
}
func (i IOGPUMetalCommandBuffer) CommitEncoder() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitEncoder"))
}
func (i IOGPUMetalCommandBuffer) DidCompleteWithStartTimeEndTimeError(time uint64, time2 uint64, error_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("didCompleteWithStartTime:endTime:error:"), time, time2, error_)
}
func (i IOGPUMetalCommandBuffer) DoCorruptCBSPI(cbspi int) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("doCorruptCBSPI:"), cbspi)
}
func (i IOGPUMetalCommandBuffer) EncodeConditionalAbortEvent(event objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeConditionalAbortEvent:"), event)
}
func (i IOGPUMetalCommandBuffer) EncodeSignalEventValueAgentMask(event objectivec.IObject, value uint64, mask uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeSignalEvent:value:agentMask:"), event, value, mask)
}
func (i IOGPUMetalCommandBuffer) EncodeSignalEventScheduledValue(scheduled objectivec.IObject, value uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeSignalEventScheduled:value:"), scheduled, value)
}
func (i IOGPUMetalCommandBuffer) EncodeSubmitSleepMS(ms uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeSubmitSleepMS:"), ms)
}
func (i IOGPUMetalCommandBuffer) EncodeWaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("encodeWaitForEvent:value:timeout:"), event, value, timeout)
}
func (i IOGPUMetalCommandBuffer) EndCurrentSegment() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("endCurrentSegment"))
}
func (i IOGPUMetalCommandBuffer) FillCommandBufferArgsCommandQueue(args *IOGPUCommandQueueCommandBufferArgs, queue objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("fillCommandBufferArgs:commandQueue:"), unsafe.Pointer(args), queue)
}
func (i IOGPUMetalCommandBuffer) GetCurrentKernelCommandBufferPointerEnd(pointer unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getCurrentKernelCommandBufferPointer:end:"), pointer, end)
}
func (i IOGPUMetalCommandBuffer) GetCurrentKernelCommandBufferStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getCurrentKernelCommandBufferStart:current:end:"), start, current, end)
}
func (i IOGPUMetalCommandBuffer) GetDebugBufferPointerStartEnd(start unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getDebugBufferPointerStart:end:"), start, end)
}
func (i IOGPUMetalCommandBuffer) GetSegmentListHeader() *IOGPUSegmentListHeader {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getSegmentListHeader"))
	return (*IOGPUSegmentListHeader)(rv)
}
func (i IOGPUMetalCommandBuffer) GetSegmentListLockedPeerIndex(index *uint32) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("getSegmentListLockedPeerIndex:"), unsafe.Pointer(index))
	return rv
}
func (i IOGPUMetalCommandBuffer) GetSegmentListPointerStartCurrentEnd(start unsafe.Pointer, current unsafe.Pointer, end unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getSegmentListPointerStart:current:end:"), start, current, end)
}
func (i IOGPUMetalCommandBuffer) GrowDebugBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growDebugBuffer:"), buffer)
}
func (i IOGPUMetalCommandBuffer) GrowKernelCommandBuffer(buffer uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growKernelCommandBuffer:"), buffer)
}
func (i IOGPUMetalCommandBuffer) GrowSegmentList() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growSegmentList"))
}
func (i IOGPUMetalCommandBuffer) GrowSidebandBuffer(buffer uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("growSidebandBuffer:"), buffer)
}
func (i IOGPUMetalCommandBuffer) KernelCommandCollectTimeStamp() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("kernelCommandCollectTimeStamp"))
}
func (i IOGPUMetalCommandBuffer) KprintfBytesLength(bytes string, length uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("kprintfBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
}
func (i IOGPUMetalCommandBuffer) ProtectionOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("protectionOptions"))
	return rv
}
func (i IOGPUMetalCommandBuffer) SetCurrentCommandEncoder(encoder objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentCommandEncoder:"), encoder)
}
func (i IOGPUMetalCommandBuffer) SetCurrentKernelCommandBufferPointer(pointer unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentKernelCommandBufferPointer:"), pointer)
}
func (i IOGPUMetalCommandBuffer) SetCurrentSegmentListPointer(pointer unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentSegmentListPointer:"), pointer)
}
func (i IOGPUMetalCommandBuffer) SetProtectionOptions(options uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setProtectionOptions:"), options)
}
func (i IOGPUMetalCommandBuffer) SetResponsibleTaskIDsCount(iDs []uint32, count uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setResponsibleTaskIDs:count:"), objc.CArray(iDs), count)
}
func (i IOGPUMetalCommandBuffer) SetSegmentListLockedPeerIndex(index uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setSegmentListLockedPeerIndex:"), index)
}
func (i IOGPUMetalCommandBuffer) UseInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useInternalResidencySet:"), set)
}
func (i IOGPUMetalCommandBuffer) UseInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetalCommandBuffer) Validate() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("validate"))
}
func (i IOGPUMetalCommandBuffer) InitWithQueueRetainedReferences(queue objectivec.IObject, references bool) IOGPUMetalCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](i.ID, objc.Sel("initWithQueue:retainedReferences:"), queue, references)
	return rv
}
func (i IOGPUMetalCommandBuffer) InitWithQueueRetainedReferencesSynchronousDebugMode(queue objectivec.IObject, references bool, mode bool) IOGPUMetalCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalCommandBuffer](i.ID, objc.Sel("initWithQueue:retainedReferences:synchronousDebugMode:"), queue, references, mode)
	return rv
}

func (i IOGPUMetalCommandBuffer) CommandBufferResourceInfo() *IOGPUMetalCommandBufferResourceInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("commandBufferResourceInfo"))
	return (*IOGPUMetalCommandBufferResourceInfo)(rv)
}
func (i IOGPUMetalCommandBuffer) CommandBufferStorage() *IOGPUMetalCommandBufferStorage {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("commandBufferStorage"))
	return (*IOGPUMetalCommandBufferStorage)(rv)
}
func (i IOGPUMetalCommandBuffer) GpuFaultAddress() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuFaultAddress"))
	return rv
}
func (i IOGPUMetalCommandBuffer) IoGPUResourceList() *IOGPUResourceList {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("ioGPUResourceList"))
	return (*IOGPUResourceList)(rv)
}
