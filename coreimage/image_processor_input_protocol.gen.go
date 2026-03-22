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

// A container of image data and information for use in a custom image processor.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput
type CIImageProcessorInput interface {
	objectivec.IObject

	// The base address of CPU memory that your Core Image Processor Kernel can read pixels from.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/baseAddress
	BaseAddress() unsafe.Pointer

	// A MTLTexture object that can be bound for input using Metal.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/metalTexture
	MetalTexture() metal.MTLTexture

	// An input pixel buffer object that your Core Image Processor Kernel can read from.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/pixelBuffer
	PixelBuffer() corevideo.CVImageBufferRef

	// An input surface object that your Core Image Processor Kernel can read from.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/surface
	Surface() iosurface.IOSurfaceRef

	// The rectangular region of the input image that your Core Image Processor Kernel can use to provide the output.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/region
	Region() corefoundation.CGRect

	// The bytes per row of the CPU memory that your Core Image Processor Kernel can read pixelsfrom.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/bytesPerRow
	BytesPerRow() uintptr

	// The pixel format of the CPU memory that your Core Image Processor Kernel can read pixels from.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/format
	Format() CIFormat

	// A 64-bit digest that uniquely describes the contents of the input to a processor.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/digest
	Digest() uint64

	// This property tells a tiled-input processor how many input tiles will be processed.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/roiTileCount
	RoiTileCount() uint

	// This property tells a tiled-input processor which input tile index is being processed.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/roiTileIndex
	RoiTileIndex() uint
}

// CIImageProcessorInputObject wraps an existing Objective-C object that conforms to the CIImageProcessorInput protocol.
type CIImageProcessorInputObject struct {
	objectivec.Object
}
func (o CIImageProcessorInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIImageProcessorInputObjectFromID constructs a [CIImageProcessorInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIImageProcessorInputObjectFromID(id objc.ID) CIImageProcessorInputObject {
	return CIImageProcessorInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The base address of CPU memory that your Core Image Processor Kernel can
// read pixels from.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/baseAddress
func (o CIImageProcessorInputObject) BaseAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("baseAddress"))
	return rv
	}
// A MTLTexture object that can be bound for input using Metal.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/metalTexture
func (o CIImageProcessorInputObject) MetalTexture() metal.MTLTexture {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metalTexture"))
	return metal.MTLTextureObjectFromID(rv)
	}
// An input pixel buffer object that your Core Image Processor Kernel can read
// from.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/pixelBuffer
func (o CIImageProcessorInputObject) PixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](o.ID, objc.Sel("pixelBuffer"))
	return rv
	}
// An input surface object that your Core Image Processor Kernel can read
// from.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/surface
func (o CIImageProcessorInputObject) Surface() iosurface.IOSurfaceRef {
	rv := objc.Send[iosurface.IOSurfaceRef](o.ID, objc.Sel("surface"))
	return rv
	}
// The rectangular region of the input image that your Core Image Processor
// Kernel can use to provide the output.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/region
func (o CIImageProcessorInputObject) Region() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("region"))
	return rv
	}
// The bytes per row of the CPU memory that your Core Image Processor Kernel
// can read pixelsfrom.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/bytesPerRow
func (o CIImageProcessorInputObject) BytesPerRow() uintptr {
	rv := objc.Send[uintptr](o.ID, objc.Sel("bytesPerRow"))
	return rv
	}
// The pixel format of the CPU memory that your Core Image Processor Kernel
// can read pixels from.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/format
func (o CIImageProcessorInputObject) Format() CIFormat {
	rv := objc.Send[CIFormat](o.ID, objc.Sel("format"))
	return rv
	}
// A 64-bit digest that uniquely describes the contents of the input to a
// processor.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/digest
func (o CIImageProcessorInputObject) Digest() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("digest"))
	return rv
	}
// This property tells a tiled-input processor how many input tiles will be
// processed.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/roiTileCount
func (o CIImageProcessorInputObject) RoiTileCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("roiTileCount"))
	return rv
	}
// This property tells a tiled-input processor which input tile index is being
// processed.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorInput/roiTileIndex
func (o CIImageProcessorInputObject) RoiTileIndex() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("roiTileIndex"))
	return rv
	}

