// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSCoreAudioQueue] class.
var (
	_TextToSpeechTTSCoreAudioQueueClass     TextToSpeechTTSCoreAudioQueueClass
	_TextToSpeechTTSCoreAudioQueueClassOnce sync.Once
)

func getTextToSpeechTTSCoreAudioQueueClass() TextToSpeechTTSCoreAudioQueueClass {
	_TextToSpeechTTSCoreAudioQueueClassOnce.Do(func() {
		_TextToSpeechTTSCoreAudioQueueClass = TextToSpeechTTSCoreAudioQueueClass{class: objc.GetClass("TextToSpeech.TTSCoreAudioQueue")}
	})
	return _TextToSpeechTTSCoreAudioQueueClass
}

// GetTextToSpeechTTSCoreAudioQueueClass returns the class object for TextToSpeech.TTSCoreAudioQueue.
func GetTextToSpeechTTSCoreAudioQueueClass() TextToSpeechTTSCoreAudioQueueClass {
	return getTextToSpeechTTSCoreAudioQueueClass()
}

type TextToSpeechTTSCoreAudioQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSCoreAudioQueueClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSCoreAudioQueueClass) Alloc() TextToSpeechTTSCoreAudioQueue {
	rv := objc.SendIfResponds[TextToSpeechTTSCoreAudioQueue](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechTTSCoreAudioQueue struct {
	objectivec.Object
}

// TextToSpeechTTSCoreAudioQueueFromID constructs a [TextToSpeechTTSCoreAudioQueue] from an objc.ID.
func TextToSpeechTTSCoreAudioQueueFromID(id objc.ID) TextToSpeechTTSCoreAudioQueue {
	return TextToSpeechTTSCoreAudioQueue{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechTTSCoreAudioQueue embeds objectivec.Object because the parent type is
// unavailable, but ITextToSpeechTTSCoreAudioQueue embeds ISwiftNativeNSObject, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [TextToSpeechTTSCoreAudioQueue] class.
type ITextToSpeechTTSCoreAudioQueue interface {
	ISwiftNativeNSObject
}

// Init initializes the instance.
func (t TextToSpeechTTSCoreAudioQueue) Init() TextToSpeechTTSCoreAudioQueue {
	rv := objc.SendIfResponds[TextToSpeechTTSCoreAudioQueue](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSCoreAudioQueue) Autorelease() TextToSpeechTTSCoreAudioQueue {
	rv := objc.SendIfResponds[TextToSpeechTTSCoreAudioQueue](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSCoreAudioQueue creates a new TextToSpeechTTSCoreAudioQueue instance.
func NewTextToSpeechTTSCoreAudioQueue() TextToSpeechTTSCoreAudioQueue {
	class := getTextToSpeechTTSCoreAudioQueueClass()
	rv := objc.SendIfResponds[TextToSpeechTTSCoreAudioQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}
