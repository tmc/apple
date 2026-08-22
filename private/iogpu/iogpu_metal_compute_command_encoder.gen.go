// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalComputeCommandEncoder] class.
var (
	_IOGPUMetalComputeCommandEncoderClass     IOGPUMetalComputeCommandEncoderClass
	_IOGPUMetalComputeCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalComputeCommandEncoderClass() IOGPUMetalComputeCommandEncoderClass {
	_IOGPUMetalComputeCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalComputeCommandEncoderClass = IOGPUMetalComputeCommandEncoderClass{class: objc.GetClass("IOGPUMetalComputeCommandEncoder")}
	})
	return _IOGPUMetalComputeCommandEncoderClass
}

// GetIOGPUMetalComputeCommandEncoderClass returns the class object for IOGPUMetalComputeCommandEncoder.
func GetIOGPUMetalComputeCommandEncoderClass() IOGPUMetalComputeCommandEncoderClass {
	return getIOGPUMetalComputeCommandEncoderClass()
}

type IOGPUMetalComputeCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalComputeCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalComputeCommandEncoderClass) Alloc() IOGPUMetalComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalComputeCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalComputeCommandEncoder.ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset]
//   - [IOGPUMetalComputeCommandEncoder.ExecuteCommandsInBufferWithRange]
//   - [IOGPUMetalComputeCommandEncoder.GetBufferContentsAtIndex]
//   - [IOGPUMetalComputeCommandEncoder.GetComputePipelineState]
//   - [IOGPUMetalComputeCommandEncoder.GetType]
//   - [IOGPUMetalComputeCommandEncoder.MemoryBarrierWithResourcesCount]
//   - [IOGPUMetalComputeCommandEncoder.MemoryBarrierWithScope]
//   - [IOGPUMetalComputeCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IOGPUMetalComputeCommandEncoder.SetAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalComputeCommandEncoder.SetFunctionTableAtIndex]
//   - [IOGPUMetalComputeCommandEncoder.SetFunctionTablesWithRange]
//   - [IOGPUMetalComputeCommandEncoder.SetIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalComputeCommandEncoder.SetIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalComputeCommandEncoder.SetVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalComputeCommandEncoder.SetVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalComputeCommandEncoder.UpdateFence]
//   - [IOGPUMetalComputeCommandEncoder.UseHeap]
//   - [IOGPUMetalComputeCommandEncoder.UseHeapsCount]
//   - [IOGPUMetalComputeCommandEncoder.UseResourceUsage]
//   - [IOGPUMetalComputeCommandEncoder.UseResourcesCountUsage]
//   - [IOGPUMetalComputeCommandEncoder.WaitForFence]
type IOGPUMetalComputeCommandEncoder struct {
	IOGPUMetalCommandEncoder
}

// IOGPUMetalComputeCommandEncoderFromID constructs a [IOGPUMetalComputeCommandEncoder] from an objc.ID.
func IOGPUMetalComputeCommandEncoderFromID(id objc.ID) IOGPUMetalComputeCommandEncoder {
	return IOGPUMetalComputeCommandEncoder{IOGPUMetalCommandEncoder: IOGPUMetalCommandEncoderFromID(id)}
}

// Ensure IOGPUMetalComputeCommandEncoder implements IIOGPUMetalComputeCommandEncoder.
var _ IIOGPUMetalComputeCommandEncoder = IOGPUMetalComputeCommandEncoder{}

