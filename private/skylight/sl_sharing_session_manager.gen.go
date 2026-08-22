// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSharingSessionManager] class.
var (
	_SLSharingSessionManagerClass     SLSharingSessionManagerClass
	_SLSharingSessionManagerClassOnce sync.Once
)

func getSLSharingSessionManagerClass() SLSharingSessionManagerClass {
	_SLSharingSessionManagerClassOnce.Do(func() {
		_SLSharingSessionManagerClass = SLSharingSessionManagerClass{class: objc.GetClass("SLSharingSessionManager")}
	})
	return _SLSharingSessionManagerClass
}

// GetSLSharingSessionManagerClass returns the class object for SLSharingSessionManager.
func GetSLSharingSessionManagerClass() SLSharingSessionManagerClass {
	return getSLSharingSessionManagerClass()
}

type SLSharingSessionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSharingSessionManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSharingSessionManagerClass) Alloc() SLSharingSessionManager {
	rv := objc.SendIfResponds[SLSharingSessionManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSharingSessionManager.BeginNoPillSharingSessionWithTitle]
//   - [SLSharingSessionManager.BeginSharingSessionWithTitle]
//   - [SLSharingSessionManager.BeginSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications]
//   - [SLSharingSessionManager.ClientContexts]
//   - [SLSharingSessionManager.SetClientContexts]
//   - [SLSharingSessionManager.ContextForPayload]
//   - [SLSharingSessionManager.CopyAllSessions]
//   - [SLSharingSessionManager.Delegate]
//   - [SLSharingSessionManager.SetDelegate]
//   - [SLSharingSessionManager.DispatchToClientDelegate]
//   - [SLSharingSessionManager.EndSharingSession]
//   - [SLSharingSessionManager.NotificationDictionary]
//   - [SLSharingSessionManager.PickerCanceledSession]
//   - [SLSharingSessionManager.RegisterNotification]
//   - [SLSharingSessionManager.SetDelegateBlock]
//   - [SLSharingSessionManager.SystemDelegate]
//   - [SLSharingSessionManager.SetSystemDelegate]
//   - [SLSharingSessionManager.UnregisterNotification]
//   - [SLSharingSessionManager.InitPrivate]
type SLSharingSessionManager struct {
	objectivec.Object
}

// SLSharingSessionManagerFromID constructs a [SLSharingSessionManager] from an objc.ID.
func SLSharingSessionManagerFromID(id objc.ID) SLSharingSessionManager {
	return SLSharingSessionManager{objectivec.Object{ID: id}}
}

// Ensure SLSharingSessionManager implements ISLSharingSessionManager.
var _ ISLSharingSessionManager = SLSharingSessionManager{}

// An interface definition for the [SLSharingSessionManager] class.
//
// # Methods
//
//   - [ISLSharingSessionManager.BeginNoPillSharingSessionWithTitle]
//   - [ISLSharingSessionManager.BeginSharingSessionWithTitle]
//   - [ISLSharingSessionManager.BeginSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications]
//   - [ISLSharingSessionManager.ClientContexts]
//   - [ISLSharingSessionManager.SetClientContexts]
//   - [ISLSharingSessionManager.ContextForPayload]
//   - [ISLSharingSessionManager.CopyAllSessions]
//   - [ISLSharingSessionManager.Delegate]
//   - [ISLSharingSessionManager.SetDelegate]
//   - [ISLSharingSessionManager.DispatchToClientDelegate]
//   - [ISLSharingSessionManager.EndSharingSession]
//   - [ISLSharingSessionManager.NotificationDictionary]
//   - [ISLSharingSessionManager.PickerCanceledSession]
//   - [ISLSharingSessionManager.RegisterNotification]
//   - [ISLSharingSessionManager.SetDelegateBlock]
//   - [ISLSharingSessionManager.SystemDelegate]
//   - [ISLSharingSessionManager.SetSystemDelegate]
//   - [ISLSharingSessionManager.UnregisterNotification]
//   - [ISLSharingSessionManager.InitPrivate]
type ISLSharingSessionManager interface {
	objectivec.IObject

	// Topic: Methods

	BeginNoPillSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject
	BeginSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject
	BeginSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) objectivec.IObject
	ClientContexts() foundation.NSMapTable
	SetClientContexts(value foundation.NSMapTable)
	ContextForPayload(payload objectivec.IObject) objectivec.IObject
	CopyAllSessions() objectivec.IObject
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DispatchToClientDelegate(delegate objectivec.IObject)
	EndSharingSession(session objectivec.IObject)
	NotificationDictionary() objectivec.IObject
	PickerCanceledSession(session objectivec.IObject)
	RegisterNotification()
	SetDelegateBlock(block VoidHandler)
	SystemDelegate() unsafe.Pointer
	SetSystemDelegate(value unsafe.Pointer)
	UnregisterNotification()
	InitPrivate() SLSharingSessionManager
}

