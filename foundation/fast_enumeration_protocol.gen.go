// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that objects adopt to support fast enumeration.
//
// See: https://developer.apple.com/documentation/Foundation/NSFastEnumeration
type NSFastEnumeration interface {
	objectivec.IObject
}

// NSFastEnumerationObject wraps an existing Objective-C object that conforms to the NSFastEnumeration protocol.
type NSFastEnumerationObject struct {
	objectivec.Object
}

func (o NSFastEnumerationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFastEnumerationObjectFromID constructs a [NSFastEnumerationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFastEnumerationObjectFromID(id objc.ID) NSFastEnumerationObject {
	return NSFastEnumerationObject{
		Object: objectivec.ObjectFromID(id),
	}
}
