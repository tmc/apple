// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalResidencySet] class.
var (
	_IOGPUMetalResidencySetClass     IOGPUMetalResidencySetClass
	_IOGPUMetalResidencySetClassOnce sync.Once
)

func getIOGPUMetalResidencySetClass() IOGPUMetalResidencySetClass {
	_IOGPUMetalResidencySetClassOnce.Do(func() {
		_IOGPUMetalResidencySetClass = IOGPUMetalResidencySetClass{class: objc.GetClass("IOGPUMetalResidencySet")}
	})
	return _IOGPUMetalResidencySetClass
}

// GetIOGPUMetalResidencySetClass returns the class object for IOGPUMetalResidencySet.
func GetIOGPUMetalResidencySetClass() IOGPUMetalResidencySetClass {
	return getIOGPUMetalResidencySetClass()
}

type IOGPUMetalResidencySetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalResidencySetClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalResidencySetClass) Alloc() IOGPUMetalResidencySet {
	rv := objc.SendIfResponds[IOGPUMetalResidencySet](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalResidencySet._commitAddedAllocationsCountRemovedAllocationsCount]
//   - [IOGPUMetalResidencySet._commitAddedHeapsCountRemovedHeapsCount]
//   - [IOGPUMetalResidencySet._commitAddedResourcesCountRemovedResourcesCount]
//   - [IOGPUMetalResidencySet.AddAllocation]
//   - [IOGPUMetalResidencySet.AddAllocationsCount]
//   - [IOGPUMetalResidencySet.AllAllocations]
//   - [IOGPUMetalResidencySet.AllCommittedAllocations]
//   - [IOGPUMetalResidencySet.AllocatedSize]
//   - [IOGPUMetalResidencySet.AllocationCount]
//   - [IOGPUMetalResidencySet.Commit]
//   - [IOGPUMetalResidencySet.ContainsAllocation]
//   - [IOGPUMetalResidencySet.CountForAllocation]
//   - [IOGPUMetalResidencySet.CurrentGeneration]
//   - [IOGPUMetalResidencySet.SetCurrentGeneration]
//   - [IOGPUMetalResidencySet.Device]
//   - [IOGPUMetalResidencySet.EndResidency]
//   - [IOGPUMetalResidencySet.ExpiredGeneration]
//   - [IOGPUMetalResidencySet.SetExpiredGeneration]
//   - [IOGPUMetalResidencySet.GenerationForAllocation]
//   - [IOGPUMetalResidencySet.RemoveAllAllocations]
//   - [IOGPUMetalResidencySet.RemoveAllocation]
//   - [IOGPUMetalResidencySet.RemoveAllocationsCount]
//   - [IOGPUMetalResidencySet.RequestResidency]
//   - [IOGPUMetalResidencySet.InitWithDeviceDescriptorArgsArgsSize]
//   - [IOGPUMetalResidencySet.DebugDescription]
//   - [IOGPUMetalResidencySet.Description]
//   - [IOGPUMetalResidencySet.Hash]
//   - [IOGPUMetalResidencySet.Label]
//   - [IOGPUMetalResidencySet.Superclass]
type IOGPUMetalResidencySet struct {
	objectivec.Object
}

// IOGPUMetalResidencySetFromID constructs a [IOGPUMetalResidencySet] from an objc.ID.
func IOGPUMetalResidencySetFromID(id objc.ID) IOGPUMetalResidencySet {
	return IOGPUMetalResidencySet{objectivec.Object{ID: id}}
}

// NOTE: IOGPUMetalResidencySet embeds objectivec.Object because the parent type is
// unavailable, but IIOGPUMetalResidencySet embeds IMTLObjectWithLabel, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [IOGPUMetalResidencySet] class.
//
// # Methods
//
//   - [IIOGPUMetalResidencySet._commitAddedAllocationsCountRemovedAllocationsCount]
//   - [IIOGPUMetalResidencySet._commitAddedHeapsCountRemovedHeapsCount]
//   - [IIOGPUMetalResidencySet._commitAddedResourcesCountRemovedResourcesCount]
//   - [IIOGPUMetalResidencySet.AddAllocation]
//   - [IIOGPUMetalResidencySet.AddAllocationsCount]
//   - [IIOGPUMetalResidencySet.AllAllocations]
//   - [IIOGPUMetalResidencySet.AllCommittedAllocations]
//   - [IIOGPUMetalResidencySet.AllocatedSize]
//   - [IIOGPUMetalResidencySet.AllocationCount]
//   - [IIOGPUMetalResidencySet.Commit]
//   - [IIOGPUMetalResidencySet.ContainsAllocation]
//   - [IIOGPUMetalResidencySet.CountForAllocation]
//   - [IIOGPUMetalResidencySet.CurrentGeneration]
//   - [IIOGPUMetalResidencySet.SetCurrentGeneration]
//   - [IIOGPUMetalResidencySet.Device]
//   - [IIOGPUMetalResidencySet.EndResidency]
//   - [IIOGPUMetalResidencySet.ExpiredGeneration]
//   - [IIOGPUMetalResidencySet.SetExpiredGeneration]
//   - [IIOGPUMetalResidencySet.GenerationForAllocation]
//   - [IIOGPUMetalResidencySet.RemoveAllAllocations]
//   - [IIOGPUMetalResidencySet.RemoveAllocation]
//   - [IIOGPUMetalResidencySet.RemoveAllocationsCount]
//   - [IIOGPUMetalResidencySet.RequestResidency]
//   - [IIOGPUMetalResidencySet.InitWithDeviceDescriptorArgsArgsSize]
//   - [IIOGPUMetalResidencySet.DebugDescription]
//   - [IIOGPUMetalResidencySet.Description]
//   - [IIOGPUMetalResidencySet.Hash]
//   - [IIOGPUMetalResidencySet.Label]
//   - [IIOGPUMetalResidencySet.Superclass]
type IIOGPUMetalResidencySet interface {
	IMTLObjectWithLabel

	// Topic: Methods

	_commitAddedAllocationsCountRemovedAllocationsCount(allocations []objectivec.IObject, count uint64, allocations2 []objectivec.IObject, count2 uint64) bool
	_commitAddedHeapsCountRemovedHeapsCount(heaps []objectivec.IObject, count uint64, heaps2 []objectivec.IObject, count2 uint64) bool
	_commitAddedResourcesCountRemovedResourcesCount(resources []objectivec.IObject, count uint64, resources2 []objectivec.IObject, count2 uint64) bool
	AddAllocation(allocation objectivec.IObject)
	AddAllocationsCount(allocations []objectivec.IObject, count uint64)
	AllAllocations() foundation.INSArray
	AllCommittedAllocations() foundation.INSArray
	AllocatedSize() uint64
	AllocationCount() uint64
	Commit()
	ContainsAllocation(allocation objectivec.IObject) bool
	CountForAllocation(allocation objectivec.IObject) uint64
	CurrentGeneration() uint64
	SetCurrentGeneration(value uint64)
	Device() unsafe.Pointer
	EndResidency()
	ExpiredGeneration() uint64
	SetExpiredGeneration(value uint64)
	GenerationForAllocation(allocation objectivec.IObject) uint64
	RemoveAllAllocations()
	RemoveAllocation(allocation objectivec.IObject)
	RemoveAllocationsCount(allocations []objectivec.IObject, count uint64)
	RequestResidency()
	InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUNewResourceArgs, size uint64) IOGPUMetalResidencySet
	DebugDescription() string
	Description() string
	Hash() uint64
	Label() string
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i IOGPUMetalResidencySet) Init() IOGPUMetalResidencySet {
	rv := objc.SendIfResponds[IOGPUMetalResidencySet](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalResidencySet) Autorelease() IOGPUMetalResidencySet {
	rv := objc.SendIfResponds[IOGPUMetalResidencySet](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalResidencySet creates a new IOGPUMetalResidencySet instance.
func NewIOGPUMetalResidencySet() IOGPUMetalResidencySet {
	class := getIOGPUMetalResidencySetClass()
	rv := objc.SendIfResponds[IOGPUMetalResidencySet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalResidencySetWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUNewResourceArgs, size uint64) IOGPUMetalResidencySet {
	instance := getIOGPUMetalResidencySetClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return IOGPUMetalResidencySetFromID(rv)
}

func (i IOGPUMetalResidencySet) _commitAddedAllocationsCountRemovedAllocationsCount(allocations []objectivec.IObject, count uint64, allocations2 []objectivec.IObject, count2 uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("_commitAddedAllocations:count:removedAllocations:count:"), objc.CArray(allocations), count, objectivec.IObjectSliceToNSArray(allocations2), count2)
	return rv
}

// CommitAddedAllocationsCountRemovedAllocationsCount is an exported wrapper for the private method _commitAddedAllocationsCountRemovedAllocationsCount.
func (i IOGPUMetalResidencySet) CommitAddedAllocationsCountRemovedAllocationsCount(allocations []objectivec.IObject, count uint64, allocations2 []objectivec.IObject, count2 uint64) (bool, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedAllocations:count:removedAllocations:count:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_commitAddedAllocations:count:removedAllocations:count:"}
		return false, err
	}
	return i._commitAddedAllocationsCountRemovedAllocationsCount(allocations, count, allocations2, count2), nil
}

// CanCommitAddedAllocationsCountRemovedAllocationsCount reports whether the receiver responds to the private selector _commitAddedAllocations:count:removedAllocations:count:.
func (i IOGPUMetalResidencySet) CanCommitAddedAllocationsCountRemovedAllocationsCount() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedAllocations:count:removedAllocations:count:"))
}
func (i IOGPUMetalResidencySet) _commitAddedHeapsCountRemovedHeapsCount(heaps []objectivec.IObject, count uint64, heaps2 []objectivec.IObject, count2 uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("_commitAddedHeaps:count:removedHeaps:count:"), objc.CArray(heaps), count, objectivec.IObjectSliceToNSArray(heaps2), count2)
	return rv
}

