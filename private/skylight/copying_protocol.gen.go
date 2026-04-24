// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSCopying protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/NSCopying
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
