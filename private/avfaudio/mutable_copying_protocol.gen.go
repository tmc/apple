// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSMutableCopying protocol.
//
// See: https://developer.apple.com/documentation/AVFAudio/NSMutableCopying
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
