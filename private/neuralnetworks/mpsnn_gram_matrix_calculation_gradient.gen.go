// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGramMatrixCalculationGradient] class.
var (
	_MPSNNGramMatrixCalculationGradientClass     MPSNNGramMatrixCalculationGradientClass
	_MPSNNGramMatrixCalculationGradientClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationGradientClass() MPSNNGramMatrixCalculationGradientClass {
	_MPSNNGramMatrixCalculationGradientClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationGradientClass = MPSNNGramMatrixCalculationGradientClass{class: objc.GetClass("MPSNNGramMatrixCalculationGradient")}
	})
	return _MPSNNGramMatrixCalculationGradientClass
}

// GetMPSNNGramMatrixCalculationGradientClass returns the class object for MPSNNGramMatrixCalculationGradient.
func GetMPSNNGramMatrixCalculationGradientClass() MPSNNGramMatrixCalculationGradientClass {
	return getMPSNNGramMatrixCalculationGradientClass()
}

type MPSNNGramMatrixCalculationGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationGradientClass) Alloc() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGramMatrixCalculationGradient.Alpha]
//   - [MPSNNGramMatrixCalculationGradient.SetAlpha]
//   - [MPSNNGramMatrixCalculationGradient.CopyWithZoneDevice]
//   - [MPSNNGramMatrixCalculationGradient.InitWithCoderDevice]
//   - [MPSNNGramMatrixCalculationGradient.InitWithDevice]
//   - [MPSNNGramMatrixCalculationGradient.InitWithDeviceAlpha]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient
type MPSNNGramMatrixCalculationGradient struct {
	MPSCNNGradientKernel
}

// MPSNNGramMatrixCalculationGradientFromID constructs a [MPSNNGramMatrixCalculationGradient] from an objc.ID.
func MPSNNGramMatrixCalculationGradientFromID(id objc.ID) MPSNNGramMatrixCalculationGradient {
	return MPSNNGramMatrixCalculationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}
// Ensure MPSNNGramMatrixCalculationGradient implements IMPSNNGramMatrixCalculationGradient.
var _ IMPSNNGramMatrixCalculationGradient = MPSNNGramMatrixCalculationGradient{}





// An interface definition for the [MPSNNGramMatrixCalculationGradient] class.
//
// # Methods
//
//   - [IMPSNNGramMatrixCalculationGradient.Alpha]
//   - [IMPSNNGramMatrixCalculationGradient.SetAlpha]
//   - [IMPSNNGramMatrixCalculationGradient.CopyWithZoneDevice]
//   - [IMPSNNGramMatrixCalculationGradient.InitWithCoderDevice]
//   - [IMPSNNGramMatrixCalculationGradient.InitWithDevice]
//   - [IMPSNNGramMatrixCalculationGradient.InitWithDeviceAlpha]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient
type IMPSNNGramMatrixCalculationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Methods

	Alpha() float32
	SetAlpha(value float32)
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculationGradient
	InitWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculationGradient
	InitWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradient
}





// Init initializes the instance.
func (g MPSNNGramMatrixCalculationGradient) Init() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationGradient) Autorelease() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationGradient creates a new MPSNNGramMatrixCalculationGradient instance.
func NewMPSNNGramMatrixCalculationGradient() MPSNNGramMatrixCalculationGradient {
	class := getMPSNNGramMatrixCalculationGradientClass()
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithCoder:device:
func NewGramMatrixCalculationGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithDevice:
func NewGramMatrixCalculationGradientWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithDevice:alpha:
func NewGramMatrixCalculationGradientWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/copyWithZone:device:
func (g MPSNNGramMatrixCalculationGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithCoder:device:
func (g MPSNNGramMatrixCalculationGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithDevice:
func (g MPSNNGramMatrixCalculationGradient) InitWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/initWithDevice:alpha:
func (g MPSNNGramMatrixCalculationGradient) InitWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/libraryInfo:
func (_MPSNNGramMatrixCalculationGradientClass MPSNNGramMatrixCalculationGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNGramMatrixCalculationGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradient/alpha
func (g MPSNNGramMatrixCalculationGradient) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}
func (g MPSNNGramMatrixCalculationGradient) SetAlpha(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAlpha:"), value)
}

















