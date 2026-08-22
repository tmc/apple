// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIOCommandQueue] class.
var (
	_IOGPUMetalIOCommandQueueClass     IOGPUMetalIOCommandQueueClass
	_IOGPUMetalIOCommandQueueClassOnce sync.Once
)

func getIOGPUMetalIOCommandQueueClass() IOGPUMetalIOCommandQueueClass {
	_IOGPUMetalIOCommandQueueClassOnce.Do(func() {
		_IOGPUMetalIOCommandQueueClass = IOGPUMetalIOCommandQueueClass{class: objc.GetClass("IOGPUMetalIOCommandQueue")}
	})
	return _IOGPUMetalIOCommandQueueClass
}

// GetIOGPUMetalIOCommandQueueClass returns the class object for IOGPUMetalIOCommandQueue.
func GetIOGPUMetalIOCommandQueueClass() IOGPUMetalIOCommandQueueClass {
	return getIOGPUMetalIOCommandQueueClass()
}

type IOGPUMetalIOCommandQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIOCommandQueueClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIOCommandQueueClass) Alloc() IOGPUMetalIOCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandQueue](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIOCommandQueue._submitAvailableCommands]
//   - [IOGPUMetalIOCommandQueue.Barrier]
//   - [IOGPUMetalIOCommandQueue.CommandBufferComplete]
//   - [IOGPUMetalIOCommandQueue.CommitCommandBuffer]
//   - [IOGPUMetalIOCommandQueue.DidCompleteWithStatus]
//   - [IOGPUMetalIOCommandQueue.EnqueueCommandBuffer]
//   - [IOGPUMetalIOCommandQueue.GetDecompressionQueue]
//   - [IOGPUMetalIOCommandQueue.GetPriority]
//   - [IOGPUMetalIOCommandQueue.GlobalTraceObjectID]
//   - [IOGPUMetalIOCommandQueue.LaunchIOGPUIOThreads]
//   - [IOGPUMetalIOCommandQueue.LaunchIOWorkerThreads]
//   - [IOGPUMetalIOCommandQueue.SubmitAvailableCommands]
//   - [IOGPUMetalIOCommandQueue.InitWithDeviceDescriptor]
type IOGPUMetalIOCommandQueue struct {
	metal.MTLIOCommandQueueObject
}

// IOGPUMetalIOCommandQueueFromID constructs a [IOGPUMetalIOCommandQueue] from an objc.ID.
func IOGPUMetalIOCommandQueueFromID(id objc.ID) IOGPUMetalIOCommandQueue {
	return IOGPUMetalIOCommandQueue{MTLIOCommandQueueObject: metal.MTLIOCommandQueueObjectFromID(id)}
}

// Ensure IOGPUMetalIOCommandQueue implements IIOGPUMetalIOCommandQueue.
var _ IIOGPUMetalIOCommandQueue = IOGPUMetalIOCommandQueue{}

// An interface definition for the [IOGPUMetalIOCommandQueue] class.
//
// # Methods
//
//   - [IIOGPUMetalIOCommandQueue._submitAvailableCommands]
//   - [IIOGPUMetalIOCommandQueue.Barrier]
//   - [IIOGPUMetalIOCommandQueue.CommandBufferComplete]
//   - [IIOGPUMetalIOCommandQueue.CommitCommandBuffer]
//   - [IIOGPUMetalIOCommandQueue.DidCompleteWithStatus]
//   - [IIOGPUMetalIOCommandQueue.EnqueueCommandBuffer]
//   - [IIOGPUMetalIOCommandQueue.GetDecompressionQueue]
//   - [IIOGPUMetalIOCommandQueue.GetPriority]
//   - [IIOGPUMetalIOCommandQueue.GlobalTraceObjectID]
//   - [IIOGPUMetalIOCommandQueue.LaunchIOGPUIOThreads]
//   - [IIOGPUMetalIOCommandQueue.LaunchIOWorkerThreads]
//   - [IIOGPUMetalIOCommandQueue.SubmitAvailableCommands]
//   - [IIOGPUMetalIOCommandQueue.InitWithDeviceDescriptor]
type IIOGPUMetalIOCommandQueue interface {
	metal.MTLIOCommandQueue

	// Topic: Methods

	_submitAvailableCommands(commands unsafe.Pointer)
	Barrier()
	CommandBufferComplete()
	CommitCommandBuffer(buffer objectivec.IObject)
	DidCompleteWithStatus(complete objectivec.IObject, status int64)
	EnqueueCommandBuffer(buffer objectivec.IObject)
	GetDecompressionQueue() objectivec.IObject
	GetPriority() int64
	GlobalTraceObjectID() uint64
	LaunchIOGPUIOThreads()
	LaunchIOWorkerThreads()
	SubmitAvailableCommands()
	InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalIOCommandQueue
}

// Init initializes the instance.
func (i IOGPUMetalIOCommandQueue) Init() IOGPUMetalIOCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandQueue](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIOCommandQueue) Autorelease() IOGPUMetalIOCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandQueue](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIOCommandQueue creates a new IOGPUMetalIOCommandQueue instance.
func NewIOGPUMetalIOCommandQueue() IOGPUMetalIOCommandQueue {
	class := getIOGPUMetalIOCommandQueueClass()
	rv := objc.SendIfResponds[IOGPUMetalIOCommandQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIOCommandQueueWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalIOCommandQueue {
	instance := getIOGPUMetalIOCommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return IOGPUMetalIOCommandQueueFromID(rv)
}

func (i IOGPUMetalIOCommandQueue) _submitAvailableCommands(commands unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_submitAvailableCommands:"), commands)
}
func (i IOGPUMetalIOCommandQueue) Barrier() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("barrier"))
}
func (i IOGPUMetalIOCommandQueue) CommandBufferComplete() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commandBufferComplete"))
}
func (i IOGPUMetalIOCommandQueue) CommitCommandBuffer(buffer objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commitCommandBuffer:"), buffer)
}
func (i IOGPUMetalIOCommandQueue) DidCompleteWithStatus(complete objectivec.IObject, status int64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("didComplete:withStatus:"), complete, status)
}
func (i IOGPUMetalIOCommandQueue) EnqueueCommandBuffer(buffer objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("enqueueCommandBuffer:"), buffer)
}
func (i IOGPUMetalIOCommandQueue) GetDecompressionQueue() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getDecompressionQueue"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalIOCommandQueue) GetPriority() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("getPriority"))
	return rv
}
func (i IOGPUMetalIOCommandQueue) GlobalTraceObjectID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalTraceObjectID"))
	return rv
}
func (i IOGPUMetalIOCommandQueue) LaunchIOGPUIOThreads() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("launchIOGPUIOThreads"))
}
func (i IOGPUMetalIOCommandQueue) LaunchIOWorkerThreads() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("launchIOWorkerThreads"))
}
func (i IOGPUMetalIOCommandQueue) SubmitAvailableCommands() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("submitAvailableCommands"))
}
func (i IOGPUMetalIOCommandQueue) InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalIOCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalIOCommandQueue](i.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}
