// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPermuteNode] class.
var (
	_MPSNNPermuteNodeClass     MPSNNPermuteNodeClass
	_MPSNNPermuteNodeClassOnce sync.Once
)

func getMPSNNPermuteNodeClass() MPSNNPermuteNodeClass {
	_MPSNNPermuteNodeClassOnce.Do(func() {
		_MPSNNPermuteNodeClass = MPSNNPermuteNodeClass{class: objc.GetClass("MPSNNPermuteNode")}
	})
	return _MPSNNPermuteNodeClass
}

// GetMPSNNPermuteNodeClass returns the class object for MPSNNPermuteNode.
func GetMPSNNPermuteNodeClass() MPSNNPermuteNodeClass {
	return getMPSNNPermuteNodeClass()
}

type MPSNNPermuteNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPermuteNodeClass) Alloc() MPSNNPermuteNode {
	rv := objc.Send[MPSNNPermuteNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPermuteNode.InitWithSourceDimensionOrder]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteNode
type MPSNNPermuteNode struct {
	MPSNNFilterNode
}

// MPSNNPermuteNodeFromID constructs a [MPSNNPermuteNode] from an objc.ID.
func MPSNNPermuteNodeFromID(id objc.ID) MPSNNPermuteNode {
	return MPSNNPermuteNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNPermuteNode implements IMPSNNPermuteNode.
var _ IMPSNNPermuteNode = MPSNNPermuteNode{}





// An interface definition for the [MPSNNPermuteNode] class.
//
// # Methods
//
//   - [IMPSNNPermuteNode.InitWithSourceDimensionOrder]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteNode
type IMPSNNPermuteNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	InitWithSourceDimensionOrder(source objectivec.IObject, order objectivec.IObject) MPSNNPermuteNode
}





// Init initializes the instance.
func (p MPSNNPermuteNode) Init() MPSNNPermuteNode {
	rv := objc.Send[MPSNNPermuteNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPermuteNode) Autorelease() MPSNNPermuteNode {
	rv := objc.Send[MPSNNPermuteNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPermuteNode creates a new MPSNNPermuteNode instance.
func NewMPSNNPermuteNode() MPSNNPermuteNode {
	class := getMPSNNPermuteNodeClass()
	rv := objc.Send[MPSNNPermuteNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteNode/initWithSource:dimensionOrder:
func NewPermuteNodeWithSourceDimensionOrder(source objectivec.IObject, order objectivec.IObject) MPSNNPermuteNode {
	instance := getMPSNNPermuteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:dimensionOrder:"), source, order)
	return MPSNNPermuteNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewPermuteNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNPermuteNode {
	instance := getMPSNNPermuteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNPermuteNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteNode/initWithSource:dimensionOrder:
func (p MPSNNPermuteNode) InitWithSourceDimensionOrder(source objectivec.IObject, order objectivec.IObject) MPSNNPermuteNode {
	rv := objc.Send[MPSNNPermuteNode](p.ID, objc.Sel("initWithSource:dimensionOrder:"), source, order)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPermuteNode/nodeWithSource:dimensionOrder:
func (_MPSNNPermuteNodeClass MPSNNPermuteNodeClass) NodeWithSourceDimensionOrder(source objectivec.IObject, order objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPermuteNodeClass.class), objc.Sel("nodeWithSource:dimensionOrder:"), source, order)
	return objectivec.Object{ID: rv}
}






















