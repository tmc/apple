// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNDetectHumanHandPoseRequest] class.
var (
	_VNDetectHumanHandPoseRequestClass     VNDetectHumanHandPoseRequestClass
	_VNDetectHumanHandPoseRequestClassOnce sync.Once
)

func getVNDetectHumanHandPoseRequestClass() VNDetectHumanHandPoseRequestClass {
	_VNDetectHumanHandPoseRequestClassOnce.Do(func() {
		_VNDetectHumanHandPoseRequestClass = VNDetectHumanHandPoseRequestClass{class: objc.GetClass("VNDetectHumanHandPoseRequest")}
	})
	return _VNDetectHumanHandPoseRequestClass
}

// GetVNDetectHumanHandPoseRequestClass returns the class object for VNDetectHumanHandPoseRequest.
func GetVNDetectHumanHandPoseRequestClass() VNDetectHumanHandPoseRequestClass {
	return getVNDetectHumanHandPoseRequestClass()
}

type VNDetectHumanHandPoseRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectHumanHandPoseRequestClass) Alloc() VNDetectHumanHandPoseRequest {
	rv := objc.Send[VNDetectHumanHandPoseRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects a human hand pose.
//
// # Overview
// 
// The framework provides the detected hand pose as a
// [VNIdentifiedPointsObservation].
//
// # Configuring the Request
//
//   - [VNDetectHumanHandPoseRequest.MaximumHandCount]: The maximum number of hands to detect in an image.
//   - [VNDetectHumanHandPoseRequest.SetMaximumHandCount]
//
// # Determining Supported Joints
//
//   - [VNDetectHumanHandPoseRequest.SupportedJointNames]: Retrieves the supported joint names.
//   - [VNDetectHumanHandPoseRequest.SetSupportedJointNames]
//   - [VNDetectHumanHandPoseRequest.SupportedJointsGroupNames]: Retrieves the supported joint group names.
//   - [VNDetectHumanHandPoseRequest.SetSupportedJointsGroupNames]
//
// # Identifying Hand Pose Revisions
//
//   - [VNDetectHumanHandPoseRequest.VNDetectHumanHandPoseRequestRevision1]: A constant for specifying revision 1 of the hand pose detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest
type VNDetectHumanHandPoseRequest struct {
	VNImageBasedRequest
}

// VNDetectHumanHandPoseRequestFromID constructs a [VNDetectHumanHandPoseRequest] from an objc.ID.
//
// A request that detects a human hand pose.
func VNDetectHumanHandPoseRequestFromID(id objc.ID) VNDetectHumanHandPoseRequest {
	return VNDetectHumanHandPoseRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectHumanHandPoseRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectHumanHandPoseRequest] class.
//
// # Configuring the Request
//
//   - [IVNDetectHumanHandPoseRequest.MaximumHandCount]: The maximum number of hands to detect in an image.
//   - [IVNDetectHumanHandPoseRequest.SetMaximumHandCount]
//
// # Determining Supported Joints
//
//   - [IVNDetectHumanHandPoseRequest.SupportedJointNames]: Retrieves the supported joint names.
//   - [IVNDetectHumanHandPoseRequest.SetSupportedJointNames]
//   - [IVNDetectHumanHandPoseRequest.SupportedJointsGroupNames]: Retrieves the supported joint group names.
//   - [IVNDetectHumanHandPoseRequest.SetSupportedJointsGroupNames]
//
// # Identifying Hand Pose Revisions
//
//   - [IVNDetectHumanHandPoseRequest.VNDetectHumanHandPoseRequestRevision1]: A constant for specifying revision 1 of the hand pose detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest
type IVNDetectHumanHandPoseRequest interface {
	IVNImageBasedRequest

	// Topic: Configuring the Request

	// The maximum number of hands to detect in an image.
	MaximumHandCount() uint
	SetMaximumHandCount(value uint)

	// Topic: Determining Supported Joints

	// Retrieves the supported joint names.
	SupportedJointNames() VNHumanHandPoseObservationJointName
	SetSupportedJointNames(value VNHumanHandPoseObservationJointName)
	// Retrieves the supported joint group names.
	SupportedJointsGroupNames() VNHumanHandPoseObservationJointsGroupName
	SetSupportedJointsGroupNames(value VNHumanHandPoseObservationJointsGroupName)

	// Topic: Identifying Hand Pose Revisions

	// A constant for specifying revision 1 of the hand pose detection request.
	VNDetectHumanHandPoseRequestRevision1() int

	// Retrieves the supported joint names.
	SupportedJointNamesAndReturnError() ([]string, error)
	// Retrieves the supported joint group names.
	SupportedJointsGroupNamesAndReturnError() ([]string, error)
}

// Init initializes the instance.
func (d VNDetectHumanHandPoseRequest) Init() VNDetectHumanHandPoseRequest {
	rv := objc.Send[VNDetectHumanHandPoseRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectHumanHandPoseRequest) Autorelease() VNDetectHumanHandPoseRequest {
	rv := objc.Send[VNDetectHumanHandPoseRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectHumanHandPoseRequest creates a new VNDetectHumanHandPoseRequest instance.
func NewVNDetectHumanHandPoseRequest() VNDetectHumanHandPoseRequest {
	class := getVNDetectHumanHandPoseRequestClass()
	rv := objc.Send[VNDetectHumanHandPoseRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectHumanHandPoseRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectHumanHandPoseRequest {
	instance := getVNDetectHumanHandPoseRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectHumanHandPoseRequestFromID(rv)
}

// Retrieves the supported joint names.
//
// error: If an error occurs, an error object that describes the error; otherwise,
// `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointNamesAndReturnError:
func (d VNDetectHumanHandPoseRequest) SupportedJointNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Retrieves the supported joint group names.
//
// error: If an error occurs, an error object that describes the error; otherwise,
// `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d VNDetectHumanHandPoseRequest) SupportedJointsGroupNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// The maximum number of hands to detect in an image.
//
// # Discussion
// 
// The request orders detected hands by relative size, with only the largest
// ones having key points determined.
// 
// The default value is 2.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/maximumHandCount
func (d VNDetectHumanHandPoseRequest) MaximumHandCount() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("maximumHandCount"))
	return rv
}
func (d VNDetectHumanHandPoseRequest) SetMaximumHandCount(value uint) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumHandCount:"), value)
}

// Retrieves the supported joint names.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d VNDetectHumanHandPoseRequest) SupportedJointNames() VNHumanHandPoseObservationJointName {
	rv := objc.Send[VNHumanHandPoseObservationJointName](d.ID, objc.Sel("supportedJointNames"))
	return VNHumanHandPoseObservationJointName(rv)
}
func (d VNDetectHumanHandPoseRequest) SetSupportedJointNames(value VNHumanHandPoseObservationJointName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointNames:"), value)
}

// Retrieves the supported joint group names.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d VNDetectHumanHandPoseRequest) SupportedJointsGroupNames() VNHumanHandPoseObservationJointsGroupName {
	rv := objc.Send[VNHumanHandPoseObservationJointsGroupName](d.ID, objc.Sel("supportedJointsGroupNames"))
	return VNHumanHandPoseObservationJointsGroupName(rv)
}
func (d VNDetectHumanHandPoseRequest) SetSupportedJointsGroupNames(value VNHumanHandPoseObservationJointsGroupName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}

// A constant for specifying revision 1 of the hand pose detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequestrevision1
func (d VNDetectHumanHandPoseRequest) VNDetectHumanHandPoseRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectHumanHandPoseRequestRevision1"))
	return rv
}

