// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionFeatureChannelsMinNode] class.
var (
	_MPSNNReductionFeatureChannelsMinNodeClass     MPSNNReductionFeatureChannelsMinNodeClass
	_MPSNNReductionFeatureChannelsMinNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsMinNodeClass() MPSNNReductionFeatureChannelsMinNodeClass {
	_MPSNNReductionFeatureChannelsMinNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsMinNodeClass = MPSNNReductionFeatureChannelsMinNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsMinNode")}
	})
	return _MPSNNReductionFeatureChannelsMinNodeClass
}

// GetMPSNNReductionFeatureChannelsMinNodeClass returns the class object for MPSNNReductionFeatureChannelsMinNode.
func GetMPSNNReductionFeatureChannelsMinNodeClass() MPSNNReductionFeatureChannelsMinNodeClass {
	return getMPSNNReductionFeatureChannelsMinNodeClass()
}

type MPSNNReductionFeatureChannelsMinNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsMinNodeClass) Alloc() MPSNNReductionFeatureChannelsMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMinNode
type MPSNNReductionFeatureChannelsMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsMinNodeFromID constructs a [MPSNNReductionFeatureChannelsMinNode] from an objc.ID.
func MPSNNReductionFeatureChannelsMinNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsMinNode {
	return MPSNNReductionFeatureChannelsMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionFeatureChannelsMinNode implements IMPSNNReductionFeatureChannelsMinNode.
var _ IMPSNNReductionFeatureChannelsMinNode = MPSNNReductionFeatureChannelsMinNode{}





// An interface definition for the [MPSNNReductionFeatureChannelsMinNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsMinNode
type IMPSNNReductionFeatureChannelsMinNode interface {
	IMPSNNUnaryReductionNode
}





// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsMinNode) Init() MPSNNReductionFeatureChannelsMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsMinNode) Autorelease() MPSNNReductionFeatureChannelsMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsMinNode creates a new MPSNNReductionFeatureChannelsMinNode instance.
func NewMPSNNReductionFeatureChannelsMinNode() MPSNNReductionFeatureChannelsMinNode {
	class := getMPSNNReductionFeatureChannelsMinNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionFeatureChannelsMinNodeWithSource(source objectivec.IObject) MPSNNReductionFeatureChannelsMinNode {
	instance := getMPSNNReductionFeatureChannelsMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionFeatureChannelsMinNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionFeatureChannelsMinNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionFeatureChannelsMinNode {
	instance := getMPSNNReductionFeatureChannelsMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionFeatureChannelsMinNodeFromID(rv)
}
































