// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronLogarithmNode] class.
var (
	_MPSCNNNeuronLogarithmNodeClass     MPSCNNNeuronLogarithmNodeClass
	_MPSCNNNeuronLogarithmNodeClassOnce sync.Once
)

func getMPSCNNNeuronLogarithmNodeClass() MPSCNNNeuronLogarithmNodeClass {
	_MPSCNNNeuronLogarithmNodeClassOnce.Do(func() {
		_MPSCNNNeuronLogarithmNodeClass = MPSCNNNeuronLogarithmNodeClass{class: objc.GetClass("MPSCNNNeuronLogarithmNode")}
	})
	return _MPSCNNNeuronLogarithmNodeClass
}

// GetMPSCNNNeuronLogarithmNodeClass returns the class object for MPSCNNNeuronLogarithmNode.
func GetMPSCNNNeuronLogarithmNodeClass() MPSCNNNeuronLogarithmNodeClass {
	return getMPSCNNNeuronLogarithmNodeClass()
}

type MPSCNNNeuronLogarithmNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronLogarithmNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronLogarithmNodeClass) Alloc() MPSCNNNeuronLogarithmNode {
	rv := objc.Send[MPSCNNNeuronLogarithmNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a logarithm neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronLogarithmNode.InitWithSource]
//   - [MPSCNNNeuronLogarithmNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode
type MPSCNNNeuronLogarithmNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronLogarithmNodeFromID constructs a [MPSCNNNeuronLogarithmNode] from an objc.ID.
//
// A representation of a logarithm neuron filter.
func MPSCNNNeuronLogarithmNodeFromID(id objc.ID) MPSCNNNeuronLogarithmNode {
	return MPSCNNNeuronLogarithmNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronLogarithmNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronLogarithmNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronLogarithmNode.InitWithSource]
//   - [IMPSCNNNeuronLogarithmNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode
type IMPSCNNNeuronLogarithmNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLogarithmNode
	InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronLogarithmNode
}

// Init initializes the instance.
func (c MPSCNNNeuronLogarithmNode) Init() MPSCNNNeuronLogarithmNode {
	rv := objc.Send[MPSCNNNeuronLogarithmNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronLogarithmNode) Autorelease() MPSCNNNeuronLogarithmNode {
	rv := objc.Send[MPSCNNNeuronLogarithmNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronLogarithmNode creates a new MPSCNNNeuronLogarithmNode instance.
func NewMPSCNNNeuronLogarithmNode() MPSCNNNeuronLogarithmNode {
	class := getMPSCNNNeuronLogarithmNodeClass()
	rv := objc.Send[MPSCNNNeuronLogarithmNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/init(source:)
func NewCNNNeuronLogarithmNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLogarithmNode {
	instance := getMPSCNNNeuronLogarithmNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronLogarithmNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/init(source:a:b:c:)
func NewCNNNeuronLogarithmNodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronLogarithmNode {
	instance := getMPSCNNNeuronLogarithmNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronLogarithmNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronLogarithmNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronLogarithmNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronLogarithmNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronLogarithmNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/init(source:)
func (c MPSCNNNeuronLogarithmNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLogarithmNode {
	rv := objc.Send[MPSCNNNeuronLogarithmNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/init(source:a:b:c:)
func (c_ MPSCNNNeuronLogarithmNode) InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronLogarithmNode {
	rv := objc.Send[MPSCNNNeuronLogarithmNode](c_.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/nodeWithSource:
func (_MPSCNNNeuronLogarithmNodeClass MPSCNNNeuronLogarithmNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronLogarithmNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronLogarithmNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronLogarithmNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithmNode/nodeWithSource:a:b:c:
func (_MPSCNNNeuronLogarithmNodeClass MPSCNNNeuronLogarithmNodeClass) NodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronLogarithmNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronLogarithmNodeClass.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronLogarithmNodeFromID(rv)
}
