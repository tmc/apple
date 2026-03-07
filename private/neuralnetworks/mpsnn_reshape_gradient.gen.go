// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReshapeGradient] class.
var (
	_MPSNNReshapeGradientClass     MPSNNReshapeGradientClass
	_MPSNNReshapeGradientClassOnce sync.Once
)

func getMPSNNReshapeGradientClass() MPSNNReshapeGradientClass {
	_MPSNNReshapeGradientClassOnce.Do(func() {
		_MPSNNReshapeGradientClass = MPSNNReshapeGradientClass{class: objc.GetClass("MPSNNReshapeGradient")}
	})
	return _MPSNNReshapeGradientClass
}

// GetMPSNNReshapeGradientClass returns the class object for MPSNNReshapeGradient.
func GetMPSNNReshapeGradientClass() MPSNNReshapeGradientClass {
	return getMPSNNReshapeGradientClass()
}

type MPSNNReshapeGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeGradientClass) Alloc() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReshapeGradient.CopyWithZoneDevice]
//   - [MPSNNReshapeGradient.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [MPSNNReshapeGradient.InitWithCoderDevice]
//   - [MPSNNReshapeGradient.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient
type MPSNNReshapeGradient struct {
	MPSCNNGradientKernel
}

// MPSNNReshapeGradientFromID constructs a [MPSNNReshapeGradient] from an objc.ID.
func MPSNNReshapeGradientFromID(id objc.ID) MPSNNReshapeGradient {
	return MPSNNReshapeGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}
// Ensure MPSNNReshapeGradient implements IMPSNNReshapeGradient.
var _ IMPSNNReshapeGradient = MPSNNReshapeGradient{}





// An interface definition for the [MPSNNReshapeGradient] class.
//
// # Methods
//
//   - [IMPSNNReshapeGradient.CopyWithZoneDevice]
//   - [IMPSNNReshapeGradient.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [IMPSNNReshapeGradient.InitWithCoderDevice]
//   - [IMPSNNReshapeGradient.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient
type IMPSNNReshapeGradient interface {
	IMPSCNNGradientKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshapeGradient
	InitWithDevice(device objectivec.IObject) MPSNNReshapeGradient
}





// Init initializes the instance.
func (r MPSNNReshapeGradient) Init() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeGradient) Autorelease() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeGradient creates a new MPSNNReshapeGradient instance.
func NewMPSNNReshapeGradient() MPSNNReshapeGradient {
	class := getMPSNNReshapeGradientClass()
	rv := objc.Send[MPSNNReshapeGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/initWithCoder:device:
func NewReshapeGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshapeGradient {
	instance := getMPSNNReshapeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReshapeGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/initWithDevice:
func NewReshapeGradientWithDevice(device objectivec.IObject) MPSNNReshapeGradient {
	instance := getMPSNNReshapeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReshapeGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/copyWithZone:device:
func (r MPSNNReshapeGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:
func (r MPSNNReshapeGradient) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:"), images, states, method, offset, offset2, offset3)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/initWithCoder:device:
func (r MPSNNReshapeGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/initWithDevice:
func (r MPSNNReshapeGradient) InitWithDevice(device objectivec.IObject) MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradient/libraryInfo:
func (_MPSNNReshapeGradientClass MPSNNReshapeGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNReshapeGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}






















