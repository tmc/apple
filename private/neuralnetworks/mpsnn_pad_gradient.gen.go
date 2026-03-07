// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPadGradient] class.
var (
	_MPSNNPadGradientClass     MPSNNPadGradientClass
	_MPSNNPadGradientClassOnce sync.Once
)

func getMPSNNPadGradientClass() MPSNNPadGradientClass {
	_MPSNNPadGradientClassOnce.Do(func() {
		_MPSNNPadGradientClass = MPSNNPadGradientClass{class: objc.GetClass("MPSNNPadGradient")}
	})
	return _MPSNNPadGradientClass
}

// GetMPSNNPadGradientClass returns the class object for MPSNNPadGradient.
func GetMPSNNPadGradientClass() MPSNNPadGradientClass {
	return getMPSNNPadGradientClass()
}

type MPSNNPadGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadGradientClass) Alloc() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPadGradient.CopyWithZoneDevice]
//   - [MPSNNPadGradient.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSNNPadGradient.InitWithCoderDevice]
//   - [MPSNNPadGradient.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient
type MPSNNPadGradient struct {
	MPSCNNGradientKernel
}

// MPSNNPadGradientFromID constructs a [MPSNNPadGradient] from an objc.ID.
func MPSNNPadGradientFromID(id objc.ID) MPSNNPadGradient {
	return MPSNNPadGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}
// Ensure MPSNNPadGradient implements IMPSNNPadGradient.
var _ IMPSNNPadGradient = MPSNNPadGradient{}





// An interface definition for the [MPSNNPadGradient] class.
//
// # Methods
//
//   - [IMPSNNPadGradient.CopyWithZoneDevice]
//   - [IMPSNNPadGradient.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSNNPadGradient.InitWithCoderDevice]
//   - [IMPSNNPadGradient.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient
type IMPSNNPadGradient interface {
	IMPSCNNGradientKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPadGradient
	InitWithDevice(device objectivec.IObject) MPSNNPadGradient
}





// Init initializes the instance.
func (p MPSNNPadGradient) Init() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadGradient) Autorelease() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadGradient creates a new MPSNNPadGradient instance.
func NewMPSNNPadGradient() MPSNNPadGradient {
	class := getMPSNNPadGradientClass()
	rv := objc.Send[MPSNNPadGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/initWithCoder:device:
func NewPadGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPadGradient {
	instance := getMPSNNPadGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNPadGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/initWithDevice:
func NewPadGradientWithDevice(device objectivec.IObject) MPSNNPadGradient {
	instance := getMPSNNPadGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPadGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/copyWithZone:device:
func (p MPSNNPadGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/destinationImageDescriptorForSourceImages:sourceStates:
func (p MPSNNPadGradient) DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), images, states)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/initWithCoder:device:
func (p MPSNNPadGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/initWithDevice:
func (p MPSNNPadGradient) InitWithDevice(device objectivec.IObject) MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradient/libraryInfo:
func (_MPSNNPadGradientClass MPSNNPadGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNPadGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}






















