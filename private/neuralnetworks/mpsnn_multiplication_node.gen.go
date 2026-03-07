// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNMultiplicationNode] class.
var (
	_MPSNNMultiplicationNodeClass     MPSNNMultiplicationNodeClass
	_MPSNNMultiplicationNodeClassOnce sync.Once
)

func getMPSNNMultiplicationNodeClass() MPSNNMultiplicationNodeClass {
	_MPSNNMultiplicationNodeClassOnce.Do(func() {
		_MPSNNMultiplicationNodeClass = MPSNNMultiplicationNodeClass{class: objc.GetClass("MPSNNMultiplicationNode")}
	})
	return _MPSNNMultiplicationNodeClass
}

// GetMPSNNMultiplicationNodeClass returns the class object for MPSNNMultiplicationNode.
func GetMPSNNMultiplicationNodeClass() MPSNNMultiplicationNodeClass {
	return getMPSNNMultiplicationNodeClass()
}

type MPSNNMultiplicationNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiplicationNodeClass) Alloc() MPSNNMultiplicationNode {
	rv := objc.Send[MPSNNMultiplicationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiplicationNode
type MPSNNMultiplicationNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNMultiplicationNodeFromID constructs a [MPSNNMultiplicationNode] from an objc.ID.
func MPSNNMultiplicationNodeFromID(id objc.ID) MPSNNMultiplicationNode {
	return MPSNNMultiplicationNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}
// Ensure MPSNNMultiplicationNode implements IMPSNNMultiplicationNode.
var _ IMPSNNMultiplicationNode = MPSNNMultiplicationNode{}





// An interface definition for the [MPSNNMultiplicationNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiplicationNode
type IMPSNNMultiplicationNode interface {
	IMPSNNBinaryArithmeticNode
}





// Init initializes the instance.
func (m MPSNNMultiplicationNode) Init() MPSNNMultiplicationNode {
	rv := objc.Send[MPSNNMultiplicationNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiplicationNode) Autorelease() MPSNNMultiplicationNode {
	rv := objc.Send[MPSNNMultiplicationNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiplicationNode creates a new MPSNNMultiplicationNode instance.
func NewMPSNNMultiplicationNode() MPSNNMultiplicationNode {
	class := getMPSNNMultiplicationNodeClass()
	rv := objc.Send[MPSNNMultiplicationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewMultiplicationNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNMultiplicationNode {
	instance := getMPSNNMultiplicationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNMultiplicationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewMultiplicationNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNMultiplicationNode {
	instance := getMPSNNMultiplicationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNMultiplicationNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewMultiplicationNodeWithSources(sources objectivec.IObject) MPSNNMultiplicationNode {
	instance := getMPSNNMultiplicationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNMultiplicationNodeFromID(rv)
}
































