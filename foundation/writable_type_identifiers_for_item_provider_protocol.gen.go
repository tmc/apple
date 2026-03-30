// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An array of UTI strings representing the types of data that can be loaded for an item provider.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.type.property
type writableTypeIdentifiersForItemProvider interface {
	objectivec.IObject
}

// writableTypeIdentifiersForItemProviderObject wraps an existing Objective-C object that conforms to the writableTypeIdentifiersForItemProvider protocol.
type writableTypeIdentifiersForItemProviderObject struct {
	objectivec.Object
}

func (o writableTypeIdentifiersForItemProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// writableTypeIdentifiersForItemProviderObjectFromID constructs a [writableTypeIdentifiersForItemProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func writableTypeIdentifiersForItemProviderObjectFromID(id objc.ID) writableTypeIdentifiersForItemProviderObject {
	return writableTypeIdentifiersForItemProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}
