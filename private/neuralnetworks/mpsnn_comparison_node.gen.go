// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNComparisonNode] class.
var (
	_MPSNNComparisonNodeClass     MPSNNComparisonNodeClass
	_MPSNNComparisonNodeClassOnce sync.Once
)

func getMPSNNComparisonNodeClass() MPSNNComparisonNodeClass {
	_MPSNNComparisonNodeClassOnce.Do(func() {
		_MPSNNComparisonNodeClass = MPSNNComparisonNodeClass{class: objc.GetClass("MPSNNComparisonNode")}
	})
	return _MPSNNComparisonNodeClass
}

// GetMPSNNComparisonNodeClass returns the class object for MPSNNComparisonNode.
func GetMPSNNComparisonNodeClass() MPSNNComparisonNodeClass {
	return getMPSNNComparisonNodeClass()
}

type MPSNNComparisonNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNComparisonNodeClass) Alloc() MPSNNComparisonNode {
	rv := objc.Send[MPSNNComparisonNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNComparisonNode.ComparisonType]
//   - [MPSNNComparisonNode.SetComparisonType]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNComparisonNode
type MPSNNComparisonNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNComparisonNodeFromID constructs a [MPSNNComparisonNode] from an objc.ID.
func MPSNNComparisonNodeFromID(id objc.ID) MPSNNComparisonNode {
	return MPSNNComparisonNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}
// Ensure MPSNNComparisonNode implements IMPSNNComparisonNode.
var _ IMPSNNComparisonNode = MPSNNComparisonNode{}





// An interface definition for the [MPSNNComparisonNode] class.
//
// # Methods
//
//   - [IMPSNNComparisonNode.ComparisonType]
//   - [IMPSNNComparisonNode.SetComparisonType]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNComparisonNode
type IMPSNNComparisonNode interface {
	IMPSNNBinaryArithmeticNode

	// Topic: Methods

	ComparisonType() uint64
	SetComparisonType(value uint64)
}





// Init initializes the instance.
func (c MPSNNComparisonNode) Init() MPSNNComparisonNode {
	rv := objc.Send[MPSNNComparisonNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNComparisonNode) Autorelease() MPSNNComparisonNode {
	rv := objc.Send[MPSNNComparisonNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNComparisonNode creates a new MPSNNComparisonNode instance.
func NewMPSNNComparisonNode() MPSNNComparisonNode {
	class := getMPSNNComparisonNodeClass()
	rv := objc.Send[MPSNNComparisonNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewComparisonNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNComparisonNode {
	instance := getMPSNNComparisonNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNComparisonNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewComparisonNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNComparisonNode {
	instance := getMPSNNComparisonNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNComparisonNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewComparisonNodeWithSources(sources objectivec.IObject) MPSNNComparisonNode {
	instance := getMPSNNComparisonNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNComparisonNodeFromID(rv)
}


















// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNComparisonNode/comparisonType
func (c MPSNNComparisonNode) ComparisonType() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("comparisonType"))
	return rv
}
func (c MPSNNComparisonNode) SetComparisonType(value uint64) {
	objc.Send[struct{}](c.ID, objc.Sel("setComparisonType:"), value)
}

















