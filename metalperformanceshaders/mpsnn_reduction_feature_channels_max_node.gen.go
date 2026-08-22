// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionFeatureChannelsMaxNode] class.
var (
	_MPSNNReductionFeatureChannelsMaxNodeClass     MPSNNReductionFeatureChannelsMaxNodeClass
	_MPSNNReductionFeatureChannelsMaxNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsMaxNodeClass() MPSNNReductionFeatureChannelsMaxNodeClass {
	_MPSNNReductionFeatureChannelsMaxNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsMaxNodeClass = MPSNNReductionFeatureChannelsMaxNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsMaxNode")}
	})
	return _MPSNNReductionFeatureChannelsMaxNodeClass
}

// GetMPSNNReductionFeatureChannelsMaxNodeClass returns the class object for MPSNNReductionFeatureChannelsMaxNode.
func GetMPSNNReductionFeatureChannelsMaxNodeClass() MPSNNReductionFeatureChannelsMaxNodeClass {
	return getMPSNNReductionFeatureChannelsMaxNodeClass()
}

type MPSNNReductionFeatureChannelsMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionFeatureChannelsMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsMaxNodeClass) Alloc() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMaxNode
type MPSNNReductionFeatureChannelsMaxNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsMaxNodeFromID constructs a [MPSNNReductionFeatureChannelsMaxNode] from an objc.ID.
func MPSNNReductionFeatureChannelsMaxNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsMaxNode {
	return MPSNNReductionFeatureChannelsMaxNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionFeatureChannelsMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionFeatureChannelsMaxNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsMaxNode
type IMPSNNReductionFeatureChannelsMaxNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsMaxNode) Init() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsMaxNode) Autorelease() MPSNNReductionFeatureChannelsMaxNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsMaxNode creates a new MPSNNReductionFeatureChannelsMaxNode instance.
func NewMPSNNReductionFeatureChannelsMaxNode() MPSNNReductionFeatureChannelsMaxNode {
	class := getMPSNNReductionFeatureChannelsMaxNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionFeatureChannelsMaxNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionFeatureChannelsMaxNode {
	instance := getMPSNNReductionFeatureChannelsMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionFeatureChannelsMaxNodeFromID(rv)
}
