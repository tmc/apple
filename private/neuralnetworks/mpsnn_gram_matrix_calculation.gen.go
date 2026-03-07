// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGramMatrixCalculation] class.
var (
	_MPSNNGramMatrixCalculationClass     MPSNNGramMatrixCalculationClass
	_MPSNNGramMatrixCalculationClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationClass() MPSNNGramMatrixCalculationClass {
	_MPSNNGramMatrixCalculationClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationClass = MPSNNGramMatrixCalculationClass{class: objc.GetClass("MPSNNGramMatrixCalculation")}
	})
	return _MPSNNGramMatrixCalculationClass
}

// GetMPSNNGramMatrixCalculationClass returns the class object for MPSNNGramMatrixCalculation.
func GetMPSNNGramMatrixCalculationClass() MPSNNGramMatrixCalculationClass {
	return getMPSNNGramMatrixCalculationClass()
}

type MPSNNGramMatrixCalculationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationClass) Alloc() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGramMatrixCalculation.Alpha]
//   - [MPSNNGramMatrixCalculation.SetAlpha]
//   - [MPSNNGramMatrixCalculation.CopyWithZoneDevice]
//   - [MPSNNGramMatrixCalculation.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSNNGramMatrixCalculation.IsResultStateReusedAcrossBatch]
//   - [MPSNNGramMatrixCalculation.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSNNGramMatrixCalculation.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNGramMatrixCalculation.InitWithCoderDevice]
//   - [MPSNNGramMatrixCalculation.InitWithDevice]
//   - [MPSNNGramMatrixCalculation.InitWithDeviceAlpha]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation
type MPSNNGramMatrixCalculation struct {
	MPSCNNKernel
}

// MPSNNGramMatrixCalculationFromID constructs a [MPSNNGramMatrixCalculation] from an objc.ID.
func MPSNNGramMatrixCalculationFromID(id objc.ID) MPSNNGramMatrixCalculation {
	return MPSNNGramMatrixCalculation{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNGramMatrixCalculation implements IMPSNNGramMatrixCalculation.
var _ IMPSNNGramMatrixCalculation = MPSNNGramMatrixCalculation{}





// An interface definition for the [MPSNNGramMatrixCalculation] class.
//
// # Methods
//
//   - [IMPSNNGramMatrixCalculation.Alpha]
//   - [IMPSNNGramMatrixCalculation.SetAlpha]
//   - [IMPSNNGramMatrixCalculation.CopyWithZoneDevice]
//   - [IMPSNNGramMatrixCalculation.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSNNGramMatrixCalculation.IsResultStateReusedAcrossBatch]
//   - [IMPSNNGramMatrixCalculation.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNGramMatrixCalculation.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNGramMatrixCalculation.InitWithCoderDevice]
//   - [IMPSNNGramMatrixCalculation.InitWithDevice]
//   - [IMPSNNGramMatrixCalculation.InitWithDeviceAlpha]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation
type IMPSNNGramMatrixCalculation interface {
	IMPSCNNKernel

	// Topic: Methods

	Alpha() float32
	SetAlpha(value float32)
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	IsResultStateReusedAcrossBatch() bool
	ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculation
	InitWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculation
	InitWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculation
}





// Init initializes the instance.
func (g MPSNNGramMatrixCalculation) Init() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculation) Autorelease() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculation creates a new MPSNNGramMatrixCalculation instance.
func NewMPSNNGramMatrixCalculation() MPSNNGramMatrixCalculation {
	class := getMPSNNGramMatrixCalculationClass()
	rv := objc.Send[MPSNNGramMatrixCalculation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithCoder:device:
func NewGramMatrixCalculationWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNGramMatrixCalculationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithDevice:
func NewGramMatrixCalculationWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGramMatrixCalculationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithDevice:alpha:
func NewGramMatrixCalculationWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return MPSNNGramMatrixCalculationFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/copyWithZone:device:
func (g MPSNNGramMatrixCalculation) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/destinationImageDescriptorForSourceImages:sourceStates:
func (g MPSNNGramMatrixCalculation) DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), images, states)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/isResultStateReusedAcrossBatch
func (g MPSNNGramMatrixCalculation) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/resultStateForSourceImage:sourceStates:destinationImage:
func (g MPSNNGramMatrixCalculation) ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (g MPSNNGramMatrixCalculation) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithCoder:device:
func (g MPSNNGramMatrixCalculation) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithDevice:
func (g MPSNNGramMatrixCalculation) InitWithDevice(device objectivec.IObject) MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/initWithDevice:alpha:
func (g MPSNNGramMatrixCalculation) InitWithDeviceAlpha(device objectivec.IObject, alpha float32) MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/libraryInfo:
func (_MPSNNGramMatrixCalculationClass MPSNNGramMatrixCalculationClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNGramMatrixCalculationClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculation/alpha
func (g MPSNNGramMatrixCalculation) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}
func (g MPSNNGramMatrixCalculation) SetAlpha(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAlpha:"), value)
}

















