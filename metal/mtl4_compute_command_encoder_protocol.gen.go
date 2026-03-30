// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Encodes computation dispatches, resource copying commands, and acceleration structure building commands for a single pass into a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder
type MTL4ComputeCommandEncoder interface {
	objectivec.IObject
	MTL4CommandEncoder

	// Configures this encoder with a compute pipeline state that applies to your subsequent dispatch commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setComputePipelineState(_:)
	SetComputePipelineState(state MTLComputePipelineState)

	// Sets an argument table for the compute shader stage of this pipeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setArgumentTable(_:)
	SetArgumentTable(argumentTable MTL4ArgumentTable)

	// Configures the size of a threadgroup memory buffer for a threadgroup argument in the compute shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setThreadgroupMemoryLength(_:index:)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// Specifies the size, in pixels, of imageblock data in tile memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setImageblockSize(width:height:)
	SetImageblockWidthHeight(width uint, height uint)

	// Queries a bitmask representing the shader stages on which commands currently present in this command encoder operate.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/stages()
	Stages() MTLStages

	// Encodes a compute dispatch command using an arbitrarily-sized grid.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreads(threadsPerGrid:threadsPerThreadgroup:)
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Encodes a compute dispatch command with an arbitrarily sized grid, using an indirect buffer for arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreads(indirectBuffer:)
	DispatchThreadsWithIndirectBuffer(indirectBuffer MTLGPUAddress)

	// Encodes a compute dispatch command with a grid that aligns to threadgroup boundaries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreadgroups(threadgroupsPerGrid:threadsPerThreadgroup:)
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Encodes a compute dispatch command with a grid that aligns to threadgroup boundaries, using an indirect buffer for arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreadgroups(indirectBuffer:threadsPerThreadgroup:)
	DispatchThreadgroupsWithIndirectBufferThreadsPerThreadgroup(indirectBuffer MTLGPUAddress, threadsPerThreadgroup MTLSize)

	// Encodes a command that copies data from a buffer instance into another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceBuffer:sourceOffset:destinationBuffer:destinationOffset:size:)
	CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer MTLBuffer, sourceOffset uint, destinationBuffer MTLBuffer, destinationOffset uint, size uint)

	// Encodes a command to copy data from a tensor instance into another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTensor:sourceOrigin:sourceDimensions:destinationTensor:destinationOrigin:destinationDimensions:)
	CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor MTLTensor, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor MTLTensor, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents)

	// Encodes a command that copies data from a texture to another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:destinationTexture:)
	CopyFromTextureToTexture(sourceTexture MTLTexture, destinationTexture MTLTexture)

	// Encodes a command that copies slices of a texture to slices of another texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:sourceSlice:sourceLevel:destinationTexture:destinationSlice:destinationLevel:sliceCount:levelCount:)
	CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint)

	// Encodes a command that copies image data from a slice of a texture into a slice of another texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:destinationTexture:destinationSlice:destinationLevel:destinationOrigin:)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin)

	// Encodes a command that generates mipmaps for a texture instance from the base mipmap level up to the highest mipmap level.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/generateMipmaps(texture:)
	GenerateMipmapsForTexture(texture MTLTexture)

	// Encodes a command that modifies the contents of a texture to improve the performance of CPU accesses to its contents.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forCPUAccess:)
	OptimizeContentsForCPUAccess(texture MTLTexture)

	// Encodes a command that modifies the contents of a texture to improve the performance of CPU accesses to its contents in a specific region.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forCPUAccess:slice:level:)
	OptimizeContentsForCPUAccessSliceLevel(texture MTLTexture, slice uint, level uint)

	// Encodes a command that modifies the contents of a texture to improve the performance of GPU accesses to its contents.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forGPUAccess:)
	OptimizeContentsForGPUAccess(texture MTLTexture)

	// Encodes a command that modifies the contents of a texture instance to improve the performance of GPU accesses to its contents in a specific region.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forGPUAccess:slice:level:)
	OptimizeContentsForGPUAccessSliceLevel(texture MTLTexture, slice uint, level uint)

	// Encodes an acceleration structure build into the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/build(destinationAccelerationStructure:descriptor:scratchBuffer:)
	BuildAccelerationStructureDescriptorScratchBuffer(accelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, scratchBuffer MTL4BufferRange)

	// Encodes an acceleration structure copy operation into the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceAccelerationStructure:destinationAccelerationStructure:)
	CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure)

	// Encodes a command to copy and compact an acceleration structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copyAndCompact(sourceAccelerationStructure:destinationAccelerationStructure:)
	CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure)

	// Encodes a command to compute the size an acceleration structure can compact into, writing the result into a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/writeCompactedSize(sourceAccelerationStructure:destinationBuffer:)
	WriteCompactedAccelerationStructureSizeToBuffer(accelerationStructure MTLAccelerationStructure, buffer MTL4BufferRange)

	// Encodes an acceleration structure refit operation into the command buffer, providing additional options.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:options:)
	RefitAccelerationStructureDescriptorDestinationScratchBufferOptions(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTL4BufferRange, options MTLAccelerationStructureRefitOptions)

	// Encodes an instruction to execute commands from an indirect command buffer, using an indirect buffer for arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/executeCommands(buffer:indirectBuffer:)
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandbuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLGPUAddress)

	// Writes a GPU timestamp into a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/writeTimestamp(granularity:counterHeap:index:)
	WriteTimestampWithGranularityIntoHeapAtIndex(granularity MTL4TimestampGranularity, counterHeap MTL4CounterHeap, index uint)

	// Encode a command to attempt to improve the performance of a range of commands within an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeIndirectCommandBuffer:withRange:
	OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, range_ foundation.NSRange)

	// Encodes an acceleration structure refit into the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/refitAccelerationStructure:descriptor:destination:scratchBuffer:
	RefitAccelerationStructureDescriptorDestinationScratchBuffer(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTL4BufferRange)

	// Encodes a command that resets a range of commands in an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/resetCommandsInBuffer:withRange:
	ResetCommandsInBufferWithRange(buffer MTLIndirectCommandBuffer, range_ foundation.NSRange)
}

