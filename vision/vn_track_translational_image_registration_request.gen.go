// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNTrackTranslationalImageRegistrationRequest] class.
var (
	_VNTrackTranslationalImageRegistrationRequestClass     VNTrackTranslationalImageRegistrationRequestClass
	_VNTrackTranslationalImageRegistrationRequestClassOnce sync.Once
)

func getVNTrackTranslationalImageRegistrationRequestClass() VNTrackTranslationalImageRegistrationRequestClass {
	_VNTrackTranslationalImageRegistrationRequestClassOnce.Do(func() {
		_VNTrackTranslationalImageRegistrationRequestClass = VNTrackTranslationalImageRegistrationRequestClass{class: objc.GetClass("VNTrackTranslationalImageRegistrationRequest")}
	})
	return _VNTrackTranslationalImageRegistrationRequestClass
}

// GetVNTrackTranslationalImageRegistrationRequestClass returns the class object for VNTrackTranslationalImageRegistrationRequest.
func GetVNTrackTranslationalImageRegistrationRequestClass() VNTrackTranslationalImageRegistrationRequestClass {
	return getVNTrackTranslationalImageRegistrationRequestClass()
}

type VNTrackTranslationalImageRegistrationRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackTranslationalImageRegistrationRequestClass) Alloc() VNTrackTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTrackTranslationalImageRegistrationRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request, as a stateful request you track over time, that
// determines the affine transform necessary to align the content of two
// images.
//
// # Overview
// 
// This request is similar to [VNTranslationalImageRegistrationRequest].
// However, as a [VNStatefulRequest], it automatically computes the
// registration against the previous frame.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest
type VNTrackTranslationalImageRegistrationRequest struct {
	VNStatefulRequest
}

// VNTrackTranslationalImageRegistrationRequestFromID constructs a [VNTrackTranslationalImageRegistrationRequest] from an objc.ID.
//
// An image-analysis request, as a stateful request you track over time, that
// determines the affine transform necessary to align the content of two
// images.
func VNTrackTranslationalImageRegistrationRequestFromID(id objc.ID) VNTrackTranslationalImageRegistrationRequest {
	return VNTrackTranslationalImageRegistrationRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNTrackTranslationalImageRegistrationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackTranslationalImageRegistrationRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest
type IVNTrackTranslationalImageRegistrationRequest interface {
	IVNStatefulRequest
}

// Init initializes the instance.
func (t VNTrackTranslationalImageRegistrationRequest) Init() VNTrackTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTrackTranslationalImageRegistrationRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackTranslationalImageRegistrationRequest) Autorelease() VNTrackTranslationalImageRegistrationRequest {
	rv := objc.Send[VNTrackTranslationalImageRegistrationRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackTranslationalImageRegistrationRequest creates a new VNTrackTranslationalImageRegistrationRequest instance.
func NewVNTrackTranslationalImageRegistrationRequest() VNTrackTranslationalImageRegistrationRequest {
	class := getVNTrackTranslationalImageRegistrationRequestClass()
	rv := objc.Send[VNTrackTranslationalImageRegistrationRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new request that tracks the translational registration of two
// images, with a system callback on completion.
//
// completionHandler: The callback the system invokes when it completes the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest/init(completionHandler:)
func NewTrackTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackTranslationalImageRegistrationRequest {
	instance := getVNTrackTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackTranslationalImageRegistrationRequestFromID(rv)
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
// frameAnalysisSpacing is a [coremedia.CMTime].
func NewTrackTranslationalImageRegistrationRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing objectivec.IObject, completionHandler VNRequestCompletionHandler) VNTrackTranslationalImageRegistrationRequest {
	instance := getVNTrackTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNTrackTranslationalImageRegistrationRequestFromID(rv)
}