// CommitAddedHeapsCountRemovedHeapsCount is an exported wrapper for the private method _commitAddedHeapsCountRemovedHeapsCount.
func (i IOGPUMetalResidencySet) CommitAddedHeapsCountRemovedHeapsCount(heaps []objectivec.IObject, count uint64, heaps2 []objectivec.IObject, count2 uint64) (bool, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedHeaps:count:removedHeaps:count:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_commitAddedHeaps:count:removedHeaps:count:"}
		return false, err
	}
	return i._commitAddedHeapsCountRemovedHeapsCount(heaps, count, heaps2, count2), nil
}

// CanCommitAddedHeapsCountRemovedHeapsCount reports whether the receiver responds to the private selector _commitAddedHeaps:count:removedHeaps:count:.
func (i IOGPUMetalResidencySet) CanCommitAddedHeapsCountRemovedHeapsCount() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedHeaps:count:removedHeaps:count:"))
}
func (i IOGPUMetalResidencySet) _commitAddedResourcesCountRemovedResourcesCount(resources []objectivec.IObject, count uint64, resources2 []objectivec.IObject, count2 uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("_commitAddedResources:count:removedResources:count:"), objc.CArray(resources), count, objectivec.IObjectSliceToNSArray(resources2), count2)
	return rv
}

