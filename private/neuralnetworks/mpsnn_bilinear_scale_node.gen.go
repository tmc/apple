// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNBilinearScaleNode] class.
var (
	_MPSNNBilinearScaleNodeClass     MPSNNBilinearScaleNodeClass
	_MPSNNBilinearScaleNodeClassOnce sync.Once
)

func getMPSNNBilinearScaleNodeClass() MPSNNBilinearScaleNodeClass {
	_MPSNNBilinearScaleNodeClassOnce.Do(func() {
		_MPSNNBilinearScaleNodeClass = MPSNNBilinearScaleNodeClass{class: objc.GetClass("MPSNNBilinearScaleNode")}
	})
	return _MPSNNBilinearScaleNodeClass
}

// GetMPSNNBilinearScaleNodeClass returns the class object for MPSNNBilinearScaleNode.
func GetMPSNNBilinearScaleNodeClass() MPSNNBilinearScaleNodeClass {
	return getMPSNNBilinearScaleNodeClass()
}

type MPSNNBilinearScaleNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBilinearScaleNodeClass) Alloc() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBilinearScaleNode
type MPSNNBilinearScaleNode struct {
	MPSNNScaleNode
}

// MPSNNBilinearScaleNodeFromID constructs a [MPSNNBilinearScaleNode] from an objc.ID.
func MPSNNBilinearScaleNodeFromID(id objc.ID) MPSNNBilinearScaleNode {
	return MPSNNBilinearScaleNode{MPSNNScaleNode: MPSNNScaleNodeFromID(id)}
}
// Ensure MPSNNBilinearScaleNode implements IMPSNNBilinearScaleNode.
var _ IMPSNNBilinearScaleNode = MPSNNBilinearScaleNode{}





// An interface definition for the [MPSNNBilinearScaleNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBilinearScaleNode
type IMPSNNBilinearScaleNode interface {
	IMPSNNScaleNode
}





// Init initializes the instance.
func (b MPSNNBilinearScaleNode) Init() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBilinearScaleNode) Autorelease() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBilinearScaleNode creates a new MPSNNBilinearScaleNode instance.
func NewMPSNNBilinearScaleNode() MPSNNBilinearScaleNode {
	class := getMPSNNBilinearScaleNodeClass()
	rv := objc.Send[MPSNNBilinearScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewBilinearScaleNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNBilinearScaleNode {
	instance := getMPSNNBilinearScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNBilinearScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:outputSize:
func NewBilinearScaleNodeWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) MPSNNBilinearScaleNode {
	instance := getMPSNNBilinearScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), source, size)
	return MPSNNBilinearScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBilinearScaleNode/initWithSource:transformProvider:outputSize:
func NewBilinearScaleNodeWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) MPSNNBilinearScaleNode {
	instance := getMPSNNBilinearScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), source, provider, size)
	return MPSNNBilinearScaleNodeFromID(rv)
}
































