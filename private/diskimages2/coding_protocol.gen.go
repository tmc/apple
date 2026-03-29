// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSCoding protocol.
//
// See: https://developer.apple.com/documentation/DiskImages2/NSCoding
type NSCoding interface {
	objectivec.IObject
}

// NSCodingObject wraps an existing Objective-C object that conforms to the NSCoding protocol.
type NSCodingObject struct {
	objectivec.Object
}
func (o NSCodingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCodingObjectFromID constructs a [NSCodingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCodingObjectFromID(id objc.ID) NSCodingObject {
	return NSCodingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/DiskImages2/NSCoding/encodeWithCoder:
func (o NSCodingObject) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), coder)
	}

