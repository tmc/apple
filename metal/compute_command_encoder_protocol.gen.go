// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Encodes computation dispatch commands for a single compute pass into a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder
type MTLComputeCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Configures the compute encoder with a pipeline state for subsequent kernel calls.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setComputePipelineState(_:)
	SetComputePipelineState(state MTLComputePipelineState)

	// The dispatch type to use when submitting compute work to the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchType
	DispatchType() MTLDispatchType

	// Binds a buffer to the buffer argument table, allowing compute kernels to access its data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:index:)
	SetBufferWithOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Binds a buffer with a stride to the buffer argument table, allowing compute kernels to access its data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:attributeStride:index:)
	SetBufferWithOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint)

	// Changes where the data begins in a buffer already bound to the buffer argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBufferOffset(_:index:)
	SetBufferOffsetAtIndex(offset uint, index uint)

	// Changes where the data begins and the distance between adjacent elements in a buffer already bound to the buffer argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBufferOffset(offset:attributeStride:index:)
	SetBufferOffsetAttributeStrideAtIndex(offset uint, stride uint, index uint)

	// Copies data directly to the GPU to populate an entry in the buffer argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBytes(_:length:index:)
	SetBytesLengthAtIndex(bytes []byte, index uint)

	// Copies data with a given stride directly to the GPU to populate an entry in the buffer argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBytes(_:length:attributeStride:index:)
	SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint, index uint)

	// Binds a texture to the texture argument table, allowing compute kernels to access its data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setTexture(_:index:)
	SetTextureAtIndex(texture MTLTexture, index uint)

	// Encodes a texture sampler, allowing compute kernels to use it for sampling textures on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerState(_:index:)
	SetSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Encodes a texture sampler with a custom level of detail clamping, allowing compute kernels to use it for sampling textures on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Binds a visible function table to the buffer argument table, allowing you to call its functions on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setVisibleFunctionTable(_:bufferIndex:)
	SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable MTLVisibleFunctionTable, bufferIndex uint)

	// Binds an acceleration structure to the buffer argument table, allowing functions to access it on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setAccelerationStructure(_:bufferIndex:)
	SetAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint)

	// Binds an intersection function table to the buffer argument table, making it callable in your Metal shaders.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setIntersectionFunctionTable(_:bufferIndex:)
	SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint)

	// Ensures kernel calls that the system encodes in subsequent commands have access to a resource.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useResource(_:usage:)
	UseResourceUsage(resource MTLResource, usage MTLResourceUsage)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to all of the resources you allocate from a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useHeap(_:)
	UseHeap(heap MTLHeap)

	// Configures the size of a block of threadgroup memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setThreadgroupMemoryLength(_:index:)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// Sets the size, in pixels, of imageblock data in tile memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setImageblockWidth(_:height:)
	SetImageblockWidthHeight(width uint, height uint)

	// Sets the dimensions over the thread grid of how your compute kernel receives stage-in arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setStageInRegion(_:)
	SetStageInRegion(region MTLRegion)

	// Sets the region of the stage-in attributes to apply to a compute kernel using an indirect buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setStageInRegionWithIndirectBuffer(_:indirectBufferOffset:)
	SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a compute command using an arbitrarily sized grid.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreads(_:threadsPerThreadgroup:)
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Encodes a compute dispatch command using a grid aligned to threadgroup boundaries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreadgroups(_:threadsPerThreadgroup:)
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Encodes a dispatch call for a compute pass, using an indirect buffer that defines the size of a grid that aligns to threadgroup boundaries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreadgroups(indirectBuffer:indirectBufferOffset:threadsPerThreadgroup:)
	DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer MTLBuffer, indirectBufferOffset uint, threadsPerThreadgroup MTLSize)

	// Encodes a command that instructs the GPU to pause the compute pass until another pass updates a fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/waitForFence(_:)
	WaitForFence(fence MTLFence)

	// Encodes a command that instructs the GPU to update a fence after the compute pass completes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/updateFence(_:)
	UpdateFence(fence MTLFence)

	// Creates a memory barrier that enforces the order of write and read operations for specific resource types.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/memoryBarrier(scope:)
	MemoryBarrierWithScope(scope MTLBarrierScope)

	// Encodes a command to sample hardware counters, providing performance information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool)

	// Creates a memory barrier that enforces the order of write and read operations for specific resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/memoryBarrierWithResources:count:
	MemoryBarrierWithResourcesCount(resources []MTLResource, count uint)

	// Binds multiple buffers with data in stride to the buffer argument table at once, allowing compute kernels to access their data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffers:offsets:attributeStrides:withRange:
	SetBuffersOffsetsAttributeStridesWithRange(buffers []MTLBuffer, offsets uint, strides uint, range_ foundation.NSRange)

	// Binds multiple buffers to the buffer argument table at once, allowing compute kernels to access their data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffers:offsets:withRange:
	SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Binds multiple intersection function tables to the buffer argument table, allowing you to call their functions on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setIntersectionFunctionTables:withBufferRange:
	SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange)

	// Encodes multiple texture samplers with custom levels of detail clamping, allowing compute kernels to use them for sampling textures on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Encodes multiple texture samplers, allowing compute kernels to use them for sampling textures on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerStates:withRange:
	SetSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Binds multiple textures to the texture argument table, allowing compute kernels to access their data on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setTextures:withRange:
	SetTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Binds multiple visible function tables to the buffer argument table, allowing you to call their functions on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setVisibleFunctionTables:withBufferRange:
	SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables []MTLVisibleFunctionTable, range_ foundation.NSRange)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to all of the resources you allocate from multiple heaps.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useHeaps:count:
	UseHeapsCount(heaps []MTLHeap, count uint)

	// Ensures kernel calls that the system encodes in subsequent commands have access to multiple resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useResources:count:usage:
	UseResourcesCountUsage(resources []MTLResource, count uint, usage MTLResourceUsage)
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

