// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNArithmeticGradientStateNode] class.
var (
	_MPSNNArithmeticGradientStateNodeClass     MPSNNArithmeticGradientStateNodeClass
	_MPSNNArithmeticGradientStateNodeClassOnce sync.Once
)

func getMPSNNArithmeticGradientStateNodeClass() MPSNNArithmeticGradientStateNodeClass {
	_MPSNNArithmeticGradientStateNodeClassOnce.Do(func() {
		_MPSNNArithmeticGradientStateNodeClass = MPSNNArithmeticGradientStateNodeClass{class: objc.GetClass("MPSNNArithmeticGradientStateNode")}
	})
	return _MPSNNArithmeticGradientStateNodeClass
}

// GetMPSNNArithmeticGradientStateNodeClass returns the class object for MPSNNArithmeticGradientStateNode.
func GetMPSNNArithmeticGradientStateNodeClass() MPSNNArithmeticGradientStateNodeClass {
	return getMPSNNArithmeticGradientStateNodeClass()
}

type MPSNNArithmeticGradientStateNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNArithmeticGradientStateNodeClass) Alloc() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientStateNode
type MPSNNArithmeticGradientStateNode struct {
	MPSNNBinaryGradientStateNode
}

// MPSNNArithmeticGradientStateNodeFromID constructs a [MPSNNArithmeticGradientStateNode] from an objc.ID.
func MPSNNArithmeticGradientStateNodeFromID(id objc.ID) MPSNNArithmeticGradientStateNode {
	return MPSNNArithmeticGradientStateNode{MPSNNBinaryGradientStateNode: MPSNNBinaryGradientStateNodeFromID(id)}
}
// Ensure MPSNNArithmeticGradientStateNode implements IMPSNNArithmeticGradientStateNode.
var _ IMPSNNArithmeticGradientStateNode = MPSNNArithmeticGradientStateNode{}





// An interface definition for the [MPSNNArithmeticGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientStateNode
type IMPSNNArithmeticGradientStateNode interface {
	IMPSNNBinaryGradientStateNode
}





// Init initializes the instance.
func (a MPSNNArithmeticGradientStateNode) Init() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNArithmeticGradientStateNode) Autorelease() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNArithmeticGradientStateNode creates a new MPSNNArithmeticGradientStateNode instance.
func NewMPSNNArithmeticGradientStateNode() MPSNNArithmeticGradientStateNode {
	class := getMPSNNArithmeticGradientStateNodeClass()
	rv := objc.Send[MPSNNArithmeticGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewArithmeticGradientStateNodeWithParent(parent objectivec.IObject) MPSNNArithmeticGradientStateNode {
	instance := getMPSNNArithmeticGradientStateNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNArithmeticGradientStateNodeFromID(rv)
}
































