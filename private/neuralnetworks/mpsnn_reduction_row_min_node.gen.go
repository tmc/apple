// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionRowMinNode] class.
var (
	_MPSNNReductionRowMinNodeClass     MPSNNReductionRowMinNodeClass
	_MPSNNReductionRowMinNodeClassOnce sync.Once
)

func getMPSNNReductionRowMinNodeClass() MPSNNReductionRowMinNodeClass {
	_MPSNNReductionRowMinNodeClassOnce.Do(func() {
		_MPSNNReductionRowMinNodeClass = MPSNNReductionRowMinNodeClass{class: objc.GetClass("MPSNNReductionRowMinNode")}
	})
	return _MPSNNReductionRowMinNodeClass
}

// GetMPSNNReductionRowMinNodeClass returns the class object for MPSNNReductionRowMinNode.
func GetMPSNNReductionRowMinNodeClass() MPSNNReductionRowMinNodeClass {
	return getMPSNNReductionRowMinNodeClass()
}

type MPSNNReductionRowMinNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowMinNodeClass) Alloc() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowMinNode
type MPSNNReductionRowMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowMinNodeFromID constructs a [MPSNNReductionRowMinNode] from an objc.ID.
func MPSNNReductionRowMinNodeFromID(id objc.ID) MPSNNReductionRowMinNode {
	return MPSNNReductionRowMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionRowMinNode implements IMPSNNReductionRowMinNode.
var _ IMPSNNReductionRowMinNode = MPSNNReductionRowMinNode{}





// An interface definition for the [MPSNNReductionRowMinNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionRowMinNode
type IMPSNNReductionRowMinNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionRowMinNode) Init() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowMinNode) Autorelease() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowMinNode creates a new MPSNNReductionRowMinNode instance.
func NewMPSNNReductionRowMinNode() MPSNNReductionRowMinNode {
	class := getMPSNNReductionRowMinNodeClass()
	rv := objc.Send[MPSNNReductionRowMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionRowMinNodeWithSource(source objectivec.IObject) MPSNNReductionRowMinNode {
	instance := getMPSNNReductionRowMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionRowMinNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionRowMinNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionRowMinNode {
	instance := getMPSNNReductionRowMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionRowMinNodeFromID(rv)
}
































