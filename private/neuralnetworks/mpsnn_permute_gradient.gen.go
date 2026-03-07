// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPermuteGradient] class.
var (
	_MPSNNPermuteGradientClass     MPSNNPermuteGradientClass
	_MPSNNPermuteGradientClassOnce sync.Once
)

func getMPSNNPermuteGradientClass() MPSNNPermuteGradientClass {
	_MPSNNPermuteGradientClassOnce.Do(func() {
		_MPSNNPermuteGradientClass = MPSNNPermuteGradientClass{class: objc.GetClass("MPSNNPermuteGradient")}
	})
	return _MPSNNPermuteGradientClass
}

// GetMPSNNPermuteGradientClass returns the class object for MPSNNPermuteGradient.
func GetMPSNNPermuteGradientClass() MPSNNPermuteGradientClass {
	return getMPSNNPermuteGradientClass()
}

type MPSNNPermuteGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPermuteGradientClass) Alloc() MPSNNPermuteGradient {
	rv := objc.Send[MPSNNPermuteGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPermuteGradient.CopyWithZoneDevice]
//   - [MPSNNPermuteGradient.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages]
//   - [MPSNNPermuteGradient.EncodeBatchToCommandEncoderCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages]
//   - [MPSNNPermuteGradient.InitWithCoderDevice]
//   - [MPSNNPermuteGradient.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient
type MPSNNPermuteGradient struct {
	MPSCNNGradientKernel
}

// MPSNNPermuteGradientFromID constructs a [MPSNNPermuteGradient] from an objc.ID.
func MPSNNPermuteGradientFromID(id objc.ID) MPSNNPermuteGradient {
	return MPSNNPermuteGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}
// Ensure MPSNNPermuteGradient implements IMPSNNPermuteGradient.
var _ IMPSNNPermuteGradient = MPSNNPermuteGradient{}





// An interface definition for the [MPSNNPermuteGradient] class.
//
// # Methods
//
//   - [IMPSNNPermuteGradient.CopyWithZoneDevice]
//   - [IMPSNNPermuteGradient.EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages]
//   - [IMPSNNPermuteGradient.EncodeBatchToCommandEncoderCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages]
//   - [IMPSNNPermuteGradient.InitWithCoderDevice]
//   - [IMPSNNPermuteGradient.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient
type IMPSNNPermuteGradient interface {
	IMPSCNNGradientKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, images2 objectivec.IObject, states objectivec.IObject, images3 objectivec.IObject)
	EncodeBatchToCommandEncoderCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, images2 objectivec.IObject, states objectivec.IObject, images3 objectivec.IObject)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermuteGradient
	InitWithDevice(device objectivec.IObject) MPSNNPermuteGradient
}





// Init initializes the instance.
func (p MPSNNPermuteGradient) Init() MPSNNPermuteGradient {
	rv := objc.Send[MPSNNPermuteGradient](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPermuteGradient) Autorelease() MPSNNPermuteGradient {
	rv := objc.Send[MPSNNPermuteGradient](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPermuteGradient creates a new MPSNNPermuteGradient instance.
func NewMPSNNPermuteGradient() MPSNNPermuteGradient {
	class := getMPSNNPermuteGradientClass()
	rv := objc.Send[MPSNNPermuteGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/initWithCoder:device:
func NewPermuteGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermuteGradient {
	instance := getMPSNNPermuteGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNPermuteGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/initWithDevice:
func NewPermuteGradientWithDevice(device objectivec.IObject) MPSNNPermuteGradient {
	instance := getMPSNNPermuteGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPermuteGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/copyWithZone:device:
func (p MPSNNPermuteGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/encodeBatchToCommandBuffer:primaryImages:secondaryImages:inStates:destinationImages:
func (p MPSNNPermuteGradient) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, images2 objectivec.IObject, states objectivec.IObject, images3 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:inStates:destinationImages:"), buffer, images, images2, states, images3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/encodeBatchToCommandEncoder:commandBuffer:primaryImages:secondaryImages:inStates:destinationImages:
func (p MPSNNPermuteGradient) EncodeBatchToCommandEncoderCommandBufferPrimaryImagesSecondaryImagesInStatesDestinationImages(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, images2 objectivec.IObject, states objectivec.IObject, images3 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeBatchToCommandEncoder:commandBuffer:primaryImages:secondaryImages:inStates:destinationImages:"), encoder, buffer, images, images2, states, images3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/initWithCoder:device:
func (p MPSNNPermuteGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermuteGradient {
	rv := objc.Send[MPSNNPermuteGradient](p.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/initWithDevice:
func (p MPSNNPermuteGradient) InitWithDevice(device objectivec.IObject) MPSNNPermuteGradient {
	rv := objc.Send[MPSNNPermuteGradient](p.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradient/libraryInfo:
func (_MPSNNPermuteGradientClass MPSNNPermuteGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNPermuteGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}






















