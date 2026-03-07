// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_transpose_inner_product_weights] class.
var (
	_EspressoPass_transpose_inner_product_weightsClass     EspressoPass_transpose_inner_product_weightsClass
	_EspressoPass_transpose_inner_product_weightsClassOnce sync.Once
)

func getEspressoPass_transpose_inner_product_weightsClass() EspressoPass_transpose_inner_product_weightsClass {
	_EspressoPass_transpose_inner_product_weightsClassOnce.Do(func() {
		_EspressoPass_transpose_inner_product_weightsClass = EspressoPass_transpose_inner_product_weightsClass{class: objc.GetClass("EspressoPass_transpose_inner_product_weights")}
	})
	return _EspressoPass_transpose_inner_product_weightsClass
}

// GetEspressoPass_transpose_inner_product_weightsClass returns the class object for EspressoPass_transpose_inner_product_weights.
func GetEspressoPass_transpose_inner_product_weightsClass() EspressoPass_transpose_inner_product_weightsClass {
	return getEspressoPass_transpose_inner_product_weightsClass()
}

type EspressoPass_transpose_inner_product_weightsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_transpose_inner_product_weightsClass) Alloc() EspressoPass_transpose_inner_product_weights {
	rv := objc.Send[EspressoPass_transpose_inner_product_weights](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transpose_inner_product_weights
type EspressoPass_transpose_inner_product_weights struct {
	EspressoCustomPass
}

// EspressoPass_transpose_inner_product_weightsFromID constructs a [EspressoPass_transpose_inner_product_weights] from an objc.ID.
func EspressoPass_transpose_inner_product_weightsFromID(id objc.ID) EspressoPass_transpose_inner_product_weights {
	return EspressoPass_transpose_inner_product_weights{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_transpose_inner_product_weights implements IEspressoPass_transpose_inner_product_weights.
var _ IEspressoPass_transpose_inner_product_weights = EspressoPass_transpose_inner_product_weights{}





// An interface definition for the [EspressoPass_transpose_inner_product_weights] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transpose_inner_product_weights
type IEspressoPass_transpose_inner_product_weights interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_transpose_inner_product_weights) Init() EspressoPass_transpose_inner_product_weights {
	rv := objc.Send[EspressoPass_transpose_inner_product_weights](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_transpose_inner_product_weights) Autorelease() EspressoPass_transpose_inner_product_weights {
	rv := objc.Send[EspressoPass_transpose_inner_product_weights](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_transpose_inner_product_weights creates a new EspressoPass_transpose_inner_product_weights instance.
func NewEspressoPass_transpose_inner_product_weights() EspressoPass_transpose_inner_product_weights {
	class := getEspressoPass_transpose_inner_product_weightsClass()
	rv := objc.Send[EspressoPass_transpose_inner_product_weights](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































