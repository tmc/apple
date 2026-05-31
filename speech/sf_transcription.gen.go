// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFTranscription] class.
var (
	_SFTranscriptionClass     SFTranscriptionClass
	_SFTranscriptionClassOnce sync.Once
)

func getSFTranscriptionClass() SFTranscriptionClass {
	_SFTranscriptionClassOnce.Do(func() {
		_SFTranscriptionClass = SFTranscriptionClass{class: objc.GetClass("SFTranscription")}
	})
	return _SFTranscriptionClass
}

// GetSFTranscriptionClass returns the class object for SFTranscription.
func GetSFTranscriptionClass() SFTranscriptionClass {
	return getSFTranscriptionClass()
}

type SFTranscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFTranscriptionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFTranscriptionClass) Alloc() SFTranscription {
	rv := objc.Send[SFTranscription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A textual representation of the specified speech in its entirety, as
// recognized by the speech recognizer.
//
// # Overview
//
// Use [SFTranscription] to obtain all the recognized utterances from your
// audio content. An is a vocalized word or group of words that represent a
// single meaning to the speech recognizer ([SFSpeechRecognizer]).
//
// Use the [SFTranscription.FormattedString] property to retrieve the entire
// transcription of utterances, or use the [SFTranscription.Segments] property
// to retrieve an individual utterance ([SFTranscriptionSegment]).
//
// You don’t create an [SFTranscription] directly. Instead, you retrieve it
// from an [SFSpeechRecognitionResult] instance. The speech recognizer sends a
// speech recognition result to your app in one of two ways, depending on how
// your app started a speech recognition task.
//
// You can start a speech recognition task by using the speech recognizer’s
// [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler] method. When
// the task is complete, the speech recognizer sends an
// [SFSpeechRecognitionResult] instance to your `resultHandler` closure.
// Alternatively, you can use the speech recognizer’s
// [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate] method to start a
// speech recognition task. When the task is complete, the speech recognizer
// uses your [SFSpeechRecognitionTaskDelegate] to send an
// [SFSpeechRecognitionResult] by using the delegate’s
// [SpeechRecognitionTaskDidFinishRecognition] method.
//
// An [SFTranscription] represents only a potential version of the speech. It
// might not be an accurate representation of the utterances.
//
// # Transcribing utterances
//
//   - [SFTranscription.FormattedString]: The entire transcription of utterances, formatted into a single, user-displayable string.
//
// # Getting individual utterances
//
//   - [SFTranscription.Segments]: An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscription
type SFTranscription struct {
	objectivec.Object
}

// SFTranscriptionFromID constructs a [SFTranscription] from an objc.ID.
//
// A textual representation of the specified speech in its entirety, as
// recognized by the speech recognizer.
func SFTranscriptionFromID(id objc.ID) SFTranscription {
	return SFTranscription{objectivec.Object{ID: id}}
}

// NOTE: SFTranscription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFTranscription] class.
//
// # Transcribing utterances
//
//   - [ISFTranscription.FormattedString]: The entire transcription of utterances, formatted into a single, user-displayable string.
//
// # Getting individual utterances
//
//   - [ISFTranscription.Segments]: An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscription
type ISFTranscription interface {
	objectivec.IObject

	// Topic: Transcribing utterances

	// The entire transcription of utterances, formatted into a single, user-displayable string.
	FormattedString() string

	// Topic: Getting individual utterances

	// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
	Segments() []SFTranscriptionSegment

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SFTranscription) Init() SFTranscription {
	rv := objc.Send[SFTranscription](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SFTranscription) Autorelease() SFTranscription {
	rv := objc.Send[SFTranscription](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFTranscription creates a new SFTranscription instance.
func NewSFTranscription() SFTranscription {
	class := getSFTranscriptionClass()
	rv := objc.Send[SFTranscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t SFTranscription) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The entire transcription of utterances, formatted into a single,
// user-displayable string.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscription/formattedString
func (t SFTranscription) FormattedString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("formattedString"))
	return foundation.NSStringFromID(rv).String()
}

// An array of transcription segments that represent the parts of the
// transcription, as identified by the speech recognizer.
//
// # Discussion
//
// The order of the segments in the array matches the order in which the
// corresponding utterances occur in the spoken content.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscription/segments
func (t SFTranscription) Segments() []SFTranscriptionSegment {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("segments"))
	return objc.ConvertSlice(rv, func(id objc.ID) SFTranscriptionSegment {
		return SFTranscriptionSegmentFromID(id)
	})
}
