// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNForwardLoss] class.
var (
	_MPSNNForwardLossClass     MPSNNForwardLossClass
	_MPSNNForwardLossClassOnce sync.Once
)

func getMPSNNForwardLossClass() MPSNNForwardLossClass {
	_MPSNNForwardLossClassOnce.Do(func() {
		_MPSNNForwardLossClass = MPSNNForwardLossClass{class: objc.GetClass("MPSNNForwardLoss")}
	})
	return _MPSNNForwardLossClass
}

// GetMPSNNForwardLossClass returns the class object for MPSNNForwardLoss.
func GetMPSNNForwardLossClass() MPSNNForwardLossClass {
	return getMPSNNForwardLossClass()
}

type MPSNNForwardLossClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNForwardLossClass) Alloc() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNForwardLoss.CopyWithZoneDevice]
//   - [MPSNNForwardLoss.Delta]
//   - [MPSNNForwardLoss.SetDelta]
//   - [MPSNNForwardLoss.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages]
//   - [MPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary]
//   - [MPSNNForwardLoss.Epsilon]
//   - [MPSNNForwardLoss.SetEpsilon]
//   - [MPSNNForwardLoss.IsResultStateReusedAcrossBatch]
//   - [MPSNNForwardLoss.LabelSmoothing]
//   - [MPSNNForwardLoss.SetLabelSmoothing]
//   - [MPSNNForwardLoss.LossType]
//   - [MPSNNForwardLoss.MaxBatchSize]
//   - [MPSNNForwardLoss.NumberOfClasses]
//   - [MPSNNForwardLoss.ReduceAcrossBatch]
//   - [MPSNNForwardLoss.ReductionType]
//   - [MPSNNForwardLoss.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [MPSNNForwardLoss.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [MPSNNForwardLoss.Weight]
//   - [MPSNNForwardLoss.SetWeight]
//   - [MPSNNForwardLoss.InitWithCoderDevice]
//   - [MPSNNForwardLoss.InitWithDevice]
//   - [MPSNNForwardLoss.InitWithDeviceLossDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss
type MPSNNForwardLoss struct {
	MPSCNNKernel
}

// MPSNNForwardLossFromID constructs a [MPSNNForwardLoss] from an objc.ID.
func MPSNNForwardLossFromID(id objc.ID) MPSNNForwardLoss {
	return MPSNNForwardLoss{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNForwardLoss implements IMPSNNForwardLoss.
var _ IMPSNNForwardLoss = MPSNNForwardLoss{}





// An interface definition for the [MPSNNForwardLoss] class.
//
// # Methods
//
//   - [IMPSNNForwardLoss.CopyWithZoneDevice]
//   - [IMPSNNForwardLoss.Delta]
//   - [IMPSNNForwardLoss.SetDelta]
//   - [IMPSNNForwardLoss.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages]
//   - [IMPSNNForwardLoss.EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary]
//   - [IMPSNNForwardLoss.Epsilon]
//   - [IMPSNNForwardLoss.SetEpsilon]
//   - [IMPSNNForwardLoss.IsResultStateReusedAcrossBatch]
//   - [IMPSNNForwardLoss.LabelSmoothing]
//   - [IMPSNNForwardLoss.SetLabelSmoothing]
//   - [IMPSNNForwardLoss.LossType]
//   - [IMPSNNForwardLoss.MaxBatchSize]
//   - [IMPSNNForwardLoss.NumberOfClasses]
//   - [IMPSNNForwardLoss.ReduceAcrossBatch]
//   - [IMPSNNForwardLoss.ReductionType]
//   - [IMPSNNForwardLoss.ResultStateForSourceImageSourceStatesDestinationImage]
//   - [IMPSNNForwardLoss.TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage]
//   - [IMPSNNForwardLoss.Weight]
//   - [IMPSNNForwardLoss.SetWeight]
//   - [IMPSNNForwardLoss.InitWithCoderDevice]
//   - [IMPSNNForwardLoss.InitWithDevice]
//   - [IMPSNNForwardLoss.InitWithDeviceLossDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss
type IMPSNNForwardLoss interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	Delta() float32
	SetDelta(value float32)
	DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject)
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(buffer objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states []objectivec.IObject, temporary bool) objectivec.IObject
	Epsilon() float32
	SetEpsilon(value float32)
	IsResultStateReusedAcrossBatch() bool
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() uint32
	MaxBatchSize() uint64
	NumberOfClasses() uint64
	ReduceAcrossBatch() bool
	ReductionType() int
	ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject
	Weight() float32
	SetWeight(value float32)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNForwardLoss
	InitWithDevice(device objectivec.IObject) MPSNNForwardLoss
	InitWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLoss
}





