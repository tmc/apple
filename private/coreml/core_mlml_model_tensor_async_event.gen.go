// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLModelTensorAsyncEvent] class.
var (
	_CoreMLMLModelTensorAsyncEventClass     CoreMLMLModelTensorAsyncEventClass
	_CoreMLMLModelTensorAsyncEventClassOnce sync.Once
)

func getCoreMLMLModelTensorAsyncEventClass() CoreMLMLModelTensorAsyncEventClass {
	_CoreMLMLModelTensorAsyncEventClassOnce.Do(func() {
		_CoreMLMLModelTensorAsyncEventClass = CoreMLMLModelTensorAsyncEventClass{class: objc.GetClass("CoreML.MLModelTensorAsyncEvent")}
	})
	return _CoreMLMLModelTensorAsyncEventClass
}

// GetCoreMLMLModelTensorAsyncEventClass returns the class object for CoreML.MLModelTensorAsyncEvent.
func GetCoreMLMLModelTensorAsyncEventClass() CoreMLMLModelTensorAsyncEventClass {
	return getCoreMLMLModelTensorAsyncEventClass()
}

type CoreMLMLModelTensorAsyncEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLModelTensorAsyncEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLModelTensorAsyncEventClass) Alloc() CoreMLMLModelTensorAsyncEvent {
	rv := objc.Send[CoreMLMLModelTensorAsyncEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MLModelTensorAsyncEvent
type CoreMLMLModelTensorAsyncEvent struct {
	objectivec.Object
}

// CoreMLMLModelTensorAsyncEventFromID constructs a [CoreMLMLModelTensorAsyncEvent] from an objc.ID.
func CoreMLMLModelTensorAsyncEventFromID(id objc.ID) CoreMLMLModelTensorAsyncEvent {
	return CoreMLMLModelTensorAsyncEvent{objectivec.Object{ID: id}}
}

// NOTE: CoreMLMLModelTensorAsyncEvent struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMLModelTensorAsyncEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMLModelTensorAsyncEvent] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLModelTensorAsyncEvent
type ICoreMLMLModelTensorAsyncEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLModelTensorAsyncEvent) Init() CoreMLMLModelTensorAsyncEvent {
	rv := objc.Send[CoreMLMLModelTensorAsyncEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLModelTensorAsyncEvent) Autorelease() CoreMLMLModelTensorAsyncEvent {
	rv := objc.Send[CoreMLMLModelTensorAsyncEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLModelTensorAsyncEvent creates a new CoreMLMLModelTensorAsyncEvent instance.
func NewCoreMLMLModelTensorAsyncEvent() CoreMLMLModelTensorAsyncEvent {
	class := getCoreMLMLModelTensorAsyncEventClass()
	rv := objc.Send[CoreMLMLModelTensorAsyncEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
