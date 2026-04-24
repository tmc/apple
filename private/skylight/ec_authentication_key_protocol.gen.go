// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationKey protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationKey
type ECAuthenticationKey interface {
	objectivec.IObject
}

// ECAuthenticationKeyObject wraps an existing Objective-C object that conforms to the ECAuthenticationKey protocol.
type ECAuthenticationKeyObject struct {
	objectivec.Object
}

func (o ECAuthenticationKeyObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECAuthenticationKeyObjectFromID constructs a [ECAuthenticationKeyObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECAuthenticationKeyObjectFromID(id objc.ID) ECAuthenticationKeyObject {
	return ECAuthenticationKeyObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationKey/createSignatureForMessage:
func (o ECAuthenticationKeyObject) CreateSignatureForMessage(message objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createSignatureForMessage:"), message)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationKey/signingContext
func (o ECAuthenticationKeyObject) SigningContext() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("signingContext"))
	return objectivec.Object{ID: rv}
}
