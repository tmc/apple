// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIImage] class.
var (
	_CIImageClass     CIImageClass
	_CIImageClassOnce sync.Once
)

func getCIImageClass() CIImageClass {
	_CIImageClassOnce.Do(func() {
		_CIImageClass = CIImageClass{class: objc.GetClass("CIImage")}
	})
	return _CIImageClass
}

// GetCIImageClass returns the class object for CIImage.
func GetCIImageClass() CIImageClass {
	return getCIImageClass()
}

type CIImageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIImageClass) Alloc() CIImage {
	rv := objc.Send[CIImage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an image to be processed or produced by Core Image
// filters.
//
// # Overview
// 
// You use [CIImage] objects in conjunction with other Core Image
// classes—such as [CIFilter], [CIContext], [CIVector], and [CIColor]—to
// take advantage of the built-in Core Image filters when processing images.
// You can create [CIImage] objects with data supplied from a variety of
// sources, including Quartz 2D images, Core Video image buffers
// ([CVImageBuffer]), URL-based objects, and [NSData] objects.
// 
// Although a [CIImage] object has image data associated with it, it is not an
// image. You can think of a [CIImage] object as an image “recipe.” A
// [CIImage] object has all the information necessary to produce an image, but
// Core Image doesn’t actually render an image until it is told to do so.
// This lazy evaluation allows Core Image to operate as efficiently as
// possible. To show a [CIImage] object as an on-screen image, you can display
// it as a [UIImage] in [UIImageView]:
// 
// [CIContext] and [CIImage] objects are immutable, which means each can be
// shared safely among threads. Multiple threads can use the same GPU or CPU
// [CIContext] object to render [CIImage] objects. However, this is not the
// case for [CIFilter] objects, which are mutable. A [CIFilter] object cannot
// be shared safely among threads. If you app is multithreaded, each thread
// must create its own [CIFilter] objects. Otherwise, your app could behave
// unexpectedly.
// 
// Core Image also provides auto-adjustment methods. These methods analyze an
// image for common deficiencies and return a set of filters to correct those
// deficiencies. The filters are preset with values for improving image
// quality by altering values for skin tones, saturation, contrast, and
// shadows and for removing red-eye or other artifacts caused by flash. (See
// Getting Autoadjustment Filters.)
// 
// For a discussion of all the methods you can use to create [CIImage] objects
// on iOS and macOS, see [Core Image Programming Guide].
//
// [CVImageBuffer]: https://developer.apple.com/documentation/CoreVideo/CVImageBuffer
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
// [UIImageView]: https://developer.apple.com/documentation/UIKit/UIImageView
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
//
// # Creating an Image
//
//   - [CIImage.InitWithContentsOfURL]: Initializes an image object by reading an image from a URL.
//   - [CIImage.InitWithContentsOfURLOptions]: Initializes an image object by reading an image from a URL, using the specified options.
//   - [CIImage.InitWithCGImage]: Initializes an image object with a Quartz 2D image.
//   - [CIImage.InitWithCGImageOptions]: Initializes an image object with a Quartz 2D image, using the specified options.
//   - [CIImage.InitWithCGImageSourceIndexOptions]
//   - [CIImage.InitWithData]: Initializes an image object with the supplied image data.
//   - [CIImage.InitWithDataOptions]: Initializes an image object with the supplied image data, using the specified options.
//   - [CIImage.InitWithBitmapDataBytesPerRowSizeFormatColorSpace]: Initializes an image object with bitmap data.
//   - [CIImage.InitWithBitmapImageRep]: Initializes an image object with the specified bitmap image representation.
//   - [CIImage.InitWithImageProviderSizeFormatColorSpaceOptions]: Initializes an image object based on pixels from an image provider object.
//   - [CIImage.InitWithDepthData]
//   - [CIImage.InitWithDepthDataOptions]
//   - [CIImage.InitWithPortaitEffectsMatte]
//   - [CIImage.InitWithPortaitEffectsMatteOptions]
//   - [CIImage.InitWithSemanticSegmentationMatte]
//   - [CIImage.InitWithSemanticSegmentationMatteOptions]
//   - [CIImage.InitWithCVImageBuffer]: Initializes an image object from the contents of a Core Video image buffer.
//   - [CIImage.InitWithCVImageBufferOptions]: Initializes an image object from the contents of a Core Video image buffer, using the specified options.
//   - [CIImage.InitWithCVPixelBuffer]: Initializes an image object from the contents of a Core Video pixel buffer.
//   - [CIImage.InitWithCVPixelBufferOptions]: Initializes an image object from the contents of a Core Video pixel buffer using the specified options.
//   - [CIImage.InitWithMTLTextureOptions]: Initializes an image object with data supplied by a Metal texture.
//   - [CIImage.InitWithIOSurface]: Initializes an image with the contents of an IOSurface.
//   - [CIImage.InitWithIOSurfaceOptions]: Initializes, using the specified options, an image with the contents of an IOSurface.
//
// # Creating an Image by Modifying an Existing Image
//
//   - [CIImage.ImageByApplyingFilterWithInputParameters]: Returns a new image created by applying a filter to the original image with the specified name and parameters.
//   - [CIImage.ImageByApplyingFilter]: Applies the filter to an image and returns the output.
//   - [CIImage.ImageByApplyingTransform]: Returns a new image that represents the original image after applying an affine transform.
//   - [CIImage.ImageByApplyingTransformHighQualityDownsample]
//   - [CIImage.ImageByCroppingToRect]: Returns a new image with a cropped portion of the original image.
//   - [CIImage.ImageByApplyingOrientation]: Returns a new image created by transforming the original image to the specified EXIF orientation.
//   - [CIImage.ImageByClampingToExtent]: Returns a new image created by making the pixel colors along its edges extend infinitely in all directions.
//   - [CIImage.ImageByClampingToRect]: Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions.
//   - [CIImage.ImageByCompositingOverImage]: Returns a new image created by compositing the original image over the specified destination image.
//   - [CIImage.ImageByConvertingWorkingSpaceToLab]
//   - [CIImage.ImageByConvertingLabToWorkingSpace]
//   - [CIImage.ImageByColorMatchingColorSpaceToWorkingSpace]: Returns a new image created by color matching from the specified color space to the context’s working color space.
//   - [CIImage.ImageByColorMatchingWorkingSpaceToColorSpace]: Returns a new image created by color matching from the context’s working color space to the specified color space.
//   - [CIImage.ImageByPremultiplyingAlpha]: Returns a new image created by multiplying the image’s RGB values by its alpha values.
//   - [CIImage.ImageByUnpremultiplyingAlpha]: Returns a new image created by dividing the image’s RGB values by its alpha values.
//   - [CIImage.ImageBySettingAlphaOneInExtent]: Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area.
//   - [CIImage.ImageByApplyingGaussianBlurWithSigma]: Create an image by applying a gaussian blur to the receiver.
//   - [CIImage.ImageBySettingProperties]: Return a new image by changing the receiver’s metadata properties.
//   - [CIImage.ImageByInsertingIntermediate]: Create an image that inserts a intermediate that is cacheable
//   - [CIImage.ImageByInsertingIntermediateWithCache]: Create an image that inserts a intermediate that is cacheable.
//
// # Creating Solid Colors
//
//   - [CIImage.InitWithColor]: Initializes an image of infinite extent whose entire content is the specified color.
//
// # Getting Image Information
//
//   - [CIImage.Definition]: Returns a filter shape object that represents the domain of definition of the image.
//   - [CIImage.Extent]: A rectangle that specifies the extent of the image.
//   - [CIImage.Properties]: Returns the metadata properties dictionary of the image.
//   - [CIImage.Url]: The URL from which the image was loaded.
//   - [CIImage.ColorSpace]: The color space of the image.
//   - [CIImage.ImageTransformForOrientation]: Returns the transformation needed to reorient the image to the specified orientation.
//
// # Drawing Images
//
//   - [CIImage.DrawAtPointFromRectOperationFraction]: Draws all or part of the image at the specified point in the current coordinate system.
//   - [CIImage.DrawInRectFromRectOperationFraction]: Draws all or part of the image in the specified rectangle in the current coordinate system
//
// # Getting Autoadjustment Filters
//
//   - [CIImage.AutoAdjustmentFilters]: Returns all possible automatically selected and configured filters for adjusting the image.
//   - [CIImage.AutoAdjustmentFiltersWithOptions]: Returns a subset of automatically selected and configured filters for adjusting the image.
//
// # Working with Filter Regions of Interest
//
//   - [CIImage.RegionOfInterestForImageInRect]: Returns the region of interest for the filter chain that generates the image.
//
// # Working with Orientation
//
//   - [CIImage.ImageByApplyingCGOrientation]: Transforms the original image by a given orientation.
//   - [CIImage.ImageTransformForCGOrientation]: The affine transform for changing the image to the given orientation.
//
// # Sampling the Image
//
//   - [CIImage.ImageBySamplingNearest]: Create an image by changing the receiver’s sample mode to nearest neighbor.
//   - [CIImage.ImageBySamplingLinear]: Create an image by changing the receiver’s sample mode to bilinear interpolation.
//
// # Accessing Original Image Content
//
//   - [CIImage.CGImage]: The CoreGraphics image object this image was created from, if applicable.
//   - [CIImage.PixelBuffer]: The CoreVideo pixel buffer this image was created from, if applicable.
//   - [CIImage.DepthData]: Depth data associated with the image.
//   - [CIImage.PortraitEffectsMatte]: The portrait effects matte associated with the image.
//   - [CIImage.SemanticSegmentationMatte]
//
// # Deprecated
//
//   - [CIImage.TextureTarget]: The key for an OpenGL texture target.
//   - [CIImage.TextureFormat]: The key for an OpenGL texture format.
//
// # Instance Properties
//
//   - [CIImage.ContentHeadroom]: Returns the content headroom of the image.
//   - [CIImage.Opaque]: Returns YES if the image is known to have and alpha value of `1.0` over the entire image extent.
//   - [CIImage.MetalTexture]
//   - [CIImage.ContentAverageLightLevel]: Returns the content average light level of the image.
//
// # Instance Methods
//
//   - [CIImage.ImageByApplyingGainMap]: Create an image that applies a gain map Core Image image to the received Core Image image.
//   - [CIImage.ImageByApplyingGainMapHeadroom]: Create an image that applies a gain map Core Image image with a specified headroom to the received Core Image image.
//   - [CIImage.ImageByInsertingTiledIntermediate]: Create an image that inserts a intermediate that is cached in tiles
//   - [CIImage.ImageBySettingContentAverageLightLevel]: Create an image by changing the receiver’s contentAverageLightLevel property.
//   - [CIImage.ImageBySettingContentHeadroom]: Create an image by changing the receiver’s contentHeadroom property.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage
type CIImage struct {
	objectivec.Object
}

// CIImageFromID constructs a [CIImage] from an objc.ID.
//
// A representation of an image to be processed or produced by Core Image
// filters.
func CIImageFromID(id objc.ID) CIImage {
	return CIImage{objectivec.Object{ID: id}}
}
// NOTE: CIImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIImage] class.
//
// # Creating an Image
//
//   - [ICIImage.InitWithContentsOfURL]: Initializes an image object by reading an image from a URL.
//   - [ICIImage.InitWithContentsOfURLOptions]: Initializes an image object by reading an image from a URL, using the specified options.
//   - [ICIImage.InitWithCGImage]: Initializes an image object with a Quartz 2D image.
//   - [ICIImage.InitWithCGImageOptions]: Initializes an image object with a Quartz 2D image, using the specified options.
//   - [ICIImage.InitWithCGImageSourceIndexOptions]
//   - [ICIImage.InitWithData]: Initializes an image object with the supplied image data.
//   - [ICIImage.InitWithDataOptions]: Initializes an image object with the supplied image data, using the specified options.
//   - [ICIImage.InitWithBitmapDataBytesPerRowSizeFormatColorSpace]: Initializes an image object with bitmap data.
//   - [ICIImage.InitWithBitmapImageRep]: Initializes an image object with the specified bitmap image representation.
//   - [ICIImage.InitWithImageProviderSizeFormatColorSpaceOptions]: Initializes an image object based on pixels from an image provider object.
//   - [ICIImage.InitWithDepthData]
//   - [ICIImage.InitWithDepthDataOptions]
//   - [ICIImage.InitWithPortaitEffectsMatte]
//   - [ICIImage.InitWithPortaitEffectsMatteOptions]
//   - [ICIImage.InitWithSemanticSegmentationMatte]
//   - [ICIImage.InitWithSemanticSegmentationMatteOptions]
//   - [ICIImage.InitWithCVImageBuffer]: Initializes an image object from the contents of a Core Video image buffer.
//   - [ICIImage.InitWithCVImageBufferOptions]: Initializes an image object from the contents of a Core Video image buffer, using the specified options.
//   - [ICIImage.InitWithCVPixelBuffer]: Initializes an image object from the contents of a Core Video pixel buffer.
//   - [ICIImage.InitWithCVPixelBufferOptions]: Initializes an image object from the contents of a Core Video pixel buffer using the specified options.
//   - [ICIImage.InitWithMTLTextureOptions]: Initializes an image object with data supplied by a Metal texture.
//   - [ICIImage.InitWithIOSurface]: Initializes an image with the contents of an IOSurface.
//   - [ICIImage.InitWithIOSurfaceOptions]: Initializes, using the specified options, an image with the contents of an IOSurface.
//
// # Creating an Image by Modifying an Existing Image
//
//   - [ICIImage.ImageByApplyingFilterWithInputParameters]: Returns a new image created by applying a filter to the original image with the specified name and parameters.
//   - [ICIImage.ImageByApplyingFilter]: Applies the filter to an image and returns the output.
//   - [ICIImage.ImageByApplyingTransform]: Returns a new image that represents the original image after applying an affine transform.
//   - [ICIImage.ImageByApplyingTransformHighQualityDownsample]
//   - [ICIImage.ImageByCroppingToRect]: Returns a new image with a cropped portion of the original image.
//   - [ICIImage.ImageByApplyingOrientation]: Returns a new image created by transforming the original image to the specified EXIF orientation.
//   - [ICIImage.ImageByClampingToExtent]: Returns a new image created by making the pixel colors along its edges extend infinitely in all directions.
//   - [ICIImage.ImageByClampingToRect]: Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions.
//   - [ICIImage.ImageByCompositingOverImage]: Returns a new image created by compositing the original image over the specified destination image.
//   - [ICIImage.ImageByConvertingWorkingSpaceToLab]
//   - [ICIImage.ImageByConvertingLabToWorkingSpace]
//   - [ICIImage.ImageByColorMatchingColorSpaceToWorkingSpace]: Returns a new image created by color matching from the specified color space to the context’s working color space.
//   - [ICIImage.ImageByColorMatchingWorkingSpaceToColorSpace]: Returns a new image created by color matching from the context’s working color space to the specified color space.
//   - [ICIImage.ImageByPremultiplyingAlpha]: Returns a new image created by multiplying the image’s RGB values by its alpha values.
//   - [ICIImage.ImageByUnpremultiplyingAlpha]: Returns a new image created by dividing the image’s RGB values by its alpha values.
//   - [ICIImage.ImageBySettingAlphaOneInExtent]: Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area.
//   - [ICIImage.ImageByApplyingGaussianBlurWithSigma]: Create an image by applying a gaussian blur to the receiver.
//   - [ICIImage.ImageBySettingProperties]: Return a new image by changing the receiver’s metadata properties.
//   - [ICIImage.ImageByInsertingIntermediate]: Create an image that inserts a intermediate that is cacheable
//   - [ICIImage.ImageByInsertingIntermediateWithCache]: Create an image that inserts a intermediate that is cacheable.
//
// # Creating Solid Colors
//
//   - [ICIImage.InitWithColor]: Initializes an image of infinite extent whose entire content is the specified color.
//
// # Getting Image Information
//
//   - [ICIImage.Definition]: Returns a filter shape object that represents the domain of definition of the image.
//   - [ICIImage.Extent]: A rectangle that specifies the extent of the image.
//   - [ICIImage.Properties]: Returns the metadata properties dictionary of the image.
//   - [ICIImage.Url]: The URL from which the image was loaded.
//   - [ICIImage.ColorSpace]: The color space of the image.
//   - [ICIImage.ImageTransformForOrientation]: Returns the transformation needed to reorient the image to the specified orientation.
//
// # Drawing Images
//
//   - [ICIImage.DrawAtPointFromRectOperationFraction]: Draws all or part of the image at the specified point in the current coordinate system.
//   - [ICIImage.DrawInRectFromRectOperationFraction]: Draws all or part of the image in the specified rectangle in the current coordinate system
//
// # Getting Autoadjustment Filters
//
//   - [ICIImage.AutoAdjustmentFilters]: Returns all possible automatically selected and configured filters for adjusting the image.
//   - [ICIImage.AutoAdjustmentFiltersWithOptions]: Returns a subset of automatically selected and configured filters for adjusting the image.
//
// # Working with Filter Regions of Interest
//
//   - [ICIImage.RegionOfInterestForImageInRect]: Returns the region of interest for the filter chain that generates the image.
//
// # Working with Orientation
//
//   - [ICIImage.ImageByApplyingCGOrientation]: Transforms the original image by a given orientation.
//   - [ICIImage.ImageTransformForCGOrientation]: The affine transform for changing the image to the given orientation.
//
// # Sampling the Image
//
//   - [ICIImage.ImageBySamplingNearest]: Create an image by changing the receiver’s sample mode to nearest neighbor.
//   - [ICIImage.ImageBySamplingLinear]: Create an image by changing the receiver’s sample mode to bilinear interpolation.
//
// # Accessing Original Image Content
//
//   - [ICIImage.CGImage]: The CoreGraphics image object this image was created from, if applicable.
//   - [ICIImage.PixelBuffer]: The CoreVideo pixel buffer this image was created from, if applicable.
//   - [ICIImage.DepthData]: Depth data associated with the image.
//   - [ICIImage.PortraitEffectsMatte]: The portrait effects matte associated with the image.
//   - [ICIImage.SemanticSegmentationMatte]
//
// # Deprecated
//
//   - [ICIImage.TextureTarget]: The key for an OpenGL texture target.
//   - [ICIImage.TextureFormat]: The key for an OpenGL texture format.
//
// # Instance Properties
//
//   - [ICIImage.ContentHeadroom]: Returns the content headroom of the image.
//   - [ICIImage.Opaque]: Returns YES if the image is known to have and alpha value of `1.0` over the entire image extent.
//   - [ICIImage.MetalTexture]
//   - [ICIImage.ContentAverageLightLevel]: Returns the content average light level of the image.
//
// # Instance Methods
//
//   - [ICIImage.ImageByApplyingGainMap]: Create an image that applies a gain map Core Image image to the received Core Image image.
//   - [ICIImage.ImageByApplyingGainMapHeadroom]: Create an image that applies a gain map Core Image image with a specified headroom to the received Core Image image.
//   - [ICIImage.ImageByInsertingTiledIntermediate]: Create an image that inserts a intermediate that is cached in tiles
//   - [ICIImage.ImageBySettingContentAverageLightLevel]: Create an image by changing the receiver’s contentAverageLightLevel property.
//   - [ICIImage.ImageBySettingContentHeadroom]: Create an image by changing the receiver’s contentHeadroom property.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage
type ICIImage interface {
	objectivec.IObject

	// Topic: Creating an Image

	// Initializes an image object by reading an image from a URL.
	InitWithContentsOfURL(url foundation.INSURL) CIImage
	// Initializes an image object by reading an image from a URL, using the specified options.
	InitWithContentsOfURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIImage
	// Initializes an image object with a Quartz 2D image.
	InitWithCGImage(image coregraphics.CGImageRef) CIImage
	// Initializes an image object with a Quartz 2D image, using the specified options.
	InitWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) CIImage
	InitWithCGImageSourceIndexOptions(source objectivec.IObject, index uintptr, dict foundation.INSDictionary) CIImage
	// Initializes an image object with the supplied image data.
	InitWithData(data foundation.INSData) CIImage
	// Initializes an image object with the supplied image data, using the specified options.
	InitWithDataOptions(data foundation.INSData, options foundation.INSDictionary) CIImage
	// Initializes an image object with bitmap data.
	InitWithBitmapDataBytesPerRowSizeFormatColorSpace(data foundation.INSData, bytesPerRow uintptr, size corefoundation.CGSize, format int, colorSpace coregraphics.CGColorSpaceRef) CIImage
	// Initializes an image object with the specified bitmap image representation.
	InitWithBitmapImageRep(bitmapImageRep objectivec.IObject) CIImage
	// Initializes an image object based on pixels from an image provider object.
	InitWithImageProviderSizeFormatColorSpaceOptions(provider objectivec.IObject, width uintptr, height uintptr, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) CIImage
	InitWithDepthData(data objectivec.IObject) CIImage
	InitWithDepthDataOptions(data objectivec.IObject, options foundation.INSDictionary) CIImage
	InitWithPortaitEffectsMatte(matte objectivec.IObject) CIImage
	InitWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage
	InitWithSemanticSegmentationMatte(matte objectivec.IObject) CIImage
	InitWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage
	// Initializes an image object from the contents of a Core Video image buffer.
	InitWithCVImageBuffer(imageBuffer corevideo.CVImageBufferRef) CIImage
	// Initializes an image object from the contents of a Core Video image buffer, using the specified options.
	InitWithCVImageBufferOptions(imageBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage
	// Initializes an image object from the contents of a Core Video pixel buffer.
	InitWithCVPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIImage
	// Initializes an image object from the contents of a Core Video pixel buffer using the specified options.
	InitWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage
	// Initializes an image object with data supplied by a Metal texture.
	InitWithMTLTextureOptions(texture objectivec.IObject, options foundation.INSDictionary) CIImage
	// Initializes an image with the contents of an IOSurface.
	InitWithIOSurface(surface iosurface.IOSurfaceRef) CIImage
	// Initializes, using the specified options, an image with the contents of an IOSurface.
	InitWithIOSurfaceOptions(surface iosurface.IOSurfaceRef, options foundation.INSDictionary) CIImage

	// Topic: Creating an Image by Modifying an Existing Image

	// Returns a new image created by applying a filter to the original image with the specified name and parameters.
	ImageByApplyingFilterWithInputParameters(filterName string, params foundation.INSDictionary) ICIImage
	// Applies the filter to an image and returns the output.
	ImageByApplyingFilter(filterName string) ICIImage
	// Returns a new image that represents the original image after applying an affine transform.
	ImageByApplyingTransform(matrix corefoundation.CGAffineTransform) ICIImage
	ImageByApplyingTransformHighQualityDownsample(matrix corefoundation.CGAffineTransform, highQualityDownsample bool) ICIImage
	// Returns a new image with a cropped portion of the original image.
	ImageByCroppingToRect(rect corefoundation.CGRect) ICIImage
	// Returns a new image created by transforming the original image to the specified EXIF orientation.
	ImageByApplyingOrientation(orientation int) ICIImage
	// Returns a new image created by making the pixel colors along its edges extend infinitely in all directions.
	ImageByClampingToExtent() ICIImage
	// Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions.
	ImageByClampingToRect(rect corefoundation.CGRect) ICIImage
	// Returns a new image created by compositing the original image over the specified destination image.
	ImageByCompositingOverImage(dest ICIImage) ICIImage
	ImageByConvertingWorkingSpaceToLab() ICIImage
	ImageByConvertingLabToWorkingSpace() ICIImage
	// Returns a new image created by color matching from the specified color space to the context’s working color space.
	ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace coregraphics.CGColorSpaceRef) ICIImage
	// Returns a new image created by color matching from the context’s working color space to the specified color space.
	ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace coregraphics.CGColorSpaceRef) ICIImage
	// Returns a new image created by multiplying the image’s RGB values by its alpha values.
	ImageByPremultiplyingAlpha() ICIImage
	// Returns a new image created by dividing the image’s RGB values by its alpha values.
	ImageByUnpremultiplyingAlpha() ICIImage
	// Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area.
	ImageBySettingAlphaOneInExtent(extent corefoundation.CGRect) ICIImage
	// Create an image by applying a gaussian blur to the receiver.
	ImageByApplyingGaussianBlurWithSigma(sigma float64) ICIImage
	// Return a new image by changing the receiver’s metadata properties.
	ImageBySettingProperties(properties foundation.INSDictionary) ICIImage
	// Create an image that inserts a intermediate that is cacheable
	ImageByInsertingIntermediate() ICIImage
	// Create an image that inserts a intermediate that is cacheable.
	ImageByInsertingIntermediateWithCache(cache bool) ICIImage

	// Topic: Creating Solid Colors

	// Initializes an image of infinite extent whose entire content is the specified color.
	InitWithColor(color ICIColor) CIImage

	// Topic: Getting Image Information

	// Returns a filter shape object that represents the domain of definition of the image.
	Definition() ICIFilterShape
	// A rectangle that specifies the extent of the image.
	Extent() corefoundation.CGRect
	// Returns the metadata properties dictionary of the image.
	Properties() foundation.INSDictionary
	// The URL from which the image was loaded.
	Url() foundation.INSURL
	// The color space of the image.
	ColorSpace() coregraphics.CGColorSpaceRef
	// Returns the transformation needed to reorient the image to the specified orientation.
	ImageTransformForOrientation(orientation int) corefoundation.CGAffineTransform

	// Topic: Drawing Images

	// Draws all or part of the image at the specified point in the current coordinate system.
	DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op objectivec.IObject, delta float64)
	// Draws all or part of the image in the specified rectangle in the current coordinate system
	DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op objectivec.IObject, delta float64)

	// Topic: Getting Autoadjustment Filters

	// Returns all possible automatically selected and configured filters for adjusting the image.
	AutoAdjustmentFilters() []CIFilter
	// Returns a subset of automatically selected and configured filters for adjusting the image.
	AutoAdjustmentFiltersWithOptions(options foundation.INSDictionary) []CIFilter

	// Topic: Working with Filter Regions of Interest

	// Returns the region of interest for the filter chain that generates the image.
	RegionOfInterestForImageInRect(image ICIImage, rect corefoundation.CGRect) corefoundation.CGRect

	// Topic: Working with Orientation

	// Transforms the original image by a given orientation.
	ImageByApplyingCGOrientation(orientation objectivec.IObject) ICIImage
	// The affine transform for changing the image to the given orientation.
	ImageTransformForCGOrientation(orientation objectivec.IObject) corefoundation.CGAffineTransform

	// Topic: Sampling the Image

	// Create an image by changing the receiver’s sample mode to nearest neighbor.
	ImageBySamplingNearest() ICIImage
	// Create an image by changing the receiver’s sample mode to bilinear interpolation.
	ImageBySamplingLinear() ICIImage

	// Topic: Accessing Original Image Content

	// The CoreGraphics image object this image was created from, if applicable.
	CGImage() coregraphics.CGImageRef
	// The CoreVideo pixel buffer this image was created from, if applicable.
	PixelBuffer() corevideo.CVImageBufferRef
	// Depth data associated with the image.
	DepthData() objectivec.IObject
	// The portrait effects matte associated with the image.
	PortraitEffectsMatte() objectivec.IObject
	SemanticSegmentationMatte() objectivec.IObject

	// Topic: Deprecated

	// The key for an OpenGL texture target.
	TextureTarget() CIImageOption
	// The key for an OpenGL texture format.
	TextureFormat() CIImageOption

	// Topic: Instance Properties

	// Returns the content headroom of the image.
	ContentHeadroom() float32
	// Returns YES if the image is known to have and alpha value of `1.0` over the entire image extent.
	Opaque() bool
	MetalTexture() objectivec.IObject
	// Returns the content average light level of the image.
	ContentAverageLightLevel() float32

	// Topic: Instance Methods

	// Create an image that applies a gain map Core Image image to the received Core Image image.
	ImageByApplyingGainMap(gainmap ICIImage) ICIImage
	// Create an image that applies a gain map Core Image image with a specified headroom to the received Core Image image.
	ImageByApplyingGainMapHeadroom(gainmap ICIImage, headroom float32) ICIImage
	// Create an image that inserts a intermediate that is cached in tiles
	ImageByInsertingTiledIntermediate() ICIImage
	// Create an image by changing the receiver’s contentAverageLightLevel property.
	ImageBySettingContentAverageLightLevel(average float32) ICIImage
	// Create an image by changing the receiver’s contentHeadroom property.
	ImageBySettingContentHeadroom(headroom float32) ICIImage

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i CIImage) Init() CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i CIImage) Autorelease() CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIImage creates a new CIImage instance.
func NewCIImage() CIImage {
	class := getCIImageClass()
	rv := objc.Send[CIImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an image object with bitmap data.
//
// data: The bitmap data to use for the image. The data you supply must be
// premultiplied.
//
// bytesPerRow: The number of bytes per row.
//
// size: The size of the image data.
//
// format: A pixel format constant. See `Pixel Formats`.
//
// colorSpace: The color space that the image is defined in. It must be a Quartz 2D color
// space ([CGColorSpace]). Pass `nil` for images that don’t contain color
// data (such as elevation maps, normal vector maps, and sampled function
// tables).
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapData:bytesPerRow:size:format:colorSpace:)
func NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data foundation.INSData, bytesPerRow uintptr, size corefoundation.CGSize, format int, colorSpace coregraphics.CGColorSpaceRef) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return CIImageFromID(rv)
}

