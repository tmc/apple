// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An interface you can use to encode argument data into an argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder
type MTLArgumentEncoder interface {
	objectivec.IObject

	// Specifies the position in a buffer where the encoder writes argument data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setArgumentBuffer(_:offset:)
	SetArgumentBufferOffset(argumentBuffer MTLBuffer, offset uint)

	// Specifies an array element within a buffer where the encoder writes argument data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setArgumentBuffer(_:startOffset:arrayElement:)
	SetArgumentBufferStartOffsetArrayElement(argumentBuffer MTLBuffer, startOffset uint, arrayElement uint)

	// The number of bytes required to store the encoded resources of an argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/encodedLength
	EncodedLength() uint

	// Encodes a reference to a buffer into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setBuffer(_:offset:index:)
	SetBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Encodes a reference to a texture into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setTexture(_:index:)
	SetTextureAtIndex(texture MTLTexture, index uint)

	// Encodes a sampler into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setSamplerState(_:index:)
	SetSamplerStateAtIndex(sampler MTLSamplerState, index uint)

	// Encodes a reference to a render pipeline state into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setRenderPipelineState(_:index:)
	SetRenderPipelineStateAtIndex(pipeline MTLRenderPipelineState, index uint)

	// Encodes a reference to a compute pipeline state into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setComputePipelineState(_:index:)
	SetComputePipelineStateAtIndex(pipeline MTLComputePipelineState, index uint)

	// Returns a pointer to an inline, constant-data argument within the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/constantData(at:)
	ConstantDataAtIndex(index uint)

	// Encodes a reference to an indirect command buffer into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIndirectCommandBuffer(_:index:)
	SetIndirectCommandBufferAtIndex(indirectCommandBuffer MTLIndirectCommandBuffer, index uint)

	// Encodes a reference to an acceleration structure into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setAccelerationStructure(_:index:)
	SetAccelerationStructureAtIndex(accelerationStructure MTLAccelerationStructure, index uint)

	// Encodes a reference to a visible-function table into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setVisibleFunctionTable(_:index:)
	SetVisibleFunctionTableAtIndex(visibleFunctionTable MTLVisibleFunctionTable, index uint)

	// Encodes a reference to a ray-tracing intersection-function table into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIntersectionFunctionTable(_:index:)
	SetIntersectionFunctionTableAtIndex(intersectionFunctionTable MTLIntersectionFunctionTable, index uint)

	// Creates a new argument encoder for a nested argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/makeArgumentEncoderForBuffer(atIndex:)
	NewArgumentEncoderForBufferAtIndex(index uint) MTLArgumentEncoder

	// The alignment, in bytes, required for storing the encoded resources of an argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/alignment
	Alignment() uint

	// A string that identifies the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/label
	Label() string

	// The device object that created the argument encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/device
	Device() MTLDevice

	// SetDepthStencilStateAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setDepthStencilState(_:index:)
	SetDepthStencilStateAtIndex(depthStencilState MTLDepthStencilState, index uint)

	// Encodes references to an array of buffers into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setBuffers:offsets:withRange:
	SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Encodes references to an array of compute pipeline states into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setComputePipelineStates:withRange:
	SetComputePipelineStatesWithRange(pipelines []MTLComputePipelineState, range_ foundation.NSRange)

	// SetDepthStencilStatesWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setDepthStencilStates:withRange:
	SetDepthStencilStatesWithRange(depthStencilStates []MTLDepthStencilState, range_ foundation.NSRange)

	// Encodes an array of indirect command buffers into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIndirectCommandBuffers:withRange:
	SetIndirectCommandBuffersWithRange(buffers []MTLIndirectCommandBuffer, range_ foundation.NSRange)

	// Encodes references to an array of ray-tracing intersection-function tables into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIntersectionFunctionTables:withRange:
	SetIntersectionFunctionTablesWithRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange)

	// Encodes references to an array of render pipeline states into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setRenderPipelineStates:withRange:
	SetRenderPipelineStatesWithRange(pipelines []MTLRenderPipelineState, range_ foundation.NSRange)

	// Encodes an array of samplers into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setSamplerStates:withRange:
	SetSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange)

	// Encodes references to an array of textures into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setTextures:withRange:
	SetTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange)

	// Encodes references to an array of visible function tables into the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setVisibleFunctionTables:withRange:
	SetVisibleFunctionTablesWithRange(visibleFunctionTables []MTLVisibleFunctionTable, range_ foundation.NSRange)

	// A string that identifies the argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/label
	SetLabel(value string)
}

