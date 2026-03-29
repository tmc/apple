// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LeftEyePosition] class.
var (
	_LeftEyePositionClass     LeftEyePositionClass
	_LeftEyePositionClassOnce sync.Once
)

func getLeftEyePositionClass() LeftEyePositionClass {
	_LeftEyePositionClassOnce.Do(func() {
		_LeftEyePositionClass = LeftEyePositionClass{class: objc.GetClass("leftEyePosition")}
	})
	return _LeftEyePositionClass
}

// GetLeftEyePositionClass returns the class object for leftEyePosition.
func GetLeftEyePositionClass() LeftEyePositionClass {
	return getLeftEyePositionClass()
}

type LeftEyePositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LeftEyePositionClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LeftEyePositionClass) Alloc() LeftEyePosition {
	rv := objc.Send[LeftEyePosition](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-c.ivar
type LeftEyePosition struct {
	objectivec.Object
}

// LeftEyePositionFromID constructs a [LeftEyePosition] from an objc.ID.
func LeftEyePositionFromID(id objc.ID) LeftEyePosition {
	return LeftEyePosition{objectivec.Object{ID: id}}
}
// Ensure LeftEyePosition implements ILeftEyePosition.
var _ ILeftEyePosition = LeftEyePosition{}

// An interface definition for the [LeftEyePosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-c.ivar
type ILeftEyePosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l LeftEyePosition) Init() LeftEyePosition {
	rv := objc.Send[LeftEyePosition](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LeftEyePosition) Autorelease() LeftEyePosition {
	rv := objc.Send[LeftEyePosition](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeftEyePosition creates a new LeftEyePosition instance.
func NewLeftEyePosition() LeftEyePosition {
	class := getLeftEyePositionClass()
	rv := objc.Send[LeftEyePosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

