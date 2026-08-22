// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNAdditionNode] class.
var (
	_MPSNNAdditionNodeClass     MPSNNAdditionNodeClass
	_MPSNNAdditionNodeClassOnce sync.Once
)

func getMPSNNAdditionNodeClass() MPSNNAdditionNodeClass {
	_MPSNNAdditionNodeClassOnce.Do(func() {
		_MPSNNAdditionNodeClass = MPSNNAdditionNodeClass{class: objc.GetClass("MPSNNAdditionNode")}
	})
	return _MPSNNAdditionNodeClass
}

// GetMPSNNAdditionNodeClass returns the class object for MPSNNAdditionNode.
func GetMPSNNAdditionNodeClass() MPSNNAdditionNodeClass {
	return getMPSNNAdditionNodeClass()
}

type MPSNNAdditionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNAdditionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNAdditionNodeClass) Alloc() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an addition operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionNode
type MPSNNAdditionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNAdditionNodeFromID constructs a [MPSNNAdditionNode] from an objc.ID.
//
// A representation of an addition operator.
func MPSNNAdditionNodeFromID(id objc.ID) MPSNNAdditionNode {
	return MPSNNAdditionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}

// NOTE: MPSNNAdditionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNAdditionNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionNode
type IMPSNNAdditionNode interface {
	IMPSNNBinaryArithmeticNode
}

// Init initializes the instance.
func (a MPSNNAdditionNode) Init() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNAdditionNode) Autorelease() MPSNNAdditionNode {
	rv := objc.Send[MPSNNAdditionNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNAdditionNode creates a new MPSNNAdditionNode instance.
func NewMPSNNAdditionNode() MPSNNAdditionNode {
	class := getMPSNNAdditionNodeClass()
	rv := objc.Send[MPSNNAdditionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func NewAdditionNodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNAdditionNode {
	instance := getMPSNNAdditionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return MPSNNAdditionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func NewAdditionNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNAdditionNode {
	instance := getMPSNNAdditionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNAdditionNodeFromID(rv)
}
