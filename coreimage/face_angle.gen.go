// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FaceAngle] class.
var (
	_FaceAngleClass     FaceAngleClass
	_FaceAngleClassOnce sync.Once
)

func getFaceAngleClass() FaceAngleClass {
	_FaceAngleClassOnce.Do(func() {
		_FaceAngleClass = FaceAngleClass{class: objc.GetClass("faceAngle")}
	})
	return _FaceAngleClass
}

// GetFaceAngleClass returns the class object for faceAngle.
func GetFaceAngleClass() FaceAngleClass {
	return getFaceAngleClass()
}

type FaceAngleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (fc FaceAngleClass) Alloc() FaceAngle {
	rv := objc.Send[FaceAngle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-c.ivar
type FaceAngle struct {
	objectivec.Object
}

// FaceAngleFromID constructs a [FaceAngle] from an objc.ID.
func FaceAngleFromID(id objc.ID) FaceAngle {
	return FaceAngle{objectivec.Object{ID: id}}
}
// Ensure FaceAngle implements IFaceAngle.
var _ IFaceAngle = FaceAngle{}

// An interface definition for the [FaceAngle] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-c.ivar
type IFaceAngle interface {
	objectivec.IObject
}

// Init initializes the instance.
func (f FaceAngle) Init() FaceAngle {
	rv := objc.Send[FaceAngle](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FaceAngle) Autorelease() FaceAngle {
	rv := objc.Send[FaceAngle](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceAngle creates a new FaceAngle instance.
func NewFaceAngle() FaceAngle {
	class := getFaceAngleClass()
	rv := objc.Send[FaceAngle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

