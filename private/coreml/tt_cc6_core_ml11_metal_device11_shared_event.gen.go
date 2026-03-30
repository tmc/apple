// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC6CoreML11MetalDevice11SharedEvent] class.
var (
	_TtCC6CoreML11MetalDevice11SharedEventClass     TtCC6CoreML11MetalDevice11SharedEventClass
	_TtCC6CoreML11MetalDevice11SharedEventClassOnce sync.Once
)

func getTtCC6CoreML11MetalDevice11SharedEventClass() TtCC6CoreML11MetalDevice11SharedEventClass {
	_TtCC6CoreML11MetalDevice11SharedEventClassOnce.Do(func() {
		_TtCC6CoreML11MetalDevice11SharedEventClass = TtCC6CoreML11MetalDevice11SharedEventClass{class: objc.GetClass("_TtCC6CoreML11MetalDevice11SharedEvent")}
	})
	return _TtCC6CoreML11MetalDevice11SharedEventClass
}

// GetTtCC6CoreML11MetalDevice11SharedEventClass returns the class object for _TtCC6CoreML11MetalDevice11SharedEvent.
func GetTtCC6CoreML11MetalDevice11SharedEventClass() TtCC6CoreML11MetalDevice11SharedEventClass {
	return getTtCC6CoreML11MetalDevice11SharedEventClass()
}

type TtCC6CoreML11MetalDevice11SharedEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC6CoreML11MetalDevice11SharedEventClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC6CoreML11MetalDevice11SharedEventClass) Alloc() TtCC6CoreML11MetalDevice11SharedEvent {
	rv := objc.Send[TtCC6CoreML11MetalDevice11SharedEvent](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML11MetalDevice11SharedEvent
type TtCC6CoreML11MetalDevice11SharedEvent struct {
	objectivec.Object
}

// TtCC6CoreML11MetalDevice11SharedEventFromID constructs a [TtCC6CoreML11MetalDevice11SharedEvent] from an objc.ID.
func TtCC6CoreML11MetalDevice11SharedEventFromID(id objc.ID) TtCC6CoreML11MetalDevice11SharedEvent {
	return TtCC6CoreML11MetalDevice11SharedEvent{objectivec.Object{ID: id}}
}

// NOTE: TtCC6CoreML11MetalDevice11SharedEvent struct embeds objectivec.Object (parent type unavailable) but
// ITtCC6CoreML11MetalDevice11SharedEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC6CoreML11MetalDevice11SharedEvent] class.
//
// See: https://developer.apple.com/documentation/CoreML/_TtCC6CoreML11MetalDevice11SharedEvent
type ITtCC6CoreML11MetalDevice11SharedEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC6CoreML11MetalDevice11SharedEvent) Init() TtCC6CoreML11MetalDevice11SharedEvent {
	rv := objc.Send[TtCC6CoreML11MetalDevice11SharedEvent](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC6CoreML11MetalDevice11SharedEvent) Autorelease() TtCC6CoreML11MetalDevice11SharedEvent {
	rv := objc.Send[TtCC6CoreML11MetalDevice11SharedEvent](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC6CoreML11MetalDevice11SharedEvent creates a new TtCC6CoreML11MetalDevice11SharedEvent instance.
func NewTtCC6CoreML11MetalDevice11SharedEvent() TtCC6CoreML11MetalDevice11SharedEvent {
	class := getTtCC6CoreML11MetalDevice11SharedEventClass()
	rv := objc.Send[TtCC6CoreML11MetalDevice11SharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
