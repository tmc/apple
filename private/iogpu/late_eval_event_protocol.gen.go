// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLLateEvalEvent protocol.
type MTLLateEvalEvent interface {
	objectivec.IObject
}

// MTLLateEvalEventObject wraps an existing Objective-C object that conforms to the MTLLateEvalEvent protocol.
type MTLLateEvalEventObject struct {
	objectivec.Object
}

func (o MTLLateEvalEventObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLLateEvalEventObjectFromID constructs a [MTLLateEvalEventObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLLateEvalEventObjectFromID(id objc.ID) MTLLateEvalEventObject {
	return MTLLateEvalEventObject{
		Object: objectivec.ObjectFromID(id),
	}
}
