// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[CustomizeVoicesWindowController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CustomizeVoicesWindowController) Init() CustomizeVoicesWindowController {
	rv := objc.SendIfResponds[CustomizeVoicesWindowController](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CustomizeVoicesWindowController) Autorelease() CustomizeVoicesWindowController {
	rv := objc.SendIfResponds[CustomizeVoicesWindowController](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomizeVoicesWindowController creates a new CustomizeVoicesWindowController instance.
func NewCustomizeVoicesWindowController() CustomizeVoicesWindowController {
	class := getCustomizeVoicesWindowControllerClass()
	rv := objc.SendIfResponds[CustomizeVoicesWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CustomizeVoicesWindowController) _anyVoicePlaying() bool {
	rv := objc.SendIfResponds[bool](c.ID, objc.Sel("_anyVoicePlaying"))
	return rv
}

// AnyVoicePlaying is an exported wrapper for the private method _anyVoicePlaying.
func (c CustomizeVoicesWindowController) AnyVoicePlaying() (bool, error) {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_anyVoicePlaying")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_anyVoicePlaying"}
		return false, err
	}
	return c._anyVoicePlaying(), nil
}

// CanAnyVoicePlaying reports whether the receiver responds to the private selector _anyVoicePlaying.
func (c CustomizeVoicesWindowController) CanAnyVoicePlaying() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_anyVoicePlaying"))
}
func (c CustomizeVoicesWindowController) _delayedPopUpUpdate() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_delayedPopUpUpdate"))
}

// DelayedPopUpUpdate is an exported wrapper for the private method _delayedPopUpUpdate.
func (c CustomizeVoicesWindowController) DelayedPopUpUpdate() error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_delayedPopUpUpdate")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_delayedPopUpUpdate"}
		return err
	}
	c._delayedPopUpUpdate()
	return nil
}

// CanDelayedPopUpUpdate reports whether the receiver responds to the private selector _delayedPopUpUpdate.
func (c CustomizeVoicesWindowController) CanDelayedPopUpUpdate() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_delayedPopUpUpdate"))
}
func (c CustomizeVoicesWindowController) _isSampleAvailableForVoiceObject(object objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](c.ID, objc.Sel("_isSampleAvailableForVoiceObject:"), object)
	return rv
}

// IsSampleAvailableForVoiceObject is an exported wrapper for the private method _isSampleAvailableForVoiceObject.
func (c CustomizeVoicesWindowController) IsSampleAvailableForVoiceObject(object objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_isSampleAvailableForVoiceObject:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isSampleAvailableForVoiceObject:"}
		return false, err
	}
	return c._isSampleAvailableForVoiceObject(object), nil
}

// CanIsSampleAvailableForVoiceObject reports whether the receiver responds to the private selector _isSampleAvailableForVoiceObject:.
func (c CustomizeVoicesWindowController) CanIsSampleAvailableForVoiceObject() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_isSampleAvailableForVoiceObject:"))
}
func (c CustomizeVoicesWindowController) _propagateCheckboxSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_propagateCheckboxSelection:"), selection)
}

// PropagateCheckboxSelection is an exported wrapper for the private method _propagateCheckboxSelection.
func (c CustomizeVoicesWindowController) PropagateCheckboxSelection(selection objectivec.IObject) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_propagateCheckboxSelection:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_propagateCheckboxSelection:"}
		return err
	}
	c._propagateCheckboxSelection(selection)
	return nil
}

// CanPropagateCheckboxSelection reports whether the receiver responds to the private selector _propagateCheckboxSelection:.
func (c CustomizeVoicesWindowController) CanPropagateCheckboxSelection() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_propagateCheckboxSelection:"))
}
func (c CustomizeVoicesWindowController) _propagateDownloadCheckboxSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_propagateDownloadCheckboxSelection:"), selection)
}

// PropagateDownloadCheckboxSelection is an exported wrapper for the private method _propagateDownloadCheckboxSelection.
func (c CustomizeVoicesWindowController) PropagateDownloadCheckboxSelection(selection objectivec.IObject) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_propagateDownloadCheckboxSelection:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_propagateDownloadCheckboxSelection:"}
		return err
	}
	c._propagateDownloadCheckboxSelection(selection)
	return nil
}

