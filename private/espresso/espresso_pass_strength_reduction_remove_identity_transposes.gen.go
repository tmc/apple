// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionRemoveIdentityTransposes] class.
var (
	_EspressoPassStrengthReductionRemoveIdentityTransposesClass     EspressoPassStrengthReductionRemoveIdentityTransposesClass
	_EspressoPassStrengthReductionRemoveIdentityTransposesClassOnce sync.Once
)

func getEspressoPassStrengthReductionRemoveIdentityTransposesClass() EspressoPassStrengthReductionRemoveIdentityTransposesClass {
	_EspressoPassStrengthReductionRemoveIdentityTransposesClassOnce.Do(func() {
		_EspressoPassStrengthReductionRemoveIdentityTransposesClass = EspressoPassStrengthReductionRemoveIdentityTransposesClass{class: objc.GetClass("EspressoPass_strength_reduction_remove_identity_transposes")}
	})
	return _EspressoPassStrengthReductionRemoveIdentityTransposesClass
}

// GetEspressoPassStrengthReductionRemoveIdentityTransposesClass returns the class object for EspressoPass_strength_reduction_remove_identity_transposes.
func GetEspressoPassStrengthReductionRemoveIdentityTransposesClass() EspressoPassStrengthReductionRemoveIdentityTransposesClass {
	return getEspressoPassStrengthReductionRemoveIdentityTransposesClass()
}

type EspressoPassStrengthReductionRemoveIdentityTransposesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionRemoveIdentityTransposesClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionRemoveIdentityTransposesClass) Alloc() EspressoPassStrengthReductionRemoveIdentityTransposes {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionRemoveIdentityTransposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionRemoveIdentityTransposes struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionRemoveIdentityTransposesFromID constructs a [EspressoPassStrengthReductionRemoveIdentityTransposes] from an objc.ID.
func EspressoPassStrengthReductionRemoveIdentityTransposesFromID(id objc.ID) EspressoPassStrengthReductionRemoveIdentityTransposes {
	return EspressoPassStrengthReductionRemoveIdentityTransposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_remove_identity_transposesFromID is an alias for [EspressoPassStrengthReductionRemoveIdentityTransposesFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_remove_identity_transposesFromID(id objc.ID) EspressoPassStrengthReductionRemoveIdentityTransposes {
	return EspressoPassStrengthReductionRemoveIdentityTransposesFromID(id)
}

// Ensure EspressoPassStrengthReductionRemoveIdentityTransposes implements IEspressoPassStrengthReductionRemoveIdentityTransposes.
var _ IEspressoPassStrengthReductionRemoveIdentityTransposes = EspressoPassStrengthReductionRemoveIdentityTransposes{}

// An interface definition for the [EspressoPassStrengthReductionRemoveIdentityTransposes] class.
type IEspressoPassStrengthReductionRemoveIdentityTransposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionRemoveIdentityTransposes) Init() EspressoPassStrengthReductionRemoveIdentityTransposes {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionRemoveIdentityTransposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionRemoveIdentityTransposes) Autorelease() EspressoPassStrengthReductionRemoveIdentityTransposes {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionRemoveIdentityTransposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionRemoveIdentityTransposes creates a new EspressoPassStrengthReductionRemoveIdentityTransposes instance.
func NewEspressoPassStrengthReductionRemoveIdentityTransposes() EspressoPassStrengthReductionRemoveIdentityTransposes {
	class := getEspressoPassStrengthReductionRemoveIdentityTransposesClass()
	rv := objc.SendIfResponds[EspressoPassStrengthReductionRemoveIdentityTransposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}
