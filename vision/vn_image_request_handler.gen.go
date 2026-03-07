// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNImageRequestHandler] class.
var (
	_VNImageRequestHandlerClass     VNImageRequestHandlerClass
	_VNImageRequestHandlerClassOnce sync.Once
)

func getVNImageRequestHandlerClass() VNImageRequestHandlerClass {
	_VNImageRequestHandlerClassOnce.Do(func() {
		_VNImageRequestHandlerClass = VNImageRequestHandlerClass{class: objc.GetClass("VNImageRequestHandler")}
	})
	return _VNImageRequestHandlerClass
}

// GetVNImageRequestHandlerClass returns the class object for VNImageRequestHandler.
func GetVNImageRequestHandlerClass() VNImageRequestHandlerClass {
	return getVNImageRequestHandlerClass()
}

type VNImageRequestHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageRequestHandlerClass) Alloc() VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that processes one or more image-analysis request pertaining to a
// single image.
//
// # Overview
// 
// Instantiate this handler to perform Vision requests on a single image. You
// specify the image and, optionally, a completion handler at the time of
// creation, and call [VNImageRequestHandler.PerformRequestsError] to begin executing the request.
//
// # Creating a Request Handler
//
//   - [VNImageRequestHandler.InitWithCGImageOptions]: Creates a handler to be used for performing requests on Core Graphics images.
//   - [VNImageRequestHandler.InitWithCGImageOrientationOptions]: Creates a handler to be used for performing requests on a Core Graphics image with known orientation.
//   - [VNImageRequestHandler.InitWithCIImageOptions]: Creates a handler to use for performing requests on Core Image image data.
//   - [VNImageRequestHandler.InitWithCIImageOrientationOptions]: Creates a handler to be used for performing requests on Core Image image data of a known orientation.
//   - [VNImageRequestHandler.InitWithCVPixelBufferOptions]: Creates a handler for performing requests on a Core Video pixel buffer.
//   - [VNImageRequestHandler.InitWithCVPixelBufferOrientationOptions]: Creates a handler for performing requests on a Core Video pixel buffer of a known orientation.
//   - [VNImageRequestHandler.InitWithCVPixelBufferDepthDataOrientationOptions]
//   - [VNImageRequestHandler.InitWithCMSampleBufferOptions]: Creates a request handler that performs requests on an image contained within a sample buffer.
//   - [VNImageRequestHandler.InitWithCMSampleBufferOrientationOptions]: Creates a request handler that performs requests on an image of a specified orientation contained within a sample buffer.
//   - [VNImageRequestHandler.InitWithCMSampleBufferDepthDataOrientationOptions]: Creates a request handler that performs requests on an image in a sample buffer that contains depth data.
//   - [VNImageRequestHandler.InitWithDataOptions]: Creates a handler to use for performing requests on an image in a data object.
//   - [VNImageRequestHandler.InitWithDataOrientationOptions]: Creates a handler to use for performing requests on an image of known orientation.
//   - [VNImageRequestHandler.InitWithURLOptions]: Creates a handler to be used for performing requests on an image at the specified URL.
//   - [VNImageRequestHandler.InitWithURLOrientationOptions]: Creates a handler to be used for performing requests on an image with known orientation, at the specified URL.
//
// # Executing a Request Handler
//
//   - [VNImageRequestHandler.PerformRequestsError]: Schedules Vision requests to perform.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler
type VNImageRequestHandler struct {
	objectivec.Object
}

