// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods implemented by the data provider of a pasteboard item to provide the data for a particular UTI type.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItemDataProvider
type NSPasteboardItemDataProvider interface {
	objectivec.IObject

	// Asks the receiver to provide data for a specified type to a given pasteboard.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItemDataProvider/pasteboard(_:item:provideDataForType:)
	PasteboardItemProvideDataForType(pasteboard INSPasteboard, item INSPasteboardItem, type_ NSPasteboardType)
}

// NSPasteboardItemDataProviderObject wraps an existing Objective-C object that conforms to the NSPasteboardItemDataProvider protocol.
type NSPasteboardItemDataProviderObject struct {
	objectivec.Object
}

func (o NSPasteboardItemDataProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPasteboardItemDataProviderObjectFromID constructs a [NSPasteboardItemDataProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPasteboardItemDataProviderObjectFromID(id objc.ID) NSPasteboardItemDataProviderObject {
	return NSPasteboardItemDataProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the receiver to provide data for a specified type to a given
// pasteboard.
//
// pasteboard: A pasteboard to which the receiver has promised to provide data.
//
// item: A pasteboard item for which the receiver has promised to provide data
//
// type: A UTI type string.
//
// # Discussion
//
// The receiver was previously set as the provider using
// [SetDataProviderForTypes].
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItemDataProvider/pasteboard(_:item:provideDataForType:)
func (o NSPasteboardItemDataProviderObject) PasteboardItemProvideDataForType(pasteboard INSPasteboard, item INSPasteboardItem, type_ NSPasteboardType) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboard:item:provideDataForType:"), pasteboard, item, objc.String(string(type_)))
}

// Informs the receiver that the pasteboard no longer needs the data provider
// for any of its pasteboard items.
//
// pasteboard: A pasteboard.
//
// # Discussion
//
// One data provider can provide data for more than one pasteboard item. This
// method is called when the pasteboard no longer needs the data provider for
// any of its pasteboard items. This can be either because the data provider
// has fulfilled all promises, or because ownership of the pasteboard has
// changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItemDataProvider/pasteboardFinishedWithDataProvider(_:)
func (o NSPasteboardItemDataProviderObject) PasteboardFinishedWithDataProvider(pasteboard INSPasteboard) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboardFinishedWithDataProvider:"), pasteboard)
}
