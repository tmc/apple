// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLTensorStorageView] class.
var (
	_CoreMLMLTensorStorageViewClass     CoreMLMLTensorStorageViewClass
	_CoreMLMLTensorStorageViewClassOnce sync.Once
)

func getCoreMLMLTensorStorageViewClass() CoreMLMLTensorStorageViewClass {
	_CoreMLMLTensorStorageViewClassOnce.Do(func() {
		_CoreMLMLTensorStorageViewClass = CoreMLMLTensorStorageViewClass{class: objc.GetClass("CoreML.MLTensorStorageView")}
	})
	return _CoreMLMLTensorStorageViewClass
}

// GetCoreMLMLTensorStorageViewClass returns the class object for CoreML.MLTensorStorageView.
func GetCoreMLMLTensorStorageViewClass() CoreMLMLTensorStorageViewClass {
	return getCoreMLMLTensorStorageViewClass()
}

type CoreMLMLTensorStorageViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLTensorStorageViewClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLTensorStorageViewClass) Alloc() CoreMLMLTensorStorageView {
	rv := objc.SendIfResponds[CoreMLMLTensorStorageView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type CoreMLMLTensorStorageView struct {
	objectivec.Object
}

// CoreMLMLTensorStorageViewFromID constructs a [CoreMLMLTensorStorageView] from an objc.ID.
func CoreMLMLTensorStorageViewFromID(id objc.ID) CoreMLMLTensorStorageView {
	return CoreMLMLTensorStorageView{objectivec.Object{ID: id}}
}

// Ensure CoreMLMLTensorStorageView implements ICoreMLMLTensorStorageView.
var _ ICoreMLMLTensorStorageView = CoreMLMLTensorStorageView{}

// An interface definition for the [CoreMLMLTensorStorageView] class.
type ICoreMLMLTensorStorageView interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLTensorStorageView) Init() CoreMLMLTensorStorageView {
	rv := objc.SendIfResponds[CoreMLMLTensorStorageView](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLTensorStorageView) Autorelease() CoreMLMLTensorStorageView {
	rv := objc.SendIfResponds[CoreMLMLTensorStorageView](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLTensorStorageView creates a new CoreMLMLTensorStorageView instance.
func NewCoreMLMLTensorStorageView() CoreMLMLTensorStorageView {
	class := getCoreMLMLTensorStorageViewClass()
	rv := objc.SendIfResponds[CoreMLMLTensorStorageView](objc.ID(class.class), objc.Sel("new"))
	return rv
}
