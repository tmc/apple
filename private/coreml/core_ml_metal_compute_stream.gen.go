// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMetalComputeStream] class.
var (
	_CoreMLMetalComputeStreamClass     CoreMLMetalComputeStreamClass
	_CoreMLMetalComputeStreamClassOnce sync.Once
)

func getCoreMLMetalComputeStreamClass() CoreMLMetalComputeStreamClass {
	_CoreMLMetalComputeStreamClassOnce.Do(func() {
		_CoreMLMetalComputeStreamClass = CoreMLMetalComputeStreamClass{class: objc.GetClass("CoreML.MetalComputeStream")}
	})
	return _CoreMLMetalComputeStreamClass
}

// GetCoreMLMetalComputeStreamClass returns the class object for CoreML.MetalComputeStream.
func GetCoreMLMetalComputeStreamClass() CoreMLMetalComputeStreamClass {
	return getCoreMLMetalComputeStreamClass()
}

type CoreMLMetalComputeStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMetalComputeStreamClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMetalComputeStreamClass) Alloc() CoreMLMetalComputeStream {
	rv := objc.Send[CoreMLMetalComputeStream](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalComputeStream
type CoreMLMetalComputeStream struct {
	objectivec.Object
}

// CoreMLMetalComputeStreamFromID constructs a [CoreMLMetalComputeStream] from an objc.ID.
func CoreMLMetalComputeStreamFromID(id objc.ID) CoreMLMetalComputeStream {
	return CoreMLMetalComputeStream{objectivec.Object{ID: id}}
}

// NOTE: CoreMLMetalComputeStream struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMetalComputeStream embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMetalComputeStream] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MetalComputeStream
type ICoreMLMetalComputeStream interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMetalComputeStream) Init() CoreMLMetalComputeStream {
	rv := objc.Send[CoreMLMetalComputeStream](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMetalComputeStream) Autorelease() CoreMLMetalComputeStream {
	rv := objc.Send[CoreMLMetalComputeStream](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMetalComputeStream creates a new CoreMLMetalComputeStream instance.
func NewCoreMLMetalComputeStream() CoreMLMetalComputeStream {
	class := getCoreMLMetalComputeStreamClass()
	rv := objc.Send[CoreMLMetalComputeStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}
