// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNHumanObservation] class.
var (
	_VNHumanObservationClass     VNHumanObservationClass
	_VNHumanObservationClassOnce sync.Once
)

func getVNHumanObservationClass() VNHumanObservationClass {
	_VNHumanObservationClassOnce.Do(func() {
		_VNHumanObservationClass = VNHumanObservationClass{class: objc.GetClass("VNHumanObservation")}
	})
	return _VNHumanObservationClass
}

// GetVNHumanObservationClass returns the class object for VNHumanObservation.
func GetVNHumanObservationClass() VNHumanObservationClass {
	return getVNHumanObservationClass()
}

type VNHumanObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNHumanObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHumanObservationClass) Alloc() VNHumanObservation {
	rv := objc.Send[VNHumanObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a person that the request detects.
//
// # Inspecting the Observation
//
//   - [VNHumanObservation.UpperBodyOnly]: A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanObservation
type VNHumanObservation struct {
	VNDetectedObjectObservation
}

// VNHumanObservationFromID constructs a [VNHumanObservation] from an objc.ID.
//
// An object that represents a person that the request detects.
func VNHumanObservationFromID(id objc.ID) VNHumanObservation {
	return VNHumanObservation{VNDetectedObjectObservation: VNDetectedObjectObservationFromID(id)}
}

// NOTE: VNHumanObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHumanObservation] class.
//
// # Inspecting the Observation
//
//   - [IVNHumanObservation.UpperBodyOnly]: A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanObservation
type IVNHumanObservation interface {
	IVNDetectedObjectObservation

	// Topic: Inspecting the Observation

	// A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle.
	UpperBodyOnly() bool
}

// Init initializes the instance.
func (h VNHumanObservation) Init() VNHumanObservation {
	rv := objc.Send[VNHumanObservation](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHumanObservation) Autorelease() VNHumanObservation {
	rv := objc.Send[VNHumanObservation](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHumanObservation creates a new VNHumanObservation instance.
func NewVNHumanObservation() VNHumanObservation {
	class := getVNHumanObservationClass()
	rv := objc.Send[VNHumanObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewHumanObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNHumanObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNHumanObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNHumanObservationFromID(rv)
}

// Creates an observation with a revision number and bounding box.
//
// requestRevision: The revision of the request to use.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func NewHumanObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNHumanObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNHumanObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNHumanObservationFromID(rv)
}

// A Boolean value that indicates whether the observation represents an
// upper-body or full-body rectangle.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanObservation/upperBodyOnly
func (h VNHumanObservation) UpperBodyOnly() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("upperBodyOnly"))
	return rv
}