// Init initializes the instance.
func (s SLSharingSessionManager) Init() SLSharingSessionManager {
	rv := objc.SendIfResponds[SLSharingSessionManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSharingSessionManager) Autorelease() SLSharingSessionManager {
	rv := objc.SendIfResponds[SLSharingSessionManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSharingSessionManager creates a new SLSharingSessionManager instance.
func NewSLSharingSessionManager() SLSharingSessionManager {
	class := getSLSharingSessionManagerClass()
	rv := objc.SendIfResponds[SLSharingSessionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSharingSessionManagerPrivate() SLSharingSessionManager {
	instance := getSLSharingSessionManagerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initPrivate"))
	return SLSharingSessionManagerFromID(rv)
}

func (s SLSharingSessionManager) BeginNoPillSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("beginNoPillSharingSessionWithTitle:"), title)
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) BeginSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("beginSharingSessionWithTitle:"), title)
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) BeginSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("beginSharingSessionWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) ContextForPayload(payload objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("contextForPayload:"), payload)
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) CopyAllSessions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyAllSessions"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) DispatchToClientDelegate(delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("dispatchToClientDelegate:"), delegate)
}
func (s SLSharingSessionManager) EndSharingSession(session objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("endSharingSession:"), session)
}
func (s SLSharingSessionManager) NotificationDictionary() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("notificationDictionary"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) PickerCanceledSession(session objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("pickerCanceledSession:"), session)
}
func (s SLSharingSessionManager) RegisterNotification() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("registerNotification"))
}

var _slsharingsessionmanager_setdelegateblock_p0_key byte

func (s SLSharingSessionManager) SetDelegateBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setDelegateBlock:"), _block0)
}
func (s SLSharingSessionManager) UnregisterNotification() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("unregisterNotification"))
}
func (s SLSharingSessionManager) InitPrivate() SLSharingSessionManager {
	rv := objc.SendIfResponds[SLSharingSessionManager](s.ID, objc.Sel("initPrivate"))
	return rv
}

func (_SLSharingSessionManagerClass SLSharingSessionManagerClass) Shared() SLSharingSessionManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLSharingSessionManagerClass.class), objc.Sel("shared"))
	return SLSharingSessionManagerFromID(rv)
}

func (s SLSharingSessionManager) ClientContexts() foundation.NSMapTable {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("clientContexts"))
	return foundation.NSMapTableFromID(objc.ID(rv))
}
func (s SLSharingSessionManager) SetClientContexts(value foundation.NSMapTable) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setClientContexts:"), value)
}
func (s SLSharingSessionManager) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("delegate"))
	return rv
}
func (s SLSharingSessionManager) SetDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
func (s SLSharingSessionManager) SystemDelegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("systemDelegate"))
	return rv
}
func (s SLSharingSessionManager) SetSystemDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setSystemDelegate:"), value)
}

// SetDelegateBlockSync is a synchronous wrapper around [SLSharingSessionManager.SetDelegateBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSharingSessionManager) SetDelegateBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetDelegateBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
