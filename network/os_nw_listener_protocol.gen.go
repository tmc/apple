// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_listener protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_listener
type OS_nw_listener interface {
	objectivec.IObject
}

// OS_nw_listenerObject wraps an existing Objective-C object that conforms to the OS_nw_listener protocol.
type OS_nw_listenerObject struct {
	objectivec.Object
}

func (o OS_nw_listenerObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_listenerObjectFromID constructs a [OS_nw_listenerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_listenerObjectFromID(id objc.ID) OS_nw_listenerObject {
	return OS_nw_listenerObject{
		Object: objectivec.ObjectFromID(id),
	}
}
