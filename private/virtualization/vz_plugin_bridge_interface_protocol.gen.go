// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZPluginBridgeInterface protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPluginBridgeInterface
type VZPluginBridgeInterface interface {
	objectivec.IObject
}

// VZPluginBridgeInterfaceObject wraps an existing Objective-C object that conforms to the VZPluginBridgeInterface protocol.
type VZPluginBridgeInterfaceObject struct {
	objectivec.Object
}

func (o VZPluginBridgeInterfaceObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZPluginBridgeInterfaceObjectFromID constructs a [VZPluginBridgeInterfaceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZPluginBridgeInterfaceObjectFromID(id objc.ID) VZPluginBridgeInterfaceObject {
	return VZPluginBridgeInterfaceObject{
		Object: objectivec.ObjectFromID(id),
	}
}
