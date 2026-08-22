// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalVisibleFunctionTable] class.
var (
	_IOGPUMetalVisibleFunctionTableClass     IOGPUMetalVisibleFunctionTableClass
	_IOGPUMetalVisibleFunctionTableClassOnce sync.Once
)

func getIOGPUMetalVisibleFunctionTableClass() IOGPUMetalVisibleFunctionTableClass {
	_IOGPUMetalVisibleFunctionTableClassOnce.Do(func() {
		_IOGPUMetalVisibleFunctionTableClass = IOGPUMetalVisibleFunctionTableClass{class: objc.GetClass("IOGPUMetalVisibleFunctionTable")}
	})
	return _IOGPUMetalVisibleFunctionTableClass
}

// GetIOGPUMetalVisibleFunctionTableClass returns the class object for IOGPUMetalVisibleFunctionTable.
func GetIOGPUMetalVisibleFunctionTableClass() IOGPUMetalVisibleFunctionTableClass {
	return getIOGPUMetalVisibleFunctionTableClass()
}

type IOGPUMetalVisibleFunctionTableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalVisibleFunctionTableClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalVisibleFunctionTableClass) Alloc() IOGPUMetalVisibleFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalVisibleFunctionTable](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalVisibleFunctionTable.BufferAddressAtIndex]
//   - [IOGPUMetalVisibleFunctionTable.GlobalBuffer]
//   - [IOGPUMetalVisibleFunctionTable.SetGlobalBuffer]
//   - [IOGPUMetalVisibleFunctionTable.GlobalBufferOffset]
//   - [IOGPUMetalVisibleFunctionTable.SetGlobalBufferOffset]
//   - [IOGPUMetalVisibleFunctionTable.GpuHandle]
//   - [IOGPUMetalVisibleFunctionTable.GpuResourceID]
//   - [IOGPUMetalVisibleFunctionTable.SetBufferOffsetAtIndex]
//   - [IOGPUMetalVisibleFunctionTable.SetBuffersOffsetsWithRange]
//   - [IOGPUMetalVisibleFunctionTable.SetFunctionAtIndex]
//   - [IOGPUMetalVisibleFunctionTable.SetFunctionsWithRange]
//   - [IOGPUMetalVisibleFunctionTable.SetValueAtIndex]
//   - [IOGPUMetalVisibleFunctionTable.SetValueWithRange]
//   - [IOGPUMetalVisibleFunctionTable.SetVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalVisibleFunctionTable.SetVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalVisibleFunctionTable.ResourceIndex]
//   - [IOGPUMetalVisibleFunctionTable.UniqueIdentifier]
type IOGPUMetalVisibleFunctionTable struct {
	IOGPUMetalBuffer
}

// IOGPUMetalVisibleFunctionTableFromID constructs a [IOGPUMetalVisibleFunctionTable] from an objc.ID.
func IOGPUMetalVisibleFunctionTableFromID(id objc.ID) IOGPUMetalVisibleFunctionTable {
	return IOGPUMetalVisibleFunctionTable{IOGPUMetalBuffer: IOGPUMetalBufferFromID(id)}
}

// Ensure IOGPUMetalVisibleFunctionTable implements IIOGPUMetalVisibleFunctionTable.
var _ IIOGPUMetalVisibleFunctionTable = IOGPUMetalVisibleFunctionTable{}

// An interface definition for the [IOGPUMetalVisibleFunctionTable] class.
//
// # Methods
//
//   - [IIOGPUMetalVisibleFunctionTable.BufferAddressAtIndex]
//   - [IIOGPUMetalVisibleFunctionTable.GlobalBuffer]
//   - [IIOGPUMetalVisibleFunctionTable.SetGlobalBuffer]
//   - [IIOGPUMetalVisibleFunctionTable.GlobalBufferOffset]
//   - [IIOGPUMetalVisibleFunctionTable.SetGlobalBufferOffset]
//   - [IIOGPUMetalVisibleFunctionTable.GpuHandle]
//   - [IIOGPUMetalVisibleFunctionTable.GpuResourceID]
//   - [IIOGPUMetalVisibleFunctionTable.SetBufferOffsetAtIndex]
//   - [IIOGPUMetalVisibleFunctionTable.SetBuffersOffsetsWithRange]
//   - [IIOGPUMetalVisibleFunctionTable.SetFunctionAtIndex]
//   - [IIOGPUMetalVisibleFunctionTable.SetFunctionsWithRange]
//   - [IIOGPUMetalVisibleFunctionTable.SetValueAtIndex]
//   - [IIOGPUMetalVisibleFunctionTable.SetValueWithRange]
//   - [IIOGPUMetalVisibleFunctionTable.SetVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalVisibleFunctionTable.SetVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalVisibleFunctionTable.ResourceIndex]
//   - [IIOGPUMetalVisibleFunctionTable.UniqueIdentifier]
type IIOGPUMetalVisibleFunctionTable interface {
	IIOGPUMetalBuffer

	// Topic: Methods

	BufferAddressAtIndex(index uint64) uint64
	GlobalBuffer() unsafe.Pointer
	SetGlobalBuffer(value unsafe.Pointer)
	GlobalBufferOffset() uint64
	SetGlobalBufferOffset(value uint64)
	GpuHandle() uint64
	GpuResourceID() metal.MTLResourceID
	SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64)
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange)
	SetFunctionAtIndex(function objectivec.IObject, index uint64)
	SetFunctionsWithRange(functions []objectivec.IObject, range_ foundation.NSRange)
	SetValueAtIndex(value uint64, index uint64)
	SetValueWithRange(value uint64, range_ foundation.NSRange)
	SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	ResourceIndex() uint64
	UniqueIdentifier() uint64
}