// Configures the compute encoder with a pipeline state for subsequent kernel
// calls.
//
// state: An [MTLComputePipelineState] instance.
//
// # Discussion
//
// Create your pipeline state through one of the [MTLDevice] methods in
// Creating Compute Pipeline States.
//
// A compute pipeline state provides information Metal uses to compile and run
// encoded commands. You can change the pipeline state at any time, allowing
// you to encode multiple kernel calls in a single command buffer. Changing
// the pipeline state doesn’t affect any previously encoded commands.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setComputePipelineState(_:)
func (o MTLComputeCommandEncoderObject) SetComputePipelineState(state MTLComputePipelineState) {
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:"), state)
}

// The dispatch type to use when submitting compute work to the GPU.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchType
func (o MTLComputeCommandEncoderObject) DispatchType() MTLDispatchType {
	rv := objc.Send[MTLDispatchType](o.ID, objc.Sel("dispatchType"))
	return rv
}

// Binds a buffer to the buffer argument table, allowing compute kernels to
// access its data on the GPU.
//
// buffer: The [MTLBuffer] instance to bind to the argument table.
//
// offset: The number of bytes to skip in the buffer before the first element of data.
//
// index: The index the buffer binds to in the argument table.
//
// # Discussion
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// Rebinding an already bound buffer causes a Metal error.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:index:)
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
}

// Binds a buffer with a stride to the buffer argument table, allowing compute
// kernels to access its data on the GPU.
//
// buffer: The [MTLBuffer] instance to bind to the argument table.
//
// offset: The number of bytes to skip in the buffer before the first element of data.
//
// stride: The number of bytes between the start of one element and the start of the
// next.
//
// index: The index the buffer binds to in the argument table.
//
// # Discussion
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// Rebinding an already bound buffer causes a Metal error.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:attributeStride:index:)
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBufferWithOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:attributeStride:atIndex:"), buffer, offset, stride, index)
}

// Changes where the data begins in a buffer already bound to the buffer
// argument table.
//
// offset: Where the data to bind begins, in bytes, from the start of the bound
// buffer.
//
// index: The argument table entry to change.
//
// # Discussion
//
// Prefer calling this method to unbinding and then rebinding data.
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBufferOffset(_:index:)
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAtIndex(offset uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:atIndex:"), offset, index)
}

// Changes where the data begins and the distance between adjacent elements in
// a buffer already bound to the buffer argument table.
//
// offset: Where the data to bind begins, in bytes, from the start of the bound
// buffer.
//
// stride: The number of bytes between the start of one element and the start of the
// next.
//
// index: The index of the buffer to change in the argument table.
//
// # Discussion
//
// _ _Prefer calling this method to unbinding and then rebinding data.
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBufferOffset(offset:attributeStride:index:)
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBufferOffsetAttributeStrideAtIndex(offset uint, stride uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBufferOffset:attributeStride:atIndex:"), offset, stride, index)
}

// Copies data directly to the GPU to populate an entry in the buffer argument
// table.
//
// bytes: A pointer to where the data to copy starts.
//
// length: The number of bytes to copy.
//
// index: The index the data binds to in the argument table.
//
// # Discussion
//
// This method allows Metal to copy data efficiently onto the GPU without the
// need for your own buffer. Binding data directly can improve performance,
// especially when making many small allocations.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBytes(_:length:index:)
func (o MTLComputeCommandEncoderObject) SetBytesLengthAtIndex(bytes []byte, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
}

