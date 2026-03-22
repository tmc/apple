// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// NSComboBoxCellDataSource protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCellDataSource
type NSComboBoxCellDataSource interface {
	objectivec.IObject
}

// NSComboBoxCellDataSourceObject wraps an existing Objective-C object that conforms to the NSComboBoxCellDataSource protocol.
type NSComboBoxCellDataSourceObject struct {
	objectivec.Object
}
func (o NSComboBoxCellDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSComboBoxCellDataSourceObjectFromID constructs a [NSComboBoxCellDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSComboBoxCellDataSourceObjectFromID(id objc.ID) NSComboBoxCellDataSourceObject {
	return NSComboBoxCellDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the item from the combo box’s pop-up list that matches the text
// entered by the user.
//
// comboBoxCell: The combo box cell.
//
// uncompletedString: The substring containing the text the user typed into the text field of the
// combo box cell.
//
// # Return Value
// 
// The completed string, from the items in the pop-up list, that matches the
// text entered by the user. Your implementation should return the first
// complete string that starts with `uncompletedString`.
//
// # Discussion
// 
// An [NSComboBoxCell] object uses this method to perform incremental—or
// “smart”—searches when the user types into the text field.
// 
// As the user types in the text field, the receiver uses this method to
// search for items from the pop-up list that start with what the user has
// typed. The receiver adds the new text to the end of the field and selects
// the new text, so when the user types another character, it replaces the new
// text.
// 
// If you don’t implement this method, the receiver does not perform
// incremental searches.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCellDataSource/comboBoxCell(_:completedString:)
func (o NSComboBoxCellDataSourceObject) ComboBoxCellCompletedString(comboBoxCell INSComboBoxCell, uncompletedString string) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comboBoxCell:completedString:"), comboBoxCell, objc.String(uncompletedString))
	return foundation.NSStringFromID(rv).String()
	}
// Invoked by an [NSComboBoxCell] object to synchronize the pop-up list’s
// selected item with the text field’s contents.
//
// comboBoxCell: The combo box cell.
//
// string: The string to match. If [ComboBoxCellCompletedString] is implemented,
// `aString` is the string returned by that method. Otherwise, `aString` is
// the text that the user has typed.
//
// # Return Value
// 
// The index for the pop-up list item matching `aString`, or [NSNotFound] if
// no item matches.
//
// # Discussion
// 
// If you don’t implement this method, the receiver does not synchronize the
// pop-up list’s selected item with the text field’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCellDataSource/comboBoxCell(_:indexOfItemWithStringValue:)
func (o NSComboBoxCellDataSourceObject) ComboBoxCellIndexOfItemWithStringValue(comboBoxCell INSComboBoxCell, string_ string) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("comboBoxCell:indexOfItemWithStringValue:"), comboBoxCell, objc.String(string_))
	return rv
	}
// Returns the object that corresponds to the item at the given index in the
// combo box cell.
//
// comboBoxCell: The combo box cell for which to return the item.
//
// index: The index of the item to return.
//
// # Return Value
// 
// The object corresponding to the item at the specified index in the given
// combo box cell.
//
// # Discussion
// 
// An [NSComboBoxCell] object uses this method to populate the items displayed
// in its pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCellDataSource/comboBoxCell(_:objectValueForItemAt:)
func (o NSComboBoxCellDataSourceObject) ComboBoxCellObjectValueForItemAtIndex(comboBoxCell INSComboBoxCell, index int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comboBoxCell:objectValueForItemAtIndex:"), comboBoxCell, index)
	return objectivec.Object{ID: rv}
	}
// Returns the number of items managed for the combo box cell by your data
// source object.
//
// comboBoxCell: The combo box cell for which your data source manages items.
//
// # Return Value
// 
// The number of items your data source object manages.
//
// # Discussion
// 
// An [NSComboBoxCell] object uses this method to determine how many items it
// should display in its pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCellDataSource/numberOfItems(in:)
func (o NSComboBoxCellDataSourceObject) NumberOfItemsInComboBoxCell(comboBoxCell INSComboBoxCell) int {
	rv := objc.Send[int](o.ID, objc.Sel("numberOfItemsInComboBoxCell:"), comboBoxCell)
	return rv
	}

