// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalHeap] class.
var (
	_IOGPUMetalHeapClass     IOGPUMetalHeapClass
	_IOGPUMetalHeapClassOnce sync.Once
)

func getIOGPUMetalHeapClass() IOGPUMetalHeapClass {
	_IOGPUMetalHeapClassOnce.Do(func() {
		_IOGPUMetalHeapClass = IOGPUMetalHeapClass{class: objc.GetClass("IOGPUMetalHeap")}
	})
	return _IOGPUMetalHeapClass
}

// GetIOGPUMetalHeapClass returns the class object for IOGPUMetalHeap.
func GetIOGPUMetalHeapClass() IOGPUMetalHeapClass {
	return getIOGPUMetalHeapClass()
}

type IOGPUMetalHeapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalHeapClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalHeapClass) Alloc() IOGPUMetalHeap {
	rv := objc.SendIfResponds[IOGPUMetalHeap](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalHeap.DeallocHeapSubResource]
//   - [IOGPUMetalHeap.DetachBacking]
//   - [IOGPUMetalHeap.GlobalTraceObjectID]
//   - [IOGPUMetalHeap.GpuAddress]
//   - [IOGPUMetalHeap.MaxCompatiblePlacementSparsePageSize]
//   - [IOGPUMetalHeap.NewAccelerationStructureWithSizeOffsetResourceIndex]
//   - [IOGPUMetalHeap.NewAccelerationStructureWithSizeResourceIndex]
//   - [IOGPUMetalHeap.NewSubResourceAtOffsetWithLengthAlignmentOptions]
//   - [IOGPUMetalHeap.NewSubResourceWithLengthAlignmentOptionsOffset]
//   - [IOGPUMetalHeap.ProtectionOptions]
//   - [IOGPUMetalHeap.ReplaceBackingWithRangesReadOnly]
//   - [IOGPUMetalHeap.UnpinMemoryAtOffsetWithLength]
//   - [IOGPUMetalHeap.InitWithDeviceSizeOptionsArgsArgsSize]
//   - [IOGPUMetalHeap.InitWithDeviceSizeOptionsArgsArgsSizeDesc]
//   - [IOGPUMetalHeap.UnfilteredResourceOptions]
type IOGPUMetalHeap struct {
	metal.MTLHeapObject
}

// IOGPUMetalHeapFromID constructs a [IOGPUMetalHeap] from an objc.ID.
func IOGPUMetalHeapFromID(id objc.ID) IOGPUMetalHeap {
	return IOGPUMetalHeap{MTLHeapObject: metal.MTLHeapObjectFromID(id)}
}

// Ensure IOGPUMetalHeap implements IIOGPUMetalHeap.
var _ IIOGPUMetalHeap = IOGPUMetalHeap{}

// An interface definition for the [IOGPUMetalHeap] class.
//
// # Methods
//
//   - [IIOGPUMetalHeap.DeallocHeapSubResource]
//   - [IIOGPUMetalHeap.DetachBacking]
//   - [IIOGPUMetalHeap.GlobalTraceObjectID]
//   - [IIOGPUMetalHeap.GpuAddress]
//   - [IIOGPUMetalHeap.MaxCompatiblePlacementSparsePageSize]
//   - [IIOGPUMetalHeap.NewAccelerationStructureWithSizeOffsetResourceIndex]
//   - [IIOGPUMetalHeap.NewAccelerationStructureWithSizeResourceIndex]
//   - [IIOGPUMetalHeap.NewSubResourceAtOffsetWithLengthAlignmentOptions]
//   - [IIOGPUMetalHeap.NewSubResourceWithLengthAlignmentOptionsOffset]
//   - [IIOGPUMetalHeap.ProtectionOptions]
//   - [IIOGPUMetalHeap.ReplaceBackingWithRangesReadOnly]
//   - [IIOGPUMetalHeap.UnpinMemoryAtOffsetWithLength]
//   - [IIOGPUMetalHeap.InitWithDeviceSizeOptionsArgsArgsSize]
//   - [IIOGPUMetalHeap.InitWithDeviceSizeOptionsArgsArgsSizeDesc]
//   - [IIOGPUMetalHeap.UnfilteredResourceOptions]
type IIOGPUMetalHeap interface {
	metal.MTLHeap

	// Topic: Methods

	DeallocHeapSubResource()
	DetachBacking() bool
	GlobalTraceObjectID() uint64
	GpuAddress() uint64
	MaxCompatiblePlacementSparsePageSize() int64
	NewAccelerationStructureWithSizeOffsetResourceIndex(size uint64, offset uint64, index uint64) objectivec.IObject
	NewAccelerationStructureWithSizeResourceIndex(size uint64, index uint64) objectivec.IObject
	NewSubResourceAtOffsetWithLengthAlignmentOptions(offset uint64, length uint64, alignment uint64, options uint64) objectivec.IObject
	NewSubResourceWithLengthAlignmentOptionsOffset(length uint64, alignment uint64, options uint64, offset *uint64) objectivec.IObject
	ProtectionOptions() uint64
	ReplaceBackingWithRangesReadOnly(ranges objectivec.IObject, only bool) bool
	UnpinMemoryAtOffsetWithLength(offset uint64, length uint64)
	InitWithDeviceSizeOptionsArgsArgsSize(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalHeap
	InitWithDeviceSizeOptionsArgsArgsSizeDesc(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32, desc objectivec.IObject) IOGPUMetalHeap
	UnfilteredResourceOptions() uint64
}

// Init initializes the instance.
func (i IOGPUMetalHeap) Init() IOGPUMetalHeap {
	rv := objc.SendIfResponds[IOGPUMetalHeap](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalHeap) Autorelease() IOGPUMetalHeap {
	rv := objc.SendIfResponds[IOGPUMetalHeap](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalHeap creates a new IOGPUMetalHeap instance.
func NewIOGPUMetalHeap() IOGPUMetalHeap {
	class := getIOGPUMetalHeapClass()
	rv := objc.SendIfResponds[IOGPUMetalHeap](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalHeapWithDeviceSizeOptionsArgsArgsSize(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalHeap {
	instance := getIOGPUMetalHeapClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:size:options:args:argsSize:"), device, size, options, unsafe.Pointer(args), size2)
	return IOGPUMetalHeapFromID(rv)
}

func NewGPUMetalHeapWithDeviceSizeOptionsArgsArgsSizeDesc(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32, desc objectivec.IObject) IOGPUMetalHeap {
	instance := getIOGPUMetalHeapClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:size:options:args:argsSize:desc:"), device, size, options, unsafe.Pointer(args), size2, desc)
	return IOGPUMetalHeapFromID(rv)
}

func (i IOGPUMetalHeap) DeallocHeapSubResource() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("deallocHeapSubResource"))
}
func (i IOGPUMetalHeap) DetachBacking() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("detachBacking"))
	return rv
}
func (i IOGPUMetalHeap) GpuAddress() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuAddress"))
	return rv
}
func (i IOGPUMetalHeap) NewAccelerationStructureWithSizeOffsetResourceIndex(size uint64, offset uint64, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newAccelerationStructureWithSize:offset:resourceIndex:"), size, offset, index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalHeap) NewAccelerationStructureWithSizeResourceIndex(size uint64, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newAccelerationStructureWithSize:resourceIndex:"), size, index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalHeap) NewSubResourceAtOffsetWithLengthAlignmentOptions(offset uint64, length uint64, alignment uint64, options uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newSubResourceAtOffset:withLength:alignment:options:"), offset, length, alignment, options)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalHeap) NewSubResourceWithLengthAlignmentOptionsOffset(length uint64, alignment uint64, options uint64, offset *uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newSubResourceWithLength:alignment:options:offset:"), length, alignment, options, unsafe.Pointer(offset))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalHeap) ReplaceBackingWithRangesReadOnly(ranges objectivec.IObject, only bool) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("replaceBackingWithRanges:readOnly:"), ranges, only)
	return rv
}
func (i IOGPUMetalHeap) UnpinMemoryAtOffsetWithLength(offset uint64, length uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("unpinMemoryAtOffset:withLength:"), offset, length)
}
func (i IOGPUMetalHeap) InitWithDeviceSizeOptionsArgsArgsSize(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalHeap {
	rv := objc.SendIfResponds[IOGPUMetalHeap](i.ID, objc.Sel("initWithDevice:size:options:args:argsSize:"), device, size, options, unsafe.Pointer(args), size2)
	return rv
}
func (i IOGPUMetalHeap) InitWithDeviceSizeOptionsArgsArgsSizeDesc(device objectivec.IObject, size uint64, options uint64, args *IOGPUNewResourceArgs, size2 uint32, desc objectivec.IObject) IOGPUMetalHeap {
	rv := objc.SendIfResponds[IOGPUMetalHeap](i.ID, objc.Sel("initWithDevice:size:options:args:argsSize:desc:"), device, size, options, unsafe.Pointer(args), size2, desc)
	return rv
}

func (i IOGPUMetalHeap) GlobalTraceObjectID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalTraceObjectID"))
	return rv
}
func (i IOGPUMetalHeap) MaxCompatiblePlacementSparsePageSize() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("maxCompatiblePlacementSparsePageSize"))
	return rv
}
func (i IOGPUMetalHeap) ProtectionOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("protectionOptions"))
	return rv
}
func (i IOGPUMetalHeap) UnfilteredResourceOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("unfilteredResourceOptions"))
	return rv
}
