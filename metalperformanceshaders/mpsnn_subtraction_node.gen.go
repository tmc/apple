// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNSubtractionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSubtractionNodeClass) Alloc() MPSNNSubtractionNode {
	rv := objc.Send[MPSNNSubtractionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an subtraction operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionNode
type MPSNNSubtractionNode struct {
	MPSNNBinaryArithmeticNode
}

// MPSNNSubtractionNodeFromID constructs a [MPSNNSubtractionNode] from an objc.ID.
//
// A representation of an subtraction operator.
func MPSNNSubtractionNodeFromID(id objc.ID) MPSNNSubtractionNode {
	return MPSNNSubtractionNode{MPSNNBinaryArithmeticNode: MPSNNBinaryArithmeticNodeFromID(id)}
}

// NOTE: MPSNNSubtractionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNSubtractionNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionNode
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func NewSubtractionNodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNSubtractionNode {
	instance := getMPSNNSubtractionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return MPSNNSubtractionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func NewSubtractionNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNSubtractionNode {
	instance := getMPSNNSubtractionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNSubtractionNodeFromID(rv)
}
