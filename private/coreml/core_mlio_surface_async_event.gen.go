// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLIOSurfaceAsyncEvent] class.
var (
	_CoreMLIOSurfaceAsyncEventClass     CoreMLIOSurfaceAsyncEventClass
	_CoreMLIOSurfaceAsyncEventClassOnce sync.Once
)

func getCoreMLIOSurfaceAsyncEventClass() CoreMLIOSurfaceAsyncEventClass {
	_CoreMLIOSurfaceAsyncEventClassOnce.Do(func() {
		_CoreMLIOSurfaceAsyncEventClass = CoreMLIOSurfaceAsyncEventClass{class: objc.GetClass("CoreML.IOSurfaceAsyncEvent")}
	})
	return _CoreMLIOSurfaceAsyncEventClass
}

// GetCoreMLIOSurfaceAsyncEventClass returns the class object for CoreML.IOSurfaceAsyncEvent.
func GetCoreMLIOSurfaceAsyncEventClass() CoreMLIOSurfaceAsyncEventClass {
	return getCoreMLIOSurfaceAsyncEventClass()
}

type CoreMLIOSurfaceAsyncEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLIOSurfaceAsyncEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLIOSurfaceAsyncEventClass) Alloc() CoreMLIOSurfaceAsyncEvent {
	rv := objc.Send[CoreMLIOSurfaceAsyncEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceAsyncEvent
type CoreMLIOSurfaceAsyncEvent struct {
	objectivec.Object
}

// CoreMLIOSurfaceAsyncEventFromID constructs a [CoreMLIOSurfaceAsyncEvent] from an objc.ID.
func CoreMLIOSurfaceAsyncEventFromID(id objc.ID) CoreMLIOSurfaceAsyncEvent {
	return CoreMLIOSurfaceAsyncEvent{objectivec.Object{ID: id}}
}

// NOTE: CoreMLIOSurfaceAsyncEvent struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLIOSurfaceAsyncEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLIOSurfaceAsyncEvent] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.IOSurfaceAsyncEvent
type ICoreMLIOSurfaceAsyncEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLIOSurfaceAsyncEvent) Init() CoreMLIOSurfaceAsyncEvent {
	rv := objc.Send[CoreMLIOSurfaceAsyncEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLIOSurfaceAsyncEvent) Autorelease() CoreMLIOSurfaceAsyncEvent {
	rv := objc.Send[CoreMLIOSurfaceAsyncEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLIOSurfaceAsyncEvent creates a new CoreMLIOSurfaceAsyncEvent instance.
func NewCoreMLIOSurfaceAsyncEvent() CoreMLIOSurfaceAsyncEvent {
	class := getCoreMLIOSurfaceAsyncEventClass()
	rv := objc.Send[CoreMLIOSurfaceAsyncEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
