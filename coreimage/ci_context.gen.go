// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIContext] class.
var (
	_CIContextClass     CIContextClass
	_CIContextClassOnce sync.Once
)

func getCIContextClass() CIContextClass {
	_CIContextClassOnce.Do(func() {
		_CIContextClass = CIContextClass{class: objc.GetClass("CIContext")}
	})
	return _CIContextClass
}

// GetCIContextClass returns the class object for CIContext.
func GetCIContextClass() CIContextClass {
	return getCIContextClass()
}

type CIContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIContextClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIContextClass) Alloc() CIContext {
	rv := objc.Send[CIContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The Core Image context class provides an evaluation context for Core Image
// processing with Metal, OpenGL, or OpenCL.
//
// # Overview
//
// You use a [CIContext] instance to render a [CIImage] instance which
// represents a graph of image processing operations which are built using
// other Core Image classes, such as [CIFilter], [CIKernel], [CIColor] and
// [CIImage]. You can also use a [CIContext] with the [CIDetector] class to
// analyze images — for example, to detect faces or barcodes.
//
// Contexts support automatic color management by performing all processing
// operations in a working color space. This means that unless told otherwise:
//
// - All input images are color matched from the input’s color space to the
// working space. - All renders are color matched from the working space to
// the destination space. (For more information on [CGColorSpace] see
// [CGColorSpace])
//
// [CIContext] and [CIImage] instances are immutable, so multiple threads can
// use the same [CIContext] instance to render [CIImage] instances. However,
// [CIFilter] instances are mutable and thus cannot be shared safely among
// threads. Each thread must take case not to access or modify a [CIFilter]
// instance while it is being used by another thread.
//
// The [CIContext] manages various internal state such as [MTLCommandQueue]
// and caches for compiled kernels and intermediate buffers. For this reason
// it is not recommended to create many [CIContext] instances. As a rule, it
// recommended that you create one [CIContext] instance for each view that
// renders [CIImage] or each background task.
//
// # Rendering Images
//
//   - [CIContext.CreateCGImageFromRect]: Creates a Core Graphics image from a region of a Core Image image instance.
//   - [CIContext.CreateCGImageFromRectFormatColorSpace]: Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling the pixel format and color space of the [CGImage].
//   - [CIContext.CreateCGImageFromRectFormatColorSpaceDeferred]: Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling when the image is rendered.
//   - [CIContext.RenderToBitmapRowBytesBoundsFormatColorSpace]: Renders to the given bitmap.
//   - [CIContext.RenderToCVPixelBuffer]: Renders an image into a pixel buffer.
//   - [CIContext.RenderToCVPixelBufferBoundsColorSpace]: Renders a region of an image into a pixel buffer.
//   - [CIContext.RenderToIOSurfaceBoundsColorSpace]: Renders a region of an image into an IOSurface object.
//   - [CIContext.RenderToMTLTextureCommandBufferBoundsColorSpace]: Renders a region of an image to a Metal texture.
//
// # Drawing Images
//
//   - [CIContext.DrawImageInRectFromRect]: Renders a region of an image to a rectangle in the context destination.
//
// # Managing Resources
//
//   - [CIContext.ClearCaches]: Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
//   - [CIContext.ReclaimResources]: Runs the garbage collector to reclaim any resources that the context no longer requires.
//   - [CIContext.WorkingColorSpace]: The working color space of the Core Image context.
//   - [CIContext.WorkingFormat]: The working pixel format of the Core Image context.
//
// # Rendering Images for Data or File Export
//
//   - [CIContext.TIFFRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in TIFF format.
//   - [CIContext.JPEGRepresentationOfImageColorSpaceOptions]: Renders the image and exports the resulting image data in JPEG format.
//   - [CIContext.PNGRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in PNG format.
//   - [CIContext.HEIFRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in HEIF format.
//   - [CIContext.HEIF10RepresentationOfImageColorSpaceOptionsError]: Renders the image and exports the resulting image data in HEIF10 format.
//   - [CIContext.OpenEXRRepresentationOfImageOptionsError]: Renders the image and exports the resulting image data in open EXR format.
//   - [CIContext.WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in TIFF format.
//   - [CIContext.WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in JPEG format.
//   - [CIContext.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in PNG format.
//   - [CIContext.WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in HEIF format.
//   - [CIContext.WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in HEIF10 format.
//   - [CIContext.WriteOpenEXRRepresentationOfImageToURLOptionsError]: Renders the image and exports the resulting image data as a file in open EXR format.
//
// # Creating Depth Blur Filters
//
//   - [CIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [CIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [CIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [CIContext.DepthBlurEffectFilterForImageDataOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect.
//   - [CIContext.DepthBlurEffectFilterForImageURLOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image URL that can be used to apply a depth blur effect.
//
// # Customizing Render Destination
//
//   - [CIContext.PrepareRenderFromRectToDestinationAtPointError]: An optional call to warm up a [CIContext](<doc://com.apple.coreimage/documentation/CoreImage/CIContext>) so that subsequent calls to render with the same arguments run more efficiently.
//   - [CIContext.StartTaskToClearError]: Fills the entire destination with black or clear depending on its [alphaMode](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderDestination/alphaMode>).
//   - [CIContext.StartTaskToRenderFromRectToDestinationAtPointError]: Renders a portion of an image to a point in the destination.
//   - [CIContext.StartTaskToRenderToDestinationError]: Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination.
//
// # Initializers
//
//   - [CIContext.InitWithOptions]: Initializes a context without a specific rendering destination, using the specified options.
//
// # Instance Methods
//
//   - [CIContext.CalculateHDRStatsForCGImage]: Given a Core Graphics image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Graphics image that has the calculated values.
//   - [CIContext.CalculateHDRStatsForIOSurface]: Given an IOSurface, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the surface’s attachments to store the values.
//   - [CIContext.CalculateHDRStatsForCVPixelBuffer]: Given a CVPixelBuffer, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the buffers’s attachments to store the values.
//   - [CIContext.CalculateHDRStatsForImage]: Given a Core Image image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Image image that has the calculated values.
//   - [CIContext.CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats]: Creates a Core Graphics image from a region of a Core Image image instance with an option for calculating HDR statistics.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
type CIContext struct {
	objectivec.Object
}

// CIContextFromID constructs a [CIContext] from an objc.ID.
//
// The Core Image context class provides an evaluation context for Core Image
// processing with Metal, OpenGL, or OpenCL.
func CIContextFromID(id objc.ID) CIContext {
	return CIContext{objectivec.Object{ID: id}}
}

// NOTE: CIContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIContext] class.
//
// # Rendering Images
//
//   - [ICIContext.CreateCGImageFromRect]: Creates a Core Graphics image from a region of a Core Image image instance.
//   - [ICIContext.CreateCGImageFromRectFormatColorSpace]: Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling the pixel format and color space of the [CGImage].
//   - [ICIContext.CreateCGImageFromRectFormatColorSpaceDeferred]: Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling when the image is rendered.
//   - [ICIContext.RenderToBitmapRowBytesBoundsFormatColorSpace]: Renders to the given bitmap.
//   - [ICIContext.RenderToCVPixelBuffer]: Renders an image into a pixel buffer.
//   - [ICIContext.RenderToCVPixelBufferBoundsColorSpace]: Renders a region of an image into a pixel buffer.
//   - [ICIContext.RenderToIOSurfaceBoundsColorSpace]: Renders a region of an image into an IOSurface object.
//   - [ICIContext.RenderToMTLTextureCommandBufferBoundsColorSpace]: Renders a region of an image to a Metal texture.
//
// # Drawing Images
//
//   - [ICIContext.DrawImageInRectFromRect]: Renders a region of an image to a rectangle in the context destination.
//
// # Managing Resources
//
//   - [ICIContext.ClearCaches]: Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
//   - [ICIContext.ReclaimResources]: Runs the garbage collector to reclaim any resources that the context no longer requires.
//   - [ICIContext.WorkingColorSpace]: The working color space of the Core Image context.
//   - [ICIContext.WorkingFormat]: The working pixel format of the Core Image context.
//
// # Rendering Images for Data or File Export
//
//   - [ICIContext.TIFFRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in TIFF format.
//   - [ICIContext.JPEGRepresentationOfImageColorSpaceOptions]: Renders the image and exports the resulting image data in JPEG format.
//   - [ICIContext.PNGRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in PNG format.
//   - [ICIContext.HEIFRepresentationOfImageFormatColorSpaceOptions]: Renders the image and exports the resulting image data in HEIF format.
//   - [ICIContext.HEIF10RepresentationOfImageColorSpaceOptionsError]: Renders the image and exports the resulting image data in HEIF10 format.
//   - [ICIContext.OpenEXRRepresentationOfImageOptionsError]: Renders the image and exports the resulting image data in open EXR format.
//   - [ICIContext.WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in TIFF format.
//   - [ICIContext.WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in JPEG format.
//   - [ICIContext.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in PNG format.
//   - [ICIContext.WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in HEIF format.
//   - [ICIContext.WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError]: Renders the image and exports the resulting image data as a file in HEIF10 format.
//   - [ICIContext.WriteOpenEXRRepresentationOfImageToURLOptionsError]: Renders the image and exports the resulting image data as a file in open EXR format.
//
// # Creating Depth Blur Filters
//
//   - [ICIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [ICIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [ICIContext.DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
//   - [ICIContext.DepthBlurEffectFilterForImageDataOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect.
//   - [ICIContext.DepthBlurEffectFilterForImageURLOptions]: Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image URL that can be used to apply a depth blur effect.
//
// # Customizing Render Destination
//
//   - [ICIContext.PrepareRenderFromRectToDestinationAtPointError]: An optional call to warm up a [CIContext](<doc://com.apple.coreimage/documentation/CoreImage/CIContext>) so that subsequent calls to render with the same arguments run more efficiently.
//   - [ICIContext.StartTaskToClearError]: Fills the entire destination with black or clear depending on its [alphaMode](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderDestination/alphaMode>).
//   - [ICIContext.StartTaskToRenderFromRectToDestinationAtPointError]: Renders a portion of an image to a point in the destination.
//   - [ICIContext.StartTaskToRenderToDestinationError]: Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination.
//
// # Initializers
//
//   - [ICIContext.InitWithOptions]: Initializes a context without a specific rendering destination, using the specified options.
//
// # Instance Methods
//
//   - [ICIContext.CalculateHDRStatsForCGImage]: Given a Core Graphics image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Graphics image that has the calculated values.
//   - [ICIContext.CalculateHDRStatsForIOSurface]: Given an IOSurface, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the surface’s attachments to store the values.
//   - [ICIContext.CalculateHDRStatsForCVPixelBuffer]: Given a CVPixelBuffer, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the buffers’s attachments to store the values.
//   - [ICIContext.CalculateHDRStatsForImage]: Given a Core Image image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Image image that has the calculated values.
//   - [ICIContext.CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats]: Creates a Core Graphics image from a region of a Core Image image instance with an option for calculating HDR statistics.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext
type ICIContext interface {
	objectivec.IObject

	// Topic: Rendering Images

	// Creates a Core Graphics image from a region of a Core Image image instance.
	CreateCGImageFromRect(image ICIImage, fromRect corefoundation.CGRect) coregraphics.CGImageRef
	// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling the pixel format and color space of the [CGImage].
	CreateCGImageFromRectFormatColorSpace(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) coregraphics.CGImageRef
	// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling when the image is rendered.
	CreateCGImageFromRectFormatColorSpaceDeferred(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef, deferred bool) coregraphics.CGImageRef
	// Renders to the given bitmap.
	RenderToBitmapRowBytesBoundsFormatColorSpace(image ICIImage, data unsafe.Pointer, rowBytes int, bounds corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef)
	// Renders an image into a pixel buffer.
	RenderToCVPixelBuffer(image ICIImage, buffer corevideo.CVImageBufferRef)
	// Renders a region of an image into a pixel buffer.
	RenderToCVPixelBufferBoundsColorSpace(image ICIImage, buffer corevideo.CVImageBufferRef, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef)
	// Renders a region of an image into an IOSurface object.
	RenderToIOSurfaceBoundsColorSpace(image ICIImage, surface iosurface.IOSurfaceRef, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef)
	// Renders a region of an image to a Metal texture.
	RenderToMTLTextureCommandBufferBoundsColorSpace(image ICIImage, texture metal.MTLTexture, commandBuffer metal.MTLCommandBuffer, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef)

	// Topic: Drawing Images

	// Renders a region of an image to a rectangle in the context destination.
	DrawImageInRectFromRect(image ICIImage, inRect corefoundation.CGRect, fromRect corefoundation.CGRect)

	// Topic: Managing Resources

	// Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
	ClearCaches()
	// Runs the garbage collector to reclaim any resources that the context no longer requires.
	ReclaimResources()
	// The working color space of the Core Image context.
	WorkingColorSpace() coregraphics.CGColorSpaceRef
	// The working pixel format of the Core Image context.
	WorkingFormat() CIFormat

	// Topic: Rendering Images for Data or File Export

	// Renders the image and exports the resulting image data in TIFF format.
	TIFFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData
	// Renders the image and exports the resulting image data in JPEG format.
	JPEGRepresentationOfImageColorSpaceOptions(image ICIImage, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData
	// Renders the image and exports the resulting image data in PNG format.
	PNGRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData
	// Renders the image and exports the resulting image data in HEIF format.
	HEIFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData
	// Renders the image and exports the resulting image data in HEIF10 format.
	HEIF10RepresentationOfImageColorSpaceOptionsError(image ICIImage, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (foundation.INSData, error)
	// Renders the image and exports the resulting image data in open EXR format.
	OpenEXRRepresentationOfImageOptionsError(image ICIImage, options foundation.INSDictionary) (foundation.INSData, error)
	// Renders the image and exports the resulting image data as a file in TIFF format.
	WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error)
	// Renders the image and exports the resulting image data as a file in JPEG format.
	WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url foundation.INSURL, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error)
	// Renders the image and exports the resulting image data as a file in PNG format.
	WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error)
	// Renders the image and exports the resulting image data as a file in HEIF format.
	WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error)
	// Renders the image and exports the resulting image data as a file in HEIF10 format.
	WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url foundation.INSURL, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error)
	// Renders the image and exports the resulting image data as a file in open EXR format.
	WriteOpenEXRRepresentationOfImageToURLOptionsError(image ICIImage, url foundation.INSURL, options foundation.INSDictionary) (bool, error)

	// Topic: Creating Depth Blur Filters

	// Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, glassesMatte ICIImage, gainMap ICIImage, orientation uint, options foundation.INSDictionary) CIFilter
	// Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, orientation uint, options foundation.INSDictionary) CIFilter
	// Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect created with the supplied auxiliary images.
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, orientation uint, options foundation.INSDictionary) CIFilter
	// Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image data that can be used to apply a depth blur effect.
	DepthBlurEffectFilterForImageDataOptions(data foundation.INSData, options foundation.INSDictionary) CIFilter
	// Create a [CIFilter](<doc://com.apple.coreimage/documentation/CoreImage/CIFilter-swift.class>) instance for the supplied image URL that can be used to apply a depth blur effect.
	DepthBlurEffectFilterForImageURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIFilter

	// Topic: Customizing Render Destination

	// An optional call to warm up a [CIContext](<doc://com.apple.coreimage/documentation/CoreImage/CIContext>) so that subsequent calls to render with the same arguments run more efficiently.
	PrepareRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint) (bool, error)
	// Fills the entire destination with black or clear depending on its [alphaMode](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderDestination/alphaMode>).
	StartTaskToClearError(destination ICIRenderDestination) (ICIRenderTask, error)
	// Renders a portion of an image to a point in the destination.
	StartTaskToRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint) (ICIRenderTask, error)
	// Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination.
	StartTaskToRenderToDestinationError(image ICIImage, destination ICIRenderDestination) (ICIRenderTask, error)

	// Topic: Initializers

	// Initializes a context without a specific rendering destination, using the specified options.
	InitWithOptions(options foundation.INSDictionary) CIContext

	// Topic: Instance Methods

	// Given a Core Graphics image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Graphics image that has the calculated values.
	CalculateHDRStatsForCGImage(cgimage coregraphics.CGImageRef) coregraphics.CGImageRef
	// Given an IOSurface, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the surface’s attachments to store the values.
	CalculateHDRStatsForIOSurface(surface iosurface.IOSurfaceRef)
	// Given a CVPixelBuffer, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the buffers’s attachments to store the values.
	CalculateHDRStatsForCVPixelBuffer(buffer corevideo.CVImageBufferRef)
	// Given a Core Image image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Image image that has the calculated values.
	CalculateHDRStatsForImage(image ICIImage) ICIImage
	// Creates a Core Graphics image from a region of a Core Image image instance with an option for calculating HDR statistics.
	CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef, deferred bool, calculateHDRStats bool) coregraphics.CGImageRef

	// The render destination’s representation of alpha (transparency) values.
	AlphaMode() CIRenderDestinationAlphaMode
	SetAlphaMode(value CIRenderDestinationAlphaMode)
}

