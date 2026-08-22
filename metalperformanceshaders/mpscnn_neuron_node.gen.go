// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronNode] class.
var (
	_MPSCNNNeuronNodeClass     MPSCNNNeuronNodeClass
	_MPSCNNNeuronNodeClassOnce sync.Once
)

func getMPSCNNNeuronNodeClass() MPSCNNNeuronNodeClass {
	_MPSCNNNeuronNodeClassOnce.Do(func() {
		_MPSCNNNeuronNodeClass = MPSCNNNeuronNodeClass{class: objc.GetClass("MPSCNNNeuronNode")}
	})
	return _MPSCNNNeuronNodeClass
}

// GetMPSCNNNeuronNodeClass returns the class object for MPSCNNNeuronNode.
func GetMPSCNNNeuronNodeClass() MPSCNNNeuronNodeClass {
	return getMPSCNNNeuronNodeClass()
}

type MPSCNNNeuronNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronNodeClass) Alloc() MPSCNNNeuronNode {
	rv := objc.Send[MPSCNNNeuronNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The virtual base class for MPS CNN neuron nodes.
//
// # Instance Properties
//
//   - [MPSCNNNeuronNode.A]
//   - [MPSCNNNeuronNode.B]
//   - [MPSCNNNeuronNode.C]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode
type MPSCNNNeuronNode struct {
	MPSNNFilterNode
}

// MPSCNNNeuronNodeFromID constructs a [MPSCNNNeuronNode] from an objc.ID.
//
// The virtual base class for MPS CNN neuron nodes.
func MPSCNNNeuronNodeFromID(id objc.ID) MPSCNNNeuronNode {
	return MPSCNNNeuronNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNNeuronNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronNode] class.
//
// # Instance Properties
//
//   - [IMPSCNNNeuronNode.A]
//   - [IMPSCNNNeuronNode.B]
//   - [IMPSCNNNeuronNode.C]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode
type IMPSCNNNeuronNode interface {
	IMPSNNFilterNode

	// Topic: Instance Properties

	A() float32
	B() float32
	C() float32
}

// Init initializes the instance.
func (c MPSCNNNeuronNode) Init() MPSCNNNeuronNode {
	rv := objc.Send[MPSCNNNeuronNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronNode) Autorelease() MPSCNNNeuronNode {
	rv := objc.Send[MPSCNNNeuronNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronNode creates a new MPSCNNNeuronNode instance.
func NewMPSCNNNeuronNode() MPSCNNNeuronNode {
	class := getMPSCNNNeuronNodeClass()
	rv := objc.Send[MPSCNNNeuronNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/init(source:descriptor:)
func NewCNNNeuronNodeWithSourceDescriptor(sourceNode IMPSNNImageNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronNode {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNNeuronNodeClass().class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return MPSCNNNeuronNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/a
func (c MPSCNNNeuronNode) A() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("a"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/b
func (c MPSCNNNeuronNode) B() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("b"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode/c
func (c MPSCNNNeuronNode) C() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("c"))
	return rv
}
