// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionFeatureChannelsSumNode] class.
var (
	_MPSNNReductionFeatureChannelsSumNodeClass     MPSNNReductionFeatureChannelsSumNodeClass
	_MPSNNReductionFeatureChannelsSumNodeClassOnce sync.Once
)

func getMPSNNReductionFeatureChannelsSumNodeClass() MPSNNReductionFeatureChannelsSumNodeClass {
	_MPSNNReductionFeatureChannelsSumNodeClassOnce.Do(func() {
		_MPSNNReductionFeatureChannelsSumNodeClass = MPSNNReductionFeatureChannelsSumNodeClass{class: objc.GetClass("MPSNNReductionFeatureChannelsSumNode")}
	})
	return _MPSNNReductionFeatureChannelsSumNodeClass
}

// GetMPSNNReductionFeatureChannelsSumNodeClass returns the class object for MPSNNReductionFeatureChannelsSumNode.
func GetMPSNNReductionFeatureChannelsSumNodeClass() MPSNNReductionFeatureChannelsSumNodeClass {
	return getMPSNNReductionFeatureChannelsSumNodeClass()
}

type MPSNNReductionFeatureChannelsSumNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionFeatureChannelsSumNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionFeatureChannelsSumNodeClass) Alloc() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNNReductionFeatureChannelsSumNode.Weight]
//   - [MPSNNReductionFeatureChannelsSumNode.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode
type MPSNNReductionFeatureChannelsSumNode struct {
	MPSNNUnaryReductionNode
}

// MPSNNReductionFeatureChannelsSumNodeFromID constructs a [MPSNNReductionFeatureChannelsSumNode] from an objc.ID.
func MPSNNReductionFeatureChannelsSumNodeFromID(id objc.ID) MPSNNReductionFeatureChannelsSumNode {
	return MPSNNReductionFeatureChannelsSumNode{MPSNNUnaryReductionNode: MPSNNUnaryReductionNodeFromID(id)}
}

// NOTE: MPSNNReductionFeatureChannelsSumNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionFeatureChannelsSumNode] class.
//
// # Instance Properties
//
//   - [IMPSNNReductionFeatureChannelsSumNode.Weight]
//   - [IMPSNNReductionFeatureChannelsSumNode.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode
type IMPSNNReductionFeatureChannelsSumNode interface {
	IMPSNNUnaryReductionNode

	// Topic: Instance Properties

	Weight() float32
	SetWeight(value float32)
}

// Init initializes the instance.
func (r MPSNNReductionFeatureChannelsSumNode) Init() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionFeatureChannelsSumNode) Autorelease() MPSNNReductionFeatureChannelsSumNode {
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionFeatureChannelsSumNode creates a new MPSNNReductionFeatureChannelsSumNode instance.
func NewMPSNNReductionFeatureChannelsSumNode() MPSNNReductionFeatureChannelsSumNode {
	class := getMPSNNReductionFeatureChannelsSumNodeClass()
	rv := objc.Send[MPSNNReductionFeatureChannelsSumNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewReductionFeatureChannelsSumNodeWithSource(sourceNode IMPSNNImageNode) MPSNNReductionFeatureChannelsSumNode {
	instance := getMPSNNReductionFeatureChannelsSumNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNReductionFeatureChannelsSumNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionFeatureChannelsSumNode/weight
func (r MPSNNReductionFeatureChannelsSumNode) Weight() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("weight"))
	return rv
}
func (r MPSNNReductionFeatureChannelsSumNode) SetWeight(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setWeight:"), value)
}