// Initializes an image object with the specified bitmap image representation.
//
// bitmapImageRep: An image representation object containing the bitmap data.
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapImageRep:)
// bitmapImageRep is a [appkit.NSBitmapImageRep].
func NewImageWithBitmapImageRep(bitmapImageRep objectivec.IObject) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitmapImageRep:"), bitmapImageRep)
	return CIImageFromID(rv)
}

// Initializes an image object with a Quartz 2D image.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:)
func NewImageWithCGImage(image coregraphics.CGImageRef) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:"), image)
	return CIImageFromID(rv)
}

// Initializes an image object with a Quartz 2D image, using the specified
// options.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:options:)
func NewImageWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:options:"), image, options)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImageSource:index:options:)
// source is a [imageio.CGImageSourceRef].
func NewImageWithCGImageSourceIndexOptions(source objectivec.IObject, index uintptr, dict foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImageSource:index:options:"), source, index, dict)
	return CIImageFromID(rv)
}

// Initializes an image object from the contents of a Core Video image buffer.
//
// imageBuffer: A [CVImageBuffer] object in a supported pixel format constant. For more
// information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// # Return Value
// 
// The initialized image object.
//
// # Discussion
// 
// The `imageBuffer` parameter must be in one of the following formats:
// 
// - [kCVPixelFormatType_32ARGB] - [kCVPixelFormatType_422YpCbCr8] -
// [kCVPixelFormatType_32BGRA]
//
// [kCVPixelFormatType_32ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_422YpCbCr8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:)
func NewImageWithCVImageBuffer(imageBuffer corevideo.CVImageBufferRef) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVImageBuffer:"), imageBuffer)
	return CIImageFromID(rv)
}

