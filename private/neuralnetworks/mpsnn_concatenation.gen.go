// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenation] class.
var (
	_MPSNNConcatenationClass     MPSNNConcatenationClass
	_MPSNNConcatenationClassOnce sync.Once
)

func getMPSNNConcatenationClass() MPSNNConcatenationClass {
	_MPSNNConcatenationClassOnce.Do(func() {
		_MPSNNConcatenationClass = MPSNNConcatenationClass{class: objc.GetClass("MPSNNConcatenation")}
	})
	return _MPSNNConcatenationClass
}

// GetMPSNNConcatenationClass returns the class object for MPSNNConcatenation.
func GetMPSNNConcatenationClass() MPSNNConcatenationClass {
	return getMPSNNConcatenationClass()
}

type MPSNNConcatenationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationClass) Alloc() MPSNNConcatenation {
	rv := objc.Send[MPSNNConcatenation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNConcatenation.CopyToGradientStateSourceImageSourceStatesDestinationImage]
//   - [MPSNNConcatenation.CopyToGradientStateSourceImagesSourceStatesDestinationImage]
//   - [MPSNNConcatenation.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNConcatenation.EncodeBatchToCommandBufferSourceImagesDestinationImage]
//   - [MPSNNConcatenation.EncodeToCommandBufferSourceImagesDestinationImage]
//   - [MPSNNConcatenation.IsResultStateReusedAcrossBatch]
//   - [MPSNNConcatenation.ResultStateBatchForSourceImageSourceStatesDestinationImage]
//   - [MPSNNConcatenation.ResultStateBatchForSourceImagesSourceStatesDestinationImage]
//   - [MPSNNConcatenation.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSNNConcatenation.ResultStateForSourceImagesSourceStatesDestinationImage]
//   - [MPSNNConcatenation.TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNConcatenation.TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [MPSNNConcatenation.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNConcatenation.TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [MPSNNConcatenation.InitWithCoderDevice]
//   - [MPSNNConcatenation.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation
type MPSNNConcatenation struct {
	MPSCNNKernel
}

// MPSNNConcatenationFromID constructs a [MPSNNConcatenation] from an objc.ID.
func MPSNNConcatenationFromID(id objc.ID) MPSNNConcatenation {
	return MPSNNConcatenation{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNConcatenation implements IMPSNNConcatenation.
var _ IMPSNNConcatenation = MPSNNConcatenation{}





// An interface definition for the [MPSNNConcatenation] class.
//
// # Methods
//
//   - [IMPSNNConcatenation.CopyToGradientStateSourceImageSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.CopyToGradientStateSourceImagesSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNConcatenation.EncodeBatchToCommandBufferSourceImagesDestinationImage]
//   - [IMPSNNConcatenation.EncodeToCommandBufferSourceImagesDestinationImage]
//   - [IMPSNNConcatenation.IsResultStateReusedAcrossBatch]
//   - [IMPSNNConcatenation.ResultStateBatchForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.ResultStateBatchForSourceImagesSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.ResultStateForSourceImagesSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage]
//   - [IMPSNNConcatenation.InitWithCoderDevice]
//   - [IMPSNNConcatenation.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation
type IMPSNNConcatenation interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyToGradientStateSourceImageSourceStatesDestinationImage(state objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject)
	CopyToGradientStateSourceImagesSourceStatesDestinationImage(state objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject)
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, image objectivec.IObject)
	EncodeToCommandBufferSourceImagesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, image objectivec.IObject)
	IsResultStateReusedAcrossBatch() bool
	ResultStateBatchForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	ResultStateBatchForSourceImagesSourceStatesDestinationImage(images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject
	ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	ResultStateForSourceImagesSourceStatesDestinationImage(images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenation
	InitWithDevice(device objectivec.IObject) MPSNNConcatenation
}





// Init initializes the instance.
func (c MPSNNConcatenation) Init() MPSNNConcatenation {
	rv := objc.Send[MPSNNConcatenation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenation) Autorelease() MPSNNConcatenation {
	rv := objc.Send[MPSNNConcatenation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenation creates a new MPSNNConcatenation instance.
func NewMPSNNConcatenation() MPSNNConcatenation {
	class := getMPSNNConcatenationClass()
	rv := objc.Send[MPSNNConcatenation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/initWithCoder:device:
func NewConcatenationWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenation {
	instance := getMPSNNConcatenationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNConcatenationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/initWithDevice:
func NewConcatenationWithDevice(device objectivec.IObject) MPSNNConcatenation {
	instance := getMPSNNConcatenationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNConcatenationFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/copyToGradientState:sourceImage:sourceStates:destinationImage:
func (c MPSNNConcatenation) CopyToGradientStateSourceImageSourceStatesDestinationImage(state objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("copyToGradientState:sourceImage:sourceStates:destinationImage:"), state, image, states, image2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/copyToGradientState:sourceImages:sourceStates:destinationImage:
func (c MPSNNConcatenation) CopyToGradientStateSourceImagesSourceStatesDestinationImage(state objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("copyToGradientState:sourceImages:sourceStates:destinationImage:"), state, images, states, image)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (c MPSNNConcatenation) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/encodeBatchToCommandBuffer:sourceImages:destinationImage:
func (c MPSNNConcatenation) EncodeBatchToCommandBufferSourceImagesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, image objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationImage:"), buffer, images, image)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/encodeToCommandBuffer:sourceImages:destinationImage:
func (c MPSNNConcatenation) EncodeToCommandBufferSourceImagesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, image objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImages:destinationImage:"), buffer, images, image)
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/isResultStateReusedAcrossBatch
func (c MPSNNConcatenation) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/resultStateBatchForSourceImage:sourceStates:destinationImage:
func (c MPSNNConcatenation) ResultStateBatchForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/resultStateBatchForSourceImages:sourceStates:destinationImage:
func (c MPSNNConcatenation) ResultStateBatchForSourceImagesSourceStatesDestinationImage(images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateBatchForSourceImages:sourceStates:destinationImage:"), images, states, image)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/resultStateForSourceImage:sourceStates:destinationImage:
func (c MPSNNConcatenation) ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/resultStateForSourceImages:sourceStates:destinationImage:
func (c MPSNNConcatenation) ResultStateForSourceImagesSourceStatesDestinationImage(images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resultStateForSourceImages:sourceStates:destinationImage:"), images, states, image)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (c MPSNNConcatenation) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:
func (c MPSNNConcatenation) TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:"), buffer, images, states, image)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (c MPSNNConcatenation) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:
func (c MPSNNConcatenation) TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:"), buffer, images, states, image)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/initWithCoder:device:
func (c MPSNNConcatenation) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNConcatenation {
	rv := objc.Send[MPSNNConcatenation](c.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/initWithDevice:
func (c MPSNNConcatenation) InitWithDevice(device objectivec.IObject) MPSNNConcatenation {
	rv := objc.Send[MPSNNConcatenation](c.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenation/libraryInfo:
func (_MPSNNConcatenationClass MPSNNConcatenationClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNConcatenationClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}






















