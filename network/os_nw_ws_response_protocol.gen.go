// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_ws_response protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_ws_response
type OS_nw_ws_response interface {
	objectivec.IObject
}

// OS_nw_ws_responseObject wraps an existing Objective-C object that conforms to the OS_nw_ws_response protocol.
type OS_nw_ws_responseObject struct {
	objectivec.Object
}

func (o OS_nw_ws_responseObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_ws_responseObjectFromID constructs a [OS_nw_ws_responseObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_ws_responseObjectFromID(id objc.ID) OS_nw_ws_responseObject {
	return OS_nw_ws_responseObject{
		Object: objectivec.ObjectFromID(id),
	}
}
