// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNComparisonNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNComparisonNodeClass) Alloc() MPSNNComparisonNode {
	rv := objc.Send[MPSNNComparisonNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNNComparisonNode.ComparisonType]
//   - [MPSNNComparisonNode.SetComparisonType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode
type MPSNNComparisonNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNComparisonNodeFromID constructs a [MPSNNComparisonNode] from an objc.ID.
func MPSNNComparisonNodeFromID(id objc.ID) MPSNNComparisonNode {
	return MPSNNComparisonNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}

// NOTE: MPSNNComparisonNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNComparisonNode] class.
//
// # Instance Properties
//
//   - [IMPSNNComparisonNode.ComparisonType]
//   - [IMPSNNComparisonNode.SetComparisonType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode
type IMPSNNComparisonNode interface {
	IMPSNNBinaryArithmeticNode

	// Topic: Instance Properties

	ComparisonType() MPSNNComparisonType
	SetComparisonType(value MPSNNComparisonType)
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func NewComparisonNodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNComparisonNode {
	instance := getMPSNNComparisonNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return MPSNNComparisonNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func NewComparisonNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNComparisonNode {
	instance := getMPSNNComparisonNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNComparisonNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode/comparisonType
func (c MPSNNComparisonNode) ComparisonType() MPSNNComparisonType {
	rv := objc.Send[MPSNNComparisonType](c.ID, objc.Sel("comparisonType"))
	return MPSNNComparisonType(rv)
}
func (c MPSNNComparisonNode) SetComparisonType(value MPSNNComparisonType) {
	objc.Send[struct{}](c.ID, objc.Sel("setComparisonType:"), value)
}
