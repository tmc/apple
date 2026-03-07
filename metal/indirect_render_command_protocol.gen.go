// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A render command in an indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand
type MTLIndirectRenderCommand interface {
	objectivec.IObject

	// Sets the render pipeline state for the command.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setRenderPipelineState(_:)
	SetRenderPipelineState(pipelineState MTLRenderPipelineState)

	// Sets a vertex buffer argument for the command.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setVertexBuffer(_:offset:at:)
	SetVertexBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Sets a fragment buffer argument for the command.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setFragmentBuffer(_:offset:at:)
	SetFragmentBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Encodes a command to render a number of instances of primitives using vertex data in contiguous array elements, starting from the base instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawPrimitives(_:vertexStart:vertexCount:instanceCount:baseInstance:)
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)

	// Encodes a command to render a number of instances of primitives using an index list specified in a buffer, starting from the base vertex of the base instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawIndexedPrimitives(_:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)

	// Encodes a command to render a number of instances of tessellated patches.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawPatches(_:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:)
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer MTLBuffer, offset uint, instanceStride uint)

	// Resets the command to its default state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/reset()
	Reset()

	// ClearBarrier protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/clearBarrier()
	ClearBarrier()

	// DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawMeshThreadgroups(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawMeshThreads(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// SetBarrier protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setBarrier()
	SetBarrier()

	// SetCullMode protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setCullMode(_:)
	SetCullMode(cullMode MTLCullMode)

	// SetDepthBiasSlopeScaleClamp protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthBias(_:slopeScale:clamp:)
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)

	// SetDepthClipMode protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthClipMode(_:)
	SetDepthClipMode(depthClipMode MTLDepthClipMode)

	// SetDepthStencilState protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthStencilState(_:)
	SetDepthStencilState(depthStencilState MTLDepthStencilState)

	// SetFrontFacingWinding protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setFrontFacing(_:)
	SetFrontFacingWinding(frontFacingWindning MTLWinding)

	// SetMeshBufferOffsetAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setMeshBuffer(_:offset:at:)
	SetMeshBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// SetObjectBufferOffsetAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setObjectBuffer(_:offset:at:)
	SetObjectBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// SetObjectThreadgroupMemoryLengthAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setObjectThreadgroupMemoryLength(_:index:)
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// SetTriangleFillMode protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setTriangleFillMode(_:)
	SetTriangleFillMode(fillMode MTLTriangleFillMode)

	// SetVertexBufferOffsetAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setVertexBuffer(_:offset:attributeStride:at:)
	SetVertexBufferOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint)

	// Encodes a command to render a number of instances of tessellated patches, using a control point index buffer.
	//
	// See: /documentation/metal/mtlindirectrendercommand/drawindexedpatches(_:patchstart:patchcount:patchindexbuffer:patchindexbufferoffset:controlpointindexbuffer:controlpointindexbufferoffset:instancecount:baseinstance:tessellationfactorbuffer:tessellationfactorbufferoffset:tessellationfactorbu-4mdz8
	DrawIndexedPatches()
}



