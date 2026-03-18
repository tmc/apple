// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNTrackOpticalFlowRequest] class.
var (
	_VNTrackOpticalFlowRequestClass     VNTrackOpticalFlowRequestClass
	_VNTrackOpticalFlowRequestClassOnce sync.Once
)

func getVNTrackOpticalFlowRequestClass() VNTrackOpticalFlowRequestClass {
	_VNTrackOpticalFlowRequestClassOnce.Do(func() {
		_VNTrackOpticalFlowRequestClass = VNTrackOpticalFlowRequestClass{class: objc.GetClass("VNTrackOpticalFlowRequest")}
	})
	return _VNTrackOpticalFlowRequestClass
}

// GetVNTrackOpticalFlowRequestClass returns the class object for VNTrackOpticalFlowRequest.
func GetVNTrackOpticalFlowRequestClass() VNTrackOpticalFlowRequestClass {
	return getVNTrackOpticalFlowRequestClass()
}

type VNTrackOpticalFlowRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackOpticalFlowRequestClass) Alloc() VNTrackOpticalFlowRequest {
	rv := objc.Send[VNTrackOpticalFlowRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that determines the direction change of vectors for each pixel
// from a previous to current image.
//
// # Overview
// 
// This request works at the pixel level, so both images must have the same
// dimensions to successfully perform the request.
// 
// Setting a region of interest isolates where to perform the change
// determination.
//
// # Configuring the Request
//
//   - [VNTrackOpticalFlowRequest.ComputationAccuracy]: The level of accuracy to compute the optical flow.
//   - [VNTrackOpticalFlowRequest.SetComputationAccuracy]
//   - [VNTrackOpticalFlowRequest.KeepNetworkOutput]: A Boolean value that indicates the raw pixel buffer continues to emit from the network.
//   - [VNTrackOpticalFlowRequest.SetKeepNetworkOutput]
//   - [VNTrackOpticalFlowRequest.OutputPixelFormat]: The pixel format type of the output value.
//   - [VNTrackOpticalFlowRequest.SetOutputPixelFormat]
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest
type VNTrackOpticalFlowRequest struct {
	VNStatefulRequest
}

// VNTrackOpticalFlowRequestFromID constructs a [VNTrackOpticalFlowRequest] from an objc.ID.
//
// An object that determines the direction change of vectors for each pixel
// from a previous to current image.
func VNTrackOpticalFlowRequestFromID(id objc.ID) VNTrackOpticalFlowRequest {
	return VNTrackOpticalFlowRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNTrackOpticalFlowRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackOpticalFlowRequest] class.
//
// # Configuring the Request
//
//   - [IVNTrackOpticalFlowRequest.ComputationAccuracy]: The level of accuracy to compute the optical flow.
//   - [IVNTrackOpticalFlowRequest.SetComputationAccuracy]
//   - [IVNTrackOpticalFlowRequest.KeepNetworkOutput]: A Boolean value that indicates the raw pixel buffer continues to emit from the network.
//   - [IVNTrackOpticalFlowRequest.SetKeepNetworkOutput]
//   - [IVNTrackOpticalFlowRequest.OutputPixelFormat]: The pixel format type of the output value.
//   - [IVNTrackOpticalFlowRequest.SetOutputPixelFormat]
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest
type IVNTrackOpticalFlowRequest interface {
	IVNStatefulRequest

	// Topic: Configuring the Request

	// The level of accuracy to compute the optical flow.
	ComputationAccuracy() VNTrackOpticalFlowRequestComputationAccuracy
	SetComputationAccuracy(value VNTrackOpticalFlowRequestComputationAccuracy)
	// A Boolean value that indicates the raw pixel buffer continues to emit from the network.
	KeepNetworkOutput() bool
	SetKeepNetworkOutput(value bool)
	// The pixel format type of the output value.
	OutputPixelFormat() uint32
	SetOutputPixelFormat(value uint32)
}

// Init initializes the instance.
func (t VNTrackOpticalFlowRequest) Init() VNTrackOpticalFlowRequest {
	rv := objc.Send[VNTrackOpticalFlowRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackOpticalFlowRequest) Autorelease() VNTrackOpticalFlowRequest {
	rv := objc.Send[VNTrackOpticalFlowRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackOpticalFlowRequest creates a new VNTrackOpticalFlowRequest instance.
func NewVNTrackOpticalFlowRequest() VNTrackOpticalFlowRequest {
	class := getVNTrackOpticalFlowRequestClass()
	rv := objc.Send[VNTrackOpticalFlowRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new request that tracks the optical from one image to another,
// with a system callback on completion.
//
// completionHandler: The callback the system invokes when it completes the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/init(completionHandler:)
func NewTrackOpticalFlowRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNTrackOpticalFlowRequest {
	instance := getVNTrackOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNTrackOpticalFlowRequestFromID(rv)
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
func NewTrackOpticalFlowRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing objectivec.IObject, completionHandler VNRequestCompletionHandler) VNTrackOpticalFlowRequest {
	instance := getVNTrackOpticalFlowRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNTrackOpticalFlowRequestFromID(rv)
}

// The level of accuracy to compute the optical flow.
//
// # Discussion
// 
// The computational time trends with accuracy level. The default value is
// [TrackOpticalFlowRequestComputationAccuracyMedium].
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/computationAccuracy-swift.property
func (t VNTrackOpticalFlowRequest) ComputationAccuracy() VNTrackOpticalFlowRequestComputationAccuracy {
	rv := objc.Send[VNTrackOpticalFlowRequestComputationAccuracy](t.ID, objc.Sel("computationAccuracy"))
	return VNTrackOpticalFlowRequestComputationAccuracy(rv)
}
func (t VNTrackOpticalFlowRequest) SetComputationAccuracy(value VNTrackOpticalFlowRequestComputationAccuracy) {
	objc.Send[struct{}](t.ID, objc.Sel("setComputationAccuracy:"), value)
}

// A Boolean value that indicates the raw pixel buffer continues to emit from
// the network.
//
// # Discussion
// 
// The default value is [false]; otherwise, the request ignores
// [OutputPixelFormat].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/keepNetworkOutput
func (t VNTrackOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("keepNetworkOutput"))
	return rv
}
func (t VNTrackOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setKeepNetworkOutput:"), value)
}

// The pixel format type of the output value.
//
// # Discussion
// 
// The valid values are [kCVPixelFormatType_TwoComponent32Float] and
// [kCVPixelFormatType_TwoComponent16Half]. The default value is
// [kCVPixelFormatType_TwoComponent32Float].
//
// [kCVPixelFormatType_TwoComponent16Half]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_TwoComponent16Half
// [kCVPixelFormatType_TwoComponent32Float]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_TwoComponent32Float
//
// See: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/outputPixelFormat
func (t VNTrackOpticalFlowRequest) OutputPixelFormat() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("outputPixelFormat"))
	return rv
}
func (t VNTrackOpticalFlowRequest) SetOutputPixelFormat(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setOutputPixelFormat:"), value)
}

