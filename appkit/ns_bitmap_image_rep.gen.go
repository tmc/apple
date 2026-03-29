// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBitmapImageRep] class.
var (
	_NSBitmapImageRepClass     NSBitmapImageRepClass
	_NSBitmapImageRepClassOnce sync.Once
)

func getNSBitmapImageRepClass() NSBitmapImageRepClass {
	_NSBitmapImageRepClassOnce.Do(func() {
		_NSBitmapImageRepClass = NSBitmapImageRepClass{class: objc.GetClass("NSBitmapImageRep")}
	})
	return _NSBitmapImageRepClass
}

// GetNSBitmapImageRepClass returns the class object for NSBitmapImageRep.
func GetNSBitmapImageRepClass() NSBitmapImageRepClass {
	return getNSBitmapImageRepClass()
}

type NSBitmapImageRepClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBitmapImageRepClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBitmapImageRepClass) Alloc() NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that renders an image from bitmap data.
//
// # Overview
// 
// Supported bitmap data formats include GIF, JPEG, TIFF, PNG, and various
// permutations of raw bitmap data.
// 
// # Alpha Premultiplication and Bitmap Formats
// 
// When creating a bitmap using a premultiplied format, if a coverage (alpha)
// plane exists, the bitmap’s color components are premultiplied with it. In
// this case, if you modify the contents of the bitmap, you are therefore
// responsible for premultiplying the data. Note that premultiplying generally
// has negligible effect on output quality. For floating-point image data,
// premultiplying color components is a lossless operation, but for
// fixed-point image data, premultiplication can introduce small rounding
// errors. In either case, more rounding errors may appear when compositing
// many premultiplied images; however, such errors are generally not readily
// visible.
// 
// For this reason, you should not use an [NSBitmapImageRep] object if you
// want to manipulate image data. To work with data that is not premultiplied,
// use the Core Graphics framework instead. (Specifically, create images using
// the
// [init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)]
// function and [CGImageAlphaInfo.last] parameter.) Alternatively, include the
// [NSAlphaNonpremultipliedBitmapFormat] flag when creating the bitmap.
//
// [CGImageAlphaInfo.last]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/last
// [NSAlphaNonpremultipliedBitmapFormat]: https://developer.apple.com/documentation/AppKit/NSAlphaNonpremultipliedBitmapFormat
// [init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
//
// # Creating Bitmap Representations of Images
//
//   - [NSBitmapImageRep.ColorizeByMappingGrayToColorBlackMappingWhiteMapping]: Colorizes a grayscale image.
//   - [NSBitmapImageRep.InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel]: Initializes a newly allocated bitmap image representation so it can render the specified image.
//   - [NSBitmapImageRep.InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel]: Initializes a newly allocated bitmap image representation so it can render the specified image.
//   - [NSBitmapImageRep.InitWithCGImage]: Returns a bitmap image representation from a Core Graphics image object.
//   - [NSBitmapImageRep.InitWithCIImage]: Returns a bitmap image representation from a Core Image object.
//   - [NSBitmapImageRep.InitWithData]: Initializes a newly allocated bitmap image representation from the specified data.
//   - [NSBitmapImageRep.InitForIncrementalLoad]: Initializes a newly allocated bitmap image representation for incremental loading.
//
// # Getting Information About Images
//
//   - [NSBitmapImageRep.BitmapFormat]: The format of the bitmap image representation.
//   - [NSBitmapImageRep.BitsPerPixel]: The number of bits allocated for each pixel in each plane of data.
//   - [NSBitmapImageRep.BytesPerPlane]: The number of bytes in each plane or channel of data.
//   - [NSBitmapImageRep.BytesPerRow]: The minimum number of bytes required to specify a scan line in each data plane.
//   - [NSBitmapImageRep.Planar]: A Boolean value that indicates whether the image data is in a planar configuration.
//   - [NSBitmapImageRep.NumberOfPlanes]: The number of separate planes into which the image data is organized.
//   - [NSBitmapImageRep.SamplesPerPixel]: The number of components for each pixel.
//
// # Getting the Bitmap Data
//
//   - [NSBitmapImageRep.BitmapData]: A pointer to the bitmap data.
//   - [NSBitmapImageRep.GetBitmapDataPlanes]: Returns by indirection bitmap data of the bitmap image representation separated into planes.
//
// # Producing Other Representations of Images
//
//   - [NSBitmapImageRep.TIFFRepresentation]: A TIFF representation of the bitmap image data.
//   - [NSBitmapImageRep.TIFFRepresentationUsingCompressionFactor]: Returns a TIFF representation of the image using the specified compression.
//   - [NSBitmapImageRep.RepresentationUsingTypeProperties]: Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object.
//
// # Managing Compression Types
//
//   - [NSBitmapImageRep.CanBeCompressedUsing]: Tests whether the bitmap image representation can be compressed by the specified compression scheme.
//   - [NSBitmapImageRep.SetCompressionFactor]: Sets the bitmap image representation’s compression type and compression factor.
//   - [NSBitmapImageRep.GetCompressionFactor]: Returns by indirection the bitmap image representation’s compression type and compression factor.
//   - [NSBitmapImageRep.SetPropertyWithValue]: Sets the specified property of the bitmap image representation to the specified value.
//   - [NSBitmapImageRep.ValueForProperty]: Returns the value for the specified property.
//
// # Loading Images Incrementally
//
//   - [NSBitmapImageRep.IncrementalLoadFromDataComplete]: Loads the current image data into an incrementally-loaded image representation and returns the current status of the image.
//
// # Managing Pixel Values
//
//   - [NSBitmapImageRep.SetColorAtXY]: Changes the color of the pixel at the specified coordinates.
//   - [NSBitmapImageRep.ColorAtXY]: Returns the color of the pixel at the specified coordinates.
//   - [NSBitmapImageRep.SetPixelAtXY]: Sets the bitmap image representation’s pixel at the specified coordinates to the specified raw pixel values.
//   - [NSBitmapImageRep.GetPixelAtXY]: Returns by indirection the pixel data for the specified location in the bitmap image representation.
//
// # Getting Core Graphics Images
//
//   - [NSBitmapImageRep.CGImage]: A Core Graphics image object based on the bitmap image representation’s data.
//
// # Managing Color Spaces
//
//   - [NSBitmapImageRep.BitmapImageRepByConvertingToColorSpaceRenderingIntent]: Converts the bitmap image representation to the specified color space.
//   - [NSBitmapImageRep.BitmapImageRepByRetaggingWithColorSpace]: Changes the color space tag of the bitmap image representation.
//   - [NSBitmapImageRep.ColorSpace]: The color space of the bitmap.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep
type NSBitmapImageRep struct {
	NSImageRep
}