// MTL4ComputeCommandEncoderObject wraps an existing Objective-C object that conforms to the MTL4ComputeCommandEncoder protocol.
type MTL4ComputeCommandEncoderObject struct {
	objectivec.Object
}

func (o MTL4ComputeCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4ComputeCommandEncoderObjectFromID constructs a [MTL4ComputeCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4ComputeCommandEncoderObjectFromID(id objc.ID) MTL4ComputeCommandEncoderObject {
	return MTL4ComputeCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Configures this encoder with a compute pipeline state that applies to your
// subsequent dispatch commands.
//
// state: A non-`nil` [MTLComputePipelineState].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setComputePipelineState(_:)
func (o MTL4ComputeCommandEncoderObject) SetComputePipelineState(state MTLComputePipelineState) {
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:"), state)
}

// Sets an argument table for the compute shader stage of this pipeline.
//
// argumentTable: A [MTL4ArgumentTable] to set on the command encoder.
//
// # Discussion
//
// Metal takes a snapshot of the resources in the argument table when you make
// dispatch or execute calls on this encoder instance. Metal makes the
// snapshot contents available to the compute shader function of the current
// pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setArgumentTable(_:)
func (o MTL4ComputeCommandEncoderObject) SetArgumentTable(argumentTable MTL4ArgumentTable) {
	objc.Send[struct{}](o.ID, objc.Sel("setArgumentTable:"), argumentTable)
}

// Configures the size of a threadgroup memory buffer for a threadgroup
// argument in the compute shader function.
//
// length: The size of the threadgroup memory, in bytes. Use a multiple of `16` bytes.
//
// index: An integer that corresponds to the index of the argument you annotate with
// attribute `[[threadgroup(index)]]` in the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setThreadgroupMemoryLength(_:index:)
func (o MTL4ComputeCommandEncoderObject) SetThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
}

// Specifies the size, in pixels, of imageblock data in tile memory.
//
// width: The width of the imageblock, in pixels.
//
// height: The height of the imageblock, in pixels.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/setImageblockSize(width:height:)
func (o MTL4ComputeCommandEncoderObject) SetImageblockWidthHeight(width uint, height uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setImageblockWidth:height:"), width, height)
}

// Queries a bitmask representing the shader stages on which commands
// currently present in this command encoder operate.
//
// # Return Value
//
// A bitmask representing shader stages that commands currently present in
// this command encoder operate on.
//
// # Discussion
//
// Metal dynamically updates this property based on the commands you encode
// into the command encoder, for example, it sets the bit [MTLStageDispatch]
// if this encoder contains any commands that dispatch a compute kernel.
//
// Similarly, it sets the bit [MTLStageBlit] if this encoder contains any
// commands to copy or modify buffers, textures, or indirect command buffers.
//
// Finally, Metal sets the bit [MTLStageAccelerationStructure] if this encoder
// contains any commands that build, copy, or refit acceleration structures.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/stages()
func (o MTL4ComputeCommandEncoderObject) Stages() MTLStages {
	rv := objc.Send[MTLStages](o.ID, objc.Sel("stages"))
	return rv
}

// Encodes a compute dispatch command using an arbitrarily-sized grid.
//
// threadsPerGrid: An [MTLSize] instance that represents the number of threads in the grid, in
// each dimension.
//
// threadsPerThreadgroup: An [MTLSize] instance that represents the number of threads in one
// threadgroup, in each dimension.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreads(threadsPerGrid:threadsPerThreadgroup:)
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
func (o MTL4ComputeCommandEncoderObject) DispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreads:threadsPerThreadgroup:"), threadsPerGrid, threadsPerThreadgroup)
}

// Encodes a compute dispatch command with an arbitrarily sized grid, using an
// indirect buffer for arguments.
//
// indirectBuffer: GPUAddress of a [MTLBuffer] instance providing arguments. Lay out the data
// in this buffer as described in the [MTLDispatchThreadsIndirectArguments]
// structure. This address requires 4-byte alignment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreads(indirectBuffer:)
//
// [MTLDispatchThreadsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadsIndirectArguments
func (o MTL4ComputeCommandEncoderObject) DispatchThreadsWithIndirectBuffer(indirectBuffer MTLGPUAddress) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadsWithIndirectBuffer:"), indirectBuffer)
}

