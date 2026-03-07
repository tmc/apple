// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_optimize_coreml_conv1d_lrelu] class.
var (
	_EspressoPass_optimize_coreml_conv1d_lreluClass     EspressoPass_optimize_coreml_conv1d_lreluClass
	_EspressoPass_optimize_coreml_conv1d_lreluClassOnce sync.Once
)

func getEspressoPass_optimize_coreml_conv1d_lreluClass() EspressoPass_optimize_coreml_conv1d_lreluClass {
	_EspressoPass_optimize_coreml_conv1d_lreluClassOnce.Do(func() {
		_EspressoPass_optimize_coreml_conv1d_lreluClass = EspressoPass_optimize_coreml_conv1d_lreluClass{class: objc.GetClass("EspressoPass_optimize_coreml_conv1d_lrelu")}
	})
	return _EspressoPass_optimize_coreml_conv1d_lreluClass
}

// GetEspressoPass_optimize_coreml_conv1d_lreluClass returns the class object for EspressoPass_optimize_coreml_conv1d_lrelu.
func GetEspressoPass_optimize_coreml_conv1d_lreluClass() EspressoPass_optimize_coreml_conv1d_lreluClass {
	return getEspressoPass_optimize_coreml_conv1d_lreluClass()
}

type EspressoPass_optimize_coreml_conv1d_lreluClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_optimize_coreml_conv1d_lreluClass) Alloc() EspressoPass_optimize_coreml_conv1d_lrelu {
	rv := objc.Send[EspressoPass_optimize_coreml_conv1d_lrelu](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_optimize_coreml_conv1d_lrelu
type EspressoPass_optimize_coreml_conv1d_lrelu struct {
	EspressoCustomPass
}

// EspressoPass_optimize_coreml_conv1d_lreluFromID constructs a [EspressoPass_optimize_coreml_conv1d_lrelu] from an objc.ID.
func EspressoPass_optimize_coreml_conv1d_lreluFromID(id objc.ID) EspressoPass_optimize_coreml_conv1d_lrelu {
	return EspressoPass_optimize_coreml_conv1d_lrelu{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_optimize_coreml_conv1d_lrelu implements IEspressoPass_optimize_coreml_conv1d_lrelu.
var _ IEspressoPass_optimize_coreml_conv1d_lrelu = EspressoPass_optimize_coreml_conv1d_lrelu{}





// An interface definition for the [EspressoPass_optimize_coreml_conv1d_lrelu] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_optimize_coreml_conv1d_lrelu
type IEspressoPass_optimize_coreml_conv1d_lrelu interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_optimize_coreml_conv1d_lrelu) Init() EspressoPass_optimize_coreml_conv1d_lrelu {
	rv := objc.Send[EspressoPass_optimize_coreml_conv1d_lrelu](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_optimize_coreml_conv1d_lrelu) Autorelease() EspressoPass_optimize_coreml_conv1d_lrelu {
	rv := objc.Send[EspressoPass_optimize_coreml_conv1d_lrelu](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_optimize_coreml_conv1d_lrelu creates a new EspressoPass_optimize_coreml_conv1d_lrelu instance.
func NewEspressoPass_optimize_coreml_conv1d_lrelu() EspressoPass_optimize_coreml_conv1d_lrelu {
	class := getEspressoPass_optimize_coreml_conv1d_lreluClass()
	rv := objc.Send[EspressoPass_optimize_coreml_conv1d_lrelu](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































