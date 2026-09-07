// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplaySyncSessionController] class.
var (
	_SLSDisplaySyncSessionControllerClass     SLSDisplaySyncSessionControllerClass
	_SLSDisplaySyncSessionControllerClassOnce sync.Once
)

func getSLSDisplaySyncSessionControllerClass() SLSDisplaySyncSessionControllerClass {
	_SLSDisplaySyncSessionControllerClassOnce.Do(func() {
		_SLSDisplaySyncSessionControllerClass = SLSDisplaySyncSessionControllerClass{class: objc.GetClass("SLSDisplaySyncSessionController")}
	})
	return _SLSDisplaySyncSessionControllerClass
}

// GetSLSDisplaySyncSessionControllerClass returns the class object for SLSDisplaySyncSessionController.
func GetSLSDisplaySyncSessionControllerClass() SLSDisplaySyncSessionControllerClass {
	return getSLSDisplaySyncSessionControllerClass()
}

type SLSDisplaySyncSessionControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplaySyncSessionControllerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplaySyncSessionControllerClass) Alloc() SLSDisplaySyncSessionController {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplaySyncSessionController.AttemptConnect]
//   - [SLSDisplaySyncSessionController.DispatchToClientUuid]
//   - [SLSDisplaySyncSessionController.HandleNotificationUuid]
//   - [SLSDisplaySyncSessionController.InjectNotificationWithUUID]
//   - [SLSDisplaySyncSessionController.OnSessionChangedNotify]
//   - [SLSDisplaySyncSessionController.RegisterForNotificationsWithQueueBlock]
//   - [SLSDisplaySyncSessionController.RequestWithResponse]
//   - [SLSDisplaySyncSessionController.TearDownService]
//   - [SLSDisplaySyncSessionController.UnregisterNotification]
type SLSDisplaySyncSessionController struct {
	SLSDisplayControlClient
}

// SLSDisplaySyncSessionControllerFromID constructs a [SLSDisplaySyncSessionController] from an objc.ID.
func SLSDisplaySyncSessionControllerFromID(id objc.ID) SLSDisplaySyncSessionController {
	return SLSDisplaySyncSessionController{SLSDisplayControlClient: SLSDisplayControlClientFromID(id)}
}

// Ensure SLSDisplaySyncSessionController implements ISLSDisplaySyncSessionController.
var _ ISLSDisplaySyncSessionController = SLSDisplaySyncSessionController{}

// An interface definition for the [SLSDisplaySyncSessionController] class.
//
// # Methods
//
//   - [ISLSDisplaySyncSessionController.AttemptConnect]
//   - [ISLSDisplaySyncSessionController.DispatchToClientUuid]
//   - [ISLSDisplaySyncSessionController.HandleNotificationUuid]
//   - [ISLSDisplaySyncSessionController.InjectNotificationWithUUID]
//   - [ISLSDisplaySyncSessionController.OnSessionChangedNotify]
//   - [ISLSDisplaySyncSessionController.RegisterForNotificationsWithQueueBlock]
//   - [ISLSDisplaySyncSessionController.RequestWithResponse]
//   - [ISLSDisplaySyncSessionController.TearDownService]
//   - [ISLSDisplaySyncSessionController.UnregisterNotification]
type ISLSDisplaySyncSessionController interface {
	ISLSDisplayControlClient

	// Topic: Methods

	AttemptConnect()
	DispatchToClientUuid(client objectivec.IObject, uuid objectivec.IObject)
	HandleNotificationUuid(notification int32, uuid objectivec.IObject)
	InjectNotificationWithUUID(uuid objectivec.IObject)
	OnSessionChangedNotify()
	RegisterForNotificationsWithQueueBlock(queue objectivec.IObject, block VoidHandler)
	RequestWithResponse(request uint64, response []objectivec.IObject) bool
	TearDownService()
	UnregisterNotification()
}

// Init initializes the instance.
func (s SLSDisplaySyncSessionController) Init() SLSDisplaySyncSessionController {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplaySyncSessionController) Autorelease() SLSDisplaySyncSessionController {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplaySyncSessionController creates a new SLSDisplaySyncSessionController instance.
func NewSLSDisplaySyncSessionController() SLSDisplaySyncSessionController {
	class := getSLSDisplaySyncSessionControllerClass()
	rv := objc.SendIfResponds[SLSDisplaySyncSessionController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SLSDisplaySyncSessionController) AttemptConnect() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("attemptConnect"))
}
func (s SLSDisplaySyncSessionController) DispatchToClientUuid(client objectivec.IObject, uuid objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("dispatchToClient:uuid:"), client, uuid)
}
func (s SLSDisplaySyncSessionController) HandleNotificationUuid(notification int32, uuid objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("handleNotification:uuid:"), notification, uuid)
}
func (s SLSDisplaySyncSessionController) InjectNotificationWithUUID(uuid objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("injectNotificationWithUUID:"), uuid)
}
func (s SLSDisplaySyncSessionController) OnSessionChangedNotify() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("onSessionChangedNotify"))
}
func (s SLSDisplaySyncSessionController) RegisterForNotificationsWithQueueBlock(queue objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("registerForNotificationsWithQueue:block:"), queue, _block1)
}
func (s SLSDisplaySyncSessionController) RequestWithResponse(request uint64, response []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("request:withResponse:"), request, objectivec.IObjectSliceToNSArray(response))
	return rv
}
func (s SLSDisplaySyncSessionController) TearDownService() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("tearDownService"))
}
func (s SLSDisplaySyncSessionController) UnregisterNotification() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("unregisterNotification"))
}

// RegisterForNotificationsWithQueueBlockSync is a synchronous wrapper around [SLSDisplaySyncSessionController.RegisterForNotificationsWithQueueBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSDisplaySyncSessionController) RegisterForNotificationsWithQueueBlockSync(ctx context.Context, queue objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.RegisterForNotificationsWithQueueBlock(queue, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
