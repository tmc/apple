// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechEmojiLocaleDataCache] class.
var (
	_TextToSpeechEmojiLocaleDataCacheClass     TextToSpeechEmojiLocaleDataCacheClass
	_TextToSpeechEmojiLocaleDataCacheClassOnce sync.Once
)

func getTextToSpeechEmojiLocaleDataCacheClass() TextToSpeechEmojiLocaleDataCacheClass {
	_TextToSpeechEmojiLocaleDataCacheClassOnce.Do(func() {
		_TextToSpeechEmojiLocaleDataCacheClass = TextToSpeechEmojiLocaleDataCacheClass{class: objc.GetClass("TextToSpeech.EmojiLocaleDataCache")}
	})
	return _TextToSpeechEmojiLocaleDataCacheClass
}

// GetTextToSpeechEmojiLocaleDataCacheClass returns the class object for TextToSpeech.EmojiLocaleDataCache.
func GetTextToSpeechEmojiLocaleDataCacheClass() TextToSpeechEmojiLocaleDataCacheClass {
	return getTextToSpeechEmojiLocaleDataCacheClass()
}

type TextToSpeechEmojiLocaleDataCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechEmojiLocaleDataCacheClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechEmojiLocaleDataCacheClass) Alloc() TextToSpeechEmojiLocaleDataCache {
	rv := objc.SendIfResponds[TextToSpeechEmojiLocaleDataCache](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechEmojiLocaleDataCache struct {
	objectivec.Object
}

// TextToSpeechEmojiLocaleDataCacheFromID constructs a [TextToSpeechEmojiLocaleDataCache] from an objc.ID.
func TextToSpeechEmojiLocaleDataCacheFromID(id objc.ID) TextToSpeechEmojiLocaleDataCache {
	return TextToSpeechEmojiLocaleDataCache{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechEmojiLocaleDataCache implements ITextToSpeechEmojiLocaleDataCache.
var _ ITextToSpeechEmojiLocaleDataCache = TextToSpeechEmojiLocaleDataCache{}

// An interface definition for the [TextToSpeechEmojiLocaleDataCache] class.
type ITextToSpeechEmojiLocaleDataCache interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechEmojiLocaleDataCache) Init() TextToSpeechEmojiLocaleDataCache {
	rv := objc.SendIfResponds[TextToSpeechEmojiLocaleDataCache](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechEmojiLocaleDataCache) Autorelease() TextToSpeechEmojiLocaleDataCache {
	rv := objc.SendIfResponds[TextToSpeechEmojiLocaleDataCache](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechEmojiLocaleDataCache creates a new TextToSpeechEmojiLocaleDataCache instance.
func NewTextToSpeechEmojiLocaleDataCache() TextToSpeechEmojiLocaleDataCache {
	class := getTextToSpeechEmojiLocaleDataCacheClass()
	rv := objc.SendIfResponds[TextToSpeechEmojiLocaleDataCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}
