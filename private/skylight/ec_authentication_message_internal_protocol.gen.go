// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationMessageInternal protocol.
type ECAuthenticationMessageInternal interface {
	objectivec.IObject

	// CopySignedByKey protocol.
	CopySignedByKey(key objectivec.IObject) objectivec.IObject

	// Signature protocol.
	Signature() objectivec.IObject
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

func (o ECAuthenticationMessageInternalObject) CopySignedByKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("copySignedByKey:"), key)
	return objectivec.Object{ID: rv}
}
func (o ECAuthenticationMessageInternalObject) Signature() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("signature"))
	return objectivec.Object{ID: rv}
}
