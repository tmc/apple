// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNForwardLossNode] class.
var (
	_MPSNNForwardLossNodeClass     MPSNNForwardLossNodeClass
	_MPSNNForwardLossNodeClassOnce sync.Once
)

func getMPSNNForwardLossNodeClass() MPSNNForwardLossNodeClass {
	_MPSNNForwardLossNodeClassOnce.Do(func() {
		_MPSNNForwardLossNodeClass = MPSNNForwardLossNodeClass{class: objc.GetClass("MPSNNForwardLossNode")}
	})
	return _MPSNNForwardLossNodeClass
}

// GetMPSNNForwardLossNodeClass returns the class object for MPSNNForwardLossNode.
func GetMPSNNForwardLossNodeClass() MPSNNForwardLossNodeClass {
	return getMPSNNForwardLossNodeClass()
}

type MPSNNForwardLossNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNForwardLossNodeClass) Alloc() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNForwardLossNode.Delta]
//   - [MPSNNForwardLossNode.Epsilon]
//   - [MPSNNForwardLossNode.LabelSmoothing]
//   - [MPSNNForwardLossNode.LossType]
//   - [MPSNNForwardLossNode.NumberOfClasses]
//   - [MPSNNForwardLossNode.PropertyCallBack]
//   - [MPSNNForwardLossNode.SetPropertyCallBack]
//   - [MPSNNForwardLossNode.ReduceAcrossBatch]
//   - [MPSNNForwardLossNode.ReductionType]
//   - [MPSNNForwardLossNode.Weight]
//   - [MPSNNForwardLossNode.InitWithSourceLabelsLossDescriptor]
//   - [MPSNNForwardLossNode.InitWithSourceLabelsWeightsLossDescriptor]
//   - [MPSNNForwardLossNode.InitWithSourcesLossDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode
type MPSNNForwardLossNode struct {
	MPSNNFilterNode
}

// MPSNNForwardLossNodeFromID constructs a [MPSNNForwardLossNode] from an objc.ID.
func MPSNNForwardLossNodeFromID(id objc.ID) MPSNNForwardLossNode {
	return MPSNNForwardLossNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNForwardLossNode implements IMPSNNForwardLossNode.
var _ IMPSNNForwardLossNode = MPSNNForwardLossNode{}





// An interface definition for the [MPSNNForwardLossNode] class.
//
// # Methods
//
//   - [IMPSNNForwardLossNode.Delta]
//   - [IMPSNNForwardLossNode.Epsilon]
//   - [IMPSNNForwardLossNode.LabelSmoothing]
//   - [IMPSNNForwardLossNode.LossType]
//   - [IMPSNNForwardLossNode.NumberOfClasses]
//   - [IMPSNNForwardLossNode.PropertyCallBack]
//   - [IMPSNNForwardLossNode.SetPropertyCallBack]
//   - [IMPSNNForwardLossNode.ReduceAcrossBatch]
//   - [IMPSNNForwardLossNode.ReductionType]
//   - [IMPSNNForwardLossNode.Weight]
//   - [IMPSNNForwardLossNode.InitWithSourceLabelsLossDescriptor]
//   - [IMPSNNForwardLossNode.InitWithSourceLabelsWeightsLossDescriptor]
//   - [IMPSNNForwardLossNode.InitWithSourcesLossDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode
type IMPSNNForwardLossNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	Delta() float32
	Epsilon() float32
	LabelSmoothing() float32
	LossType() uint32
	NumberOfClasses() uint64
	PropertyCallBack() unsafe.Pointer
	SetPropertyCallBack(value unsafe.Pointer)
	ReduceAcrossBatch() bool
	ReductionType() int
	Weight() float32
	InitWithSourceLabelsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode
	InitWithSourceLabelsWeightsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode
	InitWithSourcesLossDescriptor(sources objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode
}





// Init initializes the instance.
func (f MPSNNForwardLossNode) Init() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSNNForwardLossNode) Autorelease() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNForwardLossNode creates a new MPSNNForwardLossNode instance.
func NewMPSNNForwardLossNode() MPSNNForwardLossNode {
	class := getMPSNNForwardLossNodeClass()
	rv := objc.Send[MPSNNForwardLossNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewForwardLossNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNForwardLossNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSource:labels:lossDescriptor:
func NewForwardLossNodeWithSourceLabelsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSource:labels:weights:lossDescriptor:
func NewForwardLossNodeWithSourceLabelsWeightsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSources:lossDescriptor:
func NewForwardLossNodeWithSourcesLossDescriptor(sources objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:lossDescriptor:"), sources, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSource:labels:lossDescriptor:
func (f MPSNNForwardLossNode) InitWithSourceLabelsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSource:labels:weights:lossDescriptor:
func (f MPSNNForwardLossNode) InitWithSourceLabelsWeightsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/initWithSources:lossDescriptor:
func (f MPSNNForwardLossNode) InitWithSourcesLossDescriptor(sources objectivec.IObject, descriptor objectivec.IObject) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSources:lossDescriptor:"), sources, descriptor)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/nodeWithSource:labels:lossDescriptor:
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourceLabelsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/nodeWithSource:labels:weights:lossDescriptor:
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourceLabelsWeightsLossDescriptor(source objectivec.IObject, labels objectivec.IObject, weights objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/nodeWithSources:lossDescriptor:
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourcesLossDescriptor(sources objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSources:lossDescriptor:"), sources, descriptor)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/delta
func (f MPSNNForwardLossNode) Delta() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("delta"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/epsilon
func (f MPSNNForwardLossNode) Epsilon() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("epsilon"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/labelSmoothing
func (f MPSNNForwardLossNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("labelSmoothing"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/lossType
func (f MPSNNForwardLossNode) LossType() uint32 {
	rv := objc.Send[uint32](f.ID, objc.Sel("lossType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/numberOfClasses
func (f MPSNNForwardLossNode) NumberOfClasses() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("numberOfClasses"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/propertyCallBack
func (f MPSNNForwardLossNode) PropertyCallBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f.ID, objc.Sel("propertyCallBack"))
	return rv
}
func (f MPSNNForwardLossNode) SetPropertyCallBack(value unsafe.Pointer) {
	objc.Send[struct{}](f.ID, objc.Sel("setPropertyCallBack:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/reduceAcrossBatch
func (f MPSNNForwardLossNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/reductionType
func (f MPSNNForwardLossNode) ReductionType() int {
	rv := objc.Send[int](f.ID, objc.Sel("reductionType"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNForwardLossNode/weight
func (f MPSNNForwardLossNode) Weight() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("weight"))
	return rv
}

