// Init initializes the instance.
func (c CIContext) Init() CIContext {
	rv := objc.Send[CIContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CIContext) Autorelease() CIContext {
	rv := objc.Send[CIContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIContext creates a new CIContext instance.
func NewCIContext() CIContext {
	class := getCIContextClass()
	rv := objc.Send[CIContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Core Image context from a Quartz context, using the specified
// options.
//
// cgctx: A Quartz graphics context.
//
// options: A dictionary that contains color space information. You can pass any of the
// keys defined in `Context Options` along with the appropriate value.
//
// # Discussion
//
// After calling this method, Core Image draws content to the specified Quartz
// graphics context.
//
// When you create a [CIContext] object using a Quartz graphics context, any
// transformations that are already set on the Quartz graphics context affect
// drawing to that context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(cgContext:options:)
func NewContextWithCGContextOptions(cgctx coregraphics.CGContextRef, options foundation.INSDictionary) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return CIContextFromID(rv)
}

// Creates a Core Image context from an EAGL context.
//
// eaglContext: The EAGL context to render to.
//
// # Return Value
//
// A Core Image context that targets OpenGL ES.
//
// # Discussion
//
// The OpenGL ES context must support OpenGL ES 2.0. All drawing performed
// using the methods listed in Drawing Images is rendered directly into the
// context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:)
func NewContextWithEAGLContext(eaglContext unsafe.Pointer) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithEAGLContext:"), eaglContext)
	return CIContextFromID(rv)
}

// Creates a Core Image context from an EAGL context using the specified
// options.
//
// eaglContext: The EAGL context to render to.
//
// options: A dictionary that contains options for creating a [CIContext] object. You
// can pass any of the keys defined in `Context Options` along with the
// appropriate value.
//
// # Return Value
//
// A Core Image context that targets OpenGL ES.
//
// # Discussion
//
// The OpenGL ES context must support OpenGL ES 2.0. All drawing performed
// using the methods listed in Drawing Images is rendered directly into the
// context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:options:)
func NewContextWithEAGLContextOptions(eaglContext unsafe.Pointer, options foundation.INSDictionary) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithEAGLContext:options:"), eaglContext, options)
	return CIContextFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:)
func NewContextWithMTLCommandQueue(commandQueue metal.MTLCommandQueue) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithMTLCommandQueue:"), commandQueue)
	return CIContextFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:options:)
func NewContextWithMTLCommandQueueOptions(commandQueue metal.MTLCommandQueue, options foundation.INSDictionary) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithMTLCommandQueue:options:"), commandQueue, options)
	return CIContextFromID(rv)
}

