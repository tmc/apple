// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// These are os_object compatible and ARC-able wrappers around existing CoreFoundation Security types, including: SecTrustRef, SecIdentityRef, and SecCertificateRef. They allow clients to use these types in os_object-type APIs and data structures. The underlying CoreFoundation types may be extracted and used by clients as needed.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_trust
type OS_sec_trust interface {
	objectivec.IObject
}

// OS_sec_trustObject wraps an existing Objective-C object that conforms to the OS_sec_trust protocol.
type OS_sec_trustObject struct {
	objectivec.Object
}

func (o OS_sec_trustObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_trustObjectFromID constructs a [OS_sec_trustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_trustObjectFromID(id objc.ID) OS_sec_trustObject {
	return OS_sec_trustObject{
		Object: objectivec.ObjectFromID(id),
	}
}
