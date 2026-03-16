// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_gelu_with_erf] class.
var (
	_EspressoPass_fuse_gelu_with_erfClass     EspressoPass_fuse_gelu_with_erfClass
	_EspressoPass_fuse_gelu_with_erfClassOnce sync.Once
)

func getEspressoPass_fuse_gelu_with_erfClass() EspressoPass_fuse_gelu_with_erfClass {
	_EspressoPass_fuse_gelu_with_erfClassOnce.Do(func() {
		_EspressoPass_fuse_gelu_with_erfClass = EspressoPass_fuse_gelu_with_erfClass{class: objc.GetClass("EspressoPass_fuse_gelu_with_erf")}
	})
	return _EspressoPass_fuse_gelu_with_erfClass
}

// GetEspressoPass_fuse_gelu_with_erfClass returns the class object for EspressoPass_fuse_gelu_with_erf.
func GetEspressoPass_fuse_gelu_with_erfClass() EspressoPass_fuse_gelu_with_erfClass {
	return getEspressoPass_fuse_gelu_with_erfClass()
}

type EspressoPass_fuse_gelu_with_erfClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_gelu_with_erfClass) Alloc() EspressoPass_fuse_gelu_with_erf {
	rv := objc.Send[EspressoPass_fuse_gelu_with_erf](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_gelu_with_erf
type EspressoPass_fuse_gelu_with_erf struct {
	EspressoCustomPass
}

// EspressoPass_fuse_gelu_with_erfFromID constructs a [EspressoPass_fuse_gelu_with_erf] from an objc.ID.
func EspressoPass_fuse_gelu_with_erfFromID(id objc.ID) EspressoPass_fuse_gelu_with_erf {
	return EspressoPass_fuse_gelu_with_erf{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_gelu_with_erf implements IEspressoPass_fuse_gelu_with_erf.
var _ IEspressoPass_fuse_gelu_with_erf = EspressoPass_fuse_gelu_with_erf{}

// An interface definition for the [EspressoPass_fuse_gelu_with_erf] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_gelu_with_erf
type IEspressoPass_fuse_gelu_with_erf interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_gelu_with_erf) Init() EspressoPass_fuse_gelu_with_erf {
	rv := objc.Send[EspressoPass_fuse_gelu_with_erf](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_gelu_with_erf) Autorelease() EspressoPass_fuse_gelu_with_erf {
	rv := objc.Send[EspressoPass_fuse_gelu_with_erf](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_gelu_with_erf creates a new EspressoPass_fuse_gelu_with_erf instance.
func NewEspressoPass_fuse_gelu_with_erf() EspressoPass_fuse_gelu_with_erf {
	class := getEspressoPass_fuse_gelu_with_erfClass()
	rv := objc.Send[EspressoPass_fuse_gelu_with_erf](objc.ID(class.class), objc.Sel("new"))
	return rv
}

