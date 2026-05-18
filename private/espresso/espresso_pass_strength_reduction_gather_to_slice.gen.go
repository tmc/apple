// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassStrengthReductionGatherToSlice] class.
var (
	_EspressoPassStrengthReductionGatherToSliceClass     EspressoPassStrengthReductionGatherToSliceClass
	_EspressoPassStrengthReductionGatherToSliceClassOnce sync.Once
)

func getEspressoPassStrengthReductionGatherToSliceClass() EspressoPassStrengthReductionGatherToSliceClass {
	_EspressoPassStrengthReductionGatherToSliceClassOnce.Do(func() {
		_EspressoPassStrengthReductionGatherToSliceClass = EspressoPassStrengthReductionGatherToSliceClass{class: objc.GetClass("EspressoPass_strength_reduction_gather_to_slice")}
	})
	return _EspressoPassStrengthReductionGatherToSliceClass
}

// GetEspressoPassStrengthReductionGatherToSliceClass returns the class object for EspressoPass_strength_reduction_gather_to_slice.
func GetEspressoPassStrengthReductionGatherToSliceClass() EspressoPassStrengthReductionGatherToSliceClass {
	return getEspressoPassStrengthReductionGatherToSliceClass()
}

type EspressoPassStrengthReductionGatherToSliceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassStrengthReductionGatherToSliceClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassStrengthReductionGatherToSliceClass) Alloc() EspressoPassStrengthReductionGatherToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGatherToSlice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_slice
type EspressoPassStrengthReductionGatherToSlice struct {
	EspressoCustomPass
}

// EspressoPassStrengthReductionGatherToSliceFromID constructs a [EspressoPassStrengthReductionGatherToSlice] from an objc.ID.
func EspressoPassStrengthReductionGatherToSliceFromID(id objc.ID) EspressoPassStrengthReductionGatherToSlice {
	return EspressoPassStrengthReductionGatherToSlice{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_strength_reduction_gather_to_sliceFromID is an alias for [EspressoPassStrengthReductionGatherToSliceFromID] for cross-framework compatibility.
func EspressoPass_strength_reduction_gather_to_sliceFromID(id objc.ID) EspressoPassStrengthReductionGatherToSlice {
	return EspressoPassStrengthReductionGatherToSliceFromID(id)
}

// Ensure EspressoPassStrengthReductionGatherToSlice implements IEspressoPassStrengthReductionGatherToSlice.
var _ IEspressoPassStrengthReductionGatherToSlice = EspressoPassStrengthReductionGatherToSlice{}

// An interface definition for the [EspressoPassStrengthReductionGatherToSlice] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_strength_reduction_gather_to_slice
type IEspressoPassStrengthReductionGatherToSlice interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassStrengthReductionGatherToSlice) Init() EspressoPassStrengthReductionGatherToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGatherToSlice](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassStrengthReductionGatherToSlice) Autorelease() EspressoPassStrengthReductionGatherToSlice {
	rv := objc.Send[EspressoPassStrengthReductionGatherToSlice](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassStrengthReductionGatherToSlice creates a new EspressoPassStrengthReductionGatherToSlice instance.
func NewEspressoPassStrengthReductionGatherToSlice() EspressoPassStrengthReductionGatherToSlice {
	class := getEspressoPassStrengthReductionGatherToSliceClass()
	rv := objc.Send[EspressoPassStrengthReductionGatherToSlice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
