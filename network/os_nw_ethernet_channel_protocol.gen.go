// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_ethernet_channel protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_ethernet_channel
type OS_nw_ethernet_channel interface {
	objectivec.IObject
}

// OS_nw_ethernet_channelObject wraps an existing Objective-C object that conforms to the OS_nw_ethernet_channel protocol.
type OS_nw_ethernet_channelObject struct {
	objectivec.Object
}

func (o OS_nw_ethernet_channelObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_ethernet_channelObjectFromID constructs a [OS_nw_ethernet_channelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_ethernet_channelObjectFromID(id objc.ID) OS_nw_ethernet_channelObject {
	return OS_nw_ethernet_channelObject{
		Object: objectivec.ObjectFromID(id),
	}
}
