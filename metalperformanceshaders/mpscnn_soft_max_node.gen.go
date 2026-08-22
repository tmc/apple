// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSoftMaxNode] class.
var (
	_MPSCNNSoftMaxNodeClass     MPSCNNSoftMaxNodeClass
	_MPSCNNSoftMaxNodeClassOnce sync.Once
)

func getMPSCNNSoftMaxNodeClass() MPSCNNSoftMaxNodeClass {
	_MPSCNNSoftMaxNodeClassOnce.Do(func() {
		_MPSCNNSoftMaxNodeClass = MPSCNNSoftMaxNodeClass{class: objc.GetClass("MPSCNNSoftMaxNode")}
	})
	return _MPSCNNSoftMaxNodeClass
}

// GetMPSCNNSoftMaxNodeClass returns the class object for MPSCNNSoftMaxNode.
func GetMPSCNNSoftMaxNodeClass() MPSCNNSoftMaxNodeClass {
	return getMPSCNNSoftMaxNodeClass()
}

type MPSCNNSoftMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSoftMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSoftMaxNodeClass) Alloc() MPSCNNSoftMaxNode {
	rv := objc.Send[MPSCNNSoftMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a softmax filter.
//
// # Initializers
//
//   - [MPSCNNSoftMaxNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode
type MPSCNNSoftMaxNode struct {
	MPSNNFilterNode
}

// MPSCNNSoftMaxNodeFromID constructs a [MPSCNNSoftMaxNode] from an objc.ID.
//
// A representation of a softmax filter.
func MPSCNNSoftMaxNodeFromID(id objc.ID) MPSCNNSoftMaxNode {
	return MPSCNNSoftMaxNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNSoftMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSoftMaxNode] class.
//
// # Initializers
//
//   - [IMPSCNNSoftMaxNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode
type IMPSCNNSoftMaxNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNSoftMaxNode
}

// Init initializes the instance.
func (c MPSCNNSoftMaxNode) Init() MPSCNNSoftMaxNode {
	rv := objc.Send[MPSCNNSoftMaxNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSoftMaxNode) Autorelease() MPSCNNSoftMaxNode {
	rv := objc.Send[MPSCNNSoftMaxNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSoftMaxNode creates a new MPSCNNSoftMaxNode instance.
func NewMPSCNNSoftMaxNode() MPSCNNSoftMaxNode {
	class := getMPSCNNSoftMaxNodeClass()
	rv := objc.Send[MPSCNNSoftMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode/init(source:)
func NewCNNSoftMaxNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNSoftMaxNode {
	instance := getMPSCNNSoftMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNSoftMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode/init(source:)
func (c MPSCNNSoftMaxNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNSoftMaxNode {
	rv := objc.Send[MPSCNNSoftMaxNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxNode/nodeWithSource:
func (_MPSCNNSoftMaxNodeClass MPSCNNSoftMaxNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNSoftMaxNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNSoftMaxNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNSoftMaxNodeFromID(rv)
}
