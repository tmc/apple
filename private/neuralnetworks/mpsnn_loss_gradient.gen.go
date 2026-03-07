// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLossGradient] class.
var (
	_MPSNNLossGradientClass     MPSNNLossGradientClass
	_MPSNNLossGradientClassOnce sync.Once
)

func getMPSNNLossGradientClass() MPSNNLossGradientClass {
	_MPSNNLossGradientClassOnce.Do(func() {
		_MPSNNLossGradientClass = MPSNNLossGradientClass{class: objc.GetClass("MPSNNLossGradient")}
	})
	return _MPSNNLossGradientClass
}

// GetMPSNNLossGradientClass returns the class object for MPSNNLossGradient.
func GetMPSNNLossGradientClass() MPSNNLossGradientClass {
	return getMPSNNLossGradientClass()
}

type MPSNNLossGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLossGradientClass) Alloc() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNLossGradient.ComputeLabelGradients]
//   - [MPSNNLossGradient.SetComputeLabelGradients]
//   - [MPSNNLossGradient.CopyWithZoneDevice]
//   - [MPSNNLossGradient.Delta]
//   - [MPSNNLossGradient.SetDelta]
//   - [MPSNNLossGradient.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [MPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates]
//   - [MPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients]
//   - [MPSNNLossGradient.Epsilon]
//   - [MPSNNLossGradient.SetEpsilon]
//   - [MPSNNLossGradient.LabelSmoothing]
//   - [MPSNNLossGradient.SetLabelSmoothing]
//   - [MPSNNLossGradient.LossType]
//   - [MPSNNLossGradient.MaxBatchSize]
//   - [MPSNNLossGradient.NumberOfClasses]
//   - [MPSNNLossGradient.ReduceAcrossBatch]
//   - [MPSNNLossGradient.ReductionType]
//   - [MPSNNLossGradient.Weight]
//   - [MPSNNLossGradient.SetWeight]
//   - [MPSNNLossGradient.InitWithCoderDevice]
//   - [MPSNNLossGradient.InitWithDevice]
//   - [MPSNNLossGradient.InitWithDeviceLossDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient
type MPSNNLossGradient struct {
	MPSCNNBinaryKernel
}

