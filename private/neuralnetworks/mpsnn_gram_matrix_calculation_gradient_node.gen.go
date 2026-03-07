// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGramMatrixCalculationGradientNode] class.
var (
	_MPSNNGramMatrixCalculationGradientNodeClass     MPSNNGramMatrixCalculationGradientNodeClass
	_MPSNNGramMatrixCalculationGradientNodeClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationGradientNodeClass() MPSNNGramMatrixCalculationGradientNodeClass {
	_MPSNNGramMatrixCalculationGradientNodeClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationGradientNodeClass = MPSNNGramMatrixCalculationGradientNodeClass{class: objc.GetClass("MPSNNGramMatrixCalculationGradientNode")}
	})
	return _MPSNNGramMatrixCalculationGradientNodeClass
}

// GetMPSNNGramMatrixCalculationGradientNodeClass returns the class object for MPSNNGramMatrixCalculationGradientNode.
func GetMPSNNGramMatrixCalculationGradientNodeClass() MPSNNGramMatrixCalculationGradientNodeClass {
	return getMPSNNGramMatrixCalculationGradientNodeClass()
}

type MPSNNGramMatrixCalculationGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationGradientNodeClass) Alloc() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGramMatrixCalculationGradientNode.Alpha]
//   - [MPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientState]
//   - [MPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientStateAlpha]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode
type MPSNNGramMatrixCalculationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNGramMatrixCalculationGradientNodeFromID constructs a [MPSNNGramMatrixCalculationGradientNode] from an objc.ID.
func MPSNNGramMatrixCalculationGradientNodeFromID(id objc.ID) MPSNNGramMatrixCalculationGradientNode {
	return MPSNNGramMatrixCalculationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNGramMatrixCalculationGradientNode implements IMPSNNGramMatrixCalculationGradientNode.
var _ IMPSNNGramMatrixCalculationGradientNode = MPSNNGramMatrixCalculationGradientNode{}





// An interface definition for the [MPSNNGramMatrixCalculationGradientNode] class.
//
// # Methods
//
//   - [IMPSNNGramMatrixCalculationGradientNode.Alpha]
//   - [IMPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientState]
//   - [IMPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientStateAlpha]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode
type IMPSNNGramMatrixCalculationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

	Alpha() float32
	InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNGramMatrixCalculationGradientNode
	InitWithSourceGradientSourceImageGradientStateAlpha(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradientNode
}





// Init initializes the instance.
func (g MPSNNGramMatrixCalculationGradientNode) Init() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationGradientNode) Autorelease() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationGradientNode creates a new MPSNNGramMatrixCalculationGradientNode instance.
func NewMPSNNGramMatrixCalculationGradientNode() MPSNNGramMatrixCalculationGradientNode {
	class := getMPSNNGramMatrixCalculationGradientNodeClass()
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/initWithGradientImages:forwardFilter:
func NewGramMatrixCalculationGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewGramMatrixCalculationGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewGramMatrixCalculationGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/initWithSourceGradient:sourceImage:gradientState:
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/initWithSourceGradient:sourceImage:gradientState:alpha:
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientStateAlpha(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:alpha:"), gradient, image, state, alpha)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewGramMatrixCalculationGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/initWithSourceGradient:sourceImage:gradientState:
func (g MPSNNGramMatrixCalculationGradientNode) InitWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/initWithSourceGradient:sourceImage:gradientState:alpha:
func (g MPSNNGramMatrixCalculationGradientNode) InitWithSourceGradientSourceImageGradientStateAlpha(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, alpha float32) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:alpha:"), gradient, image, state, alpha)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNGramMatrixCalculationGradientNodeClass MPSNNGramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradient, image, state)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:alpha:
func (_MPSNNGramMatrixCalculationGradientNodeClass MPSNNGramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateAlpha(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, alpha float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:alpha:"), gradient, image, state, alpha)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGramMatrixCalculationGradientNode/alpha
func (g MPSNNGramMatrixCalculationGradientNode) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}

















