// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SOSRLanguageRowCheckboxButton] class.
var (
	_SOSRLanguageRowCheckboxButtonClass     SOSRLanguageRowCheckboxButtonClass
	_SOSRLanguageRowCheckboxButtonClassOnce sync.Once
)

func getSOSRLanguageRowCheckboxButtonClass() SOSRLanguageRowCheckboxButtonClass {
	_SOSRLanguageRowCheckboxButtonClassOnce.Do(func() {
		_SOSRLanguageRowCheckboxButtonClass = SOSRLanguageRowCheckboxButtonClass{class: objc.GetClass("SOSRLanguageRowCheckboxButton")}
	})
	return _SOSRLanguageRowCheckboxButtonClass
}

// GetSOSRLanguageRowCheckboxButtonClass returns the class object for SOSRLanguageRowCheckboxButton.
func GetSOSRLanguageRowCheckboxButtonClass() SOSRLanguageRowCheckboxButtonClass {
	return getSOSRLanguageRowCheckboxButtonClass()
}

type SOSRLanguageRowCheckboxButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguageRowCheckboxButtonClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguageRowCheckboxButtonClass) Alloc() SOSRLanguageRowCheckboxButton {
	rv := objc.Send[SOSRLanguageRowCheckboxButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSRLanguageRowCheckboxButton.LocaleIdentifier]
//   - [SOSRLanguageRowCheckboxButton.SetLocaleIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRowCheckboxButton
type SOSRLanguageRowCheckboxButton struct {
	appkit.NSButton
}

// SOSRLanguageRowCheckboxButtonFromID constructs a [SOSRLanguageRowCheckboxButton] from an objc.ID.
func SOSRLanguageRowCheckboxButtonFromID(id objc.ID) SOSRLanguageRowCheckboxButton {
	return SOSRLanguageRowCheckboxButton{NSButton: appkit.NSButtonFromID(id)}
}

// Ensure SOSRLanguageRowCheckboxButton implements ISOSRLanguageRowCheckboxButton.
var _ ISOSRLanguageRowCheckboxButton = SOSRLanguageRowCheckboxButton{}

// An interface definition for the [SOSRLanguageRowCheckboxButton] class.
//
// # Methods
//
//   - [ISOSRLanguageRowCheckboxButton.LocaleIdentifier]
//   - [ISOSRLanguageRowCheckboxButton.SetLocaleIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRowCheckboxButton
type ISOSRLanguageRowCheckboxButton interface {
	appkit.INSButton

	// Topic: Methods

	LocaleIdentifier() string
	SetLocaleIdentifier(value string)
}

// Init initializes the instance.
func (s SOSRLanguageRowCheckboxButton) Init() SOSRLanguageRowCheckboxButton {
	rv := objc.Send[SOSRLanguageRowCheckboxButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguageRowCheckboxButton) Autorelease() SOSRLanguageRowCheckboxButton {
	rv := objc.Send[SOSRLanguageRowCheckboxButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguageRowCheckboxButton creates a new SOSRLanguageRowCheckboxButton instance.
func NewSOSRLanguageRowCheckboxButton() SOSRLanguageRowCheckboxButton {
	class := getSOSRLanguageRowCheckboxButtonClass()
	rv := objc.Send[SOSRLanguageRowCheckboxButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRowCheckboxButton/localeIdentifier
func (s SOSRLanguageRowCheckboxButton) LocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOSRLanguageRowCheckboxButton) SetLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLocaleIdentifier:"), objc.String(value))
}
