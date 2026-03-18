// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNDetectTrajectoriesRequest] class.
var (
	_VNDetectTrajectoriesRequestClass     VNDetectTrajectoriesRequestClass
	_VNDetectTrajectoriesRequestClassOnce sync.Once
)

func getVNDetectTrajectoriesRequestClass() VNDetectTrajectoriesRequestClass {
	_VNDetectTrajectoriesRequestClassOnce.Do(func() {
		_VNDetectTrajectoriesRequestClass = VNDetectTrajectoriesRequestClass{class: objc.GetClass("VNDetectTrajectoriesRequest")}
	})
	return _VNDetectTrajectoriesRequestClass
}

// GetVNDetectTrajectoriesRequestClass returns the class object for VNDetectTrajectoriesRequest.
func GetVNDetectTrajectoriesRequestClass() VNDetectTrajectoriesRequestClass {
	return getVNDetectTrajectoriesRequestClass()
}

type VNDetectTrajectoriesRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectTrajectoriesRequestClass) Alloc() VNDetectTrajectoriesRequest {
	rv := objc.Send[VNDetectTrajectoriesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects the trajectories of shapes moving along a parabolic
// path.
//
// # Overview
// 
// After the request detects a trajectory, it produces an observation that
// contains the shape’s detected points and an equation describing the
// parabola.
//
// # Creating a Request
//
//   - [VNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler]: Creates a new request to detect trajectories.
//
// # Configuring the Request
//
//   - [VNDetectTrajectoriesRequest.TargetFrameTime]: The requested target frame time for processing trajectory detection.
//   - [VNDetectTrajectoriesRequest.SetTargetFrameTime]
//   - [VNDetectTrajectoriesRequest.TrajectoryLength]: The number of points to detect before calculating a trajectory.
//   - [VNDetectTrajectoriesRequest.ObjectMinimumNormalizedRadius]: The minimum radius of the bounding circle of the object to track.
//   - [VNDetectTrajectoriesRequest.SetObjectMinimumNormalizedRadius]
//   - [VNDetectTrajectoriesRequest.ObjectMaximumNormalizedRadius]: The maximum radius of the bounding circle of the object to track.
//   - [VNDetectTrajectoriesRequest.SetObjectMaximumNormalizedRadius]
//
// # Identifying Request Revisions
//
//   - [VNDetectTrajectoriesRequest.VNDetectTrajectoriesRequestRevision1]: A constant for specifying revision 1 of the trajectories detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest
type VNDetectTrajectoriesRequest struct {
	VNStatefulRequest
}

// VNDetectTrajectoriesRequestFromID constructs a [VNDetectTrajectoriesRequest] from an objc.ID.
//
// A request that detects the trajectories of shapes moving along a parabolic
// path.
func VNDetectTrajectoriesRequestFromID(id objc.ID) VNDetectTrajectoriesRequest {
	return VNDetectTrajectoriesRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNDetectTrajectoriesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectTrajectoriesRequest] class.
//
// # Creating a Request
//
//   - [IVNDetectTrajectoriesRequest.InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler]: Creates a new request to detect trajectories.
//
// # Configuring the Request
//
//   - [IVNDetectTrajectoriesRequest.TargetFrameTime]: The requested target frame time for processing trajectory detection.
//   - [IVNDetectTrajectoriesRequest.SetTargetFrameTime]
//   - [IVNDetectTrajectoriesRequest.TrajectoryLength]: The number of points to detect before calculating a trajectory.
//   - [IVNDetectTrajectoriesRequest.ObjectMinimumNormalizedRadius]: The minimum radius of the bounding circle of the object to track.
//   - [IVNDetectTrajectoriesRequest.SetObjectMinimumNormalizedRadius]
//   - [IVNDetectTrajectoriesRequest.ObjectMaximumNormalizedRadius]: The maximum radius of the bounding circle of the object to track.
//   - [IVNDetectTrajectoriesRequest.SetObjectMaximumNormalizedRadius]
//
// # Identifying Request Revisions
//
//   - [IVNDetectTrajectoriesRequest.VNDetectTrajectoriesRequestRevision1]: A constant for specifying revision 1 of the trajectories detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest
type IVNDetectTrajectoriesRequest interface {
	IVNStatefulRequest

	// Topic: Creating a Request

	// Creates a new request to detect trajectories.
	InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing objectivec.IObject, trajectoryLength int, completionHandler ErrorHandler) VNDetectTrajectoriesRequest

	// Topic: Configuring the Request

	// The requested target frame time for processing trajectory detection.
	TargetFrameTime() objectivec.IObject
	SetTargetFrameTime(value objectivec.IObject)
	// The number of points to detect before calculating a trajectory.
	TrajectoryLength() int
	// The minimum radius of the bounding circle of the object to track.
	ObjectMinimumNormalizedRadius() float32
	SetObjectMinimumNormalizedRadius(value float32)
	// The maximum radius of the bounding circle of the object to track.
	ObjectMaximumNormalizedRadius() float32
	SetObjectMaximumNormalizedRadius(value float32)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 1 of the trajectories detection request.
	VNDetectTrajectoriesRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectTrajectoriesRequest) Init() VNDetectTrajectoriesRequest {
	rv := objc.Send[VNDetectTrajectoriesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectTrajectoriesRequest) Autorelease() VNDetectTrajectoriesRequest {
	rv := objc.Send[VNDetectTrajectoriesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectTrajectoriesRequest creates a new VNDetectTrajectoriesRequest instance.
func NewVNDetectTrajectoriesRequest() VNDetectTrajectoriesRequest {
	class := getVNDetectTrajectoriesRequestClass()
	rv := objc.Send[VNDetectTrajectoriesRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectTrajectoriesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectTrajectoriesRequest {
	instance := getVNDetectTrajectoriesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectTrajectoriesRequestFromID(rv)
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
func NewDetectTrajectoriesRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing objectivec.IObject, completionHandler VNRequestCompletionHandler) VNDetectTrajectoriesRequest {
	instance := getVNDetectTrajectoriesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNDetectTrajectoriesRequestFromID(rv)
}

// Creates a new request to detect trajectories.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// trajectoryLength: The number of points required to analyze to determine that a shape follows
// a parabolic path. This argument value must be at least 5.
//
// completionHandler: A closure that’s invoked after the request completes its processing. The
// system invokes the completion handler on the same dispatch queue that the
// request uses to perform its processing.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/init(frameAnalysisSpacing:trajectoryLength:completionHandler:)
// frameAnalysisSpacing is a [coremedia.CMTime].
func NewDetectTrajectoriesRequestWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing objectivec.IObject, trajectoryLength int, completionHandler VNRequestCompletionHandler) VNDetectTrajectoriesRequest {
	instance := getVNDetectTrajectoriesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:trajectoryLength:completionHandler:"), frameAnalysisSpacing, trajectoryLength, completionHandler)
	return VNDetectTrajectoriesRequestFromID(rv)
}

// Creates a new request to detect trajectories.
//
// frameAnalysisSpacing: A [CMTime] value that indicates the duration between analysis operations.
// Increase this value to reduce the number of frames analyzed on slower
// devices. Set this argument to [zero] to analyze all frames.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// trajectoryLength: The number of points required to analyze to determine that a shape follows
// a parabolic path. This argument value must be at least 5.
//
// completionHandler: A closure that’s invoked after the request completes its processing. The
// system invokes the completion handler on the same dispatch queue that the
// request uses to perform its processing.
//
// frameAnalysisSpacing is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/init(frameAnalysisSpacing:trajectoryLength:completionHandler:)
func (d VNDetectTrajectoriesRequest) InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing objectivec.IObject, trajectoryLength int, completionHandler ErrorHandler) VNDetectTrajectoriesRequest {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithFrameAnalysisSpacing:trajectoryLength:completionHandler:"), frameAnalysisSpacing, trajectoryLength, _block2)
	return VNDetectTrajectoriesRequestFromID(rv)
}

// The requested target frame time for processing trajectory detection.
//
// # Discussion
// 
// Use this property value for real-time processing of frames, which requires
// execution within a specific amount of time. The request evaluates from
// frame-to-frame. If processing takes longer than the targeted time for the
// current frame, it attempts to decrease the overall time by reducing the
// accuracy (down to a set minimum) for the next frame. If a frame takes less
// time than the targeted time, the request increases the accuracy (up to a
// set maximum) of the next frame.
// 
// The default value is [indefinite], which indicates that accuracy stays at
// the predefined maximum.
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/targetFrameTime
func (d VNDetectTrajectoriesRequest) TargetFrameTime() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("targetFrameTime"))
	return objectivec.Object{ID: rv}
}
func (d VNDetectTrajectoriesRequest) SetTargetFrameTime(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setTargetFrameTime:"), value)
}

// The number of points to detect before calculating a trajectory.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/trajectoryLength
func (d VNDetectTrajectoriesRequest) TrajectoryLength() int {
	rv := objc.Send[int](d.ID, objc.Sel("trajectoryLength"))
	return rv
}

// The minimum radius of the bounding circle of the object to track.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMinimumNormalizedRadius
func (d VNDetectTrajectoriesRequest) ObjectMinimumNormalizedRadius() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("objectMinimumNormalizedRadius"))
	return rv
}
func (d VNDetectTrajectoriesRequest) SetObjectMinimumNormalizedRadius(value float32) {
	objc.Send[struct{}](d.ID, objc.Sel("setObjectMinimumNormalizedRadius:"), value)
}

// The maximum radius of the bounding circle of the object to track.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMaximumNormalizedRadius
func (d VNDetectTrajectoriesRequest) ObjectMaximumNormalizedRadius() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("objectMaximumNormalizedRadius"))
	return rv
}
func (d VNDetectTrajectoriesRequest) SetObjectMaximumNormalizedRadius(value float32) {
	objc.Send[struct{}](d.ID, objc.Sel("setObjectMaximumNormalizedRadius:"), value)
}

// A constant for specifying revision 1 of the trajectories detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequestrevision1
func (d VNDetectTrajectoriesRequest) VNDetectTrajectoriesRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectTrajectoriesRequestRevision1"))
	return rv
}

