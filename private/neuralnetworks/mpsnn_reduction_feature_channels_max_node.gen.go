// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionFeatureChannelsMaxNode] class.
var (
	_MPSNNReductionFeatureChannelsMaxNodeClass     MPSNNReductionFeatureChannelsMaxNodeClass
	_MPSNNReductionFeatureChannelsMaxNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsMaxNodeClass() MPSNNReductionFeatureChannelsMaxNodeClass {
	_MPSNNReductionFeatureChannelsMaxNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsMaxNodeClass = MPSNNReductionFeatureChannelsMaxNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsMaxNode")}
	})
	return _MPSNNReductionFeatureChannelsMaxNodeClass
}

// GetMPSNNReductionFeatureChannelsMaxNodeClass returns the class object for MPSNNReductionFeatureChannelsMaxNode.
func GetMPSNNReductionFeatureChannelsMaxNodeClass() MPSNNReductionFeatureChannelsMaxNodeClass {
	return getMPSNNReductionFeatureChannelsMaxNodeClass()
}

type MPSNNReductionFeatureChannelsMaxNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsMaxNodeClass) Alloc() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMaxNode
type MPSNNReductionFeatureChannelsMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsMaxNodeFromID constructs a [MPSNNReductionFeatureChannelsMaxNode] from an objc.ID.
func MPSNNReductionFeatureChannelsMaxNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsMaxNode {
	return MPSNNReductionFeatureChannelsMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionFeatureChannelsMaxNode implements IMPSNNReductionFeatureChannelsMaxNode.
var _ IMPSNNReductionFeatureChannelsMaxNode = MPSNNReductionFeatureChannelsMaxNode{}





// An interface definition for the [MPSNNReductionFeatureChannelsMaxNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMaxNode
type IMPSNNReductionFeatureChannelsMaxNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsMaxNode) Init() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsMaxNode) Autorelease() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsMaxNode creates a new MPSNNReductionFeatureChannelsMaxNode instance.
func NewMPSNNReductionFeatureChannelsMaxNode() MPSNNReductionFeatureChannelsMaxNode {
	class := getMPSNNReductionFeatureChannelsMaxNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionFeatureChannelsMaxNodeWithSource(source objectivec.IObject) MPSNNReductionFeatureChannelsMaxNode {
	instance := getMPSNNReductionFeatureChannelsMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionFeatureChannelsMaxNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionFeatureChannelsMaxNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionFeatureChannelsMaxNode {
	instance := getMPSNNReductionFeatureChannelsMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionFeatureChannelsMaxNodeFromID(rv)
}
































