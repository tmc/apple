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

// The class instance for the [VNSequenceRequestHandler] class.
var (
	_VNSequenceRequestHandlerClass     VNSequenceRequestHandlerClass
	_VNSequenceRequestHandlerClassOnce sync.Once
)

func getVNSequenceRequestHandlerClass() VNSequenceRequestHandlerClass {
	_VNSequenceRequestHandlerClassOnce.Do(func() {
		_VNSequenceRequestHandlerClass = VNSequenceRequestHandlerClass{class: objc.GetClass("VNSequenceRequestHandler")}
	})
	return _VNSequenceRequestHandlerClass
}

// GetVNSequenceRequestHandlerClass returns the class object for VNSequenceRequestHandler.
func GetVNSequenceRequestHandlerClass() VNSequenceRequestHandlerClass {
	return getVNSequenceRequestHandlerClass()
}

type VNSequenceRequestHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNSequenceRequestHandlerClass) Alloc() VNSequenceRequestHandler {
	rv := objc.Send[VNSequenceRequestHandler](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that processes image-analysis requests for each frame in a
// sequence.
//
// # Overview
// 
// Instantiate this handler to perform Vision requests on a series of images.
// Unlike the [VNImageRequestHandler], you don’t specify the image on
// creation. Instead, you supply each image frame one by one as you continue
// to call one of the `perform` methods.
//
// # Performing a Sequence Request
//
//   - [VNSequenceRequestHandler.PerformRequestsOnCGImageError]: Schedules Vision requests to be performed on a Core Graphics image.
//   - [VNSequenceRequestHandler.PerformRequestsOnCGImageOrientationError]: Schedules one or more Vision requests to be performed on a Core Graphics image with known orientation.
//   - [VNSequenceRequestHandler.PerformRequestsOnCIImageError]: Schedules one or more Vision requests to be performed on Core Image image data.
//   - [VNSequenceRequestHandler.PerformRequestsOnCIImageOrientationError]: Schedules one or more Vision requests to be performed on Core Image image data with known orientation.
//   - [VNSequenceRequestHandler.PerformRequestsOnCVPixelBufferError]: Schedules one or more Vision requests to be performed on a Core Video pixel buffer.
//   - [VNSequenceRequestHandler.PerformRequestsOnCVPixelBufferOrientationError]: Schedules one or more Vision requests to be performed on a Core Video pixel buffer with known orientation.
//   - [VNSequenceRequestHandler.PerformRequestsOnCMSampleBufferError]: Performs one or more requests on an image contained within a sample buffer.
//   - [VNSequenceRequestHandler.PerformRequestsOnCMSampleBufferOrientationError]: Performs one or more requests on an image of a specified orientation contained within a sample buffer.
//   - [VNSequenceRequestHandler.PerformRequestsOnImageDataError]: Schedules one or more Vision requests to be performed on raw image data.
//   - [VNSequenceRequestHandler.PerformRequestsOnImageDataOrientationError]: Schedules one or more Vision requests to be performed on raw data containing an image with known orientation.
//   - [VNSequenceRequestHandler.PerformRequestsOnImageURLError]: Schedules one or more Vision requests to be performed on an image.
//   - [VNSequenceRequestHandler.PerformRequestsOnImageURLOrientationError]: Schedules one or more Vision requests to be performed on an image with known orientation, at a specific URL.
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler
type VNSequenceRequestHandler struct {
	objectivec.Object
}

// VNSequenceRequestHandlerFromID constructs a [VNSequenceRequestHandler] from an objc.ID.
//
// An object that processes image-analysis requests for each frame in a
// sequence.
func VNSequenceRequestHandlerFromID(id objc.ID) VNSequenceRequestHandler {
	return VNSequenceRequestHandler{objectivec.Object{ID: id}}
}
// NOTE: VNSequenceRequestHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNSequenceRequestHandler] class.
//
// # Performing a Sequence Request
//
//   - [IVNSequenceRequestHandler.PerformRequestsOnCGImageError]: Schedules Vision requests to be performed on a Core Graphics image.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCGImageOrientationError]: Schedules one or more Vision requests to be performed on a Core Graphics image with known orientation.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCIImageError]: Schedules one or more Vision requests to be performed on Core Image image data.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCIImageOrientationError]: Schedules one or more Vision requests to be performed on Core Image image data with known orientation.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCVPixelBufferError]: Schedules one or more Vision requests to be performed on a Core Video pixel buffer.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCVPixelBufferOrientationError]: Schedules one or more Vision requests to be performed on a Core Video pixel buffer with known orientation.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCMSampleBufferError]: Performs one or more requests on an image contained within a sample buffer.
//   - [IVNSequenceRequestHandler.PerformRequestsOnCMSampleBufferOrientationError]: Performs one or more requests on an image of a specified orientation contained within a sample buffer.
//   - [IVNSequenceRequestHandler.PerformRequestsOnImageDataError]: Schedules one or more Vision requests to be performed on raw image data.
//   - [IVNSequenceRequestHandler.PerformRequestsOnImageDataOrientationError]: Schedules one or more Vision requests to be performed on raw data containing an image with known orientation.
//   - [IVNSequenceRequestHandler.PerformRequestsOnImageURLError]: Schedules one or more Vision requests to be performed on an image.
//   - [IVNSequenceRequestHandler.PerformRequestsOnImageURLOrientationError]: Schedules one or more Vision requests to be performed on an image with known orientation, at a specific URL.
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler
type IVNSequenceRequestHandler interface {
	objectivec.IObject

	// Topic: Performing a Sequence Request

	// Schedules Vision requests to be performed on a Core Graphics image.
	PerformRequestsOnCGImageError(requests []VNRequest, image coregraphics.CGImageRef) (bool, error)
	// Schedules one or more Vision requests to be performed on a Core Graphics image with known orientation.
	PerformRequestsOnCGImageOrientationError(requests []VNRequest, image coregraphics.CGImageRef, orientation objectivec.IObject) (bool, error)
	// Schedules one or more Vision requests to be performed on Core Image image data.
	PerformRequestsOnCIImageError(requests []VNRequest, image objectivec.IObject) (bool, error)
	// Schedules one or more Vision requests to be performed on Core Image image data with known orientation.
	PerformRequestsOnCIImageOrientationError(requests []VNRequest, image objectivec.IObject, orientation objectivec.IObject) (bool, error)
	// Schedules one or more Vision requests to be performed on a Core Video pixel buffer.
	PerformRequestsOnCVPixelBufferError(requests []VNRequest, pixelBuffer corevideo.CVImageBufferRef) (bool, error)
	// Schedules one or more Vision requests to be performed on a Core Video pixel buffer with known orientation.
	PerformRequestsOnCVPixelBufferOrientationError(requests []VNRequest, pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject) (bool, error)
	// Performs one or more requests on an image contained within a sample buffer.
	PerformRequestsOnCMSampleBufferError(requests []VNRequest, sampleBuffer objectivec.IObject) (bool, error)
	// Performs one or more requests on an image of a specified orientation contained within a sample buffer.
	PerformRequestsOnCMSampleBufferOrientationError(requests []VNRequest, sampleBuffer objectivec.IObject, orientation objectivec.IObject) (bool, error)
	// Schedules one or more Vision requests to be performed on raw image data.
	PerformRequestsOnImageDataError(requests []VNRequest, imageData foundation.INSData) (bool, error)
	// Schedules one or more Vision requests to be performed on raw data containing an image with known orientation.
	PerformRequestsOnImageDataOrientationError(requests []VNRequest, imageData foundation.INSData, orientation objectivec.IObject) (bool, error)
	// Schedules one or more Vision requests to be performed on an image.
	PerformRequestsOnImageURLError(requests []VNRequest, imageURL foundation.INSURL) (bool, error)
	// Schedules one or more Vision requests to be performed on an image with known orientation, at a specific URL.
	PerformRequestsOnImageURLOrientationError(requests []VNRequest, imageURL foundation.INSURL, orientation objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (s VNSequenceRequestHandler) Init() VNSequenceRequestHandler {
	rv := objc.Send[VNSequenceRequestHandler](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VNSequenceRequestHandler) Autorelease() VNSequenceRequestHandler {
	rv := objc.Send[VNSequenceRequestHandler](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNSequenceRequestHandler creates a new VNSequenceRequestHandler instance.
func NewVNSequenceRequestHandler() VNSequenceRequestHandler {
	class := getVNSequenceRequestHandlerClass()
	rv := objc.Send[VNSequenceRequestHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Schedules Vision requests to be performed on a Core Graphics image.
//
// requests: An array of [VNRequest] requests to perform.
//
// image: The input [CGImage] on which to perform the request.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-3zt7l
func (s VNSequenceRequestHandler) PerformRequestsOnCGImageError(requests []VNRequest, image coregraphics.CGImageRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCGImage:error:"), objectivec.IObjectSliceToNSArray(requests), image, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCGImage:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on a Core Graphics
// image with known orientation.
//
// requests: An array of [VNRequest] requests to perform.
//
// image: The input [CGImage] on which to perform the request.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// orientation: The orientation of the input `image`.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-3gcmv
func (s VNSequenceRequestHandler) PerformRequestsOnCGImageOrientationError(requests []VNRequest, image coregraphics.CGImageRef, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCGImage:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), image, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCGImage:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on Core Image image
// data.
//
// requests: An array of [VNRequest] requests to perform.
//
// image: The input [CIImage] on which to perform the request.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// image is a [coreimage.CIImage].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-9jtgj
func (s VNSequenceRequestHandler) PerformRequestsOnCIImageError(requests []VNRequest, image objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCIImage:error:"), objectivec.IObjectSliceToNSArray(requests), image, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCIImage:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on Core Image image
// data with known orientation.
//
// requests: An array of [VNRequest] requests to perform.
//
// image: The input [CIImage] on which to perform the request.
// //
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
//
// orientation: The orientation of the input `image`.
//
// image is a [coreimage.CIImage].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-1bkm1
func (s VNSequenceRequestHandler) PerformRequestsOnCIImageOrientationError(requests []VNRequest, image objectivec.IObject, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCIImage:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), image, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCIImage:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on a Core Video pixel
// buffer.
//
// requests: An array of [VNRequest] requests to perform.
//
// pixelBuffer: The input [CVPixelBuffer] on which to perform the request.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-3d7nt
func (s VNSequenceRequestHandler) PerformRequestsOnCVPixelBufferError(requests []VNRequest, pixelBuffer corevideo.CVImageBufferRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCVPixelBuffer:error:"), objectivec.IObjectSliceToNSArray(requests), pixelBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCVPixelBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on a Core Video pixel
// buffer with known orientation.
//
// requests: An array of [VNRequest] requests to perform.
//
// pixelBuffer: The input [CVPixelBuffer] on which to perform the request.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// orientation: The orientation of the input `image`.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-2wvt8
func (s VNSequenceRequestHandler) PerformRequestsOnCVPixelBufferOrientationError(requests []VNRequest, pixelBuffer corevideo.CVImageBufferRef, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCVPixelBuffer:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), pixelBuffer, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCVPixelBuffer:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Performs one or more requests on an image contained within a sample buffer.
//
// requests: The array of requests to perform.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-45e73
func (s VNSequenceRequestHandler) PerformRequestsOnCMSampleBufferError(requests []VNRequest, sampleBuffer objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCMSampleBuffer:error:"), objectivec.IObjectSliceToNSArray(requests), sampleBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCMSampleBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Performs one or more requests on an image of a specified orientation
// contained within a sample buffer.
//
// requests: The array of requests to perform.
//
// sampleBuffer: A sample buffer containing a valid [imageBuffer].
// //
// [imageBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/imageBuffer
//
// orientation: The orientation of the image contained within the sample buffer.
//
// sampleBuffer is a [coremedia.CMSampleBufferRef].
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-6b7rk
func (s VNSequenceRequestHandler) PerformRequestsOnCMSampleBufferOrientationError(requests []VNRequest, sampleBuffer objectivec.IObject, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onCMSampleBuffer:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), sampleBuffer, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onCMSampleBuffer:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on raw image data.
//
// requests: An array of [VNRequest] requests to perform.
//
// imageData: The input image data on which to perform the request.
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageData:)
func (s VNSequenceRequestHandler) PerformRequestsOnImageDataError(requests []VNRequest, imageData foundation.INSData) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onImageData:error:"), objectivec.IObjectSliceToNSArray(requests), imageData, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onImageData:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on raw data
// containing an image with known orientation.
//
// requests: An array of [VNRequest] requests to perform.
//
// imageData: The input image data on which to perform the request.
//
// orientation: The orientation of the input image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageData:orientation:)
func (s VNSequenceRequestHandler) PerformRequestsOnImageDataOrientationError(requests []VNRequest, imageData foundation.INSData, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onImageData:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), imageData, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onImageData:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on an image.
//
// requests: An array of [VNRequest] requests to perform.
//
// imageURL: A URL pointing to the image on which to perform the request.
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageURL:)
func (s VNSequenceRequestHandler) PerformRequestsOnImageURLError(requests []VNRequest, imageURL foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onImageURL:error:"), objectivec.IObjectSliceToNSArray(requests), imageURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onImageURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules one or more Vision requests to be performed on an image with
// known orientation, at a specific URL.
//
// requests: An array of [VNRequest] requests to perform.
//
// imageURL: A URL pointing to the image on which to perform the request.
//
// orientation: The orientation of the input image.
//
// orientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageURL:orientation:)
func (s VNSequenceRequestHandler) PerformRequestsOnImageURLOrientationError(requests []VNRequest, imageURL foundation.INSURL, orientation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("performRequests:onImageURL:orientation:error:"), objectivec.IObjectSliceToNSArray(requests), imageURL, orientation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performRequests:onImageURL:orientation:error: returned NO with nil NSError")
	}
	return rv, nil

}

