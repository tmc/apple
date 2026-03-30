// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSRLanguagePopUpButton] class.
var (
	_SOSRLanguagePopUpButtonClass     SOSRLanguagePopUpButtonClass
	_SOSRLanguagePopUpButtonClassOnce sync.Once
)

func getSOSRLanguagePopUpButtonClass() SOSRLanguagePopUpButtonClass {
	_SOSRLanguagePopUpButtonClassOnce.Do(func() {
		_SOSRLanguagePopUpButtonClass = SOSRLanguagePopUpButtonClass{class: objc.GetClass("SOSRLanguagePopUpButton")}
	})
	return _SOSRLanguagePopUpButtonClass
}

// GetSOSRLanguagePopUpButtonClass returns the class object for SOSRLanguagePopUpButton.
func GetSOSRLanguagePopUpButtonClass() SOSRLanguagePopUpButtonClass {
	return getSOSRLanguagePopUpButtonClass()
}

type SOSRLanguagePopUpButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguagePopUpButtonClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguagePopUpButtonClass) Alloc() SOSRLanguagePopUpButton {
	rv := objc.Send[SOSRLanguagePopUpButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSRLanguagePopUpButton._clearDownloadStatusFieldAfterPreHeat]
//   - [SOSRLanguagePopUpButton._initCommon]
//   - [SOSRLanguagePopUpButton._startDelayedPopUpUpdate]
//   - [SOSRLanguagePopUpButton._statusStringForActiveDownloads]
//   - [SOSRLanguagePopUpButton._updateDownloadStatusFields]
//   - [SOSRLanguagePopUpButton._updateSRLanguageMenu]
//   - [SOSRLanguagePopUpButton.BuildPopUpButtonAndSelectLocaleIdentifierNetworkSupportedLocaleIdentifiersOfflineSupportedLocaleIdentifiers]
//   - [SOSRLanguagePopUpButton.InstallationFinished]
//   - [SOSRLanguagePopUpButton.LanguagesAreDownloadable]
//   - [SOSRLanguagePopUpButton.SetLanguagesAreDownloadable]
//   - [SOSRLanguagePopUpButton.NetworkBasedLocaleIdentifiers]
//   - [SOSRLanguagePopUpButton.SetNetworkBasedLocaleIdentifiers]
//   - [SOSRLanguagePopUpButton.OfflineBasedLocaleIdentifiers]
//   - [SOSRLanguagePopUpButton.SetOfflineBasedLocaleIdentifiers]
//   - [SOSRLanguagePopUpButton.PreviouslyChosenLocaleIdentifier]
//   - [SOSRLanguagePopUpButton.SetPreviouslyChosenLocaleIdentifier]
//   - [SOSRLanguagePopUpButton.SelectedLanguageItem]
//   - [SOSRLanguagePopUpButton.ShowOnlyNetworkSupportedItems]
//   - [SOSRLanguagePopUpButton.SetShowOnlyNetworkSupportedItems]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton
type SOSRLanguagePopUpButton struct {
	appkit.NSPopUpButton
}

// SOSRLanguagePopUpButtonFromID constructs a [SOSRLanguagePopUpButton] from an objc.ID.
func SOSRLanguagePopUpButtonFromID(id objc.ID) SOSRLanguagePopUpButton {
	return SOSRLanguagePopUpButton{NSPopUpButton: appkit.NSPopUpButtonFromID(id)}
}

// Ensure SOSRLanguagePopUpButton implements ISOSRLanguagePopUpButton.
var _ ISOSRLanguagePopUpButton = SOSRLanguagePopUpButton{}

// An interface definition for the [SOSRLanguagePopUpButton] class.
//
// # Methods
//
//   - [ISOSRLanguagePopUpButton._clearDownloadStatusFieldAfterPreHeat]
//   - [ISOSRLanguagePopUpButton._initCommon]
//   - [ISOSRLanguagePopUpButton._startDelayedPopUpUpdate]
//   - [ISOSRLanguagePopUpButton._statusStringForActiveDownloads]
//   - [ISOSRLanguagePopUpButton._updateDownloadStatusFields]
//   - [ISOSRLanguagePopUpButton._updateSRLanguageMenu]
//   - [ISOSRLanguagePopUpButton.BuildPopUpButtonAndSelectLocaleIdentifierNetworkSupportedLocaleIdentifiersOfflineSupportedLocaleIdentifiers]
//   - [ISOSRLanguagePopUpButton.InstallationFinished]
//   - [ISOSRLanguagePopUpButton.LanguagesAreDownloadable]
//   - [ISOSRLanguagePopUpButton.SetLanguagesAreDownloadable]
//   - [ISOSRLanguagePopUpButton.NetworkBasedLocaleIdentifiers]
//   - [ISOSRLanguagePopUpButton.SetNetworkBasedLocaleIdentifiers]
//   - [ISOSRLanguagePopUpButton.OfflineBasedLocaleIdentifiers]
//   - [ISOSRLanguagePopUpButton.SetOfflineBasedLocaleIdentifiers]
//   - [ISOSRLanguagePopUpButton.PreviouslyChosenLocaleIdentifier]
//   - [ISOSRLanguagePopUpButton.SetPreviouslyChosenLocaleIdentifier]
//   - [ISOSRLanguagePopUpButton.SelectedLanguageItem]
//   - [ISOSRLanguagePopUpButton.ShowOnlyNetworkSupportedItems]
//   - [ISOSRLanguagePopUpButton.SetShowOnlyNetworkSupportedItems]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton
type ISOSRLanguagePopUpButton interface {
	appkit.INSPopUpButton

	// Topic: Methods

	_clearDownloadStatusFieldAfterPreHeat()
	_initCommon()
	_startDelayedPopUpUpdate()
	_statusStringForActiveDownloads() objectivec.IObject
	_updateDownloadStatusFields()
	_updateSRLanguageMenu()
	BuildPopUpButtonAndSelectLocaleIdentifierNetworkSupportedLocaleIdentifiersOfflineSupportedLocaleIdentifiers(identifier objectivec.IObject, identifiers objectivec.IObject, identifiers2 objectivec.IObject)
	InstallationFinished(finished objectivec.IObject)
	LanguagesAreDownloadable() bool
	SetLanguagesAreDownloadable(value bool)
	NetworkBasedLocaleIdentifiers() foundation.INSArray
	SetNetworkBasedLocaleIdentifiers(value foundation.INSArray)
	OfflineBasedLocaleIdentifiers() foundation.INSArray
	SetOfflineBasedLocaleIdentifiers(value foundation.INSArray)
	PreviouslyChosenLocaleIdentifier() string
	SetPreviouslyChosenLocaleIdentifier(value string)
	SelectedLanguageItem() objectivec.IObject
	ShowOnlyNetworkSupportedItems() bool
	SetShowOnlyNetworkSupportedItems(value bool)
}

// Init initializes the instance.
func (s SOSRLanguagePopUpButton) Init() SOSRLanguagePopUpButton {
	rv := objc.Send[SOSRLanguagePopUpButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguagePopUpButton) Autorelease() SOSRLanguagePopUpButton {
	rv := objc.Send[SOSRLanguagePopUpButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguagePopUpButton creates a new SOSRLanguagePopUpButton instance.
func NewSOSRLanguagePopUpButton() SOSRLanguagePopUpButton {
	class := getSOSRLanguagePopUpButtonClass()
	rv := objc.Send[SOSRLanguagePopUpButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/initWithCoder:
func NewSOSRLanguagePopUpButtonWithCoder(coder objectivec.IObject) SOSRLanguagePopUpButton {
	instance := getSOSRLanguagePopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SOSRLanguagePopUpButtonFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/initWithFrame:
func NewSOSRLanguagePopUpButtonWithFrame(frame corefoundation.CGRect) SOSRLanguagePopUpButton {
	instance := getSOSRLanguagePopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frame)
	return SOSRLanguagePopUpButtonFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_clearDownloadStatusFieldAfterPreHeat
func (s SOSRLanguagePopUpButton) _clearDownloadStatusFieldAfterPreHeat() {
	objc.Send[objc.ID](s.ID, objc.Sel("_clearDownloadStatusFieldAfterPreHeat"))
}

// ClearDownloadStatusFieldAfterPreHeat is an exported wrapper for the private method _clearDownloadStatusFieldAfterPreHeat.
func (s SOSRLanguagePopUpButton) ClearDownloadStatusFieldAfterPreHeat() {
	s._clearDownloadStatusFieldAfterPreHeat()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_initCommon
func (s SOSRLanguagePopUpButton) _initCommon() {
	objc.Send[objc.ID](s.ID, objc.Sel("_initCommon"))
}

// InitCommon is an exported wrapper for the private method _initCommon.
func (s SOSRLanguagePopUpButton) InitCommon() {
	s._initCommon()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_startDelayedPopUpUpdate
func (s SOSRLanguagePopUpButton) _startDelayedPopUpUpdate() {
	objc.Send[objc.ID](s.ID, objc.Sel("_startDelayedPopUpUpdate"))
}

// StartDelayedPopUpUpdate is an exported wrapper for the private method _startDelayedPopUpUpdate.
func (s SOSRLanguagePopUpButton) StartDelayedPopUpUpdate() {
	s._startDelayedPopUpUpdate()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_statusStringForActiveDownloads
func (s SOSRLanguagePopUpButton) _statusStringForActiveDownloads() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_statusStringForActiveDownloads"))
	return objectivec.Object{ID: rv}
}

// StatusStringForActiveDownloads is an exported wrapper for the private method _statusStringForActiveDownloads.
func (s SOSRLanguagePopUpButton) StatusStringForActiveDownloads() objectivec.IObject {
	return s._statusStringForActiveDownloads()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_updateDownloadStatusFields
func (s SOSRLanguagePopUpButton) _updateDownloadStatusFields() {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateDownloadStatusFields"))
}

// UpdateDownloadStatusFields is an exported wrapper for the private method _updateDownloadStatusFields.
func (s SOSRLanguagePopUpButton) UpdateDownloadStatusFields() {
	s._updateDownloadStatusFields()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/_updateSRLanguageMenu
func (s SOSRLanguagePopUpButton) _updateSRLanguageMenu() {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateSRLanguageMenu"))
}

// UpdateSRLanguageMenu is an exported wrapper for the private method _updateSRLanguageMenu.
func (s SOSRLanguagePopUpButton) UpdateSRLanguageMenu() {
	s._updateSRLanguageMenu()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/buildPopUpButtonAndSelectLocaleIdentifier:networkSupportedLocaleIdentifiers:offlineSupportedLocaleIdentifiers:
func (s SOSRLanguagePopUpButton) BuildPopUpButtonAndSelectLocaleIdentifierNetworkSupportedLocaleIdentifiersOfflineSupportedLocaleIdentifiers(identifier objectivec.IObject, identifiers objectivec.IObject, identifiers2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("buildPopUpButtonAndSelectLocaleIdentifier:networkSupportedLocaleIdentifiers:offlineSupportedLocaleIdentifiers:"), identifier, identifiers, identifiers2)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/installationFinished:
func (s SOSRLanguagePopUpButton) InstallationFinished(finished objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("installationFinished:"), finished)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/selectedLanguageItem
func (s SOSRLanguagePopUpButton) SelectedLanguageItem() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectedLanguageItem"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/languagesAreDownloadable
func (s SOSRLanguagePopUpButton) LanguagesAreDownloadable() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("languagesAreDownloadable"))
	return rv
}
func (s SOSRLanguagePopUpButton) SetLanguagesAreDownloadable(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setLanguagesAreDownloadable:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/networkBasedLocaleIdentifiers
func (s SOSRLanguagePopUpButton) NetworkBasedLocaleIdentifiers() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("networkBasedLocaleIdentifiers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SOSRLanguagePopUpButton) SetNetworkBasedLocaleIdentifiers(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setNetworkBasedLocaleIdentifiers:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/offlineBasedLocaleIdentifiers
func (s SOSRLanguagePopUpButton) OfflineBasedLocaleIdentifiers() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("offlineBasedLocaleIdentifiers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SOSRLanguagePopUpButton) SetOfflineBasedLocaleIdentifiers(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setOfflineBasedLocaleIdentifiers:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/previouslyChosenLocaleIdentifier
func (s SOSRLanguagePopUpButton) PreviouslyChosenLocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("previouslyChosenLocaleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOSRLanguagePopUpButton) SetPreviouslyChosenLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreviouslyChosenLocaleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguagePopUpButton/showOnlyNetworkSupportedItems
func (s SOSRLanguagePopUpButton) ShowOnlyNetworkSupportedItems() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showOnlyNetworkSupportedItems"))
	return rv
}
func (s SOSRLanguagePopUpButton) SetShowOnlyNetworkSupportedItems(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowOnlyNetworkSupportedItems:"), value)
}
