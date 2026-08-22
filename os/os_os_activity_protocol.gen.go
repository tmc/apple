// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_activity protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_activity
type OS_os_activity interface {
	objectivec.IObject
}

// OS_os_activityObject wraps an existing Objective-C object that conforms to the OS_os_activity protocol.
type OS_os_activityObject struct {
	objectivec.Object
}

func (o OS_os_activityObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_activityObjectFromID constructs a [OS_os_activityObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_activityObjectFromID(id objc.ID) OS_os_activityObject {
	return OS_os_activityObject{
		Object: objectivec.ObjectFromID(id),
	}
}
