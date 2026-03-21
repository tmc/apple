// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIRenderDestination] class.
var (
	_CIRenderDestinationClass     CIRenderDestinationClass
	_CIRenderDestinationClassOnce sync.Once
)

func getCIRenderDestinationClass() CIRenderDestinationClass {
	_CIRenderDestinationClassOnce.Do(func() {
		_CIRenderDestinationClass = CIRenderDestinationClass{class: objc.GetClass("CIRenderDestination")}
	})
	return _CIRenderDestinationClass
}

// GetCIRenderDestinationClass returns the class object for CIRenderDestination.
func GetCIRenderDestinationClass() CIRenderDestinationClass {
	return getCIRenderDestinationClass()
}

type CIRenderDestinationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIRenderDestinationClass) Alloc() CIRenderDestination {
	rv := objc.Send[CIRenderDestination](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A specification for configuring all attributes of a render task’s
// destination and issuing asynchronous render tasks.
//
// # Overview
// 
// The [CIRenderDestination] class provides an API for specifying a render
// task destination’s properties, such as buffer format, alpha mode,
// clamping behavior, blending, and color space, properties formerly tied to
// [CIContext].
// 
// You can create a [CIRenderDestination] object for each surface or buffer to
// which you must render. You can also render multiple times to a single
// destination with different settings such as colorspace and blend mode by
// mutating a single [CIRenderDestination] object between renders.
// 
// Renders issued to a [CIRenderDestination] return to the caller as soon as
// the CPU has issued the task, rather than after the GPU has performed the
// task, so you can start render tasks on subsequent frames without waiting
// for previous renders to finish. If the render fails, a [CIRenderTask] will
// return immediately.
//
// # Creating a Render Destination
//
//   - [CIRenderDestination.InitWithPixelBuffer]: Creates a render destination based on a Core Video pixel buffer.
//   - [CIRenderDestination.InitWithIOSurface]: Creates a render destination based on an [IOSurface] object.
//   - [CIRenderDestination.InitWithMTLTextureCommandBuffer]: Creates a render destination based on a Metal texture.
//   - [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]: Creates a render destination based on a Metal texture with specified pixel format.
//   - [CIRenderDestination.InitWithGLTextureTargetWidthHeight]: Creates a render destination based on an OpenGL texture.
//   - [CIRenderDestination.InitWithBitmapDataWidthHeightBytesPerRowFormat]: Creates a render destination based on a client-managed buffer.
//
// # Customizing Rendering
//
//   - [CIRenderDestination.AlphaMode]: The render destination’s representation of alpha (transparency) values.
//   - [CIRenderDestination.SetAlphaMode]
//   - [CIRenderDestination.BlendKernel]: The destination’s blend kernel.
//   - [CIRenderDestination.SetBlendKernel]
//   - [CIRenderDestination.BlendsInDestinationColorSpace]: Indicator of whether to blend in the destination’s color space.
//   - [CIRenderDestination.SetBlendsInDestinationColorSpace]
//   - [CIRenderDestination.ColorSpace]: The destination’s color space.
//   - [CIRenderDestination.SetColorSpace]
//   - [CIRenderDestination.Width]: The render destination’s row width.
//   - [CIRenderDestination.Height]: The render destination’s buffer height.
//   - [CIRenderDestination.Clamped]: Indicator of whether or not the destination clamps.
//   - [CIRenderDestination.SetClamped]
//   - [CIRenderDestination.Dithered]: Indicator of whether or not the destination dithers.
//   - [CIRenderDestination.SetDithered]
//   - [CIRenderDestination.Flipped]: Indicator of whether the destination is flipped.
//   - [CIRenderDestination.SetFlipped]
//
// # Instance Properties
//
//   - [CIRenderDestination.CaptureTraceURL]: Tell the next render using this destination to capture a Metal trace.
//   - [CIRenderDestination.SetCaptureTraceURL]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination
type CIRenderDestination struct {
	objectivec.Object
}

// CIRenderDestinationFromID constructs a [CIRenderDestination] from an objc.ID.
//
// A specification for configuring all attributes of a render task’s
// destination and issuing asynchronous render tasks.
func CIRenderDestinationFromID(id objc.ID) CIRenderDestination {
	return CIRenderDestination{objectivec.Object{ID: id}}
}
// NOTE: CIRenderDestination adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIRenderDestination] class.
//
// # Creating a Render Destination
//
//   - [ICIRenderDestination.InitWithPixelBuffer]: Creates a render destination based on a Core Video pixel buffer.
//   - [ICIRenderDestination.InitWithIOSurface]: Creates a render destination based on an [IOSurface] object.
//   - [ICIRenderDestination.InitWithMTLTextureCommandBuffer]: Creates a render destination based on a Metal texture.
//   - [ICIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]: Creates a render destination based on a Metal texture with specified pixel format.
//   - [ICIRenderDestination.InitWithGLTextureTargetWidthHeight]: Creates a render destination based on an OpenGL texture.
//   - [ICIRenderDestination.InitWithBitmapDataWidthHeightBytesPerRowFormat]: Creates a render destination based on a client-managed buffer.
//
// # Customizing Rendering
//
//   - [ICIRenderDestination.AlphaMode]: The render destination’s representation of alpha (transparency) values.
//   - [ICIRenderDestination.SetAlphaMode]
//   - [ICIRenderDestination.BlendKernel]: The destination’s blend kernel.
//   - [ICIRenderDestination.SetBlendKernel]
//   - [ICIRenderDestination.BlendsInDestinationColorSpace]: Indicator of whether to blend in the destination’s color space.
//   - [ICIRenderDestination.SetBlendsInDestinationColorSpace]
//   - [ICIRenderDestination.ColorSpace]: The destination’s color space.
//   - [ICIRenderDestination.SetColorSpace]
//   - [ICIRenderDestination.Width]: The render destination’s row width.
//   - [ICIRenderDestination.Height]: The render destination’s buffer height.
//   - [ICIRenderDestination.Clamped]: Indicator of whether or not the destination clamps.
//   - [ICIRenderDestination.SetClamped]
//   - [ICIRenderDestination.Dithered]: Indicator of whether or not the destination dithers.
//   - [ICIRenderDestination.SetDithered]
//   - [ICIRenderDestination.Flipped]: Indicator of whether the destination is flipped.
//   - [ICIRenderDestination.SetFlipped]
//
// # Instance Properties
//
//   - [ICIRenderDestination.CaptureTraceURL]: Tell the next render using this destination to capture a Metal trace.
//   - [ICIRenderDestination.SetCaptureTraceURL]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination
type ICIRenderDestination interface {
	objectivec.IObject

	// Topic: Creating a Render Destination

	// Creates a render destination based on a Core Video pixel buffer.
	InitWithPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIRenderDestination
	// Creates a render destination based on an [IOSurface] object.
	InitWithIOSurface(surface iosurface.IOSurface) CIRenderDestination
	// Creates a render destination based on a Metal texture.
	InitWithMTLTextureCommandBuffer(texture objectivec.IObject, commandBuffer objectivec.IObject) CIRenderDestination
	// Creates a render destination based on a Metal texture with specified pixel format.
	InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat objectivec.IObject, commandBuffer objectivec.IObject, block VoidHandler) CIRenderDestination
	// Creates a render destination based on an OpenGL texture.
	InitWithGLTextureTargetWidthHeight(texture uint32, target uint32, width uint, height uint) CIRenderDestination
	// Creates a render destination based on a client-managed buffer.
	InitWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format int) CIRenderDestination

	// Topic: Customizing Rendering

	// The render destination’s representation of alpha (transparency) values.
	AlphaMode() CIRenderDestinationAlphaMode
	SetAlphaMode(value CIRenderDestinationAlphaMode)
	// The destination’s blend kernel.
	BlendKernel() ICIBlendKernel
	SetBlendKernel(value ICIBlendKernel)
	// Indicator of whether to blend in the destination’s color space.
	BlendsInDestinationColorSpace() bool
	SetBlendsInDestinationColorSpace(value bool)
	// The destination’s color space.
	ColorSpace() coregraphics.CGColorSpaceRef
	SetColorSpace(value coregraphics.CGColorSpaceRef)
	// The render destination’s row width.
	Width() uint
	// The render destination’s buffer height.
	Height() uint
	// Indicator of whether or not the destination clamps.
	Clamped() bool
	SetClamped(value bool)
	// Indicator of whether or not the destination dithers.
	Dithered() bool
	SetDithered(value bool)
	// Indicator of whether the destination is flipped.
	Flipped() bool
	SetFlipped(value bool)

	// Topic: Instance Properties

	// Tell the next render using this destination to capture a Metal trace.
	CaptureTraceURL() foundation.INSURL
	SetCaptureTraceURL(value foundation.INSURL)
}

