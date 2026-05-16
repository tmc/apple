// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Creates a virtual machine configuration object.
//
// See: https://developer.apple.com/documentation/Hypervisor/OS_hv_vm_config
type OS_hv_vm_config interface {
	objectivec.IObject
}

// OS_hv_vm_configObject wraps an existing Objective-C object that conforms to the OS_hv_vm_config protocol.
type OS_hv_vm_configObject struct {
	objectivec.Object
}

func (o OS_hv_vm_configObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_hv_vm_configObjectFromID constructs a [OS_hv_vm_configObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_hv_vm_configObjectFromID(id objc.ID) OS_hv_vm_configObject {
	return OS_hv_vm_configObject{
		Object: objectivec.ObjectFromID(id),
	}
}
