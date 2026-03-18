// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_sec_identity protocol.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_identity
type OS_sec_identity interface {
	objectivec.IObject
}

// OS_sec_identityObject wraps an existing Objective-C object that conforms to the OS_sec_identity protocol.
type OS_sec_identityObject struct {
	objectivec.Object
}
func (o OS_sec_identityObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_identityObjectFromID constructs a [OS_sec_identityObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_identityObjectFromID(id objc.ID) OS_sec_identityObject {
	return OS_sec_identityObject{
		Object: objectivec.ObjectFromID(id),
	}
}

