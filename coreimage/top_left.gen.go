// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TopLeft] class.
var (
	_TopLeftClass     TopLeftClass
	_TopLeftClassOnce sync.Once
)

func getTopLeftClass() TopLeftClass {
	_TopLeftClassOnce.Do(func() {
		_TopLeftClass = TopLeftClass{class: objc.GetClass("topLeft")}
	})
	return _TopLeftClass
}

// GetTopLeftClass returns the class object for topLeft.
func GetTopLeftClass() TopLeftClass {
	return getTopLeftClass()
}

type TopLeftClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TopLeftClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TopLeftClass) Alloc() TopLeft {
	rv := objc.Send[TopLeft](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topLeft-c.ivar
type TopLeft struct {
	objectivec.Object
}

// TopLeftFromID constructs a [TopLeft] from an objc.ID.
func TopLeftFromID(id objc.ID) TopLeft {
	return TopLeft{objectivec.Object{ID: id}}
}

// Ensure TopLeft implements ITopLeft.
var _ ITopLeft = TopLeft{}

// An interface definition for the [TopLeft] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topLeft-c.ivar
type ITopLeft interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TopLeft) Init() TopLeft {
	rv := objc.Send[TopLeft](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TopLeft) Autorelease() TopLeft {
	rv := objc.Send[TopLeft](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTopLeft creates a new TopLeft instance.
func NewTopLeft() TopLeft {
	class := getTopLeftClass()
	rv := objc.Send[TopLeft](objc.ID(class.class), objc.Sel("new"))
	return rv
}
