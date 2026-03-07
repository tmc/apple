// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionFeatureChannelsArgumentMaxNode] class.
var (
	_MPSNNReductionFeatureChannelsArgumentMaxNodeClass     MPSNNReductionFeatureChannelsArgumentMaxNodeClass
	_MPSNNReductionFeatureChannelsArgumentMaxNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsArgumentMaxNodeClass() MPSNNReductionFeatureChannelsArgumentMaxNodeClass {
	_MPSNNReductionFeatureChannelsArgumentMaxNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsArgumentMaxNodeClass = MPSNNReductionFeatureChannelsArgumentMaxNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsArgumentMaxNode")}
	})
	return _MPSNNReductionFeatureChannelsArgumentMaxNodeClass
}

// GetMPSNNReductionFeatureChannelsArgumentMaxNodeClass returns the class object for MPSNNReductionFeatureChannelsArgumentMaxNode.
func GetMPSNNReductionFeatureChannelsArgumentMaxNodeClass() MPSNNReductionFeatureChannelsArgumentMaxNodeClass {
	return getMPSNNReductionFeatureChannelsArgumentMaxNodeClass()
}

type MPSNNReductionFeatureChannelsArgumentMaxNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsArgumentMaxNodeClass) Alloc() MPSNNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsArgumentMaxNode
type MPSNNReductionFeatureChannelsArgumentMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsArgumentMaxNodeFromID constructs a [MPSNNReductionFeatureChannelsArgumentMaxNode] from an objc.ID.
func MPSNNReductionFeatureChannelsArgumentMaxNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsArgumentMaxNode {
	return MPSNNReductionFeatureChannelsArgumentMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionFeatureChannelsArgumentMaxNode implements IMPSNNReductionFeatureChannelsArgumentMaxNode.
var _ IMPSNNReductionFeatureChannelsArgumentMaxNode = MPSNNReductionFeatureChannelsArgumentMaxNode{}





// An interface definition for the [MPSNNReductionFeatureChannelsArgumentMaxNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsArgumentMaxNode
type IMPSNNReductionFeatureChannelsArgumentMaxNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsArgumentMaxNode) Init() MPSNNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsArgumentMaxNode) Autorelease() MPSNNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsArgumentMaxNode creates a new MPSNNReductionFeatureChannelsArgumentMaxNode instance.
func NewMPSNNReductionFeatureChannelsArgumentMaxNode() MPSNNReductionFeatureChannelsArgumentMaxNode {
	class := getMPSNNReductionFeatureChannelsArgumentMaxNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionFeatureChannelsArgumentMaxNodeWithSource(source objectivec.IObject) MPSNNReductionFeatureChannelsArgumentMaxNode {
	instance := getMPSNNReductionFeatureChannelsArgumentMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionFeatureChannelsArgumentMaxNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionFeatureChannelsArgumentMaxNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionFeatureChannelsArgumentMaxNode {
	instance := getMPSNNReductionFeatureChannelsArgumentMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionFeatureChannelsArgumentMaxNodeFromID(rv)
}
































