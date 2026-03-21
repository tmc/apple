// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasFaceAngle] class.
var (
	_HasFaceAngleClass     HasFaceAngleClass
	_HasFaceAngleClassOnce sync.Once
)

func getHasFaceAngleClass() HasFaceAngleClass {
	_HasFaceAngleClassOnce.Do(func() {
		_HasFaceAngleClass = HasFaceAngleClass{class: objc.GetClass("hasFaceAngle")}
	})
	return _HasFaceAngleClass
}

// GetHasFaceAngleClass returns the class object for hasFaceAngle.
func GetHasFaceAngleClass() HasFaceAngleClass {
	return getHasFaceAngleClass()
}

type HasFaceAngleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasFaceAngleClass) Alloc() HasFaceAngle {
	rv := objc.Send[HasFaceAngle](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-c.ivar
type HasFaceAngle struct {
	objectivec.Object
}

// HasFaceAngleFromID constructs a [HasFaceAngle] from an objc.ID.
func HasFaceAngleFromID(id objc.ID) HasFaceAngle {
	return HasFaceAngle{objectivec.Object{ID: id}}
}
// Ensure HasFaceAngle implements IHasFaceAngle.
var _ IHasFaceAngle = HasFaceAngle{}

// An interface definition for the [HasFaceAngle] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-c.ivar
type IHasFaceAngle interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasFaceAngle) Init() HasFaceAngle {
	rv := objc.Send[HasFaceAngle](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasFaceAngle) Autorelease() HasFaceAngle {
	rv := objc.Send[HasFaceAngle](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasFaceAngle creates a new HasFaceAngle instance.
func NewHasFaceAngle() HasFaceAngle {
	class := getHasFaceAngleClass()
	rv := objc.Send[HasFaceAngle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

