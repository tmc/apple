// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_workgroup_interval protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.protocol
type OS_os_workgroup_intervalProtocol interface {
	objectivec.IObject
}

// OS_os_workgroup_intervalProtocolObject wraps an existing Objective-C object that conforms to the OS_os_workgroup_intervalProtocol protocol.
type OS_os_workgroup_intervalProtocolObject struct {
	objectivec.Object
}

func (o OS_os_workgroup_intervalProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_workgroup_intervalProtocolObjectFromID constructs a [OS_os_workgroup_intervalProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_workgroup_intervalProtocolObjectFromID(id objc.ID) OS_os_workgroup_intervalProtocolObject {
	return OS_os_workgroup_intervalProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}
