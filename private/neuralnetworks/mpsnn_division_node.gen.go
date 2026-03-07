// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNDivisionNode] class.
var (
	_MPSNNDivisionNodeClass     MPSNNDivisionNodeClass
	_MPSNNDivisionNodeClassOnce sync.Once
)

func getMPSNNDivisionNodeClass() MPSNNDivisionNodeClass {
	_MPSNNDivisionNodeClassOnce.Do(func() {
		_MPSNNDivisionNodeClass = MPSNNDivisionNodeClass{class: objc.GetClass("MPSNNDivisionNode")}
	})
	return _MPSNNDivisionNodeClass
}

// GetMPSNNDivisionNodeClass returns the class object for MPSNNDivisionNode.
func GetMPSNNDivisionNodeClass() MPSNNDivisionNodeClass {
	return getMPSNNDivisionNodeClass()
}

type MPSNNDivisionNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNDivisionNodeClass) Alloc() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDivisionNode
type MPSNNDivisionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNDivisionNodeFromID constructs a [MPSNNDivisionNode] from an objc.ID.
func MPSNNDivisionNodeFromID(id objc.ID) MPSNNDivisionNode {
	return MPSNNDivisionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}
// Ensure MPSNNDivisionNode implements IMPSNNDivisionNode.
var _ IMPSNNDivisionNode = MPSNNDivisionNode{}





// An interface definition for the [MPSNNDivisionNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDivisionNode
type IMPSNNDivisionNode interface {
	IMPSNNBinaryArithmeticNode
}





// Init initializes the instance.
func (d MPSNNDivisionNode) Init() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MPSNNDivisionNode) Autorelease() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNDivisionNode creates a new MPSNNDivisionNode instance.
func NewMPSNNDivisionNode() MPSNNDivisionNode {
	class := getMPSNNDivisionNodeClass()
	rv := objc.Send[MPSNNDivisionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewDivisionNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNDivisionNode {
	instance := getMPSNNDivisionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNDivisionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewDivisionNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNDivisionNode {
	instance := getMPSNNDivisionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNDivisionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewDivisionNodeWithSources(sources objectivec.IObject) MPSNNDivisionNode {
	instance := getMPSNNDivisionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNDivisionNodeFromID(rv)
}
































