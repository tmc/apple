// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSFallbackLoader] class.
var (
	_TTSFallbackLoaderClass     TTSFallbackLoaderClass
	_TTSFallbackLoaderClassOnce sync.Once
)

func getTTSFallbackLoaderClass() TTSFallbackLoaderClass {
	_TTSFallbackLoaderClassOnce.Do(func() {
		_TTSFallbackLoaderClass = TTSFallbackLoaderClass{class: objc.GetClass("TTSFallbackLoader")}
	})
	return _TTSFallbackLoaderClass
}

// GetTTSFallbackLoaderClass returns the class object for TTSFallbackLoader.
func GetTTSFallbackLoaderClass() TTSFallbackLoaderClass {
	return getTTSFallbackLoaderClass()
}

type TTSFallbackLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSFallbackLoaderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSFallbackLoaderClass) Alloc() TTSFallbackLoader {
	rv := objc.Send[TTSFallbackLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFallbackLoader
type TTSFallbackLoader struct {
	objectivec.Object
}

// TTSFallbackLoaderFromID constructs a [TTSFallbackLoader] from an objc.ID.
func TTSFallbackLoaderFromID(id objc.ID) TTSFallbackLoader {
	return TTSFallbackLoader{objectivec.Object{ID: id}}
}

// Ensure TTSFallbackLoader implements ITTSFallbackLoader.
var _ ITTSFallbackLoader = TTSFallbackLoader{}

// An interface definition for the [TTSFallbackLoader] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSFallbackLoader
type ITTSFallbackLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSFallbackLoader) Init() TTSFallbackLoader {
	rv := objc.Send[TTSFallbackLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSFallbackLoader) Autorelease() TTSFallbackLoader {
	rv := objc.Send[TTSFallbackLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSFallbackLoader creates a new TTSFallbackLoader instance.
func NewTTSFallbackLoader() TTSFallbackLoader {
	class := getTTSFallbackLoaderClass()
	rv := objc.Send[TTSFallbackLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFallbackLoader/fallbackRendererClass
func (_TTSFallbackLoaderClass TTSFallbackLoaderClass) FallbackRendererClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_TTSFallbackLoaderClass.class), objc.Sel("fallbackRendererClass"))
	return rv
}
