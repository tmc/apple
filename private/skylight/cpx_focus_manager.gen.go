// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXFocusManager] class.
var (
	_CPXFocusManagerClass     CPXFocusManagerClass
	_CPXFocusManagerClassOnce sync.Once
)

func getCPXFocusManagerClass() CPXFocusManagerClass {
	_CPXFocusManagerClassOnce.Do(func() {
		_CPXFocusManagerClass = CPXFocusManagerClass{class: objc.GetClass("CPXFocusManager")}
	})
	return _CPXFocusManagerClass
}

// GetCPXFocusManagerClass returns the class object for CPXFocusManager.
func GetCPXFocusManagerClass() CPXFocusManagerClass {
	return getCPXFocusManagerClass()
}

type CPXFocusManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXFocusManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXFocusManagerClass) Alloc() CPXFocusManager {
	rv := objc.Send[CPXFocusManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXFocusManager._fixBadForegroundProcess]
//   - [CPXFocusManager.AddToPermittedFrontList]
//   - [CPXFocusManager.CleanupForProcessDeath]
//   - [CPXFocusManager.FocusController]
//   - [CPXFocusManager.FrontVisibleProcess]
//   - [CPXFocusManager.FrontmostProcess]
//   - [CPXFocusManager.GetProcessToBringForwardAtNextCheckin]
//   - [CPXFocusManager.IsProcessPermittedToBeFrontmost]
//   - [CPXFocusManager.IsProcessToBringForwardAtNextCheckin]
//   - [CPXFocusManager.KeyThiefConnectionID]
//   - [CPXFocusManager.ProcessDidUpdateConnectionOldConnectionID]
//   - [CPXFocusManager.ReleaseAllKeyThiefInstancesNotPermittedFrontmost]
//   - [CPXFocusManager.RemoveFromPermittedFrontList]
//   - [CPXFocusManager.SetProcessToBringForwardAtNextCheckinPSN]
//   - [CPXFocusManager.SuppressDeferringPolicyUpdatesForReason]
//   - [CPXFocusManager.InitWithSessionDeferringManagerLaunchServicesProviderProcessManager]
//   - [CPXFocusManager.InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler]
//   - [CPXFocusManager.DebugDescription]
//   - [CPXFocusManager.Description]
//   - [CPXFocusManager.Hash]
//   - [CPXFocusManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager
type CPXFocusManager struct {
	objectivec.Object
}

// CPXFocusManagerFromID constructs a [CPXFocusManager] from an objc.ID.
func CPXFocusManagerFromID(id objc.ID) CPXFocusManager {
	return CPXFocusManager{objectivec.Object{ID: id}}
}

// Ensure CPXFocusManager implements ICPXFocusManager.
var _ ICPXFocusManager = CPXFocusManager{}

// An interface definition for the [CPXFocusManager] class.
//
// # Methods
//
//   - [ICPXFocusManager._fixBadForegroundProcess]
//   - [ICPXFocusManager.AddToPermittedFrontList]
//   - [ICPXFocusManager.CleanupForProcessDeath]
//   - [ICPXFocusManager.FocusController]
//   - [ICPXFocusManager.FrontVisibleProcess]
//   - [ICPXFocusManager.FrontmostProcess]
//   - [ICPXFocusManager.GetProcessToBringForwardAtNextCheckin]
//   - [ICPXFocusManager.IsProcessPermittedToBeFrontmost]
//   - [ICPXFocusManager.IsProcessToBringForwardAtNextCheckin]
//   - [ICPXFocusManager.KeyThiefConnectionID]
//   - [ICPXFocusManager.ProcessDidUpdateConnectionOldConnectionID]
//   - [ICPXFocusManager.ReleaseAllKeyThiefInstancesNotPermittedFrontmost]
//   - [ICPXFocusManager.RemoveFromPermittedFrontList]
//   - [ICPXFocusManager.SetProcessToBringForwardAtNextCheckinPSN]
//   - [ICPXFocusManager.SuppressDeferringPolicyUpdatesForReason]
//   - [ICPXFocusManager.InitWithSessionDeferringManagerLaunchServicesProviderProcessManager]
//   - [ICPXFocusManager.InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler]
//   - [ICPXFocusManager.DebugDescription]
//   - [ICPXFocusManager.Description]
//   - [ICPXFocusManager.Hash]
//   - [ICPXFocusManager.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager
type ICPXFocusManager interface {
	objectivec.IObject

	// Topic: Methods

	_fixBadForegroundProcess(process *CPSProcessRecRef)
	AddToPermittedFrontList(list objectivec.IObject) int16
	CleanupForProcessDeath(death *CPSProcessRecRef)
	FocusController() objectivec.IObject
	FrontVisibleProcess() *CPSProcessRecRef
	FrontmostProcess() *CPSProcessRecRef
	GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool
	IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool
	KeyThiefConnectionID() uint32
	ProcessDidUpdateConnectionOldConnectionID(connection *CPSProcessRecRef, id uint32)
	ReleaseAllKeyThiefInstancesNotPermittedFrontmost()
	RemoveFromPermittedFrontList(list objectivec.IObject) int16
	SetProcessToBringForwardAtNextCheckinPSN(psn objectivec.IObject) int16
	SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject
	InitWithSessionDeferringManagerLaunchServicesProviderProcessManager(session unsafe.Pointer, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager
	InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session unsafe.Pointer, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXFocusManager) Init() CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXFocusManager) Autorelease() CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXFocusManager creates a new CPXFocusManager instance.
func NewCPXFocusManager() CPXFocusManager {
	class := getCPXFocusManagerClass()
	rv := objc.Send[CPXFocusManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/initWithSession:deferringManager:launchServicesProvider:processManager:
func NewCPXFocusManagerWithSessionDeferringManagerLaunchServicesProviderProcessManager(session unsafe.Pointer, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager {
	instance := getCPXFocusManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:deferringManager:launchServicesProvider:processManager:"), session, manager, provider, manager2)
	return CPXFocusManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:
func NewCPXFocusManagerWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session unsafe.Pointer, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager {
	instance := getCPXFocusManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:"), session, policy, sanitizer, manager, source, manager2, scheduler)
	return CPXFocusManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/_fixBadForegroundProcess:
func (c CPXFocusManager) _fixBadForegroundProcess(process *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("_fixBadForegroundProcess:"), process)
}

// FixBadForegroundProcess is an exported wrapper for the private method _fixBadForegroundProcess.
func (c CPXFocusManager) FixBadForegroundProcess(process *CPSProcessRecRef) {
	c._fixBadForegroundProcess(process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/addToPermittedFrontList:
func (c CPXFocusManager) AddToPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/cleanupForProcessDeath:
func (c CPXFocusManager) CleanupForProcessDeath(death *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), death)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/getProcessToBringForwardAtNextCheckin:
func (c CPXFocusManager) GetProcessToBringForwardAtNextCheckin(checkin unsafe.Pointer) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/isProcessPermittedToBeFrontmost:
func (c CPXFocusManager) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRecRef) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/isProcessToBringForwardAtNextCheckin:
func (c CPXFocusManager) IsProcessToBringForwardAtNextCheckin(checkin objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/processDidUpdateConnection:oldConnectionID:
func (c CPXFocusManager) ProcessDidUpdateConnectionOldConnectionID(connection *CPSProcessRecRef, id uint32) {
	objc.Send[objc.ID](c.ID, objc.Sel("processDidUpdateConnection:oldConnectionID:"), connection, id)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/releaseAllKeyThiefInstancesNotPermittedFrontmost
func (c CPXFocusManager) ReleaseAllKeyThiefInstancesNotPermittedFrontmost() {
	objc.Send[objc.ID](c.ID, objc.Sel("releaseAllKeyThiefInstancesNotPermittedFrontmost"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/removeFromPermittedFrontList:
func (c CPXFocusManager) RemoveFromPermittedFrontList(list objectivec.IObject) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/setProcessToBringForwardAtNextCheckinPSN:
func (c CPXFocusManager) SetProcessToBringForwardAtNextCheckinPSN(psn objectivec.IObject) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("setProcessToBringForwardAtNextCheckinPSN:"), psn)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/suppressDeferringPolicyUpdatesForReason:
func (c CPXFocusManager) SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("suppressDeferringPolicyUpdatesForReason:"), reason)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/initWithSession:deferringManager:launchServicesProvider:processManager:
func (c CPXFocusManager) InitWithSessionDeferringManagerLaunchServicesProviderProcessManager(session unsafe.Pointer, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("initWithSession:deferringManager:launchServicesProvider:processManager:"), session, manager, provider, manager2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:
func (c CPXFocusManager) InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session unsafe.Pointer, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:"), session, policy, sanitizer, manager, source, manager2, scheduler)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/debugDescription
func (c CPXFocusManager) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/description
func (c CPXFocusManager) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/focusController
func (c CPXFocusManager) FocusController() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("focusController"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/frontVisibleProcess
func (c CPXFocusManager) FrontVisibleProcess() *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontVisibleProcess"))
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/frontmostProcess
func (c CPXFocusManager) FrontmostProcess() *CPSProcessRecRef {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontmostProcess"))
	return (*CPSProcessRecRef)(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/hash
func (c CPXFocusManager) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/keyThiefConnectionID
func (c CPXFocusManager) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXFocusManager/superclass
func (c CPXFocusManager) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
