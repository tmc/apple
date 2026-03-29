// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSiriVoiceLoader] class.
var (
	_TextToSpeechSiriVoiceLoaderClass     TextToSpeechSiriVoiceLoaderClass
	_TextToSpeechSiriVoiceLoaderClassOnce sync.Once
)

func getTextToSpeechSiriVoiceLoaderClass() TextToSpeechSiriVoiceLoaderClass {
	_TextToSpeechSiriVoiceLoaderClassOnce.Do(func() {
		_TextToSpeechSiriVoiceLoaderClass = TextToSpeechSiriVoiceLoaderClass{class: objc.GetClass("TextToSpeech.SiriVoiceLoader")}
	})
	return _TextToSpeechSiriVoiceLoaderClass
}

// GetTextToSpeechSiriVoiceLoaderClass returns the class object for TextToSpeech.SiriVoiceLoader.
func GetTextToSpeechSiriVoiceLoaderClass() TextToSpeechSiriVoiceLoaderClass {
	return getTextToSpeechSiriVoiceLoaderClass()
}

type TextToSpeechSiriVoiceLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSiriVoiceLoaderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSiriVoiceLoaderClass) Alloc() TextToSpeechSiriVoiceLoader {
	rv := objc.Send[TextToSpeechSiriVoiceLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SiriVoiceLoader
type TextToSpeechSiriVoiceLoader struct {
	objectivec.Object
}

// TextToSpeechSiriVoiceLoaderFromID constructs a [TextToSpeechSiriVoiceLoader] from an objc.ID.
func TextToSpeechSiriVoiceLoaderFromID(id objc.ID) TextToSpeechSiriVoiceLoader {
	return TextToSpeechSiriVoiceLoader{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechSiriVoiceLoader struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechSiriVoiceLoader embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechSiriVoiceLoader] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SiriVoiceLoader
type ITextToSpeechSiriVoiceLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSiriVoiceLoader) Init() TextToSpeechSiriVoiceLoader {
	rv := objc.Send[TextToSpeechSiriVoiceLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSiriVoiceLoader) Autorelease() TextToSpeechSiriVoiceLoader {
	rv := objc.Send[TextToSpeechSiriVoiceLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSiriVoiceLoader creates a new TextToSpeechSiriVoiceLoader instance.
func NewTextToSpeechSiriVoiceLoader() TextToSpeechSiriVoiceLoader {
	class := getTextToSpeechSiriVoiceLoaderClass()
	rv := objc.Send[TextToSpeechSiriVoiceLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

