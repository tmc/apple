// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that support interaction with items users can share through a sharing service.
//
// See: https://developer.apple.com/documentation/AppKit/NSServicesMenuRequestor
type NSServicesMenuRequestor interface {
	objectivec.IObject
}

// NSServicesMenuRequestorObject wraps an existing Objective-C object that conforms to the NSServicesMenuRequestor protocol.
type NSServicesMenuRequestorObject struct {
	objectivec.Object
}
func (o NSServicesMenuRequestorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSServicesMenuRequestorObjectFromID constructs a [NSServicesMenuRequestorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSServicesMenuRequestorObjectFromID(id objc.ID) NSServicesMenuRequestorObject {
	return NSServicesMenuRequestorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Reads data from the pasteboard and uses it to replace the current
// selection.
//
// pboard: The pasteboard containing the data to read.
//
// # Return Value
// 
// [true] if your implementation was able to read the pasteboard data
// successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You implement this method to replace your application’s current selection
// (that is, the text or objects that are currently selected) with the data on
// the pasteboard. The data would have been placed in the pasteboard by
// another application in response to a remote message from the Services menu.
// A [ReadSelectionFromPasteboard] message is sent to the same object that
// previously received a [WriteSelectionToPasteboardTypes] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSServicesMenuRequestor/readSelection(from:)

func (o NSServicesMenuRequestorObject) ReadSelectionFromPasteboard(pboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("readSelectionFromPasteboard:"), pboard)
	return rv
	}

// Writes the current selection to the pasteboard.
//
// pboard: The pasteboard to receive your data.
//
// types: An array of [NSString] objects listing the types of data that you should
// write to the pasteboard. You should write data to the pasteboard for as
// many of the types as you support.
//
// # Return Value
// 
// [true] if your implementation was able to write one or more types to the
// pasteboard; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A [WriteSelectionToPasteboardTypes] message is sent to the first responder
// when the user chooses a command from the Services menu, but only if the
// receiver didn’t return `nil` to a previous
// [ValidRequestorForSendTypeReturnType] message.
// 
// After your method writes the data to the pasteboard, a remote message is
// sent to the application that provides the service the user requested. If
// the service provider supplies return data to replace the selection, the
// first responder will then receive a [ReadSelectionFromPasteboard] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSServicesMenuRequestor/writeSelection(to:types:)

func (o NSServicesMenuRequestorObject) WriteSelectionToPasteboardTypes(pboard INSPasteboard, types []string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("writeSelectionToPasteboard:types:"), pboard, objectivec.StringSliceToNSArray(types))
	return rv
	}

