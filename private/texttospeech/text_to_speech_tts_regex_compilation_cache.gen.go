// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSRegexCompilationCache] class.
var (
	_TextToSpeechTTSRegexCompilationCacheClass     TextToSpeechTTSRegexCompilationCacheClass
	_TextToSpeechTTSRegexCompilationCacheClassOnce sync.Once
)

func getTextToSpeechTTSRegexCompilationCacheClass() TextToSpeechTTSRegexCompilationCacheClass {
	_TextToSpeechTTSRegexCompilationCacheClassOnce.Do(func() {
		_TextToSpeechTTSRegexCompilationCacheClass = TextToSpeechTTSRegexCompilationCacheClass{class: objc.GetClass("TextToSpeech.TTSRegexCompilationCache")}
	})
	return _TextToSpeechTTSRegexCompilationCacheClass
}

// GetTextToSpeechTTSRegexCompilationCacheClass returns the class object for TextToSpeech.TTSRegexCompilationCache.
func GetTextToSpeechTTSRegexCompilationCacheClass() TextToSpeechTTSRegexCompilationCacheClass {
	return getTextToSpeechTTSRegexCompilationCacheClass()
}

type TextToSpeechTTSRegexCompilationCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSRegexCompilationCacheClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSRegexCompilationCacheClass) Alloc() TextToSpeechTTSRegexCompilationCache {
	rv := objc.SendIfResponds[TextToSpeechTTSRegexCompilationCache](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechTTSRegexCompilationCache struct {
	objectivec.Object
}

// TextToSpeechTTSRegexCompilationCacheFromID constructs a [TextToSpeechTTSRegexCompilationCache] from an objc.ID.
func TextToSpeechTTSRegexCompilationCacheFromID(id objc.ID) TextToSpeechTTSRegexCompilationCache {
	return TextToSpeechTTSRegexCompilationCache{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechTTSRegexCompilationCache implements ITextToSpeechTTSRegexCompilationCache.
var _ ITextToSpeechTTSRegexCompilationCache = TextToSpeechTTSRegexCompilationCache{}

// An interface definition for the [TextToSpeechTTSRegexCompilationCache] class.
type ITextToSpeechTTSRegexCompilationCache interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSRegexCompilationCache) Init() TextToSpeechTTSRegexCompilationCache {
	rv := objc.SendIfResponds[TextToSpeechTTSRegexCompilationCache](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSRegexCompilationCache) Autorelease() TextToSpeechTTSRegexCompilationCache {
	rv := objc.SendIfResponds[TextToSpeechTTSRegexCompilationCache](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSRegexCompilationCache creates a new TextToSpeechTTSRegexCompilationCache instance.
func NewTextToSpeechTTSRegexCompilationCache() TextToSpeechTTSRegexCompilationCache {
	class := getTextToSpeechTTSRegexCompilationCacheClass()
	rv := objc.SendIfResponds[TextToSpeechTTSRegexCompilationCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}
