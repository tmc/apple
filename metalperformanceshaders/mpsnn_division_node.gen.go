// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNDivisionNode] class.
var (
	_MPSNNDivisionNodeClass     MPSNNDivisionNodeClass
	_MPSNNDivisionNodeClassOnce sync.Once
)

func getMPSNNDivisionNodeClass() MPSNNDivisionNodeClass {
	_MPSNNDivisionNodeClassOnce.Do(func() {
		_MPSNNDivisionNodeClass = MPSNNDivisionNodeClass{class: objc.GetClass("MPSNNDivisionNode")}
	})
	return _MPSNNDivisionNodeClass
}

// GetMPSNNDivisionNodeClass returns the class object for MPSNNDivisionNode.
func GetMPSNNDivisionNodeClass() MPSNNDivisionNodeClass {
	return getMPSNNDivisionNodeClass()
}

type MPSNNDivisionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNDivisionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNDivisionNodeClass) Alloc() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a division operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDivisionNode
type MPSNNDivisionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNDivisionNodeFromID constructs a [MPSNNDivisionNode] from an objc.ID.
//
// A representation of a division operator.
func MPSNNDivisionNodeFromID(id objc.ID) MPSNNDivisionNode {
	return MPSNNDivisionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}

// NOTE: MPSNNDivisionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNDivisionNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDivisionNode
type IMPSNNDivisionNode interface {
	IMPSNNBinaryArithmeticNode
}

// Init initializes the instance.
func (d MPSNNDivisionNode) Init() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MPSNNDivisionNode) Autorelease() MPSNNDivisionNode {
	rv := objc.Send[MPSNNDivisionNode](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNDivisionNode creates a new MPSNNDivisionNode instance.
func NewMPSNNDivisionNode() MPSNNDivisionNode {
	class := getMPSNNDivisionNodeClass()
	rv := objc.Send[MPSNNDivisionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func NewDivisionNodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNDivisionNode {
	instance := getMPSNNDivisionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return MPSNNDivisionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func NewDivisionNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNDivisionNode {
	instance := getMPSNNDivisionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNDivisionNodeFromID(rv)
}
