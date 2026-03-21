// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BottomRight] class.
var (
	_BottomRightClass     BottomRightClass
	_BottomRightClassOnce sync.Once
)

func getBottomRightClass() BottomRightClass {
	_BottomRightClassOnce.Do(func() {
		_BottomRightClass = BottomRightClass{class: objc.GetClass("bottomRight")}
	})
	return _BottomRightClass
}

// GetBottomRightClass returns the class object for bottomRight.
func GetBottomRightClass() BottomRightClass {
	return getBottomRightClass()
}

type BottomRightClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BottomRightClass) Alloc() BottomRight {
	rv := objc.Send[BottomRight](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomRight-c.ivar
type BottomRight struct {
	objectivec.Object
}

// BottomRightFromID constructs a [BottomRight] from an objc.ID.
func BottomRightFromID(id objc.ID) BottomRight {
	return BottomRight{objectivec.Object{ID: id}}
}
// Ensure BottomRight implements IBottomRight.
var _ IBottomRight = BottomRight{}

// An interface definition for the [BottomRight] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomRight-c.ivar
type IBottomRight interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BottomRight) Init() BottomRight {
	rv := objc.Send[BottomRight](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BottomRight) Autorelease() BottomRight {
	rv := objc.Send[BottomRight](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBottomRight creates a new BottomRight instance.
func NewBottomRight() BottomRight {
	class := getBottomRightClass()
	rv := objc.Send[BottomRight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

