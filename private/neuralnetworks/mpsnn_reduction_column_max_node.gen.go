// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionColumnMaxNode] class.
var (
	_MPSNNReductionColumnMaxNodeClass     MPSNNReductionColumnMaxNodeClass
	_MPSNNReductionColumnMaxNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMaxNodeClass() MPSNNReductionColumnMaxNodeClass {
	_MPSNNReductionColumnMaxNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMaxNodeClass = MPSNNReductionColumnMaxNodeClass{class: objc.GetClass("MPSNNReductionColumnMaxNode")}
	})
	return _MPSNNReductionColumnMaxNodeClass
}

// GetMPSNNReductionColumnMaxNodeClass returns the class object for MPSNNReductionColumnMaxNode.
func GetMPSNNReductionColumnMaxNodeClass() MPSNNReductionColumnMaxNodeClass {
	return getMPSNNReductionColumnMaxNodeClass()
}

type MPSNNReductionColumnMaxNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMaxNodeClass) Alloc() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMaxNode
type MPSNNReductionColumnMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMaxNodeFromID constructs a [MPSNNReductionColumnMaxNode] from an objc.ID.
func MPSNNReductionColumnMaxNodeFromID(id objc.ID) MPSNNReductionColumnMaxNode {
	return MPSNNReductionColumnMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionColumnMaxNode implements IMPSNNReductionColumnMaxNode.
var _ IMPSNNReductionColumnMaxNode = MPSNNReductionColumnMaxNode{}





// An interface definition for the [MPSNNReductionColumnMaxNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionColumnMaxNode
type IMPSNNReductionColumnMaxNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionColumnMaxNode) Init() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMaxNode) Autorelease() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMaxNode creates a new MPSNNReductionColumnMaxNode instance.
func NewMPSNNReductionColumnMaxNode() MPSNNReductionColumnMaxNode {
	class := getMPSNNReductionColumnMaxNodeClass()
	rv := objc.Send[MPSNNReductionColumnMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionColumnMaxNodeWithSource(source objectivec.IObject) MPSNNReductionColumnMaxNode {
	instance := getMPSNNReductionColumnMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionColumnMaxNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionColumnMaxNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionColumnMaxNode {
	instance := getMPSNNReductionColumnMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionColumnMaxNodeFromID(rv)
}
