// Copies data with a given stride directly to the GPU to populate an entry in
// the buffer argument table.
//
// bytes: A pointer to the memory where the data to copy starts.
//
// length: The number of bytes to copy.
//
// stride: The number of bytes between the start of one element and the start of the
// next.
//
// index: The index the data binds to in the argument table.
//
// # Discussion
//
// This method allows Metal to copy data directly onto the GPU, rather than
// creating a new [MTLBuffer] instance and binding it. Binding data directly
// can improve performance, especially when making many small allocations.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBytes(_:length:attributeStride:index:)
func (o MTLComputeCommandEncoderObject) SetBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBytes:length:attributeStride:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), stride, index)
}

// Binds a texture to the texture argument table, allowing compute kernels to
// access its data on the GPU.
//
// texture: An [MTLTexture] instance to bind to the texture argument table.
//
// index: The index the texture binds to in the texture argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setTexture(_:index:)
func (o MTLComputeCommandEncoderObject) SetTextureAtIndex(texture MTLTexture, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTexture:atIndex:"), texture, index)
}

// Encodes a texture sampler, allowing compute kernels to use it for sampling
// textures on the GPU.
//
// sampler: An [MTLSamplerState] instance to bind to the sampler argument table.
//
// index: The index in the sampler argument table to bind the sampler to.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerState(_:index:)
func (o MTLComputeCommandEncoderObject) SetSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:atIndex:"), sampler, index)
}

// Encodes a texture sampler with a custom level of detail clamping, allowing
// compute kernels to use it for sampling textures on the GPU.
//
// sampler: An [MTLSamplerState] instance to bind to the sampler argument table.
//
// lodMinClamp: The minimum level of detail used when sampling a texture.
//
// lodMaxClamp: The maximum level of detail used when sampling a texture.
//
// index: The index in the sampler argument table to bind the sampler to.
//
// # Discussion
//
// Calling this method ignores the [LodMinClamp] and [LodMaxClamp] properties
// of the sampler, using the provided levels of detail instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLComputeCommandEncoderObject) SetSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
}

// Binds a visible function table to the buffer argument table, allowing you
// to call its functions on the GPU.
//
// visibleFunctionTable: The [MTLVisibleFunctionTable] to bind.
//
// bufferIndex: The index the function table binds to in the buffer argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setVisibleFunctionTable(_:bufferIndex:)
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable MTLVisibleFunctionTable, bufferIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), visibleFunctionTable, bufferIndex)
}

// Binds an acceleration structure to the buffer argument table, allowing
// functions to access it on the GPU.
//
// accelerationStructure: An [MTLAccelerationStructure] instance to bind to the argument table.
//
// bufferIndex: The index the structure binds to in the argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setAccelerationStructure(_:bufferIndex:)
func (o MTLComputeCommandEncoderObject) SetAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccelerationStructure:atBufferIndex:"), accelerationStructure, bufferIndex)
}

// Binds an intersection function table to the buffer argument table, making
// it callable in your Metal shaders.
//
// intersectionFunctionTable: The [MTLIntersectionFunctionTable] to bind.
//
// bufferIndex: The index in the buffer argument table the intersection function table
// binds to.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setIntersectionFunctionTable(_:bufferIndex:)
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), intersectionFunctionTable, bufferIndex)
}

// Ensures kernel calls that the system encodes in subsequent commands have
// access to a resource.
//
// resource: An [MTLResource] instance used in an argument buffer.
//
// usage: How the compute pass can access data, including [MTLResourceUsageRead] and
// [MTLResourceUsageWrite] permission.
//
// For applicable resources, you may be able to prevent the GPU from
// unnecessarily decompressing color attachments on some devices by setting
// `usage` to [MTLResourceUsageRead].
//
// # Discussion
//
// You can make a resource (available in GPU memory) for the remaining
// duration of the compute pass by calling this method. Call the method before
// encoding function calls that may access the `resource` through an argument
// buffer. The method ensures the resource is in a format that’s compatible
// with the kernels that depend on it.
//
// The method also informs Metal when to apply hazard tracking for a resource
// you create with [MTLHazardTrackingModeTracked]. For a resource you create
// with [MTLHazardTrackingModeUntracked], you need to apply an [MTLFence] or
// an [MTLEvent] to account for potential reading and writing hazards.
//
// You can reconfigure an individual resource’s `usage` options for
// subsequent draw calls in the same render pass by calling this method again.
//
// Apps typically call this method for a resource in an argument buffer as a
// part of their implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useResource(_:usage:)
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
func (o MTLComputeCommandEncoderObject) UseResourceUsage(resource MTLResource, usage MTLResourceUsage) {
	objc.Send[struct{}](o.ID, objc.Sel("useResource:usage:"), resource, usage)
}