// Init initializes the instance.
func (r CIRenderDestination) Init() CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CIRenderDestination) Autorelease() CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIRenderDestination creates a new CIRenderDestination instance.
func NewCIRenderDestination() CIRenderDestination {
	class := getCIRenderDestinationClass()
	rv := objc.Send[CIRenderDestination](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a render destination based on a client-managed buffer.
//
// data: Pointer to raw bits of a client-managed buffer that is at least
// (`bytesPerRow` * `height`) bytes in size.
//
// width: Width of the bitmap image in pixels.
//
// height: Height of the bitmap image in pixels.
//
// bytesPerRow: Number of bytes per row of data.
//
// format: Color format specifying how the colors are laid out in memory (for example,
// [RGBA8]).
// //
// [RGBA8]: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBA8
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a client-managed buffer.
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(bitmapData:width:height:bytesPerRow:format:)
func NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format int) CIRenderDestination {
	instance := getCIRenderDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitmapData:width:height:bytesPerRow:format:"), data, width, height, bytesPerRow, format)
	return CIRenderDestinationFromID(rv)
}

// Creates a render destination based on an OpenGL texture.
//
// texture: [GLTexture]-backed texture data.
//
// target: A value denoting the type of destination. Use `GL_TEXTURE_2D` if your
// texture dimensions are a power of two, or `GL_TEXTURE_RECTANGLE_EXT`
// otherwise.
//
// width: Width of the texture in texels.
//
// height: Height of the texture in texels.
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a [GLTexture] supported by
// [GLContext]-backed [CIContext].
//
// # Discussion
// 
// Rendering to a [GLTexture]-backed [CIRenderDestination] is supported by
// only [GLContext]-backed [CIContext].
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(glTexture:target:width:height:)
func NewRenderDestinationWithGLTextureTargetWidthHeight(texture uint32, target uint32, width uint, height uint) CIRenderDestination {
	instance := getCIRenderDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGLTexture:target:width:height:"), texture, target, width, height)
	return CIRenderDestinationFromID(rv)
}

// Creates a render destination based on an [IOSurface] object.
//
// surface: The [IOSurface] render target.
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to an [IOSurface] object.
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created by querying the [IOSurface] object’s attributes.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(ioSurface:)
func NewRenderDestinationWithIOSurface(surface iosurface.IOSurface) CIRenderDestination {
	instance := getCIRenderDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	return CIRenderDestinationFromID(rv)
}

// Creates a render destination based on a Metal texture.
//
// texture: The [MTLTexture] object for rendering with [MTLTextureType] of
// [MTLTextureType.type2D].
// //
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// commandBuffer: An optional [MTLCommandBuffer] to use for rendering to the [MTLTexture]
// destination.
// //
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a Metal buffer.
//
// # Discussion
// 
// Rendering to a [MTLTexture]-backed [CIRenderDestination] is supported by
// only [MTLTexture]-backed [CIContext] objects. The texture must have
// [MTLTextureType] of [MTLTextureType.type2D].
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(mtlTexture:commandBuffer:)
func NewRenderDestinationWithMTLTextureCommandBuffer(texture objectivec.IObject, commandBuffer objectivec.IObject) CIRenderDestination {
	instance := getCIRenderDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMTLTexture:commandBuffer:"), texture, commandBuffer)
	return CIRenderDestinationFromID(rv)
}

// Creates a render destination based on a Core Video pixel buffer.
//
// pixelBuffer: The [CVPixelBuffer] render target.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a [CVPixelBuffer].
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created by querying the [CVPixelBuffer] object’s attributes.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(pixelBuffer:)
func NewRenderDestinationWithPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIRenderDestination {
	instance := getCIRenderDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPixelBuffer:"), pixelBuffer)
	return CIRenderDestinationFromID(rv)
}

// Creates a render destination based on a Core Video pixel buffer.
//
// pixelBuffer: The [CVPixelBuffer] render target.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a [CVPixelBuffer].
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created by querying the [CVPixelBuffer] object’s attributes.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(pixelBuffer:)
func (r CIRenderDestination) InitWithPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("initWithPixelBuffer:"), pixelBuffer)
	return rv
}
// Creates a render destination based on an [IOSurface] object.
//
// surface: The [IOSurface] render target.
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to an [IOSurface] object.
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created by querying the [IOSurface] object’s attributes.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(ioSurface:)
func (r CIRenderDestination) InitWithIOSurface(surface iosurface.IOSurface) CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("initWithIOSurface:"), surface)
	return rv
}
// Creates a render destination based on a Metal texture.
//
// texture: The [MTLTexture] object for rendering with [MTLTextureType] of
// [MTLTextureType.type2D].
// //
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// commandBuffer: An optional [MTLCommandBuffer] to use for rendering to the [MTLTexture]
// destination.
// //
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a Metal buffer.
//
// # Discussion
// 
// Rendering to a [MTLTexture]-backed [CIRenderDestination] is supported by
// only [MTLTexture]-backed [CIContext] objects. The texture must have
// [MTLTextureType] of [MTLTextureType.type2D].
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(mtlTexture:commandBuffer:)
func (r CIRenderDestination) InitWithMTLTextureCommandBuffer(texture objectivec.IObject, commandBuffer objectivec.IObject) CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("initWithMTLTexture:commandBuffer:"), texture, commandBuffer)
	return rv
}
// Creates a render destination based on a Metal texture with specified pixel
// format.
//
// width: Width of the [MTLTexture] that will be returned by block.
// //
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// height: Height of the [MTLTexture] that will be returned by block.
// //
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// pixelFormat: Pixel format of the [MTLTexture] that will be returned by block.
// //
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// commandBuffer: An optional [MTLCommandBuffer] used for rendering to the [MTLTexture].
// //
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// block: [MTLTexture]-rendering provider block to be called lazily when the
// destination is rendered to. The block must return a texture of
// [MTLTextureType] of [MTLTextureType.type2D].
// //
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a Metal texture.
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(width:height:pixelFormat:commandBuffer:mtlTextureProvider:)
func (r CIRenderDestination) InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat objectivec.IObject, commandBuffer objectivec.IObject, block VoidHandler) CIRenderDestination {
_block4, _cleanup4 := NewVoidBlock(block)
	defer _cleanup4()
	rv := objc.Send[objc.ID](r.ID, objc.Sel("initWithWidth:height:pixelFormat:commandBuffer:mtlTextureProvider:"), width, height, pixelFormat, commandBuffer, _block4)
	return CIRenderDestinationFromID(rv)
}
// Creates a render destination based on an OpenGL texture.
//
// texture: [GLTexture]-backed texture data.
//
// target: A value denoting the type of destination. Use `GL_TEXTURE_2D` if your
// texture dimensions are a power of two, or `GL_TEXTURE_RECTANGLE_EXT`
// otherwise.
//
// width: Width of the texture in texels.
//
// height: Height of the texture in texels.
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a [GLTexture] supported by
// [GLContext]-backed [CIContext].
//
// # Discussion
// 
// Rendering to a [GLTexture]-backed [CIRenderDestination] is supported by
// only [GLContext]-backed [CIContext].
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(glTexture:target:width:height:)
func (r CIRenderDestination) InitWithGLTextureTargetWidthHeight(texture uint32, target uint32, width uint, height uint) CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("initWithGLTexture:target:width:height:"), texture, target, width, height)
	return rv
}
// Creates a render destination based on a client-managed buffer.
//
// data: Pointer to raw bits of a client-managed buffer that is at least
// (`bytesPerRow` * `height`) bytes in size.
//
// width: Width of the bitmap image in pixels.
//
// height: Height of the bitmap image in pixels.
//
// bytesPerRow: Number of bytes per row of data.
//
// format: Color format specifying how the colors are laid out in memory (for example,
// [RGBA8]).
// //
// [RGBA8]: https://developer.apple.com/documentation/CoreImage/CIFormat/RGBA8
//
// # Return Value
// 
// A [CIRenderDestination] object for rendering to a client-managed buffer.
//
// # Discussion
// 
// The destination’s [ColorSpace] property will default to a [CGColorSpace]
// created with [sRGB], [extendedSRGB], or [genericGrayGamma2_2].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [extendedSRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/extendedSRGB
// [genericGrayGamma2_2]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/genericGrayGamma2_2
// [sRGB]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/sRGB
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(bitmapData:width:height:bytesPerRow:format:)
func (r CIRenderDestination) InitWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format int) CIRenderDestination {
	rv := objc.Send[CIRenderDestination](r.ID, objc.Sel("initWithBitmapData:width:height:bytesPerRow:format:"), data, width, height, bytesPerRow, format)
	return rv
}