// Creates a Core Image context using the specified Metal device.
//
// device: The Metal device object to use for rendering.
//
// # Return Value
//
// A Core Image context.
//
// # Discussion
//
// Use this method to choose a specific Metal device for rendering when a
// system contains multiple Metal devices. To create a Metal-based context
// using the system’s default Metal device, use the [ContextWithOptions]
// method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:)
func NewContextWithMTLDevice(device metal.MTLDevice) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithMTLDevice:"), device)
	return CIContextFromID(rv)
}

// Creates a Core Image context using the specified Metal device and options.
//
// device: The Metal device object to use for rendering.
//
// options: A dictionary that contains options for creating a [CIContext] object. You
// can pass any of the keys defined in `Context Options` along with the
// appropriate value.
//
// # Return Value
//
// A Core Image context.
//
// # Discussion
//
// Use this method to choose a specific Metal device for rendering when a
// system contains multiple Metal devices. To create a Metal-based context
// using the system’s default Metal device, use the [ContextWithOptions]
// method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:options:)
func NewContextWithMTLDeviceOptions(device metal.MTLDevice, options foundation.INSDictionary) CIContext {
	rv := objc.Send[objc.ID](objc.ID(getCIContextClass().class), objc.Sel("contextWithMTLDevice:options:"), device, options)
	return CIContextFromID(rv)
}

