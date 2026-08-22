// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4CommandQueue] class.
var (
	_IOGPUMetal4CommandQueueClass     IOGPUMetal4CommandQueueClass
	_IOGPUMetal4CommandQueueClassOnce sync.Once
)

func getIOGPUMetal4CommandQueueClass() IOGPUMetal4CommandQueueClass {
	_IOGPUMetal4CommandQueueClassOnce.Do(func() {
		_IOGPUMetal4CommandQueueClass = IOGPUMetal4CommandQueueClass{class: objc.GetClass("IOGPUMetal4CommandQueue")}
	})
	return _IOGPUMetal4CommandQueueClass
}

// GetIOGPUMetal4CommandQueueClass returns the class object for IOGPUMetal4CommandQueue.
func GetIOGPUMetal4CommandQueueClass() IOGPUMetal4CommandQueueClass {
	return getIOGPUMetal4CommandQueueClass()
}

type IOGPUMetal4CommandQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4CommandQueueClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4CommandQueueClass) Alloc() IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4CommandQueue._commitCountCommitFeedback]
//   - [IOGPUMetal4CommandQueue.AddInternalResidencySet]
//   - [IOGPUMetal4CommandQueue.AddInternalResidencySetsCount]
//   - [IOGPUMetal4CommandQueue.AllocateMappingCommandBuffer]
//   - [IOGPUMetal4CommandQueue.CommandAllocator]
//   - [IOGPUMetal4CommandQueue.CommandBufferCompleteCommandAllocatorCommandAllocatorGenerationStorageSubmissionIDStartTimeCompletionTimeKernelErrorErrorCommitFeedbackCommitComplete]
//   - [IOGPUMetal4CommandQueue.CommitFillArgsCountArgsArgsSizeCommitFeedback]
//   - [IOGPUMetal4CommandQueue.CommitMappingCommandBuffer]
//   - [IOGPUMetal4CommandQueue.EndTier1MappingCommands]
//   - [IOGPUMetal4CommandQueue.PreCommitCountOptions]
//   - [IOGPUMetal4CommandQueue.RemoveInternalResidencySet]
//   - [IOGPUMetal4CommandQueue.RemoveInternalResidencySetsCount]
//   - [IOGPUMetal4CommandQueue.SetScheduledHandler]
//   - [IOGPUMetal4CommandQueue.SupportsBackgroundAppRole]
//   - [IOGPUMetal4CommandQueue.WaitForEventValueTimeout]
//   - [IOGPUMetal4CommandQueue.InitIOGPUMTL4CommandQueueDescriptorArgsArgsSize]
//   - [IOGPUMetal4CommandQueue.InitWithDevice]
//   - [IOGPUMetal4CommandQueue.InitWithDeviceDescriptor]
//   - [IOGPUMetal4CommandQueue.InitWithDeviceDescriptorArgsArgsSize]
type IOGPUMetal4CommandQueue struct {
	metal.MTL4CommandQueueObject
}

// IOGPUMetal4CommandQueueFromID constructs a [IOGPUMetal4CommandQueue] from an objc.ID.
func IOGPUMetal4CommandQueueFromID(id objc.ID) IOGPUMetal4CommandQueue {
	return IOGPUMetal4CommandQueue{MTL4CommandQueueObject: metal.MTL4CommandQueueObjectFromID(id)}
}

// Ensure IOGPUMetal4CommandQueue implements IIOGPUMetal4CommandQueue.
var _ IIOGPUMetal4CommandQueue = IOGPUMetal4CommandQueue{}

