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

// The class instance for the [VNHomographicImageRegistrationRequest] class.
var (
	_VNHomographicImageRegistrationRequestClass     VNHomographicImageRegistrationRequestClass
	_VNHomographicImageRegistrationRequestClassOnce sync.Once
)

func getVNHomographicImageRegistrationRequestClass() VNHomographicImageRegistrationRequestClass {
	_VNHomographicImageRegistrationRequestClassOnce.Do(func() {
		_VNHomographicImageRegistrationRequestClass = VNHomographicImageRegistrationRequestClass{class: objc.GetClass("VNHomographicImageRegistrationRequest")}
	})
	return _VNHomographicImageRegistrationRequestClass
}

// GetVNHomographicImageRegistrationRequestClass returns the class object for VNHomographicImageRegistrationRequest.
func GetVNHomographicImageRegistrationRequestClass() VNHomographicImageRegistrationRequestClass {
	return getVNHomographicImageRegistrationRequestClass()
}

type VNHomographicImageRegistrationRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNHomographicImageRegistrationRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHomographicImageRegistrationRequestClass) Alloc() VNHomographicImageRegistrationRequest {
	rv := objc.Send[VNHomographicImageRegistrationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that determines the perspective warp matrix
// necessary to align the content of two images.
//
// # Overview
// 
// Create and perform a homographic image registration request to align
// content in two images through a homography. A is an isomorphism of
// projected spaces, a bijection that maps lines to lines.
//
// # Identifying Request Revisions
//
//   - [VNHomographicImageRegistrationRequest.VNHomographicImageRegistrationRequestRevision1]: A constant for specifying revision 1 of the homographic image registration request.
//
// See: https://developer.apple.com/documentation/Vision/VNHomographicImageRegistrationRequest
type VNHomographicImageRegistrationRequest struct {
	VNImageRegistrationRequest
}

// VNHomographicImageRegistrationRequestFromID constructs a [VNHomographicImageRegistrationRequest] from an objc.ID.
//
// An image-analysis request that determines the perspective warp matrix
// necessary to align the content of two images.
func VNHomographicImageRegistrationRequestFromID(id objc.ID) VNHomographicImageRegistrationRequest {
	return VNHomographicImageRegistrationRequest{VNImageRegistrationRequest: VNImageRegistrationRequestFromID(id)}
}
// NOTE: VNHomographicImageRegistrationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHomographicImageRegistrationRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNHomographicImageRegistrationRequest.VNHomographicImageRegistrationRequestRevision1]: A constant for specifying revision 1 of the homographic image registration request.
//
// See: https://developer.apple.com/documentation/Vision/VNHomographicImageRegistrationRequest
type IVNHomographicImageRegistrationRequest interface {
	IVNImageRegistrationRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the homographic image registration request.
	VNHomographicImageRegistrationRequestRevision1() int
}

// Init initializes the instance.
func (h VNHomographicImageRegistrationRequest) Init() VNHomographicImageRegistrationRequest {
	rv := objc.Send[VNHomographicImageRegistrationRequest](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHomographicImageRegistrationRequest) Autorelease() VNHomographicImageRegistrationRequest {
	rv := objc.Send[VNHomographicImageRegistrationRequest](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHomographicImageRegistrationRequest creates a new VNHomographicImageRegistrationRequest instance.
func NewVNHomographicImageRegistrationRequest() VNHomographicImageRegistrationRequest {
	class := getVNHomographicImageRegistrationRequestClass()
	rv := objc.Send[VNHomographicImageRegistrationRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewHomographicImageRegistrationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewHomographicImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.CIImage, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.CIImage, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.CIImage, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewHomographicImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewHomographicImageRegistrationRequestWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
}

// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewHomographicImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return VNHomographicImageRegistrationRequestFromID(rv)
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
func NewHomographicImageRegistrationRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNHomographicImageRegistrationRequest {
	instance := getVNHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return VNHomographicImageRegistrationRequestFromID(rv)
}

// A constant for specifying revision 1 of the homographic image registration
// request.
//
// See: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequestrevision1
func (h VNHomographicImageRegistrationRequest) VNHomographicImageRegistrationRequestRevision1() int {
	rv := objc.Send[int](h.ID, objc.Sel("VNHomographicImageRegistrationRequestRevision1"))
	return rv
}

