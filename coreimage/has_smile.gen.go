// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasSmile] class.
var (
	_HasSmileClass     HasSmileClass
	_HasSmileClassOnce sync.Once
)

func getHasSmileClass() HasSmileClass {
	_HasSmileClassOnce.Do(func() {
		_HasSmileClass = HasSmileClass{class: objc.GetClass("hasSmile")}
	})
	return _HasSmileClass
}

// GetHasSmileClass returns the class object for hasSmile.
func GetHasSmileClass() HasSmileClass {
	return getHasSmileClass()
}

type HasSmileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HasSmileClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasSmileClass) Alloc() HasSmile {
	rv := objc.Send[HasSmile](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-c.ivar
type HasSmile struct {
	objectivec.Object
}

// HasSmileFromID constructs a [HasSmile] from an objc.ID.
func HasSmileFromID(id objc.ID) HasSmile {
	return HasSmile{objectivec.Object{ID: id}}
}
// Ensure HasSmile implements IHasSmile.
var _ IHasSmile = HasSmile{}

// An interface definition for the [HasSmile] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-c.ivar
type IHasSmile interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasSmile) Init() HasSmile {
	rv := objc.Send[HasSmile](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasSmile) Autorelease() HasSmile {
	rv := objc.Send[HasSmile](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasSmile creates a new HasSmile instance.
func NewHasSmile() HasSmile {
	class := getHasSmileClass()
	rv := objc.Send[HasSmile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

