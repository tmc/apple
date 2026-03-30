// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechAudioUnitCache] class.
var (
	_TextToSpeechAudioUnitCacheClass     TextToSpeechAudioUnitCacheClass
	_TextToSpeechAudioUnitCacheClassOnce sync.Once
)

func getTextToSpeechAudioUnitCacheClass() TextToSpeechAudioUnitCacheClass {
	_TextToSpeechAudioUnitCacheClassOnce.Do(func() {
		_TextToSpeechAudioUnitCacheClass = TextToSpeechAudioUnitCacheClass{class: objc.GetClass("TextToSpeech.AudioUnitCache")}
	})
	return _TextToSpeechAudioUnitCacheClass
}

// GetTextToSpeechAudioUnitCacheClass returns the class object for TextToSpeech.AudioUnitCache.
func GetTextToSpeechAudioUnitCacheClass() TextToSpeechAudioUnitCacheClass {
	return getTextToSpeechAudioUnitCacheClass()
}

type TextToSpeechAudioUnitCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechAudioUnitCacheClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechAudioUnitCacheClass) Alloc() TextToSpeechAudioUnitCache {
	rv := objc.Send[TextToSpeechAudioUnitCache](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AudioUnitCache
type TextToSpeechAudioUnitCache struct {
	objectivec.Object
}

// TextToSpeechAudioUnitCacheFromID constructs a [TextToSpeechAudioUnitCache] from an objc.ID.
func TextToSpeechAudioUnitCacheFromID(id objc.ID) TextToSpeechAudioUnitCache {
	return TextToSpeechAudioUnitCache{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechAudioUnitCache struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechAudioUnitCache embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechAudioUnitCache] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.AudioUnitCache
type ITextToSpeechAudioUnitCache interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechAudioUnitCache) Init() TextToSpeechAudioUnitCache {
	rv := objc.Send[TextToSpeechAudioUnitCache](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechAudioUnitCache) Autorelease() TextToSpeechAudioUnitCache {
	rv := objc.Send[TextToSpeechAudioUnitCache](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechAudioUnitCache creates a new TextToSpeechAudioUnitCache instance.
func NewTextToSpeechAudioUnitCache() TextToSpeechAudioUnitCache {
	class := getTextToSpeechAudioUnitCacheClass()
	rv := objc.Send[TextToSpeechAudioUnitCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}
