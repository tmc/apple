// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXProcessManager] class.
var (
	_CPXProcessManagerClass     CPXProcessManagerClass
	_CPXProcessManagerClassOnce sync.Once
)

func getCPXProcessManagerClass() CPXProcessManagerClass {
	_CPXProcessManagerClassOnce.Do(func() {
		_CPXProcessManagerClass = CPXProcessManagerClass{class: objc.GetClass("CPXProcessManager")}
	})
	return _CPXProcessManagerClass
}

// GetCPXProcessManagerClass returns the class object for CPXProcessManager.
func GetCPXProcessManagerClass() CPXProcessManagerClass {
	return getCPXProcessManagerClass()
}

type CPXProcessManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXProcessManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXProcessManagerClass) Alloc() CPXProcessManager {
	rv := objc.Send[CPXProcessManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXProcessManager.IsPSNEqualToPSN]
//   - [CPXProcessManager.IsValidConnectionIDForPSN]
//   - [CPXProcessManager.ProcessForPID]
//   - [CPXProcessManager.ProcessForPSN]
//   - [CPXProcessManager.ProcessOwningConnection]
//   - [CPXProcessManager.ProcessOwningConnectionID]
//   - [CPXProcessManager.ProcessPendingKill]
//   - [CPXProcessManager.SetProcessPendingKill]
//   - [CPXProcessManager.ProcessRepresentedByConnection]
//   - [CPXProcessManager.ProcessRepresentedByConnectionID]
//   - [CPXProcessManager.UpdateProcessApplicationTypeIfNecessary]
//   - [CPXProcessManager.InitWithSessionConnectionManager]
//   - [CPXProcessManager.DebugDescription]
//   - [CPXProcessManager.Description]
//   - [CPXProcessManager.Hash]
//   - [CPXProcessManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager
type CPXProcessManager struct {
	objectivec.Object
}

// CPXProcessManagerFromID constructs a [CPXProcessManager] from an objc.ID.
func CPXProcessManagerFromID(id objc.ID) CPXProcessManager {
	return CPXProcessManager{objectivec.Object{ID: id}}
}

// Ensure CPXProcessManager implements ICPXProcessManager.
var _ ICPXProcessManager = CPXProcessManager{}

// An interface definition for the [CPXProcessManager] class.
//
// # Methods
//
//   - [ICPXProcessManager.IsPSNEqualToPSN]
//   - [ICPXProcessManager.IsValidConnectionIDForPSN]
//   - [ICPXProcessManager.ProcessForPID]
//   - [ICPXProcessManager.ProcessForPSN]
//   - [ICPXProcessManager.ProcessOwningConnection]
//   - [ICPXProcessManager.ProcessOwningConnectionID]
//   - [ICPXProcessManager.ProcessPendingKill]
//   - [ICPXProcessManager.SetProcessPendingKill]
//   - [ICPXProcessManager.ProcessRepresentedByConnection]
//   - [ICPXProcessManager.ProcessRepresentedByConnectionID]
//   - [ICPXProcessManager.UpdateProcessApplicationTypeIfNecessary]
//   - [ICPXProcessManager.InitWithSessionConnectionManager]
//   - [ICPXProcessManager.DebugDescription]
//   - [ICPXProcessManager.Description]
//   - [ICPXProcessManager.Hash]
//   - [ICPXProcessManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager
type ICPXProcessManager interface {
	objectivec.IObject

	// Topic: Methods

	IsPSNEqualToPSN(psn objectivec.IObject, psn2 objectivec.IObject) bool
	IsValidConnectionIDForPSN(id uint32, psn objectivec.IObject) bool
	ProcessForPID(pid int) *CPSProcessRecRef
	ProcessForPSN(psn objectivec.IObject) *CPSProcessRecRef
	ProcessOwningConnection(connection unsafe.Pointer) *CPSProcessRecRef
	ProcessOwningConnectionID(id uint32) *CPSProcessRecRef
	ProcessPendingKill() *CPSProcessRecRef
	SetProcessPendingKill(value *CPSProcessRecRef)
	ProcessRepresentedByConnection(connection unsafe.Pointer) *CPSProcessRecRef
	ProcessRepresentedByConnectionID(id uint32) *CPSProcessRecRef
	UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRecRef) byte
	InitWithSessionConnectionManager(session unsafe.Pointer, manager objectivec.IObject) CPXProcessManager
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXProcessManager) Init() CPXProcessManager {
	rv := objc.Send[CPXProcessManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXProcessManager) Autorelease() CPXProcessManager {
	rv := objc.Send[CPXProcessManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXProcessManager creates a new CPXProcessManager instance.
func NewCPXProcessManager() CPXProcessManager {
	class := getCPXProcessManagerClass()
	rv := objc.Send[CPXProcessManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/initWithSession:connectionManager:
func NewCPXProcessManagerWithSessionConnectionManager(session unsafe.Pointer, manager objectivec.IObject) CPXProcessManager {
	instance := getCPXProcessManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:connectionManager:"), session, manager)
	return CPXProcessManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/isPSN:equalToPSN:
func (c CPXProcessManager) IsPSNEqualToPSN(psn objectivec.IObject, psn2 objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPSN:equalToPSN:"), psn, psn2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/isValidConnectionID:forPSN:
func (c CPXProcessManager) IsValidConnectionIDForPSN(id uint32, psn objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isValidConnectionID:forPSN:"), id, psn)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processForPID:
func (c CPXProcessManager) ProcessForPID(pid int) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processForPID:"), pid)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processForPSN:
func (c CPXProcessManager) ProcessForPSN(psn objectivec.IObject) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processForPSN:"), psn)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processOwningConnection:
func (c CPXProcessManager) ProcessOwningConnection(connection unsafe.Pointer) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processOwningConnection:"), connection)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processOwningConnectionID:
func (c CPXProcessManager) ProcessOwningConnectionID(id uint32) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processOwningConnectionID:"), id)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processRepresentedByConnection:
func (c CPXProcessManager) ProcessRepresentedByConnection(connection unsafe.Pointer) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processRepresentedByConnection:"), connection)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processRepresentedByConnectionID:
func (c CPXProcessManager) ProcessRepresentedByConnectionID(id uint32) *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processRepresentedByConnectionID:"), id)
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/updateProcessApplicationTypeIfNecessary:
func (c CPXProcessManager) UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRecRef) byte {
	rv := objc.Send[byte](c.ID, objc.Sel("updateProcessApplicationTypeIfNecessary:"), necessary)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/initWithSession:connectionManager:
func (c CPXProcessManager) InitWithSessionConnectionManager(session unsafe.Pointer, manager objectivec.IObject) CPXProcessManager {
	rv := objc.Send[CPXProcessManager](c.ID, objc.Sel("initWithSession:connectionManager:"), session, manager)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/debugDescription
func (c CPXProcessManager) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/description
func (c CPXProcessManager) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/hash
func (c CPXProcessManager) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/processPendingKill
func (c CPXProcessManager) ProcessPendingKill() *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processPendingKill"))
	return (*CPSProcessRecRef)(rv)
}
func (c CPXProcessManager) SetProcessPendingKill(value *CPSProcessRecRef) {
	objc.Send[struct{}](c.ID, objc.Sel("setProcessPendingKill:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXProcessManager/superclass
func (c CPXProcessManager) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
