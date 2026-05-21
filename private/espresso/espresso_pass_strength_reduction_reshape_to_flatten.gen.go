// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionReshapeToFlatten] class.
var (
	_EspressoPassStrengthReductionReshapeToFlattenClass     EspressoPassStrengthReductionReshapeToFlattenClass
	_EspressoPassStrengthReductionReshapeToFlattenClassOnce sync.Once
)

func getEspressoPassStrengthReductionReshapeToFlattenClass() EspressoPassStrengthReductionReshapeToFlattenClass {
	_EspressoPassStrengthReductionReshapeToFlattenClassOnce.Do(func() {
		_EspressoPassStrengthReductionReshapeToFlattenClass = EspressoPassStrengthReductionReshapeToFlattenClass{class: objc.GetClass("EspressoPass_strength_reduction_reshape_to_flatten")}
	})
	return _EspressoPassStrengthReductionReshapeToFlattenClass
}

// GetEspressoPassStrengthReductionReshapeToFlattenClass returns the class object for EspressoPass_strength_reduction_reshape_to_flatten.
func GetEspressoPassStrengthReductionReshapeToFlattenClass() EspressoPassStrengthReductionReshapeToFlattenClass {
	return getEspressoPassStrengthReductionReshapeToFlattenClass()
}

type EspressoPassStrengthReductionReshapeToFlattenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionReshapeToFlattenClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionReshapeToFlattenClass) Alloc() EspressoPassStrengthReductionReshapeToFlatten {
	rv := objc.Send[EspressoPassStrengthReductionReshapeToFlatten](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionReshapeToFlatten struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionReshapeToFlattenFromID constructs a [EspressoPassStrengthReductionReshapeToFlatten] from an objc.ID.
func EspressoPassStrengthReductionReshapeToFlattenFromID(id objc.ID) EspressoPassStrengthReductionReshapeToFlatten {
	return EspressoPassStrengthReductionReshapeToFlatten{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_reshape_to_flattenFromID is an alias for [EspressoPassStrengthReductionReshapeToFlattenFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_reshape_to_flattenFromID(id objc.ID) EspressoPassStrengthReductionReshapeToFlatten {
	return EspressoPassStrengthReductionReshapeToFlattenFromID(id)
}

// Ensure EspressoPassStrengthReductionReshapeToFlatten implements IEspressoPassStrengthReductionReshapeToFlatten.
var _ IEspressoPassStrengthReductionReshapeToFlatten = EspressoPassStrengthReductionReshapeToFlatten{}

// An interface definition for the [EspressoPassStrengthReductionReshapeToFlatten] class.
type IEspressoPassStrengthReductionReshapeToFlatten interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionReshapeToFlatten) Init() EspressoPassStrengthReductionReshapeToFlatten {
	rv := objc.Send[EspressoPassStrengthReductionReshapeToFlatten](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionReshapeToFlatten) Autorelease() EspressoPassStrengthReductionReshapeToFlatten {
	rv := objc.Send[EspressoPassStrengthReductionReshapeToFlatten](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionReshapeToFlatten creates a new EspressoPassStrengthReductionReshapeToFlatten instance.
func NewEspressoPassStrengthReductionReshapeToFlatten() EspressoPassStrengthReductionReshapeToFlatten {
	class := getEspressoPassStrengthReductionReshapeToFlattenClass()
	rv := objc.Send[EspressoPassStrengthReductionReshapeToFlatten](objc.ID(class.class), objc.Sel("new"))
	return rv
}
