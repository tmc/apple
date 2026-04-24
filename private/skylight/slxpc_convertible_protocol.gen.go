// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLXPCConvertible protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLXPCConvertible
type SLXPCConvertible interface {
	objectivec.IObject
}

// SLXPCConvertibleObject wraps an existing Objective-C object that conforms to the SLXPCConvertible protocol.
type SLXPCConvertibleObject struct {
	objectivec.Object
}

func (o SLXPCConvertibleObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLXPCConvertibleObjectFromID constructs a [SLXPCConvertibleObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLXPCConvertibleObjectFromID(id objc.ID) SLXPCConvertibleObject {
	return SLXPCConvertibleObject{
		Object: objectivec.ObjectFromID(id),
	}
}