// Encodes a compute dispatch command with a grid that aligns to threadgroup
// boundaries.
//
// threadgroupsPerGrid: An [MTLSize] instance that represents the number of threadgroups in the
// grid, in each dimension.
//
// threadsPerThreadgroup: An [MTLSize] instance that represents the number of threads in one
// threadgroup, in each dimension.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreadgroups(threadgroupsPerGrid:threadsPerThreadgroup:)
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
func (o MTL4ComputeCommandEncoderObject) DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"), threadgroupsPerGrid, threadsPerThreadgroup)
}

// Encodes a compute dispatch command with a grid that aligns to threadgroup
// boundaries, using an indirect buffer for arguments.
//
// indirectBuffer: GPUAddress of a [MTLBuffer] instance providing compute parameters. Lay out
// the data in this buffer as described in the
// [MTLDispatchThreadgroupsIndirectArguments] structure. This address requires
// 4-byte alignment.
//
// threadsPerThreadgroup: A [MTLSize] instance that represents the number of threads in one
// threadgroup, in each dimension.
//
// # Discussion
//
// This method allows you to supply the threadgroups-per-grid counts
// indirectly via an [MTLBuffer] index. This enables you to calculate this
// value in the GPU timeline from a shader function, enabling GPU-driven
// workflows.
//
// Metal assumes that the buffer contents correspond to the layout of struct
// [MTLDispatchThreadgroupsIndirectArguments]. You are responsible for
// ensuring this address aligns to 4-bytes.
//
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectBuffer` parameter references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/dispatchThreadgroups(indirectBuffer:threadsPerThreadgroup:)
//
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
func (o MTL4ComputeCommandEncoderObject) DispatchThreadgroupsWithIndirectBufferThreadsPerThreadgroup(indirectBuffer MTLGPUAddress, threadsPerThreadgroup MTLSize) {
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadgroupsWithIndirectBuffer:threadsPerThreadgroup:"), indirectBuffer, threadsPerThreadgroup)
}

// Encodes a command that copies data from a buffer instance into another.
//
// sourceBuffer: An [MTLBuffer] instance the command copies data from.
//
// sourceOffset: A byte offset within `sourceBuffer` the command copies from.
//
// destinationBuffer: An [MTLBuffer] instance the command copies data to.
//
// destinationOffset: A byte offset within `destinationBuffer` the command copies to.
//
// size: The number of bytes the command copies from `sourceBuffer` to
// `destinationBuffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceBuffer:sourceOffset:destinationBuffer:destinationOffset:size:)
func (o MTL4ComputeCommandEncoderObject) CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer MTLBuffer, sourceOffset uint, destinationBuffer MTLBuffer, destinationOffset uint, size uint) {
	objc.Send[struct{}](o.ID, objc.Sel("copyFromBuffer:sourceOffset:toBuffer:destinationOffset:size:"), sourceBuffer, sourceOffset, destinationBuffer, destinationOffset, size)
}

// Encodes a command to copy data from a tensor instance into another.
//
// sourceTensor: An [MTLTensor] instance the command copies data from.
//
// destinationTensor: An [MTLTensor] instance the command copies data to.
//
// # Discussion
//
// If the `sourceTensor` and `destinationTensor` instances are not aliasable,
// this command applies the correct reshapes to enable this operation.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTensor:sourceOrigin:sourceDimensions:destinationTensor:destinationOrigin:destinationDimensions:)
func (o MTL4ComputeCommandEncoderObject) CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor MTLTensor, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor MTLTensor, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents) {
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTensor:sourceOrigin:sourceDimensions:toTensor:destinationOrigin:destinationDimensions:"), sourceTensor, sourceOrigin, sourceDimensions, destinationTensor, destinationOrigin, destinationDimensions)
}

// Encodes a command that copies data from a texture to another.
//
// sourceTexture: An [MTLTexture] instance the command copies data from.
//
// destinationTexture: Another [MTLTexture] instance the command copies the data into that has the
// same [PixelFormat] and [SampleCount] as `sourceTexture`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:destinationTexture:)
func (o MTL4ComputeCommandEncoderObject) CopyFromTextureToTexture(sourceTexture MTLTexture, destinationTexture MTLTexture) {
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:toTexture:"), sourceTexture, destinationTexture)
}