// MTLArgumentEncoderObject wraps an existing Objective-C object that conforms to the MTLArgumentEncoder protocol.
type MTLArgumentEncoderObject struct {
	objectivec.Object
}
func (o MTLArgumentEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLArgumentEncoderObjectFromID constructs a [MTLArgumentEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLArgumentEncoderObjectFromID(id objc.ID) MTLArgumentEncoderObject {
	return MTLArgumentEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Specifies the position in a buffer where the encoder writes argument data.
//
// argumentBuffer: The destination buffer that represents an argument buffer.
//
// offset: The byte offset of the buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setArgumentBuffer(_:offset:)
func (o MTLArgumentEncoderObject) SetArgumentBufferOffset(argumentBuffer MTLBuffer, offset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setArgumentBuffer:offset:"), argumentBuffer, offset)
	}
// Specifies an array element within a buffer where the encoder writes
// argument data.
//
// argumentBuffer: The destination buffer that represents an argument buffer.
//
// startOffset: The starting byte offset of the buffer data.
//
// arrayElement: The desired element of the argument buffer array targeted by encoding.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setArgumentBuffer(_:startOffset:arrayElement:)
func (o MTLArgumentEncoderObject) SetArgumentBufferStartOffsetArrayElement(argumentBuffer MTLBuffer, startOffset uint, arrayElement uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setArgumentBuffer:startOffset:arrayElement:"), argumentBuffer, startOffset, arrayElement)
	}
// The number of bytes required to store the encoded resources of an argument
// buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/encodedLength
func (o MTLArgumentEncoderObject) EncodedLength() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("encodedLength"))
	return rv
	}
// Encodes a reference to a buffer into the argument buffer.
//
// buffer: A buffer the method encodes.
//
// offset: A byte offset for `buffer`.
//
// index: The index of a buffer within the argument buffer. The value corresponds to
// either the index ID of a declaration in Metal Shading Language (MSL) or the
// [Index] property of an [MTLArgumentDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setBuffer(_:offset:index:)
func (o MTLArgumentEncoderObject) SetBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
	}
// Encodes a reference to a texture into the argument buffer.
//
// texture: A texture the method encodes.
//
// index: The index of a texture within the argument buffer. The value corresponds to
// either the index ID of a declaration in Metal Shading Language (MSL) or the
// [Index] property of an [MTLArgumentDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setTexture(_:index:)
func (o MTLArgumentEncoderObject) SetTextureAtIndex(texture MTLTexture, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTexture:atIndex:"), texture, index)
	}
// Encodes a sampler into the argument buffer.
//
// sampler: A sampler the method encodes.
//
// index: The index of a sampler within the argument buffer. The value corresponds to
// either the index ID of a declaration in Metal Shading Language (MSL) or the
// [Index] property of an [MTLArgumentDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setSamplerState(_:index:)
func (o MTLArgumentEncoderObject) SetSamplerStateAtIndex(sampler MTLSamplerState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:atIndex:"), sampler, index)
	}
// Encodes a reference to a render pipeline state into the argument buffer.
//
// pipeline: A pipeline state the method encodes.
//
// index: The index of a pipeline state within the argument buffer. The value
// corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setRenderPipelineState(_:index:)
func (o MTLArgumentEncoderObject) SetRenderPipelineStateAtIndex(pipeline MTLRenderPipelineState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setRenderPipelineState:atIndex:"), pipeline, index)
	}
// Encodes a reference to a compute pipeline state into the argument buffer.
//
// pipeline: A pipeline state the method encodes.
//
// index: The index of a pipeline state within the argument buffer. The value
// corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setComputePipelineState(_:index:)
func (o MTLArgumentEncoderObject) SetComputePipelineStateAtIndex(pipeline MTLComputePipelineState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:atIndex:"), pipeline, index)
	}
