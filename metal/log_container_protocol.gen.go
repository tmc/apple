// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A collection of logged messages, created when a Metal device runs a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogContainer-c.protocol
type MTLLogContainer interface {
	objectivec.IObject
	foundation.NSFastEnumeration
}

// MTLLogContainerObject wraps an existing Objective-C object that conforms to the MTLLogContainer protocol.
type MTLLogContainerObject struct {
	foundation.NSFastEnumerationObject
}

func (o MTLLogContainerObject) BaseObject() objectivec.Object {
	return o.NSFastEnumerationObject.BaseObject()
}

// MTLLogContainerObjectFromID constructs a [MTLLogContainerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLLogContainerObjectFromID(id objc.ID) MTLLogContainerObject {
	return MTLLogContainerObject{
		NSFastEnumerationObject: foundation.NSFastEnumerationObjectFromID(id),
	}
}
