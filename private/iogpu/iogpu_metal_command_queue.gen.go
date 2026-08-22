// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalCommandQueue] class.
var (
	_IOGPUMetalCommandQueueClass     IOGPUMetalCommandQueueClass
	_IOGPUMetalCommandQueueClassOnce sync.Once
)

func getIOGPUMetalCommandQueueClass() IOGPUMetalCommandQueueClass {
	_IOGPUMetalCommandQueueClassOnce.Do(func() {
		_IOGPUMetalCommandQueueClass = IOGPUMetalCommandQueueClass{class: objc.GetClass("IOGPUMetalCommandQueue")}
	})
	return _IOGPUMetalCommandQueueClass
}

// GetIOGPUMetalCommandQueueClass returns the class object for IOGPUMetalCommandQueue.
func GetIOGPUMetalCommandQueueClass() IOGPUMetalCommandQueueClass {
	return getIOGPUMetalCommandQueueClass()
}

type IOGPUMetalCommandQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalCommandQueueClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalCommandQueueClass) Alloc() IOGPUMetalCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalCommandQueue._setGPUPriorityBackgroundPriority]
//   - [IOGPUMetalCommandQueue._submitCommandBuffersCount]
//   - [IOGPUMetalCommandQueue.AddInternalResidencySet]
//   - [IOGPUMetalCommandQueue.AddInternalResidencySetsCount]
//   - [IOGPUMetalCommandQueue.DispatchAvailableCompletionNotifications]
//   - [IOGPUMetalCommandQueue.GetBackgroundGPUPriority]
//   - [IOGPUMetalCommandQueue.GetGPUPriority]
//   - [IOGPUMetalCommandQueue.RemoveInternalResidencySet]
//   - [IOGPUMetalCommandQueue.RemoveInternalResidencySetsCount]
//   - [IOGPUMetalCommandQueue.SetBackgroundGPUPriority]
//   - [IOGPUMetalCommandQueue.SetBackgroundGPUPriorityOffset]
//   - [IOGPUMetalCommandQueue.SetCompletionQueue]
//   - [IOGPUMetalCommandQueue.SetGPUPriority]
//   - [IOGPUMetalCommandQueue.SetGPUPriorityOffset]
//   - [IOGPUMetalCommandQueue.SubmitCommandBuffersCount]
//   - [IOGPUMetalCommandQueue.SupportsBackgroundAppRole]
//   - [IOGPUMetalCommandQueue.InitWithDeviceDescriptor]
//   - [IOGPUMetalCommandQueue.InitWithDeviceDescriptorArgsArgsSize]
type IOGPUMetalCommandQueue struct {
	metal.MTLCommandQueueObject
}

// IOGPUMetalCommandQueueFromID constructs a [IOGPUMetalCommandQueue] from an objc.ID.
func IOGPUMetalCommandQueueFromID(id objc.ID) IOGPUMetalCommandQueue {
	return IOGPUMetalCommandQueue{MTLCommandQueueObject: metal.MTLCommandQueueObjectFromID(id)}
}

// Ensure IOGPUMetalCommandQueue implements IIOGPUMetalCommandQueue.
var _ IIOGPUMetalCommandQueue = IOGPUMetalCommandQueue{}

// An interface definition for the [IOGPUMetalCommandQueue] class.
//
// # Methods
//
//   - [IIOGPUMetalCommandQueue._setGPUPriorityBackgroundPriority]
//   - [IIOGPUMetalCommandQueue._submitCommandBuffersCount]
//   - [IIOGPUMetalCommandQueue.AddInternalResidencySet]
//   - [IIOGPUMetalCommandQueue.AddInternalResidencySetsCount]
//   - [IIOGPUMetalCommandQueue.DispatchAvailableCompletionNotifications]
//   - [IIOGPUMetalCommandQueue.GetBackgroundGPUPriority]
//   - [IIOGPUMetalCommandQueue.GetGPUPriority]
//   - [IIOGPUMetalCommandQueue.RemoveInternalResidencySet]
//   - [IIOGPUMetalCommandQueue.RemoveInternalResidencySetsCount]
//   - [IIOGPUMetalCommandQueue.SetBackgroundGPUPriority]
//   - [IIOGPUMetalCommandQueue.SetBackgroundGPUPriorityOffset]
//   - [IIOGPUMetalCommandQueue.SetCompletionQueue]
//   - [IIOGPUMetalCommandQueue.SetGPUPriority]
//   - [IIOGPUMetalCommandQueue.SetGPUPriorityOffset]
//   - [IIOGPUMetalCommandQueue.SubmitCommandBuffersCount]
//   - [IIOGPUMetalCommandQueue.SupportsBackgroundAppRole]
//   - [IIOGPUMetalCommandQueue.InitWithDeviceDescriptor]
//   - [IIOGPUMetalCommandQueue.InitWithDeviceDescriptorArgsArgsSize]
type IIOGPUMetalCommandQueue interface {
	metal.MTLCommandQueue

	// Topic: Methods

	_setGPUPriorityBackgroundPriority(gPUPriority uint64, priority uint64) bool
	_submitCommandBuffersCount(buffers []objectivec.IObject, count uint64)
	AddInternalResidencySet(set objectivec.IObject)
	AddInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	DispatchAvailableCompletionNotifications()
	GetBackgroundGPUPriority() uint64
	GetGPUPriority() uint64
	RemoveInternalResidencySet(set objectivec.IObject)
	RemoveInternalResidencySetsCount(sets []objectivec.IObject, count uint64)
	SetBackgroundGPUPriority(gPUPriority uint64) bool
	SetBackgroundGPUPriorityOffset(gPUPriority uint64, offset uint16) bool
	SetCompletionQueue(queue objectivec.IObject)
	SetGPUPriority(gPUPriority uint64) bool
	SetGPUPriorityOffset(gPUPriority uint64, offset uint16) bool
	SubmitCommandBuffersCount(buffers []objectivec.IObject, count uint64)
	SupportsBackgroundAppRole() bool
	InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalCommandQueue
	InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetalCommandQueue
}

