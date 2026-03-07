// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPermuteGradientNode] class.
var (
	_MPSNNPermuteGradientNodeClass     MPSNNPermuteGradientNodeClass
	_MPSNNPermuteGradientNodeClassOnce sync.Once
)

func getMPSNNPermuteGradientNodeClass() MPSNNPermuteGradientNodeClass {
	_MPSNNPermuteGradientNodeClassOnce.Do(func() {
		_MPSNNPermuteGradientNodeClass = MPSNNPermuteGradientNodeClass{class: objc.GetClass("MPSNNPermuteGradientNode")}
	})
	return _MPSNNPermuteGradientNodeClass
}

// GetMPSNNPermuteGradientNodeClass returns the class object for MPSNNPermuteGradientNode.
func GetMPSNNPermuteGradientNodeClass() MPSNNPermuteGradientNodeClass {
	return getMPSNNPermuteGradientNodeClass()
}

type MPSNNPermuteGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPermuteGradientNodeClass) Alloc() MPSNNPermuteGradientNode {
	rv := objc.Send[MPSNNPermuteGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPermuteGradientNode.InitWithSourceGradientSourceImageGradientState]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientNode
type MPSNNPermuteGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNPermuteGradientNodeFromID constructs a [MPSNNPermuteGradientNode] from an objc.ID.
func MPSNNPermuteGradientNodeFromID(id objc.ID) MPSNNPermuteGradientNode {
	return MPSNNPermuteGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNPermuteGradientNode implements IMPSNNPermuteGradientNode.
var _ IMPSNNPermuteGradientNode = MPSNNPermuteGradientNode{}





// An interface definition for the [MPSNNPermuteGradientNode] class.
//
// # Methods
//
//   - [IMPSNNPermuteGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientNode
type IMPSNNPermuteGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPermuteGradientNode
}





// Init initializes the instance.
func (p MPSNNPermuteGradientNode) Init() MPSNNPermuteGradientNode {
	rv := objc.Send[MPSNNPermuteGradientNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPermuteGradientNode) Autorelease() MPSNNPermuteGradientNode {
	rv := objc.Send[MPSNNPermuteGradientNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPermuteGradientNode creates a new MPSNNPermuteGradientNode instance.
func NewMPSNNPermuteGradientNode() MPSNNPermuteGradientNode {
	class := getMPSNNPermuteGradientNodeClass()
	rv := objc.Send[MPSNNPermuteGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewPermuteGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNPermuteGradientNode {
	instance := getMPSNNPermuteGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNPermuteGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewPermuteGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNPermuteGradientNode {
	instance := getMPSNNPermuteGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNPermuteGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewPermuteGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNPermuteGradientNode {
	instance := getMPSNNPermuteGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNPermuteGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewPermuteGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPermuteGradientNode {
	instance := getMPSNNPermuteGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNPermuteGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewPermuteGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNPermuteGradientNode {
	instance := getMPSNNPermuteGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNPermuteGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (p MPSNNPermuteGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPermuteGradientNode {
	rv := objc.Send[MPSNNPermuteGradientNode](p.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNPermuteGradientNodeClass MPSNNPermuteGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPermuteGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}






















