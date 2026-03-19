// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
)

// A resource that holds formatted image data.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture
type MTLTexture interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// Copies pixel data into a section of a texture slice.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/replace(region:mipmapLevel:slice:withBytes:bytesPerRow:bytesPerImage:)
	ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region MTLRegion, level uint, slice uint, pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint)

	// Copies a block of pixels into a section of texture slice 0.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/replace(region:mipmapLevel:withBytes:bytesPerRow:)
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region MTLRegion, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint)

	// Copies pixel data from the texture to a buffer in system memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/getBytes(_:bytesPerRow:bytesPerImage:from:mipmapLevel:slice:)
	GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region MTLRegion, level uint, slice uint)

	// Copies pixel data from the first slice of the texture to a buffer in system memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/getBytes(_:bytesPerRow:from:mipmapLevel:)
	GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes unsafe.Pointer, bytesPerRow uint, region MTLRegion, level uint)

	// Creates a new view of the texture, reinterpreting its data using a different pixel format.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeTextureView(pixelFormat:)
	NewTextureViewWithPixelFormat(pixelFormat MTLPixelFormat) MTLTexture

	// The dimension and arrangement of the texture image data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/textureType
	TextureType() MTLTextureType

	// The format of pixels in the texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/pixelFormat
	PixelFormat() MTLPixelFormat

	// The width of the texture image for the base level mipmap, in pixels.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/width
	Width() uint

	// The height of the texture image for the base level mipmap, in pixels.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/height
	Height() uint

	// The depth of the texture image for the base level mipmap, in pixels.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/depth
	Depth() uint

	// The number of mipmap levels in the texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/mipmapLevelCount
	MipmapLevelCount() uint

	// The number of slices in the texture array.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/arrayLength
	ArrayLength() uint

	// The number of samples in each pixel.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/sampleCount
	SampleCount() uint

	// A Boolean value that indicates whether the texture can only be used as a render target.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/isFramebufferOnly
	IsFramebufferOnly() bool

	// Options that determine how you can use the texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/usage
	Usage() MTLTextureUsage

	// A Boolean value indicating whether the GPU is allowed to adjust the contents of the texture to improve GPU performance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/allowGPUOptimizedContents
	AllowGPUOptimizedContents() bool

	// A Boolean indicating whether this texture can be shared with other processes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/isShareable
	IsShareable() bool

	// The pattern that the GPU applies to pixels when you read or sample pixels from the texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/swizzle
	Swizzle() MTLTextureSwizzleChannels

	// A reference to the underlying surface instance for the texture, if applicable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/iosurface
	Iosurface() iosurface.IOSurfaceRef

	// The number of a plane within the underlying surface instance for the texture, if applicable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/iosurfacePlane
	IosurfacePlane() uint

	// The parent texture used to create this texture, if any.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/parent
	ParentTexture() MTLTexture

	// The base level of the parent texture used to create this texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/parentRelativeLevel
	ParentRelativeLevel() uint

	// The base slice of the parent texture used to create this texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/parentRelativeSlice
	ParentRelativeSlice() uint

	// The source buffer used to create this texture, if any.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/buffer
	Buffer() MTLBuffer

	// The offset in the source buffer where the texture’s data comes from.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/bufferOffset
	BufferOffset() uint

	// The number of bytes in each row of the texture’s source buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/bufferBytesPerRow
	BufferBytesPerRow() uint

	// The resource that owns the storage for this texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/rootResource
	RootResource() MTLResource

	// Creates a new texture handle from a shareable texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeSharedTextureHandle()
	NewSharedTextureHandle() IMTLSharedTextureHandle

	// Creates a remote texture view for another GPU in the same peer group.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeRemoteTextureView(_:)
	NewRemoteTextureViewForDevice(device MTLDevice) MTLTexture

	// The texture on another GPU that the texture was created from, if any.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/remoteStorageTexture
	RemoteStorageTexture() MTLTexture

	// A Boolean value that indicates whether this is a sparse texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/isSparse
	IsSparse() bool

	// The index of the first mipmap in the tail.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/firstMipmapInTail
	FirstMipmapInTail() uint

	// The size of the sparse texture tail, in bytes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/tailSizeInBytes
	TailSizeInBytes() uint

	// CompressionType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/compressionType
	CompressionType() MTLTextureCompressionType

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/gpuResourceID
	GpuResourceID() MTLResourceID

	// SparseTextureTier protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/sparseTextureTier
	SparseTextureTier() MTLTextureSparseTier

	// NewTextureViewWithDescriptor protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureView(with:)
	NewTextureViewWithDescriptor(descriptor IMTLTextureViewDescriptor) MTLTexture

	// Creates a new view of the texture, reinterpreting a subset of its data using a different type and pixel format.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureViewWithPixelFormat:textureType:levels:slices:
	NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat MTLPixelFormat, textureType MTLTextureType, levelRange foundation.NSRange, sliceRange foundation.NSRange) MTLTexture

	// Creates a new view of the texture, reinterpreting a subset of its data using a different type, pixel format, and swizzle pattern.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:
	NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat MTLPixelFormat, textureType MTLTextureType, levelRange foundation.NSRange, sliceRange foundation.NSRange, swizzle MTLTextureSwizzleChannels) MTLTexture
}

