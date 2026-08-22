// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronTanHNode] class.
var (
	_MPSCNNNeuronTanHNodeClass     MPSCNNNeuronTanHNodeClass
	_MPSCNNNeuronTanHNodeClassOnce sync.Once
)

func getMPSCNNNeuronTanHNodeClass() MPSCNNNeuronTanHNodeClass {
	_MPSCNNNeuronTanHNodeClassOnce.Do(func() {
		_MPSCNNNeuronTanHNodeClass = MPSCNNNeuronTanHNodeClass{class: objc.GetClass("MPSCNNNeuronTanHNode")}
	})
	return _MPSCNNNeuronTanHNodeClass
}

// GetMPSCNNNeuronTanHNodeClass returns the class object for MPSCNNNeuronTanHNode.
func GetMPSCNNNeuronTanHNodeClass() MPSCNNNeuronTanHNodeClass {
	return getMPSCNNNeuronTanHNodeClass()
}

type MPSCNNNeuronTanHNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronTanHNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronTanHNodeClass) Alloc() MPSCNNNeuronTanHNode {
	rv := objc.Send[MPSCNNNeuronTanHNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a hyperbolic tangent neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronTanHNode.InitWithSourceAB]
//   - [MPSCNNNeuronTanHNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode
type MPSCNNNeuronTanHNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronTanHNodeFromID constructs a [MPSCNNNeuronTanHNode] from an objc.ID.
//
// A representation of a hyperbolic tangent neuron filter.
func MPSCNNNeuronTanHNodeFromID(id objc.ID) MPSCNNNeuronTanHNode {
	return MPSCNNNeuronTanHNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronTanHNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronTanHNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronTanHNode.InitWithSourceAB]
//   - [IMPSCNNNeuronTanHNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode
type IMPSCNNNeuronTanHNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronTanHNode
	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronTanHNode
}

// Init initializes the instance.
func (c MPSCNNNeuronTanHNode) Init() MPSCNNNeuronTanHNode {
	rv := objc.Send[MPSCNNNeuronTanHNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronTanHNode) Autorelease() MPSCNNNeuronTanHNode {
	rv := objc.Send[MPSCNNNeuronTanHNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronTanHNode creates a new MPSCNNNeuronTanHNode instance.
func NewMPSCNNNeuronTanHNode() MPSCNNNeuronTanHNode {
	class := getMPSCNNNeuronTanHNodeClass()
	rv := objc.Send[MPSCNNNeuronTanHNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/init(source:)
func NewCNNNeuronTanHNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronTanHNode {
	instance := getMPSCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronTanHNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/init(source:a:b:)
func NewCNNNeuronTanHNodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronTanHNode {
	instance := getMPSCNNNeuronTanHNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronTanHNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronTanHNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronTanHNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronTanHNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronTanHNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/init(source:a:b:)
func (c MPSCNNNeuronTanHNode) InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronTanHNode {
	rv := objc.Send[MPSCNNNeuronTanHNode](c.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/init(source:)
func (c MPSCNNNeuronTanHNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronTanHNode {
	rv := objc.Send[MPSCNNNeuronTanHNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/nodeWithSource:
func (_MPSCNNNeuronTanHNodeClass MPSCNNNeuronTanHNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronTanHNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronTanHNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronTanHNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanHNode/nodeWithSource:a:b:
func (_MPSCNNNeuronTanHNodeClass MPSCNNNeuronTanHNodeClass) NodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronTanHNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronTanHNodeClass.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronTanHNodeFromID(rv)
}
