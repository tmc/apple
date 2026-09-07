// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSGestureEventCollector] class.
var (
	_WSGestureEventCollectorClass     WSGestureEventCollectorClass
	_WSGestureEventCollectorClassOnce sync.Once
)

func getWSGestureEventCollectorClass() WSGestureEventCollectorClass {
	_WSGestureEventCollectorClassOnce.Do(func() {
		_WSGestureEventCollectorClass = WSGestureEventCollectorClass{class: objc.GetClass("WSGestureEventCollector")}
	})
	return _WSGestureEventCollectorClass
}

// GetWSGestureEventCollectorClass returns the class object for WSGestureEventCollector.
func GetWSGestureEventCollectorClass() WSGestureEventCollectorClass {
	return getWSGestureEventCollectorClass()
}

type WSGestureEventCollectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSGestureEventCollectorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSGestureEventCollectorClass) Alloc() WSGestureEventCollector {
	rv := objc.SendIfResponds[WSGestureEventCollector](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSGestureEventCollector.AddGestureEventForTargetIdentifier]
//   - [WSGestureEventCollector.AddTopLevelEventForTargetIdentifierConnectionID]
//   - [WSGestureEventCollector.BatchingAssertion]
//   - [WSGestureEventCollector.SetBatchingAssertion]
//   - [WSGestureEventCollector.DispatchEventsIfNecessary]
//   - [WSGestureEventCollector.HidEventDispatcher]
//   - [WSGestureEventCollector.StartBatchingForReason]
//   - [WSGestureEventCollector.InitWithHIDEventDispatcher]
type WSGestureEventCollector struct {
	objectivec.Object
}

// WSGestureEventCollectorFromID constructs a [WSGestureEventCollector] from an objc.ID.
func WSGestureEventCollectorFromID(id objc.ID) WSGestureEventCollector {
	return WSGestureEventCollector{objectivec.Object{ID: id}}
}

// Ensure WSGestureEventCollector implements IWSGestureEventCollector.
var _ IWSGestureEventCollector = WSGestureEventCollector{}

// An interface definition for the [WSGestureEventCollector] class.
//
// # Methods
//
//   - [IWSGestureEventCollector.AddGestureEventForTargetIdentifier]
//   - [IWSGestureEventCollector.AddTopLevelEventForTargetIdentifierConnectionID]
//   - [IWSGestureEventCollector.BatchingAssertion]
//   - [IWSGestureEventCollector.SetBatchingAssertion]
//   - [IWSGestureEventCollector.DispatchEventsIfNecessary]
//   - [IWSGestureEventCollector.HidEventDispatcher]
//   - [IWSGestureEventCollector.StartBatchingForReason]
//   - [IWSGestureEventCollector.InitWithHIDEventDispatcher]
type IWSGestureEventCollector interface {
	objectivec.IObject

	// Topic: Methods

	AddGestureEventForTargetIdentifier(event uintptr, identifier unsafe.Pointer)
	AddTopLevelEventForTargetIdentifierConnectionID(event uintptr, identifier unsafe.Pointer, id uint32)
	BatchingAssertion() objectivec.IObject
	SetBatchingAssertion(value objectivec.IObject)
	DispatchEventsIfNecessary()
	HidEventDispatcher() unsafe.Pointer
	StartBatchingForReason(reason objectivec.IObject) objectivec.IObject
	InitWithHIDEventDispatcher(dispatcher objectivec.IObject) WSGestureEventCollector
}

// Init initializes the instance.
func (w WSGestureEventCollector) Init() WSGestureEventCollector {
	rv := objc.SendIfResponds[WSGestureEventCollector](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSGestureEventCollector) Autorelease() WSGestureEventCollector {
	rv := objc.SendIfResponds[WSGestureEventCollector](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSGestureEventCollector creates a new WSGestureEventCollector instance.
func NewWSGestureEventCollector() WSGestureEventCollector {
	class := getWSGestureEventCollectorClass()
	rv := objc.SendIfResponds[WSGestureEventCollector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSGestureEventCollectorWithHIDEventDispatcher(dispatcher objectivec.IObject) WSGestureEventCollector {
	instance := getWSGestureEventCollectorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHIDEventDispatcher:"), dispatcher)
	return WSGestureEventCollectorFromID(rv)
}

func (w WSGestureEventCollector) AddGestureEventForTargetIdentifier(event uintptr, identifier unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("addGestureEvent:forTargetIdentifier:"), event, identifier)
}
func (w WSGestureEventCollector) AddTopLevelEventForTargetIdentifierConnectionID(event uintptr, identifier unsafe.Pointer, id uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("addTopLevelEvent:forTargetIdentifier:connectionID:"), event, identifier, id)
}
func (w WSGestureEventCollector) DispatchEventsIfNecessary() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("dispatchEventsIfNecessary"))
}
func (w WSGestureEventCollector) StartBatchingForReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("startBatchingForReason:"), reason)
	return objectivec.Object{ID: rv}
}
func (w WSGestureEventCollector) InitWithHIDEventDispatcher(dispatcher objectivec.IObject) WSGestureEventCollector {
	rv := objc.SendIfResponds[WSGestureEventCollector](w.ID, objc.Sel("initWithHIDEventDispatcher:"), dispatcher)
	return rv
}

func (w WSGestureEventCollector) BatchingAssertion() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("batchingAssertion"))
	return objectivec.Object{ID: rv}
}
func (w WSGestureEventCollector) SetBatchingAssertion(value objectivec.IObject) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setBatchingAssertion:"), value)
}
func (w WSGestureEventCollector) HidEventDispatcher() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("hidEventDispatcher"))
	return rv
}
