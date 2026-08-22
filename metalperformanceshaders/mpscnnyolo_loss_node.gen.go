// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNYOLOLossNode] class.
var (
	_MPSCNNYOLOLossNodeClass     MPSCNNYOLOLossNodeClass
	_MPSCNNYOLOLossNodeClassOnce sync.Once
)

func getMPSCNNYOLOLossNodeClass() MPSCNNYOLOLossNodeClass {
	_MPSCNNYOLOLossNodeClassOnce.Do(func() {
		_MPSCNNYOLOLossNodeClass = MPSCNNYOLOLossNodeClass{class: objc.GetClass("MPSCNNYOLOLossNode")}
	})
	return _MPSCNNYOLOLossNodeClass
}

// GetMPSCNNYOLOLossNodeClass returns the class object for MPSCNNYOLOLossNode.
func GetMPSCNNYOLOLossNodeClass() MPSCNNYOLOLossNodeClass {
	return getMPSCNNYOLOLossNodeClass()
}

type MPSCNNYOLOLossNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNYOLOLossNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNYOLOLossNodeClass) Alloc() MPSCNNYOLOLossNode {
	rv := objc.Send[MPSCNNYOLOLossNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a YOLO loss kernel.
//
// # Initializers
//
//   - [MPSCNNYOLOLossNode.InitWithSourceLossDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNYOLOLossNode.InputLabels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode
type MPSCNNYOLOLossNode struct {
	MPSNNFilterNode
}

// MPSCNNYOLOLossNodeFromID constructs a [MPSCNNYOLOLossNode] from an objc.ID.
//
// A representation of a YOLO loss kernel.
func MPSCNNYOLOLossNodeFromID(id objc.ID) MPSCNNYOLOLossNode {
	return MPSCNNYOLOLossNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNYOLOLossNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNYOLOLossNode] class.
//
// # Initializers
//
//   - [IMPSCNNYOLOLossNode.InitWithSourceLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNYOLOLossNode.InputLabels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode
type IMPSCNNYOLOLossNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLossNode

	// Topic: Instance Properties

	InputLabels() IMPSNNLabelsNode
}

// Init initializes the instance.
func (c MPSCNNYOLOLossNode) Init() MPSCNNYOLOLossNode {
	rv := objc.Send[MPSCNNYOLOLossNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNYOLOLossNode) Autorelease() MPSCNNYOLOLossNode {
	rv := objc.Send[MPSCNNYOLOLossNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNYOLOLossNode creates a new MPSCNNYOLOLossNode instance.
func NewMPSCNNYOLOLossNode() MPSCNNYOLOLossNode {
	class := getMPSCNNYOLOLossNodeClass()
	rv := objc.Send[MPSCNNYOLOLossNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode/init(source:lossDescriptor:)
func NewCNNYOLOLossNodeWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLossNode {
	instance := getMPSCNNYOLOLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	return MPSCNNYOLOLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode/init(source:lossDescriptor:)
func (c MPSCNNYOLOLossNode) InitWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLossNode {
	rv := objc.Send[MPSCNNYOLOLossNode](c.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode/nodeWithSource:lossDescriptor:
func (_MPSCNNYOLOLossNodeClass MPSCNNYOLOLossNodeClass) NodeWithSourceLossDescriptor(source IMPSNNImageNode, descriptor IMPSCNNYOLOLossDescriptor) MPSCNNYOLOLossNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNYOLOLossNodeClass.class), objc.Sel("nodeWithSource:lossDescriptor:"), source, descriptor)
	return MPSCNNYOLOLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode/inputLabels
func (c MPSCNNYOLOLossNode) InputLabels() IMPSNNLabelsNode {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputLabels"))
	return MPSNNLabelsNodeFromID(objc.ID(rv))
}