// Init initializes the instance.
func (i IOGPUMetalCommandQueue) Init() IOGPUMetalCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalCommandQueue) Autorelease() IOGPUMetalCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalCommandQueue creates a new IOGPUMetalCommandQueue instance.
func NewIOGPUMetalCommandQueue() IOGPUMetalCommandQueue {
	class := getIOGPUMetalCommandQueueClass()
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalCommandQueueWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalCommandQueue {
	instance := getIOGPUMetalCommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return IOGPUMetalCommandQueueFromID(rv)
}

func NewGPUMetalCommandQueueWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetalCommandQueue {
	instance := getIOGPUMetalCommandQueueClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return IOGPUMetalCommandQueueFromID(rv)
}

func (i IOGPUMetalCommandQueue) _setGPUPriorityBackgroundPriority(gPUPriority uint64, priority uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("_setGPUPriority:backgroundPriority:"), gPUPriority, priority)
	return rv
}

// SetGPUPriorityBackgroundPriority is an exported wrapper for the private method _setGPUPriorityBackgroundPriority.
func (i IOGPUMetalCommandQueue) SetGPUPriorityBackgroundPriority(gPUPriority uint64, priority uint64) (bool, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_setGPUPriority:backgroundPriority:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGPUPriority:backgroundPriority:"}
		return false, err
	}
	return i._setGPUPriorityBackgroundPriority(gPUPriority, priority), nil
}

// CanSetGPUPriorityBackgroundPriority reports whether the receiver responds to the private selector _setGPUPriority:backgroundPriority:.
func (i IOGPUMetalCommandQueue) CanSetGPUPriorityBackgroundPriority() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_setGPUPriority:backgroundPriority:"))
}
func (i IOGPUMetalCommandQueue) _submitCommandBuffersCount(buffers []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_submitCommandBuffers:count:"), objc.CArray(buffers), count)
}
func (i IOGPUMetalCommandQueue) AddInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addInternalResidencySet:"), set)
}
func (i IOGPUMetalCommandQueue) AddInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetalCommandQueue) DispatchAvailableCompletionNotifications() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("dispatchAvailableCompletionNotifications"))
}
func (i IOGPUMetalCommandQueue) GetBackgroundGPUPriority() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getBackgroundGPUPriority"))
	return rv
}
func (i IOGPUMetalCommandQueue) GetGPUPriority() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getGPUPriority"))
	return rv
}
func (i IOGPUMetalCommandQueue) RemoveInternalResidencySet(set objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeInternalResidencySet:"), set)
}
func (i IOGPUMetalCommandQueue) RemoveInternalResidencySetsCount(sets []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeInternalResidencySets:count:"), objc.CArray(sets), count)
}
func (i IOGPUMetalCommandQueue) SetBackgroundGPUPriority(gPUPriority uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("setBackgroundGPUPriority:"), gPUPriority)
	return rv
}
func (i IOGPUMetalCommandQueue) SetBackgroundGPUPriorityOffset(gPUPriority uint64, offset uint16) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("setBackgroundGPUPriority:offset:"), gPUPriority, offset)
	return rv
}
func (i IOGPUMetalCommandQueue) SetCompletionQueue(queue objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCompletionQueue:"), queue)
}
func (i IOGPUMetalCommandQueue) SetGPUPriority(gPUPriority uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("setGPUPriority:"), gPUPriority)
	return rv
}
func (i IOGPUMetalCommandQueue) SetGPUPriorityOffset(gPUPriority uint64, offset uint16) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("setGPUPriority:offset:"), gPUPriority, offset)
	return rv
}
func (i IOGPUMetalCommandQueue) SubmitCommandBuffersCount(buffers []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("submitCommandBuffers:count:"), objc.CArray(buffers), count)
}
func (i IOGPUMetalCommandQueue) SupportsBackgroundAppRole() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsBackgroundAppRole"))
	return rv
}
func (i IOGPUMetalCommandQueue) InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](i.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}
func (i IOGPUMetalCommandQueue) InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUDeviceNewCommandQueueArgs, size uint32) IOGPUMetalCommandQueue {
	rv := objc.SendIfResponds[IOGPUMetalCommandQueue](i.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return rv
}
