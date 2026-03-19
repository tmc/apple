// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Encodes commands that copy and modify resources for a single blit pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder
type MTLBlitCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Encodes a command that generates mipmaps for a texture from the base mipmap level up to the highest mipmap level.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/generateMipmaps(for:)
	GenerateMipmapsForTexture(texture MTLTexture)

	// Encodes a command that copies data from one buffer into another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:to:destinationOffset:size:)
	CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer MTLBuffer, sourceOffset uint, destinationBuffer MTLBuffer, destinationOffset uint, size uint)

	// Encodes a command that copies data from one texture to another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:to:)
	CopyFromTextureToTexture(sourceTexture MTLTexture, destinationTexture MTLTexture)

	// Encodes a command that copies slices of a texture to another texture’s slices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:to:destinationSlice:destinationLevel:sliceCount:levelCount:)
	CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint)

	// Encodes a command that copies image data from a texture’s slice into another slice.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin)

	// Encodes a command to copy data from a slice of one tensor into a slice of another tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOrigin:sourceDimensions:to:destinationOrigin:destinationDimensions:)
	CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor MTLTensor, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor MTLTensor, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents)

	// Encodes a command to copy image data from a source buffer into a destination texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:)
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceBuffer MTLBuffer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin)

	// Encodes a command to copy image data from a source buffer into a destination texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:options:)
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions(sourceBuffer MTLBuffer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin, options MTLBlitOption)

	// Encodes a command that copies image data from a texture slice to a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationBuffer MTLBuffer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint)

	// Encodes a command that copies image data from a texture slice to a buffer, and provides options for special texture formats.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:options:)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationBuffer MTLBuffer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint, options MTLBlitOption)

	// Encodes a command that improves the performance of GPU memory operations with a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForGPUAccess(texture:)
	OptimizeContentsForGPUAccess(texture MTLTexture)

	// Encodes a command that improves the performance of GPU memory operations with a specific portion of a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForGPUAccess(texture:slice:level:)
	OptimizeContentsForGPUAccessSliceLevel(texture MTLTexture, slice uint, level uint)

	// Encodes a command that improves the performance of CPU memory operations with a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForCPUAccess(texture:)
	OptimizeContentsForCPUAccess(texture MTLTexture)

	// Encodes a command that improves the performance of CPU memory operations with a specific portion of a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForCPUAccess(texture:slice:level:)
	OptimizeContentsForCPUAccessSliceLevel(texture MTLTexture, slice uint, level uint)

	// Encodes a command that synchronizes the CPU’s copy of a managed resource, such as a buffer or texture, so that it matches the GPU’s copy.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/synchronize(resource:)
	SynchronizeResource(resource MTLResource)

	// Encodes a command that synchronizes a part of the CPU’s copy of a texture so that it matches the GPU’s copy.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/synchronize(texture:slice:level:)
	SynchronizeTextureSliceLevel(texture MTLTexture, slice uint, level uint)

	// Encodes a command that instructs the GPU to pause the blit pass until another pass updates a fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/waitForFence(_:)
	WaitForFence(fence MTLFence)

	// Encodes a command that instructs the GPU to update a fence after the blit pass completes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/updateFence(_:)
	UpdateFence(fence MTLFence)

	// Encodes a command that samples the GPU’s hardware counters during a blit pass and stores the data in a counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool)

	// Encodes a command that retrieves a sparse texture’s access data for a specific region, mipmap level, and slice.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/getTextureAccessCounters(_:region:mipLevel:slice:resetCounters:countersBuffer:countersBufferOffset:)
	GetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset(texture MTLTexture, region MTLRegion, mipLevel uint, slice uint, resetCounters bool, countersBuffer MTLBuffer, countersBufferOffset uint)

	// Encodes a command that resets a sparse texture’s access data for a specific region, mipmap level, and slice.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resetTextureAccessCounters(_:region:mipLevel:slice:)
	ResetTextureAccessCountersRegionMipLevelSlice(texture MTLTexture, region MTLRegion, mipLevel uint, slice uint)

	// Encodes a command that copies commands from one indirect command buffer into another.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:
	CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source MTLIndirectCommandBuffer, sourceRange foundation.NSRange, destination MTLIndirectCommandBuffer, destinationIndex uint)

	// Encodes a command that fills a buffer with a constant value for each byte.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/fillBuffer:range:value:
	FillBufferRangeValue(buffer MTLBuffer, range_ foundation.NSRange, value uint8)

	// Encodes a command that can improve the performance of a range of commands within an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeIndirectCommandBuffer:withRange:
	OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, range_ foundation.NSRange)

	// Encodes a command that resets a range of commands in an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resetCommandsInBuffer:withRange:
	ResetCommandsInBufferWithRange(buffer MTLIndirectCommandBuffer, range_ foundation.NSRange)

	// Encodes a command that resolves the data from the samples in a sample counter buffer and stores the results into a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resolveCounters:inRange:destinationBuffer:destinationOffset:
	ResolveCountersInRangeDestinationBufferDestinationOffset(sampleBuffer MTLCounterSampleBuffer, range_ foundation.NSRange, destinationBuffer MTLBuffer, destinationOffset uint)
}

// MTLBlitCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLBlitCommandEncoder protocol.
type MTLBlitCommandEncoderObject struct {
	objectivec.Object
}
func (o MTLBlitCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLBlitCommandEncoderObjectFromID constructs a [MTLBlitCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLBlitCommandEncoderObjectFromID(id objc.ID) MTLBlitCommandEncoderObject {
	return MTLBlitCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes a command that generates mipmaps for a texture from the base mipmap
// level up to the highest mipmap level.
//
// texture: A texture instance the command generates mipmaps for that has:
// 
// - A [MipmapLevelCount] property that’s greater than `1` - A [PixelFormat]
// that’s color-renderable and color-filterable
//
// # Discussion
// 
// The command generates with scaled images for all levels up to the highest
// mipmap level.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/generateMipmaps(for:)
func (o MTLBlitCommandEncoderObject) GenerateMipmapsForTexture(texture MTLTexture) {
	
	objc.Send[struct{}](o.ID, objc.Sel("generateMipmapsForTexture:"), texture)
	}
// Encodes a command that copies data from one buffer into another.
//
// sourceBuffer: A buffer the command copies data from.
//
// sourceOffset: A byte offset within `sourceBuffer` the command copies from. In macOS,
// `sourceOffset` needs to be a multiple of `4`, but can be any value in iOS
// and tvOS.
//
// destinationBuffer: The destination buffer for the copy operation.
//
// destinationOffset: A byte offset within `destinationBuffer` the command copies to. In macOS,
// `destinationOffset` needs to be a multiple of `4`, but can be any value in
// iOS and tvOS.
//
// size: The number of bytes the command copies from `sourceBuffer` to
// `destinationBuffer`. In macOS, `size` needs to be a multiple of `4`, but
// can be any value in iOS and tvOS.
//
// # Discussion
// 
// You can pass the same buffer to the `sourceBuffer` and `destinationBuffer`
// parameters if `size` is less than the distance between `sourceOffset` and
// `destinationOffset`.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:to:destinationOffset:size:)
func (o MTLBlitCommandEncoderObject) CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer MTLBuffer, sourceOffset uint, destinationBuffer MTLBuffer, destinationOffset uint, size uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromBuffer:sourceOffset:toBuffer:destinationOffset:size:"), sourceBuffer, sourceOffset, destinationBuffer, destinationOffset, size)
	}
// Encodes a command that copies data from one texture to another.
//
// sourceTexture: A texture the command copies data from.
//
// destinationTexture: Another texture the command copies the data to that has the same pixel
// format and sample count as `sourceTexture`.
//
// # Discussion
// 
// The textures can be different sizes as long as the larger texture has a
// mipmap level that’s the same size as the smaller texture’s level `0`
// mipmap.
// 
// The command copies all identical mipmap sizes. If both textures are arrays,
// the command copies as many texture slices (array elements) as possible.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:to:)
func (o MTLBlitCommandEncoderObject) CopyFromTextureToTexture(sourceTexture MTLTexture, destinationTexture MTLTexture) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:toTexture:"), sourceTexture, destinationTexture)
	}
// Encodes a command that copies slices of a texture to another texture’s
// slices.
//
// sourceTexture: A texture the command copies data from.
//
// sourceSlice: A slice within `sourceTexture` the command uses as a starting point to copy
// data from.
// 
// Set this to `0` if `sourceTexture` isn’t a texture array or a cube
// texture.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// destinationTexture: Another texture the command copies the data to that has the same pixel
// format and sample count as `sourceTexture`.
//
// destinationSlice: A slice within `destinationTexture` the command uses as its starting point
// for coping data.
// 
// Set this to `0` if `destinationTexture` isn’t a texture array or a cube
// texture.
//
// destinationLevel: A mipmap level within `destinationTexture` that has the same size as the
// source texture’s `sourceLevel` mipmap.
//
// sliceCount: The number of slices the command copies so that it satisfies these
// conditions:
// 
// - The sum of `sourceLevel` and `sourceSlice` doesn’t exceed the number of
// slices in `sourceTexture`. - The sum of `destinationLevel` and
// `destinationSlice` doesn’t exceed the number of slices in
// `destinationTexture`.
//
// levelCount: The number of mipmap levels the command copies so that it satisfies these
// conditions:
// 
// - The sum of `levelCount` and `sourceLevel` doesn’t exceed the number of
// mipmap levels in `sourceTexture`. - The sum of `levelCount` and
// `destinationLevel` doesn’t exceed the number of mipmap levels in
// `destinationTexture`.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:to:destinationSlice:destinationLevel:sliceCount:levelCount:)
func (o MTLBlitCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:toTexture:destinationSlice:destinationLevel:sliceCount:levelCount:"), sourceTexture, sourceSlice, sourceLevel, destinationTexture, destinationSlice, destinationLevel, sliceCount, levelCount)
	}
// Encodes a command that copies image data from a texture’s slice into
// another slice.
//
// sourceTexture: A texture with an [IsFramebufferOnly] property value of [false] that the
// command copies data from.
// 
// For a texture that uses a compressed pixel format, align the copy region
// (`sourceOrigin` and `sourceSize`) to the pixel format’s block size.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// sourceSlice: A slice within `sourceTexture`.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// sourceOrigin: A location within `sourceTexture` that the command begins copying data
// from.
// 
// Assign `0` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the origin’s [z] property to
// `0`. - If the source texture is a 1D texture, set the origin’s [y] and
// [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// sourceSize: An [MTLSize] instance, which can represent a 3D region, that instructs the
// command how many pixels to copy from `sourceTexture`, starting at
// `sourceOrigin`.
// 
// Assign `1` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the size’s [depth] property
// to `1`. - If the source texture is a 1D texture, set the size’s [height]
// and [depth] properties to `1`.
// 
// If `sourceTexture` uses a compressed pixel format, set `sourceSize` to a
// multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
//
// destinationTexture: A texture the command copies data to that has the following configuration:
// 
// - The [IsFramebufferOnly] property value is [false]. - The pixel format is
// the same as `sourceTexture`. - The sample count is the same as
// `sourceTexture`.
// 
// For a texture that uses a compressed pixel format, align the copy region
// (`destinationOrigin`) to the pixel format’s block size.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// destinationSlice: A slice within `destinationTexture`.
//
// destinationLevel: A mipmap level within `destinationTexture`.
//
// destinationOrigin: A location within `destinationTexture` that the command begins copying data
// to.
// 
// Assign `0` to each dimension that’s not relevant to `destinationTexture`.
// For example:
// 
// - If the destination texture is a 2D texture, set the origin’s [z]
// property to `0`. - If the destination texture is a 1D texture, set the
// origin’s [y] and [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// # Discussion
// 
// For textures that use a PVRTC pixel format, you can use this method to copy
// the entire texture, but not a subregion of the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:)
func (o MTLBlitCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:"), sourceTexture, sourceSlice, sourceLevel, sourceOrigin, sourceSize, destinationTexture, destinationSlice, destinationLevel, destinationOrigin)
	}
// Encodes a command to copy data from a slice of one tensor into a slice of
// another tensor.
//
// sourceTensor: A tensor instance that this command copies data from.
//
// sourceOrigin: An array of offsets, in elements, to the first element of the slice of
// `sourceTensor` that this command copies data from.
//
// sourceDimensions: An array of sizes, in elements, of the slice `sourceTensor` that this
// command copies data from.
//
// destinationTensor: A tensor instance that this command copies data to.
//
// destinationOrigin: An array of offsets, in elements, to the first element of the slice of
// `destinationTensor` that this command copies data to.
//
// destinationDimensions: An array of sizes, in elements, of the slice of `destinationTensor` that
// this command copies data to.
//
// # Discussion
// 
// This command applies reshapes if `sourceTensor` and `destinationTensor` are
// not aliasable.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOrigin:sourceDimensions:to:destinationOrigin:destinationDimensions:)
func (o MTLBlitCommandEncoderObject) CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor MTLTensor, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor MTLTensor, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTensor:sourceOrigin:sourceDimensions:toTensor:destinationOrigin:destinationDimensions:"), sourceTensor, sourceOrigin, sourceDimensions, destinationTensor, destinationOrigin, destinationDimensions)
	}
