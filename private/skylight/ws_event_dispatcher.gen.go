// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventDispatcher] class.
var (
	_WSEventDispatcherClass     WSEventDispatcherClass
	_WSEventDispatcherClassOnce sync.Once
)

func getWSEventDispatcherClass() WSEventDispatcherClass {
	_WSEventDispatcherClassOnce.Do(func() {
		_WSEventDispatcherClass = WSEventDispatcherClass{class: objc.GetClass("WSEventDispatcher")}
	})
	return _WSEventDispatcherClass
}

// GetWSEventDispatcherClass returns the class object for WSEventDispatcher.
func GetWSEventDispatcherClass() WSEventDispatcherClass {
	return getWSEventDispatcherClass()
}

type WSEventDispatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventDispatcherClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventDispatcherClass) Alloc() WSEventDispatcher {
	rv := objc.SendIfResponds[WSEventDispatcher](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventDispatcher.PostBackgroundEvent]
//   - [WSEventDispatcher.PostEventToConnectionID]
//   - [WSEventDispatcher.PostEventToDestination]
//   - [WSEventDispatcher.DebugDescription]
//   - [WSEventDispatcher.Description]
//   - [WSEventDispatcher.Hash]
//   - [WSEventDispatcher.Superclass]
type WSEventDispatcher struct {
	objectivec.Object
}

// WSEventDispatcherFromID constructs a [WSEventDispatcher] from an objc.ID.
func WSEventDispatcherFromID(id objc.ID) WSEventDispatcher {
	return WSEventDispatcher{objectivec.Object{ID: id}}
}

// Ensure WSEventDispatcher implements IWSEventDispatcher.
var _ IWSEventDispatcher = WSEventDispatcher{}

// An interface definition for the [WSEventDispatcher] class.
//
// # Methods
//
//   - [IWSEventDispatcher.PostBackgroundEvent]
//   - [IWSEventDispatcher.PostEventToConnectionID]
//   - [IWSEventDispatcher.PostEventToDestination]
//   - [IWSEventDispatcher.DebugDescription]
//   - [IWSEventDispatcher.Description]
//   - [IWSEventDispatcher.Hash]
//   - [IWSEventDispatcher.Superclass]
type IWSEventDispatcher interface {
	objectivec.IObject

	// Topic: Methods

	PostBackgroundEvent(event *SLSEventRecord)
	PostEventToConnectionID(event *SLSEventRecord, id uint32)
	PostEventToDestination(event *SLSEventRecord, destination objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSEventDispatcher) Init() WSEventDispatcher {
	rv := objc.SendIfResponds[WSEventDispatcher](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventDispatcher) Autorelease() WSEventDispatcher {
	rv := objc.SendIfResponds[WSEventDispatcher](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventDispatcher creates a new WSEventDispatcher instance.
func NewWSEventDispatcher() WSEventDispatcher {
	class := getWSEventDispatcherClass()
	rv := objc.SendIfResponds[WSEventDispatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSEventDispatcher) PostBackgroundEvent(event *SLSEventRecord) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("postBackgroundEvent:"), unsafe.Pointer(event))
}
func (w WSEventDispatcher) PostEventToConnectionID(event *SLSEventRecord, id uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("postEvent:toConnectionID:"), unsafe.Pointer(event), id)
}
func (w WSEventDispatcher) PostEventToDestination(event *SLSEventRecord, destination objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("postEvent:toDestination:"), unsafe.Pointer(event), destination)
}

func (w WSEventDispatcher) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventDispatcher) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventDispatcher) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSEventDispatcher) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
