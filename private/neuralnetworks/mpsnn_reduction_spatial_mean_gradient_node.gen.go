// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionSpatialMeanGradientNode] class.
var (
	_MPSNNReductionSpatialMeanGradientNodeClass     MPSNNReductionSpatialMeanGradientNodeClass
	_MPSNNReductionSpatialMeanGradientNodeClassOnce sync.Once
)

func getMPSNNReductionSpatialMeanGradientNodeClass() MPSNNReductionSpatialMeanGradientNodeClass {
	_MPSNNReductionSpatialMeanGradientNodeClassOnce.Do(func() {
		_MPSNNReductionSpatialMeanGradientNodeClass = MPSNNReductionSpatialMeanGradientNodeClass{class: objc.GetClass("MPSNNReductionSpatialMeanGradientNode")}
	})
	return _MPSNNReductionSpatialMeanGradientNodeClass
}

// GetMPSNNReductionSpatialMeanGradientNodeClass returns the class object for MPSNNReductionSpatialMeanGradientNode.
func GetMPSNNReductionSpatialMeanGradientNodeClass() MPSNNReductionSpatialMeanGradientNodeClass {
	return getMPSNNReductionSpatialMeanGradientNodeClass()
}

type MPSNNReductionSpatialMeanGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionSpatialMeanGradientNodeClass) Alloc() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReductionSpatialMeanGradientNode.InitWithSourceGradientSourceImageGradientState]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanGradientNode
type MPSNNReductionSpatialMeanGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNReductionSpatialMeanGradientNodeFromID constructs a [MPSNNReductionSpatialMeanGradientNode] from an objc.ID.
func MPSNNReductionSpatialMeanGradientNodeFromID(id objc.ID) MPSNNReductionSpatialMeanGradientNode {
	return MPSNNReductionSpatialMeanGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNReductionSpatialMeanGradientNode implements IMPSNNReductionSpatialMeanGradientNode.
var _ IMPSNNReductionSpatialMeanGradientNode = MPSNNReductionSpatialMeanGradientNode{}





// An interface definition for the [MPSNNReductionSpatialMeanGradientNode] class.
//
// # Methods
//
//   - [IMPSNNReductionSpatialMeanGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanGradientNode
type IMPSNNReductionSpatialMeanGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReductionSpatialMeanGradientNode
}





// Init initializes the instance.
func (r MPSNNReductionSpatialMeanGradientNode) Init() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionSpatialMeanGradientNode) Autorelease() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionSpatialMeanGradientNode creates a new MPSNNReductionSpatialMeanGradientNode instance.
func NewMPSNNReductionSpatialMeanGradientNode() MPSNNReductionSpatialMeanGradientNode {
	class := getMPSNNReductionSpatialMeanGradientNodeClass()
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewReductionSpatialMeanGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewReductionSpatialMeanGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewReductionSpatialMeanGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewReductionSpatialMeanGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionSpatialMeanGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (r MPSNNReductionSpatialMeanGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNReductionSpatialMeanGradientNodeClass MPSNNReductionSpatialMeanGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReductionSpatialMeanGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}






















