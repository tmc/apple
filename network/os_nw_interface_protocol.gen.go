// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_interface protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_interface
type OS_nw_interface interface {
	objectivec.IObject
}

// OS_nw_interfaceObject wraps an existing Objective-C object that conforms to the OS_nw_interface protocol.
type OS_nw_interfaceObject struct {
	objectivec.Object
}

func (o OS_nw_interfaceObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_interfaceObjectFromID constructs a [OS_nw_interfaceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_interfaceObjectFromID(id objc.ID) OS_nw_interfaceObject {
	return OS_nw_interfaceObject{
		Object: objectivec.ObjectFromID(id),
	}
}
