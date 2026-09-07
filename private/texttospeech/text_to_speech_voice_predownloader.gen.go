// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoicePredownloader] class.
var (
	_TextToSpeechVoicePredownloaderClass     TextToSpeechVoicePredownloaderClass
	_TextToSpeechVoicePredownloaderClassOnce sync.Once
)

func getTextToSpeechVoicePredownloaderClass() TextToSpeechVoicePredownloaderClass {
	_TextToSpeechVoicePredownloaderClassOnce.Do(func() {
		_TextToSpeechVoicePredownloaderClass = TextToSpeechVoicePredownloaderClass{class: objc.GetClass("TextToSpeech.VoicePredownloader")}
	})
	return _TextToSpeechVoicePredownloaderClass
}

// GetTextToSpeechVoicePredownloaderClass returns the class object for TextToSpeech.VoicePredownloader.
func GetTextToSpeechVoicePredownloaderClass() TextToSpeechVoicePredownloaderClass {
	return getTextToSpeechVoicePredownloaderClass()
}

type TextToSpeechVoicePredownloaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoicePredownloaderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoicePredownloaderClass) Alloc() TextToSpeechVoicePredownloader {
	rv := objc.SendIfResponds[TextToSpeechVoicePredownloader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechVoicePredownloader struct {
	objectivec.Object
}

// TextToSpeechVoicePredownloaderFromID constructs a [TextToSpeechVoicePredownloader] from an objc.ID.
func TextToSpeechVoicePredownloaderFromID(id objc.ID) TextToSpeechVoicePredownloader {
	return TextToSpeechVoicePredownloader{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechVoicePredownloader implements ITextToSpeechVoicePredownloader.
var _ ITextToSpeechVoicePredownloader = TextToSpeechVoicePredownloader{}

// An interface definition for the [TextToSpeechVoicePredownloader] class.
type ITextToSpeechVoicePredownloader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoicePredownloader) Init() TextToSpeechVoicePredownloader {
	rv := objc.SendIfResponds[TextToSpeechVoicePredownloader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoicePredownloader) Autorelease() TextToSpeechVoicePredownloader {
	rv := objc.SendIfResponds[TextToSpeechVoicePredownloader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoicePredownloader creates a new TextToSpeechVoicePredownloader instance.
func NewTextToSpeechVoicePredownloader() TextToSpeechVoicePredownloader {
	class := getTextToSpeechVoicePredownloaderClass()
	rv := objc.SendIfResponds[TextToSpeechVoicePredownloader](objc.ID(class.class), objc.Sel("new"))
	return rv
}
