// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFVoiceAnalytics] class.
var (
	_SFVoiceAnalyticsClass     SFVoiceAnalyticsClass
	_SFVoiceAnalyticsClassOnce sync.Once
)

func getSFVoiceAnalyticsClass() SFVoiceAnalyticsClass {
	_SFVoiceAnalyticsClassOnce.Do(func() {
		_SFVoiceAnalyticsClass = SFVoiceAnalyticsClass{class: objc.GetClass("SFVoiceAnalytics")}
	})
	return _SFVoiceAnalyticsClass
}

// GetSFVoiceAnalyticsClass returns the class object for SFVoiceAnalytics.
func GetSFVoiceAnalyticsClass() SFVoiceAnalyticsClass {
	return getSFVoiceAnalyticsClass()
}

type SFVoiceAnalyticsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFVoiceAnalyticsClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFVoiceAnalyticsClass) Alloc() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A collection of vocal analysis metrics.
//
// # Overview
//
// Use an [SFAcousticFeature] object to access the [SFVoiceAnalytics]
// insights. Voice analytics include the following features:
//
// - Use [SFVoiceAnalytics.Jitter] to measure how pitch varies in audio. - Use
// [SFVoiceAnalytics.Shimmer] to measure how amplitude varies in audio. - Use
// [SFVoiceAnalytics.Pitch] to measure the highness and lowness of the tone. -
// Use [SFVoiceAnalytics.Voicing] to identify voiced regions in speech.
//
// These results are part of the [SFTranscriptionSegment] object and are
// available when the system sends the [SFSpeechRecognitionResult.Final] flag.
//
// # Analyzing voice
//
//   - [SFVoiceAnalytics.Voicing]: The likelihood of a voice in each frame of a transcription segment.
//   - [SFVoiceAnalytics.Pitch]: The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.
//   - [SFVoiceAnalytics.Jitter]: The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.
//   - [SFVoiceAnalytics.Shimmer]: The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics
type SFVoiceAnalytics struct {
	objectivec.Object
}

// SFVoiceAnalyticsFromID constructs a [SFVoiceAnalytics] from an objc.ID.
//
// A collection of vocal analysis metrics.
func SFVoiceAnalyticsFromID(id objc.ID) SFVoiceAnalytics {
	return SFVoiceAnalytics{objectivec.Object{ID: id}}
}

// NOTE: SFVoiceAnalytics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFVoiceAnalytics] class.
//
// # Analyzing voice
//
//   - [ISFVoiceAnalytics.Voicing]: The likelihood of a voice in each frame of a transcription segment.
//   - [ISFVoiceAnalytics.Pitch]: The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.
//   - [ISFVoiceAnalytics.Jitter]: The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.
//   - [ISFVoiceAnalytics.Shimmer]: The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics
type ISFVoiceAnalytics interface {
	objectivec.IObject

	// Topic: Analyzing voice

	// The likelihood of a voice in each frame of a transcription segment.
	Voicing() ISFAcousticFeature
	// The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.
	Pitch() ISFAcousticFeature
	// The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.
	Jitter() ISFAcousticFeature
	// The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.
	Shimmer() ISFAcousticFeature

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v SFVoiceAnalytics) Init() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v SFVoiceAnalytics) Autorelease() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFVoiceAnalytics creates a new SFVoiceAnalytics instance.
func NewSFVoiceAnalytics() SFVoiceAnalytics {
	class := getSFVoiceAnalyticsClass()
	rv := objc.Send[SFVoiceAnalytics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v SFVoiceAnalytics) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The likelihood of a voice in each frame of a transcription segment.
//
// # Discussion
//
// The `voicing` value is expressed as a probability in the range `[0.0,
// 1.0]`.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/voicing
func (v SFVoiceAnalytics) Voicing() ISFAcousticFeature {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("voicing"))
	return SFAcousticFeatureFromID(objc.ID(rv))
}

// The highness or lowness of the tone (fundamental frequency) in each frame
// of a transcription segment, expressed as a logarithm.
//
// # Discussion
//
// The value is a logarithm (base `e`) of the normalized pitch estimate for
// each frame.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/pitch
func (v SFVoiceAnalytics) Pitch() ISFAcousticFeature {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pitch"))
	return SFAcousticFeatureFromID(objc.ID(rv))
}

// The variation in pitch in each frame of a transcription segment, expressed
// as a percentage of the frame’s fundamental frequency.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/jitter
func (v SFVoiceAnalytics) Jitter() ISFAcousticFeature {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("jitter"))
	return SFAcousticFeatureFromID(objc.ID(rv))
}

// The variation in vocal volume stability (amplitude) in each frame of a
// transcription segment, expressed in decibels.
//
// See: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/shimmer
func (v SFVoiceAnalytics) Shimmer() ISFAcousticFeature {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("shimmer"))
	return SFAcousticFeatureFromID(objc.ID(rv))
}
