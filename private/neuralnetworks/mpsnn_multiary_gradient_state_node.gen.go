// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNMultiaryGradientStateNode] class.
var (
	_MPSNNMultiaryGradientStateNodeClass     MPSNNMultiaryGradientStateNodeClass
	_MPSNNMultiaryGradientStateNodeClassOnce sync.Once
)

func getMPSNNMultiaryGradientStateNodeClass() MPSNNMultiaryGradientStateNodeClass {
	_MPSNNMultiaryGradientStateNodeClassOnce.Do(func() {
		_MPSNNMultiaryGradientStateNodeClass = MPSNNMultiaryGradientStateNodeClass{class: objc.GetClass("MPSNNMultiaryGradientStateNode")}
	})
	return _MPSNNMultiaryGradientStateNodeClass
}

// GetMPSNNMultiaryGradientStateNodeClass returns the class object for MPSNNMultiaryGradientStateNode.
func GetMPSNNMultiaryGradientStateNodeClass() MPSNNMultiaryGradientStateNodeClass {
	return getMPSNNMultiaryGradientStateNodeClass()
}

type MPSNNMultiaryGradientStateNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiaryGradientStateNodeClass) Alloc() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientStateNode
type MPSNNMultiaryGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNMultiaryGradientStateNodeFromID constructs a [MPSNNMultiaryGradientStateNode] from an objc.ID.
func MPSNNMultiaryGradientStateNodeFromID(id objc.ID) MPSNNMultiaryGradientStateNode {
	return MPSNNMultiaryGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}
// Ensure MPSNNMultiaryGradientStateNode implements IMPSNNMultiaryGradientStateNode.
var _ IMPSNNMultiaryGradientStateNode = MPSNNMultiaryGradientStateNode{}





// An interface definition for the [MPSNNMultiaryGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNMultiaryGradientStateNode
type IMPSNNMultiaryGradientStateNode interface {
	IMPSNNStateNode
}





// Init initializes the instance.
func (m MPSNNMultiaryGradientStateNode) Init() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiaryGradientStateNode) Autorelease() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiaryGradientStateNode creates a new MPSNNMultiaryGradientStateNode instance.
func NewMPSNNMultiaryGradientStateNode() MPSNNMultiaryGradientStateNode {
	class := getMPSNNMultiaryGradientStateNodeClass()
	rv := objc.Send[MPSNNMultiaryGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewMultiaryGradientStateNodeWithParent(parent objectivec.IObject) MPSNNMultiaryGradientStateNode {
	instance := getMPSNNMultiaryGradientStateNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNMultiaryGradientStateNodeFromID(rv)
}
































