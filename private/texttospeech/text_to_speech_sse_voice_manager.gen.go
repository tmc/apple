// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSSEVoiceManager] class.
var (
	_TextToSpeechSSEVoiceManagerClass     TextToSpeechSSEVoiceManagerClass
	_TextToSpeechSSEVoiceManagerClassOnce sync.Once
)

func getTextToSpeechSSEVoiceManagerClass() TextToSpeechSSEVoiceManagerClass {
	_TextToSpeechSSEVoiceManagerClassOnce.Do(func() {
		_TextToSpeechSSEVoiceManagerClass = TextToSpeechSSEVoiceManagerClass{class: objc.GetClass("TextToSpeech.SSEVoiceManager")}
	})
	return _TextToSpeechSSEVoiceManagerClass
}

// GetTextToSpeechSSEVoiceManagerClass returns the class object for TextToSpeech.SSEVoiceManager.
func GetTextToSpeechSSEVoiceManagerClass() TextToSpeechSSEVoiceManagerClass {
	return getTextToSpeechSSEVoiceManagerClass()
}

type TextToSpeechSSEVoiceManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSEVoiceManagerClass) Alloc() TextToSpeechSSEVoiceManager {
	rv := objc.Send[TextToSpeechSSEVoiceManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSEVoiceManager
type TextToSpeechSSEVoiceManager struct {
	objectivec.Object
}

// TextToSpeechSSEVoiceManagerFromID constructs a [TextToSpeechSSEVoiceManager] from an objc.ID.
func TextToSpeechSSEVoiceManagerFromID(id objc.ID) TextToSpeechSSEVoiceManager {
	return TextToSpeechSSEVoiceManager{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechSSEVoiceManager implements ITextToSpeechSSEVoiceManager.
var _ ITextToSpeechSSEVoiceManager = TextToSpeechSSEVoiceManager{}

// An interface definition for the [TextToSpeechSSEVoiceManager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SSEVoiceManager
type ITextToSpeechSSEVoiceManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSSEVoiceManager) Init() TextToSpeechSSEVoiceManager {
	rv := objc.Send[TextToSpeechSSEVoiceManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSEVoiceManager) Autorelease() TextToSpeechSSEVoiceManager {
	rv := objc.Send[TextToSpeechSSEVoiceManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSEVoiceManager creates a new TextToSpeechSSEVoiceManager instance.
func NewTextToSpeechSSEVoiceManager() TextToSpeechSSEVoiceManager {
	class := getTextToSpeechSSEVoiceManagerClass()
	rv := objc.Send[TextToSpeechSSEVoiceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

