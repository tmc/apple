// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionColumnMinNode] class.
var (
	_MPSNNReductionColumnMinNodeClass     MPSNNReductionColumnMinNodeClass
	_MPSNNReductionColumnMinNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMinNodeClass() MPSNNReductionColumnMinNodeClass {
	_MPSNNReductionColumnMinNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMinNodeClass = MPSNNReductionColumnMinNodeClass{class: objc.GetClass("MPSNNReductionColumnMinNode")}
	})
	return _MPSNNReductionColumnMinNodeClass
}

// GetMPSNNReductionColumnMinNodeClass returns the class object for MPSNNReductionColumnMinNode.
func GetMPSNNReductionColumnMinNodeClass() MPSNNReductionColumnMinNodeClass {
	return getMPSNNReductionColumnMinNodeClass()
}

type MPSNNReductionColumnMinNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMinNodeClass) Alloc() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMinNode
type MPSNNReductionColumnMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMinNodeFromID constructs a [MPSNNReductionColumnMinNode] from an objc.ID.
func MPSNNReductionColumnMinNodeFromID(id objc.ID) MPSNNReductionColumnMinNode {
	return MPSNNReductionColumnMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionColumnMinNode implements IMPSNNReductionColumnMinNode.
var _ IMPSNNReductionColumnMinNode = MPSNNReductionColumnMinNode{}





// An interface definition for the [MPSNNReductionColumnMinNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMinNode
type IMPSNNReductionColumnMinNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionColumnMinNode) Init() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMinNode) Autorelease() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMinNode creates a new MPSNNReductionColumnMinNode instance.
func NewMPSNNReductionColumnMinNode() MPSNNReductionColumnMinNode {
	class := getMPSNNReductionColumnMinNodeClass()
	rv := objc.Send[MPSNNReductionColumnMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionColumnMinNodeWithSource(source objectivec.IObject) MPSNNReductionColumnMinNode {
	instance := getMPSNNReductionColumnMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionColumnMinNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionColumnMinNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionColumnMinNode {
	instance := getMPSNNReductionColumnMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionColumnMinNodeFromID(rv)
}
































