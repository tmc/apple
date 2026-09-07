// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechEventBatcher] class.
var (
	_TextToSpeechEventBatcherClass     TextToSpeechEventBatcherClass
	_TextToSpeechEventBatcherClassOnce sync.Once
)

func getTextToSpeechEventBatcherClass() TextToSpeechEventBatcherClass {
	_TextToSpeechEventBatcherClassOnce.Do(func() {
		_TextToSpeechEventBatcherClass = TextToSpeechEventBatcherClass{class: objc.GetClass("TextToSpeech.EventBatcher")}
	})
	return _TextToSpeechEventBatcherClass
}

// GetTextToSpeechEventBatcherClass returns the class object for TextToSpeech.EventBatcher.
func GetTextToSpeechEventBatcherClass() TextToSpeechEventBatcherClass {
	return getTextToSpeechEventBatcherClass()
}

type TextToSpeechEventBatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechEventBatcherClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechEventBatcherClass) Alloc() TextToSpeechEventBatcher {
	rv := objc.SendIfResponds[TextToSpeechEventBatcher](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechEventBatcher struct {
	objectivec.Object
}

// TextToSpeechEventBatcherFromID constructs a [TextToSpeechEventBatcher] from an objc.ID.
func TextToSpeechEventBatcherFromID(id objc.ID) TextToSpeechEventBatcher {
	return TextToSpeechEventBatcher{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechEventBatcher implements ITextToSpeechEventBatcher.
var _ ITextToSpeechEventBatcher = TextToSpeechEventBatcher{}

// An interface definition for the [TextToSpeechEventBatcher] class.
type ITextToSpeechEventBatcher interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechEventBatcher) Init() TextToSpeechEventBatcher {
	rv := objc.SendIfResponds[TextToSpeechEventBatcher](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechEventBatcher) Autorelease() TextToSpeechEventBatcher {
	rv := objc.SendIfResponds[TextToSpeechEventBatcher](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechEventBatcher creates a new TextToSpeechEventBatcher instance.
func NewTextToSpeechEventBatcher() TextToSpeechEventBatcher {
	class := getTextToSpeechEventBatcherClass()
	rv := objc.SendIfResponds[TextToSpeechEventBatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}
