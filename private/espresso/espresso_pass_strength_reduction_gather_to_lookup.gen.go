// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionGatherToLookup] class.
var (
	_EspressoPassStrengthReductionGatherToLookupClass     EspressoPassStrengthReductionGatherToLookupClass
	_EspressoPassStrengthReductionGatherToLookupClassOnce sync.Once
)

func getEspressoPassStrengthReductionGatherToLookupClass() EspressoPassStrengthReductionGatherToLookupClass {
	_EspressoPassStrengthReductionGatherToLookupClassOnce.Do(func() {
		_EspressoPassStrengthReductionGatherToLookupClass = EspressoPassStrengthReductionGatherToLookupClass{class: objc.GetClass("EspressoPass_strength_reduction_gather_to_lookup")}
	})
	return _EspressoPassStrengthReductionGatherToLookupClass
}

// GetEspressoPassStrengthReductionGatherToLookupClass returns the class object for EspressoPass_strength_reduction_gather_to_lookup.
func GetEspressoPassStrengthReductionGatherToLookupClass() EspressoPassStrengthReductionGatherToLookupClass {
	return getEspressoPassStrengthReductionGatherToLookupClass()
}

type EspressoPassStrengthReductionGatherToLookupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionGatherToLookupClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionGatherToLookupClass) Alloc() EspressoPassStrengthReductionGatherToLookup {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionGatherToLookup](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionGatherToLookup struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionGatherToLookupFromID constructs a [EspressoPassStrengthReductionGatherToLookup] from an objc.ID.
func EspressoPassStrengthReductionGatherToLookupFromID(id objc.ID) EspressoPassStrengthReductionGatherToLookup {
	return EspressoPassStrengthReductionGatherToLookup{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_gather_to_lookupFromID is an alias for [EspressoPassStrengthReductionGatherToLookupFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_gather_to_lookupFromID(id objc.ID) EspressoPassStrengthReductionGatherToLookup {
	return EspressoPassStrengthReductionGatherToLookupFromID(id)
}

// Ensure EspressoPassStrengthReductionGatherToLookup implements IEspressoPassStrengthReductionGatherToLookup.
var _ IEspressoPassStrengthReductionGatherToLookup = EspressoPassStrengthReductionGatherToLookup{}

// An interface definition for the [EspressoPassStrengthReductionGatherToLookup] class.
type IEspressoPassStrengthReductionGatherToLookup interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionGatherToLookup) Init() EspressoPassStrengthReductionGatherToLookup {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionGatherToLookup](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionGatherToLookup) Autorelease() EspressoPassStrengthReductionGatherToLookup {
	rv := objc.SendIfResponds[EspressoPassStrengthReductionGatherToLookup](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionGatherToLookup creates a new EspressoPassStrengthReductionGatherToLookup instance.
func NewEspressoPassStrengthReductionGatherToLookup() EspressoPassStrengthReductionGatherToLookup {
	class := getEspressoPassStrengthReductionGatherToLookupClass()
	rv := objc.SendIfResponds[EspressoPassStrengthReductionGatherToLookup](objc.ID(class.class), objc.Sel("new"))
	return rv
}
