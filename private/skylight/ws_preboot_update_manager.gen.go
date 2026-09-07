// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSPrebootUpdateManager] class.
var (
	_WSPrebootUpdateManagerClass     WSPrebootUpdateManagerClass
	_WSPrebootUpdateManagerClassOnce sync.Once
)

func getWSPrebootUpdateManagerClass() WSPrebootUpdateManagerClass {
	_WSPrebootUpdateManagerClassOnce.Do(func() {
		_WSPrebootUpdateManagerClass = WSPrebootUpdateManagerClass{class: objc.GetClass("WSPrebootUpdateManager")}
	})
	return _WSPrebootUpdateManagerClass
}

// GetWSPrebootUpdateManagerClass returns the class object for WSPrebootUpdateManager.
func GetWSPrebootUpdateManagerClass() WSPrebootUpdateManagerClass {
	return getWSPrebootUpdateManagerClass()
}

type WSPrebootUpdateManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSPrebootUpdateManagerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSPrebootUpdateManagerClass) Alloc() WSPrebootUpdateManager {
	rv := objc.SendIfResponds[WSPrebootUpdateManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSPrebootUpdateManager.ApfsManager]
//   - [WSPrebootUpdateManager.SetApfsManager]
//   - [WSPrebootUpdateManager.BackgroundAdminPerformAsyncUpdate]
//   - [WSPrebootUpdateManager.BackgroundAdminRequestAsyncUpdate]
//   - [WSPrebootUpdateManager.BackgroundCancel]
//   - [WSPrebootUpdateManager.BackgroundHandleCancelRequest]
//   - [WSPrebootUpdateManager.BackgroundMonitorRequestedAsyncUpdate]
//   - [WSPrebootUpdateManager.BackgroundRemoveAndClearRunLoopSources]
//   - [WSPrebootUpdateManager.BackgroundScheduleRunLoopSources]
//   - [WSPrebootUpdateManager.BackgroundUpdateDidBegin]
//   - [WSPrebootUpdateManager.BackgroundUpdateDidCompleteWithErrorDetailErrorDictionary]
//   - [WSPrebootUpdateManager.BegunGroup]
//   - [WSPrebootUpdateManager.CancelPort]
//   - [WSPrebootUpdateManager.CancelRequested]
//   - [WSPrebootUpdateManager.SetCancelRequested]
//   - [WSPrebootUpdateManager.CompleteGroup]
//   - [WSPrebootUpdateManager.DiskManager]
//   - [WSPrebootUpdateManager.SetDiskManager]
//   - [WSPrebootUpdateManager.DmAsyncFinishedForDiskMainErrorDetailErrorDictionary]
//   - [WSPrebootUpdateManager.DmAsyncMessageForDiskStringDictionary]
//   - [WSPrebootUpdateManager.DmAsyncProgressForDiskBarberPolePercent]
//   - [WSPrebootUpdateManager.DmAsyncStartedForDisk]
//   - [WSPrebootUpdateManager.HandleMachMessage]
//   - [WSPrebootUpdateManager.MainAdminBeginAsyncUpdate]
//   - [WSPrebootUpdateManager.MainAdminPerformSyncUpdate]
//   - [WSPrebootUpdateManager.MainBeginAsyncUpdate]
//   - [WSPrebootUpdateManager.MainCancelInFlightUpdate]
//   - [WSPrebootUpdateManager.MainCancelPendingUpdateRequests]
//   - [WSPrebootUpdateManager.MainFinalizeUpdate]
//   - [WSPrebootUpdateManager.MainIsUpdateInFlight]
//   - [WSPrebootUpdateManager.MainIsUpdateRequestPending]
//   - [WSPrebootUpdateManager.MainPerformSyncUpdate]
//   - [WSPrebootUpdateManager.MainRequestAsyncUpdate]
//   - [WSPrebootUpdateManager.MainRescheduleForFinishedUpdate]
//   - [WSPrebootUpdateManager.MainSendCancel]
//   - [WSPrebootUpdateManager.MainUpdateDidCancel]
//   - [WSPrebootUpdateManager.MainUpdateDidCompleteWithErrorDetailErrorDictionary]
//   - [WSPrebootUpdateManager.RequestGroup]
//   - [WSPrebootUpdateManager.Stage]
//   - [WSPrebootUpdateManager.SetStage]
//   - [WSPrebootUpdateManager.DebugDescription]
//   - [WSPrebootUpdateManager.Description]
//   - [WSPrebootUpdateManager.Hash]
//   - [WSPrebootUpdateManager.Superclass]
type WSPrebootUpdateManager struct {
	objectivec.Object
}

// WSPrebootUpdateManagerFromID constructs a [WSPrebootUpdateManager] from an objc.ID.
func WSPrebootUpdateManagerFromID(id objc.ID) WSPrebootUpdateManager {
	return WSPrebootUpdateManager{objectivec.Object{ID: id}}
}

// Ensure WSPrebootUpdateManager implements IWSPrebootUpdateManager.
var _ IWSPrebootUpdateManager = WSPrebootUpdateManager{}

// An interface definition for the [WSPrebootUpdateManager] class.
//
// # Methods
//
//   - [IWSPrebootUpdateManager.ApfsManager]
//   - [IWSPrebootUpdateManager.SetApfsManager]
//   - [IWSPrebootUpdateManager.BackgroundAdminPerformAsyncUpdate]
//   - [IWSPrebootUpdateManager.BackgroundAdminRequestAsyncUpdate]
//   - [IWSPrebootUpdateManager.BackgroundCancel]
//   - [IWSPrebootUpdateManager.BackgroundHandleCancelRequest]
//   - [IWSPrebootUpdateManager.BackgroundMonitorRequestedAsyncUpdate]
//   - [IWSPrebootUpdateManager.BackgroundRemoveAndClearRunLoopSources]
//   - [IWSPrebootUpdateManager.BackgroundScheduleRunLoopSources]
//   - [IWSPrebootUpdateManager.BackgroundUpdateDidBegin]
//   - [IWSPrebootUpdateManager.BackgroundUpdateDidCompleteWithErrorDetailErrorDictionary]
//   - [IWSPrebootUpdateManager.BegunGroup]
//   - [IWSPrebootUpdateManager.CancelPort]
//   - [IWSPrebootUpdateManager.CancelRequested]
//   - [IWSPrebootUpdateManager.SetCancelRequested]
//   - [IWSPrebootUpdateManager.CompleteGroup]
//   - [IWSPrebootUpdateManager.DiskManager]
//   - [IWSPrebootUpdateManager.SetDiskManager]
//   - [IWSPrebootUpdateManager.DmAsyncFinishedForDiskMainErrorDetailErrorDictionary]
//   - [IWSPrebootUpdateManager.DmAsyncMessageForDiskStringDictionary]
//   - [IWSPrebootUpdateManager.DmAsyncProgressForDiskBarberPolePercent]
//   - [IWSPrebootUpdateManager.DmAsyncStartedForDisk]
//   - [IWSPrebootUpdateManager.HandleMachMessage]
//   - [IWSPrebootUpdateManager.MainAdminBeginAsyncUpdate]
//   - [IWSPrebootUpdateManager.MainAdminPerformSyncUpdate]
//   - [IWSPrebootUpdateManager.MainBeginAsyncUpdate]
//   - [IWSPrebootUpdateManager.MainCancelInFlightUpdate]
//   - [IWSPrebootUpdateManager.MainCancelPendingUpdateRequests]
//   - [IWSPrebootUpdateManager.MainFinalizeUpdate]
//   - [IWSPrebootUpdateManager.MainIsUpdateInFlight]
//   - [IWSPrebootUpdateManager.MainIsUpdateRequestPending]
//   - [IWSPrebootUpdateManager.MainPerformSyncUpdate]
//   - [IWSPrebootUpdateManager.MainRequestAsyncUpdate]
//   - [IWSPrebootUpdateManager.MainRescheduleForFinishedUpdate]
//   - [IWSPrebootUpdateManager.MainSendCancel]
//   - [IWSPrebootUpdateManager.MainUpdateDidCancel]
//   - [IWSPrebootUpdateManager.MainUpdateDidCompleteWithErrorDetailErrorDictionary]
//   - [IWSPrebootUpdateManager.RequestGroup]
//   - [IWSPrebootUpdateManager.Stage]
//   - [IWSPrebootUpdateManager.SetStage]
//   - [IWSPrebootUpdateManager.DebugDescription]
//   - [IWSPrebootUpdateManager.Description]
//   - [IWSPrebootUpdateManager.Hash]
//   - [IWSPrebootUpdateManager.Superclass]
type IWSPrebootUpdateManager interface {
	objectivec.IObject

	// Topic: Methods

	ApfsManager() unsafe.Pointer
	SetApfsManager(value unsafe.Pointer)
	BackgroundAdminPerformAsyncUpdate()
	BackgroundAdminRequestAsyncUpdate() bool
	BackgroundCancel()
	BackgroundHandleCancelRequest()
	BackgroundMonitorRequestedAsyncUpdate()
	BackgroundRemoveAndClearRunLoopSources()
	BackgroundScheduleRunLoopSources()
	BackgroundUpdateDidBegin()
	BackgroundUpdateDidCompleteWithErrorDetailErrorDictionary(withError int32, detailError int32, dictionary objectivec.IObject)
	BegunGroup() objectivec.Object
	CancelPort() foundation.NSMachPort
	CancelRequested() bool
	SetCancelRequested(value bool)
	CompleteGroup() objectivec.Object
	DiskManager() unsafe.Pointer
	SetDiskManager(value unsafe.Pointer)
	DmAsyncFinishedForDiskMainErrorDetailErrorDictionary(disk DADiskRef, mainError int32, detailError int32, dictionary objectivec.IObject)
	DmAsyncMessageForDiskStringDictionary(disk DADiskRef, string_ objectivec.IObject, dictionary objectivec.IObject)
	DmAsyncProgressForDiskBarberPolePercent(disk DADiskRef, pole bool, percent float32)
	DmAsyncStartedForDisk(disk DADiskRef)
	HandleMachMessage(message unsafe.Pointer)
	MainAdminBeginAsyncUpdate()
	MainAdminPerformSyncUpdate()
	MainBeginAsyncUpdate()
	MainCancelInFlightUpdate()
	MainCancelPendingUpdateRequests()
	MainFinalizeUpdate()
	MainIsUpdateInFlight() bool
	MainIsUpdateRequestPending() bool
	MainPerformSyncUpdate()
	MainRequestAsyncUpdate()
	MainRescheduleForFinishedUpdate()
	MainSendCancel()
	MainUpdateDidCancel()
	MainUpdateDidCompleteWithErrorDetailErrorDictionary(withError int32, detailError int32, dictionary objectivec.IObject)
	RequestGroup() objectivec.Object
	Stage() uint32
	SetStage(value uint32)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSPrebootUpdateManager) Init() WSPrebootUpdateManager {
	rv := objc.SendIfResponds[WSPrebootUpdateManager](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSPrebootUpdateManager) Autorelease() WSPrebootUpdateManager {
	rv := objc.SendIfResponds[WSPrebootUpdateManager](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSPrebootUpdateManager creates a new WSPrebootUpdateManager instance.
func NewWSPrebootUpdateManager() WSPrebootUpdateManager {
	class := getWSPrebootUpdateManagerClass()
	rv := objc.SendIfResponds[WSPrebootUpdateManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSPrebootUpdateManager) BackgroundAdminPerformAsyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundAdminPerformAsyncUpdate"))
}
func (w WSPrebootUpdateManager) BackgroundAdminRequestAsyncUpdate() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("backgroundAdminRequestAsyncUpdate"))
	return rv
}
func (w WSPrebootUpdateManager) BackgroundCancel() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundCancel"))
}
func (w WSPrebootUpdateManager) BackgroundHandleCancelRequest() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundHandleCancelRequest"))
}
func (w WSPrebootUpdateManager) BackgroundMonitorRequestedAsyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundMonitorRequestedAsyncUpdate"))
}
func (w WSPrebootUpdateManager) BackgroundRemoveAndClearRunLoopSources() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundRemoveAndClearRunLoopSources"))
}
func (w WSPrebootUpdateManager) BackgroundScheduleRunLoopSources() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundScheduleRunLoopSources"))
}
func (w WSPrebootUpdateManager) BackgroundUpdateDidBegin() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundUpdateDidBegin"))
}
func (w WSPrebootUpdateManager) BackgroundUpdateDidCompleteWithErrorDetailErrorDictionary(withError int32, detailError int32, dictionary objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("backgroundUpdateDidCompleteWithError:detailError:dictionary:"), withError, detailError, dictionary)
}
func (w WSPrebootUpdateManager) DmAsyncFinishedForDiskMainErrorDetailErrorDictionary(disk DADiskRef, mainError int32, detailError int32, dictionary objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("dmAsyncFinishedForDisk:mainError:detailError:dictionary:"), disk, mainError, detailError, dictionary)
}
func (w WSPrebootUpdateManager) DmAsyncMessageForDiskStringDictionary(disk DADiskRef, string_ objectivec.IObject, dictionary objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("dmAsyncMessageForDisk:string:dictionary:"), disk, string_, dictionary)
}
func (w WSPrebootUpdateManager) DmAsyncProgressForDiskBarberPolePercent(disk DADiskRef, pole bool, percent float32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("dmAsyncProgressForDisk:barberPole:percent:"), disk, pole, percent)
}
func (w WSPrebootUpdateManager) DmAsyncStartedForDisk(disk DADiskRef) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("dmAsyncStartedForDisk:"), disk)
}
func (w WSPrebootUpdateManager) HandleMachMessage(message unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleMachMessage:"), message)
}
func (w WSPrebootUpdateManager) MainAdminBeginAsyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainAdminBeginAsyncUpdate"))
}
func (w WSPrebootUpdateManager) MainAdminPerformSyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainAdminPerformSyncUpdate"))
}
func (w WSPrebootUpdateManager) MainBeginAsyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainBeginAsyncUpdate"))
}
func (w WSPrebootUpdateManager) MainCancelInFlightUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainCancelInFlightUpdate"))
}
func (w WSPrebootUpdateManager) MainCancelPendingUpdateRequests() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainCancelPendingUpdateRequests"))
}
func (w WSPrebootUpdateManager) MainFinalizeUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainFinalizeUpdate"))
}
func (w WSPrebootUpdateManager) MainIsUpdateInFlight() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("mainIsUpdateInFlight"))
	return rv
}
func (w WSPrebootUpdateManager) MainIsUpdateRequestPending() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("mainIsUpdateRequestPending"))
	return rv
}
func (w WSPrebootUpdateManager) MainPerformSyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainPerformSyncUpdate"))
}
func (w WSPrebootUpdateManager) MainRequestAsyncUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainRequestAsyncUpdate"))
}
func (w WSPrebootUpdateManager) MainRescheduleForFinishedUpdate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainRescheduleForFinishedUpdate"))
}
func (w WSPrebootUpdateManager) MainSendCancel() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainSendCancel"))
}
func (w WSPrebootUpdateManager) MainUpdateDidCancel() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainUpdateDidCancel"))
}
func (w WSPrebootUpdateManager) MainUpdateDidCompleteWithErrorDetailErrorDictionary(withError int32, detailError int32, dictionary objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("mainUpdateDidCompleteWithError:detailError:dictionary:"), withError, detailError, dictionary)
}