// Encodes a command to copy image data from a source buffer into a
// destination texture.
//
// sourceBuffer: A buffer the command copies data from.
//
// sourceOffset: A byte offset within `sourceBuffer` that the command copies from, which
// needs to be a multiple of the destination texture’s pixel size, in bytes.
//
// sourceBytesPerRow: The number of bytes between adjacent rows of pixels in the source
// buffer’s memory, which needs to be:
// 
// - A multiple of the source texture’s pixel size, in bytes - Less than or
// equal to the product of the destination texture’s pixel size, in bytes,
// and the largest pixel width the destination texture’s type allows
// 
// If `destinationTexture` uses a compressed pixel format, set
// `sourceBytesPerRow` to the number of bytes between the starts of two row
// blocks.
//
// sourceBytesPerImage: The number of bytes between each 2D image of a 3D texture. This value needs
// to be a multiple of the source texture’s pixel size, in bytes.
// 
// Set this value to `0` for 2D textures, which means `sourceSize.`[depth] is
// equal to `1`.
// //
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
//
// sourceSize: An [MTLSize] instance, which can represent a 3D region, that instructs the
// command how many pixels to copy to `destinationTexture`, starting at
// `destinationOrigin`.
// 
// Assign `1` to each dimension that’s not relevant to `destinationTexture`.
// For example:
// 
// - If the destination texture is a 2D texture, set the size’s [depth]
// property to `1`. - If the destination texture is a 1D texture, set the
// size’s [height] and [depth] properties to `1`.
// 
// If `destinationTexture` uses a compressed pixel format, set `sourceSize` to
// a multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
//
// destinationTexture: A texture with an [IsFramebufferOnly] property value of [false] that the
// command copies data to.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// destinationSlice: A slice within `destinationTexture`.
// 
// For textures that use a combined depth/stencil pixel format, call the
// [CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions]
// method instead. Configure that method’s `options` parameter
// appropriately.
//
// destinationLevel: A mipmap level within `destinationTexture`.
//
// destinationOrigin: A location within `destinationTexture` that the command begins copying data
// to.
// 
// Assign `0` to each dimension that’s not relevant to `destinationTexture`.
// For example:
// 
// - If the destination texture is a 2D texture, set the origin’s [z]
// property to `0`. - If the destination texture is a 1D texture, set the
// origin’s [y] and [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// # Discussion
// 
// This method is the equivalent of passing an empty [OptionSet] to the
// `options` parameter of
// [CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions].
// In Swift, pass `[]` to represent an empty option set, and in Objective-C,
// pass [BlitOptionNone].
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:)
func (o MTLBlitCommandEncoderObject) CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceBuffer MTLBuffer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromBuffer:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:"), sourceBuffer, sourceOffset, sourceBytesPerRow, sourceBytesPerImage, sourceSize, destinationTexture, destinationSlice, destinationLevel, destinationOrigin)
	}
