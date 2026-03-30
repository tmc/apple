// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOCustomizeSRLanguagesWindowController] class.
var (
	_SOCustomizeSRLanguagesWindowControllerClass     SOCustomizeSRLanguagesWindowControllerClass
	_SOCustomizeSRLanguagesWindowControllerClassOnce sync.Once
)

func getSOCustomizeSRLanguagesWindowControllerClass() SOCustomizeSRLanguagesWindowControllerClass {
	_SOCustomizeSRLanguagesWindowControllerClassOnce.Do(func() {
		_SOCustomizeSRLanguagesWindowControllerClass = SOCustomizeSRLanguagesWindowControllerClass{class: objc.GetClass("SOCustomizeSRLanguagesWindowController")}
	})
	return _SOCustomizeSRLanguagesWindowControllerClass
}

// GetSOCustomizeSRLanguagesWindowControllerClass returns the class object for SOCustomizeSRLanguagesWindowController.
func GetSOCustomizeSRLanguagesWindowControllerClass() SOCustomizeSRLanguagesWindowControllerClass {
	return getSOCustomizeSRLanguagesWindowControllerClass()
}

type SOCustomizeSRLanguagesWindowControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOCustomizeSRLanguagesWindowControllerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOCustomizeSRLanguagesWindowControllerClass) Alloc() SOCustomizeSRLanguagesWindowController {
	rv := objc.Send[SOCustomizeSRLanguagesWindowController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOCustomizeSRLanguagesWindowController._propagateCheckboxSelection]
//   - [SOCustomizeSRLanguagesWindowController._propagateDownloadVariantSelection]
//   - [SOCustomizeSRLanguagesWindowController._rebuildList]
//   - [SOCustomizeSRLanguagesWindowController._setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected]
//   - [SOCustomizeSRLanguagesWindowController._updateButtonStatesOnlyIfDownloadRequired]
//   - [SOCustomizeSRLanguagesWindowController._updateDisplayUsingFilterString]
//   - [SOCustomizeSRLanguagesWindowController._updateRowDownloadStatus]
//   - [SOCustomizeSRLanguagesWindowController.AcceptSelection]
//   - [SOCustomizeSRLanguagesWindowController.CancelSelection]
//   - [SOCustomizeSRLanguagesWindowController.NumberOfRowsInTableView]
//   - [SOCustomizeSRLanguagesWindowController.SearchFieldChanged]
//   - [SOCustomizeSRLanguagesWindowController.ShowSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems]
//   - [SOCustomizeSRLanguagesWindowController.TableViewIsGroupRow]
//   - [SOCustomizeSRLanguagesWindowController.TableViewShouldSelectRow]
//   - [SOCustomizeSRLanguagesWindowController.TableViewViewForTableColumnRow]
//   - [SOCustomizeSRLanguagesWindowController.TableViewSelectionDidChange]
//   - [SOCustomizeSRLanguagesWindowController.DebugDescription]
//   - [SOCustomizeSRLanguagesWindowController.Description]
//   - [SOCustomizeSRLanguagesWindowController.Hash]
//   - [SOCustomizeSRLanguagesWindowController.Superclass]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController
type SOCustomizeSRLanguagesWindowController struct {
	appkit.NSWindowController
}

// SOCustomizeSRLanguagesWindowControllerFromID constructs a [SOCustomizeSRLanguagesWindowController] from an objc.ID.
func SOCustomizeSRLanguagesWindowControllerFromID(id objc.ID) SOCustomizeSRLanguagesWindowController {
	return SOCustomizeSRLanguagesWindowController{NSWindowController: appkit.NSWindowControllerFromID(id)}
}

// Ensure SOCustomizeSRLanguagesWindowController implements ISOCustomizeSRLanguagesWindowController.
var _ ISOCustomizeSRLanguagesWindowController = SOCustomizeSRLanguagesWindowController{}

// An interface definition for the [SOCustomizeSRLanguagesWindowController] class.
//
// # Methods
//
//   - [ISOCustomizeSRLanguagesWindowController._propagateCheckboxSelection]
//   - [ISOCustomizeSRLanguagesWindowController._propagateDownloadVariantSelection]
//   - [ISOCustomizeSRLanguagesWindowController._rebuildList]
//   - [ISOCustomizeSRLanguagesWindowController._setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected]
//   - [ISOCustomizeSRLanguagesWindowController._updateButtonStatesOnlyIfDownloadRequired]
//   - [ISOCustomizeSRLanguagesWindowController._updateDisplayUsingFilterString]
//   - [ISOCustomizeSRLanguagesWindowController._updateRowDownloadStatus]
//   - [ISOCustomizeSRLanguagesWindowController.AcceptSelection]
//   - [ISOCustomizeSRLanguagesWindowController.CancelSelection]
//   - [ISOCustomizeSRLanguagesWindowController.NumberOfRowsInTableView]
//   - [ISOCustomizeSRLanguagesWindowController.SearchFieldChanged]
//   - [ISOCustomizeSRLanguagesWindowController.ShowSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems]
//   - [ISOCustomizeSRLanguagesWindowController.TableViewIsGroupRow]
//   - [ISOCustomizeSRLanguagesWindowController.TableViewShouldSelectRow]
//   - [ISOCustomizeSRLanguagesWindowController.TableViewViewForTableColumnRow]
//   - [ISOCustomizeSRLanguagesWindowController.TableViewSelectionDidChange]
//   - [ISOCustomizeSRLanguagesWindowController.DebugDescription]
//   - [ISOCustomizeSRLanguagesWindowController.Description]
//   - [ISOCustomizeSRLanguagesWindowController.Hash]
//   - [ISOCustomizeSRLanguagesWindowController.Superclass]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController
type ISOCustomizeSRLanguagesWindowController interface {
	appkit.INSWindowController

	// Topic: Methods

	_propagateCheckboxSelection(selection objectivec.IObject)
	_propagateDownloadVariantSelection(selection objectivec.IObject)
	_rebuildList()
	_setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view objectivec.IObject, button objectivec.IObject, item objectivec.IObject, selected bool)
	_updateButtonStatesOnlyIfDownloadRequired(required bool)
	_updateDisplayUsingFilterString(string_ objectivec.IObject)
	_updateRowDownloadStatus()
	AcceptSelection(selection objectivec.IObject)
	CancelSelection(selection objectivec.IObject)
	NumberOfRowsInTableView(view objectivec.IObject) int64
	SearchFieldChanged(changed objectivec.IObject)
	ShowSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems(window objectivec.IObject, identifiers objectivec.IObject, identifier objectivec.IObject, downloads bool, items bool)
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
func (s SOCustomizeSRLanguagesWindowController) Init() SOCustomizeSRLanguagesWindowController {
	rv := objc.Send[SOCustomizeSRLanguagesWindowController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOCustomizeSRLanguagesWindowController) Autorelease() SOCustomizeSRLanguagesWindowController {
	rv := objc.Send[SOCustomizeSRLanguagesWindowController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOCustomizeSRLanguagesWindowController creates a new SOCustomizeSRLanguagesWindowController instance.
func NewSOCustomizeSRLanguagesWindowController() SOCustomizeSRLanguagesWindowController {
	class := getSOCustomizeSRLanguagesWindowControllerClass()
	rv := objc.Send[SOCustomizeSRLanguagesWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_propagateCheckboxSelection:
func (s SOCustomizeSRLanguagesWindowController) _propagateCheckboxSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_propagateCheckboxSelection:"), selection)
}

// PropagateCheckboxSelection is an exported wrapper for the private method _propagateCheckboxSelection.
func (s SOCustomizeSRLanguagesWindowController) PropagateCheckboxSelection(selection objectivec.IObject) {
	s._propagateCheckboxSelection(selection)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_propagateDownloadVariantSelection:
func (s SOCustomizeSRLanguagesWindowController) _propagateDownloadVariantSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_propagateDownloadVariantSelection:"), selection)
}

// PropagateDownloadVariantSelection is an exported wrapper for the private method _propagateDownloadVariantSelection.
func (s SOCustomizeSRLanguagesWindowController) PropagateDownloadVariantSelection(selection objectivec.IObject) {
	s._propagateDownloadVariantSelection(selection)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_rebuildList
func (s SOCustomizeSRLanguagesWindowController) _rebuildList() {
	objc.Send[objc.ID](s.ID, objc.Sel("_rebuildList"))
}

// RebuildList is an exported wrapper for the private method _rebuildList.
func (s SOCustomizeSRLanguagesWindowController) RebuildList() {
	s._rebuildList()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:
func (s SOCustomizeSRLanguagesWindowController) _setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view objectivec.IObject, button objectivec.IObject, item objectivec.IObject, selected bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:"), view, button, item, selected)
}

// SetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected is an exported wrapper for the private method _setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected.
func (s SOCustomizeSRLanguagesWindowController) SetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view objectivec.IObject, button objectivec.IObject, item objectivec.IObject, selected bool) {
	s._setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view, button, item, selected)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_updateButtonStatesOnlyIfDownloadRequired:
func (s SOCustomizeSRLanguagesWindowController) _updateButtonStatesOnlyIfDownloadRequired(required bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateButtonStatesOnlyIfDownloadRequired:"), required)
}

// UpdateButtonStatesOnlyIfDownloadRequired is an exported wrapper for the private method _updateButtonStatesOnlyIfDownloadRequired.
func (s SOCustomizeSRLanguagesWindowController) UpdateButtonStatesOnlyIfDownloadRequired(required bool) {
	s._updateButtonStatesOnlyIfDownloadRequired(required)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_updateDisplayUsingFilterString:
func (s SOCustomizeSRLanguagesWindowController) _updateDisplayUsingFilterString(string_ objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateDisplayUsingFilterString:"), string_)
}

// UpdateDisplayUsingFilterString is an exported wrapper for the private method _updateDisplayUsingFilterString.
func (s SOCustomizeSRLanguagesWindowController) UpdateDisplayUsingFilterString(string_ objectivec.IObject) {
	s._updateDisplayUsingFilterString(string_)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/_updateRowDownloadStatus
func (s SOCustomizeSRLanguagesWindowController) _updateRowDownloadStatus() {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateRowDownloadStatus"))
}

// UpdateRowDownloadStatus is an exported wrapper for the private method _updateRowDownloadStatus.
func (s SOCustomizeSRLanguagesWindowController) UpdateRowDownloadStatus() {
	s._updateRowDownloadStatus()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/acceptSelection:
func (s SOCustomizeSRLanguagesWindowController) AcceptSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("acceptSelection:"), selection)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/cancelSelection:
func (s SOCustomizeSRLanguagesWindowController) CancelSelection(selection objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelSelection:"), selection)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/numberOfRowsInTableView:
func (s SOCustomizeSRLanguagesWindowController) NumberOfRowsInTableView(view objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("numberOfRowsInTableView:"), view)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/searchFieldChanged:
func (s SOCustomizeSRLanguagesWindowController) SearchFieldChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("searchFieldChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/showSheetForWindow:networkSupportedLocaleIdentifiers:requiredLocaleIdentifier:supportDownloads:showOnlyNetworkSupportedItems:
func (s SOCustomizeSRLanguagesWindowController) ShowSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems(window objectivec.IObject, identifiers objectivec.IObject, identifier objectivec.IObject, downloads bool, items bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("showSheetForWindow:networkSupportedLocaleIdentifiers:requiredLocaleIdentifier:supportDownloads:showOnlyNetworkSupportedItems:"), window, identifiers, identifier, downloads, items)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/tableView:isGroupRow:
func (s SOCustomizeSRLanguagesWindowController) TableViewIsGroupRow(view objectivec.IObject, row int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("tableView:isGroupRow:"), view, row)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/tableView:shouldSelectRow:
func (s SOCustomizeSRLanguagesWindowController) TableViewShouldSelectRow(view objectivec.IObject, row int64) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("tableView:shouldSelectRow:"), view, row)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/tableView:viewForTableColumn:row:
func (s SOCustomizeSRLanguagesWindowController) TableViewViewForTableColumnRow(view objectivec.IObject, column objectivec.IObject, row int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("tableView:viewForTableColumn:row:"), view, column, row)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/tableViewSelectionDidChange:
func (s SOCustomizeSRLanguagesWindowController) TableViewSelectionDidChange(change objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("tableViewSelectionDidChange:"), change)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/debugDescription
func (s SOCustomizeSRLanguagesWindowController) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/description
func (s SOCustomizeSRLanguagesWindowController) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/hash
func (s SOCustomizeSRLanguagesWindowController) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOCustomizeSRLanguagesWindowController/superclass
func (s SOCustomizeSRLanguagesWindowController) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
