// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAXVoiceLoader] class.
var (
	_TextToSpeechAXVoiceLoaderClass     TextToSpeechAXVoiceLoaderClass
	_TextToSpeechAXVoiceLoaderClassOnce sync.Once
)

func getTextToSpeechAXVoiceLoaderClass() TextToSpeechAXVoiceLoaderClass {
	_TextToSpeechAXVoiceLoaderClassOnce.Do(func() {
		_TextToSpeechAXVoiceLoaderClass = TextToSpeechAXVoiceLoaderClass{class: objc.GetClass("TextToSpeech.AXVoiceLoader")}
	})
	return _TextToSpeechAXVoiceLoaderClass
}

// GetTextToSpeechAXVoiceLoaderClass returns the class object for TextToSpeech.AXVoiceLoader.
func GetTextToSpeechAXVoiceLoaderClass() TextToSpeechAXVoiceLoaderClass {
	return getTextToSpeechAXVoiceLoaderClass()
}

type TextToSpeechAXVoiceLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAXVoiceLoaderClass) Alloc() TextToSpeechAXVoiceLoader {
	rv := objc.Send[TextToSpeechAXVoiceLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXVoiceLoader
type TextToSpeechAXVoiceLoader struct {
	objectivec.Object
}

// TextToSpeechAXVoiceLoaderFromID constructs a [TextToSpeechAXVoiceLoader] from an objc.ID.
func TextToSpeechAXVoiceLoaderFromID(id objc.ID) TextToSpeechAXVoiceLoader {
	return TextToSpeechAXVoiceLoader{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechAXVoiceLoader implements ITextToSpeechAXVoiceLoader.
var _ ITextToSpeechAXVoiceLoader = TextToSpeechAXVoiceLoader{}

// An interface definition for the [TextToSpeechAXVoiceLoader] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXVoiceLoader
type ITextToSpeechAXVoiceLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAXVoiceLoader) Init() TextToSpeechAXVoiceLoader {
	rv := objc.Send[TextToSpeechAXVoiceLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAXVoiceLoader) Autorelease() TextToSpeechAXVoiceLoader {
	rv := objc.Send[TextToSpeechAXVoiceLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAXVoiceLoader creates a new TextToSpeechAXVoiceLoader instance.
func NewTextToSpeechAXVoiceLoader() TextToSpeechAXVoiceLoader {
	class := getTextToSpeechAXVoiceLoaderClass()
	rv := objc.Send[TextToSpeechAXVoiceLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

