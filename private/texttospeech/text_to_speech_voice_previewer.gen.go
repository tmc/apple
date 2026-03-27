// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoicePreviewer] class.
var (
	_TextToSpeechVoicePreviewerClass     TextToSpeechVoicePreviewerClass
	_TextToSpeechVoicePreviewerClassOnce sync.Once
)

func getTextToSpeechVoicePreviewerClass() TextToSpeechVoicePreviewerClass {
	_TextToSpeechVoicePreviewerClassOnce.Do(func() {
		_TextToSpeechVoicePreviewerClass = TextToSpeechVoicePreviewerClass{class: objc.GetClass("TextToSpeech.VoicePreviewer")}
	})
	return _TextToSpeechVoicePreviewerClass
}

// GetTextToSpeechVoicePreviewerClass returns the class object for TextToSpeech.VoicePreviewer.
func GetTextToSpeechVoicePreviewerClass() TextToSpeechVoicePreviewerClass {
	return getTextToSpeechVoicePreviewerClass()
}

type TextToSpeechVoicePreviewerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoicePreviewerClass) Alloc() TextToSpeechVoicePreviewer {
	rv := objc.Send[TextToSpeechVoicePreviewer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechVoicePreviewer.AudioPlayerDidFinishPlayingSuccessfully]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoicePreviewer
type TextToSpeechVoicePreviewer struct {
	SwiftNativeNSObject
}

// TextToSpeechVoicePreviewerFromID constructs a [TextToSpeechVoicePreviewer] from an objc.ID.
func TextToSpeechVoicePreviewerFromID(id objc.ID) TextToSpeechVoicePreviewer {
	return TextToSpeechVoicePreviewer{SwiftNativeNSObject: SwiftNativeNSObjectFromID(id)}
}
// Ensure TextToSpeechVoicePreviewer implements ITextToSpeechVoicePreviewer.
var _ ITextToSpeechVoicePreviewer = TextToSpeechVoicePreviewer{}

// An interface definition for the [TextToSpeechVoicePreviewer] class.
//
// # Methods
//
//   - [ITextToSpeechVoicePreviewer.AudioPlayerDidFinishPlayingSuccessfully]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoicePreviewer
type ITextToSpeechVoicePreviewer interface {
	ISwiftNativeNSObject

	// Topic: Methods

	AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool)
}

// Init initializes the instance.
func (t TextToSpeechVoicePreviewer) Init() TextToSpeechVoicePreviewer {
	rv := objc.Send[TextToSpeechVoicePreviewer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoicePreviewer) Autorelease() TextToSpeechVoicePreviewer {
	rv := objc.Send[TextToSpeechVoicePreviewer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoicePreviewer creates a new TextToSpeechVoicePreviewer instance.
func NewTextToSpeechVoicePreviewer() TextToSpeechVoicePreviewer {
	class := getTextToSpeechVoicePreviewerClass()
	rv := objc.Send[TextToSpeechVoicePreviewer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoicePreviewer/audioPlayerDidFinishPlaying:successfully:
func (t TextToSpeechVoicePreviewer) AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("audioPlayerDidFinishPlaying:successfully:"), playing, successfully)
}

