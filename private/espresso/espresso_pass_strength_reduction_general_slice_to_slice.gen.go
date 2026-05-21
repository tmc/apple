// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionGeneralSliceToSlice] class.
var (
	_EspressoPassStrengthReductionGeneralSliceToSliceClass     EspressoPassStrengthReductionGeneralSliceToSliceClass
	_EspressoPassStrengthReductionGeneralSliceToSliceClassOnce sync.Once
)

func getEspressoPassStrengthReductionGeneralSliceToSliceClass() EspressoPassStrengthReductionGeneralSliceToSliceClass {
	_EspressoPassStrengthReductionGeneralSliceToSliceClassOnce.Do(func() {
		_EspressoPassStrengthReductionGeneralSliceToSliceClass = EspressoPassStrengthReductionGeneralSliceToSliceClass{class: objc.GetClass("EspressoPass_strength_reduction_general_slice_to_slice")}
	})
	return _EspressoPassStrengthReductionGeneralSliceToSliceClass
}

// GetEspressoPassStrengthReductionGeneralSliceToSliceClass returns the class object for EspressoPass_strength_reduction_general_slice_to_slice.
func GetEspressoPassStrengthReductionGeneralSliceToSliceClass() EspressoPassStrengthReductionGeneralSliceToSliceClass {
	return getEspressoPassStrengthReductionGeneralSliceToSliceClass()
}

type EspressoPassStrengthReductionGeneralSliceToSliceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionGeneralSliceToSliceClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionGeneralSliceToSliceClass) Alloc() EspressoPassStrengthReductionGeneralSliceToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGeneralSliceToSlice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassStrengthReductionGeneralSliceToSlice struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionGeneralSliceToSliceFromID constructs a [EspressoPassStrengthReductionGeneralSliceToSlice] from an objc.ID.
func EspressoPassStrengthReductionGeneralSliceToSliceFromID(id objc.ID) EspressoPassStrengthReductionGeneralSliceToSlice {
	return EspressoPassStrengthReductionGeneralSliceToSlice{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_general_slice_to_sliceFromID is an alias for [EspressoPassStrengthReductionGeneralSliceToSliceFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_general_slice_to_sliceFromID(id objc.ID) EspressoPassStrengthReductionGeneralSliceToSlice {
	return EspressoPassStrengthReductionGeneralSliceToSliceFromID(id)
}

// Ensure EspressoPassStrengthReductionGeneralSliceToSlice implements IEspressoPassStrengthReductionGeneralSliceToSlice.
var _ IEspressoPassStrengthReductionGeneralSliceToSlice = EspressoPassStrengthReductionGeneralSliceToSlice{}

// An interface definition for the [EspressoPassStrengthReductionGeneralSliceToSlice] class.
type IEspressoPassStrengthReductionGeneralSliceToSlice interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionGeneralSliceToSlice) Init() EspressoPassStrengthReductionGeneralSliceToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGeneralSliceToSlice](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionGeneralSliceToSlice) Autorelease() EspressoPassStrengthReductionGeneralSliceToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGeneralSliceToSlice](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionGeneralSliceToSlice creates a new EspressoPassStrengthReductionGeneralSliceToSlice instance.
func NewEspressoPassStrengthReductionGeneralSliceToSlice() EspressoPassStrengthReductionGeneralSliceToSlice {
	class := getEspressoPassStrengthReductionGeneralSliceToSliceClass()
	rv := objc.Send[EspressoPassStrengthReductionGeneralSliceToSlice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
