// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNPadNode] class.
var (
	_MPSNNPadNodeClass     MPSNNPadNodeClass
	_MPSNNPadNodeClassOnce sync.Once
)

func getMPSNNPadNodeClass() MPSNNPadNodeClass {
	_MPSNNPadNodeClassOnce.Do(func() {
		_MPSNNPadNodeClass = MPSNNPadNodeClass{class: objc.GetClass("MPSNNPadNode")}
	})
	return _MPSNNPadNodeClass
}

// GetMPSNNPadNodeClass returns the class object for MPSNNPadNode.
func GetMPSNNPadNodeClass() MPSNNPadNodeClass {
	return getMPSNNPadNodeClass()
}

type MPSNNPadNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadNodeClass) Alloc() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNPadNode.FillValue]
//   - [MPSNNPadNode.SetFillValue]
//   - [MPSNNPadNode.InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode
type MPSNNPadNode struct {
	MPSNNFilterNode
}

// MPSNNPadNodeFromID constructs a [MPSNNPadNode] from an objc.ID.
func MPSNNPadNodeFromID(id objc.ID) MPSNNPadNode {
	return MPSNNPadNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNPadNode implements IMPSNNPadNode.
var _ IMPSNNPadNode = MPSNNPadNode{}





// An interface definition for the [MPSNNPadNode] class.
//
// # Methods
//
//   - [IMPSNNPadNode.FillValue]
//   - [IMPSNNPadNode.SetFillValue]
//   - [IMPSNNPadNode.InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode
type IMPSNNPadNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	FillValue() float32
	SetFillValue(value float32)
	InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source objectivec.IObject, before objectivec.IObject, after objectivec.IObject, mode uint64) MPSNNPadNode
}





// Init initializes the instance.
func (p MPSNNPadNode) Init() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadNode) Autorelease() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadNode creates a new MPSNNPadNode instance.
func NewMPSNNPadNode() MPSNNPadNode {
	class := getMPSNNPadNodeClass()
	rv := objc.Send[MPSNNPadNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewPadNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNPadNode {
	instance := getMPSNNPadNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNPadNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode/initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:
func NewPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source objectivec.IObject, before objectivec.IObject, after objectivec.IObject, mode uint64) MPSNNPadNode {
	instance := getMPSNNPadNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, before, after, mode)
	return MPSNNPadNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode/initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:
func (p MPSNNPadNode) InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source objectivec.IObject, before objectivec.IObject, after objectivec.IObject, mode uint64) MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, before, after, mode)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode/nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:
func (_MPSNNPadNodeClass MPSNNPadNodeClass) NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source objectivec.IObject, before objectivec.IObject, after objectivec.IObject, mode uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPadNodeClass.class), objc.Sel("nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, before, after, mode)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNPadNode/fillValue
func (p MPSNNPadNode) FillValue() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("fillValue"))
	return rv
}
func (p MPSNNPadNode) SetFillValue(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setFillValue:"), value)
}

















