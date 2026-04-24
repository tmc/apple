// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECEventType protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECEventType
type ECEventType interface {
	objectivec.IObject

	// CgSubType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECEventType/cgSubType
	CgSubType() uint64

	// CgType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECEventType/cgType
	CgType() uint32

	// HidType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECEventType/hidType
	HidType() uint32

	// IsCGType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECEventType/isCGType
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

// See: https://developer.apple.com/documentation/SkyLight/ECEventType/cgSubType
func (o ECEventTypeObject) CgSubType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("cgSubType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventType/cgType
func (o ECEventTypeObject) CgType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("cgType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventType/hidType
func (o ECEventTypeObject) HidType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("hidType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventType/isCGType
func (o ECEventTypeObject) IsCGType() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isCGType"))
	return rv
}
