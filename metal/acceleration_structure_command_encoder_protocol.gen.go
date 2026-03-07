// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Encodes commands that build and refit acceleration structures for a single pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder
type MTLAccelerationStructureCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Encodes a command to build a new acceleration structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/build(accelerationStructure:descriptor:scratchBuffer:scratchBufferOffset:)
	BuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset(accelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, scratchBuffer MTLBuffer, scratchBufferOffset uint)

	// Encodes a command to copy the data from one acceleration structure to another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/copy(sourceAccelerationStructure:destinationAccelerationStructure:)
	CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure)

	// Encodes a command to calculate the compacted size of an acceleration structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/writeCompactedSize(accelerationStructure:buffer:offset:)
	WriteCompactedAccelerationStructureSizeToBufferOffset(accelerationStructure MTLAccelerationStructure, buffer MTLBuffer, offset uint)

	// Encodes a command to calculate the compacted size of an acceleration structure, taking into account the size of the output data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/writeCompactedSize(accelerationStructure:buffer:offset:sizeDataType:)
	WriteCompactedAccelerationStructureSizeToBufferOffsetSizeDataType(accelerationStructure MTLAccelerationStructure, buffer MTLBuffer, offset uint, sizeDataType MTLDataType)

	// Encodes a command to compact an acceleration structure’s data and copy it into a different acceleration structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/copyAndCompact(sourceAccelerationStructure:destinationAccelerationStructure:)
	CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure)

	// Updates an acceleration structure with new geometry or instance data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:scratchBufferOffset:)
	RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTLBuffer, scratchBufferOffset uint)

	// Updates an acceleration structure with new geometry or instance data, with options that control the refitting process.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:scratchBufferOffset:options:)
	RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffsetOptions(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTLBuffer, scratchBufferOffset uint, options MTLAccelerationStructureRefitOptions)

	// Encodes a command that instructs the GPU to update a fence after the acceleration structure pass completes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/updateFence(_:)
	UpdateFence(fence MTLFence)

	// Encodes a command that instructs the GPU to pause the acceleration structure pass until another pass updates a fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/waitForFence(_:)
	WaitForFence(fence MTLFence)

	// Makes the resources contained in the specified heap available to the acceleration structure pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useHeap(_:)
	UseHeap(heap MTLHeap)

	// Makes a resource available to the acceleration structure pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useResource(_:usage:)
	UseResourceUsage(resource MTLResource, usage MTLResourceUsage)

	// Encodes a command to sample hardware counters at this point in the acceleration structure pass and store the samples into a counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool)

	// Specifies that an array of heaps containing resources in an argument buffer can be safely used by the acceleration structure pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useHeaps:count:
	UseHeapsCount(heaps []MTLHeap, count uint)

	// Specifies that an array of resources in an argument buffer can be safely used by the acceleration structure pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useResources:count:usage:
	UseResourcesCountUsage(resources []MTLResource, count uint, usage MTLResourceUsage)
}



// MTLAccelerationStructureCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLAccelerationStructureCommandEncoder protocol.
type MTLAccelerationStructureCommandEncoderObject struct {
	objectivec.Object
}
func (o MTLAccelerationStructureCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLAccelerationStructureCommandEncoderObjectFromID constructs a [MTLAccelerationStructureCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLAccelerationStructureCommandEncoderObjectFromID(id objc.ID) MTLAccelerationStructureCommandEncoderObject {
	return MTLAccelerationStructureCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Encodes a command to build a new acceleration structure.
//
// accelerationStructure: The acceleration structure to build into.
//
// descriptor: A description of the new acceleration structure.
//
// scratchBuffer: A buffer used to hold data while building the acceleration structure.
//
// scratchBufferOffset: An offset, in, bytes, in the scratch buffer where the scratch memory
// starts.
//
// # Discussion
// 
// The destination acceleration structure and the scratch buffer needs enough
// space in memory to hold the acceleration structure data. Call the
// [AccelerationStructureSizesWithDescriptor] method on the Metal device
// object to get the required space.
// 
// The resulting acceleration structure contains references to any other
// acceleration structures referenced by the descriptor, but not any other
// underlying buffer data. As part of creating the structure, the GPU
// overwrites data in the scratch buffer.
// 
// By default, Metal automatically synchronizes GPU access to the acceleration
// structure and any resources contained in the acceleration descriptor. For
// example, you can use a single encoder to build multiple geometry
// acceleration structures and then build the instanced structure that uses
// them. Similarly, you can use the acceleration structure in an encoder
// scheduled after this encoder.
// 
// If you are using untracked resources, you are responsible for synchronizing
// access to any untracked resources, including the acceleration structure.
// For example, you could use one encoder to build the geometry acceleration
// structures and update a fence, and a second encoder that waits on the fence
// and builds the instance acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/build(accelerationStructure:descriptor:scratchBuffer:scratchBufferOffset:)

func (o MTLAccelerationStructureCommandEncoderObject) BuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset(accelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, scratchBuffer MTLBuffer, scratchBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("buildAccelerationStructure:descriptor:scratchBuffer:scratchBufferOffset:"), accelerationStructure, descriptor, scratchBuffer, scratchBufferOffset)
	}

// Encodes a command to copy the data from one acceleration structure to
// another.
//
// sourceAccelerationStructure: The source acceleration structure.
//
// destinationAccelerationStructure: The destination acceleration structure.
//
// # Discussion
// 
// The destination acceleration structure needs to be at least as large as the
// source acceleration structure, unless you’re compacting the source
// acceleration structure. In that case, the destination acceleration
// structure needs be at least as large as the compact size of the source
// acceleration structure.
// 
// If the source acceleration structure contains references to other
// acceleration structures, the copy of the acceleration structure also refers
// to the same child structures.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/copy(sourceAccelerationStructure:destinationAccelerationStructure:)

func (o MTLAccelerationStructureCommandEncoderObject) CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyAccelerationStructure:toAccelerationStructure:"), sourceAccelerationStructure, destinationAccelerationStructure)
	}

// Encodes a command to calculate the compacted size of an acceleration
// structure.
//
// accelerationStructure: The acceleration structure to measure.
//
// buffer: The buffer to write the size into.
//
// offset: An offset, in bytes, where the GPU should write the result.
//
// # Discussion
// 
// The GPU writes the compacted size to the buffer as a 32-bit unsigned
// integer representing the compacted size in bytes. The compacted size may be
// smaller than the source acceleration structure.
// 
// To compact an acceleration structure, encode a command to get the minimum
// size. After the command completes, read the size from the buffer and
// allocate a new acceleration structure with at least that much storage. Then
// create another encoder and call the
// [CopyAndCompactAccelerationStructureToAccelerationStructure] method to copy
// it into the new structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/writeCompactedSize(accelerationStructure:buffer:offset:)

func (o MTLAccelerationStructureCommandEncoderObject) WriteCompactedAccelerationStructureSizeToBufferOffset(accelerationStructure MTLAccelerationStructure, buffer MTLBuffer, offset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writeCompactedAccelerationStructureSize:toBuffer:offset:"), accelerationStructure, buffer, offset)
	}

// Encodes a command to calculate the compacted size of an acceleration
// structure, taking into account the size of the output data.
//
// accelerationStructure: The acceleration structure to measure.
//
// buffer: The buffer to write the size into.
//
// offset: An offset, in bytes, where the GPU should write the result.
//
// sizeDataType: The data type of the resulting data.
//
// # Discussion
// 
// The GPU writes the compacted size to the buffer, in bytes, using the
// `sizeDataType` parameter to determine the size of the output data. The
// compacted size may be smaller than the source acceleration structure.
// 
// To compact an acceleration structure, encode a command to get the minimum
// size. After the command completes, read the size from the buffer and
// allocate a new acceleration structure with at least that much storage. Then
// create another encoder and call the
// [CopyAndCompactAccelerationStructureToAccelerationStructure] method to copy
// it into the new structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/writeCompactedSize(accelerationStructure:buffer:offset:sizeDataType:)

func (o MTLAccelerationStructureCommandEncoderObject) WriteCompactedAccelerationStructureSizeToBufferOffsetSizeDataType(accelerationStructure MTLAccelerationStructure, buffer MTLBuffer, offset uint, sizeDataType MTLDataType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writeCompactedAccelerationStructureSize:toBuffer:offset:sizeDataType:"), accelerationStructure, buffer, offset, sizeDataType)
	}

// Encodes a command to compact an acceleration structure’s data and copy it
// into a different acceleration structure.
//
// sourceAccelerationStructure: The source acceleration structure.
//
// destinationAccelerationStructure: The destination acceleration structure.
//
// # Discussion
// 
// The source and destination acceleration structures can’t overlap in
// memory. The destination acceleration structure needs to be at least as
// large as the compact size of the source acceleration structure, which you
// obtain by using the [WriteCompactedAccelerationStructureSizeToBufferOffset]
// method.
// 
// If the source acceleration structure contains references to other
// acceleration structures, the copy of the acceleration structure refers to
// the same child structures.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/copyAndCompact(sourceAccelerationStructure:destinationAccelerationStructure:)

func (o MTLAccelerationStructureCommandEncoderObject) CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure MTLAccelerationStructure, destinationAccelerationStructure MTLAccelerationStructure) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyAndCompactAccelerationStructure:toAccelerationStructure:"), sourceAccelerationStructure, destinationAccelerationStructure)
	}

