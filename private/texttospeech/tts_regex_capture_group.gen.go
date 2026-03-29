// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRegexCaptureGroup] class.
var (
	_TTSRegexCaptureGroupClass     TTSRegexCaptureGroupClass
	_TTSRegexCaptureGroupClassOnce sync.Once
)

func getTTSRegexCaptureGroupClass() TTSRegexCaptureGroupClass {
	_TTSRegexCaptureGroupClassOnce.Do(func() {
		_TTSRegexCaptureGroupClass = TTSRegexCaptureGroupClass{class: objc.GetClass("TTSRegexCaptureGroup")}
	})
	return _TTSRegexCaptureGroupClass
}

// GetTTSRegexCaptureGroupClass returns the class object for TTSRegexCaptureGroup.
func GetTTSRegexCaptureGroupClass() TTSRegexCaptureGroupClass {
	return getTTSRegexCaptureGroupClass()
}

type TTSRegexCaptureGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSRegexCaptureGroupClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRegexCaptureGroupClass) Alloc() TTSRegexCaptureGroup {
	rv := objc.Send[TTSRegexCaptureGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRegexCaptureGroup.Utf8Range]
//   - [TTSRegexCaptureGroup.SetUtf8Range]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCaptureGroup
type TTSRegexCaptureGroup struct {
	objectivec.Object
}

// TTSRegexCaptureGroupFromID constructs a [TTSRegexCaptureGroup] from an objc.ID.
func TTSRegexCaptureGroupFromID(id objc.ID) TTSRegexCaptureGroup {
	return TTSRegexCaptureGroup{objectivec.Object{ID: id}}
}
// Ensure TTSRegexCaptureGroup implements ITTSRegexCaptureGroup.
var _ ITTSRegexCaptureGroup = TTSRegexCaptureGroup{}

// An interface definition for the [TTSRegexCaptureGroup] class.
//
// # Methods
//
//   - [ITTSRegexCaptureGroup.Utf8Range]
//   - [ITTSRegexCaptureGroup.SetUtf8Range]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCaptureGroup
type ITTSRegexCaptureGroup interface {
	objectivec.IObject

	// Topic: Methods

	Utf8Range() foundation.NSRange
	SetUtf8Range(value foundation.NSRange)
}

// Init initializes the instance.
func (t TTSRegexCaptureGroup) Init() TTSRegexCaptureGroup {
	rv := objc.Send[TTSRegexCaptureGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRegexCaptureGroup) Autorelease() TTSRegexCaptureGroup {
	rv := objc.Send[TTSRegexCaptureGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRegexCaptureGroup creates a new TTSRegexCaptureGroup instance.
func NewTTSRegexCaptureGroup() TTSRegexCaptureGroup {
	class := getTTSRegexCaptureGroupClass()
	rv := objc.Send[TTSRegexCaptureGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCaptureGroup/utf8Range
func (t TTSRegexCaptureGroup) Utf8Range() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("utf8Range"))
	return foundation.NSRange(rv)
}
func (t TTSRegexCaptureGroup) SetUtf8Range(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setUtf8Range:"), value)
}

