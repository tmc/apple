// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechStreamingTokenizer] class.
var (
	_TextToSpeechStreamingTokenizerClass     TextToSpeechStreamingTokenizerClass
	_TextToSpeechStreamingTokenizerClassOnce sync.Once
)

func getTextToSpeechStreamingTokenizerClass() TextToSpeechStreamingTokenizerClass {
	_TextToSpeechStreamingTokenizerClassOnce.Do(func() {
		_TextToSpeechStreamingTokenizerClass = TextToSpeechStreamingTokenizerClass{class: objc.GetClass("TextToSpeech.StreamingTokenizer")}
	})
	return _TextToSpeechStreamingTokenizerClass
}

// GetTextToSpeechStreamingTokenizerClass returns the class object for TextToSpeech.StreamingTokenizer.
func GetTextToSpeechStreamingTokenizerClass() TextToSpeechStreamingTokenizerClass {
	return getTextToSpeechStreamingTokenizerClass()
}

type TextToSpeechStreamingTokenizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechStreamingTokenizerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechStreamingTokenizerClass) Alloc() TextToSpeechStreamingTokenizer {
	rv := objc.Send[TextToSpeechStreamingTokenizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.StreamingTokenizer
type TextToSpeechStreamingTokenizer struct {
	objectivec.Object
}

// TextToSpeechStreamingTokenizerFromID constructs a [TextToSpeechStreamingTokenizer] from an objc.ID.
func TextToSpeechStreamingTokenizerFromID(id objc.ID) TextToSpeechStreamingTokenizer {
	return TextToSpeechStreamingTokenizer{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechStreamingTokenizer struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechStreamingTokenizer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechStreamingTokenizer] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.StreamingTokenizer
type ITextToSpeechStreamingTokenizer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechStreamingTokenizer) Init() TextToSpeechStreamingTokenizer {
	rv := objc.Send[TextToSpeechStreamingTokenizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechStreamingTokenizer) Autorelease() TextToSpeechStreamingTokenizer {
	rv := objc.Send[TextToSpeechStreamingTokenizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechStreamingTokenizer creates a new TextToSpeechStreamingTokenizer instance.
func NewTextToSpeechStreamingTokenizer() TextToSpeechStreamingTokenizer {
	class := getTextToSpeechStreamingTokenizerClass()
	rv := objc.Send[TextToSpeechStreamingTokenizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

