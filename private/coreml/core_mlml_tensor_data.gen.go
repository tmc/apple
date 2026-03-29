// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLTensorData] class.
var (
	_CoreMLMLTensorDataClass     CoreMLMLTensorDataClass
	_CoreMLMLTensorDataClassOnce sync.Once
)

func getCoreMLMLTensorDataClass() CoreMLMLTensorDataClass {
	_CoreMLMLTensorDataClassOnce.Do(func() {
		_CoreMLMLTensorDataClass = CoreMLMLTensorDataClass{class: objc.GetClass("CoreML.MLTensorData")}
	})
	return _CoreMLMLTensorDataClass
}

// GetCoreMLMLTensorDataClass returns the class object for CoreML.MLTensorData.
func GetCoreMLMLTensorDataClass() CoreMLMLTensorDataClass {
	return getCoreMLMLTensorDataClass()
}

type CoreMLMLTensorDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLTensorDataClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLTensorDataClass) Alloc() CoreMLMLTensorData {
	rv := objc.Send[CoreMLMLTensorData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreML.MLTensorData
type CoreMLMLTensorData struct {
	objectivec.Object
}

// CoreMLMLTensorDataFromID constructs a [CoreMLMLTensorData] from an objc.ID.
func CoreMLMLTensorDataFromID(id objc.ID) CoreMLMLTensorData {
	return CoreMLMLTensorData{objectivec.Object{ID: id}}
}
// NOTE: CoreMLMLTensorData struct embeds objectivec.Object (parent type unavailable) but
// ICoreMLMLTensorData embeds the parent interface; skip compile-time assertion.

// An interface definition for the [CoreMLMLTensorData] class.
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLTensorData
type ICoreMLMLTensorData interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLTensorData) Init() CoreMLMLTensorData {
	rv := objc.Send[CoreMLMLTensorData](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLTensorData) Autorelease() CoreMLMLTensorData {
	rv := objc.Send[CoreMLMLTensorData](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLTensorData creates a new CoreMLMLTensorData instance.
func NewCoreMLMLTensorData() CoreMLMLTensorData {
	class := getCoreMLMLTensorDataClass()
	rv := objc.Send[CoreMLMLTensorData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

