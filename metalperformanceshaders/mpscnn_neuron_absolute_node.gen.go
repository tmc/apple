// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronAbsoluteNode] class.
var (
	_MPSCNNNeuronAbsoluteNodeClass     MPSCNNNeuronAbsoluteNodeClass
	_MPSCNNNeuronAbsoluteNodeClassOnce sync.Once
)

func getMPSCNNNeuronAbsoluteNodeClass() MPSCNNNeuronAbsoluteNodeClass {
	_MPSCNNNeuronAbsoluteNodeClassOnce.Do(func() {
		_MPSCNNNeuronAbsoluteNodeClass = MPSCNNNeuronAbsoluteNodeClass{class: objc.GetClass("MPSCNNNeuronAbsoluteNode")}
	})
	return _MPSCNNNeuronAbsoluteNodeClass
}

// GetMPSCNNNeuronAbsoluteNodeClass returns the class object for MPSCNNNeuronAbsoluteNode.
func GetMPSCNNNeuronAbsoluteNodeClass() MPSCNNNeuronAbsoluteNodeClass {
	return getMPSCNNNeuronAbsoluteNodeClass()
}

type MPSCNNNeuronAbsoluteNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronAbsoluteNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronAbsoluteNodeClass) Alloc() MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[MPSCNNNeuronAbsoluteNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an absolute neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronAbsoluteNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode
type MPSCNNNeuronAbsoluteNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronAbsoluteNodeFromID constructs a [MPSCNNNeuronAbsoluteNode] from an objc.ID.
//
// A representation of an absolute neuron filter.
func MPSCNNNeuronAbsoluteNodeFromID(id objc.ID) MPSCNNNeuronAbsoluteNode {
	return MPSCNNNeuronAbsoluteNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronAbsoluteNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronAbsoluteNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronAbsoluteNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode
type IMPSCNNNeuronAbsoluteNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronAbsoluteNode
}

// Init initializes the instance.
func (c MPSCNNNeuronAbsoluteNode) Init() MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[MPSCNNNeuronAbsoluteNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronAbsoluteNode) Autorelease() MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[MPSCNNNeuronAbsoluteNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronAbsoluteNode creates a new MPSCNNNeuronAbsoluteNode instance.
func NewMPSCNNNeuronAbsoluteNode() MPSCNNNeuronAbsoluteNode {
	class := getMPSCNNNeuronAbsoluteNodeClass()
	rv := objc.Send[MPSCNNNeuronAbsoluteNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode/init(source:)
func NewCNNNeuronAbsoluteNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronAbsoluteNode {
	instance := getMPSCNNNeuronAbsoluteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronAbsoluteNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronAbsoluteNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronAbsoluteNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronAbsoluteNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode/init(source:)
func (c MPSCNNNeuronAbsoluteNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[MPSCNNNeuronAbsoluteNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsoluteNode/nodeWithSource:
func (_MPSCNNNeuronAbsoluteNodeClass MPSCNNNeuronAbsoluteNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronAbsoluteNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronAbsoluteNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronAbsoluteNodeFromID(rv)
}
