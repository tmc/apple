// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_endpoint protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_endpoint
type OS_nw_endpoint interface {
	objectivec.IObject
}

// OS_nw_endpointObject wraps an existing Objective-C object that conforms to the OS_nw_endpoint protocol.
type OS_nw_endpointObject struct {
	objectivec.Object
}

func (o OS_nw_endpointObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_endpointObjectFromID constructs a [OS_nw_endpointObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_endpointObjectFromID(id objc.ID) OS_nw_endpointObject {
	return OS_nw_endpointObject{
		Object: objectivec.ObjectFromID(id),
	}
}
