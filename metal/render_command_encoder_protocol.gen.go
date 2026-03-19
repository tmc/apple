// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Encodes configuration and draw commands for a single render pass into a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder
type MTLRenderCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Encodes a draw command that renders an instance of a geometric primitive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:)
	DrawPrimitivesVertexStartVertexCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:instanceCount:)
	DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive that starts with a custom instance identification number.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:instanceCount:baseInstance:)
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:indirectBuffer:indirectBufferOffset:)
	DrawPrimitivesIndirectBufferIndirectBufferOffset(primitiveType MTLPrimitiveType, indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a draw command that renders an instance of a geometric primitive with indexed vertices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffset(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices, starting with a custom vertex and instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices and indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexType:indexBuffer:indexBufferOffset:indirectBuffer:indirectBufferOffset:)
	DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferOffsetIndirectBufferIndirectBufferOffset(primitiveType MTLPrimitiveType, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with a grid of threads.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreads(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with a grid of threadgroups.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreadgroups(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreadgroups(indirectBuffer:indirectBufferOffset:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer MTLBuffer, indirectBufferOffset uint, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a draw command that renders multiple instances of tessellated patches.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPatches(numberOfPatchControlPoints:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:)
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint)

	// Encodes a draw command that renders multiple instances of tessellated patches with indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPatches(numberOfPatchControlPoints:patchIndexBuffer:patchIndexBufferOffset:indirectBuffer:indirectBufferOffset:)
	DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a draw command that renders multiple instances of tessellated patches with a control point index buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPatches(numberOfPatchControlPoints:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:instanceCount:baseInstance:)
	DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer MTLBuffer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint)

	// Encodes a draw command that renders multiple instances of tessellated patches with a control point index buffer and indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPatches(numberOfPatchControlPoints:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:indirectBuffer:indirectBufferOffset:)
	DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer MTLBuffer, controlPointIndexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a command that invokes GPU functions from the encoder’s current tile render pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/dispatchThreadsPerTile(_:)
	DispatchThreadsPerTile(threadsPerTile MTLSize)

	// The width of the tiles, in pixels, for the render command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/tileWidth
	TileWidth() uint

	// The height of the tiles, in pixels, for the render command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/tileHeight
	TileHeight() uint

	// Encodes a command that instructs the GPU to pause before starting one or more stages of the render pass until a pass updates a fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/waitForFence(_:before:)
	WaitForFenceBeforeStages(fence MTLFence, stages MTLRenderStages)

	// Encodes a command that instructs the GPU to update a fence after one or more stages, which can unblock other passes waiting for the fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/updateFence(_:after:)
	UpdateFenceAfterStages(fence MTLFence, stages MTLRenderStages)

	// Creates a memory barrier that enforces the order of write and read operations for specific resource types.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/memoryBarrier(scope:after:before:)
	MemoryBarrierWithScopeAfterStagesBeforeStages(scope MTLBarrierScope, after MTLRenderStages, before MTLRenderStages)

	// Encodes a command that samples hardware counters during the render pass and stores the data into a counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool)

	// Encodes a command that runs an indirect range of commands from an indirect command buffer (ICB).
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:
	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(indirectCommandbuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a command that runs a range of commands from an indirect command buffer (ICB).
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/executeCommandsInBuffer:withRange:
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, executionRange foundation.NSRange)

	// Creates a memory barrier that enforces the order of write and read operations for specific resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/memoryBarrierWithResources:count:afterStages:beforeStages:
	MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources []MTLResource, count uint, after MTLRenderStages, before MTLRenderStages)

	// Configures each pixel component value, including alpha, for the render pipeline’s constant blend color.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setBlendColor(red:green:blue:alpha:)
	SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32)

	// Sets the mapping from logical shader color output to physical render pass color attachments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorAttachmentMap(_:)
	SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap)

	// Configures the store action for a color attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorStoreAction(_:index:)
	SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint)

	// Configures the store action options for a color attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorStoreActionOptions(_:index:)
	SetColorStoreActionOptionsAtIndex(storeActionOptions MTLStoreActionOptions, colorAttachmentIndex uint)

	// Configures how the render pipeline determines which primitives to remove.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setCullMode(_:)
	SetCullMode(cullMode MTLCullMode)

	// Configures the adjustments a render pass applies to depth values from fragment functions by a scaling factor and bias.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthBias(_:slopeScale:clamp:)
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)

	// Configures how the render pipeline handles fragments outside the near and far planes of the view frustum.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthClipMode(_:)
	SetDepthClipMode(depthClipMode MTLDepthClipMode)

	// Configures the combined depth and stencil state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStencilState(_:)
	SetDepthStencilState(depthStencilState MTLDepthStencilState)

	// Configures the store action for the depth attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStoreAction(_:)
	SetDepthStoreAction(storeAction MTLStoreAction)

	// Configures the store action options for the depth attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStoreActionOptions(_:)
	SetDepthStoreActionOptions(storeActionOptions MTLStoreActionOptions)

	// Configures the minimum and maximum bounds for depth bounds testing.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthTestMinBound:maxBound:
	SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32)

	// Assigns an acceleration structure to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentAccelerationStructure(_:bufferIndex:)
	SetFragmentAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint)

	// Updates an entry in the fragment shader argument table with a new location within the entry’s current buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBufferOffset(_:index:)
	SetFragmentBufferOffsetAtIndex(offset uint, index uint)

	// Assigns multiple buffers to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBuffers:offsets:withRange:
	SetFragmentBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Creates a buffer from bytes and assigns it to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBytes(_:length:index:)
	SetFragmentBytesLengthAtIndex(bytes []byte, index uint)

	// Assigns an intersection function table to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentIntersectionFunctionTable(_:bufferIndex:)
	SetFragmentIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint)

	// Assigns multiple intersection function tables to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentIntersectionFunctionTables:withBufferRange:
	SetFragmentIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange)

	// Assigns a sampler state to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerState(_:index:)
	SetFragmentSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Assigns a sampler state and clamp values to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetFragmentSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Assigns multiple sampler states and clamp values to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetFragmentSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Assigns multiple sampler states to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerStates:withRange:
	SetFragmentSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Assigns a texture to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentTexture(_:index:)
	SetFragmentTextureAtIndex(texture MTLTexture, index uint)

	// Assigns multiple textures to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentTextures:withRange:
	SetFragmentTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Assigns a visible function table to an entry in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentVisibleFunctionTable(_:bufferIndex:)
	SetFragmentVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint)

	// Assigns multiple visible function tables to a range of entries in the fragment shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentVisibleFunctionTables:withBufferRange:
	SetFragmentVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange)

	// Configures which face of a primitive, such as a triangle, is the front.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFrontFacing(_:)
	SetFrontFacingWinding(frontFacingWinding MTLWinding)

	// Updates an entry in the mesh shader argument table with a new location within the entry’s current buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBufferOffset(_:index:)
	SetMeshBufferOffsetAtIndex(offset uint, index uint)

	// Assigns multiple buffers to a range of entries in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBuffers:offsets:withRange:
	SetMeshBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Creates a buffer from bytes and assigns it to an entry in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBytes(_:length:index:)
	SetMeshBytesLengthAtIndex(bytes []byte, index uint)

	// Assigns a sampler state to an entry in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerState(_:index:)
	SetMeshSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Assigns a sampler state and clamp values to an entry in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetMeshSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Assigns multiple sampler states and clamp values to a range of entries in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetMeshSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Assigns multiple sampler states to a range of entries in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerStates:withRange:
	SetMeshSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Assigns a texture to an entry in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshTexture(_:index:)
	SetMeshTextureAtIndex(texture MTLTexture, index uint)

	// Assigns multiple textures to a range of entries in the mesh shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshTextures:withRange:
	SetMeshTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Updates an entry in the object shader argument table with a new location within the entry’s current buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBufferOffset(_:index:)
	SetObjectBufferOffsetAtIndex(offset uint, index uint)

	// Encodes a command that assigns multiple buffers to a range of entries in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBuffers:offsets:withRange:
	SetObjectBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Creates a buffer from bytes and assigns it to an entry in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBytes(_:length:index:)
	SetObjectBytesLengthAtIndex(bytes []byte, index uint)

	// Assigns a sampler state to an entry in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerState(_:index:)
	SetObjectSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Assigns a sampler state and clamp values to an entry in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetObjectSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Assigns multiple sampler states and clamp values to a range of entries in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetObjectSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Assigns multiple sampler states to a range of entries in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerStates:withRange:
	SetObjectSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Assigns a texture to an entry in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectTexture(_:index:)
	SetObjectTextureAtIndex(texture MTLTexture, index uint)

	// Assigns multiple textures to a range of entries in the object shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectTextures:withRange:
	SetObjectTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Configures the size of a threadgroup memory buffer for an entry in the object argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectThreadgroupMemoryLength(_:index:)
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// Configures the encoder with a render or tile pipeline state that applies to your subsequent draw commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setRenderPipelineState(_:)
	SetRenderPipelineState(pipelineState MTLRenderPipelineState)

	// Configures a rectangle for the fragment scissor test.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setScissorRect(_:)
	SetScissorRect(rect MTLScissorRect)

	// Configures multiple rectangles for the fragment scissor test.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setScissorRects:count:
	SetScissorRectsCount(scissorRects []MTLScissorRect, count uint)

	// Configures different comparison values for front- and back-facing primitives.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilReferenceValues(front:back:)
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32)

	// Configures the same comparison value for front- and back-facing primitives.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilReferenceValue(_:)
	SetStencilReferenceValue(referenceValue uint32)

	// Configures the store action for the stencil attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilStoreAction(_:)
	SetStencilStoreAction(storeAction MTLStoreAction)

	// Configures the store action options for the stencil attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilStoreActionOptions(_:)
	SetStencilStoreActionOptions(storeActionOptions MTLStoreActionOptions)

	// Configures the per-patch tessellation factors for any subsequent patch-drawing commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTessellationFactorBuffer(_:offset:instanceStride:)
	SetTessellationFactorBufferOffsetInstanceStride(buffer MTLBuffer, offset uint, instanceStride uint)

	// Configures the scale factor for per-patch tessellation factors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTessellationFactorScale(_:)
	SetTessellationFactorScale(scale float32)

	// Configures the size of a threadgroup memory buffer for an entry in the fragment or tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setThreadgroupMemoryLength(_:offset:index:)
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)

	// Assigns an acceleration structure to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileAccelerationStructure(_:bufferIndex:)
	SetTileAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint)

	// Updates an entry in the tile shader argument table with a new location within the entry’s current buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBufferOffset(_:index:)
	SetTileBufferOffsetAtIndex(offset uint, index uint)

	// Assigns multiple buffers to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBuffers:offsets:withRange:
	SetTileBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Creates a buffer from bytes and assigns it to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBytes(_:length:index:)
	SetTileBytesLengthAtIndex(bytes []byte, index uint)

	// Assigns an intersection function table to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileIntersectionFunctionTable(_:bufferIndex:)
	SetTileIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint)

	// Assigns multiple intersection function tables to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileIntersectionFunctionTables:withBufferRange:
	SetTileIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange)

	// Assigns a sampler state to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerState(_:index:)
	SetTileSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Assigns a sampler state and clamp values to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetTileSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Assigns multiple sampler states and clamp values to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Assigns multiple sampler states to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerStates:withRange:
	SetTileSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Assigns a texture to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileTexture(_:index:)
	SetTileTextureAtIndex(texture MTLTexture, index uint)

	// Assigns multiple textures to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileTextures:withRange:
	SetTileTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Assigns a visible function table to an entry in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileVisibleFunctionTable(_:bufferIndex:)
	SetTileVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint)

	// Assigns multiple visible function tables to a range of entries in the tile shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileVisibleFunctionTables:withBufferRange:
	SetTileVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange)

	// Configures how subsequent draw commands rasterize triangle and triangle strip primitives.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTriangleFillMode(_:)
	SetTriangleFillMode(fillMode MTLTriangleFillMode)

	// Assigns an acceleration structure to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexAccelerationStructure(_:bufferIndex:)
	SetVertexAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint)

	// Configures the number of output vertices the render pipeline produces for each input vertex, optionally with render target and viewport offsets.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexAmplificationCount(_:viewMappings:)
	SetVertexAmplificationCountViewMappings(count uint, viewMappings *MTLVertexAmplificationViewMapping)

	// Updates an entry in the vertex shader argument table with a new location within the entry’s current buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBufferOffset(_:index:)
	SetVertexBufferOffsetAtIndex(offset uint, index uint)

	// SetVertexBufferOffsetAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBufferOffset(offset:attributeStride:index:)
	SetVertexBufferOffsetAttributeStrideAtIndex(offset uint, stride uint, index uint)

	// SetVertexBuffersOffsetsAttributeStridesWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffers:offsets:attributeStrides:withRange:
	SetVertexBuffersOffsetsAttributeStridesWithRange(buffers []MTLBuffer, offsets uint, strides uint, range_ foundation.NSRange)

	// Assigns multiple buffers to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffers:offsets:withRange:
	SetVertexBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Creates a buffer from bytes and assigns it to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBytes(_:length:index:)
	SetVertexBytesLengthAtIndex(bytes []byte, index uint)

	// SetVertexBytesLengthAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBytes(_:length:attributeStride:index:)
	SetVertexBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint, index uint)

	// Assigns an intersection function table to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexIntersectionFunctionTable(_:bufferIndex:)
	SetVertexIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint)

	// Assigns multiple intersection function tables to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexIntersectionFunctionTables:withBufferRange:
	SetVertexIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange)

	// Assigns a sampler state to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerState(_:index:)
	SetVertexSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Assigns a sampler state and clamp values to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerState(_:lodMinClamp:lodMaxClamp:index:)
	SetVertexSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint)

	// Assigns multiple sampler states and clamp values to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerStates:lodMinClamps:lodMaxClamps:withRange:
	SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange)

	// Assigns multiple sampler states to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerStates:withRange:
	SetVertexSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Assigns a texture to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexTexture(_:index:)
	SetVertexTextureAtIndex(texture MTLTexture, index uint)

	// Assigns multiple textures to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexTextures:withRange:
	SetVertexTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Assigns a visible function table to an entry in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexVisibleFunctionTable(_:bufferIndex:)
	SetVertexVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint)

	// Assigns multiple visible function tables to a range of entries in the vertex shader argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexVisibleFunctionTables:withBufferRange:
	SetVertexVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange)

	// Configures the render pipeline with a viewport that applies a transformation and a clipping rectangle.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setViewport(_:)
	SetViewport(viewport MTLViewport)

	// Configures the render pipeline with multiple viewports that apply transformations and clipping rectangles.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setViewports:count:
	SetViewportsCount(viewports []MTLViewport, count uint)

	// Configures which visibility test the GPU runs and the destination for any results it generates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVisibilityResultMode(_:offset:)
	SetVisibilityResultModeOffset(mode MTLVisibilityResultMode, offset uint)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to the resources you allocate from a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useHeap(_:stages:)
	UseHeapStages(heap MTLHeap, stages MTLRenderStages)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to the resources you allocate from multiple heaps.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useHeaps:count:stages:
	UseHeapsCountStages(heaps []MTLHeap, count uint, stages MTLRenderStages)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to a resource.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useResource(_:usage:stages:)
	UseResourceUsageStages(resource MTLResource, usage MTLResourceUsage, stages MTLRenderStages)

	// Ensures the shaders in the render pass’s subsequent draw commands have access to multiple resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useResources:count:usage:stages:
	UseResourcesCountUsageStages(resources []MTLResource, count uint, usage MTLResourceUsage, stages MTLRenderStages)
}

// MTLRenderCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLRenderCommandEncoder protocol.
type MTLRenderCommandEncoderObject struct {
	objectivec.Object
}
func (o MTLRenderCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLRenderCommandEncoderObjectFromID constructs a [MTLRenderCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLRenderCommandEncoderObjectFromID(id objc.ID) MTLRenderCommandEncoderObject {
	return MTLRenderCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes a draw command that renders an instance of a geometric primitive.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// vertexStart: The lowest value the command passes to your vertex shader’s parameter
// with the `vertex_id` attribute. The command assigns each vertex a unique
// `vertex_id` value within its drawing instance that increases from
// `vertexStart` through `(vertexStart + vertexCount - 1)`. Your shader can
// use that value to identify a vertex in each drawing instance.
// 
// For more information about the `vertex_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws.
//
// # Discussion
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:)
func (o MTLRenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:"), primitiveType, vertexStart, vertexCount)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// vertexStart: The lowest value the command passes to your vertex shader’s parameter
// with the `vertex_id` attribute. The command assigns each vertex a unique
// `vertex_id` value within its drawing instance that increases from
// `vertexStart` through `(vertexStart + vertexCount - 1)`. Your shader can
// use that value to identify a vertex in each drawing instance.
// 
// For more information about the `vertex_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws per instance.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `vertexCount` vertices.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `0` through `(instanceCount - 1)`. Your shader can use that
// value to identify which instance the vertex belongs to.
// 
// For more information about the `instance_id` argument attribute for vertex
// shaders, the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Discussion
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:instanceCount:)
func (o MTLRenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:"), primitiveType, vertexStart, vertexCount, instanceCount)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive that starts with a custom instance identification number.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// vertexStart: The lowest value the command passes to your vertex shader’s parameter
// with the `vertex_id` attribute. The command assigns each vertex a unique
// `vertex_id` value within its drawing instance that increases from
// `vertexStart` through `(vertexStart + vertexCount - 1)`. Your shader can
// use that value to identify a vertex in each drawing instance.
// 
// For more information about the `vertex_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws per instance.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `vertexCount` vertices.
//
// baseInstance: The lowest value the command passes to your vertex shader’s parameter
// with the `instance_id` attribute.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `baseInstance` through `(baseInstance + instanceCount - 1)`.
// Your shader can use that value to identify which instance the vertex
// belongs to.
// 
// For more information about the `instance_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Discussion
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:vertexStart:vertexCount:instanceCount:baseInstance:)
func (o MTLRenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:baseInstance:"), primitiveType, vertexStart, vertexCount, instanceCount, baseInstance)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive with indirect arguments.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// indirectBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLDrawPrimitivesIndirectArguments] structure.
// //
// [MTLDrawPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawPrimitivesIndirectArguments
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectBuffer` where the indirect arguments structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// # Discussion
// 
// Indirect drawing methods may help your app avoid expensive latency costs.
// This is because the command reads arguments from an [MTLBuffer] instance
// instead of using the CPU to pass parameters to the command.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPrimitives(type:indirectBuffer:indirectBufferOffset:)
func (o MTLRenderCommandEncoderObject) DrawPrimitivesIndirectBufferIndirectBufferOffset(primitiveType MTLPrimitiveType, indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:indirectBuffer:indirectBufferOffset:"), primitiveType, indirectBuffer, indirectBufferOffset)
	}
// Encodes a draw command that renders an instance of a geometric primitive
// with indexed vertices.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer`.
//
// indexType: An [MTLIndexType] instance that represents the index’s format, including
// [IndexTypeUInt16] and [IndexTypeUInt32].
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: An [MTLBuffer] instance that contains the `indexCount` vertex indices of
// the `indexType` format.
//
// indexBufferOffset: An integer that represents the location that’s a multiple of the index
// size from the start of `indexBuffer` where the vertex indices begin.
//
// # Discussion
// 
// You can complete a primitive and start a new one by passing a sentinel
// index value that’s the largest unsigned integer possible for `indexType`.
// For example, the largest unsigned integer for [IndexTypeUInt16] and
// [IndexTypeUInt32] is `0xFFFF` and `0xFFFFFFFF`, respectively. The command
// finishes the current primitive and begins drawing a new one each time the
// command reads a sentinel index value.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffset(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferOffset)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer` for each instance.
//
// indexType: An [MTLIndexType] instance that represents the index’s format, including
// [IndexTypeUInt16] and [IndexTypeUInt32].
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: An [MTLBuffer] instance that contains the `indexCount` vertex indices of
// the `indexType` format.
//
// indexBufferOffset: An integer that represents the location that’s a multiple of the index
// size from the start of `indexBuffer` where the vertex indices begin.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `indexCount` vertices.
//
// # Discussion
// 
// You can complete a primitive and start a new one by passing a sentinel
// index value that’s the largest unsigned integer possible for `indexType`.
// For example, the largest unsigned integer for [IndexTypeUInt16] and
// [IndexTypeUInt32] is `0xFFFF` and `0xFFFFFFFF`, respectively. The command
// finishes the current primitive and begins drawing a new one each time the
// command reads a sentinel index value.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferOffset, instanceCount)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices, starting with a custom vertex and
// instance.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer` for each instance.
//
// indexType: An [MTLIndexType] instance that represents the index’s format, including
// [IndexTypeUInt16] and [IndexTypeUInt32].
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: An [MTLBuffer] instance that contains the `indexCount` vertex indices of
// the `indexType` format.
//
// indexBufferOffset: An integer that represents the location that’s a multiple of the index
// size from the start of `indexBuffer` where the vertex indices begin.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `indexCount` vertices.
//
// baseVertex: The lowest value the command passes to your vertex shader’s parameter
// with the `vertex_id` attribute. The command assigns each vertex a unique
// `vertex_id` value that increases from `baseVertex` through `(baseVertex +
// indexCount - 1)`. Your shader can use that value to identify each vertex in
// the `primitiveType` instance.
// 
// For more information about the `vertex_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// baseInstance: The lowest value the command passes to your vertex shader’s parameter
// with the `instance_id` attribute.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `baseInstance` through `(baseInstance + instanceCount - 1)`.
// Your shader can use that value to identify which instance the vertex
// belongs to.
// 
// For more information about the `instance_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Discussion
// 
// You can complete a primitive and start a new one by passing a sentinel
// index value that’s the largest unsigned integer possible for `indexType`.
// For example, the largest unsigned integer for [IndexTypeUInt16] and
// [IndexTypeUInt32] is `0xFFFF` and `0xFFFFFFFF`, respectively. The command
// finishes the current primitive and begins drawing a new one each time the
// command reads a sentinel index value.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferOffset, instanceCount, baseVertex, baseInstance)
	}
// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices and indirect arguments.
//
// primitiveType: An [MTLPrimitiveType] instance that represents how the command interprets
// vertex argument data.
// 
// See the [setVertexBuffer(_:offset:index:)] method and its siblings for more
// information about setting an entry in the vertex shader argument table for
// buffers.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// indexType: An [MTLIndexType] instance that represents the index’s format, including
// [IndexTypeUInt16] and [IndexTypeUInt32].
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: An [MTLBuffer] instance that contains the vertex indices of the `indexType`
// format.
//
// indexBufferOffset: An integer that represents the location that’s a multiple of the index
// size from the start of `indexBuffer` where the vertex indices begin.
//
// indirectBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLDrawIndexedPrimitivesIndirectArguments] structure.
// //
// [MTLDrawIndexedPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawIndexedPrimitivesIndirectArguments
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectBuffer` where the indirect arguments structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// # Discussion
// 
// Indirect drawing methods may help your app avoid expensive latency costs.
// This is because the command reads arguments from an [MTLBuffer] instance
// instead of using the CPU to pass parameters to the command.
// 
// You can complete a primitive and start a new one by passing a sentinel
// index value that’s the largest unsigned integer possible for `indexType`.
// For example, the largest unsigned integer for [IndexTypeUInt16] and
// [IndexTypeUInt32] is `0xFFFF` and `0xFFFFFFFF`, respectively. The command
// finishes the current primitive and begins drawing a new one each time the
// command reads a sentinel index value.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPrimitives(type:indexType:indexBuffer:indexBufferOffset:indirectBuffer:indirectBufferOffset:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferOffsetIndirectBufferIndirectBufferOffset(primitiveType MTLPrimitiveType, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexType:indexBuffer:indexBufferOffset:indirectBuffer:indirectBufferOffset:"), primitiveType, indexType, indexBuffer, indexBufferOffset, indirectBuffer, indirectBufferOffset)
	}
// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with a grid of threads.
//
// threadsPerGrid: An [MTLSize] instance that represents the number of threads for each grid
// dimension.
// 
// For mesh shaders, the command rounds the value down to the nearest multiple
// of `threadsPerMeshThreadgroup` for each dimension.
// 
// For object shaders, the value doesn’t need to be a multiple of
// `threadsPerObjectThreadgroup`.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerObjectThreadgroup: An [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: An [MTLSize] instance that represents the number of threads in a mesh
// shader threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreads(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
func (o MTLRenderCommandEncoderObject) DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreads:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}
// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with a grid of threadgroups.
//
// threadgroupsPerGrid: An [MTLSize] instance that represents the number of threadgroups for each
// grid dimension.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerObjectThreadgroup: An [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: An [MTLSize] instance that represents the number of threads in a mesh
// shader threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreadgroups(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
func (o MTLRenderCommandEncoderObject) DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreadgroups:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadgroupsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}
// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with indirect arguments.
//
// indirectBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLDispatchThreadgroupsIndirectArguments] structure.
// //
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectBuffer` where the indirect arguments structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// threadsPerObjectThreadgroup: An [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: An [MTLSize] instance that represents the number of threads in a mesh
// shader threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawMeshThreadgroups(indirectBuffer:indirectBufferOffset:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
func (o MTLRenderCommandEncoderObject) DrawMeshThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer MTLBuffer, indirectBufferOffset uint, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), indirectBuffer, indirectBufferOffset, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}
// Encodes a draw command that renders multiple instances of tessellated
// patches.
//
// numberOfPatchControlPoints: The number of control points for each patch, which needs to be in the range
// `[0, 32]`.
//
// patchStart: The first patch the command draws.
//
// patchCount: The number of patches the command draws for each instance.
//
// patchIndexBuffer: An [MTLBuffer] instance that contains the indices to patches.
//
// patchIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `patchIndexBuffer` where the patch indices begin.
//
// instanceCount: The number of times the command draws `patchCount` patches.
//
// baseInstance: The lowest value the command passes to your vertex shader’s parameter
// with the `instance_id` attribute.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `baseInstance` through `(baseInstance + instanceCount - 1)`.
// Your shader can use that value to identify which instance the vertex
// belongs to.
// 
// For more information about the `instance_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Discussion
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPatches(numberOfPatchControlPoints:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:)
func (o MTLRenderCommandEncoderObject) DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:"), numberOfPatchControlPoints, patchStart, patchCount, patchIndexBuffer, patchIndexBufferOffset, instanceCount, baseInstance)
	}
