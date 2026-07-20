// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Preferences] class.
var (
	_PreferencesClass     PreferencesClass
	_PreferencesClassOnce sync.Once
)

func getPreferencesClass() PreferencesClass {
	_PreferencesClassOnce.Do(func() {
		_PreferencesClass = PreferencesClass{class: objc.GetClass("_TtCC12TextToSpeech13VoiceDatabase11Preferences")}
	})
	return _PreferencesClass
}

// GetPreferencesClass returns the class object for _TtCC12TextToSpeech13VoiceDatabase11Preferences.
func GetPreferencesClass() PreferencesClass {
	return getPreferencesClass()
}

type PreferencesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PreferencesClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PreferencesClass) Alloc() Preferences {
	rv := objc.Send[Preferences](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type Preferences struct {
	objectivec.Object
}

// PreferencesFromID constructs a [Preferences] from an objc.ID.
func PreferencesFromID(id objc.ID) Preferences {
	return Preferences{objectivec.Object{ID: id}}
}

// NOTE: Preferences struct embeds objectivec.Object (parent type unavailable) but
// IPreferences embeds the parent interface; skip compile-time assertion.

// An interface definition for the [Preferences] class.
type IPreferences interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p Preferences) Init() Preferences {
	rv := objc.Send[Preferences](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Preferences) Autorelease() Preferences {
	rv := objc.Send[Preferences](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreferences creates a new Preferences instance.
func NewPreferences() Preferences {
	class := getPreferencesClass()
	rv := objc.Send[Preferences](objc.ID(class.class), objc.Sel("new"))
	return rv
}
