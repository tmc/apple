// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLanczosScaleNode] class.
var (
	_MPSNNLanczosScaleNodeClass     MPSNNLanczosScaleNodeClass
	_MPSNNLanczosScaleNodeClassOnce sync.Once
)

func getMPSNNLanczosScaleNodeClass() MPSNNLanczosScaleNodeClass {
	_MPSNNLanczosScaleNodeClassOnce.Do(func() {
		_MPSNNLanczosScaleNodeClass = MPSNNLanczosScaleNodeClass{class: objc.GetClass("MPSNNLanczosScaleNode")}
	})
	return _MPSNNLanczosScaleNodeClass
}

// GetMPSNNLanczosScaleNodeClass returns the class object for MPSNNLanczosScaleNode.
func GetMPSNNLanczosScaleNodeClass() MPSNNLanczosScaleNodeClass {
	return getMPSNNLanczosScaleNodeClass()
}

type MPSNNLanczosScaleNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLanczosScaleNodeClass) Alloc() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLanczosScaleNode
type MPSNNLanczosScaleNode struct {
	MPSNNScaleNode
}

// MPSNNLanczosScaleNodeFromID constructs a [MPSNNLanczosScaleNode] from an objc.ID.
func MPSNNLanczosScaleNodeFromID(id objc.ID) MPSNNLanczosScaleNode {
	return MPSNNLanczosScaleNode{MPSNNScaleNode: MPSNNScaleNodeFromID(id)}
}
// Ensure MPSNNLanczosScaleNode implements IMPSNNLanczosScaleNode.
var _ IMPSNNLanczosScaleNode = MPSNNLanczosScaleNode{}





// An interface definition for the [MPSNNLanczosScaleNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLanczosScaleNode
type IMPSNNLanczosScaleNode interface {
	IMPSNNScaleNode
}





// Init initializes the instance.
func (l MPSNNLanczosScaleNode) Init() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLanczosScaleNode) Autorelease() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLanczosScaleNode creates a new MPSNNLanczosScaleNode instance.
func NewMPSNNLanczosScaleNode() MPSNNLanczosScaleNode {
	class := getMPSNNLanczosScaleNodeClass()
	rv := objc.Send[MPSNNLanczosScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewLanczosScaleNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNLanczosScaleNode {
	instance := getMPSNNLanczosScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNLanczosScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:outputSize:
func NewLanczosScaleNodeWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) MPSNNLanczosScaleNode {
	instance := getMPSNNLanczosScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), source, size)
	return MPSNNLanczosScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLanczosScaleNode/initWithSource:transformProvider:outputSize:
func NewLanczosScaleNodeWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) MPSNNLanczosScaleNode {
	instance := getMPSNNLanczosScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), source, provider, size)
	return MPSNNLanczosScaleNodeFromID(rv)
}
































