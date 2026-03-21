// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// A container for writing image data and information produced by a custom image processor.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput
type CIImageProcessorOutput interface {
	objectivec.IObject

	// The base address of CPU memory that your Core Image Processor Kernel can write pixels to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/baseAddress
	BaseAddress() unsafe.Pointer

	// A Metal texture object that can be bound for output using Metal.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/metalTexture
	MetalTexture() metal.MTLTexture

	// An output pixelBuffer object that your Core Image Processor Kernel can write to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/pixelBuffer
	PixelBuffer() corevideo.CVImageBufferRef

	// An output surface object that your Core Image Processor Kernel can write to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/surface
	Surface() iosurface.IOSurfaceRef

	// The rectangular region of the output image that your Core Image Processor Kernel must provide.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/region
	Region() corefoundation.CGRect

	// Returns a Metal command buffer object that can be used for encoding commands.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/metalCommandBuffer
	MetalCommandBuffer() metal.MTLCommandBuffer

	// The bytes per row of the CPU memory that your Core Image Processor Kernel can write pixels to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/bytesPerRow
	BytesPerRow() uintptr

	// The pixel format of the CPU memory that your Core Image Processor Kernel can write pixels to.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/format
	Format() CIFormat

	// A 64-bit digest that uniquely describes the contents of the output of a processor.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/digest
	Digest() uint64
}

// CIImageProcessorOutputObject wraps an existing Objective-C object that conforms to the CIImageProcessorOutput protocol.
type CIImageProcessorOutputObject struct {
	objectivec.Object
}
func (o CIImageProcessorOutputObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIImageProcessorOutputObjectFromID constructs a [CIImageProcessorOutputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIImageProcessorOutputObjectFromID(id objc.ID) CIImageProcessorOutputObject {
	return CIImageProcessorOutputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The base address of CPU memory that your Core Image Processor Kernel can
// write pixels to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/baseAddress
func (o CIImageProcessorOutputObject) BaseAddress() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("baseAddress"))
	return rv
	}
// A Metal texture object that can be bound for output using Metal.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/metalTexture
func (o CIImageProcessorOutputObject) MetalTexture() metal.MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metalTexture"))
	return metal.MTLTextureObjectFromID(rv)
	}
// An output pixelBuffer object that your Core Image Processor Kernel can
// write to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/pixelBuffer
func (o CIImageProcessorOutputObject) PixelBuffer() corevideo.CVImageBufferRef {
	
	rv := objc.Send[corevideo.CVImageBufferRef](o.ID, objc.Sel("pixelBuffer"))
	return rv
	}
// An output surface object that your Core Image Processor Kernel can write
// to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/surface
func (o CIImageProcessorOutputObject) Surface() iosurface.IOSurfaceRef {
	
	rv := objc.Send[iosurface.IOSurfaceRef](o.ID, objc.Sel("surface"))
	return rv
	}
// The rectangular region of the output image that your Core Image Processor
// Kernel must provide.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/region
func (o CIImageProcessorOutputObject) Region() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("region"))
	return rv
	}
// Returns a Metal command buffer object that can be used for encoding
// commands.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/metalCommandBuffer
func (o CIImageProcessorOutputObject) MetalCommandBuffer() metal.MTLCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metalCommandBuffer"))
	return metal.MTLCommandBufferObjectFromID(rv)
	}
// The bytes per row of the CPU memory that your Core Image Processor Kernel
// can write pixels to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/bytesPerRow
func (o CIImageProcessorOutputObject) BytesPerRow() uintptr {
	
	rv := objc.Send[uintptr](o.ID, objc.Sel("bytesPerRow"))
	return rv
	}
// The pixel format of the CPU memory that your Core Image Processor Kernel
// can write pixels to.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/format
func (o CIImageProcessorOutputObject) Format() CIFormat {
	
	rv := objc.Send[CIFormat](o.ID, objc.Sel("format"))
	return rv
	}
// A 64-bit digest that uniquely describes the contents of the output of a
// processor.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorOutput/digest
func (o CIImageProcessorOutputObject) Digest() uint64 {
	
	rv := objc.Send[uint64](o.ID, objc.Sel("digest"))
	return rv
	}

