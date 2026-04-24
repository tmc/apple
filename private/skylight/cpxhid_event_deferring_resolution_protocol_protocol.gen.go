// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXHIDEventDeferringResolutionProtocol protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol
type CPXHIDEventDeferringResolutionProtocol interface {
	objectivec.IObject

	// ConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/connectionID
	ConnectionID() uint32

	// Pid protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/pid
	Pid() int

	// ProcessRecord protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/processRecord
	ProcessRecord() unsafe.Pointer
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

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/connectionID
func (o CPXHIDEventDeferringResolutionProtocolObject) ConnectionID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("connectionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/environment
func (o CPXHIDEventDeferringResolutionProtocolObject) Environment() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("environment"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/pid
func (o CPXHIDEventDeferringResolutionProtocolObject) Pid() int {
	rv := objc.Send[int](o.ID, objc.Sel("pid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/processRecord
func (o CPXHIDEventDeferringResolutionProtocolObject) ProcessRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processRecord"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXHIDEventDeferringResolutionProtocol/token
func (o CPXHIDEventDeferringResolutionProtocolObject) Token() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("token"))
	return objectivec.Object{ID: rv}
}
