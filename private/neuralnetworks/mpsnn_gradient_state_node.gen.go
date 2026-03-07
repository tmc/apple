// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGradientStateNode] class.
var (
	_MPSNNGradientStateNodeClass     MPSNNGradientStateNodeClass
	_MPSNNGradientStateNodeClassOnce sync.Once
)

func getMPSNNGradientStateNodeClass() MPSNNGradientStateNodeClass {
	_MPSNNGradientStateNodeClassOnce.Do(func() {
		_MPSNNGradientStateNodeClass = MPSNNGradientStateNodeClass{class: objc.GetClass("MPSNNGradientStateNode")}
	})
	return _MPSNNGradientStateNodeClass
}

// GetMPSNNGradientStateNodeClass returns the class object for MPSNNGradientStateNode.
func GetMPSNNGradientStateNodeClass() MPSNNGradientStateNodeClass {
	return getMPSNNGradientStateNodeClass()
}

type MPSNNGradientStateNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientStateNodeClass) Alloc() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientStateNode
type MPSNNGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNGradientStateNodeFromID constructs a [MPSNNGradientStateNode] from an objc.ID.
func MPSNNGradientStateNodeFromID(id objc.ID) MPSNNGradientStateNode {
	return MPSNNGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}
// Ensure MPSNNGradientStateNode implements IMPSNNGradientStateNode.
var _ IMPSNNGradientStateNode = MPSNNGradientStateNode{}





// An interface definition for the [MPSNNGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientStateNode
type IMPSNNGradientStateNode interface {
	IMPSNNStateNode
}





// Init initializes the instance.
func (g MPSNNGradientStateNode) Init() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientStateNode) Autorelease() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientStateNode creates a new MPSNNGradientStateNode instance.
func NewMPSNNGradientStateNode() MPSNNGradientStateNode {
	class := getMPSNNGradientStateNodeClass()
	rv := objc.Send[MPSNNGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewGradientStateNodeWithParent(parent objectivec.IObject) MPSNNGradientStateNode {
	instance := getMPSNNGradientStateNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNGradientStateNodeFromID(rv)
}
