func (_WSPrebootUpdateManagerClass WSPrebootUpdateManagerClass) AdminGetDiskManagerAPFSManagerAndBootDisk(manager []objectivec.IObject, pFSManager []objectivec.IObject, disk *DADiskRef) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_WSPrebootUpdateManagerClass.class), objc.Sel("adminGetDiskManager:APFSManager:andBootDisk:"), objectivec.IObjectSliceToNSArray(manager), objectivec.IObjectSliceToNSArray(pFSManager), unsafe.Pointer(disk))
	return rv
}
func (_WSPrebootUpdateManagerClass WSPrebootUpdateManagerClass) SharedManager() WSPrebootUpdateManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSPrebootUpdateManagerClass.class), objc.Sel("sharedManager"))
	return WSPrebootUpdateManagerFromID(rv)
}
func (_WSPrebootUpdateManagerClass WSPrebootUpdateManagerClass) TreatBootDiskAsAPFSWithManager(disk objectivec.IObject, manager objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_WSPrebootUpdateManagerClass.class), objc.Sel("treatBootDisk:asAPFSWithManager:"), disk, manager)
	return rv
}
func (_WSPrebootUpdateManagerClass WSPrebootUpdateManagerClass) TreatBootDiskAsAPFS() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_WSPrebootUpdateManagerClass.class), objc.Sel("treatBootDiskAsAPFS"))
	return rv
}

