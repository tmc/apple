// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RightEyePosition] class.
var (
	_RightEyePositionClass     RightEyePositionClass
	_RightEyePositionClassOnce sync.Once
)

func getRightEyePositionClass() RightEyePositionClass {
	_RightEyePositionClassOnce.Do(func() {
		_RightEyePositionClass = RightEyePositionClass{class: objc.GetClass("rightEyePosition")}
	})
	return _RightEyePositionClass
}

// GetRightEyePositionClass returns the class object for rightEyePosition.
func GetRightEyePositionClass() RightEyePositionClass {
	return getRightEyePositionClass()
}

type RightEyePositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RightEyePositionClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RightEyePositionClass) Alloc() RightEyePosition {
	rv := objc.Send[RightEyePosition](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-c.ivar
type RightEyePosition struct {
	objectivec.Object
}

// RightEyePositionFromID constructs a [RightEyePosition] from an objc.ID.
func RightEyePositionFromID(id objc.ID) RightEyePosition {
	return RightEyePosition{objectivec.Object{ID: id}}
}
// Ensure RightEyePosition implements IRightEyePosition.
var _ IRightEyePosition = RightEyePosition{}

// An interface definition for the [RightEyePosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-c.ivar
type IRightEyePosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RightEyePosition) Init() RightEyePosition {
	rv := objc.Send[RightEyePosition](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RightEyePosition) Autorelease() RightEyePosition {
	rv := objc.Send[RightEyePosition](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRightEyePosition creates a new RightEyePosition instance.
func NewRightEyePosition() RightEyePosition {
	class := getRightEyePositionClass()
	rv := objc.Send[RightEyePosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

