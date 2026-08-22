// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusControlling protocol.
type CPXFocusControlling interface {
	objectivec.IObject

	// AddToPermittedFrontList protocol.
	AddToPermittedFrontList(list CPSProcessSerNum) int16

	// RemoveFromPermittedFrontList protocol.
	RemoveFromPermittedFrontList(list CPSProcessSerNum) int16

	// SetFrontmostProcess protocol.
	SetFrontmostProcess(process *CPSProcessRec) int16

	// SetKeyThiefConnectionID protocol.
	SetKeyThiefConnectionID(id uint32)

	// SetProcessToBringForwardAtNextCheckinPSN protocol.
	SetProcessToBringForwardAtNextCheckinPSN(psn CPSProcessSerNum) int16
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

func (o CPXFocusControllingObject) AddToPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.SendIfResponds[int16](o.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}
func (o CPXFocusControllingObject) RemoveFromPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.SendIfResponds[int16](o.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}
func (o CPXFocusControllingObject) SetFrontmostProcess(process *CPSProcessRec) int16 {
	rv := objc.SendIfResponds[int16](o.ID, objc.Sel("setFrontmostProcess:"), unsafe.Pointer(process))
	return rv
}
func (o CPXFocusControllingObject) SetKeyThiefConnectionID(id uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setKeyThiefConnectionID:"), id)
}
func (o CPXFocusControllingObject) SetProcessToBringForwardAtNextCheckinPSN(psn CPSProcessSerNum) int16 {
	rv := objc.SendIfResponds[int16](o.ID, objc.Sel("setProcessToBringForwardAtNextCheckinPSN:"), psn)
	return rv
}
