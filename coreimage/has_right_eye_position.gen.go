// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasRightEyePosition] class.
var (
	_HasRightEyePositionClass     HasRightEyePositionClass
	_HasRightEyePositionClassOnce sync.Once
)

func getHasRightEyePositionClass() HasRightEyePositionClass {
	_HasRightEyePositionClassOnce.Do(func() {
		_HasRightEyePositionClass = HasRightEyePositionClass{class: objc.GetClass("hasRightEyePosition")}
	})
	return _HasRightEyePositionClass
}

// GetHasRightEyePositionClass returns the class object for hasRightEyePosition.
func GetHasRightEyePositionClass() HasRightEyePositionClass {
	return getHasRightEyePositionClass()
}

type HasRightEyePositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HasRightEyePositionClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasRightEyePositionClass) Alloc() HasRightEyePosition {
	rv := objc.Send[HasRightEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-c.ivar
type HasRightEyePosition struct {
	objectivec.Object
}

// HasRightEyePositionFromID constructs a [HasRightEyePosition] from an objc.ID.
func HasRightEyePositionFromID(id objc.ID) HasRightEyePosition {
	return HasRightEyePosition{objectivec.Object{ID: id}}
}

// Ensure HasRightEyePosition implements IHasRightEyePosition.
var _ IHasRightEyePosition = HasRightEyePosition{}

// An interface definition for the [HasRightEyePosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-c.ivar
type IHasRightEyePosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasRightEyePosition) Init() HasRightEyePosition {
	rv := objc.Send[HasRightEyePosition](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasRightEyePosition) Autorelease() HasRightEyePosition {
	rv := objc.Send[HasRightEyePosition](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasRightEyePosition creates a new HasRightEyePosition instance.
func NewHasRightEyePosition() HasRightEyePosition {
	class := getHasRightEyePositionClass()
	rv := objc.Send[HasRightEyePosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}