// Encodes a command to copy image data from a source buffer into a
// destination texture.
//
// sourceBuffer: A buffer the command copies data from.
//
// sourceOffset: A byte offset within `sourceBuffer` that the command copies from, which
// needs to be a multiple of the destination texture’s pixel size, in bytes.
//
// sourceBytesPerRow: The number of bytes between adjacent rows of pixels in the source
// buffer’s memory, which needs to be:
// 
// - A multiple of the source texture’s pixel size, in bytes - Less than or
// equal to the product of the destination texture’s pixel size, in bytes,
// and the largest pixel width the destination texture’s type allows
// 
// If `destinationTexture` uses a compressed pixel format, set
// `sourceBytesPerRow` to the number of bytes between the starts of two row
// blocks.
//
// sourceBytesPerImage: The number of bytes between each 2D image of a 3D texture. This value needs
// to be a multiple of the source texture’s pixel size, in bytes.
// 
// Set this value to `0` for 2D textures, which means `sourceSize.`[depth] is
// equal to `1`.
// //
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
//
// sourceSize: An [MTLSize] instance, which can represent a 3D region, that instructs the
// command how many pixels to copy to `destinationTexture`, starting at
// `destinationOrigin`.
// 
// Assign `1` to each dimension that’s not relevant to `destinationTexture`.
// For example:
// 
// - If the destination texture is a 2D texture, set the size’s [depth]
// property to `1`. - If the destination texture is a 1D texture, set the
// size’s [height] and [depth] properties to `1`.
// 
// If `destinationTexture` uses a compressed pixel format, set `sourceSize` to
// a multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
//
// destinationTexture: A texture with an [IsFramebufferOnly] property value of [false] that the
// command copies data to.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// destinationSlice: A slice within `destinationTexture`.
// 
// For textures that use a combined depth/stencil pixel format, configure the
// `options` parameter appropriately.
//
// destinationLevel: A mipmap level within `destinationTexture`.
//
// destinationOrigin: A location within `destinationTexture` that the command begins copying data
// to.
// 
// Assign `0` to each dimension that’s not relevant to `destinationTexture`.
// For example:
// 
// - If the destination texture is a 2D texture, set the origin’s [z]
// property to `0`. - If the destination texture is a 1D texture, set the
// origin’s [y] and [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// options: An option set that applies to textures with applicable pixel formats, such
// as combined depth/stencil or PVRTC formats.
// 
// If the texture’s pixel format is a combined depth/stencil format, set
// `options` to either [BlitOptionDepthFromDepthStencil] or
// [BlitOptionStencilFromDepthStencil], but not both.
// 
// If the texture’s pixel format is a PVRTC format, set `options` to
// [BlitOptionRowLinearPVRTC].
//
// # Discussion
// 
// Passing an empty [OptionSet] to the `options` parameter is the equivalent
// of calling
// [CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin].
// In Swift, pass `[]` to represent an empty option set, and in Objective-C,
// pass [BlitOptionNone].
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:to:destinationSlice:destinationLevel:destinationOrigin:options:)
func (o MTLBlitCommandEncoderObject) CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions(sourceBuffer MTLBuffer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize MTLSize, destinationTexture MTLTexture, destinationSlice uint, destinationLevel uint, destinationOrigin MTLOrigin, options MTLBlitOption) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromBuffer:sourceOffset:sourceBytesPerRow:sourceBytesPerImage:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:options:"), sourceBuffer, sourceOffset, sourceBytesPerRow, sourceBytesPerImage, sourceSize, destinationTexture, destinationSlice, destinationLevel, destinationOrigin, options)
	}
