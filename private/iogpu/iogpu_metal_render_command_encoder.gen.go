// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalRenderCommandEncoder] class.
var (
	_IOGPUMetalRenderCommandEncoderClass     IOGPUMetalRenderCommandEncoderClass
	_IOGPUMetalRenderCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalRenderCommandEncoderClass() IOGPUMetalRenderCommandEncoderClass {
	_IOGPUMetalRenderCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalRenderCommandEncoderClass = IOGPUMetalRenderCommandEncoderClass{class: objc.GetClass("IOGPUMetalRenderCommandEncoder")}
	})
	return _IOGPUMetalRenderCommandEncoderClass
}

// GetIOGPUMetalRenderCommandEncoderClass returns the class object for IOGPUMetalRenderCommandEncoder.
func GetIOGPUMetalRenderCommandEncoderClass() IOGPUMetalRenderCommandEncoderClass {
	return getIOGPUMetalRenderCommandEncoderClass()
}

type IOGPUMetalRenderCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalRenderCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalRenderCommandEncoderClass) Alloc() IOGPUMetalRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalRenderCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalRenderCommandEncoder.ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset]
//   - [IOGPUMetalRenderCommandEncoder.ExecuteCommandsInBufferWithRange]
//   - [IOGPUMetalRenderCommandEncoder.GetFragmentBufferContentsAtIndex]
//   - [IOGPUMetalRenderCommandEncoder.GetRenderPipelineState]
//   - [IOGPUMetalRenderCommandEncoder.GetType]
//   - [IOGPUMetalRenderCommandEncoder.GetVertexBufferContentsAtIndex]
//   - [IOGPUMetalRenderCommandEncoder.IsMemorylessRender]
//   - [IOGPUMetalRenderCommandEncoder.MemoryBarrierWithResourcesCountAfterStagesBeforeStages]
//   - [IOGPUMetalRenderCommandEncoder.MemoryBarrierWithScopeAfterStagesBeforeStages]
//   - [IOGPUMetalRenderCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IOGPUMetalRenderCommandEncoder.SetColorStoreActionAtIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetColorStoreActionOptionsAtIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetDepthCleared]
//   - [IOGPUMetalRenderCommandEncoder.SetDepthClipModeSPI]
//   - [IOGPUMetalRenderCommandEncoder.SetDepthStoreAction]
//   - [IOGPUMetalRenderCommandEncoder.SetDepthStoreActionOptions]
//   - [IOGPUMetalRenderCommandEncoder.SetFragmentAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetFragmentIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetFragmentIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetFragmentVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetFragmentVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetMeshAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetMeshIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetMeshIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetObjectAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetObjectIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetObjectIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetStencilCleared]
//   - [IOGPUMetalRenderCommandEncoder.SetStencilStoreAction]
//   - [IOGPUMetalRenderCommandEncoder.SetStencilStoreActionOptions]
//   - [IOGPUMetalRenderCommandEncoder.SetTileAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetTileIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetTileIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetTileVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetTileVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexAccelerationStructureAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexAmplificationCountViewMappings]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexIntersectionFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexIntersectionFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexVisibleFunctionTableAtBufferIndex]
//   - [IOGPUMetalRenderCommandEncoder.SetVertexVisibleFunctionTablesWithBufferRange]
//   - [IOGPUMetalRenderCommandEncoder.TextureBarrier]
//   - [IOGPUMetalRenderCommandEncoder.UpdateFenceAfterStages]
//   - [IOGPUMetalRenderCommandEncoder.UseHeap]
//   - [IOGPUMetalRenderCommandEncoder.UseHeapStages]
//   - [IOGPUMetalRenderCommandEncoder.UseHeapsCount]
//   - [IOGPUMetalRenderCommandEncoder.UseHeapsCountStages]
//   - [IOGPUMetalRenderCommandEncoder.UseResourceUsage]
//   - [IOGPUMetalRenderCommandEncoder.UseResourceUsageStages]
//   - [IOGPUMetalRenderCommandEncoder.UseResourcesCountUsage]
//   - [IOGPUMetalRenderCommandEncoder.UseResourcesCountUsageStages]
//   - [IOGPUMetalRenderCommandEncoder.WaitForFenceBeforeStages]
//   - [IOGPUMetalRenderCommandEncoder.InitWithCommandBufferDescriptor]
type IOGPUMetalRenderCommandEncoder struct {
	IOGPUMetalCommandEncoder
}

