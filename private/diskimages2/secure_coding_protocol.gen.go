// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSSecureCoding protocol.
//
// See: https://developer.apple.com/documentation/DiskImages2/NSSecureCoding
type NSSecureCoding interface {
	objectivec.IObject

	// SupportsSecureCoding protocol.
	//
	// See: https://developer.apple.com/documentation/DiskImages2/NSSecureCoding/supportsSecureCoding
	SupportsSecureCoding() bool
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

// See: https://developer.apple.com/documentation/DiskImages2/NSSecureCoding/supportsSecureCoding
func (o NSSecureCodingObject) SupportsSecureCoding() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsSecureCoding"))
	return rv
	}

