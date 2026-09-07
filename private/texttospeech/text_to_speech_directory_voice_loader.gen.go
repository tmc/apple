// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechDirectoryVoiceLoader] class.
var (
	_TextToSpeechDirectoryVoiceLoaderClass     TextToSpeechDirectoryVoiceLoaderClass
	_TextToSpeechDirectoryVoiceLoaderClassOnce sync.Once
)

func getTextToSpeechDirectoryVoiceLoaderClass() TextToSpeechDirectoryVoiceLoaderClass {
	_TextToSpeechDirectoryVoiceLoaderClassOnce.Do(func() {
		_TextToSpeechDirectoryVoiceLoaderClass = TextToSpeechDirectoryVoiceLoaderClass{class: objc.GetClass("TextToSpeech.DirectoryVoiceLoader")}
	})
	return _TextToSpeechDirectoryVoiceLoaderClass
}

// GetTextToSpeechDirectoryVoiceLoaderClass returns the class object for TextToSpeech.DirectoryVoiceLoader.
func GetTextToSpeechDirectoryVoiceLoaderClass() TextToSpeechDirectoryVoiceLoaderClass {
	return getTextToSpeechDirectoryVoiceLoaderClass()
}

type TextToSpeechDirectoryVoiceLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechDirectoryVoiceLoaderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechDirectoryVoiceLoaderClass) Alloc() TextToSpeechDirectoryVoiceLoader {
	rv := objc.SendIfResponds[TextToSpeechDirectoryVoiceLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechDirectoryVoiceLoader struct {
	objectivec.Object
}

// TextToSpeechDirectoryVoiceLoaderFromID constructs a [TextToSpeechDirectoryVoiceLoader] from an objc.ID.
func TextToSpeechDirectoryVoiceLoaderFromID(id objc.ID) TextToSpeechDirectoryVoiceLoader {
	return TextToSpeechDirectoryVoiceLoader{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechDirectoryVoiceLoader implements ITextToSpeechDirectoryVoiceLoader.
var _ ITextToSpeechDirectoryVoiceLoader = TextToSpeechDirectoryVoiceLoader{}

// An interface definition for the [TextToSpeechDirectoryVoiceLoader] class.
type ITextToSpeechDirectoryVoiceLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechDirectoryVoiceLoader) Init() TextToSpeechDirectoryVoiceLoader {
	rv := objc.SendIfResponds[TextToSpeechDirectoryVoiceLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechDirectoryVoiceLoader) Autorelease() TextToSpeechDirectoryVoiceLoader {
	rv := objc.SendIfResponds[TextToSpeechDirectoryVoiceLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechDirectoryVoiceLoader creates a new TextToSpeechDirectoryVoiceLoader instance.
func NewTextToSpeechDirectoryVoiceLoader() TextToSpeechDirectoryVoiceLoader {
	class := getTextToSpeechDirectoryVoiceLoaderClass()
	rv := objc.SendIfResponds[TextToSpeechDirectoryVoiceLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}
