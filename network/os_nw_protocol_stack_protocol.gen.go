// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_protocol_stack protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_protocol_stack
type OS_nw_protocol_stack interface {
	objectivec.IObject
}

// OS_nw_protocol_stackObject wraps an existing Objective-C object that conforms to the OS_nw_protocol_stack protocol.
type OS_nw_protocol_stackObject struct {
	objectivec.Object
}

func (o OS_nw_protocol_stackObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_protocol_stackObjectFromID constructs a [OS_nw_protocol_stackObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_protocol_stackObjectFromID(id objc.ID) OS_nw_protocol_stackObject {
	return OS_nw_protocol_stackObject{
		Object: objectivec.ObjectFromID(id),
	}
}
