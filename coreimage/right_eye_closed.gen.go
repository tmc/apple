// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RightEyeClosed] class.
var (
	_RightEyeClosedClass     RightEyeClosedClass
	_RightEyeClosedClassOnce sync.Once
)

func getRightEyeClosedClass() RightEyeClosedClass {
	_RightEyeClosedClassOnce.Do(func() {
		_RightEyeClosedClass = RightEyeClosedClass{class: objc.GetClass("rightEyeClosed")}
	})
	return _RightEyeClosedClass
}

// GetRightEyeClosedClass returns the class object for rightEyeClosed.
func GetRightEyeClosedClass() RightEyeClosedClass {
	return getRightEyeClosedClass()
}

type RightEyeClosedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RightEyeClosedClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RightEyeClosedClass) Alloc() RightEyeClosed {
	rv := objc.Send[RightEyeClosed](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-c.ivar
type RightEyeClosed struct {
	objectivec.Object
}

// RightEyeClosedFromID constructs a [RightEyeClosed] from an objc.ID.
func RightEyeClosedFromID(id objc.ID) RightEyeClosed {
	return RightEyeClosed{objectivec.Object{ID: id}}
}

// Ensure RightEyeClosed implements IRightEyeClosed.
var _ IRightEyeClosed = RightEyeClosed{}

// An interface definition for the [RightEyeClosed] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-c.ivar
type IRightEyeClosed interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RightEyeClosed) Init() RightEyeClosed {
	rv := objc.Send[RightEyeClosed](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RightEyeClosed) Autorelease() RightEyeClosed {
	rv := objc.Send[RightEyeClosed](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRightEyeClosed creates a new RightEyeClosed instance.
func NewRightEyeClosed() RightEyeClosed {
	class := getRightEyeClosedClass()
	rv := objc.Send[RightEyeClosed](objc.ID(class.class), objc.Sel("new"))
	return rv
}
