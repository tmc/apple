// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNTrackObjectRequest] class.
var (
	_VNTrackObjectRequestClass     VNTrackObjectRequestClass
	_VNTrackObjectRequestClassOnce sync.Once
)

func getVNTrackObjectRequestClass() VNTrackObjectRequestClass {
	_VNTrackObjectRequestClassOnce.Do(func() {
		_VNTrackObjectRequestClass = VNTrackObjectRequestClass{class: objc.GetClass("VNTrackObjectRequest")}
	})
	return _VNTrackObjectRequestClass
}

// GetVNTrackObjectRequestClass returns the class object for VNTrackObjectRequest.
func GetVNTrackObjectRequestClass() VNTrackObjectRequestClass {
	return getVNTrackObjectRequestClass()
}

type VNTrackObjectRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackObjectRequestClass) Alloc() VNTrackObjectRequest {
	rv := objc.Send[VNTrackObjectRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that tracks the movement of a previously
// identified object across multiple images or video frames.
//
// # Overview
// 
// Use this type of request to track the bounding boxes around objects
// previously identified in an image. Vision attempts to locate the same
// object from the input observation throughout the sequence.
//
// # Initializing an Object Tracking Request
//
//   - [VNTrackObjectRequest.InitWithDetectedObjectObservation]: Creates a new object tracking request with a detected object observation.
//   - [VNTrackObjectRequest.InitWithDetectedObjectObservationCompletionHandler]: Creates a new object tracking request with a detected object observation.
//
// # Identifying Request Revisions
//
//   - [VNTrackObjectRequest.VNTrackObjectRequestRevision2]: A constant for specifying revision 2 of the object tracking request.
//   - [VNTrackObjectRequest.VNTrackObjectRequestRevision1]: A constant for specifying revision 1 of the object tracking request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest
type VNTrackObjectRequest struct {
	VNTrackingRequest
}

// VNTrackObjectRequestFromID constructs a [VNTrackObjectRequest] from an objc.ID.
//
// An image-analysis request that tracks the movement of a previously
// identified object across multiple images or video frames.
func VNTrackObjectRequestFromID(id objc.ID) VNTrackObjectRequest {
	return VNTrackObjectRequest{VNTrackingRequest: VNTrackingRequestFromID(id)}
}
// NOTE: VNTrackObjectRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackObjectRequest] class.
//
// # Initializing an Object Tracking Request
//
//   - [IVNTrackObjectRequest.InitWithDetectedObjectObservation]: Creates a new object tracking request with a detected object observation.
//   - [IVNTrackObjectRequest.InitWithDetectedObjectObservationCompletionHandler]: Creates a new object tracking request with a detected object observation.
//
// # Identifying Request Revisions
//
//   - [IVNTrackObjectRequest.VNTrackObjectRequestRevision2]: A constant for specifying revision 2 of the object tracking request.
//   - [IVNTrackObjectRequest.VNTrackObjectRequestRevision1]: A constant for specifying revision 1 of the object tracking request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest
type IVNTrackObjectRequest interface {
	IVNTrackingRequest

	// Topic: Initializing an Object Tracking Request

	// Creates a new object tracking request with a detected object observation.
	InitWithDetectedObjectObservation(observation IVNDetectedObjectObservation) VNTrackObjectRequest
	// Creates a new object tracking request with a detected object observation.
	InitWithDetectedObjectObservationCompletionHandler(observation IVNDetectedObjectObservation, completionHandler ErrorHandler) VNTrackObjectRequest

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 2 of the object tracking request.
	VNTrackObjectRequestRevision2() int
	// A constant for specifying revision 1 of the object tracking request.
	VNTrackObjectRequestRevision1() int
}

// Init initializes the instance.
func (t VNTrackObjectRequest) Init() VNTrackObjectRequest {
	rv := objc.Send[VNTrackObjectRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackObjectRequest) Autorelease() VNTrackObjectRequest {
	rv := objc.Send[VNTrackObjectRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackObjectRequest creates a new VNTrackObjectRequest instance.
func NewVNTrackObjectRequest() VNTrackObjectRequest {
	class := getVNTrackObjectRequestClass()
	rv := objc.Send[VNTrackObjectRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewTrackObjectRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackObjectRequest {
	instance := getVNTrackObjectRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackObjectRequestFromID(rv)
}

// Creates a new object tracking request with a detected object observation.
//
// observation: A detected object observation with bounding box information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:)
func NewTrackObjectRequestWithDetectedObjectObservation(observation IVNDetectedObjectObservation) VNTrackObjectRequest {
	instance := getVNTrackObjectRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDetectedObjectObservation:"), observation)
	return VNTrackObjectRequestFromID(rv)
}

// Creates a new object tracking request with a detected object observation.
//
// observation: A detected object observation with bounding box information.
//
// completionHandler: The callback to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:completionHandler:)
func NewTrackObjectRequestWithDetectedObjectObservationCompletionHandler(observation IVNDetectedObjectObservation, completionHandler VNRequestCompletionHandler) VNTrackObjectRequest {
	instance := getVNTrackObjectRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDetectedObjectObservation:completionHandler:"), observation, completionHandler)
	return VNTrackObjectRequestFromID(rv)
}

// Creates a new object tracking request with a detected object observation.
//
// observation: A detected object observation with bounding box information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:)
func (t VNTrackObjectRequest) InitWithDetectedObjectObservation(observation IVNDetectedObjectObservation) VNTrackObjectRequest {
	rv := objc.Send[VNTrackObjectRequest](t.ID, objc.Sel("initWithDetectedObjectObservation:"), observation)
	return rv
}

// Creates a new object tracking request with a detected object observation.
//
// observation: A detected object observation with bounding box information.
//
// completionHandler: The callback to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:completionHandler:)
func (t VNTrackObjectRequest) InitWithDetectedObjectObservationCompletionHandler(observation IVNDetectedObjectObservation, completionHandler ErrorHandler) VNTrackObjectRequest {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithDetectedObjectObservation:completionHandler:"), observation, _block1)
	return VNTrackObjectRequestFromID(rv)
}

// A constant for specifying revision 2 of the object tracking request.
//
// See: https://developer.apple.com/documentation/vision/vntrackobjectrequestrevision2
func (t VNTrackObjectRequest) VNTrackObjectRequestRevision2() int {
	rv := objc.Send[int](t.ID, objc.Sel("VNTrackObjectRequestRevision2"))
	return rv
}

// A constant for specifying revision 1 of the object tracking request.
//
// See: https://developer.apple.com/documentation/vision/vntrackobjectrequestrevision1
func (t VNTrackObjectRequest) VNTrackObjectRequestRevision1() int {
	rv := objc.Send[int](t.ID, objc.Sel("VNTrackObjectRequestRevision1"))
	return rv
}

