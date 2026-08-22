// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionColumnMeanNode] class.
var (
	_MPSNNReductionColumnMeanNodeClass     MPSNNReductionColumnMeanNodeClass
	_MPSNNReductionColumnMeanNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMeanNodeClass() MPSNNReductionColumnMeanNodeClass {
	_MPSNNReductionColumnMeanNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMeanNodeClass = MPSNNReductionColumnMeanNodeClass{class: objc.GetClass("MPSNNReductionColumnMeanNode")}
	})
	return _MPSNNReductionColumnMeanNodeClass
}

// GetMPSNNReductionColumnMeanNodeClass returns the class object for MPSNNReductionColumnMeanNode.
func GetMPSNNReductionColumnMeanNodeClass() MPSNNReductionColumnMeanNodeClass {
	return getMPSNNReductionColumnMeanNodeClass()
}

type MPSNNReductionColumnMeanNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionColumnMeanNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMeanNodeClass) Alloc() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMeanNode
type MPSNNReductionColumnMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMeanNodeFromID constructs a [MPSNNReductionColumnMeanNode] from an objc.ID.
func MPSNNReductionColumnMeanNodeFromID(id objc.ID) MPSNNReductionColumnMeanNode {
	return MPSNNReductionColumnMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionColumnMeanNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionColumnMeanNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMeanNode
type IMPSNNReductionColumnMeanNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionColumnMeanNode) Init() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMeanNode) Autorelease() MPSNNReductionColumnMeanNode {
	rv := objc.Send[MPSNNReductionColumnMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMeanNode creates a new MPSNNReductionColumnMeanNode instance.
func NewMPSNNReductionColumnMeanNode() MPSNNReductionColumnMeanNode {
	class := getMPSNNReductionColumnMeanNodeClass()
	rv := objc.Send[MPSNNReductionColumnMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionColumnMeanNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionColumnMeanNode {
	instance := getMPSNNReductionColumnMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionColumnMeanNodeFromID(rv)
}
