// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LayerCount] class.
var (
	_LayerCountClass     LayerCountClass
	_LayerCountClassOnce sync.Once
)

func getLayerCountClass() LayerCountClass {
	_LayerCountClassOnce.Do(func() {
		_LayerCountClass = LayerCountClass{class: objc.GetClass("layerCount")}
	})
	return _LayerCountClass
}

// GetLayerCountClass returns the class object for layerCount.
func GetLayerCountClass() LayerCountClass {
	return getLayerCountClass()
}

type LayerCountClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LayerCountClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LayerCountClass) Alloc() LayerCount {
	rv := objc.Send[LayerCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar
type LayerCount struct {
	objectivec.Object
}

// LayerCountFromID constructs a [LayerCount] from an objc.ID.
func LayerCountFromID(id objc.ID) LayerCount {
	return LayerCount{objectivec.Object{ID: id}}
}
// Ensure LayerCount implements ILayerCount.
var _ ILayerCount = LayerCount{}

// An interface definition for the [LayerCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar
type ILayerCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l LayerCount) Init() LayerCount {
	rv := objc.Send[LayerCount](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LayerCount) Autorelease() LayerCount {
	rv := objc.Send[LayerCount](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayerCount creates a new LayerCount instance.
func NewLayerCount() LayerCount {
	class := getLayerCountClass()
	rv := objc.Send[LayerCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}

