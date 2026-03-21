// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Vec] class.
var (
	_VecClass     VecClass
	_VecClassOnce sync.Once
)

func getVecClass() VecClass {
	_VecClassOnce.Do(func() {
		_VecClass = VecClass{class: objc.GetClass("vec")}
	})
	return _VecClass
}

// GetVecClass returns the class object for vec.
func GetVecClass() VecClass {
	return getVecClass()
}

type VecClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VecClass) Alloc() Vec {
	rv := objc.Send[Vec](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/vec
type Vec struct {
	objectivec.Object
}

// VecFromID constructs a [Vec] from an objc.ID.
func VecFromID(id objc.ID) Vec {
	return Vec{objectivec.Object{ID: id}}
}
// Ensure Vec implements IVec.
var _ IVec = Vec{}

// An interface definition for the [Vec] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/vec
type IVec interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v Vec) Init() Vec {
	rv := objc.Send[Vec](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v Vec) Autorelease() Vec {
	rv := objc.Send[Vec](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVec creates a new Vec instance.
func NewVec() Vec {
	class := getVecClass()
	rv := objc.Send[Vec](objc.ID(class.class), objc.Sel("new"))
	return rv
}

