// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that a custom class can adopt to manage the automatic enablement of a UI control.
//
// See: https://developer.apple.com/documentation/AppKit/NSValidatedUserInterfaceItem
type NSValidatedUserInterfaceItem interface {
	objectivec.IObject

	// Returns the selector of the receiver’s action method.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSValidatedUserInterfaceItem/action
	Action() objc.SEL

	// Returns the receiver’s tag integer.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSValidatedUserInterfaceItem/tag
	Tag() int
}

// NSValidatedUserInterfaceItemObject wraps an existing Objective-C object that conforms to the NSValidatedUserInterfaceItem protocol.
type NSValidatedUserInterfaceItemObject struct {
	objectivec.Object
}
func (o NSValidatedUserInterfaceItemObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSValidatedUserInterfaceItemObjectFromID constructs a [NSValidatedUserInterfaceItemObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSValidatedUserInterfaceItemObjectFromID(id objc.ID) NSValidatedUserInterfaceItemObject {
	return NSValidatedUserInterfaceItemObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the selector of the receiver’s action method.
//
// See: https://developer.apple.com/documentation/AppKit/NSValidatedUserInterfaceItem/action
func (o NSValidatedUserInterfaceItemObject) Action() objc.SEL {
	rv := objc.Send[objc.SEL](o.ID, objc.Sel("action"))
	return rv
	}
// Returns the receiver’s tag integer.
//
// See: https://developer.apple.com/documentation/AppKit/NSValidatedUserInterfaceItem/tag
func (o NSValidatedUserInterfaceItemObject) Tag() int {
	rv := objc.Send[int](o.ID, objc.Sel("tag"))
	return rv
	}

