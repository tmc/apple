// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_protocol_definition protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_protocol_definition
type OS_nw_protocol_definition interface {
	objectivec.IObject
}

// OS_nw_protocol_definitionObject wraps an existing Objective-C object that conforms to the OS_nw_protocol_definition protocol.
type OS_nw_protocol_definitionObject struct {
	objectivec.Object
}

func (o OS_nw_protocol_definitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_protocol_definitionObjectFromID constructs a [OS_nw_protocol_definitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_protocol_definitionObjectFromID(id objc.ID) OS_nw_protocol_definitionObject {
	return OS_nw_protocol_definitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}