// Initializes an image object from the contents of a Core Video image buffer,
// using the specified options.
//
// imageBuffer: A [CVImageBuffer] object in a supported pixel format constant. For more
// information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// options: A dictionary that contains options for creating an image object. (See
// `Image Dictionary Keys`.) The pixel format is supplied by the
// [CVImageBuffer] object.)
//
// # Return Value
// 
// The initialized image object.
//
// # Discussion
// 
// The `imageBuffer` parameter must be in one of the following formats:
// 
// - [kCVPixelFormatType_32ARGB] - [kCVPixelFormatType_422YpCbCr8] -
// [kCVPixelFormatType_32BGRA]
//
// [kCVPixelFormatType_32ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_422YpCbCr8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:options:)
func NewImageWithCVImageBufferOptions(imageBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVImageBuffer:options:"), imageBuffer, options)
	return CIImageFromID(rv)
}

// Initializes an image object from the contents of a Core Video pixel buffer.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:)
func NewImageWithCVPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:"), pixelBuffer)
	return CIImageFromID(rv)
}

// Initializes an image object from the contents of a Core Video pixel buffer
// using the specified options.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// options: A dictionary that contains options for creating an image object. (See
// `Image Dictionary Keys`.) The pixel format is supplied by the
// [CVPixelBuffer] object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:options:)
func NewImageWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	return CIImageFromID(rv)
}

