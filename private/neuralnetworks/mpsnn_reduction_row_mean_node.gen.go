// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionRowMeanNode] class.
var (
	_MPSNNReductionRowMeanNodeClass     MPSNNReductionRowMeanNodeClass
	_MPSNNReductionRowMeanNodeClassOnce sync.Once
)

func getMPSNNReductionRowMeanNodeClass() MPSNNReductionRowMeanNodeClass {
	_MPSNNReductionRowMeanNodeClassOnce.Do(func() {
		_MPSNNReductionRowMeanNodeClass = MPSNNReductionRowMeanNodeClass{class: objc.GetClass("MPSNNReductionRowMeanNode")}
	})
	return _MPSNNReductionRowMeanNodeClass
}

// GetMPSNNReductionRowMeanNodeClass returns the class object for MPSNNReductionRowMeanNode.
func GetMPSNNReductionRowMeanNodeClass() MPSNNReductionRowMeanNodeClass {
	return getMPSNNReductionRowMeanNodeClass()
}

type MPSNNReductionRowMeanNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowMeanNodeClass) Alloc() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowMeanNode
type MPSNNReductionRowMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowMeanNodeFromID constructs a [MPSNNReductionRowMeanNode] from an objc.ID.
func MPSNNReductionRowMeanNodeFromID(id objc.ID) MPSNNReductionRowMeanNode {
	return MPSNNReductionRowMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionRowMeanNode implements IMPSNNReductionRowMeanNode.
var _ IMPSNNReductionRowMeanNode = MPSNNReductionRowMeanNode{}





// An interface definition for the [MPSNNReductionRowMeanNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowMeanNode
type IMPSNNReductionRowMeanNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionRowMeanNode) Init() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowMeanNode) Autorelease() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowMeanNode creates a new MPSNNReductionRowMeanNode instance.
func NewMPSNNReductionRowMeanNode() MPSNNReductionRowMeanNode {
	class := getMPSNNReductionRowMeanNodeClass()
	rv := objc.Send[MPSNNReductionRowMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionRowMeanNodeWithSource(source objectivec.IObject) MPSNNReductionRowMeanNode {
	instance := getMPSNNReductionRowMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionRowMeanNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionRowMeanNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionRowMeanNode {
	instance := getMPSNNReductionRowMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionRowMeanNodeFromID(rv)
}
































