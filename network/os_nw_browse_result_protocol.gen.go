// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_browse_result protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_browse_result
type OS_nw_browse_result interface {
	objectivec.IObject
}

// OS_nw_browse_resultObject wraps an existing Objective-C object that conforms to the OS_nw_browse_result protocol.
type OS_nw_browse_resultObject struct {
	objectivec.Object
}

func (o OS_nw_browse_resultObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_browse_resultObjectFromID constructs a [OS_nw_browse_resultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_browse_resultObjectFromID(id objc.ID) OS_nw_browse_resultObject {
	return OS_nw_browse_resultObject{
		Object: objectivec.ObjectFromID(id),
	}
}
