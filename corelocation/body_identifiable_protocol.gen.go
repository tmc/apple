// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CLBodyIdentifiable protocol.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLBodyIdentifiable
type CLBodyIdentifiable interface {
	objectivec.IObject
}

// CLBodyIdentifiableObject wraps an existing Objective-C object that conforms to the CLBodyIdentifiable protocol.
type CLBodyIdentifiableObject struct {
	objectivec.Object
}

func (o CLBodyIdentifiableObject) BaseObject() objectivec.Object {
	return o.Object
}

// CLBodyIdentifiableObjectFromID constructs a [CLBodyIdentifiableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CLBodyIdentifiableObjectFromID(id objc.ID) CLBodyIdentifiableObject {
	return CLBodyIdentifiableObject{
		Object: objectivec.ObjectFromID(id),
	}
}
