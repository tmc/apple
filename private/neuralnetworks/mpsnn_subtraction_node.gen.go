// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNSubtractionNode] class.
var (
	_MPSNNSubtractionNodeClass     MPSNNSubtractionNodeClass
	_MPSNNSubtractionNodeClassOnce sync.Once
)

func getMPSNNSubtractionNodeClass() MPSNNSubtractionNodeClass {
	_MPSNNSubtractionNodeClassOnce.Do(func() {
		_MPSNNSubtractionNodeClass = MPSNNSubtractionNodeClass{class: objc.GetClass("MPSNNSubtractionNode")}
	})
	return _MPSNNSubtractionNodeClass
}

// GetMPSNNSubtractionNodeClass returns the class object for MPSNNSubtractionNode.
func GetMPSNNSubtractionNodeClass() MPSNNSubtractionNodeClass {
	return getMPSNNSubtractionNodeClass()
}

type MPSNNSubtractionNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSubtractionNodeClass) Alloc() MPSNNSubtractionNode {
	rv := objc.Send[MPSNNSubtractionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSubtractionNode
type MPSNNSubtractionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNSubtractionNodeFromID constructs a [MPSNNSubtractionNode] from an objc.ID.
func MPSNNSubtractionNodeFromID(id objc.ID) MPSNNSubtractionNode {
	return MPSNNSubtractionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}
// Ensure MPSNNSubtractionNode implements IMPSNNSubtractionNode.
var _ IMPSNNSubtractionNode = MPSNNSubtractionNode{}





// An interface definition for the [MPSNNSubtractionNode] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSubtractionNode
type IMPSNNSubtractionNode interface {
	IMPSNNBinaryArithmeticNode
}





// Init initializes the instance.
func (s MPSNNSubtractionNode) Init() MPSNNSubtractionNode {
	rv := objc.Send[MPSNNSubtractionNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNSubtractionNode) Autorelease() MPSNNSubtractionNode {
	rv := objc.Send[MPSNNSubtractionNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNSubtractionNode creates a new MPSNNSubtractionNode instance.
func NewMPSNNSubtractionNode() MPSNNSubtractionNode {
	class := getMPSNNSubtractionNodeClass()
	rv := objc.Send[MPSNNSubtractionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewSubtractionNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNSubtractionNode {
	instance := getMPSNNSubtractionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNSubtractionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewSubtractionNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNSubtractionNode {
	instance := getMPSNNSubtractionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNSubtractionNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewSubtractionNodeWithSources(sources objectivec.IObject) MPSNNSubtractionNode {
	instance := getMPSNNSubtractionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNSubtractionNodeFromID(rv)
}
