// Returns a pointer to an inline, constant-data argument within the argument
// buffer.
//
// index: The index of an inline, constant-data argument within the argument buffer.
// The value corresponds to either the index ID of a declaration in Metal
// Shading Language (MSL) or the [Index] property of an
// [MTLArgumentDescriptor] instance.
//
// # Return Value
// 
// A pointer to the location in the buffer to which you should write the
// constant data.
//
// # Discussion
// 
// Constants declared contiguously in the Metal shading language (in an array
// or structure) are contiguous in memory. You can encode contiguous ranges of
// inlined constant data through a pointer to the first constant.
// 
// To encode inlined constant data into the argument buffer, perform a memory
// copy operation from your data’s source pointer to the returned
// destination pointer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/constantData(at:)
func (o MTLArgumentEncoderObject) ConstantDataAtIndex(index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("constantDataAtIndex:"), index)
	}
// Encodes a reference to an indirect command buffer into the argument buffer.
//
// indirectCommandBuffer: An indirect command-buffer the method encodes.
//
// index: The index of an inline, constant-data argument within the argument buffer.
// The value corresponds to either the index ID of a declaration in Metal
// Shading Language (MSL) or the [Index] property of an
// [MTLArgumentDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIndirectCommandBuffer(_:index:)
func (o MTLArgumentEncoderObject) SetIndirectCommandBufferAtIndex(indirectCommandBuffer MTLIndirectCommandBuffer, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setIndirectCommandBuffer:atIndex:"), indirectCommandBuffer, index)
	}
// Encodes a reference to an acceleration structure into the argument buffer.
//
// accelerationStructure: An acceleration structure the method encodes.
//
// index: The index of an acceleration structure within the argument buffer. The
// value corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setAccelerationStructure(_:index:)
func (o MTLArgumentEncoderObject) SetAccelerationStructureAtIndex(accelerationStructure MTLAccelerationStructure, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccelerationStructure:atIndex:"), accelerationStructure, index)
	}
// Encodes a reference to a visible-function table into the argument buffer.
//
// visibleFunctionTable: A visible-function table the method encodes.
//
// index: The index of a visible-function table within the argument buffer. The value
// corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setVisibleFunctionTable(_:index:)
func (o MTLArgumentEncoderObject) SetVisibleFunctionTableAtIndex(visibleFunctionTable MTLVisibleFunctionTable, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atIndex:"), visibleFunctionTable, index)
	}
// Encodes a reference to a ray-tracing intersection-function table into the
// argument buffer.
//
// intersectionFunctionTable: An intersection-function table the method encodes.
//
// index: An index of an intersection-function table within the argument buffer. The
// value corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIntersectionFunctionTable(_:index:)
func (o MTLArgumentEncoderObject) SetIntersectionFunctionTableAtIndex(intersectionFunctionTable MTLIntersectionFunctionTable, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTable:atIndex:"), intersectionFunctionTable, index)
	}
// Creates a new argument encoder for a nested argument buffer.
//
// index: The index of a nested argument-buffer within the argument buffer. The value
// corresponds to either the index ID of a declaration in Metal Shading
// Language (MSL) or the [Index] property of an [MTLArgumentDescriptor]
// instance.
//
// # Return Value
// 
// An argument encoder targeting the nested argument buffer.
//
// # Discussion
// 
// If an argument buffer contains nested argument buffers in its structure,
// then each nested argument buffer needs to use its own [MTLArgumentEncoder]
// object to encode its individual resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/makeArgumentEncoderForBuffer(atIndex:)
func (o MTLArgumentEncoderObject) NewArgumentEncoderForBufferAtIndex(index uint) MTLArgumentEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newArgumentEncoderForBufferAtIndex:"), index)
	return MTLArgumentEncoderObjectFromID(rv)
	}
// The alignment, in bytes, required for storing the encoded resources of an
// argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/alignment
func (o MTLArgumentEncoderObject) Alignment() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("alignment"))
	return rv
	}
