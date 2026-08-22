// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIntersectionFunctionTable] class.
var (
	_IOGPUMetalIntersectionFunctionTableClass     IOGPUMetalIntersectionFunctionTableClass
	_IOGPUMetalIntersectionFunctionTableClassOnce sync.Once
)

func getIOGPUMetalIntersectionFunctionTableClass() IOGPUMetalIntersectionFunctionTableClass {
	_IOGPUMetalIntersectionFunctionTableClassOnce.Do(func() {
		_IOGPUMetalIntersectionFunctionTableClass = IOGPUMetalIntersectionFunctionTableClass{class: objc.GetClass("IOGPUMetalIntersectionFunctionTable")}
	})
	return _IOGPUMetalIntersectionFunctionTableClass
}

// GetIOGPUMetalIntersectionFunctionTableClass returns the class object for IOGPUMetalIntersectionFunctionTable.
func GetIOGPUMetalIntersectionFunctionTableClass() IOGPUMetalIntersectionFunctionTableClass {
	return getIOGPUMetalIntersectionFunctionTableClass()
}

type IOGPUMetalIntersectionFunctionTableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIntersectionFunctionTableClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIntersectionFunctionTableClass) Alloc() IOGPUMetalIntersectionFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalIntersectionFunctionTable](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIntersectionFunctionTable.BufferAddressAtIndex]
//   - [IOGPUMetalIntersectionFunctionTable.GlobalBuffer]
//   - [IOGPUMetalIntersectionFunctionTable.SetGlobalBuffer]
//   - [IOGPUMetalIntersectionFunctionTable.GlobalBufferOffset]
//   - [IOGPUMetalIntersectionFunctionTable.SetGlobalBufferOffset]
//   - [IOGPUMetalIntersectionFunctionTable.GpuHandle]
//   - [IOGPUMetalIntersectionFunctionTable.GpuResourceID]
//   - [IOGPUMetalIntersectionFunctionTable.ResourceIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetBufferOffsetAtIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetBuffersOffsetsWithRange]
//   - [IOGPUMetalIntersectionFunctionTable.SetFunctionAtIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetFunctionsWithRange]
//   - [IOGPUMetalIntersectionFunctionTable.SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetOpaqueCurveIntersectionFunctionWithSignatureWithRange]
//   - [IOGPUMetalIntersectionFunctionTable.SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange]
//   - [IOGPUMetalIntersectionFunctionTable.SetVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalIntersectionFunctionTable.SetVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalIntersectionFunctionTable.UniqueIdentifier]
//   - [IOGPUMetalIntersectionFunctionTable.VisibleFunctionTable]
//   - [IOGPUMetalIntersectionFunctionTable.InitWithVisibleFunctionTable]
type IOGPUMetalIntersectionFunctionTable struct {
	IOGPUMetalResource
}

// IOGPUMetalIntersectionFunctionTableFromID constructs a [IOGPUMetalIntersectionFunctionTable] from an objc.ID.
func IOGPUMetalIntersectionFunctionTableFromID(id objc.ID) IOGPUMetalIntersectionFunctionTable {
	return IOGPUMetalIntersectionFunctionTable{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalIntersectionFunctionTable implements IIOGPUMetalIntersectionFunctionTable.
var _ IIOGPUMetalIntersectionFunctionTable = IOGPUMetalIntersectionFunctionTable{}

// An interface definition for the [IOGPUMetalIntersectionFunctionTable] class.
//
// # Methods
//
//   - [IIOGPUMetalIntersectionFunctionTable.BufferAddressAtIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.GlobalBuffer]
//   - [IIOGPUMetalIntersectionFunctionTable.SetGlobalBuffer]
//   - [IIOGPUMetalIntersectionFunctionTable.GlobalBufferOffset]
//   - [IIOGPUMetalIntersectionFunctionTable.SetGlobalBufferOffset]
//   - [IIOGPUMetalIntersectionFunctionTable.GpuHandle]
//   - [IIOGPUMetalIntersectionFunctionTable.GpuResourceID]
//   - [IIOGPUMetalIntersectionFunctionTable.ResourceIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetBufferOffsetAtIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetBuffersOffsetsWithRange]
//   - [IIOGPUMetalIntersectionFunctionTable.SetFunctionAtIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetFunctionsWithRange]
//   - [IIOGPUMetalIntersectionFunctionTable.SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetOpaqueCurveIntersectionFunctionWithSignatureWithRange]
//   - [IIOGPUMetalIntersectionFunctionTable.SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange]
//   - [IIOGPUMetalIntersectionFunctionTable.SetVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalIntersectionFunctionTable.SetVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalIntersectionFunctionTable.UniqueIdentifier]
//   - [IIOGPUMetalIntersectionFunctionTable.VisibleFunctionTable]
//   - [IIOGPUMetalIntersectionFunctionTable.InitWithVisibleFunctionTable]
type IIOGPUMetalIntersectionFunctionTable interface {
	IIOGPUMetalResource

	// Topic: Methods

	BufferAddressAtIndex(index uint64) uint64
	GlobalBuffer() unsafe.Pointer
	SetGlobalBuffer(value unsafe.Pointer)
	GlobalBufferOffset() uint64
	SetGlobalBufferOffset(value uint64)
	GpuHandle() uint64
	GpuResourceID() metal.MTLResourceID
	ResourceIndex() uint64
	SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64)
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange)
	SetFunctionAtIndex(function objectivec.IObject, index uint64)
	SetFunctionsWithRange(functions []objectivec.IObject, range_ foundation.NSRange)
	SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature uint64, index uint64)
	SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature uint64, range_ foundation.NSRange)
	SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature uint64, index uint64)
	SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature uint64, range_ foundation.NSRange)
	SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	UniqueIdentifier() uint64
	VisibleFunctionTable() IIOGPUMetalResource
	InitWithVisibleFunctionTable(table objectivec.IObject) IOGPUMetalIntersectionFunctionTable
}

