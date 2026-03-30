// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that a custom class can adopt to manage the enabled state of a UI element.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations
type NSUserInterfaceValidations interface {
	objectivec.IObject

	// Returns a Boolean value that indicates whether the sender should be enabled.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
	ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool
}

// NSUserInterfaceValidationsObject wraps an existing Objective-C object that conforms to the NSUserInterfaceValidations protocol.
type NSUserInterfaceValidationsObject struct {
	objectivec.Object
}

func (o NSUserInterfaceValidationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserInterfaceValidationsObjectFromID constructs a [NSUserInterfaceValidationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserInterfaceValidationsObjectFromID(id objc.ID) NSUserInterfaceValidationsObject {
	return NSUserInterfaceValidationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
//
// true if the user interface item should be enabled, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (o NSUserInterfaceValidationsObject) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}
