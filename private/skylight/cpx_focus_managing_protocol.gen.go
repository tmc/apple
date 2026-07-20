// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusManaging protocol.
type CPXFocusManaging interface {
	objectivec.IObject

	// CleanupForProcessDeath protocol.
	CleanupForProcessDeath(death *CPSProcessRec)

	// FrontVisibleProcess protocol.
	FrontVisibleProcess() unsafe.Pointer

	// FrontmostProcess protocol.
	FrontmostProcess() unsafe.Pointer

	// GetProcessToBringForwardAtNextCheckin protocol.
	GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool

	// IsProcessPermittedToBeFrontmost protocol.
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool

	// IsProcessToBringForwardAtNextCheckin protocol.
	IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool

	// KeyThiefConnectionID protocol.
	KeyThiefConnectionID() uint32

	// ReleaseAllKeyThiefInstancesNotPermittedFrontmost protocol.
	ReleaseAllKeyThiefInstancesNotPermittedFrontmost()
}

// CPXFocusManagingObject wraps an existing Objective-C object that conforms to the CPXFocusManaging protocol.
type CPXFocusManagingObject struct {
	objectivec.Object
}

func (o CPXFocusManagingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXFocusManagingObjectFromID constructs a [CPXFocusManagingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXFocusManagingObjectFromID(id objc.ID) CPXFocusManagingObject {
	return CPXFocusManagingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXFocusManagingObject) CleanupForProcessDeath(death *CPSProcessRec) {
	objc.Send[struct{}](o.ID, objc.Sel("cleanupForProcessDeath:"), death)
}
func (o CPXFocusManagingObject) FocusController() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("focusController"))
	return objectivec.Object{ID: rv}
}
func (o CPXFocusManagingObject) FrontVisibleProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontVisibleProcess"))
	return rv
}
func (o CPXFocusManagingObject) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontmostProcess"))
	return rv
}
func (o CPXFocusManagingObject) GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (o CPXFocusManagingObject) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}
func (o CPXFocusManagingObject) IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (o CPXFocusManagingObject) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (o CPXFocusManagingObject) ReleaseAllKeyThiefInstancesNotPermittedFrontmost() {
	objc.Send[struct{}](o.ID, objc.Sel("releaseAllKeyThiefInstancesNotPermittedFrontmost"))
}
func (o CPXFocusManagingObject) SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("suppressDeferringPolicyUpdatesForReason:"), reason)
	return objectivec.Object{ID: rv}
}