// VNImageRequestHandlerFromID constructs a [VNImageRequestHandler] from an objc.ID.
//
// An object that processes one or more image-analysis request pertaining to a
// single image.
func VNImageRequestHandlerFromID(id objc.ID) VNImageRequestHandler {
	return VNImageRequestHandler{objectivec.Object{id}}
}
// NOTE: VNImageRequestHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNImageRequestHandler] class.
//
// # Creating a Request Handler
//
//   - [IVNImageRequestHandler.InitWithCGImageOptions]: Creates a handler to be used for performing requests on Core Graphics images.
//   - [IVNImageRequestHandler.InitWithCGImageOrientationOptions]: Creates a handler to be used for performing requests on a Core Graphics image with known orientation.
//   - [IVNImageRequestHandler.InitWithCIImageOptions]: Creates a handler to use for performing requests on Core Image image data.
//   - [IVNImageRequestHandler.InitWithCIImageOrientationOptions]: Creates a handler to be used for performing requests on Core Image image data of a known orientation.
//   - [IVNImageRequestHandler.InitWithCVPixelBufferOptions]: Creates a handler for performing requests on a Core Video pixel buffer.
//   - [IVNImageRequestHandler.InitWithCVPixelBufferOrientationOptions]: Creates a handler for performing requests on a Core Video pixel buffer of a known orientation.
//   - [IVNImageRequestHandler.InitWithCVPixelBufferDepthDataOrientationOptions]
//   - [IVNImageRequestHandler.InitWithCMSampleBufferOptions]: Creates a request handler that performs requests on an image contained within a sample buffer.
//   - [IVNImageRequestHandler.InitWithCMSampleBufferOrientationOptions]: Creates a request handler that performs requests on an image of a specified orientation contained within a sample buffer.
//   - [IVNImageRequestHandler.InitWithCMSampleBufferDepthDataOrientationOptions]: Creates a request handler that performs requests on an image in a sample buffer that contains depth data.
//   - [IVNImageRequestHandler.InitWithDataOptions]: Creates a handler to use for performing requests on an image in a data object.
//   - [IVNImageRequestHandler.InitWithDataOrientationOptions]: Creates a handler to use for performing requests on an image of known orientation.
//   - [IVNImageRequestHandler.InitWithURLOptions]: Creates a handler to be used for performing requests on an image at the specified URL.
//   - [IVNImageRequestHandler.InitWithURLOrientationOptions]: Creates a handler to be used for performing requests on an image with known orientation, at the specified URL.
//
// # Executing a Request Handler
//
//   - [IVNImageRequestHandler.PerformRequestsError]: Schedules Vision requests to perform.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler
type IVNImageRequestHandler interface {
	objectivec.IObject

	// Topic: Creating a Request Handler

	// Creates a handler to be used for performing requests on Core Graphics images.
	InitWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to be used for performing requests on a Core Graphics image with known orientation.
	InitWithCGImageOrientationOptions(image coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to use for performing requests on Core Image image data.
	InitWithCIImageOptions(image objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to be used for performing requests on Core Image image data of a known orientation.
	InitWithCIImageOrientationOptions(image objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler for performing requests on a Core Video pixel buffer.
	InitWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler for performing requests on a Core Video pixel buffer of a known orientation.
	InitWithCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	InitWithCVPixelBufferDepthDataOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a request handler that performs requests on an image contained within a sample buffer.
	InitWithCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a request handler that performs requests on an image of a specified orientation contained within a sample buffer.
	InitWithCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a request handler that performs requests on an image in a sample buffer that contains depth data.
	InitWithCMSampleBufferDepthDataOrientationOptions(sampleBuffer objectivec.IObject, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to use for performing requests on an image in a data object.
	InitWithDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to use for performing requests on an image of known orientation.
	InitWithDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to be used for performing requests on an image at the specified URL.
	InitWithURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNImageRequestHandler
	// Creates a handler to be used for performing requests on an image with known orientation, at the specified URL.
	InitWithURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler

	// Topic: Executing a Request Handler

	// Schedules Vision requests to perform.
	PerformRequestsError(requests []VNRequest) (bool, error)
}





// Init initializes the instance.
func (i VNImageRequestHandler) Init() VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageRequestHandler) Autorelease() VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageRequestHandler creates a new VNImageRequestHandler instance.
func NewVNImageRequestHandler() VNImageRequestHandler {
	class := getVNImageRequestHandlerClass()
	rv := objc.Send[VNImageRequestHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a handler to be used for performing requests on Core Graphics
// images.
//
// image: A [CGImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:options:)
func NewImageRequestHandlerWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:options:"), image, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to be used for performing requests on a Core Graphics
// image with known orientation.
//
// image: A [CGImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:orientation:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCGImageOrientationOptions(image coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:orientation:options:"), image, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to use for performing requests on Core Image image data.
//
// image: A [CIImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: An optional dictionary containing [properties] keys to auxiliary image
// data.
// //
// [properties]: https://developer.apple.com/documentation/Vision/VNImageOption/properties
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:options:)
// image is a [coreimage.CIImage].
func NewImageRequestHandlerWithCIImageOptions(image objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCIImage:options:"), image, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to be used for performing requests on Core Image image
// data of a known orientation.
//
// image: A [CIImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:orientation:options:)
// image is a [coreimage.CIImage].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCIImageOrientationOptions(image objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCIImage:orientation:options:"), image, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a request handler that performs requests on an image in a sample
// buffer that contains depth data.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// depthData: The depth data to associate with the image.
//
// orientation: The EXIF orientation of the image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:depthData:orientation:options:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// depthData is a [avfoundation.AVDepthData].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCMSampleBufferDepthDataOrientationOptions(sampleBuffer objectivec.IObject, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCMSampleBuffer:depthData:orientation:options:"), sampleBuffer, depthData, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a request handler that performs requests on an image contained
// within a sample buffer.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:options:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
func NewImageRequestHandlerWithCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCMSampleBuffer:options:"), sampleBuffer, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a request handler that performs requests on an image of a specified
// orientation contained within a sample buffer.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// orientation: The EXIF orientation of the image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:orientation:options:)
// sampleBuffer is a [coremedia.CMSampleBufferRef].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:depthData:orientation:options:)
// depthData is a [avfoundation.AVDepthData].
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCVPixelBufferDepthDataOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:depthData:orientation:options:"), pixelBuffer, depthData, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler for performing requests on a Core Video pixel buffer.
//
// pixelBuffer: A pixel buffer that contains the image to use for performing the requests.
// The contents can’t change for the lifetime of the request handler.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:options:)
func NewImageRequestHandlerWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler for performing requests on a Core Video pixel buffer of a
// known orientation.
//
// pixelBuffer: A pixel buffer that contains the image to use for performing the requests.
// The contents can’t change for the lifetime of the request handler.
//
// orientation: The orientation of the input image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:orientation:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to use for performing requests on an image in a data
// object.
//
// imageData: Data containing the image to be used for performing the requests. Image
// content is immutable.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// # Discussion
// 
// The intended use cases of this type of initializer include compressed
// images and network downloads, where a client may receive a JPEG from a
// website or the cloud.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:options:)
func NewImageRequestHandlerWithDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:"), imageData, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to use for performing requests on an image of known
// orientation.
//
// imageData: Data containing the image to be used for performing the requests. Image
// content is immutable.
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// # Discussion
// 
// The intended use cases of this type of initializer include compressed
// images and network downloads, where a client may receive a JPEG from a
// website or the cloud.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:orientation:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:orientation:options:"), imageData, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to be used for performing requests on an image at the
// specified URL.
//
// imageURL: A URL pointing to the image to be used for performing the requests. The
// image must be in a format supported by [Image I/O]. Image content is
// immutable.
// //
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:options:)
func NewImageRequestHandlerWithURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), imageURL, options)
	return VNImageRequestHandlerFromID(rv)
}


// Creates a handler to be used for performing requests on an image with known
// orientation, at the specified URL.
//
// imageURL: A URL pointing to the image to be used for performing the requests. The
// image must be in a format supported by [Image I/O]. Image content is
// immutable.
// //
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:orientation:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewImageRequestHandlerWithURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	instance := getVNImageRequestHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:orientation:options:"), imageURL, orientation, options)
	return VNImageRequestHandlerFromID(rv)
}







