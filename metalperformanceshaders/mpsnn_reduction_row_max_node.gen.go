// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionRowMaxNode] class.
var (
	_MPSNNReductionRowMaxNodeClass     MPSNNReductionRowMaxNodeClass
	_MPSNNReductionRowMaxNodeClassOnce sync.Once
)

func getMPSNNReductionRowMaxNodeClass() MPSNNReductionRowMaxNodeClass {
	_MPSNNReductionRowMaxNodeClassOnce.Do(func() {
		_MPSNNReductionRowMaxNodeClass = MPSNNReductionRowMaxNodeClass{class: objc.GetClass("MPSNNReductionRowMaxNode")}
	})
	return _MPSNNReductionRowMaxNodeClass
}

// GetMPSNNReductionRowMaxNodeClass returns the class object for MPSNNReductionRowMaxNode.
func GetMPSNNReductionRowMaxNodeClass() MPSNNReductionRowMaxNodeClass {
	return getMPSNNReductionRowMaxNodeClass()
}

type MPSNNReductionRowMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionRowMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionRowMaxNodeClass) Alloc() MPSNNReductionRowMaxNode {
	rv := objc.Send[MPSNNReductionRowMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMaxNode
type MPSNNReductionRowMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionRowMaxNodeFromID constructs a [MPSNNReductionRowMaxNode] from an objc.ID.
func MPSNNReductionRowMaxNodeFromID(id objc.ID) MPSNNReductionRowMaxNode {
	return MPSNNReductionRowMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionRowMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionRowMaxNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionRowMaxNode
type IMPSNNReductionRowMaxNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionRowMaxNode) Init() MPSNNReductionRowMaxNode {
	rv := objc.Send[MPSNNReductionRowMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionRowMaxNode) Autorelease() MPSNNReductionRowMaxNode {
	rv := objc.Send[MPSNNReductionRowMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionRowMaxNode creates a new MPSNNReductionRowMaxNode instance.
func NewMPSNNReductionRowMaxNode() MPSNNReductionRowMaxNode {
	class := getMPSNNReductionRowMaxNodeClass()
	rv := objc.Send[MPSNNReductionRowMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionRowMaxNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionRowMaxNode {
	instance := getMPSNNReductionRowMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionRowMaxNodeFromID(rv)
}
