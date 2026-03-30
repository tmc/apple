// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SOVoiceRowCheckboxButton] class.
var (
	_SOVoiceRowCheckboxButtonClass     SOVoiceRowCheckboxButtonClass
	_SOVoiceRowCheckboxButtonClassOnce sync.Once
)

func getSOVoiceRowCheckboxButtonClass() SOVoiceRowCheckboxButtonClass {
	_SOVoiceRowCheckboxButtonClassOnce.Do(func() {
		_SOVoiceRowCheckboxButtonClass = SOVoiceRowCheckboxButtonClass{class: objc.GetClass("SOVoiceRowCheckboxButton")}
	})
	return _SOVoiceRowCheckboxButtonClass
}

// GetSOVoiceRowCheckboxButtonClass returns the class object for SOVoiceRowCheckboxButton.
func GetSOVoiceRowCheckboxButtonClass() SOVoiceRowCheckboxButtonClass {
	return getSOVoiceRowCheckboxButtonClass()
}

type SOVoiceRowCheckboxButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOVoiceRowCheckboxButtonClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOVoiceRowCheckboxButtonClass) Alloc() SOVoiceRowCheckboxButton {
	rv := objc.Send[SOVoiceRowCheckboxButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOVoiceRowCheckboxButton.VoiceIdentifier]
//   - [SOVoiceRowCheckboxButton.SetVoiceIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceRowCheckboxButton
type SOVoiceRowCheckboxButton struct {
	appkit.NSButton
}

// SOVoiceRowCheckboxButtonFromID constructs a [SOVoiceRowCheckboxButton] from an objc.ID.
func SOVoiceRowCheckboxButtonFromID(id objc.ID) SOVoiceRowCheckboxButton {
	return SOVoiceRowCheckboxButton{NSButton: appkit.NSButtonFromID(id)}
}

// Ensure SOVoiceRowCheckboxButton implements ISOVoiceRowCheckboxButton.
var _ ISOVoiceRowCheckboxButton = SOVoiceRowCheckboxButton{}

// An interface definition for the [SOVoiceRowCheckboxButton] class.
//
// # Methods
//
//   - [ISOVoiceRowCheckboxButton.VoiceIdentifier]
//   - [ISOVoiceRowCheckboxButton.SetVoiceIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceRowCheckboxButton
type ISOVoiceRowCheckboxButton interface {
	appkit.INSButton

	// Topic: Methods

	VoiceIdentifier() string
	SetVoiceIdentifier(value string)
}

// Init initializes the instance.
func (s SOVoiceRowCheckboxButton) Init() SOVoiceRowCheckboxButton {
	rv := objc.Send[SOVoiceRowCheckboxButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOVoiceRowCheckboxButton) Autorelease() SOVoiceRowCheckboxButton {
	rv := objc.Send[SOVoiceRowCheckboxButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOVoiceRowCheckboxButton creates a new SOVoiceRowCheckboxButton instance.
func NewSOVoiceRowCheckboxButton() SOVoiceRowCheckboxButton {
	class := getSOVoiceRowCheckboxButtonClass()
	rv := objc.Send[SOVoiceRowCheckboxButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceRowCheckboxButton/voiceIdentifier
func (s SOVoiceRowCheckboxButton) VoiceIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOVoiceRowCheckboxButton) SetVoiceIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setVoiceIdentifier:"), objc.String(value))
}
