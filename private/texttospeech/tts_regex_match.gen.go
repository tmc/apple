// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRegexMatch] class.
var (
	_TTSRegexMatchClass     TTSRegexMatchClass
	_TTSRegexMatchClassOnce sync.Once
)

func getTTSRegexMatchClass() TTSRegexMatchClass {
	_TTSRegexMatchClassOnce.Do(func() {
		_TTSRegexMatchClass = TTSRegexMatchClass{class: objc.GetClass("TTSRegexMatch")}
	})
	return _TTSRegexMatchClass
}

// GetTTSRegexMatchClass returns the class object for TTSRegexMatch.
func GetTTSRegexMatchClass() TTSRegexMatchClass {
	return getTTSRegexMatchClass()
}

type TTSRegexMatchClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRegexMatchClass) Alloc() TTSRegexMatch {
	rv := objc.Send[TTSRegexMatch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRegexMatch.CaptureGroups]
//   - [TTSRegexMatch.SetCaptureGroups]
//   - [TTSRegexMatch.Utf8Range]
//   - [TTSRegexMatch.SetUtf8Range]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexMatch
type TTSRegexMatch struct {
	objectivec.Object
}

// TTSRegexMatchFromID constructs a [TTSRegexMatch] from an objc.ID.
func TTSRegexMatchFromID(id objc.ID) TTSRegexMatch {
	return TTSRegexMatch{objectivec.Object{ID: id}}
}
// Ensure TTSRegexMatch implements ITTSRegexMatch.
var _ ITTSRegexMatch = TTSRegexMatch{}

// An interface definition for the [TTSRegexMatch] class.
//
// # Methods
//
//   - [ITTSRegexMatch.CaptureGroups]
//   - [ITTSRegexMatch.SetCaptureGroups]
//   - [ITTSRegexMatch.Utf8Range]
//   - [ITTSRegexMatch.SetUtf8Range]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexMatch
type ITTSRegexMatch interface {
	objectivec.IObject

	// Topic: Methods

	CaptureGroups() foundation.INSArray
	SetCaptureGroups(value foundation.INSArray)
	Utf8Range() foundation.NSRange
	SetUtf8Range(value foundation.NSRange)
}

// Init initializes the instance.
func (t TTSRegexMatch) Init() TTSRegexMatch {
	rv := objc.Send[TTSRegexMatch](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRegexMatch) Autorelease() TTSRegexMatch {
	rv := objc.Send[TTSRegexMatch](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRegexMatch creates a new TTSRegexMatch instance.
func NewTTSRegexMatch() TTSRegexMatch {
	class := getTTSRegexMatchClass()
	rv := objc.Send[TTSRegexMatch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexMatch/captureGroups
func (t TTSRegexMatch) CaptureGroups() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("captureGroups"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSRegexMatch) SetCaptureGroups(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setCaptureGroups:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexMatch/utf8Range
func (t TTSRegexMatch) Utf8Range() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("utf8Range"))
	return foundation.NSRange(rv)
}
func (t TTSRegexMatch) SetUtf8Range(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setUtf8Range:"), value)
}

