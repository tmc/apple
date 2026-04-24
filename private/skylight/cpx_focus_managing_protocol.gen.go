// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusManaging protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging
type CPXFocusManaging interface {
	objectivec.IObject

	// CleanupForProcessDeath protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/cleanupForProcessDeath:
	CleanupForProcessDeath(death *CPSProcessRecRef)

	// FrontVisibleProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/frontVisibleProcess
	FrontVisibleProcess() unsafe.Pointer

	// FrontmostProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/frontmostProcess
	FrontmostProcess() unsafe.Pointer

	// GetProcessToBringForwardAtNextCheckin protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/getProcessToBringForwardAtNextCheckin:
	GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool

	// IsProcessPermittedToBeFrontmost protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/isProcessPermittedToBeFrontmost:
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool

	// IsProcessToBringForwardAtNextCheckin protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/isProcessToBringForwardAtNextCheckin:
	IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool

	// KeyThiefConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/keyThiefConnectionID
	KeyThiefConnectionID() uint32

	// ReleaseAllKeyThiefInstancesNotPermittedFrontmost protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/releaseAllKeyThiefInstancesNotPermittedFrontmost
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

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/cleanupForProcessDeath:
func (o CPXFocusManagingObject) CleanupForProcessDeath(death *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("cleanupForProcessDeath:"), death)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/focusController
func (o CPXFocusManagingObject) FocusController() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("focusController"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/frontVisibleProcess
func (o CPXFocusManagingObject) FrontVisibleProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontVisibleProcess"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/frontmostProcess
func (o CPXFocusManagingObject) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontmostProcess"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/getProcessToBringForwardAtNextCheckin:
func (o CPXFocusManagingObject) GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/isProcessPermittedToBeFrontmost:
func (o CPXFocusManagingObject) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/isProcessToBringForwardAtNextCheckin:
func (o CPXFocusManagingObject) IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/keyThiefConnectionID
func (o CPXFocusManagingObject) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/releaseAllKeyThiefInstancesNotPermittedFrontmost
func (o CPXFocusManagingObject) ReleaseAllKeyThiefInstancesNotPermittedFrontmost() {
	objc.Send[struct{}](o.ID, objc.Sel("releaseAllKeyThiefInstancesNotPermittedFrontmost"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManaging/suppressDeferringPolicyUpdatesForReason:
func (o CPXFocusManagingObject) SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("suppressDeferringPolicyUpdatesForReason:"), reason)
	return objectivec.Object{ID: rv}
}
