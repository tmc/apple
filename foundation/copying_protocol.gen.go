// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that objects adopt to provide functional copies of themselves.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopying
type NSCopying interface {
	objectivec.IObject
}

// NSCopyingObject wraps an existing Objective-C object that conforms to the NSCopying protocol.
type NSCopyingObject struct {
	objectivec.Object
}
func (o NSCopyingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCopyingObjectFromID constructs a [NSCopyingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCopyingObjectFromID(id objc.ID) NSCopyingObject {
	return NSCopyingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

