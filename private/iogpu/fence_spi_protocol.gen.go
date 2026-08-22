// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLFenceSPI protocol.
type MTLFenceSPI interface {
	objectivec.IObject
}

// MTLFenceSPIObject wraps an existing Objective-C object that conforms to the MTLFenceSPI protocol.
type MTLFenceSPIObject struct {
	objectivec.Object
}

func (o MTLFenceSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFenceSPIObjectFromID constructs a [MTLFenceSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFenceSPIObjectFromID(id objc.ID) MTLFenceSPIObject {
	return MTLFenceSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}