// CommitAddedResourcesCountRemovedResourcesCount is an exported wrapper for the private method _commitAddedResourcesCountRemovedResourcesCount.
func (i IOGPUMetalResidencySet) CommitAddedResourcesCountRemovedResourcesCount(resources []objectivec.IObject, count uint64, resources2 []objectivec.IObject, count2 uint64) (bool, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedResources:count:removedResources:count:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_commitAddedResources:count:removedResources:count:"}
		return false, err
	}
	return i._commitAddedResourcesCountRemovedResourcesCount(resources, count, resources2, count2), nil
}

// CanCommitAddedResourcesCountRemovedResourcesCount reports whether the receiver responds to the private selector _commitAddedResources:count:removedResources:count:.
func (i IOGPUMetalResidencySet) CanCommitAddedResourcesCountRemovedResourcesCount() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_commitAddedResources:count:removedResources:count:"))
}
func (i IOGPUMetalResidencySet) AddAllocation(allocation objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addAllocation:"), allocation)
}
func (i IOGPUMetalResidencySet) AddAllocationsCount(allocations []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addAllocations:count:"), objc.CArray(allocations), count)
}
func (i IOGPUMetalResidencySet) Commit() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("commit"))
}
func (i IOGPUMetalResidencySet) ContainsAllocation(allocation objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("containsAllocation:"), allocation)
	return rv
}
func (i IOGPUMetalResidencySet) CountForAllocation(allocation objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("countForAllocation:"), allocation)
	return rv
}
func (i IOGPUMetalResidencySet) EndResidency() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("endResidency"))
}
func (i IOGPUMetalResidencySet) GenerationForAllocation(allocation objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("generationForAllocation:"), allocation)
	return rv
}
func (i IOGPUMetalResidencySet) RemoveAllAllocations() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeAllAllocations"))
}
func (i IOGPUMetalResidencySet) RemoveAllocation(allocation objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeAllocation:"), allocation)
}
func (i IOGPUMetalResidencySet) RemoveAllocationsCount(allocations []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeAllocations:count:"), objc.CArray(allocations), count)
}
func (i IOGPUMetalResidencySet) RequestResidency() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("requestResidency"))
}
func (i IOGPUMetalResidencySet) InitWithDeviceDescriptorArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, args *IOGPUNewResourceArgs, size uint64) IOGPUMetalResidencySet {
	rv := objc.SendIfResponds[IOGPUMetalResidencySet](i.ID, objc.Sel("initWithDevice:descriptor:args:argsSize:"), device, descriptor, unsafe.Pointer(args), size)
	return rv
}

func (i IOGPUMetalResidencySet) AllAllocations() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allAllocations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (i IOGPUMetalResidencySet) AllCommittedAllocations() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allCommittedAllocations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (i IOGPUMetalResidencySet) AllocatedSize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("allocatedSize"))
	return rv
}
func (i IOGPUMetalResidencySet) AllocationCount() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("allocationCount"))
	return rv
}
func (i IOGPUMetalResidencySet) CurrentGeneration() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("currentGeneration"))
	return rv
}
func (i IOGPUMetalResidencySet) SetCurrentGeneration(value uint64) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setCurrentGeneration:"), value)
}
func (i IOGPUMetalResidencySet) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalResidencySet) Description() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalResidencySet) Device() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("device"))
	return rv
}
func (i IOGPUMetalResidencySet) ExpiredGeneration() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("expiredGeneration"))
	return rv
}
func (i IOGPUMetalResidencySet) SetExpiredGeneration(value uint64) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setExpiredGeneration:"), value)
}
func (i IOGPUMetalResidencySet) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i IOGPUMetalResidencySet) Label() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalResidencySet) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
