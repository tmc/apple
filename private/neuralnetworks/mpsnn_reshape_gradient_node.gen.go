// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReshapeGradientNode] class.
var (
	_MPSNNReshapeGradientNodeClass     MPSNNReshapeGradientNodeClass
	_MPSNNReshapeGradientNodeClassOnce sync.Once
)

func getMPSNNReshapeGradientNodeClass() MPSNNReshapeGradientNodeClass {
	_MPSNNReshapeGradientNodeClassOnce.Do(func() {
		_MPSNNReshapeGradientNodeClass = MPSNNReshapeGradientNodeClass{class: objc.GetClass("MPSNNReshapeGradientNode")}
	})
	return _MPSNNReshapeGradientNodeClass
}

// GetMPSNNReshapeGradientNodeClass returns the class object for MPSNNReshapeGradientNode.
func GetMPSNNReshapeGradientNodeClass() MPSNNReshapeGradientNodeClass {
	return getMPSNNReshapeGradientNodeClass()
}

type MPSNNReshapeGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeGradientNodeClass) Alloc() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReshapeGradientNode.InitWithSourceGradientSourceImageGradientState]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradientNode
type MPSNNReshapeGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNReshapeGradientNodeFromID constructs a [MPSNNReshapeGradientNode] from an objc.ID.
func MPSNNReshapeGradientNodeFromID(id objc.ID) MPSNNReshapeGradientNode {
	return MPSNNReshapeGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNReshapeGradientNode implements IMPSNNReshapeGradientNode.
var _ IMPSNNReshapeGradientNode = MPSNNReshapeGradientNode{}





// An interface definition for the [MPSNNReshapeGradientNode] class.
//
// # Methods
//
//   - [IMPSNNReshapeGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradientNode
type IMPSNNReshapeGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReshapeGradientNode
}





// Init initializes the instance.
func (r MPSNNReshapeGradientNode) Init() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeGradientNode) Autorelease() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeGradientNode creates a new MPSNNReshapeGradientNode instance.
func NewMPSNNReshapeGradientNode() MPSNNReshapeGradientNode {
	class := getMPSNNReshapeGradientNodeClass()
	rv := objc.Send[MPSNNReshapeGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewReshapeGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNReshapeGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewReshapeGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNReshapeGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewReshapeGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNReshapeGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewReshapeGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNReshapeGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReshapeGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReshapeGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (r MPSNNReshapeGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNReshapeGradientNodeClass MPSNNReshapeGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReshapeGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}






















