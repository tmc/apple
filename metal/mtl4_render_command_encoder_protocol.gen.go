// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Encodes configuration and draw commands for a single render pass into a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder
type MTL4RenderCommandEncoder interface {
	objectivec.IObject
	MTL4CommandEncoder

	// Configures this encoder with a render pipeline state that applies to your subsequent draw commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setRenderPipelineState(_:)
	SetRenderPipelineState(pipelineState MTLRenderPipelineState)

	// Configures the store action for a color attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setColorStoreAction(_:index:)
	SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint)

	// Configures the store action for the depth attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthStoreAction(_:)
	SetDepthStoreAction(storeAction MTLStoreAction)

	// Configures the store action for the stencil attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilStoreAction(_:)
	SetStencilStoreAction(storeAction MTLStoreAction)

	// Configures each pixel component value, including alpha, for the render pipeline’s constant blend color.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setBlendColor(red:green:blue:alpha:)
	SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32)

	// Sets the mapping from logical shader color output to physical render pass color attachments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setColorAttachmentMap(_:)
	SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap)

	// Configures how subsequent draw commands rasterize triangle and triangle strip primitives.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setTriangleFillMode(_:)
	SetTriangleFillMode(fillMode MTLTriangleFillMode)

	// Configures the vertex winding order that determines which face of a geometric primitive is the front one.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setFrontFacing(_:)
	SetFrontFacingWinding(frontFacingWinding MTLWinding)

	// Controls whether Metal culls front facing primitives, back facing primitives, or culls no primitives at all.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setCullMode(_:)
	SetCullMode(cullMode MTLCullMode)

	// Configures this encoder with a depth stencil state that applies to your subsequent draw commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthStencilState(_:)
	SetDepthStencilState(depthStencilState MTLDepthStencilState)

	// Configures the adjustments a render pass applies to depth values from fragment shader functions by a scaling factor and bias.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthBias(_:slopeScale:clamp:)
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)

	// Controls the behavior for fragments outside of the near or far planes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthClipMode(_:)
	SetDepthClipMode(depthClipMode MTLDepthClipMode)

	// Configures this encoder with a reference value for stencil testing.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilReferenceValue(_:)
	SetStencilReferenceValue(referenceValue uint32)

	// Configures the encoder with different stencil test reference values for front-facing and back-facing primitives.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilReferenceValue(front:back:)
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32)

	// Sets the viewport which that transforms vertices from normalized device coordinates to window coordinates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setViewport(_:)
	SetViewport(viewport MTLViewport)

	// Sets a scissor rectangle to discard fragments outside a specific area.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setScissorRect(_:)
	SetScissorRect(rect MTLScissorRect)

	// Configures a visibility test for Metal to run, and the destination for any results it generates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setVisibilityResultMode(_:offset:)
	SetVisibilityResultModeOffset(mode MTLVisibilityResultMode, offset uint)

	// Configures the size of a threadgroup memory buffer for a threadgroup argument in the object shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setObjectThreadgroupMemoryLength(_:index:)
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// Configures the size of a threadgroup memory buffer for a threadgroup argument in the fragment and tile shader functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setThreadgroupMemoryLength(_:offset:index:)
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)

	// Associates an argument table with a set of render stages.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setArgumentTable(_:stages:)
	SetArgumentTableAtStages(argumentTable MTL4ArgumentTable, stages MTLRenderStages)

	// Encodes a draw command that renders an instance of a geometric primitive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:)
	DrawPrimitivesVertexStartVertexCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:instanceCount:)
	DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive, starting with a custom instance identification number.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:instanceCount:baseInstance:)
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:indirectBuffer:)
	DrawPrimitivesIndirectBuffer(primitiveType MTLPrimitiveType, indirectBuffer MTLGPUAddress)

	// Encodes a draw command that renders an instance of a geometric primitive with indexed vertices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLength(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCount(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, instanceCount uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices, starting with a custom vertex and instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:baseVertex:baseInstance:)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, instanceCount uint, baseVertex int, baseInstance uint)

	// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices and indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexType:indexBuffer:indexBufferLength:indirectBuffer:)
	DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferLengthIndirectBuffer(primitiveType MTLPrimitiveType, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, indirectBuffer MTLGPUAddress)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with a grid of threads.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreads(threadsPerGrid:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with a grid of threadgroups.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreadgroups(threadgroupsPerGrid:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a draw command that invokes a mesh shader and, optionally, an object shader with indirect arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreadgroups(indirectBuffer:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)
	DrawMeshThreadgroupsWithIndirectBufferThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer MTLGPUAddress, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize)

	// Encodes a command that invokes a tile shader function from the encoder’s current tile render pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/dispatchThreadsPerTile(_:)
	DispatchThreadsPerTile(threadsPerTile MTLSize)

	// Sets the width of a tile for this render pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/tileWidth
	TileWidth() uint

	// Sets the height of a tile for this render pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/tileHeight
	TileHeight() uint

	// Encodes a command that runs an indirect range of commands from an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/executeCommands(buffer:indirectBuffer:)
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandBuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLGPUAddress)

	// Writes a GPU timestamp into the given [MTL4CounterHeap](<doc://com.apple.metal/documentation/Metal/MTL4CounterHeap>) at `index` after `stage` completes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/writeTimestamp(granularity:after:counterHeap:index:)
	WriteTimestampWithGranularityAfterStageIntoHeapAtIndex(granularity MTL4TimestampGranularity, stage MTLRenderStages, counterHeap MTL4CounterHeap, index uint)

	// Encodes a command that runs a range of commands from an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/executeCommandsInBuffer:withRange:
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, executionRange foundation.NSRange)

	// Configures the minimum and maximum bounds for depth bounds testing.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthTestMinBound:maxBound:
	SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32)

	// Sets an array of scissor rectangles for a fragment scissor test.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setScissorRects:count:
	SetScissorRectsCount(scissorRects []MTLScissorRect, count uint)

	// Sets the vertex amplification count and its view mapping for each amplification ID.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setVertexAmplificationCount:viewMappings:
	SetVertexAmplificationCountViewMappings(count uint, viewMappings *MTLVertexAmplificationViewMapping)

	// Sets an array of viewports to transform vertices from normalized device coordinates to window coordinates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setViewports:count:
	SetViewportsCount(viewports []MTLViewport, count uint)
}