// An interface definition for the [IOGPUMetalComputeCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalComputeCommandEncoder.ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset]
//   - [IIOGPUMetalComputeCommandEncoder.ExecuteCommandsInBufferWithRange]
//   - [IIOGPUMetalComputeCommandEncoder.GetBufferContentsAtIndex]
//   - [IIOGPUMetalComputeCommandEncoder.GetComputePipelineState]
//   - [IIOGPUMetalComputeCommandEncoder.GetType]
//   - [IIOGPUMetalComputeCommandEncoder.MemoryBarrierWithResourcesCount]
//   - [IIOGPUMetalComputeCommandEncoder.MemoryBarrierWithScope]
//   - [IIOGPUMetalComputeCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IIOGPUMetalComputeCommandEncoder.SetAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalComputeCommandEncoder.SetFunctionTableAtIndex]
//   - [IIOGPUMetalComputeCommandEncoder.SetFunctionTablesWithRange]
//   - [IIOGPUMetalComputeCommandEncoder.SetIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalComputeCommandEncoder.SetIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalComputeCommandEncoder.SetVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalComputeCommandEncoder.SetVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalComputeCommandEncoder.UpdateFence]
//   - [IIOGPUMetalComputeCommandEncoder.UseHeap]
//   - [IIOGPUMetalComputeCommandEncoder.UseHeapsCount]
//   - [IIOGPUMetalComputeCommandEncoder.UseResourceUsage]
//   - [IIOGPUMetalComputeCommandEncoder.UseResourcesCountUsage]
//   - [IIOGPUMetalComputeCommandEncoder.WaitForFence]
type IIOGPUMetalComputeCommandEncoder interface {
	IIOGPUMetalCommandEncoder

	// Topic: Methods

	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64)
	ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange)
	GetBufferContentsAtIndex(index uint64) unsafe.Pointer
	GetComputePipelineState() objectivec.IObject
	GetType() uint64
	MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64)
	MemoryBarrierWithScope(scope uint64)
	SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool)
	SetAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetFunctionTableAtIndex(table objectivec.IObject, index uint64)
	SetFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	UpdateFence(fence objectivec.IObject)
	UseHeap(heap objectivec.IObject)
	UseHeapsCount(heaps []objectivec.IObject, count uint64)
	UseResourceUsage(resource objectivec.IObject, usage uint64)
	UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64)
	WaitForFence(fence objectivec.IObject)
}

// Init initializes the instance.
func (i IOGPUMetalComputeCommandEncoder) Init() IOGPUMetalComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalComputeCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalComputeCommandEncoder) Autorelease() IOGPUMetalComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalComputeCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalComputeCommandEncoder creates a new IOGPUMetalComputeCommandEncoder instance.
func NewIOGPUMetalComputeCommandEncoder() IOGPUMetalComputeCommandEncoder {
	class := getIOGPUMetalComputeCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalComputeCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalComputeCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalComputeCommandEncoder {
	instance := getIOGPUMetalComputeCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalComputeCommandEncoderFromID(rv)
}

func (i IOGPUMetalComputeCommandEncoder) ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:"), buffer, buffer2, offset)
}
func (i IOGPUMetalComputeCommandEncoder) ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("executeCommandsInBuffer:withRange:"), buffer, range_)
}
func (i IOGPUMetalComputeCommandEncoder) GetBufferContentsAtIndex(index uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getBufferContentsAtIndex:"), index)
	return rv
}
func (i IOGPUMetalComputeCommandEncoder) GetComputePipelineState() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getComputePipelineState"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalComputeCommandEncoder) GetType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetalComputeCommandEncoder) MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("memoryBarrierWithResources:count:"), objc.CArray(resources), count)
}
func (i IOGPUMetalComputeCommandEncoder) MemoryBarrierWithScope(scope uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("memoryBarrierWithScope:"), scope)
}
func (i IOGPUMetalComputeCommandEncoder) SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), buffer, index, barrier)
}
func (i IOGPUMetalComputeCommandEncoder) SetAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalComputeCommandEncoder) SetFunctionTableAtIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctionTable:atIndex:"), table, index)
}
func (i IOGPUMetalComputeCommandEncoder) SetFunctionTablesWithRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFunctionTables:withRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalComputeCommandEncoder) SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalComputeCommandEncoder) SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalComputeCommandEncoder) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalComputeCommandEncoder) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalComputeCommandEncoder) UpdateFence(fence objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateFence:"), fence)
}
func (i IOGPUMetalComputeCommandEncoder) UseHeap(heap objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeap:"), heap)
}
func (i IOGPUMetalComputeCommandEncoder) UseHeapsCount(heaps []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
}
func (i IOGPUMetalComputeCommandEncoder) UseResourceUsage(resource objectivec.IObject, usage uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResource:usage:"), resource, usage)
}
func (i IOGPUMetalComputeCommandEncoder) UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResources:count:usage:"), objc.CArray(resources), count, usage)
}
func (i IOGPUMetalComputeCommandEncoder) WaitForFence(fence objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("waitForFence:"), fence)
}
