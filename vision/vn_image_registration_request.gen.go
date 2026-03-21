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

// The class instance for the [VNImageRegistrationRequest] class.
var (
	_VNImageRegistrationRequestClass     VNImageRegistrationRequestClass
	_VNImageRegistrationRequestClassOnce sync.Once
)

func getVNImageRegistrationRequestClass() VNImageRegistrationRequestClass {
	_VNImageRegistrationRequestClassOnce.Do(func() {
		_VNImageRegistrationRequestClass = VNImageRegistrationRequestClass{class: objc.GetClass("VNImageRegistrationRequest")}
	})
	return _VNImageRegistrationRequestClass
}

// GetVNImageRegistrationRequestClass returns the class object for VNImageRegistrationRequest.
func GetVNImageRegistrationRequestClass() VNImageRegistrationRequestClass {
	return getVNImageRegistrationRequestClass()
}

type VNImageRegistrationRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageRegistrationRequestClass) Alloc() VNImageRegistrationRequest {
	rv := objc.Send[VNImageRegistrationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for image-analysis requests that align images
// according to their content.
//
// # Overview
// 
// This abstract superclass forms the basis of image alignment or registration
// requests. Make specific requests through one of its subclasses,
// [VNTranslationalImageRegistrationRequest] or
// [VNHomographicImageRegistrationRequest]. Don’t create an instance of this
// superclass yourself.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRegistrationRequest
type VNImageRegistrationRequest struct {
	VNTargetedImageRequest
}

// VNImageRegistrationRequestFromID constructs a [VNImageRegistrationRequest] from an objc.ID.
//
// The abstract superclass for image-analysis requests that align images
// according to their content.
func VNImageRegistrationRequestFromID(id objc.ID) VNImageRegistrationRequest {
	return VNImageRegistrationRequest{VNTargetedImageRequest: VNTargetedImageRequestFromID(id)}
}
// NOTE: VNImageRegistrationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNImageRegistrationRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRegistrationRequest
type IVNImageRegistrationRequest interface {
	IVNTargetedImageRequest
}

// Init initializes the instance.
func (i VNImageRegistrationRequest) Init() VNImageRegistrationRequest {
	rv := objc.Send[VNImageRegistrationRequest](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageRegistrationRequest) Autorelease() VNImageRegistrationRequest {
	rv := objc.Send[VNImageRegistrationRequest](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageRegistrationRequest creates a new VNImageRegistrationRequest instance.
func NewVNImageRegistrationRequest() VNImageRegistrationRequest {
	class := getVNImageRegistrationRequestClass()
	rv := objc.Send[VNImageRegistrationRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewImageRegistrationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.CIImage, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.CIImage, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewImageRegistrationRequestWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return VNImageRegistrationRequestFromID(rv)
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
func NewImageRegistrationRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNImageRegistrationRequest {
	instance := getVNImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return VNImageRegistrationRequestFromID(rv)
}