// MTL4RenderCommandEncoderObject wraps an existing Objective-C object that conforms to the MTL4RenderCommandEncoder protocol.
type MTL4RenderCommandEncoderObject struct {
	objectivec.Object
}
func (o MTL4RenderCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4RenderCommandEncoderObjectFromID constructs a [MTL4RenderCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4RenderCommandEncoderObjectFromID(id objc.ID) MTL4RenderCommandEncoderObject {
	return MTL4RenderCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Configures this encoder with a render pipeline state that applies to your
// subsequent draw commands.
//
// pipelineState: A non-`nil` [MTLRenderPipelineState] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setRenderPipelineState(_:)

func (o MTL4RenderCommandEncoderObject) SetRenderPipelineState(pipelineState MTLRenderPipelineState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setRenderPipelineState:"), pipelineState)
	}

// Configures the store action for a color attachment.
//
// storeAction: A store action for the color attachment that can’t be
// [StoreActionUnknown].
//
// colorAttachmentIndex: The index of a color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setColorStoreAction(_:index:)

func (o MTL4RenderCommandEncoderObject) SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setColorStoreAction:atIndex:"), storeAction, colorAttachmentIndex)
	}

// Configures the store action for the depth attachment.
//
// storeAction: A store action for the depth attachment that can’t be
// [StoreActionUnknown].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthStoreAction(_:)

func (o MTL4RenderCommandEncoderObject) SetDepthStoreAction(storeAction MTLStoreAction) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStoreAction:"), storeAction)
	}

// Configures the store action for the stencil attachment.
//
// storeAction: A store action for the stencil attachment that can’t be
// [StoreActionUnknown].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilStoreAction(_:)

func (o MTL4RenderCommandEncoderObject) SetStencilStoreAction(storeAction MTLStoreAction) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilStoreAction:"), storeAction)
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setBlendColor(red:green:blue:alpha:)

func (o MTL4RenderCommandEncoderObject) SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32) {
	
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setColorAttachmentMap(_:)

func (o MTL4RenderCommandEncoderObject) SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setColorAttachmentMap:"), mapping)
	}

// Configures how subsequent draw commands rasterize triangle and triangle
// strip primitives.
//
// fillMode: [MTLTriangleFillMode] the render pass applies to draw commands that
// rasterize triangles or triangle strips.
// //
// [MTLTriangleFillMode]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setTriangleFillMode(_:)

func (o MTL4RenderCommandEncoderObject) SetTriangleFillMode(fillMode MTLTriangleFillMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTriangleFillMode:"), fillMode)
	}

// Configures the vertex winding order that determines which face of a
// geometric primitive is the front one.
//
// frontFacingWinding: A [MTLWinding] value that determines which side of a primitive the render
// pipeline interprets as front facing.
// //
// [MTLWinding]: https://developer.apple.com/documentation/Metal/MTLWinding
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setFrontFacing(_:)

func (o MTL4RenderCommandEncoderObject) SetFrontFacingWinding(frontFacingWinding MTLWinding) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFrontFacingWinding:"), frontFacingWinding)
	}

// Controls whether Metal culls front facing primitives, back facing
// primitives, or culls no primitives at all.
//
// cullMode: [MTLCullMode] to set.
// //
// [MTLCullMode]: https://developer.apple.com/documentation/Metal/MTLCullMode
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setCullMode(_:)

func (o MTL4RenderCommandEncoderObject) SetCullMode(cullMode MTLCullMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setCullMode:"), cullMode)
	}

// Configures this encoder with a depth stencil state that applies to your
// subsequent draw commands.
//
// depthStencilState: The [MTLDepthStencilState] instance to set.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthStencilState(_:)

func (o MTL4RenderCommandEncoderObject) SetDepthStencilState(depthStencilState MTLDepthStencilState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStencilState:"), depthStencilState)
	}

// Configures the adjustments a render pass applies to depth values from
// fragment shader functions by a scaling factor and bias.
//
// depthBias: A constant bias the render pipeline applies to all fragments.
//
// slopeScale: A bias coefficient that scales with the depth of the primitive relative to
// the camera.
//
// clamp: A value that limits the bias value the render pipeline can apply to a
// fragment. Pass a positive or negative value to limit the largest magnitude
// of a positive or negative bias, respectively. Set this value to `0` to
// disable bias clamping.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthBias(_:slopeScale:clamp:)

func (o MTL4RenderCommandEncoderObject) SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthBias:slopeScale:clamp:"), depthBias, slopeScale, clamp)
	}

// Controls the behavior for fragments outside of the near or far planes.
//
// depthClipMode: [MTLDepthClipMode] to set.
// //
// [MTLDepthClipMode]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthClipMode(_:)

