// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_path_monitor protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_path_monitor
type OS_nw_path_monitor interface {
	objectivec.IObject
}

// OS_nw_path_monitorObject wraps an existing Objective-C object that conforms to the OS_nw_path_monitor protocol.
type OS_nw_path_monitorObject struct {
	objectivec.Object
}

func (o OS_nw_path_monitorObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_path_monitorObjectFromID constructs a [OS_nw_path_monitorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_path_monitorObjectFromID(id objc.ID) OS_nw_path_monitorObject {
	return OS_nw_path_monitorObject{
		Object: objectivec.ObjectFromID(id),
	}
}