// Ensures the shaders in the render pass’s subsequent draw commands have
// access to all of the resources you allocate from a heap.
//
// heap: An [MTLHeap] instance containing resources used in an argument buffer.
//
// # Discussion
//
// You can make the resources in `heap` (available in GPU memory) for the
// remaining duration of the render pass by calling this method. Call the
// method before encoding draw calls that may access resources within `heap`
// through an argument buffer. The method ensures each resource is in a format
// that’s compatible with the shaders that depend on it.
//
// This method applies the [MTLResourceUsageRead] resource usage option to all
// of the resources within the `heaps`, except for textures. The method
// ignores any texture that has [MTLTextureUsageRenderTarget],
// [MTLTextureUsageShaderWrite], or both in its [Usage] property. For all
// other textures in each of the `heaps`, the method optimizes each
// texture’s memory layout for rendering with a sampler. However, your
// kernels can’t read from those textures by calling this method because the
// texture needs a different memory layout that’s suitable for reading.
//
// Methods that apply a usage option for resources (see Encoding Resident
// Resources) override any previous calls that apply to a resource. For
// example, you can change the usage option of any heap in `heaps` to
// [MTLResourceUsageWrite] by passing it to [UseResourceUsageStages] after
// calling this method. However, you can’t reverse the call order because
// this method resets the usage for all resources within `heap` to
// [MTLResourceUsageRead], overriding previous calls to [UseResourceUsage].
//
// This method instructs Metal to apply hazard tracking for resources you
// allocate from a heap that you create with [MTLHazardTrackingModeTracked].
// However, for untracked resources — which come from heaps you create with
// [MTLHazardTrackingModeUntracked] — you need to account for hazards by
// applying [MTLFence] or [MTLEvent] instances.
//
// Apps typically call the method for heaps that have resources in argument
// buffers for a implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useHeap(_:)
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
func (o MTLComputeCommandEncoderObject) UseHeap(heap MTLHeap) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeap:"), heap)
}

// Configures the size of a block of threadgroup memory.
//
// length: The size of the threadgroup memory, in bytes, which needs to be a multiple
// of `16` bytes.
//
// index: The index in the threadgroup memory argument table using this allocation.
//
// # Discussion
//
// The `threadgroup` memory space allows for sharing data between multiple
// threads in a threadgroup, which can be faster than using `device` memory in
// your kernels. Before using any threadgroup memory, call this method to
// configure the threadgroup memory argument table. Kernels accessing their
// arguments from threadgroup memory have the `[[threadgroup]]` attribute.
//
// To learn more about using the threadgroup address space, see the [Metal
// Shading Language Specification] section 4.4.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setThreadgroupMemoryLength(_:index:)
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/metal-shading-language-specification.pdf#//apple_ref/doc/uid/TP40014364-CH4-SW5
func (o MTLComputeCommandEncoderObject) SetThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
}

// Sets the size, in pixels, of imageblock data in tile memory.
//
// width: The width of the imageblock, in pixels.
//
// height: The height of the imageblock, in pixels.
//
// # Discussion
//
// Both imageblocks and threadgroup memory share the available space you can
// reserve in tile memory, so the sum of these allocations can’t exceed the
// maximum total tile memory limit. To find the amount of memory used by an
// imageblock, call [ImageblockMemoryLengthForDimensions]. Kernels accessing
// an imageblock argument from threadgroup memory have the
// `[[threadgroup_imageblock]]` attribute.
//
// To learn more about using imageblocks, see the following sections in the
// [Metal Shading Language Specification]:
//
// - For information on the `threadgroup_imageblock` address space, see
// Section 4.5. - For information on the `imageblock` type, see Section 2.11.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setImageblockWidth(_:height:)
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/metal-shading-language-specification.pdf#//apple_ref/doc/uid/TP40014364-CH4-SW5
func (o MTLComputeCommandEncoderObject) SetImageblockWidthHeight(width uint, height uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setImageblockWidth:height:"), width, height)
}

