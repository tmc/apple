// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAUParamWrapper] class.
var (
	_TextToSpeechAUParamWrapperClass     TextToSpeechAUParamWrapperClass
	_TextToSpeechAUParamWrapperClassOnce sync.Once
)

func getTextToSpeechAUParamWrapperClass() TextToSpeechAUParamWrapperClass {
	_TextToSpeechAUParamWrapperClassOnce.Do(func() {
		_TextToSpeechAUParamWrapperClass = TextToSpeechAUParamWrapperClass{class: objc.GetClass("TextToSpeech.AUParamWrapper")}
	})
	return _TextToSpeechAUParamWrapperClass
}

// GetTextToSpeechAUParamWrapperClass returns the class object for TextToSpeech.AUParamWrapper.
func GetTextToSpeechAUParamWrapperClass() TextToSpeechAUParamWrapperClass {
	return getTextToSpeechAUParamWrapperClass()
}

type TextToSpeechAUParamWrapperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAUParamWrapperClass) Alloc() TextToSpeechAUParamWrapper {
	rv := objc.Send[TextToSpeechAUParamWrapper](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AUParamWrapper
type TextToSpeechAUParamWrapper struct {
	objectivec.Object
}

// TextToSpeechAUParamWrapperFromID constructs a [TextToSpeechAUParamWrapper] from an objc.ID.
func TextToSpeechAUParamWrapperFromID(id objc.ID) TextToSpeechAUParamWrapper {
	return TextToSpeechAUParamWrapper{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechAUParamWrapper implements ITextToSpeechAUParamWrapper.
var _ ITextToSpeechAUParamWrapper = TextToSpeechAUParamWrapper{}

// An interface definition for the [TextToSpeechAUParamWrapper] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AUParamWrapper
type ITextToSpeechAUParamWrapper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAUParamWrapper) Init() TextToSpeechAUParamWrapper {
	rv := objc.Send[TextToSpeechAUParamWrapper](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAUParamWrapper) Autorelease() TextToSpeechAUParamWrapper {
	rv := objc.Send[TextToSpeechAUParamWrapper](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAUParamWrapper creates a new TextToSpeechAUParamWrapper instance.
func NewTextToSpeechAUParamWrapper() TextToSpeechAUParamWrapper {
	class := getTextToSpeechAUParamWrapperClass()
	rv := objc.Send[TextToSpeechAUParamWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