// Encodes a command that copies slices of a texture to slices of another
// texture.
//
// sourceTexture: A [MTLTexture] texture that the command copies data from. To read the
// source texture contents, you need to set its [IsFramebufferOnly] property
// to false prior to drawing into it.
//
// sourceSlice: A slice within `sourceTexture` the command uses as a starting point to copy
// data from. Set this to `0` if `sourceTexture` isn’t a texture array or a
// cube texture.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// destinationTexture: Another [MTLTexture] the command copies the data to that has the same
// [PixelFormat] and [SampleCount] as `sourceTexture`. To write the contents
// into this texture, you need to set its [IsFramebufferOnly] property to
// false.
//
// destinationSlice: A slice within `destinationTexture` the command uses as its starting point
// for copying data to. Set this to `0` if `destinationTexture` isn’t a
// texture array or a cube texture.
//
// destinationLevel: A mipmap level within `destinationTexture`. The mipmap level you reference
// needs to have the same size as the `sourceTexture` slice’s mipmap at
// `sourceLevel`.
//
// sliceCount: The number of slices the command copies so that it satisfies the conditions
// that the sum of `sourceSlice` and `sliceCount` doesn’t exceed the number
// of slices in `sourceTexture` and the sum of `destinationSlice` and
// `sliceCount` doesn’t exceed the number of slices in `destinationTexture`.
//
// levelCount: The number of mipmap levels the command copies so that it satisfies the
// conditions that the sum of `sourceLevel` and `levelCount` doesn’t exceed
// the number of mipmap levels in `sourceTexture` and the sum of
// `destinationLevel` and `levelCount` doesn’t exceed the number of mipmap
// levels in `destinationTexture`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:sourceSlice:sourceLevel:destinationTexture:destinationSlice:destinationLevel:sliceCount:levelCount:)
func (o MTL4ComputeCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint) {
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:toTexture:destinationSlice:destinationLevel:sliceCount:levelCount:"), sourceTexture, sourceSlice, sourceLevel, destinationTexture, destinationSlice, destinationLevel, sliceCount, levelCount)
}

// Encodes a command that copies image data from a slice of a texture into a
// slice of another texture.
//
// sourceTexture: An [MTLTexture] texture that the command copies data from. To read the
// source texture contents, you need to set its [IsFramebufferOnly] property
// to false prior to drawing into it.
//
// sourceSlice: A slice within `sourceTexture` the command uses as a starting point to copy
// data from. Set this to `0` if `sourceTexture` isn’t a texture array or a
// cube texture.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// sourceOrigin: An [MTLOrigin] instance that represents a location within `sourceTexture`
// that the command begins copying data from. Assign `0` to each dimension
// that’s not relevant to `sourceTexture`.
//
// sourceSize: An [MTLSize] instance that represents the size of the region, in pixels,
// that the command copies from `sourceTexture`, starting at `sourceOrigin`.
// Assign `1` to each dimension that’s not relevant to `sourceTexture`. If
// sourceTexture uses a compressed pixel format, set `sourceSize` to a
// multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
//
// destinationTexture: Another [MTLTexture] the command copies the data to that has the same
// [PixelFormat] and [SampleCount] as `sourceTexture`. To write the contents
// into this texture, you need to set its [IsFramebufferOnly] property to
// false.
//
// destinationSlice: A slice within `destinationTexture` the command uses as its starting point
// for copying data to. Set this to `0` if `destinationTexture` isn’t a
// texture array or a cube texture.
//
// destinationLevel: A mipmap level within `destinationTexture`. The mipmap level you reference
// needs to have the same size as the `sourceTexture` slice’s mipmap at
// `sourceLevel`.
//
// destinationOrigin: An [MTLOrigin] instance that represents a location within
// `destinationTexture` that the command begins copying data to. Assign `0` to
// each dimension that’s not relevant to `destinationTexture`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:destinationTexture:destinationSlice:destinationLevel:destinationOrigin:)
//
// [MTLOrigin]: https://developer.apple.com/documentation/Metal/MTLOrigin
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// [MTLOrigin]: https://developer.apple.com/documentation/Metal/MTLOrigin
func (o MTL4ComputeCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin) {
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:"), sourceTexture, sourceSlice, sourceLevel, sourceOrigin, sourceSize, destinationTexture, destinationSlice, destinationLevel, destinationOrigin)
}

// Encodes a command that generates mipmaps for a texture instance from the
// base mipmap level up to the highest mipmap level.
//
// texture: A mipmapped, color-renderable or color-filterable [MTLTexture] instance the
// command generates mipmaps for.
//
// # Discussion
//
// This method generates mipmaps for a mipmapped texture. The texture you
// provide needs to have a [MipmapLevelCount] greater than `1`, and a
// color-renderable or color-filterable [PixelFormat].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/generateMipmaps(texture:)
func (o MTL4ComputeCommandEncoderObject) GenerateMipmapsForTexture(texture MTLTexture) {
	objc.Send[struct{}](o.ID, objc.Sel("generateMipmapsForTexture:"), texture)
}

// Encodes a command that modifies the contents of a texture to improve the
// performance of CPU accesses to its contents.
//
// texture: A [MTLTexture] instance the command optimizes for CPU access.
//
// # Discussion
//
// Optimizing a texture for CPU access may affect the performance of GPU
// accesses, however, the data the GPU retrieves from the texture remains
// consistent.
//
// You typically use this command for:
//
// - Textures the CPU accesses for an extended period of time. - Textures with
// a [StorageMode] property that’s [MTLStorageModeShared] or
// [MTLStorageModeManaged].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forCPUAccess:)
func (o MTL4ComputeCommandEncoderObject) OptimizeContentsForCPUAccess(texture MTLTexture) {
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForCPUAccess:"), texture)
}

