// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseAffineScale] class.
var (
	_EspressoPassFuseAffineScaleClass     EspressoPassFuseAffineScaleClass
	_EspressoPassFuseAffineScaleClassOnce sync.Once
)

func getEspressoPassFuseAffineScaleClass() EspressoPassFuseAffineScaleClass {
	_EspressoPassFuseAffineScaleClassOnce.Do(func() {
		_EspressoPassFuseAffineScaleClass = EspressoPassFuseAffineScaleClass{class: objc.GetClass("EspressoPass_fuse_affine_scale")}
	})
	return _EspressoPassFuseAffineScaleClass
}

// GetEspressoPassFuseAffineScaleClass returns the class object for EspressoPass_fuse_affine_scale.
func GetEspressoPassFuseAffineScaleClass() EspressoPassFuseAffineScaleClass {
	return getEspressoPassFuseAffineScaleClass()
}

type EspressoPassFuseAffineScaleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseAffineScaleClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseAffineScaleClass) Alloc() EspressoPassFuseAffineScale {
	rv := objc.SendIfResponds[EspressoPassFuseAffineScale](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseAffineScale struct {
	EspressoCustomPass
}

// EspressoPassFuseAffineScaleFromID constructs a [EspressoPassFuseAffineScale] from an objc.ID.
func EspressoPassFuseAffineScaleFromID(id objc.ID) EspressoPassFuseAffineScale {
	return EspressoPassFuseAffineScale{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_affine_scaleFromID is an alias for [EspressoPassFuseAffineScaleFromID] for cross-framework compatibility.
func EspressoPass_fuse_affine_scaleFromID(id objc.ID) EspressoPassFuseAffineScale {
	return EspressoPassFuseAffineScaleFromID(id)
}

// Ensure EspressoPassFuseAffineScale implements IEspressoPassFuseAffineScale.
var _ IEspressoPassFuseAffineScale = EspressoPassFuseAffineScale{}

// An interface definition for the [EspressoPassFuseAffineScale] class.
type IEspressoPassFuseAffineScale interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseAffineScale) Init() EspressoPassFuseAffineScale {
	rv := objc.SendIfResponds[EspressoPassFuseAffineScale](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseAffineScale) Autorelease() EspressoPassFuseAffineScale {
	rv := objc.SendIfResponds[EspressoPassFuseAffineScale](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseAffineScale creates a new EspressoPassFuseAffineScale instance.
func NewEspressoPassFuseAffineScale() EspressoPassFuseAffineScale {
	class := getEspressoPassFuseAffineScaleClass()
	rv := objc.SendIfResponds[EspressoPassFuseAffineScale](objc.ID(class.class), objc.Sel("new"))
	return rv
}
