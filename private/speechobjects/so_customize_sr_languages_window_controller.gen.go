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
	rv := objc.SendIfResponds[SOCustomizeSRLanguagesWindowController](objc.ID(sc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (s SOCustomizeSRLanguagesWindowController) Init() SOCustomizeSRLanguagesWindowController {
	rv := objc.SendIfResponds[SOCustomizeSRLanguagesWindowController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOCustomizeSRLanguagesWindowController) Autorelease() SOCustomizeSRLanguagesWindowController {
	rv := objc.SendIfResponds[SOCustomizeSRLanguagesWindowController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOCustomizeSRLanguagesWindowController creates a new SOCustomizeSRLanguagesWindowController instance.
func NewSOCustomizeSRLanguagesWindowController() SOCustomizeSRLanguagesWindowController {
	class := getSOCustomizeSRLanguagesWindowControllerClass()
	rv := objc.SendIfResponds[SOCustomizeSRLanguagesWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SOCustomizeSRLanguagesWindowController) _propagateCheckboxSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_propagateCheckboxSelection:"), selection)
}

// PropagateCheckboxSelection is an exported wrapper for the private method _propagateCheckboxSelection.
func (s SOCustomizeSRLanguagesWindowController) PropagateCheckboxSelection(selection objectivec.IObject) error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_propagateCheckboxSelection:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_propagateCheckboxSelection:"}
		return err
	}
	s._propagateCheckboxSelection(selection)
	return nil
}

// CanPropagateCheckboxSelection reports whether the receiver responds to the private selector _propagateCheckboxSelection:.
func (s SOCustomizeSRLanguagesWindowController) CanPropagateCheckboxSelection() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_propagateCheckboxSelection:"))
}
func (s SOCustomizeSRLanguagesWindowController) _propagateDownloadVariantSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_propagateDownloadVariantSelection:"), selection)
}

// PropagateDownloadVariantSelection is an exported wrapper for the private method _propagateDownloadVariantSelection.
func (s SOCustomizeSRLanguagesWindowController) PropagateDownloadVariantSelection(selection objectivec.IObject) error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_propagateDownloadVariantSelection:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_propagateDownloadVariantSelection:"}
		return err
	}
	s._propagateDownloadVariantSelection(selection)
	return nil
}

// CanPropagateDownloadVariantSelection reports whether the receiver responds to the private selector _propagateDownloadVariantSelection:.
func (s SOCustomizeSRLanguagesWindowController) CanPropagateDownloadVariantSelection() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_propagateDownloadVariantSelection:"))
}
func (s SOCustomizeSRLanguagesWindowController) _rebuildList() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_rebuildList"))
}

// RebuildList is an exported wrapper for the private method _rebuildList.
func (s SOCustomizeSRLanguagesWindowController) RebuildList() error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_rebuildList")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rebuildList"}
		return err
	}
	s._rebuildList()
	return nil
}

// CanRebuildList reports whether the receiver responds to the private selector _rebuildList.
func (s SOCustomizeSRLanguagesWindowController) CanRebuildList() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_rebuildList"))
}
func (s SOCustomizeSRLanguagesWindowController) _setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view objectivec.IObject, button objectivec.IObject, item objectivec.IObject, selected bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:"), view, button, item, selected)
}

// SetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected is an exported wrapper for the private method _setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected.
func (s SOCustomizeSRLanguagesWindowController) SetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view objectivec.IObject, button objectivec.IObject, item objectivec.IObject, selected bool) error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:"}
		return err
	}
	s._setRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected(view, button, item, selected)
	return nil
}

// CanSetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected reports whether the receiver responds to the private selector _setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:.
func (s SOCustomizeSRLanguagesWindowController) CanSetRowStatusFieldViewVariantPopUpButtonSpeechItemIsSelected() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_setRowStatusFieldView:variantPopUpButton:speechItem:isSelected:"))
}
func (s SOCustomizeSRLanguagesWindowController) _updateButtonStatesOnlyIfDownloadRequired(required bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_updateButtonStatesOnlyIfDownloadRequired:"), required)
}

