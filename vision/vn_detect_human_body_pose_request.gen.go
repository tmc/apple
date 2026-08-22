// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNDetectHumanBodyPoseRequest] class.
var (
	_VNDetectHumanBodyPoseRequestClass     VNDetectHumanBodyPoseRequestClass
	_VNDetectHumanBodyPoseRequestClassOnce sync.Once
)

func getVNDetectHumanBodyPoseRequestClass() VNDetectHumanBodyPoseRequestClass {
	_VNDetectHumanBodyPoseRequestClassOnce.Do(func() {
		_VNDetectHumanBodyPoseRequestClass = VNDetectHumanBodyPoseRequestClass{class: objc.GetClass("VNDetectHumanBodyPoseRequest")}
	})
	return _VNDetectHumanBodyPoseRequestClass
}

// GetVNDetectHumanBodyPoseRequestClass returns the class object for VNDetectHumanBodyPoseRequest.
func GetVNDetectHumanBodyPoseRequestClass() VNDetectHumanBodyPoseRequestClass {
	return getVNDetectHumanBodyPoseRequestClass()
}

type VNDetectHumanBodyPoseRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectHumanBodyPoseRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectHumanBodyPoseRequestClass) Alloc() VNDetectHumanBodyPoseRequest {
	rv := objc.Send[VNDetectHumanBodyPoseRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects a human body pose.
//
// # Overview
//
// The framework provides the detected body pose as a
// [VNHumanBodyPoseObservation].
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest
type VNDetectHumanBodyPoseRequest struct {
	VNImageBasedRequest
}

// VNDetectHumanBodyPoseRequestFromID constructs a [VNDetectHumanBodyPoseRequest] from an objc.ID.
//
// A request that detects a human body pose.
func VNDetectHumanBodyPoseRequestFromID(id objc.ID) VNDetectHumanBodyPoseRequest {
	return VNDetectHumanBodyPoseRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNDetectHumanBodyPoseRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectHumanBodyPoseRequest] class.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest
type IVNDetectHumanBodyPoseRequest interface {
	IVNImageBasedRequest

	// Retrieves the supported joint names.
	SupportedJointNamesAndReturnError() ([]string, error)
	// Retrieves the supported joint group names.
	SupportedJointsGroupNamesAndReturnError() ([]string, error)
}

// Init initializes the instance.
func (d VNDetectHumanBodyPoseRequest) Init() VNDetectHumanBodyPoseRequest {
	rv := objc.Send[VNDetectHumanBodyPoseRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectHumanBodyPoseRequest) Autorelease() VNDetectHumanBodyPoseRequest {
	rv := objc.Send[VNDetectHumanBodyPoseRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectHumanBodyPoseRequest creates a new VNDetectHumanBodyPoseRequest instance.
func NewVNDetectHumanBodyPoseRequest() VNDetectHumanBodyPoseRequest {
	class := getVNDetectHumanBodyPoseRequestClass()
	rv := objc.Send[VNDetectHumanBodyPoseRequest](objc.ID(class.class), objc.Sel("new"))
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
// [VNImageRequestHandler.PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewDetectHumanBodyPoseRequestWithCompletionHandler(completionHandler VNRequestErrorHandler) VNDetectHumanBodyPoseRequest {
	_block0, _ := NewVNRequestErrorBlock(completionHandler)
	instance := getVNDetectHumanBodyPoseRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), _block0)
	return VNDetectHumanBodyPoseRequestFromID(rv)
}

// Retrieves the supported joint names.
//
// error: If an error occurs, an error object that describes the error; otherwise,
// `nil`.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointNamesAndReturnError:
func (d VNDetectHumanBodyPoseRequest) SupportedJointNamesAndReturnError() ([]string, error) {
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
// See: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d VNDetectHumanBodyPoseRequest) SupportedJointsGroupNamesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}
