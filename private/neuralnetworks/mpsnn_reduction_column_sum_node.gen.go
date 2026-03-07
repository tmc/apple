// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionColumnSumNode] class.
var (
	_MPSNNReductionColumnSumNodeClass     MPSNNReductionColumnSumNodeClass
	_MPSNNReductionColumnSumNodeClassOnce sync.Once
)

func getMPSNNReductionColumnSumNodeClass() MPSNNReductionColumnSumNodeClass {
	_MPSNNReductionColumnSumNodeClassOnce.Do(func() {
		_MPSNNReductionColumnSumNodeClass = MPSNNReductionColumnSumNodeClass{class: objc.GetClass("MPSNNReductionColumnSumNode")}
	})
	return _MPSNNReductionColumnSumNodeClass
}

// GetMPSNNReductionColumnSumNodeClass returns the class object for MPSNNReductionColumnSumNode.
func GetMPSNNReductionColumnSumNodeClass() MPSNNReductionColumnSumNodeClass {
	return getMPSNNReductionColumnSumNodeClass()
}

type MPSNNReductionColumnSumNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnSumNodeClass) Alloc() MPSNNReductionColumnSumNode {
	rv := objc.Send[MPSNNReductionColumnSumNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnSumNode
type MPSNNReductionColumnSumNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnSumNodeFromID constructs a [MPSNNReductionColumnSumNode] from an objc.ID.
func MPSNNReductionColumnSumNodeFromID(id objc.ID) MPSNNReductionColumnSumNode {
	return MPSNNReductionColumnSumNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionColumnSumNode implements IMPSNNReductionColumnSumNode.
var _ IMPSNNReductionColumnSumNode = MPSNNReductionColumnSumNode{}





// An interface definition for the [MPSNNReductionColumnSumNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnSumNode
type IMPSNNReductionColumnSumNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionColumnSumNode) Init() MPSNNReductionColumnSumNode {
	rv := objc.Send[MPSNNReductionColumnSumNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnSumNode) Autorelease() MPSNNReductionColumnSumNode {
	rv := objc.Send[MPSNNReductionColumnSumNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnSumNode creates a new MPSNNReductionColumnSumNode instance.
func NewMPSNNReductionColumnSumNode() MPSNNReductionColumnSumNode {
	class := getMPSNNReductionColumnSumNodeClass()
	rv := objc.Send[MPSNNReductionColumnSumNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionColumnSumNodeWithSource(source objectivec.IObject) MPSNNReductionColumnSumNode {
	instance := getMPSNNReductionColumnSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionColumnSumNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionColumnSumNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionColumnSumNode {
	instance := getMPSNNReductionColumnSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionColumnSumNodeFromID(rv)
}
































