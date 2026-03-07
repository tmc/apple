// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (ec EspressoTrainingConvolutionWeightsForMPSClass) Alloc() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.Send[EspressoTrainingConvolutionWeightsForMPS](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [EspressoTrainingConvolutionWeightsForMPS.BiasesBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.SetBiasesBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.WeightsBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.SetWeightsBuffer]
//   - [EspressoTrainingConvolutionWeightsForMPS.WeightsLayout]
//   - [EspressoTrainingConvolutionWeightsForMPS.InitWithParamsForMode]
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS
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
//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS
type IEspressoTrainingConvolutionWeightsForMPS interface {
	IEspressoConvolutionWeightsForMPS

	// Topic: Methods

	BiasesBuffer() unsafe.Pointer
	SetBiasesBuffer(value unsafe.Pointer)
	WeightsBuffer() unsafe.Pointer
	SetWeightsBuffer(value unsafe.Pointer)
	WeightsLayout() uint32
	InitWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingConvolutionWeightsForMPS
}





// Init initializes the instance.
func (e EspressoTrainingConvolutionWeightsForMPS) Init() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.Send[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoTrainingConvolutionWeightsForMPS) Autorelease() EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.Send[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoTrainingConvolutionWeightsForMPS creates a new EspressoTrainingConvolutionWeightsForMPS instance.
func NewEspressoTrainingConvolutionWeightsForMPS() EspressoTrainingConvolutionWeightsForMPS {
	class := getEspressoTrainingConvolutionWeightsForMPSClass()
	rv := objc.Send[EspressoTrainingConvolutionWeightsForMPS](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Espresso/EspressoConvolutionWeightsForMPS/initWithParams:
func NewEspressoTrainingConvolutionWeightsForMPSWithParams(params objectivec.IObject) EspressoTrainingConvolutionWeightsForMPS {
	instance := getEspressoTrainingConvolutionWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return EspressoTrainingConvolutionWeightsForMPSFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS/initWithParams:forMode:
func NewEspressoTrainingConvolutionWeightsForMPSWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingConvolutionWeightsForMPS {
	instance := getEspressoTrainingConvolutionWeightsForMPSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return EspressoTrainingConvolutionWeightsForMPSFromID(rv)
}







// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS/weightsLayout
func (e EspressoTrainingConvolutionWeightsForMPS) WeightsLayout() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("weightsLayout"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS/initWithParams:forMode:
func (e EspressoTrainingConvolutionWeightsForMPS) InitWithParamsForMode(params objectivec.IObject, mode bool) EspressoTrainingConvolutionWeightsForMPS {
	rv := objc.Send[EspressoTrainingConvolutionWeightsForMPS](e.ID, objc.Sel("initWithParams:forMode:"), params, mode)
	return rv
}












// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS/biasesBuffer
func (e EspressoTrainingConvolutionWeightsForMPS) BiasesBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("biasesBuffer"))
	return rv
}
func (e EspressoTrainingConvolutionWeightsForMPS) SetBiasesBuffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setBiasesBuffer:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoTrainingConvolutionWeightsForMPS/weightsBuffer
func (e EspressoTrainingConvolutionWeightsForMPS) WeightsBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("weightsBuffer"))
	return rv
}
func (e EspressoTrainingConvolutionWeightsForMPS) SetWeightsBuffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setWeightsBuffer:"), value)
}

















