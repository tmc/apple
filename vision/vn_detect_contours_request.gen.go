// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNDetectContoursRequest] class.
var (
	_VNDetectContoursRequestClass     VNDetectContoursRequestClass
	_VNDetectContoursRequestClassOnce sync.Once
)

func getVNDetectContoursRequestClass() VNDetectContoursRequestClass {
	_VNDetectContoursRequestClassOnce.Do(func() {
		_VNDetectContoursRequestClass = VNDetectContoursRequestClass{class: objc.GetClass("VNDetectContoursRequest")}
	})
	return _VNDetectContoursRequestClass
}

// GetVNDetectContoursRequestClass returns the class object for VNDetectContoursRequest.
func GetVNDetectContoursRequestClass() VNDetectContoursRequestClass {
	return getVNDetectContoursRequestClass()
}

type VNDetectContoursRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectContoursRequestClass) Alloc() VNDetectContoursRequest {
	rv := objc.Send[VNDetectContoursRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A request that detects the contours of the edges of an image.
//
// # Configuring the Request
//
//   - [VNDetectContoursRequest.ContrastAdjustment]: The amount by which to adjust the image contrast.
//   - [VNDetectContoursRequest.SetContrastAdjustment]
//   - [VNDetectContoursRequest.ContrastPivot]: The pixel value to use as a pivot for the contrast.
//   - [VNDetectContoursRequest.SetContrastPivot]
//   - [VNDetectContoursRequest.DetectsDarkOnLight]: A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
//   - [VNDetectContoursRequest.SetDetectsDarkOnLight]
//   - [VNDetectContoursRequest.MaximumImageDimension]: The maximum image dimension to use for contour detection.
//   - [VNDetectContoursRequest.SetMaximumImageDimension]
//
// # Identifying Request Revisions
//
//   - [VNDetectContoursRequest.VNDetectContourRequestRevision1]: A constant for specifying revision 1 of the contours detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest
type VNDetectContoursRequest struct {
	VNImageBasedRequest
}

// VNDetectContoursRequestFromID constructs a [VNDetectContoursRequest] from an objc.ID.
//
// A request that detects the contours of the edges of an image.
func VNDetectContoursRequestFromID(id objc.ID) VNDetectContoursRequest {
	return VNDetectContoursRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectContoursRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNDetectContoursRequest] class.
//
// # Configuring the Request
//
//   - [IVNDetectContoursRequest.ContrastAdjustment]: The amount by which to adjust the image contrast.
//   - [IVNDetectContoursRequest.SetContrastAdjustment]
//   - [IVNDetectContoursRequest.ContrastPivot]: The pixel value to use as a pivot for the contrast.
//   - [IVNDetectContoursRequest.SetContrastPivot]
//   - [IVNDetectContoursRequest.DetectsDarkOnLight]: A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
//   - [IVNDetectContoursRequest.SetDetectsDarkOnLight]
//   - [IVNDetectContoursRequest.MaximumImageDimension]: The maximum image dimension to use for contour detection.
//   - [IVNDetectContoursRequest.SetMaximumImageDimension]
//
// # Identifying Request Revisions
//
//   - [IVNDetectContoursRequest.VNDetectContourRequestRevision1]: A constant for specifying revision 1 of the contours detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest
type IVNDetectContoursRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring the Request

	// The amount by which to adjust the image contrast.
	ContrastAdjustment() float32
	SetContrastAdjustment(value float32)
	// The pixel value to use as a pivot for the contrast.
	ContrastPivot() foundation.NSNumber
	SetContrastPivot(value foundation.NSNumber)
	// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
	DetectsDarkOnLight() bool
	SetDetectsDarkOnLight(value bool)
	// The maximum image dimension to use for contour detection.
	MaximumImageDimension() uint
	SetMaximumImageDimension(value uint)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the contours detection request.
	VNDetectContourRequestRevision1() int
}





// Init initializes the instance.
func (d VNDetectContoursRequest) Init() VNDetectContoursRequest {
	rv := objc.Send[VNDetectContoursRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectContoursRequest) Autorelease() VNDetectContoursRequest {
	rv := objc.Send[VNDetectContoursRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectContoursRequest creates a new VNDetectContoursRequest instance.
func NewVNDetectContoursRequest() VNDetectContoursRequest {
	class := getVNDetectContoursRequestClass()
	rv := objc.Send[VNDetectContoursRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectContoursRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectContoursRequest {
	instance := getVNDetectContoursRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectContoursRequestFromID(rv)
}


















// The amount by which to adjust the image contrast.
//
// # Discussion
// 
// Contour detection works best with high-contrast images. The default value
// of this property is `2.0`, which doubles the image contrast to achieve the
// most accurate results.
// 
// This property supports a value range from `0.0` to `3.0`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastAdjustment
func (d VNDetectContoursRequest) ContrastAdjustment() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("contrastAdjustment"))
	return rv
}
func (d VNDetectContoursRequest) SetContrastAdjustment(value float32) {
	objc.Send[struct{}](d.ID, objc.Sel("setContrastAdjustment:"), value)
}



// The pixel value to use as a pivot for the contrast.
//
// # Discussion
// 
// Numeric values range from `0.0` to `1.0`. You can also specify `nil` to
// have the framework automatically detect the value according to image
// intensity.
// 
// The default value is `0.5`, which indicates the pixel center.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastPivot
func (d VNDetectContoursRequest) ContrastPivot() foundation.NSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("contrastPivot"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (d VNDetectContoursRequest) SetContrastPivot(value foundation.NSNumber) {
	objc.Send[struct{}](d.ID, objc.Sel("setContrastPivot:"), value)
}



// A Boolean value that indicates whether the request detects a dark object on
// a light background to aid in detection.
//
// # Discussion
// 
// The default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/detectsDarkOnLight
func (d VNDetectContoursRequest) DetectsDarkOnLight() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("detectsDarkOnLight"))
	return rv
}
func (d VNDetectContoursRequest) SetDetectsDarkOnLight(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDetectsDarkOnLight:"), value)
}



// The maximum image dimension to use for contour detection.
//
// # Discussion
// 
// Contour detection is computationally intensive. To improve performance,
// Vision scales the input image down, while maintaining its aspect ratio,
// such that its maximum dimension is the value of this property. Vision never
// scales the image up, so specifying the maximum value ensures that the image
// processes in its original size and not as a downscaled version.
// 
// This property supports values from 64 to [NSUIntegerMax]. The default value
// is 512.
//
// [NSUIntegerMax]: https://developer.apple.com/documentation/ObjectiveC/NSUIntegerMax
//
// See: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/maximumImageDimension
func (d VNDetectContoursRequest) MaximumImageDimension() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("maximumImageDimension"))
	return rv
}
func (d VNDetectContoursRequest) SetMaximumImageDimension(value uint) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumImageDimension:"), value)
}



// A constant for specifying revision 1 of the contours detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectcontourrequestrevision1
func (d VNDetectContoursRequest) VNDetectContourRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectContourRequestRevision1"))
	return rv
}
























