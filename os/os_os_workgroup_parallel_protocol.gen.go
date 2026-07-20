// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_workgroup_parallel protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_workgroup_parallel-c.protocol
type OS_os_workgroup_parallelProtocol interface {
	objectivec.IObject
}

// OS_os_workgroup_parallelProtocolObject wraps an existing Objective-C object that conforms to the OS_os_workgroup_parallelProtocol protocol.
type OS_os_workgroup_parallelProtocolObject struct {
	objectivec.Object
}

func (o OS_os_workgroup_parallelProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_workgroup_parallelProtocolObjectFromID constructs a [OS_os_workgroup_parallelProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_workgroup_parallelProtocolObjectFromID(id objc.ID) OS_os_workgroup_parallelProtocolObject {
	return OS_os_workgroup_parallelProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}