// Initializes a context without a specific rendering destination, using the
// specified options.
//
// options: A dictionary containing options for the context. For applicable keys and
// values, see `Context Options`.
//
// # Return Value
//
// An initialized Core Image context.
//
// # Discussion
//
// If you create a context without specifying a rendering destination, Core
// Image automatically chooses and internally manages a rendering destination
// based on the current device’s capabilities and your settings in the
// `options` dictionary. You cannot use a context without an explicit
// destination for the methods listed in Drawing Images. Instead, use the
// methods listed in Rendering Images.
//
// The `options` dictionary defines behaviors for the context, such as color
// space and rendering quality. For example, to create a CPU-based context,
// use the [useSoftwareRenderer] key.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(options:)
//
// [useSoftwareRenderer]: https://developer.apple.com/documentation/CoreImage/CIContextOption/useSoftwareRenderer
func NewContextWithOptions(options foundation.INSDictionary) CIContext {
	instance := getCIContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return CIContextFromID(rv)
}

// Creates a Core Graphics image from a region of a Core Image image instance.
//
// image: A [CIImage] image instance for which to create a [CGImage].
//
// fromRect: The [CGRect] region of the `image` to use. This region relative to the
// cartesean coordinate system of `image`. This region will be intersected
// with integralized and intersected with `image.Extent()`.
//
// # Return Value
//
// Returns a new [CGImage] instance. You are responsible for releasing the
// returned image when you no longer need it. The returned value will be
// `null` if the extent is empty or too big.
//
// # Discussion
//
// The color space of the created [CGImage] will be sRGB unless the receiving
// [CIContext] was created with a `kCIContextOutputColorSpace` option.
//
// Normally the pixel format of the created CGImage will be 8
// bits-per-component. It will be 16 bits-per-component float if the above
// color space is HDR.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:)
func (c CIContext) CreateCGImageFromRect(image ICIImage, fromRect corefoundation.CGRect) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("createCGImage:fromRect:"), image, fromRect)
	return coregraphics.CGImageRef(rv)
}

// Creates a Core Graphics image from a region of a Core Image image instance
// with an option for controlling the pixel format and color space of the
// [CGImage].
//
// image: A [CIImage] image instance for which to create a [CGImage].
//
// fromRect: The [CGRect] region of the `image` to use. This region relative to the
// cartesean coordinate system of `image`. This region will be intersected
// with integralized and intersected with `image.Extent()`.
//
// format: A [CIFormat] to specify the pixel format of the created [CGImage]. For
// example, if `kCIFormatRGBX16` is specified, then the created [CGImage] will
// be 16 bits-per-component and opaque.
//
// colorSpace: The [CGColorSpace] for the output image. This color space must have either
// `CGColorSpaceModel.Rgb()` or `CGColorSpaceModel.Monochrome()` and be
// compatible with the specified pixel format.
//
// # Return Value
//
// Returns a new [CGImage] instance. You are responsible for releasing the
// returned image when you no longer need it. The returned value will be
// `null` if the extent is empty or too big.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:)
func (c CIContext) CreateCGImageFromRectFormatColorSpace(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:"), image, fromRect, format, colorSpace)
	return coregraphics.CGImageRef(rv)
}

// Creates a Core Graphics image from a region of a Core Image image instance
// with an option for controlling when the image is rendered.
//
// image: A [CIImage] image instance for which to create a [CGImage].
//
// fromRect: The [CGRect] region of the `image` to use. This region relative to the
// cartesean coordinate system of `image`. This region will be intersected
// with integralized and intersected with `image.Extent()`.
//
// format: A [CIFormat] to specify the pixel format of the created [CGImage]. For
// example, if `kCIFormatRGBX16` is specified, then the created [CGImage] will
// be 16 bits-per-component and opaque.
//
// colorSpace: The [CGColorSpace] for the output image. This color space must have either
// `CGColorSpaceModel.Rgb()` or `CGColorSpaceModel.Monochrome()` and be
// compatible with the specified pixel format.
//
// deferred: Controls when Core Image renders `image`.
//
// - True: rendering of `image` is deferred until the created [CGImage]
// rendered. - False: the `image` is rendered immediately.
//
// # Return Value
//
// Returns a new [CGImage] instance. You are responsible for releasing the
// returned image when you no longer need it. The returned value will be
// `null` if the extent is empty or too big.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:)
func (c CIContext) CreateCGImageFromRectFormatColorSpaceDeferred(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef, deferred bool) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:"), image, fromRect, format, colorSpace, deferred)
	return coregraphics.CGImageRef(rv)
}

// Renders to the given bitmap.
//
// image: A Core Image image object.
//
// data: Storage for the bitmap data.
//
// rowBytes: The bytes per row.
//
// bounds: The bounds of the bitmap data.
//
// format: The format of the bitmap data.
//
// colorSpace: The color space for the data. Pass [NULL] if you want to use the output
// color space of the context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:toBitmap:rowBytes:bounds:format:colorSpace:)
func (c CIContext) RenderToBitmapRowBytesBoundsFormatColorSpace(image ICIImage, data unsafe.Pointer, rowBytes int, bounds corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("render:toBitmap:rowBytes:bounds:format:colorSpace:"), image, data, rowBytes, bounds, format, colorSpace)
}

// Renders an image into a pixel buffer.
//
// image: A Core Image image object.
//
// buffer: The destination pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:)
func (c CIContext) RenderToCVPixelBuffer(image ICIImage, buffer corevideo.CVImageBufferRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("render:toCVPixelBuffer:"), image, buffer)
}

// Renders a region of an image into a pixel buffer.
//
// image: A Core Image image object.
//
// buffer: The destination pixel buffer.
//
// bounds: The rectangle in the destination pixel buffer to draw into.
//
// colorSpace: The color space of the destination pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-2k8l2
func (c CIContext) RenderToCVPixelBufferBoundsColorSpace(image ICIImage, buffer corevideo.CVImageBufferRef, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("render:toCVPixelBuffer:bounds:colorSpace:"), image, buffer, bounds, colorSpace)
}

