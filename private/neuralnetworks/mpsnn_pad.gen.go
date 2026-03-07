// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPad] class.
var (
	_MPSNNPadClass     MPSNNPadClass
	_MPSNNPadClassOnce sync.Once
)

func getMPSNNPadClass() MPSNNPadClass {
	_MPSNNPadClassOnce.Do(func() {
		_MPSNNPadClass = MPSNNPadClass{class: objc.GetClass("MPSNNPad")}
	})
	return _MPSNNPadClass
}

// GetMPSNNPadClass returns the class object for MPSNNPad.
func GetMPSNNPadClass() MPSNNPadClass {
	return getMPSNNPadClass()
}

type MPSNNPadClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadClass) Alloc() MPSNNPad {
	rv := objc.Send[MPSNNPad](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPad.CopyWithZoneDevice]
//   - [MPSNNPad.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSNNPad.FillValue]
//   - [MPSNNPad.SetFillValue]
//   - [MPSNNPad.IsResultStateReusedAcrossBatch]
//   - [MPSNNPad.PaddingSizeAfter]
//   - [MPSNNPad.SetPaddingSizeAfter]
//   - [MPSNNPad.PaddingSizeBefore]
//   - [MPSNNPad.SetPaddingSizeBefore]
//   - [MPSNNPad.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSNNPad.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNPad.InitWithCoderDevice]
//   - [MPSNNPad.InitWithDevice]
//   - [MPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfter]
//   - [MPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad
type MPSNNPad struct {
	MPSCNNKernel
}

// MPSNNPadFromID constructs a [MPSNNPad] from an objc.ID.
func MPSNNPadFromID(id objc.ID) MPSNNPad {
	return MPSNNPad{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNPad implements IMPSNNPad.
var _ IMPSNNPad = MPSNNPad{}





// An interface definition for the [MPSNNPad] class.
//
// # Methods
//
//   - [IMPSNNPad.CopyWithZoneDevice]
//   - [IMPSNNPad.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSNNPad.FillValue]
//   - [IMPSNNPad.SetFillValue]
//   - [IMPSNNPad.IsResultStateReusedAcrossBatch]
//   - [IMPSNNPad.PaddingSizeAfter]
//   - [IMPSNNPad.SetPaddingSizeAfter]
//   - [IMPSNNPad.PaddingSizeBefore]
//   - [IMPSNNPad.SetPaddingSizeBefore]
//   - [IMPSNNPad.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNPad.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNPad.InitWithCoderDevice]
//   - [IMPSNNPad.InitWithDevice]
//   - [IMPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfter]
//   - [IMPSNNPad.InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad
type IMPSNNPad interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	FillValue() float32
	SetFillValue(value float32)
	IsResultStateReusedAcrossBatch() bool
	PaddingSizeAfter() objectivec.IObject
	SetPaddingSizeAfter(value objectivec.IObject)
	PaddingSizeBefore() objectivec.IObject
	SetPaddingSizeBefore(value objectivec.IObject)
	ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPad
	InitWithDevice(device objectivec.IObject) MPSNNPad
	InitWithDevicePaddingSizeBeforePaddingSizeAfter(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject) MPSNNPad
	InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject, array objectivec.IObject) MPSNNPad
}





// Init initializes the instance.
func (p MPSNNPad) Init() MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPad) Autorelease() MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPad creates a new MPSNNPad instance.
func NewMPSNNPad() MPSNNPad {
	class := getMPSNNPadClass()
	rv := objc.Send[MPSNNPad](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithCoder:device:
func NewPadWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNPadFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:
func NewPadWithDevice(device objectivec.IObject) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPadFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:paddingSizeBefore:paddingSizeAfter:
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfter(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:"), device, before, after)
	return MPSNNPadFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:
func NewPadWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject, array objectivec.IObject) MPSNNPad {
	instance := getMPSNNPadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:"), device, before, after, array)
	return MPSNNPadFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/copyWithZone:device:
func (p MPSNNPad) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/destinationImageDescriptorForSourceImages:sourceStates:
func (p MPSNNPad) DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), images, states)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/isResultStateReusedAcrossBatch
func (p MPSNNPad) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/resultStateForSourceImage:sourceStates:destinationImage:
func (p MPSNNPad) ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (p MPSNNPad) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithCoder:device:
func (p MPSNNPad) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:
func (p MPSNNPad) InitWithDevice(device objectivec.IObject) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:paddingSizeBefore:paddingSizeAfter:
func (p MPSNNPad) InitWithDevicePaddingSizeBeforePaddingSizeAfter(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:"), device, before, after)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:
func (p MPSNNPad) InitWithDevicePaddingSizeBeforePaddingSizeAfterFillValueArray(device objectivec.IObject, before objectivec.IObject, after objectivec.IObject, array objectivec.IObject) MPSNNPad {
	rv := objc.Send[MPSNNPad](p.ID, objc.Sel("initWithDevice:paddingSizeBefore:paddingSizeAfter:fillValueArray:"), device, before, after, array)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/libraryInfo:
func (_MPSNNPadClass MPSNNPadClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNPadClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/fillValue
func (p MPSNNPad) FillValue() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("fillValue"))
	return rv
}
func (p MPSNNPad) SetFillValue(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setFillValue:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/paddingSizeAfter
func (p MPSNNPad) PaddingSizeAfter() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("paddingSizeAfter"))
	return objectivec.Object{ID: rv}
}
func (p MPSNNPad) SetPaddingSizeAfter(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaddingSizeAfter:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPad/paddingSizeBefore
func (p MPSNNPad) PaddingSizeBefore() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("paddingSizeBefore"))
	return objectivec.Object{ID: rv}
}
func (p MPSNNPad) SetPaddingSizeBefore(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaddingSizeBefore:"), value)
}

















