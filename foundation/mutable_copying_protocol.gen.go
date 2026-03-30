// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that mutable objects adopt to provide functional copies of themselves.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableCopying
type NSMutableCopying interface {
	objectivec.IObject
}

// NSMutableCopyingObject wraps an existing Objective-C object that conforms to the NSMutableCopying protocol.
type NSMutableCopyingObject struct {
	objectivec.Object
}

func (o NSMutableCopyingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSMutableCopyingObjectFromID constructs a [NSMutableCopyingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMutableCopyingObjectFromID(id objc.ID) NSMutableCopyingObject {
	return NSMutableCopyingObject{
		Object: objectivec.ObjectFromID(id),
	}
}
