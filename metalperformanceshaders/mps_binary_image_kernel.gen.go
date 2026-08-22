// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSBinaryImageKernel] class.
var (
	_MPSBinaryImageKernelClass     MPSBinaryImageKernelClass
	_MPSBinaryImageKernelClassOnce sync.Once
)

func getMPSBinaryImageKernelClass() MPSBinaryImageKernelClass {
	_MPSBinaryImageKernelClassOnce.Do(func() {
		_MPSBinaryImageKernelClass = MPSBinaryImageKernelClass{class: objc.GetClass("MPSBinaryImageKernel")}
	})
	return _MPSBinaryImageKernelClass
}

// GetMPSBinaryImageKernelClass returns the class object for MPSBinaryImageKernel.
func GetMPSBinaryImageKernelClass() MPSBinaryImageKernelClass {
	return getMPSBinaryImageKernelClass()
}

type MPSBinaryImageKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSBinaryImageKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSBinaryImageKernelClass) Alloc() MPSBinaryImageKernel {
	rv := objc.Send[MPSBinaryImageKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that consumes two textures and produces one texture.
//
// # Overview
//
// [MPSBinaryImageKernel] defines shared behavior for most image processing
// kernels (filters) such as edging modes, clipping, and tiling support for
// image operations that consume two source textures. It is not meant to be
// used directly, but provides API abstraction and in some cases may allow
// some level of polymorphic manipulation of image kernel objects.
//
// # Methods
//
//   - [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [MPSBinaryImageKernel.EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture]: Encodes a kernel into a command buffer, out-of-place.
//   - [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage]
//   - [MPSBinaryImageKernel.PrimarySourceRegionForDestinationSize]: Determines the region of the primary source texture that will be read for an encode operation.
//   - [MPSBinaryImageKernel.SecondarySourceRegionForDestinationSize]: Determines the region of the secondary source texture that will be read for an encode operation.
//
// # Properties
//
//   - [MPSBinaryImageKernel.PrimaryOffset]: The position of the destination clip rectangle origin relative to the primary source buffer.
//   - [MPSBinaryImageKernel.SetPrimaryOffset]
//   - [MPSBinaryImageKernel.SecondaryOffset]: The position of the destination clip rectangle origin relative to the secondary source buffer.
//   - [MPSBinaryImageKernel.SetSecondaryOffset]
//   - [MPSBinaryImageKernel.PrimaryEdgeMode]: The edge mode to use when texture reads stray off the edge of the primary source image.
//   - [MPSBinaryImageKernel.SetPrimaryEdgeMode]
//   - [MPSBinaryImageKernel.SecondaryEdgeMode]: The edge mode to use when texture reads stray off the edge of the secondary source image.
//   - [MPSBinaryImageKernel.SetSecondaryEdgeMode]
//   - [MPSBinaryImageKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//   - [MPSBinaryImageKernel.SetClipRect]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel
type MPSBinaryImageKernel struct {
	MPSKernel
}

// MPSBinaryImageKernelFromID constructs a [MPSBinaryImageKernel] from an objc.ID.
//
// A kernel that consumes two textures and produces one texture.
func MPSBinaryImageKernelFromID(id objc.ID) MPSBinaryImageKernel {
	return MPSBinaryImageKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSBinaryImageKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSBinaryImageKernel] class.
//
// # Methods
//
//   - [IMPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [IMPSBinaryImageKernel.EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [IMPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture]: Encodes a kernel into a command buffer, out-of-place.
//   - [IMPSBinaryImageKernel.EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage]
//   - [IMPSBinaryImageKernel.PrimarySourceRegionForDestinationSize]: Determines the region of the primary source texture that will be read for an encode operation.
//   - [IMPSBinaryImageKernel.SecondarySourceRegionForDestinationSize]: Determines the region of the secondary source texture that will be read for an encode operation.
//
// # Properties
//
//   - [IMPSBinaryImageKernel.PrimaryOffset]: The position of the destination clip rectangle origin relative to the primary source buffer.
//   - [IMPSBinaryImageKernel.SetPrimaryOffset]
//   - [IMPSBinaryImageKernel.SecondaryOffset]: The position of the destination clip rectangle origin relative to the secondary source buffer.
//   - [IMPSBinaryImageKernel.SetSecondaryOffset]
//   - [IMPSBinaryImageKernel.PrimaryEdgeMode]: The edge mode to use when texture reads stray off the edge of the primary source image.
//   - [IMPSBinaryImageKernel.SetPrimaryEdgeMode]
//   - [IMPSBinaryImageKernel.SecondaryEdgeMode]: The edge mode to use when texture reads stray off the edge of the secondary source image.
//   - [IMPSBinaryImageKernel.SetSecondaryEdgeMode]
//   - [IMPSBinaryImageKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//   - [IMPSBinaryImageKernel.SetClipRect]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel
type IMPSBinaryImageKernel interface {
	IMPSKernel

	// Topic: Methods

	// This method attempts to apply a kernel in place on a texture.
	EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, primaryTexture metal.MTLTexture, inPlaceSecondaryTexture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool
	// This method attempts to apply a kernel in place on a texture.
	EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, inPlacePrimaryTexture metal.MTLTexture, secondaryTexture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool
	// Encodes a kernel into a command buffer, out-of-place.
	EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, primaryTexture metal.MTLTexture, secondaryTexture metal.MTLTexture, destinationTexture metal.MTLTexture)
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationImage IMPSImage)
	// Determines the region of the primary source texture that will be read for an encode operation.
	PrimarySourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion
	// Determines the region of the secondary source texture that will be read for an encode operation.
	SecondarySourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion

	// Topic: Properties

	// The position of the destination clip rectangle origin relative to the primary source buffer.
	PrimaryOffset() MPSOffset
	SetPrimaryOffset(value MPSOffset)
	// The position of the destination clip rectangle origin relative to the secondary source buffer.
	SecondaryOffset() MPSOffset
	SetSecondaryOffset(value MPSOffset)
	// The edge mode to use when texture reads stray off the edge of the primary source image.
	PrimaryEdgeMode() MPSImageEdgeMode
	SetPrimaryEdgeMode(value MPSImageEdgeMode)
	// The edge mode to use when texture reads stray off the edge of the secondary source image.
	SecondaryEdgeMode() MPSImageEdgeMode
	SetSecondaryEdgeMode(value MPSImageEdgeMode)
	// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
	ClipRect() metal.MTLRegion
	SetClipRect(value metal.MTLRegion)
}

