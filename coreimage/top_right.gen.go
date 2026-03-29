// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TopRight] class.
var (
	_TopRightClass     TopRightClass
	_TopRightClassOnce sync.Once
)

func getTopRightClass() TopRightClass {
	_TopRightClassOnce.Do(func() {
		_TopRightClass = TopRightClass{class: objc.GetClass("topRight")}
	})
	return _TopRightClass
}

// GetTopRightClass returns the class object for topRight.
func GetTopRightClass() TopRightClass {
	return getTopRightClass()
}

type TopRightClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TopRightClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TopRightClass) Alloc() TopRight {
	rv := objc.Send[TopRight](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-c.ivar
type TopRight struct {
	objectivec.Object
}

// TopRightFromID constructs a [TopRight] from an objc.ID.
func TopRightFromID(id objc.ID) TopRight {
	return TopRight{objectivec.Object{ID: id}}
}
// Ensure TopRight implements ITopRight.
var _ ITopRight = TopRight{}

// An interface definition for the [TopRight] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-c.ivar
type ITopRight interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TopRight) Init() TopRight {
	rv := objc.Send[TopRight](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TopRight) Autorelease() TopRight {
	rv := objc.Send[TopRight](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTopRight creates a new TopRight instance.
func NewTopRight() TopRight {
	class := getTopRightClass()
	rv := objc.Send[TopRight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

