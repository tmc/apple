// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXHIDEventDeferringResolutionProtocol protocol.
type CPXHIDEventDeferringResolutionProtocol interface {
	objectivec.IObject

	// ConnectionID protocol.
	ConnectionID() uint32

	// Environment protocol.
	Environment() objectivec.IObject

	// Pid protocol.
	Pid() int

	// ProcessRecord protocol.
	ProcessRecord() *CPSProcessRec

	// Token protocol.
	Token() objectivec.IObject
}

// CPXHIDEventDeferringResolutionProtocolObject wraps an existing Objective-C object that conforms to the CPXHIDEventDeferringResolutionProtocol protocol.
type CPXHIDEventDeferringResolutionProtocolObject struct {
	objectivec.Object
}

func (o CPXHIDEventDeferringResolutionProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXHIDEventDeferringResolutionProtocolObjectFromID constructs a [CPXHIDEventDeferringResolutionProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXHIDEventDeferringResolutionProtocolObjectFromID(id objc.ID) CPXHIDEventDeferringResolutionProtocolObject {
	return CPXHIDEventDeferringResolutionProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXHIDEventDeferringResolutionProtocolObject) ConnectionID() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("connectionID"))
	return rv
}
func (o CPXHIDEventDeferringResolutionProtocolObject) Environment() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("environment"))
	return objectivec.Object{ID: rv}
}
func (o CPXHIDEventDeferringResolutionProtocolObject) Pid() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("pid"))
	return rv
}
func (o CPXHIDEventDeferringResolutionProtocolObject) ProcessRecord() *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processRecord"))
	return (*CPSProcessRec)(rv)
}
func (o CPXHIDEventDeferringResolutionProtocolObject) Token() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("token"))
	return objectivec.Object{ID: rv}
}
