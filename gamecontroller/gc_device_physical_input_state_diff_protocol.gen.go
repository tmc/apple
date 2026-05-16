// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common functions for objects that contain the differences between a current and previous input state object.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputStateDiff
type GCDevicePhysicalInputStateDiff interface {
	objectivec.IObject

	// Returns whether the input value of an element changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputStateDiff/change(for:)
	ChangeForElement(element GCPhysicalInputElement) GCDevicePhysicalInputElementChange
}

// GCDevicePhysicalInputStateDiffObject wraps an existing Objective-C object that conforms to the GCDevicePhysicalInputStateDiff protocol.
type GCDevicePhysicalInputStateDiffObject struct {
	objectivec.Object
}

func (o GCDevicePhysicalInputStateDiffObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDevicePhysicalInputStateDiffObjectFromID constructs a [GCDevicePhysicalInputStateDiffObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDevicePhysicalInputStateDiffObjectFromID(id objc.ID) GCDevicePhysicalInputStateDiffObject {
	return GCDevicePhysicalInputStateDiffObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns whether the input value of an element changes.
//
// element: The element whose value changes.
//
// # Return Value
//
// Description of the change to the element.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputStateDiff/change(for:)
func (o GCDevicePhysicalInputStateDiffObject) ChangeForElement(element GCPhysicalInputElement) GCDevicePhysicalInputElementChange {
	rv := objc.Send[GCDevicePhysicalInputElementChange](o.ID, objc.Sel("changeForElement:"), element)
	return rv
}