// An interface definition for the [IOGPUMetal4CommandQueue] class.
//
// # Methods
//
//   - [IIOGPUMetal4CommandQueue._commitCountCommitFeedback]
//   - [IIOGPUMetal4CommandQueue.AddInternalResidencySet]
//   - [IIOGPUMetal4CommandQueue.AddInternalResidencySetsCount]
//   - [IIOGPUMetal4CommandQueue.AllocateMappingCommandBuffer]
//   - [IIOGPUMetal4CommandQueue.CommandAllocator]
//   - [IIOGPUMetal4CommandQueue.CommandBufferCompleteCommandAllocatorCommandAllocatorGenerationStorageSubmissionIDStartTimeCompletionTimeKernelErrorErrorCommitFeedbackCommitComplete]
//   - [IIOGPUMetal4CommandQueue.CommitFillArgsCountArgsArgsSizeCommitFeedback]
//   - [IIOGPUMetal4CommandQueue.CommitMappingCommandBuffer]
//   - [IIOGPUMetal4CommandQueue.EndTier1MappingCommands]
//   - [IIOGPUMetal4CommandQueue.PreCommitCountOptions]
//   - [IIOGPUMetal4CommandQueue.RemoveInternalResidencySet]
//   - [IIOGPUMetal4CommandQueue.RemoveInternalResidencySetsCount]
//   - [IIOGPUMetal4CommandQueue.SetScheduledHandler]
//   - [IIOGPUMetal4CommandQueue.SupportsBackgroundAppRole]
//   - [IIOGPUMetal4CommandQueue.WaitForEventValueTimeout]
//   - [IIOGPUMetal4CommandQueue.InitIOGPUMTL4CommandQueueDescriptorArgsArgsSize]
//   - [IIOGPUMetal4CommandQueue.InitWithDevice]
//   - [IIOGPUMetal4CommandQueue.InitWithDeviceDescriptor]
//   - [IIOGPUMetal4CommandQueue.InitWithDeviceDescriptorArgsArgsSize]
type IIOGPUMetal4CommandQueue interface {
	metal.MTL4CommandQueue

	// Topic: Methods

	_commitCountCommitFeedback(_commit []objectivec.IObject, count uint64, feedback objectivec.IObject)
	AddInternalResidencySet(set objectivec.IObject)
	AddInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	AllocateMappingCommandBuffer()
	CommandAllocator() objectivec.IObject
	CommandBufferCompleteCommandAllocatorCommandAllocatorGenerationStorageSubmissionIDStartTimeCompletionTimeKernelErrorErrorCommitFeedbackCommitComplete(complete objectivec.IObject, allocator objectivec.IObject, generation uint32, storage *IOGPUMetalCommandBufferStorage, id uint64, time uint64, time2 uint64, kernelError uint32, underlyingError objectivec.IObject, feedback objectivec.IObject, complete2 bool)
	CommitFillArgsCountArgsArgsSizeCommitFeedback(args []objectivec.IObject, count uint64, args2 uint64, size uint32, feedback objectivec.IObject)
	CommitMappingCommandBuffer()
	EndTier1MappingCommands()
	PreCommitCountOptions(commit []objectivec.IObject, count uint64, options objectivec.IObject) objectivec.IObject
	RemoveInternalResidencySet(set objectivec.IObject)
	RemoveInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	SetScheduledHandler(handler VoidHandler)
	SupportsBackgroundAppRole() bool
	WaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint16)
	InitIOGPUMTL4CommandQueueDescriptorArgsArgsSize(queue objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue
	InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandQueue
	InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandQueue
	InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue
}

// Init initializes the instance.
func (i IOGPUMetal4CommandQueue) Init() IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4CommandQueue) Autorelease() IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4CommandQueue creates a new IOGPUMetal4CommandQueue instance.
func NewIOGPUMetal4CommandQueue() IOGPUMetal4CommandQueue {
	class := getIOGPUMetal4CommandQueueClass()
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4CommandQueueIOGPUMTL4CommandQueueDescriptorArgsArgsSize(queue objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue {
	instance := getIOGPUMetal4CommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initIOGPUMTL4CommandQueue:descriptor:args:argsSize:"), queue, descriptor, unsafe.Pointer(args), size)
	return IOGPUMetal4CommandQueueFromID(rv)
}

func NewGPUMetal4CommandQueueWithDevice(device objectivec.IObject) IOGPUMetal4CommandQueue {
	instance := getIOGPUMetal4CommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetal4CommandQueueFromID(rv)
}

func NewGPUMetal4CommandQueueWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandQueue {
	instance := getIOGPUMetal4CommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return IOGPUMetal4CommandQueueFromID(rv)
}

func NewGPUMetal4CommandQueueWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue {
	instance := getIOGPUMetal4CommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return IOGPUMetal4CommandQueueFromID(rv)
}

func (i IOGPUMetal4CommandQueue) _commitCountCommitFeedback(_commit []objectivec.IObject, count uint64, feedback objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_commit:count:commitFeedback:"), objc.CArray(_commit), count, feedback)
}

// CommitCountCommitFeedback is an exported wrapper for the private method _commitCountCommitFeedback.
func (i IOGPUMetal4CommandQueue) CommitCountCommitFeedback(_commit []objectivec.IObject, count uint64, feedback objectivec.IObject) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_commit:count:commitFeedback:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_commit:count:commitFeedback:"}
		return err
	}
	i._commitCountCommitFeedback(_commit, count, feedback)
	return nil
}

