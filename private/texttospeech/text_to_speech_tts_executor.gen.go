// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSExecutor] class.
var (
	_TextToSpeechTTSExecutorClass     TextToSpeechTTSExecutorClass
	_TextToSpeechTTSExecutorClassOnce sync.Once
)

func getTextToSpeechTTSExecutorClass() TextToSpeechTTSExecutorClass {
	_TextToSpeechTTSExecutorClassOnce.Do(func() {
		_TextToSpeechTTSExecutorClass = TextToSpeechTTSExecutorClass{class: objc.GetClass("TextToSpeech.TTSExecutor")}
	})
	return _TextToSpeechTTSExecutorClass
}

// GetTextToSpeechTTSExecutorClass returns the class object for TextToSpeech.TTSExecutor.
func GetTextToSpeechTTSExecutorClass() TextToSpeechTTSExecutorClass {
	return getTextToSpeechTTSExecutorClass()
}

type TextToSpeechTTSExecutorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSExecutorClass) Alloc() TextToSpeechTTSExecutor {
	rv := objc.Send[TextToSpeechTTSExecutor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSExecutor
type TextToSpeechTTSExecutor struct {
	objectivec.Object
}

// TextToSpeechTTSExecutorFromID constructs a [TextToSpeechTTSExecutor] from an objc.ID.
func TextToSpeechTTSExecutorFromID(id objc.ID) TextToSpeechTTSExecutor {
	return TextToSpeechTTSExecutor{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechTTSExecutor implements ITextToSpeechTTSExecutor.
var _ ITextToSpeechTTSExecutor = TextToSpeechTTSExecutor{}

// An interface definition for the [TextToSpeechTTSExecutor] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSExecutor
type ITextToSpeechTTSExecutor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSExecutor) Init() TextToSpeechTTSExecutor {
	rv := objc.Send[TextToSpeechTTSExecutor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSExecutor) Autorelease() TextToSpeechTTSExecutor {
	rv := objc.Send[TextToSpeechTTSExecutor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSExecutor creates a new TextToSpeechTTSExecutor instance.
func NewTextToSpeechTTSExecutor() TextToSpeechTTSExecutor {
	class := getTextToSpeechTTSExecutorClass()
	rv := objc.Send[TextToSpeechTTSExecutor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