// Updates an acceleration structure with new geometry or instance data.
//
// sourceAccelerationStructure: The source acceleration structure.
//
// descriptor: A description of the updated acceleration structure.
//
// destinationAccelerationStructure: The destination to write the new acceleration structure to. Pass the same
// acceleration structure or `nil` to refit the structure in place.
//
// scratchBuffer: A buffer used to hold data while building the acceleration structure. Pass
// `nil` if [refitScratchBufferSize] returns zero.
// //
// [refitScratchBufferSize]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureSizes/refitScratchBufferSize
//
// scratchBufferOffset: An offset, in bytes, in the scratch buffer where the scratch memory starts.
//
// # Discussion
// 
// Use refitting to update an acceleration structure when you make small
// changes to the underlying geometry. Refitting performs much faster than
// rebuilding an acceleration structure from scratch. However, ray-tracing
// performance may degrade, based on how many changes you make to the geometry
// data.
// 
// You can’t use refitting to add or remove geometry in the acceleration
// structure.
// 
// If the source and destination acceleration structures aren’t the same,
// they can’t overlap in memory. The destination acceleration structure and
// the scratch buffer need to have enough space in memory to hold the
// acceleration structure data. Get the minimum amount of space it needs by
// calling the [AccelerationStructureSizesWithDescriptor] method of the Metal
// device instance. If you’re compacting the source structure, the
// destination needs to be at least as large as the compact size of the source
// acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:scratchBufferOffset:)

func (o MTLAccelerationStructureCommandEncoderObject) RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTLBuffer, scratchBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:scratchBufferOffset:"), sourceAccelerationStructure, descriptor, destinationAccelerationStructure, scratchBuffer, scratchBufferOffset)
	}

