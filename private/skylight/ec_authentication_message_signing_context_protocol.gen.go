// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationMessageSigningContext protocol.
type ECAuthenticationMessageSigningContext interface {
	objectivec.IObject

	// UpdateSigningContextWithBytesLength protocol.
	UpdateSigningContextWithBytesLength(bytes []byte, length uint64)
}

// ECAuthenticationMessageSigningContextObject wraps an existing Objective-C object that conforms to the ECAuthenticationMessageSigningContext protocol.
type ECAuthenticationMessageSigningContextObject struct {
	objectivec.Object
}

func (o ECAuthenticationMessageSigningContextObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECAuthenticationMessageSigningContextObjectFromID constructs a [ECAuthenticationMessageSigningContextObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECAuthenticationMessageSigningContextObjectFromID(id objc.ID) ECAuthenticationMessageSigningContextObject {
	return ECAuthenticationMessageSigningContextObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECAuthenticationMessageSigningContextObject) FinalizedData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("finalizedData"))
	return objectivec.Object{ID: rv}
}
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithBytesLength(bytes []byte, length uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithBytes:length:"), bytes, length)
}
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithData(data objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithData:"), data)
}
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithObject(object objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithObject:"), object)
}
