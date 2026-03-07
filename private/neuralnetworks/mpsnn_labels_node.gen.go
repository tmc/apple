// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLabelsNode] class.
var (
	_MPSNNLabelsNodeClass     MPSNNLabelsNodeClass
	_MPSNNLabelsNodeClassOnce sync.Once
)

func getMPSNNLabelsNodeClass() MPSNNLabelsNodeClass {
	_MPSNNLabelsNodeClassOnce.Do(func() {
		_MPSNNLabelsNodeClass = MPSNNLabelsNodeClass{class: objc.GetClass("MPSNNLabelsNode")}
	})
	return _MPSNNLabelsNodeClass
}

// GetMPSNNLabelsNodeClass returns the class object for MPSNNLabelsNode.
func GetMPSNNLabelsNodeClass() MPSNNLabelsNodeClass {
	return getMPSNNLabelsNodeClass()
}

type MPSNNLabelsNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLabelsNodeClass) Alloc() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLabelsNode
type MPSNNLabelsNode struct {
	MPSNNStateNode
}

// MPSNNLabelsNodeFromID constructs a [MPSNNLabelsNode] from an objc.ID.
func MPSNNLabelsNodeFromID(id objc.ID) MPSNNLabelsNode {
	return MPSNNLabelsNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}
// Ensure MPSNNLabelsNode implements IMPSNNLabelsNode.
var _ IMPSNNLabelsNode = MPSNNLabelsNode{}





// An interface definition for the [MPSNNLabelsNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLabelsNode
type IMPSNNLabelsNode interface {
	IMPSNNStateNode
}





// Init initializes the instance.
func (l MPSNNLabelsNode) Init() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLabelsNode) Autorelease() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLabelsNode creates a new MPSNNLabelsNode instance.
func NewMPSNNLabelsNode() MPSNNLabelsNode {
	class := getMPSNNLabelsNodeClass()
	rv := objc.Send[MPSNNLabelsNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewLabelsNodeWithParent(parent objectivec.IObject) MPSNNLabelsNode {
	instance := getMPSNNLabelsNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNLabelsNodeFromID(rv)
}
































