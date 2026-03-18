// Code generated from Apple documentation. DO NOT EDIT.

package corevideo

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

type CVBufferRef uintptr

type CVDisplayLinkOutputCallback = func(uintptr, *CVTimeStamp, *CVTimeStamp, uint64, *uint64, unsafe.Pointer) int

type CVDisplayLinkOutputHandler = func(objectivec.IObject, *CVTimeStamp, *CVTimeStamp, uint64, *uint64) int

type CVDisplayLinkRef uintptr

type CVFillExtendedPixelsCallBack = func(uintptr, unsafe.Pointer) uint8

type CVImageBufferRef uintptr

type CVMetalBufferCacheRef uintptr

type CVMetalBufferRef uintptr

type CVMetalTextureCacheRef uintptr

type CVMetalTextureRef uintptr

type CVOpenGLBufferPoolRef uintptr

type CVOpenGLBufferRef uintptr

type CVOpenGLTextureCacheRef uintptr

type CVOpenGLTextureRef uintptr

type CVOptionFlags = uint64

type CVPixelBufferPoolRef uintptr

type CVPixelBufferRef uintptr

type CVPixelBufferReleaseBytesCallback = func(unsafe.Pointer, unsafe.Pointer)

type CVPixelBufferReleasePlanarBytesCallback = func(unsafe.Pointer, unsafe.Pointer, uint, uint, unsafe.Pointer)

type CVReturn = int32

