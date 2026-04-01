// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_path protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_path
type OS_nw_path interface {
	objectivec.IObject
}

// OS_nw_pathObject wraps an existing Objective-C object that conforms to the OS_nw_path protocol.
type OS_nw_pathObject struct {
	objectivec.Object
}

func (o OS_nw_pathObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_pathObjectFromID constructs a [OS_nw_pathObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_pathObjectFromID(id objc.ID) OS_nw_pathObject {
	return OS_nw_pathObject{
		Object: objectivec.ObjectFromID(id),
	}
}