// Init initializes the instance.
func (b MPSBinaryImageKernel) Init() MPSBinaryImageKernel {
	rv := objc.Send[MPSBinaryImageKernel](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSBinaryImageKernel) Autorelease() MPSBinaryImageKernel {
	rv := objc.Send[MPSBinaryImageKernel](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSBinaryImageKernel creates a new MPSBinaryImageKernel instance.
func NewMPSBinaryImageKernel() MPSBinaryImageKernel {
	class := getMPSBinaryImageKernelClass()
	rv := objc.Send[MPSBinaryImageKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewBinaryImageKernelWithCoder(aDecoder foundation.INSCoder) MPSBinaryImageKernel {
	instance := getMPSBinaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSBinaryImageKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewBinaryImageKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSBinaryImageKernel {
	instance := getMPSBinaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSBinaryImageKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(device:)
func NewBinaryImageKernelWithDevice(device metal.MTLDevice) MPSBinaryImageKernel {
	instance := getMPSBinaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSBinaryImageKernelFromID(rv)
}

// This method attempts to apply a kernel in place on a texture.
//
// commandBuffer: A valid command buffer to receive the encoded kernel.
//
// primaryTexture: A pointer to a valid texture containing the primary source image. It will
// not be overwritten.
//
// inPlaceSecondaryTexture: A pointer to a valid texture containing the secondary image. On success,
// the image contents and possibly the texture itself will be replaced with
// the result image.
//
// copyAllocator: An optional block to allocate a new texture to hold the operation results,
// in case in-place operation is not possible. The allocator may use a
// different pixel format or size than the original texture. You may enqueue
// operations on the provided command buffer using the provided compute
// command encoder to initialize the texture contents.
//
// # Return Value
//
// true if the operation succeeded (the texture may have been replaced with a
// new texture if a copy allocator was provided). false if the operation
// failed (the texture is unmodified).
//
// # Discussion
//
// This method attempts to apply the kernel in place on a texture. In-place
// operation means that the same texture is used both to hold the input image
// and the results. Operating in-place can be an excellent way to reduce
// resource utilization, and save time and energy. While simple Metal kernels
// can not operate in place because textures can not be readable and writable
// at the same time, some Metal Performance Shaders kernels can operate in
// place because they use multi-pass algorithms. Whether a kernel can operate
// in-place can depend on current hardware, OS version, and the parameters and
// properties passed to it. You should never assume that a kernel will
// continue to work in place, even if you have observed it doing so before.
//
// If the in-place operation succeeds, this method returns true. If the
// in-place operation fails and no copy allocator is provided, then this
// method returns false. Without a fallback copy allocator, in neither case is
// the pointer held at `texture` modified.
//
// Failure during in-place operation is common. You may find it simplifies
// your code to provide a copy allocator. When an in-place operation fails,
// your copy allocator will be invoked to create a new texture in which to
// write the results, allowing the kernel to proceed reliably out-of-place.
// The original texture will be released, replaced with a pointer to the new
// texture and true will be returned. If the copy allocator returns an invalid
// texture, it is released, `texture` remains unmodified, and false is
// returned.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/encode(commandBuffer:primaryTexture:inPlaceSecondaryTexture:fallbackCopyAllocator:)
func (b MPSBinaryImageKernel) EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, primaryTexture metal.MTLTexture, inPlaceSecondaryTexture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool {
	_block3, _ := NewMTLTextureMPSKernelMTLCommandBufferMTLTextureBlock(copyAllocator)
	rv := objc.Send[bool](b.ID, objc.Sel("encodeToCommandBuffer:primaryTexture:inPlaceSecondaryTexture:fallbackCopyAllocator:"), commandBuffer, primaryTexture, inPlaceSecondaryTexture, _block3)
	return rv
}

// This method attempts to apply a kernel in place on a texture.
//
// commandBuffer: A valid command buffer to receive the encoded kernel.
//
// inPlacePrimaryTexture: A pointer to a valid texture containing the primary image. On success, the
// image contents and possibly the texture itself will be replaced with the
// result image.
//
// secondaryTexture: A pointer to a valid texture containing the secondary source image. It will
// not be overwritten.
//
// copyAllocator: An optional block to allocate a new texture to hold the operation results,
// in case in-place operation is not possible. The allocator may use a
// different pixel format or size than the original texture. You may enqueue
// operations on the provided command buffer using the provided compute
// command encoder to initialize the texture contents.
//
// # Return Value
//
// true if the operation succeeded (the texture may have been replaced with a
// new texture if a copy allocator was provided). false if the operation
// failed (the texture is unmodified).
//
// # Discussion
//
// This method attempts to apply the kernel in place on a texture. In-place
// operation means that the same texture is used both to hold the input image
// and the results. Operating in-place can be an excellent way to reduce
// resource utilization, and save time and energy. While simple Metal kernels
// can not operate in place because textures can not be readable and writable
// at the same time, some Metal Performance Shaders kernels can operate in
// place because they use multi-pass algorithms. Whether a kernel can operate
// in-place can depend on current hardware, OS version, and the parameters and
// properties passed to it. You should never assume that a kernel will
// continue to work in place, even if you have observed it doing so before.
//
// If the in-place operation succeeds, this method returns true. If the
// in-place operation fails and no copy allocator is provided, then this
// method returns false. Without a fallback copy allocator, in neither case is
// the pointer held at `texture` modified.
//
// Failure during in-place operation is common. You may find it simplifies
// your code to provide a copy allocator. When an in-place operation fails,
// your copy allocator will be invoked to create a new texture in which to
// write the results, allowing the kernel to proceed reliably out-of-place.
// The original texture will be released, replaced with a pointer to the new
// texture and true will be returned. If the copy allocator returns an invalid
// texture, it is released, `texture` remains unmodified, and false is
// returned.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/encode(commandBuffer:inPlacePrimaryTexture:secondaryTexture:fallbackCopyAllocator:)
func (b MPSBinaryImageKernel) EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, inPlacePrimaryTexture metal.MTLTexture, secondaryTexture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool {
	_block3, _ := NewMTLTextureMPSKernelMTLCommandBufferMTLTextureBlock(copyAllocator)
	rv := objc.Send[bool](b.ID, objc.Sel("encodeToCommandBuffer:inPlacePrimaryTexture:secondaryTexture:fallbackCopyAllocator:"), commandBuffer, inPlacePrimaryTexture, secondaryTexture, _block3)
	return rv
}

// Encodes a kernel into a command buffer, out-of-place.
//
// commandBuffer: A valid command buffer to receive the encoded kernel.
//
// primaryTexture: A valid texture containing the primary source image.
//
// secondaryTexture: A valid texture containing the secondary source image.
//
// destinationTexture: A valid texture to be overwritten by the result image. `destinationTexture`
// may not alias `primaryTexture` nor `secondaryTexture`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/encode(commandBuffer:primaryTexture:secondaryTexture:destinationTexture:)
func (b MPSBinaryImageKernel) EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, primaryTexture metal.MTLTexture, secondaryTexture metal.MTLTexture, destinationTexture metal.MTLTexture) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeToCommandBuffer:primaryTexture:secondaryTexture:destinationTexture:"), commandBuffer, primaryTexture, secondaryTexture, destinationTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/encode(commandBuffer:primaryImage:secondaryImage:destinationImage:)
func (b MPSBinaryImageKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.MTLCommandBuffer, primaryImage IMPSImage, secondaryImage IMPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationImage)
}

// Determines the region of the primary source texture that will be read for
// an encode operation.
//
// destinationSize: The size of the full virtual destination image.
//
// # Return Value
//
// The area in the virtual source image that will be read.
//
// # Discussion
//
// This method is used to determine which region of the primary source texture
// will be read by the
// [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture]
// method when the filter runs. This information may be needed if the primary
// source image is broken into multiple textures. The size of the full
// (untiled) destination image is provided. The region of the full (untiled)
// source image that will be read is returned. You can then piece together an
// appropriate texture containing that information for use in your tiled
// context.
//
// This method will consult the [MPSBinaryImageKernel.PrimaryOffset] and
// [MPSBinaryImageKernel.ClipRect] properties to determine the full region
// read by the function. Other properties, such as kernel height and width,
// will be consulted as necessary. All properties should be set to their
// intended values prior to calling this method.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/primarySourceRegion(forDestinationSize:)
func (b MPSBinaryImageKernel) PrimarySourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion {
	rv := objc.Send[MPSRegion](b.ID, objc.Sel("primarySourceRegionForDestinationSize:"), destinationSize)
	return MPSRegion(rv)
}

// Determines the region of the secondary source texture that will be read for
// an encode operation.
//
// destinationSize: The size of the full virtual destination image.
//
// # Return Value
//
// The area in the virtual source image that will be read.
//
// # Discussion
//
// This method is used to determine which region of the secondary source
// texture will be read by the
// [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture]
// method when the filter runs. This information may be needed if the
// secondary source image is broken into multiple textures. The size of the
// full (untiled) destination image is provided. The region of the full
// (untiled) source image that will be read is returned. You can then piece
// together an appropriate texture containing that information for use in your
// tiled context.
//
// This method will consult the [MPSBinaryImageKernel.SecondaryOffset] and
// [MPSBinaryImageKernel.ClipRect] properties to determine the full region
// read by the function. Other properties, such as kernel height and width,
// will be consulted as necessary. All properties should be set to their
// intended values prior to calling this method.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/secondarySourceRegion(forDestinationSize:)
func (b MPSBinaryImageKernel) SecondarySourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion {
	rv := objc.Send[MPSRegion](b.ID, objc.Sel("secondarySourceRegionForDestinationSize:"), destinationSize)
	return MPSRegion(rv)
}

// The position of the destination clip rectangle origin relative to the
// primary source buffer.
//
// # Discussion
//
// The offset is defined to be the position of the `origin` value of
// [MPSBinaryImageKernel.ClipRect], in source coordinates.
//
// The default value is `{0, 0, 0}`, indicating that the top left corners of
// the clip rectangle and the primary source image align.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/primaryOffset
func (b MPSBinaryImageKernel) PrimaryOffset() MPSOffset {
	rv := objc.Send[MPSOffset](b.ID, objc.Sel("primaryOffset"))
	return MPSOffset(rv)
}
func (b MPSBinaryImageKernel) SetPrimaryOffset(value MPSOffset) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryOffset:"), value)
}

// The position of the destination clip rectangle origin relative to the
// secondary source buffer.
//
// # Discussion
//
// The offset is defined to be the position of the `origin` value of
// [MPSBinaryImageKernel.ClipRect], in source coordinates.
//
// The default value is `{0, 0, 0}`, indicating that the top left corners of
// the clip rectangle and the secondary source image align.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/secondaryOffset
func (b MPSBinaryImageKernel) SecondaryOffset() MPSOffset {
	rv := objc.Send[MPSOffset](b.ID, objc.Sel("secondaryOffset"))
	return MPSOffset(rv)
}
func (b MPSBinaryImageKernel) SetSecondaryOffset(value MPSOffset) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryOffset:"), value)
}

// The edge mode to use when texture reads stray off the edge of the primary
// source image.
//
// # Discussion
//
// Most kernel objects can read off the edge of a source image. This can
// happen because of a negative offset property, because the `offset +
// clipRect.Size()` is larger than the source image, or because the filter
// uses neighboring pixels in its calculations (e.g. convolution filters).
//
// The default value is usually [MPSImageEdgeModeZero], but some kernels
// default to the [MPSImageEdgeModeClamp] value instead if an edge mode of
// zero is either unsupported or undefined.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/primaryEdgeMode
func (b MPSBinaryImageKernel) PrimaryEdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](b.ID, objc.Sel("primaryEdgeMode"))
	return MPSImageEdgeMode(rv)
}
func (b MPSBinaryImageKernel) SetPrimaryEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}

