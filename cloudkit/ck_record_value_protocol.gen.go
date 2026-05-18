// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol that provides strong type-checking for objects that the CloudKit framework stores on the server.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordValue-c.protocol
type CKRecordValue interface {
	objectivec.IObject
}

// CKRecordValueObject wraps an existing Objective-C object that conforms to the CKRecordValue protocol.
type CKRecordValueObject struct {
	objectivec.Object
}

func (o CKRecordValueObject) BaseObject() objectivec.Object {
	return o.Object
}

// CKRecordValueObjectFromID constructs a [CKRecordValueObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CKRecordValueObjectFromID(id objc.ID) CKRecordValueObject {
	return CKRecordValueObject{
		Object: objectivec.ObjectFromID(id),
	}
}
