// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSPronunciationUtils] class.
var (
	_TextToSpeechTTSPronunciationUtilsClass     TextToSpeechTTSPronunciationUtilsClass
	_TextToSpeechTTSPronunciationUtilsClassOnce sync.Once
)

func getTextToSpeechTTSPronunciationUtilsClass() TextToSpeechTTSPronunciationUtilsClass {
	_TextToSpeechTTSPronunciationUtilsClassOnce.Do(func() {
		_TextToSpeechTTSPronunciationUtilsClass = TextToSpeechTTSPronunciationUtilsClass{class: objc.GetClass("TextToSpeech.TTSPronunciationUtils")}
	})
	return _TextToSpeechTTSPronunciationUtilsClass
}

// GetTextToSpeechTTSPronunciationUtilsClass returns the class object for TextToSpeech.TTSPronunciationUtils.
func GetTextToSpeechTTSPronunciationUtilsClass() TextToSpeechTTSPronunciationUtilsClass {
	return getTextToSpeechTTSPronunciationUtilsClass()
}

type TextToSpeechTTSPronunciationUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSPronunciationUtilsClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSPronunciationUtilsClass) Alloc() TextToSpeechTTSPronunciationUtils {
	rv := objc.Send[TextToSpeechTTSPronunciationUtils](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSPronunciationUtils
type TextToSpeechTTSPronunciationUtils struct {
	objectivec.Object
}

// TextToSpeechTTSPronunciationUtilsFromID constructs a [TextToSpeechTTSPronunciationUtils] from an objc.ID.
func TextToSpeechTTSPronunciationUtilsFromID(id objc.ID) TextToSpeechTTSPronunciationUtils {
	return TextToSpeechTTSPronunciationUtils{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechTTSPronunciationUtils struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechTTSPronunciationUtils embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechTTSPronunciationUtils] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSPronunciationUtils
type ITextToSpeechTTSPronunciationUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSPronunciationUtils) Init() TextToSpeechTTSPronunciationUtils {
	rv := objc.Send[TextToSpeechTTSPronunciationUtils](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSPronunciationUtils) Autorelease() TextToSpeechTTSPronunciationUtils {
	rv := objc.Send[TextToSpeechTTSPronunciationUtils](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSPronunciationUtils creates a new TextToSpeechTTSPronunciationUtils instance.
func NewTextToSpeechTTSPronunciationUtils() TextToSpeechTTSPronunciationUtils {
	class := getTextToSpeechTTSPronunciationUtilsClass()
	rv := objc.Send[TextToSpeechTTSPronunciationUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}
