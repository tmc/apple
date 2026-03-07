// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionFeatureChannelsMeanNode] class.
var (
	_MPSNNReductionFeatureChannelsMeanNodeClass     MPSNNReductionFeatureChannelsMeanNodeClass
	_MPSNNReductionFeatureChannelsMeanNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsMeanNodeClass() MPSNNReductionFeatureChannelsMeanNodeClass {
	_MPSNNReductionFeatureChannelsMeanNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsMeanNodeClass = MPSNNReductionFeatureChannelsMeanNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsMeanNode")}
	})
	return _MPSNNReductionFeatureChannelsMeanNodeClass
}

// GetMPSNNReductionFeatureChannelsMeanNodeClass returns the class object for MPSNNReductionFeatureChannelsMeanNode.
func GetMPSNNReductionFeatureChannelsMeanNodeClass() MPSNNReductionFeatureChannelsMeanNodeClass {
	return getMPSNNReductionFeatureChannelsMeanNodeClass()
}

type MPSNNReductionFeatureChannelsMeanNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsMeanNodeClass) Alloc() MPSNNReductionFeatureChannelsMeanNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMeanNode
type MPSNNReductionFeatureChannelsMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsMeanNodeFromID constructs a [MPSNNReductionFeatureChannelsMeanNode] from an objc.ID.
func MPSNNReductionFeatureChannelsMeanNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsMeanNode {
	return MPSNNReductionFeatureChannelsMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionFeatureChannelsMeanNode implements IMPSNNReductionFeatureChannelsMeanNode.
var _ IMPSNNReductionFeatureChannelsMeanNode = MPSNNReductionFeatureChannelsMeanNode{}





// An interface definition for the [MPSNNReductionFeatureChannelsMeanNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMeanNode
type IMPSNNReductionFeatureChannelsMeanNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsMeanNode) Init() MPSNNReductionFeatureChannelsMeanNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsMeanNode) Autorelease() MPSNNReductionFeatureChannelsMeanNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsMeanNode creates a new MPSNNReductionFeatureChannelsMeanNode instance.
func NewMPSNNReductionFeatureChannelsMeanNode() MPSNNReductionFeatureChannelsMeanNode {
	class := getMPSNNReductionFeatureChannelsMeanNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionFeatureChannelsMeanNodeWithSource(source objectivec.IObject) MPSNNReductionFeatureChannelsMeanNode {
	instance := getMPSNNReductionFeatureChannelsMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionFeatureChannelsMeanNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionFeatureChannelsMeanNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionFeatureChannelsMeanNode {
	instance := getMPSNNReductionFeatureChannelsMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionFeatureChannelsMeanNodeFromID(rv)
}
