// MTLIndirectRenderCommandObject wraps an existing Objective-C object that conforms to the MTLIndirectRenderCommand protocol.
type MTLIndirectRenderCommandObject struct {
	objectivec.Object
}
func (o MTLIndirectRenderCommandObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLIndirectRenderCommandObjectFromID constructs a [MTLIndirectRenderCommandObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIndirectRenderCommandObjectFromID(id objc.ID) MTLIndirectRenderCommandObject {
	return MTLIndirectRenderCommandObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Sets the render pipeline state for the command.
//
// pipelineState: The rendering pipeline state object to use.
//
// # Discussion
// 
// You don’t need to call this method if you create an indirect command
// buffer with its [InheritPipelineState] property equal to [true]. The
// command gets the pipeline state from the parent encoder when it runs.
// 
// If you created the indirect command buffer with [InheritPipelineState] set
// to [false], you need to set the pipeline state prior to encoding the
// drawing command.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setRenderPipelineState(_:)

func (o MTLIndirectRenderCommandObject) SetRenderPipelineState(pipelineState MTLRenderPipelineState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setRenderPipelineState:"), pipelineState)
	}

// Sets a vertex buffer argument for the command.
//
// buffer: The buffer to set in the buffer argument table.
//
// offset: The location, in bytes relative to start of `buffer`, of the first byte of
// data for the vertex shader.
//
// index: An index in the buffer argument table. The maximum index is determined when
// you created the indirect command buffer.
//
// # Discussion
// 
// You don’t need to call this method if you create an indirect command
// buffer with its [InheritBuffers] property equal to [true]. The command gets
// the arguments from the parent encoder when it runs.
// 
// If you need to pass other kinds of parameters to your shader, such as
// textures and samplers, create an argument buffer and pass it to the shader
// using this method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setVertexBuffer(_:offset:at:)

func (o MTLIndirectRenderCommandObject) SetVertexBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBuffer:offset:atIndex:"), buffer, offset, index)
	}

// Sets a fragment buffer argument for the command.
//
// buffer: The buffer to set in the buffer argument table.
//
// offset: The location, in bytes relative to start of `buffer`, of the first byte of
// data for the fragment shader.
//
// index: An index in the buffer argument table. The maximum index is determined when
// you created the indirect command buffer.
//
// # Discussion
// 
// You don’t need to call this method if you create an indirect command
// buffer with its [InheritBuffers] equal to [true]. The command gets the
// arguments from the parent encoder when it runs.
// 
// If you need to pass other kinds of parameters to your shader, such as
// textures and samplers, create an argument buffer and pass it to the shader
// using this method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setFragmentBuffer(_:offset:at:)

func (o MTLIndirectRenderCommandObject) SetFragmentBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFragmentBuffer:offset:atIndex:"), buffer, offset, index)
	}

// Encodes a command to render a number of instances of primitives using
// vertex data in contiguous array elements, starting from the base instance.
//
// primitiveType: The type of primitives that the vertices are assembled into.
//
// vertexStart: The first vertex to draw.
//
// vertexCount: The number of vertices to draw.
//
// instanceCount: The number of instances to draw.
//
// baseInstance: The first instance to draw.
//
// # Discussion
// 
// The command generated by this method is equivalent to calling
// [DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance].
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawPrimitives(_:vertexStart:vertexCount:instanceCount:baseInstance:)

func (o MTLIndirectRenderCommandObject) DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:baseInstance:"), primitiveType, vertexStart, vertexCount, instanceCount, baseInstance)
	}

// Encodes a command to render a number of instances of primitives using an
// index list specified in a buffer, starting from the base vertex of the base
// instance.
//
// primitiveType: The type of primitive that the vertices are assembled into.
//
// indexCount: For each instance, the number of indices to read from the index buffer.
//
// indexType: The data type of the indices.
//
// indexBuffer: A buffer that contains indices to vertices.
//
// indexBufferOffset: Byte offset within indexBuffer to start reading indices from.
//
// instanceCount: The number of instances to draw.
//
// baseVertex: The first vertex to draw.
//
// baseInstance: The first instance to draw.
//
// # Discussion
// 
// The command generated by this method is equivalent to calling
// [DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance].
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawIndexedPrimitives(_:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:)

func (o MTLIndirectRenderCommandObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLBuffer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferOffset, instanceCount, baseVertex, baseInstance)
	}

// Encodes a command to render a number of instances of tessellated patches.
//
// numberOfPatchControlPoints: The number of control points per patch. This value needs to be between `0`
// and `32`, inclusive.
//
// patchStart: The first patch to draw.
//
// patchCount: The number of patches in each instance.
//
// patchIndexBuffer: A buffer that contains indices to patches.
//
// patchIndexBufferOffset: The byte offset within `patchIndexBuffer` to start reading indices from.
//
// instanceCount: The number of instances to draw.
//
// baseInstance: The first instance to draw.
//
// buffer: The per-patch tessellation factors buffer.
//
// offset: The distance, in bytes, between the start of the data and the start of the
// buffer. This value needs to be a multiple of `4` bytes.
//
// instanceStride: The distance, in bytes, between two instances of data in the buffer.
//
// # Discussion
// 
// The command generated by this method is equivalent to calling
// [SetTessellationFactorBufferOffsetInstanceStride] followed by
// [DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstance].
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawPatches(_:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:)

