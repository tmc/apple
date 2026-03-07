// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNMultiplicationGradientNode] class.
var (
	_MPSNNMultiplicationGradientNodeClass     MPSNNMultiplicationGradientNodeClass
	_MPSNNMultiplicationGradientNodeClassOnce sync.Once
)

func getMPSNNMultiplicationGradientNodeClass() MPSNNMultiplicationGradientNodeClass {
	_MPSNNMultiplicationGradientNodeClassOnce.Do(func() {
		_MPSNNMultiplicationGradientNodeClass = MPSNNMultiplicationGradientNodeClass{class: objc.GetClass("MPSNNMultiplicationGradientNode")}
	})
	return _MPSNNMultiplicationGradientNodeClass
}

// GetMPSNNMultiplicationGradientNodeClass returns the class object for MPSNNMultiplicationGradientNode.
func GetMPSNNMultiplicationGradientNodeClass() MPSNNMultiplicationGradientNodeClass {
	return getMPSNNMultiplicationGradientNodeClass()
}

type MPSNNMultiplicationGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiplicationGradientNodeClass) Alloc() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiplicationGradientNode
type MPSNNMultiplicationGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNMultiplicationGradientNodeFromID constructs a [MPSNNMultiplicationGradientNode] from an objc.ID.
func MPSNNMultiplicationGradientNodeFromID(id objc.ID) MPSNNMultiplicationGradientNode {
	return MPSNNMultiplicationGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}
// Ensure MPSNNMultiplicationGradientNode implements IMPSNNMultiplicationGradientNode.
var _ IMPSNNMultiplicationGradientNode = MPSNNMultiplicationGradientNode{}





// An interface definition for the [MPSNNMultiplicationGradientNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiplicationGradientNode
type IMPSNNMultiplicationGradientNode interface {
	IMPSNNArithmeticGradientNode
}





// Init initializes the instance.
func (m MPSNNMultiplicationGradientNode) Init() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiplicationGradientNode) Autorelease() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiplicationGradientNode creates a new MPSNNMultiplicationGradientNode instance.
func NewMPSNNMultiplicationGradientNode() MPSNNMultiplicationGradientNode {
	class := getMPSNNMultiplicationGradientNodeClass()
	rv := objc.Send[MPSNNMultiplicationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewMultiplicationGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithGradientImages:forwardFilter:isSecondarySourceFilter:
func NewMultiplicationGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), images, filter, filter2)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewMultiplicationGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewMultiplicationGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func NewMultiplicationGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewMultiplicationGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}
