// NSBitmapImageRepFromID constructs a [NSBitmapImageRep] from an objc.ID.
//
// An object that renders an image from bitmap data.
func NSBitmapImageRepFromID(id objc.ID) NSBitmapImageRep {
	return NSBitmapImageRep{NSImageRep: NSImageRepFromID(id)}
}
// NOTE: NSBitmapImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBitmapImageRep] class.
//
// # Creating Bitmap Representations of Images
//
//   - [INSBitmapImageRep.ColorizeByMappingGrayToColorBlackMappingWhiteMapping]: Colorizes a grayscale image.
//   - [INSBitmapImageRep.InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel]: Initializes a newly allocated bitmap image representation so it can render the specified image.
//   - [INSBitmapImageRep.InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel]: Initializes a newly allocated bitmap image representation so it can render the specified image.
//   - [INSBitmapImageRep.InitWithCGImage]: Returns a bitmap image representation from a Core Graphics image object.
//   - [INSBitmapImageRep.InitWithCIImage]: Returns a bitmap image representation from a Core Image object.
//   - [INSBitmapImageRep.InitWithData]: Initializes a newly allocated bitmap image representation from the specified data.
//   - [INSBitmapImageRep.InitForIncrementalLoad]: Initializes a newly allocated bitmap image representation for incremental loading.
//
// # Getting Information About Images
//
//   - [INSBitmapImageRep.BitmapFormat]: The format of the bitmap image representation.
//   - [INSBitmapImageRep.BitsPerPixel]: The number of bits allocated for each pixel in each plane of data.
//   - [INSBitmapImageRep.BytesPerPlane]: The number of bytes in each plane or channel of data.
//   - [INSBitmapImageRep.BytesPerRow]: The minimum number of bytes required to specify a scan line in each data plane.
//   - [INSBitmapImageRep.Planar]: A Boolean value that indicates whether the image data is in a planar configuration.
//   - [INSBitmapImageRep.NumberOfPlanes]: The number of separate planes into which the image data is organized.
//   - [INSBitmapImageRep.SamplesPerPixel]: The number of components for each pixel.
//
// # Getting the Bitmap Data
//
//   - [INSBitmapImageRep.BitmapData]: A pointer to the bitmap data.
//   - [INSBitmapImageRep.GetBitmapDataPlanes]: Returns by indirection bitmap data of the bitmap image representation separated into planes.
//
// # Producing Other Representations of Images
//
//   - [INSBitmapImageRep.TIFFRepresentation]: A TIFF representation of the bitmap image data.
//   - [INSBitmapImageRep.TIFFRepresentationUsingCompressionFactor]: Returns a TIFF representation of the image using the specified compression.
//   - [INSBitmapImageRep.RepresentationUsingTypeProperties]: Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object.
//
// # Managing Compression Types
//
//   - [INSBitmapImageRep.CanBeCompressedUsing]: Tests whether the bitmap image representation can be compressed by the specified compression scheme.
//   - [INSBitmapImageRep.SetCompressionFactor]: Sets the bitmap image representation’s compression type and compression factor.
//   - [INSBitmapImageRep.GetCompressionFactor]: Returns by indirection the bitmap image representation’s compression type and compression factor.
//   - [INSBitmapImageRep.SetPropertyWithValue]: Sets the specified property of the bitmap image representation to the specified value.
//   - [INSBitmapImageRep.ValueForProperty]: Returns the value for the specified property.
//
// # Loading Images Incrementally
//
//   - [INSBitmapImageRep.IncrementalLoadFromDataComplete]: Loads the current image data into an incrementally-loaded image representation and returns the current status of the image.
//
// # Managing Pixel Values
//
//   - [INSBitmapImageRep.SetColorAtXY]: Changes the color of the pixel at the specified coordinates.
//   - [INSBitmapImageRep.ColorAtXY]: Returns the color of the pixel at the specified coordinates.
//   - [INSBitmapImageRep.SetPixelAtXY]: Sets the bitmap image representation’s pixel at the specified coordinates to the specified raw pixel values.
//   - [INSBitmapImageRep.GetPixelAtXY]: Returns by indirection the pixel data for the specified location in the bitmap image representation.
//
// # Getting Core Graphics Images
//
//   - [INSBitmapImageRep.CGImage]: A Core Graphics image object based on the bitmap image representation’s data.
//
// # Managing Color Spaces
//
//   - [INSBitmapImageRep.BitmapImageRepByConvertingToColorSpaceRenderingIntent]: Converts the bitmap image representation to the specified color space.
//   - [INSBitmapImageRep.BitmapImageRepByRetaggingWithColorSpace]: Changes the color space tag of the bitmap image representation.
//   - [INSBitmapImageRep.ColorSpace]: The color space of the bitmap.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep
type INSBitmapImageRep interface {
	INSImageRep

	// Topic: Creating Bitmap Representations of Images

	// Colorizes a grayscale image.
	ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor INSColor, shadowColor INSColor, lightColor INSColor)
	// Initializes a newly allocated bitmap image representation so it can render the specified image.
	InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, bitmapFormat NSBitmapFormat, rBytes int, pBits int) NSBitmapImageRep
	// Initializes a newly allocated bitmap image representation so it can render the specified image.
	InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, rBytes int, pBits int) NSBitmapImageRep
	// Returns a bitmap image representation from a Core Graphics image object.
	InitWithCGImage(cgImage coregraphics.CGImageRef) NSBitmapImageRep
	// Returns a bitmap image representation from a Core Image object.
	InitWithCIImage(ciImage coreimage.CIImage) NSBitmapImageRep
	// Initializes a newly allocated bitmap image representation from the specified data.
	InitWithData(data foundation.INSData) NSBitmapImageRep
	// Initializes a newly allocated bitmap image representation for incremental loading.
	InitForIncrementalLoad() NSBitmapImageRep

	// Topic: Getting Information About Images

	// The format of the bitmap image representation.
	BitmapFormat() NSBitmapFormat
	// The number of bits allocated for each pixel in each plane of data.
	BitsPerPixel() int
	// The number of bytes in each plane or channel of data.
	BytesPerPlane() int
	// The minimum number of bytes required to specify a scan line in each data plane.
	BytesPerRow() int
	// A Boolean value that indicates whether the image data is in a planar configuration.
	Planar() bool
	// The number of separate planes into which the image data is organized.
	NumberOfPlanes() int
	// The number of components for each pixel.
	SamplesPerPixel() int

	// Topic: Getting the Bitmap Data

	// A pointer to the bitmap data.
	BitmapData() string
	// Returns by indirection bitmap data of the bitmap image representation separated into planes.
	GetBitmapDataPlanes(data unsafe.Pointer)

	// Topic: Producing Other Representations of Images

	// A TIFF representation of the bitmap image data.
	TIFFRepresentation() foundation.INSData
	// Returns a TIFF representation of the image using the specified compression.
	TIFFRepresentationUsingCompressionFactor(comp NSTIFFCompression, factor float32) foundation.INSData
	// Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object.
	RepresentationUsingTypeProperties(storageType NSBitmapImageFileType, properties foundation.INSDictionary) foundation.INSData

	// Topic: Managing Compression Types

	// Tests whether the bitmap image representation can be compressed by the specified compression scheme.
	CanBeCompressedUsing(compression NSTIFFCompression) bool
	// Sets the bitmap image representation’s compression type and compression factor.
	SetCompressionFactor(compression NSTIFFCompression, factor float32)
	// Returns by indirection the bitmap image representation’s compression type and compression factor.
	GetCompressionFactor(compression NSTIFFCompression, factor unsafe.Pointer)
	// Sets the specified property of the bitmap image representation to the specified value.
	SetPropertyWithValue(property NSBitmapImageRepPropertyKey, value objectivec.IObject)
	// Returns the value for the specified property.
	ValueForProperty(property NSBitmapImageRepPropertyKey) objectivec.IObject

	// Topic: Loading Images Incrementally

	// Loads the current image data into an incrementally-loaded image representation and returns the current status of the image.
	IncrementalLoadFromDataComplete(data foundation.INSData, complete bool) int

	// Topic: Managing Pixel Values

	// Changes the color of the pixel at the specified coordinates.
	SetColorAtXY(color INSColor, x int, y int)
	// Returns the color of the pixel at the specified coordinates.
	ColorAtXY(x int, y int) INSColor
	// Sets the bitmap image representation’s pixel at the specified coordinates to the specified raw pixel values.
	SetPixelAtXY(p uint, x int, y int)
	// Returns by indirection the pixel data for the specified location in the bitmap image representation.
	GetPixelAtXY(p uint, x int, y int)

	// Topic: Getting Core Graphics Images

	// A Core Graphics image object based on the bitmap image representation’s data.
	CGImage() coregraphics.CGImageRef

	// Topic: Managing Color Spaces

	// Converts the bitmap image representation to the specified color space.
	BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace INSColorSpace, renderingIntent NSColorRenderingIntent) INSBitmapImageRep
	// Changes the color space tag of the bitmap image representation.
	BitmapImageRepByRetaggingWithColorSpace(newSpace INSColorSpace) INSBitmapImageRep
	// The color space of the bitmap.
	ColorSpace() INSColorSpace
}

