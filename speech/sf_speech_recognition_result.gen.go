// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechRecognitionResult] class.
var (
	_SFSpeechRecognitionResultClass     SFSpeechRecognitionResultClass
	_SFSpeechRecognitionResultClassOnce sync.Once
)

func getSFSpeechRecognitionResultClass() SFSpeechRecognitionResultClass {
	_SFSpeechRecognitionResultClassOnce.Do(func() {
		_SFSpeechRecognitionResultClass = SFSpeechRecognitionResultClass{class: objc.GetClass("SFSpeechRecognitionResult")}
	})
	return _SFSpeechRecognitionResultClass
}

// GetSFSpeechRecognitionResultClass returns the class object for SFSpeechRecognitionResult.
func GetSFSpeechRecognitionResultClass() SFSpeechRecognitionResultClass {
	return getSFSpeechRecognitionResultClass()
}

type SFSpeechRecognitionResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechRecognitionResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechRecognitionResultClass) Alloc() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the partial or final results of a speech
// recognition request.
//
// # Overview
//
// Use an [SFSpeechRecognitionResult] object to retrieve the results of a
// speech recognition request. You don’t create these objects directly.
// Instead, the Speech framework creates them and passes them to the handler
// block or delegate object you specified when starting your speech
// recognition task.
//
// A speech recognition result object contains one or more
// [SFSpeechRecognitionResult.Transcriptions] of the current utterance. Each
// transcription has a confidence rating indicating how likely it is to be
// correct. You can also get the transcription with the highest rating
// directly from the [SFSpeechRecognitionResult.BestTranscription] property.
//
// If you requested partial results from the speech recognizer, the
// transcriptions may represent only part of the total audio content. Use the
// [SFSpeechRecognitionResult.Final] property to determine if the request
// contains partial or final results.
//
// # Getting transcriptions
//
//   - [SFSpeechRecognitionResult.BestTranscription]: The transcription with the highest confidence level.
//   - [SFSpeechRecognitionResult.Transcriptions]: An array of potential transcriptions, sorted in descending order of confidence.
//   - [SFSpeechRecognitionResult.SpeechRecognitionMetadata]: An object that contains the metadata results for a speech recognition request.
//
// # Determining whether transcriptions are final
//
//   - [SFSpeechRecognitionResult.IsFinal]: A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult
type SFSpeechRecognitionResult struct {
	objectivec.Object
}

// SFSpeechRecognitionResultFromID constructs a [SFSpeechRecognitionResult] from an objc.ID.
//
// An object that contains the partial or final results of a speech
// recognition request.
func SFSpeechRecognitionResultFromID(id objc.ID) SFSpeechRecognitionResult {
	return SFSpeechRecognitionResult{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechRecognitionResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechRecognitionResult] class.
//
// # Getting transcriptions
//
//   - [ISFSpeechRecognitionResult.BestTranscription]: The transcription with the highest confidence level.
//   - [ISFSpeechRecognitionResult.Transcriptions]: An array of potential transcriptions, sorted in descending order of confidence.
//   - [ISFSpeechRecognitionResult.SpeechRecognitionMetadata]: An object that contains the metadata results for a speech recognition request.
//
// # Determining whether transcriptions are final
//
//   - [ISFSpeechRecognitionResult.IsFinal]: A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult
type ISFSpeechRecognitionResult interface {
	objectivec.IObject

	// Topic: Getting transcriptions

	// The transcription with the highest confidence level.
	BestTranscription() ISFTranscription
	// An array of potential transcriptions, sorted in descending order of confidence.
	Transcriptions() []SFTranscription
	// An object that contains the metadata results for a speech recognition request.
	SpeechRecognitionMetadata() ISFSpeechRecognitionMetadata

	// Topic: Determining whether transcriptions are final

	// A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
	IsFinal() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s SFSpeechRecognitionResult) Init() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechRecognitionResult) Autorelease() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionResult creates a new SFSpeechRecognitionResult instance.
func NewSFSpeechRecognitionResult() SFSpeechRecognitionResult {
	class := getSFSpeechRecognitionResultClass()
	rv := objc.Send[SFSpeechRecognitionResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SFSpeechRecognitionResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The transcription with the highest confidence level.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult/bestTranscription
func (s SFSpeechRecognitionResult) BestTranscription() ISFTranscription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bestTranscription"))
	return SFTranscriptionFromID(objc.ID(rv))
}

// An array of potential transcriptions, sorted in descending order of
// confidence.
//
// # Discussion
//
// All transcriptions correspond to the same utterance, which can be a partial
// or final result of the overall request. The first transcription in the
// array has the highest confidence rating, followed by transcriptions with
// decreasing confidence ratings.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult/transcriptions
func (s SFSpeechRecognitionResult) Transcriptions() []SFTranscription {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("transcriptions"))
	return objc.ConvertSlice(rv, func(id objc.ID) SFTranscription {
		return SFTranscriptionFromID(id)
	})
}

// An object that contains the metadata results for a speech recognition
// request.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult/speechRecognitionMetadata
func (s SFSpeechRecognitionResult) SpeechRecognitionMetadata() ISFSpeechRecognitionMetadata {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechRecognitionMetadata"))
	return SFSpeechRecognitionMetadataFromID(objc.ID(rv))
}

// A Boolean value that indicates whether speech recognition is complete and
// whether the transcriptions are final.
//
// # Discussion
//
// When a speech recognition request is final, its transcriptions don’t
// change.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult/isFinal
func (s SFSpeechRecognitionResult) IsFinal() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isFinal"))
	return rv
}
