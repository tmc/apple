// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A Boolean value that indicates whether or not the class supports secure coding.
//
// See: https://developer.apple.com/documentation/Foundation/NSSecureCoding/supportsSecureCoding
type supportsSecureCoding interface {
	objectivec.IObject
}

// supportsSecureCodingObject wraps an existing Objective-C object that conforms to the supportsSecureCoding protocol.
type supportsSecureCodingObject struct {
	objectivec.Object
}
func (o supportsSecureCodingObject) BaseObject() objectivec.Object {
	return o.Object
}

// supportsSecureCodingObjectFromID constructs a [supportsSecureCodingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func supportsSecureCodingObjectFromID(id objc.ID) supportsSecureCodingObject {
	return supportsSecureCodingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

