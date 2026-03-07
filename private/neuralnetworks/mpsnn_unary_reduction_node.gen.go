// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNUnaryReductionNode] class.
var (
	_MPSNNUnaryReductionNodeClass     MPSNNUnaryReductionNodeClass
	_MPSNNUnaryReductionNodeClassOnce sync.Once
)

func getMPSNNUnaryReductionNodeClass() MPSNNUnaryReductionNodeClass {
	_MPSNNUnaryReductionNodeClassOnce.Do(func() {
		_MPSNNUnaryReductionNodeClass = MPSNNUnaryReductionNodeClass{class: objc.GetClass("MPSNNUnaryReductionNode")}
	})
	return _MPSNNUnaryReductionNodeClass
}

// GetMPSNNUnaryReductionNodeClass returns the class object for MPSNNUnaryReductionNode.
func GetMPSNNUnaryReductionNodeClass() MPSNNUnaryReductionNodeClass {
	return getMPSNNUnaryReductionNodeClass()
}

type MPSNNUnaryReductionNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNUnaryReductionNodeClass) Alloc() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNUnaryReductionNode.ClipRectSource]
//   - [MPSNNUnaryReductionNode.SetClipRectSource]
//   - [MPSNNUnaryReductionNode.InitWithSource]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode
type MPSNNUnaryReductionNode struct {
	MPSNNFilterNode
}

// MPSNNUnaryReductionNodeFromID constructs a [MPSNNUnaryReductionNode] from an objc.ID.
func MPSNNUnaryReductionNodeFromID(id objc.ID) MPSNNUnaryReductionNode {
	return MPSNNUnaryReductionNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNUnaryReductionNode implements IMPSNNUnaryReductionNode.
var _ IMPSNNUnaryReductionNode = MPSNNUnaryReductionNode{}





// An interface definition for the [MPSNNUnaryReductionNode] class.
//
// # Methods
//
//   - [IMPSNNUnaryReductionNode.ClipRectSource]
//   - [IMPSNNUnaryReductionNode.SetClipRectSource]
//   - [IMPSNNUnaryReductionNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode
type IMPSNNUnaryReductionNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	ClipRectSource() objectivec.IObject
	SetClipRectSource(value objectivec.IObject)
	InitWithSource(source objectivec.IObject) MPSNNUnaryReductionNode
}





// Init initializes the instance.
func (u MPSNNUnaryReductionNode) Init() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MPSNNUnaryReductionNode) Autorelease() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNUnaryReductionNode creates a new MPSNNUnaryReductionNode instance.
func NewMPSNNUnaryReductionNode() MPSNNUnaryReductionNode {
	class := getMPSNNUnaryReductionNodeClass()
	rv := objc.Send[MPSNNUnaryReductionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewUnaryReductionNodeWithSource(source objectivec.IObject) MPSNNUnaryReductionNode {
	instance := getMPSNNUnaryReductionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNUnaryReductionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewUnaryReductionNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNUnaryReductionNode {
	instance := getMPSNNUnaryReductionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNUnaryReductionNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func (u MPSNNUnaryReductionNode) InitWithSource(source objectivec.IObject) MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("initWithSource:"), source)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/nodeWithSource:
func (_MPSNNUnaryReductionNodeClass MPSNNUnaryReductionNodeClass) NodeWithSource(source objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNUnaryReductionNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/clipRectSource
func (u MPSNNUnaryReductionNode) ClipRectSource() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("clipRectSource"))
	return objectivec.Object{ID: rv}
}
func (u MPSNNUnaryReductionNode) SetClipRectSource(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setClipRectSource:"), value)
}

















