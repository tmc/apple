// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Bounds] class.
var (
	_BoundsClass     BoundsClass
	_BoundsClassOnce sync.Once
)

func getBoundsClass() BoundsClass {
	_BoundsClassOnce.Do(func() {
		_BoundsClass = BoundsClass{class: objc.GetClass("bounds")}
	})
	return _BoundsClass
}

// GetBoundsClass returns the class object for bounds.
func GetBoundsClass() BoundsClass {
	return getBoundsClass()
}

type BoundsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BoundsClass) Alloc() Bounds {
	rv := objc.Send[Bounds](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/bounds-c.ivar
type Bounds struct {
	objectivec.Object
}

// BoundsFromID constructs a [Bounds] from an objc.ID.
func BoundsFromID(id objc.ID) Bounds {
	return Bounds{objectivec.Object{ID: id}}
}
// Ensure Bounds implements IBounds.
var _ IBounds = Bounds{}

// An interface definition for the [Bounds] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/bounds-c.ivar
type IBounds interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b Bounds) Init() Bounds {
	rv := objc.Send[Bounds](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b Bounds) Autorelease() Bounds {
	rv := objc.Send[Bounds](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBounds creates a new Bounds instance.
func NewBounds() Bounds {
	class := getBoundsClass()
	rv := objc.Send[Bounds](objc.ID(class.class), objc.Sel("new"))
	return rv
}

