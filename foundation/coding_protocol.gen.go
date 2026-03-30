// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that enables an object to be encoded and decoded for archiving and distribution.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding
type NSCoding interface {
	objectivec.IObject

	// Encodes the receiver using a given archiver.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
	EncodeWithCoder(coder INSCoder)
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

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (o NSCodingObject) EncodeWithCoder(coder INSCoder) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeWithCoder:"), coder)
}
