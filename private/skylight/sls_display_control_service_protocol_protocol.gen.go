// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSDisplayControlServiceProtocol protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlServiceProtocol
type SLSDisplayControlServiceProtocol interface {
	objectivec.IObject
}

// SLSDisplayControlServiceProtocolObject wraps an existing Objective-C object that conforms to the SLSDisplayControlServiceProtocol protocol.
type SLSDisplayControlServiceProtocolObject struct {
	objectivec.Object
}

func (o SLSDisplayControlServiceProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSDisplayControlServiceProtocolObjectFromID constructs a [SLSDisplayControlServiceProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSDisplayControlServiceProtocolObjectFromID(id objc.ID) SLSDisplayControlServiceProtocolObject {
	return SLSDisplayControlServiceProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlServiceProtocol/service
func (o SLSDisplayControlServiceProtocolObject) Service() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("service"))
	return objectivec.Object{ID: rv}
}
