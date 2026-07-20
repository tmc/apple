// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SharedEvent] class.
var (
	_SharedEventClass     SharedEventClass
	_SharedEventClassOnce sync.Once
)

func getSharedEventClass() SharedEventClass {
	_SharedEventClassOnce.Do(func() {
		_SharedEventClass = SharedEventClass{class: objc.GetClass("_TtCC6CoreML11MetalDevice11SharedEvent")}
	})
	return _SharedEventClass
}

// GetSharedEventClass returns the class object for _TtCC6CoreML11MetalDevice11SharedEvent.
func GetSharedEventClass() SharedEventClass {
	return getSharedEventClass()
}

type SharedEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SharedEventClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SharedEventClass) Alloc() SharedEvent {
	rv := objc.Send[SharedEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SharedEvent struct {
	objectivec.Object
}

// SharedEventFromID constructs a [SharedEvent] from an objc.ID.
func SharedEventFromID(id objc.ID) SharedEvent {
	return SharedEvent{objectivec.Object{ID: id}}
}

// NOTE: SharedEvent struct embeds objectivec.Object (parent type unavailable) but
// ISharedEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SharedEvent] class.
type ISharedEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SharedEvent) Init() SharedEvent {
	rv := objc.Send[SharedEvent](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SharedEvent) Autorelease() SharedEvent {
	rv := objc.Send[SharedEvent](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharedEvent creates a new SharedEvent instance.
func NewSharedEvent() SharedEvent {
	class := getSharedEventClass()
	rv := objc.Send[SharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
