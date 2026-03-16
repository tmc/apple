// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_reshape_around_rank2_softmax] class.
var (
	_EspressoPass_remove_reshape_around_rank2_softmaxClass     EspressoPass_remove_reshape_around_rank2_softmaxClass
	_EspressoPass_remove_reshape_around_rank2_softmaxClassOnce sync.Once
)

func getEspressoPass_remove_reshape_around_rank2_softmaxClass() EspressoPass_remove_reshape_around_rank2_softmaxClass {
	_EspressoPass_remove_reshape_around_rank2_softmaxClassOnce.Do(func() {
		_EspressoPass_remove_reshape_around_rank2_softmaxClass = EspressoPass_remove_reshape_around_rank2_softmaxClass{class: objc.GetClass("EspressoPass_remove_reshape_around_rank2_softmax")}
	})
	return _EspressoPass_remove_reshape_around_rank2_softmaxClass
}

// GetEspressoPass_remove_reshape_around_rank2_softmaxClass returns the class object for EspressoPass_remove_reshape_around_rank2_softmax.
func GetEspressoPass_remove_reshape_around_rank2_softmaxClass() EspressoPass_remove_reshape_around_rank2_softmaxClass {
	return getEspressoPass_remove_reshape_around_rank2_softmaxClass()
}

type EspressoPass_remove_reshape_around_rank2_softmaxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_reshape_around_rank2_softmaxClass) Alloc() EspressoPass_remove_reshape_around_rank2_softmax {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank2_softmax](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank2_softmax
type EspressoPass_remove_reshape_around_rank2_softmax struct {
	EspressoCustomPass
}

// EspressoPass_remove_reshape_around_rank2_softmaxFromID constructs a [EspressoPass_remove_reshape_around_rank2_softmax] from an objc.ID.
func EspressoPass_remove_reshape_around_rank2_softmaxFromID(id objc.ID) EspressoPass_remove_reshape_around_rank2_softmax {
	return EspressoPass_remove_reshape_around_rank2_softmax{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_remove_reshape_around_rank2_softmax implements IEspressoPass_remove_reshape_around_rank2_softmax.
var _ IEspressoPass_remove_reshape_around_rank2_softmax = EspressoPass_remove_reshape_around_rank2_softmax{}

// An interface definition for the [EspressoPass_remove_reshape_around_rank2_softmax] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank2_softmax
type IEspressoPass_remove_reshape_around_rank2_softmax interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_reshape_around_rank2_softmax) Init() EspressoPass_remove_reshape_around_rank2_softmax {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank2_softmax](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_reshape_around_rank2_softmax) Autorelease() EspressoPass_remove_reshape_around_rank2_softmax {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank2_softmax](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_reshape_around_rank2_softmax creates a new EspressoPass_remove_reshape_around_rank2_softmax instance.
func NewEspressoPass_remove_reshape_around_rank2_softmax() EspressoPass_remove_reshape_around_rank2_softmax {
	class := getEspressoPass_remove_reshape_around_rank2_softmaxClass()
	rv := objc.Send[EspressoPass_remove_reshape_around_rank2_softmax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

