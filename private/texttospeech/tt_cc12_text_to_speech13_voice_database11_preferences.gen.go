// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC12TextToSpeech13VoiceDatabase11Preferences] class.
var (
	_TtCC12TextToSpeech13VoiceDatabase11PreferencesClass     TtCC12TextToSpeech13VoiceDatabase11PreferencesClass
	_TtCC12TextToSpeech13VoiceDatabase11PreferencesClassOnce sync.Once
)

func getTtCC12TextToSpeech13VoiceDatabase11PreferencesClass() TtCC12TextToSpeech13VoiceDatabase11PreferencesClass {
	_TtCC12TextToSpeech13VoiceDatabase11PreferencesClassOnce.Do(func() {
		_TtCC12TextToSpeech13VoiceDatabase11PreferencesClass = TtCC12TextToSpeech13VoiceDatabase11PreferencesClass{class: objc.GetClass("_TtCC12TextToSpeech13VoiceDatabase11Preferences")}
	})
	return _TtCC12TextToSpeech13VoiceDatabase11PreferencesClass
}

// GetTtCC12TextToSpeech13VoiceDatabase11PreferencesClass returns the class object for _TtCC12TextToSpeech13VoiceDatabase11Preferences.
func GetTtCC12TextToSpeech13VoiceDatabase11PreferencesClass() TtCC12TextToSpeech13VoiceDatabase11PreferencesClass {
	return getTtCC12TextToSpeech13VoiceDatabase11PreferencesClass()
}

type TtCC12TextToSpeech13VoiceDatabase11PreferencesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC12TextToSpeech13VoiceDatabase11PreferencesClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC12TextToSpeech13VoiceDatabase11PreferencesClass) Alloc() TtCC12TextToSpeech13VoiceDatabase11Preferences {
	rv := objc.Send[TtCC12TextToSpeech13VoiceDatabase11Preferences](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech13VoiceDatabase11Preferences
type TtCC12TextToSpeech13VoiceDatabase11Preferences struct {
	objectivec.Object
}

// TtCC12TextToSpeech13VoiceDatabase11PreferencesFromID constructs a [TtCC12TextToSpeech13VoiceDatabase11Preferences] from an objc.ID.
func TtCC12TextToSpeech13VoiceDatabase11PreferencesFromID(id objc.ID) TtCC12TextToSpeech13VoiceDatabase11Preferences {
	return TtCC12TextToSpeech13VoiceDatabase11Preferences{objectivec.Object{ID: id}}
}
// NOTE: TtCC12TextToSpeech13VoiceDatabase11Preferences struct embeds objectivec.Object (parent type unavailable) but
// ITtCC12TextToSpeech13VoiceDatabase11Preferences embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC12TextToSpeech13VoiceDatabase11Preferences] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech13VoiceDatabase11Preferences
type ITtCC12TextToSpeech13VoiceDatabase11Preferences interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC12TextToSpeech13VoiceDatabase11Preferences) Init() TtCC12TextToSpeech13VoiceDatabase11Preferences {
	rv := objc.Send[TtCC12TextToSpeech13VoiceDatabase11Preferences](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC12TextToSpeech13VoiceDatabase11Preferences) Autorelease() TtCC12TextToSpeech13VoiceDatabase11Preferences {
	rv := objc.Send[TtCC12TextToSpeech13VoiceDatabase11Preferences](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC12TextToSpeech13VoiceDatabase11Preferences creates a new TtCC12TextToSpeech13VoiceDatabase11Preferences instance.
func NewTtCC12TextToSpeech13VoiceDatabase11Preferences() TtCC12TextToSpeech13VoiceDatabase11Preferences {
	class := getTtCC12TextToSpeech13VoiceDatabase11PreferencesClass()
	rv := objc.Send[TtCC12TextToSpeech13VoiceDatabase11Preferences](objc.ID(class.class), objc.Sel("new"))
	return rv
}

