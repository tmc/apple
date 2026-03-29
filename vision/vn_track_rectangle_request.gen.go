// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNTrackRectangleRequest] class.
var (
	_VNTrackRectangleRequestClass     VNTrackRectangleRequestClass
	_VNTrackRectangleRequestClassOnce sync.Once
)

func getVNTrackRectangleRequestClass() VNTrackRectangleRequestClass {
	_VNTrackRectangleRequestClassOnce.Do(func() {
		_VNTrackRectangleRequestClass = VNTrackRectangleRequestClass{class: objc.GetClass("VNTrackRectangleRequest")}
	})
	return _VNTrackRectangleRequestClass
}

// GetVNTrackRectangleRequestClass returns the class object for VNTrackRectangleRequest.
func GetVNTrackRectangleRequestClass() VNTrackRectangleRequestClass {
	return getVNTrackRectangleRequestClass()
}

type VNTrackRectangleRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNTrackRectangleRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackRectangleRequestClass) Alloc() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that tracks movement of a previously identified
// rectangular object across multiple images or video frames.
//
// # Overview
// 
// Use this type of request to track the bounding boxes of rectangles
// throughout a sequence of images. Vision returns locations for rectangles
// found in all orientations and sizes.
//
// # Initializing a Rectangle Tracking Request
//
//   - [VNTrackRectangleRequest.InitWithRectangleObservation]: Creates a new rectangle tracking request with a rectangle observation.
//   - [VNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]: Creates a new rectangle tracking request with a rectangle observation.
//
// # Identifying Request Revisions
//
//   - [VNTrackRectangleRequest.VNTrackRectangleRequestRevision1]: A constant for specifying revision 1 of the rectangling tracking request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest
type VNTrackRectangleRequest struct {
	VNTrackingRequest
}

// VNTrackRectangleRequestFromID constructs a [VNTrackRectangleRequest] from an objc.ID.
//
// An image-analysis request that tracks movement of a previously identified
// rectangular object across multiple images or video frames.
func VNTrackRectangleRequestFromID(id objc.ID) VNTrackRectangleRequest {
	return VNTrackRectangleRequest{VNTrackingRequest: VNTrackingRequestFromID(id)}
}
// NOTE: VNTrackRectangleRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackRectangleRequest] class.
//
// # Initializing a Rectangle Tracking Request
//
//   - [IVNTrackRectangleRequest.InitWithRectangleObservation]: Creates a new rectangle tracking request with a rectangle observation.
//   - [IVNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]: Creates a new rectangle tracking request with a rectangle observation.
//
// # Identifying Request Revisions
//
//   - [IVNTrackRectangleRequest.VNTrackRectangleRequestRevision1]: A constant for specifying revision 1 of the rectangling tracking request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest
type IVNTrackRectangleRequest interface {
	IVNTrackingRequest

	// Topic: Initializing a Rectangle Tracking Request

	// Creates a new rectangle tracking request with a rectangle observation.
	InitWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest
	// Creates a new rectangle tracking request with a rectangle observation.
	InitWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler ErrorHandler) VNTrackRectangleRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the rectangling tracking request.
	VNTrackRectangleRequestRevision1() int
}

// Init initializes the instance.
func (t VNTrackRectangleRequest) Init() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackRectangleRequest) Autorelease() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackRectangleRequest creates a new VNTrackRectangleRequest instance.
func NewVNTrackRectangleRequest() VNTrackRectangleRequest {
	class := getVNTrackRectangleRequestClass()
	rv := objc.Send[VNTrackRectangleRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewTrackRectangleRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackRectangleRequest {
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:)
func NewTrackRectangleRequestWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest {
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRectangleObservation:"), observation)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// completionHandler: The block to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func NewTrackRectangleRequestWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler VNRequestCompletionHandler) VNTrackRectangleRequest {
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, completionHandler)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:)
func (t VNTrackRectangleRequest) InitWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("initWithRectangleObservation:"), observation)
	return rv
}
// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// completionHandler: The block to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func (t VNTrackRectangleRequest) InitWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler ErrorHandler) VNTrackRectangleRequest {
_block1, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, _block1)
	return VNTrackRectangleRequestFromID(rv)
}

// A constant for specifying revision 1 of the rectangling tracking request.
//
// See: https://developer.apple.com/documentation/vision/vntrackrectanglerequestrevision1
func (t VNTrackRectangleRequest) VNTrackRectangleRequestRevision1() int {
	rv := objc.Send[int](t.ID, objc.Sel("VNTrackRectangleRequestRevision1"))
	return rv
}

