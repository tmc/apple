// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionRowMeanNode] class.
var (
	_MPSNNReductionRowMeanNodeClass     MPSNNReductionRowMeanNodeClass
	_MPSNNReductionRowMeanNodeClassOnce sync.Once
)

func getMPSNNReductionRowMeanNodeClass() MPSNNReductionRowMeanNodeClass {
	_MPSNNReductionRowMeanNodeClassOnce.Do(func() {
		_MPSNNReductionRowMeanNodeClass = MPSNNReductionRowMeanNodeClass{class: objc.GetClass("MPSNNReductionRowMeanNode")}
	})
	return _MPSNNReductionRowMeanNodeClass
}

// GetMPSNNReductionRowMeanNodeClass returns the class object for MPSNNReductionRowMeanNode.
func GetMPSNNReductionRowMeanNodeClass() MPSNNReductionRowMeanNodeClass {
	return getMPSNNReductionRowMeanNodeClass()
}

type MPSNNReductionRowMeanNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionRowMeanNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowMeanNodeClass) Alloc() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMeanNode
type MPSNNReductionRowMeanNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowMeanNodeFromID constructs a [MPSNNReductionRowMeanNode] from an objc.ID.
func MPSNNReductionRowMeanNodeFromID(id objc.ID) MPSNNReductionRowMeanNode {
	return MPSNNReductionRowMeanNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionRowMeanNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionRowMeanNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMeanNode
type IMPSNNReductionRowMeanNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionRowMeanNode) Init() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowMeanNode) Autorelease() MPSNNReductionRowMeanNode {
	rv := objc.Send[MPSNNReductionRowMeanNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowMeanNode creates a new MPSNNReductionRowMeanNode instance.
func NewMPSNNReductionRowMeanNode() MPSNNReductionRowMeanNode {
	class := getMPSNNReductionRowMeanNodeClass()
	rv := objc.Send[MPSNNReductionRowMeanNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionRowMeanNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionRowMeanNode {
	instance := getMPSNNReductionRowMeanNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionRowMeanNodeFromID(rv)
}
