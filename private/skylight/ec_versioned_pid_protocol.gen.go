// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECVersionedPID protocol.
type ECVersionedPID interface {
	objectivec.IObject

	// Pid protocol.
	Pid() int

	// Version protocol.
	Version() uint32
}

// ECVersionedPIDObject wraps an existing Objective-C object that conforms to the ECVersionedPID protocol.
type ECVersionedPIDObject struct {
	objectivec.Object
}

func (o ECVersionedPIDObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECVersionedPIDObjectFromID constructs a [ECVersionedPIDObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECVersionedPIDObjectFromID(id objc.ID) ECVersionedPIDObject {
	return ECVersionedPIDObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECVersionedPIDObject) Pid() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("pid"))
	return rv
}
func (o ECVersionedPIDObject) Version() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("version"))
	return rv
}
