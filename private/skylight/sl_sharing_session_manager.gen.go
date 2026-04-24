// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

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
	rv := objc.Send[SLSharingSessionManager](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager
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
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DispatchToClientDelegate(delegate objectivec.IObject)
	EndSharingSession(session objectivec.IObject)
	NotificationDictionary() objectivec.IObject
	PickerCanceledSession(session objectivec.IObject)
	RegisterNotification()
	SetDelegateBlock(block VoidHandler)
	SystemDelegate() objectivec.IObject
	SetSystemDelegate(value objectivec.IObject)
	UnregisterNotification()
	InitPrivate() SLSharingSessionManager
}

// Init initializes the instance.
func (s SLSharingSessionManager) Init() SLSharingSessionManager {
	rv := objc.Send[SLSharingSessionManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSharingSessionManager) Autorelease() SLSharingSessionManager {
	rv := objc.Send[SLSharingSessionManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSharingSessionManager creates a new SLSharingSessionManager instance.
func NewSLSharingSessionManager() SLSharingSessionManager {
	class := getSLSharingSessionManagerClass()
	rv := objc.Send[SLSharingSessionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/initPrivate
func NewSLSharingSessionManagerPrivate() SLSharingSessionManager {
	instance := getSLSharingSessionManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initPrivate"))
	return SLSharingSessionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/beginNoPillSharingSessionWithTitle:
func (s SLSharingSessionManager) BeginNoPillSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("beginNoPillSharingSessionWithTitle:"), title)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/beginSharingSessionWithTitle:
func (s SLSharingSessionManager) BeginSharingSessionWithTitle(title objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("beginSharingSessionWithTitle:"), title)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/beginSharingSessionWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:
func (s SLSharingSessionManager) BeginSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("beginSharingSessionWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/contextForPayload:
func (s SLSharingSessionManager) ContextForPayload(payload objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("contextForPayload:"), payload)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/copyAllSessions
func (s SLSharingSessionManager) CopyAllSessions() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyAllSessions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/dispatchToClientDelegate:
func (s SLSharingSessionManager) DispatchToClientDelegate(delegate objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("dispatchToClientDelegate:"), delegate)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/endSharingSession:
func (s SLSharingSessionManager) EndSharingSession(session objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("endSharingSession:"), session)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/notificationDictionary
func (s SLSharingSessionManager) NotificationDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("notificationDictionary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/pickerCanceledSession:
func (s SLSharingSessionManager) PickerCanceledSession(session objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("pickerCanceledSession:"), session)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/registerNotification
func (s SLSharingSessionManager) RegisterNotification() {
	objc.Send[objc.ID](s.ID, objc.Sel("registerNotification"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/setDelegateBlock:
func (s SLSharingSessionManager) SetDelegateBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setDelegateBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/unregisterNotification
func (s SLSharingSessionManager) UnregisterNotification() {
	objc.Send[objc.ID](s.ID, objc.Sel("unregisterNotification"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/initPrivate
func (s SLSharingSessionManager) InitPrivate() SLSharingSessionManager {
	rv := objc.Send[SLSharingSessionManager](s.ID, objc.Sel("initPrivate"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/shared
func (_SLSharingSessionManagerClass SLSharingSessionManagerClass) Shared() SLSharingSessionManager {
	rv := objc.Send[objc.ID](objc.ID(_SLSharingSessionManagerClass.class), objc.Sel("shared"))
	return SLSharingSessionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/clientContexts
func (s SLSharingSessionManager) ClientContexts() foundation.NSMapTable {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("clientContexts"))
	return foundation.NSMapTableFromID(objc.ID(rv))
}
func (s SLSharingSessionManager) SetClientContexts(value foundation.NSMapTable) {
	objc.Send[struct{}](s.ID, objc.Sel("setClientContexts:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/delegate
func (s SLSharingSessionManager) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionManager/systemDelegate
func (s SLSharingSessionManager) SystemDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("systemDelegate"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSessionManager) SetSystemDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setSystemDelegate:"), value)
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
