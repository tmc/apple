// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusPolicy protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusPolicy
type CPXFocusPolicyProtocol interface {
	objectivec.IObject

	// BringNextApplicationToFrontInternal protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusPolicy/bringNextApplicationToFrontInternal:
	BringNextApplicationToFrontInternal(internal *CPSProcessRecRef)

	// BringNextProcessToFront protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusPolicy/bringNextProcessToFront:
	BringNextProcessToFront(front *CPSProcessRecRef)
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

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusPolicy/bringNextApplicationToFrontInternal:
func (o CPXFocusPolicyProtocolObject) BringNextApplicationToFrontInternal(internal *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("bringNextApplicationToFrontInternal:"), internal)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusPolicy/bringNextProcessToFront:
func (o CPXFocusPolicyProtocolObject) BringNextProcessToFront(front *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("bringNextProcessToFront:"), front)
}
