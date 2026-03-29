// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_l2_normalize_5d_to_instancenorm] class.
var (
	_EspressoPass_l2_normalize_5d_to_instancenormClass     EspressoPass_l2_normalize_5d_to_instancenormClass
	_EspressoPass_l2_normalize_5d_to_instancenormClassOnce sync.Once
)

func getEspressoPass_l2_normalize_5d_to_instancenormClass() EspressoPass_l2_normalize_5d_to_instancenormClass {
	_EspressoPass_l2_normalize_5d_to_instancenormClassOnce.Do(func() {
		_EspressoPass_l2_normalize_5d_to_instancenormClass = EspressoPass_l2_normalize_5d_to_instancenormClass{class: objc.GetClass("EspressoPass_l2_normalize_5d_to_instancenorm")}
	})
	return _EspressoPass_l2_normalize_5d_to_instancenormClass
}

// GetEspressoPass_l2_normalize_5d_to_instancenormClass returns the class object for EspressoPass_l2_normalize_5d_to_instancenorm.
func GetEspressoPass_l2_normalize_5d_to_instancenormClass() EspressoPass_l2_normalize_5d_to_instancenormClass {
	return getEspressoPass_l2_normalize_5d_to_instancenormClass()
}

type EspressoPass_l2_normalize_5d_to_instancenormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_l2_normalize_5d_to_instancenormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_l2_normalize_5d_to_instancenormClass) Alloc() EspressoPass_l2_normalize_5d_to_instancenorm {
	rv := objc.Send[EspressoPass_l2_normalize_5d_to_instancenorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_l2_normalize_5d_to_instancenorm
type EspressoPass_l2_normalize_5d_to_instancenorm struct {
	EspressoCustomPass
}

// EspressoPass_l2_normalize_5d_to_instancenormFromID constructs a [EspressoPass_l2_normalize_5d_to_instancenorm] from an objc.ID.
func EspressoPass_l2_normalize_5d_to_instancenormFromID(id objc.ID) EspressoPass_l2_normalize_5d_to_instancenorm {
	return EspressoPass_l2_normalize_5d_to_instancenorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_l2_normalize_5d_to_instancenorm implements IEspressoPass_l2_normalize_5d_to_instancenorm.
var _ IEspressoPass_l2_normalize_5d_to_instancenorm = EspressoPass_l2_normalize_5d_to_instancenorm{}

// An interface definition for the [EspressoPass_l2_normalize_5d_to_instancenorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_l2_normalize_5d_to_instancenorm
type IEspressoPass_l2_normalize_5d_to_instancenorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_l2_normalize_5d_to_instancenorm) Init() EspressoPass_l2_normalize_5d_to_instancenorm {
	rv := objc.Send[EspressoPass_l2_normalize_5d_to_instancenorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_l2_normalize_5d_to_instancenorm) Autorelease() EspressoPass_l2_normalize_5d_to_instancenorm {
	rv := objc.Send[EspressoPass_l2_normalize_5d_to_instancenorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_l2_normalize_5d_to_instancenorm creates a new EspressoPass_l2_normalize_5d_to_instancenorm instance.
func NewEspressoPass_l2_normalize_5d_to_instancenorm() EspressoPass_l2_normalize_5d_to_instancenorm {
	class := getEspressoPass_l2_normalize_5d_to_instancenormClass()
	rv := objc.Send[EspressoPass_l2_normalize_5d_to_instancenorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

