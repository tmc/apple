// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLBNNSComputeStream] class.
var (
	_CoreMLBNNSComputeStreamClass     CoreMLBNNSComputeStreamClass
	_CoreMLBNNSComputeStreamClassOnce sync.Once
)

func getCoreMLBNNSComputeStreamClass() CoreMLBNNSComputeStreamClass {
	_CoreMLBNNSComputeStreamClassOnce.Do(func() {
		_CoreMLBNNSComputeStreamClass = CoreMLBNNSComputeStreamClass{class: objc.GetClass("CoreML.BNNSComputeStream")}
	})
	return _CoreMLBNNSComputeStreamClass
}

// GetCoreMLBNNSComputeStreamClass returns the class object for CoreML.BNNSComputeStream.
func GetCoreMLBNNSComputeStreamClass() CoreMLBNNSComputeStreamClass {
	return getCoreMLBNNSComputeStreamClass()
}

type CoreMLBNNSComputeStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLBNNSComputeStreamClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLBNNSComputeStreamClass) Alloc() CoreMLBNNSComputeStream {
	rv := objc.Send[CoreMLBNNSComputeStream](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.BNNSComputeStream
type CoreMLBNNSComputeStream struct {
	objectivec.Object
}

// CoreMLBNNSComputeStreamFromID constructs a [CoreMLBNNSComputeStream] from an objc.ID.
func CoreMLBNNSComputeStreamFromID(id objc.ID) CoreMLBNNSComputeStream {
	return CoreMLBNNSComputeStream{objectivec.Object{ID: id}}
}
// NOTE: CoreMLBNNSComputeStream struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLBNNSComputeStream embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLBNNSComputeStream] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.BNNSComputeStream
type ICoreMLBNNSComputeStream interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLBNNSComputeStream) Init() CoreMLBNNSComputeStream {
	rv := objc.Send[CoreMLBNNSComputeStream](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLBNNSComputeStream) Autorelease() CoreMLBNNSComputeStream {
	rv := objc.Send[CoreMLBNNSComputeStream](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLBNNSComputeStream creates a new CoreMLBNNSComputeStream instance.
func NewCoreMLBNNSComputeStream() CoreMLBNNSComputeStream {
	class := getCoreMLBNNSComputeStreamClass()
	rv := objc.Send[CoreMLBNNSComputeStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

