// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNStatefulRequest] class.
var (
	_VNStatefulRequestClass     VNStatefulRequestClass
	_VNStatefulRequestClassOnce sync.Once
)

func getVNStatefulRequestClass() VNStatefulRequestClass {
	_VNStatefulRequestClassOnce.Do(func() {
		_VNStatefulRequestClass = VNStatefulRequestClass{class: objc.GetClass("VNStatefulRequest")}
	})
	return _VNStatefulRequestClass
}

// GetVNStatefulRequestClass returns the class object for VNStatefulRequest.
func GetVNStatefulRequestClass() VNStatefulRequestClass {
	return getVNStatefulRequestClass()
}

type VNStatefulRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNStatefulRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNStatefulRequestClass) Alloc() VNStatefulRequest {
	rv := objc.Send[VNStatefulRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An abstract request type that builds evidence of a condition over time.
//
// # Initializing a Request
//
//   - [VNStatefulRequest.InitWithFrameAnalysisSpacingCompletionHandler]: Initializes a video-based request.
//
// # Configuring the Request
//
//   - [VNStatefulRequest.MinimumLatencyFrameCount]: The minimum number of frames a request processes before reporting an observation.
//   - [VNStatefulRequest.FrameAnalysisSpacing]: A time value that indicates the interval between analysis operations.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest
type VNStatefulRequest struct {
	VNImageBasedRequest
}

// VNStatefulRequestFromID constructs a [VNStatefulRequest] from an objc.ID.
//
// An abstract request type that builds evidence of a condition over time.
func VNStatefulRequestFromID(id objc.ID) VNStatefulRequest {
	return VNStatefulRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNStatefulRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNStatefulRequest] class.
//
// # Initializing a Request
//
//   - [IVNStatefulRequest.InitWithFrameAnalysisSpacingCompletionHandler]: Initializes a video-based request.
//
// # Configuring the Request
//
//   - [IVNStatefulRequest.MinimumLatencyFrameCount]: The minimum number of frames a request processes before reporting an observation.
//   - [IVNStatefulRequest.FrameAnalysisSpacing]: A time value that indicates the interval between analysis operations.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest
type IVNStatefulRequest interface {
	IVNImageBasedRequest

	// Topic: Initializing a Request

	// Initializes a video-based request.
	InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.CMTime, completionHandler ErrorHandler) VNStatefulRequest

	// Topic: Configuring the Request

	// The minimum number of frames a request processes before reporting an observation.
	MinimumLatencyFrameCount() int
	// A time value that indicates the interval between analysis operations.
	FrameAnalysisSpacing() coremedia.CMTime
}

// Init initializes the instance.
func (s VNStatefulRequest) Init() VNStatefulRequest {
	rv := objc.Send[VNStatefulRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VNStatefulRequest) Autorelease() VNStatefulRequest {
	rv := objc.Send[VNStatefulRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNStatefulRequest creates a new VNStatefulRequest instance.
func NewVNStatefulRequest() VNStatefulRequest {
	class := getVNStatefulRequestClass()
	rv := objc.Send[VNStatefulRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewStatefulRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNStatefulRequest {
	instance := getVNStatefulRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNStatefulRequestFromID(rv)
}

// Initializes a video-based request.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
//
// completionHandler: A closure that’s invoked after the request has completed its processing.
// The system invokes the completion handler on the same dispatch queue as the
// request performs its processing.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
//
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func NewStatefulRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.CMTime, completionHandler VNRequestCompletionHandler) VNStatefulRequest {
	instance := getVNStatefulRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNStatefulRequestFromID(rv)
}

// Initializes a video-based request.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
//
// completionHandler: A closure that’s invoked after the request has completed its processing.
// The system invokes the completion handler on the same dispatch queue as the
// request performs its processing.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
//
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (s VNStatefulRequest) InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.CMTime, completionHandler ErrorHandler) VNStatefulRequest {
	_block1, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, _block1)
	return VNStatefulRequestFromID(rv)
}

// The minimum number of frames a request processes before reporting an
// observation.
//
// # Discussion
//
// Video-based requests often need a minimum number of frames before they can
// report an observation. For example, a movement detection request requires a
// minimum of five frames before it can generate an observation. This value
// indicates how responsive a request is at processing incoming data.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/minimumLatencyFrameCount
func (s VNStatefulRequest) MinimumLatencyFrameCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("minimumLatencyFrameCount"))
	return rv
}

// A time value that indicates the interval between analysis operations.
//
// See: https://developer.apple.com/documentation/Vision/VNStatefulRequest/frameAnalysisSpacing
func (s VNStatefulRequest) FrameAnalysisSpacing() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("frameAnalysisSpacing"))
	return coremedia.CMTime(rv)
}
