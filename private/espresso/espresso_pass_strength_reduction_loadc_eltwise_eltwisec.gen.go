// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_loadc_eltwise__eltwisec] class.
var (
	_EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass     EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass
	_EspressoPass_strength_reduction_loadc_eltwise__eltwisecClassOnce sync.Once
)

func getEspressoPass_strength_reduction_loadc_eltwise__eltwisecClass() EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass {
	_EspressoPass_strength_reduction_loadc_eltwise__eltwisecClassOnce.Do(func() {
		_EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass = EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass{class: objc.GetClass("EspressoPass_strength_reduction_loadc_eltwise__eltwisec")}
	})
	return _EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass
}

// GetEspressoPass_strength_reduction_loadc_eltwise__eltwisecClass returns the class object for EspressoPass_strength_reduction_loadc_eltwise__eltwisec.
func GetEspressoPass_strength_reduction_loadc_eltwise__eltwisecClass() EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass {
	return getEspressoPass_strength_reduction_loadc_eltwise__eltwisecClass()
}

type EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_loadc_eltwise__eltwisecClass) Alloc() EspressoPass_strength_reduction_loadc_eltwise__eltwisec {
	rv := objc.Send[EspressoPass_strength_reduction_loadc_eltwise__eltwisec](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_loadc_eltwise__eltwisec
type EspressoPass_strength_reduction_loadc_eltwise__eltwisec struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_loadc_eltwise__eltwisecFromID constructs a [EspressoPass_strength_reduction_loadc_eltwise__eltwisec] from an objc.ID.
func EspressoPass_strength_reduction_loadc_eltwise__eltwisecFromID(id objc.ID) EspressoPass_strength_reduction_loadc_eltwise__eltwisec {
	return EspressoPass_strength_reduction_loadc_eltwise__eltwisec{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_strength_reduction_loadc_eltwise__eltwisec implements IEspressoPass_strength_reduction_loadc_eltwise__eltwisec.
var _ IEspressoPass_strength_reduction_loadc_eltwise__eltwisec = EspressoPass_strength_reduction_loadc_eltwise__eltwisec{}

// An interface definition for the [EspressoPass_strength_reduction_loadc_eltwise__eltwisec] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_loadc_eltwise__eltwisec
type IEspressoPass_strength_reduction_loadc_eltwise__eltwisec interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_strength_reduction_loadc_eltwise__eltwisec) Init() EspressoPass_strength_reduction_loadc_eltwise__eltwisec {
	rv := objc.Send[EspressoPass_strength_reduction_loadc_eltwise__eltwisec](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_loadc_eltwise__eltwisec) Autorelease() EspressoPass_strength_reduction_loadc_eltwise__eltwisec {
	rv := objc.Send[EspressoPass_strength_reduction_loadc_eltwise__eltwisec](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_loadc_eltwise__eltwisec creates a new EspressoPass_strength_reduction_loadc_eltwise__eltwisec instance.
func NewEspressoPass_strength_reduction_loadc_eltwise__eltwisec() EspressoPass_strength_reduction_loadc_eltwise__eltwisec {
	class := getEspressoPass_strength_reduction_loadc_eltwise__eltwisecClass()
	rv := objc.Send[EspressoPass_strength_reduction_loadc_eltwise__eltwisec](objc.ID(class.class), objc.Sel("new"))
	return rv
}
