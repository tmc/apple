// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXFocusManagerDataSource protocol.
type CPXFocusManagerDataSource interface {
	objectivec.IObject

	// AddToPermittedFrontList protocol.
	AddToPermittedFrontList(list CPSProcessSerNum) int16

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

	// RemoveFromPermittedFrontList protocol.
	RemoveFromPermittedFrontList(list CPSProcessSerNum) int16

	// SetFrontmostProcess protocol.
	SetFrontmostProcess(process *CPSProcessRec) int16

	// SetKeyThiefConnectionID protocol.
	SetKeyThiefConnectionID(id uint32)

	// SetProcessToBringForwardAtNextCheckin protocol.
	SetProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) int
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

func (o CPXFocusManagerDataSourceObject) AddToPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}
func (o CPXFocusManagerDataSourceObject) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("frontmostProcess"))
	return rv
}
func (o CPXFocusManagerDataSourceObject) GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (o CPXFocusManagerDataSourceObject) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}
func (o CPXFocusManagerDataSourceObject) IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (o CPXFocusManagerDataSourceObject) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (o CPXFocusManagerDataSourceObject) RemoveFromPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}
func (o CPXFocusManagerDataSourceObject) SetFrontmostProcess(process *CPSProcessRec) int16 {
	rv := objc.Send[int16](o.ID, objc.Sel("setFrontmostProcess:"), process)
	return rv
}
func (o CPXFocusManagerDataSourceObject) SetKeyThiefConnectionID(id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("setKeyThiefConnectionID:"), id)
}
func (o CPXFocusManagerDataSourceObject) SetProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) int {
	rv := objc.Send[int](o.ID, objc.Sel("setProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
