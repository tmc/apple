// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The general interface for classes that can act as containers in an object hierarchy.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLObjectContainerComponent
type MDLObjectContainerComponent interface {
	objectivec.IObject
	MDLComponent
	foundation.NSFastEnumeration
}

// MDLObjectContainerComponentObject wraps an existing Objective-C object that conforms to the MDLObjectContainerComponent protocol.
type MDLObjectContainerComponentObject struct {
	foundation.NSFastEnumerationObject
}

func (o MDLObjectContainerComponentObject) BaseObject() objectivec.Object {
	return o.NSFastEnumerationObject.BaseObject()
}

// MDLObjectContainerComponentObjectFromID constructs a [MDLObjectContainerComponentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLObjectContainerComponentObjectFromID(id objc.ID) MDLObjectContainerComponentObject {
	return MDLObjectContainerComponentObject{
		NSFastEnumerationObject: foundation.NSFastEnumerationObjectFromID(id),
	}
}