// Sets the dimensions over the thread grid of how your compute kernel
// receives stage-in arguments.
//
// region: The [MTLRegion] defining how to interpret a thread’s location as a
// coordinate for stage-in data.
//
// # Discussion
//
// The region’s origin point, starting from `(0,0,0)` in the upper left of
// the bound data, determines the final index of `[[stage_in]]` data. Note
// that the total number of threads Metal launches may be larger than your
// stage-in data.
//
// To determine the index used to fetch `[[stage_in]]` data for a given
// thread, the GPU adds the values specified by the region’s origin to the
// thread position in the grid. Threads in the grid outside of the maximum
// stage-in data size have undefined behavior when accessing the stage-in
// memory region.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setStageInRegion(_:)
//
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
func (o MTLComputeCommandEncoderObject) SetStageInRegion(region MTLRegion) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegion:"), region)
}

// Sets the region of the stage-in attributes to apply to a compute kernel
// using an indirect buffer.
//
// indirectBuffer: The [MTLRegion] defining how to interpret a thread’s location as a
// coordinate for stage-in data.
//
// indirectBufferOffset: Where the data begins, in bytes, from the start of the buffer.
//
// # Discussion
//
// The region’s origin point, starting from `(0,0,0)` in the upper left of
// the bound data, determines the final index of `[[stage_in]]` data. Note
// that the total number of threads Metal launches may be larger than your
// stage-in data.
//
// To determine the index used to fetch `[[stage_in]]` data for a given
// thread, the GPU adds the values specified by the region’s origin to the
// thread position in the grid. Threads in the grid outside of the maximum
// stage-in data size have undefined behavior when accessing the stage-in
// memory region.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setStageInRegionWithIndirectBuffer(_:indirectBufferOffset:)
//
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
func (o MTLComputeCommandEncoderObject) SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegionWithIndirectBuffer:indirectBufferOffset:"), indirectBuffer, indirectBufferOffset)
}

// Encodes a compute command using an arbitrarily sized grid.
//
// threadsPerGrid: The number of threads in the grid, in each dimension.
//
// threadsPerThreadgroup: The number of threads in one threadgroup, in each dimension.
//
// # Discussion
//
// This method encodes a call that uses an arbitrary number of threads in its
// execution grid. Metal calculates the number of threadgroups needed,
// providing partial threadgroups if necessary. Prefer this method to
// [DispatchThreadgroupsThreadsPerThreadgroup] if your app requires bounds
// checking or you need extra data allocations to saturate a uniform grid.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreads(_:threadsPerThreadgroup:)
func (o MTLComputeCommandEncoderObject) DispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreads:threadsPerThreadgroup:"), threadsPerGrid, threadsPerThreadgroup)
}

// Encodes a compute dispatch command using a grid aligned to threadgroup
// boundaries.
//
// threadgroupsPerGrid: An [MTLSize] instance that represents the number of threads for each grid
// dimension.
//
// threadsPerThreadgroup: An [MTLSize] instance that represents the number of threads in a
// threadgroup.
//
// # Discussion
//
// Metal calculates the number of threads in a grid by multiplying
// `threadsPerThreadgroup` by `threadgroupsPerGrid`.
//
// If the size of your data doesn’t match the size of the grid, perform
// boundary checks in your compute function to avoid accessing data out of
// bounds. See [Calculating threadgroup and grid sizes] for an example.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreadgroups(_:threadsPerThreadgroup:)
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [Calculating threadgroup and grid sizes]: https://developer.apple.com/documentation/Metal/calculating-threadgroup-and-grid-sizes
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"), threadgroupsPerGrid, threadsPerThreadgroup)
}

// Encodes a dispatch call for a compute pass, using an indirect buffer that
// defines the size of a grid that aligns to threadgroup boundaries.
//
// indirectBuffer: An [MTLBuffer] instance providing compute parameters. Lay out the data in
// this buffer as described in the [MTLDispatchThreadgroupsIndirectArguments]
// structure.
//
// indirectBufferOffset: Where the data begins, in bytes, from the start of the buffer. This value
// needs to be a multiple of `4`.
//
// threadsPerThreadgroup: The number of threads in one threadgroup, in each dimension.
//
// # Discussion
//
// The GPU fetches parameters from the indirect buffer just before the thread
// grid starts. This process lets the compute function run based on GPU
// feedback, without latency from data transfer between the CPU and the GPU.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/dispatchThreadgroups(indirectBuffer:indirectBufferOffset:threadsPerThreadgroup:)
//
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
func (o MTLComputeCommandEncoderObject) DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer MTLBuffer, indirectBufferOffset uint, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:"), indirectBuffer, indirectBufferOffset, threadsPerThreadgroup)
}

