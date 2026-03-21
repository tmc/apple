// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasLeftEyePosition] class.
var (
	_HasLeftEyePositionClass     HasLeftEyePositionClass
	_HasLeftEyePositionClassOnce sync.Once
)

func getHasLeftEyePositionClass() HasLeftEyePositionClass {
	_HasLeftEyePositionClassOnce.Do(func() {
		_HasLeftEyePositionClass = HasLeftEyePositionClass{class: objc.GetClass("hasLeftEyePosition")}
	})
	return _HasLeftEyePositionClass
}

// GetHasLeftEyePositionClass returns the class object for hasLeftEyePosition.
func GetHasLeftEyePositionClass() HasLeftEyePositionClass {
	return getHasLeftEyePositionClass()
}

type HasLeftEyePositionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasLeftEyePositionClass) Alloc() HasLeftEyePosition {
	rv := objc.Send[HasLeftEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-c.ivar
type HasLeftEyePosition struct {
	objectivec.Object
}

// HasLeftEyePositionFromID constructs a [HasLeftEyePosition] from an objc.ID.
func HasLeftEyePositionFromID(id objc.ID) HasLeftEyePosition {
	return HasLeftEyePosition{objectivec.Object{ID: id}}
}
// Ensure HasLeftEyePosition implements IHasLeftEyePosition.
var _ IHasLeftEyePosition = HasLeftEyePosition{}

// An interface definition for the [HasLeftEyePosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-c.ivar
type IHasLeftEyePosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasLeftEyePosition) Init() HasLeftEyePosition {
	rv := objc.Send[HasLeftEyePosition](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasLeftEyePosition) Autorelease() HasLeftEyePosition {
	rv := objc.Send[HasLeftEyePosition](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasLeftEyePosition creates a new HasLeftEyePosition instance.
func NewHasLeftEyePosition() HasLeftEyePosition {
	class := getHasLeftEyePositionClass()
	rv := objc.Send[HasLeftEyePosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

