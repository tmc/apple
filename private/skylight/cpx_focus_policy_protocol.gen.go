// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusPolicy protocol.
type CPXFocusPolicyProtocol interface {
	objectivec.IObject

	// BringNextApplicationToFrontInternal protocol.
	BringNextApplicationToFrontInternal(internal *CPSProcessRec)

	// BringNextProcessToFront protocol.
	BringNextProcessToFront(front *CPSProcessRec)
}

// CPXFocusPolicyProtocolObject wraps an existing Objective-C object that conforms to the CPXFocusPolicyProtocol protocol.
type CPXFocusPolicyProtocolObject struct {
	objectivec.Object
}

func (o CPXFocusPolicyProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXFocusPolicyProtocolObjectFromID constructs a [CPXFocusPolicyProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXFocusPolicyProtocolObjectFromID(id objc.ID) CPXFocusPolicyProtocolObject {
	return CPXFocusPolicyProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXFocusPolicyProtocolObject) BringNextApplicationToFrontInternal(internal *CPSProcessRec) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("bringNextApplicationToFrontInternal:"), unsafe.Pointer(internal))
}
func (o CPXFocusPolicyProtocolObject) BringNextProcessToFront(front *CPSProcessRec) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("bringNextProcessToFront:"), unsafe.Pointer(front))
}