// The edge mode to use when texture reads stray off the edge of the secondary
// source image.
//
// # Discussion
//
// Most kernel objects can read off the edge of a source image. This can
// happen because of a negative offset property, because the `offset +
// clipRect.Size()` is larger than the source image, or because the filter
// uses neighboring pixels in its calculations (e.g. convolution filters).
//
// The default value is usually [MPSImageEdgeModeZero], but some kernels
// default to the [MPSImageEdgeModeClamp] value instead if an edge mode of
// zero is either unsupported or undefined.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/secondaryEdgeMode
func (b MPSBinaryImageKernel) SecondaryEdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](b.ID, objc.Sel("secondaryEdgeMode"))
	return MPSImageEdgeMode(rv)
}
func (b MPSBinaryImageKernel) SetSecondaryEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}

// An optional clip rectangle to use when writing data. Only the pixels in the
// rectangle will be overwritten.
//
// # Discussion
//
// This value indicates which part of the destination to overwrite. If the
// clip rectangle does not lie completely within the destination image, then
// the intersection between the clip rectangle and destination bounds is used
// instead.
//
// The default value is [MPSRectNoClip], indicating that the entire image is
// used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/clipRect
//
// [MPSRectNoClip]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRectNoClip
func (b MPSBinaryImageKernel) ClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](b.ID, objc.Sel("clipRect"))
	return metal.MTLRegion(rv)
}
func (b MPSBinaryImageKernel) SetClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](b.ID, objc.Sel("setClipRect:"), value)
}
