// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_complete_fixpoint_optimization] class.
var (
	_EspressoPass_complete_fixpoint_optimizationClass     EspressoPass_complete_fixpoint_optimizationClass
	_EspressoPass_complete_fixpoint_optimizationClassOnce sync.Once
)

func getEspressoPass_complete_fixpoint_optimizationClass() EspressoPass_complete_fixpoint_optimizationClass {
	_EspressoPass_complete_fixpoint_optimizationClassOnce.Do(func() {
		_EspressoPass_complete_fixpoint_optimizationClass = EspressoPass_complete_fixpoint_optimizationClass{class: objc.GetClass("EspressoPass_complete_fixpoint_optimization")}
	})
	return _EspressoPass_complete_fixpoint_optimizationClass
}

// GetEspressoPass_complete_fixpoint_optimizationClass returns the class object for EspressoPass_complete_fixpoint_optimization.
func GetEspressoPass_complete_fixpoint_optimizationClass() EspressoPass_complete_fixpoint_optimizationClass {
	return getEspressoPass_complete_fixpoint_optimizationClass()
}

type EspressoPass_complete_fixpoint_optimizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_complete_fixpoint_optimizationClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_complete_fixpoint_optimizationClass) Alloc() EspressoPass_complete_fixpoint_optimization {
	rv := objc.Send[EspressoPass_complete_fixpoint_optimization](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_complete_fixpoint_optimization
type EspressoPass_complete_fixpoint_optimization struct {
	EspressoCustomPass
}

// EspressoPass_complete_fixpoint_optimizationFromID constructs a [EspressoPass_complete_fixpoint_optimization] from an objc.ID.
func EspressoPass_complete_fixpoint_optimizationFromID(id objc.ID) EspressoPass_complete_fixpoint_optimization {
	return EspressoPass_complete_fixpoint_optimization{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_complete_fixpoint_optimization implements IEspressoPass_complete_fixpoint_optimization.
var _ IEspressoPass_complete_fixpoint_optimization = EspressoPass_complete_fixpoint_optimization{}

// An interface definition for the [EspressoPass_complete_fixpoint_optimization] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_complete_fixpoint_optimization
type IEspressoPass_complete_fixpoint_optimization interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_complete_fixpoint_optimization) Init() EspressoPass_complete_fixpoint_optimization {
	rv := objc.Send[EspressoPass_complete_fixpoint_optimization](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_complete_fixpoint_optimization) Autorelease() EspressoPass_complete_fixpoint_optimization {
	rv := objc.Send[EspressoPass_complete_fixpoint_optimization](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_complete_fixpoint_optimization creates a new EspressoPass_complete_fixpoint_optimization instance.
func NewEspressoPass_complete_fixpoint_optimization() EspressoPass_complete_fixpoint_optimization {
	class := getEspressoPass_complete_fixpoint_optimizationClass()
	rv := objc.Send[EspressoPass_complete_fixpoint_optimization](objc.ID(class.class), objc.Sel("new"))
	return rv
}
