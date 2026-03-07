// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNBinaryGradientStateNode] class.
var (
	_MPSNNBinaryGradientStateNodeClass     MPSNNBinaryGradientStateNodeClass
	_MPSNNBinaryGradientStateNodeClassOnce sync.Once
)

func getMPSNNBinaryGradientStateNodeClass() MPSNNBinaryGradientStateNodeClass {
	_MPSNNBinaryGradientStateNodeClassOnce.Do(func() {
		_MPSNNBinaryGradientStateNodeClass = MPSNNBinaryGradientStateNodeClass{class: objc.GetClass("MPSNNBinaryGradientStateNode")}
	})
	return _MPSNNBinaryGradientStateNodeClass
}

// GetMPSNNBinaryGradientStateNodeClass returns the class object for MPSNNBinaryGradientStateNode.
func GetMPSNNBinaryGradientStateNodeClass() MPSNNBinaryGradientStateNodeClass {
	return getMPSNNBinaryGradientStateNodeClass()
}

type MPSNNBinaryGradientStateNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryGradientStateNodeClass) Alloc() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryGradientStateNode
type MPSNNBinaryGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNBinaryGradientStateNodeFromID constructs a [MPSNNBinaryGradientStateNode] from an objc.ID.
func MPSNNBinaryGradientStateNodeFromID(id objc.ID) MPSNNBinaryGradientStateNode {
	return MPSNNBinaryGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}
// Ensure MPSNNBinaryGradientStateNode implements IMPSNNBinaryGradientStateNode.
var _ IMPSNNBinaryGradientStateNode = MPSNNBinaryGradientStateNode{}





// An interface definition for the [MPSNNBinaryGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryGradientStateNode
type IMPSNNBinaryGradientStateNode interface {
	IMPSNNStateNode
}





// Init initializes the instance.
func (b MPSNNBinaryGradientStateNode) Init() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBinaryGradientStateNode) Autorelease() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBinaryGradientStateNode creates a new MPSNNBinaryGradientStateNode instance.
func NewMPSNNBinaryGradientStateNode() MPSNNBinaryGradientStateNode {
	class := getMPSNNBinaryGradientStateNodeClass()
	rv := objc.Send[MPSNNBinaryGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewBinaryGradientStateNodeWithParent(parent objectivec.IObject) MPSNNBinaryGradientStateNode {
	instance := getMPSNNBinaryGradientStateNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNBinaryGradientStateNodeFromID(rv)
}
































