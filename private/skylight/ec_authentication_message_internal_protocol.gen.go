// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationMessageInternal protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageInternal
type ECAuthenticationMessageInternal interface {
	objectivec.IObject
}

// ECAuthenticationMessageInternalObject wraps an existing Objective-C object that conforms to the ECAuthenticationMessageInternal protocol.
type ECAuthenticationMessageInternalObject struct {
	objectivec.Object
}

func (o ECAuthenticationMessageInternalObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECAuthenticationMessageInternalObjectFromID constructs a [ECAuthenticationMessageInternalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECAuthenticationMessageInternalObjectFromID(id objc.ID) ECAuthenticationMessageInternalObject {
	return ECAuthenticationMessageInternalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageInternal/copySignedByKey:
func (o ECAuthenticationMessageInternalObject) CopySignedByKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copySignedByKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageInternal/signature
func (o ECAuthenticationMessageInternalObject) Signature() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("signature"))
	return objectivec.Object{ID: rv}
}
