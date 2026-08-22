// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSigmoidNode] class.
var (
	_MPSCNNNeuronSigmoidNodeClass     MPSCNNNeuronSigmoidNodeClass
	_MPSCNNNeuronSigmoidNodeClassOnce sync.Once
)

func getMPSCNNNeuronSigmoidNodeClass() MPSCNNNeuronSigmoidNodeClass {
	_MPSCNNNeuronSigmoidNodeClassOnce.Do(func() {
		_MPSCNNNeuronSigmoidNodeClass = MPSCNNNeuronSigmoidNodeClass{class: objc.GetClass("MPSCNNNeuronSigmoidNode")}
	})
	return _MPSCNNNeuronSigmoidNodeClass
}

// GetMPSCNNNeuronSigmoidNodeClass returns the class object for MPSCNNNeuronSigmoidNode.
func GetMPSCNNNeuronSigmoidNodeClass() MPSCNNNeuronSigmoidNodeClass {
	return getMPSCNNNeuronSigmoidNodeClass()
}

type MPSCNNNeuronSigmoidNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSigmoidNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSigmoidNodeClass) Alloc() MPSCNNNeuronSigmoidNode {
	rv := objc.Send[MPSCNNNeuronSigmoidNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a sigmoid neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronSigmoidNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode
type MPSCNNNeuronSigmoidNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronSigmoidNodeFromID constructs a [MPSCNNNeuronSigmoidNode] from an objc.ID.
//
// A representation of a sigmoid neuron filter.
func MPSCNNNeuronSigmoidNodeFromID(id objc.ID) MPSCNNNeuronSigmoidNode {
	return MPSCNNNeuronSigmoidNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronSigmoidNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSigmoidNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronSigmoidNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode
type IMPSCNNNeuronSigmoidNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSigmoidNode
}

// Init initializes the instance.
func (c MPSCNNNeuronSigmoidNode) Init() MPSCNNNeuronSigmoidNode {
	rv := objc.Send[MPSCNNNeuronSigmoidNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSigmoidNode) Autorelease() MPSCNNNeuronSigmoidNode {
	rv := objc.Send[MPSCNNNeuronSigmoidNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSigmoidNode creates a new MPSCNNNeuronSigmoidNode instance.
func NewMPSCNNNeuronSigmoidNode() MPSCNNNeuronSigmoidNode {
	class := getMPSCNNNeuronSigmoidNodeClass()
	rv := objc.Send[MPSCNNNeuronSigmoidNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode/init(source:)
func NewCNNNeuronSigmoidNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSigmoidNode {
	instance := getMPSCNNNeuronSigmoidNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronSigmoidNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSigmoidNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronSigmoidNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronSigmoidNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode/init(source:)
func (c MPSCNNNeuronSigmoidNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSigmoidNode {
	rv := objc.Send[MPSCNNNeuronSigmoidNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoidNode/nodeWithSource:
func (_MPSCNNNeuronSigmoidNodeClass MPSCNNNeuronSigmoidNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSigmoidNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronSigmoidNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronSigmoidNodeFromID(rv)
}
