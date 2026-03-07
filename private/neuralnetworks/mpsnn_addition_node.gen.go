// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNAdditionNode] class.
var (
	_MPSNNAdditionNodeClass     MPSNNAdditionNodeClass
	_MPSNNAdditionNodeClassOnce sync.Once
)

func getMPSNNAdditionNodeClass() MPSNNAdditionNodeClass {
	_MPSNNAdditionNodeClassOnce.Do(func() {
		_MPSNNAdditionNodeClass = MPSNNAdditionNodeClass{class: objc.GetClass("MPSNNAdditionNode")}
	})
	return _MPSNNAdditionNodeClass
}

// GetMPSNNAdditionNodeClass returns the class object for MPSNNAdditionNode.
func GetMPSNNAdditionNodeClass() MPSNNAdditionNodeClass {
	return getMPSNNAdditionNodeClass()
}

type MPSNNAdditionNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNAdditionNodeClass) Alloc() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNAdditionNode
type MPSNNAdditionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNAdditionNodeFromID constructs a [MPSNNAdditionNode] from an objc.ID.
func MPSNNAdditionNodeFromID(id objc.ID) MPSNNAdditionNode {
	return MPSNNAdditionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}
// Ensure MPSNNAdditionNode implements IMPSNNAdditionNode.
var _ IMPSNNAdditionNode = MPSNNAdditionNode{}





// An interface definition for the [MPSNNAdditionNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNAdditionNode
type IMPSNNAdditionNode interface {
	IMPSNNBinaryArithmeticNode
}





// Init initializes the instance.
func (a MPSNNAdditionNode) Init() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNAdditionNode) Autorelease() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNAdditionNode creates a new MPSNNAdditionNode instance.
func NewMPSNNAdditionNode() MPSNNAdditionNode {
	class := getMPSNNAdditionNodeClass()
	rv := objc.Send[MPSNNAdditionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewAdditionNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNAdditionNode {
	instance := getMPSNNAdditionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNAdditionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewAdditionNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNAdditionNode {
	instance := getMPSNNAdditionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNAdditionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewAdditionNodeWithSources(sources objectivec.IObject) MPSNNAdditionNode {
	instance := getMPSNNAdditionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNAdditionNodeFromID(rv)
}
































