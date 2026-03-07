// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNHumanBodyPoseObservation] class.
var (
	_VNHumanBodyPoseObservationClass     VNHumanBodyPoseObservationClass
	_VNHumanBodyPoseObservationClassOnce sync.Once
)

func getVNHumanBodyPoseObservationClass() VNHumanBodyPoseObservationClass {
	_VNHumanBodyPoseObservationClassOnce.Do(func() {
		_VNHumanBodyPoseObservationClass = VNHumanBodyPoseObservationClass{class: objc.GetClass("VNHumanBodyPoseObservation")}
	})
	return _VNHumanBodyPoseObservationClass
}

// GetVNHumanBodyPoseObservationClass returns the class object for VNHumanBodyPoseObservation.
func GetVNHumanBodyPoseObservationClass() VNHumanBodyPoseObservationClass {
	return getVNHumanBodyPoseObservationClass()
}

type VNHumanBodyPoseObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHumanBodyPoseObservationClass) Alloc() VNHumanBodyPoseObservation {
	rv := objc.Send[VNHumanBodyPoseObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An observation that provides the body points the analysis recognized.
//
// # Accessing Points
//
//   - [VNHumanBodyPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [VNHumanBodyPoseObservation.AvailableJointsGroupNames]: The available joint group names in the observation.
//   - [VNHumanBodyPoseObservation.RecognizedPointForJointNameError]: Retrieves the recognized point for a joint name.
//   - [VNHumanBodyPoseObservation.RecognizedPointsForJointsGroupNameError]: Retrieves the recognized points associated with the joint group name.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation
type VNHumanBodyPoseObservation struct {
	VNRecognizedPointsObservation
}

// VNHumanBodyPoseObservationFromID constructs a [VNHumanBodyPoseObservation] from an objc.ID.
//
// An observation that provides the body points the analysis recognized.
func VNHumanBodyPoseObservationFromID(id objc.ID) VNHumanBodyPoseObservation {
	return VNHumanBodyPoseObservation{VNRecognizedPointsObservation: VNRecognizedPointsObservationFromID(id)}
}
// NOTE: VNHumanBodyPoseObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNHumanBodyPoseObservation] class.
//
// # Accessing Points
//
//   - [IVNHumanBodyPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [IVNHumanBodyPoseObservation.AvailableJointsGroupNames]: The available joint group names in the observation.
//   - [IVNHumanBodyPoseObservation.RecognizedPointForJointNameError]: Retrieves the recognized point for a joint name.
//   - [IVNHumanBodyPoseObservation.RecognizedPointsForJointsGroupNameError]: Retrieves the recognized points associated with the joint group name.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation
type IVNHumanBodyPoseObservation interface {
	IVNRecognizedPointsObservation

	// Topic: Accessing Points

	// The names of the available joints in the observation.
	AvailableJointNames() []string
	// The available joint group names in the observation.
	AvailableJointsGroupNames() []string
	// Retrieves the recognized point for a joint name.
	RecognizedPointForJointNameError(jointName VNHumanBodyPoseObservationJointName) (IVNRecognizedPoint, error)
	// Retrieves the recognized points associated with the joint group name.
	RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanBodyPoseObservationJointsGroupName) (foundation.INSDictionary, error)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (h VNHumanBodyPoseObservation) Init() VNHumanBodyPoseObservation {
	rv := objc.Send[VNHumanBodyPoseObservation](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHumanBodyPoseObservation) Autorelease() VNHumanBodyPoseObservation {
	rv := objc.Send[VNHumanBodyPoseObservation](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHumanBodyPoseObservation creates a new VNHumanBodyPoseObservation instance.
func NewVNHumanBodyPoseObservation() VNHumanBodyPoseObservation {
	class := getVNHumanBodyPoseObservationClass()
	rv := objc.Send[VNHumanBodyPoseObservation](objc.ID(class.class), objc.Sel("new"))
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
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoint(_:)
func (h VNHumanBodyPoseObservation) RecognizedPointForJointNameError(jointName VNHumanBodyPoseObservationJointName) (IVNRecognizedPoint, error) {
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
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoints(_:)
func (h VNHumanBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName VNHumanBodyPoseObservationJointsGroupName) (foundation.INSDictionary, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](h.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}
func (h VNHumanBodyPoseObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](h.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The names of the available joints in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointNames
func (h VNHumanBodyPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointNames"))
	return objc.ConvertSliceToStrings(rv)
}



// The available joint group names in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointsGroupNames
func (h VNHumanBodyPoseObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("availableJointsGroupNames"))
	return objc.ConvertSliceToStrings(rv)
}



























