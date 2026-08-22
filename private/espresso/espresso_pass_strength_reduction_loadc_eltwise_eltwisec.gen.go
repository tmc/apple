// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionLoadcEltwiseEltwisec] class.
var (
	_EspressoPassStrengthReductionLoadcEltwiseEltwisecClass     EspressoPassStrengthReductionLoadcEltwiseEltwisecClass
	_EspressoPassStrengthReductionLoadcEltwiseEltwisecClassOnce sync.Once
)

func getEspressoPassStrengthReductionLoadcEltwiseEltwisecClass() EspressoPassStrengthReductionLoadcEltwiseEltwisecClass {
	_EspressoPassStrengthReductionLoadcEltwiseEltwisecClassOnce.Do(func() {
		_EspressoPassStrengthReductionLoadcEltwiseEltwisecClass = EspressoPassStrengthReductionLoadcEltwiseEltwisecClass{class: objc.GetClass("EspressoPass_strength_reduction_loadc_eltwise__eltwisec")}
	})
	return _EspressoPassStrengthReductionLoadcEltwiseEltwisecClass
}

// GetEspressoPassStrengthReductionLoadcEltwiseEltwisecClass returns the class object for EspressoPass_strength_reduction_loadc_eltwise__eltwisec.
func GetEspressoPassStrengthReductionLoadcEltwiseEltwisecClass() EspressoPassStrengthReductionLoadcEltwiseEltwisecClass {
	return getEspressoPassStrengthReductionLoadcEltwiseEltwisecClass()
}

type EspressoPassStrengthReductionLoadcEltwiseEltwisecClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionLoadcEltwiseEltwisecClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionLoadcEltwiseEltwisecClass) Alloc() EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionLoadcEltwiseEltwisec](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionLoadcEltwiseEltwisec struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionLoadcEltwiseEltwisecFromID constructs a [EspressoPassStrengthReductionLoadcEltwiseEltwisec] from an objc.ID.
func EspressoPassStrengthReductionLoadcEltwiseEltwisecFromID(id objc.ID) EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	return EspressoPassStrengthReductionLoadcEltwiseEltwisec{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_loadc_eltwise__eltwisecFromID is an alias for [EspressoPassStrengthReductionLoadcEltwiseEltwisecFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_loadc_eltwise__eltwisecFromID(id objc.ID) EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	return EspressoPassStrengthReductionLoadcEltwiseEltwisecFromID(id)
}

// Ensure EspressoPassStrengthReductionLoadcEltwiseEltwisec implements IEspressoPassStrengthReductionLoadcEltwiseEltwisec.
var _ IEspressoPassStrengthReductionLoadcEltwiseEltwisec = EspressoPassStrengthReductionLoadcEltwiseEltwisec{}

// An interface definition for the [EspressoPassStrengthReductionLoadcEltwiseEltwisec] class.
type IEspressoPassStrengthReductionLoadcEltwiseEltwisec interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionLoadcEltwiseEltwisec) Init() EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionLoadcEltwiseEltwisec](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionLoadcEltwiseEltwisec) Autorelease() EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionLoadcEltwiseEltwisec](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionLoadcEltwiseEltwisec creates a new EspressoPassStrengthReductionLoadcEltwiseEltwisec instance.
func NewEspressoPassStrengthReductionLoadcEltwiseEltwisec() EspressoPassStrengthReductionLoadcEltwiseEltwisec {
	class := getEspressoPassStrengthReductionLoadcEltwiseEltwisecClass()
	rv := objc.SendIfResponds[EspressoPassStrengthReductionLoadcEltwiseEltwisec](objc.ID(class.class), objc.Sel("new"))
	return rv
}
