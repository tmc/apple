// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that act as device sources.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionDeviceSource
type CMIOExtensionDeviceSource interface {
	objectivec.IObject
}

// CMIOExtensionDeviceSourceObject wraps an existing Objective-C object that conforms to the CMIOExtensionDeviceSource protocol.
type CMIOExtensionDeviceSourceObject struct {
	objectivec.Object
}
func (o CMIOExtensionDeviceSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CMIOExtensionDeviceSourceObjectFromID constructs a [CMIOExtensionDeviceSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CMIOExtensionDeviceSourceObjectFromID(id objc.ID) CMIOExtensionDeviceSourceObject {
	return CMIOExtensionDeviceSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

