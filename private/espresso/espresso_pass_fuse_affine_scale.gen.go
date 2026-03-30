// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_affine_scale] class.
var (
	_EspressoPass_fuse_affine_scaleClass     EspressoPass_fuse_affine_scaleClass
	_EspressoPass_fuse_affine_scaleClassOnce sync.Once
)

func getEspressoPass_fuse_affine_scaleClass() EspressoPass_fuse_affine_scaleClass {
	_EspressoPass_fuse_affine_scaleClassOnce.Do(func() {
		_EspressoPass_fuse_affine_scaleClass = EspressoPass_fuse_affine_scaleClass{class: objc.GetClass("EspressoPass_fuse_affine_scale")}
	})
	return _EspressoPass_fuse_affine_scaleClass
}

// GetEspressoPass_fuse_affine_scaleClass returns the class object for EspressoPass_fuse_affine_scale.
func GetEspressoPass_fuse_affine_scaleClass() EspressoPass_fuse_affine_scaleClass {
	return getEspressoPass_fuse_affine_scaleClass()
}

type EspressoPass_fuse_affine_scaleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fuse_affine_scaleClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_affine_scaleClass) Alloc() EspressoPass_fuse_affine_scale {
	rv := objc.Send[EspressoPass_fuse_affine_scale](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_affine_scale
type EspressoPass_fuse_affine_scale struct {
	EspressoCustomPass
}

// EspressoPass_fuse_affine_scaleFromID constructs a [EspressoPass_fuse_affine_scale] from an objc.ID.
func EspressoPass_fuse_affine_scaleFromID(id objc.ID) EspressoPass_fuse_affine_scale {
	return EspressoPass_fuse_affine_scale{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_fuse_affine_scale implements IEspressoPass_fuse_affine_scale.
var _ IEspressoPass_fuse_affine_scale = EspressoPass_fuse_affine_scale{}

// An interface definition for the [EspressoPass_fuse_affine_scale] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_affine_scale
type IEspressoPass_fuse_affine_scale interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_affine_scale) Init() EspressoPass_fuse_affine_scale {
	rv := objc.Send[EspressoPass_fuse_affine_scale](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_affine_scale) Autorelease() EspressoPass_fuse_affine_scale {
	rv := objc.Send[EspressoPass_fuse_affine_scale](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_affine_scale creates a new EspressoPass_fuse_affine_scale instance.
func NewEspressoPass_fuse_affine_scale() EspressoPass_fuse_affine_scale {
	class := getEspressoPass_fuse_affine_scaleClass()
	rv := objc.Send[EspressoPass_fuse_affine_scale](objc.ID(class.class), objc.Sel("new"))
	return rv
}