// Encodes a command that modifies the contents of a texture to improve the
// performance of CPU accesses to its contents in a specific region.
//
// texture: A [MTLTexture] the command optimizes for CPU access.
//
// slice: A slice within `texture`.
//
// level: A mipmap level within `texture`.
//
// # Discussion
//
// Optimizing a texture for CPU access may affect the performance of GPU
// accesses, however, the data the GPU retrieves from the texture remains
// consistent.
//
// You typically use this command for:
//
// - Textures the CPU accesses for an extended period of time. - Textures with
// a [StorageMode] property that’s [MTLStorageModeShared] or
// [MTLStorageModeManaged].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forCPUAccess:slice:level:)
func (o MTL4ComputeCommandEncoderObject) OptimizeContentsForCPUAccessSliceLevel(texture MTLTexture, slice uint, level uint) {
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForCPUAccess:slice:level:"), texture, slice, level)
}

// Encodes a command that modifies the contents of a texture to improve the
// performance of GPU accesses to its contents.
//
// texture: A [MTLTexture] instance the command optimizes for GPU access.
//
// # Discussion
//
// Optimizing a texture for GPU access may affect the performance of CPU
// accesses, however, the data the CPU retrieves from the texture remains
// consistent.
//
// You typically run this command for:
//
// - Textures the GPU accesses for an extended period of time. - Textures with
// a [StorageMode] property that’s [MTLStorageModeShared] or
// [MTLStorageModeManaged].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forGPUAccess:)
func (o MTL4ComputeCommandEncoderObject) OptimizeContentsForGPUAccess(texture MTLTexture) {
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForGPUAccess:"), texture)
}

// Encodes a command that modifies the contents of a texture instance to
// improve the performance of GPU accesses to its contents in a specific
// region.
//
// texture: A [MTLTexture] the command optimizes for GPU access.
//
// slice: A slice within `texture`.
//
// level: A mipmap level within `texture`.
//
// # Discussion
//
// Optimizing a texture for GPU access may affect the performance of CPU
// accesses, however, the data the CPU retrieves from the texture remains
// consistent.
//
// You typically run this command for:
//
// - Textures the GPU accesses for an extended period of time. - Textures with
// a [StorageMode] property that’s [MTLStorageModeShared] or
// [MTLStorageModeManaged].
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeContents(forGPUAccess:slice:level:)
func (o MTL4ComputeCommandEncoderObject) OptimizeContentsForGPUAccessSliceLevel(texture MTLTexture, slice uint, level uint) {
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForGPUAccess:slice:level:"), texture, slice, level)
}

// Encodes an acceleration structure build into the command buffer.
//
// accelerationStructure: Acceleration structure storage to build into.
//
// descriptor: A descriptor for the acceleration structure Metal builds.
//
// scratchBuffer: Scratch buffer Metal can use while building the acceleration structure.
// Metal may overwrite the contents of this buffer, and you should consider
// them “undefined” after the refit operation starts and completes.
//
// # Discussion
//
// Before you build an instance acceleration structure, you are responsible
// for ensuring the build operations for all primitive acceleration structures
// is complete. The built acceleration structure doesn’t retain any
// references to the input buffers of the descriptor, such as the vertex
// buffer or instance buffer, among others.
//
// The acceleration structure build process may continue as long as the
// command buffer is not completed. However, you can safely encode ray tracing
// work against the acceleration structure if you schedule and synchronize the
// command buffers that contain this ray tracing work such that the command
// buffer with the build command is complete by the time ray tracing starts.
//
// You are responsible for ensuring that the acceleration structure and
// scratch buffer are at least the size that the query
// [accelerationStructureSizes(descriptor:)] returns.
//
// Use an instance of [MTLResidencySet] to mark residency of the scratch
// buffer the `scratchBuffer` parameter references, as well as for all the
// primitive acceleration structures you directly and indirectly reference.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/build(destinationAccelerationStructure:descriptor:scratchBuffer:)
//
// [accelerationStructureSizes(descriptor:)]: https://developer.apple.com/documentation/Metal/MTLDevice/accelerationStructureSizes(descriptor:)
func (o MTL4ComputeCommandEncoderObject) BuildAccelerationStructureDescriptorScratchBuffer(accelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, scratchBuffer MTL4BufferRange) {
	objc.Send[struct{}](o.ID, objc.Sel("buildAccelerationStructure:descriptor:scratchBuffer:"), accelerationStructure, descriptor, scratchBuffer)
}

