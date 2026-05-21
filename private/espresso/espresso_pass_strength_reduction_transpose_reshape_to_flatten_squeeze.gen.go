// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze] class.
var (
	_EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass     EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass
	_EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClassOnce sync.Once
)

func getEspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass {
	_EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClassOnce.Do(func() {
		_EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass = EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass{class: objc.GetClass("EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze")}
	})
	return _EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass
}

// GetEspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass returns the class object for EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeeze.
func GetEspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass {
	return getEspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass()
}

type EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass) Alloc() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	rv := objc.Send[EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeFromID constructs a [EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze] from an objc.ID.
func EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeFromID(id objc.ID) EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	return EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeFromID is an alias for [EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_transpose_reshape_to_flatten_squeezeFromID(id objc.ID) EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	return EspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeFromID(id)
}

// Ensure EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze implements IEspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze.
var _ IEspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze = EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze{}

// An interface definition for the [EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze] class.
type IEspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze) Init() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	rv := objc.Send[EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze) Autorelease() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	rv := objc.Send[EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze creates a new EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze instance.
func NewEspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze() EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze {
	class := getEspressoPassStrengthReductionTransposeReshapeToFlattenSqueezeClass()
	rv := objc.Send[EspressoPassStrengthReductionTransposeReshapeToFlattenSqueeze](objc.ID(class.class), objc.Sel("new"))
	return rv
}
