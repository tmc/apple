// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassTransposeInnerProductWeights] class.
var (
	_EspressoPassTransposeInnerProductWeightsClass     EspressoPassTransposeInnerProductWeightsClass
	_EspressoPassTransposeInnerProductWeightsClassOnce sync.Once
)

func getEspressoPassTransposeInnerProductWeightsClass() EspressoPassTransposeInnerProductWeightsClass {
	_EspressoPassTransposeInnerProductWeightsClassOnce.Do(func() {
		_EspressoPassTransposeInnerProductWeightsClass = EspressoPassTransposeInnerProductWeightsClass{class: objc.GetClass("EspressoPass_transpose_inner_product_weights")}
	})
	return _EspressoPassTransposeInnerProductWeightsClass
}

// GetEspressoPassTransposeInnerProductWeightsClass returns the class object for EspressoPass_transpose_inner_product_weights.
func GetEspressoPassTransposeInnerProductWeightsClass() EspressoPassTransposeInnerProductWeightsClass {
	return getEspressoPassTransposeInnerProductWeightsClass()
}

type EspressoPassTransposeInnerProductWeightsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassTransposeInnerProductWeightsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassTransposeInnerProductWeightsClass) Alloc() EspressoPassTransposeInnerProductWeights {
	rv := objc.Send[EspressoPassTransposeInnerProductWeights](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transpose_inner_product_weights
type EspressoPassTransposeInnerProductWeights struct {
	EspressoCustomPass
}

// EspressoPassTransposeInnerProductWeightsFromID constructs a [EspressoPassTransposeInnerProductWeights] from an objc.ID.
func EspressoPassTransposeInnerProductWeightsFromID(id objc.ID) EspressoPassTransposeInnerProductWeights {
	return EspressoPassTransposeInnerProductWeights{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_transpose_inner_product_weightsFromID is an alias for [EspressoPassTransposeInnerProductWeightsFromID] for cross-framework compatibility.
func EspressoPass_transpose_inner_product_weightsFromID(id objc.ID) EspressoPassTransposeInnerProductWeights {
	return EspressoPassTransposeInnerProductWeightsFromID(id)
}

// Ensure EspressoPassTransposeInnerProductWeights implements IEspressoPassTransposeInnerProductWeights.
var _ IEspressoPassTransposeInnerProductWeights = EspressoPassTransposeInnerProductWeights{}

// An interface definition for the [EspressoPassTransposeInnerProductWeights] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transpose_inner_product_weights
type IEspressoPassTransposeInnerProductWeights interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassTransposeInnerProductWeights) Init() EspressoPassTransposeInnerProductWeights {
	rv := objc.Send[EspressoPassTransposeInnerProductWeights](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassTransposeInnerProductWeights) Autorelease() EspressoPassTransposeInnerProductWeights {
	rv := objc.Send[EspressoPassTransposeInnerProductWeights](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassTransposeInnerProductWeights creates a new EspressoPassTransposeInnerProductWeights instance.
func NewEspressoPassTransposeInnerProductWeights() EspressoPassTransposeInnerProductWeights {
	class := getEspressoPassTransposeInnerProductWeightsClass()
	rv := objc.Send[EspressoPassTransposeInnerProductWeights](objc.ID(class.class), objc.Sel("new"))
	return rv
}
