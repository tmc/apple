// Code generated from Apple documentation. DO NOT EDIT.

package corevideo

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

// CVBufferRef is a reference to a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBuffer
type CVBufferRef uintptr

// CVDisplayLinkOutputCallback is a type for a display link callback function that the system invokes when it’s time for the app to output a video frame.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkOutputCallback
type CVDisplayLinkOutputCallback = func(uintptr, *CVTimeStamp, *CVTimeStamp, uint64, *uint64, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkOutputHandler
type CVDisplayLinkOutputHandler = func(objectivec.IObject, *CVTimeStamp, *CVTimeStamp, uint64, *uint64) int

// CVDisplayLinkRef is a reference to a display link object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLink
type CVDisplayLinkRef uintptr

// CVFillExtendedPixelsCallBack is defines a pointer to a custom extended pixel-fill function, which is called whenever the system needs to pad a buffer holding your custom pixel format.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVFillExtendedPixelsCallBack
type CVFillExtendedPixelsCallBack = func(uintptr, unsafe.Pointer) uint8

// CVImageBufferRef is a reference to a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBuffer
type CVImageBufferRef uintptr

// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCache
type CVMetalBufferCacheRef uintptr

// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBuffer
type CVMetalBufferRef uintptr

// CVMetalTextureCacheRef is a reference to a Core Video Metal texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCache
type CVMetalTextureCacheRef uintptr

// CVMetalTextureRef is a reference to a CoreVideo Metal texture-based image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTexture
type CVMetalTextureRef uintptr

// CVOpenGLBufferPoolRef is a reference to an OpenGL buffer pool object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPool
type CVOpenGLBufferPoolRef uintptr

// CVOpenGLBufferRef is a reference to a Core Video OpenGL buffer object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBuffer
type CVOpenGLBufferRef uintptr

// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCache
type CVOpenGLTextureCacheRef uintptr

// CVOpenGLTextureRef is a reference to an OpenGL texture-based image buffer object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTexture
type CVOpenGLTextureRef uintptr

// CVOptionFlags is the flags to be used for the display link output callback function.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOptionFlags
type CVOptionFlags = uint64

// CVPixelBufferPoolRef is a reference to a pixel buffer pool object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPool
type CVPixelBufferPoolRef uintptr

// CVPixelBufferRef is a reference to a Core Video pixel buffer object.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
type CVPixelBufferRef uintptr

// CVPixelBufferReleaseBytesCallback is a type that defines a release callback function.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleaseBytesCallback
type CVPixelBufferReleaseBytesCallback = func(unsafe.Pointer, unsafe.Pointer)

// CVPixelBufferReleasePlanarBytesCallback is defines a pointer to a pixel buffer release callback function, which is called when a pixel buffer created by [CVPixelBufferCreateWithPlanarBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)] is released.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleasePlanarBytesCallback
type CVPixelBufferReleasePlanarBytesCallback = func(unsafe.Pointer, unsafe.Pointer, uint, uint, unsafe.Pointer)

// CVReturn is a Core Video error type return value.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVReturn
type CVReturn = int32

