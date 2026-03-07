// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNRecognizedPoints3DObservation] class.
var (
	_VNRecognizedPoints3DObservationClass     VNRecognizedPoints3DObservationClass
	_VNRecognizedPoints3DObservationClassOnce sync.Once
)

func getVNRecognizedPoints3DObservationClass() VNRecognizedPoints3DObservationClass {
	_VNRecognizedPoints3DObservationClassOnce.Do(func() {
		_VNRecognizedPoints3DObservationClass = VNRecognizedPoints3DObservationClass{class: objc.GetClass("VNRecognizedPoints3DObservation")}
	})
	return _VNRecognizedPoints3DObservationClass
}

// GetVNRecognizedPoints3DObservationClass returns the class object for VNRecognizedPoints3DObservation.
func GetVNRecognizedPoints3DObservationClass() VNRecognizedPoints3DObservationClass {
	return getVNRecognizedPoints3DObservationClass()
}

type VNRecognizedPoints3DObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedPoints3DObservationClass) Alloc() VNRecognizedPoints3DObservation {
	rv := objc.Send[VNRecognizedPoints3DObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An observation that provides the 3D points for a request.
//
// # Inspecting the Observation
//
//   - [VNRecognizedPoints3DObservation.AvailableKeys]: The available point keys in the observation.
//   - [VNRecognizedPoints3DObservation.AvailableGroupKeys]: The available point group keys in the observation.
//   - [VNRecognizedPoints3DObservation.RecognizedPointForKeyError]: Returns a point for a key you specify.
//   - [VNRecognizedPoints3DObservation.RecognizedPointsForGroupKeyError]: Returns a point for a group key you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation
type VNRecognizedPoints3DObservation struct {
	VNObservation
}

// VNRecognizedPoints3DObservationFromID constructs a [VNRecognizedPoints3DObservation] from an objc.ID.
//
// An observation that provides the 3D points for a request.
func VNRecognizedPoints3DObservationFromID(id objc.ID) VNRecognizedPoints3DObservation {
	return VNRecognizedPoints3DObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNRecognizedPoints3DObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNRecognizedPoints3DObservation] class.
//
// # Inspecting the Observation
//
//   - [IVNRecognizedPoints3DObservation.AvailableKeys]: The available point keys in the observation.
//   - [IVNRecognizedPoints3DObservation.AvailableGroupKeys]: The available point group keys in the observation.
//   - [IVNRecognizedPoints3DObservation.RecognizedPointForKeyError]: Returns a point for a key you specify.
//   - [IVNRecognizedPoints3DObservation.RecognizedPointsForGroupKeyError]: Returns a point for a group key you specify.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation
type IVNRecognizedPoints3DObservation interface {
	IVNObservation

	// Topic: Inspecting the Observation

	// The available point keys in the observation.
	AvailableKeys() []string
	// The available point group keys in the observation.
	AvailableGroupKeys() []string
	// Returns a point for a key you specify.
	RecognizedPointForKeyError(pointKey VNRecognizedPointKey) (IVNRecognizedPoint3D, error)
	// Returns a point for a group key you specify.
	RecognizedPointsForGroupKeyError(groupKey VNRecognizedPointGroupKey) (foundation.INSDictionary, error)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (r VNRecognizedPoints3DObservation) Init() VNRecognizedPoints3DObservation {
	rv := objc.Send[VNRecognizedPoints3DObservation](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedPoints3DObservation) Autorelease() VNRecognizedPoints3DObservation {
	rv := objc.Send[VNRecognizedPoints3DObservation](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedPoints3DObservation creates a new VNRecognizedPoints3DObservation instance.
func NewVNRecognizedPoints3DObservation() VNRecognizedPoints3DObservation {
	class := getVNRecognizedPoints3DObservationClass()
	rv := objc.Send[VNRecognizedPoints3DObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a point for a key you specify.
//
// pointKey: The key of the point to retrieve.
//
// # Return Value
// 
// The point the observation associates with the key.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/recognizedPoint(forKey:)
func (r VNRecognizedPoints3DObservation) RecognizedPointForKeyError(pointKey VNRecognizedPointKey) (IVNRecognizedPoint3D, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recognizedPointForKey:error:"), objc.String(string(pointKey)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNRecognizedPoint3D{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNRecognizedPoint3DFromID(rv), nil

}

// Returns a point for a group key you specify.
//
// groupKey: The group key to retrieve points for.
//
// # Return Value
// 
// A dictionary of labeled points for the group.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/recognizedPoints(forGroupKey:)
func (r VNRecognizedPoints3DObservation) RecognizedPointsForGroupKeyError(groupKey VNRecognizedPointGroupKey) (foundation.INSDictionary, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recognizedPointsForGroupKey:error:"), objc.String(string(groupKey)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}
func (r VNRecognizedPoints3DObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The available point keys in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/availableKeys
func (r VNRecognizedPoints3DObservation) AvailableKeys() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableKeys"))
	return objc.ConvertSliceToStrings(rv)
}



// The available point group keys in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/availableGroupKeys
func (r VNRecognizedPoints3DObservation) AvailableGroupKeys() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableGroupKeys"))
	return objc.ConvertSliceToStrings(rv)
}



























