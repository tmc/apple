// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MTLComputeCommandEncoder protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder
type MTLComputeCommandEncoder interface {
	objectivec.IObject

	// DispatchThreadgroupsThreadsPerThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchThreadgroups:threadsPerThreadgroup:
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroups objectivec.IObject, threadgroup objectivec.IObject)

	// DispatchThreadsThreadsPerThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchThreads:threadsPerThreadgroup:
	DispatchThreadsThreadsPerThreadgroup(threads objectivec.IObject, threadgroup objectivec.IObject)

	// DispatchType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchType
	DispatchType() uint64

	// MemoryBarrierWithResourcesCount protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/memoryBarrierWithResources:count:
	MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64)

	// MemoryBarrierWithScope protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/memoryBarrierWithScope:
	MemoryBarrierWithScope(scope uint64)

	// SetBufferOffsetAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBufferOffset:atIndex:
	SetBufferOffsetAtIndex(offset uint64, index uint64)

	// SetBufferOffsetAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBufferOffset:attributeStride:atIndex:
	SetBufferOffsetAttributeStrideAtIndex(offset uint64, stride uint64, index uint64)

	// SetBuffersOffsetsAttributeStridesWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffers:offsets:attributeStrides:withRange:
	SetBuffersOffsetsAttributeStridesWithRange(buffers []objectivec.IObject, offsets unsafe.Pointer, strides unsafe.Pointer, range_ foundation.NSRange)

	// SetBuffersOffsetsWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffers:offsets:withRange:
	SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets unsafe.Pointer, range_ foundation.NSRange)

	// SetBytesLengthAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBytes:length:atIndex:
	SetBytesLengthAtIndex(bytes []byte, index uint64)

	// SetBytesLengthAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBytes:length:attributeStride:atIndex:
	SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint64, index uint64)

	// SetImageblockWidthHeight protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setImageblockWidth:height:
	SetImageblockWidthHeight(width uint64, height uint64)

	// SetIntersectionFunctionTablesWithBufferRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setIntersectionFunctionTables:withBufferRange:
	SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)

	// SetSamplerStatesLodMinClampsLodMaxClampsWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetSamplerStatesLodMinClampsLodMaxClampsWithRange(states []objectivec.IObject, clamps unsafe.Pointer, clamps2 unsafe.Pointer, range_ foundation.NSRange)

	// SetSamplerStatesWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerStates:withRange:
	SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange)

	// SetStageInRegion protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setStageInRegion:
	SetStageInRegion(region objectivec.IObject)

	// SetTexturesWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setTextures:withRange:
	SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange)

	// SetThreadgroupMemoryLengthAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setThreadgroupMemoryLength:atIndex:
	SetThreadgroupMemoryLengthAtIndex(length uint64, index uint64)

	// SetVisibleFunctionTablesWithBufferRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setVisibleFunctionTables:withBufferRange:
	SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange)

	// UseHeapsCount protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useHeaps:count:
	UseHeapsCount(heaps []objectivec.IObject, count uint64)

	// UseResourcesCountUsage protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useResources:count:usage:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchThreadgroups:threadsPerThreadgroup:
