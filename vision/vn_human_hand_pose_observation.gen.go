// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNHumanHandPoseObservation] class.
var (
	_VNHumanHandPoseObservationClass     VNHumanHandPoseObservationClass
	_VNHumanHandPoseObservationClassOnce sync.Once
)

func getVNHumanHandPoseObservationClass() VNHumanHandPoseObservationClass {
	_VNHumanHandPoseObservationClassOnce.Do(func() {
		_VNHumanHandPoseObservationClass = VNHumanHandPoseObservationClass{class: objc.GetClass("VNHumanHandPoseObservation")}
	})
	return _VNHumanHandPoseObservationClass
}

// GetVNHumanHandPoseObservationClass returns the class object for VNHumanHandPoseObservation.
func GetVNHumanHandPoseObservationClass() VNHumanHandPoseObservationClass {
	return getVNHumanHandPoseObservationClass()
}

type VNHumanHandPoseObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNHumanHandPoseObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHumanHandPoseObservationClass) Alloc() VNHumanHandPoseObservation {
	rv := objc.Send[VNHumanHandPoseObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that provides the hand points the analysis recognized.
//
// # Retrieving Points
//
//   - [VNHumanHandPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [VNHumanHandPoseObservation.AvailableJointsGroupNames]: The joint group names available in the observation.
//   - [VNHumanHandPoseObservation.RecognizedPointForJointNameError]: Retrieves the recognized point for a joint name.
//   - [VNHumanHandPoseObservation.RecognizedPointsForJointsGroupNameError]: Retrieves the recognized points associated with the joint group name.
//
// # Determining the Chirality
//
//   - [VNHumanHandPoseObservation.Chirality]: The chirality, or handedness, of a pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation
type VNHumanHandPoseObservation struct {
	VNRecognizedPointsObservation
}

// VNHumanHandPoseObservationFromID constructs a [VNHumanHandPoseObservation] from an objc.ID.
//
// An observation that provides the hand points the analysis recognized.
func VNHumanHandPoseObservationFromID(id objc.ID) VNHumanHandPoseObservation {
	return VNHumanHandPoseObservation{VNRecognizedPointsObservation: VNRecognizedPointsObservationFromID(id)}
}
// NOTE: VNHumanHandPoseObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHumanHandPoseObservation] class.
//
// # Retrieving Points
//
//   - [IVNHumanHandPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [IVNHumanHandPoseObservation.AvailableJointsGroupNames]: The joint group names available in the observation.
//   - [IVNHumanHandPoseObservation.RecognizedPointForJointNameError]: Retrieves the recognized point for a joint name.
//   - [IVNHumanHandPoseObservation.RecognizedPointsForJointsGroupNameError]: Retrieves the recognized points associated with the joint group name.
//
// # Determining the Chirality
//
//   - [IVNHumanHandPoseObservation.Chirality]: The chirality, or handedness, of a pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation
type IVNHumanHandPoseObservation interface {
	IVNRecognizedPointsObservation

	// Topic: Retrieving Points

	// The names of the available joints in the observation.
	AvailableJointNames() []string
	// The joint group names available in the observation.
	AvailableJointsGroupNames() []string
	// Retrieves the recognized point for a joint name.
	RecognizedPointForJointNameError(jointName VNHumanHandPoseObservationJointName) (IVNRecognizedPoint, error)
	// Retrieves the recognized points associated with the joint group name.
	RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanHandPoseObservationJointsGroupName) (foundation.INSDictionary, error)

	// Topic: Determining the Chirality

	// The chirality, or handedness, of a pose.
	Chirality() VNChirality
}

// Init initializes the instance.
func (h VNHumanHandPoseObservation) Init() VNHumanHandPoseObservation {
	rv := objc.Send[VNHumanHandPoseObservation](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHumanHandPoseObservation) Autorelease() VNHumanHandPoseObservation {
	rv := objc.Send[VNHumanHandPoseObservation](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHumanHandPoseObservation creates a new VNHumanHandPoseObservation instance.
func NewVNHumanHandPoseObservation() VNHumanHandPoseObservation {
	class := getVNHumanHandPoseObservationClass()
	rv := objc.Send[VNHumanHandPoseObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the recognized point for a joint name.
//
// jointName: The joint name of the point to retrieve.
//
// # Return Value
// 
// The point for the joint name.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/recognizedPoint(_:)
func (h VNHumanHandPoseObservation) RecognizedPointForJointNameError(jointName VNHumanHandPoseObservationJointName) (IVNRecognizedPoint, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNRecognizedPoint{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNRecognizedPointFromID(rv), nil

}
// Retrieves the recognized points associated with the joint group name.
//
// jointsGroupName: The joint group name of the points to retrieve.
//
// # Return Value
// 
// The array of points associated with the joint group name.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/recognizedPoints(_:)
func (h VNHumanHandPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanHandPoseObservationJointsGroupName) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// The names of the available joints in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/availableJointNames
func (h VNHumanHandPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointNames"))
	return objc.ConvertSliceToStrings(rv)
}
// The joint group names available in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/availableJointsGroupNames
func (h VNHumanHandPoseObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointsGroupNames"))
	return objc.ConvertSliceToStrings(rv)
}
// The chirality, or handedness, of a pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/chirality
func (h VNHumanHandPoseObservation) Chirality() VNChirality {
	rv := objc.Send[VNChirality](h.ID, objc.Sel("chirality"))
	return VNChirality(rv)
}