// Encodes a command that instructs the GPU to pause the compute pass until
// another pass updates a fence.
//
// fence: A fence that the pass waits for before it runs any of its commands.
//
// # Discussion
//
// You can synchronize memory operations of a compute pass that access
// resources with an [MTLFence]. This method instructs the GPU to wait until
// another pass updates `fence` before running the compute pass. The fence
// indicates when the pass can access those resources without a race
// condition.
//
// For more information about synchronization with fences, see:
//
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
//
// # Reuse a fence by waiting first and updating second
//
// When encoding a compute pass that reuses a fence, wait for other passes to
// update the fence before repurposing that fence to notify subsequent passes
// with an update:
//
// - Call the [WaitForFence] method before encoding commands that need to wait
// for other passes. - Call the [UpdateFence] method after encoding commands
// that later passes depend on.
//
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/waitForFence(_:)
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
func (o MTLComputeCommandEncoderObject) WaitForFence(fence MTLFence) {
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
}

// Encodes a command that instructs the GPU to update a fence after the
// compute pass completes.
//
// fence: A fence the pass updates after it completes.
//
// # Discussion
//
// You can synchronize memory operations of a compute pass that access
// resources with an [MTLFence]. This method instructs the pass to update
// `fence` after it runs all its memory store operations to the resources it
// accesses. The fence indicates when other passes can access those resources
// without a race condition.
//
// For more information about synchronization with fences, see:
//
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
//
// # Reuse a fence by waiting first and updating second
//
// When encoding a compute pass that reuses a fence, wait for other passes to
// update the fence before repurposing that fence to notify subsequent passes
// with an update:
//
// - Call the [WaitForFence] method before encoding commands that need to wait
// for other passes. - Call the [UpdateFence] method after encoding commands
// that later passes depend on.
//
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/updateFence(_:)
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
func (o MTLComputeCommandEncoderObject) UpdateFence(fence MTLFence) {
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
}

// Creates a memory barrier that enforces the order of write and read
// operations for specific resource types.
//
// scope: An [MTLBarrierScope] instance that represents the resource types the
// barrier synchronizes operations on.
//
// # Discussion
//
// Memory barriers ensure the relevant passes finish updating resources before
// starting the stages of subsequent commands that depend on those resources.
//
// To determine whether a GPU supports memory barriers, see the [Metal feature
// set tables (PDF)].
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/memoryBarrier(scope:)
//
// [MTLBarrierScope]: https://developer.apple.com/documentation/Metal/MTLBarrierScope
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithScope(scope MTLBarrierScope) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithScope:"), scope)
}

// Encodes a command to sample hardware counters, providing performance
// information.
//
// sampleBuffer: An [MTLCounterSampleBuffer] instance that stores the GPU hardware data.
//
// sampleIndex: An index within `sampleBuffer` the command stores the data to.
//
// barrier: Whether or not the command inserts a barrier before sampling the
// counter’s data.
//
// A barrier ensures that the commands you encode before this one complete
// before the GPU samples the hardware counters, but can negatively impact
// runtime performance.
//
// Running this command without a barrier means the GPU can sample counters
// concurrently with other commands from the encoder.
//
// The `barrier` parameter for the command has no impact on sampling commands
// from other passes.
//
// # Discussion
//
// See [GPU counters and counter sample buffers], [Sampling GPU data into
// counter sample buffers], and [MTLCounter] for more information.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
//
// [GPU counters and counter sample buffers]: https://developer.apple.com/documentation/Metal/gpu-counters-and-counter-sample-buffers
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
func (o MTLComputeCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool) {
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), sampleBuffer, sampleIndex, barrier)
}

// Creates a memory barrier that enforces the order of write and read
// operations for specific resources.
//
// resources: A C array of [MTLResource] instances the barrier applies to.
//
// count: The number of resources in the array.
//
// # Discussion
//
// Memory barriers ensure the relevant passes finish updating resources before
// starting the stages of subsequent commands that depend on those resources.
//
// To determine whether a GPU supports memory barriers, see the [Metal feature
// set tables (PDF)].
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/memoryBarrierWithResources:count:
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) MemoryBarrierWithResourcesCount(resources []MTLResource, count uint) {
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithResources:count:"), objc.CArray(resources), count)
}

// Binds multiple buffers with data in stride to the buffer argument table at
// once, allowing compute kernels to access their data on the GPU.
//
// buffers: The [MTLBuffer] instances to bind to the buffer argument table.
//
// offsets: An array of offsets, each of which specifies where the data begins, in
// bytes, from the start of its corresponding buffer.
//
// strides: An array of strides, each of which specifies the number of bytes between
// the start of one element and the start of the next for its corresponding
// buffer.
//
// range: The argument table indices to bind each of the `buffers` to, in the order
// they appear.
//
// # Discussion
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// Rebinding an already bound buffer causes a Metal error.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffers:offsets:attributeStrides:withRange:
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsAttributeStridesWithRange(buffers []MTLBuffer, offsets uint, strides uint, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:attributeStrides:withRange:"), buffers, offsets, strides, range_)
}