func (o MTL4RenderCommandEncoderObject) SetDepthClipMode(depthClipMode MTLDepthClipMode) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthClipMode:"), depthClipMode)
	}

// Configures this encoder with a reference value for stencil testing.
//
// referenceValue: A stencil test comparison value.
//
// # Discussion
// 
// The render pipeline applies this reference value to both front-facing and
// back-facing primitives.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilReferenceValue(_:)

func (o MTL4RenderCommandEncoderObject) SetStencilReferenceValue(referenceValue uint32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilReferenceValue:"), referenceValue)
	}

// Configures the encoder with different stencil test reference values for
// front-facing and back-facing primitives.
//
// frontReferenceValue: A stencil test comparison value the render pipeline applies to front-facing
// primitives.
//
// backReferenceValue: A stencil test comparison value the render pipeline applies to back-facing
// primitives.
//
// # Discussion
// 
// The render pipeline applies `frontReferenceValue` to front-facing
// primitives and `backReferenceValue` to back-facing primitives.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setStencilReferenceValue(front:back:)

func (o MTL4RenderCommandEncoderObject) SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStencilFrontReferenceValue:backReferenceValue:"), frontReferenceValue, backReferenceValue)
	}

// Sets the viewport which that transforms vertices from normalized device
// coordinates to window coordinates.
//
// viewport: [MTLViewport] to set.
// //
// [MTLViewport]: https://developer.apple.com/documentation/Metal/MTLViewport
//
// # Discussion
// 
// Metal clips fragments that lie outside this viewport, and optionally clamps
// fragments outside of z-near/z-far range depending on the value you assign
// to [SetDepthClipMode].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setViewport(_:)

func (o MTL4RenderCommandEncoderObject) SetViewport(viewport MTLViewport) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setViewport:"), viewport)
	}

// Sets a scissor rectangle to discard fragments outside a specific area.
//
// rect: [MTLScissorRect] rectangle to specify. This rectangle needs to lie
// completely within the current render attachment.
// //
// [MTLScissorRect]: https://developer.apple.com/documentation/Metal/MTLScissorRect
//
// # Discussion
// 
// Metal performs a scissor test and discards all fragments outside of the
// scissor rect.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setScissorRect(_:)

func (o MTL4RenderCommandEncoderObject) SetScissorRect(rect MTLScissorRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setScissorRect:"), rect)
	}

// Configures a visibility test for Metal to run, and the destination for any
// results it generates.
//
// mode: A [MTLVisibilityResultMode] that configures which visibility test results
// the render pass saves to a buffer, or disables visibility testing.
// //
// [MTLVisibilityResultMode]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode
//
// offset: A location, in bytes, relative to the start of [VisibilityResultBuffer] The
// GPU stores the result of a visibility test at `offset`, which needs to be a
// multiple of `8`.
//
// # Discussion
// 
// You use the `mode` parameter to enable or disable the visibility test, and
// determine if it produces a boolean response for passing fragments, or if it
// counts the number of fragments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setVisibilityResultMode(_:offset:)

func (o MTL4RenderCommandEncoderObject) SetVisibilityResultModeOffset(mode MTLVisibilityResultMode, offset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVisibilityResultMode:offset:"), mode, offset)
	}

// Configures the size of a threadgroup memory buffer for a threadgroup
// argument in the object shader function.
//
// length: The size of the threadgroup memory, in bytes.
//
// index: An integer that corresponds to the index of the argument you annotate with
// attribute `[[threadgroup(index)]]` in the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setObjectThreadgroupMemoryLength(_:index:)

func (o MTL4RenderCommandEncoderObject) SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setObjectThreadgroupMemoryLength:atIndex:"), length, index)
	}

// Configures the size of a threadgroup memory buffer for a threadgroup
// argument in the fragment and tile shader functions.
//
// length: The size of the threadgroup memory, in bytes.
//
// offset: An integer that represents the location, in bytes, from the start of the
// threadgroup memory buffer at `index` where the threadgroup memory begins.
//
// index: An integer that corresponds to the index of the argument you annotate with
// attribute `[[threadgroup(index)]]` in the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setThreadgroupMemoryLength(_:offset:index:)

func (o MTL4RenderCommandEncoderObject) SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:offset:atIndex:"), length, offset, index)
	}

// Associates an argument table with a set of render stages.
//
// argumentTable: [MTL4ArgumentTable] to set.
//
// stages: A [MTLRenderStages] bitmask that specifies the shader stages with
// visibility over the table.
// //
// [MTLRenderStages]: https://developer.apple.com/documentation/Metal/MTLRenderStages
//
// # Discussion
// 
// Metal takes a snapshot of the resources in the argument table when you
// encode a draw, dispatch, or execute command. This snapshot becomes
// available to the `stages` you specify to this method.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setArgumentTable(_:stages:)

func (o MTL4RenderCommandEncoderObject) SetArgumentTableAtStages(argumentTable MTL4ArgumentTable, stages MTLRenderStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setArgumentTable:atStages:"), argumentTable, stages)
	}

// Encodes a draw command that renders an instance of a geometric primitive.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// vertexStart: The lowest value the command passes to your vertex shader function’s
// parameter with the `[[vertex_id]]` attribute.
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws.
//
// # Discussion
// 
// This command assigns each vertex a unique `vertex_id` value that increases
// from `vertexStart` through `(vertexStart + vertexCount - 1)`.
// 
// Your vertex shader function can use this value to uniquely identify each
// vertex.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:)

