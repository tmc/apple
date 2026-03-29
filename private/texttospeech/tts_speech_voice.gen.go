// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TTSSpeechVoice] class.
var (
	_TTSSpeechVoiceClass     TTSSpeechVoiceClass
	_TTSSpeechVoiceClassOnce sync.Once
)

func getTTSSpeechVoiceClass() TTSSpeechVoiceClass {
	_TTSSpeechVoiceClassOnce.Do(func() {
		_TTSSpeechVoiceClass = TTSSpeechVoiceClass{class: objc.GetClass("TTSSpeechVoice")}
	})
	return _TTSSpeechVoiceClass
}

// GetTTSSpeechVoiceClass returns the class object for TTSSpeechVoice.
func GetTTSSpeechVoiceClass() TTSSpeechVoiceClass {
	return getTTSSpeechVoiceClass()
}

type TTSSpeechVoiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechVoiceClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechVoiceClass) Alloc() TTSSpeechVoice {
	rv := objc.Send[TTSSpeechVoice](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechVoice
type TTSSpeechVoice struct {
	TTSAXResource
}

// TTSSpeechVoiceFromID constructs a [TTSSpeechVoice] from an objc.ID.
func TTSSpeechVoiceFromID(id objc.ID) TTSSpeechVoice {
	return TTSSpeechVoice{TTSAXResource: TTSAXResourceFromID(id)}
}
// Ensure TTSSpeechVoice implements ITTSSpeechVoice.
var _ ITTSSpeechVoice = TTSSpeechVoice{}

// An interface definition for the [TTSSpeechVoice] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechVoice
type ITTSSpeechVoice interface {
	ITTSAXResource
}

// Init initializes the instance.
func (t TTSSpeechVoice) Init() TTSSpeechVoice {
	rv := objc.Send[TTSSpeechVoice](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechVoice) Autorelease() TTSSpeechVoice {
	rv := objc.Send[TTSSpeechVoice](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechVoice creates a new TTSSpeechVoice instance.
func NewTTSSpeechVoice() TTSSpeechVoice {
	class := getTTSSpeechVoiceClass()
	rv := objc.Send[TTSSpeechVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

