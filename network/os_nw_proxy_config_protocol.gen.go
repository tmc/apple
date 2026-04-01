// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_proxy_config protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_proxy_config
type OS_nw_proxy_config interface {
	objectivec.IObject
}

// OS_nw_proxy_configObject wraps an existing Objective-C object that conforms to the OS_nw_proxy_config protocol.
type OS_nw_proxy_configObject struct {
	objectivec.Object
}

func (o OS_nw_proxy_configObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_proxy_configObjectFromID constructs a [OS_nw_proxy_configObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_proxy_configObjectFromID(id objc.ID) OS_nw_proxy_configObject {
	return OS_nw_proxy_configObject{
		Object: objectivec.ObjectFromID(id),
	}
}
