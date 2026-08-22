// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSoftSignNode] class.
var (
	_MPSCNNNeuronSoftSignNodeClass     MPSCNNNeuronSoftSignNodeClass
	_MPSCNNNeuronSoftSignNodeClassOnce sync.Once
)

func getMPSCNNNeuronSoftSignNodeClass() MPSCNNNeuronSoftSignNodeClass {
	_MPSCNNNeuronSoftSignNodeClassOnce.Do(func() {
		_MPSCNNNeuronSoftSignNodeClass = MPSCNNNeuronSoftSignNodeClass{class: objc.GetClass("MPSCNNNeuronSoftSignNode")}
	})
	return _MPSCNNNeuronSoftSignNodeClass
}

// GetMPSCNNNeuronSoftSignNodeClass returns the class object for MPSCNNNeuronSoftSignNode.
func GetMPSCNNNeuronSoftSignNodeClass() MPSCNNNeuronSoftSignNodeClass {
	return getMPSCNNNeuronSoftSignNodeClass()
}

type MPSCNNNeuronSoftSignNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSoftSignNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSoftSignNodeClass) Alloc() MPSCNNNeuronSoftSignNode {
	rv := objc.Send[MPSCNNNeuronSoftSignNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a softsign neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronSoftSignNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode
type MPSCNNNeuronSoftSignNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronSoftSignNodeFromID constructs a [MPSCNNNeuronSoftSignNode] from an objc.ID.
//
// A representation of a softsign neuron filter.
func MPSCNNNeuronSoftSignNodeFromID(id objc.ID) MPSCNNNeuronSoftSignNode {
	return MPSCNNNeuronSoftSignNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronSoftSignNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSoftSignNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronSoftSignNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode
type IMPSCNNNeuronSoftSignNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftSignNode
}

// Init initializes the instance.
func (c MPSCNNNeuronSoftSignNode) Init() MPSCNNNeuronSoftSignNode {
	rv := objc.Send[MPSCNNNeuronSoftSignNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSoftSignNode) Autorelease() MPSCNNNeuronSoftSignNode {
	rv := objc.Send[MPSCNNNeuronSoftSignNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSoftSignNode creates a new MPSCNNNeuronSoftSignNode instance.
func NewMPSCNNNeuronSoftSignNode() MPSCNNNeuronSoftSignNode {
	class := getMPSCNNNeuronSoftSignNodeClass()
	rv := objc.Send[MPSCNNNeuronSoftSignNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode/init(source:)
func NewCNNNeuronSoftSignNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftSignNode {
	instance := getMPSCNNNeuronSoftSignNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronSoftSignNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronSoftSignNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSoftSignNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronSoftSignNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronSoftSignNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode/init(source:)
func (c MPSCNNNeuronSoftSignNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftSignNode {
	rv := objc.Send[MPSCNNNeuronSoftSignNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSignNode/nodeWithSource:
func (_MPSCNNNeuronSoftSignNodeClass MPSCNNNeuronSoftSignNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronSoftSignNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronSoftSignNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronSoftSignNodeFromID(rv)
}
