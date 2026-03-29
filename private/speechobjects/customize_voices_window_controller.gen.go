// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CustomizeVoicesWindowController] class.
var (
	_CustomizeVoicesWindowControllerClass     CustomizeVoicesWindowControllerClass
	_CustomizeVoicesWindowControllerClassOnce sync.Once
)

func getCustomizeVoicesWindowControllerClass() CustomizeVoicesWindowControllerClass {
	_CustomizeVoicesWindowControllerClassOnce.Do(func() {
		_CustomizeVoicesWindowControllerClass = CustomizeVoicesWindowControllerClass{class: objc.GetClass("CustomizeVoicesWindowController")}
	})
	return _CustomizeVoicesWindowControllerClass
}

// GetCustomizeVoicesWindowControllerClass returns the class object for CustomizeVoicesWindowController.
func GetCustomizeVoicesWindowControllerClass() CustomizeVoicesWindowControllerClass {
	return getCustomizeVoicesWindowControllerClass()
}

type CustomizeVoicesWindowControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CustomizeVoicesWindowControllerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CustomizeVoicesWindowControllerClass) Alloc() CustomizeVoicesWindowController {
	rv := objc.Send[CustomizeVoicesWindowController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [CustomizeVoicesWindowController._anyVoicePlaying]
//   - [CustomizeVoicesWindowController._delayedPopUpUpdate]
//   - [CustomizeVoicesWindowController._isSampleAvailableForVoiceObject]
//   - [CustomizeVoicesWindowController._propagateCheckboxSelection]
//   - [CustomizeVoicesWindowController._propagateDownloadCheckboxSelection]
//   - [CustomizeVoicesWindowController._propagateDownloadVariantSelection]
//   - [CustomizeVoicesWindowController._rebuildVoiceList]
//   - [CustomizeVoicesWindowController._setRowDownloadCheckboxVoiceObjectIsSelected]
//   - [CustomizeVoicesWindowController._setRowStatusFieldViewVoiceObjectIsSelected]
//   - [CustomizeVoicesWindowController._shouldAllowRemovalOfVoiceObject]
//   - [CustomizeVoicesWindowController._showPlayStopButtonAsPlaying]
//   - [CustomizeVoicesWindowController._stopAndResetAllVoicePlaying]
//   - [CustomizeVoicesWindowController._updateButtonStates]
//   - [CustomizeVoicesWindowController._updateDisplayedVoicesUsingFilterString]
//   - [CustomizeVoicesWindowController._updateRowDownloadStatus]
//   - [CustomizeVoicesWindowController._voiceObjectForCurrentlySelectedRow]
//   - [CustomizeVoicesWindowController.AcceptVoiceSelection]
//   - [CustomizeVoicesWindowController.CancelVoiceSelection]
//   - [CustomizeVoicesWindowController.NumberOfRowsInTableView]
//   - [CustomizeVoicesWindowController.SearchFieldChanged]
//   - [CustomizeVoicesWindowController.ShowSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved]
//   - [CustomizeVoicesWindowController.SoundDidFinishPlaying]
//   - [CustomizeVoicesWindowController.SpeechSynthesizerDidFinishSpeaking]
//   - [CustomizeVoicesWindowController.StartStopPlayingSelection]
//   - [CustomizeVoicesWindowController.TableViewIsGroupRow]
//   - [CustomizeVoicesWindowController.TableViewShouldSelectRow]
//   - [CustomizeVoicesWindowController.TableViewViewForTableColumnRow]
//   - [CustomizeVoicesWindowController.TableViewSelectionDidChange]
//   - [CustomizeVoicesWindowController.DebugDescription]
//   - [CustomizeVoicesWindowController.Description]
//   - [CustomizeVoicesWindowController.Hash]
//   - [CustomizeVoicesWindowController.Superclass]
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController
type CustomizeVoicesWindowController struct {
	appkit.NSWindowController
}

// CustomizeVoicesWindowControllerFromID constructs a [CustomizeVoicesWindowController] from an objc.ID.
func CustomizeVoicesWindowControllerFromID(id objc.ID) CustomizeVoicesWindowController {
	return CustomizeVoicesWindowController{NSWindowController: appkit.NSWindowControllerFromID(id)}
}
// Ensure CustomizeVoicesWindowController implements ICustomizeVoicesWindowController.
var _ ICustomizeVoicesWindowController = CustomizeVoicesWindowController{}

// An interface definition for the [CustomizeVoicesWindowController] class.
//
// # Methods
//
//   - [ICustomizeVoicesWindowController._anyVoicePlaying]
//   - [ICustomizeVoicesWindowController._delayedPopUpUpdate]
//   - [ICustomizeVoicesWindowController._isSampleAvailableForVoiceObject]
//   - [ICustomizeVoicesWindowController._propagateCheckboxSelection]
//   - [ICustomizeVoicesWindowController._propagateDownloadCheckboxSelection]
//   - [ICustomizeVoicesWindowController._propagateDownloadVariantSelection]
//   - [ICustomizeVoicesWindowController._rebuildVoiceList]
//   - [ICustomizeVoicesWindowController._setRowDownloadCheckboxVoiceObjectIsSelected]
//   - [ICustomizeVoicesWindowController._setRowStatusFieldViewVoiceObjectIsSelected]
//   - [ICustomizeVoicesWindowController._shouldAllowRemovalOfVoiceObject]
//   - [ICustomizeVoicesWindowController._showPlayStopButtonAsPlaying]
//   - [ICustomizeVoicesWindowController._stopAndResetAllVoicePlaying]
//   - [ICustomizeVoicesWindowController._updateButtonStates]
//   - [ICustomizeVoicesWindowController._updateDisplayedVoicesUsingFilterString]
//   - [ICustomizeVoicesWindowController._updateRowDownloadStatus]
//   - [ICustomizeVoicesWindowController._voiceObjectForCurrentlySelectedRow]
//   - [ICustomizeVoicesWindowController.AcceptVoiceSelection]
//   - [ICustomizeVoicesWindowController.CancelVoiceSelection]
//   - [ICustomizeVoicesWindowController.NumberOfRowsInTableView]
//   - [ICustomizeVoicesWindowController.SearchFieldChanged]
//   - [ICustomizeVoicesWindowController.ShowSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved]
//   - [ICustomizeVoicesWindowController.SoundDidFinishPlaying]
//   - [ICustomizeVoicesWindowController.SpeechSynthesizerDidFinishSpeaking]
//   - [ICustomizeVoicesWindowController.StartStopPlayingSelection]
//   - [ICustomizeVoicesWindowController.TableViewIsGroupRow]
//   - [ICustomizeVoicesWindowController.TableViewShouldSelectRow]
//   - [ICustomizeVoicesWindowController.TableViewViewForTableColumnRow]
//   - [ICustomizeVoicesWindowController.TableViewSelectionDidChange]
//   - [ICustomizeVoicesWindowController.DebugDescription]
//   - [ICustomizeVoicesWindowController.Description]
//   - [ICustomizeVoicesWindowController.Hash]
//   - [ICustomizeVoicesWindowController.Superclass]
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController
type ICustomizeVoicesWindowController interface {
	appkit.INSWindowController

	// Topic: Methods

	_anyVoicePlaying() bool
	_delayedPopUpUpdate()
	_isSampleAvailableForVoiceObject(object objectivec.IObject) bool
	_propagateCheckboxSelection(selection objectivec.IObject)
	_propagateDownloadCheckboxSelection(selection objectivec.IObject)
	_propagateDownloadVariantSelection(selection objectivec.IObject)
	_rebuildVoiceList()
	_setRowDownloadCheckboxVoiceObjectIsSelected(checkbox objectivec.IObject, object objectivec.IObject, selected bool)
	_setRowStatusFieldViewVoiceObjectIsSelected(view objectivec.IObject, object objectivec.IObject, selected bool)
	_shouldAllowRemovalOfVoiceObject(object objectivec.IObject) bool
	_showPlayStopButtonAsPlaying(playing bool)
	_stopAndResetAllVoicePlaying()
	_updateButtonStates()
	_updateDisplayedVoicesUsingFilterString(string_ objectivec.IObject)
	_updateRowDownloadStatus()
	_voiceObjectForCurrentlySelectedRow() objectivec.IObject
	AcceptVoiceSelection(selection objectivec.IObject)
	CancelVoiceSelection(selection objectivec.IObject)
	NumberOfRowsInTableView(view objectivec.IObject) int64
	SearchFieldChanged(changed objectivec.IObject)
	ShowSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved(window objectivec.IObject, qualities bool, removed objectivec.IObject)
	SoundDidFinishPlaying(sound objectivec.IObject, playing bool)
	SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool)
	StartStopPlayingSelection(selection objectivec.IObject)
	TableViewIsGroupRow(view objectivec.IObject, row int) bool
	TableViewShouldSelectRow(view objectivec.IObject, row int64) bool
	TableViewViewForTableColumnRow(view objectivec.IObject, column objectivec.IObject, row int64) objectivec.IObject
	TableViewSelectionDidChange(change objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CustomizeVoicesWindowController) Init() CustomizeVoicesWindowController {
	rv := objc.Send[CustomizeVoicesWindowController](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CustomizeVoicesWindowController) Autorelease() CustomizeVoicesWindowController {
	rv := objc.Send[CustomizeVoicesWindowController](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomizeVoicesWindowController creates a new CustomizeVoicesWindowController instance.
func NewCustomizeVoicesWindowController() CustomizeVoicesWindowController {
	class := getCustomizeVoicesWindowControllerClass()
	rv := objc.Send[CustomizeVoicesWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_anyVoicePlaying
func (c CustomizeVoicesWindowController) _anyVoicePlaying() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("_anyVoicePlaying"))
	return rv
}

// AnyVoicePlaying is an exported wrapper for the private method _anyVoicePlaying.
func (c CustomizeVoicesWindowController) AnyVoicePlaying() bool {
	return c._anyVoicePlaying()
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_delayedPopUpUpdate
func (c CustomizeVoicesWindowController) _delayedPopUpUpdate() {
	objc.Send[objc.ID](c.ID, objc.Sel("_delayedPopUpUpdate"))
}

// DelayedPopUpUpdate is an exported wrapper for the private method _delayedPopUpUpdate.
func (c CustomizeVoicesWindowController) DelayedPopUpUpdate() {
	c._delayedPopUpUpdate()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_isSampleAvailableForVoiceObject:
func (c CustomizeVoicesWindowController) _isSampleAvailableForVoiceObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("_isSampleAvailableForVoiceObject:"), object)
	return rv
}

// IsSampleAvailableForVoiceObject is an exported wrapper for the private method _isSampleAvailableForVoiceObject.
func (c CustomizeVoicesWindowController) IsSampleAvailableForVoiceObject(object objectivec.IObject) bool {
	return c._isSampleAvailableForVoiceObject(object)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_propagateCheckboxSelection:
func (c CustomizeVoicesWindowController) _propagateCheckboxSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_propagateCheckboxSelection:"), selection)
}

// PropagateCheckboxSelection is an exported wrapper for the private method _propagateCheckboxSelection.
func (c CustomizeVoicesWindowController) PropagateCheckboxSelection(selection objectivec.IObject) {
	c._propagateCheckboxSelection(selection)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_propagateDownloadCheckboxSelection:
func (c CustomizeVoicesWindowController) _propagateDownloadCheckboxSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_propagateDownloadCheckboxSelection:"), selection)
}

// PropagateDownloadCheckboxSelection is an exported wrapper for the private method _propagateDownloadCheckboxSelection.
func (c CustomizeVoicesWindowController) PropagateDownloadCheckboxSelection(selection objectivec.IObject) {
	c._propagateDownloadCheckboxSelection(selection)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_propagateDownloadVariantSelection:
func (c CustomizeVoicesWindowController) _propagateDownloadVariantSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_propagateDownloadVariantSelection:"), selection)
}

// PropagateDownloadVariantSelection is an exported wrapper for the private method _propagateDownloadVariantSelection.
func (c CustomizeVoicesWindowController) PropagateDownloadVariantSelection(selection objectivec.IObject) {
	c._propagateDownloadVariantSelection(selection)
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_rebuildVoiceList
func (c CustomizeVoicesWindowController) _rebuildVoiceList() {
	objc.Send[objc.ID](c.ID, objc.Sel("_rebuildVoiceList"))
}

// RebuildVoiceList is an exported wrapper for the private method _rebuildVoiceList.
func (c CustomizeVoicesWindowController) RebuildVoiceList() {
	c._rebuildVoiceList()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_setRowDownloadCheckbox:voiceObject:isSelected:
func (c CustomizeVoicesWindowController) _setRowDownloadCheckboxVoiceObjectIsSelected(checkbox objectivec.IObject, object objectivec.IObject, selected bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("_setRowDownloadCheckbox:voiceObject:isSelected:"), checkbox, object, selected)
}

// SetRowDownloadCheckboxVoiceObjectIsSelected is an exported wrapper for the private method _setRowDownloadCheckboxVoiceObjectIsSelected.
func (c CustomizeVoicesWindowController) SetRowDownloadCheckboxVoiceObjectIsSelected(checkbox objectivec.IObject, object objectivec.IObject, selected bool) {
	c._setRowDownloadCheckboxVoiceObjectIsSelected(checkbox, object, selected)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_setRowStatusFieldView:voiceObject:isSelected:
func (c CustomizeVoicesWindowController) _setRowStatusFieldViewVoiceObjectIsSelected(view objectivec.IObject, object objectivec.IObject, selected bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("_setRowStatusFieldView:voiceObject:isSelected:"), view, object, selected)
}

// SetRowStatusFieldViewVoiceObjectIsSelected is an exported wrapper for the private method _setRowStatusFieldViewVoiceObjectIsSelected.
func (c CustomizeVoicesWindowController) SetRowStatusFieldViewVoiceObjectIsSelected(view objectivec.IObject, object objectivec.IObject, selected bool) {
	c._setRowStatusFieldViewVoiceObjectIsSelected(view, object, selected)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_shouldAllowRemovalOfVoiceObject:
func (c CustomizeVoicesWindowController) _shouldAllowRemovalOfVoiceObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("_shouldAllowRemovalOfVoiceObject:"), object)
	return rv
}

// ShouldAllowRemovalOfVoiceObject is an exported wrapper for the private method _shouldAllowRemovalOfVoiceObject.
func (c CustomizeVoicesWindowController) ShouldAllowRemovalOfVoiceObject(object objectivec.IObject) bool {
	return c._shouldAllowRemovalOfVoiceObject(object)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_showPlayStopButtonAsPlaying:
func (c CustomizeVoicesWindowController) _showPlayStopButtonAsPlaying(playing bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("_showPlayStopButtonAsPlaying:"), playing)
}

// ShowPlayStopButtonAsPlaying is an exported wrapper for the private method _showPlayStopButtonAsPlaying.
func (c CustomizeVoicesWindowController) ShowPlayStopButtonAsPlaying(playing bool) {
	c._showPlayStopButtonAsPlaying(playing)
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_stopAndResetAllVoicePlaying
func (c CustomizeVoicesWindowController) _stopAndResetAllVoicePlaying() {
	objc.Send[objc.ID](c.ID, objc.Sel("_stopAndResetAllVoicePlaying"))
}

// StopAndResetAllVoicePlaying is an exported wrapper for the private method _stopAndResetAllVoicePlaying.
func (c CustomizeVoicesWindowController) StopAndResetAllVoicePlaying() {
	c._stopAndResetAllVoicePlaying()
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_updateButtonStates
func (c CustomizeVoicesWindowController) _updateButtonStates() {
	objc.Send[objc.ID](c.ID, objc.Sel("_updateButtonStates"))
}

// UpdateButtonStates is an exported wrapper for the private method _updateButtonStates.
func (c CustomizeVoicesWindowController) UpdateButtonStates() {
	c._updateButtonStates()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_updateDisplayedVoicesUsingFilterString:
func (c CustomizeVoicesWindowController) _updateDisplayedVoicesUsingFilterString(string_ objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("_updateDisplayedVoicesUsingFilterString:"), string_)
}

// UpdateDisplayedVoicesUsingFilterString is an exported wrapper for the private method _updateDisplayedVoicesUsingFilterString.
func (c CustomizeVoicesWindowController) UpdateDisplayedVoicesUsingFilterString(string_ objectivec.IObject) {
	c._updateDisplayedVoicesUsingFilterString(string_)
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_updateRowDownloadStatus
func (c CustomizeVoicesWindowController) _updateRowDownloadStatus() {
	objc.Send[objc.ID](c.ID, objc.Sel("_updateRowDownloadStatus"))
}

// UpdateRowDownloadStatus is an exported wrapper for the private method _updateRowDownloadStatus.
func (c CustomizeVoicesWindowController) UpdateRowDownloadStatus() {
	c._updateRowDownloadStatus()
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/_voiceObjectForCurrentlySelectedRow
func (c CustomizeVoicesWindowController) _voiceObjectForCurrentlySelectedRow() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_voiceObjectForCurrentlySelectedRow"))
	return objectivec.Object{ID: rv}
}

// VoiceObjectForCurrentlySelectedRow is an exported wrapper for the private method _voiceObjectForCurrentlySelectedRow.
func (c CustomizeVoicesWindowController) VoiceObjectForCurrentlySelectedRow() objectivec.IObject {
	return c._voiceObjectForCurrentlySelectedRow()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/acceptVoiceSelection:
func (c CustomizeVoicesWindowController) AcceptVoiceSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("acceptVoiceSelection:"), selection)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/cancelVoiceSelection:
func (c CustomizeVoicesWindowController) CancelVoiceSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("cancelVoiceSelection:"), selection)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/numberOfRowsInTableView:
func (c CustomizeVoicesWindowController) NumberOfRowsInTableView(view objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("numberOfRowsInTableView:"), view)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/searchFieldChanged:
func (c CustomizeVoicesWindowController) SearchFieldChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("searchFieldChanged:"), changed)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/showSheetForWindow:showIndividualVoiceQualities:voiceIdentifiersNotToBeRemoved:
func (c CustomizeVoicesWindowController) ShowSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved(window objectivec.IObject, qualities bool, removed objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("showSheetForWindow:showIndividualVoiceQualities:voiceIdentifiersNotToBeRemoved:"), window, qualities, removed)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/sound:didFinishPlaying:
func (c CustomizeVoicesWindowController) SoundDidFinishPlaying(sound objectivec.IObject, playing bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("sound:didFinishPlaying:"), sound, playing)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/speechSynthesizer:didFinishSpeaking:
func (c CustomizeVoicesWindowController) SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), synthesizer, speaking)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/startStopPlayingSelection:
func (c CustomizeVoicesWindowController) StartStopPlayingSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("startStopPlayingSelection:"), selection)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/tableView:isGroupRow:
func (c CustomizeVoicesWindowController) TableViewIsGroupRow(view objectivec.IObject, row int) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("tableView:isGroupRow:"), view, row)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/tableView:shouldSelectRow:
func (c CustomizeVoicesWindowController) TableViewShouldSelectRow(view objectivec.IObject, row int64) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("tableView:shouldSelectRow:"), view, row)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/tableView:viewForTableColumn:row:
func (c CustomizeVoicesWindowController) TableViewViewForTableColumnRow(view objectivec.IObject, column objectivec.IObject, row int64) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("tableView:viewForTableColumn:row:"), view, column, row)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/tableViewSelectionDidChange:
func (c CustomizeVoicesWindowController) TableViewSelectionDidChange(change objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("tableViewSelectionDidChange:"), change)
}

// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/debugDescription
func (c CustomizeVoicesWindowController) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/description
func (c CustomizeVoicesWindowController) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/hash
func (c CustomizeVoicesWindowController) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/CustomizeVoicesWindowController/superclass
func (c CustomizeVoicesWindowController) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

