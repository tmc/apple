// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An encoder that encodes commands that modify resource configurations.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder
type MTLResourceStateCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Encodes a command to update the texture mappings for a region in a single texture mipmap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMapping(_:mode:region:mipLevel:slice:)
	UpdateTextureMappingModeRegionMipLevelSlice(texture MTLTexture, mode MTLSparseTextureMappingMode, region MTLRegion, mipLevel uint, slice uint)

	// Encodes a command to update memory mappings for multiple regions inside a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMappings(_:mode:regions:mipLevels:slices:numRegions:)
	UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture MTLTexture, mode MTLSparseTextureMappingMode, regions []MTLRegion, mipLevels uint, slices uint, numRegions uint)

	// Encodes a command to update a texture’s memory mappings, specifying the parameters indirectly.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMapping(_:mode:indirectBuffer:indirectBufferOffset:)
	UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture MTLTexture, mode MTLSparseTextureMappingMode, indirectBuffer MTLBuffer, indirectBufferOffset uint)

	// Encodes a command that instructs the GPU to update a fence, which signals passes waiting on the fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/update(_:)
	UpdateFence(fence MTLFence)

	// Encodes a command that instructs the GPU to pause before starting the resource state commands until another pass updates a fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/wait(for:)
	WaitForFence(fence MTLFence)

	// MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/moveTextureMappings(sourceTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:destinationTexture:destinationSlice:destinationLevel:destinationOrigin:)
	MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin)
}

// MTLResourceStateCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLResourceStateCommandEncoder protocol.
type MTLResourceStateCommandEncoderObject struct {
	objectivec.Object
}
func (o MTLResourceStateCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLResourceStateCommandEncoderObjectFromID constructs a [MTLResourceStateCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLResourceStateCommandEncoderObjectFromID(id objc.ID) MTLResourceStateCommandEncoderObject {
	return MTLResourceStateCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes a command to update the texture mappings for a region in a single
// texture mipmap.
//
// texture: The sparse texture to update.
//
// mode: A mode that indicates whether the method allocates or frees a memory tile
// in the texture.
//
// region: A region, in tile coordinates, that describes the part of the mipmap to
// update.
//
// mipLevel: The mipmap to update.
//
// slice: The slice in the texture to update.
//
// # Discussion
// 
// When the GPU executes the command that updates the texture’s memory
// mapping, the GPU gets details about the region from the `region` parameter.
// 
// To allocate tiles from the heap, pass [SparseTextureMappingModeMap] as the
// `mode` parameter, and to free files back to the heap, pass
// [SparseTextureMappingModeUnmap].
// 
// If you encode other commands that use the texture’s contents, such as
// rendering to the texture or sampling from a texture, synchronize the
// texture’s mapping updates with those commands to avoid race conditions.
// See [Resource synchronization].
// 
// If you encode commands with multiple resource state passes, synchronize the
// resources to run the commands in the passes sequentially. See the
// [MTLResourceStateCommandEncoder] protocol.
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMapping(_:mode:region:mipLevel:slice:)
func (o MTLResourceStateCommandEncoderObject) UpdateTextureMappingModeRegionMipLevelSlice(texture MTLTexture, mode MTLSparseTextureMappingMode, region MTLRegion, mipLevel uint, slice uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateTextureMapping:mode:region:mipLevel:slice:"), texture, mode, region, mipLevel, slice)
	}
// Encodes a command to update memory mappings for multiple regions inside a
// texture.
//
// texture: The sparse texture to update.
//
// mode: The change to make to the texture mapping.
//
// regions: A pointer to an array of regions to change. You need to provide as many
// regions as you specify in the `numRegions` parameter.
//
// mipLevels: A pointer to an array of mipmap levels to change. You need to provide as
// many entries as you specify in the `numRegions` parameter.
//
// slices: A pointer to an array of slices to change. You need to provide as many
// entries as you specify in the `numRegions` parameter.
//
// numRegions: The number of regions to update.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMappings(_:mode:regions:mipLevels:slices:numRegions:)
func (o MTLResourceStateCommandEncoderObject) UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture MTLTexture, mode MTLSparseTextureMappingMode, regions []MTLRegion, mipLevels uint, slices uint, numRegions uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"), texture, mode, regions, mipLevels, slices, numRegions)
	}
// Encodes a command to update a texture’s memory mappings, specifying the
// parameters indirectly.
//
// texture: The sparse texture to update.
//
// mode: A mode that indicates whether the method allocates or frees a memory tile
// in the texture.
//
// indirectBuffer: A buffer that contains an array of mapping arguments that are instances of
// the [MTLMapIndirectArguments] structure.
// //
// [MTLMapIndirectArguments]: https://developer.apple.com/documentation/Metal/MTLMapIndirectArguments
//
// indirectBufferOffset: The offset, in bytes, where the first argument begins in the
// `indirectBuffer` parameter.
//
// # Discussion
// 
// When the GPU executes the command that updates the texture’s memory
// mapping, the GPU gets details about the region to update from the
// `indirectBuffer` parameter.
// 
// To allocate tiles from the heap, pass [SparseTextureMappingModeMap] as the
// `mode` parameter, and to free files back to the heap, pass
// [SparseTextureMappingModeUnmap].
// 
// If you encode other commands that use the texture’s contents, such as
// rendering to the texture or sampling from a texture, synchronize the
// texture’s mapping updates with those commands to avoid race conditions.
// See [Resource synchronization].
// 
// If you encode commands with multiple resource state passes, synchronize the
// resources to run the commands in the passes sequentially. See the
// [MTLResourceStateCommandEncoder] protocol.
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/updateTextureMapping(_:mode:indirectBuffer:indirectBufferOffset:)
func (o MTLResourceStateCommandEncoderObject) UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture MTLTexture, mode MTLSparseTextureMappingMode, indirectBuffer MTLBuffer, indirectBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"), texture, mode, indirectBuffer, indirectBufferOffset)
	}
// Encodes a command that instructs the GPU to update a fence, which signals
// passes waiting on the fence.
//
// fence: An [MTLFence] instance to update.
//
// # Discussion
// 
// Fences maintain order to prevent GPU data hazards as the GPU runs various
// passes within the same command queue. This encoded command notifies any
// passes waiting for `fence`.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/update(_:)
func (o MTLResourceStateCommandEncoderObject) UpdateFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
	}
// Encodes a command that instructs the GPU to pause before starting the
// resource state commands until another pass updates a fence.
//
// fence: An [MTLFence] instance to pause execution on until updated.
//
// # Discussion
// 
// Fences maintain order to prevent GPU data hazards as the GPU runs various
// passes within the same command queue. The encoded resource state commands
// wait for a pass to update `fence` before running.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/wait(for:)
func (o MTLResourceStateCommandEncoderObject) WaitForFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStateCommandEncoder/moveTextureMappings(sourceTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:destinationTexture:destinationSlice:destinationLevel:destinationOrigin:)
func (o MTLResourceStateCommandEncoderObject) MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin) {
	
	objc.Send[struct{}](o.ID, objc.Sel("moveTextureMappingsFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:"), sourceTexture, sourceSlice, sourceLevel, sourceOrigin, sourceSize, destinationTexture, destinationSlice, destinationLevel, destinationOrigin)
	}
// Declares that all command generation from the encoder is completed.
//
// # Discussion
// 
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLResourceStateCommandEncoderObject) EndEncoding() {
	
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
func (o MTLResourceStateCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
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
func (o MTLResourceStateCommandEncoderObject) PushDebugGroup(string_ string) {
	
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
func (o MTLResourceStateCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}
// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLResourceStateCommandEncoderObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLResourceStateCommandEncoderObject) Label() string {
	
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
func (o MTLResourceStateCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
	}

func (o MTLResourceStateCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

