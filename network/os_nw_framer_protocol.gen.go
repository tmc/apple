// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_framer protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_framer
type OS_nw_framer interface {
	objectivec.IObject
}

// OS_nw_framerObject wraps an existing Objective-C object that conforms to the OS_nw_framer protocol.
type OS_nw_framerObject struct {
	objectivec.Object
}

func (o OS_nw_framerObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_framerObjectFromID constructs a [OS_nw_framerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_framerObjectFromID(id objc.ID) OS_nw_framerObject {
	return OS_nw_framerObject{
		Object: objectivec.ObjectFromID(id),
	}
}
