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

	// FocusController protocol.
	FocusController() objectivec.IObject

	// FrontVisibleProcess protocol.
	FrontVisibleProcess() *CPSProcessRec

	// FrontmostProcess protocol.
	FrontmostProcess() *CPSProcessRec

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

	// SuppressDeferringPolicyUpdatesForReason protocol.
	SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject
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
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cleanupForProcessDeath:"), unsafe.Pointer(death))
}
func (o CPXFocusManagingObject) FocusController() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("focusController"))
	return objectivec.Object{ID: rv}
}
func (o CPXFocusManagingObject) FrontVisibleProcess() *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("frontVisibleProcess"))
	return (*CPSProcessRec)(rv)
}
func (o CPXFocusManagingObject) FrontmostProcess() *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("frontmostProcess"))
	return (*CPSProcessRec)(rv)
}
func (o CPXFocusManagingObject) GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), unsafe.Pointer(checkin))
	return rv
}
func (o CPXFocusManagingObject) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), unsafe.Pointer(frontmost))
	return rv
}
func (o CPXFocusManagingObject) IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (o CPXFocusManagingObject) KeyThiefConnectionID() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (o CPXFocusManagingObject) ReleaseAllKeyThiefInstancesNotPermittedFrontmost() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("releaseAllKeyThiefInstancesNotPermittedFrontmost"))
}
func (o CPXFocusManagingObject) SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("suppressDeferringPolicyUpdatesForReason:"), reason)
	return objectivec.Object{ID: rv}
}
