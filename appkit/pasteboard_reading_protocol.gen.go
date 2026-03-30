// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that defines the interface for initializing an object from a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading
type NSPasteboardReading interface {
	objectivec.IObject
}

// NSPasteboardReadingObject wraps an existing Objective-C object that conforms to the NSPasteboardReading protocol.
type NSPasteboardReadingObject struct {
	objectivec.Object
}

func (o NSPasteboardReadingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPasteboardReadingObjectFromID constructs a [NSPasteboardReadingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPasteboardReadingObjectFromID(id objc.ID) NSPasteboardReadingObject {
	return NSPasteboardReadingObject{
		Object: objectivec.ObjectFromID(id),
	}
}
