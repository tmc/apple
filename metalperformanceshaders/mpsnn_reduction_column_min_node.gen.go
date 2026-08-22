// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionColumnMinNode] class.
var (
	_MPSNNReductionColumnMinNodeClass     MPSNNReductionColumnMinNodeClass
	_MPSNNReductionColumnMinNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMinNodeClass() MPSNNReductionColumnMinNodeClass {
	_MPSNNReductionColumnMinNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMinNodeClass = MPSNNReductionColumnMinNodeClass{class: objc.GetClass("MPSNNReductionColumnMinNode")}
	})
	return _MPSNNReductionColumnMinNodeClass
}

// GetMPSNNReductionColumnMinNodeClass returns the class object for MPSNNReductionColumnMinNode.
func GetMPSNNReductionColumnMinNodeClass() MPSNNReductionColumnMinNodeClass {
	return getMPSNNReductionColumnMinNodeClass()
}

type MPSNNReductionColumnMinNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionColumnMinNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMinNodeClass) Alloc() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMinNode
type MPSNNReductionColumnMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMinNodeFromID constructs a [MPSNNReductionColumnMinNode] from an objc.ID.
func MPSNNReductionColumnMinNodeFromID(id objc.ID) MPSNNReductionColumnMinNode {
	return MPSNNReductionColumnMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionColumnMinNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionColumnMinNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMinNode
type IMPSNNReductionColumnMinNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionColumnMinNode) Init() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMinNode) Autorelease() MPSNNReductionColumnMinNode {
	rv := objc.Send[MPSNNReductionColumnMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMinNode creates a new MPSNNReductionColumnMinNode instance.
func NewMPSNNReductionColumnMinNode() MPSNNReductionColumnMinNode {
	class := getMPSNNReductionColumnMinNodeClass()
	rv := objc.Send[MPSNNReductionColumnMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionColumnMinNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionColumnMinNode {
	instance := getMPSNNReductionColumnMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionColumnMinNodeFromID(rv)
}
