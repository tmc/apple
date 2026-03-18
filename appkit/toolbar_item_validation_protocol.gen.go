// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Validation of a toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemValidation
type NSToolbarItemValidation interface {
	objectivec.IObject

	// Determines whether to enable or disable the toolbar item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemValidation/validateToolbarItem(_:)
	ValidateToolbarItem(item INSToolbarItem) bool
}

// NSToolbarItemValidationObject wraps an existing Objective-C object that conforms to the NSToolbarItemValidation protocol.
type NSToolbarItemValidationObject struct {
	objectivec.Object
}
func (o NSToolbarItemValidationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSToolbarItemValidationObjectFromID constructs a [NSToolbarItemValidationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSToolbarItemValidationObjectFromID(id objc.ID) NSToolbarItemValidationObject {
	return NSToolbarItemValidationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Determines whether to enable or disable the toolbar item.
//
// # Discussion
// 
// If this method is implemented and returns [false], [NSToolbar] will disable
// `item`. Returning [true] causes `item` to be enabled.
// 
// [NSToolbar] only calls this method for image items.
// 
// If the receiver is the `target` for the actions of multiple toolbar items,
// it’s necessary to determine which toolbar item `item` refers to by
// testing the `itemIdentifier`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemValidation/validateToolbarItem(_:)

func (o NSToolbarItemValidationObject) ValidateToolbarItem(item INSToolbarItem) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("validateToolbarItem:"), item)
	return rv
	}

