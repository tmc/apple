// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_object protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_object
type OS_nw_object interface {
	objectivec.IObject
}

// OS_nw_objectObject wraps an existing Objective-C object that conforms to the OS_nw_object protocol.
type OS_nw_objectObject struct {
	objectivec.Object
}

func (o OS_nw_objectObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_objectObjectFromID constructs a [OS_nw_objectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_objectObjectFromID(id objc.ID) OS_nw_objectObject {
	return OS_nw_objectObject{
		Object: objectivec.ObjectFromID(id),
	}
}
