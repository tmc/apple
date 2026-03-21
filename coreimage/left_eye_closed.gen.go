// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LeftEyeClosed] class.
var (
	_LeftEyeClosedClass     LeftEyeClosedClass
	_LeftEyeClosedClassOnce sync.Once
)

func getLeftEyeClosedClass() LeftEyeClosedClass {
	_LeftEyeClosedClassOnce.Do(func() {
		_LeftEyeClosedClass = LeftEyeClosedClass{class: objc.GetClass("leftEyeClosed")}
	})
	return _LeftEyeClosedClass
}

// GetLeftEyeClosedClass returns the class object for leftEyeClosed.
func GetLeftEyeClosedClass() LeftEyeClosedClass {
	return getLeftEyeClosedClass()
}

type LeftEyeClosedClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (lc LeftEyeClosedClass) Alloc() LeftEyeClosed {
	rv := objc.Send[LeftEyeClosed](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-c.ivar
type LeftEyeClosed struct {
	objectivec.Object
}

// LeftEyeClosedFromID constructs a [LeftEyeClosed] from an objc.ID.
func LeftEyeClosedFromID(id objc.ID) LeftEyeClosed {
	return LeftEyeClosed{objectivec.Object{ID: id}}
}
// Ensure LeftEyeClosed implements ILeftEyeClosed.
var _ ILeftEyeClosed = LeftEyeClosed{}

// An interface definition for the [LeftEyeClosed] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-c.ivar
type ILeftEyeClosed interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l LeftEyeClosed) Init() LeftEyeClosed {
	rv := objc.Send[LeftEyeClosed](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LeftEyeClosed) Autorelease() LeftEyeClosed {
	rv := objc.Send[LeftEyeClosed](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeftEyeClosed creates a new LeftEyeClosed instance.
func NewLeftEyeClosed() LeftEyeClosed {
	class := getLeftEyeClosedClass()
	rv := objc.Send[LeftEyeClosed](objc.ID(class.class), objc.Sel("new"))
	return rv
}