func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsThreadsPerThreadgroup(threadgroups objectivec.IObject, threadgroup objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"), threadgroups, threadgroup)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:
func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(buffer objectivec.IObject, offset uint64, threadgroup objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:"), buffer, offset, threadgroup)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchThreads:threadsPerThreadgroup:
func (o MTLComputeCommandEncoderObject) DispatchThreadsThreadsPerThreadgroup(threads objectivec.IObject, threadgroup objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreads:threadsPerThreadgroup:"), threads, threadgroup)
	}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/dispatchType
func (o MTLComputeCommandEncoderObject) DispatchType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("dispatchType"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:
func (o MTLComputeCommandEncoderObject) ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, buffer2 objectivec.IObject, offset uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:"), buffer, buffer2, offset)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/executeCommandsInBuffer:withRange:
func (o MTLComputeCommandEncoderObject) ExecuteCommandsInBufferWithRange(buffer objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:withRange:"), buffer, range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/memoryBarrierWithResources:count:
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithResourcesCount(resources []objectivec.IObject, count uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithResources:count:"), objc.CArray(resources), count)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/memoryBarrierWithScope:
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithScope(scope uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithScope:"), scope)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/sampleCountersInBuffer:atSampleIndex:withBarrier:
func (o MTLComputeCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(buffer objectivec.IObject, index uint64, barrier bool) {
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), buffer, index, barrier)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setAccelerationStructure:atBufferIndex:
func (o MTLComputeCommandEncoderObject) SetAccelerationStructureAtBufferIndex(structure objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccelerationStructure:atBufferIndex:"), structure, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffer:offset:atIndex:
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAtIndex(buffer objectivec.IObject, offset uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffer:offset:attributeStride:atIndex:
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAttributeStrideAtIndex(buffer objectivec.IObject, offset uint64, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:attributeStride:atIndex:"), buffer, offset, stride, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBufferOffset:atIndex:
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAtIndex(offset uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:atIndex:"), offset, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBufferOffset:attributeStride:atIndex:
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAttributeStrideAtIndex(offset uint64, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:attributeStride:atIndex:"), offset, stride, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffers:offsets:attributeStrides:withRange:
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsAttributeStridesWithRange(buffers []objectivec.IObject, offsets unsafe.Pointer, strides unsafe.Pointer, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:attributeStrides:withRange:"), objectivec.IObjectSliceToNSArray(buffers), offsets, strides, range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBuffers:offsets:withRange:
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsWithRange(buffers []objectivec.IObject, offsets unsafe.Pointer, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), objectivec.IObjectSliceToNSArray(buffers), offsets, range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBytes:length:atIndex:
func (o MTLComputeCommandEncoderObject) SetBytesLengthAtIndex(bytes []byte, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setBytes:length:attributeStride:atIndex:
func (o MTLComputeCommandEncoderObject) SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:attributeStride:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), stride, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setComputePipelineState:
func (o MTLComputeCommandEncoderObject) SetComputePipelineState(state objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:"), state)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setImageblockWidth:height:
func (o MTLComputeCommandEncoderObject) SetImageblockWidthHeight(width uint64, height uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setImageblockWidth:height:"), width, height)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setIntersectionFunctionTable:atBufferIndex:
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), table, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setIntersectionFunctionTables:withBufferRange:
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerState:atIndex:
func (o MTLComputeCommandEncoderObject) SetSamplerStateAtIndex(state objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:atIndex:"), state, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerState:lodMinClamp:lodMaxClamp:atIndex:
func (o MTLComputeCommandEncoderObject) SetSamplerStateLodMinClampLodMaxClampAtIndex(state objectivec.IObject, clamp float32, clamp2 float32, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), state, clamp, clamp2, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLComputeCommandEncoderObject) SetSamplerStatesLodMinClampsLodMaxClampsWithRange(states []objectivec.IObject, clamps unsafe.Pointer, clamps2 unsafe.Pointer, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), objectivec.IObjectSliceToNSArray(states), clamps, clamps2, range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setSamplerStates:withRange:
func (o MTLComputeCommandEncoderObject) SetSamplerStatesWithRange(states []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:withRange:"), objectivec.IObjectSliceToNSArray(states), range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setStageInRegion:
func (o MTLComputeCommandEncoderObject) SetStageInRegion(region objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegion:"), region)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setStageInRegionWithIndirectBuffer:indirectBufferOffset:
func (o MTLComputeCommandEncoderObject) SetStageInRegionWithIndirectBufferIndirectBufferOffset(buffer objectivec.IObject, offset uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegionWithIndirectBuffer:indirectBufferOffset:"), buffer, offset)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setTexture:atIndex:
func (o MTLComputeCommandEncoderObject) SetTextureAtIndex(texture objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setTexture:atIndex:"), texture, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setTextures:withRange:
func (o MTLComputeCommandEncoderObject) SetTexturesWithRange(textures []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextures:withRange:"), objectivec.IObjectSliceToNSArray(textures), range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setThreadgroupMemoryLength:atIndex:
func (o MTLComputeCommandEncoderObject) SetThreadgroupMemoryLengthAtIndex(length uint64, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setVisibleFunctionTable:atBufferIndex:
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTableAtBufferIndex(table objectivec.IObject, index uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), table, index)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/setVisibleFunctionTables:withBufferRange:
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTablesWithBufferRange(tables []objectivec.IObject, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), objectivec.IObjectSliceToNSArray(tables), range_)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/updateFence:
func (o MTLComputeCommandEncoderObject) UpdateFence(fence objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useHeap:
func (o MTLComputeCommandEncoderObject) UseHeap(heap objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeap:"), heap)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useHeaps:count:
func (o MTLComputeCommandEncoderObject) UseHeapsCount(heaps []objectivec.IObject, count uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useResource:usage:
func (o MTLComputeCommandEncoderObject) UseResourceUsage(resource objectivec.IObject, usage uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useResource:usage:"), resource, usage)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/useResources:count:usage:
func (o MTLComputeCommandEncoderObject) UseResourcesCountUsage(resources []objectivec.IObject, count uint64, usage uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("useResources:count:usage:"), objc.CArray(resources), count, usage)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/MTLComputeCommandEncoder/waitForFence:
func (o MTLComputeCommandEncoderObject) WaitForFence(fence objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
	}

