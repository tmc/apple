// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_workgroup_interval protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_interval-c.protocol
type OS_os_workgroup_interval interface {
	objectivec.IObject
}

// OS_os_workgroup_intervalObject wraps an existing Objective-C object that conforms to the OS_os_workgroup_interval protocol.
type OS_os_workgroup_intervalObject struct {
	objectivec.Object
}

func (o OS_os_workgroup_intervalObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_workgroup_intervalObjectFromID constructs a [OS_os_workgroup_intervalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_workgroup_intervalObjectFromID(id objc.ID) OS_os_workgroup_intervalObject {
	return OS_os_workgroup_intervalObject{
		Object: objectivec.ObjectFromID(id),
	}
}
