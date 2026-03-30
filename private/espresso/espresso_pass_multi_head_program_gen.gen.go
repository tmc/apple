// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_multi_head_program_gen] class.
var (
	_EspressoPass_multi_head_program_genClass     EspressoPass_multi_head_program_genClass
	_EspressoPass_multi_head_program_genClassOnce sync.Once
)

func getEspressoPass_multi_head_program_genClass() EspressoPass_multi_head_program_genClass {
	_EspressoPass_multi_head_program_genClassOnce.Do(func() {
		_EspressoPass_multi_head_program_genClass = EspressoPass_multi_head_program_genClass{class: objc.GetClass("EspressoPass_multi_head_program_gen")}
	})
	return _EspressoPass_multi_head_program_genClass
}

// GetEspressoPass_multi_head_program_genClass returns the class object for EspressoPass_multi_head_program_gen.
func GetEspressoPass_multi_head_program_genClass() EspressoPass_multi_head_program_genClass {
	return getEspressoPass_multi_head_program_genClass()
}

type EspressoPass_multi_head_program_genClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_multi_head_program_genClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_multi_head_program_genClass) Alloc() EspressoPass_multi_head_program_gen {
	rv := objc.Send[EspressoPass_multi_head_program_gen](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_program_gen
type EspressoPass_multi_head_program_gen struct {
	EspressoCustomPass
}

// EspressoPass_multi_head_program_genFromID constructs a [EspressoPass_multi_head_program_gen] from an objc.ID.
func EspressoPass_multi_head_program_genFromID(id objc.ID) EspressoPass_multi_head_program_gen {
	return EspressoPass_multi_head_program_gen{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_multi_head_program_gen implements IEspressoPass_multi_head_program_gen.
var _ IEspressoPass_multi_head_program_gen = EspressoPass_multi_head_program_gen{}

// An interface definition for the [EspressoPass_multi_head_program_gen] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_program_gen
type IEspressoPass_multi_head_program_gen interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_multi_head_program_gen) Init() EspressoPass_multi_head_program_gen {
	rv := objc.Send[EspressoPass_multi_head_program_gen](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_multi_head_program_gen) Autorelease() EspressoPass_multi_head_program_gen {
	rv := objc.Send[EspressoPass_multi_head_program_gen](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_multi_head_program_gen creates a new EspressoPass_multi_head_program_gen instance.
func NewEspressoPass_multi_head_program_gen() EspressoPass_multi_head_program_gen {
	class := getEspressoPass_multi_head_program_genClass()
	rv := objc.Send[EspressoPass_multi_head_program_gen](objc.ID(class.class), objc.Sel("new"))
	return rv
}
