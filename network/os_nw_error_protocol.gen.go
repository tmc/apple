// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_error protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_error
type OS_nw_error interface {
	objectivec.IObject
}

// OS_nw_errorObject wraps an existing Objective-C object that conforms to the OS_nw_error protocol.
type OS_nw_errorObject struct {
	objectivec.Object
}

func (o OS_nw_errorObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_errorObjectFromID constructs a [OS_nw_errorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_errorObjectFromID(id objc.ID) OS_nw_errorObject {
	return OS_nw_errorObject{
		Object: objectivec.ObjectFromID(id),
	}
}