// Initializes an image of infinite extent whose entire content is the
// specified color.
//
// color: A color object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(color:)
func NewImageWithColor(color ICIColor) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColor:"), color)
	return CIImageFromID(rv)
}

// Initializes an image object by reading an image from a URL.
//
// url: The location of the image file to read.
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:)
func NewImageWithContentsOfURL(url foundation.INSURL) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return CIImageFromID(rv)
}

// Initializes an image object by reading an image from a URL, using the
// specified options.
//
// url: The location of the image file to read.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:options:)
func NewImageWithContentsOfURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:"), url, options)
	return CIImageFromID(rv)
}

// Initializes an image object with the supplied image data.
//
// data: The image data. The data you supply must be premultiplied.
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:)
func NewImageWithData(data foundation.INSData) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return CIImageFromID(rv)
}

// Initializes an image object with the supplied image data, using the
// specified options.
//
// data: The image data. The data you supply must be premultiplied.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:options:)
func NewImageWithDataOptions(data foundation.INSData, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:"), data, options)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:)
func NewImageWithDepthData(data objectivec.IObject) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDepthData:"), data)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:options:)
func NewImageWithDepthDataOptions(data objectivec.IObject, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDepthData:options:"), data, options)
	return CIImageFromID(rv)
}

// Initializes an image with the contents of an IOSurface.
//
// surface: An IOSurface object.
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface object.
//
// # Discussion
// 
// An IOSurface object is a framebuffer object that is suitable for sharing
// across process boundaries. You can use it to allow your app to move complex
// image decompression and drawing logic into a separate process for the
// purpose of increasing security.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:)
func NewImageWithIOSurface(surface iosurface.IOSurfaceRef) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	return CIImageFromID(rv)
}

// Initializes, using the specified options, an image with the contents of an
// IOSurface.
//
// surface: An IOSurface object.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:options:)
func NewImageWithIOSurfaceOptions(surface iosurface.IOSurfaceRef, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurface:options:"), surface, options)
	return CIImageFromID(rv)
}

// Initializes an image object with the specified UIKit image object.
//
// image: An image containing the source data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:)
func NewImageWithImage(image objectivec.IObject) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:"), image)
	return CIImageFromID(rv)
}

// Initializes an image object with the specified UIKit image object, using
// the specified options.
//
// image: An image containing the source data.
//
// options: A dictionary that contains options for creating an image object. You can
// supply such options as a pixel format and a color space. See `Image
// Dictionary Keys`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:options:)
func NewImageWithImageOptions(image objectivec.IObject, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	return CIImageFromID(rv)
}

// Initializes an image object based on pixels from an image provider object.
//
// provider: An object that implements the [CIImageProvider] protocol.
//
// width: The width of the image.
//
// height: The height of the image.
//
// format: The [CIFormat] of the provided pixels.
//
// colorSpace: The color space that the image is defined in. If `nil`, then the pixels
// will not be is not color matched to the Core Image working color space.
//
// options: A dictionary that contains various [CIImageOption] keys that affect the
// resulting [CIImage]. The option [providerTileSize] controls if and how the
// provider object is called in tiles. The option [providerUserInfo] allows
// additional state to be passed to the provider object.
// //
// [providerTileSize]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerTileSize
// [providerUserInfo]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerUserInfo
//
// # Return Value
// 
// An initialized [CIImage] object based on the data provider.
//
// # Discussion
// 
// Core Image retains the provider object until the image is deallocated. The
// image provider object will not be called until the image is rendered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(imageProvider:size:_:format:colorSpace:options:)
func NewImageWithImageProviderSizeFormatColorSpaceOptions(provider objectivec.IObject, width uintptr, height uintptr, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	return CIImageFromID(rv)
}

// Initializes an image object with data supplied by a Metal texture.
//
// texture: The Metal texture from which to use image data.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the image could not be
// initialized.
//
// # Discussion
// 
// To render the image using Metal, use this image with a Metal-based
// [CIContext] object created with the [ContextWithMTLDevice] method, and call
// the [RenderToMTLTextureCommandBufferBoundsColorSpace] method to create an
// output image in another Metal texture object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(mtlTexture:options:)
func NewImageWithMTLTextureOptions(texture objectivec.IObject, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMTLTexture:options:"), texture, options)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:)
func NewImageWithPortaitEffectsMatte(matte objectivec.IObject) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPortaitEffectsMatte:"), matte)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:options:)
func NewImageWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPortaitEffectsMatte:options:"), matte, options)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:)
func NewImageWithSemanticSegmentationMatte(matte objectivec.IObject) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:"), matte)
	return CIImageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:options:)
func NewImageWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	instance := getCIImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:options:"), matte, options)
	return CIImageFromID(rv)
}