// Init initializes the instance.
func (i IOGPUMetalVisibleFunctionTable) Init() IOGPUMetalVisibleFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalVisibleFunctionTable](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalVisibleFunctionTable) Autorelease() IOGPUMetalVisibleFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalVisibleFunctionTable](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalVisibleFunctionTable creates a new IOGPUMetalVisibleFunctionTable instance.
func NewIOGPUMetalVisibleFunctionTable() IOGPUMetalVisibleFunctionTable {
	class := getIOGPUMetalVisibleFunctionTableClass()
	rv := objc.SendIfResponds[IOGPUMetalVisibleFunctionTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalVisibleFunctionTableMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableStandinWithDevice(device objectivec.IObject) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize(device objectivec.IObject, ranges *IOGPUAddressRange, count uint64, length uint64, options uint64, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:addressRanges:addressRangeCount:length:options:gpuAddress:args:argsSize:"), device, unsafe.Pointer(ranges), count, length, options, address, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceIosurfaceArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:iosurface:args:argsSize:"), device, iosurface, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceIosurfaceGpuAddressArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:iosurface:gpuAddress:args:argsSize:"), device, iosurface, address, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceRemoteStorageBufferArgsArgsSize(device objectivec.IObject, buffer objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageBuffer:args:argsSize:"), device, buffer, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithHeapResourceOffsetLength(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHeap:resource:offset:length:"), heap, resource, offset, length)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithHeapResourceOffsetLengthGpuTag(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, tag uint64) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHeap:resource:offset:length:gpuTag:"), heap, resource, offset, length, tag)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32, tag uint64) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:gpuTag:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size, tag)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func NewGPUMetalVisibleFunctionTableWithResource(resource objectivec.IObject) IOGPUMetalVisibleFunctionTable {
	instance := getIOGPUMetalVisibleFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalVisibleFunctionTableFromID(rv)
}

func (i IOGPUMetalVisibleFunctionTable) BufferAddressAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferAddressAtIndex:"), index)
	return rv
}
func (i IOGPUMetalVisibleFunctionTable) SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}
func (i IOGPUMetalVisibleFunctionTable) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), unsafe.Pointer(offsets), range_)
}
func (i IOGPUMetalVisibleFunctionTable) SetFunctionAtIndex(function objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunction:atIndex:"), function, index)
}
func (i IOGPUMetalVisibleFunctionTable) SetFunctionsWithRange(functions []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctions:withRange:"), objectivec.IObjectSliceToNSArray(functions), range_)
}
func (i IOGPUMetalVisibleFunctionTable) SetValueAtIndex(value uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setValue:atIndex:"), value, index)
}
func (i IOGPUMetalVisibleFunctionTable) SetValueWithRange(value uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setValue:withRange:"), value, range_)
}
func (i IOGPUMetalVisibleFunctionTable) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalVisibleFunctionTable) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}

func (i IOGPUMetalVisibleFunctionTable) GlobalBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("globalBuffer"))
	return rv
}
func (i IOGPUMetalVisibleFunctionTable) SetGlobalBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setGlobalBuffer:"), value)
}
func (i IOGPUMetalVisibleFunctionTable) GlobalBufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalBufferOffset"))
	return rv
}
func (i IOGPUMetalVisibleFunctionTable) SetGlobalBufferOffset(value uint64) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setGlobalBufferOffset:"), value)
}
func (i IOGPUMetalVisibleFunctionTable) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuHandle"))
	return rv
}
func (i IOGPUMetalVisibleFunctionTable) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalVisibleFunctionTable) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("resourceIndex"))
	return rv
}
func (i IOGPUMetalVisibleFunctionTable) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