// Init initializes the instance.
func (b NSBitmapImageRep) Init() NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBitmapImageRep) Autorelease() NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBitmapImageRep creates a new NSBitmapImageRep instance.
func NewNSBitmapImageRep() NSBitmapImageRep {
	class := getNSBitmapImageRepClass()
	rv := objc.Send[NSBitmapImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated bitmap image representation for incremental
// loading.
//
// # Discussion
// 
// The receiver returns itself after setting its size and data buffer to zero.
// You can then call [IncrementalLoadFromDataComplete] to incrementally add
// image data.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(forIncrementalLoad:)
func NewBitmapImageRepForIncrementalLoad() NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForIncrementalLoad"))
	return NSBitmapImageRepFromID(rv)
}

// Initializes a newly allocated bitmap image representation so it can render
// the specified image.
//
// planes: An array of character pointers, each of which points to a buffer containing
// raw image data. If the data is in planar configuration, each buffer holds
// one component—one plane—of the data. Color planes are arranged in the
// standard order—for example, red before green before blue for RGB color.
// All color planes precede the coverage plane. If a coverage plane exists,
// the bitmap’s color components must be premultiplied with it. If the data
// is in meshed configuration (that is, `isPlanar` is [false]), only the first
// buffer is read.
// 
// If `planes` is [NULL] or an array of [NULL] pointers, this method allocates
// enough memory to hold the image described by the other arguments. You can
// then obtain pointers to this memory (with the [GetPixelAtXY] method or
// [BitmapData] property) and fill in the image data. In this case, the
// allocated memory will belong to the object and will be freed when it’s
// freed.
// 
// If `planes` is not [NULL] and the array contains at least one data pointer,
// the returned object will only reference the image data; it will not copy
// it. The object treats the image data in the buffers as immutable and will
// not attempt to alter it. When the object itself is freed, it will not
// attempt to free the buffers.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// width: The width of the image in pixels. This value must be greater than 0.
//
// height: The height of the image in pixels. This value must be greater than 0.
//
// bps: The number of bits used to specify one pixel in a single component of the
// data. All components are assumed to have the same bits per sample. `bps`
// should be one of these values: 1, 2, 4, 8, 12, or 16.
//
// spp: The number of data components, or samples per pixel. This value includes
// both color components and the coverage component (alpha), if present.
// Meaningful values range from 1 through 5. An image with cyan, magenta,
// yellow, and black (CMYK) color components plus a coverage component would
// have an `spp` of 5; a grayscale image that lacks a coverage component would
// have an `spp` of 1.
//
// alpha: [true] if one of the components counted in the number of samples per pixel
// (`spp`) is a coverage (alpha) component, and [false] if there is no
// coverage component. If [true], the color components in the bitmap data must
// be premultiplied with their coverage component.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// isPlanar: [true] if the data components are laid out in a series of separate
// “planes” or channels (“planar configuration”) and [false] if
// component values are interwoven in a single channel (“meshed
// configuration”). If [false], only the first buffer of `planes` is read.
// 
// For example, in meshed configuration, the red, green, blue, and coverage
// values for the first pixel of an image would precede the red, green, blue,
// and coverage values for the second pixel, and so on. In planar
// configuration, red values for all the pixels in the image would precede all
// green values, which would precede all blue values, which would precede all
// coverage values.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// colorSpaceName: A constant that indicates how data values are to be interpreted. It should
// be one of the constants in [NSColorSpaceName].
// 
// If `bps` is 12, you cannot specify the monochrome color space.
//
// bitmapFormat: An integer that specifies the ordering of the bitmap components. It is a
// mask created by combining the [NSBitmapImageRep.Format] constants
// [BitmapFormatAlphaFirst], [BitmapFormatAlphaNonpremultiplied] and
// [BitmapFormatFloatingPointSamples] using the C bitwise OR operator.
// //
// [NSBitmapImageRep.Format]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format
//
// rBytes: The number of bytes that are allocated for each scan line in each plane of
// data. A scan line is a single row of pixels spanning the width of the
// image.
// 
// Normally, `rowBytes` can be figured from the width of the image, the number
// of bits per pixel in each sample (`bps`), and, if the data is in a meshed
// configuration, the number of samples per pixel (`spp`). However, if the
// data for each row is aligned on word or other boundaries, it may have been
// necessary to allocate more memory for each row than there is data to fill
// it. `rowBytes` lets the object know whether that’s the case.
// 
// If you pass in a `rowBytes` value of 0, the bitmap data allocated may be
// padded to fall on long word or larger boundaries for performance. If your
// code wants to advance row by row, use [BytesPerRow] and do not assume the
// data is packed. Passing in a non-zero value allows you to specify exact row
// advances.
//
// pBits: This integer value informs [NSBitmapImageRep] how many bits are actually
// allocated per pixel in each plane of data. If the data is in planar
// configuration, this normally equals `bps` (bits per sample). If the data is
// in meshed configuration, it normally equals `bps` times `spp` (samples per
// pixel). However, it’s possible for a pixel specification to be followed
// by some meaningless bits (empty space), as may happen, for example, if
// pixel data is aligned on byte boundaries. [NSBitmapImageRep] supports only
// a limited number of `pixelBits` values (other than the default): for RGB
// images with 4 `bps`, `pixelBits` may be 16; for RGB images with 8 `bps`,
// `pixelBits` may be 32. The legal values for `pixelBits` are system
// dependent.
// 
// If you specify 0 for this parameter, the object interprets the number of
// bits per pixel using the values in the `bps` and `spp` parameters, as
// described in the preceding paragraph, without any meaningless bits.
//
// # Return Value
// 
// An initialized [NSBitmapImageRep] object or `nil` if the object cannot be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:)
func NewBitmapImageRepWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, bitmapFormat NSBitmapFormat, rBytes int, pBits int) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, objc.String(string(colorSpaceName)), bitmapFormat, rBytes, pBits)
	return NSBitmapImageRepFromID(rv)
}

// Initializes a newly allocated bitmap image representation so it can render
// the specified image.
//
// planes: An array of character pointers, each of which points to a buffer containing
// raw image data. If the data is in planar configuration, each buffer holds
// one component—one plane—of the data. Color planes are arranged in the
// standard order—for example, red before green before blue for RGB color.
// All color planes precede the coverage plane. If a coverage plane exists,
// the bitmap’s color components must be premultiplied with it. If the data
// is in meshed configuration (that is, `isPlanar` is [false]), only the first
// buffer is read.
// 
// If `planes` is [NULL] or an array of [NULL] pointers, this method allocates
// enough memory to hold the image described by the other arguments. You can
// then obtain pointers to this memory (with the [GetPixelAtXY] method or
// [BitmapData] property) and fill in the image data. In this case, the
// allocated memory will belong to the object and will be freed when it’s
// freed.
// 
// If `planes` is not [NULL] and the array contains at least one data pointer,
// the returned object will only reference the image data; it will not copy
// it. The object treats the image data in the buffers as immutable and will
// not attempt to alter it. When the object itself is freed, it will not
// attempt to free the buffers.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// width: The width of the image in pixels. This value must be greater than 0.
//
// height: The height of the image in pixels. This value must be greater than 0.
//
// bps: The number of bits used to specify one pixel in a single component of the
// data. All components are assumed to have the same bits per sample. `bps`
// should be one of these values: 1, 2, 4, 8, 12, or 16.
//
// spp: The number of data components, or samples per pixel. This value includes
// both color components and the coverage component (alpha), if present.
// Meaningful values range from 1 through 5. An image with cyan, magenta,
// yellow, and black (CMYK) color components plus a coverage component would
// have an `spp` of 5; a grayscale image that lacks a coverage component would
// have an `spp` of 1.
//
// alpha: [true] if one of the components counted in the number of samples per pixel
// (`spp`) is a coverage (alpha) component, and [false] if there is no
// coverage component. If [true], the color components in the bitmap data must
// be premultiplied with their coverage component.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// isPlanar: [true] if the data components are laid out in a series of separate
// “planes” or channels (“planar configuration”) and [false] if
// component values are interwoven in a single channel (“meshed
// configuration”). If [false], only the first buffer of `planes` is read.
// 
// For example, in meshed configuration, the red, green, blue, and coverage
// values for the first pixel of an image would precede the red, green, blue,
// and coverage values for the second pixel, and so on. In planar
// configuration, red values for all the pixels in the image would precede all
// green values, which would precede all blue values, which would precede all
// coverage values.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// colorSpaceName: A constant that indicates how data values are to be interpreted. It should
// be one of the constants in [NSColorSpaceName].
// 
// If `bps` is 12, you cannot specify the monochrome color space.
//
// rBytes: The number of bytes that are allocated for each scan line in each plane of
// data. A scan line is a single row of pixels spanning the width of the
// image.
// 
// Normally, `rowBytes` can be figured from the width of the image, the number
// of bits per pixel in each sample (`bps`), and, if the data is in a meshed
// configuration, the number of samples per pixel (`spp`). However, if the
// data for each row is aligned on word or other boundaries, it may have been
// necessary to allocate more memory for each row than there is data to fill
// it. `rowBytes` lets the object know whether that’s the case.
// 
// If you pass in a `rowBytes` value of 0, the bitmap data allocated may be
// padded to fall on long word or larger boundaries for performance. If your
// code wants to advance row by row, use [BytesPerRow] and do not assume the
// data is packed. Passing in a non-zero value allows you to specify exact row
// advances.
//
// pBits: This integer value informs [NSBitmapImageRep] how many bits are actually
// allocated per pixel in each plane of data. If the data is in planar
// configuration, this normally equals `bps` (bits per sample). If the data is
// in meshed configuration, it normally equals `bps` times `spp` (samples per
// pixel). However, it’s possible for a pixel specification to be followed
// by some meaningless bits (empty space), as may happen, for example, if
// pixel data is aligned on byte boundaries. [NSBitmapImageRep] supports only
// a limited number of `pixelBits` values (other than the default): for RGB
// images with 4 `bps`, `pixelBits` may be 16; for RGB images with 8 `bps`,
// `pixelBits` may be 32. The legal values for `pixelBits` are system
// dependent.
// 
// If you specify 0 for this parameter, the object interprets the number of
// bits per pixel using the values in the `bps` and `spp` parameters, as
// described in the preceding paragraph, without any meaningless bits.
//
// # Return Value
// 
// An initialized [NSBitmapImageRep] object or `nil` if the object cannot be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:)
func NewBitmapImageRepWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, rBytes int, pBits int) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, objc.String(string(colorSpaceName)), rBytes, pBits)
	return NSBitmapImageRepFromID(rv)
}

