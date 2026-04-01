// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_nw_protocol_metadata protocol.
//
// See: https://developer.apple.com/documentation/Network/OS_nw_protocol_metadata
type OS_nw_protocol_metadata interface {
	objectivec.IObject
}

// OS_nw_protocol_metadataObject wraps an existing Objective-C object that conforms to the OS_nw_protocol_metadata protocol.
type OS_nw_protocol_metadataObject struct {
	objectivec.Object
}

func (o OS_nw_protocol_metadataObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_nw_protocol_metadataObjectFromID constructs a [OS_nw_protocol_metadataObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_nw_protocol_metadataObjectFromID(id objc.ID) OS_nw_protocol_metadataObject {
	return OS_nw_protocol_metadataObject{
		Object: objectivec.ObjectFromID(id),
	}
}