// Encodes an acceleration structure copy operation into the command buffer.
//
// sourceAccelerationStructure: Acceleration structure to copy from.
//
// destinationAccelerationStructure: Acceleration structure to copy to.
//
// # Discussion
//
// You are responsible for ensuring the source and destination acceleration
// structures don’t overlap in memory. If this is an instance acceleration
// structure, Metal preserves references to the primitive acceleration
// structures it references.
//
// Typically, the destination acceleration structure is at least as large as
// the source acceleration structure, except in cases where you compact the
// source acceleration structure. In this case, you need to allocate the
// destination acceleration to be at least as large as the compacted size of
// the source acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copy(sourceAccelerationStructure:destinationAccelerationStructure:)
func (o MTL4ComputeCommandEncoderObject) CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure) {
	objc.Send[struct{}](o.ID, objc.Sel("copyAccelerationStructure:toAccelerationStructure:"), sourceAccelerationStructure, destinationAccelerationStructure)
}

// Encodes a command to copy and compact an acceleration structure.
//
// sourceAccelerationStructure: Acceleration structure to copy and compact.
//
// destinationAccelerationStructure: Acceleration structure to copy to.
//
// # Discussion
//
// You are responsible for ensuring that the source and destination
// acceleration structures don’t overlap in memory. If this is an instance
// acceleration structure, Metal preserves references to primitive
// acceleration structures it references.
//
// This operation requires that the destination acceleration structure is at
// least as large as the compacted size of the source acceleration structure.
// You can compute this size by calling the
// [WriteCompactedAccelerationStructureSizeToBuffer] method.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/copyAndCompact(sourceAccelerationStructure:destinationAccelerationStructure:)
func (o MTL4ComputeCommandEncoderObject) CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure) {
	objc.Send[struct{}](o.ID, objc.Sel("copyAndCompactAccelerationStructure:toAccelerationStructure:"), sourceAccelerationStructure, destinationAccelerationStructure)
}

// Encodes a command to compute the size an acceleration structure can compact
// into, writing the result into a buffer.
//
// accelerationStructure: Source acceleration structure.
//
// buffer: Destination size buffer. Metal writes the compacted size as a 64-bit
// unsigned integer value, representing the compacted size in bytes.
//
// # Discussion
//
// This size is potentially smaller than the acceleration structure. To
// perform compaction, you typically read this size from the buffer once the
// command buffer completes. You then use it to allocate a new, potentially
// smaller acceleration structure. Finally, you call the
// [CopyAndCompactAccelerationStructureToAccelerationStructure] method to
// perform the copy.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/writeCompactedSize(sourceAccelerationStructure:destinationBuffer:)
func (o MTL4ComputeCommandEncoderObject) WriteCompactedAccelerationStructureSizeToBuffer(accelerationStructure MTLAccelerationStructure, buffer MTL4BufferRange) {
	objc.Send[struct{}](o.ID, objc.Sel("writeCompactedAccelerationStructureSize:toBuffer:"), accelerationStructure, buffer)
}

// Encodes an acceleration structure refit operation into the command buffer,
// providing additional options.
//
// sourceAccelerationStructure: Acceleration structure to refit.
//
// descriptor: A descriptor for the acceleration structure to refit.
//
// destinationAccelerationStructure: Acceleration structure to store the refit result into. If
// `destinationAccelerationStructure` is `nil`, Metal performs an in-place
// refit operation of the `sourceAccelerationStructure`.
//
// scratchBuffer: Scratch buffer Metal can use while refitting the acceleration structure.
// Metal may overwrite the contents of this buffer, and you should consider
// them “undefined” after the refit operation starts and completes.
//
// options: Options specifying the elements of the acceleration structure to refit.
//
// # Discussion
//
// You refit an acceleration structure to update it when the geometry it
// references changes. This operation is typically much faster than rebuilding
// the acceleration structure from scratch. The trade off is that after you
// refit the acceleration structure, its quality, as well as the performance
// of any subsequent ray tracing operation degrades, depending on how much the
// geometry changes.
//
// After certain operations, refitting an acceleration structure may not be
// possible, for example, after adding or removing geometry.
//
// When you refit an acceleration structure, you can do so in place, by
// specifying the same source and destination acceleration structures, or by
// providing a `nil` destination acceleration structure. If the source and
// destination acceleration structures aren’t the same, then you are
// responsible for ensuring they don’t overlap in memory.
//
// Typically, the destination acceleration structure is at least as large as
// the source acceleration structure, except in cases where you compact the
// source acceleration structure. In this case, you need to allocate the
// destination acceleration to be at least as large as the compacted size of
// the source acceleration structure.
//
// The scratch buffer you provide for the refit operation needs to be at least
// as large as the size that the query
// [accelerationStructureSizes(descriptor:)] returns. If the size this query
// returns is zero, you can omit providing a scratch buffer by passing `0` as
// the address to the `scratchBuffer` parameter.
//
// Use an instance of [MTLResidencySet] to mark residency of the scratch
// buffer the `scratchBuffer` parameter references, as well as for all the
// instance and primitive acceleration structures you directly and indirectly
// reference.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:options:)
//
// [accelerationStructureSizes(descriptor:)]: https://developer.apple.com/documentation/Metal/MTLDevice/accelerationStructureSizes(descriptor:)
func (o MTL4ComputeCommandEncoderObject) RefitAccelerationStructureDescriptorDestinationScratchBufferOptions(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTL4BufferRange, options MTLAccelerationStructureRefitOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:options:"), sourceAccelerationStructure, descriptor, destinationAccelerationStructure, scratchBuffer, options)
}