// MTLTextureObject wraps an existing Objective-C object that conforms to the MTLTexture protocol.
type MTLTextureObject struct {
	objectivec.Object
}
func (o MTLTextureObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTextureObjectFromID constructs a [MTLTextureObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTextureObjectFromID(id objc.ID) MTLTextureObject {
	return MTLTextureObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Copies pixel data into a section of a texture slice.
//
// region: The location of a block of pixels in the texture slice. The region needs to
// be within the dimensions of the slice.
//
// level: A zero-indexed value that specifies which mipmap level is the destination.
// If the texture doesn’t have mipmaps, use `0`.
//
// slice: A zero-indexed value that specifies which texture slice is the destination:
// 
// - For a cube texture, `slice` is a value between `0` and `5`, inclusive,
// that defines which cube face is the destination. - For a texture array,
// `slice` is the element index. - For a cube texture array, slice defines
// both the cube face and an array index. To determine the correct slice for a
// cube texture array, treat it as having a stride of `6`: `slice = cubeFace +
// arrayIndex * 6.` - For all other texture types, use `0`.
//
// pixelBytes: A pointer to the bytes in memory to copy.
//
// bytesPerRow: The stride, in bytes, of one row in the source data. For [TextureType1D]
// and [TextureType1DArray], use `0`. For raw and packed pixel types, the
// stride is the number of pixels in one row. For compressed pixel formats,
// the stride is the number of bytes from the beginning of one row of blocks
// to the beginning of the next. When source data consists of only a single
// row, use `0`.
// 
// Your data type determines how you should compute `bytesPerRow`:
// 
// - For raw or packed pixel data, use a value greater than or equal to the
// size of data in one row, and less than `32767 * pixel size`. - For
// compressed pixel data, use a multiple of the compression block size. When
// working with PowerVR Texture Compression (PVRTC), use `0.`
// 
// Nonzero values smaller than the texture width or not a multiple of the
// pixel size cause an error.
//
// bytesPerImage: The stride, in bytes, between images in the source data. Supply a nonzero
// value only when you copy data to an [TextureType3D] type texture. Your data
// type determines how you should compute `bytesPerImage`:
// 
// - For data that consists only of a single image, use `0`. - For ordinary or
// packed pixel formats, use a multiple of pixel size. - For compressed pixel
// data, use a multiple of the compression block size. When working with
// PVRTC, use `0`.
// 
// When copying data to a type of texture other than [TextureType3D], use `0`.
//
// # Discussion
// 
// This method runs on the CPU and immediately copies the pixel data into the
// texture. It doesn’t synchronize against any GPU memory operations to the
// texture. Ensure all operations that write or render to the texture complete
// before reading the texture’s contents using one of the following methods:
// 
// - Synchronize on the GPU with a [SynchronizeResource] or
// [SynchronizeTextureSliceLevel] command in an [MTLBlitCommandEncoder]. -
// Synchronize on the CPU with a callback passed to the [AddCompletedHandler]
// method to handle completion asynchronously, or the [WaitUntilCompleted]
// method to block thread execution until the GPU work completes.
// 
// If the texture image has a compressed pixel format, only write to
// block-aligned regions. If the size of a dimension of region isn’t a
// multiple of the block size, include both the edge block and additional
// space up to the texture dimensions in `bytesPerRow`.
// 
// To copy your data to a private texture, copy your data to a temporary
// texture with non-private storage, and then use an [MTLBlitCommandEncoder]
// to copy the data to the private texture for GPU use.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/replace(region:mipmapLevel:slice:withBytes:bytesPerRow:bytesPerImage:)
func (o MTLTextureObject) ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region MTLRegion, level uint, slice uint, pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("replaceRegion:mipmapLevel:slice:withBytes:bytesPerRow:bytesPerImage:"), region, level, slice, pixelBytes, bytesPerRow, bytesPerImage)
	}
// Copies a block of pixels into a section of texture slice 0.
//
// region: The location of a block of pixels in the texture slice. The region needs to
// be within the dimensions of the slice.
//
// level: A zero-indexed value that specifies which mipmap level is the destination.
// If the texture doesn’t have mipmaps, use `0`.
//
// pixelBytes: A pointer to the bytes in memory to copy.
//
// bytesPerRow: The stride, in bytes, of one row in the source data. For [TextureType1D]
// and [TextureType1DArray], use `0`. For raw and packed pixel types, the
// stride is the number of pixels in one row. For compressed pixel formats,
// the stride is the number of bytes from the beginning of one row of blocks
// to the beginning of the next. When source data consists of only a single
// row, use `0`.
// 
// Your data type determines how you should compute `bytesPerRow`:
// 
// - For raw or packed pixel data, use a value greater than or equal to the
// size of data in one row, and less than [max] `* pixel size`. - For
// compressed pixel data, use a multiple of the compression block size. When
// working with PowerVR Texture Compression (PVRTC), use `0.`
// 
// Nonzero values smaller than the texture width or not a multiple of the
// pixel size cause an error.
// //
// [max]: https://developer.apple.com/documentation/Swift/Int32/max
//
// # Discussion
// 
// This method runs on the CPU and immediately copies the pixel data into the
// texture. It doesn’t synchronize against any GPU memory operations to the
// texture. Ensure all operations that write or render to the texture complete
// before reading the texture’s contents using one of the following methods:
// 
// - Synchronize on the GPU with a [SynchronizeResource] or
// [SynchronizeTextureSliceLevel] command in an [MTLBlitCommandEncoder]. -
// Synchronize on the CPU with a callback passed to the [AddCompletedHandler]
// method to handle completion asynchronously, or the [WaitUntilCompleted]
// method to block thread execution until the GPU work completes.
// 
// If the texture image has a compressed pixel format, only write to
// block-aligned regions. If the size of a dimension of region isn’t a
// multiple of the block size, then include both the edge block and space up
// to the texture dimensions in `bytesPerRow`.
// 
// To copy your data to a private texture, copy your data to a temporary
// texture with non-private storage, and then use an [MTLBlitCommandEncoder]
// to copy the data to the private texture for GPU use.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/replace(region:mipmapLevel:withBytes:bytesPerRow:)
func (o MTLTextureObject) ReplaceRegionMipmapLevelWithBytesBytesPerRow(region MTLRegion, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"), region, level, pixelBytes, bytesPerRow)
	}
// Copies pixel data from the texture to a buffer in system memory.
//
// pixelBytes: A pointer to a destination buffer in system memory.
//
// bytesPerRow: The number of bytes () between two adjacent rows of pixel data in the
// destination buffer. For [TextureType1D] and [TextureType1DArray], use `0`.
// For raw and packed pixel types, the stride is the number of pixels in one
// row. For compressed pixel formats, the stride is the number of bytes from
// the beginning of one row of blocks to the beginning of the next.
// 
// Your data type determines how you should compute `bytesPerRow`:
// 
// - For raw or packed pixel data, use a multiple of the pixel size less than
// [max] `* pixel size`. - For compressed pixel data, use a multiple of the
// compression block size. When working with PowerVR Texture Compression
// (PVRTC), use `0.`
// 
// Nonzero values smaller than the texture width or any values not a multiple
// of the pixel or block size cause an error.
// //
// [max]: https://developer.apple.com/documentation/Swift/Int32/max
//
// bytesPerImage: The stride between adjacent images in the destination buffer.
//
// region: The location of a block of pixels in the texture slice. For textures
// compressed as PVRTC, use the entire texture for the region.
//
// level: A zero-indexed value that selects the texture’s mipmap level as the
// method’s data source. Use `0` for textures that don’t have mipmaps.
//
// slice: A zero-indexed value specifying the destination texture slice:
// 
// - For a cube texture, `slice` is a value between `0` and `5`, inclusive,
// that defines which cube face is the source. - For a texture array, `slice`
// is the element index. - For a cube texture array, slice defines both the
// cube face and an array index. To determine the correct slice for a cube
// texture array, treat it as having a stride of `6`: `slice = cubeFace +
// arrayIndex * 6.` - For all other texture types, use `0`.
//
// # Discussion
// 
// This method runs on the CPU and immediately copies the pixel data from the
// texture to system memory, but it doesn’t synchronize with any GPU texture
// memory operations. Ensure all operations that write or render to the
// texture complete before reading the texture’s contents using one of the
// following methods:
// 
// - Synchronize on the GPU with a [SynchronizeResource] or
// [SynchronizeTextureSliceLevel] command in an [MTLBlitCommandEncoder]. -
// Synchronize on the CPU with a callback passed to the [AddCompletedHandler]
// method to handle completion asynchronously, or the [WaitUntilCompleted]
// method to block thread execution until the GPU work completes.
// 
// For multisample textures, the method consecutively positions each sample
// within a pixel in memory and treats the pixels as part of one row.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/getBytes(_:bytesPerRow:bytesPerImage:from:mipmapLevel:slice:)
func (o MTLTextureObject) GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region MTLRegion, level uint, slice uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("getBytes:bytesPerRow:bytesPerImage:fromRegion:mipmapLevel:slice:"), pixelBytes, bytesPerRow, bytesPerImage, region, level, slice)
	}
