// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
)

// The class instance for the [VNTrackHomographicImageRegistrationRequest] class.
var (
	_VNTrackHomographicImageRegistrationRequestClass     VNTrackHomographicImageRegistrationRequestClass
	_VNTrackHomographicImageRegistrationRequestClassOnce sync.Once
)

func getVNTrackHomographicImageRegistrationRequestClass() VNTrackHomographicImageRegistrationRequestClass {
	_VNTrackHomographicImageRegistrationRequestClassOnce.Do(func() {
		_VNTrackHomographicImageRegistrationRequestClass = VNTrackHomographicImageRegistrationRequestClass{class: objc.GetClass("VNTrackHomographicImageRegistrationRequest")}
	})
	return _VNTrackHomographicImageRegistrationRequestClass
}

// GetVNTrackHomographicImageRegistrationRequestClass returns the class object for VNTrackHomographicImageRegistrationRequest.
func GetVNTrackHomographicImageRegistrationRequestClass() VNTrackHomographicImageRegistrationRequestClass {
	return getVNTrackHomographicImageRegistrationRequestClass()
}

type VNTrackHomographicImageRegistrationRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNTrackHomographicImageRegistrationRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackHomographicImageRegistrationRequestClass) Alloc() VNTrackHomographicImageRegistrationRequest {
	rv := objc.Send[VNTrackHomographicImageRegistrationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request, as a stateful request you track over time, that
// determines the perspective warp matrix necessary to align the content of
// two images.
//
// # Overview
// 
// This request is similar to [VNHomographicImageRegistrationRequest].
// However, as a [VNStatefulRequest], it automatically computes the
// registration against the previous frame.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest
type VNTrackHomographicImageRegistrationRequest struct {
	VNStatefulRequest
}

// VNTrackHomographicImageRegistrationRequestFromID constructs a [VNTrackHomographicImageRegistrationRequest] from an objc.ID.
//
// An image-analysis request, as a stateful request you track over time, that
// determines the perspective warp matrix necessary to align the content of
// two images.
func VNTrackHomographicImageRegistrationRequestFromID(id objc.ID) VNTrackHomographicImageRegistrationRequest {
	return VNTrackHomographicImageRegistrationRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNTrackHomographicImageRegistrationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackHomographicImageRegistrationRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest
type IVNTrackHomographicImageRegistrationRequest interface {
	IVNStatefulRequest
}

// Init initializes the instance.
func (t VNTrackHomographicImageRegistrationRequest) Init() VNTrackHomographicImageRegistrationRequest {
	rv := objc.Send[VNTrackHomographicImageRegistrationRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackHomographicImageRegistrationRequest) Autorelease() VNTrackHomographicImageRegistrationRequest {
	rv := objc.Send[VNTrackHomographicImageRegistrationRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackHomographicImageRegistrationRequest creates a new VNTrackHomographicImageRegistrationRequest instance.
func NewVNTrackHomographicImageRegistrationRequest() VNTrackHomographicImageRegistrationRequest {
	class := getVNTrackHomographicImageRegistrationRequestClass()
	rv := objc.Send[VNTrackHomographicImageRegistrationRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new request that tracks the homographic transformation of two
// images, with a system callback on completion.
//
// completionHandler: The callback the system invokes when it completes the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest/init(completionHandler:)
func NewTrackHomographicImageRegistrationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackHomographicImageRegistrationRequest {
	instance := getVNTrackHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackHomographicImageRegistrationRequestFromID(rv)
}

// Initializes a video-based request.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// completionHandler: A closure that’s invoked after the request has completed its processing.
// The system invokes the completion handler on the same dispatch queue as the
// request performs its processing.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
func NewTrackHomographicImageRegistrationRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.CMTime, completionHandler VNRequestCompletionHandler) VNTrackHomographicImageRegistrationRequest {
	instance := getVNTrackHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNTrackHomographicImageRegistrationRequestFromID(rv)
}

