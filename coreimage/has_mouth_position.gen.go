// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasMouthPosition] class.
var (
	_HasMouthPositionClass     HasMouthPositionClass
	_HasMouthPositionClassOnce sync.Once
)

func getHasMouthPositionClass() HasMouthPositionClass {
	_HasMouthPositionClassOnce.Do(func() {
		_HasMouthPositionClass = HasMouthPositionClass{class: objc.GetClass("hasMouthPosition")}
	})
	return _HasMouthPositionClass
}

// GetHasMouthPositionClass returns the class object for hasMouthPosition.
func GetHasMouthPositionClass() HasMouthPositionClass {
	return getHasMouthPositionClass()
}

type HasMouthPositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HasMouthPositionClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasMouthPositionClass) Alloc() HasMouthPosition {
	rv := objc.Send[HasMouthPosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-c.ivar
type HasMouthPosition struct {
	objectivec.Object
}

// HasMouthPositionFromID constructs a [HasMouthPosition] from an objc.ID.
func HasMouthPositionFromID(id objc.ID) HasMouthPosition {
	return HasMouthPosition{objectivec.Object{ID: id}}
}

// Ensure HasMouthPosition implements IHasMouthPosition.
var _ IHasMouthPosition = HasMouthPosition{}

// An interface definition for the [HasMouthPosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-c.ivar
type IHasMouthPosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasMouthPosition) Init() HasMouthPosition {
	rv := objc.Send[HasMouthPosition](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasMouthPosition) Autorelease() HasMouthPosition {
	rv := objc.Send[HasMouthPosition](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasMouthPosition creates a new HasMouthPosition instance.
func NewHasMouthPosition() HasMouthPosition {
	class := getHasMouthPositionClass()
	rv := objc.Send[HasMouthPosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}
