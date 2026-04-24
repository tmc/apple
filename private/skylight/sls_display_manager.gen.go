// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayManager] class.
var (
	_SLSDisplayManagerClass     SLSDisplayManagerClass
	_SLSDisplayManagerClassOnce sync.Once
)

func getSLSDisplayManagerClass() SLSDisplayManagerClass {
	_SLSDisplayManagerClassOnce.Do(func() {
		_SLSDisplayManagerClass = SLSDisplayManagerClass{class: objc.GetClass("SLSDisplayManager")}
	})
	return _SLSDisplayManagerClass
}

// GetSLSDisplayManagerClass returns the class object for SLSDisplayManager.
func GetSLSDisplayManagerClass() SLSDisplayManagerClass {
	return getSLSDisplayManagerClass()
}

type SLSDisplayManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayManagerClass) Alloc() SLSDisplayManager {
	rv := objc.Send[SLSDisplayManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayManager.Client]
//   - [SLSDisplayManager.SetClient]
//   - [SLSDisplayManager.Connected]
//   - [SLSDisplayManager.SetConnected]
//   - [SLSDisplayManager.DeliverPowerStateNotificationPayload]
//   - [SLSDisplayManager.DispatchPowerStateNotificationPayload]
//   - [SLSDisplayManager.Enabled]
//   - [SLSDisplayManager.SetEnabled]
//   - [SLSDisplayManager.IdleDisplays]
//   - [SLSDisplayManager.LastPowerStateNotification]
//   - [SLSDisplayManager.SetLastPowerStateNotification]
//   - [SLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [SLSDisplayManager.RegisteredNotifiers]
//   - [SLSDisplayManager.SetRegisteredNotifiers]
//   - [SLSDisplayManager.RunningInServer]
//   - [SLSDisplayManager.SetRunningInServer]
//   - [SLSDisplayManager.Semaphore]
//   - [SLSDisplayManager.SetSemaphore]
//   - [SLSDisplayManager.UnregisterPowerStateNotificationRegistrationIDNotificationOption]
//   - [SLSDisplayManager.InitWith]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager
type SLSDisplayManager struct {
	objectivec.Object
}

// SLSDisplayManagerFromID constructs a [SLSDisplayManager] from an objc.ID.
func SLSDisplayManagerFromID(id objc.ID) SLSDisplayManager {
	return SLSDisplayManager{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayManager implements ISLSDisplayManager.
var _ ISLSDisplayManager = SLSDisplayManager{}

// An interface definition for the [SLSDisplayManager] class.
//
// # Methods
//
//   - [ISLSDisplayManager.Client]
//   - [ISLSDisplayManager.SetClient]
//   - [ISLSDisplayManager.Connected]
//   - [ISLSDisplayManager.SetConnected]
//   - [ISLSDisplayManager.DeliverPowerStateNotificationPayload]
//   - [ISLSDisplayManager.DispatchPowerStateNotificationPayload]
//   - [ISLSDisplayManager.Enabled]
//   - [ISLSDisplayManager.SetEnabled]
//   - [ISLSDisplayManager.IdleDisplays]
//   - [ISLSDisplayManager.LastPowerStateNotification]
//   - [ISLSDisplayManager.SetLastPowerStateNotification]
//   - [ISLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [ISLSDisplayManager.RegisteredNotifiers]
//   - [ISLSDisplayManager.SetRegisteredNotifiers]
//   - [ISLSDisplayManager.RunningInServer]
//   - [ISLSDisplayManager.SetRunningInServer]
//   - [ISLSDisplayManager.Semaphore]
//   - [ISLSDisplayManager.SetSemaphore]
//   - [ISLSDisplayManager.UnregisterPowerStateNotificationRegistrationIDNotificationOption]
//   - [ISLSDisplayManager.InitWith]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager
type ISLSDisplayManager interface {
	objectivec.IObject

	// Topic: Methods

	Client() ISLSGUIClient
	SetClient(value ISLSGUIClient)
	Connected() bool
	SetConnected(value bool)
	DeliverPowerStateNotificationPayload(notification byte, payload objectivec.IObject)
	DispatchPowerStateNotificationPayload(notification byte, payload objectivec.IObject)
	Enabled() bool
	SetEnabled(value bool)
	IdleDisplays(displays []objectivec.IObject) bool
	LastPowerStateNotification() byte
	SetLastPowerStateNotification(value byte)
	RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock(notification []objectivec.IObject, id objectivec.IObject, state bool, queue objectivec.IObject, refcon uintptr, option uint32, block VoidHandler, block2 VoidHandler) bool
	RegisteredNotifiers() unsafe.Pointer
	SetRegisteredNotifiers(value unsafe.Pointer)
	RunningInServer() bool
	SetRunningInServer(value bool)
	Semaphore() objectivec.Object
	SetSemaphore(value objectivec.Object)
	UnregisterPowerStateNotificationRegistrationIDNotificationOption(notification []objectivec.IObject, id objectivec.IObject, option uint32) bool
	InitWith(with []objectivec.IObject) SLSDisplayManager
}

// Init initializes the instance.
func (s SLSDisplayManager) Init() SLSDisplayManager {
	rv := objc.Send[SLSDisplayManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayManager) Autorelease() SLSDisplayManager {
	rv := objc.Send[SLSDisplayManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayManager creates a new SLSDisplayManager instance.
func NewSLSDisplayManager() SLSDisplayManager {
	class := getSLSDisplayManagerClass()
	rv := objc.Send[SLSDisplayManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/initWith:
func NewSLSDisplayManagerWith(with []objectivec.IObject) SLSDisplayManager {
	instance := getSLSDisplayManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:"), objectivec.IObjectSliceToNSArray(with))
	return SLSDisplayManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/deliverPowerStateNotification:payload:
func (s SLSDisplayManager) DeliverPowerStateNotificationPayload(notification byte, payload objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("deliverPowerStateNotification:payload:"), notification, payload)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/dispatchPowerStateNotification:payload:
func (s SLSDisplayManager) DispatchPowerStateNotificationPayload(notification byte, payload objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("dispatchPowerStateNotification:payload:"), notification, payload)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/idleDisplays:
func (s SLSDisplayManager) IdleDisplays(displays []objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("idleDisplays:"), objectivec.IObjectSliceToNSArray(displays))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/registerPowerStateNotification:registrationID:sendInitialState:queue:refcon:notificationOption:notificationBlock:notificationPayloadBlock:
func (s SLSDisplayManager) RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock(notification []objectivec.IObject, id objectivec.IObject, state bool, queue objectivec.IObject, refcon uintptr, option uint32, block VoidHandler, block2 VoidHandler) bool {
	_block6, _ := NewVoidBlock(block)
	_block7, _ := NewVoidBlock(block2)
	rv := objc.Send[bool](s.ID, objc.Sel("registerPowerStateNotification:registrationID:sendInitialState:queue:refcon:notificationOption:notificationBlock:notificationPayloadBlock:"), notification, id, state, queue, refcon, option, _block6, _block7)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/unregisterPowerStateNotification:registrationID:notificationOption:
func (s SLSDisplayManager) UnregisterPowerStateNotificationRegistrationIDNotificationOption(notification []objectivec.IObject, id objectivec.IObject, option uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("unregisterPowerStateNotification:registrationID:notificationOption:"), objectivec.IObjectSliceToNSArray(notification), id, option)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/initWith:
func (s SLSDisplayManager) InitWith(with []objectivec.IObject) SLSDisplayManager {
	rv := objc.Send[SLSDisplayManager](s.ID, objc.Sel("initWith:"), objectivec.IObjectSliceToNSArray(with))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/broadcastPowerStateChangeEvent:payload:
func (_SLSDisplayManagerClass SLSDisplayManagerClass) BroadcastPowerStateChangeEventPayload(event byte, payload objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SLSDisplayManagerClass.class), objc.Sel("broadcastPowerStateChangeEvent:payload:"), event, payload)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/client
func (s SLSDisplayManager) Client() ISLSGUIClient {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("client"))
	return SLSGUIClientFromID(objc.ID(rv))
}
func (s SLSDisplayManager) SetClient(value ISLSGUIClient) {
	objc.Send[struct{}](s.ID, objc.Sel("setClient:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/connected
func (s SLSDisplayManager) Connected() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("connected"))
	return rv
}
func (s SLSDisplayManager) SetConnected(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setConnected:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/enabled
func (s SLSDisplayManager) Enabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("enabled"))
	return rv
}
func (s SLSDisplayManager) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/lastPowerStateNotification
func (s SLSDisplayManager) LastPowerStateNotification() byte {
	rv := objc.Send[byte](s.ID, objc.Sel("lastPowerStateNotification"))
	return rv
}
func (s SLSDisplayManager) SetLastPowerStateNotification(value byte) {
	objc.Send[struct{}](s.ID, objc.Sel("setLastPowerStateNotification:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/registeredNotifiers
func (s SLSDisplayManager) RegisteredNotifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("registeredNotifiers"))
	return rv
}
func (s SLSDisplayManager) SetRegisteredNotifiers(value unsafe.Pointer) {
	objc.Send[struct{}](s.ID, objc.Sel("setRegisteredNotifiers:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/runningInServer
func (s SLSDisplayManager) RunningInServer() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("runningInServer"))
	return rv
}
func (s SLSDisplayManager) SetRunningInServer(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRunningInServer:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayManager/semaphore
func (s SLSDisplayManager) Semaphore() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("semaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLSDisplayManager) SetSemaphore(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setSemaphore:"), value)
}
