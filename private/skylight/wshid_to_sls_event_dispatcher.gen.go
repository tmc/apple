// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDToSLSEventDispatcher] class.
var (
	_WSHIDToSLSEventDispatcherClass     WSHIDToSLSEventDispatcherClass
	_WSHIDToSLSEventDispatcherClassOnce sync.Once
)

func getWSHIDToSLSEventDispatcherClass() WSHIDToSLSEventDispatcherClass {
	_WSHIDToSLSEventDispatcherClassOnce.Do(func() {
		_WSHIDToSLSEventDispatcherClass = WSHIDToSLSEventDispatcherClass{class: objc.GetClass("WSHIDToSLSEventDispatcher")}
	})
	return _WSHIDToSLSEventDispatcherClass
}

// GetWSHIDToSLSEventDispatcherClass returns the class object for WSHIDToSLSEventDispatcher.
func GetWSHIDToSLSEventDispatcherClass() WSHIDToSLSEventDispatcherClass {
	return getWSHIDToSLSEventDispatcherClass()
}

type WSHIDToSLSEventDispatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDToSLSEventDispatcherClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDToSLSEventDispatcherClass) Alloc() WSHIDToSLSEventDispatcher {
	rv := objc.SendIfResponds[WSHIDToSLSEventDispatcher](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDToSLSEventDispatcher.PostHIDEventForTargetIdentifierToConnectionID]
//   - [WSHIDToSLSEventDispatcher.DebugDescription]
//   - [WSHIDToSLSEventDispatcher.Description]
//   - [WSHIDToSLSEventDispatcher.Hash]
//   - [WSHIDToSLSEventDispatcher.Superclass]
type WSHIDToSLSEventDispatcher struct {
	objectivec.Object
}

// WSHIDToSLSEventDispatcherFromID constructs a [WSHIDToSLSEventDispatcher] from an objc.ID.
func WSHIDToSLSEventDispatcherFromID(id objc.ID) WSHIDToSLSEventDispatcher {
	return WSHIDToSLSEventDispatcher{objectivec.Object{ID: id}}
}

// Ensure WSHIDToSLSEventDispatcher implements IWSHIDToSLSEventDispatcher.
var _ IWSHIDToSLSEventDispatcher = WSHIDToSLSEventDispatcher{}

// An interface definition for the [WSHIDToSLSEventDispatcher] class.
//
// # Methods
//
//   - [IWSHIDToSLSEventDispatcher.PostHIDEventForTargetIdentifierToConnectionID]
//   - [IWSHIDToSLSEventDispatcher.DebugDescription]
//   - [IWSHIDToSLSEventDispatcher.Description]
//   - [IWSHIDToSLSEventDispatcher.Hash]
//   - [IWSHIDToSLSEventDispatcher.Superclass]
type IWSHIDToSLSEventDispatcher interface {
	objectivec.IObject

	// Topic: Methods

	PostHIDEventForTargetIdentifierToConnectionID(hIDEvent uintptr, identifier unsafe.Pointer, id uint32)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSHIDToSLSEventDispatcher) Init() WSHIDToSLSEventDispatcher {
	rv := objc.SendIfResponds[WSHIDToSLSEventDispatcher](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDToSLSEventDispatcher) Autorelease() WSHIDToSLSEventDispatcher {
	rv := objc.SendIfResponds[WSHIDToSLSEventDispatcher](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDToSLSEventDispatcher creates a new WSHIDToSLSEventDispatcher instance.
func NewWSHIDToSLSEventDispatcher() WSHIDToSLSEventDispatcher {
	class := getWSHIDToSLSEventDispatcherClass()
	rv := objc.SendIfResponds[WSHIDToSLSEventDispatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSHIDToSLSEventDispatcher) PostHIDEventForTargetIdentifierToConnectionID(hIDEvent uintptr, identifier unsafe.Pointer, id uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("postHIDEvent:forTargetIdentifier:toConnectionID:"), hIDEvent, identifier, id)
}

func (w WSHIDToSLSEventDispatcher) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDToSLSEventDispatcher) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDToSLSEventDispatcher) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSHIDToSLSEventDispatcher) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