// Encodes a draw command that renders multiple instances of tessellated
// patches with indirect arguments.
//
// numberOfPatchControlPoints: The number of control points for each patch, which needs to be in the range
// `[0, 32]`.
//
// patchIndexBuffer: An [MTLBuffer] instance that contains the indices to patches.
//
// patchIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `patchIndexBuffer` where the patch indices begin.
//
// indirectBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLDrawPatchIndirectArguments] structure.
// //
// [MTLDrawPatchIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawPatchIndirectArguments
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectBuffer` where the indirect arguments structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// # Discussion
// 
// Indirect drawing methods may help your app avoid expensive latency costs.
// This is because the command reads arguments from an [MTLBuffer] instance
// instead of using the CPU to pass parameters to the command.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawPatches(numberOfPatchControlPoints:patchIndexBuffer:patchIndexBufferOffset:indirectBuffer:indirectBufferOffset:)
func (o MTLRenderCommandEncoderObject) DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPatches:patchIndexBuffer:patchIndexBufferOffset:indirectBuffer:indirectBufferOffset:"), numberOfPatchControlPoints, patchIndexBuffer, patchIndexBufferOffset, indirectBuffer, indirectBufferOffset)
	}
// Encodes a draw command that renders multiple instances of tessellated
// patches with a control point index buffer.
//
// numberOfPatchControlPoints: The number of control points for each patch, which needs to be in the range
// `[0, 32]`.
//
// patchStart: The patch start index.
//
// patchCount: The number of patches in each instance.
//
// patchIndexBuffer: An [MTLBuffer] instance that contains the indices to patches.
//
// patchIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `patchIndexBuffer` where the patch indices begin.
//
// controlPointIndexBuffer: An [MTLBuffer] instance that contains the indices to control points.
//
// controlPointIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `controlPointIndexBuffer` where the control point indices begin.
//
// instanceCount: The number of times the command draws `patchCount` patches.
//
// baseInstance: The lowest value the command passes to your vertex shader’s parameter
// with the `instance_id` attribute.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `baseInstance` through `(baseInstance + instanceCount - 1)`.
// Your shader can use that value to identify which instance the vertex
// belongs to.
// 
// For more information about the `instance_id` argument attribute for vertex
// shaders, see the [Metal Shading Language Specification (PDF)].
// //
// [Metal Shading Language Specification (PDF)]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Discussion
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPatches(numberOfPatchControlPoints:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:instanceCount:baseInstance:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer MTLBuffer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:instanceCount:baseInstance:"), numberOfPatchControlPoints, patchStart, patchCount, patchIndexBuffer, patchIndexBufferOffset, controlPointIndexBuffer, controlPointIndexBufferOffset, instanceCount, baseInstance)
	}
// Encodes a draw command that renders multiple instances of tessellated
// patches with a control point index buffer and indirect arguments.
//
// numberOfPatchControlPoints: The number of control points for each patch, which needs to be in the range
// `[0, 32]`.
//
// patchIndexBuffer: An [MTLBuffer] instance that contains the indices to patches.
//
// patchIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `patchIndexBuffer` where the patch indices begin.
//
// controlPointIndexBuffer: An [MTLBuffer] instance that contains the indices to control points.
//
// controlPointIndexBufferOffset: An integer that represents the location, in bytes, from the start of
// `controlPointIndexBuffer` where the control point indices begin.
//
// indirectBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLDrawPatchIndirectArguments] structure.
// //
// [MTLDrawPatchIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawPatchIndirectArguments
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectBuffer` where the indirect arguments structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// # Discussion
// 
// Indirect drawing methods may help your app avoid expensive latency costs.
// This is because the command reads arguments from an [MTLBuffer] instance
// instead of using the CPU to pass parameters to the command.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/drawIndexedPatches(numberOfPatchControlPoints:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:indirectBuffer:indirectBufferOffset:)
func (o MTLRenderCommandEncoderObject) DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer MTLBuffer, controlPointIndexBufferOffset uint, indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPatches:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:indirectBuffer:indirectBufferOffset:"), numberOfPatchControlPoints, patchIndexBuffer, patchIndexBufferOffset, controlPointIndexBuffer, controlPointIndexBufferOffset, indirectBuffer, indirectBufferOffset)
	}
// Encodes a command that invokes GPU functions from the encoder’s current
// tile render pipeline state.
//
// threadsPerTile: An [MTLSize] instance that represents the number of threads the render pass
// uses per tile.
// 
// Set the size’s [width] and [height] properties to values that are less
// than or equal to [TileWidth] and [TileHeight], respectively. Some GPU
// families only support square tile dispatches and require the same value for
// [width] and [height]. See the [Metal feature set tables (PDF)] to check
// which GPU families support nonsquare dispatches.
// 
// Set the [depth] property to `1`.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
// [width]: https://developer.apple.com/documentation/Metal/MTLSize/width
//
// # Discussion
// 
// The command invokes the GPU function that’s in the encoder’s current
// tile render pipeline state. You can configure that state with the following
// steps:
// 
// - Configure an [MTLTileRenderPipelineDescriptor] instance. - Create a tile
// render pipeline state by calling one of the applicable methods of an
// [MTLDevice] instance, including
// [NewRenderPipelineStateWithTileDescriptorOptionsReflectionError]. - Apply
// that tile render pipeline state by calling the [SetRenderPipelineState]
// method.
// 
// The method records the encoder’s current rendering state and resources
// the command needs as it runs. You can safely change the encoder’s render
// pipeline state to encode other commands after calling this method.
// Subsequent changes to the state don’t affect the commands already in the
// encoder’s [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/dispatchThreadsPerTile(_:)
func (o MTLRenderCommandEncoderObject) DispatchThreadsPerTile(threadsPerTile MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadsPerTile:"), threadsPerTile)
	}
// The width of the tiles, in pixels, for the render command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/tileWidth
func (o MTLRenderCommandEncoderObject) TileWidth() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("tileWidth"))
	return rv
	}
// The height of the tiles, in pixels, for the render command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/tileHeight
func (o MTLRenderCommandEncoderObject) TileHeight() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("tileHeight"))
	return rv
	}
// Encodes a command that instructs the GPU to pause before starting one or
// more stages of the render pass until a pass updates a fence.
//
// fence: A fence that the pass waits for before running the stages you pass to
// `stages`.
//
// stages: The render stages that need to wait for another pass to update `fence`
// before they run.
//
// # Discussion
// 
// Synchronize memory operations of a render pass that access resources with
// an [MTLFence]. This method instructs the GPU to wait until another pass
// updates `fence` before running the stages you pass to the `stages`
// parameter. The fence indicates when the pass can access those resources
// without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a render pass that reuses a fence, wait for other passes to
// update the fence before repurposing that fence to notify subsequent passes
// with an update:
// 
// - Call the [WaitForFenceBeforeStages] method before encoding commands that
// need to wait for other passes. - Call the [UpdateFenceAfterStages] method
// after encoding commands that later passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
// 
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/waitForFence(_:before:)
func (o MTLRenderCommandEncoderObject) WaitForFenceBeforeStages(fence MTLFence, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:beforeStages:"), fence, stages)
	}
// Encodes a command that instructs the GPU to update a fence after one or
// more stages, which can unblock other passes waiting for the fence.
//
// fence: A fence the pass updates after the stages in `stages` complete.
//
// stages: The render stages that need to complete before the pass updates `fence`.
//
// # Discussion
// 
// You can synchronize memory operations of a render pass that access
// resources with an [MTLFence]. This method instructs the pass to update
// `fence` after the stages you pass to the `stages` run all their memory
// store operations to the resources it accesses. The fence indicates when
// other passes can access those resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a render pass that reuses a fence, wait for other passes to
// update the fence before repurposing that fence to notify subsequent passes
// with an update:
// 
// - Call the [WaitForFenceBeforeStages] method before encoding commands that
// need to wait for other passes. - Call the [UpdateFenceAfterStages] method
// after encoding commands that later passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
// 
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/updateFence(_:after:)
func (o MTLRenderCommandEncoderObject) UpdateFenceAfterStages(fence MTLFence, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:afterStages:"), fence, stages)
	}
// Creates a memory barrier that enforces the order of write and read
// operations for specific resource types.
//
// scope: An [MTLBarrierScope] instance that represents the resource types the
// barrier synchronizes operations on.
// //
// [MTLBarrierScope]: https://developer.apple.com/documentation/Metal/MTLBarrierScope
//
// after: The render stages of previous draw commands that modify resources of the
// types that `scope` defines.
//
// before: The render stages of subsequent draw commands that read or modify resources
// of the types that `scope` defines.
//
// # Discussion
// 
// Memory barriers ensure the relevant stages of prior draw commands finish
// updating resources before starting the stages of subsequent commands that
// depend on those resources.
// 
// To determine whether a GPU supports memory barriers, see the [Metal feature
// set tables (PDF)].
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/memoryBarrier(scope:after:before:)
func (o MTLRenderCommandEncoderObject) MemoryBarrierWithScopeAfterStagesBeforeStages(scope MTLBarrierScope, after MTLRenderStages, before MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithScope:afterStages:beforeStages:"), scope, after, before)
	}
// Encodes a command that samples hardware counters during the render pass and
// stores the data into a counter sample buffer.
//
// sampleBuffer: An [MTLCounterSampleBuffer] instance that stores the GPU hardware data.
//
// sampleIndex: An index within `sampleBuffer` the command stores the data to.
//
// barrier: A Boolean value that indicates whether the command inserts a barrier before
// sampling the counter’s data.
// 
// A barrier ensures that the commands you encode before this one complete
// before the GPU samples the hardware counters, but can negatively impact
// runtime performance.
// 
// Running this command without a barrier means the GPU can sample counters
// concurrently with other commands from the encoder.
// 
// Either way, the `barrier` parameter for the command has no impact on
// sampling commands from other passes.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
func (o MTLRenderCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), sampleBuffer, sampleIndex, barrier)
	}
// Encodes a command that runs an indirect range of commands from an indirect
// command buffer (ICB).
//
// indirectCommandbuffer: An [MTLIndirectCommandBuffer] instance that contains other commands the
// current command runs.
//
// indirectRangeBuffer: An [MTLBuffer] instance with data that matches the layout of the
// [MTLIndirectCommandBufferExecutionRange] structure.
// 
// The [length] property of that structure needs to be less than or equal to
// `0x4000` (`16,384`).
// //
// [MTLIndirectCommandBufferExecutionRange]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
// [length]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange/length
//
// indirectBufferOffset: An integer that represents the location, in bytes, from the start of
// `indirectRangeBuffer` where the execution range structure begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:
func (o MTLRenderCommandEncoderObject) ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(indirectCommandbuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:indirectBufferOffset:"), indirectCommandbuffer, indirectRangeBuffer, indirectBufferOffset)
	}
// Encodes a command that runs a range of commands from an indirect command
// buffer (ICB).
//
// indirectCommandBuffer: An [MTLIndirectCommandBuffer] instance that contains other commands the
// current command runs.
//
// executionRange: A span of integers that represent the command entries in `buffer` the
// current command runs. The number of commands needs to be less than or equal
// to `0x4000` (`16,384`).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/executeCommandsInBuffer:withRange:
func (o MTLRenderCommandEncoderObject) ExecuteCommandsInBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, executionRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:withRange:"), indirectCommandBuffer, executionRange)
	}
// Creates a memory barrier that enforces the order of write and read
// operations for specific resources.
//
// resources: A C array of [MTLResource] instances the barrier applies to.
//
// count: The number of elements in the resources array.
//
// after: The render stages of previous draw commands that modify `resources`.
//
// before: The render stages of subsequent draw commands that read or modify
// `resources`.
//
// # Discussion
// 
// Memory barriers ensure the relevant stages of prior draw commands finish
// modifying resources before starting the stages of subsequent commands that
// depend on those resources.
// 
// To determine whether a GPU supports memory barriers, see the [Metal feature
// set tables (PDF)].
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/memoryBarrierWithResources:count:afterStages:beforeStages:
func (o MTLRenderCommandEncoderObject) MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources []MTLResource, count uint, after MTLRenderStages, before MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("memoryBarrierWithResources:count:afterStages:beforeStages:"), objc.CArray(resources), count, after, before)
	}
// Configures each pixel component value, including alpha, for the render
// pipeline’s constant blend color.
//
// red: A value for the red component for the blend color constant.
//
// green: A value for the green component for the blend color constant.
//
// blue: A value for the blue component for the blend color constant.
//
// alpha: A value for the alpha component for the blend color constant.
//
// # Discussion
// 
// The alpha and color values apply to all the render pass’s attachments.
// The `red`, `green`, and `blue` color parameters apply to the
// [BlendFactorBlendColor] and [BlendFactorOneMinusBlendColor] blend factors.
// 
// The `alpha` parameter applies to the [BlendFactorBlendAlpha] and
// [BlendFactorOneMinusBlendAlpha] blend factors.
// 
// The render pipeline’s default blend color value is `0.0` for each
// parameter, which is equivalent to [BlendFactorZero]. For other blending
// factor values, see [MTLBlendFactor].
//
// [MTLBlendFactor]: https://developer.apple.com/documentation/Metal/MTLBlendFactor
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setBlendColor(red:green:blue:alpha:)
func (o MTLRenderCommandEncoderObject) SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setBlendColorRed:green:blue:alpha:"), red, green, blue, alpha)
	}
// Sets the mapping from logical shader color output to physical render pass
// color attachments.
//
// mapping: Mapping from logical shader outputs to physical outputs.
//
// # Discussion
// 
// Use this method to define how the physical color attachments you specify
// via [ColorAttachments] map to the logical color output the fragment shader
// writes to.
// 
// To use this feature, make sure to set [SupportColorAttachmentMapping] to
// [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorAttachmentMap(_:)
func (o MTLRenderCommandEncoderObject) SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setColorAttachmentMap:"), mapping)
	}
// Configures the store action for a color attachment.
//
// storeAction: A store action for the color attachment that can’t be
// [StoreActionUnknown].
//
// colorAttachmentIndex: The index of a color attachment.
//
// # Discussion
// 
// This method changes the render command encoder’s store action for a color
// attachment. You can assign the default store action for a color attachment
// by configuring the [StoreAction] property of its
// [MTLRenderPassColorAttachmentDescriptor] (see [MTLRenderPassDescriptor] and
// its [ColorAttachments] property).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorStoreAction(_:index:)
func (o MTLRenderCommandEncoderObject) SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setColorStoreAction:atIndex:"), storeAction, colorAttachmentIndex)
	}
// Configures the store action options for a color attachment.
//
// storeActionOptions: Additional options for the store action of a color attachment.
//
// colorAttachmentIndex: The index of a color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setColorStoreActionOptions(_:index:)
func (o MTLRenderCommandEncoderObject) SetColorStoreActionOptionsAtIndex(storeActionOptions MTLStoreActionOptions, colorAttachmentIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setColorStoreActionOptions:atIndex:"), storeActionOptions, colorAttachmentIndex)
	}
