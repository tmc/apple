// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A `sec_object` is a generic, ARC-able type wrapper for common CoreFoundation Security types.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_object
type OS_sec_object interface {
	objectivec.IObject
}

// OS_sec_objectObject wraps an existing Objective-C object that conforms to the OS_sec_object protocol.
type OS_sec_objectObject struct {
	objectivec.Object
}
func (o OS_sec_objectObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_objectObjectFromID constructs a [OS_sec_objectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_objectObjectFromID(id objc.ID) OS_sec_objectObject {
	return OS_sec_objectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

