// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenationGradientNode] class.
var (
	_MPSNNConcatenationGradientNodeClass     MPSNNConcatenationGradientNodeClass
	_MPSNNConcatenationGradientNodeClassOnce sync.Once
)

func getMPSNNConcatenationGradientNodeClass() MPSNNConcatenationGradientNodeClass {
	_MPSNNConcatenationGradientNodeClassOnce.Do(func() {
		_MPSNNConcatenationGradientNodeClass = MPSNNConcatenationGradientNodeClass{class: objc.GetClass("MPSNNConcatenationGradientNode")}
	})
	return _MPSNNConcatenationGradientNodeClass
}

// GetMPSNNConcatenationGradientNodeClass returns the class object for MPSNNConcatenationGradientNode.
func GetMPSNNConcatenationGradientNodeClass() MPSNNConcatenationGradientNodeClass {
	return getMPSNNConcatenationGradientNodeClass()
}

type MPSNNConcatenationGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationGradientNodeClass) Alloc() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNConcatenationGradientNode.InitWithSourceGradientSourceImageGradientState]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientNode
type MPSNNConcatenationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNConcatenationGradientNodeFromID constructs a [MPSNNConcatenationGradientNode] from an objc.ID.
func MPSNNConcatenationGradientNodeFromID(id objc.ID) MPSNNConcatenationGradientNode {
	return MPSNNConcatenationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNConcatenationGradientNode implements IMPSNNConcatenationGradientNode.
var _ IMPSNNConcatenationGradientNode = MPSNNConcatenationGradientNode{}





// An interface definition for the [MPSNNConcatenationGradientNode] class.
//
// # Methods
//
//   - [IMPSNNConcatenationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientNode
type IMPSNNConcatenationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNConcatenationGradientNode
}





// Init initializes the instance.
func (c MPSNNConcatenationGradientNode) Init() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationGradientNode) Autorelease() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationGradientNode creates a new MPSNNConcatenationGradientNode instance.
func NewMPSNNConcatenationGradientNode() MPSNNConcatenationGradientNode {
	class := getMPSNNConcatenationGradientNodeClass()
	rv := objc.Send[MPSNNConcatenationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewConcatenationGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNConcatenationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewConcatenationGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNConcatenationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewConcatenationGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNConcatenationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewConcatenationGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNConcatenationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewConcatenationGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNConcatenationGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (c MPSNNConcatenationGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNConcatenationGradientNodeClass MPSNNConcatenationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}






















