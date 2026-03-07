// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPadGradientNode] class.
var (
	_MPSNNPadGradientNodeClass     MPSNNPadGradientNodeClass
	_MPSNNPadGradientNodeClassOnce sync.Once
)

func getMPSNNPadGradientNodeClass() MPSNNPadGradientNodeClass {
	_MPSNNPadGradientNodeClassOnce.Do(func() {
		_MPSNNPadGradientNodeClass = MPSNNPadGradientNodeClass{class: objc.GetClass("MPSNNPadGradientNode")}
	})
	return _MPSNNPadGradientNodeClass
}

// GetMPSNNPadGradientNodeClass returns the class object for MPSNNPadGradientNode.
func GetMPSNNPadGradientNodeClass() MPSNNPadGradientNodeClass {
	return getMPSNNPadGradientNodeClass()
}

type MPSNNPadGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadGradientNodeClass) Alloc() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPadGradientNode.InitWithSourceGradientSourceImageGradientState]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientNode
type MPSNNPadGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNPadGradientNodeFromID constructs a [MPSNNPadGradientNode] from an objc.ID.
func MPSNNPadGradientNodeFromID(id objc.ID) MPSNNPadGradientNode {
	return MPSNNPadGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNPadGradientNode implements IMPSNNPadGradientNode.
var _ IMPSNNPadGradientNode = MPSNNPadGradientNode{}





// An interface definition for the [MPSNNPadGradientNode] class.
//
// # Methods
//
//   - [IMPSNNPadGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientNode
type IMPSNNPadGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPadGradientNode
}





// Init initializes the instance.
func (p MPSNNPadGradientNode) Init() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadGradientNode) Autorelease() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadGradientNode creates a new MPSNNPadGradientNode instance.
func NewMPSNNPadGradientNode() MPSNNPadGradientNode {
	class := getMPSNNPadGradientNodeClass()
	rv := objc.Send[MPSNNPadGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewPadGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNPadGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewPadGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNPadGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewPadGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNPadGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewPadGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNPadGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewPadGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNPadGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (p MPSNNPadGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNPadGradientNodeClass MPSNNPadGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPadGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}






















