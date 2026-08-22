// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFAcousticFeature] class.
var (
	_SFAcousticFeatureClass     SFAcousticFeatureClass
	_SFAcousticFeatureClassOnce sync.Once
)

func getSFAcousticFeatureClass() SFAcousticFeatureClass {
	_SFAcousticFeatureClassOnce.Do(func() {
		_SFAcousticFeatureClass = SFAcousticFeatureClass{class: objc.GetClass("SFAcousticFeature")}
	})
	return _SFAcousticFeatureClass
}

// GetSFAcousticFeatureClass returns the class object for SFAcousticFeature.
func GetSFAcousticFeatureClass() SFAcousticFeatureClass {
	return getSFAcousticFeatureClass()
}

type SFAcousticFeatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFAcousticFeatureClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFAcousticFeatureClass) Alloc() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The value of a voice analysis metric.
//
// # Inspecting a feature
//
//   - [SFAcousticFeature.FrameDuration]: The duration of the audio frame.
//   - [SFAcousticFeature.AcousticFeatureValuePerFrame]: An array of feature values, one value per audio frame, corresponding to a transcript segment of recorded audio.
//   - [SFAcousticFeature.SetAcousticFeatureValuePerFrame]
//
// See: https://developer.apple.com/documentation/Speech/SFAcousticFeature
type SFAcousticFeature struct {
	objectivec.Object
}

// SFAcousticFeatureFromID constructs a [SFAcousticFeature] from an objc.ID.
//
// The value of a voice analysis metric.
func SFAcousticFeatureFromID(id objc.ID) SFAcousticFeature {
	return SFAcousticFeature{objectivec.Object{ID: id}}
}

// NOTE: SFAcousticFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFAcousticFeature] class.
//
// # Inspecting a feature
//
//   - [ISFAcousticFeature.FrameDuration]: The duration of the audio frame.
//   - [ISFAcousticFeature.AcousticFeatureValuePerFrame]: An array of feature values, one value per audio frame, corresponding to a transcript segment of recorded audio.
//   - [ISFAcousticFeature.SetAcousticFeatureValuePerFrame]
//
// See: https://developer.apple.com/documentation/Speech/SFAcousticFeature
type ISFAcousticFeature interface {
	objectivec.IObject

	// Topic: Inspecting a feature

	// The duration of the audio frame.
	FrameDuration() foundation.NSTimeInterval
	// An array of feature values, one value per audio frame, corresponding to a transcript segment of recorded audio.
	AcousticFeatureValuePerFrame() []float64
	SetAcousticFeatureValuePerFrame(value []float64)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a SFAcousticFeature) Init() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SFAcousticFeature) Autorelease() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAcousticFeature creates a new SFAcousticFeature instance.
func NewSFAcousticFeature() SFAcousticFeature {
	class := getSFAcousticFeatureClass()
	rv := objc.Send[SFAcousticFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a SFAcousticFeature) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The duration of the audio frame.
//
// See: https://developer.apple.com/documentation/Speech/SFAcousticFeature/frameDuration
func (a SFAcousticFeature) FrameDuration() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("frameDuration"))
	return foundation.NSTimeInterval(rv)
}

// An array of feature values, one value per audio frame, corresponding to a
// transcript segment of recorded audio.
//
// See: https://developer.apple.com/documentation/speech/sfacousticfeature/acousticfeaturevalueperframe-5krkk
func (a SFAcousticFeature) AcousticFeatureValuePerFrame() []float64 {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("acousticFeatureValuePerFrame"))
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(objc.Send[float64](id, objc.Sel("doubleValue")))
	})
}
func (a SFAcousticFeature) SetAcousticFeatureValuePerFrame(value []float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setAcousticFeatureValuePerFrame:"), objectivec.NumberSliceToNSArray(value))
}