// UpdateButtonStatesOnlyIfDownloadRequired is an exported wrapper for the private method _updateButtonStatesOnlyIfDownloadRequired.
func (s SOCustomizeSRLanguagesWindowController) UpdateButtonStatesOnlyIfDownloadRequired(required bool) error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_updateButtonStatesOnlyIfDownloadRequired:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateButtonStatesOnlyIfDownloadRequired:"}
		return err
	}
	s._updateButtonStatesOnlyIfDownloadRequired(required)
	return nil
}

// CanUpdateButtonStatesOnlyIfDownloadRequired reports whether the receiver responds to the private selector _updateButtonStatesOnlyIfDownloadRequired:.
func (s SOCustomizeSRLanguagesWindowController) CanUpdateButtonStatesOnlyIfDownloadRequired() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_updateButtonStatesOnlyIfDownloadRequired:"))
}
func (s SOCustomizeSRLanguagesWindowController) _updateDisplayUsingFilterString(string_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_updateDisplayUsingFilterString:"), string_)
}

// UpdateDisplayUsingFilterString is an exported wrapper for the private method _updateDisplayUsingFilterString.
func (s SOCustomizeSRLanguagesWindowController) UpdateDisplayUsingFilterString(string_ objectivec.IObject) error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_updateDisplayUsingFilterString:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateDisplayUsingFilterString:"}
		return err
	}
	s._updateDisplayUsingFilterString(string_)
	return nil
}

// CanUpdateDisplayUsingFilterString reports whether the receiver responds to the private selector _updateDisplayUsingFilterString:.
func (s SOCustomizeSRLanguagesWindowController) CanUpdateDisplayUsingFilterString() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_updateDisplayUsingFilterString:"))
}
func (s SOCustomizeSRLanguagesWindowController) _updateRowDownloadStatus() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_updateRowDownloadStatus"))
}

// UpdateRowDownloadStatus is an exported wrapper for the private method _updateRowDownloadStatus.
func (s SOCustomizeSRLanguagesWindowController) UpdateRowDownloadStatus() error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_updateRowDownloadStatus")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateRowDownloadStatus"}
		return err
	}
	s._updateRowDownloadStatus()
	return nil
}

// CanUpdateRowDownloadStatus reports whether the receiver responds to the private selector _updateRowDownloadStatus.
func (s SOCustomizeSRLanguagesWindowController) CanUpdateRowDownloadStatus() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_updateRowDownloadStatus"))
}
func (s SOCustomizeSRLanguagesWindowController) AcceptSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("acceptSelection:"), selection)
}
func (s SOCustomizeSRLanguagesWindowController) CancelSelection(selection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("cancelSelection:"), selection)
}
func (s SOCustomizeSRLanguagesWindowController) NumberOfRowsInTableView(view objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](s.ID, objc.Sel("numberOfRowsInTableView:"), view)
	return rv
}
func (s SOCustomizeSRLanguagesWindowController) SearchFieldChanged(changed objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("searchFieldChanged:"), changed)
}
func (s SOCustomizeSRLanguagesWindowController) ShowSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems(window objectivec.IObject, identifiers objectivec.IObject, identifier objectivec.IObject, downloads bool, items bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("showSheetForWindow:networkSupportedLocaleIdentifiers:requiredLocaleIdentifier:supportDownloads:showOnlyNetworkSupportedItems:"), window, identifiers, identifier, downloads, items)
}
func (s SOCustomizeSRLanguagesWindowController) TableViewIsGroupRow(view objectivec.IObject, row int) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("tableView:isGroupRow:"), view, row)
	return rv
}
func (s SOCustomizeSRLanguagesWindowController) TableViewShouldSelectRow(view objectivec.IObject, row int64) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("tableView:shouldSelectRow:"), view, row)
	return rv
}
func (s SOCustomizeSRLanguagesWindowController) TableViewViewForTableColumnRow(view objectivec.IObject, column objectivec.IObject, row int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("tableView:viewForTableColumn:row:"), view, column, row)
	return objectivec.Object{ID: rv}
}
func (s SOCustomizeSRLanguagesWindowController) TableViewSelectionDidChange(change objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("tableViewSelectionDidChange:"), change)
}

func (s SOCustomizeSRLanguagesWindowController) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOCustomizeSRLanguagesWindowController) Description() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOCustomizeSRLanguagesWindowController) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SOCustomizeSRLanguagesWindowController) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
