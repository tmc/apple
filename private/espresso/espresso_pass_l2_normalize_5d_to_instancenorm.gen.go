// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassL2Normalize5dToInstancenorm] class.
var (
	_EspressoPassL2Normalize5dToInstancenormClass     EspressoPassL2Normalize5dToInstancenormClass
	_EspressoPassL2Normalize5dToInstancenormClassOnce sync.Once
)

func getEspressoPassL2Normalize5dToInstancenormClass() EspressoPassL2Normalize5dToInstancenormClass {
	_EspressoPassL2Normalize5dToInstancenormClassOnce.Do(func() {
		_EspressoPassL2Normalize5dToInstancenormClass = EspressoPassL2Normalize5dToInstancenormClass{class: objc.GetClass("EspressoPass_l2_normalize_5d_to_instancenorm")}
	})
	return _EspressoPassL2Normalize5dToInstancenormClass
}

// GetEspressoPassL2Normalize5dToInstancenormClass returns the class object for EspressoPass_l2_normalize_5d_to_instancenorm.
func GetEspressoPassL2Normalize5dToInstancenormClass() EspressoPassL2Normalize5dToInstancenormClass {
	return getEspressoPassL2Normalize5dToInstancenormClass()
}

type EspressoPassL2Normalize5dToInstancenormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassL2Normalize5dToInstancenormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassL2Normalize5dToInstancenormClass) Alloc() EspressoPassL2Normalize5dToInstancenorm {
	rv := objc.Send[EspressoPassL2Normalize5dToInstancenorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_l2_normalize_5d_to_instancenorm
type EspressoPassL2Normalize5dToInstancenorm struct {
	EspressoCustomPass
}

// EspressoPassL2Normalize5dToInstancenormFromID constructs a [EspressoPassL2Normalize5dToInstancenorm] from an objc.ID.
func EspressoPassL2Normalize5dToInstancenormFromID(id objc.ID) EspressoPassL2Normalize5dToInstancenorm {
	return EspressoPassL2Normalize5dToInstancenorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_l2_normalize_5d_to_instancenormFromID is an alias for [EspressoPassL2Normalize5dToInstancenormFromID] for cross-framework compatibility.
func EspressoPass_l2_normalize_5d_to_instancenormFromID(id objc.ID) EspressoPassL2Normalize5dToInstancenorm {
	return EspressoPassL2Normalize5dToInstancenormFromID(id)
}

// Ensure EspressoPassL2Normalize5dToInstancenorm implements IEspressoPassL2Normalize5dToInstancenorm.
var _ IEspressoPassL2Normalize5dToInstancenorm = EspressoPassL2Normalize5dToInstancenorm{}

// An interface definition for the [EspressoPassL2Normalize5dToInstancenorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_l2_normalize_5d_to_instancenorm
type IEspressoPassL2Normalize5dToInstancenorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassL2Normalize5dToInstancenorm) Init() EspressoPassL2Normalize5dToInstancenorm {
	rv := objc.Send[EspressoPassL2Normalize5dToInstancenorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassL2Normalize5dToInstancenorm) Autorelease() EspressoPassL2Normalize5dToInstancenorm {
	rv := objc.Send[EspressoPassL2Normalize5dToInstancenorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassL2Normalize5dToInstancenorm creates a new EspressoPassL2Normalize5dToInstancenorm instance.
func NewEspressoPassL2Normalize5dToInstancenorm() EspressoPassL2Normalize5dToInstancenorm {
	class := getEspressoPassL2Normalize5dToInstancenormClass()
	rv := objc.Send[EspressoPassL2Normalize5dToInstancenorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}
