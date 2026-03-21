// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BottomLeft] class.
var (
	_BottomLeftClass     BottomLeftClass
	_BottomLeftClassOnce sync.Once
)

func getBottomLeftClass() BottomLeftClass {
	_BottomLeftClassOnce.Do(func() {
		_BottomLeftClass = BottomLeftClass{class: objc.GetClass("bottomLeft")}
	})
	return _BottomLeftClass
}

// GetBottomLeftClass returns the class object for bottomLeft.
func GetBottomLeftClass() BottomLeftClass {
	return getBottomLeftClass()
}

type BottomLeftClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BottomLeftClass) Alloc() BottomLeft {
	rv := objc.Send[BottomLeft](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-c.ivar
type BottomLeft struct {
	objectivec.Object
}

// BottomLeftFromID constructs a [BottomLeft] from an objc.ID.
func BottomLeftFromID(id objc.ID) BottomLeft {
	return BottomLeft{objectivec.Object{ID: id}}
}
// Ensure BottomLeft implements IBottomLeft.
var _ IBottomLeft = BottomLeft{}

// An interface definition for the [BottomLeft] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-c.ivar
type IBottomLeft interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BottomLeft) Init() BottomLeft {
	rv := objc.Send[BottomLeft](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BottomLeft) Autorelease() BottomLeft {
	rv := objc.Send[BottomLeft](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBottomLeft creates a new BottomLeft instance.
func NewBottomLeft() BottomLeft {
	class := getBottomLeftClass()
	rv := objc.Send[BottomLeft](objc.ID(class.class), objc.Sel("new"))
	return rv
}

