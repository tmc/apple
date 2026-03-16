// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_reshape_around_rank3_inner_product] class.
var (
	_EspressoPass_remove_reshape_around_rank3_inner_productClass     EspressoPass_remove_reshape_around_rank3_inner_productClass
	_EspressoPass_remove_reshape_around_rank3_inner_productClassOnce sync.Once
)

func getEspressoPass_remove_reshape_around_rank3_inner_productClass() EspressoPass_remove_reshape_around_rank3_inner_productClass {
	_EspressoPass_remove_reshape_around_rank3_inner_productClassOnce.Do(func() {
		_EspressoPass_remove_reshape_around_rank3_inner_productClass = EspressoPass_remove_reshape_around_rank3_inner_productClass{class: objc.GetClass("EspressoPass_remove_reshape_around_rank3_inner_product")}
	})
	return _EspressoPass_remove_reshape_around_rank3_inner_productClass
}

// GetEspressoPass_remove_reshape_around_rank3_inner_productClass returns the class object for EspressoPass_remove_reshape_around_rank3_inner_product.
func GetEspressoPass_remove_reshape_around_rank3_inner_productClass() EspressoPass_remove_reshape_around_rank3_inner_productClass {
	return getEspressoPass_remove_reshape_around_rank3_inner_productClass()
}

type EspressoPass_remove_reshape_around_rank3_inner_productClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_reshape_around_rank3_inner_productClass) Alloc() EspressoPass_remove_reshape_around_rank3_inner_product {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank3_inner_product](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank3_inner_product
type EspressoPass_remove_reshape_around_rank3_inner_product struct {
	EspressoCustomPass
}

// EspressoPass_remove_reshape_around_rank3_inner_productFromID constructs a [EspressoPass_remove_reshape_around_rank3_inner_product] from an objc.ID.
func EspressoPass_remove_reshape_around_rank3_inner_productFromID(id objc.ID) EspressoPass_remove_reshape_around_rank3_inner_product {
	return EspressoPass_remove_reshape_around_rank3_inner_product{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_remove_reshape_around_rank3_inner_product implements IEspressoPass_remove_reshape_around_rank3_inner_product.
var _ IEspressoPass_remove_reshape_around_rank3_inner_product = EspressoPass_remove_reshape_around_rank3_inner_product{}

// An interface definition for the [EspressoPass_remove_reshape_around_rank3_inner_product] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank3_inner_product
type IEspressoPass_remove_reshape_around_rank3_inner_product interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_reshape_around_rank3_inner_product) Init() EspressoPass_remove_reshape_around_rank3_inner_product {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank3_inner_product](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_reshape_around_rank3_inner_product) Autorelease() EspressoPass_remove_reshape_around_rank3_inner_product {
	rv := objc.Send[EspressoPass_remove_reshape_around_rank3_inner_product](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_reshape_around_rank3_inner_product creates a new EspressoPass_remove_reshape_around_rank3_inner_product instance.
func NewEspressoPass_remove_reshape_around_rank3_inner_product() EspressoPass_remove_reshape_around_rank3_inner_product {
	class := getEspressoPass_remove_reshape_around_rank3_inner_productClass()
	rv := objc.Send[EspressoPass_remove_reshape_around_rank3_inner_product](objc.ID(class.class), objc.Sel("new"))
	return rv
}