func (o MTLIndirectRenderCommandObject) DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer MTLBuffer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer MTLBuffer, offset uint, instanceStride uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:"), numberOfPatchControlPoints, patchStart, patchCount, patchIndexBuffer, patchIndexBufferOffset, instanceCount, baseInstance, buffer, offset, instanceStride)
	}

// Resets the command to its default state.
//
// # Discussion
// 
// A command that has been reset loses any state that you previously set and
// does nothing when executed.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/reset()

func (o MTLIndirectRenderCommandObject) Reset() {
	
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
	}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/clearBarrier()

func (o MTLIndirectRenderCommandObject) ClearBarrier() {
	
	objc.Send[struct{}](o.ID, objc.Sel("clearBarrier"))
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawMeshThreadgroups(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)

func (o MTLIndirectRenderCommandObject) DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreadgroups:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadgroupsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/drawMeshThreads(_:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)

func (o MTLIndirectRenderCommandObject) DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreads:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setBarrier()

func (o MTLIndirectRenderCommandObject) SetBarrier() {
	
	objc.Send[struct{}](o.ID, objc.Sel("setBarrier"))
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setCullMode(_:)

func (o MTLIndirectRenderCommandObject) SetCullMode(cullMode MTLCullMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setCullMode:"), cullMode)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthBias(_:slopeScale:clamp:)

func (o MTLIndirectRenderCommandObject) SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthBias:slopeScale:clamp:"), depthBias, slopeScale, clamp)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthClipMode(_:)

func (o MTLIndirectRenderCommandObject) SetDepthClipMode(depthClipMode MTLDepthClipMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthClipMode:"), depthClipMode)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setDepthStencilState(_:)

func (o MTLIndirectRenderCommandObject) SetDepthStencilState(depthStencilState MTLDepthStencilState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStencilState:"), depthStencilState)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setFrontFacing(_:)

func (o MTLIndirectRenderCommandObject) SetFrontFacingWinding(frontFacingWindning MTLWinding) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFrontFacingWinding:"), frontFacingWindning)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setMeshBuffer(_:offset:at:)

func (o MTLIndirectRenderCommandObject) SetMeshBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setMeshBuffer:offset:atIndex:"), buffer, offset, index)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setObjectBuffer(_:offset:at:)

func (o MTLIndirectRenderCommandObject) SetObjectBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectBuffer:offset:atIndex:"), buffer, offset, index)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setObjectThreadgroupMemoryLength(_:index:)

func (o MTLIndirectRenderCommandObject) SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectThreadgroupMemoryLength:atIndex:"), length, index)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setTriangleFillMode(_:)

func (o MTLIndirectRenderCommandObject) SetTriangleFillMode(fillMode MTLTriangleFillMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTriangleFillMode:"), fillMode)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectRenderCommand/setVertexBuffer(_:offset:attributeStride:at:)

func (o MTLIndirectRenderCommandObject) SetVertexBufferOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexBuffer:offset:attributeStride:atIndex:"), buffer, offset, stride, index)
	}

// Encodes a command to render a number of instances of tessellated patches,
// using a control point index buffer.
//
// See: /documentation/metal/mtlindirectrendercommand/drawindexedpatches(_:patchstart:patchcount:patchindexbuffer:patchindexbufferoffset:controlpointindexbuffer:controlpointindexbufferoffset:instancecount:baseinstance:tessellationfactorbuffer:tessellationfactorbufferoffset:tessellationfactorbu-4mdz8

func (o MTLIndirectRenderCommandObject) DrawIndexedPatches() {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPatches"))
	}