// Encodes a command that copies image data from a texture slice to a buffer.
//
// sourceTexture: A texture with an [IsFramebufferOnly] property value of [false] that the
// command copies data from.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// sourceSlice: A slice within `sourceTexture`.
// 
// For textures that use a combined depth/stencil pixel format, call the
// [CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions]
// method instead. Configure that method’s `options` parameter
// appropriately.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// sourceOrigin: A location within `sourceTexture` that the command begins copying data
// from.
// 
// Assign `0` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the origin’s [z] property to
// `0`. - If the source texture is a 1D texture, set the origin’s [y] and
// [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// sourceSize: An [MTLSize] instance, which can represent a 3D region, that instructs the
// command how many pixels to copy from `sourceTexture`, starting at
// `sourceOrigin`.
// 
// Assign `1` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the size’s [depth] property
// to `1`. - If the source texture is a 1D texture, set the size’s [height]
// and [depth] properties to `1`.
// 
// If `sourceTexture` uses a compressed pixel format, set `sourceSize` to a
// multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
//
// destinationBuffer: A buffer the command copies data to.
//
// destinationOffset: A byte offset within `destinationBuffer` the command copies to, which needs
// to be a multiple of the source texture’s pixel size, in bytes.
//
// destinationBytesPerRow: The number of bytes between adjacent rows of pixels in the destination
// buffer’s memory, which needs to be:
// 
// - A multiple of the source texture’s pixel size, in bytes - Less than or
// equal to the product of the source texture’s pixel size, in bytes, and
// the largest pixel width the source texture’s type allows
// 
// If `sourceTexture` uses a compressed pixel format, set
// `destinationBytesPerRow` to the number of bytes between the starts of two
// row blocks.
//
// destinationBytesPerImage: The number of bytes between each 2D image of a 3D texture. This value needs
// to be a multiple of the source texture’s pixel size, in bytes.
// 
// Set this value to `0` for 2D textures, which means `sourceSize.`[depth] is
// equal to `1`.
// //
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
//
// # Discussion
// 
// This method is the equivalent of passing an empty [OptionSet] to the
// `options` parameter of
// [CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions].
// In Swift, pass `[]` to represent an empty option set, and in Objective-C,
// pass [BlitOptionNone].
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:)
func (o MTLBlitCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationBuffer MTLBuffer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toBuffer:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:"), sourceTexture, sourceSlice, sourceLevel, sourceOrigin, sourceSize, destinationBuffer, destinationOffset, destinationBytesPerRow, destinationBytesPerImage)
	}
// Encodes a command that copies image data from a texture slice to a buffer,
// and provides options for special texture formats.
//
// sourceTexture: A texture with an [IsFramebufferOnly] property value of [false] that the
// command copies data from.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// sourceSlice: A slice within `sourceTexture`.
// 
// For textures that use a combined depth/stencil pixel format, configure the
// `options` parameter appropriately.
//
// sourceLevel: A mipmap level within `sourceTexture`.
//
// sourceOrigin: A location within `sourceTexture` that the command begins copying data
// from.
// 
// Assign `0` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the origin’s [z] property to
// `0`. - If the source texture is a 1D texture, set the origin’s [y] and
// [z] properties to `0`.
// //
// [y]: https://developer.apple.com/documentation/Metal/MTLOrigin/y
// [z]: https://developer.apple.com/documentation/Metal/MTLOrigin/z
//
// sourceSize: An [MTLSize] instance, which can represent a 3D region, that instructs the
// command how many pixels to copy from `sourceTexture`, starting at
// `sourceOrigin`.
// 
// Assign `1` to each dimension that’s not relevant to `sourceTexture`. For
// example:
// 
// - If the source texture is a 2D texture, set the size’s [depth] property
// to `1`. - If the source texture is a 1D texture, set the size’s [height]
// and [depth] properties to `1`.
// 
// If `sourceTexture` uses a compressed pixel format, set `sourceSize` to a
// multiple of the pixel format’s block size. If the block extends outside
// the bounds of the texture, clamp `sourceSize` to the edge of the texture.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
//
// destinationBuffer: A buffer the command copies data to.
//
// destinationOffset: A byte offset within `destinationBuffer` the command copies to, which needs
// to be a multiple of the source texture’s pixel size, in bytes.
//
// destinationBytesPerRow: The number of bytes between adjacent rows of pixels in the destination
// buffer’s memory, which needs to be:
// 
// - A multiple of the source texture’s pixel size, in bytes - Less than or
// equal to the product of the source texture’s pixel size, in bytes, and
// the largest pixel width the source texture’s type allows
// 
// If `sourceTexture` uses a compressed pixel format, set
// `destinationBytesPerRow` to the number of bytes between the starts of two
// row blocks.
//
// destinationBytesPerImage: The number of bytes between each 2D image of a 3D texture. This value needs
// to be a multiple of the source texture’s pixel size, in bytes.
// 
// Set this value to `0` for 2D textures, which means `sourceSize.`[depth] is
// equal to `1`.
// //
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
//
// options: An option set that applies to textures with applicable pixel formats, such
// as combined depth/stencil or PVRTC formats.
// 
// If the texture’s pixel format is a combined depth/stencil format, set
// `options` to either [BlitOptionDepthFromDepthStencil] or
// [BlitOptionStencilFromDepthStencil], but not both.
// 
// If the texture’s pixel format is a PVRTC format, set `options` to
// [BlitOptionRowLinearPVRTC].
//
// # Discussion
// 
// Passing an empty [OptionSet] to the `options` parameter is the equivalent
// of calling
// [CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage].
// In Swift, pass `[]` to represent an empty option set, and in Objective-C,
// pass [BlitOptionNone].
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copy(from:sourceSlice:sourceLevel:sourceOrigin:sourceSize:to:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:options:)
func (o MTLBlitCommandEncoderObject) CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions(sourceTexture MTLTexture, sourceSlice uint, sourceLevel uint, sourceOrigin MTLOrigin, sourceSize MTLSize, destinationBuffer MTLBuffer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint, options MTLBlitOption) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toBuffer:destinationOffset:destinationBytesPerRow:destinationBytesPerImage:options:"), sourceTexture, sourceSlice, sourceLevel, sourceOrigin, sourceSize, destinationBuffer, destinationOffset, destinationBytesPerRow, destinationBytesPerImage, options)
	}
