// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_concat_fuse_fix] class.
var (
	_EspressoPass_concat_fuse_fixClass     EspressoPass_concat_fuse_fixClass
	_EspressoPass_concat_fuse_fixClassOnce sync.Once
)

func getEspressoPass_concat_fuse_fixClass() EspressoPass_concat_fuse_fixClass {
	_EspressoPass_concat_fuse_fixClassOnce.Do(func() {
		_EspressoPass_concat_fuse_fixClass = EspressoPass_concat_fuse_fixClass{class: objc.GetClass("EspressoPass_concat_fuse_fix")}
	})
	return _EspressoPass_concat_fuse_fixClass
}

// GetEspressoPass_concat_fuse_fixClass returns the class object for EspressoPass_concat_fuse_fix.
func GetEspressoPass_concat_fuse_fixClass() EspressoPass_concat_fuse_fixClass {
	return getEspressoPass_concat_fuse_fixClass()
}

type EspressoPass_concat_fuse_fixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_concat_fuse_fixClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_concat_fuse_fixClass) Alloc() EspressoPass_concat_fuse_fix {
	rv := objc.Send[EspressoPass_concat_fuse_fix](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_concat_fuse_fix
type EspressoPass_concat_fuse_fix struct {
	EspressoCustomPass
}

// EspressoPass_concat_fuse_fixFromID constructs a [EspressoPass_concat_fuse_fix] from an objc.ID.
func EspressoPass_concat_fuse_fixFromID(id objc.ID) EspressoPass_concat_fuse_fix {
	return EspressoPass_concat_fuse_fix{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_concat_fuse_fix implements IEspressoPass_concat_fuse_fix.
var _ IEspressoPass_concat_fuse_fix = EspressoPass_concat_fuse_fix{}

// An interface definition for the [EspressoPass_concat_fuse_fix] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_concat_fuse_fix
type IEspressoPass_concat_fuse_fix interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_concat_fuse_fix) Init() EspressoPass_concat_fuse_fix {
	rv := objc.Send[EspressoPass_concat_fuse_fix](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_concat_fuse_fix) Autorelease() EspressoPass_concat_fuse_fix {
	rv := objc.Send[EspressoPass_concat_fuse_fix](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_concat_fuse_fix creates a new EspressoPass_concat_fuse_fix instance.
func NewEspressoPass_concat_fuse_fix() EspressoPass_concat_fuse_fix {
	class := getEspressoPass_concat_fuse_fixClass()
	rv := objc.Send[EspressoPass_concat_fuse_fix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