// Configures how the render pipeline determines which primitives to remove.
//
// cullMode: An [MTLCullMode] value that configures how the render pipeline determines
// which primitives to remove from the pipeline.
// //
// [MTLCullMode]: https://developer.apple.com/documentation/Metal/MTLCullMode
//
// # Discussion
// 
// This method configures which primitives the render pipeline removes, if
// any, based on the direction of each primitive’s face relative to the
// scene’s camera. For example, you can correctly cull hidden surfaces on
// some geometric models, such as a sphere made of filled triangles, if it
// uses orientable surfaces. A surface is if its primitives consistently use
// the same ordering for its vertices. Metal defines vertex ordering with the
// [MTLWinding] type, which includes [WindingClockwise] and
// [WindingCounterClockwise]. You can tell the render pipeline which direction
// your primitives face by calling the [SetFrontFacingWinding] method, which
// affects the primitives the culling mode removes.
// 
// The render pass’s default culling mode is [CullModeNone].
//
// [MTLWinding]: https://developer.apple.com/documentation/Metal/MTLWinding
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setCullMode(_:)
func (o MTLRenderCommandEncoderObject) SetCullMode(cullMode MTLCullMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setCullMode:"), cullMode)
	}
// Configures the adjustments a render pass applies to depth values from
// fragment functions by a scaling factor and bias.
//
// depthBias: A constant bias the render pipeline applies to all fragments.
//
// slopeScale: A bias coefficient that scales with the depth of the primitive relative to
// the camera.
//
// clamp: A value that limits the bias value the render pipeline can apply to a
// fragment. Pass a positive or negative value to limit the largest magnitude
// of a positive or negative bias, respectively.
// 
// You can disable the bias clamping functionality by passing `0.0`.
//
// # Discussion
// 
// Call this method to have the render pipeline apply a bias to the rasterized
// depth after the clipping stage. The bias affects both depth testing and the
// values the render pipeline writes to the depth render target. If you
// don’t explicitly call this method, the pipeline doesn’t apply a scale
// or a bias to a depth value.
// 
// Set a depth bias to improve the quality of techniques such as shadow
// mapping and avoid depth artifacts like shadow acne.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthBias(_:slopeScale:clamp:)
func (o MTLRenderCommandEncoderObject) SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthBias:slopeScale:clamp:"), depthBias, slopeScale, clamp)
	}
// Configures how the render pipeline handles fragments outside the near and
// far planes of the view frustum.
//
// depthClipMode: The mode that determines how to handle fragments outside the near and far
// planes.
//
// # Discussion
// 
// You can use depth clipping to ignore fragments outside the z-axis
// boundaries of a viewing volume.
// 
// The render pass’s default clip mode is [DepthClipModeClip].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthClipMode(_:)
func (o MTLRenderCommandEncoderObject) SetDepthClipMode(depthClipMode MTLDepthClipMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthClipMode:"), depthClipMode)
	}
// Configures the combined depth and stencil state.
//
// depthStencilState: An instance that conforms to the [MTLDepthStencilState] protocol.
//
// # Discussion
// 
// This method changes the combined depth and stencil state for the render
// command encoder that’s compatible with its depth and stencil attachment
// configuration. For example, if the new state enables depth testing or depth
// writing, the render pass needs to have a depth attachment. Similarly, if
// the new state enables stencil testing or stencil writing, the render
// pass’s stencil needs to have a stencil attachment. You create depth and
// stencil attachments for a render pass by assigning the [DepthAttachment]
// and [StencilAttachment] properties of the [MTLRenderPassDescriptor]
// instance that creates it.
// 
// Pass `nil` to clear the state from the previous call, which restores a
// state that’s equivalent to the default values of an
// [MTLDepthStencilDescriptor] instance’s properties.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStencilState(_:)
func (o MTLRenderCommandEncoderObject) SetDepthStencilState(depthStencilState MTLDepthStencilState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStencilState:"), depthStencilState)
	}
// Configures the store action for the depth attachment.
//
// storeAction: A store action for the depth attachment that can’t be
// [StoreActionUnknown].
//
// # Discussion
// 
// This method changes the render command encoder’s store action for the
// depth attachment. You can assign the default store action for the depth
// attachment by configuring the [StoreAction] property of its
// [MTLRenderPassDepthAttachmentDescriptor] (see [MTLRenderPassDescriptor] and
// its [DepthAttachment] property).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStoreAction(_:)
func (o MTLRenderCommandEncoderObject) SetDepthStoreAction(storeAction MTLStoreAction) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStoreAction:"), storeAction)
	}
// Configures the store action options for the depth attachment.
//
// storeActionOptions: Additional options for the store action of the depth attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthStoreActionOptions(_:)
func (o MTLRenderCommandEncoderObject) SetDepthStoreActionOptions(storeActionOptions MTLStoreActionOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStoreActionOptions:"), storeActionOptions)
	}
// Configures the minimum and maximum bounds for depth bounds testing.
//
// minBound: A minimum bound for depth testing, which discards fragments with a stored
// depth that is less than `minBound`.
//
// maxBound: A maximum bound for depth testing, which discards fragments with a stored
// depth that is greater than `maxBound`.
//
// # Discussion
// 
// The render command encoder disables depth bounds testing by default. The
// render command encoder also disables depth bounds testing when all of the
// following properties equal a specific value:
// 
// - The `minBound` property is equal to `0.0f`. - The `maxBound` property is
// equal to `1.0f`. Both `minBound` and `maxBound` need to be within `[0.0f,
// 1.0f]`, and `minBound` needs to be less than or equal to `maxBound`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setDepthTestMinBound:maxBound:
func (o MTLRenderCommandEncoderObject) SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthTestMinBound:maxBound:"), minBound, maxBound)
	}
// Assigns an acceleration structure to an entry in the fragment shader
// argument table.
//
// accelerationStructure: An [MTLAccelerationStructure] instance the command assigns to an entry in
// the fragment shader argument table for acceleration structures.
//
// bufferIndex: An integer that represents the entry in the fragment shader argument table
// for acceleration structures that stores a record of
// `accelerationStructure`.
//
// # Discussion
// 
// By default, the acceleration structure at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentAccelerationStructure(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetFragmentAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentAccelerationStructure:atBufferIndex:"), accelerationStructure, bufferIndex)
	}
// Updates an entry in the fragment shader argument table with a new location
// within the entry’s current buffer.
//
// offset: An integer that represents the location, in bytes, from the start of
// `buffer` where the fragment shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// index: An integer that represents the entry in the fragment shader argument table
// for buffers that already stores a record of an [MTLBuffer].
//
// # Discussion
// 
// The command this method encodes changes the offset for a fragment buffer
// that already has a previous assignment from one of your earlier commands.
// 
// For more information, see:
// 
// - [setFragmentBuffer(_:offset:index:)] -
// [setFragmentBuffers(_:offsets:range:)] (Swift) -
// [SetFragmentBuffersOffsetsWithRange] (Objective-C)
// 
// The command can also adjust the offset for an entry that you previously set
// with the [SetFragmentBytesLengthAtIndex] method.
// 
// By default, the buffer at each index is `nil`.
//
// [setFragmentBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBuffer(_:offset:index:)
// [setFragmentBuffers(_:offsets:range:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBuffers(_:offsets:range:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBufferOffset(_:index:)
func (o MTLRenderCommandEncoderObject) SetFragmentBufferOffsetAtIndex(offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentBufferOffset:atIndex:"), offset, index)
	}
// Assigns multiple buffers to a range of entries in the fragment shader
// argument table.
//
// buffers: A pointer to a C array of [MTLBuffer] instances the command assigns to
// entries in the fragment shader argument table for buffers.
//
// offsets: A pointer to a C array of unsigned integers. Each element represents the
// location, in bytes, from the start of the corresponding [MTLBuffer] element
// in `buffers` where the fragment shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for buffers. Each entry stores a record of the corresponding
// element in `buffers` and `offsets`.
//
// # Discussion
// 
// By default, the buffer at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBuffers:offsets:withRange:
func (o MTLRenderCommandEncoderObject) SetFragmentBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Creates a buffer from bytes and assigns it to an entry in the fragment
// shader argument table.
//
// bytes: A pointer to argument data the method copies to an [MTLBuffer] and assigns
// to an entry in the fragment shader argument table for buffers.
//
// length: The number of bytes the method copies from the `bytes` pointer.
//
// index: An integer that represents the entry in the fragment shader argument table
// for buffers that stores a record of the [MTLBuffer] the method creates from
// `bytes`.
//
// # Discussion
// 
// The method is equivalent to creating an [MTLBuffer] instance that contains
// the same data as `bytes` and calling the
// [setFragmentBuffer(_:offset:index:)] method. However, this method avoids
// the overhead of creating a buffer to store your data; instead, Metal
// manages the data.
// 
// For data that’s more than 4 KB, create an [MTLBuffer] instance and pass
// it to [setFragmentBuffer(_:offset:index:)].
// 
// By default, the buffer at each index is `nil`.
//
// [setFragmentBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBuffer(_:offset:index:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentBytes(_:length:index:)
func (o MTLRenderCommandEncoderObject) SetFragmentBytesLengthAtIndex(bytes []byte, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
// Assigns an intersection function table to an entry in the fragment shader
// argument table.
//
// intersectionFunctionTable: An [MTLIntersectionFunctionTable] instance the command assigns to an entry
// in the fragment shader argument table for intersection function tables.
//
// bufferIndex: An integer that represents the entry in the fragment shader argument table
// for intersection function tables that stores a record of
// `intersectionFunctionTable`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentIntersectionFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetFragmentIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentIntersectionFunctionTable:atBufferIndex:"), intersectionFunctionTable, bufferIndex)
	}
// Assigns multiple intersection function tables to a range of entries in the
// fragment shader argument table.
//
// intersectionFunctionTables: A pointer to a C array of [MTLIntersectionFunctionTable] instances the
// command assigns to entries in the fragment shader argument table for
// intersection function tables.
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for intersection function tables. Each entry stores a record
// of the corresponding element in `intersectionFunctionTables`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentIntersectionFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetFragmentIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentIntersectionFunctionTables:withBufferRange:"), intersectionFunctionTables, range_)
	}
// Assigns a sampler state to an entry in the fragment shader argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the
// fragment shader argument table for sampler states.
//
// index: An integer that represents the entry in the fragment shader argument table
// for sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerState(_:index:)
func (o MTLRenderCommandEncoderObject) SetFragmentSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentSamplerState:atIndex:"), sampler, index)
	}
// Assigns a sampler state and clamp values to an entry in the fragment shader
// argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the
// fragment shader argument table for sampler states.
//
// lodMinClamp: The smallest level of detail value a fragment shader can use when it
// samples a texture.
//
// lodMaxClamp: The largest level of detail value a fragment shader can use when it samples
// a texture.
//
// index: An integer that represents the entry in the fragment shader argument table
// for sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// The method’s `lodMinClamp` and `lodMaxClamp` parameters override the
// default values for `sampler`. You can set the sampler’s default values by
// configuring the [LodMinClamp] and [LodMaxClamp] properties of
// [MTLSamplerDescriptor] before you create the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLRenderCommandEncoderObject) SetFragmentSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
	}
// Assigns multiple sampler states and clamp values to a range of entries in
// the fragment shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the fragment shader argument table for sampler states.
//
// lodMinClamps: A pointer to a C array of floating-point values. Each element is the
// smallest level of detail value a fragment shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// lodMaxClamps: A pointer to a C array of floating-point values. Each element is the
// largest level of detail value a fragment shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for sampler states. Each entry stores a record of the
// corresponding element in `samplers`.
//
// # Discussion
// 
// Each element of the method’s `lodMinClamps` and `lodMaxClamps` parameters
// overrides the default values for the corresponding sampler in `samplers`.
// You can set a sampler’s default values by configuring the [LodMinClamp]
// and [LodMaxClamp] properties of [MTLSamplerDescriptor] before you create
// the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLRenderCommandEncoderObject) SetFragmentSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
	}
// Assigns multiple sampler states to a range of entries in the fragment
// shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the fragment shader argument table for sampler states.
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for sampler states. Each entry stores a record of the
// corresponding element in `samplers`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentSamplerStates:withRange:
func (o MTLRenderCommandEncoderObject) SetFragmentSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentSamplerStates:withRange:"), samplers, range_)
	}
// Assigns a texture to an entry in the fragment shader argument table.
//
// texture: An [MTLTexture] instance the command assigns to an entry in the fragment
// shader argument table for textures.
//
// index: An integer that represents the entry in the fragment shader argument table
// for textures that stores a record of `texture`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentTexture(_:index:)
func (o MTLRenderCommandEncoderObject) SetFragmentTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentTexture:atIndex:"), texture, index)
	}
// Assigns multiple textures to a range of entries in the fragment shader
// argument table.
//
// textures: A pointer to a C array of [MTLTexture] instances the command assigns to
// entries in the fragment shader argument table for textures.
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for textures. Each entry stores a record of the
// corresponding element in `textures`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentTextures:withRange:
func (o MTLRenderCommandEncoderObject) SetFragmentTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentTextures:withRange:"), textures, range_)
	}
// Assigns a visible function table to an entry in the fragment shader
// argument table.
//
// functionTable: An [MTLVisibleFunctionTable] instance the command assigns to an entry in
// the fragment shader argument table for visible function tables.
//
// bufferIndex: An integer that represents the entry in the fragment shader argument table
// for visible function tables that stores a record of `functionTable`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentVisibleFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetFragmentVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentVisibleFunctionTable:atBufferIndex:"), functionTable, bufferIndex)
	}
// Assigns multiple visible function tables to a range of entries in the
// fragment shader argument table.
//
// functionTables: A pointer to a C array of [MTLVisibleFunctionTable] instances the command
// assigns to entries in the fragment shader argument table for visible
// function tables.
//
// range: A span of integers that represent the entries in the fragment shader
// argument table for visible function tables. Each entry stores a record of
// the corresponding element in `functionTables`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFragmentVisibleFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetFragmentVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentVisibleFunctionTables:withBufferRange:"), functionTables, range_)
	}
// Configures which face of a primitive, such as a triangle, is the front.
//
// frontFacingWinding: An [MTLWinding] value that configures how the render pipeline defines which
// side of a primitive is its front.
// //
// [MTLWinding]: https://developer.apple.com/documentation/Metal/MTLWinding
//
// # Discussion
// 
// The render pass’s default front-facing mode is [WindingClockwise].
// 
// The winding direction of a primitive determines whether the render pass
// culls it (see [SetCullMode]).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setFrontFacing(_:)
func (o MTLRenderCommandEncoderObject) SetFrontFacingWinding(frontFacingWinding MTLWinding) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFrontFacingWinding:"), frontFacingWinding)
	}