// Init initializes the instance.
func (f MPSNNForwardLoss) Init() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSNNForwardLoss) Autorelease() MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNForwardLoss creates a new MPSNNForwardLoss instance.
func NewMPSNNForwardLoss() MPSNNForwardLoss {
	class := getMPSNNForwardLossClass()
	rv := objc.Send[MPSNNForwardLoss](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithCoder:device:
func NewForwardLossWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNForwardLossFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithDevice:
func NewForwardLossWithDevice(device objectivec.IObject) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNForwardLossFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithDevice:lossDescriptor:
func NewForwardLossWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLoss {
	instance := getMPSNNForwardLossClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, descriptor)
	return MPSNNForwardLossFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/copyWithZone:device:
func (f MPSNNForwardLoss) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/destinationImageDescriptorForSourceImages:sourceStates:
func (f MPSNNForwardLoss) DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), images, states)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:
func (f MPSNNForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(buffer objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:"), buffer, images, labels, weights, states, images2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:
func (f MPSNNForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(buffer objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states []objectivec.IObject, temporary bool) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), buffer, images, labels, weights, objectivec.IObjectSliceToNSArray(states), temporary)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/isResultStateReusedAcrossBatch
func (f MPSNNForwardLoss) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/maxBatchSize
func (f MPSNNForwardLoss) MaxBatchSize() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("maxBatchSize"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/resultStateForSourceImage:sourceStates:destinationImage:
func (f MPSNNForwardLoss) ResultStateForSourceImageSourceStatesDestinationImage(image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:
func (f MPSNNForwardLoss) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(buffer objectivec.IObject, image objectivec.IObject, states objectivec.IObject, image2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), buffer, image, states, image2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithCoder:device:
func (f MPSNNForwardLoss) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithDevice:
func (f MPSNNForwardLoss) InitWithDevice(device objectivec.IObject) MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/initWithDevice:lossDescriptor:
func (f MPSNNForwardLoss) InitWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLoss {
	rv := objc.Send[MPSNNForwardLoss](f.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, descriptor)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/libraryInfo:
func (_MPSNNForwardLossClass MPSNNForwardLossClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNForwardLossClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/delta
func (f MPSNNForwardLoss) Delta() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("delta"))
	return rv
}
func (f MPSNNForwardLoss) SetDelta(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelta:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/epsilon
func (f MPSNNForwardLoss) Epsilon() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("epsilon"))
	return rv
}
func (f MPSNNForwardLoss) SetEpsilon(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setEpsilon:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/labelSmoothing
func (f MPSNNForwardLoss) LabelSmoothing() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("labelSmoothing"))
	return rv
}
func (f MPSNNForwardLoss) SetLabelSmoothing(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setLabelSmoothing:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/lossType
func (f MPSNNForwardLoss) LossType() uint32 {
	rv := objc.Send[uint32](f.ID, objc.Sel("lossType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/numberOfClasses
func (f MPSNNForwardLoss) NumberOfClasses() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("numberOfClasses"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/reduceAcrossBatch
func (f MPSNNForwardLoss) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/reductionType
func (f MPSNNForwardLoss) ReductionType() int {
	rv := objc.Send[int](f.ID, objc.Sel("reductionType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLoss/weight
func (f MPSNNForwardLoss) Weight() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("weight"))
	return rv
}
func (f MPSNNForwardLoss) SetWeight(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setWeight:"), value)
}

















