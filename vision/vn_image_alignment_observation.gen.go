// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNImageAlignmentObservation] class.
var (
	_VNImageAlignmentObservationClass     VNImageAlignmentObservationClass
	_VNImageAlignmentObservationClassOnce sync.Once
)

func getVNImageAlignmentObservationClass() VNImageAlignmentObservationClass {
	_VNImageAlignmentObservationClassOnce.Do(func() {
		_VNImageAlignmentObservationClass = VNImageAlignmentObservationClass{class: objc.GetClass("VNImageAlignmentObservation")}
	})
	return _VNImageAlignmentObservationClass
}

// GetVNImageAlignmentObservationClass returns the class object for VNImageAlignmentObservation.
func GetVNImageAlignmentObservationClass() VNImageAlignmentObservationClass {
	return getVNImageAlignmentObservationClass()
}

type VNImageAlignmentObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageAlignmentObservationClass) Alloc() VNImageAlignmentObservation {
	rv := objc.Send[VNImageAlignmentObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for image-analysis results that describe the
// relative alignment of two images.
//
// # Overview
// 
// This abstract superclass forms the basis of image alignment or registration
// output. You receive its subclasses, such as
// [VNImageTranslationAlignmentObservation] and
// [VNImageHomographicAlignmentObservation], by performing specific
// registration requests. Don’t create one of these classes yourself.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAlignmentObservation
type VNImageAlignmentObservation struct {
	VNObservation
}

// VNImageAlignmentObservationFromID constructs a [VNImageAlignmentObservation] from an objc.ID.
//
// The abstract superclass for image-analysis results that describe the
// relative alignment of two images.
func VNImageAlignmentObservationFromID(id objc.ID) VNImageAlignmentObservation {
	return VNImageAlignmentObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNImageAlignmentObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNImageAlignmentObservation] class.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAlignmentObservation
type IVNImageAlignmentObservation interface {
	IVNObservation
}

// Init initializes the instance.
func (i VNImageAlignmentObservation) Init() VNImageAlignmentObservation {
	rv := objc.Send[VNImageAlignmentObservation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageAlignmentObservation) Autorelease() VNImageAlignmentObservation {
	rv := objc.Send[VNImageAlignmentObservation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageAlignmentObservation creates a new VNImageAlignmentObservation instance.
func NewVNImageAlignmentObservation() VNImageAlignmentObservation {
	class := getVNImageAlignmentObservationClass()
	rv := objc.Send[VNImageAlignmentObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

