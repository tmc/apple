// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionColumnMeanNode] class.
var (
	_MPSNNReductionColumnMeanNodeClass     MPSNNReductionColumnMeanNodeClass
	_MPSNNReductionColumnMeanNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMeanNodeClass() MPSNNReductionColumnMeanNodeClass {
	_MPSNNReductionColumnMeanNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMeanNodeClass = MPSNNReductionColumnMeanNodeClass{class: objc.GetClass("MPSNNReductionColumnMeanNode")}
	})
	return _MPSNNReductionColumnMeanNodeClass
}

// GetMPSNNReductionColumnMeanNodeClass returns the class object for MPSNNReductionColumnMeanNode.
func GetMPSNNReductionColumnMeanNodeClass() MPSNNReductionColumnMeanNodeClass {
	return getMPSNNReductionColumnMeanNodeClass()
}

type MPSNNReductionColumnMeanNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMeanNodeClass) Alloc() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMeanNode
type MPSNNReductionColumnMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMeanNodeFromID constructs a [MPSNNReductionColumnMeanNode] from an objc.ID.
func MPSNNReductionColumnMeanNodeFromID(id objc.ID) MPSNNReductionColumnMeanNode {
	return MPSNNReductionColumnMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionColumnMeanNode implements IMPSNNReductionColumnMeanNode.
var _ IMPSNNReductionColumnMeanNode = MPSNNReductionColumnMeanNode{}





// An interface definition for the [MPSNNReductionColumnMeanNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMeanNode
type IMPSNNReductionColumnMeanNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionColumnMeanNode) Init() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMeanNode) Autorelease() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMeanNode creates a new MPSNNReductionColumnMeanNode instance.
func NewMPSNNReductionColumnMeanNode() MPSNNReductionColumnMeanNode {
	class := getMPSNNReductionColumnMeanNodeClass()
	rv := objc.Send[MPSNNReductionColumnMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionColumnMeanNodeWithSource(source objectivec.IObject) MPSNNReductionColumnMeanNode {
	instance := getMPSNNReductionColumnMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionColumnMeanNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionColumnMeanNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionColumnMeanNode {
	instance := getMPSNNReductionColumnMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionColumnMeanNodeFromID(rv)
}
































