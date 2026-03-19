// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoTrainingInnerProductWeightsForMPS] class.
var (
	_EspressoTrainingInnerProductWeightsForMPSClass     EspressoTrainingInnerProductWeightsForMPSClass
	_EspressoTrainingInnerProductWeightsForMPSClassOnce sync.Once
)

func getEspressoTrainingInnerProductWeightsForMPSClass() EspressoTrainingInnerProductWeightsForMPSClass {
	_EspressoTrainingInnerProductWeightsForMPSClassOnce.Do(func() {
		_EspressoTrainingInnerProductWeightsForMPSClass = EspressoTrainingInnerProductWeightsForMPSClass{class: objc.GetClass("EspressoTrainingInnerProductWeightsForMPS")}
	})
	return _EspressoTrainingInnerProductWeightsForMPSClass
}

// GetEspressoTrainingInnerProductWeightsForMPSClass returns the class object for EspressoTrainingInnerProductWeightsForMPS.
func GetEspressoTrainingInnerProductWeightsForMPSClass() EspressoTrainingInnerProductWeightsForMPSClass {
	return getEspressoTrainingInnerProductWeightsForMPSClass()
}

type EspressoTrainingInnerProductWeightsForMPSClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoTrainingInnerProductWeightsForMPSClass) Alloc() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.Send[EspressoTrainingInnerProductWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoTrainingInnerProductWeightsForMPS.BiasesBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.SetBiasesBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.WeightsBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.SetWeightsBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.InitWithParamsForMode]
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS
type EspressoTrainingInnerProductWeightsForMPS struct {
	EspressoInnerProductWeightsForMPS
}

// EspressoTrainingInnerProductWeightsForMPSFromID constructs a [EspressoTrainingInnerProductWeightsForMPS] from an objc.ID.
func EspressoTrainingInnerProductWeightsForMPSFromID(id objc.ID) EspressoTrainingInnerProductWeightsForMPS {
	return EspressoTrainingInnerProductWeightsForMPS{EspressoInnerProductWeightsForMPS: EspressoInnerProductWeightsForMPSFromID(id)}
}
// Ensure EspressoTrainingInnerProductWeightsForMPS implements IEspressoTrainingInnerProductWeightsForMPS.
var _ IEspressoTrainingInnerProductWeightsForMPS = EspressoTrainingInnerProductWeightsForMPS{}

// An interface definition for the [EspressoTrainingInnerProductWeightsForMPS] class.
//
// # Methods
//
//   - [IEspressoTrainingInnerProductWeightsForMPS.BiasesBuffer]
//   - [IEspressoTrainingInnerProductWeightsForMPS.SetBiasesBuffer]
//   - [IEspressoTrainingInnerProductWeightsForMPS.WeightsBuffer]
//   - [IEspressoTrainingInnerProductWeightsForMPS.SetWeightsBuffer]
//   - [IEspressoTrainingInnerProductWeightsForMPS.InitWithParamsForMode]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS
type IEspressoTrainingInnerProductWeightsForMPS interface {
	IEspressoInnerProductWeightsForMPS

	// Topic: Methods

	BiasesBuffer() objectivec.IObject
	SetBiasesBuffer(value objectivec.IObject)
	WeightsBuffer() objectivec.IObject
	SetWeightsBuffer(value objectivec.IObject)
	InitWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingInnerProductWeightsForMPS
}

// Init initializes the instance.
func (e EspressoTrainingInnerProductWeightsForMPS) Init() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.Send[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoTrainingInnerProductWeightsForMPS) Autorelease() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.Send[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoTrainingInnerProductWeightsForMPS creates a new EspressoTrainingInnerProductWeightsForMPS instance.
func NewEspressoTrainingInnerProductWeightsForMPS() EspressoTrainingInnerProductWeightsForMPS {
	class := getEspressoTrainingInnerProductWeightsForMPSClass()
	rv := objc.Send[EspressoTrainingInnerProductWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoInnerProductWeightsForMPS/initWithParams:
func NewEspressoTrainingInnerProductWeightsForMPSWithParams(params objectivec.IObject) EspressoTrainingInnerProductWeightsForMPS {
	instance := getEspressoTrainingInnerProductWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoTrainingInnerProductWeightsForMPSFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS/initWithParams:forMode:
func NewEspressoTrainingInnerProductWeightsForMPSWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingInnerProductWeightsForMPS {
	instance := getEspressoTrainingInnerProductWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return EspressoTrainingInnerProductWeightsForMPSFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS/initWithParams:forMode:
func (e EspressoTrainingInnerProductWeightsForMPS) InitWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.Send[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS/biasesBuffer
func (e EspressoTrainingInnerProductWeightsForMPS) BiasesBuffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("biasesBuffer"))
	return objectivec.Object{ID: rv}
}
func (e EspressoTrainingInnerProductWeightsForMPS) SetBiasesBuffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setBiasesBuffer:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingInnerProductWeightsForMPS/weightsBuffer
func (e EspressoTrainingInnerProductWeightsForMPS) WeightsBuffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("weightsBuffer"))
	return objectivec.Object{ID: rv}
}
func (e EspressoTrainingInnerProductWeightsForMPS) SetWeightsBuffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setWeightsBuffer:"), value)
}

