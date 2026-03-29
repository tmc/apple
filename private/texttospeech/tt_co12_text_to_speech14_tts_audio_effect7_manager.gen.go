// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCO12TextToSpeech14TTSAudioEffect7Manager] class.
var (
	_TtCO12TextToSpeech14TTSAudioEffect7ManagerClass     TtCO12TextToSpeech14TTSAudioEffect7ManagerClass
	_TtCO12TextToSpeech14TTSAudioEffect7ManagerClassOnce sync.Once
)

func getTtCO12TextToSpeech14TTSAudioEffect7ManagerClass() TtCO12TextToSpeech14TTSAudioEffect7ManagerClass {
	_TtCO12TextToSpeech14TTSAudioEffect7ManagerClassOnce.Do(func() {
		_TtCO12TextToSpeech14TTSAudioEffect7ManagerClass = TtCO12TextToSpeech14TTSAudioEffect7ManagerClass{class: objc.GetClass("_TtCO12TextToSpeech14TTSAudioEffect7Manager")}
	})
	return _TtCO12TextToSpeech14TTSAudioEffect7ManagerClass
}

// GetTtCO12TextToSpeech14TTSAudioEffect7ManagerClass returns the class object for _TtCO12TextToSpeech14TTSAudioEffect7Manager.
func GetTtCO12TextToSpeech14TTSAudioEffect7ManagerClass() TtCO12TextToSpeech14TTSAudioEffect7ManagerClass {
	return getTtCO12TextToSpeech14TTSAudioEffect7ManagerClass()
}

type TtCO12TextToSpeech14TTSAudioEffect7ManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCO12TextToSpeech14TTSAudioEffect7ManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCO12TextToSpeech14TTSAudioEffect7ManagerClass) Alloc() TtCO12TextToSpeech14TTSAudioEffect7Manager {
	rv := objc.Send[TtCO12TextToSpeech14TTSAudioEffect7Manager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCO12TextToSpeech14TTSAudioEffect7Manager
type TtCO12TextToSpeech14TTSAudioEffect7Manager struct {
	objectivec.Object
}

// TtCO12TextToSpeech14TTSAudioEffect7ManagerFromID constructs a [TtCO12TextToSpeech14TTSAudioEffect7Manager] from an objc.ID.
func TtCO12TextToSpeech14TTSAudioEffect7ManagerFromID(id objc.ID) TtCO12TextToSpeech14TTSAudioEffect7Manager {
	return TtCO12TextToSpeech14TTSAudioEffect7Manager{objectivec.Object{ID: id}}
}
// NOTE: TtCO12TextToSpeech14TTSAudioEffect7Manager struct embeds objectivec.Object (parent type unavailable) but
// ITtCO12TextToSpeech14TTSAudioEffect7Manager embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCO12TextToSpeech14TTSAudioEffect7Manager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCO12TextToSpeech14TTSAudioEffect7Manager
type ITtCO12TextToSpeech14TTSAudioEffect7Manager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCO12TextToSpeech14TTSAudioEffect7Manager) Init() TtCO12TextToSpeech14TTSAudioEffect7Manager {
	rv := objc.Send[TtCO12TextToSpeech14TTSAudioEffect7Manager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCO12TextToSpeech14TTSAudioEffect7Manager) Autorelease() TtCO12TextToSpeech14TTSAudioEffect7Manager {
	rv := objc.Send[TtCO12TextToSpeech14TTSAudioEffect7Manager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCO12TextToSpeech14TTSAudioEffect7Manager creates a new TtCO12TextToSpeech14TTSAudioEffect7Manager instance.
func NewTtCO12TextToSpeech14TTSAudioEffect7Manager() TtCO12TextToSpeech14TTSAudioEffect7Manager {
	class := getTtCO12TextToSpeech14TTSAudioEffect7ManagerClass()
	rv := objc.Send[TtCO12TextToSpeech14TTSAudioEffect7Manager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

