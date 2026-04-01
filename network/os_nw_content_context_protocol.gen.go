// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_content_context protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_content_context
type OS_nw_content_context interface {
	objectivec.IObject
}

// OS_nw_content_contextObject wraps an existing Objective-C object that conforms to the OS_nw_content_context protocol.
type OS_nw_content_contextObject struct {
	objectivec.Object
}

func (o OS_nw_content_contextObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_content_contextObjectFromID constructs a [OS_nw_content_contextObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_content_contextObjectFromID(id objc.ID) OS_nw_content_contextObject {
	return OS_nw_content_contextObject{
		Object: objectivec.ObjectFromID(id),
	}
}
