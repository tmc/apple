// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for implementing a class to allow an item provider to create an instance of the class.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderReading
type NSItemProviderReading interface {
	objectivec.IObject
}

// NSItemProviderReadingObject wraps an existing Objective-C object that conforms to the NSItemProviderReading protocol.
type NSItemProviderReadingObject struct {
	objectivec.Object
}
func (o NSItemProviderReadingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSItemProviderReadingObjectFromID constructs a [NSItemProviderReadingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSItemProviderReadingObjectFromID(id objc.ID) NSItemProviderReadingObject {
	return NSItemProviderReadingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

