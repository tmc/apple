// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZGuestAgent protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGuestAgent
type VZGuestAgent interface {
	objectivec.IObject
}

// VZGuestAgentObject wraps an existing Objective-C object that conforms to the VZGuestAgent protocol.
type VZGuestAgentObject struct {
	objectivec.Object
}
func (o VZGuestAgentObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZGuestAgentObjectFromID constructs a [VZGuestAgentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZGuestAgentObjectFromID(id objc.ID) VZGuestAgentObject {
	return VZGuestAgentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

