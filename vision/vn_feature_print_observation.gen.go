// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNFeaturePrintObservation] class.
var (
	_VNFeaturePrintObservationClass     VNFeaturePrintObservationClass
	_VNFeaturePrintObservationClassOnce sync.Once
)

func getVNFeaturePrintObservationClass() VNFeaturePrintObservationClass {
	_VNFeaturePrintObservationClassOnce.Do(func() {
		_VNFeaturePrintObservationClass = VNFeaturePrintObservationClass{class: objc.GetClass("VNFeaturePrintObservation")}
	})
	return _VNFeaturePrintObservationClass
}

// GetVNFeaturePrintObservationClass returns the class object for VNFeaturePrintObservation.
func GetVNFeaturePrintObservationClass() VNFeaturePrintObservationClass {
	return getVNFeaturePrintObservationClass()
}

type VNFeaturePrintObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFeaturePrintObservationClass) Alloc() VNFeaturePrintObservation {
	rv := objc.Send[VNFeaturePrintObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that provides the recognized feature print.
//
// # Fetching Feature Print Data
//
//   - [VNFeaturePrintObservation.Data]: The feature print data.
//   - [VNFeaturePrintObservation.ElementCount]: The total number of elements in the data.
//
// # Determining Types of Feature Prints
//
//   - [VNFeaturePrintObservation.ElementType]: The type of each element in the data.
//
// # Computing Distance Between Features
//
//   - [VNFeaturePrintObservation.ComputeDistanceToFeaturePrintObservationError]: Computes the distance between two feature print observations.
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation
type VNFeaturePrintObservation struct {
	VNObservation
}

// VNFeaturePrintObservationFromID constructs a [VNFeaturePrintObservation] from an objc.ID.
//
// An observation that provides the recognized feature print.
func VNFeaturePrintObservationFromID(id objc.ID) VNFeaturePrintObservation {
	return VNFeaturePrintObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNFeaturePrintObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNFeaturePrintObservation] class.
//
// # Fetching Feature Print Data
//
//   - [IVNFeaturePrintObservation.Data]: The feature print data.
//   - [IVNFeaturePrintObservation.ElementCount]: The total number of elements in the data.
//
// # Determining Types of Feature Prints
//
//   - [IVNFeaturePrintObservation.ElementType]: The type of each element in the data.
//
// # Computing Distance Between Features
//
//   - [IVNFeaturePrintObservation.ComputeDistanceToFeaturePrintObservationError]: Computes the distance between two feature print observations.
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation
type IVNFeaturePrintObservation interface {
	IVNObservation

	// Topic: Fetching Feature Print Data

	// The feature print data.
	Data() foundation.INSData
	// The total number of elements in the data.
	ElementCount() uint

	// Topic: Determining Types of Feature Prints

	// The type of each element in the data.
	ElementType() VNElementType

	// Topic: Computing Distance Between Features

	// Computes the distance between two feature print observations.
	ComputeDistanceToFeaturePrintObservationError(featurePrint VNFeaturePrintObservation) (float32, error)
}

// Init initializes the instance.
func (f VNFeaturePrintObservation) Init() VNFeaturePrintObservation {
	rv := objc.Send[VNFeaturePrintObservation](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFeaturePrintObservation) Autorelease() VNFeaturePrintObservation {
	rv := objc.Send[VNFeaturePrintObservation](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFeaturePrintObservation creates a new VNFeaturePrintObservation instance.
func NewVNFeaturePrintObservation() VNFeaturePrintObservation {
	class := getVNFeaturePrintObservationClass()
	rv := objc.Send[VNFeaturePrintObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Computes the distance between two feature print observations.
//
// outDistance: A pointer to store the calculated distance value.
//
// featurePrint: The feature print object whose distance to calculate.
//
// # Discussion
// 
// Shorter distances indicate greater similarity between feature prints.
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/computeDistance(_:to:)
func (f VNFeaturePrintObservation) ComputeDistanceToFeaturePrintObservationError(featurePrint VNFeaturePrintObservation) (float32, error) {
	var outDistance float32
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("computeDistance:toFeaturePrintObservation:error:"), unsafe.Pointer(&outDistance), featurePrint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("computeDistance:toFeaturePrintObservation:error: returned NO with nil NSError")
	}
	return outDistance, nil
}

// The feature print data.
//
// # Discussion
// 
// The data is divided into separate elements. Determine the type of element
// using [ElementType], and the number of elements using [ElementCount].
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/data
func (f VNFeaturePrintObservation) Data() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The total number of elements in the data.
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/elementCount
func (f VNFeaturePrintObservation) ElementCount() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("elementCount"))
	return rv
}

// The type of each element in the data.
//
// See: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/elementType
func (f VNFeaturePrintObservation) ElementType() VNElementType {
	rv := objc.Send[VNElementType](f.ID, objc.Sel("elementType"))
	return VNElementType(rv)
}

