// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_connection_group protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_connection_group
type OS_nw_connection_group interface {
	objectivec.IObject
}

// OS_nw_connection_groupObject wraps an existing Objective-C object that conforms to the OS_nw_connection_group protocol.
type OS_nw_connection_groupObject struct {
	objectivec.Object
}

func (o OS_nw_connection_groupObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_connection_groupObjectFromID constructs a [OS_nw_connection_groupObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_connection_groupObjectFromID(id objc.ID) OS_nw_connection_groupObject {
	return OS_nw_connection_groupObject{
		Object: objectivec.ObjectFromID(id),
	}
}