// Binds multiple buffers to the buffer argument table at once, allowing
// compute kernels to access their data on the GPU.
//
// buffers: The [MTLBuffer] instances to bind to the buffer argument table.
//
// offsets: An array of offsets, each of which specifies where the data begins, in
// bytes, from the start of its corresponding buffer.
//
// range: The argument table indices to bind each of the `buffers` to, in the order
// they appear.
//
// # Discussion
//
// For buffers binding to an argument using the `device` address space, align
// the offset to the data type’s size. The maximum size for an offset is
// `16` bytes.
//
// For buffers in the `constant` address space, the minimum alignment depends
// on the hardware running your app. See the [Metal feature set tables (PDF)]
// for information on each Apple GPU family.
//
// Rebinding an already bound buffer causes a Metal error.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffers:offsets:withRange:
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLComputeCommandEncoderObject) SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), buffers, offsets, range_)
}

// Binds multiple intersection function tables to the buffer argument table,
// allowing you to call their functions on the GPU.
//
// intersectionFunctionTables: An array of [MTLIntersectionFunctionTable] instances to bind.
//
// range: The argument buffer table indices to bind each of the
// `intersectionFunctionTables` to, in the order they appear.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setIntersectionFunctionTables:withBufferRange:
func (o MTLComputeCommandEncoderObject) SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), intersectionFunctionTables, range_)
}

// Encodes multiple texture samplers with custom levels of detail clamping,
// allowing compute kernels to use them for sampling textures on the GPU.
//
// samplers: An array of [MTLSamplerState] instances to bind to the sampler argument
// table.
//
// lodMinClamps: An array of minimum levels of detail to use for the corresponding sampler
// in the `samplers` array.
//
// lodMaxClamps: An array of maximum levels of detail to use for the corresponding sampler
// in the `samplers` array.
//
// range: The sampler table indicies to bind each of the `samplers` to, in the order
// they appear.
//
// # Discussion
//
// Calling this method ignores the [LodMinClamp] and [LodMaxClamp] properties
// of the samplers, using the provided levels of detail instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLComputeCommandEncoderObject) SetSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
}

// Encodes multiple texture samplers, allowing compute kernels to use them for
// sampling textures on the GPU.
//
// samplers: An array of [MTLSamplerState] instances to bind to the sampler argument
// table.
//
// range: The sampler table indexes to bind each of the `samplers` to, in the order
// they appear.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setSamplerStates:withRange:
func (o MTLComputeCommandEncoderObject) SetSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:withRange:"), samplers, range_)
}

// Binds multiple textures to the texture argument table, allowing compute
// kernels to access their data on the GPU.
//
// textures: An array of [MTLTexture] instances to bind to the texture argument table.
//
// range: The texture table indices to bind each of the `textures` to, in the order
// they appear.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setTextures:withRange:
func (o MTLComputeCommandEncoderObject) SetTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextures:withRange:"), textures, range_)
}

// Binds multiple visible function tables to the buffer argument table,
// allowing you to call their functions on the GPU.
//
// visibleFunctionTables: An array of [MTLVisibleFunctionTable] instances to bind.
//
// range: The buffer argument table indices to bind each of the
// `visibleFunctionTables` to, in the order they appear.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setVisibleFunctionTables:withBufferRange:
func (o MTLComputeCommandEncoderObject) SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables []MTLVisibleFunctionTable, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), visibleFunctionTables, range_)
}