// Copies pixel data from the first slice of the texture to a buffer in system
// memory.
//
// pixelBytes: A pointer to a destination buffer in system memory.
//
// bytesPerRow: The number of bytes () between two adjacent rows of pixel data in the
// destination buffer. For [TextureType1D] and [TextureType1DArray], use `0`.
// For raw and packed pixel types, the stride is the number of pixels in one
// row. For compressed pixel formats, the stride is the number of bytes from
// the beginning of one row of blocks to the beginning of the next.
// 
// Your data type determines how you should compute `bytesPerRow`:
// 
// - For raw or packed pixel data, use a multiple of the pixel size less than
// [max] `* pixel size`. - For compressed pixel data, use a multiple of the
// compression block size. When working with PowerVR Texture Compression
// (PVRTC), use `0.`
// 
// Nonzero values smaller than the texture width or any values not a multiple
// of the pixel or block size cause an error.
// //
// [max]: https://developer.apple.com/documentation/Swift/Int32/max
//
// region: The location of a block of pixels in the texture slice. For textures
// compressed as PVRTC, use the entire texture for the region.
//
// level: A zero-indexed value that selects the texture’s mipmap level as the
// method’s data source. Use `0` for textures that don’t have mipmaps.
//
// # Discussion
// 
// This method runs on the CPU and immediately copies the pixel data from the
// texture to system memory, but it doesn’t synchronize with any GPU texture
// memory operations. Ensure all operations that write or render to the
// texture complete before reading the texture’s contents using one of the
// following methods:
// 
// - Synchronize on the GPU with a [SynchronizeResource] or
// [SynchronizeTextureSliceLevel] command in an [MTLBlitCommandEncoder]. -
// Synchronize on the CPU with a callback passed to the [AddCompletedHandler]
// method to handle completion asynchronously, or the [WaitUntilCompleted]
// method to block thread execution until the GPU work completes.
// 
// For multisample textures, the method consecutively positions each sample
// within a pixel in memory and treats the pixels as part of one row.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/getBytes(_:bytesPerRow:from:mipmapLevel:)
func (o MTLTextureObject) GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes unsafe.Pointer, bytesPerRow uint, region MTLRegion, level uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("getBytes:bytesPerRow:fromRegion:mipmapLevel:"), pixelBytes, bytesPerRow, region, level)
	}
