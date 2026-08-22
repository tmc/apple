// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveReshapeAroundRank3InnerProduct] class.
var (
	_EspressoPassRemoveReshapeAroundRank3InnerProductClass     EspressoPassRemoveReshapeAroundRank3InnerProductClass
	_EspressoPassRemoveReshapeAroundRank3InnerProductClassOnce sync.Once
)

func getEspressoPassRemoveReshapeAroundRank3InnerProductClass() EspressoPassRemoveReshapeAroundRank3InnerProductClass {
	_EspressoPassRemoveReshapeAroundRank3InnerProductClassOnce.Do(func() {
		_EspressoPassRemoveReshapeAroundRank3InnerProductClass = EspressoPassRemoveReshapeAroundRank3InnerProductClass{class: objc.GetClass("EspressoPass_remove_reshape_around_rank3_inner_product")}
	})
	return _EspressoPassRemoveReshapeAroundRank3InnerProductClass
}

// GetEspressoPassRemoveReshapeAroundRank3InnerProductClass returns the class object for EspressoPass_remove_reshape_around_rank3_inner_product.
func GetEspressoPassRemoveReshapeAroundRank3InnerProductClass() EspressoPassRemoveReshapeAroundRank3InnerProductClass {
	return getEspressoPassRemoveReshapeAroundRank3InnerProductClass()
}

type EspressoPassRemoveReshapeAroundRank3InnerProductClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveReshapeAroundRank3InnerProductClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveReshapeAroundRank3InnerProductClass) Alloc() EspressoPassRemoveReshapeAroundRank3InnerProduct {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeAroundRank3InnerProduct](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassRemoveReshapeAroundRank3InnerProduct struct {
	EspressoCustomPass
}

// EspressoPassRemoveReshapeAroundRank3InnerProductFromID constructs a [EspressoPassRemoveReshapeAroundRank3InnerProduct] from an objc.ID.
func EspressoPassRemoveReshapeAroundRank3InnerProductFromID(id objc.ID) EspressoPassRemoveReshapeAroundRank3InnerProduct {
	return EspressoPassRemoveReshapeAroundRank3InnerProduct{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_reshape_around_rank3_inner_productFromID is an alias for [EspressoPassRemoveReshapeAroundRank3InnerProductFromID] for cross-framework compatibility.
func EspressoPass_remove_reshape_around_rank3_inner_productFromID(id objc.ID) EspressoPassRemoveReshapeAroundRank3InnerProduct {
	return EspressoPassRemoveReshapeAroundRank3InnerProductFromID(id)
}

// Ensure EspressoPassRemoveReshapeAroundRank3InnerProduct implements IEspressoPassRemoveReshapeAroundRank3InnerProduct.
var _ IEspressoPassRemoveReshapeAroundRank3InnerProduct = EspressoPassRemoveReshapeAroundRank3InnerProduct{}

// An interface definition for the [EspressoPassRemoveReshapeAroundRank3InnerProduct] class.
type IEspressoPassRemoveReshapeAroundRank3InnerProduct interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveReshapeAroundRank3InnerProduct) Init() EspressoPassRemoveReshapeAroundRank3InnerProduct {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeAroundRank3InnerProduct](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveReshapeAroundRank3InnerProduct) Autorelease() EspressoPassRemoveReshapeAroundRank3InnerProduct {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeAroundRank3InnerProduct](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveReshapeAroundRank3InnerProduct creates a new EspressoPassRemoveReshapeAroundRank3InnerProduct instance.
func NewEspressoPassRemoveReshapeAroundRank3InnerProduct() EspressoPassRemoveReshapeAroundRank3InnerProduct {
	class := getEspressoPassRemoveReshapeAroundRank3InnerProductClass()
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeAroundRank3InnerProduct](objc.ID(class.class), objc.Sel("new"))
	return rv
}
