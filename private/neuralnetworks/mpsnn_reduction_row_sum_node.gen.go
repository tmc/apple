// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionRowSumNode] class.
var (
	_MPSNNReductionRowSumNodeClass     MPSNNReductionRowSumNodeClass
	_MPSNNReductionRowSumNodeClassOnce sync.Once
)

func getMPSNNReductionRowSumNodeClass() MPSNNReductionRowSumNodeClass {
	_MPSNNReductionRowSumNodeClassOnce.Do(func() {
		_MPSNNReductionRowSumNodeClass = MPSNNReductionRowSumNodeClass{class: objc.GetClass("MPSNNReductionRowSumNode")}
	})
	return _MPSNNReductionRowSumNodeClass
}

// GetMPSNNReductionRowSumNodeClass returns the class object for MPSNNReductionRowSumNode.
func GetMPSNNReductionRowSumNodeClass() MPSNNReductionRowSumNodeClass {
	return getMPSNNReductionRowSumNodeClass()
}

type MPSNNReductionRowSumNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowSumNodeClass) Alloc() MPSNNReductionRowSumNode {
	rv := objc.Send[MPSNNReductionRowSumNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowSumNode
type MPSNNReductionRowSumNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowSumNodeFromID constructs a [MPSNNReductionRowSumNode] from an objc.ID.
func MPSNNReductionRowSumNodeFromID(id objc.ID) MPSNNReductionRowSumNode {
	return MPSNNReductionRowSumNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionRowSumNode implements IMPSNNReductionRowSumNode.
var _ IMPSNNReductionRowSumNode = MPSNNReductionRowSumNode{}





// An interface definition for the [MPSNNReductionRowSumNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowSumNode
type IMPSNNReductionRowSumNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionRowSumNode) Init() MPSNNReductionRowSumNode {
	rv := objc.Send[MPSNNReductionRowSumNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowSumNode) Autorelease() MPSNNReductionRowSumNode {
	rv := objc.Send[MPSNNReductionRowSumNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowSumNode creates a new MPSNNReductionRowSumNode instance.
func NewMPSNNReductionRowSumNode() MPSNNReductionRowSumNode {
	class := getMPSNNReductionRowSumNodeClass()
	rv := objc.Send[MPSNNReductionRowSumNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionRowSumNodeWithSource(source objectivec.IObject) MPSNNReductionRowSumNode {
	instance := getMPSNNReductionRowSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionRowSumNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionRowSumNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionRowSumNode {
	instance := getMPSNNReductionRowSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionRowSumNodeFromID(rv)
}
