// CanPropagateDownloadCheckboxSelection reports whether the receiver responds to the private selector _propagateDownloadCheckboxSelection:.
func (c CustomizeVoicesWindowController) CanPropagateDownloadCheckboxSelection() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_propagateDownloadCheckboxSelection:"))
}
func (c CustomizeVoicesWindowController) _propagateDownloadVariantSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_propagateDownloadVariantSelection:"), selection)
}

// PropagateDownloadVariantSelection is an exported wrapper for the private method _propagateDownloadVariantSelection.
func (c CustomizeVoicesWindowController) PropagateDownloadVariantSelection(selection objectivec.IObject) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_propagateDownloadVariantSelection:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_propagateDownloadVariantSelection:"}
		return err
	}
	c._propagateDownloadVariantSelection(selection)
	return nil
}

// CanPropagateDownloadVariantSelection reports whether the receiver responds to the private selector _propagateDownloadVariantSelection:.
func (c CustomizeVoicesWindowController) CanPropagateDownloadVariantSelection() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_propagateDownloadVariantSelection:"))
}
func (c CustomizeVoicesWindowController) _rebuildVoiceList() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_rebuildVoiceList"))
}

// RebuildVoiceList is an exported wrapper for the private method _rebuildVoiceList.
func (c CustomizeVoicesWindowController) RebuildVoiceList() error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_rebuildVoiceList")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rebuildVoiceList"}
		return err
	}
	c._rebuildVoiceList()
	return nil
}

// CanRebuildVoiceList reports whether the receiver responds to the private selector _rebuildVoiceList.
func (c CustomizeVoicesWindowController) CanRebuildVoiceList() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_rebuildVoiceList"))
}
func (c CustomizeVoicesWindowController) _setRowDownloadCheckboxVoiceObjectIsSelected(checkbox objectivec.IObject, object objectivec.IObject, selected bool) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_setRowDownloadCheckbox:voiceObject:isSelected:"), checkbox, object, selected)
}

// SetRowDownloadCheckboxVoiceObjectIsSelected is an exported wrapper for the private method _setRowDownloadCheckboxVoiceObjectIsSelected.
func (c CustomizeVoicesWindowController) SetRowDownloadCheckboxVoiceObjectIsSelected(checkbox objectivec.IObject, object objectivec.IObject, selected bool) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_setRowDownloadCheckbox:voiceObject:isSelected:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRowDownloadCheckbox:voiceObject:isSelected:"}
		return err
	}
	c._setRowDownloadCheckboxVoiceObjectIsSelected(checkbox, object, selected)
	return nil
}

// CanSetRowDownloadCheckboxVoiceObjectIsSelected reports whether the receiver responds to the private selector _setRowDownloadCheckbox:voiceObject:isSelected:.
func (c CustomizeVoicesWindowController) CanSetRowDownloadCheckboxVoiceObjectIsSelected() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_setRowDownloadCheckbox:voiceObject:isSelected:"))
}
func (c CustomizeVoicesWindowController) _setRowStatusFieldViewVoiceObjectIsSelected(view objectivec.IObject, object objectivec.IObject, selected bool) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_setRowStatusFieldView:voiceObject:isSelected:"), view, object, selected)
}

// SetRowStatusFieldViewVoiceObjectIsSelected is an exported wrapper for the private method _setRowStatusFieldViewVoiceObjectIsSelected.
func (c CustomizeVoicesWindowController) SetRowStatusFieldViewVoiceObjectIsSelected(view objectivec.IObject, object objectivec.IObject, selected bool) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_setRowStatusFieldView:voiceObject:isSelected:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRowStatusFieldView:voiceObject:isSelected:"}
		return err
	}
	c._setRowStatusFieldViewVoiceObjectIsSelected(view, object, selected)
	return nil
}

