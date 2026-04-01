// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNHumanBodyPose3DObservation] class.
var (
	_VNHumanBodyPose3DObservationClass     VNHumanBodyPose3DObservationClass
	_VNHumanBodyPose3DObservationClassOnce sync.Once
)

func getVNHumanBodyPose3DObservationClass() VNHumanBodyPose3DObservationClass {
	_VNHumanBodyPose3DObservationClassOnce.Do(func() {
		_VNHumanBodyPose3DObservationClass = VNHumanBodyPose3DObservationClass{class: objc.GetClass("VNHumanBodyPose3DObservation")}
	})
	return _VNHumanBodyPose3DObservationClass
}

// GetVNHumanBodyPose3DObservationClass returns the class object for VNHumanBodyPose3DObservation.
func GetVNHumanBodyPose3DObservationClass() VNHumanBodyPose3DObservationClass {
	return getVNHumanBodyPose3DObservationClass()
}

type VNHumanBodyPose3DObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNHumanBodyPose3DObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHumanBodyPose3DObservationClass) Alloc() VNHumanBodyPose3DObservation {
	rv := objc.Send[VNHumanBodyPose3DObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that provides the 3D body points the request recognizes.
//
// # Accessing Points
//
//   - [VNHumanBodyPose3DObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [VNHumanBodyPose3DObservation.AvailableJointsGroupNames]: The available joint group names in the observation.
//   - [VNHumanBodyPose3DObservation.RecognizedPointForJointNameError]: Returns the point for a joint name that the observation recognizes.
//   - [VNHumanBodyPose3DObservation.RecognizedPointsForJointsGroupNameError]: Returns a collection of points for the group name you specify.
//
// # Getting the Joint Position
//
//   - [VNHumanBodyPose3DObservation.PointInImageForJointNameError]: Returns a 2D point for the joint name you specify, relative to the input image.
//
// # Getting the Parent Joint Name
//
//   - [VNHumanBodyPose3DObservation.ParentJointNameForJointName]: Returns the parent joint of the joint name you specify.
//
// # Getting the Body Height
//
//   - [VNHumanBodyPose3DObservation.HeightEstimation]: The technique the framework uses to estimate body height.
//   - [VNHumanBodyPose3DObservation.BodyHeight]: The estimated human body height, in meters.
//
// # Getting the Camera Position
//
//   - [VNHumanBodyPose3DObservation.CameraOriginMatrix]: A transform from the skeleton hip to the camera.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation
type VNHumanBodyPose3DObservation struct {
	VNRecognizedPoints3DObservation
}

// VNHumanBodyPose3DObservationFromID constructs a [VNHumanBodyPose3DObservation] from an objc.ID.
//
// An observation that provides the 3D body points the request recognizes.
func VNHumanBodyPose3DObservationFromID(id objc.ID) VNHumanBodyPose3DObservation {
	return VNHumanBodyPose3DObservation{VNRecognizedPoints3DObservation: VNRecognizedPoints3DObservationFromID(id)}
}

// NOTE: VNHumanBodyPose3DObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHumanBodyPose3DObservation] class.
//
// # Accessing Points
//
//   - [IVNHumanBodyPose3DObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [IVNHumanBodyPose3DObservation.AvailableJointsGroupNames]: The available joint group names in the observation.
//   - [IVNHumanBodyPose3DObservation.RecognizedPointForJointNameError]: Returns the point for a joint name that the observation recognizes.
//   - [IVNHumanBodyPose3DObservation.RecognizedPointsForJointsGroupNameError]: Returns a collection of points for the group name you specify.
//
// # Getting the Joint Position
//
//   - [IVNHumanBodyPose3DObservation.PointInImageForJointNameError]: Returns a 2D point for the joint name you specify, relative to the input image.
//
// # Getting the Parent Joint Name
//
//   - [IVNHumanBodyPose3DObservation.ParentJointNameForJointName]: Returns the parent joint of the joint name you specify.
//
// # Getting the Body Height
//
//   - [IVNHumanBodyPose3DObservation.HeightEstimation]: The technique the framework uses to estimate body height.
//   - [IVNHumanBodyPose3DObservation.BodyHeight]: The estimated human body height, in meters.
//
// # Getting the Camera Position
//
//   - [IVNHumanBodyPose3DObservation.CameraOriginMatrix]: A transform from the skeleton hip to the camera.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation
type IVNHumanBodyPose3DObservation interface {
	IVNRecognizedPoints3DObservation

	// Topic: Accessing Points

	// The names of the available joints in the observation.
	AvailableJointNames() []string
	// The available joint group names in the observation.
	AvailableJointsGroupNames() []string
	// Returns the point for a joint name that the observation recognizes.
	RecognizedPointForJointNameError(jointName VNHumanBodyPose3DObservationJointName) (IVNHumanBodyRecognizedPoint3D, error)
	// Returns a collection of points for the group name you specify.
	RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanBodyPose3DObservationJointsGroupName) (foundation.INSDictionary, error)

	// Topic: Getting the Joint Position

	// Returns a 2D point for the joint name you specify, relative to the input image.
	PointInImageForJointNameError(jointName VNHumanBodyPose3DObservationJointName) (IVNPoint, error)

	// Topic: Getting the Parent Joint Name

	// Returns the parent joint of the joint name you specify.
	ParentJointNameForJointName(jointName VNHumanBodyPose3DObservationJointName) VNHumanBodyPose3DObservationJointName

	// Topic: Getting the Body Height

	// The technique the framework uses to estimate body height.
	HeightEstimation() VNHumanBodyPose3DObservationHeightEstimation
	// The estimated human body height, in meters.
	BodyHeight() float32

	// Topic: Getting the Camera Position

	// A transform from the skeleton hip to the camera.
	CameraOriginMatrix() unsafe.Pointer

	// Gets a position relative to the camera for the body joint you specify.
	GetCameraRelativePositionForJointNameError(modelPositionOut unsafe.Pointer, jointName VNHumanBodyPose3DObservationJointName) (bool, error)
}

// Init initializes the instance.
func (h VNHumanBodyPose3DObservation) Init() VNHumanBodyPose3DObservation {
	rv := objc.Send[VNHumanBodyPose3DObservation](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHumanBodyPose3DObservation) Autorelease() VNHumanBodyPose3DObservation {
	rv := objc.Send[VNHumanBodyPose3DObservation](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHumanBodyPose3DObservation creates a new VNHumanBodyPose3DObservation instance.
func NewVNHumanBodyPose3DObservation() VNHumanBodyPose3DObservation {
	class := getVNHumanBodyPose3DObservationClass()
	rv := objc.Send[VNHumanBodyPose3DObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the point for a joint name that the observation recognizes.
//
// jointName: The joint name to retrieve.
//
// # Return Value
//
// The point for the joint name.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/recognizedPoint(_:)
func (h VNHumanBodyPose3DObservation) RecognizedPointForJointNameError(jointName VNHumanBodyPose3DObservationJointName) (IVNHumanBodyRecognizedPoint3D, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNHumanBodyRecognizedPoint3D{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNHumanBodyRecognizedPoint3DFromID(rv), nil

}

// Returns a collection of points for the group name you specify.
//
// jointsGroupName: The name of the human body joints group.
//
// # Return Value
//
// A collection of points.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/recognizedPoints(_:)
func (h VNHumanBodyPose3DObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanBodyPose3DObservationJointsGroupName) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Returns a 2D point for the joint name you specify, relative to the input
// image.
//
// jointName: The name of the human body joint.
//
// # Return Value
//
// A projection of the 3D position onto the original 2D image in normalized,
// lower left origin coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/pointInImage(_:)
func (h VNHumanBodyPose3DObservation) PointInImageForJointNameError(jointName VNHumanBodyPose3DObservationJointName) (IVNPoint, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("pointInImageForJointName:error:"), jointName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNPoint{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNPointFromID(rv), nil

}

// Returns the parent joint of the joint name you specify.
//
// jointName: The name of the body joint to return the parent of.
//
// # Return Value
//
// The name of the parent joint.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/parentJointName(_:)
func (h VNHumanBodyPose3DObservation) ParentJointNameForJointName(jointName VNHumanBodyPose3DObservationJointName) VNHumanBodyPose3DObservationJointName {
	rv := objc.Send[VNHumanBodyPose3DObservationJointName](h.ID, objc.Sel("parentJointNameForJointName:"), jointName)
	return VNHumanBodyPose3DObservationJointName(rv)
}

// Gets a position relative to the camera for the body joint you specify.
//
// modelPositionOut: The reference to the position.
//
// jointName: The name of the humany body joint.
//
// error: If an error occurs, an error object that describes the error; otherwise,
// `nil`.
//
// modelPositionOut is a [*simd.simd_float4x4].
//
// # Return Value
//
// A Boolean value that indicates the success of determining the position.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/getCameraRelativePosition:forJointName:error:
func (h VNHumanBodyPose3DObservation) GetCameraRelativePositionForJointNameError(modelPositionOut unsafe.Pointer, jointName VNHumanBodyPose3DObservationJointName) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](h.ID, objc.Sel("getCameraRelativePosition:forJointName:error:"), modelPositionOut, jointName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getCameraRelativePosition:forJointName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The names of the available joints in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/availableJointNames
func (h VNHumanBodyPose3DObservation) AvailableJointNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointNames"))
	return objc.ConvertSliceToStrings(rv)
}

// The available joint group names in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/availableJointsGroupNames
func (h VNHumanBodyPose3DObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointsGroupNames"))
	return objc.ConvertSliceToStrings(rv)
}

// The technique the framework uses to estimate body height.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/heightEstimation-swift.property
func (h VNHumanBodyPose3DObservation) HeightEstimation() VNHumanBodyPose3DObservationHeightEstimation {
	rv := objc.Send[VNHumanBodyPose3DObservationHeightEstimation](h.ID, objc.Sel("heightEstimation"))
	return VNHumanBodyPose3DObservationHeightEstimation(rv)
}

// The estimated human body height, in meters.
//
// # Discussion
//
// The framework returns an accurate height if [HeightEstimation] is
// [VNHumanBodyPose3DObservationHeightEstimationMeasured]; otherwise, it
// returns a [VNHumanBodyPose3DObservationHeightEstimationReference] height.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/bodyHeight
func (h VNHumanBodyPose3DObservation) BodyHeight() float32 {
	rv := objc.Send[float32](h.ID, objc.Sel("bodyHeight"))
	return rv
}

// A transform from the skeleton hip to the camera.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/cameraOriginMatrix
func (h VNHumanBodyPose3DObservation) CameraOriginMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h.ID, objc.Sel("cameraOriginMatrix"))
	return rv
}