// Returns a bitmap image representation from a Core Graphics image object.
//
// cgImage: A Core Graphics image object (an opaque type) from which to create the
// receiver. This opaque type is retained.
//
// # Return Value
// 
// An [NSBitmapImageRep] object initialized from the contents of the Core
// Graphics image.
//
// # Discussion
// 
// If you use this method, you should treat the resulting bitmap
// [NSBitmapImageRep] object as read only. Because it only retains the value
// in the `cgImage` parameter, rather than unpacking the data, accessing the
// pixel data requires the creation of a copy of that data in memory. Changes
// to that data are not saved back to the Core Graphics image.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(cgImage:)
func NewBitmapImageRepWithCGImage(cgImage coregraphics.CGImageRef) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:"), cgImage)
	return NSBitmapImageRepFromID(rv)
}

// Returns a bitmap image representation from a Core Image object.
//
// ciImage: A Core Image object whose contents are to be copied to the receiver. This
// image rectangle must be of a finite size.
//
// # Return Value
// 
// An [NSBitmapImageRep] object initialized from the contents of the Core
// Image ([CIImage]) object.
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// # Discussion
// 
// The image in the `ciImage` parameter must be fully rendered before the
// receiver can be initialized. If you specify an object whose rendering was
// deferred (and thus does not have any pixels available now), this method
// forces the image to be rendered immediately. Rendering the image could
// result in a performance penalty if the image has a complex rendering chain
// or accelerated rendering hardware is not available. Rendering uses the
// current graphics context in the thread from which this method is called; to
// ensure consistent results across multiple threads, set the current context
// using the [NSGraphicsContext] class before calling this method.
// 
// By the time this method returns, the resultant [NSBitmapImageRep] object
// can have its raw pixel data inspected, can be put on the pasteboard, and
// can be encoded to any of the standard image formats that [NSBitmapImageRep]
// supports (JPEG, TIFF, and so on.)
// 
// If you pass in a [CIImage] object whose extents are not finite, this method
// raises an exception.
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(ciImage:)
func NewBitmapImageRepWithCIImage(ciImage coreimage.CIImage) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCIImage:"), ciImage)
	return NSBitmapImageRepFromID(rv)
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewBitmapImageRepWithCoder(coder foundation.INSCoder) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSBitmapImageRepFromID(rv)
}

// Initializes a newly allocated bitmap image representation from the
// specified data.
//
// data: A data object containing image data. The contents of `bitmapData` can be
// any supported bitmap format. For TIFF data, the [NSBitmapImageRep] is
// initialized from the first header and image data found in `bitmapData`.
//
// # Return Value
// 
// Returns an initialized [NSBitmapImageRep] if the initialization was
// successful or `nil` if it was unable to interpret the contents of
// `bitmapData`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(data:)
func NewBitmapImageRepWithData(data foundation.INSData) NSBitmapImageRep {
	instance := getNSBitmapImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSBitmapImageRepFromID(rv)
}

