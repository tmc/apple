// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNTargetedImageRequest] class.
var (
	_VNTargetedImageRequestClass     VNTargetedImageRequestClass
	_VNTargetedImageRequestClassOnce sync.Once
)

func getVNTargetedImageRequestClass() VNTargetedImageRequestClass {
	_VNTargetedImageRequestClassOnce.Do(func() {
		_VNTargetedImageRequestClass = VNTargetedImageRequestClass{class: objc.GetClass("VNTargetedImageRequest")}
	})
	return _VNTargetedImageRequestClass
}

// GetVNTargetedImageRequestClass returns the class object for VNTargetedImageRequest.
func GetVNTargetedImageRequestClass() VNTargetedImageRequestClass {
	return getVNTargetedImageRequestClass()
}

type VNTargetedImageRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNTargetedImageRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTargetedImageRequestClass) Alloc() VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for image analysis requests that operate on both
// the processed image and a secondary image.
//
// # Overview
// 
// Other Vision request handlers that operate on both the processed image and
// a secondary image inherit from this abstract base class. Instantiate one of
// its subclasses to perform image analysis, and pass in auxiliary image data
// by filling in the `options` dictionary at initialization.
//
// # Creating a Request
//
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOptionsCompletionHandler]: Creates a new request targeting a Core Graphics image, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]: Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOptionsCompletionHandler]: Creates a new request targeting a Core Image image.
//   - [VNTargetedImageRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]: Creates a new request targeting a Core Image image of known orientation, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]: Creates a new request targeting an image in a pixel buffer.
//   - [VNTargetedImageRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]: Creates a new request targeting an image in a pixel buffer of known orientation.
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]: Creates a new request with a completion handler that targets an image in a sample buffer.
//   - [VNTargetedImageRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]: Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer.
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOptionsCompletionHandler]: Creates a new request targeting an image as raw data, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]: Creates a new request targeting a raw data image of known orientation, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOptionsCompletionHandler]: Creates a new request targeting an image at the specified URL, executing the completion handler when done.
//   - [VNTargetedImageRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]: Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest
type VNTargetedImageRequest struct {
	VNImageBasedRequest
}

// VNTargetedImageRequestFromID constructs a [VNTargetedImageRequest] from an objc.ID.
//
// The abstract superclass for image analysis requests that operate on both
// the processed image and a secondary image.
func VNTargetedImageRequestFromID(id objc.ID) VNTargetedImageRequest {
	return VNTargetedImageRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNTargetedImageRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTargetedImageRequest] class.
