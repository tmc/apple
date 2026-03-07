// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNTranslationalImageRegistrationRequest] class.
var (
	_VNTranslationalImageRegistrationRequestClass     VNTranslationalImageRegistrationRequestClass
	_VNTranslationalImageRegistrationRequestClassOnce sync.Once
)

func getVNTranslationalImageRegistrationRequestClass() VNTranslationalImageRegistrationRequestClass {
	_VNTranslationalImageRegistrationRequestClassOnce.Do(func() {
		_VNTranslationalImageRegistrationRequestClass = VNTranslationalImageRegistrationRequestClass{class: objc.GetClass("VNTranslationalImageRegistrationRequest")}
	})
	return _VNTranslationalImageRegistrationRequestClass
}

// GetVNTranslationalImageRegistrationRequestClass returns the class object for VNTranslationalImageRegistrationRequest.
func GetVNTranslationalImageRegistrationRequestClass() VNTranslationalImageRegistrationRequestClass {
	return getVNTranslationalImageRegistrationRequestClass()
}

type VNTranslationalImageRegistrationRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTranslationalImageRegistrationRequestClass) Alloc() VNTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTranslationalImageRegistrationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An image-analysis request that determines the affine transform necessary to
// align the content of two images.
//
// # Overview
// 
// Create and perform a translational image registration request to align
// content in two images through translation.
//
// See: https://developer.apple.com/documentation/Vision/VNTranslationalImageRegistrationRequest
type VNTranslationalImageRegistrationRequest struct {
	VNImageRegistrationRequest
}

// VNTranslationalImageRegistrationRequestFromID constructs a [VNTranslationalImageRegistrationRequest] from an objc.ID.
//
// An image-analysis request that determines the affine transform necessary to
// align the content of two images.
func VNTranslationalImageRegistrationRequestFromID(id objc.ID) VNTranslationalImageRegistrationRequest {
	return VNTranslationalImageRegistrationRequest{VNImageRegistrationRequest: VNImageRegistrationRequestFromID(id)}
}
// NOTE: VNTranslationalImageRegistrationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNTranslationalImageRegistrationRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNTranslationalImageRegistrationRequest
type IVNTranslationalImageRegistrationRequest interface {
	IVNImageRegistrationRequest
}





// Init initializes the instance.
func (t VNTranslationalImageRegistrationRequest) Init() VNTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTranslationalImageRegistrationRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTranslationalImageRegistrationRequest) Autorelease() VNTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTranslationalImageRegistrationRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTranslationalImageRegistrationRequest creates a new VNTranslationalImageRegistrationRequest instance.
func NewVNTranslationalImageRegistrationRequest() VNTranslationalImageRegistrationRequest {
	class := getVNTranslationalImageRegistrationRequestClass()
	rv := objc.Send[VNTranslationalImageRegistrationRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
}


// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
// ciImage is a [coreimage.CIImage].
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOptions(ciImage objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
// ciImage is a [coreimage.CIImage].
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOptionsCompletionHandler(ciImage objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
// ciImage is a [coreimage.CIImage].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOrientationOptions(ciImage objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
// ciImage is a [coreimage.CIImage].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
}


// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
}


// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
}


// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return VNTranslationalImageRegistrationRequestFromID(rv)
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
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNTranslationalImageRegistrationRequest {
	instance := getVNTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return VNTranslationalImageRegistrationRequestFromID(rv)
}







