func (o MTL4RenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:"), primitiveType, vertexStart, vertexCount)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive.
//
// primitiveType: A [MTLPrimitiveType] represents how the command interprets vertex argument
// data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// vertexStart: The lowest value the command passes to your vertex shader function’s
// parameter with the `vertex_id` attribute.
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` primitives with `vertexCount` vertices.
//
// # Discussion
// 
// The command assigns each vertex a unique `vertex_id` value within its
// drawing instance that increases from `vertexStart` through `(vertexStart +
// vertexCount - 1)`.
// 
// Additionally, the command assigns each drawing instance a unique
// `instance_id` value that increases from `0` through `(instanceCount - 1)`.
// 
// Your vertex shader can use the `vertex_id` value to uniquely identify each
// vertex in each drawing instance, and the `instance_id` value to identify
// which instance that vertex belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:instanceCount:)

func (o MTL4RenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:"), primitiveType, vertexStart, vertexCount, instanceCount)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive, starting with a custom instance identification number.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// vertexStart: The lowest value the command passes to your vertex shader function’s
// parameter with the `vertex_id` attribute.
//
// vertexCount: An integer that represents the number of vertices of `primitiveType` the
// command draws.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `vertexCount` vertices.
//
// baseInstance: The lowest value the command passes to your vertex shader function’s
// parameter with the `instance_id` attribute.
//
// # Discussion
// 
// The command assigns each vertex a unique `vertex_id` value within its
// drawing instance that increases from `vertexStart` through `(vertexStart +
// vertexCount - 1)`.
// 
// Additionally, the command assigns each drawing instance a unique
// `instance_id` value that increases from `baseInstance` through
// `(baseInstance + instanceCount - 1)`.
// 
// Your vertex shader can use the `vertex_id` value to uniquely identify each
// vertex in each drawing instance, and the `instance_id` value to identify
// which instance that vertex belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:vertexStart:vertexCount:instanceCount:baseInstance:)

func (o MTL4RenderCommandEncoderObject) DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType MTLPrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:baseInstance:"), primitiveType, vertexStart, vertexCount, instanceCount, baseInstance)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive with indirect arguments.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// indirectBuffer: GPUAddress of a [MTLBuffer] instance with data that matches the layout of
// the [MTLDrawPrimitivesIndirectArguments] structure. You are responsible for
// ensuring that the alignment of this address is 4 bytes.
// //
// [MTLDrawPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawPrimitivesIndirectArguments
//
// # Discussion
// 
// When you use this function, Metal reads the parameters to the draw command
// from an [MTLBuffer] instance, allowing you to implement a GPU-driven
// workflow where a compute pipeline state determines the draw arguments.
// 
// You are responsible for ensuring that the address of the indirect buffer
// you provide to this method has 4-byte alignment.
// 
// Because this is a non-indexed draw call, Metal interprets the contents of
// the indirect buffer to match the layout of struct
// [MTLDrawPrimitivesIndirectArguments].
// 
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectBuffer` parameter references.
//
// [MTLDrawPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawPrimitivesIndirectArguments
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawPrimitives(primitiveType:indirectBuffer:)

func (o MTL4RenderCommandEncoderObject) DrawPrimitivesIndirectBuffer(primitiveType MTLPrimitiveType, indirectBuffer MTLGPUAddress) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawPrimitives:indirectBuffer:"), primitiveType, indirectBuffer)
	}

// Encodes a draw command that renders an instance of a geometric primitive
// with indexed vertices.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer`.
//
// indexType: A [MTLIndexType] instance that represents the index format.
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: GPUAddress of a [MTLBuffer] instance that contains `indexCount` indices of
// `indexType` format. You are responsible for ensuring this address is
// aligned to 2 bytes if the `indexType` format is [IndexTypeUInt16], and
// aligned to 4 bytes if the format is [IndexTypeUInt32].
//
// indexBufferLength: An integer that represents the length of `indexBuffer`, in bytes. You are
// responsible for ensuring this this size is a multiple of 2 if the
// `indexType` format is [IndexTypeUInt16], and a multiple of 4 if the format
// is [IndexTypeUInt32]. If this draw call causes Metal to read indices at or
// beyond the `indexBufferLength`, Metal continues to execute them assigning a
// value of `0` to the `vertex_id` attribute.
//
// # Discussion
// 
// Use this method to perform indexed drawing, where an index buffer
// determines how Metal assembles primitives.
// 
// Metal imposes some restrictions on the index buffer’s address, which
// needs to be 2- or 4-byte aligned, and its length in bytes, which needs to
// be a multiple of 2 or 4, depending on whether the format of the index is
// [IndexTypeUInt16] or [IndexTypeUInt32].
// 
// Use an instance of [MTLResidencySet] to mark residency of the index buffer
// the `indexBuffer` parameter references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:)

func (o MTL4RenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLength(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferLength:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferLength)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer`.
//
// indexType: A [MTLIndexType] instance that represents the index format.
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: GPUAddress of a [MTLBuffer] instance that contains `indexCount` indices of
// `indexType` format. You are responsible for ensuring this address is
// aligned to 2 bytes if the `indexType` format is [IndexTypeUInt16], and
// aligned to 4 bytes if the format is [IndexTypeUInt32].
//
// indexBufferLength: An integer that represents the length of `indexBuffer`, in bytes. You are
// responsible for ensuring this this size is a multiple of 2 if the
// `indexType` format is [IndexTypeUInt16], and a multiple of 4 if the format
// is [IndexTypeUInt32]. Metal disregards this value and assigns `0` to the
// `vertex_id` attribute for all primitives that require loading indices at a
// byte offset of `indexBufferLength` or greater.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `indexCount` vertices.
//
// # Discussion
// 
// Use this method to perform instanced indexed drawing, where an index buffer
// determines how Metal assembles primitives.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `0` through `(instanceCount - 1)`. Your shader can use this
// value to identify which instance the vertex belongs to.
// 
// Metal imposes some restrictions on the index buffer’s address, which
// needs to be 2- or 4-byte aligned, and its length in bytes, which needs to
// be a multiple of 2 or 4, depending on whether the format of the index is
// [IndexTypeUInt16] or [IndexTypeUInt32].
// 
// Use an instance of [MTLResidencySet] to mark residency of the index buffer
// the `indexBuffer` parameter references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:)

func (o MTL4RenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCount(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, instanceCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferLength, instanceCount)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices, starting with a custom vertex and
// instance.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// indexCount: An integer that represents the number of vertices the command reads from
// `indexBuffer`.
//
// indexType: A [MTLIndexType] instance that represents the index format.
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: GPUAddress of a [MTLBuffer] instance that contains `indexCount` indices of
// `indexType` format. You are responsible for ensuring this address is
// aligned to 2 bytes if the `indexType` format is [IndexTypeUInt16], and
// aligned to 4 bytes if the format is [IndexTypeUInt32].
//
// indexBufferLength: An integer that represents the length of `indexBuffer`, in bytes. You are
// responsible for ensuring this this size is a multiple of 2 if the
// `indexType` format is [IndexTypeUInt16], and a multiple of 4 if the format
// is [IndexTypeUInt32]. If this draw call causes Metal to read indices at or
// beyond the `indexBufferLength`, Metal continues to execute them assigning a
// value of `0` to the `vertex_id` attribute.
//
// instanceCount: An integer that represents the number of times the command draws
// `primitiveType` with `indexCount` vertices.
//
// baseVertex: The lowest value the command passes to your vertex shader functions’s
// parameter with the `vertex_id` attribute. Metal disregards this value and
// assigns `0` to the `vertex_id` attribute for all primitives that require
// loading indices at a byte offset of `indexBufferLength` or greater.
//
// baseInstance: The lowest value the command passes to your vertex shader’s parameter
// with the `instance_id` attribute.
//
// # Discussion
// 
// Use this method to perform instanced indexed drawing, where an index buffer
// determines how Metal assembles primitives whilst customizing the base
// vertex and base instance value Metal passes to the vertex shader function.
// 
// The command assigns each drawing instance a unique `instance_id` value that
// increases from `baseInstance` through `(baseInstance + instanceCount - 1)`.
// Your shader can use this value to identify which instance the vertex
// belongs to.
// 
// Metal imposes some restrictions on the index buffer’s address, which
// needs to be 2- or 4-byte aligned, and its length in bytes, which needs to
// be a multiple of 2 or 4, depending on whether the format of the index is
// [IndexTypeUInt16] or [IndexTypeUInt32].
// 
// Use an instance of [MTLResidencySet] to mark residency of the index buffer
// the `indexBuffer` parameter references.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:baseVertex:baseInstance:)

func (o MTL4RenderCommandEncoderObject) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCountBaseVertexBaseInstance(primitiveType MTLPrimitiveType, indexCount uint, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, instanceCount uint, baseVertex int, baseInstance uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferLength:instanceCount:baseVertex:baseInstance:"), primitiveType, indexCount, indexType, indexBuffer, indexBufferLength, instanceCount, baseVertex, baseInstance)
	}

// Encodes a draw command that renders multiple instances of a geometric
// primitive with indexed vertices and indirect arguments.
//
// primitiveType: A [MTLPrimitiveType] representing how the command interprets vertex
// argument data.
// //
// [MTLPrimitiveType]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
//
// indexType: A [MTLIndexType] instance that represents the index format.
// //
// [MTLIndexType]: https://developer.apple.com/documentation/Metal/MTLIndexType
//
// indexBuffer: GPUAddress of a [MTLBuffer] instance that contains `indexCount` indices of
// `indexType` format. You are responsible for ensuring this address is
// aligned to 2 bytes if the `indexType` format is [IndexTypeUInt16], and
// aligned to 4 bytes if the format is [IndexTypeUInt32].
//
// indexBufferLength: An integer that represents the length of `indexBuffer`, in bytes. You are
// responsible for ensuring this this size is a multiple of 2 if the
// `indexType` format is [IndexTypeUInt16], and a multiple of 4 if the format
// is [IndexTypeUInt32]. If this draw call causes Metal to read indices at or
// beyond the `indexBufferLength`, Metal continues to execute them assigning a
// value of `0` to the `vertex_id` attribute.
//
// indirectBuffer: GPUAddress of an [MTLBuffer] instance with data that matches the layout of
// the [MTLDrawIndexedPrimitivesIndirectArguments] structure. This address
// requires 4-byte alignment.
// //
// [MTLDrawIndexedPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawIndexedPrimitivesIndirectArguments
//
// # Discussion
// 
// When you use this function, Metal reads the parameters to the draw command
// from an [MTLBuffer] instance, allowing you to implement a GPU-driven
// workflow where a compute pipeline state determines the draw arguments.
// 
// Because this is an indexed draw call, Metal interprets the contents of the
// indirect buffer to match the layout of struct
// [MTLDrawIndexedPrimitivesIndirectArguments], which includes `indexStart`
// and `indexCount` members, denoting a range within the index buffer you
// provide in the `indexBuffer` parameter.
// 
// The range of indices within the `indexBuffer` form the primitives Metal
// draws.
// 
// Metal imposes some restrictions on the index buffer’s address, which
// needs to be 2- or 4-byte aligned, and its length in bytes, which needs to
// be a multiple of 2 or 4, depending on whether the format of the index is
// [IndexTypeUInt16] or [IndexTypeUInt32].
// 
// Similarly, you are responsible for ensuring the indirect buffer’s address
// has 4-byte alignment.
// 
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectBuffer` parameter references, and of the index
// buffer the `indexBuffer` parameter references.
//
// [MTLDrawIndexedPrimitivesIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDrawIndexedPrimitivesIndirectArguments
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawIndexedPrimitives(primitiveType:indexType:indexBuffer:indexBufferLength:indirectBuffer:)

func (o MTL4RenderCommandEncoderObject) DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferLengthIndirectBuffer(primitiveType MTLPrimitiveType, indexType MTLIndexType, indexBuffer MTLGPUAddress, indexBufferLength uint, indirectBuffer MTLGPUAddress) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawIndexedPrimitives:indexType:indexBuffer:indexBufferLength:indirectBuffer:"), primitiveType, indexType, indexBuffer, indexBufferLength, indirectBuffer)
	}

// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with a grid of threads.
//
// threadsPerGrid: A [MTLSize] instance that represents the number of threads for each grid
// dimension. For mesh shaders, the command rounds the value down to the
// nearest multiple of `threadsPerMeshThreadgroup` for each dimension. For
// object shaders, the value doesn’t need to be a multiple of
// `threadsPerObjectThreadgroup`.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerObjectThreadgroup: A [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: A [MTLSize] instance that represents the number of threads in a mesh shader
// threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreads(threadsPerGrid:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)

func (o MTL4RenderCommandEncoderObject) DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreads:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}

// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with a grid of threadgroups.
//
// threadgroupsPerGrid: A [MTLSize] instance that represents the number of threadgroups for each
// grid dimension.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerObjectThreadgroup: A [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: A [MTLSize] instance that represents the number of threads in a mesh shader
// threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreadgroups(threadgroupsPerGrid:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)

func (o MTL4RenderCommandEncoderObject) DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid MTLSize, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreadgroups:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), threadgroupsPerGrid, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}