// Updates an entry in the mesh shader argument table with a new location
// within the entry’s current buffer.
//
// offset: An integer that represents the location, in bytes, from the start of
// `buffer` where the mesh shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// index: An integer that represents the entry in the mesh shader argument table for
// buffers that already stores a record of an [MTLBuffer].
//
// # Discussion
// 
// The command this method encodes changes the offset for a mesh buffer that
// already has a previous assignment from one of your earlier commands.
// 
// For more information, see:
// 
// - [setMeshBuffer(_:offset:index:)] - [setMeshBuffers(_:offsets:range:)]
// (Swift) - [SetMeshBuffersOffsetsWithRange] (Objective-C)
// 
// The command can also adjust the offset for an entry that you previously set
// with the [SetMeshBytesLengthAtIndex] method.
//
// [setMeshBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBuffer(_:offset:index:)
// [setMeshBuffers(_:offsets:range:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBuffers(_:offsets:range:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBufferOffset(_:index:)
func (o MTLRenderCommandEncoderObject) SetMeshBufferOffsetAtIndex(offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshBufferOffset:atIndex:"), offset, index)
	}
// Assigns multiple buffers to a range of entries in the mesh shader argument
// table.
//
// buffers: A pointer to a C array of [MTLBuffer] instances the command assigns to
// entries in the mesh shader argument table for buffers.
//
// offsets: A pointer to a C array of unsigned integers. Each element represents the
// location, in bytes, from the start of the corresponding [MTLBuffer] element
// in `buffers` where the mesh shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// range: A span of integers that represent the entries in the mesh shader argument
// table for buffers. Each entry stores a record of the corresponding element
// in `buffers` and `offsets`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBuffers:offsets:withRange:
func (o MTLRenderCommandEncoderObject) SetMeshBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Creates a buffer from bytes and assigns it to an entry in the mesh shader
// argument table.
//
// bytes: A pointer to argument data the method copies to an [MTLBuffer] and assigns
// to an entry in the mesh shader argument table for buffers.
//
// length: The number of bytes the method copies from the `bytes` pointer.
//
// index: An integer that represents the entry in the mesh shader argument table for
// buffers that stores a record of the [MTLBuffer] the method creates from
// `bytes`.
//
// # Discussion
// 
// The method is equivalent to creating an [MTLBuffer] instance that contains
// the same data as `bytes` and calling the [setMeshBuffer(_:offset:index:)]
// method. However, this method avoids the overhead of creating a buffer to
// store your data; instead, Metal manages the data.
// 
// For data that’s more than 4 KB, create an [MTLBuffer] instance and pass
// it to [SetMeshBufferOffsetAtIndex].
//
// [setMeshBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBuffer(_:offset:index:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshBytes(_:length:index:)
func (o MTLRenderCommandEncoderObject) SetMeshBytesLengthAtIndex(bytes []byte, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
// Assigns a sampler state to an entry in the mesh shader argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the mesh
// shader argument table for sampler states.
//
// index: An integer that represents the entry in the mesh shader argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerState(_:index:)
func (o MTLRenderCommandEncoderObject) SetMeshSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshSamplerState:atIndex:"), sampler, index)
	}
// Assigns a sampler state and clamp values to an entry in the mesh shader
// argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the mesh
// shader argument table for sampler states.
//
// lodMinClamp: The smallest level of detail value a mesh shader can use when it samples a
// texture.
//
// lodMaxClamp: The largest level of detail value a mesh shader can use when it samples a
// texture.
//
// index: An integer that represents the entry in the mesh shader argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// The method’s `lodMinClamp` and `lodMaxClamp` parameters override the
// default values for `sampler`. You can set the sampler’s default values by
// configuring the [LodMinClamp] and [LodMaxClamp] properties of
// [MTLSamplerDescriptor] before you create the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLRenderCommandEncoderObject) SetMeshSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
	}
// Assigns multiple sampler states and clamp values to a range of entries in
// the mesh shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the mesh shader argument table for sampler states.
//
// lodMinClamps: A pointer to a C array of floating-point values. Each element is the
// smallest level of detail value a mesh shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// lodMaxClamps: A pointer to a C array of floating-point values. Each element is the
// largest level of detail value a mesh shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// range: A span of integers that represent the entries in the mesh shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// Each element of the method’s `lodMinClamps` and `lodMaxClamps` parameters
// overrides the default values for the corresponding sampler in `samplers`.
// You can set a sampler’s default values by configuring the [LodMinClamp]
// and [LodMaxClamp] properties of [MTLSamplerDescriptor] before you create
// the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLRenderCommandEncoderObject) SetMeshSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
	}
// Assigns multiple sampler states to a range of entries in the mesh shader
// argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the mesh shader argument table for sampler states.
//
// range: A span of integers that represent the entries in the mesh shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshSamplerStates:withRange:
func (o MTLRenderCommandEncoderObject) SetMeshSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshSamplerStates:withRange:"), samplers, range_)
	}
// Assigns a texture to an entry in the mesh shader argument table.
//
// texture: An [MTLTexture] instance the command assigns to an entry in the mesh shader
// argument table for textures.
//
// index: An integer that represents the entry in the mesh shader argument table for
// textures that stores a record of `texture`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshTexture(_:index:)
func (o MTLRenderCommandEncoderObject) SetMeshTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshTexture:atIndex:"), texture, index)
	}
// Assigns multiple textures to a range of entries in the mesh shader argument
// table.
//
// textures: A pointer to a C array of [MTLTexture] instances the command assigns to
// entries in the mesh shader argument table for textures.
//
// range: A span of integers that represent the entries in the mesh shader argument
// table for textures. Each entry stores a record of the corresponding element
// in `textures`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setMeshTextures:withRange:
func (o MTLRenderCommandEncoderObject) SetMeshTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshTextures:withRange:"), textures, range_)
	}
// Updates an entry in the object shader argument table with a new location
// within the entry’s current buffer.
//
// offset: An integer that represents the location, in bytes, from the start of
// `buffer` where the object shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// index: An integer that represents the entry in the object shader argument table
// for buffers that already stores a record of an [MTLBuffer].
//
// # Discussion
// 
// The command this method encodes changes the offset for a mesh buffer that
// already has a previous assignment from one of your earlier commands.
// 
// For more information, see:
// 
// - [setObjectBuffer(_:offset:index:)] - [setObjectBuffers(_:offsets:range:)]
// (Swift) - [SetObjectBuffersOffsetsWithRange] (Objective-C)
// 
// The command can also adjust the offset for an entry that you previously set
// with the [SetObjectBytesLengthAtIndex] method.
//
// [setObjectBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBuffer(_:offset:index:)
// [setObjectBuffers(_:offsets:range:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBuffers(_:offsets:range:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBufferOffset(_:index:)
func (o MTLRenderCommandEncoderObject) SetObjectBufferOffsetAtIndex(offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectBufferOffset:atIndex:"), offset, index)
	}
// Encodes a command that assigns multiple buffers to a range of entries in
// the object shader argument table.
//
// buffers: A pointer to a C array of [MTLBuffer] instances the command assigns to
// entries in the object shader argument table for buffers.
//
// offsets: A pointer to a C array of unsigned integers. Each element represents the
// location, in bytes, from the start of the corresponding [MTLBuffer] element
// in `buffers` where the object shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// range: A span of integers that represent the entries in the object shader argument
// table for buffers. Each entry stores a record of the corresponding element
// in `buffers` and `offsets`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBuffers:offsets:withRange:
func (o MTLRenderCommandEncoderObject) SetObjectBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Creates a buffer from bytes and assigns it to an entry in the object shader
// argument table.
//
// bytes: A pointer to argument data the method copies to an [MTLBuffer] and assigns
// to an entry in the object shader argument table for buffers.
//
// length: The number of bytes the method copies from the `bytes` pointer.
//
// index: An integer that represents the entry in the object shader argument table
// for buffers that stores a record of the [MTLBuffer] the method creates from
// `bytes`.
//
// # Discussion
// 
// The method is equivalent to creating an [MTLBuffer] instance that contains
// the same data as `bytes` and calling the [SetObjectBufferOffsetAtIndex]
// method. However, this method avoids the overhead of creating a buffer to
// store your data; instead, Metal manages the data.
// 
// For data that’s more than 4 KB, create an [MTLBuffer] instance and pass
// it to [setObjectBuffer(_:offset:index:)].
//
// [setObjectBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBuffer(_:offset:index:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectBytes(_:length:index:)
func (o MTLRenderCommandEncoderObject) SetObjectBytesLengthAtIndex(bytes []byte, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
// Assigns a sampler state to an entry in the object shader argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the object
// shader argument table for sampler states.
//
// index: An integer that represents the entry in the object argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerState(_:index:)
func (o MTLRenderCommandEncoderObject) SetObjectSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectSamplerState:atIndex:"), sampler, index)
	}
// Assigns a sampler state and clamp values to an entry in the object shader
// argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the object
// shader argument table for sampler states.
//
// lodMinClamp: The smallest level of detail value an object shader can use when it samples
// a texture.
//
// lodMaxClamp: The largest level of detail value an object shader can use when it samples
// a texture.
//
// index: An integer that represents the entry in the object argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// The method’s `lodMinClamp` and `lodMaxClamp` parameters override the
// default values for `sampler`. You can set the sampler’s default values by
// configuring the [LodMinClamp] and [LodMaxClamp] properties of
// [MTLSamplerDescriptor] before you create the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLRenderCommandEncoderObject) SetObjectSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
	}
// Assigns multiple sampler states and clamp values to a range of entries in
// the object shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the object shader argument table for sampler states.
//
// lodMinClamps: A pointer to a C array of floating-point values. Each element is the
// smallest level of detail value an object shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// lodMaxClamps: A pointer to a C array of floating-point values. Each element is the
// largest level of detail value an object shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// range: A span of integers that represent the entries in the object shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// Each element of the method’s `lodMinClamps` and `lodMaxClamps` parameters
// overrides the default values for the corresponding sampler in `samplers`.
// You can set a sampler’s default values by configuring the [LodMinClamp]
// and [LodMaxClamp] properties of [MTLSamplerDescriptor] before you create
// the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLRenderCommandEncoderObject) SetObjectSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
	}
// Assigns multiple sampler states to a range of entries in the object shader
// argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the object shader argument table for sampler states.
//
// range: A span of integers that represent the entries in the object shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectSamplerStates:withRange:
func (o MTLRenderCommandEncoderObject) SetObjectSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectSamplerStates:withRange:"), samplers, range_)
	}
// Assigns a texture to an entry in the object shader argument table.
//
// texture: An [MTLTexture] instance the command assigns to an entry in the object
// shader argument table for textures.
//
// index: An integer that represents the entry in the object shader argument table
// for textures that stores a record of `texture`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectTexture(_:index:)
func (o MTLRenderCommandEncoderObject) SetObjectTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectTexture:atIndex:"), texture, index)
	}
// Assigns multiple textures to a range of entries in the object shader
// argument table.
//
// textures: A pointer to a C array of [MTLTexture] instances the command assigns to
// entries in the object shader argument table for textures.
//
// range: A span of integers that represent the entries in the object shader argument
// table for textures. Each entry stores a record of the corresponding element
// in `textures`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectTextures:withRange:
func (o MTLRenderCommandEncoderObject) SetObjectTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectTextures:withRange:"), textures, range_)
	}
// Configures the size of a threadgroup memory buffer for an entry in the
// object argument table.
//
// length: The threadgroup memory length, in bytes.
//
// index: An integer that represents an entry in the object argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setObjectThreadgroupMemoryLength(_:index:)
func (o MTLRenderCommandEncoderObject) SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectThreadgroupMemoryLength:atIndex:"), length, index)
	}
// Configures the encoder with a render or tile pipeline state that applies to
// your subsequent draw commands.
//
// pipelineState: A render pipeline state that you create by calling an [MTLDevice] methods
// (see [Pipeline state creation]).
// //
// [Pipeline state creation]: https://developer.apple.com/documentation/Metal/pipeline-state-creation
//
// # Discussion
// 
// Set the render pass’s render pipeline state before encoding any draw or
// tile commands by calling this method because the default pipeline state is
// `nil`.
// 
// You can change which pipeline state the encoder uses multiple times during
// its lifetime. For example, your app may want render some things with a
// vertex shader, and render others with an object and mesh shader. Changing
// the pipeline state only affects the subsequent commands and has no effect
// on the commands you encode before changing the state.
// 
// The render pipeline you pass to this method needs to be compatible with the
// render pass’s attachments. You configure these attachments with the
// properties of an [MTLRenderPassDescriptor] instance, including
// [ColorAttachments], [DepthAttachment], and [StencilAttachment].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setRenderPipelineState(_:)
func (o MTLRenderCommandEncoderObject) SetRenderPipelineState(pipelineState MTLRenderPipelineState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setRenderPipelineState:"), pipelineState)
	}
// Configures a rectangle for the fragment scissor test.
//
// rect: An [MTLScissorRect] instance that represents a rectangle that needs to lie
// completely within the current render attachment.
// //
// [MTLScissorRect]: https://developer.apple.com/documentation/Metal/MTLScissorRect
//
// # Discussion
// 
// The rendering pipeline discards any fragments that lie outside the scissor
// rectangle.
// 
// The default scissor rectangle is the same size as the current render
// attachment, with its origin coordinates in the upper-left corner at `(0,
// 0)`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setScissorRect(_:)
func (o MTLRenderCommandEncoderObject) SetScissorRect(rect MTLScissorRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setScissorRect:"), rect)
	}
