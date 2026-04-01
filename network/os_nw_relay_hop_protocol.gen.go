// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_relay_hop protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_relay_hop
type OS_nw_relay_hop interface {
	objectivec.IObject
}

// OS_nw_relay_hopObject wraps an existing Objective-C object that conforms to the OS_nw_relay_hop protocol.
type OS_nw_relay_hopObject struct {
	objectivec.Object
}

func (o OS_nw_relay_hopObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_relay_hopObjectFromID constructs a [OS_nw_relay_hopObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_relay_hopObjectFromID(id objc.ID) OS_nw_relay_hopObject {
	return OS_nw_relay_hopObject{
		Object: objectivec.ObjectFromID(id),
	}
}
