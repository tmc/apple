// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSTaskRunner] class.
var (
	_TextToSpeechTTSTaskRunnerClass     TextToSpeechTTSTaskRunnerClass
	_TextToSpeechTTSTaskRunnerClassOnce sync.Once
)

func getTextToSpeechTTSTaskRunnerClass() TextToSpeechTTSTaskRunnerClass {
	_TextToSpeechTTSTaskRunnerClassOnce.Do(func() {
		_TextToSpeechTTSTaskRunnerClass = TextToSpeechTTSTaskRunnerClass{class: objc.GetClass("TextToSpeech.TTSTaskRunner")}
	})
	return _TextToSpeechTTSTaskRunnerClass
}

// GetTextToSpeechTTSTaskRunnerClass returns the class object for TextToSpeech.TTSTaskRunner.
func GetTextToSpeechTTSTaskRunnerClass() TextToSpeechTTSTaskRunnerClass {
	return getTextToSpeechTTSTaskRunnerClass()
}

type TextToSpeechTTSTaskRunnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSTaskRunnerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSTaskRunnerClass) Alloc() TextToSpeechTTSTaskRunner {
	rv := objc.SendIfResponds[TextToSpeechTTSTaskRunner](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechTTSTaskRunner struct {
	objectivec.Object
}

// TextToSpeechTTSTaskRunnerFromID constructs a [TextToSpeechTTSTaskRunner] from an objc.ID.
func TextToSpeechTTSTaskRunnerFromID(id objc.ID) TextToSpeechTTSTaskRunner {
	return TextToSpeechTTSTaskRunner{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechTTSTaskRunner implements ITextToSpeechTTSTaskRunner.
var _ ITextToSpeechTTSTaskRunner = TextToSpeechTTSTaskRunner{}

// An interface definition for the [TextToSpeechTTSTaskRunner] class.
type ITextToSpeechTTSTaskRunner interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSTaskRunner) Init() TextToSpeechTTSTaskRunner {
	rv := objc.SendIfResponds[TextToSpeechTTSTaskRunner](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSTaskRunner) Autorelease() TextToSpeechTTSTaskRunner {
	rv := objc.SendIfResponds[TextToSpeechTTSTaskRunner](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSTaskRunner creates a new TextToSpeechTTSTaskRunner instance.
func NewTextToSpeechTTSTaskRunner() TextToSpeechTTSTaskRunner {
	class := getTextToSpeechTTSTaskRunnerClass()
	rv := objc.SendIfResponds[TextToSpeechTTSTaskRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}