//
// # Creating a Request
//
//   - [IVNTargetedImageRequest.InitWithTargetedCGImageOptionsCompletionHandler]: Creates a new request targeting a Core Graphics image, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedCGImageOrientationOptionsCompletionHandler]: Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedCIImageOptionsCompletionHandler]: Creates a new request targeting a Core Image image.
//   - [IVNTargetedImageRequest.InitWithTargetedCIImageOrientationOptionsCompletionHandler]: Creates a new request targeting a Core Image image of known orientation, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedCVPixelBufferOptionsCompletionHandler]: Creates a new request targeting an image in a pixel buffer.
//   - [IVNTargetedImageRequest.InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler]: Creates a new request targeting an image in a pixel buffer of known orientation.
//   - [IVNTargetedImageRequest.InitWithTargetedCMSampleBufferOptionsCompletionHandler]: Creates a new request with a completion handler that targets an image in a sample buffer.
//   - [IVNTargetedImageRequest.InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler]: Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer.
//   - [IVNTargetedImageRequest.InitWithTargetedImageDataOptionsCompletionHandler]: Creates a new request targeting an image as raw data, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedImageDataOrientationOptionsCompletionHandler]: Creates a new request targeting a raw data image of known orientation, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedImageURLOptionsCompletionHandler]: Creates a new request targeting an image at the specified URL, executing the completion handler when done.
//   - [IVNTargetedImageRequest.InitWithTargetedImageURLOrientationOptionsCompletionHandler]: Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest
type IVNTargetedImageRequest interface {
	IVNImageBasedRequest

	// Topic: Creating a Request

	// Creates a new request targeting a Core Graphics image, executing the completion handler when done.
	InitWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done.
	InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting a Core Image image.
	InitWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.CIImage, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting a Core Image image of known orientation, executing the completion handler when done.
	InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting an image in a pixel buffer.
	InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting an image in a pixel buffer of known orientation.
	InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request with a completion handler that targets an image in a sample buffer.
	InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer.
	InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting an image as raw data, executing the completion handler when done.
	InitWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting a raw data image of known orientation, executing the completion handler when done.
	InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting an image at the specified URL, executing the completion handler when done.
	InitWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest
	// Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done.
	InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest

	// Creates a new request targeting a Core Graphics image.
	InitWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting a Core Graphics image of known orientation.
	InitWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting a Core Image image.
	InitWithTargetedCIImageOptions(ciImage coreimage.CIImage, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting a Core Image image of known orientation.
	InitWithTargetedCIImageOrientationOptions(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request that targets an image in a sample buffer.
	InitWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request that targets an image of a known orientation in a sample buffer.
	InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting an image in a pixel buffer.
	InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting an image in a pixel buffer of known orientation.
	InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting an image as raw data.
	InitWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting a raw data image of known orientation.
	InitWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting an image at the specified URL.
	InitWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNTargetedImageRequest
	// Creates a new request targeting an image of known orientation, at the specified URL.
	InitWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest
}

// Init initializes the instance.
func (t VNTargetedImageRequest) Init() VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTargetedImageRequest) Autorelease() VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTargetedImageRequest creates a new VNTargetedImageRequest instance.
func NewVNTargetedImageRequest() VNTargetedImageRequest {
	class := getVNTargetedImageRequestClass()
	rv := objc.Send[VNTargetedImageRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new Vision request with an optional completion handler.
//
// completionHandler: The block to invoke after the request finishes processing.
//
// # Discussion
// 
// Vision executes the completion handler on the same queue that it executes
// the request; however, this queue differs from the one where you called
// [PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewTargetedImageRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewTargetedImageRequestWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image, executing the
// completion handler when done.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image of known orientation.
//
// cgImage: The targeted Core Graphics image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image of known orientation,
// executing the completion handler when done.
//
// cgImage: The targeted Core Graphics image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Image image.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:options:
func NewTargetedImageRequestWithTargetedCIImageOptions(ciImage coreimage.CIImage, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Image image.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.CIImage, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Image image of known orientation.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Image image of known orientation,
// executing the completion handler when done.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request that targets an image in a sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:options:
// sampleBuffer is a [coremedia.CMSampleBufferRef].
func NewTargetedImageRequestWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request with a completion handler that targets an image in a
// sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The callback the system invokes when the request finishes executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:options:completionHandler:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
func NewTargetedImageRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request that targets an image of a known orientation in a
// sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// orientation: The EXIF orientation of the image. See [CGImagePropertyOrientation] for
// supported orientation values.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:orientation:options:
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request with a completion handler that targets an image of a
// known orientation in a sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// orientation: The EXIF orientation of the image. See [CGImagePropertyOrientation] for
// supported orientations.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The callback the system invokes when the request finishes executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:orientation:options:completionHandler:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewTargetedImageRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer of known
// orientation.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer of known
// orientation.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewTargetedImageRequestWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image as raw data, executing the
// completion handler when done.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a raw data image of known orientation.
//
// imageData: The data containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a raw data image of known orientation,
// executing the completion handler when done.
//
// imageData: The data containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewTargetedImageRequestWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image at the specified URL, executing
// the completion handler when done.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image of known orientation, at the
// specified URL.
//
// imageURL: The URL of the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting an image of known orientation, at the
// specified URL, executing the completion handler when done.
//
// imageURL: The URL of the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTargetedImageRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTargetedImageRequest {
	instance := getVNTargetedImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return VNTargetedImageRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image, executing the
// completion handler when done.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:options:completionHandler:)
func (t VNTargetedImageRequest) InitWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting a Core Graphics image of known orientation,
// executing the completion handler when done.
//
// cgImage: The targeted Core Graphics image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting a Core Image image.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:options:completionHandler:)
func (t VNTargetedImageRequest) InitWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.CIImage, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting a Core Image image of known orientation,
// executing the completion handler when done.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:options:completionHandler:)
func (t VNTargetedImageRequest) InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting an image in a pixel buffer of known
// orientation.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request with a completion handler that targets an image in a
// sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The callback the system invokes when the request finishes executing.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:options:completionHandler:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
func (t VNTargetedImageRequest) InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request with a completion handler that targets an image of a
// known orientation in a sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// orientation: The EXIF orientation of the image. See [CGImagePropertyOrientation] for
// supported orientations.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The callback the system invokes when the request finishes executing.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:orientation:options:completionHandler:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting an image as raw data, executing the
// completion handler when done.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:options:completionHandler:)
func (t VNTargetedImageRequest) InitWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting a raw data image of known orientation,
// executing the completion handler when done.
//
// imageData: The data containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting an image at the specified URL, executing
// the completion handler when done.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:options:completionHandler:)
func (t VNTargetedImageRequest) InitWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, _block2)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting an image of known orientation, at the
// specified URL, executing the completion handler when done.
//
// imageURL: The URL of the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// completionHandler: The block to invoke when the request has finished executing.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:orientation:options:completionHandler:)
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) VNTargetedImageRequest {
_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, _block3)
	return VNTargetedImageRequestFromID(rv)
}
// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func (t VNTargetedImageRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}
// Creates a new request targeting a Core Graphics image of known orientation.
//
// cgImage: The targeted Core Graphics image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return rv
}
// Creates a new request targeting a Core Image image.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:options:
func (t VNTargetedImageRequest) InitWithTargetedCIImageOptions(ciImage coreimage.CIImage, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return rv
}
// Creates a new request targeting a Core Image image of known orientation.
//
// ciImage: The [CIImage] encapsulating the targeted image.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCIImageOrientationOptions(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return rv
}
// Creates a new request that targets an image in a sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:options:
// sampleBuffer is a [coremedia.CMSampleBufferRef].
func (t VNTargetedImageRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}
// Creates a new request that targets an image of a known orientation in a
// sample buffer.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// orientation: The EXIF orientation of the image. See [CGImagePropertyOrientation] for
// supported orientation values.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:orientation:options:
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return rv
}
// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func (t VNTargetedImageRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}
// Creates a new request targeting an image in a pixel buffer of known
// orientation.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return rv
}
// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func (t VNTargetedImageRequest) InitWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}
// Creates a new request targeting a raw data image of known orientation.
//
// imageData: The data containing the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return rv
}
// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func (t VNTargetedImageRequest) InitWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return rv
}
// Creates a new request targeting an image of known orientation, at the
// specified URL.
//
// imageURL: The URL of the targeted image.
//
// orientation: The orientation of the image buffer, based on EXIF specification and
// superseding other orientation information. The value must be an integer
// from `1` to `8`; see [kCGImagePropertyOrientation] for details.
// //
// [kCGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyOrientation
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:orientation:options:
// orientation is a [imageio.CGImagePropertyOrientation].
func (t VNTargetedImageRequest) InitWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNTargetedImageRequest {
	rv := objc.Send[VNTargetedImageRequest](t.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return rv
}

