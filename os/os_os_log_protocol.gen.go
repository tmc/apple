// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_os_log protocol.
//
// See: https://developer.apple.com/documentation/os/OS_os_log
type OS_os_log interface {
	objectivec.IObject
}

// OS_os_logObject wraps an existing Objective-C object that conforms to the OS_os_log protocol.
type OS_os_logObject struct {
	objectivec.Object
}

func (o OS_os_logObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_os_logObjectFromID constructs a [OS_os_logObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_os_logObjectFromID(id objc.ID) OS_os_logObject {
	return OS_os_logObject{
		Object: objectivec.ObjectFromID(id),
	}
}
