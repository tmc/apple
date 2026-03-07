// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenationGradient] class.
var (
	_MPSNNConcatenationGradientClass     MPSNNConcatenationGradientClass
	_MPSNNConcatenationGradientClassOnce sync.Once
)

func getMPSNNConcatenationGradientClass() MPSNNConcatenationGradientClass {
	_MPSNNConcatenationGradientClassOnce.Do(func() {
		_MPSNNConcatenationGradientClass = MPSNNConcatenationGradientClass{class: objc.GetClass("MPSNNConcatenationGradient")}
	})
	return _MPSNNConcatenationGradientClass
}

// GetMPSNNConcatenationGradientClass returns the class object for MPSNNConcatenationGradient.
func GetMPSNNConcatenationGradientClass() MPSNNConcatenationGradientClass {
	return getMPSNNConcatenationGradientClass()
}

type MPSNNConcatenationGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationGradientClass) Alloc() MPSNNConcatenationGradient {
	rv := objc.Send[MPSNNConcatenationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNConcatenationGradient.CopyWithZoneDevice]
//   - [MPSNNConcatenationGradient.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [MPSNNConcatenationGradient.InitWithCoderDevice]
//   - [MPSNNConcatenationGradient.InitWithDeviceSourceIndex]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient
type MPSNNConcatenationGradient struct {
	MPSCNNGradientKernel
}

// MPSNNConcatenationGradientFromID constructs a [MPSNNConcatenationGradient] from an objc.ID.
func MPSNNConcatenationGradientFromID(id objc.ID) MPSNNConcatenationGradient {
	return MPSNNConcatenationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}
// Ensure MPSNNConcatenationGradient implements IMPSNNConcatenationGradient.
var _ IMPSNNConcatenationGradient = MPSNNConcatenationGradient{}





// An interface definition for the [MPSNNConcatenationGradient] class.
//
// # Methods
//
//   - [IMPSNNConcatenationGradient.CopyWithZoneDevice]
//   - [IMPSNNConcatenationGradient.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [IMPSNNConcatenationGradient.InitWithCoderDevice]
//   - [IMPSNNConcatenationGradient.InitWithDeviceSourceIndex]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient
type IMPSNNConcatenationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenationGradient
	InitWithDeviceSourceIndex(device objectivec.IObject, index uint64) MPSNNConcatenationGradient
}





// Init initializes the instance.
func (c MPSNNConcatenationGradient) Init() MPSNNConcatenationGradient {
	rv := objc.Send[MPSNNConcatenationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationGradient) Autorelease() MPSNNConcatenationGradient {
	rv := objc.Send[MPSNNConcatenationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationGradient creates a new MPSNNConcatenationGradient instance.
func NewMPSNNConcatenationGradient() MPSNNConcatenationGradient {
	class := getMPSNNConcatenationGradientClass()
	rv := objc.Send[MPSNNConcatenationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/initWithCoder:device:
func NewConcatenationGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenationGradient {
	instance := getMPSNNConcatenationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNConcatenationGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/initWithDevice:sourceIndex:
func NewConcatenationGradientWithDeviceSourceIndex(device objectivec.IObject, index uint64) MPSNNConcatenationGradient {
	instance := getMPSNNConcatenationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceIndex:"), device, index)
	return MPSNNConcatenationGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/copyWithZone:device:
func (c MPSNNConcatenationGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:
func (c MPSNNConcatenationGradient) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:"), images, states, method, offset, offset2, offset3)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/initWithCoder:device:
func (c MPSNNConcatenationGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenationGradient {
	rv := objc.Send[MPSNNConcatenationGradient](c.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/initWithDevice:sourceIndex:
func (c MPSNNConcatenationGradient) InitWithDeviceSourceIndex(device objectivec.IObject, index uint64) MPSNNConcatenationGradient {
	rv := objc.Send[MPSNNConcatenationGradient](c.ID, objc.Sel("initWithDevice:sourceIndex:"), device, index)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradient/libraryInfo:
func (_MPSNNConcatenationGradientClass MPSNNConcatenationGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNConcatenationGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}






















