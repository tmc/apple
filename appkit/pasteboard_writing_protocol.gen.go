// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that defines the interface for retrieving a representation of an object that can be written to a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting
type NSPasteboardWriting interface {
	objectivec.IObject

	// Returns an array of UTI strings of data types the receiver can write to a given pasteboard.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
	WritableTypesForPasteboard(pasteboard INSPasteboard) []string
}

// NSPasteboardWritingObject wraps an existing Objective-C object that conforms to the NSPasteboardWriting protocol.
type NSPasteboardWritingObject struct {
	objectivec.Object
}

func (o NSPasteboardWritingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPasteboardWritingObjectFromID constructs a [NSPasteboardWritingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPasteboardWritingObjectFromID(id objc.ID) NSPasteboardWritingObject {
	return NSPasteboardWritingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
//
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
//
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [NSPasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (o NSPasteboardWritingObject) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
//
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
//
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (o NSPasteboardWritingObject) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// # Discussion
//
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
func (o NSPasteboardWritingObject) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](o.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return rv
}
