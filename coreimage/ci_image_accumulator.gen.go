// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIImageAccumulator] class.
var (
	_CIImageAccumulatorClass     CIImageAccumulatorClass
	_CIImageAccumulatorClassOnce sync.Once
)

func getCIImageAccumulatorClass() CIImageAccumulatorClass {
	_CIImageAccumulatorClassOnce.Do(func() {
		_CIImageAccumulatorClass = CIImageAccumulatorClass{class: objc.GetClass("CIImageAccumulator")}
	})
	return _CIImageAccumulatorClass
}

// GetCIImageAccumulatorClass returns the class object for CIImageAccumulator.
func GetCIImageAccumulatorClass() CIImageAccumulatorClass {
	return getCIImageAccumulatorClass()
}

type CIImageAccumulatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIImageAccumulatorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIImageAccumulatorClass) Alloc() CIImageAccumulator {
	rv := objc.Send[CIImageAccumulator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages feedback-based image processing for tasks such as
// painting or fluid simulation.
//
// # Overview
// 
// The [CIImageAccumulator] class enables feedback-based image processing for
// such things as iterative painting operations or fluid dynamics simulations.
// You use [CIImageAccumulator] objects in conjunction with other Core Image
// classes, such as [CIFilter], [CIImage], [CIVector], and [CIContext], to
// take advantage of the built-in Core Image filters when processing images.
//
// # Initializing an Image Accumulator
//
//   - [CIImageAccumulator.InitWithExtentFormat]: Initializes an image accumulator with the specified extent and pixel format.
//   - [CIImageAccumulator.InitWithExtentFormatColorSpace]: Initializes an image accumulator with the specified extent, pixel format, and color space.
//
// # Setting an Image
//
//   - [CIImageAccumulator.SetImage]: Sets the contents of the image accumulator to the contents of the specified image object.
//   - [CIImageAccumulator.SetImageDirtyRect]: Updates an image accumulator with a subregion of an image object.
//
// # Obtaining Data From an Image Accumulator
//
//   - [CIImageAccumulator.Extent]: The extent of the image associated with the image accumulator.
//   - [CIImageAccumulator.Format]: The pixel format of the image accumulator.
//   - [CIImageAccumulator.Image]: Returns the current contents of the image accumulator.
//
// # Resetting an Accumulator
//
//   - [CIImageAccumulator.Clear]: Resets the accumulator, discarding any pending updates and the current content.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator
type CIImageAccumulator struct {
	objectivec.Object
}

// CIImageAccumulatorFromID constructs a [CIImageAccumulator] from an objc.ID.
//
// An object that manages feedback-based image processing for tasks such as
// painting or fluid simulation.
func CIImageAccumulatorFromID(id objc.ID) CIImageAccumulator {
	return CIImageAccumulator{objectivec.Object{ID: id}}
}
// NOTE: CIImageAccumulator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIImageAccumulator] class.
//
// # Initializing an Image Accumulator
//
//   - [ICIImageAccumulator.InitWithExtentFormat]: Initializes an image accumulator with the specified extent and pixel format.
//   - [ICIImageAccumulator.InitWithExtentFormatColorSpace]: Initializes an image accumulator with the specified extent, pixel format, and color space.
//
// # Setting an Image
//
//   - [ICIImageAccumulator.SetImage]: Sets the contents of the image accumulator to the contents of the specified image object.
//   - [ICIImageAccumulator.SetImageDirtyRect]: Updates an image accumulator with a subregion of an image object.
//
// # Obtaining Data From an Image Accumulator
//
//   - [ICIImageAccumulator.Extent]: The extent of the image associated with the image accumulator.
//   - [ICIImageAccumulator.Format]: The pixel format of the image accumulator.
//   - [ICIImageAccumulator.Image]: Returns the current contents of the image accumulator.
//
// # Resetting an Accumulator
//
//   - [ICIImageAccumulator.Clear]: Resets the accumulator, discarding any pending updates and the current content.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator
type ICIImageAccumulator interface {
	objectivec.IObject

	// Topic: Initializing an Image Accumulator

	// Initializes an image accumulator with the specified extent and pixel format.
	InitWithExtentFormat(extent corefoundation.CGRect, format int) CIImageAccumulator
	// Initializes an image accumulator with the specified extent, pixel format, and color space.
	InitWithExtentFormatColorSpace(extent corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) CIImageAccumulator

	// Topic: Setting an Image

	// Sets the contents of the image accumulator to the contents of the specified image object.
	SetImage(image ICIImage)
	// Updates an image accumulator with a subregion of an image object.
	SetImageDirtyRect(image ICIImage, dirtyRect corefoundation.CGRect)

	// Topic: Obtaining Data From an Image Accumulator

	// The extent of the image associated with the image accumulator.
	Extent() corefoundation.CGRect
	// The pixel format of the image accumulator.
	Format() CIFormat
	// Returns the current contents of the image accumulator.
	Image() ICIImage

	// Topic: Resetting an Accumulator

	// Resets the accumulator, discarding any pending updates and the current content.
	Clear()
}

// Init initializes the instance.
func (i CIImageAccumulator) Init() CIImageAccumulator {
	rv := objc.Send[CIImageAccumulator](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i CIImageAccumulator) Autorelease() CIImageAccumulator {
	rv := objc.Send[CIImageAccumulator](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIImageAccumulator creates a new CIImageAccumulator instance.
func NewCIImageAccumulator() CIImageAccumulator {
	class := getCIImageAccumulatorClass()
	rv := objc.Send[CIImageAccumulator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an image accumulator with the specified extent and pixel
// format.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such askCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// # Return Value
// 
// The initialized image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:)
func NewImageAccumulatorWithExtentFormat(extent corefoundation.CGRect, format int) CIImageAccumulator {
	instance := getCIImageAccumulatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExtent:format:"), extent, format)
	return CIImageAccumulatorFromID(rv)
}

// Initializes an image accumulator with the specified extent, pixel format,
// and color space.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such askCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// colorSpace: A [CGColorSpace] object describing the color space for the image
// accumulator.
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:colorSpace:)
func NewImageAccumulatorWithExtentFormatColorSpace(extent corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) CIImageAccumulator {
	instance := getCIImageAccumulatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExtent:format:colorSpace:"), extent, format, colorSpace)
	return CIImageAccumulatorFromID(rv)
}

// Initializes an image accumulator with the specified extent and pixel
// format.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such askCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// # Return Value
// 
// The initialized image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:)
func (i CIImageAccumulator) InitWithExtentFormat(extent corefoundation.CGRect, format int) CIImageAccumulator {
	rv := objc.Send[CIImageAccumulator](i.ID, objc.Sel("initWithExtent:format:"), extent, format)
	return rv
}
// Initializes an image accumulator with the specified extent, pixel format,
// and color space.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such askCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// colorSpace: A [CGColorSpace] object describing the color space for the image
// accumulator.
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:colorSpace:)
func (i CIImageAccumulator) InitWithExtentFormatColorSpace(extent corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) CIImageAccumulator {
	rv := objc.Send[CIImageAccumulator](i.ID, objc.Sel("initWithExtent:format:colorSpace:"), extent, format, colorSpace)
	return rv
}
// Sets the contents of the image accumulator to the contents of the specified
// image object.
//
// image: The image object whose contents you want to assign to the image
// accumulator.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/setImage(_:)
func (i CIImageAccumulator) SetImage(image ICIImage) {
	objc.Send[objc.ID](i.ID, objc.Sel("setImage:"), image)
}
// Updates an image accumulator with a subregion of an image object.
//
// image: The image object whose contents you want to assign to the image
// accumulator.
//
// dirtyRect: A rectangle that defines the subregion of the image object that’s changed
// since the last time you updated the image accumulator. You must guarantee
// that the new contents differ from the old only within the region specified
// by the this argument.
//
// # Discussion
// 
// For additional details on using this method, see “Imaging Dynamical
// Systems” in [Core Image Programming Guide].
//
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/setImage(_:dirtyRect:)
func (i CIImageAccumulator) SetImageDirtyRect(image ICIImage, dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](i.ID, objc.Sel("setImage:dirtyRect:"), image, dirtyRect)
}
// Returns the current contents of the image accumulator.
//
// # Return Value
// 
// The image object that represents the current contents of the image
// accumulator.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/image()
func (i CIImageAccumulator) Image() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("image"))
	return CIImageFromID(rv)
}
// Resets the accumulator, discarding any pending updates and the current
// content.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/clear()
func (i CIImageAccumulator) Clear() {
	objc.Send[objc.ID](i.ID, objc.Sel("clear"))
}

// Creates an image accumulator with the specified extent and pixel format.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such as kCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// # Return Value
// 
// The image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/imageAccumulatorWithExtent:format:
func (_CIImageAccumulatorClass CIImageAccumulatorClass) ImageAccumulatorWithExtentFormat(extent corefoundation.CGRect, format int) CIImageAccumulator {
	rv := objc.Send[objc.ID](objc.ID(_CIImageAccumulatorClass.class), objc.Sel("imageAccumulatorWithExtent:format:"), extent, format)
	return CIImageAccumulatorFromID(rv)
}
// Creates an image accumulator with the specified extent, pixel format, and
// color space.
//
// extent: A rectangle that specifies the x-value of the rectangle origin, the y-value
// of the rectangle origin, and the width and height.
//
// format: The format and size of each pixel. You must supply a pixel format constant,
// such as kCIFormatARGB8 (32 bit-per-pixel, fixed-point pixel format) or
// kCIFormatRGBAf (128 bit-per-pixel, floating-point pixel format). See
// [CIImage] for more information about pixel format constants.
//
// colorSpace: A [CGColorSpace] object describing the color space for the image
// accumulator.
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The image accumulator object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/imageAccumulatorWithExtent:format:colorSpace:
func (_CIImageAccumulatorClass CIImageAccumulatorClass) ImageAccumulatorWithExtentFormatColorSpace(extent corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) CIImageAccumulator {
	rv := objc.Send[objc.ID](objc.ID(_CIImageAccumulatorClass.class), objc.Sel("imageAccumulatorWithExtent:format:colorSpace:"), extent, format, colorSpace)
	return CIImageAccumulatorFromID(rv)
}

// The extent of the image associated with the image accumulator.
//
// # Discussion
// 
// Extent is a rectangle that specifies the size of the image associated with
// the image accumulator. This rectangle is the size of the complete region of
// the working coordinate space, and is a fixed area. It specifies the x-value
// of the rectangle origin, the y-value of the rectangle origin, and the width
// and height.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/extent
func (i CIImageAccumulator) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("extent"))
	return corefoundation.CGRect(rv)
}
// The pixel format of the image accumulator.
//
// # Discussion
// 
// For applicable values, see Pixel Formats.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/format
func (i CIImageAccumulator) Format() CIFormat {
	rv := objc.Send[CIFormat](i.ID, objc.Sel("format"))
	return CIFormat(rv)
}

