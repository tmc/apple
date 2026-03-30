// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An object that serves as a data provider for data types that use lazy data fulfillment from a pasteboard request.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardTypeOwner
type NSPasteboardTypeOwner interface {
	objectivec.IObject

	// Requests that the object provide data for the data type to the pasteboard.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboardTypeOwner/pasteboard(_:provideDataForType:)
	PasteboardProvideDataForType(sender INSPasteboard, type_ NSPasteboardType)
}

// NSPasteboardTypeOwnerObject wraps an existing Objective-C object that conforms to the NSPasteboardTypeOwner protocol.
type NSPasteboardTypeOwnerObject struct {
	objectivec.Object
}

func (o NSPasteboardTypeOwnerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPasteboardTypeOwnerObjectFromID constructs a [NSPasteboardTypeOwnerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPasteboardTypeOwnerObjectFromID(id objc.ID) NSPasteboardTypeOwnerObject {
	return NSPasteboardTypeOwnerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests that the object provide data for the data type to the pasteboard.
//
// sender: The pasteboard requesting data.
//
// type: The data type.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardTypeOwner/pasteboard(_:provideDataForType:)
func (o NSPasteboardTypeOwnerObject) PasteboardProvideDataForType(sender INSPasteboard, type_ NSPasteboardType) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboard:provideDataForType:"), sender, objc.String(string(type_)))
}

// Notifies the object that the pasteboard’s owner changed.
//
// sender: The pasteboard whose owner changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardTypeOwner/pasteboardChangedOwner(_:)
func (o NSPasteboardTypeOwnerObject) PasteboardChangedOwner(sender INSPasteboard) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboardChangedOwner:"), sender)
}
