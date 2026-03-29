// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMetalAsyncEvent] class.
var (
	_CoreMLMetalAsyncEventClass     CoreMLMetalAsyncEventClass
	_CoreMLMetalAsyncEventClassOnce sync.Once
)

func getCoreMLMetalAsyncEventClass() CoreMLMetalAsyncEventClass {
	_CoreMLMetalAsyncEventClassOnce.Do(func() {
		_CoreMLMetalAsyncEventClass = CoreMLMetalAsyncEventClass{class: objc.GetClass("CoreML.MetalAsyncEvent")}
	})
	return _CoreMLMetalAsyncEventClass
}

// GetCoreMLMetalAsyncEventClass returns the class object for CoreML.MetalAsyncEvent.
func GetCoreMLMetalAsyncEventClass() CoreMLMetalAsyncEventClass {
	return getCoreMLMetalAsyncEventClass()
}

type CoreMLMetalAsyncEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMetalAsyncEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMetalAsyncEventClass) Alloc() CoreMLMetalAsyncEvent {
	rv := objc.Send[CoreMLMetalAsyncEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalAsyncEvent
type CoreMLMetalAsyncEvent struct {
	objectivec.Object
}

// CoreMLMetalAsyncEventFromID constructs a [CoreMLMetalAsyncEvent] from an objc.ID.
func CoreMLMetalAsyncEventFromID(id objc.ID) CoreMLMetalAsyncEvent {
	return CoreMLMetalAsyncEvent{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMetalAsyncEvent struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMetalAsyncEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMetalAsyncEvent] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalAsyncEvent
type ICoreMLMetalAsyncEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMetalAsyncEvent) Init() CoreMLMetalAsyncEvent {
	rv := objc.Send[CoreMLMetalAsyncEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMetalAsyncEvent) Autorelease() CoreMLMetalAsyncEvent {
	rv := objc.Send[CoreMLMetalAsyncEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMetalAsyncEvent creates a new CoreMLMetalAsyncEvent instance.
func NewCoreMLMetalAsyncEvent() CoreMLMetalAsyncEvent {
	class := getCoreMLMetalAsyncEventClass()
	rv := objc.Send[CoreMLMetalAsyncEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