// IOGPUMetalRenderCommandEncoderFromID constructs a [IOGPUMetalRenderCommandEncoder] from an objc.ID.
func IOGPUMetalRenderCommandEncoderFromID(id objc.ID) IOGPUMetalRenderCommandEncoder {
	return IOGPUMetalRenderCommandEncoder{IOGPUMetalCommandEncoder: IOGPUMetalCommandEncoderFromID(id)}
}

// Ensure IOGPUMetalRenderCommandEncoder implements IIOGPUMetalRenderCommandEncoder.
var _ IIOGPUMetalRenderCommandEncoder = IOGPUMetalRenderCommandEncoder{}

// An interface definition for the [IOGPUMetalRenderCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalRenderCommandEncoder.ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset]
//   - [IIOGPUMetalRenderCommandEncoder.ExecuteCommandsInBufferWithRange]
//   - [IIOGPUMetalRenderCommandEncoder.GetFragmentBufferContentsAtIndex]
//   - [IIOGPUMetalRenderCommandEncoder.GetRenderPipelineState]
//   - [IIOGPUMetalRenderCommandEncoder.GetType]
//   - [IIOGPUMetalRenderCommandEncoder.GetVertexBufferContentsAtIndex]
//   - [IIOGPUMetalRenderCommandEncoder.IsMemorylessRender]
//   - [IIOGPUMetalRenderCommandEncoder.MemoryBarrierWithResourcesCountAfterStagesBeforeStages]
//   - [IIOGPUMetalRenderCommandEncoder.MemoryBarrierWithScopeAfterStagesBeforeStages]
//   - [IIOGPUMetalRenderCommandEncoder.SampleCountersInBufferAtSampleIndexWithBarrier]
//   - [IIOGPUMetalRenderCommandEncoder.SetColorStoreActionAtIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetColorStoreActionOptionsAtIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetDepthCleared]
//   - [IIOGPUMetalRenderCommandEncoder.SetDepthClipModeSPI]
//   - [IIOGPUMetalRenderCommandEncoder.SetDepthStoreAction]
//   - [IIOGPUMetalRenderCommandEncoder.SetDepthStoreActionOptions]
//   - [IIOGPUMetalRenderCommandEncoder.SetFragmentAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetFragmentIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetFragmentIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetFragmentVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetFragmentVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetMeshAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetMeshIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetMeshIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetObjectAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetObjectIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetObjectIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetStencilCleared]
//   - [IIOGPUMetalRenderCommandEncoder.SetStencilStoreAction]
//   - [IIOGPUMetalRenderCommandEncoder.SetStencilStoreActionOptions]
//   - [IIOGPUMetalRenderCommandEncoder.SetTileAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetTileIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetTileIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetTileVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetTileVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexAccelerationStructureAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexAmplificationCountViewMappings]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexIntersectionFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexIntersectionFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexVisibleFunctionTableAtBufferIndex]
//   - [IIOGPUMetalRenderCommandEncoder.SetVertexVisibleFunctionTablesWithBufferRange]
//   - [IIOGPUMetalRenderCommandEncoder.TextureBarrier]
//   - [IIOGPUMetalRenderCommandEncoder.UpdateFenceAfterStages]
//   - [IIOGPUMetalRenderCommandEncoder.UseHeap]
//   - [IIOGPUMetalRenderCommandEncoder.UseHeapStages]
//   - [IIOGPUMetalRenderCommandEncoder.UseHeapsCount]
//   - [IIOGPUMetalRenderCommandEncoder.UseHeapsCountStages]
//   - [IIOGPUMetalRenderCommandEncoder.UseResourceUsage]
//   - [IIOGPUMetalRenderCommandEncoder.UseResourceUsageStages]
//   - [IIOGPUMetalRenderCommandEncoder.UseResourcesCountUsage]
//   - [IIOGPUMetalRenderCommandEncoder.UseResourcesCountUsageStages]
//   - [IIOGPUMetalRenderCommandEncoder.WaitForFenceBeforeStages]
//   - [IIOGPUMetalRenderCommandEncoder.InitWithCommandBufferDescriptor]
type IIOGPUMetalRenderCommandEncoder interface {
	IIOGPUMetalCommandEncoder

	// Topic: Methods

	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64)
	ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange)
	GetFragmentBufferContentsAtIndex(index uint64) unsafe.Pointer
	GetRenderPipelineState() objectivec.IObject
	GetType() uint64
	GetVertexBufferContentsAtIndex(index uint64) unsafe.Pointer
	IsMemorylessRender() bool
	MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources []objectivec.IObject, count uint64, stages uint64, stages2 uint64)
	MemoryBarrierWithScopeAfterStagesBeforeStages(scope uint64, stages uint64, stages2 uint64)
	SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool)
	SetColorStoreActionAtIndex(action uint64, index uint64)
	SetColorStoreActionOptionsAtIndex(options uint64, index uint64)
	SetDepthCleared()
	SetDepthClipModeSPI(spi uint64)
	SetDepthStoreAction(action uint64)
	SetDepthStoreActionOptions(options uint64)
	SetFragmentAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetFragmentIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetFragmentIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetFragmentVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetFragmentVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetMeshAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetMeshIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetMeshIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetObjectAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetObjectIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetObjectIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetStencilCleared()
	SetStencilStoreAction(action uint64)
	SetStencilStoreActionOptions(options uint64)
	SetTileAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetTileIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetTileIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetTileVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetTileVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetVertexAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64)
	SetVertexAmplificationCountViewMappings(count uint64, mappings unsafe.Pointer)
	SetVertexIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVertexIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	SetVertexVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64)
	SetVertexVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)
	TextureBarrier()
	UpdateFenceAfterStages(fence objectivec.IObject, stages uint64)
	UseHeap(heap objectivec.IObject)
	UseHeapStages(heap objectivec.IObject, stages uint64)
	UseHeapsCount(heaps []objectivec.IObject, count uint64)
	UseHeapsCountStages(heaps []objectivec.IObject, count uint64, stages uint64)
	UseResourceUsage(resource objectivec.IObject, usage uint64)
	UseResourceUsageStages(resource objectivec.IObject, usage uint64, stages uint64)
	UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64)
	UseResourcesCountUsageStages(resources []objectivec.IObject, count uint64, usage uint64, stages uint64)
	WaitForFenceBeforeStages(fence objectivec.IObject, stages uint64)
	InitWithCommandBufferDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalRenderCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetalRenderCommandEncoder) Init() IOGPUMetalRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalRenderCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalRenderCommandEncoder) Autorelease() IOGPUMetalRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalRenderCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalRenderCommandEncoder creates a new IOGPUMetalRenderCommandEncoder instance.
