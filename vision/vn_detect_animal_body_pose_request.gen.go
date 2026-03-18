// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNDetectAnimalBodyPoseRequest] class.
var (
	_VNDetectAnimalBodyPoseRequestClass     VNDetectAnimalBodyPoseRequestClass
	_VNDetectAnimalBodyPoseRequestClassOnce sync.Once
)

func getVNDetectAnimalBodyPoseRequestClass() VNDetectAnimalBodyPoseRequestClass {
	_VNDetectAnimalBodyPoseRequestClassOnce.Do(func() {
		_VNDetectAnimalBodyPoseRequestClass = VNDetectAnimalBodyPoseRequestClass{class: objc.GetClass("VNDetectAnimalBodyPoseRequest")}
	})
	return _VNDetectAnimalBodyPoseRequestClass
}

// GetVNDetectAnimalBodyPoseRequestClass returns the class object for VNDetectAnimalBodyPoseRequest.
func GetVNDetectAnimalBodyPoseRequestClass() VNDetectAnimalBodyPoseRequestClass {
	return getVNDetectAnimalBodyPoseRequestClass()
}

type VNDetectAnimalBodyPoseRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectAnimalBodyPoseRequestClass) Alloc() VNDetectAnimalBodyPoseRequest {
	rv := objc.Send[VNDetectAnimalBodyPoseRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects an animal body pose.
//
// # Determining Supported Joints
//
//   - [VNDetectAnimalBodyPoseRequest.SupportedJointNames]: Retrieves the joint names the request supports.
//   - [VNDetectAnimalBodyPoseRequest.SetSupportedJointNames]
//   - [VNDetectAnimalBodyPoseRequest.SupportedJointsGroupNames]: Retrieves the joint group names the request supports.
//   - [VNDetectAnimalBodyPoseRequest.SetSupportedJointsGroupNames]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest
type VNDetectAnimalBodyPoseRequest struct {
	VNImageBasedRequest
}

// VNDetectAnimalBodyPoseRequestFromID constructs a [VNDetectAnimalBodyPoseRequest] from an objc.ID.
//
// A request that detects an animal body pose.
func VNDetectAnimalBodyPoseRequestFromID(id objc.ID) VNDetectAnimalBodyPoseRequest {
	return VNDetectAnimalBodyPoseRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectAnimalBodyPoseRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectAnimalBodyPoseRequest] class.
//
// # Determining Supported Joints
//
//   - [IVNDetectAnimalBodyPoseRequest.SupportedJointNames]: Retrieves the joint names the request supports.
//   - [IVNDetectAnimalBodyPoseRequest.SetSupportedJointNames]
//   - [IVNDetectAnimalBodyPoseRequest.SupportedJointsGroupNames]: Retrieves the joint group names the request supports.
//   - [IVNDetectAnimalBodyPoseRequest.SetSupportedJointsGroupNames]
//
// See: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest
type IVNDetectAnimalBodyPoseRequest interface {
	IVNImageBasedRequest

	// Topic: Determining Supported Joints

	// Retrieves the joint names the request supports.
	SupportedJointNames() VNAnimalBodyPoseObservationJointName
	SetSupportedJointNames(value VNAnimalBodyPoseObservationJointName)
	// Retrieves the joint group names the request supports.
	SupportedJointsGroupNames() VNAnimalBodyPoseObservationJointsGroupName
	SetSupportedJointsGroupNames(value VNAnimalBodyPoseObservationJointsGroupName)

	// Retrieves the joint names the request supports.
	SupportedJointNamesAndReturnError() ([]string, error)
	// Retrieves the joint group names the request supports.
	SupportedJointsGroupNamesAndReturnError() ([]string, error)
}

// Init initializes the instance.
func (d VNDetectAnimalBodyPoseRequest) Init() VNDetectAnimalBodyPoseRequest {
	rv := objc.Send[VNDetectAnimalBodyPoseRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectAnimalBodyPoseRequest) Autorelease() VNDetectAnimalBodyPoseRequest {
	rv := objc.Send[VNDetectAnimalBodyPoseRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectAnimalBodyPoseRequest creates a new VNDetectAnimalBodyPoseRequest instance.
func NewVNDetectAnimalBodyPoseRequest() VNDetectAnimalBodyPoseRequest {
	class := getVNDetectAnimalBodyPoseRequestClass()
	rv := objc.Send[VNDetectAnimalBodyPoseRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewDetectAnimalBodyPoseRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectAnimalBodyPoseRequest {
	instance := getVNDetectAnimalBodyPoseRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectAnimalBodyPoseRequestFromID(rv)
}

// Retrieves the joint names the request supports.
//
// error: If an error occurs, an object that describes the error; otherwise, `nil`.
//
// # Return Value
// 
// The array of joint names.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointNamesAndReturnError:
func (d VNDetectAnimalBodyPoseRequest) SupportedJointNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Retrieves the joint group names the request supports.
//
// error: If an error occurs, an object that describes the error; otherwise, `nil`.
//
// # Return Value
// 
// The array of joint group names.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d VNDetectAnimalBodyPoseRequest) SupportedJointsGroupNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// Retrieves the joint names the request supports.
//
// See: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointnames
func (d VNDetectAnimalBodyPoseRequest) SupportedJointNames() VNAnimalBodyPoseObservationJointName {
	rv := objc.Send[VNAnimalBodyPoseObservationJointName](d.ID, objc.Sel("supportedJointNames"))
	return VNAnimalBodyPoseObservationJointName(rv)
}
func (d VNDetectAnimalBodyPoseRequest) SetSupportedJointNames(value VNAnimalBodyPoseObservationJointName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointNames:"), value)
}

// Retrieves the joint group names the request supports.
//
// See: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointsgroupnames
func (d VNDetectAnimalBodyPoseRequest) SupportedJointsGroupNames() VNAnimalBodyPoseObservationJointsGroupName {
	rv := objc.Send[VNAnimalBodyPoseObservationJointsGroupName](d.ID, objc.Sel("supportedJointsGroupNames"))
	return VNAnimalBodyPoseObservationJointsGroupName(rv)
}
func (d VNDetectAnimalBodyPoseRequest) SetSupportedJointsGroupNames(value VNAnimalBodyPoseObservationJointsGroupName) {
	objc.Send[struct{}](d.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}