// Renders a region of an image into an IOSurface object.
//
// image: A Core Image image object.
//
// surface: The destination IOSurface object.
//
// bounds: The rectangle in the destination IOSurface object to draw into.
//
// colorSpace: The color space of the destination IOSurface object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-54b9l
func (c CIContext) RenderToIOSurfaceBoundsColorSpace(image ICIImage, surface iosurface.IOSurfaceRef, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("render:toIOSurface:bounds:colorSpace:"), image, surface, bounds, colorSpace)
}

// Renders a region of an image to a Metal texture.
//
// image: A Core Image image object.
//
// texture: The destination Metal texture object.
//
// commandBuffer: The Metal command buffer into which to schedule Core Image rendering tasks.
//
// bounds: The rectangle in the destination texture to draw into.
//
// colorSpace: The color space of the destination texture.
//
// # Discussion
//
// If you specify `nil` for the `commandBuffer` parameter, Core Image manages
// its own Metal command buffer. To combine Core Image rendering with other
// Metal rendering tasks—for example, to use Core Image filters on textures
// whose content is generated by a Metal render-to-texture operation, or to
// use Core Image output later in the same Metal rendering pass—pass the
// same [MTLCommandBuffer] object you use for those tasks.
//
// Rendering to a Metal texture requires a Metal-based context created with
// the [ContextWithMTLDevice] or [ContextWithMTLDeviceOptions] method. Calling
// this method on any other context raises an exception. This method renders
// only to Metal textures whose texture type is [MTLTextureType.type2D] and
// whose [sampleCount] value is 1.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:commandBuffer:bounds:colorSpace:)
//
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [sampleCount]: https://developer.apple.com/documentation/Metal/MTLTexture/sampleCount
func (c CIContext) RenderToMTLTextureCommandBufferBoundsColorSpace(image ICIImage, texture metal.MTLTexture, commandBuffer metal.MTLCommandBuffer, bounds corefoundation.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("render:toMTLTexture:commandBuffer:bounds:colorSpace:"), image, texture, commandBuffer, bounds, colorSpace)
}

// Renders a region of an image to a rectangle in the context destination.
//
// image: A Core Image image object.
//
// inRect: The rectangle in the context destination to draw into. The image is scaled
// to fill the destination rectangle.
//
// fromRect: The subregion of the image that you want to draw into the context, with the
// origin and target size defined by the `dest` parameter. This rectangle is
// always in pixel dimensions.
//
// # Discussion
//
// In iOS, this method draws the [CIImage] object into a renderbuffer for the
// OpenGL ES context. Use this method only if the [CIContext] object is
// created with “ and if you are rendering to a CAEAGLayer. This method is
// asynchronous for apps linked against the iOS 6 or later SDK.
//
// In macOS, you need to be aware of whether the [CIContext] object is created
// with a [CGContextRef] or a [CGLContext] object. If you create the
// [CIContext] object with a [CGContextRef], the dimensions of the destination
// rectangle are in points. If you create the [CIContext] object with a
// [CGLContext] object, the dimensions are in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/draw(_:in:from:)
func (c CIContext) DrawImageInRectFromRect(image ICIImage, inRect corefoundation.CGRect, fromRect corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawImage:inRect:fromRect:"), image, inRect, fromRect)
}

// Frees any cached data, such as temporary images, associated with the
// context and runs the garbage collector.
//
// # Discussion
//
// You can use this method to remove textures from the texture cache that
// reference deleted images.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/clearCaches()
func (c CIContext) ClearCaches() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearCaches"))
}

// Runs the garbage collector to reclaim any resources that the context no
// longer requires.
//
// # Discussion
//
// The system calls this method automatically after every rendering operation.
// You can use this method to remove textures from the texture cache that
// reference deleted images.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/reclaimResources()
func (c CIContext) ReclaimResources() {
	objc.Send[objc.ID](c.ID, objc.Sel("reclaimResources"))
}

// Renders the image and exports the resulting image data in TIFF format.
//
// image: The image object to render.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export.
//
// # Return Value
//
// A data representation of the rendered image in TIFF format, or `nil` if the
// image could not be rendered.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/tiffRepresentation(of:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
func (c CIContext) TIFFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("TIFFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return foundation.NSDataFromID(rv)
}

// Renders the image and exports the resulting image data in JPEG format.
//
// image: The image object to render.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export. Use the
// [kCGImageDestinationLossyCompressionQuality] key to specify JPEG
// compression level. Other supported keys include [avDepthData],
// [depthImage], and [disparityImage].
//
// # Return Value
//
// A data representation of the rendered image in JPEG format, or `nil` if the
// image could not be rendered.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/jpegRepresentation(of:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
// [avDepthData]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/avDepthData
// [depthImage]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/depthImage
// [disparityImage]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/disparityImage
// [kCGImageDestinationLossyCompressionQuality]: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationLossyCompressionQuality
func (c CIContext) JPEGRepresentationOfImageColorSpaceOptions(image ICIImage, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("JPEGRepresentationOfImage:colorSpace:options:"), image, colorSpace, options)
	return foundation.NSDataFromID(rv)
}

// Renders the image and exports the resulting image data in PNG format.
//
// image: The image object to render.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: No options keys are supported at this time.
//
// # Return Value
//
// A data representation of the rendered image in PNG format, or `nil` if the
// image could not be rendered.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/pngRepresentation(of:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
func (c CIContext) PNGRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("PNGRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return foundation.NSDataFromID(rv)
}

// Renders the image and exports the resulting image data in HEIF format.
//
// image: The image object to render.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export. Supported keys include
// [kCGImageDestinationLossyCompressionQuality], [avDepthData], [depthImage],
// and [disparityImage].
//
// # Return Value
//
// A data representation of the rendered image in HEIF format, or `nil` if the
// image could not be rendered.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/heifRepresentation(of:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
// [avDepthData]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/avDepthData
// [depthImage]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/depthImage
// [disparityImage]: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption/disparityImage
// [kCGImageDestinationLossyCompressionQuality]: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationLossyCompressionQuality
func (c CIContext) HEIFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("HEIFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return foundation.NSDataFromID(rv)
}

// Renders the image and exports the resulting image data in HEIF10 format.
//
// image: The image object to render.
//
// colorSpace: The color space in which to render the output image.
//
// options: A dictionary with additional options for export.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/heif10Representation(of:colorSpace:options:)

// DEBUG_TEMPLATE_STRUCT HEIF10RepresentationOfImageColorSpaceOptionsError Params: image ICIImage, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary

func (c CIContext) HEIF10RepresentationOfImageColorSpaceOptionsError(image ICIImage, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("HEIF10RepresentationOfImage:colorSpace:options:error:"), image, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}

// Renders the image and exports the resulting image data in open EXR format.
//
// image: The image object to render.
//
// options: A dictionary with additional options for export.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/openEXRRepresentation(of:options:)
func (c CIContext) OpenEXRRepresentationOfImageOptionsError(image ICIImage, options foundation.INSDictionary) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("OpenEXRRepresentationOfImage:options:error:"), image, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}

