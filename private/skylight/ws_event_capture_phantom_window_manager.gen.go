// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventCapturePhantomWindowManager] class.
var (
	_WSEventCapturePhantomWindowManagerClass     WSEventCapturePhantomWindowManagerClass
	_WSEventCapturePhantomWindowManagerClassOnce sync.Once
)

func getWSEventCapturePhantomWindowManagerClass() WSEventCapturePhantomWindowManagerClass {
	_WSEventCapturePhantomWindowManagerClassOnce.Do(func() {
		_WSEventCapturePhantomWindowManagerClass = WSEventCapturePhantomWindowManagerClass{class: objc.GetClass("WSEventCapturePhantomWindowManager")}
	})
	return _WSEventCapturePhantomWindowManagerClass
}

// GetWSEventCapturePhantomWindowManagerClass returns the class object for WSEventCapturePhantomWindowManager.
func GetWSEventCapturePhantomWindowManagerClass() WSEventCapturePhantomWindowManagerClass {
	return getWSEventCapturePhantomWindowManagerClass()
}

type WSEventCapturePhantomWindowManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventCapturePhantomWindowManagerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventCapturePhantomWindowManagerClass) Alloc() WSEventCapturePhantomWindowManager {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventCapturePhantomWindowManager._lock_acquirePhantomWindowLeaseForWindowIDReason]
//   - [WSEventCapturePhantomWindowManager._lock_retirePhantomWindowConnectionID]
//   - [WSEventCapturePhantomWindowManager.AcquirePhantomWindowLeaseForConnectionReason]
//   - [WSEventCapturePhantomWindowManager.AcquirePhantomWindowLeaseForWindowReason]
//   - [WSEventCapturePhantomWindowManager.InitWithProvider]
type WSEventCapturePhantomWindowManager struct {
	objectivec.Object
}

// WSEventCapturePhantomWindowManagerFromID constructs a [WSEventCapturePhantomWindowManager] from an objc.ID.
func WSEventCapturePhantomWindowManagerFromID(id objc.ID) WSEventCapturePhantomWindowManager {
	return WSEventCapturePhantomWindowManager{objectivec.Object{ID: id}}
}

// Ensure WSEventCapturePhantomWindowManager implements IWSEventCapturePhantomWindowManager.
var _ IWSEventCapturePhantomWindowManager = WSEventCapturePhantomWindowManager{}

// An interface definition for the [WSEventCapturePhantomWindowManager] class.
//
// # Methods
//
//   - [IWSEventCapturePhantomWindowManager._lock_acquirePhantomWindowLeaseForWindowIDReason]
//   - [IWSEventCapturePhantomWindowManager._lock_retirePhantomWindowConnectionID]
//   - [IWSEventCapturePhantomWindowManager.AcquirePhantomWindowLeaseForConnectionReason]
//   - [IWSEventCapturePhantomWindowManager.AcquirePhantomWindowLeaseForWindowReason]
//   - [IWSEventCapturePhantomWindowManager.InitWithProvider]
type IWSEventCapturePhantomWindowManager interface {
	objectivec.IObject

	// Topic: Methods

	_lock_acquirePhantomWindowLeaseForWindowIDReason(id uint32, reason objectivec.IObject) objectivec.IObject
	_lock_retirePhantomWindowConnectionID(window uint32, id uint32)
	AcquirePhantomWindowLeaseForConnectionReason(connection *CGXConnection, reason objectivec.IObject) objectivec.IObject
	AcquirePhantomWindowLeaseForWindowReason(window uint32, reason objectivec.IObject) objectivec.IObject
	InitWithProvider(provider objectivec.IObject) WSEventCapturePhantomWindowManager
}

