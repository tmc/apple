// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods used to associate a unique identifier with objects in your user interface.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification
type NSUserInterfaceItemIdentification interface {
	objectivec.IObject

	// A string that identifies the user interface item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
	Identifier() NSUserInterfaceItemIdentifier

	// A string that identifies the user interface item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
	SetIdentifier(value NSUserInterfaceItemIdentifier)
}

// NSUserInterfaceItemIdentificationObject wraps an existing Objective-C object that conforms to the NSUserInterfaceItemIdentification protocol.
type NSUserInterfaceItemIdentificationObject struct {
	objectivec.Object
}

func (o NSUserInterfaceItemIdentificationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserInterfaceItemIdentificationObjectFromID constructs a [NSUserInterfaceItemIdentificationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserInterfaceItemIdentificationObjectFromID(id objc.ID) NSUserInterfaceItemIdentificationObject {
	return NSUserInterfaceItemIdentificationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (o NSUserInterfaceItemIdentificationObject) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}

// A string that identifies the user interface item.
//
// # Discussion
//
// Identifiers are used during window restoration operations to uniquely
// identify the windows of the application. You can set the value of this
// string programmatically or in Interface Builder. If you create an item in
// Interface Builder and do not set a value for this string, a unique value is
// created for the item when the nib file is loaded. For programmatically
// created views, you typically set this value after creating the item but
// before adding it to a window.
//
// You should not change the value of a window’s identifier after adding any
// views to the window. For views and controls in a window, the value you
// specify for this string must be unique on a per-window basis.
//
// The slash (`/`), backslash (`\`), or colon (`:`) characters are reserved
// and must not be used in your custom identifiers. Similarly, Apple reserves
// all identifiers beginning with an underscore (`_`) character. Applications
// and frameworks should use a consistent prefix for their identifiers to
// avoid collisions with other frameworks. For a list of prefixes used by the
// system frameworks, see [OS X Frameworks] in [Mac Technology Overview].
//
// If you are subclassing a class from one of the system frameworks, do not
// override the accessor methods of this protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
//
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
func (o NSUserInterfaceItemIdentificationObject) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}
