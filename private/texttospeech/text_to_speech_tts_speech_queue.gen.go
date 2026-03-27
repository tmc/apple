// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSpeechQueue] class.
var (
	_TextToSpeechTTSSpeechQueueClass     TextToSpeechTTSSpeechQueueClass
	_TextToSpeechTTSSpeechQueueClassOnce sync.Once
)

func getTextToSpeechTTSSpeechQueueClass() TextToSpeechTTSSpeechQueueClass {
	_TextToSpeechTTSSpeechQueueClassOnce.Do(func() {
		_TextToSpeechTTSSpeechQueueClass = TextToSpeechTTSSpeechQueueClass{class: objc.GetClass("TextToSpeech.TTSSpeechQueue")}
	})
	return _TextToSpeechTTSSpeechQueueClass
}

// GetTextToSpeechTTSSpeechQueueClass returns the class object for TextToSpeech.TTSSpeechQueue.
func GetTextToSpeechTTSSpeechQueueClass() TextToSpeechTTSSpeechQueueClass {
	return getTextToSpeechTTSSpeechQueueClass()
}

type TextToSpeechTTSSpeechQueueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSpeechQueueClass) Alloc() TextToSpeechTTSSpeechQueue {
	rv := objc.Send[TextToSpeechTTSSpeechQueue](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechQueue
type TextToSpeechTTSSpeechQueue struct {
	objectivec.Object
}

// TextToSpeechTTSSpeechQueueFromID constructs a [TextToSpeechTTSSpeechQueue] from an objc.ID.
func TextToSpeechTTSSpeechQueueFromID(id objc.ID) TextToSpeechTTSSpeechQueue {
	return TextToSpeechTTSSpeechQueue{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechTTSSpeechQueue implements ITextToSpeechTTSSpeechQueue.
var _ ITextToSpeechTTSSpeechQueue = TextToSpeechTTSSpeechQueue{}

// An interface definition for the [TextToSpeechTTSSpeechQueue] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechQueue
type ITextToSpeechTTSSpeechQueue interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSpeechQueue) Init() TextToSpeechTTSSpeechQueue {
	rv := objc.Send[TextToSpeechTTSSpeechQueue](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSpeechQueue) Autorelease() TextToSpeechTTSSpeechQueue {
	rv := objc.Send[TextToSpeechTTSSpeechQueue](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSpeechQueue creates a new TextToSpeechTTSSpeechQueue instance.
func NewTextToSpeechTTSSpeechQueue() TextToSpeechTTSSpeechQueue {
	class := getTextToSpeechTTSSpeechQueueClass()
	rv := objc.Send[TextToSpeechTTSSpeechQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

