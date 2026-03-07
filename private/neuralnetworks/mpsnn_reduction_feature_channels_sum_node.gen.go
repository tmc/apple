// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReductionFeatureChannelsSumNode] class.
var (
	_MPSNNReductionFeatureChannelsSumNodeClass     MPSNNReductionFeatureChannelsSumNodeClass
	_MPSNNReductionFeatureChannelsSumNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsSumNodeClass() MPSNNReductionFeatureChannelsSumNodeClass {
	_MPSNNReductionFeatureChannelsSumNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsSumNodeClass = MPSNNReductionFeatureChannelsSumNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsSumNode")}
	})
	return _MPSNNReductionFeatureChannelsSumNodeClass
}

// GetMPSNNReductionFeatureChannelsSumNodeClass returns the class object for MPSNNReductionFeatureChannelsSumNode.
func GetMPSNNReductionFeatureChannelsSumNodeClass() MPSNNReductionFeatureChannelsSumNodeClass {
	return getMPSNNReductionFeatureChannelsSumNodeClass()
}

type MPSNNReductionFeatureChannelsSumNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsSumNodeClass) Alloc() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReductionFeatureChannelsSumNode.Weight]
//   - [MPSNNReductionFeatureChannelsSumNode.SetWeight]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsSumNode
type MPSNNReductionFeatureChannelsSumNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsSumNodeFromID constructs a [MPSNNReductionFeatureChannelsSumNode] from an objc.ID.
func MPSNNReductionFeatureChannelsSumNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsSumNode {
	return MPSNNReductionFeatureChannelsSumNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}
// Ensure MPSNNReductionFeatureChannelsSumNode implements IMPSNNReductionFeatureChannelsSumNode.
var _ IMPSNNReductionFeatureChannelsSumNode = MPSNNReductionFeatureChannelsSumNode{}





// An interface definition for the [MPSNNReductionFeatureChannelsSumNode] class.
//
// # Methods
//
//   - [IMPSNNReductionFeatureChannelsSumNode.Weight]
//   - [IMPSNNReductionFeatureChannelsSumNode.SetWeight]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsSumNode
type IMPSNNReductionFeatureChannelsSumNode interface {
	IMPSNNUnaryReductionNode

	// Topic: Methods

	Weight() float32
	SetWeight(value float32)
}





// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsSumNode) Init() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsSumNode) Autorelease() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsSumNode creates a new MPSNNReductionFeatureChannelsSumNode instance.
func NewMPSNNReductionFeatureChannelsSumNode() MPSNNReductionFeatureChannelsSumNode {
	class := getMPSNNReductionFeatureChannelsSumNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNUnaryReductionNode/initWithSource:
func NewReductionFeatureChannelsSumNodeWithSource(source objectivec.IObject) MPSNNReductionFeatureChannelsSumNode {
	instance := getMPSNNReductionFeatureChannelsSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNReductionFeatureChannelsSumNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReductionFeatureChannelsSumNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReductionFeatureChannelsSumNode {
	instance := getMPSNNReductionFeatureChannelsSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReductionFeatureChannelsSumNodeFromID(rv)
}


















// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReductionFeatureChannelsSumNode/weight
func (r MPSNNReductionFeatureChannelsSumNode) Weight() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("weight"))
	return rv
}
func (r MPSNNReductionFeatureChannelsSumNode) SetWeight(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setWeight:"), value)
}

