// Renders the image and exports the resulting image data as a file in TIFF
// format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output TIFF file.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writeTIFFRepresentation(of:to:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
func (c CIContext) WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeTIFFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeTIFFRepresentationOfImage:toURL:format:colorSpace:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Renders the image and exports the resulting image data as a file in JPEG
// format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output JPEG file.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export. Use the
// [kCGImageDestinationLossyCompressionQuality] key to specify JPEG
// compression level.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writeJPEGRepresentation(of:to:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
// [kCGImageDestinationLossyCompressionQuality]: https://developer.apple.com/documentation/ImageIO/kCGImageDestinationLossyCompressionQuality
func (c CIContext) WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url foundation.INSURL, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeJPEGRepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeJPEGRepresentationOfImage:toURL:colorSpace:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Renders the image and exports the resulting image data as a file in PNG
// format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output PNG file.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writePNGRepresentation(of:to:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
func (c CIContext) WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writePNGRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writePNGRepresentationOfImage:toURL:format:colorSpace:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Renders the image and exports the resulting image data as a file in HEIF
// format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output HEIF file.
//
// format: The pixel format for the output image.
//
// colorSpace: The color space in which to render the output image. This color space must
// conform to either the [CGColorSpaceModel.rgb] or
// [CGColorSpaceModel.monochrome] model and must be compatible with the
// specified pixel format.
//
// options: A dictionary with additional options for export.
//
// # Discussion
//
// To render an image for export, the image’s contents must not be empty and
// its [Extent] dimensions must be finite. To export after applying a filter
// whose output has infinite extent, see the [ImageByClampingToExtent] method.
//
// In Objective-C `writeHEIFRepresentationOfImage` returns true if the file
// export succeeded. If false, examine the `errorPtr` parameter for possible
// failure reasons.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIFRepresentation(of:to:format:colorSpace:options:)
//
// [CGColorSpaceModel.monochrome]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
func (c CIContext) WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url foundation.INSURL, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeHEIFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeHEIFRepresentationOfImage:toURL:format:colorSpace:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Renders the image and exports the resulting image data as a file in HEIF10
// format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output HEIF file.
//
// colorSpace: The color space in which to render the output image.
//
// options: A dictionary with additional options for export.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIF10Representation(of:to:colorSpace:options:)
func (c CIContext) WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url foundation.INSURL, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeHEIF10RepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeHEIF10RepresentationOfImage:toURL:colorSpace:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Renders the image and exports the resulting image data as a file in open
// EXR format.
//
// image: The image object to render.
//
// url: The file URL at which to write the output HEIF file.
//
// options: A dictionary with additional options for export.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/writeOpenEXRRepresentation(of:to:options:)
func (c CIContext) WriteOpenEXRRepresentationOfImageToURLOptionsError(image ICIImage, url foundation.INSURL, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeOpenEXRRepresentationOfImage:toURL:options:error:"), image, url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeOpenEXRRepresentationOfImage:toURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Create a [CIFilter] instance for the supplied image data that can be used
// to apply a depth blur effect created with the supplied auxiliary images.
//
// image: The image object to apply the depth blur effect to.
//
// disparityImage: The auxiliary disparity image. For more information, see
// [auxiliaryDisparity].
//
// portraitEffectsMatte: The auxiliary portrait effects matte image. For more information, see
// [auxiliaryPortraitEffectsMatte].
//
// hairSemanticSegmentation: The auxiliary semantic segmentation hair matte image. For more information,
// see [auxiliarySemanticSegmentationHairMatte].
//
// glassesMatte: The auxiliary glasses matte image. For more information, see
// [auxiliarySemanticSegmentationGlassesMatte].
//
// gainMap: The auxiliary gain map image. For more information, see
// [auxiliaryHDRGainMap].
//
// orientation: The intended display orientation for the image.
//
// options: Reserved for future use.
//
// # Discussion
//
// The receiver context is used to render the image in order to get the facial
// landmarks used to create the effect. The auxiliary images used to create
// the filter can be obtained from a JPEG or HEIC file containing embedded
// portrait effects matte data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:)
//
// [auxiliaryDisparity]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDisparity
// [auxiliaryPortraitEffectsMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryPortraitEffectsMatte
// [auxiliarySemanticSegmentationHairMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationHairMatte
// [auxiliarySemanticSegmentationGlassesMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationGlassesMatte
// [auxiliaryHDRGainMap]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryHDRGainMap
func (c CIContext) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, glassesMatte ICIImage, gainMap ICIImage, orientation uint, options foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, glassesMatte, gainMap, orientation, options)
	return CIFilterFromID(rv)
}

// Create a [CIFilter] instance for the supplied image data that can be used
// to apply a depth blur effect created with the supplied auxiliary images.
//
// image: The image object to apply the depth blur effect to.
//
// disparityImage: The auxiliary disparity image. For more information, see
// [auxiliaryDisparity].
//
// portraitEffectsMatte: The auxiliary portrait effects matte image. For more information, see
// [auxiliaryPortraitEffectsMatte].
//
// hairSemanticSegmentation: The auxiliary semantic segmentation hair matte image. For more information,
// see [auxiliarySemanticSegmentationHairMatte].
//
// orientation: The intended display orientation for the image.
//
// options: Reserved for future use.
//
// # Discussion
//
// The receiver context is used to render the image in order to get the facial
// landmarks used to create the effect. The auxiliary images used to create
// the filter can be obtained from a JPEG or HEIC file containing embedded
// portrait effects matte data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:)
//
// [auxiliaryDisparity]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDisparity
// [auxiliaryPortraitEffectsMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryPortraitEffectsMatte
// [auxiliarySemanticSegmentationHairMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliarySemanticSegmentationHairMatte
func (c CIContext) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, orientation uint, options foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, orientation, options)
	return CIFilterFromID(rv)
}

// Create a [CIFilter] instance for the supplied image data that can be used
// to apply a depth blur effect created with the supplied auxiliary images.
//
// image: The image object to apply the depth blur effect to.
//
// disparityImage: The auxiliary disparity image. For more information, see
// [auxiliaryDisparity].
//
// portraitEffectsMatte: The portrait effects matte image. For more information, see
// [auxiliaryPortraitEffectsMatte].
//
// orientation: The intended display orientation for the image.
//
// options: Reserved for future use.
//
// # Discussion
//
// The receiver context is used to render the image in order to get the facial
// landmarks used to create the effect. The auxiliary images used to create
// the filter can be obtained from a JPEG or HEIC file containing embedded
// portrait effects matte data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:orientation:options:)
//
// [auxiliaryDisparity]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDisparity
// [auxiliaryPortraitEffectsMatte]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryPortraitEffectsMatte
func (c CIContext) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, orientation uint, options foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:orientation:options:"), image, disparityImage, portraitEffectsMatte, orientation, options)
	return CIFilterFromID(rv)
}