// MPSNNLossGradientFromID constructs a [MPSNNLossGradient] from an objc.ID.
func MPSNNLossGradientFromID(id objc.ID) MPSNNLossGradient {
	return MPSNNLossGradient{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}
// Ensure MPSNNLossGradient implements IMPSNNLossGradient.
var _ IMPSNNLossGradient = MPSNNLossGradient{}





// An interface definition for the [MPSNNLossGradient] class.
//
// # Methods
//
//   - [IMPSNNLossGradient.ComputeLabelGradients]
//   - [IMPSNNLossGradient.SetComputeLabelGradients]
//   - [IMPSNNLossGradient.CopyWithZoneDevice]
//   - [IMPSNNLossGradient.Delta]
//   - [IMPSNNLossGradient.SetDelta]
//   - [IMPSNNLossGradient.DestinationImageDescriptorForSourceImagesSourceStates]
//   - [IMPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates]
//   - [IMPSNNLossGradient.EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients]
//   - [IMPSNNLossGradient.Epsilon]
//   - [IMPSNNLossGradient.SetEpsilon]
//   - [IMPSNNLossGradient.LabelSmoothing]
//   - [IMPSNNLossGradient.SetLabelSmoothing]
//   - [IMPSNNLossGradient.LossType]
//   - [IMPSNNLossGradient.MaxBatchSize]
//   - [IMPSNNLossGradient.NumberOfClasses]
//   - [IMPSNNLossGradient.ReduceAcrossBatch]
//   - [IMPSNNLossGradient.ReductionType]
//   - [IMPSNNLossGradient.Weight]
//   - [IMPSNNLossGradient.SetWeight]
//   - [IMPSNNLossGradient.InitWithCoderDevice]
//   - [IMPSNNLossGradient.InitWithDevice]
//   - [IMPSNNLossGradient.InitWithDeviceLossDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient
type IMPSNNLossGradient interface {
	IMPSCNNBinaryKernel

	// Topic: Methods

	ComputeLabelGradients() bool
	SetComputeLabelGradients(value bool)
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	Delta() float32
	SetDelta(value float32)
	DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(buffer objectivec.IObject, gradients objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(buffer objectivec.IObject, gradients objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject, gradients2 objectivec.IObject)
	Epsilon() float32
	SetEpsilon(value float32)
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() uint32
	MaxBatchSize() uint64
	NumberOfClasses() uint64
	ReduceAcrossBatch() bool
	ReductionType() int
	Weight() float32
	SetWeight(value float32)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNLossGradient
	InitWithDevice(device objectivec.IObject) MPSNNLossGradient
	InitWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNLossGradient
}





// Init initializes the instance.
func (l MPSNNLossGradient) Init() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLossGradient) Autorelease() MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLossGradient creates a new MPSNNLossGradient instance.
func NewMPSNNLossGradient() MPSNNLossGradient {
	class := getMPSNNLossGradientClass()
	rv := objc.Send[MPSNNLossGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithCoder:device:
func NewLossGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNLossGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithDevice:
func NewLossGradientWithDevice(device objectivec.IObject) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNLossGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithDevice:lossDescriptor:
func NewLossGradientWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNLossGradient {
	instance := getMPSNNLossGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, descriptor)
	return MPSNNLossGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/copyWithZone:device:
func (l MPSNNLossGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/destinationImageDescriptorForSourceImages:sourceStates:
func (l MPSNNLossGradient) DestinationImageDescriptorForSourceImagesSourceStates(images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), images, states)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:
func (l MPSNNLossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(buffer objectivec.IObject, gradients objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:"), buffer, gradients, images, labels, weights, states)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:
func (l MPSNNLossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(buffer objectivec.IObject, gradients objectivec.IObject, images objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, states objectivec.IObject, gradients2 objectivec.IObject) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), buffer, gradients, images, labels, weights, states, gradients2)
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/maxBatchSize
func (l MPSNNLossGradient) MaxBatchSize() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("maxBatchSize"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithCoder:device:
func (l MPSNNLossGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithDevice:
func (l MPSNNLossGradient) InitWithDevice(device objectivec.IObject) MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/initWithDevice:lossDescriptor:
func (l MPSNNLossGradient) InitWithDeviceLossDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNLossGradient {
	rv := objc.Send[MPSNNLossGradient](l.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, descriptor)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/libraryInfo:
func (_MPSNNLossGradientClass MPSNNLossGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNLossGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/computeLabelGradients
func (l MPSNNLossGradient) ComputeLabelGradients() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("computeLabelGradients"))
	return rv
}
func (l MPSNNLossGradient) SetComputeLabelGradients(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setComputeLabelGradients:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/delta
func (l MPSNNLossGradient) Delta() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("delta"))
	return rv
}
func (l MPSNNLossGradient) SetDelta(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setDelta:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/epsilon
func (l MPSNNLossGradient) Epsilon() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("epsilon"))
	return rv
}
func (l MPSNNLossGradient) SetEpsilon(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setEpsilon:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/labelSmoothing
func (l MPSNNLossGradient) LabelSmoothing() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("labelSmoothing"))
	return rv
}
func (l MPSNNLossGradient) SetLabelSmoothing(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setLabelSmoothing:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/lossType
func (l MPSNNLossGradient) LossType() uint32 {
	rv := objc.Send[uint32](l.ID, objc.Sel("lossType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/numberOfClasses
func (l MPSNNLossGradient) NumberOfClasses() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("numberOfClasses"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/reduceAcrossBatch
func (l MPSNNLossGradient) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/reductionType
func (l MPSNNLossGradient) ReductionType() int {
	rv := objc.Send[int](l.ID, objc.Sel("reductionType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradient/weight
func (l MPSNNLossGradient) Weight() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("weight"))
	return rv
}
func (l MPSNNLossGradient) SetWeight(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setWeight:"), value)
}

