// Creates a new view of the texture, reinterpreting its data using a
// different pixel format.
//
// pixelFormat: A new pixel format, which needs to be compatible with the original pixel
// format.
//
// # Return Value
// 
// A new texture object that shares the same storage allocation of the
// texture.
//
// # Discussion
// 
// When you create a texture normally, Metal allocates memory for the
// textureʼs pixel data. These storage allocations can be quite large. You
// can reduce memory use and avoid copying texture data by using a —a
// texture object that shares another textureʼs storage allocation,
// reinterpreting the pixel data in some other format.
// 
// Not all pixel formats are compatible with one another. Reinterpretation of
// image data between pixel formats is supported within the following groups:
// 
// - All 8-, 16-, 32-, 64-, and 128-bit color formats are compatible with
// other formats with the same bit length. - sRGB and non-sRGB forms of the
// same compressed format (for example, [PixelFormatBC1_RGBA] and
// [PixelFormatBC1_RGBA_sRGB]) - Combined depth-stencil texture formats and
// the related format used to access the stencil from a shader (for example,
// [PixelFormatDepth24Unorm_Stencil8] and [PixelFormatX24_Stencil8])
// 
// This method doesn’t change the original texture image data in any way,
// but it may drastically change how the data is interpreted. For example,
// given a texture with the [PixelFormatRG16Uint] pixel format that contains
// image data for Red `0xFFFE` and Green `0x0001`, this method would
// reinterpret that data in an [PixelFormatR32Uint] format as Red
// `0x0001FFFE`.
// 
// Some format reinterpretations are supported but may not be useful. For
// example, this method considers the 32-bit packed color formats
// [PixelFormatBGR10A2Unorm] and [DataTypeRG11B10Float] to be compatible, but
// it’s unlikely that the same data can be interpreted by both formats in a
// meaningful way.
// 
// Some format reinterpretations require you to create the source texture with
// a special usage flag. Set that flag only when necessary, as it can affect
// performance. For more details, see [TextureUsagePixelFormatView].
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeTextureView(pixelFormat:)
func (o MTLTextureObject) NewTextureViewWithPixelFormat(pixelFormat MTLPixelFormat) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureViewWithPixelFormat:"), pixelFormat)
	return MTLTextureObjectFromID(rv)
	}
