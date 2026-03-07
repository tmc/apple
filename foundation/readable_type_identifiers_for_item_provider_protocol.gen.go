// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An array of UTI strings representing the data types supported by the class.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderReading/readableTypeIdentifiersForItemProvider
type readableTypeIdentifiersForItemProvider interface {
	objectivec.IObject
}



// readableTypeIdentifiersForItemProviderObject wraps an existing Objective-C object that conforms to the readableTypeIdentifiersForItemProvider protocol.
type readableTypeIdentifiersForItemProviderObject struct {
	objectivec.Object
}
func (o readableTypeIdentifiersForItemProviderObject) BaseObject() objectivec.Object {
	return o.Object
}



// readableTypeIdentifiersForItemProviderObjectFromID constructs a [readableTypeIdentifiersForItemProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func readableTypeIdentifiersForItemProviderObjectFromID(id objc.ID) readableTypeIdentifiersForItemProviderObject {
	return readableTypeIdentifiersForItemProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}