// CanSetRowStatusFieldViewVoiceObjectIsSelected reports whether the receiver responds to the private selector _setRowStatusFieldView:voiceObject:isSelected:.
func (c CustomizeVoicesWindowController) CanSetRowStatusFieldViewVoiceObjectIsSelected() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_setRowStatusFieldView:voiceObject:isSelected:"))
}
func (c CustomizeVoicesWindowController) _shouldAllowRemovalOfVoiceObject(object objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](c.ID, objc.Sel("_shouldAllowRemovalOfVoiceObject:"), object)
	return rv
}

// ShouldAllowRemovalOfVoiceObject is an exported wrapper for the private method _shouldAllowRemovalOfVoiceObject.
func (c CustomizeVoicesWindowController) ShouldAllowRemovalOfVoiceObject(object objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_shouldAllowRemovalOfVoiceObject:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_shouldAllowRemovalOfVoiceObject:"}
		return false, err
	}
	return c._shouldAllowRemovalOfVoiceObject(object), nil
}

// CanShouldAllowRemovalOfVoiceObject reports whether the receiver responds to the private selector _shouldAllowRemovalOfVoiceObject:.
func (c CustomizeVoicesWindowController) CanShouldAllowRemovalOfVoiceObject() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_shouldAllowRemovalOfVoiceObject:"))
}
func (c CustomizeVoicesWindowController) _showPlayStopButtonAsPlaying(playing bool) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_showPlayStopButtonAsPlaying:"), playing)
}

// ShowPlayStopButtonAsPlaying is an exported wrapper for the private method _showPlayStopButtonAsPlaying.
func (c CustomizeVoicesWindowController) ShowPlayStopButtonAsPlaying(playing bool) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_showPlayStopButtonAsPlaying:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_showPlayStopButtonAsPlaying:"}
		return err
	}
	c._showPlayStopButtonAsPlaying(playing)
	return nil
}

// CanShowPlayStopButtonAsPlaying reports whether the receiver responds to the private selector _showPlayStopButtonAsPlaying:.
func (c CustomizeVoicesWindowController) CanShowPlayStopButtonAsPlaying() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_showPlayStopButtonAsPlaying:"))
}
func (c CustomizeVoicesWindowController) _stopAndResetAllVoicePlaying() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_stopAndResetAllVoicePlaying"))
}

// StopAndResetAllVoicePlaying is an exported wrapper for the private method _stopAndResetAllVoicePlaying.
func (c CustomizeVoicesWindowController) StopAndResetAllVoicePlaying() error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_stopAndResetAllVoicePlaying")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_stopAndResetAllVoicePlaying"}
		return err
	}
	c._stopAndResetAllVoicePlaying()
	return nil
}

// CanStopAndResetAllVoicePlaying reports whether the receiver responds to the private selector _stopAndResetAllVoicePlaying.
func (c CustomizeVoicesWindowController) CanStopAndResetAllVoicePlaying() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_stopAndResetAllVoicePlaying"))
}
func (c CustomizeVoicesWindowController) _updateButtonStates() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_updateButtonStates"))
}

// UpdateButtonStates is an exported wrapper for the private method _updateButtonStates.
func (c CustomizeVoicesWindowController) UpdateButtonStates() error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_updateButtonStates")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateButtonStates"}
		return err
	}
	c._updateButtonStates()
	return nil
}

// CanUpdateButtonStates reports whether the receiver responds to the private selector _updateButtonStates.
func (c CustomizeVoicesWindowController) CanUpdateButtonStates() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_updateButtonStates"))
}
func (c CustomizeVoicesWindowController) _updateDisplayedVoicesUsingFilterString(string_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_updateDisplayedVoicesUsingFilterString:"), string_)
}

// UpdateDisplayedVoicesUsingFilterString is an exported wrapper for the private method _updateDisplayedVoicesUsingFilterString.
func (c CustomizeVoicesWindowController) UpdateDisplayedVoicesUsingFilterString(string_ objectivec.IObject) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_updateDisplayedVoicesUsingFilterString:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateDisplayedVoicesUsingFilterString:"}
		return err
	}
	c._updateDisplayedVoicesUsingFilterString(string_)
	return nil
}

