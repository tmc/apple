// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusManagerDataSource protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource
type CPXFocusManagerDataSource interface {
	objectivec.IObject

	// AddToPermittedFrontList protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/addToPermittedFrontList:
	AddToPermittedFrontList(list objectivec.IObject) int16

	// FrontmostProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/frontmostProcess
	FrontmostProcess() unsafe.Pointer

	// GetProcessToBringForwardAtNextCheckin protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/getProcessToBringForwardAtNextCheckin:
	GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool

	// IsProcessPermittedToBeFrontmost protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/isProcessPermittedToBeFrontmost:
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool

	// IsProcessToBringForwardAtNextCheckin protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/isProcessToBringForwardAtNextCheckin:
	IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool

	// KeyThiefConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/keyThiefConnectionID
	KeyThiefConnectionID() uint32

	// RemoveFromPermittedFrontList protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/removeFromPermittedFrontList:
	RemoveFromPermittedFrontList(list objectivec.IObject) int16

	// SetFrontmostProcess protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setFrontmostProcess:
	SetFrontmostProcess(process *CPSProcessRecRef) int16

	// SetKeyThiefConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setKeyThiefConnectionID:
	SetKeyThiefConnectionID(id uint32)

	// SetProcessToBringForwardAtNextCheckin protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setProcessToBringForwardAtNextCheckin:
	SetProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) int
}

// CPXFocusManagerDataSourceObject wraps an existing Objective-C object that conforms to the CPXFocusManagerDataSource protocol.
type CPXFocusManagerDataSourceObject struct {
	objectivec.Object
}

func (o CPXFocusManagerDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXFocusManagerDataSourceObjectFromID constructs a [CPXFocusManagerDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXFocusManagerDataSourceObjectFromID(id objc.ID) CPXFocusManagerDataSourceObject {
	return CPXFocusManagerDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/addToPermittedFrontList:
func (o CPXFocusManagerDataSourceObject) AddToPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/frontmostProcess
func (o CPXFocusManagerDataSourceObject) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontmostProcess"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/getProcessToBringForwardAtNextCheckin:
func (o CPXFocusManagerDataSourceObject) GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/isProcessPermittedToBeFrontmost:
func (o CPXFocusManagerDataSourceObject) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/isProcessToBringForwardAtNextCheckin:
func (o CPXFocusManagerDataSourceObject) IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/keyThiefConnectionID
func (o CPXFocusManagerDataSourceObject) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/removeFromPermittedFrontList:
func (o CPXFocusManagerDataSourceObject) RemoveFromPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setFrontmostProcess:
func (o CPXFocusManagerDataSourceObject) SetFrontmostProcess(process *CPSProcessRecRef) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("setFrontmostProcess:"), process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setKeyThiefConnectionID:
func (o CPXFocusManagerDataSourceObject) SetKeyThiefConnectionID(id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("setKeyThiefConnectionID:"), id)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManagerDataSource/setProcessToBringForwardAtNextCheckin:
func (o CPXFocusManagerDataSourceObject) SetProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("setProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
