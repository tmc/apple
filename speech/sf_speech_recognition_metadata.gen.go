// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechRecognitionMetadata] class.
var (
	_SFSpeechRecognitionMetadataClass     SFSpeechRecognitionMetadataClass
	_SFSpeechRecognitionMetadataClassOnce sync.Once
)

func getSFSpeechRecognitionMetadataClass() SFSpeechRecognitionMetadataClass {
	_SFSpeechRecognitionMetadataClassOnce.Do(func() {
		_SFSpeechRecognitionMetadataClass = SFSpeechRecognitionMetadataClass{class: objc.GetClass("SFSpeechRecognitionMetadata")}
	})
	return _SFSpeechRecognitionMetadataClass
}

// GetSFSpeechRecognitionMetadataClass returns the class object for SFSpeechRecognitionMetadata.
func GetSFSpeechRecognitionMetadataClass() SFSpeechRecognitionMetadataClass {
	return getSFSpeechRecognitionMetadataClass()
}

type SFSpeechRecognitionMetadataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechRecognitionMetadataClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechRecognitionMetadataClass) Alloc() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The metadata of speech in the audio of a speech recognition request.
//
// # Getting audio timing information
//
//   - [SFSpeechRecognitionMetadata.AveragePauseDuration]: The average pause duration between words, measured in seconds.
//   - [SFSpeechRecognitionMetadata.SpeakingRate]: The number of words spoken per minute.
//   - [SFSpeechRecognitionMetadata.SpeechDuration]: The duration in seconds of speech in the audio.
//   - [SFSpeechRecognitionMetadata.SpeechStartTimestamp]: The start timestamp of speech in the audio.
//
// # Analyzing voice
//
//   - [SFSpeechRecognitionMetadata.VoiceAnalytics]: An analysis of the transcription segment’s vocal properties.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata
type SFSpeechRecognitionMetadata struct {
	objectivec.Object
}

// SFSpeechRecognitionMetadataFromID constructs a [SFSpeechRecognitionMetadata] from an objc.ID.
//
// The metadata of speech in the audio of a speech recognition request.
func SFSpeechRecognitionMetadataFromID(id objc.ID) SFSpeechRecognitionMetadata {
	return SFSpeechRecognitionMetadata{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechRecognitionMetadata adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechRecognitionMetadata] class.
//
// # Getting audio timing information
//
//   - [ISFSpeechRecognitionMetadata.AveragePauseDuration]: The average pause duration between words, measured in seconds.
//   - [ISFSpeechRecognitionMetadata.SpeakingRate]: The number of words spoken per minute.
//   - [ISFSpeechRecognitionMetadata.SpeechDuration]: The duration in seconds of speech in the audio.
//   - [ISFSpeechRecognitionMetadata.SpeechStartTimestamp]: The start timestamp of speech in the audio.
//
// # Analyzing voice
//
//   - [ISFSpeechRecognitionMetadata.VoiceAnalytics]: An analysis of the transcription segment’s vocal properties.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata
type ISFSpeechRecognitionMetadata interface {
	objectivec.IObject

	// Topic: Getting audio timing information

	// The average pause duration between words, measured in seconds.
	AveragePauseDuration() foundation.NSTimeInterval
	// The number of words spoken per minute.
	SpeakingRate() float64
	// The duration in seconds of speech in the audio.
	SpeechDuration() foundation.NSTimeInterval
	// The start timestamp of speech in the audio.
	SpeechStartTimestamp() foundation.NSTimeInterval

	// Topic: Analyzing voice

	// An analysis of the transcription segment’s vocal properties.
	VoiceAnalytics() ISFVoiceAnalytics

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s SFSpeechRecognitionMetadata) Init() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechRecognitionMetadata) Autorelease() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionMetadata creates a new SFSpeechRecognitionMetadata instance.
func NewSFSpeechRecognitionMetadata() SFSpeechRecognitionMetadata {
	class := getSFSpeechRecognitionMetadataClass()
	rv := objc.Send[SFSpeechRecognitionMetadata](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SFSpeechRecognitionMetadata) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The average pause duration between words, measured in seconds.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/averagePauseDuration
func (s SFSpeechRecognitionMetadata) AveragePauseDuration() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](s.ID, objc.Sel("averagePauseDuration"))
	return foundation.NSTimeInterval(rv)
}

// The number of words spoken per minute.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speakingRate
func (s SFSpeechRecognitionMetadata) SpeakingRate() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("speakingRate"))
	return rv
}

// The duration in seconds of speech in the audio.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speechDuration
func (s SFSpeechRecognitionMetadata) SpeechDuration() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](s.ID, objc.Sel("speechDuration"))
	return foundation.NSTimeInterval(rv)
}

// The start timestamp of speech in the audio.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speechStartTimestamp
func (s SFSpeechRecognitionMetadata) SpeechStartTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](s.ID, objc.Sel("speechStartTimestamp"))
	return foundation.NSTimeInterval(rv)
}

// An analysis of the transcription segment’s vocal properties.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/voiceAnalytics
func (s SFSpeechRecognitionMetadata) VoiceAnalytics() ISFVoiceAnalytics {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceAnalytics"))
	return SFVoiceAnalyticsFromID(objc.ID(rv))
}