// Colorizes a grayscale image.
//
// midPoint: A float value representing the midpoint of the grayscale image.
//
// midPointColor: A color object representing the midpoint of the color to map the image to.
//
// shadowColor: A color object representing the black mapping to use for shadows.
//
// lightColor: A color object representing the white mapping to be used in the image.
//
// # Discussion
// 
// This method maps the receiver such that:
// 
// - Gray value of `midPoint` –> `midPointColor`; - black –>
// `shadowColor`; - white –> `lightColor`.
// 
// It works on images with 8-bit SPP, and thus supports either 8-bit gray or
// 24-bit color (with optional alpha).
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorize(byMappingGray:to:blackMapping:whiteMapping:)
func (b NSBitmapImageRep) ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor INSColor, shadowColor INSColor, lightColor INSColor) {
	objc.Send[objc.ID](b.ID, objc.Sel("colorizeByMappingGray:toColor:blackMapping:whiteMapping:"), midPoint, midPointColor, shadowColor, lightColor)
}
// Initializes a newly allocated bitmap image representation so it can render
// the specified image.
//
// planes: An array of character pointers, each of which points to a buffer containing
// raw image data. If the data is in planar configuration, each buffer holds
// one component—one plane—of the data. Color planes are arranged in the
// standard order—for example, red before green before blue for RGB color.
// All color planes precede the coverage plane. If a coverage plane exists,
// the bitmap’s color components must be premultiplied with it. If the data
// is in meshed configuration (that is, `isPlanar` is [false]), only the first
// buffer is read.
// 
// If `planes` is [NULL] or an array of [NULL] pointers, this method allocates
// enough memory to hold the image described by the other arguments. You can
// then obtain pointers to this memory (with the [GetPixelAtXY] method or
// [BitmapData] property) and fill in the image data. In this case, the
// allocated memory will belong to the object and will be freed when it’s
// freed.
// 
// If `planes` is not [NULL] and the array contains at least one data pointer,
// the returned object will only reference the image data; it will not copy
// it. The object treats the image data in the buffers as immutable and will
// not attempt to alter it. When the object itself is freed, it will not
// attempt to free the buffers.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// width: The width of the image in pixels. This value must be greater than 0.
//
// height: The height of the image in pixels. This value must be greater than 0.
//
// bps: The number of bits used to specify one pixel in a single component of the
// data. All components are assumed to have the same bits per sample. `bps`
// should be one of these values: 1, 2, 4, 8, 12, or 16.
//
// spp: The number of data components, or samples per pixel. This value includes
// both color components and the coverage component (alpha), if present.
// Meaningful values range from 1 through 5. An image with cyan, magenta,
// yellow, and black (CMYK) color components plus a coverage component would
// have an `spp` of 5; a grayscale image that lacks a coverage component would
// have an `spp` of 1.
//
// alpha: [true] if one of the components counted in the number of samples per pixel
// (`spp`) is a coverage (alpha) component, and [false] if there is no
// coverage component. If [true], the color components in the bitmap data must
// be premultiplied with their coverage component.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// isPlanar: [true] if the data components are laid out in a series of separate
// “planes” or channels (“planar configuration”) and [false] if
// component values are interwoven in a single channel (“meshed
// configuration”). If [false], only the first buffer of `planes` is read.
// 
// For example, in meshed configuration, the red, green, blue, and coverage
// values for the first pixel of an image would precede the red, green, blue,
// and coverage values for the second pixel, and so on. In planar
// configuration, red values for all the pixels in the image would precede all
// green values, which would precede all blue values, which would precede all
// coverage values.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// colorSpaceName: A constant that indicates how data values are to be interpreted. It should
// be one of the constants in [NSColorSpaceName].
// 
// If `bps` is 12, you cannot specify the monochrome color space.
//
// bitmapFormat: An integer that specifies the ordering of the bitmap components. It is a
// mask created by combining the [NSBitmapImageRep.Format] constants
// [BitmapFormatAlphaFirst], [BitmapFormatAlphaNonpremultiplied] and
// [BitmapFormatFloatingPointSamples] using the C bitwise OR operator.
// //
// [NSBitmapImageRep.Format]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format
//
// rBytes: The number of bytes that are allocated for each scan line in each plane of
// data. A scan line is a single row of pixels spanning the width of the
// image.
// 
// Normally, `rowBytes` can be figured from the width of the image, the number
// of bits per pixel in each sample (`bps`), and, if the data is in a meshed
// configuration, the number of samples per pixel (`spp`). However, if the
// data for each row is aligned on word or other boundaries, it may have been
// necessary to allocate more memory for each row than there is data to fill
// it. `rowBytes` lets the object know whether that’s the case.
// 
// If you pass in a `rowBytes` value of 0, the bitmap data allocated may be
// padded to fall on long word or larger boundaries for performance. If your
// code wants to advance row by row, use [BytesPerRow] and do not assume the
// data is packed. Passing in a non-zero value allows you to specify exact row
// advances.
//
// pBits: This integer value informs [NSBitmapImageRep] how many bits are actually
// allocated per pixel in each plane of data. If the data is in planar
// configuration, this normally equals `bps` (bits per sample). If the data is
// in meshed configuration, it normally equals `bps` times `spp` (samples per
// pixel). However, it’s possible for a pixel specification to be followed
// by some meaningless bits (empty space), as may happen, for example, if
// pixel data is aligned on byte boundaries. [NSBitmapImageRep] supports only
// a limited number of `pixelBits` values (other than the default): for RGB
// images with 4 `bps`, `pixelBits` may be 16; for RGB images with 8 `bps`,
// `pixelBits` may be 32. The legal values for `pixelBits` are system
// dependent.
// 
// If you specify 0 for this parameter, the object interprets the number of
// bits per pixel using the values in the `bps` and `spp` parameters, as
// described in the preceding paragraph, without any meaningless bits.
//
// # Return Value
// 
// An initialized [NSBitmapImageRep] object or `nil` if the object cannot be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:)
func (b NSBitmapImageRep) InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, bitmapFormat NSBitmapFormat, rBytes int, pBits int) NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, objc.String(string(colorSpaceName)), bitmapFormat, rBytes, pBits)
	return rv
}
// Initializes a newly allocated bitmap image representation so it can render
// the specified image.
//
// planes: An array of character pointers, each of which points to a buffer containing
// raw image data. If the data is in planar configuration, each buffer holds
// one component—one plane—of the data. Color planes are arranged in the
// standard order—for example, red before green before blue for RGB color.
// All color planes precede the coverage plane. If a coverage plane exists,
// the bitmap’s color components must be premultiplied with it. If the data
// is in meshed configuration (that is, `isPlanar` is [false]), only the first
// buffer is read.
// 
// If `planes` is [NULL] or an array of [NULL] pointers, this method allocates
// enough memory to hold the image described by the other arguments. You can
// then obtain pointers to this memory (with the [GetPixelAtXY] method or
// [BitmapData] property) and fill in the image data. In this case, the
// allocated memory will belong to the object and will be freed when it’s
// freed.
// 
// If `planes` is not [NULL] and the array contains at least one data pointer,
// the returned object will only reference the image data; it will not copy
// it. The object treats the image data in the buffers as immutable and will
// not attempt to alter it. When the object itself is freed, it will not
// attempt to free the buffers.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// width: The width of the image in pixels. This value must be greater than 0.
//
// height: The height of the image in pixels. This value must be greater than 0.
//
// bps: The number of bits used to specify one pixel in a single component of the
// data. All components are assumed to have the same bits per sample. `bps`
// should be one of these values: 1, 2, 4, 8, 12, or 16.
//
// spp: The number of data components, or samples per pixel. This value includes
// both color components and the coverage component (alpha), if present.
// Meaningful values range from 1 through 5. An image with cyan, magenta,
// yellow, and black (CMYK) color components plus a coverage component would
// have an `spp` of 5; a grayscale image that lacks a coverage component would
// have an `spp` of 1.
//
// alpha: [true] if one of the components counted in the number of samples per pixel
// (`spp`) is a coverage (alpha) component, and [false] if there is no
// coverage component. If [true], the color components in the bitmap data must
// be premultiplied with their coverage component.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// isPlanar: [true] if the data components are laid out in a series of separate
// “planes” or channels (“planar configuration”) and [false] if
// component values are interwoven in a single channel (“meshed
// configuration”). If [false], only the first buffer of `planes` is read.
// 
// For example, in meshed configuration, the red, green, blue, and coverage
// values for the first pixel of an image would precede the red, green, blue,
// and coverage values for the second pixel, and so on. In planar
// configuration, red values for all the pixels in the image would precede all
// green values, which would precede all blue values, which would precede all
// coverage values.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// colorSpaceName: A constant that indicates how data values are to be interpreted. It should
// be one of the constants in [NSColorSpaceName].
// 
// If `bps` is 12, you cannot specify the monochrome color space.
//
// rBytes: The number of bytes that are allocated for each scan line in each plane of
// data. A scan line is a single row of pixels spanning the width of the
// image.
// 
// Normally, `rowBytes` can be figured from the width of the image, the number
// of bits per pixel in each sample (`bps`), and, if the data is in a meshed
// configuration, the number of samples per pixel (`spp`). However, if the
// data for each row is aligned on word or other boundaries, it may have been
// necessary to allocate more memory for each row than there is data to fill
// it. `rowBytes` lets the object know whether that’s the case.
// 
// If you pass in a `rowBytes` value of 0, the bitmap data allocated may be
// padded to fall on long word or larger boundaries for performance. If your
// code wants to advance row by row, use [BytesPerRow] and do not assume the
// data is packed. Passing in a non-zero value allows you to specify exact row
// advances.
//
// pBits: This integer value informs [NSBitmapImageRep] how many bits are actually
// allocated per pixel in each plane of data. If the data is in planar
// configuration, this normally equals `bps` (bits per sample). If the data is
// in meshed configuration, it normally equals `bps` times `spp` (samples per
// pixel). However, it’s possible for a pixel specification to be followed
// by some meaningless bits (empty space), as may happen, for example, if
// pixel data is aligned on byte boundaries. [NSBitmapImageRep] supports only
// a limited number of `pixelBits` values (other than the default): for RGB
// images with 4 `bps`, `pixelBits` may be 16; for RGB images with 8 `bps`,
// `pixelBits` may be 32. The legal values for `pixelBits` are system
// dependent.
// 
// If you specify 0 for this parameter, the object interprets the number of
// bits per pixel using the values in the `bps` and `spp` parameters, as
// described in the preceding paragraph, without any meaningless bits.
//
// # Return Value
// 
// An initialized [NSBitmapImageRep] object or `nil` if the object cannot be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:)
func (b NSBitmapImageRep) InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName NSColorSpaceName, rBytes int, pBits int) NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, objc.String(string(colorSpaceName)), rBytes, pBits)
	return rv
}
// Returns a bitmap image representation from a Core Graphics image object.
//
// cgImage: A Core Graphics image object (an opaque type) from which to create the
// receiver. This opaque type is retained.
//
// # Return Value
// 
// An [NSBitmapImageRep] object initialized from the contents of the Core
// Graphics image.
//
// # Discussion
// 
// If you use this method, you should treat the resulting bitmap
// [NSBitmapImageRep] object as read only. Because it only retains the value
// in the `cgImage` parameter, rather than unpacking the data, accessing the
// pixel data requires the creation of a copy of that data in memory. Changes
// to that data are not saved back to the Core Graphics image.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(cgImage:)
func (b NSBitmapImageRep) InitWithCGImage(cgImage coregraphics.CGImageRef) NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initWithCGImage:"), cgImage)
	return rv
}
// Returns a bitmap image representation from a Core Image object.
//
// ciImage: A Core Image object whose contents are to be copied to the receiver. This
// image rectangle must be of a finite size.
//
// # Return Value
// 
// An [NSBitmapImageRep] object initialized from the contents of the Core
// Image ([CIImage]) object.
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// # Discussion
// 
// The image in the `ciImage` parameter must be fully rendered before the
// receiver can be initialized. If you specify an object whose rendering was
// deferred (and thus does not have any pixels available now), this method
// forces the image to be rendered immediately. Rendering the image could
// result in a performance penalty if the image has a complex rendering chain
// or accelerated rendering hardware is not available. Rendering uses the
// current graphics context in the thread from which this method is called; to
// ensure consistent results across multiple threads, set the current context
// using the [NSGraphicsContext] class before calling this method.
// 
// By the time this method returns, the resultant [NSBitmapImageRep] object
// can have its raw pixel data inspected, can be put on the pasteboard, and
// can be encoded to any of the standard image formats that [NSBitmapImageRep]
// supports (JPEG, TIFF, and so on.)
// 
// If you pass in a [CIImage] object whose extents are not finite, this method
// raises an exception.
//
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(ciImage:)
func (b NSBitmapImageRep) InitWithCIImage(ciImage coreimage.CIImage) NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initWithCIImage:"), ciImage)
	return rv
}
// Initializes a newly allocated bitmap image representation from the
// specified data.
//
// data: A data object containing image data. The contents of `bitmapData` can be
// any supported bitmap format. For TIFF data, the [NSBitmapImageRep] is
// initialized from the first header and image data found in `bitmapData`.
//
// # Return Value
// 
// Returns an initialized [NSBitmapImageRep] if the initialization was
// successful or `nil` if it was unable to interpret the contents of
// `bitmapData`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(data:)
func (b NSBitmapImageRep) InitWithData(data foundation.INSData) NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initWithData:"), data)
	return rv
}
// Initializes a newly allocated bitmap image representation for incremental
// loading.
//
// # Discussion
// 
// The receiver returns itself after setting its size and data buffer to zero.
// You can then call [IncrementalLoadFromDataComplete] to incrementally add
// image data.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(forIncrementalLoad:)
func (b NSBitmapImageRep) InitForIncrementalLoad() NSBitmapImageRep {
	rv := objc.Send[NSBitmapImageRep](b.ID, objc.Sel("initForIncrementalLoad"))
	return rv
}
// Returns by indirection bitmap data of the bitmap image representation
// separated into planes.
//
// data: On return, a C array of five character pointers. If the bitmap data is in
// planar configuration, each pointer will be initialized to point to one of
// the data planes. If there are less than five planes, the remaining pointers
// will be set to [NULL]. If the bitmap data is in meshed configuration, only
// the first pointer will be initialized; the others will be [NULL].
//
// # Discussion
// 
// Color components in planar configuration are arranged in the expected
// order—for example, red before green before blue for RGB color. All color
// planes precede the coverage plane. For bitmaps whose [BitmapFormat] mask
// does not include [BitmapFormatAlphaNonpremultiplied], if a coverage plane
// exists, the bitmap’s color components are premultiplied with it. In this
// case, if you modify the contents of the bitmap, you are responsible for
// premultiplying the data.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getBitmapDataPlanes(_:)
func (b NSBitmapImageRep) GetBitmapDataPlanes(data unsafe.Pointer) {
	objc.Send[objc.ID](b.ID, objc.Sel("getBitmapDataPlanes:"), data)
}
// Returns a TIFF representation of the image using the specified compression.
//
// comp: An enum constant that represents a TIFF data-compression scheme. Legal
// values for `compression` can be found in
// [NSBitmapImageRep.TIFFCompression].
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// factor: A `float` value that provides a hint for those compression types that
// implement variable compression ratios.
// 
// Currently only JPEG compression uses a compression factor. JPEG compression
// in TIFF files is not supported, and `factor` is ignored.
//
// # Discussion
// 
// If the compression type isn’t supported for writing TIFF data (for
// example, [NSBitmapImageRep.TIFFCompression.next]), the stored compression
// is changed to [NSBitmapImageRep.TIFFCompression.none] before the TIFF
// representation is generated.
// 
// If a problem is encountered during generation of the TIFF,
// [TIFFRepresentationUsingCompressionFactor] raises an [NSTIFFException] or
// an [NSBadBitmapParametersException].
//
// [NSBadBitmapParametersException]: https://developer.apple.com/documentation/AppKit/NSBadBitmapParametersException
// [NSBitmapImageRep.TIFFCompression.next]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/next
// [NSBitmapImageRep.TIFFCompression.none]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/none
// [NSTIFFException]: https://developer.apple.com/documentation/AppKit/NSTIFFException
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation(using:factor:)
func (b NSBitmapImageRep) TIFFRepresentationUsingCompressionFactor(comp NSTIFFCompression, factor float32) foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return foundation.NSDataFromID(rv)
}
// Formats the bitmap representation’s image data using the specified
// storage type and properties and returns it in a data object.
//
// storageType: A constant that specifies a file type for bitmap images. It can be
// [NSBMPFileType], [NSGIFFileType], [NSJPEGFileType], [NSPNGFileType], or
// [NSTIFFFileType].
// //
// [NSBMPFileType]: https://developer.apple.com/documentation/AppKit/NSBMPFileType
// [NSGIFFileType]: https://developer.apple.com/documentation/AppKit/NSGIFFileType
// [NSJPEGFileType]: https://developer.apple.com/documentation/AppKit/NSJPEGFileType
// [NSPNGFileType]: https://developer.apple.com/documentation/AppKit/NSPNGFileType
// [NSTIFFFileType]: https://developer.apple.com/documentation/AppKit/NSTIFFFileType
//
// properties: A dictionary that contains key-value pairs specifying image properties.
// These string constants used as keys and the valid values are described in
// [NSBitmapImageRepPropertyKey].
//
// # Return Value
// 
// A data object containing the receiver’s image data in the specified
// format. You can write this data to a file or use it to create a new
// [NSBitmapImageRep] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/representation(using:properties:)
func (b NSBitmapImageRep) RepresentationUsingTypeProperties(storageType NSBitmapImageFileType, properties foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("representationUsingType:properties:"), storageType, properties)
	return foundation.NSDataFromID(rv)
}
// Tests whether the bitmap image representation can be compressed by the
// specified compression scheme.
//
// compression: A TIFF compression type. For more information, see the constants in
// [NSBitmapImageRep.TIFFCompression].
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// # Return Value
// 
// [true] if the receiver’s data matches `compression` with this type,
// [false] if the data doesn’t match `compression` or if `compression` is
// unsupported.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns [true] if the receiver’s data matches `compression`;
// for example, if `compression` is
// [NSBitmapImageRep.TIFFCompression.ccittfax3], then the data must be 1 bit
// per sample and 1 sample per pixel.
//
// [NSBitmapImageRep.TIFFCompression.ccittfax3]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/ccittfax3
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/canBeCompressed(using:)
func (b NSBitmapImageRep) CanBeCompressedUsing(compression NSTIFFCompression) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("canBeCompressedUsing:"), compression)
	return rv
}
// Sets the bitmap image representation’s compression type and compression
// factor.
//
// compression: An `enum` constant that identifies one of the supported compression types
// as described in [NSBitmapImageRep.TIFFCompression].
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// factor: A floating point value that is specific to the compression type. Many types
// of compression don’t support varying degrees of compression and thus
// ignore `factor`. JPEG compression allows a compression factor ranging from
// 0.0 to 1.0, with 0.0 being the lowest and 1.0 being the highest.
//
// # Discussion
// 
// When an [NSBitmapImageRep] is created, the instance stores the compression
// type and factor for the source data. The [TIFFRepresentation] property and
// [TIFFRepresentationOfImageRepsInArray] class method try to use the stored
// compression type and factor. Use this method to change the compression type
// and factor.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setCompression(_:factor:)
func (b NSBitmapImageRep) SetCompressionFactor(compression NSTIFFCompression, factor float32) {
	objc.Send[objc.ID](b.ID, objc.Sel("setCompression:factor:"), compression, factor)
}
// Returns by indirection the bitmap image representation’s compression type
// and compression factor.
//
// compression: On return, an `enum` constant that represents the compression type used on
// the data; it corresponds to one of the values returned by the class method
// [GetTIFFCompressionTypesCount].
//
// factor: A float value that is specific to the compression type. Many types of
// compression don’t support varying degrees of compression and thus ignore
// `factor`. JPEG compression allows a compression factor ranging from 0.0 to
// 1.0, with 0.0 being the lowest and 1.0 being the highest.
//
// # Discussion
// 
// Use this method to get information on the compression type for the source
// image data.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getCompression(_:factor:)
func (b NSBitmapImageRep) GetCompressionFactor(compression NSTIFFCompression, factor unsafe.Pointer) {
	objc.Send[objc.ID](b.ID, objc.Sel("getCompression:factor:"), compression, factor)
}
// Sets the specified property of the bitmap image representation to the
// specified value.
//
// property: A string constant used as a key for an image property. These properties are
// described in [NSBitmapImageRepPropertyKey].
//
// value: A value specific to `property`. If `value` is `nil`, the value of the
// property is cleared.
//
// # Discussion
// 
// The properties can affect how the image is read in and saved to file.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setProperty(_:withValue:)
func (b NSBitmapImageRep) SetPropertyWithValue(property NSBitmapImageRepPropertyKey, value objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("setProperty:withValue:"), objc.String(string(property)), value)
}
// Returns the value for the specified property.
//
// property: A string constant used as a key for an image property. These properties are
// described in [NSBitmapImageRepPropertyKey].
//
// # Return Value
// 
// A value specific to `property`, or `nil` if the property is not set for the
// bitmap.
//
// # Discussion
// 
// Image properties can affect how an image is read in and saved to file. When
// retrieving the bitmap image properties defined in
// [NSBitmapImageRepPropertyKey], be sure to check the return value of this
// method for a `nil` value. If a particular value is not set for the image,
// this method may return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/value(forProperty:)
func (b NSBitmapImageRep) ValueForProperty(property NSBitmapImageRepPropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("valueForProperty:"), objc.String(string(property)))
	return objectivec.Object{ID: rv}
}
// Loads the current image data into an incrementally-loaded image
// representation and returns the current status of the image.
//
// data: A data object that contains the image to be loaded.
//
// complete: [true] if the image is entirely downloaded, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An integer constant indicating the status of the image during the load
// operation. See the discussion for details.
//
// # Discussion
// 
// After initializing the receiver with [InitForIncrementalLoad], you should
// call this method to incrementally load the image. Call this method each
// time new data becomes available. Always pass the entire image data buffer
// in `data`, not just the newest data, because the image decompressor may
// need the original data in order to backtrack. This method will block until
// the data is decompressed; it will decompress as much of the image as
// possible based on the length of the data. The image rep does not retain
// `data`, so you must ensure that `data` is not released for the duration of
// this method call. Pass [false] for `complete` until the entire image is
// downloaded, at which time you should pass [true]. You should also pass
// [true] for `complete` if you have only partially downloaded the data, but
// cannot finish the download.
// 
// This method returns [NSBitmapImageRep.LoadStatus.unknownType] if you did
// not pass enough data to determine the image format; you should continue to
// invoke this method with additional data.
// 
// This method returns [NSBitmapImageRep.LoadStatus.readingHeader] if it has
// enough data to determine the image format, but needs more data to determine
// the size and depth and other characteristics of the image. You should
// continue to invoke this method with additional data.
// 
// This method returns [NSBitmapImageRep.LoadStatus.willNeedAllData] if the
// image format does not support incremental loading or the Application Kit
// does not yet implement incremental loading for the image format. You may
// continue to invoke this method in this case, but until you pass [true] for
// `complete`, this method will continue to return
// [NSBitmapImageRep.LoadStatus.willNeedAllData], and will perform no
// decompression. Once you pass [true], the image will be decompressed and one
// of the final three status messages will be returned.
// 
// If the image format does support incremental loading, then once enough data
// has been read, the image is decompressed from the top down a row at a time.
// In this case, instead of a status value, this method returns the number of
// pixel rows that have been decompressed, starting from the top of the image.
// You can use this information to draw the part of the image that is valid.
// The rest of the image is filled with opaque white. Note that if the image
// is progressive (as in a progressive JPEG or 2D interlaced PNG), this method
// may quickly return the full height of the image, but the image may still be
// loading, so do not use this return value as an indication of how much of
// the image remains to be decompressed.
// 
// If an error occurred while decompressing, this method returns
// [NSBitmapImageRep.LoadStatus.invalidData]. If `complete` is [true] but not
// enough data was available for decompression,
// [NSBitmapImageRep.LoadStatus.unexpectedEOF] is returned. If enough data has
// been provided (regardless of the `complete` flag), then
// [NSBitmapImageRep.LoadStatus.completed] is returned. When any of these
// three status results are returned, this method has adjusted the
// [NSBitmapImageRep] so that [PixelsHigh] and [Size], as well as the bitmap
// data, only contains the pixels that are valid, if any.
// 
// To cancel decompression, just pass in the existing data or `nil` and [true]
// for `complete`. This method stops decompression immediately, adjusts the
// image size, and returns [NSBitmapImageRep.LoadStatus.unexpectedEOF]. This
// method returns [NSBitmapImageRep.LoadStatus.completed] if you call it after
// receiving any error results ([NSBitmapImageRep.LoadStatus.invalidData] or
// [NSBitmapImageRep.LoadStatus.unexpectedEOF]) or if you call it on an
// [NSBitmapImageRep] that was not initialized with [InitForIncrementalLoad].
//
// [NSBitmapImageRep.LoadStatus.completed]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/completed
// [NSBitmapImageRep.LoadStatus.invalidData]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/invalidData
// [NSBitmapImageRep.LoadStatus.readingHeader]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/readingHeader
// [NSBitmapImageRep.LoadStatus.unexpectedEOF]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/unexpectedEOF
// [NSBitmapImageRep.LoadStatus.unknownType]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/unknownType
// [NSBitmapImageRep.LoadStatus.willNeedAllData]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/willNeedAllData
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/incrementalLoad(from:complete:)
func (b NSBitmapImageRep) IncrementalLoadFromDataComplete(data foundation.INSData, complete bool) int {
	rv := objc.Send[int](b.ID, objc.Sel("incrementalLoadFromData:complete:"), data, complete)
	return rv
}
// Changes the color of the pixel at the specified coordinates.
//
// color: A color object representing the color to be set.
//
// x: The x-axis coordinate of the pixel.
//
// y: The y-axis coordinate of the pixel.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setColor(_:atX:y:)
func (b NSBitmapImageRep) SetColorAtXY(color INSColor, x int, y int) {
	objc.Send[objc.ID](b.ID, objc.Sel("setColor:atX:y:"), color, x, y)
}
// Returns the color of the pixel at the specified coordinates.
//
// x: The x-axis coordinate.
//
// y: The y-axis coordinate.
//
// # Return Value
// 
// A color object representing the color at the specified coordinates.
//
// # Discussion
// 
// Calling this method creates a new [NSColor] object. The overhead of object
// creation means this method is best suited for infrequent color sampling. If
// you instead need to work with large numbers of pixels, access the bitmap
// data directly using the [BitmapData] property or the [GetPixelAtXY] method
// for better performance.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorAt(x:y:)
func (b NSBitmapImageRep) ColorAtXY(x int, y int) INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("colorAtX:y:"), x, y)
	return NSColorFromID(rv)
}
// Sets the bitmap image representation’s pixel at the specified coordinates
// to the specified raw pixel values.
//
// p: An array of integers representing the raw pixel values. The values must be
// in an order appropriate to the object’s bitmap format. Small pixel sample
// values should be passed as an integer value. Floating point values should
// be cast `int[]`.
//
// x: The x-axis coordinate of the pixel.
//
// y: The y-axis coordinate of the pixel.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setPixel(_:atX:y:)
func (b NSBitmapImageRep) SetPixelAtXY(p uint, x int, y int) {
	objc.Send[objc.ID](b.ID, objc.Sel("setPixel:atX:y:"), p, x, y)
}
// Returns by indirection the pixel data for the specified location in the
// bitmap image representation.
//
// p: On return, an array of integers containing raw pixel data in the
// appropriate order according to the object’s bitmap format. Smaller
// integer samples, such as 4-bit, are returned as an integer. Floating point
// values are cast to integer values and returned.
//
// x: The x-axis coordinate of the pixel.
//
// y: The y-axis coordinate of the pixel.
//
// # Discussion
// 
// The origin is in the top-left corner.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getPixel(_:atX:y:)
func (b NSBitmapImageRep) GetPixelAtXY(p uint, x int, y int) {
	objc.Send[objc.ID](b.ID, objc.Sel("getPixel:atX:y:"), p, x, y)
}
// Converts the bitmap image representation to the specified color space.
//
// targetSpace: The new color space.
//
// renderingIntent: The rendering intent specifies how to handle colors that are not located
// within the target color space. The supported values are
// [NSColorRenderingIntent].
// //
// [NSColorRenderingIntent]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
//
// # Return Value
// 
// An [NSBitmapImageRep], or `nil`, if the conversion fails. If the original
// [NSBitmapImageRep] already uses that color space, it is returned as is.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/converting(to:renderingIntent:)
func (b NSBitmapImageRep) BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace INSColorSpace, renderingIntent NSColorRenderingIntent) INSBitmapImageRep {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bitmapImageRepByConvertingToColorSpace:renderingIntent:"), targetSpace, renderingIntent)
	return NSBitmapImageRepFromID(rv)
}
// Changes the color space tag of the bitmap image representation.
//
// newSpace: The desired color space.
//
// # Return Value
// 
// An [NSBitmapImageRep], or `nil`, if the conversion fails. If the original
// [NSBitmapImageRep] already uses that color space, it is returned as is.
//
// # Discussion
// 
// This method will definitely fail if you pass a color space that has a
// different color space model than the receiver. That is, if your original
// image is sRGB, you can only retag with some other RGB colorspace.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/retagging(with:)
func (b NSBitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace INSColorSpace) INSBitmapImageRep {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bitmapImageRepByRetaggingWithColorSpace:"), newSpace)
	return NSBitmapImageRepFromID(rv)
}