// Initializes an image object by reading an image from a URL.
//
// url: The location of the image file to read.
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:)
func (i CIImage) InitWithContentsOfURL(url foundation.INSURL) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}
// Initializes an image object by reading an image from a URL, using the
// specified options.
//
// url: The location of the image file to read.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:options:)
func (i CIImage) InitWithContentsOfURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithContentsOfURL:options:"), url, options)
	return rv
}
// Initializes an image object with a Quartz 2D image.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:)
func (i CIImage) InitWithCGImage(image coregraphics.CGImageRef) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCGImage:"), image)
	return rv
}
// Initializes an image object with a Quartz 2D image, using the specified
// options.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:options:)
func (i CIImage) InitWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCGImage:options:"), image, options)
	return rv
}
//
// source is a [imageio.CGImageSourceRef].
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImageSource:index:options:)
func (i CIImage) InitWithCGImageSourceIndexOptions(source objectivec.IObject, index uintptr, dict foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCGImageSource:index:options:"), source, index, dict)
	return rv
}
// Initializes an image object with the supplied image data.
//
// data: The image data. The data you supply must be premultiplied.
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:)
func (i CIImage) InitWithData(data foundation.INSData) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithData:"), data)
	return rv
}
// Initializes an image object with the supplied image data, using the
// specified options.
//
// data: The image data. The data you supply must be premultiplied.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:options:)
func (i CIImage) InitWithDataOptions(data foundation.INSData, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithData:options:"), data, options)
	return rv
}
// Initializes an image object with bitmap data.
//
// data: The bitmap data to use for the image. The data you supply must be
// premultiplied.
//
// bytesPerRow: The number of bytes per row.
//
// size: The size of the image data.
//
// format: A pixel format constant. See `Pixel Formats`.
//
// colorSpace: The color space that the image is defined in. It must be a Quartz 2D color
// space ([CGColorSpace]). Pass `nil` for images that don’t contain color
// data (such as elevation maps, normal vector maps, and sampled function
// tables).
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapData:bytesPerRow:size:format:colorSpace:)
func (i CIImage) InitWithBitmapDataBytesPerRowSizeFormatColorSpace(data foundation.INSData, bytesPerRow uintptr, size corefoundation.CGSize, format int, colorSpace coregraphics.CGColorSpaceRef) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return rv
}
// Initializes an image object with the specified bitmap image representation.
//
// bitmapImageRep: An image representation object containing the bitmap data.
//
// bitmapImageRep is a [appkit.NSBitmapImageRep].
//
// # Return Value
// 
// The initialized image object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapImageRep:)
func (i CIImage) InitWithBitmapImageRep(bitmapImageRep objectivec.IObject) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithBitmapImageRep:"), bitmapImageRep)
	return rv
}
// Initializes an image object based on pixels from an image provider object.
//
// provider: An object that implements the [CIImageProvider] protocol.
//
// width: The width of the image.
//
// height: The height of the image.
//
// format: The [CIFormat] of the provided pixels.
//
// colorSpace: The color space that the image is defined in. If `nil`, then the pixels
// will not be is not color matched to the Core Image working color space.
//
// options: A dictionary that contains various [CIImageOption] keys that affect the
// resulting [CIImage]. The option [providerTileSize] controls if and how the
// provider object is called in tiles. The option [providerUserInfo] allows
// additional state to be passed to the provider object.
// //
// [providerTileSize]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerTileSize
// [providerUserInfo]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerUserInfo
//
// # Return Value
// 
// An initialized [CIImage] object based on the data provider.
//
// # Discussion
// 
// Core Image retains the provider object until the image is deallocated. The
// image provider object will not be called until the image is rendered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(imageProvider:size:_:format:colorSpace:options:)
func (i CIImage) InitWithImageProviderSizeFormatColorSpaceOptions(provider objectivec.IObject, width uintptr, height uintptr, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:)
func (i CIImage) InitWithDepthData(data objectivec.IObject) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithDepthData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:options:)
func (i CIImage) InitWithDepthDataOptions(data objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithDepthData:options:"), data, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:)
func (i CIImage) InitWithPortaitEffectsMatte(matte objectivec.IObject) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithPortaitEffectsMatte:"), matte)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:options:)
func (i CIImage) InitWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithPortaitEffectsMatte:options:"), matte, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:)
func (i CIImage) InitWithSemanticSegmentationMatte(matte objectivec.IObject) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithSemanticSegmentationMatte:"), matte)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:options:)
func (i CIImage) InitWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithSemanticSegmentationMatte:options:"), matte, options)
	return rv
}
// Initializes an image object from the contents of a Core Video image buffer.
//
// imageBuffer: A [CVImageBuffer] object in a supported pixel format constant. For more
// information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// # Return Value
// 
// The initialized image object.
//
// # Discussion
// 
// The `imageBuffer` parameter must be in one of the following formats:
// 
// - [kCVPixelFormatType_32ARGB] - [kCVPixelFormatType_422YpCbCr8] -
// [kCVPixelFormatType_32BGRA]
//
// [kCVPixelFormatType_32ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_422YpCbCr8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:)
func (i CIImage) InitWithCVImageBuffer(imageBuffer corevideo.CVImageBufferRef) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCVImageBuffer:"), imageBuffer)
	return rv
}
// Initializes an image object from the contents of a Core Video image buffer,
// using the specified options.
//
// imageBuffer: A [CVImageBuffer] object in a supported pixel format constant. For more
// information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// options: A dictionary that contains options for creating an image object. (See
// `Image Dictionary Keys`.) The pixel format is supplied by the
// [CVImageBuffer] object.)
//
// # Return Value
// 
// The initialized image object.
//
// # Discussion
// 
// The `imageBuffer` parameter must be in one of the following formats:
// 
// - [kCVPixelFormatType_32ARGB] - [kCVPixelFormatType_422YpCbCr8] -
// [kCVPixelFormatType_32BGRA]
//
// [kCVPixelFormatType_32ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_422YpCbCr8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:options:)
func (i CIImage) InitWithCVImageBufferOptions(imageBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCVImageBuffer:options:"), imageBuffer, options)
	return rv
}
// Initializes an image object from the contents of a Core Video pixel buffer.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:)
func (i CIImage) InitWithCVPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCVPixelBuffer:"), pixelBuffer)
	return rv
}
// Initializes an image object from the contents of a Core Video pixel buffer
// using the specified options.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// options: A dictionary that contains options for creating an image object. (See
// `Image Dictionary Keys`.) The pixel format is supplied by the
// [CVPixelBuffer] object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:options:)
func (i CIImage) InitWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}
// Initializes an image object with data supplied by a Metal texture.
//
// texture: The Metal texture from which to use image data.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// The initialized image object, or `nil` if the image could not be
// initialized.
//
// # Discussion
// 
// To render the image using Metal, use this image with a Metal-based
// [CIContext] object created with the [ContextWithMTLDevice] method, and call
// the [RenderToMTLTextureCommandBufferBoundsColorSpace] method to create an
// output image in another Metal texture object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(mtlTexture:options:)
func (i CIImage) InitWithMTLTextureOptions(texture objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithMTLTexture:options:"), texture, options)
	return rv
}
// Initializes an image with the contents of an IOSurface.
//
// surface: An IOSurface object.
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface object.
//
// # Discussion
// 
// An IOSurface object is a framebuffer object that is suitable for sharing
// across process boundaries. You can use it to allow your app to move complex
// image decompression and drawing logic into a separate process for the
// purpose of increasing security.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:)
func (i CIImage) InitWithIOSurface(surface iosurface.IOSurfaceRef) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithIOSurface:"), surface)
	return rv
}
// Initializes, using the specified options, an image with the contents of an
// IOSurface.
//
// surface: An IOSurface object.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:options:)
func (i CIImage) InitWithIOSurfaceOptions(surface iosurface.IOSurfaceRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithIOSurface:options:"), surface, options)
	return rv
}
// Returns a new image created by applying a filter to the original image with
// the specified name and parameters.
//
// filterName: The name of the filter to apply, as used when creating a [CIFilter]
// instance with the [FilterWithName] method.
//
// params: A dictionary whose key-value pairs are set as input values to the filter.
// Each key is a constant that specifies the name of an input parameter for
// the filter, and the corresponding value is the value for that parameter.
// See [Core Image Filter Reference] for built-in filters and their allowed
// parameters.
// //
// [Core Image Filter Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/uid/TP40004346
//
// # Return Value
// 
// An image object representing the result of applying the filter.
//
// # Discussion
// 
// Calling this method is equivalent to the following sequence of steps:
// 
// - Creating a [CIFilter] instance - Setting the original image as the
// filter’s `inputImage` parameter - Setting the remaining filter parameters
// from the `params` dictionary - Retrieving the [OutputImage] object from the
// filter
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:parameters:)
func (i CIImage) ImageByApplyingFilterWithInputParameters(filterName string, params foundation.INSDictionary) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingFilter:withInputParameters:"), objc.String(filterName), params)
	return CIImageFromID(rv)
}
// Applies the filter to an image and returns the output.
//
// # Discussion
// 
// A convenience method for applying a single filter to the method receiver
// and returning the output image. Identical to
// [ImageByApplyingFilterWithInputParameters] with default parameters.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:)
func (i CIImage) ImageByApplyingFilter(filterName string) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingFilter:"), objc.String(filterName))
	return CIImageFromID(rv)
}
// Returns a new image that represents the original image after applying an
// affine transform.
//
// matrix: An affine transform.
//
// # Return Value
// 
// The transformed image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:)
func (i CIImage) ImageByApplyingTransform(matrix corefoundation.CGAffineTransform) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingTransform:"), matrix)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:highQualityDownsample:)
func (i CIImage) ImageByApplyingTransformHighQualityDownsample(matrix corefoundation.CGAffineTransform, highQualityDownsample bool) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingTransform:highQualityDownsample:"), matrix, highQualityDownsample)
	return CIImageFromID(rv)
}
// Returns a new image with a cropped portion of the original image.
//
// rect: The rectangle, in image coordinates, to which to crop the image.
//
// # Return Value
// 
// An image object cropped to the specified rectangle.
//
// # Discussion
// 
// [media-2951307]
// 
// # Discussion
// 
// Due to Core Image’s coordinate system mismatch with [UIKit], this
// filtering approach may yield unexpected results when displayed in a
// [UIImageView] with [contentMode]. Be sure to back it with a [CGImage] so
// that it handles [contentMode] properly.
// 
// If you are displaying or processing your image primarily as a [CGImage] or
// [UIImage], with no additional Core Image application, consider cropping in
// Core Graphics using the [cropping(to:)] function to save processing
// overhead from conversion of images to [CIImage]. It makes most sense to use
// [ImageByCroppingToRect] when you already have [CIImage] in your pipeline.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [UIImageView]: https://developer.apple.com/documentation/UIKit/UIImageView
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
// [UIKit]: https://developer.apple.com/library/archive/releasenotes/General/WhatsNewIniOS/Articles/iOS5.html#//apple_ref/doc/uid/TP30915195-SW41
// [contentMode]: https://developer.apple.com/documentation/UIKit/UIView/contentMode-swift.property
// [cropping(to:)]: https://developer.apple.com/documentation/CoreGraphics/CGImage/cropping(to:)
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/cropped(to:)
func (i CIImage) ImageByCroppingToRect(rect corefoundation.CGRect) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByCroppingToRect:"), rect)
	return CIImageFromID(rv)
}
// Returns a new image created by transforming the original image to the
// specified EXIF orientation.
//
// orientation: An integer specifying an image orientation according to the EXIF
// specification. For details, see [kCGImagePropertyOrientation].
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// # Return Value
// 
// An image object representing the result of rotating or mirroring the image
// to the target orientation.
//
// # Discussion
// 
// This method determines and then applies the transformation needed to
// reorient the image to the specified orientation. If you plan to also apply
// other transformations, you can retrieve the transformation this method
// would use by calling the [ImageTransformForOrientation] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(forExifOrientation:)
func (i CIImage) ImageByApplyingOrientation(orientation int) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingOrientation:"), orientation)
	return CIImageFromID(rv)
}
// Returns a new image created by making the pixel colors along its edges
// extend infinitely in all directions.
//
// # Return Value
// 
// An image object representing the result of the clamp operation.
//
// # Discussion
// 
// Calling this method is equivalent to using the [CIAffineClamp] filter,
// which creates an image of infinite extent by repeating pixel colors from
// the edges of the original image.
// 
// This operation can be useful when using the image as input to other
// filters. When an image has finite extent, Core Image treats the area
// outside the extent as if it were filled with empty (black, zero alpha)
// pixels. If you apply a filter that samples from outside the image’s
// extent, those empty pixels affect the result of the filter.
// 
// For example, applying the [CIGaussianBlur] filter to an image softens the
// edges of the blurred image, because the opaque pixels at the edges of the
// image blur into the transparent pixels outside the image’s extent.
// Applying a clamp effect before the blur filter avoids edge softening by
// making the original image opaque in all directions. (However, the blurred
// image will also have infinite extent. Use the [ImageByCroppingToRect]
// method to return to the original image’s dimensions while retaining hard
// edges.)
//
// [CIAffineClamp]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CIAffineClamp
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/clampedToExtent()
func (i CIImage) ImageByClampingToExtent() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByClampingToExtent"))
	return CIImageFromID(rv)
}
// Returns a new image created by cropping to a specified area, then making
// the pixel colors along the edges of the cropped image extend infinitely in
// all directions.
//
// rect: The rectangle, in image coordinates, to which to crop the image.
//
// # Return Value
// 
// An image object representing the result of the clamp operation.
//
// # Discussion
// 
// Calling this method is equivalent to cropping the image (with the
// [ImageByCroppingToRect] method or the [CICrop] filter), then using the
// [ImageByClampingToExtent] method (or the [CIAffineClamp] filter), which
// creates an image of infinite extent by repeating pixel colors from the
// edges of the cropped image.
//
// [CIAffineClamp]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CIAffineClamp
// [CICrop]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CICrop
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/clamped(to:)
func (i CIImage) ImageByClampingToRect(rect corefoundation.CGRect) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByClampingToRect:"), rect)
	return CIImageFromID(rv)
}
// Returns a new image created by compositing the original image over the
// specified destination image.
//
// dest: An image to serve as the destination of the compositing operation.
//
// # Return Value
// 
// An image object representing the result of the compositing operation.
//
// # Discussion
// 
// Calling this method is equivalent to using the [CISourceOverCompositing]
// filter. To use other compositing operations and blending modes, create a
// [CIFilter] object using one of the built-in filters from the
// [CICategoryCompositeOperation] category. For details, see [Core Image
// Filter Reference].
//
// [CICategoryCompositeOperation]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/uid/TP30000136-SW71
// [CISourceOverCompositing]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CISourceOverCompositing
// [Core Image Filter Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/uid/TP40004346
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/composited(over:)
func (i CIImage) ImageByCompositingOverImage(dest ICIImage) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByCompositingOverImage:"), dest)
	return CIImageFromID(rv)
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/convertingWorkingSpaceToLab()
func (i CIImage) ImageByConvertingWorkingSpaceToLab() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByConvertingWorkingSpaceToLab"))
	return CIImageFromID(rv)
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/convertingLabToWorkingSpace()
func (i CIImage) ImageByConvertingLabToWorkingSpace() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByConvertingLabToWorkingSpace"))
	return CIImageFromID(rv)
}
// Returns a new image created by color matching from the specified color
// space to the context’s working color space.
//
// colorSpace: The color space to be converted from. This color space must conform to the
// [CGColorSpaceModel.rgb] color space model.
// //
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
//
// # Return Value
// 
// An image object representing the result of the color matching operation, or
// `nil` if the color spaces to be converted are not compatible.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/matchedToWorkingSpace(from:)
func (i CIImage) ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace coregraphics.CGColorSpaceRef) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByColorMatchingColorSpaceToWorkingSpace:"), colorSpace)
	return CIImageFromID(rv)
}
// Returns a new image created by color matching from the context’s working
// color space to the specified color space.
//
// colorSpace: The color space to be converted to. This color space must conform to the
// [CGColorSpaceModel.rgb] color space model.
// //
// [CGColorSpaceModel.rgb]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
//
// # Return Value
// 
// An image object representing the result of the color matching operation, or
// `nil` if the color spaces to be converted are not compatible.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/matchedFromWorkingSpace(to:)
func (i CIImage) ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace coregraphics.CGColorSpaceRef) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByColorMatchingWorkingSpaceToColorSpace:"), colorSpace)
	return CIImageFromID(rv)
}
// Returns a new image created by multiplying the image’s RGB values by its
// alpha values.
//
// # Return Value
// 
// An image object representing the result of the operation.
//
// # Discussion
// 
// Premultiplied alpha speeds up the rendering of images, so Core Image
// filters require that input image data be premultiplied. If you have an
// image without premultiplied alpha that you want to feed into a filter, use
// this method before applying the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/premultiplyingAlpha()
func (i CIImage) ImageByPremultiplyingAlpha() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByPremultiplyingAlpha"))
	return CIImageFromID(rv)
}
// Returns a new image created by dividing the image’s RGB values by its
// alpha values.
//
// # Return Value
// 
// An image object representing the result of the operation.
//
// # Discussion
// 
// Premultiplied alpha speeds up the rendering of images, but some custom
// filter routines can be expressed more efficiently with non-premultiplied
// RGB values. Use this method if you need to apply such a filter to an image
// that has premultiplied alpha.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/unpremultiplyingAlpha()
func (i CIImage) ImageByUnpremultiplyingAlpha() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByUnpremultiplyingAlpha"))
	return CIImageFromID(rv)
}
// Returns a new image created by setting all alpha values to 1.0 within the
// specified rectangle and to 0.0 outside of that area.
//
// extent: The rectangular area in the image to be set to full opacity.
//
// # Return Value
// 
// An image object representing the result of the operation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/settingAlphaOne(in:)
func (i CIImage) ImageBySettingAlphaOneInExtent(extent corefoundation.CGRect) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySettingAlphaOneInExtent:"), extent)
	return CIImageFromID(rv)
}
// Create an image by applying a gaussian blur to the receiver.
//
// sigma: The sigma of the gaussian blur to apply to the receiver. If the sigma is
// very small (less than `0.16`) then the receiver is returned.
//
// # Return Value
// 
// An autoreleased [CIImage] instance or the received image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGaussianBlur(sigma:)
func (i CIImage) ImageByApplyingGaussianBlurWithSigma(sigma float64) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingGaussianBlurWithSigma:"), sigma)
	return CIImageFromID(rv)
}
// Return a new image by changing the receiver’s metadata properties.
//
// properties: A dictionary of metadata properties akin to the
// `CGImageSourceCopyPropertiesAtIndex()` function.
//
// # Return Value
// 
// An autoreleased [CIImage] instance with a copy of the new properties.
//
// # Discussion
// 
// When you create an image, Core Image sets an image’s properties to a
// metadata dictionary as described here: [Properties]. Use this method to
// override an image’s metadata properties with new values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/settingProperties(_:)
func (i CIImage) ImageBySettingProperties(properties foundation.INSDictionary) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySettingProperties:"), properties)
	return CIImageFromID(rv)
}
// Create an image that inserts a intermediate that is cacheable
//
// # Return Value
// 
// An autoreleased [CIImage].
//
// # Discussion
// 
// This intermediate will be not be cached if [cacheIntermediates] is false.
//
// [cacheIntermediates]: https://developer.apple.com/documentation/CoreImage/CIContextOption/cacheIntermediates
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate()
func (i CIImage) ImageByInsertingIntermediate() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByInsertingIntermediate"))
	return CIImageFromID(rv)
}
// Create an image that inserts a intermediate that is cacheable.
//
// cache: Controls if Core Image caches the returned image. * [YES] : This
// intermediate will be cacheable even if [cacheIntermediates] is false. *
// [NO] : the intermediate will be not be cached if [cacheIntermediates] is
// false.
// //
// [cacheIntermediates]: https://developer.apple.com/documentation/CoreImage/CIContextOption/cacheIntermediates
//
// # Return Value
// 
// An autoreleased [CIImage].
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate(cache:)
func (i CIImage) ImageByInsertingIntermediateWithCache(cache bool) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByInsertingIntermediate:"), cache)
	return CIImageFromID(rv)
}
// Initializes an image of infinite extent whose entire content is the
// specified color.
//
// color: A color object.
//
// # Return Value
// 
// The initialized image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/init(color:)
func (i CIImage) InitWithColor(color ICIColor) CIImage {
	rv := objc.Send[CIImage](i.ID, objc.Sel("initWithColor:"), color)
	return rv
}
// Returns the transformation needed to reorient the image to the specified
// orientation.
//
// orientation: An integer specifying an image orientation according to the EXIF
// specification. For details, see [kCGImagePropertyOrientation].
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// # Return Value
// 
// An affine transform that will rotate or mirror the image to match the
// specified orientation when applied.
//
// # Discussion
// 
// This method determines the transformation needed to match the specified
// orientation, but does not apply that transformation to the image. To apply
// the transformation (possibly after concatenating it with other
// transformations), use the [ImageByApplyingTransform] method or the
// [CIAffineTransform] filter. To determine and apply the transformation in a
// single step, use the [ImageByApplyingOrientation] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(forExifOrientation:)
func (i CIImage) ImageTransformForOrientation(orientation int) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i.ID, objc.Sel("imageTransformForOrientation:"), orientation)
	return corefoundation.CGAffineTransform(rv)
}
// Draws all or part of the image at the specified point in the current
// coordinate system.
//
// point: The location in the current coordinate system at which to draw the image.
//
// fromRect: The source rectangle specifying the portion of the image you want to draw.
// The coordinates of this rectangle must be specified using the image’s own
// coordinate system.
//
// op: The compositing operation to use when drawing the image. For details, see
// [NSCompositingOperation].
// //
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
//
// delta: The opacity of the image, specified as a value from `0.0` to `1.0`.
// Specifying a value of `0.0` draws the image as fully transparent while a
// value of `1.0` draws the image as fully opaque. Values greater than `1.0`
// are interpreted as `1.0`.
//
// op is a [appkit.NSCompositingOperation].
//
// # Discussion
// 
// The image content is drawn at its current resolution and is not scaled
// unless the CTM of the current coordinate system itself contains a scaling
// factor. The image is otherwise positioned and oriented using the current
// coordinate system.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/draw(at:from:operation:fraction:)
func (i CIImage) DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op objectivec.IObject, delta float64) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}
// Draws all or part of the image in the specified rectangle in the current
// coordinate system
//
// rect: The rectangle in which to draw the image.
//
// fromRect: The source rectangle specifying the portion of the image you want to draw.
// The coordinates of this rectangle must be specified using the image’s own
// coordinate system.
//
// op: The compositing operation to use when drawing the image. For details, see
// [NSCompositingOperation].
// //
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
//
// delta: The opacity of the image, specified as a value from `0.0` to `1.0`.
// Specifying a value of `0.0` draws the image as fully transparent while a
// value of `1.0` draws the image as fully opaque. Values greater than `1.0`
// are interpreted as `1.0`.
//
// op is a [appkit.NSCompositingOperation].
//
// # Discussion
// 
// If the `srcRect` and `dstRect` rectangles have different sizes, the source
// portion of the image is scaled to fit the specified destination rectangle.
// The image is otherwise positioned and oriented using the current coordinate
// system.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/draw(in:from:operation:fraction:)
func (i CIImage) DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op objectivec.IObject, delta float64) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}
// Returns all possible automatically selected and configured filters for
// adjusting the image.
//
// # Return Value
// 
// An array of [CIFilter] instances preconfigured for correcting deficiencies
// in the supplied image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters()
func (i CIImage) AutoAdjustmentFilters() []CIFilter {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("autoAdjustmentFilters"))
	return objc.ConvertSlice(rv, func(id objc.ID) CIFilter {
		return CIFilterFromID(id)
	})
}
// Returns a subset of automatically selected and configured filters for
// adjusting the image.
//
// options: You can control which filters are returned by supplying one or more of the
// keys described in [Autoadjustment Keys].
// 
// The options dictionary can also contain a [CIDetectorImageOrientation] key.
// Because some autoadjustment filters rely on face detection, you should
// specify an image orientation if you want to enable these filters for an
// image containing face whose orientation does not match that of the image.
// //
// [Autoadjustment Keys]: https://developer.apple.com/documentation/CoreImage/autoadjustment-keys
// [CIDetectorImageOrientation]: https://developer.apple.com/documentation/CoreImage/CIDetectorImageOrientation
//
// # Return Value
// 
// An array of [CIFilter] instances preconfigured for correcting deficiencies
// in the supplied image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters(options:)
func (i CIImage) AutoAdjustmentFiltersWithOptions(options foundation.INSDictionary) []CIFilter {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("autoAdjustmentFiltersWithOptions:"), options)
	return objc.ConvertSlice(rv, func(id objc.ID) CIFilter {
		return CIFilterFromID(id)
	})
}
// Returns the region of interest for the filter chain that generates the
// image.
//
// image: Another image that is part of the filter chain that generates the image.
//
// rect: A rectangle in the image’s coordinate space.
//
// # Return Value
// 
// A rectangle in the coordinate space of the input image (the `im`
// parameter).
//
// # Discussion
// 
// The region of interest is the rectangle containing pixel data in a source
// image (the `im` parameter) necessary to produce a corresponding rectangle
// in the output image. If the image is not the output of a filter (or of a
// chain or graph of several [CIFilter] objects), or the image in the `im`
// parameter is not an input to that filter, the rectangle returned is the
// same as that in the `r` parameter.
// 
// For example,
// 
// - If the image is the output of a filter that doubles the size of its input
// image, the rectangle returned will be half the size of that in the `r`
// parameter. (Upscaling causes every pixel in the input image to correspond
// to multiple pixels in the output image.) - If the image is the output of a
// blur filter, the rectangle returned will be slightly larger than that in
// the `r` parameter. (In a blur filter, each pixel in the output image is
// produced using information from the corresponding pixel and those
// immediately surrounding it in the input image.)
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/regionOfInterest(for:in:)
func (i CIImage) RegionOfInterestForImageInRect(image ICIImage, rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("regionOfInterestForImage:inRect:"), image, rect)
	return corefoundation.CGRect(rv)
}
// Transforms the original image by a given orientation.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// # Discussion
// 
// Returns a new image representing the original image transformed for the
// given [CGImagePropertyOrientation].
//
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(_:)
func (i CIImage) ImageByApplyingCGOrientation(orientation objectivec.IObject) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingCGOrientation:"), orientation)
	return CIImageFromID(rv)
}
// The affine transform for changing the image to the given orientation.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// # Discussion
// 
// Returns a [CGAffineTransform] for the [CGImagePropertyOrientation] value to
// apply to the image.
//
// [CGAffineTransform]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(for:)
func (i CIImage) ImageTransformForCGOrientation(orientation objectivec.IObject) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i.ID, objc.Sel("imageTransformForCGOrientation:"), orientation)
	return corefoundation.CGAffineTransform(rv)
}
// Create an image by changing the receiver’s sample mode to nearest
// neighbor.
//
// # Return Value
// 
// An autoreleased [CIImage] instance with a nearest sampling.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/samplingNearest()
func (i CIImage) ImageBySamplingNearest() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySamplingNearest"))
	return CIImageFromID(rv)
}
// Create an image by changing the receiver’s sample mode to bilinear
// interpolation.
//
// # Return Value
// 
// An autoreleased [CIImage] instance with a bilinear sampling.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/samplingLinear()
func (i CIImage) ImageBySamplingLinear() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySamplingLinear"))
	return CIImageFromID(rv)
}
// Create an image that applies a gain map Core Image image to the received
// Core Image image.
//
// # Return Value
// 
// An autoreleased [CIImage] instance or the received image.
//
// # Discussion
// 
// The gain map image can be obtained by creating a [CIImage] instance from
// [NSURL]/[NSData] and setting the [auxiliaryHDRGainMap] option set to
// `@YES`.
// 
// If the gain map [CIImage] instance doesn’t have the needed [Properties]
// metadata, the received image will be returned as-is.
//
// [auxiliaryHDRGainMap]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryHDRGainMap
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:)
func (i CIImage) ImageByApplyingGainMap(gainmap ICIImage) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingGainMap:"), gainmap)
	return CIImageFromID(rv)
}
// Create an image that applies a gain map Core Image image with a specified
// headroom to the received Core Image image.
//
// gainmap: The gain map [CIImage] instance to apply to the receiver.
//
// headroom: A float value that specify how much headroom the resulting image should
// have. The headroom value will be limited to between 1.0 (i.e. SDR) and the
// full headroom allowed by the gain map.
//
// # Return Value
// 
// An autoreleased [CIImage] instance or the received image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:headroom:)
func (i CIImage) ImageByApplyingGainMapHeadroom(gainmap ICIImage, headroom float32) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByApplyingGainMap:headroom:"), gainmap, headroom)
	return CIImageFromID(rv)
}
// Create an image that inserts a intermediate that is cached in tiles
//
// # Return Value
// 
// An autoreleased [CIImage].
//
// # Discussion
// 
// This intermediate will be cacheable even if [cacheIntermediates] is false.
//
// [cacheIntermediates]: https://developer.apple.com/documentation/CoreImage/CIContextOption/cacheIntermediates
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/insertingTiledIntermediate()
func (i CIImage) ImageByInsertingTiledIntermediate() ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageByInsertingTiledIntermediate"))
	return CIImageFromID(rv)
}
// Create an image by changing the receiver’s contentAverageLightLevel
// property.
//
// # Return Value
// 
// An autoreleased [CIImage].
//
// # Discussion
// 
// Changing this value will alter the behavior of the [CIToneMapHeadroom] and
// [CISystemToneMap] filters.
// 
// - If the value is set to 0.0 or less then the returned image’s
// [ContentAverageLightLevel] is unknown.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentAverageLightLevel(_:)
func (i CIImage) ImageBySettingContentAverageLightLevel(average float32) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySettingContentAverageLightLevel:"), average)
	return CIImageFromID(rv)
}
// Create an image by changing the receiver’s contentHeadroom property.
//
// # Return Value
// 
// An autoreleased [CIImage].
//
// # Discussion
// 
// Changing this value will alter the behavior of the [CIToneMapHeadroom] and
// [CISystemToneMap] filters.
// 
// - If the value is set to 0.0 then the returned image’s headroom is
// unknown. - If the value is set to 1.0 then the returned image is SDR. - If
// the value is set to greater 1.0 then the returned image is HDR. - Otherwise
// the returned image’s headroom is unknown.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentHeadroom(_:)
func (i CIImage) ImageBySettingContentHeadroom(headroom float32) ICIImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageBySettingContentHeadroom:"), headroom)
	return CIImageFromID(rv)
}
func (i CIImage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an empty image object.
//
// # Return Value
// 
// An image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/empty()
func (_CIImageClass CIImageClass) EmptyImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("emptyImage"))
	return CIImageFromID(rv)
}
// Creates and returns an image object from bitmap data.
//
// data: The bitmap data for the image. This data must be premultiplied.
//
// bytesPerRow: The number of bytes per row.
//
// size: The dimensions of the image.
//
// format: The format and size of each pixel. You must supply a pixel format constant.
// See `Pixel Formats`.
//
// colorSpace: The color space that the image is defined in. If this value is `nil`, the
// image is not color matched. Pass `nil` for images that don’t contain
// color data (such as elevation maps, normal vector maps, and sampled
// function tables).
//
// # Return Value
// 
// An image object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithBitmapData:bytesPerRow:size:format:colorSpace:
func (_CIImageClass CIImageClass) ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data foundation.INSData, bytesPerRow uintptr, size corefoundation.CGSize, format int, colorSpace coregraphics.CGColorSpaceRef) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return CIImageFromID(rv)
}
// Creates and returns an image object from a Quartz 2D image.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// # Return Value
// 
// An image object initialized with the contents of the Quartz 2D image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:
func (_CIImageClass CIImageClass) ImageWithCGImage(image coregraphics.CGImageRef) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCGImage:"), image)
	return CIImageFromID(rv)
}
// Creates and returns an image object from a Quartz 2D image using the
// specified options.
//
// image: A Quartz 2D image ([CGImage]) object. For more information, see [Quartz 2D
// Programming Guide] and [CGImage].
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the contents of the Quartz 2D image and
// the specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:options:
func (_CIImageClass CIImageClass) ImageWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCGImage:options:"), image, options)
	return CIImageFromID(rv)
}
//
// source is a [imageio.CGImageSourceRef].
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImageSource:index:options:
func (_CIImageClass CIImageClass) ImageWithCGImageSourceIndexOptions(source objectivec.IObject, index uintptr, dict foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCGImageSource:index:options:"), source, index, dict)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of [CVImageBuffer]
// object.
//
// imageBuffer: A [CVImageBuffer] object. For more information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// # Return Value
// 
// An image object initialized with the contents of the image buffer object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:
func (_CIImageClass CIImageClass) ImageWithCVImageBuffer(imageBuffer corevideo.CVImageBufferRef) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCVImageBuffer:"), imageBuffer)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of [CVImageBuffer]
// object, using the specified options.
//
// imageBuffer: A [CVImageBuffer] object. For more information, see [Core Video].
// //
// [Core Video]: https://developer.apple.com/documentation/CoreVideo
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the contents of the image buffer object
// and set up with the specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:options:
func (_CIImageClass CIImageClass) ImageWithCVImageBufferOptions(imageBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCVImageBuffer:options:"), imageBuffer, options)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of [CVPixelBuffer]
// object.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// # Return Value
// 
// An image object initialized with the contents of the image buffer object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:
func (_CIImageClass CIImageClass) ImageWithCVPixelBuffer(pixelBuffer corevideo.CVImageBufferRef) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCVPixelBuffer:"), pixelBuffer)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of [CVPixelBuffer]
// object, using the specified options.
//
// pixelBuffer: A [CVPixelBuffer] object.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the contents of the image buffer object
// and set up with the specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:options:
func (_CIImageClass CIImageClass) ImageWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithCVPixelBuffer:options:"), pixelBuffer, options)
	return CIImageFromID(rv)
}
// Creates and returns an image of infinite extent whose entire content is the
// specified color.
//
// color: A color object.
//
// # Return Value
// 
// The image object initialized with the color represented by the [CIColor]
// object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithColor:
func (_CIImageClass CIImageClass) ImageWithColor(color ICIColor) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithColor:"), color)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of a file.
//
// url: The location of the file.
//
// # Return Value
// 
// An image object initialized with the contents of the file.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:
func (_CIImageClass CIImageClass) ImageWithContentsOfURL(url foundation.INSURL) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithContentsOfURL:"), url)
	return CIImageFromID(rv)
}
// Creates and returns an image object from the contents of a file, using the
// specified options.
//
// url: The location of the file.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the contents of the file and set up with
// the specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:options:
func (_CIImageClass CIImageClass) ImageWithContentsOfURLOptions(url foundation.INSURL, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithContentsOfURL:options:"), url, options)
	return CIImageFromID(rv)
}
// Creates and returns an image object initialized with the supplied image
// data.
//
// data: The data object that holds the contents of an image file (such as TIFF,
// GIF, JPG, or whatever else the system supports). The image data must be
// premultiplied.
//
// # Return Value
// 
// An image object initialized with the supplied data, or `nil` if the method
// cannot create an image representation from the contents of the supplied
// data object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:
func (_CIImageClass CIImageClass) ImageWithData(data foundation.INSData) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithData:"), data)
	return CIImageFromID(rv)
}
// Creates and returns an image object initialized with the supplied image
// data, using the specified options.
//
// data: A pointer to the image data. The data must be premultiplied.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the supplied data and set up with the
// specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:options:
func (_CIImageClass CIImageClass) ImageWithDataOptions(data foundation.INSData, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithData:options:"), data, options)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:
func (_CIImageClass CIImageClass) ImageWithDepthData(data objectivec.IObject) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithDepthData:"), data)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:options:
func (_CIImageClass CIImageClass) ImageWithDepthDataOptions(data objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithDepthData:options:"), data, options)
	return CIImageFromID(rv)
}
// Creates and returns an image from the contents of an IOSurface.
//
// surface: An IOSurface object.
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface object.
//
// # Discussion
// 
// An IOSurface object is a framebuffer object that is suitable for sharing
// across process boundaries. You can use it to allow your app to move complex
// image decompression and drawing logic into a separate process for the
// purpose of increasing security.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:
func (_CIImageClass CIImageClass) ImageWithIOSurface(surface iosurface.IOSurfaceRef) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithIOSurface:"), surface)
	return CIImageFromID(rv)
}
// Creates, using the specified options, and returns an image from the
// contents of an IOSurface.
//
// surface: An IOSurface object.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the data from the IOSurface.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:options:
func (_CIImageClass CIImageClass) ImageWithIOSurfaceOptions(surface iosurface.IOSurfaceRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithIOSurface:options:"), surface, options)
	return CIImageFromID(rv)
}
// Create an image object based on pixels from an image provider object.
//
// provider: An object that implements the [CIImageProvider] protocol.
//
// width: The width of the image.
//
// height: The height of the image.
//
// format: The [CIFormat] of the provided pixels.
//
// colorSpace: The color space that the image is defined in. If `nil`, then the pixels
// will not be is not color matched to the Core Image working color space.
//
// options: A dictionary that contains various [CIImageOption] keys that affect the
// resulting [CIImage]. The option [providerTileSize] controls if and how the
// provider object is called in tiles. The option [providerUserInfo] allows
// additional state to be passed to the provider object.
// //
// [providerTileSize]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerTileSize
// [providerUserInfo]: https://developer.apple.com/documentation/CoreImage/CIImageOption/providerUserInfo
//
// # Return Value
// 
// An autoreleased [CIImage] object based on the data provider.
//
// # Discussion
// 
// Core Image retains the provider object until the image is deallocated. The
// image provider object will not be called until the image is rendered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithImageProvider:size::format:colorSpace:options:
func (_CIImageClass CIImageClass) ImageWithImageProviderSizeFormatColorSpaceOptions(provider objectivec.IObject, width uintptr, height uintptr, format int, colorSpace coregraphics.CGColorSpaceRef, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	return CIImageFromID(rv)
}
// Creates and returns an image object with data supplied by a Metal texture.
//
// texture: The Metal texture from which to use image data.
//
// options: A dictionary specifying image options. (See `Image Dictionary Keys`.)
//
// # Return Value
// 
// An image object initialized with the texture data.
//
// # Discussion
// 
// To also render using Metal, use this image with a Metal-based [CIContext]
// object created with the [ContextWithMTLDevice] method, and call the
// [RenderToMTLTextureCommandBufferBoundsColorSpace] method to create an
// output image in another Metal texture object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithMTLTexture:options:
func (_CIImageClass CIImageClass) ImageWithMTLTextureOptions(texture objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithMTLTexture:options:"), texture, options)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:
func (_CIImageClass CIImageClass) ImageWithPortaitEffectsMatte(matte objectivec.IObject) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithPortaitEffectsMatte:"), matte)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:options:
func (_CIImageClass CIImageClass) ImageWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithPortaitEffectsMatte:options:"), matte, options)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:
func (_CIImageClass CIImageClass) ImageWithSemanticSegmentationMatte(matte objectivec.IObject) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithSemanticSegmentationMatte:"), matte)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:options:
func (_CIImageClass CIImageClass) ImageWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.INSDictionary) CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("imageWithSemanticSegmentationMatte:options:"), matte, options)
	return CIImageFromID(rv)
}

