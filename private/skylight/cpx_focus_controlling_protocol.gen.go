// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusControlling protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling
type CPXFocusControlling interface {
	objectivec.IObject

	// AddToPermittedFrontList protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/addToPermittedFrontList:
	AddToPermittedFrontList(list objectivec.IObject) int16

	// RemoveFromPermittedFrontList protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/removeFromPermittedFrontList:
	RemoveFromPermittedFrontList(list objectivec.IObject) int16

	// SetFrontmostProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setFrontmostProcess:
	SetFrontmostProcess(process *CPSProcessRecRef) int16

	// SetKeyThiefConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setKeyThiefConnectionID:
	SetKeyThiefConnectionID(id uint32)

	// SetProcessToBringForwardAtNextCheckinPSN protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setProcessToBringForwardAtNextCheckinPSN:
	SetProcessToBringForwardAtNextCheckinPSN(psn objectivec.IObject) int16
}

// CPXFocusControllingObject wraps an existing Objective-C object that conforms to the CPXFocusControlling protocol.
type CPXFocusControllingObject struct {
	objectivec.Object
}

func (o CPXFocusControllingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXFocusControllingObjectFromID constructs a [CPXFocusControllingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXFocusControllingObjectFromID(id objc.ID) CPXFocusControllingObject {
	return CPXFocusControllingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/addToPermittedFrontList:
func (o CPXFocusControllingObject) AddToPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/removeFromPermittedFrontList:
func (o CPXFocusControllingObject) RemoveFromPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setFrontmostProcess:
func (o CPXFocusControllingObject) SetFrontmostProcess(process *CPSProcessRecRef) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("setFrontmostProcess:"), process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setKeyThiefConnectionID:
func (o CPXFocusControllingObject) SetKeyThiefConnectionID(id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("setKeyThiefConnectionID:"), id)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusControlling/setProcessToBringForwardAtNextCheckinPSN:
func (o CPXFocusControllingObject) SetProcessToBringForwardAtNextCheckinPSN(psn objectivec.IObject) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("setProcessToBringForwardAtNextCheckinPSN:"), psn)
	return rv
}
