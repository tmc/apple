// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_txt_record protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_txt_record
type OS_nw_txt_record interface {
	objectivec.IObject
}

// OS_nw_txt_recordObject wraps an existing Objective-C object that conforms to the OS_nw_txt_record protocol.
type OS_nw_txt_recordObject struct {
	objectivec.Object
}

func (o OS_nw_txt_recordObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_txt_recordObjectFromID constructs a [OS_nw_txt_recordObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_txt_recordObjectFromID(id objc.ID) OS_nw_txt_recordObject {
	return OS_nw_txt_recordObject{
		Object: objectivec.ObjectFromID(id),
	}
}
