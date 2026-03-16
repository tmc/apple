// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze] class.
var (
	_EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass     EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass
	_EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClassOnce sync.Once
)

func getEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass {
	_EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClassOnce.Do(func() {
		_EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass = EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass{class: objc.GetClass("EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze")}
	})
	return _EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass
}

// GetEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass returns the class object for EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze.
func GetEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass {
	return getEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass()
}

type EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass) Alloc() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze {
	rv := objc.Send[EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze
type EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeFromID constructs a [EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze] from an objc.ID.
func EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeFromID(id objc.ID) EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze {
	return EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze implements IEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze.
var _ IEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze = EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze{}

// An interface definition for the [EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze
type IEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze) Init() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze {
	rv := objc.Send[EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze) Autorelease() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze {
	rv := objc.Send[EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze creates a new EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze instance.
func NewEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze() EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze {
	class := getEspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeClass()
	rv := objc.Send[EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze](objc.ID(class.class), objc.Sel("new"))
	return rv
}