// Updates an acceleration structure with new geometry or instance data, with
// options that control the refitting process.
//
// sourceAccelerationStructure: The source acceleration structure.
//
// descriptor: A description of the updated acceleration structure.
//
// destinationAccelerationStructure: The destination to write the new acceleration structure to. Pass the same
// acceleration structure or `nil` to refit the structure in place.
//
// scratchBuffer: A buffer used to hold data while building the acceleration structure. Pass
// `nil` if [refitScratchBufferSize] returns zero.
// //
// [refitScratchBufferSize]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureSizes/refitScratchBufferSize
//
// scratchBufferOffset: An offset, in bytes, in the scratch buffer where the scratch memory starts.
//
// options: Options that control the refitting process.
//
// # Discussion
// 
// Use refitting to update an acceleration structure when you make small
// changes to the underlying geometry. Refitting performs much faster than
// rebuilding an acceleration structure from scratch. However, ray-tracing
// performance may degrade, based on how many changes you make to the geometry
// data.
// 
// You can’t use refitting to add or remove geometry in the acceleration
// structure.
// 
// If the source and destination acceleration structures are not the same,
// they need to avoid overlapping in memory. The destination acceleration
// structure and the scratch buffer need to have enough space in memory to
// hold the acceleration structure data. Call the
// [AccelerationStructureSizesWithDescriptor] method on the Metal device
// object to get the required space. If you compact the source structure, the
// destination needs to be at least as large as the compacted size of the
// source acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/refit(sourceAccelerationStructure:descriptor:destinationAccelerationStructure:scratchBuffer:scratchBufferOffset:options:)

func (o MTLAccelerationStructureCommandEncoderObject) RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffsetOptions(sourceAccelerationStructure MTLAccelerationStructure, descriptor IMTLAccelerationStructureDescriptor, destinationAccelerationStructure MTLAccelerationStructure, scratchBuffer MTLBuffer, scratchBufferOffset uint, options MTLAccelerationStructureRefitOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:scratchBufferOffset:options:"), sourceAccelerationStructure, descriptor, destinationAccelerationStructure, scratchBuffer, scratchBufferOffset, options)
	}

// Encodes a command that instructs the GPU to update a fence after the
// acceleration structure pass completes.
//
// fence: A fence the pass updates after it completes.
//
// # Discussion
// 
// You can synchronize memory operations of an acceleration structure pass
// that access resources with an [MTLFence]. This method instructs the pass to
// update `fence` after it runs all its memory store operations to the
// resources it accesses. The fence indicates when other passes can access
// those resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding an acceleration structure pass that reuses a fence, wait for
// other passes to update the fence before repurposing that fence to notify
// subsequent passes with an update:
// 
// - Call the [WaitForFence] method before encoding commands that need to wait
// for other passes. - Call the [UpdateFence] method after encoding commands
// that later passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/updateFence(_:)

func (o MTLAccelerationStructureCommandEncoderObject) UpdateFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
	}

// Encodes a command that instructs the GPU to pause the acceleration
// structure pass until another pass updates a fence.
//
// fence: A fence that the pass waits for before it runs any of its commands.
//
// # Discussion
// 
// You can synchronize memory operations of an acceleration structure pass
// that access resources with an [MTLFence]. This method instructs the GPU to
// wait until another pass updates `fence` before running the acceleration
// structure pass. The fence indicates when the pass can access those
// resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding an acceleration structure pass that reuses a fence, wait for
// other passes to update the fence before repurposing that fence to notify
// subsequent passes with an update:
// 
// - Call the [WaitForFence] method before encoding commands that need to wait
// for other passes. - Call the [UpdateFence] method after encoding commands
// that later passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/waitForFence(_:)

func (o MTLAccelerationStructureCommandEncoderObject) WaitForFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
	}

// Makes the resources contained in the specified heap available to the
// acceleration structure pass.
//
// heap: A heap that contains resources within an argument buffer.
//
// # Discussion
// 
// This method makes all the resources in the heap resident for the duration
// of a compute pass and ensures that they’re in a format compatible with
// the compute function.
// 
// Call this method before issuing any dispatch calls that may access the
// resources in the heap.
// 
// You can only read or sample resources in the specified heap. This method
// ignores render targets (textures that specify a [TextureUsageRenderTarget]
// usage option) and writable textures (textures that specify a
// [TextureUsageShaderWrite] usage option) within the heap. To use these
// resources, you need to call the [UseResourceUsage] method instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useHeap(_:)

func (o MTLAccelerationStructureCommandEncoderObject) UseHeap(heap MTLHeap) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useHeap:"), heap)
	}

