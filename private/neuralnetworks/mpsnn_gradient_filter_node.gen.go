// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGradientFilterNode] class.
var (
	_MPSNNGradientFilterNodeClass     MPSNNGradientFilterNodeClass
	_MPSNNGradientFilterNodeClassOnce sync.Once
)

func getMPSNNGradientFilterNodeClass() MPSNNGradientFilterNodeClass {
	_MPSNNGradientFilterNodeClassOnce.Do(func() {
		_MPSNNGradientFilterNodeClass = MPSNNGradientFilterNodeClass{class: objc.GetClass("MPSNNGradientFilterNode")}
	})
	return _MPSNNGradientFilterNodeClass
}

// GetMPSNNGradientFilterNodeClass returns the class object for MPSNNGradientFilterNode.
func GetMPSNNGradientFilterNodeClass() MPSNNGradientFilterNodeClass {
	return getMPSNNGradientFilterNodeClass()
}

type MPSNNGradientFilterNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientFilterNodeClass) Alloc() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGradientFilterNode.InitWithGradientImagesForwardFilter]
//   - [MPSNNGradientFilterNode.InitWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy]
//   - [MPSNNGradientFilterNode.InitWithGradientImagesSourceImagesGradientStatePaddingPolicy]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode
type MPSNNGradientFilterNode struct {
	MPSNNFilterNode
}

// MPSNNGradientFilterNodeFromID constructs a [MPSNNGradientFilterNode] from an objc.ID.
func MPSNNGradientFilterNodeFromID(id objc.ID) MPSNNGradientFilterNode {
	return MPSNNGradientFilterNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNGradientFilterNode implements IMPSNNGradientFilterNode.
var _ IMPSNNGradientFilterNode = MPSNNGradientFilterNode{}





// An interface definition for the [MPSNNGradientFilterNode] class.
//
// # Methods
//
//   - [IMPSNNGradientFilterNode.InitWithGradientImagesForwardFilter]
//   - [IMPSNNGradientFilterNode.InitWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy]
//   - [IMPSNNGradientFilterNode.InitWithGradientImagesSourceImagesGradientStatePaddingPolicy]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode
type IMPSNNGradientFilterNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	InitWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNGradientFilterNode
	InitWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode
	InitWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode
}





// Init initializes the instance.
func (g MPSNNGradientFilterNode) Init() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientFilterNode) Autorelease() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientFilterNode creates a new MPSNNGradientFilterNode instance.
func NewMPSNNGradientFilterNode() MPSNNGradientFilterNode {
	class := getMPSNNGradientFilterNodeClass()
	rv := objc.Send[MPSNNGradientFilterNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewGradientFilterNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNGradientFilterNode {
	instance := getMPSNNGradientFilterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNGradientFilterNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewGradientFilterNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode {
	instance := getMPSNNGradientFilterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNGradientFilterNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewGradientFilterNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode {
	instance := getMPSNNGradientFilterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNGradientFilterNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewGradientFilterNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode {
	instance := getMPSNNGradientFilterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNGradientFilterNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func (g MPSNNGradientFilterNode) InitWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func (g MPSNNGradientFilterNode) InitWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func (g MPSNNGradientFilterNode) InitWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return rv
}


