// Configures multiple rectangles for the fragment scissor test.
//
// scissorRects: An array of [MTLScissorRect] instances the command applies to the render
// pipeline for clipping.
// //
// [MTLScissorRect]: https://developer.apple.com/documentation/Metal/MTLScissorRect
//
// count: The number of elements in the `scissorRects` array.
//
// # Discussion
// 
// The rendering pipeline discards any fragments that lie outside the scissor
// rectangle. The default scissor rectangle is the same size as the current
// render attachment, with its origin coordinates in the upper-left corner at
// `(0, 0)`.
// 
// Use this method to configure a different scissor rectangle for multiple
// viewports you configure with the [setViewports(_:)] method. Multiple
// viewports give your app the ability to draw into separate areas of an image
// with a single draw call. You can either set a single scissor rectangle for
// all viewports with the [SetScissorRect] method, or set each viewport’s
// rectangle with this method.
// 
// The maximum number of viewports and scissor rectangles a GPU supports
// varies by device family. For more information, see [MTLGPUFamily] and
// [Detecting GPU features and Metal software versions].
// 
// The rendering pipeline sends each primitive to a single viewport and its
// associated scissor rectangle. You can select which viewport each primitive
// uses in your vertex shader by adding the `[[viewport_array_index]]`
// attribute to an output value.
// 
// The [SetScissorRect] method is equivalent to calling this method with a
// single element in the `scissorRects` array.
//
// [Detecting GPU features and Metal software versions]: https://developer.apple.com/documentation/Metal/detecting-gpu-features-and-metal-software-versions
// [MTLGPUFamily]: https://developer.apple.com/documentation/Metal/MTLGPUFamily
// [setViewports(_:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setViewports(_:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setScissorRects:count:
func (o MTLRenderCommandEncoderObject) SetScissorRectsCount(scissorRects []MTLScissorRect, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setScissorRects:count:"), objc.CArray(scissorRects), count)
	}
// Configures different comparison values for front- and back-facing
// primitives.
//
// frontReferenceValue: A stencil test comparison value the render pipeline applies to only
// front-facing primitives.
//
// backReferenceValue: A stencil test comparison value the render pipeline applies to only
// back-facing primitives.
//
// # Discussion
// 
// The command sets separate reference values for front- and back-facing
// primitives (see [StencilCompareFunction], [FrontFaceStencil], and
// [BackFaceStencil]). These reference values apply to the stencil state you
// set with the [SetDepthStencilState] method.
// 
// The render pass’s default reference value for the front and back stencil
// compare function is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilReferenceValues(front:back:)
func (o MTLRenderCommandEncoderObject) SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilFrontReferenceValue:backReferenceValue:"), frontReferenceValue, backReferenceValue)
	}
// Configures the same comparison value for front- and back-facing primitives.
//
// referenceValue: A stencil test comparison value the render pipeline applies to both front-
// and back-facing primitives.
//
// # Discussion
// 
// The command sets the same reference value for front- and back-facing
// primitives (see [StencilCompareFunction], [FrontFaceStencil], and
// [BackFaceStencil]). This reference value applies to the stencil state you
// set with the [SetDepthStencilState] method.
// 
// The render pass’s default reference value for the front and back stencil
// compare function is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilReferenceValue(_:)
func (o MTLRenderCommandEncoderObject) SetStencilReferenceValue(referenceValue uint32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilReferenceValue:"), referenceValue)
	}
// Configures the store action for the stencil attachment.
//
// storeAction: A store action for the stencil attachment that can’t be
// [StoreActionUnknown].
//
// # Discussion
// 
// This method changes the render command encoder’s store action for the
// stencil attachment. You can assign the default store action for the stencil
// attachment by configuring the [StoreAction] property of its
// [MTLRenderPassStencilAttachmentDescriptor] (see [MTLRenderPassDescriptor]
// and its [StencilAttachment] property).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilStoreAction(_:)
func (o MTLRenderCommandEncoderObject) SetStencilStoreAction(storeAction MTLStoreAction) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilStoreAction:"), storeAction)
	}
// Configures the store action options for the stencil attachment.
//
// storeActionOptions: Additional options for the store action of the stencil attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setStencilStoreActionOptions(_:)
func (o MTLRenderCommandEncoderObject) SetStencilStoreActionOptions(storeActionOptions MTLStoreActionOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilStoreActionOptions:"), storeActionOptions)
	}
// Configures the per-patch tessellation factors for any subsequent
// patch-drawing commands.
//
// buffer: An [MTLBuffer] instance that stores the per-patch tessellation factors,
// which can’t be empty or `nil`.
//
// offset: The distance, in bytes, between the start of the data and the start of the
// buffer, which needs to be a multiple of `4`.
//
// instanceStride: The number of bytes between two instances of data in `buffer`, which needs
// to be a multiple of `4`.
//
// # Discussion
// 
// Call this method before encoding patch-drawing commands.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTessellationFactorBuffer(_:offset:instanceStride:)
func (o MTLRenderCommandEncoderObject) SetTessellationFactorBufferOffsetInstanceStride(buffer MTLBuffer, offset uint, instanceStride uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTessellationFactorBuffer:offset:instanceStride:"), buffer, offset, instanceStride)
	}
// Configures the scale factor for per-patch tessellation factors.
//
// scale: A positive, normal floating-point scale factor the render pass applies to
// the per-patch tessellation factors.
// 
// The value of `scale` can’t be negative, infinite, equal to [NaN] (not a
// number), or a denormalized number.
//
// # Discussion
// 
// The command converts `scale` to a half-precision floating-point value
// before it applies it to the per-patch tessellation factors (see
// [SetTessellationFactorBufferOffsetInstanceStride]).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTessellationFactorScale(_:)
func (o MTLRenderCommandEncoderObject) SetTessellationFactorScale(scale float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTessellationFactorScale:"), scale)
	}
// Configures the size of a threadgroup memory buffer for an entry in the
// fragment or tile shader argument table.
//
// length: The threadgroup memory length, in bytes.
//
// offset: An integer that represents the location, in bytes, from the start of the
// buffer at `index` where the threadgroup memory begins.
//
// index: An integer that represents an entry in the buffer argument table.
//
// # Discussion
// 
// You can only change the threadgroup memory’s size between tile dispatches
// (see [DispatchThreadsPerTile]).
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setThreadgroupMemoryLength(_:offset:index:)
func (o MTLRenderCommandEncoderObject) SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:offset:atIndex:"), length, offset, index)
	}
// Assigns an acceleration structure to an entry in the tile shader argument
// table.
//
// accelerationStructure: An [MTLAccelerationStructure] instance the command assigns to an entry in
// the tile shader argument table for acceleration structures.
//
// bufferIndex: An integer that represents the entry in the tile shader argument table for
// acceleration structures that stores a record of `accelerationStructure`.
//
// # Discussion
// 
// By default, the acceleration structure at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileAccelerationStructure(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetTileAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileAccelerationStructure:atBufferIndex:"), accelerationStructure, bufferIndex)
	}
// Updates an entry in the tile shader argument table with a new location
// within the entry’s current buffer.
//
// offset: An integer that represents the location, in bytes, from the start of
// `buffer` where the tile shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// index: An integer that represents the entry in the tile shader argument table for
// buffers that already stores a record of an [MTLBuffer].
//
// # Discussion
// 
// The command this method encodes changes the offset for a fragment buffer
// that already has a previous assignment from one of your earlier commands.
// 
// For more information, see:
// 
// - [setTileBuffer(_:offset:index:)] - [setTileBuffers(_:offsets:range:)]
// (Swift) - [SetTileBuffersOffsetsWithRange] (Objective-C)
// 
// The command can also adjust the offset for an entry that you previously set
// with the [SetTileBytesLengthAtIndex] method.
// 
// By default, the buffer at each index is `nil`.
//
// [setTileBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBuffer(_:offset:index:)
// [setTileBuffers(_:offsets:range:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBuffers(_:offsets:range:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBufferOffset(_:index:)
func (o MTLRenderCommandEncoderObject) SetTileBufferOffsetAtIndex(offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileBufferOffset:atIndex:"), offset, index)
	}
// Assigns multiple buffers to a range of entries in the tile shader argument
// table.
//
// buffers: A pointer to a C array of [MTLBuffer] instances the command assigns to
// entries in the tile shader argument table for buffers.
//
// offsets: A pointer to a C array of unsigned integers. Each element represents the
// location, in bytes, from the start of the corresponding [MTLBuffer] element
// in `buffers` where the tile shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// range: A span of integers that represent the entries in the tile shader argument
// table for buffers. Each entry stores a record of the corresponding element
// in `buffers` and `offsets`.
//
// # Discussion
// 
// By default, the buffer at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBuffers:offsets:withRange:
func (o MTLRenderCommandEncoderObject) SetTileBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Creates a buffer from bytes and assigns it to an entry in the tile shader
// argument table.
//
// bytes: A pointer to argument data the method copies to an [MTLBuffer] and assigns
// to an entry in the tile shader argument table for buffers.
//
// length: The number of bytes the method copies from the `bytes` pointer.
//
// index: An integer that represents the entry in the tile shader argument table for
// buffers that stores a record of the [MTLBuffer] the method creates from
// `bytes`.
//
// # Discussion
// 
// The method is equivalent to creating an [MTLBuffer] instance that contains
// the same data as `bytes` and calling the [setTileBuffer(_:offset:index:)]
// method. However, this method avoids the overhead of creating a buffer to
// store your data; instead, Metal manages the data.
// 
// For data that’s more than 4 KB, create an [MTLBuffer] instance and pass
// it to [setTileBuffer(_:offset:index:)].
// 
// By default, the buffer at each index is `nil`.
//
// [setTileBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBuffer(_:offset:index:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileBytes(_:length:index:)
func (o MTLRenderCommandEncoderObject) SetTileBytesLengthAtIndex(bytes []byte, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
// Assigns an intersection function table to an entry in the tile shader
// argument table.
//
// intersectionFunctionTable: An [MTLIntersectionFunctionTable] instance the command assigns to an entry
// in the tile shader argument table for intersection function tables.
//
// bufferIndex: An integer that represents the entry in the tile shader argument table for
// intersection function tables that stores a record of
// `intersectionFunctionTable`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileIntersectionFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetTileIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileIntersectionFunctionTable:atBufferIndex:"), intersectionFunctionTable, bufferIndex)
	}
// Assigns multiple intersection function tables to a range of entries in the
// tile shader argument table.
//
// intersectionFunctionTables: A pointer to a C array of [MTLIntersectionFunctionTable] instances the
// command assigns to entries in the tile shader argument table for
// intersection function tables.
//
// range: A span of integers that represent the entries in the tile shader argument
// table for intersection function tables. Each entry stores a record of the
// corresponding element in `intersectionFunctionTables`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileIntersectionFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetTileIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileIntersectionFunctionTables:withBufferRange:"), intersectionFunctionTables, range_)
	}
// Assigns a sampler state to an entry in the tile shader argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the tile
// shader argument table for sampler states.
//
// index: An integer that represents the entry in the tile shader argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerState(_:index:)
func (o MTLRenderCommandEncoderObject) SetTileSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileSamplerState:atIndex:"), sampler, index)
	}
// Assigns a sampler state and clamp values to an entry in the tile shader
// argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the tile
// shader argument table for sampler states.
//
// lodMinClamp: The smallest level of detail value a tile shader can use when it samples a
// texture.
//
// lodMaxClamp: The largest level of detail value a tile shader can use when it samples a
// texture.
//
// index: An integer that represents the entry in the tile shader argument table for
// sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// The method’s `lodMinClamp` and `lodMaxClamp` parameters override the
// default values for `sampler`. You can set the sampler’s default values by
// configuring the [LodMinClamp] and [LodMaxClamp] properties of
// [MTLSamplerDescriptor] before you create the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLRenderCommandEncoderObject) SetTileSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
	}
// Assigns multiple sampler states and clamp values to a range of entries in
// the tile shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the tile shader argument table for sampler states.
//
// lodMinClamps: A pointer to a C array of floating-point values. Each element is the
// smallest level of detail value a tile shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// lodMaxClamps: A pointer to a C array of floating-point values. Each element is the
// largest level of detail value a tile shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// range: A span of integers that represent the entries in the tile shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// Each element of the method’s `lodMinClamps` and `lodMaxClamps` parameters
// overrides the default values for the corresponding sampler in `samplers`.
// You can set a sampler’s default values by configuring the [LodMinClamp]
// and [LodMaxClamp] properties of [MTLSamplerDescriptor] before you create
// the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLRenderCommandEncoderObject) SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
	}
// Assigns multiple sampler states to a range of entries in the tile shader
// argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the tile shader argument table for sampler states.
//
// range: A span of integers that represent the entries in the tile shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileSamplerStates:withRange:
func (o MTLRenderCommandEncoderObject) SetTileSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileSamplerStates:withRange:"), samplers, range_)
	}
// Assigns a texture to an entry in the tile shader argument table.
//
// texture: An [MTLTexture] instance the command assigns to an entry in the tile shader
// argument table for textures.
//
// index: An integer that represents the entry in the tile shader argument table for
// textures that stores a record of `texture`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileTexture(_:index:)
func (o MTLRenderCommandEncoderObject) SetTileTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileTexture:atIndex:"), texture, index)
	}
// Assigns multiple textures to a range of entries in the tile shader argument
// table.
//
// textures: A pointer to a C array of [MTLTexture] instances the command assigns to
// entries in the tile shader argument table for textures.
//
// range: A span of integers that represent the entries in the tile shader argument
// table for textures. Each entry stores a record of the corresponding element
// in `textures`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileTextures:withRange:
func (o MTLRenderCommandEncoderObject) SetTileTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileTextures:withRange:"), textures, range_)
	}
// Assigns a visible function table to an entry in the tile shader argument
// table.
//
// functionTable: An [MTLVisibleFunctionTable] instance the command assigns to an entry in
// the tile shader argument table for visible function tables.
//
// bufferIndex: An integer that represents the entry in the tile shader argument table for
// visible function tables that stores a record of `functionTable`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileVisibleFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetTileVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileVisibleFunctionTable:atBufferIndex:"), functionTable, bufferIndex)
	}
// Assigns multiple visible function tables to a range of entries in the tile
// shader argument table.
//
// functionTables: A pointer to a C array of [MTLVisibleFunctionTable] instances the command
// assigns to entries in the tile shader argument table for visible function
// tables.
//
// range: A span of integers that represent the entries in the tile shader argument
// table for visible function tables. Each entry stores a record of the
// corresponding element in `functionTables`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTileVisibleFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetTileVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTileVisibleFunctionTables:withBufferRange:"), functionTables, range_)
	}
