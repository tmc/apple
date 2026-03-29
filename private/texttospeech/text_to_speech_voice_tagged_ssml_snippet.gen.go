// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceTaggedSSMLSnippet] class.
var (
	_TextToSpeechVoiceTaggedSSMLSnippetClass     TextToSpeechVoiceTaggedSSMLSnippetClass
	_TextToSpeechVoiceTaggedSSMLSnippetClassOnce sync.Once
)

func getTextToSpeechVoiceTaggedSSMLSnippetClass() TextToSpeechVoiceTaggedSSMLSnippetClass {
	_TextToSpeechVoiceTaggedSSMLSnippetClassOnce.Do(func() {
		_TextToSpeechVoiceTaggedSSMLSnippetClass = TextToSpeechVoiceTaggedSSMLSnippetClass{class: objc.GetClass("TextToSpeech.VoiceTaggedSSMLSnippet")}
	})
	return _TextToSpeechVoiceTaggedSSMLSnippetClass
}

// GetTextToSpeechVoiceTaggedSSMLSnippetClass returns the class object for TextToSpeech.VoiceTaggedSSMLSnippet.
func GetTextToSpeechVoiceTaggedSSMLSnippetClass() TextToSpeechVoiceTaggedSSMLSnippetClass {
	return getTextToSpeechVoiceTaggedSSMLSnippetClass()
}

type TextToSpeechVoiceTaggedSSMLSnippetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceTaggedSSMLSnippetClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceTaggedSSMLSnippetClass) Alloc() TextToSpeechVoiceTaggedSSMLSnippet {
	rv := objc.Send[TextToSpeechVoiceTaggedSSMLSnippet](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechVoiceTaggedSSMLSnippet.Language]
//   - [TextToSpeechVoiceTaggedSSMLSnippet.SetLanguage]
//   - [TextToSpeechVoiceTaggedSSMLSnippet.Ssml]
//   - [TextToSpeechVoiceTaggedSSMLSnippet.SetSsml]
//   - [TextToSpeechVoiceTaggedSSMLSnippet.VoiceName]
//   - [TextToSpeechVoiceTaggedSSMLSnippet.SetVoiceName]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceTaggedSSMLSnippet
type TextToSpeechVoiceTaggedSSMLSnippet struct {
	objectivec.Object
}

// TextToSpeechVoiceTaggedSSMLSnippetFromID constructs a [TextToSpeechVoiceTaggedSSMLSnippet] from an objc.ID.
func TextToSpeechVoiceTaggedSSMLSnippetFromID(id objc.ID) TextToSpeechVoiceTaggedSSMLSnippet {
	return TextToSpeechVoiceTaggedSSMLSnippet{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechVoiceTaggedSSMLSnippet implements ITextToSpeechVoiceTaggedSSMLSnippet.
var _ ITextToSpeechVoiceTaggedSSMLSnippet = TextToSpeechVoiceTaggedSSMLSnippet{}

// An interface definition for the [TextToSpeechVoiceTaggedSSMLSnippet] class.
//
// # Methods
//
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.Language]
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.SetLanguage]
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.Ssml]
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.SetSsml]
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.VoiceName]
//   - [ITextToSpeechVoiceTaggedSSMLSnippet.SetVoiceName]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceTaggedSSMLSnippet
type ITextToSpeechVoiceTaggedSSMLSnippet interface {
	objectivec.IObject

	// Topic: Methods

	Language() string
	SetLanguage(value string)
	Ssml() string
	SetSsml(value string)
	VoiceName() string
	SetVoiceName(value string)
}

// Init initializes the instance.
func (t TextToSpeechVoiceTaggedSSMLSnippet) Init() TextToSpeechVoiceTaggedSSMLSnippet {
	rv := objc.Send[TextToSpeechVoiceTaggedSSMLSnippet](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceTaggedSSMLSnippet) Autorelease() TextToSpeechVoiceTaggedSSMLSnippet {
	rv := objc.Send[TextToSpeechVoiceTaggedSSMLSnippet](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceTaggedSSMLSnippet creates a new TextToSpeechVoiceTaggedSSMLSnippet instance.
func NewTextToSpeechVoiceTaggedSSMLSnippet() TextToSpeechVoiceTaggedSSMLSnippet {
	class := getTextToSpeechVoiceTaggedSSMLSnippetClass()
	rv := objc.Send[TextToSpeechVoiceTaggedSSMLSnippet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceTaggedSSMLSnippet/language
func (t TextToSpeechVoiceTaggedSSMLSnippet) Language() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechVoiceTaggedSSMLSnippet) SetLanguage(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLanguage:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceTaggedSSMLSnippet/ssml
func (t TextToSpeechVoiceTaggedSSMLSnippet) Ssml() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ssml"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechVoiceTaggedSSMLSnippet) SetSsml(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setSsml:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceTaggedSSMLSnippet/voiceName
func (t TextToSpeechVoiceTaggedSSMLSnippet) VoiceName() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceName"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechVoiceTaggedSSMLSnippet) SetVoiceName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceName:"), objc.String(value))
}

