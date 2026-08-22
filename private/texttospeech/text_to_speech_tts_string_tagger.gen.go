// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSStringTagger] class.
var (
	_TextToSpeechTTSStringTaggerClass     TextToSpeechTTSStringTaggerClass
	_TextToSpeechTTSStringTaggerClassOnce sync.Once
)

func getTextToSpeechTTSStringTaggerClass() TextToSpeechTTSStringTaggerClass {
	_TextToSpeechTTSStringTaggerClassOnce.Do(func() {
		_TextToSpeechTTSStringTaggerClass = TextToSpeechTTSStringTaggerClass{class: objc.GetClass("TextToSpeech.TTSStringTagger")}
	})
	return _TextToSpeechTTSStringTaggerClass
}

// GetTextToSpeechTTSStringTaggerClass returns the class object for TextToSpeech.TTSStringTagger.
func GetTextToSpeechTTSStringTaggerClass() TextToSpeechTTSStringTaggerClass {
	return getTextToSpeechTTSStringTaggerClass()
}

type TextToSpeechTTSStringTaggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSStringTaggerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSStringTaggerClass) Alloc() TextToSpeechTTSStringTagger {
	rv := objc.SendIfResponds[TextToSpeechTTSStringTagger](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechTTSStringTagger struct {
	objectivec.Object
}

// TextToSpeechTTSStringTaggerFromID constructs a [TextToSpeechTTSStringTagger] from an objc.ID.
func TextToSpeechTTSStringTaggerFromID(id objc.ID) TextToSpeechTTSStringTagger {
	return TextToSpeechTTSStringTagger{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechTTSStringTagger implements ITextToSpeechTTSStringTagger.
var _ ITextToSpeechTTSStringTagger = TextToSpeechTTSStringTagger{}

// An interface definition for the [TextToSpeechTTSStringTagger] class.
type ITextToSpeechTTSStringTagger interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSStringTagger) Init() TextToSpeechTTSStringTagger {
	rv := objc.SendIfResponds[TextToSpeechTTSStringTagger](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSStringTagger) Autorelease() TextToSpeechTTSStringTagger {
	rv := objc.SendIfResponds[TextToSpeechTTSStringTagger](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSStringTagger creates a new TextToSpeechTTSStringTagger instance.
func NewTextToSpeechTTSStringTagger() TextToSpeechTTSStringTagger {
	class := getTextToSpeechTTSStringTaggerClass()
	rv := objc.SendIfResponds[TextToSpeechTTSStringTagger](objc.ID(class.class), objc.Sel("new"))
	return rv
}
