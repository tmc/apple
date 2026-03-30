// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceSmuggler] class.
var (
	_TextToSpeechVoiceSmugglerClass     TextToSpeechVoiceSmugglerClass
	_TextToSpeechVoiceSmugglerClassOnce sync.Once
)

func getTextToSpeechVoiceSmugglerClass() TextToSpeechVoiceSmugglerClass {
	_TextToSpeechVoiceSmugglerClassOnce.Do(func() {
		_TextToSpeechVoiceSmugglerClass = TextToSpeechVoiceSmugglerClass{class: objc.GetClass("TextToSpeech.VoiceSmuggler")}
	})
	return _TextToSpeechVoiceSmugglerClass
}

// GetTextToSpeechVoiceSmugglerClass returns the class object for TextToSpeech.VoiceSmuggler.
func GetTextToSpeechVoiceSmugglerClass() TextToSpeechVoiceSmugglerClass {
	return getTextToSpeechVoiceSmugglerClass()
}

type TextToSpeechVoiceSmugglerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceSmugglerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceSmugglerClass) Alloc() TextToSpeechVoiceSmuggler {
	rv := objc.Send[TextToSpeechVoiceSmuggler](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceSmuggler
type TextToSpeechVoiceSmuggler struct {
	objectivec.Object
}

// TextToSpeechVoiceSmugglerFromID constructs a [TextToSpeechVoiceSmuggler] from an objc.ID.
func TextToSpeechVoiceSmugglerFromID(id objc.ID) TextToSpeechVoiceSmuggler {
	return TextToSpeechVoiceSmuggler{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechVoiceSmuggler implements ITextToSpeechVoiceSmuggler.
var _ ITextToSpeechVoiceSmuggler = TextToSpeechVoiceSmuggler{}

// An interface definition for the [TextToSpeechVoiceSmuggler] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceSmuggler
type ITextToSpeechVoiceSmuggler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechVoiceSmuggler) Init() TextToSpeechVoiceSmuggler {
	rv := objc.Send[TextToSpeechVoiceSmuggler](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceSmuggler) Autorelease() TextToSpeechVoiceSmuggler {
	rv := objc.Send[TextToSpeechVoiceSmuggler](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceSmuggler creates a new TextToSpeechVoiceSmuggler instance.
func NewTextToSpeechVoiceSmuggler() TextToSpeechVoiceSmuggler {
	class := getTextToSpeechVoiceSmugglerClass()
	rv := objc.Send[TextToSpeechVoiceSmuggler](objc.ID(class.class), objc.Sel("new"))
	return rv
}