// CanCommitCountCommitFeedback reports whether the receiver responds to the private selector _commit:count:commitFeedback:.
func (i IOGPUMetal4CommandQueue) CanCommitCountCommitFeedback() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_commit:count:commitFeedback:"))
}
func (i IOGPUMetal4CommandQueue) AddInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addInternalResidencySet:"), set)
}
func (i IOGPUMetal4CommandQueue) AddInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetal4CommandQueue) AllocateMappingCommandBuffer() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocateMappingCommandBuffer"))
}
func (i IOGPUMetal4CommandQueue) CommandAllocator() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commandAllocator"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetal4CommandQueue) CommandBufferCompleteCommandAllocatorCommandAllocatorGenerationStorageSubmissionIDStartTimeCompletionTimeKernelErrorErrorCommitFeedbackCommitComplete(complete objectivec.IObject, allocator objectivec.IObject, generation uint32, storage *IOGPUMetalCommandBufferStorage, id uint64, time uint64, time2 uint64, kernelError uint32, underlyingError objectivec.IObject, feedback objectivec.IObject, complete2 bool) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commandBufferComplete:commandAllocator:commandAllocatorGeneration:storage:submissionID:startTime:completionTime:kernelError:error:commitFeedback:commitComplete:"), complete, allocator, generation, unsafe.Pointer(storage), id, time, time2, kernelError, underlyingError, feedback, complete2)
}
func (i IOGPUMetal4CommandQueue) CommitFillArgsCountArgsArgsSizeCommitFeedback(args []objectivec.IObject, count uint64, args2 uint64, size uint32, feedback objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitFillArgs:count:args:argsSize:commitFeedback:"), objc.CArray(args), count, args2, size, feedback)
}
func (i IOGPUMetal4CommandQueue) CommitMappingCommandBuffer() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitMappingCommandBuffer"))
}
func (i IOGPUMetal4CommandQueue) EndTier1MappingCommands() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("endTier1MappingCommands"))
}
func (i IOGPUMetal4CommandQueue) PreCommitCountOptions(commit []objectivec.IObject, count uint64, options objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("preCommit:count:options:"), objc.CArray(commit), count, options)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetal4CommandQueue) RemoveInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeInternalResidencySet:"), set)
}
func (i IOGPUMetal4CommandQueue) RemoveInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetal4CommandQueue) SetScheduledHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setScheduledHandler:"), _block0)
}
func (i IOGPUMetal4CommandQueue) SupportsBackgroundAppRole() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsBackgroundAppRole"))
	return rv
}
func (i IOGPUMetal4CommandQueue) WaitForEventValueTimeout(event objectivec.IObject, value uint64, timeout uint16) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("waitForEvent:value:timeout:"), event, value, timeout)
}
func (i IOGPUMetal4CommandQueue) InitIOGPUMTL4CommandQueueDescriptorArgsArgsSize(queue objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("initIOGPUMTL4CommandQueue:descriptor:args:argsSize:"), queue, descriptor, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetal4CommandQueue) InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
func (i IOGPUMetal4CommandQueue) InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}
func (i IOGPUMetal4CommandQueue) InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetal4CommandQueue {
	rv := objc.SendIfResponds[IOGPUMetal4CommandQueue](i.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return rv
}

// SetScheduledHandlerSync is a synchronous wrapper around [IOGPUMetal4CommandQueue.SetScheduledHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetal4CommandQueue) SetScheduledHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	i.SetScheduledHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
