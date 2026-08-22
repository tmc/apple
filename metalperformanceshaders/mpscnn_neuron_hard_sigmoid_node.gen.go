// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronHardSigmoidNode] class.
var (
	_MPSCNNNeuronHardSigmoidNodeClass     MPSCNNNeuronHardSigmoidNodeClass
	_MPSCNNNeuronHardSigmoidNodeClassOnce sync.Once
)

func getMPSCNNNeuronHardSigmoidNodeClass() MPSCNNNeuronHardSigmoidNodeClass {
	_MPSCNNNeuronHardSigmoidNodeClassOnce.Do(func() {
		_MPSCNNNeuronHardSigmoidNodeClass = MPSCNNNeuronHardSigmoidNodeClass{class: objc.GetClass("MPSCNNNeuronHardSigmoidNode")}
	})
	return _MPSCNNNeuronHardSigmoidNodeClass
}

// GetMPSCNNNeuronHardSigmoidNodeClass returns the class object for MPSCNNNeuronHardSigmoidNode.
func GetMPSCNNNeuronHardSigmoidNodeClass() MPSCNNNeuronHardSigmoidNodeClass {
	return getMPSCNNNeuronHardSigmoidNodeClass()
}

type MPSCNNNeuronHardSigmoidNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronHardSigmoidNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronHardSigmoidNodeClass) Alloc() MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a hard sigmoid neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronHardSigmoidNode.InitWithSourceAB]
//   - [MPSCNNNeuronHardSigmoidNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode
type MPSCNNNeuronHardSigmoidNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronHardSigmoidNodeFromID constructs a [MPSCNNNeuronHardSigmoidNode] from an objc.ID.
//
// A representation of a hard sigmoid neuron filter.
func MPSCNNNeuronHardSigmoidNodeFromID(id objc.ID) MPSCNNNeuronHardSigmoidNode {
	return MPSCNNNeuronHardSigmoidNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronHardSigmoidNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronHardSigmoidNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronHardSigmoidNode.InitWithSourceAB]
//   - [IMPSCNNNeuronHardSigmoidNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode
type IMPSCNNNeuronHardSigmoidNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronHardSigmoidNode
	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronHardSigmoidNode
}

// Init initializes the instance.
func (c MPSCNNNeuronHardSigmoidNode) Init() MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronHardSigmoidNode) Autorelease() MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronHardSigmoidNode creates a new MPSCNNNeuronHardSigmoidNode instance.
func NewMPSCNNNeuronHardSigmoidNode() MPSCNNNeuronHardSigmoidNode {
	class := getMPSCNNNeuronHardSigmoidNodeClass()
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/init(source:)
func NewCNNNeuronHardSigmoidNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronHardSigmoidNode {
	instance := getMPSCNNNeuronHardSigmoidNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronHardSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/init(source:a:b:)
func NewCNNNeuronHardSigmoidNodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronHardSigmoidNode {
	instance := getMPSCNNNeuronHardSigmoidNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronHardSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronHardSigmoidNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronHardSigmoidNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronHardSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/init(source:a:b:)
func (c MPSCNNNeuronHardSigmoidNode) InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](c.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/init(source:)
func (c MPSCNNNeuronHardSigmoidNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[MPSCNNNeuronHardSigmoidNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/nodeWithSource:
func (_MPSCNNNeuronHardSigmoidNodeClass MPSCNNNeuronHardSigmoidNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronHardSigmoidNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronHardSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoidNode/nodeWithSource:a:b:
func (_MPSCNNNeuronHardSigmoidNodeClass MPSCNNNeuronHardSigmoidNodeClass) NodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronHardSigmoidNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronHardSigmoidNodeClass.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronHardSigmoidNodeFromID(rv)
}
