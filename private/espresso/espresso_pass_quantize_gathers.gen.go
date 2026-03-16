// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_quantize_gathers] class.
var (
	_EspressoPass_quantize_gathersClass     EspressoPass_quantize_gathersClass
	_EspressoPass_quantize_gathersClassOnce sync.Once
)

func getEspressoPass_quantize_gathersClass() EspressoPass_quantize_gathersClass {
	_EspressoPass_quantize_gathersClassOnce.Do(func() {
		_EspressoPass_quantize_gathersClass = EspressoPass_quantize_gathersClass{class: objc.GetClass("EspressoPass_quantize_gathers")}
	})
	return _EspressoPass_quantize_gathersClass
}

// GetEspressoPass_quantize_gathersClass returns the class object for EspressoPass_quantize_gathers.
func GetEspressoPass_quantize_gathersClass() EspressoPass_quantize_gathersClass {
	return getEspressoPass_quantize_gathersClass()
}

type EspressoPass_quantize_gathersClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_quantize_gathersClass) Alloc() EspressoPass_quantize_gathers {
	rv := objc.Send[EspressoPass_quantize_gathers](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_quantize_gathers
type EspressoPass_quantize_gathers struct {
	EspressoCustomPass
}

// EspressoPass_quantize_gathersFromID constructs a [EspressoPass_quantize_gathers] from an objc.ID.
func EspressoPass_quantize_gathersFromID(id objc.ID) EspressoPass_quantize_gathers {
	return EspressoPass_quantize_gathers{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_quantize_gathers implements IEspressoPass_quantize_gathers.
var _ IEspressoPass_quantize_gathers = EspressoPass_quantize_gathers{}

// An interface definition for the [EspressoPass_quantize_gathers] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_quantize_gathers
type IEspressoPass_quantize_gathers interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_quantize_gathers) Init() EspressoPass_quantize_gathers {
	rv := objc.Send[EspressoPass_quantize_gathers](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_quantize_gathers) Autorelease() EspressoPass_quantize_gathers {
	rv := objc.Send[EspressoPass_quantize_gathers](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_quantize_gathers creates a new EspressoPass_quantize_gathers instance.
func NewEspressoPass_quantize_gathers() EspressoPass_quantize_gathers {
	class := getEspressoPass_quantize_gathersClass()
	rv := objc.Send[EspressoPass_quantize_gathers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

