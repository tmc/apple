// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNClassification] class.
var (
	_SNClassificationClass     SNClassificationClass
	_SNClassificationClassOnce sync.Once
)

func getSNClassificationClass() SNClassificationClass {
	_SNClassificationClassOnce.Do(func() {
		_SNClassificationClass = SNClassificationClass{class: objc.GetClass("SNClassification")}
	})
	return _SNClassificationClass
}

// GetSNClassificationClass returns the class object for SNClassification.
func GetSNClassificationClass() SNClassificationClass {
	return getSNClassificationClass()
}

type SNClassificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNClassificationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNClassificationClass) Alloc() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A type that pairs a sound classifier’s prediction with its confidence in
// that prediction.
//
// # Overview
//
// An [SNClassification] represents a single sound classification prediction,
// and the sound classifier model’s confidence in that prediction.
//
// # Inspecting a Classification
//
//   - [SNClassification.Identifier]: A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
//   - [SNClassification.Confidence]: The confidence value the model has in its prediction.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassification
type SNClassification struct {
	objectivec.Object
}

// SNClassificationFromID constructs a [SNClassification] from an objc.ID.
//
// A type that pairs a sound classifier’s prediction with its confidence in
// that prediction.
func SNClassificationFromID(id objc.ID) SNClassification {
	return SNClassification{objectivec.Object{ID: id}}
}

// NOTE: SNClassification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SNClassification] class.
//
// # Inspecting a Classification
//
//   - [ISNClassification.Identifier]: A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
//   - [ISNClassification.Confidence]: The confidence value the model has in its prediction.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassification
type ISNClassification interface {
	objectivec.IObject

	// Topic: Inspecting a Classification

	// A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
	Identifier() string
	// The confidence value the model has in its prediction.
	Confidence() float64
}

// Init initializes the instance.
func (c SNClassification) Init() SNClassification {
	rv := objc.Send[SNClassification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SNClassification) Autorelease() SNClassification {
	rv := objc.Send[SNClassification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassification creates a new SNClassification instance.
func NewSNClassification() SNClassification {
	class := getSNClassificationClass()
	rv := objc.Send[SNClassification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A prediction label that’s one of the classifications a sound
// classifier’s underlying model defines.
//
// # Discussion
//
// An example `identifier` might be a string like `laughter` or `applause`.
// The sound classifier’s underlying model defines the possible string
// values, which are typically technical names that you don’t directly
// present in your app’s user interface.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/identifier
func (c SNClassification) Identifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The confidence value the model has in its prediction.
//
// # Discussion
//
// The model assigns confidence values in the range `[0, 1.0]`, where `1.0`
// represents 100% confidence.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/confidence
func (c SNClassification) Confidence() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("confidence"))
	return rv
}
