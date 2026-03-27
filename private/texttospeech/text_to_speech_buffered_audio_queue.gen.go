// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechBufferedAudioQueue] class.
var (
	_TextToSpeechBufferedAudioQueueClass     TextToSpeechBufferedAudioQueueClass
	_TextToSpeechBufferedAudioQueueClassOnce sync.Once
)

func getTextToSpeechBufferedAudioQueueClass() TextToSpeechBufferedAudioQueueClass {
	_TextToSpeechBufferedAudioQueueClassOnce.Do(func() {
		_TextToSpeechBufferedAudioQueueClass = TextToSpeechBufferedAudioQueueClass{class: objc.GetClass("TextToSpeech.BufferedAudioQueue")}
	})
	return _TextToSpeechBufferedAudioQueueClass
}

// GetTextToSpeechBufferedAudioQueueClass returns the class object for TextToSpeech.BufferedAudioQueue.
func GetTextToSpeechBufferedAudioQueueClass() TextToSpeechBufferedAudioQueueClass {
	return getTextToSpeechBufferedAudioQueueClass()
}

type TextToSpeechBufferedAudioQueueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechBufferedAudioQueueClass) Alloc() TextToSpeechBufferedAudioQueue {
	rv := objc.Send[TextToSpeechBufferedAudioQueue](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.BufferedAudioQueue
type TextToSpeechBufferedAudioQueue struct {
	objectivec.Object
}

// TextToSpeechBufferedAudioQueueFromID constructs a [TextToSpeechBufferedAudioQueue] from an objc.ID.
func TextToSpeechBufferedAudioQueueFromID(id objc.ID) TextToSpeechBufferedAudioQueue {
	return TextToSpeechBufferedAudioQueue{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechBufferedAudioQueue implements ITextToSpeechBufferedAudioQueue.
var _ ITextToSpeechBufferedAudioQueue = TextToSpeechBufferedAudioQueue{}

// An interface definition for the [TextToSpeechBufferedAudioQueue] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.BufferedAudioQueue
type ITextToSpeechBufferedAudioQueue interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechBufferedAudioQueue) Init() TextToSpeechBufferedAudioQueue {
	rv := objc.Send[TextToSpeechBufferedAudioQueue](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechBufferedAudioQueue) Autorelease() TextToSpeechBufferedAudioQueue {
	rv := objc.Send[TextToSpeechBufferedAudioQueue](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechBufferedAudioQueue creates a new TextToSpeechBufferedAudioQueue instance.
func NewTextToSpeechBufferedAudioQueue() TextToSpeechBufferedAudioQueue {
	class := getTextToSpeechBufferedAudioQueueClass()
	rv := objc.Send[TextToSpeechBufferedAudioQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

