// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_resolution_report protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_resolution_report
type OS_nw_resolution_report interface {
	objectivec.IObject
}

// OS_nw_resolution_reportObject wraps an existing Objective-C object that conforms to the OS_nw_resolution_report protocol.
type OS_nw_resolution_reportObject struct {
	objectivec.Object
}

func (o OS_nw_resolution_reportObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_resolution_reportObjectFromID constructs a [OS_nw_resolution_reportObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_resolution_reportObjectFromID(id objc.ID) OS_nw_resolution_reportObject {
	return OS_nw_resolution_reportObject{
		Object: objectivec.ObjectFromID(id),
	}
}