// Makes a resource available to the acceleration structure pass.
//
// resource: A specific resource within an argument buffer.
//
// usage: The options that describe how the compute function uses the resource.
//
// # Discussion
// 
// This method makes the resource resident for the duration of a compute pass
// and ensures that it’s in a format compatible with the compute function.
// 
// Call this method before issuing any dispatch calls that may access the
// resource. Calling this method again, or calling [UseHeap], overwrites any
// previously specified usage options for future dispatch calls within the
// same compute command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useResource(_:usage:)

func (o MTLAccelerationStructureCommandEncoderObject) UseResourceUsage(resource MTLResource, usage MTLResourceUsage) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResource:usage:"), resource, usage)
	}

// Encodes a command to sample hardware counters at this point in the
// acceleration structure pass and store the samples into a counter sample
// buffer.
//
// sampleBuffer: The sample buffer to sample into.
//
// sampleIndex: The index in the counter buffer to write the sample.
//
// barrier: A Boolean value that states whether to insert a barrier before taking the
// sample.
//
// # Discussion
// 
// Inserting a barrier ensures that any work you encoded with this encoder is
// complete before the GPU samples the hardware counters. If you don’t
// insert a barrier, the GPU can sample the counters concurrently with other
// commands encoded by this encoder. Using a barrier leads to more repeatable
// counter results but can negatively impact performance.
// 
// Regardless of whether you set a barrier, the GPU doesn’t isolate the
// sampling from work encoded by other encoders.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)

func (o MTLAccelerationStructureCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), sampleBuffer, sampleIndex, barrier)
	}

// Specifies that an array of heaps containing resources in an argument buffer
// can be safely used by the acceleration structure pass.
//
// heaps: An array of heaps that contains resources within an argument buffer.
//
// count: The number of heaps in the array.
//
// # Discussion
// 
// This method makes all the resources in the array of heaps resident for the
// duration of a compute pass and ensures that they’re in a format
// compatible with the compute function.
// 
// Call this method before issuing any dispatch calls that may access the
// resources in the array of heaps.
// 
// Resources within the specified array of heaps can only be read or sampled
// from. This method ignores render targets (textures that specify a
// [TextureUsageRenderTarget] usage option) and writable textures (textures
// that specify a [TextureUsageShaderWrite] usage option) within the array of
// heaps. To use these resources, you need to call the [UseResourceUsage]
// method instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useHeaps:count:

func (o MTLAccelerationStructureCommandEncoderObject) UseHeapsCount(heaps []MTLHeap, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useHeaps:count:"), objc.CArray(heaps), count)
	}

// Specifies that an array of resources in an argument buffer can be safely
// used by the acceleration structure pass.
//
// resources: An array of resources within an argument buffer.
//
// count: The number of resource elements in `resources`.
//
// usage: Options that indicate how a GPU function accesses each resource in
// `resources`.
//
// # Discussion
// 
// This method makes the array of resources resident for the duration of a
// compute pass and ensures that it’s in a format compatible with the
// compute function.
// 
// Call this method before issuing any dispatch calls that may access the
// array of resources. Calling this method again, or calling [UseHeap],
// overwrites any previously specified usage options for future dispatch calls
// within the same compute command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCommandEncoder/useResources:count:usage:

func (o MTLAccelerationStructureCommandEncoderObject) UseResourcesCountUsage(resources []MTLResource, count uint, usage MTLResourceUsage) {
	
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

func (o MTLAccelerationStructureCommandEncoderObject) EndEncoding() {
	
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

func (o MTLAccelerationStructureCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
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

func (o MTLAccelerationStructureCommandEncoderObject) PushDebugGroup(string_ string) {
	
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

func (o MTLAccelerationStructureCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}

// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device

func (o MTLAccelerationStructureCommandEncoderObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label

func (o MTLAccelerationStructureCommandEncoderObject) Label() string {
	
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

func (o MTLAccelerationStructureCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
	}






func (o MTLAccelerationStructureCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}





