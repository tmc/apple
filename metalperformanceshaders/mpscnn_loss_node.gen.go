// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLossNode] class.
var (
	_MPSCNNLossNodeClass     MPSCNNLossNodeClass
	_MPSCNNLossNodeClassOnce sync.Once
)

func getMPSCNNLossNodeClass() MPSCNNLossNodeClass {
	_MPSCNNLossNodeClassOnce.Do(func() {
		_MPSCNNLossNodeClass = MPSCNNLossNodeClass{class: objc.GetClass("MPSCNNLossNode")}
	})
	return _MPSCNNLossNodeClass
}

// GetMPSCNNLossNodeClass returns the class object for MPSCNNLossNode.
func GetMPSCNNLossNodeClass() MPSCNNLossNodeClass {
	return getMPSCNNLossNodeClass()
}

type MPSCNNLossNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLossNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLossNodeClass) Alloc() MPSCNNLossNode {
	rv := objc.Send[MPSCNNLossNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a loss kernel.
//
// # Initializers
//
//   - [MPSCNNLossNode.InitWithSourceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNLossNode.InputLabels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode
type MPSCNNLossNode struct {
	MPSNNFilterNode
}

// MPSCNNLossNodeFromID constructs a [MPSCNNLossNode] from an objc.ID.
//
// A representation of a loss kernel.
func MPSCNNLossNodeFromID(id objc.ID) MPSCNNLossNode {
	return MPSCNNLossNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNLossNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLossNode] class.
//
// # Initializers
//
//   - [IMPSCNNLossNode.InitWithSourceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNLossNode.InputLabels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode
type IMPSCNNLossNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSCNNLossNode

	// Topic: Instance Properties

	InputLabels() IMPSNNLabelsNode
}

// Init initializes the instance.
func (c MPSCNNLossNode) Init() MPSCNNLossNode {
	rv := objc.Send[MPSCNNLossNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLossNode) Autorelease() MPSCNNLossNode {
	rv := objc.Send[MPSCNNLossNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLossNode creates a new MPSCNNLossNode instance.
func NewMPSCNNLossNode() MPSCNNLossNode {
	class := getMPSCNNLossNodeClass()
	rv := objc.Send[MPSCNNLossNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode/init(source:lossDescriptor:)
func NewCNNLossNodeWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSCNNLossNode {
	instance := getMPSCNNLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	return MPSCNNLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode/init(source:lossDescriptor:)
func (c MPSCNNLossNode) InitWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSCNNLossNode {
	rv := objc.Send[MPSCNNLossNode](c.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode/nodeWithSource:lossDescriptor:
func (_MPSCNNLossNodeClass MPSCNNLossNodeClass) NodeWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSCNNLossNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNLossNodeClass.class), objc.Sel("nodeWithSource:lossDescriptor:"), source, descriptor)
	return MPSCNNLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode/inputLabels
func (c MPSCNNLossNode) InputLabels() IMPSNNLabelsNode {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputLabels"))
	return MPSNNLabelsNodeFromID(objc.ID(rv))
}
