// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_browser protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_browser
type OS_nw_browser interface {
	objectivec.IObject
}

// OS_nw_browserObject wraps an existing Objective-C object that conforms to the OS_nw_browser protocol.
type OS_nw_browserObject struct {
	objectivec.Object
}

func (o OS_nw_browserObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_browserObjectFromID constructs a [OS_nw_browserObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_browserObjectFromID(id objc.ID) OS_nw_browserObject {
	return OS_nw_browserObject{
		Object: objectivec.ObjectFromID(id),
	}
}
