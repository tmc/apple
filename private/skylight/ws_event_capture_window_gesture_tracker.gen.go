// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventCaptureWindowGestureTracker] class.
var (
	_WSEventCaptureWindowGestureTrackerClass     WSEventCaptureWindowGestureTrackerClass
	_WSEventCaptureWindowGestureTrackerClassOnce sync.Once
)

func getWSEventCaptureWindowGestureTrackerClass() WSEventCaptureWindowGestureTrackerClass {
	_WSEventCaptureWindowGestureTrackerClassOnce.Do(func() {
		_WSEventCaptureWindowGestureTrackerClass = WSEventCaptureWindowGestureTrackerClass{class: objc.GetClass("WSEventCaptureWindowGestureTracker")}
	})
	return _WSEventCaptureWindowGestureTrackerClass
}

// GetWSEventCaptureWindowGestureTrackerClass returns the class object for WSEventCaptureWindowGestureTracker.
func GetWSEventCaptureWindowGestureTrackerClass() WSEventCaptureWindowGestureTrackerClass {
	return getWSEventCaptureWindowGestureTrackerClass()
}

type WSEventCaptureWindowGestureTrackerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventCaptureWindowGestureTrackerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventCaptureWindowGestureTrackerClass) Alloc() WSEventCaptureWindowGestureTracker {
	rv := objc.SendIfResponds[WSEventCaptureWindowGestureTracker](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

type WSEventCaptureWindowGestureTracker struct {
	objectivec.Object
}

// WSEventCaptureWindowGestureTrackerFromID constructs a [WSEventCaptureWindowGestureTracker] from an objc.ID.
func WSEventCaptureWindowGestureTrackerFromID(id objc.ID) WSEventCaptureWindowGestureTracker {
	return WSEventCaptureWindowGestureTracker{objectivec.Object{ID: id}}
}

// Ensure WSEventCaptureWindowGestureTracker implements IWSEventCaptureWindowGestureTracker.
var _ IWSEventCaptureWindowGestureTracker = WSEventCaptureWindowGestureTracker{}

// An interface definition for the [WSEventCaptureWindowGestureTracker] class.
type IWSEventCaptureWindowGestureTracker interface {
	objectivec.IObject
}

// Init initializes the instance.
func (w WSEventCaptureWindowGestureTracker) Init() WSEventCaptureWindowGestureTracker {
	rv := objc.SendIfResponds[WSEventCaptureWindowGestureTracker](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventCaptureWindowGestureTracker) Autorelease() WSEventCaptureWindowGestureTracker {
	rv := objc.SendIfResponds[WSEventCaptureWindowGestureTracker](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventCaptureWindowGestureTracker creates a new WSEventCaptureWindowGestureTracker instance.
func NewWSEventCaptureWindowGestureTracker() WSEventCaptureWindowGestureTracker {
	class := getWSEventCaptureWindowGestureTrackerClass()
	rv := objc.SendIfResponds[WSEventCaptureWindowGestureTracker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_WSEventCaptureWindowGestureTrackerClass WSEventCaptureWindowGestureTrackerClass) AddEventCaptureTrackerToWindowID(id uint32) {
	objc.SendIfResponds[objc.ID](objc.ID(_WSEventCaptureWindowGestureTrackerClass.class), objc.Sel("addEventCaptureTrackerToWindowID:"), id)
}
func (_WSEventCaptureWindowGestureTrackerClass WSEventCaptureWindowGestureTrackerClass) AddLifetimeTrackerToPhantomWindowIDConnectionManager(id uint32, connection unsafe.Pointer, manager objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_WSEventCaptureWindowGestureTrackerClass.class), objc.Sel("addLifetimeTrackerToPhantomWindowID:connection:manager:"), id, connection, manager)
}
func (_WSEventCaptureWindowGestureTrackerClass WSEventCaptureWindowGestureTrackerClass) RemoveEventCaptureTrackerFromWindowID(id uint32) {
	objc.SendIfResponds[objc.ID](objc.ID(_WSEventCaptureWindowGestureTrackerClass.class), objc.Sel("removeEventCaptureTrackerFromWindowID:"), id)
}
