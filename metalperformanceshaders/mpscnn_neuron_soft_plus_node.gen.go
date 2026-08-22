// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSoftPlusNode] class.
var (
	_MPSCNNNeuronSoftPlusNodeClass     MPSCNNNeuronSoftPlusNodeClass
	_MPSCNNNeuronSoftPlusNodeClassOnce sync.Once
)

func getMPSCNNNeuronSoftPlusNodeClass() MPSCNNNeuronSoftPlusNodeClass {
	_MPSCNNNeuronSoftPlusNodeClassOnce.Do(func() {
		_MPSCNNNeuronSoftPlusNodeClass = MPSCNNNeuronSoftPlusNodeClass{class: objc.GetClass("MPSCNNNeuronSoftPlusNode")}
	})
	return _MPSCNNNeuronSoftPlusNodeClass
}

// GetMPSCNNNeuronSoftPlusNodeClass returns the class object for MPSCNNNeuronSoftPlusNode.
func GetMPSCNNNeuronSoftPlusNodeClass() MPSCNNNeuronSoftPlusNodeClass {
	return getMPSCNNNeuronSoftPlusNodeClass()
}

type MPSCNNNeuronSoftPlusNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSoftPlusNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSoftPlusNodeClass) Alloc() MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a parametric softplus neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronSoftPlusNode.InitWithSourceAB]
//   - [MPSCNNNeuronSoftPlusNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode
type MPSCNNNeuronSoftPlusNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronSoftPlusNodeFromID constructs a [MPSCNNNeuronSoftPlusNode] from an objc.ID.
//
// A representation of a parametric softplus neuron filter.
func MPSCNNNeuronSoftPlusNodeFromID(id objc.ID) MPSCNNNeuronSoftPlusNode {
	return MPSCNNNeuronSoftPlusNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronSoftPlusNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSoftPlusNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronSoftPlusNode.InitWithSourceAB]
//   - [IMPSCNNNeuronSoftPlusNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode
type IMPSCNNNeuronSoftPlusNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronSoftPlusNode
	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftPlusNode
}

// Init initializes the instance.
func (c MPSCNNNeuronSoftPlusNode) Init() MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSoftPlusNode) Autorelease() MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSoftPlusNode creates a new MPSCNNNeuronSoftPlusNode instance.
func NewMPSCNNNeuronSoftPlusNode() MPSCNNNeuronSoftPlusNode {
	class := getMPSCNNNeuronSoftPlusNodeClass()
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/init(source:)
func NewCNNNeuronSoftPlusNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftPlusNode {
	instance := getMPSCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronSoftPlusNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/init(source:a:b:)
func NewCNNNeuronSoftPlusNodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronSoftPlusNode {
	instance := getMPSCNNNeuronSoftPlusNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronSoftPlusNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronSoftPlusNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronSoftPlusNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronSoftPlusNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/init(source:a:b:)
func (c MPSCNNNeuronSoftPlusNode) InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](c.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/init(source:)
func (c MPSCNNNeuronSoftPlusNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[MPSCNNNeuronSoftPlusNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/nodeWithSource:
func (_MPSCNNNeuronSoftPlusNodeClass MPSCNNNeuronSoftPlusNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronSoftPlusNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronSoftPlusNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlusNode/nodeWithSource:a:b:
func (_MPSCNNNeuronSoftPlusNodeClass MPSCNNNeuronSoftPlusNodeClass) NodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronSoftPlusNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronSoftPlusNodeClass.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronSoftPlusNodeFromID(rv)
}
