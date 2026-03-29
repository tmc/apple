// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNRecognizedPointsObservation] class.
var (
	_VNRecognizedPointsObservationClass     VNRecognizedPointsObservationClass
	_VNRecognizedPointsObservationClassOnce sync.Once
)

func getVNRecognizedPointsObservationClass() VNRecognizedPointsObservationClass {
	_VNRecognizedPointsObservationClassOnce.Do(func() {
		_VNRecognizedPointsObservationClass = VNRecognizedPointsObservationClass{class: objc.GetClass("VNRecognizedPointsObservation")}
	})
	return _VNRecognizedPointsObservationClass
}

// GetVNRecognizedPointsObservationClass returns the class object for VNRecognizedPointsObservation.
func GetVNRecognizedPointsObservationClass() VNRecognizedPointsObservationClass {
	return getVNRecognizedPointsObservationClass()
}

type VNRecognizedPointsObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRecognizedPointsObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedPointsObservationClass) Alloc() VNRecognizedPointsObservation {
	rv := objc.Send[VNRecognizedPointsObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that provides the points the analysis recognized.
//
// # Inspecting the Observation
//
//   - [VNRecognizedPointsObservation.AvailableKeys]: The available point keys in the observation.
//   - [VNRecognizedPointsObservation.AvailableGroupKeys]: The available point group keys in the observation.
//   - [VNRecognizedPointsObservation.RecognizedPointForKeyError]: Retrieves a recognized point for a key.
//   - [VNRecognizedPointsObservation.RecognizedPointsForGroupKeyError]: Retrieves the recognized points for a key.
//
// # Converting Points for Core ML
//
//   - [VNRecognizedPointsObservation.KeypointsMultiArrayAndReturnError]: Retrieves the grouping of normalized point coordinates and confidence scores in a format compatible with Core ML.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation
type VNRecognizedPointsObservation struct {
	VNObservation
}

// VNRecognizedPointsObservationFromID constructs a [VNRecognizedPointsObservation] from an objc.ID.
//
// An observation that provides the points the analysis recognized.
func VNRecognizedPointsObservationFromID(id objc.ID) VNRecognizedPointsObservation {
	return VNRecognizedPointsObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNRecognizedPointsObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizedPointsObservation] class.
//
// # Inspecting the Observation
//
//   - [IVNRecognizedPointsObservation.AvailableKeys]: The available point keys in the observation.
//   - [IVNRecognizedPointsObservation.AvailableGroupKeys]: The available point group keys in the observation.
//   - [IVNRecognizedPointsObservation.RecognizedPointForKeyError]: Retrieves a recognized point for a key.
//   - [IVNRecognizedPointsObservation.RecognizedPointsForGroupKeyError]: Retrieves the recognized points for a key.
//
// # Converting Points for Core ML
//
//   - [IVNRecognizedPointsObservation.KeypointsMultiArrayAndReturnError]: Retrieves the grouping of normalized point coordinates and confidence scores in a format compatible with Core ML.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation
type IVNRecognizedPointsObservation interface {
	IVNObservation

	// Topic: Inspecting the Observation

	// The available point keys in the observation.
	AvailableKeys() []string
	// The available point group keys in the observation.
	AvailableGroupKeys() []string
	// Retrieves a recognized point for a key.
	RecognizedPointForKeyError(pointKey VNRecognizedPointKey) (IVNRecognizedPoint, error)
	// Retrieves the recognized points for a key.
	RecognizedPointsForGroupKeyError(groupKey VNRecognizedPointGroupKey) (foundation.INSDictionary, error)

	// Topic: Converting Points for Core ML

	// Retrieves the grouping of normalized point coordinates and confidence scores in a format compatible with Core ML.
	KeypointsMultiArrayAndReturnError() (coreml.MLMultiArray, error)
}

// Init initializes the instance.
func (r VNRecognizedPointsObservation) Init() VNRecognizedPointsObservation {
	rv := objc.Send[VNRecognizedPointsObservation](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedPointsObservation) Autorelease() VNRecognizedPointsObservation {
	rv := objc.Send[VNRecognizedPointsObservation](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedPointsObservation creates a new VNRecognizedPointsObservation instance.
func NewVNRecognizedPointsObservation() VNRecognizedPointsObservation {
	class := getVNRecognizedPointsObservationClass()
	rv := objc.Send[VNRecognizedPointsObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves a recognized point for a key.
//
// pointKey: The key of the point to retrieve.
//
// # Return Value
// 
// The recognized point associated with the key.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/recognizedPoint(forKey:)
func (r VNRecognizedPointsObservation) RecognizedPointForKeyError(pointKey VNRecognizedPointKey) (IVNRecognizedPoint, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recognizedPointForKey:error:"), objc.String(string(pointKey)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNRecognizedPoint{}, foundation.NSErrorFrom(errorPtr)
	}
	return VNRecognizedPointFromID(rv), nil

}
// Retrieves the recognized points for a key.
//
// groupKey: The group key to retrieve recognized points for.
//
// # Return Value
// 
// A dictionary of labeled points for the group.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/recognizedPoints(forGroupKey:)
func (r VNRecognizedPointsObservation) RecognizedPointsForGroupKeyError(groupKey VNRecognizedPointGroupKey) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recognizedPointsForGroupKey:error:"), objc.String(string(groupKey)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}
// Retrieves the grouping of normalized point coordinates and confidence
// scores in a format compatible with Core ML.
//
// # Return Value
// 
// The key points converted to an [MLMultiArray].
//
// [MLMultiArray]: https://developer.apple.com/documentation/CoreML/MLMultiArray
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/keypointsMultiArray()
func (r VNRecognizedPointsObservation) KeypointsMultiArrayAndReturnError() (coreml.MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("keypointsMultiArrayAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return coreml.MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return coreml.MLMultiArrayFromID(rv), nil

}

// The available point keys in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/availableKeys
func (r VNRecognizedPointsObservation) AvailableKeys() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableKeys"))
	return objc.ConvertSliceToStrings(rv)
}
// The available point group keys in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/availableGroupKeys
func (r VNRecognizedPointsObservation) AvailableGroupKeys() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableGroupKeys"))
	return objc.ConvertSliceToStrings(rv)
}

