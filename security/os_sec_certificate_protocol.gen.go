// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_sec_certificate protocol.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_certificate
type OS_sec_certificate interface {
	objectivec.IObject
}

// OS_sec_certificateObject wraps an existing Objective-C object that conforms to the OS_sec_certificate protocol.
type OS_sec_certificateObject struct {
	objectivec.Object
}

func (o OS_sec_certificateObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_certificateObjectFromID constructs a [OS_sec_certificateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_certificateObjectFromID(id objc.ID) OS_sec_certificateObject {
	return OS_sec_certificateObject{
		Object: objectivec.ObjectFromID(id),
	}
}
