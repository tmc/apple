// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectDocumentSegmentationRequest] class.
var (
	_VNDetectDocumentSegmentationRequestClass     VNDetectDocumentSegmentationRequestClass
	_VNDetectDocumentSegmentationRequestClassOnce sync.Once
)

func getVNDetectDocumentSegmentationRequestClass() VNDetectDocumentSegmentationRequestClass {
	_VNDetectDocumentSegmentationRequestClassOnce.Do(func() {
		_VNDetectDocumentSegmentationRequestClass = VNDetectDocumentSegmentationRequestClass{class: objc.GetClass("VNDetectDocumentSegmentationRequest")}
	})
	return _VNDetectDocumentSegmentationRequestClass
}

// GetVNDetectDocumentSegmentationRequestClass returns the class object for VNDetectDocumentSegmentationRequest.
func GetVNDetectDocumentSegmentationRequestClass() VNDetectDocumentSegmentationRequestClass {
	return getVNDetectDocumentSegmentationRequestClass()
}

type VNDetectDocumentSegmentationRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectDocumentSegmentationRequestClass) Alloc() VNDetectDocumentSegmentationRequest {
	rv := objc.Send[VNDetectDocumentSegmentationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that detects rectangular regions that contain text in the input
// image.
//
// # Overview
// 
// Perform this request to detect a document in an image. The result that the
// request generates contains the four corner points of a document’s
// quadrilateral and saliency mask.
//
// # Identifying Request Revisions
//
//   - [VNDetectDocumentSegmentationRequest.VNDetectDocumentSegmentationRequestRevision1]: A constant for specifying revision 1 of the document segmentation request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectDocumentSegmentationRequest
type VNDetectDocumentSegmentationRequest struct {
	VNImageBasedRequest
}

// VNDetectDocumentSegmentationRequestFromID constructs a [VNDetectDocumentSegmentationRequest] from an objc.ID.
//
// An object that detects rectangular regions that contain text in the input
// image.
func VNDetectDocumentSegmentationRequestFromID(id objc.ID) VNDetectDocumentSegmentationRequest {
	return VNDetectDocumentSegmentationRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectDocumentSegmentationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectDocumentSegmentationRequest] class.
//
// # Identifying Request Revisions
//
//   - [IVNDetectDocumentSegmentationRequest.VNDetectDocumentSegmentationRequestRevision1]: A constant for specifying revision 1 of the document segmentation request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectDocumentSegmentationRequest
type IVNDetectDocumentSegmentationRequest interface {
	IVNImageBasedRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the document segmentation request.
	VNDetectDocumentSegmentationRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectDocumentSegmentationRequest) Init() VNDetectDocumentSegmentationRequest {
	rv := objc.Send[VNDetectDocumentSegmentationRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectDocumentSegmentationRequest) Autorelease() VNDetectDocumentSegmentationRequest {
	rv := objc.Send[VNDetectDocumentSegmentationRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectDocumentSegmentationRequest creates a new VNDetectDocumentSegmentationRequest instance.
func NewVNDetectDocumentSegmentationRequest() VNDetectDocumentSegmentationRequest {
	class := getVNDetectDocumentSegmentationRequestClass()
	rv := objc.Send[VNDetectDocumentSegmentationRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectDocumentSegmentationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectDocumentSegmentationRequest {
	instance := getVNDetectDocumentSegmentationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectDocumentSegmentationRequestFromID(rv)
}

// A constant for specifying revision 1 of the document segmentation request.
//
// See: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequestrevision1
func (d VNDetectDocumentSegmentationRequest) VNDetectDocumentSegmentationRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectDocumentSegmentationRequestRevision1"))
	return rv
}

