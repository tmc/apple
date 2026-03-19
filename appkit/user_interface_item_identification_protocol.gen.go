// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

func (o NSUserInterfaceItemIdentificationObject) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