// Encodes a command that improves the performance of GPU memory operations
// with a texture.
//
// texture: A texture the command optimizes.
//
// # Discussion
// 
// This command can reduce the time it takes the GPU to access a texture. Apps
// typically run the command for:
// 
// - Textures the GPU accesses for an extended period of time - Textures with
// a [StorageMode] property that’s [StorageModeShared] or
// [StorageModeManaged]
// 
// When a blit pass runs this command, the GPU only applies lossless changes
// to the texture’s underlying data.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForGPUAccess(texture:)
func (o MTLBlitCommandEncoderObject) OptimizeContentsForGPUAccess(texture MTLTexture) {
	
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForGPUAccess:"), texture)
	}
// Encodes a command that improves the performance of GPU memory operations
// with a specific portion of a texture.
//
// texture: A texture the command optimizes.
//
// slice: A slice within `texture`.
//
// level: A mipmap level within `texture`.
//
// # Discussion
// 
// This command can reduce the time it takes the GPU to access a texture. Apps
// typically run the command for:
// 
// - Textures the GPU accesses for an extended period of time - Textures with
// a [StorageMode] property that’s [StorageModeShared] or
// [StorageModeManaged]
// 
// When a blit pass runs this command, the GPU only applies lossless changes
// to the texture’s underlying data.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForGPUAccess(texture:slice:level:)
func (o MTLBlitCommandEncoderObject) OptimizeContentsForGPUAccessSliceLevel(texture MTLTexture, slice uint, level uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForGPUAccess:slice:level:"), texture, slice, level)
	}
// Encodes a command that improves the performance of CPU memory operations
// with a texture.
//
// texture: A texture the command optimizes.
//
// # Discussion
// 
// This command can reduce the time it takes the CPU to access a texture. Apps
// typically run the command for:
// 
// - Textures the CPU accesses for an extended period of time - Textures with
// a [StorageMode] property that’s [StorageModeShared] or
// [StorageModeManaged]
// 
// When a blit pass runs this command, the GPU only applies lossless changes
// to the texture’s underlying data.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForCPUAccess(texture:)
func (o MTLBlitCommandEncoderObject) OptimizeContentsForCPUAccess(texture MTLTexture) {
	
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForCPUAccess:"), texture)
	}
// Encodes a command that improves the performance of CPU memory operations
// with a specific portion of a texture.
//
// texture: A texture the command optimizes.
//
// slice: A slice within `texture`.
//
// level: A mipmap level within `texture`.
//
// # Discussion
// 
// This command can reduce the time it takes the CPU to access a texture. Apps
// typically run the command for:
// 
// - Textures the CPU accesses for an extended period of time - Textures with
// a [StorageMode] property that’s [StorageModeShared] or
// [StorageModeManaged]
// 
// When a blit pass runs this command, the GPU only applies lossless changes
// to the texture’s underlying data.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeContentsForCPUAccess(texture:slice:level:)
func (o MTLBlitCommandEncoderObject) OptimizeContentsForCPUAccessSliceLevel(texture MTLTexture, slice uint, level uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("optimizeContentsForCPUAccess:slice:level:"), texture, slice, level)
	}
// Encodes a command that synchronizes the CPU’s copy of a managed resource,
// such as a buffer or texture, so that it matches the GPU’s copy.
//
// resource: An [MTLResource] instance — such as an [MTLBuffer] or [MTLTexture] —
// with a [StorageMode] property that’s equal to [StorageModeManaged].
//
// # Discussion
// 
// This method ensures the CPU can correctly read all the changes a GPU makes
// to a resource that uses the managed storage mode. For the resources you
// create with [StorageModeManaged], the CPU and GPU each have a copy of that
// resource. As the GPU modifies its copy, the CPU’s copy remains unchanged
// until you synchronize with a command, such as this one.
// 
// The CPU can access the updated content from its copy of the resources after
// the synchronization command completes.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/synchronize(resource:)
func (o MTLBlitCommandEncoderObject) SynchronizeResource(resource MTLResource) {
	
	objc.Send[struct{}](o.ID, objc.Sel("synchronizeResource:"), resource)
	}
// Encodes a command that synchronizes a part of the CPU’s copy of a texture
// so that it matches the GPU’s copy.
//
// texture: An [MTLTexture] instance with a [StorageMode] property that’s equal to
// [StorageModeManaged].
//
// slice: A slice within `texture`.
//
// level: A mipmap level within `texture`.
//
// # Discussion
// 
// This method ensures the CPU can correctly read the changes a GPU makes to a
// slice of a texture that uses the managed storage mode. For the resources
// you create with [StorageModeManaged], the CPU and GPU each have a copy of
// that resource. As the GPU modifies its copy, the CPU’s copy remains
// unchanged until you synchronize with a command, such as this one.
// 
// The CPU can access the updated content from its copy of the texture after
// the synchronization command completes.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/synchronize(texture:slice:level:)
func (o MTLBlitCommandEncoderObject) SynchronizeTextureSliceLevel(texture MTLTexture, slice uint, level uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("synchronizeTexture:slice:level:"), texture, slice, level)
	}
