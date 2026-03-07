// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReshapeNode] class.
var (
	_MPSNNReshapeNodeClass     MPSNNReshapeNodeClass
	_MPSNNReshapeNodeClassOnce sync.Once
)

func getMPSNNReshapeNodeClass() MPSNNReshapeNodeClass {
	_MPSNNReshapeNodeClassOnce.Do(func() {
		_MPSNNReshapeNodeClass = MPSNNReshapeNodeClass{class: objc.GetClass("MPSNNReshapeNode")}
	})
	return _MPSNNReshapeNodeClass
}

// GetMPSNNReshapeNodeClass returns the class object for MPSNNReshapeNode.
func GetMPSNNReshapeNodeClass() MPSNNReshapeNodeClass {
	return getMPSNNReshapeNodeClass()
}

type MPSNNReshapeNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeNodeClass) Alloc() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReshapeNode.InitWithSourceResultWidthResultHeightResultFeatureChannels]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeNode
type MPSNNReshapeNode struct {
	MPSNNFilterNode
}

// MPSNNReshapeNodeFromID constructs a [MPSNNReshapeNode] from an objc.ID.
func MPSNNReshapeNodeFromID(id objc.ID) MPSNNReshapeNode {
	return MPSNNReshapeNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNReshapeNode implements IMPSNNReshapeNode.
var _ IMPSNNReshapeNode = MPSNNReshapeNode{}





// An interface definition for the [MPSNNReshapeNode] class.
//
// # Methods
//
//   - [IMPSNNReshapeNode.InitWithSourceResultWidthResultHeightResultFeatureChannels]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeNode
type IMPSNNReshapeNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	InitWithSourceResultWidthResultHeightResultFeatureChannels(source objectivec.IObject, width uint64, height uint64, channels uint64) MPSNNReshapeNode
}





// Init initializes the instance.
func (r MPSNNReshapeNode) Init() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeNode) Autorelease() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeNode creates a new MPSNNReshapeNode instance.
func NewMPSNNReshapeNode() MPSNNReshapeNode {
	class := getMPSNNReshapeNodeClass()
	rv := objc.Send[MPSNNReshapeNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewReshapeNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNReshapeNode {
	instance := getMPSNNReshapeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNReshapeNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeNode/initWithSource:resultWidth:resultHeight:resultFeatureChannels:
func NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source objectivec.IObject, width uint64, height uint64, channels uint64) MPSNNReshapeNode {
	instance := getMPSNNReshapeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, width, height, channels)
	return MPSNNReshapeNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeNode/initWithSource:resultWidth:resultHeight:resultFeatureChannels:
func (r MPSNNReshapeNode) InitWithSourceResultWidthResultHeightResultFeatureChannels(source objectivec.IObject, width uint64, height uint64, channels uint64) MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, width, height, channels)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshapeNode/nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:
func (_MPSNNReshapeNodeClass MPSNNReshapeNodeClass) NodeWithSourceResultWidthResultHeightResultFeatureChannels(source objectivec.IObject, width uint64, height uint64, channels uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReshapeNodeClass.class), objc.Sel("nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, width, height, channels)
	return objectivec.Object{ID: rv}
}






















