// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLossGradientNode] class.
var (
	_MPSNNLossGradientNodeClass     MPSNNLossGradientNodeClass
	_MPSNNLossGradientNodeClassOnce sync.Once
)

func getMPSNNLossGradientNodeClass() MPSNNLossGradientNodeClass {
	_MPSNNLossGradientNodeClassOnce.Do(func() {
		_MPSNNLossGradientNodeClass = MPSNNLossGradientNodeClass{class: objc.GetClass("MPSNNLossGradientNode")}
	})
	return _MPSNNLossGradientNodeClass
}

// GetMPSNNLossGradientNodeClass returns the class object for MPSNNLossGradientNode.
func GetMPSNNLossGradientNodeClass() MPSNNLossGradientNodeClass {
	return getMPSNNLossGradientNodeClass()
}

type MPSNNLossGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLossGradientNodeClass) Alloc() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNLossGradientNode.Delta]
//   - [MPSNNLossGradientNode.Epsilon]
//   - [MPSNNLossGradientNode.IsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.LabelSmoothing]
//   - [MPSNNLossGradientNode.LossType]
//   - [MPSNNLossGradientNode.NumberOfClasses]
//   - [MPSNNLossGradientNode.PropertyCallBack]
//   - [MPSNNLossGradientNode.SetPropertyCallBack]
//   - [MPSNNLossGradientNode.ReduceAcrossBatch]
//   - [MPSNNLossGradientNode.ReductionType]
//   - [MPSNNLossGradientNode.Weight]
//   - [MPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode
type MPSNNLossGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNLossGradientNodeFromID constructs a [MPSNNLossGradientNode] from an objc.ID.
func MPSNNLossGradientNodeFromID(id objc.ID) MPSNNLossGradientNode {
	return MPSNNLossGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNLossGradientNode implements IMPSNNLossGradientNode.
var _ IMPSNNLossGradientNode = MPSNNLossGradientNode{}





// An interface definition for the [MPSNNLossGradientNode] class.
//
// # Methods
//
//   - [IMPSNNLossGradientNode.Delta]
//   - [IMPSNNLossGradientNode.Epsilon]
//   - [IMPSNNLossGradientNode.IsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.LabelSmoothing]
//   - [IMPSNNLossGradientNode.LossType]
//   - [IMPSNNLossGradientNode.NumberOfClasses]
//   - [IMPSNNLossGradientNode.PropertyCallBack]
//   - [IMPSNNLossGradientNode.SetPropertyCallBack]
//   - [IMPSNNLossGradientNode.ReduceAcrossBatch]
//   - [IMPSNNLossGradientNode.ReductionType]
//   - [IMPSNNLossGradientNode.Weight]
//   - [IMPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode
type IMPSNNLossGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	Delta() float32
	Epsilon() float32
	IsLabelsGradientFilter() bool
	LabelSmoothing() float32
	LossType() uint32
	NumberOfClasses() uint64
	PropertyCallBack() unsafe.Pointer
	SetPropertyCallBack(value unsafe.Pointer)
	ReduceAcrossBatch() bool
	ReductionType() int
	Weight() float32
	InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode
	InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode
	InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sources objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode
}





// Init initializes the instance.
func (l MPSNNLossGradientNode) Init() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLossGradientNode) Autorelease() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLossGradientNode creates a new MPSNNLossGradientNode instance.
func NewMPSNNLossGradientNode() MPSNNLossGradientNode {
	class := getMPSNNLossGradientNodeClass()
	rv := objc.Send[MPSNNLossGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewLossGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewLossGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewLossGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:
func NewLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, state, descriptor, filter)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:
func NewLossGradientNodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, weights, state, descriptor, filter)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewLossGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNLossGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func NewLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sources objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sources, state, descriptor, filter)
	return MPSNNLossGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:
func (l MPSNNLossGradientNode) InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, state, descriptor, filter)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:
func (l MPSNNLossGradientNode) InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, weights, state, descriptor, filter)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (l MPSNNLossGradientNode) InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sources objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sources, state, descriptor, filter)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, state, descriptor, filter)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/nodeWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(gradient objectivec.IObject, image objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), gradient, image, labels, weights, state, descriptor, filter)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sources objectivec.IObject, state objectivec.IObject, descriptor objectivec.IObject, filter bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sources, state, descriptor, filter)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/delta
func (l MPSNNLossGradientNode) Delta() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("delta"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/epsilon
func (l MPSNNLossGradientNode) Epsilon() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("epsilon"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/isLabelsGradientFilter
func (l MPSNNLossGradientNode) IsLabelsGradientFilter() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/labelSmoothing
func (l MPSNNLossGradientNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("labelSmoothing"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/lossType
func (l MPSNNLossGradientNode) LossType() uint32 {
	rv := objc.Send[uint32](l.ID, objc.Sel("lossType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/numberOfClasses
func (l MPSNNLossGradientNode) NumberOfClasses() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("numberOfClasses"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/propertyCallBack
func (l MPSNNLossGradientNode) PropertyCallBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l.ID, objc.Sel("propertyCallBack"))
	return rv
}
func (l MPSNNLossGradientNode) SetPropertyCallBack(value unsafe.Pointer) {
	objc.Send[struct{}](l.ID, objc.Sel("setPropertyCallBack:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/reduceAcrossBatch
func (l MPSNNLossGradientNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/reductionType
func (l MPSNNLossGradientNode) ReductionType() int {
	rv := objc.Send[int](l.ID, objc.Sel("reductionType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLossGradientNode/weight
func (l MPSNNLossGradientNode) Weight() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("weight"))
	return rv
}

