// Creates a handler to be used for performing requests on Core Graphics
// images.
//
// image: A [CGImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:options:)
func (i VNImageRequestHandler) InitWithCGImageOptions(image coregraphics.CGImageRef, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCGImage:options:"), image, options)
	return rv
}

// Creates a handler to be used for performing requests on a Core Graphics
// image with known orientation.
//
// image: A [CGImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:orientation:options:)
func (i VNImageRequestHandler) InitWithCGImageOrientationOptions(image coregraphics.CGImageRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCGImage:orientation:options:"), image, orientation, options)
	return rv
}

// Creates a handler to use for performing requests on Core Image image data.
//
// image: A [CIImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// options: An optional dictionary containing [properties] keys to auxiliary image
// data.
// //
// [properties]: https://developer.apple.com/documentation/Vision/VNImageOption/properties
//
// image is a [coreimage.CIImage].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:options:)
func (i VNImageRequestHandler) InitWithCIImageOptions(image objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCIImage:options:"), image, options)
	return rv
}

// Creates a handler to be used for performing requests on Core Image image
// data of a known orientation.
//
// image: A [CIImage] containing the image to be used for performing the requests.
// Image content is immutable.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// image is a [coreimage.CIImage].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:orientation:options:)
func (i VNImageRequestHandler) InitWithCIImageOrientationOptions(image objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCIImage:orientation:options:"), image, orientation, options)
	return rv
}

// Creates a handler for performing requests on a Core Video pixel buffer.
//
// pixelBuffer: A pixel buffer that contains the image to use for performing the requests.
// The contents can’t change for the lifetime of the request handler.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:options:)
func (i VNImageRequestHandler) InitWithCVPixelBufferOptions(pixelBuffer corevideo.CVImageBufferRef, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a handler for performing requests on a Core Video pixel buffer of a
// known orientation.
//
// pixelBuffer: A pixel buffer that contains the image to use for performing the requests.
// The contents can’t change for the lifetime of the request handler.
//
// orientation: The orientation of the input image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:orientation:options:)
func (i VNImageRequestHandler) InitWithCVPixelBufferOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return rv
}

