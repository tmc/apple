// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNTrackingRequest] class.
var (
	_VNTrackingRequestClass     VNTrackingRequestClass
	_VNTrackingRequestClassOnce sync.Once
)

func getVNTrackingRequestClass() VNTrackingRequestClass {
	_VNTrackingRequestClassOnce.Do(func() {
		_VNTrackingRequestClass = VNTrackingRequestClass{class: objc.GetClass("VNTrackingRequest")}
	})
	return _VNTrackingRequestClass
}

// GetVNTrackingRequestClass returns the class object for VNTrackingRequest.
func GetVNTrackingRequestClass() VNTrackingRequestClass {
	return getVNTrackingRequestClass()
}

type VNTrackingRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackingRequestClass) Alloc() VNTrackingRequest {
	rv := objc.Send[VNTrackingRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The abstract superclass for image-analysis requests that track unique
// features across multiple images or video frames.
//
// # Overview
// 
// Instantiate a tracking request subclass to perform object tracking across
// multiple frames of an image. After initialization, configure the degree of
// accuracy by setting [VNTrackingRequest.TrackingLevel], and provide observations you’d like
// to track by setting the [VNTrackingRequest.InputObservation] initial bounding box.
//
// # Configuring a Tracking Request
//
//   - [VNTrackingRequest.InputObservation]: The observation object defining a region to track.
//   - [VNTrackingRequest.SetInputObservation]
//   - [VNTrackingRequest.TrackingLevel]: A value for specifying whether to prioritize speed or location accuracy.
//   - [VNTrackingRequest.SetTrackingLevel]
//   - [VNTrackingRequest.LastFrame]: A Boolean that indicates the last frame in a tracking sequence.
//   - [VNTrackingRequest.SetLastFrame]
//
// # Getting the Number of Trackers
//
//   - [VNTrackingRequest.SupportedNumberOfTrackersAndReturnError]: Returns the maximum number of simultaneous trackers for the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest
type VNTrackingRequest struct {
	VNImageBasedRequest
}

// VNTrackingRequestFromID constructs a [VNTrackingRequest] from an objc.ID.
//
// The abstract superclass for image-analysis requests that track unique
// features across multiple images or video frames.
func VNTrackingRequestFromID(id objc.ID) VNTrackingRequest {
	return VNTrackingRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNTrackingRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNTrackingRequest] class.
//
// # Configuring a Tracking Request
//
//   - [IVNTrackingRequest.InputObservation]: The observation object defining a region to track.
//   - [IVNTrackingRequest.SetInputObservation]
//   - [IVNTrackingRequest.TrackingLevel]: A value for specifying whether to prioritize speed or location accuracy.
//   - [IVNTrackingRequest.SetTrackingLevel]
//   - [IVNTrackingRequest.LastFrame]: A Boolean that indicates the last frame in a tracking sequence.
//   - [IVNTrackingRequest.SetLastFrame]
//
// # Getting the Number of Trackers
//
//   - [IVNTrackingRequest.SupportedNumberOfTrackersAndReturnError]: Returns the maximum number of simultaneous trackers for the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest
type IVNTrackingRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring a Tracking Request

	// The observation object defining a region to track.
	InputObservation() IVNDetectedObjectObservation
	SetInputObservation(value IVNDetectedObjectObservation)
	// A value for specifying whether to prioritize speed or location accuracy.
	TrackingLevel() VNRequestTrackingLevel
	SetTrackingLevel(value VNRequestTrackingLevel)
	// A Boolean that indicates the last frame in a tracking sequence.
	LastFrame() bool
	SetLastFrame(value bool)

	// Topic: Getting the Number of Trackers

	// Returns the maximum number of simultaneous trackers for the request.
	SupportedNumberOfTrackersAndReturnError() (uint, error)
}





// Init initializes the instance.
func (t VNTrackingRequest) Init() VNTrackingRequest {
	rv := objc.Send[VNTrackingRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackingRequest) Autorelease() VNTrackingRequest {
	rv := objc.Send[VNTrackingRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackingRequest creates a new VNTrackingRequest instance.
func NewVNTrackingRequest() VNTrackingRequest {
	class := getVNTrackingRequestClass()
	rv := objc.Send[VNTrackingRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewTrackingRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackingRequest {
	instance := getVNTrackingRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackingRequestFromID(rv)
}







// Returns the maximum number of simultaneous trackers for the request.
//
// error: An error that contains the reason why a failure occurs.
//
// # Return Value
// 
// The maximum number of trackers given a combination; or `0` if such
// combination doesn’t exist.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest/supportedNumber(ofTrackersAndReturnError:)
func (t VNTrackingRequest) SupportedNumberOfTrackersAndReturnError() (uint, error) {
			var errorPtr objc.ID
	rv := objc.Send[uint](t.ID, objc.Sel("supportedNumberOfTrackersAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}












// The observation object defining a region to track.
//
// # Discussion
// 
// Providing an observation not returned from a tracker, such as a
// user-defined observation, begins a new tracker for the sequence. Providing
// an observation that was returned from a tracker continues the use of that
// tracker, to track the region to the next frame.
// 
// In general, unless specified in the request’s documentation or header
// file, you must define the rectangle in normalized coordinates, with the
// origin at the lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest/inputObservation
func (t VNTrackingRequest) InputObservation() IVNDetectedObjectObservation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("inputObservation"))
	return VNDetectedObjectObservationFromID(objc.ID(rv))
}
func (t VNTrackingRequest) SetInputObservation(value IVNDetectedObjectObservation) {
	objc.Send[struct{}](t.ID, objc.Sel("setInputObservation:"), value)
}



// A value for specifying whether to prioritize speed or location accuracy.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest/trackingLevel
func (t VNTrackingRequest) TrackingLevel() VNRequestTrackingLevel {
	rv := objc.Send[VNRequestTrackingLevel](t.ID, objc.Sel("trackingLevel"))
	return VNRequestTrackingLevel(rv)
}
func (t VNTrackingRequest) SetTrackingLevel(value VNRequestTrackingLevel) {
	objc.Send[struct{}](t.ID, objc.Sel("setTrackingLevel:"), value)
}



// A Boolean that indicates the last frame in a tracking sequence.
//
// # Discussion
// 
// If set to [true], the current tracker will be released to the pool of
// available trackers when the current frame finishes processing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNTrackingRequest/isLastFrame
func (t VNTrackingRequest) LastFrame() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isLastFrame"))
	return rv
}
func (t VNTrackingRequest) SetLastFrame(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setLastFrame:"), value)
}
