// The dimension and arrangement of the texture image data.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/textureType
func (o MTLTextureObject) TextureType() MTLTextureType {
	
	rv := objc.Send[MTLTextureType](o.ID, objc.Sel("textureType"))
	return rv
	}
// The format of pixels in the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/pixelFormat
func (o MTLTextureObject) PixelFormat() MTLPixelFormat {
	
	rv := objc.Send[MTLPixelFormat](o.ID, objc.Sel("pixelFormat"))
	return rv
	}
// The width of the texture image for the base level mipmap, in pixels.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/width
func (o MTLTextureObject) Width() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("width"))
	return rv
	}
// The height of the texture image for the base level mipmap, in pixels.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/height
func (o MTLTextureObject) Height() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("height"))
	return rv
	}
// The depth of the texture image for the base level mipmap, in pixels.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/depth
func (o MTLTextureObject) Depth() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("depth"))
	return rv
	}
// The number of mipmap levels in the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/mipmapLevelCount
func (o MTLTextureObject) MipmapLevelCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("mipmapLevelCount"))
	return rv
	}
// The number of slices in the texture array.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/arrayLength
func (o MTLTextureObject) ArrayLength() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("arrayLength"))
	return rv
	}
// The number of samples in each pixel.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/sampleCount
func (o MTLTextureObject) SampleCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("sampleCount"))
	return rv
	}
// A Boolean value that indicates whether the texture can only be used as a
// render target.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/isFramebufferOnly
func (o MTLTextureObject) IsFramebufferOnly() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isFramebufferOnly"))
	return rv
	}
// Options that determine how you can use the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/usage
func (o MTLTextureObject) Usage() MTLTextureUsage {
	
	rv := objc.Send[MTLTextureUsage](o.ID, objc.Sel("usage"))
	return rv
	}
