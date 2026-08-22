// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronGeLUNode] class.
var (
	_MPSCNNNeuronGeLUNodeClass     MPSCNNNeuronGeLUNodeClass
	_MPSCNNNeuronGeLUNodeClassOnce sync.Once
)

func getMPSCNNNeuronGeLUNodeClass() MPSCNNNeuronGeLUNodeClass {
	_MPSCNNNeuronGeLUNodeClassOnce.Do(func() {
		_MPSCNNNeuronGeLUNodeClass = MPSCNNNeuronGeLUNodeClass{class: objc.GetClass("MPSCNNNeuronGeLUNode")}
	})
	return _MPSCNNNeuronGeLUNodeClass
}

// GetMPSCNNNeuronGeLUNodeClass returns the class object for MPSCNNNeuronGeLUNode.
func GetMPSCNNNeuronGeLUNodeClass() MPSCNNNeuronGeLUNodeClass {
	return getMPSCNNNeuronGeLUNodeClass()
}

type MPSCNNNeuronGeLUNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronGeLUNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronGeLUNodeClass) Alloc() MPSCNNNeuronGeLUNode {
	rv := objc.Send[MPSCNNNeuronGeLUNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNNeuronGeLUNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode
type MPSCNNNeuronGeLUNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronGeLUNodeFromID constructs a [MPSCNNNeuronGeLUNode] from an objc.ID.
func MPSCNNNeuronGeLUNodeFromID(id objc.ID) MPSCNNNeuronGeLUNode {
	return MPSCNNNeuronGeLUNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronGeLUNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronGeLUNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronGeLUNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode
type IMPSCNNNeuronGeLUNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronGeLUNode
}

// Init initializes the instance.
func (c MPSCNNNeuronGeLUNode) Init() MPSCNNNeuronGeLUNode {
	rv := objc.Send[MPSCNNNeuronGeLUNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronGeLUNode) Autorelease() MPSCNNNeuronGeLUNode {
	rv := objc.Send[MPSCNNNeuronGeLUNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronGeLUNode creates a new MPSCNNNeuronGeLUNode instance.
func NewMPSCNNNeuronGeLUNode() MPSCNNNeuronGeLUNode {
	class := getMPSCNNNeuronGeLUNodeClass()
	rv := objc.Send[MPSCNNNeuronGeLUNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode/init(source:)
func NewCNNNeuronGeLUNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronGeLUNode {
	instance := getMPSCNNNeuronGeLUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNeuronGeLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronGeLUNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGeLUNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronGeLUNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronGeLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode/init(source:)
func (c MPSCNNNeuronGeLUNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronGeLUNode {
	rv := objc.Send[MPSCNNNeuronGeLUNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGeLUNode/nodeWithSource:
func (_MPSCNNNeuronGeLUNodeClass MPSCNNNeuronGeLUNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNeuronGeLUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronGeLUNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNeuronGeLUNodeFromID(rv)
}
