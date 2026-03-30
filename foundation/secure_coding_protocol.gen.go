// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that enables encoding and decoding in a manner that is robust against object substitution attacks.
//
// See: https://developer.apple.com/documentation/Foundation/NSSecureCoding
type NSSecureCoding interface {
	objectivec.IObject
	NSCoding
}

// NSSecureCodingObject wraps an existing Objective-C object that conforms to the NSSecureCoding protocol.
type NSSecureCodingObject struct {
	objectivec.Object
}

func (o NSSecureCodingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSecureCodingObjectFromID constructs a [NSSecureCodingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSecureCodingObjectFromID(id objc.ID) NSSecureCodingObject {
	return NSSecureCodingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (o NSSecureCodingObject) EncodeWithCoder(coder INSCoder) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), coder)
}