//
// depthData is a [avfoundation.AVDepthData].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:depthData:orientation:options:)
func (i VNImageRequestHandler) InitWithCVPixelBufferDepthDataOrientationOptions(pixelBuffer corevideo.CVImageBufferRef, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCVPixelBuffer:depthData:orientation:options:"), pixelBuffer, depthData, orientation, options)
	return rv
}

// Creates a request handler that performs requests on an image contained
// within a sample buffer.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:options:)
func (i VNImageRequestHandler) InitWithCMSampleBufferOptions(sampleBuffer objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a request handler that performs requests on an image of a specified
// orientation contained within a sample buffer.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// orientation: The EXIF orientation of the image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:orientation:options:)
func (i VNImageRequestHandler) InitWithCMSampleBufferOrientationOptions(sampleBuffer objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return rv
}

// Creates a request handler that performs requests on an image in a sample
// buffer that contains depth data.
//
// sampleBuffer: The sample buffer that contains the image to analyze. If the sample buffer
// doesn’t contain an image buffer with image data, the system raises an
// error.
//
// depthData: The depth data to associate with the image.
//
// orientation: The EXIF orientation of the image.
//
// options: A dictionary that specifies auxiliary information about the image.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// depthData is a [avfoundation.AVDepthData].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// # Discussion
// 
// Sample buffers may contain metadata, like the camera intrinsics. Vision
// algorithms that support this metadata use it in their analysis, unless
// overwritten by the options you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:depthData:orientation:options:)
func (i VNImageRequestHandler) InitWithCMSampleBufferDepthDataOrientationOptions(sampleBuffer objectivec.IObject, depthData objectivec.IObject, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithCMSampleBuffer:depthData:orientation:options:"), sampleBuffer, depthData, orientation, options)
	return rv
}

// Creates a handler to use for performing requests on an image in a data
// object.
//
// imageData: Data containing the image to be used for performing the requests. Image
// content is immutable.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// # Discussion
// 
// The intended use cases of this type of initializer include compressed
// images and network downloads, where a client may receive a JPEG from a
// website or the cloud.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:options:)
func (i VNImageRequestHandler) InitWithDataOptions(imageData foundation.INSData, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithData:options:"), imageData, options)
	return rv
}

// Creates a handler to use for performing requests on an image of known
// orientation.
//
// imageData: Data containing the image to be used for performing the requests. Image
// content is immutable.
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// # Discussion
// 
// The intended use cases of this type of initializer include compressed
// images and network downloads, where a client may receive a JPEG from a
// website or the cloud.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:orientation:options:)
func (i VNImageRequestHandler) InitWithDataOrientationOptions(imageData foundation.INSData, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithData:orientation:options:"), imageData, orientation, options)
	return rv
}

// Creates a handler to be used for performing requests on an image at the
// specified URL.
//
// imageURL: A URL pointing to the image to be used for performing the requests. The
// image must be in a format supported by [Image I/O]. Image content is
// immutable.
// //
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:options:)
func (i VNImageRequestHandler) InitWithURLOptions(imageURL foundation.INSURL, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithURL:options:"), imageURL, options)
	return rv
}

// Creates a handler to be used for performing requests on an image with known
// orientation, at the specified URL.
//
// imageURL: A URL pointing to the image to be used for performing the requests. The
// image must be in a format supported by [Image I/O]. Image content is
// immutable.
// //
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// orientation: The orientation of the input `image`.
//
// options: An optional dictionary containing [VNImageOption] keys to auxiliary image
// data.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:orientation:options:)
func (i VNImageRequestHandler) InitWithURLOrientationOptions(imageURL foundation.INSURL, orientation objectivec.IObject, options foundation.INSDictionary) VNImageRequestHandler {
	rv := objc.Send[VNImageRequestHandler](i.ID, objc.Sel("initWithURL:orientation:options:"), imageURL, orientation, options)
	return rv
}

// Schedules Vision requests to perform.
//
// requests: An array of Vision requests to perform.
//
// # Discussion
// 
// The function returns after all requests have either completed or failed.
// Check individual requests and errors for their respective successes and
// failures.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/perform(_:)
func (i VNImageRequestHandler) PerformRequestsError(requests []VNRequest) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("performRequests:error:"), objectivec.IObjectSliceToNSArray(requests), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:error: returned NO with nil NSError")
	}
	return rv, nil

}
































