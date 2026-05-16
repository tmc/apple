// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods that provide information on the hypervisor state.
//
// See: https://developer.apple.com/documentation/Hypervisor/OS_hv_gic_state
type OS_hv_gic_state interface {
	objectivec.IObject
}

// OS_hv_gic_stateObject wraps an existing Objective-C object that conforms to the OS_hv_gic_state protocol.
type OS_hv_gic_stateObject struct {
	objectivec.Object
}

func (o OS_hv_gic_stateObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_hv_gic_stateObjectFromID constructs a [OS_hv_gic_stateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_hv_gic_stateObjectFromID(id objc.ID) OS_hv_gic_stateObject {
	return OS_hv_gic_stateObject{
		Object: objectivec.ObjectFromID(id),
	}
}
