// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionRowMinNode] class.
var (
	_MPSNNReductionRowMinNodeClass     MPSNNReductionRowMinNodeClass
	_MPSNNReductionRowMinNodeClassOnce sync.Once
)

func getMPSNNReductionRowMinNodeClass() MPSNNReductionRowMinNodeClass {
	_MPSNNReductionRowMinNodeClassOnce.Do(func() {
		_MPSNNReductionRowMinNodeClass = MPSNNReductionRowMinNodeClass{class: objc.GetClass("MPSNNReductionRowMinNode")}
	})
	return _MPSNNReductionRowMinNodeClass
}

// GetMPSNNReductionRowMinNodeClass returns the class object for MPSNNReductionRowMinNode.
func GetMPSNNReductionRowMinNodeClass() MPSNNReductionRowMinNodeClass {
	return getMPSNNReductionRowMinNodeClass()
}

type MPSNNReductionRowMinNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionRowMinNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowMinNodeClass) Alloc() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMinNode
type MPSNNReductionRowMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowMinNodeFromID constructs a [MPSNNReductionRowMinNode] from an objc.ID.
func MPSNNReductionRowMinNodeFromID(id objc.ID) MPSNNReductionRowMinNode {
	return MPSNNReductionRowMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionRowMinNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionRowMinNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMinNode
type IMPSNNReductionRowMinNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionRowMinNode) Init() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowMinNode) Autorelease() MPSNNReductionRowMinNode {
	rv := objc.Send[MPSNNReductionRowMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowMinNode creates a new MPSNNReductionRowMinNode instance.
func NewMPSNNReductionRowMinNode() MPSNNReductionRowMinNode {
	class := getMPSNNReductionRowMinNodeClass()
	rv := objc.Send[MPSNNReductionRowMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionRowMinNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionRowMinNode {
	instance := getMPSNNReductionRowMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionRowMinNodeFromID(rv)
}