// A string that identifies the argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/label
func (o MTLArgumentEncoderObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The device object that created the argument encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/device
func (o MTLArgumentEncoderObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
//
// # Discussion
// 
// Sets a depth stencil state at a given bind point index
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setDepthStencilState(_:index:)
func (o MTLArgumentEncoderObject) SetDepthStencilStateAtIndex(depthStencilState MTLDepthStencilState, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStencilState:atIndex:"), depthStencilState, index)
	}
// Encodes references to an array of buffers into the argument buffer.
//
// buffers: An array of buffers the method encodes.
//
// offsets: An array of byte offsets for each element in `buffers`.
//
// range: A range of indices within the argument buffer for each element in
// `buffers`. The values correspond to either the index IDs of declarations in
// Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setBuffers:offsets:withRange:
func (o MTLArgumentEncoderObject) SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Encodes references to an array of compute pipeline states into the argument
// buffer.
//
// pipelines: An array of pipeline states the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `pipelines`. The values correspond to either the index IDs of declarations
// in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setComputePipelineStates:withRange:
func (o MTLArgumentEncoderObject) SetComputePipelineStatesWithRange(pipelines []MTLComputePipelineState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineStates:withRange:"), pipelines, range_)
	}
//
// # Discussion
// 
// Sets an array of depth stencil states at a given buffer index range
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setDepthStencilStates:withRange:
func (o MTLArgumentEncoderObject) SetDepthStencilStatesWithRange(depthStencilStates []MTLDepthStencilState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStencilStates:withRange:"), depthStencilStates, range_)
	}
// Encodes an array of indirect command buffers into the argument buffer.
//
// buffers: An array of indirect command buffers the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `buffers`. The values correspond to either the index IDs of declarations in
// Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIndirectCommandBuffers:withRange:
func (o MTLArgumentEncoderObject) SetIndirectCommandBuffersWithRange(buffers []MTLIndirectCommandBuffer, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setIndirectCommandBuffers:withRange:"), buffers, range_)
	}
// Encodes references to an array of ray-tracing intersection-function tables
// into the argument buffer.
//
// intersectionFunctionTables: An array of intersection-function tables the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `intersectionFunctionTables`. The values correspond to either the index IDs
// of declarations in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setIntersectionFunctionTables:withRange:
func (o MTLArgumentEncoderObject) SetIntersectionFunctionTablesWithRange(intersectionFunctionTables []MTLIntersectionFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setIntersectionFunctionTables:withRange:"), intersectionFunctionTables, range_)
	}
// Encodes references to an array of render pipeline states into the argument
// buffer.
//
// pipelines: An array of pipeline states the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `pipelines`. The values correspond to either the index IDs of declarations
// in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setRenderPipelineStates:withRange:
func (o MTLArgumentEncoderObject) SetRenderPipelineStatesWithRange(pipelines []MTLRenderPipelineState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setRenderPipelineStates:withRange:"), pipelines, range_)
	}
// Encodes an array of samplers into the argument buffer.
//
// samplers: An array of samplers the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `samplers`. The values correspond to either the index IDs of declarations
// in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setSamplerStates:withRange:
func (o MTLArgumentEncoderObject) SetSamplerStatesWithRange(samplers []MTLSamplerState, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerStates:withRange:"), samplers, range_)
	}
// Encodes references to an array of textures into the argument buffer.
//
// textures: An array of textures the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `textures`. The values correspond to either the index IDs of declarations
// in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setTextures:withRange:
func (o MTLArgumentEncoderObject) SetTexturesWithRange(textures []MTLTexture, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setTextures:withRange:"), textures, range_)
	}
// Encodes references to an array of visible function tables into the argument
// buffer.
//
// visibleFunctionTables: An array of visible-function tables the method encodes.
//
// range: A range of indices within the argument buffer for each element in
// `visibleFunctionTables`. The values correspond to either the index IDs of
// declarations in Metal Shading Language (MSL) or the [Index] property of
// [MTLArgumentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentEncoder/setVisibleFunctionTables:withRange:
func (o MTLArgumentEncoderObject) SetVisibleFunctionTablesWithRange(visibleFunctionTables []MTLVisibleFunctionTable, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withRange:"), visibleFunctionTables, range_)
	}

func (o MTLArgumentEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

