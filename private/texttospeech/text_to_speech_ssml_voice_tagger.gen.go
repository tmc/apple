// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSSMLVoiceTagger] class.
var (
	_TextToSpeechSSMLVoiceTaggerClass     TextToSpeechSSMLVoiceTaggerClass
	_TextToSpeechSSMLVoiceTaggerClassOnce sync.Once
)

func getTextToSpeechSSMLVoiceTaggerClass() TextToSpeechSSMLVoiceTaggerClass {
	_TextToSpeechSSMLVoiceTaggerClassOnce.Do(func() {
		_TextToSpeechSSMLVoiceTaggerClass = TextToSpeechSSMLVoiceTaggerClass{class: objc.GetClass("TextToSpeech.SSMLVoiceTagger")}
	})
	return _TextToSpeechSSMLVoiceTaggerClass
}

// GetTextToSpeechSSMLVoiceTaggerClass returns the class object for TextToSpeech.SSMLVoiceTagger.
func GetTextToSpeechSSMLVoiceTaggerClass() TextToSpeechSSMLVoiceTaggerClass {
	return getTextToSpeechSSMLVoiceTaggerClass()
}

type TextToSpeechSSMLVoiceTaggerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSMLVoiceTaggerClass) Alloc() TextToSpeechSSMLVoiceTagger {
	rv := objc.Send[TextToSpeechSSMLVoiceTagger](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechSSMLVoiceTagger.TagSSML]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLVoiceTagger
type TextToSpeechSSMLVoiceTagger struct {
	objectivec.Object
}

// TextToSpeechSSMLVoiceTaggerFromID constructs a [TextToSpeechSSMLVoiceTagger] from an objc.ID.
func TextToSpeechSSMLVoiceTaggerFromID(id objc.ID) TextToSpeechSSMLVoiceTagger {
	return TextToSpeechSSMLVoiceTagger{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechSSMLVoiceTagger implements ITextToSpeechSSMLVoiceTagger.
var _ ITextToSpeechSSMLVoiceTagger = TextToSpeechSSMLVoiceTagger{}

// An interface definition for the [TextToSpeechSSMLVoiceTagger] class.
//
// # Methods
//
//   - [ITextToSpeechSSMLVoiceTagger.TagSSML]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLVoiceTagger
type ITextToSpeechSSMLVoiceTagger interface {
	objectivec.IObject

	// Topic: Methods

	TagSSML(ssml objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSSMLVoiceTagger) Init() TextToSpeechSSMLVoiceTagger {
	rv := objc.Send[TextToSpeechSSMLVoiceTagger](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSMLVoiceTagger) Autorelease() TextToSpeechSSMLVoiceTagger {
	rv := objc.Send[TextToSpeechSSMLVoiceTagger](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSMLVoiceTagger creates a new TextToSpeechSSMLVoiceTagger instance.
func NewTextToSpeechSSMLVoiceTagger() TextToSpeechSSMLVoiceTagger {
	class := getTextToSpeechSSMLVoiceTaggerClass()
	rv := objc.Send[TextToSpeechSSMLVoiceTagger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSMLVoiceTagger/tagSSML:
func (t TextToSpeechSSMLVoiceTagger) TagSSML(ssml objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tagSSML:"), ssml)
	return objectivec.Object{ID: rv}
}

