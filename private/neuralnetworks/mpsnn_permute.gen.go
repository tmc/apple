// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPermute] class.
var (
	_MPSNNPermuteClass     MPSNNPermuteClass
	_MPSNNPermuteClassOnce sync.Once
)

func getMPSNNPermuteClass() MPSNNPermuteClass {
	_MPSNNPermuteClassOnce.Do(func() {
		_MPSNNPermuteClass = MPSNNPermuteClass{class: objc.GetClass("MPSNNPermute")}
	})
	return _MPSNNPermuteClass
}

// GetMPSNNPermuteClass returns the class object for MPSNNPermute.
func GetMPSNNPermuteClass() MPSNNPermuteClass {
	return getMPSNNPermuteClass()
}

type MPSNNPermuteClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPermuteClass) Alloc() MPSNNPermute {
	rv := objc.Send[MPSNNPermute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPermute.AppendBatchBarrier]
//   - [MPSNNPermute.CopyWithZoneDevice]
//   - [MPSNNPermute.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNPermute.DimensionOrder]
//   - [MPSNNPermute.SetDimensionOrder]
//   - [MPSNNPermute.EncodeBatchInternalToCommandEncoderCommandBufferSourceImagesInStatesDestinationImagesSrcSizeDestSizeTestClipRectTestMaxClipRect]
//   - [MPSNNPermute.EncodeBatchToCommandBufferSourceImagesInStatesDestinationImages]
//   - [MPSNNPermute.EncodeBatchToCommandEncoderCommandBufferSourceImagesInStatesDestinationImages]
//   - [MPSNNPermute.IsResultStateReusedAcrossBatch]
//   - [MPSNNPermute.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSNNPermute.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNPermute.InitWithCoderDevice]
//   - [MPSNNPermute.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute
type MPSNNPermute struct {
	MPSCNNKernel
}

// MPSNNPermuteFromID constructs a [MPSNNPermute] from an objc.ID.
func MPSNNPermuteFromID(id objc.ID) MPSNNPermute {
	return MPSNNPermute{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNPermute implements IMPSNNPermute.
var _ IMPSNNPermute = MPSNNPermute{}





// An interface definition for the [MPSNNPermute] class.
//
// # Methods
//
//   - [IMPSNNPermute.AppendBatchBarrier]
//   - [IMPSNNPermute.CopyWithZoneDevice]
//   - [IMPSNNPermute.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNPermute.DimensionOrder]
//   - [IMPSNNPermute.SetDimensionOrder]
//   - [IMPSNNPermute.EncodeBatchInternalToCommandEncoderCommandBufferSourceImagesInStatesDestinationImagesSrcSizeDestSizeTestClipRectTestMaxClipRect]
//   - [IMPSNNPermute.EncodeBatchToCommandBufferSourceImagesInStatesDestinationImages]
//   - [IMPSNNPermute.EncodeBatchToCommandEncoderCommandBufferSourceImagesInStatesDestinationImages]
//   - [IMPSNNPermute.IsResultStateReusedAcrossBatch]
//   - [IMPSNNPermute.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNPermute.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNPermute.InitWithCoderDevice]
//   - [IMPSNNPermute.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute
type IMPSNNPermute interface {
	IMPSCNNKernel

	// Topic: Methods

	AppendBatchBarrier() bool
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	DimensionOrder() objectivec.IObject
	SetDimensionOrder(value objectivec.IObject)
	EncodeBatchInternalToCommandEncoderCommandBufferSourceImagesInStatesDestinationImagesSrcSizeDestSizeTestClipRectTestMaxClipRect(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, size objectivec.IObject, size2 objectivec.IObject, rect objectivec.IObject, rect2 objectivec.IObject)
	EncodeBatchToCommandBufferSourceImagesInStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject)
	EncodeBatchToCommandEncoderCommandBufferSourceImagesInStatesDestinationImages(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject)
	IsResultStateReusedAcrossBatch() bool
	ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermute
	InitWithDevice(device objectivec.IObject) MPSNNPermute
}





// Init initializes the instance.
func (p MPSNNPermute) Init() MPSNNPermute {
	rv := objc.Send[MPSNNPermute](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPermute) Autorelease() MPSNNPermute {
	rv := objc.Send[MPSNNPermute](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPermute creates a new MPSNNPermute instance.
func NewMPSNNPermute() MPSNNPermute {
	class := getMPSNNPermuteClass()
	rv := objc.Send[MPSNNPermute](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/initWithCoder:device:
func NewPermuteWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermute {
	instance := getMPSNNPermuteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNPermuteFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/initWithDevice:
func NewPermuteWithDevice(device objectivec.IObject) MPSNNPermute {
	instance := getMPSNNPermuteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPermuteFromID(rv)
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/appendBatchBarrier
func (p MPSNNPermute) AppendBatchBarrier() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("appendBatchBarrier"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/copyWithZone:device:
func (p MPSNNPermute) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (p MPSNNPermute) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/encodeBatchInternalToCommandEncoder:commandBuffer:sourceImages:inStates:destinationImages:srcSize:destSize:testClipRect:testMaxClipRect:
func (p MPSNNPermute) EncodeBatchInternalToCommandEncoderCommandBufferSourceImagesInStatesDestinationImagesSrcSizeDestSizeTestClipRectTestMaxClipRect(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, size objectivec.IObject, size2 objectivec.IObject, rect objectivec.IObject, rect2 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeBatchInternalToCommandEncoder:commandBuffer:sourceImages:inStates:destinationImages:srcSize:destSize:testClipRect:testMaxClipRect:"), encoder, buffer, images, states, images2, size, size2, rect, rect2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/encodeBatchToCommandBuffer:sourceImages:inStates:destinationImages:
func (p MPSNNPermute) EncodeBatchToCommandBufferSourceImagesInStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:inStates:destinationImages:"), buffer, images, states, images2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/encodeBatchToCommandEncoder:commandBuffer:sourceImages:inStates:destinationImages:
func (p MPSNNPermute) EncodeBatchToCommandEncoderCommandBufferSourceImagesInStatesDestinationImages(encoder objectivec.IObject, buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeBatchToCommandEncoder:commandBuffer:sourceImages:inStates:destinationImages:"), encoder, buffer, images, states, images2)
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/isResultStateReusedAcrossBatch
func (p MPSNNPermute) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/resultStateForSourceImage:sourceStates:destinationImage:
func (p MPSNNPermute) ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (p MPSNNPermute) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/initWithCoder:device:
func (p MPSNNPermute) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNPermute {
	rv := objc.Send[MPSNNPermute](p.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/initWithDevice:
func (p MPSNNPermute) InitWithDevice(device objectivec.IObject) MPSNNPermute {
	rv := objc.Send[MPSNNPermute](p.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/libraryInfo:
func (_MPSNNPermuteClass MPSNNPermuteClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNPermuteClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermute/dimensionOrder
func (p MPSNNPermute) DimensionOrder() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dimensionOrder"))
	return objectivec.Object{ID: rv}
}
func (p MPSNNPermute) SetDimensionOrder(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setDimensionOrder:"), value)
}

















