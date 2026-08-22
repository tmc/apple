// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNClassificationResult] class.
var (
	_SNClassificationResultClass     SNClassificationResultClass
	_SNClassificationResultClassOnce sync.Once
)

func getSNClassificationResultClass() SNClassificationResultClass {
	_SNClassificationResultClassOnce.Do(func() {
		_SNClassificationResultClass = SNClassificationResultClass{class: objc.GetClass("SNClassificationResult")}
	})
	return _SNClassificationResultClass
}

// GetSNClassificationResultClass returns the class object for SNClassificationResult.
func GetSNClassificationResultClass() SNClassificationResultClass {
	return getSNClassificationResultClass()
}

type SNClassificationResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNClassificationResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNClassificationResultClass) Alloc() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A result that contains the highest-ranking classifications in a time range.
//
// # Overview
//
// An [SNClassificationResult] represents the predictions that a sound
// classification model made for a time span in an audio file or stream. Each
// result contains one or more classification predictions and a time range
// within the audio data.
//
// An audio analyzer, such as [SNAudioFileAnalyzer] and
// [SNAudioStreamAnalyzer], produces an [SNClassificationResult] each time it
// recognizes a sound for any of its [SNClassifySoundRequest] instances.
//
// # Inspecting the Result
//
//   - [SNClassificationResult.TimeRange]: The time span that corresponds to the result’s classifications.
//   - [SNClassificationResult.Classifications]: A sorted array of the request’s top classification candidates.
//   - [SNClassificationResult.ClassificationForIdentifier]: Returns the classification for an identifier.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult
type SNClassificationResult struct {
	objectivec.Object
}

// SNClassificationResultFromID constructs a [SNClassificationResult] from an objc.ID.
//
// A result that contains the highest-ranking classifications in a time range.
func SNClassificationResultFromID(id objc.ID) SNClassificationResult {
	return SNClassificationResult{objectivec.Object{ID: id}}
}

// NOTE: SNClassificationResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SNClassificationResult] class.
//
// # Inspecting the Result
//
//   - [ISNClassificationResult.TimeRange]: The time span that corresponds to the result’s classifications.
//   - [ISNClassificationResult.Classifications]: A sorted array of the request’s top classification candidates.
//   - [ISNClassificationResult.ClassificationForIdentifier]: Returns the classification for an identifier.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult
type ISNClassificationResult interface {
	objectivec.IObject
	SNResult

	// Topic: Inspecting the Result

	// The time span that corresponds to the result’s classifications.
	TimeRange() coremedia.CMTimeRange
	// A sorted array of the request’s top classification candidates.
	Classifications() []SNClassification
	// Returns the classification for an identifier.
	ClassificationForIdentifier(identifier string) ISNClassification
}

// Init initializes the instance.
func (c SNClassificationResult) Init() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SNClassificationResult) Autorelease() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassificationResult creates a new SNClassificationResult instance.
func NewSNClassificationResult() SNClassificationResult {
	class := getSNClassificationResultClass()
	rv := objc.Send[SNClassificationResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the classification for an identifier.
//
// identifier: A sound classification label.
//
// # Return Value
//
// A sound classification with a corresponding identifier if it exists in the
// result; otherwise, `nil`.
//
// # Discussion
//
// The `identifier` parameter corresponds to the [SNClassification.Identifier]
// property in an [SNClassification].
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/classification(forIdentifier:)
func (c SNClassificationResult) ClassificationForIdentifier(identifier string) ISNClassification {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("classificationForIdentifier:"), objc.String(identifier))
	return SNClassificationFromID(rv)
}

// The time span that corresponds to the result’s classifications.
//
// # Discussion
//
// The time range’s [CMTime] values are the number of audio frames at the
// analyzer’s sample rate. Use these time indices to determine where, in
// time, the result corresponds to the original audio.
//
// A result’s time range typically refers to audio older than its most
// recent audio because the request gathers the data into blocks before
// sending them to the model.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/timeRange
//
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
func (c SNClassificationResult) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](c.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}

// A sorted array of the request’s top classification candidates.
//
// # Discussion
//
// [SNClassificationResult] sorts its classifications in descending confidence
// score order.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/classifications
func (c SNClassificationResult) Classifications() []SNClassification {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("classifications"))
	return objc.ConvertSlice(rv, func(id objc.ID) SNClassification {
		return SNClassificationFromID(id)
	})
}

// Protocol methods for SNResult
