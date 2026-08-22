// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_workgroup_parallel protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.protocol
type OS_os_workgroup_parallel interface {
	objectivec.IObject
}

// OS_os_workgroup_parallelObject wraps an existing Objective-C object that conforms to the OS_os_workgroup_parallel protocol.
type OS_os_workgroup_parallelObject struct {
	objectivec.Object
}

func (o OS_os_workgroup_parallelObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_workgroup_parallelObjectFromID constructs a [OS_os_workgroup_parallelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_workgroup_parallelObjectFromID(id objc.ID) OS_os_workgroup_parallelObject {
	return OS_os_workgroup_parallelObject{
		Object: objectivec.ObjectFromID(id),
	}
}