// The render destination’s representation of alpha (transparency) values.
//
// # Discussion
// 
// This property defaults to an appropriate value given the object with which
// you initialized the [CIRenderDestination].
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/alphaMode
func (r CIRenderDestination) AlphaMode() CIRenderDestinationAlphaMode {
	rv := objc.Send[CIRenderDestinationAlphaMode](r.ID, objc.Sel("alphaMode"))
	return CIRenderDestinationAlphaMode(rv)
}
func (r CIRenderDestination) SetAlphaMode(value CIRenderDestinationAlphaMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setAlphaMode:"), value)
}
// The destination’s blend kernel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendKernel
func (r CIRenderDestination) BlendKernel() ICIBlendKernel {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("blendKernel"))
	return CIBlendKernelFromID(objc.ID(rv))
}
func (r CIRenderDestination) SetBlendKernel(value ICIBlendKernel) {
	objc.Send[struct{}](r.ID, objc.Sel("setBlendKernel:"), value)
}
// Indicator of whether to blend in the destination’s color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendsInDestinationColorSpace
func (r CIRenderDestination) BlendsInDestinationColorSpace() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("blendsInDestinationColorSpace"))
	return rv
}
func (r CIRenderDestination) SetBlendsInDestinationColorSpace(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setBlendsInDestinationColorSpace:"), value)
}
// The destination’s color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/colorSpace
func (r CIRenderDestination) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](r.ID, objc.Sel("colorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}
func (r CIRenderDestination) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](r.ID, objc.Sel("setColorSpace:"), value)
}
// The render destination’s row width.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/width
func (r CIRenderDestination) Width() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("width"))
	return rv
}
// The render destination’s buffer height.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/height
func (r CIRenderDestination) Height() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("height"))
	return rv
}
// Indicator of whether or not the destination clamps.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isClamped
func (r CIRenderDestination) Clamped() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isClamped"))
	return rv
}
func (r CIRenderDestination) SetClamped(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setClamped:"), value)
}
// Indicator of whether or not the destination dithers.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isDithered
func (r CIRenderDestination) Dithered() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isDithered"))
	return rv
}
func (r CIRenderDestination) SetDithered(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setDithered:"), value)
}
// Indicator of whether the destination is flipped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isFlipped
func (r CIRenderDestination) Flipped() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isFlipped"))
	return rv
}
func (r CIRenderDestination) SetFlipped(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setFlipped:"), value)
}
// Tell the next render using this destination to capture a Metal trace.
//
// # Discussion
// 
// If this property is set to a file-based URL, then the next render using
// this destination will capture a Metal trace, deleting any existing file if
// present. This property is nil by default.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/captureTraceURL
func (r CIRenderDestination) CaptureTraceURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("captureTraceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (r CIRenderDestination) SetCaptureTraceURL(value foundation.INSURL) {
	objc.Send[struct{}](r.ID, objc.Sel("setCaptureTraceURL:"), value)
}

// InitWithWidthHeightPixelFormatCommandBufferMtlTextureProviderSync is a synchronous wrapper around [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (r CIRenderDestination) InitWithWidthHeightPixelFormatCommandBufferMtlTextureProviderSync(ctx context.Context, width uint, height uint, pixelFormat objectivec.IObject, commandBuffer objectivec.IObject) error {
	done := make(chan struct{}, 1)
	r.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width, height, pixelFormat, commandBuffer, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

