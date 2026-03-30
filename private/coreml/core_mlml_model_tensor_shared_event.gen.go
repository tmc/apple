// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLModelTensorSharedEvent] class.
var (
	_CoreMLMLModelTensorSharedEventClass     CoreMLMLModelTensorSharedEventClass
	_CoreMLMLModelTensorSharedEventClassOnce sync.Once
)

func getCoreMLMLModelTensorSharedEventClass() CoreMLMLModelTensorSharedEventClass {
	_CoreMLMLModelTensorSharedEventClassOnce.Do(func() {
		_CoreMLMLModelTensorSharedEventClass = CoreMLMLModelTensorSharedEventClass{class: objc.GetClass("CoreML.MLModelTensorSharedEvent")}
	})
	return _CoreMLMLModelTensorSharedEventClass
}

// GetCoreMLMLModelTensorSharedEventClass returns the class object for CoreML.MLModelTensorSharedEvent.
func GetCoreMLMLModelTensorSharedEventClass() CoreMLMLModelTensorSharedEventClass {
	return getCoreMLMLModelTensorSharedEventClass()
}

type CoreMLMLModelTensorSharedEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLModelTensorSharedEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLModelTensorSharedEventClass) Alloc() CoreMLMLModelTensorSharedEvent {
	rv := objc.Send[CoreMLMLModelTensorSharedEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MLModelTensorSharedEvent
type CoreMLMLModelTensorSharedEvent struct {
	objectivec.Object
}

// CoreMLMLModelTensorSharedEventFromID constructs a [CoreMLMLModelTensorSharedEvent] from an objc.ID.
func CoreMLMLModelTensorSharedEventFromID(id objc.ID) CoreMLMLModelTensorSharedEvent {
	return CoreMLMLModelTensorSharedEvent{objectivec.Object{ID: id}}
}

// NOTE: CoreMLMLModelTensorSharedEvent struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMLModelTensorSharedEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMLModelTensorSharedEvent] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLModelTensorSharedEvent
type ICoreMLMLModelTensorSharedEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLModelTensorSharedEvent) Init() CoreMLMLModelTensorSharedEvent {
	rv := objc.Send[CoreMLMLModelTensorSharedEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLModelTensorSharedEvent) Autorelease() CoreMLMLModelTensorSharedEvent {
	rv := objc.Send[CoreMLMLModelTensorSharedEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLModelTensorSharedEvent creates a new CoreMLMLModelTensorSharedEvent instance.
func NewCoreMLMLModelTensorSharedEvent() CoreMLMLModelTensorSharedEvent {
	class := getCoreMLMLModelTensorSharedEventClass()
	rv := objc.Send[CoreMLMLModelTensorSharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
