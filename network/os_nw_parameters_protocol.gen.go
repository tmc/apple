// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_parameters protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_parameters
type OS_nw_parameters interface {
	objectivec.IObject
}

// OS_nw_parametersObject wraps an existing Objective-C object that conforms to the OS_nw_parameters protocol.
type OS_nw_parametersObject struct {
	objectivec.Object
}

func (o OS_nw_parametersObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_parametersObjectFromID constructs a [OS_nw_parametersObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_parametersObjectFromID(id objc.ID) OS_nw_parametersObject {
	return OS_nw_parametersObject{
		Object: objectivec.ObjectFromID(id),
	}
}
