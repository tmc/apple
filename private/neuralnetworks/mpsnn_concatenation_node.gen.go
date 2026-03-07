// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenationNode] class.
var (
	_MPSNNConcatenationNodeClass     MPSNNConcatenationNodeClass
	_MPSNNConcatenationNodeClassOnce sync.Once
)

func getMPSNNConcatenationNodeClass() MPSNNConcatenationNodeClass {
	_MPSNNConcatenationNodeClassOnce.Do(func() {
		_MPSNNConcatenationNodeClass = MPSNNConcatenationNodeClass{class: objc.GetClass("MPSNNConcatenationNode")}
	})
	return _MPSNNConcatenationNodeClass
}

// GetMPSNNConcatenationNodeClass returns the class object for MPSNNConcatenationNode.
func GetMPSNNConcatenationNodeClass() MPSNNConcatenationNodeClass {
	return getMPSNNConcatenationNodeClass()
}

type MPSNNConcatenationNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationNodeClass) Alloc() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNConcatenationNode.InitWithSource]
//   - [MPSNNConcatenationNode.InitWithSources]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode
type MPSNNConcatenationNode struct {
	MPSNNFilterNode
}

// MPSNNConcatenationNodeFromID constructs a [MPSNNConcatenationNode] from an objc.ID.
func MPSNNConcatenationNodeFromID(id objc.ID) MPSNNConcatenationNode {
	return MPSNNConcatenationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNConcatenationNode implements IMPSNNConcatenationNode.
var _ IMPSNNConcatenationNode = MPSNNConcatenationNode{}





// An interface definition for the [MPSNNConcatenationNode] class.
//
// # Methods
//
//   - [IMPSNNConcatenationNode.InitWithSource]
//   - [IMPSNNConcatenationNode.InitWithSources]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode
type IMPSNNConcatenationNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	InitWithSource(source objectivec.IObject) MPSNNConcatenationNode
	InitWithSources(sources objectivec.IObject) MPSNNConcatenationNode
}





// Init initializes the instance.
func (c MPSNNConcatenationNode) Init() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationNode) Autorelease() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationNode creates a new MPSNNConcatenationNode instance.
func NewMPSNNConcatenationNode() MPSNNConcatenationNode {
	class := getMPSNNConcatenationNodeClass()
	rv := objc.Send[MPSNNConcatenationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/initWithSource:
func NewConcatenationNodeWithSource(source objectivec.IObject) MPSNNConcatenationNode {
	instance := getMPSNNConcatenationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNConcatenationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewConcatenationNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNConcatenationNode {
	instance := getMPSNNConcatenationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNConcatenationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/initWithSources:
func NewConcatenationNodeWithSources(sources objectivec.IObject) MPSNNConcatenationNode {
	instance := getMPSNNConcatenationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNConcatenationNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/initWithSource:
func (c MPSNNConcatenationNode) InitWithSource(source objectivec.IObject) MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("initWithSource:"), source)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/initWithSources:
func (c MPSNNConcatenationNode) InitWithSources(sources objectivec.IObject) MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("initWithSources:"), sources)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/nodeWithSource:
func (_MPSNNConcatenationNodeClass MPSNNConcatenationNodeClass) NodeWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNConcatenationNode/nodeWithSources:
func (_MPSNNConcatenationNodeClass MPSNNConcatenationNodeClass) NodeWithSources(sources objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationNodeClass.class), objc.Sel("nodeWithSources:"), sources)
	return objectivec.Object{ID: rv}
}






















