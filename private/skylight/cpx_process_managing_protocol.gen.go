// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXProcessManaging protocol.
type CPXProcessManaging interface {
	objectivec.IObject

	// IsPSNEqualToPSN protocol.
	IsPSNEqualToPSN(psn CPSProcessSerNum, psn2 CPSProcessSerNum) bool

	// IsValidConnectionIDForPSN protocol.
	IsValidConnectionIDForPSN(id uint32, psn CPSProcessSerNum) bool

	// ProcessForPID protocol.
	ProcessForPID(pid int) *CPSProcessRec

	// ProcessForPSN protocol.
	ProcessForPSN(psn CPSProcessSerNum) *CPSProcessRec

	// ProcessOwningConnection protocol.
	ProcessOwningConnection(connection *CGXConnection) *CPSProcessRec

	// ProcessOwningConnectionID protocol.
	ProcessOwningConnectionID(id uint32) *CPSProcessRec

	// ProcessPendingKill protocol.
	ProcessPendingKill() *CPSProcessRec

	// ProcessRepresentedByConnection protocol.
	ProcessRepresentedByConnection(connection *CGXConnection) *CPSProcessRec

	// ProcessRepresentedByConnectionID protocol.
	ProcessRepresentedByConnectionID(id uint32) *CPSProcessRec

	// SetProcessPendingKill protocol.
	SetProcessPendingKill(kill *CPSProcessRec)

	// UpdateProcessApplicationTypeIfNecessary protocol.
	UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRec) byte
}

// CPXProcessManagingObject wraps an existing Objective-C object that conforms to the CPXProcessManaging protocol.
type CPXProcessManagingObject struct {
	objectivec.Object
}

func (o CPXProcessManagingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXProcessManagingObjectFromID constructs a [CPXProcessManagingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXProcessManagingObjectFromID(id objc.ID) CPXProcessManagingObject {
	return CPXProcessManagingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXProcessManagingObject) IsPSNEqualToPSN(psn CPSProcessSerNum, psn2 CPSProcessSerNum) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isPSN:equalToPSN:"), psn, psn2)
	return rv
}
func (o CPXProcessManagingObject) IsValidConnectionIDForPSN(id uint32, psn CPSProcessSerNum) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isValidConnectionID:forPSN:"), id, psn)
	return rv
}
func (o CPXProcessManagingObject) ProcessForPID(pid int) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processForPID:"), pid)
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessForPSN(psn CPSProcessSerNum) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processForPSN:"), psn)
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessOwningConnection(connection *CGXConnection) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processOwningConnection:"), unsafe.Pointer(connection))
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessOwningConnectionID(id uint32) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processOwningConnectionID:"), id)
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessPendingKill() *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processPendingKill"))
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessRepresentedByConnection(connection *CGXConnection) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnection:"), unsafe.Pointer(connection))
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) ProcessRepresentedByConnectionID(id uint32) *CPSProcessRec {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnectionID:"), id)
	return (*CPSProcessRec)(rv)
}
func (o CPXProcessManagingObject) SetProcessPendingKill(kill *CPSProcessRec) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setProcessPendingKill:"), unsafe.Pointer(kill))
}
func (o CPXProcessManagingObject) UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRec) byte {
	rv := objc.SendIfResponds[byte](o.ID, objc.Sel("updateProcessApplicationTypeIfNecessary:"), unsafe.Pointer(necessary))
	return rv
}
