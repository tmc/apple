// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_protocol_options protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_protocol_options
type OS_nw_protocol_options interface {
	objectivec.IObject
}

// OS_nw_protocol_optionsObject wraps an existing Objective-C object that conforms to the OS_nw_protocol_options protocol.
type OS_nw_protocol_optionsObject struct {
	objectivec.Object
}

func (o OS_nw_protocol_optionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_protocol_optionsObjectFromID constructs a [OS_nw_protocol_optionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_protocol_optionsObjectFromID(id objc.ID) OS_nw_protocol_optionsObject {
	return OS_nw_protocol_optionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}
