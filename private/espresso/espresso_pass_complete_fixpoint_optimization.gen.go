// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassCompleteFixpointOptimization] class.
var (
	_EspressoPassCompleteFixpointOptimizationClass     EspressoPassCompleteFixpointOptimizationClass
	_EspressoPassCompleteFixpointOptimizationClassOnce sync.Once
)

func getEspressoPassCompleteFixpointOptimizationClass() EspressoPassCompleteFixpointOptimizationClass {
	_EspressoPassCompleteFixpointOptimizationClassOnce.Do(func() {
		_EspressoPassCompleteFixpointOptimizationClass = EspressoPassCompleteFixpointOptimizationClass{class: objc.GetClass("EspressoPass_complete_fixpoint_optimization")}
	})
	return _EspressoPassCompleteFixpointOptimizationClass
}

// GetEspressoPassCompleteFixpointOptimizationClass returns the class object for EspressoPass_complete_fixpoint_optimization.
func GetEspressoPassCompleteFixpointOptimizationClass() EspressoPassCompleteFixpointOptimizationClass {
	return getEspressoPassCompleteFixpointOptimizationClass()
}

type EspressoPassCompleteFixpointOptimizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassCompleteFixpointOptimizationClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassCompleteFixpointOptimizationClass) Alloc() EspressoPassCompleteFixpointOptimization {
	rv := objc.SendIfResponds[EspressoPassCompleteFixpointOptimization](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassCompleteFixpointOptimization struct {
	EspressoCustomPass
}

// EspressoPassCompleteFixpointOptimizationFromID constructs a [EspressoPassCompleteFixpointOptimization] from an objc.ID.
func EspressoPassCompleteFixpointOptimizationFromID(id objc.ID) EspressoPassCompleteFixpointOptimization {
	return EspressoPassCompleteFixpointOptimization{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_complete_fixpoint_optimizationFromID is an alias for [EspressoPassCompleteFixpointOptimizationFromID] for cross-framework compatibility.
func EspressoPass_complete_fixpoint_optimizationFromID(id objc.ID) EspressoPassCompleteFixpointOptimization {
	return EspressoPassCompleteFixpointOptimizationFromID(id)
}

// Ensure EspressoPassCompleteFixpointOptimization implements IEspressoPassCompleteFixpointOptimization.
var _ IEspressoPassCompleteFixpointOptimization = EspressoPassCompleteFixpointOptimization{}

// An interface definition for the [EspressoPassCompleteFixpointOptimization] class.
type IEspressoPassCompleteFixpointOptimization interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassCompleteFixpointOptimization) Init() EspressoPassCompleteFixpointOptimization {
	rv := objc.SendIfResponds[EspressoPassCompleteFixpointOptimization](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassCompleteFixpointOptimization) Autorelease() EspressoPassCompleteFixpointOptimization {
	rv := objc.SendIfResponds[EspressoPassCompleteFixpointOptimization](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassCompleteFixpointOptimization creates a new EspressoPassCompleteFixpointOptimization instance.
func NewEspressoPassCompleteFixpointOptimization() EspressoPassCompleteFixpointOptimization {
	class := getEspressoPassCompleteFixpointOptimizationClass()
	rv := objc.SendIfResponds[EspressoPassCompleteFixpointOptimization](objc.ID(class.class), objc.Sel("new"))
	return rv
}
