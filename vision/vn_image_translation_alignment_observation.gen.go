// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNImageTranslationAlignmentObservation] class.
var (
	_VNImageTranslationAlignmentObservationClass     VNImageTranslationAlignmentObservationClass
	_VNImageTranslationAlignmentObservationClassOnce sync.Once
)

func getVNImageTranslationAlignmentObservationClass() VNImageTranslationAlignmentObservationClass {
	_VNImageTranslationAlignmentObservationClassOnce.Do(func() {
		_VNImageTranslationAlignmentObservationClass = VNImageTranslationAlignmentObservationClass{class: objc.GetClass("VNImageTranslationAlignmentObservation")}
	})
	return _VNImageTranslationAlignmentObservationClass
}

// GetVNImageTranslationAlignmentObservationClass returns the class object for VNImageTranslationAlignmentObservation.
func GetVNImageTranslationAlignmentObservationClass() VNImageTranslationAlignmentObservationClass {
	return getVNImageTranslationAlignmentObservationClass()
}

type VNImageTranslationAlignmentObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNImageTranslationAlignmentObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageTranslationAlignmentObservationClass) Alloc() VNImageTranslationAlignmentObservation {
	rv := objc.Send[VNImageTranslationAlignmentObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Affine transform information that an image-alignment request produces.
//
// # Overview
//
// This type of observation results from a
// [VNTranslationalImageRegistrationRequest], informing the
// [VNImageTranslationAlignmentObservation.AlignmentTransform] performed to
// align the input images.
//
// # Determining Alignment
//
//   - [VNImageTranslationAlignmentObservation.AlignmentTransform]: The alignment transform to align the floating image with the reference image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation
type VNImageTranslationAlignmentObservation struct {
	VNImageAlignmentObservation
}

// VNImageTranslationAlignmentObservationFromID constructs a [VNImageTranslationAlignmentObservation] from an objc.ID.
//
// Affine transform information that an image-alignment request produces.
func VNImageTranslationAlignmentObservationFromID(id objc.ID) VNImageTranslationAlignmentObservation {
	return VNImageTranslationAlignmentObservation{VNImageAlignmentObservation: VNImageAlignmentObservationFromID(id)}
}

// NOTE: VNImageTranslationAlignmentObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNImageTranslationAlignmentObservation] class.
//
// # Determining Alignment
//
//   - [IVNImageTranslationAlignmentObservation.AlignmentTransform]: The alignment transform to align the floating image with the reference image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation
type IVNImageTranslationAlignmentObservation interface {
	IVNImageAlignmentObservation

	// Topic: Determining Alignment

	// The alignment transform to align the floating image with the reference image.
	AlignmentTransform() corefoundation.CGAffineTransform
}

// Init initializes the instance.
func (i VNImageTranslationAlignmentObservation) Init() VNImageTranslationAlignmentObservation {
	rv := objc.Send[VNImageTranslationAlignmentObservation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageTranslationAlignmentObservation) Autorelease() VNImageTranslationAlignmentObservation {
	rv := objc.Send[VNImageTranslationAlignmentObservation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageTranslationAlignmentObservation creates a new VNImageTranslationAlignmentObservation instance.
func NewVNImageTranslationAlignmentObservation() VNImageTranslationAlignmentObservation {
	class := getVNImageTranslationAlignmentObservationClass()
	rv := objc.Send[VNImageTranslationAlignmentObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Vision/VNObservation/init(coder:)
func NewImageTranslationAlignmentObservationWithCoder(coder foundation.INSCoder) VNImageTranslationAlignmentObservation {
	instance := getVNImageTranslationAlignmentObservationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VNImageTranslationAlignmentObservationFromID(rv)
}

// The alignment transform to align the floating image with the reference
// image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation/alignmentTransform
func (i VNImageTranslationAlignmentObservation) AlignmentTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i.ID, objc.Sel("alignmentTransform"))
	return corefoundation.CGAffineTransform(rv)
}
