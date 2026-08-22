// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronLinearNode] class.
var (
	_MPSCNNNeuronLinearNodeClass     MPSCNNNeuronLinearNodeClass
	_MPSCNNNeuronLinearNodeClassOnce sync.Once
)

func getMPSCNNNeuronLinearNodeClass() MPSCNNNeuronLinearNodeClass {
	_MPSCNNNeuronLinearNodeClassOnce.Do(func() {
		_MPSCNNNeuronLinearNodeClass = MPSCNNNeuronLinearNodeClass{class: objc.GetClass("MPSCNNNeuronLinearNode")}
	})
	return _MPSCNNNeuronLinearNodeClass
}

// GetMPSCNNNeuronLinearNodeClass returns the class object for MPSCNNNeuronLinearNode.
func GetMPSCNNNeuronLinearNodeClass() MPSCNNNeuronLinearNodeClass {
	return getMPSCNNNeuronLinearNodeClass()
}

type MPSCNNNeuronLinearNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronLinearNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronLinearNodeClass) Alloc() MPSCNNNeuronLinearNode {
	rv := objc.Send[MPSCNNNeuronLinearNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a linear neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronLinearNode.InitWithSourceAB]
//   - [MPSCNNNeuronLinearNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode
type MPSCNNNeuronLinearNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronLinearNodeFromID constructs a [MPSCNNNeuronLinearNode] from an objc.ID.
//
// A representation of a linear neuron filter.
func MPSCNNNeuronLinearNodeFromID(id objc.ID) MPSCNNNeuronLinearNode {
	return MPSCNNNeuronLinearNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronLinearNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronLinearNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronLinearNode.InitWithSourceAB]
//   - [IMPSCNNNeuronLinearNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode
type IMPSCNNNeuronLinearNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronLinearNode
	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLinearNode
}

// Init initializes the instance.
func (c MPSCNNNeuronLinearNode) Init() MPSCNNNeuronLinearNode {
	rv := objc.Send[MPSCNNNeuronLinearNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronLinearNode) Autorelease() MPSCNNNeuronLinearNode {
	rv := objc.Send[MPSCNNNeuronLinearNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronLinearNode creates a new MPSCNNNeuronLinearNode instance.
func NewMPSCNNNeuronLinearNode() MPSCNNNeuronLinearNode {
	class := getMPSCNNNeuronLinearNodeClass()
	rv := objc.Send[MPSCNNNeuronLinearNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/init(source:)
func NewCNNNeuronLinearNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLinearNode {
	instance := getMPSCNNNeuronLinearNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronLinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/init(source:a:b:)
func NewCNNNeuronLinearNodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronLinearNode {
	instance := getMPSCNNNeuronLinearNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronLinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronLinearNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronLinearNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronLinearNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronLinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/init(source:a:b:)
func (c MPSCNNNeuronLinearNode) InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronLinearNode {
	rv := objc.Send[MPSCNNNeuronLinearNode](c.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/init(source:)
func (c MPSCNNNeuronLinearNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLinearNode {
	rv := objc.Send[MPSCNNNeuronLinearNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/nodeWithSource:
func (_MPSCNNNeuronLinearNodeClass MPSCNNNeuronLinearNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLinearNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronLinearNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronLinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinearNode/nodeWithSource:a:b:
func (_MPSCNNNeuronLinearNodeClass MPSCNNNeuronLinearNodeClass) NodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronLinearNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronLinearNodeClass.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronLinearNodeFromID(rv)
}
