// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXEventDeferringManager] class.
var (
	_CPXEventDeferringManagerClass     CPXEventDeferringManagerClass
	_CPXEventDeferringManagerClassOnce sync.Once
)

func getCPXEventDeferringManagerClass() CPXEventDeferringManagerClass {
	_CPXEventDeferringManagerClassOnce.Do(func() {
		_CPXEventDeferringManagerClass = CPXEventDeferringManagerClass{class: objc.GetClass("CPXEventDeferringManager")}
	})
	return _CPXEventDeferringManagerClass
}

// GetCPXEventDeferringManagerClass returns the class object for CPXEventDeferringManager.
func GetCPXEventDeferringManagerClass() CPXEventDeferringManagerClass {
	return getCPXEventDeferringManagerClass()
}

type CPXEventDeferringManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXEventDeferringManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXEventDeferringManagerClass) Alloc() CPXEventDeferringManager {
	rv := objc.Send[CPXEventDeferringManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXEventDeferringManager.EnforcedPolicy]
//   - [CPXEventDeferringManager.UpdatePolicyReason]
//   - [CPXEventDeferringManager.InitWithDeliveryManagerProcessManagerConnectionManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager
type CPXEventDeferringManager struct {
	objectivec.Object
}

// CPXEventDeferringManagerFromID constructs a [CPXEventDeferringManager] from an objc.ID.
func CPXEventDeferringManagerFromID(id objc.ID) CPXEventDeferringManager {
	return CPXEventDeferringManager{objectivec.Object{ID: id}}
}

// Ensure CPXEventDeferringManager implements ICPXEventDeferringManager.
var _ ICPXEventDeferringManager = CPXEventDeferringManager{}

// An interface definition for the [CPXEventDeferringManager] class.
//
// # Methods
//
//   - [ICPXEventDeferringManager.EnforcedPolicy]
//   - [ICPXEventDeferringManager.UpdatePolicyReason]
//   - [ICPXEventDeferringManager.InitWithDeliveryManagerProcessManagerConnectionManager]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager
type ICPXEventDeferringManager interface {
	objectivec.IObject

	// Topic: Methods

	EnforcedPolicy() ICPXEventDeferringPolicy
	UpdatePolicyReason(policy objectivec.IObject, reason objectivec.IObject)
	InitWithDeliveryManagerProcessManagerConnectionManager(manager objectivec.IObject, manager2 objectivec.IObject, manager3 objectivec.IObject) CPXEventDeferringManager
}

// Init initializes the instance.
func (c CPXEventDeferringManager) Init() CPXEventDeferringManager {
	rv := objc.Send[CPXEventDeferringManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXEventDeferringManager) Autorelease() CPXEventDeferringManager {
	rv := objc.Send[CPXEventDeferringManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXEventDeferringManager creates a new CPXEventDeferringManager instance.
func NewCPXEventDeferringManager() CPXEventDeferringManager {
	class := getCPXEventDeferringManagerClass()
	rv := objc.Send[CPXEventDeferringManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager/initWithDeliveryManager:processManager:connectionManager:
func NewCPXEventDeferringManagerWithDeliveryManagerProcessManagerConnectionManager(manager objectivec.IObject, manager2 objectivec.IObject, manager3 objectivec.IObject) CPXEventDeferringManager {
	instance := getCPXEventDeferringManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:processManager:connectionManager:"), manager, manager2, manager3)
	return CPXEventDeferringManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager/updatePolicy:reason:
func (c CPXEventDeferringManager) UpdatePolicyReason(policy objectivec.IObject, reason objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("updatePolicy:reason:"), policy, reason)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager/initWithDeliveryManager:processManager:connectionManager:
func (c CPXEventDeferringManager) InitWithDeliveryManagerProcessManagerConnectionManager(manager objectivec.IObject, manager2 objectivec.IObject, manager3 objectivec.IObject) CPXEventDeferringManager {
	rv := objc.Send[CPXEventDeferringManager](c.ID, objc.Sel("initWithDeliveryManager:processManager:connectionManager:"), manager, manager2, manager3)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager/_symbolicLinkTokenForProcess:
func (_CPXEventDeferringManagerClass CPXEventDeferringManagerClass) _symbolicLinkTokenForProcess(process *CPSProcessRecRef) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_CPXEventDeferringManagerClass.class), objc.Sel("_symbolicLinkTokenForProcess:"), process)
	return objectivec.Object{ID: rv}
}

// SymbolicLinkTokenForProcess is an exported wrapper for the private method _symbolicLinkTokenForProcess.
func (_CPXEventDeferringManagerClass CPXEventDeferringManagerClass) SymbolicLinkTokenForProcess(process *CPSProcessRecRef) objectivec.IObject {
	return _CPXEventDeferringManagerClass._symbolicLinkTokenForProcess(process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDeferringManager/enforcedPolicy
func (c CPXEventDeferringManager) EnforcedPolicy() ICPXEventDeferringPolicy {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("enforcedPolicy"))
	return CPXEventDeferringPolicyFromID(objc.ID(rv))
}
