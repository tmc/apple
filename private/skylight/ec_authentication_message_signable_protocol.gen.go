// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationMessageSignable protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSignable
type ECAuthenticationMessageSignable interface {
	objectivec.IObject
}

// ECAuthenticationMessageSignableObject wraps an existing Objective-C object that conforms to the ECAuthenticationMessageSignable protocol.
type ECAuthenticationMessageSignableObject struct {
	objectivec.Object
}

func (o ECAuthenticationMessageSignableObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECAuthenticationMessageSignableObjectFromID constructs a [ECAuthenticationMessageSignableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECAuthenticationMessageSignableObjectFromID(id objc.ID) ECAuthenticationMessageSignableObject {
	return ECAuthenticationMessageSignableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSignable/addToSigningContext:
func (o ECAuthenticationMessageSignableObject) AddToSigningContext(context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("addToSigningContext:"), context)
}