// Ensures the shaders in the render pass’s subsequent draw commands have
// access to all of the resources you allocate from multiple heaps.
//
// heaps: An array of [MTLHeap] instances, each of which contain resources used in an
// argument buffer.
//
// count: The number of heaps in the array.
//
// # Discussion
//
// You can make the resources in each of the `heaps` (available in GPU memory)
// for the remaining duration of the render pass by calling this method. Call
// the method before encoding draw calls that may access resources within the
// `heaps` through an argument buffer. The method ensures each resource is in
// a format that’s compatible with the kernels that depend on it.
//
// This method applies the [MTLResourceUsageRead] resource usage option to all
// of the resources within `heap`, except for textures. The method ignores any
// texture that has [MTLTextureUsageRenderTarget],
// [MTLTextureUsageShaderWrite], or both in its [Usage] property. For all
// other textures in `heap`, the method optimizes each texture’s memory
// layout for rendering with a sampler. However, your kernels can’t read
// from those textures by calling this method because the texture needs a
// different memory layout that’s suitable for reading.
//
// Methods that apply a usage option for resources (see Encoding Resident
// Resources) override any previous calls that apply to a resource. For
// example, you can change the usage option for a buffer allocated in `heap`
// to [MTLResourceUsageWrite] by passing it to [UseResourcesCountUsage] after
// calling this method. However, you can’t reverse the call order because
// this method resets the usage for all resources within `heap` to
// [MTLResourceUsageRead], overriding previous calls to [UseResourceUsage].
//
// This method instructs Metal to apply hazard tracking for resources you
// allocate from a heap that you create with [MTLHazardTrackingModeTracked].
// However, for untracked resources — which come from heaps you create with
// [MTLHazardTrackingModeUntracked] — you need to account for hazards by
// applying [MTLFence] or [MTLEvent] instances.
//
// Apps typically call the method for heaps that have resources in argument
// buffers for a implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useHeaps:count:
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
func (o MTLComputeCommandEncoderObject) UseHeapsCount(heaps []MTLHeap, count uint) {
	objc.Send[struct{}](o.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
}

// Ensures kernel calls that the system encodes in subsequent commands have
// access to multiple resources.
//
// resources: An array of [MTLResource] instances used in one or more argument buffers.
//
// count: The number of resources in the array.
//
// usage: All the applicable access types the compute pass’s use of these
// resources, including [MTLResourceUsageRead] and [MTLResourceUsageWrite].
// Your resource usage type applies to all resources passed to this method
// call.
//
// For applicable resources, you may be able to prevent the GPU from
// unnecessarily decompressing color attachments on some devices by setting
// `usage` to [MTLResourceUsageRead].
//
// # Discussion
//
// You can make many resources (available in GPU memory) for the remaining
// duration of the compute pass by calling this method. Call the method before
// encoding function calls that may access these `resources` through an
// argument buffer. The method ensures the resource is in a format that’s
// compatible with the kernels that depend on it.
//
// The method also informs Metal when to apply hazard tracking for a resource
// you create with [MTLHazardTrackingModeTracked]. For a resource you create
// with [MTLHazardTrackingModeUntracked], you need to apply an [MTLFence] or
// an [MTLEvent] to account for potential reading and writing hazards.
//
// You can reconfigure an individual resource’s `usage` options for
// subsequent draw calls with the [UseResourceUsage] method.
//
// Apps typically call this method for a resource in an argument buffer as a
// part of their implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/useResources:count:usage:
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
func (o MTLComputeCommandEncoderObject) UseResourcesCountUsage(resources []MTLResource, count uint, usage MTLResourceUsage) {
	objc.Send[struct{}](o.ID, objc.Sel("useResources:count:usage:"), objc.CArray(resources), count, usage)
}

// Declares that all command generation from the encoder is completed.
//
// # Discussion
//
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLComputeCommandEncoderObject) EndEncoding() {
	objc.Send[struct{}](o.ID, objc.Sel("endEncoding"))
}

// Inserts a debug string into the captured frame data.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/insertDebugSignpost(_:)
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLComputeCommandEncoderObject) InsertDebugSignpost(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
}

// Pushes a specific string onto a stack of debug group strings for the
// command encoder.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/pushDebugGroup(_:)
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLComputeCommandEncoderObject) PushDebugGroup(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
}

// Pops the latest string off of a stack of debug group strings for the
// command encoder.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/popDebugGroup()
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLComputeCommandEncoderObject) PopDebugGroup() {
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
}

// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLComputeCommandEncoderObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLComputeCommandEncoderObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Encodes a consumer barrier on work you commit to the same command queue.
//
// afterQueueStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in prior
// command encoders, and not for the current encoder.
//
// beforeStages: [MTLStages] mask that represents the stages of work that wait. This
// argument applies to work you encode in the current command encoder.
//
// # Discussion
//
// Encode a barrier that guarantees that any subsequent work you encode in the
// current command encoder that corresponds to the `beforeStages` stages
// doesn’t proceed until Metal completes all work prior to the current
// command encoder corresponding to the `afterQueueStages` stages, completes.
//
// Metal can reorder the exact point where it applies the barrier, so use this
// method for synchronizing between different passes.
//
// If you need to synchronize work within a pass that you encode with an
// instance of a subclass of [MTLCommandEncoder], use memory barriers instead.
// For subclasses of [MTL4CommandEncoder], use encoder barriers.
//
// You can specify `afterQueueStages` and `beforeStages` that contain
// [MTLStages] unrelated to the current command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/barrier(afterQueueStages:beforeStages:)
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
func (o MTLComputeCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
}

// A string that labels the command encoder.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLComputeCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
