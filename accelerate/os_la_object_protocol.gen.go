// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OS_la_object protocol.
//
// See: https://developer.apple.com/documentation/Accelerate/OS_la_object
type OS_la_object interface {
	objectivec.IObject
}

// OS_la_objectObject wraps an existing Objective-C object that conforms to the OS_la_object protocol.
type OS_la_objectObject struct {
	objectivec.Object
}

func (o OS_la_objectObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_la_objectObjectFromID constructs a [OS_la_objectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_la_objectObjectFromID(id objc.ID) OS_la_objectObject {
	return OS_la_objectObject{
		Object: objectivec.ObjectFromID(id),
	}
}
