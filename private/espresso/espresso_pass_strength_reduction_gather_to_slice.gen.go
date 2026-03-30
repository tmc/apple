// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_strength_reduction_gather_to_slice] class.
var (
	_EspressoPass_strength_reduction_gather_to_sliceClass     EspressoPass_strength_reduction_gather_to_sliceClass
	_EspressoPass_strength_reduction_gather_to_sliceClassOnce sync.Once
)

func getEspressoPass_strength_reduction_gather_to_sliceClass() EspressoPass_strength_reduction_gather_to_sliceClass {
	_EspressoPass_strength_reduction_gather_to_sliceClassOnce.Do(func() {
		_EspressoPass_strength_reduction_gather_to_sliceClass = EspressoPass_strength_reduction_gather_to_sliceClass{class: objc.GetClass("EspressoPass_strength_reduction_gather_to_slice")}
	})
	return _EspressoPass_strength_reduction_gather_to_sliceClass
}

// GetEspressoPass_strength_reduction_gather_to_sliceClass returns the class object for EspressoPass_strength_reduction_gather_to_slice.
func GetEspressoPass_strength_reduction_gather_to_sliceClass() EspressoPass_strength_reduction_gather_to_sliceClass {
	return getEspressoPass_strength_reduction_gather_to_sliceClass()
}

type EspressoPass_strength_reduction_gather_to_sliceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_strength_reduction_gather_to_sliceClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_strength_reduction_gather_to_sliceClass) Alloc() EspressoPass_strength_reduction_gather_to_slice {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_slice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_slice
type EspressoPass_strength_reduction_gather_to_slice struct {
	EspressoCustomPass
}

// EspressoPass_strength_reduction_gather_to_sliceFromID constructs a [EspressoPass_strength_reduction_gather_to_slice] from an objc.ID.
func EspressoPass_strength_reduction_gather_to_sliceFromID(id objc.ID) EspressoPass_strength_reduction_gather_to_slice {
	return EspressoPass_strength_reduction_gather_to_slice{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_strength_reduction_gather_to_slice implements IEspressoPass_strength_reduction_gather_to_slice.
var _ IEspressoPass_strength_reduction_gather_to_slice = EspressoPass_strength_reduction_gather_to_slice{}

// An interface definition for the [EspressoPass_strength_reduction_gather_to_slice] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_slice
type IEspressoPass_strength_reduction_gather_to_slice interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_strength_reduction_gather_to_slice) Init() EspressoPass_strength_reduction_gather_to_slice {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_slice](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_strength_reduction_gather_to_slice) Autorelease() EspressoPass_strength_reduction_gather_to_slice {
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_slice](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_strength_reduction_gather_to_slice creates a new EspressoPass_strength_reduction_gather_to_slice instance.
func NewEspressoPass_strength_reduction_gather_to_slice() EspressoPass_strength_reduction_gather_to_slice {
	class := getEspressoPass_strength_reduction_gather_to_sliceClass()
	rv := objc.Send[EspressoPass_strength_reduction_gather_to_slice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
