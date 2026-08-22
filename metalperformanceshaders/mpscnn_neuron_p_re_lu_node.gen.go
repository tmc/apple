// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronPReLUNode] class.
var (
	_MPSCNNNeuronPReLUNodeClass     MPSCNNNeuronPReLUNodeClass
	_MPSCNNNeuronPReLUNodeClassOnce sync.Once
)

func getMPSCNNNeuronPReLUNodeClass() MPSCNNNeuronPReLUNodeClass {
	_MPSCNNNeuronPReLUNodeClassOnce.Do(func() {
		_MPSCNNNeuronPReLUNodeClass = MPSCNNNeuronPReLUNodeClass{class: objc.GetClass("MPSCNNNeuronPReLUNode")}
	})
	return _MPSCNNNeuronPReLUNodeClass
}

// GetMPSCNNNeuronPReLUNodeClass returns the class object for MPSCNNNeuronPReLUNode.
func GetMPSCNNNeuronPReLUNodeClass() MPSCNNNeuronPReLUNodeClass {
	return getMPSCNNNeuronPReLUNodeClass()
}

type MPSCNNNeuronPReLUNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronPReLUNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronPReLUNodeClass) Alloc() MPSCNNNeuronPReLUNode {
	rv := objc.Send[MPSCNNNeuronPReLUNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation a PReLU neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronPReLUNode.InitWithSourceAData]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode
type MPSCNNNeuronPReLUNode struct {
	MPSCNNNeuronNode
}

// MPSCNNNeuronPReLUNodeFromID constructs a [MPSCNNNeuronPReLUNode] from an objc.ID.
//
// A representation a PReLU neuron filter.
func MPSCNNNeuronPReLUNodeFromID(id objc.ID) MPSCNNNeuronPReLUNode {
	return MPSCNNNeuronPReLUNode{MPSCNNNeuronNode: MPSCNNNeuronNodeFromID(id)}
}

// NOTE: MPSCNNNeuronPReLUNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronPReLUNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronPReLUNode.InitWithSourceAData]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode
type IMPSCNNNeuronPReLUNode interface {
	IMPSCNNNeuronNode

	// Topic: Initializers

	InitWithSourceAData(sourceNode IMPSNNImageNode, aData foundation.NSData) MPSCNNNeuronPReLUNode
}

// Init initializes the instance.
func (c MPSCNNNeuronPReLUNode) Init() MPSCNNNeuronPReLUNode {
	rv := objc.Send[MPSCNNNeuronPReLUNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronPReLUNode) Autorelease() MPSCNNNeuronPReLUNode {
	rv := objc.Send[MPSCNNNeuronPReLUNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronPReLUNode creates a new MPSCNNNeuronPReLUNode instance.
func NewMPSCNNNeuronPReLUNode() MPSCNNNeuronPReLUNode {
	class := getMPSCNNNeuronPReLUNodeClass()
	rv := objc.Send[MPSCNNNeuronPReLUNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode/init(source:aData:)
func NewCNNNeuronPReLUNodeWithSourceAData(sourceNode IMPSNNImageNode, aData foundation.NSData) MPSCNNNeuronPReLUNode {
	instance := getMPSCNNNeuronPReLUNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:aData:"), sourceNode, aData)
	return MPSCNNNeuronPReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronPReLUNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronPReLUNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronPReLUNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronPReLUNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode/init(source:aData:)
func (c MPSCNNNeuronPReLUNode) InitWithSourceAData(sourceNode IMPSNNImageNode, aData foundation.NSData) MPSCNNNeuronPReLUNode {
	rv := objc.Send[MPSCNNNeuronPReLUNode](c.ID, objc.Sel("initWithSource:aData:"), sourceNode, aData)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLUNode/nodeWithSource:aData:
func (_MPSCNNNeuronPReLUNodeClass MPSCNNNeuronPReLUNodeClass) NodeWithSourceAData(sourceNode IMPSNNImageNode, aData foundation.NSData) MPSCNNNeuronPReLUNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronPReLUNodeClass.class), objc.Sel("nodeWithSource:aData:"), sourceNode, aData)
	return MPSCNNNeuronPReLUNodeFromID(rv)
}
