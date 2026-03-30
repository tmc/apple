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
//   - [AVSpeechUtterance.InitWithCoder]
//   - [AVSpeechUtterance.AttributedSpeechString]
//   - [AVSpeechUtterance.SetAttributedSpeechString]
//   - [AVSpeechUtterance.SpeechString]
//   - [AVSpeechUtterance.SetSpeechString]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance
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
//   - [IAVSpeechUtterance.InitWithCoder]
//   - [IAVSpeechUtterance.AttributedSpeechString]
//   - [IAVSpeechUtterance.SetAttributedSpeechString]
//   - [IAVSpeechUtterance.SpeechString]
//   - [IAVSpeechUtterance.SetSpeechString]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance
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
	InitWithCoder(coder foundation.INSCoder) AVSpeechUtterance
	AttributedSpeechString() foundation.NSAttributedString
	SetAttributedSpeechString(value foundation.NSAttributedString)
	SpeechString() string
	SetSpeechString(value string)
}

// Init initializes the instance.
func (s AVSpeechUtterance) Init() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechUtterance) Autorelease() AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechUtterance creates a new AVSpeechUtterance instance.
func NewAVSpeechUtterance() AVSpeechUtterance {
	class := getAVSpeechUtteranceClass()
	rv := objc.Send[AVSpeechUtterance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/initWithCoder:
func NewSpeechUtteranceWithCoder(coder objectivec.IObject) AVSpeechUtterance {
	instance := getAVSpeechUtteranceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVSpeechUtteranceFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/audioBufferCallback
func (s AVSpeechUtterance) AudioBufferCallback() {
	objc.Send[objc.ID](s.ID, objc.Sel("audioBufferCallback"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/markerCallback
func (s AVSpeechUtterance) MarkerCallback() {
	objc.Send[objc.ID](s.ID, objc.Sel("markerCallback"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/prefersAssistiveTechnologyExceptions
func (s AVSpeechUtterance) PrefersAssistiveTechnologyExceptions() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("prefersAssistiveTechnologyExceptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/processEmoticons
func (s AVSpeechUtterance) ProcessEmoticons() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("processEmoticons"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setAudioBufferCallback:
func (s AVSpeechUtterance) SetAudioBufferCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](s.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setMarkerCallback:
func (s AVSpeechUtterance) SetMarkerCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](s.ID, objc.Sel("setMarkerCallback:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setPrefersAssistiveTechnologyExceptions:
func (s AVSpeechUtterance) SetPrefersAssistiveTechnologyExceptions(exceptions objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setPrefersAssistiveTechnologyExceptions:"), exceptions)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setProcessEmoticons:
func (s AVSpeechUtterance) SetProcessEmoticons(emoticons bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setProcessEmoticons:"), emoticons)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setSsmlRepresentation:
func (s AVSpeechUtterance) SetSsmlRepresentation(representation objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSsmlRepresentation:"), representation)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/setVoiceSelection:
func (s AVSpeechUtterance) SetVoiceSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setVoiceSelection:"), selection)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/ssmlRepresentation
func (s AVSpeechUtterance) SsmlRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("ssmlRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/voiceSelection
func (s AVSpeechUtterance) VoiceSelection() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceSelection"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/initWithCoder:
func (s AVSpeechUtterance) InitWithCoder(coder foundation.INSCoder) AVSpeechUtterance {
	rv := objc.Send[AVSpeechUtterance](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/supportsSecureCoding
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/transformUtteranceBasedOnSSMLIfDetected:
func (_AVSpeechUtteranceClass AVSpeechUtteranceClass) TransformUtteranceBasedOnSSMLIfDetected(detected objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechUtteranceClass.class), objc.Sel("transformUtteranceBasedOnSSMLIfDetected:"), detected)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/action
func (s AVSpeechUtterance) Action() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("action"))
	return objectivec.Object{ID: rv}
}
func (s AVSpeechUtterance) SetAction(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setAction:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/attributedSpeechString
func (s AVSpeechUtterance) AttributedSpeechString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attributedSpeechString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (s AVSpeechUtterance) SetAttributedSpeechString(value foundation.NSAttributedString) {
	objc.Send[struct{}](s.ID, objc.Sel("setAttributedSpeechString:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechUtterance/speechString
func (s AVSpeechUtterance) SpeechString() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechString"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechUtterance) SetSpeechString(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechString:"), objc.String(value))
}

// SetAudioBufferCallbackSync is a synchronous wrapper around [AVSpeechUtterance.SetAudioBufferCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSpeechUtterance) SetAudioBufferCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetAudioBufferCallback(func() {
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
func (s AVSpeechUtterance) SetMarkerCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetMarkerCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
