// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSettings] class.
var (
	_TextToSpeechTTSSettingsClass     TextToSpeechTTSSettingsClass
	_TextToSpeechTTSSettingsClassOnce sync.Once
)

func getTextToSpeechTTSSettingsClass() TextToSpeechTTSSettingsClass {
	_TextToSpeechTTSSettingsClassOnce.Do(func() {
		_TextToSpeechTTSSettingsClass = TextToSpeechTTSSettingsClass{class: objc.GetClass("TextToSpeech.TTSSettings")}
	})
	return _TextToSpeechTTSSettingsClass
}

// GetTextToSpeechTTSSettingsClass returns the class object for TextToSpeech.TTSSettings.
func GetTextToSpeechTTSSettingsClass() TextToSpeechTTSSettingsClass {
	return getTextToSpeechTTSSettingsClass()
}

type TextToSpeechTTSSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSSettingsClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSettingsClass) Alloc() TextToSpeechTTSSettings {
	rv := objc.Send[TextToSpeechTTSSettings](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSettings
type TextToSpeechTTSSettings struct {
	objectivec.Object
}

// TextToSpeechTTSSettingsFromID constructs a [TextToSpeechTTSSettings] from an objc.ID.
func TextToSpeechTTSSettingsFromID(id objc.ID) TextToSpeechTTSSettings {
	return TextToSpeechTTSSettings{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechTTSSettings struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechTTSSettings embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechTTSSettings] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSettings
type ITextToSpeechTTSSettings interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSettings) Init() TextToSpeechTTSSettings {
	rv := objc.Send[TextToSpeechTTSSettings](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSettings) Autorelease() TextToSpeechTTSSettings {
	rv := objc.Send[TextToSpeechTTSSettings](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSettings creates a new TextToSpeechTTSSettings instance.
func NewTextToSpeechTTSSettings() TextToSpeechTTSSettings {
	class := getTextToSpeechTTSSettingsClass()
	rv := objc.Send[TextToSpeechTTSSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

