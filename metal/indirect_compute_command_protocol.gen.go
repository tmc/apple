// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A compute command in an indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand
type MTLIndirectComputeCommand interface {
	objectivec.IObject

	// Sets the command’s compute pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setComputePipelineState(_:)
	SetComputePipelineState(pipelineState MTLComputePipelineState)

	// Sets the size, in pixels, of the imageblock.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setImageblockWidth(_:height:)
	SetImageblockWidthHeight(width uint, height uint)

	// Sets a buffer for the compute function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setKernelBuffer(_:offset:at:)
	SetKernelBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Sets the size of a block of threadgroup memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setThreadgroupMemoryLength(_:index:)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)

	// Sets the region of the stage-in attributes to apply to the compute kernel.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setStageInRegion(_:)
	SetStageInRegion(region MTLRegion)

	// Adds a barrier to ensure that commands executed prior to this command are complete before this command executes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setBarrier()
	SetBarrier()

	// Removes any barrier set on the command.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/clearBarrier()
	ClearBarrier()

	// Encodes a compute command using a grid aligned to threadgroup boundaries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/concurrentDispatchThreadgroups(_:threadsPerThreadgroup:)
	ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Encodes a compute command using an arbitrarily sized grid.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/concurrentDispatchThreads(_:threadsPerThreadgroup:)
	ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize)

	// Resets the command to its default state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/reset()
	Reset()

	// SetKernelBufferOffsetAttributeStrideAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setKernelBuffer(_:offset:attributeStride:at:)
	SetKernelBufferOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint)
}

// MTLIndirectComputeCommandObject wraps an existing Objective-C object that conforms to the MTLIndirectComputeCommand protocol.
type MTLIndirectComputeCommandObject struct {
	objectivec.Object
}
func (o MTLIndirectComputeCommandObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIndirectComputeCommandObjectFromID constructs a [MTLIndirectComputeCommandObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIndirectComputeCommandObjectFromID(id objc.ID) MTLIndirectComputeCommandObject {
	return MTLIndirectComputeCommandObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sets the command’s compute pipeline state.
//
// pipelineState: A compute pipeline state instance.
//
// # Discussion
// 
// You don’t need to call this method if you create an indirect command
// buffer with its [InheritPipelineState] property equal to [true]. The
// command gets the pipeline state from the parent encoder when you run the
// command.
// 
// If you create an indirect command buffer with its [InheritPipelineState]
// property equal to [false], you need to set the pipeline state prior to
// encoding a drawing command.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setComputePipelineState(_:)
func (o MTLIndirectComputeCommandObject) SetComputePipelineState(pipelineState MTLComputePipelineState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setComputePipelineState:"), pipelineState)
	}
// Sets the size, in pixels, of the imageblock.
//
// width: The width of the imageblock.
//
// height: The height of the imageblock.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setImageblockWidth(_:height:)
func (o MTLIndirectComputeCommandObject) SetImageblockWidthHeight(width uint, height uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setImageblockWidth:height:"), width, height)
	}
// Sets a buffer for the compute function.
//
// buffer: The buffer to set in the buffer argument table.
//
// offset: Where the data begins, in bytes, from the start of the buffer.
//
// index: An index in the buffer argument table.
//
// # Discussion
// 
// If you created the indirect command buffer with [InheritBuffers] set to
// [true], don’t call this method. The command gets the arguments from the
// parent encoder when you execute the command.
// 
// If you need to pass other kinds of parameters to your shader, such as
// textures and samplers, create an argument buffer and pass it to the shader
// using this method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setKernelBuffer(_:offset:at:)
func (o MTLIndirectComputeCommandObject) SetKernelBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setKernelBuffer:offset:atIndex:"), buffer, offset, index)
	}
// Sets the size of a block of threadgroup memory.
//
// length: The size of the threadgroup memory, in bytes, which needs to be a multiple
// of 16 bytes.
//
// index: The index in the threadgroup memory argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setThreadgroupMemoryLength(_:index:)
func (o MTLIndirectComputeCommandObject) SetThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
	}
// Sets the region of the stage-in attributes to apply to the compute kernel.
//
// region: The offset and maximum size of the grid over which compute threads that
// read per-thread stage-in data are launched.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setStageInRegion(_:)
func (o MTLIndirectComputeCommandObject) SetStageInRegion(region MTLRegion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setStageInRegion:"), region)
	}
// Adds a barrier to ensure that commands executed prior to this command are
// complete before this command executes.
//
// # Discussion
// 
// Set or clear barriers (as needed) before encoding the command.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setBarrier()
func (o MTLIndirectComputeCommandObject) SetBarrier() {
	
	objc.Send[struct{}](o.ID, objc.Sel("setBarrier"))
	}
// Removes any barrier set on the command.
//
// # Discussion
// 
// You need to set or clear barriers (as needed) before executing any of the
// commands in the indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/clearBarrier()
func (o MTLIndirectComputeCommandObject) ClearBarrier() {
	
	objc.Send[struct{}](o.ID, objc.Sel("clearBarrier"))
	}
// Encodes a compute command using a grid aligned to threadgroup boundaries.
//
// threadgroupsPerGrid: The number of threadgroups in the grid, in each dimension.
//
// threadsPerThreadgroup: The number of threads in one threadgroup, in each dimension.
//
// # Discussion
// 
// The command generated by this method is equivalent to calling
// [DispatchThreadgroupsThreadsPerThreadgroup].
// 
// Compute commands encoded into an indirect command buffer don’t
// automatically serialize access to resources with other commands in the
// indirect command buffer. Use barriers to serialize access to resources
// within a range of commands. See [SetBarrier].
// 
// When you execute the commands in an indirect command buffer, the
// [DispatchType] property of the compute command encoder determines whether
// the compute command encoder adds additional barriers. If [DispatchType] is
// [DispatchTypeSerial], the compute command encoder adds a barrier before and
// after the range of commands. If [DispatchType] is [DispatchTypeConcurrent],
// then the compute command encoder does nothing, and you are responsible for
// synchronizing access to resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/concurrentDispatchThreadgroups(_:threadsPerThreadgroup:)
func (o MTLIndirectComputeCommandObject) ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("concurrentDispatchThreadgroups:threadsPerThreadgroup:"), threadgroupsPerGrid, threadsPerThreadgroup)
	}
