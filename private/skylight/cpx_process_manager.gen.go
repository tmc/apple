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
type ICPXProcessManager interface {
	objectivec.IObject

	// Topic: Methods

	IsPSNEqualToPSN(psn CPSProcessSerNum, psn2 CPSProcessSerNum) bool
	IsValidConnectionIDForPSN(id uint32, psn CPSProcessSerNum) bool
	ProcessForPID(pid int) unsafe.Pointer
	ProcessForPSN(psn CPSProcessSerNum) unsafe.Pointer
	ProcessOwningConnection(connection *CGXConnection) unsafe.Pointer
	ProcessOwningConnectionID(id uint32) unsafe.Pointer
	ProcessPendingKill() unsafe.Pointer
	SetProcessPendingKill(value *CPSProcessRec)
	ProcessRepresentedByConnection(connection *CGXConnection) unsafe.Pointer
	ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer
	UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRec) byte
	InitWithSessionConnectionManager(session *CGXSession, manager objectivec.IObject) CPXProcessManager
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func NewCPXProcessManagerWithSessionConnectionManager(session *CGXSession, manager objectivec.IObject) CPXProcessManager {
	instance := getCPXProcessManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:connectionManager:"), session, manager)
	return CPXProcessManagerFromID(rv)
}

func (c CPXProcessManager) IsPSNEqualToPSN(psn CPSProcessSerNum, psn2 CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPSN:equalToPSN:"), psn, psn2)
	return rv
}
func (c CPXProcessManager) IsValidConnectionIDForPSN(id uint32, psn CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isValidConnectionID:forPSN:"), id, psn)
	return rv
}
func (c CPXProcessManager) ProcessForPID(pid int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processForPID:"), pid)
	return rv
}
func (c CPXProcessManager) ProcessForPSN(psn CPSProcessSerNum) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processForPSN:"), psn)
	return rv
}
func (c CPXProcessManager) ProcessOwningConnection(connection *CGXConnection) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processOwningConnection:"), connection)
	return rv
}
func (c CPXProcessManager) ProcessOwningConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processOwningConnectionID:"), id)
	return rv
}
func (c CPXProcessManager) ProcessRepresentedByConnection(connection *CGXConnection) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processRepresentedByConnection:"), connection)
	return rv
}
func (c CPXProcessManager) ProcessRepresentedByConnectionID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processRepresentedByConnectionID:"), id)
	return rv
}
func (c CPXProcessManager) UpdateProcessApplicationTypeIfNecessary(necessary *CPSProcessRec) byte {
	rv := objc.Send[byte](c.ID, objc.Sel("updateProcessApplicationTypeIfNecessary:"), necessary)
	return rv
}
func (c CPXProcessManager) InitWithSessionConnectionManager(session *CGXSession, manager objectivec.IObject) CPXProcessManager {
	rv := objc.Send[CPXProcessManager](c.ID, objc.Sel("initWithSession:connectionManager:"), session, manager)
	return rv
}

func (c CPXProcessManager) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXProcessManager) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXProcessManager) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXProcessManager) ProcessPendingKill() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("processPendingKill"))
	return rv
}
func (c CPXProcessManager) SetProcessPendingKill(value *CPSProcessRec) {
	objc.Send[struct{}](c.ID, objc.Sel("setProcessPendingKill:"), value)
}
func (c CPXProcessManager) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
