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
type ICPXFocusManager interface {
	objectivec.IObject

	// Topic: Methods

	_fixBadForegroundProcess(process *CPSProcessRec)
	AddToPermittedFrontList(list CPSProcessSerNum) int16
	CleanupForProcessDeath(death *CPSProcessRec)
	FocusController() unsafe.Pointer
	FrontVisibleProcess() unsafe.Pointer
	FrontmostProcess() unsafe.Pointer
	GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool
	IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool
	IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool
	KeyThiefConnectionID() uint32
	ProcessDidUpdateConnectionOldConnectionID(connection *CPSProcessRec, id uint32)
	ReleaseAllKeyThiefInstancesNotPermittedFrontmost()
	RemoveFromPermittedFrontList(list CPSProcessSerNum) int16
	SetProcessToBringForwardAtNextCheckinPSN(psn CPSProcessSerNum) int16
	SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject
	InitWithSessionDeferringManagerLaunchServicesProviderProcessManager(session *CGXSession, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager
	InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session *CGXSession, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func NewCPXFocusManagerWithSessionDeferringManagerLaunchServicesProviderProcessManager(session *CGXSession, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager {
	instance := getCPXFocusManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:deferringManager:launchServicesProvider:processManager:"), session, manager, provider, manager2)
	return CPXFocusManagerFromID(rv)
}

func NewCPXFocusManagerWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session *CGXSession, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager {
	instance := getCPXFocusManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:"), session, policy, sanitizer, manager, source, manager2, scheduler)
	return CPXFocusManagerFromID(rv)
}

func (c CPXFocusManager) _fixBadForegroundProcess(process *CPSProcessRec) {
	objc.Send[objc.ID](c.ID, objc.Sel("_fixBadForegroundProcess:"), process)
}

// FixBadForegroundProcess is an exported wrapper for the private method _fixBadForegroundProcess.
func (c CPXFocusManager) FixBadForegroundProcess(process *CPSProcessRec) error {
	if !objc.RespondsToSelector(c.ID, objc.Sel("_fixBadForegroundProcess:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_fixBadForegroundProcess:"}
		return err
	}
	c._fixBadForegroundProcess(process)
	return nil
}

// CanFixBadForegroundProcess reports whether the receiver responds to the private selector _fixBadForegroundProcess:.
func (c CPXFocusManager) CanFixBadForegroundProcess() bool {
	return objc.RespondsToSelector(c.ID, objc.Sel("_fixBadForegroundProcess:"))
}
func (c CPXFocusManager) AddToPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("addToPermittedFrontList:"), list)
	return rv
}
func (c CPXFocusManager) CleanupForProcessDeath(death *CPSProcessRec) {
	objc.Send[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), death)
}
func (c CPXFocusManager) GetProcessToBringForwardAtNextCheckin(checkin *CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("getProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (c CPXFocusManager) IsProcessPermittedToBeFrontmost(frontmost *CPSProcessRec) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessPermittedToBeFrontmost:"), frontmost)
	return rv
}
func (c CPXFocusManager) IsProcessToBringForwardAtNextCheckin(checkin CPSProcessSerNum) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isProcessToBringForwardAtNextCheckin:"), checkin)
	return rv
}
func (c CPXFocusManager) ProcessDidUpdateConnectionOldConnectionID(connection *CPSProcessRec, id uint32) {
	objc.Send[objc.ID](c.ID, objc.Sel("processDidUpdateConnection:oldConnectionID:"), connection, id)
}
func (c CPXFocusManager) ReleaseAllKeyThiefInstancesNotPermittedFrontmost() {
	objc.Send[objc.ID](c.ID, objc.Sel("releaseAllKeyThiefInstancesNotPermittedFrontmost"))
}
func (c CPXFocusManager) RemoveFromPermittedFrontList(list CPSProcessSerNum) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("removeFromPermittedFrontList:"), list)
	return rv
}
func (c CPXFocusManager) SetProcessToBringForwardAtNextCheckinPSN(psn CPSProcessSerNum) int16 {
	rv := objc.Send[int16](c.ID, objc.Sel("setProcessToBringForwardAtNextCheckinPSN:"), psn)
	return rv
}
func (c CPXFocusManager) SuppressDeferringPolicyUpdatesForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("suppressDeferringPolicyUpdatesForReason:"), reason)
	return objectivec.Object{ID: rv}
}
func (c CPXFocusManager) InitWithSessionDeferringManagerLaunchServicesProviderProcessManager(session *CGXSession, manager objectivec.IObject, provider objectivec.IObject, manager2 objectivec.IObject) CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("initWithSession:deferringManager:launchServicesProvider:processManager:"), session, manager, provider, manager2)
	return rv
}
func (c CPXFocusManager) InitWithSessionPolicyDeferringPolicySanitizerDeferringManagerDataSourceProcessManagerCallbackScheduler(session *CGXSession, policy objectivec.IObject, sanitizer objectivec.IObject, manager objectivec.IObject, source objectivec.IObject, manager2 objectivec.IObject, scheduler objectivec.IObject) CPXFocusManager {
	rv := objc.Send[CPXFocusManager](c.ID, objc.Sel("initWithSession:policy:deferringPolicySanitizer:deferringManager:dataSource:processManager:callbackScheduler:"), session, policy, sanitizer, manager, source, manager2, scheduler)
	return rv
}

func (c CPXFocusManager) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXFocusManager) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXFocusManager) FocusController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("focusController"))
	return rv
}
func (c CPXFocusManager) FrontVisibleProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontVisibleProcess"))
	return rv
}
func (c CPXFocusManager) FrontmostProcess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("frontmostProcess"))
	return rv
}
func (c CPXFocusManager) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXFocusManager) KeyThiefConnectionID() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("keyThiefConnectionID"))
	return rv
}
func (c CPXFocusManager) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