// Configures how subsequent draw commands rasterize triangle and triangle
// strip primitives.
//
// fillMode: A triangle filling mode the render pass applies to draw commands that
// rasterize triangles or triangle strips.
//
// # Discussion
// 
// The render pass’s default mode is [TriangleFillModeFill].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setTriangleFillMode(_:)
func (o MTLRenderCommandEncoderObject) SetTriangleFillMode(fillMode MTLTriangleFillMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTriangleFillMode:"), fillMode)
	}
// Assigns an acceleration structure to an entry in the vertex shader argument
// table.
//
// accelerationStructure: An [MTLAccelerationStructure] instance the command assigns to an entry in
// the vertex shader argument table for acceleration structures.
//
// bufferIndex: An integer that represents the entry in the vertex shader argument table
// for acceleration structures that stores a record of
// `accelerationStructure`.
//
// # Discussion
// 
// By default, the acceleration structure at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexAccelerationStructure(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetVertexAccelerationStructureAtBufferIndex(accelerationStructure MTLAccelerationStructure, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexAccelerationStructure:atBufferIndex:"), accelerationStructure, bufferIndex)
	}
// Configures the number of output vertices the render pipeline produces for
// each input vertex, optionally with render target and viewport offsets.
//
// count: The number of outputs to create.
//
// viewMappings: An optional pointer to a C array that has at least `count`
// [MTLVertexAmplificationViewMapping] elements. Each element in the array
// provides per-output offsets to a specific render target and viewport.
// //
// [MTLVertexAmplificationViewMapping]: https://developer.apple.com/documentation/Metal/MTLVertexAmplificationViewMapping
//
// # Discussion
// 
// With , you can encode drawing commands that process the same vertex
// multiple times, one per render target. You can configure the render
// pipeline’s vertex amplification multiplier by calling this method with a
// `count` argument that’s greater than `1`.
// 
// For more information about vertex amplification and how to use the
// `viewMappings` parameter, see [Improving rendering performance with vertex
// amplification].
//
// [Improving rendering performance with vertex amplification]: https://developer.apple.com/documentation/Metal/improving-rendering-performance-with-vertex-amplification
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexAmplificationCount(_:viewMappings:)
func (o MTLRenderCommandEncoderObject) SetVertexAmplificationCountViewMappings(count uint, viewMappings *MTLVertexAmplificationViewMapping) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexAmplificationCount:viewMappings:"), count, viewMappings)
	}
// Updates an entry in the vertex shader argument table with a new location
// within the entry’s current buffer.
//
// offset: An integer that represents the location, in bytes, from the start of
// `buffer` where the vertex shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// index: An integer that represents the entry in the vertex shader argument table
// for buffers that already stores a record of an [MTLBuffer].
//
// # Discussion
// 
// The command this method encodes changes the offset for a fragment buffer
// that already has a previous assignment from one of your earlier commands.
// 
// For more information, see:
// 
// - [setVertexBuffer(_:offset:index:)] - [setVertexBuffers(_:offsets:range:)]
// (Swift) - [SetVertexBuffersOffsetsWithRange] (Objective-C)
// 
// The command can also adjust the offset for an entry that you previously set
// with the [SetVertexBytesLengthAtIndex] method.
// 
// By default, the buffer at each index is `nil`.
//
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
// [setVertexBuffers(_:offsets:range:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffers(_:offsets:range:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBufferOffset(_:index:)
func (o MTLRenderCommandEncoderObject) SetVertexBufferOffsetAtIndex(offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBufferOffset:atIndex:"), offset, index)
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBufferOffset(offset:attributeStride:index:)
func (o MTLRenderCommandEncoderObject) SetVertexBufferOffsetAttributeStrideAtIndex(offset uint, stride uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBufferOffset:attributeStride:atIndex:"), offset, stride, index)
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffers:offsets:attributeStrides:withRange:
func (o MTLRenderCommandEncoderObject) SetVertexBuffersOffsetsAttributeStridesWithRange(buffers []MTLBuffer, offsets uint, strides uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBuffers:offsets:attributeStrides:withRange:"), buffers, offsets, strides, range_)
	}
// Assigns multiple buffers to a range of entries in the vertex shader
// argument table.
//
// buffers: A pointer to a C array of [MTLBuffer] instances the command assigns to
// entries in the vertex shader argument table for buffers.
//
// offsets: A pointer to a C array of unsigned integers. Each element represents the
// location, in bytes, from the start of the corresponding [MTLBuffer] element
// in `buffers` where the vertex shader argument data begins.
// 
// See the [Metal feature set tables (PDF)] to check for offset alignment
// requirements for buffers in `device` and `constant` address space.
// //
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for buffers. Each entry stores a record of the corresponding element
// in `buffers` and `offsets`.
//
// # Discussion
// 
// By default, the buffer at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffers:offsets:withRange:
func (o MTLRenderCommandEncoderObject) SetVertexBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Creates a buffer from bytes and assigns it to an entry in the vertex shader
// argument table.
//
// bytes: A pointer to argument data the method copies to an [MTLBuffer] and assigns
// to an entry in the vertex shader argument table for buffers.
//
// length: The number of bytes the method copies from the `bytes` pointer.
//
// index: An integer that represents the entry in the vertex shader argument table
// for buffers that stores a record of the [MTLBuffer] the method creates from
// `bytes`.
//
// # Discussion
// 
// The method is equivalent to creating an [MTLBuffer] instance that contains
// the same data as `bytes` and calling the [setVertexBuffer(_:offset:index:)]
// method. However, this method avoids the overhead of creating a buffer to
// store your data; instead, Metal manages the data.
// 
// For data that’s more than 4 KB, create an [MTLBuffer] instance and pass
// it to [setVertexBuffer(_:offset:index:)].
// 
// By default, the buffer at each index is `nil`.
//
// [setVertexBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBuffer(_:offset:index:)
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBytes(_:length:index:)
func (o MTLRenderCommandEncoderObject) SetVertexBytesLengthAtIndex(bytes []byte, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBytes:length:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), index)
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexBytes(_:length:attributeStride:index:)
func (o MTLRenderCommandEncoderObject) SetVertexBytesLengthAttributeStrideAtIndex(bytes []byte, stride uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBytes:length:attributeStride:atIndex:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)), stride, index)
	}
// Assigns an intersection function table to an entry in the vertex shader
// argument table.
//
// intersectionFunctionTable: An [MTLIntersectionFunctionTable] instance the command assigns to an entry
// in the vertex shader argument table for intersection function tables.
//
// bufferIndex: An integer that represents the entry in the vertex shader argument table
// for intersection function tables that stores a record of
// `intersectionFunctionTable`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexIntersectionFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetVertexIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable MTLIntersectionFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexIntersectionFunctionTable:atBufferIndex:"), intersectionFunctionTable, bufferIndex)
	}
// Assigns multiple intersection function tables to a range of entries in the
// vertex shader argument table.
//
// intersectionFunctionTables: A pointer to a C array of [MTLIntersectionFunctionTable] instances the
// command assigns to entries in the vertex shader argument table for
// intersection function tables.
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for intersection function tables. Each entry stores a record of the
// corresponding element in `intersectionFunctionTables`.
//
// # Discussion
// 
// By default, the intersection function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexIntersectionFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetVertexIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexIntersectionFunctionTables:withBufferRange:"), intersectionFunctionTables, range_)
	}
// Assigns a sampler state to an entry in the vertex shader argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the vertex
// shader argument table for sampler states.
//
// index: An integer that represents the entry in the vertex shader argument table
// for sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerState(_:index:)
func (o MTLRenderCommandEncoderObject) SetVertexSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexSamplerState:atIndex:"), sampler, index)
	}
// Assigns a sampler state and clamp values to an entry in the vertex shader
// argument table.
//
// sampler: An [MTLSamplerState] instance the command assigns to an entry in the vertex
// shader argument table for sampler states.
//
// lodMinClamp: The smallest level of detail value a vertex shader can use when it samples
// a texture.
//
// lodMaxClamp: The largest level of detail value a vertex shader can use when it samples a
// texture.
//
// index: An integer that represents the entry in the vertex shader argument table
// for sampler states that stores a record of `sampler`.
//
// # Discussion
// 
// The method’s `lodMinClamp` and `lodMaxClamp` parameters override the
// default values for `sampler`. You can set the sampler’s default values by
// configuring the [LodMinClamp] and [LodMaxClamp] properties of
// [MTLSamplerDescriptor] before you create the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerState(_:lodMinClamp:lodMaxClamp:index:)
func (o MTLRenderCommandEncoderObject) SetVertexSamplerStateLodMinClampLodMaxClampAtIndex(sampler MTLSamplerState, lodMinClamp float32, lodMaxClamp float32, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), sampler, lodMinClamp, lodMaxClamp, index)
	}
// Assigns multiple sampler states and clamp values to a range of entries in
// the vertex shader argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the vertex shader argument table for sampler states.
//
// lodMinClamps: A pointer to a C array of floating-point values. Each element is the
// smallest level of detail value a vertex shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// lodMaxClamps: A pointer to a C array of floating-point values. Each element is the
// largest level of detail value a vertex shader can use when it samples a
// texture with the corresponding element in `samplers`.
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// Each element of the method’s `lodMinClamps` and `lodMaxClamps` parameters
// overrides the default values for the corresponding sampler in `samplers`.
// You can set a sampler’s default values by configuring the [LodMinClamp]
// and [LodMaxClamp] properties of [MTLSamplerDescriptor] before you create
// the sampler.
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerStates:lodMinClamps:lodMaxClamps:withRange:
func (o MTLRenderCommandEncoderObject) SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []MTLSamplerState, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), samplers, lodMinClamps, lodMaxClamps, range_)
	}
// Assigns multiple sampler states to a range of entries in the vertex shader
// argument table.
//
// samplers: A pointer to a C array of [MTLSamplerState] instances the command assigns
// to entries in the vertex shader argument table for sampler states.
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for sampler states. Each entry stores a record of the corresponding
// element in `samplers`.
//
// # Discussion
// 
// By default, the sampler state at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexSamplerStates:withRange:
func (o MTLRenderCommandEncoderObject) SetVertexSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexSamplerStates:withRange:"), samplers, range_)
	}
// Assigns a texture to an entry in the vertex shader argument table.
//
// texture: An [MTLTexture] instance the command assigns to an entry in the vertex
// shader argument table for textures.
//
// index: An integer that represents the entry in the vertex shader argument table
// for textures that stores a record of `texture`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexTexture(_:index:)
func (o MTLRenderCommandEncoderObject) SetVertexTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexTexture:atIndex:"), texture, index)
	}
// Assigns multiple textures to a range of entries in the vertex shader
// argument table.
//
// textures: A pointer to a C array of [MTLTexture] instances the command assigns to
// entries in the vertex shader argument table for textures.
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for textures. Each entry stores a record of the corresponding element
// in `textures`.
//
// # Discussion
// 
// By default, the texture at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexTextures:withRange:
func (o MTLRenderCommandEncoderObject) SetVertexTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexTextures:withRange:"), textures, range_)
	}
// Assigns a visible function table to an entry in the vertex shader argument
// table.
//
// functionTable: An [MTLVisibleFunctionTable] instance the command assigns to an entry in
// the vertex shader argument table for visible function tables.
//
// bufferIndex: An integer that represents the entry in the vertex shader argument table
// for visible function tables that stores a record of `functionTable`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexVisibleFunctionTable(_:bufferIndex:)
func (o MTLRenderCommandEncoderObject) SetVertexVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexVisibleFunctionTable:atBufferIndex:"), functionTable, bufferIndex)
	}
// Assigns multiple visible function tables to a range of entries in the
// vertex shader argument table.
//
// functionTables: A pointer to a C array of [MTLVisibleFunctionTable] instances the command
// assigns to entries in the vertex shader argument table for visible function
// tables.
//
// range: A span of integers that represent the entries in the vertex shader argument
// table for visible function tables. Each entry stores a record of the
// corresponding element in `functionTables`.
//
// # Discussion
// 
// By default, the visible function table at each index is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVertexVisibleFunctionTables:withBufferRange:
func (o MTLRenderCommandEncoderObject) SetVertexVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexVisibleFunctionTables:withBufferRange:"), functionTables, range_)
	}
// Configures the render pipeline with a viewport that applies a
// transformation and a clipping rectangle.
//
// viewport: An [MTLViewport] instance the command applies to the render pipeline for
// transformations and clipping.
// //
// [MTLViewport]: https://developer.apple.com/documentation/Metal/MTLViewport
//
// # Discussion
// 
// The render pipeline linearly maps vertex positions from normalized device
// coordinates to viewport coordinates by applying a viewport during the
// rasterization stage. It applies the transform first and then rasterizes the
// primitive while clipping any fragments outside the scissor rectangle (see
// [SetScissorRect]) or the render target’s extents.
// 
// The viewport’s [originX] and [originY] properties, which default to
// `0.0`, represent the number of pixels from the top-left corner of the
// render target. Positive [originX] values go to the right and positive
// [originY] values go downward. The default values for its [width] and
// [height] properties are the render target’s width and height,
// respectively. The default values for its [znear] and [zfar] properties are
// `0.0` and `1.0`, respectively, which you can flip.
//
// [height]: https://developer.apple.com/documentation/Metal/MTLViewport/height
// [originX]: https://developer.apple.com/documentation/Metal/MTLViewport/originX
// [originY]: https://developer.apple.com/documentation/Metal/MTLViewport/originY
// [width]: https://developer.apple.com/documentation/Metal/MTLViewport/width
// [zfar]: https://developer.apple.com/documentation/Metal/MTLViewport/zfar
// [znear]: https://developer.apple.com/documentation/Metal/MTLViewport/znear
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setViewport(_:)
func (o MTLRenderCommandEncoderObject) SetViewport(viewport MTLViewport) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setViewport:"), viewport)
	}
