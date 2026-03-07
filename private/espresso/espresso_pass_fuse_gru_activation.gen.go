// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_gru_activation] class.
var (
	_EspressoPass_fuse_gru_activationClass     EspressoPass_fuse_gru_activationClass
	_EspressoPass_fuse_gru_activationClassOnce sync.Once
)

func getEspressoPass_fuse_gru_activationClass() EspressoPass_fuse_gru_activationClass {
	_EspressoPass_fuse_gru_activationClassOnce.Do(func() {
		_EspressoPass_fuse_gru_activationClass = EspressoPass_fuse_gru_activationClass{class: objc.GetClass("EspressoPass_fuse_gru_activation")}
	})
	return _EspressoPass_fuse_gru_activationClass
}

// GetEspressoPass_fuse_gru_activationClass returns the class object for EspressoPass_fuse_gru_activation.
func GetEspressoPass_fuse_gru_activationClass() EspressoPass_fuse_gru_activationClass {
	return getEspressoPass_fuse_gru_activationClass()
}

type EspressoPass_fuse_gru_activationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_gru_activationClass) Alloc() EspressoPass_fuse_gru_activation {
	rv := objc.Send[EspressoPass_fuse_gru_activation](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_gru_activation
type EspressoPass_fuse_gru_activation struct {
	EspressoCustomPass
}

// EspressoPass_fuse_gru_activationFromID constructs a [EspressoPass_fuse_gru_activation] from an objc.ID.
func EspressoPass_fuse_gru_activationFromID(id objc.ID) EspressoPass_fuse_gru_activation {
	return EspressoPass_fuse_gru_activation{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_gru_activation implements IEspressoPass_fuse_gru_activation.
var _ IEspressoPass_fuse_gru_activation = EspressoPass_fuse_gru_activation{}





// An interface definition for the [EspressoPass_fuse_gru_activation] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_gru_activation
type IEspressoPass_fuse_gru_activation interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_fuse_gru_activation) Init() EspressoPass_fuse_gru_activation {
	rv := objc.Send[EspressoPass_fuse_gru_activation](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_gru_activation) Autorelease() EspressoPass_fuse_gru_activation {
	rv := objc.Send[EspressoPass_fuse_gru_activation](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_gru_activation creates a new EspressoPass_fuse_gru_activation instance.
func NewEspressoPass_fuse_gru_activation() EspressoPass_fuse_gru_activation {
	class := getEspressoPass_fuse_gru_activationClass()
	rv := objc.Send[EspressoPass_fuse_gru_activation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































