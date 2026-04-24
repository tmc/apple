// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXProcessManaging protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging
type CPXProcessManaging interface {
	objectivec.IObject

	// IsPSNEqualToPSN protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/isPSN:equalToPSN:
	IsPSNEqualToPSN(psn objectivec.IObject, psn2 objectivec.IObject) bool

	// IsValidConnectionIDForPSN protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/isValidConnectionID:forPSN:
	IsValidConnectionIDForPSN(id uint32, psn objectivec.IObject) bool

	// ProcessForPID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processForPID:
	ProcessForPID(pid int) unsafe.Pointer

	// ProcessForPSN protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processForPSN:
	ProcessForPSN(psn objectivec.IObject) unsafe.Pointer

	// ProcessOwningConnection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processOwningConnection:
	ProcessOwningConnection(connection unsafe.Pointer) unsafe.Pointer

	// ProcessOwningConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processOwningConnectionID:
	ProcessOwningConnectionID(id uint32) unsafe.Pointer

	// ProcessPendingKill protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processPendingKill
	ProcessPendingKill() unsafe.Pointer

	// ProcessRepresentedByConnection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processRepresentedByConnection:
	ProcessRepresentedByConnection(connection unsafe.Pointer) unsafe.Pointer

	// ProcessRepresentedByConnectionID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processRepresentedByConnectionID:
	ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer

	// SetProcessPendingKill protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/setProcessPendingKill:
	SetProcessPendingKill(kill *CPSProcessRecRef)

	// UpdateProcessApplicationTypeIfNecessary protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/updateProcessApplicationTypeIfNecessary:
	UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRecRef) byte
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

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/isPSN:equalToPSN:
func (o CPXProcessManagingObject) IsPSNEqualToPSN(psn objectivec.IObject, psn2 objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPSN:equalToPSN:"), psn, psn2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/isValidConnectionID:forPSN:
func (o CPXProcessManagingObject) IsValidConnectionIDForPSN(id uint32, psn objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isValidConnectionID:forPSN:"), id, psn)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processForPID:
func (o CPXProcessManagingObject) ProcessForPID(pid int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processForPID:"), pid)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processForPSN:
func (o CPXProcessManagingObject) ProcessForPSN(psn objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processForPSN:"), psn)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processOwningConnection:
func (o CPXProcessManagingObject) ProcessOwningConnection(connection unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processOwningConnection:"), connection)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processOwningConnectionID:
func (o CPXProcessManagingObject) ProcessOwningConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processOwningConnectionID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processPendingKill
func (o CPXProcessManagingObject) ProcessPendingKill() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processPendingKill"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processRepresentedByConnection:
func (o CPXProcessManagingObject) ProcessRepresentedByConnection(connection unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnection:"), connection)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/processRepresentedByConnectionID:
func (o CPXProcessManagingObject) ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("processRepresentedByConnectionID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/setProcessPendingKill:
func (o CPXProcessManagingObject) SetProcessPendingKill(kill *CPSProcessRecRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setProcessPendingKill:"), kill)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManaging/updateProcessApplicationTypeIfNecessary:
func (o CPXProcessManagingObject) UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRecRef) byte {
	rv := objc.Send[byte](o.ID, objc.Sel("updateProcessApplicationTypeIfNecessary:"), necessary)
	return rv
}