// Configures the render pipeline with multiple viewports that apply
// transformations and clipping rectangles.
//
// viewports: An array of [MTLViewport] instances the command applies to the render
// pipeline for transformations and clipping.
// //
// [MTLViewport]: https://developer.apple.com/documentation/Metal/MTLViewport
//
// count: The number of elements in the `viewports` array.
//
// # Discussion
// 
// Use this method to configure multiple active viewports and corresponding
// scissor rectangles. Multiple viewports give your app the ability to draw
// into separate areas of an image with a single draw call. You can either set
// a single scissor rectangle for all viewports with the [SetScissorRect]
// method, or set each viewport’s rectangle with the [setScissorRects(_:)]
// method.
// 
// The maximum number of viewports and scissor rectangles a GPU supports
// varies by device family. For more information, see [MTLGPUFamily] and
// [Detecting GPU features and Metal software versions].
// 
// The rendering pipeline sends each primitive to a single viewport and its
// associated scissor rectangle. You can select which viewport each primitive
// uses in your vertex shader by adding the `[[viewport_array_index]]`
// attribute to an output value.
// 
// The render pipeline linearly maps vertex positions from normalized device
// coordinates to viewport coordinates by applying a viewport during the
// rasterization stage. It applies the transform first and then rasterizes the
// primitive while clipping any fragments outside the scissor rectangle (see
// [SetScissorRect]) or the render target’s extents.
// 
// The viewport’s [originX] and [originY] properties, which default to
// `0.0`, represent the number of pixels from the top-left corner of the
// render target. Positive [originX] values go to the right and positive
// [originY] values go downward. The default values for its [width] and
// [height] properties are the render target’s width and height,
// respectively. The default values for its [znear] and [zfar] properties are
// `0.0` and `1.0`, respectively, which you can flip.
// 
// The [SetViewport] method is equivalent to calling this method with a single
// viewport element in the `viewports` array.
//
// [Detecting GPU features and Metal software versions]: https://developer.apple.com/documentation/Metal/detecting-gpu-features-and-metal-software-versions
// [MTLGPUFamily]: https://developer.apple.com/documentation/Metal/MTLGPUFamily
// [height]: https://developer.apple.com/documentation/Metal/MTLViewport/height
// [originX]: https://developer.apple.com/documentation/Metal/MTLViewport/originX
// [originY]: https://developer.apple.com/documentation/Metal/MTLViewport/originY
// [setScissorRects(_:)]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setScissorRects(_:)
// [width]: https://developer.apple.com/documentation/Metal/MTLViewport/width
// [zfar]: https://developer.apple.com/documentation/Metal/MTLViewport/zfar
// [znear]: https://developer.apple.com/documentation/Metal/MTLViewport/znear
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setViewports:count:
func (o MTLRenderCommandEncoderObject) SetViewportsCount(viewports []MTLViewport, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setViewports:count:"), objc.CArray(viewports), count)
	}
// Configures which visibility test the GPU runs and the destination for any
// results it generates.
//
// mode: An [MTLVisibilityResultMode] that configures which visibility test results
// the render pass saves to a buffer, or disables visibility testing.
// //
// [MTLVisibilityResultMode]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode
//
// offset: A location, in bytes, relative to the start of the render pass’s
// [VisibilityResultBuffer]. The GPU stores the result of a visibility test at
// `offset`, which needs to be a multiple of 8.
//
// # Discussion
// 
// To create a render pass that can enable visibility testing, assign an
// [MTLBuffer] instance to the [VisibilityResultBuffer] property of an
// [MTLRenderPassDescriptor].
// 
// You can monitor one or more drawing commands with a visibility test by
// calling this method before the drawing commands. The encoder uses the new
// visibility mode and offset for subsequent drawing commands until you change
// the configuration by calling the method again. For example, you can change
// the offset or entirely disable visibility tests for subsequent commands by
// passing [VisibilityResultModeDisabled].
// 
// The default mode for a render pass is [VisibilityResultModeDisabled].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/setVisibilityResultMode(_:offset:)
func (o MTLRenderCommandEncoderObject) SetVisibilityResultModeOffset(mode MTLVisibilityResultMode, offset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVisibilityResultMode:offset:"), mode, offset)
	}
// Ensures the shaders in the render pass’s subsequent draw commands have
// access to the resources you allocate from a heap.
//
// heap: An [MTLHeap] instance with resources that subsequent draw commands depend
// on.
//
// stages: All the render stages that depend on resources within `heap`, including
// [RenderStageObject], [RenderStageMesh], [RenderStageVertex],
// [RenderStageFragment], and [RenderStageTile].
//
// # Discussion
// 
// You can make the resources in `heap` (available in GPU memory) for the
// remaining duration of the render pass by calling this method. Call the
// method before encoding draw calls that may access resources within `heap`
// through an argument buffer. The method ensures each resource is in a format
// that’s compatible with the shaders that depend on it.
// 
// The method’s applies the [ResourceUsageRead] resource usage option to all
// of the resources within `heap`, except for textures. The method ignores any
// texture that has [TextureUsageRenderTarget], [TextureUsageShaderWrite], or
// both in its [Usage] property. For all other textures in `heap`, the method
// optimizes each texture’s memory layout for rendering with a sampler.
// However, your shaders can’t read from those textures by calling this
// method because the texture needs a different memory layout that’s
// suitable for reading.
// 
// Methods that apply a usage option for resources (see [Argument buffer
// resource preparation commands]) override any previous calls that apply to a
// resource. For example, you can change the usage option for buffer in `heap`
// to [ResourceUsageWrite] by passing it to [UseResourceUsageStages] after
// calling this method. However, you can’t reverse the call order because
// this method resets the usage for all resources within `heap` to
// [ResourceUsageRead], overriding previous calls to [UseResourceUsageStages].
// 
// The method instructs Metal to apply hazard tracking for resources you
// allocate from a heap that you create with [HazardTrackingModeTracked].
// However, for untracked resources — which come from heaps you create with
// [HazardTrackingModeUntracked] — you need to account for hazards by
// applying [MTLFence] or [MTLEvent] instances.
// 
// Apps typically call the method for heaps that have resources in argument
// buffers for a implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// [Argument buffer resource preparation commands]: https://developer.apple.com/documentation/Metal/argument-buffer-resource-preparation-commands
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useHeap(_:stages:)
func (o MTLRenderCommandEncoderObject) UseHeapStages(heap MTLHeap, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useHeap:stages:"), heap, stages)
	}
// Ensures the shaders in the render pass’s subsequent draw commands have
// access to the resources you allocate from multiple heaps.
//
// heaps: A pointer to a C array of [MTLHeap] instances with resources that
// subsequent draw commands depend on.
//
// count: The number of elements in `heaps`.
//
// stages: All the render stages that depend on resources within `heaps`, including
// [RenderStageObject], [RenderStageMesh], [RenderStageVertex],
// [RenderStageFragment], and [RenderStageTile].
//
// # Discussion
// 
// You can make the resources in `heaps` (available in GPU memory) for the
// remaining duration of the render pass by calling this method. Call the
// method before encoding draw calls that may access resources within `heaps`
// through an argument buffer. The method ensures each resource is in a format
// that’s compatible with the shaders that depend on it.
// 
// The method’s applies the [ResourceUsageRead] resource usage option to all
// of the resources within `heaps`, except for textures. The method ignores
// any texture that has [TextureUsageRenderTarget], [TextureUsageShaderWrite],
// or both in its [Usage] property. For all other textures in `heaps`, the
// method optimizes each texture’s memory layout for rendering with a
// sampler. However, your shaders can’t read from those textures by calling
// this method because the texture needs a different memory layout that’s
// suitable for reading.
// 
// Methods that apply a usage option for resources (see [Argument buffer
// resource preparation commands]) override any previous calls that apply to a
// resource. For example, you can change the usage option for a buffer to
// [ResourceUsageWrite] by passing it to [UseResourceUsageStages] after
// calling this method. However, you can’t reverse the call order because
// this method resets the usage for all resources within `heaps` to
// [ResourceUsageRead], overriding previous calls to [UseResourceUsageStages].
// 
// The method instructs Metal to apply hazard tracking for resources you
// allocate from a heap that you create with [HazardTrackingModeTracked].
// However, for untracked resources — which come from heaps you create with
// [HazardTrackingModeUntracked] — you need to account for hazards by
// applying [MTLFence] or [MTLEvent] instances.
// 
// Apps typically call the method for heaps that have resources in argument
// buffers for a implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// [Argument buffer resource preparation commands]: https://developer.apple.com/documentation/Metal/argument-buffer-resource-preparation-commands
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useHeaps:count:stages:
func (o MTLRenderCommandEncoderObject) UseHeapsCountStages(heaps []MTLHeap, count uint, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useHeaps:count:stages:"), objc.CArray(heaps), count, stages)
	}
// Ensures the shaders in the render pass’s subsequent draw commands have
// access to a resource.
//
// resource: An [MTLResource] instance that subsequent draw commands depend on.
//
// usage: All the applicable access types the render pass’s shaders use for the
// resource, including [ResourceUsageRead] and [ResourceUsageWrite].
// 
// For applicable resources, you may be able to prevent the GPU from
// unnecessarily decompressing color attachments on some devices by setting
// `usage` to [ResourceUsageRead].
//
// stages: All the render stages that depend on `resource`, including
// [RenderStageObject], [RenderStageMesh], [RenderStageVertex],
// [RenderStageFragment], and [RenderStageTile].
//
// # Discussion
// 
// You can make a resource (available in GPU memory) for the remaining
// duration of the render pass by calling this method. Call the method before
// encoding draw calls that may access `resource` through an argument buffer.
// The method ensures the resource is in a format that’s compatible with the
// shaders that depend on it.
// 
// For example, you can explicitly bind resources for the vertex stage with
// the methods in the [Vertex shader resource preparation commands]
// collection.
// 
// The method also informs Metal when to apply hazard tracking for a resource
// you create with [HazardTrackingModeTracked]. For a resource you create with
// [HazardTrackingModeUntracked], you need to apply an [MTLFence] or an
// [MTLEvent] to account for potential reading and writing hazards.
// 
// You can reconfigure an individual resource’s `usage` options for
// subsequent draw calls in the same render pass by calling this method again.
// 
// Apps typically call the method for a resource in an argument buffer as a
// part of their implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
// [Vertex shader resource preparation commands]: https://developer.apple.com/documentation/Metal/vertex-shader-resource-preparation-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useResource(_:usage:stages:)
func (o MTLRenderCommandEncoderObject) UseResourceUsageStages(resource MTLResource, usage MTLResourceUsage, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResource:usage:stages:"), resource, usage, stages)
	}
// Ensures the shaders in the render pass’s subsequent draw commands have
// access to multiple resources.
//
// resources: A pointer to a C array of [MTLResource] instances that subsequent draw
// commands depend on.
//
// count: The number of elements in `resources`.
//
// usage: All the applicable access types the render pass’s shaders use for
// `resource`, including [ResourceUsageRead] and [ResourceUsageWrite].
// 
// For applicable resources, you may be able to prevent the GPU from
// unnecessarily decompressing color attachments on some devices by setting
// `usage` to [ResourceUsageRead].
//
// stages: All the render stages that depend on elements in `resources`, including
// [RenderStageObject], [RenderStageMesh], [RenderStageVertex],
// [RenderStageFragment], and [RenderStageTile].
//
// # Discussion
// 
// You can make multiple resources (available in GPU memory) for the remaining
// duration of the render pass by calling this method. Call the method before
// encoding draw calls that may access the elements of `resources` through an
// argument buffer. The method ensures each resource is in a format that’s
// compatible with the shaders that depend on it.
// 
// For example, you can explicitly bind resources for the vertex stage with
// the methods in the [Vertex shader resource preparation commands]
// collection.
// 
// The method also informs Metal when to apply hazard tracking for the
// resources you create with [HazardTrackingModeTracked]. For resources you
// create with [HazardTrackingModeUntracked], you need to apply an [MTLFence]
// or an [MTLEvent] to account for potential reading and writing hazards.
// 
// You can reconfigure an individual resource’s `usage` options for
// subsequent draw calls in the same render pass by calling this method again.
// 
// Apps typically call the method for resources in an argument buffer as a
// part of their implementation. For more information about argument buffers
// and bindless implementations, see [Improving CPU performance by using
// argument buffers] and [Go bindless with Metal 3], respectively.
//
// [Go bindless with Metal 3]: https://developer.apple.com/videos/play/wwdc2022/10101/
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
// [Vertex shader resource preparation commands]: https://developer.apple.com/documentation/Metal/vertex-shader-resource-preparation-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder/useResources:count:usage:stages:
func (o MTLRenderCommandEncoderObject) UseResourcesCountUsageStages(resources []MTLResource, count uint, usage MTLResourceUsage, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResources:count:usage:stages:"), objc.CArray(resources), count, usage, stages)
	}
// Declares that all command generation from the encoder is completed.
//
// # Discussion
// 
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLRenderCommandEncoderObject) EndEncoding() {
	
	objc.Send[struct{}](o.ID, objc.Sel("endEncoding"))
	}
// Inserts a debug string into the captured frame data.
//
// # Discussion
// 
// For more information, see [Naming resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/insertDebugSignpost(_:)
func (o MTLRenderCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
	}
// Pushes a specific string onto a stack of debug group strings for the
// command encoder.
//
// # Discussion
// 
// For more information, see [Naming resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/pushDebugGroup(_:)
func (o MTLRenderCommandEncoderObject) PushDebugGroup(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
	}
// Pops the latest string off of a stack of debug group strings for the
// command encoder.
//
// # Discussion
// 
// For more information, see [Naming resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/popDebugGroup()
func (o MTLRenderCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}
// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLRenderCommandEncoderObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLRenderCommandEncoderObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Encodes a consumer barrier on work you commit to the same command queue.
//
// afterQueueStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in prior
// command encoders, and not for the current encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeStages: [MTLStages] mask that represents the stages of work that wait. This
// argument applies to work you encode in the current command encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
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
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/barrier(afterQueueStages:beforeStages:)
func (o MTLRenderCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
	}

func (o MTLRenderCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

