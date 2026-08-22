// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
)

// The class instance for the [EspressoTrainingConvolutionWeightsForMPS] class.
var (
	_EspressoTrainingConvolutionWeightsForMPSClass     EspressoTrainingConvolutionWeightsForMPSClass
	_EspressoTrainingConvolutionWeightsForMPSClassOnce sync.Once
)

func getEspressoTrainingConvolutionWeightsForMPSClass() EspressoTrainingConvolutionWeightsForMPSClass {
	_EspressoTrainingConvolutionWeightsForMPSClassOnce.Do(func() {
		_EspressoTrainingConvolutionWeightsForMPSClass = EspressoTrainingConvolutionWeightsForMPSClass{class: objc.GetClass("EspressoTrainingConvolutionWeightsForMPS")}
	})
	return _EspressoTrainingConvolutionWeightsForMPSClass
}

// GetEspressoTrainingConvolutionWeightsForMPSClass returns the class object for EspressoTrainingConvolutionWeightsForMPS.
func GetEspressoTrainingConvolutionWeightsForMPSClass() EspressoTrainingConvolutionWeightsForMPSClass {
	return getEspressoTrainingConvolutionWeightsForMPSClass()
}

type EspressoTrainingConvolutionWeightsForMPSClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoTrainingConvolutionWeightsForMPSClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoTrainingConvolutionWeightsForMPSClass) Alloc() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingConvolutionWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoTrainingConvolutionWeightsForMPS.BiasesBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.SetBiasesBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.WeightsBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.SetWeightsBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.WeightsLayout]
//   - [EspressoTrainingConvolutionWeightsForMPS.InitWithParamsForMode]
type EspressoTrainingConvolutionWeightsForMPS struct {
	EspressoConvolutionWeightsForMPS
}

// EspressoTrainingConvolutionWeightsForMPSFromID constructs a [EspressoTrainingConvolutionWeightsForMPS] from an objc.ID.
func EspressoTrainingConvolutionWeightsForMPSFromID(id objc.ID) EspressoTrainingConvolutionWeightsForMPS {
	return EspressoTrainingConvolutionWeightsForMPS{EspressoConvolutionWeightsForMPS: EspressoConvolutionWeightsForMPSFromID(id)}
}

// Ensure EspressoTrainingConvolutionWeightsForMPS implements IEspressoTrainingConvolutionWeightsForMPS.
var _ IEspressoTrainingConvolutionWeightsForMPS = EspressoTrainingConvolutionWeightsForMPS{}

// An interface definition for the [EspressoTrainingConvolutionWeightsForMPS] class.
//
// # Methods
//
//   - [IEspressoTrainingConvolutionWeightsForMPS.BiasesBuffer]
//   - [IEspressoTrainingConvolutionWeightsForMPS.SetBiasesBuffer]
//   - [IEspressoTrainingConvolutionWeightsForMPS.WeightsBuffer]
//   - [IEspressoTrainingConvolutionWeightsForMPS.SetWeightsBuffer]
//   - [IEspressoTrainingConvolutionWeightsForMPS.WeightsLayout]
//   - [IEspressoTrainingConvolutionWeightsForMPS.InitWithParamsForMode]
type IEspressoTrainingConvolutionWeightsForMPS interface {
	IEspressoConvolutionWeightsForMPS

	// Topic: Methods

	BiasesBuffer() unsafe.Pointer
	SetBiasesBuffer(value unsafe.Pointer)
	WeightsBuffer() unsafe.Pointer
	SetWeightsBuffer(value unsafe.Pointer)
	WeightsLayout() uint32
	InitWithParamsForMode(params appleneuralengine.ConvolutionUniforms, mode bool) EspressoTrainingConvolutionWeightsForMPS
}

// Init initializes the instance.
func (e EspressoTrainingConvolutionWeightsForMPS) Init() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoTrainingConvolutionWeightsForMPS) Autorelease() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoTrainingConvolutionWeightsForMPS creates a new EspressoTrainingConvolutionWeightsForMPS instance.
func NewEspressoTrainingConvolutionWeightsForMPS() EspressoTrainingConvolutionWeightsForMPS {
	class := getEspressoTrainingConvolutionWeightsForMPSClass()
	rv := objc.SendIfResponds[EspressoTrainingConvolutionWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoTrainingConvolutionWeightsForMPSWithParams(params appleneuralengine.ConvolutionUniforms) EspressoTrainingConvolutionWeightsForMPS {
	instance := getEspressoTrainingConvolutionWeightsForMPSClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoTrainingConvolutionWeightsForMPSFromID(rv)
}

func NewEspressoTrainingConvolutionWeightsForMPSWithParamsForMode(params appleneuralengine.ConvolutionUniforms, mode bool) EspressoTrainingConvolutionWeightsForMPS {
	instance := getEspressoTrainingConvolutionWeightsForMPSClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return EspressoTrainingConvolutionWeightsForMPSFromID(rv)
}

func (e EspressoTrainingConvolutionWeightsForMPS) WeightsLayout() uint32 {
	rv := objc.SendIfResponds[uint32](e.ID, objc.Sel("weightsLayout"))
	return rv
}
func (e EspressoTrainingConvolutionWeightsForMPS) InitWithParamsForMode(params appleneuralengine.ConvolutionUniforms, mode bool) EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.SendIfResponds[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return rv
}

func (e EspressoTrainingConvolutionWeightsForMPS) BiasesBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("biasesBuffer"))
	return rv
}
func (e EspressoTrainingConvolutionWeightsForMPS) SetBiasesBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setBiasesBuffer:"), value)
}
func (e EspressoTrainingConvolutionWeightsForMPS) WeightsBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("weightsBuffer"))
	return rv
}
func (e EspressoTrainingConvolutionWeightsForMPS) SetWeightsBuffer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setWeightsBuffer:"), value)
}