// Encodes a compute command using an arbitrarily sized grid.
//
// threadsPerGrid: The number of threads in the grid, in each dimension.
//
// threadsPerThreadgroup: The number of threads in one threadgroup, in each dimension.
//
// # Discussion
// 
// The command generated by this method is equivalent to calling
// [DispatchThreadsThreadsPerThreadgroup].
// 
// Compute commands encoded into an indirect command buffer don’t
// automatically serialize access to resources with other commands in the
// indirect command buffer. Use barriers to serialize access to resources
// within a range of commands. See [SetBarrier].
// 
// When you execute the commands in an indirect command buffer, the
// [DispatchType] property of the compute command encoder determines whether
// the compute command encoder adds additional barriers. If [DispatchType] is
// [DispatchTypeSerial], the compute command encoder adds a barrier before and
// after the range of commands. If [DispatchType] is [DispatchTypeConcurrent],
// then the compute command encoder does nothing, and you are responsible for
// synchronizing access to resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/concurrentDispatchThreads(_:threadsPerThreadgroup:)
func (o MTLIndirectComputeCommandObject) ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid MTLSize, threadsPerThreadgroup MTLSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("concurrentDispatchThreads:threadsPerThreadgroup:"), threadsPerGrid, threadsPerThreadgroup)
	}
// Resets the command to its default state.
//
// # Discussion
// 
// A command that has been reset loses any state that you previously set and
// does nothing when executed.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/reset()
func (o MTLIndirectComputeCommandObject) Reset() {
	
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectComputeCommand/setKernelBuffer(_:offset:attributeStride:at:)
func (o MTLIndirectComputeCommandObject) SetKernelBufferOffsetAttributeStrideAtIndex(buffer MTLBuffer, offset uint, stride uint, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setKernelBuffer:offset:attributeStride:atIndex:"), buffer, offset, stride, index)
	}

