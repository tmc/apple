// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_advertise_descriptor protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_advertise_descriptor
type OS_nw_advertise_descriptor interface {
	objectivec.IObject
}

// OS_nw_advertise_descriptorObject wraps an existing Objective-C object that conforms to the OS_nw_advertise_descriptor protocol.
type OS_nw_advertise_descriptorObject struct {
	objectivec.Object
}

func (o OS_nw_advertise_descriptorObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_advertise_descriptorObjectFromID constructs a [OS_nw_advertise_descriptorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_advertise_descriptorObjectFromID(id objc.ID) OS_nw_advertise_descriptorObject {
	return OS_nw_advertise_descriptorObject{
		Object: objectivec.ObjectFromID(id),
	}
}
