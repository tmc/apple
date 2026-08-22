// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionFeatureChannelsArgumentMinNode] class.
var (
	_MPSNNReductionFeatureChannelsArgumentMinNodeClass     MPSNNReductionFeatureChannelsArgumentMinNodeClass
	_MPSNNReductionFeatureChannelsArgumentMinNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsArgumentMinNodeClass() MPSNNReductionFeatureChannelsArgumentMinNodeClass {
	_MPSNNReductionFeatureChannelsArgumentMinNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsArgumentMinNodeClass = MPSNNReductionFeatureChannelsArgumentMinNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsArgumentMinNode")}
	})
	return _MPSNNReductionFeatureChannelsArgumentMinNodeClass
}

// GetMPSNNReductionFeatureChannelsArgumentMinNodeClass returns the class object for MPSNNReductionFeatureChannelsArgumentMinNode.
func GetMPSNNReductionFeatureChannelsArgumentMinNodeClass() MPSNNReductionFeatureChannelsArgumentMinNodeClass {
	return getMPSNNReductionFeatureChannelsArgumentMinNodeClass()
}

type MPSNNReductionFeatureChannelsArgumentMinNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionFeatureChannelsArgumentMinNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsArgumentMinNodeClass) Alloc() MPSNNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMinNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMinNode
type MPSNNReductionFeatureChannelsArgumentMinNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsArgumentMinNodeFromID constructs a [MPSNNReductionFeatureChannelsArgumentMinNode] from an objc.ID.
func MPSNNReductionFeatureChannelsArgumentMinNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsArgumentMinNode {
	return MPSNNReductionFeatureChannelsArgumentMinNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionFeatureChannelsArgumentMinNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionFeatureChannelsArgumentMinNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsArgumentMinNode
type IMPSNNReductionFeatureChannelsArgumentMinNode interface {
	IMPSNNUnaryReductionNode
}

// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsArgumentMinNode) Init() MPSNNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMinNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsArgumentMinNode) Autorelease() MPSNNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMinNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsArgumentMinNode creates a new MPSNNReductionFeatureChannelsArgumentMinNode instance.
func NewMPSNNReductionFeatureChannelsArgumentMinNode() MPSNNReductionFeatureChannelsArgumentMinNode {
	class := getMPSNNReductionFeatureChannelsArgumentMinNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsArgumentMinNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionFeatureChannelsArgumentMinNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionFeatureChannelsArgumentMinNode {
	instance := getMPSNNReductionFeatureChannelsArgumentMinNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionFeatureChannelsArgumentMinNodeFromID(rv)
}
