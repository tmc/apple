// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
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

// Class returns the underlying Objective-C class pointer.
func (ec EspressoTrainingInnerProductWeightsForMPSClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoTrainingInnerProductWeightsForMPSClass) Alloc() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingInnerProductWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoTrainingInnerProductWeightsForMPS.BiasesBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.SetBiasesBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.WeightsBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.SetWeightsBuffer]
//   - [EspressoTrainingInnerProductWeightsForMPS.InitWithParamsForMode]
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
type IEspressoTrainingInnerProductWeightsForMPS interface {
	IEspressoInnerProductWeightsForMPS

	// Topic: Methods

	BiasesBuffer() unsafe.Pointer
	SetBiasesBuffer(value unsafe.Pointer)
	WeightsBuffer() unsafe.Pointer
	SetWeightsBuffer(value unsafe.Pointer)
	InitWithParamsForMode(params appleneuralengine.InnerProductUniforms, mode bool) EspressoTrainingInnerProductWeightsForMPS
}

// Init initializes the instance.
func (e EspressoTrainingInnerProductWeightsForMPS) Init() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoTrainingInnerProductWeightsForMPS) Autorelease() EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoTrainingInnerProductWeightsForMPS creates a new EspressoTrainingInnerProductWeightsForMPS instance.
func NewEspressoTrainingInnerProductWeightsForMPS() EspressoTrainingInnerProductWeightsForMPS {
	class := getEspressoTrainingInnerProductWeightsForMPSClass()
	rv := objc.SendIfResponds[EspressoTrainingInnerProductWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoTrainingInnerProductWeightsForMPSWithParams(params appleneuralengine.InnerProductUniforms) EspressoTrainingInnerProductWeightsForMPS {
	instance := getEspressoTrainingInnerProductWeightsForMPSClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoTrainingInnerProductWeightsForMPSFromID(rv)
}

func NewEspressoTrainingInnerProductWeightsForMPSWithParamsForMode(params appleneuralengine.InnerProductUniforms, mode bool) EspressoTrainingInnerProductWeightsForMPS {
	instance := getEspressoTrainingInnerProductWeightsForMPSClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return EspressoTrainingInnerProductWeightsForMPSFromID(rv)
}

func (e EspressoTrainingInnerProductWeightsForMPS) InitWithParamsForMode(params appleneuralengine.InnerProductUniforms, mode bool) EspressoTrainingInnerProductWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingInnerProductWeightsForMPS](e.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return rv
}

func (e EspressoTrainingInnerProductWeightsForMPS) BiasesBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("biasesBuffer"))
	return rv
}
func (e EspressoTrainingInnerProductWeightsForMPS) SetBiasesBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setBiasesBuffer:"), value)
}
func (e EspressoTrainingInnerProductWeightsForMPS) WeightsBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("weightsBuffer"))
	return rv
}
func (e EspressoTrainingInnerProductWeightsForMPS) SetWeightsBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setWeightsBuffer:"), value)
}