// Creates and returns an array of bitmap image representation objects that
// correspond to the images in the specified data.
//
// data: A data object containing one or more bitmapped images or `nil` if the class
// is unable to create an image representation. The `bitmapData` parameter can
// contain data in any supported bitmap format.
//
// # Return Value
// 
// An array of [NSBitmapImageRep] instances or an empty array if the class is
// unable to create any image representations.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/imageReps(with:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) ImageRepsWithData(data foundation.INSData) []NSImageRep {
	rv := objc.Send[[]objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("imageRepsWithData:"), data)
	return objc.ConvertSlice(rv, func(id objc.ID) NSImageRep {
		return NSImageRepFromID(id)
	})
}
// Returns a TIFF representation of the specified images.
//
// array: An array containing objects representing bitmap image representations.
//
// # Return Value
// 
// A data object containing a TIFF image representation.
//
// # Discussion
// 
// This method uses the compression returned by [GetCompressionFactor] (if
// applicable). If a problem is encountered during generation of the TIFF,
// this method raises an [NSTIFFException] or an
// [NSBadBitmapParametersException].
//
// [NSBadBitmapParametersException]: https://developer.apple.com/documentation/AppKit/NSBadBitmapParametersException
// [NSTIFFException]: https://developer.apple.com/documentation/AppKit/NSTIFFException
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentationOfImageReps(in:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []NSImageRep) foundation.NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("TIFFRepresentationOfImageRepsInArray:"), objectivec.IObjectSliceToNSArray(array))
	return foundation.NSDataFromID(rv)
}
// Returns a TIFF representation of the specified images using the specified
// compression scheme and factor.
//
// array: An array containing objects representing bitmap image representations.
//
// comp: An enum constant that represents a TIFF data-compression scheme. Legal
// values for `compression` can be found in
// [NSBitmapImageRep.TIFFCompression].
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// factor: A `float` value that provides a hint for those compression types that
// implement variable compression ratios.
// 
// Currently only JPEG compression uses a compression factor. JPEG compression
// in TIFF files is not supported, and `factor` is ignored.
//
// # Return Value
// 
// A data object containing a TIFF image representation.
//
// # Discussion
// 
// If the specified compression isn’t applicable, no compression is used. If
// a problem is encountered during generation of the TIFF, the method raises
// an [NSTIFFException] or an [NSBadBitmapParametersException].
//
// [NSBadBitmapParametersException]: https://developer.apple.com/documentation/AppKit/NSBadBitmapParametersException
// [NSTIFFException]: https://developer.apple.com/documentation/AppKit/NSTIFFException
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentationOfImageReps(in:using:factor:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) TIFFRepresentationOfImageRepsInArrayUsingCompressionFactor(array []NSImageRep, comp NSTIFFCompression, factor float32) foundation.NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("TIFFRepresentationOfImageRepsInArray:usingCompression:factor:"), objectivec.IObjectSliceToNSArray(array), comp, factor)
	return foundation.NSDataFromID(rv)
}
// Formats the specified bitmap images using the specified storage type and
// properties and returns them in a data object.
//
// imageReps: An array of [NSBitmapImageRep] objects.
//
// storageType: An [NSBitmapImageRep.FileType] constant specifying a file type for bitmap
// images.
// //
// [NSBitmapImageRep.FileType]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType
//
// properties: A dictionary that contains key-value pairs specifying image properties.
// These string constants used as keys and the valid values are described in
// [NSBitmapImageRepPropertyKey].
//
// # Return Value
// 
// A data object containing the bitmap image data in the specified format. You
// can write this data to a file or use it to create a new [NSBitmapImageRep]
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/representationOfImageReps(in:using:properties:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []NSImageRep, storageType NSBitmapImageFileType, properties foundation.INSDictionary) foundation.NSData {
	rv := objc.Send[objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("representationOfImageRepsInArray:usingType:properties:"), objectivec.IObjectSliceToNSArray(imageReps), storageType, properties)
	return foundation.NSDataFromID(rv)
}
// Returns by indirection an array of all available compression types that can
// be used when writing a TIFF image.
//
// list: On return, a C array of [NSBitmapImageRep.TIFFCompression] constants. This
// array belongs to the [NSBitmapImageRep] class; it shouldn’t be freed or
// altered. See [NSBitmapImageRep.TIFFCompression] for the supported TIFF
// compression types.
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// numTypes: The number of constants in list.
//
// # Discussion
// 
// Note that not all compression types can be used for all images:
// [NSBitmapImageRep.TIFFCompression.next] can be used only to retrieve image
// data. Because future releases may include other compression types, always
// use this method to get the available compression types—for example, when
// you implement a user interface for selecting compression types.
//
// [NSBitmapImageRep.TIFFCompression.next]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/next
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getTIFFCompressionTypes(_:count:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) GetTIFFCompressionTypesCount(list []NSTIFFCompression, numTypes unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("getTIFFCompressionTypes:count:"), objc.CArray(list), numTypes)
}
// Returns an autoreleased string containing the localized name for the
// specified compression type.
//
// compression: A TIFF compression type. For more information, see the constants in
// [NSBitmapImageRep.TIFFCompression].
// //
// [NSBitmapImageRep.TIFFCompression]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
//
// # Return Value
// 
// A localized name for `compression` or `nil` if `compression` is
// unrecognized.
//
// # Discussion
// 
// When implementing a user interface for selecting TIFF compression types,
// use [GetTIFFCompressionTypesCount] to get the list of supported compression
// types, then use this method to get the localized names for each compression
// type.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/localizedName(forTIFFCompressionType:)
func (_NSBitmapImageRepClass NSBitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression NSTIFFCompression) string {
	rv := objc.Send[objc.ID](objc.ID(_NSBitmapImageRepClass.class), objc.Sel("localizedNameForTIFFCompressionType:"), compression)
	return foundation.NSStringFromID(rv).String()
}

