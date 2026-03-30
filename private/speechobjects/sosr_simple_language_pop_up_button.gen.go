// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSRSimpleLanguagePopUpButton] class.
var (
	_SOSRSimpleLanguagePopUpButtonClass     SOSRSimpleLanguagePopUpButtonClass
	_SOSRSimpleLanguagePopUpButtonClassOnce sync.Once
)

func getSOSRSimpleLanguagePopUpButtonClass() SOSRSimpleLanguagePopUpButtonClass {
	_SOSRSimpleLanguagePopUpButtonClassOnce.Do(func() {
		_SOSRSimpleLanguagePopUpButtonClass = SOSRSimpleLanguagePopUpButtonClass{class: objc.GetClass("SOSRSimpleLanguagePopUpButton")}
	})
	return _SOSRSimpleLanguagePopUpButtonClass
}

// GetSOSRSimpleLanguagePopUpButtonClass returns the class object for SOSRSimpleLanguagePopUpButton.
func GetSOSRSimpleLanguagePopUpButtonClass() SOSRSimpleLanguagePopUpButtonClass {
	return getSOSRSimpleLanguagePopUpButtonClass()
}

type SOSRSimpleLanguagePopUpButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRSimpleLanguagePopUpButtonClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRSimpleLanguagePopUpButtonClass) Alloc() SOSRSimpleLanguagePopUpButton {
	rv := objc.Send[SOSRSimpleLanguagePopUpButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSRSimpleLanguagePopUpButton._rowsFromSRLanguageItems]
//   - [SOSRSimpleLanguagePopUpButton._startDelayedPopUpUpdate]
//   - [SOSRSimpleLanguagePopUpButton._updateSRLanguageMenu]
//   - [SOSRSimpleLanguagePopUpButton.BuildPopUpButtonAndSelectLocaleIdentifierSupportedLocaleIdentifiers]
//   - [SOSRSimpleLanguagePopUpButton.PreviouslyChosenLocaleIdentifier]
//   - [SOSRSimpleLanguagePopUpButton.SetPreviouslyChosenLocaleIdentifier]
//   - [SOSRSimpleLanguagePopUpButton.SelectedLanguageItem]
//   - [SOSRSimpleLanguagePopUpButton.SupportedLocaleIdentifiers]
//   - [SOSRSimpleLanguagePopUpButton.SetSupportedLocaleIdentifiers]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton
type SOSRSimpleLanguagePopUpButton struct {
	appkit.NSPopUpButton
}

// SOSRSimpleLanguagePopUpButtonFromID constructs a [SOSRSimpleLanguagePopUpButton] from an objc.ID.
func SOSRSimpleLanguagePopUpButtonFromID(id objc.ID) SOSRSimpleLanguagePopUpButton {
	return SOSRSimpleLanguagePopUpButton{NSPopUpButton: appkit.NSPopUpButtonFromID(id)}
}

// Ensure SOSRSimpleLanguagePopUpButton implements ISOSRSimpleLanguagePopUpButton.
var _ ISOSRSimpleLanguagePopUpButton = SOSRSimpleLanguagePopUpButton{}

// An interface definition for the [SOSRSimpleLanguagePopUpButton] class.
//
// # Methods
//
//   - [ISOSRSimpleLanguagePopUpButton._rowsFromSRLanguageItems]
//   - [ISOSRSimpleLanguagePopUpButton._startDelayedPopUpUpdate]
//   - [ISOSRSimpleLanguagePopUpButton._updateSRLanguageMenu]
//   - [ISOSRSimpleLanguagePopUpButton.BuildPopUpButtonAndSelectLocaleIdentifierSupportedLocaleIdentifiers]
//   - [ISOSRSimpleLanguagePopUpButton.PreviouslyChosenLocaleIdentifier]
//   - [ISOSRSimpleLanguagePopUpButton.SetPreviouslyChosenLocaleIdentifier]
//   - [ISOSRSimpleLanguagePopUpButton.SelectedLanguageItem]
//   - [ISOSRSimpleLanguagePopUpButton.SupportedLocaleIdentifiers]
//   - [ISOSRSimpleLanguagePopUpButton.SetSupportedLocaleIdentifiers]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton
type ISOSRSimpleLanguagePopUpButton interface {
	appkit.INSPopUpButton

	// Topic: Methods

	_rowsFromSRLanguageItems(items objectivec.IObject) objectivec.IObject
	_startDelayedPopUpUpdate()
	_updateSRLanguageMenu()
	BuildPopUpButtonAndSelectLocaleIdentifierSupportedLocaleIdentifiers(identifier objectivec.IObject, identifiers objectivec.IObject)
	PreviouslyChosenLocaleIdentifier() string
	SetPreviouslyChosenLocaleIdentifier(value string)
	SelectedLanguageItem() objectivec.IObject
	SupportedLocaleIdentifiers() foundation.INSArray
	SetSupportedLocaleIdentifiers(value foundation.INSArray)
}

// Init initializes the instance.
func (s SOSRSimpleLanguagePopUpButton) Init() SOSRSimpleLanguagePopUpButton {
	rv := objc.Send[SOSRSimpleLanguagePopUpButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRSimpleLanguagePopUpButton) Autorelease() SOSRSimpleLanguagePopUpButton {
	rv := objc.Send[SOSRSimpleLanguagePopUpButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRSimpleLanguagePopUpButton creates a new SOSRSimpleLanguagePopUpButton instance.
func NewSOSRSimpleLanguagePopUpButton() SOSRSimpleLanguagePopUpButton {
	class := getSOSRSimpleLanguagePopUpButtonClass()
	rv := objc.Send[SOSRSimpleLanguagePopUpButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/initWithCoder:
func NewSOSRSimpleLanguagePopUpButtonWithCoder(coder objectivec.IObject) SOSRSimpleLanguagePopUpButton {
	instance := getSOSRSimpleLanguagePopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SOSRSimpleLanguagePopUpButtonFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/_rowsFromSRLanguageItems:
func (s SOSRSimpleLanguagePopUpButton) _rowsFromSRLanguageItems(items objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_rowsFromSRLanguageItems:"), items)
	return objectivec.Object{ID: rv}
}

// RowsFromSRLanguageItems is an exported wrapper for the private method _rowsFromSRLanguageItems.
func (s SOSRSimpleLanguagePopUpButton) RowsFromSRLanguageItems(items objectivec.IObject) objectivec.IObject {
	return s._rowsFromSRLanguageItems(items)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/_startDelayedPopUpUpdate
func (s SOSRSimpleLanguagePopUpButton) _startDelayedPopUpUpdate() {
	objc.Send[objc.ID](s.ID, objc.Sel("_startDelayedPopUpUpdate"))
}

// StartDelayedPopUpUpdate is an exported wrapper for the private method _startDelayedPopUpUpdate.
func (s SOSRSimpleLanguagePopUpButton) StartDelayedPopUpUpdate() {
	s._startDelayedPopUpUpdate()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/_updateSRLanguageMenu
func (s SOSRSimpleLanguagePopUpButton) _updateSRLanguageMenu() {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateSRLanguageMenu"))
}

// UpdateSRLanguageMenu is an exported wrapper for the private method _updateSRLanguageMenu.
func (s SOSRSimpleLanguagePopUpButton) UpdateSRLanguageMenu() {
	s._updateSRLanguageMenu()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/buildPopUpButtonAndSelectLocaleIdentifier:supportedLocaleIdentifiers:
func (s SOSRSimpleLanguagePopUpButton) BuildPopUpButtonAndSelectLocaleIdentifierSupportedLocaleIdentifiers(identifier objectivec.IObject, identifiers objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("buildPopUpButtonAndSelectLocaleIdentifier:supportedLocaleIdentifiers:"), identifier, identifiers)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/selectedLanguageItem
func (s SOSRSimpleLanguagePopUpButton) SelectedLanguageItem() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectedLanguageItem"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/previouslyChosenLocaleIdentifier
func (s SOSRSimpleLanguagePopUpButton) PreviouslyChosenLocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("previouslyChosenLocaleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOSRSimpleLanguagePopUpButton) SetPreviouslyChosenLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreviouslyChosenLocaleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRSimpleLanguagePopUpButton/supportedLocaleIdentifiers
func (s SOSRSimpleLanguagePopUpButton) SupportedLocaleIdentifiers() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedLocaleIdentifiers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SOSRSimpleLanguagePopUpButton) SetSupportedLocaleIdentifiers(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setSupportedLocaleIdentifiers:"), value)
}
