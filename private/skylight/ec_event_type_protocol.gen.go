// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECEventType protocol.
type ECEventType interface {
	objectivec.IObject

	// CgSubType protocol.
	CgSubType() uint64

	// CgType protocol.
	CgType() uint32

	// HidType protocol.
	HidType() uint32

	// IsCGType protocol.
	IsCGType() bool
}

// ECEventTypeObject wraps an existing Objective-C object that conforms to the ECEventType protocol.
type ECEventTypeObject struct {
	objectivec.Object
}

func (o ECEventTypeObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECEventTypeObjectFromID constructs a [ECEventTypeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECEventTypeObjectFromID(id objc.ID) ECEventTypeObject {
	return ECEventTypeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECEventTypeObject) CgSubType() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("cgSubType"))
	return rv
}
func (o ECEventTypeObject) CgType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("cgType"))
	return rv
}
func (o ECEventTypeObject) HidType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("hidType"))
	return rv
}
func (o ECEventTypeObject) IsCGType() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isCGType"))
	return rv
}
