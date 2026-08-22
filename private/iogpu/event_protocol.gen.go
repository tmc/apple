// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLEvent protocol.
type MTLEvent interface {
	objectivec.IObject

	// Device protocol.
	Device() objectivec.IObject

	// Label protocol.
	Label() objectivec.IObject

	// SetLabel protocol.
	SetLabel(label objectivec.IObject)
}

// MTLEventObject wraps an existing Objective-C object that conforms to the MTLEvent protocol.
type MTLEventObject struct {
	objectivec.Object
}

func (o MTLEventObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLEventObjectFromID constructs a [MTLEventObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLEventObjectFromID(id objc.ID) MTLEventObject {
	return MTLEventObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLEventObject) Device() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("device"))
	return objectivec.Object{ID: rv}
}
func (o MTLEventObject) Label() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}
func (o MTLEventObject) SetLabel(label objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setLabel:"), label)
}
