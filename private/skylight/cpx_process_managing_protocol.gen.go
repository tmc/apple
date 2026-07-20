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
	ProcessForPID(pid int) unsafe.Pointer

	// ProcessForPSN protocol.
	ProcessForPSN(psn CPSProcessSerNum) unsafe.Pointer

	// ProcessOwningConnection protocol.
	ProcessOwningConnection(connection *CGXConnection) unsafe.Pointer

	// ProcessOwningConnectionID protocol.
	ProcessOwningConnectionID(id uint32) unsafe.Pointer

	// ProcessPendingKill protocol.
	ProcessPendingKill() unsafe.Pointer

	// ProcessRepresentedByConnection protocol.
	ProcessRepresentedByConnection(connection *CGXConnection) unsafe.Pointer

	// ProcessRepresentedByConnectionID protocol.
	ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer

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
	rv := objc.Send[bool](o.ID, objc.Sel("isPSN:equalToPSN:"), psn, psn2)
	return rv
}
func (o CPXProcessManagingObject) IsValidConnectionIDForPSN(id uint32, psn CPSProcessSerNum) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isValidConnectionID:forPSN:"), id, psn)
	return rv
}
func (o CPXProcessManagingObject) ProcessForPID(pid int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processForPID:"), pid)
	return rv
}
func (o CPXProcessManagingObject) ProcessForPSN(psn CPSProcessSerNum) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processForPSN:"), psn)
	return rv
}
func (o CPXProcessManagingObject) ProcessOwningConnection(connection *CGXConnection) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processOwningConnection:"), connection)
	return rv
}
func (o CPXProcessManagingObject) ProcessOwningConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processOwningConnectionID:"), id)
	return rv
}
func (o CPXProcessManagingObject) ProcessPendingKill() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processPendingKill"))
	return rv
}
func (o CPXProcessManagingObject) ProcessRepresentedByConnection(connection *CGXConnection) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnection:"), connection)
	return rv
}
func (o CPXProcessManagingObject) ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnectionID:"), id)
	return rv
}
func (o CPXProcessManagingObject) SetProcessPendingKill(kill *CPSProcessRec) {
	objc.Send[struct{}](o.ID, objc.Sel("setProcessPendingKill:"), kill)
}
func (o CPXProcessManagingObject) UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRec) byte {
	rv := objc.Send[byte](o.ID, objc.Sel("updateProcessApplicationTypeIfNecessary:"), necessary)
	return rv
}
