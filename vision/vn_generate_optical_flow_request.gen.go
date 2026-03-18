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

// The class instance for the [VNGenerateOpticalFlowRequest] class.
var (
	_VNGenerateOpticalFlowRequestClass     VNGenerateOpticalFlowRequestClass
	_VNGenerateOpticalFlowRequestClassOnce sync.Once
)

func getVNGenerateOpticalFlowRequestClass() VNGenerateOpticalFlowRequestClass {
	_VNGenerateOpticalFlowRequestClassOnce.Do(func() {
		_VNGenerateOpticalFlowRequestClass = VNGenerateOpticalFlowRequestClass{class: objc.GetClass("VNGenerateOpticalFlowRequest")}
	})
	return _VNGenerateOpticalFlowRequestClass
}

// GetVNGenerateOpticalFlowRequestClass returns the class object for VNGenerateOpticalFlowRequest.
func GetVNGenerateOpticalFlowRequestClass() VNGenerateOpticalFlowRequestClass {
	return getVNGenerateOpticalFlowRequestClass()
}

type VNGenerateOpticalFlowRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNGenerateOpticalFlowRequestClass) Alloc() VNGenerateOpticalFlowRequest {
	rv := objc.Send[VNGenerateOpticalFlowRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that generates directional change vectors for each pixel in the
// targeted image.
//
// # Overview
// 
// This request operates at a pixel level, so both images need to have the
// same dimensions to successfully perform the analysis. Setting a region of
// interest limits the region in which the analysis occurs. However, the
// system reports the resulting observation at full resolution.
// 
// Optical flow requests are resource-intensive, so create only one request at
// a time, and release it immediately after generating optical flows.
//
// # Configuring the Request
//
//   - [VNGenerateOpticalFlowRequest.ComputationAccuracy]: The accuracy level for computing optical flow.
//   - [VNGenerateOpticalFlowRequest.SetComputationAccuracy]
//   - [VNGenerateOpticalFlowRequest.OutputPixelFormat]: The output buffer’s pixel format.
//   - [VNGenerateOpticalFlowRequest.SetOutputPixelFormat]
//   - [VNGenerateOpticalFlowRequest.KeepNetworkOutput]: A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
//   - [VNGenerateOpticalFlowRequest.SetKeepNetworkOutput]
//
// # Identifying Request Revisions
//
//   - [VNGenerateOpticalFlowRequest.VNGenerateOpticalFlowRequestRevision2]: A constant for specifying revision 2 of the optical flow generation request.
//   - [VNGenerateOpticalFlowRequest.VNGenerateOpticalFlowRequestRevision1]: A constant for specifying revision 1 of the optical flow generation request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest
type VNGenerateOpticalFlowRequest struct {
	VNTargetedImageRequest
}

// VNGenerateOpticalFlowRequestFromID constructs a [VNGenerateOpticalFlowRequest] from an objc.ID.
//
// An object that generates directional change vectors for each pixel in the
// targeted image.
func VNGenerateOpticalFlowRequestFromID(id objc.ID) VNGenerateOpticalFlowRequest {
	return VNGenerateOpticalFlowRequest{VNTargetedImageRequest: VNTargetedImageRequestFromID(id)}
}
// NOTE: VNGenerateOpticalFlowRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNGenerateOpticalFlowRequest] class.
//
// # Configuring the Request
//
//   - [IVNGenerateOpticalFlowRequest.ComputationAccuracy]: The accuracy level for computing optical flow.
//   - [IVNGenerateOpticalFlowRequest.SetComputationAccuracy]
//   - [IVNGenerateOpticalFlowRequest.OutputPixelFormat]: The output buffer’s pixel format.
//   - [IVNGenerateOpticalFlowRequest.SetOutputPixelFormat]
//   - [IVNGenerateOpticalFlowRequest.KeepNetworkOutput]: A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
//   - [IVNGenerateOpticalFlowRequest.SetKeepNetworkOutput]
//
// # Identifying Request Revisions
//
//   - [IVNGenerateOpticalFlowRequest.VNGenerateOpticalFlowRequestRevision2]: A constant for specifying revision 2 of the optical flow generation request.
//   - [IVNGenerateOpticalFlowRequest.VNGenerateOpticalFlowRequestRevision1]: A constant for specifying revision 1 of the optical flow generation request.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest
type IVNGenerateOpticalFlowRequest interface {
	IVNTargetedImageRequest

	// Topic: Configuring the Request

	// The accuracy level for computing optical flow.
	ComputationAccuracy() VNGenerateOpticalFlowRequestComputationAccuracy
	SetComputationAccuracy(value VNGenerateOpticalFlowRequestComputationAccuracy)
	// The output buffer’s pixel format.
	OutputPixelFormat() uint32
	SetOutputPixelFormat(value uint32)
	// A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
	KeepNetworkOutput() bool
	SetKeepNetworkOutput(value bool)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 2 of the optical flow generation request.
	VNGenerateOpticalFlowRequestRevision2() int
	// A constant for specifying revision 1 of the optical flow generation request.
	VNGenerateOpticalFlowRequestRevision1() int
}

// Init initializes the instance.
func (g VNGenerateOpticalFlowRequest) Init() VNGenerateOpticalFlowRequest {
	rv := objc.Send[VNGenerateOpticalFlowRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VNGenerateOpticalFlowRequest) Autorelease() VNGenerateOpticalFlowRequest {
	rv := objc.Send[VNGenerateOpticalFlowRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNGenerateOpticalFlowRequest creates a new VNGenerateOpticalFlowRequest instance.
func NewVNGenerateOpticalFlowRequest() VNGenerateOpticalFlowRequest {
	class := getVNGenerateOpticalFlowRequestClass()
	rv := objc.Send[VNGenerateOpticalFlowRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewGenerateOpticalFlowRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
}

// Creates a new request targeting a Core Graphics image.
//
// cgImage: The targeted Core Graphics image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewGenerateOpticalFlowRequestWithTargetedCGImageOptions(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCIImageOptions(ciImage objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCIImageOptionsCompletionHandler(ciImage objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCIImageOrientationOptions(ciImage objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
}

// Creates a new request targeting an image in a pixel buffer.
//
// pixelBuffer: The pixel buffer containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewGenerateOpticalFlowRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
}

// Creates a new request targeting an image as raw data.
//
// imageData: The data containing the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewGenerateOpticalFlowRequestWithTargetedImageDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.INSData, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
}

// Creates a new request targeting an image at the specified URL.
//
// imageURL: The URL of the targeted image.
//
// options: A dictionary with options specifying auxiliary information for the image.
//
// See: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewGenerateOpticalFlowRequestWithTargetedImageURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.INSURL, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return VNGenerateOpticalFlowRequestFromID(rv)
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
func NewGenerateOpticalFlowRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary, completionHandler VNRequestCompletionHandler) VNGenerateOpticalFlowRequest {
	instance := getVNGenerateOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return VNGenerateOpticalFlowRequestFromID(rv)
}

// The accuracy level for computing optical flow.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/computationAccuracy-swift.property
func (g VNGenerateOpticalFlowRequest) ComputationAccuracy() VNGenerateOpticalFlowRequestComputationAccuracy {
	rv := objc.Send[VNGenerateOpticalFlowRequestComputationAccuracy](g.ID, objc.Sel("computationAccuracy"))
	return VNGenerateOpticalFlowRequestComputationAccuracy(rv)
}
func (g VNGenerateOpticalFlowRequest) SetComputationAccuracy(value VNGenerateOpticalFlowRequestComputationAccuracy) {
	objc.Send[struct{}](g.ID, objc.Sel("setComputationAccuracy:"), value)
}

// The output buffer’s pixel format.
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/outputPixelFormat
func (g VNGenerateOpticalFlowRequest) OutputPixelFormat() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("outputPixelFormat"))
	return rv
}
func (g VNGenerateOpticalFlowRequest) SetOutputPixelFormat(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputPixelFormat:"), value)
}

// A Boolean value that indicates whether to keep the raw pixel buffer coming
// from the machine learning network.
//
// # Discussion
// 
// The default is [false]. When you set this to [true], the system ignores
// [OutputPixelFormat]. Setting this for revision 1 has no effect because
// it’s not machine learning-based.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/keepNetworkOutput
func (g VNGenerateOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("keepNetworkOutput"))
	return rv
}
func (g VNGenerateOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setKeepNetworkOutput:"), value)
}

// A constant for specifying revision 2 of the optical flow generation
// request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision2
func (g VNGenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision2() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision2"))
	return rv
}

// A constant for specifying revision 1 of the optical flow generation
// request.
//
// See: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision1
func (g VNGenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision1() int {
	rv := objc.Send[int](g.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision1"))
	return rv
}

