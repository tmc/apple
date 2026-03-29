// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNDetectHumanBodyPose3DRequest] class.
var (
	_VNDetectHumanBodyPose3DRequestClass     VNDetectHumanBodyPose3DRequestClass
	_VNDetectHumanBodyPose3DRequestClassOnce sync.Once
)

func getVNDetectHumanBodyPose3DRequestClass() VNDetectHumanBodyPose3DRequestClass {
	_VNDetectHumanBodyPose3DRequestClassOnce.Do(func() {
		_VNDetectHumanBodyPose3DRequestClass = VNDetectHumanBodyPose3DRequestClass{class: objc.GetClass("VNDetectHumanBodyPose3DRequest")}
	})
	return _VNDetectHumanBodyPose3DRequestClass
}

// GetVNDetectHumanBodyPose3DRequestClass returns the class object for VNDetectHumanBodyPose3DRequest.
func GetVNDetectHumanBodyPose3DRequestClass() VNDetectHumanBodyPose3DRequestClass {
	return getVNDetectHumanBodyPose3DRequestClass()
}

type VNDetectHumanBodyPose3DRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectHumanBodyPose3DRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectHumanBodyPose3DRequestClass) Alloc() VNDetectHumanBodyPose3DRequest {
	rv := objc.Send[VNDetectHumanBodyPose3DRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects points on human bodies in 3D space, relative to the
// camera.
//
// # Overview
// 
// This request generates a collection of [VNHumanBodyPose3DObservation]
// objects that describe the position of each body the request detects. If the
// system allows it, the request uses [AVDepthData] information to improve the
// accuracy.
//
// [AVDepthData]: https://developer.apple.com/documentation/AVFoundation/AVDepthData
//
// # Determining Supported Joints
//
//   - [VNDetectHumanBodyPose3DRequest.SupportedJointsGroupNames]: Returns the joint names the request supports.
//   - [VNDetectHumanBodyPose3DRequest.SetSupportedJointsGroupNames]
//   - [VNDetectHumanBodyPose3DRequest.SupportedJointNames]: Returns the joint group names the request supports.
//   - [VNDetectHumanBodyPose3DRequest.SetSupportedJointNames]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest
type VNDetectHumanBodyPose3DRequest struct {
	VNStatefulRequest
}

// VNDetectHumanBodyPose3DRequestFromID constructs a [VNDetectHumanBodyPose3DRequest] from an objc.ID.
//
// A request that detects points on human bodies in 3D space, relative to the
// camera.
func VNDetectHumanBodyPose3DRequestFromID(id objc.ID) VNDetectHumanBodyPose3DRequest {
	return VNDetectHumanBodyPose3DRequest{VNStatefulRequest: VNStatefulRequestFromID(id)}
}
// NOTE: VNDetectHumanBodyPose3DRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectHumanBodyPose3DRequest] class.
//
// # Determining Supported Joints
//
//   - [IVNDetectHumanBodyPose3DRequest.SupportedJointsGroupNames]: Returns the joint names the request supports.
//   - [IVNDetectHumanBodyPose3DRequest.SetSupportedJointsGroupNames]
//   - [IVNDetectHumanBodyPose3DRequest.SupportedJointNames]: Returns the joint group names the request supports.
//   - [IVNDetectHumanBodyPose3DRequest.SetSupportedJointNames]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest
type IVNDetectHumanBodyPose3DRequest interface {
	IVNStatefulRequest

	// Topic: Determining Supported Joints

	// Returns the joint names the request supports.
	SupportedJointsGroupNames() VNHumanBodyPose3DObservationJointsGroupName
	SetSupportedJointsGroupNames(value VNHumanBodyPose3DObservationJointsGroupName)
	// Returns the joint group names the request supports.
	SupportedJointNames() VNHumanBodyPose3DObservationJointName
	SetSupportedJointNames(value VNHumanBodyPose3DObservationJointName)

	// Returns the joint names the request supports.
	SupportedJointNamesAndReturnError() ([]string, error)
	// Returns the joint group names the request supports.
	SupportedJointsGroupNamesAndReturnError() ([]string, error)
}

// Init initializes the instance.
func (d VNDetectHumanBodyPose3DRequest) Init() VNDetectHumanBodyPose3DRequest {
	rv := objc.Send[VNDetectHumanBodyPose3DRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectHumanBodyPose3DRequest) Autorelease() VNDetectHumanBodyPose3DRequest {
	rv := objc.Send[VNDetectHumanBodyPose3DRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectHumanBodyPose3DRequest creates a new VNDetectHumanBodyPose3DRequest instance.
func NewVNDetectHumanBodyPose3DRequest() VNDetectHumanBodyPose3DRequest {
	class := getVNDetectHumanBodyPose3DRequestClass()
	rv := objc.Send[VNDetectHumanBodyPose3DRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new 3D body pose request with a completion handler.
//
// completionHandler: The block to invoke after the request finishes processing.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/init(completionHandler:)
func NewDetectHumanBodyPose3DRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectHumanBodyPose3DRequest {
	instance := getVNDetectHumanBodyPose3DRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectHumanBodyPose3DRequestFromID(rv)
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
func NewDetectHumanBodyPose3DRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing objectivec.IObject, completionHandler VNRequestCompletionHandler) VNDetectHumanBodyPose3DRequest {
	instance := getVNDetectHumanBodyPose3DRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return VNDetectHumanBodyPose3DRequestFromID(rv)
}

// Returns the joint names the request supports.
//
// # Return Value
// 
// The array of joint name objects.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/supportedJointNamesAndReturnError:
func (d VNDetectHumanBodyPose3DRequest) SupportedJointNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}
// Returns the joint group names the request supports.
//
// # Return Value
// 
// The array of joint group name objects.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/supportedJointsGroupNamesAndReturnError:
func (d VNDetectHumanBodyPose3DRequest) SupportedJointsGroupNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Returns the joint names the request supports.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointsgroupnames
func (d VNDetectHumanBodyPose3DRequest) SupportedJointsGroupNames() VNHumanBodyPose3DObservationJointsGroupName {
	rv := objc.Send[VNHumanBodyPose3DObservationJointsGroupName](d.ID, objc.Sel("supportedJointsGroupNames"))
	return VNHumanBodyPose3DObservationJointsGroupName(rv)
}
func (d VNDetectHumanBodyPose3DRequest) SetSupportedJointsGroupNames(value VNHumanBodyPose3DObservationJointsGroupName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}
// Returns the joint group names the request supports.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointnames
func (d VNDetectHumanBodyPose3DRequest) SupportedJointNames() VNHumanBodyPose3DObservationJointName {
	rv := objc.Send[VNHumanBodyPose3DObservationJointName](d.ID, objc.Sel("supportedJointNames"))
	return VNHumanBodyPose3DObservationJointName(rv)
}
func (d VNDetectHumanBodyPose3DRequest) SetSupportedJointNames(value VNHumanBodyPose3DObservationJointName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointNames:"), value)
}