// The format of the bitmap image representation.
//
// # Discussion
// 
// Returns 0 by default. The return value can indicate several different
// attributes, which are described in [NSBitmapImageRep.Format].
//
// [NSBitmapImageRep.Format]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitmapFormat
func (b NSBitmapImageRep) BitmapFormat() NSBitmapFormat {
	rv := objc.Send[NSBitmapFormat](b.ID, objc.Sel("bitmapFormat"))
	return NSBitmapFormat(rv)
}
// The number of bits allocated for each pixel in each plane of data.
//
// # Discussion
// 
// This number is normally equal to the number of bits per sample or, if the
// data is in meshed configuration, the number of bits per sample times the
// number of samples per pixel. It can be explicitly set to another value (in
// [InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel])
// in case extra memory is allocated for each pixel. This may be the case, for
// example, if pixel data is aligned on byte boundaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitsPerPixel
func (b NSBitmapImageRep) BitsPerPixel() int {
	rv := objc.Send[int](b.ID, objc.Sel("bitsPerPixel"))
	return rv
}
// The number of bytes in each plane or channel of data.
//
// # Discussion
// 
// This number is calculated from the number of bytes per row and the height
// of the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bytesPerPlane
func (b NSBitmapImageRep) BytesPerPlane() int {
	rv := objc.Send[int](b.ID, objc.Sel("bytesPerPlane"))
	return rv
}
// The minimum number of bytes required to specify a scan line in each data
// plane.
//
// # Discussion
// 
// A scan line is a single row of pixels spanning the width of the image. If
// not explicitly set to another value (in
// [InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel]),
// this number will be figured from the width of the image, the number of bits
// per sample, and, if the data is in a meshed configuration, the number of
// samples per pixel. It can be set to another value to indicate that each row
// of data is aligned on word or other boundaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bytesPerRow
func (b NSBitmapImageRep) BytesPerRow() int {
	rv := objc.Send[int](b.ID, objc.Sel("bytesPerRow"))
	return rv
}
// A Boolean value that indicates whether the image data is in a planar
// configuration.
//
// # Discussion
// 
// The value of this property is [true] if the data is in a planar
// configuration or [false] if it is in a meshed configuration. In a planar
// configuration, the image data is segregated into a separate plane for each
// color and coverage component. In a meshed configuration, the data is
// integrated into a single plane.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/isPlanar
func (b NSBitmapImageRep) Planar() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isPlanar"))
	return rv
}
// The number of separate planes into which the image data is organized.
//
// # Discussion
// 
// If the data has a separate plane for each component—that is, [Planar] is
// [true]—the value of this property is the number of samples per pixel. If
// the data is meshed, the value of this property is `1`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/numberOfPlanes
func (b NSBitmapImageRep) NumberOfPlanes() int {
	rv := objc.Send[int](b.ID, objc.Sel("numberOfPlanes"))
	return rv
}
// The number of components for each pixel.
//
// # Discussion
// 
// This property reflects both the number of color components and the coverage
// component, if present.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/samplesPerPixel
func (b NSBitmapImageRep) SamplesPerPixel() int {
	rv := objc.Send[int](b.ID, objc.Sel("samplesPerPixel"))
	return rv
}
// A pointer to the bitmap data.
//
// # Discussion
// 
// For planar data, the value in this property points to the first plane.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitmapData
func (b NSBitmapImageRep) BitmapData() string {
	rv := objc.Send[*byte](b.ID, objc.Sel("bitmapData"))
	return objc.GoString(rv)
}
// A TIFF representation of the bitmap image data.
//
// # Discussion
// 
// Accessing this property results in a call to the
// [TIFFRepresentationUsingCompressionFactor] method using the stored
// compression type and factor retrieved from the initial image data or
// changed using the [SetCompressionFactor] method. If the stored compression
// type isn’t supported for writing TIFF data (for example,
// [NSBitmapImageRep.TIFFCompression.next]), the stored compression is changed
// to [NSBitmapImageRep.TIFFCompression.none] before calling the
// [TIFFRepresentationUsingCompressionFactor] method using the compression
// that’s returned by [GetCompressionFactor] (if applicable).
// 
// If a problem is encountered during generation of the TIFF, an
// [NSTIFFException] or an [NSBadBitmapParametersException] is raised.
//
// [NSBadBitmapParametersException]: https://developer.apple.com/documentation/AppKit/NSBadBitmapParametersException
// [NSBitmapImageRep.TIFFCompression.next]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/next
// [NSBitmapImageRep.TIFFCompression.none]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/none
// [NSTIFFException]: https://developer.apple.com/documentation/AppKit/NSTIFFException
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation
func (b NSBitmapImageRep) TIFFRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("TIFFRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A Core Graphics image object based on the bitmap image representation’s
// data.
//
// # Discussion
// 
// The autoreleased [CGImage] opaque type in this property has pixel
// dimensions that are identical to those of the bitmap image rep object. If
// an existing [CGImage] opaque type is not available, accessing this property
// creates a new one. If you change the bitmap image rep’s contents later,
// accessing this property again might return a different [CGImage] opaque
// type.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/cgImage
func (b NSBitmapImageRep) CGImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](b.ID, objc.Sel("CGImage"))
	return coregraphics.CGImageRef(rv)
}
// The color space of the bitmap.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorSpace
func (b NSBitmapImageRep) ColorSpace() INSColorSpace {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("colorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}

