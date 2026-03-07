// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNScaleNode] class.
var (
	_MPSNNScaleNodeClass     MPSNNScaleNodeClass
	_MPSNNScaleNodeClassOnce sync.Once
)

func getMPSNNScaleNodeClass() MPSNNScaleNodeClass {
	_MPSNNScaleNodeClassOnce.Do(func() {
		_MPSNNScaleNodeClass = MPSNNScaleNodeClass{class: objc.GetClass("MPSNNScaleNode")}
	})
	return _MPSNNScaleNodeClass
}

// GetMPSNNScaleNodeClass returns the class object for MPSNNScaleNode.
func GetMPSNNScaleNodeClass() MPSNNScaleNodeClass {
	return getMPSNNScaleNodeClass()
}

type MPSNNScaleNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNScaleNodeClass) Alloc() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNScaleNode.PrivateInitWithSourceTransformProviderOutputSize]
//   - [MPSNNScaleNode.InitWithSourceOutputSize]
//   - [MPSNNScaleNode.InitWithSourceTransformProviderOutputSize]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode
type MPSNNScaleNode struct {
	MPSNNFilterNode
}

// MPSNNScaleNodeFromID constructs a [MPSNNScaleNode] from an objc.ID.
func MPSNNScaleNodeFromID(id objc.ID) MPSNNScaleNode {
	return MPSNNScaleNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNScaleNode implements IMPSNNScaleNode.
var _ IMPSNNScaleNode = MPSNNScaleNode{}





// An interface definition for the [MPSNNScaleNode] class.
//
// # Methods
//
//   - [IMPSNNScaleNode.PrivateInitWithSourceTransformProviderOutputSize]
//   - [IMPSNNScaleNode.InitWithSourceOutputSize]
//   - [IMPSNNScaleNode.InitWithSourceTransformProviderOutputSize]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode
type IMPSNNScaleNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	PrivateInitWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) objectivec.IObject
	InitWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) MPSNNScaleNode
	InitWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) MPSNNScaleNode
}





// Init initializes the instance.
func (s MPSNNScaleNode) Init() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNScaleNode) Autorelease() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNScaleNode creates a new MPSNNScaleNode instance.
func NewMPSNNScaleNode() MPSNNScaleNode {
	class := getMPSNNScaleNodeClass()
	rv := objc.Send[MPSNNScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewScaleNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNScaleNode {
	instance := getMPSNNScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:outputSize:
func NewScaleNodeWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) MPSNNScaleNode {
	instance := getMPSNNScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), source, size)
	return MPSNNScaleNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:transformProvider:outputSize:
func NewScaleNodeWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) MPSNNScaleNode {
	instance := getMPSNNScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), source, provider, size)
	return MPSNNScaleNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/privateInitWithSource:transformProvider:outputSize:
func (s MPSNNScaleNode) PrivateInitWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("privateInitWithSource:transformProvider:outputSize:"), source, provider, size)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:outputSize:
func (s MPSNNScaleNode) InitWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("initWithSource:outputSize:"), source, size)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/initWithSource:transformProvider:outputSize:
func (s MPSNNScaleNode) InitWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), source, provider, size)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/nodeWithSource:outputSize:
func (_MPSNNScaleNodeClass MPSNNScaleNodeClass) NodeWithSourceOutputSize(source objectivec.IObject, size objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNScaleNodeClass.class), objc.Sel("nodeWithSource:outputSize:"), source, size)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScaleNode/nodeWithSource:transformProvider:outputSize:
func (_MPSNNScaleNodeClass MPSNNScaleNodeClass) NodeWithSourceTransformProviderOutputSize(source objectivec.IObject, provider objectivec.IObject, size objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNScaleNodeClass.class), objc.Sel("nodeWithSource:transformProvider:outputSize:"), source, provider, size)
	return objectivec.Object{ID: rv}
}






