// Encodes an instruction to execute commands from an indirect command buffer,
// using an indirect buffer for arguments.
//
// indirectCommandbuffer: [MTLIndirectCommandBuffer] instance containing the commands to execute.
//
// indirectRangeBuffer: GPUAddress of a [MTLBuffer] containing the execution range. Lay out the
// data in this buffer as described in the
// [MTLIndirectCommandBufferExecutionRange] structure. This address requires
// 4-byte alignment.
//
// # Discussion
//
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectRangeBuffer` parameter references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/executeCommands(buffer:indirectBuffer:)
//
// [MTLIndirectCommandBufferExecutionRange]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
func (o MTL4ComputeCommandEncoderObject) ExecuteCommandsInBufferIndirectBuffer(indirectCommandbuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLGPUAddress) {
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:"), indirectCommandbuffer, indirectRangeBuffer)
}

// Writes a GPU timestamp into a heap.
//
// granularity: [MTL4TimestampGranularity] hint to Metal about acceptable the level of
// precision.
//
// counterHeap: [MTL4CounterHeap] to write timestamps into.
//
// index: The index value into which Metal writes the timestamp.
//
// # Discussion
//
// The method ensures that any prior work finishes, but doesn’t delay any
// subsequent work.
//
// You can alter this command’s behavior through the `granularity`
// parameter.
//
// - Pass [MTL4TimestampGranularityRelaxed] to allow Metal to provide
// timestamps with minimal impact to runtime performance, but with less
// detail. For example, the command may group all timestamps for a pass
// together. - Pass [MTL4TimestampGranularityPrecise] to request that Metal
// provides timestamps with the most detail. This can affect runtime
// performance.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/writeTimestamp(granularity:counterHeap:index:)
//
// [MTL4TimestampGranularity]: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity
func (o MTL4ComputeCommandEncoderObject) WriteTimestampWithGranularityIntoHeapAtIndex(granularity MTL4TimestampGranularity, counterHeap MTL4CounterHeap, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("writeTimestampWithGranularity:intoHeap:atIndex:"), granularity, counterHeap, index)
}

// Encode a command to attempt to improve the performance of a range of
// commands within an indirect command buffer.
//
// indirectCommandBuffer: An [MTLIndirectCommandBuffer] instance that this command optimizes.
//
// range: A range of commands within `indirectCommandBuffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/optimizeIndirectCommandBuffer:withRange:
func (o MTL4ComputeCommandEncoderObject) OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("optimizeIndirectCommandBuffer:withRange:"), indirectCommandBuffer, range_)
}

// Encodes an acceleration structure refit into the command buffer.
//
// sourceAccelerationStructure: Acceleration structure to refit.
//
// descriptor: A descriptor for the acceleration structure to refit.
//
// destinationAccelerationStructure: Acceleration structure to store the refit result into. If
// `destinationAccelerationStructure` is `nil`, Metal performs an in-place
// refit operation of the `sourceAccelerationStructure`.
//
// scratchBuffer: Scratch buffer Metal can use while refitting the acceleration structure.
// Metal may overwrite the contents of this buffer, and you should consider
// them “undefined” after the refit operation starts and completes.
//
// # Discussion
//
// You refit an acceleration structure to update it when the geometry it
// references changes. This operation is typically much faster than rebuilding
// the acceleration structure from scratch. The trade off is that after you
// refit the acceleration structure, its quality, as well as the performance
// of any subsequent ray tracing operation degrades, depending on how much the
// geometry changes.
//
// After certain operations, refitting an acceleration structure may not be
// possible, for example, after adding or removing geometry.
//
// When you refit an acceleration structure, you can do so in place, by
// specifying the same source and destination acceleration structures, or by
// providing a `nil` destination acceleration structure. If the source and
// destination acceleration structures aren’t the same, then you are
// responsible for ensuring they don’t overlap in memory.
//
// Typically, the destination acceleration structure is at least as large as
// the source acceleration structure, except in cases where you compact the
// source acceleration structure. In this case, you need to allocate the
// destination acceleration to be at least as large as the compacted size of
// the source acceleration structure.
//
// The scratch buffer you provide for the refit operation needs to be at least
// as large as the size that the query
// [accelerationStructureSizes(descriptor:)] returns. If the size this query
// returns is zero, you can omit providing a scratch buffer by passing `0` as
// the address to the `scratchBuffer` parameter.
//
// Use an instance of [MTLResidencySet] to mark residency of the scratch
// buffer the `scratchBuffer` parameter references, as well as for all the
// instance and primitive acceleration structures you directly and indirectly
// reference.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/refitAccelerationStructure:descriptor:destination:scratchBuffer:
//
// [accelerationStructureSizes(descriptor:)]: https://developer.apple.com/documentation/Metal/MTLDevice/accelerationStructureSizes(descriptor:)
func (o MTL4ComputeCommandEncoderObject) RefitAccelerationStructureDescriptorDestinationScratchBuffer(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTL4BufferRange) {
	objc.Send[struct{}](o.ID, objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:"), sourceAccelerationStructure, descriptor, destinationAccelerationStructure, scratchBuffer)
}

// Encodes a command that resets a range of commands in an indirect command
// buffer.
//
// buffer: An [MTLIndirectCommandBuffer] the command resets.
//
// range: A range of commands within `buffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputeCommandEncoder/resetCommandsInBuffer:withRange:
func (o MTL4ComputeCommandEncoderObject) ResetCommandsInBufferWithRange(buffer MTLIndirectCommandBuffer, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("resetCommandsInBuffer:withRange:"), buffer, range_)
}