// Encodes a command that instructs the GPU to pause the blit pass until
// another pass updates a fence.
//
// fence: A fence that the pass waits for before it runs any of its commands.
//
// # Discussion
// 
// You can synchronize memory operations of a blit pass that access resources
// with an [MTLFence]. This method instructs the GPU to wait until another
// pass updates `fence` before running the blit pass. The fence indicates when
// the pass can access those resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a blit pass that reuses a fence, wait for other passes to
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
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/waitForFence(_:)
func (o MTLBlitCommandEncoderObject) WaitForFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:"), fence)
	}
// Encodes a command that instructs the GPU to update a fence after the blit
// pass completes.
//
// fence: A fence the pass updates after it completes.
//
// # Discussion
// 
// You can synchronize memory operations of a blit pass that access resources
// with an [MTLFence]. This method instructs the pass to update `fence` after
// it runs all its memory store operations to the resources it accesses. The
// fence indicates when other passes can access those resources without a race
// condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a blit pass that reuses a fence, wait for other passes to
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
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/updateFence(_:)
func (o MTLBlitCommandEncoderObject) UpdateFence(fence MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:"), fence)
	}
// Encodes a command that samples the GPU’s hardware counters during a blit
// pass and stores the data in a counter sample buffer.
//
// sampleBuffer: A counter sample buffer where the command stores the sample data.
//
// sampleIndex: A location within `sampleBuffer` where the command stores the sample data.
//
// barrier: A Boolean value that indicates whether the command inserts a barrier before
// taking the sample.
//
// # Discussion
// 
// Inserting a barrier ensures that any work you encode with this encoder is
// complete before the GPU samples the hardware counters. If you don’t
// insert a barrier, the GPU can sample the counters concurrently with other
// commands you encode with this encoder. Using a barrier can help the counter
// results be more predictable and repeatable, but it may adversely affect
// your app’s runtime performance.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/sampleCounters(sampleBuffer:sampleIndex:barrier:)
func (o MTLBlitCommandEncoderObject) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer MTLCounterSampleBuffer, sampleIndex uint, barrier bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), sampleBuffer, sampleIndex, barrier)
	}
// Encodes a command that retrieves a sparse texture’s access data for a
// specific region, mipmap level, and slice.
//
// texture: A sparse texture instance.
//
// region: A region within the sparse texture’s `mipLevel`, in sparse tile
// coordinates.
//
// mipLevel: A mipmap level within the sparse texture.
//
// slice: A slice within the sparse texture.
//
// resetCounters: A Boolean value that indicates whether the command resets the counters
// after it completes.
//
// countersBuffer: A destination buffer where the command stores the sparse texture’s access
// counter data.
//
// countersBufferOffset: A starting offset, in bytes, within `countersBuffer` where the command
// writes the first byte of the sparse texture’s access counter data.
//
// # Discussion
// 
// The GPU returns a counter for each sparse tile in the region you specify.
// Each counter is a [uint32_t] in row-major order. Provide space in the
// buffer for each counter you request.
// 
// When the GPU samples a texture and fails to find data in its internal
// caches, the GPU increments the access counter for the sparse tile. The GPU
// then attempts to fetch a new cache line from device memory that contains
// those pixels.
// 
// The counter doesn’t track memory operations to data that’s already in
// the GPU’s caches. You can ignore differences in cache line sizes or pixel
// formats because the GPU driver normalizes the access counts. Each count
// represents the number of pixels the GPU fetches into memory.
//
// [uint32_t]: https://developer.apple.com/documentation/kernel/uint32_t
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/getTextureAccessCounters(_:region:mipLevel:slice:resetCounters:countersBuffer:countersBufferOffset:)
func (o MTLBlitCommandEncoderObject) GetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset(texture MTLTexture, region MTLRegion, mipLevel uint, slice uint, resetCounters bool, countersBuffer MTLBuffer, countersBufferOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("getTextureAccessCounters:region:mipLevel:slice:resetCounters:countersBuffer:countersBufferOffset:"), texture, region, mipLevel, slice, resetCounters, countersBuffer, countersBufferOffset)
	}
// Encodes a command that resets a sparse texture’s access data for a
// specific region, mipmap level, and slice.
//
// texture: A sparse texture instance.
//
// region: A region within the sparse texture’s `mipLevel`, in sparse tile
// coordinates.
//
// mipLevel: A mipmap level within the sparse texture.
//
// slice: A slice within the sparse texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resetTextureAccessCounters(_:region:mipLevel:slice:)
func (o MTLBlitCommandEncoderObject) ResetTextureAccessCountersRegionMipLevelSlice(texture MTLTexture, region MTLRegion, mipLevel uint, slice uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("resetTextureAccessCounters:region:mipLevel:slice:"), texture, region, mipLevel, slice)
	}
