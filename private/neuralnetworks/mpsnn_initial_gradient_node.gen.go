// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNInitialGradientNode] class.
var (
	_MPSNNInitialGradientNodeClass     MPSNNInitialGradientNodeClass
	_MPSNNInitialGradientNodeClassOnce sync.Once
)

func getMPSNNInitialGradientNodeClass() MPSNNInitialGradientNodeClass {
	_MPSNNInitialGradientNodeClassOnce.Do(func() {
		_MPSNNInitialGradientNodeClass = MPSNNInitialGradientNodeClass{class: objc.GetClass("MPSNNInitialGradientNode")}
	})
	return _MPSNNInitialGradientNodeClass
}

// GetMPSNNInitialGradientNodeClass returns the class object for MPSNNInitialGradientNode.
func GetMPSNNInitialGradientNodeClass() MPSNNInitialGradientNodeClass {
	return getMPSNNInitialGradientNodeClass()
}

type MPSNNInitialGradientNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNInitialGradientNodeClass) Alloc() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNInitialGradientNode.InitWithSource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradientNode
type MPSNNInitialGradientNode struct {
	MPSNNFilterNode
}

// MPSNNInitialGradientNodeFromID constructs a [MPSNNInitialGradientNode] from an objc.ID.
func MPSNNInitialGradientNodeFromID(id objc.ID) MPSNNInitialGradientNode {
	return MPSNNInitialGradientNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNInitialGradientNode implements IMPSNNInitialGradientNode.
var _ IMPSNNInitialGradientNode = MPSNNInitialGradientNode{}





// An interface definition for the [MPSNNInitialGradientNode] class.
//
// # Methods
//
//   - [IMPSNNInitialGradientNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradientNode
type IMPSNNInitialGradientNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	InitWithSource(source objectivec.IObject) MPSNNInitialGradientNode
}





// Init initializes the instance.
func (i MPSNNInitialGradientNode) Init() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSNNInitialGradientNode) Autorelease() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNInitialGradientNode creates a new MPSNNInitialGradientNode instance.
func NewMPSNNInitialGradientNode() MPSNNInitialGradientNode {
	class := getMPSNNInitialGradientNodeClass()
	rv := objc.Send[MPSNNInitialGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradientNode/initWithSource:
func NewInitialGradientNodeWithSource(source objectivec.IObject) MPSNNInitialGradientNode {
	instance := getMPSNNInitialGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNInitialGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewInitialGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNInitialGradientNode {
	instance := getMPSNNInitialGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNInitialGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradientNode/initWithSource:
func (i MPSNNInitialGradientNode) InitWithSource(source objectivec.IObject) MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("initWithSource:"), source)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradientNode/nodeWithSource:
func (_MPSNNInitialGradientNodeClass MPSNNInitialGradientNodeClass) NodeWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNInitialGradientNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return objectivec.Object{ID: rv}
}






















