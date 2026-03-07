// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNAnimalBodyPoseObservation] class.
var (
	_VNAnimalBodyPoseObservationClass     VNAnimalBodyPoseObservationClass
	_VNAnimalBodyPoseObservationClassOnce sync.Once
)

func getVNAnimalBodyPoseObservationClass() VNAnimalBodyPoseObservationClass {
	_VNAnimalBodyPoseObservationClassOnce.Do(func() {
		_VNAnimalBodyPoseObservationClass = VNAnimalBodyPoseObservationClass{class: objc.GetClass("VNAnimalBodyPoseObservation")}
	})
	return _VNAnimalBodyPoseObservationClass
}

// GetVNAnimalBodyPoseObservationClass returns the class object for VNAnimalBodyPoseObservation.
func GetVNAnimalBodyPoseObservationClass() VNAnimalBodyPoseObservationClass {
	return getVNAnimalBodyPoseObservationClass()
}

type VNAnimalBodyPoseObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNAnimalBodyPoseObservationClass) Alloc() VNAnimalBodyPoseObservation {
	rv := objc.Send[VNAnimalBodyPoseObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An observation that provides the animal body points the analysis
// recognizes.
//
// # Accessing Points
//
//   - [VNAnimalBodyPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [VNAnimalBodyPoseObservation.AvailableJointGroupNames]: The available joint group names in the observation.
//   - [VNAnimalBodyPoseObservation.RecognizedPointForJointNameError]: Returns the point for a joint name the observation recognizes.
//   - [VNAnimalBodyPoseObservation.RecognizedPointsForJointsGroupNameError]: Returns the points for a joint group name the observation recognizes.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation
type VNAnimalBodyPoseObservation struct {
	VNRecognizedPointsObservation
}

// VNAnimalBodyPoseObservationFromID constructs a [VNAnimalBodyPoseObservation] from an objc.ID.
//
// An observation that provides the animal body points the analysis
// recognizes.
func VNAnimalBodyPoseObservationFromID(id objc.ID) VNAnimalBodyPoseObservation {
	return VNAnimalBodyPoseObservation{VNRecognizedPointsObservation: VNRecognizedPointsObservationFromID(id)}
}
// NOTE: VNAnimalBodyPoseObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNAnimalBodyPoseObservation] class.
//
// # Accessing Points
//
//   - [IVNAnimalBodyPoseObservation.AvailableJointNames]: The names of the available joints in the observation.
//   - [IVNAnimalBodyPoseObservation.AvailableJointGroupNames]: The available joint group names in the observation.
//   - [IVNAnimalBodyPoseObservation.RecognizedPointForJointNameError]: Returns the point for a joint name the observation recognizes.
//   - [IVNAnimalBodyPoseObservation.RecognizedPointsForJointsGroupNameError]: Returns the points for a joint group name the observation recognizes.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation
type IVNAnimalBodyPoseObservation interface {
	IVNRecognizedPointsObservation

	// Topic: Accessing Points

	// The names of the available joints in the observation.
	AvailableJointNames() []string
	// The available joint group names in the observation.
	AvailableJointGroupNames() []string
	// Returns the point for a joint name the observation recognizes.
	RecognizedPointForJointNameError(jointName VNAnimalBodyPoseObservationJointName) (IVNRecognizedPoint, error)
	// Returns the points for a joint group name the observation recognizes.
	RecognizedPointsForJointsGroupNameError(jointsGroupName VNAnimalBodyPoseObservationJointsGroupName) (foundation.INSDictionary, error)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (a VNAnimalBodyPoseObservation) Init() VNAnimalBodyPoseObservation {
	rv := objc.Send[VNAnimalBodyPoseObservation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a VNAnimalBodyPoseObservation) Autorelease() VNAnimalBodyPoseObservation {
	rv := objc.Send[VNAnimalBodyPoseObservation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNAnimalBodyPoseObservation creates a new VNAnimalBodyPoseObservation instance.
func NewVNAnimalBodyPoseObservation() VNAnimalBodyPoseObservation {
	class := getVNAnimalBodyPoseObservationClass()
	rv := objc.Send[VNAnimalBodyPoseObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the point for a joint name the observation recognizes.
//
// jointName: The joint name to retrieve.
//
// # Return Value
// 
// The point for the joint name.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoint(_:)
func (a VNAnimalBodyPoseObservation) RecognizedPointForJointNameError(jointName VNAnimalBodyPoseObservationJointName) (IVNRecognizedPoint, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNRecognizedPoint{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNRecognizedPointFromID(rv), nil

}

// Returns the points for a joint group name the observation recognizes.
//
// jointsGroupName: The joint group of the points to retrieve.
//
// # Return Value
// 
// The dictionary of points the observation associates with the group name.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoints(_:)
func (a VNAnimalBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName VNAnimalBodyPoseObservationJointsGroupName) (foundation.INSDictionary, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}
func (a VNAnimalBodyPoseObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The names of the available joints in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/availableJointNames
func (a VNAnimalBodyPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableJointNames"))
	return objc.ConvertSliceToStrings(rv)
}



// The available joint group names in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/availableJointGroupNames
func (a VNAnimalBodyPoseObservation) AvailableJointGroupNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableJointGroupNames"))
	return objc.ConvertSliceToStrings(rv)
}



