// Create a [CIFilter] instance for the supplied image data that can be used
// to apply a depth blur effect.
//
// data: The image file data to apply the depth blur effect to.
//
// options: Reserved for future use.
//
// # Discussion
//
// The receiver context is used to render the image in order to get the facial
// landmarks used to create the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageData:options:)
func (c CIContext) DepthBlurEffectFilterForImageDataOptions(data foundation.INSData, options foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("depthBlurEffectFilterForImageData:options:"), data, options)
	return CIFilterFromID(rv)
}

// Create a [CIFilter] instance for the supplied image URL that can be used to
// apply a depth blur effect.
//
// url: The URL location of the image to apply the depth blur effect to.
//
// options: Reserved for future use.
//
// # Discussion
//
// The receiver context is used to render the image in order to get the facial
// landmarks used to create the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageURL:options:)
func (c CIContext) DepthBlurEffectFilterForImageURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("depthBlurEffectFilterForImageURL:options:"), url, options)
	return CIFilterFromID(rv)
}

// An optional call to warm up a [CIContext] so that subsequent calls to
// render with the same arguments run more efficiently.
//
// image: [CIImage] to prepare to render.
//
// fromRect: A [CGRect] defining the region to render.
//
// destination: The [CIRenderDestination] to which you are preparing to render.
//
// atPoint: The [CGPoint] at which you are preparing to render.
//
// # Discussion
//
// By making this call, the Core Image framework ensures that any needed
// kernels are compiled, and any intermediate buffers are allocated and marked
// volatile up front.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/prepareRender(_:from:to:at:)
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (c CIContext) PrepareRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("prepareRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareRender:fromRect:toDestination:atPoint:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Fills the entire destination with black or clear depending on its
// [AlphaMode].
//
// destination: The [CIRenderDestination] to clear.
//
// # Return Value
//
// The asynchronous [CIRenderTask] for clearing the destination.
//
// # Discussion
//
// If the destination’s [AlphaMode] is [CIRenderDestinationAlphaNone], this
// command fills the entire destination with black `(0, 0, 0, 1)`.
//
// If the destination’s [AlphaMode] is
// [CIRenderDestinationAlphaPremultiplied] or
// [CIRenderDestinationAlphaUnpremultiplied], this command fills the entire
// destination with clear `(0, 0, 0, 0)`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toClear:)
func (c CIContext) StartTaskToClearError(destination ICIRenderDestination) (ICIRenderTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("startTaskToClear:error:"), destination, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIRenderTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIRenderTaskFromID(rv), nil

}

// Renders a portion of an image to a point in the destination.
//
// image: A [CIImage] to render.
//
// fromRect: The part of the image to render, as if cropped.
//
// destination: A [CIRenderDestination] into which to render the image.
//
// atPoint: An origin point in the destination at which to place the image.
//
// # Return Value
//
// An asynchronous [CIRenderTask] to render the image to the specified
// destination.
//
// # Discussion
//
// This method crops the image to the specified rectangle and renders the
// result at the indicated origin point. If the image’s [Extent] property
// and `fromRect` argument values are infinite, this call renders the
// image’s (0, 0) point starting from the origin `atPoint`.
//
// You must use an [MTLTexture]-backed [CIContext] to support an
// [MTLTexture]-backed [CIRenderDestination]. Similarly, you must use
// [GLContext]-backed [CIContext] to support a [GLTexture]-backed
// [CIRenderDestination].
//
// This call returns as soon as it enqueues all work required to render the
// image on the context’s device. In many situations, after issuing a
// render, you may need to wait for it to complete. In these cases, use the
// returned [CIRenderTask] as follows:
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:from:to:at:)
//
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
func (c CIContext) StartTaskToRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint) (ICIRenderTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("startTaskToRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIRenderTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIRenderTaskFromID(rv), nil

}

// Renders an image to a destination so that point (0, 0) of the image maps to
// point (0, 0) of the destination.
//
// image: [CIImage] to prepare to render.
//
// destination: The [CIRenderDestination] to which to render.
//
// # Return Value
//
// The asynchronous [CIRenderTask] to render the image to the specified
// destination.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:to:)
func (c CIContext) StartTaskToRenderToDestinationError(image ICIImage, destination ICIRenderDestination) (ICIRenderTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("startTaskToRender:toDestination:error:"), image, destination, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIRenderTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIRenderTaskFromID(rv), nil

}

// Initializes a context without a specific rendering destination, using the
// specified options.
//
// options: A dictionary containing options for the context. For applicable keys and
// values, see `Context Options`.
//
// # Return Value
//
// An initialized Core Image context.
//
// # Discussion
//
// If you create a context without specifying a rendering destination, Core
// Image automatically chooses and internally manages a rendering destination
// based on the current device’s capabilities and your settings in the
// `options` dictionary. You cannot use a context without an explicit
// destination for the methods listed in Drawing Images. Instead, use the
// methods listed in Rendering Images.
//
// The `options` dictionary defines behaviors for the context, such as color
// space and rendering quality. For example, to create a CPU-based context,
// use the [useSoftwareRenderer] key.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/init(options:)
//
// [useSoftwareRenderer]: https://developer.apple.com/documentation/CoreImage/CIContextOption/useSoftwareRenderer
func (c CIContext) InitWithOptions(options foundation.INSDictionary) CIContext {
	rv := objc.Send[CIContext](c.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// Given a Core Graphics image, use the receiving Core Image context to
// calculate its HDR statistics (content headroom and content average light
// level) and then return a new Core Graphics image that has the calculated
// values.
//
// cgimage: An immutable [CGImage] for which to calculate statistics.
//
// # Return Value
//
// Returns a new [CGImage] instance that has the calculated statistics
// attached.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-3ia7r
func (c CIContext) CalculateHDRStatsForCGImage(cgimage coregraphics.CGImageRef) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("calculateHDRStatsForCGImage:"), cgimage)
	return coregraphics.CGImageRef(rv)
}

// Given an IOSurface, use the receiving Core Image context to calculate its
// HDR statistics (content headroom and content average light level) and then
// update the surface’s attachments to store the values.
//
// surface: A mutable [IOSurfaceRef] for which to calculate and attach statistics.
//
// # Discussion
//
// If the [IOSurface] has a Clean Aperture rectangle then only pixels within
// that rectangle are considered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-6lwmz
func (c CIContext) CalculateHDRStatsForIOSurface(surface iosurface.IOSurfaceRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("calculateHDRStatsForIOSurface:"), surface)
}

// Given a CVPixelBuffer, use the receiving Core Image context to calculate
// its HDR statistics (content headroom and content average light level) and
// then update the buffers’s attachments to store the values.
//
// buffer: A mutable [CVPixelBuffer] for which to calculate and attach statistics.
//
// # Discussion
//
// If the [CVPixelBuffer] has a Clean Aperture rectangle then only pixels
// within that rectangle are considered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-7bcki
func (c CIContext) CalculateHDRStatsForCVPixelBuffer(buffer corevideo.CVImageBufferRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("calculateHDRStatsForCVPixelBuffer:"), buffer)
}

