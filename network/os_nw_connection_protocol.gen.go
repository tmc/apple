// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_connection protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_connection
type OS_nw_connection interface {
	objectivec.IObject
}

// OS_nw_connectionObject wraps an existing Objective-C object that conforms to the OS_nw_connection protocol.
type OS_nw_connectionObject struct {
	objectivec.Object
}

func (o OS_nw_connectionObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_connectionObjectFromID constructs a [OS_nw_connectionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_connectionObjectFromID(id objc.ID) OS_nw_connectionObject {
	return OS_nw_connectionObject{
		Object: objectivec.ObjectFromID(id),
	}
}