// Encodes a command that copies commands from one indirect command buffer
// into another.
//
// source: An indirect command buffer the command copies from.
//
// sourceRange: The range of commands in the source buffer to copy. The source range needs
// to start on a valid execution point.
//
// destination: Another indirect command buffer the command copies to.
//
// destinationIndex: An index in `destination` where the command copies content from `source`
// to. The destination index needs to be a valid execution point with enough
// remaining space in `destination` to accommodate `sourceRange.Count()`
// indices.
//
// # Discussion
// 
// You can copy commands from one indirect command buffer to another, but only
// a compatible one. You can create compatible indirect command buffers by
// passing [MTLIndirectCommandBufferDescriptor] instances with the same
// configuration to the
// [NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions] method of
// [MTLDevice].
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:
func (o MTLBlitCommandEncoderObject) CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source MTLIndirectCommandBuffer, sourceRange foundation.NSRange, destination MTLIndirectCommandBuffer, destinationIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:"), source, sourceRange, destination, destinationIndex)
	}
// Encodes a command that fills a buffer with a constant value for each byte.
//
// buffer: A buffer instance the command assigns each byte in `range` to `value`.
//
// range: A range of bytes within the `buffer` the command assigns `value` to. The
// range’s [count] property needs to be greater than `0`. The range’s
// [count], [lowerBound], and [upperBound] properties need to be a multiple of
// `4` in macOS, but can be any value in iOS and tvOS.
// //
// [count]: https://developer.apple.com/documentation/Swift/Collection/count
// [lowerBound]: https://developer.apple.com/documentation/Swift/Range/lowerBound
// [upperBound]: https://developer.apple.com/documentation/Swift/Range/upperBound
//
// value: The value to write to each byte.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/fillBuffer:range:value:
func (o MTLBlitCommandEncoderObject) FillBufferRangeValue(buffer MTLBuffer, range_ foundation.NSRange, value uint8) {
	
	objc.Send[struct{}](o.ID, objc.Sel("fillBuffer:range:value:"), buffer, range_, value)
	}
// Encodes a command that can improve the performance of a range of commands
// within an indirect command buffer.
//
// indirectCommandBuffer: An indirect command buffer the command optimizes.
//
// range: A range of commands within `indirectCommandBuffer`.
//
// # Discussion
// 
// This command can reduce the time it takes the GPU to run the commands
// within an indirect command buffer by removing its redundancies. For
// example, an indirect command buffer may have empty commands or commands
// that duplicate identical state. Redundancies like these can come from
// multiple compute functions that encode commands in parallel, which can
// sometimes reset commands or configure identical states multiple times.
// 
// You can’t run any commands that start or end at an index within that
// range, or that cross into another optimized range. However, you can reuse
// the range you optimize by resetting it and then encoding new commands to
// it.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/optimizeIndirectCommandBuffer:withRange:
func (o MTLBlitCommandEncoderObject) OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer MTLIndirectCommandBuffer, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("optimizeIndirectCommandBuffer:withRange:"), indirectCommandBuffer, range_)
	}
// Encodes a command that resets a range of commands in an indirect command
// buffer.
//
// buffer: An indirect command buffer the command resets.
//
// range: A range of commands within `buffer`.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resetCommandsInBuffer:withRange:
func (o MTLBlitCommandEncoderObject) ResetCommandsInBufferWithRange(buffer MTLIndirectCommandBuffer, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("resetCommandsInBuffer:withRange:"), buffer, range_)
	}
// Encodes a command that resolves the data from the samples in a sample
// counter buffer and stores the results into a buffer.
//
// sampleBuffer: A counter sample buffer source that contains the sample data.
//
// range: A range that indicates which of the buffer’s samples the command
// resolves.
//
// destinationBuffer: A destination buffer where the command stores the data it resolves.
//
// destinationOffset: A starting offset, in bytes, within `destinationBuffer` where the blit pass
// writes the first byte of the data it resolves.
//
// # Discussion
// 
// For an example of how and when to use this method, see [Converting a
// GPU’s counter data into a readable format].
//
// [Converting a GPU’s counter data into a readable format]: https://developer.apple.com/documentation/Metal/converting-a-gpus-counter-data-into-a-readable-format
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitCommandEncoder/resolveCounters:inRange:destinationBuffer:destinationOffset:
func (o MTLBlitCommandEncoderObject) ResolveCountersInRangeDestinationBufferDestinationOffset(sampleBuffer MTLCounterSampleBuffer, range_ foundation.NSRange, destinationBuffer MTLBuffer, destinationOffset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("resolveCounters:inRange:destinationBuffer:destinationOffset:"), sampleBuffer, range_, destinationBuffer, destinationOffset)
	}
// Declares that all command generation from the encoder is completed.
//
// # Discussion
// 
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLBlitCommandEncoderObject) EndEncoding() {
	
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
func (o MTLBlitCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
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
func (o MTLBlitCommandEncoderObject) PushDebugGroup(string_ string) {
	
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
func (o MTLBlitCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}
// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLBlitCommandEncoderObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLBlitCommandEncoderObject) Label() string {
	
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
func (o MTLBlitCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
	}

func (o MTLBlitCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

