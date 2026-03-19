// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectTextRectanglesRequest] class.
var (
	_VNDetectTextRectanglesRequestClass     VNDetectTextRectanglesRequestClass
	_VNDetectTextRectanglesRequestClassOnce sync.Once
)

func getVNDetectTextRectanglesRequestClass() VNDetectTextRectanglesRequestClass {
	_VNDetectTextRectanglesRequestClassOnce.Do(func() {
		_VNDetectTextRectanglesRequestClass = VNDetectTextRectanglesRequestClass{class: objc.GetClass("VNDetectTextRectanglesRequest")}
	})
	return _VNDetectTextRectanglesRequestClass
}

// GetVNDetectTextRectanglesRequestClass returns the class object for VNDetectTextRectanglesRequest.
func GetVNDetectTextRectanglesRequestClass() VNDetectTextRectanglesRequestClass {
	return getVNDetectTextRectanglesRequestClass()
}

type VNDetectTextRectanglesRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectTextRectanglesRequestClass) Alloc() VNDetectTextRectanglesRequest {
	rv := objc.Send[VNDetectTextRectanglesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that finds regions of visible text in an image.
//
// # Overview
// 
// This request returns detected text characters as rectangular bounding boxes
// with origin and size.
//
// # Configuring a Request
//
//   - [VNDetectTextRectanglesRequest.ReportCharacterBoxes]: A Boolean value that indicates whether the request detects character bounding boxes.
//   - [VNDetectTextRectanglesRequest.SetReportCharacterBoxes]
//
// # Identifying Request Revisions
//
//   - [VNDetectTextRectanglesRequest.VNDetectTextRectanglesRequestRevision1]: A constant for specifying revision 1 of the text rectangles detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest
type VNDetectTextRectanglesRequest struct {
	VNImageBasedRequest
}

// VNDetectTextRectanglesRequestFromID constructs a [VNDetectTextRectanglesRequest] from an objc.ID.
//
// An image-analysis request that finds regions of visible text in an image.
func VNDetectTextRectanglesRequestFromID(id objc.ID) VNDetectTextRectanglesRequest {
	return VNDetectTextRectanglesRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectTextRectanglesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectTextRectanglesRequest] class.
//
// # Configuring a Request
//
//   - [IVNDetectTextRectanglesRequest.ReportCharacterBoxes]: A Boolean value that indicates whether the request detects character bounding boxes.
//   - [IVNDetectTextRectanglesRequest.SetReportCharacterBoxes]
//
// # Identifying Request Revisions
//
//   - [IVNDetectTextRectanglesRequest.VNDetectTextRectanglesRequestRevision1]: A constant for specifying revision 1 of the text rectangles detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest
type IVNDetectTextRectanglesRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring a Request

	// A Boolean value that indicates whether the request detects character bounding boxes.
	ReportCharacterBoxes() bool
	SetReportCharacterBoxes(value bool)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the text rectangles detection request.
	VNDetectTextRectanglesRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectTextRectanglesRequest) Init() VNDetectTextRectanglesRequest {
	rv := objc.Send[VNDetectTextRectanglesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectTextRectanglesRequest) Autorelease() VNDetectTextRectanglesRequest {
	rv := objc.Send[VNDetectTextRectanglesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectTextRectanglesRequest creates a new VNDetectTextRectanglesRequest instance.
func NewVNDetectTextRectanglesRequest() VNDetectTextRectanglesRequest {
	class := getVNDetectTextRectanglesRequestClass()
	rv := objc.Send[VNDetectTextRectanglesRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectTextRectanglesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectTextRectanglesRequest {
	instance := getVNDetectTextRectanglesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectTextRectanglesRequestFromID(rv)
}

// A Boolean value that indicates whether the request detects character
// bounding boxes.
//
// # Discussion
// 
// Set the value to [true] to have the detector return character bounding
// boxes as an array of [VNRectangleObservation] objects.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest/reportCharacterBoxes
func (d VNDetectTextRectanglesRequest) ReportCharacterBoxes() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("reportCharacterBoxes"))
	return rv
}
func (d VNDetectTextRectanglesRequest) SetReportCharacterBoxes(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setReportCharacterBoxes:"), value)
}
// A constant for specifying revision 1 of the text rectangles detection
// request.
//
// See: https://developer.apple.com/documentation/vision/vndetecttextrectanglesrequestrevision1
func (d VNDetectTextRectanglesRequest) VNDetectTextRectanglesRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectTextRectanglesRequestRevision1"))
	return rv
}

