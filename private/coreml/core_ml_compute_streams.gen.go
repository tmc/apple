// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLComputeStreams] class.
var (
	_CoreMLComputeStreamsClass     CoreMLComputeStreamsClass
	_CoreMLComputeStreamsClassOnce sync.Once
)

func getCoreMLComputeStreamsClass() CoreMLComputeStreamsClass {
	_CoreMLComputeStreamsClassOnce.Do(func() {
		_CoreMLComputeStreamsClass = CoreMLComputeStreamsClass{class: objc.GetClass("CoreML.ComputeStreams")}
	})
	return _CoreMLComputeStreamsClass
}

// GetCoreMLComputeStreamsClass returns the class object for CoreML.ComputeStreams.
func GetCoreMLComputeStreamsClass() CoreMLComputeStreamsClass {
	return getCoreMLComputeStreamsClass()
}

type CoreMLComputeStreamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLComputeStreamsClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLComputeStreamsClass) Alloc() CoreMLComputeStreams {
	rv := objc.Send[CoreMLComputeStreams](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.ComputeStreams
type CoreMLComputeStreams struct {
	objectivec.Object
}

// CoreMLComputeStreamsFromID constructs a [CoreMLComputeStreams] from an objc.ID.
func CoreMLComputeStreamsFromID(id objc.ID) CoreMLComputeStreams {
	return CoreMLComputeStreams{objectivec.Object{ID: id}}
}

// NOTE: CoreMLComputeStreams struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLComputeStreams embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLComputeStreams] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.ComputeStreams
type ICoreMLComputeStreams interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLComputeStreams) Init() CoreMLComputeStreams {
	rv := objc.Send[CoreMLComputeStreams](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLComputeStreams) Autorelease() CoreMLComputeStreams {
	rv := objc.Send[CoreMLComputeStreams](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLComputeStreams creates a new CoreMLComputeStreams instance.
func NewCoreMLComputeStreams() CoreMLComputeStreams {
	class := getCoreMLComputeStreamsClass()
	rv := objc.Send[CoreMLComputeStreams](objc.ID(class.class), objc.Sel("new"))
	return rv
}