// Returns the command buffer that is currently encoding commands.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/commandBuffer
func (o MTL4ComputeCommandEncoderObject) CommandBuffer() MTL4CommandBuffer {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBuffer"))
	return MTL4CommandBufferObjectFromID(rv)
}

// Provides an optional label to assign to the command encoder for debug
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/label
func (o MTL4ComputeCommandEncoderObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Declares that all command generation from this encoder is complete.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/endEncoding()
func (o MTL4ComputeCommandEncoderObject) EndEncoding() {
	objc.Send[struct{}](o.ID, objc.Sel("endEncoding"))
}

// Inserts a debug string into the frame data to aid debugging.
//
// string: The debug string to insert as a signpost.
//
// # Discussion
//
// Calling this method doesn’t change any behaviors, but can be useful for
// debugging purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/insertDebugSignpost(_:)
func (o MTL4ComputeCommandEncoderObject) InsertDebugSignpost(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
}

// Pops the latest debug group string from this encoder’s stack of debug
// groups.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/popDebugGroup()
func (o MTL4ComputeCommandEncoderObject) PopDebugGroup() {
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
}

// Pushes a string onto this encoder’s stack of debug groups.
//
// string: The debug string to push.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/pushDebugGroup(_:)
func (o MTL4ComputeCommandEncoderObject) PushDebugGroup(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
}

// Encodes a command that instructs the GPU to update a fence after one or
// more stages, which can unblock other passes waiting for the fence.
//
// fence: A fence the pass updates after the stages in `afterEncoderStages` complete.
//
// afterEncoderStages: The encoder stages that need to complete before the pass updates `fence`.
//
// # Discussion
//
// You can synchronize memory operations of a pass that access resources with
// an [MTLFence]. This method instructs the pass to update `fence` after the
// stages you pass to the `afterEncoderStages` run all their memory store
// operations to the resources it accesses. The fence indicates when other
// passes can access those resources without a race condition.
//
// For more information about synchronization with fences, see:
//
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
//
// # Reuse a fence by waiting first and updating second
//
// When encoding a pass that reuses a fence, wait for other passes to update
// the fence before repurposing that fence to notify subsequent passes with an
// update:
//
// - Call the [WaitForFenceBeforeEncoderStages] method before encoding
// commands that need to wait for other passes. - Call the
// [UpdateFenceAfterEncoderStages] method after encoding commands that later
// passes depend on.
//
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/updateFence(_:afterEncoderStages:)
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
func (o MTL4ComputeCommandEncoderObject) UpdateFenceAfterEncoderStages(fence MTLFence, afterEncoderStages MTLStages) {
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:afterEncoderStages:"), fence, afterEncoderStages)
}

// Encodes a command that instructs the GPU to pause before starting one or
// more stages of the pass until a pass updates a fence.
//
// fence: A fence that the pass waits for before running the stages you pass to
// `beforeEncoderStages`.
//
// beforeEncoderStages: The encoder stages that need to wait for another pass to update `fence`
// before they run.
//
// # Discussion
//
// You can synchronize memory operations of a pass that access resources with
// an [MTLFence]. This method instructs the GPU to wait until another pass
// updates `fence` before running the stages you pass to the
// `beforeEncoderStages` parameter. The fence indicates when the pass can
// access those resources without a race condition.
//
// For more information about synchronization with fences, see:
//
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
//
// # Reuse a fence by waiting first and updating second
//
// When encoding a pass that reuses a fence, wait for other passes to update
// the fence before repurposing that fence to notify subsequent passes with an
// update:
//
// - Call the [WaitForFenceBeforeEncoderStages] method before encoding
// commands that need to wait for other passes. - Call the
// [UpdateFenceAfterEncoderStages] method after encoding commands that later
// passes depend on.
//
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/waitForFence(_:beforeEncoderStages:)
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
func (o MTL4ComputeCommandEncoderObject) WaitForFenceBeforeEncoderStages(fence MTLFence, beforeEncoderStages MTLStages) {
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:beforeEncoderStages:"), fence, beforeEncoderStages)
}

// Provides an optional label to assign to the command encoder for debug
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/label
func (o MTL4ComputeCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