// Init initializes the instance.
func (w WSEventCapturePhantomWindowManager) Init() WSEventCapturePhantomWindowManager {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowManager](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventCapturePhantomWindowManager) Autorelease() WSEventCapturePhantomWindowManager {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowManager](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventCapturePhantomWindowManager creates a new WSEventCapturePhantomWindowManager instance.
func NewWSEventCapturePhantomWindowManager() WSEventCapturePhantomWindowManager {
	class := getWSEventCapturePhantomWindowManagerClass()
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSEventCapturePhantomWindowManagerWithProvider(provider objectivec.IObject) WSEventCapturePhantomWindowManager {
	instance := getWSEventCapturePhantomWindowManagerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProvider:"), provider)
	return WSEventCapturePhantomWindowManagerFromID(rv)
}

func (w WSEventCapturePhantomWindowManager) _lock_acquirePhantomWindowLeaseForWindowIDReason(id uint32, reason objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_lock_acquirePhantomWindowLeaseForWindowID:reason:"), id, reason)
	return objectivec.Object{ID: rv}
}

// Lock_acquirePhantomWindowLeaseForWindowIDReason is an exported wrapper for the private method _lock_acquirePhantomWindowLeaseForWindowIDReason.
func (w WSEventCapturePhantomWindowManager) Lock_acquirePhantomWindowLeaseForWindowIDReason(id uint32, reason objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_lock_acquirePhantomWindowLeaseForWindowID:reason:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_lock_acquirePhantomWindowLeaseForWindowID:reason:"}
		return nil, err
	}
	return w._lock_acquirePhantomWindowLeaseForWindowIDReason(id, reason), nil
}

// CanLock_acquirePhantomWindowLeaseForWindowIDReason reports whether the receiver responds to the private selector _lock_acquirePhantomWindowLeaseForWindowID:reason:.
func (w WSEventCapturePhantomWindowManager) CanLock_acquirePhantomWindowLeaseForWindowIDReason() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_lock_acquirePhantomWindowLeaseForWindowID:reason:"))
}
func (w WSEventCapturePhantomWindowManager) _lock_retirePhantomWindowConnectionID(window uint32, id uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_lock_retirePhantomWindow:connectionID:"), window, id)
}

// Lock_retirePhantomWindowConnectionID is an exported wrapper for the private method _lock_retirePhantomWindowConnectionID.
func (w WSEventCapturePhantomWindowManager) Lock_retirePhantomWindowConnectionID(window uint32, id uint32) error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_lock_retirePhantomWindow:connectionID:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_lock_retirePhantomWindow:connectionID:"}
		return err
	}
	w._lock_retirePhantomWindowConnectionID(window, id)
	return nil
}

// CanLock_retirePhantomWindowConnectionID reports whether the receiver responds to the private selector _lock_retirePhantomWindow:connectionID:.
func (w WSEventCapturePhantomWindowManager) CanLock_retirePhantomWindowConnectionID() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_lock_retirePhantomWindow:connectionID:"))
}
func (w WSEventCapturePhantomWindowManager) AcquirePhantomWindowLeaseForConnectionReason(connection *CGXConnection, reason objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("acquirePhantomWindowLeaseForConnection:reason:"), unsafe.Pointer(connection), reason)
	return objectivec.Object{ID: rv}
}
func (w WSEventCapturePhantomWindowManager) AcquirePhantomWindowLeaseForWindowReason(window uint32, reason objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("acquirePhantomWindowLeaseForWindow:reason:"), window, reason)
	return objectivec.Object{ID: rv}
}
func (w WSEventCapturePhantomWindowManager) InitWithProvider(provider objectivec.IObject) WSEventCapturePhantomWindowManager {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowManager](w.ID, objc.Sel("initWithProvider:"), provider)
	return rv
}

func (_WSEventCapturePhantomWindowManagerClass WSEventCapturePhantomWindowManagerClass) SharedManager() WSEventCapturePhantomWindowManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSEventCapturePhantomWindowManagerClass.class), objc.Sel("sharedManager"))
	return WSEventCapturePhantomWindowManagerFromID(rv)
}
