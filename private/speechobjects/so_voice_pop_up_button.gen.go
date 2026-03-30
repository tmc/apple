// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOVoicePopUpButton] class.
var (
	_SOVoicePopUpButtonClass     SOVoicePopUpButtonClass
	_SOVoicePopUpButtonClassOnce sync.Once
)

func getSOVoicePopUpButtonClass() SOVoicePopUpButtonClass {
	_SOVoicePopUpButtonClassOnce.Do(func() {
		_SOVoicePopUpButtonClass = SOVoicePopUpButtonClass{class: objc.GetClass("SOVoicePopUpButton")}
	})
	return _SOVoicePopUpButtonClass
}

// GetSOVoicePopUpButtonClass returns the class object for SOVoicePopUpButton.
func GetSOVoicePopUpButtonClass() SOVoicePopUpButtonClass {
	return getSOVoicePopUpButtonClass()
}

type SOVoicePopUpButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOVoicePopUpButtonClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOVoicePopUpButtonClass) Alloc() SOVoicePopUpButton {
	rv := objc.Send[SOVoicePopUpButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOVoicePopUpButton._forcePopupsToAdoptCachedMenuExceptPopUp]
//   - [SOVoicePopUpButton._handleSpeechDataInstallationNotification]
//   - [SOVoicePopUpButton._previouslyChosenVoiceIdentifier]
//   - [SOVoicePopUpButton._setPreviouslyChosenVoiceIdentifier]
//   - [SOVoicePopUpButton._statusStringForActiveDownloads]
//   - [SOVoicePopUpButton._updateDownloadStatusFields]
//   - [SOVoicePopUpButton.AddExcludedIdentifier]
//   - [SOVoicePopUpButton.BuildPopUpButtonWithSelectVoiceIdentifier]
//   - [SOVoicePopUpButton.IsSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmationParentWindowForSheet]
//   - [SOVoicePopUpButton.SelectedVoiceAttributes]
//   - [SOVoicePopUpButton.SetAllowSystemVoiceChoice]
//   - [SOVoicePopUpButton.SetShowIndividualVoiceQualities]
//   - [SOVoicePopUpButton.SetSystemVoiceLocalizedTextMenuItemTag]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton
type SOVoicePopUpButton struct {
	appkit.NSPopUpButton
}

// SOVoicePopUpButtonFromID constructs a [SOVoicePopUpButton] from an objc.ID.
func SOVoicePopUpButtonFromID(id objc.ID) SOVoicePopUpButton {
	return SOVoicePopUpButton{NSPopUpButton: appkit.NSPopUpButtonFromID(id)}
}

// Ensure SOVoicePopUpButton implements ISOVoicePopUpButton.
var _ ISOVoicePopUpButton = SOVoicePopUpButton{}

// An interface definition for the [SOVoicePopUpButton] class.
//
// # Methods
//
//   - [ISOVoicePopUpButton._forcePopupsToAdoptCachedMenuExceptPopUp]
//   - [ISOVoicePopUpButton._handleSpeechDataInstallationNotification]
//   - [ISOVoicePopUpButton._previouslyChosenVoiceIdentifier]
//   - [ISOVoicePopUpButton._setPreviouslyChosenVoiceIdentifier]
//   - [ISOVoicePopUpButton._statusStringForActiveDownloads]
//   - [ISOVoicePopUpButton._updateDownloadStatusFields]
//   - [ISOVoicePopUpButton.AddExcludedIdentifier]
//   - [ISOVoicePopUpButton.BuildPopUpButtonWithSelectVoiceIdentifier]
//   - [ISOVoicePopUpButton.IsSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmationParentWindowForSheet]
//   - [ISOVoicePopUpButton.SelectedVoiceAttributes]
//   - [ISOVoicePopUpButton.SetAllowSystemVoiceChoice]
//   - [ISOVoicePopUpButton.SetShowIndividualVoiceQualities]
//   - [ISOVoicePopUpButton.SetSystemVoiceLocalizedTextMenuItemTag]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton
type ISOVoicePopUpButton interface {
	appkit.INSPopUpButton

	// Topic: Methods

	_forcePopupsToAdoptCachedMenuExceptPopUp(up objectivec.IObject)
	_handleSpeechDataInstallationNotification()
	_previouslyChosenVoiceIdentifier() objectivec.IObject
	_setPreviouslyChosenVoiceIdentifier(identifier objectivec.IObject)
	_statusStringForActiveDownloads() objectivec.IObject
	_updateDownloadStatusFields()
	AddExcludedIdentifier(identifier objectivec.IObject)
	BuildPopUpButtonWithSelectVoiceIdentifier(identifier objectivec.IObject)
	IsSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmationParentWindowForSheet(confirmation bool, sheet objectivec.IObject) bool
	SelectedVoiceAttributes() objectivec.IObject
	SetAllowSystemVoiceChoice(choice bool)
	SetShowIndividualVoiceQualities(qualities bool)
	SetSystemVoiceLocalizedTextMenuItemTag(text objectivec.IObject, tag int64)
}

// Init initializes the instance.
func (s SOVoicePopUpButton) Init() SOVoicePopUpButton {
	rv := objc.Send[SOVoicePopUpButton](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOVoicePopUpButton) Autorelease() SOVoicePopUpButton {
	rv := objc.Send[SOVoicePopUpButton](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOVoicePopUpButton creates a new SOVoicePopUpButton instance.
func NewSOVoicePopUpButton() SOVoicePopUpButton {
	class := getSOVoicePopUpButtonClass()
	rv := objc.Send[SOVoicePopUpButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/initWithCoder:
func NewSOVoicePopUpButtonWithCoder(coder objectivec.IObject) SOVoicePopUpButton {
	instance := getSOVoicePopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SOVoicePopUpButtonFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_forcePopupsToAdoptCachedMenuExceptPopUp:
func (s SOVoicePopUpButton) _forcePopupsToAdoptCachedMenuExceptPopUp(up objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_forcePopupsToAdoptCachedMenuExceptPopUp:"), up)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_handleSpeechDataInstallationNotification
func (s SOVoicePopUpButton) _handleSpeechDataInstallationNotification() {
	objc.Send[objc.ID](s.ID, objc.Sel("_handleSpeechDataInstallationNotification"))
}

// HandleSpeechDataInstallationNotification is an exported wrapper for the private method _handleSpeechDataInstallationNotification.
func (s SOVoicePopUpButton) HandleSpeechDataInstallationNotification() {
	s._handleSpeechDataInstallationNotification()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_previouslyChosenVoiceIdentifier
func (s SOVoicePopUpButton) _previouslyChosenVoiceIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_previouslyChosenVoiceIdentifier"))
	return objectivec.Object{ID: rv}
}

// PreviouslyChosenVoiceIdentifier is an exported wrapper for the private method _previouslyChosenVoiceIdentifier.
func (s SOVoicePopUpButton) PreviouslyChosenVoiceIdentifier() objectivec.IObject {
	return s._previouslyChosenVoiceIdentifier()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_setPreviouslyChosenVoiceIdentifier:
func (s SOVoicePopUpButton) _setPreviouslyChosenVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_setPreviouslyChosenVoiceIdentifier:"), identifier)
}

// SetPreviouslyChosenVoiceIdentifier is an exported wrapper for the private method _setPreviouslyChosenVoiceIdentifier.
func (s SOVoicePopUpButton) SetPreviouslyChosenVoiceIdentifier(identifier objectivec.IObject) {
	s._setPreviouslyChosenVoiceIdentifier(identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_statusStringForActiveDownloads
func (s SOVoicePopUpButton) _statusStringForActiveDownloads() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_statusStringForActiveDownloads"))
	return objectivec.Object{ID: rv}
}

// StatusStringForActiveDownloads is an exported wrapper for the private method _statusStringForActiveDownloads.
func (s SOVoicePopUpButton) StatusStringForActiveDownloads() objectivec.IObject {
	return s._statusStringForActiveDownloads()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_updateDownloadStatusFields
func (s SOVoicePopUpButton) _updateDownloadStatusFields() {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateDownloadStatusFields"))
}

// UpdateDownloadStatusFields is an exported wrapper for the private method _updateDownloadStatusFields.
func (s SOVoicePopUpButton) UpdateDownloadStatusFields() {
	s._updateDownloadStatusFields()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/addExcludedIdentifier:
func (s SOVoicePopUpButton) AddExcludedIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addExcludedIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/buildPopUpButtonWithSelectVoiceIdentifier:
func (s SOVoicePopUpButton) BuildPopUpButtonWithSelectVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("buildPopUpButtonWithSelectVoiceIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/isSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmation:parentWindowForSheet:
func (s SOVoicePopUpButton) IsSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmationParentWindowForSheet(confirmation bool, sheet objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSelectedVoiceAppropriateForCurrentLanguageWithUserConfirmation:parentWindowForSheet:"), confirmation, sheet)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/selectedVoiceAttributes
func (s SOVoicePopUpButton) SelectedVoiceAttributes() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectedVoiceAttributes"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/setAllowSystemVoiceChoice:
func (s SOVoicePopUpButton) SetAllowSystemVoiceChoice(choice bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAllowSystemVoiceChoice:"), choice)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/setShowIndividualVoiceQualities:
func (s SOVoicePopUpButton) SetShowIndividualVoiceQualities(qualities bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setShowIndividualVoiceQualities:"), qualities)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/setSystemVoiceLocalizedText:menuItemTag:
func (s SOVoicePopUpButton) SetSystemVoiceLocalizedTextMenuItemTag(text objectivec.IObject, tag int64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSystemVoiceLocalizedText:menuItemTag:"), text, tag)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_forceAllVoicePopupsToUpdate
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) _forceAllVoicePopupsToUpdate() {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("_forceAllVoicePopupsToUpdate"))
}

// ForceAllVoicePopupsToUpdate is an exported wrapper for the private method _forceAllVoicePopupsToUpdate.
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) ForceAllVoicePopupsToUpdate() {
	_SOVoicePopUpButtonClass._forceAllVoicePopupsToUpdate()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/_startDelayedForceAllVoicePopupsToUpdate
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) _startDelayedForceAllVoicePopupsToUpdate() {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("_startDelayedForceAllVoicePopupsToUpdate"))
}

// StartDelayedForceAllVoicePopupsToUpdate is an exported wrapper for the private method _startDelayedForceAllVoicePopupsToUpdate.
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) StartDelayedForceAllVoicePopupsToUpdate() {
	_SOVoicePopUpButtonClass._startDelayedForceAllVoicePopupsToUpdate()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/addExcludedVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) AddExcludedVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("addExcludedVoiceIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/addRequiredVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) AddRequiredVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("addRequiredVoiceIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/buildSharedVoicesMenuShowingVOVoices:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) BuildSharedVoicesMenuShowingVOVoices(vOVoices bool) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("buildSharedVoicesMenuShowingVOVoices:"), vOVoices)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/displayNameForGender:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) DisplayNameForGender(gender objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("displayNameForGender:"), gender)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/forcePopupsToAdoptCachedMenuExceptPopUp:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) ForcePopupsToAdoptCachedMenuExceptPopUp(up objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("forcePopupsToAdoptCachedMenuExceptPopUp:"), up)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/isRequiredVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) IsRequiredVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("isRequiredVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/isSiriVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) IsSiriVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("isSiriVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/removeRequiredVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) RemoveRequiredVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("removeRequiredVoiceIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/setFallbackVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) SetFallbackVoiceIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("setFallbackVoiceIdentifier:"), identifier)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoicePopUpButton/shouldExcludeVoiceIdentifier:
func (_SOVoicePopUpButtonClass SOVoicePopUpButtonClass) ShouldExcludeVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoicePopUpButtonClass.class), objc.Sel("shouldExcludeVoiceIdentifier:"), identifier)
	return rv
}
