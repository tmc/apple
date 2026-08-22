// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronReLUNNode] class.
var (
	_MPSCNNNeuronReLUNNodeClass     MPSCNNNeuronReLUNNodeClass
	_MPSCNNNeuronReLUNNodeClassOnce sync.Once
)

func getMPSCNNNeuronReLUNNodeClass() MPSCNNNeuronReLUNNodeClass {
	_MPSCNNNeuronReLUNNodeClassOnce.Do(func() {
		_MPSCNNNeuronReLUNNodeClass = MPSCNNNeuronReLUNNodeClass{class: objc.GetClass("MPSCNNNeuronReLUNNode")}
	})
	return _MPSCNNNeuronReLUNNodeClass
}

// GetMPSCNNNeuronReLUNNodeClass returns the class object for MPSCNNNeuronReLUNNode.
func GetMPSCNNNeuronReLUNNodeClass() MPSCNNNeuronReLUNNodeClass {
	return getMPSCNNNeuronReLUNNodeClass()
}

type MPSCNNNeuronReLUNNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronReLUNNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronReLUNNodeClass) Alloc() MPSCNNNeuronReLUNNode {
	rv := objc.Send[MPSCNNNeuronReLUNNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation a ReLUN neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronReLUNNode.InitWithSource]
//   - [MPSCNNNeuronReLUNNode.InitWithSourceAB]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode
type MPSCNNNeuronReLUNNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronReLUNNodeFromID constructs a [MPSCNNNeuronReLUNNode] from an objc.ID.
//
// A representation a ReLUN neuron filter.
func MPSCNNNeuronReLUNNodeFromID(id objc.ID) MPSCNNNeuronReLUNNode {
	return MPSCNNNeuronReLUNNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronReLUNNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronReLUNNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronReLUNNode.InitWithSource]
//   - [IMPSCNNNeuronReLUNNode.InitWithSourceAB]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode
type IMPSCNNNeuronReLUNNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNNode
	InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronReLUNNode
}

// Init initializes the instance.
func (c MPSCNNNeuronReLUNNode) Init() MPSCNNNeuronReLUNNode {
	rv := objc.Send[MPSCNNNeuronReLUNNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronReLUNNode) Autorelease() MPSCNNNeuronReLUNNode {
	rv := objc.Send[MPSCNNNeuronReLUNNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronReLUNNode creates a new MPSCNNNeuronReLUNNode instance.
func NewMPSCNNNeuronReLUNNode() MPSCNNNeuronReLUNNode {
	class := getMPSCNNNeuronReLUNNodeClass()
	rv := objc.Send[MPSCNNNeuronReLUNNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/init(source:)
func NewCNNNeuronReLUNNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNNode {
	instance := getMPSCNNNeuronReLUNNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronReLUNNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/init(source:a:b:)
func NewCNNNeuronReLUNNodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronReLUNNode {
	instance := getMPSCNNNeuronReLUNNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronReLUNNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronReLUNNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronReLUNNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronReLUNNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronReLUNNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/init(source:)
func (c MPSCNNNeuronReLUNNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNNode {
	rv := objc.Send[MPSCNNNeuronReLUNNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/init(source:a:b:)
func (c MPSCNNNeuronReLUNNode) InitWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronReLUNNode {
	rv := objc.Send[MPSCNNNeuronReLUNNode](c.ID, objc.Sel("initWithSource:a:b:"), sourceNode, a, b)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/nodeWithSource:
func (_MPSCNNNeuronReLUNNodeClass MPSCNNNeuronReLUNNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronReLUNNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronReLUNNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronReLUNNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUNNode/nodeWithSource:a:b:
func (_MPSCNNNeuronReLUNNodeClass MPSCNNNeuronReLUNNodeClass) NodeWithSourceAB(sourceNode IMPSNNImageNode, a float32, b float32) MPSCNNNeuronReLUNNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronReLUNNodeClass.class), objc.Sel("nodeWithSource:a:b:"), sourceNode, a, b)
	return MPSCNNNeuronReLUNNodeFromID(rv)
}
