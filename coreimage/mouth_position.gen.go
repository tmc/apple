// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MouthPosition] class.
var (
	_MouthPositionClass     MouthPositionClass
	_MouthPositionClassOnce sync.Once
)

func getMouthPositionClass() MouthPositionClass {
	_MouthPositionClassOnce.Do(func() {
		_MouthPositionClass = MouthPositionClass{class: objc.GetClass("mouthPosition")}
	})
	return _MouthPositionClass
}

// GetMouthPositionClass returns the class object for mouthPosition.
func GetMouthPositionClass() MouthPositionClass {
	return getMouthPositionClass()
}

type MouthPositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MouthPositionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MouthPositionClass) Alloc() MouthPosition {
	rv := objc.Send[MouthPosition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-c.ivar
type MouthPosition struct {
	objectivec.Object
}

// MouthPositionFromID constructs a [MouthPosition] from an objc.ID.
func MouthPositionFromID(id objc.ID) MouthPosition {
	return MouthPosition{objectivec.Object{ID: id}}
}

// Ensure MouthPosition implements IMouthPosition.
var _ IMouthPosition = MouthPosition{}

// An interface definition for the [MouthPosition] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-c.ivar
type IMouthPosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MouthPosition) Init() MouthPosition {
	rv := objc.Send[MouthPosition](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MouthPosition) Autorelease() MouthPosition {
	rv := objc.Send[MouthPosition](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMouthPosition creates a new MouthPosition instance.
func NewMouthPosition() MouthPosition {
	class := getMouthPositionClass()
	rv := objc.Send[MouthPosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}