// CanUpdateDisplayedVoicesUsingFilterString reports whether the receiver responds to the private selector _updateDisplayedVoicesUsingFilterString:.
func (c CustomizeVoicesWindowController) CanUpdateDisplayedVoicesUsingFilterString() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_updateDisplayedVoicesUsingFilterString:"))
}
func (c CustomizeVoicesWindowController) _updateRowDownloadStatus() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_updateRowDownloadStatus"))
}

// UpdateRowDownloadStatus is an exported wrapper for the private method _updateRowDownloadStatus.
func (c CustomizeVoicesWindowController) UpdateRowDownloadStatus() error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_updateRowDownloadStatus")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateRowDownloadStatus"}
		return err
	}
	c._updateRowDownloadStatus()
	return nil
}

// CanUpdateRowDownloadStatus reports whether the receiver responds to the private selector _updateRowDownloadStatus.
func (c CustomizeVoicesWindowController) CanUpdateRowDownloadStatus() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_updateRowDownloadStatus"))
}
func (c CustomizeVoicesWindowController) _voiceObjectForCurrentlySelectedRow() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("_voiceObjectForCurrentlySelectedRow"))
	return objectivec.Object{ID: rv}
}

// VoiceObjectForCurrentlySelectedRow is an exported wrapper for the private method _voiceObjectForCurrentlySelectedRow.
func (c CustomizeVoicesWindowController) VoiceObjectForCurrentlySelectedRow() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_voiceObjectForCurrentlySelectedRow")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_voiceObjectForCurrentlySelectedRow"}
		return nil, err
	}
	return c._voiceObjectForCurrentlySelectedRow(), nil
}

// CanVoiceObjectForCurrentlySelectedRow reports whether the receiver responds to the private selector _voiceObjectForCurrentlySelectedRow.
func (c CustomizeVoicesWindowController) CanVoiceObjectForCurrentlySelectedRow() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_voiceObjectForCurrentlySelectedRow"))
}
func (c CustomizeVoicesWindowController) AcceptVoiceSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("acceptVoiceSelection:"), selection)
}
func (c CustomizeVoicesWindowController) CancelVoiceSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("cancelVoiceSelection:"), selection)
}
func (c CustomizeVoicesWindowController) NumberOfRowsInTableView(view objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("numberOfRowsInTableView:"), view)
	return rv
}
func (c CustomizeVoicesWindowController) SearchFieldChanged(changed objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("searchFieldChanged:"), changed)
}
func (c CustomizeVoicesWindowController) ShowSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved(window objectivec.IObject, qualities bool, removed objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("showSheetForWindow:showIndividualVoiceQualities:voiceIdentifiersNotToBeRemoved:"), window, qualities, removed)
}
func (c CustomizeVoicesWindowController) SoundDidFinishPlaying(sound objectivec.IObject, playing bool) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("sound:didFinishPlaying:"), sound, playing)
}
func (c CustomizeVoicesWindowController) SpeechSynthesizerDidFinishSpeaking(synthesizer objectivec.IObject, speaking bool) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), synthesizer, speaking)
}
func (c CustomizeVoicesWindowController) StartStopPlayingSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("startStopPlayingSelection:"), selection)
}
func (c CustomizeVoicesWindowController) TableViewIsGroupRow(view objectivec.IObject, row int) bool {
	rv := objc.SendIfResponds[bool](c.ID, objc.Sel("tableView:isGroupRow:"), view, row)
	return rv
}
func (c CustomizeVoicesWindowController) TableViewShouldSelectRow(view objectivec.IObject, row int64) bool {
	rv := objc.SendIfResponds[bool](c.ID, objc.Sel("tableView:shouldSelectRow:"), view, row)
	return rv
}
func (c CustomizeVoicesWindowController) TableViewViewForTableColumnRow(view objectivec.IObject, column objectivec.IObject, row int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("tableView:viewForTableColumn:row:"), view, column, row)
	return objectivec.Object{ID: rv}
}
func (c CustomizeVoicesWindowController) TableViewSelectionDidChange(change objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("tableViewSelectionDidChange:"), change)
}

func (c CustomizeVoicesWindowController) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CustomizeVoicesWindowController) Description() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CustomizeVoicesWindowController) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CustomizeVoicesWindowController) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