func NewIOGPUMetalRenderCommandEncoder() IOGPUMetalRenderCommandEncoder {
	class := getIOGPUMetalRenderCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalRenderCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalRenderCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalRenderCommandEncoder {
	instance := getIOGPUMetalRenderCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalRenderCommandEncoderFromID(rv)
}

func NewGPUMetalRenderCommandEncoderWithCommandBufferDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalRenderCommandEncoder {
	instance := getIOGPUMetalRenderCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:descriptor:"), buffer, descriptor)
	return IOGPUMetalRenderCommandEncoderFromID(rv)
}

func (i IOGPUMetalRenderCommandEncoder) ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:"), buffer, buffer2, offset)
}
func (i IOGPUMetalRenderCommandEncoder) ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("executeCommandsInBuffer:withRange:"), buffer, range_)
}
func (i IOGPUMetalRenderCommandEncoder) GetFragmentBufferContentsAtIndex(index uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getFragmentBufferContentsAtIndex:"), index)
	return rv
}
func (i IOGPUMetalRenderCommandEncoder) GetRenderPipelineState() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getRenderPipelineState"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalRenderCommandEncoder) GetType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetalRenderCommandEncoder) GetVertexBufferContentsAtIndex(index uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getVertexBufferContentsAtIndex:"), index)
	return rv
}
func (i IOGPUMetalRenderCommandEncoder) IsMemorylessRender() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isMemorylessRender"))
	return rv
}
func (i IOGPUMetalRenderCommandEncoder) MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources []objectivec.IObject, count uint64, stages uint64, stages2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("memoryBarrierWithResources:count:afterStages:beforeStages:"), objc.CArray(resources), count, stages, stages2)
}
func (i IOGPUMetalRenderCommandEncoder) MemoryBarrierWithScopeAfterStagesBeforeStages(scope uint64, stages uint64, stages2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("memoryBarrierWithScope:afterStages:beforeStages:"), scope, stages, stages2)
}
func (i IOGPUMetalRenderCommandEncoder) SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), buffer, index, barrier)
}
func (i IOGPUMetalRenderCommandEncoder) SetColorStoreActionAtIndex(action uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setColorStoreAction:atIndex:"), action, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetColorStoreActionOptionsAtIndex(options uint64, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setColorStoreActionOptions:atIndex:"), options, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetDepthCleared() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setDepthCleared"))
}
func (i IOGPUMetalRenderCommandEncoder) SetDepthClipModeSPI(spi uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setDepthClipModeSPI:"), spi)
}
func (i IOGPUMetalRenderCommandEncoder) SetDepthStoreAction(action uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setDepthStoreAction:"), action)
}
func (i IOGPUMetalRenderCommandEncoder) SetDepthStoreActionOptions(options uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setDepthStoreActionOptions:"), options)
}
func (i IOGPUMetalRenderCommandEncoder) SetFragmentAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFragmentAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetFragmentIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFragmentIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetFragmentIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFragmentIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetFragmentVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFragmentVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetFragmentVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setFragmentVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetMeshAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setMeshAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetMeshIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setMeshIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetMeshIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setMeshIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetObjectAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setObjectAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetObjectIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setObjectIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetObjectIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setObjectIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetStencilCleared() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setStencilCleared"))
}
func (i IOGPUMetalRenderCommandEncoder) SetStencilStoreAction(action uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setStencilStoreAction:"), action)
}
func (i IOGPUMetalRenderCommandEncoder) SetStencilStoreActionOptions(options uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setStencilStoreActionOptions:"), options)
}
func (i IOGPUMetalRenderCommandEncoder) SetTileAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTileAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetTileIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTileIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetTileIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTileIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetTileVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTileVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetTileVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setTileVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexAccelerationStructure:atBufferIndex:"), structure, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexAmplificationCountViewMappings(count uint64, mappings unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexAmplificationCount:viewMappings:"), count, mappings)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (i IOGPUMetalRenderCommandEncoder) SetVertexVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setVertexVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (i IOGPUMetalRenderCommandEncoder) TextureBarrier() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("textureBarrier"))
}
func (i IOGPUMetalRenderCommandEncoder) UpdateFenceAfterStages(fence objectivec.IObject, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateFence:afterStages:"), fence, stages)
}
func (i IOGPUMetalRenderCommandEncoder) UseHeap(heap objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeap:"), heap)
}
func (i IOGPUMetalRenderCommandEncoder) UseHeapStages(heap objectivec.IObject, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeap:stages:"), heap, stages)
}
func (i IOGPUMetalRenderCommandEncoder) UseHeapsCount(heaps []objectivec.IObject, count uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
}
func (i IOGPUMetalRenderCommandEncoder) UseHeapsCountStages(heaps []objectivec.IObject, count uint64, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useHeaps:count:stages:"), objc.CArray(heaps), count, stages)
}
func (i IOGPUMetalRenderCommandEncoder) UseResourceUsage(resource objectivec.IObject, usage uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResource:usage:"), resource, usage)
}
func (i IOGPUMetalRenderCommandEncoder) UseResourceUsageStages(resource objectivec.IObject, usage uint64, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResource:usage:stages:"), resource, usage, stages)
}
func (i IOGPUMetalRenderCommandEncoder) UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResources:count:usage:"), objc.CArray(resources), count, usage)
}
func (i IOGPUMetalRenderCommandEncoder) UseResourcesCountUsageStages(resources []objectivec.IObject, count uint64, usage uint64, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("useResources:count:usage:stages:"), objc.CArray(resources), count, usage, stages)
}
func (i IOGPUMetalRenderCommandEncoder) WaitForFenceBeforeStages(fence objectivec.IObject, stages uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("waitForFence:beforeStages:"), fence, stages)
}
func (i IOGPUMetalRenderCommandEncoder) InitWithCommandBufferDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalRenderCommandEncoder](i.ID, objc.Sel("initWithCommandBuffer:descriptor:"), buffer, descriptor)
	return rv
}
