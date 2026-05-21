// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionBatchMatmulToInnerProduct] class.
var (
	_EspressoPassStrengthReductionBatchMatmulToInnerProductClass     EspressoPassStrengthReductionBatchMatmulToInnerProductClass
	_EspressoPassStrengthReductionBatchMatmulToInnerProductClassOnce sync.Once
)

func getEspressoPassStrengthReductionBatchMatmulToInnerProductClass() EspressoPassStrengthReductionBatchMatmulToInnerProductClass {
	_EspressoPassStrengthReductionBatchMatmulToInnerProductClassOnce.Do(func() {
		_EspressoPassStrengthReductionBatchMatmulToInnerProductClass = EspressoPassStrengthReductionBatchMatmulToInnerProductClass{class: objc.GetClass("EspressoPass_strength_reduction_batch_matmul_to_inner_product")}
	})
	return _EspressoPassStrengthReductionBatchMatmulToInnerProductClass
}

// GetEspressoPassStrengthReductionBatchMatmulToInnerProductClass returns the class object for EspressoPass_strength_reduction_batch_matmul_to_inner_product.
func GetEspressoPassStrengthReductionBatchMatmulToInnerProductClass() EspressoPassStrengthReductionBatchMatmulToInnerProductClass {
	return getEspressoPassStrengthReductionBatchMatmulToInnerProductClass()
}

type EspressoPassStrengthReductionBatchMatmulToInnerProductClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionBatchMatmulToInnerProductClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionBatchMatmulToInnerProductClass) Alloc() EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	rv := objc.Send[EspressoPassStrengthReductionBatchMatmulToInnerProduct](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionBatchMatmulToInnerProduct struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionBatchMatmulToInnerProductFromID constructs a [EspressoPassStrengthReductionBatchMatmulToInnerProduct] from an objc.ID.
func EspressoPassStrengthReductionBatchMatmulToInnerProductFromID(id objc.ID) EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	return EspressoPassStrengthReductionBatchMatmulToInnerProduct{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_batch_matmul_to_inner_productFromID is an alias for [EspressoPassStrengthReductionBatchMatmulToInnerProductFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_batch_matmul_to_inner_productFromID(id objc.ID) EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	return EspressoPassStrengthReductionBatchMatmulToInnerProductFromID(id)
}

// Ensure EspressoPassStrengthReductionBatchMatmulToInnerProduct implements IEspressoPassStrengthReductionBatchMatmulToInnerProduct.
var _ IEspressoPassStrengthReductionBatchMatmulToInnerProduct = EspressoPassStrengthReductionBatchMatmulToInnerProduct{}

// An interface definition for the [EspressoPassStrengthReductionBatchMatmulToInnerProduct] class.
type IEspressoPassStrengthReductionBatchMatmulToInnerProduct interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionBatchMatmulToInnerProduct) Init() EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	rv := objc.Send[EspressoPassStrengthReductionBatchMatmulToInnerProduct](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionBatchMatmulToInnerProduct) Autorelease() EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	rv := objc.Send[EspressoPassStrengthReductionBatchMatmulToInnerProduct](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionBatchMatmulToInnerProduct creates a new EspressoPassStrengthReductionBatchMatmulToInnerProduct instance.
func NewEspressoPassStrengthReductionBatchMatmulToInnerProduct() EspressoPassStrengthReductionBatchMatmulToInnerProduct {
	class := getEspressoPassStrengthReductionBatchMatmulToInnerProductClass()
	rv := objc.Send[EspressoPassStrengthReductionBatchMatmulToInnerProduct](objc.ID(class.class), objc.Sel("new"))
	return rv
}
