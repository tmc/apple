// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// NSComboBoxDataSource protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxDataSource
type NSComboBoxDataSource interface {
	objectivec.IObject
}

// NSComboBoxDataSourceObject wraps an existing Objective-C object that conforms to the NSComboBoxDataSource protocol.
type NSComboBoxDataSourceObject struct {
	objectivec.Object
}
func (o NSComboBoxDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSComboBoxDataSourceObjectFromID constructs a [NSComboBoxDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSComboBoxDataSourceObjectFromID(id objc.ID) NSComboBoxDataSourceObject {
	return NSComboBoxDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the first item from the pop-up list that starts with the text the
// user has typed.
//
// comboBox: The combo box.
//
// string: The string to match against items in the combo box’s pop-up list. This is
// text that the user has typed.
//
// # Return Value
// 
// The first complete string from the items in the combo box’s pop-up list
// that starts with the string in `uncompletedString`.
//
// # Discussion
// 
// An [NSComboBox] object uses this method to perform incremental—or
// “smart”—searches when the user types into the text field. As the user
// types in the text field, the receiver uses this method to search for items
// from the pop-up list that start with what the user has typed. The receiver
// adds the new text to the end of the field and selects the new text, so when
// the user types another character, it replaces the new text.
// 
// This method is optional. If you don’t implement it, the receiver does not
// perform incremental searches.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxDataSource/comboBox(_:completedString:)
func (o NSComboBoxDataSourceObject) ComboBoxCompletedString(comboBox INSComboBox, string_ string) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comboBox:completedString:"), comboBox, objc.String(string_))
	return foundation.NSStringFromID(rv).String()
	}
// Returns the index of the combo box item matching the specified string.
//
// comboBox: The combo box.
//
// string: The string to match against the items in the combo box. If the datasource
// implements[ComboBoxCompletedString], this is the string returned by that
// method. Otherwise, it is the text that the user has typed.
//
// # Return Value
// 
// The index for the item that matches the specified string, or [NSNotFound]
// if no item matches.
//
// # Discussion
// 
// An [NSComboBox] object uses this method to synchronize the pop-up list’s
// selected item with the text field’s contents. If you don’t implement
// this method the receiver does not synchronize the pop-up list’s selected
// item with the text field’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxDataSource/comboBox(_:indexOfItemWithStringValue:)
func (o NSComboBoxDataSourceObject) ComboBoxIndexOfItemWithStringValue(comboBox INSComboBox, string_ string) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("comboBox:indexOfItemWithStringValue:"), comboBox, objc.String(string_))
	return rv
	}
// Returns the object that corresponds to the item at the specified index in
// the combo box.
//
// comboBox: The combo box.
//
// index: The index of the item to return.
//
// # Return Value
// 
// The object corresponding to the specified index number.
//
// # Discussion
// 
// An [NSComboBox] object uses this method to populate the items displayed in
// its pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxDataSource/comboBox(_:objectValueForItemAt:)
func (o NSComboBoxDataSourceObject) ComboBoxObjectValueForItemAtIndex(comboBox INSComboBox, index int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comboBox:objectValueForItemAtIndex:"), comboBox, index)
	return objectivec.Object{ID: rv}
	}
// Returns the number of items that the data source manages for the combo box.
//
// comboBox: The combo box.
//
// # Return Value
// 
// The number of items that the data source object manages for the specified
// combo box.
//
// # Discussion
// 
// An [NSComboBox] object uses this method to determine how many items it
// should display in its pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxDataSource/numberOfItems(in:)
func (o NSComboBoxDataSourceObject) NumberOfItemsInComboBox(comboBox INSComboBox) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("numberOfItemsInComboBox:"), comboBox)
	return rv
	}

