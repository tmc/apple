// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNDetectFaceCaptureQualityRequest] class.
var (
	_VNDetectFaceCaptureQualityRequestClass     VNDetectFaceCaptureQualityRequestClass
	_VNDetectFaceCaptureQualityRequestClassOnce sync.Once
)

func getVNDetectFaceCaptureQualityRequestClass() VNDetectFaceCaptureQualityRequestClass {
	_VNDetectFaceCaptureQualityRequestClassOnce.Do(func() {
		_VNDetectFaceCaptureQualityRequestClass = VNDetectFaceCaptureQualityRequestClass{class: objc.GetClass("VNDetectFaceCaptureQualityRequest")}
	})
	return _VNDetectFaceCaptureQualityRequestClass
}

// GetVNDetectFaceCaptureQualityRequestClass returns the class object for VNDetectFaceCaptureQualityRequest.
func GetVNDetectFaceCaptureQualityRequestClass() VNDetectFaceCaptureQualityRequestClass {
	return getVNDetectFaceCaptureQualityRequestClass()
}

type VNDetectFaceCaptureQualityRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectFaceCaptureQualityRequestClass) Alloc() VNDetectFaceCaptureQualityRequest {
	rv := objc.Send[VNDetectFaceCaptureQualityRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that produces a floating-point number that represents the capture
// quality of a face in a photo.
//
// # Overview
// 
// This request produces or updates a [VNFaceObservation] object’s property
// [faceCaptureQuality] with a floating-point value. The value ranges from `0`
// to `1`. Faces with quality closer to `1` are better lit, sharper, and more
// centrally positioned than faces with quality closer to `0`.
// 
// If you don’t execute the request, or the request fails, the property
// [faceCaptureQuality] is [nil].
//
// [faceCaptureQuality]: https://developer.apple.com/documentation/Vision/VNFaceObservation/faceCaptureQuality-bjg5
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Identifying Request Revisions
//
//   - [VNDetectFaceCaptureQualityRequest.VNDetectFaceCaptureQualityRequestRevision2]: Revision 2 of the request algorithm.
//   - [VNDetectFaceCaptureQualityRequest.VNDetectFaceCaptureQualityRequestRevision1]: A constant for specifying revision 1 of the face capture detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceCaptureQualityRequest
type VNDetectFaceCaptureQualityRequest struct {
	VNImageBasedRequest
}

// VNDetectFaceCaptureQualityRequestFromID constructs a [VNDetectFaceCaptureQualityRequest] from an objc.ID.
//
// A request that produces a floating-point number that represents the capture
// quality of a face in a photo.
func VNDetectFaceCaptureQualityRequestFromID(id objc.ID) VNDetectFaceCaptureQualityRequest {
	return VNDetectFaceCaptureQualityRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectFaceCaptureQualityRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectFaceCaptureQualityRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNDetectFaceCaptureQualityRequest.VNDetectFaceCaptureQualityRequestRevision2]: Revision 2 of the request algorithm.
//   - [IVNDetectFaceCaptureQualityRequest.VNDetectFaceCaptureQualityRequestRevision1]: A constant for specifying revision 1 of the face capture detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectFaceCaptureQualityRequest
type IVNDetectFaceCaptureQualityRequest interface {
	IVNImageBasedRequest
	VNFaceObservationAccepting

	// Topic: Identifying Request Revisions

	// Revision 2 of the request algorithm.
	VNDetectFaceCaptureQualityRequestRevision2() int
	// A constant for specifying revision 1 of the face capture detection request.
	VNDetectFaceCaptureQualityRequestRevision1() int

	// A value that indicates the quality of the face capture.
	FaceCaptureQuality() float32
	SetFaceCaptureQuality(value float32)
}

// Init initializes the instance.
func (d VNDetectFaceCaptureQualityRequest) Init() VNDetectFaceCaptureQualityRequest {
	rv := objc.Send[VNDetectFaceCaptureQualityRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectFaceCaptureQualityRequest) Autorelease() VNDetectFaceCaptureQualityRequest {
	rv := objc.Send[VNDetectFaceCaptureQualityRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectFaceCaptureQualityRequest creates a new VNDetectFaceCaptureQualityRequest instance.
func NewVNDetectFaceCaptureQualityRequest() VNDetectFaceCaptureQualityRequest {
	class := getVNDetectFaceCaptureQualityRequestClass()
	rv := objc.Send[VNDetectFaceCaptureQualityRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectFaceCaptureQualityRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectFaceCaptureQualityRequest {
	instance := getVNDetectFaceCaptureQualityRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectFaceCaptureQualityRequestFromID(rv)
}

// Revision 2 of the request algorithm.
//
// See: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequestrevision2
func (d VNDetectFaceCaptureQualityRequest) VNDetectFaceCaptureQualityRequestRevision2() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectFaceCaptureQualityRequestRevision2"))
	return rv
}
// A constant for specifying revision 1 of the face capture detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequestrevision1
func (d VNDetectFaceCaptureQualityRequest) VNDetectFaceCaptureQualityRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectFaceCaptureQualityRequestRevision1"))
	return rv
}
// A value that indicates the quality of the face capture.
//
// See: https://developer.apple.com/documentation/vision/vnfaceobservation/facecapturequality-bjg5
func (d VNDetectFaceCaptureQualityRequest) FaceCaptureQuality() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("faceCaptureQuality"))
	return rv
}
func (d VNDetectFaceCaptureQualityRequest) SetFaceCaptureQuality(value float32) {
	objc.Send[struct{}](d.ID, objc.Sel("setFaceCaptureQuality:"), value)
}
// An array of [VNFaceObservation] objects to process as part of the request.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceObservationAccepting/inputFaceObservations
func (d VNDetectFaceCaptureQualityRequest) InputFaceObservations() []VNFaceObservation {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("inputFaceObservations"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNFaceObservation {
		return VNFaceObservationFromID(id)
	})
}
func (d VNDetectFaceCaptureQualityRequest) SetInputFaceObservations(value []VNFaceObservation) {
	objc.Send[struct{}](d.ID, objc.Sel("setInputFaceObservations:"), objectivec.IObjectSliceToNSArray(value))
}

			// Protocol methods for VNFaceObservationAccepting
			

