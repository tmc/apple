// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionColumnMaxNode] class.
var (
	_MPSNNReductionColumnMaxNodeClass     MPSNNReductionColumnMaxNodeClass
	_MPSNNReductionColumnMaxNodeClassOnce sync.Once
)

func getMPSNNReductionColumnMaxNodeClass() MPSNNReductionColumnMaxNodeClass {
	_MPSNNReductionColumnMaxNodeClassOnce.Do(func() {
		_MPSNNReductionColumnMaxNodeClass = MPSNNReductionColumnMaxNodeClass{class: objc.GetClass("MPSNNReductionColumnMaxNode")}
	})
	return _MPSNNReductionColumnMaxNodeClass
}

// GetMPSNNReductionColumnMaxNodeClass returns the class object for MPSNNReductionColumnMaxNode.
func GetMPSNNReductionColumnMaxNodeClass() MPSNNReductionColumnMaxNodeClass {
	return getMPSNNReductionColumnMaxNodeClass()
}

type MPSNNReductionColumnMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionColumnMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionColumnMaxNodeClass) Alloc() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMaxNode
type MPSNNReductionColumnMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionColumnMaxNodeFromID constructs a [MPSNNReductionColumnMaxNode] from an objc.ID.
func MPSNNReductionColumnMaxNodeFromID(id objc.ID) MPSNNReductionColumnMaxNode {
	return MPSNNReductionColumnMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionColumnMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionColumnMaxNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionColumnMaxNode
type IMPSNNReductionColumnMaxNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionColumnMaxNode) Init() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionColumnMaxNode) Autorelease() MPSNNReductionColumnMaxNode {
	rv := objc.Send[MPSNNReductionColumnMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionColumnMaxNode creates a new MPSNNReductionColumnMaxNode instance.
func NewMPSNNReductionColumnMaxNode() MPSNNReductionColumnMaxNode {
	class := getMPSNNReductionColumnMaxNodeClass()
	rv := objc.Send[MPSNNReductionColumnMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionColumnMaxNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionColumnMaxNode {
	instance := getMPSNNReductionColumnMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionColumnMaxNodeFromID(rv)
}
