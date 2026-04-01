// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_ws_request protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_ws_request
type OS_nw_ws_request interface {
	objectivec.IObject
}

// OS_nw_ws_requestObject wraps an existing Objective-C object that conforms to the OS_nw_ws_request protocol.
type OS_nw_ws_requestObject struct {
	objectivec.Object
}

func (o OS_nw_ws_requestObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_ws_requestObjectFromID constructs a [OS_nw_ws_requestObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_ws_requestObjectFromID(id objc.ID) OS_nw_ws_requestObject {
	return OS_nw_ws_requestObject{
		Object: objectivec.ObjectFromID(id),
	}
}
