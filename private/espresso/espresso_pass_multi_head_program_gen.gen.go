// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMultiHeadProgramGen] class.
var (
	_EspressoPassMultiHeadProgramGenClass     EspressoPassMultiHeadProgramGenClass
	_EspressoPassMultiHeadProgramGenClassOnce sync.Once
)

func getEspressoPassMultiHeadProgramGenClass() EspressoPassMultiHeadProgramGenClass {
	_EspressoPassMultiHeadProgramGenClassOnce.Do(func() {
		_EspressoPassMultiHeadProgramGenClass = EspressoPassMultiHeadProgramGenClass{class: objc.GetClass("EspressoPass_multi_head_program_gen")}
	})
	return _EspressoPassMultiHeadProgramGenClass
}

// GetEspressoPassMultiHeadProgramGenClass returns the class object for EspressoPass_multi_head_program_gen.
func GetEspressoPassMultiHeadProgramGenClass() EspressoPassMultiHeadProgramGenClass {
	return getEspressoPassMultiHeadProgramGenClass()
}

type EspressoPassMultiHeadProgramGenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMultiHeadProgramGenClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMultiHeadProgramGenClass) Alloc() EspressoPassMultiHeadProgramGen {
	rv := objc.Send[EspressoPassMultiHeadProgramGen](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassMultiHeadProgramGen struct {
	EspressoCustomPass
}

// EspressoPassMultiHeadProgramGenFromID constructs a [EspressoPassMultiHeadProgramGen] from an objc.ID.
func EspressoPassMultiHeadProgramGenFromID(id objc.ID) EspressoPassMultiHeadProgramGen {
	return EspressoPassMultiHeadProgramGen{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_multi_head_program_genFromID is an alias for [EspressoPassMultiHeadProgramGenFromID] for cross-framework compatibility.
func EspressoPass_multi_head_program_genFromID(id objc.ID) EspressoPassMultiHeadProgramGen {
	return EspressoPassMultiHeadProgramGenFromID(id)
}

// Ensure EspressoPassMultiHeadProgramGen implements IEspressoPassMultiHeadProgramGen.
var _ IEspressoPassMultiHeadProgramGen = EspressoPassMultiHeadProgramGen{}

// An interface definition for the [EspressoPassMultiHeadProgramGen] class.
type IEspressoPassMultiHeadProgramGen interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMultiHeadProgramGen) Init() EspressoPassMultiHeadProgramGen {
	rv := objc.Send[EspressoPassMultiHeadProgramGen](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMultiHeadProgramGen) Autorelease() EspressoPassMultiHeadProgramGen {
	rv := objc.Send[EspressoPassMultiHeadProgramGen](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMultiHeadProgramGen creates a new EspressoPassMultiHeadProgramGen instance.
func NewEspressoPassMultiHeadProgramGen() EspressoPassMultiHeadProgramGen {
	class := getEspressoPassMultiHeadProgramGenClass()
	rv := objc.Send[EspressoPassMultiHeadProgramGen](objc.ID(class.class), objc.Sel("new"))
	return rv
}