// Given a Core Image image, use the receiving Core Image context to calculate
// its HDR statistics (content headroom and content average light level) and
// then return a new Core Image image that has the calculated values.
//
// image: An immutable [CIImage] for which to calculate statistics.
//
// # Return Value
//
// Returns a new [CIImage] instance that has the calculated statistics
// attached.
//
// # Discussion
//
// If the image extent is not finite, then nil will be returned.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-l1rj
func (c CIContext) CalculateHDRStatsForImage(image ICIImage) ICIImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("calculateHDRStatsForImage:"), image)
	return CIImageFromID(rv)
}

// Creates a Core Graphics image from a region of a Core Image image instance
// with an option for calculating HDR statistics.
//
// image: A [CIImage] image instance for which to create a [CGImage].
//
// fromRect: The [CGRect] region of the `image` to use. This region relative to the
// cartesean coordinate system of `image`. This region will be intersected
// with integralized and intersected with `image.Extent()`.
//
// format: A [CIFormat] to specify the pixel format of the created [CGImage]. For
// example, if `kCIFormatRGBX16` is specified, then the created [CGImage] will
// be 16 bits-per-component and opaque.
//
// colorSpace: The [CGColorSpace] for the output image. This color space must have either
// `CGColorSpaceModel.Rgb()` or `CGColorSpaceModel.Monochrome()` and be
// compatible with the specified pixel format.
//
// deferred: Controls when Core Image renders `image`.
//
// - True: rendering of `image` is deferred until the created [CGImage]
// rendered. - False: the `image` is rendered immediately.
//
// calculateHDRStats: Controls if Core Image calculates HDR statistics.
//
// - True: Core Image will immediately render `image`, calculate the HDR
// statistics and create a [CGImage] that has the calculated values. - False:
// the created [CGImage] will not have any HDR statistics.
//
// # Return Value
//
// Returns a new [CGImage] instance. You are responsible for releasing the
// returned image when you no longer need it. The returned value will be
// `null` if the extent is empty or too big.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:calculateHDRStats:)
func (c CIContext) CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image ICIImage, fromRect corefoundation.CGRect, format int, colorSpace coregraphics.CGColorSpaceRef, deferred bool, calculateHDRStats bool) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:calculateHDRStats:"), image, fromRect, format, colorSpace, deferred, calculateHDRStats)
	return coregraphics.CGImageRef(rv)
}

// Returns the number of GPUs not currently driving a display.
//
// # Return Value
//
// The number of offline GPU devices.
//
// # Discussion
//
// If this count is greater than zero, the system has attached GPU devices
// that are not currently driving a display. You can use these devices for
// Core Image rendering by creating a context with the
// [init(forOfflineGPUAtIndex:)]
// or[init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/offlineGPUCount()
//
// [init(forOfflineGPUAtIndex:)]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:)
// [init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)
func (_CIContextClass CIContextClass) OfflineGPUCount() uint32 {
	rv := objc.Send[uint32](objc.ID(_CIContextClass.class), objc.Sel("offlineGPUCount"))
	return rv
}

// Creates a context without a specific rendering destination, using default
// options.
//
// # Return Value
//
// A new Core Image context.
//
// # Discussion
//
// If you create a context without specifying a rendering destination, Core
// Image automatically chooses and internally manages a rendering destination
// based on the current device’s capabilities. You cannot use a context
// without an explicit destination for the methods listed in Drawing Images.
// Instead, use the methods listed in Rendering Images.
//
// To specify additional options for the context, use the [ContextWithOptions]
// method instead.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/context
func (_CIContextClass CIContextClass) Context() CIContext {
	rv := objc.Send[objc.ID](objc.ID(_CIContextClass.class), objc.Sel("context"))
	return CIContextFromID(rv)
}

// Initializes a context without a specific rendering destination, using the
// specified options.
//
// options: A dictionary containing options for the context. For applicable keys and
// values, see [CIContextOption].
//
// # Return Value
//
// An initialized Core Image context.
//
// # Discussion
//
// If you create a context without specifying a rendering destination, Core
// Image automatically chooses and internally manages a rendering destination
// based on the current device’s capabilities and your settings in the
// `options` dictionary. You cannot use a context without an explicit
// destination for the methods listed in Drawing Images. Instead, use the
// methods listed in Rendering Images.
//
// The `options` dictionary defines behaviors for the context, such as color
// space and rendering quality. For example, to create a CPU-based context,
// use the [useSoftwareRenderer] key.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/contextWithOptions:
//
// [useSoftwareRenderer]: https://developer.apple.com/documentation/CoreImage/CIContextOption/useSoftwareRenderer
func (_CIContextClass CIContextClass) ContextWithOptions(options foundation.INSDictionary) CIContext {
	rv := objc.Send[objc.ID](objc.ID(_CIContextClass.class), objc.Sel("contextWithOptions:"), options)
	return CIContextFromID(rv)
}

// The working color space of the Core Image context.
//
// # Discussion
//
// The working color space determines the color space used when executing
// filter kernels; Core Image automatically converts to and from the source
// and destination color spaces of input images and output contexts. You
// specify a working color space using the [workingColorSpace] key in the
// `options` dictionary when creating a Core Image context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/workingColorSpace
//
// [workingColorSpace]: https://developer.apple.com/documentation/CoreImage/CIContextOption/workingColorSpace
func (c CIContext) WorkingColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](c.ID, objc.Sel("workingColorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}

// The working pixel format of the Core Image context.
//
// # Discussion
//
// The working format determines the pixel format that Core Image uses to
// create intermediate buffers for executing filter kernels. Core Image
// automatically converts to and from the source and destination pixel formats
// of input images and output contexts. You specify a working pixel format
// using the [workingFormat] key in the `options` dictionary when creating a
// Core Image context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/workingFormat
//
// [workingFormat]: https://developer.apple.com/documentation/CoreImage/CIContextOption/workingFormat
func (c CIContext) WorkingFormat() CIFormat {
	rv := objc.Send[CIFormat](c.ID, objc.Sel("workingFormat"))
	return CIFormat(rv)
}

// The render destination’s representation of alpha (transparency) values.
//
// See: https://developer.apple.com/documentation/coreimage/cirenderdestination/alphamode
func (c CIContext) AlphaMode() CIRenderDestinationAlphaMode {
	rv := objc.Send[CIRenderDestinationAlphaMode](c.ID, objc.Sel("alphaMode"))
	return CIRenderDestinationAlphaMode(rv)
}
func (c CIContext) SetAlphaMode(value CIRenderDestinationAlphaMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlphaMode:"), value)
}
