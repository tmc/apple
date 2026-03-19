// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNImageAestheticsScoresObservation] class.
var (
	_VNImageAestheticsScoresObservationClass     VNImageAestheticsScoresObservationClass
	_VNImageAestheticsScoresObservationClassOnce sync.Once
)

func getVNImageAestheticsScoresObservationClass() VNImageAestheticsScoresObservationClass {
	_VNImageAestheticsScoresObservationClassOnce.Do(func() {
		_VNImageAestheticsScoresObservationClass = VNImageAestheticsScoresObservationClass{class: objc.GetClass("VNImageAestheticsScoresObservation")}
	})
	return _VNImageAestheticsScoresObservationClass
}

// GetVNImageAestheticsScoresObservationClass returns the class object for VNImageAestheticsScoresObservation.
func GetVNImageAestheticsScoresObservationClass() VNImageAestheticsScoresObservationClass {
	return getVNImageAestheticsScoresObservationClass()
}

type VNImageAestheticsScoresObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNImageAestheticsScoresObservationClass) Alloc() VNImageAestheticsScoresObservation {
	rv := objc.Send[VNImageAestheticsScoresObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the overall score of aesthetic attributes for an
// image.
//
// # Parsing Observation Content
//
//   - [VNImageAestheticsScoresObservation.OverallScore]: A score which incorporates aesthetic score, failure score, and utility labels.
//   - [VNImageAestheticsScoresObservation.IsUtility]: A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation
type VNImageAestheticsScoresObservation struct {
	VNObservation
}

// VNImageAestheticsScoresObservationFromID constructs a [VNImageAestheticsScoresObservation] from an objc.ID.
//
// An object that represents the overall score of aesthetic attributes for an
// image.
func VNImageAestheticsScoresObservationFromID(id objc.ID) VNImageAestheticsScoresObservation {
	return VNImageAestheticsScoresObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNImageAestheticsScoresObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNImageAestheticsScoresObservation] class.
//
// # Parsing Observation Content
//
//   - [IVNImageAestheticsScoresObservation.OverallScore]: A score which incorporates aesthetic score, failure score, and utility labels.
//   - [IVNImageAestheticsScoresObservation.IsUtility]: A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation
type IVNImageAestheticsScoresObservation interface {
	IVNObservation

	// Topic: Parsing Observation Content

	// A score which incorporates aesthetic score, failure score, and utility labels.
	OverallScore() float32
	// A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.
	IsUtility() bool

	// The results of the aesthetics request.
	Results() IVNImageAestheticsScoresObservation
	SetResults(value IVNImageAestheticsScoresObservation)
}

// Init initializes the instance.
func (i VNImageAestheticsScoresObservation) Init() VNImageAestheticsScoresObservation {
	rv := objc.Send[VNImageAestheticsScoresObservation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i VNImageAestheticsScoresObservation) Autorelease() VNImageAestheticsScoresObservation {
	rv := objc.Send[VNImageAestheticsScoresObservation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNImageAestheticsScoresObservation creates a new VNImageAestheticsScoresObservation instance.
func NewVNImageAestheticsScoresObservation() VNImageAestheticsScoresObservation {
	class := getVNImageAestheticsScoresObservationClass()
	rv := objc.Send[VNImageAestheticsScoresObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A score which incorporates aesthetic score, failure score, and utility
// labels.
//
// # Discussion
// 
// This returns a value within the range of `-1` and `1`, where `-1` is least
// desirable and `1` is most desirable.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation/overallScore
func (i VNImageAestheticsScoresObservation) OverallScore() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("overallScore"))
	return rv
}
// A Boolean value that represents images that are not necessarily of poor
// image quality, but may not have memorable or exciting content.
//
// See: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation/isUtility
func (i VNImageAestheticsScoresObservation) IsUtility() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isUtility"))
	return rv
}
// The results of the aesthetics request.
//
// See: https://developer.apple.com/documentation/vision/vncalculateimageaestheticsscoresrequest/results
func (i VNImageAestheticsScoresObservation) Results() IVNImageAestheticsScoresObservation {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("results"))
	return VNImageAestheticsScoresObservationFromID(objc.ID(rv))
}
func (i VNImageAestheticsScoresObservation) SetResults(value IVNImageAestheticsScoresObservation) {
	objc.Send[struct{}](i.ID, objc.Sel("setResults:"), value)
}

