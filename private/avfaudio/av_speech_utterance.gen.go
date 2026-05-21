// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechUtterance] class.
var (
	_AVSpeechUtteranceClass     AVSpeechUtteranceClass
	_AVSpeechUtteranceClassOnce sync.Once
)

func getAVSpeechUtteranceClass() AVSpeechUtteranceClass {
	_AVSpeechUtteranceClassOnce.Do(func() {
		_AVSpeechUtteranceClass = AVSpeechUtteranceClass{class: objc.GetClass("AVSpeechUtterance")}
	})
	return _AVSpeechUtteranceClass
}

// GetAVSpeechUtteranceClass returns the class object for AVSpeechUtterance.
func GetAVSpeechUtteranceClass() AVSpeechUtteranceClass {
	return getAVSpeechUtteranceClass()
}

type AVSpeechUtteranceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechUtteranceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechUtteranceClass) Alloc() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVSpeechUtterance.Action]
//   - [AVSpeechUtterance.SetAction]
//   - [AVSpeechUtterance.AudioBufferCallback]
//   - [AVSpeechUtterance.MarkerCallback]
//   - [AVSpeechUtterance.PrefersAssistiveTechnologyExceptions]
//   - [AVSpeechUtterance.ProcessEmoticons]
//   - [AVSpeechUtterance.SetAudioBufferCallback]
//   - [AVSpeechUtterance.SetMarkerCallback]
//   - [AVSpeechUtterance.SetPrefersAssistiveTechnologyExceptions]
//   - [AVSpeechUtterance.SetProcessEmoticons]
//   - [AVSpeechUtterance.SetSsmlRepresentation]
//   - [AVSpeechUtterance.SetVoiceSelection]
//   - [AVSpeechUtterance.SsmlRepresentation]
//   - [AVSpeechUtterance.VoiceSelection]
//   - [AVSpeechUtterance.AttributedSpeechString]
//   - [AVSpeechUtterance.SetAttributedSpeechString]
//   - [AVSpeechUtterance.SpeechString]
//   - [AVSpeechUtterance.SetSpeechString]
type AVSpeechUtterance struct {
	objectivec.Object
}

// AVSpeechUtteranceFromID constructs a [AVSpeechUtterance] from an objc.ID.
func AVSpeechUtteranceFromID(id objc.ID) AVSpeechUtterance {
	return AVSpeechUtterance{objectivec.Object{ID: id}}
}

// Ensure AVSpeechUtterance implements IAVSpeechUtterance.
var _ IAVSpeechUtterance = AVSpeechUtterance{}

// An interface definition for the [AVSpeechUtterance] class.
//
// # Methods
//
//   - [IAVSpeechUtterance.Action]
//   - [IAVSpeechUtterance.SetAction]
//   - [IAVSpeechUtterance.AudioBufferCallback]
//   - [IAVSpeechUtterance.MarkerCallback]
//   - [IAVSpeechUtterance.PrefersAssistiveTechnologyExceptions]
//   - [IAVSpeechUtterance.ProcessEmoticons]
//   - [IAVSpeechUtterance.SetAudioBufferCallback]
//   - [IAVSpeechUtterance.SetMarkerCallback]
//   - [IAVSpeechUtterance.SetPrefersAssistiveTechnologyExceptions]
//   - [IAVSpeechUtterance.SetProcessEmoticons]
//   - [IAVSpeechUtterance.SetSsmlRepresentation]
//   - [IAVSpeechUtterance.SetVoiceSelection]
//   - [IAVSpeechUtterance.SsmlRepresentation]
//   - [IAVSpeechUtterance.VoiceSelection]
//   - [IAVSpeechUtterance.AttributedSpeechString]
//   - [IAVSpeechUtterance.SetAttributedSpeechString]
//   - [IAVSpeechUtterance.SpeechString]
//   - [IAVSpeechUtterance.SetSpeechString]
type IAVSpeechUtterance interface {
	objectivec.IObject

	// Topic: Methods

	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	AudioBufferCallback()
	MarkerCallback()
	PrefersAssistiveTechnologyExceptions() objectivec.IObject
	ProcessEmoticons() bool
	SetAudioBufferCallback(callback VoidHandler)
	SetMarkerCallback(callback VoidHandler)
	SetPrefersAssistiveTechnologyExceptions(exceptions objectivec.IObject)
	SetProcessEmoticons(emoticons bool)
	SetSsmlRepresentation(representation objectivec.IObject)
	SetVoiceSelection(selection objectivec.IObject)
	SsmlRepresentation() objectivec.IObject
	VoiceSelection() objectivec.IObject
	AttributedSpeechString() foundation.NSAttributedString
	SetAttributedSpeechString(value foundation.NSAttributedString)
	SpeechString() string
	SetSpeechString(value string)
}

// Init initializes the instance.
func (a AVSpeechUtterance) Init() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechUtterance) Autorelease() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechUtterance creates a new AVSpeechUtterance instance.
func NewAVSpeechUtterance() AVSpeechUtterance {
	class := getAVSpeechUtteranceClass()
	rv := objc.Send[AVSpeechUtterance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVSpeechUtterance) AudioBufferCallback() {
	objc.Send[objc.ID](a.ID, objc.Sel("audioBufferCallback"))
}
func (a AVSpeechUtterance) MarkerCallback() {
	objc.Send[objc.ID](a.ID, objc.Sel("markerCallback"))
}
func (a AVSpeechUtterance) PrefersAssistiveTechnologyExceptions() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("prefersAssistiveTechnologyExceptions"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechUtterance) ProcessEmoticons() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("processEmoticons"))
	return rv
}
func (a AVSpeechUtterance) SetAudioBufferCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](a.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}
func (a AVSpeechUtterance) SetMarkerCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](a.ID, objc.Sel("setMarkerCallback:"), _block0)
}
func (a AVSpeechUtterance) SetPrefersAssistiveTechnologyExceptions(exceptions objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setPrefersAssistiveTechnologyExceptions:"), exceptions)
}
func (a AVSpeechUtterance) SetProcessEmoticons(emoticons bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setProcessEmoticons:"), emoticons)
}
func (a AVSpeechUtterance) SetSsmlRepresentation(representation objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSsmlRepresentation:"), representation)
}
func (a AVSpeechUtterance) SetVoiceSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setVoiceSelection:"), selection)
}
func (a AVSpeechUtterance) SsmlRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("ssmlRepresentation"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechUtterance) VoiceSelection() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("voiceSelection"))
	return objectivec.Object{ID: rv}
}

func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) TransformUtteranceBasedOnSSMLIfDetected(detected objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("transformUtteranceBasedOnSSMLIfDetected:"), detected)
}

func (a AVSpeechUtterance) Action() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("action"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechUtterance) SetAction(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setAction:"), value)
}
func (a AVSpeechUtterance) AttributedSpeechString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedSpeechString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (a AVSpeechUtterance) SetAttributedSpeechString(value foundation.NSAttributedString) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributedSpeechString:"), value)
}
func (a AVSpeechUtterance) SpeechString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("speechString"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechUtterance) SetSpeechString(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSpeechString:"), objc.String(value))
}

// SetAudioBufferCallbackSync is a synchronous wrapper around [AVSpeechUtterance.SetAudioBufferCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVSpeechUtterance) SetAudioBufferCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetAudioBufferCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetMarkerCallbackSync is a synchronous wrapper around [AVSpeechUtterance.SetMarkerCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVSpeechUtterance) SetMarkerCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetMarkerCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
