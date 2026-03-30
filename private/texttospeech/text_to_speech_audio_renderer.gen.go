// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAudioRenderer] class.
var (
	_TextToSpeechAudioRendererClass     TextToSpeechAudioRendererClass
	_TextToSpeechAudioRendererClassOnce sync.Once
)

func getTextToSpeechAudioRendererClass() TextToSpeechAudioRendererClass {
	_TextToSpeechAudioRendererClassOnce.Do(func() {
		_TextToSpeechAudioRendererClass = TextToSpeechAudioRendererClass{class: objc.GetClass("TextToSpeech.AudioRenderer")}
	})
	return _TextToSpeechAudioRendererClass
}

// GetTextToSpeechAudioRendererClass returns the class object for TextToSpeech.AudioRenderer.
func GetTextToSpeechAudioRendererClass() TextToSpeechAudioRendererClass {
	return getTextToSpeechAudioRendererClass()
}

type TextToSpeechAudioRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechAudioRendererClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAudioRendererClass) Alloc() TextToSpeechAudioRenderer {
	rv := objc.Send[TextToSpeechAudioRenderer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AudioRenderer
type TextToSpeechAudioRenderer struct {
	objectivec.Object
}

// TextToSpeechAudioRendererFromID constructs a [TextToSpeechAudioRenderer] from an objc.ID.
func TextToSpeechAudioRendererFromID(id objc.ID) TextToSpeechAudioRenderer {
	return TextToSpeechAudioRenderer{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechAudioRenderer struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechAudioRenderer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechAudioRenderer] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AudioRenderer
type ITextToSpeechAudioRenderer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAudioRenderer) Init() TextToSpeechAudioRenderer {
	rv := objc.Send[TextToSpeechAudioRenderer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAudioRenderer) Autorelease() TextToSpeechAudioRenderer {
	rv := objc.Send[TextToSpeechAudioRenderer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAudioRenderer creates a new TextToSpeechAudioRenderer instance.
func NewTextToSpeechAudioRenderer() TextToSpeechAudioRenderer {
	class := getTextToSpeechAudioRendererClass()
	rv := objc.Send[TextToSpeechAudioRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
