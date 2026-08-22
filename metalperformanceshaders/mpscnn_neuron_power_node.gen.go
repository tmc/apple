// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronPowerNode] class.
var (
	_MPSCNNNeuronPowerNodeClass     MPSCNNNeuronPowerNodeClass
	_MPSCNNNeuronPowerNodeClassOnce sync.Once
)

func getMPSCNNNeuronPowerNodeClass() MPSCNNNeuronPowerNodeClass {
	_MPSCNNNeuronPowerNodeClassOnce.Do(func() {
		_MPSCNNNeuronPowerNodeClass = MPSCNNNeuronPowerNodeClass{class: objc.GetClass("MPSCNNNeuronPowerNode")}
	})
	return _MPSCNNNeuronPowerNodeClass
}

// GetMPSCNNNeuronPowerNodeClass returns the class object for MPSCNNNeuronPowerNode.
func GetMPSCNNNeuronPowerNodeClass() MPSCNNNeuronPowerNodeClass {
	return getMPSCNNNeuronPowerNodeClass()
}

type MPSCNNNeuronPowerNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronPowerNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronPowerNodeClass) Alloc() MPSCNNNeuronPowerNode {
	rv := objc.Send[MPSCNNNeuronPowerNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a power neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronPowerNode.InitWithSource]
//   - [MPSCNNNeuronPowerNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode
type MPSCNNNeuronPowerNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronPowerNodeFromID constructs a [MPSCNNNeuronPowerNode] from an objc.ID.
//
// A representation of a power neuron filter.
func MPSCNNNeuronPowerNodeFromID(id objc.ID) MPSCNNNeuronPowerNode {
	return MPSCNNNeuronPowerNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronPowerNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronPowerNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronPowerNode.InitWithSource]
//   - [IMPSCNNNeuronPowerNode.InitWithSourceABC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode
type IMPSCNNNeuronPowerNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronPowerNode
	InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronPowerNode
}

// Init initializes the instance.
func (c MPSCNNNeuronPowerNode) Init() MPSCNNNeuronPowerNode {
	rv := objc.Send[MPSCNNNeuronPowerNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronPowerNode) Autorelease() MPSCNNNeuronPowerNode {
	rv := objc.Send[MPSCNNNeuronPowerNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronPowerNode creates a new MPSCNNNeuronPowerNode instance.
func NewMPSCNNNeuronPowerNode() MPSCNNNeuronPowerNode {
	class := getMPSCNNNeuronPowerNodeClass()
	rv := objc.Send[MPSCNNNeuronPowerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/init(source:)
func NewCNNNeuronPowerNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronPowerNode {
	instance := getMPSCNNNeuronPowerNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronPowerNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/init(source:a:b:c:)
func NewCNNNeuronPowerNodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronPowerNode {
	instance := getMPSCNNNeuronPowerNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronPowerNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronPowerNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronPowerNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronPowerNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronPowerNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/init(source:)
func (c MPSCNNNeuronPowerNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronPowerNode {
	rv := objc.Send[MPSCNNNeuronPowerNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/init(source:a:b:c:)
func (c_ MPSCNNNeuronPowerNode) InitWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronPowerNode {
	rv := objc.Send[MPSCNNNeuronPowerNode](c_.ID, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/nodeWithSource:
func (_MPSCNNNeuronPowerNodeClass MPSCNNNeuronPowerNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronPowerNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronPowerNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronPowerNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPowerNode/nodeWithSource:a:b:c:
func (_MPSCNNNeuronPowerNodeClass MPSCNNNeuronPowerNodeClass) NodeWithSourceABC(sourceNode IMPSNNImageNode, a float32, b float32, c float32) MPSCNNNeuronPowerNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronPowerNodeClass.class), objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return MPSCNNNeuronPowerNodeFromID(rv)
}
