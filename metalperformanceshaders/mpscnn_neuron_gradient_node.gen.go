// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronGradientNode] class.
var (
	_MPSCNNNeuronGradientNodeClass     MPSCNNNeuronGradientNodeClass
	_MPSCNNNeuronGradientNodeClassOnce sync.Once
)

func getMPSCNNNeuronGradientNodeClass() MPSCNNNeuronGradientNodeClass {
	_MPSCNNNeuronGradientNodeClassOnce.Do(func() {
		_MPSCNNNeuronGradientNodeClass = MPSCNNNeuronGradientNodeClass{class: objc.GetClass("MPSCNNNeuronGradientNode")}
	})
	return _MPSCNNNeuronGradientNodeClass
}

// GetMPSCNNNeuronGradientNodeClass returns the class object for MPSCNNNeuronGradientNode.
func GetMPSCNNNeuronGradientNodeClass() MPSCNNNeuronGradientNodeClass {
	return getMPSCNNNeuronGradientNodeClass()
}

type MPSCNNNeuronGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronGradientNodeClass) Alloc() MPSCNNNeuronGradientNode {
	rv := objc.Send[MPSCNNNeuronGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient exponential neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronGradientNode.InitWithSourceGradientSourceImageGradientStateDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNNeuronGradientNode.Descriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode
type MPSCNNNeuronGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNNeuronGradientNodeFromID constructs a [MPSCNNNeuronGradientNode] from an objc.ID.
//
// A representation of a gradient exponential neuron filter.
func MPSCNNNeuronGradientNodeFromID(id objc.ID) MPSCNNNeuronGradientNode {
	return MPSCNNNeuronGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNNeuronGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronGradientNode.InitWithSourceGradientSourceImageGradientStateDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNNeuronGradientNode.Descriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode
type IMPSCNNNeuronGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradientNode

	// Topic: Instance Properties

	Descriptor() IMPSNNNeuronDescriptor
}

// Init initializes the instance.
func (c MPSCNNNeuronGradientNode) Init() MPSCNNNeuronGradientNode {
	rv := objc.Send[MPSCNNNeuronGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronGradientNode) Autorelease() MPSCNNNeuronGradientNode {
	rv := objc.Send[MPSCNNNeuronGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronGradientNode creates a new MPSCNNNeuronGradientNode instance.
func NewMPSCNNNeuronGradientNode() MPSCNNNeuronGradientNode {
	class := getMPSCNNNeuronGradientNodeClass()
	rv := objc.Send[MPSCNNNeuronGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode/init(sourceGradient:sourceImage:gradientState:descriptor:)
func NewCNNNeuronGradientNodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradientNode {
	instance := getMPSCNNNeuronGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return MPSCNNNeuronGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode/init(sourceGradient:sourceImage:gradientState:descriptor:)
func (c MPSCNNNeuronGradientNode) InitWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradientNode {
	rv := objc.Send[MPSCNNNeuronGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode/nodeWithSourceGradient:sourceImage:gradientState:descriptor:
func (_MPSCNNNeuronGradientNodeClass MPSCNNNeuronGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNeuronGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return MPSCNNNeuronGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode/descriptor
func (c MPSCNNNeuronGradientNode) Descriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("descriptor"))
	return MPSNNNeuronDescriptorFromID(objc.ID(rv))
}