// A Boolean value indicating whether the GPU is allowed to adjust the
// contents of the texture to improve GPU performance.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/allowGPUOptimizedContents
func (o MTLTextureObject) AllowGPUOptimizedContents() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
	}
// A Boolean indicating whether this texture can be shared with other
// processes.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/isShareable
func (o MTLTextureObject) IsShareable() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isShareable"))
	return rv
	}
// The pattern that the GPU applies to pixels when you read or sample pixels
// from the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/swizzle
func (o MTLTextureObject) Swizzle() MTLTextureSwizzleChannels {
	
	rv := objc.Send[MTLTextureSwizzleChannels](o.ID, objc.Sel("swizzle"))
	return rv
	}
// A reference to the underlying surface instance for the texture, if
// applicable.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/iosurface
func (o MTLTextureObject) Iosurface() iosurface.IOSurfaceRef {
	
	rv := objc.Send[iosurface.IOSurfaceRef](o.ID, objc.Sel("iosurface"))
	return rv
	}
// The number of a plane within the underlying surface instance for the
// texture, if applicable.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/iosurfacePlane
func (o MTLTextureObject) IosurfacePlane() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("iosurfacePlane"))
	return rv
	}
// The parent texture used to create this texture, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/parent
func (o MTLTextureObject) ParentTexture() MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parentTexture"))
	return MTLTextureObjectFromID(rv)
	}
// The base level of the parent texture used to create this texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/parentRelativeLevel
func (o MTLTextureObject) ParentRelativeLevel() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("parentRelativeLevel"))
	return rv
	}
// The base slice of the parent texture used to create this texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/parentRelativeSlice
func (o MTLTextureObject) ParentRelativeSlice() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("parentRelativeSlice"))
	return rv
	}
// The source buffer used to create this texture, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/buffer
func (o MTLTextureObject) Buffer() MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buffer"))
	return MTLBufferObjectFromID(rv)
	}
// The offset in the source buffer where the texture’s data comes from.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/bufferOffset
func (o MTLTextureObject) BufferOffset() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("bufferOffset"))
	return rv
	}
// The number of bytes in each row of the texture’s source buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/bufferBytesPerRow
func (o MTLTextureObject) BufferBytesPerRow() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("bufferBytesPerRow"))
	return rv
	}
// The resource that owns the storage for this texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/rootResource
func (o MTLTextureObject) RootResource() MTLResource {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("rootResource"))
	return MTLResourceObjectFromID(rv)
	}
// Creates a new texture handle from a shareable texture.
//
// # Discussion
// 
// If the texture is not shareable, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeSharedTextureHandle()
func (o MTLTextureObject) NewSharedTextureHandle() IMTLSharedTextureHandle {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedTextureHandle"))
	return MTLSharedTextureHandleFromID(rv)
	}
// Creates a remote texture view for another GPU in the same peer group.
//
// # Discussion
// 
// The device instance that created this texture and the device instance
// passed into this method need to have the same nonzero peer group identifier
// ([peerGroupID]). This texture needs to either use the private storage mode
// ([StorageModePrivate]) or be backed by an [IOSurface].
// 
// A remote view doesn’t allocate any storage for the new texture; it
// references the memory allocated for the original texture. You can use
// remote views only as a source for copy commands encoded by an
// [MTLBlitCommandEncoder]. For more information, see [Transferring data
// between connected GPUs].
//
// [IOSurface]: https://developer.apple.com/documentation/IOSurface/IOSurface
// [Transferring data between connected GPUs]: https://developer.apple.com/documentation/Metal/transferring-data-between-connected-gpus
// [peerGroupID]: https://developer.apple.com/documentation/Metal/MTLDevice/peerGroupID
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/makeRemoteTextureView(_:)
func (o MTLTextureObject) NewRemoteTextureViewForDevice(device MTLDevice) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRemoteTextureViewForDevice:"), device)
	return MTLTextureObjectFromID(rv)
	}
// The texture on another GPU that the texture was created from, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/remoteStorageTexture
func (o MTLTextureObject) RemoteStorageTexture() MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("remoteStorageTexture"))
	return MTLTextureObjectFromID(rv)
	}