// Encodes a draw command that invokes a mesh shader and, optionally, an
// object shader with indirect arguments.
//
// indirectBuffer: GPUAddress of an [MTLBuffer] instance with data that matches the layout of
// the [MTLDispatchThreadgroupsIndirectArguments] structure. This address
// requires 4-byte alignment.
// //
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
//
// threadsPerObjectThreadgroup: A [MTLSize] instance that represents the number of threads in an object
// shader threadgroup, if applicable.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// threadsPerMeshThreadgroup: A [MTLSize] instance that represents the number of threads in a mesh shader
// threadgroup.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// # Discussion
// 
// This method enables you to determine the number of threadgroups per grid
// indirectly, in the GPU timeline. Metal expects this buffer’s contents to
// match the layout of structure [MTLDispatchThreadgroupsIndirectArguments].
// You are responsible for ensuring the address of this buffer has 4-byte
// alignment.
// 
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectBuffer` parameter references.
//
// [MTLDispatchThreadgroupsIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/drawMeshThreadgroups(indirectBuffer:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:)

func (o MTL4RenderCommandEncoderObject) DrawMeshThreadgroupsWithIndirectBufferThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer MTLGPUAddress, threadsPerObjectThreadgroup MTLSize, threadsPerMeshThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawMeshThreadgroupsWithIndirectBuffer:threadsPerObjectThreadgroup:threadsPerMeshThreadgroup:"), indirectBuffer, threadsPerObjectThreadgroup, threadsPerMeshThreadgroup)
	}

// Encodes a command that invokes a tile shader function from the encoder’s
// current tile render pipeline state.
//
// threadsPerTile: A [MTLSize] instance that represents the number of threads the render pass
// uses per tile. Set the size’s [width] and [height] properties to values
// that are less than or equal to [TileWidth] and [TileHeight], respectively.
// Some GPU families only support square tile dispatches and require the same
// value for width and height. Set [depth] to `1`.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
// [width]: https://developer.apple.com/documentation/Metal/MTLSize/width
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/dispatchThreadsPerTile(_:)

func (o MTL4RenderCommandEncoderObject) DispatchThreadsPerTile(threadsPerTile MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("dispatchThreadsPerTile:"), threadsPerTile)
	}

// Sets the width of a tile for this render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/tileWidth

func (o MTL4RenderCommandEncoderObject) TileWidth() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("tileWidth"))
	return rv
	}

// Sets the height of a tile for this render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/tileHeight

func (o MTL4RenderCommandEncoderObject) TileHeight() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("tileHeight"))
	return rv
	}

// Encodes a command that runs an indirect range of commands from an indirect
// command buffer.
//
// indirectCommandBuffer: A [MTLIndirectCommandBuffer] instance that contains other commands the
// current command runs.
//
// indirectRangeBuffer: GPUAddress of a [MTLBuffer] instance with data that matches the layout of
// the [MTLIndirectCommandBufferExecutionRange] structure. This address
// requires 4-byte alignment.
// //
// [MTLIndirectCommandBufferExecutionRange]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
//
// # Discussion
// 
// Use this method to indicate to Metal the span of indices in the command
// buffer to execute indirectly via an [MTLBuffer] instance you provide in the
// `indirectRangeBuffer` parameter. This allows you to calculate the span of
// commands Metal executes in the GPU timeline, enabling GPU-driven workflows.
// 
// Metal requires that the contents of this buffer match the layout of struct
// [MTLIndirectCommandBufferExecutionRange], which specifies a location and a
// length within the indirect command buffer. You are responsible for ensuring
// the address of this buffer has 4-byte alignment.
// 
// Use an instance of [MTLResidencySet] to mark residency of the indirect
// buffer that the `indirectRangeBuffer` parameter references.
//
// [MTLIndirectCommandBufferExecutionRange]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/executeCommands(buffer:indirectBuffer:)

func (o MTL4RenderCommandEncoderObject) ExecuteCommandsInBufferIndirectBuffer(indirectCommandBuffer MTLIndirectCommandBuffer, indirectRangeBuffer MTLGPUAddress) {
	
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:indirectBuffer:"), indirectCommandBuffer, indirectRangeBuffer)
	}

// Writes a GPU timestamp into the given [MTL4CounterHeap] at `index` after
// `stage` completes.
//
// granularity: A [MTL4TimestampGranularity] hint.
// //
// [MTL4TimestampGranularity]: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity
//
// stage: [MTLRenderStages] that need to complete before Metal writes the timestamp.
// This may also include later stages that are related, for example
// [RenderStageMesh] may include [RenderStageVertex].
// //
// [MTLRenderStages]: https://developer.apple.com/documentation/Metal/MTLRenderStages
//
// counterHeap: [MTL4CounterHeap] into which Metal writes timestamps.
//
// index: The index value into which Metal writes this timestamp.
//
// # Discussion
// 
// This command only guarantees all draws prior to this command are complete
// when Metal writes the timestamp into the counter heap you provide in the
// `counterHeap` parameter. The timestamp may also include subsequent
// operations.
// 
// If you call this method before any draw calls, Metal writes a timestamp
// before the stage you specify in the `stage` parameter begins.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/writeTimestamp(granularity:after:counterHeap:index:)

func (o MTL4RenderCommandEncoderObject) WriteTimestampWithGranularityAfterStageIntoHeapAtIndex(granularity MTL4TimestampGranularity, stage MTLRenderStages, counterHeap MTL4CounterHeap, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writeTimestampWithGranularity:afterStage:intoHeap:atIndex:"), granularity, stage, counterHeap, index)
	}

// Encodes a command that runs a range of commands from an indirect command
// buffer.
//
// indirectCommandBuffer: A [MTLIndirectCommandBuffer] instance containing other commands that the
// current command runs.
//
// executionRange: A span of integers that represent the command entries in the buffer that
// the current command runs.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/executeCommandsInBuffer:withRange:

func (o MTL4RenderCommandEncoderObject) ExecuteCommandsInBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, executionRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("executeCommandsInBuffer:withRange:"), indirectCommandBuffer, executionRange)
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setDepthTestMinBound:maxBound:

func (o MTL4RenderCommandEncoderObject) SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthTestMinBound:maxBound:"), minBound, maxBound)
	}

// Sets an array of scissor rectangles for a fragment scissor test.
//
// scissorRects: Array of [MTLScissorRect] structures.
// //
// [MTLScissorRect]: https://developer.apple.com/documentation/Metal/MTLScissorRect
//
// count: Number of [MTLScissorRect] structures in the array.
// //
// [MTLScissorRect]: https://developer.apple.com/documentation/Metal/MTLScissorRect
//
// # Discussion
// 
// Metal uses the specific scissor rectangle corresponding to the index you
// specify via the `[[ viewport_array_index ]]` output attribute of the vertex
// shader function in the Metal Shading Language, discarding all fragments
// outside of the scissor rect.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setScissorRects:count:

func (o MTL4RenderCommandEncoderObject) SetScissorRectsCount(scissorRects []MTLScissorRect, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setScissorRects:count:"), objc.CArray(scissorRects), count)
	}

// Sets the vertex amplification count and its view mapping for each
// amplification ID.
//
// count: The number of outputs to create. The maximum value is `2`.
//
// viewMappings: Array of [MTLVertexAmplificationViewMapping] instances. Each instance
// provides per-output offsets to a specific render target and viewport.
// //
// [MTLVertexAmplificationViewMapping]: https://developer.apple.com/documentation/Metal/MTLVertexAmplificationViewMapping
//
// # Discussion
// 
// Each view mapping element describes how to route the corresponding
// amplification ID to a specific viewport and render target array index by
// using offsets from the base array index provided by the `[[
// render_target_array_index ]]` and/or `[[ viewport_array_index ]]` output
// attributes in the vertex shader. This allows Metal to route each amplified
// vertex to a different `[[ render_target_array_index ]]` and `[[
// viewport_array_index ]]`, even though you can’t directly amplify these
// attributes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setVertexAmplificationCount:viewMappings:

func (o MTL4RenderCommandEncoderObject) SetVertexAmplificationCountViewMappings(count uint, viewMappings *MTLVertexAmplificationViewMapping) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVertexAmplificationCount:viewMappings:"), count, viewMappings)
	}

// Sets an array of viewports to transform vertices from normalized device
// coordinates to window coordinates.
//
// viewports: Array of [MTLViewport] instances.
// //
// [MTLViewport]: https://developer.apple.com/documentation/Metal/MTLViewport
//
// count: Number of [MTLViewport] instances in the array.
// //
// [MTLViewport]: https://developer.apple.com/documentation/Metal/MTLViewport
//
// # Discussion
// 
// Metal clips fragments that lie outside of the viewport, and optionally
// clamps fragments outside of z-near/z-far range, depending on the value you
// assign to [SetDepthClipMode].
// 
// Metal selects the viewport to use from the `[[ viewport_array_index ]]`
// attribute you specify in the pipeline state’s vertex shader function in
// the Metal Shading Language.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderCommandEncoder/setViewports:count:

func (o MTL4RenderCommandEncoderObject) SetViewportsCount(viewports []MTLViewport, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setViewports:count:"), objc.CArray(viewports), count)
	}

// Returns the command buffer that is currently encoding commands.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/commandBuffer

func (o MTL4RenderCommandEncoderObject) CommandBuffer() MTL4CommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBuffer"))
	return MTL4CommandBufferObjectFromID(rv)
	}

