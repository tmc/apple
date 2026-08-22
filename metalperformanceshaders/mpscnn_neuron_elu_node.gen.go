// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronELUNode] class.
var (
	_MPSCNNNeuronELUNodeClass     MPSCNNNeuronELUNodeClass
	_MPSCNNNeuronELUNodeClassOnce sync.Once
)

func getMPSCNNNeuronELUNodeClass() MPSCNNNeuronELUNodeClass {
	_MPSCNNNeuronELUNodeClassOnce.Do(func() {
		_MPSCNNNeuronELUNodeClass = MPSCNNNeuronELUNodeClass{class: objc.GetClass("MPSCNNNeuronELUNode")}
	})
	return _MPSCNNNeuronELUNodeClass
}

// GetMPSCNNNeuronELUNodeClass returns the class object for MPSCNNNeuronELUNode.
func GetMPSCNNNeuronELUNodeClass() MPSCNNNeuronELUNodeClass {
	return getMPSCNNNeuronELUNodeClass()
}

type MPSCNNNeuronELUNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronELUNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronELUNodeClass) Alloc() MPSCNNNeuronELUNode {
	rv := objc.Send[MPSCNNNeuronELUNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a parametric ELU neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronELUNode.InitWithSource]
//   - [MPSCNNNeuronELUNode.InitWithSourceA]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode
type MPSCNNNeuronELUNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronELUNodeFromID constructs a [MPSCNNNeuronELUNode] from an objc.ID.
//
// A representation of a parametric ELU neuron filter.
func MPSCNNNeuronELUNodeFromID(id objc.ID) MPSCNNNeuronELUNode {
	return MPSCNNNeuronELUNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronELUNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronELUNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronELUNode.InitWithSource]
//   - [IMPSCNNNeuronELUNode.InitWithSourceA]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode
type IMPSCNNNeuronELUNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronELUNode
	InitWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronELUNode
}

// Init initializes the instance.
func (c MPSCNNNeuronELUNode) Init() MPSCNNNeuronELUNode {
	rv := objc.Send[MPSCNNNeuronELUNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronELUNode) Autorelease() MPSCNNNeuronELUNode {
	rv := objc.Send[MPSCNNNeuronELUNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronELUNode creates a new MPSCNNNeuronELUNode instance.
func NewMPSCNNNeuronELUNode() MPSCNNNeuronELUNode {
	class := getMPSCNNNeuronELUNodeClass()
	rv := objc.Send[MPSCNNNeuronELUNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/init(source:)
func NewCNNNeuronELUNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronELUNode {
	instance := getMPSCNNNeuronELUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronELUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/init(source:a:)
func NewCNNNeuronELUNodeWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronELUNode {
	instance := getMPSCNNNeuronELUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	return MPSCNNNeuronELUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronELUNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronELUNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronELUNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronELUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/init(source:)
func (c MPSCNNNeuronELUNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronELUNode {
	rv := objc.Send[MPSCNNNeuronELUNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/init(source:a:)
func (c MPSCNNNeuronELUNode) InitWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronELUNode {
	rv := objc.Send[MPSCNNNeuronELUNode](c.ID, objc.Sel("initWithSource:a:"), sourceNode, a)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/nodeWithSource:
func (_MPSCNNNeuronELUNodeClass MPSCNNNeuronELUNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronELUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronELUNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronELUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELUNode/nodeWithSource:a:
func (_MPSCNNNeuronELUNodeClass MPSCNNNeuronELUNodeClass) NodeWithSourceA(sourceNode IMPSNNImageNode, a float32) MPSCNNNeuronELUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronELUNodeClass.class), objc.Sel("nodeWithSource:a:"), sourceNode, a)
	return MPSCNNNeuronELUNodeFromID(rv)
}
