// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSMenuItemValidation protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation
type NSMenuItemValidation interface {
	objectivec.IObject

	// Implemented to override the default action of enabling or disabling a specific menu item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
	ValidateMenuItem(menuItem INSMenuItem) bool
}

// NSMenuItemValidationObject wraps an existing Objective-C object that conforms to the NSMenuItemValidation protocol.
type NSMenuItemValidationObject struct {
	objectivec.Object
}
func (o NSMenuItemValidationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSMenuItemValidationObjectFromID constructs a [NSMenuItemValidationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMenuItemValidationObjectFromID(id objc.ID) NSMenuItemValidationObject {
	return NSMenuItemValidationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
// 
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (o NSMenuItemValidationObject) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
	}

