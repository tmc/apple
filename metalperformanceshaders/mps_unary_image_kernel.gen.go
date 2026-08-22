// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSUnaryImageKernel] class.
var (
	_MPSUnaryImageKernelClass     MPSUnaryImageKernelClass
	_MPSUnaryImageKernelClassOnce sync.Once
)

func getMPSUnaryImageKernelClass() MPSUnaryImageKernelClass {
	_MPSUnaryImageKernelClassOnce.Do(func() {
		_MPSUnaryImageKernelClass = MPSUnaryImageKernelClass{class: objc.GetClass("MPSUnaryImageKernel")}
	})
	return _MPSUnaryImageKernelClass
}

// GetMPSUnaryImageKernelClass returns the class object for MPSUnaryImageKernel.
func GetMPSUnaryImageKernelClass() MPSUnaryImageKernelClass {
	return getMPSUnaryImageKernelClass()
}

type MPSUnaryImageKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSUnaryImageKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSUnaryImageKernelClass) Alloc() MPSUnaryImageKernel {
	rv := objc.Send[MPSUnaryImageKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that consumes one texture and produces one texture.
//
// # Overview
//
// [MPSUnaryImageKernel] defines shared behavior for most image processing
// kernels (filters) such as edging modes, clipping, and tiling support for
// image operations that consumes a single source textures. It is not meant to
// be used directly, but provides API abstraction and in some cases may allow
// some level of polymorphic manipulation of image kernel objects.
//
// # Methods
//
//   - [MPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [MPSUnaryImageKernel.EncodeToCommandBufferSourceImageDestinationImage]
//   - [MPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]: Encodes a kernel into a command buffer, out of place.
//   - [MPSUnaryImageKernel.SourceRegionForDestinationSize]: Determines the region of the source texture that will be read for an encode operation.
//
// # Properties
//
//   - [MPSUnaryImageKernel.Offset]: The position of the destination clip rectangle origin relative to the source buffer.
//   - [MPSUnaryImageKernel.SetOffset]
//   - [MPSUnaryImageKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//   - [MPSUnaryImageKernel.SetClipRect]
//   - [MPSUnaryImageKernel.EdgeMode]: The edge mode to use when texture reads stray off the edge of an image.
//   - [MPSUnaryImageKernel.SetEdgeMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel
type MPSUnaryImageKernel struct {
	MPSKernel
}

// MPSUnaryImageKernelFromID constructs a [MPSUnaryImageKernel] from an objc.ID.
//
// A kernel that consumes one texture and produces one texture.
func MPSUnaryImageKernelFromID(id objc.ID) MPSUnaryImageKernel {
	return MPSUnaryImageKernel{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSUnaryImageKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSUnaryImageKernel] class.
//
// # Methods
//
//   - [IMPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]: This method attempts to apply a kernel in place on a texture.
//   - [IMPSUnaryImageKernel.EncodeToCommandBufferSourceImageDestinationImage]
//   - [IMPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]: Encodes a kernel into a command buffer, out of place.
//   - [IMPSUnaryImageKernel.SourceRegionForDestinationSize]: Determines the region of the source texture that will be read for an encode operation.
//
// # Properties
//
//   - [IMPSUnaryImageKernel.Offset]: The position of the destination clip rectangle origin relative to the source buffer.
//   - [IMPSUnaryImageKernel.SetOffset]
//   - [IMPSUnaryImageKernel.ClipRect]: An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//   - [IMPSUnaryImageKernel.SetClipRect]
//   - [IMPSUnaryImageKernel.EdgeMode]: The edge mode to use when texture reads stray off the edge of an image.
//   - [IMPSUnaryImageKernel.SetEdgeMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel
type IMPSUnaryImageKernel interface {
	IMPSKernel

	// Topic: Methods

	// This method attempts to apply a kernel in place on a texture.
	EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, texture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationImage IMPSImage)
	// Encodes a kernel into a command buffer, out of place.
	EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture)
	// Determines the region of the source texture that will be read for an encode operation.
	SourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion

	// Topic: Properties

	// The position of the destination clip rectangle origin relative to the source buffer.
	Offset() MPSOffset
	SetOffset(value MPSOffset)
	// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
	ClipRect() metal.MTLRegion
	SetClipRect(value metal.MTLRegion)
	// The edge mode to use when texture reads stray off the edge of an image.
	EdgeMode() MPSImageEdgeMode
	SetEdgeMode(value MPSImageEdgeMode)
}

// Init initializes the instance.
func (u MPSUnaryImageKernel) Init() MPSUnaryImageKernel {
	rv := objc.Send[MPSUnaryImageKernel](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MPSUnaryImageKernel) Autorelease() MPSUnaryImageKernel {
	rv := objc.Send[MPSUnaryImageKernel](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSUnaryImageKernel creates a new MPSUnaryImageKernel instance.
func NewMPSUnaryImageKernel() MPSUnaryImageKernel {
	class := getMPSUnaryImageKernelClass()
	rv := objc.Send[MPSUnaryImageKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewUnaryImageKernelWithCoder(aDecoder foundation.INSCoder) MPSUnaryImageKernel {
	instance := getMPSUnaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSUnaryImageKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewUnaryImageKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSUnaryImageKernel {
	instance := getMPSUnaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSUnaryImageKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewUnaryImageKernelWithDevice(device metal.MTLDevice) MPSUnaryImageKernel {
	instance := getMPSUnaryImageKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSUnaryImageKernelFromID(rv)
}

// This method attempts to apply a kernel in place on a texture.
//
// commandBuffer: A valid command buffer to receive the encoded kernel.
//
// texture: A pointer to a valid texture containing the source image. On success, the
// image contents and possibly the texture itself will be replaced with the
// result image.
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
// Failure during in-place operation is very common and will occur
// inconsistently across different hardware platforms and OS versions. Without
// a fallback copy allocator, operating in place may require significant error
// handling code to accompany each call to this method, further complicating
// your code.
//
// You may find it simplifies your code to provide a fallback copy allocator
// so that the operation can proceed reliably even when it can not complete
// in-place. When an in-place filter fails, the copy allocator will be invoked
// to create a new texture in which to write the results, allowing the filter
// to proceed reliably out-of-place. The original texture will be released,
// replaced with a pointer to the new texture and true will be returned. If
// the copy allocator returns an invalid texture, it is released, `texture`
// remains unmodified, and false is returned.
//
// Listing 1. In-Place Operation Example
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:inPlaceTexture:fallbackCopyAllocator:)
func (u MPSUnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer metal.MTLCommandBuffer, texture metal.MTLTexture, copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool {
	_block2, _ := NewMTLTextureMPSKernelMTLCommandBufferMTLTextureBlock(copyAllocator)
	rv := objc.Send[bool](u.ID, objc.Sel("encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator:"), commandBuffer, texture, _block2)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:sourceImage:destinationImage:)
func (u MPSUnaryImageKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}

// Encodes a kernel into a command buffer, out of place.
//
// commandBuffer: A valid command buffer to receive the encoded kernel.
//
// sourceTexture: A valid texture containing the source image.
//
// destinationTexture: A valid texture to be overwritten by the result image. `destinationTexture`
// may not alias `sourceTexture`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:sourceTexture:destinationTexture:)
func (u MPSUnaryImageKernel) EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:"), commandBuffer, sourceTexture, destinationTexture)
}

// Determines the region of the source texture that will be read for an encode
// operation.
//
// destinationSize: The size of the full virtual destination image.
//
// # Return Value
//
// The area in the virtual source image that will be read.
//
// # Discussion
//
// This method is used to determine which region of the source texture will be
// read by the
// [MPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]
// method when the filter runs. This information may be needed if the source
// image is broken into multiple textures. The size of the full (untiled)
// destination image is provided. The region of the full (untiled) source
// image that will be read is returned. You can then piece together an
// appropriate texture containing that information for use in your tiled
// context.
//
// This method will consult the [MPSUnaryImageKernel.Offset] and
// [MPSUnaryImageKernel.ClipRect] properties to determine the full region read
// by the function. Other properties, such as kernel height and width, will be
// consulted as necessary. All properties should be set to their intended
// values prior to calling this method.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/sourceRegion(destinationSize:)
func (u MPSUnaryImageKernel) SourceRegionForDestinationSize(destinationSize metal.MTLSize) MPSRegion {
	rv := objc.Send[MPSRegion](u.ID, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return MPSRegion(rv)
}

// The position of the destination clip rectangle origin relative to the
// source buffer.
//
// # Discussion
//
// The offset is defined to be the position of the `origin` value of
// [MPSUnaryImageKernel.ClipRect], in source coordinates.
//
// The default value is `{0, 0, 0}`, indicating that the top left corners of
// the clip rectangle and the source image align.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/offset
func (u MPSUnaryImageKernel) Offset() MPSOffset {
	rv := objc.Send[MPSOffset](u.ID, objc.Sel("offset"))
	return MPSOffset(rv)
}
func (u MPSUnaryImageKernel) SetOffset(value MPSOffset) {
	objc.Send[struct{}](u.ID, objc.Sel("setOffset:"), value)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/clipRect
//
// [MPSRectNoClip]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRectNoClip
func (u MPSUnaryImageKernel) ClipRect() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](u.ID, objc.Sel("clipRect"))
	return metal.MTLRegion(rv)
}
func (u MPSUnaryImageKernel) SetClipRect(value metal.MTLRegion) {
	objc.Send[struct{}](u.ID, objc.Sel("setClipRect:"), value)
}

// The edge mode to use when texture reads stray off the edge of an image.
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/edgeMode
func (u MPSUnaryImageKernel) EdgeMode() MPSImageEdgeMode {
	rv := objc.Send[MPSImageEdgeMode](u.ID, objc.Sel("edgeMode"))
	return MPSImageEdgeMode(rv)
}
func (u MPSUnaryImageKernel) SetEdgeMode(value MPSImageEdgeMode) {
	objc.Send[struct{}](u.ID, objc.Sel("setEdgeMode:"), value)
}