func (w WSPrebootUpdateManager) ApfsManager() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("apfsManager"))
	return rv
}
func (w WSPrebootUpdateManager) SetApfsManager(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setApfsManager:"), value)
}
func (w WSPrebootUpdateManager) BegunGroup() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("begunGroup"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (w WSPrebootUpdateManager) CancelPort() foundation.NSMachPort {
	rv := objc.SendIfResponds[foundation.NSMachPort](w.ID, objc.Sel("cancelPort"))
	return foundation.NSMachPort(rv)
}
func (w WSPrebootUpdateManager) CancelRequested() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("cancelRequested"))
	return rv
}
func (w WSPrebootUpdateManager) SetCancelRequested(value bool) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setCancelRequested:"), value)
}
func (w WSPrebootUpdateManager) CompleteGroup() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("completeGroup"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (w WSPrebootUpdateManager) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSPrebootUpdateManager) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSPrebootUpdateManager) DiskManager() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("diskManager"))
	return rv
}
func (w WSPrebootUpdateManager) SetDiskManager(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setDiskManager:"), value)
}
func (w WSPrebootUpdateManager) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSPrebootUpdateManager) RequestGroup() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("requestGroup"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (w WSPrebootUpdateManager) Stage() uint32 {
	rv := objc.SendIfResponds[uint32](w.ID, objc.Sel("stage"))
	return rv
}
func (w WSPrebootUpdateManager) SetStage(value uint32) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setStage:"), value)
}
func (w WSPrebootUpdateManager) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