// Returns a filter shape object that represents the domain of definition of
// the image.
//
// # Return Value
// 
// A filter shape object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/definition
func (i CIImage) Definition() ICIFilterShape {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("definition"))
	return CIFilterShapeFromID(objc.ID(rv))
}
// A rectangle that specifies the extent of the image.
//
// # Discussion
// 
// This rectangle specifies the extent of the image in working space
// coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/extent
func (i CIImage) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("extent"))
	return corefoundation.CGRect(rv)
}
// Returns the metadata properties dictionary of the image.
//
// # Discussion
// 
// If the [CIImage] was created from [NSURL] or [NSData] then this dictionary
// is determined by calling `CGImageSourceCopyPropertiesAtIndex()`.
// 
// If the [CIImage] was created with the [properties] option, then that
// dictionary is returned.
// 
// If the [CIImage] was created by applying [CIFilter] or [CIKernel] then the
// properties of the root inputImage will be returned.
//
// [properties]: https://developer.apple.com/documentation/CoreImage/CIImageOption/properties
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/properties
func (i CIImage) Properties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("properties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The URL from which the image was loaded.
//
// # Discussion
// 
// A URL is available only if the image object was created with a URL (such as
// with the [InitWithContentsOfURL] method or related methods). Otherwise,
// this property’s value is `nil`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/url
func (i CIImage) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The color space of the image.
//
// # Discussion
// 
// This property’s value is `nil` if the image’s color space cannot be
// determined.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/colorSpace
func (i CIImage) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](i.ID, objc.Sel("colorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}
// The CoreGraphics image object this image was created from, if applicable.
//
// # Discussion
// 
// If this image was created using the [InitWithCGImage] or
// [InitWithContentsOfURL] initializer, this property’s value is the
// [CGImage] object that provides the image’s underlying image data.
// Otherwise, this property’s value is `nil`—in this case you can obtain a
// CoreGraphics image by rendering the image with the [CIContext]
// [CreateCGImageFromRect] method.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/cgImage
func (i CIImage) CGImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](i.ID, objc.Sel("CGImage"))
	return coregraphics.CGImageRef(rv)
}
// The CoreVideo pixel buffer this image was created from, if applicable.
//
// # Discussion
// 
// If this image was create using the [InitWithCVPixelBuffer] initializer,
// this property’s value is the [CVPixelBuffer] object that provides the
// image’s underlying image data. Do not modify the contents of this pixel
// buffer; doing so will cause undefined rendering results.
// 
// Otherwise, this property’s value is `nil`—in this case you can obtain a
// pixel buffer by rendering the image with the [CIContext]
// [RenderToCVPixelBuffer] method.
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/pixelBuffer
func (i CIImage) PixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](i.ID, objc.Sel("pixelBuffer"))
	return corevideo.CVImageBufferRef(rv)
}
// Depth data associated with the image.
//
// # Discussion
// 
// Returns an [AVDepthData] if the [CIImage] was created with [ImageWithData]
// or [ImageWithContentsOfURL] and one of the options [auxiliaryDepth] or
// [auxiliaryDisparity], otherwise [nil].
//
// [AVDepthData]: https://developer.apple.com/documentation/AVFoundation/AVDepthData
// [auxiliaryDepth]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDepth
// [auxiliaryDisparity]: https://developer.apple.com/documentation/CoreImage/CIImageOption/auxiliaryDisparity
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/depthData
func (i CIImage) DepthData() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("depthData"))
	return objectivec.Object{ID: rv}
}
// The portrait effects matte associated with the image.
//
// # Discussion
// 
// [AVPortraitEffectsMatte] representation of portrait effects.
//
// [AVPortraitEffectsMatte]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/portraitEffectsMatte
func (i CIImage) PortraitEffectsMatte() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("portraitEffectsMatte"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/semanticSegmentationMatte
func (i CIImage) SemanticSegmentationMatte() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("semanticSegmentationMatte"))
	return objectivec.Object{ID: rv}
}
// The key for an OpenGL texture target.
//
// See: https://developer.apple.com/documentation/coreimage/ciimageoption/texturetarget
func (i CIImage) TextureTarget() CIImageOption {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("kCIImageTextureTarget"))
	return CIImageOption(foundation.NSStringFromID(rv).String())
}
// The key for an OpenGL texture format.
//
// See: https://developer.apple.com/documentation/coreimage/ciimageoption/textureformat
func (i CIImage) TextureFormat() CIImageOption {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("kCIImageTextureFormat"))
	return CIImageOption(foundation.NSStringFromID(rv).String())
}
// Returns the content headroom of the image.
//
// # Discussion
// 
// If the image headroom is unknown, then the value 0.0 will be returned.
// 
// If the image headroom is known, then a value greater than or equal to 1.0
// will be returned. A value of 1.0 will be returned if the image is SDR. A
// value greater than 1.0 will be returned if the image is HDR.
// 
// The image headroom may known when a CIImage is first initialized. If the a
// CIImage is initialized using:
// 
// - [NSURL] or [NSData] : the headroom may be determined by associated
// metadata or deduced from pixel format or colorSpace information. -
// [CGImage] : headroom may be determined by `CGImageGetHeadroomInfo()` or
// deduced from pixel format or colorSpace information. - [IOSurface] : then
// the headroom will be determined by `kIOSurfaceContentHeadroom`. or deduced
// from pixel format or colorSpace information. - [CVPixelBuffer] : then the
// headroom will be determined by `kCVImageBufferContentLightLevelInfoKey`. or
// deduced from pixel format or colorSpace information. - [BitmapData] :
// headroom may be deduced from pixel format or colorSpace information.
// 
// If the image is the result of applying a [CIFilter] or [CIKernel], this
// method will return `0.0`.
// 
// There are exceptions to this. Applying a `CIWarpKernel`` or certain
// ``CIFilter-class`` (e.G(). `CIGaussianBlur`, `CILanczosScaleTransform`,
// `CIAreaAverage`and some others) to an image will result in a ``CIImage``
// instance with the same`contentHeadroom` property value.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/contentHeadroom
func (i CIImage) ContentHeadroom() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("contentHeadroom"))
	return rv
}
// Returns YES if the image is known to have and alpha value of `1.0` over the
// entire image extent.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/isOpaque
func (i CIImage) Opaque() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isOpaque"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/metalTexture
func (i CIImage) MetalTexture() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("metalTexture"))
	return objectivec.Object{ID: rv}
}
// Returns the content average light level of the image.
//
// # Discussion
// 
// If the image average light level is unknown, then the value 0.0 will be
// returned.
// 
// If the image headroom is known, then a value greater than or equal to 0.0
// will be returned.
// 
// The image average light level may known when a CIImage is first
// initialized. If the a CIImage is initialized with a:
// 
// - [CGImage] : then the headroom will be determined by
// `CGImageGetContentAverageLightLevel()`. - [CVPixelBuffer] : then the
// headroom will be determined by `kCVImageBufferContentLightLevelInfoKey`.
// 
// If the image is the result of applying a [CIFilter] or [CIKernel], this
// property will return `0.0`.
// 
// There are exceptions to this. Applying a [CIWarpKernel] or certain
// [CIFilter] (e.g. [CIGaussianBlur], [CILanczosScaleTransform],
// [CIAreaAverage] and some others) to an image will result in a [CIImage]
// instance with the same `contentAverageLightLevel` property value.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImage/contentAverageLightLevel
func (i CIImage) ContentAverageLightLevel() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("contentAverageLightLevel"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIImage/black
func (_CIImageClass CIImageClass) BlackImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("blackImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/blue
func (_CIImageClass CIImageClass) BlueImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("blueImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/clear
func (_CIImageClass CIImageClass) ClearImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("clearImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/cyan
func (_CIImageClass CIImageClass) CyanImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("cyanImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/gray
func (_CIImageClass CIImageClass) GrayImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("grayImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/green
func (_CIImageClass CIImageClass) GreenImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("greenImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/magenta
func (_CIImageClass CIImageClass) MagentaImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("magentaImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/red
func (_CIImageClass CIImageClass) RedImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("redImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/white
func (_CIImageClass CIImageClass) WhiteImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("whiteImage"))
	return CIImageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreImage/CIImage/yellow
func (_CIImageClass CIImageClass) YellowImage() CIImage {
	rv := objc.Send[objc.ID](objc.ID(_CIImageClass.class), objc.Sel("yellowImage"))
	return CIImageFromID(objc.ID(rv))
}

