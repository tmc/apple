// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLComputeCommandEncoder protocol.
type MTLComputeCommandEncoder interface {
	objectivec.IObject

	// DispatchThreadgroupsThreadsPerThreadgroup protocol.
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroups unsafe.Pointer, threadgroup unsafe.Pointer)

	// DispatchThreadsThreadsPerThreadgroup protocol.
	DispatchThreadsThreadsPerThreadgroup(threads unsafe.Pointer, threadgroup unsafe.Pointer)

	// DispatchType protocol.
	DispatchType() uint64

	// MemoryBarrierWithResourcesCount protocol.
	MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64)

	// MemoryBarrierWithScope protocol.
	MemoryBarrierWithScope(scope uint64)

	// SetBufferOffsetAtIndex protocol.
	SetBufferOffsetAtIndex(offset uint64, index uint64)

	// SetBufferOffsetAttributeStrideAtIndex protocol.
	SetBufferOffsetAttributeStrideAtIndex(offset uint64, stride uint64, index uint64)

	// SetBuffersOffsetsAttributeStridesWithRange protocol.
	SetBuffersOffsetsAttributeStridesWithRange(buffers []objectivec.IObject, offsets *uint64, strides *uint64, range_ foundation.NSRange)

	// SetBuffersOffsetsWithRange protocol.
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange)

	// SetBytesLengthAtIndex protocol.
	SetBytesLengthAtIndex(bytes []byte, index uint64)

	// SetBytesLengthAttributeStrideAtIndex protocol.
	SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint64, index uint64)

	// SetImageblockWidthHeight protocol.
	SetImageblockWidthHeight(width uint64, height uint64)

	// SetIntersectionFunctionTablesWithBufferRange protocol.
	SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)

	// SetSamplerStatesLodMinClampsLodMaxClampsWithRange protocol.
	SetSamplerStatesLodMinClampsLodMaxClampsWithRange(states []objectivec.IObject, clamps *float32, clamps2 *float32, range_ foundation.NSRange)

	// SetSamplerStatesWithRange protocol.
	SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange)

	// SetStageInRegion protocol.
	SetStageInRegion(region unsafe.Pointer)

	// SetTexturesWithRange protocol.
	SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange)

	// SetThreadgroupMemoryLengthAtIndex protocol.
	SetThreadgroupMemoryLengthAtIndex(length uint64, index uint64)

	// SetVisibleFunctionTablesWithBufferRange protocol.
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)

	// UseHeapsCount protocol.
	UseHeapsCount(heaps []objectivec.IObject, count uint64)

	// UseResourcesCountUsage protocol.
	UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64)
}

// MTLComputeCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLComputeCommandEncoder protocol.
type MTLComputeCommandEncoderObject struct {
	objectivec.Object
}

func (o MTLComputeCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLComputeCommandEncoderObjectFromID constructs a [MTLComputeCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLComputeCommandEncoderObjectFromID(id objc.ID) MTLComputeCommandEncoderObject {
	return MTLComputeCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsThreadsPerThreadgroup(threadgroups unsafe.Pointer, threadgroup unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"), threadgroups, threadgroup)
}
func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(buffer objectivec.IObject, offset uint64, threadgroup unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:"), buffer, offset, threadgroup)
}
func (o MTLComputeCommandEncoderObject) DispatchThreadsThreadsPerThreadgroup(threads unsafe.Pointer, threadgroup unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreads:threadsPerThreadgroup:"), threads, threadgroup)
}
func (o MTLComputeCommandEncoderObject) DispatchType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("dispatchType"))
	return rv
}
func (o MTLComputeCommandEncoderObject) ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:"), buffer, buffer2, offset)
}
func (o MTLComputeCommandEncoderObject) ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:withRange:"), buffer, range_)
}
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithResources:count:"), objc.CArray(resources), count)
}
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithScope(scope uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithScope:"), scope)
}
func (o MTLComputeCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool) {
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), buffer, index, barrier)
}
func (o MTLComputeCommandEncoderObject) SetAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccelerationStructure:atBufferIndex:"), structure, index)
}
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAttributeStrideAtIndex(buffer objectivec.IObject, offset uint64, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:attributeStride:atIndex:"), buffer, offset, stride, index)
}
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAtIndex(offset uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:atIndex:"), offset, index)
}
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAttributeStrideAtIndex(offset uint64, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:attributeStride:atIndex:"), offset, stride, index)
}
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsAttributeStridesWithRange(buffers []objectivec.IObject, offsets *uint64, strides *uint64, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:attributeStrides:withRange:"), objectivec.IObjectSliceToNSArray(buffers), offsets, strides, range_)
}
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets *uint64, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), offsets, range_)
}
func (o MTLComputeCommandEncoderObject) SetBytesLengthAtIndex(bytes []byte, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
}
func (o MTLComputeCommandEncoderObject) SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:attributeStride:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), stride, index)
}
func (o MTLComputeCommandEncoderObject) SetComputePipelineState(state objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:"), state)
}
func (o MTLComputeCommandEncoderObject) SetImageblockWidthHeight(width uint64, height uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setImageblockWidth:height:"), width, height)
}
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), table, index)
}
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (o MTLComputeCommandEncoderObject) SetSamplerStateAtIndex(state objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:atIndex:"), state, index)
}
func (o MTLComputeCommandEncoderObject) SetSamplerStateLodMinClampLodMaxClampAtIndex(state objectivec.IObject, clamp float32, clamp2 float32, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), state, clamp, clamp2, index)
}
func (o MTLComputeCommandEncoderObject) SetSamplerStatesLodMinClampsLodMaxClampsWithRange(states []objectivec.IObject, clamps *float32, clamps2 *float32, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), objectivec.IObjectSliceToNSArray(states), clamps, clamps2, range_)
}
func (o MTLComputeCommandEncoderObject) SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:withRange:"), objectivec.IObjectSliceToNSArray(states), range_)
}
func (o MTLComputeCommandEncoderObject) SetStageInRegion(region unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegion:"), region)
}
func (o MTLComputeCommandEncoderObject) SetStageInRegionWithIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, offset uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegionWithIndirectBuffer:indirectBufferOffset:"), buffer, offset)
}
func (o MTLComputeCommandEncoderObject) SetTextureAtIndex(texture objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setTexture:atIndex:"), texture, index)
}
func (o MTLComputeCommandEncoderObject) SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextures:withRange:"), objectivec.IObjectSliceToNSArray(textures), range_)
}
func (o MTLComputeCommandEncoderObject) SetThreadgroupMemoryLengthAtIndex(length uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
}
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
}
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
}
func (o MTLComputeCommandEncoderObject) UpdateFence(fence objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
}
func (o MTLComputeCommandEncoderObject) UseHeap(heap objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeap:"), heap)
}
func (o MTLComputeCommandEncoderObject) UseHeapsCount(heaps []objectivec.IObject, count uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
}
func (o MTLComputeCommandEncoderObject) UseResourceUsage(resource objectivec.IObject, usage uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useResource:usage:"), resource, usage)
}
func (o MTLComputeCommandEncoderObject) UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useResources:count:usage:"), objc.CArray(resources), count, usage)
}
func (o MTLComputeCommandEncoderObject) WaitForFence(fence objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
}
