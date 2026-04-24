// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECAuthenticationMessageSigningContext protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext
type ECAuthenticationMessageSigningContext interface {
	objectivec.IObject

	// UpdateSigningContextWithBytesLength protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext/updateSigningContextWithBytes:length:
	UpdateSigningContextWithBytesLength(bytes []byte)
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

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext/finalizedData
func (o ECAuthenticationMessageSigningContextObject) FinalizedData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("finalizedData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext/updateSigningContextWithBytes:length:
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithBytesLength(bytes []byte) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithBytes:length:"), unsafe.Pointer(unsafe.SliceData(bytes)), uint(len(bytes)))
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext/updateSigningContextWithData:
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithData(data objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithData:"), data)
}

// See: https://developer.apple.com/documentation/SkyLight/ECAuthenticationMessageSigningContext/updateSigningContextWithObject:
func (o ECAuthenticationMessageSigningContextObject) UpdateSigningContextWithObject(object objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("updateSigningContextWithObject:"), object)
}
