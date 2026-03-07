// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNSubtractionGradientNode] class.
var (
	_MPSNNSubtractionGradientNodeClass     MPSNNSubtractionGradientNodeClass
	_MPSNNSubtractionGradientNodeClassOnce sync.Once
)

func getMPSNNSubtractionGradientNodeClass() MPSNNSubtractionGradientNodeClass {
	_MPSNNSubtractionGradientNodeClassOnce.Do(func() {
		_MPSNNSubtractionGradientNodeClass = MPSNNSubtractionGradientNodeClass{class: objc.GetClass("MPSNNSubtractionGradientNode")}
	})
	return _MPSNNSubtractionGradientNodeClass
}

// GetMPSNNSubtractionGradientNodeClass returns the class object for MPSNNSubtractionGradientNode.
func GetMPSNNSubtractionGradientNodeClass() MPSNNSubtractionGradientNodeClass {
	return getMPSNNSubtractionGradientNodeClass()
}

type MPSNNSubtractionGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSubtractionGradientNodeClass) Alloc() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSubtractionGradientNode
type MPSNNSubtractionGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNSubtractionGradientNodeFromID constructs a [MPSNNSubtractionGradientNode] from an objc.ID.
func MPSNNSubtractionGradientNodeFromID(id objc.ID) MPSNNSubtractionGradientNode {
	return MPSNNSubtractionGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}
// Ensure MPSNNSubtractionGradientNode implements IMPSNNSubtractionGradientNode.
var _ IMPSNNSubtractionGradientNode = MPSNNSubtractionGradientNode{}





// An interface definition for the [MPSNNSubtractionGradientNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSubtractionGradientNode
type IMPSNNSubtractionGradientNode interface {
	IMPSNNArithmeticGradientNode
}





// Init initializes the instance.
func (s MPSNNSubtractionGradientNode) Init() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNSubtractionGradientNode) Autorelease() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNSubtractionGradientNode creates a new MPSNNSubtractionGradientNode instance.
func NewMPSNNSubtractionGradientNode() MPSNNSubtractionGradientNode {
	class := getMPSNNSubtractionGradientNodeClass()
	rv := objc.Send[MPSNNSubtractionGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewSubtractionGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNSubtractionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithGradientImages:forwardFilter:isSecondarySourceFilter:
func NewSubtractionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), images, filter, filter2)
	return MPSNNSubtractionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewSubtractionGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNSubtractionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewSubtractionGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNSubtractionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func NewSubtractionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return MPSNNSubtractionGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewSubtractionGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNSubtractionGradientNodeFromID(rv)
}
