// Init initializes the instance.
func (i IOGPUMetalIntersectionFunctionTable) Init() IOGPUMetalIntersectionFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalIntersectionFunctionTable](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIntersectionFunctionTable) Autorelease() IOGPUMetalIntersectionFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalIntersectionFunctionTable](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIntersectionFunctionTable creates a new IOGPUMetalIntersectionFunctionTable instance.
func NewIOGPUMetalIntersectionFunctionTable() IOGPUMetalIntersectionFunctionTable {
	class := getIOGPUMetalIntersectionFunctionTableClass()
	rv := objc.SendIfResponds[IOGPUMetalIntersectionFunctionTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIntersectionFunctionTableMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func NewGPUMetalIntersectionFunctionTableStandinWithDevice(device objectivec.IObject) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func NewGPUMetalIntersectionFunctionTableWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func NewGPUMetalIntersectionFunctionTableWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func NewGPUMetalIntersectionFunctionTableWithResource(resource objectivec.IObject) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func NewGPUMetalIntersectionFunctionTableWithVisibleFunctionTable(table objectivec.IObject) IOGPUMetalIntersectionFunctionTable {
	instance := getIOGPUMetalIntersectionFunctionTableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVisibleFunctionTable:"), table)
	return IOGPUMetalIntersectionFunctionTableFromID(rv)
}

func (i IOGPUMetalIntersectionFunctionTable) BufferAddressAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferAddressAtIndex:"), index)
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) SetBufferOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}
func (i IOGPUMetalIntersectionFunctionTable) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), unsafe.Pointer(offsets), range_)
}
func (i IOGPUMetalIntersectionFunctionTable) SetFunctionAtIndex(function objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunction:atIndex:"), function, index)
}
func (i IOGPUMetalIntersectionFunctionTable) SetFunctionsWithRange(functions []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctions:withRange:"), objectivec.IObjectSliceToNSArray(functions), range_)
}
func (i IOGPUMetalIntersectionFunctionTable) SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setOpaqueCurveIntersectionFunctionWithSignature:atIndex:"), signature, index)
}
func (i IOGPUMetalIntersectionFunctionTable) SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setOpaqueCurveIntersectionFunctionWithSignature:withRange:"), signature, range_)
}
func (i IOGPUMetalIntersectionFunctionTable) SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:atIndex:"), signature, index)
}
func (i IOGPUMetalIntersectionFunctionTable) SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature uint64, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:withRange:"), signature, range_)
}
func (i IOGPUMetalIntersectionFunctionTable) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalIntersectionFunctionTable) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalIntersectionFunctionTable) InitWithVisibleFunctionTable(table objectivec.IObject) IOGPUMetalIntersectionFunctionTable {
	rv := objc.SendIfResponds[IOGPUMetalIntersectionFunctionTable](i.ID, objc.Sel("initWithVisibleFunctionTable:"), table)
	return rv
}

func (i IOGPUMetalIntersectionFunctionTable) GlobalBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("globalBuffer"))
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) SetGlobalBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setGlobalBuffer:"), value)
}
func (i IOGPUMetalIntersectionFunctionTable) GlobalBufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalBufferOffset"))
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) SetGlobalBufferOffset(value uint64) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setGlobalBufferOffset:"), value)
}
func (i IOGPUMetalIntersectionFunctionTable) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuHandle"))
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalIntersectionFunctionTable) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("resourceIndex"))
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
func (i IOGPUMetalIntersectionFunctionTable) VisibleFunctionTable() IIOGPUMetalResource {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("visibleFunctionTable"))
	return IOGPUMetalResourceFromID(objc.ID(rv))
}
