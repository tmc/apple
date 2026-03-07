// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionSpatialMeanNode] class.
var (
	_MPSNNReductionSpatialMeanNodeClass     MPSNNReductionSpatialMeanNodeClass
	_MPSNNReductionSpatialMeanNodeClassOnce sync.Once
)

func getMPSNNReductionSpatialMeanNodeClass() MPSNNReductionSpatialMeanNodeClass {
	_MPSNNReductionSpatialMeanNodeClassOnce.Do(func() {
		_MPSNNReductionSpatialMeanNodeClass = MPSNNReductionSpatialMeanNodeClass{class: objc.GetClass("MPSNNReductionSpatialMeanNode")}
	})
	return _MPSNNReductionSpatialMeanNodeClass
}

// GetMPSNNReductionSpatialMeanNodeClass returns the class object for MPSNNReductionSpatialMeanNode.
func GetMPSNNReductionSpatialMeanNodeClass() MPSNNReductionSpatialMeanNodeClass {
	return getMPSNNReductionSpatialMeanNodeClass()
}

type MPSNNReductionSpatialMeanNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionSpatialMeanNodeClass) Alloc() MPSNNReductionSpatialMeanNode {
	rv := objc.Send[MPSNNReductionSpatialMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanNode
type MPSNNReductionSpatialMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionSpatialMeanNodeFromID constructs a [MPSNNReductionSpatialMeanNode] from an objc.ID.
func MPSNNReductionSpatialMeanNodeFromID(id objc.ID) MPSNNReductionSpatialMeanNode {
	return MPSNNReductionSpatialMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionSpatialMeanNode implements IMPSNNReductionSpatialMeanNode.
var _ IMPSNNReductionSpatialMeanNode = MPSNNReductionSpatialMeanNode{}





// An interface definition for the [MPSNNReductionSpatialMeanNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionSpatialMeanNode
type IMPSNNReductionSpatialMeanNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionSpatialMeanNode) Init() MPSNNReductionSpatialMeanNode {
	rv := objc.Send[MPSNNReductionSpatialMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionSpatialMeanNode) Autorelease() MPSNNReductionSpatialMeanNode {
	rv := objc.Send[MPSNNReductionSpatialMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionSpatialMeanNode creates a new MPSNNReductionSpatialMeanNode instance.
func NewMPSNNReductionSpatialMeanNode() MPSNNReductionSpatialMeanNode {
	class := getMPSNNReductionSpatialMeanNodeClass()
	rv := objc.Send[MPSNNReductionSpatialMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionSpatialMeanNodeWithSource(source objectivec.IObject) MPSNNReductionSpatialMeanNode {
	instance := getMPSNNReductionSpatialMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionSpatialMeanNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionSpatialMeanNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionSpatialMeanNode {
	instance := getMPSNNReductionSpatialMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionSpatialMeanNodeFromID(rv)
}
































