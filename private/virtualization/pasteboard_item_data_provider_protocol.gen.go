// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSPasteboardItemDataProvider protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/NSPasteboardItemDataProvider
type NSPasteboardItemDataProvider interface {
	objectivec.IObject
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

// See: https://developer.apple.com/documentation/Virtualization/NSPasteboardItemDataProvider/pasteboard:item:provideDataForType:
func (o NSPasteboardItemDataProviderObject) PasteboardItemProvideDataForType(pasteboard objectivec.IObject, item objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboard:item:provideDataForType:"), pasteboard, item, type_)
}

// See: https://developer.apple.com/documentation/Virtualization/NSPasteboardItemDataProvider/pasteboardFinishedWithDataProvider:
func (o NSPasteboardItemDataProviderObject) PasteboardFinishedWithDataProvider(provider objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pasteboardFinishedWithDataProvider:"), provider)
}
