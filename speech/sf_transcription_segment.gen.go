// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFTranscriptionSegment] class.
var (
	_SFTranscriptionSegmentClass     SFTranscriptionSegmentClass
	_SFTranscriptionSegmentClassOnce sync.Once
)

func getSFTranscriptionSegmentClass() SFTranscriptionSegmentClass {
	_SFTranscriptionSegmentClassOnce.Do(func() {
		_SFTranscriptionSegmentClass = SFTranscriptionSegmentClass{class: objc.GetClass("SFTranscriptionSegment")}
	})
	return _SFTranscriptionSegmentClass
}

// GetSFTranscriptionSegmentClass returns the class object for SFTranscriptionSegment.
func GetSFTranscriptionSegmentClass() SFTranscriptionSegmentClass {
	return getSFTranscriptionSegmentClass()
}

type SFTranscriptionSegmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFTranscriptionSegmentClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFTranscriptionSegmentClass) Alloc() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A discrete part of an entire transcription, as identified by the speech
// recognizer.
//
// # Overview
//
// Use [SFTranscriptionSegment] to get details about a part of an overall
// [SFTranscription]. An [SFTranscriptionSegment] represents an utterance,
// which is a vocalized word or group of words that represent a single meaning
// to the speech recognizer ([SFSpeechRecognizer]).
//
// You don’t create transcription object segments directly. Instead, you
// access them from a transcription’s [SFTranscription.Segments] property.
//
// A transcription segment includes the following information:
//
// - The text of the utterance, plus any alternative interpretations of the
// spoken word. - The character range of the segment within the
// [SFTranscription.FormattedString] of its parent [SFTranscription]. - A
// [SFTranscriptionSegment.Confidence] value, indicating how likely it is that
// the specified string matches the audible speech. - A
// [SFTranscriptionSegment.Timestamp] and [SFTranscriptionSegment.Duration]
// value, indicating the position of the segment within the provided audio
// stream.
//
// # Transcribing the segment
//
//   - [SFTranscriptionSegment.Substring]: The string representation of the utterance in the transcription segment.
//   - [SFTranscriptionSegment.SubstringRange]: The range information for the transcription segment’s substring, relative to the overall transcription.
//   - [SFTranscriptionSegment.AlternativeSubstrings]: An array of alternate interpretations of the utterance in the transcription segment.
//
// # Assessing the recognition confidence level
//
//   - [SFTranscriptionSegment.Confidence]: The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// # Getting audio timing information
//
//   - [SFTranscriptionSegment.Timestamp]: The start time of the segment in the processed audio stream.
//   - [SFTranscriptionSegment.Duration]: The number of seconds it took for the user to speak the utterance represented by the segment.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment
type SFTranscriptionSegment struct {
	objectivec.Object
}

// SFTranscriptionSegmentFromID constructs a [SFTranscriptionSegment] from an objc.ID.
//
// A discrete part of an entire transcription, as identified by the speech
// recognizer.
func SFTranscriptionSegmentFromID(id objc.ID) SFTranscriptionSegment {
	return SFTranscriptionSegment{objectivec.Object{ID: id}}
}

// NOTE: SFTranscriptionSegment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFTranscriptionSegment] class.
//
// # Transcribing the segment
//
//   - [ISFTranscriptionSegment.Substring]: The string representation of the utterance in the transcription segment.
//   - [ISFTranscriptionSegment.SubstringRange]: The range information for the transcription segment’s substring, relative to the overall transcription.
//   - [ISFTranscriptionSegment.AlternativeSubstrings]: An array of alternate interpretations of the utterance in the transcription segment.
//
// # Assessing the recognition confidence level
//
//   - [ISFTranscriptionSegment.Confidence]: The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// # Getting audio timing information
//
//   - [ISFTranscriptionSegment.Timestamp]: The start time of the segment in the processed audio stream.
//   - [ISFTranscriptionSegment.Duration]: The number of seconds it took for the user to speak the utterance represented by the segment.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment
type ISFTranscriptionSegment interface {
	objectivec.IObject

	// Topic: Transcribing the segment

	// The string representation of the utterance in the transcription segment.
	Substring() string
	// The range information for the transcription segment’s substring, relative to the overall transcription.
	SubstringRange() foundation.NSRange
	// An array of alternate interpretations of the utterance in the transcription segment.
	AlternativeSubstrings() []string

	// Topic: Assessing the recognition confidence level

	// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
	Confidence() float32

	// Topic: Getting audio timing information

	// The start time of the segment in the processed audio stream.
	Timestamp() foundation.NSTimeInterval
	// The number of seconds it took for the user to speak the utterance represented by the segment.
	Duration() foundation.NSTimeInterval

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SFTranscriptionSegment) Init() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SFTranscriptionSegment) Autorelease() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFTranscriptionSegment creates a new SFTranscriptionSegment instance.
func NewSFTranscriptionSegment() SFTranscriptionSegment {
	class := getSFTranscriptionSegmentClass()
	rv := objc.Send[SFTranscriptionSegment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t SFTranscriptionSegment) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The string representation of the utterance in the transcription segment.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/substring
func (t SFTranscriptionSegment) Substring() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("substring"))
	return foundation.NSStringFromID(rv).String()
}

// The range information for the transcription segment’s substring, relative
// to the overall transcription.
//
// # Discussion
//
// Use the range information to find the position of the segment within the
// [SFTranscription.FormattedString] property of the [SFTranscription] object
// containing this segment.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/substringRange
func (t SFTranscriptionSegment) SubstringRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("substringRange"))
	return foundation.NSRange(rv)
}

// An array of alternate interpretations of the utterance in the transcription
// segment.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/alternativeSubstrings
func (t SFTranscriptionSegment) AlternativeSubstrings() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("alternativeSubstrings"))
	return objc.ConvertSliceToStrings(rv)
}

// The level of confidence the speech recognizer has in its recognition of the
// speech transcribed for the segment.
//
// # Discussion
//
// This property reflects the overall confidence in the recognition of the
// entire phrase. The value is `0` if there was no recognition, and it is
// closer to `1` when there is a high certainty that a transcription matches
// the user’s speech exactly. For example, a confidence value of `0.94`
// represents a very high confidence level, and is more likely to be correct
// than a transcription with a confidence value of `0.72`.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/confidence
func (t SFTranscriptionSegment) Confidence() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("confidence"))
	return rv
}

// The start time of the segment in the processed audio stream.
//
// # Discussion
//
// The [SFTranscriptionSegment.Timestamp] is the number of seconds between the
// beginning of the audio content and when the user spoke the word represented
// by the segment. For example, if the user said the word “time” one
// second into the transcription “What time is it”, the timestamp would be
// equal to `1.0`.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/timestamp
func (t SFTranscriptionSegment) Timestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](t.ID, objc.Sel("timestamp"))
	return foundation.NSTimeInterval(rv)
}

// The number of seconds it took for the user to speak the utterance
// represented by the segment.
//
// # Discussion
//
// The [SFTranscriptionSegment.Duration] contains the number of seconds it
// took for the user to speak the one or more words (utterance) represented by
// the segment. For example, the [SFSpeechRecognizer] sets
// [SFTranscriptionSegment.Duration] to `0.6` if the user took `0.6` seconds
// to say `“time”` in the transcription of `“What time is it?"`.
//
// See: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/duration
func (t SFTranscriptionSegment) Duration() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](t.ID, objc.Sel("duration"))
	return foundation.NSTimeInterval(rv)
}
