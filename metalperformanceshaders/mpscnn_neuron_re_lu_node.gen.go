// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronReLUNode] class.
var (
	_MPSCNNNeuronReLUNodeClass     MPSCNNNeuronReLUNodeClass
	_MPSCNNNeuronReLUNodeClassOnce sync.Once
)

func getMPSCNNNeuronReLUNodeClass() MPSCNNNeuronReLUNodeClass {
	_MPSCNNNeuronReLUNodeClassOnce.Do(func() {
		_MPSCNNNeuronReLUNodeClass = MPSCNNNeuronReLUNodeClass{class: objc.GetClass("MPSCNNNeuronReLUNode")}
	})
	return _MPSCNNNeuronReLUNodeClass
}

// GetMPSCNNNeuronReLUNodeClass returns the class object for MPSCNNNeuronReLUNode.
func GetMPSCNNNeuronReLUNodeClass() MPSCNNNeuronReLUNodeClass {
	return getMPSCNNNeuronReLUNodeClass()
}

type MPSCNNNeuronReLUNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronReLUNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronReLUNodeClass) Alloc() MPSCNNNeuronReLUNode {
	rv := objc.Send[MPSCNNNeuronReLUNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation a ReLU neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronReLUNode.InitWithSource]
//   - [MPSCNNNeuronReLUNode.InitWithSourceA]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode
type MPSCNNNeuronReLUNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronReLUNodeFromID constructs a [MPSCNNNeuronReLUNode] from an objc.ID.
//
// A representation a ReLU neuron filter.
func MPSCNNNeuronReLUNodeFromID(id objc.ID) MPSCNNNeuronReLUNode {
	return MPSCNNNeuronReLUNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronReLUNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronReLUNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronReLUNode.InitWithSource]
//   - [IMPSCNNNeuronReLUNode.InitWithSourceA]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode
type IMPSCNNNeuronReLUNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNode
	InitWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronReLUNode
}

// Init initializes the instance.
func (c MPSCNNNeuronReLUNode) Init() MPSCNNNeuronReLUNode {
	rv := objc.Send[MPSCNNNeuronReLUNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronReLUNode) Autorelease() MPSCNNNeuronReLUNode {
	rv := objc.Send[MPSCNNNeuronReLUNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronReLUNode creates a new MPSCNNNeuronReLUNode instance.
func NewMPSCNNNeuronReLUNode() MPSCNNNeuronReLUNode {
	class := getMPSCNNNeuronReLUNodeClass()
	rv := objc.Send[MPSCNNNeuronReLUNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/init(source:)
func NewCNNNeuronReLUNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNode {
	instance := getMPSCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/init(source:a:)
func NewCNNNeuronReLUNodeWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronReLUNode {
	instance := getMPSCNNNeuronReLUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	return MPSCNNNeuronReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronReLUNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronReLUNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronReLUNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/init(source:)
func (c MPSCNNNeuronReLUNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNode {
	rv := objc.Send[MPSCNNNeuronReLUNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/init(source:a:)
func (c MPSCNNNeuronReLUNode) InitWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronReLUNode {
	rv := objc.Send[MPSCNNNeuronReLUNode](c.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/nodeWithSource:
func (_MPSCNNNeuronReLUNodeClass MPSCNNNeuronReLUNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronReLUNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNode/nodeWithSource:a:
func (_MPSCNNNeuronReLUNodeClass MPSCNNNeuronReLUNodeClass) NodeWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronReLUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronReLUNodeClass.class), objc.Sel("nodeWithSource:a:"), sourceNode, a)
	return MPSCNNNeuronReLUNodeFromID(rv)
}