// Provides an optional label to assign to the command encoder for debug
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/label

func (o MTL4RenderCommandEncoderObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Declares that all command generation from this encoder is complete.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/endEncoding()

func (o MTL4RenderCommandEncoderObject) EndEncoding() {
	
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

func (o MTL4RenderCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
	}

// Pops the latest debug group string from this encoder’s stack of debug
// groups.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/popDebugGroup()

func (o MTL4RenderCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}

// Pushes a string onto this encoder’s stack of debug groups.
//
// string: The debug string to push.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/pushDebugGroup(_:)

func (o MTL4RenderCommandEncoderObject) PushDebugGroup(string_ string) {
	
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
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/updateFence(_:afterEncoderStages:)

func (o MTL4RenderCommandEncoderObject) UpdateFenceAfterEncoderStages(fence MTLFence, afterEncoderStages MTLStages) {
	
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
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/waitForFence(_:beforeEncoderStages:)

func (o MTL4RenderCommandEncoderObject) WaitForFenceBeforeEncoderStages(fence MTLFence, beforeEncoderStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:beforeEncoderStages:"), fence, beforeEncoderStages)
	}

// Encodes an intra-pass barrier.
//
// afterEncoderStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument only applies to subsequent work you encode in the current command
// encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeEncoderStages: [MTLStages] mask that represents the stages of work that wait. This
// argument only applies to work you encode in the current command encoder
// prior to this barrier.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// visibilityOptions: [MTL4VisibilityOptions] of the barrier, controlling cache flush behavior.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// Encode a barrier that guarantees that any subsequent work you encode in the
// , corresponding to `beforeEncoderStages`, doesn’t begin until all prior
// commands in this command encoder, corresponding to `afterEncoderStages`,
// completes.
// 
// When calling this method, it’s your responsibility to ensure parameters
// `afterEncoderStages` and `beforeEncoderStages` contain a combination of
// [MTLStages] for which this encoder can encode commands. For example, for a
// [MTL4ComputeCommandEncoder] instance, you can provide any combination of
// [StageDispatch], [StageBlit] and [StageAccelerationStructure].
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterEncoderStages:beforeEncoderStages:visibilityOptions:

func (o MTL4RenderCommandEncoderObject) BarrierAfterEncoderStagesBeforeEncoderStagesVisibilityOptions(afterEncoderStages MTLStages, beforeEncoderStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterEncoderStages:beforeEncoderStages:visibilityOptions:"), afterEncoderStages, beforeEncoderStages, visibilityOptions)
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
// visibilityOptions: [MTL4VisibilityOptions] of the barrier.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// Encode a barrier that guarantees that any subsequent work you encode in the
// current command encoder that corresponds to the `beforeStages` stages
// doesn’t proceed until Metal completes all work prior to the current
// command encoder corresponding to the `afterQueueStages` stages, completes.
// 
// Metal can reorder the exact point where it applies the barrier, so encode
// the barrier as close to the command that consumes the resource as possible.
// Don’t use this method for synchronizing resource access within the same
// pass.
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
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterQueueStages:beforeStages:visibilityOptions:

func (o MTL4RenderCommandEncoderObject) BarrierAfterQueueStagesBeforeStagesVisibilityOptions(afterQueueStages MTLStages, beforeStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:visibilityOptions:"), afterQueueStages, beforeStages, visibilityOptions)
	}

// Encodes a producer barrier on work committed to the same command queue.
//
// afterStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in the
// current command encoder prior to this barrier command.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeQueueStages: [MTLStages] mask that represents the stages of work that need to wait. This
// argument applies to subsequent encoders and not to work in the current
// command encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// visibilityOptions: [MTL4VisibilityOptions] of the barrier, controlling cache flush behavior.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// This method encodes a barrier that guarantees that any work you encode
// using , corresponding to `beforeQueueStages`, don’t begin until all
// commands you previously encode in the current encoder (and prior encoders),
// corresponding to `afterStages`, complete.
// 
// When calling this method, you can pass any [MTLStages] to parameters
// `afterStages` and `beforeQueueStages`, even stages that don’t relate to
// the current or prior command encoders.
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterStages:beforeQueueStages:visibilityOptions:

func (o MTL4RenderCommandEncoderObject) BarrierAfterStagesBeforeQueueStagesVisibilityOptions(afterStages MTLStages, beforeQueueStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterStages:beforeQueueStages:visibilityOptions:"), afterStages, beforeQueueStages, visibilityOptions)
	}

func (o MTL4RenderCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

