// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronExponentialNode] class.
var (
	_MPSCNNNeuronExponentialNodeClass     MPSCNNNeuronExponentialNodeClass
	_MPSCNNNeuronExponentialNodeClassOnce sync.Once
)

func getMPSCNNNeuronExponentialNodeClass() MPSCNNNeuronExponentialNodeClass {
	_MPSCNNNeuronExponentialNodeClassOnce.Do(func() {
		_MPSCNNNeuronExponentialNodeClass = MPSCNNNeuronExponentialNodeClass{class: objc.GetClass("MPSCNNNeuronExponentialNode")}
	})
	return _MPSCNNNeuronExponentialNodeClass
}

// GetMPSCNNNeuronExponentialNodeClass returns the class object for MPSCNNNeuronExponentialNode.
func GetMPSCNNNeuronExponentialNodeClass() MPSCNNNeuronExponentialNodeClass {
	return getMPSCNNNeuronExponentialNodeClass()
}

type MPSCNNNeuronExponentialNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronExponentialNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronExponentialNodeClass) Alloc() MPSCNNNeuronExponentialNode {
	rv := objc.Send[MPSCNNNeuronExponentialNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an exponential neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronExponentialNode.InitWithSource]
//   - [MPSCNNNeuronExponentialNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode
type MPSCNNNeuronExponentialNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronExponentialNodeFromID constructs a [MPSCNNNeuronExponentialNode] from an objc.ID.
//
// A representation of an exponential neuron filter.
func MPSCNNNeuronExponentialNodeFromID(id objc.ID) MPSCNNNeuronExponentialNode {
	return MPSCNNNeuronExponentialNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronExponentialNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronExponentialNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronExponentialNode.InitWithSource]
//   - [IMPSCNNNeuronExponentialNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode
type IMPSCNNNeuronExponentialNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronExponentialNode
	InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronExponentialNode
}

// Init initializes the instance.
func (c MPSCNNNeuronExponentialNode) Init() MPSCNNNeuronExponentialNode {
	rv := objc.Send[MPSCNNNeuronExponentialNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronExponentialNode) Autorelease() MPSCNNNeuronExponentialNode {
	rv := objc.Send[MPSCNNNeuronExponentialNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronExponentialNode creates a new MPSCNNNeuronExponentialNode instance.
func NewMPSCNNNeuronExponentialNode() MPSCNNNeuronExponentialNode {
	class := getMPSCNNNeuronExponentialNodeClass()
	rv := objc.Send[MPSCNNNeuronExponentialNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/init(source:)
func NewCNNNeuronExponentialNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronExponentialNode {
	instance := getMPSCNNNeuronExponentialNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronExponentialNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/init(source:a:b:c:)
func NewCNNNeuronExponentialNodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronExponentialNode {
	instance := getMPSCNNNeuronExponentialNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronExponentialNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronExponentialNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronExponentialNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronExponentialNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronExponentialNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/init(source:)
func (c MPSCNNNeuronExponentialNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronExponentialNode {
	rv := objc.Send[MPSCNNNeuronExponentialNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/init(source:a:b:c:)
func (c_ MPSCNNNeuronExponentialNode) InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronExponentialNode {
	rv := objc.Send[MPSCNNNeuronExponentialNode](c_.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/nodeWithSource:
func (_MPSCNNNeuronExponentialNodeClass MPSCNNNeuronExponentialNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronExponentialNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronExponentialNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronExponentialNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponentialNode/nodeWithSource:a:b:c:
func (_MPSCNNNeuronExponentialNodeClass MPSCNNNeuronExponentialNodeClass) NodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronExponentialNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronExponentialNodeClass.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronExponentialNodeFromID(rv)
}
