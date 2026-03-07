// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNAdditionGradientNode] class.
var (
	_MPSNNAdditionGradientNodeClass     MPSNNAdditionGradientNodeClass
	_MPSNNAdditionGradientNodeClassOnce sync.Once
)

func getMPSNNAdditionGradientNodeClass() MPSNNAdditionGradientNodeClass {
	_MPSNNAdditionGradientNodeClassOnce.Do(func() {
		_MPSNNAdditionGradientNodeClass = MPSNNAdditionGradientNodeClass{class: objc.GetClass("MPSNNAdditionGradientNode")}
	})
	return _MPSNNAdditionGradientNodeClass
}

// GetMPSNNAdditionGradientNodeClass returns the class object for MPSNNAdditionGradientNode.
func GetMPSNNAdditionGradientNodeClass() MPSNNAdditionGradientNodeClass {
	return getMPSNNAdditionGradientNodeClass()
}

type MPSNNAdditionGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNAdditionGradientNodeClass) Alloc() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNAdditionGradientNode
type MPSNNAdditionGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNAdditionGradientNodeFromID constructs a [MPSNNAdditionGradientNode] from an objc.ID.
func MPSNNAdditionGradientNodeFromID(id objc.ID) MPSNNAdditionGradientNode {
	return MPSNNAdditionGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}
// Ensure MPSNNAdditionGradientNode implements IMPSNNAdditionGradientNode.
var _ IMPSNNAdditionGradientNode = MPSNNAdditionGradientNode{}





// An interface definition for the [MPSNNAdditionGradientNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNAdditionGradientNode
type IMPSNNAdditionGradientNode interface {
	IMPSNNArithmeticGradientNode
}





// Init initializes the instance.
func (a MPSNNAdditionGradientNode) Init() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNAdditionGradientNode) Autorelease() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNAdditionGradientNode creates a new MPSNNAdditionGradientNode instance.
func NewMPSNNAdditionGradientNode() MPSNNAdditionGradientNode {
	class := getMPSNNAdditionGradientNodeClass()
	rv := objc.Send[MPSNNAdditionGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewAdditionGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNAdditionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithGradientImages:forwardFilter:isSecondarySourceFilter:
func NewAdditionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), images, filter, filter2)
	return MPSNNAdditionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewAdditionGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNAdditionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewAdditionGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNAdditionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func NewAdditionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return MPSNNAdditionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewAdditionGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNAdditionGradientNodeFromID(rv)
}
