// A Boolean value that indicates whether this is a sparse texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/isSparse
func (o MTLTextureObject) IsSparse() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isSparse"))
	return rv
	}
// The index of the first mipmap in the tail.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/firstMipmapInTail
func (o MTLTextureObject) FirstMipmapInTail() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("firstMipmapInTail"))
	return rv
	}
// The size of the sparse texture tail, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/tailSizeInBytes
func (o MTLTextureObject) TailSizeInBytes() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("tailSizeInBytes"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTexture/compressionType
func (o MTLTextureObject) CompressionType() MTLTextureCompressionType {
	
	rv := objc.Send[MTLTextureCompressionType](o.ID, objc.Sel("compressionType"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTexture/gpuResourceID
func (o MTLTextureObject) GpuResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLTexture/sparseTextureTier
func (o MTLTextureObject) SparseTextureTier() MTLTextureSparseTier {
	
	rv := objc.Send[MTLTextureSparseTier](o.ID, objc.Sel("sparseTextureTier"))
	return rv
	}
//
// # Discussion
// 
// Create a new texture which shares the same storage as the source texture,
// but with different (but compatible) properties specified by the descriptor
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureView(with:)
func (o MTLTextureObject) NewTextureViewWithDescriptor(descriptor IMTLTextureViewDescriptor) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureViewWithDescriptor:"), descriptor)
	return MTLTextureObjectFromID(rv)
	}
// Creates a new view of the texture, reinterpreting a subset of its data
// using a different type and pixel format.
//
// pixelFormat: A new pixel format, which needs to be compatible with the original pixel
// format.
//
// textureType: A new texture type, which can be cast according to the original texture
// type as listed in the table below.
//
// levelRange: A new base level range that restricts which mipmap levels are visible in
// the new texture.
//
// sliceRange: A new base slice range that restricts which array slices are visible in the
// new texture.
//
// # Return Value
// 
// A new texture object that shares the same storage allocation of the calling
// texture object.
//
// # Discussion
// 
// The texture type can be cast between the targets listed in the following
// table.
// 
// [Table data omitted]
// 
// The `length` value of the `sliceRange` parameter needs to be `6` if the new
// texture type value is [TextureTypeCube], or a multiple of `6` if the new
// texture type value is [TextureTypeCubeArray].
// 
// For more information on pixel format restrictions, see
// [NewTextureViewWithPixelFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureViewWithPixelFormat:textureType:levels:slices:
func (o MTLTextureObject) NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat MTLPixelFormat, textureType MTLTextureType, levelRange foundation.NSRange, sliceRange foundation.NSRange) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:"), pixelFormat, textureType, levelRange, sliceRange)
	return MTLTextureObjectFromID(rv)
	}
// Creates a new view of the texture, reinterpreting a subset of its data
// using a different type, pixel format, and swizzle pattern.
//
// pixelFormat: A new pixel format, which needs to be compatible with the original pixel
// format.
//
// textureType: A new texture type.
//
// levelRange: A new base level range that restricts which mipmap levels are visible in
// the new texture.
//
// sliceRange: A new base slice range that restricts which array slices are visible in the
// new texture.
//
// swizzle: The swizzle pattern the GPU uses to reorder the data when sampling or
// reading the texture.
//
// # Return Value
// 
// A new texture view.
//
// # Discussion
// 
// For more information on texture views, see
// [NewTextureViewWithPixelFormatTextureTypeLevelsSlices].
// 
// The swizzle pattern of the view is combined with that of the parent texture
// to generate the final swizzle pattern. For example: An `[R,G,A,B]` swizzle
// of a texture with a `[R,1,1,G]` swizzle pattern is `[R,1,G,1]`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTexture/newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:
func (o MTLTextureObject) NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat MTLPixelFormat, textureType MTLTextureType, levelRange foundation.NSRange, sliceRange foundation.NSRange, swizzle MTLTextureSwizzleChannels) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:"), pixelFormat, textureType, levelRange, sliceRange, swizzle)
	return MTLTextureObjectFromID(rv)
	}
// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLTextureObject) AllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}
// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device
func (o MTLTextureObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
func (o MTLTextureObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode
func (o MTLTextureObject) CpuCacheMode() MTLCPUCacheMode {
	
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}
// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode
func (o MTLTextureObject) StorageMode() MTLStorageMode {
	
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}
// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode
func (o MTLTextureObject) HazardTrackingMode() MTLHazardTrackingMode {
	
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}
// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions
func (o MTLTextureObject) ResourceOptions() MTLResourceOptions {
	
	rv := objc.Send[MTLResourceOptions](o.ID, objc.Sel("resourceOptions"))
	return rv
	}
// Specifies or queries the resource’s purgeable state.
//
// state: The desired purgeable state of a resource.
//
// # Return Value
// 
// The prior purgeable state of the resource.
//
// # Discussion
// 
// If `state` is [PurgeableStateKeepCurrent], the method returns the current
// purgeable state without changing it.
// 
// If `state` is [PurgeableStateNonVolatile], the resource is marked to inform
// the caller that the data should not be discarded.
// 
// If `state` is [PurgeableStateEmpty], the resource is marked as data that
// can be discarded, because the caller no longer needs the contents of the
// resource.
// 
// If `state` is [PurgeableStateVolatile], the resource is marked as data that
// can be discarded, even if the caller may need the resource. [MTLResource]
// objects can be made purgeable, even if the caller may need the resource,
// where the implementation can reclaim the underlying storage at any time
// without informing the app. Purgeable resources may enable an app to keep
// larger caches of idle memory that may be useful again in the future without
// the risk of preventing the allocation of more important memory.
// 
// When you use purgeable memory, make sure the block of memory is locked
// before you access it. This locking mechanism ensures that auto-removal
// policies don’t discard the data while you are accessing it. Similarly,
// the locking mechanism ensures that the virtual memory system has not
// already discarded the data.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setPurgeableState(_:)
func (o MTLTextureObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}
// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset
func (o MTLTextureObject) HeapOffset() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
	}
// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap
func (o MTLTextureObject) Heap() MTLHeap {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("heap"))
	return MTLHeapObjectFromID(rv)
	}
// Allows future heap resource allocations to alias against the resource’s
// memory, reusing it.
//
// # Discussion
// 
// Resource instances marked as aliased have backing memory available for use
// in new allocations to the heap. One common use case is to make a single
// large resource aliasable for reuse of memory by smaller and more frequent
// resource allocations. For situations where you need fine-grained control
// over your memory management, you might want to use a heap with the
// allocation type [HeapTypePlacement] and manage memory yourself instead.
// 
// Aliased resources can’t be un-aliased or moved. If you use an aliased
// resource instance to read or write data, it results in undefined behavior.
// 
// When working with resources possibly backed by aliased memory, you should
// take great care that the system doesn’t access resources from multiple
// aliases concurrently. Use an [MTLEvent] or [MTLFence] instance to protect
// access to resources that you’ve either already aliased or intend to
// alias.
// 
// The general process to reuse memory from aliased resources is:
// 
// - Allocate an [MTLHeap] instance to hold your task’s resources, using the
// [NewHeapWithDescriptor] method. Your heap should be big enough to store the
// maximum amount of concurrently loaded data you expect. - Allocate your
// resource(s) using a heap method that returns an [MTLResource] instance. -
// Perform your stage on the GPU, and when it completes, mark the resource
// allocation(s) as aliasable by calling this method. - For each successive
// stage of your overall pass, repeat steps 2 and 3. Ensure that the prior
// stage fully completes before making any new resources on an aliasable heap
// through an event or fence.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/makeAliasable()
func (o MTLTextureObject) MakeAliasable() {
	
	objc.Send[struct{}](o.ID, objc.Sel("makeAliasable"))
	}
// A Boolean value that indicates whether future heap resource allocations may
// alias against the resource’s memory.
//
// # Return Value
// 
// The default value is [false]. The value is [true] only if the
// [MakeAliasable] method was previously called on this resource.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/isAliasable()
func (o MTLTextureObject) IsAliasable() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:
func (o MTLTextureObject) SetOwnerWithIdentity(task_id_token objectivec.IObject) int32 {
	
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
	}

func (o MTLTextureObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

