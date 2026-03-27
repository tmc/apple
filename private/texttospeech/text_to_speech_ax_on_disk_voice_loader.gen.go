// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAXOnDiskVoiceLoader] class.
var (
	_TextToSpeechAXOnDiskVoiceLoaderClass     TextToSpeechAXOnDiskVoiceLoaderClass
	_TextToSpeechAXOnDiskVoiceLoaderClassOnce sync.Once
)

func getTextToSpeechAXOnDiskVoiceLoaderClass() TextToSpeechAXOnDiskVoiceLoaderClass {
	_TextToSpeechAXOnDiskVoiceLoaderClassOnce.Do(func() {
		_TextToSpeechAXOnDiskVoiceLoaderClass = TextToSpeechAXOnDiskVoiceLoaderClass{class: objc.GetClass("TextToSpeech.AXOnDiskVoiceLoader")}
	})
	return _TextToSpeechAXOnDiskVoiceLoaderClass
}

// GetTextToSpeechAXOnDiskVoiceLoaderClass returns the class object for TextToSpeech.AXOnDiskVoiceLoader.
func GetTextToSpeechAXOnDiskVoiceLoaderClass() TextToSpeechAXOnDiskVoiceLoaderClass {
	return getTextToSpeechAXOnDiskVoiceLoaderClass()
}

type TextToSpeechAXOnDiskVoiceLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAXOnDiskVoiceLoaderClass) Alloc() TextToSpeechAXOnDiskVoiceLoader {
	rv := objc.Send[TextToSpeechAXOnDiskVoiceLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXOnDiskVoiceLoader
type TextToSpeechAXOnDiskVoiceLoader struct {
	objectivec.Object
}

// TextToSpeechAXOnDiskVoiceLoaderFromID constructs a [TextToSpeechAXOnDiskVoiceLoader] from an objc.ID.
func TextToSpeechAXOnDiskVoiceLoaderFromID(id objc.ID) TextToSpeechAXOnDiskVoiceLoader {
	return TextToSpeechAXOnDiskVoiceLoader{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechAXOnDiskVoiceLoader implements ITextToSpeechAXOnDiskVoiceLoader.
var _ ITextToSpeechAXOnDiskVoiceLoader = TextToSpeechAXOnDiskVoiceLoader{}

// An interface definition for the [TextToSpeechAXOnDiskVoiceLoader] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AXOnDiskVoiceLoader
type ITextToSpeechAXOnDiskVoiceLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAXOnDiskVoiceLoader) Init() TextToSpeechAXOnDiskVoiceLoader {
	rv := objc.Send[TextToSpeechAXOnDiskVoiceLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAXOnDiskVoiceLoader) Autorelease() TextToSpeechAXOnDiskVoiceLoader {
	rv := objc.Send[TextToSpeechAXOnDiskVoiceLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAXOnDiskVoiceLoader creates a new TextToSpeechAXOnDiskVoiceLoader instance.
func NewTextToSpeechAXOnDiskVoiceLoader() TextToSpeechAXOnDiskVoiceLoader {
	class := getTextToSpeechAXOnDiskVoiceLoaderClass()
	rv := objc.Send[TextToSpeechAXOnDiskVoiceLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

