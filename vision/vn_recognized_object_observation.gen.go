// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VNRecognizedObjectObservation] class.
var (
	_VNRecognizedObjectObservationClass     VNRecognizedObjectObservationClass
	_VNRecognizedObjectObservationClassOnce sync.Once
)

func getVNRecognizedObjectObservationClass() VNRecognizedObjectObservationClass {
	_VNRecognizedObjectObservationClassOnce.Do(func() {
		_VNRecognizedObjectObservationClass = VNRecognizedObjectObservationClass{class: objc.GetClass("VNRecognizedObjectObservation")}
	})
	return _VNRecognizedObjectObservationClass
}

// GetVNRecognizedObjectObservationClass returns the class object for VNRecognizedObjectObservation.
func GetVNRecognizedObjectObservationClass() VNRecognizedObjectObservationClass {
	return getVNRecognizedObjectObservationClass()
}

type VNRecognizedObjectObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRecognizedObjectObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedObjectObservationClass) Alloc() VNRecognizedObjectObservation {
	rv := objc.Send[VNRecognizedObjectObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A detected object observation with an array of classification labels that
// classify the recognized object.
//
// # Overview
// 
// The confidence of the classifications sum up to `1.0.` Multiply the
// classification confidence with the confidence of this observation.
//
// # Classifying a Recognized Object
//
//   - [VNRecognizedObjectObservation.Labels]: An array of observations that classify the recognized object.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation
type VNRecognizedObjectObservation struct {
	VNDetectedObjectObservation
}

// VNRecognizedObjectObservationFromID constructs a [VNRecognizedObjectObservation] from an objc.ID.
//
// A detected object observation with an array of classification labels that
// classify the recognized object.
func VNRecognizedObjectObservationFromID(id objc.ID) VNRecognizedObjectObservation {
	return VNRecognizedObjectObservation{VNDetectedObjectObservation: VNDetectedObjectObservationFromID(id)}
}
// NOTE: VNRecognizedObjectObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizedObjectObservation] class.
//
// # Classifying a Recognized Object
//
//   - [IVNRecognizedObjectObservation.Labels]: An array of observations that classify the recognized object.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation
type IVNRecognizedObjectObservation interface {
	IVNDetectedObjectObservation

	// Topic: Classifying a Recognized Object

	// An array of observations that classify the recognized object.
	Labels() []VNClassificationObservation
}

// Init initializes the instance.
func (r VNRecognizedObjectObservation) Init() VNRecognizedObjectObservation {
	rv := objc.Send[VNRecognizedObjectObservation](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedObjectObservation) Autorelease() VNRecognizedObjectObservation {
	rv := objc.Send[VNRecognizedObjectObservation](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedObjectObservation creates a new VNRecognizedObjectObservation instance.
func NewVNRecognizedObjectObservation() VNRecognizedObjectObservation {
	class := getVNRecognizedObjectObservationClass()
	rv := objc.Send[VNRecognizedObjectObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewRecognizedObjectObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNRecognizedObjectObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedObjectObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNRecognizedObjectObservationFromID(rv)
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
func NewRecognizedObjectObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNRecognizedObjectObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNRecognizedObjectObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNRecognizedObjectObservationFromID(rv)
}

// An array of observations that classify the recognized object.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation/labels
func (r VNRecognizedObjectObservation) Labels() []VNClassificationObservation {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("labels"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNClassificationObservation {
		return VNClassificationObservationFromID(id)
	})
}

